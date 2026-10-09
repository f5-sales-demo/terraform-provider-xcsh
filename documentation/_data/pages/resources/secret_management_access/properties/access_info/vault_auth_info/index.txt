---
page_title: "access_info.vault_auth_info"
subcategory: ""
description: "Authentication parameters for Hashicorp Vault hosts."
xcsh_docs: {"aliases": ["access info vault auth info"], "body_bytes": 1432, "body_sha256": "sha256:7a2b32be94c26cccc180e063c10c3adc0ed3f15bf4a36065a0b501c292ffa4c0", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info", "path": "documentation/resources/secret_management_access/properties/access_info/vault_auth_info/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2300210120332323-0231220231322210-3032002112232110-2333033211233200-1011201213321300-0120301210302002-0332332302303131-3032233113000312", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "vault_auth_info"], "schema_version": 1, "sections": [{"aliases": ["access info vault auth info app role auth"], "anchor": "section", "description": "AppRoleAuthInfoType contains parameters for AppRole authentication in Hashicorp Vault.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "vault_auth_info", "app_role_auth"], "syntax": "block", "type": "object"}, {"aliases": ["access info vault auth info token"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "vault_auth_info", "token"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/vault_auth_info/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Authentication parameters for Hashicorp Vault hosts.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
