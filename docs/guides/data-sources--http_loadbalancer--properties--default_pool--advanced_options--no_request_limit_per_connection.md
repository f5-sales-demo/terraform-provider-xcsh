---
page_title: "default_pool.advanced_options.no_request_limit_per_connection"
subcategory: "Load Balancing"
description: "default_pool.advanced_options.no_request_limit_per_connection for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1297, "body_sha256": "sha256:b1b68cd5d91ad221615da701c890f5aa9d2914d827b3b8878847e71275793cad", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:no_request_limit_per_connection", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:no_request_limit_per_connection", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--advanced_options--no_request_limit_per_connection.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "advanced_options", "no_request_limit_per_connection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/advanced_options/no_request_limit_per_connection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.advanced_options.no_request_limit_per_connection for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.no_request_limit_per_connection

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- [default_pool.advanced_options](data-sources--http_loadbalancer--properties--default_pool--advanced_options.md)
- default_pool.advanced_options.no_request_limit_per_connection

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no request limit per connection. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool.advanced_options](data-sources--http_loadbalancer--properties--default_pool--advanced_options.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
