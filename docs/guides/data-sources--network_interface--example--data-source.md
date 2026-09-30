---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1039, "body_sha256": "sha256:626eb246f53d0e023b99fbb9c95085179aa6b180e1d94f907f17c504370deb8a", "canonical_id": "xcsh-docs:data-sources:network_interface:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:2b23ed57ede50484e6ffd1c4ce6fdf9ba70d1bbaddf2b7e8ceb68afb488d14e2", "source_path": "examples/data-sources/xcsh_network_interface/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_interface:example:data-source", "parent_id": "xcsh-docs:data-sources:network_interface:examples", "path": "docs/guides/data-sources--network_interface--example--data-source.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md)
- [Examples](data-sources--network_interface--examples.md)
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

## Next pages

- [Examples](data-sources--network_interface--examples.md)
- [xcsh_network_interface](../data-sources/network_interface.md)
