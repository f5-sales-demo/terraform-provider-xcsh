---
page_title: "sli_to_slo_snat"
subcategory: "Networking"
description: "sli_to_slo_snat for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1495, "body_sha256": "sha256:c247e12d378e129575b41961b94dab106404955cf2cb6a7a1ce656660cbda49f", "canonical_id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat", "child_ids": ["xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:default_gw_snat", "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:interface_ip"], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat", "parent_id": "xcsh-docs:resources:network_connector:reference", "path": "docs/guides/resources--network_connector--properties--sli_to_slo_snat.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sli_to_slo_snat"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/sli_to_slo_snat/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sli_to_slo_snat for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sli_to_slo_snat

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md)
- [Property reference](resources--network_connector--reference.md)
- sli_to_slo_snat

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sli to slo snat.

Upstream description:

X-example: "" description.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-pool_choice": "[\"interface_ip\"]",
  "x-ves-oneof-field-routing_choice": "[\"default_gw_snat\"]"
}
```

Terraform syntax:

```terraform
sli_to_slo_snat {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_gw_snat](resources--network_connector--properties--sli_to_slo_snat--default_gw_snat.md): complete subsection reference.

- [interface_ip](resources--network_connector--properties--sli_to_slo_snat--interface_ip.md): complete subsection reference.

## Next pages

- [sli_to_slo_snat.default_gw_snat](resources--network_connector--properties--sli_to_slo_snat--default_gw_snat.md)
- [sli_to_slo_snat.interface_ip](resources--network_connector--properties--sli_to_slo_snat--interface_ip.md)
- [Property reference](resources--network_connector--reference.md)
- [xcsh_network_connector](../resources/network_connector.md)
