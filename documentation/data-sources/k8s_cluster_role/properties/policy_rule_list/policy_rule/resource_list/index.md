---
page_title: "policy_rule_list.policy_rule.resource_list"
subcategory: "Container"
description: "List of resources in terms of API groups/resource types/resource instances and verbs allowed."
xcsh_docs: {"aliases": ["policy rule list policy rule resource list"], "body_bytes": 6209, "body_sha256": "sha256:8996e522670cd0f498a047decba945f412e61a2334895e76336b0f34a3395d4c", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "parent_id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule", "path": "documentation/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/resource_list/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-3312111220300133-2211133102232101-0011222300333020-2211023000112002-2230223223320322-2233123201230121-3310312202322121-1231113002303221", "registry_path": "docs/guides/data-sources--k8s_cluster_role--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_rule_list", "policy_rule", "resource_list"], "schema_version": 1, "sections": [{"aliases": ["policy rule list policy rule resource list api groups"], "anchor": "schema-policy_rule_list--policy_rule--resource_list--api_groups", "description": "Allowed list of API group that contains resources, all resources of a given API group.", "document_id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_rule_list", "policy_rule", "resource_list", "api_groups"], "syntax": "attribute", "type": "list"}, {"aliases": ["policy rule list policy rule resource list resource instances"], "anchor": "schema-policy_rule_list--policy_rule--resource_list--resource_instances", "description": "Allowed list of resource instances within the resource types.", "document_id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_rule_list", "policy_rule", "resource_list", "resource_instances"], "syntax": "attribute", "type": "list"}, {"aliases": ["policy rule list policy rule resource list resource types"], "anchor": "schema-policy_rule_list--policy_rule--resource_list--resource_types", "description": "Allowed list of resource types within the API groups.", "document_id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_rule_list", "policy_rule", "resource_list", "resource_types"], "syntax": "attribute", "type": "list"}, {"aliases": ["policy rule list policy rule resource list verbs"], "anchor": "schema-policy_rule_list--policy_rule--resource_list--verbs", "description": "Allowed list of verbs(operations) on resources. Use * for all operations.", "document_id": "xcsh-docs:data-sources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_rule_list", "policy_rule", "resource_list", "verbs"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/resource_list/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "List of resources in terms of API groups/resource types/resource instances and verbs allowed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_rule_list.policy_rule.resource_list

Breadcrumbs:

- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/)
- [policy_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/policy_rule_list/)
- [policy_rule_list.policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster_role/properties/policy_rule_list/policy_rule/)
- policy_rule_list.policy_rule.resource_list

<a id="section"></a>

Type: `"single"`. Computed.

List of resources in terms of API groups/resource types/resource instances and verbs allowed.

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

<a id="schema-policy_rule_list--policy_rule--resource_list--api_groups"></a>

### api_groups property

Type: `["list", "string"]`. Computed.

Allowed list of API group that contains resources, all resources of a given API group.

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
      "validatedAt": "2026-10-29T12:14:42+00:00"
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

<a id="schema-policy_rule_list--policy_rule--resource_list--resource_instances"></a>

### resource_instances property

Type: `["list", "string"]`. Computed.

Allowed list of resource instances within the resource types.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-29T12:14:42+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-policy_rule_list--policy_rule--resource_list--resource_types"></a>

### resource_types property

Type: `["list", "string"]`. Computed.

Allowed list of resource types within the API groups.

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
      "validatedAt": "2026-10-29T12:14:42+00:00"
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

<a id="schema-policy_rule_list--policy_rule--resource_list--verbs"></a>

### verbs property

Type: `["list", "string"]`. Computed.

Allowed list of verbs(operations) on resources. Use \* for all operations.

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
      "validatedAt": "2026-10-29T12:14:42+00:00"
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
