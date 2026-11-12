---
page_title: "access_info.vault_auth_info.app_role_auth"
subcategory: ""
description: "AppRoleAuthInfoType contains parameters for AppRole authentication in Hashicorp Vault."
xcsh_docs: {"aliases": ["access info vault auth info app role auth"], "body_bytes": 2002, "body_sha256": "sha256:8fa0add62612f5ba1d793e529b8b1202a7ebb72009c140cd67b0298e54952b9c", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info", "path": "documentation/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-0030032121233032-3133222130022033-3020230323331202-3211220030111033-0100123113222122-2331103233022011-2000121002233003-1333030132032231", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "vault_auth_info", "app_role_auth"], "schema_version": 1, "sections": [{"aliases": ["access info vault auth info app role auth role id"], "anchor": "schema-access_info--vault_auth_info--app_role_auth--role_id", "description": "Role-ID to be used for authentication.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "vault_auth_info", "app_role_auth", "role_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["access info vault auth info app role auth secret id"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "vault_auth_info", "app_role_auth", "secret_id"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "AppRoleAuthInfoType contains parameters for AppRole authentication in Hashicorp Vault.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.vault_auth_info.app_role_auth

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- [access_info.vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/)
- access_info.vault_auth_info.app_role_auth

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
app_role_auth {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-access_info--vault_auth_info--app_role_auth--role_id"></a>

### role_id property

Type: `"string"`. Optional.

Role ID. Role-ID to be used for authentication.

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
      "validatedAt": "2026-10-29T12:14:42+00:00"
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

- [secret_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/): complete subsection reference.
