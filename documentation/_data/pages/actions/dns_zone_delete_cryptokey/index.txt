---
page_title: "xcsh_dns_zone_delete_cryptokey"
subcategory: ""
description: "Deletes a cryptographic key from a DNS zone."
xcsh_docs: {"aliases": ["dns zone delete cryptokey"], "body_bytes": 1307, "body_sha256": "sha256:15a16013dcc138679c7eb8504fe0dd3ea9c24819bf52ed414c503f1736ab1a81", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:dns_zone_delete_cryptokey:reference", "xcsh-docs:actions:dns_zone_delete_cryptokey:examples", "xcsh-docs:actions:dns_zone_delete_cryptokey:lifecycle"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:dns_zone_delete_cryptokey:collection", "completeness": "complete", "id": "xcsh-docs:actions:dns_zone_delete_cryptokey:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/dns_zone_delete_cryptokey/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_delete_cryptokey", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "actions", "registry_anchor": "canonical-3100330101232133-3311113231220332-0022020131232233-2130012232123321-2121220121030020-0323223001111032-0003233110133020-2000133100010203", "registry_path": "docs/actions/dns_zone_delete_cryptokey.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_delete_cryptokey/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Deletes a cryptographic key from a DNS zone.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_zone_delete_cryptokey

Breadcrumbs:

- xcsh_dns_zone_delete_cryptokey

Deletes a cryptographic key from a DNS zone.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSZoneDeleteCryptokey Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_dns_zone_delete_cryptokey" "example" {
  config {
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_delete_cryptokey/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_delete_cryptokey/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_delete_cryptokey/lifecycle/)
