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
	"github.com/BlueRainSoftware/id4i-cli/api_models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	newCollectionGuidLength int32
	newCollectionLabel      string
	newCollectionType       string
)

var collectionCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new collection",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Creating collection ...")

		orga := viper.GetString("organization")
		validateCollectionType(newCollectionType)

		params := collections.NewCreateCollectionParams().
			WithCreateInfo(
				&api_models.CreateCollectionRequest{
					Label:          newCollectionLabel,
					Length:         &newCollectionGuidLength,
					OrganizationID: &orga,
					Type:           &newCollectionType,
				})

		ok, created, accepted, err := ID4i.Collections.CreateCollection(params, Bearer())
		DieOnError(err)

		// TODO check whether created & accepted really occurs
		log.Info(created)
		log.Info(accepted)

		if ok != nil {
			log.Info("Collection created.")
			OutputResult(ok.Payload)
		}
	},
}

func init() {
	collectionsCmd.AddCommand(collectionCreateCmd)

	collectionCreateCmd.Flags().StringVarP(&newCollectionType, "type", "t", "", "Collection type, ROUTING_COLLECTION, LOGISTIC_COLLECTION or LABELLED_COLLECTION")
	collectionCreateCmd.MarkFlagRequired("type")

	collectionCreateCmd.Flags().StringVarP(&newCollectionLabel, "label", "a", "", "Collection label")
	collectionCreateCmd.MarkFlagRequired("label")

	collectionCreateCmd.Flags().Int32VarP(&newCollectionGuidLength, "length", "l", 10, "Collection ID length")
}
