---
page_title: "advanced_options.http2_options"
subcategory: "Load Balancing"
description: "advanced_options.http2_options for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1276, "body_sha256": "sha256:79b90e1dd94b5e227d947e088a3a82e5464d24f7db83515cd61d9a1fd000dfba", "canonical_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http2_options", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http2_options", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "path": "docs/guides/resources--origin_pool--properties--advanced_options--http2_options.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options", "http2_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/http2_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.http2_options for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.http2_options

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- advanced_options.http2_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http2_options {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-advanced_options--http2_options--enabled"></a>

### enabled property

Type: `"bool"`. Optional.

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

- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
