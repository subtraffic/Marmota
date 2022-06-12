package application

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "marmota",
	Short: "start a dns auto-update server",
}

func init() {
	echoIpCmd := &cobra.Command{
		Use:   "echo",
		Short: "show your ip according to the server you provide",
		Run: func(cmd *cobra.Command, args []string) {
			// echo your ip
			fmt.Println("your ip is:")
		},
	}

	webCmd := &cobra.Command{
		Use:   "server",
		Short: "run marmoto server",
		Run: func(cmd *cobra.Command, args []string) {
			// about you webserver
		},
	}

	webCmd.Flags().BoolP("ui", "u", true, "running with ui web")
	webCmd.Flags().BoolP("server", "s", true, "provide echo ip server for public website")
	webCmd.Flags().StringP("configfile", "c", "~/.marmoto/marmoto.yml", "the path of configfile")
	webCmd.Flags().StringP("datafile", "d", "~/.marmoto/data.yml", "the data of user seting and running status")
	webCmd.Flags().String("host", "0.0.0.0", "the host of server")
	webCmd.Flags().Int("port", 9787, "the defautl server port")

	rootCmd.AddCommand(echoIpCmd)
	rootCmd.AddCommand(webCmd)
}

func Start() {
	rootCmd.Execute()
}
