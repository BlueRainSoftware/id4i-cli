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
	"github.com/BlueRainSoftware/id4i-cli/api_client/storage"
	"github.com/go-openapi/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

var (
	srcFile      string
	destFile     string
	shareDocWith []string
	publishDoc   bool
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload new document",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Creating document ...")
		orga := viper.GetString("organization")

		f, err := os.Open(srcFile)
		if err != nil {
			log.Fatal(err)
		}
		runtime.NamedReader("test", f)
		params := storage.NewCreateDocumentParams().
			WithOrganizationID(orga).
			WithID4N(globParamId4n).
			WithContent(f)

		ok, accepted, err := ID4i.Storage.CreateDocument(params, Bearer())

		if err != nil {
			OutputError(err)
		}
		if accepted != nil {
			log.Info(accepted)
		}

		if ok != nil {
			log.Info("Document uploaded")
			OutputResult(ok)
		}
	},
}

func init() {
	storageCmd.AddCommand(uploadCmd)
	listCmd.Flags().StringVarP(&srcFile, "file", "f", "", "Path to file to upload")
	listCmd.Flags().StringVarP(&destFile, "dest", "d", "", "Destination file name on ID4i")
	addCmd.Flags().StringArrayVarP(&shareDocWith, "share-with", "s", []string{}, "Share document other organization(s). Repeat for sharing with multiple organizations.")
	addCmd.Flags().BoolVarP(&publishDoc, "public", "p", false, "Publish document after uploading")
}
