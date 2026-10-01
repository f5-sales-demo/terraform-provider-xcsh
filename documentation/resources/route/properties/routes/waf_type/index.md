---
page_title: "routes.waf_type"
subcategory: ""
description: "routes.waf_type for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 2419, "body_sha256": "sha256:0825d6e921c40f0993bf9a9097bf5cd930224e03a80b17b0c8d8f8a429245bdf", "child_ids": ["xcsh-docs:resources:route:properties:routes:waf_type:app_firewall", "xcsh-docs:resources:route:properties:routes:waf_type:disable_waf", "xcsh-docs:resources:route:properties:routes:waf_type:inherit_waf"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:waf_type", "parent_id": "xcsh-docs:resources:route:properties:routes", "path": "documentation/resources/route/properties/routes/waf_type/index.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["routes", "waf_type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/waf_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.waf_type for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.waf_type

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
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

- [app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/app_firewall/): complete subsection reference.

- [disable_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/disable_waf/): complete subsection reference.

- [inherit_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/inherit_waf/): complete subsection reference.

## Next pages

- [routes.waf_type.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/app_firewall/)
- [routes.waf_type.disable_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/disable_waf/)
- [routes.waf_type.inherit_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/inherit_waf/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
