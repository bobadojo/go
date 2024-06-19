// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	storespb "github.com/bobadojo/go/pkg/stores/v1/storespb"
)

var GetStoreInput storespb.GetStoreRequest

var GetStoreFromFile string

func init() {
	StoreLocatorServiceCmd.AddCommand(GetStoreCmd)

	GetStoreCmd.Flags().StringVar(&GetStoreInput.Name, "name", "", "Required. The ID of the store resource to retrieve.")

	GetStoreCmd.Flags().StringVar(&GetStoreFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var GetStoreCmd = &cobra.Command{
	Use:   "get-store",
	Short: "Returns a specific store.",
	Long:  "Returns a specific store.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if GetStoreFromFile == "" {

			cmd.MarkFlagRequired("name")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if GetStoreFromFile != "" {
			in, err = os.Open(GetStoreFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &GetStoreInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("StoreLocator", "GetStore", &GetStoreInput)
		}
		resp, err := StoreLocatorClient.GetStore(ctx, &GetStoreInput)
		if err != nil {
			return err
		}

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
