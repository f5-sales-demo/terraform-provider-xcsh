---
page_title: "dedicated_interface.monitor"
subcategory: ""
description: "Link Quality Monitoring configuration for a network interface."
xcsh_docs: {"aliases": ["dedicated interface monitor"], "body_bytes": 1258, "body_sha256": "sha256:88dff1e1d086ea5aab9fab73b0a7afb0932c23e03cb2b5345811e7d74e61f771", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:monitor", "parent_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface", "path": "documentation/resources/network_interface/properties/dedicated_interface/monitor/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3230321221323212-3020031231321210-2321212001020302-3332321332202330-0202313030101032-3113030003102211-1301123123303333-1031203002101112", "registry_path": "docs/guides/resources--network_interface--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dedicated_interface", "monitor"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/dedicated_interface/monitor/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Link Quality Monitoring configuration for a network interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dedicated_interface.monitor

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/)
- dedicated_interface.monitor

<a id="section"></a>

Type: `["object", {}]`. Optional.

Link Quality Monitoring configuration for a network interface.

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
monitor = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
