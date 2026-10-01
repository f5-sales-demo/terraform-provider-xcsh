---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_bot_defense."
xcsh_docs: {"aliases": [], "body_bytes": 1175, "body_sha256": "sha256:44886f6bfb8036fb8986f4ba4bb8bb8a37377c76e6d499a398c5031e1e1ea25e", "canonical_id": "xcsh-docs:data-sources:network_bot_defense:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_bot_defense:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a428d9d820fed630e17a7faf15c103fb2b68704b5517319077da6fc62dccc081", "source_path": "examples/data-sources/xcsh_network_bot_defense/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_bot_defense:example:data-source", "parent_id": "xcsh-docs:data-sources:network_bot_defense:examples", "path": "docs/guides/data-sources--network_bot_defense--example--data-source.md", "provider_name": "network_bot_defense", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_bot_defense/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_bot_defense.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_bot_defense](../data-sources/network_bot_defense.md)
- [Examples](data-sources--network_bot_defense--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_bot_defense/data-source.tf`; digest `sha256:a428d9d820fed630e17a7faf15c103fb2b68704b5517319077da6fc62dccc081`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_bot_defense" "proxy" {}

# Configure these exact domains in a suffix-aware proxy or FQDN firewall.
output "bot_defense_https_proxy_rule" {
  value = {
    direction = "egress"
    protocol  = "tcp"
    port      = 443
    domains   = data.xcsh_network_bot_defense.proxy.domains
  }
}
```

## Next pages

- [Examples](data-sources--network_bot_defense--examples.md)
- [xcsh_network_bot_defense](../data-sources/network_bot_defense.md)
