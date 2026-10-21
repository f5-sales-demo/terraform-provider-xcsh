---
page_title: "equinix.not_managed.node_list.interface_list.ethernet_interface"
subcategory: ""
description: "Configuration parameter for ethernet interface."
xcsh_docs: {"aliases": ["equinix not managed node list interface list ethernet interface"], "body_bytes": 3413, "body_sha256": "sha256:e6f9046e9b7e2e932e956c7edf770a514e9d40fd4a7fdaa53ec0f60808177ac5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ethernet_interface", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list", "path": "documentation/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/ethernet_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0030022103120212-2330122332133101-3323000322033223-2203313012323233-1122132021130130-2021010032223303-1312302223202331-0031202333110110", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "ethernet_interface"], "schema_version": 1, "sections": [{"aliases": ["equinix not managed node list interface list ethernet interface device"], "anchor": "schema-equinix--not_managed--node_list--interface_list--ethernet_interface--device", "description": "Select an Ethernet device from the discovered interfaces to configure. Once configured, this interface will be part of this sites dataplane and can participate in the networking services configured on this site.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ethernet_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "ethernet_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["equinix not managed node list interface list ethernet interface mac"], "anchor": "schema-equinix--not_managed--node_list--interface_list--ethernet_interface--mac", "description": "Configuration parameter for mac", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ethernet_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "ethernet_interface", "mac"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/ethernet_interface/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Configuration parameter for ethernet interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# equinix.not_managed.node_list.interface_list.ethernet_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [equinix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/)
- [equinix.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/not_managed/)
- [equinix.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/)
- [equinix.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/)
- equinix.not_managed.node_list.interface_list.ethernet_interface

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

<a id="schema-equinix--not_managed--node_list--interface_list--ethernet_interface--device"></a>

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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

<a id="schema-equinix--not_managed--node_list--interface_list--ethernet_interface--mac"></a>

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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
