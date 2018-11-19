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
	"github.com/BlueRainSoftware/id4i-cli/api_client/history"
	"github.com/BlueRainSoftware/id4i-cli/api_models"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var (
	historyType string
	shareWith   []string
	public bool
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add history item",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Creating history item ...")

		item := api_models.HistoryItem{
			OrganizationID: &globParamOrganization,
			Type:           &historyType,
			Visibility: &api_models.Visibility{
				SharedOrganizationIds: shareWith,
				Public: public,
			},
		}

		err := item.Validate(nil)
		if err != nil {
			OutputValidationError(err)
		}

		params := history.NewAddItemParams().
			WithID4N(globParamId4n).WithHistoryItem(&item)

		ok, created, accepted, err := ID4i.History.AddItem(params, Bearer())

		if err != nil {
			OutputError(err)
		}
		if accepted != nil {
			log.Info(accepted)
		}
		if created != nil {
			log.Info(created)
		}
		if ok != nil {
			log.Info("History item created")
			log.Info(ok)
			OutputResult(item)
		}

	},
}

func init() {
	historyCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&historyType, "type", "t", "", "History item type")
	addCmd.MarkFlagRequired("type")

	addCmd.Flags().StringArrayVarP(&shareWith, "shareWith", "s", []string{}, "Share with other organization(s). Repeat for sharing with multiple organizations")

	addCmd.Flags().BoolVarP(&public, "public", "p", false, "Make history item public")
}
