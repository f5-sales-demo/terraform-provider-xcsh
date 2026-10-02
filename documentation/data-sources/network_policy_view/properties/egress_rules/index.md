---
page_title: "egress_rules"
subcategory: ""
description: "Ordered list of rules applied to connections from policy endpoints."
xcsh_docs: {"aliases": ["egress rules"], "body_bytes": 7012, "body_sha256": "sha256:3d80a0f131dedc89cf6abc7f5973b904558b9d6f19978939846481916ee4e32c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_policy_view:properties:egress_rules:adv_action", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:all_tcp_traffic", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:all_traffic", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:all_udp_traffic", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:any", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:applications", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:inside_endpoints", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:ip_prefix_set", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:label_matcher", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:label_selector", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:metadata", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:outside_endpoints", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:prefix_list", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:protocol_port_range"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules", "parent_id": "xcsh-docs:data-sources:network_policy_view:reference", "path": "documentation/data-sources/network_policy_view/properties/egress_rules/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022", "registry_path": "docs/guides/data-sources--network_policy_view--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["egress_rules"], "schema_version": 1, "sections": [{"aliases": ["action"], "anchor": "schema-egress_rules--action", "description": "Network policy rule action configures the action to be taken on rule match Apply deny action on rule match Apply allow action on rule match.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "action"], "syntax": "attribute", "type": "string"}, {"aliases": ["adv action"], "anchor": "section", "description": "Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and PBRRuleAction.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:adv_action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["egress_rules", "adv_action"], "syntax": "attribute", "type": "object"}, {"aliases": ["all tcp traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:all_tcp_traffic", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "all_tcp_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["all traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:all_traffic", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "all_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["all udp traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:all_udp_traffic", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "all_udp_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:any", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "any"], "syntax": "attribute", "type": "object"}, {"aliases": ["applications"], "anchor": "section", "description": "Application protocols like HTTP, SNMP.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:applications", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["egress_rules", "applications"], "syntax": "attribute", "type": "object"}, {"aliases": ["inside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:inside_endpoints", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "inside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["ip prefix set"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:ip_prefix_set", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["egress_rules", "ip_prefix_set"], "syntax": "attribute", "type": "object"}, {"aliases": ["label matcher"], "anchor": "section", "description": "A label matcher specifies a list of label keys whose values need to match for source/client and destination/server. Note that the actual label values are not specified and do not matter. This allows an ability to scope grouping by the label key name.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:label_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["egress_rules", "label_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:label_selector", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["egress_rules", "label_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["egress_rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["outside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:outside_endpoints", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "outside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:prefix_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["egress_rules", "prefix_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["protocol port range"], "anchor": "section", "description": "Protocol and Port ranges.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:protocol_port_range", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["egress_rules", "protocol_port_range"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/properties/egress_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Ordered list of rules applied to connections from policy endpoints.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# egress_rules

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/)
- egress_rules

<a id="section"></a>

Type: `"list"`. Computed.

Ordered list of rules applied to connections from policy endpoints.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

<a id="schema-egress_rules--action"></a>

### action property

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [adv_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/adv_action/): complete subsection reference.

- [all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/all_tcp_traffic/): complete subsection reference.

- [all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/all_traffic/): complete subsection reference.

- [all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/all_udp_traffic/): complete subsection reference.

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/any/): complete subsection reference.

- [applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/applications/): complete subsection reference.

- [inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/inside_endpoints/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/ip_prefix_set/): complete subsection reference.

- [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/label_matcher/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/label_selector/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/metadata/): complete subsection reference.

- [outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/outside_endpoints/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/prefix_list/): complete subsection reference.

- [protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/protocol_port_range/): complete subsection reference.

## Next pages

- [egress_rules.adv_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/adv_action/)
- [egress_rules.all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/all_tcp_traffic/)
- [egress_rules.all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/all_traffic/)
- [egress_rules.all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/all_udp_traffic/)
- [egress_rules.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/any/)
- [egress_rules.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/applications/)
- [egress_rules.inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/inside_endpoints/)
- [egress_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/ip_prefix_set/)
- [egress_rules.label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/label_matcher/)
- [egress_rules.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/label_selector/)
- [egress_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/metadata/)
- [egress_rules.outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/outside_endpoints/)
- [egress_rules.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/prefix_list/)
- [egress_rules.protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/protocol_port_range/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/)
- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
