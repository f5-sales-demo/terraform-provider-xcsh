---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cluster."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1216, "body_sha256": "sha256:049fa665357ddfeefd152de777574c4156b41940dc285dea4ec55c05d74431f0", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:657c300c6f1b904145d0c5ed8c458dbdc20cbe652caf3ef6b0a40bc6ebb5e4ee", "source_path": "examples/data-sources/xcsh_cluster/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cluster:example:data-source", "parent_id": "xcsh-docs:data-sources:cluster:examples", "path": "documentation/data-sources/cluster/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3113213201302132-3020232133203000-1212202212233003-0033333300300310-1110111323132220-1001202322033223-0120321221121102-2221021311131021", "registry_path": "docs/guides/data-sources--cluster--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cluster/data-source.tf`; digest `sha256:657c300c6f1b904145d0c5ed8c458dbdc20cbe652caf3ef6b0a40bc6ebb5e4ee`.

```terraform
# Cluster Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cluster by name
data "xcsh_cluster" "example" {
  name      = "example-cluster"
  namespace = "staging"
}

output "cluster_id" {
  value = data.xcsh_cluster.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/examples/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
