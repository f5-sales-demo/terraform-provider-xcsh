---
page_title: "where.site.disable_internet_vip"
subcategory: ""
description: "where.site.disable_internet_vip for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 976, "body_sha256": "sha256:6e35178eac3fa262cc8f9e9bc67aefa76601ef1856b89dea22ae67fdd1d4d193", "canonical_id": "xcsh-docs:resources:bgp:properties:where:site:disable_internet_vip", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:where:site:disable_internet_vip", "parent_id": "xcsh-docs:resources:bgp:properties:where:site", "path": "docs/guides/resources--bgp--properties--where--site--disable_internet_vip.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["where", "site", "disable_internet_vip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/where/site/disable_internet_vip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where.site.disable_internet_vip for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.site.disable_internet_vip

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [where](resources--bgp--properties--where.md)
- [where.site](resources--bgp--properties--where--site.md)
- where.site.disable_internet_vip

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
disable_internet_vip = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [where.site](resources--bgp--properties--where--site.md)
- [xcsh_bgp](../resources/bgp.md)
