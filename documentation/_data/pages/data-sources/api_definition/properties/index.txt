---
page_title: "Property reference"
subcategory: "API Management"
description: "Property reference for xcsh_api_definition."
xcsh_docs: {"aliases": [], "body_bytes": 12486, "body_sha256": "sha256:940012e0b281493da99bc15efbb422c38863e52bd1d03bc58de08d6399028633", "child_ids": ["xcsh-docs:data-sources:api_definition:properties:api_inventory_exclusion_list", "xcsh-docs:data-sources:api_definition:properties:api_inventory_inclusion_list", "xcsh-docs:data-sources:api_definition:properties:mixed_schema_origin", "xcsh-docs:data-sources:api_definition:properties:non_api_endpoints", "xcsh-docs:data-sources:api_definition:properties:strict_schema_origin"], "collection_id": "xcsh-docs:data-sources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_definition:reference", "parent_id": "xcsh-docs:data-sources:api_definition:fundamentals", "path": "documentation/data-sources/api_definition/properties/index.md", "provider_name": "api_definition", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_definition/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_api_definition.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_definitionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/)
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

- [api_inventory_exclusion_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/api_inventory_exclusion_list/): complete subsection reference.

- [api_inventory_inclusion_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/api_inventory_inclusion_list/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the APIDefinition.

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

- [mixed_schema_origin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/mixed_schema_origin/): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the APIDefinition.

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

Namespace where the APIDefinition exists.

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

- [non_api_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/non_api_endpoints/): complete subsection reference.

- [strict_schema_origin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/strict_schema_origin/): complete subsection reference.

<a id="schema-swagger_specs"></a>

### swagger_specs property

Type: `["list", "string"]`. Computed.

URLs of versioned OpenAPI files uploaded through Web App &amp; API Protection &gt; Files &gt;
Swagger Files. The 512-byte item limit is a URL-length limit; inline string:/// OpenAPI content is
rejected and does not create an API-definition object. Defaults to \`\[\]\`. Server applies default
when omitted.

Upstream description:

URLs of versioned OpenAPI files uploaded through Web App &amp; API Protection &gt; Files &gt;
Swagger Files. The 512-byte item limit is a URL-length limit; inline string:/// OpenAPI content is
rejected and does not create an API-definition object.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [api_inventory_exclusion_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/api_inventory_exclusion_list/)
- [api_inventory_inclusion_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/api_inventory_inclusion_list/)
- [mixed_schema_origin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/mixed_schema_origin/)
- [non_api_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/non_api_endpoints/)
- [strict_schema_origin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/strict_schema_origin/)
- [xcsh_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/)
