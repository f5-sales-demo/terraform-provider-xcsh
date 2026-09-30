---
page_title: "routes.match.incoming_port.no_port_match"
subcategory: ""
description: "routes.match.incoming_port.no_port_match for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1037, "body_sha256": "sha256:ef6473ca477606a2e9eaa40e4fc8363529423b1a9d0e33d4b56a1b38a0ae74b3", "canonical_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port:no_port_match", "child_ids": [], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:match:incoming_port:no_port_match", "parent_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port", "path": "docs/guides/resources--route--properties--routes--match--incoming_port--no_port_match.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "match", "incoming_port", "no_port_match"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/match/incoming_port/no_port_match/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.match.incoming_port.no_port_match for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
