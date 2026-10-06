---
page_title: "no_policer"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["no policer"], "body_bytes": 1209, "body_sha256": "sha256:31f22c41b18d317956d22e7a9926c747915ee01b8735f39030d8424484b8a5fb", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:forwarding_class:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forwarding_class:properties:no_policer", "parent_id": "xcsh-docs:data-sources:forwarding_class:reference", "path": "documentation/data-sources/forwarding_class/properties/no_policer/index.md", "product": "distributed-cloud", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-0132002012332231-3230221000012120-0132211221330131-0112303112310211-2132222113213031-0223021031202132-3200211212100232-0203202200211203", "registry_path": "docs/guides/data-sources--forwarding_class--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["no_policer"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forwarding_class/properties/no_policer/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# no_policer

Breadcrumbs:

- [xcsh_forwarding_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/)
- no_policer

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: no\_policer, policer; Default: no\_policer\] Enable this option

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

- [no_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/no_policer/#section)
- [policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/policer/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
