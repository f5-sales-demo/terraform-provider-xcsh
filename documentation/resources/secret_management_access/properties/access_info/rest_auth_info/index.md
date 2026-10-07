---
page_title: "access_info.rest_auth_info"
subcategory: ""
description: "Authentication parameters for REST based hosts."
xcsh_docs: {"aliases": ["access info rest auth info"], "body_bytes": 2026, "body_sha256": "sha256:d58d53b86ad21205fc6653659489b32f01d6bb779186c2b70cabaabb89d86d20", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:headers_auth", "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info", "path": "documentation/resources/secret_management_access/properties/access_info/rest_auth_info/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:basic_auth,headers_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:basic_auth,query_params_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:basic_auth,headers_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:headers_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:headers_auth,query_params_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:headers_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:basic_auth,query_params_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.rest_auth_info:ConflictingObjectAttributes:headers_auth,query_params_auth", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "rest_auth_info"], "schema_version": 1, "sections": [{"aliases": ["access info rest auth info basic auth"], "anchor": "section", "description": "AuthnTypeBasicAuth is used for using basic_auth mode of HTTP authentication.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "rest_auth_info", "basic_auth"], "syntax": "block", "type": "object"}, {"aliases": ["access info rest auth info headers auth"], "anchor": "section", "description": "AuthnTypeHeaders is used for setting headers for authentication.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:headers_auth", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "rest_auth_info", "headers_auth"], "syntax": "block", "type": "object"}, {"aliases": ["access info rest auth info query params auth"], "anchor": "section", "description": "AuthnTypeQueryParams is used for setting query_params for authentication.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "rest_auth_info", "query_params_auth"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/rest_auth_info/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Authentication parameters for REST based hosts.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
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
EnumExtractionComplete: false
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
