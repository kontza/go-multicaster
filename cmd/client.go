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
	"net"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start a multicast client, i.e. an app that send payloads to a multicast server",
	Run:   multicastClient,
}

func init() {
	rootCmd.AddCommand(clientCmd)
}

func multicastClient(cmd *cobra.Command, args []string) {
	addressToUse, _ := getParams()
	addr, err := net.ResolveUDPAddr("udp", addressToUse)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to ResolveUDPAddr")
	}
	var c *net.UDPConn
	c, err = net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to DialUDP")
	}
	log.Info().Msg("Entering send loop...")
	for {
		payload := time.Now().Format(time.RFC3339)
		log.Info().Str("payload", payload).Msg("Sending")
		c.Write([]byte(payload))
		time.Sleep(3140 * time.Millisecond)
	}
}
