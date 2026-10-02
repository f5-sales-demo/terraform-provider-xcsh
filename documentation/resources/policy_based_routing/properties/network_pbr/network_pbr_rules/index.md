---
page_title: "network_pbr.network_pbr_rules"
subcategory: ""
description: "Network(L3/L4) routing policy rule."
xcsh_docs: {"aliases": ["network pbr network pbr rules"], "body_bytes": 8701, "body_sha256": "sha256:6537a3eb79cfc7139912b2ce978b7b523d3b23dfc7c74b2d4021c20d963e51b5", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_tcp_traffic", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_traffic", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_udp_traffic", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:any", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:applications", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:forwarding_class_list", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:metadata", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:prefix_list", "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:protocol_port_range"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules", "parent_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr", "path": "documentation/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1333313200330032-3321032201303330-1002002123303122-3101012200333212-1032101131030110-1112013101010232-0110222033230303-1212302032133000", "registry_path": "docs/guides/resources--policy_based_routing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["network_pbr", "network_pbr_rules"], "schema_version": 1, "sections": [{"aliases": ["all tcp traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_tcp_traffic", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "all_tcp_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["all traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_traffic", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "all_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["all udp traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:all_udp_traffic", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "all_udp_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:any", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "any"], "syntax": "attribute", "type": "object"}, {"aliases": ["applications"], "anchor": "section", "description": "Application protocols like HTTP, SNMP.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:applications", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "applications"], "syntax": "block", "type": "object"}, {"aliases": ["dns name"], "anchor": "schema-network_pbr--network_pbr_rules--dns_name", "description": "Exclusive with Resolve hostname to GET the IP.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "dns_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["forwarding class list"], "anchor": "section", "description": "Ordered list of forwarding Class to be used if rule match.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:forwarding_class_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "forwarding_class_list"], "syntax": "block", "type": "object"}, {"aliases": ["ip prefix set"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "ip_prefix_set"], "syntax": "block", "type": "object"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-network_pbr--network_pbr_rules--metadata--name", "enforcement": "provider-schema", "group": "network_pbr.network_pbr_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:metadata", "type": "requires"}], "schema_path": ["network_pbr", "network_pbr_rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:prefix_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "prefix_list"], "syntax": "block", "type": "object"}, {"aliases": ["protocol port range"], "anchor": "section", "description": "Protocol and Port ranges.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:protocol_port_range", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "protocol_port_range"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Network(L3/L4) routing policy rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

Network(L3/L4) routing policy rule.

Provider validators and defaults (from schema source):

```go
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

Upstream description:

Exclusive with \[any ip\_prefix\_set prefix\_list\] Resolve hostname to GET the IP.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [network_pbr.network_pbr_rules.all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/all_tcp_traffic/)
- [network_pbr.network_pbr_rules.all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/all_traffic/)
- [network_pbr.network_pbr_rules.all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/all_udp_traffic/)
- [network_pbr.network_pbr_rules.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/any/)
- [network_pbr.network_pbr_rules.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/applications/)
- [network_pbr.network_pbr_rules.forwarding_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/forwarding_class_list/)
- [network_pbr.network_pbr_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/)
- [network_pbr.network_pbr_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/metadata/)
- [network_pbr.network_pbr_rules.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/prefix_list/)
- [network_pbr.network_pbr_rules.protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/protocol_port_range/)
- [network_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/)
- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
