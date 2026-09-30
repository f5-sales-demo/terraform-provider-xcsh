---
page_title: "routes"
subcategory: ""
description: "routes for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 9644, "body_sha256": "sha256:2ff7f3f737d5390ffa5adc37de0b776208e7c647db201dc98a1e1a3947b50601", "canonical_id": "xcsh-docs:data-sources:route:properties:routes", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection", "xcsh-docs:data-sources:route:properties:routes:inherited_bot_defense_javascript_injection", "xcsh-docs:data-sources:route:properties:routes:inherited_waf_exclusion", "xcsh-docs:data-sources:route:properties:routes:match", "xcsh-docs:data-sources:route:properties:routes:request_cookies_to_add", "xcsh-docs:data-sources:route:properties:routes:request_headers_to_add", "xcsh-docs:data-sources:route:properties:routes:response_cookies_to_add", "xcsh-docs:data-sources:route:properties:routes:response_headers_to_add", "xcsh-docs:data-sources:route:properties:routes:route_destination", "xcsh-docs:data-sources:route:properties:routes:route_direct_response", "xcsh-docs:data-sources:route:properties:routes:route_redirect", "xcsh-docs:data-sources:route:properties:routes:service_policy", "xcsh-docs:data-sources:route:properties:routes:waf_exclusion_policy", "xcsh-docs:data-sources:route:properties:routes:waf_type"], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes", "parent_id": "xcsh-docs:data-sources:route:reference", "path": "docs/guides/data-sources--route--properties--routes.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes

Breadcrumbs:

- [xcsh_route](../data-sources/route.md)
- [Property reference](data-sources--route--reference.md)
- routes

<a id="section"></a>

Type: `"list"`. Computed.

List of routes to match for incoming request.

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

## Direct properties

- [bot_defense_javascript_injection](data-sources--route--properties--routes--bot_defense_javascript_injection.md): complete subsection reference.

<a id="schema-routes--disable_location_add"></a>

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

- [inherited_bot_defense_javascript_injection](data-sources--route--properties--routes--inherited_bot_defense_javascript_injection.md): complete subsection reference.

- [inherited_waf_exclusion](data-sources--route--properties--routes--inherited_waf_exclusion.md): complete subsection reference.

- [match](data-sources--route--properties--routes--match.md): complete subsection reference.

- [request_cookies_to_add](data-sources--route--properties--routes--request_cookies_to_add.md): complete subsection reference.

<a id="schema-routes--request_cookies_to_remove"></a>

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

- [request_headers_to_add](data-sources--route--properties--routes--request_headers_to_add.md): complete subsection reference.

<a id="schema-routes--request_headers_to_remove"></a>

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

- [response_cookies_to_add](data-sources--route--properties--routes--response_cookies_to_add.md): complete subsection reference.

<a id="schema-routes--response_cookies_to_remove"></a>

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

- [response_headers_to_add](data-sources--route--properties--routes--response_headers_to_add.md): complete subsection reference.

<a id="schema-routes--response_headers_to_remove"></a>

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

- [route_destination](data-sources--route--properties--routes--route_destination.md): complete subsection reference.

- [route_direct_response](data-sources--route--properties--routes--route_direct_response.md): complete subsection reference.

- [route_redirect](data-sources--route--properties--routes--route_redirect.md): complete subsection reference.

- [service_policy](data-sources--route--properties--routes--service_policy.md): complete subsection reference.

- [waf_exclusion_policy](data-sources--route--properties--routes--waf_exclusion_policy.md): complete subsection reference.

- [waf_type](data-sources--route--properties--routes--waf_type.md): complete subsection reference.

## Next pages

- [routes.bot_defense_javascript_injection](data-sources--route--properties--routes--bot_defense_javascript_injection.md)
- [routes.inherited_bot_defense_javascript_injection](data-sources--route--properties--routes--inherited_bot_defense_javascript_injection.md)
- [routes.inherited_waf_exclusion](data-sources--route--properties--routes--inherited_waf_exclusion.md)
- [routes.match](data-sources--route--properties--routes--match.md)
- [routes.request_cookies_to_add](data-sources--route--properties--routes--request_cookies_to_add.md)
- [routes.request_headers_to_add](data-sources--route--properties--routes--request_headers_to_add.md)
- [routes.response_cookies_to_add](data-sources--route--properties--routes--response_cookies_to_add.md)
- [routes.response_headers_to_add](data-sources--route--properties--routes--response_headers_to_add.md)
- [routes.route_destination](data-sources--route--properties--routes--route_destination.md)
- [routes.route_direct_response](data-sources--route--properties--routes--route_direct_response.md)
- [routes.route_redirect](data-sources--route--properties--routes--route_redirect.md)
- [routes.service_policy](data-sources--route--properties--routes--service_policy.md)
- [routes.waf_exclusion_policy](data-sources--route--properties--routes--waf_exclusion_policy.md)
- [routes.waf_type](data-sources--route--properties--routes--waf_type.md)
- [Property reference](data-sources--route--reference.md)
- [xcsh_route](../data-sources/route.md)
