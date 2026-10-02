---
page_title: "stateful_service.advertise_options.advertise_custom.ports.port.info"
subcategory: "Container"
description: "Port information."
xcsh_docs: {"aliases": ["stateful service advertise options advertise custom ports port info"], "body_bytes": 5893, "body_sha256": "sha256:35fcc4d3879f974d8ab01853f47e709a918a88949f905fae760789cea93f4300", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:port:info:same_as_port"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:port:info", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:port", "path": "documentation/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/port/info/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2110220130101010-3030301022012202-0021201311302013-2010021213133010-3003201320101102-3211032223013120-2011131020111320-2311023211121203", "registry_path": "docs/guides/data-sources--workload--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "port", "info"], "schema_version": 1, "sections": [{"aliases": ["port"], "anchor": "schema-stateful_service--advertise_options--advertise_custom--ports--port--info--port", "description": "Port the workload can be reached on.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:port:info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "port", "info", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["protocol"], "anchor": "schema-stateful_service--advertise_options--advertise_custom--ports--port--info--protocol", "description": "Type of protocol - PROTOCOL_TCP: TCP TCP - PROTOCOL_HTTP: HTTP HTTP - PROTOCOL_HTTP2: HTTP2 HTTP2 - PROTOCOL_TLS_WITH_SNI: TLS with SNI TLS with SNI - PROTOCOL_UDP: UDP UDP.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:port:info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "port", "info", "protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["same as port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:port:info:same_as_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "port", "info", "same_as_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["target port"], "anchor": "schema-stateful_service--advertise_options--advertise_custom--ports--port--info--target_port", "description": "Exclusive with Port the workload is listening on.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:port:info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "port", "info", "target_port"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/port/info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Port information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_custom.ports.port.info

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/)
- [stateful_service.advertise_options.advertise_custom.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/)
- [stateful_service.advertise_options.advertise_custom.ports.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/port/)
- stateful_service.advertise_options.advertise_custom.ports.port.info

<a id="section"></a>

Type: `"single"`. Computed.

Port Information. Port information.

Upstream description:

Port information.

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

<a id="schema-stateful_service--advertise_options--advertise_custom--ports--port--info--port"></a>

### port property

Type: `"number"`. Computed.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

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

<a id="schema-stateful_service--advertise_options--advertise_custom--ports--port--info--protocol"></a>

### protocol property

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

TCP &#8203;- PROTOCOL\_HTTP: HTTP

HTTP &#8203;- PROTOCOL\_HTTP2: HTTP2

HTTP2 &#8203;- PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI

TLS with SNI &#8203;- PROTOCOL\_UDP: UDP

UDP.

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

- [same_as_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/port/info/same_as_port/): complete subsection reference.

<a id="schema-stateful_service--advertise_options--advertise_custom--ports--port--info--target_port"></a>

### target_port property

Type: `"number"`. Computed.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

## Next pages

- [stateful_service.advertise_options.advertise_custom.ports.port.info.same_as_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/port/info/same_as_port/)
- [stateful_service.advertise_options.advertise_custom.ports.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/port/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
