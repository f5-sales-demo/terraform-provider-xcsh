---
page_title: "Property reference"
subcategory: "API Management"
description: "Property reference for xcsh_api_definition."
xcsh_docs: {"aliases": ["api definition"], "body_bytes": 12049, "body_sha256": "sha256:66d01976a61c9e99e31d454e8269fc27171853091f5065a6f17409638f732287", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_definition:properties:api_inventory_exclusion_list", "xcsh-docs:data-sources:api_definition:properties:api_inventory_inclusion_list", "xcsh-docs:data-sources:api_definition:properties:mixed_schema_origin", "xcsh-docs:data-sources:api_definition:properties:non_api_endpoints", "xcsh-docs:data-sources:api_definition:properties:strict_schema_origin"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_definition:reference", "parent_id": "xcsh-docs:data-sources:api_definition:fundamentals", "path": "documentation/data-sources/api_definition/properties/index.md", "product": "distributed-cloud", "provider_name": "api_definition", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3322310000110031-3303331030310132-3201123301111320-0223031203232010-1233333113000003-1023202201203132-0311103213220013-0232300212311123", "registry_path": "docs/guides/data-sources--api_definition--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:api_definition:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["api inventory exclusion list"], "anchor": "section", "description": "List of API Endpoints excluded from the API Inventory.", "document_id": "xcsh-docs:data-sources:api_definition:properties:api_inventory_exclusion_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_inventory_exclusion_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["api inventory inclusion list"], "anchor": "section", "description": "List of API Endpoints included in the API Inventory. Typically, discovered API endpoints are added to the API Inventory using this list.", "document_id": "xcsh-docs:data-sources:api_definition:properties:api_inventory_inclusion_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_inventory_inclusion_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:api_definition:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:api_definition:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:api_definition:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["backend servers", "mixed schema origin", "origin servers", "upstream servers"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:api_definition:properties:mixed_schema_origin", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mixed_schema_origin"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:api_definition:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:api_definition:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["non api endpoints"], "anchor": "section", "description": "List of Non-API Endpoints.", "document_id": "xcsh-docs:data-sources:api_definition:properties:non_api_endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["non_api_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "origin servers", "strict schema origin", "upstream servers"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:api_definition:properties:strict_schema_origin", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["strict_schema_origin"], "syntax": "attribute", "type": "object"}, {"aliases": ["swagger specs"], "anchor": "schema-swagger_specs", "description": "URLs of versioned OpenAPI files uploaded through Web App & API Protection > Files > Swagger Files. The 512-byte item limit is a URL-length limit; inline string:/// OpenAPI content is rejected and does not create an API-definition object.", "document_id": "xcsh-docs:data-sources:api_definition:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["swagger_specs"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_definition/properties/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Property reference for xcsh_api_definition.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["api_definitionCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

- [api_inventory_exclusion_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/api_inventory_exclusion_list/): complete subsection reference.

- [api_inventory_inclusion_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/api_inventory_inclusion_list/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the APIDefinition.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

- [mixed_schema_origin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/mixed_schema_origin/): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the APIDefinition.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Namespace where the APIDefinition exists.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [non_api_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/non_api_endpoints/): complete subsection reference.

- [strict_schema_origin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/strict_schema_origin/): complete subsection reference.

<a id="schema-swagger_specs"></a>

### swagger_specs property

Type: `["list", "string"]`. Computed.

URLs of versioned OpenAPI files uploaded through Web App &amp; API Protection &gt; Files &gt;
Swagger Files. The 512-byte item limit is a URL-length limit; inline string:/// OpenAPI content is
rejected and does not create an API-definition object. Defaults to \`\[\]\`. Server applies default
when omitted.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "512",
    "ves.io.schema.rules.repeated.items.string.pattern": "/api/object_store/namespaces/([a-z]([-a-z0-9]*[a-z0-9])?)/stored_objects/swagger/([a-z]([-a-z0-9]*[a-z0-9])?)/(v|V)[0-9]+(-[0-9]{2}){3}$",
    "ves.io.schema.rules.repeated.max_items": "20",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/#schema-annotations) |
| `api_inventory_exclusion_list` | [api_inventory_exclusion_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/api_inventory_exclusion_list/#section) |
| `api_inventory_exclusion_list.method` | [api_inventory_exclusion_list.method](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/api_inventory_exclusion_list/#schema-api_inventory_exclusion_list--method) |
| `api_inventory_exclusion_list.path` | [api_inventory_exclusion_list.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/api_inventory_exclusion_list/#schema-api_inventory_exclusion_list--path) |
| `api_inventory_inclusion_list` | [api_inventory_inclusion_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/api_inventory_inclusion_list/#section) |
| `api_inventory_inclusion_list.method` | [api_inventory_inclusion_list.method](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/api_inventory_inclusion_list/#schema-api_inventory_inclusion_list--method) |
| `api_inventory_inclusion_list.path` | [api_inventory_inclusion_list.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/api_inventory_inclusion_list/#schema-api_inventory_inclusion_list--path) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/#schema-labels) |
| `mixed_schema_origin` | [mixed_schema_origin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/mixed_schema_origin/#section) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/#schema-namespace) |
| `non_api_endpoints` | [non_api_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/non_api_endpoints/#section) |
| `non_api_endpoints.method` | [non_api_endpoints.method](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/non_api_endpoints/#schema-non_api_endpoints--method) |
| `non_api_endpoints.path` | [non_api_endpoints.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/non_api_endpoints/#schema-non_api_endpoints--path) |
| `strict_schema_origin` | [strict_schema_origin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/strict_schema_origin/#section) |
| `swagger_specs` | [swagger_specs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/#schema-swagger_specs) |
