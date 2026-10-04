---
page_title: "routes"
subcategory: ""
description: "List of routes to match for incoming request."
xcsh_docs: {"aliases": ["routes"], "body_bytes": 11351, "body_sha256": "sha256:14605523057de005187c78cc8f94253d00a13a1b68c4774e1f6c59742a0f5262", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection", "xcsh-docs:data-sources:route:properties:routes:inherited_bot_defense_javascript_injection", "xcsh-docs:data-sources:route:properties:routes:inherited_waf_exclusion", "xcsh-docs:data-sources:route:properties:routes:match", "xcsh-docs:data-sources:route:properties:routes:request_cookies_to_add", "xcsh-docs:data-sources:route:properties:routes:request_headers_to_add", "xcsh-docs:data-sources:route:properties:routes:response_cookies_to_add", "xcsh-docs:data-sources:route:properties:routes:response_headers_to_add", "xcsh-docs:data-sources:route:properties:routes:route_destination", "xcsh-docs:data-sources:route:properties:routes:route_direct_response", "xcsh-docs:data-sources:route:properties:routes:route_redirect", "xcsh-docs:data-sources:route:properties:routes:service_policy", "xcsh-docs:data-sources:route:properties:routes:waf_exclusion_policy", "xcsh-docs:data-sources:route:properties:routes:waf_type"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes", "parent_id": "xcsh-docs:data-sources:route:reference", "path": "documentation/data-sources/route/properties/routes/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021", "registry_path": "docs/guides/data-sources--route--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes"], "schema_version": 1, "sections": [{"aliases": ["routes bot defense javascript injection"], "anchor": "section", "description": "Bot Defense Javascript Injection Configuration for inline bot defense deployments.", "document_id": "xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "bot_defense_javascript_injection"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes disable location add"], "anchor": "schema-routes--disable_location_add", "description": "Disables append of x-F5 Distributed Cloud-location = <RE-site-name> at route level, if it is configured at virtual-host level. This configuration is ignored on CE sites.", "document_id": "xcsh-docs:data-sources:route:properties:routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "disable_location_add"], "syntax": "attribute", "type": "bool"}, {"aliases": ["routes inherited bot defense javascript injection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:route:properties:routes:inherited_bot_defense_javascript_injection", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "inherited_bot_defense_javascript_injection"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes inherited waf exclusion"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:route:properties:routes:inherited_waf_exclusion", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "inherited_waf_exclusion"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes match"], "anchor": "section", "description": "Route match condition.", "document_id": "xcsh-docs:data-sources:route:properties:routes:match", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "match"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes request cookies to add"], "anchor": "section", "description": "Cookies are key-value pairs to be added to HTTP request being routed towards upstream.", "document_id": "xcsh-docs:data-sources:route:properties:routes:request_cookies_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "request_cookies_to_add"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes request cookies to remove"], "anchor": "schema-routes--request_cookies_to_remove", "description": "List of keys of Cookies to be removed from the HTTP request being sent towards upstream.", "document_id": "xcsh-docs:data-sources:route:properties:routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "request_cookies_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["routes request headers to add"], "anchor": "section", "description": "Headers are key-value pairs to be added to HTTP requests being sent towards upstream. Headers specified at this level are applied before headers from the enclosing VirtualHost object level.", "document_id": "xcsh-docs:data-sources:route:properties:routes:request_headers_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "request_headers_to_add"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes request headers to remove"], "anchor": "schema-routes--request_headers_to_remove", "description": "List of keys of Headers to be removed from the HTTP request being sent towards upstream.", "document_id": "xcsh-docs:data-sources:route:properties:routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "request_headers_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["routes response cookies to add"], "anchor": "section", "description": "Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream.", "document_id": "xcsh-docs:data-sources:route:properties:routes:response_cookies_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "response_cookies_to_add"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes response cookies to remove"], "anchor": "schema-routes--response_cookies_to_remove", "description": "List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed.", "document_id": "xcsh-docs:data-sources:route:properties:routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "response_cookies_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["routes response headers to add"], "anchor": "section", "description": "Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers specified at this level are applied before headers from the enclosing VirtualHost object level.", "document_id": "xcsh-docs:data-sources:route:properties:routes:response_headers_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "response_headers_to_add"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes response headers to remove"], "anchor": "schema-routes--response_headers_to_remove", "description": "List of keys of Headers to be removed from the HTTP response being sent towards downstream.", "document_id": "xcsh-docs:data-sources:route:properties:routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "response_headers_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["routes route destination"], "anchor": "section", "description": "List of destination to choose if the route is match.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "route_destination"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route direct response"], "anchor": "section", "description": "Send this direct response in case of route match action is direct response.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_direct_response", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "route_direct_response"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route redirect"], "anchor": "section", "description": "Route redirect parameters when match action is redirect.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_redirect", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "route_redirect"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes service policy"], "anchor": "section", "description": "ServicePolicy configuration details at route level.", "document_id": "xcsh-docs:data-sources:route:properties:routes:service_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "service_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes waf exclusion policy"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:route:properties:routes:waf_exclusion_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "waf_exclusion_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes waf type"], "anchor": "section", "description": "WAF instance will be pointing to an app_firewall object.", "document_id": "xcsh-docs:data-sources:route:properties:routes:waf_type", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "waf_type"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "List of routes to match for incoming request.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["routeCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/bot_defense_javascript_injection/): complete subsection reference.

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

- [inherited_bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/inherited_bot_defense_javascript_injection/): complete subsection reference.

- [inherited_waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/inherited_waf_exclusion/): complete subsection reference.

- [match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/match/): complete subsection reference.

- [request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/request_cookies_to_add/): complete subsection reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/request_headers_to_add/): complete subsection reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/response_cookies_to_add/): complete subsection reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/response_headers_to_add/): complete subsection reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/): complete subsection reference.

- [route_direct_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_direct_response/): complete subsection reference.

- [route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_redirect/): complete subsection reference.

- [service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/service_policy/): complete subsection reference.

- [waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/waf_exclusion_policy/): complete subsection reference.

- [waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/waf_type/): complete subsection reference.

## Next pages

- [routes.bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/bot_defense_javascript_injection/)
- [routes.inherited_bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/inherited_bot_defense_javascript_injection/)
- [routes.inherited_waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/inherited_waf_exclusion/)
- [routes.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/match/)
- [routes.request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/request_cookies_to_add/)
- [routes.request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/request_headers_to_add/)
- [routes.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/response_cookies_to_add/)
- [routes.response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/response_headers_to_add/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/)
- [routes.route_direct_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_direct_response/)
- [routes.route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_redirect/)
- [routes.service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/service_policy/)
- [routes.waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/waf_exclusion_policy/)
- [routes.waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/waf_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
