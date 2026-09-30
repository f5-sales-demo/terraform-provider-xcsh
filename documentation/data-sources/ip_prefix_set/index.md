---
page_title: "xcsh_ip_prefix_set"
subcategory: ""
description: "xcsh_ip_prefix_set for xcsh_ip_prefix_set."
xcsh_docs: {"aliases": [], "body_bytes": 1271, "body_sha256": "sha256:0a7465e7bed8eca6dc9fa7d5cdd8aba3d64b34f09c827a06cce3a3795ec203c5", "child_ids": ["xcsh-docs:data-sources:ip_prefix_set:reference", "xcsh-docs:data-sources:ip_prefix_set:examples"], "collection_id": "xcsh-docs:data-sources:ip_prefix_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ip_prefix_set:fundamentals", "parent_id": null, "path": "documentation/data-sources/ip_prefix_set/index.md", "provider_name": "ip_prefix_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_ip_prefix_set for xcsh_ip_prefix_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ip_prefix_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_ip_prefix_set

Breadcrumbs:

- xcsh_ip_prefix_set

Manages ip\_prefix\_set creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IPPrefixSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IPPrefixSet by name
data "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}

output "ip_prefix_set_id" {
  value = data.xcsh_ip_prefix_set.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ip_prefix_set/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ip_prefix_set/examples/)
