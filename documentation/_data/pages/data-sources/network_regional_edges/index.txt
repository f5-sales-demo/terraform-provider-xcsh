---
page_title: "xcsh_network_regional_edges"
subcategory: ""
description: "xcsh_network_regional_edges for xcsh_network_regional_edges."
xcsh_docs: {"aliases": [], "body_bytes": 1674, "body_sha256": "sha256:8b05f140698a753fc1f820117efb453744793ad1e5aaae7244534c995cfce472", "child_ids": ["xcsh-docs:data-sources:network_regional_edges:reference", "xcsh-docs:data-sources:network_regional_edges:examples"], "collection_id": "xcsh-docs:data-sources:network_regional_edges:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_regional_edges:fundamentals", "parent_id": null, "path": "documentation/data-sources/network_regional_edges/index.md", "provider_name": "network_regional_edges", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_regional_edges/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_network_regional_edges for xcsh_network_regional_edges.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_regional_edges

Breadcrumbs:

- xcsh_network_regional_edges

Regional Edge IPv4 networks for origin ingress allowlists. Values are bundled from the pinned
OpenAPI release; this data source performs no network request. Ports and traffic direction are not
encoded in the manifest.

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

# Select the Regional Edge source networks that may initiate HTTPS connections
# to an origin. Omit regions to return all published regions.
data "xcsh_network_regional_edges" "origin_ingress" {
  regions = ["americas", "europe"]
}

output "https_origin_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_regional_edges.origin_ingress.cidr_blocks
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/examples/)
