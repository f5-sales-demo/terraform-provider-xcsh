---
page_title: "access_info.vault_auth_info.token"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["access info vault auth info token"], "body_bytes": 2626, "body_sha256": "sha256:1b18dd7c8150bf5e242319ff9bff435c20b39c919ee759772b48f67bddfd8cdb", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token:blindfold_secret_info", "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info", "path": "documentation/resources/secret_management_access/properties/access_info/vault_auth_info/token/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3231030210203130-0323032211032312-1020202213120211-2121013313233310-1031103301332032-2331132032211122-0023130320122100-1020130002320101", "registry_path": "docs/guides/resources--secret_management_access--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "access_info.vault_auth_info.token:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.vault_auth_info.token:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "vault_auth_info", "token"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-access_info--vault_auth_info--token--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "access_info.vault_auth_info.token.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token:blindfold_secret_info", "type": "requires"}], "schema_path": ["access_info", "vault_auth_info", "token", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-access_info--vault_auth_info--token--clear_secret_info--url", "enforcement": "provider-schema", "group": "access_info.vault_auth_info.token.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:vault_auth_info:token:clear_secret_info", "type": "requires"}], "schema_path": ["access_info", "vault_auth_info", "token", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/vault_auth_info/token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.vault_auth_info.token

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- [access_info.vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/)
- access_info.vault_auth_info.token

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
token {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/token/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/token/clear_secret_info/): complete subsection reference.

## Next pages

- [access_info.vault_auth_info.token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/token/blindfold_secret_info/)
- [access_info.vault_auth_info.token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/token/clear_secret_info/)
- [access_info.vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/vault_auth_info/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
