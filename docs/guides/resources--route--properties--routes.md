---
page_title: "routes"
subcategory: ""
description: "routes for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 10891, "body_sha256": "sha256:3a8ed57250413ca7e454a024a6ecba852384640c482a3b4bd599daef65d63f34", "canonical_id": "xcsh-docs:resources:route:properties:routes", "child_ids": ["xcsh-docs:resources:route:properties:routes:bot_defense_javascript_injection", "xcsh-docs:resources:route:properties:routes:inherited_bot_defense_javascript_injection", "xcsh-docs:resources:route:properties:routes:inherited_waf_exclusion", "xcsh-docs:resources:route:properties:routes:match", "xcsh-docs:resources:route:properties:routes:request_cookies_to_add", "xcsh-docs:resources:route:properties:routes:request_headers_to_add", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "xcsh-docs:resources:route:properties:routes:response_headers_to_add", "xcsh-docs:resources:route:properties:routes:route_destination", "xcsh-docs:resources:route:properties:routes:route_direct_response", "xcsh-docs:resources:route:properties:routes:route_redirect", "xcsh-docs:resources:route:properties:routes:service_policy", "xcsh-docs:resources:route:properties:routes:waf_exclusion_policy", "xcsh-docs:resources:route:properties:routes:waf_type"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes", "parent_id": "xcsh-docs:resources:route:reference", "path": "docs/guides/resources--route--properties--routes.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
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

- [bot_defense_javascript_injection](resources--route--properties--routes--bot_defense_javascript_injection.md): complete subsection reference.

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

- [inherited_bot_defense_javascript_injection](resources--route--properties--routes--inherited_bot_defense_javascript_injection.md): complete subsection reference.

- [inherited_waf_exclusion](resources--route--properties--routes--inherited_waf_exclusion.md): complete subsection reference.

- [match](resources--route--properties--routes--match.md): complete subsection reference.

- [request_cookies_to_add](resources--route--properties--routes--request_cookies_to_add.md): complete subsection reference.

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

- [request_headers_to_add](resources--route--properties--routes--request_headers_to_add.md): complete subsection reference.

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

- [response_cookies_to_add](resources--route--properties--routes--response_cookies_to_add.md): complete subsection reference.

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

- [response_headers_to_add](resources--route--properties--routes--response_headers_to_add.md): complete subsection reference.

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

- [route_destination](resources--route--properties--routes--route_destination.md): complete subsection reference.

- [route_direct_response](resources--route--properties--routes--route_direct_response.md): complete subsection reference.

- [route_redirect](resources--route--properties--routes--route_redirect.md): complete subsection reference.

- [service_policy](resources--route--properties--routes--service_policy.md): complete subsection reference.

- [waf_exclusion_policy](resources--route--properties--routes--waf_exclusion_policy.md): complete subsection reference.

- [waf_type](resources--route--properties--routes--waf_type.md): complete subsection reference.

## Next pages

- [routes.bot_defense_javascript_injection](resources--route--properties--routes--bot_defense_javascript_injection.md)
- [routes.inherited_bot_defense_javascript_injection](resources--route--properties--routes--inherited_bot_defense_javascript_injection.md)
- [routes.inherited_waf_exclusion](resources--route--properties--routes--inherited_waf_exclusion.md)
- [routes.match](resources--route--properties--routes--match.md)
- [routes.request_cookies_to_add](resources--route--properties--routes--request_cookies_to_add.md)
- [routes.request_headers_to_add](resources--route--properties--routes--request_headers_to_add.md)
- [routes.response_cookies_to_add](resources--route--properties--routes--response_cookies_to_add.md)
- [routes.response_headers_to_add](resources--route--properties--routes--response_headers_to_add.md)
- [routes.route_destination](resources--route--properties--routes--route_destination.md)
- [routes.route_direct_response](resources--route--properties--routes--route_direct_response.md)
- [routes.route_redirect](resources--route--properties--routes--route_redirect.md)
- [routes.service_policy](resources--route--properties--routes--service_policy.md)
- [routes.waf_exclusion_policy](resources--route--properties--routes--waf_exclusion_policy.md)
- [routes.waf_type](resources--route--properties--routes--waf_type.md)
- [Property reference](resources--route--reference.md)
- [xcsh_route](../resources/route.md)
