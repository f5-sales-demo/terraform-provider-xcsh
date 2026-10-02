---
page_title: "access_info.rest_auth_info.query_params_auth"
subcategory: ""
description: "AuthnTypeQueryParams is used for setting query_params for authentication."
xcsh_docs: {"aliases": ["access info rest auth info query params auth", "authentication", "credential setup", "credentials"], "body_bytes": 1947, "body_sha256": "sha256:bc66182ef4817976b85c2f363d8ad8f81c2eb467f6768f581b3e50ecd67b49db", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth:query_params"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info", "path": "documentation/resources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1032023300203021-1102132320110003-2303022031202333-1031100030012103-2230031103203233-0222012031222211-1112333310230121-1030332213021130", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "rest_auth_info", "query_params_auth"], "schema_version": 1, "sections": [{"aliases": ["authentication", "credential setup", "credentials", "query params"], "anchor": "section", "description": "The set of authentication parameters to be passed as query parameters.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth:query_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "rest_auth_info", "query_params_auth", "query_params"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "AuthnTypeQueryParams is used for setting query_params for authentication.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.rest_auth_info.query_params_auth

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- [access_info.rest_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/)
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

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/query_params/): complete subsection reference.

## Next pages

- [access_info.rest_auth_info.query_params_auth.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/query_params/)
- [access_info.rest_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
