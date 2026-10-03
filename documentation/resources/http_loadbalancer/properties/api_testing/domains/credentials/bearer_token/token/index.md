---
page_title: "api_testing.domains.credentials.bearer_token.token"
subcategory: "Load Balancing"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["api testing domains credentials bearer token token"], "body_bytes": 3052, "body_sha256": "sha256:838906951c247a03a43bf6eff31cf3b926c33265c8e471268bc01386c92be4ce", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token:token:blindfold_secret_info", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token:token:clear_secret_info"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token:token", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token", "path": "documentation/resources/http_loadbalancer/properties/api_testing/domains/credentials/bearer_token/token/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3223203032111230-2332221333133011-0012311033133223-3010031123130122-2111202011233232-0303301122313233-3020123022030000-3220203310330102", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials.bearer_token.token:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token:token:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials.bearer_token.token:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token:token:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_testing", "domains", "credentials", "bearer_token", "token"], "schema_version": 1, "sections": [{"aliases": ["api testing domains credentials bearer token token blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token:token:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_testing--domains--credentials--bearer_token--token--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "api_testing.domains.credentials.bearer_token.token.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token:token:blindfold_secret_info", "type": "requires"}], "schema_path": ["api_testing", "domains", "credentials", "bearer_token", "token", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["api testing domains credentials bearer token token clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token:token:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_testing--domains--credentials--bearer_token--token--clear_secret_info--url", "enforcement": "provider-schema", "group": "api_testing.domains.credentials.bearer_token.token.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token:token:clear_secret_info", "type": "requires"}], "schema_path": ["api_testing", "domains", "credentials", "bearer_token", "token", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_testing/domains/credentials/bearer_token/token/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_testing.domains.credentials.bearer_token.token

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/)
- [api_testing.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/)
- [api_testing.domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/)
- [api_testing.domains.credentials.bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/bearer_token/)
- api_testing.domains.credentials.bearer_token.token

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
token {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/bearer_token/token/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/bearer_token/token/clear_secret_info/): complete subsection reference.

## Next pages

- [api_testing.domains.credentials.bearer_token.token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/bearer_token/token/blindfold_secret_info/)
- [api_testing.domains.credentials.bearer_token.token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/bearer_token/token/clear_secret_info/)
- [api_testing.domains.credentials.bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/bearer_token/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
