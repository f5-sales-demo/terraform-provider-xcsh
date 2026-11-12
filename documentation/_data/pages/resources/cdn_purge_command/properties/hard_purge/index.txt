---
page_title: "hard_purge"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["hard purge"], "body_bytes": 1243, "body_sha256": "sha256:1e2cde5a34bb0c72981952d189f08b8fdc150f4da6abf3455af3c4be3ecc04a9", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cdn_purge_command:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_purge_command:properties:hard_purge", "parent_id": "xcsh-docs:resources:cdn_purge_command:reference", "path": "documentation/resources/cdn_purge_command/properties/hard_purge/index.md", "product": "distributed-cloud", "provider_name": "cdn_purge_command", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-1131033021013223-2101330132310303-1131303103232322-3010100202300132-2222302013012332-0020231111311021-0230211131321232-1221221210312303", "registry_path": "docs/guides/resources--cdn_purge_command--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["hard_purge"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_purge_command/properties/hard_purge/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["cdn_purge_commandCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# hard_purge

Breadcrumbs:

- [xcsh_cdn_purge_command](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_purge_command/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_purge_command/properties/)
- hard_purge

<a id="section"></a>

Type: `["object", {}]`. Optional.

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

- [hard_purge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_purge_command/properties/hard_purge/#section)
- [soft_purge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_purge_command/properties/soft_purge/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
hard_purge = {}
```

This is an empty object or choice marker. It has no direct properties.
