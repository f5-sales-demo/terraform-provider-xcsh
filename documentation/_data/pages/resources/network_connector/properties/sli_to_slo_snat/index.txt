---
page_title: "sli_to_slo_snat"
subcategory: "Networking"
description: "X-example: \"\" description."
xcsh_docs: {"aliases": ["sli to slo snat"], "body_bytes": 1903, "body_sha256": "sha256:322a9dc301582f83465da46c54c1700e1f9dd3f242c12d5f4febb42a5b33c23d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:default_gw_snat", "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:interface_ip"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat", "parent_id": "xcsh-docs:resources:network_connector:reference", "path": "documentation/resources/network_connector/properties/sli_to_slo_snat/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0311230213310203-1002131223302201-1301011222023101-2212122300321021-2103313222121202-2010010213310323-3111100010201210-0122233201102021", "registry_path": "docs/guides/resources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sli_to_slo_snat"], "schema_version": 1, "sections": [{"aliases": ["default gw snat"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:default_gw_snat", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sli_to_slo_snat", "default_gw_snat"], "syntax": "block", "type": "object"}, {"aliases": ["interface ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:interface_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sli_to_slo_snat", "interface_ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/sli_to_slo_snat/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "X-example: \"\" description.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
