---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_infraprotect_mitigation_ips."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1409, "body_sha256": "sha256:303911c507bf833902df7b37c046b9ec9be160b54522191073d9bd0bc6a3bf73", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0666bca16af2fe6a216b181ad7ffa5f07262937f90b186cb1b67917f3286333d", "source_path": "examples/data-sources/xcsh_infraprotect_mitigation_ips/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:example:data-source", "parent_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:examples", "path": "documentation/data-sources/infraprotect_mitigation_ips/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_mitigation_ips", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2323222212111121-1020330310021221-2100022220100331-2103120000311113-3220321301201120-0113320021300110-2202332100312021-0113102330213302", "registry_path": "docs/guides/data-sources--infraprotect_mitigation_ips--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/infraprotect_mitigation_ips/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_infraprotect_mitigation_ips.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_infraprotect_mitigation_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/examples/)
- [xcsh_infraprotect_mitigation_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/)
