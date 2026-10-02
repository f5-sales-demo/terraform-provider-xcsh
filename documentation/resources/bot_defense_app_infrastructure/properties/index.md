---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_defense_app_infrastructure."
xcsh_docs: {"aliases": ["bot defense app infrastructure"], "body_bytes": 16235, "body_sha256": "sha256:4b51d652c30c6744127a839e5c09b8865bd2889abac0bf503c16127e081df2f5", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted", "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted", "xcsh-docs:resources:bot_defense_app_infrastructure:properties:timeouts"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_defense_app_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_defense_app_infrastructure:reference", "parent_id": "xcsh-docs:resources:bot_defense_app_infrastructure:fundamentals", "path": "documentation/resources/bot_defense_app_infrastructure/properties/index.md", "product": "distributed-cloud", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2223101122200310-0021021103103022-3232011233203112-1202002312211021-1012103011330123-1022210000013032-1301123330002320-3003231322323010", "registry_path": "docs/guides/resources--bot_defense_app_infrastructure--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["cloud hosted"], "anchor": "section", "description": "Infra F5 Hosted.", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloud_hosted--infra_host_name", "enforcement": "provider-schema", "group": "cloud_hosted:RequiredObjectAttributes:egress,infra_host_name,ingress", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloud_hosted:RequiredObjectAttributes:egress,infra_host_name,ingress", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted:egress", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloud_hosted:RequiredObjectAttributes:egress,infra_host_name,ingress", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted:ingress", "type": "requires"}], "schema_path": ["cloud_hosted"], "syntax": "block", "type": "object"}, {"aliases": ["data center hosted"], "anchor": "section", "description": "Infra F5 Hosted.", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-data_center_hosted--infra_host_name", "enforcement": "provider-schema", "group": "data_center_hosted:RequiredObjectAttributes:egress,infra_host_name,ingress", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "data_center_hosted:RequiredObjectAttributes:egress,infra_host_name,ingress", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted:egress", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "data_center_hosted:RequiredObjectAttributes:egress,infra_host_name,ingress", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted:ingress", "type": "requires"}], "schema_path": ["data_center_hosted"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["environment type"], "anchor": "schema-environment_type", "description": "Environment Type Production environment Testing environment.", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["environment_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["traffic type"], "anchor": "schema-traffic_type", "description": "Traffic Type Web traffic Mobile traffic.", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["traffic_type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_defense_app_infrastructure/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_bot_defense_app_infrastructure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/)
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

- [cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/): complete subsection reference.

- [data_center_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/): complete subsection reference.

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

<a id="schema-environment_type"></a>

### environment_type property

Type: `"string"`. Optional, Computed.

\[Enum: PRODUCTION|TESTING\] Environment Type Production environment Testing environment. Possible
values are \`PRODUCTION\`, \`TESTING\`. Defaults to \`PRODUCTION\`.

Upstream description:

Environment Type

Production environment Testing environment.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PRODUCTION",
    "TESTING"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PRODUCTION",
  "enum": [
    "PRODUCTION",
    "TESTING"
  ],
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

Name of the Bot Defense App Infrastructure. Must be unique within the namespace.

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

Namespace where the Bot Defense App Infrastructure is created.

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

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/timeouts/): complete subsection reference.

<a id="schema-traffic_type"></a>

### traffic_type property

Type: `"string"`. Optional, Computed.

\[Enum: WEB|MOBILE\] Traffic Type Web traffic Mobile traffic. Possible values are \`WEB\`,
\`MOBILE\`. Defaults to \`WEB\`.

Upstream description:

Traffic Type

Web traffic Mobile traffic.

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
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/#schema-annotations) |
| `cloud_hosted` | [cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/#section) |
| `cloud_hosted.egress` | [cloud_hosted.egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/egress/#section) |
| `cloud_hosted.egress.ip_address` | [cloud_hosted.egress.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/egress/#schema-cloud_hosted--egress--ip_address) |
| `cloud_hosted.egress.location` | [cloud_hosted.egress.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/egress/#schema-cloud_hosted--egress--location) |
| `cloud_hosted.infra_host_name` | [cloud_hosted.infra_host_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/#schema-cloud_hosted--infra_host_name) |
| `cloud_hosted.ingress` | [cloud_hosted.ingress](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/ingress/#section) |
| `cloud_hosted.ingress.host_name` | [cloud_hosted.ingress.host_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/ingress/#schema-cloud_hosted--ingress--host_name) |
| `cloud_hosted.ingress.ip_address` | [cloud_hosted.ingress.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/ingress/#schema-cloud_hosted--ingress--ip_address) |
| `cloud_hosted.ingress.location` | [cloud_hosted.ingress.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/ingress/#schema-cloud_hosted--ingress--location) |
| `cloud_hosted.region` | [cloud_hosted.region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/#schema-cloud_hosted--region) |
| `data_center_hosted` | [data_center_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/#section) |
| `data_center_hosted.egress` | [data_center_hosted.egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/egress/#section) |
| `data_center_hosted.egress.ip_address` | [data_center_hosted.egress.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/egress/#schema-data_center_hosted--egress--ip_address) |
| `data_center_hosted.egress.location` | [data_center_hosted.egress.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/egress/#schema-data_center_hosted--egress--location) |
| `data_center_hosted.infra_host_name` | [data_center_hosted.infra_host_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/#schema-data_center_hosted--infra_host_name) |
| `data_center_hosted.ingress` | [data_center_hosted.ingress](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/ingress/#section) |
| `data_center_hosted.ingress.host_name` | [data_center_hosted.ingress.host_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/ingress/#schema-data_center_hosted--ingress--host_name) |
| `data_center_hosted.ingress.ip_address` | [data_center_hosted.ingress.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/ingress/#schema-data_center_hosted--ingress--ip_address) |
| `data_center_hosted.ingress.location` | [data_center_hosted.ingress.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/ingress/#schema-data_center_hosted--ingress--location) |
| `data_center_hosted.region` | [data_center_hosted.region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/#schema-data_center_hosted--region) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/#schema-disable) |
| `environment_type` | [environment_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/#schema-environment_type) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/#schema-namespace) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/timeouts/#schema-timeouts--update) |
| `traffic_type` | [traffic_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/#schema-traffic_type) |

## Next pages

- [cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/)
- [data_center_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/timeouts/)
- [xcsh_bot_defense_app_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/)
