---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_segment."
xcsh_docs: {"aliases": [], "body_bytes": 1087, "body_sha256": "sha256:952abc9d8545843b73c5592d021a31dfc1758d7acf243529d48663d62b9bb412", "child_ids": [], "collection_id": "xcsh-docs:resources:segment:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d784ed069273afd06c3680c3f92fabd8acd51be0aa4d921e6bfa8799ea0d6232", "source_path": "examples/resources/xcsh_segment/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:segment:example:resource", "parent_id": "xcsh-docs:resources:segment:examples", "path": "documentation/resources/segment/examples/resource/index.md", "provider_name": "segment", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/segment/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_segment.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["segmentCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_segment/resource.tf`; digest `sha256:d784ed069273afd06c3680c3f92fabd8acd51be0aa4d921e6bfa8799ea0d6232`.

```terraform
# Segment Resource Example
# Manages a Segment resource in F5 Distributed Cloud for segment.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Segment configuration
resource "xcsh_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/examples/)
- [xcsh_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/)
