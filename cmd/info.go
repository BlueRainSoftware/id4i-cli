// Copyright © 2018 Wolfgang Werner <mail@wolfgang-werner.net>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgShowConfig      bool
	cfgShowBackendInfo bool
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Configuration and backend API information",
	Run: func(cmd *cobra.Command, args []string) {

		if cfgShowConfig {
			log.Info("Configuration")
			log.WithField("file", viper.ConfigFileUsed()).Info()

			for _, element := range viper.AllKeys() {
				switch element {
				case "secret":
					log.WithField(element, "*****").Info()
				default:
					log.WithField(element, viper.Get(element)).Info()
				}
			}
		}

		if cfgShowBackendInfo {
			log.Info("Fetching backend information")

			resp, accepted, err := ID4i.MetaInformation.ApplicationInfo(nil, Bearer())
			DieOnError(err)

			if accepted != nil {
				log.Info(accepted)
			}
			OutputResult(resp.Payload)

		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	infoCmd.Flags().BoolVarP(&cfgShowConfig, "show-config", "", true, "Show configuration")
	infoCmd.Flags().BoolVarP(&cfgShowBackendInfo, "show-backend", "", true, "Show backend information")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
