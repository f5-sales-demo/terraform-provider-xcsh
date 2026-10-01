---
page_title: "xcsh_dns_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy reference."
---

# xcsh_dns_proxy reference

<a id="canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26c56de0ea538e405486277d0d5e9ee43efe30cac758f84a2efc9a5624eadce3"></a>

## Property reference — Property reference / 4ab100afcb95 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- Property reference

<a id="canonical-7ec9799cee360da0975bc16dbc7114b160459847735afc8c795caadcb2a10a08"></a>

## Direct properties — Property reference / 4ab100afcb95 / 3

<a id="canonical-09bdb53350a0060f672c3635a117270ae8e2c88c3d9ec53fcacf1100173bd98e"></a>

<a id="canonical-d67b08c3d186c3cc367f1c14ebe6647da88d877ba9b5f4f7f32155f9fb71d783"></a>

## annotations property — Property reference / 4ab100afcb95 / 4

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

- [cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-9b8331703e912880af23b8d21216c343c0f6f0c28f256e527f682650226f0a9b): complete subsection reference.

- [ddos_profile](data-sources--dns_proxy--reference--group-001.md#canonical-2e1cb6af469fb44b81586f0cc731388b8574518c894a2e18104e9a1c380203d0): complete subsection reference.

<a id="canonical-8cc0e9a4c66890b02e95cab87ac4b6347d875170efa1f28147c16eb062ca669f"></a>

<a id="canonical-1d6f99cafa5543e19791d66cdac937bb826b0aa263e5cae42fb452a8eb209ca5"></a>

## description property — Property reference / 4ab100afcb95 / 5

Type: `"string"`. Computed.

Description of the DNSProxy.

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

<a id="canonical-a3cf2b9f7987a57698bf4f4c69b78aad98a4e6fc7497134e1e615c72016e6c9c"></a>

<a id="canonical-bbe1848bc58d32bfa73465f142ecfda5f8a96f01adec3ff9a3ff844af43b9008"></a>

## id property — Property reference / 4ab100afcb95 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](data-sources--dns_proxy--reference--group-001.md#canonical-7b1150e1330b61c4d3436316083ae76e767a999d7648148beb7acbff7654f112): complete subsection reference.

<a id="canonical-4f89679840860a5582677d669408562e1284de8adb18f448238e1c99793cc6d4"></a>

<a id="canonical-e341fd4b8d9affa86421ad782786125d412e86c39f77a326cf339c2705419058"></a>

## labels property — Property reference / 4ab100afcb95 / 7

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

- [lb_algorithm](data-sources--dns_proxy--reference--group-001.md#canonical-879501040bd92674fb08f9bdd32caaa6247162e032ad593e239557236f049bc7): complete subsection reference.

<a id="canonical-3d32beee1d657702c5122bf325a90f9144d6811116bb3eb2c2da6aa5306409fc"></a>

<a id="canonical-bd6c2bdda2d3e11d2a38a90f47329d28804b5010f6d54deb2194a00760da2269"></a>

## name property — Property reference / 4ab100afcb95 / 8

Type: `"string"`. Required.

Name of the DNSProxy.

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

<a id="canonical-68f49b454bebe80d382d3ea71561e877a47c7d0897b8ea107bc771afac27b32c"></a>

<a id="canonical-6505580a9d94b2fa4a08afc77233a1f33dcd556685e3dc200f7f53fab21555bb"></a>

## namespace property — Property reference / 4ab100afcb95 / 9

Type: `"string"`. Optional, Computed.

Namespace where the DNSProxy exists.

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

- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844): complete subsection reference.

- [protocol_inspection](data-sources--dns_proxy--reference--group-001.md#canonical-6ad138bdca9dd8c2037ea1d3f939f5ee70ff85ce14bff91ab9f5ebaba79d304c): complete subsection reference.

- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f): complete subsection reference.

<a id="canonical-79e209b25a748c04037bbc20aca4521f179671f836980507009e768353b36b88"></a>

<a id="canonical-f95523107cad8cec66c7aac23f6fe937c1ce15deb80213b67f37040f26bad57e"></a>

## transport_type property — Property reference / 4ab100afcb95 / 10

Type: `"string"`. Computed.

\[Enum: UDP|TCP|BothTCPAndUDP\] Transport Type - UDP: UDP - TCP: TCP - BothTCPAndUDP: Both TCP and
UDP. Possible values are \`UDP\`, \`TCP\`, \`BothTCPAndUDP\`. Defaults to \`UDP\`.

Upstream description:

Transport Type

&#8203;- UDP: UDP

&#8203;- TCP: TCP

&#8203;- BothTCPAndUDP: Both TCP and UDP.

Receipt-pinned upstream constraints:

```json
{
  "default": "UDP",
  "enum": [
    "UDP",
    "TCP",
    "BothTCPAndUDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-bb02a6661780aae33a377f68c2d04b139157d7adfc25e2c1d36a313b1ba02cda"></a>

## All schema paths — Property reference / 4ab100afcb95 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--dns_proxy--reference--group-001.md#canonical-09bdb53350a0060f672c3635a117270ae8e2c88c3d9ec53fcacf1100173bd98e) |
| `cache_profile` | [cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-f958d1f4b2c1b45f71096efa946362febc938405ae758abf8ab1817d4db80b08) |
| `cache_profile.cache_size` | [cache_profile.cache_size](data-sources--dns_proxy--reference--group-001.md#canonical-43f22c45557bd0277c63e5618b72e34b76bf905dd68e39d190de5c98bfb20673) |
| `cache_profile.disable_cache_profile` | [cache_profile.disable_cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-afdd403ae0c0b4281761aeca0e5f50fd9b6f194e526cf7dcdaacc40c8d0f37a6) |
| `ddos_profile` | [ddos_profile](data-sources--dns_proxy--reference--group-001.md#canonical-4aac37425e48c3e77ba4beb99a3d9b948a41e46e9bf31ee56264ebaa36fa4901) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](data-sources--dns_proxy--reference--group-001.md#canonical-7929f4308f1a78651e357d9c1270000f4c0a7b8786d77e9530391b4a7995f762) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](data-sources--dns_proxy--reference--group-001.md#canonical-12800f8a6a539cd135ab1abcc9f4e3da268dd3c7fb70053ee062b1230ee93138) |
| `description` | [description](data-sources--dns_proxy--reference--group-001.md#canonical-8cc0e9a4c66890b02e95cab87ac4b6347d875170efa1f28147c16eb062ca669f) |
| `id` | [id](data-sources--dns_proxy--reference--group-001.md#canonical-a3cf2b9f7987a57698bf4f4c69b78aad98a4e6fc7497134e1e615c72016e6c9c) |
| `irules` | [irules](data-sources--dns_proxy--reference--group-001.md#canonical-825788726606ffe729450ee5433fa230cfdf9a8415a82a96774ed4e37dfecf5f) |
| `irules.name` | [irules.name](data-sources--dns_proxy--reference--group-001.md#canonical-467be693ab8373b80d8192cccc56781f467980482f40989f75d8ddcfea8ccca3) |
| `irules.namespace` | [irules.namespace](data-sources--dns_proxy--reference--group-001.md#canonical-45d55157f40e3612c4b5522d9b3a8ab20cbb60ef70715d29573f2244496ef689) |
| `irules.tenant` | [irules.tenant](data-sources--dns_proxy--reference--group-001.md#canonical-4148ed888d1811512f9764b8c000516d076ca61e94948e229949a81299365e2a) |
| `labels` | [labels](data-sources--dns_proxy--reference--group-001.md#canonical-4f89679840860a5582677d669408562e1284de8adb18f448238e1c99793cc6d4) |
| `lb_algorithm` | [lb_algorithm](data-sources--dns_proxy--reference--group-001.md#canonical-4e41c5176bf46a40ba875a9d334a930bca67b714c89646024a4986c9a8534e8f) |
| `lb_algorithm.round_robin` | [lb_algorithm.round_robin](data-sources--dns_proxy--reference--group-001.md#canonical-b037dead96f53773b1e0bfad63e993e8c0efc18095f44c849a5f3de7e7dbdbb3) |
| `name` | [name](data-sources--dns_proxy--reference--group-001.md#canonical-3d32beee1d657702c5122bf325a90f9144d6811116bb3eb2c2da6aa5306409fc) |
| `namespace` | [namespace](data-sources--dns_proxy--reference--group-001.md#canonical-68f49b454bebe80d382d3ea71561e877a47c7d0897b8ea107bc771afac27b32c) |
| `origin_servers` | [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-bce57a5e93c61e9256499bd4e82de48e29d6871554904b0a1d30757f2634d8db) |
| `origin_servers.health_checks` | [origin_servers.health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-802a3716a1933d4fee4475119c04f09a4b9d169d747bfc8d00835ec4ba0e1c8e) |
| `origin_servers.health_checks.health_check` | [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-44afac441ff4d82a745488e78a27f2e60712932c45b573d9c672c28fa96c7c53) |
| `origin_servers.health_checks.health_check.dns_health_check` | [origin_servers.health_checks.health_check.dns_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-64fce48f919e2dff8080624076db87c2c4fa77a8c79501f378b3deffbc57e885) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_rcode` | [origin_servers.health_checks.health_check.dns_health_check.expected_rcode](data-sources--dns_proxy--reference--group-001.md#canonical-76c3bb693a9dac35781cc6a11fc9428d0cf592ef08909bb33acea209e631cccb) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_record_type` | [origin_servers.health_checks.health_check.dns_health_check.expected_record_type](data-sources--dns_proxy--reference--group-001.md#canonical-7c7577a4ddb974df8433ef6b68b47e6afb46ea7610e1012fd733610dea18fb60) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_response` | [origin_servers.health_checks.health_check.dns_health_check.expected_response](data-sources--dns_proxy--reference--group-001.md#canonical-f9d93e3d37f618b6a0f5f8005d6e751e2f8b49cb6bdbf792369a836199413454) |
| `origin_servers.health_checks.health_check.dns_health_check.query_name` | [origin_servers.health_checks.health_check.dns_health_check.query_name](data-sources--dns_proxy--reference--group-001.md#canonical-2a3519f6304ea3b390f1bf373ba19ee9adab91aeb4397c9c599b62db9517da72) |
| `origin_servers.health_checks.health_check.dns_health_check.query_type` | [origin_servers.health_checks.health_check.dns_health_check.query_type](data-sources--dns_proxy--reference--group-001.md#canonical-c6b7fecdb32118a29c7534bed9011b9d2c663fed0e719a8fb0ffc66ef90c7164) |
| `origin_servers.health_checks.health_check.dns_health_check.reverse` | [origin_servers.health_checks.health_check.dns_health_check.reverse](data-sources--dns_proxy--reference--group-001.md#canonical-b9ae22bd9e8c3f59361e58ed00cb1b02a0cafb812720dce8e6d6a1e4fb2479ad) |
| `origin_servers.health_checks.health_check.icmp_health_check` | [origin_servers.health_checks.health_check.icmp_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-6fc474781eb7c582b36e784d2c3cf23bb357558ef766f3f9457700fcd4367c95) |
| `origin_servers.health_checks.health_check.tcp_health_check` | [origin_servers.health_checks.health_check.tcp_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-a8abc3c91143ab8b98fcf0bd222f5a09f8c73f32314320b7650f96a526404d03) |
| `origin_servers.health_checks.health_check.tcp_health_check.expected_response` | [origin_servers.health_checks.health_check.tcp_health_check.expected_response](data-sources--dns_proxy--reference--group-001.md#canonical-ee1b8948cf71765bd6fd399ef12983ac0c9e5c802c001cbaca28f2ed7776631d) |
| `origin_servers.health_checks.health_check.tcp_health_check.send_payload` | [origin_servers.health_checks.health_check.tcp_health_check.send_payload](data-sources--dns_proxy--reference--group-001.md#canonical-6effb77574703667ae66985ec7b3b4286a180ff5c87e6b2c77d1f67e16b96b82) |
| `origin_servers.health_checks.healthy_threshold` | [origin_servers.health_checks.healthy_threshold](data-sources--dns_proxy--reference--group-001.md#canonical-8679c3516065d61052f2072201a0f1f3e9af1d0cedd09c51cc208c089b72ac75) |
| `origin_servers.health_checks.interval` | [origin_servers.health_checks.interval](data-sources--dns_proxy--reference--group-001.md#canonical-86a01b9d2340899e5fcd5ae2a83f8f3fe9c8a9d0a94e035df06a8962acf262a5) |
| `origin_servers.health_checks.timeout` | [origin_servers.health_checks.timeout](data-sources--dns_proxy--reference--group-001.md#canonical-f7d9809a212e31b15e73cafd3193aee7061beecafa981dd76fd9e6daf1eeadf7) |
| `origin_servers.health_checks.unhealthy_threshold` | [origin_servers.health_checks.unhealthy_threshold](data-sources--dns_proxy--reference--group-001.md#canonical-ae0cfb9078d5176f930c8999171e7372795f806c25a15864bc2586a5078aa878) |
| `origin_servers.origin_servers` | [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-50b585653ffa072eead0278aa156375a1af3e1e91cd4b2c322da0416e7d53e96) |
| `origin_servers.origin_servers.k8s_service` | [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-1f24b75f7c722e61a4b1c6b897b91352627e864a54d8ec278f699e944eeca669) |
| `origin_servers.origin_servers.k8s_service.inside_network` | [origin_servers.origin_servers.k8s_service.inside_network](data-sources--dns_proxy--reference--group-001.md#canonical-90963d2cfd5a791b030742eb77fd241edaeae46f918a8413c502489fa1f6450b) |
| `origin_servers.origin_servers.k8s_service.outside_network` | [origin_servers.origin_servers.k8s_service.outside_network](data-sources--dns_proxy--reference--group-001.md#canonical-ab9b5b76b16032268df758d1ac805a08f815c6053d060748e0c564027cda620b) |
| `origin_servers.origin_servers.k8s_service.protocol` | [origin_servers.origin_servers.k8s_service.protocol](data-sources--dns_proxy--reference--group-001.md#canonical-768044e37707c0b8a08ea6696ae3409c838ffe464cb9da228f21abadb9268960) |
| `origin_servers.origin_servers.k8s_service.service_name` | [origin_servers.origin_servers.k8s_service.service_name](data-sources--dns_proxy--reference--group-001.md#canonical-4932f684ac66e431e6280fae6e55e9e62e8d5ccf467d2989c7d6e8a7eafcb68d) |
| `origin_servers.origin_servers.k8s_service.site_locator` | [origin_servers.origin_servers.k8s_service.site_locator](data-sources--dns_proxy--reference--group-001.md#canonical-845bce43f53f284b9546fbd101f6bd9beeca7541a6ca8c4ec70e9c1d59852986) |
| `origin_servers.origin_servers.k8s_service.site_locator.site` | [origin_servers.origin_servers.k8s_service.site_locator.site](data-sources--dns_proxy--reference--group-001.md#canonical-b4d091c829e4bae161853d32b51628c97936b77d2202885cc15312fd1014af99) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.name` | [origin_servers.origin_servers.k8s_service.site_locator.site.name](data-sources--dns_proxy--reference--group-001.md#canonical-5873a78e824f582d8e509a910ec52a80c4ece01e6a5a4925d4f2102dc82cecbb) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.namespace` | [origin_servers.origin_servers.k8s_service.site_locator.site.namespace](data-sources--dns_proxy--reference--group-001.md#canonical-46025812d2af665aec60c8adccca02429c5de9db6c70ca01abdcb40eed091088) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.tenant` | [origin_servers.origin_servers.k8s_service.site_locator.site.tenant](data-sources--dns_proxy--reference--group-001.md#canonical-b8b1a13cc54c61aeb58af63c724ca37e5390e1dfee39cf455a8558c69b910d41) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site](data-sources--dns_proxy--reference--group-001.md#canonical-566a189fe86f083a2314b99cbecd984a5e3b65bbf247af0bcaebca80852cb067) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name](data-sources--dns_proxy--reference--group-001.md#canonical-5a44860f7387d1db16be0467406b7b42c13912daff52197c4968ee0cf46ff33a) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace](data-sources--dns_proxy--reference--group-001.md#canonical-4a45008ed87364952e9fd670ce3c505958e4b417a5d0986d1456801011db2d4c) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant](data-sources--dns_proxy--reference--group-001.md#canonical-349d2008602b8c11eb0759fd36276fe66115519a802a18a9d91ece252211677a) |
| `origin_servers.origin_servers.k8s_service.snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-27352dce247acde15d4525a1816200ab66bcdd32849070d7d865c5060772e0d5) |
| `origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-064dc1bd4138d0b920b36cc0031b8a694ce5ef59ef4e8697d41cee47041d66c3) |
| `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-adf5a6bb0e75de9a3a28e78cc47c544070ef236bed379d225580fd8eb14b7aa8) |
| `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes](data-sources--dns_proxy--reference--group-001.md#canonical-c9367fdb9c78e49907dd5fcf6417c4279876517940f7b26a99ab5299ea882f50) |
| `origin_servers.origin_servers.k8s_service.vk8s_networks` | [origin_servers.origin_servers.k8s_service.vk8s_networks](data-sources--dns_proxy--reference--group-001.md#canonical-5816d102dd92e02105dec308ce8ab22db47b59de41cba8943584705138e3a9de) |
| `origin_servers.origin_servers.no_preference` | [origin_servers.origin_servers.no_preference](data-sources--dns_proxy--reference--group-001.md#canonical-4cb943213ed34b50186c27c0a56885a0de60aa8212ffc7fbab5f0f0e8f3db4a9) |
| `origin_servers.origin_servers.public_ip` | [origin_servers.origin_servers.public_ip](data-sources--dns_proxy--reference--group-001.md#canonical-3904e9626ff586742b3c7b06e611ec5aaece615773c533e3e8959f3224245a68) |
| `origin_servers.origin_servers.public_ip.ip` | [origin_servers.origin_servers.public_ip.ip](data-sources--dns_proxy--reference--group-001.md#canonical-405dc5ea9c6ec1021ef24549485e622500a2a1a20d43ab6bfaca884d7c318164) |
| `origin_servers.origin_servers.public_name` | [origin_servers.origin_servers.public_name](data-sources--dns_proxy--reference--group-001.md#canonical-2d0e83c7a64ebf8c22069fd4080745d1da0ca677451c1a4dd6a8f6ab51912887) |
| `origin_servers.origin_servers.public_name.dns_name` | [origin_servers.origin_servers.public_name.dns_name](data-sources--dns_proxy--reference--group-001.md#canonical-625d344ceb8d49236becc738475e4e6cbd387722104369f76e97319a72e2ac2d) |
| `origin_servers.origin_servers.public_name.refresh_interval` | [origin_servers.origin_servers.public_name.refresh_interval](data-sources--dns_proxy--reference--group-001.md#canonical-3a1bf226a19d2cd014cc0743d47ccb9fdcee77b223a5239ef40d2ea309d7becc) |
| `origin_servers.origin_servers.site_preferences` | [origin_servers.origin_servers.site_preferences](data-sources--dns_proxy--reference--group-001.md#canonical-75c19670f1d19a48e4ab165faa222cbb9f1a9793db0e2da7a330ad4b407e92bf) |
| `origin_servers.origin_servers.site_preferences.refs` | [origin_servers.origin_servers.site_preferences.refs](data-sources--dns_proxy--reference--group-001.md#canonical-1c3ee67964cba41aa446325b0024383512aa352497449dfd09e1bd4a78f841fa) |
| `origin_servers.origin_servers.site_preferences.refs.name` | [origin_servers.origin_servers.site_preferences.refs.name](data-sources--dns_proxy--reference--group-001.md#canonical-56d2bd12bf306c126a03c8ee1fc7c533915d2483b45ff2987b9c66fcea341338) |
| `origin_servers.origin_servers.site_preferences.refs.namespace` | [origin_servers.origin_servers.site_preferences.refs.namespace](data-sources--dns_proxy--reference--group-001.md#canonical-9871f99fb97c814ea70dccb7204581075090bc5d2d9b526b6f1e770919fd9139) |
| `origin_servers.origin_servers.site_preferences.refs.tenant` | [origin_servers.origin_servers.site_preferences.refs.tenant](data-sources--dns_proxy--reference--group-001.md#canonical-92d2c9fd59b630debfe56529de89c0886cdee13246e2d5cf3f09a17a7c34626f) |
| `protocol_inspection` | [protocol_inspection](data-sources--dns_proxy--reference--group-001.md#canonical-14e8a5ab2a9cb962646f1fa25ded818a5281c1e8449b0be4e6f342e8ad5fd59f) |
| `protocol_inspection.name` | [protocol_inspection.name](data-sources--dns_proxy--reference--group-001.md#canonical-e94348c0436854d7a955b1d8c30911b6540f1ab95d2fc833cb1084d54df406b3) |
| `protocol_inspection.namespace` | [protocol_inspection.namespace](data-sources--dns_proxy--reference--group-001.md#canonical-aa24e770c5a0ed73b074e1cfd38b75013ad32918dcbe2196401a7eaa96a1c247) |
| `protocol_inspection.tenant` | [protocol_inspection.tenant](data-sources--dns_proxy--reference--group-001.md#canonical-1b4af943ecc39f837e3651ef179cfd00cb1b23e7272b8b5798494cf58716ccba) |
| `proxy_advertisement` | [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-4d7a03319bb87c24886f8bd21c7a8489c284cd37fb5fb77156b7de6df939bee4) |
| `proxy_advertisement.advertise_custom` | [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-fcf0a55395e10c6b278b6a03cccc9b2b9d444c2b6b5a704c94aefa216423b3b3) |
| `proxy_advertisement.advertise_custom.advertise_where` | [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3d631337c524759727ea6936c97c6feec543ccb4491f9035a14e77c6941628b0) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-001.md#canonical-ff16fa2f1e3733d3607350c268adffeae4f4acff7740c89c0a0fb930b6db61e6) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--dns_proxy--reference--group-001.md#canonical-e9f9d200ab2f3cafdaf145e770fd9e29578fbbf086f5409e3bc160119ac6538d) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](data-sources--dns_proxy--reference--group-001.md#canonical-d78b0fd58d039571f6cffff1c3b39ae227fe28e56d71b39cb8bef95e885ed6cc) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](data-sources--dns_proxy--reference--group-001.md#canonical-191648537ed3b6aa822345e86468b78f899b3ba2ab9bf882c3b9b3b3f2ceb00b) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-aadb4c600dd3cb9edef908b78b9d33208abc3e063cb907dc87a565a68cf5069e) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-12e07e62837889ed4d39ce59f84682c944c562a642feb7701f1f22147ca29d22) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-3d3f07eac8f5b47d5303bf72ea279b5a1b531898745fcd55b03266a3ed2fbef6) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name](data-sources--dns_proxy--reference--group-002.md#canonical-18b1e76a95303cad7cc1df145c3fce504474918cbaa87da419f458c356a14e01) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-885171a9603ed02608e96fd0a96e9d509cf02a36c8fe390ebf912a8e7066883e) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-8f4174e48a44cb0a028a6740a31dd0bc2f46cfb19a1fc8aabf84c0b07eda6eb3) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-b518214d6b63aace97742963c6479668989c17b1674937eb9ae08e812fff9048) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-f37e6841bb8722e6cdbba344a31ca49a639e132b0ada680157c4de08cb0d605a) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](data-sources--dns_proxy--reference--group-002.md#canonical-94f7136c4ace82b6ef88cd7429afa063435a3fb0bf38c99e3ca6d362d66924a0) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-6d4ce4d1a41d0d759cca03df1680ee47ad9c3933edf9a092fc956e09a91340fa) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-2ff842a1b3fac92e3a29a977dd3e757f3798ea214d3e3d7d92d7dfb228fe4379) |
| `proxy_advertisement.advertise_custom.advertise_where.port` | [proxy_advertisement.advertise_custom.advertise_where.port](data-sources--dns_proxy--reference--group-001.md#canonical-627808d74fce31d3019eba286d6cb1a4a8119e5320541f34ffa154f51347f6fc) |
| `proxy_advertisement.advertise_custom.advertise_where.port_ranges` | [proxy_advertisement.advertise_custom.advertise_where.port_ranges](data-sources--dns_proxy--reference--group-001.md#canonical-59ff688ba3d639b0b2fe8a4a5dc948349ca36e1cb803fe27e139a2e2b8a696d4) |
| `proxy_advertisement.advertise_custom.advertise_where.site` | [proxy_advertisement.advertise_custom.advertise_where.site](data-sources--dns_proxy--reference--group-002.md#canonical-c2b13ce75cbd2e5d8e68cc86bdeb81b34a9c98848ed9cd74a3c677f68e1f2bfb) |
| `proxy_advertisement.advertise_custom.advertise_where.site.ip` | [proxy_advertisement.advertise_custom.advertise_where.site.ip](data-sources--dns_proxy--reference--group-002.md#canonical-5ea60f4cb67b4dca02af884a2af7281f4b37e2651bb97049c3a264386d5472f1) |
| `proxy_advertisement.advertise_custom.advertise_where.site.network` | [proxy_advertisement.advertise_custom.advertise_where.site.network](data-sources--dns_proxy--reference--group-002.md#canonical-0e0615fca8a23d720f1090b482b72a691ce0dfb0385f001240c34806e898aa37) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site` | [proxy_advertisement.advertise_custom.advertise_where.site.site](data-sources--dns_proxy--reference--group-002.md#canonical-fcca4ceeda1816921bd59423be986322b6ff3ac9c9aa0e605dbe3e1d98392cf1) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.name` | [proxy_advertisement.advertise_custom.advertise_where.site.site.name](data-sources--dns_proxy--reference--group-002.md#canonical-615fd44cb8ee54d42d4f2c291fcd6080b225fccdabe21027668208f55c687b79) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.site.site.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-31eea8ea53af5ce0f630e183ebcd73140e2837a49880f10d40258ebb1d6465d3) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.site.site.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-021275f7217f86434b881f7885c3e97dc2f4b927a3353a52ab8585d22b7b7f83) |
| `proxy_advertisement.advertise_custom.advertise_where.use_default_port` | [proxy_advertisement.advertise_custom.advertise_where.use_default_port](data-sources--dns_proxy--reference--group-002.md#canonical-0cad932476e3e3a59ea1593fd8197877611659762cee3fbd90756b65a7c09789) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-bdea9c85c542a40e96ae779e1cbf4a6c7fef3fd1f88bbe26ee26c4793ff60a32) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--dns_proxy--reference--group-002.md#canonical-98c0fafe260d1c3091846a59174d0717658e30efa312af2d7e3a07317e463777) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](data-sources--dns_proxy--reference--group-002.md#canonical-fdc8083742a7642e1898f73a4d4f4ce478a9a766b96f26d1fc484b030e7377d5) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip](data-sources--dns_proxy--reference--group-002.md#canonical-63046b4ba0617e110ac53cff9ce83d5e4a06f968b4fcb1bca9f5b074d90fb3d9) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip](data-sources--dns_proxy--reference--group-002.md#canonical-048cc0609c288d02cac3e14522b9d6f9a3e65e3906092776b833bb7a0f4dca96) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-2312675bbcc2ecd253dca45e251257199f83f52cdc8ca2d349e7aafe60fe492b) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name](data-sources--dns_proxy--reference--group-002.md#canonical-b57f7467897d325bbb08f318c50db31cf756e4ed39d84003a7deb9c9d0c04ef2) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-c3344bd38705e44eaaa9f5ea2c951ab5169d602efc5944a028dc6a7c4817b6cb) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-24948a1ee0db876d4e16fc208ae943ce53401cd6a4788a947f029474ef29a994) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-8c4b14b4fa228b4ecf09bb574c2adf62eddb13a7b12ef2b3cd7617a7a83a03db) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.network](data-sources--dns_proxy--reference--group-002.md#canonical-3e8f07772f57d0adcc4d3fecf137db8614372f8a5a045d9d2af798dee77910bb) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-6cca8f02619d21f6bf23a340b3b6bc4f07474a6d517e31bbec21446b108dd2ee) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name](data-sources--dns_proxy--reference--group-002.md#canonical-8e31420706552bf805bd98ba9aa35ddbdc67f02d361f4b703ada7a61fcd5798b) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-2a23245da545d52e85a4011e84df9b38c27f4f666de368a9ba8233d4c2d24d9a) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-8effe30ce233bb1bf9a424442a0f78873f36b382dbf89dd44e56c8e0e3cfbea1) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](data-sources--dns_proxy--reference--group-002.md#canonical-7e408690a758b100baf4411c57051ddd5ee5c44593114fc7854dbbf37628a4af) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip](data-sources--dns_proxy--reference--group-002.md#canonical-03f412f66798f79b2ea4a6aa0cb7d1c6c504f40f23aedc1fc0726d3e410a77a1) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network](data-sources--dns_proxy--reference--group-002.md#canonical-c901df74975eda39a4878ac06ee786ab6f838f9f09269efd4fc95d32431c6ffe) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-395fbf136ae9f1182f7a9155796944a456b79d33a34959f157d007b23cb5bc49) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](data-sources--dns_proxy--reference--group-002.md#canonical-a6495a51d0c36a5b50239cce60b691cc79cea0f3bf5612843fe010040e47198b) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-8cd1208ad2f7b14c9ea6a8435df058779e1672e289630a116829fd69113f9b47) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-512e80cfb7d90efce4e6bdb0a88c3c9e7a737ecfb4465d590fdca850e10e61af) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-4f5906fe871268ae7f9483e71839524b2f3c93d1630865a594a262f96cc006ad) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](data-sources--dns_proxy--reference--group-002.md#canonical-d2942572aa976f5a0de9ea7dc6b406d36bb63f29520e9fc9ffd2eeea80cb699c) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name](data-sources--dns_proxy--reference--group-002.md#canonical-e0ed1363e5494274fbcde0dbfc9050280fcb44f673ec15e6c37aeec3221138db) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-d0b2e8bc67bd9ce4ab991a8a59e40563fdcea8fb27f89904451260c645363415) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-e5d431a3f01daf5fd1b6a16bc5e66c7ceede05e8b7afbb23df73106961e20aea) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-48f211ab470ef50498ded4604fb9f1e5b6e3bb73a7e1276b95dcfdf774411897) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name](data-sources--dns_proxy--reference--group-002.md#canonical-489161520e4ce1a194d14f8b1c8f263a8eb6a508ad7a37b5b3c78cbd9e705c1a) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-8a6e5413b1957ee27906b24c21fd05ba77d41b6d8c882036272deb6e8d997c84) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-7faa77b13198ae0ed186b293dada3c33fb718764ef05a609e98acd59eadcccd4) |
| `proxy_advertisement.advertise_dualstack_on_public` | [proxy_advertisement.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3963aa67ecea5a48cf21756cfc739b8be6c92876e2ed22f255a1eedabb82310a) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_dualstack_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-de345e6a6a71c0e945ee67bf60d7fba1cf96646eb909eca7e06762831e5f5778) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.name](data-sources--dns_proxy--reference--group-002.md#canonical-3c941c66c5c84ca725f28948ae2b9cb10a84b38f05c61aee927c6487bd0b241b) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-11dd5f8b7b19f5f3aa7ac388ef028bd8e0b95e4631569edea55bdde3e3cd47bb) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-b085403f66c25aa61678a1a91d0f14c702c2271e3bd5f91a20a5d852a8549eb2) |
| `proxy_advertisement.advertise_on_public` | [proxy_advertisement.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-a157cacdd0fc6cf07a4d118bd70ca4ab97f66330910a5cc2b64359aee88e18e3) |
| `proxy_advertisement.advertise_on_public.public_ip` | [proxy_advertisement.advertise_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-355300169a51d13fb2df11b96f46cdaa9dc37e3bebb21eda1ee200f3568c9936) |
| `proxy_advertisement.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_on_public.public_ip.name](data-sources--dns_proxy--reference--group-002.md#canonical-37c4cf38e6283d5c1572010f7e0c6b0d48fc382870117fa23083a85e824c079f) |
| `proxy_advertisement.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_on_public.public_ip.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-9011ac00453e3ff14410b2bc8d339638d835b35e6971bb39d45b1a5db1ea8730) |
| `proxy_advertisement.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_on_public.public_ip.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-c4d2c4d07d18be7ab0cc59f69767a299d31f0793cfd20271ec2b41db71d795cc) |
| `proxy_advertisement.advertise_on_public_default_dualstack_vip` | [proxy_advertisement.advertise_on_public_default_dualstack_vip](data-sources--dns_proxy--reference--group-002.md#canonical-02754fc93b9f773c33a8f617d96c6d013429b8e5defe5bea09a9aa6235fe7237) |
| `proxy_advertisement.advertise_on_public_default_ipv6_vip` | [proxy_advertisement.advertise_on_public_default_ipv6_vip](data-sources--dns_proxy--reference--group-002.md#canonical-90b526a24342cafc2151115834fff21f24d4fb250fab4289016b4dd022949b7d) |
| `proxy_advertisement.advertise_on_public_default_vip` | [proxy_advertisement.advertise_on_public_default_vip](data-sources--dns_proxy--reference--group-002.md#canonical-32600ac7754133bb5c5a9fce8a9d05daef8ef970807c5035742fdbfb9972e173) |
| `proxy_advertisement.advertise_v6_on_public` | [proxy_advertisement.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-51c3c45931d9e05c50604070c4e664ee44024ebcbb67efa90e05f666aff1c924) |
| `proxy_advertisement.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_v6_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-d0a3b2c79b45bc6af09c709ea5dc47b11084076b3e12e93bcd575e471417a170) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_v6_on_public.public_ip.name](data-sources--dns_proxy--reference--group-002.md#canonical-0d82b4e6b2bbb311beec40b43b98b07979b625015fdd9d8dfd282e72813e462f) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_v6_on_public.public_ip.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-fb4e6a1d7e18dfaa81dde820309d7ad2fb853e7f785883c2489a3fc8877f427d) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_v6_on_public.public_ip.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-a12a0e5f259d49576c93b2deb2adf2052f599a6ee860e2a800947886b7b94aaa) |
| `proxy_advertisement.do_not_advertise` | [proxy_advertisement.do_not_advertise](data-sources--dns_proxy--reference--group-002.md#canonical-6049129f23630cf5b93004df2c0a123064ba72308bc0b0b6e9fe725b55754d62) |
| `transport_type` | [transport_type](data-sources--dns_proxy--reference--group-001.md#canonical-79e209b25a748c04037bbc20aca4521f179671f836980507009e768353b36b88) |

<a id="canonical-3aeab6f284790f39ccdc8fc63fc6ecc2a6aa4f49256c337ec067301cf77a028e"></a>

## Next pages — Property reference / 4ab100afcb95 / 12

- [cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-9b8331703e912880af23b8d21216c343c0f6f0c28f256e527f682650226f0a9b)
- [ddos_profile](data-sources--dns_proxy--reference--group-001.md#canonical-2e1cb6af469fb44b81586f0cc731388b8574518c894a2e18104e9a1c380203d0)
- [irules](data-sources--dns_proxy--reference--group-001.md#canonical-7b1150e1330b61c4d3436316083ae76e767a999d7648148beb7acbff7654f112)
- [lb_algorithm](data-sources--dns_proxy--reference--group-001.md#canonical-879501040bd92674fb08f9bdd32caaa6247162e032ad593e239557236f049bc7)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [protocol_inspection](data-sources--dns_proxy--reference--group-001.md#canonical-6ad138bdca9dd8c2037ea1d3f939f5ee70ff85ce14bff91ab9f5ebaba79d304c)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-9b8331703e912880af23b8d21216c343c0f6f0c28f256e527f682650226f0a9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6441e135031457fb182bc723d0d005a26b8b42d2c9b8a1127b70cb5e53e5edb6"></a>

## cache_profile — cache_profile / 6b1bd38b3038 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- cache_profile

<a id="canonical-f958d1f4b2c1b45f71096efa946362febc938405ae758abf8ab1817d4db80b08"></a>

Type: `"single"`. Computed.

DNS Cache specifies cache configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_profile_choice": "[\"cache_size\",\"disable_cache_profile\"]"
}
```

<a id="canonical-e73212d9a4dfacec20374f06f6db574a8c2486b0c43f3aeb01bf2734c08336f7"></a>

## Direct properties — cache_profile / 6b1bd38b3038 / 3

<a id="canonical-43f22c45557bd0277c63e5618b72e34b76bf905dd68e39d190de5c98bfb20673"></a>

<a id="canonical-2561e0c8f35640a8b752262f5bdbd7af0393d81c129c609fafde66987bffa37c"></a>

## cache_size property — cache_profile / 6b1bd38b3038 / 4

Type: `"number"`. Computed.

Exclusive with \[disable\_cache\_profile\] cache size.

Upstream description:

Exclusive with \[disable\_cache\_profile\] cache size.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10240,
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
    "ves.io.schema.rules.uint32.lte": "10240"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10240"
  }
}
```

- [disable_cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-6f59c722c0d0b90fc71e63fe69a0c6229274da7800004ed6aae7bc86d1bf2814): complete subsection reference.

<a id="canonical-4abbccec9d7d7da310d40de775767905d22087525c76210887824d5f3fc26574"></a>

## Next pages — cache_profile / 6b1bd38b3038 / 5

- [cache_profile.disable_cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-6f59c722c0d0b90fc71e63fe69a0c6229274da7800004ed6aae7bc86d1bf2814)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-6f59c722c0d0b90fc71e63fe69a0c6229274da7800004ed6aae7bc86d1bf2814"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d306b9b50863d747e915c176cb62d8e8c02dff38b4619402a10b2acd09f0f97a"></a>

## cache_profile.disable_cache_profile — cache_profile.disable_cache_profile / 5e0c9ee997da / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-9b8331703e912880af23b8d21216c343c0f6f0c28f256e527f682650226f0a9b)
- cache_profile.disable_cache_profile

<a id="canonical-afdd403ae0c0b4281761aeca0e5f50fd9b6f194e526cf7dcdaacc40c8d0f37a6"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable cache profile.

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

<a id="canonical-01cf35de1a282add77e2bd377e873fabcb60b48b79a879f76bba8bd5f7f99c44"></a>

## Direct properties — cache_profile.disable_cache_profile / 5e0c9ee997da / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0a1b5785f4c239739c990bcc7a7a4a1e4c14d6c06c30ff6e5594bf4c87174ec3"></a>

## Next pages — cache_profile.disable_cache_profile / 5e0c9ee997da / 4

- [cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-9b8331703e912880af23b8d21216c343c0f6f0c28f256e527f682650226f0a9b)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-2e1cb6af469fb44b81586f0cc731388b8574518c894a2e18104e9a1c380203d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ffe717f7ae6d9bea5b4bb962b3a913378042e5c1510f2136566a3b9dea5b5f9"></a>

## ddos_profile — ddos_profile / eeb36df2baa8 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- ddos_profile

<a id="canonical-4aac37425e48c3e77ba4beb99a3d9b948a41e46e9bf31ee56264ebaa36fa4901"></a>

Type: `"single"`. Computed.

Configuration parameter for ddos profile.

Upstream description:

DDoS Protection Rule for DNS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ddos_mitigation_choice": "[\"disable_ddos_mitigation\",\"enable_ddos_mitigation\"]"
}
```

<a id="canonical-9b60a37a2412ba63df0631ac836d34e1754dd5261a7eaa33ae6a7bc417d83148"></a>

## Direct properties — ddos_profile / eeb36df2baa8 / 3

- [disable_ddos_mitigation](data-sources--dns_proxy--reference--group-001.md#canonical-d1fe0045d74b08d4fa9825eb839354fa6b694454b1892eb1465a55f90af569cf): complete subsection reference.

- [enable_ddos_mitigation](data-sources--dns_proxy--reference--group-001.md#canonical-a2e795225967d6d2465e9cc7e5a0b3abee1a87e728d7933a5c5b3ba6f83cad00): complete subsection reference.

<a id="canonical-1bf5260285852eef3dbee5fd2c2ba13f03ae12d743e9d14f4842597855e019ae"></a>

## Next pages — ddos_profile / eeb36df2baa8 / 4

- [ddos_profile.disable_ddos_mitigation](data-sources--dns_proxy--reference--group-001.md#canonical-d1fe0045d74b08d4fa9825eb839354fa6b694454b1892eb1465a55f90af569cf)
- [ddos_profile.enable_ddos_mitigation](data-sources--dns_proxy--reference--group-001.md#canonical-a2e795225967d6d2465e9cc7e5a0b3abee1a87e728d7933a5c5b3ba6f83cad00)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-d1fe0045d74b08d4fa9825eb839354fa6b694454b1892eb1465a55f90af569cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-678695411586928c30ac8e59245ae6737b87db2d76f9813c77bcea07355b61df"></a>

## ddos_profile.disable_ddos_mitigation — ddos_profile.disable_ddos_mitigation / 23ddac935d44 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [ddos_profile](data-sources--dns_proxy--reference--group-001.md#canonical-2e1cb6af469fb44b81586f0cc731388b8574518c894a2e18104e9a1c380203d0)
- ddos_profile.disable_ddos_mitigation

<a id="canonical-7929f4308f1a78651e357d9c1270000f4c0a7b8786d77e9530391b4a7995f762"></a>

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

<a id="canonical-47013e6a8a1969345cdc363c9f462b61a7dd6a1cb2753ba9f8057ca259361fae"></a>

## Direct properties — ddos_profile.disable_ddos_mitigation / 23ddac935d44 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-14236886807bb44847f05da01f46dfa6e254cb62463bfa7cb93d1fdb9df6e19e"></a>

## Next pages — ddos_profile.disable_ddos_mitigation / 23ddac935d44 / 4

- [ddos_profile](data-sources--dns_proxy--reference--group-001.md#canonical-2e1cb6af469fb44b81586f0cc731388b8574518c894a2e18104e9a1c380203d0)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-a2e795225967d6d2465e9cc7e5a0b3abee1a87e728d7933a5c5b3ba6f83cad00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f90c7f19b0b69f997e01688cd41135cc262d7bbc8b5c3feaee1874d07fd586df"></a>

## ddos_profile.enable_ddos_mitigation — ddos_profile.enable_ddos_mitigation / 7791ba47d8a9 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [ddos_profile](data-sources--dns_proxy--reference--group-001.md#canonical-2e1cb6af469fb44b81586f0cc731388b8574518c894a2e18104e9a1c380203d0)
- ddos_profile.enable_ddos_mitigation

<a id="canonical-12800f8a6a539cd135ab1abcc9f4e3da268dd3c7fb70053ee062b1230ee93138"></a>

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

<a id="canonical-76e3a239f3f0c869c6df5496411b9450fc4f6ba7c20107918b94bf1df87dbdd0"></a>

## Direct properties — ddos_profile.enable_ddos_mitigation / 7791ba47d8a9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-780c7a8261978003192c6d32652cd059725f8bd3aaf3dbde823bb7c4b8d490ff"></a>

## Next pages — ddos_profile.enable_ddos_mitigation / 7791ba47d8a9 / 4

- [ddos_profile](data-sources--dns_proxy--reference--group-001.md#canonical-2e1cb6af469fb44b81586f0cc731388b8574518c894a2e18104e9a1c380203d0)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-7b1150e1330b61c4d3436316083ae76e767a999d7648148beb7acbff7654f112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8de9d3583df8f3b871ceba71f000a8cd8d297a4195dc92ffbdcd4e32cc880c5b"></a>

## irules — irules / f4e298563932 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- irules

<a id="canonical-825788726606ffe729450ee5433fa230cfdf9a8415a82a96774ed4e37dfecf5f"></a>

Type: `"list"`. Computed.

OPTIONS for attaching iRules to DNS proxy.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0a61cc792de50c200fecb47b26edd5247b2b9d536abee1dd56b7337987bf71e5"></a>

## Direct properties — irules / f4e298563932 / 3

<a id="canonical-467be693ab8373b80d8192cccc56781f467980482f40989f75d8ddcfea8ccca3"></a>

<a id="canonical-4be4fd62f0c498b30f100595c237a981ee304af11bb00269b0ee195fb0bbd30d"></a>

## name property — irules / f4e298563932 / 4

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

<a id="canonical-45d55157f40e3612c4b5522d9b3a8ab20cbb60ef70715d29573f2244496ef689"></a>

<a id="canonical-f139451e355b822325cf8be34857dafc163e11d8492768971e0fc9eae2429365"></a>

## namespace property — irules / f4e298563932 / 5

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

<a id="canonical-4148ed888d1811512f9764b8c000516d076ca61e94948e229949a81299365e2a"></a>

<a id="canonical-dae389069cca9e2d3b8eac46b2610b5f6c0e7aae33d5e466c50051e2fe6b91be"></a>

## tenant property — irules / f4e298563932 / 6

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

<a id="canonical-1820ba1f93cc8ebd71d9964b7d0d2dc345d0312b6a463530f659a3665eb55543"></a>

## Next pages — irules / f4e298563932 / 7

- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-879501040bd92674fb08f9bdd32caaa6247162e032ad593e239557236f049bc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2278414832a1ce599a0e3feb3c594ddeafd240c43d1faedaadd0817f74ae65d2"></a>

## lb_algorithm — lb_algorithm / 86f979c05f92 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- lb_algorithm

<a id="canonical-4e41c5176bf46a40ba875a9d334a930bca67b714c89646024a4986c9a8534e8f"></a>

Type: `"single"`. Computed.

Configuration parameter for lb algorithm.

Upstream description:

Load Balancing Algorithm Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lb_algorithm_choice": "[\"round_robin\"]"
}
```

<a id="canonical-213ec18dfce9d28954d42a172dead6d55c5bc32377aaae9e6abe2dffdd740362"></a>

## Direct properties — lb_algorithm / 86f979c05f92 / 3

- [round_robin](data-sources--dns_proxy--reference--group-001.md#canonical-40fea004bea8c2c81335e81660be125a6a47531996c66742e3d08b6e3da613e0): complete subsection reference.

<a id="canonical-30b2095d6a848e171ff345aa2f774118cb38da8e291d90cbe2cc2739667c5a77"></a>

## Next pages — lb_algorithm / 86f979c05f92 / 4

- [lb_algorithm.round_robin](data-sources--dns_proxy--reference--group-001.md#canonical-40fea004bea8c2c81335e81660be125a6a47531996c66742e3d08b6e3da613e0)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-40fea004bea8c2c81335e81660be125a6a47531996c66742e3d08b6e3da613e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3090de50638510e603440c32fae140965d8addc8f2243aae0117d2f700419a7"></a>

## lb_algorithm.round_robin — lb_algorithm.round_robin / 2f47cda0474a / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [lb_algorithm](data-sources--dns_proxy--reference--group-001.md#canonical-879501040bd92674fb08f9bdd32caaa6247162e032ad593e239557236f049bc7)
- lb_algorithm.round_robin

<a id="canonical-b037dead96f53773b1e0bfad63e993e8c0efc18095f44c849a5f3de7e7dbdbb3"></a>

Type: `"single"`. Computed.

Configuration parameter for round robin.

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

<a id="canonical-3fe28bd50e9a3e6278774069d46d3b46b1d92f892aaf9659170a3047353273bb"></a>

## Direct properties — lb_algorithm.round_robin / 2f47cda0474a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-67d2793ad8f2540287abaa3a64d95683f157eda646e1781c276e0a18990c048a"></a>

## Next pages — lb_algorithm.round_robin / 2f47cda0474a / 4

- [lb_algorithm](data-sources--dns_proxy--reference--group-001.md#canonical-879501040bd92674fb08f9bdd32caaa6247162e032ad593e239557236f049bc7)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba5b6bf8f9c540acd84b6742a94fc40c4afcbd7abff56782e54e477490dc8c8a"></a>

## origin_servers — origin_servers / 961ecab789a2 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- origin_servers

<a id="canonical-bce57a5e93c61e9256499bd4e82de48e29d6871554904b0a1d30757f2634d8db"></a>

Type: `"single"`. Computed.

List of origin Servers for the DNS proxy.

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

<a id="canonical-f0835da965c94d748f7aa9f5a90bb40284fc02647ba7d45855d0cbe9d34a2eee"></a>

## Direct properties — origin_servers / 961ecab789a2 / 3

- [health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-7eaa8c8d913c4e01cd893d4928edc718847eee234c1406469a52d8245eedad6d): complete subsection reference.

- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7): complete subsection reference.

<a id="canonical-31fc37c77ce9c46bf6ac432aa6e2e8154c67802a366492f7ec36d73d28de84cd"></a>

## Next pages — origin_servers / 961ecab789a2 / 4

- [origin_servers.health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-7eaa8c8d913c4e01cd893d4928edc718847eee234c1406469a52d8245eedad6d)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-7eaa8c8d913c4e01cd893d4928edc718847eee234c1406469a52d8245eedad6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16d89651350192691f552bb189e8623d9ec91d30aea08fa2e4f25be87af6f9ea"></a>

## origin_servers.health_checks — origin_servers.health_checks / 142135948f7f / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- origin_servers.health_checks

<a id="canonical-802a3716a1933d4fee4475119c04f09a4b9d169d747bfc8d00835ec4ba0e1c8e"></a>

Type: `"single"`. Computed.

Configuration parameter for health checks.

Upstream description:

Origin Server Health Checks.

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

<a id="canonical-beab37a5b61a76c35f55e2d11e286b18be426b095c642e115b64f290fcab7354"></a>

## Direct properties — origin_servers.health_checks / 142135948f7f / 3

- [health_check](data-sources--dns_proxy--reference--group-001.md#canonical-1ebfdc9c0d1409263f20f7c4fd31bf0e40dd8c74dcc287a884005cdab92902de): complete subsection reference.

<a id="canonical-8679c3516065d61052f2072201a0f1f3e9af1d0cedd09c51cc208c089b72ac75"></a>

<a id="canonical-17ffc4d527ff188be36ebf4dc9762b4d94a334f13403648586a35b353cb5ae8a"></a>

## healthy_threshold property — origin_servers.health_checks / 142135948f7f / 4

Type: `"number"`. Computed.

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

Upstream description:

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-86a01b9d2340899e5fcd5ae2a83f8f3fe9c8a9d0a94e035df06a8962acf262a5"></a>

<a id="canonical-153f6ff24b62f91fb32704b2f0eeb255f4e6411f241bc324df832f1dacb68bed"></a>

## interval property — origin_servers.health_checks / 142135948f7f / 5

Type: `"number"`. Computed.

Time interval in seconds between two healthcheck requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-f7d9809a212e31b15e73cafd3193aee7061beecafa981dd76fd9e6daf1eeadf7"></a>

<a id="canonical-0c1f0dd0fbf0bd47b571ec63f555d537a95692eb9b9389ea4c7aabbfc805c524"></a>

## timeout property — origin_servers.health_checks / 142135948f7f / 6

Type: `"number"`. Computed.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-ae0cfb9078d5176f930c8999171e7372795f806c25a15864bc2586a5078aa878"></a>

<a id="canonical-1f21c66ec149f77ba68efbd4adfb29e36d99ffd817782879194a420ea4d83ec9"></a>

## unhealthy_threshold property — origin_servers.health_checks / 142135948f7f / 7

Type: `"number"`. Computed.

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health
checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

Upstream description:

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health
checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-cabce8fc0b1bc5e1e94d5d7414a42cc127d158f5648a75898c222e9d3ac916ac"></a>

## Next pages — origin_servers.health_checks / 142135948f7f / 8

- [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-1ebfdc9c0d1409263f20f7c4fd31bf0e40dd8c74dcc287a884005cdab92902de)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-1ebfdc9c0d1409263f20f7c4fd31bf0e40dd8c74dcc287a884005cdab92902de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0ace29bd45ddee69da50e0ee8fa0673d55a7cc869277964cf72a1c9f876e57a"></a>

## origin_servers.health_checks.health_check — origin_servers.health_checks.health_check / 1c326bed9453 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-7eaa8c8d913c4e01cd893d4928edc718847eee234c1406469a52d8245eedad6d)
- origin_servers.health_checks.health_check

<a id="canonical-44afac441ff4d82a745488e78a27f2e60712932c45b573d9c672c28fa96c7c53"></a>

Type: `"list"`. Computed.

List of Health Checks. List of Health Checks.

Upstream description:

List of Health Checks.

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
    "minItems": 0,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f84c9f1a9beb7415ae74d1f4e59efd5487ecd5f772847a8a5bb764d4c4433da8"></a>

## Direct properties — origin_servers.health_checks.health_check / 1c326bed9453 / 3

- [dns_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-360a69b44f2054adde87f1742d7645006f6c872e93b756d9a0cda32454bc9e21): complete subsection reference.

- [icmp_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-7efba39b7ee218253da594186bee26c8055e6e33b972e8477c1fb3cad20c3c2a): complete subsection reference.

- [tcp_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-24a8dcb5e325cef59182190c31e2dcee57a7ba8768b3ae7e31e5c7f5aa34e5d0): complete subsection reference.

<a id="canonical-9a6c6d14bb90b96056949897ca5a06eaae493acb83bc7eccf632685d431df3dc"></a>

## Next pages — origin_servers.health_checks.health_check / 1c326bed9453 / 4

- [origin_servers.health_checks.health_check.dns_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-360a69b44f2054adde87f1742d7645006f6c872e93b756d9a0cda32454bc9e21)
- [origin_servers.health_checks.health_check.icmp_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-7efba39b7ee218253da594186bee26c8055e6e33b972e8477c1fb3cad20c3c2a)
- [origin_servers.health_checks.health_check.tcp_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-24a8dcb5e325cef59182190c31e2dcee57a7ba8768b3ae7e31e5c7f5aa34e5d0)
- [origin_servers.health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-7eaa8c8d913c4e01cd893d4928edc718847eee234c1406469a52d8245eedad6d)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-360a69b44f2054adde87f1742d7645006f6c872e93b756d9a0cda32454bc9e21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-326685eff731dc94b93240641a0ba67a1a38b4809ff75988e510e7e255311b16"></a>

## origin_servers.health_checks.health_check.dns_health_check — origin_servers.health_checks.health_check.dns_health_check / c218e2a83dbb / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-7eaa8c8d913c4e01cd893d4928edc718847eee234c1406469a52d8245eedad6d)
- [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-1ebfdc9c0d1409263f20f7c4fd31bf0e40dd8c74dcc287a884005cdab92902de)
- origin_servers.health_checks.health_check.dns_health_check

<a id="canonical-64fce48f919e2dff8080624076db87c2c4fa77a8c79501f378b3deffbc57e885"></a>

Type: `"single"`. Computed.

DNS health check reports healthy if DNS query is successful and response header and answer matches
the given value.

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

<a id="canonical-0ffdac38d25909eea13e8a0f5bca63b0e7b6229b2139894ed6c4fa570cbc3f2d"></a>

## Direct properties — origin_servers.health_checks.health_check.dns_health_check / c218e2a83dbb / 3

<a id="canonical-76c3bb693a9dac35781cc6a11fc9428d0cf592ef08909bb33acea209e631cccb"></a>

<a id="canonical-dfcb1077c7387db13d4ab2f76bc81b5b7d48b31f8b4628f832cd58550ab28140"></a>

## expected_rcode property — origin_servers.health_checks.health_check.dns_health_check / c218e2a83dbb / 4

Type: `"string"`. Computed.

\[Enum: DNS\_RES\_RCODE\_NOERROR|DNS\_RES\_RCODE\_ANY\] Expected DNS Response Rcode Type -
DNS\_RES\_RCODE\_NOERROR: Rcode NOERROR - DNS\_RES\_RCODE\_ANY: RCODE ANY. Possible values are
\`DNS\_RES\_RCODE\_NOERROR\`, \`DNS\_RES\_RCODE\_ANY\`. Defaults to \`DNS\_RES\_RCODE\_NOERROR\`.

Upstream description:

Expected DNS Response Rcode Type

&#8203;- DNS\_RES\_RCODE\_NOERROR: Rcode NOERROR

&#8203;- DNS\_RES\_RCODE\_ANY: RCODE ANY.

Receipt-pinned upstream constraints:

```json
{
  "default": "DNS_RES_RCODE_NOERROR",
  "enum": [
    "DNS_RES_RCODE_NOERROR",
    "DNS_RES_RCODE_ANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7c7577a4ddb974df8433ef6b68b47e6afb46ea7610e1012fd733610dea18fb60"></a>

<a id="canonical-aa039e27f86e3e29f3add173fef5e9fc75c08915d18e896451781436777f79d2"></a>

## expected_record_type property — origin_servers.health_checks.health_check.dns_health_check / c218e2a83dbb / 5

Type: `"string"`. Computed.

\[Enum: DNS\_REQUESTED\_QUERY\_TYPE|DNS\_RES\_RECORD\_TYPE\_ANY\] DNS Response Record Type -
DNS\_REQUESTED\_QUERY\_TYPE: Requested Query Type - DNS\_RES\_RECORD\_TYPE\_ANY: Any Record Type.
Possible values are \`DNS\_REQUESTED\_QUERY\_TYPE\`, \`DNS\_RES\_RECORD\_TYPE\_ANY\`. Defaults to
\`DNS\_REQUESTED\_QUERY\_TYPE\`.

Upstream description:

DNS Response Record Type

&#8203;- DNS\_REQUESTED\_QUERY\_TYPE: Requested Query Type

&#8203;- DNS\_RES\_RECORD\_TYPE\_ANY: Any Record Type.

Receipt-pinned upstream constraints:

```json
{
  "default": "DNS_REQUESTED_QUERY_TYPE",
  "enum": [
    "DNS_REQUESTED_QUERY_TYPE",
    "DNS_RES_RECORD_TYPE_ANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f9d93e3d37f618b6a0f5f8005d6e751e2f8b49cb6bdbf792369a836199413454"></a>

<a id="canonical-70cefcbfc145594423904e998b4383d46d530e59af7468c0facea283f16e7468"></a>

## expected_response property — origin_servers.health_checks.health_check.dns_health_check / c218e2a83dbb / 6

Type: `"string"`. Computed.

Specifies an IPv4 or IPv6 address in the answer section of DNS Response.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_len": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  }
}
```

<a id="canonical-2a3519f6304ea3b390f1bf373ba19ee9adab91aeb4397c9c599b62db9517da72"></a>

<a id="canonical-496221481a522af43a052228c79a21f4821d15ce76205557184195d58966d9ee"></a>

## query_name property — origin_servers.health_checks.health_check.dns_health_check / c218e2a83dbb / 7

Type: `"string"`. Computed.

The query name that the monitor sends a DNS query for.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-c6b7fecdb32118a29c7534bed9011b9d2c663fed0e719a8fb0ffc66ef90c7164"></a>

<a id="canonical-4cace5ec1e672b0ed0668c69f06e876878888a8c4d33a1de700de9a4b5fa240b"></a>

## query_type property — origin_servers.health_checks.health_check.dns_health_check / c218e2a83dbb / 8

Type: `"string"`. Computed.

\[Enum: DNS\_QTYPE\_A|DNS\_QTYPE\_AAAA\] DNS Query Type - DNS\_QTYPE\_A: Query Type A -
DNS\_QTYPE\_AAAA: Query Type AAAA. Possible values are \`DNS\_QTYPE\_A\`, \`DNS\_QTYPE\_AAAA\`.
Defaults to \`DNS\_QTYPE\_A\`.

Upstream description:

DNS Query Type

&#8203;- DNS\_QTYPE\_A: Query Type A

&#8203;- DNS\_QTYPE\_AAAA: Query Type AAAA.

Receipt-pinned upstream constraints:

```json
{
  "default": "DNS_QTYPE_A",
  "enum": [
    "DNS_QTYPE_A",
    "DNS_QTYPE_AAAA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b9ae22bd9e8c3f59361e58ed00cb1b02a0cafb812720dce8e6d6a1e4fb2479ad"></a>

<a id="canonical-ff77ecc9b84c44575a10609abdedf14a41d57d2f7a4ea07c0b4ebbedbababa63"></a>

## reverse property — origin_servers.health_checks.health_check.dns_health_check / c218e2a83dbb / 9

Type: `"bool"`. Computed.

Enable the monitor operation in reverse mode. When the monitor is in reverse mode, a successful
receive string match marks the monitored object down instead of up.

Upstream description:

Enable the monitor operation in reverse mode. When the monitor is in reverse mode, a successful
receive string match marks the monitored object down instead of up.

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

<a id="canonical-2364907f1d320ade4d139e3416b1ca20d56d611f48f211f6c72f83b9ce7d7392"></a>

## Next pages — origin_servers.health_checks.health_check.dns_health_check / c218e2a83dbb / 10

- [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-1ebfdc9c0d1409263f20f7c4fd31bf0e40dd8c74dcc287a884005cdab92902de)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-7efba39b7ee218253da594186bee26c8055e6e33b972e8477c1fb3cad20c3c2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59d69feb80974d30d2e74f065a9fb155cc0b8584a582ed8e23a4bd0290eacb63"></a>

## origin_servers.health_checks.health_check.icmp_health_check — origin_servers.health_checks.health_check.icmp_health_check / 7e9ef99ae96c / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-7eaa8c8d913c4e01cd893d4928edc718847eee234c1406469a52d8245eedad6d)
- [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-1ebfdc9c0d1409263f20f7c4fd31bf0e40dd8c74dcc287a884005cdab92902de)
- origin_servers.health_checks.health_check.icmp_health_check

<a id="canonical-6fc474781eb7c582b36e784d2c3cf23bb357558ef766f3f9457700fcd4367c95"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for icmp health check.

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

<a id="canonical-82332ae85ed4349f3b0910fb1f684eabae3569d44a5ab32b95d926d928d0f0c7"></a>

## Direct properties — origin_servers.health_checks.health_check.icmp_health_check / 7e9ef99ae96c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-94b3a82759e9e1cd149b8c033229115c7af8ccd71466045a2615a8497f914b4a"></a>

## Next pages — origin_servers.health_checks.health_check.icmp_health_check / 7e9ef99ae96c / 4

- [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-1ebfdc9c0d1409263f20f7c4fd31bf0e40dd8c74dcc287a884005cdab92902de)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-24a8dcb5e325cef59182190c31e2dcee57a7ba8768b3ae7e31e5c7f5aa34e5d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-adb1ec3ea415430dba8de3c0556dfc8cda61e4fb18c7337e31147f730bb0c0f1"></a>

## origin_servers.health_checks.health_check.tcp_health_check — origin_servers.health_checks.health_check.tcp_health_check / c3f2e887a011 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-7eaa8c8d913c4e01cd893d4928edc718847eee234c1406469a52d8245eedad6d)
- [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-1ebfdc9c0d1409263f20f7c4fd31bf0e40dd8c74dcc287a884005cdab92902de)
- origin_servers.health_checks.health_check.tcp_health_check

<a id="canonical-a8abc3c91143ab8b98fcf0bd222f5a09f8c73f32314320b7650f96a526404d03"></a>

Type: `"single"`. Computed.

Monitor reports healthy status if UDP connection is successful and response payload matches expected
response pattern.

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

<a id="canonical-80e9fc30d2ffbcce17dc6e8f387caecbc524c4bcdb2bf42925895ac4e8e2f7b0"></a>

## Direct properties — origin_servers.health_checks.health_check.tcp_health_check / c3f2e887a011 / 3

<a id="canonical-ee1b8948cf71765bd6fd399ef12983ac0c9e5c802c001cbaca28f2ed7776631d"></a>

<a id="canonical-94fe8720551717b84fce1bb0213210dd411f46dc3c60d5fbab3770ef37df3142"></a>

## expected_response property — origin_servers.health_checks.health_check.tcp_health_check / c3f2e887a011 / 4

Type: `"string"`. Computed.

Specifies a regular expression pattern which will be matched against response payload.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-6effb77574703667ae66985ec7b3b4286a180ff5c87e6b2c77d1f67e16b96b82"></a>

<a id="canonical-68cd2acab23d26596e3acf79c3d529a1ada799987cf16f44b3b4c1b464ada6d2"></a>

## send_payload property — origin_servers.health_checks.health_check.tcp_health_check / c3f2e887a011 / 5

Type: `"string"`. Computed.

Send string. Text string sent in the request.

Upstream description:

Text string sent in the request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-f63aacba16f2a35f3ef0fc524b7363ebc6a9e0e578a087c56ff8d444fc6291ad"></a>

## Next pages — origin_servers.health_checks.health_check.tcp_health_check / c3f2e887a011 / 6

- [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-1ebfdc9c0d1409263f20f7c4fd31bf0e40dd8c74dcc287a884005cdab92902de)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0718c8eed3964f14f1027b46327541a7cd588355923ef45e429db286dd2d6d1e"></a>

## origin_servers.origin_servers — origin_servers.origin_servers / 8a4fb14176bf / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- origin_servers.origin_servers

<a id="canonical-50b585653ffa072eead0278aa156375a1af3e1e91cd4b2c322da0416e7d53e96"></a>

Type: `"list"`. Computed.

List Of Origin Servers. List of origin servers for Proxy.

Upstream description:

List of origin servers for Proxy.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5aabb6b73dd4b474334c5f3bc9686544ffffed560f446317639855d2c8750dc9"></a>

## Direct properties — origin_servers.origin_servers / 8a4fb14176bf / 3

- [k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6): complete subsection reference.

- [no_preference](data-sources--dns_proxy--reference--group-001.md#canonical-ea080b6c3ad3ffcc664262c5df6ccd54d54691d1d30de213e3758f7eaa68954c): complete subsection reference.

- [public_ip](data-sources--dns_proxy--reference--group-001.md#canonical-5caa19762d55c8c60d29f57d2ce91d62ec2175e7229cdad1ab2c3e52a349ad5b): complete subsection reference.

- [public_name](data-sources--dns_proxy--reference--group-001.md#canonical-a100ae9190b9ebd0445bb3f2051aa099db59162f9d6c7781db1ebefd2bae1506): complete subsection reference.

- [site_preferences](data-sources--dns_proxy--reference--group-001.md#canonical-8c0e123b7b205156804e291f333bc5323e0bdcb1dc72e572ea3d13377d63c352): complete subsection reference.

<a id="canonical-71b8494b4fe6c4a16a32d1545598248817b5174675c13fa492dd4845b5c6a5df"></a>

## Next pages — origin_servers.origin_servers / 8a4fb14176bf / 4

- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- [origin_servers.origin_servers.no_preference](data-sources--dns_proxy--reference--group-001.md#canonical-ea080b6c3ad3ffcc664262c5df6ccd54d54691d1d30de213e3758f7eaa68954c)
- [origin_servers.origin_servers.public_ip](data-sources--dns_proxy--reference--group-001.md#canonical-5caa19762d55c8c60d29f57d2ce91d62ec2175e7229cdad1ab2c3e52a349ad5b)
- [origin_servers.origin_servers.public_name](data-sources--dns_proxy--reference--group-001.md#canonical-a100ae9190b9ebd0445bb3f2051aa099db59162f9d6c7781db1ebefd2bae1506)
- [origin_servers.origin_servers.site_preferences](data-sources--dns_proxy--reference--group-001.md#canonical-8c0e123b7b205156804e291f333bc5323e0bdcb1dc72e572ea3d13377d63c352)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf7cae6db67a0aea9e0883c85039840bb50107de7e96917e6f2c6a20db6b20a5"></a>

## origin_servers.origin_servers.k8s_service — origin_servers.origin_servers.k8s_service / 3b75a53122b3 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- origin_servers.origin_servers.k8s_service

<a id="canonical-1f24b75f7c722e61a4b1c6b897b91352627e864a54d8ec278f699e944eeca669"></a>

Type: `"single"`. Computed.

Specify origin server with K8s service name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

<a id="canonical-26450776b90b94e883b661f091a9303beac6856457d676f60305fecb83d4e79e"></a>

## Direct properties — origin_servers.origin_servers.k8s_service / 3b75a53122b3 / 3

- [inside_network](data-sources--dns_proxy--reference--group-001.md#canonical-1ce5bf02b385e41f44fbf29dbcfbd61641e0d4485cc23d67497d1225cb3a6b76): complete subsection reference.

- [outside_network](data-sources--dns_proxy--reference--group-001.md#canonical-132b27324ae418e15e6edd2bfae94c5c1a257830febd1c41766e11a4a13dc1e3): complete subsection reference.

<a id="canonical-768044e37707c0b8a08ea6696ae3409c838ffe464cb9da228f21abadb9268960"></a>

<a id="canonical-8a09c0a6c2b07d0301c592d2e5e71be04dd78749922245752aa57cb5090b4a75"></a>

## protocol property — origin_servers.origin_servers.k8s_service / 3b75a53122b3 / 4

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

&#8203;- PROTOCOL\_UDP: UDP.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4932f684ac66e431e6280fae6e55e9e62e8d5ccf467d2989c7d6e8a7eafcb68d"></a>

<a id="canonical-1b9c6e8d0adc2c587e2fd040183f762d7d98e382c560c20147b1541a35bb7c59"></a>

## service_name property — origin_servers.origin_servers.k8s_service / 3b75a53122b3 / 5

Type: `"string"`. Computed.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Upstream description:

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](data-sources--dns_proxy--reference--group-001.md#canonical-7f036a0c19afed56bfe6d004997cb68e39a1dde6c80c9d6bcccf74f94055c6cc): complete subsection reference.

- [snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-ff4c559166eda98df9eb06cf5a18463ae9ea77410b2e421cd4d799563f30505e): complete subsection reference.

- [vk8s_networks](data-sources--dns_proxy--reference--group-001.md#canonical-017bb0183d42558fb7a35b48a5b0dd286ebc77384ac11a8d72643cf25e95bdd8): complete subsection reference.

<a id="canonical-7ea9f5a858a3bd04dd17a0995b2adfde93d5c7dd2862499f937cb0b2bb9feea1"></a>

## Next pages — origin_servers.origin_servers.k8s_service / 3b75a53122b3 / 6

- [origin_servers.origin_servers.k8s_service.inside_network](data-sources--dns_proxy--reference--group-001.md#canonical-1ce5bf02b385e41f44fbf29dbcfbd61641e0d4485cc23d67497d1225cb3a6b76)
- [origin_servers.origin_servers.k8s_service.outside_network](data-sources--dns_proxy--reference--group-001.md#canonical-132b27324ae418e15e6edd2bfae94c5c1a257830febd1c41766e11a4a13dc1e3)
- [origin_servers.origin_servers.k8s_service.site_locator](data-sources--dns_proxy--reference--group-001.md#canonical-7f036a0c19afed56bfe6d004997cb68e39a1dde6c80c9d6bcccf74f94055c6cc)
- [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-ff4c559166eda98df9eb06cf5a18463ae9ea77410b2e421cd4d799563f30505e)
- [origin_servers.origin_servers.k8s_service.vk8s_networks](data-sources--dns_proxy--reference--group-001.md#canonical-017bb0183d42558fb7a35b48a5b0dd286ebc77384ac11a8d72643cf25e95bdd8)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-1ce5bf02b385e41f44fbf29dbcfbd61641e0d4485cc23d67497d1225cb3a6b76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d7b20e56daf03703039056c9d5a4db4c8bf94e1d14d8268c8c4f2728e0463ea"></a>

## origin_servers.origin_servers.k8s_service.inside_network — origin_servers.origin_servers.k8s_service.inside_network / 4efcc4b48732 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- origin_servers.origin_servers.k8s_service.inside_network

<a id="canonical-90963d2cfd5a791b030742eb77fd241edaeae46f918a8413c502489fa1f6450b"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

<a id="canonical-9396d7a51725aaa7e76bdabedaf669dc9bbbd194acfe249d1a94add5633569a0"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.inside_network / 4efcc4b48732 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-76d25ddd9c55bc67eb5f10d407dfcdb52fa0afbffc20c382ed8fb69f91d43bf1"></a>

## Next pages — origin_servers.origin_servers.k8s_service.inside_network / 4efcc4b48732 / 4

- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-132b27324ae418e15e6edd2bfae94c5c1a257830febd1c41766e11a4a13dc1e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa94c8d9f4418b15542eb67d47c624c80a9f28baa19c9bb623bce330101cb995"></a>

## origin_servers.origin_servers.k8s_service.outside_network — origin_servers.origin_servers.k8s_service.outside_network / 8bf13d82ad95 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- origin_servers.origin_servers.k8s_service.outside_network

<a id="canonical-ab9b5b76b16032268df758d1ac805a08f815c6053d060748e0c564027cda620b"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-a24cd74d4f28b259b47ce113d99c7a22f5bd397a12be2edfe832df04ed992014"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.outside_network / 8bf13d82ad95 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-907057399788f72ee1b43a6a5f233496ff96eeb56ab72a1144c0e11afa134399"></a>

## Next pages — origin_servers.origin_servers.k8s_service.outside_network / 8bf13d82ad95 / 4

- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-7f036a0c19afed56bfe6d004997cb68e39a1dde6c80c9d6bcccf74f94055c6cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d771917f4e5008db12d68c4d04e9654a96f67685b92bd19bf3ab74943bc950d5"></a>

## origin_servers.origin_servers.k8s_service.site_locator — origin_servers.origin_servers.k8s_service.site_locator / 1e0832f07756 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- origin_servers.origin_servers.k8s_service.site_locator

<a id="canonical-845bce43f53f284b9546fbd101f6bd9beeca7541a6ca8c4ec70e9c1d59852986"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

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

<a id="canonical-6a42150d6259c8fde23716205bb25c9bf2b3b409d1c9867a8dc83dc148624069"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.site_locator / 1e0832f07756 / 3

- [site](data-sources--dns_proxy--reference--group-001.md#canonical-8f9937dd77595c241052aa85d2384144a697da71569c2cb8a8ee7d3d3b782991): complete subsection reference.

- [virtual_site](data-sources--dns_proxy--reference--group-001.md#canonical-f5b5e5fc55a7359ef7590a07883c08c85e7086eb3a2b3756bdc458c8501af006): complete subsection reference.

<a id="canonical-6dd92e26261d3e9b008dae13de1a4915182b6a366c97fe320728f26870f5e715"></a>

## Next pages — origin_servers.origin_servers.k8s_service.site_locator / 1e0832f07756 / 4

- [origin_servers.origin_servers.k8s_service.site_locator.site](data-sources--dns_proxy--reference--group-001.md#canonical-8f9937dd77595c241052aa85d2384144a697da71569c2cb8a8ee7d3d3b782991)
- [origin_servers.origin_servers.k8s_service.site_locator.virtual_site](data-sources--dns_proxy--reference--group-001.md#canonical-f5b5e5fc55a7359ef7590a07883c08c85e7086eb3a2b3756bdc458c8501af006)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-8f9937dd77595c241052aa85d2384144a697da71569c2cb8a8ee7d3d3b782991"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b798479af388462ea3f0651b751c935bc4c108f91321e02d25e11d99a45cfeaa"></a>

## origin_servers.origin_servers.k8s_service.site_locator.site — origin_servers.origin_servers.k8s_service.site_locator.site / 8a6835740f8f / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- [origin_servers.origin_servers.k8s_service.site_locator](data-sources--dns_proxy--reference--group-001.md#canonical-7f036a0c19afed56bfe6d004997cb68e39a1dde6c80c9d6bcccf74f94055c6cc)
- origin_servers.origin_servers.k8s_service.site_locator.site

<a id="canonical-b4d091c829e4bae161853d32b51628c97936b77d2202885cc15312fd1014af99"></a>

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

<a id="canonical-98de1efd39019dc0567c6b9c87d82214111b266adfb039000d3f75a1006c8fe0"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.site_locator.site / 8a6835740f8f / 3

<a id="canonical-5873a78e824f582d8e509a910ec52a80c4ece01e6a5a4925d4f2102dc82cecbb"></a>

<a id="canonical-abb3232a9c010aa110e350d209985b0f0192163f8369af8e4f7af561517cfc81"></a>

## name property — origin_servers.origin_servers.k8s_service.site_locator.site / 8a6835740f8f / 4

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

<a id="canonical-46025812d2af665aec60c8adccca02429c5de9db6c70ca01abdcb40eed091088"></a>

<a id="canonical-f269f4852927c62192100255221ea5d69a7797ea2ea3e919ce31f3b2c4a93156"></a>

## namespace property — origin_servers.origin_servers.k8s_service.site_locator.site / 8a6835740f8f / 5

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

<a id="canonical-b8b1a13cc54c61aeb58af63c724ca37e5390e1dfee39cf455a8558c69b910d41"></a>

<a id="canonical-1885497a567cea9f370b7d49de57c55a2f679d96c8bf768d646c599ed6e6b0a5"></a>

## tenant property — origin_servers.origin_servers.k8s_service.site_locator.site / 8a6835740f8f / 6

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

<a id="canonical-b7697d894675de26449bf79ba5b38c2f66eb8544fda79f10c6ba73e0c97cfb34"></a>

## Next pages — origin_servers.origin_servers.k8s_service.site_locator.site / 8a6835740f8f / 7

- [origin_servers.origin_servers.k8s_service.site_locator](data-sources--dns_proxy--reference--group-001.md#canonical-7f036a0c19afed56bfe6d004997cb68e39a1dde6c80c9d6bcccf74f94055c6cc)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-f5b5e5fc55a7359ef7590a07883c08c85e7086eb3a2b3756bdc458c8501af006"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8406feb95446dc3b6aa8d37a449d16eeae41f04b811ed5829db1e8defcd66358"></a>

## origin_servers.origin_servers.k8s_service.site_locator.virtual_site — origin_servers.origin_servers.k8s_service.site_locator.virtual_site / c961de7e8642 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- [origin_servers.origin_servers.k8s_service.site_locator](data-sources--dns_proxy--reference--group-001.md#canonical-7f036a0c19afed56bfe6d004997cb68e39a1dde6c80c9d6bcccf74f94055c6cc)
- origin_servers.origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-566a189fe86f083a2314b99cbecd984a5e3b65bbf247af0bcaebca80852cb067"></a>

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

<a id="canonical-79b745836a381675da8f03b5c599aeef5579e8aa06072fbaebc9278689cf6e51"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.site_locator.virtual_site / c961de7e8642 / 3

<a id="canonical-5a44860f7387d1db16be0467406b7b42c13912daff52197c4968ee0cf46ff33a"></a>

<a id="canonical-88dec5af6c42cc96759de9d5c907088da53904b070e30cf492acff8fe05551ae"></a>

## name property — origin_servers.origin_servers.k8s_service.site_locator.virtual_site / c961de7e8642 / 4

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

<a id="canonical-4a45008ed87364952e9fd670ce3c505958e4b417a5d0986d1456801011db2d4c"></a>

<a id="canonical-7cc5991490b5deac9142360700bb5a0b462d6b9dc12d3e3ed3d40d8ab908d21b"></a>

## namespace property — origin_servers.origin_servers.k8s_service.site_locator.virtual_site / c961de7e8642 / 5

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

<a id="canonical-349d2008602b8c11eb0759fd36276fe66115519a802a18a9d91ece252211677a"></a>

<a id="canonical-37f94720bff2bb722c3ce164b517917a06d1f3ca8c2ecdf46dcf0ccb238758c6"></a>

## tenant property — origin_servers.origin_servers.k8s_service.site_locator.virtual_site / c961de7e8642 / 6

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

<a id="canonical-f901ef4bf7af4ddb87cf567268e1e0b81dc3acb803a7a8def2467a098d94babd"></a>

## Next pages — origin_servers.origin_servers.k8s_service.site_locator.virtual_site / c961de7e8642 / 7

- [origin_servers.origin_servers.k8s_service.site_locator](data-sources--dns_proxy--reference--group-001.md#canonical-7f036a0c19afed56bfe6d004997cb68e39a1dde6c80c9d6bcccf74f94055c6cc)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-ff4c559166eda98df9eb06cf5a18463ae9ea77410b2e421cd4d799563f30505e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50331f00bf1284903a229b8d59ce3438ec7a2c0b7144661400fedf3277e70c45"></a>

## origin_servers.origin_servers.k8s_service.snat_pool — origin_servers.origin_servers.k8s_service.snat_pool / 254602a52586 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- origin_servers.origin_servers.k8s_service.snat_pool

<a id="canonical-27352dce247acde15d4525a1816200ab66bcdd32849070d7d865c5060772e0d5"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-0004c8ee36cba05b7ee3fd0898a52a16b7171df7015190bbec6ea1e111e30bf4"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.snat_pool / 254602a52586 / 3

- [no_snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-93d2cff86f4c090d9234fe2ad8b482880f1baf9694bdbfd9580135bef96be858): complete subsection reference.

- [snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-979fd9e78510b7d24682fb66c48860a1623039166bf83cbaf83b6261a1632e6d): complete subsection reference.

<a id="canonical-556c6394ec0bdcf5ba4c1a0c3312ccaefce38e197de1b2268222443f61ff0566"></a>

## Next pages — origin_servers.origin_servers.k8s_service.snat_pool / 254602a52586 / 4

- [origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-93d2cff86f4c090d9234fe2ad8b482880f1baf9694bdbfd9580135bef96be858)
- [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-979fd9e78510b7d24682fb66c48860a1623039166bf83cbaf83b6261a1632e6d)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-93d2cff86f4c090d9234fe2ad8b482880f1baf9694bdbfd9580135bef96be858"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4a6706d6b9a94456b86c4b1635db286c41d4b7e93a757bf57ff93866cdf854d"></a>

## origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool — origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool / 75676e966b4b / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-ff4c559166eda98df9eb06cf5a18463ae9ea77410b2e421cd4d799563f30505e)
- origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-064dc1bd4138d0b920b36cc0031b8a694ce5ef59ef4e8697d41cee47041d66c3"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

<a id="canonical-daed3e9f47fa6efebabecd166c6c29b057267b5131a9bd2ed1bd99623a55e547"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool / 75676e966b4b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-27ff127641808573e36d018cd8ce60dc57aa5a77acb98d5d888b9d6c83867a82"></a>

## Next pages — origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool / 75676e966b4b / 4

- [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-ff4c559166eda98df9eb06cf5a18463ae9ea77410b2e421cd4d799563f30505e)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-979fd9e78510b7d24682fb66c48860a1623039166bf83cbaf83b6261a1632e6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95baf0fb7ccb4d6603f558e88693ce5685f6fedcee995c5692efed7fd715c3d3"></a>

## origin_servers.origin_servers.k8s_service.snat_pool.snat_pool — origin_servers.origin_servers.k8s_service.snat_pool.snat_pool / dce8035e3abf / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-ff4c559166eda98df9eb06cf5a18463ae9ea77410b2e421cd4d799563f30505e)
- origin_servers.origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-adf5a6bb0e75de9a3a28e78cc47c544070ef236bed379d225580fd8eb14b7aa8"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-aa2a989fd07dbc94093990246243def32802f4efb7653b78219a5cb0ccd56644"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.snat_pool.snat_pool / dce8035e3abf / 3

<a id="canonical-c9367fdb9c78e49907dd5fcf6417c4279876517940f7b26a99ab5299ea882f50"></a>

<a id="canonical-e346919982b0f1e7a4174087e6d820b3a4b6809a459bf84316ae93f085d6f08e"></a>

## prefixes property — origin_servers.origin_servers.k8s_service.snat_pool.snat_pool / dce8035e3abf / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7ed2e741ea2da393d575f17136f98c45ea0d4f8a6c5339e36a57d9e79253e280"></a>

## Next pages — origin_servers.origin_servers.k8s_service.snat_pool.snat_pool / dce8035e3abf / 5

- [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-ff4c559166eda98df9eb06cf5a18463ae9ea77410b2e421cd4d799563f30505e)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-017bb0183d42558fb7a35b48a5b0dd286ebc77384ac11a8d72643cf25e95bdd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c604da020751c595c3a150b5817222b3691cbccc3f0f380b43c81dec7d774eee"></a>

## origin_servers.origin_servers.k8s_service.vk8s_networks — origin_servers.origin_servers.k8s_service.vk8s_networks / 86d480a26145 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- origin_servers.origin_servers.k8s_service.vk8s_networks

<a id="canonical-5816d102dd92e02105dec308ce8ab22db47b59de41cba8943584705138e3a9de"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vk8s networks.

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

<a id="canonical-6d2e85fc71af53effcd09980d6b744f8c67bc26dffb9d0640e32ae57c55eafa6"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.vk8s_networks / 86d480a26145 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-513baf5db019b313bfc0f56d5df1a32a6b550c9e89696b742d2c532a7903367e"></a>

## Next pages — origin_servers.origin_servers.k8s_service.vk8s_networks / 86d480a26145 / 4

- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-c7e46c1c8187f58bbda50f20e6d708210a5da0e277d12fac219fdefcac693fb6)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-ea080b6c3ad3ffcc664262c5df6ccd54d54691d1d30de213e3758f7eaa68954c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c13bd128dbfdf1467d9772341fa400ddd00966a24baa7e895ccf31d7e1e548b"></a>

## origin_servers.origin_servers.no_preference — origin_servers.origin_servers.no_preference / b73ecf23fd98 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- origin_servers.origin_servers.no_preference

<a id="canonical-4cb943213ed34b50186c27c0a56885a0de60aa8212ffc7fbab5f0f0e8f3db4a9"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no preference.

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

<a id="canonical-6c612ddad0e11a5d069fd8c1495ea0d99f3f358d7d1b7618605aebd73ad1c0e8"></a>

## Direct properties — origin_servers.origin_servers.no_preference / b73ecf23fd98 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-72712dd05bb525cd3f1117234c8cc9029137d34d17de410df0305cbf24872f28"></a>

## Next pages — origin_servers.origin_servers.no_preference / b73ecf23fd98 / 4

- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-5caa19762d55c8c60d29f57d2ce91d62ec2175e7229cdad1ab2c3e52a349ad5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e82af5cdc8a5968aa7470f2deb02f94d4544334b5f15e7bd56ecb2aac33f0cf"></a>

## origin_servers.origin_servers.public_ip — origin_servers.origin_servers.public_ip / d00811571e8b / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- origin_servers.origin_servers.public_ip

<a id="canonical-3904e9626ff586742b3c7b06e611ec5aaece615773c533e3e8959f3224245a68"></a>

Type: `"single"`. Computed.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-208e73af51d5b7498c730851e1c686818f8546495458649786b1b375566f76db"></a>

## Direct properties — origin_servers.origin_servers.public_ip / d00811571e8b / 3

<a id="canonical-405dc5ea9c6ec1021ef24549485e622500a2a1a20d43ab6bfaca884d7c318164"></a>

<a id="canonical-d09aceddcad47d12f3ad16a68bbc31bb4c460cf7e9d445fdacdf78f4d5f548f7"></a>

## ip property — origin_servers.origin_servers.public_ip / d00811571e8b / 4

Type: `"string"`. Computed.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Upstream description:

Exclusive with \[\] Public IPv4 address.

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

<a id="canonical-b2d45faec4333398e59d67196bc8160213a01bef7e077674b4b4de7d0ca1b9f4"></a>

## Next pages — origin_servers.origin_servers.public_ip / d00811571e8b / 5

- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-a100ae9190b9ebd0445bb3f2051aa099db59162f9d6c7781db1ebefd2bae1506"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da81daacc8aafaac64b487669d35a83dbbc80b8bf51e9e9af8e4287f71dd7d65"></a>

## origin_servers.origin_servers.public_name — origin_servers.origin_servers.public_name / 9f5ae6150b5c / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- origin_servers.origin_servers.public_name

<a id="canonical-2d0e83c7a64ebf8c22069fd4080745d1da0ca677451c1a4dd6a8f6ab51912887"></a>

Type: `"single"`. Computed.

Specify origin server with public DNS name.

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

<a id="canonical-76551bc2fc6a7a4f788fe80ecf1d320e48660b80e13ac3bf34efca74a7d0a578"></a>

## Direct properties — origin_servers.origin_servers.public_name / 9f5ae6150b5c / 3

<a id="canonical-625d344ceb8d49236becc738475e4e6cbd387722104369f76e97319a72e2ac2d"></a>

<a id="canonical-23c7409efc61830e51495858f8d519f4168bc2b85ead815f162aac09ba87db98"></a>

## dns_name property — origin_servers.origin_servers.public_name / 9f5ae6150b5c / 4

Type: `"string"`. Computed.

DNS Name. DNS Name

Upstream description:

DNS Name

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3a1bf226a19d2cd014cc0743d47ccb9fdcee77b223a5239ef40d2ea309d7becc"></a>

<a id="canonical-1840a37ae566925c6f5f560672262cc0fbb430735f69e5de3084342a4d266c8a"></a>

## refresh_interval property — origin_servers.origin_servers.public_name / 9f5ae6150b5c / 5

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-e931899f9939007c7eac77ecbd9d643545ef353c3aab4fd32b661318ac47c7d4"></a>

## Next pages — origin_servers.origin_servers.public_name / 9f5ae6150b5c / 6

- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-8c0e123b7b205156804e291f333bc5323e0bdcb1dc72e572ea3d13377d63c352"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-546a14396ef71c06f19a8e9296e87a68ed392f4e810a626e13eca56083b05074"></a>

## origin_servers.origin_servers.site_preferences — origin_servers.origin_servers.site_preferences / 182da27f0ae0 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- origin_servers.origin_servers.site_preferences

<a id="canonical-75c19670f1d19a48e4ab165faa222cbb9f1a9793db0e2da7a330ad4b407e92bf"></a>

Type: `"single"`. Computed.

Carries the references to one or more sites.

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

<a id="canonical-81d5461946be0efd4510480a38e8071e634ea99eee44a8f670ba8b0533251e96"></a>

## Direct properties — origin_servers.origin_servers.site_preferences / 182da27f0ae0 / 3

- [refs](data-sources--dns_proxy--reference--group-001.md#canonical-46f25dd70b82dab6023dd2685a847826c53ea5d714d3c49d90ae479ddada200b): complete subsection reference.

<a id="canonical-a024e920d7f13295ea2138feacaa6113d2004f64a0b7e457f534d7506bd57490"></a>

## Next pages — origin_servers.origin_servers.site_preferences / 182da27f0ae0 / 4

- [origin_servers.origin_servers.site_preferences.refs](data-sources--dns_proxy--reference--group-001.md#canonical-46f25dd70b82dab6023dd2685a847826c53ea5d714d3c49d90ae479ddada200b)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-46f25dd70b82dab6023dd2685a847826c53ea5d714d3c49d90ae479ddada200b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e98dff38d21ae8dea7ddbfc2faacd5de0d0b78dea233d0a2be6baa38fbc95a76"></a>

## origin_servers.origin_servers.site_preferences.refs — origin_servers.origin_servers.site_preferences.refs / 26e3b71ec9b6 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-d2a77c98b07e783d3807e5c7507e3f5ded52f99c7698f300f22dfe678b27f844)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-7f775a3f6d620b15740def1a724b3aee7d7522eac30069b5b03fd256057524e7)
- [origin_servers.origin_servers.site_preferences](data-sources--dns_proxy--reference--group-001.md#canonical-8c0e123b7b205156804e291f333bc5323e0bdcb1dc72e572ea3d13377d63c352)
- origin_servers.origin_servers.site_preferences.refs

<a id="canonical-1c3ee67964cba41aa446325b0024383512aa352497449dfd09e1bd4a78f841fa"></a>

Type: `"list"`. Computed.

Site References. Reference to one or more sites.

Upstream description:

Reference to one or more sites.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64"
  }
}
```

<a id="canonical-391808b82357d48ee2fbec3d81c26e0657c44afc0e12dc8da95f9a6ff718f7de"></a>

## Direct properties — origin_servers.origin_servers.site_preferences.refs / 26e3b71ec9b6 / 3

<a id="canonical-56d2bd12bf306c126a03c8ee1fc7c533915d2483b45ff2987b9c66fcea341338"></a>

<a id="canonical-c95dc86dc13f91b718495c8e49b2028fb983c59ecd5ae8d58a393ef440c17001"></a>

## name property — origin_servers.origin_servers.site_preferences.refs / 26e3b71ec9b6 / 4

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

<a id="canonical-9871f99fb97c814ea70dccb7204581075090bc5d2d9b526b6f1e770919fd9139"></a>

<a id="canonical-5e421b66ec0f32184c49df523aeac7534abe1cae5bfb8383f9d79e31bc24a137"></a>

## namespace property — origin_servers.origin_servers.site_preferences.refs / 26e3b71ec9b6 / 5

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

<a id="canonical-92d2c9fd59b630debfe56529de89c0886cdee13246e2d5cf3f09a17a7c34626f"></a>

<a id="canonical-7beb2482ef3da905c993e268ff692b29d2003bf2fc6e988b00368c4eddf6ae43"></a>

## tenant property — origin_servers.origin_servers.site_preferences.refs / 26e3b71ec9b6 / 6

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

<a id="canonical-496d1d5a0a024dfd03e5fe58fbcc8e2bba42e87f177ebbc99425f51784afe8eb"></a>

## Next pages — origin_servers.origin_servers.site_preferences.refs / 26e3b71ec9b6 / 7

- [origin_servers.origin_servers.site_preferences](data-sources--dns_proxy--reference--group-001.md#canonical-8c0e123b7b205156804e291f333bc5323e0bdcb1dc72e572ea3d13377d63c352)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-6ad138bdca9dd8c2037ea1d3f939f5ee70ff85ce14bff91ab9f5ebaba79d304c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b19d111cf2b0cf2db9fc2097a713636a4b70213e4b5efca6990f6bef3adcecf5"></a>

## protocol_inspection — protocol_inspection / ab07a285b4a7 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- protocol_inspection

<a id="canonical-14e8a5ab2a9cb962646f1fa25ded818a5281c1e8449b0be4e6f342e8ad5fd59f"></a>

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

<a id="canonical-512ad69bbab8817f1e43c63d14c0feb47114fde1e5403bb1428e27a49971f996"></a>

## Direct properties — protocol_inspection / ab07a285b4a7 / 3

<a id="canonical-e94348c0436854d7a955b1d8c30911b6540f1ab95d2fc833cb1084d54df406b3"></a>

<a id="canonical-16ee555b658bfba3003007b30a08db9d9da63a97b9e21a5787d3a54a0e1fbcf7"></a>

## name property — protocol_inspection / ab07a285b4a7 / 4

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

<a id="canonical-aa24e770c5a0ed73b074e1cfd38b75013ad32918dcbe2196401a7eaa96a1c247"></a>

<a id="canonical-d1743ac643dd5974776fef0f621b2d80aaf423fbb68727564ed2b378cf8a17f7"></a>

## namespace property — protocol_inspection / ab07a285b4a7 / 5

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

<a id="canonical-1b4af943ecc39f837e3651ef179cfd00cb1b23e7272b8b5798494cf58716ccba"></a>

<a id="canonical-b9845670eb3344f58d7beeab6c8d0097ea5d35d3d186b144330cca21b593e479"></a>

## tenant property — protocol_inspection / ab07a285b4a7 / 6

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

<a id="canonical-b6c39e43cea8d4f5b2d2e65e94a1fc1d6329381ef53a18d83e8ccfebe65453fd"></a>

## Next pages — protocol_inspection / ab07a285b4a7 / 7

- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-524fc80ed9032f73824da9140453c219e20a68f40cc1e49ea39746b059652924"></a>

## proxy_advertisement — proxy_advertisement / d5405141f516 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- proxy_advertisement

<a id="canonical-4d7a03319bb87c24886f8bd21c7a8489c284cd37fb5fb77156b7de6df939bee4"></a>

Type: `"single"`. Computed.

Configuration parameter for proxy advertisement.

Upstream description:

Proxy Advertisement Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_dualstack_on_public\",\"advertise_on_public\",\"advertise_on_public_default_dualstack_vip\",\"advertise_on_public_default_ipv6_vip\",\"advertise_on_public_default_vip\",\"advertise_v6_on_public\",\"do_not_advertise\"]"
}
```

<a id="canonical-5c45828132a6b3249ee80a8119e8cf3d85bc761256d2cbd717a53650cbbb425c"></a>

## Direct properties — proxy_advertisement / d5405141f516 / 3

- [advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291): complete subsection reference.

- [advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-21fdcf8716231761892b7e61c7bdcd6bcac228499dc6d0cbd9efd8eb82480f99): complete subsection reference.

- [advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-f45289809fd7b085fcca1d1df5f4bc10cd0202d1393eb7ba629de991b309111c): complete subsection reference.

- [advertise_on_public_default_dualstack_vip](data-sources--dns_proxy--reference--group-002.md#canonical-5da82d70aa98c07006c1346e0d2ee465f2cb4f5bb1621b7b588ef1deb4a21f24): complete subsection reference.

- [advertise_on_public_default_ipv6_vip](data-sources--dns_proxy--reference--group-002.md#canonical-04548d7062502d76d540d22defd914edb9470bd895a0730923176e8d8ec67d58): complete subsection reference.

- [advertise_on_public_default_vip](data-sources--dns_proxy--reference--group-002.md#canonical-4ae88fb7618e9af821b1d9b75009b587b713589b606525f666c57e6e1bb0a7b8): complete subsection reference.

- [advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-cb10f0b6654a042f07f268fa0f73ecaa482009ef09967a66690c1fa4dad64ce7): complete subsection reference.

- [do_not_advertise](data-sources--dns_proxy--reference--group-002.md#canonical-467fcbabf8498a0bc252658d3c5ba41ee23f6e92f141c3f3e050d16febb486fa): complete subsection reference.

<a id="canonical-0b14b4d77c70fb0649e38785d05966c04212b6fbbb0bc4c178c5c49adf531517"></a>

## Next pages — proxy_advertisement / d5405141f516 / 4

- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-21fdcf8716231761892b7e61c7bdcd6bcac228499dc6d0cbd9efd8eb82480f99)
- [proxy_advertisement.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-f45289809fd7b085fcca1d1df5f4bc10cd0202d1393eb7ba629de991b309111c)
- [proxy_advertisement.advertise_on_public_default_dualstack_vip](data-sources--dns_proxy--reference--group-002.md#canonical-5da82d70aa98c07006c1346e0d2ee465f2cb4f5bb1621b7b588ef1deb4a21f24)
- [proxy_advertisement.advertise_on_public_default_ipv6_vip](data-sources--dns_proxy--reference--group-002.md#canonical-04548d7062502d76d540d22defd914edb9470bd895a0730923176e8d8ec67d58)
- [proxy_advertisement.advertise_on_public_default_vip](data-sources--dns_proxy--reference--group-002.md#canonical-4ae88fb7618e9af821b1d9b75009b587b713589b606525f666c57e6e1bb0a7b8)
- [proxy_advertisement.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-cb10f0b6654a042f07f268fa0f73ecaa482009ef09967a66690c1fa4dad64ce7)
- [proxy_advertisement.do_not_advertise](data-sources--dns_proxy--reference--group-002.md#canonical-467fcbabf8498a0bc252658d3c5ba41ee23f6e92f141c3f3e050d16febb486fa)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-723ef6ef03ab397e8021e659afd1dc75df7f2894addbb98062d72712b1177a78"></a>

## proxy_advertisement.advertise_custom — proxy_advertisement.advertise_custom / b9f0ee311aa2 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- proxy_advertisement.advertise_custom

<a id="canonical-fcf0a55395e10c6b278b6a03cccc9b2b9d444c2b6b5a704c94aefa216423b3b3"></a>

Type: `"single"`. Computed.

Defines a way to advertise a VIP on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

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

<a id="canonical-2ed3d42c35ccf2cc9555e5747a1ff2a85973563360acb27fb633a061822a6571"></a>

## Direct properties — proxy_advertisement.advertise_custom / b9f0ee311aa2 / 3

- [advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8): complete subsection reference.

<a id="canonical-1436f52200ff3783ece182e87debfb9c4811772643c3fc1e6e3b083dbb73d868"></a>

## Next pages — proxy_advertisement.advertise_custom / b9f0ee311aa2 / 4

- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbd1a16da91af6745f8044e37a59bcf66ea0b160c0ead94d6ea9620c1aefad1e"></a>

## proxy_advertisement.advertise_custom.advertise_where — proxy_advertisement.advertise_custom.advertise_where / 719c90d64c0e / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- proxy_advertisement.advertise_custom.advertise_where

<a id="canonical-3d631337c524759727ea6936c97c6feec543ccb4491f9035a14e77c6941628b0"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b6ed90643f40b97bd07e4592a3bfd40437fc558cdc50d43fc51de732d7759fee"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where / 719c90d64c0e / 3

- [advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-001.md#canonical-1e6f3c3032c0be1872b90c1fa8a945972be9da7c70a6be7a39a1d0b7f9b6f08d): complete subsection reference.

- [advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-ff29e9aad584e4a7470b72fd27b4d2e4aac34f8b47bc801929171b469fb73ad3): complete subsection reference.

- [advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-c5e01fcffb12a3eab6d30a8dd212340e3a4f649e74a0a0e5a1afe8bdc6e3fe1f): complete subsection reference.

<a id="canonical-627808d74fce31d3019eba286d6cb1a4a8119e5320541f34ffa154f51347f6fc"></a>

<a id="canonical-0cd50d37c0ab4083577cdc5a496bd3c1efd0e86a0dc2b250bccea5ff8b696421"></a>

## port property — proxy_advertisement.advertise_custom.advertise_where / 719c90d64c0e / 4

Type: `"number"`. Computed.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Upstream description:

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-59ff688ba3d639b0b2fe8a4a5dc948349ca36e1cb803fe27e139a2e2b8a696d4"></a>

<a id="canonical-1e1c3622cea1fb7ac95f6a30da59c728d800ce4acedd47bd0d4742f5b50c3a66"></a>

## port_ranges property — proxy_advertisement.advertise_custom.advertise_where / 719c90d64c0e / 5

Type: `"string"`. Computed.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by "-".

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

- [site](data-sources--dns_proxy--reference--group-002.md#canonical-621fe6793fe6ca6f30379d7d229bc4ce815d52ecb218219841c5cd35762a06e6): complete subsection reference.

- [use_default_port](data-sources--dns_proxy--reference--group-002.md#canonical-7a12a16bffa8a3ad33139db767a1f14f3467d4bd47171ea90fc95fb865d3fd07): complete subsection reference.

- [virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-64e5fd38895b34a6b87c36ef654132398bed54dfe40d6c9b03f8b491cfc979dc): complete subsection reference.

- [virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-0e8f2c87d5a27244330d2a8fa89a3ab266680a51ffadd3b03efc669447be6eb8): complete subsection reference.

- [virtual_site_with_vip](data-sources--dns_proxy--reference--group-002.md#canonical-0f754bb182bfae26d5bc167a8b0754deee7eb47ca231f12695a151fb2fa37ddf): complete subsection reference.

- [vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-8f16f574f366343c6077bbc105796cde89af60b22ed3ca94986db24ccc10bcdc): complete subsection reference.

<a id="canonical-11e8c7872ac02a5d8da7aac7362456c0da899443568ea7f340f629b38acb6ad8"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where / 719c90d64c0e / 6

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-001.md#canonical-1e6f3c3032c0be1872b90c1fa8a945972be9da7c70a6be7a39a1d0b7f9b6f08d)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-ff29e9aad584e4a7470b72fd27b4d2e4aac34f8b47bc801929171b469fb73ad3)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-c5e01fcffb12a3eab6d30a8dd212340e3a4f649e74a0a0e5a1afe8bdc6e3fe1f)
- [proxy_advertisement.advertise_custom.advertise_where.site](data-sources--dns_proxy--reference--group-002.md#canonical-621fe6793fe6ca6f30379d7d229bc4ce815d52ecb218219841c5cd35762a06e6)
- [proxy_advertisement.advertise_custom.advertise_where.use_default_port](data-sources--dns_proxy--reference--group-002.md#canonical-7a12a16bffa8a3ad33139db767a1f14f3467d4bd47171ea90fc95fb865d3fd07)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-64e5fd38895b34a6b87c36ef654132398bed54dfe40d6c9b03f8b491cfc979dc)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-0e8f2c87d5a27244330d2a8fa89a3ab266680a51ffadd3b03efc669447be6eb8)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](data-sources--dns_proxy--reference--group-002.md#canonical-0f754bb182bfae26d5bc167a8b0754deee7eb47ca231f12695a151fb2fa37ddf)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-8f16f574f366343c6077bbc105796cde89af60b22ed3ca94986db24ccc10bcdc)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-1e6f3c3032c0be1872b90c1fa8a945972be9da7c70a6be7a39a1d0b7f9b6f08d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-709138eadfeba1a6663a7073fb5f10fa4f5a117084c29799603b1c2fc11bd254"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 15611ed94ffc / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-ff16fa2f1e3733d3607350c268adffeae4f4acff7740c89c0a0fb930b6db61e6"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-ef98f45cbd9faee7ecc329323e15d31ba154f6e7ae1702843ebf9f146b398291"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 15611ed94ffc / 3

- [public_ip](data-sources--dns_proxy--reference--group-001.md#canonical-3dba65f8e991a7b978b6972696643d5610e2d2f7028fe18723a3d5be3d0d209a): complete subsection reference.

<a id="canonical-efe7f577e29d6272cb84a280ebb35e7ecf91a3fa18cc0e5d5add3e18e34131fb"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 15611ed94ffc / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--dns_proxy--reference--group-001.md#canonical-3dba65f8e991a7b978b6972696643d5610e2d2f7028fe18723a3d5be3d0d209a)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-3dba65f8e991a7b978b6972696643d5610e2d2f7028fe18723a3d5be3d0d209a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bfd8a2fe651995c22485192ee2fbdb06e5dee19e9eb393095ef902595ac6d63f"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 2f956c520f24 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-001.md#canonical-1e6f3c3032c0be1872b90c1fa8a945972be9da7c70a6be7a39a1d0b7f9b6f08d)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-e9f9d200ab2f3cafdaf145e770fd9e29578fbbf086f5409e3bc160119ac6538d"></a>

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

<a id="canonical-c4547a49d3e25608b62a70df41f8fc1617bf14095050256e232dd3630264bd89"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 2f956c520f24 / 3

<a id="canonical-d78b0fd58d039571f6cffff1c3b39ae227fe28e56d71b39cb8bef95e885ed6cc"></a>

<a id="canonical-ace5eae40c459554efbb718d57cabee4c0135e885ebf5d5e67610604b5c3eb05"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 2f956c520f24 / 4

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

<a id="canonical-191648537ed3b6aa822345e86468b78f899b3ba2ab9bf882c3b9b3b3f2ceb00b"></a>
