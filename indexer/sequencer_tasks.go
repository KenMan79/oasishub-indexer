package indexer

import (
	"context"
	"fmt"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/figment-networks/oasishub-indexer/model"
	"github.com/figment-networks/oasishub-indexer/store"
	"github.com/figment-networks/oasishub-indexer/utils/logger"
)

const (
	TaskNameBlockSeqCreator               = "BlockSeqCreator"
	TaskNameValidatorSeqCreator           = "ValidatorSeqCreator"
	TaskNameTransactionSeqCreator         = "TransactionSeqCreator"
	TaskNameStakingSeqCreator             = "StakingSeqCreator"
	TaskNameDelegationSeqCreator          = "DelegationSeqCreator"
	TaskNameDebondingDelegationSeqCreator = "DebondingDelegationSeqCreator"
)

var (
	_ pipeline.Task = (*blockSeqCreatorTask)(nil)
	_ pipeline.Task = (*validatorSeqCreatorTask)(nil)
	_ pipeline.Task = (*transactionSeqCreatorTask)(nil)
	_ pipeline.Task = (*stakingSeqCreatorTask)(nil)
	_ pipeline.Task = (*delegationSeqCreatorTask)(nil)
	_ pipeline.Task = (*debondingDelegationSeqCreatorTask)(nil)
)

func NewBlockSeqCreatorTask(db BlockSeqCreatorTaskStore) *blockSeqCreatorTask {
	return &blockSeqCreatorTask{
		db: db,
	}
}

type blockSeqCreatorTask struct {
	db BlockSeqCreatorTaskStore
}

type BlockSeqCreatorTaskStore interface {
	FindByHeight(height int64) (*model.BlockSeq, error)
}

func (t *blockSeqCreatorTask) GetName() string {
	return TaskNameBlockSeqCreator
}

func (t *blockSeqCreatorTask) Run(ctx context.Context, p pipeline.Payload) error {

	payload := p.(*payload)

	logger.Info(fmt.Sprintf("running indexer task [stage=%s] [task=%s] [height=%d]", pipeline.StageSequencer, t.GetName(), payload.CurrentHeight))

	rawBlockSeq, err := BlockToSequence(payload.Syncable, payload.ParsedBlock)
	if err != nil {
		return err
	}

	blockSeq, err := t.db.FindByHeight(payload.CurrentHeight)
	if err != nil {
		if err == store.ErrNotFound {
			payload.NewBlockSequence = rawBlockSeq
			return nil
		} else {
			return err
		}
	}

	blockSeq.Update(*rawBlockSeq)
	payload.UpdatedBlockSequence = blockSeq

	return nil
}

func NewValidatorSeqCreatorTask(db ValidatorSeqCreatorTaskStore) *validatorSeqCreatorTask {
	return &validatorSeqCreatorTask{
		db: db,
	}
}

type validatorSeqCreatorTask struct {
	db ValidatorSeqCreatorTaskStore
}

type ValidatorSeqCreatorTaskStore interface {
	FindByHeightAndEntityUID(h int64, key string) (*model.ValidatorSeq, error)
}

func (t *validatorSeqCreatorTask) GetName() string {
	return TaskNameValidatorSeqCreator
}

func (t *validatorSeqCreatorTask) Run(ctx context.Context, p pipeline.Payload) error {

	payload := p.(*payload)

	logger.Info(fmt.Sprintf("running indexer task [stage=%s] [task=%s] [height=%d]", pipeline.StageSequencer, t.GetName(), payload.CurrentHeight))

	rawValidatorSeqs, err := ValidatorToSequence(payload.Syncable, payload.RawValidators, payload.ParsedValidators)
	if err != nil {
		return err
	}

	var newValidatorSeqs []model.ValidatorSeq
	var updatedValidatorSeqs []model.ValidatorSeq
	for _, rawValidatorSeq := range rawValidatorSeqs {
		validatorSeq, err := t.db.FindByHeightAndEntityUID(payload.CurrentHeight, rawValidatorSeq.EntityUID)
		if err != nil {
			if err == store.ErrNotFound {
				newValidatorSeqs = append(newValidatorSeqs, rawValidatorSeq)
				continue
			} else {
				return err
			}
		}

		validatorSeq.Update(rawValidatorSeq)
		updatedValidatorSeqs = append(updatedValidatorSeqs, *validatorSeq)
	}

	payload.NewValidatorSequences = newValidatorSeqs
	payload.UpdatedValidatorSequences = updatedValidatorSeqs

	return nil
}

func NewTransactionSeqCreatorTask(db TransactionSeqCreatorTaskStore) *transactionSeqCreatorTask {
	return &transactionSeqCreatorTask{
		db: db,
	}
}

type transactionSeqCreatorTask struct {
	db TransactionSeqCreatorTaskStore
}

type TransactionSeqCreatorTaskStore interface {
	Create(record interface{}) error
	FindByHeight(h int64) ([]model.TransactionSeq, error)
}

func (t *transactionSeqCreatorTask) GetName() string {
	return TaskNameTransactionSeqCreator
}

func (t *transactionSeqCreatorTask) Run(ctx context.Context, p pipeline.Payload) error {

	payload := p.(*payload)

	logger.Info(fmt.Sprintf("running indexer task [stage=%s] [task=%s] [height=%d]", pipeline.StageSequencer, t.GetName(), payload.CurrentHeight))

	var res []model.TransactionSeq
	sequenced, err := t.db.FindByHeight(payload.CurrentHeight)
	if err != nil {
		return err
	}

	toSequence, err := TransactionToSequence(payload.Syncable, payload.RawTransactions)
	if err != nil {
		return err
	}

	// Nothing to sequence
	if len(toSequence) == 0 {
		payload.TransactionSequences = res
		return nil
	}

	// Everything sequenced and saved to persistence
	if len(sequenced) == len(toSequence) {
		payload.TransactionSequences = sequenced
		return nil
	}

	isSequenced := func(vs model.TransactionSeq) bool {
		for _, sv := range sequenced {
			if sv.Equal(vs) {
				return true
			}
		}
		return false
	}

	for _, vs := range toSequence {
		if !isSequenced(vs) {
			if err := t.db.Create(&vs); err != nil {
				return err
			}
		}
		res = append(res, vs)
	}
	payload.TransactionSequences = res
	return nil
}

func NewStakingSeqCreatorTask(db StakingSeqCreatorTaskStore) *stakingSeqCreatorTask {
	return &stakingSeqCreatorTask{
		db: db,
	}
}

type stakingSeqCreatorTask struct {
	db StakingSeqCreatorTaskStore
}

type StakingSeqCreatorTaskStore interface {
	Create(record interface{}) error
	FindByHeight(height int64) (*model.StakingSeq, error)
}

func (t *stakingSeqCreatorTask) GetName() string {
	return TaskNameStakingSeqCreator
}

func (t *stakingSeqCreatorTask) Run(ctx context.Context, p pipeline.Payload) error {

	payload := p.(*payload)

	logger.Info(fmt.Sprintf("running indexer task [stage=%s] [task=%s] [height=%d]", pipeline.StageSequencer, t.GetName(), payload.CurrentHeight))

	sequenced, err := t.db.FindByHeight(payload.CurrentHeight)
	if err != nil {
		if err == store.ErrNotFound {
			toSequence, err := StakingToSequence(payload.Syncable, payload.RawState.GetStaking())
			if err != nil {
				return err
			}
			if err := t.db.Create(toSequence); err != nil {
				return err
			}
			payload.StakingSequence = toSequence
			return nil
		}
		return err
	}
	payload.StakingSequence = sequenced
	return nil
}

type delegationSeqCreatorTask struct {
	db DelegationSeqCreatorTaskStore
}

type DelegationSeqCreatorTaskStore interface {
	Create(record interface{}) error
	FindByHeight(h int64) ([]model.DelegationSeq, error)
}

func (t *delegationSeqCreatorTask) GetName() string {
	return TaskNameDelegationSeqCreator
}

func NewDelegationsSeqCreatorTask(db DelegationSeqCreatorTaskStore) *delegationSeqCreatorTask {
	return &delegationSeqCreatorTask{
		db: db,
	}
}

func (t *delegationSeqCreatorTask) Run(ctx context.Context, p pipeline.Payload) error {

	payload := p.(*payload)

	logger.Info(fmt.Sprintf("running indexer task [stage=%s] [task=%s] [height=%d]", pipeline.StageSequencer, t.GetName(), payload.CurrentHeight))

	var res []model.DelegationSeq
	sequenced, err := t.db.FindByHeight(payload.CurrentHeight)
	if err != nil {
		return err
	}

	toSequence, err := DelegationToSequence(payload.Syncable, payload.RawState)
	if err != nil {
		return err
	}

	// Nothing to sequence
	if len(toSequence) == 0 {
		payload.DelegationSequences = res
		return nil
	}

	// Everything sequenced and saved to persistence
	if len(sequenced) == len(toSequence) {
		payload.DelegationSequences = sequenced
		return nil
	}

	isSequenced := func(vs model.DelegationSeq) bool {
		for _, sv := range sequenced {
			if sv.Equal(vs) {
				return true
			}
		}
		return false
	}

	for _, vs := range toSequence {
		if !isSequenced(vs) {
			if err := t.db.Create(&vs); err != nil {
				return err
			}
		}
		res = append(res, vs)
	}
	payload.DelegationSequences = res
	return nil
}

func NewDebondingDelegationsSeqCreatorTask(db DebondingDelegationSeqCreatorTaskStore) *debondingDelegationSeqCreatorTask {
	return &debondingDelegationSeqCreatorTask{
		db: db,
	}
}

type debondingDelegationSeqCreatorTask struct {
	db DebondingDelegationSeqCreatorTaskStore
}

type DebondingDelegationSeqCreatorTaskStore interface {
	Create(record interface{}) error
	FindByHeight(h int64) ([]model.DebondingDelegationSeq, error)
}

func (t *debondingDelegationSeqCreatorTask) GetName() string {
	return TaskNameDebondingDelegationSeqCreator
}

func (t *debondingDelegationSeqCreatorTask) Run(ctx context.Context, p pipeline.Payload) error {

	payload := p.(*payload)

	logger.Info(fmt.Sprintf("running indexer task [stage=%s] [task=%s] [height=%d]", pipeline.StageSequencer, t.GetName(), payload.CurrentHeight))

	var res []model.DebondingDelegationSeq
	sequenced, err := t.db.FindByHeight(payload.CurrentHeight)
	if err != nil {
		return err
	}

	toSequence, err := DebondingDelegationToSequence(payload.Syncable, payload.RawState)
	if err != nil {
		return err
	}

	// Nothing to sequence
	if len(toSequence) == 0 {
		payload.DebondingDelegationSequences = res
		return nil
	}

	// Everything sequenced and saved to persistence
	if len(sequenced) == len(toSequence) {
		payload.DebondingDelegationSequences = sequenced
		return nil
	}

	isSequenced := func(vs model.DebondingDelegationSeq) bool {
		for _, sv := range sequenced {
			if sv.Equal(vs) {
				return true
			}
		}
		return false
	}

	for _, vs := range toSequence {
		if !isSequenced(vs) {
			if err := t.db.Create(&vs); err != nil {
				return err
			}
		}
		res = append(res, vs)
	}
	payload.DebondingDelegationSequences = res
	return nil
}
