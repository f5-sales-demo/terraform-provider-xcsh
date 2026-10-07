---
page_title: "rule_list.rules.spec"
subcategory: "Security"
description: "Shape of service_policy_rule in the storage backend."
xcsh_docs: {"aliases": ["rule list rules spec"], "body_bytes": 11513, "body_sha256": "sha256:0acbf9f6b6837b4457a4030158b7849e1bacd8f73c50dd86b824d4ac16f77914", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:any_asn", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:any_client", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:any_ip", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:api_group_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:arg_matchers", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:asn_list", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:asn_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:body_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:bot_action", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:client_name_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:client_selector", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:cookie_matchers", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:domain_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:headers", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:http_method", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ip_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ip_prefix_list", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ip_threat_category_list", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ja4_tls_fingerprint", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:jwt_claims", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:label_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:mum_action", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:path", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:port_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:query_params", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:segment_policy", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:tls_fingerprint_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:user_identity_matcher", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec", "parent_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules", "path": "documentation/data-sources/service_policy/properties/rule_list/rules/spec/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210", "registry_path": "docs/guides/data-sources--service_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec"], "schema_version": 1, "sections": [{"aliases": ["rule list rules spec action"], "anchor": "schema-rule_list--rules--spec--action", "description": "The rule action determines the disposition of the input request API. If a policy matches a rule with an ALLOW action, the processing of the request proceeds forward. If it matches a rule with a DENY action, the processing of the request is terminated and an appropriate message/code returned to the originator. If it", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "action"], "syntax": "attribute", "type": "string"}, {"aliases": ["rule list rules spec any asn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:any_asn", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "any_asn"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec any client"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:any_client", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "any_client"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec any ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:any_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "any_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec api group matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies a list of values for matching an input string. The match is considered successful if the input value is present in the list. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:api_group_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "api_group_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec arg matchers"], "anchor": "section", "description": "A list of predicates for all POST args that need to be matched. The criteria for matching each arg are described in individual instances of ArgMatcherType. The actual arg values are extracted from the request API as a list of strings for each arg selector name. Note that all specified arg matcher predicates must", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:arg_matchers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "arg_matchers"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec asn list"], "anchor": "section", "description": "An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:asn_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "asn_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec asn matcher"], "anchor": "section", "description": "Match any AS number contained in the list of bgp_asn_sets.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:asn_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "asn_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec body matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:body_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "body_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec bot action"], "anchor": "section", "description": "Modify Bot protection behavior for a matching request. The modification could be to entirely skip Bot processing.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:bot_action", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "bot_action"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec client name"], "anchor": "schema-rule_list--rules--spec--client_name", "description": "Exclusive with The expected name of the client invoking the request API. The predicate evaluates to true if any of the actual names is the same as the expected client name.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "client_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["rule list rules spec client name matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:client_name_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "client_name_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec client selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:client_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "client_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec cookie matchers"], "anchor": "section", "description": "A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:cookie_matchers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "cookie_matchers"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec domain matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:domain_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "domain_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec expiration timestamp"], "anchor": "schema-rule_list--rules--spec--expiration_timestamp", "description": "The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "expiration_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["rule list rules spec headers"], "anchor": "section", "description": "A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:headers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "headers"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec http method", "succeeded", "success", "successful"], "anchor": "section", "description": "A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:http_method", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "http_method"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec ip matcher"], "anchor": "section", "description": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ip_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "ip_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec ip prefix list"], "anchor": "section", "description": "List of IP Prefix strings to match against.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ip_prefix_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "ip_prefix_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec ip threat category list"], "anchor": "section", "description": "List of IP threat categories.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ip_threat_category_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "ip_threat_category_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec ja4 tls fingerprint"], "anchor": "section", "description": "An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a different structure and length.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ja4_tls_fingerprint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "ja4_tls_fingerprint"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec jwt claims"], "anchor": "section", "description": "A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:jwt_claims", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "jwt_claims"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec label matcher"], "anchor": "section", "description": "A label matcher specifies a list of label keys whose values need to match for source/client and destination/server. Note that the actual label values are not specified and do not matter. This allows an ability to scope grouping by the label key name.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:label_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "label_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec log rule evaluation"], "anchor": "schema-rule_list--rules--spec--log_rule_evaluation", "description": "Log the rule match details along with the request and continue to evaluate rules in the sequence.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "log_rule_evaluation"], "syntax": "attribute", "type": "bool"}, {"aliases": ["rule list rules spec mum action"], "anchor": "section", "description": "Modify behavior for a matching request. The modification could be to entirely skip processing.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:mum_action", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "mum_action"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec path", "succeeded", "success", "successful"], "anchor": "section", "description": "A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of path prefixes, a list of exact path values and a list of regular expressions.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:path", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "path"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec port matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A port matcher specifies a list of port ranges as match criteria. The match is considered successful if the input port falls within any of the port ranges. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:port_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "port_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec query params"], "anchor": "section", "description": "A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:query_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "query_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec request constraints"], "anchor": "section", "description": "Configuration parameter for request constraints.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "request_constraints"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec segment policy"], "anchor": "section", "description": "Configure source and destination segment for policy.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:segment_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "segment_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec tls fingerprint matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positive match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:tls_fingerprint_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "tls_fingerprint_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec user identity matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:user_identity_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "user_identity_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec waf action"], "anchor": "section", "description": "Modify App Firewall behavior for a matching request. The modification could either be to entirely skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall Rule Control settings.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "waf_action"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/rules/spec/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Shape of service_policy_rule in the storage backend.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["service_policyCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/)
- rule_list.rules.spec

<a id="section"></a>

Type: `"single"`. Computed.

Shape of service\_policy\_rule in the storage backend.

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
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_name\",\"client_name_matcher\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-dst_asn_choice": "[]",
  "x-ves-oneof-field-dst_ip_choice": "[]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"ja4_tls_fingerprint\",\"tls_fingerprint_matcher\"]"
}
```

## Direct properties

<a id="schema-rule_list--rules--spec--action"></a>

### action property

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW|NEXT\_POLICY\] The rule action determines the disposition of the input request
API. If a policy matches a rule with an ALLOW action, the processing of the request proceeds
forward. If it matches a rule with a DENY action, the processing of the request is terminated and an
appropriate message/code returned to.. Possible values are \`DENY\`, \`ALLOW\`, \`NEXT\_POLICY\`.
Defaults to \`DENY\`.

Additional upstream details:

The rule action determines the disposition of the input request API. If it matches a rule with a
DENY action, the processing of the request is terminated and an appropriate message/code returned to
the originator. If it matches a rule with a NEXT\_POLICY\_SET action, evaluation of the current
policy set terminates and evaluation of the next policy set in the chain begins.

&#8203;- DENY: DENY

Deny the request. &#8203;- ALLOW: ALLOW

Allow the request to proceed. &#8203;- NEXT\_POLICY\_SET: NEXT\_POLICY\_SET

Terminate evaluation of the current policy set and begin evaluating the next policy set in the
chain. Note that the evaluation of any remaining policies in the current policy set is skipped.
&#8203;- NEXT\_POLICY: NEXT\_POLICY

Terminate evaluation of the current policy and begin evaluating the next policy in the policy set.
Note that the evaluation of any remaining rules in the current policy is skipped. &#8203;-
LAST\_POLICY: LAST\_POLICY

Terminate evaluation of the current policy and begin evaluating the last policy in the policy set.
Note that the evaluation of any remaining rules in the current policy is skipped. &#8203;-
GOTO\_POLICY: GOTO\_POLICY

Terminate evaluation of the current policy and begin evaluating a specific policy in the policy set.
The policy is specified using the goto\_policy field in the rule and must be after the current
policy in the policy set.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW",
    "NEXT_POLICY"
  ],
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  }
}
```

- [any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/any_asn/): complete subsection reference.

- [any_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/any_client/): complete subsection reference.

- [any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/any_ip/): complete subsection reference.

- [api_group_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/api_group_matcher/): complete subsection reference.

- [arg_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/arg_matchers/): complete subsection reference.

- [asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/asn_list/): complete subsection reference.

- [asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/asn_matcher/): complete subsection reference.

- [body_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/body_matcher/): complete subsection reference.

- [bot_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/bot_action/): complete subsection reference.

<a id="schema-rule_list--rules--spec--client_name"></a>

### client_name property

Type: `"string"`. Computed.

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [client_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/client_name_matcher/): complete subsection reference.

- [client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/client_selector/): complete subsection reference.

- [cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/cookie_matchers/): complete subsection reference.

- [domain_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/domain_matcher/): complete subsection reference.

<a id="schema-rule_list--rules--spec--expiration_timestamp"></a>

### expiration_timestamp property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/headers/): complete subsection reference.

- [http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/http_method/): complete subsection reference.

- [ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/ip_matcher/): complete subsection reference.

- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/ip_prefix_list/): complete subsection reference.

- [ip_threat_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/ip_threat_category_list/): complete subsection reference.

- [ja4_tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/ja4_tls_fingerprint/): complete subsection reference.

- [jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/jwt_claims/): complete subsection reference.

- [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/label_matcher/): complete subsection reference.

<a id="schema-rule_list--rules--spec--log_rule_evaluation"></a>

### log_rule_evaluation property

Type: `"bool"`. Computed.

Log the rule match details along with the request and continue to evaluate rules in the sequence.

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

- [mum_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/mum_action/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/path/): complete subsection reference.

- [port_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/port_matcher/): complete subsection reference.

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/query_params/): complete subsection reference.

- [request_constraints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/request_constraints/): complete subsection reference.

- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/segment_policy/): complete subsection reference.

- [tls_fingerprint_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/tls_fingerprint_matcher/): complete subsection reference.

- [user_identity_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/user_identity_matcher/): complete subsection reference.

- [waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/): complete subsection reference.
