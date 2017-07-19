package backup

import (
	"fmt"

	"strings"

	sh "github.com/codeskyblue/go-sh"
	logging "github.com/op/go-logging"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var log = logging.MustGetLogger("guidor")

func init() {
	Cmd.Flags().StringP("user", "u", "", "user for backup")
	Cmd.Flags().StringP("password", "p", "", "password for backup")
	Cmd.Flags().StringP("mode", "", "temp", "backup mode: temp or schedule")

	viper.BindPFlag("database.user", Cmd.Flags().Lookup("user"))
	viper.BindPFlag("database.password", Cmd.Flags().Lookup("password"))
}

// Cmd Backup
var Cmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup action: guidor backup",
	Run:   backup,
}

func backup(cmd *cobra.Command, args []string) {
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	argsStr := fmt.Sprintf("--user=%s --password=%s --default-character-set=utf8 --all-databases --result-file=backup.sql", user, password)
	argsBackup := strings.Split(argsStr, " ")
	err := sh.Command("mysqldump", argsBackup).Run()
	if err != nil {
		log.Error("mysqldump process error")
	}
	log.Info("mysqldump process success")
}
