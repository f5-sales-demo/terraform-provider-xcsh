---
page_title: "default_pool.same_as_endpoint_port"
subcategory: "Load Balancing"
description: "default_pool.same_as_endpoint_port for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1053, "body_sha256": "sha256:4c5056c28af361ee303d013be7042296728e6ba7991bc648ff82e359e95e853e", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:same_as_endpoint_port", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:same_as_endpoint_port", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--same_as_endpoint_port.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "same_as_endpoint_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/same_as_endpoint_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.same_as_endpoint_port for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.same_as_endpoint_port

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- default_pool.same_as_endpoint_port

<a id="section"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
