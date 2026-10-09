---
page_title: "domains"
subcategory: ""
description: "Add and configure testing domains and credentials."
xcsh_docs: {"aliases": ["domains"], "body_bytes": 3443, "body_sha256": "sha256:b206a2b6414610499fe6e772d5dfc7305af22370cd8f09f3dd98e6190e2a7d83", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_testing:properties:domains:credentials"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:domains", "parent_id": "xcsh-docs:resources:api_testing:reference", "path": "documentation/resources/api_testing/properties/domains/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2223022033230132-1100002031201033-2300302100332300-0002201023313131-3033303322000202-3203110011331213-1121223222031011-2320202032031303", "registry_path": "docs/guides/resources--api_testing--reference--group-001.md", "relationships": [{"anchor": "schema-domains--domain", "enforcement": "provider-schema", "group": "domains:RequiredListObjectAttributes:credentials,domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains:RequiredListObjectAttributes:credentials,domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["domains"], "schema_version": 1, "sections": [{"aliases": ["domains allow destructive methods"], "anchor": "schema-domains--allow_destructive_methods", "description": "Enable to allow API Testing to execute against destructive methods. Use with caution as these may modify or DELETE data.", "document_id": "xcsh-docs:resources:api_testing:properties:domains", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "allow_destructive_methods"], "syntax": "attribute", "type": "bool"}, {"aliases": ["authentication", "credential setup", "credentials", "domains credentials"], "anchor": "section", "description": "Add credentials for API testing to use in the selected environment.", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:admin,standard", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:admin", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:api_key,basic_auth", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:api_key,bearer_token", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:api_key,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:api_key,basic_auth", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:basic_auth,bearer_token", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:basic_auth,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:api_key,bearer_token", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:bearer_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:basic_auth,bearer_token", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:bearer_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:bearer_token,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:bearer_token", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:api_key,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:basic_auth,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:bearer_token,login_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials:ConflictingListObjectAttributes:admin,standard", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:standard", "type": "conflicts"}, {"anchor": "schema-domains--credentials--credential_name", "enforcement": "provider-schema", "group": "domains.credentials:RequiredListObjectAttributes:credential_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials", "type": "requires"}], "schema_path": ["domains", "credentials"], "syntax": "block", "type": "object"}, {"aliases": ["domains domain"], "anchor": "schema-domains--domain", "description": "Add your testing environment domain. Be aware that running tests on a production domain can impact live applications, as API testing cannot distinguish between production and testing environments.", "document_id": "xcsh-docs:resources:api_testing:properties:domains", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "domain"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/domains/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Add and configure testing domains and credentials.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["api_testingCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/)
- domains

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Add and configure testing domains and credentials.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("credentials",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-domains--allow_destructive_methods"></a>

### allow_destructive_methods property

Type: `"bool"`. Optional.

Enable to allow API Testing to execute against destructive methods. Use with caution as these may
modify or DELETE data.

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

- [credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/): complete subsection reference.

<a id="schema-domains--domain"></a>

### domain property

Type: `"string"`. Optional.

Add your testing environment domain. Be aware that running tests on a production domain can impact
live applications, as API testing cannot distinguish between production and testing environments.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```
