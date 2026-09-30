---
page_title: "site_acl.outside_network"
subcategory: ""
description: "site_acl.outside_network for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 861, "body_sha256": "sha256:d6c2cc162fb524a74b478210d65c0107abcfa81546d21d9b4ae27cd749e85085", "canonical_id": "xcsh-docs:resources:fast_acl:properties:site_acl:outside_network", "child_ids": [], "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:site_acl:outside_network", "parent_id": "xcsh-docs:resources:fast_acl:properties:site_acl", "path": "docs/guides/resources--fast_acl--properties--site_acl--outside_network.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_acl", "outside_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/site_acl/outside_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_acl.outside_network for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# site_acl.outside_network

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md)
- [Property reference](resources--fast_acl--reference.md)
- [site_acl](resources--fast_acl--properties--site_acl.md)
- site_acl.outside_network

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [site_acl](resources--fast_acl--properties--site_acl.md)
- [xcsh_fast_acl](../resources/fast_acl.md)
