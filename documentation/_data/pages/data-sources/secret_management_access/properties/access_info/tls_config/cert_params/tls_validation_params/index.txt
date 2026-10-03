---
page_title: "access_info.tls_config.cert_params.tls_validation_params"
subcategory: ""
description: "This includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names for verification."
xcsh_docs: {"aliases": ["access info tls config cert params tls validation params"], "body_bytes": 4947, "body_sha256": "sha256:61ab1426ac3057b6b3b16bac961f6b27f19d4bd107fd4ababdf9021725fda732", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params:trusted_ca"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params", "path": "documentation/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0212121101200211-3311122021003102-3323022130010010-1220003310202223-2133133021200212-3203110132212111-0110011212300112-1130202332133031", "registry_path": "docs/guides/data-sources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params"], "schema_version": 1, "sections": [{"aliases": ["access info tls config cert params tls validation params skip hostname verification"], "anchor": "schema-access_info--tls_config--cert_params--tls_validation_params--skip_hostname_verification", "description": "When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to the connecting hostname.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params", "skip_hostname_verification"], "syntax": "attribute", "type": "bool"}, {"aliases": ["access info tls config cert params tls validation params trusted ca"], "anchor": "section", "description": "Reference to Root CA Certificate.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params:trusted_ca", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params", "trusted_ca"], "syntax": "attribute", "type": "object"}, {"aliases": ["access info tls config cert params tls validation params trusted ca url"], "anchor": "schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca_url", "description": "Exclusive with Inline Root CA Certificate.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params", "trusted_ca_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config cert params tls validation params verify subject alt names"], "anchor": "schema-access_info--tls_config--cert_params--tls_validation_params--verify_subject_alt_names", "description": "List of acceptable Subject Alt Names/CN in the peer's certificate. When skip_hostname_verification is false and verify_subject_alt_names is empty, the hostname of the peer will be used for matching against SAN/CN of peer's certificate.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params", "verify_subject_alt_names"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names for verification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.cert_params.tls_validation_params

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/)
- [access_info.tls_config.cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/)
- access_info.tls_config.cert_params.tls_validation_params

<a id="section"></a>

Type: `"single"`. Computed.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

## Direct properties

<a id="schema-access_info--tls_config--cert_params--tls_validation_params--skip_hostname_verification"></a>

### skip_hostname_verification property

Type: `"bool"`. Computed.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Upstream description:

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

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

- [trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/): complete subsection reference.

<a id="schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="schema-access_info--tls_config--cert_params--tls_validation_params--verify_subject_alt_names"></a>

### verify_subject_alt_names property

Type: `["list", "string"]`. Computed.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Upstream description:

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

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

## Next pages

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/)
- [access_info.tls_config.cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
