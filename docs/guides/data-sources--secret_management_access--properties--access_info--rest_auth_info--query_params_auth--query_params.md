---
page_title: "access_info.rest_auth_info.query_params_auth.query_params"
subcategory: ""
description: "access_info.rest_auth_info.query_params_auth.query_params for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 1766, "body_sha256": "sha256:116ba87a8094af6460b1468bd2286119dd68874cab50407d9a59392f9ca2e2db", "canonical_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth:query_params", "child_ids": [], "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth:query_params", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "path": "docs/guides/data-sources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth--query_params.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "rest_auth_info", "query_params_auth", "query_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/query_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.rest_auth_info.query_params_auth.query_params for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.rest_auth_info.query_params_auth.query_params

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
- [Property reference](data-sources--secret_management_access--reference.md)
- [access_info](data-sources--secret_management_access--properties--access_info.md)
- [access_info.rest_auth_info](data-sources--secret_management_access--properties--access_info--rest_auth_info.md)
- [access_info.rest_auth_info.query_params_auth](data-sources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth.md)
- access_info.rest_auth_info.query_params_auth.query_params

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [access_info.rest_auth_info.query_params_auth](data-sources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth.md)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
