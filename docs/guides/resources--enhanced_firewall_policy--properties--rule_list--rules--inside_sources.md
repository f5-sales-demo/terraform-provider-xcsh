---
page_title: "rule_list.rules.inside_sources"
subcategory: ""
description: "rule_list.rules.inside_sources for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1189, "body_sha256": "sha256:1aa014a16c252b427bed9c5193ec615b0c108e71cc303a65711281cab2b82786", "canonical_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:inside_sources", "child_ids": [], "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:inside_sources", "parent_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules", "path": "docs/guides/resources--enhanced_firewall_policy--properties--rule_list--rules--inside_sources.md", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "inside_sources"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/properties/rule_list/rules/inside_sources/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.inside_sources for xcsh_enhanced_firewall_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.inside_sources

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md)
- [Property reference](resources--enhanced_firewall_policy--reference.md)
- [rule_list](resources--enhanced_firewall_policy--properties--rule_list.md)
- [rule_list.rules](resources--enhanced_firewall_policy--properties--rule_list--rules.md)
- rule_list.rules.inside_sources

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside sources.

Upstream description:

This can be used for messages where no values are needed.

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
inside_sources = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rule_list.rules](resources--enhanced_firewall_policy--properties--rule_list--rules.md)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md)
