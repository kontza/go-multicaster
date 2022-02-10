package cmd

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func loggerInit() {
	verbose, _ := rootCmd.Flags().GetBool("verbose")
	structured, _ := rootCmd.Flags().GetBool("structured")

	// Set time field format to ISO-8601 with T being replaced with a space,
	// and milliseconds containing trailing zeros.
	var sb strings.Builder
	sb.WriteString(strings.ReplaceAll(time.RFC3339[:19], "T", " "))
	sb.WriteString(".000")
	zerolog.TimeFieldFormat = sb.String()

	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if structured {
		log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: zerolog.TimeFieldFormat}
		log.Logger = zerolog.New(output).With().Timestamp().Logger()
	}

	log.Debug().Str("logging level", zerolog.GlobalLevel().String()).Msg("Current")
}
