---
page_title: "routes.route_destination"
subcategory: ""
description: "routes.route_destination for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 11369, "body_sha256": "sha256:5aec5c4f7fe42faf2731f4a2244115685049f25329ad7364a18af5a843ca9f02", "canonical_id": "xcsh-docs:resources:route:properties:routes:route_destination", "child_ids": ["xcsh-docs:resources:route:properties:routes:route_destination:buffer_policy", "xcsh-docs:resources:route:properties:routes:route_destination:cors_policy", "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy", "xcsh-docs:resources:route:properties:routes:route_destination:destinations", "xcsh-docs:resources:route:properties:routes:route_destination:do_not_retract_cluster", "xcsh-docs:resources:route:properties:routes:route_destination:endpoint_subsets", "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy", "xcsh-docs:resources:route:properties:routes:route_destination:query_params", "xcsh-docs:resources:route:properties:routes:route_destination:regex_rewrite", "xcsh-docs:resources:route:properties:routes:route_destination:retract_cluster", "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy", "xcsh-docs:resources:route:properties:routes:route_destination:spdy_config", "xcsh-docs:resources:route:properties:routes:route_destination:web_socket_config"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination", "parent_id": "xcsh-docs:resources:route:properties:routes", "path": "docs/guides/resources--route--properties--routes--route_destination.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "route_destination"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.route_destination

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- routes.route_destination

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of destination to choose if the route is match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("destinations"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("do_not_retract_cluster",
    "retract_cluster"),
  validators.ConflictingObjectAttributes("prefix_rewrite",
    "regex_rewrite")}
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
  "x-ves-oneof-field-cluster_retract_choice": "[\"do_not_retract_cluster\",\"retract_cluster\"]",
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"host_rewrite\"]",
  "x-ves-oneof-field-route_destination_rewrite": "[\"prefix_rewrite\",\"regex_rewrite\"]"
}
```

Terraform syntax:

```terraform
route_destination {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--route_destination--auto_host_rewrite"></a>

### auto_host_rewrite property

Type: `"bool"`. Optional.

Exclusive with \[host\_rewrite\] Indicates that during forwarding, the host header will be swapped
with the hostname of the upstream host chosen by the cluster.

Upstream description:

Exclusive with \[host\_rewrite\] Indicates that during forwarding, the host header will be swapped
with the hostname of the upstream host chosen by the cluster.

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

- [buffer_policy](resources--route--properties--routes--route_destination--buffer_policy.md): complete subsection reference.

- [cors_policy](resources--route--properties--routes--route_destination--cors_policy.md): complete subsection reference.

- [csrf_policy](resources--route--properties--routes--route_destination--csrf_policy.md): complete subsection reference.

- [destinations](resources--route--properties--routes--route_destination--destinations.md): complete subsection reference.

- [do_not_retract_cluster](resources--route--properties--routes--route_destination--do_not_retract_cluster.md): complete subsection reference.

- [endpoint_subsets](resources--route--properties--routes--route_destination--endpoint_subsets.md): complete subsection reference.

- [hash_policy](resources--route--properties--routes--route_destination--hash_policy.md): complete subsection reference.

<a id="schema-routes--route_destination--host_rewrite"></a>

### host_rewrite property

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite\] Indicates that during forwarding, the host header will be
swapped with this value.

Upstream description:

Exclusive with \[auto\_host\_rewrite\] Indicates that during forwarding, the host header will be
swapped with this value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [mirror_policy](resources--route--properties--routes--route_destination--mirror_policy.md): complete subsection reference.

<a id="schema-routes--route_destination--prefix_rewrite"></a>

### prefix_rewrite property

Type: `"string"`. Optional.

Exclusive with \[regex\_rewrite\] prefix\_rewrite indicates that during forwarding, the matched
prefix (or path) should be swapped with its value. When using regex path matching, the entire path
(not including the query string) will be swapped with this value. This option allows application
URLs to..

Upstream description:

Exclusive with \[regex\_rewrite\] prefix\_rewrite indicates that during forwarding, the matched
prefix (or path) should be swapped with its value. When using regex path matching, the entire path
(not including the query string) will be swapped with this value. This option allows application
URLs to be rooted at a different path from those exposed at the reverse proxy layer.

Example : gcSpec: routes: &#8203;- match: &#8203;- headers: \[\] path: prefix : /register/
query\_params: \[\] &#8203;- headers: \[\] path: prefix: /register query\_params: \[\]
routeDestination: prefixRewrite: "/" destinations: &#8203;- cluster: &#8203;- kind: cluster.object
uid: cluster-1

Having above entries in the config, requests to /register will be stripped to /, while requests to
/register/public will be stripped to /public.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-routes--route_destination--priority"></a>

### priority property

Type: `"string"`. Optional.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Upstream description:

Priority routing for each request. Different connection pools are used based on the priority
selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based
on selected priority.

Default routing mechanism High-Priority routing mechanism.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DEFAULT",
    "HIGH"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DEFAULT",
  "enum": [
    "DEFAULT",
    "HIGH"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [query_params](resources--route--properties--routes--route_destination--query_params.md): complete subsection reference.

- [regex_rewrite](resources--route--properties--routes--route_destination--regex_rewrite.md): complete subsection reference.

- [retract_cluster](resources--route--properties--routes--route_destination--retract_cluster.md): complete subsection reference.

- [retry_policy](resources--route--properties--routes--route_destination--retry_policy.md): complete subsection reference.

- [spdy_config](resources--route--properties--routes--route_destination--spdy_config.md): complete subsection reference.

<a id="schema-routes--route_destination--timeout"></a>

### timeout property

Type: `"number"`. Optional.

Specifies the timeout for the route in milliseconds. This timeout includes all retries. For server
side streaming, configure this field with higher value or leave it un-configured for infinite
timeout.

Upstream description:

Specifies the timeout for the route in milliseconds. This timeout includes all retries. For server
side streaming, configure this field with higher value or leave it un-configured for infinite
timeout.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1800000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1800000,
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
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

- [web_socket_config](resources--route--properties--routes--route_destination--web_socket_config.md): complete subsection reference.

## Next pages

- [routes.route_destination.buffer_policy](resources--route--properties--routes--route_destination--buffer_policy.md)
- [routes.route_destination.cors_policy](resources--route--properties--routes--route_destination--cors_policy.md)
- [routes.route_destination.csrf_policy](resources--route--properties--routes--route_destination--csrf_policy.md)
- [routes.route_destination.destinations](resources--route--properties--routes--route_destination--destinations.md)
- [routes.route_destination.do_not_retract_cluster](resources--route--properties--routes--route_destination--do_not_retract_cluster.md)
- [routes.route_destination.endpoint_subsets](resources--route--properties--routes--route_destination--endpoint_subsets.md)
- [routes.route_destination.hash_policy](resources--route--properties--routes--route_destination--hash_policy.md)
- [routes.route_destination.mirror_policy](resources--route--properties--routes--route_destination--mirror_policy.md)
- [routes.route_destination.query_params](resources--route--properties--routes--route_destination--query_params.md)
- [routes.route_destination.regex_rewrite](resources--route--properties--routes--route_destination--regex_rewrite.md)
- [routes.route_destination.retract_cluster](resources--route--properties--routes--route_destination--retract_cluster.md)
- [routes.route_destination.retry_policy](resources--route--properties--routes--route_destination--retry_policy.md)
- [routes.route_destination.spdy_config](resources--route--properties--routes--route_destination--spdy_config.md)
- [routes.route_destination.web_socket_config](resources--route--properties--routes--route_destination--web_socket_config.md)
- [routes](resources--route--properties--routes.md)
- [xcsh_route](../resources/route.md)
