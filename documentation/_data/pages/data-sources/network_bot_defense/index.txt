---
page_title: "xcsh_network_bot_defense"
subcategory: ""
description: "Bot Defense domains for an FQDN-aware firewall or proxy. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest."
xcsh_docs: {"aliases": ["network bot defense"], "body_bytes": 1541, "body_sha256": "sha256:b57a836038125bf818a6704bae3d6c4e6a0d95d093445237e3e6fe1d64a5fd22", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_bot_defense:reference", "xcsh-docs:data-sources:network_bot_defense:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_bot_defense:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_bot_defense:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_bot_defense/index.md", "product": "distributed-cloud", "provider_name": "network_bot_defense", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0033010203002120-1333301320122310-3023133010230310-1013120133033132-3201111033200121-2113002210321222-0020023233301212-2222022300321031", "registry_path": "docs/data-sources/network_bot_defense.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_bot_defense/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Bot Defense domains for an FQDN-aware firewall or proxy. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_bot_defense/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_bot_defense/examples/)
