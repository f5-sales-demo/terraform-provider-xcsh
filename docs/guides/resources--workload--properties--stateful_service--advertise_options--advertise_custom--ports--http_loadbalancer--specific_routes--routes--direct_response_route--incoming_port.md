---
page_title: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 6158, "body_sha256": "sha256:2b59ec7fb41d9db2d2941bf2131a3531c5856cff0ccb5e402965e89737c881ef", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route:incoming_port", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route:incoming_port:no_port_match"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route:incoming_port", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--incoming_port.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "direct_response_route", "incoming_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/direct_response_route/incoming_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_custom](resources--workload--properties--stateful_service--advertise_options--advertise_custom.md)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route.md)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
  validators.ConflictingObjectAttributes("port",
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
incoming_port {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_port_match](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--incoming_port--no_port_match.md): complete subsection reference.

<a id="schema-stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--incoming_port--port"></a>

### port property

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Provider validators and defaults (from schema source):

```go
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

<a id="schema-stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--incoming_port--port_ranges"></a>

### port_ranges property

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

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

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--incoming_port--no_port_match.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route.md)
- [xcsh_workload](../resources/workload.md)
