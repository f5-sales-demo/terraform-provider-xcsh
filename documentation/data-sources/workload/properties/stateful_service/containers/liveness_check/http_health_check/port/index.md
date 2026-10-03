---
page_title: "stateful_service.containers.liveness_check.http_health_check.port"
subcategory: "Container"
description: "Port"
xcsh_docs: {"aliases": ["stateful service containers liveness check http health check port"], "body_bytes": 4207, "body_sha256": "sha256:c1582a7b5b04fb21142c9692154e7053d597f152b3500fac1543ea1366c8201f", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers:liveness_check:http_health_check:port", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers:liveness_check:http_health_check", "path": "documentation/data-sources/workload/properties/stateful_service/containers/liveness_check/http_health_check/port/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1310032210323030-3033100031312022-2131302223001113-1332031022200130-3333002311103212-1122010200122032-1101020331132220-3102131320020331", "registry_path": "docs/guides/data-sources--workload--reference--group-028.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "containers", "liveness_check", "http_health_check", "port"], "schema_version": 1, "sections": [{"aliases": ["stateful service containers liveness check http health check port name"], "anchor": "schema-stateful_service--containers--liveness_check--http_health_check--port--name", "description": "Exclusive with Port Name.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers:liveness_check:http_health_check:port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "liveness_check", "http_health_check", "port", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful service containers liveness check http health check port num"], "anchor": "schema-stateful_service--containers--liveness_check--http_health_check--port--num", "description": "Exclusive with Port number.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers:liveness_check:http_health_check:port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "liveness_check", "http_health_check", "port", "num"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/containers/liveness_check/http_health_check/port/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Port", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.containers.liveness_check.http_health_check.port

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [stateful_service.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/)
- [stateful_service.containers.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/liveness_check/)
- [stateful_service.containers.liveness_check.http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/liveness_check/http_health_check/)
- stateful_service.containers.liveness_check.http_health_check.port

<a id="section"></a>

Type: `"single"`. Computed.

Port. Port

Upstream description:

Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

## Direct properties

<a id="schema-stateful_service--containers--liveness_check--http_health_check--port--name"></a>

### name property

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="schema-stateful_service--containers--liveness_check--http_health_check--port--num"></a>

### num property

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

## Next pages

- [stateful_service.containers.liveness_check.http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/liveness_check/http_health_check/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
