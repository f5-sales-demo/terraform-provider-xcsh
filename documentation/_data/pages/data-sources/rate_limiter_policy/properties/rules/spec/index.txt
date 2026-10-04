---
page_title: "rules.spec"
subcategory: "Security"
description: "Shape of Rate Limiter Rule."
xcsh_docs: {"aliases": ["rules spec"], "body_bytes": 6723, "body_sha256": "sha256:7b92dab9a1c0ba7f8df5d8580b8061f2a2a9898075805109e0f4d2debac443e2", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:any_asn", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:any_country", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:any_ip", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:apply_rate_limiter", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:asn_list", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:asn_matcher", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:bypass_rate_limiter", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:country_list", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:custom_rate_limiter", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:domain_matcher", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:headers", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:http_method", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:ip_matcher", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:ip_prefix_list", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:path", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec", "parent_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules", "path": "documentation/data-sources/rate_limiter_policy/properties/rules/spec/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323", "registry_path": "docs/guides/data-sources--rate_limiter_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "spec"], "schema_version": 1, "sections": [{"aliases": ["rules spec any asn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:any_asn", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "any_asn"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec any country"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:any_country", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "any_country"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec any ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:any_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "any_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec apply rate limiter"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:apply_rate_limiter", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "apply_rate_limiter"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec asn list"], "anchor": "section", "description": "An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:asn_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "spec", "asn_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec asn matcher"], "anchor": "section", "description": "Match any AS number contained in the list of bgp_asn_sets.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:asn_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "spec", "asn_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec bypass rate limiter"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:bypass_rate_limiter", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "bypass_rate_limiter"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec country list"], "anchor": "section", "description": "List of Country Codes to match against.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:country_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "spec", "country_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec custom rate limiter"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:custom_rate_limiter", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "spec", "custom_rate_limiter"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec domain matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:domain_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "spec", "domain_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec headers"], "anchor": "section", "description": "A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "spec", "headers"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec http method", "succeeded", "success", "successful"], "anchor": "section", "description": "A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:http_method", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "spec", "http_method"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec ip matcher"], "anchor": "section", "description": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:ip_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "spec", "ip_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec ip prefix list"], "anchor": "section", "description": "List of IP Prefix strings to match against.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:ip_prefix_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "spec", "ip_prefix_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec path", "succeeded", "success", "successful"], "anchor": "section", "description": "A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of path prefixes, a list of exact path values and a list of regular expressions.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "spec", "path"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules spec segment policy"], "anchor": "section", "description": "Configure source and destination segment for policy.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "spec", "segment_policy"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter_policy/properties/rules/spec/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Shape of Rate Limiter Rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/)
- rules.spec

<a id="section"></a>

Type: `"single"`. Computed.

Rate Limiter Rule Specification. Shape of Rate Limiter Rule.

Upstream description:

Shape of Rate Limiter Rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"apply_rate_limiter\",\"bypass_rate_limiter\",\"custom_rate_limiter\"]",
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-country_choice": "[\"any_country\",\"country_list\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

## Direct properties

- [any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/any_asn/): complete subsection reference.

- [any_country](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/any_country/): complete subsection reference.

- [any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/any_ip/): complete subsection reference.

- [apply_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/apply_rate_limiter/): complete subsection reference.

- [asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/asn_list/): complete subsection reference.

- [asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/asn_matcher/): complete subsection reference.

- [bypass_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/bypass_rate_limiter/): complete subsection reference.

- [country_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/country_list/): complete subsection reference.

- [custom_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/custom_rate_limiter/): complete subsection reference.

- [domain_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/domain_matcher/): complete subsection reference.

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/headers/): complete subsection reference.

- [http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/http_method/): complete subsection reference.

- [ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/ip_matcher/): complete subsection reference.

- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/ip_prefix_list/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/path/): complete subsection reference.

- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/): complete subsection reference.

## Next pages

- [rules.spec.any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/any_asn/)
- [rules.spec.any_country](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/any_country/)
- [rules.spec.any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/any_ip/)
- [rules.spec.apply_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/apply_rate_limiter/)
- [rules.spec.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/asn_list/)
- [rules.spec.asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/asn_matcher/)
- [rules.spec.bypass_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/bypass_rate_limiter/)
- [rules.spec.country_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/country_list/)
- [rules.spec.custom_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/custom_rate_limiter/)
- [rules.spec.domain_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/domain_matcher/)
- [rules.spec.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/headers/)
- [rules.spec.http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/http_method/)
- [rules.spec.ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/ip_matcher/)
- [rules.spec.ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/ip_prefix_list/)
- [rules.spec.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/path/)
- [rules.spec.segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/)
- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
