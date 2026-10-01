---
page_title: "custom_network_config.sli_config.no_v6_static_routes"
subcategory: ""
description: "custom_network_config.sli_config.no_v6_static_routes for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1263, "body_sha256": "sha256:478c3089a87462bee2f5e329f29ca721d24b90746c6f72c11b575e1253445a0f", "canonical_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:no_v6_static_routes", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:no_v6_static_routes", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config", "path": "docs/guides/resources--securemesh_site--properties--custom_network_config--sli_config--no_v6_static_routes.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "sli_config", "no_v6_static_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/sli_config/no_v6_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.sli_config.no_v6_static_routes for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.sli_config.no_v6_static_routes

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md)
- [Property reference](resources--securemesh_site--reference.md)
- [custom_network_config](resources--securemesh_site--properties--custom_network_config.md)
- [custom_network_config.sli_config](resources--securemesh_site--properties--custom_network_config--sli_config.md)
- custom_network_config.sli_config.no_v6_static_routes

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no v6 static routes.

Upstream description:

This can be used for messages where no values are needed.

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
no_v6_static_routes = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [custom_network_config.sli_config](resources--securemesh_site--properties--custom_network_config--sli_config.md)
- [xcsh_securemesh_site](../resources/securemesh_site.md)
