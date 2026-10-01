---
page_title: "xcsh_waf_attack_signatures"
subcategory: ""
description: "xcsh_waf_attack_signatures for xcsh_waf_attack_signatures."
xcsh_docs: {"aliases": [], "body_bytes": 1145, "body_sha256": "sha256:4dc0ee6621c76ef8441d8ea860d7b9ab662ac910d531eeeb18e0d819d69a0dac", "canonical_id": "xcsh-docs:data-sources:waf_attack_signatures:fundamentals", "child_ids": ["xcsh-docs:data-sources:waf_attack_signatures:reference", "xcsh-docs:data-sources:waf_attack_signatures:examples"], "collection_id": "xcsh-docs:data-sources:waf_attack_signatures:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_attack_signatures:fundamentals", "parent_id": null, "path": "docs/data-sources/waf_attack_signatures.md", "provider_name": "waf_attack_signatures", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_attack_signatures/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_waf_attack_signatures for xcsh_waf_attack_signatures.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_waf_attack_signatures

Breadcrumbs:

- xcsh_waf_attack_signatures

Resource retrieval operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFAttackSignatures DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_attack_signatures" "example" {
}

output "waf_attack_signatures_result" {
  value = data.xcsh_waf_attack_signatures.example
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--waf_attack_signatures--reference.md)
- [Examples](../guides/data-sources--waf_attack_signatures--examples.md)
