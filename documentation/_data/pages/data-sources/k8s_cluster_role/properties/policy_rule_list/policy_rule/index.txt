---
page_title: "policy_rule_list.policy_rule"
subcategory: "Container"
description: "policy_rule_list.policy_rule for xcsh_k8s_cluster_role."
xcsh_docs: {"aliases": [], "body_bytes": 2788, "body_sha256": "sha256:f4eb9c08db2665e0f0e89e3985c1b91116fba8554838c3d77e1461c125658b97", "child_ids": ["xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list"], "collection_id": "xcsh-docs:data-sources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule", "parent_id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list", "path": "documentation/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/index.md", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["policy_rule_list", "policy_rule"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_rule_list.policy_rule for xcsh_k8s_cluster_role.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

List of rules for role permissions.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [policy_rule_list.policy_rule.non_resource_url_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/non_resource_url_list/)
- [policy_rule_list.policy_rule.resource_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/resource_list/)
- [policy_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/policy_rule_list/)
- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/)
