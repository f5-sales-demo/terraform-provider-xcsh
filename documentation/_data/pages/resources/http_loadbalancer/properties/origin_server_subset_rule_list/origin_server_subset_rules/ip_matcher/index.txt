---
page_title: "origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher"
subcategory: "Load Balancing"
description: "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true."
xcsh_docs: {"aliases": ["origin server subset rule list origin server subset rules ip matcher"], "body_bytes": 1975, "body_sha256": "sha256:8e3c6301fff42bb6b7ff961a049ab2bb4a7ddacbd34690c8623bfc6f93b17434", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher:prefix_sets"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "path": "documentation/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/ip_matcher/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-0203121202130031-2203220200231131-2001322323233203-0133212303033123-1123320021332113-2112022122112233-2331321110200131-3131031333310133", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-023.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "ip_matcher"], "schema_version": 1, "sections": [{"aliases": ["origin server subset rule list origin server subset rules ip matcher invert matcher"], "anchor": "schema-origin_server_subset_rule_list--origin_server_subset_rules--ip_matcher--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "ip_matcher", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["origin server subset rule list origin server subset rules ip matcher prefix sets"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher:prefix_sets", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "ip_matcher", "prefix_sets"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/ip_matcher/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [origin_server_subset_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/)
- [origin_server_subset_rule_list.origin_server_subset_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-origin_server_subset_rule_list--origin_server_subset_rules--ip_matcher--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

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

- [prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/ip_matcher/prefix_sets/): complete subsection reference.
