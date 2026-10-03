---
page_title: "ethernet_interface.static_ip"
subcategory: ""
description: "Configure Static IP parameters."
xcsh_docs: {"aliases": ["ethernet interface static ip"], "body_bytes": 2341, "body_sha256": "sha256:fffea350575f0f01e8838e20dbd223f744e84d5d9432b9bc39bbc23218833da8", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ip:cluster_static_ip", "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ip:node_static_ip"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ip", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface", "path": "documentation/resources/network_interface/properties/ethernet_interface/static_ip/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3022232102130132-0113030032123321-1012123313020210-2230003031231301-1032311210332033-2222203030010301-2030212000100303-0203232233101020", "registry_path": "docs/guides/resources--network_interface--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.static_ip:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ip:cluster_static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.static_ip:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ip:node_static_ip", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ethernet_interface", "static_ip"], "schema_version": 1, "sections": [{"aliases": ["ethernet interface static ip cluster static ip"], "anchor": "section", "description": "Configure Static IP parameters for cluster.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ip:cluster_static_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ethernet_interface", "static_ip", "cluster_static_ip"], "syntax": "block", "type": "object"}, {"aliases": ["ethernet interface static ip node static ip"], "anchor": "section", "description": "Configure Static IP parameters for a node.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ip:node_static_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ethernet_interface--static_ip--node_static_ip--ip_address", "enforcement": "provider-schema", "group": "ethernet_interface.static_ip.node_static_ip:RequiredObjectAttributes:ip_address", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ip:node_static_ip", "type": "requires"}], "schema_path": ["ethernet_interface", "static_ip", "node_static_ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/static_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure Static IP parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.static_ip

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/)
- ethernet_interface.static_ip

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

- [cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/static_ip/cluster_static_ip/): complete subsection reference.

- [node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/static_ip/node_static_ip/): complete subsection reference.

## Next pages

- [ethernet_interface.static_ip.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/static_ip/cluster_static_ip/)
- [ethernet_interface.static_ip.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/static_ip/node_static_ip/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
