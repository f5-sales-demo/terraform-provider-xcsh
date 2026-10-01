---
page_title: "xcsh_addon_service_activation_status landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_addon_service_activation_status landing."
---

# xcsh_addon_service_activation_status landing

<a id="canonical-d23bd9a2e3d652134f2d7eae64e927317103e2f5329e77bf01c22114d878a506"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-465315a473bbd9ac1d1144581392e4ffb2adc239bbf742af021f474da1e0ae3a"></a>

## xcsh_addon_service_activation_status — xcsh_addon_service_activation_status / 4dc401abc9be / 2

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

<a id="canonical-f38aa391e6bb05bf1274f6db963e79d182d042135ff4dee440d572f0cfba8b13"></a>

## Prerequisites — xcsh_addon_service_activation_status / 4dc401abc9be / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-86ad71cec148f85d1e602eb75f9d71e27dd4ce5a932309bf3a67498ffd547809"></a>

## Minimal configuration — xcsh_addon_service_activation_status / 4dc401abc9be / 4

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

<a id="canonical-7e106cde9575aa0d131c8a6a2286bcfbfcf439e5d379ddbe62b79296e4027049"></a>

## Root configuration — xcsh_addon_service_activation_status / 4dc401abc9be / 5

Required root properties: `addon_service`. Full root flags and choices appear in the property reference.

<a id="canonical-4aae7b5c4a4e340ef8409a5793869dc8f2fd01dea07de293f9c285e08af92708"></a>

## Next pages — xcsh_addon_service_activation_status / 4dc401abc9be / 6

- [Property reference](../guides/data-sources--addon_service_activation_status--reference--group-001.md#canonical-f41bda08979dbea479c6a4df981142036fa7233a29104511d138789f181b019b)
- [Examples](../guides/data-sources--addon_service_activation_status--examples--group-001.md#canonical-d291fe604a5df6f19ada2116def1f779abd33918a13909945c2a13c9a8ccc44e)
