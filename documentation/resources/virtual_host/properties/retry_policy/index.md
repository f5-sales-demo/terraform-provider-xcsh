---
page_title: "retry_policy"
subcategory: ""
description: "Retry policy configuration for route destination."
xcsh_docs: {"aliases": ["retry policy"], "body_bytes": 7869, "body_sha256": "sha256:3b8e3508fec0cca3ab3a4c1eb16fe1cdc971ffed4641f35922d72c7eb067fba5", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:properties:retry_policy:back_off"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:retry_policy", "parent_id": "xcsh-docs:resources:virtual_host:reference", "path": "documentation/resources/virtual_host/properties/retry_policy/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3120210313011012-3220120012303030-2320221131003020-2003111301323011-0131032003311130-3311130033201323-1332312222032333-0100213232033200", "registry_path": "docs/guides/resources--virtual_host--reference--group-003.md", "relationships": [{"anchor": "schema-retry_policy--retry_condition", "enforcement": "provider-schema", "group": "retry_policy:RequiredObjectAttributes:retry_condition", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:retry_policy", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["retry_policy"], "schema_version": 1, "sections": [{"aliases": ["back off"], "anchor": "section", "description": "Specifies parameters that control retry back off.", "document_id": "xcsh-docs:resources:virtual_host:properties:retry_policy:back_off", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["retry_policy", "back_off"], "syntax": "block", "type": "object"}, {"aliases": ["num retries"], "anchor": "schema-retry_policy--num_retries", "description": "Specifies the allowed number of retries. Defaults to 1. Retries can be done any number of times. An exponential back-off algorithm is used between each retry.", "document_id": "xcsh-docs:resources:virtual_host:properties:retry_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["retry_policy", "num_retries"], "syntax": "attribute", "type": "number"}, {"aliases": ["duration", "operation timeout", "per try timeout"], "anchor": "schema-retry_policy--per_try_timeout", "description": "Specifies a non-zero timeout per retry attempt. In milliseconds.", "document_id": "xcsh-docs:resources:virtual_host:properties:retry_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["retry_policy", "per_try_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["retriable status codes"], "anchor": "schema-retry_policy--retriable_status_codes", "description": "HTTP status codes that should trigger a retry in addition to those specified by retry_on.", "document_id": "xcsh-docs:resources:virtual_host:properties:retry_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["retry_policy", "retriable_status_codes"], "syntax": "attribute", "type": "list"}, {"aliases": ["duration", "operation timeout", "retry condition"], "anchor": "schema-retry_policy--retry_condition", "description": "Specifies the conditions under which retry takes place. Retries can be on different types of condition depending on application requirements. For example, network failure, all 5xx response codes, idempotent 4xx response codes, etc The possible values are \"5xx\" : Retry will be done if the upstream server responds with", "document_id": "xcsh-docs:resources:virtual_host:properties:retry_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["retry_policy", "retry_condition"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/retry_policy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Retry policy configuration for route destination.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# retry_policy

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- retry_policy

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

- [back_off](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/retry_policy/back_off/): complete subsection reference.

<a id="schema-retry_policy--num_retries"></a>

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

<a id="schema-retry_policy--per_try_timeout"></a>

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

<a id="schema-retry_policy--retriable_status_codes"></a>

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

<a id="schema-retry_policy--retry_condition"></a>

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

- [retry_policy.back_off](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/retry_policy/back_off/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
