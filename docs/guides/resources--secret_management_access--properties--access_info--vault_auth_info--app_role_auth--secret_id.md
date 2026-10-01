---
page_title: "access_info.vault_auth_info.app_role_auth.secret_id"
subcategory: ""
description: "access_info.vault_auth_info.app_role_auth.secret_id for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 2452, "body_sha256": "sha256:900393bc3198d5d3ca2ef8047fd4bc3c0edd90181935508f089e464ccba48c0a", "canonical_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id:blindfold_secret_info", "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id:clear_secret_info"], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "path": "docs/guides/resources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "vault_auth_info", "app_role_auth", "secret_id"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.vault_auth_info.app_role_auth.secret_id for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.vault_auth_info.app_role_auth.secret_id

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md)
- [Property reference](resources--secret_management_access--reference.md)
- [access_info](resources--secret_management_access--properties--access_info.md)
- [access_info.vault_auth_info](resources--secret_management_access--properties--access_info--vault_auth_info.md)
- [access_info.vault_auth_info.app_role_auth](resources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth.md)
- access_info.vault_auth_info.app_role_auth.secret_id

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
secret_id {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id--clear_secret_info.md): complete subsection reference.

## Next pages

- [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info](resources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id--blindfold_secret_info.md)
- [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info](resources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id--clear_secret_info.md)
- [access_info.vault_auth_info.app_role_auth](resources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth.md)
- [xcsh_secret_management_access](../resources/secret_management_access.md)
