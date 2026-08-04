terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# F5 Distributed Cloud Provider - Addon Activation Outputs
# =========================================================

# =============================================================================
# ADDON SERVICE INFORMATION
# =============================================================================

output "bot_defense_info" {
  description = "Bot Defense addon service details"
  value = var.enable_bot_defense ? {
    display_name    = try(data.xcsh_addon_service.bot_defense[0].display_name, "N/A")
    tier            = try(data.xcsh_addon_service.bot_defense[0].tier, "N/A")
    activation_type = try(data.xcsh_addon_service.bot_defense[0].activation_type, "N/A")
  } : null
}

output "client_side_defense_info" {
  description = "Client-Side Defense addon service details"
  value = var.enable_client_side_defense ? {
    display_name    = try(data.xcsh_addon_service.client_side_defense[0].display_name, "N/A")
    tier            = try(data.xcsh_addon_service.client_side_defense[0].tier, "N/A")
    activation_type = try(data.xcsh_addon_service.client_side_defense[0].activation_type, "N/A")
  } : null
}

output "waap_info" {
  description = "WAAP addon service details"
  value = var.enable_waap ? {
    display_name    = try(data.xcsh_addon_service.waap[0].display_name, "N/A")
    tier            = try(data.xcsh_addon_service.waap[0].tier, "N/A")
    activation_type = try(data.xcsh_addon_service.waap[0].activation_type, "N/A")
  } : null
}

# =============================================================================
# ACTIVATION STATUS
# =============================================================================

output "bot_defense_status" {
  description = "Bot Defense activation status"
  value = var.enable_bot_defense ? {
    state        = try(data.xcsh_addon_service_activation_status.bot_defense[0].state, "NOT_CHECKED")
    can_activate = try(data.xcsh_addon_service_activation_status.bot_defense[0].can_activate, false)
    message      = try(data.xcsh_addon_service_activation_status.bot_defense[0].message, "Not checked")
  } : null
}

output "client_side_defense_status" {
  description = "Client-Side Defense activation status"
  value = var.enable_client_side_defense ? {
    state        = try(data.xcsh_addon_service_activation_status.client_side_defense[0].state, "NOT_CHECKED")
    can_activate = try(data.xcsh_addon_service_activation_status.client_side_defense[0].can_activate, false)
    message      = try(data.xcsh_addon_service_activation_status.client_side_defense[0].message, "Not checked")
  } : null
}

output "waap_status" {
  description = "WAAP activation status"
  value = var.enable_waap ? {
    state        = try(data.xcsh_addon_service_activation_status.waap[0].state, "NOT_CHECKED")
    can_activate = try(data.xcsh_addon_service_activation_status.waap[0].can_activate, false)
    message      = try(data.xcsh_addon_service_activation_status.waap[0].message, "Not checked")
  } : null
}

# =============================================================================
# ACTIVATION SUMMARY
# =============================================================================
#
# Reported, not effected: activation happens in the F5 Distributed Cloud console.
# `pending_addons` is the actionable output — anything listed there was requested
# in terraform.tfvars but is not active on the tenant yet.

output "active_addons" {
  description = "Requested addons that are active on this tenant"
  value       = local.active_addons
}

output "pending_addons" {
  description = "Requested addons that are not active yet; activate these in the F5 Distributed Cloud console"
  value       = local.pending_addons
}

output "activation_summary" {
  description = "Per-addon requested and active state"
  value = {
    total_requested = length(local.requested_addons)
    total_active    = length(local.active_addons)
    bot_defense = {
      requested = var.enable_bot_defense
      active    = contains(local.active_addons, "f5xc-bot-defense-standard")
    }
    client_side_defense = {
      requested = var.enable_client_side_defense
      active    = contains(local.active_addons, "f5xc-client-side-defense-standard")
    }
    waap = {
      requested = var.enable_waap
      active    = contains(local.active_addons, "f5xc-waap-standard")
    }
  }
}

# =============================================================================
# DEMO NAMESPACE
# =============================================================================

output "demo_namespace" {
  description = "Demo namespace details, created only once every requested addon is active"
  value = length(xcsh_namespace.demo) > 0 ? {
    name        = xcsh_namespace.demo[0].name
    description = xcsh_namespace.demo[0].description
  } : null
}
