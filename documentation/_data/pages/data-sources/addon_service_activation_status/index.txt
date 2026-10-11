---
page_title: "xcsh_addon_service_activation_status"
subcategory: ""
description: "Checks the activation status of an F5 Distributed Cloud Addon Service. Use this data source to determine if an addon service can be activated for your tenant and what the current subscription state is. **Possible state values:** | State | Description | | --------------- | ---------------------------------------- | |"
xcsh_docs: {"aliases": ["addon service activation status"], "body_bytes": 1958, "body_sha256": "sha256:e185916a04b8fb22a470f73752e8f78cc7db2d5e768dc09056a3e8fe246a877a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:addon_service_activation_status:reference", "xcsh-docs:data-sources:addon_service_activation_status:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:addon_service_activation_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:addon_service_activation_status:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/addon_service_activation_status/index.md", "product": "distributed-cloud", "provider_name": "addon_service_activation_status", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3102032331212202-3203311211020103-1033023113322232-1210322102130301-1301000332023311-0302213213132333-0001300202010110-3120132022110012", "registry_path": "docs/data-sources/addon_service_activation_status.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/addon_service_activation_status/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Checks the activation status of an F5 Distributed Cloud Addon Service. Use this data source to determine if an addon service can be activated for your tenant and what the current subscription state is. **Possible state values:** | State | Description | | --------------- | ---------------------------------------- | |", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service_activation_status/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service_activation_status/examples/)
