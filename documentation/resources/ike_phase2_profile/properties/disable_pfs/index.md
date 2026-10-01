---
page_title: "disable_pfs"
subcategory: ""
description: "disable_pfs for xcsh_ike_phase2_profile."
xcsh_docs: {"aliases": [], "body_bytes": 1134, "body_sha256": "sha256:c5143e465e2fd54c730129a6a971949cdbc1d98fe083476957e24fe58d4f2192", "child_ids": [], "collection_id": "xcsh-docs:resources:ike_phase2_profile:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike_phase2_profile:properties:disable_pfs", "parent_id": "xcsh-docs:resources:ike_phase2_profile:reference", "path": "documentation/resources/ike_phase2_profile/properties/disable_pfs/index.md", "provider_name": "ike_phase2_profile", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["disable_pfs"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase2_profile/properties/disable_pfs/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_pfs for xcsh_ike_phase2_profile.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase2_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_pfs

Breadcrumbs:

- [xcsh_ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/)
- disable_pfs

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable pfs.

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
disable_pfs = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/)
- [xcsh_ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/)
