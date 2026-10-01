---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_network_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1124, "body_sha256": "sha256:69e98d0dc75bc326eaaa0f52f90003dbcb2283ddc9c9120f9190b9f4a01fe1f0", "canonical_id": "xcsh-docs:data-sources:network_firewall:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5641721d230eef959a292217d55745f5f8be26fc7ce1feff6b83cbca8e44625c", "source_path": "examples/data-sources/xcsh_network_firewall/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_firewall:example:data-source", "parent_id": "xcsh-docs:data-sources:network_firewall:examples", "path": "docs/guides/data-sources--network_firewall--example--data-source.md", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_firewall/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md)
- [Examples](data-sources--network_firewall--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_firewall/data-source.tf`; digest `sha256:5641721d230eef959a292217d55745f5f8be26fc7ce1feff6b83cbca8e44625c`.

```terraform
# NetworkFirewall Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkFirewall by name
data "xcsh_network_firewall" "example" {
  name      = "example-network-firewall"
  namespace = "system"
}

output "network_firewall_id" {
  value = data.xcsh_network_firewall.example.id
}
```

## Next pages

- [Examples](data-sources--network_firewall--examples.md)
- [xcsh_network_firewall](../data-sources/network_firewall.md)
