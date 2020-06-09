package store

import (
	"github.com/figment-networks/oasishub-indexer/types"
	"github.com/jinzhu/gorm"
	"time"

	"github.com/figment-networks/oasishub-indexer/model"
)

func NewValidatorSummaryStore(db *gorm.DB) *ValidatorSummaryStore {
	return &ValidatorSummaryStore{scoped(db, model.ValidatorSummary{})}
}

// ValidatorSummaryStore handles operations on validators
type ValidatorSummaryStore struct {
	baseStore
}

// Find find validator summary by query
func (s ValidatorSummaryStore) Find(query *model.ValidatorSummary) (*model.ValidatorSummary, error) {
	var result model.ValidatorSummary

	err := s.db.
		Where(query).
		First(&result).
		Error

	return &result, checkErr(err)
}

// FindByEntityUID finds last validator summary for given entity uid
func (s ValidatorSummaryStore) FindByEntityUID(key string) ([]model.ValidatorSummary, error) {
	q := model.ValidatorSummary{
		EntityUID: key,
	}
	var result []model.ValidatorSummary

	err := s.db.
		Where(&q).
		Find(&result).
		Error

	return result, checkErr(err)
}

type ValidatorSummaryRow struct {
	TimeBucket      string         `json:"time_bucket"`
	TimeInterval    string         `json:"time_interval"`
	VotingPowerAvg  float64        `json:"voting_power_avg"`
	VotingPowerMax  float64        `json:"voting_power_max"`
	VotingPowerMin  float64        `json:"voting_power_min"`
	TotalSharesAvg  types.Quantity `json:"total_shares_avg"`
	TotalSharesMax  types.Quantity `json:"total_shares_max"`
	TotalSharesMin  types.Quantity `json:"total_shares_min"`
	ValidatedSum    int64          `json:"validated_sum"`
	NotValidatedSum int64          `json:"not_validated_sum"`
	ProposedSum     int64          `json:"proposed_sum"`
	UptimeAvg       float64        `json:"uptime_avg"`
}

// FindSummary gets summary for validator summary
func (s *ValidatorSummaryStore) FindSummary(interval types.SummaryInterval, period string) ([]ValidatorSummaryRow, error) {
	defer logQueryDuration(time.Now(), "ValidatorSummaryStore_FindSummary")

	rows, err := s.db.
		Raw(allValidatorsSummaryForIntervalQuery, interval, period, interval).
		Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []ValidatorSummaryRow
	for rows.Next() {
		var row ValidatorSummaryRow
		if err := s.db.ScanRows(rows, &row); err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil
}

// FindSummaryByEntityUID gets summary for given validator
func (s *ValidatorSummaryStore) FindSummaryByEntityUID(key string, interval types.SummaryInterval, period string) ([]model.ValidatorSummary, error) {
	defer logQueryDuration(time.Now(), "ValidatorSummaryStore_GetSummaryByEntityUID")

	rows, err := s.db.Raw(validatorSummaryForIntervalQuery, interval, period, key, interval).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []model.ValidatorSummary
	for rows.Next() {
		var row model.ValidatorSummary
		if err := s.db.ScanRows(rows, &row); err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil
}

// FindMostRecent finds most recent validator summary
func (s *ValidatorSummaryStore) FindMostRecent() (*model.ValidatorSummary, error) {
	validatorSummary := &model.ValidatorSummary{}
	err := findMostRecent(s.db, "time_bucket", validatorSummary)
	return validatorSummary, checkErr(err)
}

// FindMostRecentByInterval finds most recent validator summary for interval
func (s *ValidatorSummaryStore) FindMostRecentByInterval(interval types.SummaryInterval) (*model.ValidatorSummary, error) {
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
func (s *ValidatorSummaryStore) DeleteOlderThan(interval types.SummaryInterval, purgeThreshold time.Time) (*int64, error) {
	statement := s.db.
		Unscoped().
		Where("time_interval = ? AND time_bucket < ?", interval, purgeThreshold).
		Delete(&model.ValidatorSummary{})

	if statement.Error != nil {
		return nil, checkErr(statement.Error)
	}

	return &statement.RowsAffected, nil
}
