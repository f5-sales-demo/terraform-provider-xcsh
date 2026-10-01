---
page_title: "default_pool.advanced_options.disable_outlier_detection"
subcategory: "Load Balancing"
description: "default_pool.advanced_options.disable_outlier_detection for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1279, "body_sha256": "sha256:baa3ba0d4d9a4ab436fe078056cef8aca88d71602526d9fa2903a20884b1730a", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:disable_outlier_detection", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:disable_outlier_detection", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--advanced_options--disable_outlier_detection.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "advanced_options", "disable_outlier_detection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/advanced_options/disable_outlier_detection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.advanced_options.disable_outlier_detection for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.disable_outlier_detection

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- [default_pool.advanced_options](data-sources--http_loadbalancer--properties--default_pool--advanced_options.md)
- default_pool.advanced_options.disable_outlier_detection

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable outlier detection. Defaults to \`map\[\]\`. Server applies
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
