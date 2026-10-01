---
page_title: "xcsh_securemesh_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site reference."
---

# xcsh_securemesh_site reference

<a id="canonical-e1d414e59cee8cddaa62c498ce2914b35f428b9d941e8cf7154cfc15b199fe2f"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 1e51c7d9b571 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--securemesh_site--reference--group-002.md#canonical-584cd07eb5c1bbc7a632c828fef51e2438c72be6f2a6e6b50b49da3a400cd8fb)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--securemesh_site--reference--group-002.md#canonical-413c915394c3c08864c60e94b96381707de2da6d13ec426144304db4f2d21184)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--securemesh_site--reference--group-002.md#canonical-f7ce949eeef3aa751c3814fdf4cca4f49a0e04eee1b224a7ff0f22bcc1aff379)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site--reference--group-002.md#canonical-726f3548ce632766f6ecf5a475a42a4a7e01367414705cd7c2dc26bf4fb62918)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-7f01decd91d1b1087d101dd9015eb34c8a5a45ad51f985d2ba30765318b799f0"></a>

Type: `"list"`. Computed.

List of non overlapping IP address ranges.

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
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0b3d5cf33e13355611c3cc9ca827ecd00b75f6346a41e1a433d3ab68b10f0448"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 1e51c7d9b571 / 3

<a id="canonical-bff4039de2f9bd15ec00be8626c2f3e5e1a851b23e32159bd783d56e0610fcaa"></a>

<a id="canonical-fed25c3ea321fd98ae9d5e9a74ce2f8c8de8205fc0870dab0d702447d0ecd946"></a>

## end_ip property — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 1e51c7d9b571 / 4

Type: `"string"`. Computed.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

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

<a id="canonical-a3b36d5a23f8fbfc599d7d5feb8191e3778613e1674246b46e54369ed6081d6d"></a>

<a id="canonical-b854020f4421abebf79702ab9611dc37ec15231d13cc2c732f5c1c7b991be5b0"></a>

## start_ip property — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 1e51c7d9b571 / 5

Type: `"string"`. Computed.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

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

<a id="canonical-364f71096e55ab3edb6239b140d3aa8169d299f08f673e14bddda2e86eecab07"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 1e51c7d9b571 / 6

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site--reference--group-002.md#canonical-726f3548ce632766f6ecf5a475a42a4a7e01367414705cd7c2dc26bf4fb62918)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-52559152c44e304e146edf28973007ad316ab7b965c97b2235a542bd3c34c360"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7720f00a586ba95b567a2a95e50d91285d4e09e6d4aecd182d2690d7f77b58e"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 09075d8b0b51 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--securemesh_site--reference--group-002.md#canonical-584cd07eb5c1bbc7a632c828fef51e2438c72be6f2a6e6b50b49da3a400cd8fb)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--securemesh_site--reference--group-002.md#canonical-413c915394c3c08864c60e94b96381707de2da6d13ec426144304db4f2d21184)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--securemesh_site--reference--group-002.md#canonical-f7ce949eeef3aa751c3814fdf4cca4f49a0e04eee1b224a7ff0f22bcc1aff379)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-3275d37cb0fc1589acfb0f55300d3aa7cb4d58da66957a3467c91c2bd2b773ea"></a>

Type: `"single"`. Computed.

Map of Interface IPv6 assignments per node.

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

<a id="canonical-ef41b1ce776c46b4ce767312211fe34bf8e05ca715d0f5d6bc57a284413a4f37"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 09075d8b0b51 / 3

<a id="canonical-693205ba911acb04c2189d205e5ef12b23afbd1e653e54ee084f0bbd83272157"></a>

<a id="canonical-fd3261eec4e25b0ca1e44aab9d4df6aff9d437352c2b69ade84f27d25e049999"></a>

## interface_ip_map property — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 09075d8b0b51 / 4

Type: `["map", "string"]`. Computed.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

<a id="canonical-e67e1a4c51cfefbf3fa4ac31f1116ab4d7b53699516adbdff442071ff383851e"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 09075d8b0b51 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--securemesh_site--reference--group-002.md#canonical-f7ce949eeef3aa751c3814fdf4cca4f49a0e04eee1b224a7ff0f22bcc1aff379)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-804b5013c613bb739643517e79be3ce887e293a24115d985cf2652d00f2ed6ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8e4602b75b94127ffc929c77d98cd08e543123b1319997cfe15805088eea2d3"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.is_primary — custom_network_config.interface_list.interfaces.ethernet_interface.is_primary / d1e765d84bb8 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- custom_network_config.interface_list.interfaces.ethernet_interface.is_primary

<a id="canonical-d9a0c9bddada38952ddea537b0f1e34a4b815df6315e3f1efc5298d236c77d55"></a>

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

<a id="canonical-b9a9e7b08760b53feee7d7ae2f618389b4684ee44d21edb97f7b754b275fb4d3"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.is_primary / d1e765d84bb8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-355e46a477a8315fc7a4bc6d2fd28ea4fdf287e62c5b06135645c616e5e401a7"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.is_primary / d1e765d84bb8 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-1067f8106c50e6b12647e2e6a0cb8fa4e5ca81d01dc636730973bac9eb102196"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9755b0f4cd69017df354e6e44e34a337eb9b136bea042d3c6cce30b69470e1af"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.monitor — custom_network_config.interface_list.interfaces.ethernet_interface.monitor / 058579eb65f8 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- custom_network_config.interface_list.interfaces.ethernet_interface.monitor

<a id="canonical-83976c7edea019acacc2bd8d4ee018571ab5e7d9553c7dc163c7496ac98f7ea3"></a>

Type: `["object", {}]`. Computed.

Link Quality Monitoring configuration for a network interface.

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

<a id="canonical-96797bc0fe26c1c8bb39e80c39545f1f845c9c1986a07739c8f42fbf6b2b927f"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.monitor / 058579eb65f8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa65e7008c9232c9579f245e20756a1086f59be77f369d1c615e9533417d11ca"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.monitor / 058579eb65f8 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-c72c2de4d802defa4b84de12bf585d3a000dfc608500356f19e8a1d7105723ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab7887c22f89150672d0615ae39cb4af290f989b2d2389f98cf4973909d88832"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled — custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disab / 578f3f263c88 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled

<a id="canonical-4f9cb9f5d9977d35f0dbfd99870317eab9fd37e470d291d0966d48d106d89534"></a>

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

<a id="canonical-014aab5f3a6c066911c370c36ceccf02b856da9536c7039acf813d101f048158"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disab / 578f3f263c88 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-97b60fa9183768bc44dc61b7789407d74f9adaac4ad61b124380e2a4011d2d3e"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disab / 578f3f263c88 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-092c1c708d20866c458b826c0047f322eaec05e88e8db15ac5f0c648fac17003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-203e086be39a489d1f056e1fac448bc9f15b14c7f1b611efaa4ca5385b1f949e"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address — custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_addre / 44b1844b9608 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address

<a id="canonical-e3ed183dad80ca56240a802e6b25538c39fe3df5eb05bb31d717e682fdbb5682"></a>

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

<a id="canonical-48aa22860c7788281eccb87d1723b9f66ec42aa1c062d0f1f35dd140b770a909"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_addre / 44b1844b9608 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b922982d6d600cfca8c28b424f05805bef638144d980067d09ba6ad13473b431"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_addre / 44b1844b9608 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-9f5c68214bb6e2c41dfa34f582a0fb5a82d618c7e53963665fe7c8885520ca96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02aeb8bb8ce9a608c31efae35d5fb0210953252c180b9ea75cf34e266971d5df"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.not_primary — custom_network_config.interface_list.interfaces.ethernet_interface.not_primary / 691b53910a5d / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- custom_network_config.interface_list.interfaces.ethernet_interface.not_primary

<a id="canonical-ae4bb61206cb86ad080c93baa19c9863af7feaf4bdd9d70bcd6b2cd46aa0700f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for not primary.

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

<a id="canonical-992ec51eb651caa331ed4f74483a5825659da69253d02c5fb117caf3b0a66965"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.not_primary / 691b53910a5d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b7d539cbf8237cf0b63fc3a1d9be246b279935326c90b434b23b2e96c7ac7fd3"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.not_primary / 691b53910a5d / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-35b3b7b81256e74344aad9ca69eac464d25901d7e25a4bff305a861eb97ccba2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0f92dcf85746b06c8eb8402672acf6b4a005df76bd9d2aeb35470776b3b4b03"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network — custom_network_config.interface_list.interfaces.ethernet_interface.site_local_in / 7949accce49e / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network

<a id="canonical-eb74b147827f3f01c38881184945f081d5e6b75acbc0905bc4d5baf5ea32dff9"></a>

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

<a id="canonical-84b5f5c8900acce04f92cbda41ab1a7aecd4884fd904f97ed126c6ef4cba5969"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.site_local_in / 7949accce49e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-de570f9f313cba704ff09efe577d88aec1744f8b8a72e4e9acabea29b58c56a8"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.site_local_in / 7949accce49e / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-f637a30e5a24ac4c9cdd1c00a8196e4c2c2c9c33d961b92d992ee650adf0ce72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5701a2600ab86cb1090bd51165992bb1aed459d9b3e2140aefbe3ad605a4c472"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network — custom_network_config.interface_list.interfaces.ethernet_interface.site_local_ne / 6aeefb4f818f / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network

<a id="canonical-584fe5d36c3fe5c93468ac51223c96471f5af955286f8902263f8877e1796f6a"></a>

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

<a id="canonical-3dfccfd72e6d4e6ab0c4e3a77b8a8d05f0a5aa3bc90b9f770ebfbe0f1e0e6d2f"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.site_local_ne / 6aeefb4f818f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-751cc0f70b308d48d615d11782e86e171963279b9a9d9d1913569afd12d07dc1"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.site_local_ne / 6aeefb4f818f / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-185c2dcc7aca0542f14763c642c361508c3e65c56177f8b0767314328a64a026"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eff84215541a9aa0fa8525db7b2535c1db9d4f1b9d70d9e94f8d3beb55fb7629"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip / 3f4385621214 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip

<a id="canonical-3e9be8938e5d2e7109c95df009ae5874e0c0289766f8fe88bfc8348e2741d576"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

<a id="canonical-c4f5baada5f2d9f508a62bd8de7f0c34ac4b0336d73a9d502ec55b48d8870809"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip / 3f4385621214 / 3

- [cluster_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-d1675f9e5bbeba3ecc596a51f65238baa00082f2096886fbdd9b2cd4ec40179f): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-11855064e53b6d74c883057128dcb303f492579c2ad56e7dcecde5c721a314a9): complete subsection reference.

<a id="canonical-d020618f0ed5992ea637ce9ae07baeeceea171525745313080ff5c7dfbd1a54f"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip / 3f4385621214 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-d1675f9e5bbeba3ecc596a51f65238baa00082f2096886fbdd9b2cd4ec40179f)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-11855064e53b6d74c883057128dcb303f492579c2ad56e7dcecde5c721a314a9)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-d1675f9e5bbeba3ecc596a51f65238baa00082f2096886fbdd9b2cd4ec40179f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5603cdab9cb6c3da98affd12403030bf57d01bbcc45685b33fda6901ef228a58"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.clu / 18a8170b7fc2 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-185c2dcc7aca0542f14763c642c361508c3e65c56177f8b0767314328a64a026)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip

<a id="canonical-ab830cd9c1254e56d774087e21e47eda3b308a09710fa1c368cc0493e80ab883"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for cluster.

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

<a id="canonical-5cdce7156df64fe1731ddb5416aa712d1c224bb43ad53cacfd04d5d4d37be369"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.clu / 18a8170b7fc2 / 3

<a id="canonical-255ddd6acc5bb8e732e90b019097bb5e20e7f99e61b3eff96d0d24960f6b2f1f"></a>

<a id="canonical-926378ab61d216c1d051991869522c70e598d22c7d169743439e6398b18c7662"></a>

## interface_ip_map property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.clu / 18a8170b7fc2 / 4

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-5f6ce22839ac619972ae296e4fef37a69ef907aee596a89f11227b836241bfa0"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.clu / 18a8170b7fc2 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-185c2dcc7aca0542f14763c642c361508c3e65c56177f8b0767314328a64a026)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-11855064e53b6d74c883057128dcb303f492579c2ad56e7dcecde5c721a314a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04bb9e38ad0a870f69c33d7b7520260a6805ccc88492f375394b74f2c4be44ad"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.nod / 27b3c766d17a / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-185c2dcc7aca0542f14763c642c361508c3e65c56177f8b0767314328a64a026)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip

<a id="canonical-c97b4054e6ce2a36196d8c9d9c124b97395a4d3c6d142fea1c3706c617bb6671"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-336dd1e5889a9c13ede2ed02dd2991d171204afef2f5fbf0b771e9f60e84ef84"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.nod / 27b3c766d17a / 3

<a id="canonical-2d4525542c8a80b86a356439dd0bc4f1f5151125f999674516b7e1b8e86c8b80"></a>

<a id="canonical-e31d4582987611e5e216ac7286b6a68940c14d91c5143853216d82bb80bde4ef"></a>

## default_gw property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.nod / 27b3c766d17a / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-d88bad542174292a13375003500f3e2a5b8d1b50b17917b871bf58df9514d933"></a>

<a id="canonical-ce84cd9c0153d1c388346355bac546868da809ea6c8b3185e4a0e958f28014ed"></a>

## dns_server property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.nod / 27b3c766d17a / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-9ddebae7912aa71a6c7af82015f648d6ee92f33202f0e25937a1a379e28c5922"></a>

<a id="canonical-9b49278cc4dad700545b1c773b3dbccb4c63d2d809f2164183d8da5c088b7d1e"></a>

## ip_address property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.nod / 27b3c766d17a / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-d64f8e6748dec51c1ff2781f3c63a580c404889617e2e05b99cc4f5cbdea7594"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.nod / 27b3c766d17a / 7

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-185c2dcc7aca0542f14763c642c361508c3e65c56177f8b0767314328a64a026)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-e1863eb24e789ac6c1cec91fa1feab74111f52b96f759e7c9cbbb8e0dd13648d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f52849546d925d769696dbb9f2e666698b6177436b18a362305c19bfa0ebb3a"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 65d0d63b4fbb / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address

<a id="canonical-76309463a9eb4288cc639edd4c97f01b6bfdbc7db497a197f6890df3e5f8183a"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

<a id="canonical-69c4d2f25e695318fa4fd500b620982f0763d2beb03c94f4aae22882ded9b143"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 65d0d63b4fbb / 3

- [cluster_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-403e04207086b9187e661237fc03f753e94b6021751576f226d8c8c3bbe4b984): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-afa2c8233948a7a8e0a4858579fee8a7c70962d136c086df1d48c396dfbf0b18): complete subsection reference.

<a id="canonical-fe5ac8389fe8da6a2ad0c3d59a0970e95a283a542930704d1fff9eefeff37025"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 65d0d63b4fbb / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-403e04207086b9187e661237fc03f753e94b6021751576f226d8c8c3bbe4b984)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-afa2c8233948a7a8e0a4858579fee8a7c70962d136c086df1d48c396dfbf0b18)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-403e04207086b9187e661237fc03f753e94b6021751576f226d8c8c3bbe4b984"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31bcfadeedd2942f5a764c7e5ceea71ea8bc4905fa753db059719c097d7bf1b1"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / b965ab2e3c95 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--securemesh_site--reference--group-003.md#canonical-e1863eb24e789ac6c1cec91fa1feab74111f52b96f759e7c9cbbb8e0dd13648d)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip

<a id="canonical-50344c9373104dedf3571a6c7641ad691725e4a22cd2655b37be84ebcc72412b"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for cluster.

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

<a id="canonical-48c2925d2ebd6ddf9ff42fc2436c7505c22a65e812f600d9ea3a202bd5269c0b"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / b965ab2e3c95 / 3

<a id="canonical-6900d66811c7d895cee10ea5f5b53670f3b8023285fae310400bc0dc333de83d"></a>

<a id="canonical-b8dd35d7d9fc85af77f43a92e00dc1274613fe52fe885efbe5e0fa26e6260627"></a>

## interface_ip_map property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / b965ab2e3c95 / 4

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-4d76621ce2e1565237cb9b1c843162fa34f3ae4d58ac74c3fa9446e7f1f2ad17"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / b965ab2e3c95 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--securemesh_site--reference--group-003.md#canonical-e1863eb24e789ac6c1cec91fa1feab74111f52b96f759e7c9cbbb8e0dd13648d)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-afa2c8233948a7a8e0a4858579fee8a7c70962d136c086df1d48c396dfbf0b18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bed32c262b5b3dae57c56f2c92bd16564d421072ac857fb6158b5ef12fcdc28d"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / f0296c32ba76 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--securemesh_site--reference--group-003.md#canonical-e1863eb24e789ac6c1cec91fa1feab74111f52b96f759e7c9cbbb8e0dd13648d)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip

<a id="canonical-a1788d446f910e9178656fab814d001e9b9edf0fa4d5f5b5bcccbb40e3f20912"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-cceaee524c5851f9a77c461ae186e0744c98cbe265bd0521d3722dd0bbaf701c"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / f0296c32ba76 / 3

<a id="canonical-483587d8a2032d6c121981a6d50399a6853ada690a2d38d4a3fce0e5372ff6af"></a>

<a id="canonical-64e0f59bc29ac4588ec28be9148f4650884112df08e799c21196d3ddd9609a72"></a>

## default_gw property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / f0296c32ba76 / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-682a1d427b71e65cd86bf97a03127068a5ea380cdf95fc9d1ddea9d0735f9bda"></a>

<a id="canonical-1f021b74f91dd2f27b1fa618b52e894caa5dbe94980adb2b3354430abca1bcf9"></a>

## dns_server property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / f0296c32ba76 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-5864303d62f409df8b2bea6b1c34d083d31bba79e88cdcfaed47e41022567022"></a>

<a id="canonical-ad401b02092e4d9d370a71ffcd6c6a19dfe06bbf0eadb375e90600bf08dfc8b5"></a>

## ip_address property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / f0296c32ba76 / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-8151aab0428aa421a11211fc227ccbc8abfc0a78b2722f39e8930c0a9043bdf8"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / f0296c32ba76 / 7

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--securemesh_site--reference--group-003.md#canonical-e1863eb24e789ac6c1cec91fa1feab74111f52b96f759e7c9cbbb8e0dd13648d)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-dd4d4e980caa5b08349fa08f1b240421f301d4d69ce6e42505ac0b74a2e977a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e48a308eea92f8c13abade7524ce3b224bbc4e8532bc26d8ade7f29a192a080"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.storage_network — custom_network_config.interface_list.interfaces.ethernet_interface.storage_netwo / 0d462dd5bd2a / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- custom_network_config.interface_list.interfaces.ethernet_interface.storage_network

<a id="canonical-4a4336861767b4f82e51e07aa8c6c8e49ac8a36e456ebf34f7c40d2d64e222c2"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for storage network.

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

<a id="canonical-f103f5c88ece13eac5ff7fff1b4e7e71efd9fd23261ad5e08752940f0a004a8e"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.storage_netwo / 0d462dd5bd2a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6a275a292b701205ab56fb77623d86419770c12ad6d2c666e5a679ac40234eaa"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.storage_netwo / 0d462dd5bd2a / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-b098988ed0dd58570ac269abc56b98f3766613fab734daea08df821efd9f3a24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d0ee6b7188b23a645606344950d170a840013b5c1027a894d7d60902848f578"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.untagged — custom_network_config.interface_list.interfaces.ethernet_interface.untagged / b4a05f55af27 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-5acccc604aea7b50995107b9e397f7f48f31410fa282ff460cef6d3b3b719729)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- custom_network_config.interface_list.interfaces.ethernet_interface.untagged

<a id="canonical-9c01f47a493f6625846b721e638783890116d194794269588936503a1ccfde40"></a>

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

<a id="canonical-0dfb0ac81d83b476270d0c73ff1badb2704ba7c8778f4979ad0fd21d35ace23a"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.untagged / b4a05f55af27 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-44b9ab6b9eeea4ee4a983c04a13fefc5b299cf335772fb8005fe0e759865b9a4"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.untagged / b4a05f55af27 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-fb4ca6a7b69d3992eb832b33b4e1577b23ff60d9ae420fcc8b0662e799fb7279)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-a5790b2e3e396057bdf6d186f8e82a3b184aac39d74624bbbab99441acfe18d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-998982d0c58958f1a98e0c999936c6011bcfa4ebfe3d57ff5a569f120b1d2983"></a>

## custom_network_config.no_forward_proxy — custom_network_config.no_forward_proxy / 8a2537886709 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- custom_network_config.no_forward_proxy

<a id="canonical-c24c3f389db003ceba4862f4371f5d8e68034899f1f7d78128de7ecbed1a24e7"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no forward proxy.

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

<a id="canonical-b215cc60ad09c5881904a5ce75f9530a1ef7680940fc376ccd1a703e56d92002"></a>

## Direct properties — custom_network_config.no_forward_proxy / 8a2537886709 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65a5cebec641674eb7b42478118edddb5c8206e03d900d3d604981cd6d5217a4"></a>

## Next pages — custom_network_config.no_forward_proxy / 8a2537886709 / 4

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-e3b6fd24f6cd0da5263135b5a51a86c63437e29faacb665a8ac4fb03dbc49a79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b66cb12a9cf83f7b8761ed513e6046bb911da13dd587c78f687bed723f86f26e"></a>

## custom_network_config.no_global_network — custom_network_config.no_global_network / 5b16ea71d8dc / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- custom_network_config.no_global_network

<a id="canonical-55f14b72bf209f1a5f886ad8a2b75c5ffd770074f68933fb6a72f616d6b2c421"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no global network.

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

<a id="canonical-e00fec736b8d5a294629b7eaa3e89b9eddaf85100fd022fbf04e2b704d20c466"></a>

## Direct properties — custom_network_config.no_global_network / 5b16ea71d8dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-84c77a069a680d0d3d51c655c27a310636eea06b8370a126060e5212ca71de1e"></a>

## Next pages — custom_network_config.no_global_network / 5b16ea71d8dc / 4

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-25609b502097505129f9ca46cf6a92e8033f27ba9ed7e85f880578f8430ee0ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78033528e3ca1f467f9deb86c6ff33867dec29c500791bd1a8595adc112f4daf"></a>

## custom_network_config.no_network_policy — custom_network_config.no_network_policy / edb927c223eb / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- custom_network_config.no_network_policy

<a id="canonical-f1b53cebcb3feb816b10b931dfd28d4ed9aaaec0489448dd7f50f527387b2956"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-ace55b0afa0afda280e8a66052f820914f41d04406adf86c5f4f243f630143a6"></a>

## Direct properties — custom_network_config.no_network_policy / edb927c223eb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5747898ba96a4c2d8a94be8fd27eb104ec9172dabbb2f7b9ed612bc3c9b7290e"></a>

## Next pages — custom_network_config.no_network_policy / edb927c223eb / 4

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83d43f32972da09cc902d7a03cf95a909da62b891605eaffcdbbe059b8760628"></a>

## custom_network_config.sli_config — custom_network_config.sli_config / 4ff3abe6be55 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- custom_network_config.sli_config

<a id="canonical-ad21dacbb4a422228fbbacf96c6fcbac28ae73faf722c956a1d8751301b904ba"></a>

Type: `"single"`. Computed.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

<a id="canonical-563b2d6fece5ba9d76b37aef97b04807b5b7a963ec78b44dfa404aa68262c373"></a>

## Direct properties — custom_network_config.sli_config / 4ff3abe6be55 / 3

- [dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-e6f7ef578f3263cbf53be629a4eb386fb255e2ccd5ce736bd3b18576751b9dc1): complete subsection reference.

<a id="canonical-08b8a1aa7d5ed74d348fd480cd025852f09047d162fdd1024d64bde978395bb3"></a>

<a id="canonical-c991dbb1905272f5aec387a26d9bced7aa303cb724af986af5f7b5a35400ace8"></a>

## labels property — custom_network_config.sli_config / 4ff3abe6be55 / 4

Type: `["map", "string"]`. Computed.

Add Labels for this network, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-7d6e685d92142ef6dbee013e06baf4ffe4f833ef6290a473dd4ae4fa3ba34929"></a>

<a id="canonical-ef591c00e0054d359f825d21069d55e6fdf44a602aa473e5f689121abbc161e1"></a>

## nameserver property — custom_network_config.sli_config / 4ff3abe6be55 / 5

Type: `"string"`. Computed.

Optional DNS V4 server IP to be used for name resolution.

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

- [no_dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-d41a20fec55d434525f8387e1f0b974b5009b1fd6c7694325bc0a087756279ad): complete subsection reference.

- [no_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-f18365586fee5f40f0d075f886cb4334c60bdcc68fc2453d2001798d632fd8a5): complete subsection reference.

- [no_v6_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-c1b99d89e264225064475cd3ad28433dd8819f13c35d5af3db9d65b919737580): complete subsection reference.

- [static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-32c921db064e14dda63b3d17fbbb3d50320715910dd154f2f0d30854bd14a810): complete subsection reference.

- [static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-e27563b5f2d7d3556a21c4911f6ee19f9563140ba007b04884d93232f8dd8169): complete subsection reference.

<a id="canonical-1056bdc637939bcb5ec7a19b4b29cb610aacd9573653ae2af80007fcb994f1ce"></a>

<a id="canonical-651b170e1d402c3e2bb78e54cd3107abfd943afacac7ce4e85b000852aeb6a2b"></a>

## vip property — custom_network_config.sli_config / 4ff3abe6be55 / 6

Type: `"string"`. Computed.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

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

<a id="canonical-ab6e1861ee201bb0954be116608032840f7d8a9df4d5eac5e1187a4f39312d40"></a>

## Next pages — custom_network_config.sli_config / 4ff3abe6be55 / 7

- [custom_network_config.sli_config.dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-e6f7ef578f3263cbf53be629a4eb386fb255e2ccd5ce736bd3b18576751b9dc1)
- [custom_network_config.sli_config.no_dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-d41a20fec55d434525f8387e1f0b974b5009b1fd6c7694325bc0a087756279ad)
- [custom_network_config.sli_config.no_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-f18365586fee5f40f0d075f886cb4334c60bdcc68fc2453d2001798d632fd8a5)
- [custom_network_config.sli_config.no_v6_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-c1b99d89e264225064475cd3ad28433dd8819f13c35d5af3db9d65b919737580)
- [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-32c921db064e14dda63b3d17fbbb3d50320715910dd154f2f0d30854bd14a810)
- [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-e27563b5f2d7d3556a21c4911f6ee19f9563140ba007b04884d93232f8dd8169)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-e6f7ef578f3263cbf53be629a4eb386fb255e2ccd5ce736bd3b18576751b9dc1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96e446c084803e4427fe5f6e22431b585590cb775f5d90ce909f9c3f569147b6"></a>

## custom_network_config.sli_config.dc_cluster_group — custom_network_config.sli_config.dc_cluster_group / 195b5f7b7290 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- custom_network_config.sli_config.dc_cluster_group

<a id="canonical-d08b3fa380cb3c68b528778ce30e03e8e38ad59511db665b6e98925d55c7bb46"></a>

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

<a id="canonical-47fe34e2845de7b449b0e8974886b74f5b19eb0803117b371d33acf6f747f15d"></a>

## Direct properties — custom_network_config.sli_config.dc_cluster_group / 195b5f7b7290 / 3

<a id="canonical-73bc5d5ed3ba162f9d8f5e99f9edd1320eaf1aeef9a5bd8af73405d0b7b4155a"></a>

<a id="canonical-4617ba01a8cb28dd47b0f4417985cb05d3000c048488f8e3ca1e80d39f18198e"></a>

## name property — custom_network_config.sli_config.dc_cluster_group / 195b5f7b7290 / 4

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

<a id="canonical-cfd541c446cabb90a49b1af7532aafd32b907d12d260b3cf9e29396554e7fd14"></a>

<a id="canonical-1d2964f94cf6923315b9a8f098f0bf78e061bb58825e4ffb753757eab3448491"></a>

## namespace property — custom_network_config.sli_config.dc_cluster_group / 195b5f7b7290 / 5

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

<a id="canonical-a60f093f4c85fae23025969bffe8d0bad09c85818ef28d4b77e6babaed340975"></a>

<a id="canonical-7e506dbe43d0b152d6d189f1918c965ecbdd153dfc6acd47b6b4f3630bd7e9d0"></a>

## tenant property — custom_network_config.sli_config.dc_cluster_group / 195b5f7b7290 / 6

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

<a id="canonical-5b18ef3787230b7d76a96675c89fd45acb79753a497b8e4b6f10593804a325fd"></a>

## Next pages — custom_network_config.sli_config.dc_cluster_group / 195b5f7b7290 / 7

- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-d41a20fec55d434525f8387e1f0b974b5009b1fd6c7694325bc0a087756279ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a16ce5aef8541f3232f8b7b71ea7e4a89d410325aa893e899f167dd828f0b22"></a>

## custom_network_config.sli_config.no_dc_cluster_group — custom_network_config.sli_config.no_dc_cluster_group / 685cd23b9f62 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- custom_network_config.sli_config.no_dc_cluster_group

<a id="canonical-7ec87a5320874fa159378a56ab42222d0a879270b64e560d90effabebccd718b"></a>

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

<a id="canonical-924c2beacdbbe02f81fe89c2e557e14b2b7f204fa344230cb4bf27d6a011f345"></a>

## Direct properties — custom_network_config.sli_config.no_dc_cluster_group / 685cd23b9f62 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed5a9810ff7d540bbb8160a238a02fadbfeb9feb2bf6694c1bf1ea8154d69723"></a>

## Next pages — custom_network_config.sli_config.no_dc_cluster_group / 685cd23b9f62 / 4

- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-f18365586fee5f40f0d075f886cb4334c60bdcc68fc2453d2001798d632fd8a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1293aa677e4ebb3fd9a8afef529c8f294003388e02e0aabc9197c865e9b47157"></a>

## custom_network_config.sli_config.no_static_routes — custom_network_config.sli_config.no_static_routes / f11af6d66c8a / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- custom_network_config.sli_config.no_static_routes

<a id="canonical-ecd077a78286ff354c227bd5b1195f7baa7ba3280bdae5f32744608bdd99d922"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no static routes.

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

<a id="canonical-6b8e51facebb4f744c9792aec3d8fbce54e64e93cb513b8830795253b938c926"></a>

## Direct properties — custom_network_config.sli_config.no_static_routes / f11af6d66c8a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9c6d9488503bc4146e8a9147460593b8be25e7ac9bb532e22fa10c6d5e5b702"></a>

## Next pages — custom_network_config.sli_config.no_static_routes / f11af6d66c8a / 4

- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-c1b99d89e264225064475cd3ad28433dd8819f13c35d5af3db9d65b919737580"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4306576aa5caba1e3f157d49770ad2f7c174cf47c923560dfe898a6d9c108e97"></a>

## custom_network_config.sli_config.no_v6_static_routes — custom_network_config.sli_config.no_v6_static_routes / 9447e6b138cc / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- custom_network_config.sli_config.no_v6_static_routes

<a id="canonical-6081515011db362567dd29b8452de982587e54c7dfa225e4f89d6fe9df2cc61a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no v6 static routes.

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

<a id="canonical-9c7473a147900f921a3e519e5e9781664c1ca0e022ee859f4e66a8a923dc96fb"></a>

## Direct properties — custom_network_config.sli_config.no_v6_static_routes / 9447e6b138cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-df0402ec81c44744d4be1893687a7e0b7d30c7286a935e873e2510594373e74d"></a>

## Next pages — custom_network_config.sli_config.no_v6_static_routes / 9447e6b138cc / 4

- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-32c921db064e14dda63b3d17fbbb3d50320715910dd154f2f0d30854bd14a810"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d057ccb18c5dfed2533924991afcc656655db824beafe3680264646bd355d4c"></a>

## custom_network_config.sli_config.static_routes — custom_network_config.sli_config.static_routes / e65baccafdbf / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- custom_network_config.sli_config.static_routes

<a id="canonical-adb603cec0bc90888ea38115d2ba3ef6e674a3e3891a00ee33deb0fd70ba9f23"></a>

Type: `"single"`. Computed.

Configuration parameter for static routes.

Upstream description:

List of static routes.

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

<a id="canonical-4bba892d507c4c98e5f70c6f2768adccd66f71868e6ad8739865417dc77ad978"></a>

## Direct properties — custom_network_config.sli_config.static_routes / e65baccafdbf / 3

- [static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-a3635915609adef0a0a347c6b635d2786874b04e2f598e992ed259e0f6493099): complete subsection reference.

<a id="canonical-1769caea6e0a0c41cf56f840f03e30f122b5deb177624b6988c2d10cffa4d1f9"></a>

## Next pages — custom_network_config.sli_config.static_routes / e65baccafdbf / 4

- [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-a3635915609adef0a0a347c6b635d2786874b04e2f598e992ed259e0f6493099)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-a3635915609adef0a0a347c6b635d2786874b04e2f598e992ed259e0f6493099"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efac0e6a5cbb8514e42af6d6305c809e53b2c267cf747be4754349bf37949663"></a>

## custom_network_config.sli_config.static_routes.static_routes — custom_network_config.sli_config.static_routes.static_routes / fbfb1742a720 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-32c921db064e14dda63b3d17fbbb3d50320715910dd154f2f0d30854bd14a810)
- custom_network_config.sli_config.static_routes.static_routes

<a id="canonical-e94cf72a48b636be51b388208261d48d1555cf3cd6221f2db7dc2b2cfc18dc83"></a>

Type: `"list"`. Computed.

Static Routes. List of static routes.

Upstream description:

List of static routes.

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

<a id="canonical-6d51e98535781ab577ada2a2b689c76ae21fca2d77c96de92e9fac57b5e32dd0"></a>

## Direct properties — custom_network_config.sli_config.static_routes.static_routes / fbfb1742a720 / 3

<a id="canonical-137659a3f6c579862c49de826bb6ffaf05d31ea9297c048865c33345e24f605e"></a>

<a id="canonical-fb2ccaeef0b06b43070ba6bac38fcc532feedcf157e4265601da696dcf30b490"></a>

## attrs property — custom_network_config.sli_config.static_routes.static_routes / fbfb1742a720 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-51f0ca346435dd304b367ea056dac1a776ed0f7569987c4bffb310598c6bf28a): complete subsection reference.

<a id="canonical-16b10c5bba3868fb09944cb521b8c1e31bd28d81d3d6642a8ad95cc1588b999c"></a>

<a id="canonical-f3290843ce62cbfed808b108115d6aeb6dba83135e2aadc0e1d5acd9c73dd217"></a>

## ip_address property — custom_network_config.sli_config.static_routes.static_routes / fbfb1742a720 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-f717bb66faf3feed8f987641003b6cc9dfa4ac6b40ce4a2d3fd1f46c4fd91b06"></a>

<a id="canonical-36864bd16cb5e3d931747951abab0782c315e901bd7ca6df8f80bf9a9694bba3"></a>

## ip_prefixes property — custom_network_config.sli_config.static_routes.static_routes / fbfb1742a720 / 6

Type: `["list", "string"]`. Computed.

List of route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-dfa09640e2bd7bfc4f7b08110d4191ca350673205b1dbbaa9ce7ad3b849736ea): complete subsection reference.

<a id="canonical-ae3b0934f48de15736de434e4da63cab6f0197cc8dc47d0dfef2e1e96b161796"></a>

## Next pages — custom_network_config.sli_config.static_routes.static_routes / fbfb1742a720 / 7

- [custom_network_config.sli_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-51f0ca346435dd304b367ea056dac1a776ed0f7569987c4bffb310598c6bf28a)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-dfa09640e2bd7bfc4f7b08110d4191ca350673205b1dbbaa9ce7ad3b849736ea)
- [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-32c921db064e14dda63b3d17fbbb3d50320715910dd154f2f0d30854bd14a810)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-51f0ca346435dd304b367ea056dac1a776ed0f7569987c4bffb310598c6bf28a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61f84279b1b93fd27e638e06180219ef1d5498900584fe76f92f32248033a0bb"></a>

## custom_network_config.sli_config.static_routes.static_routes.default_gateway — custom_network_config.sli_config.static_routes.static_routes.default_gateway / 78a07d1358e0 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-32c921db064e14dda63b3d17fbbb3d50320715910dd154f2f0d30854bd14a810)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-a3635915609adef0a0a347c6b635d2786874b04e2f598e992ed259e0f6493099)
- custom_network_config.sli_config.static_routes.static_routes.default_gateway

<a id="canonical-6f1fafee96d79afdcd1a8d5b16d1ce96f96b2e922e1d8fac40369bd832226893"></a>

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

<a id="canonical-3bbeb5aab7844076934049032a7db9e95befdd6b3afae4c236b6c813236bd692"></a>

## Direct properties — custom_network_config.sli_config.static_routes.static_routes.default_gateway / 78a07d1358e0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-47a1745618f1f6d211d6577d59d62b9584b8ffa2578520dddecb8fbf203b7590"></a>

## Next pages — custom_network_config.sli_config.static_routes.static_routes.default_gateway / 78a07d1358e0 / 4

- [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-a3635915609adef0a0a347c6b635d2786874b04e2f598e992ed259e0f6493099)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-dfa09640e2bd7bfc4f7b08110d4191ca350673205b1dbbaa9ce7ad3b849736ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-319883f7d5a73d46d2dcb2cf793d7e11e097986fe438b7f644d43d8892913b5c"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface — custom_network_config.sli_config.static_routes.static_routes.node_interface / d672b02e833a / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-32c921db064e14dda63b3d17fbbb3d50320715910dd154f2f0d30854bd14a810)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-a3635915609adef0a0a347c6b635d2786874b04e2f598e992ed259e0f6493099)
- custom_network_config.sli_config.static_routes.static_routes.node_interface

<a id="canonical-4d10f24d2dca60465a0f9bea7fbba60047c6f0ea006da69d6b30168b45846ff1"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

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

<a id="canonical-96d110425dc2894f30d3332ec01d7f2ed866436d28b6a9e5e63f9bfb67660bb9"></a>

## Direct properties — custom_network_config.sli_config.static_routes.static_routes.node_interface / d672b02e833a / 3

- [list](data-sources--securemesh_site--reference--group-003.md#canonical-3754d8f0426e275e4c285f28e94b5986baeb8af3d33dc161bd73b7449ee87715): complete subsection reference.

<a id="canonical-58b7c06537adcbbd2155c243508f64634d864f177b32b4303e42c2809ae4ec26"></a>

## Next pages — custom_network_config.sli_config.static_routes.static_routes.node_interface / d672b02e833a / 4

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-3754d8f0426e275e4c285f28e94b5986baeb8af3d33dc161bd73b7449ee87715)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-a3635915609adef0a0a347c6b635d2786874b04e2f598e992ed259e0f6493099)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-3754d8f0426e275e4c285f28e94b5986baeb8af3d33dc161bd73b7449ee87715"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-878c97c6fae0034f1241d76cb37796bacd959b49028b0b725b4a91796f4c029e"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface.list — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / c3cd672c74c9 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-32c921db064e14dda63b3d17fbbb3d50320715910dd154f2f0d30854bd14a810)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-a3635915609adef0a0a347c6b635d2786874b04e2f598e992ed259e0f6493099)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-dfa09640e2bd7bfc4f7b08110d4191ca350673205b1dbbaa9ce7ad3b849736ea)
- custom_network_config.sli_config.static_routes.static_routes.node_interface.list

<a id="canonical-bc1dc6bf9bfc9a71ab0bf847ebc5532dbe9b98b80c27fc006c4b44b061bcbe54"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-ae0a2870b4da7f01f62e12f46df3fb1b35db60d886ec74d4709aa32bb2862ed5"></a>

## Direct properties — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / c3cd672c74c9 / 3

- [interface](data-sources--securemesh_site--reference--group-003.md#canonical-edd24de6d62c10a7480fc632510da26d6f36f928440559b07a703d8c0ca8e4c5): complete subsection reference.

<a id="canonical-b18f282dc325325982f045a7f2fe8809334b4c7f1083176da467a89ab33b7670"></a>

<a id="canonical-8dc423dafb76456ac4d92fc0032dd8ca9b7c0c4fb8a790d2d74f5ee21d313f04"></a>

## node property — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / c3cd672c74c9 / 4

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-eee59bdd69b59fcf26bf027739cfecdee5ee44085b43a92d325a875a325c6818"></a>

## Next pages — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / c3cd672c74c9 / 5

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-003.md#canonical-edd24de6d62c10a7480fc632510da26d6f36f928440559b07a703d8c0ca8e4c5)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-dfa09640e2bd7bfc4f7b08110d4191ca350673205b1dbbaa9ce7ad3b849736ea)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-edd24de6d62c10a7480fc632510da26d6f36f928440559b07a703d8c0ca8e4c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56d876593a9d7151a15a6b555eef516861b3bccd887c6eff27163fbcd2c7f95d"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / c3081eda5215 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-32c921db064e14dda63b3d17fbbb3d50320715910dd154f2f0d30854bd14a810)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-a3635915609adef0a0a347c6b635d2786874b04e2f598e992ed259e0f6493099)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-dfa09640e2bd7bfc4f7b08110d4191ca350673205b1dbbaa9ce7ad3b849736ea)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-3754d8f0426e275e4c285f28e94b5986baeb8af3d33dc161bd73b7449ee87715)
- custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-bd3d4ca8f03bfa327d340db87d5ad2ecace4cb2035845a6ddf07ee66ba4a57b2"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-7144b4ca9818276a5f08141e42b824fa567c386cae44feb6a0441db1368f8cd2"></a>

## Direct properties — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / c3081eda5215 / 3

<a id="canonical-10de66f004658d6b2a1fa712d5deb861cfc940921f5e7ac94451b79f674895e8"></a>

<a id="canonical-4f46d4cb894c54462c3cf6857664d0deb286c02b5e9b4bf7cc8ba30b6b589591"></a>

## kind property — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / c3081eda5215 / 4

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

<a id="canonical-c50ae05cc75086f9aeca08672ce72ce206315033839493422f325dfb09de1e57"></a>

<a id="canonical-a003f671fb6833411f58e6d82dea9033cd86672e4f455ca0e7b8f2e259a9f34f"></a>

## name property — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / c3081eda5215 / 5

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

<a id="canonical-c1f88b730341cedea27532e751d3fd8bfb10920597cf58ff8829a782543d81ca"></a>

<a id="canonical-5b342f0357a1769098c58f7232c834a6519eb4ee369e6c6f7b7991a75c14115d"></a>

## namespace property — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / c3081eda5215 / 6

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

<a id="canonical-66e151c192cfcc4308bb9a1f37ece792213d18998695fbbd872ac863dbfa45d8"></a>

<a id="canonical-dbec47dd5b39ed7031a168d127e989e856dc1a2dff2bcc72a146cc71427d6f0b"></a>

## tenant property — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / c3081eda5215 / 7

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

<a id="canonical-cf54920870cabc4214cc9241acea1cd596be4c0d09ca5a3eb2f54539c5c8d279"></a>

<a id="canonical-210aaf7898264cbbbf2e835256e2cb4f1404bbe38b0fae62a4c816a451ab1e36"></a>

## uid property — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / c3081eda5215 / 8

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

<a id="canonical-b10fd31753edeb35f53cdd938ee65e7706d29f8980478bfb89a36ebef56c9960"></a>

## Next pages — custom_network_config.sli_config.static_routes.static_routes.node_interface.list / c3081eda5215 / 9

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-3754d8f0426e275e4c285f28e94b5986baeb8af3d33dc161bd73b7449ee87715)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-e27563b5f2d7d3556a21c4911f6ee19f9563140ba007b04884d93232f8dd8169"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5143776f9daa72ed7749a1950389ed27b5dfc5dcfe96ae1038bc585e8a7edf1"></a>

## custom_network_config.sli_config.static_v6_routes — custom_network_config.sli_config.static_v6_routes / f49b473ac975 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- custom_network_config.sli_config.static_v6_routes

<a id="canonical-45d825941c73e246d7a438069d8cb46ea47af85f4e378089f0df3b7faea7d435"></a>

Type: `"single"`. Computed.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-081078251a583dc9bab55c4f6a16c7065cc04d350a6e909fc2030f4bd9ed7a3b"></a>

## Direct properties — custom_network_config.sli_config.static_v6_routes / f49b473ac975 / 3

- [static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-24d19674240a022ef43900dfa34cc90a9c272ef73cb6e7c504c79fcac058f980): complete subsection reference.

<a id="canonical-388515220743e7ffeac933183fa46374fd4713e44d277e51d8b9407963487ab4"></a>

## Next pages — custom_network_config.sli_config.static_v6_routes / f49b473ac975 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-24d19674240a022ef43900dfa34cc90a9c272ef73cb6e7c504c79fcac058f980)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-24d19674240a022ef43900dfa34cc90a9c272ef73cb6e7c504c79fcac058f980"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2c37e448896d1055247bec9de2425a4dad2f9217810c6a7448ccbfe921e72cb"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes — custom_network_config.sli_config.static_v6_routes.static_routes / 993d6fd05886 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-e27563b5f2d7d3556a21c4911f6ee19f9563140ba007b04884d93232f8dd8169)
- custom_network_config.sli_config.static_v6_routes.static_routes

<a id="canonical-5ffcc4765f8903478318a70ca8888fa5a7548145edec140512fe16a912398771"></a>

Type: `"list"`. Computed.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-b605fcc90e6af92b8fff53148ced886b1bd39f4b5463d7f705c00dc7818733a0"></a>

## Direct properties — custom_network_config.sli_config.static_v6_routes.static_routes / 993d6fd05886 / 3

<a id="canonical-a44136706f17dc64287bcbbe12d82be740321c1860616b4845d120939bbc693e"></a>

<a id="canonical-ccc9d980f0202e46e17ddf4315a37f6e4d5fedaacca050d569ff8db8b28ea037"></a>

## attrs property — custom_network_config.sli_config.static_v6_routes.static_routes / 993d6fd05886 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-2950e74fbc6be1629e82b967d69d75a203445c73dec6ab7039e5eb46fd67fbea): complete subsection reference.

<a id="canonical-91e04b48e19227ea98983d404b497c38ae103dc4932559c7ef4e87a2bd0d9ee7"></a>

<a id="canonical-2d0e87e36e2ff498a400c327964f82bea6094d2d6422edf3f01e9d4d2bf6075e"></a>

## ip_address property — custom_network_config.sli_config.static_v6_routes.static_routes / 993d6fd05886 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-e711b95653f841089e1f6413b1e9c9eac94661b906724511d92166027f0421a2"></a>

<a id="canonical-4b634ad6c35a65e88e0aa0c9c858bbbb3f5d733ee562e2a905c1799615a43692"></a>

## ip_prefixes property — custom_network_config.sli_config.static_v6_routes.static_routes / 993d6fd05886 / 6

Type: `["list", "string"]`. Computed.

List of IPv6 route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-cc56e78b5d0afbee189186bd64572dccc5cd7bc2e0ee0f1e143d7e9573c1c106): complete subsection reference.

<a id="canonical-3bed774455b90526c50db2880e3461d1fff77bb77cdea665389dc452c7f8c2a3"></a>

## Next pages — custom_network_config.sli_config.static_v6_routes.static_routes / 993d6fd05886 / 7

- [custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-2950e74fbc6be1629e82b967d69d75a203445c73dec6ab7039e5eb46fd67fbea)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-cc56e78b5d0afbee189186bd64572dccc5cd7bc2e0ee0f1e143d7e9573c1c106)
- [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-e27563b5f2d7d3556a21c4911f6ee19f9563140ba007b04884d93232f8dd8169)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-2950e74fbc6be1629e82b967d69d75a203445c73dec6ab7039e5eb46fd67fbea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a71e15b8700ce4959b6362067bca98843cb4ca355be67eaf674bef583d421562"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway — custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway / 8a0b1411740d / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-e27563b5f2d7d3556a21c4911f6ee19f9563140ba007b04884d93232f8dd8169)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-24d19674240a022ef43900dfa34cc90a9c272ef73cb6e7c504c79fcac058f980)
- custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-9ff71083b1e631079b742fb238062edf0c62323280f2eb931a4e01b32d89bcd8"></a>

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

<a id="canonical-19e7186803caaf0325dad7111596dd0caa52cf1306a36eae2b5db2896dac7962"></a>

## Direct properties — custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway / 8a0b1411740d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d5b95275362cc2bbb7d60e864ae34e028ab1305ad2071a16444b92142dcca70c"></a>

## Next pages — custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway / 8a0b1411740d / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-24d19674240a022ef43900dfa34cc90a9c272ef73cb6e7c504c79fcac058f980)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-cc56e78b5d0afbee189186bd64572dccc5cd7bc2e0ee0f1e143d7e9573c1c106"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b0146c1ee458293555818b8599675e20f35331464ba9664c5555a221a79efad"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface / 14d92f917158 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-e27563b5f2d7d3556a21c4911f6ee19f9563140ba007b04884d93232f8dd8169)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-24d19674240a022ef43900dfa34cc90a9c272ef73cb6e7c504c79fcac058f980)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface

<a id="canonical-930b7b967774d9e7a1cb08b85607322f7426fb1794faeeed0791ffeeb6918808"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

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

<a id="canonical-48cea9233e95f0eb40316534cf08492506b2bd811a650e6f77a7bfb1179e0dba"></a>

## Direct properties — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface / 14d92f917158 / 3

- [list](data-sources--securemesh_site--reference--group-003.md#canonical-9bd67353f07b745dac492295f62a6ba0fa2283770ca41a3b5b1c0d30e23ff7eb): complete subsection reference.

<a id="canonical-e188470ca7f73d96d3ea92d71aae22804eb5bf35c6b3805a34c4f04a0a550dae"></a>

## Next pages — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface / 14d92f917158 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-9bd67353f07b745dac492295f62a6ba0fa2283770ca41a3b5b1c0d30e23ff7eb)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-24d19674240a022ef43900dfa34cc90a9c272ef73cb6e7c504c79fcac058f980)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-9bd67353f07b745dac492295f62a6ba0fa2283770ca41a3b5b1c0d30e23ff7eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae5a1a05e4b0fbddb5ba868f9551b0b47dd6044fe08939174fd9186919b179cb"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / e48464e994f9 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-e27563b5f2d7d3556a21c4911f6ee19f9563140ba007b04884d93232f8dd8169)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-24d19674240a022ef43900dfa34cc90a9c272ef73cb6e7c504c79fcac058f980)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-cc56e78b5d0afbee189186bd64572dccc5cd7bc2e0ee0f1e143d7e9573c1c106)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-ee96d256e035486e4daaef3eed84b4149f5070bfe991248de11237294b076696"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-7820266da0aaeecd9e53887bcc7df60d4f639817c49575868f4b13ba69c062cb"></a>

## Direct properties — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / e48464e994f9 / 3

- [interface](data-sources--securemesh_site--reference--group-003.md#canonical-5af0c5ba1a5aaf153d41ed484b959cb6aa3c46f4e99705be9253e82549652d68): complete subsection reference.

<a id="canonical-e889913afad0cea90ac71bb496a22b7b9b921401617ef7b3084b5e8c2aada853"></a>

<a id="canonical-3f52fed7ff91848fcf27f57f1b4ed749b78e89ab95deafd375b28411f2a238ad"></a>

## node property — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / e48464e994f9 / 4

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-6c2206e6ea30623423edc638424ed85a806545afc4152db81eb5ce4408021a76"></a>

## Next pages — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / e48464e994f9 / 5

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-003.md#canonical-5af0c5ba1a5aaf153d41ed484b959cb6aa3c46f4e99705be9253e82549652d68)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-cc56e78b5d0afbee189186bd64572dccc5cd7bc2e0ee0f1e143d7e9573c1c106)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-5af0c5ba1a5aaf153d41ed484b959cb6aa3c46f4e99705be9253e82549652d68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fbc6d8c5ab98a8e56c4d2ae85c9d1698728f8fc9c81967379df83fd58f5e0a5"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 27e4c0a802f4 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-e27563b5f2d7d3556a21c4911f6ee19f9563140ba007b04884d93232f8dd8169)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-24d19674240a022ef43900dfa34cc90a9c272ef73cb6e7c504c79fcac058f980)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-cc56e78b5d0afbee189186bd64572dccc5cd7bc2e0ee0f1e143d7e9573c1c106)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-9bd67353f07b745dac492295f62a6ba0fa2283770ca41a3b5b1c0d30e23ff7eb)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-63505ded2366e6514427473c7b70e7b49489d86dba1fa7c62f8c8f514bb415d2"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-95b4e9a5d5db6ea28c67dab85d21f970d45789f216305384a6f764b2d622be45"></a>

## Direct properties — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 27e4c0a802f4 / 3

<a id="canonical-e8f2c340d574b8347992f184d2ff75665ec782f74fca740afe859dab8fdd00fd"></a>

<a id="canonical-46312914eea4e51e8f77df9b70e76aabe4b4958dbf7f9de35745bbaefc67aea3"></a>

## kind property — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 27e4c0a802f4 / 4

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

<a id="canonical-58ef9dc6754c094e33ea17c027317eb0632818a3ac865ee90a3674129ce98027"></a>

<a id="canonical-39ce5e821259407b9b591d59f6adc5990bfbe14ed2ad0909b73ad38f53876563"></a>

## name property — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 27e4c0a802f4 / 5

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

<a id="canonical-35cc441b1b0ffdd6249bc9a29305d78d9c1fa2eb6eeea6e31fdcd4c92a0b2761"></a>

<a id="canonical-a698b226705bee861f693ca997bbfd72b9ac453ac23674751ff6a6e15ac9896d"></a>

## namespace property — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 27e4c0a802f4 / 6

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

<a id="canonical-1c772cc372e25465a7be1379224126cf2712fb9c9c243d57c9ca284c893e2ffa"></a>

<a id="canonical-e71d7f4c9eb34c46ac9761be93cc4d3fe478b8d6e3288ac048e3b6e33a3f1160"></a>

## tenant property — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 27e4c0a802f4 / 7

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

<a id="canonical-29ca2689dcb48c4b23a98309d21ddf971f0a3baec84caa7cf2f735575bf21fb6"></a>

<a id="canonical-2641defa30add5e972510fcc08e86757a2f85258367d81afea47a426f8d4e8a5"></a>

## uid property — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 27e4c0a802f4 / 8

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

<a id="canonical-490b16bbef4a58df1247cfd6667ab55248db24555e91e20d842c2186a6feb0f5"></a>

## Next pages — custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.l / 27e4c0a802f4 / 9

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-9bd67353f07b745dac492295f62a6ba0fa2283770ca41a3b5b1c0d30e23ff7eb)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2bc1525c9266b0d72e157adb8404e145baed0d3e4d1061c4e3bd34a79a970f7d"></a>

## custom_network_config.slo_config — custom_network_config.slo_config / 3d8116911464 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- custom_network_config.slo_config

<a id="canonical-e3b059f3609c33ca360a0878715c906262327bdd93dcc755e0d2d944d8bde980"></a>

Type: `"single"`. Computed.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

<a id="canonical-277d7263c643113b1469865d4b2fb86490aea6425ad877d84b1b870876277f90"></a>

## Direct properties — custom_network_config.slo_config / 3d8116911464 / 3

- [dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-091ed995bb144ccc2832ad356fc0247484c59593fd328ef745e079873a238c1c): complete subsection reference.

<a id="canonical-a4ca39a26a53a101141f90aade25f331256c15d97956234ebccb138a0975049f"></a>

<a id="canonical-cf04b4f563a3efa591bc43e2d677af9360af1da40627eda66f1b16da256fc12a"></a>

## labels property — custom_network_config.slo_config / 3d8116911464 / 4

Type: `["map", "string"]`. Computed.

Add Labels for this network, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-a90e9f45f87aac081027077e55c03013f516c54f313307a96f4e58f3f797ccb9"></a>

<a id="canonical-e83f51db3801cb5aeb94389ce579c0b9dfc1f4a7cc540905beb78ca3944abd84"></a>

## nameserver property — custom_network_config.slo_config / 3d8116911464 / 5

Type: `"string"`. Computed.

Optional DNS V4 server IP to be used for name resolution.

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

- [no_dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-3da3735b5cb5513deae65afdc8a4c6eb9a4dcdaa0124ceb799d3a763bdaaa72a): complete subsection reference.

- [no_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-756e1ce6272ce153fb3b81adbe21b5febb5f81fcfa29769f7729396205adcbba): complete subsection reference.

- [no_v6_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3ebc15d15170d5851c62d14f0500b6282e65d654c4e64b937d10ea203b90c958): complete subsection reference.

- [static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-ebd2256dc9a55c693f5d456be7cb12d1a12a3ecbf41fb0628e8266095b5e1772): complete subsection reference.

- [static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-4264d8bbba705a76f895a4e46cc501f9d26b3c652823436945770b53d6ec8cfd): complete subsection reference.

<a id="canonical-e4276c21b447f6d3bbbeaa4cf194ad55017e22fdb81cd0fc4db40edc29c559be"></a>

<a id="canonical-319c7ea85b285ec93ca73d69eae057e6b5f7684162f4df8c4c402696b6e8f1bf"></a>

## vip property — custom_network_config.slo_config / 3d8116911464 / 6

Type: `"string"`. Computed.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

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

<a id="canonical-247c5a0492944a78aeff6092164032e8bb4814a51dbc7104d18b4afb1f579975"></a>

## Next pages — custom_network_config.slo_config / 3d8116911464 / 7

- [custom_network_config.slo_config.dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-091ed995bb144ccc2832ad356fc0247484c59593fd328ef745e079873a238c1c)
- [custom_network_config.slo_config.no_dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-3da3735b5cb5513deae65afdc8a4c6eb9a4dcdaa0124ceb799d3a763bdaaa72a)
- [custom_network_config.slo_config.no_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-756e1ce6272ce153fb3b81adbe21b5febb5f81fcfa29769f7729396205adcbba)
- [custom_network_config.slo_config.no_v6_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3ebc15d15170d5851c62d14f0500b6282e65d654c4e64b937d10ea203b90c958)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-ebd2256dc9a55c693f5d456be7cb12d1a12a3ecbf41fb0628e8266095b5e1772)
- [custom_network_config.slo_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-4264d8bbba705a76f895a4e46cc501f9d26b3c652823436945770b53d6ec8cfd)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-091ed995bb144ccc2832ad356fc0247484c59593fd328ef745e079873a238c1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c2f80e7fdd51b83a739275643e02aaf74ba5f4732fa5610b47fa3e51ffb8387"></a>

## custom_network_config.slo_config.dc_cluster_group — custom_network_config.slo_config.dc_cluster_group / 291371257b7f / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- custom_network_config.slo_config.dc_cluster_group

<a id="canonical-967ffd7a08952211431144b72deee745fa776c7dc7f4982e10b6c1d855a65dc1"></a>

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

<a id="canonical-b0bad68158383e9167f405f2c4d2e8699a52998ad9a7f0ab8b78e2311a2e1a4f"></a>

## Direct properties — custom_network_config.slo_config.dc_cluster_group / 291371257b7f / 3

<a id="canonical-8a5a1baadd41579c709460052eb962a10891c7592340693bbd193f83c2096533"></a>

<a id="canonical-fd98b0e2c58b696f3fb1633610cbdfd151ef05767b3dcddedced307dabb821fd"></a>

## name property — custom_network_config.slo_config.dc_cluster_group / 291371257b7f / 4

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

<a id="canonical-8ecada2131f5da55bf129dc362f57793bf637946c8dc3d2ca4bca5fc3936e956"></a>

<a id="canonical-05dbbd84acce0756055486034feb89c3829ed649271f9f47394067b3735b41bd"></a>

## namespace property — custom_network_config.slo_config.dc_cluster_group / 291371257b7f / 5

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

<a id="canonical-8116cc256a8031da33403eb44680059285397c78b080d10d19175c29945623b1"></a>

<a id="canonical-2f1a0cc0547134196fc6d4c2dd4623b4648bb16d3b7d3007dc1d36733d64f5fd"></a>

## tenant property — custom_network_config.slo_config.dc_cluster_group / 291371257b7f / 6

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

<a id="canonical-9ed5d355f02e6536e5d7a0ba5eac84d2aafabc1b7c188ba0f3f6ee6dc1516fe6"></a>

## Next pages — custom_network_config.slo_config.dc_cluster_group / 291371257b7f / 7

- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-3da3735b5cb5513deae65afdc8a4c6eb9a4dcdaa0124ceb799d3a763bdaaa72a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a072ff6636ef66bdd9d533750fa992a9f72f144ebe760bbf2fd65690ccd287a"></a>

## custom_network_config.slo_config.no_dc_cluster_group — custom_network_config.slo_config.no_dc_cluster_group / 82f995e2f691 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- custom_network_config.slo_config.no_dc_cluster_group

<a id="canonical-1eda53017d7fb33956a16e356459fe493c8515247e7304e6f8be0a7821b3429e"></a>

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

<a id="canonical-abe90dfa3c9004c9959483e8b6ace3bcc8692169dd4ef08fa39ebc9032fd66a6"></a>

## Direct properties — custom_network_config.slo_config.no_dc_cluster_group / 82f995e2f691 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-765071fde06d947c2f07048c4897105e74c6f8d5a4693d8e74426d2670509524"></a>

## Next pages — custom_network_config.slo_config.no_dc_cluster_group / 82f995e2f691 / 4

- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-756e1ce6272ce153fb3b81adbe21b5febb5f81fcfa29769f7729396205adcbba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4d78389e6c27bac1ac418b5c7695b6a92ac7d1a18251947c54eb70ea8644081"></a>

## custom_network_config.slo_config.no_static_routes — custom_network_config.slo_config.no_static_routes / 8067d5a9cc34 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- custom_network_config.slo_config.no_static_routes

<a id="canonical-fe5dcf9ccefb1c6052017eb4ca8b097c55a51f782e55e52efc2213ef02e6ecf2"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no static routes.

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

<a id="canonical-d37cac8ca3af7337602e421d51db341cd8ed84bf3f3eba2b39ecfd2c37d34bd9"></a>

## Direct properties — custom_network_config.slo_config.no_static_routes / 8067d5a9cc34 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa74fcfeadde8ac9adfef583ad0c79d03c8844422cb9ef86218e55a544294c6b"></a>

## Next pages — custom_network_config.slo_config.no_static_routes / 8067d5a9cc34 / 4

- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-3ebc15d15170d5851c62d14f0500b6282e65d654c4e64b937d10ea203b90c958"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f32d201147a2e61337b082c815a9262a80728d0e58d96d1b0ba249874ef789c"></a>

## custom_network_config.slo_config.no_v6_static_routes — custom_network_config.slo_config.no_v6_static_routes / fddfef15769c / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- custom_network_config.slo_config.no_v6_static_routes

<a id="canonical-690266c99cf90a29b3be1ceda330b9d5650b8b480a5563db3703130813896583"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no v6 static routes.

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

<a id="canonical-8cd6dcea04a7e366715739d4ab5d5bfbe04f6ad31d95462329af86c0ecaf31e9"></a>

## Direct properties — custom_network_config.slo_config.no_v6_static_routes / fddfef15769c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-42590c620915e6a21ae8e7370012b183eebcc77865da3ef93e46b3e9918c1a21"></a>

## Next pages — custom_network_config.slo_config.no_v6_static_routes / fddfef15769c / 4

- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-ebd2256dc9a55c693f5d456be7cb12d1a12a3ecbf41fb0628e8266095b5e1772"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c14c2e6be2e4c603a5f0223fc4f8adc4ff48d7a47377ab4ee53fc4f45dba4da"></a>

## custom_network_config.slo_config.static_routes — custom_network_config.slo_config.static_routes / a2afcb736bb0 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- custom_network_config.slo_config.static_routes

<a id="canonical-a4165b957f93bcbbd02833473aee8ddc6d1628848af38ad91d9c0330c7a034c8"></a>

Type: `"single"`. Computed.

Configuration parameter for static routes.

Upstream description:

List of static routes.

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

<a id="canonical-50804989c6d7359613f2af2cc2273ff4a34a24bedac9ec5763a245ff6ae696f1"></a>

## Direct properties — custom_network_config.slo_config.static_routes / a2afcb736bb0 / 3

- [static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-021e207896f45efe583d32ed49cc1ed6cf809f6db9c9b52d0c65130d87033128): complete subsection reference.

<a id="canonical-b1dea0793848d8c9584c8edbe7e879f81cfc921860fbda53df2a89f48451d522"></a>

## Next pages — custom_network_config.slo_config.static_routes / a2afcb736bb0 / 4

- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-021e207896f45efe583d32ed49cc1ed6cf809f6db9c9b52d0c65130d87033128)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-021e207896f45efe583d32ed49cc1ed6cf809f6db9c9b52d0c65130d87033128"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-acf35dbfe528e6e71b8993e01e65f5df395d51fa025313276c19459bc6f0a323"></a>

## custom_network_config.slo_config.static_routes.static_routes — custom_network_config.slo_config.static_routes.static_routes / f74c0b750b94 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-ebd2256dc9a55c693f5d456be7cb12d1a12a3ecbf41fb0628e8266095b5e1772)
- custom_network_config.slo_config.static_routes.static_routes

<a id="canonical-f9f9299c7e689d88d81aefa636d5c8620049a1023d4121e1d9246caebf9c365f"></a>

Type: `"list"`. Computed.

Static Routes. List of static routes.

Upstream description:

List of static routes.

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

<a id="canonical-3106a472e2846d367a47b6de9e4e6d12879fbd99e7f5dc95402267cb905d5302"></a>

## Direct properties — custom_network_config.slo_config.static_routes.static_routes / f74c0b750b94 / 3

<a id="canonical-5e7c239fa5fd00390db6e7e713d6ee9eb30dd11c9613d6f73e76a709aded9c05"></a>

<a id="canonical-6aea9c7330d04e52631bfea9484ced48e65a89fc5587298f15bc7cb2200f2cbf"></a>

## attrs property — custom_network_config.slo_config.static_routes.static_routes / f74c0b750b94 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-f765490147d6b89ed3ffe4a00857c1ce6a608af08c3c6b1ad1ed7921bba632d0): complete subsection reference.

<a id="canonical-dfc810eca90e875d5135a166f52b82d77a4e31dc1621fc69e5f7e2e077810107"></a>

<a id="canonical-3c5d0d2d30007d0ffcabdb6779800787ba936a97595f7bcefce15fbff654c6fd"></a>

## ip_address property — custom_network_config.slo_config.static_routes.static_routes / f74c0b750b94 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-01582deba78e34b071fefff2db483562c172c1df5c0f314eb7a6ced6acdc7c72"></a>

<a id="canonical-2ca319020c5b01518b44105d4ee45f7eb328cb9ada902342f471ef957dfe4e4b"></a>

## ip_prefixes property — custom_network_config.slo_config.static_routes.static_routes / f74c0b750b94 / 6

Type: `["list", "string"]`. Computed.

List of route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-28222d58951046048ba30d8ed3150d79bb89cb3c15dde18992866096a4b3843b): complete subsection reference.

<a id="canonical-340392c05f5ab954cef08f02f870dc4a1857a859a2d681dd24a0e4f1a3ceb5e6"></a>

## Next pages — custom_network_config.slo_config.static_routes.static_routes / f74c0b750b94 / 7

- [custom_network_config.slo_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-f765490147d6b89ed3ffe4a00857c1ce6a608af08c3c6b1ad1ed7921bba632d0)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-28222d58951046048ba30d8ed3150d79bb89cb3c15dde18992866096a4b3843b)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-ebd2256dc9a55c693f5d456be7cb12d1a12a3ecbf41fb0628e8266095b5e1772)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-f765490147d6b89ed3ffe4a00857c1ce6a608af08c3c6b1ad1ed7921bba632d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2772a3fa935c77d11d4594fd04f2d09a3b312accd29002d766fa5d6dc841ebb2"></a>

## custom_network_config.slo_config.static_routes.static_routes.default_gateway — custom_network_config.slo_config.static_routes.static_routes.default_gateway / 93f007ce39df / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-ebd2256dc9a55c693f5d456be7cb12d1a12a3ecbf41fb0628e8266095b5e1772)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-021e207896f45efe583d32ed49cc1ed6cf809f6db9c9b52d0c65130d87033128)
- custom_network_config.slo_config.static_routes.static_routes.default_gateway

<a id="canonical-f946d448c4109e8859e91b0930de75087abfd7644baf5547474d6ca44729d5d5"></a>

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

<a id="canonical-2d95fbbe0cdc5775a47aad0166d8e2fbcdd4be0dbebf1ced9fe4fc9a2abdbced"></a>

## Direct properties — custom_network_config.slo_config.static_routes.static_routes.default_gateway / 93f007ce39df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-359c7270cdf2cc9b044d38eeee41d3897dff524f4d15a7529c156fa944973bb4"></a>

## Next pages — custom_network_config.slo_config.static_routes.static_routes.default_gateway / 93f007ce39df / 4

- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-021e207896f45efe583d32ed49cc1ed6cf809f6db9c9b52d0c65130d87033128)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-28222d58951046048ba30d8ed3150d79bb89cb3c15dde18992866096a4b3843b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3e4a4aa778aaca7779671b6f8d411895e7eab536138a45146840a9b01df9da2"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface — custom_network_config.slo_config.static_routes.static_routes.node_interface / 75b2d951e48f / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-ebd2256dc9a55c693f5d456be7cb12d1a12a3ecbf41fb0628e8266095b5e1772)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-021e207896f45efe583d32ed49cc1ed6cf809f6db9c9b52d0c65130d87033128)
- custom_network_config.slo_config.static_routes.static_routes.node_interface

<a id="canonical-35ac2a106b98ac51d4e79f38451206bba0e4dfbead7e9366db7c4c798403af81"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

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

<a id="canonical-75bfa3acda3c944a0d97190ac5d5a46483ba76ffe2262aaa83f5feafc32a5894"></a>

## Direct properties — custom_network_config.slo_config.static_routes.static_routes.node_interface / 75b2d951e48f / 3

- [list](data-sources--securemesh_site--reference--group-003.md#canonical-5f3949c65f41755db6afaf606e103f5dde2bb805711680d184c9c6f79355121f): complete subsection reference.

<a id="canonical-0960525fb1956bad780e57abc02e9280eb060970ce79231c0e82c57d4861216c"></a>

## Next pages — custom_network_config.slo_config.static_routes.static_routes.node_interface / 75b2d951e48f / 4

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-5f3949c65f41755db6afaf606e103f5dde2bb805711680d184c9c6f79355121f)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-021e207896f45efe583d32ed49cc1ed6cf809f6db9c9b52d0c65130d87033128)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-5f3949c65f41755db6afaf606e103f5dde2bb805711680d184c9c6f79355121f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c30788ff20709f4a03d8d0fea97f2d89d9522f2fb65381570427dd2d35024d44"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface.list — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / d186d2381c00 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-ebd2256dc9a55c693f5d456be7cb12d1a12a3ecbf41fb0628e8266095b5e1772)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-021e207896f45efe583d32ed49cc1ed6cf809f6db9c9b52d0c65130d87033128)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-28222d58951046048ba30d8ed3150d79bb89cb3c15dde18992866096a4b3843b)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list

<a id="canonical-a82d2e4d99e655751c07eb6219fc7d4d05630a788e1c9a6e879fe3f6bdc16016"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-77e8c5141c2c447b0b1c3be4374fbc6d5e3910b74c124cd6a7dfacc55cea9466"></a>

## Direct properties — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / d186d2381c00 / 3

- [interface](data-sources--securemesh_site--reference--group-003.md#canonical-13301d397e5fee574e6c8287b78bf5d616cc2b4a6e5fcd3dcf6f78084209e0a8): complete subsection reference.

<a id="canonical-3bdf1d146b591d9d6d9a2dad83a39beb2a3ed0c8c60f9917107239a053792aa4"></a>

<a id="canonical-b688857001aa713d4a117d26310d93bf8b9b0b8ed3ea01530f7a30412a5672c4"></a>

## node property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / d186d2381c00 / 4

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-f61b5f652ec4500c43b66523354621b83dece001a4e1bf5528c20701a8255ffe"></a>

## Next pages — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / d186d2381c00 / 5

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-003.md#canonical-13301d397e5fee574e6c8287b78bf5d616cc2b4a6e5fcd3dcf6f78084209e0a8)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-28222d58951046048ba30d8ed3150d79bb89cb3c15dde18992866096a4b3843b)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-13301d397e5fee574e6c8287b78bf5d616cc2b4a6e5fcd3dcf6f78084209e0a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e6136feaa375d63a9e71c58e8476a1d561b5283e8c70ccdc15685ff6544b480"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / d3e4a68c33f8 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-ebd2256dc9a55c693f5d456be7cb12d1a12a3ecbf41fb0628e8266095b5e1772)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-021e207896f45efe583d32ed49cc1ed6cf809f6db9c9b52d0c65130d87033128)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-28222d58951046048ba30d8ed3150d79bb89cb3c15dde18992866096a4b3843b)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-5f3949c65f41755db6afaf606e103f5dde2bb805711680d184c9c6f79355121f)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-6599f5d20b584c63491150bfb78ba7b9fdc9f6119a2bb39ec49184483bb12f5b"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-59407cde35f0db03c2ad27aa964a14c6ad6ca49040ffb46a355d3bc8ba20a036"></a>

## Direct properties — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / d3e4a68c33f8 / 3

<a id="canonical-b0480853f4d726a558c5ec360a52d902007630392ad761143b154620950ca901"></a>

<a id="canonical-dbee337ae6ca6261c28cc0c3d04482e74382d726148aa3205bfcbd40b17cb166"></a>

## kind property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / d3e4a68c33f8 / 4

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

<a id="canonical-3c844f8299751bf4fbc53d16c2e0fde436b2502183b464ed5f39b48fa1012361"></a>

<a id="canonical-68d0a6462a69bd3de0673585a9311070627d07b359f0e30f2efbbf315ab42def"></a>

## name property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / d3e4a68c33f8 / 5

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

<a id="canonical-38069ee7489f90b0ed0242c9f2d84bb668265540f94da535ac3c87c197435dc8"></a>

<a id="canonical-28de511b4db83ae0e827537e4a198e24c12afb27b178a4acb6e68f26b7433fa4"></a>

## namespace property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / d3e4a68c33f8 / 6

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

<a id="canonical-e68b9204b155d318d5c9f4c41338b5b04b1e32d5cb69d6c9f9ddebbb1ff42f89"></a>

<a id="canonical-d3a495d4e4a4ff0873d439717db104abb69c4affada6a07432adacef502e2ac3"></a>

## tenant property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / d3e4a68c33f8 / 7

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

<a id="canonical-a3da5e05dd13b314f814e4abbe2904748960b5ad2c6beef72d59d917fee6fe00"></a>

<a id="canonical-69698a6363ff35b5722f8097017db398aff11dd46ae34848f77b5fada5e8d46a"></a>

## uid property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / d3e4a68c33f8 / 8

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

<a id="canonical-68d34114747946ccd09c9fff4b2e568cd4e7a78dea485ee0f665fa42dbaee8c4"></a>

## Next pages — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / d3e4a68c33f8 / 9

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-5f3949c65f41755db6afaf606e103f5dde2bb805711680d184c9c6f79355121f)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-4264d8bbba705a76f895a4e46cc501f9d26b3c652823436945770b53d6ec8cfd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2ff95280e6042ab5c27900e4e412e6ef78aab4add966df1ace1f0568019448b"></a>

## custom_network_config.slo_config.static_v6_routes — custom_network_config.slo_config.static_v6_routes / 8d66b1bb2b77 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- custom_network_config.slo_config.static_v6_routes

<a id="canonical-662415e19a270a5840fd84198c70e2a7e7022221911c96ba7dd02164251acd12"></a>

Type: `"single"`. Computed.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-31c355444c9e9a263299a2c4e0d0ee5a85fa311018e7d32333387ed87abbecfb"></a>

## Direct properties — custom_network_config.slo_config.static_v6_routes / 8d66b1bb2b77 / 3

- [static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-6c77e115eccc7dd11160a5f59067c8b16c4b8d8325e592126688da59a876f6f6): complete subsection reference.

<a id="canonical-9a372d205c8153563e764f236a4c29bf0006bccc4c62907bb4cc05ec0d74793d"></a>

## Next pages — custom_network_config.slo_config.static_v6_routes / 8d66b1bb2b77 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-6c77e115eccc7dd11160a5f59067c8b16c4b8d8325e592126688da59a876f6f6)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-6c77e115eccc7dd11160a5f59067c8b16c4b8d8325e592126688da59a876f6f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82ed232065f280121d05c9c4f7d942302ceb2957a090da9600a98b854219b0a3"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes — custom_network_config.slo_config.static_v6_routes.static_routes / d45c076af1a9 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- [custom_network_config.slo_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-4264d8bbba705a76f895a4e46cc501f9d26b3c652823436945770b53d6ec8cfd)
- custom_network_config.slo_config.static_v6_routes.static_routes

<a id="canonical-66072100241f318c771d752b8152e1fee090301578eb1dddf61fff7246758684"></a>

Type: `"list"`. Computed.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-00d146bee1bf2beccc7f243a6f7f6bd8a664658b392491a2c7acbacd7a6bbe82"></a>

## Direct properties — custom_network_config.slo_config.static_v6_routes.static_routes / d45c076af1a9 / 3

<a id="canonical-7250af9de5b3c385a73b3b2ea300a4d48d61cb46dc15f3f52d840f03eab4986e"></a>

<a id="canonical-39ff4481d78fe0ca6f6b5a626f61d7a6b3108bf783f594ce57ef53d8f8703e65"></a>

## attrs property — custom_network_config.slo_config.static_v6_routes.static_routes / d45c076af1a9 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-6f03af4c10afa3314c6502d03cc71f475a6aceea1361a0491d3974c947414ffd): complete subsection reference.

<a id="canonical-dac2a660d6add17fe6c48cc8b2fd52b0ab2dd8a4eb104a157b432e420edaa0c6"></a>

<a id="canonical-003dc8ffac42ae53535e5a6ab6100ada905efb3be3257bd30c7a354025cbe174"></a>

## ip_address property — custom_network_config.slo_config.static_v6_routes.static_routes / d45c076af1a9 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-9288327cfc719dfb0830ff12abe6f9d07dca161a1911e056d6f78239b0e3eaac"></a>

<a id="canonical-f95f5a21bfeed9871d47f65d5d241b0c659726ab660f43076c7c23c43263b282"></a>

## ip_prefixes property — custom_network_config.slo_config.static_v6_routes.static_routes / d45c076af1a9 / 6

Type: `["list", "string"]`. Computed.

List of IPv6 route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--securemesh_site--reference--group-004.md#canonical-3d219709b701bf9e95100e77a733e22105b8758ec89aaec49d717f58beb7d8e7): complete subsection reference.

<a id="canonical-f77a1c6788bea59a2464de19509942ebc61e06a99b9372de603e4adad40cb1be"></a>

## Next pages — custom_network_config.slo_config.static_v6_routes.static_routes / d45c076af1a9 / 7

- [custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-6f03af4c10afa3314c6502d03cc71f475a6aceea1361a0491d3974c947414ffd)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-004.md#canonical-3d219709b701bf9e95100e77a733e22105b8758ec89aaec49d717f58beb7d8e7)
- [custom_network_config.slo_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-4264d8bbba705a76f895a4e46cc501f9d26b3c652823436945770b53d6ec8cfd)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-6f03af4c10afa3314c6502d03cc71f475a6aceea1361a0491d3974c947414ffd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
