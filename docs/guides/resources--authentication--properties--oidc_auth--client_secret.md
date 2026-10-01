---
page_title: "oidc_auth.client_secret"
subcategory: ""
description: "oidc_auth.client_secret for xcsh_authentication."
xcsh_docs: {"aliases": [], "body_bytes": 1787, "body_sha256": "sha256:6e17cdedf5417a20d03835835c7ec0a7a491ca836c299ccc2a93105f273c418b", "canonical_id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret", "child_ids": ["xcsh-docs:resources:authentication:properties:oidc_auth:client_secret:blindfold_secret_info", "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret:clear_secret_info"], "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:properties:oidc_auth:client_secret", "parent_id": "xcsh-docs:resources:authentication:properties:oidc_auth", "path": "docs/guides/resources--authentication--properties--oidc_auth--client_secret.md", "provider_name": "authentication", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["oidc_auth", "client_secret"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/properties/oidc_auth/client_secret/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "oidc_auth.client_secret for xcsh_authentication.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# oidc_auth.client_secret

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md)
- [Property reference](resources--authentication--reference.md)
- [oidc_auth](resources--authentication--properties--oidc_auth.md)
- oidc_auth.client_secret

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
client_secret {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--authentication--properties--oidc_auth--client_secret--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--authentication--properties--oidc_auth--client_secret--clear_secret_info.md): complete subsection reference.

## Next pages

- [oidc_auth.client_secret.blindfold_secret_info](resources--authentication--properties--oidc_auth--client_secret--blindfold_secret_info.md)
- [oidc_auth.client_secret.clear_secret_info](resources--authentication--properties--oidc_auth--client_secret--clear_secret_info.md)
- [oidc_auth](resources--authentication--properties--oidc_auth.md)
- [xcsh_authentication](../resources/authentication.md)
