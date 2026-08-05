---
page_title: "Guide: Addon Service Activation"
subcategory: "Guides"
description: |-
  Report F5XC addon service activation with Terraform and gate resources on it.
  Covers Bot Defense, Client-Side Defense, WAAP, and more.
---

# Addon Service Activation

This guide covers F5 Distributed Cloud addon services in Terraform. The provider exposes the addon API read-only, so Terraform reports activation and gates on it rather than performing it. By the end, you'll understand how to:

- **Check activation eligibility** - Determine if an addon can be activated
- **Activate an addon** - where activation actually happens, and why it is not a Terraform resource
- **Handle managed activation** - Services requiring sales contact
- **Gate dependent resources** - create them only once the addon is active

## Overview

F5 Distributed Cloud addon services are additional security and performance features that can be activated for your tenant. These include:

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

- **Terraform >= 1.0** - The F5XC provider requires Terraform 1.0 or later
- **F5 Distributed Cloud Account** - Sign up at <https://www.f5.com/cloud/products/distributed-cloud-console>
- **API Credentials** - Token or P12 certificate authentication configured
- **Appropriate Subscription Tier** - Most addon services require ADVANCED tier

### Authentication Setup

Configure one of these authentication methods via environment variables:

#### Option 1: API Token (Recommended for development)

```bash
export XCSH_API_URL="https://your-tenant.console.ves.volterra.io"
export XCSH_API_TOKEN="your-api-token"
```

#### Option 2: P12 Certificate (Recommended for production)

```bash
export XCSH_API_URL="https://your-tenant.console.ves.volterra.io"
export XCSH_P12_FILE="/path/to/your-credentials.p12"
export XCSH_P12_PASSWORD="your-p12-password"  # pragma: allowlist secret
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

Edit `terraform.tfvars` to choose the addon services you want reported:

```hcl
# Enable Bot Defense activation
enable_bot_defense = true

# Enable Client-Side Defense
enable_client_side_defense = false
```

### Step 3: Initialize and Apply

```bash
terraform init
terraform plan
terraform apply
```

## Checking Activation Eligibility

Check where an addon stands for your tenant before relying on it.

### Using the Activation Status Data Source

```hcl
# Check if Bot Defense can be activated
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
# Get detailed information about an addon service
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

The provider exposes the addon API read-only: there is no `xcsh_addon_subscription`
resource, so Terraform reports activation rather than performing it. Activation
happens in the F5 Distributed Cloud console under **Administration > Billing &
Subscriptions**, or through your F5 account team for services that report managed
activation.

That split is deliberate rather than a gap. Activation is a billing event against
the tenant, not a piece of per-environment infrastructure, so it does not belong
in a plan that several environments apply independently.

The workflow is therefore:

1. Read `can_activate` and `state` to find out where an addon stands.
2. Activate it in the console, once, for the tenant.
3. Gate the resources that depend on the addon on `state` reaching `AS_SUBSCRIBED`.

### Reporting Pending Addons

Collect the addons a configuration expects and surface the ones that still need
attention. `pending` is the actionable output — anything in it has to be
activated in the console before dependent resources can be created.

```terraform
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
  description = "Activate these in the F5 Distributed Cloud console"
  value       = local.pending
}
```

### Gating Dependent Resources

Use the reported state as the condition on anything that needs the addon. A plan
run against a tenant where the addon is not active then creates nothing that
would depend on a feature that is not there.

```terraform
data "xcsh_addon_service_activation_status" "bot_defense" {
  addon_service = "f5xc-bot-defense-standard"
}

resource "xcsh_http_loadbalancer" "with_bot_defense" {
  count = data.xcsh_addon_service_activation_status.bot_defense.state == "AS_SUBSCRIBED" ? 1 : 0

  name      = "my-protected-app"
  namespace = "production"
  domains   = ["app.example.com"]

  bot_defense {
    policy {
      name      = "my-bot-policy"
      namespace = "shared"
    }
  }

  http {
    port = 80
  }
}
```

### Failing Loudly Instead of Silently Skipping

`count = 0` skips quietly, which is the right behaviour for an optional feature
and the wrong one for a hard requirement. When the addon is mandatory, assert it
so the plan stops with a readable message rather than applying a half-configured
environment.

```terraform
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

For addon services requiring sales contact, use Terraform to monitor status after F5 activates the service.

### Verifying Managed Addon Status

```hcl
# Managed addons are activated by F5; Terraform reports the resulting state
data "xcsh_addon_service_activation_status" "managed_addon" {
  addon_service = "some_managed_addon"
}

output "managed_addon_status" {
  value = {
    active  = data.xcsh_addon_service_activation_status.managed_addon.state == "AS_SUBSCRIBED"
    message = data.xcsh_addon_service_activation_status.managed_addon.message
  }
}

# Use conditional logic based on activation status
resource "xcsh_http_loadbalancer" "with_managed_feature" {
  count = data.xcsh_addon_service_activation_status.managed_addon.state == "AS_SUBSCRIBED" ? 1 : 0

  # Configuration that uses the managed addon feature
  name      = "lb-with-managed-addon"
  namespace = "production"
  # ... rest of configuration
}
```

## Using Addon Features

Once an addon is activated, you can use its features in your configurations.

### Bot Defense in HTTP Load Balancer

```hcl
resource "xcsh_http_loadbalancer" "with_bot_defense" {
  count = data.xcsh_addon_service_activation_status.bot_defense.state == "AS_SUBSCRIBED" ? 1 : 0

  name      = "my-protected-app"
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
      name      = "my-bot-policy"
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

  name      = "my-protected-app"
  namespace = "production"

  domains = ["app.example.com"]

  # Enable Client-Side Defense
  enable_client_side_defense = true

  # ... rest of configuration
}
```

## Troubleshooting

### Common Issues

#### Access denied when reading activation status

- Verify your API token has addon management permissions
- Check that your subscription tier supports the addon

#### Activation stuck in AS_PENDING

- For partially managed addons, contact F5 support
- For self-activation, wait and retry after a few minutes

#### State shows AS_ERROR

- Check F5XC console for detailed error messages
- Contact F5 support with your tenant ID

### Debugging Tips

```hcl
# Output detailed status for debugging
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

1. **Always check eligibility first** - Use the activation status data source before attempting activation
2. **Use conditional resource creation** - Gate dependent resources on `state` being `AS_SUBSCRIBED`
3. **Handle dependencies properly** - Use `depends_on` to ensure addons are active before using features
4. **Monitor activation state** - For partially managed addons, monitor the state for completion
5. **Document addon requirements** - Clearly document which addons your configuration requires

## Complete Example

See the [addon-activation example](https://github.com/f5-sales-demo/terraform-provider-xcsh/tree/main/examples/guides/addon-activation) for a complete, working Terraform configuration.

## Related Resources

- [HTTP Load Balancer Resource](../resources/http_loadbalancer.md)
