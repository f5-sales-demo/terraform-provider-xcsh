---
page_title: "access_info.tls_config.common_params.tls_certificates"
subcategory: ""
description: "Set of TLS certificates."
xcsh_docs: {"aliases": ["access info tls config common params tls certificates", "cert", "certificate", "existing certificates", "tls certificates"], "body_bytes": 5232, "body_sha256": "sha256:ccf8afb7e4b911af5a9604b578a5179d9b557779e93e43e2a3f97f00d6d96e87", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:custom_hash_algorithms", "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:disable_ocsp_stapling", "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key", "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:use_system_defaults"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params", "path": "documentation/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312", "registry_path": "docs/guides/data-sources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "certificate url", "existing certificates", "tls certificates"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--certificate_url", "description": "TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom hash algorithms"], "anchor": "section", "description": "Specifies the hash algorithms to be used.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:custom_hash_algorithms", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "custom_hash_algorithms"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "description spec", "existing certificates", "tls certificates"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--description_spec", "description": "Description. Description for the certificate.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable ocsp stapling"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:disable_ocsp_stapling", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "disable_ocsp_stapling"], "syntax": "attribute", "type": "object"}, {"aliases": ["private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "private_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["use system defaults"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:use_system_defaults", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "use_system_defaults"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Set of TLS certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.common_params.tls_certificates

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/)
- [access_info.tls_config.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/)
- access_info.tls_config.common_params.tls_certificates

<a id="section"></a>

Type: `"list"`. Computed.

TLS Certificates. Set of TLS certificates.

Upstream description:

Set of TLS certificates.

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

<a id="schema-access_info--tls_config--common_params--tls_certificates--certificate_url"></a>

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

- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/custom_hash_algorithms/): complete subsection reference.

<a id="schema-access_info--tls_config--common_params--tls_certificates--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/disable_ocsp_stapling/): complete subsection reference.

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/): complete subsection reference.

- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/use_system_defaults/): complete subsection reference.

## Next pages

- [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/custom_hash_algorithms/)
- [access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/disable_ocsp_stapling/)
- [access_info.tls_config.common_params.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/)
- [access_info.tls_config.common_params.tls_certificates.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/use_system_defaults/)
- [access_info.tls_config.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
