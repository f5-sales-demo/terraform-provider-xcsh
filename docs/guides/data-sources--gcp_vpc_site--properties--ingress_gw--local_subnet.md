---
page_title: "ingress_gw.local_subnet"
subcategory: "Infrastructure"
description: "ingress_gw.local_subnet for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1461, "body_sha256": "sha256:33bd40ac6f0a06965c2fe3311c54c64f2bd37a3487bd1db22288f00a4ea5dec7", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:local_subnet", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:local_subnet:existing_subnet", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:local_subnet:new_subnet"], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:local_subnet", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw", "path": "docs/guides/data-sources--gcp_vpc_site--properties--ingress_gw--local_subnet.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_gw", "local_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/ingress_gw/local_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_gw.local_subnet for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw.local_subnet

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [ingress_gw](data-sources--gcp_vpc_site--properties--ingress_gw.md)
- ingress_gw.local_subnet

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

- [existing_subnet](data-sources--gcp_vpc_site--properties--ingress_gw--local_subnet--existing_subnet.md): complete subsection reference.

- [new_subnet](data-sources--gcp_vpc_site--properties--ingress_gw--local_subnet--new_subnet.md): complete subsection reference.

## Next pages

- [ingress_gw.local_subnet.existing_subnet](data-sources--gcp_vpc_site--properties--ingress_gw--local_subnet--existing_subnet.md)
- [ingress_gw.local_subnet.new_subnet](data-sources--gcp_vpc_site--properties--ingress_gw--local_subnet--new_subnet.md)
- [ingress_gw](data-sources--gcp_vpc_site--properties--ingress_gw.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
