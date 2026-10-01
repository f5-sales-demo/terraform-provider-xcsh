---
page_title: "default_pool.use_tls.tls_config"
subcategory: "Load Balancing"
description: "default_pool.use_tls.tls_config for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2751, "body_sha256": "sha256:6926d9326fcd5e64269ca5c30d867771295044a468bcde8ac6c2eeb956298a96", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config:custom_security", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config:default_security", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config:low_security", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config:medium_security"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--use_tls--tls_config.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "use_tls", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/use_tls/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.use_tls.tls_config for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.use_tls.tls_config

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- [default_pool.use_tls](data-sources--http_loadbalancer--properties--default_pool--use_tls.md)
- default_pool.use_tls.tls_config

<a id="section"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

## Direct properties

- [custom_security](data-sources--http_loadbalancer--properties--default_pool--use_tls--tls_config--custom_security.md): complete subsection reference.

- [default_security](data-sources--http_loadbalancer--properties--default_pool--use_tls--tls_config--default_security.md): complete subsection reference.

- [low_security](data-sources--http_loadbalancer--properties--default_pool--use_tls--tls_config--low_security.md): complete subsection reference.

- [medium_security](data-sources--http_loadbalancer--properties--default_pool--use_tls--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [default_pool.use_tls.tls_config.custom_security](data-sources--http_loadbalancer--properties--default_pool--use_tls--tls_config--custom_security.md)
- [default_pool.use_tls.tls_config.default_security](data-sources--http_loadbalancer--properties--default_pool--use_tls--tls_config--default_security.md)
- [default_pool.use_tls.tls_config.low_security](data-sources--http_loadbalancer--properties--default_pool--use_tls--tls_config--low_security.md)
- [default_pool.use_tls.tls_config.medium_security](data-sources--http_loadbalancer--properties--default_pool--use_tls--tls_config--medium_security.md)
- [default_pool.use_tls](data-sources--http_loadbalancer--properties--default_pool--use_tls.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
