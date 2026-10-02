---
page_title: "dedicated_interface"
subcategory: ""
description: "Dedicated Interface Configuration."
xcsh_docs: {"aliases": ["dedicated interface"], "body_bytes": 8133, "body_sha256": "sha256:cb9a9b161c872d0687044de23aae153c5a52a61fcb4b2029725f73cc50bc12c0", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:dedicated_interface:cluster", "xcsh-docs:data-sources:network_interface:properties:dedicated_interface:is_primary", "xcsh-docs:data-sources:network_interface:properties:dedicated_interface:monitor", "xcsh-docs:data-sources:network_interface:properties:dedicated_interface:monitor_disabled", "xcsh-docs:data-sources:network_interface:properties:dedicated_interface:not_primary"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:dedicated_interface", "parent_id": "xcsh-docs:data-sources:network_interface:reference", "path": "documentation/data-sources/network_interface/properties/dedicated_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321", "registry_path": "docs/guides/data-sources--network_interface--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dedicated_interface"], "schema_version": 1, "sections": [{"aliases": ["cluster"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_interface:cluster", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["device"], "anchor": "schema-dedicated_interface--device", "description": "Name of the device for which interface is configured. Use wwan0 for 4G/LTE.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["is primary"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_interface:is_primary", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "is_primary"], "syntax": "attribute", "type": "object"}, {"aliases": ["monitor"], "anchor": "section", "description": "Link Quality Monitoring configuration for a network interface.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_interface:monitor", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "monitor"], "syntax": "attribute", "type": "object"}, {"aliases": ["monitor disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_interface:monitor_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "monitor_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["mtu"], "anchor": "schema-dedicated_interface--mtu", "description": "Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between 512 and 9000.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "mtu"], "syntax": "attribute", "type": "number"}, {"aliases": ["node"], "anchor": "schema-dedicated_interface--node", "description": "Exclusive with Configuration will apply to a device on the given node of the site.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "node"], "syntax": "attribute", "type": "string"}, {"aliases": ["not primary"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_interface:not_primary", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "not_primary"], "syntax": "attribute", "type": "object"}, {"aliases": ["priority"], "anchor": "schema-dedicated_interface--priority", "description": "Priority of the network interface when multiple network interfaces are present in outside network Greater the value, higher the priority.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_interface", "priority"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/dedicated_interface/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Dedicated Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dedicated_interface

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/)
- dedicated_interface

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: dedicated\_interface, dedicated\_management\_interface, ethernet\_interface,
layer2\_interface, tunnel\_interface\] Configuration parameter for dedicated interface.

Upstream description:

Dedicated Interface Configuration.

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

- [dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/#section)
- [dedicated_management_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_management_interface/#section)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/#section)
- [layer2_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/#section)
- [tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/cluster/): complete subsection reference.

<a id="schema-dedicated_interface--device"></a>

### device property

Type: `"string"`. Computed.

Name of the device for which interface is configured. Use wwan0 for 4G/LTE.

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

- [is_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/is_primary/): complete subsection reference.

- [monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/monitor/): complete subsection reference.

- [monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/monitor_disabled/): complete subsection reference.

<a id="schema-dedicated_interface--mtu"></a>

### mtu property

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

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

<a id="schema-dedicated_interface--node"></a>

### node property

Type: `"string"`. Computed.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

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

- [not_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/not_primary/): complete subsection reference.

<a id="schema-dedicated_interface--priority"></a>

### priority property

Type: `"number"`. Computed.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

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

## Next pages

- [dedicated_interface.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/cluster/)
- [dedicated_interface.is_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/is_primary/)
- [dedicated_interface.monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/monitor/)
- [dedicated_interface.monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/monitor_disabled/)
- [dedicated_interface.not_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/not_primary/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
