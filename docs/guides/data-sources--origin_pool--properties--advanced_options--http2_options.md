---
page_title: "advanced_options.http2_options"
subcategory: "Load Balancing"
description: "advanced_options.http2_options for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1070, "body_sha256": "sha256:fb6102ddfbe787c6facd4f1ad3144786fc1a0e42f49173dad6620b9781f3925f", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http2_options", "child_ids": [], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http2_options", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options", "path": "docs/guides/data-sources--origin_pool--properties--advanced_options--http2_options.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options", "http2_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/advanced_options/http2_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.http2_options for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# advanced_options.http2_options

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [advanced_options](data-sources--origin_pool--properties--advanced_options.md)
- advanced_options.http2_options

<a id="section"></a>

Type: `"single"`. Computed.

Http2 Protocol OPTIONS for upstream connections.

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

<a id="schema-advanced_options--http2_options--enabled"></a>

### enabled property

Type: `"bool"`. Computed.

Enable/disable HTTP2 Protocol for upstream connections.

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

## Next pages

- [advanced_options](data-sources--origin_pool--properties--advanced_options.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
