---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_dc_cluster_group."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1052, "body_sha256": "sha256:c8b4dfbf62d33d268360880ee31ffdef3d238095c50d0c1389484f4b0ba68509", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dc_cluster_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8a0876385980d3e39be1eb4ce88a5f57c0c3917ba68858a952fc5d4f99aaef53", "source_path": "examples/resources/xcsh_dc_cluster_group/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dc_cluster_group:example:resource", "parent_id": "xcsh-docs:resources:dc_cluster_group:examples", "path": "documentation/resources/dc_cluster_group/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0223321230310323-2220020220301011-0123200033210203-2212212210021211-1120230132131231-0101003332213021-3112103031101223-1031333233113323", "registry_path": "docs/guides/resources--dc_cluster_group--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dc_cluster_group/examples/resource/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Resource for xcsh_dc_cluster_group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dc_cluster_group/resource.tf`; digest `sha256:8a0876385980d3e39be1eb4ce88a5f57c0c3917ba68858a952fc5d4f99aaef53`.

```terraform
# DcClusterGroup Resource Example
# Manages DC Cluster group in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DcClusterGroup configuration
resource "xcsh_dc_cluster_group" "example" {
  name      = "example-dc-cluster-group"
  namespace = "system"
}
```
