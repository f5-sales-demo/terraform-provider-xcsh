---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_infraprotect_mitigation_ips."
xcsh_docs: {"aliases": [], "body_bytes": 1203, "body_sha256": "sha256:fe533b0f5d4a82ece8c04561017aaaf023a325db733400d08011800de5b8f60e", "canonical_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0666bca16af2fe6a216b181ad7ffa5f07262937f90b186cb1b67917f3286333d", "source_path": "examples/data-sources/xcsh_infraprotect_mitigation_ips/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:example:data-source", "parent_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:examples", "path": "docs/guides/data-sources--infraprotect_mitigation_ips--example--data-source.md", "provider_name": "infraprotect_mitigation_ips", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/infraprotect_mitigation_ips/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_infraprotect_mitigation_ips.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_infraprotect_mitigation_ips](../data-sources/infraprotect_mitigation_ips.md)
- [Examples](data-sources--infraprotect_mitigation_ips--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_infraprotect_mitigation_ips/data-source.tf`; digest `sha256:0666bca16af2fe6a216b181ad7ffa5f07262937f90b186cb1b67917f3286333d`.

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

## Next pages

- [Examples](data-sources--infraprotect_mitigation_ips--examples.md)
- [xcsh_infraprotect_mitigation_ips](../data-sources/infraprotect_mitigation_ips.md)
