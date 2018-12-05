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
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

var (
	historyTypes   []string
	fromDate       string
	toDate         string
	qualifier      []string
	includePrivate bool
	organization   string
)

// copied from api_models/history_item.go:81
var historyItemTypeAllowedValues = strings.Join([]string{
	"CREATED", "DESTROYED", "RECYCLED", "SHIPMENT_PREPARED", "STORED", "RETRIEVED_FROM_STORAGE", "PACKAGED",
	"DISPATCHED", "RECEIVED", "REPROCESSING_STARTED", "REPROCESSING_STEP_STARTED", "REPROCESSING_STEP_CANCELLED",
	"REPROCESSING_STEP_FINISHED", "REPROCESSING_CANCELLED", "REPROCESSING_FINISHED", "DISASSEMBLED",
	"MAINTENANCE_STARTED", "MAINTENANCE_STEP_STARTED", "MAINTENANCE_STEP_CANCELLED", "MAINTENANCE_STEP_FINISHED",
	"MAINTENANCE_CANCELLED", "MAINTENANCE_FINISHED", "PRODUCTION_STARTED", "PRODUCTION_CANCELLED",
	"PRODUCTION_FINISHED", "PRODUCTION_STEP_STARTED", "PRODUCTION_STEP_CANCELLED", "PRODUCTION_STEP_FINISHED",
	"QUALITY_CHECK_PERFORMED"}, " ")

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List ID history",
	Run: func(cmd *cobra.Command, args []string) {

		validateTypeParams()
		var start, end = parseDurationParams()
		log.Info("Retrieving history ...")

		filter := history.NewFilteredListParams().
			WithID4N(globParamId4n).
			WithFromDate(start).
			WithToDate(end).
			WithType(historyTypes).
			WithQualifier(qualifier).
			WithIncludePrivate(&includePrivate).
			WithOrganization(&organization).
			WithOffset(&globParamOffset).
			WithLimit(&globParamLimit)

		ok, _, err := ID4i.History.FilteredList(filter, Bearer())
		DieOnError(err)

		if ok != nil {
			log.Info("History retrieved")
			OutputResult(ok.Payload)
		}

	},
}

func validateTypeParams() {
	for _, t := range historyTypes {
		if ! strings.Contains(historyItemTypeAllowedValues, t) {
			log.WithFields(log.Fields{"type": t}).Fatal("Unknown history item type used for filtering")
		}
	}
}

func parseDurationParams() (*strfmt.DateTime, *strfmt.DateTime) {
	start, startErr := strfmt.ParseDateTime(fromDate)
	if startErr != nil {
		log.WithField("start", fromDate).Fatal("Start date could not be parsed using ISO8601 format (e.g. 2006-01-02T15:04:05.000Z07:00)")
	}
	end, endErr := strfmt.ParseDateTime(toDate)
	if endErr != nil {
		log.WithField("end", toDate).Fatal("End date could not be parsed using ISO8601 format (e.g. 2006-01-02T15:04:05.000Z07:00)")
	}

	return &start, &end
}

func init() {
	historyCmd.AddCommand(listCmd)

	listCmd.Flags().StringArrayVarP(&historyTypes, "type", "t", []string{}, "History item type. Repeat to filter multiple types.")
	listCmd.Flags().StringArrayVarP(&qualifier, "qualifier", "q", []string{}, "Filter by qualifier. Repeat to filter multiple qualifiers.")
	listCmd.Flags().StringVarP(&organization, "creator", "", "", "Filter by creator organization")
	listCmd.Flags().BoolVarP(&includePrivate, "include-private", "p", false, "Show private items")
	listCmd.Flags().StringVarP(&fromDate, "start", "s", "2017-01-01T00:00:42Z", "Start date in ISO8601 format (e.g. 2006-01-02T15:04:05.000Z07:00)")
	listCmd.Flags().StringVarP(&toDate, "end", "e", time.Now().UTC().Format(time.RFC3339), "End date in ISO8601 format (e.g. 2006-01-02T15:04:05.000Z07:00)")

	listCmd.MarkPersistentFlagRequired("id")
	listCmd.MarkPersistentFlagRequired("limit")
	listCmd.MarkPersistentFlagRequired("offset")

}
