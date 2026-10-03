---
page_title: "slack"
subcategory: ""
description: "Slack configuration to send alert notifications."
xcsh_docs: {"aliases": ["slack"], "body_bytes": 2460, "body_sha256": "sha256:3a7355d29890ac9eccf2be54248ac00d600167c9c8488a0af26e189e4a10fa58", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:slack:url"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:slack", "parent_id": "xcsh-docs:resources:alert_receiver:reference", "path": "documentation/resources/alert_receiver/properties/slack/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1132230210120230-1021200013202111-0112203322033000-1023212002032331-0132121001023021-0031132230200201-0322203213301311-1103312012003011", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [{"anchor": "schema-slack--channel", "enforcement": "provider-schema", "group": "slack:RequiredObjectAttributes:channel", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:slack", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["slack"], "schema_version": 1, "sections": [{"aliases": ["slack channel"], "anchor": "schema-slack--channel", "description": "Channel or user to send notifications to.", "document_id": "xcsh-docs:resources:alert_receiver:properties:slack", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["slack", "channel"], "syntax": "attribute", "type": "string"}, {"aliases": ["slack url"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:alert_receiver:properties:slack:url", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "slack.url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:slack:url:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "slack.url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:slack:url:clear_secret_info", "type": "conflicts"}], "schema_path": ["slack", "url"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/slack/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Slack configuration to send alert notifications.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# slack

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- slack

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Slack configuration to send alert notifications.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("channel")}
```

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

Terraform syntax:

```terraform
slack {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-slack--channel"></a>

### channel property

Type: `"string"`. Optional.

Channel or user to send notifications to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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

- [url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/url/): complete subsection reference.

## Next pages

- [slack.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/url/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
