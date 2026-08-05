# Addon Activation Example

This example demonstrates how to activate F5 Distributed Cloud addon services using Terraform.

## Overview

F5XC addon services are additional security and performance features that can be activated for your tenant. The provider exposes the addon API read-only, so this example reports activation state rather than changing it. It shows how to:

1. Query addon service details and tier requirements
2. Check activation eligibility for your tenant
3. Report which requested addons are active and which are still pending
4. Gate dependent resources on an addon actually being active

## Supported Addon Services

| Addon Service                       | Description                                   | Tier Required |
| ----------------------------------- | --------------------------------------------- | ------------- |
| `f5xc-bot-defense-standard`         | Protect applications from automated attacks   | STANDARD      |
| `f5xc-bot-defense-advanced`         | Bot defense with advanced ML detection        | ADVANCED      |
| `f5xc-client-side-defense-standard` | Protect against Magecart and formjacking      | STANDARD      |
| `f5xc-waap-standard`                | Web App and API Protection with API Discovery | STANDARD      |
| `f5xc-waap-advanced`                | WAAP with full API security features          | ADVANCED      |

## Prerequisites

- Terraform >= 1.0
- F5 Distributed Cloud account with appropriate subscription tier
- API credentials configured

## Quick Start

### 1. Configure Authentication

Set environment variables for authentication:

```bash
# Option 1: API Token
export XCSH_API_URL="https://your-tenant.console.ves.volterra.io"
export XCSH_API_TOKEN="your-api-token"

# Option 2: P12 Certificate
export XCSH_API_URL="https://your-tenant.console.ves.volterra.io"
export XCSH_P12_FILE="/path/to/credentials.p12"
export XCSH_P12_PASSWORD="your-p12-password"  # gitleaks:allow
```

### 2. Configure Variables

```bash
cp terraform.tfvars.example terraform.tfvars
```

Edit `terraform.tfvars` to enable desired addons:

```hcl
enable_bot_defense = true
enable_client_side_defense = false
enable_waap = false
```

### 3. Deploy

```bash
terraform init
terraform plan
terraform apply
```

## Outputs

After applying, you'll see:

- **Addon service info**: Details about each addon (tier, activation type)
- **Activation status**: Current state and whether activation is possible
- **Activation summary**: Which addons were activated

## How It Works

1. **Check eligibility**: The example uses `xcsh_addon_service_activation_status` data source to check if each addon can be activated
2. **Conditional activation**: Subscriptions are only created if the addon is available and not already active
3. **Wait for propagation**: An optional delay ensures addons are fully active before dependent resources are created

## Customization

### Activate Additional Addons

To report on other addon services, add a lookup and a status check to `main.tf`:

```hcl
# Example: Adding WAAP Advanced tier
data "xcsh_addon_service" "waap_advanced" {
  count = var.enable_waap_advanced ? 1 : 0
  name  = "f5xc-waap-advanced"
}

data "xcsh_addon_service_activation_status" "waap_advanced" {
  count         = var.enable_waap_advanced ? 1 : 0
  addon_service = "f5xc-waap-advanced"
}
```

Then add it to the `requested_addons` and `active_addons` lists in the `locals`
block so it shows up in the `pending_addons` output.

### Activating an Addon

The provider exposes the addon API read-only, so activation is done outside
Terraform:

1. Run `terraform apply` and read the `pending_addons` output.
2. For each addon listed, open the F5 Distributed Cloud console under
   Administration > Billing & Subscriptions and activate it. When the status
   check reports that self-service activation is unavailable, contact your F5
   account team instead.
3. Re-run `terraform apply`. Once `pending_addons` is empty, dependent
   resources gated on it are created.

## Troubleshooting

### Addon Not Activating

1. Check the activation status output for the `can_activate` and `state` values
2. Verify your subscription tier supports the addon
3. Check F5XC console for any pending approvals

### State Shows "AS_PENDING"

Some addons require SRE approval. Wait for the approval process to complete.

### State Shows "AS_ERROR"

Contact F5 support with your tenant ID and the specific error message.

## Related Documentation

- [Addon Activation Guide](https://registry.terraform.io/providers/f5-sales-demo/xcsh/latest/docs/guides/addon-activation)
- [xcsh_addon_service Data Source](https://registry.terraform.io/providers/f5-sales-demo/xcsh/latest/docs/data-sources/addon_service)
- [xcsh_addon_service_activation_status Data Source](https://registry.terraform.io/providers/f5-sales-demo/xcsh/latest/docs/data-sources/addon_service_activation_status)
