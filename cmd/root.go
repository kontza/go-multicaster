/*
Copyright © 2022 Juha Ruotsalainen <kontza@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "multicaster",
	Short: "Start a multicaster client or server",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(loggerInit)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.multicaster.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	var address string
	rootCmd.PersistentFlags().StringVarP(&address, "address", "a", "224.0.0.1:9999", "address:port to use")
	var maxDatagramSize int
	rootCmd.PersistentFlags().IntVarP(&maxDatagramSize, "datagram-size", "d", 8192, "maximum datagram size")
}

func getParams() (string, int) {
	var addressToUse string
	var maxDatagramSize int
	if address, err := rootCmd.Flags().GetString("address"); err != nil {
		log.Fatal().Err(err)
	} else {
		addressToUse = address
	}
	if datagramSize, err := rootCmd.Flags().GetInt("datagram-size"); err != nil {
		log.Fatal().Err(err)
	} else {
		maxDatagramSize = datagramSize
	}
	log.Info().
		Str("address", addressToUse).
		Int("datagram size", maxDatagramSize).
		Msg("Using")
	return addressToUse, maxDatagramSize
}
