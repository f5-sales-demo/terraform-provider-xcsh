---
page_title: "access_info.vault_auth_info"
subcategory: ""
description: "Authentication parameters for Hashicorp Vault hosts."
xcsh_docs: {"aliases": ["access info vault auth info"], "body_bytes": 1978, "body_sha256": "sha256:16452847e39eaaa15e784f17fbeb460586855c64bbc7fdbf45460eec9cedeea8", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:token"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info", "path": "documentation/data-sources/secret_management_access/properties/access_info/vault_auth_info/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022", "registry_path": "docs/guides/data-sources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "vault_auth_info"], "schema_version": 1, "sections": [{"aliases": ["access info vault auth info app role auth"], "anchor": "section", "description": "AppRoleAuthInfoType contains parameters for AppRole authentication in Hashicorp Vault.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "vault_auth_info", "app_role_auth"], "syntax": "attribute", "type": "object"}, {"aliases": ["access info vault auth info token"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:token", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "vault_auth_info", "token"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/vault_auth_info/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Authentication parameters for Hashicorp Vault hosts.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.vault_auth_info

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- access_info.vault_auth_info

<a id="section"></a>

Type: `"single"`. Computed.

Authentication parameters for Hashicorp Vault hosts.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_params": "[\"app_role_auth\",\"token\"]"
}
```

## Direct properties

- [app_role_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/): complete subsection reference.

- [token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/): complete subsection reference.

## Next pages

- [access_info.vault_auth_info.app_role_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/)
- [access_info.vault_auth_info.token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
