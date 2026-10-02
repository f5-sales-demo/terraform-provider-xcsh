---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_ike2."
xcsh_docs: {"aliases": ["ike2"], "body_bytes": 11539, "body_sha256": "sha256:c52de5b24d6ef722f50e64fb376b2a10dae2ccc83f38e33556ae249967282131", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:ike2:properties:dh_group_set", "xcsh-docs:resources:ike2:properties:disable_pfs", "xcsh-docs:resources:ike2:properties:ike_keylifetime_hours", "xcsh-docs:resources:ike2:properties:ike_keylifetime_minutes", "xcsh-docs:resources:ike2:properties:timeouts", "xcsh-docs:resources:ike2:properties:use_default_keylifetime"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike2:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike2:reference", "parent_id": "xcsh-docs:resources:ike2:fundamentals", "path": "documentation/resources/ike2/properties/index.md", "product": "distributed-cloud", "provider_name": "ike2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0121100022303232-1202012003120020-0122303122032001-2001323012112121-3102220103221003-0032210001132223-1221211112013122-1321230021230203", "registry_path": "docs/guides/resources--ike2--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:ike2:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:ike2:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["dh group set"], "anchor": "section", "description": "Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of this profile.", "document_id": "xcsh-docs:resources:ike2:properties:dh_group_set", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dh_group_set"], "syntax": "block", "type": "object"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:ike2:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["disable pfs"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:ike2:properties:disable_pfs", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable_pfs"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:ike2:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["ike keylifetime hours"], "anchor": "section", "description": "Input Hours.", "document_id": "xcsh-docs:resources:ike2:properties:ike_keylifetime_hours", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ike_keylifetime_hours--duration", "enforcement": "provider-schema", "group": "ike_keylifetime_hours:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:ike2:properties:ike_keylifetime_hours", "type": "requires"}], "schema_path": ["ike_keylifetime_hours"], "syntax": "block", "type": "object"}, {"aliases": ["ike keylifetime minutes"], "anchor": "section", "description": "Set IKE Key Lifetime in minutes.", "document_id": "xcsh-docs:resources:ike2:properties:ike_keylifetime_minutes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ike_keylifetime_minutes--duration", "enforcement": "provider-schema", "group": "ike_keylifetime_minutes:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:ike2:properties:ike_keylifetime_minutes", "type": "requires"}], "schema_path": ["ike_keylifetime_minutes"], "syntax": "block", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:ike2:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:ike2:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:ike2:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:ike2:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["use default keylifetime"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:ike2:properties:use_default_keylifetime", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_default_keylifetime"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike2/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_ike2.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_ike2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/)
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

- [dh_group_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/dh_group_set/): complete subsection reference.

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

- [disable_pfs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/disable_pfs/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/ike_keylifetime_hours/): complete subsection reference.

- [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/ike_keylifetime_minutes/): complete subsection reference.

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

Name of the Ike2. Must be unique within the namespace.

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

Namespace where the Ike2 is created.

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

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/timeouts/): complete subsection reference.

- [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/use_default_keylifetime/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/#schema-description) |
| `dh_group_set` | [dh_group_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/dh_group_set/#section) |
| `dh_group_set.dh_groups` | [dh_group_set.dh_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/dh_group_set/#schema-dh_group_set--dh_groups) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/#schema-disable) |
| `disable_pfs` | [disable_pfs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/disable_pfs/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/#schema-id) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/ike_keylifetime_hours/#section) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/ike_keylifetime_hours/#schema-ike_keylifetime_hours--duration) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/ike_keylifetime_minutes/#section) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/ike_keylifetime_minutes/#schema-ike_keylifetime_minutes--duration) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/#schema-namespace) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/timeouts/#schema-timeouts--update) |
| `use_default_keylifetime` | [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/use_default_keylifetime/#section) |

## Next pages

- [dh_group_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/dh_group_set/)
- [disable_pfs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/disable_pfs/)
- [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/ike_keylifetime_hours/)
- [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/ike_keylifetime_minutes/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/timeouts/)
- [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/use_default_keylifetime/)
- [xcsh_ike2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/)
