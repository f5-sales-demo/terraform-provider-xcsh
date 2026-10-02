---
page_title: "tunnel_interface.static_ip"
subcategory: ""
description: "Configure Static IP parameters."
xcsh_docs: {"aliases": ["tunnel interface static ip"], "body_bytes": 2317, "body_sha256": "sha256:a5b423a8903b1eea0e36fc0196d2a909ff8b5152ce8a774a7284601b1448c5e7", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip:cluster_static_ip", "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip:node_static_ip"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip", "parent_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface", "path": "documentation/resources/network_interface/properties/tunnel_interface/static_ip/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3000301202120000-2120302233320331-1032223120223111-2120030310022133-0301211132223101-1120232202101331-0120120333222230-2033213332133001", "registry_path": "docs/guides/resources--network_interface--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tunnel_interface.static_ip:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip:cluster_static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tunnel_interface.static_ip:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip:node_static_ip", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tunnel_interface", "static_ip"], "schema_version": 1, "sections": [{"aliases": ["cluster static ip"], "anchor": "section", "description": "Configure Static IP parameters for cluster.", "document_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip:cluster_static_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tunnel_interface", "static_ip", "cluster_static_ip"], "syntax": "block", "type": "object"}, {"aliases": ["node static ip"], "anchor": "section", "description": "Configure Static IP parameters for a node.", "document_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip:node_static_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-tunnel_interface--static_ip--node_static_ip--ip_address", "enforcement": "provider-schema", "group": "tunnel_interface.static_ip.node_static_ip:RequiredObjectAttributes:ip_address", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:static_ip:node_static_ip", "type": "requires"}], "schema_path": ["tunnel_interface", "static_ip", "node_static_ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/tunnel_interface/static_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure Static IP parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tunnel_interface.static_ip

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/)
- tunnel_interface.static_ip

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

- [cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/static_ip/cluster_static_ip/): complete subsection reference.

- [node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/static_ip/node_static_ip/): complete subsection reference.

## Next pages

- [tunnel_interface.static_ip.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/static_ip/cluster_static_ip/)
- [tunnel_interface.static_ip.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/static_ip/node_static_ip/)
- [tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
