---
page_title: "service.containers.liveness_check.http_health_check.port"
subcategory: "Container"
description: "Port"
xcsh_docs: {"aliases": ["service containers liveness check http health check port"], "body_bytes": 4081, "body_sha256": "sha256:0a01b7c9484f247773a102a5265faa3f29ed90ea86efedbd979c3a481f12fed8", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:containers:liveness_check:http_health_check:port", "parent_id": "xcsh-docs:data-sources:workload:properties:service:containers:liveness_check:http_health_check", "path": "documentation/data-sources/workload/properties/service/containers/liveness_check/http_health_check/port/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2323332311033013-0102132333310313-1220303311232030-2323132012203211-0203132131332323-3022101301331332-1333333132132203-0022111011012021", "registry_path": "docs/guides/data-sources--workload--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "containers", "liveness_check", "http_health_check", "port"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-service--containers--liveness_check--http_health_check--port--name", "description": "Exclusive with Port Name.", "document_id": "xcsh-docs:data-sources:workload:properties:service:containers:liveness_check:http_health_check:port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "containers", "liveness_check", "http_health_check", "port", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["num"], "anchor": "schema-service--containers--liveness_check--http_health_check--port--num", "description": "Exclusive with Port number.", "document_id": "xcsh-docs:data-sources:workload:properties:service:containers:liveness_check:http_health_check:port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "containers", "liveness_check", "http_health_check", "port", "num"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/containers/liveness_check/http_health_check/port/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Port", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.containers.liveness_check.http_health_check.port

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/containers/)
- [service.containers.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/containers/liveness_check/)
- [service.containers.liveness_check.http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/containers/liveness_check/http_health_check/)
- service.containers.liveness_check.http_health_check.port

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

<a id="schema-service--containers--liveness_check--http_health_check--port--name"></a>

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

<a id="schema-service--containers--liveness_check--http_health_check--port--num"></a>

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

- [service.containers.liveness_check.http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/containers/liveness_check/http_health_check/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
