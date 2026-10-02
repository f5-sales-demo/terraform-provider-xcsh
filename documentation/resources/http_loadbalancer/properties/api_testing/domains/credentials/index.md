---
page_title: "api_testing.domains.credentials"
subcategory: "Load Balancing"
description: "Add credentials for API testing to use in the selected environment."
xcsh_docs: {"aliases": ["api testing domains credentials", "authentication", "credential setup", "credentials"], "body_bytes": 5544, "body_sha256": "sha256:042f5b0d170a748672442d563793ae48cd5db3887d4a033c4d107a9110dedbdb", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:admin", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:api_key", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:standard"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "path": "documentation/resources/http_loadbalancer/properties/api_testing/domains/credentials/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_testing", "domains", "credentials"], "schema_version": 1, "sections": [{"aliases": ["admin"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:admin", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "domains", "credentials", "admin"], "syntax": "attribute", "type": "object"}, {"aliases": ["api key"], "anchor": "section", "description": "API Key", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:api_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_testing--domains--credentials--api_key--key", "enforcement": "provider-schema", "group": "api_testing.domains.credentials.api_key:RequiredObjectAttributes:key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:api_key", "type": "requires"}], "schema_path": ["api_testing", "domains", "credentials", "api_key"], "syntax": "block", "type": "object"}, {"aliases": ["authentication", "basic auth", "credential setup", "credentials"], "anchor": "section", "description": "Basic Authentication.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_testing--domains--credentials--basic_auth--user", "enforcement": "provider-schema", "group": "api_testing.domains.credentials.basic_auth:RequiredObjectAttributes:user", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth", "type": "requires"}], "schema_path": ["api_testing", "domains", "credentials", "basic_auth"], "syntax": "block", "type": "object"}, {"aliases": ["bearer token"], "anchor": "section", "description": "Configuration parameter for bearer token.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_testing", "domains", "credentials", "bearer_token"], "syntax": "block", "type": "object"}, {"aliases": ["authentication", "credential name", "credential setup", "credentials"], "anchor": "schema-api_testing--domains--credentials--credential_name", "description": "Enter a unique name for the credentials used in API testing.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "domains", "credentials", "credential_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["login", "login endpoint", "login result", "sign in"], "anchor": "section", "description": "Login Endpoint.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_testing--domains--credentials--login_endpoint--path", "enforcement": "provider-schema", "group": "api_testing.domains.credentials.login_endpoint:RequiredObjectAttributes:path,token_response_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint", "type": "requires"}, {"anchor": "schema-api_testing--domains--credentials--login_endpoint--token_response_key", "enforcement": "provider-schema", "group": "api_testing.domains.credentials.login_endpoint:RequiredObjectAttributes:path,token_response_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint", "type": "requires"}], "schema_path": ["api_testing", "domains", "credentials", "login_endpoint"], "syntax": "block", "type": "object"}, {"aliases": ["standard"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:standard", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "domains", "credentials", "standard"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_testing/domains/credentials/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Add credentials for API testing to use in the selected environment.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_testing.domains.credentials

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/)
- [api_testing.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/)
- api_testing.domains.credentials

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Add credentials for API testing to use in the selected environment.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("credential_name"),
  validators.ConflictingListObjectAttributes("admin",
    "standard"),
  validators.ConflictingListObjectAttributes("api_key",
    "basic_auth"),
  validators.ConflictingListObjectAttributes("api_key",
    "bearer_token"),
  validators.ConflictingListObjectAttributes("api_key",
    "login_endpoint"),
  validators.ConflictingListObjectAttributes("basic_auth",
    "bearer_token"),
  validators.ConflictingListObjectAttributes("basic_auth",
    "login_endpoint"),
  validators.ConflictingListObjectAttributes("bearer_token",
    "login_endpoint")}
```

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

Terraform syntax:

```terraform
credentials {
  # Configure direct properties listed below.
}
```

## Direct properties

- [admin](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/admin/): complete subsection reference.

- [api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/api_key/): complete subsection reference.

- [basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/basic_auth/): complete subsection reference.

- [bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/bearer_token/): complete subsection reference.

<a id="schema-api_testing--domains--credentials--credential_name"></a>

### credential_name property

Type: `"string"`. Optional.

Enter a unique name for the credentials used in API testing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

- [login_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/login_endpoint/): complete subsection reference.

- [standard](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/standard/): complete subsection reference.

## Next pages

- [api_testing.domains.credentials.admin](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/admin/)
- [api_testing.domains.credentials.api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/api_key/)
- [api_testing.domains.credentials.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/basic_auth/)
- [api_testing.domains.credentials.bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/bearer_token/)
- [api_testing.domains.credentials.login_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/login_endpoint/)
- [api_testing.domains.credentials.standard](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/standard/)
- [api_testing.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
