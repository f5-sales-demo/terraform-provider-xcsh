---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-5b9b47ff4a22b36e0f3dfdde90952eaf8cd7ca8ea43357402eaceabfb926f7ca"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 634cc9462f15 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-8b2881b89d6b7161b0018eb07e4064f729d67e6e32a1041032b0e7bc09d2631b)
- stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site

<a id="canonical-2b6cd24f003dbd67f8f99db3104f6ccd03df11983a4a667c5f7520bea001cfcf"></a>

Type: `"single"`. Computed.

Defines a reference to a customer site virtual site along with network type where a load balancer
could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

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

<a id="canonical-59bb950267ecec0aa1435fc2e2bc44d0ac30e159c918e00a608d59706a27a678"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 634cc9462f15 / 3

<a id="canonical-4b3bde70f73c0bdfaab28cb1c56a6d4466880d797c4f6d1c7320434c6470a928"></a>

<a id="canonical-1e0ac5ddc94e6790552302a7059e7e77676b96276c1b521f0e9bb93df8191302"></a>

## network property — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 634cc9462f15 / 4

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--workload--reference--group-017.md#canonical-d46d9b79d99792e8af1d687c9943ce0e066c68c8dc4156be1bac07fdcdff77d4): complete subsection reference.

<a id="canonical-4940fa898d14a082a2e25d9802dbef40d4e8c0fe62a8d78e293bb25083191f06"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 634cc9462f15 / 5

- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--workload--reference--group-017.md#canonical-d46d9b79d99792e8af1d687c9943ce0e066c68c8dc4156be1bac07fdcdff77d4)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-8b2881b89d6b7161b0018eb07e4064f729d67e6e32a1041032b0e7bc09d2631b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d46d9b79d99792e8af1d687c9943ce0e066c68c8dc4156be1bac07fdcdff77d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd576827dcf4ecf1d8d5b1cb2114b5049d19890bf3bad71011c72e220daa94c4"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 8ac67c4527e7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-8b2881b89d6b7161b0018eb07e4064f729d67e6e32a1041032b0e7bc09d2631b)
- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--reference--group-016.md#canonical-8e28a9c5bcdf92777ecf5f5af305e6166d9df51941c591590a2eeb5362ba7d3d)
- stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-ee2cf56754ba792351fef1a66b2aa7511323494929de373700b482062b56afc8"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0e12edd7e9050a330469a41efed5b397ef3fdf7e41c5e43a4c69867b9e149f56"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 8ac67c4527e7 / 3

<a id="canonical-7443a7f85e413fef8182b5eacabb3ee6d0a76a7dbbd2ea0554312d51cc6d913d"></a>

<a id="canonical-7db2da2fd5768fc63a80898b6014108341a9504b3ce0e4217eb9f41977e149dc"></a>

## name property — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 8ac67c4527e7 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-908d8eab4fe2241d908550172a8d779018ec4b6dd3082aa7f956f364760b6dca"></a>

<a id="canonical-13403e36d50808963088452cf41fe818e4b849bebfddb421002faa5d40ae282f"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 8ac67c4527e7 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-a492e11ce97b86f72d71a6757e6b27b5d9d901763d36cdd8c2d4ac146668e7a3"></a>

<a id="canonical-1d20421166662d9572226fd53878bac743d1681e43211a3fe34bc3a4a927c463"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 8ac67c4527e7 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1c5bc529f324b901b403e73d16d2ba7e8684d7f50ef57d23cb76e644db343db7"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site / 8ac67c4527e7 / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--reference--group-016.md#canonical-8e28a9c5bcdf92777ecf5f5af305e6166d9df51941c591590a2eeb5362ba7d3d)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c58489be85f9fe13706372bcc4b0121fba090f246f0ca6c549e0ce1d95574573"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-714e605bcf8fbd51de09ab93676ec8200004fa33bcbe5e2708d44fca11bc9bd6"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 2ffc2201fadf / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-8b2881b89d6b7161b0018eb07e4064f729d67e6e32a1041032b0e7bc09d2631b)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service

<a id="canonical-61d2ed85b4196ffadc4b14e466eb19ce0d93e73f8828058d9af347c895f86af5"></a>

Type: `"single"`. Computed.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-49d67a7c6a74d6833482f0c6ec0adf04996e174d32f4f162b784a10ea29a277a"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 2ffc2201fadf / 3

- [site](data-sources--workload--reference--group-017.md#canonical-147071c554c8c4cc3986cf65eecf4fd9f45aaddd4a1426915ea376eb4e670395): complete subsection reference.

- [virtual_site](data-sources--workload--reference--group-017.md#canonical-aacd6a573deb67c9cfec9a82aefb0914d9420380bf395a94580b6ff3fa69a867): complete subsection reference.

<a id="canonical-ba3a61e1b2cb8449e9e05b8e2c9359f1a6e3b30df4459b647ff4658c9af19a0f"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 2ffc2201fadf / 4

- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site](data-sources--workload--reference--group-017.md#canonical-147071c554c8c4cc3986cf65eecf4fd9f45aaddd4a1426915ea376eb4e670395)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--workload--reference--group-017.md#canonical-aacd6a573deb67c9cfec9a82aefb0914d9420380bf395a94580b6ff3fa69a867)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-8b2881b89d6b7161b0018eb07e4064f729d67e6e32a1041032b0e7bc09d2631b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-147071c554c8c4cc3986cf65eecf4fd9f45aaddd4a1426915ea376eb4e670395"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4b9de88e16f0cf4ced00c1013e6498c891ad93ade5e2a64fe0ef68ad9aac036"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 34b8fb65671f / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-8b2881b89d6b7161b0018eb07e4064f729d67e6e32a1041032b0e7bc09d2631b)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-017.md#canonical-c58489be85f9fe13706372bcc4b0121fba090f246f0ca6c549e0ce1d95574573)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-9552f77978a1eaf12663bb65a58803bdf99f9404e0b28a0c6f5f48944411f908"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-10e4d3230f07276526b822d5b2e2e23392cb1cc05d883a527f7e97a77f570941"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 34b8fb65671f / 3

<a id="canonical-3ac88049edb8ec452c5bda17742532290754c77b9b822cf29873b89809f9743b"></a>

<a id="canonical-c31e65bab0b3b86a53f1e99b161906e3d6d6c54e87e65d9add3213002abd6835"></a>

## name property — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 34b8fb65671f / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-d343100bfedc28c167617eaa720ea84e93f7adfa121a9ae3c8dca38ec613c7c1"></a>

<a id="canonical-3460eb3c20459152b48cbbf3dfeb8cfc7fbd2131d280f5e90c322fefd4317136"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 34b8fb65671f / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-8534cc46df49070c157b4b305d5602e8ff356b0db02a70bee90e960dc665fc77"></a>

<a id="canonical-eabb5172dd279d9705ed387840fba2cf0037e9ed02c5e5062af3ebfb8ab2dca2"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 34b8fb65671f / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-617311e30198810e518c425016093274719679e8c12cc3c4ba51413f23df2ed5"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / 34b8fb65671f / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-017.md#canonical-c58489be85f9fe13706372bcc4b0121fba090f246f0ca6c549e0ce1d95574573)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-aacd6a573deb67c9cfec9a82aefb0914d9420380bf395a94580b6ff3fa69a867"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49c9066a9014b6277e6c3f88e2592d57d6c5ffb1e977219aee9878d22a8fd165"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / c9b3ad8a074f / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-8b2881b89d6b7161b0018eb07e4064f729d67e6e32a1041032b0e7bc09d2631b)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-017.md#canonical-c58489be85f9fe13706372bcc4b0121fba090f246f0ca6c549e0ce1d95574573)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-41f4cf3d6b82087f79f7c208e8f78942b1af034361160f136fe84b2b04d2d7b9"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-63730038d70b448659edfbd8d5af47016c1413fbbee4902b3455ee307657af21"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / c9b3ad8a074f / 3

<a id="canonical-82a708b1cbdf2a2711833150762c88deb01ffc5066bf664c1083db9c112e5349"></a>

<a id="canonical-0f1f05428c41fbaa89e411ac712f31a303b0d546817e6ba4a0fb3d750686d24e"></a>

## name property — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / c9b3ad8a074f / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-30ce149dbd2fc18524e2a9c25b2c460f18c9777819aaa45434911404c32bb96a"></a>

<a id="canonical-32ce2f9f13b754cd0a96422d3c36d0c141720d50e9bb56c9fe66bce6630e973c"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / c9b3ad8a074f / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-308328c6b92dfbe30775d9f052c4c9d9928675209da1da61d5d86a93ab52f206"></a>

<a id="canonical-f2b7847b2324491092f7166117d70ecae67e565b4adc62343b994089935c5a2f"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / c9b3ad8a074f / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-40af213e3219225798aa999588e7d170d0dd2bb2144e702d8c18ee29ee142253"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service / c9b3ad8a074f / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-017.md#canonical-c58489be85f9fe13706372bcc4b0121fba090f246f0ca6c549e0ce1d95574573)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3280561c530737514b080f9a139178e2c813f5e78a3dc3459c916a99e3cb4c4"></a>

## stateful_service.advertise_options.advertise_custom.ports — stateful_service.advertise_options.advertise_custom.ports / c670c502bd59 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- stateful_service.advertise_options.advertise_custom.ports

<a id="canonical-be34b1b4651966fa2c38d665edc6d8feca10ec3d3bf3b6394b77e44f5bc8ccdd"></a>

Type: `"list"`. Computed.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-09cbbb3b7bf727e5d98f60d42d7676e803ec32c5ede2d798f047d30c86b07f43"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports / c670c502bd59 / 3

- [http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21): complete subsection reference.

- [port](data-sources--workload--reference--group-020.md#canonical-c88ea326cf5dca54b8e30f8541332e128ae6cd0016217e6ac24bac80e9a9a351): complete subsection reference.

- [tcp_loadbalancer](data-sources--workload--reference--group-020.md#canonical-c919047ddcb76e16db814aaa9279d9c936af12cb1088b8da2562f50b3e853092): complete subsection reference.

<a id="canonical-42d39676153ed1efbb0478f5129a3eeec50de0c986b117f93ae5fd176de2e9ce"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports / c670c502bd59 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.port](data-sources--workload--reference--group-020.md#canonical-c88ea326cf5dca54b8e30f8541332e128ae6cd0016217e6ac24bac80e9a9a351)
- [stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer](data-sources--workload--reference--group-020.md#canonical-c919047ddcb76e16db814aaa9279d9c936af12cb1088b8da2562f50b3e853092)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ec1bdd1af0e83d34fae0ba6925321c15424d003490872d5cf3781e6b60404d3"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer / 1dd48738db78 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer

<a id="canonical-8a6b1693c0e042161a80206f8db51d116b307c1b98e6499bd664d851d4d97fcb"></a>

Type: `"single"`. Computed.

Configuration parameter for http loadbalancer.

Upstream description:

HTTP/HTTPS Load balancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

<a id="canonical-12f914c44b1fd8ece3a208896d48adafd39a4379ae8138f00ce222c69636e03c"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer / 1dd48738db78 / 3

- [default_route](data-sources--workload--reference--group-017.md#canonical-88abba157d963880e120946115e54ac76ceeb5972b2d0abb326dbb52a3b72d68): complete subsection reference.

<a id="canonical-3cca7e65958cfa8b82cfec3580cde2c552b5f5e3c12c9e90c37c9e11ea7b44f2"></a>

<a id="canonical-12f0a68fc4dc87a9dd4e5d61464bea8cb6a161649f06cf4956816312e29fb2b6"></a>

## domains property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer / 1dd48738db78 / 4

Type: `["list", "string"]`. Computed.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\`
and \`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
Wildcards must match a whole DNS label. E.g. \`\`\*.example.com\`\` and \*.bar.example.com are
valid, however \`\`\*bar.example.com\`\` or \`\`\*-bar.example.com\`\` is invalid

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [http](data-sources--workload--reference--group-017.md#canonical-2401b506233fc23606f8b5a674666a8be42edb2f6785581f6905b59145940933): complete subsection reference.

- [https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643): complete subsection reference.

- [https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50): complete subsection reference.

- [specific_routes](data-sources--workload--reference--group-019.md#canonical-47950d9bb3055b7b3775e6803f87ccac4dee024056bd2f8cd22fb9b784307644): complete subsection reference.

<a id="canonical-bc292df27b40f0f6ba6bc3d376958ab66448e7755b629cd24abc993322241d40"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer / 1dd48738db78 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-017.md#canonical-88abba157d963880e120946115e54ac76ceeb5972b2d0abb326dbb52a3b72d68)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http](data-sources--workload--reference--group-017.md#canonical-2401b506233fc23606f8b5a674666a8be42edb2f6785581f6905b59145940933)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-018.md#canonical-aad260f9f514def74f6d2d0b5de8b2a3348f266b26c8f24655dde803b35aec50)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-019.md#canonical-47950d9bb3055b7b3775e6803f87ccac4dee024056bd2f8cd22fb9b784307644)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-88abba157d963880e120946115e54ac76ceeb5972b2d0abb326dbb52a3b72d68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b332c1a3300e7dc6b227e1216a3b32fa9b8fff714e1430e01d52f6e12ac15c08"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 2461cc8bbfe9 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route

<a id="canonical-7d04e6d16e742f61c69a846df303dc260be553980cbf1da1cbb31394628fd4b7"></a>

Type: `"single"`. Computed.

Configuration parameter for default route.

Upstream description:

Default route matching all APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

<a id="canonical-811812b8e69d3b320234a8f36546987b68364cf6e472fd6421e85e7d9d2a1591"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 2461cc8bbfe9 / 3

- [auto_host_rewrite](data-sources--workload--reference--group-017.md#canonical-c833fbbe17b3ee0f2393d0a7603dd57e15e9c6eb7bbc70b32eecdc513aa745c2): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-017.md#canonical-df064198406be2a0de2f02b0a5c0fc9ddd1baadcdd38693499c5a3e2aaee7e73): complete subsection reference.

<a id="canonical-59f2891bb24dc5ef4518952405f0d7ee646d213b5dff6046ded838712288c0a6"></a>

<a id="canonical-68c341044abc638cd7366d1f04e8671d15c955a6369e64ed55d0511a54185c68"></a>

## host_rewrite property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 2461cc8bbfe9 / 4

Type: `"string"`. Computed.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-62f841af4056592945ecc7a5397e2b2a09385061c013f9790068966e54283606"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 2461cc8bbfe9 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite](data-sources--workload--reference--group-017.md#canonical-c833fbbe17b3ee0f2393d0a7603dd57e15e9c6eb7bbc70b32eecdc513aa745c2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite](data-sources--workload--reference--group-017.md#canonical-df064198406be2a0de2f02b0a5c0fc9ddd1baadcdd38693499c5a3e2aaee7e73)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c833fbbe17b3ee0f2393d0a7603dd57e15e9c6eb7bbc70b32eecdc513aa745c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cdd83031bb3967d4dcd7f41f96c52f7d2633697d74aea88f68632562209872e0"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 1ef5ca43f9e2 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-017.md#canonical-88abba157d963880e120946115e54ac76ceeb5972b2d0abb326dbb52a3b72d68)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-4debca50c47695f63699faa3367723744c32a71dd6190c30bd732aa9bb0a54b0"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-f5c3f9f749a20cc6be03c55cb236dddb840a90be0bde9533060f42c6b2a9c7af"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 1ef5ca43f9e2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-351c2843fb4cae6ff12adce15fc217e760a71ef0c7661555bc8f39eeac33a23b"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / 1ef5ca43f9e2 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-017.md#canonical-88abba157d963880e120946115e54ac76ceeb5972b2d0abb326dbb52a3b72d68)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-df064198406be2a0de2f02b0a5c0fc9ddd1baadcdd38693499c5a3e2aaee7e73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f517529b04e199495de87376c28e00a6b454cc8b8e24a69b630f337206c6b182"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / bddbdbbde724 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-017.md#canonical-88abba157d963880e120946115e54ac76ceeb5972b2d0abb326dbb52a3b72d68)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-059da5ca5778a20965e277e7437d71462fe35ccca78d62320bdeef2747ff0164"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-5be52a5a9d6b2594a43ca3a44db4833e5d2a80510cf8c4da00bcdf1909d093a3"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / bddbdbbde724 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8d7e5db26cf417fb2fb1092b189a47a00fc730ffbb62b6ac334624361dd51d4c"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.defa / bddbdbbde724 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-017.md#canonical-88abba157d963880e120946115e54ac76ceeb5972b2d0abb326dbb52a3b72d68)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-2401b506233fc23606f8b5a674666a8be42edb2f6785581f6905b59145940933"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37fc64c0970241743dbe981ba3d3abd4a6f31888640cdd768dbd937b81ef9c96"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5fc83971a9ae / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http

<a id="canonical-9da5a4d49e462f7917b139ad3610af94f2625bac0a4862ad4a0297f7cb917fd5"></a>

Type: `"single"`. Computed.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

<a id="canonical-89e8e5e880db1d4cac43f3ae3bf5d001ba8d72cf0a129b18c75ee010e2ab6bc9"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5fc83971a9ae / 3

<a id="canonical-105c457af6385d7b6edf15fe55853ab52e6fd82d1618ca9705d0d5618dfd1118"></a>

<a id="canonical-252bc08e1ceba21ed2e5395a4e91e272f8ce87dd1c344da65651dd59003d35c7"></a>

## dns_volterra_managed property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5fc83971a9ae / 4

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

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

<a id="canonical-16b6144e17986e110d23c820a20406428a407905e7e2d433380e1239079e02df"></a>

<a id="canonical-116fbb86e01922f57b1aa8ef7287b4a312c3d8f47f49920c1aeff4f0711c1228"></a>

## port property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5fc83971a9ae / 5

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3020cbce44ae3a7eb2a194be698b7faa635083c375381512cb4903285de1ec34"></a>

<a id="canonical-e23da4578ac19dbc61d20b0d9323e0c0dc6e11d7154d5a93dda889aac40f7a21"></a>

## port_ranges property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5fc83971a9ae / 6

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-ad6b617c0866ec7f3c98813e0d7a480e5f728c7cd6ea181bd2e1c9888ada0301"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5fc83971a9ae / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6aa8c6d42de8e18219419d5e3082be853e8fd26795fd2d8d1ec44992d69797c"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c0f04387bf74 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https

<a id="canonical-12cf21ec7ce2fbf49768d5c6dc08b8df8772d845be633c9a7aca527881dfc865"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

<a id="canonical-cc624de746a39383acf3a1e7780743188bd85c74471671f0f15cf0c51585f406"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c0f04387bf74 / 3

<a id="canonical-8b94c8cb3e94f340d70541f4effbece03a9a1625f3059341cd720673aa5320d7"></a>

<a id="canonical-8212a66624e27c9a93e887f3cb91bbf9e72d08a29f3751fd3b9a424e3b0cd9d3"></a>

## add_hsts property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c0f04387bf74 / 4

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header.

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

<a id="canonical-3f771d19578acfa11246251100b05bf291eea376b0f79679a078dcda2338bcdd"></a>

<a id="canonical-87dad5320c1530fea9f444245b6e755b77d04aad1d2af2dc7baba4b2b5a14f48"></a>

## append_server_name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c0f04387bf74 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--workload--reference--group-017.md#canonical-eab35555c85fc1ac02c8e5cae974cdcf9de6d7ca1e94e11b113dec43a0018e63): complete subsection reference.

<a id="canonical-62b3fb68b8a44d1304d177097b5efb4304936e64d2a7cd6138a64452a872cac7"></a>

<a id="canonical-ea615541fcdb313b243e2bcbcf58e957a8fb40b3cab197733dab0748a04d372a"></a>

## connection_idle_timeout property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c0f04387bf74 / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](data-sources--workload--reference--group-017.md#canonical-0bbb4c20eb37fdd9a6627c947f8d2538c3d86802a20d9aa04ddb085e0d6b8eb7): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-017.md#canonical-06053321e381e97aee9e3febf65001d3c4874b31309e26861319adc24b9e1ff0): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-017.md#canonical-efa526835c974df7bb1002be7676d493d1e495f07b83b2c66c457fca22fd6d02): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-017.md#canonical-87ae98bcdeb8efbe5e8bc884e464e7848f883a8b79a18477a19bb7963d0b2c08): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-017.md#canonical-58a8503f25407ccd30104faf4dc0fef4cc409f54733ecd128433ff1251b71fbd): complete subsection reference.

<a id="canonical-bbbea728ea385825ebd0025be523e98546f70160aee6c9017e386a6b1f07f2ce"></a>

<a id="canonical-ef09e5a142d15557e6b2d4766bef8eaffa424d66857b2d2d6f9a546bde602990"></a>

## http_redirect property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c0f04387bf74 / 7

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

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

- [non_default_loadbalancer](data-sources--workload--reference--group-017.md#canonical-af9a8f4cff7b1bf25d574c044831d0678a0ff6824bb27e0f122d6bc53bb3bcbb): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-017.md#canonical-e455a3bb5129453d4b86c3d9850fe1d04224e9b738019ceed1d809fbee7fbd32): complete subsection reference.

<a id="canonical-360e547ae00ffbbf09485a48e3dc2dc9294dd2c7515f14506ba3fb963d8ddb48"></a>

<a id="canonical-4f1d4c7059a5b5ad7aa5839055605dad7dddf890af62f5b41059f2f4d3a3910d"></a>

## port property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c0f04387bf74 / 8

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-63541911ec6f7ebade32423311e08a2e6d096d6fefbfc2915b3cba8953398ee7"></a>

<a id="canonical-2c3e5874873e13f3924870b6b6d189667d57cedd596e989cc794a8a72a8b3c06"></a>

## port_ranges property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c0f04387bf74 / 9

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-731087a66208c449a75905460a6d3bcfaf100351e0e8eb2238957561ed6cf36e"></a>

<a id="canonical-21efe164b148c02c32ad20b60b18641ccce86928d3882d8f51cd413c6da649d0"></a>

## server_name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c0f04387bf74 / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78): complete subsection reference.

- [tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2): complete subsection reference.

<a id="canonical-40030d27af4bcb305cbebab9b478c70ff69d007e407934b529cba603a0736166"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c0f04387bf74 / 11

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-017.md#canonical-eab35555c85fc1ac02c8e5cae974cdcf9de6d7ca1e94e11b113dec43a0018e63)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header](data-sources--workload--reference--group-017.md#canonical-0bbb4c20eb37fdd9a6627c947f8d2538c3d86802a20d9aa04ddb085e0d6b8eb7)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer](data-sources--workload--reference--group-017.md#canonical-06053321e381e97aee9e3febf65001d3c4874b31309e26861319adc24b9e1ff0)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize](data-sources--workload--reference--group-017.md#canonical-efa526835c974df7bb1002be7676d493d1e495f07b83b2c66c457fca22fd6d02)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize](data-sources--workload--reference--group-017.md#canonical-87ae98bcdeb8efbe5e8bc884e464e7848f883a8b79a18477a19bb7963d0b2c08)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-58a8503f25407ccd30104faf4dc0fef4cc409f54733ecd128433ff1251b71fbd)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer](data-sources--workload--reference--group-017.md#canonical-af9a8f4cff7b1bf25d574c044831d0678a0ff6824bb27e0f122d6bc53bb3bcbb)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through](data-sources--workload--reference--group-017.md#canonical-e455a3bb5129453d4b86c3d9850fe1d04224e9b738019ceed1d809fbee7fbd32)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-ecd89b3af405ecf7541968a7fceb933b59a697d8b94962ff85990a43286152b2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-eab35555c85fc1ac02c8e5cae974cdcf9de6d7ca1e94e11b113dec43a0018e63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e36681af2dc1916e298257ceca3bc3b0cc44077612da28583813c6640413d36e"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f70e490f497d / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-0561aa24f0f7b5a2d566070aa117bf084a5466abbd6defbdf7c37b6d34513501"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

<a id="canonical-54fa4ddbb74d2ab07a74b645d3db56279624de509d743d1cba6e03f2308fdfc1"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f70e490f497d / 3

- [default_coalescing](data-sources--workload--reference--group-017.md#canonical-f295b81bf32cf5ca342cae68f1282eae9907842d953a26101e7d0b5141dfcce2): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-017.md#canonical-493e5abde869cad76c89e178ce5f4bbf329f91922ea2ae01f42f85d13f082366): complete subsection reference.

<a id="canonical-420a84905a6082cc32d5f041a1ddd303fd4d321034b44f7b703ed86c2efb00d2"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f70e490f497d / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing](data-sources--workload--reference--group-017.md#canonical-f295b81bf32cf5ca342cae68f1282eae9907842d953a26101e7d0b5141dfcce2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing](data-sources--workload--reference--group-017.md#canonical-493e5abde869cad76c89e178ce5f4bbf329f91922ea2ae01f42f85d13f082366)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f295b81bf32cf5ca342cae68f1282eae9907842d953a26101e7d0b5141dfcce2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d66f403a5fcf79ea7fea2e844d78f74d441b5eabe7e97ab3e04a1eec9a1406e3"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a85bdf9869c3 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-017.md#canonical-eab35555c85fc1ac02c8e5cae974cdcf9de6d7ca1e94e11b113dec43a0018e63)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-c02d726ff7d29d9e501f019e93a5f16260012e40feeeba4a3368d5762d0c3018"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-26bc185a9751412500ad2bd5523c089b4345ade19f65cf3bd09a295253ddb779"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a85bdf9869c3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-320fa7a1df838c8391466b57ac4f55474a6c7be7f929fde7d8eef1f5e67c1a77"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a85bdf9869c3 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-017.md#canonical-eab35555c85fc1ac02c8e5cae974cdcf9de6d7ca1e94e11b113dec43a0018e63)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-493e5abde869cad76c89e178ce5f4bbf329f91922ea2ae01f42f85d13f082366"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2351e8087e0bce20f1f622b441289a21f8d8f3d962ac6d4b7b22f4fcce4ca56"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 416efa1362d6 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-017.md#canonical-eab35555c85fc1ac02c8e5cae974cdcf9de6d7ca1e94e11b113dec43a0018e63)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-5331d8857e9af2620b820b849fcf9fa5d47cfadc3fe103ba3fbffed58ef959a8"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-b573eb1d3d80ec34ee1d7703df06862945f13afba8c92341048295ec72b9eb65"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 416efa1362d6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-48cd6a02e877a39ba4ef6086d911d9b0c694e065e2d3b54a7a53200aba36c992"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 416efa1362d6 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-017.md#canonical-eab35555c85fc1ac02c8e5cae974cdcf9de6d7ca1e94e11b113dec43a0018e63)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-0bbb4c20eb37fdd9a6627c947f8d2538c3d86802a20d9aa04ddb085e0d6b8eb7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-804cd30fa090b863906d6b5ed83d0b4a58fd519f365e0a693c12681a7533234f"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c77afa1d3bd7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header

<a id="canonical-22616b09bc97808d88974c7f548a891986d3aa6086278487b580d759ad8cd2f1"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-f73574de7f32e1e653c468f9e97588ddb9a0c57ad1372d7c9098d15249e2db1a"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c77afa1d3bd7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-26509a00a1af57850f223b4821cdd3283c79dbf383ab73df7e251f32f32b4e83"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c77afa1d3bd7 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-06053321e381e97aee9e3febf65001d3c4874b31309e26861319adc24b9e1ff0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13a130ae1a17a45ee7d6a9723cd43f167d94809b0cbd26438acde8e427e43db6"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 24dfd4cec710 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-87ff249ba14eb44edee35c62b6baf34df6c1bcfeb40cb531490a1b7e3f076cb1"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

<a id="canonical-fe65868a0837dd544d12b88531d1998e7f93fc870bbfb66ec5740f83d45605f2"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 24dfd4cec710 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f9e82243cab8d98203682b45b363a0700b2e75b460641de34d1d3875864f63f0"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 24dfd4cec710 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-efa526835c974df7bb1002be7676d493d1e495f07b83b2c66c457fca22fd6d02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0de1d9ed075fee9671dcc5c04372b0d60432ffa7ac719b50f2bc4297555da0c0"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 24f15850cec8 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-dcfc7b3af91f7236557e53671c9ccafccacfa520971a65e4ac01b5aed7a91344"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-7ab41d594f5c71410e8d67d4fe49a263f4ffd9f249f30ec20269eb769c095ea4"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 24f15850cec8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0721c472b69d8f92ced983981174f036d5a8e12672e69e851b7a02bdb7efdeb2"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 24f15850cec8 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-87ae98bcdeb8efbe5e8bc884e464e7848f883a8b79a18477a19bb7963d0b2c08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a0fcfe3ba19c1ba9db9da1692b1283fa6c041b79235ad4b92c8543596f22f41"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 37cfed727524 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-e34989dee8ee6819c24b999cbd1e4728c36bbc47756ead59f4d212fc133564d1"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-529b5d492fff1f4224f81d0638642ab5790ece32144cfd744d78f6294b2fe10e"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 37cfed727524 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8ad63e785940297f488b052c0203bbafa19c447e1fd7235925900d3e10b8e923"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 37cfed727524 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-58a8503f25407ccd30104faf4dc0fef4cc409f54733ecd128433ff1251b71fbd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-566a51afdd12912d7b22475de3ff8156a7f7a68e0880c94977bd2770c48574c0"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a293642ec3d5 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-50c3a98b53d23291b7e5d50eaaa75853969916d5e48b15c120c13b99e33d325d"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

<a id="canonical-2afd15362241d81e6c7a87be35f74595f338f5de1f5d540f6d99d9c7753eeb28"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a293642ec3d5 / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-017.md#canonical-8cba06549036d57c8be2c79030e4ddb92a7391daa230f97a7cd0ccd2b6252e5d): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-017.md#canonical-582b6fa3f021dcecff5e3ab7fee75e3153d7f41ef88d740d287bddab47308d0c): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-017.md#canonical-823bb02db5e2465e23a0b739376b4b03ac4bfa817c72f09a428216c76e5d7938): complete subsection reference.

<a id="canonical-6b02c0184e997558215ee90ae6ae09b7f292feeca4c598bd84918c9ba6e0c0eb"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a293642ec3d5 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-017.md#canonical-8cba06549036d57c8be2c79030e4ddb92a7391daa230f97a7cd0ccd2b6252e5d)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-017.md#canonical-582b6fa3f021dcecff5e3ab7fee75e3153d7f41ef88d740d287bddab47308d0c)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-017.md#canonical-823bb02db5e2465e23a0b739376b4b03ac4bfa817c72f09a428216c76e5d7938)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-8cba06549036d57c8be2c79030e4ddb92a7391daa230f97a7cd0ccd2b6252e5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d244a03fd1d0cda998a9acbfa0b326da88816729f1a6558753193bd5945dc85"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 38fd32a4410c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-58a8503f25407ccd30104faf4dc0fef4cc409f54733ecd128433ff1251b71fbd)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-fbed4db99bf5eee1fd9d27bc87dcd9fbb5bdf3b0602b67b5de88e32da6fb1503"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

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

<a id="canonical-e271391abdf68f4644a531be8ff3a67be5c1e2af22852c7ad5fc7ff611cd6e6e"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 38fd32a4410c / 3

- [header_transformation](data-sources--workload--reference--group-017.md#canonical-a08f8cdd8b56e8ae65c0a8a440db41208d7601113867fe910a53e44f67145430): complete subsection reference.

<a id="canonical-cadd5ad1687f9745a7760da5cd4c470eb709c8453879563cdf17afd55579c8e0"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 38fd32a4410c / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-017.md#canonical-a08f8cdd8b56e8ae65c0a8a440db41208d7601113867fe910a53e44f67145430)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-58a8503f25407ccd30104faf4dc0fef4cc409f54733ecd128433ff1251b71fbd)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a08f8cdd8b56e8ae65c0a8a440db41208d7601113867fe910a53e44f67145430"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e60fd9f588273ec71923a4774ad6ca1cf4174c025e5751d3706b53ec16f67f22"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bef635822366 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-58a8503f25407ccd30104faf4dc0fef4cc409f54733ecd128433ff1251b71fbd)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-017.md#canonical-8cba06549036d57c8be2c79030e4ddb92a7391daa230f97a7cd0ccd2b6252e5d)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-06e0da2113bc74c6ddc374403b45d4ae4d196195659c6e2908f1bc19adada82b"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

<a id="canonical-66cda05d1042db23c61c3702ed73f1217d355605c86fcc867f8933421ab6469b"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bef635822366 / 3

- [default_header_transformation](data-sources--workload--reference--group-017.md#canonical-78f620037bce20caddcebbe5d515d8041f68232ce185098baafe8ac548c16829): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-017.md#canonical-139a2be8be5f9ba45a2138b4e838136406fbba7e797e2b68250e44016f8b85fe): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-017.md#canonical-9f874ff5ccca299d035a95b12be6dcb8ac14ab2609a4d7c0b3a2e0e8a6a83b2b): complete subsection reference.

<a id="canonical-5e75f781b75c422e9291b1228db7b68c023d58687e5ae9ccd1dd7d98db2465d2"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bef635822366 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--reference--group-017.md#canonical-78f620037bce20caddcebbe5d515d8041f68232ce185098baafe8ac548c16829)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--reference--group-017.md#canonical-139a2be8be5f9ba45a2138b4e838136406fbba7e797e2b68250e44016f8b85fe)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--reference--group-017.md#canonical-9f874ff5ccca299d035a95b12be6dcb8ac14ab2609a4d7c0b3a2e0e8a6a83b2b)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-017.md#canonical-8cba06549036d57c8be2c79030e4ddb92a7391daa230f97a7cd0ccd2b6252e5d)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-78f620037bce20caddcebbe5d515d8041f68232ce185098baafe8ac548c16829"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dba3633237eb8354f69aa4663dabdf8136e7d03b3c68bf34ccccf601b320642d"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f60dc827aa26 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-58a8503f25407ccd30104faf4dc0fef4cc409f54733ecd128433ff1251b71fbd)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-017.md#canonical-8cba06549036d57c8be2c79030e4ddb92a7391daa230f97a7cd0ccd2b6252e5d)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-017.md#canonical-a08f8cdd8b56e8ae65c0a8a440db41208d7601113867fe910a53e44f67145430)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-57ddfd026b406cb02738cf9e0da50fd6876c4ccd4f1ccef34ab27a4480889b96"></a>

Type: `["object", {}]`. Computed.

Use the platform's current default HTTP header transformation behavior.

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

<a id="canonical-1f1360e4523f14b4a5885b373f2f8018a66f175ecd61c8c81afc2164017df8ce"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f60dc827aa26 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65f47dc9c57a9c4a2bb0071203e09da69804fd485ee478adb833cf25b357ed6b"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f60dc827aa26 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-017.md#canonical-a08f8cdd8b56e8ae65c0a8a440db41208d7601113867fe910a53e44f67145430)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-139a2be8be5f9ba45a2138b4e838136406fbba7e797e2b68250e44016f8b85fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c424928d824cc144c2a810c496cf6b6ac8178050fa34f47cc4884104e770d66"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c9c850980a32 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-58a8503f25407ccd30104faf4dc0fef4cc409f54733ecd128433ff1251b71fbd)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-017.md#canonical-8cba06549036d57c8be2c79030e4ddb92a7391daa230f97a7cd0ccd2b6252e5d)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-017.md#canonical-a08f8cdd8b56e8ae65c0a8a440db41208d7601113867fe910a53e44f67145430)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-64c9ffbfc710539b1229814470cc954187e6e1353f0376507c43e9c0e6016525"></a>

Type: `["object", {}]`. Computed.

Preserve HTTP header-name case when upstream case must remain unchanged.

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

<a id="canonical-41c6f1e2cd6764e3d753cdbb65b303ad457d09ccf34c980cae62071e329cc3a2"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c9c850980a32 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0465715d73f647a574309c030eee7a87929c37cdda62b19ab37760e7ff2c1721"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c9c850980a32 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-017.md#canonical-a08f8cdd8b56e8ae65c0a8a440db41208d7601113867fe910a53e44f67145430)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9f874ff5ccca299d035a95b12be6dcb8ac14ab2609a4d7c0b3a2e0e8a6a83b2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42b9eec6d68b83912d5cdf00e5fb07e21e4e17aa7ebbe9872eda07c4547987bf"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bbc8d12405ea / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-58a8503f25407ccd30104faf4dc0fef4cc409f54733ecd128433ff1251b71fbd)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-017.md#canonical-8cba06549036d57c8be2c79030e4ddb92a7391daa230f97a7cd0ccd2b6252e5d)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-017.md#canonical-a08f8cdd8b56e8ae65c0a8a440db41208d7601113867fe910a53e44f67145430)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-f559903a021ff3cda36eb472f714324df164f9b88c68ccd87e717221d905882b"></a>

Type: `["object", {}]`. Computed.

Transform HTTP header names to proper case when explicit transformation is required.

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

<a id="canonical-b4625b782c2850cdce00f38c22c241be1edac947819743854ea764646e760fc2"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bbc8d12405ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ffd7ac156cca9631eac30fedcdb3168ba7a61f585b1eee4965d25c88fd3b929a"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bbc8d12405ea / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-017.md#canonical-a08f8cdd8b56e8ae65c0a8a440db41208d7601113867fe910a53e44f67145430)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-582b6fa3f021dcecff5e3ab7fee75e3153d7f41ef88d740d287bddab47308d0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c720421e49f69f436dcbbf71c8322e576b333145872f1c07758091167d4e668c"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2 — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 85f586725d52 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-58a8503f25407ccd30104faf4dc0fef4cc409f54733ecd128433ff1251b71fbd)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-37da564733c5a48b3fd014fd50aa11720dea6c8d8d1506c20859f78c97fba950"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

<a id="canonical-70454627a0c97dea98317c0d706be7ed5474237ebe9c4da4e1590c3a341dcd2e"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 85f586725d52 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-319821416600ac9d1681f061b1cf8f5f28f2e57b20f502e16750468e3da2eab4"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 85f586725d52 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-58a8503f25407ccd30104faf4dc0fef4cc409f54733ecd128433ff1251b71fbd)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-823bb02db5e2465e23a0b739376b4b03ac4bfa817c72f09a428216c76e5d7938"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-548a2a9b676218a3ade0f37ed56f505c4f168e428d4a60e230d2a2db7695a7e2"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5922c19e966d / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-58a8503f25407ccd30104faf4dc0fef4cc409f54733ecd128433ff1251b71fbd)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-f9842d40b63dac102a262f50f07e32efa070dd303ca53ec01f1e96ccc1b6dc6f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-99d7ee7f6799bc3f05029bba29f4961d7af52c04e1ed8b5001c856b86f610fe4"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5922c19e966d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-347d83021a4a4524055ecfbb0742297ccc6ea9d4ff256e709468b83dcc72910d"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5922c19e966d / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-58a8503f25407ccd30104faf4dc0fef4cc409f54733ecd128433ff1251b71fbd)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-af9a8f4cff7b1bf25d574c044831d0678a0ff6824bb27e0f122d6bc53bb3bcbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39e2c2d4d8a305898616913b2d1656e129f6c228b4ff5021602249156d1481e9"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f76b0a2f85c5 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-fc0df16bd62073cb820921c967451782d19c357a1734ef53e3f53df3058d2b62"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

<a id="canonical-2d86e1f9382455d94c254e637f8ebb9ccad124578a433556f31e25c65ea2d974"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f76b0a2f85c5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b2f464f110bce2bd8d5744996ed77de4e1705e4b3df1dfe249e4fcf94784eb1b"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f76b0a2f85c5 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e455a3bb5129453d4b86c3d9850fe1d04224e9b738019ceed1d809fbee7fbd32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f69c092657e836e173e0abef07499948d876a9cdecc27fa6a40970681866467"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 67523707c986 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through

<a id="canonical-b610ef5c64aac361ceb98b72397e7d2e2ce9ec0446ddac8e6b07fd071b3f6202"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

<a id="canonical-8644441817e08012def794dc29a03b6bd2e2f40692af0524292b6214bfdd98e0"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 67523707c986 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7e44de519863b682fc0185177f3bd45f1f42728eef03114087a61c7fc774a886"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 67523707c986 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40a641bab7957975781692b00dc4a7f2a69b6bb8b818da2df475621bf1c280d5"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 0c63b92a83e3 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params

<a id="canonical-ec84936a343878e7a300424868038487e679a84d9e687b514598467df5e6dec8"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-941deec4a6fb657697e79c8c81914c61c469301bc12e4c4cfac87c28cf7cbce1"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 0c63b92a83e3 / 3

- [certificates](data-sources--workload--reference--group-017.md#canonical-20e0f9df0807986f229ce22712184676eaf1d88e5370b061bf5ce86699626662): complete subsection reference.

- [no_mtls](data-sources--workload--reference--group-017.md#canonical-1ef341187594073186095cac3a00a7f76761f7a012522c30529bf3ccfba683da): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-017.md#canonical-aee05c2a3328d0b64ae539b71f0fc8df698eec53a968a2992a3f987429f0bc99): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-017.md#canonical-b7c1216887c2c56b699d1fcaef1d8773467395aa0f4bbb9adf802d0b0fe80ad3): complete subsection reference.

<a id="canonical-168e49a34dd44157e284705a0a2ebd10ab8e08453c6b6c410a4482fd267cd81e"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 0c63b92a83e3 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates](data-sources--workload--reference--group-017.md#canonical-20e0f9df0807986f229ce22712184676eaf1d88e5370b061bf5ce86699626662)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls](data-sources--workload--reference--group-017.md#canonical-1ef341187594073186095cac3a00a7f76761f7a012522c30529bf3ccfba683da)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-017.md#canonical-aee05c2a3328d0b64ae539b71f0fc8df698eec53a968a2992a3f987429f0bc99)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-017.md#canonical-b7c1216887c2c56b699d1fcaef1d8773467395aa0f4bbb9adf802d0b0fe80ad3)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-20e0f9df0807986f229ce22712184676eaf1d88e5370b061bf5ce86699626662"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4eb68ef82a728ae53d68bb66e7adaa413e722ad56edb0cb2e2205f3998683343"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bfe075398b32 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-598cbcccc009ae003b551e95d9b2856b560d4adfe2abcd2306b5c445be594f8c"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5b3d19b80eba69fa716d64cb20b13fa8049ffa6656ae376d9456ae807799cb95"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bfe075398b32 / 3

<a id="canonical-54963c26889b7a813727625e9e4e32e414ea11faf56d535b8e70afdc2409dd20"></a>

<a id="canonical-958383507375f4f9e36580dbfb8ffeef3b4d5a9b07c663ca722cf1863e4dea57"></a>

## name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bfe075398b32 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-4e94f03e0cc0768773ccb1f64ff283ba18af394b77723391798e97711ffbd911"></a>

<a id="canonical-07a2cc27a31f77e18cd47b18a454cdc37b5ee48d97733a33e7e32279cbb0e393"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bfe075398b32 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-9a427c4c1c8f8b4c69a93f0b6409345905ff71a9ed8f733a683799b679a632bc"></a>

<a id="canonical-40fc9ff74f63520d2a87f7a00a466d9f9f44b4cb867a18bc49c54bc2b42e92f0"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bfe075398b32 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-fa043c0fa8ade28f421081aa18cc095ed2acee41454e146c7c29918f3f53f552"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bfe075398b32 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-1ef341187594073186095cac3a00a7f76761f7a012522c30529bf3ccfba683da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc289ed117e1c6290e1dd888ab910ded9d602dd5aade47a3e17f3062a079a0cd"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 362ad582a16c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-29de3d0386b1c3383e64285392a7c464b68df574696380738845d99aff4d6838"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-3c28b04acac49146a723549ad9238728c576a760b0c72537279e936f0556b396"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 362ad582a16c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b35870c641e4f4a3541ba51c39d93da9ae7496de180641d01ebd28f6a621150d"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 362ad582a16c / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-aee05c2a3328d0b64ae539b71f0fc8df698eec53a968a2992a3f987429f0bc99"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d4b1440730c3a91916e81f38a13b9514a57c45defee3e7e1577b0c7573d5b01"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 057476f3a503 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-f28a1b10e6bcb7e1cd0e17e0d103e0937ab0624861cc66186ae6bc18c0419878"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-02edae5a162d451e1c7472ed92a7aa8a64e8188ee7360cfd363d92d33debf944"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 057476f3a503 / 3

- [custom_security](data-sources--workload--reference--group-017.md#canonical-59b22914efd9b7298230944f0b5fce2c028dc0e472e92e6413510f53c3fd1ccb): complete subsection reference.

- [default_security](data-sources--workload--reference--group-017.md#canonical-9be7d64271940a5580b49a85bab383072639c393b52fd3fd005b3868ecfa9134): complete subsection reference.

- [low_security](data-sources--workload--reference--group-017.md#canonical-0ecd0756d510f0a56f21dc8e97ceef269de8ed1e683b040ffb1c3155c9038967): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-017.md#canonical-578bbc1e657ca31b642d1c51aa8d8e253aeffbb386af2db53bc4fd5788ac8908): complete subsection reference.

<a id="canonical-d3ce5914fdfeb60ab4f5a765d87c4a0156f7d5038accbe28b76d4bbd288e6176"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 057476f3a503 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](data-sources--workload--reference--group-017.md#canonical-59b22914efd9b7298230944f0b5fce2c028dc0e472e92e6413510f53c3fd1ccb)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security](data-sources--workload--reference--group-017.md#canonical-9be7d64271940a5580b49a85bab383072639c393b52fd3fd005b3868ecfa9134)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security](data-sources--workload--reference--group-017.md#canonical-0ecd0756d510f0a56f21dc8e97ceef269de8ed1e683b040ffb1c3155c9038967)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](data-sources--workload--reference--group-017.md#canonical-578bbc1e657ca31b642d1c51aa8d8e253aeffbb386af2db53bc4fd5788ac8908)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-59b22914efd9b7298230944f0b5fce2c028dc0e472e92e6413510f53c3fd1ccb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-771d9c811feaace782074935a4715f4a626daef98e192942cf2df81c7cafd9cd"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8e4c9737da58 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-017.md#canonical-aee05c2a3328d0b64ae539b71f0fc8df698eec53a968a2992a3f987429f0bc99)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-2f861623e676473db46ef869393f5a3c578408dc7a59516c0ace1e5baae76dce"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-9556b33c88e794536867fe222f9f20b3238ce07c11a2587527b05705d4c29146"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8e4c9737da58 / 3

<a id="canonical-9a694c7c5eb5f03ffd6fe5d7ff6401899ce15659b80a97a2dedf45d4a685a218"></a>

<a id="canonical-320c6c7bd3026e73f6fd7e6e3e70e94a3fb6fe74666e5a84f35322c5485d82d3"></a>

## cipher_suites property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8e4c9737da58 / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c0ca484fed2ad066ddfb5cbae95a0ec15f74c2988a1a4be9e5f3b977ba7c12f8"></a>

<a id="canonical-80d17943e2517898280658fae3cbdcd14993c2d967b312b48c34805894880252"></a>

## max_version property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8e4c9737da58 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c1c96ca5290c8f60c146247cac8d8b7246f13cdff4c75c422ce914e669f31ea8"></a>

<a id="canonical-853f5fe7789bf47b96b4759b25fe8c969103fa51b92fc7eb5056da6ea8a94fbe"></a>

## min_version property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8e4c9737da58 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-30ecd762e7d9ae70b14f09876e4424aaa6e2ac60c8b9234f83a2115dd394bdbb"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8e4c9737da58 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-017.md#canonical-aee05c2a3328d0b64ae539b71f0fc8df698eec53a968a2992a3f987429f0bc99)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9be7d64271940a5580b49a85bab383072639c393b52fd3fd005b3868ecfa9134"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8a0eff0b2a6be9a8228a14d4ed439c518685653b90ad8fbd4b6e0d893f676d8"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 9ce4f9a88eda / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-017.md#canonical-aee05c2a3328d0b64ae539b71f0fc8df698eec53a968a2992a3f987429f0bc99)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-bc7ede91beec7889f0e801868c1a1ba19c0a20aa2f125e22dd5375a5d5fa6334"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-a6da43e464560ff2a37863ea9352474d04d8a41040fe4ff6d60c3272682931c4"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 9ce4f9a88eda / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-76351dfc1e97123e479925a7bc150849ea6196d9fed470a73951caa2c237781d"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 9ce4f9a88eda / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-017.md#canonical-aee05c2a3328d0b64ae539b71f0fc8df698eec53a968a2992a3f987429f0bc99)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-0ecd0756d510f0a56f21dc8e97ceef269de8ed1e683b040ffb1c3155c9038967"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2906d984a403c0d83cba90ade2d58e057b022fb2d73d9f5a5991007fb4b03e0d"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / eeec2da856e7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-017.md#canonical-aee05c2a3328d0b64ae539b71f0fc8df698eec53a968a2992a3f987429f0bc99)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-0e70b2bca0b2087983988c00bda5c66dbf3ffdc1d158f418afb4e51f91e02ac5"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-4498f45e254bfdf37bb68c3634f779d2caf01b698a904fc6c27ee1f373806f14"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / eeec2da856e7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f11515b24160122dabfcbe0ee9050c15c949dda902dfca1912e740888f84e122"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / eeec2da856e7 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-017.md#canonical-aee05c2a3328d0b64ae539b71f0fc8df698eec53a968a2992a3f987429f0bc99)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-578bbc1e657ca31b642d1c51aa8d8e253aeffbb386af2db53bc4fd5788ac8908"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02e6a73bb08d4c601a70f914e50c59eed173115f8d6ebd4613d241c69db0cbb0"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 99ee31dd5661 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-017.md#canonical-aee05c2a3328d0b64ae539b71f0fc8df698eec53a968a2992a3f987429f0bc99)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-1da5a2b34b45251968489a1e1f47b97a4ae439adf9fc216cb8aae6d60bb799f4"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-cb0822c91d48f1310e6bcd32bf31114787405f40709516b8fd76f0408f70086a"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 99ee31dd5661 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7bb0f15de5c2ceaae968a66f86d0e6bf0415d3e4651911ca4fe5c223f5757c22"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 99ee31dd5661 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-017.md#canonical-aee05c2a3328d0b64ae539b71f0fc8df698eec53a968a2992a3f987429f0bc99)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-b7c1216887c2c56b699d1fcaef1d8773467395aa0f4bbb9adf802d0b0fe80ad3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-302540f9538fe65a31e4e1592822033e9427e8cb46c2b6a51309ad1d224c3df0"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8a73b670f314 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-55964c8b5cc713561f2c955aa0e888d12b8c3c518a677f777bc053b52221a17b)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-07f03bf6539a934eaf606b59d1d5c0bfaa9674e2e9b9e2086f14e7dcdc590929)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-3bfc81105e68c8b128bdd2ea199c3e52c04a58309a8f95301d47b780a8333a21)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-4f97493dda7438b685cc41a364f2d80513cbe28302221a4a3cd45f0f11e8d643)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-017.md#canonical-84045b21c8578756bf25fa652058fae0c30afa065bb3e0c1132ec32e6269fc78)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-3929db8b6e21cf5830b3ec1694ed43a86689c3cff6177e133a8d4dd1a5af5dc0"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-427ac4b2f8f4a030b634b8c6574b10fab4328a56603ea05caa71f39b602e6e1f"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8a73b670f314 / 3

<a id="canonical-3af4ebe22b04327e68ebe842ef71f1e5cae7151943e1cd55efc7e3ebc22b2a32"></a>

<a id="canonical-512576d06cf582a96e5607403c0208c55bb97c4cc290c1e696e8a7e799639f8b"></a>

## client_certificate_optional property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8a73b670f314 / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](data-sources--workload--reference--group-018.md#canonical-1d67abd8348bfcdd4c0024b590371e2f6014232daceb4c0fb36a1d1dc84747cd): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-018.md#canonical-7bb8d8864d1806147ba0eb63643fcc179e30f75a656fd059d86d6db6c67c1663): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-018.md#canonical-a8e0718b1f2c2483f240bd02ff224a46e4f78e75db80065e7c9ad601944ea7c5): complete subsection reference.

<a id="canonical-0a1fa67c4617573a683a44c1c70c6cf6dc53c3ccdf1e88f9479e690d0c6bfedd"></a>

<a id="canonical-2b663e17c67bcd9d6123d4709aa0df8037cba60041133971cd6d5a05674ca5f7"></a>

## trusted_ca_url property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 8a73b670f314 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--workload--reference--group-018.md#canonical-3724c86831aa76f3960aa7df9df8c88832dba6d97564a1b378c8ae14a705e4e3): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-018.md#canonical-6a8b473d36dd8abdfe8b143b1fc1f75515626b43b7dc31fdf7221254ff06f9b6): complete subsection reference.
