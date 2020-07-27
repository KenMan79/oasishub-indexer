package indexer

import (
	"context"
	"fmt"
	"time"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/figment-networks/oasishub-indexer/metric"
	"github.com/figment-networks/oasishub-indexer/model"
	"github.com/figment-networks/oasishub-indexer/utils/logger"
)

const (
	SyncerPersistorTaskName       = "SyncerPersistor"
	BlockSeqPersistorTaskName     = "BlockSeqPersistor"
	ValidatorSeqPersistorTaskName = "ValidatorSeqPersistor"
	ValidatorAggPersistorTaskName = "ValidatorAggPersistor"
)

func NewSyncerPersistorTask(db SyncerPersistorTaskStore) pipeline.Task {
	return &syncerPersistorTask{
		db: db,
	}
}

type SyncerPersistorTaskStore interface {
	CreateOrUpdate(val *model.Syncable) error
}

type syncerPersistorTask struct {
	db SyncerPersistorTaskStore
}

func (t *syncerPersistorTask) GetName() string {
	return SyncerPersistorTaskName
}

func (t *syncerPersistorTask) Run(ctx context.Context, p pipeline.Payload) error {
	defer metric.LogIndexerTaskDuration(time.Now(), t.GetName())

	payload := p.(*payload)

	logger.Info(fmt.Sprintf("running indexer task [stage=%s] [task=%s] [height=%d]", pipeline.StagePersistor, t.GetName(), payload.CurrentHeight))

	return t.db.CreateOrUpdate(payload.Syncable)
}

func NewBlockSeqPersistorTask(db BlockSeqPersistorTaskStore) pipeline.Task {
	return &blockSeqPersistorTask{
		db: db,
	}
}

type blockSeqPersistorTask struct {
	db BlockSeqPersistorTaskStore
}

type BlockSeqPersistorTaskStore interface {
	Create(record interface{}) error
	Save(record interface{}) error
}

func (t *blockSeqPersistorTask) GetName() string {
	return BlockSeqPersistorTaskName
}

func (t *blockSeqPersistorTask) Run(ctx context.Context, p pipeline.Payload) error {
	defer metric.LogIndexerTaskDuration(time.Now(), t.GetName())

	payload := p.(*payload)

	logger.Info(fmt.Sprintf("running indexer task [stage=%s] [task=%s] [height=%d]", pipeline.StagePersistor, t.GetName(), payload.CurrentHeight))

	if payload.NewBlockSequence != nil {
		return t.db.Create(payload.NewBlockSequence)
	}

	if payload.UpdatedBlockSequence != nil {
		return t.db.Save(payload.UpdatedBlockSequence)
	}

	return nil
}

func NewValidatorSeqPersistorTask(db ValidatorSeqPersistorTaskStore) pipeline.Task {
	return &validatorSeqPersistorTask{
		db: db,
	}
}

type ValidatorSeqPersistorTaskStore interface {
	Create(record interface{}) error
	Save(record interface{}) error
}

type validatorSeqPersistorTask struct {
	db ValidatorSeqPersistorTaskStore
}

func (t *validatorSeqPersistorTask) GetName() string {
	return ValidatorSeqPersistorTaskName
}

func (t *validatorSeqPersistorTask) Run(ctx context.Context, p pipeline.Payload) error {
	defer metric.LogIndexerTaskDuration(time.Now(), t.GetName())

	payload := p.(*payload)

	logger.Info(fmt.Sprintf("running indexer task [stage=%s] [task=%s] [height=%d]", pipeline.StagePersistor, t.GetName(), payload.CurrentHeight))

	for _, sequence := range payload.NewValidatorSequences {
		if err := t.db.Create(&sequence); err != nil {
			return err
		}
	}

	for _, sequence := range payload.UpdatedValidatorSequences {
		if err := t.db.Save(&sequence); err != nil {
			return err
		}
	}

	return nil
}

func NewValidatorAggPersistorTask(db ValidatorAggPersistorTaskStore) pipeline.Task {
	return &validatorAggPersistorTask{
		db: db,
	}
}

type ValidatorAggPersistorTaskStore interface {
	Create(record interface{}) error
	Save(record interface{}) error
}

type validatorAggPersistorTask struct {
	db ValidatorAggPersistorTaskStore
}

func (t *validatorAggPersistorTask) GetName() string {
	return ValidatorAggPersistorTaskName
}

func (t *validatorAggPersistorTask) Run(ctx context.Context, p pipeline.Payload) error {
	defer metric.LogIndexerTaskDuration(time.Now(), t.GetName())

	payload := p.(*payload)

	logger.Info(fmt.Sprintf("running indexer task [stage=%s] [task=%s] [height=%d]", pipeline.StagePersistor, t.GetName(), payload.CurrentHeight))

	for _, aggregate := range payload.NewAggregatedValidators {
		if err := t.db.Create(&aggregate); err != nil {
			return err
		}
	}

	for _, aggregate := range payload.UpdatedAggregatedValidators {
		if err := t.db.Save(&aggregate); err != nil {
			return err
		}
	}

	return nil
}
