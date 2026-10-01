---
page_title: "policy_rule_list.policy_rule.non_resource_url_list"
subcategory: "Container"
description: "policy_rule_list.policy_rule.non_resource_url_list for xcsh_k8s_cluster_role."
xcsh_docs: {"aliases": [], "body_bytes": 4048, "body_sha256": "sha256:7e8471e370e845e59fede401afe592edfff185dd234c1e44bdee5c1672055a73", "child_ids": [], "collection_id": "xcsh-docs:data-sources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "parent_id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule", "path": "documentation/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/non_resource_url_list/index.md", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["policy_rule_list", "policy_rule", "non_resource_url_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/non_resource_url_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_rule_list.policy_rule.non_resource_url_list for xcsh_k8s_cluster_role.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_rule_list.policy_rule.non_resource_url_list

Breadcrumbs:

- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/)
- [policy_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/policy_rule_list/)
- [policy_rule_list.policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/)
- policy_rule_list.policy_rule.non_resource_url_list

<a id="section"></a>

Type: `"single"`. Computed.

Permissions for URL(s) that do not represent K8s resource.

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

<a id="schema-policy_rule_list--policy_rule--non_resource_url_list--urls"></a>

### urls property

Type: `["list", "string"]`. Computed.

Allowed URL(s) that do not represent any K8s resource. URL can be suffix or regex.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-policy_rule_list--policy_rule--non_resource_url_list--verbs"></a>

### verbs property

Type: `["list", "string"]`. Computed.

Allowed list of verbs(operations) on resources. Use VerbAll for all operations.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [policy_rule_list.policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/)
- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/)
