---
page_title: "origin_pool.use_tls"
subcategory: "Load Balancing"
description: "origin_pool.use_tls for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 6456, "body_sha256": "sha256:6b264251cdccf622582dfeba696726b1803b066bc4b1297efc54ab0056b0ca33", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:default_session_key_caching", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:disable_session_key_caching", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:disable_sni", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:no_mtls", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:skip_server_verification", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:tls_config", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_host_header_as_sni", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls_obj", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_server_verification", "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:volterra_trusted_ca"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--origin_pool--use_tls.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pool", "use_tls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pool.use_tls for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.use_tls

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [origin_pool](data-sources--cdn_loadbalancer--properties--origin_pool.md)
- origin_pool.use_tls

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

- [default_session_key_caching](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--default_session_key_caching.md): complete subsection reference.

- [disable_session_key_caching](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--disable_session_key_caching.md): complete subsection reference.

- [disable_sni](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--disable_sni.md): complete subsection reference.

<a id="schema-origin_pool--use_tls--max_session_keys"></a>

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

- [no_mtls](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--no_mtls.md): complete subsection reference.

- [skip_server_verification](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--skip_server_verification.md): complete subsection reference.

<a id="schema-origin_pool--use_tls--sni"></a>

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

- [tls_config](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--tls_config.md): complete subsection reference.

- [use_host_header_as_sni](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--use_host_header_as_sni.md): complete subsection reference.

- [use_mtls](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls.md): complete subsection reference.

- [use_mtls_obj](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls_obj.md): complete subsection reference.

- [use_server_verification](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--use_server_verification.md): complete subsection reference.

- [volterra_trusted_ca](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--volterra_trusted_ca.md): complete subsection reference.

## Next pages

- [origin_pool.use_tls.default_session_key_caching](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--default_session_key_caching.md)
- [origin_pool.use_tls.disable_session_key_caching](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--disable_session_key_caching.md)
- [origin_pool.use_tls.disable_sni](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--disable_sni.md)
- [origin_pool.use_tls.no_mtls](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--no_mtls.md)
- [origin_pool.use_tls.skip_server_verification](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--skip_server_verification.md)
- [origin_pool.use_tls.tls_config](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--tls_config.md)
- [origin_pool.use_tls.use_host_header_as_sni](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--use_host_header_as_sni.md)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls.md)
- [origin_pool.use_tls.use_mtls_obj](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls_obj.md)
- [origin_pool.use_tls.use_server_verification](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--use_server_verification.md)
- [origin_pool.use_tls.volterra_trusted_ca](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--volterra_trusted_ca.md)
- [origin_pool](data-sources--cdn_loadbalancer--properties--origin_pool.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
