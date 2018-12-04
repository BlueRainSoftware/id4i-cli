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
	"strings"
)

var collectionTypes = strings.Join([]string{
	"ROUTING_COLLECTION", "LOGISTIC_COLLECTION", "LABELLED_COLLECTION"}, " ")

var collectionsCmd = &cobra.Command{
	Use:   "collections",
	Short: "Operations on labelled, logistic and routing collections",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(collectionsCmd)
	collectionsCmd.PersistentFlags().StringVarP(&globParamId4n, "id", "i", "", "ID4i ID of Collection to operate on")
}

func validateCollectionType(t string) {
	if ! strings.Contains(collectionTypes, t) {
		log.WithFields(log.Fields{"type": t}).Fatal("Unknown history item type used for filtering")
	}
}
