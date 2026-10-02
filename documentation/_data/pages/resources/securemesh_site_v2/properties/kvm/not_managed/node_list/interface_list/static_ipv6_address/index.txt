---
page_title: "kvm.not_managed.node_list.interface_list.static_ipv6_address"
subcategory: ""
description: "Configure Static IP parameters."
xcsh_docs: {"aliases": ["kvm not managed node list interface list static ipv6 address"], "body_bytes": 3104, "body_sha256": "sha256:b20d76a49ce6a52423bcae280af17416df6d83b4520ce09a68350aa813b67ff2", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:static_ipv6_address", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list", "path": "documentation/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/static_ipv6_address/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2333320220330222-3223331321123102-1230203212000212-2200200003213010-2321301123233220-2002223331120023-3032023201233013-2313023210133321", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-011.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.static_ipv6_address:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.static_ipv6_address:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "static_ipv6_address"], "schema_version": 1, "sections": [{"aliases": ["cluster static ip"], "anchor": "section", "description": "Configure Static IP parameters for cluster.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "static_ipv6_address", "cluster_static_ip"], "syntax": "block", "type": "object"}, {"aliases": ["node static ip"], "anchor": "section", "description": "Configure Static IP parameters for a node.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-kvm--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--ip_address", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip:RequiredObjectAttributes:ip_address", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "type": "requires"}], "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "static_ipv6_address", "node_static_ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/static_ipv6_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configure Static IP parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kvm.not_managed.node_list.interface_list.static_ipv6_address

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [kvm](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/)
- [kvm.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/)
- [kvm.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/)
- [kvm.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/)
- kvm.not_managed.node_list.interface_list.static_ipv6_address

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
static_ipv6_address {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/static_ipv6_address/cluster_static_ip/): complete subsection reference.

- [node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/static_ipv6_address/node_static_ip/): complete subsection reference.

## Next pages

- [kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/static_ipv6_address/cluster_static_ip/)
- [kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/static_ipv6_address/node_static_ip/)
- [kvm.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
