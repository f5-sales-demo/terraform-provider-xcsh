---
page_title: "voltstack_cluster.site_local_network"
subcategory: "Infrastructure"
description: "voltstack_cluster.site_local_network for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1976, "body_sha256": "sha256:bc71804e016a0fe5f90bbcec9eff56a923315dc1cdb05c22e5d609352f0eef2a", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_network", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:existing_network", "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network", "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network_autogenerate"], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_network", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster", "path": "docs/guides/data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "site_local_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.site_local_network for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.site_local_network

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md)
- voltstack_cluster.site_local_network

<a id="section"></a>

Type: `"single"`. Computed.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_network\",\"new_network\",\"new_network_autogenerate\"]"
}
```

## Direct properties

- [existing_network](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network--existing_network.md): complete subsection reference.

- [new_network](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network--new_network.md): complete subsection reference.

- [new_network_autogenerate](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network--new_network_autogenerate.md): complete subsection reference.

## Next pages

- [voltstack_cluster.site_local_network.existing_network](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network--existing_network.md)
- [voltstack_cluster.site_local_network.new_network](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network--new_network.md)
- [voltstack_cluster.site_local_network.new_network_autogenerate](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network--new_network_autogenerate.md)
- [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
