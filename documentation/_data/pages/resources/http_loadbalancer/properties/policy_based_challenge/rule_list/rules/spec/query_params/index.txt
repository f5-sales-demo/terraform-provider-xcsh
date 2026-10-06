---
page_title: "policy_based_challenge.rule_list.rules.spec.query_params"
subcategory: "Load Balancing"
description: "A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that"
xcsh_docs: {"aliases": ["policy based challenge rule list rules spec query params"], "body_bytes": 4920, "body_sha256": "sha256:4aef210e2a2dd1244912a55be468453bdc82e90d9026682d06134b85a77a3e3e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:check_not_present", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:check_present", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:item"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "path": "documentation/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/query_params/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-023.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.query_params:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.query_params:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.query_params:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.query_params:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.query_params:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:item", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.query_params:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:item", "type": "conflicts"}, {"anchor": "schema-policy_based_challenge--rule_list--rules--spec--query_params--key", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.query_params:RequiredListObjectAttributes:key", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "query_params"], "schema_version": 1, "sections": [{"aliases": ["policy based challenge rule list rules spec query params check not present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:check_not_present", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "query_params", "check_not_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec query params check present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:check_present", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "query_params", "check_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec query params invert matcher"], "anchor": "schema-policy_based_challenge--rule_list--rules--spec--query_params--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "query_params", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["policy based challenge rule list rules spec query params item", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:item", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "query_params", "item"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec query params key"], "anchor": "schema-policy_based_challenge--rule_list--rules--spec--query_params--key", "description": "A case-sensitive HTTP query parameter name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "query_params", "key"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/query_params/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules.spec.query_params

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/)
- [policy_based_challenge.rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/)
- [policy_based_challenge.rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/)
- [policy_based_challenge.rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/)
- policy_based_challenge.rule_list.rules.spec.query_params

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

## Direct properties

- [check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/query_params/check_not_present/): complete subsection reference.

- [check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/query_params/check_present/): complete subsection reference.

<a id="schema-policy_based_challenge--rule_list--rules--spec--query_params--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/query_params/item/): complete subsection reference.

<a id="schema-policy_based_challenge--rule_list--rules--spec--query_params--key"></a>

### key property

Type: `"string"`. Optional.

A case-sensitive HTTP query parameter name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```
