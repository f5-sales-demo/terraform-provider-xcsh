---
page_title: "ethernet_interface.no_ipv6_address"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["ethernet interface no ipv6 address"], "body_bytes": 1031, "body_sha256": "sha256:0b0c560aaf050b9a098cdc43498fcf7946eb5b8ef8e5f7b689bb5fe352d2c616", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:no_ipv6_address", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface", "path": "documentation/resources/network_interface/properties/ethernet_interface/no_ipv6_address/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1333032111011330-0031133113021320-3311010213022003-0330032310232013-0203001200201102-0021320311032100-3220302232131331-2200200132331213", "registry_path": "docs/guides/resources--network_interface--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ethernet_interface", "no_ipv6_address"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/no_ipv6_address/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.no_ipv6_address

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/)
- ethernet_interface.no_ipv6_address

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
no_ipv6_address = {}
```

This is an empty object or choice marker. It has no direct properties.
