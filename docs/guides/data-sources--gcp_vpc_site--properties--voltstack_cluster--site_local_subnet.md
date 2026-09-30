---
page_title: "voltstack_cluster.site_local_subnet"
subcategory: "Infrastructure"
description: "voltstack_cluster.site_local_subnet for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1486, "body_sha256": "sha256:c10776e5968bbdf41e131427449f6f5897686c00853115f54edcd5df5259af85", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet:existing_subnet", "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet:new_subnet"], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster", "path": "docs/guides/data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_subnet.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "site_local_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/voltstack_cluster/site_local_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.site_local_subnet for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster.site_local_subnet

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md)
- voltstack_cluster.site_local_subnet

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
  "x-ves-oneof-field-choice": "[\"existing_subnet\",\"new_subnet\"]"
}
```

## Direct properties

- [existing_subnet](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_subnet--existing_subnet.md): complete subsection reference.

- [new_subnet](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_subnet--new_subnet.md): complete subsection reference.

## Next pages

- [voltstack_cluster.site_local_subnet.existing_subnet](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_subnet--existing_subnet.md)
- [voltstack_cluster.site_local_subnet.new_subnet](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_subnet--new_subnet.md)
- [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
