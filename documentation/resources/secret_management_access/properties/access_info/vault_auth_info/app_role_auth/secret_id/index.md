---
page_title: "access_info.vault_auth_info.app_role_auth.secret_id"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["access info vault auth info app role auth secret id"], "body_bytes": 2994, "body_sha256": "sha256:2ada28ac98d3bae1561d04372329ff2a61fc3febf970908ce80fb9d70735a67d", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id:blindfold_secret_info", "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "path": "documentation/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2203120011331101-1223102021222200-1333211322323230-0010033101100020-3021211120222202-0131123113232120-0100231113111333-1100312311212212", "registry_path": "docs/guides/resources--secret_management_access--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "access_info.vault_auth_info.app_role_auth.secret_id:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.vault_auth_info.app_role_auth.secret_id:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "vault_auth_info", "app_role_auth", "secret_id"], "schema_version": 1, "sections": [{"aliases": ["access info vault auth info app role auth secret id blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-access_info--vault_auth_info--app_role_auth--secret_id--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id:blindfold_secret_info", "type": "requires"}], "schema_path": ["access_info", "vault_auth_info", "app_role_auth", "secret_id", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["access info vault auth info app role auth secret id clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-access_info--vault_auth_info--app_role_auth--secret_id--clear_secret_info--url", "enforcement": "provider-schema", "group": "access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id:clear_secret_info", "type": "requires"}], "schema_path": ["access_info", "vault_auth_info", "app_role_auth", "secret_id", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.vault_auth_info.app_role_auth.secret_id

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- [access_info.vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/)
- [access_info.vault_auth_info.app_role_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/)
- access_info.vault_auth_info.app_role_auth.secret_id

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_id {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/clear_secret_info/): complete subsection reference.

## Next pages

- [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/blindfold_secret_info/)
- [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/clear_secret_info/)
- [access_info.vault_auth_info.app_role_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
