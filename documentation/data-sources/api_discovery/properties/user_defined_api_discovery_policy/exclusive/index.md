---
page_title: "user_defined_api_discovery_policy.exclusive"
subcategory: ""
description: "Configuration for exclusion action."
xcsh_docs: {"aliases": ["user defined api discovery policy exclusive"], "body_bytes": 1360, "body_sha256": "sha256:2450407daaa01d4a10925f0396801877275bd3423395273fb1331451c301d988", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:archive", "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:ignore"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "parent_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy", "path": "documentation/data-sources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2300322012120002-0201232221212210-2123211221021011-2200001022101221-2331331221020112-0010131103023101-3003002313313310-3030102012302313", "registry_path": "docs/guides/data-sources--api_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "exclusive"], "schema_version": 1, "sections": [{"aliases": ["user defined api discovery policy exclusive archive"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:archive", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "exclusive", "archive"], "syntax": "attribute", "type": "object"}, {"aliases": ["user defined api discovery policy exclusive ignore"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:ignore", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "exclusive", "ignore"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Configuration for exclusion action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy.exclusive

Breadcrumbs:

- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/)
- [user_defined_api_discovery_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/)
- user_defined_api_discovery_policy.exclusive

<a id="section"></a>

Type: `"single"`. Computed.

Exclusion Configuration. Configuration for exclusion action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"archive\",\"ignore\"]"
}
```

## Direct properties

- [archive](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/archive/): complete subsection reference.

- [ignore](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/ignore/): complete subsection reference.
