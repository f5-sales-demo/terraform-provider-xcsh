---
page_title: "default_pool.no_tls"
subcategory: "Load Balancing"
description: "default_pool.no_tls for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 924, "body_sha256": "sha256:3f2e13023f4f3a2ea46dd75d69de7002a3d2367e64617d13f4be936549b88c3d", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:no_tls", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:no_tls", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--no_tls.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "no_tls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/no_tls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.no_tls for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# default_pool.no_tls

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- default_pool.no_tls

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
