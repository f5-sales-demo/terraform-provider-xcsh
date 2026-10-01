---
page_title: "slack"
subcategory: ""
description: "slack for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1771, "body_sha256": "sha256:c040071235445e749dbf1a83e86d43a3b04350688e8f6a62b01170549d2fd9d0", "canonical_id": "xcsh-docs:data-sources:alert_receiver:properties:slack", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:slack:url"], "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:slack", "parent_id": "xcsh-docs:data-sources:alert_receiver:reference", "path": "docs/guides/data-sources--alert_receiver--properties--slack.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["slack"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/slack/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "slack for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# slack

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
- [Property reference](data-sources--alert_receiver--reference.md)
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [url](data-sources--alert_receiver--properties--slack--url.md): complete subsection reference.

## Next pages

- [slack.url](data-sources--alert_receiver--properties--slack--url.md)
- [Property reference](data-sources--alert_receiver--reference.md)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
