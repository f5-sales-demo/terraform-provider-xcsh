---
page_title: "api_protection_rules.api_groups_rules.client_matcher.ip_matcher"
subcategory: "Load Balancing"
description: "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true."
xcsh_docs: {"aliases": ["api protection rules api groups rules client matcher ip matcher"], "body_bytes": 2800, "body_sha256": "sha256:357228c0a51e12d713796c72be077f2237dabba25260ff573c590c8a6fd1ca9c", "capabilities": ["load-balancing", "security.api-protection"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_matcher:prefix_sets"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_matcher", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher", "path": "documentation/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/ip_matcher/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3230103330333033-3322012100123200-1033002123321202-3021031102133112-0101210323133312-3330202200011213-1333010330031101-0333000310020020", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_protection_rules", "api_groups_rules", "client_matcher", "ip_matcher"], "schema_version": 1, "sections": [{"aliases": ["api protection rules api groups rules client matcher ip matcher invert matcher"], "anchor": "schema-api_protection_rules--api_groups_rules--client_matcher--ip_matcher--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "client_matcher", "ip_matcher", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["api protection rules api groups rules client matcher ip matcher prefix sets"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_matcher:prefix_sets", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "client_matcher", "ip_matcher", "prefix_sets"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/ip_matcher/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_groups_rules.client_matcher.ip_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/)
- [api_protection_rules.api_groups_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/)
- [api_protection_rules.api_groups_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/)
- api_protection_rules.api_groups_rules.client_matcher.ip_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

## Direct properties

<a id="schema-api_protection_rules--api_groups_rules--client_matcher--ip_matcher--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

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

- [prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/ip_matcher/prefix_sets/): complete subsection reference.

## Next pages

- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/ip_matcher/prefix_sets/)
- [api_protection_rules.api_groups_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
