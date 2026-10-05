---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config"
subcategory: ""
description: "IPV6AutoConfigType."
xcsh_docs: {"aliases": ["custom network config interface list interfaces ethernet interface ipv6 auto config"], "body_bytes": 2986, "body_sha256": "sha256:80565b4f62e352da1337e6033f674e847685b7fdc9b1e163f3b220c215abd868", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:host", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "path": "documentation/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3210322111333323-3121032003032012-0123101311213322-1333013110031211-3333033113032213-1021032301121210-3011211311030010-0120000021311131", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config"], "schema_version": 1, "sections": [{"aliases": ["custom network config interface list interfaces ethernet interface ipv6 auto config host"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:host", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "host"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config interface list interfaces ethernet interface ipv6 auto config router"], "anchor": "section", "description": "IPV6AutoConfigRouterType.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "IPV6AutoConfigType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config

<a id="section"></a>

Type: `"single"`. Computed.

IPV6AutoConfigType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

## Direct properties

- [host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/host/): complete subsection reference.

- [router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/host/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
