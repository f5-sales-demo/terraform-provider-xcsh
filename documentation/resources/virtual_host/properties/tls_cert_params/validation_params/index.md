---
page_title: "tls_cert_params.validation_params"
subcategory: ""
description: "This includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names for verification."
xcsh_docs: {"aliases": ["tls cert params validation params"], "body_bytes": 3316, "body_sha256": "sha256:c16d026710dd29dede3fbf67c849e00497e6ea6edbcd4f8f272a28e93e763ef2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:properties:tls_cert_params:validation_params:trusted_ca"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:validation_params", "parent_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params", "path": "documentation/resources/virtual_host/properties/tls_cert_params/validation_params/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-2111212302120220-0101203001130310-1133231301002100-3121101200220032-0321213313100212-2200021303113021-0023002112012032-3022331122002112", "registry_path": "docs/guides/resources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_cert_params", "validation_params"], "schema_version": 1, "sections": [{"aliases": ["tls cert params validation params skip hostname verification"], "anchor": "schema-tls_cert_params--validation_params--skip_hostname_verification", "description": "When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to the connecting hostname.", "document_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:validation_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_cert_params", "validation_params", "skip_hostname_verification"], "syntax": "attribute", "type": "bool"}, {"aliases": ["tls cert params validation params trusted ca"], "anchor": "section", "description": "Reference to Root CA Certificate.", "document_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:validation_params:trusted_ca", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_cert_params", "validation_params", "trusted_ca"], "syntax": "block", "type": "object"}, {"aliases": ["tls cert params validation params trusted ca url"], "anchor": "schema-tls_cert_params--validation_params--trusted_ca_url", "description": "Exclusive with Inline Root CA Certificate.", "document_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:validation_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_cert_params", "validation_params", "trusted_ca_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls cert params validation params verify subject alt names"], "anchor": "schema-tls_cert_params--validation_params--verify_subject_alt_names", "description": "List of acceptable Subject Alt Names/CN in the peer's certificate. When skip_hostname_verification is false and verify_subject_alt_names is empty, the hostname of the peer will be used for matching against SAN/CN of peer's certificate.", "document_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:validation_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_cert_params", "validation_params", "verify_subject_alt_names"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/tls_cert_params/validation_params/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names for verification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_cert_params.validation_params

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/)
- tls_cert_params.validation_params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

- [trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/trusted_ca/): complete subsection reference.

<a id="schema-tls_cert_params--validation_params--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-16T10:31:02+00:00"
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
