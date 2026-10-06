---
page_title: "ingress_rules"
subcategory: ""
description: "Ordered list of rules applied to connections to policy endpoints."
xcsh_docs: {"aliases": ["ingress rules"], "body_bytes": 4370, "body_sha256": "sha256:5336b149ddc3d81072f73f386020239e7575885d565a2a49c1aebb9fa3545e17", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:adv_action", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:all_tcp_traffic", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:all_traffic", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:all_udp_traffic", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:any", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:applications", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:inside_endpoints", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:ip_prefix_set", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:label_matcher", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:label_selector", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:metadata", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:outside_endpoints", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:prefix_list", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:protocol_port_range"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules", "parent_id": "xcsh-docs:data-sources:network_policy_view:reference", "path": "documentation/data-sources/network_policy_view/properties/ingress_rules/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230", "registry_path": "docs/guides/data-sources--network_policy_view--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_rules"], "schema_version": 1, "sections": [{"aliases": ["ingress rules action"], "anchor": "schema-ingress_rules--action", "description": "Network policy rule action configures the action to be taken on rule match Apply deny action on rule match Apply allow action on rule match.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_rules", "action"], "syntax": "attribute", "type": "string"}, {"aliases": ["ingress rules adv action"], "anchor": "section", "description": "Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and PBRRuleAction.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:adv_action", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_rules", "adv_action"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress rules all tcp traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:all_tcp_traffic", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_rules", "all_tcp_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress rules all traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:all_traffic", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_rules", "all_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress rules all udp traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:all_udp_traffic", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_rules", "all_udp_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress rules any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:any", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_rules", "any"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress rules applications"], "anchor": "section", "description": "Application protocols like HTTP, SNMP.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:applications", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_rules", "applications"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress rules inside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:inside_endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_rules", "inside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress rules ip prefix set"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:ip_prefix_set", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_rules", "ip_prefix_set"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress rules label matcher"], "anchor": "section", "description": "A label matcher specifies a list of label keys whose values need to match for source/client and destination/server. Note that the actual label values are not specified and do not matter. This allows an ability to scope grouping by the label key name.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:label_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_rules", "label_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress rules label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:label_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_rules", "label_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress rules outside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:outside_endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_rules", "outside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress rules prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:prefix_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_rules", "prefix_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress rules protocol port range"], "anchor": "section", "description": "Protocol and Port ranges.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:protocol_port_range", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_rules", "protocol_port_range"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/properties/ingress_rules/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Ordered list of rules applied to connections to policy endpoints.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_rules

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/)
- ingress_rules

<a id="section"></a>

Type: `"list"`. Computed.

Ordered list of rules applied to connections to policy endpoints.

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

<a id="schema-ingress_rules--action"></a>

### action property

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

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

- [adv_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/adv_action/): complete subsection reference.

- [all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/all_tcp_traffic/): complete subsection reference.

- [all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/all_traffic/): complete subsection reference.

- [all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/all_udp_traffic/): complete subsection reference.

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/any/): complete subsection reference.

- [applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/applications/): complete subsection reference.

- [inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/inside_endpoints/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/ip_prefix_set/): complete subsection reference.

- [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/label_matcher/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/label_selector/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/metadata/): complete subsection reference.

- [outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/outside_endpoints/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/prefix_list/): complete subsection reference.

- [protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/protocol_port_range/): complete subsection reference.
