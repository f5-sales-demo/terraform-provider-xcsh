---
page_title: "access_info.vault_auth_info.app_role_auth"
subcategory: ""
description: "access_info.vault_auth_info.app_role_auth for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 2182, "body_sha256": "sha256:04420b6a0d68fbf29eaf0e7cbfbc447ab9dbe0caa838bd402fcdd2758aea57f5", "canonical_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id"], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info", "path": "docs/guides/resources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "vault_auth_info", "app_role_auth"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.vault_auth_info.app_role_auth for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.vault_auth_info.app_role_auth

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md)
- [Property reference](resources--secret_management_access--reference.md)
- [access_info](resources--secret_management_access--properties--access_info.md)
- [access_info.vault_auth_info](resources--secret_management_access--properties--access_info--vault_auth_info.md)
- access_info.vault_auth_info.app_role_auth

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

AppRoleAuthInfoType contains parameters for AppRole authentication in Hashicorp Vault.

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
app_role_auth {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-access_info--vault_auth_info--app_role_auth--role_id"></a>

### role_id property

Type: `"string"`. Optional.

Role ID. Role-ID to be used for authentication.

Upstream description:

Role-ID to be used for authentication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [secret_id](resources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id.md): complete subsection reference.

## Next pages

- [access_info.vault_auth_info.app_role_auth.secret_id](resources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id.md)
- [access_info.vault_auth_info](resources--secret_management_access--properties--access_info--vault_auth_info.md)
- [xcsh_secret_management_access](../resources/secret_management_access.md)
