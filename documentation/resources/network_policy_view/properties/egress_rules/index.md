---
page_title: "egress_rules"
subcategory: ""
description: "Ordered list of rules applied to connections from policy endpoints."
xcsh_docs: {"aliases": ["egress rules"], "body_bytes": 7058, "body_sha256": "sha256:0b01737002fd63913a65fd9caec21f01d48399816bb73b6175d45c4ebfbdc2b6", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:network_policy_view:properties:egress_rules:adv_action", "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_tcp_traffic", "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_traffic", "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_udp_traffic", "xcsh-docs:resources:network_policy_view:properties:egress_rules:any", "xcsh-docs:resources:network_policy_view:properties:egress_rules:applications", "xcsh-docs:resources:network_policy_view:properties:egress_rules:inside_endpoints", "xcsh-docs:resources:network_policy_view:properties:egress_rules:ip_prefix_set", "xcsh-docs:resources:network_policy_view:properties:egress_rules:label_matcher", "xcsh-docs:resources:network_policy_view:properties:egress_rules:label_selector", "xcsh-docs:resources:network_policy_view:properties:egress_rules:metadata", "xcsh-docs:resources:network_policy_view:properties:egress_rules:outside_endpoints", "xcsh-docs:resources:network_policy_view:properties:egress_rules:prefix_list", "xcsh-docs:resources:network_policy_view:properties:egress_rules:protocol_port_range"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:properties:egress_rules", "parent_id": "xcsh-docs:resources:network_policy_view:reference", "path": "documentation/resources/network_policy_view/properties/egress_rules/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213", "registry_path": "docs/guides/resources--network_policy_view--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_tcp_traffic,all_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_tcp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_tcp_traffic,all_udp_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_tcp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_tcp_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_tcp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_tcp_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_tcp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_tcp_traffic,all_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_traffic,all_udp_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_tcp_traffic,all_udp_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_udp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_traffic,all_udp_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_udp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_udp_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_udp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_udp_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_udp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:any,inside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:any,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:any,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:any,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:any,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_tcp_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:applications", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:applications", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_udp_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:applications", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:applications,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:applications", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:any,inside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:inside_endpoints,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:inside_endpoints,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:inside_endpoints,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:inside_endpoints,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:any,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:inside_endpoints,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:ip_prefix_set,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:ip_prefix_set,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:any,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:inside_endpoints,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:ip_prefix_set,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:label_selector,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:label_selector,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:any,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:inside_endpoints,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:ip_prefix_set,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:label_selector,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:outside_endpoints,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:any,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:inside_endpoints,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:label_selector,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:outside_endpoints,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_tcp_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:protocol_port_range", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:protocol_port_range", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:all_udp_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:protocol_port_range", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "egress_rules:ConflictingListObjectAttributes:applications,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:protocol_port_range", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["egress_rules"], "schema_version": 1, "sections": [{"aliases": ["egress rules action"], "anchor": "schema-egress_rules--action", "description": "Network policy rule action configures the action to be taken on rule match Apply deny action on rule match Apply allow action on rule match.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ALLOW", "DENY"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "action"], "syntax": "attribute", "type": "string"}, {"aliases": ["egress rules adv action"], "anchor": "section", "description": "Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and PBRRuleAction.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:adv_action", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["egress_rules", "adv_action"], "syntax": "block", "type": "object"}, {"aliases": ["egress rules all tcp traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_tcp_traffic", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "all_tcp_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["egress rules all traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_traffic", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "all_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["egress rules all udp traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:all_udp_traffic", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "all_udp_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["egress rules any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:any", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "any"], "syntax": "attribute", "type": "object"}, {"aliases": ["egress rules applications"], "anchor": "section", "description": "Application protocols like HTTP, SNMP.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:applications", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["egress_rules", "applications"], "syntax": "block", "type": "object"}, {"aliases": ["egress rules inside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:inside_endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "inside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["egress rules ip prefix set"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:ip_prefix_set", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["egress_rules", "ip_prefix_set"], "syntax": "block", "type": "object"}, {"aliases": ["egress rules label matcher"], "anchor": "section", "description": "A label matcher specifies a list of label keys whose values need to match for source/client and destination/server. Note that the actual label values are not specified and do not matter. This allows an ability to scope grouping by the label key name.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:label_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["egress_rules", "label_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["egress rules label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:label_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-egress_rules--label_selector--expressions", "enforcement": "provider-schema", "group": "egress_rules.label_selector:RequiredObjectAttributes:expressions", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:label_selector", "type": "requires"}], "schema_path": ["egress_rules", "label_selector"], "syntax": "block", "type": "object"}, {"aliases": ["egress rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-egress_rules--metadata--name", "enforcement": "provider-schema", "group": "egress_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:metadata", "type": "requires"}], "schema_path": ["egress_rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["egress rules outside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:outside_endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_rules", "outside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["egress rules prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:prefix_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["egress_rules", "prefix_list"], "syntax": "block", "type": "object"}, {"aliases": ["egress rules protocol port range"], "anchor": "section", "description": "Protocol and Port ranges.", "document_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:protocol_port_range", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["egress_rules", "protocol_port_range"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/properties/egress_rules/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Ordered list of rules applied to connections from policy endpoints.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# egress_rules

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/)
- egress_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of rules applied to connections from policy endpoints.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_tcp_traffic",
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
  validators.ConflictingListObjectAttributes("any",
    "inside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("any",
    "label_selector"),
  validators.ConflictingListObjectAttributes("any",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("applications",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "label_selector"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "label_selector"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("label_selector",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("outside_endpoints",
    "prefix_list")}
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Terraform syntax:

```terraform
egress_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-egress_rules--action"></a>

### action property

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALLOW","DENY"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW"),
}
```

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

- [adv_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/adv_action/): complete subsection reference.

- [all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/all_tcp_traffic/): complete subsection reference.

- [all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/all_traffic/): complete subsection reference.

- [all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/all_udp_traffic/): complete subsection reference.

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/any/): complete subsection reference.

- [applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/applications/): complete subsection reference.

- [inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/inside_endpoints/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/ip_prefix_set/): complete subsection reference.

- [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/label_matcher/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/label_selector/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/metadata/): complete subsection reference.

- [outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/outside_endpoints/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/prefix_list/): complete subsection reference.

- [protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/egress_rules/protocol_port_range/): complete subsection reference.
