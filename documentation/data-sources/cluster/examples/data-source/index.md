---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cluster."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1003, "body_sha256": "sha256:a61b26c14484f8d4e53dac274416d222ee7f32ea481d0fcadcea1bcdda5b9fee", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:657c300c6f1b904145d0c5ed8c458dbdc20cbe652caf3ef6b0a40bc6ebb5e4ee", "source_path": "examples/data-sources/xcsh_cluster/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cluster:example:data-source", "parent_id": "xcsh-docs:data-sources:cluster:examples", "path": "documentation/data-sources/cluster/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-3113213201302132-3020232133203000-1212202212233003-0033333300300310-1110111323132220-1001202322033223-0120321221121102-2221021311131021", "registry_path": "docs/guides/data-sources--cluster--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/examples/data-source/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Data source for xcsh_cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["clusterCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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
