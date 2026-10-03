---
page_title: "vn_config"
subcategory: ""
description: "Virtual Network Configuration."
xcsh_docs: {"aliases": ["vn config"], "body_bytes": 6071, "body_sha256": "sha256:27d78dd3c5a13daab9789a2177c694ce92d81ae7c9fb0b3e65940992f604aeb1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:dc_cluster_group_inside_vn", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:dc_cluster_group_outside_vn", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:global_network_list", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:no_dc_cluster_group", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:no_global_network", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:no_inside_static_routes", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:no_outside_static_routes", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:sm_connection_public_ip", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:sm_connection_pvt_ip"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:reference", "path": "documentation/data-sources/aws_tgw_site/properties/vn_config/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config"], "schema_version": 1, "sections": [{"aliases": ["vn config allowed vip port"], "anchor": "section", "description": "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "allowed_vip_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config allowed vip port sli"], "anchor": "section", "description": "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "allowed_vip_port_sli"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config dc cluster group inside vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:dc_cluster_group_inside_vn", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "dc_cluster_group_inside_vn"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config dc cluster group outside vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:dc_cluster_group_outside_vn", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "dc_cluster_group_outside_vn"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config global network list"], "anchor": "section", "description": "List of global network connections.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:global_network_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "global_network_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config inside static routes"], "anchor": "section", "description": "List of static routes.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "inside_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config no dc cluster group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:no_dc_cluster_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "no_dc_cluster_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config no global network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:no_global_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "no_global_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config no inside static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:no_inside_static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "no_inside_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config no outside static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:no_outside_static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "no_outside_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config outside static routes"], "anchor": "section", "description": "List of static routes.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "outside_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config sm connection public ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:sm_connection_public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "sm_connection_public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config sm connection pvt ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:sm_connection_pvt_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "sm_connection_pvt_ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vn_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Virtual Network Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- vn_config

<a id="section"></a>

Type: `"single"`. Computed.

Virtual Network Configuration. Virtual Network Configuration.

Upstream description:

Virtual Network Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

## Direct properties

- [allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port/): complete subsection reference.

- [allowed_vip_port_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/): complete subsection reference.

- [dc_cluster_group_inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/dc_cluster_group_inside_vn/): complete subsection reference.

- [dc_cluster_group_outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/dc_cluster_group_outside_vn/): complete subsection reference.

- [global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/global_network_list/): complete subsection reference.

- [inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/): complete subsection reference.

- [no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/no_dc_cluster_group/): complete subsection reference.

- [no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/no_global_network/): complete subsection reference.

- [no_inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/no_inside_static_routes/): complete subsection reference.

- [no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/no_outside_static_routes/): complete subsection reference.

- [outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/outside_static_routes/): complete subsection reference.

- [sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/sm_connection_public_ip/): complete subsection reference.

- [sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/sm_connection_pvt_ip/): complete subsection reference.

## Next pages

- [vn_config.allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port/)
- [vn_config.allowed_vip_port_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/)
- [vn_config.dc_cluster_group_inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/dc_cluster_group_inside_vn/)
- [vn_config.dc_cluster_group_outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/dc_cluster_group_outside_vn/)
- [vn_config.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/global_network_list/)
- [vn_config.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/)
- [vn_config.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/no_dc_cluster_group/)
- [vn_config.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/no_global_network/)
- [vn_config.no_inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/no_inside_static_routes/)
- [vn_config.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/no_outside_static_routes/)
- [vn_config.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/outside_static_routes/)
- [vn_config.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/sm_connection_public_ip/)
- [vn_config.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/sm_connection_pvt_ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
