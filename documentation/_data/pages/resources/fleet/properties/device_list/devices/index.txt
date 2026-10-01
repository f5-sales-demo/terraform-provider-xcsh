---
page_title: "device_list.devices"
subcategory: ""
description: "device_list.devices for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 5568, "body_sha256": "sha256:f4e46235b46f3dd6912e5ae36613a501145722aae3f374c6bb84e5665a10eb6b", "child_ids": ["xcsh-docs:resources:fleet:properties:device_list:devices:network_device"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:device_list:devices", "parent_id": "xcsh-docs:resources:fleet:properties:device_list", "path": "documentation/resources/fleet/properties/device_list/devices/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["device_list", "devices"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/device_list/devices/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "device_list.devices for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
and has an corresponding interface/device.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Upstream description:

Defines ownership for a device.

Device owner is invalid Device is owned by VER pod. Usually it will be network interface device or
accelerator like crypto engine. Device is available to be owned by vK8s workload on the site, like
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

## Next pages

- [device_list.devices.network_device](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/device_list/devices/network_device/)
- [device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/device_list/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
