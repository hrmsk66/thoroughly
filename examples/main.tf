terraform {
  required_providers {
      thoroughly = {
          version = "0.1"
          source = "fastly/edu/thoroughly"
      }
  }
}

provider "thoroughly" {}

data "thoroughly_datacenters" "dc" {

}

output "datacenters" {
    value = data.thoroughly_datacenters.dc
}

