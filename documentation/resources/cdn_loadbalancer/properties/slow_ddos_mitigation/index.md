---
page_title: "slow_ddos_mitigation"
subcategory: "Load Balancing"
description: "\"Slow and low\" attacks tie up server resources, leaving none available for servicing requests from actual users."
xcsh_docs: {"aliases": ["slow ddos mitigation"], "body_bytes": 5051, "body_sha256": "sha256:76d984f7d8eeab19dc12c6500e3d5455328ac2e2929fb95a56aa6bc4369bb472", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:slow_ddos_mitigation:disable_request_timeout"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:slow_ddos_mitigation", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/slow_ddos_mitigation/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1233003020202211-1121000220031121-0002002321202011-3200220102103002-3312211001303013-1111001323333210-3232211013001110-3302023302323112", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-014.md", "relationships": [{"anchor": "schema-slow_ddos_mitigation--request_timeout", "enforcement": "provider-schema", "group": "slow_ddos_mitigation:ConflictingObjectAttributes:disable_request_timeout,request_timeout", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:slow_ddos_mitigation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "slow_ddos_mitigation:ConflictingObjectAttributes:disable_request_timeout,request_timeout", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:slow_ddos_mitigation:disable_request_timeout", "type": "conflicts"}, {"anchor": "schema-slow_ddos_mitigation--request_headers_timeout", "enforcement": "provider-schema", "group": "slow_ddos_mitigation:RequiredObjectAttributes:request_headers_timeout", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:slow_ddos_mitigation", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["slow_ddos_mitigation"], "schema_version": 1, "sections": [{"aliases": ["disable request timeout", "duration", "operation timeout"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:slow_ddos_mitigation:disable_request_timeout", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["slow_ddos_mitigation", "disable_request_timeout"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "operation timeout", "request headers timeout"], "anchor": "schema-slow_ddos_mitigation--request_headers_timeout", "description": "The amount of time the client has to send only the headers on the request stream before the stream is cancelled. The default value is 10000 milliseconds. This setting provides protection against Slowloris attacks.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:slow_ddos_mitigation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["slow_ddos_mitigation", "request_headers_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["duration", "operation timeout", "request timeout"], "anchor": "schema-slow_ddos_mitigation--request_timeout", "description": "Exclusive with", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:slow_ddos_mitigation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["slow_ddos_mitigation", "request_timeout"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/slow_ddos_mitigation/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "\"Slow and low\" attacks tie up server resources, leaving none available for servicing requests from actual users.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# slow_ddos_mitigation

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- slow_ddos_mitigation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: slow\_ddos\_mitigation, system\_default\_timeouts; Default: system\_default\_timeouts\]
'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Upstream description:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("request_headers_timeout"),
  validators.ConflictingObjectAttributes("disable_request_timeout",
    "request_timeout")}
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
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

OneOf alternatives in this subsection:

- [slow_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/slow_ddos_mitigation/#section)
- [system_default_timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/system_default_timeouts/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
slow_ddos_mitigation {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_request_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/slow_ddos_mitigation/disable_request_timeout/): complete subsection reference.

<a id="schema-slow_ddos_mitigation--request_headers_timeout"></a>

### request_headers_timeout property

Type: `"number"`. Optional.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Upstream description:

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The default value is 10000 milliseconds. This setting provides protection against
Slowloris attacks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2000, 30000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="schema-slow_ddos_mitigation--request_timeout"></a>

### request_timeout property

Type: `"number"`. Optional.

Exclusive with \[disable\_request\_timeout\].

Upstream description:

Exclusive with \[disable\_request\_timeout\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2000, 300000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

## Next pages

- [slow_ddos_mitigation.disable_request_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/slow_ddos_mitigation/disable_request_timeout/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
