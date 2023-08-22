// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"snowmirror/internal/sdk/pkg/models/shared"
)

type GetSynchronizationActionResponse struct {
	// OK
	Actions     []shared.Action
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
