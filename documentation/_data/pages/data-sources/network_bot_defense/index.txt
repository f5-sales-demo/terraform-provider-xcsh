---
page_title: "xcsh_network_bot_defense"
subcategory: ""
description: "Bot Defense domains for an FQDN-aware firewall or proxy. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest."
xcsh_docs: {"aliases": ["network bot defense"], "body_bytes": 1528, "body_sha256": "sha256:ab4d115a6af0a189abae46b696a8c9064b7be612ea2ae9bdc7de7119e1a1db0f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_bot_defense:reference", "xcsh-docs:data-sources:network_bot_defense:examples"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_bot_defense:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_bot_defense:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_bot_defense/index.md", "product": "distributed-cloud", "provider_name": "network_bot_defense", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0033010203002120-1333301320122310-3023133010230310-1013120133033132-3201111033200121-2113002210321222-0020023233301212-2222022300321031", "registry_path": "docs/data-sources/network_bot_defense.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_bot_defense/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Bot Defense domains for an FQDN-aware firewall or proxy. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
