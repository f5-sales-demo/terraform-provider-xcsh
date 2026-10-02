---
page_title: "api_rate_limit.server_url_rules.request_matcher.headers"
subcategory: "Load Balancing"
description: "A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header"
xcsh_docs: {"aliases": ["api rate limit server url rules request matcher headers"], "body_bytes": 6787, "body_sha256": "sha256:0a48f10f68238cba181a0b4820b8ae4d31df51a703c482a1473394f620680d62", "capabilities": ["cdn", "security.rate-limiting"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers:check_not_present", "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers:check_present", "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers:item"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher", "path": "documentation/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/headers/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.request_matcher.headers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.request_matcher.headers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.request_matcher.headers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.request_matcher.headers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.request_matcher.headers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers:item", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.request_matcher.headers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers:item", "type": "conflicts"}, {"anchor": "schema-api_rate_limit--server_url_rules--request_matcher--headers--name", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.request_matcher.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "headers"], "schema_version": 1, "sections": [{"aliases": ["check not present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers:check_not_present", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "headers", "check_not_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["check present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers:check_present", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "headers", "check_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["invert matcher"], "anchor": "schema-api_rate_limit--server_url_rules--request_matcher--headers--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "headers", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["item", "login success", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers:item", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "headers", "item"], "syntax": "block", "type": "object"}, {"aliases": ["name"], "anchor": "schema-api_rate_limit--server_url_rules--request_matcher--headers--name", "description": "A case-insensitive HTTP header name.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:headers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "headers", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/headers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.server_url_rules.request_matcher.headers

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.server_url_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/)
- [api_rate_limit.server_url_rules.request_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/)
- api_rate_limit.server_url_rules.request_matcher.headers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

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
    },
    "minItems": 0,
    "uniqueItems": false
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
headers {
  # Configure direct properties listed below.
}
```

## Direct properties

- [check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/headers/check_not_present/): complete subsection reference.

- [check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/headers/check_present/): complete subsection reference.

<a id="schema-api_rate_limit--server_url_rules--request_matcher--headers--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/headers/item/): complete subsection reference.

<a id="schema-api_rate_limit--server_url_rules--request_matcher--headers--name"></a>

### name property

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

## Next pages

- [api_rate_limit.server_url_rules.request_matcher.headers.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/headers/check_not_present/)
- [api_rate_limit.server_url_rules.request_matcher.headers.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/headers/check_present/)
- [api_rate_limit.server_url_rules.request_matcher.headers.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/headers/item/)
- [api_rate_limit.server_url_rules.request_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
