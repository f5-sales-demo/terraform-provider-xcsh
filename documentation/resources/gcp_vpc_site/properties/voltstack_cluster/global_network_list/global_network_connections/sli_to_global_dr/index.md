---
page_title: "voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr"
subcategory: "Infrastructure"
description: "Global network reference for direct connection."
xcsh_docs: {"aliases": ["voltstack cluster global network list global network connections sli to global dr"], "body_bytes": 2308, "body_sha256": "sha256:2474148c69b50ba57d70e4e8ddcc02ab57262fe9639fca7e7b0b8a5c6974607f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list:global_network_connections:sli_to_global_dr:global_vn"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list:global_network_connections:sli_to_global_dr", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list:global_network_connections", "path": "documentation/resources/gcp_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1223021200110313-3032022112320302-3132320102320320-1100033103003330-1123230122103223-0233010300122312-3002131022303033-3100211113112320", "registry_path": "docs/guides/resources--gcp_vpc_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "global_network_list", "global_network_connections", "sli_to_global_dr"], "schema_version": 1, "sections": [{"aliases": ["voltstack cluster global network list global network connections sli to global dr global vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list:global_network_connections:sli_to_global_dr:global_vn", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name", "enforcement": "provider-schema", "group": "voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list:global_network_connections:sli_to_global_dr:global_vn", "type": "requires"}], "schema_path": ["voltstack_cluster", "global_network_list", "global_network_connections", "sli_to_global_dr", "global_vn"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Global network reference for direct connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/)
- [voltstack_cluster.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/global_network_list/)
- [voltstack_cluster.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr

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

- [global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/global_vn/): complete subsection reference.

## Next pages

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/global_vn/)
- [voltstack_cluster.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
