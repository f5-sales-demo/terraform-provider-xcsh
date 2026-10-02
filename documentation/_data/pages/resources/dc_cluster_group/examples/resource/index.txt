---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_dc_cluster_group."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1286, "body_sha256": "sha256:0e4c3037d1dd7b21815745a234af5ccb6ba7d84e631938b0ba01e410cd20c4cd", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dc_cluster_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8a0876385980d3e39be1eb4ce88a5f57c0c3917ba68858a952fc5d4f99aaef53", "source_path": "examples/resources/xcsh_dc_cluster_group/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dc_cluster_group:example:resource", "parent_id": "xcsh-docs:resources:dc_cluster_group:examples", "path": "documentation/resources/dc_cluster_group/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0223321230310323-2220020220301011-0123200033210203-2212212210021211-1120230132131231-0101003332213021-3112103031101223-1031333233113323", "registry_path": "docs/guides/resources--dc_cluster_group--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dc_cluster_group/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_dc_cluster_group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/examples/)
- [xcsh_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/)
