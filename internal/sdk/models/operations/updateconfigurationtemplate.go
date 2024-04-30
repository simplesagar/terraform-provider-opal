// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type UpdateConfigurationTemplateResponse struct {
	// The configuration template just updated.
	ConfigurationTemplate *shared.ConfigurationTemplate
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateConfigurationTemplateResponse) GetConfigurationTemplate() *shared.ConfigurationTemplate {
	if o == nil {
		return nil
	}
	return o.ConfigurationTemplate
}

func (o *UpdateConfigurationTemplateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateConfigurationTemplateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConfigurationTemplateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}