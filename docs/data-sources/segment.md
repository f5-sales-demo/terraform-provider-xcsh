---
page_title: "xcsh_segment"
subcategory: ""
description: "xcsh_segment for xcsh_segment."
xcsh_docs: {"aliases": [], "body_bytes": 1079, "body_sha256": "sha256:b4a6434f88919232a39873961c953ed466febd09f19d33cc994b8b23a5ffacc8", "canonical_id": "xcsh-docs:data-sources:segment:fundamentals", "child_ids": ["xcsh-docs:data-sources:segment:reference", "xcsh-docs:data-sources:segment:examples"], "collection_id": "xcsh-docs:data-sources:segment:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:segment:fundamentals", "parent_id": null, "path": "docs/data-sources/segment.md", "provider_name": "segment", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_segment for xcsh_segment.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["segmentCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_segment

Breadcrumbs:

- xcsh_segment

Manages a Segment resource in F5 Distributed Cloud for segment. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Segment Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Segment by name
data "xcsh_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}

output "segment_id" {
  value = data.xcsh_segment.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--segment--reference.md)
- [Examples](../guides/data-sources--segment--examples.md)
