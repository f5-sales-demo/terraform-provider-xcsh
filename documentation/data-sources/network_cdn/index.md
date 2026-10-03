---
page_title: "xcsh_network_cdn"
subcategory: ""
description: "CDN IPv4 networks for origin or network-firewall ingress allowlists. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest."
xcsh_docs: {"aliases": ["network cdn"], "body_bytes": 1445, "body_sha256": "sha256:e9a24d4a0a36b1998315da56381e06bd2518e87cbff9477d877d5b0d169ccd14", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_cdn:reference", "xcsh-docs:data-sources:network_cdn:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_cdn:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_cdn:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_cdn/index.md", "product": "distributed-cloud", "provider_name": "network_cdn", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3332101011301320-0312003011132032-2233133133001013-2021212311222213-1103000302133332-3011321012313330-0002302023203120-2131131001121112", "registry_path": "docs/data-sources/network_cdn.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_cdn/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "CDN IPv4 networks for origin or network-firewall ingress allowlists. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_cdn

Breadcrumbs:

- xcsh_network_cdn

CDN IPv4 networks for origin or network-firewall ingress allowlists. Values are bundled from the
pinned OpenAPI release; this data source performs no network request. Ports and traffic direction
are not encoded in the manifest.

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

data "xcsh_network_cdn" "origin_ingress" {}

output "cdn_https_origin_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_cdn.origin_ingress.cidr_blocks
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_cdn/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_cdn/examples/)
