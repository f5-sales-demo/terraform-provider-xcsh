---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_secret_management_access."
xcsh_docs: {"aliases": ["secret management access"], "body_bytes": 49493, "body_sha256": "sha256:a347faa7db70a38c359f7677fc0c141007baaa5cf7f22353fcd1f2c1b9f97f4b", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info", "xcsh-docs:data-sources:secret_management_access:properties:where"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:reference", "parent_id": "xcsh-docs:data-sources:secret_management_access:fundamentals", "path": "documentation/data-sources/secret_management_access/properties/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332", "registry_path": "docs/guides/data-sources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["access info"], "anchor": "section", "description": "HostAccessInfoType contains the information about how to connect to the remote host.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:secret_management_access:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:secret_management_access:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:secret_management_access:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:secret_management_access:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:secret_management_access:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:secret_management_access:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["backend servers", "origin servers", "provider name", "upstream servers"], "anchor": "schema-provider_name", "description": "Name given to this secret management backend. site.provider needs to be unique, and will be referenced for using this object.", "document_id": "xcsh-docs:data-sources:secret_management_access:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["provider_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["where"], "anchor": "section", "description": "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:where", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_secret_management_access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- Property reference

## Direct properties

- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/): complete subsection reference.

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

Description of the SecretManagementAccess.

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

Name of the SecretManagementAccess.

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

Namespace where the SecretManagementAccess exists.

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

<a id="schema-provider_name"></a>

### provider_name property

Type: `"string"`. Computed.

Name given to this secret management backend. site.provider needs to be unique, and will be
referenced for using this object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `access_info` | [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/#section) |
| `access_info.rest_auth_info` | [access_info.rest_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/#section) |
| `access_info.rest_auth_info.basic_auth` | [access_info.rest_auth_info.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/#section) |
| `access_info.rest_auth_info.basic_auth.password` | [access_info.rest_auth_info.basic_auth.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/#section) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/blindfold_secret_info/#section) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.decryption_provider` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/blindfold_secret_info/#schema-access_info--rest_auth_info--basic_auth--password--blindfold_secret_info--decryption_provider) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.location` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/blindfold_secret_info/#schema-access_info--rest_auth_info--basic_auth--password--blindfold_secret_info--location) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.store_provider` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/blindfold_secret_info/#schema-access_info--rest_auth_info--basic_auth--password--blindfold_secret_info--store_provider) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/clear_secret_info/#section) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info.provider_ref` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/clear_secret_info/#schema-access_info--rest_auth_info--basic_auth--password--clear_secret_info--provider_ref) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info.url` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/clear_secret_info/#schema-access_info--rest_auth_info--basic_auth--password--clear_secret_info--url) |
| `access_info.rest_auth_info.basic_auth.username` | [access_info.rest_auth_info.basic_auth.username](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/#schema-access_info--rest_auth_info--basic_auth--username) |
| `access_info.rest_auth_info.headers_auth` | [access_info.rest_auth_info.headers_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/headers_auth/#section) |
| `access_info.rest_auth_info.headers_auth.headers` | [access_info.rest_auth_info.headers_auth.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/headers_auth/headers/#section) |
| `access_info.rest_auth_info.query_params_auth` | [access_info.rest_auth_info.query_params_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/#section) |
| `access_info.rest_auth_info.query_params_auth.query_params` | [access_info.rest_auth_info.query_params_auth.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/query_params/#section) |
| `access_info.scheme` | [access_info.scheme](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/#schema-access_info--scheme) |
| `access_info.server_endpoint` | [access_info.server_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/#schema-access_info--server_endpoint) |
| `access_info.tls_config` | [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/#section) |
| `access_info.tls_config.cert_params` | [access_info.tls_config.cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/#section) |
| `access_info.tls_config.cert_params.certificates` | [access_info.tls_config.cert_params.certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/certificates/#section) |
| `access_info.tls_config.cert_params.certificates.kind` | [access_info.tls_config.cert_params.certificates.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/certificates/#schema-access_info--tls_config--cert_params--certificates--kind) |
| `access_info.tls_config.cert_params.certificates.name` | [access_info.tls_config.cert_params.certificates.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/certificates/#schema-access_info--tls_config--cert_params--certificates--name) |
| `access_info.tls_config.cert_params.certificates.namespace` | [access_info.tls_config.cert_params.certificates.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/certificates/#schema-access_info--tls_config--cert_params--certificates--namespace) |
| `access_info.tls_config.cert_params.certificates.tenant` | [access_info.tls_config.cert_params.certificates.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/certificates/#schema-access_info--tls_config--cert_params--certificates--tenant) |
| `access_info.tls_config.cert_params.certificates.uid` | [access_info.tls_config.cert_params.certificates.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/certificates/#schema-access_info--tls_config--cert_params--certificates--uid) |
| `access_info.tls_config.cert_params.cipher_suites` | [access_info.tls_config.cert_params.cipher_suites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/#schema-access_info--tls_config--cert_params--cipher_suites) |
| `access_info.tls_config.cert_params.maximum_protocol_version` | [access_info.tls_config.cert_params.maximum_protocol_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/#schema-access_info--tls_config--cert_params--maximum_protocol_version) |
| `access_info.tls_config.cert_params.minimum_protocol_version` | [access_info.tls_config.cert_params.minimum_protocol_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/#schema-access_info--tls_config--cert_params--minimum_protocol_version) |
| `access_info.tls_config.cert_params.skip_server_verification` | [access_info.tls_config.cert_params.skip_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/skip_server_verification/#section) |
| `access_info.tls_config.cert_params.tls_validation_params` | [access_info.tls_config.cert_params.tls_validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/#section) |
| `access_info.tls_config.cert_params.tls_validation_params.skip_hostname_verification` | [access_info.tls_config.cert_params.tls_validation_params.skip_hostname_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/#schema-access_info--tls_config--cert_params--tls_validation_params--skip_hostname_verification) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/#section) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/#section) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/#schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--kind) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/#schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--name) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/#schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--namespace) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/#schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--tenant) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/#schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--uid) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca_url` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/#schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca_url) |
| `access_info.tls_config.cert_params.tls_validation_params.verify_subject_alt_names` | [access_info.tls_config.cert_params.tls_validation_params.verify_subject_alt_names](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/#schema-access_info--tls_config--cert_params--tls_validation_params--verify_subject_alt_names) |
| `access_info.tls_config.cert_params.volterra_trusted_ca` | [access_info.tls_config.cert_params.volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/volterra_trusted_ca/#section) |
| `access_info.tls_config.common_params` | [access_info.tls_config.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/#section) |
| `access_info.tls_config.common_params.cipher_suites` | [access_info.tls_config.common_params.cipher_suites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/#schema-access_info--tls_config--common_params--cipher_suites) |
| `access_info.tls_config.common_params.maximum_protocol_version` | [access_info.tls_config.common_params.maximum_protocol_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/#schema-access_info--tls_config--common_params--maximum_protocol_version) |
| `access_info.tls_config.common_params.minimum_protocol_version` | [access_info.tls_config.common_params.minimum_protocol_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/#schema-access_info--tls_config--common_params--minimum_protocol_version) |
| `access_info.tls_config.common_params.tls_certificates` | [access_info.tls_config.common_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/#section) |
| `access_info.tls_config.common_params.tls_certificates.certificate_url` | [access_info.tls_config.common_params.tls_certificates.certificate_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/#schema-access_info--tls_config--common_params--tls_certificates--certificate_url) |
| `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms` | [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/custom_hash_algorithms/#section) |
| `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/custom_hash_algorithms/#schema-access_info--tls_config--common_params--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `access_info.tls_config.common_params.tls_certificates.description_spec` | [access_info.tls_config.common_params.tls_certificates.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/#schema-access_info--tls_config--common_params--tls_certificates--description_spec) |
| `access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling` | [access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/disable_ocsp_stapling/#section) |
| `access_info.tls_config.common_params.tls_certificates.private_key` | [access_info.tls_config.common_params.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/#section) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/blindfold_secret_info/#section) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/blindfold_secret_info/#schema-access_info--tls_config--common_params--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/blindfold_secret_info/#schema-access_info--tls_config--common_params--tls_certificates--private_key--blindfold_secret_info--location) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/blindfold_secret_info/#schema-access_info--tls_config--common_params--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/clear_secret_info/#section) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/clear_secret_info/#schema-access_info--tls_config--common_params--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.url` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/clear_secret_info/#schema-access_info--tls_config--common_params--tls_certificates--private_key--clear_secret_info--url) |
| `access_info.tls_config.common_params.tls_certificates.use_system_defaults` | [access_info.tls_config.common_params.tls_certificates.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/use_system_defaults/#section) |
| `access_info.tls_config.common_params.validation_params` | [access_info.tls_config.common_params.validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/#section) |
| `access_info.tls_config.common_params.validation_params.skip_hostname_verification` | [access_info.tls_config.common_params.validation_params.skip_hostname_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/#schema-access_info--tls_config--common_params--validation_params--skip_hostname_verification) |
| `access_info.tls_config.common_params.validation_params.trusted_ca` | [access_info.tls_config.common_params.validation_params.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/trusted_ca/#section) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/trusted_ca/trusted_ca_list/#section) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--kind) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--name) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--namespace) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--tenant) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-access_info--tls_config--common_params--validation_params--trusted_ca--trusted_ca_list--uid) |
| `access_info.tls_config.common_params.validation_params.trusted_ca_url` | [access_info.tls_config.common_params.validation_params.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/#schema-access_info--tls_config--common_params--validation_params--trusted_ca_url) |
| `access_info.tls_config.common_params.validation_params.verify_subject_alt_names` | [access_info.tls_config.common_params.validation_params.verify_subject_alt_names](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/#schema-access_info--tls_config--common_params--validation_params--verify_subject_alt_names) |
| `access_info.tls_config.default_session_key_caching` | [access_info.tls_config.default_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/default_session_key_caching/#section) |
| `access_info.tls_config.disable_session_key_caching` | [access_info.tls_config.disable_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/disable_session_key_caching/#section) |
| `access_info.tls_config.disable_sni` | [access_info.tls_config.disable_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/disable_sni/#section) |
| `access_info.tls_config.max_session_keys` | [access_info.tls_config.max_session_keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/#schema-access_info--tls_config--max_session_keys) |
| `access_info.tls_config.sni` | [access_info.tls_config.sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/#schema-access_info--tls_config--sni) |
| `access_info.tls_config.use_host_header_as_sni` | [access_info.tls_config.use_host_header_as_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/use_host_header_as_sni/#section) |
| `access_info.vault_auth_info` | [access_info.vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/#section) |
| `access_info.vault_auth_info.app_role_auth` | [access_info.vault_auth_info.app_role_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/#section) |
| `access_info.vault_auth_info.app_role_auth.role_id` | [access_info.vault_auth_info.app_role_auth.role_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/#schema-access_info--vault_auth_info--app_role_auth--role_id) |
| `access_info.vault_auth_info.app_role_auth.secret_id` | [access_info.vault_auth_info.app_role_auth.secret_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/#section) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/blindfold_secret_info/#section) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/blindfold_secret_info/#schema-access_info--vault_auth_info--app_role_auth--secret_id--blindfold_secret_info--decryption_provider) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/blindfold_secret_info/#schema-access_info--vault_auth_info--app_role_auth--secret_id--blindfold_secret_info--location) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/blindfold_secret_info/#schema-access_info--vault_auth_info--app_role_auth--secret_id--blindfold_secret_info--store_provider) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/clear_secret_info/#section) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/clear_secret_info/#schema-access_info--vault_auth_info--app_role_auth--secret_id--clear_secret_info--provider_ref) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/clear_secret_info/#schema-access_info--vault_auth_info--app_role_auth--secret_id--clear_secret_info--url) |
| `access_info.vault_auth_info.token` | [access_info.vault_auth_info.token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/#section) |
| `access_info.vault_auth_info.token.blindfold_secret_info` | [access_info.vault_auth_info.token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/blindfold_secret_info/#section) |
| `access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider` | [access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/blindfold_secret_info/#schema-access_info--vault_auth_info--token--blindfold_secret_info--decryption_provider) |
| `access_info.vault_auth_info.token.blindfold_secret_info.location` | [access_info.vault_auth_info.token.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/blindfold_secret_info/#schema-access_info--vault_auth_info--token--blindfold_secret_info--location) |
| `access_info.vault_auth_info.token.blindfold_secret_info.store_provider` | [access_info.vault_auth_info.token.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/blindfold_secret_info/#schema-access_info--vault_auth_info--token--blindfold_secret_info--store_provider) |
| `access_info.vault_auth_info.token.clear_secret_info` | [access_info.vault_auth_info.token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/clear_secret_info/#section) |
| `access_info.vault_auth_info.token.clear_secret_info.provider_ref` | [access_info.vault_auth_info.token.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/clear_secret_info/#schema-access_info--vault_auth_info--token--clear_secret_info--provider_ref) |
| `access_info.vault_auth_info.token.clear_secret_info.url` | [access_info.vault_auth_info.token.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/clear_secret_info/#schema-access_info--vault_auth_info--token--clear_secret_info--url) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/#schema-namespace) |
| `provider_name` | [provider_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/#schema-provider_name) |
| `where` | [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/#section) |
| `where.site` | [where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/site/#section) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/site/disable_internet_vip/#section) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/site/enable_internet_vip/#section) |
| `where.site.network_type` | [where.site.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/site/#schema-where--site--network_type) |
| `where.site.ref` | [where.site.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/site/ref/#section) |
| `where.site.ref.kind` | [where.site.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/site/ref/#schema-where--site--ref--kind) |
| `where.site.ref.name` | [where.site.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/site/ref/#schema-where--site--ref--name) |
| `where.site.ref.namespace` | [where.site.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/site/ref/#schema-where--site--ref--namespace) |
| `where.site.ref.tenant` | [where.site.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/site/ref/#schema-where--site--ref--tenant) |
| `where.site.ref.uid` | [where.site.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/site/ref/#schema-where--site--ref--uid) |
| `where.virtual_network` | [where.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_network/#section) |
| `where.virtual_network.ref` | [where.virtual_network.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_network/ref/#section) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--kind) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--name) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--namespace) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--tenant) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--uid) |
| `where.virtual_site` | [where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_site/#section) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_site/disable_internet_vip/#section) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_site/enable_internet_vip/#section) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_site/#schema-where--virtual_site--network_type) |
| `where.virtual_site.ref` | [where.virtual_site.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_site/ref/#section) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--kind) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--name) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--namespace) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--tenant) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--uid) |

## Next pages

- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/where/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
