---
page_title: "routes.service_policy"
subcategory: ""
description: "routes.service_policy for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1371, "body_sha256": "sha256:7bb58393c95616c219d3df73af2df602551b41640bffc4ecf4f56a0f27909cfd", "child_ids": [], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:service_policy", "parent_id": "xcsh-docs:resources:route:properties:routes", "path": "documentation/resources/route/properties/routes/service_policy/index.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["routes", "service_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/service_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.service_policy for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.service_policy

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- routes.service_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ServicePolicy configuration details at route level.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-service_policy_choice": "[\"disable\"]"
}
```

Terraform syntax:

```terraform
service_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--service_policy--disable_spec"></a>

### disable_spec property

Type: `"bool"`. Optional.

Exclusive with \[\] disable service policy at route level, if it is configured at virtual-host
level.

## Next pages

- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
