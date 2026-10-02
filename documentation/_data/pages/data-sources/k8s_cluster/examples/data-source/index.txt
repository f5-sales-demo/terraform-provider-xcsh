---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_k8s_cluster."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1265, "body_sha256": "sha256:8954ab134e3f083fc44fbb03f3816b3bc341f9af509b66f1909196631fa19787", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1ffe54fa3443473ef381429123673c75783a908cc4dae691369d96028c415887", "source_path": "examples/data-sources/xcsh_k8s_cluster/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:k8s_cluster:example:data-source", "parent_id": "xcsh-docs:data-sources:k8s_cluster:examples", "path": "documentation/data-sources/k8s_cluster/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3101112002210110-2023302322113310-3001112202310120-2122213213233222-2210332312102132-1030300002030131-3130113301123120-0322033210221000", "registry_path": "docs/guides/data-sources--k8s_cluster--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_k8s_cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_k8s_cluster/data-source.tf`; digest `sha256:1ffe54fa3443473ef381429123673c75783a908cc4dae691369d96028c415887`.

```terraform
# K8SCluster Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SCluster by name
data "xcsh_k8s_cluster" "example" {
  name      = "example-k8s-cluster"
  namespace = "system"
}

output "k8s_cluster_id" {
  value = data.xcsh_k8s_cluster.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/examples/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
