---
page_title: "default_pool.upstream_conn_pool_reuse_type"
subcategory: "Load Balancing"
description: "default_pool.upstream_conn_pool_reuse_type for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1811, "body_sha256": "sha256:769d40fbbb67c5624b3f9e516bb26ba0d4155823a6c34516f2e6e4d81deef66a", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:upstream_conn_pool_reuse_type", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:upstream_conn_pool_reuse_type:enable_conn_pool_reuse"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:upstream_conn_pool_reuse_type", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--upstream_conn_pool_reuse_type.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "upstream_conn_pool_reuse_type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/upstream_conn_pool_reuse_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.upstream_conn_pool_reuse_type for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.upstream_conn_pool_reuse_type

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- default_pool.upstream_conn_pool_reuse_type

<a id="section"></a>

Type: `"single"`. Computed.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

## Direct properties

- [disable_conn_pool_reuse](data-sources--http_loadbalancer--properties--default_pool--upstream_conn_pool_reuse_type--disable_conn_pool_reuse.md): complete subsection reference.

- [enable_conn_pool_reuse](data-sources--http_loadbalancer--properties--default_pool--upstream_conn_pool_reuse_type--enable_conn_pool_reuse.md): complete subsection reference.

## Next pages

- [default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse](data-sources--http_loadbalancer--properties--default_pool--upstream_conn_pool_reuse_type--disable_conn_pool_reuse.md)
- [default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse](data-sources--http_loadbalancer--properties--default_pool--upstream_conn_pool_reuse_type--enable_conn_pool_reuse.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
