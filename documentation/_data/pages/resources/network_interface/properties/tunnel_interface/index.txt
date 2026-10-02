---
page_title: "tunnel_interface"
subcategory: ""
description: "Tunnel Interface Configuration."
xcsh_docs: {"aliases": ["tunnel interface"], "body_bytes": 6479, "body_sha256": "sha256:c62faa8a82a8cc1f3a7a39725e65ca0b16143315aabd615f8f9a272397dcd754", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_interface:properties:tunnel_interface:site_local_inside_network", "xcsh-docs:resources:network_interface:properties:tunnel_interface:site_local_network", "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip", "xcsh-docs:resources:network_interface:properties:tunnel_interface:tunnel"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:tunnel_interface", "parent_id": "xcsh-docs:resources:network_interface:reference", "path": "documentation/resources/network_interface/properties/tunnel_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0031231223100101-0221202300220330-0013311301301203-3301232310130302-2330031103323032-2221301010223301-3302220320030311-3202002232312200", "registry_path": "docs/guides/resources--network_interface--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tunnel_interface:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:site_local_inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tunnel_interface:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:site_local_network", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tunnel_interface"], "schema_version": 1, "sections": [{"aliases": ["mtu"], "anchor": "schema-tunnel_interface--mtu", "description": "Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between 512 and 9000.", "document_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tunnel_interface", "mtu"], "syntax": "attribute", "type": "number"}, {"aliases": ["node"], "anchor": "schema-tunnel_interface--node", "description": "Exclusive with Configuration will apply to a given device on the given node.", "document_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tunnel_interface", "node"], "syntax": "attribute", "type": "string"}, {"aliases": ["priority"], "anchor": "schema-tunnel_interface--priority", "description": "Priority of the network interface when multiple network interfaces are present in outside network Greater the value, higher the priority.", "document_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tunnel_interface", "priority"], "syntax": "attribute", "type": "number"}, {"aliases": ["site local inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:site_local_inside_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tunnel_interface", "site_local_inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["site local network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:site_local_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tunnel_interface", "site_local_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["static ip"], "anchor": "section", "description": "Configure Static IP parameters.", "document_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tunnel_interface.static_ip:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip:cluster_static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tunnel_interface.static_ip:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip:node_static_ip", "type": "conflicts"}], "schema_path": ["tunnel_interface", "static_ip"], "syntax": "block", "type": "object"}, {"aliases": ["tunnel"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:tunnel", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-tunnel_interface--tunnel--name", "enforcement": "provider-schema", "group": "tunnel_interface.tunnel:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:tunnel", "type": "requires"}], "schema_path": ["tunnel_interface", "tunnel"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/tunnel_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Tunnel Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tunnel_interface

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- tunnel_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tunnel interface.

Upstream description:

Tunnel Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"site_local_inside_network\",\"site_local_network\"]",
  "x-ves-oneof-field-node_choice": "[\"node\"]"
}
```

Terraform syntax:

```terraform
tunnel_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-tunnel_interface--mtu"></a>

### mtu property

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="schema-tunnel_interface--node"></a>

### node property

Type: `"string"`. Optional.

Exclusive with \[\] Configuration will apply to a given device on the given node.

Upstream description:

Exclusive with \[\] Configuration will apply to a given device on the given node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-tunnel_interface--priority"></a>

### priority property

Type: `"number"`. Optional.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/site_local_inside_network/): complete subsection reference.

- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/site_local_network/): complete subsection reference.

- [static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/static_ip/): complete subsection reference.

- [tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/tunnel/): complete subsection reference.

## Next pages

- [tunnel_interface.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/site_local_inside_network/)
- [tunnel_interface.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/site_local_network/)
- [tunnel_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/static_ip/)
- [tunnel_interface.tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
