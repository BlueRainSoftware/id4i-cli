package cmd

import (
	"fmt"
	"github.com/BlueRainSoftware/id4i-cli/api_client"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/viper"
	"os"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/spf13/cobra"
)

var (
	cfgShowConfig      bool
	cfgShowBackendInfo bool
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Configuration and backend API information",
	Run: func(cmd *cobra.Command, args []string) {

		if cfgShowConfig {
			fmt.Printf("Config file: %s\n", viper.ConfigFileUsed())

			for _, element := range viper.AllKeys() {
				fmt.Printf("%s\t%s\n", element, viper.Get(element))
			}
		}

		if cfgShowBackendInfo {
			fmt.Printf("now call metadata api")


			client := api_client.NewHTTPClientWithConfig(
				strfmt.Default,
				api_client.DefaultTransportConfig().
					WithSchemes([]string {"https"}).
					WithHost("backend.id4i.de"))

			resp, accepted, err := client.MetaInformation.ApplicationInfo(nil, bearer())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(accepted)
			fmt.Printf("%#v\n", resp.Payload)
		}
	},
}

func bearer() runtime.ClientAuthInfoWriter {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,jwt.MapClaims{
		"sub": viper.Get("apikey"),
	})

	token.Header["iat"] = time.Now()
	token.Header["exp"] = time.Now().Add(time.Second * 30)
	token.Header["typ"] = "API"

	fmt.Println(token)

	tokenString, err := token.SignedString([]byte(viper.GetString("secret")))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(tokenString)

	return httptransport.BearerToken(tokenString)
}

func init() {
	rootCmd.AddCommand(infoCmd)

	infoCmd.Flags().BoolVarP(&cfgShowConfig, "show-config", "", true, "Show configuration")
	infoCmd.Flags().BoolVarP(&cfgShowBackendInfo, "show-backend", "", true, "Show backend information")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
