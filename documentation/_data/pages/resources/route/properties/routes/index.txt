---
page_title: "routes"
subcategory: ""
description: "routes for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 12499, "body_sha256": "sha256:01d3e49fcedb05f6c792cabaac001c70a484bb9137bb8f48e60d42ba311f836c", "child_ids": ["xcsh-docs:resources:route:properties:routes:bot_defense_javascript_injection", "xcsh-docs:resources:route:properties:routes:inherited_bot_defense_javascript_injection", "xcsh-docs:resources:route:properties:routes:inherited_waf_exclusion", "xcsh-docs:resources:route:properties:routes:match", "xcsh-docs:resources:route:properties:routes:request_cookies_to_add", "xcsh-docs:resources:route:properties:routes:request_headers_to_add", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "xcsh-docs:resources:route:properties:routes:response_headers_to_add", "xcsh-docs:resources:route:properties:routes:route_destination", "xcsh-docs:resources:route:properties:routes:route_direct_response", "xcsh-docs:resources:route:properties:routes:route_redirect", "xcsh-docs:resources:route:properties:routes:service_policy", "xcsh-docs:resources:route:properties:routes:waf_exclusion_policy", "xcsh-docs:resources:route:properties:routes:waf_type"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes", "parent_id": "xcsh-docs:resources:route:reference", "path": "documentation/resources/route/properties/routes/index.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- routes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of routes to match for incoming request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bot_defense_javascript_injection",
    "inherited_bot_defense_javascript_injection"),
  validators.ConflictingListObjectAttributes("inherited_waf_exclusion",
    "waf_exclusion_policy"),
  validators.ConflictingListObjectAttributes("route_destination",
    "route_direct_response"),
  validators.ConflictingListObjectAttributes("route_destination",
    "route_redirect"),
  validators.ConflictingListObjectAttributes("route_direct_response",
    "route_redirect")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 257,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 257,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "257"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "257"
  }
}
```

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/bot_defense_javascript_injection/): complete subsection reference.

<a id="schema-routes--disable_location_add"></a>

### disable_location_add property

Type: `"bool"`. Optional.

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

- [inherited_bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/inherited_bot_defense_javascript_injection/): complete subsection reference.

- [inherited_waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/inherited_waf_exclusion/): complete subsection reference.

- [match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/): complete subsection reference.

- [request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/): complete subsection reference.

<a id="schema-routes--request_cookies_to_remove"></a>

### request_cookies_to_remove property

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

- [request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/): complete subsection reference.

<a id="schema-routes--request_headers_to_remove"></a>

### request_headers_to_remove property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
    }
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

- [response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/): complete subsection reference.

<a id="schema-routes--response_cookies_to_remove"></a>

### response_cookies_to_remove property

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

- [response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/): complete subsection reference.

<a id="schema-routes--response_headers_to_remove"></a>

### response_headers_to_remove property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
    }
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

- [route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/): complete subsection reference.

- [route_direct_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_direct_response/): complete subsection reference.

- [route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_redirect/): complete subsection reference.

- [service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/service_policy/): complete subsection reference.

- [waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_exclusion_policy/): complete subsection reference.

- [waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/): complete subsection reference.

## Next pages

- [routes.bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/bot_defense_javascript_injection/)
- [routes.inherited_bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/inherited_bot_defense_javascript_injection/)
- [routes.inherited_waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/inherited_waf_exclusion/)
- [routes.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/)
- [routes.request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/)
- [routes.request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/)
- [routes.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/)
- [routes.response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- [routes.route_direct_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_direct_response/)
- [routes.route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_redirect/)
- [routes.service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/service_policy/)
- [routes.waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_exclusion_policy/)
- [routes.waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
