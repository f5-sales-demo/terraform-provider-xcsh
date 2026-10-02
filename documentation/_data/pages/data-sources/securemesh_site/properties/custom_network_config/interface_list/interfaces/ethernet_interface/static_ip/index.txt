---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.static_ip"
subcategory: ""
description: "Configure Static IP parameters."
xcsh_docs: {"aliases": ["custom network config interface list interfaces ethernet interface static ip"], "body_bytes": 3143, "body_sha256": "sha256:a3d53a3376cc829378b655c980cf588e16a975069a9905a3d89ee1d98af909b4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip:cluster_static_ip", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip:node_static_ip"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "path": "documentation/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0120113002313030-1322302200111002-3301101312033012-1002300312011100-2030033212113011-1201131333202300-1312130301100302-2022121022000212", "registry_path": "docs/guides/data-sources--securemesh_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "static_ip"], "schema_version": 1, "sections": [{"aliases": ["cluster static ip"], "anchor": "section", "description": "Configure Static IP parameters for cluster.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip:cluster_static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "static_ip", "cluster_static_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["node static ip"], "anchor": "section", "description": "Configure Static IP parameters for a node.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip:node_static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "static_ip", "node_static_ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configure Static IP parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.ethernet_interface.static_ip

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip

<a id="section"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

## Direct properties

- [cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/cluster_static_ip/): complete subsection reference.

- [node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/node_static_ip/): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/cluster_static_ip/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/node_static_ip/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
