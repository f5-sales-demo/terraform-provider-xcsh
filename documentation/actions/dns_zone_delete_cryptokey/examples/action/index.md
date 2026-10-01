---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_dns_zone_delete_cryptokey."
xcsh_docs: {"aliases": [], "body_bytes": 1190, "body_sha256": "sha256:68e437fa6563dd3784aa12cc8b0fe7f056af675471ad43616f9b4bf38e65cbd1", "child_ids": [], "collection_id": "xcsh-docs:actions:dns_zone_delete_cryptokey:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:41efc29a59313c81b918dfd8dc2bfaca42cc84009a5aa5450b2c6881b14ad7fe", "source_path": "examples/actions/xcsh_dns_zone_delete_cryptokey/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:dns_zone_delete_cryptokey:example:action", "parent_id": "xcsh-docs:actions:dns_zone_delete_cryptokey:examples", "path": "documentation/actions/dns_zone_delete_cryptokey/examples/action/index.md", "provider_name": "dns_zone_delete_cryptokey", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "role": "example", "schema_path": ["action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_delete_cryptokey/examples/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Action for xcsh_dns_zone_delete_cryptokey.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_dns_zone_delete_cryptokey](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_delete_cryptokey/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_delete_cryptokey/examples/)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_dns_zone_delete_cryptokey/action.tf`; digest `sha256:41efc29a59313c81b918dfd8dc2bfaca42cc84009a5aa5450b2c6881b14ad7fe`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_delete_cryptokey/examples/)
- [xcsh_dns_zone_delete_cryptokey](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_delete_cryptokey/)
