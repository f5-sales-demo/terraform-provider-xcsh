---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_filter_set."
xcsh_docs: {"aliases": ["filter set"], "body_bytes": 10548, "body_sha256": "sha256:baa9697f53a3182b4d6d75b352142c45ab3607a04189412aa08b150e3eb9a4ec", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:filter_set:properties:filter_fields"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:filter_set:reference", "parent_id": "xcsh-docs:data-sources:filter_set:fundamentals", "path": "documentation/data-sources/filter_set/properties/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3103220200110331-2212303313301103-2220311222333203-2213110301121201-1003211001200312-3213023223222201-1232232300201201-0121121213112121", "registry_path": "docs/guides/data-sources--filter_set--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:filter_set:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["context key"], "anchor": "schema-context_key", "description": "Indexable context key that identifies a page or page type for which the FilterSet is applicable.", "document_id": "xcsh-docs:data-sources:filter_set:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["context_key"], "syntax": "attribute", "type": "string"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:filter_set:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["filter fields"], "anchor": "section", "description": "List of fields and their values selected by the user.", "document_id": "xcsh-docs:data-sources:filter_set:properties:filter_fields", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["filter_fields"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:filter_set:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:filter_set:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:filter_set:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:filter_set:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/filter_set/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_filter_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["filter_setCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

<a id="schema-context_key"></a>

### context_key property

Type: `"string"`. Computed.

Indexable context key that identifies a page or page type for which the FilterSet is applicable.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the FilterSet.

Upstream description:

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

- [filter_fields](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

Name of the FilterSet.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

Namespace where the FilterSet exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/#schema-annotations) |
| `context_key` | [context_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/#schema-context_key) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/#schema-description) |
| `filter_fields` | [filter_fields](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/#section) |
| `filter_fields.date_field` | [filter_fields.date_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/date_field/#section) |
| `filter_fields.date_field.absolute` | [filter_fields.date_field.absolute](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/date_field/absolute/#section) |
| `filter_fields.date_field.absolute.end_date` | [filter_fields.date_field.absolute.end_date](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/date_field/absolute/#schema-filter_fields--date_field--absolute--end_date) |
| `filter_fields.date_field.absolute.start_date` | [filter_fields.date_field.absolute.start_date](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/date_field/absolute/#schema-filter_fields--date_field--absolute--start_date) |
| `filter_fields.date_field.relative` | [filter_fields.date_field.relative](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/date_field/#schema-filter_fields--date_field--relative) |
| `filter_fields.field_id` | [filter_fields.field_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/#schema-filter_fields--field_id) |
| `filter_fields.filter_expression_field` | [filter_fields.filter_expression_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/filter_expression_field/#section) |
| `filter_fields.filter_expression_field.expression` | [filter_fields.filter_expression_field.expression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/filter_expression_field/#schema-filter_fields--filter_expression_field--expression) |
| `filter_fields.string_field` | [filter_fields.string_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/string_field/#section) |
| `filter_fields.string_field.field_values` | [filter_fields.string_field.field_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/string_field/#schema-filter_fields--string_field--field_values) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/#schema-namespace) |

## Next pages

- [filter_fields](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/)
- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/)
