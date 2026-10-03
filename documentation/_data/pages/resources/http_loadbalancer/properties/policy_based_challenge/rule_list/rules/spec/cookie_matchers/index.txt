---
page_title: "policy_based_challenge.rule_list.rules.spec.cookie_matchers"
subcategory: "Load Balancing"
description: "A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must"
xcsh_docs: {"aliases": ["policy based challenge rule list rules spec cookie matchers"], "body_bytes": 6863, "body_sha256": "sha256:8b4fc3480be688690f5503c03c4cf82294b1726a3936161e4f9cf427d1428081", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers:check_not_present", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers:check_present", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers:item"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "path": "documentation/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/cookie_matchers/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-023.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.cookie_matchers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.cookie_matchers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.cookie_matchers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.cookie_matchers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.cookie_matchers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers:item", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.cookie_matchers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers:item", "type": "conflicts"}, {"anchor": "schema-policy_based_challenge--rule_list--rules--spec--cookie_matchers--name", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.cookie_matchers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "cookie_matchers"], "schema_version": 1, "sections": [{"aliases": ["policy based challenge rule list rules spec cookie matchers check not present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers:check_not_present", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "cookie_matchers", "check_not_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec cookie matchers check present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers:check_present", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "cookie_matchers", "check_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec cookie matchers invert matcher"], "anchor": "schema-policy_based_challenge--rule_list--rules--spec--cookie_matchers--invert_matcher", "description": "Invert Match of the expression defined.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "cookie_matchers", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["policy based challenge rule list rules spec cookie matchers item", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers:item", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "cookie_matchers", "item"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec cookie matchers name"], "anchor": "schema-policy_based_challenge--rule_list--rules--spec--cookie_matchers--name", "description": "A case-sensitive cookie name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "cookie_matchers", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/cookie_matchers/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules.spec.cookie_matchers

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/)
- [policy_based_challenge.rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/)
- [policy_based_challenge.rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/)
- [policy_based_challenge.rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
cookie_matchers {
  # Configure direct properties listed below.
}
```

## Direct properties

- [check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/cookie_matchers/check_not_present/): complete subsection reference.

- [check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/cookie_matchers/check_present/): complete subsection reference.

<a id="schema-policy_based_challenge--rule_list--rules--spec--cookie_matchers--invert_matcher"></a>

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

- [item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/cookie_matchers/item/): complete subsection reference.

<a id="schema-policy_based_challenge--rule_list--rules--spec--cookie_matchers--name"></a>

### name property

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

## Next pages

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/cookie_matchers/check_not_present/)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/cookie_matchers/check_present/)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/cookie_matchers/item/)
- [policy_based_challenge.rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
