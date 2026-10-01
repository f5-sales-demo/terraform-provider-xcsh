---
page_title: "custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway"
subcategory: ""
description: "custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1692, "body_sha256": "sha256:5d4647d08d67bfda418edb5b4cc62e82619985f87ac1b697346204a9dc180a70", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:static_v6_routes:static_routes:default_gateway", "child_ids": [], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:static_v6_routes:static_routes:default_gateway", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:static_v6_routes:static_routes", "path": "docs/guides/resources--voltstack_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--default_gateway.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "sli_config", "static_v6_routes", "static_routes", "default_gateway"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/default_gateway/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_network_config](resources--voltstack_site--properties--custom_network_config.md)
- [custom_network_config.sli_config](resources--voltstack_site--properties--custom_network_config--sli_config.md)
- [custom_network_config.sli_config.static_v6_routes](resources--voltstack_site--properties--custom_network_config--sli_config--static_v6_routes.md)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes.md)
- custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway

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

- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--voltstack_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
