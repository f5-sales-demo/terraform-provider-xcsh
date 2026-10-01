---
page_title: "http_loadbalancer"
subcategory: ""
description: "http_loadbalancer for xcsh_app_api_group."
xcsh_docs: {"aliases": [], "body_bytes": 1096, "body_sha256": "sha256:d5247c76ba2cb55d8070590ed16b48c1229c0935a6e8349434341df86f827c0c", "canonical_id": "xcsh-docs:resources:app_api_group:properties:http_loadbalancer", "child_ids": ["xcsh-docs:resources:app_api_group:properties:http_loadbalancer:http_loadbalancer"], "collection_id": "xcsh-docs:resources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_api_group:properties:http_loadbalancer", "parent_id": "xcsh-docs:resources:app_api_group:reference", "path": "docs/guides/resources--app_api_group--properties--http_loadbalancer.md", "provider_name": "app_api_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_loadbalancer"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_api_group/properties/http_loadbalancer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_loadbalancer for xcsh_app_api_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_loadbalancer

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md)
- [Property reference](resources--app_api_group--reference.md)
- http_loadbalancer

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Set the scope of the API Group to a specific HTTP Loadbalancer.

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
http_loadbalancer {
  # Configure direct properties listed below.
}
```

## Direct properties

- [http_loadbalancer](resources--app_api_group--properties--http_loadbalancer--http_loadbalancer.md): complete subsection reference.

## Next pages

- [http_loadbalancer.http_loadbalancer](resources--app_api_group--properties--http_loadbalancer--http_loadbalancer.md)
- [Property reference](resources--app_api_group--reference.md)
- [xcsh_app_api_group](../resources/app_api_group.md)
