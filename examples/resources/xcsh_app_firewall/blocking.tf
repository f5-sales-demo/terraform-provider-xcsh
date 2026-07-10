# Blocking — Verified Configuration Example
# This configuration is extracted from acceptance tests
# and verified against the live F5 XC API.

resource "xcsh_app_firewall" "test" {
  name      = "example"
  namespace = "system"

  # Use default detection settings
  default_detection_settings {}

  # Blocking mode - actively block malicious requests
  blocking {}

  # allow_all_response_codes / use_default_blocking_page / default_bot_setting /
  # default_anonymization are server-default oneof markers the provider import-suppresses.
  # Declaring them makes the config import-unclean (config has them, imported state does
  # not), so they are intentionally omitted here — the server still materializes them.
}