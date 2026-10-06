---
page_title: "access_info.tls_config.cert_params.tls_validation_params"
subcategory: ""
description: "This includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names for verification."
xcsh_docs: {"aliases": ["access info tls config cert params tls validation params"], "body_bytes": 4163, "body_sha256": "sha256:8cbfb7ff08089a2c619670abfd7d01ea11f770e085312ca437ba5f3799dd3048", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params:trusted_ca"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params", "path": "documentation/resources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0121031320112211-1123210202002222-2222232302211030-2030211033120231-2211100103310122-1110011313303011-1233131133202213-2132123300322200", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [{"anchor": "schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca_url", "enforcement": "provider-schema", "group": "access_info.tls_config.cert_params.tls_validation_params:ConflictingObjectAttributes:trusted_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config.cert_params.tls_validation_params:ConflictingObjectAttributes:trusted_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params:trusted_ca", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params"], "schema_version": 1, "sections": [{"aliases": ["access info tls config cert params tls validation params skip hostname verification"], "anchor": "schema-access_info--tls_config--cert_params--tls_validation_params--skip_hostname_verification", "description": "When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to the connecting hostname.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params", "skip_hostname_verification"], "syntax": "attribute", "type": "bool"}, {"aliases": ["access info tls config cert params tls validation params trusted ca"], "anchor": "section", "description": "Reference to Root CA Certificate.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params:trusted_ca", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params", "trusted_ca"], "syntax": "block", "type": "object"}, {"aliases": ["access info tls config cert params tls validation params trusted ca url"], "anchor": "schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca_url", "description": "Exclusive with Inline Root CA Certificate.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params", "trusted_ca_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info tls config cert params tls validation params verify subject alt names"], "anchor": "schema-access_info--tls_config--cert_params--tls_validation_params--verify_subject_alt_names", "description": "List of acceptable Subject Alt Names/CN in the peer's certificate. When skip_hostname_verification is false and verify_subject_alt_names is empty, the hostname of the peer will be used for matching against SAN/CN of peer's certificate.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params", "verify_subject_alt_names"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names for verification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.cert_params.tls_validation_params

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/)
- [access_info.tls_config.cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/cert_params/)
- access_info.tls_config.cert_params.tls_validation_params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
```

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

Terraform syntax:

```terraform
tls_validation_params {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-access_info--tls_config--cert_params--tls_validation_params--skip_hostname_verification"></a>

### skip_hostname_verification property

Type: `"bool"`. Optional.

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

- [trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/): complete subsection reference.

<a id="schema-access_info--tls_config--cert_params--tls_validation_params--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `["list", "string"]`. Optional.

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
