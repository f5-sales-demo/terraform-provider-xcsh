---
page_title: "ingress_egress_gw.inside_network"
subcategory: "Infrastructure"
description: "ingress_egress_gw.inside_network for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1932, "body_sha256": "sha256:481bda239022da7002aa66558a4de65f14b8883b4c18e2f2f56b0cb33dd35175", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_network", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_network:existing_network", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_network:new_network", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_network:new_network_autogenerate"], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_network", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw", "path": "docs/guides/data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "inside_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.inside_network for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.inside_network

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [ingress_egress_gw](data-sources--gcp_vpc_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.inside_network

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

- [existing_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network--existing_network.md): complete subsection reference.

- [new_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network--new_network.md): complete subsection reference.

- [new_network_autogenerate](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network--new_network_autogenerate.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.inside_network.existing_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network--existing_network.md)
- [ingress_egress_gw.inside_network.new_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network--new_network.md)
- [ingress_egress_gw.inside_network.new_network_autogenerate](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network--new_network_autogenerate.md)
- [ingress_egress_gw](data-sources--gcp_vpc_site--properties--ingress_egress_gw.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
