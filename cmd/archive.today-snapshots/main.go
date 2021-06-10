package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"jaytaylor.com/archive.today"
)

var (
	Quiet          bool
	Verbose        bool
	Wait           bool
	RequestTimeout time.Duration = archivetoday.DefaultRequestTimeout
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "Activate quiet log output")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Activate verbose log output")
	rootCmd.PersistentFlags().DurationVarP(&RequestTimeout, "request-timeout", "r", RequestTimeout, "Timeout duration for HTTP requests")
	rootCmd.PersistentFlags().StringVarP(&archivetoday.BaseURL, "base-url", "b", archivetoday.BaseURL, "Archive.today server base hostname / URL address")
	rootCmd.PersistentFlags().StringVarP(&archivetoday.UserAgent, "user-agent", "u", archivetoday.UserAgent, "'User-Agent' header to use")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		errorExit(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "archive.today-snapshots",
	Short: "search for archive.today snapshots",
	Long:  "command-line interface for searching archive.today for URL page snapshots",
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(_ *cobra.Command, _ []string) {
		initLogging()
	},
	Run: func(cmd *cobra.Command, args []string) {
		snapshots, err := archivetoday.Search(args[0], RequestTimeout)
		if err != nil {
			errorExit(err)
		}

		log.Infof("Found %v results", len(snapshots))

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "    ")

		if err := enc.Encode(&snapshots); err != nil {
			errorExit(fmt.Errorf("marshalling snapshots to JSON: %s", err))
		}
	},
}

func errorExit(err interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	os.Exit(1)
}

func initLogging() {
	level := log.InfoLevel
	if Verbose {
		level = log.DebugLevel
	}
	if Quiet {
		level = log.ErrorLevel
	}
	log.SetLevel(level)
}
