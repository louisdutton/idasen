package idasen

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func (i *Idasen) Monitor() error {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)

	go func() {
		var previousHeight = 0.0

		for true {
			h, err := i.Height()
			if err != nil {
				fmt.Println(err)
				return
			}

			if h != previousHeight {
				fmt.Printf("%.4fm\n", h)
				previousHeight = h
			}

			select {
			case <-ch:
				wg.Done()
				return
			default:

			}
		}
	}()

	wg.Wait()
	return nil
}
