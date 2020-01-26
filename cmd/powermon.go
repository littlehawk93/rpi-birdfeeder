/*
Copyright © 2020 NAME HERE github.com/littlehawk93

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/littlehawk93/rpi-birdfeeder/proc/powermon"
	"github.com/spf13/cobra"
)

// powermonCmd represents the powermon command
var powermonCmd = &cobra.Command{
	Use:   "powermon",
	Short: "Monitors power consumption and battery level and records data to InfluxDB",
	Long:  `Using the INA219 sensor, this process checks battery voltage / power levels and current power consumption by the pi on regular intervals and reports data to an InfluxDB instance`,
	Run:   buildProcess(powermon.Run),
}

func init() {
	rootCmd.AddCommand(powermonCmd)
}
