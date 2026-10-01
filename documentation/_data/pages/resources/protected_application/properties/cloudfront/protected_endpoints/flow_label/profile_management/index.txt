---
page_title: "cloudfront.protected_endpoints.flow_label.profile_management"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label.profile_management for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 3438, "body_sha256": "sha256:2ad325d563955710bcfd25d428b76bf9d0e2888fa40861f81ecc2bcdc62871e5", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:create", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:update", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:view"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "path": "documentation/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/profile_management/index.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "profile_management"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/profile_management/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label.profile_management for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.profile_management

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/)
- [cloudfront.protected_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/)
- cloudfront.protected_endpoints.flow_label.profile_management

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Profile Management Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "update"),
  validators.ConflictingObjectAttributes("create",
    "view"),
  validators.ConflictingObjectAttributes("update",
    "view")}
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
  "x-ves-oneof-field-label_choice": "[\"create\",\"update\",\"view\"]"
}
```

Terraform syntax:

```terraform
profile_management {
  # Configure direct properties listed below.
}
```

## Direct properties

- [create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/profile_management/create/): complete subsection reference.

- [update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/profile_management/update/): complete subsection reference.

- [view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/profile_management/view/): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.profile_management.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/profile_management/create/)
- [cloudfront.protected_endpoints.flow_label.profile_management.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/profile_management/update/)
- [cloudfront.protected_endpoints.flow_label.profile_management.view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/profile_management/view/)
- [cloudfront.protected_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
