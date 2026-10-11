---
page_title: "policy_rule_list.policy_rule.resource_list"
subcategory: "Container"
description: "List of resources in terms of API groups/resource types/resource instances and verbs allowed."
xcsh_docs: {"aliases": ["policy rule list policy rule resource list"], "body_bytes": 7189, "body_sha256": "sha256:d3285b68e4e5e654f6d7698da98dfca3b3ddb7779026ab58a80b0a2da6f7f8e7", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "parent_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule", "path": "documentation/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/resource_list/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1213332002311130-3231122203013000-2322301033131231-0011022002223223-2330322021103232-0312300212021310-2300212230331232-1003210302103211", "registry_path": "docs/guides/resources--k8s_cluster_role--reference--group-001.md", "relationships": [{"anchor": "schema-policy_rule_list--policy_rule--resource_list--api_groups", "enforcement": "provider-schema", "group": "policy_rule_list.policy_rule.resource_list:RequiredObjectAttributes:api_groups,resource_types,verbs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "type": "requires"}, {"anchor": "schema-policy_rule_list--policy_rule--resource_list--resource_types", "enforcement": "provider-schema", "group": "policy_rule_list.policy_rule.resource_list:RequiredObjectAttributes:api_groups,resource_types,verbs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "type": "requires"}, {"anchor": "schema-policy_rule_list--policy_rule--resource_list--verbs", "enforcement": "provider-schema", "group": "policy_rule_list.policy_rule.resource_list:RequiredObjectAttributes:api_groups,resource_types,verbs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_rule_list", "policy_rule", "resource_list"], "schema_version": 1, "sections": [{"aliases": ["policy rule list policy rule resource list api groups"], "anchor": "schema-policy_rule_list--policy_rule--resource_list--api_groups", "description": "Allowed list of API group that contains resources, all resources of a given API group.", "document_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_rule_list", "policy_rule", "resource_list", "api_groups"], "syntax": "attribute", "type": "list"}, {"aliases": ["policy rule list policy rule resource list resource instances"], "anchor": "schema-policy_rule_list--policy_rule--resource_list--resource_instances", "description": "Allowed list of resource instances within the resource types.", "document_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_rule_list", "policy_rule", "resource_list", "resource_instances"], "syntax": "attribute", "type": "list"}, {"aliases": ["policy rule list policy rule resource list resource types"], "anchor": "schema-policy_rule_list--policy_rule--resource_list--resource_types", "description": "Allowed list of resource types within the API groups.", "document_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_rule_list", "policy_rule", "resource_list", "resource_types"], "syntax": "attribute", "type": "list"}, {"aliases": ["policy rule list policy rule resource list verbs"], "anchor": "schema-policy_rule_list--policy_rule--resource_list--verbs", "description": "Allowed list of verbs(operations) on resources. Use * for all operations.", "document_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_rule_list", "policy_rule", "resource_list", "verbs"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/resource_list/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of resources in terms of API groups/resource types/resource instances and verbs allowed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_rule_list.policy_rule.resource_list

Breadcrumbs:

- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/)
- [policy_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/)
- [policy_rule_list.policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/)
- policy_rule_list.policy_rule.resource_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of resources in terms of API groups/resource types/resource instances and verbs allowed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("api_groups",
    "resource_types",
    "verbs")}
```

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

Terraform syntax:

```terraform
resource_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-policy_rule_list--policy_rule--resource_list--api_groups"></a>

### api_groups property

Type: `["list", "string"]`. Optional.

Allowed list of API group that contains resources, all resources of a given API group.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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

Type: `["list", "string"]`. Optional.

Allowed list of resource instances within the resource types.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Type: `["list", "string"]`. Optional.

Allowed list of resource types within the API groups.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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

Type: `["list", "string"]`. Optional.

Allowed list of verbs(operations) on resources. Use \* for all operations.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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
