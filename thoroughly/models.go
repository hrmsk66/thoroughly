package thoroughly

import "github.com/hashicorp/terraform-plugin-framework/types"

type Datacenters struct {
	Datacenters []Datacenter `tfsdk:"datacenters"`
}

type Datacenter struct {
	Code        types.String `tfsdk:"code"`
	Name        types.String `tfsdk:"name"`
	Group       types.String `tfsdk:"group"`
	Coordinates Coordinates  `tfsdk:"coordinates"`
	Shield      types.String `tfsdk:"shield"`
}

type Coordinates struct {
	X         types.Float64 `tfsdk:"x"`
	Y         types.Float64 `tfsdk:"y"`
	Latitude  types.Float64 `tfsdk:"latitude"`
	Longitude types.Float64 `tfsdk:"longitude"`
}
