// Copyright © 2018 Wolfgang Werner <mail@wolfgang-werner.net>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/BlueRainSoftware/id4i-cli/api_client"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

var ID4i *api_client.ID4I

var (
	globCfgFile           string
	globParamOrganization string
	globCfgApiKey         string
	globCfgApiKeySecret   string
	globCfgBackend        string
	globParamId4n         string
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
		log.Fatal(err)
	}
}

func Bearer() runtime.ClientAuthInfoWriter {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": viper.Get("apikey"),
	})

	token.Header["iat"] = time.Now()
	token.Header["exp"] = time.Now().Add(time.Second * 30)
	token.Header["typ"] = "API"

	tokenString, err := token.SignedString([]byte(viper.GetString("secret")))
	if err != nil {
		log.Fatal(err)
	}

	return httptransport.BearerToken(tokenString)
}

func OutputResult(result interface{}) {
	j, _ := json.Marshal(result)
	fmt.Println(string(j))
}

func DieOnError(error interface{}) {
	if error != nil {
		log.Error(error)
		j, _ := json.Marshal(error)
		fmt.Println(string(j))
		log.Fatal("Operation failed")
	}
}

func OutputValidationError(error interface{}) {
	log.WithFields(log.Fields{"error": error}).Fatal("Parameter validation failed")
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initClient)

	rootCmd.PersistentFlags().StringVarP(&globCfgFile, "config", "", "", "config file (default is ./.id4i, falls back to $HOME/.id4i)")
	rootCmd.PersistentFlags().StringVarP(&globParamOrganization, "organization", "o", "", "ID4i organization namespace to work in")
	rootCmd.PersistentFlags().StringVarP(&globCfgApiKey, "apikey", "", "", "ID4i API key to use")
	rootCmd.PersistentFlags().StringVarP(&globCfgApiKeySecret, "secret", "", "", "API key secret")
	rootCmd.PersistentFlags().StringVarP(&globCfgBackend, "backend", "", "sandbox.id4i.de", "ID4i Backend to use, e.g. sandbox.id4i.de")
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
			log.Fatal(err)
		}

		// Search config in home directory with name ".id4i" (without extension).
		// Supported extensions are yml, json, properties, hcl, toml
		viper.SetConfigName(".id4i")
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Fatal(err)
	}

	// Map each flag to environment variable ID4I_<FLAG>"
	viper.SetEnvPrefix("id4i")
	viper.AutomaticEnv() // read in environment variables that match

	// bind args passed via the command line to vipers
	viper.BindPFlags(rootCmd.PersistentFlags())
}

func initClient() {
	ID4i = api_client.NewHTTPClientWithConfig(
		strfmt.Default,
		api_client.DefaultTransportConfig().
			WithSchemes([]string{"https"}).
			WithHost(viper.GetString("backend")))
}
