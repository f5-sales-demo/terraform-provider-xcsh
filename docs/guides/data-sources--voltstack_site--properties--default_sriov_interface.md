---
page_title: "default_sriov_interface"
subcategory: ""
description: "default_sriov_interface for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1292, "body_sha256": "sha256:45c2ad78eefd18e81e497f62dbb30b79fd7dc7e7dad4c1ebf5562d86bd4002dc", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:default_sriov_interface", "child_ids": [], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:default_sriov_interface", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "docs/guides/data-sources--voltstack_site--properties--default_sriov_interface.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_sriov_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/default_sriov_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_sriov_interface for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_sriov_interface

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- default_sriov_interface

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_sriov\_interface, sriov\_interfaces; Default: default\_sriov\_interface\]
Configuration parameter for default sriov interface.

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

- [default_sriov_interface](data-sources--voltstack_site--properties--default_sriov_interface.md#section)
- [sriov_interfaces](data-sources--voltstack_site--properties--sriov_interfaces.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--voltstack_site--reference.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
