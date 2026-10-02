---
page_title: "rules.ingress_rules"
subcategory: "Security"
description: "Ordered list of rules applied to connections to policy endpoints."
xcsh_docs: {"aliases": ["rules ingress rules"], "body_bytes": 9634, "body_sha256": "sha256:1123b7e0f7a28d0976e37a240d85f484b7a5b90caa7f0de7068923812fd73475", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_policy:properties:rules:ingress_rules:adv_action", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_tcp_traffic", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_traffic", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_udp_traffic", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:any", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:applications", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:inside_endpoints", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:ip_prefix_set", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:label_matcher", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:label_selector", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:metadata", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:outside_endpoints", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:prefix_list", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:protocol_port_range"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules", "parent_id": "xcsh-docs:resources:network_policy:properties:rules", "path": "documentation/resources/network_policy/properties/rules/ingress_rules/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133", "registry_path": "docs/guides/resources--network_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_tcp_traffic,all_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_tcp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_tcp_traffic,all_udp_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_tcp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_tcp_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_tcp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_tcp_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_tcp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_tcp_traffic,all_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_traffic,all_udp_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_tcp_traffic,all_udp_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_udp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_traffic,all_udp_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_udp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_udp_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_udp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_udp_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_udp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:any,inside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:any,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:any,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:any,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:any,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_tcp_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:applications", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:applications", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_udp_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:applications", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:applications,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:applications", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:any,inside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:inside_endpoints,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:inside_endpoints,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:inside_endpoints,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:inside_endpoints,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:any,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:inside_endpoints,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:ip_prefix_set,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:ip_prefix_set,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:any,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:inside_endpoints,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:ip_prefix_set,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:label_selector,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:label_selector,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:any,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:inside_endpoints,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:ip_prefix_set,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:label_selector,outside_endpoints", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:outside_endpoints,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:any,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:inside_endpoints,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:label_selector,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:outside_endpoints,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_tcp_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:protocol_port_range", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:protocol_port_range", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:all_udp_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:protocol_port_range", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.ingress_rules:ConflictingListObjectAttributes:applications,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:protocol_port_range", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "ingress_rules"], "schema_version": 1, "sections": [{"aliases": ["action"], "anchor": "schema-rules--ingress_rules--action", "description": "Network policy rule action configures the action to be taken on rule match Apply deny action on rule match Apply allow action on rule match.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ingress_rules", "action"], "syntax": "attribute", "type": "string"}, {"aliases": ["adv action"], "anchor": "section", "description": "Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and PBRRuleAction.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:adv_action", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "ingress_rules", "adv_action"], "syntax": "block", "type": "object"}, {"aliases": ["all tcp traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_tcp_traffic", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ingress_rules", "all_tcp_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["all traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_traffic", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ingress_rules", "all_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["all udp traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_udp_traffic", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ingress_rules", "all_udp_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:any", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ingress_rules", "any"], "syntax": "attribute", "type": "object"}, {"aliases": ["applications"], "anchor": "section", "description": "Application protocols like HTTP, SNMP.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:applications", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "ingress_rules", "applications"], "syntax": "block", "type": "object"}, {"aliases": ["inside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:inside_endpoints", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ingress_rules", "inside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["ip prefix set"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:ip_prefix_set", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "ingress_rules", "ip_prefix_set"], "syntax": "block", "type": "object"}, {"aliases": ["label matcher"], "anchor": "section", "description": "A label matcher specifies a list of label keys whose values need to match for source/client and destination/server. Note that the actual label values are not specified and do not matter. This allows an ability to scope grouping by the label key name.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:label_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "ingress_rules", "label_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:label_selector", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--ingress_rules--label_selector--expressions", "enforcement": "provider-schema", "group": "rules.ingress_rules.label_selector:RequiredObjectAttributes:expressions", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:label_selector", "type": "requires"}], "schema_path": ["rules", "ingress_rules", "label_selector"], "syntax": "block", "type": "object"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--ingress_rules--metadata--name", "enforcement": "provider-schema", "group": "rules.ingress_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:metadata", "type": "requires"}], "schema_path": ["rules", "ingress_rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["outside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:outside_endpoints", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ingress_rules", "outside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:prefix_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "ingress_rules", "prefix_list"], "syntax": "block", "type": "object"}, {"aliases": ["protocol port range"], "anchor": "section", "description": "Protocol and Port ranges.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:protocol_port_range", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "ingress_rules", "protocol_port_range"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/rules/ingress_rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Ordered list of rules applied to connections to policy endpoints.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.ingress_rules

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/)
- rules.ingress_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of rules applied to connections to policy endpoints.

Provider validators and defaults (from schema source):

```go
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

Terraform syntax:

```terraform
ingress_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rules--ingress_rules--action"></a>

### action property

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Provider validators and defaults (from schema source):

```go
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

- [adv_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/adv_action/): complete subsection reference.

- [all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/all_tcp_traffic/): complete subsection reference.

- [all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/all_traffic/): complete subsection reference.

- [all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/all_udp_traffic/): complete subsection reference.

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/any/): complete subsection reference.

- [applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/applications/): complete subsection reference.

- [inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/inside_endpoints/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/ip_prefix_set/): complete subsection reference.

- [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/label_matcher/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/label_selector/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/metadata/): complete subsection reference.

- [outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/outside_endpoints/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/prefix_list/): complete subsection reference.

- [protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/protocol_port_range/): complete subsection reference.

## Next pages

- [rules.ingress_rules.adv_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/adv_action/)
- [rules.ingress_rules.all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/all_tcp_traffic/)
- [rules.ingress_rules.all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/all_traffic/)
- [rules.ingress_rules.all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/all_udp_traffic/)
- [rules.ingress_rules.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/any/)
- [rules.ingress_rules.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/applications/)
- [rules.ingress_rules.inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/inside_endpoints/)
- [rules.ingress_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/ip_prefix_set/)
- [rules.ingress_rules.label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/label_matcher/)
- [rules.ingress_rules.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/label_selector/)
- [rules.ingress_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/metadata/)
- [rules.ingress_rules.outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/outside_endpoints/)
- [rules.ingress_rules.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/prefix_list/)
- [rules.ingress_rules.protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/protocol_port_range/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/)
- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
