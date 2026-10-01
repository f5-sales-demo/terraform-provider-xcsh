---
page_title: "origin_pool.use_tls.tls_config.low_security"
subcategory: "Load Balancing"
description: "origin_pool.use_tls.tls_config.low_security for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1280, "body_sha256": "sha256:afcd384091f7a9c0e865deb77ba0d0cd8bf2f25315a26e5a50a7906fdcb881b5", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:tls_config:low_security", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:tls_config:low_security", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:tls_config", "path": "docs/guides/resources--cdn_loadbalancer--properties--origin_pool--use_tls--tls_config--low_security.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pool", "use_tls", "tls_config", "low_security"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/use_tls/tls_config/low_security/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pool.use_tls.tls_config.low_security for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.use_tls.tls_config.low_security

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [origin_pool](resources--cdn_loadbalancer--properties--origin_pool.md)
- [origin_pool.use_tls](resources--cdn_loadbalancer--properties--origin_pool--use_tls.md)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--properties--origin_pool--use_tls--tls_config.md)
- origin_pool.use_tls.tls_config.low_security

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
low_security = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--properties--origin_pool--use_tls--tls_config.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
