---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cminstance."
xcsh_docs: {"aliases": ["cminstance"], "body_bytes": 13998, "body_sha256": "sha256:d08f486d9dc1710a1b4909aa5b8599dd8b26a62911fc09973221ead1680c6f60", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:cminstance:properties:api_token", "xcsh-docs:data-sources:cminstance:properties:ip", "xcsh-docs:data-sources:cminstance:properties:password"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cminstance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cminstance:reference", "parent_id": "xcsh-docs:data-sources:cminstance:fundamentals", "path": "documentation/data-sources/cminstance/properties/index.md", "product": "distributed-cloud", "provider_name": "cminstance", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3021233103221211-3000330212303311-3123011101210133-2233111111213112-2200323123322113-3330123132032201-0000230310101123-2230201321203202", "registry_path": "docs/guides/data-sources--cminstance--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:cminstance:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["api token", "authentication", "credential setup", "credentials"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:cminstance:properties:api_token", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_token"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:cminstance:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:cminstance:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["ip"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:data-sources:cminstance:properties:ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:cminstance:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:cminstance:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:cminstance:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:cminstance:properties:password", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["password"], "syntax": "attribute", "type": "object"}, {"aliases": ["port"], "anchor": "schema-port", "description": "Port of the Central Manager instance to connect to.", "document_id": "xcsh-docs:data-sources:cminstance:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["port"], "syntax": "attribute", "type": "number"}, {"aliases": ["username"], "anchor": "schema-username", "description": "Username for the Central Manager instance.", "document_id": "xcsh-docs:data-sources:cminstance:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["username"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cminstance/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_cminstance.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cminstanceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/)
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

- [api_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/api_token/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the Cminstance.

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

- [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/ip/): complete subsection reference.

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

Name of the Cminstance.

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

Namespace where the Cminstance exists.

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

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/password/): complete subsection reference.

<a id="schema-port"></a>

### port property

Type: `"number"`. Computed.

Port of the Central Manager instance to connect to.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-username"></a>

### username property

Type: `"string"`. Computed.

Username for the Central Manager instance.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "4"
  }
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/#schema-annotations) |
| `api_token` | [api_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/api_token/#section) |
| `api_token.blindfold_secret_info` | [api_token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/api_token/blindfold_secret_info/#section) |
| `api_token.blindfold_secret_info.decryption_provider` | [api_token.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/api_token/blindfold_secret_info/#schema-api_token--blindfold_secret_info--decryption_provider) |
| `api_token.blindfold_secret_info.location` | [api_token.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/api_token/blindfold_secret_info/#schema-api_token--blindfold_secret_info--location) |
| `api_token.blindfold_secret_info.store_provider` | [api_token.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/api_token/blindfold_secret_info/#schema-api_token--blindfold_secret_info--store_provider) |
| `api_token.clear_secret_info` | [api_token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/api_token/clear_secret_info/#section) |
| `api_token.clear_secret_info.provider_ref` | [api_token.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/api_token/clear_secret_info/#schema-api_token--clear_secret_info--provider_ref) |
| `api_token.clear_secret_info.url` | [api_token.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/api_token/clear_secret_info/#schema-api_token--clear_secret_info--url) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/#schema-id) |
| `ip` | [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/ip/#section) |
| `ip.addr` | [ip.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/ip/#schema-ip--addr) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/#schema-namespace) |
| `password` | [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/password/#section) |
| `password.blindfold_secret_info` | [password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/password/blindfold_secret_info/#section) |
| `password.blindfold_secret_info.decryption_provider` | [password.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/password/blindfold_secret_info/#schema-password--blindfold_secret_info--decryption_provider) |
| `password.blindfold_secret_info.location` | [password.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/password/blindfold_secret_info/#schema-password--blindfold_secret_info--location) |
| `password.blindfold_secret_info.store_provider` | [password.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/password/blindfold_secret_info/#schema-password--blindfold_secret_info--store_provider) |
| `password.clear_secret_info` | [password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/password/clear_secret_info/#section) |
| `password.clear_secret_info.provider_ref` | [password.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/password/clear_secret_info/#schema-password--clear_secret_info--provider_ref) |
| `password.clear_secret_info.url` | [password.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/password/clear_secret_info/#schema-password--clear_secret_info--url) |
| `port` | [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/#schema-port) |
| `username` | [username](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/#schema-username) |

## Next pages

- [api_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/api_token/)
- [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/ip/)
- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/password/)
- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/)
