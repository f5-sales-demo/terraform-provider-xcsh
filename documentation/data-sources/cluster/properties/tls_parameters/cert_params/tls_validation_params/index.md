---
page_title: "tls_parameters.cert_params.tls_validation_params"
subcategory: ""
description: "This includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names for verification."
xcsh_docs: {"aliases": ["tls parameters cert params tls validation params"], "body_bytes": 3416, "body_sha256": "sha256:f780f29577d3cb8982fe384d2cbab4530cdb1710adb6b573429891ef3c1a1c15", "capabilities": ["load-balancing.tls"], "category": null, "child_ids": ["xcsh-docs:data-sources:cluster:properties:tls_parameters:cert_params:tls_validation_params:trusted_ca"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:cert_params:tls_validation_params", "parent_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:cert_params", "path": "documentation/data-sources/cluster/properties/tls_parameters/cert_params/tls_validation_params/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0311330302130000-1103112110213333-2011012322122311-2233023231023333-0223121323131323-0321101110002022-2023211210232300-0200112202213030", "registry_path": "docs/guides/data-sources--cluster--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "cert_params", "tls_validation_params"], "schema_version": 1, "sections": [{"aliases": ["tls parameters cert params tls validation params skip hostname verification"], "anchor": "schema-tls_parameters--cert_params--tls_validation_params--skip_hostname_verification", "description": "When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to the connecting hostname.", "document_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:cert_params:tls_validation_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "cert_params", "tls_validation_params", "skip_hostname_verification"], "syntax": "attribute", "type": "bool"}, {"aliases": ["tls parameters cert params tls validation params trusted ca"], "anchor": "section", "description": "Reference to Root CA Certificate.", "document_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:cert_params:tls_validation_params:trusted_ca", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_parameters", "cert_params", "tls_validation_params", "trusted_ca"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters cert params tls validation params trusted ca url"], "anchor": "schema-tls_parameters--cert_params--tls_validation_params--trusted_ca_url", "description": "Exclusive with Inline Root CA Certificate.", "document_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:cert_params:tls_validation_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "cert_params", "tls_validation_params", "trusted_ca_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls parameters cert params tls validation params verify subject alt names"], "anchor": "schema-tls_parameters--cert_params--tls_validation_params--verify_subject_alt_names", "description": "List of acceptable Subject Alt Names/CN in the peer's certificate. When skip_hostname_verification is false and verify_subject_alt_names is empty, the hostname of the peer will be used for matching against SAN/CN of peer's certificate.", "document_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:cert_params:tls_validation_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "cert_params", "tls_validation_params", "verify_subject_alt_names"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/tls_parameters/cert_params/tls_validation_params/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names for verification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.cert_params.tls_validation_params

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/)
- [tls_parameters.cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/cert_params/)
- tls_parameters.cert_params.tls_validation_params

<a id="section"></a>

Type: `"single"`. Computed.

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

<a id="schema-tls_parameters--cert_params--tls_validation_params--skip_hostname_verification"></a>

### skip_hostname_verification property

Type: `"bool"`. Computed.

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

- [trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/cert_params/tls_validation_params/trusted_ca/): complete subsection reference.

<a id="schema-tls_parameters--cert_params--tls_validation_params--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Computed.

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

<a id="schema-tls_parameters--cert_params--tls_validation_params--verify_subject_alt_names"></a>

### verify_subject_alt_names property

Type: `["list", "string"]`. Computed.

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
