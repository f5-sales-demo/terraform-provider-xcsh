---
page_title: "custom_network_config.slo_config"
subcategory: ""
description: "custom_network_config.slo_config for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 3648, "body_sha256": "sha256:5d7ae3130ecbfed020aadecd3cd865df2e2c9fddf15c2f0fe416054270bef34a", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:dc_cluster_group", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:labels", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:no_dc_cluster_group", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:no_static_routes", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:no_static_v6_routes", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config", "path": "docs/guides/resources--voltstack_site--properties--custom_network_config--slo_config.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "slo_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/slo_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.slo_config for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_network_config.slo_config

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_network_config](resources--voltstack_site--properties--custom_network_config.md)
- custom_network_config.slo_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_static_v6_routes",
    "static_v6_routes")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_static_v6_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
slo_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dc_cluster_group](resources--voltstack_site--properties--custom_network_config--slo_config--dc_cluster_group.md): complete subsection reference.

- [labels](resources--voltstack_site--properties--custom_network_config--slo_config--labels.md): complete subsection reference.

- [no_dc_cluster_group](resources--voltstack_site--properties--custom_network_config--slo_config--no_dc_cluster_group.md): complete subsection reference.

- [no_static_routes](resources--voltstack_site--properties--custom_network_config--slo_config--no_static_routes.md): complete subsection reference.

- [no_static_v6_routes](resources--voltstack_site--properties--custom_network_config--slo_config--no_static_v6_routes.md): complete subsection reference.

- [static_routes](resources--voltstack_site--properties--custom_network_config--slo_config--static_routes.md): complete subsection reference.

- [static_v6_routes](resources--voltstack_site--properties--custom_network_config--slo_config--static_v6_routes.md): complete subsection reference.

## Next pages

- [custom_network_config.slo_config.dc_cluster_group](resources--voltstack_site--properties--custom_network_config--slo_config--dc_cluster_group.md)
- [custom_network_config.slo_config.labels](resources--voltstack_site--properties--custom_network_config--slo_config--labels.md)
- [custom_network_config.slo_config.no_dc_cluster_group](resources--voltstack_site--properties--custom_network_config--slo_config--no_dc_cluster_group.md)
- [custom_network_config.slo_config.no_static_routes](resources--voltstack_site--properties--custom_network_config--slo_config--no_static_routes.md)
- [custom_network_config.slo_config.no_static_v6_routes](resources--voltstack_site--properties--custom_network_config--slo_config--no_static_v6_routes.md)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--properties--custom_network_config--slo_config--static_routes.md)
- [custom_network_config.slo_config.static_v6_routes](resources--voltstack_site--properties--custom_network_config--slo_config--static_v6_routes.md)
- [custom_network_config](resources--voltstack_site--properties--custom_network_config.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
