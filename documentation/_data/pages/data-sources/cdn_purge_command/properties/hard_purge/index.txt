---
page_title: "hard_purge"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["hard purge"], "body_bytes": 1202, "body_sha256": "sha256:bb96e31b08bf78819079cab63b04df72ca56f85e055a56109233286c42bf7f1d", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cdn_purge_command:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_purge_command:properties:hard_purge", "parent_id": "xcsh-docs:data-sources:cdn_purge_command:reference", "path": "documentation/data-sources/cdn_purge_command/properties/hard_purge/index.md", "product": "distributed-cloud", "provider_name": "cdn_purge_command", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-1003131013121223-0332023003220032-3022202203233103-2003100031213200-2033002123001211-1213330131031020-1010213010312013-0201320013312023", "registry_path": "docs/guides/data-sources--cdn_purge_command--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["hard_purge"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_purge_command/properties/hard_purge/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cdn_purge_commandCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.
