---
page_title: "tls_cert_params.validation_params"
subcategory: ""
description: "tls_cert_params.validation_params for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 4247, "body_sha256": "sha256:aa4c0ea28e33981db4e1354d58e69c94d38e7d1b4ee1e8d3d8a0efe73039b3f1", "canonical_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:validation_params", "child_ids": ["xcsh-docs:resources:virtual_host:properties:tls_cert_params:validation_params:trusted_ca"], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:validation_params", "parent_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params", "path": "docs/guides/resources--virtual_host--properties--tls_cert_params--validation_params.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_cert_params", "validation_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/tls_cert_params/validation_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_cert_params.validation_params for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_cert_params.validation_params

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [tls_cert_params](resources--virtual_host--properties--tls_cert_params.md)
- tls_cert_params.validation_params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

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
validation_params {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-tls_cert_params--validation_params--skip_hostname_verification"></a>

### skip_hostname_verification property

Type: `"bool"`. Optional.

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

- [trusted_ca](resources--virtual_host--properties--tls_cert_params--validation_params--trusted_ca.md): complete subsection reference.

<a id="schema-tls_cert_params--validation_params--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

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

<a id="schema-tls_cert_params--validation_params--verify_subject_alt_names"></a>

### verify_subject_alt_names property

Type: `["list", "string"]`. Optional.

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

- [tls_cert_params.validation_params.trusted_ca](resources--virtual_host--properties--tls_cert_params--validation_params--trusted_ca.md)
- [tls_cert_params](resources--virtual_host--properties--tls_cert_params.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
