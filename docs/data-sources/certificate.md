---
page_title: "xcsh_certificate"
subcategory: "Security"
description: "xcsh_certificate for xcsh_certificate."
xcsh_docs: {"aliases": [], "body_bytes": 1274, "body_sha256": "sha256:a4c5943cb2a96ac7c0235f97fa9c814511349712d3783f9b1b4b6cd83a0f0378", "canonical_id": "xcsh-docs:data-sources:certificate:fundamentals", "child_ids": ["xcsh-docs:data-sources:certificate:reference", "xcsh-docs:data-sources:certificate:examples"], "collection_id": "xcsh-docs:data-sources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certificate:fundamentals", "parent_id": null, "path": "docs/data-sources/certificate.md", "provider_name": "certificate", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certificate/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_certificate for xcsh_certificate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["certificateCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_certificate

Breadcrumbs:

- xcsh_certificate

Manages a Certificate resource in F5 Distributed Cloud for certificate. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Certificate Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Certificate by name
data "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"
}

output "certificate_id" {
  value = data.xcsh_certificate.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--certificate--reference.md)
- [Examples](../guides/data-sources--certificate--examples.md)
