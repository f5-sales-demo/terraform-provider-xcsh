---
page_title: "vn_config.global_network_list.global_network_connections.sli_to_global_dr"
subcategory: ""
description: "Global network reference for direct connection."
xcsh_docs: {"aliases": ["vn config global network list global network connections sli to global dr"], "body_bytes": 2204, "body_sha256": "sha256:6ad97cef660c522be77d2591a18254094e5ec989582b737a1a94dd5a3be35c0d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:sli_to_global_dr:global_vn"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:sli_to_global_dr", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections", "path": "documentation/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/sli_to_global_dr/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1323030132301113-1001220033133222-0222213000121020-1102103131200123-0013200000333010-1310021020300320-3032100002321120-2030322100201310", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "global_network_list", "global_network_connections", "sli_to_global_dr"], "schema_version": 1, "sections": [{"aliases": ["global vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:sli_to_global_dr:global_vn", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name", "enforcement": "provider-schema", "group": "vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:sli_to_global_dr:global_vn", "type": "requires"}], "schema_path": ["vn_config", "global_network_list", "global_network_connections", "sli_to_global_dr", "global_vn"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/sli_to_global_dr/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Global network reference for direct connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.global_network_list.global_network_connections.sli_to_global_dr

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- [vn_config.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/)
- [vn_config.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/)
- vn_config.global_network_list.global_network_connections.sli_to_global_dr

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

- [global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/sli_to_global_dr/global_vn/): complete subsection reference.

## Next pages

- [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/sli_to_global_dr/global_vn/)
- [vn_config.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
