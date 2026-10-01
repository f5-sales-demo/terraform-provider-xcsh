---
page_title: "access_info.rest_auth_info.query_params_auth"
subcategory: ""
description: "access_info.rest_auth_info.query_params_auth for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 1545, "body_sha256": "sha256:b467940d5c7836b4a309bd667f70a047c1c5eeffd807565ba8347189974745e4", "canonical_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth:query_params"], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info", "path": "docs/guides/resources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "rest_auth_info", "query_params_auth"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.rest_auth_info.query_params_auth for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.rest_auth_info.query_params_auth

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md)
- [Property reference](resources--secret_management_access--reference.md)
- [access_info](resources--secret_management_access--properties--access_info.md)
- [access_info.rest_auth_info](resources--secret_management_access--properties--access_info--rest_auth_info.md)
- access_info.rest_auth_info.query_params_auth

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

AuthnTypeQueryParams is used for setting query\_params for authentication.

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
query_params_auth {
  # Configure direct properties listed below.
}
```

## Direct properties

- [query_params](resources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth--query_params.md): complete subsection reference.

## Next pages

- [access_info.rest_auth_info.query_params_auth.query_params](resources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth--query_params.md)
- [access_info.rest_auth_info](resources--secret_management_access--properties--access_info--rest_auth_info.md)
- [xcsh_secret_management_access](../resources/secret_management_access.md)
