---
page_title: "private_connect_disabled"
subcategory: "Infrastructure"
description: "private_connect_disabled for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1272, "body_sha256": "sha256:32494d3bb3d15be00a4dfdf6352644c540372714910339860103f2ebd10cf719", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:private_connect_disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:private_connect_disabled", "parent_id": "xcsh-docs:resources:gcp_vpc_site:reference", "path": "docs/guides/resources--gcp_vpc_site--properties--private_connect_disabled.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["private_connect_disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/private_connect_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "private_connect_disabled for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_connect_disabled

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- private_connect_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: private\_connect\_disabled, private\_connectivity\] Enable this option

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

OneOf alternatives in this subsection:

- [private_connect_disabled](resources--gcp_vpc_site--properties--private_connect_disabled.md#section)
- [private_connectivity](resources--gcp_vpc_site--properties--private_connectivity.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
private_connect_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--gcp_vpc_site--reference.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
