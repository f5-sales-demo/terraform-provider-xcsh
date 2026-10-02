---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_subnet."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1203, "body_sha256": "sha256:e491b23659078c5cc9e6e42ef8414b0affa8e786997efeab000e51f80f1dced5", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:subnet:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1dd541aa45894ccb54e6cdac49fe4ca6130605728a4cc127960dcf649a0838dd", "source_path": "examples/data-sources/xcsh_subnet/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:subnet:example:data-source", "parent_id": "xcsh-docs:data-sources:subnet:examples", "path": "documentation/data-sources/subnet/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2120213323020012-3222311111220111-3012100031230012-0223130202020212-0030313111111120-0130102111112320-2021103230330123-2213101222200122", "registry_path": "docs/guides/data-sources--subnet--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/subnet/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["subnetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_subnet/data-source.tf`; digest `sha256:1dd541aa45894ccb54e6cdac49fe4ca6130605728a4cc127960dcf649a0838dd`.

```terraform
# Subnet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Subnet by name
data "xcsh_subnet" "example" {
  name      = "example-subnet"
  namespace = "staging"
}

output "subnet_id" {
  value = data.xcsh_subnet.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/examples/)
- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/)
