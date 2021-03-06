package cmd

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/cmd/data"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/configs"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/database"

	"github.com/spf13/cobra"
)

var ImportCMD = &cobra.Command{
	Use: "import",
	PreRun: func(cmd *cobra.Command, args []string) {
		database.MySQLInit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Step One
		fmt.Println(args)
		if len(args) == 0 {
			return
		}

		if args[0] == "symbol" {
			fmt.Println(data.RunForSymbol(configs.MatchSymbol))
		}
		if args[0] == "championship" || args[0] == "ch" {
			fmt.Println(data.RunChampionship(configs.MatchDescription))
		}
		if args[0] == "history" || args[0] == "his" {
			fmt.Println(data.RunFiNaHistory(configs.MatchHistory))
		}
		if args[0] == "brief" || args[0] == "brf" {
			fmt.Println(data.RunFiNaBrief(configs.MatchBrief))
		}
		if args[0] == "sports" || args[0] == "sp" {
			fmt.Println(data.RunSports(configs.MatchSportsMap))
		}
		if args[0] == "events" {
			fmt.Println(data.RunPostEvent(configs.MatchHistoryYear))
		}
		if args[0] == "records" {
			fmt.Println(data.RunRecords(configs.MatchRecords))
		}
		if args[0] == "rank" {
			fmt.Println(data.RunRank(configs.MatchRank))
		}

	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer database.MySQL.Close()
	},
}
