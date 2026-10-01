---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_segment_connection."
xcsh_docs: {"aliases": [], "body_bytes": 1151, "body_sha256": "sha256:1d41ab451a52ad1ae73823daff33be13e41b97ba7a92d45862587a3c8b1e41ad", "canonical_id": "xcsh-docs:data-sources:segment_connection:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:segment_connection:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d15a7947bb7cd84484433c713ef71c626c3d8bc137b6e6bda822fbc845d28228", "source_path": "examples/data-sources/xcsh_segment_connection/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:segment_connection:example:data-source", "parent_id": "xcsh-docs:data-sources:segment_connection:examples", "path": "docs/guides/data-sources--segment_connection--example--data-source.md", "provider_name": "segment_connection", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment_connection/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_segment_connection.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_segment_connection](../data-sources/segment_connection.md)
- [Examples](data-sources--segment_connection--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_segment_connection/data-source.tf`; digest `sha256:d15a7947bb7cd84484433c713ef71c626c3d8bc137b6e6bda822fbc845d28228`.

```terraform
# SegmentConnection Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SegmentConnection by name
data "xcsh_segment_connection" "example" {
  name      = "example-segment-connection"
  namespace = "staging"
}

output "segment_connection_id" {
  value = data.xcsh_segment_connection.example.id
}
```

## Next pages

- [Examples](data-sources--segment_connection--examples.md)
- [xcsh_segment_connection](../data-sources/segment_connection.md)
