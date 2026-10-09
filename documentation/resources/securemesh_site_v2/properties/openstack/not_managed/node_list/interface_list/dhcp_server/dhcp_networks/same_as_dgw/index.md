---
page_title: "openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["openstack not managed node list interface list dhcp server dhcp networks same as dgw"], "body_bytes": 2111, "body_sha256": "sha256:1b32a2e4b61eb90edc9ed8ddc28f1ba21a08e8262040e6c3aaa87e21db7568b5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:dhcp_server:dhcp_networks:same_as_dgw", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:dhcp_server:dhcp_networks", "path": "documentation/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/dhcp_server/dhcp_networks/same_as_dgw/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2233032222111100-2020101101123322-1233331313323223-2101330211233213-3102100023130203-2221301312112311-2002220103231000-3101020302332023", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "dhcp_server", "dhcp_networks", "same_as_dgw"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/dhcp_server/dhcp_networks/same_as_dgw/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [openstack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/)
- [openstack.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/)
- [openstack.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/)
- [openstack.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/)
- [openstack.not_managed.node_list.interface_list.dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/dhcp_server/)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/dhcp_server/dhcp_networks/)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for same as dgw.

Additional upstream details:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
same_as_dgw = {}
```

This is an empty object or choice marker. It has no direct properties.
