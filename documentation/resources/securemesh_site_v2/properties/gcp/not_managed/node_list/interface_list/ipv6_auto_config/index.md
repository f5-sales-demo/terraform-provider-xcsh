---
page_title: "gcp.not_managed.node_list.interface_list.ipv6_auto_config"
subcategory: ""
description: "IPV6AutoConfigType."
xcsh_docs: {"aliases": ["gcp not managed node list interface list ipv6 auto config"], "body_bytes": 2857, "body_sha256": "sha256:f848ecf6f3652622ea70d3f0562d1379aafb4702868867bf37e09926d67c3ed2", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config:host", "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config:router"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list", "path": "documentation/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "gcp.not_managed.node_list.interface_list.ipv6_auto_config:ConflictingObjectAttributes:host,router", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config:host", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp.not_managed.node_list.interface_list.ipv6_auto_config:ConflictingObjectAttributes:host,router", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config:router", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp", "not_managed", "node_list", "interface_list", "ipv6_auto_config"], "schema_version": 1, "sections": [{"aliases": ["gcp not managed node list interface list ipv6 auto config host"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config:host", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "host"], "syntax": "attribute", "type": "object"}, {"aliases": ["gcp not managed node list interface list ipv6 auto config router"], "anchor": "section", "description": "IPV6AutoConfigRouterType.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config:router", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-gcp--not_managed--node_list--interface_list--ipv6_auto_config--router--network_prefix", "enforcement": "provider-schema", "group": "gcp.not_managed.node_list.interface_list.ipv6_auto_config.router:ConflictingObjectAttributes:network_prefix,stateful", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config:router", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp.not_managed.node_list.interface_list.ipv6_auto_config.router:ConflictingObjectAttributes:network_prefix,stateful", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "type": "conflicts"}], "schema_path": ["gcp", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "IPV6AutoConfigType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp.not_managed.node_list.interface_list.ipv6_auto_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/)
- [gcp.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/)
- [gcp.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/)
- [gcp.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("host",
    "router")}
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
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

Terraform syntax:

```terraform
ipv6_auto_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/host/): complete subsection reference.

- [router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/router/): complete subsection reference.

## Next pages

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/host/)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [gcp.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
