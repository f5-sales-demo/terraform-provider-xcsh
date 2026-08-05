# BotAllowlistPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotAllowlistPolicy by name
data "xcsh_bot_allowlist_policy" "example" {
  name      = "example-bot-allowlist-policy"
  namespace = "staging"
}

output "bot_allowlist_policy_id" {
  value = data.xcsh_bot_allowlist_policy.example.id
}
