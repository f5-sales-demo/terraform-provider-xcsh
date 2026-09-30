---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_cdn."
xcsh_docs: {"aliases": [], "body_bytes": 957, "body_sha256": "sha256:7099152e34acb21d76ddcef5732b9650d65a71258d66fee6f4a538ae0af92b91", "canonical_id": "xcsh-docs:data-sources:network_cdn:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_cdn:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e239b2d01c82b76466c9b21dfd9551a3649d04c98c8f874243856b6d5396e241", "source_path": "examples/data-sources/xcsh_network_cdn/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_cdn:example:data-source", "parent_id": "xcsh-docs:data-sources:network_cdn:examples", "path": "docs/guides/data-sources--network_cdn--example--data-source.md", "provider_name": "network_cdn", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_cdn/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_cdn.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_network_cdn](../data-sources/network_cdn.md)
- [Examples](data-sources--network_cdn--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_cdn/data-source.tf`; digest `sha256:e239b2d01c82b76466c9b21dfd9551a3649d04c98c8f874243856b6d5396e241`.

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

## Next pages

- [Examples](data-sources--network_cdn--examples.md)
- [xcsh_network_cdn](../data-sources/network_cdn.md)
