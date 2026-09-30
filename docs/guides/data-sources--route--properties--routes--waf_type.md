---
page_title: "routes.waf_type"
subcategory: ""
description: "routes.waf_type for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1359, "body_sha256": "sha256:29384935060739d518f50a6326f7e68b154146fdcc257d6a24c3ae995a2f9025", "canonical_id": "xcsh-docs:data-sources:route:properties:routes:waf_type", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:waf_type:app_firewall", "xcsh-docs:data-sources:route:properties:routes:waf_type:disable_waf", "xcsh-docs:data-sources:route:properties:routes:waf_type:inherit_waf"], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:waf_type", "parent_id": "xcsh-docs:data-sources:route:properties:routes", "path": "docs/guides/data-sources--route--properties--routes--waf_type.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "waf_type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/waf_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.waf_type for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.waf_type

Breadcrumbs:

- [xcsh_route](../data-sources/route.md)
- [Property reference](data-sources--route--reference.md)
- [routes](data-sources--route--properties--routes.md)
- routes.waf_type

<a id="section"></a>

Type: `"single"`. Computed.

WAF instance will be pointing to an app\_firewall object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_type": "[\"app_firewall\",\"disable_waf\",\"inherit_waf\"]"
}
```

## Direct properties

- [app_firewall](data-sources--route--properties--routes--waf_type--app_firewall.md): complete subsection reference.

- [disable_waf](data-sources--route--properties--routes--waf_type--disable_waf.md): complete subsection reference.

- [inherit_waf](data-sources--route--properties--routes--waf_type--inherit_waf.md): complete subsection reference.

## Next pages

- [routes.waf_type.app_firewall](data-sources--route--properties--routes--waf_type--app_firewall.md)
- [routes.waf_type.disable_waf](data-sources--route--properties--routes--waf_type--disable_waf.md)
- [routes.waf_type.inherit_waf](data-sources--route--properties--routes--waf_type--inherit_waf.md)
- [routes](data-sources--route--properties--routes.md)
- [xcsh_route](../data-sources/route.md)
