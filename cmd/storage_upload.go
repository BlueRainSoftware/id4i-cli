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
	"github.com/BlueRainSoftware/id4i-cli/api_models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	srcFile      string
	destFile     string
	shareDocWith []string
	publishDoc   bool
	contentType  string
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload new document",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Creating document ...")
		orga := viper.GetString("organization")

		// upload
		params := storage.NewCreateDocumentParams().
			WithOrganizationID(orga).
			WithID4N(globParamId4n)

		src, err := os.Open(srcFile)
		DieOnError(err)
		defer src.Close()

		if destFile != "" {
			dest, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, 0666)
			DieOnError(err)
			defer dest.Close()
			defer os.Remove(destFile)

			_, err = io.Copy(dest, src)
			DieOnError(err)
			params.SetContent(dest)
		} else {
			params.SetContent(src)
		}

		ok, accepted, err := ID4i.Storage.CreateDocument(params, Bearer())

		DieOnError(err)
		if accepted != nil {
			log.Info(accepted)
		}

		if ok != nil {
			log.Info("Document uploaded")
			OutputResult(ok.Payload)
		}

		// share / publish
		updateDocument(params)

	},
}

func updateDocument(docParams *storage.CreateDocumentParams) {

	vis :=  api_models.VisibilityUpdate{
		SharedWithOrganizationIds: shareDocWith,
		Public:                publishDoc,
	}

	documentUpdate := api_models.DocumentUpdate{
		MimeType: contentType,
		Visibility: &vis,
	}
	updateParams := storage.NewUpdateDocumentMetadataParams().
		WithOrganizationID(docParams.OrganizationID).
		WithID4N(docParams.ID4N).
		WithDocument(&documentUpdate).
		WithFileName(docParams.Content.Name())

	ok, noContent, err := ID4i.Storage.UpdateDocumentMetadata(updateParams, Bearer())

	if noContent != nil {
		log.Info(noContent)
	}

	if ok != nil {
		log.Info("Document visibility / content-type updated")
		OutputResult(ok.Payload)
	}

	DieOnError(err)

}

func init() {
	storageCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVarP(&srcFile, "file", "f", "", "Path to file to upload")
	uploadCmd.Flags().StringVarP(&destFile, "dest", "d", "", "Destination file name on ID4i")
	uploadCmd.Flags().StringVarP(&contentType, "content-type", "c", "application/octet-stream", "Content type for the destination file. Currently inactive due to #1532")
	uploadCmd.Flags().StringArrayVarP(&shareDocWith, "share-with", "s", []string{}, "Share document other organization(s). Repeat for sharing with multiple organizations.")
	uploadCmd.Flags().BoolVarP(&publishDoc, "publish", "p", false, "Publish document after uploading")
}
