---
page_title: "enable"
subcategory: ""
description: "enable for xcsh_segment."
xcsh_docs: {"aliases": [], "body_bytes": 823, "body_sha256": "sha256:5b0e6a774a7c888beb28b1832f351ded553fd4fbcc210f80265fa4661f28cf83", "canonical_id": "xcsh-docs:resources:segment:properties:enable", "child_ids": [], "collection_id": "xcsh-docs:resources:segment:collection", "completeness": "complete", "id": "xcsh-docs:resources:segment:properties:enable", "parent_id": "xcsh-docs:resources:segment:reference", "path": "docs/guides/resources--segment--properties--enable.md", "provider_name": "segment", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/segment/properties/enable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable for xcsh_segment.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["segmentCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable

Breadcrumbs:

- [xcsh_segment](../resources/segment.md)
- [Property reference](resources--segment--reference.md)
- enable

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
enable = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--segment--reference.md)
- [xcsh_segment](../resources/segment.md)
