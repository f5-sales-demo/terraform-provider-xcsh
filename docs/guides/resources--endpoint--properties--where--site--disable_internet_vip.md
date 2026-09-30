---
page_title: "where.site.disable_internet_vip"
subcategory: "Networking"
description: "where.site.disable_internet_vip for xcsh_endpoint."
xcsh_docs: {"aliases": [], "body_bytes": 917, "body_sha256": "sha256:0e73426db82be62f6e96a7e2e8dd313b0eb18a9bdc6cb7f495e2fb4cfa5111b8", "canonical_id": "xcsh-docs:resources:endpoint:properties:where:site:disable_internet_vip", "child_ids": [], "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:properties:where:site:disable_internet_vip", "parent_id": "xcsh-docs:resources:endpoint:properties:where:site", "path": "docs/guides/resources--endpoint--properties--where--site--disable_internet_vip.md", "provider_name": "endpoint", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["where", "site", "disable_internet_vip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/properties/where/site/disable_internet_vip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "where.site.disable_internet_vip for xcsh_endpoint.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# where.site.disable_internet_vip

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md)
- [Property reference](resources--endpoint--reference.md)
- [where](resources--endpoint--properties--where.md)
- [where.site](resources--endpoint--properties--where--site.md)
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

- [where.site](resources--endpoint--properties--where--site.md)
- [xcsh_endpoint](../resources/endpoint.md)
