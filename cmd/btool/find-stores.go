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
	StoreLocatorServiceCmd.AddCommand(FindStoresCmd)

	FindStoresInput.BoundingBox = new(storespb.Box)

	FindStoresInput.BoundingBox.Max = new(storespb.Point)

	FindStoresInput.BoundingBox.Min = new(storespb.Point)

	FindStoresInput.Center = new(storespb.Point)

	FindStoresCmd.Flags().Float32Var(&FindStoresInput.BoundingBox.Max.Latitude, "bounding_box.max.latitude", 0.0, "Latitude of the point.")

	FindStoresCmd.Flags().Float32Var(&FindStoresInput.BoundingBox.Max.Longitude, "bounding_box.max.longitude", 0.0, "Longitude of the point.")

	FindStoresCmd.Flags().Float32Var(&FindStoresInput.BoundingBox.Min.Latitude, "bounding_box.min.latitude", 0.0, "Latitude of the point.")

	FindStoresCmd.Flags().Float32Var(&FindStoresInput.BoundingBox.Min.Longitude, "bounding_box.min.longitude", 0.0, "Longitude of the point.")

	FindStoresCmd.Flags().Float32Var(&FindStoresInput.Center.Latitude, "center.latitude", 0.0, "Latitude of the point.")

	FindStoresCmd.Flags().Float32Var(&FindStoresInput.Center.Longitude, "center.longitude", 0.0, "Longitude of the point.")

	FindStoresCmd.Flags().StringVar(&FindStoresFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var FindStoresCmd = &cobra.Command{
	Use:   "find-stores",
	Short: "Returns a list of all stores in the store.",
	Long:  "Returns a list of all stores in the store.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if FindStoresFromFile == "" {

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
			printVerboseInput("StoreLocator", "FindStores", &FindStoresInput)
		}
		resp, err := StoreLocatorClient.FindStores(ctx, &FindStoresInput)
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
