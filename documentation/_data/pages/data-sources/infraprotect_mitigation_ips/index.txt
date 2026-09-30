---
page_title: "xcsh_infraprotect_mitigation_ips"
subcategory: ""
description: "xcsh_infraprotect_mitigation_ips for xcsh_infraprotect_mitigation_ips."
xcsh_docs: {"aliases": [], "body_bytes": 1271, "body_sha256": "sha256:c4f813e9be5b293781e17b7cb4eb5994a702835a1d9780dead7a8d03583b00a3", "child_ids": ["xcsh-docs:data-sources:infraprotect_mitigation_ips:reference", "xcsh-docs:data-sources:infraprotect_mitigation_ips:examples"], "collection_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:fundamentals", "parent_id": null, "path": "documentation/data-sources/infraprotect_mitigation_ips/index.md", "provider_name": "infraprotect_mitigation_ips", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/infraprotect_mitigation_ips/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_infraprotect_mitigation_ips for xcsh_infraprotect_mitigation_ips.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/examples/)
