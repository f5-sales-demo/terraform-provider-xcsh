---
page_title: "access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info"
subcategory: ""
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["access info vault auth info app role auth secret id clear secret info"], "body_bytes": 4014, "body_sha256": "sha256:e8c3b00fdcfd30fefb2271ad28c761c93f9f59e87eea6688306d49945170ac4e", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id:clear_secret_info", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id", "path": "documentation/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3210131331231023-2130202300320212-3120220201021330-1100313132122113-3000202123200023-3001310010123021-0110223232203320-0003110101320221", "registry_path": "docs/guides/data-sources--secret_management_access--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "vault_auth_info", "app_role_auth", "secret_id", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["provider ref"], "anchor": "schema-access_info--vault_auth_info--app_role_auth--secret_id--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "vault_auth_info", "app_role_auth", "secret_id", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["url"], "anchor": "schema-access_info--vault_auth_info--app_role_auth--secret_id--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info:app_role_auth:secret_id:clear_secret_info", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "vault_auth_info", "app_role_auth", "secret_id", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/clear_secret_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- [access_info.vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/)
- [access_info.vault_auth_info.app_role_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/)
- [access_info.vault_auth_info.app_role_auth.secret_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/)
- access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info

<a id="section"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="schema-access_info--vault_auth_info--app_role_auth--secret_id--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-access_info--vault_auth_info--app_role_auth--secret_id--clear_secret_info--url"></a>

### url property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## Next pages

- [access_info.vault_auth_info.app_role_auth.secret_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/app_role_auth/secret_id/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
