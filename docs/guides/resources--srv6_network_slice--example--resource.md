---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_srv6_network_slice."
xcsh_docs: {"aliases": [], "body_bytes": 1188, "body_sha256": "sha256:3f69d93142b535c14aba9da27e1808cc3a7bf32de1e0276cbddb2eede164444f", "canonical_id": "xcsh-docs:resources:srv6_network_slice:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:srv6_network_slice:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:28b76f9b2c2ac199134ab8ed152b0dff8fdb07e8a894ab2aa2291993080df8d0", "source_path": "examples/resources/xcsh_srv6_network_slice/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:srv6_network_slice:example:resource", "parent_id": "xcsh-docs:resources:srv6_network_slice:examples", "path": "docs/guides/resources--srv6_network_slice--example--resource.md", "provider_name": "srv6_network_slice", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/srv6_network_slice/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_srv6_network_slice.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["srv6_network_sliceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md)
- [Examples](resources--srv6_network_slice--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_srv6_network_slice/resource.tf`; digest `sha256:28b76f9b2c2ac199134ab8ed152b0dff8fdb07e8a894ab2aa2291993080df8d0`.

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

## Next pages

- [Examples](resources--srv6_network_slice--examples.md)
- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md)
