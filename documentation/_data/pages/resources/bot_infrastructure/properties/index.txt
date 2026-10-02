---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_infrastructure."
xcsh_docs: {"aliases": ["bot infrastructure"], "body_bytes": 12279, "body_sha256": "sha256:7070fb5b03b11c5a2332002a410567e2b275588d842f6b7b4bce2447918fcfd3", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted", "xcsh-docs:resources:bot_infrastructure:properties:timeouts"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_infrastructure:reference", "parent_id": "xcsh-docs:resources:bot_infrastructure:fundamentals", "path": "documentation/resources/bot_infrastructure/properties/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3211120122133210-0122110022200232-1211012113212211-2130003230230030-1201022221000022-0012221313230312-0000102002302203-0201002321021101", "registry_path": "docs/guides/resources--bot_infrastructure--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:bot_infrastructure:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["create cloud hosted"], "anchor": "section", "description": "F5 Cloud Hosted.", "document_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "create_cloud_hosted:ConflictingObjectAttributes:production,testing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "create_cloud_hosted:ConflictingObjectAttributes:production,testing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "type": "conflicts"}], "schema_path": ["create_cloud_hosted"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:bot_infrastructure:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:bot_infrastructure:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:bot_infrastructure:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:bot_infrastructure:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:bot_infrastructure:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:bot_infrastructure:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:bot_infrastructure:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["traffic type"], "anchor": "schema-traffic_type", "description": "The type of traffic that is routed to and processed by this infrastructure (Web or Mobile). Only web traffic, including browser-based traffic from mobile devices, is routed through this Bot Defense infrastructure. Only mobile traffic from native mobile apps with the Bot Defense SDK are routed through this Bot Defense", "document_id": "xcsh-docs:resources:bot_infrastructure:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["traffic_type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_infrastructure/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_bot_infrastructure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/)
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

- [create_cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/): complete subsection reference.

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

Name of the Bot Infrastructure. Must be unique within the namespace.

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

Namespace where the Bot Infrastructure is created.

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

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/timeouts/): complete subsection reference.

<a id="schema-traffic_type"></a>

### traffic_type property

Type: `"string"`. Optional, Computed.

\[Enum: WEB|MOBILE\] The type of traffic that is routed to and processed by this infrastructure (Web
or Mobile). Only web traffic, including browser-based traffic from mobile devices, is routed through
this Bot Defense infrastructure. Only mobile traffic from native mobile apps with the Bot Defense
SDK are routed.. Possible values are \`WEB\`, \`MOBILE\`. Defaults to \`WEB\`.

Upstream description:

The type of traffic that is routed to and processed by this infrastructure (Web or Mobile).

Only web traffic, including browser-based traffic from mobile devices, is routed through this Bot
Defense infrastructure. Only mobile traffic from native mobile apps with the Bot Defense SDK are
routed through this Bot Defense infrastructure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("WEB",
    "MOBILE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "WEB",
  "enum": [
    "WEB",
    "MOBILE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/#schema-annotations) |
| `create_cloud_hosted` | [create_cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/#section) |
| `create_cloud_hosted.ip_addresses` | [create_cloud_hosted.ip_addresses](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/#schema-create_cloud_hosted--ip_addresses) |
| `create_cloud_hosted.production` | [create_cloud_hosted.production](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/production/#section) |
| `create_cloud_hosted.production.region_1` | [create_cloud_hosted.production.region_1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/production/#schema-create_cloud_hosted--production--region_1) |
| `create_cloud_hosted.production.region_2` | [create_cloud_hosted.production.region_2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/production/#schema-create_cloud_hosted--production--region_2) |
| `create_cloud_hosted.testing` | [create_cloud_hosted.testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/testing/#section) |
| `create_cloud_hosted.testing.region_1` | [create_cloud_hosted.testing.region_1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/testing/#schema-create_cloud_hosted--testing--region_1) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/#schema-namespace) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/timeouts/#schema-timeouts--update) |
| `traffic_type` | [traffic_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/#schema-traffic_type) |

## Next pages

- [create_cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/timeouts/)
- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/)
