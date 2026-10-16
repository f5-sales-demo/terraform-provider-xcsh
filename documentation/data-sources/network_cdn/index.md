---
page_title: "xcsh_network_cdn"
subcategory: ""
description: "CDN IPv4 networks for origin or network-firewall ingress allowlists. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest."
xcsh_docs: {"aliases": ["network cdn"], "body_bytes": 1458, "body_sha256": "sha256:79aadef7819f3cf31e6d528ae348f24c6bec72df74cfd28421c1d8607ddfb126", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_cdn:reference", "xcsh-docs:data-sources:network_cdn:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_cdn:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_cdn:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_cdn/index.md", "product": "distributed-cloud", "provider_name": "network_cdn", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3332101011301320-0312003011132032-2233133133001013-2021212311222213-1103000302133332-3011321012313330-0002302023203120-2131131001121112", "registry_path": "docs/data-sources/network_cdn.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_cdn/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "CDN IPv4 networks for origin or network-firewall ingress allowlists. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": [], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_cdn/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_cdn/examples/)
