---
page_title: "xcsh_srv6_network_slice"
subcategory: ""
description: "xcsh_srv6_network_slice for xcsh_srv6_network_slice."
xcsh_docs: {"aliases": [], "body_bytes": 1431, "body_sha256": "sha256:a0928bcba765f2eb33daaa7268d03d13753b0af131cebcfdb409e90429e114a0", "canonical_id": "xcsh-docs:resources:srv6_network_slice:fundamentals", "child_ids": ["xcsh-docs:resources:srv6_network_slice:reference", "xcsh-docs:resources:srv6_network_slice:examples", "xcsh-docs:resources:srv6_network_slice:import", "xcsh-docs:resources:srv6_network_slice:timeouts"], "collection_id": "xcsh-docs:resources:srv6_network_slice:collection", "completeness": "complete", "id": "xcsh-docs:resources:srv6_network_slice:fundamentals", "parent_id": null, "path": "docs/resources/srv6_network_slice.md", "provider_name": "srv6_network_slice", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/srv6_network_slice/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_srv6_network_slice for xcsh_srv6_network_slice.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["srv6_network_sliceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_srv6_network_slice

Breadcrumbs:

- xcsh_srv6_network_slice

Manages srv6\_network\_slice creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Srv6NetworkSlice Resource Example
# Manages srv6_network_slice creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Srv6NetworkSlice configuration
resource "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"

  sid_prefixes = ["example-value"]
}
```

## Root configuration

Required root properties: `name`, `sid_prefixes`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--srv6_network_slice--reference.md)
- [Examples](../guides/resources--srv6_network_slice--examples.md)
- [Import](../guides/resources--srv6_network_slice--import.md)
- [Timeouts](../guides/resources--srv6_network_slice--timeouts.md)
