---
page_title: "password"
subcategory: "Container"
description: "password for xcsh_container_registry."
xcsh_docs: {"aliases": [], "body_bytes": 1347, "body_sha256": "sha256:55555152803e9a191bfc2e01dbbb4cff3313bba0a1e41fc56ac1c87bb0d08871", "canonical_id": "xcsh-docs:data-sources:container_registry:properties:password", "child_ids": ["xcsh-docs:data-sources:container_registry:properties:password:blindfold_secret_info", "xcsh-docs:data-sources:container_registry:properties:password:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:container_registry:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:container_registry:properties:password", "parent_id": "xcsh-docs:data-sources:container_registry:reference", "path": "docs/guides/data-sources--container_registry--properties--password.md", "provider_name": "container_registry", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/container_registry/properties/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "password for xcsh_container_registry.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["container_registryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# password

Breadcrumbs:

- [xcsh_container_registry](../data-sources/container_registry.md)
- [Property reference](data-sources--container_registry--reference.md)
- password

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

## Direct properties

- [blindfold_secret_info](data-sources--container_registry--properties--password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--container_registry--properties--password--clear_secret_info.md): complete subsection reference.

## Next pages

- [password.blindfold_secret_info](data-sources--container_registry--properties--password--blindfold_secret_info.md)
- [password.clear_secret_info](data-sources--container_registry--properties--password--clear_secret_info.md)
- [Property reference](data-sources--container_registry--reference.md)
- [xcsh_container_registry](../data-sources/container_registry.md)
