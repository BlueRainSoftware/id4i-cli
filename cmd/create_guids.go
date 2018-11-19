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
	"github.com/BlueRainSoftware/id4i-cli/api_client/guids"
	"github.com/BlueRainSoftware/id4i-cli/api_models"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"
)

var (
	count  int32
	length int32
)

var createGuidsCmd = &cobra.Command{
	Use:   "create",
	Short: "Create GUIDs",
	Run: func(cmd *cobra.Command, args []string) {

		gr := api_models.CreateGUIDRequest{Count: &count, Length: &length, OrganizationID: &globCfgOrganization}

		params := guids.NewCreateGUIDParams().WithCreateGUIDInfo(&gr)
		log.Info("Creating GUIDs ...")
		resp, accepted, created, err := ID4i.Guids.CreateGUID(params, Bearer())
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if accepted != nil {
			log.Info("Request handled asynchronously")
			log.Info(accepted)
		}
		if created != nil {
			log.Info("GUIDs created")
			log.Info(created)
			log.Info("%#v\n", resp.Payload)
		}
	},
}

func init() {
	guidsCmd.AddCommand(createGuidsCmd)

	createGuidsCmd.Flags().Int32VarP(&count, "count", "c", 1, "Number of GUIDs to created between 1 and 1000")
	createGuidsCmd.MarkFlagRequired("count")
	createGuidsCmd.Flags().Int32VarP(&length, "length", "l", 7, "GUID length between 7 and 255")
	createGuidsCmd.MarkFlagRequired("length")
}
