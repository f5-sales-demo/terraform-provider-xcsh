---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-38669884ba32794378e4ea6d541d4ee3863e674ede2c8b6abc0ad8c7c3a12e00"></a>

## Next pages — rate_limit.rate_limiter / 619620e59043 / 8

- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d124da01c1ef9b16adece36dd23a5191866727187304f9c9d172ea389acd0f5f)
- [rate_limit.rate_limiter.disabled](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3dcb558dea7ce916f272501c5f6e4fb9972d6319f58168bd149a734799dbf4b8)
- [rate_limit.rate_limiter.leaky_bucket](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-e1f3f116664ad3323c1dd3b41f457eeaac6fbd6879e74ed2cc1614fd1a2a5cde)
- [rate_limit.rate_limiter.token_bucket](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-90be89f4fd03ab867e72c129359243efbbaadce07f7e14b541685d2f399ce4c3)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-12bed4f6c3fd4f4da68f6b705b22bddc00b50871ecdae9367b04e4df3ea421b3)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-d124da01c1ef9b16adece36dd23a5191866727187304f9c9d172ea389acd0f5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3d7811878db206593ce430c4c7b668c8fb0c1a6594b417e6a7cd784fb858c39"></a>

## rate_limit.rate_limiter.action_block — rate_limit.rate_limiter.action_block / 50a2182edab9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-12bed4f6c3fd4f4da68f6b705b22bddc00b50871ecdae9367b04e4df3ea421b3)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-15f8f1d719a365f9cad1f0d3ac0197709eb52844f4542ccccc9043745bdf5db3)
- rate_limit.rate_limiter.action_block

<a id="canonical-c3fdb06b41b618bf029a2ee188dad3c997f2204abd2764e2ab091a3c84acc88f"></a>

Type: `"single"`. Computed.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

<a id="canonical-3ac28401a665ff4c2f97dd9a0acd408094aba7dc72ca3ef045e5e35c87786ab9"></a>

## Direct properties — rate_limit.rate_limiter.action_block / 50a2182edab9 / 3

- [hours](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-e56fbd3845d676a123dce13dd77aaac2ce6e97857742d57688c533b97e156155): complete subsection reference.

- [minutes](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-a85b1a09283c96b2218ccd4bfc031c95467e49ef99e21f3fafa3640c5cf1e40b): complete subsection reference.

- [seconds](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-ed3c6a8498bb9fd636194cb787ad9b20de5d3119e800312969274e44a7995834): complete subsection reference.

<a id="canonical-45045bf6327dd5783c5d552adc72fb80dafbee3dddc15cab858e340875cb0262"></a>

## Next pages — rate_limit.rate_limiter.action_block / 50a2182edab9 / 4

- [rate_limit.rate_limiter.action_block.hours](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-e56fbd3845d676a123dce13dd77aaac2ce6e97857742d57688c533b97e156155)
- [rate_limit.rate_limiter.action_block.minutes](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-a85b1a09283c96b2218ccd4bfc031c95467e49ef99e21f3fafa3640c5cf1e40b)
- [rate_limit.rate_limiter.action_block.seconds](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-ed3c6a8498bb9fd636194cb787ad9b20de5d3119e800312969274e44a7995834)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-15f8f1d719a365f9cad1f0d3ac0197709eb52844f4542ccccc9043745bdf5db3)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-e56fbd3845d676a123dce13dd77aaac2ce6e97857742d57688c533b97e156155"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f61d7cc45707a6e5f945fd3978ef6556de127a44d18abc0368807ccb1a22a40d"></a>

## rate_limit.rate_limiter.action_block.hours — rate_limit.rate_limiter.action_block.hours / 2ccbdb1e438c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-12bed4f6c3fd4f4da68f6b705b22bddc00b50871ecdae9367b04e4df3ea421b3)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-15f8f1d719a365f9cad1f0d3ac0197709eb52844f4542ccccc9043745bdf5db3)
- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d124da01c1ef9b16adece36dd23a5191866727187304f9c9d172ea389acd0f5f)
- rate_limit.rate_limiter.action_block.hours

<a id="canonical-a2366f76746a586cc1df2c9e9c1da5693e746ab77566e930cc6a4be58606e5a1"></a>

Type: `"single"`. Computed.

Hours. Input Duration Hours.

Upstream description:

Input Duration Hours.

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

<a id="canonical-5b48c953a091d431bbfbe82bbff21af745ace85a44f2034706e2cb1d55746f4c"></a>

## Direct properties — rate_limit.rate_limiter.action_block.hours / 2ccbdb1e438c / 3

<a id="canonical-3504eb1a6de7abffd6e507f63271e491b12f55d26eac84e7d7deb76bbf849a2f"></a>

<a id="canonical-2ef85de742510340061aca87a58ad85663b5fd7f9ccb95a2164e229b99b7c585"></a>

## duration property — rate_limit.rate_limiter.action_block.hours / 2ccbdb1e438c / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 48,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  }
}
```

<a id="canonical-1ab76aa6899b30d5f3eb38faf5e40bf28a71650a2ea56f0864600939eb6720a1"></a>

## Next pages — rate_limit.rate_limiter.action_block.hours / 2ccbdb1e438c / 5

- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d124da01c1ef9b16adece36dd23a5191866727187304f9c9d172ea389acd0f5f)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-a85b1a09283c96b2218ccd4bfc031c95467e49ef99e21f3fafa3640c5cf1e40b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75e2411a1049e9c091c6f7794e3254f6d0f8853cec7bc87c992ea9d649eac6f6"></a>

## rate_limit.rate_limiter.action_block.minutes — rate_limit.rate_limiter.action_block.minutes / d02d4873be86 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-12bed4f6c3fd4f4da68f6b705b22bddc00b50871ecdae9367b04e4df3ea421b3)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-15f8f1d719a365f9cad1f0d3ac0197709eb52844f4542ccccc9043745bdf5db3)
- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d124da01c1ef9b16adece36dd23a5191866727187304f9c9d172ea389acd0f5f)
- rate_limit.rate_limiter.action_block.minutes

<a id="canonical-a2c62c7c442456b9eec8e1d1b4ea7bc63332c08e9aa065243f1a65167316e113"></a>

Type: `"single"`. Computed.

Minutes. Input Duration Minutes.

Upstream description:

Input Duration Minutes.

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

<a id="canonical-834fb9c4a098fc1aa2f677afcbebeee9e01d2cc16530a6f5c3bc5d1dd66d7c84"></a>

## Direct properties — rate_limit.rate_limiter.action_block.minutes / d02d4873be86 / 3

<a id="canonical-4d42b95c889b6bbb6d434e36aec92d17b6ff39b64ec2f8c35df6ff9526d5aa00"></a>

<a id="canonical-f3e4987f0c1269edaf7a845b9ca580344f82bb521c09c4c31e67c1e970c49e74"></a>

## duration property — rate_limit.rate_limiter.action_block.minutes / d02d4873be86 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  }
}
```

<a id="canonical-dc05b75add90d2cdbcf658047a959183d9376aed837d234c1937d9dd7dd819d6"></a>

## Next pages — rate_limit.rate_limiter.action_block.minutes / d02d4873be86 / 5

- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d124da01c1ef9b16adece36dd23a5191866727187304f9c9d172ea389acd0f5f)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ed3c6a8498bb9fd636194cb787ad9b20de5d3119e800312969274e44a7995834"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-376816e0fa8324bfb5d4dcadef7fcd1f70d5b8e42ac5e3072f8096feeb187177"></a>

## rate_limit.rate_limiter.action_block.seconds — rate_limit.rate_limiter.action_block.seconds / ab0333fd41c1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-12bed4f6c3fd4f4da68f6b705b22bddc00b50871ecdae9367b04e4df3ea421b3)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-15f8f1d719a365f9cad1f0d3ac0197709eb52844f4542ccccc9043745bdf5db3)
- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d124da01c1ef9b16adece36dd23a5191866727187304f9c9d172ea389acd0f5f)
- rate_limit.rate_limiter.action_block.seconds

<a id="canonical-b65e7eec7e4dea48067f782844ca592ed07f61f5658eb60eddd8ba3f9180d56f"></a>

Type: `"single"`. Computed.

Seconds. Input Duration Seconds.

Upstream description:

Input Duration Seconds.

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

<a id="canonical-0ef03c6fef171a4e9007b1ef3db20f2a8e00bca2424f41407660c46f63a545dc"></a>

## Direct properties — rate_limit.rate_limiter.action_block.seconds / ab0333fd41c1 / 3

<a id="canonical-72614962d854c345644cf0f0692c0e9cf02849ecc9df69a8c120b9bdd6dc76ee"></a>

<a id="canonical-22850b83e9b1285f664fb3dd9f4eda0a2a07b9ca573f4b9a8cacf15078d21a86"></a>

## duration property — rate_limit.rate_limiter.action_block.seconds / ab0333fd41c1 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-3b5dc56585f743e969ee5b3bc6746ee988f43b7074fa835ec0ca07a753dbe00d"></a>

## Next pages — rate_limit.rate_limiter.action_block.seconds / ab0333fd41c1 / 5

- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d124da01c1ef9b16adece36dd23a5191866727187304f9c9d172ea389acd0f5f)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3dcb558dea7ce916f272501c5f6e4fb9972d6319f58168bd149a734799dbf4b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de1e91d35564816d412f64b2410c12201c34e2790c5b8b709c5633595b0221b2"></a>

## rate_limit.rate_limiter.disabled — rate_limit.rate_limiter.disabled / ad473e50611c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-12bed4f6c3fd4f4da68f6b705b22bddc00b50871ecdae9367b04e4df3ea421b3)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-15f8f1d719a365f9cad1f0d3ac0197709eb52844f4542ccccc9043745bdf5db3)
- rate_limit.rate_limiter.disabled

<a id="canonical-f85baa6108235b7177a49c7b2aa274cde8ed31c41cc05861a4907ec4cf8ddb0e"></a>

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

<a id="canonical-481fbe57d6b41b9683cc163abee3f90377fa22ab1ea19ea48555f3c05f930f4b"></a>

## Direct properties — rate_limit.rate_limiter.disabled / ad473e50611c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a930f52dd9d8bf06cdb3813022cd00aff8d8e78410de64bbb0224cc784a715e"></a>

## Next pages — rate_limit.rate_limiter.disabled / ad473e50611c / 4

- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-15f8f1d719a365f9cad1f0d3ac0197709eb52844f4542ccccc9043745bdf5db3)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-e1f3f116664ad3323c1dd3b41f457eeaac6fbd6879e74ed2cc1614fd1a2a5cde"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb0d476de8a2f82faf5ab0a7de07df0c7926b81e37b223add4017ba85e5b1938"></a>

## rate_limit.rate_limiter.leaky_bucket — rate_limit.rate_limiter.leaky_bucket / caf139f8e2b7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-12bed4f6c3fd4f4da68f6b705b22bddc00b50871ecdae9367b04e4df3ea421b3)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-15f8f1d719a365f9cad1f0d3ac0197709eb52844f4542ccccc9043745bdf5db3)
- rate_limit.rate_limiter.leaky_bucket

<a id="canonical-f9313f5a93d8fced5c560d523fa2e7b2906914c03fd7048941f5932e23378db9"></a>

Type: `["object", {}]`. Computed.

Leaky-Bucket is the default rate limiter algorithm for F5.

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

<a id="canonical-ee3188331ef62f4f1adeee915bcd5383c9597b06b38277ef3ba00de6c66fd718"></a>

## Direct properties — rate_limit.rate_limiter.leaky_bucket / caf139f8e2b7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-614d77aabbb89083d11cbaddb5f45f5ebabf0d4c92d4c23c3ad269e0fa8bb1db"></a>

## Next pages — rate_limit.rate_limiter.leaky_bucket / caf139f8e2b7 / 4

- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-15f8f1d719a365f9cad1f0d3ac0197709eb52844f4542ccccc9043745bdf5db3)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-90be89f4fd03ab867e72c129359243efbbaadce07f7e14b541685d2f399ce4c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36822e19c43e44c246acbe7e6a07cc552036efd5edf0aa85275f1f8a12d36cab"></a>

## rate_limit.rate_limiter.token_bucket — rate_limit.rate_limiter.token_bucket / 3037117189e3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-12bed4f6c3fd4f4da68f6b705b22bddc00b50871ecdae9367b04e4df3ea421b3)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-15f8f1d719a365f9cad1f0d3ac0197709eb52844f4542ccccc9043745bdf5db3)
- rate_limit.rate_limiter.token_bucket

<a id="canonical-99132d1d968f5a8ddfcac4798a1f97452c6dffb17216389acd7e945eb7e5fea8"></a>

Type: `["object", {}]`. Computed.

Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.

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

<a id="canonical-cfb4bacd6d7fceea411ef2c2767d8083580039b4f4d20fa079f8e52c1c8f3009"></a>

## Direct properties — rate_limit.rate_limiter.token_bucket / 3037117189e3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-825092d63cf0716fed5b34115689a3a5aaa187858abefddb793648886899b1e2"></a>

## Next pages — rate_limit.rate_limiter.token_bucket / 3037117189e3 / 4

- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-15f8f1d719a365f9cad1f0d3ac0197709eb52844f4542ccccc9043745bdf5db3)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-bb2647b3428854666f974efe2d4accde407916594bdd794de0a14630e2c93d77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ae708192eabd2aec3f58d8f6b0707163bc196c0be892ce6352040bee5c43e7a"></a>

## sensitive_data_policy — sensitive_data_policy / 28181bb8daa8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- sensitive_data_policy

<a id="canonical-9abebbf10799beafcd7cf8ee9287f1ff4c96552d46c31c76ef8d1aa266f6081a"></a>

Type: `"single"`. Computed.

Policy configuration for this feature.

Upstream description:

Settings for data type policy.

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

<a id="canonical-2a1bbb8345d6e0407e89e658d271f8aaa5e12ec60a7ae1273521fa2706ab9775"></a>

## Direct properties — sensitive_data_policy / 28181bb8daa8 / 3

- [sensitive_data_policy_ref](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1a2606e1bf2a84cd4371bc149312a90afd7e6d1f64189f850e37126fc6a6ac95): complete subsection reference.

<a id="canonical-0c05ac51991dc78d021b46ef9230619bedb5fcd8d99651a9985c421b61d7ef78"></a>

## Next pages — sensitive_data_policy / 28181bb8daa8 / 4

- [sensitive_data_policy.sensitive_data_policy_ref](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1a2606e1bf2a84cd4371bc149312a90afd7e6d1f64189f850e37126fc6a6ac95)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-1a2606e1bf2a84cd4371bc149312a90afd7e6d1f64189f850e37126fc6a6ac95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85547ad9d220c0c27e8ff7e4707ffb85e5d09de6f1f941ede02c1cbdd6345760"></a>

## sensitive_data_policy.sensitive_data_policy_ref — sensitive_data_policy.sensitive_data_policy_ref / cc2592aa09bb / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [sensitive_data_policy](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-bb2647b3428854666f974efe2d4accde407916594bdd794de0a14630e2c93d77)
- sensitive_data_policy.sensitive_data_policy_ref

<a id="canonical-ad78acec19f56daa4a9275de3fed8ad4af3fbf00579247046d93de53e85c55ef"></a>

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

<a id="canonical-a066f430727324e57be489f8524cbb953d95f41394da4a2d15730f8559f4b9ed"></a>

## Direct properties — sensitive_data_policy.sensitive_data_policy_ref / cc2592aa09bb / 3

<a id="canonical-7e20fcfd7761fe56d0600d257b1d32f434c9efdad4f6f7c1ca998391e62c4729"></a>

<a id="canonical-9b68169a61ca92dc5b636050f1159ba8cc0e3b80a55c2fa7b30e694f22a47c30"></a>

## name property — sensitive_data_policy.sensitive_data_policy_ref / cc2592aa09bb / 4

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

<a id="canonical-90cba270924ba3d9afd8b185a19f196eceb1243a64761d5ff126f0f65034b5be"></a>

<a id="canonical-2f71f0d387e091e9772f012bab1bbe40279b48ffa7bf129ca84f42e9cf45ae76"></a>

## namespace property — sensitive_data_policy.sensitive_data_policy_ref / cc2592aa09bb / 5

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

<a id="canonical-d1a0347f51104dc1bc5fbf0489a1faac8067d08e915921a896ead340436a344d"></a>

<a id="canonical-b54509487784d841c9035f1ac56c5079c65beea44c93049d2b1f8f55bc327487"></a>

## tenant property — sensitive_data_policy.sensitive_data_policy_ref / cc2592aa09bb / 6

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

<a id="canonical-6d1183f37454c1641f5c1b04a1960611b5201042559efa598f9cca9e24437e58"></a>

## Next pages — sensitive_data_policy.sensitive_data_policy_ref / cc2592aa09bb / 7

- [sensitive_data_policy](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-bb2647b3428854666f974efe2d4accde407916594bdd794de0a14630e2c93d77)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-87808731574bc92e863e99557fbbe8621d30dfe501227638348645fb3d6ae8c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71da83aa52618134ec9d89b23fd41a46e6b67a873fe95dfefb2b5bed0b78f04c"></a>

## service_policies_from_namespace — service_policies_from_namespace / 80c9d89fdf36 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- service_policies_from_namespace

<a id="canonical-7dc0a0c243c66844becbeecf77c078855045544b2dd0500df2013d9ab80529c8"></a>

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

<a id="canonical-3552eb60bd910a236dfeafddf58ad4a1998f2c54acd98a21e810ac2135eae632"></a>

## Direct properties — service_policies_from_namespace / 80c9d89fdf36 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b04aee07a575eabbb3b68a4c7d79e68ab9396c5c4a7fcd8178e61778af578021"></a>

## Next pages — service_policies_from_namespace / 80c9d89fdf36 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-01c67b15d3a42bb9a1e992d885e2aaad2b664fbfb5f776a3c0c6d39c912ab599"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ca39e44b74896d92d6b254b59b2f62d646972f5483d65facf5099f08325e161"></a>

## slow_ddos_mitigation — slow_ddos_mitigation / 945ca71d9475 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- slow_ddos_mitigation

<a id="canonical-0bd1c44b2c87c574cd9c6646d3dd741969a8419c15807f0ecd12a99ed9ea8726"></a>

Type: `"single"`. Computed.

\[OneOf: slow\_ddos\_mitigation, system\_default\_timeouts; Default: system\_default\_timeouts\]
'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Upstream description:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

OneOf alternatives in this subsection:

- [slow_ddos_mitigation](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0bd1c44b2c87c574cd9c6646d3dd741969a8419c15807f0ecd12a99ed9ea8726)
- [system_default_timeouts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-c1834fca7d3688e007c8392d9264acfb91fdc1360bc97b58cddf12709834c1d0)

Select alternatives according to the provider validators above.

<a id="canonical-846d6d674fb8fbbc0607bae49241c542614d1eca25eb61b8ec99c7de011887bf"></a>

## Direct properties — slow_ddos_mitigation / 945ca71d9475 / 3

- [disable_request_timeout](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-68a262afd885d09a9e39681126169f8301dca81207cb825cfa7ac054b8f2326b): complete subsection reference.

<a id="canonical-b801a20724cf7a2ea83da0587d6ce28b3eba4401e39cce46561663c13a0c8247"></a>

<a id="canonical-0a90f137560b1f2b46fa4ca3b46178fb2507e22942446cd2f7ac205ab8c35894"></a>

## request_headers_timeout property — slow_ddos_mitigation / 945ca71d9475 / 4

Type: `"number"`. Computed.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Upstream description:

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The default value is 10000 milliseconds. This setting provides protection against
Slowloris attacks.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-2b6c06174569a7306824019df479acb593c8ece1fab9a921169c846292bf74fe"></a>

<a id="canonical-b58ab33b8e2fca8ff72bbd8c068a7c3ffc137cd2b932ddcfada50c7a0ae1bc22"></a>

## request_timeout property — slow_ddos_mitigation / 945ca71d9475 / 5

Type: `"number"`. Computed.

Exclusive with \[disable\_request\_timeout\].

Upstream description:

Exclusive with \[disable\_request\_timeout\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

<a id="canonical-c079735557f604701b3152b6d45d57fc730172fea62a6c955e17da39a96bfd25"></a>

## Next pages — slow_ddos_mitigation / 945ca71d9475 / 6

- [slow_ddos_mitigation.disable_request_timeout](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-68a262afd885d09a9e39681126169f8301dca81207cb825cfa7ac054b8f2326b)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-68a262afd885d09a9e39681126169f8301dca81207cb825cfa7ac054b8f2326b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebbd55e230690441d62f1c0325d77f79b748a79922c40471efbcc1c9cc6451a7"></a>

## slow_ddos_mitigation.disable_request_timeout — slow_ddos_mitigation.disable_request_timeout / f8a076f89f83 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [slow_ddos_mitigation](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-01c67b15d3a42bb9a1e992d885e2aaad2b664fbfb5f776a3c0c6d39c912ab599)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-ba9de1453d4fbe9b40224dd2fb2ea3871c2fa00b14735e74d1f3f70bec1fbeb4"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable request timeout.

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

<a id="canonical-6f9eb4c8a10d38a1d68837ad2e4568acd6de42ef534737d7b6ec70f9eb328a2c"></a>

## Direct properties — slow_ddos_mitigation.disable_request_timeout / f8a076f89f83 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-391add60aaa23f11da3c898fe0f4a3552cb40109e1600dca620dbe927e008ec3"></a>

## Next pages — slow_ddos_mitigation.disable_request_timeout / f8a076f89f83 / 4

- [slow_ddos_mitigation](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-01c67b15d3a42bb9a1e992d885e2aaad2b664fbfb5f776a3c0c6d39c912ab599)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-fa38e48238d43a1d52379d97249f1ee098d2c713f9459bf92678b51c370bbf0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0939dd63f9ede2d2a37c53fa0cc15194509963f63bb387ed231312fad419222"></a>

## system_default_timeouts — system_default_timeouts / 7199ddc5247e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- system_default_timeouts

<a id="canonical-c1834fca7d3688e007c8392d9264acfb91fdc1360bc97b58cddf12709834c1d0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for system default timeouts.

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

<a id="canonical-80634fe550c231fb9e6e4db27f1cfbfc7fb9b1c0352ff37aee4e62aa87f72b04"></a>

## Direct properties — system_default_timeouts / 7199ddc5247e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2697b6215e6ca8bf5a9f20a7cc0de1983f277dcb22734d88fcc5de074d5e1715"></a>

## Next pages — system_default_timeouts / 7199ddc5247e / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-42424e14e1fa4134094c259fc2bdc475057248d372dc4b9a2620501636ee872e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-851d3f32447be8b4ac07e2e9187e55a70abbb9d47676749fbded9f48db09be5c"></a>

## trusted_clients — trusted_clients / 845275e99a46 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- trusted_clients

<a id="canonical-6d0ab8d692ec615198931e406662266869210c5139def00dce2d4a4cec3eeca3"></a>

Type: `"list"`. Computed.

Define rules to skip processing of one or more features such as WAF, Bot Defense etc.

Upstream description:

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-cc775ae3b90fb261dc9673a96c8288e7e870a9df35461aeea3ef08f0a9a405b6"></a>

## Direct properties — trusted_clients / 845275e99a46 / 3

<a id="canonical-53717b57bf5b0157f51345a425ae7292e3b6bc2b4dce9aadd89f0ed1147c8a59"></a>

<a id="canonical-345d4b0b69f19ae70e23730be3d61188cdf9c6394a8a35f538aa3e844cdde5c6"></a>

## actions property — trusted_clients / 845275e99a46 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

Upstream description:

Actions that should be taken when client identifier matches the rule.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-fb90e71f922501614d73565e9c519a9c8ebb6e1d50020840688dbf7c241cf24c"></a>

<a id="canonical-4f90b12863661cd33dce81ea8c73672fddb62f0188e75a86c621071411b3ced4"></a>

## as_number property — trusted_clients / 845275e99a46 / 5

Type: `"number"`. Computed.

Exclusive with \[http\_header ip\_prefix ipv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Upstream description:

Exclusive with \[http\_header ip\_prefix ipv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 401308,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  }
}
```

- [bot_skip_processing](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-42a4bdbd9a88a6f3e313e1197f8f42a6db90c7b88e770baea5aa62454b4846b3): complete subsection reference.

<a id="canonical-ece087aed8bd5fd81b29fcc2b8f773a6039f36c1040661fd03db382fa28dd2ae"></a>

<a id="canonical-8a534cdb84964b9d3c89b57d8003a16e1da7b2db133e9a9ed294a84fad27a178"></a>

## expiration_timestamp property — trusted_clients / 845275e99a46 / 6

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-fddc55f29f963941fd3e7253d91a0d8e15c92210943c1272edef0e93f5e9a199): complete subsection reference.

<a id="canonical-cb9518e45bee58a926e396eb91788855b0d187e2d8276fea63eab3d387beb2d7"></a>

<a id="canonical-af8551d4acf4144c78d6250593bc59c3cf6504a778fb8656e0701ecf7d84acef"></a>

## ip_prefix property — trusted_clients / 845275e99a46 / 7

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ipv6\_prefix user\_identifier\] IPv4 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header ipv6\_prefix user\_identifier\] IPv4 prefix string.

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

<a id="canonical-86f2c6f166a60596d8d07060d98525979321ffb102aa0f43bff59c49b20aaeab"></a>

<a id="canonical-5fb31e53671ca12b8c3fbd8a10ca27c932e1f276f18e391c7a96925687b5fc8f"></a>

## ipv6_prefix property — trusted_clients / 845275e99a46 / 8

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-ebb4900a7b52c2f381fd8dd8b7abe32adeda0f26874f0db53972e0ccfd66eb08): complete subsection reference.

- [skip_processing](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d90b1015ce5738ac1b1e44aa0ae9f6229465095c1ee34c18b9cc6e285375db6f): complete subsection reference.

<a id="canonical-5259e0b773fcca5808ead14efdd1ac305b642a493a48d28c390e8a1d6600d320"></a>

<a id="canonical-6c2a970a2f377f7f9d1506d0cc794d43b36408e573fa2996ca0fdc7d034464bd"></a>

## user_identifier property — trusted_clients / 845275e99a46 / 9

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ip\_prefix ipv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix ipv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [waf_skip_processing](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-f2628c7a337a03bb137162e94a0629a0372a74e887809d6842eb5f5d72f7bb03): complete subsection reference.

<a id="canonical-e39e21ef0be8d538b47444a02012099890543bf765c3894bd76c66db0c81ac6f"></a>

## Next pages — trusted_clients / 845275e99a46 / 10

- [trusted_clients.bot_skip_processing](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-42a4bdbd9a88a6f3e313e1197f8f42a6db90c7b88e770baea5aa62454b4846b3)
- [trusted_clients.http_header](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-fddc55f29f963941fd3e7253d91a0d8e15c92210943c1272edef0e93f5e9a199)
- [trusted_clients.metadata](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-ebb4900a7b52c2f381fd8dd8b7abe32adeda0f26874f0db53972e0ccfd66eb08)
- [trusted_clients.skip_processing](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d90b1015ce5738ac1b1e44aa0ae9f6229465095c1ee34c18b9cc6e285375db6f)
- [trusted_clients.waf_skip_processing](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-f2628c7a337a03bb137162e94a0629a0372a74e887809d6842eb5f5d72f7bb03)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-42a4bdbd9a88a6f3e313e1197f8f42a6db90c7b88e770baea5aa62454b4846b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9338d81ba27509d9ee79d68c552430229c9516cb2266107c3ffbfa050771765"></a>

## trusted_clients.bot_skip_processing — trusted_clients.bot_skip_processing / 811f6ac265d3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-42424e14e1fa4134094c259fc2bdc475057248d372dc4b9a2620501636ee872e)
- trusted_clients.bot_skip_processing

<a id="canonical-7028595587fc7fec91ab2733af2a1f6bd3dabfcc996266a5c35d04116ee1cf03"></a>

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

<a id="canonical-7714ed3d41dbcde3b0c3255c5e0515b1c886cca86501be0f1a5a5fb87e990d66"></a>

## Direct properties — trusted_clients.bot_skip_processing / 811f6ac265d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-46f25ba80e7c60288fbade3d6fc23fe0d3a16138f46a3ad3953949cfbbe2604a"></a>

## Next pages — trusted_clients.bot_skip_processing / 811f6ac265d3 / 4

- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-42424e14e1fa4134094c259fc2bdc475057248d372dc4b9a2620501636ee872e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-fddc55f29f963941fd3e7253d91a0d8e15c92210943c1272edef0e93f5e9a199"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e3b5545076dcd932abdbe9bd6b25f141f754d146896cedf6a4401938c2789f6"></a>

## trusted_clients.http_header — trusted_clients.http_header / a9661894a979 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-42424e14e1fa4134094c259fc2bdc475057248d372dc4b9a2620501636ee872e)
- trusted_clients.http_header

<a id="canonical-4f5b06d9adc6868f98addda209160601f07dc74b9dacceebfb1fa81014e1bf88"></a>

Type: `"single"`. Computed.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

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

<a id="canonical-be00b344c339099952a112c2027ddd6e915b1566dd9f9525cb36b189e5677a7b"></a>

## Direct properties — trusted_clients.http_header / a9661894a979 / 3

- [headers](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-e3967ae5674111b7d334b80b3c2e59e5b5bcb407ee79ea374381839d542d07b8): complete subsection reference.

<a id="canonical-0b188fb90b0ff863b591c42e9ed282ae266981ef9ea0980eb8eba43a39170565"></a>

## Next pages — trusted_clients.http_header / a9661894a979 / 4

- [trusted_clients.http_header.headers](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-e3967ae5674111b7d334b80b3c2e59e5b5bcb407ee79ea374381839d542d07b8)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-42424e14e1fa4134094c259fc2bdc475057248d372dc4b9a2620501636ee872e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-e3967ae5674111b7d334b80b3c2e59e5b5bcb407ee79ea374381839d542d07b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-260075a61899f615336ae0170afc6a1c830746d665c3e35d2b20c46dd98ffa2a"></a>

## trusted_clients.http_header.headers — trusted_clients.http_header.headers / 64affa8fa205 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-42424e14e1fa4134094c259fc2bdc475057248d372dc4b9a2620501636ee872e)
- [trusted_clients.http_header](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-fddc55f29f963941fd3e7253d91a0d8e15c92210943c1272edef0e93f5e9a199)
- trusted_clients.http_header.headers

<a id="canonical-55e4802e81ea8e918debc3bec237667a2adfb6a183c9ca23b6cf024e5a7f26b4"></a>

Type: `"list"`. Computed.

List of HTTP header name and value pairs.

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
    "minItems": 0,
    "uniqueItems": false
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

<a id="canonical-866028f7b243b0eb5a4d9875061daea4faa02aa7ae950f76511df8d72b4458f5"></a>

## Direct properties — trusted_clients.http_header.headers / 64affa8fa205 / 3

<a id="canonical-0e96a0f88986753a3ffe2b2f132dcc839ec89de2da31a79cb286468e0377a881"></a>

<a id="canonical-e09ac69c540955b3f2762ef1a56e69fe1b4688ad249ec9646afeb41c01c19dbe"></a>

## exact property — trusted_clients.http_header.headers / 64affa8fa205 / 4

Type: `"string"`. Computed.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-7ed77920e435cc566ee5bc62c297ff6032be4429bca2d69a1f7f1652e4aa3fbb"></a>

<a id="canonical-d49dcabd06455f0a9f3cfee93b7eafcbc99516e6beb6fc0b27a933a1446ba404"></a>

## invert_match property — trusted_clients.http_header.headers / 64affa8fa205 / 5

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-2cc319fffb272b8c22de9aa5be93984c40b3a2a43bac597cd037e878836ad66b"></a>

<a id="canonical-3e3cbba1c8dc21d62042c3f6e2f1f1cd14761d345d17c2f812fb23c0a65c194e"></a>

## name property — trusted_clients.http_header.headers / 64affa8fa205 / 6

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-7af84aaba4414733ff3fddb57e421f9f43775e4d8749ed748699ac6e0f02c230"></a>

<a id="canonical-0e2ff519ad342a4995433ecc19f45aa093551311c52f80906d323b9d665c957f"></a>

## presence property — trusted_clients.http_header.headers / 64affa8fa205 / 7

Type: `"bool"`. Computed.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

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

<a id="canonical-4ba2155b25ff19c02a8991a801c0c458f85f8bd27ee8e1bc2a490f1cd0f4726f"></a>

<a id="canonical-5766cca3514fcb87c2fac01c3fd7bc4555fa0d31340aa77864b43f75ab7381c0"></a>

## regex property — trusted_clients.http_header.headers / 64affa8fa205 / 8

Type: `"string"`. Computed.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-578c4d6be7ca7002b19f8a39ed819bff7b3d3bcf9d5b8925e97bd91f432f591a"></a>

## Next pages — trusted_clients.http_header.headers / 64affa8fa205 / 9

- [trusted_clients.http_header](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-fddc55f29f963941fd3e7253d91a0d8e15c92210943c1272edef0e93f5e9a199)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ebb4900a7b52c2f381fd8dd8b7abe32adeda0f26874f0db53972e0ccfd66eb08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ffe09f27da28cd59d62066f04de0a0ca12efd72a464034e08b96c74cf2bc3dd"></a>

## trusted_clients.metadata — trusted_clients.metadata / 1d82798d0392 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-42424e14e1fa4134094c259fc2bdc475057248d372dc4b9a2620501636ee872e)
- trusted_clients.metadata

<a id="canonical-05d0e789d8595fd9b5d5ff219e09af509020501358071a8fb828c8b2e136ba49"></a>

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

<a id="canonical-a67ff5b6f49e75324658dbc0c25dcd7bc6a7a5e282bf2b6ea4e37d9ad727a3e5"></a>

## Direct properties — trusted_clients.metadata / 1d82798d0392 / 3

<a id="canonical-70f47cc33fcc3197a05836b2d210053c81a6abed16b318dd32bf2d9152aae556"></a>

<a id="canonical-0efdcb78096023dddb19337ed8c627dc71bd510c9084f4a76aea9390dcc3c9a3"></a>

## description_spec property — trusted_clients.metadata / 1d82798d0392 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1a502ab45388bc9e543b7b3b0a9d87b4a8585c0ebc9e3598be33c32c7cc740b6"></a>

<a id="canonical-5813116d8da25893e9f7368a9ab8ea56e975a5134d3acc3d19c9244b850b2c63"></a>

## name property — trusted_clients.metadata / 1d82798d0392 / 5

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

<a id="canonical-916f6c14ea3601a25cacc659fb75efa2a0c9951c3221d6d8520fe17116c8333f"></a>

## Next pages — trusted_clients.metadata / 1d82798d0392 / 6

- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-42424e14e1fa4134094c259fc2bdc475057248d372dc4b9a2620501636ee872e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-d90b1015ce5738ac1b1e44aa0ae9f6229465095c1ee34c18b9cc6e285375db6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-202dbde667bc8c0dacf5b1e321d736548931f243a0b1bb911acec9cff1421324"></a>

## trusted_clients.skip_processing — trusted_clients.skip_processing / 2f8c0a068f8f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-42424e14e1fa4134094c259fc2bdc475057248d372dc4b9a2620501636ee872e)
- trusted_clients.skip_processing

<a id="canonical-b9e9e070334ef755393db6128c57c672f71d48dff7e929742b646b4611fedf7e"></a>

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

<a id="canonical-1297f13f2d489b58d2a0af6f3ee0f14aae32e8d712b68ef88a5e7ea413c6a2ce"></a>

## Direct properties — trusted_clients.skip_processing / 2f8c0a068f8f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4d0769b6279aff320777e21a9dcfee130144adec4e0befeb60f396cccefe78fb"></a>

## Next pages — trusted_clients.skip_processing / 2f8c0a068f8f / 4

- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-42424e14e1fa4134094c259fc2bdc475057248d372dc4b9a2620501636ee872e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f2628c7a337a03bb137162e94a0629a0372a74e887809d6842eb5f5d72f7bb03"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a13040009dc648cb052ec9218a358f9377384a996264187f355d8eb9f4db62a2"></a>

## trusted_clients.waf_skip_processing — trusted_clients.waf_skip_processing / a777281faf75 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-42424e14e1fa4134094c259fc2bdc475057248d372dc4b9a2620501636ee872e)
- trusted_clients.waf_skip_processing

<a id="canonical-c51e820c9f1dcc51c8db5af8e9f02118c304c69ce9a228376e23f247c0d01f7c"></a>

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

<a id="canonical-75211af8d2c9d3a730f659f022eb3600526f5d8032e9869e79f59bdd4da71d9c"></a>

## Direct properties — trusted_clients.waf_skip_processing / a777281faf75 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-89875fde585ddf91f9ef3a2c9a957f42d3b81b614bf027aa6319c3f97ff66898"></a>

## Next pages — trusted_clients.waf_skip_processing / a777281faf75 / 4

- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-42424e14e1fa4134094c259fc2bdc475057248d372dc4b9a2620501636ee872e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-51bae569a01abc71d9356bedb70367bcacc4ff50a745eb8a820ce86d9ebea852"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23f5398e13c7b8b3ee8a6c1b7bc86abba81afce51b52948c58cfd1141d51eda7"></a>

## user_id_client_ip — user_id_client_ip / c78c908dbda1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- user_id_client_ip

<a id="canonical-539813de5d793fa5d4913d449f4844b83da94b3814909f9563f8b7edceb40f1b"></a>

Type: `["object", {}]`. Computed.

\[OneOf: user\_id\_client\_ip, user\_identification\] Enable this option

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

OneOf alternatives in this subsection:

- [user_id_client_ip](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-539813de5d793fa5d4913d449f4844b83da94b3814909f9563f8b7edceb40f1b)
- [user_identification](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-cbad670f0c9caf9dd64727be048ba6860f1c9d02bfef71d63c2e826b18e08cf9)

Select alternatives according to the provider validators above.

<a id="canonical-4bb949fec71216db714b12a41ee0e7f151d55eaa4414fae9378623dbe9ebb6fd"></a>

## Direct properties — user_id_client_ip / c78c908dbda1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-95ce0b7514e0828417b849be9736f7e7bbead2e63274306f91b014707594aace"></a>

## Next pages — user_id_client_ip / c78c908dbda1 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-cc94a6d81f55468ce844b7b39e7af9c0e4a0c8d3b1eb5554b18b65b5a24008a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-074935c4c07e0d88f1478645a543277ffa88b179db2b9f4d28dc20b31b0cd199"></a>

## user_identification — user_identification / f2712f102983 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- user_identification

<a id="canonical-cbad670f0c9caf9dd64727be048ba6860f1c9d02bfef71d63c2e826b18e08cf9"></a>

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

<a id="canonical-b867a3c4dccaaf21882d696e52c06eb19ccc562a31b30ead7d7e8fea150495a8"></a>

## Direct properties — user_identification / f2712f102983 / 3

<a id="canonical-630196d983600e07c9be73f047b9ace615f134089eafaf8c997b32ebcaebe32a"></a>

<a id="canonical-6a35bdf7cfb2ceda5772111c1b1ce2b181bb708a0bb4bb8f9a19ce404814befd"></a>

## name property — user_identification / f2712f102983 / 4

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

<a id="canonical-b74a1353e8a8a36bdccb8eaeace130c97bc49dd66b0f7d0617a20dee4ef7b0be"></a>

<a id="canonical-3d61f9621c7ca9d2dfd3f203cd50eb4ab023f74de0415816d5fd3165a135545e"></a>

## namespace property — user_identification / f2712f102983 / 5

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

<a id="canonical-564cf3cc5ade402cae677eacd29e558d5cd1be6f04ed213e508592a04501c6ec"></a>

<a id="canonical-bd2c8003420ee2a00141a2af923863dce02e9eab96fbf8f2ca1f5ff686c26163"></a>

## tenant property — user_identification / f2712f102983 / 6

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

<a id="canonical-2250fcddcfc843ed3c8236b293e669807cb46e5841d44e2dc7a9faa6afd4c9b2"></a>

## Next pages — user_identification / f2712f102983 / 7

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-862eda4565194ec6d9383054ae3a86bcd188f17d78c18b1808c8c947cdf555d5"></a>

## waf_exclusion — waf_exclusion / f2981a26630f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- waf_exclusion

<a id="canonical-7bb5ae97147ef0a9af84a89f082dffbc49e5e59eecd5fec8553878710c11d462"></a>

Type: `"single"`. Computed.

Configuration parameter for waf exclusion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-waf_exclusion_choice": "[\"waf_exclusion_inline_rules\",\"waf_exclusion_policy\"]"
}
```

<a id="canonical-00f940f1530ad5a4797111b6f9f486928dd17097f23c47338bdbcd2718b3862f"></a>

## Direct properties — waf_exclusion / f2981a26630f / 3

- [waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d9194f080806e69f7a565ead1caaf1ae7df6c377a284846942501e5419ac7a8a): complete subsection reference.

- [waf_exclusion_policy](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3912975fd6ba1b1c42cd8ac7d5f5c825aaff722b2985b40c803582e43cebf9bc): complete subsection reference.

<a id="canonical-41e8f72d7252d19a4189af95a2b3b049d2a08ddde113b2386fb8943f0720b34a"></a>

## Next pages — waf_exclusion / f2981a26630f / 4

- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d9194f080806e69f7a565ead1caaf1ae7df6c377a284846942501e5419ac7a8a)
- [waf_exclusion.waf_exclusion_policy](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3912975fd6ba1b1c42cd8ac7d5f5c825aaff722b2985b40c803582e43cebf9bc)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-d9194f080806e69f7a565ead1caaf1ae7df6c377a284846942501e5419ac7a8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dd002a26687f0fd0bbb7e58e18ed101820067f00e776b59b841167cbd10d94c"></a>

## waf_exclusion.waf_exclusion_inline_rules — waf_exclusion.waf_exclusion_inline_rules / d94af97f25b6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37)
- waf_exclusion.waf_exclusion_inline_rules

<a id="canonical-017e2f9d2a83d69c71beb0645e95e082600b7396e0dc34476dd727cf6aacde1b"></a>

Type: `"single"`. Computed.

List of WAF exclusion rules that will be applied inline.

Upstream description:

A list of WAF exclusion rules that will be applied inline.

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

<a id="canonical-ddacb2c9070c805b927996bddc1f90ee52af65d23e9c0237dee0af16cb68ae6a"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules / d94af97f25b6 / 3

- [rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131): complete subsection reference.

<a id="canonical-9d6f9df1174b3beb58a89ab4265106c1d70b9552138d84e5b8374112ce68e46d"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules / d94af97f25b6 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49fc1cfa56d5d98a52b84d021170064868c9ec40744403b27089cd85ead9b992"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules — waf_exclusion.waf_exclusion_inline_rules.rules / 3d47cfffe6a5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d9194f080806e69f7a565ead1caaf1ae7df6c377a284846942501e5419ac7a8a)
- waf_exclusion.waf_exclusion_inline_rules.rules

<a id="canonical-af870df6a133783e02d3defb32dd6e35fd91433450669a8d515c555fc4095dc9"></a>

Type: `"list"`. Computed.

Ordered list of WAF Exclusions specific to this Load Balancer.

Upstream description:

An ordered list of WAF Exclusions specific to this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-40a35b7d9154e0e52245e49674d2577675f8f9f5c9dba14d7e9489825311936a"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules / 3d47cfffe6a5 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-de0980f57d96b32586c2d43525b123fcaa0a11db814bad95b5c23c78af159273): complete subsection reference.

- [any_path](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-c507bab9dfacd4c196e13dac77312b133c53ce4c0a230fd7f86c6662fdfc0da5): complete subsection reference.

- [app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-27cc87332cf786bb256085f80a133ae11c8d245a7dc594446dd0755c398b3f20): complete subsection reference.

<a id="canonical-b0f10672c171545a9254bf12ef14227412b70f38db62ac09c6035ba4e0a6b4bb"></a>

<a id="canonical-654f86f1524ed3705296656c9ca3b71ae3a3dcfb3c27ac3e4e67905e560bff8e"></a>

## exact_value property — waf_exclusion.waf_exclusion_inline_rules.rules / 3d47cfffe6a5 / 4

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0db905b83d3472e2b00519eaa2a4a59b5cfd4c7067fd097f193571b756a55fbc"></a>

<a id="canonical-1fc3d3eb2103db294f41dc699a17d21ffec8db21424be2b2d721d215f6de5c9f"></a>

## expiration_timestamp property — waf_exclusion.waf_exclusion_inline_rules.rules / 3d47cfffe6a5 / 5

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

- [metadata](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5e0b47d889e695190916ac1a54e116c2da17e3c55f834c19bfb1b13c710207da): complete subsection reference.

<a id="canonical-c548f82b4345f36bb950023537b0004d89c9ef26e66c746a88449defdc028354"></a>

<a id="canonical-6e53ed929b2f3d1386171ad366347ac3c44d214b6e9c0b8b1a76789ad6813b67"></a>

## methods property — waf_exclusion.waf_exclusion_inline_rules.rules / 3d47cfffe6a5 / 6

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-dba83873783ff315e3e72bcf6edfbbb2f86560f3673aa607ea16bb9384989e09"></a>

<a id="canonical-d39fde588c05ff8c20501b51de46e5ed10803fb4e1d3b07ac699617d605cc29f"></a>

## path_prefix property — waf_exclusion.waf_exclusion_inline_rules.rules / 3d47cfffe6a5 / 7

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

Upstream description:

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0dd1695fbb5d55f52a343c7b9c66fe05386f78ea3005f33d7a2f1ed2b1b35fb2"></a>

<a id="canonical-5a7cb06522c188929a25bba24510caf0740bdffdaca0bdfba6a413b7d192427c"></a>

## path_regex property — waf_exclusion.waf_exclusion_inline_rules.rules / 3d47cfffe6a5 / 8

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

Upstream description:

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-9a1a999c94caf863aae3de219f990da53c20b969ccc80b7997e3ad97b8e1e1ae"></a>

<a id="canonical-0029d8b7252a50191eb1f80a4ad8d5819571c8d766b00750d12b71b23f83e41a"></a>

## suffix_value property — waf_exclusion.waf_exclusion_inline_rules.rules / 3d47cfffe6a5 / 9

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [waf_skip_processing](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-4e12755d168e87df4b2231d6c1c9afd6fe31e031fca8fcecf7a799574ba003da): complete subsection reference.

<a id="canonical-8f5643412e0a48516722d3992d328d3199c8253420d52ad4167fc3c0a805c9b5"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules / 3d47cfffe6a5 / 10

- [waf_exclusion.waf_exclusion_inline_rules.rules.any_domain](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-de0980f57d96b32586c2d43525b123fcaa0a11db814bad95b5c23c78af159273)
- [waf_exclusion.waf_exclusion_inline_rules.rules.any_path](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-c507bab9dfacd4c196e13dac77312b133c53ce4c0a230fd7f86c6662fdfc0da5)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-27cc87332cf786bb256085f80a133ae11c8d245a7dc594446dd0755c398b3f20)
- [waf_exclusion.waf_exclusion_inline_rules.rules.metadata](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5e0b47d889e695190916ac1a54e116c2da17e3c55f834c19bfb1b13c710207da)
- [waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-4e12755d168e87df4b2231d6c1c9afd6fe31e031fca8fcecf7a799574ba003da)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d9194f080806e69f7a565ead1caaf1ae7df6c377a284846942501e5419ac7a8a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-de0980f57d96b32586c2d43525b123fcaa0a11db814bad95b5c23c78af159273"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a532aafb0a80e6f4d2491e985bae2f76fe87a726566631ee9b87d50f777e93c"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.any_domain — waf_exclusion.waf_exclusion_inline_rules.rules.any_domain / d23339f13a01 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d9194f080806e69f7a565ead1caaf1ae7df6c377a284846942501e5419ac7a8a)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_domain

<a id="canonical-dec6e194dcc7dcdd4e3ecce0c91228d04371bf130522d9b49ef1ae3d2a75c1dd"></a>

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

<a id="canonical-476879a0adc8fa8880efe4107917678b3c7caed26866092ab4d7894d9dc2d0ba"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.any_domain / d23339f13a01 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c4902ba0b97fdebe264ce9fba0d14c41aa41546588639d5f1c49b74cb800bbc6"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.any_domain / d23339f13a01 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-c507bab9dfacd4c196e13dac77312b133c53ce4c0a230fd7f86c6662fdfc0da5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f0f9266eb996ce21acf7dc57497f1576563a2f34ade9fad8948fdb37b18eb5e"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.any_path — waf_exclusion.waf_exclusion_inline_rules.rules.any_path / f2350c389a0a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d9194f080806e69f7a565ead1caaf1ae7df6c377a284846942501e5419ac7a8a)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_path

<a id="canonical-5c2246e8fbd262d2b49caed12183df939805497ed48c5da93b4dae54d146ce54"></a>

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

<a id="canonical-b89300c559d8d985c421b733ae21c8f949de622dcb6d0f20b06d801d6e9b07bd"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.any_path / f2350c389a0a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a2e1beb978314c63301db076b35db99c6349f8288805cb8f5df870d51a562a9b"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.any_path / f2350c389a0a / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-27cc87332cf786bb256085f80a133ae11c8d245a7dc594446dd0755c398b3f20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ddc669e76d8d2cb052ea57af17cae486e3544e650ec44aef0033697fdcc6bca"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control / 26bfb1371829 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d9194f080806e69f7a565ead1caaf1ae7df6c377a284846942501e5419ac7a8a)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

<a id="canonical-fb254f9a0e948dddc74bfba9ea69d64602d5d108b377468021eb6ee058bf8139"></a>

Type: `"single"`. Computed.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

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

<a id="canonical-11fa2c981906fae11d55c84df6190f1cbe311f0cd7846033db4feb19e9c38d04"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control / 26bfb1371829 / 3

- [exclude_attack_type_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-95253544a0b81bca776781ac818f03915de4b50f97448abb7b62fd9676948ffd): complete subsection reference.

- [exclude_bot_name_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-356f542cd2175fa09310965df86c5ca18d9df6fd719bc16950730af0e1c46e25): complete subsection reference.

- [exclude_signature_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-db09047bd07094a7bebf89ca584aab70867a85eb3505c6acd8b69112f01f6f16): complete subsection reference.

- [exclude_violation_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-85a94ca8b3ad13a3f6ab6078546002e4283a697e10003c6941bd331f279ed375): complete subsection reference.

<a id="canonical-742ad062fd9f640586d44d90cf47f3d3fb495921d815846c47657cb0647cfb19"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control / 26bfb1371829 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-95253544a0b81bca776781ac818f03915de4b50f97448abb7b62fd9676948ffd)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-356f542cd2175fa09310965df86c5ca18d9df6fd719bc16950730af0e1c46e25)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-db09047bd07094a7bebf89ca584aab70867a85eb3505c6acd8b69112f01f6f16)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-85a94ca8b3ad13a3f6ab6078546002e4283a697e10003c6941bd331f279ed375)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-95253544a0b81bca776781ac818f03915de4b50f97448abb7b62fd9676948ffd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5240dcfb163cf071bb292935369f30963bab279f0e4de44de7b1a81717a5618"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / e095570537cd / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d9194f080806e69f7a565ead1caaf1ae7df6c377a284846942501e5419ac7a8a)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-27cc87332cf786bb256085f80a133ae11c8d245a7dc594446dd0755c398b3f20)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-3a2f361827697d6e9f2f4f0faa2710ba2d98ec510f01203a5ffbb42f52ea536f"></a>

Type: `"list"`. Computed.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-397c3a977dd8054bed77dc211415ba18a9ccea7c042664d44049e225dcc6b659"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / e095570537cd / 3

<a id="canonical-091f33f62a279481157690420ca42d0236d8c9e09b6857b678a84f3ac006f7e1"></a>

<a id="canonical-a75225d4c26ec78b804229afaa1156fb9d5e42a0e4186fd7c0d5d6bed28534b4"></a>

## context property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / e095570537cd / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6517821138a744f4d2abdc30b5a59159e0f49f0ff7e742b446f04ff9d48ee21f"></a>

<a id="canonical-449794ecf6c22c403bc6bc0602f4dde318db050bdca53823319183ad62f03157"></a>

## context_name property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / e095570537cd / 5

Type: `"string"`. Computed.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-bb245224e93bed82b4db7a821cd0baddd8ac6d05dc5a184f670041e986420a05"></a>

<a id="canonical-79afd4bd3b6255582ae96e51360994a55e419f39ab7cfb3a3257e40100a1cf14"></a>

## exclude_attack_type property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / e095570537cd / 6

Type: `"string"`. Computed.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Upstream description:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-39c8b37a1eac924b7f573936b5ceb5098e8d623b1d00ccadd5cd5164e86b9e59"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / e095570537cd / 7

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-27cc87332cf786bb256085f80a133ae11c8d245a7dc594446dd0755c398b3f20)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-356f542cd2175fa09310965df86c5ca18d9df6fd719bc16950730af0e1c46e25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32d5a9fac3632ce94bb510907cb8ced3392e323dabf129a86f5e5f07abee9953"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / d573338c8574 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d9194f080806e69f7a565ead1caaf1ae7df6c377a284846942501e5419ac7a8a)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-27cc87332cf786bb256085f80a133ae11c8d245a7dc594446dd0755c398b3f20)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-c89e1e9b6df91aaba200919d9160f95c419d9ac472fca25b7da2e6be4aba6ab5"></a>

Type: `"list"`. Computed.

Bot Names to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8af3314c13a3186e27c405c49cb6bb8795a99f8aa91a08870521062ba856e3c4"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / d573338c8574 / 3

<a id="canonical-268e075b5297542ff7a3bd84ea465dcc565232f7da536b90b282c8ea98386128"></a>

<a id="canonical-6a95bd8aa22a4b854bb68a29eccd7331d57996b4d0c0ba810b1a3fecc91b8c62"></a>

## bot_name property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / d573338c8574 / 4

Type: `"string"`. Computed.

Bot Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-a5292cb6425ed72a818f75fac983da6703f25c1cda424397fa9b56cbc8251b44"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / d573338c8574 / 5

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-27cc87332cf786bb256085f80a133ae11c8d245a7dc594446dd0755c398b3f20)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-db09047bd07094a7bebf89ca584aab70867a85eb3505c6acd8b69112f01f6f16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d7258798aeb7d0fc65ffd9a249cb89abfe166db2509d12f5ae2f7d33d8d0200"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 7af78837e3c4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d9194f080806e69f7a565ead1caaf1ae7df6c377a284846942501e5419ac7a8a)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-27cc87332cf786bb256085f80a133ae11c8d245a7dc594446dd0755c398b3f20)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-04929cf1ae3574bab2acdc7fb4244b98f127f7833923968f397faddd59c60780"></a>

Type: `"list"`. Computed.

Signature IDs to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-86eb26c71ffd15083bc003abe5f94957f853a0406a0c24126437b9aa301a7a5e"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 7af78837e3c4 / 3

<a id="canonical-16297596a5a4de1eabecaca5b68ae71e9290c54955538becd1049b5817f98854"></a>

<a id="canonical-429ff416ab50bc2c0f3afb8e132d7a3d1377fa51a40b1464eb69ad8b8e9f4fce"></a>

## context property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 7af78837e3c4 / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7dee601907f0da0fbc141bcec4974f854ff18ff324edff0e381bcacaef2cd529"></a>

<a id="canonical-071424520efb2fb87c52aed40246cc5a9a4499eddd83a371ada99b98fbfdcf56"></a>

## context_name property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 7af78837e3c4 / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-118341867c56f8f5c232014940a3fc144530ce86d46b6f67ee63aa8d159c3ecd"></a>

<a id="canonical-8a991d33bd097c169a90bdc3b5869e5e27e5d6f53d5fdd56ac39e68fccad853b"></a>

## signature_id property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 7af78837e3c4 / 6

Type: `"number"`. Computed.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

<a id="canonical-d6c7b04fdbf60aae43b5d6d304af58ca53ec5ef9e77a7f88355cc97279225dd8"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 7af78837e3c4 / 7

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-27cc87332cf786bb256085f80a133ae11c8d245a7dc594446dd0755c398b3f20)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-85a94ca8b3ad13a3f6ab6078546002e4283a697e10003c6941bd331f279ed375"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ebad44ca71a9c059403f6cf0904f1cb493e01aa2fab110018894a5b25abdc80"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 23ffe65fee4c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d9194f080806e69f7a565ead1caaf1ae7df6c377a284846942501e5419ac7a8a)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-27cc87332cf786bb256085f80a133ae11c8d245a7dc594446dd0755c398b3f20)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-0fd555a689022073f978d841db37701a95fdc1d0a1c79d778a26d21ecd697a7b"></a>

Type: `"list"`. Computed.

Violations to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c8ed847c18baa55243ecd993e8d813327bd1844007e19eb3434806bf623c656a"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 23ffe65fee4c / 3

<a id="canonical-7d71dd2f49e0f51710ac63b3bcbafb2c501c2be816f983f500dd38b2b837174e"></a>

<a id="canonical-d94b6c210eddf7cceec95619b1c5c474d36ce5ba50caabf9162b82b951d67dd1"></a>

## context property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 23ffe65fee4c / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-813508429925bb86aa37d318134a212f414ecdc8289d0fa1be3101fbb55099a3"></a>

<a id="canonical-e6d943aae1afcfe972f583e64982097afcf29fd244c2d769bb6ad1c8b8f30bb8"></a>

## context_name property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 23ffe65fee4c / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-51345d56a32a5e35fa95fb261a62ff59b65aa773891e27701ea12daa5cb1102d"></a>

<a id="canonical-09b576c7d00a5988b4dd4818e4123753cde182c1ce647a3edfcc61e9f2a2aa2a"></a>

## exclude_violation property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 23ffe65fee4c / 6

Type: `"string"`. Computed.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
List of all supported Violation Types VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER
VIOL\_HTTP\_RESPONSE\_STATUS VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD
VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED
VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS.. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Upstream description:

List of all supported Violation Types

VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER VIOL\_HTTP\_RESPONSE\_STATUS
VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED
VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS
VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT
VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION
VIOL\_HTTP\_PROTOCOL\_CRLF\_CHARACTERS\_BEFORE\_REQUEST\_START
VIOL\_HTTP\_PROTOCOL\_NO\_HOST\_HEADER\_IN\_HTTP\_1\_1\_REQUEST
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_PARAMETERS\_PARSING
VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS
VIOL\_HTTP\_PROTOCOL\_CONTENT\_LENGTH\_SHOULD\_BE\_A\_POSITIVE\_NUMBER
VIOL\_EVASION\_DIRECTORY\_TRAVERSALS VIOL\_MALFORMED\_REQUEST VIOL\_EVASION\_MULTIPLE\_DECODING
VIOL\_DATA\_GUARD VIOL\_EVASION\_APACHE\_WHITESPACE VIOL\_COOKIE\_MODIFIED
VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS VIOL\_EVASION\_IIS\_BACKSLASHES
VIOL\_EVASION\_PERCENT\_U\_DECODING VIOL\_EVASION\_BARE\_BYTE\_DECODING VIOL\_EVASION\_BAD\_UNESCAPE
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_FORMDATA\_REQUEST\_PARSING
VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST
VIOL\_HTTP\_PROTOCOL\_HIGH\_ASCII\_CHARACTERS\_IN\_HEADERS VIOL\_ENCODING VIOL\_COOKIE\_MALFORMED
VIOL\_GRAPHQL\_FORMAT VIOL\_GRAPHQL\_MALFORMED VIOL\_GRAPHQL\_INTROSPECTION\_QUERY.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIOL_NONE",
  "enum": [
    "VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START",
    "VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6108350f6c5d58c04a9216853675eedc9fa64a5c921aca1e235deab68dc525e7"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 23ffe65fee4c / 7

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-27cc87332cf786bb256085f80a133ae11c8d245a7dc594446dd0755c398b3f20)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-5e0b47d889e695190916ac1a54e116c2da17e3c55f834c19bfb1b13c710207da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3be93847bc65c1f9e91e9f1a0d19f8dbb98c22cfae775a67f2661a2c0072767f"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.metadata — waf_exclusion.waf_exclusion_inline_rules.rules.metadata / 980986ffdb6e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d9194f080806e69f7a565ead1caaf1ae7df6c377a284846942501e5419ac7a8a)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- waf_exclusion.waf_exclusion_inline_rules.rules.metadata

<a id="canonical-56dd374696454d893920835e6054078d6486a66e09353a79e00e1bb309489a8e"></a>

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

<a id="canonical-0edd73449e6065561562db7a87d87f37d6408fc984be3d18b7b0d86cfecc48e8"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.metadata / 980986ffdb6e / 3

<a id="canonical-1f91967dc7724a73563812f4b33e8435366f1815668744e035f969c9e89fe318"></a>

<a id="canonical-2017659e32179520006eb9d3b77d000b08aa4968e7f45b0cc6aaf07da54a1088"></a>

## description_spec property — waf_exclusion.waf_exclusion_inline_rules.rules.metadata / 980986ffdb6e / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-6bf45640d562a2895c35a4662a27bd53ffeb73e5873b9879de48e3d9ffdc6ef1"></a>

<a id="canonical-7e7601e5e2511ddd46f3c87dd515b0f4a87361449f2d8054ceb69e9e3fd0edd1"></a>

## name property — waf_exclusion.waf_exclusion_inline_rules.rules.metadata / 980986ffdb6e / 5

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

<a id="canonical-f794edb59abc95d732434c4a0c704bdfb98717ae39fe99c05d7e296f49f084e2"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.metadata / 980986ffdb6e / 6

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-4e12755d168e87df4b2231d6c1c9afd6fe31e031fca8fcecf7a799574ba003da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-259659139dc9fea6959e8bd417759049d5951075c0fb9d9e3a92452a97d56c51"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing — waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing / 51ce228835a8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-d9194f080806e69f7a565ead1caaf1ae7df6c377a284846942501e5419ac7a8a)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing

<a id="canonical-e4427cfef736acfef4ae34364bb0c9dae48b885dacf79b56becdd5b78708c8c9"></a>

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

<a id="canonical-8ea73a11e3271d06b5e0506bd4af04d3ba86e276d17025449657c3bd564ef21c"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing / 51ce228835a8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6d8dc60acc1e132c33a8f7a924021b96295536866d07be9a1e1ec8d58aa73f1e"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing / 51ce228835a8 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-5c67a9b12f8bdd1a2b02acbac0dec6a93586fa33ba6e3810b574eb109dbe7131)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3912975fd6ba1b1c42cd8ac7d5f5c825aaff722b2985b40c803582e43cebf9bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3611628d0edcb2e5730081adb402bbdc73ebecb4e991a7d7cded3b0b3e3829fa"></a>

## waf_exclusion.waf_exclusion_policy — waf_exclusion.waf_exclusion_policy / 556485f7d5b2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37)
- waf_exclusion.waf_exclusion_policy

<a id="canonical-99e700426789db729da39072dc681a209ac6bd85ca715bd4f3538404f10989b5"></a>

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

<a id="canonical-e9b8453b7d441eb45cdd7ec3570a1c91f50773550d8d5f18a017194a56c8a59d"></a>

## Direct properties — waf_exclusion.waf_exclusion_policy / 556485f7d5b2 / 3

<a id="canonical-fd327d046d3241060d18fe454ce118b7c11b3fb0cf1d0c05207716a9a472c2c3"></a>

<a id="canonical-762848ed65ab4363451669484d9137ac9bf7356a1eee8e8e6b029da6172ba5f1"></a>

## name property — waf_exclusion.waf_exclusion_policy / 556485f7d5b2 / 4

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

<a id="canonical-7ea442762abaf34f1d6f9d8e9accaa54dd19b6e5a54cc57e52d718efad36cf66"></a>

<a id="canonical-e1bb9fa68460d873f593336fe80d67bb15d6ffdd3a5926dcedd22e6b5c8a8ad4"></a>

## namespace property — waf_exclusion.waf_exclusion_policy / 556485f7d5b2 / 5

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

<a id="canonical-dd73c426a2d436380052b601ea29ce189a4a6cf49ecc718e4507992b3759081f"></a>

<a id="canonical-59b511f649cf80512b11fdb535e9f0440bdf1f0e783be322fb9560927374997e"></a>

## tenant property — waf_exclusion.waf_exclusion_policy / 556485f7d5b2 / 6

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

<a id="canonical-94d8ee5f474a2a506d1ae191f6a51123bbd08acdeb9888d3bc6dbb0c881a3bb8"></a>

## Next pages — waf_exclusion.waf_exclusion_policy / 556485f7d5b2 / 7

- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
