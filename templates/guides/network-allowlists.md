# Release-pinned network allowlists

The network allowlist data sources read values compiled into this provider from
the info.x-f5xc-network-allowlist extension in the pinned API release. They do
not fetch the [F5 network reference](https://docs.cloud.f5.com/docs-v2/platform/reference/network-cloud-ref)
or call an F5 API during planning.

The manifest contains addresses and domains, not complete firewall rules.
Choose direction, protocol, and port for the configured service. The examples
make those choices explicit. In particular, the source does not distinguish CDN
geography or Secondary DNS transfer addresses from notify addresses.

## AWS HTTPS origin ingress

Use each normalized Regional Edge CIDR as a separate ingress-rule source:

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

The [complete AWS example](../../examples/guides/network-allowlist-aws-origin/main.tf)
is validated against a mocked AWS provider.

See the canonical data-source examples for explicit DNS TCP/UDP 53 ingress,
TLS syslog TCP 6514 egress, health-check TCP 443 ingress, Bot Defense TCP 443
proxy egress, Data Intelligence TCP 443 egress, Customer Edge DNS/NTP egress,
and Secure Mesh v2 TCP 443 egress configurations.
