---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.static_ip"
subcategory: ""
description: "Configure Static IP parameters."
xcsh_docs: {"aliases": ["custom network config interface list interfaces ethernet interface static ip"], "body_bytes": 3406, "body_sha256": "sha256:3431eeede582c725bab29268ec2d20d0872721f6b0a036b986cd1fc1964f15b7", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip:cluster_static_ip", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip:node_static_ip"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "path": "documentation/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0312201332321112-1103221103313331-3101303123320303-2132101111111321-3300301332023210-1313031321303022-3112221000023123-0030002230201023", "registry_path": "docs/guides/resources--securemesh_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.static_ip:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip:cluster_static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.static_ip:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip:node_static_ip", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "static_ip"], "schema_version": 1, "sections": [{"aliases": ["custom network config interface list interfaces ethernet interface static ip cluster static ip"], "anchor": "section", "description": "Configure Static IP parameters for cluster.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip:cluster_static_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "static_ip", "cluster_static_ip"], "syntax": "block", "type": "object"}, {"aliases": ["custom network config interface list interfaces ethernet interface static ip node static ip"], "anchor": "section", "description": "Configure Static IP parameters for a node.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip:node_static_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip--ip_address", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip:RequiredObjectAttributes:ip_address", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip:node_static_ip", "type": "requires"}], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "static_ip", "node_static_ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure Static IP parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.ethernet_interface.static_ip

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
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
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ip {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/cluster_static_ip/): complete subsection reference.

- [node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/node_static_ip/): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/cluster_static_ip/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/node_static_ip/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
