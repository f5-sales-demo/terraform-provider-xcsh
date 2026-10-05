---
page_title: "use_tls.use_mtls.tls_certificates"
subcategory: "Load Balancing"
description: "MTLS Client Certificate."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "use tls use mtls tls certificates"], "body_bytes": 5152, "body_sha256": "sha256:9f345401487eea977ebf51b07f96ac20bc4deff583b65375f314961b49376dcc", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates:custom_hash_algorithms", "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates:disable_ocsp_stapling", "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key", "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates:use_system_defaults"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls", "path": "documentation/data-sources/origin_pool/properties/use_tls/use_mtls/tls_certificates/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1303112212030232-0223201010130110-2203010323313333-1333323310300102-0010003103112021-3212312311323231-0313030221021313-3302023303112222", "registry_path": "docs/guides/data-sources--origin_pool--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["use_tls", "use_mtls", "tls_certificates"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "use tls use mtls tls certificates certificate url"], "anchor": "schema-use_tls--use_mtls--tls_certificates--certificate_url", "description": "TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "use_mtls", "tls_certificates", "certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["use tls use mtls tls certificates custom hash algorithms"], "anchor": "section", "description": "Specifies the hash algorithms to be used.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates:custom_hash_algorithms", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["use_tls", "use_mtls", "tls_certificates", "custom_hash_algorithms"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls use mtls tls certificates description spec"], "anchor": "schema-use_tls--use_mtls--tls_certificates--description_spec", "description": "Description. Description for the certificate.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "use_mtls", "tls_certificates", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["use tls use mtls tls certificates disable ocsp stapling"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates:disable_ocsp_stapling", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "use_mtls", "tls_certificates", "disable_ocsp_stapling"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls use mtls tls certificates private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["use_tls", "use_mtls", "tls_certificates", "private_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls use mtls tls certificates use system defaults"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates:use_system_defaults", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "use_mtls", "tls_certificates", "use_system_defaults"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/use_tls/use_mtls/tls_certificates/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "MTLS Client Certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls.use_mtls.tls_certificates

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/)
- [use_tls.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_mtls/)
- use_tls.use_mtls.tls_certificates

<a id="section"></a>

Type: `"list"`. Computed.

MTLS Client Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

## Direct properties

<a id="schema-use_tls--use_mtls--tls_certificates--certificate_url"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_mtls/tls_certificates/custom_hash_algorithms/): complete subsection reference.

<a id="schema-use_tls--use_mtls--tls_certificates--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_mtls/tls_certificates/disable_ocsp_stapling/): complete subsection reference.

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/): complete subsection reference.

- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_mtls/tls_certificates/use_system_defaults/): complete subsection reference.

## Next pages

- [use_tls.use_mtls.tls_certificates.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_mtls/tls_certificates/custom_hash_algorithms/)
- [use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_mtls/tls_certificates/disable_ocsp_stapling/)
- [use_tls.use_mtls.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/)
- [use_tls.use_mtls.tls_certificates.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_mtls/tls_certificates/use_system_defaults/)
- [use_tls.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_mtls/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
