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

var MenusConfig *viper.Viper
var MenusClient *gapic.MenusClient
var MenusSubCommands []string = []string{
	"get-menu",
}

func init() {
	rootCmd.AddCommand(MenusServiceCmd)

	MenusConfig = viper.New()
	MenusConfig.SetEnvPrefix("BOBATOOL_MENUS")
	MenusConfig.AutomaticEnv()

	MenusServiceCmd.PersistentFlags().Bool("insecure", false, "Make insecure client connection. Or use BOBATOOL_MENUS_INSECURE. Must be used with \"address\" option")
	MenusConfig.BindPFlag("insecure", MenusServiceCmd.PersistentFlags().Lookup("insecure"))
	MenusConfig.BindEnv("insecure")

	MenusServiceCmd.PersistentFlags().String("address", "", "Set API address used by client. Or use BOBATOOL_MENUS_ADDRESS.")
	MenusConfig.BindPFlag("address", MenusServiceCmd.PersistentFlags().Lookup("address"))
	MenusConfig.BindEnv("address")

	MenusServiceCmd.PersistentFlags().String("token", "", "Set Bearer token used by the client. Or use BOBATOOL_MENUS_TOKEN.")
	MenusConfig.BindPFlag("token", MenusServiceCmd.PersistentFlags().Lookup("token"))
	MenusConfig.BindEnv("token")

	MenusServiceCmd.PersistentFlags().String("api_key", "", "Set API Key used by the client. Or use BOBATOOL_MENUS_API_KEY.")
	MenusConfig.BindPFlag("api_key", MenusServiceCmd.PersistentFlags().Lookup("api_key"))
	MenusConfig.BindEnv("api_key")
}

var MenusServiceCmd = &cobra.Command{
	Use:       "menus",
	Short:     "Get information about store menus.",
	Long:      "Get information about store menus.",
	ValidArgs: MenusSubCommands,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		var opts []option.ClientOption

		address := MenusConfig.GetString("address")
		if address != "" {
			opts = append(opts, option.WithEndpoint(address))
		}

		if MenusConfig.GetBool("insecure") {
			if address == "" {
				return fmt.Errorf("Missing address to use with insecure connection")
			}

			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				return err
			}
			opts = append(opts, option.WithGRPCConn(conn))
		}

		if token := MenusConfig.GetString("token"); token != "" {
			opts = append(opts, option.WithTokenSource(oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: token,
					TokenType:   "Bearer",
				})))
		}

		if key := MenusConfig.GetString("api_key"); key != "" {
			opts = append(opts, option.WithAPIKey(key))
		}

		MenusClient, err = gapic.NewMenusClient(ctx, opts...)
		return
	},
}
