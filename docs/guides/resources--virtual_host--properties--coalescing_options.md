---
page_title: "coalescing_options"
subcategory: ""
description: "coalescing_options for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1624, "body_sha256": "sha256:0b173dc747e1f1ae5caf82db247e85ef65444b68b141004e48f025984ac1d629", "canonical_id": "xcsh-docs:resources:virtual_host:properties:coalescing_options", "child_ids": ["xcsh-docs:resources:virtual_host:properties:coalescing_options:default_coalescing", "xcsh-docs:resources:virtual_host:properties:coalescing_options:strict_coalescing"], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:coalescing_options", "parent_id": "xcsh-docs:resources:virtual_host:reference", "path": "docs/guides/resources--virtual_host--properties--coalescing_options.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["coalescing_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/coalescing_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "coalescing_options for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# coalescing_options

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- coalescing_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_coalescing](resources--virtual_host--properties--coalescing_options--default_coalescing.md): complete subsection reference.

- [strict_coalescing](resources--virtual_host--properties--coalescing_options--strict_coalescing.md): complete subsection reference.

## Next pages

- [coalescing_options.default_coalescing](resources--virtual_host--properties--coalescing_options--default_coalescing.md)
- [coalescing_options.strict_coalescing](resources--virtual_host--properties--coalescing_options--strict_coalescing.md)
- [Property reference](resources--virtual_host--reference.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
