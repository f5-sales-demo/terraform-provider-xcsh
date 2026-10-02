---
page_title: "access_info.rest_auth_info.basic_auth"
subcategory: ""
description: "AuthnTypeBasicAuth is used for using basic_auth mode of HTTP authentication."
xcsh_docs: {"aliases": ["access info rest auth info basic auth"], "body_bytes": 2570, "body_sha256": "sha256:496e4631598837fec44376a446aecd2e7a4a6c1615fdd85eaa8a9aa8f8162d56", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:basic_auth:password"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info", "path": "documentation/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2332210232020211-2222222113310221-0101203303121203-2133122120101003-2122202120203322-3300313221010130-2220113232032322-3003233230300301", "registry_path": "docs/guides/data-sources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "rest_auth_info", "basic_auth"], "schema_version": 1, "sections": [{"aliases": ["password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:basic_auth:password", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_info", "rest_auth_info", "basic_auth", "password"], "syntax": "attribute", "type": "object"}, {"aliases": ["username"], "anchor": "schema-access_info--rest_auth_info--basic_auth--username", "description": "The username to encode in Basic Auth scheme.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "rest_auth_info", "basic_auth", "username"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "AuthnTypeBasicAuth is used for using basic_auth mode of HTTP authentication.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.rest_auth_info.basic_auth

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- [access_info.rest_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/)
- access_info.rest_auth_info.basic_auth

<a id="section"></a>

Type: `"single"`. Computed.

AuthnTypeBasicAuth is used for using basic\_auth mode of HTTP authentication.

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

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/): complete subsection reference.

<a id="schema-access_info--rest_auth_info--basic_auth--username"></a>

### username property

Type: `"string"`. Computed.

The username to encode in Basic Auth scheme.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "identity",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [access_info.rest_auth_info.basic_auth.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/basic_auth/password/)
- [access_info.rest_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/rest_auth_info/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
