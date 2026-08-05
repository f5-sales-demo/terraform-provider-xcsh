# Example: Check addon service activation status
# Use this to determine if an addon service can be activated

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}


data "xcsh_addon_service_activation_status" "bot_defense" {
  addon_service = "bot_defense"
}

# Output the activation status
output "state" {
  description = "Current subscription state (AS_NONE, AS_PENDING, AS_SUBSCRIBED, AS_ERROR)"
  value       = data.xcsh_addon_service_activation_status.bot_defense.state
}

output "can_activate" {
  description = "Whether the addon can be activated"
  value       = data.xcsh_addon_service_activation_status.bot_defense.can_activate
}

output "status_message" {
  description = "Human-readable status message"
  value       = data.xcsh_addon_service_activation_status.bot_defense.message
}

# Example: gate a dependent resource on the addon actually being active.
# Activation itself is performed in the F5 Distributed Cloud console; this data
# source reports the result, so a configuration can wait for it instead of
# building on a feature the tenant has not enabled.
resource "xcsh_namespace" "bot_defense_demo" {
  count = data.xcsh_addon_service_activation_status.bot_defense.state == "AS_SUBSCRIBED" ? 1 : 0

  name        = "bot-defense-demo"
  description = "Created only once Bot Defense is active on this tenant"
}
