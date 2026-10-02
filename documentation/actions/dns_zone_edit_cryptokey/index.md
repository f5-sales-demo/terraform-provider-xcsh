---
page_title: "xcsh_dns_zone_edit_cryptokey"
subcategory: ""
description: "Resource creation operation."
xcsh_docs: {"aliases": ["dns zone edit cryptokey"], "body_bytes": 1264, "body_sha256": "sha256:b93c049749d35f62f2efa5eda7991fb61e4e534b9105bd813f50fbc42ab5769f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:dns_zone_edit_cryptokey:reference", "xcsh-docs:actions:dns_zone_edit_cryptokey:examples", "xcsh-docs:actions:dns_zone_edit_cryptokey:lifecycle"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:collection", "completeness": "complete", "id": "xcsh-docs:actions:dns_zone_edit_cryptokey:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/dns_zone_edit_cryptokey/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_edit_cryptokey", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "registry_anchor": "canonical-1132220310333100-2000123203101123-2131221001021322-3330222102211322-3331333302020203-3030331000102011-2023032220122321-2312001020231030", "registry_path": "docs/actions/dns_zone_edit_cryptokey.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_edit_cryptokey/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource creation operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_zone_edit_cryptokey

Breadcrumbs:

- xcsh_dns_zone_edit_cryptokey

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_edit_cryptokey/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_edit_cryptokey/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_edit_cryptokey/lifecycle/)
