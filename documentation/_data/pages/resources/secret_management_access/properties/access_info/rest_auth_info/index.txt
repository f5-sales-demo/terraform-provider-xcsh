---
page_title: "access_info.rest_auth_info"
subcategory: ""
description: "Authentication parameters for REST based hosts."
xcsh_docs: {"aliases": ["access info rest auth info"], "body_bytes": 2841, "body_sha256": "sha256:378293c4a86c1e349722224ea4250f4a5d1156a765f8b425af8ea5da02245b7f", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:headers_auth", "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info", "path": "documentation/resources/secret_management_access/properties/access_info/rest_auth_info/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:basic_auth,headers_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:basic_auth,query_params_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:basic_auth,headers_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:headers_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:headers_auth,query_params_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:headers_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:basic_auth,query_params_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:headers_auth,query_params_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "rest_auth_info"], "schema_version": 1, "sections": [{"aliases": ["access info rest auth info basic auth"], "anchor": "section", "description": "AuthnTypeBasicAuth is used for using basic_auth mode of HTTP authentication.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "rest_auth_info", "basic_auth"], "syntax": "block", "type": "object"}, {"aliases": ["access info rest auth info headers auth"], "anchor": "section", "description": "AuthnTypeHeaders is used for setting headers for authentication.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:headers_auth", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "rest_auth_info", "headers_auth"], "syntax": "block", "type": "object"}, {"aliases": ["access info rest auth info query params auth"], "anchor": "section", "description": "AuthnTypeQueryParams is used for setting query_params for authentication.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "rest_auth_info", "query_params_auth"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/rest_auth_info/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Authentication parameters for REST based hosts.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.rest_auth_info

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- access_info.rest_auth_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Authentication parameters for REST based hosts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("basic_auth",
    "headers_auth"),
  validators.ConflictingObjectAttributes("basic_auth",
    "query_params_auth"),
  validators.ConflictingObjectAttributes("headers_auth",
    "query_params_auth")}
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
  "x-ves-oneof-field-auth_params": "[\"basic_auth\",\"headers_auth\",\"query_params_auth\"]"
}
```

Terraform syntax:

```terraform
rest_auth_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/): complete subsection reference.

- [headers_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/headers_auth/): complete subsection reference.

- [query_params_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/): complete subsection reference.

## Next pages

- [access_info.rest_auth_info.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/)
- [access_info.rest_auth_info.headers_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/headers_auth/)
- [access_info.rest_auth_info.query_params_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
