---
page_title: "sli_to_slo_snat.default_gw_snat"
subcategory: "Networking"
description: "sli_to_slo_snat.default_gw_snat for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1078, "body_sha256": "sha256:b29b89ee20798cfdffd912feedc793172d08bfff61183b90292ef46420b16fe0", "canonical_id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:default_gw_snat", "child_ids": [], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:default_gw_snat", "parent_id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat", "path": "docs/guides/resources--network_connector--properties--sli_to_slo_snat--default_gw_snat.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sli_to_slo_snat", "default_gw_snat"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/sli_to_slo_snat/default_gw_snat/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sli_to_slo_snat.default_gw_snat for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sli_to_slo_snat.default_gw_snat

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md)
- [Property reference](resources--network_connector--reference.md)
- [sli_to_slo_snat](resources--network_connector--properties--sli_to_slo_snat.md)
- sli_to_slo_snat.default_gw_snat

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default gw snat.

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
default_gw_snat {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [sli_to_slo_snat](resources--network_connector--properties--sli_to_slo_snat.md)
- [xcsh_network_connector](../resources/network_connector.md)
