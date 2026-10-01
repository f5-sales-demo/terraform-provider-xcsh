---
page_title: "routes"
subcategory: "Monitoring"
description: "routes for xcsh_alert_policy."
xcsh_docs: {"aliases": [], "body_bytes": 10917, "body_sha256": "sha256:51f90abdd7342817bb01a9dfd9fa4ead6d3f32b3c40003a0c48449e55037896e", "child_ids": ["xcsh-docs:data-sources:alert_policy:properties:routes:any", "xcsh-docs:data-sources:alert_policy:properties:routes:custom", "xcsh-docs:data-sources:alert_policy:properties:routes:dont_send", "xcsh-docs:data-sources:alert_policy:properties:routes:group", "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters", "xcsh-docs:data-sources:alert_policy:properties:routes:send", "xcsh-docs:data-sources:alert_policy:properties:routes:severity"], "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:properties:routes", "parent_id": "xcsh-docs:data-sources:alert_policy:reference", "path": "documentation/data-sources/alert_policy/properties/routes/index.md", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes for xcsh_alert_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/)
- routes

<a id="section"></a>

Type: `"list"`. Computed.

Set of routes to match the incoming alert. The routes are evaluated in the specified order and
terminates on the first match.

Upstream description:

Set of routes to match the incoming alert. The routes are evaluated in the specified order and
terminates on the first match.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

## Direct properties

<a id="schema-routes--alertname"></a>

### alertname property

Type: `"string"`. Computed.

\[Enum:
SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN|SITE\_PHYSICAL\_INTERFACE\_DOWN|TUNNELS\_TO\_CUSTOMER\_SITE\_DOWN|SERVICE\_SERVER\_ERROR|SERVICE\_CLIENT\_ERROR|SERVICE\_HEALTH\_LOW|SERVICE\_UNAVAILABLE|SERVICE\_SERVER\_ERROR\_PER\_SOURCE\_SITE|SERVICE\_CLIENT\_ERROR\_PER\_SOURCE\_SITE|SERVICE\_ENDPOINT\_HEALTHCHECK\_FAILURE|SYNTHETIC\_MONITOR\_HEALTH\_CRITICAL|MALICIOUS\_USER\_DETECTED|WAF\_TOO\_MANY\_ATTACKS|API\_SECURITY\_TOO\_MANY\_ATTACKS|SERVICE\_POLICY\_TOO\_MANY\_ATTACKS|WAF\_TOO\_MANY\_MALICIOUS\_BOTS|BOT\_DEFENSE\_TOO\_MANY\_SECURITY\_EVENTS|THREAT\_CAMPAIGN|VES\_CLIENT\_SIDE\_DEFENSE\_SUSPICIOUS\_DOMAIN|VES\_CLIENT\_SIDE\_DEFENSE\_SENSITIVE\_FIELD\_READ|TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_FAILURE|TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_STILL\_FAILING|TLS\_AUTOMATIC\_CERTIFICATE\_EXPIRED|TLS\_CUSTOM\_CERTIFICATE\_EXPIRING|TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\_SOON|TLS\_CUSTOM\_CERTIFICATE\_EXPIRED|L7DDOS|DNS\_ZONE\_IGNORED\_DUPLICATE\_RECORD|API\_SECURITY\_UNUSED\_API\_DETECTED|API\_SECURITY\_SHADOW\_API\_DETECTED|API\_SECURITY\_SENSITIVE\_DATA\_IN\_RESPONSE\_DETECTED|API\_SECURITY\_RISK\_SCORE\_HIGH\_DETECTED|ROUTED\_DDOS\_ALERT\_NOTIFICATION|ROUTED\_DDOS\_MITIGATION\_NOTIFICATION|ROUTED\_DDOS\_TUNNEL\_STATUS\_UPDATE\_NOTIFICATION|L7\_DDOS\_AUTO\_MITIGATION\]
List of Alert Names Customer tunnel interface down Physical Interface down Tunnel Interfaces to
Customer Site Down Virtual Host server error Virtual Host client error Service Health Low Service
Unavailable Virtual Host server error Virtual Host client error Endpoint Healthcheck failure
Synthetic.. Possible values are \`SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN\`,
\`SITE\_PHYSICAL\_INTERFACE\_DOWN\`, \`TUNNELS\_TO\_CUSTOMER\_SITE\_DOWN\`,
\`SERVICE\_SERVER\_ERROR\`, \`SERVICE\_CLIENT\_ERROR\`, \`SERVICE\_HEALTH\_LOW\`,
\`SERVICE\_UNAVAILABLE\`, \`SERVICE\_SERVER\_ERROR\_PER\_SOURCE\_SITE\`,
\`SERVICE\_CLIENT\_ERROR\_PER\_SOURCE\_SITE\`, \`SERVICE\_ENDPOINT\_HEALTHCHECK\_FAILURE\`,
\`SYNTHETIC\_MONITOR\_HEALTH\_CRITICAL\`, \`MALICIOUS\_USER\_DETECTED\`,
\`WAF\_TOO\_MANY\_ATTACKS\`, \`API\_SECURITY\_TOO\_MANY\_ATTACKS\`,
\`SERVICE\_POLICY\_TOO\_MANY\_ATTACKS\`, \`WAF\_TOO\_MANY\_MALICIOUS\_BOTS\`,
\`BOT\_DEFENSE\_TOO\_MANY\_SECURITY\_EVENTS\`, \`THREAT\_CAMPAIGN\`,
\`VES\_CLIENT\_SIDE\_DEFENSE\_SUSPICIOUS\_DOMAIN\`,
\`VES\_CLIENT\_SIDE\_DEFENSE\_SENSITIVE\_FIELD\_READ\`,
\`TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_FAILURE\`,
\`TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_STILL\_FAILING\`, \`TLS\_AUTOMATIC\_CERTIFICATE\_EXPIRED\`,
\`TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\`, \`TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\_SOON\`,
\`TLS\_CUSTOM\_CERTIFICATE\_EXPIRED\`, \`L7DDOS\`, \`DNS\_ZONE\_IGNORED\_DUPLICATE\_RECORD\`,
\`API\_SECURITY\_UNUSED\_API\_DETECTED\`, \`API\_SECURITY\_SHADOW\_API\_DETECTED\`,
\`API\_SECURITY\_SENSITIVE\_DATA\_IN\_RESPONSE\_DETECTED\`,
\`API\_SECURITY\_RISK\_SCORE\_HIGH\_DETECTED\`, \`ROUTED\_DDOS\_ALERT\_NOTIFICATION\`,
\`ROUTED\_DDOS\_MITIGATION\_NOTIFICATION\`, \`ROUTED\_DDOS\_TUNNEL\_STATUS\_UPDATE\_NOTIFICATION\`,
\`L7\_DDOS\_AUTO\_MITIGATION\`. Defaults to \`SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN\`.

Upstream description:

List of Alert Names

Customer tunnel interface down Physical Interface down Tunnel Interfaces to Customer Site Down
Virtual Host server error Virtual Host client error Service Health Low Service Unavailable Virtual
Host server error Virtual Host client error Endpoint Healthcheck failure Synthetic monitor health
critical Malicious user detected Virtual Host WAF security events detected Virtual Host API security
events detected Virtual Host Service Policy security events detected Virtual Host Many Malicious
Bots based WAF security events detected Virtual Host Many Malicious Bots based Bot Defense security
events detected Virtual Host Many Threat campaign based WAF security events detected Suspicious
domain identified by Client-Side Defense service Client-Side Defense has identified a suspicious
script that is reading sensitive form field TLS Automatic Certificate renewal is failing TLS
Automatic Certificate renewal is still failing after multiple retries TLS Automatic Certificate has
expired TLS Custom Certificate will expire in less than 28 days TLS Custom Certificate will expire
in less than 15 days TLS Custom Certificate has expired DDoS security event detected DNS Zone
Ignored a Duplicate Record Create Request Unused APIs Detected Shadow APIs Detected Endpoints With
Sensitive Data In Response Detected High Risk Score Endpoints Detected A routed DDoS traffic anomaly
has been detected A routed DDoS mitigation has been implemented to block malicious traffic A routed
DDoS tunnel status has been changed L7 DDoS attack was detected, automatic mitigation is taking
place.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN",
  "enum": [
    "SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN",
    "SITE_PHYSICAL_INTERFACE_DOWN",
    "TUNNELS_TO_CUSTOMER_SITE_DOWN",
    "SERVICE_SERVER_ERROR",
    "SERVICE_CLIENT_ERROR",
    "SERVICE_HEALTH_LOW",
    "SERVICE_UNAVAILABLE",
    "SERVICE_SERVER_ERROR_PER_SOURCE_SITE",
    "SERVICE_CLIENT_ERROR_PER_SOURCE_SITE",
    "SERVICE_ENDPOINT_HEALTHCHECK_FAILURE",
    "SYNTHETIC_MONITOR_HEALTH_CRITICAL",
    "MALICIOUS_USER_DETECTED",
    "WAF_TOO_MANY_ATTACKS",
    "API_SECURITY_TOO_MANY_ATTACKS",
    "SERVICE_POLICY_TOO_MANY_ATTACKS",
    "WAF_TOO_MANY_MALICIOUS_BOTS",
    "BOT_DEFENSE_TOO_MANY_SECURITY_EVENTS",
    "THREAT_CAMPAIGN",
    "VES_CLIENT_SIDE_DEFENSE_SUSPICIOUS_DOMAIN",
    "VES_CLIENT_SIDE_DEFENSE_SENSITIVE_FIELD_READ",
    "TLS_AUTOMATIC_CERTIFICATE_RENEWAL_FAILURE",
    "TLS_AUTOMATIC_CERTIFICATE_RENEWAL_STILL_FAILING",
    "TLS_AUTOMATIC_CERTIFICATE_EXPIRED",
    "TLS_CUSTOM_CERTIFICATE_EXPIRING",
    "TLS_CUSTOM_CERTIFICATE_EXPIRING_SOON",
    "TLS_CUSTOM_CERTIFICATE_EXPIRED",
    "L7DDOS",
    "DNS_ZONE_IGNORED_DUPLICATE_RECORD",
    "API_SECURITY_UNUSED_API_DETECTED",
    "API_SECURITY_SHADOW_API_DETECTED",
    "API_SECURITY_SENSITIVE_DATA_IN_RESPONSE_DETECTED",
    "API_SECURITY_RISK_SCORE_HIGH_DETECTED",
    "ROUTED_DDOS_ALERT_NOTIFICATION",
    "ROUTED_DDOS_MITIGATION_NOTIFICATION",
    "ROUTED_DDOS_TUNNEL_STATUS_UPDATE_NOTIFICATION",
    "L7_DDOS_AUTO_MITIGATION"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-routes--alertname_regex"></a>

### alertname_regex property

Type: `"string"`. Computed.

Exclusive with \[alertname any custom group severity\] Regular Expression match for the alertname.

Upstream description:

Exclusive with \[alertname any custom group severity\] Regular Expression match for the alertname.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/any/): complete subsection reference.

- [custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/): complete subsection reference.

- [dont_send](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/dont_send/): complete subsection reference.

- [group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/group/): complete subsection reference.

- [notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/): complete subsection reference.

- [send](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/send/): complete subsection reference.

- [severity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/severity/): complete subsection reference.

## Next pages

- [routes.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/any/)
- [routes.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/)
- [routes.dont_send](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/dont_send/)
- [routes.group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/group/)
- [routes.notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/)
- [routes.send](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/send/)
- [routes.severity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/severity/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
