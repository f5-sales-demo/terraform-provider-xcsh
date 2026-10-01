---
page_title: "xcsh_dns_zone_edit_cryptokey"
subcategory: ""
description: "xcsh_dns_zone_edit_cryptokey for xcsh_dns_zone_edit_cryptokey."
xcsh_docs: {"aliases": [], "body_bytes": 1137, "body_sha256": "sha256:842fc8f0c053f3cc4d9ae0bc65ffcf48645f3c2e6790c2dae187d7050d3ee7c9", "canonical_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:fundamentals", "child_ids": ["xcsh-docs:actions:dns_zone_edit_cryptokey:reference", "xcsh-docs:actions:dns_zone_edit_cryptokey:examples", "xcsh-docs:actions:dns_zone_edit_cryptokey:lifecycle"], "collection_id": "xcsh-docs:actions:dns_zone_edit_cryptokey:collection", "completeness": "complete", "id": "xcsh-docs:actions:dns_zone_edit_cryptokey:fundamentals", "parent_id": null, "path": "docs/actions/dns_zone_edit_cryptokey.md", "provider_name": "dns_zone_edit_cryptokey", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_edit_cryptokey/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_dns_zone_edit_cryptokey for xcsh_dns_zone_edit_cryptokey.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](../guides/actions--dns_zone_edit_cryptokey--reference.md)
- [Examples](../guides/actions--dns_zone_edit_cryptokey--examples.md)
- [Lifecycle](../guides/actions--dns_zone_edit_cryptokey--lifecycle.md)
