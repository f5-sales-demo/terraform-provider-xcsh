---
page_title: "rule_list.rules"
subcategory: ""
description: "rule_list.rules for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": [], "body_bytes": 8928, "body_sha256": "sha256:e7a2a3d909364392e2ff01f02e6f3b5d83508941db442ff1460563e151b73ae6", "canonical_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules", "child_ids": ["xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:advanced_action", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_destinations", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_sli_vips", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_slo_vips", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_sources", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_tcp_traffic", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_traffic", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_udp_traffic", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:allow", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:applications", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:deny", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_aws_vpc_ids", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_ip_prefix_set", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_label_selector", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_prefix_list", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:insert_service", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:inside_destinations", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:inside_sources", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:label_matcher", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:metadata", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:outside_destinations", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:outside_sources", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:protocol_port_range", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_aws_vpc_ids", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_ip_prefix_set", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_label_selector", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_prefix_list"], "collection_id": "xcsh-docs:data-sources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules", "parent_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list", "path": "docs/guides/data-sources--enhanced_firewall_policy--properties--rule_list--rules.md", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/enhanced_firewall_policy/properties/rule_list/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules for xcsh_enhanced_firewall_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md)
- [Property reference](data-sources--enhanced_firewall_policy--reference.md)
- [rule_list](data-sources--enhanced_firewall_policy--properties--rule_list.md)
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

- [advanced_action](data-sources--enhanced_firewall_policy--properties--rule_list--rules--advanced_action.md): complete subsection reference.

- [all_destinations](data-sources--enhanced_firewall_policy--properties--rule_list--rules--all_destinations.md): complete subsection reference.

- [all_sli_vips](data-sources--enhanced_firewall_policy--properties--rule_list--rules--all_sli_vips.md): complete subsection reference.

- [all_slo_vips](data-sources--enhanced_firewall_policy--properties--rule_list--rules--all_slo_vips.md): complete subsection reference.

- [all_sources](data-sources--enhanced_firewall_policy--properties--rule_list--rules--all_sources.md): complete subsection reference.

- [all_tcp_traffic](data-sources--enhanced_firewall_policy--properties--rule_list--rules--all_tcp_traffic.md): complete subsection reference.

- [all_traffic](data-sources--enhanced_firewall_policy--properties--rule_list--rules--all_traffic.md): complete subsection reference.

- [all_udp_traffic](data-sources--enhanced_firewall_policy--properties--rule_list--rules--all_udp_traffic.md): complete subsection reference.

- [allow](data-sources--enhanced_firewall_policy--properties--rule_list--rules--allow.md): complete subsection reference.

- [applications](data-sources--enhanced_firewall_policy--properties--rule_list--rules--applications.md): complete subsection reference.

- [deny](data-sources--enhanced_firewall_policy--properties--rule_list--rules--deny.md): complete subsection reference.

- [destination_aws_vpc_ids](data-sources--enhanced_firewall_policy--properties--rule_list--rules--destination_aws_vpc_ids.md): complete subsection reference.

- [destination_ip_prefix_set](data-sources--enhanced_firewall_policy--properties--rule_list--rules--destination_ip_prefix_set.md): complete subsection reference.

- [destination_label_selector](data-sources--enhanced_firewall_policy--properties--rule_list--rules--destination_label_selector.md): complete subsection reference.

- [destination_prefix_list](data-sources--enhanced_firewall_policy--properties--rule_list--rules--destination_prefix_list.md): complete subsection reference.

- [insert_service](data-sources--enhanced_firewall_policy--properties--rule_list--rules--insert_service.md): complete subsection reference.

- [inside_destinations](data-sources--enhanced_firewall_policy--properties--rule_list--rules--inside_destinations.md): complete subsection reference.

- [inside_sources](data-sources--enhanced_firewall_policy--properties--rule_list--rules--inside_sources.md): complete subsection reference.

- [label_matcher](data-sources--enhanced_firewall_policy--properties--rule_list--rules--label_matcher.md): complete subsection reference.

- [metadata](data-sources--enhanced_firewall_policy--properties--rule_list--rules--metadata.md): complete subsection reference.

- [outside_destinations](data-sources--enhanced_firewall_policy--properties--rule_list--rules--outside_destinations.md): complete subsection reference.

- [outside_sources](data-sources--enhanced_firewall_policy--properties--rule_list--rules--outside_sources.md): complete subsection reference.

- [protocol_port_range](data-sources--enhanced_firewall_policy--properties--rule_list--rules--protocol_port_range.md): complete subsection reference.

- [source_aws_vpc_ids](data-sources--enhanced_firewall_policy--properties--rule_list--rules--source_aws_vpc_ids.md): complete subsection reference.

- [source_ip_prefix_set](data-sources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set.md): complete subsection reference.

- [source_label_selector](data-sources--enhanced_firewall_policy--properties--rule_list--rules--source_label_selector.md): complete subsection reference.

- [source_prefix_list](data-sources--enhanced_firewall_policy--properties--rule_list--rules--source_prefix_list.md): complete subsection reference.

## Next pages

- [rule_list.rules.advanced_action](data-sources--enhanced_firewall_policy--properties--rule_list--rules--advanced_action.md)
- [rule_list.rules.all_destinations](data-sources--enhanced_firewall_policy--properties--rule_list--rules--all_destinations.md)
- [rule_list.rules.all_sli_vips](data-sources--enhanced_firewall_policy--properties--rule_list--rules--all_sli_vips.md)
- [rule_list.rules.all_slo_vips](data-sources--enhanced_firewall_policy--properties--rule_list--rules--all_slo_vips.md)
- [rule_list.rules.all_sources](data-sources--enhanced_firewall_policy--properties--rule_list--rules--all_sources.md)
- [rule_list.rules.all_tcp_traffic](data-sources--enhanced_firewall_policy--properties--rule_list--rules--all_tcp_traffic.md)
- [rule_list.rules.all_traffic](data-sources--enhanced_firewall_policy--properties--rule_list--rules--all_traffic.md)
- [rule_list.rules.all_udp_traffic](data-sources--enhanced_firewall_policy--properties--rule_list--rules--all_udp_traffic.md)
- [rule_list.rules.allow](data-sources--enhanced_firewall_policy--properties--rule_list--rules--allow.md)
- [rule_list.rules.applications](data-sources--enhanced_firewall_policy--properties--rule_list--rules--applications.md)
- [rule_list.rules.deny](data-sources--enhanced_firewall_policy--properties--rule_list--rules--deny.md)
- [rule_list.rules.destination_aws_vpc_ids](data-sources--enhanced_firewall_policy--properties--rule_list--rules--destination_aws_vpc_ids.md)
- [rule_list.rules.destination_ip_prefix_set](data-sources--enhanced_firewall_policy--properties--rule_list--rules--destination_ip_prefix_set.md)
- [rule_list.rules.destination_label_selector](data-sources--enhanced_firewall_policy--properties--rule_list--rules--destination_label_selector.md)
- [rule_list.rules.destination_prefix_list](data-sources--enhanced_firewall_policy--properties--rule_list--rules--destination_prefix_list.md)
- [rule_list.rules.insert_service](data-sources--enhanced_firewall_policy--properties--rule_list--rules--insert_service.md)
- [rule_list.rules.inside_destinations](data-sources--enhanced_firewall_policy--properties--rule_list--rules--inside_destinations.md)
- [rule_list.rules.inside_sources](data-sources--enhanced_firewall_policy--properties--rule_list--rules--inside_sources.md)
- [rule_list.rules.label_matcher](data-sources--enhanced_firewall_policy--properties--rule_list--rules--label_matcher.md)
- [rule_list.rules.metadata](data-sources--enhanced_firewall_policy--properties--rule_list--rules--metadata.md)
- [rule_list.rules.outside_destinations](data-sources--enhanced_firewall_policy--properties--rule_list--rules--outside_destinations.md)
- [rule_list.rules.outside_sources](data-sources--enhanced_firewall_policy--properties--rule_list--rules--outside_sources.md)
- [rule_list.rules.protocol_port_range](data-sources--enhanced_firewall_policy--properties--rule_list--rules--protocol_port_range.md)
- [rule_list.rules.source_aws_vpc_ids](data-sources--enhanced_firewall_policy--properties--rule_list--rules--source_aws_vpc_ids.md)
- [rule_list.rules.source_ip_prefix_set](data-sources--enhanced_firewall_policy--properties--rule_list--rules--source_ip_prefix_set.md)
- [rule_list.rules.source_label_selector](data-sources--enhanced_firewall_policy--properties--rule_list--rules--source_label_selector.md)
- [rule_list.rules.source_prefix_list](data-sources--enhanced_firewall_policy--properties--rule_list--rules--source_prefix_list.md)
- [rule_list](data-sources--enhanced_firewall_policy--properties--rule_list.md)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md)
