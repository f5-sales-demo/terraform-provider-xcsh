---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_user_identification."
xcsh_docs: {"aliases": [], "body_bytes": 8744, "body_sha256": "sha256:5a3674fa197586e90c58ca8d8438fe672524a7a61a233690a3d3c0e5289ad7ee", "canonical_id": "xcsh-docs:data-sources:user_identification:reference", "child_ids": ["xcsh-docs:data-sources:user_identification:properties:rules"], "collection_id": "xcsh-docs:data-sources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:user_identification:reference", "parent_id": "xcsh-docs:data-sources:user_identification:fundamentals", "path": "docs/guides/data-sources--user_identification--reference.md", "provider_name": "user_identification", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/user_identification/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_user_identification.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md)
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

Description of the UserIdentification.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the UserIdentification.

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

Namespace where the UserIdentification exists.

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

- [rules](data-sources--user_identification--properties--rules.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--user_identification--reference.md#schema-annotations) |
| `description` | [description](data-sources--user_identification--reference.md#schema-description) |
| `id` | [id](data-sources--user_identification--reference.md#schema-id) |
| `labels` | [labels](data-sources--user_identification--reference.md#schema-labels) |
| `name` | [name](data-sources--user_identification--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--user_identification--reference.md#schema-namespace) |
| `rules` | [rules](data-sources--user_identification--properties--rules.md#section) |
| `rules.client_asn` | [rules.client_asn](data-sources--user_identification--properties--rules--client_asn.md#section) |
| `rules.client_city` | [rules.client_city](data-sources--user_identification--properties--rules--client_city.md#section) |
| `rules.client_country` | [rules.client_country](data-sources--user_identification--properties--rules--client_country.md#section) |
| `rules.client_ip` | [rules.client_ip](data-sources--user_identification--properties--rules--client_ip.md#section) |
| `rules.client_region` | [rules.client_region](data-sources--user_identification--properties--rules--client_region.md#section) |
| `rules.cookie_name` | [rules.cookie_name](data-sources--user_identification--properties--rules.md#schema-rules--cookie_name) |
| `rules.http_header_name` | [rules.http_header_name](data-sources--user_identification--properties--rules.md#schema-rules--http_header_name) |
| `rules.ip_and_http_header_name` | [rules.ip_and_http_header_name](data-sources--user_identification--properties--rules.md#schema-rules--ip_and_http_header_name) |
| `rules.ip_and_ja4_tls_fingerprint` | [rules.ip_and_ja4_tls_fingerprint](data-sources--user_identification--properties--rules--ip_and_ja4_tls_fingerprint.md#section) |
| `rules.ip_and_tls_fingerprint` | [rules.ip_and_tls_fingerprint](data-sources--user_identification--properties--rules--ip_and_tls_fingerprint.md#section) |
| `rules.ja4_tls_fingerprint` | [rules.ja4_tls_fingerprint](data-sources--user_identification--properties--rules--ja4_tls_fingerprint.md#section) |
| `rules.jwt_claim_name` | [rules.jwt_claim_name](data-sources--user_identification--properties--rules.md#schema-rules--jwt_claim_name) |
| `rules.none` | [rules.none](data-sources--user_identification--properties--rules--none.md#section) |
| `rules.query_param_key` | [rules.query_param_key](data-sources--user_identification--properties--rules.md#schema-rules--query_param_key) |
| `rules.tls_fingerprint` | [rules.tls_fingerprint](data-sources--user_identification--properties--rules--tls_fingerprint.md#section) |

## Next pages

- [rules](data-sources--user_identification--properties--rules.md)
- [xcsh_user_identification](../data-sources/user_identification.md)
