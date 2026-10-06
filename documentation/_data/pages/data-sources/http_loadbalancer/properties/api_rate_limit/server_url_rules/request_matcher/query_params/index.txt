---
page_title: "api_rate_limit.server_url_rules.request_matcher.query_params"
subcategory: "Load Balancing"
description: "A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that"
xcsh_docs: {"aliases": ["api rate limit server url rules request matcher query params"], "body_bytes": 4077, "body_sha256": "sha256:7cf2451f48b68aea4f6c7d130bd3fa354d1d27811e701e0628be934fa67c8853", "capabilities": ["load-balancing", "security.rate-limiting"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:query_params:check_not_present", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:query_params:check_present", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:query_params:item"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:query_params", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher", "path": "documentation/data-sources/http_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/query_params/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2200312211113322-0112121111320122-2320320321112330-3123202310120301-0202103223313212-2121113323301100-2031302101213301-3311210320012002", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "query_params"], "schema_version": 1, "sections": [{"aliases": ["api rate limit server url rules request matcher query params check not present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:query_params:check_not_present", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "query_params", "check_not_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit server url rules request matcher query params check present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:query_params:check_present", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "query_params", "check_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit server url rules request matcher query params invert matcher"], "anchor": "schema-api_rate_limit--server_url_rules--request_matcher--query_params--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:query_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "query_params", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["api rate limit server url rules request matcher query params item", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:query_params:item", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "query_params", "item"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit server url rules request matcher query params key"], "anchor": "schema-api_rate_limit--server_url_rules--request_matcher--query_params--key", "description": "A case-sensitive HTTP query parameter name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:query_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "query_params", "key"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/query_params/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.server_url_rules.request_matcher.query_params

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.server_url_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/server_url_rules/)
- [api_rate_limit.server_url_rules.request_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/)
- api_rate_limit.server_url_rules.request_matcher.query_params

<a id="section"></a>

Type: `"list"`. Computed.

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

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

## Direct properties

- [check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/query_params/check_not_present/): complete subsection reference.

- [check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/query_params/check_present/): complete subsection reference.

<a id="schema-api_rate_limit--server_url_rules--request_matcher--query_params--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

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

- [item](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/query_params/item/): complete subsection reference.

<a id="schema-api_rate_limit--server_url_rules--request_matcher--query_params--key"></a>

### key property

Type: `"string"`. Computed.

A case-sensitive HTTP query parameter name.

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
