---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_authentication."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials"], "body_bytes": 18335, "body_sha256": "sha256:994275472cf71c4cd3336335a05dd6060da80918b7f239d9a245103fbc093b3b", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:authentication:properties:cookie_params", "xcsh-docs:data-sources:authentication:properties:oidc_auth"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:authentication:reference", "parent_id": "xcsh-docs:data-sources:authentication:fundamentals", "path": "documentation/data-sources/authentication/properties/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210", "registry_path": "docs/guides/data-sources--authentication--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:authentication:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["cookie params"], "anchor": "section", "description": "Specifies different cookie related config parameters for authentication.", "document_id": "xcsh-docs:data-sources:authentication:properties:cookie_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cookie_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:authentication:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:authentication:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:authentication:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:authentication:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:authentication:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["oidc auth"], "anchor": "section", "description": "OIDCAuthType.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["oidc_auth"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authentication/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_authentication.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["authenticationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/)
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

- [cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the Authentication.

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

Name of the Authentication.

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

Namespace where the Authentication exists.

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

- [oidc_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/#schema-annotations) |
| `cookie_params` | [cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/#section) |
| `cookie_params.auth_hmac` | [cookie_params.auth_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/#section) |
| `cookie_params.auth_hmac.prim_key` | [cookie_params.auth_hmac.prim_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/prim_key/#section) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/prim_key/blindfold_secret_info/#section) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/prim_key/blindfold_secret_info/#schema-cookie_params--auth_hmac--prim_key--blindfold_secret_info--decryption_provider) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.location` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/prim_key/blindfold_secret_info/#schema-cookie_params--auth_hmac--prim_key--blindfold_secret_info--location) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/prim_key/blindfold_secret_info/#schema-cookie_params--auth_hmac--prim_key--blindfold_secret_info--store_provider) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info` | [cookie_params.auth_hmac.prim_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/prim_key/clear_secret_info/#section) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref` | [cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/prim_key/clear_secret_info/#schema-cookie_params--auth_hmac--prim_key--clear_secret_info--provider_ref) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info.url` | [cookie_params.auth_hmac.prim_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/prim_key/clear_secret_info/#schema-cookie_params--auth_hmac--prim_key--clear_secret_info--url) |
| `cookie_params.auth_hmac.prim_key_expiry` | [cookie_params.auth_hmac.prim_key_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/#schema-cookie_params--auth_hmac--prim_key_expiry) |
| `cookie_params.auth_hmac.sec_key` | [cookie_params.auth_hmac.sec_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/sec_key/#section) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/sec_key/blindfold_secret_info/#section) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/sec_key/blindfold_secret_info/#schema-cookie_params--auth_hmac--sec_key--blindfold_secret_info--decryption_provider) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.location` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/sec_key/blindfold_secret_info/#schema-cookie_params--auth_hmac--sec_key--blindfold_secret_info--location) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/sec_key/blindfold_secret_info/#schema-cookie_params--auth_hmac--sec_key--blindfold_secret_info--store_provider) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info` | [cookie_params.auth_hmac.sec_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/sec_key/clear_secret_info/#section) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref` | [cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/sec_key/clear_secret_info/#schema-cookie_params--auth_hmac--sec_key--clear_secret_info--provider_ref) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info.url` | [cookie_params.auth_hmac.sec_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/sec_key/clear_secret_info/#schema-cookie_params--auth_hmac--sec_key--clear_secret_info--url) |
| `cookie_params.auth_hmac.sec_key_expiry` | [cookie_params.auth_hmac.sec_key_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/auth_hmac/#schema-cookie_params--auth_hmac--sec_key_expiry) |
| `cookie_params.cookie_expiry` | [cookie_params.cookie_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/#schema-cookie_params--cookie_expiry) |
| `cookie_params.cookie_refresh_interval` | [cookie_params.cookie_refresh_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/#schema-cookie_params--cookie_refresh_interval) |
| `cookie_params.kms_key_hmac` | [cookie_params.kms_key_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/kms_key_hmac/#section) |
| `cookie_params.session_expiry` | [cookie_params.session_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/#schema-cookie_params--session_expiry) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/#schema-namespace) |
| `oidc_auth` | [oidc_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/#section) |
| `oidc_auth.client_secret` | [oidc_auth.client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/client_secret/#section) |
| `oidc_auth.client_secret.blindfold_secret_info` | [oidc_auth.client_secret.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/client_secret/blindfold_secret_info/#section) |
| `oidc_auth.client_secret.blindfold_secret_info.decryption_provider` | [oidc_auth.client_secret.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/client_secret/blindfold_secret_info/#schema-oidc_auth--client_secret--blindfold_secret_info--decryption_provider) |
| `oidc_auth.client_secret.blindfold_secret_info.location` | [oidc_auth.client_secret.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/client_secret/blindfold_secret_info/#schema-oidc_auth--client_secret--blindfold_secret_info--location) |
| `oidc_auth.client_secret.blindfold_secret_info.store_provider` | [oidc_auth.client_secret.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/client_secret/blindfold_secret_info/#schema-oidc_auth--client_secret--blindfold_secret_info--store_provider) |
| `oidc_auth.client_secret.clear_secret_info` | [oidc_auth.client_secret.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/client_secret/clear_secret_info/#section) |
| `oidc_auth.client_secret.clear_secret_info.provider_ref` | [oidc_auth.client_secret.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/client_secret/clear_secret_info/#schema-oidc_auth--client_secret--clear_secret_info--provider_ref) |
| `oidc_auth.client_secret.clear_secret_info.url` | [oidc_auth.client_secret.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/client_secret/clear_secret_info/#schema-oidc_auth--client_secret--clear_secret_info--url) |
| `oidc_auth.oidc_auth_params` | [oidc_auth.oidc_auth_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/oidc_auth_params/#section) |
| `oidc_auth.oidc_auth_params.auth_endpoint_url` | [oidc_auth.oidc_auth_params.auth_endpoint_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/oidc_auth_params/#schema-oidc_auth--oidc_auth_params--auth_endpoint_url) |
| `oidc_auth.oidc_auth_params.end_session_endpoint_url` | [oidc_auth.oidc_auth_params.end_session_endpoint_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/oidc_auth_params/#schema-oidc_auth--oidc_auth_params--end_session_endpoint_url) |
| `oidc_auth.oidc_auth_params.token_endpoint_url` | [oidc_auth.oidc_auth_params.token_endpoint_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/oidc_auth_params/#schema-oidc_auth--oidc_auth_params--token_endpoint_url) |
| `oidc_auth.oidc_client_id` | [oidc_auth.oidc_client_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/#schema-oidc_auth--oidc_client_id) |
| `oidc_auth.oidc_well_known_config_url` | [oidc_auth.oidc_well_known_config_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/#schema-oidc_auth--oidc_well_known_config_url) |

## Next pages

- [cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/)
- [oidc_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/)
- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/)
