---
page_title: "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules"
subcategory: "Load Balancing"
description: "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 7393, "body_sha256": "sha256:00b40a6d9b70162e2e5e5e40c05529aa4eb4ae7a4765c4849ffa72689ce9a759", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:any_domain", "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:any_url", "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:api_endpoint", "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:api_groups", "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:client_matcher", "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:request_matcher"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "path": "docs/guides/resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules", "bypass_rate_limiting_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [api_rate_limit](resources--cdn_loadbalancer--properties--api_rate_limit.md)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules.md)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Upstream description:

This category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("any_url",
    "api_endpoint"),
  validators.ConflictingListObjectAttributes("any_url",
    "api_groups"),
  validators.ConflictingListObjectAttributes("any_url",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "api_groups"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_groups",
    "base_path")}
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
bypass_rate_limiting_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_domain](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--any_domain.md): complete subsection reference.

- [any_url](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--any_url.md): complete subsection reference.

- [api_endpoint](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--api_endpoint.md): complete subsection reference.

- [api_groups](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--api_groups.md): complete subsection reference.

<a id="schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--base_path"></a>

### base_path property

Type: `"string"`. Optional.

Exclusive with \[any\_url api\_endpoint api\_groups\] The base path which this validation applies
to.

Upstream description:

Exclusive with \[any\_url api\_endpoint api\_groups\] The base path which this validation applies
to.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--client_matcher.md): complete subsection reference.

- [request_matcher](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher.md): complete subsection reference.

<a id="schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--specific_domain"></a>

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

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--any_domain.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--any_url.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--api_endpoint.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--api_groups.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--client_matcher.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher.md)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
