---
page_title: "custom_network_config.interface_list.interfaces.dedicated_interface"
subcategory: ""
description: "Dedicated Interface Configuration."
xcsh_docs: {"aliases": ["custom network config interface list interfaces dedicated interface"], "body_bytes": 8719, "body_sha256": "sha256:edab25c09f727b76919673e5849fb29ae46d7aef2100267b74481821228a684e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface:cluster", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface:is_primary", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface:monitor", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface:monitor_disabled", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface:not_primary"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces", "path": "documentation/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1321311131000311-3032322101312330-2130120123203032-1213303320033230-2200220223330031-3020321320330023-2301111022100223-0021121201331131", "registry_path": "docs/guides/data-sources--securemesh_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_interface"], "schema_version": 1, "sections": [{"aliases": ["custom network config interface list interfaces dedicated interface cluster"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface:cluster", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_interface", "cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config interface list interfaces dedicated interface device"], "anchor": "schema-custom_network_config--interface_list--interfaces--dedicated_interface--device", "description": "Name of the device for which interface is configured. Use wwan0 for 4G/LTE.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom network config interface list interfaces dedicated interface is primary"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface:is_primary", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_interface", "is_primary"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config interface list interfaces dedicated interface monitor"], "anchor": "section", "description": "Link Quality Monitoring configuration for a network interface.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface:monitor", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_interface", "monitor"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config interface list interfaces dedicated interface monitor disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface:monitor_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_interface", "monitor_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config interface list interfaces dedicated interface mtu"], "anchor": "schema-custom_network_config--interface_list--interfaces--dedicated_interface--mtu", "description": "Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between 512 and 9000.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_interface", "mtu"], "syntax": "attribute", "type": "number"}, {"aliases": ["custom network config interface list interfaces dedicated interface node"], "anchor": "schema-custom_network_config--interface_list--interfaces--dedicated_interface--node", "description": "Exclusive with Configuration will apply to a device on the given node of the site.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_interface", "node"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom network config interface list interfaces dedicated interface not primary"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface:not_primary", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_interface", "not_primary"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config interface list interfaces dedicated interface priority"], "anchor": "schema-custom_network_config--interface_list--interfaces--dedicated_interface--priority", "description": "Priority of the network interface when multiple network interfaces are present in outside network Greater the value, higher the priority.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_interface", "priority"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Dedicated Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.dedicated_interface

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/)
- custom_network_config.interface_list.interfaces.dedicated_interface

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for dedicated interface.

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

## Direct properties

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/cluster/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--dedicated_interface--device"></a>

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

- [is_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/is_primary/): complete subsection reference.

- [monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/monitor/): complete subsection reference.

- [monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/monitor_disabled/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--dedicated_interface--mtu"></a>

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

<a id="schema-custom_network_config--interface_list--interfaces--dedicated_interface--node"></a>

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

- [not_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/not_primary/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--dedicated_interface--priority"></a>

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

- [custom_network_config.interface_list.interfaces.dedicated_interface.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/cluster/)
- [custom_network_config.interface_list.interfaces.dedicated_interface.is_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/is_primary/)
- [custom_network_config.interface_list.interfaces.dedicated_interface.monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/monitor/)
- [custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/monitor_disabled/)
- [custom_network_config.interface_list.interfaces.dedicated_interface.not_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/not_primary/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
