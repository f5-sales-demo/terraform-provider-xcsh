---
page_title: "domains.credentials.api_key.value"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["domains credentials api key value"], "body_bytes": 2597, "body_sha256": "sha256:785844890f48463bb5a4e480d10f5b2314e0d2260f777dd2516066cd08dd3707", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:blindfold_secret_info", "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:clear_secret_info"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value", "parent_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key", "path": "documentation/resources/api_testing/properties/domains/credentials/api_key/value/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3112312300210202-3311001121123231-0122332323322133-3231220321010003-0113121030320031-2220103203012320-3030121332213200-3312201110021123", "registry_path": "docs/guides/resources--api_testing--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials.api_key.value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials.api_key.value:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["domains", "credentials", "api_key", "value"], "schema_version": 1, "sections": [{"aliases": ["domains credentials api key value blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-domains--credentials--api_key--value--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "domains.credentials.api_key.value.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:blindfold_secret_info", "type": "requires"}], "schema_path": ["domains", "credentials", "api_key", "value", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["domains credentials api key value clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-domains--credentials--api_key--value--clear_secret_info--url", "enforcement": "provider-schema", "group": "domains.credentials.api_key.value.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:api_key:value:clear_secret_info", "type": "requires"}], "schema_path": ["domains", "credentials", "api_key", "value", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/domains/credentials/api_key/value/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["api_testingCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials.api_key.value

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/)
- [domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/)
- [domains.credentials.api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/)
- domains.credentials.api_key.value

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
value {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/value/clear_secret_info/): complete subsection reference.

## Next pages

- [domains.credentials.api_key.value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/value/blindfold_secret_info/)
- [domains.credentials.api_key.value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/value/clear_secret_info/)
- [domains.credentials.api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/api_key/)
- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
