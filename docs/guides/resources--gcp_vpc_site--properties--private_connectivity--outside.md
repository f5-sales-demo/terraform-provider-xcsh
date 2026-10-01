---
page_title: "private_connectivity.outside"
subcategory: "Infrastructure"
description: "private_connectivity.outside for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1010, "body_sha256": "sha256:4e962e36779453cffa50c70ef61576c13f7eaa9b49fa6bed808109bba3110aa1", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity:outside", "child_ids": [], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity:outside", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity", "path": "docs/guides/resources--gcp_vpc_site--properties--private_connectivity--outside.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["private_connectivity", "outside"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/private_connectivity/outside/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "private_connectivity.outside for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_connectivity.outside

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [private_connectivity](resources--gcp_vpc_site--properties--private_connectivity.md)
- private_connectivity.outside

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
outside = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [private_connectivity](resources--gcp_vpc_site--properties--private_connectivity.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
