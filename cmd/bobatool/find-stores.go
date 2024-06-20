// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	storespb "github.com/bobadojo/go/pkg/stores/v1/storespb"
)

var FindStoresInput storespb.FindStoresRequest

var FindStoresFromFile string

func init() {
	StoresServiceCmd.AddCommand(FindStoresCmd)

	FindStoresInput.Bounds = new(storespb.BoundingBox)

	FindStoresInput.Bounds.Max = new(storespb.Location)

	FindStoresInput.Bounds.Min = new(storespb.Location)

	FindStoresCmd.Flags().Float32Var(&FindStoresInput.Bounds.Max.Latitude, "bounds.max.latitude", 0.0, "Required. Latitude of the location.")

	FindStoresCmd.Flags().Float32Var(&FindStoresInput.Bounds.Max.Longitude, "bounds.max.longitude", 0.0, "Required. Longitude of the location.")

	FindStoresCmd.Flags().Float32Var(&FindStoresInput.Bounds.Min.Latitude, "bounds.min.latitude", 0.0, "Required. Latitude of the location.")

	FindStoresCmd.Flags().Float32Var(&FindStoresInput.Bounds.Min.Longitude, "bounds.min.longitude", 0.0, "Required. Longitude of the location.")

	FindStoresCmd.Flags().StringVar(&FindStoresFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var FindStoresCmd = &cobra.Command{
	Use:   "find-stores",
	Short: "Returns a list of all stores in a specified...",
	Long:  "Returns a list of all stores in a specified region.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if FindStoresFromFile == "" {

			cmd.MarkFlagRequired("bounds.max.latitude")

			cmd.MarkFlagRequired("bounds.max.longitude")

			cmd.MarkFlagRequired("bounds.min.latitude")

			cmd.MarkFlagRequired("bounds.min.longitude")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if FindStoresFromFile != "" {
			in, err = os.Open(FindStoresFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &FindStoresInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Stores", "FindStores", &FindStoresInput)
		}
		resp, err := StoresClient.FindStores(ctx, &FindStoresInput)
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
