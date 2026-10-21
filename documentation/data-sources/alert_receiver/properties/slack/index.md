---
page_title: "slack"
subcategory: ""
description: "Slack configuration to send alert notifications."
xcsh_docs: {"aliases": ["slack"], "body_bytes": 1712, "body_sha256": "sha256:9b8b763a02edb5a401745d7e951e4803cbcc6bdc2b188f74c2462e43e0dd4c79", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:slack:url"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:slack", "parent_id": "xcsh-docs:data-sources:alert_receiver:reference", "path": "documentation/data-sources/alert_receiver/properties/slack/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3033332003330212-1233103210333032-2131030123311030-0111310322012030-2122223002333001-2323111011321322-0321321013232130-1230032003100323", "registry_path": "docs/guides/data-sources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["slack"], "schema_version": 1, "sections": [{"aliases": ["slack channel"], "anchor": "schema-slack--channel", "description": "Channel or user to send notifications to.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:slack", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["slack", "channel"], "syntax": "attribute", "type": "string"}, {"aliases": ["slack url"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:slack:url", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["slack", "url"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/slack/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Slack configuration to send alert notifications.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# slack

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/)
- slack

<a id="section"></a>

Type: `"single"`. Computed.

Slack configuration to send alert notifications.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-slack--channel"></a>

### channel property

Type: `"string"`. Computed.

Channel or user to send notifications to.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-21T11:00:09+00:00"
    },
    "pattern": "^[a-z0-9-_]{1,80}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9-_]{1,80}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9-_]{1,80}$"
  }
}
```

- [url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/slack/url/): complete subsection reference.
