---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-3f765e18a68365c2ff923f3ce02ed5059df7c23801957b4f5371b2ee4e025bcb"></a>

## vpc_attachments.vpc_list — vpc_attachments.vpc_list / 54fa64c7cb9f / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vpc_attachments](data-sources--aws_tgw_site--reference--group-003.md#canonical-411569971dc5e487a1b7eb5dd5026c6a951d476dc5d2fda3bedd09ca3b953633)
- vpc_attachments.vpc_list

<a id="canonical-a09dee683a033d7161a8faea75075e45b4e60f0c239204b6225ace0cd6ec5dc6"></a>

Type: `"list"`. Computed.

List of VPC attachments to transit gateway.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-b31d25f51a6f134aab76a45f8df2ff6d0a9b1dfe6827ee3dd45351c58f2b6d46"></a>

## Direct properties — vpc_attachments.vpc_list / 54fa64c7cb9f / 3

- [labels](data-sources--aws_tgw_site--reference--group-004.md#canonical-7c006df67dcf0cfe86a0d71393499d8bcd08b856ffb65b49691009665be15c4d): complete subsection reference.

<a id="canonical-03200909b2d67a7d6a89cba0b72e76e1afc4f4a7a40607498f0f13fb676a0ef1"></a>

<a id="canonical-b0208e5a4a6345e32e87ef69da110b1d4f8ea11b72aa13a7698e5488db66277d"></a>

## vpc_id property — vpc_attachments.vpc_list / 54fa64c7cb9f / 4

Type: `"string"`. Computed.

VPC ID. Information about existing VPC.

Upstream description:

Information about existing VPC.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-2d68dc15e0b264bc9dec47d5d18c82419f58e0244404bb6f430953468be16c82"></a>

## Next pages — vpc_attachments.vpc_list / 54fa64c7cb9f / 5

- [vpc_attachments.vpc_list.labels](data-sources--aws_tgw_site--reference--group-004.md#canonical-7c006df67dcf0cfe86a0d71393499d8bcd08b856ffb65b49691009665be15c4d)
- [vpc_attachments](data-sources--aws_tgw_site--reference--group-003.md#canonical-411569971dc5e487a1b7eb5dd5026c6a951d476dc5d2fda3bedd09ca3b953633)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-7c006df67dcf0cfe86a0d71393499d8bcd08b856ffb65b49691009665be15c4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db77ec6782bd0ac92e53e4967ed3158797eb3fcb2fdf2e7c38bb058884407311"></a>

## vpc_attachments.vpc_list.labels — vpc_attachments.vpc_list.labels / 8e50b9459f21 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vpc_attachments](data-sources--aws_tgw_site--reference--group-003.md#canonical-411569971dc5e487a1b7eb5dd5026c6a951d476dc5d2fda3bedd09ca3b953633)
- [vpc_attachments.vpc_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-104588b3ea300c40ea94103c1c27bc7107084b99d8659bb14ae72f72e3d6ad08)
- vpc_attachments.vpc_list.labels

<a id="canonical-d56f24b645f0a3f71bb0d60fe00ac8d648fa5b80898c10f1790f8d71bf2958c7"></a>

Type: `"single"`. Computed.

Add labels for the VPC attachment. These labels can then be used in policies such as enhanced
firewall.

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

<a id="canonical-3376b501177343c374c9d44c77ab6ebae6581465ffe12ca322baf42a21b3461d"></a>

## Direct properties — vpc_attachments.vpc_list.labels / 8e50b9459f21 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0ce6ca07c5a14497f5a2d9714de9624c8cf1989023ebe0ff421751c195458e6c"></a>

## Next pages — vpc_attachments.vpc_list.labels / 8e50b9459f21 / 4

- [vpc_attachments.vpc_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-104588b3ea300c40ea94103c1c27bc7107084b99d8659bb14ae72f72e3d6ad08)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-0f930540b03ce3ec08b48a152be09afd0635e044541923a65add395e816e0f26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3852e3a82902c12dce48a8e4386e12ad158e703dc8d23fb02190e05770f20a97"></a>

## waf_signatures — waf_signatures / 0906332a750c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- waf_signatures

<a id="canonical-8267bda250f978b920e3a7bc785efd78537994e0f9db1cbed1c9399489be0684"></a>

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

<a id="canonical-bab92fd9b328503592ef7965744a7d21e406b6c972fcb3ea92a921509fdaf130"></a>

## Direct properties — waf_signatures / 0906332a750c / 3

- [automatic](data-sources--aws_tgw_site--reference--group-004.md#canonical-37db7291f83cf1b543d0a6135db0dcb08ce6683b6af6f1dc9471b44e4a2a1394): complete subsection reference.

- [manual](data-sources--aws_tgw_site--reference--group-004.md#canonical-ef66e13dcc1f22c21ca5e6cd7526b388bc3acf9b33dd5b6d3adead62ccde5e8a): complete subsection reference.

<a id="canonical-3b7493c449913a20b4ea2be379f7d063dc94c580d8a162ff8bca84227548b820"></a>

## Next pages — waf_signatures / 0906332a750c / 4

- [waf_signatures.automatic](data-sources--aws_tgw_site--reference--group-004.md#canonical-37db7291f83cf1b543d0a6135db0dcb08ce6683b6af6f1dc9471b44e4a2a1394)
- [waf_signatures.manual](data-sources--aws_tgw_site--reference--group-004.md#canonical-ef66e13dcc1f22c21ca5e6cd7526b388bc3acf9b33dd5b6d3adead62ccde5e8a)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-37db7291f83cf1b543d0a6135db0dcb08ce6683b6af6f1dc9471b44e4a2a1394"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd5b706577e9c0af649f120c2058e93eb70ea5235ec587fc285141591f4706fb"></a>

## waf_signatures.automatic — waf_signatures.automatic / ba9102bc4785 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-0f930540b03ce3ec08b48a152be09afd0635e044541923a65add395e816e0f26)
- waf_signatures.automatic

<a id="canonical-73bdd94497da37556150b81291e358fff53c59a6811b69b53ae9180233ecf461"></a>

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

<a id="canonical-73c955ff643ac19e2f4259ad92df4ee208e8ba79261083fec0424f992762db70"></a>

## Direct properties — waf_signatures.automatic / ba9102bc4785 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9795f329b3cdd7ee8989e310b1c17cfb59b83cb71880134cd2308630a63fdfd0"></a>

## Next pages — waf_signatures.automatic / ba9102bc4785 / 4

- [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-0f930540b03ce3ec08b48a152be09afd0635e044541923a65add395e816e0f26)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-ef66e13dcc1f22c21ca5e6cd7526b388bc3acf9b33dd5b6d3adead62ccde5e8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-884fede29abefbbb6ec59fd89cb051fb3a0fd5b6c74ce70863e828c8abf24c73"></a>

## waf_signatures.manual — waf_signatures.manual / 3f76569b635c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-0f930540b03ce3ec08b48a152be09afd0635e044541923a65add395e816e0f26)
- waf_signatures.manual

<a id="canonical-620c2c68cecd3fd166330a8ebc02629b1c1ae990fd91f5fc227b5a2b136e5877"></a>

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

<a id="canonical-485111ff79848f55193c6336bd7409a97e40ba5c85e0630bf01c8e961f683a8a"></a>

## Direct properties — waf_signatures.manual / 3f76569b635c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-57386fb597f8c12243dbe56e2040cfffb9a31c3b0277d8ab64493213dabeda43"></a>

## Next pages — waf_signatures.manual / 3f76569b635c / 4

- [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-0f930540b03ce3ec08b48a152be09afd0635e044541923a65add395e816e0f26)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
