---
page_title: "routes.match.incoming_port.no_port_match"
subcategory: ""
description: "routes.match.incoming_port.no_port_match for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1136, "body_sha256": "sha256:3a49f90e74e846fd21e04059bf3e79966731502c798497664b3a32ec0f81ef0a", "canonical_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port:no_port_match", "child_ids": [], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:match:incoming_port:no_port_match", "parent_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port", "path": "docs/guides/resources--route--properties--routes--match--incoming_port--no_port_match.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "match", "incoming_port", "no_port_match"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/match/incoming_port/no_port_match/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.match.incoming_port.no_port_match for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.match.incoming_port.no_port_match

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- [routes.match](resources--route--properties--routes--match.md)
- [routes.match.incoming_port](resources--route--properties--routes--match--incoming_port.md)
- routes.match.incoming_port.no_port_match

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

- [routes.match.incoming_port](resources--route--properties--routes--match--incoming_port.md)
- [xcsh_route](../resources/route.md)
