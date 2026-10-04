---
page_title: "dedicated_interface"
subcategory: ""
description: "Dedicated Interface Configuration."
xcsh_docs: {"aliases": ["dedicated interface"], "body_bytes": 9231, "body_sha256": "sha256:891d1b03aac56375f5b4cfae7e1957fe5a8c408f4e4ebacce504fd20c570b8e2", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_interface:properties:dedicated_interface:cluster", "xcsh-docs:resources:network_interface:properties:dedicated_interface:is_primary", "xcsh-docs:resources:network_interface:properties:dedicated_interface:monitor", "xcsh-docs:resources:network_interface:properties:dedicated_interface:monitor_disabled", "xcsh-docs:resources:network_interface:properties:dedicated_interface:not_primary"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:dedicated_interface", "parent_id": "xcsh-docs:resources:network_interface:reference", "path": "documentation/resources/network_interface/properties/dedicated_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2202312132221001-2113331212010103-2130332010130032-2110133310110321-3231313332220130-1300110303021031-1333223321003031-2113330102231203", "registry_path": "docs/guides/resources--network_interface--reference--group-001.md", "relationships": [{"anchor": "schema-dedicated_interface--node", "enforcement": "provider-schema", "group": "dedicated_interface:ConflictingObjectAttributes:cluster,node", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dedicated_interface:ConflictingObjectAttributes:cluster,node", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dedicated_interface:ConflictingObjectAttributes:is_primary,not_primary", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:is_primary", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dedicated_interface:ConflictingObjectAttributes:monitor,monitor_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:monitor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dedicated_interface:ConflictingObjectAttributes:monitor,monitor_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:monitor_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dedicated_interface:ConflictingObjectAttributes:is_primary,not_primary", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:not_primary", "type": "conflicts"}, {"anchor": "schema-dedicated_interface--device", "enforcement": "provider-schema", "group": "dedicated_interface:RequiredObjectAttributes:device", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["dedicated_interface"], "schema_version": 1, "sections": [{"aliases": ["cluster scope", "cluster wide interface", "dedicated interface cluster"], "anchor": "section", "description": "Applies dedicated interface configuration cluster-wide across the site rather than to a specific node.", "document_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:cluster", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["dedicated interface device"], "anchor": "schema-dedicated_interface--device", "description": "Name of the device for which interface is configured. Use wwan0 for 4G/LTE.", "document_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["dedicated interface is primary", "primary interface", "set primary"], "anchor": "section", "description": "Designates this dedicated network interface as the primary interface.", "document_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:is_primary", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "is_primary"], "syntax": "attribute", "type": "object"}, {"aliases": ["dedicated interface monitor"], "anchor": "section", "description": "Link Quality Monitoring configuration for a network interface.", "document_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:monitor", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "monitor"], "syntax": "attribute", "type": "object"}, {"aliases": ["dedicated interface monitor disabled", "disable interface monitoring", "disable link quality monitoring", "monitoring disabled"], "anchor": "section", "description": "Disables link quality monitoring on the dedicated network interface.", "document_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:monitor_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "monitor_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["dedicated interface mtu"], "anchor": "schema-dedicated_interface--mtu", "description": "Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between 512 and 9000.", "document_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "mtu"], "syntax": "attribute", "type": "number"}, {"aliases": ["dedicated interface node"], "anchor": "schema-dedicated_interface--node", "description": "Exclusive with Configuration will apply to a device on the given node of the site.", "document_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "node"], "syntax": "attribute", "type": "string"}, {"aliases": ["dedicated interface not primary"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:not_primary", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "not_primary"], "syntax": "attribute", "type": "object"}, {"aliases": ["dedicated interface priority"], "anchor": "schema-dedicated_interface--priority", "description": "Priority of the network interface when multiple network interfaces are present in outside network Greater the value, higher the priority.", "document_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "priority"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/dedicated_interface/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Dedicated Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dedicated_interface

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- dedicated_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dedicated\_interface, dedicated\_management\_interface, ethernet\_interface,
layer2\_interface, tunnel\_interface\] Configuration parameter for dedicated interface.

Upstream description:

Dedicated Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node"),
  validators.ConflictingObjectAttributes("is_primary",
    "not_primary"),
  validators.ConflictingObjectAttributes("monitor",
    "monitor_disabled")}
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
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]"
}
```

OneOf alternatives in this subsection:

- [dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/#section)
- [dedicated_management_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_management_interface/#section)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/#section)
- [layer2_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/layer2_interface/#section)
- [tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dedicated_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/cluster/): complete subsection reference.

<a id="schema-dedicated_interface--device"></a>

### device property

Type: `"string"`. Optional.

Name of the device for which interface is configured. Use wwan0 for 4G/LTE.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [is_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/is_primary/): complete subsection reference.

- [monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/monitor/): complete subsection reference.

- [monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/monitor_disabled/): complete subsection reference.

<a id="schema-dedicated_interface--mtu"></a>

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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="schema-dedicated_interface--node"></a>

### node property

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [not_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/not_primary/): complete subsection reference.

<a id="schema-dedicated_interface--priority"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Next pages

- [dedicated_interface.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/cluster/)
- [dedicated_interface.is_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/is_primary/)
- [dedicated_interface.monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/monitor/)
- [dedicated_interface.monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/monitor_disabled/)
- [dedicated_interface.not_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/not_primary/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
