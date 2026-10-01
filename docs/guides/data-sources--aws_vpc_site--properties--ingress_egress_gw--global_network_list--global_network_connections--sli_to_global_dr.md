---
page_title: "ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr"
subcategory: "Infrastructure"
description: "ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1762, "body_sha256": "sha256:0164e2c881571b2c86d82efdd2d5751db575c5fd46493efc1ece7a6e3f15b419", "canonical_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:global_network_list:global_network_connections:sli_to_global_dr", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:global_network_list:global_network_connections:sli_to_global_dr:global_vn"], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:global_network_list:global_network_connections:sli_to_global_dr", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:global_network_list:global_network_connections", "path": "docs/guides/data-sources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "global_network_list", "global_network_connections", "sli_to_global_dr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
- [Property reference](data-sources--aws_vpc_site--reference.md)
- [ingress_egress_gw](data-sources--aws_vpc_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--properties--ingress_egress_gw--global_network_list.md)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections.md)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [global_vn](data-sources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--properties--ingress_egress_gw--global_network_list--global_network_connections.md)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
