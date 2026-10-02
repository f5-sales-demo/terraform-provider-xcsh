---
page_title: "kvm.not_managed.node_list.interface_list.network_option"
subcategory: ""
description: "Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs, Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is optional. Global VRFs are"
xcsh_docs: {"aliases": ["kvm not managed node list interface list network option"], "body_bytes": 3774, "body_sha256": "sha256:fa4db705a5cb91747602a1a784b652b03dda1ddaa09e7341082fe7f52ba478de", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:network_option:site_local_inside_network", "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:network_option:site_local_network"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:network_option", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list", "path": "documentation/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/network_option/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2111012111300220-1101022333120303-1210013102220311-3130230133301023-2230000210021021-3003013233210230-3223322330002111-2123130321000113", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-011.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.network_option:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:network_option:site_local_inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.network_option:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:network_option:site_local_network", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "network_option"], "schema_version": 1, "sections": [{"aliases": ["site local inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:network_option:site_local_inside_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "network_option", "site_local_inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["site local network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:network_option:site_local_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "network_option", "site_local_network"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/network_option/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs, Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is optional. Global VRFs are", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kvm.not_managed.node_list.interface_list.network_option

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [kvm](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/)
- [kvm.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/)
- [kvm.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/)
- [kvm.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/)
- kvm.not_managed.node_list.interface_list.network_option

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional.

Upstream description:

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
network_option {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/network_option/site_local_inside_network/): complete subsection reference.

- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/network_option/site_local_network/): complete subsection reference.

## Next pages

- [kvm.not_managed.node_list.interface_list.network_option.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/network_option/site_local_inside_network/)
- [kvm.not_managed.node_list.interface_list.network_option.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/network_option/site_local_network/)
- [kvm.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
