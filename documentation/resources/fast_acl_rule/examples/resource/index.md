---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_fast_acl_rule."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1082, "body_sha256": "sha256:6a608498fb81b45df41062fb05c713e4254395e05762922e18a2955fbefb4bde", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c95e30b59cae8b267c4570beb6043ccc798329a2b5e981ef4b126965c5dceb06", "source_path": "examples/resources/xcsh_fast_acl_rule/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:fast_acl_rule:example:resource", "parent_id": "xcsh-docs:resources:fast_acl_rule:examples", "path": "documentation/resources/fast_acl_rule/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:500e603af7c92c2c213b5cbdb574fdea5ac4e4825db9ce520b047c223b245685", "provider_type": "resources", "registry_anchor": "canonical-3301000010031022-3000300300221222-0132212232322321-1202033210222023-3303231031322210-2032320123312232-3223111103230211-0332323310033032", "registry_path": "docs/guides/resources--fast_acl_rule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl_rule/examples/resource/index.txt", "spec_pin_digest": "sha256:cf249826c2873b41e4c4f1a2341c0bebf16f66cf6d95d5c5b67ae7259e61767b", "summary": "Resource for xcsh_fast_acl_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.6", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "06989705b3a435748dea253b9277f572cadb19c4"}}
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
