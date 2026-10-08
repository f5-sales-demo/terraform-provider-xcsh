---
page_title: "domains.credentials"
subcategory: ""
description: "Add credentials for API testing to use in the selected environment."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials", "domains credentials"], "body_bytes": 2908, "body_sha256": "sha256:8f4051e277365bb3f33cd4f6b451752a80856dc0474589457a65ac8a887191b5", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_testing:properties:domains:credentials:admin", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:api_key", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:login_endpoint", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:standard"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials", "parent_id": "xcsh-docs:data-sources:api_testing:properties:domains", "path": "documentation/data-sources/api_testing/properties/domains/credentials/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112", "registry_path": "docs/guides/data-sources--api_testing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["domains", "credentials"], "schema_version": 1, "sections": [{"aliases": ["domains credentials admin"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:admin", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "credentials", "admin"], "syntax": "attribute", "type": "object"}, {"aliases": ["domains credentials api key"], "anchor": "section", "description": "API Key", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:api_key", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "credentials", "api_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["domains credentials basic auth"], "anchor": "section", "description": "Basic Authentication.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "credentials", "basic_auth"], "syntax": "attribute", "type": "object"}, {"aliases": ["domains credentials bearer token"], "anchor": "section", "description": "Configuration parameter for bearer token.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "credentials", "bearer_token"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication", "credential setup", "credentials", "domains credentials credential name"], "anchor": "schema-domains--credentials--credential_name", "description": "Enter a unique name for the credentials used in API testing.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "credentials", "credential_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["domains credentials login endpoint", "login", "login result", "sign in"], "anchor": "section", "description": "Login Endpoint.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:login_endpoint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "credentials", "login_endpoint"], "syntax": "attribute", "type": "object"}, {"aliases": ["domains credentials standard"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:standard", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "credentials", "standard"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/properties/domains/credentials/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Add credentials for API testing to use in the selected environment.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["api_testingCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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

- [login_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/login_endpoint/): complete subsection reference.

- [standard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/standard/): complete subsection reference.
