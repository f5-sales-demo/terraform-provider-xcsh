# BotDefenseAppInfrastructure Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotDefenseAppInfrastructure by name
data "xcsh_bot_defense_app_infrastructure" "example" {
  name      = "example-bot-defense-app-infrastructure"
  namespace = "staging"
}

output "bot_defense_app_infrastructure_id" {
  value = data.xcsh_bot_defense_app_infrastructure.example.id
}
