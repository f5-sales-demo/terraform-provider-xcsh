---
page_title: "default_pool.advanced_options.enable_subsets"
subcategory: "Load Balancing"
description: "default_pool.advanced_options.enable_subsets for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2952, "body_sha256": "sha256:3e4955ac91e709be6ca14f61e2296d2a4786e76f42fb7f54ecf479bb97228a5a", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:any_endpoint", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:endpoint_subsets", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:fail_request"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "advanced_options", "enable_subsets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.advanced_options.enable_subsets for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.enable_subsets

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.advanced_options](resources--http_loadbalancer--properties--default_pool--advanced_options.md)
- default_pool.advanced_options.enable_subsets

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

- [any_endpoint](resources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--any_endpoint.md): complete subsection reference.

- [default_subset](resources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--default_subset.md): complete subsection reference.

- [endpoint_subsets](resources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--endpoint_subsets.md): complete subsection reference.

- [fail_request](resources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--fail_request.md): complete subsection reference.

## Next pages

- [default_pool.advanced_options.enable_subsets.any_endpoint](resources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--any_endpoint.md)
- [default_pool.advanced_options.enable_subsets.default_subset](resources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--default_subset.md)
- [default_pool.advanced_options.enable_subsets.endpoint_subsets](resources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--endpoint_subsets.md)
- [default_pool.advanced_options.enable_subsets.fail_request](resources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--fail_request.md)
- [default_pool.advanced_options](resources--http_loadbalancer--properties--default_pool--advanced_options.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
