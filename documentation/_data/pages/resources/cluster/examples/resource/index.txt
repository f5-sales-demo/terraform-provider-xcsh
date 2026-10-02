---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cluster."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1243, "body_sha256": "sha256:404a53c4714302c76dea96edc280bcd44966c3fd4cdf39cd53a01614f39f4375", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5826c5405f2445e4c35c5efebde0e90198d315b39c9333da9092af648210a59e", "source_path": "examples/resources/xcsh_cluster/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cluster:example:resource", "parent_id": "xcsh-docs:resources:cluster:examples", "path": "documentation/resources/cluster/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0331332022232123-1322210312322323-3011020222012033-3110310002112202-1232010201201122-0302122102013332-0301020330012110-0312002302322022", "registry_path": "docs/guides/resources--cluster--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["clusterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cluster/resource.tf`; digest `sha256:5826c5405f2445e4c35c5efebde0e90198d315b39c9333da9092af648210a59e`.

```terraform
# Cluster Resource Example
# Manages cluster will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Cluster configuration
resource "xcsh_cluster" "example" {
  name      = "example-cluster"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/examples/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
