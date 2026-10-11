---
page_title: "network_pbr.network_pbr_rules"
subcategory: ""
description: "Network(L3/L4) routing policy rule."
xcsh_docs: {"aliases": ["network pbr network pbr rules"], "body_bytes": 6413, "body_sha256": "sha256:007d7f12c542ebcc5340e3cbed84f23b6fa7639bb1f8ca685a29212fd0242223", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_tcp_traffic", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_traffic", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_udp_traffic", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:any", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:applications", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:forwarding_class_list", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:metadata", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:prefix_list", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:protocol_port_range"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules", "parent_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr", "path": "documentation/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000", "registry_path": "docs/guides/resources--policy_based_routing--reference--group-001.md", "relationships": [{"anchor": "schema-network_pbr--network_pbr_rules--dns_name", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:any,dns_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules", "type": "conflicts"}, {"anchor": "schema-network_pbr--network_pbr_rules--dns_name", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:dns_name,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules", "type": "conflicts"}, {"anchor": "schema-network_pbr--network_pbr_rules--dns_name", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:dns_name,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_tcp_traffic,all_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_tcp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_tcp_traffic,all_udp_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_tcp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_tcp_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_tcp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_tcp_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_tcp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_tcp_traffic,all_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_traffic,all_udp_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_tcp_traffic,all_udp_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_udp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_traffic,all_udp_traffic", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_udp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_udp_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_udp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_udp_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_udp_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:any,dns_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:any,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:any,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_tcp_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:applications", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:applications", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_udp_traffic,applications", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:applications", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:applications,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:applications", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:any,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:dns_name,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:any,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:dns_name,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_tcp_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:protocol_port_range", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:protocol_port_range", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:all_udp_traffic,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:protocol_port_range", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:ConflictingListObjectAttributes:applications,protocol_port_range", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:protocol_port_range", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules:RequiredListObjectAttributes:forwarding_class_list", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:forwarding_class_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["network_pbr", "network_pbr_rules"], "schema_version": 1, "sections": [{"aliases": ["network pbr network pbr rules all tcp traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_tcp_traffic", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "all_tcp_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["network pbr network pbr rules all traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_traffic", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "all_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["network pbr network pbr rules all udp traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_udp_traffic", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "all_udp_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["network pbr network pbr rules any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:any", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "any"], "syntax": "attribute", "type": "object"}, {"aliases": ["network pbr network pbr rules applications"], "anchor": "section", "description": "Application protocols like HTTP, SNMP.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:applications", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "applications"], "syntax": "block", "type": "object"}, {"aliases": ["network pbr network pbr rules dns name"], "anchor": "schema-network_pbr--network_pbr_rules--dns_name", "description": "Exclusive with Resolve hostname to GET the IP.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "dns_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["network pbr network pbr rules forwarding class list"], "anchor": "section", "description": "Ordered list of forwarding Class to be used if rule match.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:forwarding_class_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-network_pbr--network_pbr_rules--forwarding_class_list--name", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules.forwarding_class_list:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:forwarding_class_list", "type": "requires"}], "schema_path": ["network_pbr", "network_pbr_rules", "forwarding_class_list"], "syntax": "block", "type": "object"}, {"aliases": ["network pbr network pbr rules ip prefix set"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "ip_prefix_set"], "syntax": "block", "type": "object"}, {"aliases": ["network pbr network pbr rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-network_pbr--network_pbr_rules--metadata--name", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:metadata", "type": "requires"}], "schema_path": ["network_pbr", "network_pbr_rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["network pbr network pbr rules prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:prefix_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "prefix_list"], "syntax": "block", "type": "object"}, {"aliases": ["network pbr network pbr rules protocol port range"], "anchor": "section", "description": "Protocol and Port ranges.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:protocol_port_range", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "protocol_port_range"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Network(L3/L4) routing policy rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_pbr.network_pbr_rules

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/)
- [network_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/)
- network_pbr.network_pbr_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

L3/L4 Destination Routing Rules. Network(L3/L4) routing policy rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("forwarding_class_list"),
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
  validators.ConflictingListObjectAttributes("any",
    "dns_name"),
  validators.ConflictingListObjectAttributes("any",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("applications",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("dns_name",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("dns_name",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
network_pbr_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/all_tcp_traffic/): complete subsection reference.

- [all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/all_traffic/): complete subsection reference.

- [all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/all_udp_traffic/): complete subsection reference.

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/any/): complete subsection reference.

- [applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/applications/): complete subsection reference.

<a id="schema-network_pbr--network_pbr_rules--dns_name"></a>

### dns_name property

Type: `"string"`. Optional.

Exclusive with \[any ip\_prefix\_set prefix\_list\] Resolve hostname to GET the IP.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [forwarding_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/forwarding_class_list/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/metadata/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/prefix_list/): complete subsection reference.

- [protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/protocol_port_range/): complete subsection reference.
