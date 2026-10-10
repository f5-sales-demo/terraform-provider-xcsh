---
page_title: "xcsh_fast_acl_rule"
subcategory: ""
description: "Reads a Fast ACL rule with source IP and port matching and an associated action."
xcsh_docs: {"aliases": ["fast acl rule"], "body_bytes": 1348, "body_sha256": "sha256:dfcca02a4100c3effcdf88487567eed6e5953312b3983fc8a536b634d858f24a", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:fast_acl_rule:reference", "xcsh-docs:data-sources:fast_acl_rule:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl_rule:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/fast_acl_rule/index.md", "product": "distributed-cloud", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032", "registry_path": "docs/data-sources/fast_acl_rule.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl_rule/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reads a Fast ACL rule with source IP and port matching and an associated action.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_fast_acl_rule

Breadcrumbs:

- xcsh_fast_acl_rule

Reads a Fast ACL rule with source IP and port matching and an associated action.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FastACLRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FastACLRule by name
data "xcsh_fast_acl_rule" "example" {
  name      = "example-fast-acl-rule"
  namespace = "staging"
}

output "fast_acl_rule_id" {
  value = data.xcsh_fast_acl_rule.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/examples/)
