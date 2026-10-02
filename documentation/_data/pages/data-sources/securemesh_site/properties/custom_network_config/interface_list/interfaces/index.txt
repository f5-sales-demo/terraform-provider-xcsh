---
page_title: "custom_network_config.interface_list.interfaces"
subcategory: ""
description: "Configure network interfaces for this Secure Mesh site."
xcsh_docs: {"aliases": ["custom network config interface list interfaces"], "body_bytes": 5985, "body_sha256": "sha256:8639d4bebff6092946dc0fd17bb64f61aad10ecaac56afdd7c02e58a6e9f17d1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dc_cluster_group_connectivity_interface_disabled", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dc_cluster_group_connectivity_interface_enabled", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_management_interface", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list", "path": "documentation/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221", "registry_path": "docs/guides/data-sources--securemesh_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces"], "schema_version": 1, "sections": [{"aliases": ["dc cluster group connectivity interface disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dc_cluster_group_connectivity_interface_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dc_cluster_group_connectivity_interface_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["dc cluster group connectivity interface enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dc_cluster_group_connectivity_interface_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dc_cluster_group_connectivity_interface_enabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["dedicated interface"], "anchor": "section", "description": "Dedicated Interface Configuration.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["dedicated management interface"], "anchor": "section", "description": "Dedicated Interface Configuration.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_management_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_management_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["description spec"], "anchor": "schema-custom_network_config--interface_list--interfaces--description_spec", "description": "Interface Description. Description for this Interface.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["ethernet interface"], "anchor": "section", "description": "Ethernet Interface Configuration.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-custom_network_config--interface_list--interfaces--labels", "description": "Add Labels for this Interface, these labels can be used in firewall policy.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "labels"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configure network interfaces for this Secure Mesh site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/)
- custom_network_config.interface_list.interfaces

<a id="section"></a>

Type: `"list"`. Computed.

Configure network interfaces for this Secure Mesh site.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [dc_cluster_group_connectivity_interface_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dc_cluster_group_connectivity_interface_disabled/): complete subsection reference.

- [dc_cluster_group_connectivity_interface_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dc_cluster_group_connectivity_interface_enabled/): complete subsection reference.

- [dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/): complete subsection reference.

- [dedicated_management_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_management_interface/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Add Labels for this Interface, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

## Next pages

- [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dc_cluster_group_connectivity_interface_disabled/)
- [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dc_cluster_group_connectivity_interface_enabled/)
- [custom_network_config.interface_list.interfaces.dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/)
- [custom_network_config.interface_list.interfaces.dedicated_management_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_management_interface/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
