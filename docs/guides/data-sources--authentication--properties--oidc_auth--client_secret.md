---
page_title: "oidc_auth.client_secret"
subcategory: ""
description: "oidc_auth.client_secret for xcsh_authentication."
xcsh_docs: {"aliases": [], "body_bytes": 1405, "body_sha256": "sha256:fe7004ef67d15515ff4a83dbb69da21044c9fa55010b7af3c1055e9419a3eaa4", "canonical_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:client_secret", "child_ids": ["xcsh-docs:data-sources:authentication:properties:oidc_auth:client_secret:blindfold_secret_info", "xcsh-docs:data-sources:authentication:properties:oidc_auth:client_secret:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:client_secret", "parent_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth", "path": "docs/guides/data-sources--authentication--properties--oidc_auth--client_secret.md", "provider_name": "authentication", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["oidc_auth", "client_secret"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authentication/properties/oidc_auth/client_secret/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "oidc_auth.client_secret for xcsh_authentication.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# oidc_auth.client_secret

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md)
- [Property reference](data-sources--authentication--reference.md)
- [oidc_auth](data-sources--authentication--properties--oidc_auth.md)
- oidc_auth.client_secret

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

- [blindfold_secret_info](data-sources--authentication--properties--oidc_auth--client_secret--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--authentication--properties--oidc_auth--client_secret--clear_secret_info.md): complete subsection reference.

## Next pages

- [oidc_auth.client_secret.blindfold_secret_info](data-sources--authentication--properties--oidc_auth--client_secret--blindfold_secret_info.md)
- [oidc_auth.client_secret.clear_secret_info](data-sources--authentication--properties--oidc_auth--client_secret--clear_secret_info.md)
- [oidc_auth](data-sources--authentication--properties--oidc_auth.md)
- [xcsh_authentication](../data-sources/authentication.md)
