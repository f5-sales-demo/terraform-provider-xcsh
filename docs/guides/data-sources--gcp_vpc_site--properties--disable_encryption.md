---
page_title: "disable_encryption"
subcategory: "Infrastructure"
description: "disable_encryption for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1143, "body_sha256": "sha256:d46cf6972dcc63a3f47bbcd9ac4b47829b766d239455db33c129d917ebcbdd29", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:disable_encryption", "child_ids": [], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:disable_encryption", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:reference", "path": "docs/guides/data-sources--gcp_vpc_site--properties--disable_encryption.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_encryption"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/disable_encryption/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_encryption for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# disable_encryption

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- disable_encryption

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_encryption, enable\_encryption; Default: disable\_encryption\] Configuration
parameter for disable encryption.

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

- [disable_encryption](data-sources--gcp_vpc_site--properties--disable_encryption.md#section)
- [enable_encryption](data-sources--gcp_vpc_site--properties--enable_encryption.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
