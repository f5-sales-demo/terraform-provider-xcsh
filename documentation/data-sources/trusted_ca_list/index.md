---
page_title: "xcsh_trusted_ca_list"
subcategory: ""
description: "xcsh_trusted_ca_list for xcsh_trusted_ca_list."
xcsh_docs: {"aliases": [], "body_bytes": 1384, "body_sha256": "sha256:3ea1e842eb15e032295d5077aaf8ca351a37d6718df179dd4cf091817f2ee9fc", "child_ids": ["xcsh-docs:data-sources:trusted_ca_list:reference", "xcsh-docs:data-sources:trusted_ca_list:examples"], "collection_id": "xcsh-docs:data-sources:trusted_ca_list:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:trusted_ca_list:fundamentals", "parent_id": null, "path": "documentation/data-sources/trusted_ca_list/index.md", "provider_name": "trusted_ca_list", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/trusted_ca_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_trusted_ca_list for xcsh_trusted_ca_list.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["trusted_ca_listCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_trusted_ca_list

Breadcrumbs:

- xcsh_trusted_ca_list

Manages a Trusted CA List resource in F5 Distributed Cloud for trusted certificate authority list
management.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TrustedCAList Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TrustedCAList by name
data "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}

output "trusted_ca_list_id" {
  value = data.xcsh_trusted_ca_list.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/trusted_ca_list/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/trusted_ca_list/examples/)
