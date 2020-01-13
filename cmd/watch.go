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
	"log"

	"github.com/littlehawk93/rpi-birdfeeder/proc/watch"
	"github.com/spf13/cobra"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Capture videos or still images of birds landing on the bird feeder",
	Long:  "The main service for running the rpi-birdfeeder application. This command causes the Pi to continually listen for motion detected on the bird feeder and captures short videos or images of the birds when they land.",
	Run: func(cmd *cobra.Command, args []string) {

		if rootConfig == nil || rootConfig.WatchConfig == nil {
			log.Fatalln("No watch process configuration parameters provided")
		}

		watch.Run(rootConfig.WatchConfig)
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
