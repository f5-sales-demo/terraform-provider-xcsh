---
page_title: "cache_rules.rule_expression_list"
subcategory: ""
description: "Expressions are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs.."
xcsh_docs: {"aliases": ["cache rules rule expression list"], "body_bytes": 3882, "body_sha256": "sha256:96fe777fc667ac9572f5df105c9eda610ca99337d1d9ec366e2d2c437bb242a7", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list", "parent_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules", "path": "documentation/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3110102003302131-2302312133210100-3031102231130310-1111030113012121-2333122333330303-0312333103333000-2013121121313131-3232213102311232", "registry_path": "docs/guides/resources--cdn_cache_rule--reference--group-001.md", "relationships": [{"anchor": "schema-cache_rules--rule_expression_list--expression_name", "enforcement": "provider-schema", "group": "cache_rules.rule_expression_list:RequiredListObjectAttributes:cache_rule_expression,expression_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cache_rules.rule_expression_list:RequiredListObjectAttributes:cache_rule_expression,expression_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_rules", "rule_expression_list"], "schema_version": 1, "sections": [{"aliases": ["cache rules rule expression list cache rule expression"], "anchor": "section", "description": "The Cache Rule Expression Terms that are ANDed.", "document_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression"], "syntax": "block", "type": "object"}, {"aliases": ["cache rules rule expression list expression name"], "anchor": "schema-cache_rules--rule_expression_list--expression_name", "description": "Name of the Expressions items that are ANDed.", "document_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "expression_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Expressions are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs..", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.rule_expression_list

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/)
- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/)
- cache_rules.rule_expression_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Expressions are evaluated in the order in which they are specified. The evaluation stops when the
first rule match occurs.

Upstream description:

Expressions are evaluated in the order in which they are specified. The evaluation stops when the
first rule match occurs..

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("cache_rule_expression",
    "expression_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rule_expression_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cache_rule_expression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/): complete subsection reference.

<a id="schema-cache_rules--rule_expression_list--expression_name"></a>

### expression_name property

Type: `"string"`. Optional.

Name of the Expressions items that are ANDed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

## Next pages

- [cache_rules.rule_expression_list.cache_rule_expression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/)
- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/)
- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/)
