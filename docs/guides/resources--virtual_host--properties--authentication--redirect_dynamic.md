---
page_title: "authentication.redirect_dynamic"
subcategory: ""
description: "authentication.redirect_dynamic for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1028, "body_sha256": "sha256:0c77da2e88aa0ef86a78d608dbb306bcdd721ec22eb2f9a66228e96d0eff1349", "canonical_id": "xcsh-docs:resources:virtual_host:properties:authentication:redirect_dynamic", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:authentication:redirect_dynamic", "parent_id": "xcsh-docs:resources:virtual_host:properties:authentication", "path": "docs/guides/resources--virtual_host--properties--authentication--redirect_dynamic.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["authentication", "redirect_dynamic"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/authentication/redirect_dynamic/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "authentication.redirect_dynamic for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# authentication.redirect_dynamic

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [authentication](resources--virtual_host--properties--authentication.md)
- authentication.redirect_dynamic

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for redirect dynamic.

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
redirect_dynamic = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [authentication](resources--virtual_host--properties--authentication.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
