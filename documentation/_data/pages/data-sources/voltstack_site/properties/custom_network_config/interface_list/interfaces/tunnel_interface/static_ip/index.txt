---
page_title: "custom_network_config.interface_list.interfaces.tunnel_interface.static_ip"
subcategory: ""
description: "Configure Static IP parameters."
xcsh_docs: {"aliases": ["custom network config interface list interfaces tunnel interface static ip"], "body_bytes": 3105, "body_sha256": "sha256:43ffb640dbe1e8fb5150479813bcdda2fc0d33def30af0ad57ee9682afe30e87", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:static_ip:cluster_static_ip", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:static_ip:node_static_ip"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:static_ip", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface", "path": "documentation/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/static_ip/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1202133130330021-2320221232032130-1330321313212302-1231112032213110-0012312132200011-0212033122110023-2230012300211012-1231002301200010", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "tunnel_interface", "static_ip"], "schema_version": 1, "sections": [{"aliases": ["cluster static ip"], "anchor": "section", "description": "Configure Static IP parameters for cluster.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:static_ip:cluster_static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "tunnel_interface", "static_ip", "cluster_static_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["node static ip"], "anchor": "section", "description": "Configure Static IP parameters for a node.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:static_ip:node_static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "tunnel_interface", "static_ip", "node_static_ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/static_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure Static IP parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.tunnel_interface.static_ip

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config.interface_list.interfaces.tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/)
- custom_network_config.interface_list.interfaces.tunnel_interface.static_ip

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

- [cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/static_ip/cluster_static_ip/): complete subsection reference.

- [node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/static_ip/node_static_ip/): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/static_ip/cluster_static_ip/)
- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/static_ip/node_static_ip/)
- [custom_network_config.interface_list.interfaces.tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
