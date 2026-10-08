---
page_title: "policy_rule_list"
subcategory: "Container"
description: "List of rules for role permissions."
xcsh_docs: {"aliases": ["policy rule list"], "body_bytes": 868, "body_sha256": "sha256:5e07e9612ffcfa23c08594613c90e5fa85d335e97451f9d3a13f66932f28197f", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list", "parent_id": "xcsh-docs:data-sources:k8s_cluster_role:reference", "path": "documentation/data-sources/k8s_cluster_role/properties/policy_rule_list/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1221133310032133-1000323320102332-1303231100200023-3322123110022320-3120213002030302-0233133133300231-2300021312001011-3030331123221023", "registry_path": "docs/guides/data-sources--k8s_cluster_role--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_rule_list"], "schema_version": 1, "sections": [{"aliases": ["policy rule list policy rule"], "anchor": "section", "description": "List of rules for role permissions.", "document_id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["policy_rule_list", "policy_rule"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster_role/properties/policy_rule_list/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "List of rules for role permissions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_rule_list

Breadcrumbs:

- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/)
- policy_rule_list

<a id="section"></a>

Type: `"single"`. Computed.

Policy Rule List. List of rules for role permissions.

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

- [policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/): complete subsection reference.
