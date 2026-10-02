---
page_title: "voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr"
subcategory: "Infrastructure"
description: "Global network reference for direct connection."
xcsh_docs: {"aliases": ["voltstack cluster ar global network list global network connections slo to global dr"], "body_bytes": 2380, "body_sha256": "sha256:b6e7d3aaf7abe5b2266470a27152ff97671696ad5aac9990d12a22f248c1e676", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:global_network_list:global_network_connections:slo_to_global_dr:global_vn"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:global_network_list:global_network_connections:slo_to_global_dr", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:global_network_list:global_network_connections", "path": "documentation/resources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/slo_to_global_dr/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2130022003231131-2220112201210232-3303113132120110-0013022132310131-0000330100023211-0333110310003311-1031302221011330-2322220330022330", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster_ar", "global_network_list", "global_network_connections", "slo_to_global_dr"], "schema_version": 1, "sections": [{"aliases": ["global vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:global_network_list:global_network_connections:slo_to_global_dr:global_vn", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name", "enforcement": "provider-schema", "group": "voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:global_network_list:global_network_connections:slo_to_global_dr:global_vn", "type": "requires"}], "schema_path": ["voltstack_cluster_ar", "global_network_list", "global_network_connections", "slo_to_global_dr", "global_vn"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/slo_to_global_dr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Global network reference for direct connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/)
- [voltstack_cluster_ar.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/)
- [voltstack_cluster_ar.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/)
- voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

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
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

## Direct properties

- [global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/slo_to_global_dr/global_vn/): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/slo_to_global_dr/global_vn/)
- [voltstack_cluster_ar.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
