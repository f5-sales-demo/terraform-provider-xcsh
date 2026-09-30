---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_secondary_dns_zone_transfer."
xcsh_docs: {"aliases": [], "body_bytes": 1688, "body_sha256": "sha256:43632f49aa33d67db14457ff72b61953b6da23158da1c240a1671ac5d9cb849a", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:da49a1e4ce345336b0b68a09a44132872466f1849e8370d3d88d79c58a018362", "source_path": "examples/data-sources/xcsh_network_secondary_dns_zone_transfer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:example:data-source", "parent_id": "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:examples", "path": "documentation/data-sources/network_secondary_dns_zone_transfer/examples/data-source/index.md", "provider_name": "network_secondary_dns_zone_transfer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_secondary_dns_zone_transfer/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_secondary_dns_zone_transfer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_network_secondary_dns_zone_transfer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_secondary_dns_zone_transfer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_secondary_dns_zone_transfer/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_secondary_dns_zone_transfer/data-source.tf`; digest `sha256:da49a1e4ce345336b0b68a09a44132872466f1849e8370d3d88d79c58a018362`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_secondary_dns_zone_transfer" "authoritative_dns" {}

# The manifest combines transfer and notify sources, so both explicit DNS
# rules use the same published allowlist.
output "secondary_dns_rules" {
  value = [
    {
      direction = "ingress"
      protocol  = "tcp"
      port      = 53
      sources   = data.xcsh_network_secondary_dns_zone_transfer.authoritative_dns.cidr_blocks
    },
    {
      direction = "ingress"
      protocol  = "udp"
      port      = 53
      sources   = data.xcsh_network_secondary_dns_zone_transfer.authoritative_dns.cidr_blocks
    },
  ]
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_secondary_dns_zone_transfer/examples/)
- [xcsh_network_secondary_dns_zone_transfer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_secondary_dns_zone_transfer/)
