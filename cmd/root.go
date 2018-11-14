package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	globCfgFile         string
	globCfgOrganization string
	globCfgApiKey       string
	globCfgApiKeySecret string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "id4i",
	Short: "ID4i CLI",
	Long: `ID4i API commandline application.

Configuration:
  You can set global parameters in different locations (ordered by precedence):
  - As command line parameter, e.g. --apikey or -k
  - Using environment variable under the "ID4I_" prefix, e.g. "ID4I_APIKEY"
  - In a configuration file. 
    Default locations are ./.idi.<type> and ~/.id4i.<type> with <type> denoting the file format (json, yml, toml, hcl, properties)
    You can specify a different file using the --config/-c parameter.

  Refer to "id4i help" for available configuration parameters.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&globCfgFile, "config", "c", "", "config file (default is ./.id4i, falls back to $HOME/.id4i)")
	rootCmd.PersistentFlags().StringVarP(&globCfgOrganization, "organization", "o", "", "ID4i organization namespace to work in")
	rootCmd.PersistentFlags().StringVarP(&globCfgApiKey, "apikey", "k", "", "ID4i API key to use")
	rootCmd.PersistentFlags().StringVarP(&globCfgApiKeySecret, "secret", "s", "", "API key secret")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if globCfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(globCfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".id4i" (without extension).
		// Supported extensions are yml, json, properties, hcl, toml
		viper.SetConfigName(".id4i")
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}

	// Map each flag to environment variable ID4I_<FLAG>"
	viper.SetEnvPrefix("id4i")
	viper.AutomaticEnv() // read in environment variables that match

	// bind args passed via the command line to vipers
	viper.BindPFlags(rootCmd.PersistentFlags())
}
