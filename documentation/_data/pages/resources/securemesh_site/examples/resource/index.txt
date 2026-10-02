---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_securemesh_site."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1375, "body_sha256": "sha256:979019d8c7d5c04ad6573512a30a5d3a4d55300a52342ffc565ddeefc3a0ee1f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d12d5ac4d45f23102fc4cf022413c9bdb3fd5f59873115079a2e3a8fe4e2c28a", "source_path": "examples/resources/xcsh_securemesh_site/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:securemesh_site:example:resource", "parent_id": "xcsh-docs:resources:securemesh_site:examples", "path": "documentation/resources/securemesh_site/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0201022220011223-1321311300330030-2301020132113000-0030001120330303-1200200000212133-0302020330231221-3120023120220122-0022033132001000", "registry_path": "docs/guides/resources--securemesh_site--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_securemesh_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_securemesh_site/resource.tf`; digest `sha256:d12d5ac4d45f23102fc4cf022413c9bdb3fd5f59873115079a2e3a8fe4e2c28a`.

```terraform
# SecuremeshSite Resource Example
# Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with distributed security.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecuremeshSite configuration
resource "xcsh_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "staging"

  volterra_certified_hw = "example-value"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/examples/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
