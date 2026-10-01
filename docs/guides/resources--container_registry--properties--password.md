---
page_title: "password"
subcategory: "Container"
description: "password for xcsh_container_registry."
xcsh_docs: {"aliases": [], "body_bytes": 1628, "body_sha256": "sha256:82524133d54d48814b65113582b8600602383f006714a5725667ee2bfe74d3f6", "canonical_id": "xcsh-docs:resources:container_registry:properties:password", "child_ids": ["xcsh-docs:resources:container_registry:properties:password:blindfold_secret_info", "xcsh-docs:resources:container_registry:properties:password:clear_secret_info"], "collection_id": "xcsh-docs:resources:container_registry:collection", "completeness": "complete", "id": "xcsh-docs:resources:container_registry:properties:password", "parent_id": "xcsh-docs:resources:container_registry:reference", "path": "docs/guides/resources--container_registry--properties--password.md", "provider_name": "container_registry", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/container_registry/properties/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "password for xcsh_container_registry.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["container_registryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# password

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md)
- [Property reference](resources--container_registry--reference.md)
- password

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--container_registry--properties--password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--container_registry--properties--password--clear_secret_info.md): complete subsection reference.

## Next pages

- [password.blindfold_secret_info](resources--container_registry--properties--password--blindfold_secret_info.md)
- [password.clear_secret_info](resources--container_registry--properties--password--clear_secret_info.md)
- [Property reference](resources--container_registry--reference.md)
- [xcsh_container_registry](../resources/container_registry.md)
