---
page_title: "custom_network_config.interface_list.interfaces"
subcategory: ""
description: "Configure network interfaces for this App Stack site."
xcsh_docs: {"aliases": ["custom network config interface list interfaces"], "body_bytes": 5811, "body_sha256": "sha256:69a7d5a7732262dfcbb90baea6de2f0279a75f40c26ee99453c9e255c78aa923", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dc_cluster_group_connectivity_interface_disabled", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dc_cluster_group_connectivity_interface_enabled", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dedicated_management_interface", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:labels", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list", "path": "documentation/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2010011200312102-2233032202322312-0001211312111230-1020312322330000-2331013310320330-2111110021331001-1222132122030233-2203100110212020", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces"], "schema_version": 1, "sections": [{"aliases": ["dc cluster group connectivity interface disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dc_cluster_group_connectivity_interface_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dc_cluster_group_connectivity_interface_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["dc cluster group connectivity interface enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dc_cluster_group_connectivity_interface_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dc_cluster_group_connectivity_interface_enabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["dedicated interface"], "anchor": "section", "description": "Dedicated Interface Configuration.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["dedicated management interface"], "anchor": "section", "description": "Dedicated Interface Configuration.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:dedicated_management_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "dedicated_management_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["description spec"], "anchor": "schema-custom_network_config--interface_list--interfaces--description_spec", "description": "Interface Description. Description for this Interface.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["ethernet interface"], "anchor": "section", "description": "Ethernet Interface Configuration.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "section", "description": "Add Labels for this Interface, these labels can be used in firewall policy.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:labels", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["tunnel interface"], "anchor": "section", "description": "Tunnel Interface Configuration.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "tunnel_interface"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure network interfaces for this App Stack site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/)
- custom_network_config.interface_list.interfaces

<a id="section"></a>

Type: `"list"`. Computed.

Configure network interfaces for this App Stack site.

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

- [dc_cluster_group_connectivity_interface_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dc_cluster_group_connectivity_interface_disabled/): complete subsection reference.

- [dc_cluster_group_connectivity_interface_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dc_cluster_group_connectivity_interface_enabled/): complete subsection reference.

- [dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/): complete subsection reference.

- [dedicated_management_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dedicated_management_interface/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/): complete subsection reference.

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/labels/): complete subsection reference.

- [tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dc_cluster_group_connectivity_interface_disabled/)
- [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dc_cluster_group_connectivity_interface_enabled/)
- [custom_network_config.interface_list.interfaces.dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/)
- [custom_network_config.interface_list.interfaces.dedicated_management_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/dedicated_management_interface/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [custom_network_config.interface_list.interfaces.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/labels/)
- [custom_network_config.interface_list.interfaces.tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
