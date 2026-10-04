---
page_title: "peers.external.interface_list"
subcategory: ""
description: "List of network interfaces."
xcsh_docs: {"aliases": ["peers external interface list"], "body_bytes": 1761, "body_sha256": "sha256:6a94b8df6983f59ebab7a43897424a50a7e0cfc1a4f4e35177c476fecafd6539", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:external:interface_list:interfaces"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:external:interface_list", "parent_id": "xcsh-docs:resources:bgp:properties:peers:external", "path": "documentation/resources/bgp/properties/peers/external/interface_list/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0221000322321030-1113211202122102-0133131202033231-1331110121001123-0111220333300112-3122311132013122-0321033031322012-3310302200131001", "registry_path": "docs/guides/resources--bgp--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "peers.external.interface_list:RequiredObjectAttributes:interfaces", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:interface_list:interfaces", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "external", "interface_list"], "schema_version": 1, "sections": [{"aliases": ["peers external interface list interfaces"], "anchor": "section", "description": "List of network interfaces.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:interface_list:interfaces", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-peers--external--interface_list--interfaces--name", "enforcement": "provider-schema", "group": "peers.external.interface_list.interfaces:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:interface_list:interfaces", "type": "requires"}], "schema_path": ["peers", "external", "interface_list", "interfaces"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/external/interface_list/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "List of network interfaces.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["bgpCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.interface_list

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/)
- peers.external.interface_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

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

- [interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/interface_list/interfaces/): complete subsection reference.

## Next pages

- [peers.external.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/interface_list/interfaces/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
