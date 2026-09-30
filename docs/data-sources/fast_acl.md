---
page_title: "xcsh_fast_acl"
subcategory: ""
description: "xcsh_fast_acl for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 1182, "body_sha256": "sha256:61430ff28267e153e0f2b0e79cf2a9d033553b64b2122a3b96d9f52a3d45e0f4", "canonical_id": "xcsh-docs:data-sources:fast_acl:fundamentals", "child_ids": ["xcsh-docs:data-sources:fast_acl:reference", "xcsh-docs:data-sources:fast_acl:examples"], "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:fundamentals", "parent_id": null, "path": "docs/data-sources/fast_acl.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_fast_acl for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_fast_acl

Breadcrumbs:

- xcsh_fast_acl

Manages object, object contains rules to protect site from denial of service It has
destination\{destination IP, destination port) and references to in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FastACL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FastACL by name
data "xcsh_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"
}

output "fast_acl_id" {
  value = data.xcsh_fast_acl.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--fast_acl--reference.md)
- [Examples](../guides/data-sources--fast_acl--examples.md)
