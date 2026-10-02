---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_regional_edges."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1534, "body_sha256": "sha256:7efce3cc43ad32b2c51d3243660674b1045ae7651cc410bdc974c8b6299be963", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_regional_edges:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1744edb321716a35476a9a25cade01cfd9671a46b90a9dfb2f90ccff555099d6", "source_path": "examples/data-sources/xcsh_network_regional_edges/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_regional_edges:example:data-source", "parent_id": "xcsh-docs:data-sources:network_regional_edges:examples", "path": "documentation/data-sources/network_regional_edges/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_regional_edges", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3012131233211333-0113131211111320-3011322331211210-2200313333300101-3232010110232233-2021013200123312-1311122020120212-2110002123211023", "registry_path": "docs/guides/data-sources--network_regional_edges--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_regional_edges/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_regional_edges.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_regional_edges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_regional_edges/data-source.tf`; digest `sha256:1744edb321716a35476a9a25cade01cfd9671a46b90a9dfb2f90ccff555099d6`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

# Select the Regional Edge source networks that may initiate HTTPS connections
# to an origin. Omit regions to return all published regions.
data "xcsh_network_regional_edges" "origin_ingress" {
  regions = ["americas", "europe"]
}

output "https_origin_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_regional_edges.origin_ingress.cidr_blocks
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/examples/)
- [xcsh_network_regional_edges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/)
