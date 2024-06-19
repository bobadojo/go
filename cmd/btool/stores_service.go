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

var StoresConfig *viper.Viper
var StoresClient *gapic.StoresClient
var StoresSubCommands []string = []string{
	"find-stores",
	"get-store",
}

func init() {
	rootCmd.AddCommand(StoresServiceCmd)

	StoresConfig = viper.New()
	StoresConfig.SetEnvPrefix("CLI_STORES")
	StoresConfig.AutomaticEnv()

	StoresServiceCmd.PersistentFlags().Bool("insecure", false, "Make insecure client connection. Or use CLI_STORES_INSECURE. Must be used with \"address\" option")
	StoresConfig.BindPFlag("insecure", StoresServiceCmd.PersistentFlags().Lookup("insecure"))
	StoresConfig.BindEnv("insecure")

	StoresServiceCmd.PersistentFlags().String("address", "", "Set API address used by client. Or use CLI_STORES_ADDRESS.")
	StoresConfig.BindPFlag("address", StoresServiceCmd.PersistentFlags().Lookup("address"))
	StoresConfig.BindEnv("address")

	StoresServiceCmd.PersistentFlags().String("token", "", "Set Bearer token used by the client. Or use CLI_STORES_TOKEN.")
	StoresConfig.BindPFlag("token", StoresServiceCmd.PersistentFlags().Lookup("token"))
	StoresConfig.BindEnv("token")

	StoresServiceCmd.PersistentFlags().String("api_key", "", "Set API Key used by the client. Or use CLI_STORES_API_KEY.")
	StoresConfig.BindPFlag("api_key", StoresServiceCmd.PersistentFlags().Lookup("api_key"))
	StoresConfig.BindEnv("api_key")
}

var StoresServiceCmd = &cobra.Command{
	Use:       "stores",
	Short:     "A store locator API",
	Long:      "A store locator API",
	ValidArgs: StoresSubCommands,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		var opts []option.ClientOption

		address := StoresConfig.GetString("address")
		if address != "" {
			opts = append(opts, option.WithEndpoint(address))
		}

		if StoresConfig.GetBool("insecure") {
			if address == "" {
				return fmt.Errorf("Missing address to use with insecure connection")
			}

			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				return err
			}
			opts = append(opts, option.WithGRPCConn(conn))
		}

		if token := StoresConfig.GetString("token"); token != "" {
			opts = append(opts, option.WithTokenSource(oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: token,
					TokenType:   "Bearer",
				})))
		}

		if key := StoresConfig.GetString("api_key"); key != "" {
			opts = append(opts, option.WithAPIKey(key))
		}

		StoresClient, err = gapic.NewStoresClient(ctx, opts...)
		return
	},
}
