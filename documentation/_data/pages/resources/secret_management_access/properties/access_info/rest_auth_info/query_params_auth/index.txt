---
page_title: "access_info.rest_auth_info.query_params_auth"
subcategory: ""
description: "AuthnTypeQueryParams is used for setting query_params for authentication."
xcsh_docs: {"aliases": ["access info rest auth info query params auth"], "body_bytes": 1418, "body_sha256": "sha256:628cf175241ba3e7b01c3ebcff152cc536706166a6085dea92c4312a56dea7d0", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth:query_params"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info", "path": "documentation/resources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1032023300203021-1102132320110003-2303022031202333-1031100030012103-2230031103203233-0222012031222211-1112333310230121-1030332213021130", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "rest_auth_info", "query_params_auth"], "schema_version": 1, "sections": [{"aliases": ["access info rest auth info query params auth query params"], "anchor": "section", "description": "The set of authentication parameters to be passed as query parameters.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth:query_params", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "rest_auth_info", "query_params_auth", "query_params"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "AuthnTypeQueryParams is used for setting query_params for authentication.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
