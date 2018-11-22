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
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of an ongoing transfer",

	Run: func(cmd *cobra.Command, args []string) {
	params := transfer.NewGETSendInfoParams().
		WithID4N(globParamId4n)
		ok, _, err := ID4i.Transfer.GETSendInfo(params, Bearer())

		DieOnError(err)

		if ok != nil {
			log.Info("Transfer status retrieved")
			OutputResult(ok.Payload)
		}

	},
}

func init() {
	transferCmd.AddCommand(statusCmd)
}
