---
page_title: "rule_list.rules"
subcategory: ""
description: "rule_list.rules for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": [], "body_bytes": 11831, "body_sha256": "sha256:2f92cd40ffe07f24c50115ab8787569d4c13e0955b3f327c9feb7d7158bdb72e", "child_ids": ["xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:advanced_action", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_destinations", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_sli_vips", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_slo_vips", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_sources", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_tcp_traffic", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_traffic", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_udp_traffic", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:allow", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:applications", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:deny", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_aws_vpc_ids", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_ip_prefix_set", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_label_selector", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_prefix_list", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:insert_service", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:inside_destinations", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:inside_sources", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:label_matcher", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:metadata", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:outside_destinations", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:outside_sources", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:protocol_port_range", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_aws_vpc_ids", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_ip_prefix_set", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_label_selector", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_prefix_list"], "collection_id": "xcsh-docs:data-sources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules", "parent_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list", "path": "documentation/data-sources/enhanced_firewall_policy/properties/rule_list/rules/index.md", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["rule_list", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/enhanced_firewall_policy/properties/rule_list/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules for xcsh_enhanced_firewall_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/)
- rule_list.rules

<a id="section"></a>

Type: `"list"`. Computed.

Ordered List of Enhanced Firewall Policy Rules.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [advanced_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/advanced_action/): complete subsection reference.

- [all_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/all_destinations/): complete subsection reference.

- [all_sli_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/all_sli_vips/): complete subsection reference.

- [all_slo_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/all_slo_vips/): complete subsection reference.

- [all_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/all_sources/): complete subsection reference.

- [all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/all_tcp_traffic/): complete subsection reference.

- [all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/all_traffic/): complete subsection reference.

- [all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/all_udp_traffic/): complete subsection reference.

- [allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/allow/): complete subsection reference.

- [applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/applications/): complete subsection reference.

- [deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/deny/): complete subsection reference.

- [destination_aws_vpc_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/destination_aws_vpc_ids/): complete subsection reference.

- [destination_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/destination_ip_prefix_set/): complete subsection reference.

- [destination_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/destination_label_selector/): complete subsection reference.

- [destination_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/destination_prefix_list/): complete subsection reference.

- [insert_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/insert_service/): complete subsection reference.

- [inside_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/inside_destinations/): complete subsection reference.

- [inside_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/inside_sources/): complete subsection reference.

- [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/label_matcher/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/metadata/): complete subsection reference.

- [outside_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/outside_destinations/): complete subsection reference.

- [outside_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/outside_sources/): complete subsection reference.

- [protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/protocol_port_range/): complete subsection reference.

- [source_aws_vpc_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/source_aws_vpc_ids/): complete subsection reference.

- [source_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/): complete subsection reference.

- [source_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/source_label_selector/): complete subsection reference.

- [source_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/source_prefix_list/): complete subsection reference.

## Next pages

- [rule_list.rules.advanced_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/advanced_action/)
- [rule_list.rules.all_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/all_destinations/)
- [rule_list.rules.all_sli_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/all_sli_vips/)
- [rule_list.rules.all_slo_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/all_slo_vips/)
- [rule_list.rules.all_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/all_sources/)
- [rule_list.rules.all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/all_tcp_traffic/)
- [rule_list.rules.all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/all_traffic/)
- [rule_list.rules.all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/all_udp_traffic/)
- [rule_list.rules.allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/allow/)
- [rule_list.rules.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/applications/)
- [rule_list.rules.deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/deny/)
- [rule_list.rules.destination_aws_vpc_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/destination_aws_vpc_ids/)
- [rule_list.rules.destination_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/destination_ip_prefix_set/)
- [rule_list.rules.destination_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/destination_label_selector/)
- [rule_list.rules.destination_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/destination_prefix_list/)
- [rule_list.rules.insert_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/insert_service/)
- [rule_list.rules.inside_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/inside_destinations/)
- [rule_list.rules.inside_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/inside_sources/)
- [rule_list.rules.label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/label_matcher/)
- [rule_list.rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/metadata/)
- [rule_list.rules.outside_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/outside_destinations/)
- [rule_list.rules.outside_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/outside_sources/)
- [rule_list.rules.protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/protocol_port_range/)
- [rule_list.rules.source_aws_vpc_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/source_aws_vpc_ids/)
- [rule_list.rules.source_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/)
- [rule_list.rules.source_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/source_label_selector/)
- [rule_list.rules.source_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/source_prefix_list/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/)
- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/)
