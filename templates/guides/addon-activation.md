---
page_title: "Guide: Addon Service Activation"
subcategory: "Guides"
description: |-
  Report F5 Distributed Cloud addon service activation with Terraform and gate dependent resources.
  Covers Bot Defense, Client-Side Defense, WAAP, and more.
---

# Addon Service Activation

This guide explains how to activate, configure, and manage tenant addon services in F5 Distributed Cloud using Terraform. Topics covered include:

- **Check activation eligibility** — Determine whether an addon service is available for activation.
- **Understand activation workflows** — Learn where activation occurs and why activation is not managed as a Terraform resource.
- **Handle managed activation** — Coordinate services that require F5 account engagement.
- **Gate dependent resources** — Provision resources only after the addon reaches the active state.

## Overview

F5 Distributed Cloud addon services are optional security and performance features that you can activate for your tenant:

| Addon Service                        | Description                                   | Tier Required |
| ------------------------------------ | --------------------------------------------- | ------------- |
| `f5xc-bot-defense-standard`          | Protect applications from automated attacks   | STANDARD      |
| `f5xc-bot-defense-advanced`          | Bot defense with advanced ML detection        | ADVANCED      |
| `f5xc-client-side-defense-standard`  | Protect against Magecart and formjacking      | STANDARD      |
| `f5xc-waap-standard`                 | Web App and API Protection with API Discovery | STANDARD      |
| `f5xc-waap-advanced`                 | WAAP with full API security features          | ADVANCED      |
| `f5xc-malicious-user-detection`      | Identify malicious user behavior patterns     | ADVANCED      |
| `f5xc-synthetic-monitoring`          | Monitor application availability              | STANDARD      |

### Activation Types

Addon services have different activation types that determine how they can be activated:

```text
┌─────────────────────────────────────────────────────────────────────┐
│                     Activation Types                                │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  SELF-ACTIVATION                                                    │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐          │
│  │ Check Status │───►│ Activate in  │───►│ Active       │          │
│  │ (AS_NONE)    │    │ the console  │    │ (AS_SUBSCRIBED) │       │
│  └──────────────┘    └──────────────┘    └──────────────┘          │
│  Tenant admin activates; Terraform reports the result               │
│                                                                     │
│  PARTIALLY MANAGED                                                  │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐          │
│  │ Check Status │───►│ Request in   │───►│ SRE Review   │          │
│  │ (AS_NONE)    │    │ the console  │    │ (AS_PENDING) │          │
│  └──────────────┘    └──────────────┘    └──────┬───────┘          │
│                                                  │                  │
│                                          ┌──────▼───────┐          │
│                                          │ Active       │          │
│                                          │ (AS_SUBSCRIBED) │       │
│                                          └──────────────┘          │
│  User initiates, SRE team processes                                 │
│                                                                     │
│  FULLY MANAGED                                                      │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐          │
│  │ Contact      │───►│ Sales        │───►│ F5 Activates │          │
│  │ F5 Sales     │    │ Agreement    │    │ Addon        │          │
│  └──────────────┘    └──────────────┘    └──────────────┘          │
│  Requires sales engagement                                          │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

## Prerequisites

Before you begin, ensure you have:

- **Terraform >= 1.0** — Download from <https://www.terraform.io/downloads>
- **F5 Distributed Cloud Account** — Sign up at <https://www.f5.com/cloud/products/distributed-cloud-console>
- **API Credentials** — Configure API token or P12 certificate authentication
- **Appropriate Subscription Tier** — Most addon services require the ADVANCED tier

### Authentication Setup

Configure one of these authentication methods using environment variables:

#### Option 1: API Token (Recommended for development)

```bash
export XCSH_API_URL="https://<XC_TENANT>.console.ves.volterra.io"
export XCSH_API_TOKEN="<XC_API_TOKEN>"
```

#### Option 2: P12 Certificate (Recommended for production)

```bash
export XCSH_API_URL="https://<XC_TENANT>.console.ves.volterra.io"
export XCSH_P12_FILE="/path/to/credentials.p12"
export XCSH_P12_PASSWORD="<XC_P12_PASSWORD>"  # pragma: allowlist secret
```

## Quick Start

### Step 1: Clone the Repository

```bash
git clone https://github.com/f5-sales-demo/terraform-provider-xcsh.git
cd terraform-provider-xcsh/examples/guides/addon-activation
```

### Step 2: Configure Your Deployment

```bash
cp terraform.tfvars.example terraform.tfvars
```

Edit `terraform.tfvars` to select the addon services you want to report:

```hcl
# Enable Bot Defense reporting
enable_bot_defense = true

# Enable Client-Side Defense reporting
enable_client_side_defense = false
```

### Step 3: Initialize and Apply

```bash
terraform init
terraform plan
terraform apply
```

## Checking Activation Eligibility

Check where an addon stands for your tenant before relying on it in configuration.

### Using the Activation Status Data Source

```hcl
# Check whether Bot Defense can be activated
data "xcsh_addon_service_activation_status" "bot_defense" {
  addon_service = "f5xc-bot-defense-standard"
}

output "bot_defense_status" {
  value = {
    state        = data.xcsh_addon_service_activation_status.bot_defense.state
    can_activate = data.xcsh_addon_service_activation_status.bot_defense.can_activate
    message      = data.xcsh_addon_service_activation_status.bot_defense.message
  }
}
```

### State Values

| State           | Description            | Can Activate?        |
| --------------- | ---------------------- | -------------------- |
| `AS_NONE`       | Service not subscribed | Yes                  |
| `AS_PENDING`    | Activation in progress | No (wait)            |
| `AS_SUBSCRIBED` | Already active         | Already done         |
| `AS_ERROR`      | Subscription error     | No (contact support) |

### Querying Addon Service Details

```hcl
# Retrieve detailed information about an addon service
data "xcsh_addon_service" "bot_defense" {
  name = "f5xc-bot-defense-standard"
}

output "addon_details" {
  value = {
    display_name    = data.xcsh_addon_service.bot_defense.display_name
    tier            = data.xcsh_addon_service.bot_defense.tier
    activation_type = data.xcsh_addon_service.bot_defense.activation_type
  }
}
```

## Activating an Addon

The provider exposes the addon API as read-only: there is no `xcsh_addon_subscription`
resource, so Terraform reports activation rather than performing it directly. Activate
services in the F5 Distributed Cloud Console under **Administration > Billing &
Subscriptions**, or through your F5 account team for managed services.

This design is intentional. Activation is a tenant-level billing event rather than
an ephemeral infrastructure resource, so it does not belong in configuration plans
applied independently across multiple environments.

Follow this workflow:

1. Read `can_activate` and `state` to determine the current addon status.
2. Activate the service in the console once for the tenant.
3. Gate dependent resources on `state` reaching `AS_SUBSCRIBED`.

### Reporting Pending Addons

Collect the addons that a configuration requires and report any that still need
activation:

```hcl
locals {
  expected_addons = [
    "f5xc-bot-defense-standard",
    "f5xc-client-side-defense-standard",
    "f5xc-waap-standard",
  ]
}

data "xcsh_addon_service_activation_status" "expected" {
  for_each      = toset(local.expected_addons)
  addon_service = each.value
}

locals {
  active = [
    for addon in local.expected_addons : addon
    if data.xcsh_addon_service_activation_status.expected[addon].state == "AS_SUBSCRIBED"
  ]
  pending = setsubtract(local.expected_addons, local.active)
}

output "pending_addons" {
  description = "Addons requiring activation in the F5 Distributed Cloud Console"
  value       = local.pending
}
```

### Gating Dependent Resources

Use the reported state as a condition on any resource that requires the addon. When
a plan runs against a tenant where the addon is inactive, Terraform skips creating
dependent resources:

```hcl
data "xcsh_addon_service_activation_status" "bot_defense" {
  addon_service = "f5xc-bot-defense-standard"
}

resource "xcsh_http_loadbalancer" "with_bot_defense" {
  count = data.xcsh_addon_service_activation_status.bot_defense.state == "AS_SUBSCRIBED" ? 1 : 0

  name      = "example-protected-app"
  namespace = "production"
  domains   = ["app.example.com"]

  bot_defense {
    policy {
      name      = "example-bot-policy"
      namespace = "shared"
    }
  }

  http {
    port = 80
  }
}
```

### Asserting Addon Activation with Preconditions

When an addon is a mandatory requirement rather than an optional feature, use a
precondition so the plan stops with an informative message:

```hcl
resource "terraform_data" "require_bot_defense" {
  lifecycle {
    precondition {
      condition     = data.xcsh_addon_service_activation_status.bot_defense.state == "AS_SUBSCRIBED"
      error_message = "Bot Defense is not active on this tenant (state: ${data.xcsh_addon_service_activation_status.bot_defense.state}). Activate it under Administration > Billing & Subscriptions, then re-apply."
    }
  }
}
```

## Managed Activation Workflow

For addon services requiring sales contact, use Terraform to monitor status after F5 activates the service:

```hcl
# F5 activates managed addons; Terraform reports the resulting state
data "xcsh_addon_service_activation_status" "managed_addon" {
  addon_service = "some_managed_addon"
}

output "managed_addon_status" {
  value = {
    active  = data.xcsh_addon_service_activation_status.managed_addon.state == "AS_SUBSCRIBED"
    message = data.xcsh_addon_service_activation_status.managed_addon.message
  }
}

# Gate resources using conditional logic based on activation status
resource "xcsh_http_loadbalancer" "with_managed_feature" {
  count = data.xcsh_addon_service_activation_status.managed_addon.state == "AS_SUBSCRIBED" ? 1 : 0

  name      = "example-managed-app"
  namespace = "production"
  # ... remaining configuration
}
```

## Using Addon Features

After an addon is active, you can configure its features in your resources.

### Bot Defense in HTTP Load Balancer

```hcl
resource "xcsh_http_loadbalancer" "with_bot_defense" {
  count = data.xcsh_addon_service_activation_status.bot_defense.state == "AS_SUBSCRIBED" ? 1 : 0

  name      = "example-protected-app"
  namespace = "production"

  domains = ["app.example.com"]

  default_route_pools {
    pool {
      name      = xcsh_origin_pool.backend.name
      namespace = "production"
    }
    weight = 1
  }

  # Enable Bot Defense
  bot_defense {
    policy {
      name      = "example-bot-policy"
      namespace = "shared"
    }
  }

  http {
    port = 80
  }
}
```

### Client-Side Defense

```hcl
resource "xcsh_http_loadbalancer" "with_csd" {
  count = data.xcsh_addon_service_activation_status.client_side_defense.state == "AS_SUBSCRIBED" ? 1 : 0

  name      = "example-protected-app"
  namespace = "production"

  domains = ["app.example.com"]

  # Enable Client-Side Defense
  enable_client_side_defense = true

  # ... remaining configuration
}
```

## Troubleshooting

### Common Issues

#### Access denied when reading activation status

- Verify that your API credentials have addon management permissions.
- Verify that your tenant subscription tier supports the requested addon.

#### Activation remains in AS_PENDING state

- For partially managed addons, contact F5 support if review takes longer than expected.
- For self-activation, wait a few minutes and retry the plan.

#### State reports AS_ERROR

- Check the F5 Distributed Cloud Console for detailed error messages.
- Contact F5 support with your tenant ID and addon name.

### Debugging Status Output

```hcl
# Output detailed status for troubleshooting
output "debug_addon_status" {
  value = {
    addon_service = "f5xc-bot-defense-standard"
    state         = data.xcsh_addon_service_activation_status.bot_defense.state
    can_activate  = data.xcsh_addon_service_activation_status.bot_defense.can_activate
    message       = data.xcsh_addon_service_activation_status.bot_defense.message
  }
}
```

## Best Practices

1. **Check eligibility first** — Query the activation status data source before configuring dependent features.
2. **Use conditional resource creation** — Gate dependent resources on `state == "AS_SUBSCRIBED"`.
3. **Assert hard requirements** — Use lifecycle preconditions to halt execution when mandatory addons are inactive.
4. **Monitor activation state** — Track state progression for partially managed or pending addons.
5. **Document addon dependencies** — Specify required subscription tiers and addons in module documentation.

## Complete Example

See the [addon-activation example](https://github.com/f5-sales-demo/terraform-provider-xcsh/tree/main/examples/guides/addon-activation) for a complete, working Terraform configuration.

## Related Resources

- [HTTP Load Balancer Resource](../resources/http_loadbalancer.md)
