---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_fast_acl_rule."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1307, "body_sha256": "sha256:226899a97cec9f90ae2b21a965c22dfb9d0da34c01294c9cdfb6d013873a6959", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c95e30b59cae8b267c4570beb6043ccc798329a2b5e981ef4b126965c5dceb06", "source_path": "examples/resources/xcsh_fast_acl_rule/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:fast_acl_rule:example:resource", "parent_id": "xcsh-docs:resources:fast_acl_rule:examples", "path": "documentation/resources/fast_acl_rule/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3301000010031022-3000300300221222-0132212232322321-1202033210222023-3303231031322210-2032320123312232-3223111103230211-0332323310033032", "registry_path": "docs/guides/resources--fast_acl_rule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl_rule/examples/resource/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Resource for xcsh_fast_acl_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_fast_acl_rule/resource.tf`; digest `sha256:c95e30b59cae8b267c4570beb6043ccc798329a2b5e981ef4b126965c5dceb06`.

```terraform
# FastACLRule Resource Example
# Manages new Fast ACL rule, has specification to match source IP, source port and action to apply in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FastACLRule configuration
resource "xcsh_fast_acl_rule" "example" {
  name      = "example-fast-acl-rule"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/examples/)
- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/)
