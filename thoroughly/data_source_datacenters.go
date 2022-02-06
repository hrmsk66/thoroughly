package thoroughly

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type dataSourceDatacentersType struct{}

func (d dataSourceDatacentersType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"datacenters": {
				Computed: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"code": {
						Type:     types.StringType,
						Computed: true,
					},
					"name": {
						Type:     types.StringType,
						Computed: true,
					},
					"group": {
						Type:     types.StringType,
						Computed: true,
					},
					"coordinates": {
						Required: true,
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"x": {
								Type:     types.Float64Type,
								Computed: true,
							},
							"y": {
								Type:     types.Float64Type,
								Computed: true,
							},
							"latitude": {
								Type:     types.Float64Type,
								Computed: true,
							},
							"longitude": {
								Type:     types.Float64Type,
								Computed: true,
							},
						}),
					},
					"shield": {
						Type:     types.StringType,
						Computed: true,
						Optional: true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (d dataSourceDatacentersType) NewDataSource(_ context.Context, p tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return dataSourceDatacenters{
		p: *(p.(*provider)),
	}, nil
}

type dataSourceDatacenters struct {
	p provider
}

func (d dataSourceDatacenters) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var state Datacenters

	datacenters, err := d.p.client.AllDatacenters()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading datacenters",
			err.Error(),
		)
		return
	}
	fmt.Fprintf(os.Stderr, "[DEBUG]-datacenters:%+v", datacenters)

	state.Datacenters = []Datacenter{}
	for _, datacenter := range datacenters {
		state.Datacenters = append(state.Datacenters, Datacenter{
			Code:  types.String{Value: datacenter.Code},
			Name:  types.String{Value: datacenter.Name},
			Group: types.String{Value: datacenter.Group},
			Coordinates: Coordinates{
				X:         types.Float64{Value: datacenter.Coordinates.X},
				Y:         types.Float64{Value: datacenter.Coordinates.Y},
				Latitude:  types.Float64{Value: datacenter.Coordinates.Latitude},
				Longitude: types.Float64{Value: datacenter.Coordinates.Longtitude}, // "Longtitude" is a typo in fastly/go-fastly. https://github.com/fastly/go-fastly/issues/274
			},
			Shield: types.String{Value: datacenter.Shield},
		})
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
