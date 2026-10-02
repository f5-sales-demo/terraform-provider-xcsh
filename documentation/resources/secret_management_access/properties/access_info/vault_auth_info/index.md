---
page_title: "access_info.vault_auth_info"
subcategory: ""
description: "Authentication parameters for Hashicorp Vault hosts."
xcsh_docs: {"aliases": ["access info vault auth info"], "body_bytes": 2243, "body_sha256": "sha256:d10932c3c9d79bff0a4d7a567898d87f710452324fea56b90ed72830efad1f8f", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info", "path": "documentation/resources/secret_management_access/properties/access_info/vault_auth_info/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2300210120332323-0231220231322210-3032002112232110-2333033211233200-1011201213321300-0120301210302002-0332332302303131-3032233113000312", "registry_path": "docs/guides/resources--secret_management_access--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "access_info.vault_auth_info:ConflictingObjectAttributes:app_role_auth,token", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.vault_auth_info:ConflictingObjectAttributes:app_role_auth,token", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "vault_auth_info"], "schema_version": 1, "sections": [{"aliases": ["app role auth"], "anchor": "section", "description": "AppRoleAuthInfoType contains parameters for AppRole authentication in Hashicorp Vault.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "vault_auth_info", "app_role_auth"], "syntax": "block", "type": "object"}, {"aliases": ["token"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "access_info.vault_auth_info.token:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.vault_auth_info.token:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token:clear_secret_info", "type": "conflicts"}], "schema_path": ["access_info", "vault_auth_info", "token"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/vault_auth_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Authentication parameters for Hashicorp Vault hosts.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.vault_auth_info

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- access_info.vault_auth_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Authentication parameters for Hashicorp Vault hosts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_role_auth",
    "token")}
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
  "x-ves-oneof-field-auth_params": "[\"app_role_auth\",\"token\"]"
}
```

Terraform syntax:

```terraform
vault_auth_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [app_role_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/): complete subsection reference.

- [token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/token/): complete subsection reference.

## Next pages

- [access_info.vault_auth_info.app_role_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/)
- [access_info.vault_auth_info.token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/token/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
