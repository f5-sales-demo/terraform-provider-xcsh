---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_dns_zone_add_cryptokey."
xcsh_docs: {"aliases": [], "body_bytes": 957, "body_sha256": "sha256:17a9e2a5c0953ef311bb259dddb2e6ffc0e68147238c30da5e63a0ecfb21b748", "canonical_id": "xcsh-docs:actions:dns_zone_add_cryptokey:example:action", "child_ids": [], "collection_id": "xcsh-docs:actions:dns_zone_add_cryptokey:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:41e2e4052bac9b33dbdf74fdf0a795bf2458f8e723d6c49a082ace67a79ec803", "source_path": "examples/actions/xcsh_dns_zone_add_cryptokey/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:dns_zone_add_cryptokey:example:action", "parent_id": "xcsh-docs:actions:dns_zone_add_cryptokey:examples", "path": "docs/guides/actions--dns_zone_add_cryptokey--example--action.md", "provider_name": "dns_zone_add_cryptokey", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "publishing_destination": "registry", "role": "example", "schema_path": ["action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_add_cryptokey/examples/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Action for xcsh_dns_zone_add_cryptokey.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_dns_zone_add_cryptokey](../actions/dns_zone_add_cryptokey.md)
- [Examples](actions--dns_zone_add_cryptokey--examples.md)
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

## Next pages

- [Examples](actions--dns_zone_add_cryptokey--examples.md)
- [xcsh_dns_zone_add_cryptokey](../actions/dns_zone_add_cryptokey.md)
