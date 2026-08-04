# F5 Distributed Cloud Provider - Addon Activation Example
# =========================================================
#
# This example demonstrates how to inspect F5XC addon services with Terraform:
# which addons exist, which tier they require, and whether each one is already
# activated for the tenant.
#
# Activation itself is not performed here. The provider ships the read side of
# the addon API only — the xcsh_addon_service and
# xcsh_addon_service_activation_status data sources. Activate an addon in the
# F5 Distributed Cloud console (or through your F5 account team for
# managed-activation services), then use this configuration to confirm the
# state and to gate dependent resources on it.
#
# QUICK START:
# 1. Configure authentication (environment variables recommended)
# 2. Copy terraform.tfvars.example to terraform.tfvars
# 3. Enable the addons you want to report on in terraform.tfvars
# 4. Run: terraform init && terraform apply

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# =============================================================================
# PROVIDER CONFIGURATION
# =============================================================================
#
# Authentication via environment variables (recommended):
#   export XCSH_API_URL="https://your-tenant.console.ves.volterra.io"
#   export XCSH_API_TOKEN="your-api-token"
#
# Or for P12 certificate:
#   export XCSH_API_URL="https://your-tenant.console.ves.volterra.io"
#   export XCSH_P12_FILE="/path/to/credentials.p12"
#   export XCSH_P12_PASSWORD="your-p12-password"  # gitleaks:allow

provider "xcsh" {
  # Authentication via environment variables
}

# =============================================================================
# ADDON SERVICE INFORMATION
# =============================================================================

# Get details about Bot Defense addon service
data "xcsh_addon_service" "bot_defense" {
  count = var.enable_bot_defense ? 1 : 0
  name  = "f5xc-bot-defense-standard"
}

# Get details about Client-Side Defense addon service
data "xcsh_addon_service" "client_side_defense" {
  count = var.enable_client_side_defense ? 1 : 0
  name  = "f5xc-client-side-defense-standard"
}

# Get details about WAAP (Web App and API Protection) addon service
data "xcsh_addon_service" "waap" {
  count = var.enable_waap ? 1 : 0
  name  = "f5xc-waap-standard"
}

# =============================================================================
# ACTIVATION STATUS CHECKS
# =============================================================================

# Check Bot Defense activation status
data "xcsh_addon_service_activation_status" "bot_defense" {
  count         = var.enable_bot_defense ? 1 : 0
  addon_service = "f5xc-bot-defense-standard"
}

# Check Client-Side Defense activation status
data "xcsh_addon_service_activation_status" "client_side_defense" {
  count         = var.enable_client_side_defense ? 1 : 0
  addon_service = "f5xc-client-side-defense-standard"
}

# Check WAAP activation status
data "xcsh_addon_service_activation_status" "waap" {
  count         = var.enable_waap ? 1 : 0
  addon_service = "f5xc-waap-standard"
}

# =============================================================================
# ACTIVATION STATE
# =============================================================================
#
# There is no xcsh_addon_subscription resource: the provider exposes the addon
# API read-only, so Terraform reports activation rather than performing it.
# Activate an addon in the F5 Distributed Cloud console — Administration >
# Billing & Subscriptions — or through your F5 account team when the service
# reports managed activation.
#
# `state` is "AS_SUBSCRIBED" once an addon is usable. `can_activate` reports whether
# self-service activation is available to this tenant, which is the signal that
# an addon is eligible but has not been turned on yet.

locals {
  # Addons this configuration was asked to report on.
  requested_addons = compact([
    var.enable_bot_defense ? "f5xc-bot-defense-standard" : "",
    var.enable_client_side_defense ? "f5xc-client-side-defense-standard" : "",
    var.enable_waap ? "f5xc-waap-standard" : "",
  ])

  # Of those, the ones the tenant can already use.
  active_addons = compact([
    try(data.xcsh_addon_service_activation_status.bot_defense[0].state, "") == "AS_SUBSCRIBED" ? "f5xc-bot-defense-standard" : "",
    try(data.xcsh_addon_service_activation_status.client_side_defense[0].state, "") == "AS_SUBSCRIBED" ? "f5xc-client-side-defense-standard" : "",
    try(data.xcsh_addon_service_activation_status.waap[0].state, "") == "AS_SUBSCRIBED" ? "f5xc-waap-standard" : "",
  ])

  # Requested but not yet activated — act on these in the console.
  pending_addons = setsubtract(local.requested_addons, local.active_addons)
}

# =============================================================================
# EXAMPLE NAMESPACE (for demonstration)
# =============================================================================
#
# Gating on local.pending_addons is the pattern worth copying: dependent
# resources are created only once every addon they rely on is actually active,
# so a plan against an unactivated tenant does not build on a feature that is
# not there yet.

resource "xcsh_namespace" "demo" {
  count       = var.create_demo_namespace && length(local.pending_addons) == 0 ? 1 : 0
  name        = var.demo_namespace_name
  description = "Namespace for addon activation demonstration"
}
