---
page_title: "access_info.vault_auth_info"
subcategory: ""
description: "access_info.vault_auth_info for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 1790, "body_sha256": "sha256:1ec09ee799089d018f6901f1c2429931d4ed74054428778fb1a2bcbceea49af7", "canonical_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token"], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info", "path": "docs/guides/resources--secret_management_access--properties--access_info--vault_auth_info.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "vault_auth_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/vault_auth_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.vault_auth_info for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.vault_auth_info

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md)
- [Property reference](resources--secret_management_access--reference.md)
- [access_info](resources--secret_management_access--properties--access_info.md)
- access_info.vault_auth_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Authentication parameters for Hashicorp Vault hosts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_role_auth",
    "token")}
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
  "x-ves-oneof-field-auth_params": "[\"app_role_auth\",\"token\"]"
}
```

Terraform syntax:

```terraform
vault_auth_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [app_role_auth](resources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth.md): complete subsection reference.

- [token](resources--secret_management_access--properties--access_info--vault_auth_info--token.md): complete subsection reference.

## Next pages

- [access_info.vault_auth_info.app_role_auth](resources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth.md)
- [access_info.vault_auth_info.token](resources--secret_management_access--properties--access_info--vault_auth_info--token.md)
- [access_info](resources--secret_management_access--properties--access_info.md)
- [xcsh_secret_management_access](../resources/secret_management_access.md)
