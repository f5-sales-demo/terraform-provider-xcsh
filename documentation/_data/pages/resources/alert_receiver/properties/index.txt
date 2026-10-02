---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_alert_receiver."
xcsh_docs: {"aliases": ["alert receiver"], "body_bytes": 34154, "body_sha256": "sha256:8d6233f45792b67404ec640f65ef1e054157758fd4a220eb84daf0dac8867c29", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:email", "xcsh-docs:resources:alert_receiver:properties:opsgenie", "xcsh-docs:resources:alert_receiver:properties:pagerduty", "xcsh-docs:resources:alert_receiver:properties:slack", "xcsh-docs:resources:alert_receiver:properties:sms", "xcsh-docs:resources:alert_receiver:properties:timeouts", "xcsh-docs:resources:alert_receiver:properties:webhook"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:reference", "parent_id": "xcsh-docs:resources:alert_receiver:fundamentals", "path": "documentation/resources/alert_receiver/properties/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:alert_receiver:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:alert_receiver:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:alert_receiver:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["email"], "anchor": "section", "description": "Email Configuration.", "document_id": "xcsh-docs:resources:alert_receiver:properties:email", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["email"], "syntax": "block", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:alert_receiver:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:alert_receiver:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:alert_receiver:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:alert_receiver:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["opsgenie"], "anchor": "section", "description": "OpsGenie configuration to send alert notifications.", "document_id": "xcsh-docs:resources:alert_receiver:properties:opsgenie", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-opsgenie--url", "enforcement": "provider-schema", "group": "opsgenie:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:opsgenie", "type": "requires"}], "schema_path": ["opsgenie"], "syntax": "block", "type": "object"}, {"aliases": ["pagerduty"], "anchor": "section", "description": "PagerDuty configuration to send alert notifications.", "document_id": "xcsh-docs:resources:alert_receiver:properties:pagerduty", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-pagerduty--url", "enforcement": "provider-schema", "group": "pagerduty:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:pagerduty", "type": "requires"}], "schema_path": ["pagerduty"], "syntax": "block", "type": "object"}, {"aliases": ["slack"], "anchor": "section", "description": "Slack configuration to send alert notifications.", "document_id": "xcsh-docs:resources:alert_receiver:properties:slack", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-slack--channel", "enforcement": "provider-schema", "group": "slack:RequiredObjectAttributes:channel", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:slack", "type": "requires"}], "schema_path": ["slack"], "syntax": "block", "type": "object"}, {"aliases": ["sms"], "anchor": "section", "description": "SMS Configuration.", "document_id": "xcsh-docs:resources:alert_receiver:properties:sms", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sms"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:alert_receiver:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["webhook"], "anchor": "section", "description": "Webhook configuration to send alert notifications.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_alert_receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

- [email](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/email/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Alert Receiver. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the Alert Receiver is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

- [opsgenie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/): complete subsection reference.

- [pagerduty](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/): complete subsection reference.

- [slack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/): complete subsection reference.

- [sms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/sms/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/timeouts/): complete subsection reference.

- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/#schema-disable) |
| `email` | [email](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/email/#section) |
| `email.email` | [email.email](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/email/#schema-email--email) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/#schema-namespace) |
| `opsgenie` | [opsgenie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/#section) |
| `opsgenie.api_key` | [opsgenie.api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/api_key/#section) |
| `opsgenie.api_key.blindfold_secret_info` | [opsgenie.api_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/api_key/blindfold_secret_info/#section) |
| `opsgenie.api_key.blindfold_secret_info.decryption_provider` | [opsgenie.api_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/api_key/blindfold_secret_info/#schema-opsgenie--api_key--blindfold_secret_info--decryption_provider) |
| `opsgenie.api_key.blindfold_secret_info.location` | [opsgenie.api_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/api_key/blindfold_secret_info/#schema-opsgenie--api_key--blindfold_secret_info--location) |
| `opsgenie.api_key.blindfold_secret_info.store_provider` | [opsgenie.api_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/api_key/blindfold_secret_info/#schema-opsgenie--api_key--blindfold_secret_info--store_provider) |
| `opsgenie.api_key.clear_secret_info` | [opsgenie.api_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/api_key/clear_secret_info/#section) |
| `opsgenie.api_key.clear_secret_info.provider_ref` | [opsgenie.api_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/api_key/clear_secret_info/#schema-opsgenie--api_key--clear_secret_info--provider_ref) |
| `opsgenie.api_key.clear_secret_info.url` | [opsgenie.api_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/api_key/clear_secret_info/#schema-opsgenie--api_key--clear_secret_info--url) |
| `opsgenie.url` | [opsgenie.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/#schema-opsgenie--url) |
| `pagerduty` | [pagerduty](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/#section) |
| `pagerduty.routing_key` | [pagerduty.routing_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/routing_key/#section) |
| `pagerduty.routing_key.blindfold_secret_info` | [pagerduty.routing_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/routing_key/blindfold_secret_info/#section) |
| `pagerduty.routing_key.blindfold_secret_info.decryption_provider` | [pagerduty.routing_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/routing_key/blindfold_secret_info/#schema-pagerduty--routing_key--blindfold_secret_info--decryption_provider) |
| `pagerduty.routing_key.blindfold_secret_info.location` | [pagerduty.routing_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/routing_key/blindfold_secret_info/#schema-pagerduty--routing_key--blindfold_secret_info--location) |
| `pagerduty.routing_key.blindfold_secret_info.store_provider` | [pagerduty.routing_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/routing_key/blindfold_secret_info/#schema-pagerduty--routing_key--blindfold_secret_info--store_provider) |
| `pagerduty.routing_key.clear_secret_info` | [pagerduty.routing_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/routing_key/clear_secret_info/#section) |
| `pagerduty.routing_key.clear_secret_info.provider_ref` | [pagerduty.routing_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/routing_key/clear_secret_info/#schema-pagerduty--routing_key--clear_secret_info--provider_ref) |
| `pagerduty.routing_key.clear_secret_info.url` | [pagerduty.routing_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/routing_key/clear_secret_info/#schema-pagerduty--routing_key--clear_secret_info--url) |
| `pagerduty.url` | [pagerduty.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/#schema-pagerduty--url) |
| `slack` | [slack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/#section) |
| `slack.channel` | [slack.channel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/#schema-slack--channel) |
| `slack.url` | [slack.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/url/#section) |
| `slack.url.blindfold_secret_info` | [slack.url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/url/blindfold_secret_info/#section) |
| `slack.url.blindfold_secret_info.decryption_provider` | [slack.url.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/url/blindfold_secret_info/#schema-slack--url--blindfold_secret_info--decryption_provider) |
| `slack.url.blindfold_secret_info.location` | [slack.url.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/url/blindfold_secret_info/#schema-slack--url--blindfold_secret_info--location) |
| `slack.url.blindfold_secret_info.store_provider` | [slack.url.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/url/blindfold_secret_info/#schema-slack--url--blindfold_secret_info--store_provider) |
| `slack.url.clear_secret_info` | [slack.url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/url/clear_secret_info/#section) |
| `slack.url.clear_secret_info.provider_ref` | [slack.url.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/url/clear_secret_info/#schema-slack--url--clear_secret_info--provider_ref) |
| `slack.url.clear_secret_info.url` | [slack.url.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/url/clear_secret_info/#schema-slack--url--clear_secret_info--url) |
| `sms` | [sms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/sms/#section) |
| `sms.contact_number` | [sms.contact_number](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/sms/#schema-sms--contact_number) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/timeouts/#schema-timeouts--update) |
| `webhook` | [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/#section) |
| `webhook.http_config` | [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/#section) |
| `webhook.http_config.auth_token` | [webhook.http_config.auth_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/#section) |
| `webhook.http_config.auth_token.token` | [webhook.http_config.auth_token.token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/token/#section) |
| `webhook.http_config.auth_token.token.blindfold_secret_info` | [webhook.http_config.auth_token.token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/token/blindfold_secret_info/#section) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.decryption_provider` | [webhook.http_config.auth_token.token.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/token/blindfold_secret_info/#schema-webhook--http_config--auth_token--token--blindfold_secret_info--decryption_provider) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.location` | [webhook.http_config.auth_token.token.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/token/blindfold_secret_info/#schema-webhook--http_config--auth_token--token--blindfold_secret_info--location) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.store_provider` | [webhook.http_config.auth_token.token.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/token/blindfold_secret_info/#schema-webhook--http_config--auth_token--token--blindfold_secret_info--store_provider) |
| `webhook.http_config.auth_token.token.clear_secret_info` | [webhook.http_config.auth_token.token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/token/clear_secret_info/#section) |
| `webhook.http_config.auth_token.token.clear_secret_info.provider_ref` | [webhook.http_config.auth_token.token.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/token/clear_secret_info/#schema-webhook--http_config--auth_token--token--clear_secret_info--provider_ref) |
| `webhook.http_config.auth_token.token.clear_secret_info.url` | [webhook.http_config.auth_token.token.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/auth_token/token/clear_secret_info/#schema-webhook--http_config--auth_token--token--clear_secret_info--url) |
| `webhook.http_config.basic_auth` | [webhook.http_config.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/#section) |
| `webhook.http_config.basic_auth.password` | [webhook.http_config.basic_auth.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/#section) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info` | [webhook.http_config.basic_auth.password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/blindfold_secret_info/#section) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.decryption_provider` | [webhook.http_config.basic_auth.password.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/blindfold_secret_info/#schema-webhook--http_config--basic_auth--password--blindfold_secret_info--decryption_provider) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.location` | [webhook.http_config.basic_auth.password.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/blindfold_secret_info/#schema-webhook--http_config--basic_auth--password--blindfold_secret_info--location) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.store_provider` | [webhook.http_config.basic_auth.password.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/blindfold_secret_info/#schema-webhook--http_config--basic_auth--password--blindfold_secret_info--store_provider) |
| `webhook.http_config.basic_auth.password.clear_secret_info` | [webhook.http_config.basic_auth.password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/clear_secret_info/#section) |
| `webhook.http_config.basic_auth.password.clear_secret_info.provider_ref` | [webhook.http_config.basic_auth.password.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/clear_secret_info/#schema-webhook--http_config--basic_auth--password--clear_secret_info--provider_ref) |
| `webhook.http_config.basic_auth.password.clear_secret_info.url` | [webhook.http_config.basic_auth.password.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/clear_secret_info/#schema-webhook--http_config--basic_auth--password--clear_secret_info--url) |
| `webhook.http_config.basic_auth.user_name` | [webhook.http_config.basic_auth.user_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/#schema-webhook--http_config--basic_auth--user_name) |
| `webhook.http_config.client_cert_obj` | [webhook.http_config.client_cert_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/#section) |
| `webhook.http_config.client_cert_obj.use_tls_obj` | [webhook.http_config.client_cert_obj.use_tls_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/use_tls_obj/#section) |
| `webhook.http_config.client_cert_obj.use_tls_obj.kind` | [webhook.http_config.client_cert_obj.use_tls_obj.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/use_tls_obj/#schema-webhook--http_config--client_cert_obj--use_tls_obj--kind) |
| `webhook.http_config.client_cert_obj.use_tls_obj.name` | [webhook.http_config.client_cert_obj.use_tls_obj.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/use_tls_obj/#schema-webhook--http_config--client_cert_obj--use_tls_obj--name) |
| `webhook.http_config.client_cert_obj.use_tls_obj.namespace` | [webhook.http_config.client_cert_obj.use_tls_obj.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/use_tls_obj/#schema-webhook--http_config--client_cert_obj--use_tls_obj--namespace) |
| `webhook.http_config.client_cert_obj.use_tls_obj.tenant` | [webhook.http_config.client_cert_obj.use_tls_obj.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/use_tls_obj/#schema-webhook--http_config--client_cert_obj--use_tls_obj--tenant) |
| `webhook.http_config.client_cert_obj.use_tls_obj.uid` | [webhook.http_config.client_cert_obj.use_tls_obj.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/use_tls_obj/#schema-webhook--http_config--client_cert_obj--use_tls_obj--uid) |
| `webhook.http_config.enable_http2` | [webhook.http_config.enable_http2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/#schema-webhook--http_config--enable_http2) |
| `webhook.http_config.follow_redirects` | [webhook.http_config.follow_redirects](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/#schema-webhook--http_config--follow_redirects) |
| `webhook.http_config.no_authorization` | [webhook.http_config.no_authorization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/no_authorization/#section) |
| `webhook.http_config.no_tls` | [webhook.http_config.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/no_tls/#section) |
| `webhook.http_config.use_tls` | [webhook.http_config.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/#section) |
| `webhook.http_config.use_tls.disable_sni` | [webhook.http_config.use_tls.disable_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/disable_sni/#section) |
| `webhook.http_config.use_tls.max_version` | [webhook.http_config.use_tls.max_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/#schema-webhook--http_config--use_tls--max_version) |
| `webhook.http_config.use_tls.min_version` | [webhook.http_config.use_tls.min_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/#schema-webhook--http_config--use_tls--min_version) |
| `webhook.http_config.use_tls.sni` | [webhook.http_config.use_tls.sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/#schema-webhook--http_config--use_tls--sni) |
| `webhook.http_config.use_tls.use_server_verification` | [webhook.http_config.use_tls.use_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/#section) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/#section) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/trusted_ca/#section) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.kind` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/trusted_ca/#schema-webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca--kind) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.name` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/trusted_ca/#schema-webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca--name) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.namespace` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/trusted_ca/#schema-webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca--namespace) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.tenant` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/trusted_ca/#schema-webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca--tenant) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.uid` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/trusted_ca/#schema-webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca--uid) |
| `webhook.http_config.use_tls.volterra_trusted_ca` | [webhook.http_config.use_tls.volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/volterra_trusted_ca/#section) |
| `webhook.url` | [webhook.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/url/#section) |
| `webhook.url.blindfold_secret_info` | [webhook.url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/url/blindfold_secret_info/#section) |
| `webhook.url.blindfold_secret_info.decryption_provider` | [webhook.url.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/url/blindfold_secret_info/#schema-webhook--url--blindfold_secret_info--decryption_provider) |
| `webhook.url.blindfold_secret_info.location` | [webhook.url.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/url/blindfold_secret_info/#schema-webhook--url--blindfold_secret_info--location) |
| `webhook.url.blindfold_secret_info.store_provider` | [webhook.url.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/url/blindfold_secret_info/#schema-webhook--url--blindfold_secret_info--store_provider) |
| `webhook.url.clear_secret_info` | [webhook.url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/url/clear_secret_info/#section) |
| `webhook.url.clear_secret_info.provider_ref` | [webhook.url.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/url/clear_secret_info/#schema-webhook--url--clear_secret_info--provider_ref) |
| `webhook.url.clear_secret_info.url` | [webhook.url.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/url/clear_secret_info/#schema-webhook--url--clear_secret_info--url) |

## Next pages

- [email](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/email/)
- [opsgenie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/)
- [pagerduty](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/pagerduty/)
- [slack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/slack/)
- [sms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/sms/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/timeouts/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
