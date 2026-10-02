---
page_title: "user_defined_api_discovery_policy.exclusive"
subcategory: ""
description: "Configuration for exclusion action."
xcsh_docs: {"aliases": ["user defined api discovery policy exclusive"], "body_bytes": 2118, "body_sha256": "sha256:147e0c93933ceedb1b31ce49cac4a4c08d4fda5f7668c55cd202485ff223850f", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:archive", "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:ignore"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "parent_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy", "path": "documentation/data-sources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2300322012120002-0201232221212210-2123211221021011-2200001022101221-2331331221020112-0010131103023101-3003002313313310-3030102012302313", "registry_path": "docs/guides/data-sources--api_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "exclusive"], "schema_version": 1, "sections": [{"aliases": ["archive"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:archive", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "exclusive", "archive"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:exclusive:ignore", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "exclusive", "ignore"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration for exclusion action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

Configuration for exclusion action.

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

## Next pages

- [user_defined_api_discovery_policy.exclusive.archive](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/archive/)
- [user_defined_api_discovery_policy.exclusive.ignore](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/ignore/)
- [user_defined_api_discovery_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/)
- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/)
