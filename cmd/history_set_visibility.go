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
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	newShareWith []string
	newPublic    bool
	sequence	int32
)
// setVisibilityCmd represents the setVisibility command
var setVisibilityCmd = &cobra.Command{
	Use:   "set-visibility",
	Short: "Update history item visibillity",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Update history item visibility ...")

		orga := viper.GetString("organization")
		vis := api_models.Visibility{
			SharedOrganizationIds: newShareWith,
			Public:                newPublic,
		}

		params := history.NewUpdateItemVisibilityParams().
			WithID4N(globParamId4n).
			WithOrganizationID(orga).
			WithVisibility(&vis).
			WithSequenceID(sequence)

		ok, _, err := ID4i.History.UpdateItemVisibility(params, Bearer())
		if err != nil {
			OutputError(err)
		}

		if ok != nil {
			log.Info("History item visibility updated ")
			OutputResult(ok)
		}
	},
}

func init() {
	historyCmd.AddCommand(setVisibilityCmd)

	setVisibilityCmd.Flags().StringArrayVarP(&newShareWith, "share-with", "s", []string{}, "Share with other organization(s). Repeat for sharing with multiple organizations.")
	setVisibilityCmd.Flags().BoolVarP(&newPublic, "public", "p", false, "Make history item public")
	setVisibilityCmd.Flags().Int32VarP(&sequence, "sequence", "", -1, "Sequence no of the history item")

}
