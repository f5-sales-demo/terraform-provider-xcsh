---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_dns_zone_edit_cryptokey."
xcsh_docs: {"aliases": ["action"], "body_bytes": 921, "body_sha256": "sha256:c93b1dfd929da00c3869d64bee0f598782b7be2ab8dee4ad5975bfc6b4eb3e58", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:377745c7c273966a8712f1843e6db7821e8429814e456ec61ea024bb03bd4506", "source_path": "examples/actions/xcsh_dns_zone_edit_cryptokey/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:dns_zone_edit_cryptokey:example:action", "parent_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:examples", "path": "documentation/actions/dns_zone_edit_cryptokey/examples/action/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_edit_cryptokey", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "actions", "registry_anchor": "canonical-2131110033132201-0211303211102333-2013231203000120-2003230333001121-3001320202021223-1310110313112110-0130123330323322-2112303110223021", "registry_path": "docs/guides/actions--dns_zone_edit_cryptokey--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_edit_cryptokey/examples/action/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Action for xcsh_dns_zone_edit_cryptokey.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_dns_zone_edit_cryptokey](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_edit_cryptokey/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_edit_cryptokey/examples/)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_dns_zone_edit_cryptokey/action.tf`; digest `sha256:377745c7c273966a8712f1843e6db7821e8429814e456ec61ea024bb03bd4506`.

```terraform
# DNSZoneEditCryptokey Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_dns_zone_edit_cryptokey" "example" {
  config {
  }
}
```
