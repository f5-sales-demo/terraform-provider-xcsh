---
page_title: "xcsh_addon_service"
subcategory: ""
description: "Retrieves information about an F5 Distributed Cloud Addon Service. Addon services are system-managed resources that provide additional functionality such as Bot Defense, Client Side Defense, and other security features. This data source allows you to query addon service details including tier requirements and"
xcsh_docs: {"aliases": ["addon service"], "body_bytes": 1773, "body_sha256": "sha256:4856001140f6cdfd3a27b1752faaa6c614d2066cb8759de8df5111096af529ca", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:addon_service:reference", "xcsh-docs:data-sources:addon_service:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:addon_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:addon_service:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/addon_service/index.md", "product": "distributed-cloud", "provider_name": "addon_service", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2330122101113210-0031213211000013-0113320231112112-2023300013333011-1030201010131303-2332021031023321-1030302210322030-1132200333322121", "registry_path": "docs/data-sources/addon_service.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/addon_service/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Retrieves information about an F5 Distributed Cloud Addon Service. Addon services are system-managed resources that provide additional functionality such as Bot Defense, Client Side Defense, and other security features. This data source allows you to query addon service details including tier requirements and", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_addon_service

Breadcrumbs:

- xcsh_addon_service

Retrieves information about an F5 Distributed Cloud Addon Service.

Addon services are system-managed resources that provide additional functionality such as Bot
Defense, Client Side Defense, and other security features. This data source allows you to query
addon service details including tier requirements and activation type.

~&gt; \*\*Note:\*\* Addon services cannot be created or modified via Terraform. To activate or
subscribe to an addon service, please use the F5 Distributed Cloud Console or contact your account
team.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AddonService Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AddonService by name
data "xcsh_addon_service" "example" {
  name      = "example-addon-service"
  namespace = "staging"
}

output "addon_service_id" {
  value = data.xcsh_addon_service.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/examples/)
