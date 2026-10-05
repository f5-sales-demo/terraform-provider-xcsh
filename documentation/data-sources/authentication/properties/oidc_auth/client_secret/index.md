---
page_title: "oidc_auth.client_secret"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["oidc auth client secret"], "body_bytes": 1957, "body_sha256": "sha256:ad7141b1541923ffe12ddb450a4e0ad86c1051247b34ce04a6c5cc047d848d56", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:authentication:properties:oidc_auth:client_secret:blindfold_secret_info", "xcsh-docs:data-sources:authentication:properties:oidc_auth:client_secret:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:client_secret", "parent_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth", "path": "documentation/data-sources/authentication/properties/oidc_auth/client_secret/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3230011212002033-3231130123201030-0221202213110032-1110301332101132-3330101221132121-1033323000113201-1033122020221130-2030100331302021", "registry_path": "docs/guides/data-sources--authentication--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["oidc_auth", "client_secret"], "schema_version": 1, "sections": [{"aliases": ["oidc auth client secret blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:client_secret:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["oidc_auth", "client_secret", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["oidc auth client secret clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:authentication:properties:oidc_auth:client_secret:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["oidc_auth", "client_secret", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authentication/properties/oidc_auth/client_secret/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# oidc_auth.client_secret

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/)
- [oidc_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/)
- oidc_auth.client_secret

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/client_secret/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/client_secret/clear_secret_info/): complete subsection reference.

## Next pages

- [oidc_auth.client_secret.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/client_secret/blindfold_secret_info/)
- [oidc_auth.client_secret.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/client_secret/clear_secret_info/)
- [oidc_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/oidc_auth/)
- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/)
