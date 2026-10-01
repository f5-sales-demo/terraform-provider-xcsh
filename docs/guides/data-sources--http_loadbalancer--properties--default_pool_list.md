---
page_title: "default_pool_list"
subcategory: "Load Balancing"
description: "default_pool_list for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 999, "body_sha256": "sha256:777b91bb8586a2dc2bb81079a35cda325efa48155c02c859ad5dd69cd991729c", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool_list", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool_list:pools"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool_list", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool_list.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool_list for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool_list

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- default_pool_list

<a id="section"></a>

Type: `"single"`. Computed.

Origin Pool List Type. List of Origin Pools.

Upstream description:

List of Origin Pools.

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

- [pools](data-sources--http_loadbalancer--properties--default_pool_list--pools.md): complete subsection reference.

## Next pages

- [default_pool_list.pools](data-sources--http_loadbalancer--properties--default_pool_list--pools.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
