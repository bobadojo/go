// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	menuspb "github.com/bobadojo/go/pkg/menus/v1/menuspb"

	"os"
)

var GetMenuInput menuspb.GetMenuRequest

var GetMenuFromFile string

func init() {
	MenusServiceCmd.AddCommand(GetMenuCmd)

	GetMenuCmd.Flags().StringVar(&GetMenuInput.Name, "name", "", "Required. The ID of the menu resource to retrieve.")

	GetMenuCmd.Flags().StringVar(&GetMenuFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var GetMenuCmd = &cobra.Command{
	Use:   "get-menu",
	Short: "Returns a menu for a specific store.",
	Long:  "Returns a menu for a specific store.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if GetMenuFromFile == "" {

			cmd.MarkFlagRequired("name")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if GetMenuFromFile != "" {
			in, err = os.Open(GetMenuFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &GetMenuInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Menus", "GetMenu", &GetMenuInput)
		}
		resp, err := MenusClient.GetMenu(ctx, &GetMenuInput)
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
