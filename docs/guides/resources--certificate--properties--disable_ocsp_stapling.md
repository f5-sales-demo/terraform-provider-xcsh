---
page_title: "disable_ocsp_stapling"
subcategory: "Security"
description: "disable_ocsp_stapling for xcsh_certificate."
xcsh_docs: {"aliases": [], "body_bytes": 924, "body_sha256": "sha256:c37353034e06c364d0d1234e742ae32e39ca9576f83fa1b061b9b72c18bf02d0", "canonical_id": "xcsh-docs:resources:certificate:properties:disable_ocsp_stapling", "child_ids": [], "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate:properties:disable_ocsp_stapling", "parent_id": "xcsh-docs:resources:certificate:reference", "path": "docs/guides/resources--certificate--properties--disable_ocsp_stapling.md", "provider_name": "certificate", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_ocsp_stapling"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/properties/disable_ocsp_stapling/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_ocsp_stapling for xcsh_certificate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["certificateCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_ocsp_stapling

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md)
- [Property reference](resources--certificate--reference.md)
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

- [Property reference](resources--certificate--reference.md)
- [xcsh_certificate](../resources/certificate.md)
