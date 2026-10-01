---
page_title: "use_system_defaults"
subcategory: "Security"
description: "use_system_defaults for xcsh_certificate."
xcsh_docs: {"aliases": [], "body_bytes": 916, "body_sha256": "sha256:8b2b822dbe016871be41fa997ba8efb9c23eba3ed7f6dc2a3431b0998f705f95", "canonical_id": "xcsh-docs:resources:certificate:properties:use_system_defaults", "child_ids": [], "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate:properties:use_system_defaults", "parent_id": "xcsh-docs:resources:certificate:reference", "path": "docs/guides/resources--certificate--properties--use_system_defaults.md", "provider_name": "certificate", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_system_defaults"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/properties/use_system_defaults/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_system_defaults for xcsh_certificate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["certificateCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_system_defaults

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md)
- [Property reference](resources--certificate--reference.md)
- use_system_defaults

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--certificate--reference.md)
- [xcsh_certificate](../resources/certificate.md)
