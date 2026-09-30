---
page_title: "advanced_options.enable_subsets"
subcategory: "Load Balancing"
description: "advanced_options.enable_subsets for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 2443, "body_sha256": "sha256:55178c3b26d271b588819c9e95766580b8eaf3ae5dc1819eb89499686e88e1d4", "canonical_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets", "child_ids": ["xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:any_endpoint", "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:default_subset", "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:endpoint_subsets", "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:fail_request"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "path": "docs/guides/resources--origin_pool--properties--advanced_options--enable_subsets.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options", "enable_subsets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/enable_subsets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.enable_subsets for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# advanced_options.enable_subsets

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- advanced_options.enable_subsets

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure subset OPTIONS for origin pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("endpoint_subsets"),
  validators.ConflictingObjectAttributes("any_endpoint",
    "default_subset"),
  validators.ConflictingObjectAttributes("any_endpoint",
    "fail_request"),
  validators.ConflictingObjectAttributes("default_subset",
    "fail_request")}
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
  "x-ves-oneof-field-fallback_policy_choice": "[\"any_endpoint\",\"default_subset\",\"fail_request\"]"
}
```

Terraform syntax:

```terraform
enable_subsets {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_endpoint](resources--origin_pool--properties--advanced_options--enable_subsets--any_endpoint.md): complete subsection reference.

- [default_subset](resources--origin_pool--properties--advanced_options--enable_subsets--default_subset.md): complete subsection reference.

- [endpoint_subsets](resources--origin_pool--properties--advanced_options--enable_subsets--endpoint_subsets.md): complete subsection reference.

- [fail_request](resources--origin_pool--properties--advanced_options--enable_subsets--fail_request.md): complete subsection reference.

## Next pages

- [advanced_options.enable_subsets.any_endpoint](resources--origin_pool--properties--advanced_options--enable_subsets--any_endpoint.md)
- [advanced_options.enable_subsets.default_subset](resources--origin_pool--properties--advanced_options--enable_subsets--default_subset.md)
- [advanced_options.enable_subsets.endpoint_subsets](resources--origin_pool--properties--advanced_options--enable_subsets--endpoint_subsets.md)
- [advanced_options.enable_subsets.fail_request](resources--origin_pool--properties--advanced_options--enable_subsets--fail_request.md)
- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
