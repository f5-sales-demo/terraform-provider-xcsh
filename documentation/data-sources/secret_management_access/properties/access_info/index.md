---
page_title: "access_info"
subcategory: ""
description: "HostAccessInfoType contains the information about how to connect to the remote host."
xcsh_docs: {"aliases": ["access info"], "body_bytes": 3733, "body_sha256": "sha256:bc09c7b8cbac46a1293c73ff832485dda70cf442b4000f9c2386310c451bca7b", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info", "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config", "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info", "parent_id": "xcsh-docs:data-sources:secret_management_access:reference", "path": "documentation/data-sources/secret_management_access/properties/access_info/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133", "registry_path": "docs/guides/data-sources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info"], "schema_version": 1, "sections": [{"aliases": ["authentication", "credential setup", "credentials", "rest auth info"], "anchor": "section", "description": "Authentication parameters for REST based hosts.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "rest_auth_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["scheme"], "anchor": "schema-access_info--scheme", "description": "SchemeType is used to indicate URL scheme HTTP:// scheme HTTPS:// scheme.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "scheme"], "syntax": "attribute", "type": "string"}, {"aliases": ["server endpoint"], "anchor": "schema-access_info--server_endpoint", "description": "Endpoint to connect to, in host:port format.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "server_endpoint"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls config"], "anchor": "section", "description": "TLS configuration for upstream connections.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "tls_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication", "credential setup", "credentials", "vault auth info"], "anchor": "section", "description": "Authentication parameters for Hashicorp Vault hosts.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:vault_auth_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "vault_auth_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "HostAccessInfoType contains the information about how to connect to the remote host.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- access_info

<a id="section"></a>

Type: `"single"`. Computed.

HostAccessInfoType contains the information about how to connect to the remote host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_params": "[\"rest_auth_info\",\"vault_auth_info\"]"
}
```

## Direct properties

- [rest_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/): complete subsection reference.

<a id="schema-access_info--scheme"></a>

### scheme property

Type: `"string"`. Computed.

\[Enum: HTTP|HTTPS\] SchemeType is used to indicate URL scheme HTTP:// scheme HTTPS:// scheme.
Possible values are \`HTTP\`, \`HTTPS\`. Defaults to \`HTTP\`.

Upstream description:

SchemeType is used to indicate URL scheme

HTTP:// scheme HTTPS:// scheme.

Receipt-pinned upstream constraints:

```json
{
  "default": "HTTP",
  "enum": [
    "HTTP",
    "HTTPS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-access_info--server_endpoint"></a>

### server_endpoint property

Type: `"string"`. Computed.

Endpoint to connect to, in host:port format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/): complete subsection reference.

- [vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/): complete subsection reference.

## Next pages

- [access_info.rest_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/)
- [access_info.vault_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/vault_auth_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
