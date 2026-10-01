---
page_title: "api_testing.domains.credentials"
subcategory: "Load Balancing"
description: "api_testing.domains.credentials for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3755, "body_sha256": "sha256:69d5506c56ed2009ea41e427b7301f2e4e1b7236ae6f4989aba4503e13bf8883", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:admin", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:api_key", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:standard"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_testing--domains--credentials.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_testing", "domains", "credentials"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_testing.domains.credentials for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_testing.domains.credentials

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_testing](data-sources--http_loadbalancer--properties--api_testing.md)
- [api_testing.domains](data-sources--http_loadbalancer--properties--api_testing--domains.md)
- api_testing.domains.credentials

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

- [admin](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--admin.md): complete subsection reference.

- [api_key](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--api_key.md): complete subsection reference.

- [basic_auth](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--basic_auth.md): complete subsection reference.

- [bearer_token](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--bearer_token.md): complete subsection reference.

<a id="schema-api_testing--domains--credentials--credential_name"></a>

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

- [login_endpoint](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--login_endpoint.md): complete subsection reference.

- [standard](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--standard.md): complete subsection reference.

## Next pages

- [api_testing.domains.credentials.admin](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--admin.md)
- [api_testing.domains.credentials.api_key](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--api_key.md)
- [api_testing.domains.credentials.basic_auth](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--basic_auth.md)
- [api_testing.domains.credentials.bearer_token](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--bearer_token.md)
- [api_testing.domains.credentials.login_endpoint](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--login_endpoint.md)
- [api_testing.domains.credentials.standard](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--standard.md)
- [api_testing.domains](data-sources--http_loadbalancer--properties--api_testing--domains.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
