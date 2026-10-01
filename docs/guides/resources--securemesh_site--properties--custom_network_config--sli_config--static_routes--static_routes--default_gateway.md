---
page_title: "custom_network_config.sli_config.static_routes.static_routes.default_gateway"
subcategory: ""
description: "custom_network_config.sli_config.static_routes.static_routes.default_gateway for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1678, "body_sha256": "sha256:4ae50a73b40fcabb32909d5918143479f35043246f07453924af33c3ef015dfc", "canonical_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_routes:static_routes:default_gateway", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_routes:static_routes:default_gateway", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_routes:static_routes", "path": "docs/guides/resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--default_gateway.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "sli_config", "static_routes", "static_routes", "default_gateway"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/default_gateway/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.sli_config.static_routes.static_routes.default_gateway for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.sli_config.static_routes.static_routes.default_gateway

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md)
- [Property reference](resources--securemesh_site--reference.md)
- [custom_network_config](resources--securemesh_site--properties--custom_network_config.md)
- [custom_network_config.sli_config](resources--securemesh_site--properties--custom_network_config--sli_config.md)
- [custom_network_config.sli_config.static_routes](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes.md)
- [custom_network_config.sli_config.static_routes.static_routes](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes.md)
- custom_network_config.sli_config.static_routes.static_routes.default_gateway

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [custom_network_config.sli_config.static_routes.static_routes](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes.md)
- [xcsh_securemesh_site](../resources/securemesh_site.md)
