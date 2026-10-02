---
page_title: "api_testing.domains.credentials.login_endpoint.json_payload"
subcategory: "Load Balancing"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["api testing domains credentials login endpoint json payload", "authentication", "credential setup", "credentials", "login", "login result", "sign in"], "body_bytes": 3139, "body_sha256": "sha256:31ede381b179c9ef5cce416c4eee46be6ef55ad07e5307639dd2eaa7a91ea30a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint:json_payload:blindfold_secret_info", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint:json_payload:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint:json_payload", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint", "path": "documentation/resources/http_loadbalancer/properties/api_testing/domains/credentials/login_endpoint/json_payload/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3221130131220320-2123002333113211-3221300230021010-2033011110202030-1212202322200031-2201321010023033-0012320112101213-0212220311112303", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials.login_endpoint.json_payload:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint:json_payload:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains.credentials.login_endpoint.json_payload:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint:json_payload:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_testing", "domains", "credentials", "login_endpoint", "json_payload"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint:json_payload:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_testing--domains--credentials--login_endpoint--json_payload--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint:json_payload:blindfold_secret_info", "type": "requires"}], "schema_path": ["api_testing", "domains", "credentials", "login_endpoint", "json_payload", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint:json_payload:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_testing--domains--credentials--login_endpoint--json_payload--clear_secret_info--url", "enforcement": "provider-schema", "group": "api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint:json_payload:clear_secret_info", "type": "requires"}], "schema_path": ["api_testing", "domains", "credentials", "login_endpoint", "json_payload", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_testing/domains/credentials/login_endpoint/json_payload/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_testing.domains.credentials.login_endpoint.json_payload

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/)
- [api_testing.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/)
- [api_testing.domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/)
- [api_testing.domains.credentials.login_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/login_endpoint/)
- api_testing.domains.credentials.login_endpoint.json_payload

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
json_payload {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/login_endpoint/json_payload/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/login_endpoint/json_payload/clear_secret_info/): complete subsection reference.

## Next pages

- [api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/login_endpoint/json_payload/blindfold_secret_info/)
- [api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/login_endpoint/json_payload/clear_secret_info/)
- [api_testing.domains.credentials.login_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/credentials/login_endpoint/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
