---
page_title: "cloudfront.protected_endpoints.flow_label.profile_management"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label.profile_management for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2703, "body_sha256": "sha256:6dad8e5dbddca01473ea785c89b7cdecfe7e2d588c8136ca5531e68ced026ae7", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:create", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:update", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:view"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "path": "docs/guides/resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "profile_management"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/profile_management/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label.profile_management for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# cloudfront.protected_endpoints.flow_label.profile_management

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudfront](resources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](resources--protected_application--properties--cloudfront--protected_endpoints.md)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
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

- [create](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management--create.md): complete subsection reference.

- [update](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management--update.md): complete subsection reference.

- [view](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management--view.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.profile_management.create](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management--create.md)
- [cloudfront.protected_endpoints.flow_label.profile_management.update](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management--update.md)
- [cloudfront.protected_endpoints.flow_label.profile_management.view](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management--view.md)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
- [xcsh_protected_application](../resources/protected_application.md)
