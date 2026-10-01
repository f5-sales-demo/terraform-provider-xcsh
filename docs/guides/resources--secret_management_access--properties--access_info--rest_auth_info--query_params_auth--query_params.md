---
page_title: "access_info.rest_auth_info.query_params_auth.query_params"
subcategory: ""
description: "access_info.rest_auth_info.query_params_auth.query_params for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 1819, "body_sha256": "sha256:a46c31e73c6b0eb1c6bd8cdb7597c84565920a7f8f6a74f36e49a564b1d7ed6c", "canonical_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth:query_params", "child_ids": [], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth:query_params", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "path": "docs/guides/resources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth--query_params.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "rest_auth_info", "query_params_auth", "query_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/query_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.rest_auth_info.query_params_auth.query_params for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.rest_auth_info.query_params_auth.query_params

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md)
- [Property reference](resources--secret_management_access--reference.md)
- [access_info](resources--secret_management_access--properties--access_info.md)
- [access_info.rest_auth_info](resources--secret_management_access--properties--access_info--rest_auth_info.md)
- [access_info.rest_auth_info.query_params_auth](resources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth.md)
- access_info.rest_auth_info.query_params_auth.query_params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

The set of authentication parameters to be passed as query parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

Terraform syntax:

```terraform
query_params {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [access_info.rest_auth_info.query_params_auth](resources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth.md)
- [xcsh_secret_management_access](../resources/secret_management_access.md)
