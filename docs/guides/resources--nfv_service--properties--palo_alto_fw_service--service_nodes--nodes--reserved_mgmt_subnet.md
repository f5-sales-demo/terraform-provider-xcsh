---
page_title: "palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet"
subcategory: ""
description: "palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1299, "body_sha256": "sha256:e355d329e099db211cb181092af1c7fd910c5f5bed48629ba4587f2dc4c022dd", "canonical_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:reserved_mgmt_subnet", "child_ids": [], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:reserved_mgmt_subnet", "parent_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "path": "docs/guides/resources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes--reserved_mgmt_subnet.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "reserved_mgmt_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/reserved_mgmt_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
