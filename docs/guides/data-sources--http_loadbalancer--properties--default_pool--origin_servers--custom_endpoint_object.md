---
page_title: "default_pool.origin_servers.custom_endpoint_object"
subcategory: "Load Balancing"
description: "default_pool.origin_servers.custom_endpoint_object for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1377, "body_sha256": "sha256:97706d0eee391838a1ce4b2cfd13ce3cd547c25a4f4f280d9b71449153693aa7", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:custom_endpoint_object", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:custom_endpoint_object:endpoint"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:custom_endpoint_object", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--origin_servers--custom_endpoint_object.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "origin_servers", "custom_endpoint_object"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/origin_servers/custom_endpoint_object/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.origin_servers.custom_endpoint_object for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.custom_endpoint_object

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- [default_pool.origin_servers](data-sources--http_loadbalancer--properties--default_pool--origin_servers.md)
- default_pool.origin_servers.custom_endpoint_object

<a id="section"></a>

Type: `"single"`. Computed.

Specify origin server with a reference to endpoint object.

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

- [endpoint](data-sources--http_loadbalancer--properties--default_pool--origin_servers--custom_endpoint_object--endpoint.md): complete subsection reference.

## Next pages

- [default_pool.origin_servers.custom_endpoint_object.endpoint](data-sources--http_loadbalancer--properties--default_pool--origin_servers--custom_endpoint_object--endpoint.md)
- [default_pool.origin_servers](data-sources--http_loadbalancer--properties--default_pool--origin_servers.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
