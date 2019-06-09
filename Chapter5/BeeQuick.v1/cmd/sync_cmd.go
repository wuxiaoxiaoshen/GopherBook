package cmd

import (
	"fmt"

	"GopherBook/Chapter5/BeeQuick.v1/model/v1"
	"GopherBook/Chapter5/BeeQuick.v1/pkg/database.v1"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

var syncCMD = &cobra.Command{
	Use:   "sync2",
	Short: "xorm.Syn2(model)",
	Run:   sync2,
}

func sync2(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Panic("You should add one argument at least")
		return
	}
	database_v1.DataBaseInit()
	if args[0] == "db" {

		for _, i := range tables() {
			if err := database_v1.BeeQuickDatabase.Sync2(i); err != nil {
				fmt.Println(err)
			}
		}
	}
	if args[0] == "vip" {
		vipMember()
	}
}

func tables() []interface{} {
	return []interface{}{
		new(model_v1.Account),
		new(model_v1.VipMember),
	}
}

func vipMember() bool {
	if _, err := database_v1.BeeQuickDatabase.Insert(model_v1.DefaultVipMemberRecord()); err != nil {
		return false
	}
	return true
}