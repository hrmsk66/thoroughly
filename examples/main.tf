terraform {
  required_providers {
      thoroughly = {
          version = "0.1"
          source = "fastly/edu/thoroughly"
      }
  }
}

provider "thoroughly" {}

data "thoroughly_datacenters" "dc" {}

locals {
  shield_pops = { for d in data.thoroughly_datacenters.dc.datacenters : d.code => d.shield if d.shield != "" }
}

output "datacenters" {
    value = local.shield_pops
}

