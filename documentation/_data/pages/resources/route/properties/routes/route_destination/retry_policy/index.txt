---
page_title: "routes.route_destination.retry_policy"
subcategory: ""
description: "Retry policy configuration for route destination."
xcsh_docs: {"aliases": ["routes route destination retry policy"], "body_bytes": 8319, "body_sha256": "sha256:29fab8f4c48c1e3d1b4cfff8a0ca2b312bfafa72017ba33319a555521c1b590a", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:route_destination:retry_policy:back_off"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "documentation/resources/route/properties/routes/route_destination/retry_policy/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0301103202310121-1322001133231203-3330201331230333-2121120301200322-3222320330230320-1021310030201212-1120112023110002-2203110113210212", "registry_path": "docs/guides/resources--route--reference--group-002.md", "relationships": [{"anchor": "schema-routes--route_destination--retry_policy--retry_condition", "enforcement": "provider-schema", "group": "routes.route_destination.retry_policy:RequiredObjectAttributes:retry_condition", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "retry_policy"], "schema_version": 1, "sections": [{"aliases": ["back off"], "anchor": "section", "description": "Specifies parameters that control retry back off.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy:back_off", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "route_destination", "retry_policy", "back_off"], "syntax": "block", "type": "object"}, {"aliases": ["num retries"], "anchor": "schema-routes--route_destination--retry_policy--num_retries", "description": "Specifies the allowed number of retries. Defaults to 1. Retries can be done any number of times. An exponential back-off algorithm is used between each retry.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "retry_policy", "num_retries"], "syntax": "attribute", "type": "number"}, {"aliases": ["duration", "operation timeout", "per try timeout"], "anchor": "schema-routes--route_destination--retry_policy--per_try_timeout", "description": "Specifies a non-zero timeout per retry attempt. In milliseconds.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "retry_policy", "per_try_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["retriable status codes"], "anchor": "schema-routes--route_destination--retry_policy--retriable_status_codes", "description": "HTTP status codes that should trigger a retry in addition to those specified by retry_on.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "retry_policy", "retriable_status_codes"], "syntax": "attribute", "type": "list"}, {"aliases": ["duration", "operation timeout", "retry condition"], "anchor": "schema-routes--route_destination--retry_policy--retry_condition", "description": "Specifies the conditions under which retry takes place. Retries can be on different types of condition depending on application requirements. For example, network failure, all 5xx response codes, idempotent 4xx response codes, etc The possible values are \"5xx\" : Retry will be done if the upstream server responds with", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "retry_policy", "retry_condition"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/retry_policy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Retry policy configuration for route destination.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["routeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.retry_policy

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- routes.route_destination.retry_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Retry policy configuration for route destination.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("retry_condition")}
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
retry_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [back_off](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retry_policy/back_off/): complete subsection reference.

<a id="schema-routes--route_destination--retry_policy--num_retries"></a>

### num_retries property

Type: `"number"`. Optional.

Specifies the allowed number of retries. Retries can be done any number of times. An exponential
back-off algorithm is used between each retry. Defaults to \`1\`.

Upstream description:

Specifies the allowed number of retries. Defaults to 1. Retries can be done any number of times. An
exponential back-off algorithm is used between each retry.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  }
}
```

<a id="schema-routes--route_destination--retry_policy--per_try_timeout"></a>

### per_try_timeout property

Type: `"number"`. Optional.

Specifies a non-zero timeout per retry attempt. In milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="schema-routes--route_destination--retry_policy--retriable_status_codes"></a>

### retriable_status_codes property

Type: `["list", "number"]`. Optional.

HTTP status codes that should trigger a retry in addition to those specified by retry\_on.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-routes--route_destination--retry_policy--retry_condition"></a>

### retry_condition property

Type: `["list", "string"]`. Optional.

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if
the..

Upstream description:

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc

The possible values are

"5xx" : Retry will be done if the upstream server responds with any 5xx response code, or does not
respond at all (disconnect/reset/read timeout).

"gateway-error" : Retry will be done only if the upstream server responds with 502, 503 or 504
responses (Included in 5xx)

"connect-failure" : Retry will be done if the request fails because of a connection failure to the
upstream server (connect timeout, etc.). (Included in 5xx)

"refused-stream" : Retry is done if the upstream server resets the stream with a REFUSED\_STREAM
error code (Included in 5xx)

"retriable-4xx" : Retry is done if the upstream server responds with a retriable 4xx response code.
The only response code in this category is HTTP CONFLICT (409)

"retriable-status-codes" : Retry is done if the upstream server responds with any response code
matching one defined in retriable\_status\_codes field

"reset" : Retry is done if the upstream server does not respond at all (disconnect/reset/read
timeout.)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 7),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 7,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 7,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [routes.route_destination.retry_policy.back_off](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retry_policy/back_off/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
