// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/conduitio/conduit-connector-sdk/cmd/paramgen

package discord

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func (SourceConfig) Parameters() map[string]sdk.Parameter {
	return map[string]sdk.Parameter{
		"body": {
			Default:     "",
			Description: "Http body to use in the request",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"headers": {
			Default:     "",
			Description: "Http headers to use in the request, comma separated list of : separated pairs",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"method": {
			Default:     "GET",
			Description: "Http method to use in the request",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationInclusion{List: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "CONNECT", "OPTIONS", "TRACE"}},
			},
		},
		"params": {
			Default:     "",
			Description: "parameters to use in the request, & separated list of = separated pairs",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"pollingPeriod": {
			Default:     "5m",
			Description: "how often the connector will get data from the url",
			Type:        sdk.ParameterTypeDuration,
			Validations: []sdk.Validation{},
		},
		"url": {
			Default:     "",
			Description: "Http url to use in the request",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
	}
}
