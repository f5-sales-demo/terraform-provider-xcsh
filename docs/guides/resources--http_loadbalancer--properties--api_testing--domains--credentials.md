---
page_title: "api_testing.domains.credentials"
subcategory: "Load Balancing"
description: "api_testing.domains.credentials for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4662, "body_sha256": "sha256:14f0efad3bd0b5a77d7b417e0e57e9f74359c2e7a926f5031c322699c43f73f0", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:admin", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:api_key", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:basic_auth", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:bearer_token", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:login_endpoint", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials:standard"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "path": "docs/guides/resources--http_loadbalancer--properties--api_testing--domains--credentials.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_testing", "domains", "credentials"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_testing/domains/credentials/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_testing.domains.credentials for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_testing.domains.credentials

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_testing](resources--http_loadbalancer--properties--api_testing.md)
- [api_testing.domains](resources--http_loadbalancer--properties--api_testing--domains.md)
- api_testing.domains.credentials

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Add credentials for API testing to use in the selected environment.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("credential_name"),
  validators.ConflictingListObjectAttributes("admin",
    "standard"),
  validators.ConflictingListObjectAttributes("api_key",
    "basic_auth"),
  validators.ConflictingListObjectAttributes("api_key",
    "bearer_token"),
  validators.ConflictingListObjectAttributes("api_key",
    "login_endpoint"),
  validators.ConflictingListObjectAttributes("basic_auth",
    "bearer_token"),
  validators.ConflictingListObjectAttributes("basic_auth",
    "login_endpoint"),
  validators.ConflictingListObjectAttributes("bearer_token",
    "login_endpoint")}
```

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

Terraform syntax:

```terraform
credentials {
  # Configure direct properties listed below.
}
```

## Direct properties

- [admin](resources--http_loadbalancer--properties--api_testing--domains--credentials--admin.md): complete subsection reference.

- [api_key](resources--http_loadbalancer--properties--api_testing--domains--credentials--api_key.md): complete subsection reference.

- [basic_auth](resources--http_loadbalancer--properties--api_testing--domains--credentials--basic_auth.md): complete subsection reference.

- [bearer_token](resources--http_loadbalancer--properties--api_testing--domains--credentials--bearer_token.md): complete subsection reference.

<a id="schema-api_testing--domains--credentials--credential_name"></a>

### credential_name property

Type: `"string"`. Optional.

Enter a unique name for the credentials used in API testing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

- [login_endpoint](resources--http_loadbalancer--properties--api_testing--domains--credentials--login_endpoint.md): complete subsection reference.

- [standard](resources--http_loadbalancer--properties--api_testing--domains--credentials--standard.md): complete subsection reference.

## Next pages

- [api_testing.domains.credentials.admin](resources--http_loadbalancer--properties--api_testing--domains--credentials--admin.md)
- [api_testing.domains.credentials.api_key](resources--http_loadbalancer--properties--api_testing--domains--credentials--api_key.md)
- [api_testing.domains.credentials.basic_auth](resources--http_loadbalancer--properties--api_testing--domains--credentials--basic_auth.md)
- [api_testing.domains.credentials.bearer_token](resources--http_loadbalancer--properties--api_testing--domains--credentials--bearer_token.md)
- [api_testing.domains.credentials.login_endpoint](resources--http_loadbalancer--properties--api_testing--domains--credentials--login_endpoint.md)
- [api_testing.domains.credentials.standard](resources--http_loadbalancer--properties--api_testing--domains--credentials--standard.md)
- [api_testing.domains](resources--http_loadbalancer--properties--api_testing--domains.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
