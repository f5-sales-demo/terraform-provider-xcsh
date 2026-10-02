---
page_title: "custom_network_config.interface_list"
subcategory: ""
description: "Configure network interfaces for this Secure Mesh site."
xcsh_docs: {"aliases": ["custom network config interface list"], "body_bytes": 1793, "body_sha256": "sha256:79e6409a7ba251ff02a088d3dd4252706f1b59581f10a9611726b5d2d7ec76c2", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config", "path": "documentation/resources/securemesh_site/properties/custom_network_config/interface_list/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131", "registry_path": "docs/guides/resources--securemesh_site--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list:RequiredObjectAttributes:interfaces", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list"], "schema_version": 1, "sections": [{"aliases": ["interfaces"], "anchor": "section", "description": "Configure network interfaces for this Secure Mesh site.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces:ConflictingListObjectAttributes:dc_cluster_group_connectivity_interface_disabled,dc_cluster_group_connectivity_interface_enabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dc_cluster_group_connectivity_interface_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces:ConflictingListObjectAttributes:dc_cluster_group_connectivity_interface_disabled,dc_cluster_group_connectivity_interface_enabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dc_cluster_group_connectivity_interface_enabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces:ConflictingListObjectAttributes:dedicated_interface,dedicated_management_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces:ConflictingListObjectAttributes:dedicated_interface,ethernet_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces:ConflictingListObjectAttributes:dedicated_interface,dedicated_management_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_management_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces:ConflictingListObjectAttributes:dedicated_management_interface,ethernet_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:dedicated_management_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces:ConflictingListObjectAttributes:dedicated_interface,ethernet_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces:ConflictingListObjectAttributes:dedicated_management_interface,ethernet_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "type": "conflicts"}], "schema_path": ["custom_network_config", "interface_list", "interfaces"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/interface_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure network interfaces for this Secure Mesh site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/)
- custom_network_config.interface_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure network interfaces for this Secure Mesh site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces")}
```

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
interface_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
