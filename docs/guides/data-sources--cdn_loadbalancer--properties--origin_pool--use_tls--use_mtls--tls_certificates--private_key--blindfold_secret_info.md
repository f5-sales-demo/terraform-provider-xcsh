---
page_title: "origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info"
subcategory: "Load Balancing"
description: "origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4831, "body_sha256": "sha256:db57d70b547d653241dca8aaacbe28c6ac368a9d89fd79a36a8e8db4bc5bcefe", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key:blindfold_secret_info", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key:blindfold_secret_info", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "private_key", "blindfold_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/private_key/blindfold_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [origin_pool](data-sources--cdn_loadbalancer--properties--origin_pool.md)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls.md)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls.md)
- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates.md)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates--private_key.md)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="section"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info--decryption_provider"></a>

### decryption_provider property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info--location"></a>

### location property

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info--store_provider"></a>

### store_provider property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates--private_key.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
