---
page_title: "voltstack_cluster.active_network_policies"
subcategory: "Infrastructure"
description: "List of firewall policy views."
xcsh_docs: {"aliases": ["voltstack cluster active network policies"], "body_bytes": 1893, "body_sha256": "sha256:c5ac3aafcb4f5e1eb8704b3d05bb0de7929f52b19a8cf7541de315cec102fd72", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:active_network_policies:network_policies"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:active_network_policies", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster", "path": "documentation/resources/azure_vnet_site/properties/voltstack_cluster/active_network_policies/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1133002223210211-1020013333111012-3123303203202133-2010122202220132-2003322120000101-1120101311203110-1021133112011100-3010203031020220", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-008.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.active_network_policies:RequiredObjectAttributes:network_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:active_network_policies:network_policies", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "active_network_policies"], "schema_version": 1, "sections": [{"aliases": ["network policies"], "anchor": "section", "description": "Ordered List of Firewall Policies active for this network firewall.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:active_network_policies:network_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-voltstack_cluster--active_network_policies--network_policies--name", "enforcement": "provider-schema", "group": "voltstack_cluster.active_network_policies.network_policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:active_network_policies:network_policies", "type": "requires"}], "schema_path": ["voltstack_cluster", "active_network_policies", "network_policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster/active_network_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of firewall policy views.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.active_network_policies

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/)
- voltstack_cluster.active_network_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/active_network_policies/network_policies/): complete subsection reference.

## Next pages

- [voltstack_cluster.active_network_policies.network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/active_network_policies/network_policies/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
