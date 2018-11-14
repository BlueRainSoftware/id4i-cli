package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
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
				switch element {
				case "secret":
					fmt.Printf("%s\t%s\n", element, "*****")
				default:
					fmt.Printf("%s\t%s\n", element, viper.Get(element))
				}
			}
		}

		if cfgShowBackendInfo {
			fmt.Printf("now call metadata api")

			resp, accepted, err := Client().MetaInformation.ApplicationInfo(nil, Bearer())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(accepted)
			fmt.Printf("%#v\n", resp.Payload)
		}
	},
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
