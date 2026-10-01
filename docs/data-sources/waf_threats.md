---
page_title: "xcsh_waf_threats"
subcategory: ""
description: "xcsh_waf_threats for xcsh_waf_threats."
xcsh_docs: {"aliases": [], "body_bytes": 1065, "body_sha256": "sha256:15e5b25d19f0e7b54e2c4dc4c0ac70d31935ee285ce2605bcf6358dba5873af7", "canonical_id": "xcsh-docs:data-sources:waf_threats:fundamentals", "child_ids": ["xcsh-docs:data-sources:waf_threats:reference", "xcsh-docs:data-sources:waf_threats:examples"], "collection_id": "xcsh-docs:data-sources:waf_threats:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_threats:fundamentals", "parent_id": null, "path": "docs/data-sources/waf_threats.md", "provider_name": "waf_threats", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threats/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_waf_threats for xcsh_waf_threats.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_waf_threats

Breadcrumbs:

- xcsh_waf_threats

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFThreats DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_threats" "example" {
}

output "waf_threats_result" {
  value = data.xcsh_waf_threats.example
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--waf_threats--reference.md)
- [Examples](../guides/data-sources--waf_threats--examples.md)
