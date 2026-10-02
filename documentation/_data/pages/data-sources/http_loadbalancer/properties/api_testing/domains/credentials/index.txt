---
page_title: "api_testing.domains.credentials"
subcategory: "Load Balancing"
description: "Add credentials for API testing to use in the selected environment."
xcsh_docs: {"aliases": ["api testing domains credentials", "authentication", "credential setup", "credentials"], "body_bytes": 4637, "body_sha256": "sha256:009c8e4daf3b3d623788066dd7cf52329a165aa8a2b19bf0bd2fb70728806d54", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:admin", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:api_key", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:standard"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains", "path": "documentation/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-1210021130100220-3031021333331212-0312211101122113-3113123120210012-0013320121231203-1320013203200103-0110020133112032-0303113020332201", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_testing", "domains", "credentials"], "schema_version": 1, "sections": [{"aliases": ["admin"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:admin", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "domains", "credentials", "admin"], "syntax": "attribute", "type": "object"}, {"aliases": ["api key"], "anchor": "section", "description": "API Key", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:api_key", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_testing", "domains", "credentials", "api_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication", "basic auth", "credential setup", "credentials"], "anchor": "section", "description": "Basic Authentication.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_testing", "domains", "credentials", "basic_auth"], "syntax": "attribute", "type": "object"}, {"aliases": ["bearer token"], "anchor": "section", "description": "Configuration parameter for bearer token.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_testing", "domains", "credentials", "bearer_token"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication", "credential name", "credential setup", "credentials"], "anchor": "schema-api_testing--domains--credentials--credential_name", "description": "Enter a unique name for the credentials used in API testing.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "domains", "credentials", "credential_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["login", "login endpoint", "login result", "sign in"], "anchor": "section", "description": "Login Endpoint.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_testing", "domains", "credentials", "login_endpoint"], "syntax": "attribute", "type": "object"}, {"aliases": ["standard"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:standard", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "domains", "credentials", "standard"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Add credentials for API testing to use in the selected environment.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_testing.domains.credentials

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/)
- [api_testing.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/)
- api_testing.domains.credentials

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

- [admin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/admin/): complete subsection reference.

- [api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/api_key/): complete subsection reference.

- [basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/basic_auth/): complete subsection reference.

- [bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/bearer_token/): complete subsection reference.

<a id="schema-api_testing--domains--credentials--credential_name"></a>

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

- [login_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/login_endpoint/): complete subsection reference.

- [standard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/standard/): complete subsection reference.

## Next pages

- [api_testing.domains.credentials.admin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/admin/)
- [api_testing.domains.credentials.api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/api_key/)
- [api_testing.domains.credentials.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/basic_auth/)
- [api_testing.domains.credentials.bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/bearer_token/)
- [api_testing.domains.credentials.login_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/login_endpoint/)
- [api_testing.domains.credentials.standard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/standard/)
- [api_testing.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
