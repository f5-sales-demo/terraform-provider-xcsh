---
page_title: "advanced_options.proxy_protocol_v2"
subcategory: "Load Balancing"
description: "advanced_options.proxy_protocol_v2 for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1037, "body_sha256": "sha256:94dbfe7a1cef3e614f725a41b83c9e043e87114dd98c32ce7002b33564f05d48", "canonical_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v2", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v2", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "path": "docs/guides/resources--origin_pool--properties--advanced_options--proxy_protocol_v2.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options", "proxy_protocol_v2"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/proxy_protocol_v2/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.proxy_protocol_v2 for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.proxy_protocol_v2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- advanced_options.proxy_protocol_v2

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for proxy protocol v2.

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

Terraform syntax:

```terraform
proxy_protocol_v2 = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
