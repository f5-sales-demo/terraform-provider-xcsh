---
page_title: "palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["palo alto fw service service nodes nodes reserved mgmt subnet"], "body_bytes": 1443, "body_sha256": "sha256:73cc6add876601421910e00821c59833e16023e4553612e3429057d4d28131d9", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:reserved_mgmt_subnet", "parent_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "path": "documentation/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/reserved_mgmt_subnet/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2010020001033131-3003213032203223-0321221213303100-2233121110220121-0101223323203303-3231000132311120-3121300001230222-2303032301221021", "registry_path": "docs/guides/resources--nfv_service--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "reserved_mgmt_subnet"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/reserved_mgmt_subnet/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/)
- [palo_alto_fw_service.service_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/)
- [palo_alto_fw_service.service_nodes.nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/)
- palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reserved mgmt subnet.

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
reserved_mgmt_subnet = {}
```

This is an empty object or choice marker. It has no direct properties.
