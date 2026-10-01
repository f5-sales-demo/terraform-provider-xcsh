---
page_title: "xcsh_dns_zone_add_cryptokey"
subcategory: ""
description: "xcsh_dns_zone_add_cryptokey for xcsh_dns_zone_add_cryptokey."
xcsh_docs: {"aliases": [], "body_bytes": 1257, "body_sha256": "sha256:8e7c7f84411197a18d746211dec8f5300fd9cdcad2b5e3757f4a69c41a264116", "child_ids": ["xcsh-docs:actions:dns_zone_add_cryptokey:reference", "xcsh-docs:actions:dns_zone_add_cryptokey:examples", "xcsh-docs:actions:dns_zone_add_cryptokey:lifecycle"], "collection_id": "xcsh-docs:actions:dns_zone_add_cryptokey:collection", "completeness": "complete", "id": "xcsh-docs:actions:dns_zone_add_cryptokey:fundamentals", "parent_id": null, "path": "documentation/actions/dns_zone_add_cryptokey/index.md", "provider_name": "dns_zone_add_cryptokey", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_add_cryptokey/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_dns_zone_add_cryptokey for xcsh_dns_zone_add_cryptokey.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
