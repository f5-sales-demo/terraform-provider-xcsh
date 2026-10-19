---
page_title: "access_info.rest_auth_info.query_params_auth"
subcategory: ""
description: "AuthnTypeQueryParams is used for setting query_params for authentication."
xcsh_docs: {"aliases": ["access info rest auth info query params auth"], "body_bytes": 1307, "body_sha256": "sha256:4ce1b0c01ce5a0106f6ecfa70664eaafba3544e988b1d0b6c3bf846eebcaba32", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth:query_params"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info", "path": "documentation/data-sources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3123211020031221-1033310323003001-3333320020332010-2311112310201202-2122332031233311-1221220333200110-1203202212131100-1123210312001010", "registry_path": "docs/guides/data-sources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "rest_auth_info", "query_params_auth"], "schema_version": 1, "sections": [{"aliases": ["access info rest auth info query params auth query params"], "anchor": "section", "description": "The set of authentication parameters to be passed as query parameters.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth:query_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "rest_auth_info", "query_params_auth", "query_params"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "AuthnTypeQueryParams is used for setting query_params for authentication.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.rest_auth_info.query_params_auth

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- [access_info.rest_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/)
- access_info.rest_auth_info.query_params_auth

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/query_params_auth/query_params/): complete subsection reference.
