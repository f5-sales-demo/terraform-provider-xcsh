---
page_title: "where.virtual_site.enable_internet_vip"
subcategory: ""
description: "where.virtual_site.enable_internet_vip for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1069, "body_sha256": "sha256:3f88701df963fa11955624ee1f4c7c8695ef9ae05408753a1da9face33dccd4c", "canonical_id": "xcsh-docs:resources:discovery:properties:where:virtual_site:enable_internet_vip", "child_ids": [], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:where:virtual_site:enable_internet_vip", "parent_id": "xcsh-docs:resources:discovery:properties:where:virtual_site", "path": "docs/guides/resources--discovery--properties--where--virtual_site--enable_internet_vip.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["where", "virtual_site", "enable_internet_vip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/where/virtual_site/enable_internet_vip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where.virtual_site.enable_internet_vip for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.virtual_site.enable_internet_vip

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md)
- [Property reference](resources--discovery--reference.md)
- [where](resources--discovery--properties--where.md)
- [where.virtual_site](resources--discovery--properties--where--virtual_site.md)
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

- [where.virtual_site](resources--discovery--properties--where--virtual_site.md)
- [xcsh_discovery](../resources/discovery.md)
