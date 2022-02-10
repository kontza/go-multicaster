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

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a multicast server, i.e. an app that listens to multicast payloads",
	Run:   multicastServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Info().
		Int("bytes", n).
		Str("data", string(b[:n])).
		Interface("from", src).
		Msg("Received")
}

func multicastServer(cmd *cobra.Command, args []string) {
	addressToUse, maxDatagramSize := getParams()
	addr, err := net.ResolveUDPAddr("udp", addressToUse)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to ResolveUDPAddr")
	}
	l, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to ListenMulticastUDP")
	}
	l.SetReadBuffer(maxDatagramSize)
	log.Info().Msg("Entering receive loop...")
	for {
		b := make([]byte, maxDatagramSize)
		n, src, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal().Err(err).Msg("ReadFromUDP failed!")
		}
		msgHandler(src, n, b)
	}
}
