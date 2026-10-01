---
page_title: "rule_list.rules.source_ip_prefix_set"
subcategory: ""
description: "rule_list.rules.source_ip_prefix_set for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1364, "body_sha256": "sha256:40be7717bb73aea0f095be616f55a8f6765f6246b6b8c8abb7e6ed13504117d5", "canonical_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_ip_prefix_set", "child_ids": ["xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_ip_prefix_set:ref"], "collection_id": "xcsh-docs:data-sources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_ip_prefix_set", "parent_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules", "path": "docs/guides/data-sources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set.md", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "source_ip_prefix_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.source_ip_prefix_set for xcsh_enhanced_firewall_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.source_ip_prefix_set

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md)
- [Property reference](data-sources--enhanced_firewall_policy--reference.md)
- [rule_list](data-sources--enhanced_firewall_policy--properties--rule_list.md)
- [rule_list.rules](data-sources--enhanced_firewall_policy--properties--rule_list--rules.md)
- rule_list.rules.source_ip_prefix_set

<a id="section"></a>

Type: `"single"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

- [ref](data-sources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set--ref.md): complete subsection reference.

## Next pages

- [rule_list.rules.source_ip_prefix_set.ref](data-sources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set--ref.md)
- [rule_list.rules](data-sources--enhanced_firewall_policy--properties--rule_list--rules.md)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md)
