---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_segment."
xcsh_docs: {"aliases": [], "body_bytes": 1116, "body_sha256": "sha256:558827e189a7fb2bd62996d2043639d66864f8c8c435fd066ce56d894f675778", "child_ids": [], "collection_id": "xcsh-docs:data-sources:segment:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:2e9bfc7c69e50200df48c3eca01ce203db4e0516ed11b6bf987fdf485d490ac4", "source_path": "examples/data-sources/xcsh_segment/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:segment:example:data-source", "parent_id": "xcsh-docs:data-sources:segment:examples", "path": "documentation/data-sources/segment/examples/data-source/index.md", "provider_name": "segment", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_segment.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["segmentCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_segment/data-source.tf`; digest `sha256:2e9bfc7c69e50200df48c3eca01ce203db4e0516ed11b6bf987fdf485d490ac4`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment/examples/)
- [xcsh_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment/)
