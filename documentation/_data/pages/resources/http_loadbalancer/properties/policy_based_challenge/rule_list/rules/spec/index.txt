---
page_title: "policy_based_challenge.rule_list.rules.spec"
subcategory: "Load Balancing"
description: "A Challenge Rule consists of an unordered list of predicates and an action. The predicates are evaluated against a set of input fields that are extracted from or derived from an L7 request API. A request API is considered to match the rule if all predicates in the rule evaluate to true for that request. Any predicates"
xcsh_docs: {"aliases": ["policy based challenge rule list rules spec"], "body_bytes": 7293, "body_sha256": "sha256:885ae464489403e106924f36c1fd45bb6247047b2f94d145110c1430bbc84117", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_asn", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_client", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_ip", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:arg_matchers", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_list", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_matcher", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:body_matcher", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:client_selector", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:disable_challenge", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:domain_matcher", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_captcha_challenge", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_javascript_challenge", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:http_method", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_prefix_list", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:path", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:tls_fingerprint_matcher"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules", "path": "documentation/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-023.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec"], "schema_version": 1, "sections": [{"aliases": ["policy based challenge rule list rules spec any asn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_asn", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "any_asn"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec any client"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_client", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "any_client"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec any ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "any_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec arg matchers"], "anchor": "section", "description": "A list of predicates for all POST args that need to be matched. The criteria for matching each arg are described in individual instances of ArgMatcherType. The actual arg values are extracted from the request API as a list of strings for each arg selector name. Note that all specified arg matcher predicates must", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:arg_matchers", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "arg_matchers"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec asn list"], "anchor": "section", "description": "An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "asn_list"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec asn matcher"], "anchor": "section", "description": "Match any AS number contained in the list of bgp_asn_sets.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "asn_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec body matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:body_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "body_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec client selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:client_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "client_selector"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec cookie matchers"], "anchor": "section", "description": "A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "cookie_matchers"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec disable challenge"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:disable_challenge", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "disable_challenge"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec domain matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:domain_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "domain_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec enable captcha challenge"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_captcha_challenge", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "enable_captcha_challenge"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec enable javascript challenge"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_javascript_challenge", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "enable_javascript_challenge"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec expiration timestamp"], "anchor": "schema-policy_based_challenge--rule_list--rules--spec--expiration_timestamp", "description": "The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "expiration_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["policy based challenge rule list rules spec headers"], "anchor": "section", "description": "A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "headers"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec http method", "succeeded", "success", "successful"], "anchor": "section", "description": "A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:http_method", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "http_method"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec ip matcher"], "anchor": "section", "description": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "ip_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec ip prefix list"], "anchor": "section", "description": "List of IP Prefix strings to match against.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_prefix_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "ip_prefix_list"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec path", "succeeded", "success", "successful"], "anchor": "section", "description": "A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of path prefixes, a list of exact path values and a list of regular expressions.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:path", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "path"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec query params"], "anchor": "section", "description": "A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "query_params"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec tls fingerprint matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positive match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:tls_fingerprint_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "tls_fingerprint_matcher"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "A Challenge Rule consists of an unordered list of predicates and an action. The predicates are evaluated against a set of input fields that are extracted from or derived from an L7 request API. A request API is considered to match the rule if all predicates in the rule evaluate to true for that request. Any predicates", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules.spec

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/)
- [policy_based_challenge.rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/)
- [policy_based_challenge.rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/)
- policy_based_challenge.rule_list.rules.spec

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

A Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for that
request. Any predicates that are not specified in a rule are implicitly considered to be true. If a
request API matches a challenge rule, the configured challenge is enforced.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-challenge_action": "[\"disable_challenge\",\"enable_captcha_challenge\",\"enable_javascript_challenge\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"tls_fingerprint_matcher\"]"
}
```

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/any_asn/): complete subsection reference.

- [any_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/any_client/): complete subsection reference.

- [any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/any_ip/): complete subsection reference.

- [arg_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/arg_matchers/): complete subsection reference.

- [asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/asn_list/): complete subsection reference.

- [asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/asn_matcher/): complete subsection reference.

- [body_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/body_matcher/): complete subsection reference.

- [client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/client_selector/): complete subsection reference.

- [cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/cookie_matchers/): complete subsection reference.

- [disable_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/disable_challenge/): complete subsection reference.

- [domain_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/domain_matcher/): complete subsection reference.

- [enable_captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/enable_captcha_challenge/): complete subsection reference.

- [enable_javascript_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/enable_javascript_challenge/): complete subsection reference.

<a id="schema-policy_based_challenge--rule_list--rules--spec--expiration_timestamp"></a>

### expiration_timestamp property

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/headers/): complete subsection reference.

- [http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/http_method/): complete subsection reference.

- [ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/ip_matcher/): complete subsection reference.

- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/ip_prefix_list/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/path/): complete subsection reference.

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/query_params/): complete subsection reference.

- [tls_fingerprint_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/tls_fingerprint_matcher/): complete subsection reference.
