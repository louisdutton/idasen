package idasen

import (
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"time"
	"tinygo.org/x/bluetooth"
)

var deskNameRegex = regexp.MustCompile("Desk \\d+")

func GetDesk(mac string, timeout int64) (*bluetooth.ScanResult, error) {
	err := Adapter.Enable()
	if err != nil {
		return nil, fmt.Errorf("must enable adapter: %w", err)
	}

	resultCh := make(chan bluetooth.ScanResult, 1)

	// Start scanning.
	err = Adapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		if mac == "" {
			if deskNameRegex.MatchString(result.LocalName()) {
				adapter.StopScan()
				resultCh <- result
			}
		} else {
			if result.Address.String() == mac {
				adapter.StopScan()
				resultCh <- result
			}
		}

	})

	discoverTimeout := time.NewTimer(time.Duration(timeout) * time.Second)
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)

	select {
	case result := <-resultCh:
		return &result, nil
	case <-discoverTimeout.C:
		e := fmt.Errorf("Discover timeout reached after %d seconds", timeout)
		return nil, e
	case sig := <-ch:
		e := fmt.Errorf("Received signal [%v]; shutting down...", sig)
		return nil, e
	}
}
