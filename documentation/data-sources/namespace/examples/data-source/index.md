---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_namespace."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1053, "body_sha256": "sha256:b3ba0584185bf54121119cd54b5080f90db84ee00d2238eba04318b6ef893a76", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:namespace:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a2418f75dbce2847a8b4c2579a5eb124cd8b027cba5de6019e63eeb934c35492", "source_path": "examples/data-sources/xcsh_namespace/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:namespace:example:data-source", "parent_id": "xcsh-docs:data-sources:namespace:examples", "path": "documentation/data-sources/namespace/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "namespace", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0111132001320022-0023302101302010-3320002033232100-2210100313200330-2201112122230113-3012122113121202-1333011311020123-2012030121120300", "registry_path": "docs/guides/data-sources--namespace--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/namespace/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_namespace.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["namespaceCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/namespace/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/namespace/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_namespace/data-source.tf`; digest `sha256:a2418f75dbce2847a8b4c2579a5eb124cd8b027cba5de6019e63eeb934c35492`.

```terraform
# Namespace Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Credentials are supplied externally.
provider "xcsh" {}

# Look up an existing Namespace by name
data "xcsh_namespace" "example" {
  name = "example-namespace"
}

output "namespace_id" {
  value = data.xcsh_namespace.example.id
}
```
