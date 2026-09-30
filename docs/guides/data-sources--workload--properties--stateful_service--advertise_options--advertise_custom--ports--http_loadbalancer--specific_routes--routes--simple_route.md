---
page_title: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 6127, "body_sha256": "sha256:01531627c97432e48b388feb5905d4cdf1c0b7d9d7d933ce6235f6c2371226aa", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:simple_route", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:simple_route:auto_host_rewrite", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:simple_route:disable_host_rewrite", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:simple_route:path"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:simple_route", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes", "path": "docs/guides/data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--simple_route.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "simple_route"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/simple_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom.md)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes.md)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route

<a id="section"></a>

Type: `"single"`. Computed.

Simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

Upstream description:

A simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

## Direct properties

- [auto_host_rewrite](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--simple_route--auto_host_rewrite.md): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--simple_route--disable_host_rewrite.md): complete subsection reference.

<a id="schema-stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--simple_route--host_rewrite"></a>

### host_rewrite property

Type: `"string"`. Computed.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="schema-stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--simple_route--http_method"></a>

### http_method property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [path](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--simple_route--path.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--simple_route--auto_host_rewrite.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--simple_route--disable_host_rewrite.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--simple_route--path.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes.md)
- [xcsh_workload](../data-sources/workload.md)
