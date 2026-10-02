---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_fast_acl_rule."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1307, "body_sha256": "sha256:226899a97cec9f90ae2b21a965c22dfb9d0da34c01294c9cdfb6d013873a6959", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c95e30b59cae8b267c4570beb6043ccc798329a2b5e981ef4b126965c5dceb06", "source_path": "examples/resources/xcsh_fast_acl_rule/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:fast_acl_rule:example:resource", "parent_id": "xcsh-docs:resources:fast_acl_rule:examples", "path": "documentation/resources/fast_acl_rule/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3301000010031022-3000300300221222-0132212232322321-1202033210222023-3303231031322210-2032320123312232-3223111103230211-0332323310033032", "registry_path": "docs/guides/resources--fast_acl_rule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl_rule/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_fast_acl_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
