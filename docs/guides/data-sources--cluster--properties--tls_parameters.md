---
page_title: "tls_parameters"
subcategory: ""
description: "tls_parameters for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 4545, "body_sha256": "sha256:05524a75ad4f7ee3f1972de25b108a31145c0950d124b78aa85b8f64738f043f", "canonical_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters", "child_ids": ["xcsh-docs:data-sources:cluster:properties:tls_parameters:cert_params", "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params", "xcsh-docs:data-sources:cluster:properties:tls_parameters:default_session_key_caching", "xcsh-docs:data-sources:cluster:properties:tls_parameters:disable_session_key_caching", "xcsh-docs:data-sources:cluster:properties:tls_parameters:disable_sni", "xcsh-docs:data-sources:cluster:properties:tls_parameters:use_host_header_as_sni"], "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:tls_parameters", "parent_id": "xcsh-docs:data-sources:cluster:reference", "path": "docs/guides/data-sources--cluster--properties--tls_parameters.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/tls_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md)
- [Property reference](data-sources--cluster--reference.md)
- tls_parameters

<a id="section"></a>

Type: `"single"`. Computed.

TLS configuration for upstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]",
  "x-ves-oneof-field-tls_params_choice": "[\"cert_params\",\"common_params\"]"
}
```

## Direct properties

- [cert_params](data-sources--cluster--properties--tls_parameters--cert_params.md): complete subsection reference.

- [common_params](data-sources--cluster--properties--tls_parameters--common_params.md): complete subsection reference.

- [default_session_key_caching](data-sources--cluster--properties--tls_parameters--default_session_key_caching.md): complete subsection reference.

- [disable_session_key_caching](data-sources--cluster--properties--tls_parameters--disable_session_key_caching.md): complete subsection reference.

- [disable_sni](data-sources--cluster--properties--tls_parameters--disable_sni.md): complete subsection reference.

<a id="schema-tls_parameters--max_session_keys"></a>

### max_session_keys property

Type: `"number"`. Computed.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

<a id="schema-tls_parameters--sni"></a>

### sni property

Type: `"string"`. Computed.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [use_host_header_as_sni](data-sources--cluster--properties--tls_parameters--use_host_header_as_sni.md): complete subsection reference.

## Next pages

- [tls_parameters.cert_params](data-sources--cluster--properties--tls_parameters--cert_params.md)
- [tls_parameters.common_params](data-sources--cluster--properties--tls_parameters--common_params.md)
- [tls_parameters.default_session_key_caching](data-sources--cluster--properties--tls_parameters--default_session_key_caching.md)
- [tls_parameters.disable_session_key_caching](data-sources--cluster--properties--tls_parameters--disable_session_key_caching.md)
- [tls_parameters.disable_sni](data-sources--cluster--properties--tls_parameters--disable_sni.md)
- [tls_parameters.use_host_header_as_sni](data-sources--cluster--properties--tls_parameters--use_host_header_as_sni.md)
- [Property reference](data-sources--cluster--reference.md)
- [xcsh_cluster](../data-sources/cluster.md)
