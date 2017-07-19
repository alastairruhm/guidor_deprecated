package register

import (
	"github.com/op/go-logging"
	"github.com/spf13/cobra"
)

var (
	ip          string
	hostname    string
	dbtype      string
	dbversion   string
	serviceName string
)

var log = logging.MustGetLogger("guidor")

func init() {
	Cmd.Flags().StringVarP(&ip, "ip", "", "", "guidor client ip")
}

// Cmd Register
var Cmd = &cobra.Command{
	Use:     "register",
	Short:   "Register commands: guidor client register --help",
	Long:    "Register commands: guidor register [command]",
	Aliases: []string{"r"},
	Run:     registerClient,
}

func registerClient(cmd *cobra.Command, args []string) {
	client := client.DBMClient{Host: hostname}
	token, err := client.Register(hostname, "alastair", "mysql")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Info("%+v\n", token)
}
