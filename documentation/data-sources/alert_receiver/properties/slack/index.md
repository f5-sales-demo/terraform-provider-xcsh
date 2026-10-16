---
page_title: "slack"
subcategory: ""
description: "Slack configuration to send alert notifications."
xcsh_docs: {"aliases": ["slack"], "body_bytes": 1712, "body_sha256": "sha256:01ba857e2a87cc5233ef265f907976c93f3fd1a5fd26108e21f82beb0c284e9a", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:slack:url"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:slack", "parent_id": "xcsh-docs:data-sources:alert_receiver:reference", "path": "documentation/data-sources/alert_receiver/properties/slack/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3033332003330212-1233103210333032-2131030123311030-0111310322012030-2122223002333001-2323111011321322-0321321013232130-1230032003100323", "registry_path": "docs/guides/data-sources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["slack"], "schema_version": 1, "sections": [{"aliases": ["slack channel"], "anchor": "schema-slack--channel", "description": "Channel or user to send notifications to.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:slack", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["slack", "channel"], "syntax": "attribute", "type": "string"}, {"aliases": ["slack url"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:slack:url", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["slack", "url"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/slack/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Slack configuration to send alert notifications.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
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
      "validatedAt": "2026-10-16T10:31:02+00:00"
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
