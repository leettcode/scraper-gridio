package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/cobra"

	"github.com/alexgtn/go-linkshort/internal/infra/repository/nordpool"
	"github.com/alexgtn/go-linkshort/internal/infra/sqlite"
	"github.com/alexgtn/go-linkshort/internal/usecase"
	ent "github.com/alexgtn/go-linkshort/tools/ent/codegen"
)

var nordpoolCmd = &cobra.Command{
	Use:   "nordpool",
	Short: "starts nordpool scraper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nordpool started")

		client := sqlite.OpenEnt(cfg.DatabaseURL)

		defer func(client *ent.Client) {
			err := client.Close()
			if err != nil {
				log.Fatal("error closing client")
			}
		}(client)
		ctx := context.Background()

		nordpoolRepo := nordpool.NewNordpoolSqliteRepo(client)

		isDone := make(chan bool)
		scrapingInterval := time.NewTicker(time.Hour * 1)
		httpClient := resty.New() // can configure backoff etc.
		nordpoolScraperSvc := usecase.NewNordpoolScraperService(ctx,
			nordpoolRepo, scrapingInterval, isDone, cfg.NordpoolURL, httpClient,
		)

		// check sigterm
		// trigger isDone

	},
}

func init() {
	rootCmd.AddCommand(nordpoolCmd)
}
