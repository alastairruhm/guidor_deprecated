package main

import (
	"fmt"
	"os"

	"github.com/alastairruhm/guidor/cmd/client/backup"
	"github.com/alastairruhm/guidor/cmd/client/register"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("guidor")

var formatStdout = logging.MustStringFormatter(
	`%{color}%{time:2006-01-02 15:04:05.000} %{level:.8s} %{message}%{color:reset}`,
)

var formatLogfile = logging.MustStringFormatter(
	`%{time:2006-01-02 15:04:05.000} %{level:.8s} %{message}`,
)

var rootFlags struct {
	configFile string
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&rootFlags.configFile, "config", "c", "guidor.toml", "configuration file, default is guidor.toml")
}

var rootCmd = &cobra.Command{
	Use:   "guidor",
	Short: "guidor service cli tool",
	Long:  `database backup management command line tool`,
}

func main() {
	cobra.OnInitialize(initConfig, initLog)
	addCommands()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func addCommands() {
	rootCmd.AddCommand(register.Cmd)
	rootCmd.AddCommand(backup.Cmd)
}

func initConfig() {
	viper.SetConfigFile(rootFlags.configFile)
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("config file parse error")
		panic(err)
	}
	fmt.Println("using config file", viper.ConfigFileUsed())
}

// log output to stdout and file both
func initLog() {
	logFile, err := os.OpenFile(viper.GetString("log.file"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	backendLogFile := logging.NewLogBackend(logFile, "", 0)
	backendStdout := logging.NewLogBackend(os.Stdout, "", 0)

	backendLogFileFormatter := logging.NewBackendFormatter(backendLogFile, formatLogfile)
	backendStdoutFormatter := logging.NewBackendFormatter(backendStdout, formatStdout)

	//backendLogFileLeveled := logging.AddModuleLevel(backendLogFile)
	//backendLogFileLeveled.SetLevel(logging.INFO, "")

	logging.SetBackend(backendLogFileFormatter, backendStdoutFormatter)
}
