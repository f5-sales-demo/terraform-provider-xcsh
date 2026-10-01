---
page_title: "xcsh_network_data_intelligence"
subcategory: ""
description: "xcsh_network_data_intelligence for xcsh_network_data_intelligence."
xcsh_docs: {"aliases": [], "body_bytes": 1426, "body_sha256": "sha256:aef61a037980f5be10dd5dbe2fadd4f652ffa9e88167eae4c7a517dff3204ae3", "canonical_id": "xcsh-docs:data-sources:network_data_intelligence:fundamentals", "child_ids": ["xcsh-docs:data-sources:network_data_intelligence:reference", "xcsh-docs:data-sources:network_data_intelligence:examples"], "collection_id": "xcsh-docs:data-sources:network_data_intelligence:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_data_intelligence:fundamentals", "parent_id": null, "path": "docs/data-sources/network_data_intelligence.md", "provider_name": "network_data_intelligence", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_data_intelligence/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_network_data_intelligence for xcsh_network_data_intelligence.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_data_intelligence

Breadcrumbs:

- xcsh_network_data_intelligence

Regional Data Intelligence IPv4 destinations. Values are bundled from the pinned OpenAPI release;
this data source performs no network request. Ports and traffic direction are not encoded in the
manifest.

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

data "xcsh_network_data_intelligence" "us" {
  regions = ["us"]
}

output "data_intelligence_https_egress" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 443
    destinations = data.xcsh_network_data_intelligence.us.cidr_blocks
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--network_data_intelligence--reference.md)
- [Examples](../guides/data-sources--network_data_intelligence--examples.md)
