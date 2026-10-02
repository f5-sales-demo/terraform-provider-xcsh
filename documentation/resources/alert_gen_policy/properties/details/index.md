---
page_title: "details"
subcategory: ""
description: "Notification Details."
xcsh_docs: {"aliases": ["details"], "body_bytes": 4970, "body_sha256": "sha256:98d01bc2c73a4835f9773d6382886647c1f62137cc48bde4ec39c1d9f27d4bf2", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_gen_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_gen_policy:properties:details", "parent_id": "xcsh-docs:resources:alert_gen_policy:reference", "path": "documentation/resources/alert_gen_policy/properties/details/index.md", "product": "distributed-cloud", "provider_name": "alert_gen_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0320213020013121-1010322233231212-2300332320130011-3121330303302010-0001233132322230-3323233211022310-3113330122121321-2101130202032320", "registry_path": "docs/guides/resources--alert_gen_policy--reference--group-001.md", "relationships": [{"anchor": "schema-details--alert_message", "enforcement": "provider-schema", "group": "details:RequiredObjectAttributes:alert_message,alert_message_details,alert_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_gen_policy:properties:details", "type": "requires"}, {"anchor": "schema-details--alert_message_details", "enforcement": "provider-schema", "group": "details:RequiredObjectAttributes:alert_message,alert_message_details,alert_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_gen_policy:properties:details", "type": "requires"}, {"anchor": "schema-details--alert_name", "enforcement": "provider-schema", "group": "details:RequiredObjectAttributes:alert_message,alert_message_details,alert_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_gen_policy:properties:details", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["details"], "schema_version": 1, "sections": [{"aliases": ["alert message"], "anchor": "schema-details--alert_message", "description": "Alert Message.", "document_id": "xcsh-docs:resources:alert_gen_policy:properties:details", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "alert_message"], "syntax": "attribute", "type": "string"}, {"aliases": ["alert message details"], "anchor": "schema-details--alert_message_details", "description": "Detailed message of the alert.", "document_id": "xcsh-docs:resources:alert_gen_policy:properties:details", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "alert_message_details"], "syntax": "attribute", "type": "string"}, {"aliases": ["alert name"], "anchor": "schema-details--alert_name", "description": "Alert Name.", "document_id": "xcsh-docs:resources:alert_gen_policy:properties:details", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "alert_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["severity"], "anchor": "schema-details--severity", "description": "List of alert severities Minor Major Critical.", "document_id": "xcsh-docs:resources:alert_gen_policy:properties:details", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["details", "severity"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_gen_policy/properties/details/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Notification Details.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["alert_gen_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# details

Breadcrumbs:

- [xcsh_alert_gen_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/properties/)
- details

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Notification Details. Notification Details.

Upstream description:

Notification Details.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("alert_message",
    "alert_message_details",
    "alert_name")}
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
details {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-details--alert_message"></a>

### alert_message property

Type: `"string"`. Optional.

Alert Message. Alert Message.

Upstream description:

Alert Message.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="schema-details--alert_message_details"></a>

### alert_message_details property

Type: `"string"`. Optional.

Alert Message Details. Detailed message of the alert.

Upstream description:

Detailed message of the alert.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="schema-details--alert_name"></a>

### alert_name property

Type: `"string"`. Optional.

Alert Name. Alert Name.

Upstream description:

Alert Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 16,
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
    "ves.io.schema.rules.string.max_len": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "16"
  }
}
```

<a id="schema-details--severity"></a>

### severity property

Type: `"string"`. Optional.

\[Enum: MINOR|MAJOR|CRITICAL\] List of alert severities Minor Major Critical. Possible values are
\`MINOR\`, \`MAJOR\`, \`CRITICAL\`. Defaults to \`MINOR\`.

Upstream description:

List of alert severities

Minor Major Critical.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("MINOR",
    "MAJOR",
    "CRITICAL"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "MINOR",
  "enum": [
    "MINOR",
    "MAJOR",
    "CRITICAL"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/properties/)
- [xcsh_alert_gen_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_gen_policy/)
