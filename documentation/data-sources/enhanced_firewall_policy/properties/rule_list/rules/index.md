---
page_title: "rule_list.rules"
subcategory: ""
description: "Ordered List of Enhanced Firewall Policy Rules."
xcsh_docs: {"aliases": ["rule list rules"], "body_bytes": 11831, "body_sha256": "sha256:4fe3df5b9ade05b24ebd00397e78626c4b41df8e8ccfede7fb85070e0dbc48dc", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:advanced_action", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_destinations", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_sli_vips", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_slo_vips", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_sources", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_tcp_traffic", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_traffic", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_udp_traffic", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:allow", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:applications", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:deny", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_aws_vpc_ids", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_ip_prefix_set", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_label_selector", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_prefix_list", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:insert_service", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:inside_destinations", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:inside_sources", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:label_matcher", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:metadata", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:outside_destinations", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:outside_sources", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:protocol_port_range", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_aws_vpc_ids", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_ip_prefix_set", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_label_selector", "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_prefix_list"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules", "parent_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list", "path": "documentation/data-sources/enhanced_firewall_policy/properties/rule_list/rules/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3223132132303231-3200020031132030-0022221302011031-2333313023201332-1233023001030100-3030331221100130-3011231033310202-2310223110022331", "registry_path": "docs/guides/data-sources--enhanced_firewall_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules"], "schema_version": 1, "sections": [{"aliases": ["rule list rules advanced action"], "anchor": "section", "description": "Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and PBRRuleAction.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:advanced_action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "advanced_action"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules all destinations"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_destinations", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "all_destinations"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules all sli vips"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_sli_vips", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "all_sli_vips"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules all slo vips"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_slo_vips", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "all_slo_vips"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules all sources"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_sources", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "all_sources"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules all tcp traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_tcp_traffic", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "all_tcp_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules all traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_traffic", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "all_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules all udp traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:all_udp_traffic", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "all_udp_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules allow"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:allow", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "allow"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules applications"], "anchor": "section", "description": "Application protocols like HTTP, SNMP.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:applications", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "applications"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules deny"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:deny", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "deny"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules destination aws vpc ids"], "anchor": "section", "description": "List of VPC Identifiers in AWS.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_aws_vpc_ids", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "destination_aws_vpc_ids"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules destination ip prefix set"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_ip_prefix_set", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "destination_ip_prefix_set"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules destination label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_label_selector", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "destination_label_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules destination prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:destination_prefix_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "destination_prefix_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules insert service"], "anchor": "section", "description": "Action to forward traffic to external service.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:insert_service", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "insert_service"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules inside destinations"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:inside_destinations", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "inside_destinations"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules inside sources"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:inside_sources", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "inside_sources"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules label matcher"], "anchor": "section", "description": "A label matcher specifies a list of label keys whose values need to match for source/client and destination/server. Note that the actual label values are not specified and do not matter. This allows an ability to scope grouping by the label key name.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:label_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "label_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules outside destinations"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:outside_destinations", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "outside_destinations"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules outside sources"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:outside_sources", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "outside_sources"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules protocol port range"], "anchor": "section", "description": "Protocol and Port ranges.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:protocol_port_range", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "protocol_port_range"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules source aws vpc ids"], "anchor": "section", "description": "List of VPC Identifiers in AWS.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_aws_vpc_ids", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "source_aws_vpc_ids"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules source ip prefix set"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_ip_prefix_set", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "source_ip_prefix_set"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules source label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_label_selector", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "source_label_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules source prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_prefix_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "source_prefix_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/enhanced_firewall_policy/properties/rule_list/rules/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Ordered List of Enhanced Firewall Policy Rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
