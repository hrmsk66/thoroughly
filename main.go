package main

import (
	"context"
	"terraform-provider-thoroughly/thoroughly"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func main() {
	tfsdk.Serve(context.Background(), thoroughly.New, tfsdk.ServeOpts{
		Name: "throughly",
	})
}
