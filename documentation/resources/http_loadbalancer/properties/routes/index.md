---
page_title: "routes"
subcategory: "Load Balancing"
description: "Routes allow users to define match condition on a path and/or HTTP method to either forward matching traffic to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "routes", "upstream servers"], "body_bytes": 4726, "body_sha256": "sha256:0d147ceb9b0f25ee0e394234da8e2a3703ad8f9dc95393f54c01e082215502ff", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object", "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route", "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route", "xcsh-docs:resources:http_loadbalancer:properties:routes:route_state_disabled", "xcsh-docs:resources:http_loadbalancer:properties:routes:route_state_enabled", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/routes/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-024.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:custom_route_object,direct_response_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:custom_route_object,redirect_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:custom_route_object,simple_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:custom_route_object,direct_response_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:direct_response_route,redirect_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:direct_response_route,simple_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:custom_route_object,redirect_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:direct_response_route,redirect_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:redirect_route,simple_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:route_state_disabled,route_state_enabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:route_state_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:route_state_disabled,route_state_enabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:route_state_enabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:custom_route_object,simple_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:direct_response_route,simple_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:redirect_route,simple_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes"], "schema_version": 1, "sections": [{"aliases": ["custom route object"], "anchor": "section", "description": "A custom route uses a route object created outside of this view.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.custom_route_object:ConflictingObjectAttributes:caching_disable,caching_inherit", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object:caching_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.custom_route_object:ConflictingObjectAttributes:caching_disable,caching_inherit", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object:caching_inherit", "type": "conflicts"}], "schema_path": ["routes", "custom_route_object"], "syntax": "block", "type": "object"}, {"aliases": ["direct response route"], "anchor": "section", "description": "A direct response route matches on path, incoming header, incoming port and/or HTTP method and responds directly to the matching traffic.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "direct_response_route"], "syntax": "block", "type": "object"}, {"aliases": ["redirect route"], "anchor": "section", "description": "A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the matching traffic to a different URL.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "redirect_route"], "syntax": "block", "type": "object"}, {"aliases": ["route state disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:route_state_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_state_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["route state enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:route_state_enabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_state_enabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["simple route"], "anchor": "section", "description": "A simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the matching traffic to the associated pools.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--simple_route--host_rewrite", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:auto_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--host_rewrite", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:disable_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:auto_host_rewrite,disable_host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:auto_host_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:auto_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:auto_host_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:caching_disable,caching_inherit", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:caching_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:caching_disable,caching_inherit", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:caching_inherit", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:auto_host_rewrite,disable_host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:disable_host_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:disable_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:disable_host_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route:RequiredObjectAttributes:origin_pools", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:origin_pools", "type": "requires"}], "schema_path": ["routes", "simple_route"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Routes allow users to define match condition on a path and/or HTTP method to either forward matching traffic to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- routes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Routes allow users to define match condition on a path and/or HTTP method to either forward matching
traffic to origin pool or redirect matching traffic to a different URL or respond directly to
matching traffic.

Upstream description:

Routes allow users to define match condition on a path and/or HTTP method to either forward matching
traffic to origin pool or redirect matching traffic to a different URL or respond directly to
matching traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_route_object",
    "direct_response_route"),
  validators.ConflictingListObjectAttributes("custom_route_object",
    "redirect_route"),
  validators.ConflictingListObjectAttributes("custom_route_object",
    "simple_route"),
  validators.ConflictingListObjectAttributes("direct_response_route",
    "redirect_route"),
  validators.ConflictingListObjectAttributes("direct_response_route",
    "simple_route"),
  validators.ConflictingListObjectAttributes("redirect_route",
    "simple_route"),
  validators.ConflictingListObjectAttributes("route_state_disabled",
    "route_state_enabled")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
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

- [custom_route_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/custom_route_object/): complete subsection reference.

- [direct_response_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/direct_response_route/): complete subsection reference.

- [redirect_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/redirect_route/): complete subsection reference.

- [route_state_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/route_state_disabled/): complete subsection reference.

- [route_state_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/route_state_enabled/): complete subsection reference.

- [simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/): complete subsection reference.

## Next pages

- [routes.custom_route_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/custom_route_object/)
- [routes.direct_response_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/direct_response_route/)
- [routes.redirect_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/redirect_route/)
- [routes.route_state_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/route_state_disabled/)
- [routes.route_state_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/route_state_enabled/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
