---
page_title: "policy_rule_list.policy_rule"
subcategory: "Container"
description: "List of rules for role permissions."
xcsh_docs: {"aliases": ["policy rule list policy rule"], "body_bytes": 2067, "body_sha256": "sha256:7e5b94e98e3089253ce2ff516bccf13a3afcb37edc7e47596e720b9a0a0af1d1", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule", "parent_id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list", "path": "documentation/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1231123021022003-1023022010131212-1013311331321110-0120011230122130-2100122120100221-3200002302323131-0202203323303002-2102030123102011", "registry_path": "docs/guides/data-sources--k8s_cluster_role--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_rule_list", "policy_rule"], "schema_version": 1, "sections": [{"aliases": ["policy rule list policy rule non resource url list"], "anchor": "section", "description": "Permissions for URL(s) that do not represent K8s resource.", "document_id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_rule_list", "policy_rule", "non_resource_url_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy rule list policy rule resource list"], "anchor": "section", "description": "List of resources in terms of API groups/resource types/resource instances and verbs allowed.", "document_id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_rule_list", "policy_rule", "resource_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of rules for role permissions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_rule_list.policy_rule

Breadcrumbs:

- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/)
- [policy_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/policy_rule_list/)
- policy_rule_list.policy_rule

<a id="section"></a>

Type: `"list"`. Computed.

Policy Rules. List of rules for role permissions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [non_resource_url_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/non_resource_url_list/): complete subsection reference.

- [resource_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/resource_list/): complete subsection reference.
