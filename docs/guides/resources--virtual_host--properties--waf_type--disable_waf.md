---
page_title: "waf_type.disable_waf"
subcategory: ""
description: "waf_type.disable_waf for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 972, "body_sha256": "sha256:73f79fbb4db0cf1179de2f2653f713021b2fc6b23c6f71ddca825978264e2e4a", "canonical_id": "xcsh-docs:resources:virtual_host:properties:waf_type:disable_waf", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:waf_type:disable_waf", "parent_id": "xcsh-docs:resources:virtual_host:properties:waf_type", "path": "docs/guides/resources--virtual_host--properties--waf_type--disable_waf.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_type", "disable_waf"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/waf_type/disable_waf/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_type.disable_waf for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_type.disable_waf

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [waf_type](resources--virtual_host--properties--waf_type.md)
- waf_type.disable_waf

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable waf.

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
disable_waf = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [waf_type](resources--virtual_host--properties--waf_type.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
