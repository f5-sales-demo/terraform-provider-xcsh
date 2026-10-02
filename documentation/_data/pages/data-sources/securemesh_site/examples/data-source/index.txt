---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_securemesh_site."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1318, "body_sha256": "sha256:3969d37204eb582fcb6113a674fe1ddaed1c16e9dc1eab11a7d81ffe330913a5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:476e0c14fbf185d65694b51a11030a2462b91f28a9d067a59c797b6328e138a4", "source_path": "examples/data-sources/xcsh_securemesh_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:securemesh_site:example:data-source", "parent_id": "xcsh-docs:data-sources:securemesh_site:examples", "path": "documentation/data-sources/securemesh_site/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3233313300011331-2000133123211211-3121330230231301-2021321332011122-1033032303300123-3332221020032112-2010302031231132-3202113002111010", "registry_path": "docs/guides/data-sources--securemesh_site--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_securemesh_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_securemesh_site/data-source.tf`; digest `sha256:476e0c14fbf185d65694b51a11030a2462b91f28a9d067a59c797b6328e138a4`.

```terraform
# SecuremeshSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecuremeshSite by name
data "xcsh_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "staging"
}

output "securemesh_site_id" {
  value = data.xcsh_securemesh_site.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/examples/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
