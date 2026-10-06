---
page_title: "waf_type.inherit_waf"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["waf type inherit waf"], "body_bytes": 936, "body_sha256": "sha256:9ae36222b88401efcb0d3bacc209f0eb5f8d06876977f2b25a7888614fb35c13", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:waf_type:inherit_waf", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:waf_type", "path": "documentation/data-sources/virtual_host/properties/waf_type/inherit_waf/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3111320320023322-3033330033011220-2031221222111000-3211122202020221-2303103121031102-0213033333303222-0321322313130302-2022131333210103", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_type", "inherit_waf"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/waf_type/inherit_waf/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_type.inherit_waf

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/waf_type/)
- waf_type.inherit_waf

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inherit waf.

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

This is an empty object or choice marker. It has no direct properties.
