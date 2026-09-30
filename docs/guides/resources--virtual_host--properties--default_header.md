---
page_title: "default_header"
subcategory: ""
description: "default_header for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 803, "body_sha256": "sha256:8160779f6b5d4524f5659faa27e29292c2bceaddd4f26e1333ce19f1e896683d", "canonical_id": "xcsh-docs:resources:virtual_host:properties:default_header", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:default_header", "parent_id": "xcsh-docs:resources:virtual_host:reference", "path": "docs/guides/resources--virtual_host--properties--default_header.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_header"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/default_header/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_header for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# default_header

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- default_header

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default header.

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
default_header = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--virtual_host--reference.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
