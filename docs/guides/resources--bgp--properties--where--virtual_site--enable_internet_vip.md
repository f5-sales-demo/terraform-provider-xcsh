---
page_title: "where.virtual_site.enable_internet_vip"
subcategory: ""
description: "where.virtual_site.enable_internet_vip for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1021, "body_sha256": "sha256:2a4cc5fc2b0a30949a4105f687bd33a5609e92cf9318d6233313e04a7509f611", "canonical_id": "xcsh-docs:resources:bgp:properties:where:virtual_site:enable_internet_vip", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:where:virtual_site:enable_internet_vip", "parent_id": "xcsh-docs:resources:bgp:properties:where:virtual_site", "path": "docs/guides/resources--bgp--properties--where--virtual_site--enable_internet_vip.md", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["where", "virtual_site", "enable_internet_vip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/where/virtual_site/enable_internet_vip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where.virtual_site.enable_internet_vip for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.virtual_site.enable_internet_vip

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [where](resources--bgp--properties--where.md)
- [where.virtual_site](resources--bgp--properties--where--virtual_site.md)
- where.virtual_site.enable_internet_vip

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
enable_internet_vip = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [where.virtual_site](resources--bgp--properties--where--virtual_site.md)
- [xcsh_bgp](../resources/bgp.md)
