// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"snowmirror/internal/sdk"
	"snowmirror/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SynchronizationDataSource{}
var _ datasource.DataSourceWithConfigure = &SynchronizationDataSource{}

func NewSynchronizationDataSource() datasource.DataSource {
	return &SynchronizationDataSource{}
}

// SynchronizationDataSource is the data source implementation.
type SynchronizationDataSource struct {
	client *sdk.Snowmirror
}

// SynchronizationDataSourceModel describes the data model.
type SynchronizationDataSourceModel struct {
	Active                         types.Bool                                 `tfsdk:"active"`
	AllowInheritedColumns          types.Bool                                 `tfsdk:"allow_inherited_columns"`
	AttachmentDirectory            types.String                               `tfsdk:"attachment_directory"`
	AutoSchemaUpdate               types.String                               `tfsdk:"auto_schema_update"`
	Columns                        []CreateSynchronizationSyncInputColumns    `tfsdk:"columns"`
	ColumnsToExclude               []CreateSynchronizationSyncInputColumns    `tfsdk:"columns_to_exclude"`
	DeleteStrategy                 types.String                               `tfsdk:"delete_strategy"`
	EncodedQuery                   types.String                               `tfsdk:"encoded_query"`
	Format                         types.String                               `tfsdk:"format"`
	FullLoadScheduler              *SyncronizationSyncOutputFullLoadScheduler `tfsdk:"full_load_scheduler"`
	ID                             types.Int64                                `tfsdk:"id"`
	MasterTable                    types.String                               `tfsdk:"master_table"`
	MirrorTable                    types.String                               `tfsdk:"mirror_table"`
	Name                           types.String                               `tfsdk:"name"`
	ReferenceFieldType             types.String                               `tfsdk:"reference_field_type"`
	RetentionPeriod                types.Int64                                `tfsdk:"retention_period"`
	RunImmediately                 types.Bool                                 `tfsdk:"run_immediately"`
	Scheduler                      *SyncronizationSyncOutputScheduler         `tfsdk:"scheduler"`
	SchedulerPriority              types.String                               `tfsdk:"scheduler_priority"`
	SynchronizationType            types.String                               `tfsdk:"synchronization_type"`
	Table                          types.String                               `tfsdk:"table"`
	UpdateBeforeSynchronizationRun types.String                               `tfsdk:"update_before_synchronization_run"`
	View                           types.String                               `tfsdk:"view"`
}

// Metadata returns the data source type name.
func (r *SynchronizationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_synchronization"
}

// Schema defines the schema for the data source.
func (r *SynchronizationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Synchronization DataSource",

		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Computed: true,
				MarkdownDescription: `true - synchronization is active and can be scheduled to synchronize data from ServiceNow` + "\n" +
					`false - synchronization is deactivated and cannot be scheduled to synchronize data from ServiceNowNow          `,
			},
			"allow_inherited_columns": schema.BoolAttribute{
				Computed:    true,
				Description: `SnowMirror checks if columns exist in ServiceNow. If this flag is set to true,`,
			},
			"attachment_directory": schema.StringAttribute{
				Computed: true,
			},
			"auto_schema_update": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: `Configures how to check for schema changes in ServiceNow.` + "\n" +
					`` + "\n" +
					`Enabled (true) - whenever a synchronization is executed, SnowMirror checks for schema changes in ServiceNow. Automatically adds, updates (data type, max. length of a column) and removes columns. If a new column is added SnowMirror clears the mirror table and downloads all records from scratch.` + "\n" +
					`Enabled (no truncation) (ENABLED_WITHOUT_TRUNCATION) - the same as Enabled option. It handles new columns differently, though. If a new column is added SnowMirror does not clear the mirror table. Instead, it creates the column and populates it with a default value (which is defined in ServiceNow).` + "\n" +
					``,
			},
			"columns": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"columns_to_exclude": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"delete_strategy": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"AUDIT",
						"TRUNCATE",
						"DIFF",
						"NONE",
					),
				},
				Description: `must be one of ["AUDIT", "TRUNCATE", "DIFF", "NONE"]`,
			},
			"encoded_query": schema.StringAttribute{
				Computed: true,
			},
			"format": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"CSV",
						"XML",
					),
				},
				MarkdownDescription: `must be one of ["CSV", "XML"]` + "\n" +
					`How to store backups. "CSV" - comma separated file. "XML" - XML files.`,
			},
			"full_load_scheduler": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"begin_date": schema.StringAttribute{
						Computed: true,
					},
					"execution_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"CLEAN_AND_SYNCHRONIZE",
								"DIFFERENTIAL.",
							),
						},
						Description: `must be one of ["CLEAN_AND_SYNCHRONIZE", "DIFFERENTIAL."]`,
					},
					"time": schema.StringAttribute{
						Computed: true,
					},
					"type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"MANUALLY",
								"DAILY",
								"WEEKLY",
								"PERIODICALLY",
								"CRON",
							),
						},
						Description: `must be one of ["MANUALLY", "DAILY", "WEEKLY", "PERIODICALLY", "CRON"]`,
					},
					"visible": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
			"id": schema.Int64Attribute{
				Optional:    true,
				Description: `Id of the synchronization.`,
			},
			"master_table": schema.StringAttribute{
				Computed: true,
			},
			"mirror_table": schema.StringAttribute{
				Computed:    true,
				Description: `Name of the table in mirror database where the data will be migrated.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `Display name of the synchronization.`,
			},
			"reference_field_type": schema.StringAttribute{
				Computed:    true,
				Description: `Defines how to synchronize reference field types.`,
			},
			"retention_period": schema.Int64Attribute{
				Computed:    true,
				Description: `How many days to keep backups`,
			},
			"run_immediately": schema.BoolAttribute{
				Optional:    true,
				Description: `Determines whether initial synchronization should be done`,
			},
			"scheduler": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"begin_date": schema.StringAttribute{
						Computed: true,
					},
					"days": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"inc_load_execution_type": schema.StringAttribute{
						Computed: true,
					},
					"time": schema.StringAttribute{
						Computed: true,
					},
					"type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"MANUALLY",
								"DAILY",
								"WEEKLY",
								"PERIODICALLY",
								"CRON",
							),
						},
						MarkdownDescription: `must be one of ["MANUALLY", "DAILY", "WEEKLY", "PERIODICALLY", "CRON"]` + "\n" +
							`Specifies when the incremental load synchronization will run`,
					},
					"visible": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
			"scheduler_priority": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"HIGHEST",
						"HIGH",
						"NORMAL",
						"LOW",
						"LOWEST",
					),
				},
				Description: `must be one of ["HIGHEST", "HIGH", "NORMAL", "LOW", "LOWEST"]`,
			},
			"synchronization_type": schema.StringAttribute{
				Computed: true,
			},
			"table": schema.StringAttribute{
				Computed:    true,
				Description: `Name of the table in ServiceNow.`,
			},
			"update_before_synchronization_run": schema.StringAttribute{
				Computed: true,
			},
			"view": schema.StringAttribute{
				Computed:    true,
				Description: `Name of the view in ServiceNow.`,
			},
		},
	}
}

func (r *SynchronizationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Snowmirror)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Snowmirror, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SynchronizationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SynchronizationDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueInt64()
	request := operations.GetSynchronizationRequest{
		ID: id,
	}
	res, err := r.client.Synchronization.GetSynchronization(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Synchronization == nil || res.Synchronization.Sync == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.Synchronization.Sync)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
