---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_dns_zone_add_cryptokey."
xcsh_docs: {"aliases": ["action"], "body_bytes": 915, "body_sha256": "sha256:bb73638377cb829c85cbff2391447e53ac6ef766f3b345e5ad038b384e4082af", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:dns_zone_add_cryptokey:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:41e2e4052bac9b33dbdf74fdf0a795bf2458f8e723d6c49a082ace67a79ec803", "source_path": "examples/actions/xcsh_dns_zone_add_cryptokey/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:dns_zone_add_cryptokey:example:action", "parent_id": "xcsh-docs:actions:dns_zone_add_cryptokey:examples", "path": "documentation/actions/dns_zone_add_cryptokey/examples/action/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_add_cryptokey", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "actions", "registry_anchor": "canonical-1321110312233313-2202012033000310-1120121203112113-2212220301023123-0021320210303212-0123121213322322-2033231303210313-2300312122230022", "registry_path": "docs/guides/actions--dns_zone_add_cryptokey--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_add_cryptokey/examples/action/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Action for xcsh_dns_zone_add_cryptokey.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_dns_zone_add_cryptokey](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_add_cryptokey/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_add_cryptokey/examples/)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_dns_zone_add_cryptokey/action.tf`; digest `sha256:41e2e4052bac9b33dbdf74fdf0a795bf2458f8e723d6c49a082ace67a79ec803`.

```terraform
# DNSZoneAddCryptokey Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_dns_zone_add_cryptokey" "example" {
  config {
  }
}
```
