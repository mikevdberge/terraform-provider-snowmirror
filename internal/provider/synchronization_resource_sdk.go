// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"snowmirror/internal/sdk/pkg/models/shared"
)

func (r *SynchronizationResourceModel) ToCreateSDKType() *shared.CreateSynchronizationSyncInput {
	active := new(bool)
	if !r.Active.IsUnknown() && !r.Active.IsNull() {
		*active = r.Active.ValueBool()
	} else {
		active = nil
	}
	allowInheritedColumns := new(bool)
	if !r.AllowInheritedColumns.IsUnknown() && !r.AllowInheritedColumns.IsNull() {
		*allowInheritedColumns = r.AllowInheritedColumns.ValueBool()
	} else {
		allowInheritedColumns = nil
	}
	autoSchemaUpdate := new(string)
	if !r.AutoSchemaUpdate.IsUnknown() && !r.AutoSchemaUpdate.IsNull() {
		*autoSchemaUpdate = r.AutoSchemaUpdate.ValueString()
	} else {
		autoSchemaUpdate = nil
	}
	var columns []shared.CreateSynchronizationSyncInputColumns = nil
	for _, columnsItem := range r.Columns {
		name := new(string)
		if !columnsItem.Name.IsUnknown() && !columnsItem.Name.IsNull() {
			*name = columnsItem.Name.ValueString()
		} else {
			name = nil
		}
		columns = append(columns, shared.CreateSynchronizationSyncInputColumns{
			Name: name,
		})
	}
	var columnsToExclude []shared.CreateSynchronizationSyncInputColumnsToExclude = nil
	for _, columnsToExcludeItem := range r.ColumnsToExclude {
		name1 := new(string)
		if !columnsToExcludeItem.Name.IsUnknown() && !columnsToExcludeItem.Name.IsNull() {
			*name1 = columnsToExcludeItem.Name.ValueString()
		} else {
			name1 = nil
		}
		columnsToExclude = append(columnsToExclude, shared.CreateSynchronizationSyncInputColumnsToExclude{
			Name: name1,
		})
	}
	deleteStrategy := new(shared.CreateSynchronizationSyncInputDeleteStrategy)
	if !r.DeleteStrategy.IsUnknown() && !r.DeleteStrategy.IsNull() {
		*deleteStrategy = shared.CreateSynchronizationSyncInputDeleteStrategy(r.DeleteStrategy.ValueString())
	} else {
		deleteStrategy = nil
	}
	encodedQuery := new(string)
	if !r.EncodedQuery.IsUnknown() && !r.EncodedQuery.IsNull() {
		*encodedQuery = r.EncodedQuery.ValueString()
	} else {
		encodedQuery = nil
	}
	var fullLoadScheduler *shared.CreateSynchronizationSyncInputFullLoadScheduler
	if r.FullLoadScheduler != nil {
		beginDate := new(string)
		if !r.FullLoadScheduler.BeginDate.IsUnknown() && !r.FullLoadScheduler.BeginDate.IsNull() {
			*beginDate = r.FullLoadScheduler.BeginDate.ValueString()
		} else {
			beginDate = nil
		}
		executionType := new(shared.CreateSynchronizationSyncInputFullLoadSchedulerExecutionType)
		if !r.FullLoadScheduler.ExecutionType.IsUnknown() && !r.FullLoadScheduler.ExecutionType.IsNull() {
			*executionType = shared.CreateSynchronizationSyncInputFullLoadSchedulerExecutionType(r.FullLoadScheduler.ExecutionType.ValueString())
		} else {
			executionType = nil
		}
		typeVar := new(shared.CreateSynchronizationSyncInputFullLoadSchedulerType)
		if !r.FullLoadScheduler.Type.IsUnknown() && !r.FullLoadScheduler.Type.IsNull() {
			*typeVar = shared.CreateSynchronizationSyncInputFullLoadSchedulerType(r.FullLoadScheduler.Type.ValueString())
		} else {
			typeVar = nil
		}
		fullLoadScheduler = &shared.CreateSynchronizationSyncInputFullLoadScheduler{
			BeginDate:     beginDate,
			ExecutionType: executionType,
			Type:          typeVar,
		}
	}
	id := new(int64)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueInt64()
	} else {
		id = nil
	}
	mirrorTable := new(string)
	if !r.MirrorTable.IsUnknown() && !r.MirrorTable.IsNull() {
		*mirrorTable = r.MirrorTable.ValueString()
	} else {
		mirrorTable = nil
	}
	name2 := r.Name.ValueString()
	referenceFieldType := new(string)
	if !r.ReferenceFieldType.IsUnknown() && !r.ReferenceFieldType.IsNull() {
		*referenceFieldType = r.ReferenceFieldType.ValueString()
	} else {
		referenceFieldType = nil
	}
	runImmediately := new(bool)
	if !r.RunImmediately.IsUnknown() && !r.RunImmediately.IsNull() {
		*runImmediately = r.RunImmediately.ValueBool()
	} else {
		runImmediately = nil
	}
	var scheduler *shared.CreateSynchronizationSyncInputScheduler
	if r.Scheduler != nil {
		beginDate1 := new(string)
		if !r.Scheduler.BeginDate.IsUnknown() && !r.Scheduler.BeginDate.IsNull() {
			*beginDate1 = r.Scheduler.BeginDate.ValueString()
		} else {
			beginDate1 = nil
		}
		var days []shared.CreateSynchronizationSyncInputSchedulerDays = nil
		for _, daysItem := range r.Scheduler.Days {
			days = append(days, shared.CreateSynchronizationSyncInputSchedulerDays(daysItem.ValueString()))
		}
		incLoadExecutionType := new(string)
		if !r.Scheduler.IncLoadExecutionType.IsUnknown() && !r.Scheduler.IncLoadExecutionType.IsNull() {
			*incLoadExecutionType = r.Scheduler.IncLoadExecutionType.ValueString()
		} else {
			incLoadExecutionType = nil
		}
		time := new(string)
		if !r.Scheduler.Time.IsUnknown() && !r.Scheduler.Time.IsNull() {
			*time = r.Scheduler.Time.ValueString()
		} else {
			time = nil
		}
		typeVar1 := new(shared.CreateSynchronizationSyncInputSchedulerType)
		if !r.Scheduler.Type.IsUnknown() && !r.Scheduler.Type.IsNull() {
			*typeVar1 = shared.CreateSynchronizationSyncInputSchedulerType(r.Scheduler.Type.ValueString())
		} else {
			typeVar1 = nil
		}
		visible := new(bool)
		if !r.Scheduler.Visible.IsUnknown() && !r.Scheduler.Visible.IsNull() {
			*visible = r.Scheduler.Visible.ValueBool()
		} else {
			visible = nil
		}
		scheduler = &shared.CreateSynchronizationSyncInputScheduler{
			BeginDate:            beginDate1,
			Days:                 days,
			IncLoadExecutionType: incLoadExecutionType,
			Time:                 time,
			Type:                 typeVar1,
			Visible:              visible,
		}
	}
	schedulerPriority := new(shared.CreateSynchronizationSyncInputSchedulerPriority)
	if !r.SchedulerPriority.IsUnknown() && !r.SchedulerPriority.IsNull() {
		*schedulerPriority = shared.CreateSynchronizationSyncInputSchedulerPriority(r.SchedulerPriority.ValueString())
	} else {
		schedulerPriority = nil
	}
	table := new(string)
	if !r.Table.IsUnknown() && !r.Table.IsNull() {
		*table = r.Table.ValueString()
	} else {
		table = nil
	}
	view := new(string)
	if !r.View.IsUnknown() && !r.View.IsNull() {
		*view = r.View.ValueString()
	} else {
		view = nil
	}
	out := shared.CreateSynchronizationSyncInput{
		Active:                active,
		AllowInheritedColumns: allowInheritedColumns,
		AutoSchemaUpdate:      autoSchemaUpdate,
		Columns:               columns,
		ColumnsToExclude:      columnsToExclude,
		DeleteStrategy:        deleteStrategy,
		EncodedQuery:          encodedQuery,
		FullLoadScheduler:     fullLoadScheduler,
		ID:                    id,
		MirrorTable:           mirrorTable,
		Name:                  name2,
		ReferenceFieldType:    referenceFieldType,
		RunImmediately:        runImmediately,
		Scheduler:             scheduler,
		SchedulerPriority:     schedulerPriority,
		Table:                 table,
		View:                  view,
	}
	return &out
}

func (r *SynchronizationResourceModel) ToGetSDKType() *shared.CreateSynchronizationSyncInput {
	out := r.ToCreateSDKType()
	return out
}

func (r *SynchronizationResourceModel) ToUpdateSDKType() *shared.CreateSynchronizationSyncUpdate {
	active := new(bool)
	if !r.Active.IsUnknown() && !r.Active.IsNull() {
		*active = r.Active.ValueBool()
	} else {
		active = nil
	}
	allowInheritedColumns := new(bool)
	if !r.AllowInheritedColumns.IsUnknown() && !r.AllowInheritedColumns.IsNull() {
		*allowInheritedColumns = r.AllowInheritedColumns.ValueBool()
	} else {
		allowInheritedColumns = nil
	}
	autoSchemaUpdate := new(string)
	if !r.AutoSchemaUpdate.IsUnknown() && !r.AutoSchemaUpdate.IsNull() {
		*autoSchemaUpdate = r.AutoSchemaUpdate.ValueString()
	} else {
		autoSchemaUpdate = nil
	}
	var columns []shared.CreateSynchronizationSyncUpdateColumns = nil
	for _, columnsItem := range r.Columns {
		name := new(string)
		if !columnsItem.Name.IsUnknown() && !columnsItem.Name.IsNull() {
			*name = columnsItem.Name.ValueString()
		} else {
			name = nil
		}
		columns = append(columns, shared.CreateSynchronizationSyncUpdateColumns{
			Name: name,
		})
	}
	var columnsToExclude []shared.CreateSynchronizationSyncUpdateColumnsToExclude = nil
	for _, columnsToExcludeItem := range r.ColumnsToExclude {
		name1 := new(string)
		if !columnsToExcludeItem.Name.IsUnknown() && !columnsToExcludeItem.Name.IsNull() {
			*name1 = columnsToExcludeItem.Name.ValueString()
		} else {
			name1 = nil
		}
		columnsToExclude = append(columnsToExclude, shared.CreateSynchronizationSyncUpdateColumnsToExclude{
			Name: name1,
		})
	}
	deleteStrategy := new(shared.CreateSynchronizationSyncUpdateDeleteStrategy)
	if !r.DeleteStrategy.IsUnknown() && !r.DeleteStrategy.IsNull() {
		*deleteStrategy = shared.CreateSynchronizationSyncUpdateDeleteStrategy(r.DeleteStrategy.ValueString())
	} else {
		deleteStrategy = nil
	}
	encodedQuery := new(string)
	if !r.EncodedQuery.IsUnknown() && !r.EncodedQuery.IsNull() {
		*encodedQuery = r.EncodedQuery.ValueString()
	} else {
		encodedQuery = nil
	}
	var fullLoadScheduler *shared.CreateSynchronizationSyncUpdateFullLoadScheduler
	if r.FullLoadScheduler != nil {
		beginDate := new(string)
		if !r.FullLoadScheduler.BeginDate.IsUnknown() && !r.FullLoadScheduler.BeginDate.IsNull() {
			*beginDate = r.FullLoadScheduler.BeginDate.ValueString()
		} else {
			beginDate = nil
		}
		executionType := new(shared.CreateSynchronizationSyncUpdateFullLoadSchedulerExecutionType)
		if !r.FullLoadScheduler.ExecutionType.IsUnknown() && !r.FullLoadScheduler.ExecutionType.IsNull() {
			*executionType = shared.CreateSynchronizationSyncUpdateFullLoadSchedulerExecutionType(r.FullLoadScheduler.ExecutionType.ValueString())
		} else {
			executionType = nil
		}
		typeVar := new(shared.CreateSynchronizationSyncUpdateFullLoadSchedulerType)
		if !r.FullLoadScheduler.Type.IsUnknown() && !r.FullLoadScheduler.Type.IsNull() {
			*typeVar = shared.CreateSynchronizationSyncUpdateFullLoadSchedulerType(r.FullLoadScheduler.Type.ValueString())
		} else {
			typeVar = nil
		}
		fullLoadScheduler = &shared.CreateSynchronizationSyncUpdateFullLoadScheduler{
			BeginDate:     beginDate,
			ExecutionType: executionType,
			Type:          typeVar,
		}
	}
	id := new(int64)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueInt64()
	} else {
		id = nil
	}
	mirrorTable := new(string)
	if !r.MirrorTable.IsUnknown() && !r.MirrorTable.IsNull() {
		*mirrorTable = r.MirrorTable.ValueString()
	} else {
		mirrorTable = nil
	}
	name2 := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name2 = r.Name.ValueString()
	} else {
		name2 = nil
	}
	referenceFieldType := new(string)
	if !r.ReferenceFieldType.IsUnknown() && !r.ReferenceFieldType.IsNull() {
		*referenceFieldType = r.ReferenceFieldType.ValueString()
	} else {
		referenceFieldType = nil
	}
	var scheduler *shared.CreateSynchronizationSyncUpdateScheduler
	if r.Scheduler != nil {
		beginDate1 := new(string)
		if !r.Scheduler.BeginDate.IsUnknown() && !r.Scheduler.BeginDate.IsNull() {
			*beginDate1 = r.Scheduler.BeginDate.ValueString()
		} else {
			beginDate1 = nil
		}
		var days []shared.CreateSynchronizationSyncUpdateSchedulerDays = nil
		for _, daysItem := range r.Scheduler.Days {
			days = append(days, shared.CreateSynchronizationSyncUpdateSchedulerDays(daysItem.ValueString()))
		}
		incLoadExecutionType := new(string)
		if !r.Scheduler.IncLoadExecutionType.IsUnknown() && !r.Scheduler.IncLoadExecutionType.IsNull() {
			*incLoadExecutionType = r.Scheduler.IncLoadExecutionType.ValueString()
		} else {
			incLoadExecutionType = nil
		}
		time := new(string)
		if !r.Scheduler.Time.IsUnknown() && !r.Scheduler.Time.IsNull() {
			*time = r.Scheduler.Time.ValueString()
		} else {
			time = nil
		}
		typeVar1 := new(shared.CreateSynchronizationSyncUpdateSchedulerType)
		if !r.Scheduler.Type.IsUnknown() && !r.Scheduler.Type.IsNull() {
			*typeVar1 = shared.CreateSynchronizationSyncUpdateSchedulerType(r.Scheduler.Type.ValueString())
		} else {
			typeVar1 = nil
		}
		visible := new(bool)
		if !r.Scheduler.Visible.IsUnknown() && !r.Scheduler.Visible.IsNull() {
			*visible = r.Scheduler.Visible.ValueBool()
		} else {
			visible = nil
		}
		scheduler = &shared.CreateSynchronizationSyncUpdateScheduler{
			BeginDate:            beginDate1,
			Days:                 days,
			IncLoadExecutionType: incLoadExecutionType,
			Time:                 time,
			Type:                 typeVar1,
			Visible:              visible,
		}
	}
	schedulerPriority := new(shared.CreateSynchronizationSyncUpdateSchedulerPriority)
	if !r.SchedulerPriority.IsUnknown() && !r.SchedulerPriority.IsNull() {
		*schedulerPriority = shared.CreateSynchronizationSyncUpdateSchedulerPriority(r.SchedulerPriority.ValueString())
	} else {
		schedulerPriority = nil
	}
	table := new(string)
	if !r.Table.IsUnknown() && !r.Table.IsNull() {
		*table = r.Table.ValueString()
	} else {
		table = nil
	}
	view := new(string)
	if !r.View.IsUnknown() && !r.View.IsNull() {
		*view = r.View.ValueString()
	} else {
		view = nil
	}
	out := shared.CreateSynchronizationSyncUpdate{
		Active:                active,
		AllowInheritedColumns: allowInheritedColumns,
		AutoSchemaUpdate:      autoSchemaUpdate,
		Columns:               columns,
		ColumnsToExclude:      columnsToExclude,
		DeleteStrategy:        deleteStrategy,
		EncodedQuery:          encodedQuery,
		FullLoadScheduler:     fullLoadScheduler,
		ID:                    id,
		MirrorTable:           mirrorTable,
		Name:                  name2,
		ReferenceFieldType:    referenceFieldType,
		Scheduler:             scheduler,
		SchedulerPriority:     schedulerPriority,
		Table:                 table,
		View:                  view,
	}
	return &out
}

func (r *SynchronizationResourceModel) ToDeleteSDKType() *shared.CreateSynchronizationSyncInput {
	out := r.ToCreateSDKType()
	return out
}

func (r *SynchronizationResourceModel) RefreshFromGetResponse(resp *shared.SyncronizationSyncOutput) {
	if resp.Active != nil {
		r.Active = types.BoolValue(*resp.Active)
	} else {
		r.Active = types.BoolNull()
	}
	if resp.AllowInheritedColumns != nil {
		r.AllowInheritedColumns = types.BoolValue(*resp.AllowInheritedColumns)
	} else {
		r.AllowInheritedColumns = types.BoolNull()
	}
	if resp.AttachmentDirectory != nil {
		r.AttachmentDirectory = types.StringValue(*resp.AttachmentDirectory)
	} else {
		r.AttachmentDirectory = types.StringNull()
	}
	if resp.AutoSchemaUpdate != nil {
		r.AutoSchemaUpdate = types.StringValue(*resp.AutoSchemaUpdate)
	} else {
		r.AutoSchemaUpdate = types.StringNull()
	}
	if resp.DeleteStrategy != nil {
		r.DeleteStrategy = types.StringValue(string(*resp.DeleteStrategy))
	} else {
		r.DeleteStrategy = types.StringNull()
	}
	if resp.EncodedQuery != nil {
		r.EncodedQuery = types.StringValue(*resp.EncodedQuery)
	} else {
		r.EncodedQuery = types.StringNull()
	}
	if resp.Format != nil {
		r.Format = types.StringValue(string(*resp.Format))
	} else {
		r.Format = types.StringNull()
	}
	if r.FullLoadScheduler == nil {
		r.FullLoadScheduler = &SyncronizationSyncOutputFullLoadScheduler{}
	}
	if resp.FullLoadScheduler == nil {
		r.FullLoadScheduler = nil
	} else {
		r.FullLoadScheduler = &SyncronizationSyncOutputFullLoadScheduler{}
		if resp.FullLoadScheduler.BeginDate != nil {
			r.FullLoadScheduler.BeginDate = types.StringValue(*resp.FullLoadScheduler.BeginDate)
		} else {
			r.FullLoadScheduler.BeginDate = types.StringNull()
		}
		if resp.FullLoadScheduler.ExecutionType != nil {
			r.FullLoadScheduler.ExecutionType = types.StringValue(string(*resp.FullLoadScheduler.ExecutionType))
		} else {
			r.FullLoadScheduler.ExecutionType = types.StringNull()
		}
		if resp.FullLoadScheduler.Time != nil {
			r.FullLoadScheduler.Time = types.StringValue(*resp.FullLoadScheduler.Time)
		} else {
			r.FullLoadScheduler.Time = types.StringNull()
		}
		if resp.FullLoadScheduler.Type != nil {
			r.FullLoadScheduler.Type = types.StringValue(string(*resp.FullLoadScheduler.Type))
		} else {
			r.FullLoadScheduler.Type = types.StringNull()
		}
		if resp.FullLoadScheduler.Visible != nil {
			r.FullLoadScheduler.Visible = types.BoolValue(*resp.FullLoadScheduler.Visible)
		} else {
			r.FullLoadScheduler.Visible = types.BoolNull()
		}
	}
	if resp.ID != nil {
		r.ID = types.Int64Value(*resp.ID)
	} else {
		r.ID = types.Int64Null()
	}
	if resp.MasterTable != nil {
		r.MasterTable = types.StringValue(*resp.MasterTable)
	} else {
		r.MasterTable = types.StringNull()
	}
	if resp.MirrorTable != nil {
		r.MirrorTable = types.StringValue(*resp.MirrorTable)
	} else {
		r.MirrorTable = types.StringNull()
	}
	if resp.Name != nil {
		r.Name = types.StringValue(*resp.Name)
	} else {
		r.Name = types.StringNull()
	}
	if resp.ReferenceFieldType != nil {
		r.ReferenceFieldType = types.StringValue(*resp.ReferenceFieldType)
	} else {
		r.ReferenceFieldType = types.StringNull()
	}
	if resp.RetentionPeriod != nil {
		r.RetentionPeriod = types.Int64Value(*resp.RetentionPeriod)
	} else {
		r.RetentionPeriod = types.Int64Null()
	}
	if r.Scheduler == nil {
		r.Scheduler = &SyncronizationSyncOutputScheduler{}
	}
	if resp.Scheduler == nil {
		r.Scheduler = nil
	} else {
		r.Scheduler = &SyncronizationSyncOutputScheduler{}
		if resp.Scheduler.BeginDate != nil {
			r.Scheduler.BeginDate = types.StringValue(*resp.Scheduler.BeginDate)
		} else {
			r.Scheduler.BeginDate = types.StringNull()
		}
		if resp.Scheduler.IncLoadExecutionType != nil {
			r.Scheduler.IncLoadExecutionType = types.StringValue(*resp.Scheduler.IncLoadExecutionType)
		} else {
			r.Scheduler.IncLoadExecutionType = types.StringNull()
		}
		if resp.Scheduler.Time != nil {
			r.Scheduler.Time = types.StringValue(*resp.Scheduler.Time)
		} else {
			r.Scheduler.Time = types.StringNull()
		}
		if resp.Scheduler.Type != nil {
			r.Scheduler.Type = types.StringValue(string(*resp.Scheduler.Type))
		} else {
			r.Scheduler.Type = types.StringNull()
		}
		if resp.Scheduler.Visible != nil {
			r.Scheduler.Visible = types.BoolValue(*resp.Scheduler.Visible)
		} else {
			r.Scheduler.Visible = types.BoolNull()
		}
	}
	if resp.SchedulerPriority != nil {
		r.SchedulerPriority = types.StringValue(string(*resp.SchedulerPriority))
	} else {
		r.SchedulerPriority = types.StringNull()
	}
	if resp.SynchronizationType != nil {
		r.SynchronizationType = types.StringValue(*resp.SynchronizationType)
	} else {
		r.SynchronizationType = types.StringNull()
	}
	if resp.Table != nil {
		r.Table = types.StringValue(*resp.Table)
	} else {
		r.Table = types.StringNull()
	}
	if resp.UpdateBeforeSynchronizationRun != nil {
		r.UpdateBeforeSynchronizationRun = types.StringValue(*resp.UpdateBeforeSynchronizationRun)
	} else {
		r.UpdateBeforeSynchronizationRun = types.StringNull()
	}
	if resp.View != nil {
		r.View = types.StringValue(*resp.View)
	} else {
		r.View = types.StringNull()
	}
}

func (r *SynchronizationResourceModel) RefreshFromCreateResponse(resp *shared.SyncronizationSyncOutput) {
	r.RefreshFromGetResponse(resp)
}

func (r *SynchronizationResourceModel) RefreshFromUpdateResponse(resp *shared.SyncronizationSyncOutput) {
	r.RefreshFromGetResponse(resp)
}