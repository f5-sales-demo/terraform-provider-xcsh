---
page_title: "access_info.rest_auth_info"
subcategory: ""
description: "access_info.rest_auth_info for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 1765, "body_sha256": "sha256:df8b7e6015109155483c3bdfbf064ab9472bafaac0a5a529afe2ad53c34b47b1", "canonical_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:headers_auth", "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth"], "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info", "path": "docs/guides/data-sources--secret_management_access--properties--access_info--rest_auth_info.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "rest_auth_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/rest_auth_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.rest_auth_info for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# access_info.rest_auth_info

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
- [Property reference](data-sources--secret_management_access--reference.md)
- [access_info](data-sources--secret_management_access--properties--access_info.md)
- access_info.rest_auth_info

<a id="section"></a>

Type: `"single"`. Computed.

Authentication parameters for REST based hosts.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_params": "[\"basic_auth\",\"headers_auth\",\"query_params_auth\"]"
}
```

## Direct properties

- [basic_auth](data-sources--secret_management_access--properties--access_info--rest_auth_info--basic_auth.md): complete subsection reference.

- [headers_auth](data-sources--secret_management_access--properties--access_info--rest_auth_info--headers_auth.md): complete subsection reference.

- [query_params_auth](data-sources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth.md): complete subsection reference.

## Next pages

- [access_info.rest_auth_info.basic_auth](data-sources--secret_management_access--properties--access_info--rest_auth_info--basic_auth.md)
- [access_info.rest_auth_info.headers_auth](data-sources--secret_management_access--properties--access_info--rest_auth_info--headers_auth.md)
- [access_info.rest_auth_info.query_params_auth](data-sources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth.md)
- [access_info](data-sources--secret_management_access--properties--access_info.md)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
