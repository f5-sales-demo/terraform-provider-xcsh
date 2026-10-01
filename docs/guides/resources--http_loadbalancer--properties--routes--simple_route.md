---
page_title: "routes.simple_route"
subcategory: "Load Balancing"
description: "routes.simple_route for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 6628, "body_sha256": "sha256:c182e8af08fe6a5d1fd03f86f68f958336cd06b302fca882d86ae9f671989e10", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:auto_host_rewrite", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:caching_disable", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:caching_inherit", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:disable_host_rewrite", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:headers", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:origin_pools", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:path", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:query_params"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes", "path": "docs/guides/resources--http_loadbalancer--properties--routes--simple_route.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- routes.simple_route

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the
matching traffic to the associated pools.

Upstream description:

A simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the
matching traffic to the associated pools.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_pools"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("caching_disable",
    "caching_inherit"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
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
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]",
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
simple_route {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advanced_options](resources--http_loadbalancer--properties--routes--simple_route--advanced_options.md): complete subsection reference.

- [auto_host_rewrite](resources--http_loadbalancer--properties--routes--simple_route--auto_host_rewrite.md): complete subsection reference.

- [caching_disable](resources--http_loadbalancer--properties--routes--simple_route--caching_disable.md): complete subsection reference.

- [caching_inherit](resources--http_loadbalancer--properties--routes--simple_route--caching_inherit.md): complete subsection reference.

- [disable_host_rewrite](resources--http_loadbalancer--properties--routes--simple_route--disable_host_rewrite.md): complete subsection reference.

- [headers](resources--http_loadbalancer--properties--routes--simple_route--headers.md): complete subsection reference.

<a id="schema-routes--simple_route--host_rewrite"></a>

### host_rewrite property

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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

<a id="schema-routes--simple_route--http_method"></a>

### http_method property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

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

- [incoming_port](resources--http_loadbalancer--properties--routes--simple_route--incoming_port.md): complete subsection reference.

- [origin_pools](resources--http_loadbalancer--properties--routes--simple_route--origin_pools.md): complete subsection reference.

- [path](resources--http_loadbalancer--properties--routes--simple_route--path.md): complete subsection reference.

- [query_params](resources--http_loadbalancer--properties--routes--simple_route--query_params.md): complete subsection reference.

## Next pages

- [routes.simple_route.advanced_options](resources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- [routes.simple_route.auto_host_rewrite](resources--http_loadbalancer--properties--routes--simple_route--auto_host_rewrite.md)
- [routes.simple_route.caching_disable](resources--http_loadbalancer--properties--routes--simple_route--caching_disable.md)
- [routes.simple_route.caching_inherit](resources--http_loadbalancer--properties--routes--simple_route--caching_inherit.md)
- [routes.simple_route.disable_host_rewrite](resources--http_loadbalancer--properties--routes--simple_route--disable_host_rewrite.md)
- [routes.simple_route.headers](resources--http_loadbalancer--properties--routes--simple_route--headers.md)
- [routes.simple_route.incoming_port](resources--http_loadbalancer--properties--routes--simple_route--incoming_port.md)
- [routes.simple_route.origin_pools](resources--http_loadbalancer--properties--routes--simple_route--origin_pools.md)
- [routes.simple_route.path](resources--http_loadbalancer--properties--routes--simple_route--path.md)
- [routes.simple_route.query_params](resources--http_loadbalancer--properties--routes--simple_route--query_params.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
