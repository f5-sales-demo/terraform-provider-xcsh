---
page_title: "openstack.not_managed.node_list.interface_list.ethernet_interface"
subcategory: ""
description: "Configuration parameter for ethernet interface."
xcsh_docs: {"aliases": ["openstack not managed node list interface list ethernet interface"], "body_bytes": 3437, "body_sha256": "sha256:bd6dd83346a21946d202dbb295be55203a2f50f853f74b82328a2d6d39eda0cd", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ethernet_interface", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list", "path": "documentation/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/ethernet_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2033333113321230-0122003033301003-1302133221203002-3112230201223110-0223201223311010-3320113322202022-2310231230003111-0312220200021023", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "ethernet_interface"], "schema_version": 1, "sections": [{"aliases": ["openstack not managed node list interface list ethernet interface device"], "anchor": "schema-openstack--not_managed--node_list--interface_list--ethernet_interface--device", "description": "Select an Ethernet device from the discovered interfaces to configure. Once configured, this interface will be part of this sites dataplane and can participate in the networking services configured on this site.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ethernet_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "ethernet_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["openstack not managed node list interface list ethernet interface mac"], "anchor": "schema-openstack--not_managed--node_list--interface_list--ethernet_interface--mac", "description": "Configuration parameter for mac", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ethernet_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "ethernet_interface", "mac"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/ethernet_interface/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Configuration parameter for ethernet interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openstack.not_managed.node_list.interface_list.ethernet_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [openstack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openstack/)
- [openstack.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openstack/not_managed/)
- [openstack.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/)
- [openstack.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/)
- openstack.not_managed.node_list.interface_list.ethernet_interface

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for ethernet interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-openstack--not_managed--node_list--interface_list--ethernet_interface--device"></a>

### device property

Type: `"string"`. Computed.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-openstack--not_managed--node_list--interface_list--ethernet_interface--mac"></a>

### mac property

Type: `"string"`. Computed.

MAC Address. Configuration parameter for mac

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```
