---
page_title: "policy_rule_list.policy_rule.non_resource_url_list"
subcategory: "Container"
description: "Permissions for URL(s) that do not represent K8s resource."
xcsh_docs: {"aliases": ["policy rule list policy rule non resource url list"], "body_bytes": 4586, "body_sha256": "sha256:2266c59bbcadf2da43d2004528533f67e1b0d7c7a1eebf3a107353f853488a5b", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "parent_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule", "path": "documentation/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/non_resource_url_list/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0221101002211010-2332121033033023-3021032130332130-1233020012001021-2321102100233310-3202223331333212-2031221200210213-2311022221322023", "registry_path": "docs/guides/resources--k8s_cluster_role--reference--group-001.md", "relationships": [{"anchor": "schema-policy_rule_list--policy_rule--non_resource_url_list--urls", "enforcement": "provider-schema", "group": "policy_rule_list.policy_rule.non_resource_url_list:RequiredObjectAttributes:urls,verbs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "type": "requires"}, {"anchor": "schema-policy_rule_list--policy_rule--non_resource_url_list--verbs", "enforcement": "provider-schema", "group": "policy_rule_list.policy_rule.non_resource_url_list:RequiredObjectAttributes:urls,verbs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_rule_list", "policy_rule", "non_resource_url_list"], "schema_version": 1, "sections": [{"aliases": ["urls"], "anchor": "schema-policy_rule_list--policy_rule--non_resource_url_list--urls", "description": "Allowed URL(s) that do not represent any K8s resource. URL can be suffix or regex.", "document_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_rule_list", "policy_rule", "non_resource_url_list", "urls"], "syntax": "attribute", "type": "list"}, {"aliases": ["verbs"], "anchor": "schema-policy_rule_list--policy_rule--non_resource_url_list--verbs", "description": "Allowed list of verbs(operations) on resources. Use VerbAll for all operations.", "document_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_rule_list", "policy_rule", "non_resource_url_list", "verbs"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/non_resource_url_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Permissions for URL(s) that do not represent K8s resource.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_rule_list.policy_rule.non_resource_url_list

Breadcrumbs:

- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/)
- [policy_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/)
- [policy_rule_list.policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/)
- policy_rule_list.policy_rule.non_resource_url_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Permissions for URL(s) that do not represent K8s resource.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("urls",
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
non_resource_url_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-policy_rule_list--policy_rule--non_resource_url_list--urls"></a>

### urls property

Type: `["list", "string"]`. Optional.

Allowed URL(s) that do not represent any K8s resource. URL can be suffix or regex.

Provider validators and defaults (from schema source):

```go
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

Type: `["list", "string"]`. Optional.

Allowed list of verbs(operations) on resources. Use VerbAll for all operations.

Provider validators and defaults (from schema source):

```go
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

- [policy_rule_list.policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/)
- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/)
