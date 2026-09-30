---
page_title: "xcsh_network_cdn"
subcategory: ""
description: "xcsh_network_cdn for xcsh_network_cdn."
xcsh_docs: {"aliases": [], "body_bytes": 1346, "body_sha256": "sha256:5f609a2cb84c79f347051251b0cd00cc9a4c5065f97f2366cf1d8d9a6c3c5ffb", "child_ids": ["xcsh-docs:data-sources:network_cdn:reference", "xcsh-docs:data-sources:network_cdn:examples"], "collection_id": "xcsh-docs:data-sources:network_cdn:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_cdn:fundamentals", "parent_id": null, "path": "documentation/data-sources/network_cdn/index.md", "provider_name": "network_cdn", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_cdn/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_network_cdn for xcsh_network_cdn.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
