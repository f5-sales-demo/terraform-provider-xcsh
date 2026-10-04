---
page_title: "routes"
subcategory: "Monitoring"
description: "Set of routes to match the incoming alert. The routes are evaluated in the specified order and terminates on the first match."
xcsh_docs: {"aliases": ["routes"], "body_bytes": 13768, "body_sha256": "sha256:911c7c86371dfc3618067d2ad46c096573a599b72f1a7fcb28d716304bdc2465", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_policy:properties:routes:any", "xcsh-docs:resources:alert_policy:properties:routes:custom", "xcsh-docs:resources:alert_policy:properties:routes:dont_send", "xcsh-docs:resources:alert_policy:properties:routes:group", "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters", "xcsh-docs:resources:alert_policy:properties:routes:send", "xcsh-docs:resources:alert_policy:properties:routes:severity"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:properties:routes", "parent_id": "xcsh-docs:resources:alert_policy:reference", "path": "documentation/resources/alert_policy/properties/routes/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033", "registry_path": "docs/guides/resources--alert_policy--reference--group-001.md", "relationships": [{"anchor": "schema-routes--alertname", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname,alertname_regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes", "type": "conflicts"}, {"anchor": "schema-routes--alertname", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname,any", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes", "type": "conflicts"}, {"anchor": "schema-routes--alertname", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname,custom", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes", "type": "conflicts"}, {"anchor": "schema-routes--alertname", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname,group", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes", "type": "conflicts"}, {"anchor": "schema-routes--alertname", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname,severity", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes", "type": "conflicts"}, {"anchor": "schema-routes--alertname_regex", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname,alertname_regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes", "type": "conflicts"}, {"anchor": "schema-routes--alertname_regex", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname_regex,any", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes", "type": "conflicts"}, {"anchor": "schema-routes--alertname_regex", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname_regex,custom", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes", "type": "conflicts"}, {"anchor": "schema-routes--alertname_regex", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname_regex,group", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes", "type": "conflicts"}, {"anchor": "schema-routes--alertname_regex", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname_regex,severity", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname,any", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname_regex,any", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:any,custom", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:any,group", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:any,severity", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname,custom", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname_regex,custom", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:any,custom", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:custom,group", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:custom,severity", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:dont_send,send", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:dont_send", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname,group", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname_regex,group", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:any,group", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:custom,group", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:group,severity", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:dont_send,send", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:send", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname,severity", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:severity", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:alertname_regex,severity", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:severity", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:any,severity", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:severity", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:custom,severity", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:severity", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:group,severity", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:severity", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes"], "schema_version": 1, "sections": [{"aliases": ["routes alertname"], "anchor": "schema-routes--alertname", "description": "List of Alert Names Customer tunnel interface down Physical Interface down Tunnel Interfaces to Customer Site Down Virtual Host server error Virtual Host client error Service Health Low Service Unavailable Virtual Host server error Virtual Host client error Endpoint Healthcheck failure Synthetic monitor health", "document_id": "xcsh-docs:resources:alert_policy:properties:routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "alertname"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes alertname regex"], "anchor": "schema-routes--alertname_regex", "description": "Exclusive with Regular Expression match for the alertname.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "alertname_regex"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:any", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "any"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes custom"], "anchor": "section", "description": "A set of matchers an alert has to fulfill to match the route.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:custom", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "custom"], "syntax": "block", "type": "object"}, {"aliases": ["routes dont send"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:dont_send", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "dont_send"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes group"], "anchor": "section", "description": "Select one or more known group names to match the incoming alert.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:group", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "group"], "syntax": "block", "type": "object"}, {"aliases": ["routes notification parameters"], "anchor": "section", "description": "Set of notification parameters to decide how and when the alert notifications should be sent to the receivers.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:custom,default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:custom,individual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:custom,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:custom,default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:default,individual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:default,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:custom,individual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:individual", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:default,individual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:individual", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:individual,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:individual", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:custom,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:ves_io_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:default,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:ves_io_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.notification_parameters:ConflictingObjectAttributes:individual,ves_io_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:ves_io_group", "type": "conflicts"}], "schema_path": ["routes", "notification_parameters"], "syntax": "block", "type": "object"}, {"aliases": ["routes send"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:send", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "send"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes severity"], "anchor": "section", "description": "Select one or more severity levels to match the incoming alert.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:severity", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "severity"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/routes/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Set of routes to match the incoming alert. The routes are evaluated in the specified order and terminates on the first match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["alert_policyCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/)
- routes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Set of routes to match the incoming alert. The routes are evaluated in the specified order and
terminates on the first match.

Upstream description:

Set of routes to match the incoming alert. The routes are evaluated in the specified order and
terminates on the first match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("alertname",
    "alertname_regex"),
  validators.ConflictingListObjectAttributes("alertname",
    "any"),
  validators.ConflictingListObjectAttributes("alertname",
    "custom"),
  validators.ConflictingListObjectAttributes("alertname",
    "group"),
  validators.ConflictingListObjectAttributes("alertname",
    "severity"),
  validators.ConflictingListObjectAttributes("alertname_regex",
    "any"),
  validators.ConflictingListObjectAttributes("alertname_regex",
    "custom"),
  validators.ConflictingListObjectAttributes("alertname_regex",
    "group"),
  validators.ConflictingListObjectAttributes("alertname_regex",
    "severity"),
  validators.ConflictingListObjectAttributes("any",
    "custom"),
  validators.ConflictingListObjectAttributes("any",
    "group"),
  validators.ConflictingListObjectAttributes("any",
    "severity"),
  validators.ConflictingListObjectAttributes("custom",
    "group"),
  validators.ConflictingListObjectAttributes("custom",
    "severity"),
  validators.ConflictingListObjectAttributes("dont_send",
    "send"),
  validators.ConflictingListObjectAttributes("group",
    "severity")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--alertname"></a>

### alertname property

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN",
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
    "L7_DDOS_AUTO_MITIGATION"),
}
```

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

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/any/): complete subsection reference.

- [custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/): complete subsection reference.

- [dont_send](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/dont_send/): complete subsection reference.

- [group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/group/): complete subsection reference.

- [notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/): complete subsection reference.

- [send](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/send/): complete subsection reference.

- [severity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/severity/): complete subsection reference.

## Next pages

- [routes.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/any/)
- [routes.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/)
- [routes.dont_send](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/dont_send/)
- [routes.group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/group/)
- [routes.notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/)
- [routes.send](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/send/)
- [routes.severity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/severity/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
