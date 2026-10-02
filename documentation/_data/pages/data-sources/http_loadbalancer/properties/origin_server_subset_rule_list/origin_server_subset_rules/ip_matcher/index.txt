---
page_title: "origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher"
subcategory: "Load Balancing"
description: "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true."
xcsh_docs: {"aliases": ["backend servers", "origin server subset rule list origin server subset rules ip matcher", "origin servers", "upstream servers"], "body_bytes": 2689, "body_sha256": "sha256:5b2cfe052d5033986b5972ddabcd8437ee7515c799ec291317f7f95b913df550", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher:prefix_sets"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "path": "documentation/data-sources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/ip_matcher/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1303030032032131-1112131301012021-1211031323120231-1112220312122233-1330203201231213-0133101222312300-0200032222021123-2031211213323002", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-021.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "ip_matcher"], "schema_version": 1, "sections": [{"aliases": ["invert matcher"], "anchor": "schema-origin_server_subset_rule_list--origin_server_subset_rules--ip_matcher--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "ip_matcher", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["prefix sets"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher:prefix_sets", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "ip_matcher", "prefix_sets"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/ip_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [origin_server_subset_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/origin_server_subset_rule_list/)
- [origin_server_subset_rule_list.origin_server_subset_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher

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

<a id="schema-origin_server_subset_rule_list--origin_server_subset_rules--ip_matcher--invert_matcher"></a>

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

- [prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/ip_matcher/prefix_sets/): complete subsection reference.

## Next pages

- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/ip_matcher/prefix_sets/)
- [origin_server_subset_rule_list.origin_server_subset_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
