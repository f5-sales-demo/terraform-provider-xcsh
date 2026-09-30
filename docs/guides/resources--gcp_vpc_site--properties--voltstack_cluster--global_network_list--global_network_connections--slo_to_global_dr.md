---
page_title: "voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr"
subcategory: "Infrastructure"
description: "voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1761, "body_sha256": "sha256:26d623ae336fc48f9c132ecd449b9f36e1a359514c8903ba4e40a7e6145aadc3", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list:global_network_connections:slo_to_global_dr", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list:global_network_connections:slo_to_global_dr:global_vn"], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list:global_network_connections:slo_to_global_dr", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list:global_network_connections", "path": "docs/guides/resources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "global_network_list", "global_network_connections", "slo_to_global_dr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/voltstack_cluster/global_network_list/global_network_connections/slo_to_global_dr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [voltstack_cluster](resources--gcp_vpc_site--properties--voltstack_cluster.md)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--properties--voltstack_cluster--global_network_list.md)
- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections.md)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr

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

- [global_vn](resources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md): complete subsection reference.

## Next pages

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md)
- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
