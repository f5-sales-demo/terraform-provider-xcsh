---
page_title: "tls_parameters.common_params.tls_certificates"
subcategory: ""
description: "Set of TLS certificates."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "tls parameters common params tls certificates"], "body_bytes": 3366, "body_sha256": "sha256:ffbd938fbc2b417efeda8184716262762de23f0ff98c82af8d7fb1b7a73d865d", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates:custom_hash_algorithms", "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates:disable_ocsp_stapling", "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates:private_key", "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates:use_system_defaults"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params", "path": "documentation/data-sources/virtual_host/properties/tls_parameters/common_params/tls_certificates/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "tls parameters common params tls certificates certificate url"], "anchor": "schema-tls_parameters--common_params--tls_certificates--certificate_url", "description": "TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "common_params", "tls_certificates", "certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls parameters common params tls certificates custom hash algorithms"], "anchor": "section", "description": "Specifies the hash algorithms to be used.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates:custom_hash_algorithms", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_parameters", "common_params", "tls_certificates", "custom_hash_algorithms"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters common params tls certificates description spec"], "anchor": "schema-tls_parameters--common_params--tls_certificates--description_spec", "description": "Description. Description for the certificate.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "common_params", "tls_certificates", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls parameters common params tls certificates disable ocsp stapling"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates:disable_ocsp_stapling", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "common_params", "tls_certificates", "disable_ocsp_stapling"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters common params tls certificates private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates:private_key", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters common params tls certificates use system defaults"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates:use_system_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "common_params", "tls_certificates", "use_system_defaults"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/tls_parameters/common_params/tls_certificates/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Set of TLS certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.tls_certificates

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/)
- [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/common_params/)
- tls_parameters.common_params.tls_certificates

<a id="section"></a>

Type: `"list"`. Computed.

TLS Certificates. Set of TLS certificates.

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

## Direct properties

<a id="schema-tls_parameters--common_params--tls_certificates--certificate_url"></a>

### certificate_url property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/common_params/tls_certificates/custom_hash_algorithms/): complete subsection reference.

<a id="schema-tls_parameters--common_params--tls_certificates--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/common_params/tls_certificates/disable_ocsp_stapling/): complete subsection reference.

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/common_params/tls_certificates/private_key/): complete subsection reference.

- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/common_params/tls_certificates/use_system_defaults/): complete subsection reference.
