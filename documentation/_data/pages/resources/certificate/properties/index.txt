---
page_title: "Property reference"
subcategory: "Security"
description: "Property reference for xcsh_certificate."
xcsh_docs: {"aliases": ["certificate"], "body_bytes": 15462, "body_sha256": "sha256:6f2cfbfd50bf5a9e83693ff06a00158cdcfe40ad155a9d5fbd13d3681347d5c9", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:certificate:properties:certificate_chain", "xcsh-docs:resources:certificate:properties:custom_hash_algorithms", "xcsh-docs:resources:certificate:properties:disable_ocsp_stapling", "xcsh-docs:resources:certificate:properties:private_key", "xcsh-docs:resources:certificate:properties:timeouts", "xcsh-docs:resources:certificate:properties:use_system_defaults"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate:reference", "parent_id": "xcsh-docs:resources:certificate:fundamentals", "path": "documentation/resources/certificate/properties/index.md", "product": "distributed-cloud", "provider_name": "certificate", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330", "registry_path": "docs/guides/resources--certificate--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:certificate:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["cert", "certificate", "certificate chain", "existing certificates", "tls certificates"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:certificate:properties:certificate_chain", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-certificate_chain--name", "enforcement": "provider-schema", "group": "certificate_chain:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:certificate:properties:certificate_chain", "type": "requires"}], "schema_path": ["certificate_chain"], "syntax": "block", "type": "object"}, {"aliases": ["cert", "certificate", "certificate url", "existing certificates", "tls certificates"], "anchor": "schema-certificate_url", "description": "Certificate. Certificate or certificate chain in PEM format including the PEM headers.", "document_id": "xcsh-docs:resources:certificate:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom hash algorithms"], "anchor": "section", "description": "Specifies the hash algorithms to be used.", "document_id": "xcsh-docs:resources:certificate:properties:custom_hash_algorithms", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_hash_algorithms--hash_algorithms", "enforcement": "provider-schema", "group": "custom_hash_algorithms:RequiredObjectAttributes:hash_algorithms", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:certificate:properties:custom_hash_algorithms", "type": "requires"}], "schema_path": ["custom_hash_algorithms"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:certificate:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:certificate:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["disable ocsp stapling"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:certificate:properties:disable_ocsp_stapling", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable_ocsp_stapling"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:certificate:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:certificate:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:certificate:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:certificate:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:certificate:properties:private_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:certificate:properties:private_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:certificate:properties:private_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["private_key"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:certificate:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["use system defaults"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:certificate:properties:use_system_defaults", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_system_defaults"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["certificateCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/)
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

- [certificate_chain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/certificate_chain/): complete subsection reference.

<a id="schema-certificate_url"></a>

### certificate_url property

Type: `"string"`. Required.

Certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

Certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/custom_hash_algorithms/): complete subsection reference.

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

- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/disable_ocsp_stapling/): complete subsection reference.

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

Name of the Certificate. Must be unique within the namespace.

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

Namespace where the Certificate is created.

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

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/private_key/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/timeouts/): complete subsection reference.

- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/use_system_defaults/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/#schema-annotations) |
| `certificate_chain` | [certificate_chain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/certificate_chain/#section) |
| `certificate_chain.name` | [certificate_chain.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/certificate_chain/#schema-certificate_chain--name) |
| `certificate_chain.namespace` | [certificate_chain.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/certificate_chain/#schema-certificate_chain--namespace) |
| `certificate_chain.tenant` | [certificate_chain.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/certificate_chain/#schema-certificate_chain--tenant) |
| `certificate_url` | [certificate_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/#schema-certificate_url) |
| `custom_hash_algorithms` | [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/custom_hash_algorithms/#section) |
| `custom_hash_algorithms.hash_algorithms` | [custom_hash_algorithms.hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/custom_hash_algorithms/#schema-custom_hash_algorithms--hash_algorithms) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/#schema-disable) |
| `disable_ocsp_stapling` | [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/disable_ocsp_stapling/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/#schema-namespace) |
| `private_key` | [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/private_key/#section) |
| `private_key.blindfold_secret_info` | [private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/private_key/blindfold_secret_info/#section) |
| `private_key.blindfold_secret_info.decryption_provider` | [private_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/private_key/blindfold_secret_info/#schema-private_key--blindfold_secret_info--decryption_provider) |
| `private_key.blindfold_secret_info.location` | [private_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/private_key/blindfold_secret_info/#schema-private_key--blindfold_secret_info--location) |
| `private_key.blindfold_secret_info.store_provider` | [private_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/private_key/blindfold_secret_info/#schema-private_key--blindfold_secret_info--store_provider) |
| `private_key.clear_secret_info` | [private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/private_key/clear_secret_info/#section) |
| `private_key.clear_secret_info.provider_ref` | [private_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/private_key/clear_secret_info/#schema-private_key--clear_secret_info--provider_ref) |
| `private_key.clear_secret_info.url` | [private_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/private_key/clear_secret_info/#schema-private_key--clear_secret_info--url) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/timeouts/#schema-timeouts--update) |
| `use_system_defaults` | [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/use_system_defaults/#section) |

## Next pages

- [certificate_chain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/certificate_chain/)
- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/custom_hash_algorithms/)
- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/disable_ocsp_stapling/)
- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/private_key/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/timeouts/)
- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/use_system_defaults/)
- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/)
