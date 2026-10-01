---
page_title: "origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults"
subcategory: "Load Balancing"
description: "origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1532, "body_sha256": "sha256:67c18ec379c36a22264ce60a338d48eb1f8c45401904fce8c7b2f292abb4900c", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:use_system_defaults", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:use_system_defaults", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates", "path": "docs/guides/resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates--use_system_defaults.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "use_system_defaults"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/use_system_defaults/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [origin_pool](resources--cdn_loadbalancer--properties--origin_pool.md)
- [origin_pool.use_tls](resources--cdn_loadbalancer--properties--origin_pool--use_tls.md)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls.md)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates.md)
- origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
