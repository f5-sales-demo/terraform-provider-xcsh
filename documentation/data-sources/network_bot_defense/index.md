---
page_title: "xcsh_network_bot_defense"
subcategory: ""
description: "xcsh_network_bot_defense for xcsh_network_bot_defense."
xcsh_docs: {"aliases": [], "body_bytes": 1528, "body_sha256": "sha256:ab4d115a6af0a189abae46b696a8c9064b7be612ea2ae9bdc7de7119e1a1db0f", "child_ids": ["xcsh-docs:data-sources:network_bot_defense:reference", "xcsh-docs:data-sources:network_bot_defense:examples"], "collection_id": "xcsh-docs:data-sources:network_bot_defense:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_bot_defense:fundamentals", "parent_id": null, "path": "documentation/data-sources/network_bot_defense/index.md", "provider_name": "network_bot_defense", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_bot_defense/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_network_bot_defense for xcsh_network_bot_defense.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_bot_defense

Breadcrumbs:

- xcsh_network_bot_defense

Bot Defense domains for an FQDN-aware firewall or proxy. Values are bundled from the pinned OpenAPI
release; this data source performs no network request. Ports and traffic direction are not encoded
in the manifest.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_bot_defense/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_bot_defense/examples/)
