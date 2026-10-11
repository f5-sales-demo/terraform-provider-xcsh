---
page_title: "tls_intercept.custom_certificate"
subcategory: ""
description: "Handle to fetch certificate and key."
xcsh_docs: {"aliases": ["tls intercept custom certificate"], "body_bytes": 3537, "body_sha256": "sha256:d31f9b6aa8b10c23445a36508089d4228e9794dc0754d3b245eaaed88be49820", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:blindfold", "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:custom_hash_algorithms", "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:disable_ocsp_stapling", "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:private_key", "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:use_system_defaults"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate", "parent_id": "xcsh-docs:resources:proxy:properties:tls_intercept", "path": "documentation/resources/proxy/properties/tls_intercept/custom_certificate/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223", "registry_path": "docs/guides/resources--proxy--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_intercept", "custom_certificate"], "schema_version": 1, "sections": [{"aliases": ["tls intercept custom certificate blindfold"], "anchor": "section", "description": "Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with material_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates require unique IDs. Private inputs are never stored.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:blindfold", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_intercept", "custom_certificate", "blindfold"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "tls intercept custom certificate certificate url"], "anchor": "schema-tls_intercept--custom_certificate--certificate_url", "description": "TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "custom_certificate", "certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls intercept custom certificate custom hash algorithms"], "anchor": "section", "description": "Specifies the hash algorithms to be used.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:custom_hash_algorithms", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_intercept", "custom_certificate", "custom_hash_algorithms"], "syntax": "block", "type": "object"}, {"aliases": ["tls intercept custom certificate description spec"], "anchor": "schema-tls_intercept--custom_certificate--description_spec", "description": "Description. Description for the certificate.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "custom_certificate", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls intercept custom certificate disable ocsp stapling"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:disable_ocsp_stapling", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "custom_certificate", "disable_ocsp_stapling"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls intercept custom certificate private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:private_key", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_intercept", "custom_certificate", "private_key"], "syntax": "block", "type": "object"}, {"aliases": ["tls intercept custom certificate use system defaults"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:use_system_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "custom_certificate", "use_system_defaults"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/tls_intercept/custom_certificate/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Handle to fetch certificate and key.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.custom_certificate

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/)
- tls_intercept.custom_certificate

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for custom certificate.

Additional upstream details:

Handle to fetch certificate and key.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ocsp_stapling_choice": "[\"custom_hash_algorithms\",\"disable_ocsp_stapling\",\"use_system_defaults\"]"
}
```

Terraform syntax:

```terraform
custom_certificate {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/custom_certificate/blindfold/): complete subsection reference.

<a id="schema-tls_intercept--custom_certificate--certificate_url"></a>

### certificate_url property

Type: `"string"`. Optional, Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/custom_certificate/custom_hash_algorithms/): complete subsection reference.

<a id="schema-tls_intercept--custom_certificate--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/custom_certificate/disable_ocsp_stapling/): complete subsection reference.

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/custom_certificate/private_key/): complete subsection reference.

- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/custom_certificate/use_system_defaults/): complete subsection reference.
