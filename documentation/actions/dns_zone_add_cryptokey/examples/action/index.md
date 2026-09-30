---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_dns_zone_add_cryptokey."
xcsh_docs: {"aliases": [], "body_bytes": 1064, "body_sha256": "sha256:b42d8218b8d99eed0d268803d6ba4a498c6407d72360022585d9a7058b1b88be", "child_ids": [], "collection_id": "xcsh-docs:actions:dns_zone_add_cryptokey:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:41e2e4052bac9b33dbdf74fdf0a795bf2458f8e723d6c49a082ace67a79ec803", "source_path": "examples/actions/xcsh_dns_zone_add_cryptokey/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:dns_zone_add_cryptokey:example:action", "parent_id": "xcsh-docs:actions:dns_zone_add_cryptokey:examples", "path": "documentation/actions/dns_zone_add_cryptokey/examples/action/index.md", "provider_name": "dns_zone_add_cryptokey", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "role": "example", "schema_path": ["action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_add_cryptokey/examples/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Action for xcsh_dns_zone_add_cryptokey.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_add_cryptokey/examples/)
- [xcsh_dns_zone_add_cryptokey](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_add_cryptokey/)
