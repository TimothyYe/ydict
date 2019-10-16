package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/TimothyYe/ydict/lib"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	withVoice int
	withMore  bool
	withCache bool
	isQuiet   bool
	isDelete  bool
	withReset bool
)

func main() {
	//Check & load .env file
	lib.LoadEnv()

	if len(os.Args) == 1 ||
		(len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "-help")) {
		lib.DisplayLogo()
	}

	var rootCmd = &cobra.Command{
		Use: "ydict",
		Run: func(cmd *cobra.Command, args []string) {
			if args[0] == "help" {
				if err := cmd.Usage(); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				return
			}

			if withReset {
				lib.ClearCahceFiles()
				return
			}

			if isDelete {
				if err := lib.DeleteWords(args); err == nil {
					color.Green("  Word '%s' has already been removed from the cache.", strings.Join(args, " "))
				}
				return
			}

			queryP := lib.QueryParam{}
			queryP.Words = args
			queryP.WordString = strings.Join(args, " ")
			queryP.WithMore = withMore
			queryP.WithCache = withCache
			queryP.IsQuiet = isQuiet
			queryP.IsMulti = (len(args) > 1)
			queryP.IsChinese = lib.IsChinese(queryP.WordString)
			queryP.WithVoice = withVoice
			if !lib.IsAvailableOS() {
				queryP.WithVoice = 0
			}
			queryP.DoQuery()
		},
	}

	rootCmd.PersistentFlags().IntVarP(&withVoice, "voice", "v", 0, "Query with voice speech, the default voice play count is 0.")
	rootCmd.PersistentFlags().BoolVarP(&withMore, "more", "m", false, "Query with more example sentences.")
	rootCmd.PersistentFlags().BoolVarP(&withCache, "cache", "c", false, "Query with local cache, and save the query word(s) into the cache.")
	rootCmd.PersistentFlags().BoolVarP(&withReset, "reset", "r", false, "Clear all the words from the local cache.")
	rootCmd.PersistentFlags().BoolVarP(&isQuiet, "quiet", "q", false, "Query with quiet mode, don't show spinner.")
	rootCmd.PersistentFlags().BoolVarP(&isDelete, "delete", "d", false, "Remove word(s) from the cache.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
