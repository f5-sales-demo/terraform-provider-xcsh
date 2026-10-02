---
page_title: "custom_network_config.global_network_list.global_network_connections.slo_to_global_dr"
subcategory: ""
description: "Global network reference for direct connection."
xcsh_docs: {"aliases": ["custom network config global network list global network connections slo to global dr"], "body_bytes": 2382, "body_sha256": "sha256:e5193a500a1de765afd4b73a660864912b750204735fd3f8e4d361e2bf15f2db", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:slo_to_global_dr:global_vn"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:slo_to_global_dr", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections", "path": "documentation/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3210033333332303-0223103230311322-1003302201030303-1033130112100023-2112130001202000-0103130020010132-0103331000111131-2330030231212132", "registry_path": "docs/guides/resources--voltstack_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "global_network_list", "global_network_connections", "slo_to_global_dr"], "schema_version": 1, "sections": [{"aliases": ["global vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:slo_to_global_dr:global_vn", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name", "enforcement": "provider-schema", "group": "custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:slo_to_global_dr:global_vn", "type": "requires"}], "schema_path": ["custom_network_config", "global_network_list", "global_network_connections", "slo_to_global_dr", "global_vn"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Global network reference for direct connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.global_network_list.global_network_connections.slo_to_global_dr

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/global_network_list/)
- [custom_network_config.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/)
- custom_network_config.global_network_list.global_network_connections.slo_to_global_dr

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

- [global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/global_vn/): complete subsection reference.

## Next pages

- [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/global_vn/)
- [custom_network_config.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
