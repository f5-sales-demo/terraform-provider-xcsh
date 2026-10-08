---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_discovery."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1068, "body_sha256": "sha256:a93edad5dce34d907fd8df4a69f0b0ab9d749481504ffd48d1a7185ef129dab4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:639a573707f4cc4152b42dbe4ad48b3edaa1017b066d38b4ee06ee6d80b4f03c", "source_path": "examples/resources/xcsh_discovery/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:discovery:example:resource", "parent_id": "xcsh-docs:resources:discovery:examples", "path": "documentation/resources/discovery/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3231332330332133-3012022130101021-0231011222223303-3001221030232022-3032130300020103-1220010012220120-2033233200021330-3222223122312033", "registry_path": "docs/guides/resources--discovery--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/examples/resource/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Resource for xcsh_discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["discoveryCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_discovery/resource.tf`; digest `sha256:639a573707f4cc4152b42dbe4ad48b3edaa1017b066d38b4ee06ee6d80b4f03c`.

```terraform
# Discovery Resource Example
# Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Discovery configuration
resource "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}
```
