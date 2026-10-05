---
page_title: "access_info.vault_auth_info.app_role_auth"
subcategory: ""
description: "AppRoleAuthInfoType contains parameters for AppRole authentication in Hashicorp Vault."
xcsh_docs: {"aliases": ["access info vault auth info app role auth"], "body_bytes": 2486, "body_sha256": "sha256:33f143e35aba585353afba5f7363f926a14c7f76b4ab64af7020e6c1ebb2e87c", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info", "path": "documentation/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0012301320210212-3110121123113000-2122321022210103-1322002120013000-2311021002123101-3121123121201133-3032133200103013-2133120000130022", "registry_path": "docs/guides/data-sources--secret_management_access--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "vault_auth_info", "app_role_auth"], "schema_version": 1, "sections": [{"aliases": ["access info vault auth info app role auth role id"], "anchor": "schema-access_info--vault_auth_info--app_role_auth--role_id", "description": "Role-ID to be used for authentication.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "vault_auth_info", "app_role_auth", "role_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info vault auth info app role auth secret id"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "vault_auth_info", "app_role_auth", "secret_id"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "AppRoleAuthInfoType contains parameters for AppRole authentication in Hashicorp Vault.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.vault_auth_info.app_role_auth

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- [access_info.vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/)
- access_info.vault_auth_info.app_role_auth

<a id="section"></a>

Type: `"single"`. Computed.

AppRoleAuthInfoType contains parameters for AppRole authentication in Hashicorp Vault.

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

<a id="schema-access_info--vault_auth_info--app_role_auth--role_id"></a>

### role_id property

Type: `"string"`. Computed.

Role ID. Role-ID to be used for authentication.

Upstream description:

Role-ID to be used for authentication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [secret_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/): complete subsection reference.

## Next pages

- [access_info.vault_auth_info.app_role_auth.secret_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/)
- [access_info.vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
