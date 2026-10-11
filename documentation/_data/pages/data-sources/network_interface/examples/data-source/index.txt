---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_interface."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1101, "body_sha256": "sha256:2c531a9a860f9474e6deadbe07d54e68a311421377b6c687c7e3f3f44921d9c3", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:2b23ed57ede50484e6ffd1c4ce6fdf9ba70d1bbaddf2b7e8ceb68afb488d14e2", "source_path": "examples/data-sources/xcsh_network_interface/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_interface:example:data-source", "parent_id": "xcsh-docs:data-sources:network_interface:examples", "path": "documentation/data-sources/network_interface/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-1030333002233230-1030302301300310-0330313123230111-0013312003023032-0230013211000220-3022102220031021-2013001023323221-0200202201121023", "registry_path": "docs/guides/data-sources--network_interface--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_network_interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_interface/data-source.tf`; digest `sha256:2b23ed57ede50484e6ffd1c4ce6fdf9ba70d1bbaddf2b7e8ceb68afb488d14e2`.

```terraform
# NetworkInterface Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkInterface by name
data "xcsh_network_interface" "example" {
  name      = "example-network-interface"
  namespace = "staging"
}

output "network_interface_id" {
  value = data.xcsh_network_interface.example.id
}
```
