---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-ef4d7d9cd21ab8172b9eb8e85e5f53be755752072adc3e321b3183597bde1726"></a>

## Direct properties — cloudfront.protected_endpoints.domain / 31d3eb088934 / 3

<a id="canonical-ae7ff992545d343aebe6778094b2c91057c9f12a6bcdc7c19e167ac5c56950fc"></a>

<a id="canonical-a3ded592b76ebebef0ccc6ee217b26e93eca51ab15183c2b44e1d29ebed112e3"></a>

## exact_value property — cloudfront.protected_endpoints.domain / 31d3eb088934 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-922007086cd6ab4393ae21d589d7cdd9cf251e9c2c2c5d4c955f68be808f6ef8"></a>

<a id="canonical-f90a1dde138f4bbf8311af180d95f2b33ed5ad90ce422f6e106bcc19cb2a5930"></a>

## regex_value property — cloudfront.protected_endpoints.domain / 31d3eb088934 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-d0f631e795bdbe877b66ce34dc5e342ace8c9ded8cf212cbd8b57182273dca53"></a>

<a id="canonical-105154844126f7af7b10caf4134c414ccb9f8bdfe127f6a7cf87d3c634f4acc0"></a>

## suffix_value property — cloudfront.protected_endpoints.domain / 31d3eb088934 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
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

<a id="canonical-84b79b5a095a56c6a732542d4490af8ffcbaa0263a6c242eab865ecc7154fc6d"></a>

## Next pages — cloudfront.protected_endpoints.domain / 31d3eb088934 / 7

- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31becdfdec1d2ea72416da26281e302cceb3287e80669328ca2f2707f3129207"></a>

## cloudfront.protected_endpoints.flow_label — cloudfront.protected_endpoints.flow_label / 1bca4d1fd1f0 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- cloudfront.protected_endpoints.flow_label

<a id="canonical-883f954fa2a9f6c89d7e69fb2a41235f8f386662b78ba726f427dd4ac7aa4023"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

<a id="canonical-962c9f0b03742ffd3ceea2e44eaae4927cad0e8be5068b3102503997f6a527a5"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label / 1bca4d1fd1f0 / 3

- [account_management](data-sources--protected_application--reference--group-003.md#canonical-6102c482cecc1a6845d4f5611a6789ae74790bf52d852adfcdbc9b9bc109a38d): complete subsection reference.

- [authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1): complete subsection reference.

- [financial_services](data-sources--protected_application--reference--group-003.md#canonical-05a2f5487c193ebda2588693335adc40845824c27ef68cdb33c788bd8c6fe323): complete subsection reference.

- [flight](data-sources--protected_application--reference--group-003.md#canonical-ca374c077336668351a167db8d5a7361b5f794fcaca7f9ec273571e82a17e1d3): complete subsection reference.

- [profile_management](data-sources--protected_application--reference--group-003.md#canonical-91dddcc74599be6540edb5b70f4afdb7e2afe56fb1d0e11189deea45813874f3): complete subsection reference.

- [search](data-sources--protected_application--reference--group-003.md#canonical-77ac0bb1a507ea767b5b75f197887be654ecde8ed1b84838d54e3fc8f5dc5903): complete subsection reference.

- [shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939): complete subsection reference.

<a id="canonical-7cf4babe1e4b970397fa8e92af9c00292f3e73ec01fe759b79bc6014786811cc"></a>

## Next pages — cloudfront.protected_endpoints.flow_label / 1bca4d1fd1f0 / 4

- [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--reference--group-003.md#canonical-6102c482cecc1a6845d4f5611a6789ae74790bf52d852adfcdbc9b9bc109a38d)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--reference--group-003.md#canonical-05a2f5487c193ebda2588693335adc40845824c27ef68cdb33c788bd8c6fe323)
- [cloudfront.protected_endpoints.flow_label.flight](data-sources--protected_application--reference--group-003.md#canonical-ca374c077336668351a167db8d5a7361b5f794fcaca7f9ec273571e82a17e1d3)
- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-91dddcc74599be6540edb5b70f4afdb7e2afe56fb1d0e11189deea45813874f3)
- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-77ac0bb1a507ea767b5b75f197887be654ecde8ed1b84838d54e3fc8f5dc5903)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-6102c482cecc1a6845d4f5611a6789ae74790bf52d852adfcdbc9b9bc109a38d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e088d04dfe8c995d26d0ec860e71e251ab3590183618ea9fd86ab279b4b5898"></a>

## cloudfront.protected_endpoints.flow_label.account_management — cloudfront.protected_endpoints.flow_label.account_management / 67e199175f51 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- cloudfront.protected_endpoints.flow_label.account_management

<a id="canonical-e0141c194617023468bc7c89e1ab3c76e503a7fde1d813686ac1118ba8ee9854"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Account Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

<a id="canonical-1b856ae07c4256474b4359ec321a7fe554a0fccf1bcbce6cd03408c79de46b91"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.account_management / 67e199175f51 / 3

- [create](data-sources--protected_application--reference--group-003.md#canonical-e453018f45c25b365dd4e97f02831ca1702b06965359c9228e50fe9ffdbcc2d1): complete subsection reference.

- [password_reset](data-sources--protected_application--reference--group-003.md#canonical-11d0ca56aae467da60cceeb2a6b00c573613849b5ce1f8bf98d279df1a9d9fc4): complete subsection reference.

<a id="canonical-040dfb568595ebad7a35d3e7a79f8bcb2260267e2fdc17a6d4400322da5afe30"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.account_management / 67e199175f51 / 4

- [cloudfront.protected_endpoints.flow_label.account_management.create](data-sources--protected_application--reference--group-003.md#canonical-e453018f45c25b365dd4e97f02831ca1702b06965359c9228e50fe9ffdbcc2d1)
- [cloudfront.protected_endpoints.flow_label.account_management.password_reset](data-sources--protected_application--reference--group-003.md#canonical-11d0ca56aae467da60cceeb2a6b00c573613849b5ce1f8bf98d279df1a9d9fc4)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-e453018f45c25b365dd4e97f02831ca1702b06965359c9228e50fe9ffdbcc2d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43174a6b74eada40f6f654c5ff54092757ad35b3e59cc8d17873c8e4c2b2fcc5"></a>

## cloudfront.protected_endpoints.flow_label.account_management.create — cloudfront.protected_endpoints.flow_label.account_management.create / 6631112b4f16 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--reference--group-003.md#canonical-6102c482cecc1a6845d4f5611a6789ae74790bf52d852adfcdbc9b9bc109a38d)
- cloudfront.protected_endpoints.flow_label.account_management.create

<a id="canonical-6197666565d2ca11a8a7398abd5576a568cada100b933657a64c7a8fe2459d18"></a>

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

<a id="canonical-718fd5ae4f1d9abd88bcbf03d7c679c79d93ce36a96d6062f81b6e8c6b569589"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.account_management.create / 6631112b4f16 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-98bf54ec1dc993bc5c57f8d2454e91f540b8e1b987bc7b67d80344ffa13c4657"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.account_management.create / 6631112b4f16 / 4

- [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--reference--group-003.md#canonical-6102c482cecc1a6845d4f5611a6789ae74790bf52d852adfcdbc9b9bc109a38d)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-11d0ca56aae467da60cceeb2a6b00c573613849b5ce1f8bf98d279df1a9d9fc4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebc6f98c3feed061902ecd4c073fc81ec8e72793e1b64aebcc47cec153f91ef4"></a>

## cloudfront.protected_endpoints.flow_label.account_management.password_reset — cloudfront.protected_endpoints.flow_label.account_management.password_reset / ce6e52d824c7 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--reference--group-003.md#canonical-6102c482cecc1a6845d4f5611a6789ae74790bf52d852adfcdbc9b9bc109a38d)
- cloudfront.protected_endpoints.flow_label.account_management.password_reset

<a id="canonical-c908cd8075655561a494b1c9df3d2db2d7765dab3b2c3418a2e0e9eb4022d72e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for password reset.

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

<a id="canonical-279b2db7ae8409dceac734552a466ebc2eb422523aa63589592672560b44046b"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.account_management.password_reset / ce6e52d824c7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07a4c5f6b0b3f762e50334b8493901504ad46317db4391f1fe5298bd9a2d0666"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.account_management.password_reset / ce6e52d824c7 / 4

- [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--reference--group-003.md#canonical-6102c482cecc1a6845d4f5611a6789ae74790bf52d852adfcdbc9b9bc109a38d)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-815474d0aea7e0f37dbbba93327a000737c3768fbd8f8dca67777e099193a85d"></a>

## cloudfront.protected_endpoints.flow_label.authentication — cloudfront.protected_endpoints.flow_label.authentication / a4f46afd481a / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- cloudfront.protected_endpoints.flow_label.authentication

<a id="canonical-1f27742923d590ed1bd11ed4cfffe6ef8937b5be1eb8d0939377d3100824f07c"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Authentication Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

<a id="canonical-6d84bbcb35a516c9cad0d9c609703c5942438e8460298e7089da5bf616a1d2d9"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication / a4f46afd481a / 3

- [login](data-sources--protected_application--reference--group-003.md#canonical-c27d5a966d38259ae577bb0cc7a8f449041fb7a42d25c5819be546cdfb9450ea): complete subsection reference.

- [login_mfa](data-sources--protected_application--reference--group-003.md#canonical-28e7d6d342c7010bac382ac1a978eae28b40eaabb13e93dbdc443e1c438619d0): complete subsection reference.

- [login_partner](data-sources--protected_application--reference--group-003.md#canonical-cd2412f441ccbcd97fa2a036e0d6d3a5cb5d67a1f4cf3bc2823fd5267de5b311): complete subsection reference.

- [logout](data-sources--protected_application--reference--group-003.md#canonical-4699b8cefe276a78b92eab5f11fb754eb4192e5a342faf9073d1d9ba1dd977ac): complete subsection reference.

- [token_refresh](data-sources--protected_application--reference--group-003.md#canonical-aa163b767e884c2cb4a226a1ed92024045df1a11d056b49134d9528b7c178c66): complete subsection reference.

<a id="canonical-0c80ff99ab3def487675718f3d6afdbd6804140e9f37bc9f3a728ea5279afdd5"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication / a4f46afd481a / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-c27d5a966d38259ae577bb0cc7a8f449041fb7a42d25c5819be546cdfb9450ea)
- [cloudfront.protected_endpoints.flow_label.authentication.login_mfa](data-sources--protected_application--reference--group-003.md#canonical-28e7d6d342c7010bac382ac1a978eae28b40eaabb13e93dbdc443e1c438619d0)
- [cloudfront.protected_endpoints.flow_label.authentication.login_partner](data-sources--protected_application--reference--group-003.md#canonical-cd2412f441ccbcd97fa2a036e0d6d3a5cb5d67a1f4cf3bc2823fd5267de5b311)
- [cloudfront.protected_endpoints.flow_label.authentication.logout](data-sources--protected_application--reference--group-003.md#canonical-4699b8cefe276a78b92eab5f11fb754eb4192e5a342faf9073d1d9ba1dd977ac)
- [cloudfront.protected_endpoints.flow_label.authentication.token_refresh](data-sources--protected_application--reference--group-003.md#canonical-aa163b767e884c2cb4a226a1ed92024045df1a11d056b49134d9528b7c178c66)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-c27d5a966d38259ae577bb0cc7a8f449041fb7a42d25c5819be546cdfb9450ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b20e2f586a8324cb5aa7ba63918f75652deb1e50fbce683dfe2a3bf39853bf00"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login — cloudfront.protected_endpoints.flow_label.authentication.login / 91a0d3a8329c / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- cloudfront.protected_endpoints.flow_label.authentication.login

<a id="canonical-464df341b620fbed08e03dc2c827ad2822acc63f219aba91269467ea8a5fcd1a"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Upstream description:

Bot Defense Transaction Result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

<a id="canonical-cca747ac8d37f99f482c5ab5eb0ab8c8d9b0e629a9d2300a9cedc0d9db7342d8"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.login / 91a0d3a8329c / 3

- [disable_transaction_result](data-sources--protected_application--reference--group-003.md#canonical-dabf1d7931e193b3708ea2556eccb4659e575da0e864eb93a9c5fefb132a88e4): complete subsection reference.

- [transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0722f3f4f21f834f73d80bf267c81e75e9392067ea9ba7eb024ded3c2e04460a): complete subsection reference.

<a id="canonical-b2d5eba48871e2a2dab62c90949860cd3ca05729bde222a30861ba8da070b581"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.login / 91a0d3a8329c / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result](data-sources--protected_application--reference--group-003.md#canonical-dabf1d7931e193b3708ea2556eccb4659e575da0e864eb93a9c5fefb132a88e4)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0722f3f4f21f834f73d80bf267c81e75e9392067ea9ba7eb024ded3c2e04460a)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-dabf1d7931e193b3708ea2556eccb4659e575da0e864eb93a9c5fefb132a88e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98b2dee520146d0c685217343d6ed03399d062e36bf3daf0530053948847841b"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result — cloudfront.protected_endpoints.flow_label.authentication.login.disable_transacti / c5bb628d8a4a / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-c27d5a966d38259ae577bb0cc7a8f449041fb7a42d25c5819be546cdfb9450ea)
- cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-1a8250c21c9baef1526bd57f4a6f5d82ff4362ce994b1dda948fd62f2022cbad"></a>

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

<a id="canonical-e5c84d1aa49fd701bbcafbb4c6e7063ed8d8984b4086743665d9521eec332720"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.login.disable_transacti / c5bb628d8a4a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9c9b0fb58686609c28f2ee63db4da2cf047f88aa83d3a01fceea683a85f995b0"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.login.disable_transacti / c5bb628d8a4a / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-c27d5a966d38259ae577bb0cc7a8f449041fb7a42d25c5819be546cdfb9450ea)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-0722f3f4f21f834f73d80bf267c81e75e9392067ea9ba7eb024ded3c2e04460a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28020a0572270cff19260c52789e82fff97d555b1ce27160521760684d50bca1"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / c9d396fcd3fd / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-c27d5a966d38259ae577bb0cc7a8f449041fb7a42d25c5819be546cdfb9450ea)
- cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-a654a7bea255a8d5d1b7b3f8480abe73d554969ff3d31f47c75916c6b4009e9e"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

Upstream description:

Bot Defense Transaction ResultType.

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

<a id="canonical-8c90b2d7a7a462a424018eb364f3069e2330b3423b1be08b4184bc84a908fdee"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / c9d396fcd3fd / 3

- [failure_conditions](data-sources--protected_application--reference--group-003.md#canonical-dc429767ace7bfc838d5e6fa9dfe3383b3dcd2b76d5bbb8b0c664e03ae012fa3): complete subsection reference.

- [success_conditions](data-sources--protected_application--reference--group-003.md#canonical-ec08280b71a7ed9316a0a65456ae4d81e8df7bd3de3d8ba823e104c2dae66399): complete subsection reference.

<a id="canonical-01190bc6d07f469a8e0c44ffd9c430836055fb8ca1b321a916f4746a935eeab5"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / c9d396fcd3fd / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](data-sources--protected_application--reference--group-003.md#canonical-dc429767ace7bfc838d5e6fa9dfe3383b3dcd2b76d5bbb8b0c664e03ae012fa3)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions](data-sources--protected_application--reference--group-003.md#canonical-ec08280b71a7ed9316a0a65456ae4d81e8df7bd3de3d8ba823e104c2dae66399)
- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-c27d5a966d38259ae577bb0cc7a8f449041fb7a42d25c5819be546cdfb9450ea)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-dc429767ace7bfc838d5e6fa9dfe3383b3dcd2b76d5bbb8b0c664e03ae012fa3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08d0304e75a315274a41d7d15fb28074c4ef50fc5fdea69ae86a0c02d960dbce"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 6f1803c54d8f / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-c27d5a966d38259ae577bb0cc7a8f449041fb7a42d25c5819be546cdfb9450ea)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0722f3f4f21f834f73d80bf267c81e75e9392067ea9ba7eb024ded3c2e04460a)
- cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-f4a2f2b198bbf3d5ab9d76ad1fe1ca2cfffda35b6ac239e78a1c1ae52743c815"></a>

Type: `"list"`. Computed.

Failure Conditions. Failure Conditions.

Upstream description:

Failure Conditions.

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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4ebd2d834a9ea4f95cb5307fbddbc827097150319a7905c1a90f06b952b0d088"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 6f1803c54d8f / 3

<a id="canonical-38b0d0b3aa7ac998b832774b782aaa8bb18abc35a892d529c42cc63ad65a4762"></a>

<a id="canonical-146ea520a7a513bd7a1e5e6aab5b1b7434ce2f19fdef5168b4b7187e14582b46"></a>

## name property — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 6f1803c54d8f / 4

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-638c64ee4cabac18e0bd322b988edbfbdb8a90fb2eda0729a485ab1437dc140c"></a>

<a id="canonical-08b277dd243c99fadbbaa658c2af5fb21ca4db310f9907dd00ddc957fc6b8cc1"></a>

## regex_values property — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 6f1803c54d8f / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-abd850b442e540ecf7cf1b78434a2dd6f7734e99f3af7cf33dcacbc6b517da98"></a>

<a id="canonical-3044f981f3411b88cda65e969275898583b9f67d836334204d220b8a8c8b9d19"></a>

## status property — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 6f1803c54d8f / 6

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-75463086e9704aa323ca94a429e45e5356cc9995bc318458051855e67d9cdcae"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 6f1803c54d8f / 7

- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0722f3f4f21f834f73d80bf267c81e75e9392067ea9ba7eb024ded3c2e04460a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-ec08280b71a7ed9316a0a65456ae4d81e8df7bd3de3d8ba823e104c2dae66399"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f07dd36a0ae91276eec58d47be0a65ac5ceaabcec565c2ebe205325db3eb048"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 158bf2869449 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-c27d5a966d38259ae577bb0cc7a8f449041fb7a42d25c5819be546cdfb9450ea)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0722f3f4f21f834f73d80bf267c81e75e9392067ea9ba7eb024ded3c2e04460a)
- cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-c6a4fb6019cb14454f6be6cb051481590c31ef833c85e8b682f0cb5db7dae632"></a>

Type: `"list"`. Computed.

Success Conditions. Success Conditions.

Upstream description:

Success Conditions.

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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4fd3a9c0595b0566e1b13ad702cbe0ce80769f07dbfdf1b5833f64e0f7338067"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 158bf2869449 / 3

<a id="canonical-c3683e548dbc68a3c904b28e30a3c0138a85bcd469b926c4bd9ff247e1687cf5"></a>

<a id="canonical-f3573c340c3a2a1e4151a6e455d493ae9193a54e220c0b883faf9c4441cd32f4"></a>

## name property — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 158bf2869449 / 4

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-059ffdd8aca654751c5659d3fd2be063ba25a5719da4758e6592b7a9b93fff78"></a>

<a id="canonical-4ddc653c79b282132d046d552448e52b6cb3e8ce4bcdcf6870b3e15c0cfdbffe"></a>

## regex_values property — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 158bf2869449 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-83980c7397239f6665d78c52beccc64a81c96bf099233840a13c64893b7c5fae"></a>

<a id="canonical-4548f6293a7cad32501f3799aac00f458b3c8bd3f1127df1c4ee3c049637484a"></a>

## status property — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 158bf2869449 / 6

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4a2d803d5610bb427fea81f359a9acebedfaf5253e7b10c0d6490a758dee4e28"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 158bf2869449 / 7

- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0722f3f4f21f834f73d80bf267c81e75e9392067ea9ba7eb024ded3c2e04460a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-28e7d6d342c7010bac382ac1a978eae28b40eaabb13e93dbdc443e1c438619d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02f59bf64e913c1fcf03ee599ad9a89786f24469681c735fc151cb6fc2d4e604"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login_mfa — cloudfront.protected_endpoints.flow_label.authentication.login_mfa / d02a05aaedc0 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- cloudfront.protected_endpoints.flow_label.authentication.login_mfa

<a id="canonical-e83d761391be4ce942bb16f6392dbbbb1919473f2b70eeebac2d68aeea3fe97d"></a>

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

<a id="canonical-58029a96a723486cdffb8586418781da238f1d25a99927bb4175f21818a58a81"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.login_mfa / d02a05aaedc0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c6b729d5f99fc92e0081a8ccf41e15b0db5e15b355da7c0b11864281613c9e42"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.login_mfa / d02a05aaedc0 / 4

- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-cd2412f441ccbcd97fa2a036e0d6d3a5cb5d67a1f4cf3bc2823fd5267de5b311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d92c9de77d52ed99418b81c55e24bc5f581078e1848cb5cf0621266641607812"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login_partner — cloudfront.protected_endpoints.flow_label.authentication.login_partner / 45cc32af96f3 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- cloudfront.protected_endpoints.flow_label.authentication.login_partner

<a id="canonical-322c6bc7e900b8048d13c22e42e2890b4381c78b5d79525897f0e25438f6ee41"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for login partner.

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

<a id="canonical-57051a78555d493018fd36783440989e0e5db01f125c34fddac43491cdf25ec9"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.login_partner / 45cc32af96f3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6859110319cd7a18d7f53ea76e789c7afd66e9fb770380ca00086b14c588ccfc"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.login_partner / 45cc32af96f3 / 4

- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-4699b8cefe276a78b92eab5f11fb754eb4192e5a342faf9073d1d9ba1dd977ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27fdfdcca3d3ecc1e5ccb4880886b052f447de57ea6b01e7f5259581f6abc300"></a>

## cloudfront.protected_endpoints.flow_label.authentication.logout — cloudfront.protected_endpoints.flow_label.authentication.logout / 12690eafd58e / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- cloudfront.protected_endpoints.flow_label.authentication.logout

<a id="canonical-42dc2d513d755c7b1fbb70f6aad2a45b0b4738e51f4c71127c142eec52565cf1"></a>

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

<a id="canonical-936971078c7cafa92dcfdaa16976e9f395501bef6cd9ad7d163b724ab7a03423"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.logout / 12690eafd58e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0ad9fffa5860f6adc48b3d16fa0d3c3629f318e240b0cd83f8a4e35289d04841"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.logout / 12690eafd58e / 4

- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-aa163b767e884c2cb4a226a1ed92024045df1a11d056b49134d9528b7c178c66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a564999f08ceb173f87945363949ae5fcc165677bb178ba1ceaa34c139724ed"></a>

## cloudfront.protected_endpoints.flow_label.authentication.token_refresh — cloudfront.protected_endpoints.flow_label.authentication.token_refresh / 7ddb0dbc0d5d / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- cloudfront.protected_endpoints.flow_label.authentication.token_refresh

<a id="canonical-b1b5693effe5abf629f3bf60eff031c09eff0f6af8f6d1d0998fb63d28a30342"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for token refresh.

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

<a id="canonical-7d4816b10f96fab5c5dcc51efea907bd58871e511823ef209c39e6373b952764"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.token_refresh / 7ddb0dbc0d5d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4b18a8b6457cdca1b835dda59109772512beb835d9f10a2d6b34fd5018891318"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.token_refresh / 7ddb0dbc0d5d / 4

- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-794bb27f65352a0086eef0ae79e4c398de87bcadd49f8c06b7300330b57ea9a1)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-05a2f5487c193ebda2588693335adc40845824c27ef68cdb33c788bd8c6fe323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e26c445bffd3b60aa61ce9a3137e7340eda078ca969e54dcd500cd93e0df8c4"></a>

## cloudfront.protected_endpoints.flow_label.financial_services — cloudfront.protected_endpoints.flow_label.financial_services / c423aef0d5a7 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- cloudfront.protected_endpoints.flow_label.financial_services

<a id="canonical-aa2e51b746515ebd61cb823548072da32c7883bd2d2655239f8355b350bcd737"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Financial Services Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"apply\",\"money_transfer\"]"
}
```

<a id="canonical-ed20aaccdf09b5d1afecc500a4082c278cc00e7cd3cc6fe34decf6b3993e3949"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.financial_services / c423aef0d5a7 / 3

- [apply](data-sources--protected_application--reference--group-003.md#canonical-7e330c108171013379c1169735cd6111aef7277970962e1e5109d85a02bb86ca): complete subsection reference.

- [money_transfer](data-sources--protected_application--reference--group-003.md#canonical-aec5d4642be6cc6716eac44eab3f2ff7e93ec8273b4adfdf317a084034f96c2c): complete subsection reference.

<a id="canonical-b9b09d812bee7248509c4f30161e67c838ce2ccb22d2f601a81aca5366d5505f"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.financial_services / c423aef0d5a7 / 4

- [cloudfront.protected_endpoints.flow_label.financial_services.apply](data-sources--protected_application--reference--group-003.md#canonical-7e330c108171013379c1169735cd6111aef7277970962e1e5109d85a02bb86ca)
- [cloudfront.protected_endpoints.flow_label.financial_services.money_transfer](data-sources--protected_application--reference--group-003.md#canonical-aec5d4642be6cc6716eac44eab3f2ff7e93ec8273b4adfdf317a084034f96c2c)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-7e330c108171013379c1169735cd6111aef7277970962e1e5109d85a02bb86ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4fd83fcff416ab851452a905020ab44e06874f1b8446e382d99f17db2e0e719"></a>

## cloudfront.protected_endpoints.flow_label.financial_services.apply — cloudfront.protected_endpoints.flow_label.financial_services.apply / 016423e91730 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--reference--group-003.md#canonical-05a2f5487c193ebda2588693335adc40845824c27ef68cdb33c788bd8c6fe323)
- cloudfront.protected_endpoints.flow_label.financial_services.apply

<a id="canonical-5e489ca71ae4c2da7c25320cf293a7bc3f836c3f06cb124cf3db9f045be01c92"></a>

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

<a id="canonical-c39901f3c1d46d486be6271ba02ea066e8d60806aa9505a77ac7c13d41209198"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.financial_services.apply / 016423e91730 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ca3d011a12dcec00e4acc97fee21c5eed816af45bc00d632602c171bbf7ebfa5"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.financial_services.apply / 016423e91730 / 4

- [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--reference--group-003.md#canonical-05a2f5487c193ebda2588693335adc40845824c27ef68cdb33c788bd8c6fe323)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-aec5d4642be6cc6716eac44eab3f2ff7e93ec8273b4adfdf317a084034f96c2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59480595029a497bac9531ccf172ee96542615286defcc9bd13dc01f97bdd3ac"></a>

## cloudfront.protected_endpoints.flow_label.financial_services.money_transfer — cloudfront.protected_endpoints.flow_label.financial_services.money_transfer / f1d4245d86ea / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--reference--group-003.md#canonical-05a2f5487c193ebda2588693335adc40845824c27ef68cdb33c788bd8c6fe323)
- cloudfront.protected_endpoints.flow_label.financial_services.money_transfer

<a id="canonical-f2875c697c098aa377837b07cc2d8a01028c1a451040980c6542a54295bb9d09"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for money transfer.

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

<a id="canonical-80d36601cb9be6f284205bec4d930ca2a819bb665184d655d83b892f5f84597a"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.financial_services.money_transfer / f1d4245d86ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65748a1110c67e1279c1614ffcb5cd5626af8a11735272d9021ef2d41e44da3e"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.financial_services.money_transfer / f1d4245d86ea / 4

- [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--reference--group-003.md#canonical-05a2f5487c193ebda2588693335adc40845824c27ef68cdb33c788bd8c6fe323)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-ca374c077336668351a167db8d5a7361b5f794fcaca7f9ec273571e82a17e1d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-686e890d52f7e9d3bbd2a8fcca4f27a5b37440470e41f634338c21bb4dc9d999"></a>

## cloudfront.protected_endpoints.flow_label.flight — cloudfront.protected_endpoints.flow_label.flight / c49a7594019c / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- cloudfront.protected_endpoints.flow_label.flight

<a id="canonical-6e6b2bfc27fa95ddb845c400568c53676a66649c3c5c1b68de12471e51d18c2f"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category.

Upstream description:

Bot Defense Flow Label Flight Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"checkin\"]"
}
```

<a id="canonical-6cdf4e198516294cec8241f0eee24f6f3f3b48635df48c493aaebd8d97d36482"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.flight / c49a7594019c / 3

- [checkin](data-sources--protected_application--reference--group-003.md#canonical-cbf1adfcb8e36f3cf60f46ef89f1d015b2cdddb7ea76eec1281676a93f44f987): complete subsection reference.

<a id="canonical-b3cc4020b5636dfb9d531d1cb3891a9b6ba17b8babf57d8bf651fd3e71a43eae"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.flight / c49a7594019c / 4

- [cloudfront.protected_endpoints.flow_label.flight.checkin](data-sources--protected_application--reference--group-003.md#canonical-cbf1adfcb8e36f3cf60f46ef89f1d015b2cdddb7ea76eec1281676a93f44f987)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-cbf1adfcb8e36f3cf60f46ef89f1d015b2cdddb7ea76eec1281676a93f44f987"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a250005cf2562042417cd061342e0976bf98c34404d2728dd8f11c5339b4381"></a>

## cloudfront.protected_endpoints.flow_label.flight.checkin — cloudfront.protected_endpoints.flow_label.flight.checkin / cdbb1e589810 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.flight](data-sources--protected_application--reference--group-003.md#canonical-ca374c077336668351a167db8d5a7361b5f794fcaca7f9ec273571e82a17e1d3)
- cloudfront.protected_endpoints.flow_label.flight.checkin

<a id="canonical-e20c01147f4214772185dd2a0f6c2817ab0a523d844f64f215da94f1d2a05e3a"></a>

Type: `"single"`. Computed.

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

<a id="canonical-c23e49c17eba2ae7c4eff058e07858939af9816f4c3c61131a84ded6f6b32cad"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.flight.checkin / cdbb1e589810 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-21728991cd29e169f44c3f1571459e4e8bd1442bd7876ae47320ba174ba437aa"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.flight.checkin / cdbb1e589810 / 4

- [cloudfront.protected_endpoints.flow_label.flight](data-sources--protected_application--reference--group-003.md#canonical-ca374c077336668351a167db8d5a7361b5f794fcaca7f9ec273571e82a17e1d3)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-91dddcc74599be6540edb5b70f4afdb7e2afe56fb1d0e11189deea45813874f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1da17b38d890cc9db2b877a1fe01f0a8b019290174740da4b2a177190836c786"></a>

## cloudfront.protected_endpoints.flow_label.profile_management — cloudfront.protected_endpoints.flow_label.profile_management / a7846a822460 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- cloudfront.protected_endpoints.flow_label.profile_management

<a id="canonical-b35149c42e05f472ff4405a8188b28fdb034afa5bff913c21d8d1078b8d4216e"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Profile Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"update\",\"view\"]"
}
```

<a id="canonical-ee18b7a56967fedb5bce61814f612cd71ba57dd113d914201f3b00ff0058c704"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.profile_management / a7846a822460 / 3

- [create](data-sources--protected_application--reference--group-003.md#canonical-aec8913b31efb8f5a0f250d4e9c6b4aced5f649a8d6f94576a49792ba7dec579): complete subsection reference.

- [update](data-sources--protected_application--reference--group-003.md#canonical-56666e8f37680d1c7d232582c56a7b90629eccb248712adfbb5f0664712f83af): complete subsection reference.

- [view](data-sources--protected_application--reference--group-003.md#canonical-3d35e03ae63ca434a057f51a671930e75cbc09f5a5e9dc47facd108327f207a9): complete subsection reference.

<a id="canonical-7eca51e7ef55112e2b65827417dfe3870fc132be8b6600b9f0295ec49c869644"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.profile_management / a7846a822460 / 4

- [cloudfront.protected_endpoints.flow_label.profile_management.create](data-sources--protected_application--reference--group-003.md#canonical-aec8913b31efb8f5a0f250d4e9c6b4aced5f649a8d6f94576a49792ba7dec579)
- [cloudfront.protected_endpoints.flow_label.profile_management.update](data-sources--protected_application--reference--group-003.md#canonical-56666e8f37680d1c7d232582c56a7b90629eccb248712adfbb5f0664712f83af)
- [cloudfront.protected_endpoints.flow_label.profile_management.view](data-sources--protected_application--reference--group-003.md#canonical-3d35e03ae63ca434a057f51a671930e75cbc09f5a5e9dc47facd108327f207a9)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-aec8913b31efb8f5a0f250d4e9c6b4aced5f649a8d6f94576a49792ba7dec579"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6ea9fa1994b5221c1a6fd71dc8ac1ae051957ff7bf31ac822953e581aa7af88"></a>

## cloudfront.protected_endpoints.flow_label.profile_management.create — cloudfront.protected_endpoints.flow_label.profile_management.create / 28139e481339 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-91dddcc74599be6540edb5b70f4afdb7e2afe56fb1d0e11189deea45813874f3)
- cloudfront.protected_endpoints.flow_label.profile_management.create

<a id="canonical-cfb910dc1c64a4b7c3be4fdafaeec819cacf5709eb297ae4fedc87b1596458fe"></a>

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

<a id="canonical-c01295555b6a7b5745334d987beda37535c372f615afdb0ea63b86ece747c04e"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.profile_management.create / 28139e481339 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a94c6d5c7fd61cfb824bedfd9f713cb1701be4995100b6f79e98945ba522a400"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.profile_management.create / 28139e481339 / 4

- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-91dddcc74599be6540edb5b70f4afdb7e2afe56fb1d0e11189deea45813874f3)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-56666e8f37680d1c7d232582c56a7b90629eccb248712adfbb5f0664712f83af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0461a45783bb14154af40cbc71c5a18ac23c9e69c76cad363df1bf5c869483e"></a>

## cloudfront.protected_endpoints.flow_label.profile_management.update — cloudfront.protected_endpoints.flow_label.profile_management.update / a41f562dab41 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-91dddcc74599be6540edb5b70f4afdb7e2afe56fb1d0e11189deea45813874f3)
- cloudfront.protected_endpoints.flow_label.profile_management.update

<a id="canonical-8b08b61a39ed216409d3c11e7905c9bf7b54caaf2e9f455b38ee7f449792b0d2"></a>

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

<a id="canonical-e6f73a764c8c5ecd0578127862966b74f071a481ed1d2e84d4c294526b8636f2"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.profile_management.update / a41f562dab41 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-21ad8c395083e899ca3fe7209ef0d5a061079ddab29c7d4ea8591f6d2656b1c6"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.profile_management.update / a41f562dab41 / 4

- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-91dddcc74599be6540edb5b70f4afdb7e2afe56fb1d0e11189deea45813874f3)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-3d35e03ae63ca434a057f51a671930e75cbc09f5a5e9dc47facd108327f207a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-306afa29402478c0aca09fd3fc06fcbd4083971ac5dadebaa312dc0ec2ff0d69"></a>

## cloudfront.protected_endpoints.flow_label.profile_management.view — cloudfront.protected_endpoints.flow_label.profile_management.view / 1f9a2d9fa29d / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-91dddcc74599be6540edb5b70f4afdb7e2afe56fb1d0e11189deea45813874f3)
- cloudfront.protected_endpoints.flow_label.profile_management.view

<a id="canonical-0597fdc9f02ecef57bb4b875d113670f4cb7e953666384d5663e0b81c3ba7093"></a>

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

<a id="canonical-2a8ddb91189648cfc5fe15811f97d096ba97c4c42f4023bda325cfbdad3c7bcf"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.profile_management.view / 1f9a2d9fa29d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-173e5a55007f350c05c55cfae72e40758d7f4669e824727fee8bf39ecc681ba4"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.profile_management.view / 1f9a2d9fa29d / 4

- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-91dddcc74599be6540edb5b70f4afdb7e2afe56fb1d0e11189deea45813874f3)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-77ac0bb1a507ea767b5b75f197887be654ecde8ed1b84838d54e3fc8f5dc5903"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-672230199631a504f7314c40530b1ad648da6edd70aef960d453f364cacb6f5e"></a>

## cloudfront.protected_endpoints.flow_label.search — cloudfront.protected_endpoints.flow_label.search / 43dc24fe6150 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- cloudfront.protected_endpoints.flow_label.search

<a id="canonical-ac6949da5461d381ef9164f087327928840b79a9b5bdeaf48c399b08c11e298b"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category.

Upstream description:

Bot Defense Flow Label Search Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"flight_search\",\"product_search\",\"reservation_search\",\"room_search\"]"
}
```

<a id="canonical-833a2e4bb375bf3808a41808fe647005eb578d88b8d7bc5886797aa278a4a1b9"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.search / 43dc24fe6150 / 3

- [flight_search](data-sources--protected_application--reference--group-003.md#canonical-dfe7249f46ca6934f6aeeb869ade84ee9c779dde3755d618315547a9ff635b51): complete subsection reference.

- [product_search](data-sources--protected_application--reference--group-003.md#canonical-c98f6e7b70e114e7290234d04642e5a6382326bd1adaaa08a9ff7c49035f6dee): complete subsection reference.

- [reservation_search](data-sources--protected_application--reference--group-003.md#canonical-09586a54356a07ac0935d694bdcd26d23135fbd64eb6da7202955d015b538b34): complete subsection reference.

- [room_search](data-sources--protected_application--reference--group-003.md#canonical-02ba43cb961c4d5daff94fa81f3e503cf7a24318a9108eb89356b0d723675e71): complete subsection reference.

<a id="canonical-d008949a1852f32c382e94810b30dbbb5487327cbb1068435042c1c98fc6e051"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.search / 43dc24fe6150 / 4

- [cloudfront.protected_endpoints.flow_label.search.flight_search](data-sources--protected_application--reference--group-003.md#canonical-dfe7249f46ca6934f6aeeb869ade84ee9c779dde3755d618315547a9ff635b51)
- [cloudfront.protected_endpoints.flow_label.search.product_search](data-sources--protected_application--reference--group-003.md#canonical-c98f6e7b70e114e7290234d04642e5a6382326bd1adaaa08a9ff7c49035f6dee)
- [cloudfront.protected_endpoints.flow_label.search.reservation_search](data-sources--protected_application--reference--group-003.md#canonical-09586a54356a07ac0935d694bdcd26d23135fbd64eb6da7202955d015b538b34)
- [cloudfront.protected_endpoints.flow_label.search.room_search](data-sources--protected_application--reference--group-003.md#canonical-02ba43cb961c4d5daff94fa81f3e503cf7a24318a9108eb89356b0d723675e71)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-dfe7249f46ca6934f6aeeb869ade84ee9c779dde3755d618315547a9ff635b51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af53f1906fb7201cd3b83b1a630583d01fe1c0ba36b3d6caf294951006f3270b"></a>

## cloudfront.protected_endpoints.flow_label.search.flight_search — cloudfront.protected_endpoints.flow_label.search.flight_search / 8eb2bc84e194 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-77ac0bb1a507ea767b5b75f197887be654ecde8ed1b84838d54e3fc8f5dc5903)
- cloudfront.protected_endpoints.flow_label.search.flight_search

<a id="canonical-f83b5130562a452c2d4bc6ca0880565a692d9cc7c49e70aa826d34230fcbb467"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for flight search.

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

<a id="canonical-632d5a15b58304675d8e4512d4a611cd2967bcd0503cf16ca273f09a1a24eeda"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.search.flight_search / 8eb2bc84e194 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-11b38e19a091d2549f6f8ab55d361aa097f6d928a5b22d72c86100ca23e82b43"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.search.flight_search / 8eb2bc84e194 / 4

- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-77ac0bb1a507ea767b5b75f197887be654ecde8ed1b84838d54e3fc8f5dc5903)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-c98f6e7b70e114e7290234d04642e5a6382326bd1adaaa08a9ff7c49035f6dee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c60316ec9fdcfeec552e6651ec2a10c815db8213669c126842e984cafbaebee1"></a>

## cloudfront.protected_endpoints.flow_label.search.product_search — cloudfront.protected_endpoints.flow_label.search.product_search / 25ba7f74ffbc / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-77ac0bb1a507ea767b5b75f197887be654ecde8ed1b84838d54e3fc8f5dc5903)
- cloudfront.protected_endpoints.flow_label.search.product_search

<a id="canonical-86dadaece91ab3eb51f530d34a9541ae63be8f795db62e96c0bb7cdceb1f7f81"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for product search.

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

<a id="canonical-79d2739a93920f7b5240d842f863b0b46f9a9fe4d5880d8328fc02b8c15368a5"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.search.product_search / 25ba7f74ffbc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-075081446f2839347fdf03ab65afd0a69681acda1026c6067d32b9f69f9cc8c1"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.search.product_search / 25ba7f74ffbc / 4

- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-77ac0bb1a507ea767b5b75f197887be654ecde8ed1b84838d54e3fc8f5dc5903)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-09586a54356a07ac0935d694bdcd26d23135fbd64eb6da7202955d015b538b34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6663d4bc510390843707385fb95c1e504e1730fa9201d9dcb1ebdcf831609f0c"></a>

## cloudfront.protected_endpoints.flow_label.search.reservation_search — cloudfront.protected_endpoints.flow_label.search.reservation_search / 09a1e81d3abf / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-77ac0bb1a507ea767b5b75f197887be654ecde8ed1b84838d54e3fc8f5dc5903)
- cloudfront.protected_endpoints.flow_label.search.reservation_search

<a id="canonical-d52676fb7ddcb3baeae9a929385c7b72d5ec0964ac594bd836307f5750caec76"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for reservation search.

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

<a id="canonical-8dacc7692b1de6a84d8eceec5761e26bd0b47dcdc1127b214c20af5630059f74"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.search.reservation_search / 09a1e81d3abf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8a84a80fbe9b3a6083318775ce6ef3fc2504526e86037d71163fbbf4a3fb4da5"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.search.reservation_search / 09a1e81d3abf / 4

- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-77ac0bb1a507ea767b5b75f197887be654ecde8ed1b84838d54e3fc8f5dc5903)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-02ba43cb961c4d5daff94fa81f3e503cf7a24318a9108eb89356b0d723675e71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd2d7ecbfddb3fccf72c9a920fc647188a8e6e1d6788a79dceda026fe864a57d"></a>

## cloudfront.protected_endpoints.flow_label.search.room_search — cloudfront.protected_endpoints.flow_label.search.room_search / 73d0efe76a8c / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-77ac0bb1a507ea767b5b75f197887be654ecde8ed1b84838d54e3fc8f5dc5903)
- cloudfront.protected_endpoints.flow_label.search.room_search

<a id="canonical-8a56a68d4beb4547e2777a4ce603759e30038fb5216966f38f5d99e377526773"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for room search.

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

<a id="canonical-a751d02f00dbb8d8bc8898e178b24339c80b783567a7df93a6832b677141a9cf"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.search.room_search / 73d0efe76a8c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e13bec6cf5ac05701782da25899d42191f44fc91fcbcd5f9d5edd53b09893cb9"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.search.room_search / 73d0efe76a8c / 4

- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-77ac0bb1a507ea767b5b75f197887be654ecde8ed1b84838d54e3fc8f5dc5903)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d145650170a8d5e7ef577dee2e60b1f1b49381d1a602ceb5909c1c9a89bc96b1"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards — cloudfront.protected_endpoints.flow_label.shopping_gift_cards / 025907e5fe73 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards

<a id="canonical-61a1490e2aaa8f28b8fec43669981cf45f641be0ef8dff43ce74e853919b7544"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Shopping &amp; Gift Cards Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"gift_card_make_purchase_with_gift_card\",\"gift_card_validation\",\"shop_add_to_cart\",\"shop_checkout\",\"shop_choose_seat\",\"shop_enter_drawing_submission\",\"shop_make_payment\",\"shop_order\",\"shop_price_inquiry\",\"shop_promo_code_validation\",\"shop_purchase_gift_card\",\"shop_update_quantity\"]"
}
```

<a id="canonical-b605344a0059b083737c3a57da326e95fe3bc4f8ccbf2f5f2b236904f1e4ca8f"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards / 025907e5fe73 / 3

- [gift_card_make_purchase_with_gift_card](data-sources--protected_application--reference--group-003.md#canonical-d61394549b8297c696c7703b5022d09dd2220932652e424b764b55fe5e5d764a): complete subsection reference.

- [gift_card_validation](data-sources--protected_application--reference--group-003.md#canonical-757303dd39fe9c7dceb8d47b5e249db2e7603feb1a0dd49dfcf71980077469be): complete subsection reference.

- [shop_add_to_cart](data-sources--protected_application--reference--group-003.md#canonical-4e793ced89c1db85f1e4136071bcde199d72dae1a68d710efa82e5a54d9ec0a3): complete subsection reference.

- [shop_checkout](data-sources--protected_application--reference--group-003.md#canonical-e112722e236d2a96bf5c2ab2900fea7ce9771761d1b5a9f922d0f82fe37efcd2): complete subsection reference.

- [shop_choose_seat](data-sources--protected_application--reference--group-003.md#canonical-325acc227449b40a0dc9e78726b22a583fe4149346d24c245a708a2ff5763402): complete subsection reference.

- [shop_enter_drawing_submission](data-sources--protected_application--reference--group-003.md#canonical-e0ca4bf4625d58799b8cc29734dd8cc08cca02de68e9449f3c16964201cabbce): complete subsection reference.

- [shop_make_payment](data-sources--protected_application--reference--group-003.md#canonical-3645c21927dc7e6704875ed844edd087c6a2e4972caf609f2012a0c04f07b28a): complete subsection reference.

- [shop_order](data-sources--protected_application--reference--group-003.md#canonical-6166343f5b708f35d099744108a023dad1041fd0b73555f5476ca4bd8fa52d80): complete subsection reference.

- [shop_price_inquiry](data-sources--protected_application--reference--group-003.md#canonical-8d5e7c0951b7d756f8866f8f948358f7d99f4ada964312800008d24ba5f8f335): complete subsection reference.

- [shop_promo_code_validation](data-sources--protected_application--reference--group-003.md#canonical-07e98540575193a58dd95aa66580548b50ca559ff03c8e55dfa0d3509fb77581): complete subsection reference.

- [shop_purchase_gift_card](data-sources--protected_application--reference--group-003.md#canonical-4f08c83143578caa5e2042442ac5ace040dea9159b40f3446b20706244c137ab): complete subsection reference.

- [shop_update_quantity](data-sources--protected_application--reference--group-003.md#canonical-88e98f07188c28900e38c2bd8af99daf3a87490deb8b7ce5d1891a4cc7dbc0cb): complete subsection reference.

<a id="canonical-ddb64b21dc24b00021f1a0fad3b4f0fd7a0434ee4536b3cc9cbef854bc58f2f1"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards / 025907e5fe73 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](data-sources--protected_application--reference--group-003.md#canonical-d61394549b8297c696c7703b5022d09dd2220932652e424b764b55fe5e5d764a)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation](data-sources--protected_application--reference--group-003.md#canonical-757303dd39fe9c7dceb8d47b5e249db2e7603feb1a0dd49dfcf71980077469be)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](data-sources--protected_application--reference--group-003.md#canonical-4e793ced89c1db85f1e4136071bcde199d72dae1a68d710efa82e5a54d9ec0a3)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout](data-sources--protected_application--reference--group-003.md#canonical-e112722e236d2a96bf5c2ab2900fea7ce9771761d1b5a9f922d0f82fe37efcd2)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](data-sources--protected_application--reference--group-003.md#canonical-325acc227449b40a0dc9e78726b22a583fe4149346d24c245a708a2ff5763402)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](data-sources--protected_application--reference--group-003.md#canonical-e0ca4bf4625d58799b8cc29734dd8cc08cca02de68e9449f3c16964201cabbce)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment](data-sources--protected_application--reference--group-003.md#canonical-3645c21927dc7e6704875ed844edd087c6a2e4972caf609f2012a0c04f07b28a)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order](data-sources--protected_application--reference--group-003.md#canonical-6166343f5b708f35d099744108a023dad1041fd0b73555f5476ca4bd8fa52d80)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](data-sources--protected_application--reference--group-003.md#canonical-8d5e7c0951b7d756f8866f8f948358f7d99f4ada964312800008d24ba5f8f335)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](data-sources--protected_application--reference--group-003.md#canonical-07e98540575193a58dd95aa66580548b50ca559ff03c8e55dfa0d3509fb77581)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](data-sources--protected_application--reference--group-003.md#canonical-4f08c83143578caa5e2042442ac5ace040dea9159b40f3446b20706244c137ab)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](data-sources--protected_application--reference--group-003.md#canonical-88e98f07188c28900e38c2bd8af99daf3a87490deb8b7ce5d1891a4cc7dbc0cb)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-d61394549b8297c696c7703b5022d09dd2220932652e424b764b55fe5e5d764a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd47b4d49ff6a92608dade5d16adb997d14b5f84f4b737ef37167a13ff1436b6"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_pur / ad6a393f8e2d / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card

<a id="canonical-d2dad06badf63953db7ba722d5ac344dc821d8e59f09a59bc38b803ef165c7c6"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for gift card make purchase with gift card.

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

<a id="canonical-4b14d6852c7085d3e6e47c4753bad27407e2d6ed148027190364c63e0f8c468b"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_pur / ad6a393f8e2d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c25f79bf93298b0ddd508c6ca882faed0d0eecf45172f251a747f2e781280e7d"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_pur / ad6a393f8e2d / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-757303dd39fe9c7dceb8d47b5e249db2e7603feb1a0dd49dfcf71980077469be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c783e89b72f17050c5519fe7e7e89a7bf4d2b8b89a53b1ec8574005f398a1273"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validati / 1e1dee195b14 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation

<a id="canonical-9c8524837626cff600660b1ff1edd82531451a35b0367dc1cef11198d4caa330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for gift card validation.

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

<a id="canonical-f95ea581df19b418422291080a6445612b9428b1d67ec98cbe2608ac6163e4a8"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validati / 1e1dee195b14 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe23d11853e7c7e8eb4ffa69661e698756b0fb1abd65b6f61d457586fdf94fab"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validati / 1e1dee195b14 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-4e793ced89c1db85f1e4136071bcde199d72dae1a68d710efa82e5a54d9ec0a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d36b41d3c179407374c282a4dc4c87f180a46b27d3c8546fa2bf9c12b6f73205"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart / 8f949ebeef8d / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart

<a id="canonical-13158fcf2d383b852fe24bb064399af4642afd17967288a75af323217741ec72"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop add to cart.

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

<a id="canonical-4a7d69eee814dc43ce0260b82898cf76862933b6516d29f2c3337f536281e603"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart / 8f949ebeef8d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-83a12d6d2c4d6481db7284df4d3f21776260b87a5d0351cef6a0b0cf29633178"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart / 8f949ebeef8d / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-e112722e236d2a96bf5c2ab2900fea7ce9771761d1b5a9f922d0f82fe37efcd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c414e8406cb16d665272bc98bf299a5d21373670d69feb293abcf4155822a215"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout / 48d2dfc97f72 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout

<a id="canonical-1d5aaaf9e8665e61911441bd1722764c46e0d81fd41c2989da3164af8100e169"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop checkout.

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

<a id="canonical-e78492e7ded0f9d748d401a30937b0d659a071964b567144e50dbc1810766b2c"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout / 48d2dfc97f72 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c2f0beda301610d94c28dd53a6c91da1d1a909998d51944b567fd7a5cc62806"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout / 48d2dfc97f72 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-325acc227449b40a0dc9e78726b22a583fe4149346d24c245a708a2ff5763402"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39b04e16ebcfd825ec7e6e9d97f8237ca51cddf4f6c63a4d6247bed2902997fa"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat / 50418f777e9b / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat

<a id="canonical-4c29775288e5ad20377fc86afc15019e0b680dcf07a3dc5b4f1b256a1c927f15"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop choose seat.

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

<a id="canonical-5c4d4537bda302564feb2b143294e3523446c234806779f1e825423c27cd93de"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat / 50418f777e9b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1e78795f2413d7fc2b957e6f3291fbeeeb39b0cd0bc9650b7ba52617cf4135ee"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat / 50418f777e9b / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-e0ca4bf4625d58799b8cc29734dd8cc08cca02de68e9449f3c16964201cabbce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18f00f036daf8fa0c406e8e562b5be25741bcdb698c061653cdc4b805f92b0a6"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing / 6843861d0fda / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission

<a id="canonical-2c47a3a877286ed8705961b0662336f8bc36d3cdab5c752452a6fc69bc71cebe"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop enter drawing submission.

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

<a id="canonical-782192bd567987391e219ded8eb5605b9a52d054d857e7f572097a6ad58737f9"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing / 6843861d0fda / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bcec7b3b5e0cde44427be789bc9e92589ff7450e4e93eff56ecc0944c2e422aa"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing / 6843861d0fda / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-3645c21927dc7e6704875ed844edd087c6a2e4972caf609f2012a0c04f07b28a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6bdb298d002b823bd43e70c2a730d8296d996ad7fc974a983cc804ff938c656e"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment / 8dbf03900240 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment

<a id="canonical-66e6293e57a0bf41fb9aa4e690b7dc8ac103458132183fc81496ec5a72a210a8"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop make payment.

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

<a id="canonical-e4a8f7105e070f5e8b76498634ec51b90af91eca0ac4bb514a40ea80bbc260fd"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment / 8dbf03900240 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e53ca853c9e86207244ad0169278a3991a8a57af61a40872551a4d7dc97b922e"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment / 8dbf03900240 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-6166343f5b708f35d099744108a023dad1041fd0b73555f5476ca4bd8fa52d80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd207ef929efbc05ecaeb87ea2b5cdddbe3925dd9fecbf84a11213d9213c3172"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order / 61c01834fc4a / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order

<a id="canonical-3a88340827e4e858bf075078189c8f064e32606401206d4655e6cc2a536c22b8"></a>

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

<a id="canonical-9dfd0fbdf2985560723ee0ba15a553383a56eefb77774f7ddf96885b9b588874"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order / 61c01834fc4a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-045b7e39433ac149130b520c5992f17950da4a6121d9418275c43a913bfff998"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order / 61c01834fc4a / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-8d5e7c0951b7d756f8866f8f948358f7d99f4ada964312800008d24ba5f8f335"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38ae8773f0b5c78feecd914c6898376548b5e44851be166ff4dc2654d44ea6e6"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry / c7a11ece9224 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry

<a id="canonical-de7a35d7ab5f151498def5fc023797e051fdd138e6d531782aafc2e84af42be4"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop price inquiry.

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

<a id="canonical-8366cf9c985563b8a55b67ec3a21eb141f665d61e74a820a517f6a622a82bf50"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry / c7a11ece9224 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aad904c5bc978b81e7883c90bbe29d51bec2f980028ff0c9b7e73ee5e15782fc"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry / c7a11ece9224 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-07e98540575193a58dd95aa66580548b50ca559ff03c8e55dfa0d3509fb77581"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51bb0a5552914b7f409a16dfc1eb5336629cdb90eddb3091bffca563660eb100"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_va / 25f68b5655a9 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation

<a id="canonical-64da04ae6ff76e58057624039a8f3fd5b80c60f10374434a119dec27b48b9bb2"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop promo code validation.

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

<a id="canonical-1a429f5f37b0577b4a578c9012cf710434697415e0401e7ae05aaecf47ff3bf3"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_va / 25f68b5655a9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7d6055135e875e5840d753ceb7c59c9f3a563c3ce23552adc28da2f534735af5"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_va / 25f68b5655a9 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-4f08c83143578caa5e2042442ac5ace040dea9159b40f3446b20706244c137ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d49b2690654544ae16c20b65d1d38e6576fba6ddb0249f0c1e0bda763979bff1"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift / a91215ec3164 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card

<a id="canonical-ca1ba5fe987f4d28cb0d58d49c89db3f8285621a1e1e53fd74a67e31962d3876"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop purchase gift card.

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

<a id="canonical-39177be6fe30918902e3a465bf78c9af18fb69b0b2522d5015689cd12ae32e92"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift / a91215ec3164 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a17d4166b102fbff3755979a278ed89e3508cf1c1b1f8c411fe8ca7edca9b5cc"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift / a91215ec3164 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-88e98f07188c28900e38c2bd8af99daf3a87490deb8b7ce5d1891a4cc7dbc0cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f451e1a65268bfa2d57f518b4a8025c8ebbff823a107d2ed17b8b99773be03b8"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quanti / 2d05f4d3f070 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity

<a id="canonical-2721103d17b010921d8151fb2c0ebda589d2e6480af7c55457536893c3da84fd"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop update quantity.

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

<a id="canonical-56ffc0a8b160273342c1c65568b21ba8581528ef9e99c383c95877e3e7c7878f"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quanti / 2d05f4d3f070 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bf8e3baaf979f1d21a3fdd0b110981777c4bdc182856bfec82e917d2edf971fd"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quanti / 2d05f4d3f070 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-2b870022d2128ab07045747482ae04bcffbe43b179503838a38dd05027584939)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-89b33e808ca3e0c9d8b1aa72cca8cfdca797e6766a6bfc54e8c606d221d7b6ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1be6aa098ae9b5c203c792245073b55d23c70829780296d4ebdf0b65f1280dc"></a>

## cloudfront.protected_endpoints.metadata — cloudfront.protected_endpoints.metadata / fd57d2953188 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- cloudfront.protected_endpoints.metadata

<a id="canonical-5405021cd33c592ad0ff76d3fafdb26f39f8798d114e2f1c59bd531bbc0450ba"></a>

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

<a id="canonical-44c91e8d2aff09ac27119987da1287ff0429c4029e71c5bc1b1d11d34d0fc5d7"></a>

## Direct properties — cloudfront.protected_endpoints.metadata / fd57d2953188 / 3

<a id="canonical-0cb739c676bc6fa512e066079f94ad71ecbf538cc7b28162fd015fa4a49afaba"></a>

<a id="canonical-97ab021412a0a019a0579031dd034a13d7aaa5233e257006cfbfe8f50a213167"></a>

## description_spec property — cloudfront.protected_endpoints.metadata / fd57d2953188 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-c9dde01e59931b906aa510c09e13935c44e57e57390190c5cf603073395871ea"></a>

<a id="canonical-ff21a57c7cdb36fa6f46128f82de975fc8d66e3c32fc4faba43ee5e4609bb0f1"></a>

## name property — cloudfront.protected_endpoints.metadata / fd57d2953188 / 5

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

<a id="canonical-2019d2fc20c078542a0c168b516dd77748f01a774b7eeba240d8df912d4e33e1"></a>

## Next pages — cloudfront.protected_endpoints.metadata / fd57d2953188 / 6

- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-3041f54a787ce2d720d9ad74bcd8214f1b51726f556bd6613a34596b7b59f37a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-feb00bd8877c76e26e070b6f22a54ced32d1b7f7e23c1a0675ecfdb824d81085"></a>

## cloudfront.protected_endpoints.mobile_client — cloudfront.protected_endpoints.mobile_client / 9185d8956540 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- cloudfront.protected_endpoints.mobile_client

<a id="canonical-0659ed68d7590abeb654a95a10740d0b4c19ace0614b96c3f403485aeac5242e"></a>

Type: `"single"`. Computed.

Mobile Client. Mobile client configuration OPTIONS.

Upstream description:

Mobile client configuration OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\"]"
}
```

<a id="canonical-e2cf8a07ac4f2bac1dff28eac7d13d0dc43e5eff3fcb30fbfb24792f79835046"></a>

## Direct properties — cloudfront.protected_endpoints.mobile_client / 9185d8956540 / 3

- [block](data-sources--protected_application--reference--group-003.md#canonical-8dba2bf0cd067db40b889a2212ab61caaa8029797c057e989bb98c6bf5f65ac8): complete subsection reference.

- [continue](data-sources--protected_application--reference--group-003.md#canonical-0b992bcc745efd1f4156e33e0da8dcce460f1dfcdab05ec358d20a2c04f19063): complete subsection reference.

<a id="canonical-79b5f7ac259d378e9af8944ac8ea093987a78d9e0dfb44c868a833e04d72b71e"></a>

## Next pages — cloudfront.protected_endpoints.mobile_client / 9185d8956540 / 4

- [cloudfront.protected_endpoints.mobile_client.block](data-sources--protected_application--reference--group-003.md#canonical-8dba2bf0cd067db40b889a2212ab61caaa8029797c057e989bb98c6bf5f65ac8)
- [cloudfront.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-003.md#canonical-0b992bcc745efd1f4156e33e0da8dcce460f1dfcdab05ec358d20a2c04f19063)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-8dba2bf0cd067db40b889a2212ab61caaa8029797c057e989bb98c6bf5f65ac8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d9dccde5aa92d00d17f391c020eb1c3291bebc7208b161e0c30abda874728f6"></a>

## cloudfront.protected_endpoints.mobile_client.block — cloudfront.protected_endpoints.mobile_client.block / 38df124bbad4 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-3041f54a787ce2d720d9ad74bcd8214f1b51726f556bd6613a34596b7b59f37a)
- cloudfront.protected_endpoints.mobile_client.block

<a id="canonical-345dbf3c47564f8a4dd6216651a48a8726c404916eec97827a02dc82cf86f60f"></a>

Type: `"single"`. Computed.

Block Response for Mobile. Block Response.

Upstream description:

Block Response.

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

<a id="canonical-cd7861582255cb84ebca71c3961b1fac8da135730fdece9f94c9cf0f94f9f6bc"></a>

## Direct properties — cloudfront.protected_endpoints.mobile_client.block / 38df124bbad4 / 3

<a id="canonical-2654cf1701eef24bb3ad724d6ecea17fc4889186b2a641f0689834f1d71a6671"></a>

<a id="canonical-52ae8aaf81849196b1258d7edb44ca830b92732b21407121dd04742788953253"></a>

## body property — cloudfront.protected_endpoints.mobile_client.block / 38df124bbad4 / 4

Type: `"string"`. Computed.

Body. Custom body message.

Upstream description:

Custom body message.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-91fbb745c401c5b30ab351a739bd0c082a09849608fe6f26b3836312f57fbadb"></a>

<a id="canonical-133a885613814a78fe463a89c8cbe1325f7122104f2860cf9216e58b78985db1"></a>

## content_type property — cloudfront.protected_endpoints.mobile_client.block / 38df124bbad4 / 5

Type: `"string"`. Computed.

Content type to use in a block response.

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

<a id="canonical-cb901675392f067fafe2effef9a3f9d41938c3c63f5846e1360654450357e84e"></a>

<a id="canonical-a69f51921ca269ee13efad3dee09c1bc3117d32698f92138ede2e025af5cc354"></a>

## status property — cloudfront.protected_endpoints.mobile_client.block / 38df124bbad4 / 6

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-93bca79cd41e547f7d6389a7312a93525f793f50034e4e6843576fd1406ddac8"></a>

## Next pages — cloudfront.protected_endpoints.mobile_client.block / 38df124bbad4 / 7

- [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-3041f54a787ce2d720d9ad74bcd8214f1b51726f556bd6613a34596b7b59f37a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-0b992bcc745efd1f4156e33e0da8dcce460f1dfcdab05ec358d20a2c04f19063"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc59dbd38d856a4c45d3ab4af365685b2030954d5273653ffa656d6313d2e2c9"></a>

## cloudfront.protected_endpoints.mobile_client.continue — cloudfront.protected_endpoints.mobile_client.continue / d872a54bff21 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-3041f54a787ce2d720d9ad74bcd8214f1b51726f556bd6613a34596b7b59f37a)
- cloudfront.protected_endpoints.mobile_client.continue

<a id="canonical-a8798c5dcd12bb07048698624756cf76491d09124bfaf7cbc791f7d48b7d9be7"></a>

Type: `"single"`. Computed.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-add_header_choice": "[\"add_header\",\"no_header\"]"
}
```

<a id="canonical-5246dd94cd934a69ee7ee6a6b5d961195d7f4bc930ee4b9eb7aba4d35c1b15f7"></a>

## Direct properties — cloudfront.protected_endpoints.mobile_client.continue / d872a54bff21 / 3

- [add_header](data-sources--protected_application--reference--group-003.md#canonical-85212d029dc02edc88fbb9a31992fe896f99c0f4762219e6f60b7cfd4d33ce66): complete subsection reference.

- [no_header](data-sources--protected_application--reference--group-003.md#canonical-2873c85a49f24358108574217ff18bd7ec7bae9ef6de46cec0999baeb936e4a2): complete subsection reference.

<a id="canonical-54d4d591f743e7f2c0aaaa613de110cb31e22a72afa78d8e6dd97d66c8c970f8"></a>

## Next pages — cloudfront.protected_endpoints.mobile_client.continue / d872a54bff21 / 4

- [cloudfront.protected_endpoints.mobile_client.continue.add_header](data-sources--protected_application--reference--group-003.md#canonical-85212d029dc02edc88fbb9a31992fe896f99c0f4762219e6f60b7cfd4d33ce66)
- [cloudfront.protected_endpoints.mobile_client.continue.no_header](data-sources--protected_application--reference--group-003.md#canonical-2873c85a49f24358108574217ff18bd7ec7bae9ef6de46cec0999baeb936e4a2)
- [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-3041f54a787ce2d720d9ad74bcd8214f1b51726f556bd6613a34596b7b59f37a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-85212d029dc02edc88fbb9a31992fe896f99c0f4762219e6f60b7cfd4d33ce66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75db4e1c9b1054cee27ca69337fc55ba4616e1a108a95e600a611f634b95bf65"></a>

## cloudfront.protected_endpoints.mobile_client.continue.add_header — cloudfront.protected_endpoints.mobile_client.continue.add_header / 3e1cebe63659 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-3041f54a787ce2d720d9ad74bcd8214f1b51726f556bd6613a34596b7b59f37a)
- [cloudfront.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-003.md#canonical-0b992bcc745efd1f4156e33e0da8dcce460f1dfcdab05ec358d20a2c04f19063)
- cloudfront.protected_endpoints.mobile_client.continue.add_header

<a id="canonical-b4ef99d56c6718413b4b1816d50c31d1d9ce63c32c9acf85109ec2bcea3b70e9"></a>

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

<a id="canonical-e1851e593682e7bb8edc1b4d57d79094e0ac858b238cd34a13b142443f3089b8"></a>

## Direct properties — cloudfront.protected_endpoints.mobile_client.continue.add_header / 3e1cebe63659 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-642964003c67091df7c20e96bb269c0d7a03a7badc205ff02fe6a75150daa288"></a>

## Next pages — cloudfront.protected_endpoints.mobile_client.continue.add_header / 3e1cebe63659 / 4

- [cloudfront.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-003.md#canonical-0b992bcc745efd1f4156e33e0da8dcce460f1dfcdab05ec358d20a2c04f19063)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-2873c85a49f24358108574217ff18bd7ec7bae9ef6de46cec0999baeb936e4a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb068a5b2d378f819a09cf3ad311dbf88dfa624c06da82b6aadb52190ec3a191"></a>

## cloudfront.protected_endpoints.mobile_client.continue.no_header — cloudfront.protected_endpoints.mobile_client.continue.no_header / 743688a8b1db / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-3041f54a787ce2d720d9ad74bcd8214f1b51726f556bd6613a34596b7b59f37a)
- [cloudfront.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-003.md#canonical-0b992bcc745efd1f4156e33e0da8dcce460f1dfcdab05ec358d20a2c04f19063)
- cloudfront.protected_endpoints.mobile_client.continue.no_header

<a id="canonical-e3b924cbc32b6e2e3930d19dfbcac1e780f352f89f6f9c3db51a05b4a04e4852"></a>

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

<a id="canonical-63a0282939673de96b343585e2343219d6f5520a96e159f7e97f8fb4f37de6e3"></a>

## Direct properties — cloudfront.protected_endpoints.mobile_client.continue.no_header / 743688a8b1db / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331f1911619a051f983d2bf4f25e3fedee55e743d7718ccb84ac1d1b51b763f"></a>

## Next pages — cloudfront.protected_endpoints.mobile_client.continue.no_header / 743688a8b1db / 4

- [cloudfront.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-003.md#canonical-0b992bcc745efd1f4156e33e0da8dcce460f1dfcdab05ec358d20a2c04f19063)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-9cefac8a26e21e93a205a05fd4b81f09b05eb4e939dd5edcfcef6f2fc77d8cda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4234d8606275392fffde989be53a48250018f79e54219d1bf14e775d711a1552"></a>

## cloudfront.protected_endpoints.undefined_flow_label — cloudfront.protected_endpoints.undefined_flow_label / ba7d1202b8ee / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- cloudfront.protected_endpoints.undefined_flow_label

<a id="canonical-5d2db764f5baadc9931c395646f500c04ea88ffea0c4806ff58885bccf450dd7"></a>

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

<a id="canonical-dc93094a250b2c62cf9b77823aa6486c0471de7e4c3fce72be04362397032d76"></a>

## Direct properties — cloudfront.protected_endpoints.undefined_flow_label / ba7d1202b8ee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-44d7bb43a7e32f30556d5c879884d6f8f423771f8fe42718f5154ae5d70c8dd7"></a>

## Next pages — cloudfront.protected_endpoints.undefined_flow_label / ba7d1202b8ee / 4

- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-09769448fb2289c7e9c1998fcacb5f88119778cb1f52a1322920154605743268"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb72bce5dee30ab6bc0503daa6cf82f9169c9e3fa12d41ee6f717641a5ab9fd9"></a>

## cloudfront.protected_endpoints.web_client — cloudfront.protected_endpoints.web_client / e7c75d18791c / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- cloudfront.protected_endpoints.web_client

<a id="canonical-1e1c463393e17da6849f73ae17baeb59e61fe8234427995474009bea100d9afd"></a>

Type: `"single"`. Computed.

Web Client. Web client configuration OPTIONS.

Upstream description:

Web client configuration OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\",\"redirect\"]"
}
```

<a id="canonical-5c10f9cd86ec635220184bab2517d98b479554daabeae143a4f21cde141a70f1"></a>

## Direct properties — cloudfront.protected_endpoints.web_client / e7c75d18791c / 3

- [block](data-sources--protected_application--reference--group-003.md#canonical-471c9fec1df383ad6391c8d3688b655c0db14800fb0526b00919265d7cf1f358): complete subsection reference.

- [continue](data-sources--protected_application--reference--group-003.md#canonical-9c616aa10f7bfa031db985189a8e0015ee07a18ca77f5bc4ae09138108bca394): complete subsection reference.

- [redirect](data-sources--protected_application--reference--group-004.md#canonical-beba1139081b0aaa906fe01e4149d779136c0e8f496605f674da549f5d4eb50f): complete subsection reference.

<a id="canonical-568a9eb1347c7c909f796487e875808213169128ddc1ba8314f4f09d9a9abd96"></a>

## Next pages — cloudfront.protected_endpoints.web_client / e7c75d18791c / 4

- [cloudfront.protected_endpoints.web_client.block](data-sources--protected_application--reference--group-003.md#canonical-471c9fec1df383ad6391c8d3688b655c0db14800fb0526b00919265d7cf1f358)
- [cloudfront.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-003.md#canonical-9c616aa10f7bfa031db985189a8e0015ee07a18ca77f5bc4ae09138108bca394)
- [cloudfront.protected_endpoints.web_client.redirect](data-sources--protected_application--reference--group-004.md#canonical-beba1139081b0aaa906fe01e4149d779136c0e8f496605f674da549f5d4eb50f)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-471c9fec1df383ad6391c8d3688b655c0db14800fb0526b00919265d7cf1f358"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-832850f02c586519ede815d1fd43f6f07c490df64b947590c37f25c18a78f587"></a>

## cloudfront.protected_endpoints.web_client.block — cloudfront.protected_endpoints.web_client.block / e9a99a3a7eae / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.protected_endpoints.web_client](data-sources--protected_application--reference--group-003.md#canonical-09769448fb2289c7e9c1998fcacb5f88119778cb1f52a1322920154605743268)
- cloudfront.protected_endpoints.web_client.block

<a id="canonical-097aa2b346d55f7792de416891dce767a2cc3638c68fef38436c52a2df9f9586"></a>

Type: `"single"`. Computed.

Block Response. Block Response.

Upstream description:

Block Response.

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

<a id="canonical-14b0493b2267e436d995279d9b00d38efd3d3345234730d2bb1f7cad552b5160"></a>

## Direct properties — cloudfront.protected_endpoints.web_client.block / e9a99a3a7eae / 3

<a id="canonical-3702bb34cd6e1b30465c1f859a3ba0578aaaefba09e954552fccc697bcaf9b23"></a>

<a id="canonical-25c4db83e785cb7874c0d63d8336a78688257262d94d14f6921d1c545f7262fd"></a>

## body property — cloudfront.protected_endpoints.web_client.block / e9a99a3a7eae / 4

Type: `"string"`. Computed.

Body. Custom body message.

Upstream description:

Custom body message.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-366fcc37dddc44c60826b4ec30301c340d1d9ddaecd0358c0946dfe463da7122"></a>

<a id="canonical-270392bb7a9a94b980b0411a7dc368ccc1b863a20b54b128fcfaa984884ad9ec"></a>

## content_type property — cloudfront.protected_endpoints.web_client.block / e9a99a3a7eae / 5

Type: `"string"`. Computed.

Content type to use in a block response.

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

<a id="canonical-59dcfee5a792b2908d53b12b1f89a8c434569284d5aea07e6318f1656f4647a2"></a>

<a id="canonical-f5f298c6f824966e4c54cf642a37ec0be79e5113582a6e48955d98efe50ffe56"></a>

## status property — cloudfront.protected_endpoints.web_client.block / e9a99a3a7eae / 6

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e80f1b79d58498e6dadc8d0cbe41664c30723b1b41075f48fc7f525d06c4403d"></a>

## Next pages — cloudfront.protected_endpoints.web_client.block / e9a99a3a7eae / 7

- [cloudfront.protected_endpoints.web_client](data-sources--protected_application--reference--group-003.md#canonical-09769448fb2289c7e9c1998fcacb5f88119778cb1f52a1322920154605743268)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-9c616aa10f7bfa031db985189a8e0015ee07a18ca77f5bc4ae09138108bca394"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
