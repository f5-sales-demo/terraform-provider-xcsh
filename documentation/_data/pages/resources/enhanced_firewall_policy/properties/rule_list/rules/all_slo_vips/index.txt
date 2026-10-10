---
page_title: "rule_list.rules.all_slo_vips"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rule list rules all slo vips"], "body_bytes": 1166, "body_sha256": "sha256:f349ccb4c4c30789251b774f26a715007df1fb8c1a6d369ca64503bef607f3ac", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:all_slo_vips", "parent_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules", "path": "documentation/resources/enhanced_firewall_policy/properties/rule_list/rules/all_slo_vips/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2210010332032211-1132103201011133-0213113111322302-1210232313222102-3331111112121232-2231133003310122-2031222322122332-1001112012233101", "registry_path": "docs/guides/resources--enhanced_firewall_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "all_slo_vips"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/properties/rule_list/rules/all_slo_vips/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.all_slo_vips

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/)
- rule_list.rules.all_slo_vips

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
all_slo_vips = {}
```

This is an empty object or choice marker. It has no direct properties.
