terraform {
  required_providers {
    encode = {
      version = "0.1"
      source  = "registry.terraform.io/justenwalker/encode"
    }
  }
}

provider "encode" {}

data "encode_base36" "something" {
  value     = "hello"
  lowercase = true
}

output "encoded" {
  value = data.encode_base36.something.result
}
