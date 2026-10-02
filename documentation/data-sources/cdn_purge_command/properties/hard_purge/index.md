---
page_title: "hard_purge"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["hard purge"], "body_bytes": 1472, "body_sha256": "sha256:b4ac756de975767e669ee7f8fdf492c3fa4dd277511e7c027d0128568b424c38", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cdn_purge_command:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_purge_command:properties:hard_purge", "parent_id": "xcsh-docs:data-sources:cdn_purge_command:reference", "path": "documentation/data-sources/cdn_purge_command/properties/hard_purge/index.md", "product": "distributed-cloud", "provider_name": "cdn_purge_command", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1003131013121223-0332023003220032-3022202203233103-2003100031213200-2033002123001211-1213330131031020-1010213010312013-0201320013312023", "registry_path": "docs/guides/data-sources--cdn_purge_command--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["hard_purge"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_purge_command/properties/hard_purge/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_purge_commandCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# hard_purge

Breadcrumbs:

- [xcsh_cdn_purge_command](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_purge_command/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_purge_command/properties/)
- hard_purge

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: hard\_purge, soft\_purge\] Enable this option

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

OneOf alternatives in this subsection:

- [hard_purge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_purge_command/properties/hard_purge/#section)
- [soft_purge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_purge_command/properties/soft_purge/#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_purge_command/properties/)
- [xcsh_cdn_purge_command](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_purge_command/)
