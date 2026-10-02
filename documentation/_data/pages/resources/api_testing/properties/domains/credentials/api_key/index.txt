---
page_title: "domains.credentials.api_key"
subcategory: ""
description: "API Key"
xcsh_docs: {"aliases": ["domains credentials api key"], "body_bytes": 2786, "body_sha256": "sha256:c51b8ae8892dcd068498d9808c4e82853b53fcc591324f5ff9874086986d602a", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "parent_id": "xcsh-docs:resources:api_testing:properties:domains:credentials", "path": "documentation/resources/api_testing/properties/domains/credentials/api_key/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1310322331313200-2120112223003333-0312311302230121-2213021111301031-2322202103201013-1111001112101213-0113333213002202-2021030011001022", "registry_path": "docs/guides/resources--api_testing--reference--group-001.md", "relationships": [{"anchor": "schema-domains--credentials--api_key--key", "enforcement": "provider-schema", "group": "domains.credentials.api_key:RequiredObjectAttributes:key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["domains", "credentials", "api_key"], "schema_version": 1, "sections": [{"aliases": ["key"], "anchor": "schema-domains--credentials--api_key--key", "description": "Cryptographic key material", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "credentials", "api_key", "key"], "syntax": "attribute", "type": "string"}, {"aliases": ["value"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials.api_key.value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials.api_key.value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:clear_secret_info", "type": "conflicts"}], "schema_path": ["domains", "credentials", "api_key", "value"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/domains/credentials/api_key/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "API Key", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_testingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials.api_key

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/)
- [domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/)
- domains.credentials.api_key

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

API Key

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("key")}
```

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

Terraform syntax:

```terraform
api_key {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-domains--credentials--api_key--key"></a>

### key property

Type: `"string"`. Optional.

Key. Cryptographic key material

Upstream description:

Cryptographic key material

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/value/): complete subsection reference.

## Next pages

- [domains.credentials.api_key.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/value/)
- [domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/)
- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
