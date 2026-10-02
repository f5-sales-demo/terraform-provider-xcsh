---
page_title: "Property reference"
subcategory: "Security"
description: "Property reference for xcsh_certificate."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "tls certificates"], "body_bytes": 13417, "body_sha256": "sha256:9ec13df60d6d02439c2f0fe69924d1ac8425bf86ea8b92562981fe70404e002c", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:certificate:properties:certificate_chain", "xcsh-docs:data-sources:certificate:properties:custom_hash_algorithms", "xcsh-docs:data-sources:certificate:properties:disable_ocsp_stapling", "xcsh-docs:data-sources:certificate:properties:private_key", "xcsh-docs:data-sources:certificate:properties:use_system_defaults"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certificate:reference", "parent_id": "xcsh-docs:data-sources:certificate:fundamentals", "path": "documentation/data-sources/certificate/properties/index.md", "product": "distributed-cloud", "provider_name": "certificate", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000", "registry_path": "docs/guides/data-sources--certificate--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:certificate:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["cert", "certificate", "certificate chain", "existing certificates", "tls certificates"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:certificate:properties:certificate_chain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["certificate_chain"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "certificate url", "existing certificates", "tls certificates"], "anchor": "schema-certificate_url", "description": "Certificate. Certificate or certificate chain in PEM format including the PEM headers.", "document_id": "xcsh-docs:data-sources:certificate:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom hash algorithms"], "anchor": "section", "description": "Specifies the hash algorithms to be used.", "document_id": "xcsh-docs:data-sources:certificate:properties:custom_hash_algorithms", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_hash_algorithms"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:certificate:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable ocsp stapling"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:certificate:properties:disable_ocsp_stapling", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable_ocsp_stapling"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:certificate:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:certificate:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:certificate:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:certificate:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:certificate:properties:private_key", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["private_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["use system defaults"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:certificate:properties:use_system_defaults", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_system_defaults"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certificate/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["certificateCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/)
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

- [certificate_chain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/certificate_chain/): complete subsection reference.

<a id="schema-certificate_url"></a>

### certificate_url property

Type: `"string"`. Computed.

Certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

Certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/custom_hash_algorithms/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the Certificate.

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

- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/disable_ocsp_stapling/): complete subsection reference.

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

Name of the Certificate.

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

Namespace where the Certificate exists.

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

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/private_key/): complete subsection reference.

- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/use_system_defaults/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/#schema-annotations) |
| `certificate_chain` | [certificate_chain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/certificate_chain/#section) |
| `certificate_chain.name` | [certificate_chain.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/certificate_chain/#schema-certificate_chain--name) |
| `certificate_chain.namespace` | [certificate_chain.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/certificate_chain/#schema-certificate_chain--namespace) |
| `certificate_chain.tenant` | [certificate_chain.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/certificate_chain/#schema-certificate_chain--tenant) |
| `certificate_url` | [certificate_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/#schema-certificate_url) |
| `custom_hash_algorithms` | [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/custom_hash_algorithms/#section) |
| `custom_hash_algorithms.hash_algorithms` | [custom_hash_algorithms.hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/custom_hash_algorithms/#schema-custom_hash_algorithms--hash_algorithms) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/#schema-description) |
| `disable_ocsp_stapling` | [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/disable_ocsp_stapling/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/#schema-namespace) |
| `private_key` | [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/private_key/#section) |
| `private_key.blindfold_secret_info` | [private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/private_key/blindfold_secret_info/#section) |
| `private_key.blindfold_secret_info.decryption_provider` | [private_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/private_key/blindfold_secret_info/#schema-private_key--blindfold_secret_info--decryption_provider) |
| `private_key.blindfold_secret_info.location` | [private_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/private_key/blindfold_secret_info/#schema-private_key--blindfold_secret_info--location) |
| `private_key.blindfold_secret_info.store_provider` | [private_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/private_key/blindfold_secret_info/#schema-private_key--blindfold_secret_info--store_provider) |
| `private_key.clear_secret_info` | [private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/private_key/clear_secret_info/#section) |
| `private_key.clear_secret_info.provider_ref` | [private_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/private_key/clear_secret_info/#schema-private_key--clear_secret_info--provider_ref) |
| `private_key.clear_secret_info.url` | [private_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/private_key/clear_secret_info/#schema-private_key--clear_secret_info--url) |
| `use_system_defaults` | [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/use_system_defaults/#section) |

## Next pages

- [certificate_chain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/certificate_chain/)
- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/custom_hash_algorithms/)
- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/disable_ocsp_stapling/)
- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/private_key/)
- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/properties/use_system_defaults/)
- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/)
