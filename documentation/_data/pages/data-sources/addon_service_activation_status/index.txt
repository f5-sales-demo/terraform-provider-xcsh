---
page_title: "xcsh_addon_service_activation_status"
subcategory: ""
description: "Checks the activation status of an F5 Distributed Cloud Addon Service. Use this data source to determine if an addon service can be activated for your tenant and what the current subscription state is. **Possible state values:** | State | Description | | --------------- | ---------------------------------------- | |"
xcsh_docs: {"aliases": ["addon service activation status"], "body_bytes": 1945, "body_sha256": "sha256:9cafbbe62677231ffe777fb20e9ba622064fdbb0d09a6c102bbcd9c0a688875b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:addon_service_activation_status:reference", "xcsh-docs:data-sources:addon_service_activation_status:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:addon_service_activation_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:addon_service_activation_status:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/addon_service_activation_status/index.md", "product": "distributed-cloud", "provider_name": "addon_service_activation_status", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3102032331212202-3203311211020103-1033023113322232-1210322102130301-1301000332023311-0302213213132333-0001300202010110-3120132022110012", "registry_path": "docs/data-sources/addon_service_activation_status.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/addon_service_activation_status/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Checks the activation status of an F5 Distributed Cloud Addon Service. Use this data source to determine if an addon service can be activated for your tenant and what the current subscription state is. **Possible state values:** | State | Description | | --------------- | ---------------------------------------- | |", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_addon_service_activation_status

Breadcrumbs:

- xcsh_addon_service_activation_status

Checks the activation status of an F5 Distributed Cloud Addon Service.

Use this data source to determine if an addon service can be activated for your tenant and what the
current subscription state is.

\*\*Possible state values:\*\*

| State | Description | | --------------- | ---------------------------------------- | |
\`AS\_NONE\` | Default state, service not subscribed | | \`AS\_PENDING\` | Subscription request
pending activation | | \`AS\_SUBSCRIBED\` | Service is active and subscribed | | \`AS\_ERROR\` |
Subscription in error state |

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AddonServiceActivationStatus Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Check the tenant's Client-Side Defense subscription.
data "xcsh_addon_service_activation_status" "example" {
  addon_service = "f5xc-client-side-defense-standard"
}

output "addon_service_activation_state" {
  value = data.xcsh_addon_service_activation_status.example.state
}
```

## Root configuration

Required root properties: `addon_service`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service_activation_status/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service_activation_status/examples/)
