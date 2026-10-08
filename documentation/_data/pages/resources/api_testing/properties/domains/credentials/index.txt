---
page_title: "domains.credentials"
subcategory: ""
description: "Add credentials for API testing to use in the selected environment."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials", "domains credentials"], "body_bytes": 3901, "body_sha256": "sha256:117514667a52fafd3a8dbb0891a69b75b335b6836210de838673452123ec0ff8", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_testing:properties:domains:credentials:admin", "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth", "xcsh-docs:resources:api_testing:properties:domains:credentials:bearer_token", "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "xcsh-docs:resources:api_testing:properties:domains:credentials:standard"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:domains:credentials", "parent_id": "xcsh-docs:resources:api_testing:properties:domains", "path": "documentation/resources/api_testing/properties/domains/credentials/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1203220111020303-0212211321110103-2232032300111100-0122020313110201-1322013123200132-3132023321123022-3133220101003103-0222223211321301", "registry_path": "docs/guides/resources--api_testing--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:admin,standard", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:admin", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:api_key,basic_auth", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:api_key,bearer_token", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:api_key,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:api_key,basic_auth", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:basic_auth,bearer_token", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:basic_auth,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:api_key,bearer_token", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:bearer_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:basic_auth,bearer_token", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:bearer_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:bearer_token,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:bearer_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:api_key,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:basic_auth,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:bearer_token,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:admin,standard", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:standard", "type": "conflicts"}, {"anchor": "schema-domains--credentials--credential_name", "enforcement": "provider-schema", "group": "domains.credentials:RequiredListObjectAttributes:credential_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["domains", "credentials"], "schema_version": 1, "sections": [{"aliases": ["domains credentials admin"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:admin", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "credentials", "admin"], "syntax": "attribute", "type": "object"}, {"aliases": ["domains credentials api key"], "anchor": "section", "description": "API Key", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-domains--credentials--api_key--key", "enforcement": "provider-schema", "group": "domains.credentials.api_key:RequiredObjectAttributes:key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "type": "requires"}], "schema_path": ["domains", "credentials", "api_key"], "syntax": "block", "type": "object"}, {"aliases": ["domains credentials basic auth"], "anchor": "section", "description": "Basic Authentication.", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-domains--credentials--basic_auth--user", "enforcement": "provider-schema", "group": "domains.credentials.basic_auth:RequiredObjectAttributes:user", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth", "type": "requires"}], "schema_path": ["domains", "credentials", "basic_auth"], "syntax": "block", "type": "object"}, {"aliases": ["domains credentials bearer token"], "anchor": "section", "description": "Configuration parameter for bearer token.", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:bearer_token", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "credentials", "bearer_token"], "syntax": "block", "type": "object"}, {"aliases": ["authentication", "credential setup", "credentials", "domains credentials credential name"], "anchor": "schema-domains--credentials--credential_name", "description": "Enter a unique name for the credentials used in API testing.", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "credentials", "credential_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["domains credentials login endpoint", "login", "login result", "sign in"], "anchor": "section", "description": "Login Endpoint.", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-domains--credentials--login_endpoint--path", "enforcement": "provider-schema", "group": "domains.credentials.login_endpoint:RequiredObjectAttributes:path,token_response_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "type": "requires"}, {"anchor": "schema-domains--credentials--login_endpoint--token_response_key", "enforcement": "provider-schema", "group": "domains.credentials.login_endpoint:RequiredObjectAttributes:path,token_response_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "type": "requires"}], "schema_path": ["domains", "credentials", "login_endpoint"], "syntax": "block", "type": "object"}, {"aliases": ["domains credentials standard"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:standard", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "credentials", "standard"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/domains/credentials/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Add credentials for API testing to use in the selected environment.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["api_testingCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/)
- domains.credentials

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Add credentials for API testing to use in the selected environment.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [admin](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/admin/): complete subsection reference.

- [api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/): complete subsection reference.

- [basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/basic_auth/): complete subsection reference.

- [bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/bearer_token/): complete subsection reference.

<a id="schema-domains--credentials--credential_name"></a>

### credential_name property

Type: `"string"`. Optional.

Enter a unique name for the credentials used in API testing.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [login_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/login_endpoint/): complete subsection reference.

- [standard](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/standard/): complete subsection reference.
