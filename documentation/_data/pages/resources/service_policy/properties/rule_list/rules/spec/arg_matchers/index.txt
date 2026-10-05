---
page_title: "rule_list.rules.spec.arg_matchers"
subcategory: "Security"
description: "A list of predicates for all POST args that need to be matched. The criteria for matching each arg are described in individual instances of ArgMatcherType. The actual arg values are extracted from the request API as a list of strings for each arg selector name. Note that all specified arg matcher predicates must"
xcsh_docs: {"aliases": ["rule list rules spec arg matchers"], "body_bytes": 6332, "body_sha256": "sha256:bc93e1417285a123905a55cd49832e77bc0a86c7b72142f15a427c9c910ca0f7", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers:check_not_present", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers:check_present", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers:item"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "path": "documentation/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0322010120330120-3111202333230122-2110113300222122-2320031332000032-3313112022233210-3020313211100101-0221022311002011-0012110003033112", "registry_path": "docs/guides/resources--service_policy--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.arg_matchers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.arg_matchers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.arg_matchers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.arg_matchers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.arg_matchers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers:item", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.arg_matchers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers:item", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--arg_matchers--name", "enforcement": "provider-schema", "group": "rule_list.rules.spec.arg_matchers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec", "arg_matchers"], "schema_version": 1, "sections": [{"aliases": ["rule list rules spec arg matchers check not present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers:check_not_present", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "arg_matchers", "check_not_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec arg matchers check present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers:check_present", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "arg_matchers", "check_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec arg matchers invert matcher"], "anchor": "schema-rule_list--rules--spec--arg_matchers--invert_matcher", "description": "Invert Match of the expression defined.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "arg_matchers", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["rule list rules spec arg matchers item", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers:item", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "arg_matchers", "item"], "syntax": "block", "type": "object"}, {"aliases": ["rule list rules spec arg matchers name"], "anchor": "schema-rule_list--rules--spec--arg_matchers--name", "description": "A case-sensitive JSON path in the HTTP request body.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "arg_matchers", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "A list of predicates for all POST args that need to be matched. The criteria for matching each arg are described in individual instances of ArgMatcherType. The actual arg values are extracted from the request API as a list of strings for each arg selector name. Note that all specified arg matcher predicates must", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.arg_matchers

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/)
- rule_list.rules.spec.arg_matchers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all POST args that need to be matched. The criteria for matching each arg are
described in individual instances of ArgMatcherType. The actual arg values are extracted from the
request API as a list of strings for each arg selector name.

Upstream description:

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
arg_matchers {
  # Configure direct properties listed below.
}
```

## Direct properties

- [check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/check_not_present/): complete subsection reference.

- [check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/check_present/): complete subsection reference.

<a id="schema-rule_list--rules--spec--arg_matchers--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/item/): complete subsection reference.

<a id="schema-rule_list--rules--spec--arg_matchers--name"></a>

### name property

Type: `"string"`. Optional.

Case-sensitive JSON path in the HTTP request body.

Upstream description:

A case-sensitive JSON path in the HTTP request body.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
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
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

## Next pages

- [rule_list.rules.spec.arg_matchers.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/check_not_present/)
- [rule_list.rules.spec.arg_matchers.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/check_present/)
- [rule_list.rules.spec.arg_matchers.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/item/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
