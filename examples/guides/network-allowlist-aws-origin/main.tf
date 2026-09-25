terraform {
  required_version = ">= 1.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 6.0.0"
    }
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

variable "origin_security_group_id" {
  type        = string
  description = "Security group attached to the HTTPS origin."
  default     = "sg-0123456789abcdef0"
}

data "xcsh_network_regional_edges" "origin_ingress" {
  regions = ["americas"]
}

resource "aws_vpc_security_group_ingress_rule" "f5xc_regional_edge_https" {
  for_each = toset(data.xcsh_network_regional_edges.origin_ingress.cidr_blocks)

  security_group_id = var.origin_security_group_id
  cidr_ipv4         = each.value
  from_port         = 443
  to_port           = 443
  ip_protocol       = "tcp"
  description       = "HTTPS ingress from an F5XC Regional Edge"
}
