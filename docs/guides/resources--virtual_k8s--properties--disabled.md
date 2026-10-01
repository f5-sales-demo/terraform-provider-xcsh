---
page_title: "disabled"
subcategory: "Container"
description: "disabled for xcsh_virtual_k8s."
xcsh_docs: {"aliases": [], "body_bytes": 1129, "body_sha256": "sha256:a838af36be3ae661cb452823dc91a76edc9dacda1f00e4612fc54026628942a4", "canonical_id": "xcsh-docs:resources:virtual_k8s:properties:disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_k8s:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_k8s:properties:disabled", "parent_id": "xcsh-docs:resources:virtual_k8s:reference", "path": "docs/guides/resources--virtual_k8s--properties--disabled.md", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_k8s/properties/disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disabled for xcsh_virtual_k8s.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disabled

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md)
- [Property reference](resources--virtual_k8s--reference.md)
- disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disabled, isolated\] Enable this option

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

- [disabled](resources--virtual_k8s--properties--disabled.md#section)
- [isolated](resources--virtual_k8s--properties--isolated.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--virtual_k8s--reference.md)
- [xcsh_virtual_k8s](../resources/virtual_k8s.md)
