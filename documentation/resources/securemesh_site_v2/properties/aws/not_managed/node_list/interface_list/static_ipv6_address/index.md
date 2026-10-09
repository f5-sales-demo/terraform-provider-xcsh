---
page_title: "aws.not_managed.node_list.interface_list.static_ipv6_address"
subcategory: ""
description: "Configure Static IP parameters."
xcsh_docs: {"aliases": ["aws not managed node list interface list static ipv6 address"], "body_bytes": 2249, "body_sha256": "sha256:16aa4cb1e6876e03c371f35903cc5f8d131a13d7bfd9c9f612c533df5f54a286", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:static_ipv6_address", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list", "path": "documentation/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/static_ipv6_address/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0022010121023230-1010112111012231-3100032313003023-0321102200222000-0131001132132311-0202202301120022-1131012000222230-3100232030323303", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws.not_managed.node_list.interface_list.static_ipv6_address:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws.not_managed.node_list.interface_list.static_ipv6_address:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws", "not_managed", "node_list", "interface_list", "static_ipv6_address"], "schema_version": 1, "sections": [{"aliases": ["aws not managed node list interface list static ipv6 address cluster static ip"], "anchor": "section", "description": "Configure Static IP parameters for cluster.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws", "not_managed", "node_list", "interface_list", "static_ipv6_address", "cluster_static_ip"], "syntax": "block", "type": "object"}, {"aliases": ["aws not managed node list interface list static ipv6 address node static ip"], "anchor": "section", "description": "Configure Static IP parameters for a node.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--ip_address", "enforcement": "provider-schema", "group": "aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip:RequiredObjectAttributes:ip_address", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "type": "requires"}], "schema_path": ["aws", "not_managed", "node_list", "interface_list", "static_ipv6_address", "node_static_ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/static_ipv6_address/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Configure Static IP parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.not_managed.node_list.interface_list.static_ipv6_address

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/)
- [aws.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/)
- [aws.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/)
- [aws.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/)
- aws.not_managed.node_list.interface_list.static_ipv6_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
static_ipv6_address {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/static_ipv6_address/cluster_static_ip/): complete subsection reference.

- [node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/static_ipv6_address/node_static_ip/): complete subsection reference.
