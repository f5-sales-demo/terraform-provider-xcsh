---
page_title: "xcsh_dns_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy reference."
---

# xcsh_dns_proxy reference

<a id="canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ad33f79cc664e5a3eb9796f1798cb2fda0ee9bd2386a6220a42e296012d2ee5"></a>

## Property reference — Property reference / 475c3ec64193 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- Property reference

<a id="canonical-5dd0068627694476e9e4990744601590740a13e4feb38fe38b94b3a1fb37a5fd"></a>

## Direct properties — Property reference / 475c3ec64193 / 3

<a id="canonical-d857cb57ce41ae15eaa545e00100fe94fe11964e773ea9df819777decf0db415"></a>

<a id="canonical-7648d12952e18b4bf4b672447f946a48cc08fc49cd72b42651a14ddf3a59de4a"></a>

## annotations property — Property reference / 475c3ec64193 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

- [cache_profile](resources--dns_proxy--reference--group-001.md#canonical-eb85d071c67ef46568c01b90dedce2a13826f663e6d149a847d3cd7f8ddcbe7f): complete subsection reference.

- [ddos_profile](resources--dns_proxy--reference--group-001.md#canonical-940d12d9aab7d4b9a7b381cf77483ea804176e1ea925372173f509b5ecde8b50): complete subsection reference.

<a id="canonical-b4ca901530b261d0122ef1a0b633ac5266e99510fd5d63edbd5813083ff4121e"></a>

<a id="canonical-b681fbe95c63aeb43bfaff864c9aae48d7cb7a0cc7b837884d157f4d2152d564"></a>

## description property — Property reference / 475c3ec64193 / 5

Type: `"string"`. Optional.

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

<a id="canonical-4a1f424010f822cffcaa2e71be0e005847a5796eeb5ea4f1f4ec40b782c4eae7"></a>

<a id="canonical-90c560ba90178cae483c36c2b1fd13a625126bd9739a8091231b7aebeb79c9ca"></a>

## disable property — Property reference / 475c3ec64193 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="canonical-6be3f1dfd26f012dc5cf7695497534e296051e1ebe3184bf78617f27c74e3e3e"></a>

<a id="canonical-64ed3fa4d1d48adf44806a9c7b4733284ce84327968e957b4230838d26ecc7ab"></a>

## id property — Property reference / 475c3ec64193 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](resources--dns_proxy--reference--group-001.md#canonical-e8252c028cdd773408647216dec98673b2c6ffd09483adcab4803e65fa24112d): complete subsection reference.

<a id="canonical-14397982221dbfbd12229b1ebd9086995b8b34567ab6f926670d842337be59f9"></a>

<a id="canonical-0b8d466bd3de703fdfd5f9149f81c222f3e8dbf5443bdfb369e8c6c9039baea5"></a>

## labels property — Property reference / 475c3ec64193 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

- [lb_algorithm](resources--dns_proxy--reference--group-001.md#canonical-1bc8cdb070d344e0847a1e18c240288cb06645d2cf386b157175e596c71a4f6f): complete subsection reference.

<a id="canonical-18ff1e74df2f933d0abf079338e92b7f17e9599bcce82cd69ebcf9c7f4863046"></a>

<a id="canonical-ff69e14523e34845477de1249e1d4de2339a371dfee42d6fe1f05317103941d6"></a>

## name property — Property reference / 475c3ec64193 / 9

Type: `"string"`. Required.

Name of the DNS Proxy. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
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

<a id="canonical-eac05085c6692432373dd591a93d75c850bee0b681cbf0f2fb074844cafcca71"></a>

<a id="canonical-b5d17bb1854f8fc9c87bdc94e262115547fc3208f32af395e0f187724a6225af"></a>

## namespace property — Property reference / 475c3ec64193 / 10

Type: `"string"`. Optional, Computed.

Namespace for the DNS Proxy. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95): complete subsection reference.

- [protocol_inspection](resources--dns_proxy--reference--group-001.md#canonical-e8f7dfa31ff2b22f84858d14cd4bd06fb1fdc2295ca8b19b7e7d5d275e2b68b1): complete subsection reference.

- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c): complete subsection reference.

- [timeouts](resources--dns_proxy--reference--group-002.md#canonical-fbff3022aa3c1cd24e5dd973753c2686428fb896d6a2c568f5465e06f7e025a5): complete subsection reference.

<a id="canonical-8873bdf398891547daa8122b70df3ca8cee3aeb589a9382d9381fd6d0943a368"></a>

<a id="canonical-443f4dcfe5ac79cc207571c7ba377e80c94d2354e32f0925add26541b5811087"></a>

## transport_type property — Property reference / 475c3ec64193 / 11

Type: `"string"`. Optional, Computed.

\[Enum: UDP|TCP|BothTCPAndUDP\] Transport Type - UDP: UDP - TCP: TCP - BothTCPAndUDP: Both TCP and
UDP. Possible values are \`UDP\`, \`TCP\`, \`BothTCPAndUDP\`. Defaults to \`UDP\`.

Upstream description:

Transport Type

&#8203;- UDP: UDP

&#8203;- TCP: TCP

&#8203;- BothTCPAndUDP: Both TCP and UDP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UDP",
    "TCP",
    "BothTCPAndUDP"),
}
```

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

<a id="canonical-483d66a1c5918351dd2e722e79ee6c536c4db75a31f3999e80f950aa8431f589"></a>

## All schema paths — Property reference / 475c3ec64193 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--dns_proxy--reference--group-001.md#canonical-d857cb57ce41ae15eaa545e00100fe94fe11964e773ea9df819777decf0db415) |
| `cache_profile` | [cache_profile](resources--dns_proxy--reference--group-001.md#canonical-e824d724ddc6685fcf4dc68c397864fccb4ef525dc854e6c44d85ea138558965) |
| `cache_profile.cache_size` | [cache_profile.cache_size](resources--dns_proxy--reference--group-001.md#canonical-03fd3c2fba3fb25d38f5bc20e78953258890bbe615236c2efffac704f79187ce) |
| `cache_profile.disable_cache_profile` | [cache_profile.disable_cache_profile](resources--dns_proxy--reference--group-001.md#canonical-c3d962fa50f99703550de284c52d6859b543f08db191e1a3ae9e7cd046089da5) |
| `ddos_profile` | [ddos_profile](resources--dns_proxy--reference--group-001.md#canonical-01174a9048924bcb60b929dd40eab22b7d6f791842d317f80d1403aae86a0b4c) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](resources--dns_proxy--reference--group-001.md#canonical-c4a11ac558e142da5dd7ad4438206e2de594f345598cb45c6746a4afaacaeb6c) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](resources--dns_proxy--reference--group-001.md#canonical-d7ccf6f326a8e7a6f986bd003cf800958f745877e7f272ea7913f943ecc1e6bc) |
| `description` | [description](resources--dns_proxy--reference--group-001.md#canonical-b4ca901530b261d0122ef1a0b633ac5266e99510fd5d63edbd5813083ff4121e) |
| `disable` | [disable](resources--dns_proxy--reference--group-001.md#canonical-4a1f424010f822cffcaa2e71be0e005847a5796eeb5ea4f1f4ec40b782c4eae7) |
| `id` | [id](resources--dns_proxy--reference--group-001.md#canonical-6be3f1dfd26f012dc5cf7695497534e296051e1ebe3184bf78617f27c74e3e3e) |
| `irules` | [irules](resources--dns_proxy--reference--group-001.md#canonical-fdc1950928d967fc1ed196ac461933147fa13b120c0637834640a527d94ead08) |
| `irules.name` | [irules.name](resources--dns_proxy--reference--group-001.md#canonical-922adadef3a5f74af31d6eff99a376de699435a5182f3d06502ee02712e6843a) |
| `irules.namespace` | [irules.namespace](resources--dns_proxy--reference--group-001.md#canonical-1dffa57744cb6edb9fe2a6ae56dd779e724bc3f7fb25b128267003577edc55b8) |
| `irules.tenant` | [irules.tenant](resources--dns_proxy--reference--group-001.md#canonical-9b4c071fe2e2b031df65bf24e0ffa088b47d22e6381e0edeaa7b9eafb486e018) |
| `labels` | [labels](resources--dns_proxy--reference--group-001.md#canonical-14397982221dbfbd12229b1ebd9086995b8b34567ab6f926670d842337be59f9) |
| `lb_algorithm` | [lb_algorithm](resources--dns_proxy--reference--group-001.md#canonical-34452af904125e84cf9bfaf1a304fcbdba3589a66985df7a0fc6c1afaa14e781) |
| `lb_algorithm.round_robin` | [lb_algorithm.round_robin](resources--dns_proxy--reference--group-001.md#canonical-f0ebc75653699e9a3aa979ec0b39e176d47990e02f43d17df21b2ae5247017f8) |
| `name` | [name](resources--dns_proxy--reference--group-001.md#canonical-18ff1e74df2f933d0abf079338e92b7f17e9599bcce82cd69ebcf9c7f4863046) |
| `namespace` | [namespace](resources--dns_proxy--reference--group-001.md#canonical-eac05085c6692432373dd591a93d75c850bee0b681cbf0f2fb074844cafcca71) |
| `origin_servers` | [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-e97c22e56c75b3ac8ab35a8d3348819608ccaa48428fa705d9e15fa211fd0086) |
| `origin_servers.health_checks` | [origin_servers.health_checks](resources--dns_proxy--reference--group-001.md#canonical-108e9837f9087ce3d82aed35241cf4767aa22b2192bd4baec2721ad670e2a484) |
| `origin_servers.health_checks.health_check` | [origin_servers.health_checks.health_check](resources--dns_proxy--reference--group-001.md#canonical-665cf12795a0e9b24ba97dd4499c318d05c37107f44f78b776c3ef524c1386e3) |
| `origin_servers.health_checks.health_check.dns_health_check` | [origin_servers.health_checks.health_check.dns_health_check](resources--dns_proxy--reference--group-001.md#canonical-112deec437de52965e806d02d4e10f44af89d961e31b24297d87c689263ff871) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_rcode` | [origin_servers.health_checks.health_check.dns_health_check.expected_rcode](resources--dns_proxy--reference--group-001.md#canonical-915121eafc4f405d1e75b5e34fd6c450dad0b8cef75da2bceed188cbca5da049) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_record_type` | [origin_servers.health_checks.health_check.dns_health_check.expected_record_type](resources--dns_proxy--reference--group-001.md#canonical-4f6eb74d5ef9c3d0acc9eaa1626ac285995ce06a1f7c0fbc101c837e4cf27010) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_response` | [origin_servers.health_checks.health_check.dns_health_check.expected_response](resources--dns_proxy--reference--group-001.md#canonical-8e211b194b619bcd85f1785bfdf085d2cb064986b84c8158325bea10a017fe67) |
| `origin_servers.health_checks.health_check.dns_health_check.query_name` | [origin_servers.health_checks.health_check.dns_health_check.query_name](resources--dns_proxy--reference--group-001.md#canonical-6a3c8c5e486fddfed706c33ec68ea9feec8f2a607d43cd7c57294fee79debc71) |
| `origin_servers.health_checks.health_check.dns_health_check.query_type` | [origin_servers.health_checks.health_check.dns_health_check.query_type](resources--dns_proxy--reference--group-001.md#canonical-bbf682499d331afc26eb51bb1e175e8f2299c22ba251277ae04ae26ac67fd972) |
| `origin_servers.health_checks.health_check.dns_health_check.reverse` | [origin_servers.health_checks.health_check.dns_health_check.reverse](resources--dns_proxy--reference--group-001.md#canonical-1805722794fc09938da66fc7b991d31f0dcc96aec84095792b29ee0e2bba0e92) |
| `origin_servers.health_checks.health_check.icmp_health_check` | [origin_servers.health_checks.health_check.icmp_health_check](resources--dns_proxy--reference--group-001.md#canonical-f99e1d5cebcdff4626bb4f59cb416b7bb0d802597656cec614e0dbddc237539a) |
| `origin_servers.health_checks.health_check.tcp_health_check` | [origin_servers.health_checks.health_check.tcp_health_check](resources--dns_proxy--reference--group-001.md#canonical-9471cd4aa9fdba0c8f9d208c96026f571081e5f36f511065a4576c371f2376fe) |
| `origin_servers.health_checks.health_check.tcp_health_check.expected_response` | [origin_servers.health_checks.health_check.tcp_health_check.expected_response](resources--dns_proxy--reference--group-001.md#canonical-8ca853070f3a244674883fd73119d25de38b700412f6be3e33717170082e1cce) |
| `origin_servers.health_checks.health_check.tcp_health_check.send_payload` | [origin_servers.health_checks.health_check.tcp_health_check.send_payload](resources--dns_proxy--reference--group-001.md#canonical-13254c0ce743dbb703a7b6e98af20886f6d068fe0870bb8f8bef58e09e8f19cd) |
| `origin_servers.health_checks.healthy_threshold` | [origin_servers.health_checks.healthy_threshold](resources--dns_proxy--reference--group-001.md#canonical-84bbea654bbede359efe0c739adc3a79d6d8445a6a68e5db26151e37e76b1eda) |
| `origin_servers.health_checks.interval` | [origin_servers.health_checks.interval](resources--dns_proxy--reference--group-001.md#canonical-de22a390125b7b1b16fbc7cde300c0e264dea27575d82f4d9489151f8d94c2d7) |
| `origin_servers.health_checks.timeout` | [origin_servers.health_checks.timeout](resources--dns_proxy--reference--group-001.md#canonical-2b5f54287f33de670369636906cf9b484bc172ea0cce485b2078f9b2fee7b82c) |
| `origin_servers.health_checks.unhealthy_threshold` | [origin_servers.health_checks.unhealthy_threshold](resources--dns_proxy--reference--group-001.md#canonical-927d58f3164e22b91e43f7c109178ac39a42cfdb68ea3f74d4a2d47df24d3b3e) |
| `origin_servers.origin_servers` | [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-fe35ff926481b8b1ddfa85c7c26045b4787619e26e58d275576f19477b25babb) |
| `origin_servers.origin_servers.k8s_service` | [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-b534468b088f0297d7b50a4a22b90ed89ca18d99f78fa960c00df3436583ab00) |
| `origin_servers.origin_servers.k8s_service.inside_network` | [origin_servers.origin_servers.k8s_service.inside_network](resources--dns_proxy--reference--group-001.md#canonical-3d711559d5b387c2de3e08b3290fbb92d25f222e070a2ce157a79ac94deb4922) |
| `origin_servers.origin_servers.k8s_service.outside_network` | [origin_servers.origin_servers.k8s_service.outside_network](resources--dns_proxy--reference--group-001.md#canonical-37586fcd91b25124211e13ffe841ccd9670362a0396fcfbc7c8b7d88f97f6e7f) |
| `origin_servers.origin_servers.k8s_service.protocol` | [origin_servers.origin_servers.k8s_service.protocol](resources--dns_proxy--reference--group-001.md#canonical-23e28c5ea593fc2d908ac19456ebd8cb547400012267154b499603ec2605e434) |
| `origin_servers.origin_servers.k8s_service.service_name` | [origin_servers.origin_servers.k8s_service.service_name](resources--dns_proxy--reference--group-001.md#canonical-7f9cee8959dbfc294d6bbd9a5ca6dde4311e147d3b86dc6c300df790ecb2946d) |
| `origin_servers.origin_servers.k8s_service.site_locator` | [origin_servers.origin_servers.k8s_service.site_locator](resources--dns_proxy--reference--group-001.md#canonical-1d6ddc8b833d558976eaab57fce5adbdf44208c2c579ba839909bdae28215e42) |
| `origin_servers.origin_servers.k8s_service.site_locator.site` | [origin_servers.origin_servers.k8s_service.site_locator.site](resources--dns_proxy--reference--group-001.md#canonical-c62f96a7231f40af3cb0256dc57ad18449415973e444c550746b93e80c3aa5f8) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.name` | [origin_servers.origin_servers.k8s_service.site_locator.site.name](resources--dns_proxy--reference--group-001.md#canonical-dff5687ca9ec03adfada1533beebf5da999d9deefdf98bac0e8f571f870949aa) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.namespace` | [origin_servers.origin_servers.k8s_service.site_locator.site.namespace](resources--dns_proxy--reference--group-001.md#canonical-a6578f48d06e94e03023f333b7dcf9e9697b0143291bdfc3be47abdf31fe7f2b) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.tenant` | [origin_servers.origin_servers.k8s_service.site_locator.site.tenant](resources--dns_proxy--reference--group-001.md#canonical-4f2633d414e97bf1ca2ec1ae020b2ebf656270b0c46f26b1ac6a335bd082345c) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site](resources--dns_proxy--reference--group-001.md#canonical-3a8d6dc21c5f799188143f7f32633da56200bad4f9235aed67b1914ae8051da3) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name](resources--dns_proxy--reference--group-001.md#canonical-b5366ae9678a4e3894461427541a2e0ec9497c7ba9fa42705d14dcef586a4061) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace](resources--dns_proxy--reference--group-001.md#canonical-0002eed3f7d3e5e3937a97f50ce8f704dd25a4d8d2bbb21726eaa0131ffd2037) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant](resources--dns_proxy--reference--group-001.md#canonical-f1c2b533d1bf78880cff65c034e3ec9a17a7d1eb27848c8c8e3ce7f219335085) |
| `origin_servers.origin_servers.k8s_service.snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool](resources--dns_proxy--reference--group-001.md#canonical-c61121dd58b96b5c411c51c69669dfae0c4a98f639552f9e6960fab05c3103cf) |
| `origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](resources--dns_proxy--reference--group-001.md#canonical-4437abfcf13cb678e6e89c7ce3a1fe87a78ba92b76196fc44ded63611d9390ea) |
| `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](resources--dns_proxy--reference--group-001.md#canonical-149cf961e8bc7587bc14924c8e6eeaaec02e49b07e88b91f7ee154dc2365eba3) |
| `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes](resources--dns_proxy--reference--group-001.md#canonical-ee8a97c63cb8ae672e5128c0b2f4c38cb8aaaa47bfcdbe9d5df756362aa0078d) |
| `origin_servers.origin_servers.k8s_service.vk8s_networks` | [origin_servers.origin_servers.k8s_service.vk8s_networks](resources--dns_proxy--reference--group-001.md#canonical-b5c7cd1bcefd7210438ebadf09fb21eedeeeed568b30df883e3ad4e07c6dca40) |
| `origin_servers.origin_servers.no_preference` | [origin_servers.origin_servers.no_preference](resources--dns_proxy--reference--group-001.md#canonical-32b1f156370cdc82e55cd9ed6a8fc6101731db170c24c61157e48291f0fa02ef) |
| `origin_servers.origin_servers.public_ip` | [origin_servers.origin_servers.public_ip](resources--dns_proxy--reference--group-001.md#canonical-4fc592cea90e889383c00c4dcbde5de2039c184bb0d56a8b6d1b6db01d841853) |
| `origin_servers.origin_servers.public_ip.ip` | [origin_servers.origin_servers.public_ip.ip](resources--dns_proxy--reference--group-001.md#canonical-7480a00485fc0eafd02376698b36a5be33068215bdbb9fc2431a770f665c3a7c) |
| `origin_servers.origin_servers.public_name` | [origin_servers.origin_servers.public_name](resources--dns_proxy--reference--group-001.md#canonical-b4b17519ef09676e3472b3ea20ac1a7cb80af224120369a771829408b8787320) |
| `origin_servers.origin_servers.public_name.dns_name` | [origin_servers.origin_servers.public_name.dns_name](resources--dns_proxy--reference--group-001.md#canonical-569f7752e2dd0909201cbf057a647c0a6e76ad85aa1f4d500e54a296882caf72) |
| `origin_servers.origin_servers.public_name.refresh_interval` | [origin_servers.origin_servers.public_name.refresh_interval](resources--dns_proxy--reference--group-001.md#canonical-4df8bd3cf5c85648c44f69d1248a0a1167f423968e4e744a6fb7298370a9c442) |
| `origin_servers.origin_servers.site_preferences` | [origin_servers.origin_servers.site_preferences](resources--dns_proxy--reference--group-001.md#canonical-b493c645ea5aa76d43e5465506181837f6347d3d8e7180fc16e19d23dca7df45) |
| `origin_servers.origin_servers.site_preferences.refs` | [origin_servers.origin_servers.site_preferences.refs](resources--dns_proxy--reference--group-001.md#canonical-92753cb4c256cb6560631d4b35e8fc8c662e3cf4e4e3c6c7e584a4c697b6e646) |
| `origin_servers.origin_servers.site_preferences.refs.name` | [origin_servers.origin_servers.site_preferences.refs.name](resources--dns_proxy--reference--group-001.md#canonical-385144342b26d71c7afd0ae5d93c15887421c8a40cb524df1277e73395c2436e) |
| `origin_servers.origin_servers.site_preferences.refs.namespace` | [origin_servers.origin_servers.site_preferences.refs.namespace](resources--dns_proxy--reference--group-001.md#canonical-75854368b1b6f165be4533b75cdc0727ec7a47025c5ddefe0e40dd2f2b015cc4) |
| `origin_servers.origin_servers.site_preferences.refs.tenant` | [origin_servers.origin_servers.site_preferences.refs.tenant](resources--dns_proxy--reference--group-001.md#canonical-e3396942b41563842d86e6dfec5747829233029e91c1bd190b60858647237894) |
| `protocol_inspection` | [protocol_inspection](resources--dns_proxy--reference--group-001.md#canonical-b44c2972bf086ffb652866f8beb9e9f4899c9c22bdfb53a4627ecb1d77db5e6c) |
| `protocol_inspection.name` | [protocol_inspection.name](resources--dns_proxy--reference--group-001.md#canonical-40a6cd4fc1485a9b30bb7c2631ae9566aec84bedcdd9267df0441d30d67edea5) |
| `protocol_inspection.namespace` | [protocol_inspection.namespace](resources--dns_proxy--reference--group-001.md#canonical-b9d0811b390f56b63ae4c301dee7736fea2c1e616152b8e274f213feef9166cf) |
| `protocol_inspection.tenant` | [protocol_inspection.tenant](resources--dns_proxy--reference--group-001.md#canonical-3eb3fd47f1641177d12d4d0d5648b711e737f37920577fefd6bf18265dfdec85) |
| `proxy_advertisement` | [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-0986a35b91ad8935bea950dae848ae3b9a2ca711d3404b4ced97dd8335d04650) |
| `proxy_advertisement.advertise_custom` | [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-e7796a76f01bc96d1f6cd29f5cb41fab9e2deb07fa71ae3babd76ebfadd2fc00) |
| `proxy_advertisement.advertise_custom.advertise_where` | [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-c3aebc34c594d6fb61f7023560a700b466485dca1fbcfe0b8ec3e36a61efba7f) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-a7641a638b7e2ad7d0d958b3e8bf9f6ca75e9142b47a5d99a19bfb221cd67294) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-11084f7e0f4cdf6e7666142a08909d716fe03fc1f1156d2de318f9e4076fcc39) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](resources--dns_proxy--reference--group-002.md#canonical-11094d4841274368b88bc4cc91415afab2376a73c5b17625a57080e151b0c053) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](resources--dns_proxy--reference--group-002.md#canonical-4cefeeb8898061a3fe573fdb24c517c281b42c0a47d7b2b3b390910b62effc9a) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](resources--dns_proxy--reference--group-002.md#canonical-b25816f2199967924d641dc202e6b39fcd653042fcc1cad92cc0399b18f9cebe) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-b1c364f236d2079d77104c457b2d7c8854d76cdea96a02b48e233943dabd888d) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-5572bc785157d47d707cc86fa6f795731f428b660bca7640329f72bae52788d6) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name](resources--dns_proxy--reference--group-002.md#canonical-ed4719e0ad2b20a22cb2fbbfe283d2db8c56adefd89028c5c1dfc0c1f70e2057) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](resources--dns_proxy--reference--group-002.md#canonical-a28ac0787e0339d51e9ded30eda8f426e5a19a05831f55d3692bcc1a37be817f) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](resources--dns_proxy--reference--group-002.md#canonical-fb884ae872ac642f18b7ccd1eaba791fd22db37bc9f9604a64fbed6e475dd56d) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-a6c78f7a1b02d8235566e7ce4e656d383c20351cd653474860d430d92deb1f09) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-a9872ef0ff295e13a5e00d9b527465190f208baea7e11e5a91565945bbd03180) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](resources--dns_proxy--reference--group-002.md#canonical-e112615451bf5885220006e59ff915d919a9aa587f8490b21f688eade504fbd0) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](resources--dns_proxy--reference--group-002.md#canonical-d26825daf0c1d61a576bfb20b210415d5da7a544b4f210e6fe89f28379c5e40b) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](resources--dns_proxy--reference--group-002.md#canonical-88c077842f33179ed99ab72540ef2cd9c2bcac8eb586354e8c134d034a37d5c3) |
| `proxy_advertisement.advertise_custom.advertise_where.port` | [proxy_advertisement.advertise_custom.advertise_where.port](resources--dns_proxy--reference--group-001.md#canonical-5c7ac53e39a5351b6e0b401f2d57121cf77efbaa202cc22ce81defc1dedb4e96) |
| `proxy_advertisement.advertise_custom.advertise_where.port_ranges` | [proxy_advertisement.advertise_custom.advertise_where.port_ranges](resources--dns_proxy--reference--group-001.md#canonical-cf0d6ce41469a52f5b36cf83251207d829b310cc353ee4465181d0867528f02b) |
| `proxy_advertisement.advertise_custom.advertise_where.site` | [proxy_advertisement.advertise_custom.advertise_where.site](resources--dns_proxy--reference--group-002.md#canonical-b7c58531fe2ef386616ed925862a3713e7d6b6d126aa278f80f03e53a9c8328e) |
| `proxy_advertisement.advertise_custom.advertise_where.site.ip` | [proxy_advertisement.advertise_custom.advertise_where.site.ip](resources--dns_proxy--reference--group-002.md#canonical-e15601da128dfc99ceb371f45bc9d21d0c9514d31e9ca783e88b2afec468909a) |
| `proxy_advertisement.advertise_custom.advertise_where.site.network` | [proxy_advertisement.advertise_custom.advertise_where.site.network](resources--dns_proxy--reference--group-002.md#canonical-5135e3c354d4932d78b6b94782a17ce4e586c0f871e0e194c1d1ef310f636f85) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site` | [proxy_advertisement.advertise_custom.advertise_where.site.site](resources--dns_proxy--reference--group-002.md#canonical-d9189e4a4c191823dc0ada305436900ea3e670bf4565b7ce2e8dbc12af7bdbb4) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.name` | [proxy_advertisement.advertise_custom.advertise_where.site.site.name](resources--dns_proxy--reference--group-002.md#canonical-77aff12486af39edcb0a4a8fe18d1de8eeafbef6652213c45349efc8379a8b24) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.site.site.namespace](resources--dns_proxy--reference--group-002.md#canonical-6def463b6ea00e2b552ae2ac79e74a3eb48b43bcfad0404b01ee12edb65139ae) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.site.site.tenant](resources--dns_proxy--reference--group-002.md#canonical-1003845a00905ac39fa1368262e207fa45a445abb3d9945a1294d3facc73bb5f) |
| `proxy_advertisement.advertise_custom.advertise_where.use_default_port` | [proxy_advertisement.advertise_custom.advertise_where.use_default_port](resources--dns_proxy--reference--group-002.md#canonical-5f5f98d3d133309f4aa0b1a0462bd32e2599cb8fbe061121f8196b0ba451be71) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-ca180aa443b2a8bbc323d0aaa827c63017c4471510953fe633e231bf59e1620b) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--dns_proxy--reference--group-002.md#canonical-a3a080da6bf6bf1d679090a57dc285c8b8fc0d12dd821615a71c9e06b47f54c4) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](resources--dns_proxy--reference--group-002.md#canonical-c21dbdc282a723ea9812674157896af46b68f66583b7e6a984c90b3e884849cd) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip](resources--dns_proxy--reference--group-002.md#canonical-d49927a0d8bb4e63bacf2b11dcb1b43940a8cca2e170f231831259c5035957d4) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip](resources--dns_proxy--reference--group-002.md#canonical-6c007d4d9dcdc87043df472f9954090212c2845eb86a5ebc6dd16e6781d20006) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-2a9de88859f038ac8ded12e10682a96da16a07008d95111d990b620184ffbe6b) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name](resources--dns_proxy--reference--group-002.md#canonical-c031840ba96dcdcdafd7e0ee0234712af69cfb45c5405abc91416b6c6c643fa6) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace](resources--dns_proxy--reference--group-002.md#canonical-6787810ebb2262130927f15b05f04e2a7da9a9da547a0a7c73ccf677a1d24d67) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant](resources--dns_proxy--reference--group-002.md#canonical-a20fdeb4ab0de52ec011db347e14f3306bbb31c94c121a4eeb0511439d5d7de0) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-80f7220d364accf1332d905e392a4603f713feae2747d3af7e21c1399a2d856e) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.network](resources--dns_proxy--reference--group-002.md#canonical-578d4a435d89ee492fd765f8d10c91ca82875b57fbb975f0e3edfce083fe9d58) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-09402303b6b01b815aa84d443c46ff61a60b7de64ac603ae39938a7b20c0cc88) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name](resources--dns_proxy--reference--group-002.md#canonical-c8d6931a620e726cdd38925b4b89a9e2166a787c60104a7e2566d9fad553a15f) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace](resources--dns_proxy--reference--group-002.md#canonical-07a0d737c704d127a93b0fe18632fad4b10f2bdaff8efcc67a918cba5e05e255) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant](resources--dns_proxy--reference--group-002.md#canonical-138e49748931f123595157123f5ce1d62e5feb8d33ca4d94361f31bc328dfcdb) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--dns_proxy--reference--group-002.md#canonical-ba0195b00537a658c98fdc01ca2466f9dea9e5b21bc6d130cd3a18263f9fbb21) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip](resources--dns_proxy--reference--group-002.md#canonical-9e923eca821102d5f6f93f785d8fba93b67b822c1596c159acaae132417f7eed) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network](resources--dns_proxy--reference--group-002.md#canonical-a02d4095eff145223445c3826b66bb6ec60c16612df06fc05101f93efea6e4eb) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-df4d46bad32448ed50b9d8dc0ce0878b51c94b79f545c5a817574dee7ed90666) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](resources--dns_proxy--reference--group-002.md#canonical-b86bd2ce52276a614a1c7ce0cc61c6e33c8d8f61524fd8d38304e37478881fee) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](resources--dns_proxy--reference--group-002.md#canonical-456a3a4c5b5e61cbfa61dadf9b6cd37d1a8d69ddf16d994cab7e6d0c98898d23) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](resources--dns_proxy--reference--group-002.md#canonical-497f3aeca106c9f756058fae79f15c55e2e340f0cbaffbb93483fd4806ce4eea) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-1c58c9d224bcdceac97b955e16352b9dc4c9f6fa2294881dbca976f50eaa1df3) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](resources--dns_proxy--reference--group-002.md#canonical-816dc0ebd210ac2f405d272ec230ca9abf5e29b4afb0119c5dfc682bc51cd352) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name](resources--dns_proxy--reference--group-002.md#canonical-01a5e2a26a8c19bbe58f87099d36a3fa66461477065911c64e39af3332ddf986) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace](resources--dns_proxy--reference--group-002.md#canonical-4c54e33544a462d7c433d431d47d8b87fe01fe6ae7e5acaf43d5dc6d3813b49b) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant](resources--dns_proxy--reference--group-002.md#canonical-7c81efcd329b307a7328698d378d8eefe4948e50b0359fa297f72265b16adfc9) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-282ac9db2580842dfffd6856c0507080a3b3614e351ccd2da6f590227cbf0496) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name](resources--dns_proxy--reference--group-002.md#canonical-085baa29e23b35a322523acf8e20d810188aed1e5a51cce28d9a3150b05a8289) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](resources--dns_proxy--reference--group-002.md#canonical-2e1550638810dc5e6659da8ab7de9e6c3d57d4e3c38dddfc8f99e62ec9aef5cc) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](resources--dns_proxy--reference--group-002.md#canonical-03f76c87d493070fb330c35464806b73c7513eeb4321fea25b45af6ee7454fdb) |
| `proxy_advertisement.advertise_dualstack_on_public` | [proxy_advertisement.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-c820766fcdd4fb770c94ca5c5b0fd54f15cafcf06975ec25ae613b56da319ba7) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_dualstack_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-a399c938bb8bc4d822898858f8dc4f3971560e4f446651ece4993faa82d7388a) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.name](resources--dns_proxy--reference--group-002.md#canonical-a9dd5c6d15f11d94e5e2833f807625c04df13673f9e2d1092cd2aa176874bd80) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace](resources--dns_proxy--reference--group-002.md#canonical-b047659fcd0e096390aea5439f39590cd885aa1e6c4abe0b578edc3b72c5795f) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant](resources--dns_proxy--reference--group-002.md#canonical-ca1694eaf86b6d6eaab67ef64be2eb57269c82a0f08177f558e092ff0033f687) |
| `proxy_advertisement.advertise_on_public` | [proxy_advertisement.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-54e8f45ac32b59b60a7ae822690d0b26ab76fa220d1a6bdadab1c0923a818d69) |
| `proxy_advertisement.advertise_on_public.public_ip` | [proxy_advertisement.advertise_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-c124cd17917e7d5cfdee0b0c59b8f7c7e8c01c31f9df7d4edea0c3c37f85a77c) |
| `proxy_advertisement.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_on_public.public_ip.name](resources--dns_proxy--reference--group-002.md#canonical-b78ba0d652c45de3b5e0e29ca94e4bb3cbb98c8deba1ab635b5f1640ac223e37) |
| `proxy_advertisement.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_on_public.public_ip.namespace](resources--dns_proxy--reference--group-002.md#canonical-d0654a9912e1315803ede06839c8a8ac78561d358bf9bfabbcdcbf71e50563f2) |
| `proxy_advertisement.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_on_public.public_ip.tenant](resources--dns_proxy--reference--group-002.md#canonical-52b1b257ecbfda48770df3db0ff3568b2c941a9ece75c1dbdf2c86f7d5e03b55) |
| `proxy_advertisement.advertise_on_public_default_dualstack_vip` | [proxy_advertisement.advertise_on_public_default_dualstack_vip](resources--dns_proxy--reference--group-002.md#canonical-84d2d13d5761b42c5c7e54727bc052b67717e3b2adbb8be52d66b2aa667f398a) |
| `proxy_advertisement.advertise_on_public_default_ipv6_vip` | [proxy_advertisement.advertise_on_public_default_ipv6_vip](resources--dns_proxy--reference--group-002.md#canonical-7e48d8055fe6a2f1f4fd140f85356ddc41192d400af831537768ccc5e4db340f) |
| `proxy_advertisement.advertise_on_public_default_vip` | [proxy_advertisement.advertise_on_public_default_vip](resources--dns_proxy--reference--group-002.md#canonical-b229c26418bda71329600f5016b949b398b760327bbb18e88c8da04c9c067adb) |
| `proxy_advertisement.advertise_v6_on_public` | [proxy_advertisement.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-526771c2166cf95233da0632290598a5d30732d85bd6cb7985c8fb5238dd3ff2) |
| `proxy_advertisement.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_v6_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-f024433cd07fd0cb19ee3eb5768761ab712f94233cb1360b5314239281394069) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_v6_on_public.public_ip.name](resources--dns_proxy--reference--group-002.md#canonical-1ffcb351004491851561252a220ee66520830cd6fe224ae85736de2055f7c0fc) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_v6_on_public.public_ip.namespace](resources--dns_proxy--reference--group-002.md#canonical-e374814a597d550ec06cd74f075717af41293ed89621631a73d72eac3f689819) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_v6_on_public.public_ip.tenant](resources--dns_proxy--reference--group-002.md#canonical-d00b56add2a6f0fc5431966cda55582552dcda0f0e49a45ac143822d5b32d4ea) |
| `proxy_advertisement.do_not_advertise` | [proxy_advertisement.do_not_advertise](resources--dns_proxy--reference--group-002.md#canonical-c5f205abf9af66657cec11a7dc597469311c4d34a70223a22586b71ab4272fe7) |
| `timeouts` | [timeouts](resources--dns_proxy--reference--group-002.md#canonical-997c29c6fbf12e74d56064161260a45cd928709c531cea946edcba3d30f6c84f) |
| `timeouts.create` | [timeouts.create](resources--dns_proxy--reference--group-002.md#canonical-20885ebbf11b16653e443eeec8970c4b82e7092737a15d02b1ec79e2d425a301) |
| `timeouts.delete` | [timeouts.delete](resources--dns_proxy--reference--group-002.md#canonical-4a7058d79099b37eba4630afe530deacae2cc48bf3eb179d2c8553cac4d018e8) |
| `timeouts.read` | [timeouts.read](resources--dns_proxy--reference--group-002.md#canonical-8578c36fe6e7f1710d75d1f34a97521a7629e3364861d0c4a3e5de6e4841c23e) |
| `timeouts.update` | [timeouts.update](resources--dns_proxy--reference--group-002.md#canonical-ed5cfbbef02441248562d4b562ab0be69b540b31e8754582b125f8cf2a23a3e4) |
| `transport_type` | [transport_type](resources--dns_proxy--reference--group-001.md#canonical-8873bdf398891547daa8122b70df3ca8cee3aeb589a9382d9381fd6d0943a368) |

<a id="canonical-5337243679fc9fe9c7b17fb80d7c8d88e72665fbfea14825cbf566e0f63a6201"></a>

## Next pages — Property reference / 475c3ec64193 / 13

- [cache_profile](resources--dns_proxy--reference--group-001.md#canonical-eb85d071c67ef46568c01b90dedce2a13826f663e6d149a847d3cd7f8ddcbe7f)
- [ddos_profile](resources--dns_proxy--reference--group-001.md#canonical-940d12d9aab7d4b9a7b381cf77483ea804176e1ea925372173f509b5ecde8b50)
- [irules](resources--dns_proxy--reference--group-001.md#canonical-e8252c028cdd773408647216dec98673b2c6ffd09483adcab4803e65fa24112d)
- [lb_algorithm](resources--dns_proxy--reference--group-001.md#canonical-1bc8cdb070d344e0847a1e18c240288cb06645d2cf386b157175e596c71a4f6f)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [protocol_inspection](resources--dns_proxy--reference--group-001.md#canonical-e8f7dfa31ff2b22f84858d14cd4bd06fb1fdc2295ca8b19b7e7d5d275e2b68b1)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [timeouts](resources--dns_proxy--reference--group-002.md#canonical-fbff3022aa3c1cd24e5dd973753c2686428fb896d6a2c568f5465e06f7e025a5)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-eb85d071c67ef46568c01b90dedce2a13826f663e6d149a847d3cd7f8ddcbe7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6055d3c0269d2882e7a62f23d3a2e9b9b322f93afe957aeacce2045d06471e22"></a>

## cache_profile — cache_profile / 38dfd3aed999 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- cache_profile

<a id="canonical-e824d724ddc6685fcf4dc68c397864fccb4ef525dc854e6c44d85ea138558965"></a>

Type: `"object"`. single nested block, Optional.

DNS Cache specifies cache configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cache_size",
    "disable_cache_profile")}
```

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

Terraform syntax:

```terraform
cache_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-23f7d4c9510635d4d7949685f08ce51b386620b080840a50f9b2921f14c94afe"></a>

## Direct properties — cache_profile / 38dfd3aed999 / 3

<a id="canonical-03fd3c2fba3fb25d38f5bc20e78953258890bbe615236c2efffac704f79187ce"></a>

<a id="canonical-c74f8a3275fc7157faf4edd53839b8b4d10af75afc958e5be4f43bbf5d02d0c1"></a>

## cache_size property — cache_profile / 38dfd3aed999 / 4

Type: `"number"`. Optional.

Exclusive with \[disable\_cache\_profile\] cache size.

Upstream description:

Exclusive with \[disable\_cache\_profile\] cache size.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 10240),
}
```

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

- [disable_cache_profile](resources--dns_proxy--reference--group-001.md#canonical-2631a57401e813bfd5a30706568a3ff5c161a1ceb539751b54c8c7b7310341ac): complete subsection reference.

<a id="canonical-31f80b500523ecbf814952f2406ba96e6149c0387ebb3598ff78ddf8e0e4f911"></a>

## Next pages — cache_profile / 38dfd3aed999 / 5

- [cache_profile.disable_cache_profile](resources--dns_proxy--reference--group-001.md#canonical-2631a57401e813bfd5a30706568a3ff5c161a1ceb539751b54c8c7b7310341ac)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-2631a57401e813bfd5a30706568a3ff5c161a1ceb539751b54c8c7b7310341ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b37b5e29416ac27cb1499962677c3aaae39c0bbd3534ad22a33121458e436bc7"></a>

## cache_profile.disable_cache_profile — cache_profile.disable_cache_profile / 5ba00c584d9d / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [cache_profile](resources--dns_proxy--reference--group-001.md#canonical-eb85d071c67ef46568c01b90dedce2a13826f663e6d149a847d3cd7f8ddcbe7f)
- cache_profile.disable_cache_profile

<a id="canonical-c3d962fa50f99703550de284c52d6859b543f08db191e1a3ae9e7cd046089da5"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_cache_profile = {}
```

<a id="canonical-8486636a00dfe06e410330ea3772df90b5ae8762e63aee267d72c1145a8cb208"></a>

## Direct properties — cache_profile.disable_cache_profile / 5ba00c584d9d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-22aa423a3ff08f127e2fa3ac3edbccd767eaebb0bfcd053e0a681cfc719bcff1"></a>

## Next pages — cache_profile.disable_cache_profile / 5ba00c584d9d / 4

- [cache_profile](resources--dns_proxy--reference--group-001.md#canonical-eb85d071c67ef46568c01b90dedce2a13826f663e6d149a847d3cd7f8ddcbe7f)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-940d12d9aab7d4b9a7b381cf77483ea804176e1ea925372173f509b5ecde8b50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a23a41e70f0bd896dddb51fd58de597830b1e637eed4acf91cb5dcaf10fe9cec"></a>

## ddos_profile — ddos_profile / 02b6032c3c52 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- ddos_profile

<a id="canonical-01174a9048924bcb60b929dd40eab22b7d6f791842d317f80d1403aae86a0b4c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ddos profile.

Upstream description:

DDoS Protection Rule for DNS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_ddos_mitigation",
    "enable_ddos_mitigation")}
```

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

Terraform syntax:

```terraform
ddos_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-c6b2fed02748ccef72a37cdd9f0089024e37d2558c8a324bdd18da9df3d8e6b9"></a>

## Direct properties — ddos_profile / 02b6032c3c52 / 3

- [disable_ddos_mitigation](resources--dns_proxy--reference--group-001.md#canonical-ab3904f9c3fe8a7b369989044454961ff8a363d17df54592e3a5c5283e7b3223): complete subsection reference.

- [enable_ddos_mitigation](resources--dns_proxy--reference--group-001.md#canonical-17e8f51ac418d14b18334e9745c2c562c1ce6437c8572aa3bf460ab92b05cdc7): complete subsection reference.

<a id="canonical-4ee5434da20580232839ec243996c84ca961ea37af5708fe848d5ee50bb459e4"></a>

## Next pages — ddos_profile / 02b6032c3c52 / 4

- [ddos_profile.disable_ddos_mitigation](resources--dns_proxy--reference--group-001.md#canonical-ab3904f9c3fe8a7b369989044454961ff8a363d17df54592e3a5c5283e7b3223)
- [ddos_profile.enable_ddos_mitigation](resources--dns_proxy--reference--group-001.md#canonical-17e8f51ac418d14b18334e9745c2c562c1ce6437c8572aa3bf460ab92b05cdc7)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-ab3904f9c3fe8a7b369989044454961ff8a363d17df54592e3a5c5283e7b3223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6b73dad6a02bd56ede9dfd5524731206a8c48ab5a60efe8b733279ad0f2307f"></a>

## ddos_profile.disable_ddos_mitigation — ddos_profile.disable_ddos_mitigation / 5c242920119e / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [ddos_profile](resources--dns_proxy--reference--group-001.md#canonical-940d12d9aab7d4b9a7b381cf77483ea804176e1ea925372173f509b5ecde8b50)
- ddos_profile.disable_ddos_mitigation

<a id="canonical-c4a11ac558e142da5dd7ad4438206e2de594f345598cb45c6746a4afaacaeb6c"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_ddos_mitigation = {}
```

<a id="canonical-eed018982395dde1c12070df2b561ca0792d7b74e3e5a85d8368986fcd1d3575"></a>

## Direct properties — ddos_profile.disable_ddos_mitigation / 5c242920119e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-42df428cad12af04c2be44c28b737cdae28828c80aa83fe30aaa1d4a4f02f8f0"></a>

## Next pages — ddos_profile.disable_ddos_mitigation / 5c242920119e / 4

- [ddos_profile](resources--dns_proxy--reference--group-001.md#canonical-940d12d9aab7d4b9a7b381cf77483ea804176e1ea925372173f509b5ecde8b50)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-17e8f51ac418d14b18334e9745c2c562c1ce6437c8572aa3bf460ab92b05cdc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe86f4646e5d5f7db11541941aba63384f291fb3ca67c686e09b260bed0b875c"></a>

## ddos_profile.enable_ddos_mitigation — ddos_profile.enable_ddos_mitigation / 75c80b9d0647 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [ddos_profile](resources--dns_proxy--reference--group-001.md#canonical-940d12d9aab7d4b9a7b381cf77483ea804176e1ea925372173f509b5ecde8b50)
- ddos_profile.enable_ddos_mitigation

<a id="canonical-d7ccf6f326a8e7a6f986bd003cf800958f745877e7f272ea7913f943ecc1e6bc"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable_ddos_mitigation = {}
```

<a id="canonical-b293c3f4ea0c6ffc5756aad9101a465f087f5f66e57edd2df6bed6b764896b88"></a>

## Direct properties — ddos_profile.enable_ddos_mitigation / 75c80b9d0647 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-45ff31dcc3ece4a3f7ae17a16024630588fb302e1b42c3506507229be013829c"></a>

## Next pages — ddos_profile.enable_ddos_mitigation / 75c80b9d0647 / 4

- [ddos_profile](resources--dns_proxy--reference--group-001.md#canonical-940d12d9aab7d4b9a7b381cf77483ea804176e1ea925372173f509b5ecde8b50)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-e8252c028cdd773408647216dec98673b2c6ffd09483adcab4803e65fa24112d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2917788a202279d88e90047ca33f197963e35b15755cf121ad5da9eca6255bd1"></a>

## irules — irules / 0a4acba48b61 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- irules

<a id="canonical-fdc1950928d967fc1ed196ac461933147fa13b120c0637834640a527d94ead08"></a>

Type: `"object"`. list nested block, Optional.

OPTIONS for attaching iRules to DNS proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
irules {
  # Configure direct properties listed below.
}
```

<a id="canonical-ceb23ba55177e1947ed4c10930417cd3c045dfcdb92ec27e32896b442b4ff3e6"></a>

## Direct properties — irules / 0a4acba48b61 / 3

<a id="canonical-922adadef3a5f74af31d6eff99a376de699435a5182f3d06502ee02712e6843a"></a>

<a id="canonical-4e952645fca9e5c0e5374d940f67499acaae88958e0a73002a75c411ec038dfa"></a>

## name property — irules / 0a4acba48b61 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-1dffa57744cb6edb9fe2a6ae56dd779e724bc3f7fb25b128267003577edc55b8"></a>

<a id="canonical-dd0f61438c538aca75bc16c7bd400422a504bf809bacefa4b385e8ea4a0a636b"></a>

## namespace property — irules / 0a4acba48b61 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-9b4c071fe2e2b031df65bf24e0ffa088b47d22e6381e0edeaa7b9eafb486e018"></a>

<a id="canonical-85ef39849d0a7069e6bdbddf33d4273a05c20f1a10512b9f34f7e9432e51b0cc"></a>

## tenant property — irules / 0a4acba48b61 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-7607d08435c2d008e3e3fac29e8e7618fc94444c4bc2e3efebc881e0bb2e1226"></a>

## Next pages — irules / 0a4acba48b61 / 7

- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-1bc8cdb070d344e0847a1e18c240288cb06645d2cf386b157175e596c71a4f6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5e7fc2cb79137877b0c258dee804fd2a9dd80d7a0b23555cdc401fd41bfa347"></a>

## lb_algorithm — lb_algorithm / 28557ec82363 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- lb_algorithm

<a id="canonical-34452af904125e84cf9bfaf1a304fcbdba3589a66985df7a0fc6c1afaa14e781"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
lb_algorithm {
  # Configure direct properties listed below.
}
```

<a id="canonical-97471e778f3d0e4c8609024f618663b47a9d9e0febb98754cb7559bfdd1bb67f"></a>

## Direct properties — lb_algorithm / 28557ec82363 / 3

- [round_robin](resources--dns_proxy--reference--group-001.md#canonical-bead2e5f65fa403d37c3d1d596f91ffdbcb42efe8befc2db944894feb9e752ae): complete subsection reference.

<a id="canonical-5b53156a2ee725547b7644c4170f874a4b9f1ce19eb6daad973fb27a11ef6be3"></a>

## Next pages — lb_algorithm / 28557ec82363 / 4

- [lb_algorithm.round_robin](resources--dns_proxy--reference--group-001.md#canonical-bead2e5f65fa403d37c3d1d596f91ffdbcb42efe8befc2db944894feb9e752ae)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-bead2e5f65fa403d37c3d1d596f91ffdbcb42efe8befc2db944894feb9e752ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ea7211fc88a3e71697aa62e0906f26c92da7f753f697438b1c723c301889b42"></a>

## lb_algorithm.round_robin — lb_algorithm.round_robin / 4b826876b94d / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [lb_algorithm](resources--dns_proxy--reference--group-001.md#canonical-1bc8cdb070d344e0847a1e18c240288cb06645d2cf386b157175e596c71a4f6f)
- lb_algorithm.round_robin

<a id="canonical-f0ebc75653699e9a3aa979ec0b39e176d47990e02f43d17df21b2ae5247017f8"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
round_robin {}
```

<a id="canonical-ed8ce5588eb4cb3235ba061242e6d536b289cd0f07ceaae7c0bb5194bb902370"></a>

## Direct properties — lb_algorithm.round_robin / 4b826876b94d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0b17072725cbf25be90867c88f7f79045c870ee9d3372ee59592e19f98f61ea5"></a>

## Next pages — lb_algorithm.round_robin / 4b826876b94d / 4

- [lb_algorithm](resources--dns_proxy--reference--group-001.md#canonical-1bc8cdb070d344e0847a1e18c240288cb06645d2cf386b157175e596c71a4f6f)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99bf0dbc612a4740e2196f9b7a4e1483041c7852f921fea3027257a8e211cc6f"></a>

## origin_servers — origin_servers / d96669a70bcc / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- origin_servers

<a id="canonical-e97c22e56c75b3ac8ab35a8d3348819608ccaa48428fa705d9e15fa211fd0086"></a>

Type: `"object"`. single nested block, Optional.

List of origin Servers for the DNS proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_servers")}
```

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

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-6ba63209b7caca00bc141fa9d75725df16d2a6b2d676abeae19decabf26b49fd"></a>

## Direct properties — origin_servers / d96669a70bcc / 3

- [health_checks](resources--dns_proxy--reference--group-001.md#canonical-2bd317751fa784cb9ff8a110775df20082b25dc4631e8946f6ce5f8a9c5d0bac): complete subsection reference.

- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea): complete subsection reference.

<a id="canonical-2285025b1b95dcce55233f42da36c69b1ce99581b0a82053a7c2718f6da58ba0"></a>

## Next pages — origin_servers / d96669a70bcc / 4

- [origin_servers.health_checks](resources--dns_proxy--reference--group-001.md#canonical-2bd317751fa784cb9ff8a110775df20082b25dc4631e8946f6ce5f8a9c5d0bac)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-2bd317751fa784cb9ff8a110775df20082b25dc4631e8946f6ce5f8a9c5d0bac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6b72a349d7e7bb3a6f158e80cfc11bf19baeb73e0b9b6215fd229aa0ea8290c"></a>

## origin_servers.health_checks — origin_servers.health_checks / 43d3af40f571 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- origin_servers.health_checks

<a id="canonical-108e9837f9087ce3d82aed35241cf4767aa22b2192bd4baec2721ad670e2a484"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for health checks.

Upstream description:

Origin Server Health Checks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check",
    "healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold")}
```

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

Terraform syntax:

```terraform
health_checks {
  # Configure direct properties listed below.
}
```

<a id="canonical-64d2f7ccf565ea765b8448202f618ab2c115f5c02586d691448fcf49dca56063"></a>

## Direct properties — origin_servers.health_checks / 43d3af40f571 / 3

- [health_check](resources--dns_proxy--reference--group-001.md#canonical-054acab64a398d69cc22dc684f7d6a1266f1b0a5c5de05139505f8a5c039916c): complete subsection reference.

<a id="canonical-84bbea654bbede359efe0c739adc3a79d6d8445a6a68e5db26151e37e76b1eda"></a>

<a id="canonical-3818a69899decb29ab754453c35721d34f1a9123c83d6efee1368b24581fc767"></a>

## healthy_threshold property — origin_servers.health_checks / 43d3af40f571 / 4

Type: `"number"`. Optional.

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

Upstream description:

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

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

<a id="canonical-de22a390125b7b1b16fbc7cde300c0e264dea27575d82f4d9489151f8d94c2d7"></a>

<a id="canonical-5bd3fc1f85e8401f39245a74e8c9a5e1b15b38bae8ba2925a8f3733ae6907102"></a>

## interval property — origin_servers.health_checks / 43d3af40f571 / 5

Type: `"number"`. Optional.

Time interval in seconds between two healthcheck requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

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

<a id="canonical-2b5f54287f33de670369636906cf9b484bc172ea0cce485b2078f9b2fee7b82c"></a>

<a id="canonical-795a798606427e824c6df401fbce066e5109d1aa3a372dec9fd502a162073fe4"></a>

## timeout property — origin_servers.health_checks / 43d3af40f571 / 6

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

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

<a id="canonical-927d58f3164e22b91e43f7c109178ac39a42cfdb68ea3f74d4a2d47df24d3b3e"></a>

<a id="canonical-4674507eed8d92f2f72c1dd8722219359755c7fa15a694cb594d88ec6a1476a0"></a>

## unhealthy_threshold property — origin_servers.health_checks / 43d3af40f571 / 7

Type: `"number"`. Optional.

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health
checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

Upstream description:

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health
checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

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

<a id="canonical-c21a03f8644523e3f8f661c904d37740289683bbafb8511e3787a1177c79360a"></a>

## Next pages — origin_servers.health_checks / 43d3af40f571 / 8

- [origin_servers.health_checks.health_check](resources--dns_proxy--reference--group-001.md#canonical-054acab64a398d69cc22dc684f7d6a1266f1b0a5c5de05139505f8a5c039916c)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-054acab64a398d69cc22dc684f7d6a1266f1b0a5c5de05139505f8a5c039916c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c38c6d05e109c2498b5c6d41b12d0eb26838c9a5a63c57a2f9c1ed1734b555c7"></a>

## origin_servers.health_checks.health_check — origin_servers.health_checks.health_check / 1e6ce0c320ee / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.health_checks](resources--dns_proxy--reference--group-001.md#canonical-2bd317751fa784cb9ff8a110775df20082b25dc4631e8946f6ce5f8a9c5d0bac)
- origin_servers.health_checks.health_check

<a id="canonical-665cf12795a0e9b24ba97dd4499c318d05c37107f44f78b776c3ef524c1386e3"></a>

Type: `"object"`. list nested block, Optional.

List of Health Checks. List of Health Checks.

Upstream description:

List of Health Checks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dns_health_check",
    "icmp_health_check"),
  validators.ConflictingListObjectAttributes("dns_health_check",
    "tcp_health_check"),
  validators.ConflictingListObjectAttributes("icmp_health_check",
    "tcp_health_check")}
```

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

Terraform syntax:

```terraform
health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-6223afb129b3e30815c7747f15c511358ca7c10be5ee8f55176710bdc047a389"></a>

## Direct properties — origin_servers.health_checks.health_check / 1e6ce0c320ee / 3

- [dns_health_check](resources--dns_proxy--reference--group-001.md#canonical-3508b247db99d1c6ee48b6be5f7aaaee90b7f5eb3b62b1f9eb3e881e90fdfda9): complete subsection reference.

- [icmp_health_check](resources--dns_proxy--reference--group-001.md#canonical-714d06b502f2dfef5f23b44b19354cb47c9b3d8be51e890bc32ecc0cad97f68f): complete subsection reference.

- [tcp_health_check](resources--dns_proxy--reference--group-001.md#canonical-d36eef33355c945bfd324232c330465a23929a5ddf99aa527bfdb4cd5abdaf0d): complete subsection reference.

<a id="canonical-664cd207a2aa0f4744ef197f6832bdc4d4586e5c9395d7375e6fedd3812c47db"></a>

## Next pages — origin_servers.health_checks.health_check / 1e6ce0c320ee / 4

- [origin_servers.health_checks.health_check.dns_health_check](resources--dns_proxy--reference--group-001.md#canonical-3508b247db99d1c6ee48b6be5f7aaaee90b7f5eb3b62b1f9eb3e881e90fdfda9)
- [origin_servers.health_checks.health_check.icmp_health_check](resources--dns_proxy--reference--group-001.md#canonical-714d06b502f2dfef5f23b44b19354cb47c9b3d8be51e890bc32ecc0cad97f68f)
- [origin_servers.health_checks.health_check.tcp_health_check](resources--dns_proxy--reference--group-001.md#canonical-d36eef33355c945bfd324232c330465a23929a5ddf99aa527bfdb4cd5abdaf0d)
- [origin_servers.health_checks](resources--dns_proxy--reference--group-001.md#canonical-2bd317751fa784cb9ff8a110775df20082b25dc4631e8946f6ce5f8a9c5d0bac)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-3508b247db99d1c6ee48b6be5f7aaaee90b7f5eb3b62b1f9eb3e881e90fdfda9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca87f863bdbb58e4892c63a28e4b5eda85b1c0fb88d3d578c42409bb33829a45"></a>

## origin_servers.health_checks.health_check.dns_health_check — origin_servers.health_checks.health_check.dns_health_check / a8d895c0180d / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.health_checks](resources--dns_proxy--reference--group-001.md#canonical-2bd317751fa784cb9ff8a110775df20082b25dc4631e8946f6ce5f8a9c5d0bac)
- [origin_servers.health_checks.health_check](resources--dns_proxy--reference--group-001.md#canonical-054acab64a398d69cc22dc684f7d6a1266f1b0a5c5de05139505f8a5c039916c)
- origin_servers.health_checks.health_check.dns_health_check

<a id="canonical-112deec437de52965e806d02d4e10f44af89d961e31b24297d87c689263ff871"></a>

Type: `"object"`. single nested block, Optional.

DNS health check reports healthy if DNS query is successful and response header and answer matches
the given value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expected_response",
    "query_name")}
```

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

Terraform syntax:

```terraform
dns_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-504ad648914531b79502dc86a0389ac348ff9eeae5c03ba7a8bf5fb6dcb1b8de"></a>

## Direct properties — origin_servers.health_checks.health_check.dns_health_check / a8d895c0180d / 3

<a id="canonical-915121eafc4f405d1e75b5e34fd6c450dad0b8cef75da2bceed188cbca5da049"></a>

<a id="canonical-0e24fbf5a329daab4667eb3f9de014188a4ac24bd19c226447949ad1aeeb98b9"></a>

## expected_rcode property — origin_servers.health_checks.health_check.dns_health_check / a8d895c0180d / 4

Type: `"string"`. Optional.

\[Enum: DNS\_RES\_RCODE\_NOERROR|DNS\_RES\_RCODE\_ANY\] Expected DNS Response Rcode Type -
DNS\_RES\_RCODE\_NOERROR: Rcode NOERROR - DNS\_RES\_RCODE\_ANY: RCODE ANY. Possible values are
\`DNS\_RES\_RCODE\_NOERROR\`, \`DNS\_RES\_RCODE\_ANY\`. Defaults to \`DNS\_RES\_RCODE\_NOERROR\`.

Upstream description:

Expected DNS Response Rcode Type

&#8203;- DNS\_RES\_RCODE\_NOERROR: Rcode NOERROR

&#8203;- DNS\_RES\_RCODE\_ANY: RCODE ANY.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DNS_RES_RCODE_NOERROR",
    "DNS_RES_RCODE_ANY"),
}
```

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

<a id="canonical-4f6eb74d5ef9c3d0acc9eaa1626ac285995ce06a1f7c0fbc101c837e4cf27010"></a>

<a id="canonical-28bdd93378cd5a2a48ee439592fed5f80169a3e5b27e3a6a13a299c8ba7247ae"></a>

## expected_record_type property — origin_servers.health_checks.health_check.dns_health_check / a8d895c0180d / 5

Type: `"string"`. Optional.

\[Enum: DNS\_REQUESTED\_QUERY\_TYPE|DNS\_RES\_RECORD\_TYPE\_ANY\] DNS Response Record Type -
DNS\_REQUESTED\_QUERY\_TYPE: Requested Query Type - DNS\_RES\_RECORD\_TYPE\_ANY: Any Record Type.
Possible values are \`DNS\_REQUESTED\_QUERY\_TYPE\`, \`DNS\_RES\_RECORD\_TYPE\_ANY\`. Defaults to
\`DNS\_REQUESTED\_QUERY\_TYPE\`.

Upstream description:

DNS Response Record Type

&#8203;- DNS\_REQUESTED\_QUERY\_TYPE: Requested Query Type

&#8203;- DNS\_RES\_RECORD\_TYPE\_ANY: Any Record Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DNS_REQUESTED_QUERY_TYPE",
    "DNS_RES_RECORD_TYPE_ANY"),
}
```

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

<a id="canonical-8e211b194b619bcd85f1785bfdf085d2cb064986b84c8158325bea10a017fe67"></a>

<a id="canonical-ddfee4918556fc1686effe08b6eaef1808b377b5805509b103dbcc9094533b55"></a>

## expected_response property — origin_servers.health_checks.health_check.dns_health_check / a8d895c0180d / 6

Type: `"string"`. Optional.

Specifies an IPv4 or IPv6 address in the answer section of DNS Response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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

<a id="canonical-6a3c8c5e486fddfed706c33ec68ea9feec8f2a607d43cd7c57294fee79debc71"></a>

<a id="canonical-8e55d23a2c70ec8fcf271ada551c7dabcec6c927e73a72b93707f7089b454db4"></a>

## query_name property — origin_servers.health_checks.health_check.dns_health_check / a8d895c0180d / 7

Type: `"string"`. Optional.

The query name that the monitor sends a DNS query for.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 255),
}
```

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

<a id="canonical-bbf682499d331afc26eb51bb1e175e8f2299c22ba251277ae04ae26ac67fd972"></a>

<a id="canonical-8bf3e2d926875da442ebe875e913cbf6934f9e3648f6a3c550f614532cc0e432"></a>

## query_type property — origin_servers.health_checks.health_check.dns_health_check / a8d895c0180d / 8

Type: `"string"`. Optional.

\[Enum: DNS\_QTYPE\_A|DNS\_QTYPE\_AAAA\] DNS Query Type - DNS\_QTYPE\_A: Query Type A -
DNS\_QTYPE\_AAAA: Query Type AAAA. Possible values are \`DNS\_QTYPE\_A\`, \`DNS\_QTYPE\_AAAA\`.
Defaults to \`DNS\_QTYPE\_A\`.

Upstream description:

DNS Query Type

&#8203;- DNS\_QTYPE\_A: Query Type A

&#8203;- DNS\_QTYPE\_AAAA: Query Type AAAA.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DNS_QTYPE_A",
    "DNS_QTYPE_AAAA"),
}
```

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

<a id="canonical-1805722794fc09938da66fc7b991d31f0dcc96aec84095792b29ee0e2bba0e92"></a>

<a id="canonical-ce934bcc0d35e519f2b886bdfb7e31f4337eeac003de2c566db5aed43e454794"></a>

## reverse property — origin_servers.health_checks.health_check.dns_health_check / a8d895c0180d / 9

Type: `"bool"`. Optional.

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

<a id="canonical-f27416864174b9a7c165093078a6ac9de9d3dc2599f6772265f90df52603c113"></a>

## Next pages — origin_servers.health_checks.health_check.dns_health_check / a8d895c0180d / 10

- [origin_servers.health_checks.health_check](resources--dns_proxy--reference--group-001.md#canonical-054acab64a398d69cc22dc684f7d6a1266f1b0a5c5de05139505f8a5c039916c)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-714d06b502f2dfef5f23b44b19354cb47c9b3d8be51e890bc32ecc0cad97f68f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f8740b78704363f57c3fa7d9ce0467c31b47f7ad18c731abc1a6be2ae07ea76"></a>

## origin_servers.health_checks.health_check.icmp_health_check — origin_servers.health_checks.health_check.icmp_health_check / b7045d18f139 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.health_checks](resources--dns_proxy--reference--group-001.md#canonical-2bd317751fa784cb9ff8a110775df20082b25dc4631e8946f6ce5f8a9c5d0bac)
- [origin_servers.health_checks.health_check](resources--dns_proxy--reference--group-001.md#canonical-054acab64a398d69cc22dc684f7d6a1266f1b0a5c5de05139505f8a5c039916c)
- origin_servers.health_checks.health_check.icmp_health_check

<a id="canonical-f99e1d5cebcdff4626bb4f59cb416b7bb0d802597656cec614e0dbddc237539a"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
icmp_health_check = {}
```

<a id="canonical-45f1514153ff764fda21fd5173ac1967e44cd40df5d6764f0e254252093f1817"></a>

## Direct properties — origin_servers.health_checks.health_check.icmp_health_check / b7045d18f139 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cf49405d159550786174edc5ee749a4c09b10acce8b360c33e4f1d1d72f12dd4"></a>

## Next pages — origin_servers.health_checks.health_check.icmp_health_check / b7045d18f139 / 4

- [origin_servers.health_checks.health_check](resources--dns_proxy--reference--group-001.md#canonical-054acab64a398d69cc22dc684f7d6a1266f1b0a5c5de05139505f8a5c039916c)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-d36eef33355c945bfd324232c330465a23929a5ddf99aa527bfdb4cd5abdaf0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62714b51814403aa1e4488d251d0d686338b1ceb3c37262cbe888009f73e717e"></a>

## origin_servers.health_checks.health_check.tcp_health_check — origin_servers.health_checks.health_check.tcp_health_check / e945d0e3d0e9 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.health_checks](resources--dns_proxy--reference--group-001.md#canonical-2bd317751fa784cb9ff8a110775df20082b25dc4631e8946f6ce5f8a9c5d0bac)
- [origin_servers.health_checks.health_check](resources--dns_proxy--reference--group-001.md#canonical-054acab64a398d69cc22dc684f7d6a1266f1b0a5c5de05139505f8a5c039916c)
- origin_servers.health_checks.health_check.tcp_health_check

<a id="canonical-9471cd4aa9fdba0c8f9d208c96026f571081e5f36f511065a4576c371f2376fe"></a>

Type: `"object"`. single nested block, Optional.

Monitor reports healthy status if UDP connection is successful and response payload matches expected
response pattern.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expected_response",
    "send_payload")}
```

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

Terraform syntax:

```terraform
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-dff716a0743d1aca870d2b849bc65c811e2d84e0780190d9157b8ae0f38823c0"></a>

## Direct properties — origin_servers.health_checks.health_check.tcp_health_check / e945d0e3d0e9 / 3

<a id="canonical-8ca853070f3a244674883fd73119d25de38b700412f6be3e33717170082e1cce"></a>

<a id="canonical-73127340698b2395f0d0073c60d10396a7a02f402206161cf2c3aa2d060cca3e"></a>

## expected_response property — origin_servers.health_checks.health_check.tcp_health_check / e945d0e3d0e9 / 4

Type: `"string"`. Optional.

Specifies a regular expression pattern which will be matched against response payload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

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

<a id="canonical-13254c0ce743dbb703a7b6e98af20886f6d068fe0870bb8f8bef58e09e8f19cd"></a>

<a id="canonical-c0f53faa269cc96a62cb1437d9f4676e51cb5225566d74c8781671e146b978ae"></a>

## send_payload property — origin_servers.health_checks.health_check.tcp_health_check / e945d0e3d0e9 / 5

Type: `"string"`. Optional.

Send string. Text string sent in the request.

Upstream description:

Text string sent in the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

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

<a id="canonical-8140811e4853a3d1a422c52bb6b0b22fdb7d16660413d971edafd0af287df90e"></a>

## Next pages — origin_servers.health_checks.health_check.tcp_health_check / e945d0e3d0e9 / 6

- [origin_servers.health_checks.health_check](resources--dns_proxy--reference--group-001.md#canonical-054acab64a398d69cc22dc684f7d6a1266f1b0a5c5de05139505f8a5c039916c)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84ccff7e0e7ebf732a02413d8369019d86cd11c462e704996def9184222c2741"></a>

## origin_servers.origin_servers — origin_servers.origin_servers / 36a287da900d / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- origin_servers.origin_servers

<a id="canonical-fe35ff926481b8b1ddfa85c7c26045b4787619e26e58d275576f19477b25babb"></a>

Type: `"object"`. list nested block, Optional.

List Of Origin Servers. List of origin servers for Proxy.

Upstream description:

List of origin servers for Proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("k8s_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("no_preference",
    "site_preferences"),
  validators.ConflictingListObjectAttributes("public_ip",
    "public_name")}
```

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

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-050d610a7026c5f3f30cb4f42ef6329d2eb67a86d58c524fd9d468b5c44e156e"></a>

## Direct properties — origin_servers.origin_servers / 36a287da900d / 3

- [k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d): complete subsection reference.

- [no_preference](resources--dns_proxy--reference--group-001.md#canonical-a31577cf344a96f18535e9ca1114fadfceebb3163e4fc50d207178971aa0ae52): complete subsection reference.

- [public_ip](resources--dns_proxy--reference--group-001.md#canonical-3e8ea1e81ce5139586cc0d7ed6e534de47d071636db5ed375fa5fcff5ee984f5): complete subsection reference.

- [public_name](resources--dns_proxy--reference--group-001.md#canonical-56467c912bbca0c2dd3381fe73e9d8f80ec57dc81b2d627a388a97c950061139): complete subsection reference.

- [site_preferences](resources--dns_proxy--reference--group-001.md#canonical-b6945615fb9ceeb8bdbbb8816bfd90ed1d09a5d28213489e82737dbcd78270f5): complete subsection reference.

<a id="canonical-629c0aef8779f1b3e5018d2ff39a6152311e3210a0b6959d9b0095f2047a2507"></a>

## Next pages — origin_servers.origin_servers / 36a287da900d / 4

- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- [origin_servers.origin_servers.no_preference](resources--dns_proxy--reference--group-001.md#canonical-a31577cf344a96f18535e9ca1114fadfceebb3163e4fc50d207178971aa0ae52)
- [origin_servers.origin_servers.public_ip](resources--dns_proxy--reference--group-001.md#canonical-3e8ea1e81ce5139586cc0d7ed6e534de47d071636db5ed375fa5fcff5ee984f5)
- [origin_servers.origin_servers.public_name](resources--dns_proxy--reference--group-001.md#canonical-56467c912bbca0c2dd3381fe73e9d8f80ec57dc81b2d627a388a97c950061139)
- [origin_servers.origin_servers.site_preferences](resources--dns_proxy--reference--group-001.md#canonical-b6945615fb9ceeb8bdbbb8816bfd90ed1d09a5d28213489e82737dbcd78270f5)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06b01705bdf16704e2bc919cd4e97f45808383ca4de0e961ec29f72bb2ebed22"></a>

## origin_servers.origin_servers.k8s_service — origin_servers.origin_servers.k8s_service / e208c039aef3 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- origin_servers.origin_servers.k8s_service

<a id="canonical-b534468b088f0297d7b50a4a22b90ed89ca18d99f78fa960c00df3436583ab00"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with K8s service name and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "vk8s_networks"),
  validators.ConflictingObjectAttributes("outside_network",
    "vk8s_networks")}
```

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

Terraform syntax:

```terraform
k8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-8855e1fcf56312e655719c7511d5c7115834a919ba59e78229baf4aa5cd2bfca"></a>

## Direct properties — origin_servers.origin_servers.k8s_service / e208c039aef3 / 3

- [inside_network](resources--dns_proxy--reference--group-001.md#canonical-f2358752dca2714e49add5ea0e85db1051678bbf311443d8d8b27f97b90c9803): complete subsection reference.

- [outside_network](resources--dns_proxy--reference--group-001.md#canonical-e8aba1e4fd2910384424fdeab695ffbb26e809fcc9722da65ab526a464b5d21e): complete subsection reference.

<a id="canonical-23e28c5ea593fc2d908ac19456ebd8cb547400012267154b499603ec2605e434"></a>

<a id="canonical-219ca64b9734783b560c6177dc26f4334f5fcab42b882b1a9bc4eb327b951c70"></a>

## protocol property — origin_servers.origin_servers.k8s_service / e208c039aef3 / 4

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

&#8203;- PROTOCOL\_UDP: UDP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_UDP"),
}
```

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

<a id="canonical-7f9cee8959dbfc294d6bbd9a5ca6dde4311e147d3b86dc6c300df790ecb2946d"></a>

<a id="canonical-dcbc324c7cedaf079e7c5fd350afa352e81af8bb5fbcb3e9b16651f24471a374"></a>

## service_name property — origin_servers.origin_servers.k8s_service / e208c039aef3 / 5

Type: `"string"`. Optional.

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

- [site_locator](resources--dns_proxy--reference--group-001.md#canonical-ae1ba959965d67b8ed8847d0c499cc8a1e950611965fea7b2b4dfca57ecd933a): complete subsection reference.

- [snat_pool](resources--dns_proxy--reference--group-001.md#canonical-4c376bf2509d331e215246bc2821f2edebb58f5155cbb586b3883196c08a0f0a): complete subsection reference.

- [vk8s_networks](resources--dns_proxy--reference--group-001.md#canonical-6b7170715c92b81a0941536f66142ea59502546d1dd7c700366999a5c860e00a): complete subsection reference.

<a id="canonical-0a70888a2aecc53492c1fbf62702fe8098f5c1375a602be653748e556b6ec692"></a>

## Next pages — origin_servers.origin_servers.k8s_service / e208c039aef3 / 6

- [origin_servers.origin_servers.k8s_service.inside_network](resources--dns_proxy--reference--group-001.md#canonical-f2358752dca2714e49add5ea0e85db1051678bbf311443d8d8b27f97b90c9803)
- [origin_servers.origin_servers.k8s_service.outside_network](resources--dns_proxy--reference--group-001.md#canonical-e8aba1e4fd2910384424fdeab695ffbb26e809fcc9722da65ab526a464b5d21e)
- [origin_servers.origin_servers.k8s_service.site_locator](resources--dns_proxy--reference--group-001.md#canonical-ae1ba959965d67b8ed8847d0c499cc8a1e950611965fea7b2b4dfca57ecd933a)
- [origin_servers.origin_servers.k8s_service.snat_pool](resources--dns_proxy--reference--group-001.md#canonical-4c376bf2509d331e215246bc2821f2edebb58f5155cbb586b3883196c08a0f0a)
- [origin_servers.origin_servers.k8s_service.vk8s_networks](resources--dns_proxy--reference--group-001.md#canonical-6b7170715c92b81a0941536f66142ea59502546d1dd7c700366999a5c860e00a)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-f2358752dca2714e49add5ea0e85db1051678bbf311443d8d8b27f97b90c9803"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-459920b1c5ad5365690b087bd6b5f823a6c50ee578d9fe971661651363f696c1"></a>

## origin_servers.origin_servers.k8s_service.inside_network — origin_servers.origin_servers.k8s_service.inside_network / f5f0300343b3 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- origin_servers.origin_servers.k8s_service.inside_network

<a id="canonical-3d711559d5b387c2de3e08b3290fbb92d25f222e070a2ce157a79ac94deb4922"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
inside_network = {}
```

<a id="canonical-b9c5df1374bc4ee29a43d1ae173c610eedc1b56e9738cf0094a52b06b43b0a60"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.inside_network / f5f0300343b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e06034d258a13eca949a26a062361c2612397189c8df2f061673604b600b276c"></a>

## Next pages — origin_servers.origin_servers.k8s_service.inside_network / f5f0300343b3 / 4

- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-e8aba1e4fd2910384424fdeab695ffbb26e809fcc9722da65ab526a464b5d21e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5170e92fa20bd8ce45e3f0e5d801c5306fc62272f366e3bfc12339ada87a5d7b"></a>

## origin_servers.origin_servers.k8s_service.outside_network — origin_servers.origin_servers.k8s_service.outside_network / 6c957bbb820b / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- origin_servers.origin_servers.k8s_service.outside_network

<a id="canonical-37586fcd91b25124211e13ffe841ccd9670362a0396fcfbc7c8b7d88f97f6e7f"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
outside_network = {}
```

<a id="canonical-37c0b1d13ee209a77a1f986028d2c996d22eb12d4cd9e4550ac3e12b628ef526"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.outside_network / 6c957bbb820b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a957364b10b4873b14c7c6a834c4c2d6d7482ee1957a583c4095f739ee119c93"></a>

## Next pages — origin_servers.origin_servers.k8s_service.outside_network / 6c957bbb820b / 4

- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-ae1ba959965d67b8ed8847d0c499cc8a1e950611965fea7b2b4dfca57ecd933a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d05b61632a3f6416e4fa09a6276ce47e3807f4e0efb0ec9281d1ff27dc17e07f"></a>

## origin_servers.origin_servers.k8s_service.site_locator — origin_servers.origin_servers.k8s_service.site_locator / f09c8dbfbae7 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- origin_servers.origin_servers.k8s_service.site_locator

<a id="canonical-1d6ddc8b833d558976eaab57fce5adbdf44208c2c579ba839909bdae28215e42"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
```

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

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-ef5bd620c2c3bdb6e401bcdfcfcc244fea0698de143d7bbeaf4b3c7a0c95bbbc"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.site_locator / f09c8dbfbae7 / 3

- [site](resources--dns_proxy--reference--group-001.md#canonical-200832d948f7414d6ce0dce0376fe1246fa63ad396e50bf3e9c8ded0517fa830): complete subsection reference.

- [virtual_site](resources--dns_proxy--reference--group-001.md#canonical-426341f923b64597e2b75b25a0a7d8d4cb11a3511797cad5f2fd847ed1f7a1b0): complete subsection reference.

<a id="canonical-3e779a11d996f0b10fdbaf5470108a28616ef9c545f388de74388a4d5d934921"></a>

## Next pages — origin_servers.origin_servers.k8s_service.site_locator / f09c8dbfbae7 / 4

- [origin_servers.origin_servers.k8s_service.site_locator.site](resources--dns_proxy--reference--group-001.md#canonical-200832d948f7414d6ce0dce0376fe1246fa63ad396e50bf3e9c8ded0517fa830)
- [origin_servers.origin_servers.k8s_service.site_locator.virtual_site](resources--dns_proxy--reference--group-001.md#canonical-426341f923b64597e2b75b25a0a7d8d4cb11a3511797cad5f2fd847ed1f7a1b0)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-200832d948f7414d6ce0dce0376fe1246fa63ad396e50bf3e9c8ded0517fa830"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1930692ac083f10f3963e9f4d80d6b058868bbf553a38617d2eef09059f5961e"></a>

## origin_servers.origin_servers.k8s_service.site_locator.site — origin_servers.origin_servers.k8s_service.site_locator.site / 1de4bcee5856 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- [origin_servers.origin_servers.k8s_service.site_locator](resources--dns_proxy--reference--group-001.md#canonical-ae1ba959965d67b8ed8847d0c499cc8a1e950611965fea7b2b4dfca57ecd933a)
- origin_servers.origin_servers.k8s_service.site_locator.site

<a id="canonical-c62f96a7231f40af3cb0256dc57ad18449415973e444c550746b93e80c3aa5f8"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-98b83867dff05b037d18e1d7ec968bfa603ec391f646029d3724b0f3b3c9ed3e"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.site_locator.site / 1de4bcee5856 / 3

<a id="canonical-dff5687ca9ec03adfada1533beebf5da999d9deefdf98bac0e8f571f870949aa"></a>

<a id="canonical-2f5252e863d44165c4ddc450192e8be914d8f92b46989f6152a6a17bf12ddaf0"></a>

## name property — origin_servers.origin_servers.k8s_service.site_locator.site / 1de4bcee5856 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-a6578f48d06e94e03023f333b7dcf9e9697b0143291bdfc3be47abdf31fe7f2b"></a>

<a id="canonical-69a68fefb723acd7688232ece1206f0c90275239873ed4fcb21d289fe0a53621"></a>

## namespace property — origin_servers.origin_servers.k8s_service.site_locator.site / 1de4bcee5856 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-4f2633d414e97bf1ca2ec1ae020b2ebf656270b0c46f26b1ac6a335bd082345c"></a>

<a id="canonical-52d291229802bfe43e4f12e6a4118733775c892c856139e0a52af17b2a2dc4e1"></a>

## tenant property — origin_servers.origin_servers.k8s_service.site_locator.site / 1de4bcee5856 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-10a6f2353df8e3beb5e1bb92ff52f91abb64c219ee42326ced52c376ffb17f3f"></a>

## Next pages — origin_servers.origin_servers.k8s_service.site_locator.site / 1de4bcee5856 / 7

- [origin_servers.origin_servers.k8s_service.site_locator](resources--dns_proxy--reference--group-001.md#canonical-ae1ba959965d67b8ed8847d0c499cc8a1e950611965fea7b2b4dfca57ecd933a)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-426341f923b64597e2b75b25a0a7d8d4cb11a3511797cad5f2fd847ed1f7a1b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-731e8d6138fe38439a486fa05b92455c2cf706da1f44522aaee419f07d41eac7"></a>

## origin_servers.origin_servers.k8s_service.site_locator.virtual_site — origin_servers.origin_servers.k8s_service.site_locator.virtual_site / 1a38843a6054 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- [origin_servers.origin_servers.k8s_service.site_locator](resources--dns_proxy--reference--group-001.md#canonical-ae1ba959965d67b8ed8847d0c499cc8a1e950611965fea7b2b4dfca57ecd933a)
- origin_servers.origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-3a8d6dc21c5f799188143f7f32633da56200bad4f9235aed67b1914ae8051da3"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-fd07b2e302a6e95995b70c5d2375a86073e10b120e7dd46e1168ac2e5e7040ef"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.site_locator.virtual_site / 1a38843a6054 / 3

<a id="canonical-b5366ae9678a4e3894461427541a2e0ec9497c7ba9fa42705d14dcef586a4061"></a>

<a id="canonical-d2ff00375dab57b517816e26b8ec7afbea2227dc748e9b4ee3e41b07a0f2933c"></a>

## name property — origin_servers.origin_servers.k8s_service.site_locator.virtual_site / 1a38843a6054 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-0002eed3f7d3e5e3937a97f50ce8f704dd25a4d8d2bbb21726eaa0131ffd2037"></a>

<a id="canonical-92e81ebba160cce41235f067625ae443087b93a4ecd36a533260cadabd64855f"></a>

## namespace property — origin_servers.origin_servers.k8s_service.site_locator.virtual_site / 1a38843a6054 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-f1c2b533d1bf78880cff65c034e3ec9a17a7d1eb27848c8c8e3ce7f219335085"></a>

<a id="canonical-a398436b07a899fe25b55b911ee484775682d58412c03d87c2564d073779cd6c"></a>

## tenant property — origin_servers.origin_servers.k8s_service.site_locator.virtual_site / 1a38843a6054 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-41c5f68be3e4d683d6db547f9490305a82e366efa6d4289eccc6ea59d85f3a18"></a>

## Next pages — origin_servers.origin_servers.k8s_service.site_locator.virtual_site / 1a38843a6054 / 7

- [origin_servers.origin_servers.k8s_service.site_locator](resources--dns_proxy--reference--group-001.md#canonical-ae1ba959965d67b8ed8847d0c499cc8a1e950611965fea7b2b4dfca57ecd933a)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-4c376bf2509d331e215246bc2821f2edebb58f5155cbb586b3883196c08a0f0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c1fcd62d1e84404c7f850c65673c790ba2c27cf7726888a8461064d701c3158"></a>

## origin_servers.origin_servers.k8s_service.snat_pool — origin_servers.origin_servers.k8s_service.snat_pool / b968af6c6c35 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- origin_servers.origin_servers.k8s_service.snat_pool

<a id="canonical-c61121dd58b96b5c411c51c69669dfae0c4a98f639552f9e6960fab05c3103cf"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
```

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

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-43290568b784f7f8ce7a3733e65423c2b8e9f41bf9cf9fd2110273c1de8b014b"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.snat_pool / b968af6c6c35 / 3

- [no_snat_pool](resources--dns_proxy--reference--group-001.md#canonical-b64321b8d4e3343d6f413a8622935e778873577359fdb086c0d36c9a45f20dae): complete subsection reference.

- [snat_pool](resources--dns_proxy--reference--group-001.md#canonical-342f71d807641e92d5ba496952add2e12e658ec836fb10c39b72167c2cc04ff3): complete subsection reference.

<a id="canonical-2d781d4957ff32f8ce9be4e1befc64de49625acf9d3a39963d74b1a168b3d868"></a>

## Next pages — origin_servers.origin_servers.k8s_service.snat_pool / b968af6c6c35 / 4

- [origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](resources--dns_proxy--reference--group-001.md#canonical-b64321b8d4e3343d6f413a8622935e778873577359fdb086c0d36c9a45f20dae)
- [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](resources--dns_proxy--reference--group-001.md#canonical-342f71d807641e92d5ba496952add2e12e658ec836fb10c39b72167c2cc04ff3)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-b64321b8d4e3343d6f413a8622935e778873577359fdb086c0d36c9a45f20dae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b977dab8b8c49bdc4eff9043412c338a0e063e47383f265b1bdceeead99e20a6"></a>

## origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool — origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool / 0a8ad319edd3 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- [origin_servers.origin_servers.k8s_service.snat_pool](resources--dns_proxy--reference--group-001.md#canonical-4c376bf2509d331e215246bc2821f2edebb58f5155cbb586b3883196c08a0f0a)
- origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-4437abfcf13cb678e6e89c7ce3a1fe87a78ba92b76196fc44ded63611d9390ea"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_snat_pool = {}
```

<a id="canonical-09206c87487c267debd8ceb6f39390224ac19c29d1aaaab429f51ef59988b094"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool / 0a8ad319edd3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-61395934ec2fcf96f8ccb4ec8eac79e06589c2f7f76e8c1f415fdb5b742f63c2"></a>

## Next pages — origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool / 0a8ad319edd3 / 4

- [origin_servers.origin_servers.k8s_service.snat_pool](resources--dns_proxy--reference--group-001.md#canonical-4c376bf2509d331e215246bc2821f2edebb58f5155cbb586b3883196c08a0f0a)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-342f71d807641e92d5ba496952add2e12e658ec836fb10c39b72167c2cc04ff3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-523c9c8e992a5615a9e8d1fb72b3b09b6c086ca10caa6d97a4ed1c24bb99a969"></a>

## origin_servers.origin_servers.k8s_service.snat_pool.snat_pool — origin_servers.origin_servers.k8s_service.snat_pool.snat_pool / 6a43a890289b / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- [origin_servers.origin_servers.k8s_service.snat_pool](resources--dns_proxy--reference--group-001.md#canonical-4c376bf2509d331e215246bc2821f2edebb58f5155cbb586b3883196c08a0f0a)
- origin_servers.origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-149cf961e8bc7587bc14924c8e6eeaaec02e49b07e88b91f7ee154dc2365eba3"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-ed5efcca9fcdc9a3d985537ccec28a8f0c7ff0dc321f43f2319501bba14a65e6"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.snat_pool.snat_pool / 6a43a890289b / 3

<a id="canonical-ee8a97c63cb8ae672e5128c0b2f4c38cb8aaaa47bfcdbe9d5df756362aa0078d"></a>

<a id="canonical-617d502871bf634f4f8dae3901787fb79bc7acef566bfa5bd1902171a3648bcf"></a>

## prefixes property — origin_servers.origin_servers.k8s_service.snat_pool.snat_pool / 6a43a890289b / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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

<a id="canonical-c1af0e54b21fac817195de3bc116d888611efa0bd3acb2866e36fe60951ab47b"></a>

## Next pages — origin_servers.origin_servers.k8s_service.snat_pool.snat_pool / 6a43a890289b / 5

- [origin_servers.origin_servers.k8s_service.snat_pool](resources--dns_proxy--reference--group-001.md#canonical-4c376bf2509d331e215246bc2821f2edebb58f5155cbb586b3883196c08a0f0a)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-6b7170715c92b81a0941536f66142ea59502546d1dd7c700366999a5c860e00a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ac01c12245fe74399e6a0d693982245e994077b188d3075d8d85509854feaf5"></a>

## origin_servers.origin_servers.k8s_service.vk8s_networks — origin_servers.origin_servers.k8s_service.vk8s_networks / 4bbe543453d9 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- origin_servers.origin_servers.k8s_service.vk8s_networks

<a id="canonical-b5c7cd1bcefd7210438ebadf09fb21eedeeeed568b30df883e3ad4e07c6dca40"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
vk8s_networks = {}
```

<a id="canonical-3be929473002d8c757b054b029114a0b3658e05b24360ac1f453defd3e97d34d"></a>

## Direct properties — origin_servers.origin_servers.k8s_service.vk8s_networks / 4bbe543453d9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8772ff31e263edcf307073f6cf5f7ea83b0a02e81dacee8fae0501882aaf01d1"></a>

## Next pages — origin_servers.origin_servers.k8s_service.vk8s_networks / 4bbe543453d9 / 4

- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-99bacfb162f503267feb029864d80d4e510be393243985dec07df283da92563d)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-a31577cf344a96f18535e9ca1114fadfceebb3163e4fc50d207178971aa0ae52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d1edfe4c57a31bbe2b0fa4a024d81560ce66f013989df46d4247c3101123f34"></a>

## origin_servers.origin_servers.no_preference — origin_servers.origin_servers.no_preference / 45fab7091429 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- origin_servers.origin_servers.no_preference

<a id="canonical-32b1f156370cdc82e55cd9ed6a8fc6101731db170c24c61157e48291f0fa02ef"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_preference = {}
```

<a id="canonical-7aff29b45183e2c3bc814c3d1cbfe76a30e7e743b859e04d98bb1fe06f9912b2"></a>

## Direct properties — origin_servers.origin_servers.no_preference / 45fab7091429 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f6fe244e8de87e2d9a738abe02b4e68a9b8fdf35935905483142432adb13c1db"></a>

## Next pages — origin_servers.origin_servers.no_preference / 45fab7091429 / 4

- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-3e8ea1e81ce5139586cc0d7ed6e534de47d071636db5ed375fa5fcff5ee984f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4ddf7ef6555e804611eb5b0a579c2fed69b922e99b1788e3e98159f6df1e65c"></a>

## origin_servers.origin_servers.public_ip — origin_servers.origin_servers.public_ip / dacd677e9cc5 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- origin_servers.origin_servers.public_ip

<a id="canonical-4fc592cea90e889383c00c4dcbde5de2039c184bb0d56a8b6d1b6db01d841853"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-a2a5b657ac9ac8b587cff2441cd80c2947c8e1d2ffabc50a95fad514ffb551aa"></a>

## Direct properties — origin_servers.origin_servers.public_ip / dacd677e9cc5 / 3

<a id="canonical-7480a00485fc0eafd02376698b36a5be33068215bdbb9fc2431a770f665c3a7c"></a>

<a id="canonical-4b3c1df88d30d1d1ba49b04fecde44f00f5b957eeba70ab876dfd55702e0e69b"></a>

## ip property — origin_servers.origin_servers.public_ip / dacd677e9cc5 / 4

Type: `"string"`. Optional.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Upstream description:

Exclusive with \[\] Public IPv4 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-65676b601ef2df173bcc81f04f006a9840d9a7f603007c75dc99f36b4ed7aae1"></a>

## Next pages — origin_servers.origin_servers.public_ip / dacd677e9cc5 / 5

- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-56467c912bbca0c2dd3381fe73e9d8f80ec57dc81b2d627a388a97c950061139"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c6ac5f1e4a729e78a249577209e8c68cf901955fe031da078d78d916bcd4206"></a>

## origin_servers.origin_servers.public_name — origin_servers.origin_servers.public_name / d1420efddbe5 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- origin_servers.origin_servers.public_name

<a id="canonical-b4b17519ef09676e3472b3ea20ac1a7cb80af224120369a771829408b8787320"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
```

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

Terraform syntax:

```terraform
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-3fbc8d8f405f04a0189819226030885322902e926033dbb53752e2f9797360bd"></a>

## Direct properties — origin_servers.origin_servers.public_name / d1420efddbe5 / 3

<a id="canonical-569f7752e2dd0909201cbf057a647c0a6e76ad85aa1f4d500e54a296882caf72"></a>

<a id="canonical-e2940288ad5428f71022eaa81fa1415044c176934131fd92755817f9fbf9bc69"></a>

## dns_name property — origin_servers.origin_servers.public_name / d1420efddbe5 / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-4df8bd3cf5c85648c44f69d1248a0a1167f423968e4e744a6fb7298370a9c442"></a>

<a id="canonical-10aa628050bef19cce876a5a4a7b263291867dfe6aad772f802a09b3a904c23e"></a>

## refresh_interval property — origin_servers.origin_servers.public_name / d1420efddbe5 / 5

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
  ),
}
```

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

<a id="canonical-0b6ff8da97b5ee711317c7b4c88c5e336eec098405f6abab093e09693e38c7d1"></a>

## Next pages — origin_servers.origin_servers.public_name / d1420efddbe5 / 6

- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-b6945615fb9ceeb8bdbbb8816bfd90ed1d09a5d28213489e82737dbcd78270f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36475967fc5e792c1da09359f3e58edba475b33cadb1099a16a619bef9685848"></a>

## origin_servers.origin_servers.site_preferences — origin_servers.origin_servers.site_preferences / 78066426b81a / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- origin_servers.origin_servers.site_preferences

<a id="canonical-b493c645ea5aa76d43e5465506181837f6347d3d8e7180fc16e19d23dca7df45"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
site_preferences {
  # Configure direct properties listed below.
}
```

<a id="canonical-c02eaf870c2748861f2ab868d6f71c61047b10bc9e3c8b192e36997e8f41ac1a"></a>

## Direct properties — origin_servers.origin_servers.site_preferences / 78066426b81a / 3

- [refs](resources--dns_proxy--reference--group-001.md#canonical-98c313736f0950d2721e7b9d1e491fca691dbb6fb98bf0aec213f510e19b1f90): complete subsection reference.

<a id="canonical-cf03e88810cabdf8d18f92eaf6e1e79ed3da988cf883686de8090d1187613d07"></a>

## Next pages — origin_servers.origin_servers.site_preferences / 78066426b81a / 4

- [origin_servers.origin_servers.site_preferences.refs](resources--dns_proxy--reference--group-001.md#canonical-98c313736f0950d2721e7b9d1e491fca691dbb6fb98bf0aec213f510e19b1f90)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-98c313736f0950d2721e7b9d1e491fca691dbb6fb98bf0aec213f510e19b1f90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e82fee5f10d328b1d19619da69fab348863afc622ab2e71f239116a47d2bd0c1"></a>

## origin_servers.origin_servers.site_preferences.refs — origin_servers.origin_servers.site_preferences.refs / 2eaf6510a66f / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-943676c45d76ea9e183f5a02c3f7d8a7a5392c8f2403daf4ddcae415bce25f95)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-018d888f7fff261ea718d17310f96640010198ce0dcb899b2a966c98fc3fe8ea)
- [origin_servers.origin_servers.site_preferences](resources--dns_proxy--reference--group-001.md#canonical-b6945615fb9ceeb8bdbbb8816bfd90ed1d09a5d28213489e82737dbcd78270f5)
- origin_servers.origin_servers.site_preferences.refs

<a id="canonical-92753cb4c256cb6560631d4b35e8fc8c662e3cf4e4e3c6c7e584a4c697b6e646"></a>

Type: `"object"`. list nested block, Optional.

Site References. Reference to one or more sites.

Upstream description:

Reference to one or more sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-c8d6b3b9e51a9f941a7f189bac207f4faa2801ec413147019fd4e43b7fbbb15b"></a>

## Direct properties — origin_servers.origin_servers.site_preferences.refs / 2eaf6510a66f / 3

<a id="canonical-385144342b26d71c7afd0ae5d93c15887421c8a40cb524df1277e73395c2436e"></a>

<a id="canonical-d89830317903aa72c3a4367af55044b71c57ebd59f0d0ee94887d547e141c14b"></a>

## name property — origin_servers.origin_servers.site_preferences.refs / 2eaf6510a66f / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-75854368b1b6f165be4533b75cdc0727ec7a47025c5ddefe0e40dd2f2b015cc4"></a>

<a id="canonical-5cad9fbd53eddea66e89fd5a46d9e2be7a76f4b017a19da67a7f56af7b8d67c8"></a>

## namespace property — origin_servers.origin_servers.site_preferences.refs / 2eaf6510a66f / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-e3396942b41563842d86e6dfec5747829233029e91c1bd190b60858647237894"></a>

<a id="canonical-01d8ff0aa7408e0e9f93a9d9baa74e22e9d214b6bbec36aefe2104dc3a590578"></a>

## tenant property — origin_servers.origin_servers.site_preferences.refs / 2eaf6510a66f / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-85c15c0f5ba037ac5d339d6403ce615eba86f81e18998eef3ebb2eb43f081849"></a>

## Next pages — origin_servers.origin_servers.site_preferences.refs / 2eaf6510a66f / 7

- [origin_servers.origin_servers.site_preferences](resources--dns_proxy--reference--group-001.md#canonical-b6945615fb9ceeb8bdbbb8816bfd90ed1d09a5d28213489e82737dbcd78270f5)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-e8f7dfa31ff2b22f84858d14cd4bd06fb1fdc2295ca8b19b7e7d5d275e2b68b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f9279642601b4312372307b71e6d3e7d77f7b7c2e819faf41060b8048ca378d"></a>

## protocol_inspection — protocol_inspection / 4fa345e96093 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- protocol_inspection

<a id="canonical-b44c2972bf086ffb652866f8beb9e9f4899c9c22bdfb53a4627ecb1d77db5e6c"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
protocol_inspection {
  # Configure direct properties listed below.
}
```

<a id="canonical-cdc205e96c581846cbcb99af3e6295f5f8658d32ea10c68bad764305e9ad516e"></a>

## Direct properties — protocol_inspection / 4fa345e96093 / 3

<a id="canonical-40a6cd4fc1485a9b30bb7c2631ae9566aec84bedcdd9267df0441d30d67edea5"></a>

<a id="canonical-7c53da8ddf0975f9399ab5a4cf00f01524783bf2960857ed19dcd9531a511464"></a>

## name property — protocol_inspection / 4fa345e96093 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-b9d0811b390f56b63ae4c301dee7736fea2c1e616152b8e274f213feef9166cf"></a>

<a id="canonical-bc24be2f9afd7c0265bf76dd90e2c3f1314fe2885177fd9919fbcf311cf36cf2"></a>

## namespace property — protocol_inspection / 4fa345e96093 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3eb3fd47f1641177d12d4d0d5648b711e737f37920577fefd6bf18265dfdec85"></a>

<a id="canonical-aa31b8ce1fceb80a8f0f7bfb16bb7ec8113f4b8e7b218e2dc5abe3a85a81b5bb"></a>

## tenant property — protocol_inspection / 4fa345e96093 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-03803d4ca759a36ccb4db8257316e17167fedf12464748e0fd4b14876f15f143"></a>

## Next pages — protocol_inspection / 4fa345e96093 / 7

- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7777f5b673ee71471fad5954ba851d9399576892860be458515f32cca6c6ecef"></a>

## proxy_advertisement — proxy_advertisement / 9a304e4e2c78 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- proxy_advertisement

<a id="canonical-0986a35b91ad8935bea950dae848ae3b9a2ca711d3404b4ced97dd8335d04650"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for proxy advertisement.

Upstream description:

Proxy Advertisement Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_dualstack_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public_default_dualstack_vip"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public_default_ipv6_vip"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public_default_dualstack_vip"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public_default_ipv6_vip"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "advertise_on_public_default_dualstack_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "advertise_on_public_default_ipv6_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_dualstack_vip",
    "advertise_on_public_default_ipv6_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_dualstack_vip",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_dualstack_vip",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_dualstack_vip",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_ipv6_vip",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_ipv6_vip",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_ipv6_vip",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_vip",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_vip",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_v6_on_public",
    "do_not_advertise")}
```

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

Terraform syntax:

```terraform
proxy_advertisement {
  # Configure direct properties listed below.
}
```

<a id="canonical-67ecc139b78c7ad4eaa2eed7ceabb63279401c4b3b31b2fb01337532606eaeab"></a>

## Direct properties — proxy_advertisement / 9a304e4e2c78 / 3

- [advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91): complete subsection reference.

- [advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-36e413c3804a479487c72a00782daad55ec340e06316387ea5579e5f34a9cdb2): complete subsection reference.

- [advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-a8cfb9d394509b726e0d82c9c369b0980923800bee81478d2a1193a893c5ab94): complete subsection reference.

- [advertise_on_public_default_dualstack_vip](resources--dns_proxy--reference--group-002.md#canonical-d9018890345e9c38927c60cd2abc03a1c13887da81cdf0e2d11570f102d379da): complete subsection reference.

- [advertise_on_public_default_ipv6_vip](resources--dns_proxy--reference--group-002.md#canonical-a8dabdd0094f62327d00befc565720566f64203ef5501c846fa47158a9d172f0): complete subsection reference.

- [advertise_on_public_default_vip](resources--dns_proxy--reference--group-002.md#canonical-6a8eec7071e18f350fff1d62d082b941c2c8761c5e70f805cf71c559c9691d27): complete subsection reference.

- [advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-049a38713e827da66813f326a1263c600ce757ba01a9f6926640be9f17033e1b): complete subsection reference.

- [do_not_advertise](resources--dns_proxy--reference--group-002.md#canonical-3bb4e78ed32c4a8cda0a7dc8a0f096120494038738f8935d0887a0697ff759b7): complete subsection reference.

<a id="canonical-2ddf0050c6dd347456dbab8116c893f36cfc3e06600020f698e6cc478a832323"></a>

## Next pages — proxy_advertisement / 9a304e4e2c78 / 4

- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-36e413c3804a479487c72a00782daad55ec340e06316387ea5579e5f34a9cdb2)
- [proxy_advertisement.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-a8cfb9d394509b726e0d82c9c369b0980923800bee81478d2a1193a893c5ab94)
- [proxy_advertisement.advertise_on_public_default_dualstack_vip](resources--dns_proxy--reference--group-002.md#canonical-d9018890345e9c38927c60cd2abc03a1c13887da81cdf0e2d11570f102d379da)
- [proxy_advertisement.advertise_on_public_default_ipv6_vip](resources--dns_proxy--reference--group-002.md#canonical-a8dabdd0094f62327d00befc565720566f64203ef5501c846fa47158a9d172f0)
- [proxy_advertisement.advertise_on_public_default_vip](resources--dns_proxy--reference--group-002.md#canonical-6a8eec7071e18f350fff1d62d082b941c2c8761c5e70f805cf71c559c9691d27)
- [proxy_advertisement.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-049a38713e827da66813f326a1263c600ce757ba01a9f6926640be9f17033e1b)
- [proxy_advertisement.do_not_advertise](resources--dns_proxy--reference--group-002.md#canonical-3bb4e78ed32c4a8cda0a7dc8a0f096120494038738f8935d0887a0697ff759b7)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b280b42c8056c47f85176e9ba1fad7e811b313e5cdf19ed2ac1ad9520a2296f3"></a>

## proxy_advertisement.advertise_custom — proxy_advertisement.advertise_custom / 5f8109b0a47f / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- proxy_advertisement.advertise_custom

<a id="canonical-e7796a76f01bc96d1f6cd29f5cb41fab9e2deb07fa71ae3babd76ebfadd2fc00"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a VIP on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where")}
```

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

Terraform syntax:

```terraform
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-d8d92d65683d46f8d007d38ae0b463d134b3eb02828cf4e530df50811eb771bd"></a>

## Direct properties — proxy_advertisement.advertise_custom / 5f8109b0a47f / 3

- [advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8): complete subsection reference.

<a id="canonical-7fdab216df1461d1e1bdddcea260079ec9aead57a5a205bfb8b9641bdb5e86b8"></a>

## Next pages — proxy_advertisement.advertise_custom / 5f8109b0a47f / 4

- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03c01a712f46bde4ed47ba360d9aabb1205419a6eb2a71cea66e8a8e459e2501"></a>

## proxy_advertisement.advertise_custom.advertise_where — proxy_advertisement.advertise_custom.advertise_where / 6d5e952296d6 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- proxy_advertisement.advertise_custom.advertise_where

<a id="canonical-c3aebc34c594d6fb61f7023560a700b466485dca1fbcfe0b8ec3e36a61efba7f"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingListObjectAttributes("port",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("port_ranges",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site_with_vip",
    "vk8s_service")}
```

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

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-2edeae067968947c01e60c00c80dcac26d5fad9ad0bb1f0ee788ab101bb76357"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where / 6d5e952296d6 / 3

- [advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-64cad82be0b9729c28ad2caec0ed571d5167a65c50865c5ae234f216cc2787b7): complete subsection reference.

- [advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-bc85bf57f08c03a2d77a6d61eff33c6119ccc05ccc69e954c7a9df611d2108b3): complete subsection reference.

- [advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-58ffa668c93929171fa32ce6e8cad94a6d7acc87c7c420fc64ae00b2e72e587b): complete subsection reference.

<a id="canonical-5c7ac53e39a5351b6e0b401f2d57121cf77efbaa202cc22ce81defc1dedb4e96"></a>

<a id="canonical-8cccf3245f304af98a0784059ff58e644e4e0041ca83bc74ddc112b477fe16c7"></a>

## port property — proxy_advertisement.advertise_custom.advertise_where / 6d5e952296d6 / 4

Type: `"number"`. Optional.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Upstream description:

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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

<a id="canonical-cf0d6ce41469a52f5b36cf83251207d829b310cc353ee4465181d0867528f02b"></a>

<a id="canonical-e7ab703712db946dc73766c009f816fd6e202fdc56d2b663265c414cbb5e6743"></a>

## port_ranges property — proxy_advertisement.advertise_custom.advertise_where / 6d5e952296d6 / 5

Type: `"string"`. Optional.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

- [site](resources--dns_proxy--reference--group-002.md#canonical-3fa02dda6c2cf1ab054c037552de464ad85510fe3096f8358bc981e7e2306c8b): complete subsection reference.

- [use_default_port](resources--dns_proxy--reference--group-002.md#canonical-96c4edf8994da1e38c347e3702f02883bc62d3195ceb0a99c7865bf7821da37f): complete subsection reference.

- [virtual_network](resources--dns_proxy--reference--group-002.md#canonical-b448646677519bce2a1a9bfe3e0f5cb2c136e8074cae4328fccb54e8224e1efa): complete subsection reference.

- [virtual_site](resources--dns_proxy--reference--group-002.md#canonical-9913484b063cdc95a027af29b400730c1268977168fbc50afae97345195b30ae): complete subsection reference.

- [virtual_site_with_vip](resources--dns_proxy--reference--group-002.md#canonical-d893ad774dbda1e1c44ddb1f685e238f0ab412a959b195d2a3c5862a5c12c3a4): complete subsection reference.

- [vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-fd22c2efb9959a92cb3da99b917df0844bd4bf08a80b9f4522b22aa495a9223f): complete subsection reference.
