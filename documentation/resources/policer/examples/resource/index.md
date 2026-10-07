---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_policer."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1053, "body_sha256": "sha256:0493dc27f09fe219c0c91312dd78e9a16b4a5674362664a3b88fd59107962a9c", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:771424522ef5cd3615902f335201bca762afca86a171354ca7b9f6787dcaf443", "source_path": "examples/resources/xcsh_policer/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:policer:example:resource", "parent_id": "xcsh-docs:resources:policer:examples", "path": "documentation/resources/policer/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "policer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1112112102112202-1312020331023111-1012320113033311-0102113031110021-0121131031221111-0133031133130310-2102031120132310-0323020111332102", "registry_path": "docs/guides/resources--policer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policer/examples/resource/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Resource for xcsh_policer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["policerCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policer/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_policer/resource.tf`; digest `sha256:771424522ef5cd3615902f335201bca762afca86a171354ca7b9f6787dcaf443`.

```terraform
# Policer Resource Example
# Manages new policer with traffic rate limits in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Policer configuration
resource "xcsh_policer" "example" {
  name      = "example-policer"
  namespace = "staging"

  burst_size                 = 1
  committed_information_rate = 1
}
```
