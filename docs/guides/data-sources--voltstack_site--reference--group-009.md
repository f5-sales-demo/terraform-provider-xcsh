---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-92cbc03ec4bbecc82e46f92f0731faa5a316356b13ead018901d8d0d2db87920"></a>

## local_control_plane.bgp_config.peers.bfd_disabled — local_control_plane.bgp_config.peers.bfd_disabled / 85ad8ca480de / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- local_control_plane.bgp_config.peers.bfd_disabled

<a id="canonical-dfe18c0dbdda295cb7c760d309be4074cead2b7571813f5217b7130b6f7b197d"></a>

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

<a id="canonical-2d7dc38a2eeaefd3251483f094bcbfa44a4665371e9e0bb4aff931ee027f438e"></a>

## Direct properties — local_control_plane.bgp_config.peers.bfd_disabled / 85ad8ca480de / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c0a8200ed33c889e544b38e1c0d537ff75755f893a708eb4461e6550f4bf4940"></a>

## Next pages — local_control_plane.bgp_config.peers.bfd_disabled / 85ad8ca480de / 4

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-d38b3bf61cffa283fe861502ecf76d8a822b8b249062a5307f6883aa8a6fc0c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3a3a120c93599c5ef3d355793933062d4afaea83856c7301e5ab55cb14b0b84"></a>

## local_control_plane.bgp_config.peers.bfd_enabled — local_control_plane.bgp_config.peers.bfd_enabled / 5e9e2c255cda / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- local_control_plane.bgp_config.peers.bfd_enabled

<a id="canonical-0c3a5a81a53688f2ef230da894ffd7046eebd12d9a023564975f2b3a388b2363"></a>

Type: `"single"`. Computed.

BFD. BFD parameters.

Upstream description:

BFD parameters.

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

<a id="canonical-7936b4bc72b866dfa2896048263eab0250f1e6aeb17abc7ae7cc36c10e2b7b4e"></a>

## Direct properties — local_control_plane.bgp_config.peers.bfd_enabled / 5e9e2c255cda / 3

<a id="canonical-8482efa9870cc1627f3e6eb79802c3caec0bcfe876ad7bfd8305533611231538"></a>

<a id="canonical-bd0cf3b4f359a11fa6d35ecf21e748bc4876903819c17b2fac7abb629d6a8c05"></a>

## multiplier property — local_control_plane.bgp_config.peers.bfd_enabled / 5e9e2c255cda / 4

Type: `"number"`. Computed.

Specify Number of missed packets to bring session down'.

Upstream description:

Specify Number of missed packets to bring session down"

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-6c6531c1894ab346f5e2e6a0f5c18305dd9d181de6a1534bff5776d58951ac75"></a>

<a id="canonical-edce16eb882a20261c6e1fca25e061d8a1a9946d6494f7942b5be8e557d19d15"></a>

## receive_interval_milliseconds property — local_control_plane.bgp_config.peers.bfd_enabled / 5e9e2c255cda / 5

Type: `"number"`. Computed.

BFD receive interval timer, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-3af7b020bc29a78d4e08771075d656d571491c851b423058a0704985ab8f681e"></a>

<a id="canonical-9aa6e59d7dbc3bfe8cdee1ce2eafce38172fcd852189251f7a8e19d1b2468e14"></a>

## transmit_interval_milliseconds property — local_control_plane.bgp_config.peers.bfd_enabled / 5e9e2c255cda / 6

Type: `"number"`. Computed.

BFD transmit interval timer, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-e7ca53cfc2b2e890e0e50867a222c4f0dd1e100ed6976d265476e36eb92e74c8"></a>

## Next pages — local_control_plane.bgp_config.peers.bfd_enabled / 5e9e2c255cda / 7

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-59786833c29963a28df297bc9315780f00cd42f0f2f2cb0dad4c4a4f6a2ab934"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5be487b7bdf717b0e90e67da7a82f9e91a2f8007982cf3303b1c5fd5a3d9c73b"></a>

## local_control_plane.bgp_config.peers.disable_spec — local_control_plane.bgp_config.peers.disable_spec / 8746364ae62a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- local_control_plane.bgp_config.peers.disable_spec

<a id="canonical-cc3d821cfe1f0705bc18c3f30447d4593fe937eccf3ea2e0a97d32c31d4614cd"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-6b455ac1a734c0f52b39f20b298e88fb039ca65c4f038966c8ef3d5c025f8136"></a>

## Direct properties — local_control_plane.bgp_config.peers.disable_spec / 8746364ae62a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1e9aee7985022a139a562169fe1423ad170d3ffd6f8cfabcaf05c822a008d89a"></a>

## Next pages — local_control_plane.bgp_config.peers.disable_spec / 8746364ae62a / 4

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-326e52ac0c63ea0a546188ac2b12d39b1cc939e4710eba694dd4982cd0ea632b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00a57696c3d525c9caba57c141cd93fcce29f3bfc4b53d6d448a276efc15a4e9"></a>

## local_control_plane.bgp_config.peers.ebgp_multihop_disabled — local_control_plane.bgp_config.peers.ebgp_multihop_disabled / 73d4e37aee87 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- local_control_plane.bgp_config.peers.ebgp_multihop_disabled

<a id="canonical-272e2459e50e26cd6ebf5fcc8913414f25d6d800f64aceeacfb4c7555fb04779"></a>

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

<a id="canonical-20c56e462e1525637d86b8889b5fa548aa2907f2f78c4da0addc0af9193a0e0c"></a>

## Direct properties — local_control_plane.bgp_config.peers.ebgp_multihop_disabled / 73d4e37aee87 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d85965cf881f7ce8fb391f4ba71541f81af48a8a2e71cc7faf3f068118fe3706"></a>

## Next pages — local_control_plane.bgp_config.peers.ebgp_multihop_disabled / 73d4e37aee87 / 4

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-a337775efcbb590b4ee540d161def71e5f49454c17ec2d7e123b955702d853b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4256a986ccbf58aa53556e055bf011d1299734616c7e64db3eca76d6759d140f"></a>

## local_control_plane.bgp_config.peers.ebgp_multihop_enabled — local_control_plane.bgp_config.peers.ebgp_multihop_enabled / 3ce65e0639dd / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- local_control_plane.bgp_config.peers.ebgp_multihop_enabled

<a id="canonical-fbcc60907be413cd2d71ccaf82c50a34b44819256034c6a1f235f1de46f8f4b3"></a>

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

<a id="canonical-ee5b8e5ba0be64b79c5b53eeda4a0f9b1a4fc4f6cef657136d81141ce8380bdc"></a>

## Direct properties — local_control_plane.bgp_config.peers.ebgp_multihop_enabled / 3ce65e0639dd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-52a40b853f2c1c7e2e23641a7c030a72329cf2f0ab848ed0e967cc8b96bdcaa3"></a>

## Next pages — local_control_plane.bgp_config.peers.ebgp_multihop_enabled / 3ce65e0639dd / 4

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04f6636438a347d7a7a2e9de126e27ad6f32381f2262f2bd87e2ec0e9fa8102d"></a>

## local_control_plane.bgp_config.peers.external — local_control_plane.bgp_config.peers.external / 52e835399a99 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- local_control_plane.bgp_config.peers.external

<a id="canonical-8ff8aeb56abe03edfb5d022211415a8cbbd159bf43cfde54290300a7d3cde044"></a>

Type: `"single"`. Computed.

External BGP Peer. External BGP Peer parameters.

Upstream description:

External BGP Peer parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"address\",\"default_gateway\",\"disable\",\"external_connector\",\"from_site\",\"subnet_begin_offset\",\"subnet_end_offset\"]",
  "x-ves-oneof-field-address_choice_v6": "[\"address_ipv6\",\"default_gateway_v6\",\"disable_v6\",\"from_site_v6\",\"subnet_begin_offset_v6\",\"subnet_end_offset_v6\"]",
  "x-ves-oneof-field-auth_choice": "[\"md5_auth_key\",\"no_authentication\"]",
  "x-ves-oneof-field-interface_choice": "[\"interface\",\"interface_list\"]"
}
```

<a id="canonical-5a4337bc0c47a3e9575257a72597702f08c6b1fce106754e1c9b33821a9696d4"></a>

## Direct properties — local_control_plane.bgp_config.peers.external / 52e835399a99 / 3

<a id="canonical-dc5ee62da9d7af6a77381a0e0f3fd71f0e1287474c10c40caa595f2d55a50439"></a>

<a id="canonical-e3ec8cbebb1b3fe553905ea778c8dc0a3e47d34d30e2765aa185767cadd5e7fb"></a>

## address property — local_control_plane.bgp_config.peers.external / 52e835399a99 / 4

Type: `"string"`. Computed.

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

Upstream description:

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-67bdf4dab58b5e38f5d1f402bdb16637bbd58cfd1dccbce69c3db22fac287408"></a>

<a id="canonical-730c5ba9c573944e3ddf8ea05d33c66661a63996866f4876d849023404908640"></a>

## address_ipv6 property — local_control_plane.bgp_config.peers.external / 52e835399a99 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

Upstream description:

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-78aa08da48124d52c2123d312d5a710b909e75774e0de8887fd184a5b49b547b"></a>

<a id="canonical-13a5d35e26309cca10ba50d91219faed4eee98f3ec196afc1735ad181a838370"></a>

## asn property — local_control_plane.bgp_config.peers.external / 52e835399a99 / 6

Type: `"number"`. Computed.

ASN. Autonomous System Number for BGP peer.

Upstream description:

Autonomous System Number for BGP peer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [default_gateway](data-sources--voltstack_site--reference--group-009.md#canonical-bbb1d31367e6d174d90478d5926654a404183a35393fc26d0c68687cfadbea80): complete subsection reference.

- [default_gateway_v6](data-sources--voltstack_site--reference--group-009.md#canonical-eb54567d561c3e2d30cde368c1b529a2d736cb907c6b05810655141dd55e7098): complete subsection reference.

- [disable_spec](data-sources--voltstack_site--reference--group-009.md#canonical-659248fa424330d25f089d8bb1b84bc89dd32e9c3fec6f196adb2524001b292c): complete subsection reference.

- [disable_v6](data-sources--voltstack_site--reference--group-009.md#canonical-aa50ccebfb33b5caa7af59c7475596ae4e4ee92590f1066fd516b26e54442147): complete subsection reference.

- [external_connector](data-sources--voltstack_site--reference--group-009.md#canonical-0ba138e00c148b5d750ebb330a23f307511e9949ff112b1a8005158980c7712a): complete subsection reference.

- [family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-e90ff08ae2d6ac3dff1c616ef4eca5af5d75b10fc4923fcf4448ada606e2972f): complete subsection reference.

- [from_site](data-sources--voltstack_site--reference--group-009.md#canonical-618446a18ab61ec08ea6049f28686e2d4e8a0f035ec0f9cb7e6aec45a337cf3c): complete subsection reference.

- [from_site_v6](data-sources--voltstack_site--reference--group-009.md#canonical-97edea0468c34f5dfb1e511f99a0f0f98887bb255c153ec220a0dd8fbbc11edb): complete subsection reference.

- [interface](data-sources--voltstack_site--reference--group-009.md#canonical-3e4587b1b331c8080af3b4d8bbccd2562f23a1c017c7aff43f2d5ac6b70aaad4): complete subsection reference.

- [interface_list](data-sources--voltstack_site--reference--group-009.md#canonical-116ee59de0c582f666df0dc817e672bd50576a5e001480acbcb410c18e6a45a3): complete subsection reference.

<a id="canonical-e4a0a2b5be998aa19f5c799fd80d1b48048ed4d1fa4425f3022c12a41107a577"></a>

<a id="canonical-7660b9d938893bfee59a9c10d8f2aee9b30560dd5dc2285d0cc3b45d756f2a03"></a>

## md5_auth_key property — local_control_plane.bgp_config.peers.external / 52e835399a99 / 7

Type: `"string"`. Computed.

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385).

Upstream description:

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [no_authentication](data-sources--voltstack_site--reference--group-009.md#canonical-b82a22a43e45020b086a1b1c57fa9d6d9b202de7a1df57bfc171be28380834c0): complete subsection reference.

<a id="canonical-a55f4670e4187e4798639f3a0b4a0861ac44fe2d0a3ae7c30e3f8857504a6d8c"></a>

<a id="canonical-41840d5ca2dff66b59b30707488a3bc9aae03f16378e09b7be709f22dafa0065"></a>

## port property — local_control_plane.bgp_config.peers.external / 52e835399a99 / 8

Type: `"number"`. Computed.

Peer Port. Peer TCP port number.

Upstream description:

Peer TCP port number.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-c86931cffac190788181645b38bafed578104b858107b2939755ab2a213f608a"></a>

<a id="canonical-9e295ac2df8d39f049c70a745a821ee6c50a4043ff7bf38a534c4cc57a62e399"></a>

## subnet_begin_offset property — local_control_plane.bgp_config.peers.external / 52e835399a99 / 9

Type: `"number"`. Computed.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

Upstream description:

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-699f188d2a2036da5fd7648c06bcf014c8b94fe1e555904e000c5e5c4625f653"></a>

<a id="canonical-52335abfd877a962bfb38ed4ecdad8f770ddd46879a5724e9b57553ab86b72dc"></a>

## subnet_begin_offset_v6 property — local_control_plane.bgp_config.peers.external / 52e835399a99 / 10

Type: `"number"`. Computed.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

Upstream description:

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-1043bde9154d2d867ca271936e73ce9a446871a71a9435ba840d10b3006bdc0b"></a>

<a id="canonical-04ba1f3dfb8fe6565b5dbfd43adf9d0d32ed6759b59775aec25623c4e919da2c"></a>

## subnet_end_offset property — local_control_plane.bgp_config.peers.external / 52e835399a99 / 11

Type: `"number"`. Computed.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

Upstream description:

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-48eda94e6755fb499d17b9349ae49c6f179f9209c67630875ea4621766b5cdc4"></a>

<a id="canonical-0160fd1c4a75958ce0f7a4e4995ffe574fbfde4c8a6bd6b6b416a1ef753a153d"></a>

## subnet_end_offset_v6 property — local_control_plane.bgp_config.peers.external / 52e835399a99 / 12

Type: `"number"`. Computed.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

Upstream description:

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-c43102a11f759b46c705320570a0344d179fa296b0be8706c70a483a09e70b8f"></a>

## Next pages — local_control_plane.bgp_config.peers.external / 52e835399a99 / 13

- [local_control_plane.bgp_config.peers.external.default_gateway](data-sources--voltstack_site--reference--group-009.md#canonical-bbb1d31367e6d174d90478d5926654a404183a35393fc26d0c68687cfadbea80)
- [local_control_plane.bgp_config.peers.external.default_gateway_v6](data-sources--voltstack_site--reference--group-009.md#canonical-eb54567d561c3e2d30cde368c1b529a2d736cb907c6b05810655141dd55e7098)
- [local_control_plane.bgp_config.peers.external.disable_spec](data-sources--voltstack_site--reference--group-009.md#canonical-659248fa424330d25f089d8bb1b84bc89dd32e9c3fec6f196adb2524001b292c)
- [local_control_plane.bgp_config.peers.external.disable_v6](data-sources--voltstack_site--reference--group-009.md#canonical-aa50ccebfb33b5caa7af59c7475596ae4e4ee92590f1066fd516b26e54442147)
- [local_control_plane.bgp_config.peers.external.external_connector](data-sources--voltstack_site--reference--group-009.md#canonical-0ba138e00c148b5d750ebb330a23f307511e9949ff112b1a8005158980c7712a)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-e90ff08ae2d6ac3dff1c616ef4eca5af5d75b10fc4923fcf4448ada606e2972f)
- [local_control_plane.bgp_config.peers.external.from_site](data-sources--voltstack_site--reference--group-009.md#canonical-618446a18ab61ec08ea6049f28686e2d4e8a0f035ec0f9cb7e6aec45a337cf3c)
- [local_control_plane.bgp_config.peers.external.from_site_v6](data-sources--voltstack_site--reference--group-009.md#canonical-97edea0468c34f5dfb1e511f99a0f0f98887bb255c153ec220a0dd8fbbc11edb)
- [local_control_plane.bgp_config.peers.external.interface](data-sources--voltstack_site--reference--group-009.md#canonical-3e4587b1b331c8080af3b4d8bbccd2562f23a1c017c7aff43f2d5ac6b70aaad4)
- [local_control_plane.bgp_config.peers.external.interface_list](data-sources--voltstack_site--reference--group-009.md#canonical-116ee59de0c582f666df0dc817e672bd50576a5e001480acbcb410c18e6a45a3)
- [local_control_plane.bgp_config.peers.external.no_authentication](data-sources--voltstack_site--reference--group-009.md#canonical-b82a22a43e45020b086a1b1c57fa9d6d9b202de7a1df57bfc171be28380834c0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-bbb1d31367e6d174d90478d5926654a404183a35393fc26d0c68687cfadbea80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b26cbec509000dc0301d9b3928f590a7fa13606b586400d0fd07374a2ec55ec"></a>

## local_control_plane.bgp_config.peers.external.default_gateway — local_control_plane.bgp_config.peers.external.default_gateway / c03399de59c4 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- local_control_plane.bgp_config.peers.external.default_gateway

<a id="canonical-89854a99aa90a66a32804a2f59cea68d9271e421376105ee890419d3c81ec815"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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

<a id="canonical-de0437203a452c38fdb31b7085d264aee731df4835fe478754621b9861f0c09f"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.default_gateway / c03399de59c4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6c2afdb3d895f4314f2b705b149f0d5e568686dc4110b90c25c5251cdb42e95c"></a>

## Next pages — local_control_plane.bgp_config.peers.external.default_gateway / c03399de59c4 / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-eb54567d561c3e2d30cde368c1b529a2d736cb907c6b05810655141dd55e7098"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92991ddd9576cc573bac4ea9c990e5da6c182f0991c50f9ca3e21d94f3139f21"></a>

## local_control_plane.bgp_config.peers.external.default_gateway_v6 — local_control_plane.bgp_config.peers.external.default_gateway_v6 / 68b562354e3c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- local_control_plane.bgp_config.peers.external.default_gateway_v6

<a id="canonical-d6a655a192cf7dc9fda1a16dba6b764094522154c49b9c06a37618063df8f407"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway v6.

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

<a id="canonical-e2dfd08aeed6a03ac2a274397818a1f1a551f2a92ba0cef36500ecd8e8686097"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.default_gateway_v6 / 68b562354e3c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-14f7e544919ec6fa7a18c111a72262844ad27c515683491086ad9328d5092516"></a>

## Next pages — local_control_plane.bgp_config.peers.external.default_gateway_v6 / 68b562354e3c / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-659248fa424330d25f089d8bb1b84bc89dd32e9c3fec6f196adb2524001b292c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43badb3f8d374052624ddbb25cbc21f13c0691bc50c8a07a0b5f7a1a0bd53b1c"></a>

## local_control_plane.bgp_config.peers.external.disable_spec — local_control_plane.bgp_config.peers.external.disable_spec / b489448b6791 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- local_control_plane.bgp_config.peers.external.disable_spec

<a id="canonical-88a382d53828b217ac2171c5daed438cedc8d359de66f42598f08f42dde28536"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-36db45058a004065ae1a0f3ea4bf7eba89380d7a7a0bcc04518a03d09d6b6d1b"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.disable_spec / b489448b6791 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a7d2fd1758218d1cb9e6485eb30b00bfed95c24bcaa1faf66b5fdb797d538bc7"></a>

## Next pages — local_control_plane.bgp_config.peers.external.disable_spec / b489448b6791 / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-aa50ccebfb33b5caa7af59c7475596ae4e4ee92590f1066fd516b26e54442147"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ee9b2d7f26e61d3902cbf0c7c08db3ac1c61eee26df6823955e94175c010051"></a>

## local_control_plane.bgp_config.peers.external.disable_v6 — local_control_plane.bgp_config.peers.external.disable_v6 / 39713b56a299 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- local_control_plane.bgp_config.peers.external.disable_v6

<a id="canonical-579904b34f1288704fe43cca66b875d8b0c6b872ba68ab0f92a0556a5c6c7599"></a>

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

<a id="canonical-d7c4fdd82d983d0f57973c4f6966cb0d2353240effce5c4b5e3e332ebb453b25"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.disable_v6 / 39713b56a299 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a07197ac2fdc759e84708a4355e6dc830f43782f44736c7a1c975eaa87e860a2"></a>

## Next pages — local_control_plane.bgp_config.peers.external.disable_v6 / 39713b56a299 / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-0ba138e00c148b5d750ebb330a23f307511e9949ff112b1a8005158980c7712a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c1690036024ca21d046043167ef4ce9d5ff16c1d8230fea7c13cb129c15d458"></a>

## local_control_plane.bgp_config.peers.external.external_connector — local_control_plane.bgp_config.peers.external.external_connector / fe0c2c9f7e4a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- local_control_plane.bgp_config.peers.external.external_connector

<a id="canonical-c194d86d821b64abd9f5167b7a0ef482230eee5db130e9a68425a703427cc550"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for external connector.

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

<a id="canonical-99fa26601a509b591404afe0c8f48531eca2af358d18763bbf864aea482da5c4"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.external_connector / fe0c2c9f7e4a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-55c5c3f325f4544d0d691e7a9c1fdc8c9445c7f6d2506ac3f2c6a6a3cf8e418b"></a>

## Next pages — local_control_plane.bgp_config.peers.external.external_connector / fe0c2c9f7e4a / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-e90ff08ae2d6ac3dff1c616ef4eca5af5d75b10fc4923fcf4448ada606e2972f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d635502b9598da52b56df2e16ea4bf351cb59febc46cc9bf81b0277785325253"></a>

## local_control_plane.bgp_config.peers.external.family_inet — local_control_plane.bgp_config.peers.external.family_inet / 042e5c9d9666 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- local_control_plane.bgp_config.peers.external.family_inet

<a id="canonical-b93fe191eeda589613d8eec3d0f940ec11088bcf056b1fbadae6dff65fcc067e"></a>

Type: `"single"`. Computed.

Configuration parameter for family inet.

Upstream description:

Parameters for inet family.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-enable_choice": "[\"disable\",\"enable\"]"
}
```

<a id="canonical-06757dc03f2e8b7132009db452d52a328767e7057bfc70dfe30fd9931666dd17"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.family_inet / 042e5c9d9666 / 3

- [disable_spec](data-sources--voltstack_site--reference--group-009.md#canonical-513e321da4c7e840e17ab9f66f95058fb2df9601bb55cf351dd47d7756f596a7): complete subsection reference.

- [enable](data-sources--voltstack_site--reference--group-009.md#canonical-b2b8fa0f785366e93b5dc02da3161b78af5ffaaa29747c3c7a354060107f030e): complete subsection reference.

<a id="canonical-a9984dd847d888da67e1aabb026503670e2cb13782d169474f7eaefe3b600ea2"></a>

## Next pages — local_control_plane.bgp_config.peers.external.family_inet / 042e5c9d9666 / 4

- [local_control_plane.bgp_config.peers.external.family_inet.disable_spec](data-sources--voltstack_site--reference--group-009.md#canonical-513e321da4c7e840e17ab9f66f95058fb2df9601bb55cf351dd47d7756f596a7)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](data-sources--voltstack_site--reference--group-009.md#canonical-b2b8fa0f785366e93b5dc02da3161b78af5ffaaa29747c3c7a354060107f030e)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-513e321da4c7e840e17ab9f66f95058fb2df9601bb55cf351dd47d7756f596a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7d58fc3e54f3eac27e6aa9645e150f10ddd85012d33e4eabbfc638b706aaa7d"></a>

## local_control_plane.bgp_config.peers.external.family_inet.disable_spec — local_control_plane.bgp_config.peers.external.family_inet.disable_spec / c17b4c677e5e / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-e90ff08ae2d6ac3dff1c616ef4eca5af5d75b10fc4923fcf4448ada606e2972f)
- local_control_plane.bgp_config.peers.external.family_inet.disable_spec

<a id="canonical-70ce0e4118facde04000e4263b476e9da22cfc22f71b6546982f868e70088c2a"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-655cc3a3ed9d15747c33c5903f5719f094a015fa93f13bb49b2e4af449de435c"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.family_inet.disable_spec / c17b4c677e5e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c69af14a83927b9f61b5b08650f5aacd442ea13e2a4e47ac34eb11bc6802961f"></a>

## Next pages — local_control_plane.bgp_config.peers.external.family_inet.disable_spec / c17b4c677e5e / 4

- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-e90ff08ae2d6ac3dff1c616ef4eca5af5d75b10fc4923fcf4448ada606e2972f)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-b2b8fa0f785366e93b5dc02da3161b78af5ffaaa29747c3c7a354060107f030e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e94a8e15128a9c5856ea4e497b145fd5d0bf898af13b8eed5d94b8520f4ad33d"></a>

## local_control_plane.bgp_config.peers.external.family_inet.enable — local_control_plane.bgp_config.peers.external.family_inet.enable / 32af8614d10f / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-e90ff08ae2d6ac3dff1c616ef4eca5af5d75b10fc4923fcf4448ada606e2972f)
- local_control_plane.bgp_config.peers.external.family_inet.enable

<a id="canonical-961575c117e08cb67fe6ed6fe32b0d6a31d71c62c7b75491d801cc124cd2b48c"></a>

Type: `"single"`. Computed.

Unicast IPv4. IPv4 Unicast.

Upstream description:

IPv4 Unicast.

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

<a id="canonical-ee2a2955507e4d13a84090d771a048537fa636d17b1dd846447f140ebffb11d6"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.family_inet.enable / 32af8614d10f / 3

- [aggregation](data-sources--voltstack_site--reference--group-009.md#canonical-2d132bcf0f385bfc61aff3b4ee1ee7246ff3ffee9a788738a28d93c1cf4b51da): complete subsection reference.

<a id="canonical-4d647364f7afeb0fd890232bdc6df9257788d4256e9d3231a9a5988bb81027f2"></a>

## Next pages — local_control_plane.bgp_config.peers.external.family_inet.enable / 32af8614d10f / 4

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](data-sources--voltstack_site--reference--group-009.md#canonical-2d132bcf0f385bfc61aff3b4ee1ee7246ff3ffee9a788738a28d93c1cf4b51da)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-e90ff08ae2d6ac3dff1c616ef4eca5af5d75b10fc4923fcf4448ada606e2972f)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-2d132bcf0f385bfc61aff3b4ee1ee7246ff3ffee9a788738a28d93c1cf4b51da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d63823e3a599c526d33464f5590e0dc6b197f605c85cbca458b25307871667fe"></a>

## local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation / e78d48f6ec1e / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-e90ff08ae2d6ac3dff1c616ef4eca5af5d75b10fc4923fcf4448ada606e2972f)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](data-sources--voltstack_site--reference--group-009.md#canonical-b2b8fa0f785366e93b5dc02da3161b78af5ffaaa29747c3c7a354060107f030e)
- local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation

<a id="canonical-cbaa4f1644f879bd91006f8471facc87b76b478682103a3a4b98221f85156a4f"></a>

Type: `"list"`. Computed.

BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take
effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing
table and applies to outbound advertisements.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ffc9d7598f45a101be13722848287432b2d1f9eedd3212bd53a1823e670adb71"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation / e78d48f6ec1e / 3

<a id="canonical-5595bb1cd2dc705a97dbc172546af78579e8d0be95cf1e0f9e2290cac837a54c"></a>

<a id="canonical-a268de767629b870bebf2af094f58dcb5aef51ea578fb8657dd4297b078b9177"></a>

## ip_prefix property — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation / e78d48f6ec1e / 4

Type: `"string"`. Computed.

IP Prefix. Specify IPv4 subnet for aggregation.

Upstream description:

Specify IPv4 subnet for aggregation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

- [options](data-sources--voltstack_site--reference--group-009.md#canonical-3a9e626653d7170189dab86338ccb3322808dd0f109ecd23fcea52e2818f0a3f): complete subsection reference.

<a id="canonical-0f18d569def2f6e4b8d0b9fac3b655db362dfea08b675e8e3e62aab8d5d61255"></a>

## Next pages — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation / e78d48f6ec1e / 5

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options](data-sources--voltstack_site--reference--group-009.md#canonical-3a9e626653d7170189dab86338ccb3322808dd0f109ecd23fcea52e2818f0a3f)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](data-sources--voltstack_site--reference--group-009.md#canonical-b2b8fa0f785366e93b5dc02da3161b78af5ffaaa29747c3c7a354060107f030e)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-3a9e626653d7170189dab86338ccb3322808dd0f109ecd23fcea52e2818f0a3f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd38fd678d6f64cf87e8d13ddb7c015ca25acf5fe4a052f60cef19cfbd434812"></a>

## local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.opt / 266882379648 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-e90ff08ae2d6ac3dff1c616ef4eca5af5d75b10fc4923fcf4448ada606e2972f)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](data-sources--voltstack_site--reference--group-009.md#canonical-b2b8fa0f785366e93b5dc02da3161b78af5ffaaa29747c3c7a354060107f030e)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](data-sources--voltstack_site--reference--group-009.md#canonical-2d132bcf0f385bfc61aff3b4ee1ee7246ff3ffee9a788738a28d93c1cf4b51da)
- local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options

<a id="canonical-590c715639c934ee92af4f721249bb5f33d906c565dcd3f7b8c2b83c818c6dfe"></a>

Type: `"list"`. Computed.

Aggregation OPTIONS. Configuration parameter for options

Upstream description:

Configuration parameter for options

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1d9a4846bf26ba927a6f9e0940d8462db726e7fd16b366d1a0684e6580bce1fa"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.opt / 266882379648 / 3

- [summary_only](data-sources--voltstack_site--reference--group-009.md#canonical-40dce4e106dcb65a5b67f7d4b56557469d98612b8dfef34a81669b58d99cffbe): complete subsection reference.

<a id="canonical-ff8e294b6fdc9add3d5fa981be86abdf60737722e03f8c4f2c50cbbb64f8911a"></a>

## Next pages — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.opt / 266882379648 / 4

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options.summary_only](data-sources--voltstack_site--reference--group-009.md#canonical-40dce4e106dcb65a5b67f7d4b56557469d98612b8dfef34a81669b58d99cffbe)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](data-sources--voltstack_site--reference--group-009.md#canonical-2d132bcf0f385bfc61aff3b4ee1ee7246ff3ffee9a788738a28d93c1cf4b51da)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-40dce4e106dcb65a5b67f7d4b56557469d98612b8dfef34a81669b58d99cffbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b890667c87a8c9ba727a1d94f9cba03885a4749fc45aea22129f6278421b76cb"></a>

## local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options.summary_only — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.opt / e3329f8d6844 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-e90ff08ae2d6ac3dff1c616ef4eca5af5d75b10fc4923fcf4448ada606e2972f)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](data-sources--voltstack_site--reference--group-009.md#canonical-b2b8fa0f785366e93b5dc02da3161b78af5ffaaa29747c3c7a354060107f030e)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](data-sources--voltstack_site--reference--group-009.md#canonical-2d132bcf0f385bfc61aff3b4ee1ee7246ff3ffee9a788738a28d93c1cf4b51da)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options](data-sources--voltstack_site--reference--group-009.md#canonical-3a9e626653d7170189dab86338ccb3322808dd0f109ecd23fcea52e2818f0a3f)
- local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options.summary_only

<a id="canonical-0f059edae64242ab38942142d1f8a20729ab31ac12c963c8f2d98cb37e5f5c00"></a>

Type: `"single"`. Computed.

Configuration parameter for summary only.

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

<a id="canonical-b810aec69e625a8a27a623053e70ed86b9d4d2e44c51004375b41e0b8794e2f1"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.opt / e3329f8d6844 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cfb4a06c3f8a37da2afe6c31936ac405bbf17a728d6e548b19c61a3dacc1f83d"></a>

## Next pages — local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.opt / e3329f8d6844 / 4

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options](data-sources--voltstack_site--reference--group-009.md#canonical-3a9e626653d7170189dab86338ccb3322808dd0f109ecd23fcea52e2818f0a3f)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-618446a18ab61ec08ea6049f28686e2d4e8a0f035ec0f9cb7e6aec45a337cf3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e9209bd0d0a1de66aaad2a836b4e312d63eb70eec8cbe18cff75df571b494f2"></a>

## local_control_plane.bgp_config.peers.external.from_site — local_control_plane.bgp_config.peers.external.from_site / 20048c56a3e5 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- local_control_plane.bgp_config.peers.external.from_site

<a id="canonical-3e5edb7d396647772d19da36d732ded114d6544c88c85b1bee9b521438e1eabd"></a>

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

<a id="canonical-68da2148a051ddb92273c7068f7701f2e9d293e9f3f7d673ee9331fe8ea0880f"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.from_site / 20048c56a3e5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-57a495b359d0f8c1feaa9a7182efc5f7ff803eb550df3ab07fa0a0457641ea00"></a>

## Next pages — local_control_plane.bgp_config.peers.external.from_site / 20048c56a3e5 / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-97edea0468c34f5dfb1e511f99a0f0f98887bb255c153ec220a0dd8fbbc11edb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f84f41751bc34a6ab71a172ac3ab553b3771718b7722a1980ddeacc66039f90"></a>

## local_control_plane.bgp_config.peers.external.from_site_v6 — local_control_plane.bgp_config.peers.external.from_site_v6 / 47f177c33972 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- local_control_plane.bgp_config.peers.external.from_site_v6

<a id="canonical-155601d08466f4a4a614e404f6c7dfd21b2ba86f6daca3be77329afb3b96c413"></a>

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

<a id="canonical-8b1879065f01fbb2b6c6768508dd77a84e717296f21ea534067eff58c5ab99cf"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.from_site_v6 / 47f177c33972 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-607dc754fb1f71c83c668c63c2312e7c0fa303dad9e7247a27e75bffbcc031e5"></a>

## Next pages — local_control_plane.bgp_config.peers.external.from_site_v6 / 47f177c33972 / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-3e4587b1b331c8080af3b4d8bbccd2562f23a1c017c7aff43f2d5ac6b70aaad4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f831eed4734acd047cdad3ac13e8e18308595930ac5a28366a97839ab347a08"></a>

## local_control_plane.bgp_config.peers.external.interface — local_control_plane.bgp_config.peers.external.interface / 2f0fc3f285e5 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- local_control_plane.bgp_config.peers.external.interface

<a id="canonical-f4bff1f7dd577c5fa1d2d3cd6de324e1f1ba39f777c7c1668576cc58b862677f"></a>

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

<a id="canonical-0413428b1f00108c2c2a0a4b9ce113e2f7009b4923999ff6452e58886448af5c"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.interface / 2f0fc3f285e5 / 3

<a id="canonical-33f0fa0c53171bd3524b52a1e9f79201a9c01824b2437ffd483d8a9aba5f3522"></a>

<a id="canonical-a42d1d631f9d050707a2f8502eaf8f4be7645960e07dd451408831205c806091"></a>

## name property — local_control_plane.bgp_config.peers.external.interface / 2f0fc3f285e5 / 4

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

<a id="canonical-17f86de5ca08e08fa1ad5b7cbd9018e9d4f992026ce4ce94524051148986bf08"></a>

<a id="canonical-a5129499059d3c3acc199bc7f94abfedbdb1fb465b6823056bd3d2966b7cf8ab"></a>

## namespace property — local_control_plane.bgp_config.peers.external.interface / 2f0fc3f285e5 / 5

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

<a id="canonical-3165086b556fdbf0fe36d85adbe16e0951b3299b8f92349fd425e026179f3cbf"></a>

<a id="canonical-b9847bef9b4106b804914e6248825e5a946cc37cfeb911193bef4dae955e93b2"></a>

## tenant property — local_control_plane.bgp_config.peers.external.interface / 2f0fc3f285e5 / 6

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

<a id="canonical-36f990e7452147c2d4f0d06357096bbb65adba09029e2991d9960dc9ea7987a2"></a>

## Next pages — local_control_plane.bgp_config.peers.external.interface / 2f0fc3f285e5 / 7

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-116ee59de0c582f666df0dc817e672bd50576a5e001480acbcb410c18e6a45a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b6e2fa1a30e6eb9eeacbc133e619a48405c878b2da2f4d25af2b15917460497"></a>

## local_control_plane.bgp_config.peers.external.interface_list — local_control_plane.bgp_config.peers.external.interface_list / df9af214a30d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- local_control_plane.bgp_config.peers.external.interface_list

<a id="canonical-4b20028fa8559df1083ae0248ea6b0c8ff858a54001980117f3d5b92d2faeab4"></a>

Type: `"single"`. Computed.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

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

<a id="canonical-48430e57c4eda35a533abeea7010085c59b3e6120166691ccfc5245a9ec7ca42"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.interface_list / df9af214a30d / 3

- [interfaces](data-sources--voltstack_site--reference--group-009.md#canonical-c2e55af403384ea747df1a3aee2245af72ad112d323f09589a2f75e71190376a): complete subsection reference.

<a id="canonical-6825a2fdd9807a60456f000edfd7e159b162988457af153cd9639ee699a8f530"></a>

## Next pages — local_control_plane.bgp_config.peers.external.interface_list / df9af214a30d / 4

- [local_control_plane.bgp_config.peers.external.interface_list.interfaces](data-sources--voltstack_site--reference--group-009.md#canonical-c2e55af403384ea747df1a3aee2245af72ad112d323f09589a2f75e71190376a)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-c2e55af403384ea747df1a3aee2245af72ad112d323f09589a2f75e71190376a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf29181cc1dc3b43ebf27cb75a35afb837205974e170aac1714ee170524c6b33"></a>

## local_control_plane.bgp_config.peers.external.interface_list.interfaces — local_control_plane.bgp_config.peers.external.interface_list.interfaces / f200403091de / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [local_control_plane.bgp_config.peers.external.interface_list](data-sources--voltstack_site--reference--group-009.md#canonical-116ee59de0c582f666df0dc817e672bd50576a5e001480acbcb410c18e6a45a3)
- local_control_plane.bgp_config.peers.external.interface_list.interfaces

<a id="canonical-db06122ca45fc9ab9ea42b6f0e487d22cf83fe4c3d221a420b572aa671f51e86"></a>

Type: `"list"`. Computed.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3b60750d9c28c715a7681680de5ada1bdd86419f2b788d504db3e55639ae622d"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.interface_list.interfaces / f200403091de / 3

<a id="canonical-6c6b6d11fbac35779eb81645de94e3a6d6aec12804da60aeed8391e8e5307e25"></a>

<a id="canonical-3f27c64f887a317daa2e04572f04a487e227035a3ab092828c435aeb169ff4e2"></a>

## name property — local_control_plane.bgp_config.peers.external.interface_list.interfaces / f200403091de / 4

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

<a id="canonical-a057ff930f0218680d4107e02cc02a9af052142ae93033f8150e23bd8ce7c766"></a>

<a id="canonical-fdb6e3a582533b1df8dbccca1ccf17906104a7d851d1d9c688099cb250c71508"></a>

## namespace property — local_control_plane.bgp_config.peers.external.interface_list.interfaces / f200403091de / 5

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

<a id="canonical-b2c5d611f6bf938eabd53db92ff524e988d5b9cdb92e9b4e351258f16cb635c5"></a>

<a id="canonical-b70b90a672d4eecf11e75f394f7e7f28efd038fe68aa6588bdae1dce0728f240"></a>

## tenant property — local_control_plane.bgp_config.peers.external.interface_list.interfaces / f200403091de / 6

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

<a id="canonical-f3a9ad8cc91c2ecc11d7dfb2b4c10096e2b59c949ac61172dd2df6a989c683d1"></a>

## Next pages — local_control_plane.bgp_config.peers.external.interface_list.interfaces / f200403091de / 7

- [local_control_plane.bgp_config.peers.external.interface_list](data-sources--voltstack_site--reference--group-009.md#canonical-116ee59de0c582f666df0dc817e672bd50576a5e001480acbcb410c18e6a45a3)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-b82a22a43e45020b086a1b1c57fa9d6d9b202de7a1df57bfc171be28380834c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3a175825c628fac0ca591ee867e403f9ae328c0765ce58906e801f53fe97fd0"></a>

## local_control_plane.bgp_config.peers.external.no_authentication — local_control_plane.bgp_config.peers.external.no_authentication / a32a766b54b3 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- local_control_plane.bgp_config.peers.external.no_authentication

<a id="canonical-ba73e96687f04229af63066509f277d82ede7265e7e8cc3c6a658c4d1bfcd5f6"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no authentication.

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

<a id="canonical-60ea1933574b11ab6e943c730f69aa7ac83b57607adc7cfcea7b3a92f032c2dd"></a>

## Direct properties — local_control_plane.bgp_config.peers.external.no_authentication / a32a766b54b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b2b8d1aa0fd5af2a4c0b6c376495da9e700062557cbc0c0f3ccc26354d8c9d96"></a>

## Next pages — local_control_plane.bgp_config.peers.external.no_authentication / a32a766b54b3 / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-39422463d45c505c5cad6c3137d411e0b7c1cd74c66a95ca07eb5e82f07608ce)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-6482a1c2bff33aacf0b2c7706ee98b6fe19044be5e468dc936f39ff25fc38667"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ed80c486f965d5fd82f5ce8ceef4504d121cc0031483d80618cb664bbfb4252"></a>

## local_control_plane.bgp_config.peers.metadata — local_control_plane.bgp_config.peers.metadata / fe7e279f94d7 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- local_control_plane.bgp_config.peers.metadata

<a id="canonical-45fbd6fed1013ffa5eeb18476d5cc13288e20c46ba6d5c8ad1f1c4c4110a3096"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-204233dd096edc2e9ae3bd7e564efd12fbf32f50725bdf5627cc32a98027bba5"></a>

## Direct properties — local_control_plane.bgp_config.peers.metadata / fe7e279f94d7 / 3

<a id="canonical-f57b8ab78895878a44ece6b59f8577eea5e01c61cc0a6a9389135949c8d73c25"></a>

<a id="canonical-0cd947b0d97d646a903b0836dbc1c7f5d5108e2b8aa915ba85319a05aa2b2f18"></a>

## description_spec property — local_control_plane.bgp_config.peers.metadata / fe7e279f94d7 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-eb3c955ee865b1e4e9540b8152a03313d7b3e1eebfb7358d1dcc5450fc3fdd05"></a>

<a id="canonical-d8a54bba0cea9342079d00c35d8e38bd91fa2ad3909dec9200e0200cc3072043"></a>

## name property — local_control_plane.bgp_config.peers.metadata / fe7e279f94d7 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-4f012dc42ce6ad1506cd668c07e991cc764b40e8df877560114d079482f50ffe"></a>

## Next pages — local_control_plane.bgp_config.peers.metadata / fe7e279f94d7 / 6

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-cc37a0c2f4f5734c7f325e0105e03da45f8f74d1a3248dd9f95fbbacc2560e08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2da47450797b4a0b740362367e4721714a736b60ff6114c4cb9703e13d5165b6"></a>

## local_control_plane.bgp_config.peers.passive_mode_disabled — local_control_plane.bgp_config.peers.passive_mode_disabled / 1ccf5a63dcc8 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- local_control_plane.bgp_config.peers.passive_mode_disabled

<a id="canonical-1ef68c392db2295fbb3374e90cef5de6f520e0773be12bf368e05ec103293d96"></a>

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

<a id="canonical-3ff5864a4de495b28a4736a75c1e3615978a149520c8d81dc9f434c83a26c765"></a>

## Direct properties — local_control_plane.bgp_config.peers.passive_mode_disabled / 1ccf5a63dcc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e8c66160fef3e7eced410a1f33c2e9285593d7a3697b8da52abf9e603271855f"></a>

## Next pages — local_control_plane.bgp_config.peers.passive_mode_disabled / 1ccf5a63dcc8 / 4

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-47093971b2b341da17870928b0543f792823327bba8d2986b45294404479dd37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e561a2d3c3afb6b3bf9768ace6d72fd6af63336b81fc1c7b3c382f38d961fea7"></a>

## local_control_plane.bgp_config.peers.passive_mode_enabled — local_control_plane.bgp_config.peers.passive_mode_enabled / b072abb98e0a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- local_control_plane.bgp_config.peers.passive_mode_enabled

<a id="canonical-baaecea9b6bff86f0ff74046e2ef2061bbf34388a2eb47ea5734060739cc3287"></a>

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

<a id="canonical-d4c1695e4a0381d03a74afd3b8d1551bb221a6709ffd1e5af05621d0c596bd6a"></a>

## Direct properties — local_control_plane.bgp_config.peers.passive_mode_enabled / b072abb98e0a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ec650486b8043883d1b5a3f76811746e353ff202ec45d020c05dedfbbb3e3810"></a>

## Next pages — local_control_plane.bgp_config.peers.passive_mode_enabled / b072abb98e0a / 4

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-812d3ca793586e03a0c6af261d344d002bade4f06ec75d2e9de3f516ef99deee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-649fb55d79fb42f35aa9e232be220621d4c2645265ce2071bf507838cb9211f6"></a>

## local_control_plane.bgp_config.peers.routing_policies — local_control_plane.bgp_config.peers.routing_policies / 6e8bb6cd808e / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- local_control_plane.bgp_config.peers.routing_policies

<a id="canonical-2a2b32a069637c40b52b28862ea7194f9911e6c4733df4ddbc853abe1ad02860"></a>

Type: `"single"`. Computed.

List of rules which can be applied on all or particular nodes.

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

<a id="canonical-4d7ce72d9f285fb66ad207e8fd3f864aba47aacd9d7287e27c5252300769d2a8"></a>

## Direct properties — local_control_plane.bgp_config.peers.routing_policies / 6e8bb6cd808e / 3

- [route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-ceeec3a1bf6e0ca49308b3f59f95184ed4d0b4bc5592f3594ba026a597e027bb): complete subsection reference.

<a id="canonical-22b4361f3caafe50a56095a34e6faa893d686969501340c9cfc619444473e6e6"></a>

## Next pages — local_control_plane.bgp_config.peers.routing_policies / 6e8bb6cd808e / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-ceeec3a1bf6e0ca49308b3f59f95184ed4d0b4bc5592f3594ba026a597e027bb)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-ceeec3a1bf6e0ca49308b3f59f95184ed4d0b4bc5592f3594ba026a597e027bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b02f6f9cb9165e59639f83ce2d869c048c5284bc6e0b9436382d9289840d7b53"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy — local_control_plane.bgp_config.peers.routing_policies.route_policy / 5de74a44a5cf / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-812d3ca793586e03a0c6af261d344d002bade4f06ec75d2e9de3f516ef99deee)
- local_control_plane.bgp_config.peers.routing_policies.route_policy

<a id="canonical-1429b5e4911b7baf0ba56fc91e4ca36f12fcea28b4f02c52ba28f82aadc19d50"></a>

Type: `"list"`. Computed.

Policy configuration for this feature.

Upstream description:

Route policy to be applied.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-292b1dc18d5e205b10a5807ae3e88fb9fce530dbb86e3c22d59d835818a462ff"></a>

## Direct properties — local_control_plane.bgp_config.peers.routing_policies.route_policy / 5de74a44a5cf / 3

- [all_nodes](data-sources--voltstack_site--reference--group-009.md#canonical-02d2f8db13199c9ff52a5466904a1cbab700bec4a25d009d99ac16fc1045eb9a): complete subsection reference.

- [inbound](data-sources--voltstack_site--reference--group-009.md#canonical-14f9f21165d0ff4fa57e005b306f77ed3b7c20f2711e676e3f76125b9c42fa5a): complete subsection reference.

- [node_name](data-sources--voltstack_site--reference--group-009.md#canonical-da0585665d41e6220e032ad2703d9b035733b2472df9b6905bb3ab1671afc77f): complete subsection reference.

- [object_refs](data-sources--voltstack_site--reference--group-009.md#canonical-c17df504336c1565ecd6a6c2296a27cdc24cbb12e0abedb932e813f8e5b9f2c1): complete subsection reference.

- [outbound](data-sources--voltstack_site--reference--group-009.md#canonical-1bb3cdd434b68d504868a84930478ac2ab17553ee2d0f41f127045c49d4b8106): complete subsection reference.

<a id="canonical-24fb45d3952500f9080edc506d933d1e5439c3455fb06ef7806bd82f354d2250"></a>

## Next pages — local_control_plane.bgp_config.peers.routing_policies.route_policy / 5de74a44a5cf / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes](data-sources--voltstack_site--reference--group-009.md#canonical-02d2f8db13199c9ff52a5466904a1cbab700bec4a25d009d99ac16fc1045eb9a)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound](data-sources--voltstack_site--reference--group-009.md#canonical-14f9f21165d0ff4fa57e005b306f77ed3b7c20f2711e676e3f76125b9c42fa5a)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name](data-sources--voltstack_site--reference--group-009.md#canonical-da0585665d41e6220e032ad2703d9b035733b2472df9b6905bb3ab1671afc77f)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs](data-sources--voltstack_site--reference--group-009.md#canonical-c17df504336c1565ecd6a6c2296a27cdc24cbb12e0abedb932e813f8e5b9f2c1)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound](data-sources--voltstack_site--reference--group-009.md#canonical-1bb3cdd434b68d504868a84930478ac2ab17553ee2d0f41f127045c49d4b8106)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-812d3ca793586e03a0c6af261d344d002bade4f06ec75d2e9de3f516ef99deee)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-02d2f8db13199c9ff52a5466904a1cbab700bec4a25d009d99ac16fc1045eb9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0e3735ced058d790de73caa3fd332b47b0c83f0b52b5dc615e1f2789a3ece6b"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes — local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes / 8e18c3591723 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-812d3ca793586e03a0c6af261d344d002bade4f06ec75d2e9de3f516ef99deee)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-ceeec3a1bf6e0ca49308b3f59f95184ed4d0b4bc5592f3594ba026a597e027bb)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes

<a id="canonical-d1df138db1a880d0dffd3557871a6aab759a5405f3bf07b4d1c9266431c564ac"></a>

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

<a id="canonical-86d5974f34e275e53cd42d9b395f23198dda0897258b618be38df854138b7162"></a>

## Direct properties — local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes / 8e18c3591723 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b835936179f9b8dbe22d8571e9a72439a8db0f53929fb3b91301fd03d7f9fca0"></a>

## Next pages — local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes / 8e18c3591723 / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-ceeec3a1bf6e0ca49308b3f59f95184ed4d0b4bc5592f3594ba026a597e027bb)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-14f9f21165d0ff4fa57e005b306f77ed3b7c20f2711e676e3f76125b9c42fa5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6338bbd4c578c155b56d852b6dede2c4cb4d469bf7e900aed31260441937438"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound — local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound / 3be5c51531db / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-812d3ca793586e03a0c6af261d344d002bade4f06ec75d2e9de3f516ef99deee)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-ceeec3a1bf6e0ca49308b3f59f95184ed4d0b4bc5592f3594ba026a597e027bb)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound

<a id="canonical-bd9b17e372a7256b5737422070d742417b9f31c52dab26e236091d811bab85ea"></a>

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

<a id="canonical-750ff83b613d9cd2407f3978d742f2708c3bd984286a1f35fb213213e08b5b77"></a>

## Direct properties — local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound / 3be5c51531db / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-292705ab68cafd5d12ac4d0c1338f6b5261ea1d2cb949ed924f3553837050ca6"></a>

## Next pages — local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound / 3be5c51531db / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-ceeec3a1bf6e0ca49308b3f59f95184ed4d0b4bc5592f3594ba026a597e027bb)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-da0585665d41e6220e032ad2703d9b035733b2472df9b6905bb3ab1671afc77f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77bbad5a6e5cd32d4915409f98f3131aceec415ec22ea38f5790bec9bffea588"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name — local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name / 5a86c73f836d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-812d3ca793586e03a0c6af261d344d002bade4f06ec75d2e9de3f516ef99deee)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-ceeec3a1bf6e0ca49308b3f59f95184ed4d0b4bc5592f3594ba026a597e027bb)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name

<a id="canonical-402bc6b10c57302c7e319fd20f7801d7e1fd9eade017ff4ade3695e32730c48c"></a>

Type: `"single"`. Computed.

List of nodes on which BGP routing policy has to be applied.

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

<a id="canonical-5fc2d83d291b7f7c4624241c745781db7dea64e3a30619ecf971e173f3923a84"></a>

## Direct properties — local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name / 5a86c73f836d / 3

<a id="canonical-d0c03a6a5f4eeb8b695982b7ae4b872c244a623450ef939789131cb4399979ee"></a>

<a id="canonical-45b1d43bf545fc63e9037b662e335d96dd9f421f091673d7dcb62c407aa62548"></a>

## node property — local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name / 5a86c73f836d / 4

Type: `["list", "string"]`. Computed.

Select BGP Session on which policy will be applied.

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

<a id="canonical-411c9c9fa6bfe2d3a8a7c89e17b19d81bf0dea694138707f0497e92b143810c6"></a>

## Next pages — local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name / 5a86c73f836d / 5

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-ceeec3a1bf6e0ca49308b3f59f95184ed4d0b4bc5592f3594ba026a597e027bb)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-c17df504336c1565ecd6a6c2296a27cdc24cbb12e0abedb932e813f8e5b9f2c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f451498c0ecb53cd5f7ef1fd29394fee0aeb08be117803b213b28357e6b5901d"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 9ccf995dffa6 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-812d3ca793586e03a0c6af261d344d002bade4f06ec75d2e9de3f516ef99deee)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-ceeec3a1bf6e0ca49308b3f59f95184ed4d0b4bc5592f3594ba026a597e027bb)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs

<a id="canonical-e178a7df9bcd192aa4b690d40c45fbf6172486fdfe85b9bcd873291c4c51cfe1"></a>

Type: `"list"`. Computed.

BGP routing policy. Select route policy to apply.

Upstream description:

Select route policy to apply.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2bbbd62321ed3b1b81545171732139418460478c0526f284b2f04388ca150b81"></a>

## Direct properties — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 9ccf995dffa6 / 3

<a id="canonical-76222a999ded2b4153d4de8fc012b867ae17becf41cde9553a98c78bd36e7543"></a>

<a id="canonical-f3ae9a1057e56c7d2bfd378d6a378c3885c01af0089f885a63c73d50e380de14"></a>

## kind property — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 9ccf995dffa6 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-b31ad91428303d1cf7ca119908c038e371fa237c3a23ba2f1dfbaf982cc02912"></a>

<a id="canonical-518f0dcb261d346a9b8af72b0bdf0bccf7d703eeefddb82c5fa9243eb9bb4c7c"></a>

## name property — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 9ccf995dffa6 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-24cf9d673cbc89780ff432a115dd217ec65c9171ae82307d77a26b82fb68c68a"></a>

<a id="canonical-2d3dc7cffa581cd1cfb23c35c79d705800bdfdbc99e707cb8691daae86f371ee"></a>

## namespace property — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 9ccf995dffa6 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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

<a id="canonical-000bd00a8b0cfd5825cc50e55c355c2b7716e4725d85335d67cb6bd01603bd59"></a>

<a id="canonical-73e95209bd36d21252385fc29d366d75e5fd0b432c3ac175c453dcd7d2b31f62"></a>

## tenant property — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 9ccf995dffa6 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-5f4a5cbc434b6074cfb5ecd2b4d59ae8c4e381409101cc8e0b43ccbb8ed214d1"></a>

<a id="canonical-04565d725e221e4df16d6f321c873300862b248a711f8f3d575e14ad2b58df29"></a>

## uid property — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 9ccf995dffa6 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-359bdaed0d252419276b3640f8714e9eb0e2482eb74c8972a334488ac1d49588"></a>

## Next pages — local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs / 9ccf995dffa6 / 9

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-ceeec3a1bf6e0ca49308b3f59f95184ed4d0b4bc5592f3594ba026a597e027bb)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-1bb3cdd434b68d504868a84930478ac2ab17553ee2d0f41f127045c49d4b8106"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c05b625860e02e03e8619a8bc1c543ba6eb74b4eebe3eb3fd4d2495a7f1e551"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound — local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound / a6cd40d526d5 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-008.md#canonical-436f4b922ec20d380c0a2b0abf6763369d4a3003d477c3f133c67b369f6cd3d0)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-008.md#canonical-1b6e23047ea87b9ca318317180ebffafd115ccd4b9ee4d4c92db5fb57e591bb0)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-812d3ca793586e03a0c6af261d344d002bade4f06ec75d2e9de3f516ef99deee)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-ceeec3a1bf6e0ca49308b3f59f95184ed4d0b4bc5592f3594ba026a597e027bb)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound

<a id="canonical-20e35c4a64a848f34bfa93620f85e3bbe0c4d16cdf6e02aaa4d8f3183413b83b"></a>

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

<a id="canonical-15046b0f136a2885db472e9fba76d105325608289cddae476e9b93dc760a72ff"></a>

## Direct properties — local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound / a6cd40d526d5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5ff082aae80ace50c3c0e56105306ce0d0611c8de190de8bc49b44cff413bee3"></a>

## Next pages — local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound / a6cd40d526d5 / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-ceeec3a1bf6e0ca49308b3f59f95184ed4d0b4bc5592f3594ba026a597e027bb)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-1c8590fa4ff0eec10e050feee7ae8feb70316ef0a2d10b962a4f79ba14284c44"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc8c51ad9836b92a34988f163224a8e05d7d70b18ec69ddb90eea7ae45d0e34f"></a>

## local_control_plane.inside_vn — local_control_plane.inside_vn / 0c63b0252850 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- local_control_plane.inside_vn

<a id="canonical-3558972d97dd0ac843f106026ff9b43eaecfa55208b752a893212a11a026107b"></a>

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

<a id="canonical-369c094db1b186a1d939a0508066b7ff4eccbc829a375ef3ab50c7b828c64abf"></a>

## Direct properties — local_control_plane.inside_vn / 0c63b0252850 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-97520dc6abce34a7e1376b58c6f66bc5d7a76881f7d05ff5882a0bcf13e04b6e"></a>

## Next pages — local_control_plane.inside_vn / 0c63b0252850 / 4

- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-ed54e199badd7b7344fa005db28b8b06de6590dd6f08715ffc0ea9ad9221176a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0aa64241cf9675120cba2591736558cdbb37b5d680073e8671ed58ff45a6e9d8"></a>

## local_control_plane.outside_vn — local_control_plane.outside_vn / 9de2c62b2827 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- local_control_plane.outside_vn

<a id="canonical-a87fc72eda5e30c2774806b096a564f4dda1c7133734164d018e89299e1dd7e9"></a>

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

<a id="canonical-01ba8d0c0cc1b97515bd0f059f412d5842f2faafc328b9e03367911bf0a63708"></a>

## Direct properties — local_control_plane.outside_vn / 9de2c62b2827 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7be274317c77a0ae45524abd73503e38107aeb006ba6098f86f8c995679b7698"></a>

## Next pages — local_control_plane.outside_vn / 9de2c62b2827 / 4

- [local_control_plane](data-sources--voltstack_site--reference--group-008.md#canonical-074fedc0e78e2fe7b8b8cba87435c45f51880c6a1d76b5be7a85d0904f514280)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-50c701645a62b74a0ef3a5b0527f69b38d4685d66f5c40ed2a742cc015b587b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31ba99e297271b5e10bd2bd15f477ef690a1db0af3532a777f7aa90b891f8f16"></a>

## log_receiver — log_receiver / 14b8a72ac1ac / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- log_receiver

<a id="canonical-8fdd88d0d0e47ba57e5d64d2fe899b01b8e8356c686ee017f71592a113e43457"></a>

Type: `"single"`. Computed.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [log_receiver](data-sources--voltstack_site--reference--group-009.md#canonical-8fdd88d0d0e47ba57e5d64d2fe899b01b8e8356c686ee017f71592a113e43457)
- [logs_streaming_disabled](data-sources--voltstack_site--reference--group-009.md#canonical-30b3f7d4db67b3f25649ff904e2c3b2d2053f711575a940ebb123752a2137172)

Select alternatives according to the provider validators above.

<a id="canonical-af5f55560d8e8c40449d5dbb04bb606c7b56c9f646aa6fafa472a6eab166c619"></a>

## Direct properties — log_receiver / 14b8a72ac1ac / 3

<a id="canonical-67d10d852e9a738d72fa03590b94e18c5960009327366a099084ee57f0372e0d"></a>

<a id="canonical-0244f3d1cfa427ef9e01e3ff5c9ae19f3e64a27bd53da364d12878d71cf679a4"></a>

## name property — log_receiver / 14b8a72ac1ac / 4

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

<a id="canonical-4820379467603da28611eaeca3e05f17fba3dc244117d83ac71a59452b772b5a"></a>

<a id="canonical-37ac860edc9e2dd6035b75c6e293d77922d8e39620b6e802284ae611e9dc4dd9"></a>

## namespace property — log_receiver / 14b8a72ac1ac / 5

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

<a id="canonical-e46c1a8ca6adda7e8881b362f2d5775c7c7915fc2d6186c8aa054c78eabfde68"></a>

<a id="canonical-7d7be36619780207295dd4e3980a84e8d35e5f7fa223c7c1ad1a703aa89ba564"></a>

## tenant property — log_receiver / 14b8a72ac1ac / 6

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

<a id="canonical-2dc0b17708a583fba2c0a1c41dfaf3ccde700e54f1de7db2904611de1e1daa20"></a>

## Next pages — log_receiver / 14b8a72ac1ac / 7

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-50a19fa81e3567c071c56c45969bd9631a50031ab64e99f68a3981b65f12af91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd44cf393e34c29e1440d2b31a50840c0fabbf01b4ad8ebf6a5a3084ac3e241b"></a>

## logs_streaming_disabled — logs_streaming_disabled / 4a96ff5ffc4b / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- logs_streaming_disabled

<a id="canonical-30b3f7d4db67b3f25649ff904e2c3b2d2053f711575a940ebb123752a2137172"></a>

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

<a id="canonical-f0abf924d2cb7e07ff58aebc9894a6ca79c5a0599daba8aa987305ca90bf47f5"></a>

## Direct properties — logs_streaming_disabled / 4a96ff5ffc4b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a85be9b464733aa85f874bc7558a4ba6b1d5b11bb1c0e013bfeba6ef07094f99"></a>

## Next pages — logs_streaming_disabled / 4a96ff5ffc4b / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-43ac5fc6c50ea5ab410ae17a760b8be1308b94745f377e19781f4caa5374b05b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f467213c3606ecfaf909de8cdd12b1bd8ab946995b6e63f3c1699b577110afa"></a>

## master_node_configuration — master_node_configuration / 1df7cbab6978 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- master_node_configuration

<a id="canonical-8ae5bf0496808b5209450b2894a0e6550fb9e6451fcfda4d5a865a77ea3ed54e"></a>

Type: `"list"`. Computed.

Master Nodes. Configuration of master nodes.

Upstream description:

Configuration of master nodes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3"
  }
}
```

<a id="canonical-f77f1329d0c34d93e5c3ed842e50bdfccef9ae1ce6a3d3ae65e5ca9bd7d49f7a"></a>

## Direct properties — master_node_configuration / 1df7cbab6978 / 3

<a id="canonical-75cc16afc48461a02dd6a9699900c3853a57272fc8d5f76cf65124bced4c9199"></a>

<a id="canonical-56e839999147ce805ecbfb19c33d54a5d8876f4ca9c4488cc90ac34bb40912e3"></a>

## name property — master_node_configuration / 1df7cbab6978 / 4

Type: `"string"`. Computed.

Name. Names of master node.

Upstream description:

Names of master node.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-717676d2cc8b1b24c835f9cddfde628f5defb9d397517a8aa4c278d2e1322625"></a>

<a id="canonical-aae3fe60a44b03df2f482a08535b407ba4f583c3a6ad6bcd127a7471c9127ebf"></a>

## public_ip property — master_node_configuration / 1df7cbab6978 / 5

Type: `"string"`. Computed.

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

Upstream description:

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-a3b9f5f3a40727d5f10aa760389ca070e8789870ac8a0cff06a8c0b2ae9641a3"></a>

## Next pages — master_node_configuration / 1df7cbab6978 / 6

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-1821b16b75027c2ba838b2ee0ece1cbc1bc9f397cf5a263538b839a6f53f64e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fef95394edbb9079df0c574b3f6201f017961273702442816fe9902ba638bb51"></a>

## no_bond_devices — no_bond_devices / 738bb01265b0 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- no_bond_devices

<a id="canonical-d3bbb04be164482b94d938225e7f9021cf021fd84a9502b406d1bc1be577f029"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no bond devices.

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

<a id="canonical-4b2fb27b8713c7b024f39184640459babaae7a2982797765f36a165bae37ea48"></a>

## Direct properties — no_bond_devices / 738bb01265b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-057fcaffd1e00689a7e75a4aef3dd7ae74be389805d96775c5db11e5d365f042"></a>

## Next pages — no_bond_devices / 738bb01265b0 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-95766e90509223e430f27a346b3dadade1ae3ee4386baa2a652f8b2a741a573c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17edc26111e9d1ba6128d91063af1cdce36d7079e8697723dca6303714aea404"></a>

## no_k8s_cluster — no_k8s_cluster / adfdc501a1f2 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- no_k8s_cluster

<a id="canonical-47846b7a5a3d2562db3735a7c11f54ecc01b755ee58d6a52febe907b0cc1ac26"></a>

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

<a id="canonical-fc35e82895086e2a87bac821568de4def3b2b28984eecfa2c89e8110322ac042"></a>

## Direct properties — no_k8s_cluster / adfdc501a1f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-55e5f5a4ba21f1e1a3e895c48ded7f64e689cb41848cd68561061533d2eb9460"></a>

## Next pages — no_k8s_cluster / adfdc501a1f2 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-79cf117dd5b3f3194bd236124ef6a31556e3c1805f4c2ab48064434c1a7b2bea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8fce5c40439372c07b3018427f69ab7c274782010fee0a598ae2708a7b7dc9d"></a>

## no_local_control_plane — no_local_control_plane / 49111029b725 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- no_local_control_plane

<a id="canonical-42060f8c584067f5f23d4abf4f298530c8561d1cb38c8edb19f7fa9a60e58992"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no local control plane.

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

<a id="canonical-4cb394b2fc7a44f96ee7ec91ec3af242a9432497a1a977359e008546e458a9e7"></a>

## Direct properties — no_local_control_plane / 49111029b725 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ef4640551febb45a60920ea81a5e18d038c74a957f127a06bfb306f835f667f9"></a>

## Next pages — no_local_control_plane / 49111029b725 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-c68631f4a2c69397e59b3b8bc1b369582b6bc95af9c269ac5b32da9a345ae28a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d69c729c5bdbbeadba43159b86f96660638b279dc03529c495ceeeb29abfaa05"></a>

## offline_survivability_mode — offline_survivability_mode / 5c1ae25ccb17 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- offline_survivability_mode

<a id="canonical-c474347b4b657c050214637fc85a863f4f64be113570e5c13cec004f90826a03"></a>

Type: `"single"`. Computed.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

<a id="canonical-cec81861e5d3f3d1d6f6cd52619455cb70f2efccc96bb4193d13b7d57fe53bab"></a>

## Direct properties — offline_survivability_mode / 5c1ae25ccb17 / 3

- [enable_offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-7d32de9826ed0cd096aeee41540f5516da064901eb46a5037979b1a94b1879c6): complete subsection reference.

- [no_offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-a7dd422b05f8aac8f4807bd406a4da294b1569f063303fda3ca165f3fc9d66eb): complete subsection reference.

<a id="canonical-6607389adf45dc4d1cd9163f1031fa5289a79329b72a1a44f7c477567e9cced0"></a>

## Next pages — offline_survivability_mode / 5c1ae25ccb17 / 4

- [offline_survivability_mode.enable_offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-7d32de9826ed0cd096aeee41540f5516da064901eb46a5037979b1a94b1879c6)
- [offline_survivability_mode.no_offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-a7dd422b05f8aac8f4807bd406a4da294b1569f063303fda3ca165f3fc9d66eb)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-7d32de9826ed0cd096aeee41540f5516da064901eb46a5037979b1a94b1879c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cfe8118191db53f33e17b61b4b3c9ea3bd30feaa05afab516344fe7564b34655"></a>

## offline_survivability_mode.enable_offline_survivability_mode — offline_survivability_mode.enable_offline_survivability_mode / 6004886f246f / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-c68631f4a2c69397e59b3b8bc1b369582b6bc95af9c269ac5b32da9a345ae28a)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-a22b4863207ef4f9b93aea27dfe987f166d783d90b8c9907a29cda4851b9b9e8"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable offline survivability mode.

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

<a id="canonical-b4f086747bf84a50d96c80c9383c908817d21cbbe8d8def3608f9c68d9d0bf16"></a>

## Direct properties — offline_survivability_mode.enable_offline_survivability_mode / 6004886f246f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3823bfaaad6c87f6514eb5db7eafd3c43c604b463a0b88bd1056aea682058f3f"></a>

## Next pages — offline_survivability_mode.enable_offline_survivability_mode / 6004886f246f / 4

- [offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-c68631f4a2c69397e59b3b8bc1b369582b6bc95af9c269ac5b32da9a345ae28a)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-a7dd422b05f8aac8f4807bd406a4da294b1569f063303fda3ca165f3fc9d66eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d3bd024f5d2e28269e1a3673cc8cefee82515e761f4ea6d7da554ff72e0af17"></a>

## offline_survivability_mode.no_offline_survivability_mode — offline_survivability_mode.no_offline_survivability_mode / 5f0d0abc7c8f / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-c68631f4a2c69397e59b3b8bc1b369582b6bc95af9c269ac5b32da9a345ae28a)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-180d3f968026cd4727f5ecb2e52f4adf4fac05fd4bf9086ccdbd30c16d9eac44"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no offline survivability mode.

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

<a id="canonical-672d3762580aa4606da1e3c7d8572e2f0a663dd4ab5b21bc410e8932462a96a0"></a>

## Direct properties — offline_survivability_mode.no_offline_survivability_mode / 5f0d0abc7c8f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7dc6bf2c3f3c4aca8803fe1f94b91699b6316f937baadaa3ca6787b2ab2f3723"></a>

## Next pages — offline_survivability_mode.no_offline_survivability_mode / 5f0d0abc7c8f / 4

- [offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-c68631f4a2c69397e59b3b8bc1b369582b6bc95af9c269ac5b32da9a345ae28a)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-d04a7772f57dd114a68678b0130f6613df5ef337a7444cc2a60f96dd1befc6cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c05e96b13edf64034263367e53ddb9252322392d208056c68b807964d1c9556"></a>

## os — os / f9cdab287e72 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- os

<a id="canonical-a2b30dfaa6904216e541c49ccfed15193c4e2093b201e246ebe48fac1f7f10d4"></a>

Type: `"single"`. Computed.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Upstream description:

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

<a id="canonical-45bfb6bd98781a3b01ead6d88d6ee5e3956356bf72531754e99b224dd82ca19b"></a>

## Direct properties — os / f9cdab287e72 / 3

- [default_os_version](data-sources--voltstack_site--reference--group-009.md#canonical-fc9e7c22a0bfb2dcb4b38b128ee2e662b31794ab408c142b5d39f21f6ed95499): complete subsection reference.

<a id="canonical-0cf9f0122594dabf5c1998eebecd681c1197cc07aea8297529a43b7b23a0cdfc"></a>

<a id="canonical-66d434eb8f4a9b5f50090cbfc0cface4669f1bd032bca478f20fe6a57658decc"></a>

## operating_system_version property — os / f9cdab287e72 / 4

Type: `"string"`. Computed.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Upstream description:

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-f66c67abfaf143549c0557b4aacf28ff0fe2d9af0e09d83a396360610705dacd"></a>

## Next pages — os / f9cdab287e72 / 5

- [os.default_os_version](data-sources--voltstack_site--reference--group-009.md#canonical-fc9e7c22a0bfb2dcb4b38b128ee2e662b31794ab408c142b5d39f21f6ed95499)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-fc9e7c22a0bfb2dcb4b38b128ee2e662b31794ab408c142b5d39f21f6ed95499"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59af6cda489ea50f3c13cbfa317cb5ffb121ab28bc4c5d1318a6f9d506603431"></a>

## os.default_os_version — os.default_os_version / b4699830655c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [os](data-sources--voltstack_site--reference--group-009.md#canonical-d04a7772f57dd114a68678b0130f6613df5ef337a7444cc2a60f96dd1befc6cb)
- os.default_os_version

<a id="canonical-fa02d02c3a9c20da3b5152962d9f7c5d4c42a29cd2ac3642a63394e2a1ffd3d5"></a>

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

<a id="canonical-00b6449c8a40dfef05236786b5d0046340ccc46f078e286b6084fa9589702753"></a>

## Direct properties — os.default_os_version / b4699830655c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dc80e36a791acc44a576161b0b361fdf5e1968f585ffdf5fc96233592bf77777"></a>

## Next pages — os.default_os_version / b4699830655c / 4

- [os](data-sources--voltstack_site--reference--group-009.md#canonical-d04a7772f57dd114a68678b0130f6613df5ef337a7444cc2a60f96dd1befc6cb)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-539f39e53540d3d5f2b1c575b3bbfd381fd22d7eb5aa117bd8a5660f05d81ad0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65726c1b10716c15769317f1accfd992ca55a61aab53d3f8316df553ac4a98a4"></a>

## sriov_interfaces — sriov_interfaces / 5f0dc85673d7 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- sriov_interfaces

<a id="canonical-9c86d9f36c0cf58827dd163383ee5516471e5dfc8690ccd1d0e8ba362ef7204b"></a>

Type: `"single"`. Computed.

List of all custom SR-IOV interfaces configuration.

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

<a id="canonical-5b04ee953ff4ae492655ec7725a56365dc25c26e40b8968e27b5cb442d357d9e"></a>

## Direct properties — sriov_interfaces / 5f0dc85673d7 / 3

- [sriov_interface](data-sources--voltstack_site--reference--group-009.md#canonical-5867e04f37d1a0ec8f5fc77d9a85e25e4677e56e6fc07188b6fec70557d7e96b): complete subsection reference.

<a id="canonical-a197659fe632450f71d9d8c6dc5f252e51d6d967f707cc1bfc3248329589cc61"></a>

## Next pages — sriov_interfaces / 5f0dc85673d7 / 4

- [sriov_interfaces.sriov_interface](data-sources--voltstack_site--reference--group-009.md#canonical-5867e04f37d1a0ec8f5fc77d9a85e25e4677e56e6fc07188b6fec70557d7e96b)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-5867e04f37d1a0ec8f5fc77d9a85e25e4677e56e6fc07188b6fec70557d7e96b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fba3d0eb31c53428e3e56fa09d9fe612a62562a4d4a466b72f68dc290a08f2e4"></a>

## sriov_interfaces.sriov_interface — sriov_interfaces.sriov_interface / 4fb316378637 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [sriov_interfaces](data-sources--voltstack_site--reference--group-009.md#canonical-539f39e53540d3d5f2b1c575b3bbfd381fd22d7eb5aa117bd8a5660f05d81ad0)
- sriov_interfaces.sriov_interface

<a id="canonical-ca926611edf8f39223915ce2a992665ba36bd08c7f00e293ab73399d9f68089f"></a>

Type: `"list"`. Computed.

Use custom SR-IOV interfaces Configuration.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-d278cabab964909c7ae01187d917e53a297db48c4360b7c0c7657c58ccb43269"></a>

## Direct properties — sriov_interfaces.sriov_interface / 4fb316378637 / 3

<a id="canonical-c607699346eb630e93b2069d94262c629b24cb897e74b6444fcf8cb10af5d60f"></a>

<a id="canonical-7da9e23c91548b1fc2caed686635310d5068b010e3e566f89cbc6bd1171d8338"></a>

## interface_name property — sriov_interfaces.sriov_interface / 4fb316378637 / 4

Type: `"string"`. Computed.

Name of physical interface. Name of SR-IOV physical interface.

Upstream description:

Name of SR-IOV physical interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-a4d50c0fff67a89e9fda3b128830c8570dea10fd45798869cd8f7ecdccca9e28"></a>

<a id="canonical-79381cf3828df39967095cc7c471593ffebf02f39c1f2d1089e2276d38fb5f4e"></a>

## number_of_vfio_vfs property — sriov_interfaces.sriov_interface / 4fb316378637 / 5

Type: `"number"`. Computed.

Number of virtual functions reserved for VNFs and DPDK-based CNFs.

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

<a id="canonical-bf65a46000e1997d98ed095ab8a2a0765941e1e8b6adf8dae0066a36e7d53bc7"></a>

<a id="canonical-38296b12f5336db45a3e8e0179612bfce44db76f5ef218378c2a6b79aa53a28e"></a>

## number_of_vfs property — sriov_interfaces.sriov_interface / 4fb316378637 / 6

Type: `"number"`. Computed.

Total number of virtual functions. Total number of virtual functions.

Upstream description:

Total number of virtual functions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-89da9bca34d4c29112687d030b64131fab8b52bb107348dab4be2d8a93177147"></a>

## Next pages — sriov_interfaces.sriov_interface / 4fb316378637 / 7

- [sriov_interfaces](data-sources--voltstack_site--reference--group-009.md#canonical-539f39e53540d3d5f2b1c575b3bbfd381fd22d7eb5aa117bd8a5660f05d81ad0)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-dfb31ba96b2d5a71625a57f3a22d7d04c58f53983362297d97cbd10e96c73ef6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e92a1fe65531709b4b00e13e3bc3e692d350ecf3477c680d4414d61a859c48f4"></a>

## sw — sw / 0a57aafc4647 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- sw

<a id="canonical-9db2c6dced2c209a472809f003bc9d2abf5c76114931730a7a4c029c1817e4ef"></a>

Type: `"single"`. Computed.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Upstream description:

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

<a id="canonical-f19bbc98ec52f4fb58e33148ef153e60cd17b6a1ad9444cf98f508d21b26639d"></a>

## Direct properties — sw / 0a57aafc4647 / 3

- [default_sw_version](data-sources--voltstack_site--reference--group-009.md#canonical-fb00cd8f70c0dcd430a9e17e908daf169d4c06ccf08f2c07436752ffa7f7d898): complete subsection reference.

<a id="canonical-66a31819c8e6aa5cb0feed1402c83c5640da19e5a1e58e08e0fbe0f23c8ce25e"></a>

<a id="canonical-02b564c17304ebfc6bac9e747ced02e5d74f0fe1535075e7acd6aeff58d3d1e9"></a>

## volterra_software_version property — sw / 0a57aafc4647 / 4

Type: `"string"`. Computed.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-0bf27209442d618cfea59485dbe522364db4d898677b99a5f41c0e733a98309b"></a>

## Next pages — sw / 0a57aafc4647 / 5

- [sw.default_sw_version](data-sources--voltstack_site--reference--group-009.md#canonical-fb00cd8f70c0dcd430a9e17e908daf169d4c06ccf08f2c07436752ffa7f7d898)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-fb00cd8f70c0dcd430a9e17e908daf169d4c06ccf08f2c07436752ffa7f7d898"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8aa84b8b5b69d01e841835cd84c0c8f02272143b99053c74f6d93d48ce46b2d7"></a>

## sw.default_sw_version — sw.default_sw_version / 22e05660a212 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [sw](data-sources--voltstack_site--reference--group-009.md#canonical-dfb31ba96b2d5a71625a57f3a22d7d04c58f53983362297d97cbd10e96c73ef6)
- sw.default_sw_version

<a id="canonical-fc844a06c23ee7903b91905eb43fbc2564bfeb83d28bb788fb1c0e8ad52ac741"></a>

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

<a id="canonical-9a70843fe681eec9b91a0c88d2e8ebb74d2c83bd427a2df97ecfb26d73e8bb72"></a>

## Direct properties — sw.default_sw_version / 22e05660a212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c8f060ce5be89a0e49cee4aa964d2a1e113c6bd30c7334ff1456bcc9f7dadbb"></a>

## Next pages — sw.default_sw_version / 22e05660a212 / 4

- [sw](data-sources--voltstack_site--reference--group-009.md#canonical-dfb31ba96b2d5a71625a57f3a22d7d04c58f53983362297d97cbd10e96c73ef6)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-6b3bcfe958943ca6de6f5324892b713b60459ce829b7164c23fba924583017d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5530d712ed9330e44bd3796c7bdd6a96706b688371731c70c86571d77acc77b"></a>

## usb_policy — usb_policy / 2ddacc7d7102 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- usb_policy

<a id="canonical-1fd0bab96bb193e79f7e07ace7e830eabb3a33f41aed28accaf9f5823c470462"></a>

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

<a id="canonical-6388d28d4ff60352273e8cca351b12f4d616388d525a63913f29230ff58f6dda"></a>

## Direct properties — usb_policy / 2ddacc7d7102 / 3

<a id="canonical-3f7ed2edbcf15f3342fec5e9e9ae6d0aedb8006ffe704c22936e38c85175da64"></a>

<a id="canonical-843504dba07c9a23c29aeca6ee54f4d61efb225e64b68dfb9cc5a3b9f10c428e"></a>

## name property — usb_policy / 2ddacc7d7102 / 4

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

<a id="canonical-81528cbc08d40458649317bf1c5307a1cef0cd71a8b97fa3b6f60c6be966862c"></a>

<a id="canonical-6fef58114ecf6214d338bb3839c7b84450675f3918ea8cd67934e4a6751e7159"></a>

## namespace property — usb_policy / 2ddacc7d7102 / 5

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

<a id="canonical-280f63fcd2037a01e4eba3e1c7be1a090ced401929d976ee2a051fd68240ab33"></a>

<a id="canonical-e626630572e986e12d625d9a9d0a61d8a253b19ae841fa0c4feacbfdc5b8e488"></a>

## tenant property — usb_policy / 2ddacc7d7102 / 6

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

<a id="canonical-7147a8bab06fa5ac00447f7e189cab8186616832e1adb7fac562bd41b4a477ce"></a>

## Next pages — usb_policy / 2ddacc7d7102 / 7

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-5d176cdf4bc826837f7e0219c9acc47d6a8b24b94f26d46b2b52202c2d0463e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6381f4841eb4787d04757149b9c35dc7c60b9c72a02aff1284b30277dc059031"></a>

## waf_signatures — waf_signatures / 20a897682139 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- waf_signatures

<a id="canonical-66832a71e31def37e0daaff4fc46b04451dc7f445de00c0fc7200d6049865e70"></a>

Type: `"single"`. Computed.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

<a id="canonical-bed86bfe5031922186840b55ddfdfa151aa5caa7f8dcb5f09fbd673c929bba1a"></a>

## Direct properties — waf_signatures / 20a897682139 / 3

- [automatic](data-sources--voltstack_site--reference--group-009.md#canonical-3b5accfd409bbaa8405cb12e54cb489cb685c90c8ea90908029031f3904e2a07): complete subsection reference.

- [manual](data-sources--voltstack_site--reference--group-009.md#canonical-f44cfd5883844ecc5e2d2fe0d4056f10d0414625df5ed41cc9fea6783038ee0d): complete subsection reference.

<a id="canonical-2a4b205fee588c866712633a29f4df94eafb4706d203b2eca4b7706bf04dc427"></a>

## Next pages — waf_signatures / 20a897682139 / 4

- [waf_signatures.automatic](data-sources--voltstack_site--reference--group-009.md#canonical-3b5accfd409bbaa8405cb12e54cb489cb685c90c8ea90908029031f3904e2a07)
- [waf_signatures.manual](data-sources--voltstack_site--reference--group-009.md#canonical-f44cfd5883844ecc5e2d2fe0d4056f10d0414625df5ed41cc9fea6783038ee0d)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-3b5accfd409bbaa8405cb12e54cb489cb685c90c8ea90908029031f3904e2a07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1a0a20356253c04253c88a70573e826799cca393602b9d081a715d42ef32c96"></a>

## waf_signatures.automatic — waf_signatures.automatic / da91215d80e4 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [waf_signatures](data-sources--voltstack_site--reference--group-009.md#canonical-5d176cdf4bc826837f7e0219c9acc47d6a8b24b94f26d46b2b52202c2d0463e7)
- waf_signatures.automatic

<a id="canonical-9bf8cec9a6b023305d01d213b7604c2ccbc61ba40d1d56cfc74f545f64418087"></a>

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

<a id="canonical-fd3ad9f5a0a9068c7f5bd179dabcc3058e769d0afd06378be1f35f10dcce8379"></a>

## Direct properties — waf_signatures.automatic / da91215d80e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ac0e2a58487219a39a730c6746dc250bf404ec1d64786b564e2c6201da9354de"></a>

## Next pages — waf_signatures.automatic / da91215d80e4 / 4

- [waf_signatures](data-sources--voltstack_site--reference--group-009.md#canonical-5d176cdf4bc826837f7e0219c9acc47d6a8b24b94f26d46b2b52202c2d0463e7)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-f44cfd5883844ecc5e2d2fe0d4056f10d0414625df5ed41cc9fea6783038ee0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ab72bc3f7f1d74a7f6ebb328504a7bc3c3f8eb350724883baa87cdccd4cbea1"></a>

## waf_signatures.manual — waf_signatures.manual / 313e740abf09 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [waf_signatures](data-sources--voltstack_site--reference--group-009.md#canonical-5d176cdf4bc826837f7e0219c9acc47d6a8b24b94f26d46b2b52202c2d0463e7)
- waf_signatures.manual

<a id="canonical-62c65b13ed7b043c577eaf0a3ee29d9ece927d8fb39ab04abbe24200dfe7033c"></a>

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

<a id="canonical-9d3a4534c54ec15bd8d162bfb72403d20e0f4d781cb347722f4fbcb2c623f7cd"></a>

## Direct properties — waf_signatures.manual / 313e740abf09 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b21882481c05f60f6b342ae8fd5d0e8c831e475a01e10840941686024eab14dd"></a>

## Next pages — waf_signatures.manual / 313e740abf09 / 4

- [waf_signatures](data-sources--voltstack_site--reference--group-009.md#canonical-5d176cdf4bc826837f7e0219c9acc47d6a8b24b94f26d46b2b52202c2d0463e7)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
