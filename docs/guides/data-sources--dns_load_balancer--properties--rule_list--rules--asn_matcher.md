---
page_title: "rule_list.rules.asn_matcher"
subcategory: "DNS"
description: "rule_list.rules.asn_matcher for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 1111, "body_sha256": "sha256:e2c843c5bd2f07fc75fa2441d5636710fbfb5d08ebcbc345cdb2d475c1e75e4e", "canonical_id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "child_ids": ["xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:asn_matcher:asn_sets"], "collection_id": "xcsh-docs:data-sources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "parent_id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules", "path": "docs/guides/data-sources--dns_load_balancer--properties--rule_list--rules--asn_matcher.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "asn_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_load_balancer/properties/rule_list/rules/asn_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.asn_matcher for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rule_list.rules.asn_matcher

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md)
- [Property reference](data-sources--dns_load_balancer--reference.md)
- [rule_list](data-sources--dns_load_balancer--properties--rule_list.md)
- [rule_list.rules](data-sources--dns_load_balancer--properties--rule_list--rules.md)
- rule_list.rules.asn_matcher

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

- [asn_sets](data-sources--dns_load_balancer--properties--rule_list--rules--asn_matcher--asn_sets.md): complete subsection reference.

## Next pages

- [rule_list.rules.asn_matcher.asn_sets](data-sources--dns_load_balancer--properties--rule_list--rules--asn_matcher--asn_sets.md)
- [rule_list.rules](data-sources--dns_load_balancer--properties--rule_list--rules.md)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md)
