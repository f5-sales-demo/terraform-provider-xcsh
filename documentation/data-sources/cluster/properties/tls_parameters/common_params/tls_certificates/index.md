---
page_title: "tls_parameters.common_params.tls_certificates"
subcategory: ""
description: "tls_parameters.common_params.tls_certificates for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 4548, "body_sha256": "sha256:00cf64c0a9cd362b8869e0e7cbca13334cf5c22f98a63b26c333b135960d3bfc", "child_ids": ["xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params:tls_certificates:custom_hash_algorithms", "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params:tls_certificates:disable_ocsp_stapling", "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key", "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params:tls_certificates:use_system_defaults"], "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params:tls_certificates", "parent_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params", "path": "documentation/data-sources/cluster/properties/tls_parameters/common_params/tls_certificates/index.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/tls_parameters/common_params/tls_certificates/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.common_params.tls_certificates for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_parameters.common_params.tls_certificates

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/)
- [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/)
- tls_parameters.common_params.tls_certificates

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

<a id="schema-tls_parameters--common_params--tls_certificates--certificate_url"></a>

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

- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/tls_certificates/custom_hash_algorithms/): complete subsection reference.

<a id="schema-tls_parameters--common_params--tls_certificates--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/tls_certificates/disable_ocsp_stapling/): complete subsection reference.

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/): complete subsection reference.

- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/tls_certificates/use_system_defaults/): complete subsection reference.

## Next pages

- [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/tls_certificates/custom_hash_algorithms/)
- [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/tls_certificates/disable_ocsp_stapling/)
- [tls_parameters.common_params.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/)
- [tls_parameters.common_params.tls_certificates.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/tls_certificates/use_system_defaults/)
- [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
