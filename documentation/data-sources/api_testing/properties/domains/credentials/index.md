---
page_title: "domains.credentials"
subcategory: ""
description: "Add credentials for API testing to use in the selected environment."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials", "domains credentials"], "body_bytes": 4094, "body_sha256": "sha256:e133665019f8095ad30b07d8ba21fb25a786816431158811e1ff9db1f215cc40", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_testing:properties:domains:credentials:admin", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:api_key", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:login_endpoint", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:standard"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials", "parent_id": "xcsh-docs:data-sources:api_testing:properties:domains", "path": "documentation/data-sources/api_testing/properties/domains/credentials/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112", "registry_path": "docs/guides/data-sources--api_testing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["domains", "credentials"], "schema_version": 1, "sections": [{"aliases": ["admin"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:admin", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "credentials", "admin"], "syntax": "attribute", "type": "object"}, {"aliases": ["api key"], "anchor": "section", "description": "API Key", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:api_key", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "credentials", "api_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["basic auth"], "anchor": "section", "description": "Basic Authentication.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "credentials", "basic_auth"], "syntax": "attribute", "type": "object"}, {"aliases": ["bearer token"], "anchor": "section", "description": "Configuration parameter for bearer token.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "credentials", "bearer_token"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication", "credential name", "credential setup", "credentials"], "anchor": "schema-domains--credentials--credential_name", "description": "Enter a unique name for the credentials used in API testing.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "credentials", "credential_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["login", "login endpoint", "login result", "sign in"], "anchor": "section", "description": "Login Endpoint.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:login_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "credentials", "login_endpoint"], "syntax": "attribute", "type": "object"}, {"aliases": ["standard"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:standard", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "credentials", "standard"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/properties/domains/credentials/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Add credentials for API testing to use in the selected environment.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_testingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/)
- domains.credentials

<a id="section"></a>

Type: `"list"`. Computed.

Add credentials for API testing to use in the selected environment.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Direct properties

- [admin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/admin/): complete subsection reference.

- [api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/api_key/): complete subsection reference.

- [basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/basic_auth/): complete subsection reference.

- [bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/bearer_token/): complete subsection reference.

<a id="schema-domains--credentials--credential_name"></a>

### credential_name property

Type: `"string"`. Computed.

Enter a unique name for the credentials used in API testing.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [login_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/login_endpoint/): complete subsection reference.

- [standard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/standard/): complete subsection reference.

## Next pages

- [domains.credentials.admin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/admin/)
- [domains.credentials.api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/api_key/)
- [domains.credentials.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/basic_auth/)
- [domains.credentials.bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/bearer_token/)
- [domains.credentials.login_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/login_endpoint/)
- [domains.credentials.standard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/standard/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/)
- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/)
