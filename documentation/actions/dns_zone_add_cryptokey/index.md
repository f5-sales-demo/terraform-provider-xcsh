---
page_title: "xcsh_dns_zone_add_cryptokey"
subcategory: ""
description: "Resource creation operation."
xcsh_docs: {"aliases": ["dns zone add cryptokey"], "body_bytes": 1257, "body_sha256": "sha256:8e7c7f84411197a18d746211dec8f5300fd9cdcad2b5e3757f4a69c41a264116", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:dns_zone_add_cryptokey:reference", "xcsh-docs:actions:dns_zone_add_cryptokey:examples", "xcsh-docs:actions:dns_zone_add_cryptokey:lifecycle"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:dns_zone_add_cryptokey:collection", "completeness": "complete", "id": "xcsh-docs:actions:dns_zone_add_cryptokey:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/dns_zone_add_cryptokey/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_add_cryptokey", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "actions", "registry_anchor": "canonical-2011333313231130-3200122021220331-0010211322333131-1302203232313332-2231111130321232-3311022020312311-2322131130200322-0032011221212010", "registry_path": "docs/actions/dns_zone_add_cryptokey.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_add_cryptokey/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource creation operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_zone_add_cryptokey

Breadcrumbs:

- xcsh_dns_zone_add_cryptokey

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_add_cryptokey/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_add_cryptokey/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_add_cryptokey/lifecycle/)
