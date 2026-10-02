---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_trusted_ca_list."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1316, "body_sha256": "sha256:0059069f243ea5e00b099045f76142dd9d66a26d390aef34902eba29be9f8ac1", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:trusted_ca_list:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:444d0b10e4a27614d6854b1c286e2b90a9468950261cdb881f0b8b693ceeb1e7", "source_path": "examples/data-sources/xcsh_trusted_ca_list/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:trusted_ca_list:example:data-source", "parent_id": "xcsh-docs:data-sources:trusted_ca_list:examples", "path": "documentation/data-sources/trusted_ca_list/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "trusted_ca_list", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3212322202330102-2120312121312322-1122023321111303-2030002312232112-1123333303312102-3010220332113302-1000302221200323-1130000003121232", "registry_path": "docs/guides/data-sources--trusted_ca_list--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/trusted_ca_list/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_trusted_ca_list.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["trusted_ca_listCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/trusted_ca_list/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/trusted_ca_list/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_trusted_ca_list/data-source.tf`; digest `sha256:444d0b10e4a27614d6854b1c286e2b90a9468950261cdb881f0b8b693ceeb1e7`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/trusted_ca_list/examples/)
- [xcsh_trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/trusted_ca_list/)
