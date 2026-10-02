---
page_title: "policy_based_challenge.rule_list.rules.spec.ip_matcher"
subcategory: "Load Balancing"
description: "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true."
xcsh_docs: {"aliases": ["policy based challenge rule list rules spec ip matcher"], "body_bytes": 2875, "body_sha256": "sha256:fb97937f2c2c2c140df71579045bf05179721e0494212ebaf75a29d7f7c07a63", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher:prefix_sets"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "path": "documentation/data-sources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/ip_matcher/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0201030211031131-1111310113132001-1132122322201233-0010012333222220-2001233311231202-3210131212113212-2302321102233223-0202323121313221", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-013.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "ip_matcher"], "schema_version": 1, "sections": [{"aliases": ["invert matcher"], "anchor": "schema-policy_based_challenge--rule_list--rules--spec--ip_matcher--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "ip_matcher", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["prefix sets"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher:prefix_sets", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "ip_matcher", "prefix_sets"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/ip_matcher/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules.spec.ip_matcher

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/)
- [policy_based_challenge.rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/)
- [policy_based_challenge.rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/)
- [policy_based_challenge.rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/)
- policy_based_challenge.rule_list.rules.spec.ip_matcher

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

<a id="schema-policy_based_challenge--rule_list--rules--spec--ip_matcher--invert_matcher"></a>

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

- [prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/ip_matcher/prefix_sets/): complete subsection reference.

## Next pages

- [policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/ip_matcher/prefix_sets/)
- [policy_based_challenge.rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
