---
page_title: "default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse"
subcategory: "Load Balancing"
description: "default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1337, "body_sha256": "sha256:cdff0a86b872d54fa5aff30960c2f0ccf189d83fe7d72093db7676372d42679b", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:upstream_conn_pool_reuse_type", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--upstream_conn_pool_reuse_type--disable_conn_pool_reuse.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "upstream_conn_pool_reuse_type", "disable_conn_pool_reuse"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--properties--default_pool--upstream_conn_pool_reuse_type.md)
- default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable conn pool reuse.

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
disable_conn_pool_reuse = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--properties--default_pool--upstream_conn_pool_reuse_type.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
