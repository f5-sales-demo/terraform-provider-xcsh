---
page_title: "device_list.devices"
subcategory: ""
description: "Configuration for all devices in the fleet. Examples of devices are - network interfaces, cameras, scanners etc. Configuration a device is applied on VER node if the VER node is member of this fleet and has an corresponding interface/device. The mapping from device configured in fleet with interface/device in VER node"
xcsh_docs: {"aliases": ["device list devices"], "body_bytes": 4229, "body_sha256": "sha256:70b6597a7d1cfc5a16a8164423d3832037b2df4f5dbfa90650918391146feaf8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:device_list:devices:network_device"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:device_list:devices", "parent_id": "xcsh-docs:data-sources:fleet:properties:device_list", "path": "documentation/data-sources/fleet/properties/device_list/devices/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1133210223332000-2033333100133222-2032211012031033-1121222303321200-2020320210223013-3113313013313321-1123023331003311-0012002233233132", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["device_list", "devices"], "schema_version": 1, "sections": [{"aliases": ["device list devices name"], "anchor": "schema-device_list--devices--name", "description": "Name of the device including the unit number (e.g. Eth0 or disk1). The name must match name of device in host-OS of node.", "document_id": "xcsh-docs:data-sources:fleet:properties:device_list:devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["device_list", "devices", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["device list devices network device"], "anchor": "section", "description": "Represents physical network interface. The 'interface' reference points to a Network Interface object. Attributes such as Labels, MTU from Network Interface must be applied to the device. Device mapping to nodes A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An interface in", "document_id": "xcsh-docs:data-sources:fleet:properties:device_list:devices:network_device", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["device_list", "devices", "network_device"], "syntax": "attribute", "type": "object"}, {"aliases": ["device list devices owner"], "anchor": "schema-device_list--devices--owner", "description": "Defines ownership for a device. Device owner is invalid Device is owned by VER pod. Usually it will be network interface device or accelerator like crypto engine. Device is available to be owned by vK8s workload on the site, like camera GPU etc. Device is not available to be owned by vK8s or VER. Can be exposed via", "document_id": "xcsh-docs:data-sources:fleet:properties:device_list:devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["device_list", "devices", "owner"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/device_list/devices/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Configuration for all devices in the fleet. Examples of devices are - network interfaces, cameras, scanners etc. Configuration a device is applied on VER node if the VER node is member of this fleet and has an corresponding interface/device. The mapping from device configured in fleet with interface/device in VER node", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["fleetCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# device_list.devices

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/device_list/)
- device_list.devices

<a id="section"></a>

Type: `"list"`. Computed.

Configuration for all devices in the fleet. Examples of devices are - network interfaces, cameras,
scanners etc. Configuration a device is applied on VER node if the VER node is member of this fleet
and has an corresponding interface/device. The mapping from device configured in fleet with
interface/device in VER node depends on the type of device and is documented in device instance
specific sections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-device_list--devices--name"></a>

### name property

Type: `"string"`. Computed.

Name of the device including the unit number (e.g. Eth0 or disk1). The name must match name of
device in host-OS of node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [network_device](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/device_list/devices/network_device/): complete subsection reference.

<a id="schema-device_list--devices--owner"></a>

### owner property

Type: `"string"`. Computed.

\[Enum:
DEVICE\_OWNER\_INVALID|DEVICE\_OWNER\_VER|DEVICE\_OWNER\_VK8S\_WORK\_LOAD|DEVICE\_OWNER\_HOST\]
Defines ownership for a device. Device owner is invalid Device is owned by VER pod. Usually it will
be network interface device or accelerator like crypto engine. Possible values are
\`DEVICE\_OWNER\_INVALID\`, \`DEVICE\_OWNER\_VER\`, \`DEVICE\_OWNER\_VK8S\_WORK\_LOAD\`,
\`DEVICE\_OWNER\_HOST\`. Defaults to \`DEVICE\_OWNER\_INVALID\`.

Additional upstream details:

Defines ownership for a device. Device is available to be owned by vK8s workload on the site, like
camera GPU etc. Device is not available to be owned by vK8s or VER. Can be exposed via some other
service. Like TPM.

Receipt-pinned upstream constraints:

```json
{
  "default": "DEVICE_OWNER_INVALID",
  "enum": [
    "DEVICE_OWNER_INVALID",
    "DEVICE_OWNER_VER",
    "DEVICE_OWNER_VK8S_WORK_LOAD",
    "DEVICE_OWNER_HOST"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
