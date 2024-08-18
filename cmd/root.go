package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/louisdutton/idasen/pkg/idasen"
)

var verbose bool
var configPath string
var log = logrus.New()
var watch bool
var height bool

var rootCmd = &cobra.Command{
	Version: "0.1.0",
	Use:     "idasen [height]",
	Short:   "Control your IKEA IDÅSEN desk from command line",
	Long:    "A command-line tool to control the height of your IKEA IDÅSEN desk via bluetooth.",
	Args:    cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			log.SetLevel(logrus.DebugLevel)
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		macAddress := viper.GetString("mac_address")
		timeout := viper.GetInt64("timeout")

		// if macAddress == "" {
		// 	d, err := idasen.GetDesk("", timeout)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	macAddress = d.Address.String()
		// 	viper.Set("mac_address", macAddress)
		// 	err = viper.WriteConfig()
		// 	fmt.Println("Configuration saved to", viper.ConfigFileUsed())
		// }

		desk, err := idasen.New(macAddress, timeout)
		defer desk.Close()
		if err != nil {
			return err
		}

		if watch {
			return desk.Monitor()
		} else if height {
			h, err := desk.Height()
			if err != nil {
				return err
			}

			fmt.Printf("%.4fm\n", h)
			return nil
		} else {
			targetHeight, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				return fmt.Errorf("Invalid target height \"%s\" given: %s\n", args[0], err)
			}

			return desk.SetHeight(targetHeight)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.Flags().BoolVarP(&watch, "watch", "w", false, "watch and display height changes")
	rootCmd.Flags().BoolVar(&height, "height", false, "display current height")
	rootCmd.Flags().StringVar(&configPath, "config", "", "config file (default is ~/.idasen.yaml)")
	viper.SetDefault("timeout", 10)
}

func initConfig() {
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".idasen")
		viper.SetEnvPrefix("idasen")
		viper.AutomaticEnv()
	}

	_ = viper.ReadInConfig()
}
