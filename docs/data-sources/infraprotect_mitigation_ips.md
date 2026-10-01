---
page_title: "xcsh_infraprotect_mitigation_ips"
subcategory: ""
description: "xcsh_infraprotect_mitigation_ips for xcsh_infraprotect_mitigation_ips."
xcsh_docs: {"aliases": [], "body_bytes": 1285, "body_sha256": "sha256:85a0c462a185d1aa50c79269e0e54776e4b419d4ecd269769b2f4f1f6e514cc9", "canonical_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:fundamentals", "child_ids": ["xcsh-docs:data-sources:infraprotect_mitigation_ips:reference", "xcsh-docs:data-sources:infraprotect_mitigation_ips:examples"], "collection_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:fundamentals", "parent_id": null, "path": "docs/data-sources/infraprotect_mitigation_ips.md", "provider_name": "infraprotect_mitigation_ips", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/infraprotect_mitigation_ips/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_infraprotect_mitigation_ips for xcsh_infraprotect_mitigation_ips.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_infraprotect_mitigation_ips

Breadcrumbs:

- xcsh_infraprotect_mitigation_ips

Resource retrieval operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# InfraprotectMitigationIps DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_infraprotect_mitigation_ips" "example" {
  mitigation_id = "example-value"
  namespace     = "example-value"
}

output "infraprotect_mitigation_ips_result" {
  value = data.xcsh_infraprotect_mitigation_ips.example
}
```

## Root configuration

Required root properties: `mitigation_id`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--infraprotect_mitigation_ips--reference.md)
- [Examples](../guides/data-sources--infraprotect_mitigation_ips--examples.md)
