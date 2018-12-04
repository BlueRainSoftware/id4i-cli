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
	"github.com/BlueRainSoftware/id4i-cli/api_client/transfer"
	"github.com/BlueRainSoftware/id4i-cli/api_models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive an ID transfer",
	Run: func(cmd *cobra.Command, args []string) {
		orga := viper.GetString("organization")

		params := transfer.NewReceiveParams().
			WithID4N(globParamId4n).
			WithRequest(&api_models.TransferReceiveInfo{
				OrganizationID: &orga,
			})
		ok, _, err := ID4i.Transfer.Receive(params, Bearer())
		DieOnError(err)

		if ok != nil {
			log.Info("Transfer retrieved")
		}
	},
}

func init() {
	transferCmd.AddCommand(receiveCmd)
	transferCmd.MarkPersistentFlagRequired("organization")

}
