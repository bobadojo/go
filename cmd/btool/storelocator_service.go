// Code generated. DO NOT EDIT.

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/grpc"

	gapic "github.com/bobadojo/go/pkg/gapic"
)

var StoreLocatorConfig *viper.Viper
var StoreLocatorClient *gapic.StoreLocatorClient
var StoreLocatorSubCommands []string = []string{
	"find-stores",
	"get-store",
}

func init() {
	rootCmd.AddCommand(StoreLocatorServiceCmd)

	StoreLocatorConfig = viper.New()
	StoreLocatorConfig.SetEnvPrefix("CLI_STORELOCATOR")
	StoreLocatorConfig.AutomaticEnv()

	StoreLocatorServiceCmd.PersistentFlags().Bool("insecure", false, "Make insecure client connection. Or use CLI_STORELOCATOR_INSECURE. Must be used with \"address\" option")
	StoreLocatorConfig.BindPFlag("insecure", StoreLocatorServiceCmd.PersistentFlags().Lookup("insecure"))
	StoreLocatorConfig.BindEnv("insecure")

	StoreLocatorServiceCmd.PersistentFlags().String("address", "", "Set API address used by client. Or use CLI_STORELOCATOR_ADDRESS.")
	StoreLocatorConfig.BindPFlag("address", StoreLocatorServiceCmd.PersistentFlags().Lookup("address"))
	StoreLocatorConfig.BindEnv("address")

	StoreLocatorServiceCmd.PersistentFlags().String("token", "", "Set Bearer token used by the client. Or use CLI_STORELOCATOR_TOKEN.")
	StoreLocatorConfig.BindPFlag("token", StoreLocatorServiceCmd.PersistentFlags().Lookup("token"))
	StoreLocatorConfig.BindEnv("token")

	StoreLocatorServiceCmd.PersistentFlags().String("api_key", "", "Set API Key used by the client. Or use CLI_STORELOCATOR_API_KEY.")
	StoreLocatorConfig.BindPFlag("api_key", StoreLocatorServiceCmd.PersistentFlags().Lookup("api_key"))
	StoreLocatorConfig.BindEnv("api_key")
}

var StoreLocatorServiceCmd = &cobra.Command{
	Use:       "storelocator",
	Short:     "A store locator API",
	Long:      "A store locator API",
	ValidArgs: StoreLocatorSubCommands,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		var opts []option.ClientOption

		address := StoreLocatorConfig.GetString("address")
		if address != "" {
			opts = append(opts, option.WithEndpoint(address))
		}

		if StoreLocatorConfig.GetBool("insecure") {
			if address == "" {
				return fmt.Errorf("Missing address to use with insecure connection")
			}

			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				return err
			}
			opts = append(opts, option.WithGRPCConn(conn))
		}

		if token := StoreLocatorConfig.GetString("token"); token != "" {
			opts = append(opts, option.WithTokenSource(oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: token,
					TokenType:   "Bearer",
				})))
		}

		if key := StoreLocatorConfig.GetString("api_key"); key != "" {
			opts = append(opts, option.WithAPIKey(key))
		}

		StoreLocatorClient, err = gapic.NewStoreLocatorClient(ctx, opts...)
		return
	},
}
