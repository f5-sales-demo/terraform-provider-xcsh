---
page_title: "access_info.vault_auth_info.token"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["access info vault auth info token"], "body_bytes": 2354, "body_sha256": "sha256:82ddf3a9f8fda50a980b367753e7acfeb2959869c44f77348feaed090f17af63", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:token:blindfold_secret_info", "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:token:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:token", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info", "path": "documentation/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2200222113222202-3210312221321013-3230302101001032-0202022233223101-0311212122310232-3200300221230331-2113223010002011-2202110112012203", "registry_path": "docs/guides/data-sources--secret_management_access--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "vault_auth_info", "token"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:token:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "vault_auth_info", "token", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:token:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "vault_auth_info", "token", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.vault_auth_info.token

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- [access_info.vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/)
- access_info.vault_auth_info.token

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/clear_secret_info/): complete subsection reference.

## Next pages

- [access_info.vault_auth_info.token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/blindfold_secret_info/)
- [access_info.vault_auth_info.token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/token/clear_secret_info/)
- [access_info.vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
