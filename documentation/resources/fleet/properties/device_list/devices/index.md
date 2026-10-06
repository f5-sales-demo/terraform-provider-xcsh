---
page_title: "device_list.devices"
subcategory: ""
description: "Configuration for all devices in the fleet. Examples of devices are - network interfaces, cameras, scanners etc. Configuration a device is applied on VER node if the VER node is member of this fleet and has an corresponding interface/device. The mapping from device configured in fleet with interface/device in VER node"
xcsh_docs: {"aliases": ["device list devices"], "body_bytes": 4800, "body_sha256": "sha256:ca4bc35db5b0ceaa0a267ec9cf9a0d8c90b87aa0a5ba7682eab6cee35ab95dc9", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:device_list:devices:network_device"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:device_list:devices", "parent_id": "xcsh-docs:resources:fleet:properties:device_list", "path": "documentation/resources/fleet/properties/device_list/devices/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2332302030300023-0030012332113010-1332021002323222-3111212333201000-0011332200312220-0212130331121122-1202022210320310-2202333330020101", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["device_list", "devices"], "schema_version": 1, "sections": [{"aliases": ["device list devices name"], "anchor": "schema-device_list--devices--name", "description": "Name of the device including the unit number (e.g. Eth0 or disk1). The name must match name of device in host-OS of node.", "document_id": "xcsh-docs:resources:fleet:properties:device_list:devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["device_list", "devices", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["device list devices network device"], "anchor": "section", "description": "Represents physical network interface. The 'interface' reference points to a Network Interface object. Attributes such as Labels, MTU from Network Interface must be applied to the device. Device mapping to nodes A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An interface in", "document_id": "xcsh-docs:resources:fleet:properties:device_list:devices:network_device", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "device_list.devices.network_device:RequiredObjectAttributes:interface", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:device_list:devices:network_device:interface", "type": "requires"}], "schema_path": ["device_list", "devices", "network_device"], "syntax": "block", "type": "object"}, {"aliases": ["device list devices owner"], "anchor": "schema-device_list--devices--owner", "description": "Defines ownership for a device. Device owner is invalid Device is owned by VER pod. Usually it will be network interface device or accelerator like crypto engine. Device is available to be owned by vK8s workload on the site, like camera GPU etc. Device is not available to be owned by vK8s or VER. Can be exposed via", "document_id": "xcsh-docs:resources:fleet:properties:device_list:devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["device_list", "devices", "owner"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/device_list/devices/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration for all devices in the fleet. Examples of devices are - network interfaces, cameras, scanners etc. Configuration a device is applied on VER node if the VER node is member of this fleet and has an corresponding interface/device. The mapping from device configured in fleet with interface/device in VER node", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# device_list.devices

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/device_list/)
- device_list.devices

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Terraform syntax:

```terraform
devices {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-device_list--devices--name"></a>

### name property

Type: `"string"`. Optional.

Name of the device including the unit number (e.g. Eth0 or disk1). The name must match name of
device in host-OS of node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [network_device](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/device_list/devices/network_device/): complete subsection reference.

<a id="schema-device_list--devices--owner"></a>

### owner property

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DEVICE_OWNER_INVALID",
    "DEVICE_OWNER_VER",
    "DEVICE_OWNER_VK8S_WORK_LOAD",
    "DEVICE_OWNER_HOST"),
}
```

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
