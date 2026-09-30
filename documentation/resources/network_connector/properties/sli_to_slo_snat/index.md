---
page_title: "sli_to_slo_snat"
subcategory: "Networking"
description: "sli_to_slo_snat for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1804, "body_sha256": "sha256:762591a093f94a90a550b6ef479cd2ff12493eabc0a986326c49492a4c357fb9", "child_ids": ["xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:default_gw_snat", "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:interface_ip"], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat", "parent_id": "xcsh-docs:resources:network_connector:reference", "path": "documentation/resources/network_connector/properties/sli_to_slo_snat/index.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["sli_to_slo_snat"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/sli_to_slo_snat/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sli_to_slo_snat for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# sli_to_slo_snat

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/)
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

- [default_gw_snat](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_slo_snat/default_gw_snat/): complete subsection reference.

- [interface_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_slo_snat/interface_ip/): complete subsection reference.

## Next pages

- [sli_to_slo_snat.default_gw_snat](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_slo_snat/default_gw_snat/)
- [sli_to_slo_snat.interface_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_slo_snat/interface_ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
