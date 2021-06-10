package main

import (
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
	PollInterval   time.Duration = archivetoday.DefaultPollInterval
	WaitTimeout    time.Duration = time.Duration(0)
	Anyway         bool
	submitID       string
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "Activate quiet log output")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Activate verbose log output")
	rootCmd.PersistentFlags().BoolVarP(&Wait, "wait", "w", false, "Wait for crawl to finish before returning URL result")
	rootCmd.PersistentFlags().DurationVarP(&RequestTimeout, "request-timeout", "r", RequestTimeout, "Timeout duration for HTTP requests")
	rootCmd.PersistentFlags().DurationVarP(&PollInterval, "poll-interval", "p", PollInterval, "Poll interval, only applies when -w/--wait is active")
	rootCmd.PersistentFlags().DurationVarP(&WaitTimeout, "wait-timeout", "", WaitTimeout, "Maximum wait duration, only applies when -w/--wait is active (default: infinite)")
	rootCmd.PersistentFlags().BoolVarP(&Anyway, "anyway", "a", false, "Force archival even if there is already a recent snapshot of the page")
	rootCmd.PersistentFlags().StringVarP(&archivetoday.BaseURL, "base-url", "b", archivetoday.BaseURL, "Archive.today server base hostname / URL address")
	rootCmd.PersistentFlags().StringVarP(&archivetoday.HTTPHost, "http-host", "", archivetoday.HTTPHost, "'Host' header to use")
	rootCmd.PersistentFlags().StringVarP(&archivetoday.UserAgent, "user-agent", "u", archivetoday.UserAgent, "'User-Agent' header to use")
	rootCmd.PersistentFlags().StringVarP(&submitID, "submitid", "s", "", "Manually specify submitid value to use")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		errorExit(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "archive.today",
	Short: "submits URLs to archive.today for archival",
	Long:  "command-line interface to submit URLs to archive.today for webpage snapshot archival",
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(_ *cobra.Command, _ []string) {
		initLogging()
	},
	Run: func(cmd *cobra.Command, args []string) {
		cfg := archivetoday.Config{
			Anyway:   Anyway,
			Wait:     Wait,
			SubmitID: submitID,
		}

		result, err := archivetoday.Capture(args[0], cfg)

		if len(result) > 0 {
			fmt.Println(result)
		}

		if err != nil {
			errorExit(err)
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
