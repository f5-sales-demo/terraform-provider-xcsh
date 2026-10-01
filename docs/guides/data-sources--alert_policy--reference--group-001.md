---
page_title: "xcsh_alert_policy reference"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_alert_policy reference."
---

# xcsh_alert_policy reference

<a id="canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9311a8ba242bf24c1788a5691f36281e4f3b5efe6b55e9033f587a4482c0ab3"></a>

## Property reference — Property reference / a7fe95349bdc / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- Property reference

<a id="canonical-39a4d38a75dc729783bb53442f2fb2739febf15db55a99b07072c70d03cdb8b6"></a>

## Direct properties — Property reference / a7fe95349bdc / 3

<a id="canonical-b9bab96dc04bcc4e701d1be6b678cab49738d29ad9a751ef0f7a02ae5285671a"></a>

<a id="canonical-15a1b369ea52e8eb49cdcf0897c13719600b0220a4cd74a5d45396fa6cdde01b"></a>

## annotations property — Property reference / a7fe95349bdc / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-286b2fc4d3e17824c108361cd0bb7d024652674b6b26f7d7ae79336a3d8aa7e7"></a>

<a id="canonical-da44244826848a2b8a033d50e8c2516aaa6a0abf7e85066a6d42aaa2f91c0f5f"></a>

## description property — Property reference / a7fe95349bdc / 5

Type: `"string"`. Computed.

Description of the AlertPolicy.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-deeac1fee149aceb7a326f1a813721df2bad811a489a85067d293b9056b5c099"></a>

<a id="canonical-c80dc74a81fe3c112f6554d90e497f00bc1bae49c2131e98e9490494dfd4b3c9"></a>

## id property — Property reference / a7fe95349bdc / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-ca00469f8cbae888b11125b36deeb70b35d2e3876606141905d7cb47f95aadbb"></a>

<a id="canonical-a0ef8df564bb339ea9fa4e9c547096c10415e5182b707e18f1b96664925479af"></a>

## labels property — Property reference / a7fe95349bdc / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="canonical-28a99578bbcde8637d0f3530ffa6bdbe209dff7f3536d600dea9755a9aa485d4"></a>

<a id="canonical-9d35b9df7d4943d983527d47765bba0d49f4bd62c3c7561b9cd389feefefdbde"></a>

## name property — Property reference / a7fe95349bdc / 8

Type: `"string"`. Required.

Name of the AlertPolicy.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

<a id="canonical-8b3438ca1c5687ccead7519116a81798432f29dba231a6df20774affcc9cb326"></a>

<a id="canonical-8d8e9979b13e9d448ca666f869805e7f095b9f8af68f83a87d78efe0c986426d"></a>

## namespace property — Property reference / a7fe95349bdc / 9

Type: `"string"`. Required.

Namespace where the AlertPolicy exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-867b7a75fbb3878ec9962950b54cc863e180e5d8e2486731ca9249f53d5b3720): complete subsection reference.

- [receivers](data-sources--alert_policy--reference--group-001.md#canonical-d11555e3f8eb4b9c7838e3d9697c801749bd384c2925c41e4a0b080052fcca7d): complete subsection reference.

- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2): complete subsection reference.

<a id="canonical-ee7442d3fbfdaafb4615ac353243aeb2f6a33f7309c80b9bd5e889e0f9ff6915"></a>

## All schema paths — Property reference / a7fe95349bdc / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--alert_policy--reference--group-001.md#canonical-b9bab96dc04bcc4e701d1be6b678cab49738d29ad9a751ef0f7a02ae5285671a) |
| `description` | [description](data-sources--alert_policy--reference--group-001.md#canonical-286b2fc4d3e17824c108361cd0bb7d024652674b6b26f7d7ae79336a3d8aa7e7) |
| `id` | [id](data-sources--alert_policy--reference--group-001.md#canonical-deeac1fee149aceb7a326f1a813721df2bad811a489a85067d293b9056b5c099) |
| `labels` | [labels](data-sources--alert_policy--reference--group-001.md#canonical-ca00469f8cbae888b11125b36deeb70b35d2e3876606141905d7cb47f95aadbb) |
| `name` | [name](data-sources--alert_policy--reference--group-001.md#canonical-28a99578bbcde8637d0f3530ffa6bdbe209dff7f3536d600dea9755a9aa485d4) |
| `namespace` | [namespace](data-sources--alert_policy--reference--group-001.md#canonical-8b3438ca1c5687ccead7519116a81798432f29dba231a6df20774affcc9cb326) |
| `notification_parameters` | [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-f5f6b559e1075025538b938e56a33a1f0bf2a84b36789e3975edb2278058e6dd) |
| `notification_parameters.custom` | [notification_parameters.custom](data-sources--alert_policy--reference--group-001.md#canonical-ec9a54faae87c31323a0dee3101072b95fe11e7f691097a98ceffebe4f49ec0c) |
| `notification_parameters.custom.labels` | [notification_parameters.custom.labels](data-sources--alert_policy--reference--group-001.md#canonical-347a200130c7614f8983db93bb9cad505a568b8a02ac53199976c1ca005e805b) |
| `notification_parameters.default` | [notification_parameters.default](data-sources--alert_policy--reference--group-001.md#canonical-8b4e0aa6dba90446cc209bf47b8360470da4d54097b6d2b7779fef49ded06051) |
| `notification_parameters.group_interval` | [notification_parameters.group_interval](data-sources--alert_policy--reference--group-001.md#canonical-d7b6cb3a9b80c309f9382796e78a13b874fc0ddc45a6732011bb79f20ad89fb4) |
| `notification_parameters.group_wait` | [notification_parameters.group_wait](data-sources--alert_policy--reference--group-001.md#canonical-30bd6c709b205ea483f53f4ffdb170aa01985b90c4392304c6b42a0d764c0b48) |
| `notification_parameters.individual` | [notification_parameters.individual](data-sources--alert_policy--reference--group-001.md#canonical-391165b97eacb0753172bbb2fba3f93f4d9b386aec004bd096894377a8cb396f) |
| `notification_parameters.repeat_interval` | [notification_parameters.repeat_interval](data-sources--alert_policy--reference--group-001.md#canonical-d27bb88af9dc0f600567688cee578ffc3abf09d173234c6c485e1b6a91d1246f) |
| `notification_parameters.ves_io_group` | [notification_parameters.ves_io_group](data-sources--alert_policy--reference--group-001.md#canonical-2d1971c9d08e019a847c966507a87330a7a23b8b12788d18c41c72643076cff2) |
| `receivers` | [receivers](data-sources--alert_policy--reference--group-001.md#canonical-1d30c1eb7244a9ef78d1de0d3e5a3ad2d31f50cf96efd4dc7319ce67b87ba5c3) |
| `receivers.kind` | [receivers.kind](data-sources--alert_policy--reference--group-001.md#canonical-02130ac6fd48fe8080954ae38d36e8c0011c480ff4374193fac221fbda5285cb) |
| `receivers.name` | [receivers.name](data-sources--alert_policy--reference--group-001.md#canonical-caad5e4fae398841c61c7742d131d00c6c5d1e5d53605621734144bf182491c3) |
| `receivers.namespace` | [receivers.namespace](data-sources--alert_policy--reference--group-001.md#canonical-b0a683b18c1d336fea6f5b4ac4648f2e93d4af0f84178744113b6c7eed7d4bf4) |
| `receivers.tenant` | [receivers.tenant](data-sources--alert_policy--reference--group-001.md#canonical-f5c5ae148ea84c5cf07d07dc1632754f0b44bc3e1d72b95447283cc670d476d5) |
| `receivers.uid` | [receivers.uid](data-sources--alert_policy--reference--group-001.md#canonical-ef43f2f10587d89a3c94c67a9ad7b2c6e235aaec3f90305e4542ed5e0520340e) |
| `routes` | [routes](data-sources--alert_policy--reference--group-001.md#canonical-523d7ddfdf198f4ae4cd98d3bb5a37953414ed24e44786bb38d6ef13b9c31de0) |
| `routes.alertname` | [routes.alertname](data-sources--alert_policy--reference--group-001.md#canonical-124468f69d1b3ae53b9c589198d33cd28324e85936c57c17a56768653676c146) |
| `routes.alertname_regex` | [routes.alertname_regex](data-sources--alert_policy--reference--group-001.md#canonical-6e092c1c5774d9af22791f167b9f7ee7f20b84f48b33166c0d2ea0c5d57cc151) |
| `routes.any` | [routes.any](data-sources--alert_policy--reference--group-001.md#canonical-f98dfba1d133027f615bdcf1d2dacbaa79003f09de28e998271ec9dfb4d11bbe) |
| `routes.custom` | [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-507d445dc3e69fe1334c574326b1e41bdc454267e1111b01b4702a667c3562d1) |
| `routes.custom.alertlabel` | [routes.custom.alertlabel](data-sources--alert_policy--reference--group-001.md#canonical-30e389c821b163c0fd2530c4f08557ed2acca2a62c135dcde9e2e29d9034a371) |
| `routes.custom.alertname` | [routes.custom.alertname](data-sources--alert_policy--reference--group-001.md#canonical-25e04921716757354012d7ba802bc266ad2916b1e75ad02e7cab8a91ab1619f2) |
| `routes.custom.alertname.exact_match` | [routes.custom.alertname.exact_match](data-sources--alert_policy--reference--group-001.md#canonical-2c0d7d1abf0875da23cabe4cd4fb3e5e90f922266508a7f5f275a3154cdfb63a) |
| `routes.custom.alertname.regex_match` | [routes.custom.alertname.regex_match](data-sources--alert_policy--reference--group-001.md#canonical-b88acb351baabcaa3bb014fb57ea26ff98558216d412fef5238534305b3fb5ed) |
| `routes.custom.group` | [routes.custom.group](data-sources--alert_policy--reference--group-001.md#canonical-6afed69dc549049fc9998472c79f3473afb71ab6f8a746ec492998b76b7d010f) |
| `routes.custom.group.exact_match` | [routes.custom.group.exact_match](data-sources--alert_policy--reference--group-001.md#canonical-01fc7c08346b685dbc9ca687e262152a36fc082e3474f8f23788c3855ee7fcc2) |
| `routes.custom.group.regex_match` | [routes.custom.group.regex_match](data-sources--alert_policy--reference--group-001.md#canonical-ae8087a742a160ec29c2b9675771c24debdc0f70648ba7ec2d029b9c1be96c86) |
| `routes.custom.severity` | [routes.custom.severity](data-sources--alert_policy--reference--group-001.md#canonical-301f805afb6ca47249a3257993be9c8c349fd226d03c5a00e559f7d2925655e5) |
| `routes.custom.severity.exact_match` | [routes.custom.severity.exact_match](data-sources--alert_policy--reference--group-001.md#canonical-116fcb02b0eb17e6dec907137c5f1e3537a6951db99a44ae4c754ef311607a46) |
| `routes.custom.severity.regex_match` | [routes.custom.severity.regex_match](data-sources--alert_policy--reference--group-001.md#canonical-f511b16a3dbce2455b0f1a1f779c5abc654490a7bc6daadb5027016bce758220) |
| `routes.dont_send` | [routes.dont_send](data-sources--alert_policy--reference--group-001.md#canonical-1eacde35146aa92d6d0a3f272ad956d88f83f4cacdd06c4f0e25c10d7de64212) |
| `routes.group` | [routes.group](data-sources--alert_policy--reference--group-001.md#canonical-c24d54a846fcb5e4f064b1152a6bb05e0ec22b6965f555eb23a4095e80cb26d1) |
| `routes.group.groups` | [routes.group.groups](data-sources--alert_policy--reference--group-001.md#canonical-402a9e136a022250da56b588572c1c61dc05415a68c1eaf57b3800749f571a3c) |
| `routes.notification_parameters` | [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-13ff4f822130551c9821faf8ad1f336a3d947da8f90127ba85e3ed6b9674d942) |
| `routes.notification_parameters.custom` | [routes.notification_parameters.custom](data-sources--alert_policy--reference--group-001.md#canonical-11179004ff5562d4e5f45cb873d5d6a63986f03f06ac63e4dde810badc467610) |
| `routes.notification_parameters.custom.labels` | [routes.notification_parameters.custom.labels](data-sources--alert_policy--reference--group-001.md#canonical-158c1b482461a2a955be8158f97f70576b8ebdf4d6d8f965a6e72540ba0d505a) |
| `routes.notification_parameters.default` | [routes.notification_parameters.default](data-sources--alert_policy--reference--group-001.md#canonical-cd0a8f8584b7fdbc2c4587ff01964331cde74eecf048dff2d3dd2afc25a978a7) |
| `routes.notification_parameters.group_interval` | [routes.notification_parameters.group_interval](data-sources--alert_policy--reference--group-001.md#canonical-46924f62a9934e05588c289f5a3835457ded28abb2abbd367628cafd6cbbe265) |
| `routes.notification_parameters.group_wait` | [routes.notification_parameters.group_wait](data-sources--alert_policy--reference--group-001.md#canonical-c38dab52de9da9a1fd6a39020082d53e76786057af089135f4ed4cc395f6b997) |
| `routes.notification_parameters.individual` | [routes.notification_parameters.individual](data-sources--alert_policy--reference--group-001.md#canonical-a044dd6f5f400a76b768e5723a7abd09b3498f6bccbf67b7e305bcafba42d241) |
| `routes.notification_parameters.repeat_interval` | [routes.notification_parameters.repeat_interval](data-sources--alert_policy--reference--group-001.md#canonical-80ac10c5f2b6e329ccbc1b4b63a9d8d8de3a4a5fbf004b4b754e33344b2083c1) |
| `routes.notification_parameters.ves_io_group` | [routes.notification_parameters.ves_io_group](data-sources--alert_policy--reference--group-001.md#canonical-b5682adb8d196d5ae00623e190aecbb1b447c65bddab29d07ae0f72c302ef579) |
| `routes.send` | [routes.send](data-sources--alert_policy--reference--group-001.md#canonical-3c8282c810289df69bf986a016a83c176f6e75e567af6ffbadfe9b0f76bd91c9) |
| `routes.severity` | [routes.severity](data-sources--alert_policy--reference--group-001.md#canonical-e79631ac78b1a8a9a2b9bb649a1f869ffff15894888f20b0ea8402114c062942) |
| `routes.severity.severities` | [routes.severity.severities](data-sources--alert_policy--reference--group-001.md#canonical-159d2bb7856d7032a86627b3062fe769fc11dc7fb74fe2c6c971cf25a758e473) |

<a id="canonical-919624a9546fb8f9da2754258d7c86d3214e363332c4352b192043e98ed93a47"></a>

## Next pages — Property reference / a7fe95349bdc / 11

- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-867b7a75fbb3878ec9962950b54cc863e180e5d8e2486731ca9249f53d5b3720)
- [receivers](data-sources--alert_policy--reference--group-001.md#canonical-d11555e3f8eb4b9c7838e3d9697c801749bd384c2925c41e4a0b080052fcca7d)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-867b7a75fbb3878ec9962950b54cc863e180e5d8e2486731ca9249f53d5b3720"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b0c8986cea3cd03765ef2a5d8a9e299261d4940c64767a2b25a65c7f49b2bd0"></a>

## notification_parameters — notification_parameters / 9f8dc7a3e73d / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- notification_parameters

<a id="canonical-f5f6b559e1075025538b938e56a33a1f0bf2a84b36789e3975edb2278058e6dd"></a>

Type: `"single"`. Computed.

Set of notification parameters to decide how and when the alert notifications should be sent to the
receivers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-group_by": "[\"custom\",\"default\",\"individual\",\"ves_io_group\"]"
}
```

<a id="canonical-138b72a3c39315a0e1ff0e5f302a09695f5a1fbd722d17c2166ac09652c63cb1"></a>

## Direct properties — notification_parameters / 9f8dc7a3e73d / 3

- [custom](data-sources--alert_policy--reference--group-001.md#canonical-2fe6bb55180cca69c3ac072074d8746b5112eb883961ab04b40bf512b9a79874): complete subsection reference.

- [default](data-sources--alert_policy--reference--group-001.md#canonical-137b7ff84f8eefb3615a4a673ebeaf589946cb913cd15b803ed699a91331dcd0): complete subsection reference.

<a id="canonical-d7b6cb3a9b80c309f9382796e78a13b874fc0ddc45a6732011bb79f20ad89fb4"></a>

<a id="canonical-5f83d837dc42c2c50fedcacf1e598bd8e758aeb3e016e96b91949bf348747ce1"></a>

## group_interval property — notification_parameters / 9f8dc7a3e73d / 4

Type: `"string"`. Computed.

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval..

Upstream description:

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval defaults to "1m"

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
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-30bd6c709b205ea483f53f4ffdb170aa01985b90c4392304c6b42a0d764c0b48"></a>

<a id="canonical-0bfd7fb0f4f7f4408eadb452f48fa1e6e58501e6a4440e6ca59b7e66d224ec44"></a>

## group_wait property — notification_parameters / 9f8dc7a3e73d / 5

Type: `"string"`. Computed.

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to '30s'.

Upstream description:

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to "30s"

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
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [individual](data-sources--alert_policy--reference--group-001.md#canonical-be3f8495ba90b94cd8395e2ec5b892d3584d5577048bdb5f185331a537aee2ce): complete subsection reference.

<a id="canonical-d27bb88af9dc0f600567688cee578ffc3abf09d173234c6c485e1b6a91d1246f"></a>

<a id="canonical-c83c2cf0c7fe30d5cb6f26fe2da16293f370fe3757e54ff19d74ea9eaa5601bb"></a>

## repeat_interval property — notification_parameters / 9f8dc7a3e73d / 6

Type: `"string"`. Computed.

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to '4h'.

Upstream description:

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to "4h"

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
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [ves_io_group](data-sources--alert_policy--reference--group-001.md#canonical-a683747a99b08cc24a4abb0aa8a21ba4db95f89ac31453850e2e432111930dc8): complete subsection reference.

<a id="canonical-472e9055e6a420d7cb32468dcda7e4012c1ea2d73e460b0241cd64767e50b65d"></a>

## Next pages — notification_parameters / 9f8dc7a3e73d / 7

- [notification_parameters.custom](data-sources--alert_policy--reference--group-001.md#canonical-2fe6bb55180cca69c3ac072074d8746b5112eb883961ab04b40bf512b9a79874)
- [notification_parameters.default](data-sources--alert_policy--reference--group-001.md#canonical-137b7ff84f8eefb3615a4a673ebeaf589946cb913cd15b803ed699a91331dcd0)
- [notification_parameters.individual](data-sources--alert_policy--reference--group-001.md#canonical-be3f8495ba90b94cd8395e2ec5b892d3584d5577048bdb5f185331a537aee2ce)
- [notification_parameters.ves_io_group](data-sources--alert_policy--reference--group-001.md#canonical-a683747a99b08cc24a4abb0aa8a21ba4db95f89ac31453850e2e432111930dc8)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-2fe6bb55180cca69c3ac072074d8746b5112eb883961ab04b40bf512b9a79874"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63d4349bdd86497f3fd03babb4a5596bd766cee14e66e364f4ef12964887d9de"></a>

## notification_parameters.custom — notification_parameters.custom / 182ba17ef695 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-867b7a75fbb3878ec9962950b54cc863e180e5d8e2486731ca9249f53d5b3720)
- notification_parameters.custom

<a id="canonical-ec9a54faae87c31323a0dee3101072b95fe11e7f691097a98ceffebe4f49ec0c"></a>

Type: `"single"`. Computed.

Specify list of custom labels to group/aggregate the alerts.

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

<a id="canonical-6f04945b744beda111070717c9d9498509cb27ed04299c539d474388f67ecbae"></a>

## Direct properties — notification_parameters.custom / 182ba17ef695 / 3

<a id="canonical-347a200130c7614f8983db93bb9cad505a568b8a02ac53199976c1ca005e805b"></a>

<a id="canonical-a1aa11a881e751c7c8156975b8ea49312e5305b1cad65c5604b0415fc0b4fe0b"></a>

## labels property — notification_parameters.custom / 182ba17ef695 / 4

Type: `["list", "string"]`. Computed.

Name of labels to group/aggregate the alerts.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3980ba39dac7e1c66a079c8a7d951bea61e82b4ddd5470de9733a12b556bfea5"></a>

## Next pages — notification_parameters.custom / 182ba17ef695 / 5

- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-867b7a75fbb3878ec9962950b54cc863e180e5d8e2486731ca9249f53d5b3720)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-137b7ff84f8eefb3615a4a673ebeaf589946cb913cd15b803ed699a91331dcd0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7711beeab7f9bba2181e2f4a27405ac04cd6fb0c4b85946dcdd3b2b6af776ad0"></a>

## notification_parameters.default — notification_parameters.default / e30e67469401 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-867b7a75fbb3878ec9962950b54cc863e180e5d8e2486731ca9249f53d5b3720)
- notification_parameters.default

<a id="canonical-8b4e0aa6dba90446cc209bf47b8360470da4d54097b6d2b7779fef49ded06051"></a>

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

<a id="canonical-989525fae12fe8634f1ed5813ab854d533aa1a346e58b49e7e9b6bcc07e00126"></a>

## Direct properties — notification_parameters.default / e30e67469401 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cf10de06885ef64758fb5e5e0033024d9285ab4eab76f5fb3745e9af31846068"></a>

## Next pages — notification_parameters.default / e30e67469401 / 4

- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-867b7a75fbb3878ec9962950b54cc863e180e5d8e2486731ca9249f53d5b3720)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-be3f8495ba90b94cd8395e2ec5b892d3584d5577048bdb5f185331a537aee2ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23931c09664d2863401f79430a8ce1957e217e001d0529b15865090110d8726a"></a>

## notification_parameters.individual — notification_parameters.individual / 6c62363f999a / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-867b7a75fbb3878ec9962950b54cc863e180e5d8e2486731ca9249f53d5b3720)
- notification_parameters.individual

<a id="canonical-391165b97eacb0753172bbb2fba3f93f4d9b386aec004bd096894377a8cb396f"></a>

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

<a id="canonical-f7d2768e9defaccdc2f3de2a0672242f664e18a8ba2f1320986e4940eaad4200"></a>

## Direct properties — notification_parameters.individual / 6c62363f999a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a34b3da97cb995a617e33f4c7f47bab88a306713005cc9688ccd7d3f808647f1"></a>

## Next pages — notification_parameters.individual / 6c62363f999a / 4

- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-867b7a75fbb3878ec9962950b54cc863e180e5d8e2486731ca9249f53d5b3720)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-a683747a99b08cc24a4abb0aa8a21ba4db95f89ac31453850e2e432111930dc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6a13634c771143956432f485705a23bce3e213cff312a16741e5b449772fad4"></a>

## notification_parameters.ves_io_group — notification_parameters.ves_io_group / f228a8d5120a / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-867b7a75fbb3878ec9962950b54cc863e180e5d8e2486731ca9249f53d5b3720)
- notification_parameters.ves_io_group

<a id="canonical-2d1971c9d08e019a847c966507a87330a7a23b8b12788d18c41c72643076cff2"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ves io group.

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

<a id="canonical-8e9f6115c90a139188a4020beb69df49b56c1436a42fe70195c9322b93e720a0"></a>

## Direct properties — notification_parameters.ves_io_group / f228a8d5120a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ebc67d43efc6988a6000f5ec71add153f3a025371fc56aa01d40960d76a1678b"></a>

## Next pages — notification_parameters.ves_io_group / f228a8d5120a / 4

- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-867b7a75fbb3878ec9962950b54cc863e180e5d8e2486731ca9249f53d5b3720)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-d11555e3f8eb4b9c7838e3d9697c801749bd384c2925c41e4a0b080052fcca7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f35451c9c4fdb279eb31394dae2a64a58d144e06fe74752d0414609e95efd02e"></a>

## receivers — receivers / 5c151b62bb46 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- receivers

<a id="canonical-1d30c1eb7244a9ef78d1de0d3e5a3ad2d31f50cf96efd4dc7319ce67b87ba5c3"></a>

Type: `"list"`. Computed.

List of Alert Receivers where the alerts will be sent.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-1c2ba1df762ca7086204737baa0e1371108da5d0c5035f9fd7db22e0c9420f8c"></a>

## Direct properties — receivers / 5c151b62bb46 / 3

<a id="canonical-02130ac6fd48fe8080954ae38d36e8c0011c480ff4374193fac221fbda5285cb"></a>

<a id="canonical-8a4810258f5d45f328896d10ed973282ef977c4cea4b5db3ca4be7bcb7de3515"></a>

## kind property — receivers / 5c151b62bb46 / 4

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

<a id="canonical-caad5e4fae398841c61c7742d131d00c6c5d1e5d53605621734144bf182491c3"></a>

<a id="canonical-98dadfb0e3e895fe7433b105e605dc9549ed98493c9e1b442c6368c74a6750f9"></a>

## name property — receivers / 5c151b62bb46 / 5

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

<a id="canonical-b0a683b18c1d336fea6f5b4ac4648f2e93d4af0f84178744113b6c7eed7d4bf4"></a>

<a id="canonical-a09f3c16638e803585e8b3bce1197e754cd1c92306a477aec3f1ba6dd6fb6a94"></a>

## namespace property — receivers / 5c151b62bb46 / 6

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

<a id="canonical-f5c5ae148ea84c5cf07d07dc1632754f0b44bc3e1d72b95447283cc670d476d5"></a>

<a id="canonical-31d220ac6ef8db5e294b083611c2f7c1e1fca3f92f52a7050f339cf522ce1777"></a>

## tenant property — receivers / 5c151b62bb46 / 7

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

<a id="canonical-ef43f2f10587d89a3c94c67a9ad7b2c6e235aaec3f90305e4542ed5e0520340e"></a>

<a id="canonical-a68d917139f21aecce331dff9528109601a720aeaf954904cd56ee38277a4644"></a>

## uid property — receivers / 5c151b62bb46 / 8

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

<a id="canonical-005fa5ef10d29376866911d5bc59a0a2fc544fd106f3ee4bc8a000e8f6adc9b0"></a>

## Next pages — receivers / 5c151b62bb46 / 9

- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e37a9927a000c40ad80985ddd945e0b362e67e1fa591192ca3a838d22a4714f"></a>

## routes — routes / eb5f4940bfd2 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- routes

<a id="canonical-523d7ddfdf198f4ae4cd98d3bb5a37953414ed24e44786bb38d6ef13b9c31de0"></a>

Type: `"list"`. Computed.

Set of routes to match the incoming alert. The routes are evaluated in the specified order and
terminates on the first match.

Upstream description:

Set of routes to match the incoming alert. The routes are evaluated in the specified order and
terminates on the first match.

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

<a id="canonical-f01e15ca1a85c1d8694b822a5acde6884d0244b94b6ea3344b4d61f1af1ff3ea"></a>

## Direct properties — routes / eb5f4940bfd2 / 3

<a id="canonical-124468f69d1b3ae53b9c589198d33cd28324e85936c57c17a56768653676c146"></a>

<a id="canonical-61ee7b9d56bdde3cea2058d283479a3e822cd7f46b32b881dfb58852a4910b25"></a>

## alertname property — routes / eb5f4940bfd2 / 4

Type: `"string"`. Computed.

\[Enum:
SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN|SITE\_PHYSICAL\_INTERFACE\_DOWN|TUNNELS\_TO\_CUSTOMER\_SITE\_DOWN|SERVICE\_SERVER\_ERROR|SERVICE\_CLIENT\_ERROR|SERVICE\_HEALTH\_LOW|SERVICE\_UNAVAILABLE|SERVICE\_SERVER\_ERROR\_PER\_SOURCE\_SITE|SERVICE\_CLIENT\_ERROR\_PER\_SOURCE\_SITE|SERVICE\_ENDPOINT\_HEALTHCHECK\_FAILURE|SYNTHETIC\_MONITOR\_HEALTH\_CRITICAL|MALICIOUS\_USER\_DETECTED|WAF\_TOO\_MANY\_ATTACKS|API\_SECURITY\_TOO\_MANY\_ATTACKS|SERVICE\_POLICY\_TOO\_MANY\_ATTACKS|WAF\_TOO\_MANY\_MALICIOUS\_BOTS|BOT\_DEFENSE\_TOO\_MANY\_SECURITY\_EVENTS|THREAT\_CAMPAIGN|VES\_CLIENT\_SIDE\_DEFENSE\_SUSPICIOUS\_DOMAIN|VES\_CLIENT\_SIDE\_DEFENSE\_SENSITIVE\_FIELD\_READ|TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_FAILURE|TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_STILL\_FAILING|TLS\_AUTOMATIC\_CERTIFICATE\_EXPIRED|TLS\_CUSTOM\_CERTIFICATE\_EXPIRING|TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\_SOON|TLS\_CUSTOM\_CERTIFICATE\_EXPIRED|L7DDOS|DNS\_ZONE\_IGNORED\_DUPLICATE\_RECORD|API\_SECURITY\_UNUSED\_API\_DETECTED|API\_SECURITY\_SHADOW\_API\_DETECTED|API\_SECURITY\_SENSITIVE\_DATA\_IN\_RESPONSE\_DETECTED|API\_SECURITY\_RISK\_SCORE\_HIGH\_DETECTED|ROUTED\_DDOS\_ALERT\_NOTIFICATION|ROUTED\_DDOS\_MITIGATION\_NOTIFICATION|ROUTED\_DDOS\_TUNNEL\_STATUS\_UPDATE\_NOTIFICATION|L7\_DDOS\_AUTO\_MITIGATION\]
List of Alert Names Customer tunnel interface down Physical Interface down Tunnel Interfaces to
Customer Site Down Virtual Host server error Virtual Host client error Service Health Low Service
Unavailable Virtual Host server error Virtual Host client error Endpoint Healthcheck failure
Synthetic.. Possible values are \`SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN\`,
\`SITE\_PHYSICAL\_INTERFACE\_DOWN\`, \`TUNNELS\_TO\_CUSTOMER\_SITE\_DOWN\`,
\`SERVICE\_SERVER\_ERROR\`, \`SERVICE\_CLIENT\_ERROR\`, \`SERVICE\_HEALTH\_LOW\`,
\`SERVICE\_UNAVAILABLE\`, \`SERVICE\_SERVER\_ERROR\_PER\_SOURCE\_SITE\`,
\`SERVICE\_CLIENT\_ERROR\_PER\_SOURCE\_SITE\`, \`SERVICE\_ENDPOINT\_HEALTHCHECK\_FAILURE\`,
\`SYNTHETIC\_MONITOR\_HEALTH\_CRITICAL\`, \`MALICIOUS\_USER\_DETECTED\`,
\`WAF\_TOO\_MANY\_ATTACKS\`, \`API\_SECURITY\_TOO\_MANY\_ATTACKS\`,
\`SERVICE\_POLICY\_TOO\_MANY\_ATTACKS\`, \`WAF\_TOO\_MANY\_MALICIOUS\_BOTS\`,
\`BOT\_DEFENSE\_TOO\_MANY\_SECURITY\_EVENTS\`, \`THREAT\_CAMPAIGN\`,
\`VES\_CLIENT\_SIDE\_DEFENSE\_SUSPICIOUS\_DOMAIN\`,
\`VES\_CLIENT\_SIDE\_DEFENSE\_SENSITIVE\_FIELD\_READ\`,
\`TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_FAILURE\`,
\`TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_STILL\_FAILING\`, \`TLS\_AUTOMATIC\_CERTIFICATE\_EXPIRED\`,
\`TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\`, \`TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\_SOON\`,
\`TLS\_CUSTOM\_CERTIFICATE\_EXPIRED\`, \`L7DDOS\`, \`DNS\_ZONE\_IGNORED\_DUPLICATE\_RECORD\`,
\`API\_SECURITY\_UNUSED\_API\_DETECTED\`, \`API\_SECURITY\_SHADOW\_API\_DETECTED\`,
\`API\_SECURITY\_SENSITIVE\_DATA\_IN\_RESPONSE\_DETECTED\`,
\`API\_SECURITY\_RISK\_SCORE\_HIGH\_DETECTED\`, \`ROUTED\_DDOS\_ALERT\_NOTIFICATION\`,
\`ROUTED\_DDOS\_MITIGATION\_NOTIFICATION\`, \`ROUTED\_DDOS\_TUNNEL\_STATUS\_UPDATE\_NOTIFICATION\`,
\`L7\_DDOS\_AUTO\_MITIGATION\`. Defaults to \`SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN\`.

Upstream description:

List of Alert Names

Customer tunnel interface down Physical Interface down Tunnel Interfaces to Customer Site Down
Virtual Host server error Virtual Host client error Service Health Low Service Unavailable Virtual
Host server error Virtual Host client error Endpoint Healthcheck failure Synthetic monitor health
critical Malicious user detected Virtual Host WAF security events detected Virtual Host API security
events detected Virtual Host Service Policy security events detected Virtual Host Many Malicious
Bots based WAF security events detected Virtual Host Many Malicious Bots based Bot Defense security
events detected Virtual Host Many Threat campaign based WAF security events detected Suspicious
domain identified by Client-Side Defense service Client-Side Defense has identified a suspicious
script that is reading sensitive form field TLS Automatic Certificate renewal is failing TLS
Automatic Certificate renewal is still failing after multiple retries TLS Automatic Certificate has
expired TLS Custom Certificate will expire in less than 28 days TLS Custom Certificate will expire
in less than 15 days TLS Custom Certificate has expired DDoS security event detected DNS Zone
Ignored a Duplicate Record Create Request Unused APIs Detected Shadow APIs Detected Endpoints With
Sensitive Data In Response Detected High Risk Score Endpoints Detected A routed DDoS traffic anomaly
has been detected A routed DDoS mitigation has been implemented to block malicious traffic A routed
DDoS tunnel status has been changed L7 DDoS attack was detected, automatic mitigation is taking
place.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN",
  "enum": [
    "SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN",
    "SITE_PHYSICAL_INTERFACE_DOWN",
    "TUNNELS_TO_CUSTOMER_SITE_DOWN",
    "SERVICE_SERVER_ERROR",
    "SERVICE_CLIENT_ERROR",
    "SERVICE_HEALTH_LOW",
    "SERVICE_UNAVAILABLE",
    "SERVICE_SERVER_ERROR_PER_SOURCE_SITE",
    "SERVICE_CLIENT_ERROR_PER_SOURCE_SITE",
    "SERVICE_ENDPOINT_HEALTHCHECK_FAILURE",
    "SYNTHETIC_MONITOR_HEALTH_CRITICAL",
    "MALICIOUS_USER_DETECTED",
    "WAF_TOO_MANY_ATTACKS",
    "API_SECURITY_TOO_MANY_ATTACKS",
    "SERVICE_POLICY_TOO_MANY_ATTACKS",
    "WAF_TOO_MANY_MALICIOUS_BOTS",
    "BOT_DEFENSE_TOO_MANY_SECURITY_EVENTS",
    "THREAT_CAMPAIGN",
    "VES_CLIENT_SIDE_DEFENSE_SUSPICIOUS_DOMAIN",
    "VES_CLIENT_SIDE_DEFENSE_SENSITIVE_FIELD_READ",
    "TLS_AUTOMATIC_CERTIFICATE_RENEWAL_FAILURE",
    "TLS_AUTOMATIC_CERTIFICATE_RENEWAL_STILL_FAILING",
    "TLS_AUTOMATIC_CERTIFICATE_EXPIRED",
    "TLS_CUSTOM_CERTIFICATE_EXPIRING",
    "TLS_CUSTOM_CERTIFICATE_EXPIRING_SOON",
    "TLS_CUSTOM_CERTIFICATE_EXPIRED",
    "L7DDOS",
    "DNS_ZONE_IGNORED_DUPLICATE_RECORD",
    "API_SECURITY_UNUSED_API_DETECTED",
    "API_SECURITY_SHADOW_API_DETECTED",
    "API_SECURITY_SENSITIVE_DATA_IN_RESPONSE_DETECTED",
    "API_SECURITY_RISK_SCORE_HIGH_DETECTED",
    "ROUTED_DDOS_ALERT_NOTIFICATION",
    "ROUTED_DDOS_MITIGATION_NOTIFICATION",
    "ROUTED_DDOS_TUNNEL_STATUS_UPDATE_NOTIFICATION",
    "L7_DDOS_AUTO_MITIGATION"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6e092c1c5774d9af22791f167b9f7ee7f20b84f48b33166c0d2ea0c5d57cc151"></a>

<a id="canonical-e62035a6d37b6cc5346a870a8151adad61887bdd1b38d8901950ed8bcb942b6a"></a>

## alertname_regex property — routes / eb5f4940bfd2 / 5

Type: `"string"`. Computed.

Exclusive with \[alertname any custom group severity\] Regular Expression match for the alertname.

Upstream description:

Exclusive with \[alertname any custom group severity\] Regular Expression match for the alertname.

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

- [any](data-sources--alert_policy--reference--group-001.md#canonical-5571095aea92e322a9ffb5b559857230b9d46d90aebd2c937eef2e0cbf94d50e): complete subsection reference.

- [custom](data-sources--alert_policy--reference--group-001.md#canonical-953aacab2f17a9375b898defbc7396fc2d72660fa72abab271e53f1cc243f440): complete subsection reference.

- [dont_send](data-sources--alert_policy--reference--group-001.md#canonical-e7b7abe0dafd57dd7def3d6a62ff7e2bf1cdf89f0f76ff8cc28fb775bbd323d5): complete subsection reference.

- [group](data-sources--alert_policy--reference--group-001.md#canonical-afad61c5790fab12310639ff1a02a894d9ae911eb1f8ecf61427df924cb575a3): complete subsection reference.

- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-eabf5e120ceb74004bb08ed5cb926c8cc2428ed015a47c872b49808d2f622d3a): complete subsection reference.

- [send](data-sources--alert_policy--reference--group-001.md#canonical-b1ce00c62a906c0e6acf562e1d621304a84fc90ccffb2cec92886d5b13b5a982): complete subsection reference.

- [severity](data-sources--alert_policy--reference--group-001.md#canonical-0b2611d48e8c35d83c00fcf77be1b1cc5c1ff28868ad4e5e2eac60799acbc735): complete subsection reference.

<a id="canonical-70b34a37eba5d297ada442aa5149bd333a6b408c40b987f3a0fe2dc1f1fac94a"></a>

## Next pages — routes / eb5f4940bfd2 / 6

- [routes.any](data-sources--alert_policy--reference--group-001.md#canonical-5571095aea92e322a9ffb5b559857230b9d46d90aebd2c937eef2e0cbf94d50e)
- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-953aacab2f17a9375b898defbc7396fc2d72660fa72abab271e53f1cc243f440)
- [routes.dont_send](data-sources--alert_policy--reference--group-001.md#canonical-e7b7abe0dafd57dd7def3d6a62ff7e2bf1cdf89f0f76ff8cc28fb775bbd323d5)
- [routes.group](data-sources--alert_policy--reference--group-001.md#canonical-afad61c5790fab12310639ff1a02a894d9ae911eb1f8ecf61427df924cb575a3)
- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-eabf5e120ceb74004bb08ed5cb926c8cc2428ed015a47c872b49808d2f622d3a)
- [routes.send](data-sources--alert_policy--reference--group-001.md#canonical-b1ce00c62a906c0e6acf562e1d621304a84fc90ccffb2cec92886d5b13b5a982)
- [routes.severity](data-sources--alert_policy--reference--group-001.md#canonical-0b2611d48e8c35d83c00fcf77be1b1cc5c1ff28868ad4e5e2eac60799acbc735)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-5571095aea92e322a9ffb5b559857230b9d46d90aebd2c937eef2e0cbf94d50e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8bb91af7422dc3b98e6e0a7ac8c249ec82e6f86233da16def77e72431c82c9e9"></a>

## routes.any — routes.any / 7f4937978873 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- routes.any

<a id="canonical-f98dfba1d133027f615bdcf1d2dacbaa79003f09de28e998271ec9dfb4d11bbe"></a>

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

<a id="canonical-b5c6a6742878449af87f8609f02d180d76352f5ad200e8671c9a0eddaace28b7"></a>

## Direct properties — routes.any / 7f4937978873 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c2699030ff7c7224dabb0fbc79560803e4edac4f5e53e7f5d48a474220904904"></a>

## Next pages — routes.any / 7f4937978873 / 4

- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-953aacab2f17a9375b898defbc7396fc2d72660fa72abab271e53f1cc243f440"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36c5eb7478b307105ce0b92f89aa7a651bf58b3fd3b8a762d0cfd7a1e9e5185c"></a>

## routes.custom — routes.custom / 8dbf3fd492da / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- routes.custom

<a id="canonical-507d445dc3e69fe1334c574326b1e41bdc454267e1111b01b4702a667c3562d1"></a>

Type: `"single"`. Computed.

Set of matchers an alert has to fulfill to match the route.

Upstream description:

A set of matchers an alert has to fulfill to match the route.

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

<a id="canonical-8286d361589d4cfa97bc40efa61918070a9d823500a7404fa3594b9d10d1b56a"></a>

## Direct properties — routes.custom / 8dbf3fd492da / 3

- [alertlabel](data-sources--alert_policy--reference--group-001.md#canonical-1b0dba90db4ec2fe9c5e793e66b3d1e145a19bfd1b7218387767ea11c75f5d40): complete subsection reference.

- [alertname](data-sources--alert_policy--reference--group-001.md#canonical-44f4da684c4e51c5b38be3ba7998677999d3670126fc987aba8480cc798a0244): complete subsection reference.

- [group](data-sources--alert_policy--reference--group-001.md#canonical-dbf3ee26c5413fdc676479fd208f8a8e20b336c29d54de6aa5fe53f7595ad215): complete subsection reference.

- [severity](data-sources--alert_policy--reference--group-001.md#canonical-abf95e98a6ca5cfd593c60c8c638327ec1b88efff2e5ef6f28b1340aef73df08): complete subsection reference.

<a id="canonical-ac53f44581b61de53ab579725051f04bb6938c38d6dd51fc2fc7983a8dc63c7a"></a>

## Next pages — routes.custom / 8dbf3fd492da / 4

- [routes.custom.alertlabel](data-sources--alert_policy--reference--group-001.md#canonical-1b0dba90db4ec2fe9c5e793e66b3d1e145a19bfd1b7218387767ea11c75f5d40)
- [routes.custom.alertname](data-sources--alert_policy--reference--group-001.md#canonical-44f4da684c4e51c5b38be3ba7998677999d3670126fc987aba8480cc798a0244)
- [routes.custom.group](data-sources--alert_policy--reference--group-001.md#canonical-dbf3ee26c5413fdc676479fd208f8a8e20b336c29d54de6aa5fe53f7595ad215)
- [routes.custom.severity](data-sources--alert_policy--reference--group-001.md#canonical-abf95e98a6ca5cfd593c60c8c638327ec1b88efff2e5ef6f28b1340aef73df08)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-1b0dba90db4ec2fe9c5e793e66b3d1e145a19bfd1b7218387767ea11c75f5d40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-164460db36e257347cb0497b61a393223f97a1dbd59e26616f150ecff3078a6e"></a>

## routes.custom.alertlabel — routes.custom.alertlabel / 1b95362956d5 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-953aacab2f17a9375b898defbc7396fc2d72660fa72abab271e53f1cc243f440)
- routes.custom.alertlabel

<a id="canonical-30e389c821b163c0fd2530c4f08557ed2acca2a62c135dcde9e2e29d9034a371"></a>

Type: `"single"`. Computed.

AlertLabel to configure the alert policy rule.

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
    "ves.io.schema.rules.map.keys.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.map.max_pairs": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.keys.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.map.max_pairs": "3"
  }
}
```

<a id="canonical-019d0ff926bdec0f589417a0d4df8a2ccddff1f341921bc5ad6b092275177fa2"></a>

## Direct properties — routes.custom.alertlabel / 1b95362956d5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cb6884df6e78b854de63c5cf88fa40aae1b9b06b8202d8a770e7dcbdd0a07bd7"></a>

## Next pages — routes.custom.alertlabel / 1b95362956d5 / 4

- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-953aacab2f17a9375b898defbc7396fc2d72660fa72abab271e53f1cc243f440)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-44f4da684c4e51c5b38be3ba7998677999d3670126fc987aba8480cc798a0244"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44ba57ae5cd59263dbad5f7176ab5486fe09b80233280512f8b524c2d00fca05"></a>

## routes.custom.alertname — routes.custom.alertname / d9e7d3e0c0c6 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-953aacab2f17a9375b898defbc7396fc2d72660fa72abab271e53f1cc243f440)
- routes.custom.alertname

<a id="canonical-25e04921716757354012d7ba802bc266ad2916b1e75ad02e7cab8a91ab1619f2"></a>

Type: `"single"`. Computed.

Label Matcher.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-matcher_type": "[\"exact_match\",\"regex_match\"]"
}
```

<a id="canonical-3f5e3eb4c873ae58a914c07f1c66cdaf07525db69038c59d51a9588298d89e76"></a>

## Direct properties — routes.custom.alertname / d9e7d3e0c0c6 / 3

<a id="canonical-2c0d7d1abf0875da23cabe4cd4fb3e5e90f922266508a7f5f275a3154cdfb63a"></a>

<a id="canonical-a8d88184260cdba268dbc47a68239a461f1c4de624a7861d83bc742468c9dfa0"></a>

## exact_match property — routes.custom.alertname / d9e7d3e0c0c6 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_match\] Equality match value for the label.

Upstream description:

Exclusive with \[regex\_match\] Equality match value for the label.

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

<a id="canonical-b88acb351baabcaa3bb014fb57ea26ff98558216d412fef5238534305b3fb5ed"></a>

<a id="canonical-a7e8651132f14bbfd857852e22fa6e10aea719c8dfc35b22f8a23320fb6d7422"></a>

## regex_match property — routes.custom.alertname / d9e7d3e0c0c6 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_match\] Regular expression match value for the label.

Upstream description:

Exclusive with \[exact\_match\] Regular expression match value for the label.

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

<a id="canonical-a600c3cd1388b712170ae38a0e94d6ee5abd763c3eca37bf1321fac5e96a1b5f"></a>

## Next pages — routes.custom.alertname / d9e7d3e0c0c6 / 6

- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-953aacab2f17a9375b898defbc7396fc2d72660fa72abab271e53f1cc243f440)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-dbf3ee26c5413fdc676479fd208f8a8e20b336c29d54de6aa5fe53f7595ad215"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85fc192fedbb3b384e707adec44ca901e16295f5d81c4e6d01879a0b17547731"></a>

## routes.custom.group — routes.custom.group / 0d29a7154d94 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-953aacab2f17a9375b898defbc7396fc2d72660fa72abab271e53f1cc243f440)
- routes.custom.group

<a id="canonical-6afed69dc549049fc9998472c79f3473afb71ab6f8a746ec492998b76b7d010f"></a>

Type: `"single"`. Computed.

Label Matcher.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-matcher_type": "[\"exact_match\",\"regex_match\"]"
}
```

<a id="canonical-730603b1fa04423d3cab97cd67f5186173db9c08d8bb8238e106b314da6b8775"></a>

## Direct properties — routes.custom.group / 0d29a7154d94 / 3

<a id="canonical-01fc7c08346b685dbc9ca687e262152a36fc082e3474f8f23788c3855ee7fcc2"></a>

<a id="canonical-29444be0793c2ab735c66607e284cd470baae480d75bf0616f6ebcd9aa6311bb"></a>

## exact_match property — routes.custom.group / 0d29a7154d94 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_match\] Equality match value for the label.

Upstream description:

Exclusive with \[regex\_match\] Equality match value for the label.

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

<a id="canonical-ae8087a742a160ec29c2b9675771c24debdc0f70648ba7ec2d029b9c1be96c86"></a>

<a id="canonical-9dce6b53f47d175d76ff0baeb6abbd631a0af248629d5cbec5b2653e5d1e89da"></a>

## regex_match property — routes.custom.group / 0d29a7154d94 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_match\] Regular expression match value for the label.

Upstream description:

Exclusive with \[exact\_match\] Regular expression match value for the label.

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

<a id="canonical-15260fc4f494350bc50592a0c163e4442a20b34d218f131a12a10273b3eba274"></a>

## Next pages — routes.custom.group / 0d29a7154d94 / 6

- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-953aacab2f17a9375b898defbc7396fc2d72660fa72abab271e53f1cc243f440)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-abf95e98a6ca5cfd593c60c8c638327ec1b88efff2e5ef6f28b1340aef73df08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8d6ac55be02d712bf8e1374fb9b63a1e17dc28caefbd826a84cb0f3333f118c"></a>

## routes.custom.severity — routes.custom.severity / 09c9561c8bba / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-953aacab2f17a9375b898defbc7396fc2d72660fa72abab271e53f1cc243f440)
- routes.custom.severity

<a id="canonical-301f805afb6ca47249a3257993be9c8c349fd226d03c5a00e559f7d2925655e5"></a>

Type: `"single"`. Computed.

Label Matcher.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-matcher_type": "[\"exact_match\",\"regex_match\"]"
}
```

<a id="canonical-9e616662e98ab5fb29159899ee69e3ce5518a1fb7abc2648a7a200d353758427"></a>

## Direct properties — routes.custom.severity / 09c9561c8bba / 3

<a id="canonical-116fcb02b0eb17e6dec907137c5f1e3537a6951db99a44ae4c754ef311607a46"></a>

<a id="canonical-c1749ae7e5d9966fcd45573b7640368596888073461b5fc0fd8a93affb33a8d5"></a>

## exact_match property — routes.custom.severity / 09c9561c8bba / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_match\] Equality match value for the label.

Upstream description:

Exclusive with \[regex\_match\] Equality match value for the label.

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

<a id="canonical-f511b16a3dbce2455b0f1a1f779c5abc654490a7bc6daadb5027016bce758220"></a>

<a id="canonical-2fdefea60ab4dceee32fd611a631aac359dfcb19c5b303e380f492da82f9a205"></a>

## regex_match property — routes.custom.severity / 09c9561c8bba / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_match\] Regular expression match value for the label.

Upstream description:

Exclusive with \[exact\_match\] Regular expression match value for the label.

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

<a id="canonical-afde83ae2dae196124cc58bb90ff8099e879316cee5c32514ed5ae0d1498bfce"></a>

## Next pages — routes.custom.severity / 09c9561c8bba / 6

- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-953aacab2f17a9375b898defbc7396fc2d72660fa72abab271e53f1cc243f440)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-e7b7abe0dafd57dd7def3d6a62ff7e2bf1cdf89f0f76ff8cc28fb775bbd323d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ef814f4a88bb05367e641c1ad1b7ad1b31bc8d87e93da5efd2b1c8518d6602a"></a>

## routes.dont_send — routes.dont_send / 2954d8d57625 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- routes.dont_send

<a id="canonical-1eacde35146aa92d6d0a3f272ad956d88f83f4cacdd06c4f0e25c10d7de64212"></a>

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

<a id="canonical-26942c1ef482b8f5b881e5439d3386d07cb8553d57ad7e488053a79f59945742"></a>

## Direct properties — routes.dont_send / 2954d8d57625 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e1d5586b4518eaf77023522ddf2f4ce5cfcc1de25d7e6c070ec5747e1ea4d4fd"></a>

## Next pages — routes.dont_send / 2954d8d57625 / 4

- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-afad61c5790fab12310639ff1a02a894d9ae911eb1f8ecf61427df924cb575a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e083b9aee7986e68ce1e3a0fdf3cdf04c9fe78305b37854ceb407023708a77f7"></a>

## routes.group — routes.group / 9536690d5d31 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- routes.group

<a id="canonical-c24d54a846fcb5e4f064b1152a6bb05e0ec22b6965f555eb23a4095e80cb26d1"></a>

Type: `"single"`. Computed.

Select one or more known group names to match the incoming alert.

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

<a id="canonical-650bb033dbea8b13a313d0fb1575ce4e58589039ed3ff2a04eaf60a3b3d82f1d"></a>

## Direct properties — routes.group / 9536690d5d31 / 3

<a id="canonical-402a9e136a022250da56b588572c1c61dc05415a68c1eaf57b3800749f571a3c"></a>

<a id="canonical-43b6d1b67266c60b31d2c50d849531b5580fb418d5d0be349014d20ea2083e4d"></a>

## groups property — routes.group / 9536690d5d31 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
INFRASTRUCTURE|IAAS\_CAAS|VIRTUAL\_HOST|VOLT\_SHARE|UAM|SECURITY|TIMESERIES\_ANOMALY|SHAPE\_SECURITY|SECURITY\_CSD|CDN|SYNTHETIC\_MONITORS|TLS|SECURITY\_BOT\_DEFENSE|CLOUD\_LINK|DNS|ROUTED\_DDOS\]
Groups. Name of groups to match the alert. Possible values are \`INFRASTRUCTURE\`, \`IAAS\_CAAS\`,
\`VIRTUAL\_HOST\`, \`VOLT\_SHARE\`, \`UAM\`, \`SECURITY\`, \`TIMESERIES\_ANOMALY\`,
\`SHAPE\_SECURITY\`, \`SECURITY\_CSD\`, \`CDN\`, \`SYNTHETIC\_MONITORS\`, \`TLS\`,
\`SECURITY\_BOT\_DEFENSE\`, \`CLOUD\_LINK\`, \`DNS\`, \`ROUTED\_DDOS\`. Defaults to
\`INFRASTRUCTURE\`.

Upstream description:

Name of groups to match the alert.

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

<a id="canonical-69c362f39d72e60c17356bb4dda6a21aa4c9baf3d86c188c85833cfff7215a3f"></a>

## Next pages — routes.group / 9536690d5d31 / 5

- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-eabf5e120ceb74004bb08ed5cb926c8cc2428ed015a47c872b49808d2f622d3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66293d52a272b4eb4f609d80e274212d077f262b326b2d30ca4a850025e9250a"></a>

## routes.notification_parameters — routes.notification_parameters / a2256fc9c185 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- routes.notification_parameters

<a id="canonical-13ff4f822130551c9821faf8ad1f336a3d947da8f90127ba85e3ed6b9674d942"></a>

Type: `"single"`. Computed.

Set of notification parameters to decide how and when the alert notifications should be sent to the
receivers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-group_by": "[\"custom\",\"default\",\"individual\",\"ves_io_group\"]"
}
```

<a id="canonical-71741e91889d2b22f1bf0e90a9fb7998d3b2dcef7bfb3ee463cbf4170b905711"></a>

## Direct properties — routes.notification_parameters / a2256fc9c185 / 3

- [custom](data-sources--alert_policy--reference--group-001.md#canonical-44b52d220ea9f39780e572773dee83e6e27e7bc23ada8a6c4eead28ea18eea1c): complete subsection reference.

- [default](data-sources--alert_policy--reference--group-001.md#canonical-8cab2da679ceaf128f704818c99d100ac0b4d4c7cbf05257a896a62fbed49116): complete subsection reference.

<a id="canonical-46924f62a9934e05588c289f5a3835457ded28abb2abbd367628cafd6cbbe265"></a>

<a id="canonical-b487cfe36837a11c51a78c0a81b12f8a6d22c1346ae41744acaea2fb409fbfb8"></a>

## group_interval property — routes.notification_parameters / a2256fc9c185 / 4

Type: `"string"`. Computed.

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval..

Upstream description:

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval defaults to "1m"

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
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-c38dab52de9da9a1fd6a39020082d53e76786057af089135f4ed4cc395f6b997"></a>

<a id="canonical-1d2784a998dab56c544666f5d7e9082e84b3be27a127e65456a323854eeb3f0e"></a>

## group_wait property — routes.notification_parameters / a2256fc9c185 / 5

Type: `"string"`. Computed.

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to '30s'.

Upstream description:

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to "30s"

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
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [individual](data-sources--alert_policy--reference--group-001.md#canonical-d6e56be4f591caf5dbe5b7f0bae9e695f59eba598f422591f756d370caf460d0): complete subsection reference.

<a id="canonical-80ac10c5f2b6e329ccbc1b4b63a9d8d8de3a4a5fbf004b4b754e33344b2083c1"></a>

<a id="canonical-f9c9ce444779a9d213c47ab1142f67fd67a11f8dbe72e7949179c58e38136c95"></a>

## repeat_interval property — routes.notification_parameters / a2256fc9c185 / 6

Type: `"string"`. Computed.

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to '4h'.

Upstream description:

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to "4h"

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
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [ves_io_group](data-sources--alert_policy--reference--group-001.md#canonical-4c30746b90d05cdc15363a806576f14266ad9d210be7717a46fcee43f0b15801): complete subsection reference.

<a id="canonical-95412a1af0e7c2588e7e49cc5e599343ce3223b55865a72c268851845d819817"></a>

## Next pages — routes.notification_parameters / a2256fc9c185 / 7

- [routes.notification_parameters.custom](data-sources--alert_policy--reference--group-001.md#canonical-44b52d220ea9f39780e572773dee83e6e27e7bc23ada8a6c4eead28ea18eea1c)
- [routes.notification_parameters.default](data-sources--alert_policy--reference--group-001.md#canonical-8cab2da679ceaf128f704818c99d100ac0b4d4c7cbf05257a896a62fbed49116)
- [routes.notification_parameters.individual](data-sources--alert_policy--reference--group-001.md#canonical-d6e56be4f591caf5dbe5b7f0bae9e695f59eba598f422591f756d370caf460d0)
- [routes.notification_parameters.ves_io_group](data-sources--alert_policy--reference--group-001.md#canonical-4c30746b90d05cdc15363a806576f14266ad9d210be7717a46fcee43f0b15801)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-44b52d220ea9f39780e572773dee83e6e27e7bc23ada8a6c4eead28ea18eea1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfa17207cad7941ffc00fd585abf3d3c8f2dde21ee7332021b292d5680b3a2db"></a>

## routes.notification_parameters.custom — routes.notification_parameters.custom / cb3e823d0e5e / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-eabf5e120ceb74004bb08ed5cb926c8cc2428ed015a47c872b49808d2f622d3a)
- routes.notification_parameters.custom

<a id="canonical-11179004ff5562d4e5f45cb873d5d6a63986f03f06ac63e4dde810badc467610"></a>

Type: `"single"`. Computed.

Specify list of custom labels to group/aggregate the alerts.

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

<a id="canonical-67fb3c1b7d4b4075e78de39438bb420f6f4ec65a306d995d6cbffa86c08a8ba1"></a>

## Direct properties — routes.notification_parameters.custom / cb3e823d0e5e / 3

<a id="canonical-158c1b482461a2a955be8158f97f70576b8ebdf4d6d8f965a6e72540ba0d505a"></a>

<a id="canonical-281aae866d1d882b4fb2a39813f99704723e51871f934e4806b584434d2097e7"></a>

## labels property — routes.notification_parameters.custom / cb3e823d0e5e / 4

Type: `["list", "string"]`. Computed.

Name of labels to group/aggregate the alerts.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c5b981fb0d0959d96a344f962c351673d5cee01e93da3bded2e6c13a79616add"></a>

## Next pages — routes.notification_parameters.custom / cb3e823d0e5e / 5

- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-eabf5e120ceb74004bb08ed5cb926c8cc2428ed015a47c872b49808d2f622d3a)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-8cab2da679ceaf128f704818c99d100ac0b4d4c7cbf05257a896a62fbed49116"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90e8289834946e65592d9496644d94722275e32bc4c8a9512eac07076bbb6f50"></a>

## routes.notification_parameters.default — routes.notification_parameters.default / 2bb94ac108aa / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-eabf5e120ceb74004bb08ed5cb926c8cc2428ed015a47c872b49808d2f622d3a)
- routes.notification_parameters.default

<a id="canonical-cd0a8f8584b7fdbc2c4587ff01964331cde74eecf048dff2d3dd2afc25a978a7"></a>

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

<a id="canonical-404a16b3b43eb4c25ed9781cf72ae76bc80f14f08447029b2729a20655d95ad3"></a>

## Direct properties — routes.notification_parameters.default / 2bb94ac108aa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d69bf6c9fd8af35da6469a73730abbeac4ba3e3ae72f1ad20c5ad814275046cd"></a>

## Next pages — routes.notification_parameters.default / 2bb94ac108aa / 4

- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-eabf5e120ceb74004bb08ed5cb926c8cc2428ed015a47c872b49808d2f622d3a)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-d6e56be4f591caf5dbe5b7f0bae9e695f59eba598f422591f756d370caf460d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb5434729d2ab29a2edd9f52e180de87eb931c05648060a9064fe757a8ee9de1"></a>

## routes.notification_parameters.individual — routes.notification_parameters.individual / 90b490156ea1 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-eabf5e120ceb74004bb08ed5cb926c8cc2428ed015a47c872b49808d2f622d3a)
- routes.notification_parameters.individual

<a id="canonical-a044dd6f5f400a76b768e5723a7abd09b3498f6bccbf67b7e305bcafba42d241"></a>

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

<a id="canonical-cdeb8f3ede8ad5bbd7be6cc8f02d8ea121a0ace3c0ac16ad0b35c85d93a08db1"></a>

## Direct properties — routes.notification_parameters.individual / 90b490156ea1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a4d63a33b34e21466f8daff732ca946fc6685feb81fa3b7b11331c6b308d4d25"></a>

## Next pages — routes.notification_parameters.individual / 90b490156ea1 / 4

- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-eabf5e120ceb74004bb08ed5cb926c8cc2428ed015a47c872b49808d2f622d3a)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-4c30746b90d05cdc15363a806576f14266ad9d210be7717a46fcee43f0b15801"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c05f2b60b465ebeeff82bbae139ed80ac9dff8f720d8d38bf98709f2b32c1a34"></a>

## routes.notification_parameters.ves_io_group — routes.notification_parameters.ves_io_group / 8811ad680399 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-eabf5e120ceb74004bb08ed5cb926c8cc2428ed015a47c872b49808d2f622d3a)
- routes.notification_parameters.ves_io_group

<a id="canonical-b5682adb8d196d5ae00623e190aecbb1b447c65bddab29d07ae0f72c302ef579"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ves io group.

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

<a id="canonical-c42e6c38dd4326efecba6ea2df5c053c0f99f4ee5cb4649cb5b014f5ee7e7389"></a>

## Direct properties — routes.notification_parameters.ves_io_group / 8811ad680399 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cca810dd8cf0c2a19fbd3e73a7ca04ec18370db993788ce1545773d5db52a84c"></a>

## Next pages — routes.notification_parameters.ves_io_group / 8811ad680399 / 4

- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-eabf5e120ceb74004bb08ed5cb926c8cc2428ed015a47c872b49808d2f622d3a)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-b1ce00c62a906c0e6acf562e1d621304a84fc90ccffb2cec92886d5b13b5a982"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63fdf95ca5ef14daf0e38665256f1b221afcb8d8e2c402828a4984a68db1f192"></a>

## routes.send — routes.send / 31e704a57045 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- routes.send

<a id="canonical-3c8282c810289df69bf986a016a83c176f6e75e567af6ffbadfe9b0f76bd91c9"></a>

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

<a id="canonical-ad5b41dddedb47f25ce4f48400484d807380dbe4f87cf10d36c6094b7b14aef5"></a>

## Direct properties — routes.send / 31e704a57045 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1528ea561d2e46e6c51b0ee9d493ca9069156204c11ab9877efc710e95cd1752"></a>

## Next pages — routes.send / 31e704a57045 / 4

- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-0b2611d48e8c35d83c00fcf77be1b1cc5c1ff28868ad4e5e2eac60799acbc735"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7a815c5d55b597edcc46b8f71c412d7034d31361910fd6ea66177711f9ddbf3"></a>

## routes.severity — routes.severity / d09613b833d5 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- routes.severity

<a id="canonical-e79631ac78b1a8a9a2b9bb649a1f869ffff15894888f20b0ea8402114c062942"></a>

Type: `"single"`. Computed.

Select one or more severity levels to match the incoming alert.

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

<a id="canonical-8c5b9bba2e7b749d04a0fbdedcec2f05c007ec573baea8d4cb61b426f3d8726d"></a>

## Direct properties — routes.severity / d09613b833d5 / 3

<a id="canonical-159d2bb7856d7032a86627b3062fe769fc11dc7fb74fe2c6c971cf25a758e473"></a>

<a id="canonical-5170903d55d94518e8a99abd17f4e9bbfa75c35048dadfe98a891db14d1e9fb7"></a>

## severities property — routes.severity / d09613b833d5 / 4

Type: `["list", "string"]`. Computed.

\[Enum: MINOR|MAJOR|CRITICAL\] Severities. List of severity levels. Possible values are \`MINOR\`,
\`MAJOR\`, \`CRITICAL\`. Defaults to \`MINOR\`.

Upstream description:

List of severity levels.

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

<a id="canonical-62b575677254cc29616267d1f457d8b69c468a73bfba059941d6fa2d2039b89f"></a>

## Next pages — routes.severity / d09613b833d5 / 5

- [routes](data-sources--alert_policy--reference--group-001.md#canonical-49b3e35ca71ce22357dda1eeec9886a08447a9a2bd7d715297722d18839cb4e2)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
