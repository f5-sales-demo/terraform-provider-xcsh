---
page_title: "api_rate_limit.api_endpoint_rules"
subcategory: "Load Balancing"
description: "api_rate_limit.api_endpoint_rules for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 5945, "body_sha256": "sha256:ec451411f3f60bfd12aadc3d196e330b531074972af59d32365e7cc3612b01e1", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:any_domain", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:api_endpoint_method", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:ref_rate_limiter", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit", "path": "docs/guides/resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "api_endpoint_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.api_endpoint_rules for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_rate_limit.api_endpoint_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_rate_limit](resources--http_loadbalancer--properties--api_rate_limit.md)
- api_rate_limit.api_endpoint_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate\_limiter\_choice:
inline\_rate\_limiter or ref\_rate\_limiter.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("api_endpoint_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("inline_rate_limiter",
    "ref_rate_limiter")}
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

- [any_domain](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--any_domain.md): complete subsection reference.

- [api_endpoint_method](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--api_endpoint_method.md): complete subsection reference.

<a id="schema-api_rate_limit--api_endpoint_rules--api_endpoint_path"></a>

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

- [client_matcher](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher.md): complete subsection reference.

- [inline_rate_limiter](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--inline_rate_limiter.md): complete subsection reference.

- [ref_rate_limiter](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--ref_rate_limiter.md): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher.md): complete subsection reference.

<a id="schema-api_rate_limit--api_endpoint_rules--specific_domain"></a>

### specific_domain property

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

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

- [api_rate_limit.api_endpoint_rules.any_domain](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--any_domain.md)
- [api_rate_limit.api_endpoint_rules.api_endpoint_method](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--api_endpoint_method.md)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher.md)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--inline_rate_limiter.md)
- [api_rate_limit.api_endpoint_rules.ref_rate_limiter](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--ref_rate_limiter.md)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher.md)
- [api_rate_limit](resources--http_loadbalancer--properties--api_rate_limit.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
