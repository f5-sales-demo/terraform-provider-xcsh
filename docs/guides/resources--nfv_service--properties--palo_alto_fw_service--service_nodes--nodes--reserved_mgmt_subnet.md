---
page_title: "palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet"
subcategory: ""
description: "palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1398, "body_sha256": "sha256:36c086eca35638fd02de6a57b2ac60a607a41c9aed9bc034a93f561c5a3181fe", "canonical_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:reserved_mgmt_subnet", "child_ids": [], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:reserved_mgmt_subnet", "parent_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "path": "docs/guides/resources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes--reserved_mgmt_subnet.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "reserved_mgmt_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/reserved_mgmt_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [palo_alto_fw_service](resources--nfv_service--properties--palo_alto_fw_service.md)
- [palo_alto_fw_service.service_nodes](resources--nfv_service--properties--palo_alto_fw_service--service_nodes.md)
- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes.md)
- palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reserved mgmt subnet.

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
reserved_mgmt_subnet = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
