---
page_title: "rule_list.rules.spec"
subcategory: "Security"
description: "Shape of service_policy_rule in the storage backend."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "rule list rules spec", "upstream servers"], "body_bytes": 19103, "body_sha256": "sha256:403c3e9beb436e05192a30d54eb1bd920dc8c2e6230df9aa22ec2406aeabf6b4", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_asn", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_client", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_ip", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:api_group_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:asn_list", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:asn_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:body_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:bot_action", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_name_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_selector", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:cookie_matchers", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:domain_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:headers", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:http_method", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_prefix_list", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_threat_category_list", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ja4_tls_fingerprint", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:jwt_claims", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:label_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:path", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:port_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:query_params", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:tls_fingerprint_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:user_identity_matcher", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules", "path": "documentation/resources/service_policy/properties/rule_list/rules/spec/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212", "registry_path": "docs/guides/resources--service_policy--reference--group-001.md", "relationships": [{"anchor": "schema-rule_list--rules--spec--client_name", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_client,client_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--client_name", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:client_name,client_name_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--client_name", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:client_name,client_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--client_name", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:client_name,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_asn,asn_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_asn,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_client,client_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_client,client_name_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_client,client_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_client,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_asn,asn_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_asn,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_client,client_name_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_name_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:client_name,client_name_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_name_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:client_name_matcher,client_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_name_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:client_name_matcher,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_name_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_client,client_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:client_name,client_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:client_name_matcher,client_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:client_selector,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:any_client,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_threat_category_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:client_name,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_threat_category_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:client_name_matcher,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_threat_category_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:client_selector,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_threat_category_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:ja4_tls_fingerprint,tls_fingerprint_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:ConflictingObjectAttributes:ja4_tls_fingerprint,tls_fingerprint_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:tls_fingerprint_matcher", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--action", "enforcement": "provider-schema", "group": "rule_list.rules.spec:RequiredObjectAttributes:action,waf_action", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec:RequiredObjectAttributes:action,waf_action", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec"], "schema_version": 1, "sections": [{"aliases": ["action"], "anchor": "schema-rule_list--rules--spec--action", "description": "The rule action determines the disposition of the input request API. If a policy matches a rule with an ALLOW action, the processing of the request proceeds forward. If it matches a rule with a DENY action, the processing of the request is terminated and an appropriate message/code returned to the originator. If it mat", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "action"], "syntax": "attribute", "type": "string"}, {"aliases": ["any asn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_asn", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "any_asn"], "syntax": "attribute", "type": "object"}, {"aliases": ["any client"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_client", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "any_client"], "syntax": "attribute", "type": "object"}, {"aliases": ["any ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:any_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "any_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["api group matcher", "login success", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies a list of values for matching an input string. The match is considered successful if the input value is present in the list. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:api_group_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rule_list--rules--spec--api_group_matcher--match", "enforcement": "provider-schema", "group": "rule_list.rules.spec.api_group_matcher:RequiredObjectAttributes:match", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:api_group_matcher", "type": "requires"}], "schema_path": ["rule_list", "rules", "spec", "api_group_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["arg matchers"], "anchor": "section", "description": "A list of predicates for all POST args that need to be matched. The criteria for matching each arg are described in individual instances of ArgMatcherType. The actual arg values are extracted from the request API as a list of strings for each arg selector name. Note that all specified arg matcher predicates must evalua", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:arg_matchers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "arg_matchers"], "syntax": "block", "type": "object"}, {"aliases": ["asn list"], "anchor": "section", "description": "An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:asn_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rule_list--rules--spec--asn_list--as_numbers", "enforcement": "provider-schema", "group": "rule_list.rules.spec.asn_list:RequiredObjectAttributes:as_numbers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:asn_list", "type": "requires"}], "schema_path": ["rule_list", "rules", "spec", "asn_list"], "syntax": "block", "type": "object"}, {"aliases": ["asn matcher"], "anchor": "section", "description": "Match any AS number contained in the list of bgp_asn_sets.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:asn_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.asn_matcher:RequiredObjectAttributes:asn_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:asn_matcher:asn_sets", "type": "requires"}], "schema_path": ["rule_list", "rules", "spec", "asn_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["body matcher", "login success", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:body_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "body_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["bot action"], "anchor": "section", "description": "Modify Bot protection behavior for a matching request. The modification could be to entirely skip Bot processing.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:bot_action", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.bot_action:ConflictingObjectAttributes:bot_skip_processing,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:bot_action:bot_skip_processing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.bot_action:ConflictingObjectAttributes:bot_skip_processing,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:bot_action:none", "type": "conflicts"}], "schema_path": ["rule_list", "rules", "spec", "bot_action"], "syntax": "block", "type": "object"}, {"aliases": ["client name"], "anchor": "schema-rule_list--rules--spec--client_name", "description": "Exclusive with The expected name of the client invoking the request API. The predicate evaluates to true if any of the actual names is the same as the expected client name.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "client_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["client name matcher", "login success", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_name_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "client_name_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["client selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Labe", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_selector", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rule_list--rules--spec--client_selector--expressions", "enforcement": "provider-schema", "group": "rule_list.rules.spec.client_selector:RequiredObjectAttributes:expressions", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:client_selector", "type": "requires"}], "schema_path": ["rule_list", "rules", "spec", "client_selector"], "syntax": "block", "type": "object"}, {"aliases": ["cookie matchers"], "anchor": "section", "description": "A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must eva", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:cookie_matchers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "cookie_matchers"], "syntax": "block", "type": "object"}, {"aliases": ["domain matcher", "login success", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:domain_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "domain_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["expiration timestamp"], "anchor": "schema-rule_list--rules--spec--expiration_timestamp", "description": "The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "expiration_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["headers"], "anchor": "section", "description": "A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:headers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "headers"], "syntax": "block", "type": "object"}, {"aliases": ["http method", "login success", "succeeded", "success", "successful"], "anchor": "section", "description": "A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:http_method", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "http_method"], "syntax": "block", "type": "object"}, {"aliases": ["ip matcher"], "anchor": "section", "description": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.ip_matcher:RequiredObjectAttributes:prefix_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_matcher:prefix_sets", "type": "requires"}], "schema_path": ["rule_list", "rules", "spec", "ip_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["ip prefix list"], "anchor": "section", "description": "List of IP Prefix strings to match against.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_prefix_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "ip_prefix_list"], "syntax": "block", "type": "object"}, {"aliases": ["ip threat category list"], "anchor": "section", "description": "List of IP threat categories.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_threat_category_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rule_list--rules--spec--ip_threat_category_list--ip_threat_categories", "enforcement": "provider-schema", "group": "rule_list.rules.spec.ip_threat_category_list:RequiredObjectAttributes:ip_threat_categories", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_threat_category_list", "type": "requires"}], "schema_path": ["rule_list", "rules", "spec", "ip_threat_category_list"], "syntax": "block", "type": "object"}, {"aliases": ["ja4 tls fingerprint"], "anchor": "section", "description": "An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a different structure and length.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ja4_tls_fingerprint", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "ja4_tls_fingerprint"], "syntax": "block", "type": "object"}, {"aliases": ["jwt claims"], "anchor": "section", "description": "A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:jwt_claims", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "jwt_claims"], "syntax": "block", "type": "object"}, {"aliases": ["label matcher"], "anchor": "section", "description": "A label matcher specifies a list of label keys whose values need to match for source/client and destination/server. Note that the actual label values are not specified and do not matter. This allows an ability to scope grouping by the label key name.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:label_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "label_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["log rule evaluation"], "anchor": "schema-rule_list--rules--spec--log_rule_evaluation", "description": "Log the rule match details along with the request and continue to evaluate rules in the sequence.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "log_rule_evaluation"], "syntax": "attribute", "type": "bool"}, {"aliases": ["mum action"], "anchor": "section", "description": "Modify behavior for a matching request. The modification could be to entirely skip processing.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.mum_action:ConflictingObjectAttributes:default,skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action:default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.mum_action:ConflictingObjectAttributes:default,skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action:skip_processing", "type": "conflicts"}], "schema_path": ["rule_list", "rules", "spec", "mum_action"], "syntax": "block", "type": "object"}, {"aliases": ["login success", "path", "succeeded", "success", "successful"], "anchor": "section", "description": "A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of path prefixes, a list of exact path values and a list of regular expressions.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "path"], "syntax": "block", "type": "object"}, {"aliases": ["login success", "port matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A port matcher specifies a list of port ranges as match criteria. The match is considered successful if the input port falls within any of the port ranges. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:port_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rule_list--rules--spec--port_matcher--ports", "enforcement": "provider-schema", "group": "rule_list.rules.spec.port_matcher:RequiredObjectAttributes:ports", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:port_matcher", "type": "requires"}], "schema_path": ["rule_list", "rules", "spec", "port_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["query params"], "anchor": "section", "description": "A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:query_params", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "query_params"], "syntax": "block", "type": "object"}, {"aliases": ["request constraints"], "anchor": "section", "description": "Configuration parameter for request constraints.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rule_list--rules--spec--request_constraints--max_cookie_count_exceeds", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_cookie_count_exceeds,max_cookie_count_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--request_constraints--max_cookie_key_size_exceeds", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_cookie_key_size_exceeds,max_cookie_key_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--request_constraints--max_cookie_value_size_exceeds", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_cookie_value_size_exceeds,max_cookie_value_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--request_constraints--max_header_count_exceeds", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_header_count_exceeds,max_header_count_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--request_constraints--max_header_key_size_exceeds", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_header_key_size_exceeds,max_header_key_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--request_constraints--max_header_value_size_exceeds", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_header_value_size_exceeds,max_header_value_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--request_constraints--max_parameter_count_exceeds", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_parameter_count_exceeds,max_parameter_count_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--request_constraints--max_parameter_name_size_exceeds", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_parameter_name_size_exceeds,max_parameter_name_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--request_constraints--max_parameter_value_size_exceeds", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_parameter_value_size_exceeds,max_parameter_value_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--request_constraints--max_query_size_exceeds", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_query_size_exceeds,max_query_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--request_constraints--max_request_line_size_exceeds", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_request_line_size_exceeds,max_request_line_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--request_constraints--max_request_size_exceeds", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_request_size_exceeds,max_request_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--spec--request_constraints--max_url_size_exceeds", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_url_size_exceeds,max_url_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_cookie_count_exceeds,max_cookie_count_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_cookie_count_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_cookie_key_size_exceeds,max_cookie_key_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_cookie_key_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_cookie_value_size_exceeds,max_cookie_value_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_cookie_value_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_header_count_exceeds,max_header_count_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_header_count_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_header_key_size_exceeds,max_header_key_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_header_key_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_header_value_size_exceeds,max_header_value_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_header_value_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_parameter_count_exceeds,max_parameter_count_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_parameter_count_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_parameter_name_size_exceeds,max_parameter_name_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_parameter_name_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_parameter_value_size_exceeds,max_parameter_value_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_parameter_value_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_query_size_exceeds,max_query_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_query_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_request_line_size_exceeds,max_request_line_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_request_line_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_request_size_exceeds,max_request_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_request_size_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.request_constraints:ConflictingObjectAttributes:max_url_size_exceeds,max_url_size_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_url_size_none", "type": "conflicts"}], "schema_path": ["rule_list", "rules", "spec", "request_constraints"], "syntax": "block", "type": "object"}, {"aliases": ["segment policy"], "anchor": "section", "description": "Configure source and destination segment for policy.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.segment_policy:ConflictingObjectAttributes:dst_any,dst_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:dst_any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.segment_policy:ConflictingObjectAttributes:dst_any,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:dst_any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.segment_policy:ConflictingObjectAttributes:dst_any,dst_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:dst_segments", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.segment_policy:ConflictingObjectAttributes:dst_segments,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:dst_segments", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.segment_policy:ConflictingObjectAttributes:dst_any,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:intra_segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.segment_policy:ConflictingObjectAttributes:dst_segments,intra_segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:intra_segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.segment_policy:ConflictingObjectAttributes:src_any,src_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:src_any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.segment_policy:ConflictingObjectAttributes:src_any,src_segments", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:src_segments", "type": "conflicts"}], "schema_path": ["rule_list", "rules", "spec", "segment_policy"], "syntax": "block", "type": "object"}, {"aliases": ["login success", "succeeded", "success", "successful", "tls fingerprint matcher"], "anchor": "section", "description": "A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positive match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerpr", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:tls_fingerprint_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "tls_fingerprint_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["login success", "succeeded", "success", "successful", "user identity matcher"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:user_identity_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "user_identity_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["waf action"], "anchor": "section", "description": "Modify App Firewall behavior for a matching request. The modification could either be to entirely skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall Rule Control settings.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.waf_action:ConflictingObjectAttributes:app_firewall_detection_control,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.waf_action:ConflictingObjectAttributes:app_firewall_detection_control,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.waf_action:ConflictingObjectAttributes:app_firewall_detection_control,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.waf_action:ConflictingObjectAttributes:none,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.waf_action:ConflictingObjectAttributes:app_firewall_detection_control,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:waf_skip_processing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.waf_action:ConflictingObjectAttributes:none,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:waf_skip_processing", "type": "conflicts"}], "schema_path": ["rule_list", "rules", "spec", "waf_action"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Shape of service_policy_rule in the storage backend.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/)
- rule_list.rules.spec

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Shape of service\_policy\_rule in the storage backend.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("action",
    "waf_action"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_client",
    "client_name"),
  validators.ConflictingObjectAttributes("any_client",
    "client_name_matcher"),
  validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("client_name",
    "client_name_matcher"),
  validators.ConflictingObjectAttributes("client_name",
    "client_selector"),
  validators.ConflictingObjectAttributes("client_name",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("client_name_matcher",
    "client_selector"),
  validators.ConflictingObjectAttributes("client_name_matcher",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("ja4_tls_fingerprint",
    "tls_fingerprint_matcher")}
```

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

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rule_list--rules--spec--action"></a>

### action property

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW|NEXT\_POLICY\] The rule action determines the disposition of the input request
API. If a policy matches a rule with an ALLOW action, the processing of the request proceeds
forward. If it matches a rule with a DENY action, the processing of the request is terminated and an
appropriate message/code returned to.. Possible values are \`DENY\`, \`ALLOW\`, \`NEXT\_POLICY\`.
Defaults to \`DENY\`.

Upstream description:

The rule action determines the disposition of the input request API. If a policy matches a rule with
an ALLOW action, the processing of the request proceeds forward. If it matches a rule with a DENY
action, the processing of the request is terminated and an appropriate message/code returned to the
originator. If it matches a rule with a NEXT\_POLICY\_SET action, evaluation of the current policy
set terminates and evaluation of the next policy set in the chain begins.

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

- [any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_asn/): complete subsection reference.

- [any_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_client/): complete subsection reference.

- [any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_ip/): complete subsection reference.

- [api_group_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/api_group_matcher/): complete subsection reference.

- [arg_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/): complete subsection reference.

- [asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_list/): complete subsection reference.

- [asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_matcher/): complete subsection reference.

- [body_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/body_matcher/): complete subsection reference.

- [bot_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/bot_action/): complete subsection reference.

<a id="schema-rule_list--rules--spec--client_name"></a>

### client_name property

Type: `"string"`. Optional.

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

Upstream description:

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [client_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/client_name_matcher/): complete subsection reference.

- [client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/client_selector/): complete subsection reference.

- [cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/cookie_matchers/): complete subsection reference.

- [domain_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/domain_matcher/): complete subsection reference.

<a id="schema-rule_list--rules--spec--expiration_timestamp"></a>

### expiration_timestamp property

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/headers/): complete subsection reference.

- [http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/http_method/): complete subsection reference.

- [ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/): complete subsection reference.

- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_prefix_list/): complete subsection reference.

- [ip_threat_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_threat_category_list/): complete subsection reference.

- [ja4_tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ja4_tls_fingerprint/): complete subsection reference.

- [jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/jwt_claims/): complete subsection reference.

- [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/label_matcher/): complete subsection reference.

<a id="schema-rule_list--rules--spec--log_rule_evaluation"></a>

### log_rule_evaluation property

Type: `"bool"`. Optional.

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

- [mum_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/mum_action/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/path/): complete subsection reference.

- [port_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/port_matcher/): complete subsection reference.

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/query_params/): complete subsection reference.

- [request_constraints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/): complete subsection reference.

- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/): complete subsection reference.

- [tls_fingerprint_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/tls_fingerprint_matcher/): complete subsection reference.

- [user_identity_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/user_identity_matcher/): complete subsection reference.

- [waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/): complete subsection reference.

## Next pages

- [rule_list.rules.spec.any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_asn/)
- [rule_list.rules.spec.any_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_client/)
- [rule_list.rules.spec.any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_ip/)
- [rule_list.rules.spec.api_group_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/api_group_matcher/)
- [rule_list.rules.spec.arg_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/)
- [rule_list.rules.spec.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_list/)
- [rule_list.rules.spec.asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_matcher/)
- [rule_list.rules.spec.body_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/body_matcher/)
- [rule_list.rules.spec.bot_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/bot_action/)
- [rule_list.rules.spec.client_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/client_name_matcher/)
- [rule_list.rules.spec.client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/client_selector/)
- [rule_list.rules.spec.cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/cookie_matchers/)
- [rule_list.rules.spec.domain_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/domain_matcher/)
- [rule_list.rules.spec.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/headers/)
- [rule_list.rules.spec.http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/http_method/)
- [rule_list.rules.spec.ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/)
- [rule_list.rules.spec.ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_prefix_list/)
- [rule_list.rules.spec.ip_threat_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_threat_category_list/)
- [rule_list.rules.spec.ja4_tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ja4_tls_fingerprint/)
- [rule_list.rules.spec.jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/jwt_claims/)
- [rule_list.rules.spec.label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/label_matcher/)
- [rule_list.rules.spec.mum_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/mum_action/)
- [rule_list.rules.spec.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/path/)
- [rule_list.rules.spec.port_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/port_matcher/)
- [rule_list.rules.spec.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/query_params/)
- [rule_list.rules.spec.request_constraints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/)
- [rule_list.rules.spec.segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/)
- [rule_list.rules.spec.tls_fingerprint_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/tls_fingerprint_matcher/)
- [rule_list.rules.spec.user_identity_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/user_identity_matcher/)
- [rule_list.rules.spec.waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
