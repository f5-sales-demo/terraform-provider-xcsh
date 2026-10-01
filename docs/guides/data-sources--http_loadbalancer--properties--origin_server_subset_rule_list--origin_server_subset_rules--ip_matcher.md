---
page_title: "origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher"
subcategory: "Load Balancing"
description: "origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2287, "body_sha256": "sha256:73d7cb35d564f39a3bc77297f442cd0aa2f04edf355ff297c5e90f612e21e204", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher:prefix_sets"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "path": "docs/guides/data-sources--http_loadbalancer--properties--origin_server_subset_rule_list--origin_server_subset_rules--ip_matcher.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "ip_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/ip_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--properties--origin_server_subset_rule_list.md)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--properties--origin_server_subset_rule_list--origin_server_subset_rules.md)
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

- [prefix_sets](data-sources--http_loadbalancer--properties--origin_server_subset_rule_list--origin_server_subset_rules--ip_matcher--prefix_sets.md): complete subsection reference.

## Next pages

- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets](data-sources--http_loadbalancer--properties--origin_server_subset_rule_list--origin_server_subset_rules--ip_matcher--prefix_sets.md)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--properties--origin_server_subset_rule_list--origin_server_subset_rules.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
