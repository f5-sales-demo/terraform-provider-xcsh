---
page_title: "https.default_header"
subcategory: "Load Balancing"
description: "https.default_header for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 959, "body_sha256": "sha256:f6ea3b704dca2279b6b4aa8c7fa23060432d77ecc8163fdbdf9f728be0d7a76d", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:default_header", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:default_header", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https", "path": "docs/guides/data-sources--http_loadbalancer--properties--https--default_header.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "default_header"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/https/default_header/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.default_header for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.default_header

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [https](data-sources--http_loadbalancer--properties--https.md)
- https.default_header

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

- [https](data-sources--http_loadbalancer--properties--https.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
