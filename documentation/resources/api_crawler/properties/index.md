---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_api_crawler."
xcsh_docs: {"aliases": ["api crawler"], "body_bytes": 12232, "body_sha256": "sha256:f9ef95643106ff9ee804eb72281fa08f411cccfc3206ef702def5d0b522a272a", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_crawler:properties:domains", "xcsh-docs:resources:api_crawler:properties:timeouts"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_crawler:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_crawler:reference", "parent_id": "xcsh-docs:resources:api_crawler:fundamentals", "path": "documentation/resources/api_crawler/properties/index.md", "product": "distributed-cloud", "provider_name": "api_crawler", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3010222002133020-2211321021131000-2301002232300112-1123303232312202-2000011002203003-2003112030230322-2103211310110013-1020232220123011", "registry_path": "docs/guides/resources--api_crawler--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:api_crawler:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:api_crawler:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:api_crawler:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["domains"], "anchor": "section", "description": "API Crawler Configuration.", "document_id": "xcsh-docs:resources:api_crawler:properties:domains", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-domains--domain", "enforcement": "provider-schema", "group": "domains:RequiredListObjectAttributes:domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:api_crawler:properties:domains", "type": "requires"}], "schema_path": ["domains"], "syntax": "block", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:api_crawler:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:api_crawler:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:api_crawler:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:api_crawler:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:api_crawler:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_crawler/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_api_crawler.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_crawlerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/)
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

- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/): complete subsection reference.

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

Name of the API Crawler. Must be unique within the namespace.

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

Namespace where the API Crawler is created.

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

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/#schema-disable) |
| `domains` | [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/#section) |
| `domains.domain` | [domains.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/#schema-domains--domain) |
| `domains.simple_login` | [domains.simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/#section) |
| `domains.simple_login.password` | [domains.simple_login.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/password/#section) |
| `domains.simple_login.password.blindfold_secret_info` | [domains.simple_login.password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/password/blindfold_secret_info/#section) |
| `domains.simple_login.password.blindfold_secret_info.decryption_provider` | [domains.simple_login.password.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/password/blindfold_secret_info/#schema-domains--simple_login--password--blindfold_secret_info--decryption_provider) |
| `domains.simple_login.password.blindfold_secret_info.location` | [domains.simple_login.password.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/password/blindfold_secret_info/#schema-domains--simple_login--password--blindfold_secret_info--location) |
| `domains.simple_login.password.blindfold_secret_info.store_provider` | [domains.simple_login.password.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/password/blindfold_secret_info/#schema-domains--simple_login--password--blindfold_secret_info--store_provider) |
| `domains.simple_login.password.clear_secret_info` | [domains.simple_login.password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/password/clear_secret_info/#section) |
| `domains.simple_login.password.clear_secret_info.provider_ref` | [domains.simple_login.password.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/password/clear_secret_info/#schema-domains--simple_login--password--clear_secret_info--provider_ref) |
| `domains.simple_login.password.clear_secret_info.url` | [domains.simple_login.password.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/password/clear_secret_info/#schema-domains--simple_login--password--clear_secret_info--url) |
| `domains.simple_login.user` | [domains.simple_login.user](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/simple_login/#schema-domains--simple_login--user) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/#schema-namespace) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/domains/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/properties/timeouts/)
- [xcsh_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_crawler/)
