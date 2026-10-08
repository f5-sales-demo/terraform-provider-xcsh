---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_data_group."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 979, "body_sha256": "sha256:7acd1555d79d1c166d8c50adbf1badbae0581c9798b0bc827b394d28b67400ca", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1bceaa79a1e1b840495f41f1bf5adac4d3dc173e5487643b627a4cc77fc7f690", "source_path": "examples/resources/xcsh_data_group/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:data_group:example:resource", "parent_id": "xcsh-docs:resources:data_group:examples", "path": "documentation/resources/data_group/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "data_group", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2203123231131002-1010333110013022-0002230320220313-2231031310211333-1211333111103021-2310133322223011-1021233301312012-0111201321002133", "registry_path": "docs/guides/resources--data_group--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_group/examples/resource/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Resource for xcsh_data_group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["data_groupCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_data_group/resource.tf`; digest `sha256:1bceaa79a1e1b840495f41f1bf5adac4d3dc173e5487643b627a4cc77fc7f690`.

```terraform
# DataGroup Resource Example
# Manages data group in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataGroup configuration
resource "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}
```
