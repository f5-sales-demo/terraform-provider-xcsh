---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_certificate."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1268, "body_sha256": "sha256:fadfaa8ab73f0b43e81ffc40647068889b9a5da7b3567fb78d19cf4b21e73633", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:certificate:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ae1b4b8d1a231d986bf4b3d7767f79709ca2d85beaf03f7ea95f9a72d8a1d03e", "source_path": "examples/data-sources/xcsh_certificate/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:certificate:example:data-source", "parent_id": "xcsh-docs:data-sources:certificate:examples", "path": "documentation/data-sources/certificate/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "certificate", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0230322312131310-1100223021111133-2013003200312120-3001030011211123-2210201231002213-3101333023233123-2120323312113233-3103212213321302", "registry_path": "docs/guides/data-sources--certificate--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certificate/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["certificateCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_certificate/data-source.tf`; digest `sha256:ae1b4b8d1a231d986bf4b3d7767f79709ca2d85beaf03f7ea95f9a72d8a1d03e`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/examples/)
- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/)
