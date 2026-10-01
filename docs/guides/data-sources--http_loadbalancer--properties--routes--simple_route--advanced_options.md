---
page_title: "routes.simple_route.advanced_options"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 21711, "body_sha256": "sha256:a273b73fa403ff9525e2fc4da50639603a211929ba82d53e51fca82ebb9e7817", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:app_firewall", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:bot_defense_javascript_injection", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:buffer_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:common_buffering", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:common_hash_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:cors_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:default_retry_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_mirroring", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_prefix_rewrite", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_spdy", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_waf", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_web_socket_config", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:do_not_retract_cluster", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:enable_spdy", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:endpoint_subsets", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_bot_defense_javascript_injection", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_waf", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_waf_exclusion", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:no_retry_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:regex_rewrite", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:request_cookies_to_add", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:request_headers_to_add", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_headers_to_add", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:retract_cluster", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:retry_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:waf_exclusion_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:web_socket_config"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route", "path": "docs/guides/data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [routes](data-sources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](data-sources--http_loadbalancer--properties--routes--simple_route.md)
- routes.simple_route.advanced_options

<a id="section"></a>

Type: `"single"`. Computed.

Configure advanced OPTIONS for route like path rewrite, hash policy, etc.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-bot_defense_javascript_injection_choice": "[\"bot_defense_javascript_injection\",\"inherited_bot_defense_javascript_injection\"]",
  "x-ves-oneof-field-buffer_choice": "[\"buffer_policy\",\"common_buffering\"]",
  "x-ves-oneof-field-cluster_retract_choice": "[\"do_not_retract_cluster\",\"retract_cluster\"]",
  "x-ves-oneof-field-hash_policy_choice": "[\"common_hash_policy\",\"specific_hash_policy\"]",
  "x-ves-oneof-field-mirroring_choice": "[\"disable_mirroring\",\"mirror_policy\"]",
  "x-ves-oneof-field-retry_policy_choice": "[\"default_retry_policy\",\"no_retry_policy\",\"retry_policy\"]",
  "x-ves-oneof-field-rewrite_choice": "[\"disable_prefix_rewrite\",\"prefix_rewrite\",\"regex_rewrite\"]",
  "x-ves-oneof-field-spdy_choice": "[\"disable_spdy\",\"enable_spdy\"]",
  "x-ves-oneof-field-waf_choice": "[\"app_firewall\",\"disable_waf\",\"inherited_waf\"]",
  "x-ves-oneof-field-waf_exclusion_choice": "[\"inherited_waf_exclusion\",\"waf_exclusion_policy\"]",
  "x-ves-oneof-field-websocket_choice": "[\"disable_web_socket_config\",\"web_socket_config\"]"
}
```

## Direct properties

- [app_firewall](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--app_firewall.md): complete subsection reference.

- [bot_defense_javascript_injection](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--bot_defense_javascript_injection.md): complete subsection reference.

- [buffer_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--buffer_policy.md): complete subsection reference.

- [common_buffering](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--common_buffering.md): complete subsection reference.

- [common_hash_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--common_hash_policy.md): complete subsection reference.

- [cors_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--cors_policy.md): complete subsection reference.

- [csrf_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--csrf_policy.md): complete subsection reference.

- [default_retry_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--default_retry_policy.md): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--disable_location_add"></a>

### disable_location_add property

Type: `"bool"`. Computed.

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

Upstream description:

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

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

- [disable_mirroring](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--disable_mirroring.md): complete subsection reference.

- [disable_prefix_rewrite](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--disable_prefix_rewrite.md): complete subsection reference.

- [disable_spdy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--disable_spdy.md): complete subsection reference.

- [disable_waf](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--disable_waf.md): complete subsection reference.

- [disable_web_socket_config](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--disable_web_socket_config.md): complete subsection reference.

- [do_not_retract_cluster](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--do_not_retract_cluster.md): complete subsection reference.

- [enable_spdy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--enable_spdy.md): complete subsection reference.

- [endpoint_subsets](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--endpoint_subsets.md): complete subsection reference.

- [inherited_bot_defense_javascript_injection](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--inherited_bot_defense_javascript_injection.md): complete subsection reference.

- [inherited_waf](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--inherited_waf.md): complete subsection reference.

- [inherited_waf_exclusion](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--inherited_waf_exclusion.md): complete subsection reference.

- [mirror_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--mirror_policy.md): complete subsection reference.

- [no_retry_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--no_retry_policy.md): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--prefix_rewrite"></a>

### prefix_rewrite property

Type: `"string"`. Computed.

Exclusive with \[disable\_prefix\_rewrite regex\_rewrite\] prefix\_rewrite indicates that during
forwarding, the matched prefix (or path) should be swapped with its value. When using regex path
matching, the entire path (not including the query string) will be swapped with this value.

Upstream description:

Exclusive with \[disable\_prefix\_rewrite regex\_rewrite\] prefix\_rewrite indicates that during
forwarding, the matched prefix (or path) should be swapped with its value. When using regex path
matching, the entire path (not including the query string) will be swapped with this value.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-routes--simple_route--advanced_options--priority"></a>

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

- [regex_rewrite](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--regex_rewrite.md): complete subsection reference.

- [request_cookies_to_add](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--request_cookies_to_add.md): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--request_cookies_to_remove"></a>

### request_cookies_to_remove property

Type: `["list", "string"]`. Computed.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--request_headers_to_add.md): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--request_headers_to_remove"></a>

### request_headers_to_remove property

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add.md): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--response_cookies_to_remove"></a>

### response_cookies_to_remove property

Type: `["list", "string"]`. Computed.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_headers_to_add.md): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--response_headers_to_remove"></a>

### response_headers_to_remove property

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [retract_cluster](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--retract_cluster.md): complete subsection reference.

- [retry_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--retry_policy.md): complete subsection reference.

- [specific_hash_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy.md): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--timeout"></a>

### timeout property

Type: `"number"`. Computed.

The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0
(infinite timeout) for server-side streaming.

Upstream description:

The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0
(infinite timeout) for server-side streaming.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [waf_exclusion_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--waf_exclusion_policy.md): complete subsection reference.

- [web_socket_config](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--web_socket_config.md): complete subsection reference.

## Next pages

- [routes.simple_route.advanced_options.app_firewall](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--app_firewall.md)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--bot_defense_javascript_injection.md)
- [routes.simple_route.advanced_options.buffer_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--buffer_policy.md)
- [routes.simple_route.advanced_options.common_buffering](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--common_buffering.md)
- [routes.simple_route.advanced_options.common_hash_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--common_hash_policy.md)
- [routes.simple_route.advanced_options.cors_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--cors_policy.md)
- [routes.simple_route.advanced_options.csrf_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--csrf_policy.md)
- [routes.simple_route.advanced_options.default_retry_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--default_retry_policy.md)
- [routes.simple_route.advanced_options.disable_mirroring](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--disable_mirroring.md)
- [routes.simple_route.advanced_options.disable_prefix_rewrite](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--disable_prefix_rewrite.md)
- [routes.simple_route.advanced_options.disable_spdy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--disable_spdy.md)
- [routes.simple_route.advanced_options.disable_waf](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--disable_waf.md)
- [routes.simple_route.advanced_options.disable_web_socket_config](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--disable_web_socket_config.md)
- [routes.simple_route.advanced_options.do_not_retract_cluster](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--do_not_retract_cluster.md)
- [routes.simple_route.advanced_options.enable_spdy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--enable_spdy.md)
- [routes.simple_route.advanced_options.endpoint_subsets](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--endpoint_subsets.md)
- [routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--inherited_bot_defense_javascript_injection.md)
- [routes.simple_route.advanced_options.inherited_waf](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--inherited_waf.md)
- [routes.simple_route.advanced_options.inherited_waf_exclusion](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--inherited_waf_exclusion.md)
- [routes.simple_route.advanced_options.mirror_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--mirror_policy.md)
- [routes.simple_route.advanced_options.no_retry_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--no_retry_policy.md)
- [routes.simple_route.advanced_options.regex_rewrite](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--regex_rewrite.md)
- [routes.simple_route.advanced_options.request_cookies_to_add](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--request_cookies_to_add.md)
- [routes.simple_route.advanced_options.request_headers_to_add](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--request_headers_to_add.md)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add.md)
- [routes.simple_route.advanced_options.response_headers_to_add](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_headers_to_add.md)
- [routes.simple_route.advanced_options.retract_cluster](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--retract_cluster.md)
- [routes.simple_route.advanced_options.retry_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--retry_policy.md)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--specific_hash_policy.md)
- [routes.simple_route.advanced_options.waf_exclusion_policy](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--waf_exclusion_policy.md)
- [routes.simple_route.advanced_options.web_socket_config](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--web_socket_config.md)
- [routes.simple_route](data-sources--http_loadbalancer--properties--routes--simple_route.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
