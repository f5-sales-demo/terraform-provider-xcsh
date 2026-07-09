# BotDetectionRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotDetectionRule by name
data "xcsh_bot_detection_rule" "example" {
  name      = "example-bot-detection-rule"
  namespace = "staging"
}

output "bot_detection_rule_id" {
  value = data.xcsh_bot_detection_rule.example.id
}
