---
page_title: "default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults"
subcategory: "Load Balancing"
description: "default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1554, "body_sha256": "sha256:9139f43470b631eaed510c3003cc6df03b647fcf4f6765cf86cb51f3e27c0c2d", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_mtls:tls_certificates:use_system_defaults", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_mtls:tls_certificates:use_system_defaults", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_mtls:tls_certificates", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--use_tls--use_mtls--tls_certificates--use_system_defaults.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "use_tls", "use_mtls", "tls_certificates", "use_system_defaults"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/use_tls/use_mtls/tls_certificates/use_system_defaults/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.use_tls](resources--http_loadbalancer--properties--default_pool--use_tls.md)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--properties--default_pool--use_tls--use_mtls.md)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--properties--default_pool--use_tls--use_mtls--tls_certificates.md)
- default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

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

- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--properties--default_pool--use_tls--use_mtls--tls_certificates.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
