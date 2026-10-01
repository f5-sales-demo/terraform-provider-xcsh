---
page_title: "sli_to_slo_snat.interface_ip"
subcategory: "Networking"
description: "sli_to_slo_snat.interface_ip for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1043, "body_sha256": "sha256:28717982a6d849759bca62e9e81d16a22b586e1b2c4bd9796deb6d21849012b1", "canonical_id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:interface_ip", "child_ids": [], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:interface_ip", "parent_id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat", "path": "docs/guides/resources--network_connector--properties--sli_to_slo_snat--interface_ip.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sli_to_slo_snat", "interface_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/sli_to_slo_snat/interface_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sli_to_slo_snat.interface_ip for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sli_to_slo_snat.interface_ip

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md)
- [Property reference](resources--network_connector--reference.md)
- [sli_to_slo_snat](resources--network_connector--properties--sli_to_slo_snat.md)
- sli_to_slo_snat.interface_ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Enable this option

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
interface_ip {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [sli_to_slo_snat](resources--network_connector--properties--sli_to_slo_snat.md)
- [xcsh_network_connector](../resources/network_connector.md)
