---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cloud_link."
xcsh_docs: {"aliases": [], "body_bytes": 923, "body_sha256": "sha256:61defb9ebca5d6a5a0d182a411d4d34b46e4b2a5594513e3263d966763457c4a", "canonical_id": "xcsh-docs:resources:cloud_link:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1be93a99c9a3175561f8ebf72f83a5762c7f54a0a05e770026c4c08e2199ba00", "source_path": "examples/resources/xcsh_cloud_link/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cloud_link:example:resource", "parent_id": "xcsh-docs:resources:cloud_link:examples", "path": "docs/guides/resources--cloud_link--example--resource.md", "provider_name": "cloud_link", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_cloud_link.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md)
- [Examples](resources--cloud_link--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_link/resource.tf`; digest `sha256:1be93a99c9a3175561f8ebf72f83a5762c7f54a0a05e770026c4c08e2199ba00`.

```terraform
# CloudLink Resource Example
# Manages new CloudLink with configured parameters in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudLink configuration
resource "xcsh_cloud_link" "example" {
  name      = "example-cloud-link"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--cloud_link--examples.md)
- [xcsh_cloud_link](../resources/cloud_link.md)
