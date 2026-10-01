---
page_title: "cdn_loadbalancer"
subcategory: ""
description: "cdn_loadbalancer for xcsh_app_api_group."
xcsh_docs: {"aliases": [], "body_bytes": 1085, "body_sha256": "sha256:a5435143d721b5ecc134c1f51ed6762e4a2cae50ae871114c3b0ab26d4fbd9ac", "canonical_id": "xcsh-docs:resources:app_api_group:properties:cdn_loadbalancer", "child_ids": ["xcsh-docs:resources:app_api_group:properties:cdn_loadbalancer:cdn_loadbalancer"], "collection_id": "xcsh-docs:resources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_api_group:properties:cdn_loadbalancer", "parent_id": "xcsh-docs:resources:app_api_group:reference", "path": "docs/guides/resources--app_api_group--properties--cdn_loadbalancer.md", "provider_name": "app_api_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cdn_loadbalancer"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_api_group/properties/cdn_loadbalancer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cdn_loadbalancer for xcsh_app_api_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cdn_loadbalancer

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md)
- [Property reference](resources--app_api_group--reference.md)
- cdn_loadbalancer

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Set the scope of the API Group to a specific CDN Loadbalancer.

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
cdn_loadbalancer {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cdn_loadbalancer](resources--app_api_group--properties--cdn_loadbalancer--cdn_loadbalancer.md): complete subsection reference.

## Next pages

- [cdn_loadbalancer.cdn_loadbalancer](resources--app_api_group--properties--cdn_loadbalancer--cdn_loadbalancer.md)
- [Property reference](resources--app_api_group--reference.md)
- [xcsh_app_api_group](../resources/app_api_group.md)
