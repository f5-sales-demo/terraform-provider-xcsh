---
page_title: "api_testing.domains.credentials.login_endpoint.json_payload"
subcategory: "Load Balancing"
description: "api_testing.domains.credentials.login_endpoint.json_payload for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2281, "body_sha256": "sha256:95e43b3fe1b62508bd0ceac8480d028cdfa99bae565bf723488e9ac29b3a276e", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint:json_payload", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint:json_payload:blindfold_secret_info", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint:json_payload:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint:json_payload", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_testing--domains--credentials--login_endpoint--json_payload.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_testing", "domains", "credentials", "login_endpoint", "json_payload"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/login_endpoint/json_payload/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_testing.domains.credentials.login_endpoint.json_payload for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_testing.domains.credentials.login_endpoint.json_payload

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_testing](data-sources--http_loadbalancer--properties--api_testing.md)
- [api_testing.domains](data-sources--http_loadbalancer--properties--api_testing--domains.md)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--properties--api_testing--domains--credentials.md)
- [api_testing.domains.credentials.login_endpoint](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--login_endpoint.md)
- api_testing.domains.credentials.login_endpoint.json_payload

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

## Direct properties

- [blindfold_secret_info](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--login_endpoint--json_payload--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--login_endpoint--json_payload--clear_secret_info.md): complete subsection reference.

## Next pages

- [api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--login_endpoint--json_payload--blindfold_secret_info.md)
- [api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--login_endpoint--json_payload--clear_secret_info.md)
- [api_testing.domains.credentials.login_endpoint](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--login_endpoint.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
