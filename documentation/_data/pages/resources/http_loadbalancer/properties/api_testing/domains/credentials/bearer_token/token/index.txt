---
page_title: "api_testing.domains.credentials.bearer_token.token"
subcategory: "Load Balancing"
description: "api_testing.domains.credentials.bearer_token.token for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3052, "body_sha256": "sha256:838906951c247a03a43bf6eff31cf3b926c33265c8e471268bc01386c92be4ce", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token:token:blindfold_secret_info", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token:token:clear_secret_info"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token:token", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token", "path": "documentation/resources/http_loadbalancer/properties/api_testing/domains/credentials/bearer_token/token/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["api_testing", "domains", "credentials", "bearer_token", "token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_testing/domains/credentials/bearer_token/token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_testing.domains.credentials.bearer_token.token for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
