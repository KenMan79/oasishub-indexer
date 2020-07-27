package cli

import (
	"context"
	"fmt"
	"github.com/figment-networks/oasishub-indexer/config"
	"github.com/figment-networks/oasishub-indexer/usecase"
	"github.com/figment-networks/oasishub-indexer/utils/logger"
	"github.com/pkg/errors"
)

func runCmd(cfg *config.Config, flags Flags) error {
	db, err := initStore(cfg)
	if err != nil {
		return err
	}
	defer db.Close()
	client, err := initClient(cfg)
	if err != nil {
		return err
	}
	defer client.Close()

	cmdHandlers := usecase.NewCmdHandlers(cfg, db, client)

	logger.Info(fmt.Sprintf("executing cmd %s ...", flags.runCommand), logger.Field("app", "cli"))

	ctx := context.Background()
	switch flags.runCommand {
	case "status":
		cmdHandlers.GetStatus.Handle(ctx)
	case "indexer_index":
		cmdHandlers.IndexerIndex.Handle(ctx, flags.batchSize)
	case "indexer_backfill":
		cmdHandlers.IndexerBackfill.Handle(ctx, flags.parallel, flags.force, flags.versionIds, flags.targetIds)
	case "indexer_summarize":
		cmdHandlers.IndexerSummarize.Handle(ctx)
	case "indexer_purge":
		cmdHandlers.IndexerPurge.Handle(ctx)
	default:
		return errors.New(fmt.Sprintf("command %s not found", flags.runCommand))
	}
	return nil
}

