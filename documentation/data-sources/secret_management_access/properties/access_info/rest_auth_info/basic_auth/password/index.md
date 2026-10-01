---
page_title: "access_info.rest_auth_info.basic_auth.password"
subcategory: ""
description: "access_info.rest_auth_info.basic_auth.password for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 2663, "body_sha256": "sha256:4ba81c75ff88cb3e1164ee5e5e05771e4c02bd05683d33ca794aa61f77e0d07f", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:basic_auth:password:blindfold_secret_info", "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:basic_auth:password:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:basic_auth:password", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "path": "documentation/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/index.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["access_info", "rest_auth_info", "basic_auth", "password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.rest_auth_info.basic_auth.password for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.rest_auth_info.basic_auth.password

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- [access_info.rest_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/)
- [access_info.rest_auth_info.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/)
- access_info.rest_auth_info.basic_auth.password

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/clear_secret_info/): complete subsection reference.

## Next pages

- [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/blindfold_secret_info/)
- [access_info.rest_auth_info.basic_auth.password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/clear_secret_info/)
- [access_info.rest_auth_info.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
