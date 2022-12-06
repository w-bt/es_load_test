package cmd

import (
	"es_load_test/appcontext"
	"es_load_test/config"
	"es_load_test/internal/handlers"
	"es_load_test/internal/repositories/item"
	svItem "es_load_test/internal/services/item"
	"github.com/spf13/cobra"
	"log"
)

func StartAPIServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "server",
		Short:   "Start HTTP API server",
		Aliases: []string{"serve", "start"},
		Run: func(_ *cobra.Command, _ []string) {
			err := config.Load()
			if err != nil {
				log.Fatalf("Config load failed: \n%s", err.Error())
			}

			serverDeps, err := initDependencies()
			if err != nil {
				log.Fatalf("Server Dependencies initialization failed: \n%s", err.Error())
			}

			srv := newServer(handlers.NewRouter(serverDeps))
			log.Printf("server listening on %s.", config.Cfg.Addr())
			srv.Serve(config.Cfg.Addr())
		},
	}
}

func initDependencies() (*appcontext.ServerDependencies, error) {
	builderService := item.NewBuilderService()
	queriesService, err := item.NewQueriesService(builderService)
	if err != nil {
		return nil, err
	}
	itemRepo := item.NewItemRepo(queriesService)
	itemService := svItem.NewItemService(itemRepo)

	serverDependencies := &appcontext.ServerDependencies{
		Services: appcontext.Services{Item: itemService},
	}

	return serverDependencies, nil
}
