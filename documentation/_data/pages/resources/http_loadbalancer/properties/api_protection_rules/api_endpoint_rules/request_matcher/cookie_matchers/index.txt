---
page_title: "api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers"
subcategory: "Load Balancing"
description: "A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must"
xcsh_docs: {"aliases": ["api protection rules api endpoint rules request matcher cookie matchers"], "body_bytes": 5402, "body_sha256": "sha256:bcff985ac78fc6fb82a88cbb64a8b21e486279d8f4d2e17778ec8539dfaa47e0", "capabilities": ["load-balancing", "security.api-protection"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers:check_not_present", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers:check_present", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers:item"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher", "path": "documentation/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/cookie_matchers/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers:item", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers:item", "type": "conflicts"}, {"anchor": "schema-api_protection_rules--api_endpoint_rules--request_matcher--cookie_matchers--name", "enforcement": "provider-schema", "group": "api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_protection_rules", "api_endpoint_rules", "request_matcher", "cookie_matchers"], "schema_version": 1, "sections": [{"aliases": ["api protection rules api endpoint rules request matcher cookie matchers check not present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers:check_not_present", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "request_matcher", "cookie_matchers", "check_not_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["api protection rules api endpoint rules request matcher cookie matchers check present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers:check_present", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "request_matcher", "cookie_matchers", "check_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["api protection rules api endpoint rules request matcher cookie matchers invert matcher"], "anchor": "schema-api_protection_rules--api_endpoint_rules--request_matcher--cookie_matchers--invert_matcher", "description": "Invert Match of the expression defined.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "request_matcher", "cookie_matchers", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["api protection rules api endpoint rules request matcher cookie matchers item", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers:item", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "request_matcher", "cookie_matchers", "item"], "syntax": "block", "type": "object"}, {"aliases": ["api protection rules api endpoint rules request matcher cookie matchers name"], "anchor": "schema-api_protection_rules--api_endpoint_rules--request_matcher--cookie_matchers--name", "description": "A case-sensitive cookie name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "request_matcher", "cookie_matchers", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/cookie_matchers/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/)
- [api_protection_rules.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/)
- [api_protection_rules.api_endpoint_rules.request_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
cookie_matchers {
  # Configure direct properties listed below.
}
```

## Direct properties

- [check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/cookie_matchers/check_not_present/): complete subsection reference.

- [check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/cookie_matchers/check_present/): complete subsection reference.

<a id="schema-api_protection_rules--api_endpoint_rules--request_matcher--cookie_matchers--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

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

- [item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/cookie_matchers/item/): complete subsection reference.

<a id="schema-api_protection_rules--api_endpoint_rules--request_matcher--cookie_matchers--name"></a>

### name property

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
