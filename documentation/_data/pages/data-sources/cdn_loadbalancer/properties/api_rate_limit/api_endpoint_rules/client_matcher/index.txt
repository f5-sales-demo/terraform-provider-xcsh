---
page_title: "api_rate_limit.api_endpoint_rules.client_matcher"
subcategory: "Load Balancing"
description: "Client conditions for matching a rule."
xcsh_docs: {"aliases": ["api rate limit api endpoint rules client matcher"], "body_bytes": 5635, "body_sha256": "sha256:c05b8c8b3af2a50844974c812ebb2bcc269a30020476a1b2447f2718e7661b6d", "capabilities": ["cdn", "security.rate-limiting"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:any_client", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:any_ip", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:asn_list", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:asn_matcher", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:client_selector", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:ip_matcher", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:ip_prefix_list", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:ip_threat_category_list", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:tls_fingerprint_matcher"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "path": "documentation/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2333000313200333-1002311130232303-1003320210120320-3303132032032002-0222032111022122-0213320132220311-3231101100200213-1132222033312120", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher"], "schema_version": 1, "sections": [{"aliases": ["api rate limit api endpoint rules client matcher any client"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:any_client", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher", "any_client"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit api endpoint rules client matcher any ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:any_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher", "any_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit api endpoint rules client matcher asn list"], "anchor": "section", "description": "An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:asn_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher", "asn_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit api endpoint rules client matcher asn matcher"], "anchor": "section", "description": "Match any AS number contained in the list of bgp_asn_sets.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:asn_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher", "asn_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit api endpoint rules client matcher client selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:client_selector", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher", "client_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit api endpoint rules client matcher ip matcher"], "anchor": "section", "description": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:ip_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher", "ip_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit api endpoint rules client matcher ip prefix list"], "anchor": "section", "description": "List of IP Prefix strings to match against.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:ip_prefix_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher", "ip_prefix_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit api endpoint rules client matcher ip threat category list"], "anchor": "section", "description": "List of IP threat categories.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:ip_threat_category_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher", "ip_threat_category_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit api endpoint rules client matcher tls fingerprint matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positive match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:tls_fingerprint_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher", "tls_fingerprint_matcher"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Client conditions for matching a rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.api_endpoint_rules.client_matcher

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/)
- api_rate_limit.api_endpoint_rules.client_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

## Direct properties

- [any_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/any_client/): complete subsection reference.

- [any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/any_ip/): complete subsection reference.

- [asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/asn_list/): complete subsection reference.

- [asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/asn_matcher/): complete subsection reference.

- [client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/client_selector/): complete subsection reference.

- [ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/ip_matcher/): complete subsection reference.

- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/ip_prefix_list/): complete subsection reference.

- [ip_threat_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/ip_threat_category_list/): complete subsection reference.

- [tls_fingerprint_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/tls_fingerprint_matcher/): complete subsection reference.

## Next pages

- [api_rate_limit.api_endpoint_rules.client_matcher.any_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/any_client/)
- [api_rate_limit.api_endpoint_rules.client_matcher.any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/any_ip/)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/asn_list/)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/asn_matcher/)
- [api_rate_limit.api_endpoint_rules.client_matcher.client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/client_selector/)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/ip_matcher/)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/ip_prefix_list/)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/ip_threat_category_list/)
- [api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/tls_fingerprint_matcher/)
- [api_rate_limit.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
