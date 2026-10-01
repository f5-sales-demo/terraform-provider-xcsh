---
page_title: "api_testing.domains.credentials.basic_auth"
subcategory: "Load Balancing"
description: "api_testing.domains.credentials.basic_auth for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2918, "body_sha256": "sha256:8d04640c39e58d5b7b6143a1f91cfcd7e5e23c5df094c0ce190931fb06ec7a2d", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth:password"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials", "path": "docs/guides/resources--http_loadbalancer--properties--api_testing--domains--credentials--basic_auth.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_testing", "domains", "credentials", "basic_auth"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_testing/domains/credentials/basic_auth/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_testing.domains.credentials.basic_auth for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_testing.domains.credentials.basic_auth

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_testing](resources--http_loadbalancer--properties--api_testing.md)
- [api_testing.domains](resources--http_loadbalancer--properties--api_testing--domains.md)
- [api_testing.domains.credentials](resources--http_loadbalancer--properties--api_testing--domains--credentials.md)
- api_testing.domains.credentials.basic_auth

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Basic Authentication.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("user")}
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
basic_auth {
  # Configure direct properties listed below.
}
```

## Direct properties

- [password](resources--http_loadbalancer--properties--api_testing--domains--credentials--basic_auth--password.md): complete subsection reference.

<a id="schema-api_testing--domains--credentials--basic_auth--user"></a>

### user property

Type: `"string"`. Optional.

User. Configuration parameter for user

Upstream description:

Configuration parameter for user

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
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

## Next pages

- [api_testing.domains.credentials.basic_auth.password](resources--http_loadbalancer--properties--api_testing--domains--credentials--basic_auth--password.md)
- [api_testing.domains.credentials](resources--http_loadbalancer--properties--api_testing--domains--credentials.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
