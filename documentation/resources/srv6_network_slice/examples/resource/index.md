---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_srv6_network_slice."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1154, "body_sha256": "sha256:d1cd7dc11517c4849420f5847230f1064d0b2708494a6dc4a2c3fed982f8e5bc", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:srv6_network_slice:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:28b76f9b2c2ac199134ab8ed152b0dff8fdb07e8a894ab2aa2291993080df8d0", "source_path": "examples/resources/xcsh_srv6_network_slice/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:srv6_network_slice:example:resource", "parent_id": "xcsh-docs:resources:srv6_network_slice:examples", "path": "documentation/resources/srv6_network_slice/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "srv6_network_slice", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1331111120302201-1301121320111201-3320103121320102-0111112211101230-0010013013200103-1200330021230010-3021100201320131-2110302311023232", "registry_path": "docs/guides/resources--srv6_network_slice--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/srv6_network_slice/examples/resource/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Resource for xcsh_srv6_network_slice.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["srv6_network_sliceCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_srv6_network_slice](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/srv6_network_slice/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/srv6_network_slice/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_srv6_network_slice/resource.tf`; digest `sha256:28b76f9b2c2ac199134ab8ed152b0dff8fdb07e8a894ab2aa2291993080df8d0`.

```terraform
# Srv6NetworkSlice Resource Example
# Manages srv6_network_slice creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Srv6NetworkSlice configuration
resource "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"

  sid_prefixes = ["example-value"]
}
```
