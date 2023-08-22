// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CreateSynchronizationSyncInputColumns struct {
	Name *string `json:"name,omitempty"`
}

type CreateSynchronizationSyncInputColumnsToExclude struct {
	Name *string `json:"name,omitempty"`
}

type CreateSynchronizationSyncInputDeleteStrategy string

const (
	CreateSynchronizationSyncInputDeleteStrategyAudit    CreateSynchronizationSyncInputDeleteStrategy = "AUDIT"
	CreateSynchronizationSyncInputDeleteStrategyTruncate CreateSynchronizationSyncInputDeleteStrategy = "TRUNCATE"
	CreateSynchronizationSyncInputDeleteStrategyDiff     CreateSynchronizationSyncInputDeleteStrategy = "DIFF"
	CreateSynchronizationSyncInputDeleteStrategyNone     CreateSynchronizationSyncInputDeleteStrategy = "NONE"
)

func (e CreateSynchronizationSyncInputDeleteStrategy) ToPointer() *CreateSynchronizationSyncInputDeleteStrategy {
	return &e
}

func (e *CreateSynchronizationSyncInputDeleteStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AUDIT":
		fallthrough
	case "TRUNCATE":
		fallthrough
	case "DIFF":
		fallthrough
	case "NONE":
		*e = CreateSynchronizationSyncInputDeleteStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSynchronizationSyncInputDeleteStrategy: %v", v)
	}
}

type CreateSynchronizationSyncInputFullLoadSchedulerExecutionType string

const (
	CreateSynchronizationSyncInputFullLoadSchedulerExecutionTypeCleanAndSynchronize CreateSynchronizationSyncInputFullLoadSchedulerExecutionType = "CLEAN_AND_SYNCHRONIZE"
	CreateSynchronizationSyncInputFullLoadSchedulerExecutionTypeDifferential        CreateSynchronizationSyncInputFullLoadSchedulerExecutionType = "DIFFERENTIAL."
)

func (e CreateSynchronizationSyncInputFullLoadSchedulerExecutionType) ToPointer() *CreateSynchronizationSyncInputFullLoadSchedulerExecutionType {
	return &e
}

func (e *CreateSynchronizationSyncInputFullLoadSchedulerExecutionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CLEAN_AND_SYNCHRONIZE":
		fallthrough
	case "DIFFERENTIAL.":
		*e = CreateSynchronizationSyncInputFullLoadSchedulerExecutionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSynchronizationSyncInputFullLoadSchedulerExecutionType: %v", v)
	}
}

type CreateSynchronizationSyncInputFullLoadSchedulerType string

const (
	CreateSynchronizationSyncInputFullLoadSchedulerTypeManually     CreateSynchronizationSyncInputFullLoadSchedulerType = "MANUALLY"
	CreateSynchronizationSyncInputFullLoadSchedulerTypeDaily        CreateSynchronizationSyncInputFullLoadSchedulerType = "DAILY"
	CreateSynchronizationSyncInputFullLoadSchedulerTypeWeekly       CreateSynchronizationSyncInputFullLoadSchedulerType = "WEEKLY"
	CreateSynchronizationSyncInputFullLoadSchedulerTypePeriodically CreateSynchronizationSyncInputFullLoadSchedulerType = "PERIODICALLY"
	CreateSynchronizationSyncInputFullLoadSchedulerTypeCron         CreateSynchronizationSyncInputFullLoadSchedulerType = "CRON"
)

func (e CreateSynchronizationSyncInputFullLoadSchedulerType) ToPointer() *CreateSynchronizationSyncInputFullLoadSchedulerType {
	return &e
}

func (e *CreateSynchronizationSyncInputFullLoadSchedulerType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MANUALLY":
		fallthrough
	case "DAILY":
		fallthrough
	case "WEEKLY":
		fallthrough
	case "PERIODICALLY":
		fallthrough
	case "CRON":
		*e = CreateSynchronizationSyncInputFullLoadSchedulerType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSynchronizationSyncInputFullLoadSchedulerType: %v", v)
	}
}

type CreateSynchronizationSyncInputFullLoadScheduler struct {
	BeginDate     *string                                                       `json:"beginDate,omitempty"`
	ExecutionType *CreateSynchronizationSyncInputFullLoadSchedulerExecutionType `json:"executionType,omitempty"`
	Type          *CreateSynchronizationSyncInputFullLoadSchedulerType          `json:"type,omitempty"`
}

type CreateSynchronizationSyncInputSchedulerDays string

const (
	CreateSynchronizationSyncInputSchedulerDaysMonday   CreateSynchronizationSyncInputSchedulerDays = "MONDAY"
	CreateSynchronizationSyncInputSchedulerDaysTuesday  CreateSynchronizationSyncInputSchedulerDays = "TUESDAY"
	CreateSynchronizationSyncInputSchedulerDaysWednesda CreateSynchronizationSyncInputSchedulerDays = "WEDNESDA"
	CreateSynchronizationSyncInputSchedulerDaysThursday CreateSynchronizationSyncInputSchedulerDays = "THURSDAY"
	CreateSynchronizationSyncInputSchedulerDaysFriday   CreateSynchronizationSyncInputSchedulerDays = "FRIDAY"
	CreateSynchronizationSyncInputSchedulerDaysSaturday CreateSynchronizationSyncInputSchedulerDays = "SATURDAY"
	CreateSynchronizationSyncInputSchedulerDaysSunday   CreateSynchronizationSyncInputSchedulerDays = "SUNDAY"
)

func (e CreateSynchronizationSyncInputSchedulerDays) ToPointer() *CreateSynchronizationSyncInputSchedulerDays {
	return &e
}

func (e *CreateSynchronizationSyncInputSchedulerDays) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MONDAY":
		fallthrough
	case "TUESDAY":
		fallthrough
	case "WEDNESDA":
		fallthrough
	case "THURSDAY":
		fallthrough
	case "FRIDAY":
		fallthrough
	case "SATURDAY":
		fallthrough
	case "SUNDAY":
		*e = CreateSynchronizationSyncInputSchedulerDays(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSynchronizationSyncInputSchedulerDays: %v", v)
	}
}

// CreateSynchronizationSyncInputSchedulerType - Specifies when the incremental load synchronization will run
type CreateSynchronizationSyncInputSchedulerType string

const (
	CreateSynchronizationSyncInputSchedulerTypeManually     CreateSynchronizationSyncInputSchedulerType = "MANUALLY"
	CreateSynchronizationSyncInputSchedulerTypeDaily        CreateSynchronizationSyncInputSchedulerType = "DAILY"
	CreateSynchronizationSyncInputSchedulerTypeWeekly       CreateSynchronizationSyncInputSchedulerType = "WEEKLY"
	CreateSynchronizationSyncInputSchedulerTypePeriodically CreateSynchronizationSyncInputSchedulerType = "PERIODICALLY"
	CreateSynchronizationSyncInputSchedulerTypeCron         CreateSynchronizationSyncInputSchedulerType = "CRON"
)

func (e CreateSynchronizationSyncInputSchedulerType) ToPointer() *CreateSynchronizationSyncInputSchedulerType {
	return &e
}

func (e *CreateSynchronizationSyncInputSchedulerType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MANUALLY":
		fallthrough
	case "DAILY":
		fallthrough
	case "WEEKLY":
		fallthrough
	case "PERIODICALLY":
		fallthrough
	case "CRON":
		*e = CreateSynchronizationSyncInputSchedulerType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSynchronizationSyncInputSchedulerType: %v", v)
	}
}

type CreateSynchronizationSyncInputScheduler struct {
	BeginDate            *string                                       `json:"beginDate,omitempty"`
	Days                 []CreateSynchronizationSyncInputSchedulerDays `json:"days,omitempty"`
	IncLoadExecutionType *string                                       `json:"incLoadExecutionType,omitempty"`
	Time                 *string                                       `json:"time,omitempty"`
	// Specifies when the incremental load synchronization will run
	Type    *CreateSynchronizationSyncInputSchedulerType `json:"type,omitempty"`
	Visible *bool                                        `json:"visible,omitempty"`
}

type CreateSynchronizationSyncInputSchedulerPriority string

const (
	CreateSynchronizationSyncInputSchedulerPriorityHighest CreateSynchronizationSyncInputSchedulerPriority = "HIGHEST"
	CreateSynchronizationSyncInputSchedulerPriorityHigh    CreateSynchronizationSyncInputSchedulerPriority = "HIGH"
	CreateSynchronizationSyncInputSchedulerPriorityNormal  CreateSynchronizationSyncInputSchedulerPriority = "NORMAL"
	CreateSynchronizationSyncInputSchedulerPriorityLow     CreateSynchronizationSyncInputSchedulerPriority = "LOW"
	CreateSynchronizationSyncInputSchedulerPriorityLowest  CreateSynchronizationSyncInputSchedulerPriority = "LOWEST"
)

func (e CreateSynchronizationSyncInputSchedulerPriority) ToPointer() *CreateSynchronizationSyncInputSchedulerPriority {
	return &e
}

func (e *CreateSynchronizationSyncInputSchedulerPriority) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HIGHEST":
		fallthrough
	case "HIGH":
		fallthrough
	case "NORMAL":
		fallthrough
	case "LOW":
		fallthrough
	case "LOWEST":
		*e = CreateSynchronizationSyncInputSchedulerPriority(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSynchronizationSyncInputSchedulerPriority: %v", v)
	}
}

type CreateSynchronizationSyncInput struct {
	// true - synchronization is active and can be scheduled to synchronize data from ServiceNow
	// false - synchronization is deactivated and cannot be scheduled to synchronize data from ServiceNowNow
	Active *bool `json:"active,omitempty"`
	// SnowMirror checks if columns exist in ServiceNow. If this flag is set to true,
	AllowInheritedColumns *bool `json:"allowInheritedColumns,omitempty"`
	// Configures how to check for schema changes in ServiceNow.
	//
	// Enabled (true) - whenever a synchronization is executed, SnowMirror checks for schema changes in ServiceNow. Automatically adds, updates (data type, max. length of a column) and removes columns. If a new column is added SnowMirror clears the mirror table and downloads all records from scratch.
	// Enabled (no truncation) (ENABLED_WITHOUT_TRUNCATION) - the same as Enabled option. It handles new columns differently, though. If a new column is added SnowMirror does not clear the mirror table. Instead, it creates the column and populates it with a default value (which is defined in ServiceNow).
	//
	AutoSchemaUpdate  *string                                          `json:"autoSchemaUpdate,omitempty"`
	Columns           []CreateSynchronizationSyncInputColumns          `json:"columns,omitempty"`
	ColumnsToExclude  []CreateSynchronizationSyncInputColumnsToExclude `json:"columnsToExclude,omitempty"`
	DeleteStrategy    *CreateSynchronizationSyncInputDeleteStrategy    `json:"deleteStrategy,omitempty"`
	EncodedQuery      *string                                          `json:"encodedQuery,omitempty"`
	FullLoadScheduler *CreateSynchronizationSyncInputFullLoadScheduler `json:"fullLoadScheduler,omitempty"`
	// Id of the synchronization.
	ID *int64 `json:"id,omitempty"`
	// Name of the table in mirror database where the data will be migrated.
	MirrorTable *string `json:"mirrorTable,omitempty"`
	// Display name of the synchronization.
	Name string `json:"name"`
	// Defines how to synchronize reference field types.
	ReferenceFieldType *string `json:"referenceFieldType,omitempty"`
	// Determines whether initial synchronization should be done
	RunImmediately    *bool                                            `json:"runImmediately,omitempty"`
	Scheduler         *CreateSynchronizationSyncInputScheduler         `json:"scheduler,omitempty"`
	SchedulerPriority *CreateSynchronizationSyncInputSchedulerPriority `json:"schedulerPriority,omitempty"`
	// Name of the table in ServiceNow.
	Table *string `json:"table,omitempty"`
	// Name of the view in ServiceNow.
	View *string `json:"view,omitempty"`
}