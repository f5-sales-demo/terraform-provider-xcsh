---
page_title: "rules"
subcategory: ""
description: "An ordered list of rules that are evaluated sequentially against the input fields extracted from an API request in order to determine a user identifier. Evaluation of the rules is terminated once a user identifier has been extracted."
xcsh_docs: {"aliases": ["rules"], "body_bytes": 20199, "body_sha256": "sha256:e1ec093f910f055ff7097cc680e0b3243958b0a2f794398879ea0dffc11215c1", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:user_identification:properties:rules:client_asn", "xcsh-docs:resources:user_identification:properties:rules:client_city", "xcsh-docs:resources:user_identification:properties:rules:client_country", "xcsh-docs:resources:user_identification:properties:rules:client_ip", "xcsh-docs:resources:user_identification:properties:rules:client_region", "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "xcsh-docs:resources:user_identification:properties:rules:none", "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:resources:user_identification:properties:rules", "parent_id": "xcsh-docs:resources:user_identification:reference", "path": "documentation/resources/user_identification/properties/rules/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1012211320301032-3320202022331132-2112102302331031-2123331232332232-2322320133010011-1301021031230003-2202201001032002-1132320133000203", "registry_path": "docs/guides/resources--user_identification--reference--group-001.md", "relationships": [{"anchor": "schema-rules--cookie_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,cookie_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--cookie_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,cookie_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--cookie_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,cookie_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--cookie_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,cookie_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--cookie_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,cookie_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--cookie_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--cookie_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,ip_and_http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--cookie_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--cookie_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--cookie_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--cookie_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--cookie_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--cookie_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--cookie_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,ip_and_http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--ip_and_http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,ip_and_http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--ip_and_http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,ip_and_http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--ip_and_http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,ip_and_http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--ip_and_http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,ip_and_http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--ip_and_http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,ip_and_http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--ip_and_http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,ip_and_http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--ip_and_http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,ip_and_http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--ip_and_http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_http_header_name,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--ip_and_http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_http_header_name,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--ip_and_http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_http_header_name,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--ip_and_http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_http_header_name,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--ip_and_http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_http_header_name,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--ip_and_http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_http_header_name,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--ip_and_http_header_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_http_header_name,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--jwt_claim_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--jwt_claim_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--jwt_claim_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--jwt_claim_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--jwt_claim_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--jwt_claim_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--jwt_claim_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--jwt_claim_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_http_header_name,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--jwt_claim_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_ja4_tls_fingerprint,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--jwt_claim_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_tls_fingerprint,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--jwt_claim_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ja4_tls_fingerprint,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--jwt_claim_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:jwt_claim_name,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--jwt_claim_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:jwt_claim_name,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--jwt_claim_name", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:jwt_claim_name,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--query_param_key", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--query_param_key", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--query_param_key", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--query_param_key", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--query_param_key", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--query_param_key", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--query_param_key", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--query_param_key", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_http_header_name,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--query_param_key", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_ja4_tls_fingerprint,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--query_param_key", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_tls_fingerprint,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--query_param_key", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ja4_tls_fingerprint,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--query_param_key", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:jwt_claim_name,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--query_param_key", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:none,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "schema-rules--query_param_key", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:query_param_key,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,client_city", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,client_country", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,client_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,client_region", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,cookie_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,ip_and_http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,client_city", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,client_country", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,client_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,client_region", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,cookie_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,ip_and_http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,client_country", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,client_country", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,client_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,client_region", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,cookie_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,ip_and_http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,client_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,client_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,client_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,client_region", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,cookie_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,ip_and_http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,client_region", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,client_region", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,client_region", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,client_region", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,cookie_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,ip_and_http_header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_http_header_name,ip_and_ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_ja4_tls_fingerprint,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_ja4_tls_fingerprint,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_ja4_tls_fingerprint,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_ja4_tls_fingerprint,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_ja4_tls_fingerprint,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_ja4_tls_fingerprint,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_http_header_name,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_ja4_tls_fingerprint,ip_and_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_tls_fingerprint,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_tls_fingerprint,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_tls_fingerprint,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_tls_fingerprint,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_tls_fingerprint,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_http_header_name,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_ja4_tls_fingerprint,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_tls_fingerprint,ja4_tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ja4_tls_fingerprint,jwt_claim_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ja4_tls_fingerprint,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ja4_tls_fingerprint,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ja4_tls_fingerprint,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_http_header_name,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_ja4_tls_fingerprint,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_tls_fingerprint,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ja4_tls_fingerprint,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:jwt_claim_name,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:none,query_param_key", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:none,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_asn,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_city,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_country,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_ip,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:client_region,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cookie_name,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:http_header_name,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_http_header_name,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_ja4_tls_fingerprint,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ip_and_tls_fingerprint,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:ja4_tls_fingerprint,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:jwt_claim_name,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:none,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:query_param_key,tls_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules"], "schema_version": 1, "sections": [{"aliases": ["rules client asn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:user_identification:properties:rules:client_asn", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "client_asn"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules client city"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:user_identification:properties:rules:client_city", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "client_city"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules client country"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:user_identification:properties:rules:client_country", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "client_country"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules client ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:user_identification:properties:rules:client_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "client_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules client region"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "client_region"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules cookie name"], "anchor": "schema-rules--cookie_name", "description": "Exclusive with Use the HTTP cookie value for the given name as user identifier.", "document_id": "xcsh-docs:resources:user_identification:properties:rules", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "cookie_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules http header name"], "anchor": "schema-rules--http_header_name", "description": "Exclusive with Use the HTTP header value for the given name as user identifier.", "document_id": "xcsh-docs:resources:user_identification:properties:rules", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "http_header_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules ip and http header name"], "anchor": "schema-rules--ip_and_http_header_name", "description": "Exclusive with Name of HTTP header from which the value should be extracted.", "document_id": "xcsh-docs:resources:user_identification:properties:rules", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ip_and_http_header_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules ip and ja4 tls fingerprint"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ip_and_ja4_tls_fingerprint"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules ip and tls fingerprint"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ip_and_tls_fingerprint"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules ja4 tls fingerprint"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ja4_tls_fingerprint"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules jwt claim name"], "anchor": "schema-rules--jwt_claim_name", "description": "Exclusive with Use the JWT claim value as user identifier.", "document_id": "xcsh-docs:resources:user_identification:properties:rules", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "jwt_claim_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:user_identification:properties:rules:none", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "none"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules query param key"], "anchor": "schema-rules--query_param_key", "description": "Exclusive with Use the query parameter value for the given key as user identifier.", "document_id": "xcsh-docs:resources:user_identification:properties:rules", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "query_param_key"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules tls fingerprint"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "tls_fingerprint"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/properties/rules/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "An ordered list of rules that are evaluated sequentially against the input fields extracted from an API request in order to determine a user identifier. Evaluation of the rules is terminated once a user identifier has been extracted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["user_identificationCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/)
- rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

An ordered list of rules that are evaluated sequentially against the input fields extracted from an
API request in order to determine a user identifier. Evaluation of the rules is terminated once a
user identifier has been extracted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("client_asn",
    "client_city"),
  validators.ConflictingListObjectAttributes("client_asn",
    "client_country"),
  validators.ConflictingListObjectAttributes("client_asn",
    "client_ip"),
  validators.ConflictingListObjectAttributes("client_asn",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_asn",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_asn",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "none"),
  validators.ConflictingListObjectAttributes("client_asn",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_asn",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "client_country"),
  validators.ConflictingListObjectAttributes("client_city",
    "client_ip"),
  validators.ConflictingListObjectAttributes("client_city",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_city",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "none"),
  validators.ConflictingListObjectAttributes("client_city",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_city",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "client_ip"),
  validators.ConflictingListObjectAttributes("client_country",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_country",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "none"),
  validators.ConflictingListObjectAttributes("client_country",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_country",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_ip",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "none"),
  validators.ConflictingListObjectAttributes("client_ip",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_ip",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "none"),
  validators.ConflictingListObjectAttributes("client_region",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_region",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "none"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "none"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "none"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "none"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "none"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "none"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("jwt_claim_name",
    "none"),
  validators.ConflictingListObjectAttributes("jwt_claim_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("jwt_claim_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("none",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("none",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("query_param_key",
    "tls_fingerprint")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [client_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/client_asn/): complete subsection reference.

- [client_city](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/client_city/): complete subsection reference.

- [client_country](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/client_country/): complete subsection reference.

- [client_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/client_ip/): complete subsection reference.

- [client_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/client_region/): complete subsection reference.

<a id="schema-rules--cookie_name"></a>

### cookie_name property

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key
tls\_fingerprint\] Use the HTTP cookie value for the given name as user identifier.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="schema-rules--http_header_name"></a>

### http_header_name property

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint
ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Use the HTTP header
value for the given name as user identifier.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="schema-rules--ip_and_http_header_name"></a>

### ip_and_http_header_name property

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint
jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Name of HTTP header from which the value
should be extracted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [ip_and_ja4_tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/ip_and_ja4_tls_fingerprint/): complete subsection reference.

- [ip_and_tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/ip_and_tls_fingerprint/): complete subsection reference.

- [ja4_tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/ja4_tls_fingerprint/): complete subsection reference.

<a id="schema-rules--jwt_claim_name"></a>

### jwt_claim_name property

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint none query\_param\_key tls\_fingerprint\] Use the
JWT claim value as user identifier.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/none/): complete subsection reference.

<a id="schema-rules--query_param_key"></a>

### query_param_key property

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none tls\_fingerprint\] Use the
query parameter value for the given key as user identifier.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/tls_fingerprint/): complete subsection reference.
