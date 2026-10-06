---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_service_policy_rule."
xcsh_docs: {"aliases": ["service policy rule"], "body_bytes": 58701, "body_sha256": "sha256:c03907143f3a82af2f189dc5e510883771ea4b6ec5b046e4c8a34ac9b2ddd8db", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:any_asn", "xcsh-docs:resources:service_policy_rule:properties:any_client", "xcsh-docs:resources:service_policy_rule:properties:any_ip", "xcsh-docs:resources:service_policy_rule:properties:api_group_matcher", "xcsh-docs:resources:service_policy_rule:properties:arg_matchers", "xcsh-docs:resources:service_policy_rule:properties:asn_list", "xcsh-docs:resources:service_policy_rule:properties:asn_matcher", "xcsh-docs:resources:service_policy_rule:properties:body_matcher", "xcsh-docs:resources:service_policy_rule:properties:bot_action", "xcsh-docs:resources:service_policy_rule:properties:client_name_matcher", "xcsh-docs:resources:service_policy_rule:properties:client_selector", "xcsh-docs:resources:service_policy_rule:properties:cookie_matchers", "xcsh-docs:resources:service_policy_rule:properties:domain_matcher", "xcsh-docs:resources:service_policy_rule:properties:headers", "xcsh-docs:resources:service_policy_rule:properties:http_method", "xcsh-docs:resources:service_policy_rule:properties:ip_matcher", "xcsh-docs:resources:service_policy_rule:properties:ip_prefix_list", "xcsh-docs:resources:service_policy_rule:properties:ip_threat_category_list", "xcsh-docs:resources:service_policy_rule:properties:ja4_tls_fingerprint", "xcsh-docs:resources:service_policy_rule:properties:jwt_claims", "xcsh-docs:resources:service_policy_rule:properties:label_matcher", "xcsh-docs:resources:service_policy_rule:properties:mum_action", "xcsh-docs:resources:service_policy_rule:properties:path", "xcsh-docs:resources:service_policy_rule:properties:port_matcher", "xcsh-docs:resources:service_policy_rule:properties:query_params", "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "xcsh-docs:resources:service_policy_rule:properties:segment_policy", "xcsh-docs:resources:service_policy_rule:properties:timeouts", "xcsh-docs:resources:service_policy_rule:properties:tls_fingerprint_matcher", "xcsh-docs:resources:service_policy_rule:properties:waf_action"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:reference", "parent_id": "xcsh-docs:resources:service_policy_rule:fundamentals", "path": "documentation/resources/service_policy_rule/properties/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["action"], "anchor": "schema-action", "description": "The rule action determines the disposition of the input request API. If a policy matches a rule with an ALLOW action, the processing of the request proceeds forward. If it matches a rule with a DENY action, the processing of the request is terminated and an appropriate message/code returned to the originator. If it", "document_id": "xcsh-docs:resources:service_policy_rule:reference", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ALLOW", "DENY", "NEXT_POLICY"], "version": 1}], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["action"], "syntax": "attribute", "type": "string"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:service_policy_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["any asn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:any_asn", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["any_asn"], "syntax": "attribute", "type": "object"}, {"aliases": ["any client"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:any_client", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["any_client"], "syntax": "attribute", "type": "object"}, {"aliases": ["any ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:any_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["any_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["api group matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies a list of values for matching an input string. The match is considered successful if the input value is present in the list. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:api_group_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_group_matcher--match", "enforcement": "provider-schema", "group": "api_group_matcher:RequiredObjectAttributes:match", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:api_group_matcher", "type": "requires"}], "schema_path": ["api_group_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["arg matchers"], "anchor": "section", "description": "A list of predicates for all POST args that need to be matched. The criteria for matching each arg are described in individual instances of ArgMatcherType. The actual arg values are extracted from the request API as a list of strings for each arg selector name. Note that all specified arg matcher predicates must", "document_id": "xcsh-docs:resources:service_policy_rule:properties:arg_matchers", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "arg_matchers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:arg_matchers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "arg_matchers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:arg_matchers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "arg_matchers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:arg_matchers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "arg_matchers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:arg_matchers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "arg_matchers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:arg_matchers:item", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "arg_matchers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:arg_matchers:item", "type": "conflicts"}, {"anchor": "schema-arg_matchers--name", "enforcement": "provider-schema", "group": "arg_matchers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:arg_matchers", "type": "requires"}], "schema_path": ["arg_matchers"], "syntax": "block", "type": "object"}, {"aliases": ["asn list"], "anchor": "section", "description": "An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:asn_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-asn_list--as_numbers", "enforcement": "provider-schema", "group": "asn_list:RequiredObjectAttributes:as_numbers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:asn_list", "type": "requires"}], "schema_path": ["asn_list"], "syntax": "block", "type": "object"}, {"aliases": ["asn matcher"], "anchor": "section", "description": "Match any AS number contained in the list of bgp_asn_sets.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:asn_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "asn_matcher:RequiredObjectAttributes:asn_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:asn_matcher:asn_sets", "type": "requires"}], "schema_path": ["asn_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["body matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:body_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["body_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["bot action"], "anchor": "section", "description": "Modify Bot protection behavior for a matching request. The modification could be to entirely skip Bot processing.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:bot_action", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_action:ConflictingObjectAttributes:bot_skip_processing,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:bot_action:bot_skip_processing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_action:ConflictingObjectAttributes:bot_skip_processing,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:bot_action:none", "type": "conflicts"}], "schema_path": ["bot_action"], "syntax": "block", "type": "object"}, {"aliases": ["client name"], "anchor": "schema-client_name", "description": "Exclusive with The expected name of the client invoking the request API. The predicate evaluates to true if any of the actual names is the same as the expected client name.", "document_id": "xcsh-docs:resources:service_policy_rule:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["client name matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:client_name_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["client_name_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["client selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:client_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-client_selector--expressions", "enforcement": "provider-schema", "group": "client_selector:RequiredObjectAttributes:expressions", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:client_selector", "type": "requires"}], "schema_path": ["client_selector"], "syntax": "block", "type": "object"}, {"aliases": ["cookie matchers"], "anchor": "section", "description": "A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must", "document_id": "xcsh-docs:resources:service_policy_rule:properties:cookie_matchers", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cookie_matchers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:cookie_matchers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cookie_matchers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:cookie_matchers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cookie_matchers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:cookie_matchers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cookie_matchers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:cookie_matchers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cookie_matchers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:cookie_matchers:item", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cookie_matchers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:cookie_matchers:item", "type": "conflicts"}, {"anchor": "schema-cookie_matchers--name", "enforcement": "provider-schema", "group": "cookie_matchers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:cookie_matchers", "type": "requires"}], "schema_path": ["cookie_matchers"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:service_policy_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:service_policy_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["domain matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:domain_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domain_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["expiration timestamp"], "anchor": "schema-expiration_timestamp", "description": "The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore.", "document_id": "xcsh-docs:resources:service_policy_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expiration_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["headers"], "anchor": "section", "description": "A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header", "document_id": "xcsh-docs:resources:service_policy_rule:properties:headers", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "headers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:headers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "headers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:headers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "headers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:headers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "headers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:headers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "headers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:headers:item", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "headers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:headers:item", "type": "conflicts"}, {"anchor": "schema-headers--name", "enforcement": "provider-schema", "group": "headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:headers", "type": "requires"}], "schema_path": ["headers"], "syntax": "block", "type": "object"}, {"aliases": ["http method", "succeeded", "success", "successful"], "anchor": "section", "description": "A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:http_method", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_method"], "syntax": "block", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:service_policy_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["ip matcher"], "anchor": "section", "description": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:ip_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ip_matcher:RequiredObjectAttributes:prefix_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:ip_matcher:prefix_sets", "type": "requires"}], "schema_path": ["ip_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["ip prefix list"], "anchor": "section", "description": "List of IP Prefix strings to match against.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:ip_prefix_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ip_prefix_list"], "syntax": "block", "type": "object"}, {"aliases": ["ip threat category list"], "anchor": "section", "description": "List of IP threat categories.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:ip_threat_category_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ip_threat_category_list--ip_threat_categories", "enforcement": "provider-schema", "group": "ip_threat_category_list:RequiredObjectAttributes:ip_threat_categories", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:ip_threat_category_list", "type": "requires"}], "schema_path": ["ip_threat_category_list"], "syntax": "block", "type": "object"}, {"aliases": ["ja4 tls fingerprint"], "anchor": "section", "description": "An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a different structure and length.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:ja4_tls_fingerprint", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ja4_tls_fingerprint"], "syntax": "block", "type": "object"}, {"aliases": ["jwt claims"], "anchor": "section", "description": "A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:jwt_claims", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "jwt_claims:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:jwt_claims:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_claims:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:jwt_claims:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_claims:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:jwt_claims:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_claims:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:jwt_claims:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_claims:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:jwt_claims:item", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_claims:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:jwt_claims:item", "type": "conflicts"}, {"anchor": "schema-jwt_claims--name", "enforcement": "provider-schema", "group": "jwt_claims:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:jwt_claims", "type": "requires"}], "schema_path": ["jwt_claims"], "syntax": "block", "type": "object"}, {"aliases": ["label matcher"], "anchor": "section", "description": "A label matcher specifies a list of label keys whose values need to match for source/client and destination/server. Note that the actual label values are not specified and do not matter. This allows an ability to scope grouping by the label key name.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:label_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["label_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:service_policy_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["log rule evaluation"], "anchor": "schema-log_rule_evaluation", "description": "Log the rule match details along with the request and continue to evaluate rules in the sequence.", "document_id": "xcsh-docs:resources:service_policy_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["log_rule_evaluation"], "syntax": "attribute", "type": "bool"}, {"aliases": ["mum action"], "anchor": "section", "description": "Modify behavior for a matching request. The modification could be to entirely skip processing.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:mum_action", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "mum_action:ConflictingObjectAttributes:default,skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:mum_action:default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "mum_action:ConflictingObjectAttributes:default,skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:mum_action:skip_processing", "type": "conflicts"}], "schema_path": ["mum_action"], "syntax": "block", "type": "object"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:service_policy_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:service_policy_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["path", "succeeded", "success", "successful"], "anchor": "section", "description": "A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of path prefixes, a list of exact path values and a list of regular expressions.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:path", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["path"], "syntax": "block", "type": "object"}, {"aliases": ["port matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A port matcher specifies a list of port ranges as match criteria. The match is considered successful if the input port falls within any of the port ranges. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:port_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-port_matcher--ports", "enforcement": "provider-schema", "group": "port_matcher:RequiredObjectAttributes:ports", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:port_matcher", "type": "requires"}], "schema_path": ["port_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["query params"], "anchor": "section", "description": "A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that", "document_id": "xcsh-docs:resources:service_policy_rule:properties:query_params", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "query_params:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:query_params:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "query_params:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:query_params:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "query_params:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:query_params:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "query_params:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:query_params:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "query_params:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:query_params:item", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "query_params:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:query_params:item", "type": "conflicts"}, {"anchor": "schema-query_params--key", "enforcement": "provider-schema", "group": "query_params:RequiredListObjectAttributes:key", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:query_params", "type": "requires"}], "schema_path": ["query_params"], "syntax": "block", "type": "object"}, {"aliases": ["request constraints"], "anchor": "section", "description": "Configuration parameter for request constraints.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-request_constraints--max_cookie_count_exceeds", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_cookie_count_exceeds,max_cookie_count_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "type": "conflicts"}, {"anchor": "schema-request_constraints--max_cookie_key_size_exceeds", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_cookie_key_size_exceeds,max_cookie_key_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "type": "conflicts"}, {"anchor": "schema-request_constraints--max_cookie_value_size_exceeds", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_cookie_value_size_exceeds,max_cookie_value_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "type": "conflicts"}, {"anchor": "schema-request_constraints--max_header_count_exceeds", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_header_count_exceeds,max_header_count_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "type": "conflicts"}, {"anchor": "schema-request_constraints--max_header_key_size_exceeds", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_header_key_size_exceeds,max_header_key_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "type": "conflicts"}, {"anchor": "schema-request_constraints--max_header_value_size_exceeds", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_header_value_size_exceeds,max_header_value_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "type": "conflicts"}, {"anchor": "schema-request_constraints--max_parameter_count_exceeds", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_parameter_count_exceeds,max_parameter_count_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "type": "conflicts"}, {"anchor": "schema-request_constraints--max_parameter_name_size_exceeds", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_parameter_name_size_exceeds,max_parameter_name_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "type": "conflicts"}, {"anchor": "schema-request_constraints--max_parameter_value_size_exceeds", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_parameter_value_size_exceeds,max_parameter_value_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "type": "conflicts"}, {"anchor": "schema-request_constraints--max_query_size_exceeds", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_query_size_exceeds,max_query_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "type": "conflicts"}, {"anchor": "schema-request_constraints--max_request_line_size_exceeds", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_request_line_size_exceeds,max_request_line_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "type": "conflicts"}, {"anchor": "schema-request_constraints--max_request_size_exceeds", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_request_size_exceeds,max_request_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "type": "conflicts"}, {"anchor": "schema-request_constraints--max_url_size_exceeds", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_url_size_exceeds,max_url_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_cookie_count_exceeds,max_cookie_count_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_cookie_count_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_cookie_key_size_exceeds,max_cookie_key_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_cookie_key_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_cookie_value_size_exceeds,max_cookie_value_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_cookie_value_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_header_count_exceeds,max_header_count_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_header_count_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_header_key_size_exceeds,max_header_key_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_header_key_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_header_value_size_exceeds,max_header_value_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_header_value_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_parameter_count_exceeds,max_parameter_count_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_parameter_count_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_parameter_name_size_exceeds,max_parameter_name_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_parameter_name_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_parameter_value_size_exceeds,max_parameter_value_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_parameter_value_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_query_size_exceeds,max_query_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_query_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_request_line_size_exceeds,max_request_line_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_request_line_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_request_size_exceeds,max_request_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_request_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_constraints:ConflictingObjectAttributes:max_url_size_exceeds,max_url_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_url_size_none", "type": "conflicts"}], "schema_path": ["request_constraints"], "syntax": "block", "type": "object"}, {"aliases": ["segment policy"], "anchor": "section", "description": "Configure source and destination segment for policy.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:dst_any,dst_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:dst_any,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:dst_any,dst_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_segments", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:dst_segments,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_segments", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:dst_any,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:intra_segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:dst_segments,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:intra_segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:src_any,src_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy:ConflictingObjectAttributes:src_any,src_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments", "type": "conflicts"}], "schema_path": ["segment_policy"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:service_policy_rule:properties:timeouts", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["succeeded", "success", "successful", "tls fingerprint matcher"], "anchor": "section", "description": "A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positive match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input", "document_id": "xcsh-docs:resources:service_policy_rule:properties:tls_fingerprint_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_fingerprint_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["waf action"], "anchor": "section", "description": "Modify App Firewall behavior for a matching request. The modification could either be to entirely skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall Rule Control settings.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "waf_action:ConflictingObjectAttributes:app_firewall_detection_control,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_action:ConflictingObjectAttributes:app_firewall_detection_control,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_action:ConflictingObjectAttributes:app_firewall_detection_control,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_action:ConflictingObjectAttributes:none,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_action:ConflictingObjectAttributes:app_firewall_detection_control,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:waf_skip_processing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_action:ConflictingObjectAttributes:none,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:waf_skip_processing", "type": "conflicts"}], "schema_path": ["waf_action"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_service_policy_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- Property reference

## Direct properties

<a id="schema-action"></a>

### action property

Type: `"string"`. Required.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALLOW","DENY","NEXT_POLICY"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW",
    "NEXT_POLICY"),
}
```

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

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/any_asn/): complete subsection reference.

- [any_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/any_client/): complete subsection reference.

- [any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/any_ip/): complete subsection reference.

- [api_group_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/api_group_matcher/): complete subsection reference.

- [arg_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/arg_matchers/): complete subsection reference.

- [asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_list/): complete subsection reference.

- [asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_matcher/): complete subsection reference.

- [body_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/body_matcher/): complete subsection reference.

- [bot_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/bot_action/): complete subsection reference.

<a id="schema-client_name"></a>

### client_name property

Type: `"string"`. Optional, Computed.

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [client_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/client_name_matcher/): complete subsection reference.

- [client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/client_selector/): complete subsection reference.

- [cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/cookie_matchers/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

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

- [domain_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/domain_matcher/): complete subsection reference.

<a id="schema-expiration_timestamp"></a>

### expiration_timestamp property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/headers/): complete subsection reference.

- [http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/http_method/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_matcher/): complete subsection reference.

- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_prefix_list/): complete subsection reference.

- [ip_threat_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_threat_category_list/): complete subsection reference.

- [ja4_tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ja4_tls_fingerprint/): complete subsection reference.

- [jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/jwt_claims/): complete subsection reference.

- [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/label_matcher/): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-log_rule_evaluation"></a>

### log_rule_evaluation property

Type: `"bool"`. Optional, Computed.

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

- [mum_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/mum_action/): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Service Policy Rule. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the Service Policy Rule is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/path/): complete subsection reference.

- [port_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/port_matcher/): complete subsection reference.

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/query_params/): complete subsection reference.

- [request_constraints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/): complete subsection reference.

- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/timeouts/): complete subsection reference.

- [tls_fingerprint_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/tls_fingerprint_matcher/): complete subsection reference.

- [waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/#schema-action) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/#schema-annotations) |
| `any_asn` | [any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/any_asn/#section) |
| `any_client` | [any_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/any_client/#section) |
| `any_ip` | [any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/any_ip/#section) |
| `api_group_matcher` | [api_group_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/api_group_matcher/#section) |
| `api_group_matcher.invert_matcher` | [api_group_matcher.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/api_group_matcher/#schema-api_group_matcher--invert_matcher) |
| `api_group_matcher.match` | [api_group_matcher.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/api_group_matcher/#schema-api_group_matcher--match) |
| `arg_matchers` | [arg_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/arg_matchers/#section) |
| `arg_matchers.check_not_present` | [arg_matchers.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/arg_matchers/check_not_present/#section) |
| `arg_matchers.check_present` | [arg_matchers.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/arg_matchers/check_present/#section) |
| `arg_matchers.invert_matcher` | [arg_matchers.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/arg_matchers/#schema-arg_matchers--invert_matcher) |
| `arg_matchers.item` | [arg_matchers.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/arg_matchers/item/#section) |
| `arg_matchers.item.exact_values` | [arg_matchers.item.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/arg_matchers/item/#schema-arg_matchers--item--exact_values) |
| `arg_matchers.item.regex_values` | [arg_matchers.item.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/arg_matchers/item/#schema-arg_matchers--item--regex_values) |
| `arg_matchers.item.transformers` | [arg_matchers.item.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/arg_matchers/item/#schema-arg_matchers--item--transformers) |
| `arg_matchers.name` | [arg_matchers.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/arg_matchers/#schema-arg_matchers--name) |
| `asn_list` | [asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_list/#section) |
| `asn_list.as_numbers` | [asn_list.as_numbers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_list/#schema-asn_list--as_numbers) |
| `asn_matcher` | [asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_matcher/#section) |
| `asn_matcher.asn_sets` | [asn_matcher.asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_matcher/asn_sets/#section) |
| `asn_matcher.asn_sets.kind` | [asn_matcher.asn_sets.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_matcher/asn_sets/#schema-asn_matcher--asn_sets--kind) |
| `asn_matcher.asn_sets.name` | [asn_matcher.asn_sets.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_matcher/asn_sets/#schema-asn_matcher--asn_sets--name) |
| `asn_matcher.asn_sets.namespace` | [asn_matcher.asn_sets.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_matcher/asn_sets/#schema-asn_matcher--asn_sets--namespace) |
| `asn_matcher.asn_sets.tenant` | [asn_matcher.asn_sets.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_matcher/asn_sets/#schema-asn_matcher--asn_sets--tenant) |
| `asn_matcher.asn_sets.uid` | [asn_matcher.asn_sets.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_matcher/asn_sets/#schema-asn_matcher--asn_sets--uid) |
| `body_matcher` | [body_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/body_matcher/#section) |
| `body_matcher.exact_values` | [body_matcher.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/body_matcher/#schema-body_matcher--exact_values) |
| `body_matcher.regex_values` | [body_matcher.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/body_matcher/#schema-body_matcher--regex_values) |
| `body_matcher.transformers` | [body_matcher.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/body_matcher/#schema-body_matcher--transformers) |
| `bot_action` | [bot_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/bot_action/#section) |
| `bot_action.bot_skip_processing` | [bot_action.bot_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/bot_action/bot_skip_processing/#section) |
| `bot_action.none` | [bot_action.none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/bot_action/none/#section) |
| `client_name` | [client_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/#schema-client_name) |
| `client_name_matcher` | [client_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/client_name_matcher/#section) |
| `client_name_matcher.exact_values` | [client_name_matcher.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/client_name_matcher/#schema-client_name_matcher--exact_values) |
| `client_name_matcher.regex_values` | [client_name_matcher.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/client_name_matcher/#schema-client_name_matcher--regex_values) |
| `client_selector` | [client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/client_selector/#section) |
| `client_selector.expressions` | [client_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/client_selector/#schema-client_selector--expressions) |
| `cookie_matchers` | [cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/cookie_matchers/#section) |
| `cookie_matchers.check_not_present` | [cookie_matchers.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/cookie_matchers/check_not_present/#section) |
| `cookie_matchers.check_present` | [cookie_matchers.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/cookie_matchers/check_present/#section) |
| `cookie_matchers.invert_matcher` | [cookie_matchers.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/cookie_matchers/#schema-cookie_matchers--invert_matcher) |
| `cookie_matchers.item` | [cookie_matchers.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/cookie_matchers/item/#section) |
| `cookie_matchers.item.exact_values` | [cookie_matchers.item.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/cookie_matchers/item/#schema-cookie_matchers--item--exact_values) |
| `cookie_matchers.item.regex_values` | [cookie_matchers.item.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/cookie_matchers/item/#schema-cookie_matchers--item--regex_values) |
| `cookie_matchers.item.transformers` | [cookie_matchers.item.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/cookie_matchers/item/#schema-cookie_matchers--item--transformers) |
| `cookie_matchers.name` | [cookie_matchers.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/cookie_matchers/#schema-cookie_matchers--name) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/#schema-disable) |
| `domain_matcher` | [domain_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/domain_matcher/#section) |
| `domain_matcher.exact_values` | [domain_matcher.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/domain_matcher/#schema-domain_matcher--exact_values) |
| `domain_matcher.regex_values` | [domain_matcher.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/domain_matcher/#schema-domain_matcher--regex_values) |
| `expiration_timestamp` | [expiration_timestamp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/#schema-expiration_timestamp) |
| `headers` | [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/headers/#section) |
| `headers.check_not_present` | [headers.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/headers/check_not_present/#section) |
| `headers.check_present` | [headers.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/headers/check_present/#section) |
| `headers.invert_matcher` | [headers.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/headers/#schema-headers--invert_matcher) |
| `headers.item` | [headers.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/headers/item/#section) |
| `headers.item.exact_values` | [headers.item.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/headers/item/#schema-headers--item--exact_values) |
| `headers.item.regex_values` | [headers.item.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/headers/item/#schema-headers--item--regex_values) |
| `headers.item.transformers` | [headers.item.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/headers/item/#schema-headers--item--transformers) |
| `headers.name` | [headers.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/headers/#schema-headers--name) |
| `http_method` | [http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/http_method/#section) |
| `http_method.invert_matcher` | [http_method.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/http_method/#schema-http_method--invert_matcher) |
| `http_method.methods` | [http_method.methods](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/http_method/#schema-http_method--methods) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/#schema-id) |
| `ip_matcher` | [ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_matcher/#section) |
| `ip_matcher.invert_matcher` | [ip_matcher.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_matcher/#schema-ip_matcher--invert_matcher) |
| `ip_matcher.prefix_sets` | [ip_matcher.prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_matcher/prefix_sets/#section) |
| `ip_matcher.prefix_sets.kind` | [ip_matcher.prefix_sets.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_matcher/prefix_sets/#schema-ip_matcher--prefix_sets--kind) |
| `ip_matcher.prefix_sets.name` | [ip_matcher.prefix_sets.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_matcher/prefix_sets/#schema-ip_matcher--prefix_sets--name) |
| `ip_matcher.prefix_sets.namespace` | [ip_matcher.prefix_sets.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_matcher/prefix_sets/#schema-ip_matcher--prefix_sets--namespace) |
| `ip_matcher.prefix_sets.tenant` | [ip_matcher.prefix_sets.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_matcher/prefix_sets/#schema-ip_matcher--prefix_sets--tenant) |
| `ip_matcher.prefix_sets.uid` | [ip_matcher.prefix_sets.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_matcher/prefix_sets/#schema-ip_matcher--prefix_sets--uid) |
| `ip_prefix_list` | [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_prefix_list/#section) |
| `ip_prefix_list.invert_match` | [ip_prefix_list.invert_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_prefix_list/#schema-ip_prefix_list--invert_match) |
| `ip_prefix_list.ip_prefixes` | [ip_prefix_list.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_prefix_list/#schema-ip_prefix_list--ip_prefixes) |
| `ip_threat_category_list` | [ip_threat_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_threat_category_list/#section) |
| `ip_threat_category_list.ip_threat_categories` | [ip_threat_category_list.ip_threat_categories](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_threat_category_list/#schema-ip_threat_category_list--ip_threat_categories) |
| `ja4_tls_fingerprint` | [ja4_tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ja4_tls_fingerprint/#section) |
| `ja4_tls_fingerprint.exact_values` | [ja4_tls_fingerprint.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ja4_tls_fingerprint/#schema-ja4_tls_fingerprint--exact_values) |
| `jwt_claims` | [jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/jwt_claims/#section) |
| `jwt_claims.check_not_present` | [jwt_claims.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/jwt_claims/check_not_present/#section) |
| `jwt_claims.check_present` | [jwt_claims.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/jwt_claims/check_present/#section) |
| `jwt_claims.invert_matcher` | [jwt_claims.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/jwt_claims/#schema-jwt_claims--invert_matcher) |
| `jwt_claims.item` | [jwt_claims.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/jwt_claims/item/#section) |
| `jwt_claims.item.exact_values` | [jwt_claims.item.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/jwt_claims/item/#schema-jwt_claims--item--exact_values) |
| `jwt_claims.item.regex_values` | [jwt_claims.item.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/jwt_claims/item/#schema-jwt_claims--item--regex_values) |
| `jwt_claims.item.transformers` | [jwt_claims.item.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/jwt_claims/item/#schema-jwt_claims--item--transformers) |
| `jwt_claims.name` | [jwt_claims.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/jwt_claims/#schema-jwt_claims--name) |
| `label_matcher` | [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/label_matcher/#section) |
| `label_matcher.keys` | [label_matcher.keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/label_matcher/#schema-label_matcher--keys) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/#schema-labels) |
| `log_rule_evaluation` | [log_rule_evaluation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/#schema-log_rule_evaluation) |
| `mum_action` | [mum_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/mum_action/#section) |
| `mum_action.default` | [mum_action.default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/mum_action/default/#section) |
| `mum_action.skip_processing` | [mum_action.skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/mum_action/skip_processing/#section) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/#schema-namespace) |
| `path` | [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/path/#section) |
| `path.encoded_path_matcher` | [path.encoded_path_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/path/#schema-path--encoded_path_matcher) |
| `path.exact_values` | [path.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/path/#schema-path--exact_values) |
| `path.invert_matcher` | [path.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/path/#schema-path--invert_matcher) |
| `path.prefix_values` | [path.prefix_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/path/#schema-path--prefix_values) |
| `path.regex_values` | [path.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/path/#schema-path--regex_values) |
| `path.suffix_values` | [path.suffix_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/path/#schema-path--suffix_values) |
| `path.transformers` | [path.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/path/#schema-path--transformers) |
| `port_matcher` | [port_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/port_matcher/#section) |
| `port_matcher.invert_matcher` | [port_matcher.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/port_matcher/#schema-port_matcher--invert_matcher) |
| `port_matcher.ports` | [port_matcher.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/port_matcher/#schema-port_matcher--ports) |
| `query_params` | [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/query_params/#section) |
| `query_params.check_not_present` | [query_params.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/query_params/check_not_present/#section) |
| `query_params.check_present` | [query_params.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/query_params/check_present/#section) |
| `query_params.invert_matcher` | [query_params.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/query_params/#schema-query_params--invert_matcher) |
| `query_params.item` | [query_params.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/query_params/item/#section) |
| `query_params.item.exact_values` | [query_params.item.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/query_params/item/#schema-query_params--item--exact_values) |
| `query_params.item.regex_values` | [query_params.item.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/query_params/item/#schema-query_params--item--regex_values) |
| `query_params.item.transformers` | [query_params.item.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/query_params/item/#schema-query_params--item--transformers) |
| `query_params.key` | [query_params.key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/query_params/#schema-query_params--key) |
| `request_constraints` | [request_constraints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/#section) |
| `request_constraints.max_cookie_count_exceeds` | [request_constraints.max_cookie_count_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/#schema-request_constraints--max_cookie_count_exceeds) |
| `request_constraints.max_cookie_count_none` | [request_constraints.max_cookie_count_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_cookie_count_none/#section) |
| `request_constraints.max_cookie_key_size_exceeds` | [request_constraints.max_cookie_key_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/#schema-request_constraints--max_cookie_key_size_exceeds) |
| `request_constraints.max_cookie_key_size_none` | [request_constraints.max_cookie_key_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_cookie_key_size_none/#section) |
| `request_constraints.max_cookie_value_size_exceeds` | [request_constraints.max_cookie_value_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/#schema-request_constraints--max_cookie_value_size_exceeds) |
| `request_constraints.max_cookie_value_size_none` | [request_constraints.max_cookie_value_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_cookie_value_size_none/#section) |
| `request_constraints.max_header_count_exceeds` | [request_constraints.max_header_count_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/#schema-request_constraints--max_header_count_exceeds) |
| `request_constraints.max_header_count_none` | [request_constraints.max_header_count_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_header_count_none/#section) |
| `request_constraints.max_header_key_size_exceeds` | [request_constraints.max_header_key_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/#schema-request_constraints--max_header_key_size_exceeds) |
| `request_constraints.max_header_key_size_none` | [request_constraints.max_header_key_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_header_key_size_none/#section) |
| `request_constraints.max_header_value_size_exceeds` | [request_constraints.max_header_value_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/#schema-request_constraints--max_header_value_size_exceeds) |
| `request_constraints.max_header_value_size_none` | [request_constraints.max_header_value_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_header_value_size_none/#section) |
| `request_constraints.max_parameter_count_exceeds` | [request_constraints.max_parameter_count_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/#schema-request_constraints--max_parameter_count_exceeds) |
| `request_constraints.max_parameter_count_none` | [request_constraints.max_parameter_count_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_parameter_count_none/#section) |
| `request_constraints.max_parameter_name_size_exceeds` | [request_constraints.max_parameter_name_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/#schema-request_constraints--max_parameter_name_size_exceeds) |
| `request_constraints.max_parameter_name_size_none` | [request_constraints.max_parameter_name_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_parameter_name_size_none/#section) |
| `request_constraints.max_parameter_value_size_exceeds` | [request_constraints.max_parameter_value_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/#schema-request_constraints--max_parameter_value_size_exceeds) |
| `request_constraints.max_parameter_value_size_none` | [request_constraints.max_parameter_value_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_parameter_value_size_none/#section) |
| `request_constraints.max_query_size_exceeds` | [request_constraints.max_query_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/#schema-request_constraints--max_query_size_exceeds) |
| `request_constraints.max_query_size_none` | [request_constraints.max_query_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_query_size_none/#section) |
| `request_constraints.max_request_line_size_exceeds` | [request_constraints.max_request_line_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/#schema-request_constraints--max_request_line_size_exceeds) |
| `request_constraints.max_request_line_size_none` | [request_constraints.max_request_line_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_request_line_size_none/#section) |
| `request_constraints.max_request_size_exceeds` | [request_constraints.max_request_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/#schema-request_constraints--max_request_size_exceeds) |
| `request_constraints.max_request_size_none` | [request_constraints.max_request_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_request_size_none/#section) |
| `request_constraints.max_url_size_exceeds` | [request_constraints.max_url_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/#schema-request_constraints--max_url_size_exceeds) |
| `request_constraints.max_url_size_none` | [request_constraints.max_url_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_url_size_none/#section) |
| `segment_policy` | [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/#section) |
| `segment_policy.dst_any` | [segment_policy.dst_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/dst_any/#section) |
| `segment_policy.dst_segments` | [segment_policy.dst_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/dst_segments/#section) |
| `segment_policy.dst_segments.segments` | [segment_policy.dst_segments.segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/dst_segments/segments/#section) |
| `segment_policy.dst_segments.segments.name` | [segment_policy.dst_segments.segments.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/dst_segments/segments/#schema-segment_policy--dst_segments--segments--name) |
| `segment_policy.dst_segments.segments.namespace` | [segment_policy.dst_segments.segments.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/dst_segments/segments/#schema-segment_policy--dst_segments--segments--namespace) |
| `segment_policy.dst_segments.segments.tenant` | [segment_policy.dst_segments.segments.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/dst_segments/segments/#schema-segment_policy--dst_segments--segments--tenant) |
| `segment_policy.intra_segment` | [segment_policy.intra_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/intra_segment/#section) |
| `segment_policy.src_any` | [segment_policy.src_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/src_any/#section) |
| `segment_policy.src_segments` | [segment_policy.src_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/src_segments/#section) |
| `segment_policy.src_segments.segments` | [segment_policy.src_segments.segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/src_segments/segments/#section) |
| `segment_policy.src_segments.segments.name` | [segment_policy.src_segments.segments.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/src_segments/segments/#schema-segment_policy--src_segments--segments--name) |
| `segment_policy.src_segments.segments.namespace` | [segment_policy.src_segments.segments.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/src_segments/segments/#schema-segment_policy--src_segments--segments--namespace) |
| `segment_policy.src_segments.segments.tenant` | [segment_policy.src_segments.segments.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/src_segments/segments/#schema-segment_policy--src_segments--segments--tenant) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/timeouts/#schema-timeouts--update) |
| `tls_fingerprint_matcher` | [tls_fingerprint_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/tls_fingerprint_matcher/#section) |
| `tls_fingerprint_matcher.classes` | [tls_fingerprint_matcher.classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/tls_fingerprint_matcher/#schema-tls_fingerprint_matcher--classes) |
| `tls_fingerprint_matcher.exact_values` | [tls_fingerprint_matcher.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/tls_fingerprint_matcher/#schema-tls_fingerprint_matcher--exact_values) |
| `tls_fingerprint_matcher.excluded_values` | [tls_fingerprint_matcher.excluded_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/tls_fingerprint_matcher/#schema-tls_fingerprint_matcher--excluded_values) |
| `waf_action` | [waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/#section) |
| `waf_action.app_firewall_detection_control` | [waf_action.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/#section) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/#section) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/#schema-waf_action--app_firewall_detection_control--exclude_attack_type_contexts--context) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/#schema-waf_action--app_firewall_detection_control--exclude_attack_type_contexts--context_name) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/#schema-waf_action--app_firewall_detection_control--exclude_attack_type_contexts--exclude_attack_type) |
| `waf_action.app_firewall_detection_control.exclude_bot_name_contexts` | [waf_action.app_firewall_detection_control.exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/#section) |
| `waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/#schema-waf_action--app_firewall_detection_control--exclude_bot_name_contexts--bot_name) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts` | [waf_action.app_firewall_detection_control.exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_signature_contexts/#section) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.context` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.context](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_signature_contexts/#schema-waf_action--app_firewall_detection_control--exclude_signature_contexts--context) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_signature_contexts/#schema-waf_action--app_firewall_detection_control--exclude_signature_contexts--context_name) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_signature_contexts/#schema-waf_action--app_firewall_detection_control--exclude_signature_contexts--signature_id) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts` | [waf_action.app_firewall_detection_control.exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_violation_contexts/#section) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.context` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.context](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_violation_contexts/#schema-waf_action--app_firewall_detection_control--exclude_violation_contexts--context) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_violation_contexts/#schema-waf_action--app_firewall_detection_control--exclude_violation_contexts--context_name) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_violation_contexts/#schema-waf_action--app_firewall_detection_control--exclude_violation_contexts--exclude_violation) |
| `waf_action.none` | [waf_action.none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/none/#section) |
| `waf_action.waf_skip_processing` | [waf_action.waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/waf_skip_processing/#section) |
