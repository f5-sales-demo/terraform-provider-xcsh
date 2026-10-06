---
page_title: "sli_to_slo_snat"
subcategory: "Networking"
description: "X-example: \"\" description."
xcsh_docs: {"aliases": ["sli to slo snat"], "body_bytes": 1337, "body_sha256": "sha256:f626365e04ce91146b4d83b77fd9417c98e162e348e852696acde6b1679a08f9", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:default_gw_snat", "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:interface_ip"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat", "parent_id": "xcsh-docs:resources:network_connector:reference", "path": "documentation/resources/network_connector/properties/sli_to_slo_snat/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0311230213310203-1002131223302201-1301011222023101-2212122300321021-2103313222121202-2010010213310323-3111100010201210-0122233201102021", "registry_path": "docs/guides/resources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sli_to_slo_snat"], "schema_version": 1, "sections": [{"aliases": ["sli to slo snat default gw snat"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:default_gw_snat", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sli_to_slo_snat", "default_gw_snat"], "syntax": "block", "type": "object"}, {"aliases": ["sli to slo snat interface ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat:interface_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sli_to_slo_snat", "interface_ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/sli_to_slo_snat/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "X-example: \"\" description.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

Additional upstream details:

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
