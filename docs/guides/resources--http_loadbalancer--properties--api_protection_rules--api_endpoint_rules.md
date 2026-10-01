---
page_title: "api_protection_rules.api_endpoint_rules"
subcategory: "Load Balancing"
description: "api_protection_rules.api_endpoint_rules for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 6266, "body_sha256": "sha256:ad03e2085bf9e6c27f576b514f83d23d2d23caa3f282026dcde8e18887e4d370", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:action", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:any_domain", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:api_endpoint_method", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:client_matcher", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:metadata", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules", "path": "docs/guides/resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_protection_rules", "api_endpoint_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_protection_rules.api_endpoint_rules for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_endpoint_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_protection_rules](resources--http_loadbalancer--properties--api_protection_rules.md)
- api_protection_rules.api_endpoint_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Category defines specific rules per API endpoints. If request matches any of these rules, skipping
second category rules.

Upstream description:

This category defines specific rules per API endpoints. If request matches any of these rules,
skipping second category rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("api_endpoint_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

Terraform syntax:

```terraform
api_endpoint_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [action](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--action.md): complete subsection reference.

- [any_domain](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--any_domain.md): complete subsection reference.

- [api_endpoint_method](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--api_endpoint_method.md): complete subsection reference.

<a id="schema-api_protection_rules--api_endpoint_rules--api_endpoint_path"></a>

### api_endpoint_path property

Type: `"string"`. Optional.

API Endpoint. The endpoint (path) of the request.

Upstream description:

The endpoint (path) of the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

- [client_matcher](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher.md): complete subsection reference.

- [metadata](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--metadata.md): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--request_matcher.md): complete subsection reference.

<a id="schema-api_protection_rules--api_endpoint_rules--specific_domain"></a>

### specific_domain property

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

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
    "format": "fqdn",
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

## Next pages

- [api_protection_rules.api_endpoint_rules.action](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--action.md)
- [api_protection_rules.api_endpoint_rules.any_domain](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--any_domain.md)
- [api_protection_rules.api_endpoint_rules.api_endpoint_method](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--api_endpoint_method.md)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--client_matcher.md)
- [api_protection_rules.api_endpoint_rules.metadata](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--metadata.md)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules--request_matcher.md)
- [api_protection_rules](resources--http_loadbalancer--properties--api_protection_rules.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
