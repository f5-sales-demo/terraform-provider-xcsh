---
page_title: "disable_ocsp_stapling"
subcategory: "Security"
description: "disable_ocsp_stapling for xcsh_certificate."
xcsh_docs: {"aliases": [], "body_bytes": 1132, "body_sha256": "sha256:faa46a4a2d9e8e77ca283f6f99c2e2df8ea430049ba5a7afe58b72a14f8b9337", "child_ids": [], "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate:properties:disable_ocsp_stapling", "parent_id": "xcsh-docs:resources:certificate:reference", "path": "documentation/resources/certificate/properties/disable_ocsp_stapling/index.md", "provider_name": "certificate", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["disable_ocsp_stapling"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/properties/disable_ocsp_stapling/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_ocsp_stapling for xcsh_certificate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["certificateCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_ocsp_stapling

Breadcrumbs:

- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/)
- disable_ocsp_stapling

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/)
- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/)
