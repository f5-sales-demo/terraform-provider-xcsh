---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_data_type."
xcsh_docs: {"aliases": [], "body_bytes": 948, "body_sha256": "sha256:d728babfd482227d04b87c43925a54cb67c1bc1525f610113664792669099ace", "canonical_id": "xcsh-docs:resources:data_type:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:data_type:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:292681f3b6c18ec2ffbe2be1a0e17beb9c57b44e1777befb42567a6c3a8f12ee", "source_path": "examples/resources/xcsh_data_type/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:data_type:example:resource", "parent_id": "xcsh-docs:resources:data_type:examples", "path": "docs/guides/resources--data_type--example--resource.md", "provider_name": "data_type", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_type/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_data_type.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md)
- [Examples](resources--data_type--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_data_type/resource.tf`; digest `sha256:292681f3b6c18ec2ffbe2be1a0e17beb9c57b44e1777befb42567a6c3a8f12ee`.

```terraform
# DataType Resource Example
# Manages data_type creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataType configuration
resource "xcsh_data_type" "example" {
  name      = "example-data-type"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--data_type--examples.md)
- [xcsh_data_type](../resources/data_type.md)
