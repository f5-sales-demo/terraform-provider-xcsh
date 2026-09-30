---
page_title: "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 5969, "body_sha256": "sha256:023403335340335c25083c6eff0291918835d160d9576e4d7468ae18a059a117", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:direct_response_route:incoming_port", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:direct_response_route:incoming_port:no_port_match"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:direct_response_route:incoming_port", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:direct_response_route", "path": "docs/guides/data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--specific_routes--routes--direct_response_route--incoming_port.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "specific_routes", "routes", "direct_response_route", "incoming_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/direct_response_route/incoming_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--specific_routes.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--specific_routes--routes.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--specific_routes--routes--direct_response_route.md)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port

<a id="section"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

## Direct properties

- [no_port_match](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--specific_routes--routes--direct_response_route--incoming_port--no_port_match.md): complete subsection reference.

<a id="schema-stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--specific_routes--routes--direct_response_route--incoming_port--port"></a>

### port property

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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

<a id="schema-stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--specific_routes--routes--direct_response_route--incoming_port--port_ranges"></a>

### port_ranges property

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

## Next pages

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--specific_routes--routes--direct_response_route--incoming_port--no_port_match.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--specific_routes--routes--direct_response_route.md)
- [xcsh_workload](../data-sources/workload.md)
