---
page_title: "routes.route_destination"
subcategory: ""
description: "routes.route_destination for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 11998, "body_sha256": "sha256:814194e734ce61a3deae19e75dd824452f21c44817af183842bd288be48f3904", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:route_destination:buffer_policy", "xcsh-docs:data-sources:route:properties:routes:route_destination:cors_policy", "xcsh-docs:data-sources:route:properties:routes:route_destination:csrf_policy", "xcsh-docs:data-sources:route:properties:routes:route_destination:destinations", "xcsh-docs:data-sources:route:properties:routes:route_destination:do_not_retract_cluster", "xcsh-docs:data-sources:route:properties:routes:route_destination:endpoint_subsets", "xcsh-docs:data-sources:route:properties:routes:route_destination:hash_policy", "xcsh-docs:data-sources:route:properties:routes:route_destination:mirror_policy", "xcsh-docs:data-sources:route:properties:routes:route_destination:query_params", "xcsh-docs:data-sources:route:properties:routes:route_destination:regex_rewrite", "xcsh-docs:data-sources:route:properties:routes:route_destination:retract_cluster", "xcsh-docs:data-sources:route:properties:routes:route_destination:retry_policy", "xcsh-docs:data-sources:route:properties:routes:route_destination:spdy_config", "xcsh-docs:data-sources:route:properties:routes:route_destination:web_socket_config"], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination", "parent_id": "xcsh-docs:data-sources:route:properties:routes", "path": "documentation/data-sources/route/properties/routes/route_destination/index.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["routes", "route_destination"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.route_destination

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- routes.route_destination

<a id="section"></a>

Type: `"single"`. Computed.

List of destination to choose if the route is match.

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

## Direct properties

<a id="schema-routes--route_destination--auto_host_rewrite"></a>

### auto_host_rewrite property

Type: `"bool"`. Computed.

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

- [buffer_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/buffer_policy/): complete subsection reference.

- [cors_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/cors_policy/): complete subsection reference.

- [csrf_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/csrf_policy/): complete subsection reference.

- [destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/destinations/): complete subsection reference.

- [do_not_retract_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/do_not_retract_cluster/): complete subsection reference.

- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/endpoint_subsets/): complete subsection reference.

- [hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/): complete subsection reference.

<a id="schema-routes--route_destination--host_rewrite"></a>

### host_rewrite property

Type: `"string"`. Computed.

Exclusive with \[auto\_host\_rewrite\] Indicates that during forwarding, the host header will be
swapped with this value.

Upstream description:

Exclusive with \[auto\_host\_rewrite\] Indicates that during forwarding, the host header will be
swapped with this value.

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

- [mirror_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/mirror_policy/): complete subsection reference.

<a id="schema-routes--route_destination--prefix_rewrite"></a>

### prefix_rewrite property

Type: `"string"`. Computed.

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

Type: `"string"`. Computed.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Upstream description:

Priority routing for each request. Different connection pools are used based on the priority
selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based
on selected priority.

Default routing mechanism High-Priority routing mechanism.

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

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/query_params/): complete subsection reference.

- [regex_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/regex_rewrite/): complete subsection reference.

- [retract_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/retract_cluster/): complete subsection reference.

- [retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/retry_policy/): complete subsection reference.

- [spdy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/spdy_config/): complete subsection reference.

<a id="schema-routes--route_destination--timeout"></a>

### timeout property

Type: `"number"`. Computed.

Specifies the timeout for the route in milliseconds. This timeout includes all retries. For server
side streaming, configure this field with higher value or leave it un-configured for infinite
timeout.

Upstream description:

Specifies the timeout for the route in milliseconds. This timeout includes all retries. For server
side streaming, configure this field with higher value or leave it un-configured for infinite
timeout.

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

- [web_socket_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/web_socket_config/): complete subsection reference.

## Next pages

- [routes.route_destination.buffer_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/buffer_policy/)
- [routes.route_destination.cors_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/cors_policy/)
- [routes.route_destination.csrf_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/csrf_policy/)
- [routes.route_destination.destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/destinations/)
- [routes.route_destination.do_not_retract_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/do_not_retract_cluster/)
- [routes.route_destination.endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/endpoint_subsets/)
- [routes.route_destination.hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/hash_policy/)
- [routes.route_destination.mirror_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/mirror_policy/)
- [routes.route_destination.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/query_params/)
- [routes.route_destination.regex_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/regex_rewrite/)
- [routes.route_destination.retract_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/retract_cluster/)
- [routes.route_destination.retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/retry_policy/)
- [routes.route_destination.spdy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/spdy_config/)
- [routes.route_destination.web_socket_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/web_socket_config/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
