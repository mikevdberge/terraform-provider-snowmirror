// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateSynchronizationSyncInputFullLoadScheduler struct {
	BeginDate     types.String `tfsdk:"begin_date"`
	ExecutionType types.String `tfsdk:"execution_type"`
	Type          types.String `tfsdk:"type"`
}