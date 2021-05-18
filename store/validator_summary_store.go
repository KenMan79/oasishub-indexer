package store

import (
	"fmt"
	"time"

	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/figment-networks/oasishub-indexer/types"
	"github.com/jinzhu/gorm"

	"github.com/figment-networks/oasishub-indexer/model"
)

var (
	_ ValidatorSummaryStore = (*validatorSummaryStore)(nil)
)

type ValidatorSummaryStore interface {
	BaseStore

	Find(*model.ValidatorSummary) (*model.ValidatorSummary, error)
	FindActivityPeriods(types.SummaryInterval, int64) ([]ActivityPeriodRow, error)
	FindAllByTimePeriod(start, end *types.Time, addresses ...string) ([]model.ValidatorSummary, error)
	FindSummary(types.SummaryInterval, string) ([]ValidatorSummaryRow, error)
	FindSummaryByAddress(string, types.SummaryInterval, string) ([]model.ValidatorSummary, error)
	FindMostRecent() (*model.ValidatorSummary, error)
	FindMostRecentByInterval(types.SummaryInterval) (*model.ValidatorSummary, error)
	DeleteOlderThan(types.SummaryInterval, time.Time) (*int64, error)
}

func NewValidatorSummaryStore(db *gorm.DB) *validatorSummaryStore {
	return &validatorSummaryStore{scoped(db, model.ValidatorSummary{})}
}

// validatorSummaryStore handles operations on validators
type validatorSummaryStore struct {
	baseStore
}

// Find find validator summary by query
func (s validatorSummaryStore) Find(query *model.ValidatorSummary) (*model.ValidatorSummary, error) {
	var result model.ValidatorSummary

	err := s.db.
		Where(query).
		First(&result).
		Error

	return &result, checkErr(err)
}

// FindActivityPeriods Finds activity periods
func (s *validatorSummaryStore) FindActivityPeriods(interval types.SummaryInterval, indexVersion int64) ([]ActivityPeriodRow, error) {
	t := metrics.NewTimer(databaseQueryDuration.WithLabels("ValidatorSummaryStore_FindActivityPeriods"))
	defer t.ObserveDuration()

	query := getActivityPeriodsQuery(model.ValidatorSummary{}.TableName())

	var res []ActivityPeriodRow
	return res, s.db.Raw(query, fmt.Sprintf("1%s", interval), interval, indexVersion).Find(&res).Error
}

type ValidatorSummaryRow struct {
	TimeBucket             string         `json:"time_bucket"`
	TimeInterval           string         `json:"time_interval"`
	VotingPowerAvg         float64        `json:"voting_power_avg"`
	VotingPowerMax         float64        `json:"voting_power_max"`
	VotingPowerMin         float64        `json:"voting_power_min"`
	TotalSharesAvg         types.Quantity `json:"total_shares_avg"`
	TotalSharesMax         types.Quantity `json:"total_shares_max"`
	TotalSharesMin         types.Quantity `json:"total_shares_min"`
	ActiveEscrowBalanceAvg types.Quantity `json:"active_escrow_balance_avg"`
	ActiveEscrowBalanceMax types.Quantity `json:"active_escrow_balance_max"`
	ActiveEscrowBalanceMin types.Quantity `json:"active_escrow_balance_min"`
	CommissionAvg          types.Quantity `json:"commission_avg"`
	CommissionMax          types.Quantity `json:"commission_max"`
	CommissionMin          types.Quantity `json:"commission_min"`
	ValidatedSum           int64          `json:"validated_sum"`
	NotValidatedSum        int64          `json:"not_validated_sum"`
	ProposedSum            int64          `json:"proposed_sum"`
	UptimeAvg              float64        `json:"uptime_avg"`
}

// FindSummary gets summary for validator summary
func (s *validatorSummaryStore) FindSummary(interval types.SummaryInterval, period string) ([]ValidatorSummaryRow, error) {
	t := metrics.NewTimer(databaseQueryDuration.WithLabels("ValidatorSummaryStore_FindSummary"))
	defer t.ObserveDuration()

	var res []ValidatorSummaryRow
	return res, s.db.Raw(allValidatorsSummaryForIntervalQuery, interval, period, interval).Find(&res).Error
}

// FindSummaryByAddress gets summary for given validator
func (s *validatorSummaryStore) FindSummaryByAddress(address string, interval types.SummaryInterval, period string) ([]model.ValidatorSummary, error) {
	t := metrics.NewTimer(databaseQueryDuration.WithLabels("ValidatorSummaryStore_FindSummaryByAddress"))
	defer t.ObserveDuration()

	var res []model.ValidatorSummary
	return res, s.db.Raw(validatorSummaryForIntervalQuery, interval, period, address, interval).Find(&res).Error
}

// FindMostRecent finds most recent validator summary
func (s *validatorSummaryStore) FindMostRecent() (*model.ValidatorSummary, error) {
	validatorSummary := &model.ValidatorSummary{}
	err := findMostRecent(s.db, "time_bucket", validatorSummary)
	return validatorSummary, checkErr(err)
}

// FindAllByTimePeriod finds all validator summaries for a given address within the specified start and end times
func (s *validatorSummaryStore) FindAllByTimePeriod(start, end *types.Time, addresses ...string) ([]model.ValidatorSummary, error) {

	tx := s.db.
		Where("time_interval = 'day' AND address In (?)", addresses).
		Order("time_bucket")

	if !end.IsZero() {
		tx = tx.Where("time_bucket <= ?", end)
	}
	if !start.IsZero() {
		tx = tx.Where("time_bucket >= ?", start)
	}

	var res []model.ValidatorSummary
	return res, checkErr(tx.Find(&res).Error)
}

// FindMostRecentByInterval finds most recent validator summary for interval
func (s *validatorSummaryStore) FindMostRecentByInterval(interval types.SummaryInterval) (*model.ValidatorSummary, error) {
	query := &model.ValidatorSummary{
		Summary: &model.Summary{TimeInterval: interval},
	}
	result := model.ValidatorSummary{}

	err := s.db.
		Where(query).
		Order("time_bucket DESC").
		Take(&result).
		Error

	return &result, checkErr(err)
}

// DeleteOlderThan deleted validator summary records older than given threshold
func (s *validatorSummaryStore) DeleteOlderThan(interval types.SummaryInterval, purgeThreshold time.Time) (*int64, error) {
	statement := s.db.
		Unscoped().
		Where("time_interval = ? AND time_bucket < ?", interval, purgeThreshold).
		Delete(&model.ValidatorSummary{})

	if statement.Error != nil {
		return nil, checkErr(statement.Error)
	}

	return &statement.RowsAffected, nil
}
