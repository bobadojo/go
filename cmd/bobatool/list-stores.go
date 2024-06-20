// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"google.golang.org/api/iterator"

	"os"

	storespb "github.com/bobadojo/go/pkg/stores/v1/storespb"
)

var ListStoresInput storespb.ListStoresRequest

var ListStoresFromFile string

func init() {
	StoresServiceCmd.AddCommand(ListStoresCmd)

	ListStoresCmd.Flags().Int32Var(&ListStoresInput.PageSize, "page_size", 10, "Default is 10. Page size.")

	ListStoresCmd.Flags().StringVar(&ListStoresInput.PageToken, "page_token", "", "Pagination token.")

	ListStoresCmd.Flags().StringVar(&ListStoresFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var ListStoresCmd = &cobra.Command{
	Use:   "list-stores",
	Short: "List all stores.",
	Long:  "List all stores.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if ListStoresFromFile == "" {

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if ListStoresFromFile != "" {
			in, err = os.Open(ListStoresFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &ListStoresInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Stores", "ListStores", &ListStoresInput)
		}
		iter := StoresClient.ListStores(ctx, &ListStoresInput)

		// populate iterator with a page
		_, err = iter.Next()
		if err != nil && err != iterator.Done {
			return err
		}

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(iter.Response)

		return err
	},
}
