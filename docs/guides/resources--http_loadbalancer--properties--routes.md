---
page_title: "routes"
subcategory: "Load Balancing"
description: "routes for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3918, "body_sha256": "sha256:e004ea0b94134ac1f0ecd04452a8eded49a03367a3ee89252420824e0e63f4b9", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object", "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route", "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route", "xcsh-docs:resources:http_loadbalancer:properties:routes:route_state_disabled", "xcsh-docs:resources:http_loadbalancer:properties:routes:route_state_enabled", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "docs/guides/resources--http_loadbalancer--properties--routes.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
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

- [custom_route_object](resources--http_loadbalancer--properties--routes--custom_route_object.md): complete subsection reference.

- [direct_response_route](resources--http_loadbalancer--properties--routes--direct_response_route.md): complete subsection reference.

- [redirect_route](resources--http_loadbalancer--properties--routes--redirect_route.md): complete subsection reference.

- [route_state_disabled](resources--http_loadbalancer--properties--routes--route_state_disabled.md): complete subsection reference.

- [route_state_enabled](resources--http_loadbalancer--properties--routes--route_state_enabled.md): complete subsection reference.

- [simple_route](resources--http_loadbalancer--properties--routes--simple_route.md): complete subsection reference.

## Next pages

- [routes.custom_route_object](resources--http_loadbalancer--properties--routes--custom_route_object.md)
- [routes.direct_response_route](resources--http_loadbalancer--properties--routes--direct_response_route.md)
- [routes.redirect_route](resources--http_loadbalancer--properties--routes--redirect_route.md)
- [routes.route_state_disabled](resources--http_loadbalancer--properties--routes--route_state_disabled.md)
- [routes.route_state_enabled](resources--http_loadbalancer--properties--routes--route_state_enabled.md)
- [routes.simple_route](resources--http_loadbalancer--properties--routes--simple_route.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
