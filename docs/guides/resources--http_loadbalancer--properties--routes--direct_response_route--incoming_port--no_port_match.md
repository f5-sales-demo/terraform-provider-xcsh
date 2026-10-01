---
page_title: "routes.direct_response_route.incoming_port.no_port_match"
subcategory: "Load Balancing"
description: "routes.direct_response_route.incoming_port.no_port_match for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1372, "body_sha256": "sha256:f047892875e4cbb7d81b85e1e3b2a55f70210a40554439f5a07e8066f889a1f7", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:incoming_port:no_port_match", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:incoming_port:no_port_match", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:incoming_port", "path": "docs/guides/resources--http_loadbalancer--properties--routes--direct_response_route--incoming_port--no_port_match.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "direct_response_route", "incoming_port", "no_port_match"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/direct_response_route/incoming_port/no_port_match/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.direct_response_route.incoming_port.no_port_match for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.direct_response_route.incoming_port.no_port_match

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [routes.direct_response_route](resources--http_loadbalancer--properties--routes--direct_response_route.md)
- [routes.direct_response_route.incoming_port](resources--http_loadbalancer--properties--routes--direct_response_route--incoming_port.md)
- routes.direct_response_route.incoming_port.no_port_match

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
no_port_match = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.direct_response_route.incoming_port](resources--http_loadbalancer--properties--routes--direct_response_route--incoming_port.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
