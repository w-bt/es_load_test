package cmd

import (
	"context"
	"es_load_test/config"
	"es_load_test/constant"
	"es_load_test/internal/repositories/cluster"
	"github.com/spf13/cobra"
	"log"
)

func CreateIndexCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "create_index",
		Short:   "Create IndexName",
		Aliases: []string{"create"},
		Run: func(_ *cobra.Command, args []string) {
			indexName := constant.DefaultIndexName
			if len(args) == 1 {
				indexName = args[0]
			}

			err := config.Load()
			if err != nil {
				log.Fatalf("Dependencies initialization failed: \n%s", err.Error())
			}

			serverDeps, err := initDependencies()
			if err != nil {
				log.Fatalf("Dependencies initialization failed: \n%s", err.Error())
			}

			req := cluster.Request{IndexName: indexName}

			result, err := serverDeps.Services.Item.CreateIndex(context.Background(), req)
			if err != nil {
				log.Fatalf("Failed to create new index: \n%s", err.Error())
			}

			log.Printf("%+v", *result)
		},
	}
}
