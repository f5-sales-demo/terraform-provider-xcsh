---
page_title: "routes.waf_type.app_firewall"
subcategory: ""
description: "routes.waf_type.app_firewall for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1097, "body_sha256": "sha256:0be96509dd8062ad1bb0f77dd4429a2d9a96aa0a84ddaed3d26ea0045cbaa29b", "canonical_id": "xcsh-docs:data-sources:route:properties:routes:waf_type:app_firewall", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:waf_type:app_firewall:app_firewall"], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:waf_type:app_firewall", "parent_id": "xcsh-docs:data-sources:route:properties:routes:waf_type", "path": "docs/guides/data-sources--route--properties--routes--waf_type--app_firewall.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "waf_type", "app_firewall"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/waf_type/app_firewall/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.waf_type.app_firewall for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.waf_type.app_firewall

Breadcrumbs:

- [xcsh_route](../data-sources/route.md)
- [Property reference](data-sources--route--reference.md)
- [routes](data-sources--route--properties--routes.md)
- [routes.waf_type](data-sources--route--properties--routes--waf_type.md)
- routes.waf_type.app_firewall

<a id="section"></a>

Type: `"single"`. Computed.

List of references to the app\_firewall configuration objects.

Upstream description:

A list of references to the app\_firewall configuration objects.

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

## Direct properties

- [app_firewall](data-sources--route--properties--routes--waf_type--app_firewall--app_firewall.md): complete subsection reference.

## Next pages

- [routes.waf_type.app_firewall.app_firewall](data-sources--route--properties--routes--waf_type--app_firewall--app_firewall.md)
- [routes.waf_type](data-sources--route--properties--routes--waf_type.md)
- [xcsh_route](../data-sources/route.md)
