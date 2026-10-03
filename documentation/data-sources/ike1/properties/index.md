---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_ike1."
xcsh_docs: {"aliases": ["ike1"], "body_bytes": 10446, "body_sha256": "sha256:2ca985b378d2b45a6897b7816213e26eb42634d885cf8ab7d8b290002a9f4c51", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:ike1:properties:ike_keylifetime_hours", "xcsh-docs:data-sources:ike1:properties:ike_keylifetime_minutes", "xcsh-docs:data-sources:ike1:properties:reauth_disabled", "xcsh-docs:data-sources:ike1:properties:reauth_timeout_days", "xcsh-docs:data-sources:ike1:properties:reauth_timeout_hours", "xcsh-docs:data-sources:ike1:properties:use_default_keylifetime"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike1:reference", "parent_id": "xcsh-docs:data-sources:ike1:fundamentals", "path": "documentation/data-sources/ike1/properties/index.md", "product": "distributed-cloud", "provider_name": "ike1", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013", "registry_path": "docs/guides/data-sources--ike1--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:ike1:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:ike1:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:ike1:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["ike keylifetime hours"], "anchor": "section", "description": "Input Hours.", "document_id": "xcsh-docs:data-sources:ike1:properties:ike_keylifetime_hours", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ike_keylifetime_hours"], "syntax": "attribute", "type": "object"}, {"aliases": ["ike keylifetime minutes"], "anchor": "section", "description": "Set IKE Key Lifetime in minutes.", "document_id": "xcsh-docs:data-sources:ike1:properties:ike_keylifetime_minutes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ike_keylifetime_minutes"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:ike1:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:ike1:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:ike1:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["reauth disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:ike1:properties:reauth_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["reauth_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "reauth timeout days"], "anchor": "section", "description": "Set Duration in days.", "document_id": "xcsh-docs:data-sources:ike1:properties:reauth_timeout_days", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["reauth_timeout_days"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "reauth timeout hours"], "anchor": "section", "description": "Input Hours.", "document_id": "xcsh-docs:data-sources:ike1:properties:reauth_timeout_hours", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["reauth_timeout_hours"], "syntax": "attribute", "type": "object"}, {"aliases": ["use default keylifetime"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:ike1:properties:use_default_keylifetime", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_default_keylifetime"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike1/properties/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Property reference for xcsh_ike1.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["ike1CreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/)
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

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the Ike1.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/ike_keylifetime_hours/): complete subsection reference.

- [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/ike_keylifetime_minutes/): complete subsection reference.

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

Name of the Ike1.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Namespace where the Ike1 exists.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [reauth_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/reauth_disabled/): complete subsection reference.

- [reauth_timeout_days](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/reauth_timeout_days/): complete subsection reference.

- [reauth_timeout_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/reauth_timeout_hours/): complete subsection reference.

- [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/use_default_keylifetime/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/#schema-id) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/ike_keylifetime_hours/#section) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/ike_keylifetime_hours/#schema-ike_keylifetime_hours--duration) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/ike_keylifetime_minutes/#section) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/ike_keylifetime_minutes/#schema-ike_keylifetime_minutes--duration) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/#schema-namespace) |
| `reauth_disabled` | [reauth_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/reauth_disabled/#section) |
| `reauth_timeout_days` | [reauth_timeout_days](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/reauth_timeout_days/#section) |
| `reauth_timeout_days.duration` | [reauth_timeout_days.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/reauth_timeout_days/#schema-reauth_timeout_days--duration) |
| `reauth_timeout_hours` | [reauth_timeout_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/reauth_timeout_hours/#section) |
| `reauth_timeout_hours.duration` | [reauth_timeout_hours.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/reauth_timeout_hours/#schema-reauth_timeout_hours--duration) |
| `use_default_keylifetime` | [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/use_default_keylifetime/#section) |

## Next pages

- [ike_keylifetime_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/ike_keylifetime_hours/)
- [ike_keylifetime_minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/ike_keylifetime_minutes/)
- [reauth_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/reauth_disabled/)
- [reauth_timeout_days](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/reauth_timeout_days/)
- [reauth_timeout_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/reauth_timeout_hours/)
- [use_default_keylifetime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/use_default_keylifetime/)
- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/)
