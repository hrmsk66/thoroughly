package thoroughly

import (
	"context"
	"os"

	"github.com/fastly/go-fastly/v6/fastly"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var stdrr = os.Stderr

func New() tfsdk.Provider {
	return &provider{}
}

type provider struct {
	configured bool
	client     *fastly.Client
}

func (p *provider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"apikey": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
		},
	}, nil
}

type providerData struct {
	Apikey types.String `tfsdk:"apikey"`
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apikey string
	if config.Apikey.Unknown {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Cannot use unknown value as apikey",
		)
		return
	}
	if config.Apikey.Null {
		apikey = os.Getenv("FASTLY_API_KEY")
	} else {
		apikey = config.Apikey.Value
	}
	if apikey == "" {
		resp.Diagnostics.AddError(
			"Unable to find apikey",
			"apikey cannot be an empty string",
		)
		return
	}

	c, err := fastly.NewClient(apikey)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create hashicups client:\n\n"+err.Error(),
		)
		return
	}

	p.client = c
	p.configured = true
}

func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{}, nil
}

func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"thoroughly_datacenters": dataSourceDatacentersType{},
	}, nil
}
