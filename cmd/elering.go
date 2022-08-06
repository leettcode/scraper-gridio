package cmd

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/cobra"
)

var eleringCmd = &cobra.Command{
	Use:   "elering",
	Short: "starts elering scraper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("elering scraper started")

		//client := sqlite.OpenEnt(cfg.DatabaseURL)
		//defer func(client *ent.Client) {
		//	err := client.Close()
		//	if err != nil {
		//		log.Fatal("error closing client")
		//	}
		//}(client)
		//flag.Parse()

	},
}

func init() {
	rootCmd.AddCommand(eleringCmd)
}
