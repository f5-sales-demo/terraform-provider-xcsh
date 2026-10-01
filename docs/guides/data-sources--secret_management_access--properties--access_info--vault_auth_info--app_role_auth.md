---
page_title: "access_info.vault_auth_info.app_role_auth"
subcategory: ""
description: "access_info.vault_auth_info.app_role_auth for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 2084, "body_sha256": "sha256:2ba507af540aaf92fbb6ec6aa1e716ead3d4881a0124b2d05302c38596ac7b0e", "canonical_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id"], "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info", "path": "docs/guides/data-sources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "vault_auth_info", "app_role_auth"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.vault_auth_info.app_role_auth for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.vault_auth_info.app_role_auth

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
- [Property reference](data-sources--secret_management_access--reference.md)
- [access_info](data-sources--secret_management_access--properties--access_info.md)
- [access_info.vault_auth_info](data-sources--secret_management_access--properties--access_info--vault_auth_info.md)
- access_info.vault_auth_info.app_role_auth

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-access_info--vault_auth_info--app_role_auth--role_id"></a>

### role_id property

Type: `"string"`. Computed.

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

- [secret_id](data-sources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id.md): complete subsection reference.

## Next pages

- [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--properties--access_info--vault_auth_info--app_role_auth--secret_id.md)
- [access_info.vault_auth_info](data-sources--secret_management_access--properties--access_info--vault_auth_info.md)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
