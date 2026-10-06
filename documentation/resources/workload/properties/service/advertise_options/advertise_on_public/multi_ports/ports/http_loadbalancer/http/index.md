---
page_title: "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http"
subcategory: "Container"
description: "Choice for selecting HTTP proxy."
xcsh_docs: {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer http"], "body_bytes": 5545, "body_sha256": "sha256:b507ad29643eba34b8678b314837f387952435475c716c2329939ca154ae5311", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:http", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/http/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2032123312300020-0313233113231322-1312210332031202-2133210330020310-0110031022301023-1212321320020231-2033231203003230-2310210201131220", "registry_path": "docs/guides/resources--workload--reference--group-008.md", "relationships": [{"anchor": "schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--http--port", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:http", "type": "conflicts"}, {"anchor": "schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--http--port_ranges", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:http", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "http"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer http dns volterra managed"], "anchor": "schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--http--dns_volterra_managed", "description": "DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite, the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME record should be created in your DNS provider's portal.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:http", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "http", "dns_volterra_managed"], "syntax": "attribute", "type": "bool"}, {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer http port"], "anchor": "schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--http--port", "description": "Exclusive with HTTP port to Listen.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:http", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "http", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["service advertise options advertise on public multi ports ports http loadbalancer http port ranges"], "anchor": "schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--http--port_ranges", "description": "Exclusive with A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by \"-\".", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:http", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "http", "port_ranges"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/http/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Choice for selecting HTTP proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/)
- [service.advertise_options.advertise_on_public.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/)
- [service.advertise_options.advertise_on_public.multi_ports.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP Choice. Choice for selecting HTTP proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
http {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--http--dns_volterra_managed"></a>

### dns_volterra_managed property

Type: `"bool"`. Optional.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

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

<a id="schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--http--port"></a>

### port property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--http--port_ranges"></a>

### port_ranges property

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```
