---
page_title: "slow_ddos_mitigation"
subcategory: "Load Balancing"
description: "\"Slow and low\" attacks tie up server resources, leaving none available for servicing requests from actual users."
xcsh_docs: {"aliases": ["slow ddos mitigation"], "body_bytes": 4393, "body_sha256": "sha256:8916efb7c23d1a92387ee4e3fad9e64f3b3ade97c6b8341dc376bdd5acaed4a3", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:slow_ddos_mitigation:disable_request_timeout"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:slow_ddos_mitigation", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/slow_ddos_mitigation/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3333311120100023-2230031010023220-2013013300321213-0302220321233313-0222113033133022-2131111320010323-2303320333133122-3002301330000322", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [{"anchor": "schema-slow_ddos_mitigation--request_timeout", "enforcement": "provider-schema", "group": "slow_ddos_mitigation:ConflictingObjectAttributes:disable_request_timeout,request_timeout", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:slow_ddos_mitigation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "slow_ddos_mitigation:ConflictingObjectAttributes:disable_request_timeout,request_timeout", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:slow_ddos_mitigation:disable_request_timeout", "type": "conflicts"}, {"anchor": "schema-slow_ddos_mitigation--request_headers_timeout", "enforcement": "provider-schema", "group": "slow_ddos_mitigation:RequiredObjectAttributes:request_headers_timeout", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:slow_ddos_mitigation", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["slow_ddos_mitigation"], "schema_version": 1, "sections": [{"aliases": ["duration", "slow ddos mitigation disable request timeout"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:slow_ddos_mitigation:disable_request_timeout", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["slow_ddos_mitigation", "disable_request_timeout"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "slow ddos mitigation request headers timeout"], "anchor": "schema-slow_ddos_mitigation--request_headers_timeout", "description": "The amount of time the client has to send only the headers on the request stream before the stream is cancelled. The default value is 10000 milliseconds. This setting provides protection against Slowloris attacks.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:slow_ddos_mitigation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["slow_ddos_mitigation", "request_headers_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["duration", "slow ddos mitigation request timeout"], "anchor": "schema-slow_ddos_mitigation--request_timeout", "description": "Exclusive with", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:slow_ddos_mitigation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["slow_ddos_mitigation", "request_timeout"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/slow_ddos_mitigation/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "\"Slow and low\" attacks tie up server resources, leaving none available for servicing requests from actual users.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# slow_ddos_mitigation

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- slow_ddos_mitigation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: slow\_ddos\_mitigation, system\_default\_timeouts; Default: system\_default\_timeouts\]
'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Additional upstream details:

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

- [slow_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/slow_ddos_mitigation/#section)
- [system_default_timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/system_default_timeouts/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
slow_ddos_mitigation {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_request_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/slow_ddos_mitigation/disable_request_timeout/): complete subsection reference.

<a id="schema-slow_ddos_mitigation--request_headers_timeout"></a>

### request_headers_timeout property

Type: `"number"`. Optional.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Additional upstream details:

The default value is 10000 milliseconds.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
