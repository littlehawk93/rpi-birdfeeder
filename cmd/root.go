/*
Copyright © 2020 github.com/littlehawk93

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/littlehawk93/rpi-birdfeeder/conf"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootConfigFile string
var rootConfig *conf.ApplicationConfig

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rpi-birdfeeder",
	Short: "Main service and utility processes for running a Raspberry Pi powered bird-feeder",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize(onInitialize)
	rootCmd.PersistentFlags().StringVarP(&rootConfigFile, "config", "c", "", "A valid rpi-birdfeeder configuration file")
	rootCmd.MarkPersistentFlagRequired("config")
}

func onInitialize() {

	parseConfig()

	initializeLogger()
}

// parseConfig read the provided configuration file and unmarshal it into the global app configuration object
func parseConfig() {

	viper.SetConfigFile(rootConfigFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file '%s': %s\n", rootConfigFile, err.Error())
	}

	rootConfig = &conf.ApplicationConfig{}

	if err := viper.Unmarshal(rootConfig); err != nil {
		log.Fatalf("Error parsing config file '%s': %s\n", rootConfigFile, err.Error())
	}
}

// initializeLogger expects the rootAppConfig object to be populated. Initializes the global logger for the rpi-birdfeeder application
func initializeLogger() {

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	if rootConfig != nil && rootConfig.LogConfig != nil && rootConfig.LogConfig.File != "" {

		logFile, err := os.Create(rootConfig.LogConfig.File)

		if err != nil {
			log.Fatalf("Error opening log file '%s': %s\n", rootConfig.LogConfig.File, err.Error())
		}

		log.SetOutput(logFile)
	} else {
		log.SetOutput(os.Stderr)
	}
}
