---
page_title: "network-allowlists"
subcategory: ""
description: "Maintained network-allowlists guide."
xcsh_docs: {"aliases": ["network-allowlists"], "body_bytes": 1741, "body_sha256": "sha256:07af41b5d4db74629947a8ec7504bc57daf5f6c88d24dbeaa53b6c6abb27f4d7", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved"}, "collection_id": "xcsh-docs:guides:network-allowlists:collection", "completeness": "complete", "id": "xcsh-docs:guides:network-allowlists:overview", "parent_id": "xcsh-docs:provider:xcsh:navigation", "path": "documentation/guides/network-allowlists/index.md", "product": "distributed-cloud", "provider_name": "network-allowlists", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "guides", "registry_path": "docs/guides/network-allowlists.md", "relationships": [], "retrieval_version": 1, "role": "overview", "schema_path": [], "schema_version": 1, "sections": [], "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Maintained network-allowlists guide.", "tasks": ["configuration"]}
---

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

The [complete AWS example](https://github.com/f5-sales-demo/terraform-provider-xcsh/blob/main/examples/guides/network-allowlist-aws-origin/main.tf)
is validated against a mocked AWS provider.

See the canonical data-source examples for explicit DNS TCP/UDP 53 ingress,
TLS syslog TCP 6514 egress, health-check TCP 443 ingress, Bot Defense TCP 443
proxy egress, Data Intelligence TCP 443 egress, Customer Edge DNS/NTP egress,
and Secure Mesh v2 TCP 443 egress configurations.
