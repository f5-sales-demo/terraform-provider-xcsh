---
page_title: "origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher"
subcategory: "Load Balancing"
description: "origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1630, "body_sha256": "sha256:8bbc0e922383b3b01076512c19ea78ad68e26b723f43fb4db971ea28da2ef099", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_matcher", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_matcher:asn_sets"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_matcher", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "path": "docs/guides/data-sources--http_loadbalancer--properties--origin_server_subset_rule_list--origin_server_subset_rules--asn_matcher.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "asn_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/asn_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--properties--origin_server_subset_rule_list.md)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--properties--origin_server_subset_rule_list--origin_server_subset_rules.md)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

- [asn_sets](data-sources--http_loadbalancer--properties--origin_server_subset_rule_list--origin_server_subset_rules--asn_matcher--asn_sets.md): complete subsection reference.

## Next pages

- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets](data-sources--http_loadbalancer--properties--origin_server_subset_rule_list--origin_server_subset_rules--asn_matcher--asn_sets.md)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--properties--origin_server_subset_rule_list--origin_server_subset_rules.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
