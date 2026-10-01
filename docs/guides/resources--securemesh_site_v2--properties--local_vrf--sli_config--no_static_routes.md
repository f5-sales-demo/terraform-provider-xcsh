---
page_title: "local_vrf.sli_config.no_static_routes"
subcategory: ""
description: "local_vrf.sli_config.no_static_routes for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1179, "body_sha256": "sha256:46696eb479b7d3d51a947c225aad211525d9273ff68c4485d32e22ea1347afed", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config:no_static_routes", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config:no_static_routes", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config", "path": "docs/guides/resources--securemesh_site_v2--properties--local_vrf--sli_config--no_static_routes.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_vrf", "sli_config", "no_static_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/local_vrf/sli_config/no_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_vrf.sli_config.no_static_routes for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf.sli_config.no_static_routes

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [local_vrf](resources--securemesh_site_v2--properties--local_vrf.md)
- [local_vrf.sli_config](resources--securemesh_site_v2--properties--local_vrf--sli_config.md)
- local_vrf.sli_config.no_static_routes

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static routes.

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
no_static_routes = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [local_vrf.sli_config](resources--securemesh_site_v2--properties--local_vrf--sli_config.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
