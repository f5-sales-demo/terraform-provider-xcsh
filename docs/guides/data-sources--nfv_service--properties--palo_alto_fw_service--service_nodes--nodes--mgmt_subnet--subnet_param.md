---
page_title: "palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param"
subcategory: ""
description: "palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 2368, "body_sha256": "sha256:3b343cadd688882b8b3a297b4aa94f0a09103907bd9a42834f5ca0941ee88510", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet:subnet_param", "child_ids": [], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet:subnet_param", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes:mgmt_subnet", "path": "docs/guides/data-sources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes--mgmt_subnet--subnet_param.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes", "mgmt_subnet", "subnet_param"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/mgmt_subnet/subnet_param/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- [palo_alto_fw_service](data-sources--nfv_service--properties--palo_alto_fw_service.md)
- [palo_alto_fw_service.service_nodes](data-sources--nfv_service--properties--palo_alto_fw_service--service_nodes.md)
- [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes.md)
- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](data-sources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes--mgmt_subnet.md)
- palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param

<a id="section"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

## Direct properties

<a id="schema-palo_alto_fw_service--service_nodes--nodes--mgmt_subnet--subnet_param--ipv4"></a>

### ipv4 property

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

## Next pages

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](data-sources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes--mgmt_subnet.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
