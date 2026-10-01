---
page_title: "use_tls"
subcategory: "Load Balancing"
description: "use_tls for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 5765, "body_sha256": "sha256:380035413756742322a88cccdcbab4779a8c8caa7248d8f41543da13bbba2705", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:use_tls:default_session_key_caching", "xcsh-docs:data-sources:origin_pool:properties:use_tls:disable_session_key_caching", "xcsh-docs:data-sources:origin_pool:properties:use_tls:disable_sni", "xcsh-docs:data-sources:origin_pool:properties:use_tls:no_mtls", "xcsh-docs:data-sources:origin_pool:properties:use_tls:skip_server_verification", "xcsh-docs:data-sources:origin_pool:properties:use_tls:tls_config", "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_host_header_as_sni", "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls", "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls_obj", "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_server_verification", "xcsh-docs:data-sources:origin_pool:properties:use_tls:volterra_trusted_ca"], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:use_tls", "parent_id": "xcsh-docs:data-sources:origin_pool:reference", "path": "docs/guides/data-sources--origin_pool--properties--use_tls.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_tls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/use_tls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_tls for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- use_tls

<a id="section"></a>

Type: `"single"`. Computed.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Upstream description:

Upstream TLS Parameters.

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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

## Direct properties

- [default_session_key_caching](data-sources--origin_pool--properties--use_tls--default_session_key_caching.md): complete subsection reference.

- [disable_session_key_caching](data-sources--origin_pool--properties--use_tls--disable_session_key_caching.md): complete subsection reference.

- [disable_sni](data-sources--origin_pool--properties--use_tls--disable_sni.md): complete subsection reference.

<a id="schema-use_tls--max_session_keys"></a>

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

- [no_mtls](data-sources--origin_pool--properties--use_tls--no_mtls.md): complete subsection reference.

- [skip_server_verification](data-sources--origin_pool--properties--use_tls--skip_server_verification.md): complete subsection reference.

<a id="schema-use_tls--sni"></a>

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

- [tls_config](data-sources--origin_pool--properties--use_tls--tls_config.md): complete subsection reference.

- [use_host_header_as_sni](data-sources--origin_pool--properties--use_tls--use_host_header_as_sni.md): complete subsection reference.

- [use_mtls](data-sources--origin_pool--properties--use_tls--use_mtls.md): complete subsection reference.

- [use_mtls_obj](data-sources--origin_pool--properties--use_tls--use_mtls_obj.md): complete subsection reference.

- [use_server_verification](data-sources--origin_pool--properties--use_tls--use_server_verification.md): complete subsection reference.

- [volterra_trusted_ca](data-sources--origin_pool--properties--use_tls--volterra_trusted_ca.md): complete subsection reference.

## Next pages

- [use_tls.default_session_key_caching](data-sources--origin_pool--properties--use_tls--default_session_key_caching.md)
- [use_tls.disable_session_key_caching](data-sources--origin_pool--properties--use_tls--disable_session_key_caching.md)
- [use_tls.disable_sni](data-sources--origin_pool--properties--use_tls--disable_sni.md)
- [use_tls.no_mtls](data-sources--origin_pool--properties--use_tls--no_mtls.md)
- [use_tls.skip_server_verification](data-sources--origin_pool--properties--use_tls--skip_server_verification.md)
- [use_tls.tls_config](data-sources--origin_pool--properties--use_tls--tls_config.md)
- [use_tls.use_host_header_as_sni](data-sources--origin_pool--properties--use_tls--use_host_header_as_sni.md)
- [use_tls.use_mtls](data-sources--origin_pool--properties--use_tls--use_mtls.md)
- [use_tls.use_mtls_obj](data-sources--origin_pool--properties--use_tls--use_mtls_obj.md)
- [use_tls.use_server_verification](data-sources--origin_pool--properties--use_tls--use_server_verification.md)
- [use_tls.volterra_trusted_ca](data-sources--origin_pool--properties--use_tls--volterra_trusted_ca.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
