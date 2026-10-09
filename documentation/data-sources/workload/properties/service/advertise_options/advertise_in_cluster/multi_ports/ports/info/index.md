---
page_title: "service.advertise_options.advertise_in_cluster.multi_ports.ports.info"
subcategory: "Container"
description: "Port information."
xcsh_docs: {"aliases": ["service advertise options advertise in cluster multi ports ports info"], "body_bytes": 4831, "body_sha256": "sha256:dc86c5b30d01462aef8b9bf77a715a61f56e1b884124eda067ee722cd39e517c", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports:info:same_as_port"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports:info", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports", "path": "documentation/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/ports/info/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1311013020110322-3332312222230230-2111331120112121-1302023321211110-3121210102132002-3331300220020203-0032132323002103-2013103230232311", "registry_path": "docs/guides/data-sources--workload--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports", "info"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise in cluster multi ports ports info port"], "anchor": "schema-service--advertise_options--advertise_in_cluster--multi_ports--ports--info--port", "description": "Port the workload can be reached on.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports:info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports", "info", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["service advertise options advertise in cluster multi ports ports info protocol"], "anchor": "schema-service--advertise_options--advertise_in_cluster--multi_ports--ports--info--protocol", "description": "Type of protocol - PROTOCOL_TCP: TCP TCP - PROTOCOL_HTTP: HTTP HTTP - PROTOCOL_HTTP2: HTTP2 HTTP2 - PROTOCOL_TLS_WITH_SNI: TLS with SNI TLS with SNI - PROTOCOL_UDP: UDP UDP.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports:info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports", "info", "protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["service advertise options advertise in cluster multi ports ports info same as port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports:info:same_as_port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports", "info", "same_as_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise in cluster multi ports ports info target port"], "anchor": "schema-service--advertise_options--advertise_in_cluster--multi_ports--ports--info--target_port", "description": "Exclusive with Port the workload is listening on.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports:info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports", "info", "target_port"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/ports/info/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Port information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["workloadCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_in_cluster.multi_ports.ports.info

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/)
- [service.advertise_options.advertise_in_cluster.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/ports/)
- service.advertise_options.advertise_in_cluster.multi_ports.ports.info

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-service--advertise_options--advertise_in_cluster--multi_ports--ports--info--port"></a>

### port property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="schema-service--advertise_options--advertise_in_cluster--multi_ports--ports--info--protocol"></a>

### protocol property

Type: `"string"`. Computed.

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

- [same_as_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/ports/info/same_as_port/): complete subsection reference.

<a id="schema-service--advertise_options--advertise_in_cluster--multi_ports--ports--info--target_port"></a>

### target_port property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
