---
page_title: "tls_intercept.custom_certificate"
subcategory: ""
description: "Handle to fetch certificate and key."
xcsh_docs: {"aliases": ["tls intercept custom certificate"], "body_bytes": 4368, "body_sha256": "sha256:2dab68738c4cb1fc22d3851f647877cc77b53d28fd28564d864d8df569684e96", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:custom_hash_algorithms", "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:disable_ocsp_stapling", "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:private_key", "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:use_system_defaults"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate", "parent_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept", "path": "documentation/data-sources/proxy/properties/tls_intercept/custom_certificate/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220", "registry_path": "docs/guides/data-sources--proxy--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_intercept", "custom_certificate"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "certificate url", "existing certificates", "tls certificates"], "anchor": "schema-tls_intercept--custom_certificate--certificate_url", "description": "TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.", "document_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "custom_certificate", "certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom hash algorithms"], "anchor": "section", "description": "Specifies the hash algorithms to be used.", "document_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:custom_hash_algorithms", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_intercept", "custom_certificate", "custom_hash_algorithms"], "syntax": "attribute", "type": "object"}, {"aliases": ["description spec"], "anchor": "schema-tls_intercept--custom_certificate--description_spec", "description": "Description. Description for the certificate.", "document_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "custom_certificate", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable ocsp stapling"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:disable_ocsp_stapling", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "custom_certificate", "disable_ocsp_stapling"], "syntax": "attribute", "type": "object"}, {"aliases": ["private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:private_key", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_intercept", "custom_certificate", "private_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["use system defaults"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:use_system_defaults", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "custom_certificate", "use_system_defaults"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/tls_intercept/custom_certificate/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Handle to fetch certificate and key.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.custom_certificate

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/)
- tls_intercept.custom_certificate

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for custom certificate.

Upstream description:

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

## Direct properties

<a id="schema-tls_intercept--custom_certificate--certificate_url"></a>

### certificate_url property

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

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

- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/custom_certificate/custom_hash_algorithms/): complete subsection reference.

<a id="schema-tls_intercept--custom_certificate--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/custom_certificate/disable_ocsp_stapling/): complete subsection reference.

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/custom_certificate/private_key/): complete subsection reference.

- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/custom_certificate/use_system_defaults/): complete subsection reference.

## Next pages

- [tls_intercept.custom_certificate.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/custom_certificate/custom_hash_algorithms/)
- [tls_intercept.custom_certificate.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/custom_certificate/disable_ocsp_stapling/)
- [tls_intercept.custom_certificate.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/custom_certificate/private_key/)
- [tls_intercept.custom_certificate.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/custom_certificate/use_system_defaults/)
- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
