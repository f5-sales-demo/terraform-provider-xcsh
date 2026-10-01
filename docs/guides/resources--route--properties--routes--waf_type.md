---
page_title: "routes.waf_type"
subcategory: ""
description: "routes.waf_type for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1868, "body_sha256": "sha256:5f9cf5214e94511971df9ac1f8bad49c29ffbe95358288eb9b42580f1ca40542", "canonical_id": "xcsh-docs:resources:route:properties:routes:waf_type", "child_ids": ["xcsh-docs:resources:route:properties:routes:waf_type:app_firewall", "xcsh-docs:resources:route:properties:routes:waf_type:disable_waf", "xcsh-docs:resources:route:properties:routes:waf_type:inherit_waf"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:waf_type", "parent_id": "xcsh-docs:resources:route:properties:routes", "path": "docs/guides/resources--route--properties--routes--waf_type.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "waf_type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/waf_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.waf_type for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.waf_type

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- routes.waf_type

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

WAF instance will be pointing to an app\_firewall object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall",
    "disable_waf"),
  validators.ConflictingObjectAttributes("app_firewall",
    "inherit_waf"),
  validators.ConflictingObjectAttributes("disable_waf",
    "inherit_waf")}
```

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

Terraform syntax:

```terraform
waf_type {
  # Configure direct properties listed below.
}
```

## Direct properties

- [app_firewall](resources--route--properties--routes--waf_type--app_firewall.md): complete subsection reference.

- [disable_waf](resources--route--properties--routes--waf_type--disable_waf.md): complete subsection reference.

- [inherit_waf](resources--route--properties--routes--waf_type--inherit_waf.md): complete subsection reference.

## Next pages

- [routes.waf_type.app_firewall](resources--route--properties--routes--waf_type--app_firewall.md)
- [routes.waf_type.disable_waf](resources--route--properties--routes--waf_type--disable_waf.md)
- [routes.waf_type.inherit_waf](resources--route--properties--routes--waf_type--inherit_waf.md)
- [routes](resources--route--properties--routes.md)
- [xcsh_route](../resources/route.md)
