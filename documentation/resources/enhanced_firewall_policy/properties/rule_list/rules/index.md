---
page_title: "rule_list.rules"
subcategory: ""
description: "rule_list.rules for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": [], "body_bytes": 18366, "body_sha256": "sha256:3a13e6fb38ec6193c50b102c72e509cf0ae51f9929c7bcb41747ba6b0c4fc092", "child_ids": ["xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:advanced_action", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:all_destinations", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:all_sli_vips", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:all_slo_vips", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:all_sources", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:all_tcp_traffic", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:all_traffic", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:all_udp_traffic", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:allow", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:applications", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:deny", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:destination_aws_vpc_ids", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:destination_ip_prefix_set", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:destination_label_selector", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:destination_prefix_list", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:insert_service", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:inside_destinations", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:inside_sources", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:label_matcher", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:metadata", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:outside_destinations", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:outside_sources", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:protocol_port_range", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:source_aws_vpc_ids", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:source_ip_prefix_set", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:source_label_selector", "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:source_prefix_list"], "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules", "parent_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list", "path": "documentation/resources/enhanced_firewall_policy/properties/rule_list/rules/index.md", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["rule_list", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/properties/rule_list/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules for xcsh_enhanced_firewall_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/)
- rule_list.rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policy Rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_destinations",
    "all_sli_vips"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "all_slo_vips"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "destination_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "destination_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "all_slo_vips"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "destination_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "destination_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "destination_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "destination_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("all_sources",
    "inside_sources"),
  validators.ConflictingListObjectAttributes("all_sources",
    "outside_sources"),
  validators.ConflictingListObjectAttributes("all_sources",
    "source_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("all_sources",
    "source_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_sources",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("all_sources",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("allow",
    "deny"),
  validators.ConflictingListObjectAttributes("allow",
    "insert_service"),
  validators.ConflictingListObjectAttributes("applications",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("deny",
    "insert_service"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "destination_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("destination_ip_prefix_set",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("destination_ip_prefix_set",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("destination_ip_prefix_set",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("destination_ip_prefix_set",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("destination_label_selector",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("destination_label_selector",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("destination_label_selector",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("destination_prefix_list",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("destination_prefix_list",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("inside_destinations",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "outside_sources"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "source_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "source_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("outside_sources",
    "source_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("outside_sources",
    "source_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("outside_sources",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("outside_sources",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("source_aws_vpc_ids",
    "source_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("source_aws_vpc_ids",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("source_aws_vpc_ids",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("source_ip_prefix_set",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("source_ip_prefix_set",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("source_label_selector",
    "source_prefix_list")}
```

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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advanced_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/advanced_action/): complete subsection reference.

- [all_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_destinations/): complete subsection reference.

- [all_sli_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_sli_vips/): complete subsection reference.

- [all_slo_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_slo_vips/): complete subsection reference.

- [all_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_sources/): complete subsection reference.

- [all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_tcp_traffic/): complete subsection reference.

- [all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_traffic/): complete subsection reference.

- [all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_udp_traffic/): complete subsection reference.

- [allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/allow/): complete subsection reference.

- [applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/applications/): complete subsection reference.

- [deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/deny/): complete subsection reference.

- [destination_aws_vpc_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_aws_vpc_ids/): complete subsection reference.

- [destination_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_ip_prefix_set/): complete subsection reference.

- [destination_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_label_selector/): complete subsection reference.

- [destination_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_prefix_list/): complete subsection reference.

- [insert_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/insert_service/): complete subsection reference.

- [inside_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/inside_destinations/): complete subsection reference.

- [inside_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/inside_sources/): complete subsection reference.

- [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/label_matcher/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/metadata/): complete subsection reference.

- [outside_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/outside_destinations/): complete subsection reference.

- [outside_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/outside_sources/): complete subsection reference.

- [protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/protocol_port_range/): complete subsection reference.

- [source_aws_vpc_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_aws_vpc_ids/): complete subsection reference.

- [source_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/): complete subsection reference.

- [source_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_label_selector/): complete subsection reference.

- [source_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_prefix_list/): complete subsection reference.

## Next pages

- [rule_list.rules.advanced_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/advanced_action/)
- [rule_list.rules.all_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_destinations/)
- [rule_list.rules.all_sli_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_sli_vips/)
- [rule_list.rules.all_slo_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_slo_vips/)
- [rule_list.rules.all_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_sources/)
- [rule_list.rules.all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_tcp_traffic/)
- [rule_list.rules.all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_traffic/)
- [rule_list.rules.all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/all_udp_traffic/)
- [rule_list.rules.allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/allow/)
- [rule_list.rules.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/applications/)
- [rule_list.rules.deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/deny/)
- [rule_list.rules.destination_aws_vpc_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_aws_vpc_ids/)
- [rule_list.rules.destination_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_ip_prefix_set/)
- [rule_list.rules.destination_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_label_selector/)
- [rule_list.rules.destination_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_prefix_list/)
- [rule_list.rules.insert_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/insert_service/)
- [rule_list.rules.inside_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/inside_destinations/)
- [rule_list.rules.inside_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/inside_sources/)
- [rule_list.rules.label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/label_matcher/)
- [rule_list.rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/metadata/)
- [rule_list.rules.outside_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/outside_destinations/)
- [rule_list.rules.outside_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/outside_sources/)
- [rule_list.rules.protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/protocol_port_range/)
- [rule_list.rules.source_aws_vpc_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_aws_vpc_ids/)
- [rule_list.rules.source_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/)
- [rule_list.rules.source_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_label_selector/)
- [rule_list.rules.source_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_prefix_list/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/)
- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
