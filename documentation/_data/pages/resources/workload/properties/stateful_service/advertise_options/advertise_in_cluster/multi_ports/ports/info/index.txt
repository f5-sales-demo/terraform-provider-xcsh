---
page_title: "stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info"
subcategory: "Container"
description: "Port information."
xcsh_docs: {"aliases": ["stateful service advertise options advertise in cluster multi ports ports info"], "body_bytes": 5064, "body_sha256": "sha256:8a95fcbf1a110a7eb64cf153dad49b5a4fd426a5e85e7bb749cab606531c0084", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports:ports:info:same_as_port"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports:ports:info", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports:ports", "path": "documentation/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/multi_ports/ports/info/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-2321303312100200-2202211220202233-3112100312131322-0210100102330323-3111022100200021-1121332221111210-1312102132112232-2322011231321113", "registry_path": "docs/guides/resources--workload--reference--group-022.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports", "info"], "schema_version": 1, "sections": [{"aliases": ["stateful service advertise options advertise in cluster multi ports ports info port"], "anchor": "schema-stateful_service--advertise_options--advertise_in_cluster--multi_ports--ports--info--port", "description": "Port the workload can be reached on.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports:ports:info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports", "info", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["stateful service advertise options advertise in cluster multi ports ports info protocol"], "anchor": "schema-stateful_service--advertise_options--advertise_in_cluster--multi_ports--ports--info--protocol", "description": "Type of protocol - PROTOCOL_TCP: TCP TCP - PROTOCOL_HTTP: HTTP HTTP - PROTOCOL_HTTP2: HTTP2 HTTP2 - PROTOCOL_TLS_WITH_SNI: TLS with SNI TLS with SNI - PROTOCOL_UDP: UDP UDP.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports:ports:info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports", "info", "protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful service advertise options advertise in cluster multi ports ports info same as port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports:ports:info:same_as_port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports", "info", "same_as_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service advertise options advertise in cluster multi ports ports info target port"], "anchor": "schema-stateful_service--advertise_options--advertise_in_cluster--multi_ports--ports--info--target_port", "description": "Exclusive with Port the workload is listening on.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports:ports:info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports", "info", "target_port"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/multi_ports/ports/info/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Port information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/multi_ports/)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/multi_ports/ports/)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

Terraform syntax:

```terraform
info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-stateful_service--advertise_options--advertise_in_cluster--multi_ports--ports--info--port"></a>

### port property

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-stateful_service--advertise_options--advertise_in_cluster--multi_ports--ports--info--protocol"></a>

### protocol property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [same_as_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/multi_ports/ports/info/same_as_port/): complete subsection reference.

<a id="schema-stateful_service--advertise_options--advertise_in_cluster--multi_ports--ports--info--target_port"></a>

### target_port property

Type: `"number"`. Optional.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```
