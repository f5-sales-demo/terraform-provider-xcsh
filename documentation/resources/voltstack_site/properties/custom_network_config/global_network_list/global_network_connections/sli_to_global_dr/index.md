---
page_title: "custom_network_config.global_network_list.global_network_connections.sli_to_global_dr"
subcategory: ""
description: "Global network reference for direct connection."
xcsh_docs: {"aliases": ["custom network config global network list global network connections sli to global dr"], "body_bytes": 2382, "body_sha256": "sha256:1e470872c849eb0947444c8b11556538564cb4758b501fd80963d4abd40ef6f1", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:sli_to_global_dr:global_vn"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:sli_to_global_dr", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections", "path": "documentation/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/sli_to_global_dr/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0022323003132331-2112201221033031-2013120232233110-0020330200200102-2111003032232110-1022131132121003-2120102001221333-0133110231330333", "registry_path": "docs/guides/resources--voltstack_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "global_network_list", "global_network_connections", "sli_to_global_dr"], "schema_version": 1, "sections": [{"aliases": ["custom network config global network list global network connections sli to global dr global vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:sli_to_global_dr:global_vn", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name", "enforcement": "provider-schema", "group": "custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:sli_to_global_dr:global_vn", "type": "requires"}], "schema_path": ["custom_network_config", "global_network_list", "global_network_connections", "sli_to_global_dr", "global_vn"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/sli_to_global_dr/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Global network reference for direct connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.global_network_list.global_network_connections.sli_to_global_dr

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/global_network_list/)
- [custom_network_config.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/)
- custom_network_config.global_network_list.global_network_connections.sli_to_global_dr

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
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

## Direct properties

- [global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/sli_to_global_dr/global_vn/): complete subsection reference.

## Next pages

- [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/sli_to_global_dr/global_vn/)
- [custom_network_config.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
