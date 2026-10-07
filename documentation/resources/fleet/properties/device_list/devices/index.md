---
page_title: "device_list.devices"
subcategory: ""
description: "Configuration for all devices in the fleet. Examples of devices are - network interfaces, cameras, scanners etc. Configuration a device is applied on VER node if the VER node is member of this fleet and has an corresponding interface/device. The mapping from device configured in fleet with interface/device in VER node"
xcsh_docs: {"aliases": ["device list devices"], "body_bytes": 5158, "body_sha256": "sha256:cd59470708b6f3c336c883d4ef9968b97ef33b46996c44b2ef446d8e962f89d4", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:device_list:devices:network_device"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:device_list:devices", "parent_id": "xcsh-docs:resources:fleet:properties:device_list", "path": "documentation/resources/fleet/properties/device_list/devices/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2332302030300023-0030012332113010-1332021002323222-3111212333201000-0011332200312220-0212130331121122-1202022210320310-2202333330020101", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["device_list", "devices"], "schema_version": 1, "sections": [{"aliases": ["device list devices name"], "anchor": "schema-device_list--devices--name", "description": "Name of the device including the unit number (e.g. Eth0 or disk1). The name must match name of device in host-OS of node.", "document_id": "xcsh-docs:resources:fleet:properties:device_list:devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["device_list", "devices", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["device list devices network device"], "anchor": "section", "description": "Represents physical network interface. The 'interface' reference points to a Network Interface object. Attributes such as Labels, MTU from Network Interface must be applied to the device. Device mapping to nodes A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An interface in", "document_id": "xcsh-docs:resources:fleet:properties:device_list:devices:network_device", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "device_list.devices.network_device:RequiredObjectAttributes:interface", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:device_list:devices:network_device:interface", "type": "requires"}], "schema_path": ["device_list", "devices", "network_device"], "syntax": "block", "type": "object"}, {"aliases": ["device list devices owner"], "anchor": "schema-device_list--devices--owner", "description": "Defines ownership for a device. Device owner is invalid Device is owned by VER pod. Usually it will be network interface device or accelerator like crypto engine. Device is available to be owned by vK8s workload on the site, like camera GPU etc. Device is not available to be owned by vK8s or VER. Can be exposed via", "document_id": "xcsh-docs:resources:fleet:properties:device_list:devices", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["DEVICE_OWNER_HOST", "DEVICE_OWNER_INVALID", "DEVICE_OWNER_VER", "DEVICE_OWNER_VK8S_WORK_LOAD"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["device_list", "devices", "owner"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/device_list/devices/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Configuration for all devices in the fleet. Examples of devices are - network interfaces, cameras, scanners etc. Configuration a device is applied on VER node if the VER node is member of this fleet and has an corresponding interface/device. The mapping from device configured in fleet with interface/device in VER node", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["fleetCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DEVICE_OWNER_HOST","DEVICE_OWNER_INVALID","DEVICE_OWNER_VER","DEVICE_OWNER_VK8S_WORK_LOAD"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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
