---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_protocol_inspection."
xcsh_docs: {"aliases": ["protocol inspection"], "body_bytes": 13401, "body_sha256": "sha256:eb59e081235d6ea5b877e6181221f5655235e065decc9b8954ae788af78e2a7e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks", "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures", "xcsh-docs:resources:protocol_inspection:properties:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_inspection:reference", "parent_id": "xcsh-docs:resources:protocol_inspection:fundamentals", "path": "documentation/resources/protocol_inspection/properties/index.md", "product": "distributed-cloud", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3210031210322130-3133210121332000-1300320030213333-1220303203113101-3030330132202312-1103302221030102-1202010201212331-3030222101313220", "registry_path": "docs/guides/resources--protocol_inspection--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["action"], "anchor": "schema-action", "description": "Action after inspection - ALLOW: Allow Allow traffic - DENY: Deny Throw RST error for TCP and ICMP error for UDP - DROP: DROP Silently drop traffic.", "document_id": "xcsh-docs:resources:protocol_inspection:reference", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ALLOW", "DENY", "DROP"], "version": 1}], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["action"], "syntax": "attribute", "type": "string"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:protocol_inspection:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:protocol_inspection:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:protocol_inspection:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["enable disable compliance checks"], "anchor": "section", "description": "Enable Disable Compliance Checks Choice.", "document_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_disable_compliance_checks:ConflictingObjectAttributes:disable_compliance_checks,enable_compliance_checks", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks:disable_compliance_checks", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_disable_compliance_checks:ConflictingObjectAttributes:disable_compliance_checks,enable_compliance_checks", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks:enable_compliance_checks", "type": "conflicts"}], "schema_path": ["enable_disable_compliance_checks"], "syntax": "block", "type": "object"}, {"aliases": ["enable disable signatures"], "anchor": "section", "description": "Enable Disable Signature Choice.", "document_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_disable_signatures:ConflictingObjectAttributes:disable_signature,enable_signature", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures:disable_signature", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_disable_signatures:ConflictingObjectAttributes:disable_signature,enable_signature", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_signatures:enable_signature", "type": "conflicts"}], "schema_path": ["enable_disable_signatures"], "syntax": "block", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:protocol_inspection:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:protocol_inspection:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:protocol_inspection:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:protocol_inspection:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:protocol_inspection:properties:timeouts", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_inspection/properties/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Property reference for xcsh_protocol_inspection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/)
- Property reference

## Direct properties

<a id="schema-action"></a>

### action property

Type: `"string"`. Optional, Computed.

\[Enum: ALLOW|DENY|DROP\] Action after inspection - ALLOW: Allow Allow traffic - DENY: Deny Throw
RST error for TCP and ICMP error for UDP - DROP: DROP Silently drop traffic. Possible values are
\`ALLOW\`, \`DENY\`, \`DROP\`. Defaults to \`ALLOW\`. Server applies default when omitted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALLOW","DENY","DROP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ALLOW",
    "DENY",
    "DROP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ALLOW",
  "enum": [
    "ALLOW",
    "DENY",
    "DROP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Additional upstream details:

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

- [enable_disable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_compliance_checks/): complete subsection reference.

- [enable_disable_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_signatures/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Additional upstream details:

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

Name of the Protocol Inspection. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Namespace where the Protocol Inspection is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/#schema-action) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/#schema-disable) |
| `enable_disable_compliance_checks` | [enable_disable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_compliance_checks/#section) |
| `enable_disable_compliance_checks.disable_compliance_checks` | [enable_disable_compliance_checks.disable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_compliance_checks/disable_compliance_checks/#section) |
| `enable_disable_compliance_checks.enable_compliance_checks` | [enable_disable_compliance_checks.enable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_compliance_checks/enable_compliance_checks/#section) |
| `enable_disable_compliance_checks.enable_compliance_checks.name` | [enable_disable_compliance_checks.enable_compliance_checks.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_compliance_checks/enable_compliance_checks/#schema-enable_disable_compliance_checks--enable_compliance_checks--name) |
| `enable_disable_compliance_checks.enable_compliance_checks.namespace` | [enable_disable_compliance_checks.enable_compliance_checks.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_compliance_checks/enable_compliance_checks/#schema-enable_disable_compliance_checks--enable_compliance_checks--namespace) |
| `enable_disable_compliance_checks.enable_compliance_checks.tenant` | [enable_disable_compliance_checks.enable_compliance_checks.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_compliance_checks/enable_compliance_checks/#schema-enable_disable_compliance_checks--enable_compliance_checks--tenant) |
| `enable_disable_signatures` | [enable_disable_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_signatures/#section) |
| `enable_disable_signatures.disable_signature` | [enable_disable_signatures.disable_signature](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_signatures/disable_signature/#section) |
| `enable_disable_signatures.enable_signature` | [enable_disable_signatures.enable_signature](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_signatures/enable_signature/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/#schema-namespace) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/timeouts/#schema-timeouts--update) |
