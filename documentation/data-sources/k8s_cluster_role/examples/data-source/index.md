---
page_title: "Data source"
subcategory: "Container"
description: "Data source for xcsh_k8s_cluster_role."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1088, "body_sha256": "sha256:f463cd326b554d8a833de44c7c8316b15e965914bb93c8499cd8c416a1c0a6ab", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster_role:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3a2a8196bb01ed845c9acfb23f02e2761ef76a0ab6bfc6ad04f8767ce9170112", "source_path": "examples/data-sources/xcsh_k8s_cluster_role/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:k8s_cluster_role:example:data-source", "parent_id": "xcsh-docs:data-sources:k8s_cluster_role:examples", "path": "documentation/data-sources/k8s_cluster_role/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1212100012103223-0023010010321332-0031131100331110-2013103323033232-3211320211320212-1220111002313301-0232010321000002-0001231221323331", "registry_path": "docs/guides/data-sources--k8s_cluster_role--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster_role/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_k8s_cluster_role.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_k8s_cluster_role/data-source.tf`; digest `sha256:3a2a8196bb01ed845c9acfb23f02e2761ef76a0ab6bfc6ad04f8767ce9170112`.

```terraform
# K8SClusterRole Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SClusterRole by name
data "xcsh_k8s_cluster_role" "example" {
  name      = "example-k8s-cluster-role"
  namespace = "system"
}

output "k8s_cluster_role_id" {
  value = data.xcsh_k8s_cluster_role.example.id
}
```
