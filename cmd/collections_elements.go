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
	"github.com/BlueRainSoftware/id4i-cli/api_client/collections"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var collectionElementsCmd = &cobra.Command{
	Use:   "elements",
	Short: "Retrieve elements (contained GUIDs) of the collection identified by --id",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Retrieving collection elements ...")

		params := collections.NewListElementsOfCollectionParams().
			WithID4N(globParamId4n).
			WithLimit(&globParamLimit).
			WithOffset(&globParamLimit)

		ok, _, err := ID4i.Collections.ListElementsOfCollection(params, Bearer())
		DieOnError(err)

		if ok != nil {
			log.Info("Collection contents retrieved ")
			OutputResult(ok.Payload)
		}
	},
}

func init() {
	collectionsCmd.AddCommand(collectionInfoCmd)

	collectionElementsCmd.MarkPersistentFlagRequired("id")
	collectionElementsCmd.MarkPersistentFlagRequired("limit")
	collectionElementsCmd.MarkPersistentFlagRequired("offset")

}
