---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68ae45444d823b11646afd3a0b63a9b2d2b16fb7e7517ab5108ecd89daab106c"></a>

## Property reference — Property reference / 2150a054bb1b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- Property reference

<a id="canonical-3e8e7117a28f0970320eeb9833a54aa1cf33dc6e0ec57adb5ffee11d19cff171"></a>

## Direct properties — Property reference / 2150a054bb1b / 3

- [advanced_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-250c9f8b07036fc33dc6f787805f3dca2d5b508170e66d907330dacafe89c52b): complete subsection reference.

<a id="canonical-952b3fc2f672ff3ea236f9b922bf4481d975deda279504d225cf3c9bb4b8ea67"></a>

<a id="canonical-138921c427d7fda88bf57173a77cd515753a8363cd2521a4307ee93d22de2608"></a>

## annotations property — Property reference / 2150a054bb1b / 4

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

- [ddos_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-22b749b9380e691ba8bfd877b93a99f2cc8cabdce17fb4a0bda6eb12c4b504aa): complete subsection reference.

<a id="canonical-8ac6b08d0d89709ba6ee7a4a325dc1e3e9382dac8d092e4bb9f171563e9fa041"></a>

<a id="canonical-a4d8af5982dd17ef8feac14d75ac3d8bd93feed7d75baf23f72de15c04134205"></a>

## description property — Property reference / 2150a054bb1b / 5

Type: `"string"`. Computed.

Description of the BigIPHTTPProxy.

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

<a id="canonical-6b5fb2ed1f7f9e77452a19d4e8d29b02fefd12630be7bbd006515baed2640ae3"></a>

<a id="canonical-20fd15d2d6395b38fc8379d137447e7d6b8686f902d6e87570581dbd1dbb69a5"></a>

## id property — Property reference / 2150a054bb1b / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1a339565aa9a15a2645607264cca4d58d053f037f8bd0e3e846158f64d436fb3): complete subsection reference.

<a id="canonical-7f0b059a0687e30bb8bc20efeb66d3f434860c1a492b263096b03c46024beaa1"></a>

<a id="canonical-45fa0172504dfe58417909163e221f76c767d22f9b10c2db49756106170b89c8"></a>

## labels property — Property reference / 2150a054bb1b / 7

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

- [lb_algorithm](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0ed3b2daa8855618fec8a4a5e7a5ce44b295dceba3b83a658f4d469b9cb0d5c1): complete subsection reference.

<a id="canonical-ae91aab16586f2b80c497f03fc3103eb1f9e5efc5e6269dca3c6ca364df50c18"></a>

<a id="canonical-6c4338a38f7020451422684d9671098e0d7792df2ab1c17cbb35c60bf7950caf"></a>

## name property — Property reference / 2150a054bb1b / 8

Type: `"string"`. Required.

Name of the BigIPHTTPProxy.

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

<a id="canonical-ad0d3e5fd0fbe3379249860b732d93959cf6a4129e5982e1610e85ccf20b35e9"></a>

<a id="canonical-c6871803ab2aefd7382401d2239c8704eabdc49ebeb3ad5308436879188abedd"></a>

## namespace property — Property reference / 2150a054bb1b / 9

Type: `"string"`. Required.

Namespace where the BigIPHTTPProxy exists.

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

- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e): complete subsection reference.

- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-c26b7c8b16a4eda57dc835a575d989b32c52cf2b530dca551eec308ac24d31db): complete subsection reference.

- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a): complete subsection reference.

<a id="canonical-251ef561efdfda74a50af09d01919fd35a23fde506d4e7d17cab9d0f90d6f701"></a>

## All schema paths — Property reference / 2150a054bb1b / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_profile` | [advanced_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-6e2425007626682d10c3948dce4fd63e6ef4a94644c9681331a0bd5c369608f4) |
| `advanced_profile.disable_spec` | [advanced_profile.disable_spec](data-sources--bigip_http_proxy--reference--group-001.md#canonical-c970693d1e47d753d5109009b0bb3452eb973e14a1bb15c48527ff843dd233fd) |
| `advanced_profile.enable_default_profile` | [advanced_profile.enable_default_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f51e8b4fd59d7c5a80bcdcf205babc734ca7d8bc867ca44840864ac6ea437e60) |
| `annotations` | [annotations](data-sources--bigip_http_proxy--reference--group-001.md#canonical-952b3fc2f672ff3ea236f9b922bf4481d975deda279504d225cf3c9bb4b8ea67) |
| `ddos_profile` | [ddos_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f1332d9aaa34f7b5afa58b65ea9cdde426d7f01e04ff6dac5d157186c625fb07) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](data-sources--bigip_http_proxy--reference--group-001.md#canonical-38d597a3379d2f6c1dca4249f560ac1aa6f02a05c8f66b6d58768c35e7c2a66f) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](data-sources--bigip_http_proxy--reference--group-001.md#canonical-e7b7684df6299ca5c34a1c8043f9f8620cad60f39efb38409df2472b6aae105a) |
| `description` | [description](data-sources--bigip_http_proxy--reference--group-001.md#canonical-8ac6b08d0d89709ba6ee7a4a325dc1e3e9382dac8d092e4bb9f171563e9fa041) |
| `id` | [id](data-sources--bigip_http_proxy--reference--group-001.md#canonical-6b5fb2ed1f7f9e77452a19d4e8d29b02fefd12630be7bbd006515baed2640ae3) |
| `irules` | [irules](data-sources--bigip_http_proxy--reference--group-001.md#canonical-eca18e93ce30aebe953dc739dfe58fc83b49d595ac19385f06bbf300ae44061b) |
| `irules.irules` | [irules.irules](data-sources--bigip_http_proxy--reference--group-001.md#canonical-09d34740298a272e9a8966fa8ca7225e4aac821a8111af2950e189f01d7ae826) |
| `irules.irules.name` | [irules.irules.name](data-sources--bigip_http_proxy--reference--group-001.md#canonical-94988e7bffb7261e2f7067375a3b79917d4e3fea60f7e06208981c41419d31dd) |
| `irules.irules.namespace` | [irules.irules.namespace](data-sources--bigip_http_proxy--reference--group-001.md#canonical-ffcb7693a5f3f703bf4272b2d13c5c9f221fe182c3052c73a87ec76b40d5da42) |
| `irules.irules.tenant` | [irules.irules.tenant](data-sources--bigip_http_proxy--reference--group-001.md#canonical-a108d8b8d4b13d3b1cd68907b4f33b2e39fa903cc79ddf4d7da57119fb0077e2) |
| `labels` | [labels](data-sources--bigip_http_proxy--reference--group-001.md#canonical-7f0b059a0687e30bb8bc20efeb66d3f434860c1a492b263096b03c46024beaa1) |
| `lb_algorithm` | [lb_algorithm](data-sources--bigip_http_proxy--reference--group-001.md#canonical-4db29f6077710634cbc03be6d95e80e1eaf89b8c9f831425327b42e29db1eef6) |
| `lb_algorithm.round_robin` | [lb_algorithm.round_robin](data-sources--bigip_http_proxy--reference--group-001.md#canonical-a257d7c48971dad9a095eb18c00726fcf705da9c9acbd40d24fa302d64b9416d) |
| `name` | [name](data-sources--bigip_http_proxy--reference--group-001.md#canonical-ae91aab16586f2b80c497f03fc3103eb1f9e5efc5e6269dca3c6ca364df50c18) |
| `namespace` | [namespace](data-sources--bigip_http_proxy--reference--group-001.md#canonical-ad0d3e5fd0fbe3379249860b732d93959cf6a4129e5982e1610e85ccf20b35e9) |
| `origin_pools` | [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-d98b635e4877934e47542e8fde8ff65642ea9933538a0335a9d06cd44e7f9741) |
| `origin_pools.pools` | [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-60ef50d97aa68385c529dc2d2851cd4f6749b6c7e8d9bc8f84af0d7067c4968f) |
| `origin_pools.pools.name` | [origin_pools.pools.name](data-sources--bigip_http_proxy--reference--group-001.md#canonical-109de3f27cc8145458fd5f42c6bfd2245337f935fffaca622288fc83224de8d9) |
| `origin_pools.pools.origin_servers` | [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-4cc0fcc14bbd64516a356f69bf54a519d0324b1a87fbe57cd86e447d18fcd8d7) |
| `origin_pools.pools.origin_servers.automatic_port` | [origin_pools.pools.origin_servers.automatic_port](data-sources--bigip_http_proxy--reference--group-001.md#canonical-b95648a18b7ba9534f767436d523e69278405250bb0b8e2387446c15fc217c24) |
| `origin_pools.pools.origin_servers.health_checks` | [origin_pools.pools.origin_servers.health_checks](data-sources--bigip_http_proxy--reference--group-001.md#canonical-7d94e81d40f83db9b0dcb0a759d1c1ec2010da77acf3272709f22ae6221a11dc) |
| `origin_pools.pools.origin_servers.health_checks.health_check` | [origin_pools.pools.origin_servers.health_checks.health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f8e8cb7f729522e28f866dcaa5c1c36f1bd5097c9d3d520ff5c9e4f78225424f) |
| `origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check` | [origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-4608ba84d2ef94274309d51c31645c93af3b1b1b425b5ef30c929df5d680ce4c) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-213aeea9391e73c994a272adb4b2f12e16956cae5771e20e41355a072c7572ff) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.expected_response` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.expected_response](data-sources--bigip_http_proxy--reference--group-001.md#canonical-ccdb4f197002de3071e084dde3bf55b4f6633180f84efd9479d2ba64bb7b1784) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.send_payload` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.send_payload](data-sources--bigip_http_proxy--reference--group-001.md#canonical-024fa6ed46b3a523ecb7ba0083e9429ffd4650cff3d3f90a1b627dc3f140a1bc) |
| `origin_pools.pools.origin_servers.health_checks.healthy_threshold` | [origin_pools.pools.origin_servers.health_checks.healthy_threshold](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1536c70a3e0aa005c6cb509df121fd3ef7912933321c006aa172273c26031577) |
| `origin_pools.pools.origin_servers.health_checks.interval` | [origin_pools.pools.origin_servers.health_checks.interval](data-sources--bigip_http_proxy--reference--group-001.md#canonical-ba46146150729436f985f84b3ba8c6446c7d5dcb79ad0354374aed8fa4a5d531) |
| `origin_pools.pools.origin_servers.health_checks.timeout` | [origin_pools.pools.origin_servers.health_checks.timeout](data-sources--bigip_http_proxy--reference--group-001.md#canonical-898be77d5767c4d3dfdee9c97d438385688ee7ff8cc8ed97c4565f656fedb4a0) |
| `origin_pools.pools.origin_servers.health_checks.unhealthy_threshold` | [origin_pools.pools.origin_servers.health_checks.unhealthy_threshold](data-sources--bigip_http_proxy--reference--group-001.md#canonical-7631fdccbef4535c4e9b501ff9cf1a0ea740316259f1cc4936be5a5172995415) |
| `origin_pools.pools.origin_servers.lb_port` | [origin_pools.pools.origin_servers.lb_port](data-sources--bigip_http_proxy--reference--group-001.md#canonical-abcf26e9312bc007e2ba92c8f851799879404388a665c0ea2a36bafbcf5f3c3b) |
| `origin_pools.pools.origin_servers.origin_servers` | [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-fb5279501fba94f37f6a64b2522665e6ac1878440b62c6d8c8f7bb9dd024593c) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service` | [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0d13de34f9eed505a093fb50be90c5b504ab6bd31299da3739cdf3f4882db888) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network](data-sources--bigip_http_proxy--reference--group-001.md#canonical-7009de49a95e5495824730ce2d490d421e414c3024dc4ad43effb3277ad29c1a) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network](data-sources--bigip_http_proxy--reference--group-001.md#canonical-29b6a0f5e0ccc800cd3d787a35ac0c4ffcd4779ba1827f28abeadd02b8867cb6) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.protocol` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.protocol](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f3e096629bcee808a04821b2673eeb2eee894a7c6fd81ba49d707bf0993d34e5) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.service_name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.service_name](data-sources--bigip_http_proxy--reference--group-001.md#canonical-e463cf99ce14bb2229f504df45d8a68929ab085a1fd2f81dfa99ad6f2ec08d21) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](data-sources--bigip_http_proxy--reference--group-002.md#canonical-d9860e80b662bbb4d94b218977b3c43c86051d717c702d965e9e1069143303a8) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-a2eb1baf1203a23ba232b1d5428a3ce44a40c15b281322ab7c42dd407c22a48a) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-fd3db9576f5c11c8547bd22e2bba0b68271905b2a53844c0dc0597d76adb8e50) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.namespace` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-f13fef8c8fb5d1b0c51a31a0a3c337f603eee1784e6c40366c8d4746d5e80c82) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.tenant` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-4938302faf96ab06d32872ebd41540f3a6a740cf23aa3434fcaa68cd3f46fe85) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-fce70c752313880f0ad96f7bc52592ad935e8caa3dc9cf71a3889ad8175811c7) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-4099b547ff0a71a5a1f746860bfaf2eb96113fc4ac2dfcd433afb70a45b3e5d1) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-34b04b1d8416d6f24948750edb84104dac80fa93a01073970d1a082eedb0f71d) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-93956509a06b1027c990e1319f4057441687245c3a32aa142b5e12d875faf22a) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-42cdff5013e94bd5eec0ecaf4db5afc5b865b17582d254b384a97f6399072798) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2251d3bb544feef42c74641dbd9a74e798c7614d01814246b65e43c9d607112a) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-c886e712c98a4389d833311077669afafd4bd933da86c7e638df8be983ae3ed8) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0d4cfaae30f10976a6082aa022096950cb95c273aafd10746cd20aa177390c36) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks](data-sources--bigip_http_proxy--reference--group-002.md#canonical-bddbe45783ed1e8003d7375bf236b9034e7b73ab925c836dcb94b30da4216011) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip` | [origin_pools.pools.origin_servers.origin_servers.private_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-bdd167f76f834c4f5956bfc715e816382c46cc114c9b815afe7c54ce8dc48024) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network` | [origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-04ea30c56601b4f0beb0e0fcfc6a90173ff87ee4439a2559a16f92457329e497) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.ip` | [origin_pools.pools.origin_servers.origin_servers.private_ip.ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-8bccef442135c4c0b043c1382f532e43c02f8bf4d9d19347d3b41c58e6fe34f3) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network` | [origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1988ebe6747f3d427472fbbf034b86f05de3ee2b3ad9aaeeed2e289f166da4e7) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment](data-sources--bigip_http_proxy--reference--group-002.md#canonical-57de61f98a4ca7be0ef3f72cf5797480eaa7988abcd673cc574daa784392094c) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-e83887871e36c31eab986142dad7705b4474b398357abe7915d810f0dcee1230) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-f33da28cdf7e77bbb63c740f3010a6a2d703ab8be9e0e650294df96282de9248) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-12b456f606337321a6bbdf16577a36672be62535e316dc39201100092cf7b41a) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](data-sources--bigip_http_proxy--reference--group-002.md#canonical-f20b7d54ec3abf1976d2aacec54794070a88797090e9b6262d2964cdb0a67ef6) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-006f01e46563f5e31e1363e51bd090cccf21a109d378a64cc1b3191b22b74d76) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-37203e150c15f4b41c8303bfe6df1bee07edc3a7ab197eca5a19b86b65ae0759) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-539b56a2d0470b5e676d57299c806c4aca0e3f7638b6469a1f4d02572d0a4006) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-78ce4b6e04d7241e21f4c9d758f75b2b7a341df77858726109dfa3706b98de38) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-96a59f0a07c2a0a05aa7d210402965f67ca5c6439591c87b950616f27961c4d8) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-7fc147e15174184cb4b38a9bc0d80dd9398ced53ffa78d66f7343fc5af51116c) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-d0602e0ba1407b771fd859ae9bf37e496e852d919d4b8a44dc63f0f14d21e9bf) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1903415e0d9e317c7dab92759066a726440040e0ea3f55baa0e6450eeb14623f) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-8a319892c4cc8a73cdfbb7d6b508f3c94d731f07b8080ad5056f31a030103c02) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-a79b45c2ebcd9d5c0413eb08c7faef9bc49b90d0bfc41ddf14cf946a8d621b54) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2c7e4bb590c2ef121a0b39d6f239b96926fa2a7cd72b1f9adf632becfb228b35) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool.prefixes` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool.prefixes](data-sources--bigip_http_proxy--reference--group-002.md#canonical-b91a195fcdf5a01752ab276ad07e49426f9e6ebb0ef77956590184ced9fc128f) |
| `origin_pools.pools.origin_servers.origin_servers.public_ip` | [origin_pools.pools.origin_servers.origin_servers.public_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9714ed51d77d161af031d81c7ec74e0fceb395d47770119e51f46745c10202bf) |
| `origin_pools.pools.origin_servers.origin_servers.public_ip.ip` | [origin_pools.pools.origin_servers.origin_servers.public_ip.ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-cc4f737fe0390d124b611ab7fe64677bbec40f4f1325d3a3f905188c04f1dfc0) |
| `origin_pools.pools.origin_servers.origin_servers.public_name` | [origin_pools.pools.origin_servers.origin_servers.public_name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1c185e1249c99cdec0f5b9576fa81911d453d676e1f0f44f49beff29b684de10) |
| `origin_pools.pools.origin_servers.origin_servers.public_name.dns_name` | [origin_pools.pools.origin_servers.origin_servers.public_name.dns_name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-c428f2e826830cadbf3c18f35174a8e71856123e5315932ddc7349997d744e96) |
| `origin_pools.pools.origin_servers.origin_servers.public_name.refresh_interval` | [origin_pools.pools.origin_servers.origin_servers.public_name.refresh_interval](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9fa84824391a8d49a81d6af005a10abf0c3f61c4bba779a3224125d7eab992c1) |
| `origin_pools.pools.origin_servers.port` | [origin_pools.pools.origin_servers.port](data-sources--bigip_http_proxy--reference--group-001.md#canonical-668c76ceac2344a8e13a3c59880be784b4736d56bce7a346fa7154e31868c691) |
| `origin_pools.pools.priority` | [origin_pools.pools.priority](data-sources--bigip_http_proxy--reference--group-001.md#canonical-9b3b6f538036ee7a761b8517b1f37eaee47e2e22b143884121699ade57bc0bb0) |
| `origin_pools.pools.weight` | [origin_pools.pools.weight](data-sources--bigip_http_proxy--reference--group-001.md#canonical-51a602ae27991150ae4aa6c4ab7a5445e215df744e206bd0779a1f1154f59ed7) |
| `proxy_advertisement` | [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-08945195182be00dd22d607978804b464127b12113a246a7d5a78ba8f673c3ea) |
| `proxy_advertisement.advertise_custom` | [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-ec555b5e2aff52564163adbcfa51b9da4777204da4d66e34117ba130214ed4be) |
| `proxy_advertisement.advertise_custom.advertise_where` | [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-c25e4624e5c81092b0c589693409b9e1b30a3677944f12bd3bde5ca91644273d) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--bigip_http_proxy--reference--group-002.md#canonical-aae6dd8bc77b5d6e1b938608876cbcd39fb9f8378621454a4da9cde8caf6d3ea) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-cf99d95a71884489acdb63365bf25c4485b44a1fe21b286a90f094c304e23a93) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-a3df3eff7c65dbaefa8181ba2e87c601b90d143b9351aecd89df8cdbc82ef72b) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-35104d11214e1f3b85f645ef247f0be38a64a05f7151a2d80222f5a3e7b4047b) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-7ce5ea3c6f64c6abb70ff288f2ee67b8359b13875344d528c9757ef25d1f3f72) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](data-sources--bigip_http_proxy--reference--group-002.md#canonical-49d55ef00718428dbef92c7b629a77ba84d05ad3137e9461782d3a26e0e32e70) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-b5f85093b440bb01c9be21fb07a66cc31f75dfb06315560eed0f145b801fc7d4) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-183481e75731f53f96d085f4984dd2f2721f906bfc313c9366be205cbbc9ff98) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-4970bbc6647beb54c29e687619905079dd2db803a4c1ef7ebdb74bd434901273) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-b49454b2ab2d9ceba39decdce24a1092dd95f5d7ac7866bf3ed12f24910eaaf4) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](data-sources--bigip_http_proxy--reference--group-002.md#canonical-b8e465263ed8d68cc7c2f57d5071883a2cbd1bfdaa110dec53beba878d4977e7) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-a9c2ccd8d66db35d8534d2a190b1a18715ea5b7b59f2c51153c066b27c1f5bc4) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-7b488a45d40b28cb94e2a9a0ef005d0bc1327ba61972756d67d74f2b333714ee) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-21d54058eb4529f59fd40c3df48d38086ec00c9a4c95744189918043c51123a4) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-90198df71ae7431e1d0cad67b36cf60aaa564c9543b2304f0973a93d893088c6) |
| `proxy_advertisement.advertise_custom.advertise_where.port` | [proxy_advertisement.advertise_custom.advertise_where.port](data-sources--bigip_http_proxy--reference--group-002.md#canonical-f9e3cedf0916b59402df0d848f637e8fc82f05085a7683a7defe1878562a7782) |
| `proxy_advertisement.advertise_custom.advertise_where.port_ranges` | [proxy_advertisement.advertise_custom.advertise_where.port_ranges](data-sources--bigip_http_proxy--reference--group-002.md#canonical-95344f2e38e02a2aefc0a4e4ac1a3e4d37a797a3780524904c483400dda56690) |
| `proxy_advertisement.advertise_custom.advertise_where.site` | [proxy_advertisement.advertise_custom.advertise_where.site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-4e7a6c984fe48e831bff2a0c26bfb8658c553b1e58e4d8281a37822e6f25fdd7) |
| `proxy_advertisement.advertise_custom.advertise_where.site.ip` | [proxy_advertisement.advertise_custom.advertise_where.site.ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-75f5016039123fcf396af6a2f00df3b16ee8a80afc8bbb4cd403d37cbfe1b705) |
| `proxy_advertisement.advertise_custom.advertise_where.site.network` | [proxy_advertisement.advertise_custom.advertise_where.site.network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-99007de23a4347b10342af1eb4bac3c29063703f9ede2e2748194333df352705) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site` | [proxy_advertisement.advertise_custom.advertise_where.site.site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3976120cee02ab8cb96713e79c062b4bbaebd764060f0d42a5f11d3f07fcff43) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.name` | [proxy_advertisement.advertise_custom.advertise_where.site.site.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-f8e1b81a2cdf1101390923f25a716c1bf66b797fa4dd2666e3c0c03ebc349e21) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.site.site.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-c009edc295c890fa599ec23a784848b31522ff79203211c449adcb9bc94a0874) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.site.site.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-5911e2623908ae2eaa2a31afa680e25b6db51886cbb771ee9ce6eeebfb74deb0) |
| `proxy_advertisement.advertise_custom.advertise_where.use_default_port` | [proxy_advertisement.advertise_custom.advertise_where.use_default_port](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0e356a12558880208322d2cd40945f1dabdcbdc8c6a30b782ec023a1adcdd1d1) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-5eac8e94f9546679dc2b1d5b4ca4273c25e436ff34bed9d0b0ec03e04d0c9f78) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-6163aaf0ef871e704b568baf3419247e8e7e0e10a647bc6dfbf4f62f099cdeab) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-bf5342b1507df6119692576d7f6c7c0c5492ea8535c1f4daa5780d158a77a662) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-03d7578990c3dde8b1238a3113d47b5b743a6f2974ee6e3a50cea28c0489f3a6) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-95cc4102dbf3c95e548a0e3cf87f24aced7c77a8ab152a72bafa7768dcdfaa2c) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2db4394227196e294ca22de329c1ff494a620206ce3ef5bde93fa43a5098b91a) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1726079f777d8a4554900d492e8279a43723e92080cea70bd8c10febbec62ae9) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-c5498b1acc417a28db16f6f04b252db9648752c41fc60aa178f11a1fe1f50137) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-66619c3e3e6788ba2328477477bcdffc494a1c193d4aeabde70cdaa0d759ae1b) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3bac8715b5577fa18e38699d12c8ddc86cf971fd4f0674e73b3c052af49a82b1) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-79436de82ef2662c77f1a2bcea58b35ba1999c3d988a27a48153a979c749f508) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-239a7da067c73536e36963c775161f8d3c294d09a7c951c8dab421c8e4c276bf) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9f9dd724148b0c2ea79c863ec1f1c040522d599bdcd830d2f82977416d6aab2c) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-e77923c08ef34fb382a15290eab4e1e9d2aa3d2891ce4c608de2621b6a746401) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-a1d3eaf77d36194097d11a08078c14f261e724440df24eaed850c912d8333c95) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-5a7bc66e5eddc353d1e06f98298c29108b02e350d3c4359138ab32fc4afabbb4) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-33ab384f3b82b91e46e3b2402ba9dd6b76ce31d211b1ab84a25c99b67cea0395) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9cafa57fd5e9674a83f04ff44328a8aeadfbbd86494ae26617ec182145aae1fa) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-50cf4c297764bc1da130efed3402c30d756e314d7ff852d6171e9034f50795fc) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-5ecba489a578ca36168471099f1da0a931f7b5d400e34cabc9a1b478a13384d4) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-ec5c17df07fe3d7697769ba6fa59ac1e8ce85c92cbfc5a0ea4aa1b7a3950ddc0) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1416276e1ba98c5a8a89b2f7b48c9cfd5bc17e9e5b64e56bbd22cee74640f433) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--bigip_http_proxy--reference--group-002.md#canonical-f248dec74d2f14128c2295d97ee84663bff3c077ab4140fef95b8c2a6a191d4e) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0682f5bdeb3651bbbd4aba2b7a203c44864f2c3e54d0bbf47d9b57d833c2de22) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-6b8713056e0b909bb7032e0d6030ed394875b465fa2606a8accf7f46890959fa) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-fdd509d835afe2a70b29a445d595f3213e14426563b8008c70daf456ba6c3fa3) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-4ae260ce89762a9dc663dfc40637c5a7cbfecfa6e1ce12e8cafd8f2e212275cf) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-b0b7c6e6d3133c112743d38cf585b685d477120977d581d8976cf3b400bf06ac) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-dd7bd61aebcf16f191b8071f149fd0664b5c3c88bbe03001576b4e2bd7dfd7a6) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-b8ca77fc83bf8fc4d8eba285d0ca89860e9e98ce39b5b36c83f6d65974f75d08) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-d3ea9084d6607e321669995f33fb8ceca827781ab131754edf22e615a6865a0c) |
| `proxy_advertisement.do_not_advertise` | [proxy_advertisement.do_not_advertise](data-sources--bigip_http_proxy--reference--group-002.md#canonical-cddbbb48178379e40f728b5b6076e36f93cf57ca3664568ef9d71ff70a3eb8a9) |
| `proxy_config` | [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-6d93f49fb78a7e8ab4daa3dc9928efa11a611520181ac3695fae7354de1a94ba) |
| `proxy_config.domains` | [proxy_config.domains](data-sources--bigip_http_proxy--reference--group-002.md#canonical-ec9e4dc888ccf2dabb45c62077d8a357565d04b00a782c36f5eaedd8dfcf3e00) |
| `proxy_config.http` | [proxy_config.http](data-sources--bigip_http_proxy--reference--group-003.md#canonical-21eb2dfc95fb160fb5553439ab04f01ab33308041b081c0c6d61bb97fd7dddc1) |
| `proxy_config.http.dns_volterra_managed` | [proxy_config.http.dns_volterra_managed](data-sources--bigip_http_proxy--reference--group-003.md#canonical-648743fe43d44e0cdf5142b4029cc7a4220dba065d3c52e07f89f34c0a78a8b8) |
| `proxy_config.http.port` | [proxy_config.http.port](data-sources--bigip_http_proxy--reference--group-003.md#canonical-4b039b90cf6df9350fc9d1b2fcd716e89f61ed0c2f3f7794a95dae36942813be) |
| `proxy_config.http.port_ranges` | [proxy_config.http.port_ranges](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ee0cf97a35d039d283a1fc50d8f7cdb7443ee6dd25096a243501e312aa768962) |
| `proxy_config.https` | [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-dd901f55626c9e948b2f573c2c432c7ee4ba082fed149ee32bc7ad169d816664) |
| `proxy_config.https.add_hsts` | [proxy_config.https.add_hsts](data-sources--bigip_http_proxy--reference--group-003.md#canonical-4ff823f5346ae746c06df05a5f3dc9d846889c3457913b5ed8532a5f6c3e520e) |
| `proxy_config.https.append_server_name` | [proxy_config.https.append_server_name](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2b8ed58e9211d74604aa818eb866052d21b76624f7da499f93bc7668ba9fd394) |
| `proxy_config.https.coalescing_options` | [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2bb45ed6c0b9beea3ebf4634c1330f2dff816e6abf44ef6bef18f0174751843a) |
| `proxy_config.https.coalescing_options.default_coalescing` | [proxy_config.https.coalescing_options.default_coalescing](data-sources--bigip_http_proxy--reference--group-003.md#canonical-de71922070fc1906b598ae6ffae880cd7c2df79126ddc940dd58a25694ddc202) |
| `proxy_config.https.coalescing_options.strict_coalescing` | [proxy_config.https.coalescing_options.strict_coalescing](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d2ff70092c3e5b95d821b2fcd78fb7cb75c711d0e3cb5b3230a78d633a033cd3) |
| `proxy_config.https.connection_idle_timeout` | [proxy_config.https.connection_idle_timeout](data-sources--bigip_http_proxy--reference--group-003.md#canonical-aff1d4b427548484d39a43d47b072bfeccc792915077d00bb1496d3e98d12812) |
| `proxy_config.https.default_header` | [proxy_config.https.default_header](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f528313c3821e2f19438a3951bc0e57cd13b323e3c51e0bfdcd8ccbaf3a8ff0c) |
| `proxy_config.https.default_loadbalancer` | [proxy_config.https.default_loadbalancer](data-sources--bigip_http_proxy--reference--group-003.md#canonical-b3b45d0999990021ad93fce611771658f08b3c728ca86748c38b13a415083477) |
| `proxy_config.https.disable_path_normalize` | [proxy_config.https.disable_path_normalize](data-sources--bigip_http_proxy--reference--group-003.md#canonical-553211d3ccfa833e73bb06216caf83be870b78874807f9e4ff110b64816f0fd5) |
| `proxy_config.https.enable_path_normalize` | [proxy_config.https.enable_path_normalize](data-sources--bigip_http_proxy--reference--group-003.md#canonical-adb7715efa237152995753e3637fb6abb6319af6f01852d3506a9569fd6547af) |
| `proxy_config.https.http_protocol_options` | [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-88b8f8c01afc110747bc76859b953d5c3dcca88db9208418ce8cda417edde962) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-36f0474f9725c1fff5ab944bf147e3462b63ce4207ab974a43cf07f854312ae8) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-743bfceeb01bbbd69e2135921a9d8201fd5a7172135258b71021ba464d3593f8) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-c28a8c7d4bb1aaf35bf4b122b5f9e6e6318b0645b6e3b923b66b4177c05ba008) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9ed3e912acfd1c8a71e5744533de1b75ae134373b19afeb0d69731c6cdfe527f) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0abd2ef9a2719279ca5e96069f430c99f6c10ec8bdc458d6a45d3431fda352a6) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--reference--group-003.md#canonical-a7746050171947b18e44db6f304addc90fda2e120e51c7a838fc1425b6a79e37) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v2_only` | [proxy_config.https.http_protocol_options.http_protocol_enable_v2_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ae7b3dd620008e85b67443ac5fa10cbbb78ed79d50d95ab7881e9a60e6bea7eb) |
| `proxy_config.https.http_redirect` | [proxy_config.https.http_redirect](data-sources--bigip_http_proxy--reference--group-003.md#canonical-20908927e7d18777e9e355949006236eb52873ec25ec29a29d9b51a375efde8b) |
| `proxy_config.https.non_default_loadbalancer` | [proxy_config.https.non_default_loadbalancer](data-sources--bigip_http_proxy--reference--group-003.md#canonical-24aef35840aaafdeae16289b7b8b67a4e8a64e5499bea8d0422145f7f8afb136) |
| `proxy_config.https.pass_through` | [proxy_config.https.pass_through](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3e2a806e783dd1d3d4ff68e9e156cf9348d83fa690c080b6be0b9614f61067b9) |
| `proxy_config.https.port` | [proxy_config.https.port](data-sources--bigip_http_proxy--reference--group-003.md#canonical-51a1da62662fb9e82c7cf08c7c443c3fb43963456752979757317e4712116b08) |
| `proxy_config.https.port_ranges` | [proxy_config.https.port_ranges](data-sources--bigip_http_proxy--reference--group-003.md#canonical-4a4551e56df180e7a67eb1032cd755a159ee0bc3b2155b13191238c931691a75) |
| `proxy_config.https.server_name` | [proxy_config.https.server_name](data-sources--bigip_http_proxy--reference--group-003.md#canonical-7ea2faea4d0ed8df2108c81b38269d95fdf0b1a33ed3ede4857b6278a4a1656a) |
| `proxy_config.https.tls_cert_params` | [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-dad39f9ccbfc1c7a0ab0fc441dc66ffc5ac7eb3e6b0866318b196507ac076f6f) |
| `proxy_config.https.tls_cert_params.certificates` | [proxy_config.https.tls_cert_params.certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0611780e6bd773be16d67af6f7b5fd6d452d0734ae716ca368d0ae69c777cc49) |
| `proxy_config.https.tls_cert_params.certificates.name` | [proxy_config.https.tls_cert_params.certificates.name](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d798df0c0f749383d960edfedc3ec8916ab68d9a86f7c2cb4364d674467de128) |
| `proxy_config.https.tls_cert_params.certificates.namespace` | [proxy_config.https.tls_cert_params.certificates.namespace](data-sources--bigip_http_proxy--reference--group-003.md#canonical-b08a4ed4b901aa71bd430af88e1112cf5316f291bf75650972ffa71d55f8c6f3) |
| `proxy_config.https.tls_cert_params.certificates.tenant` | [proxy_config.https.tls_cert_params.certificates.tenant](data-sources--bigip_http_proxy--reference--group-003.md#canonical-24e2b2d2f62c13644cad7ea7a02bfaa062cb3482aa67b697508ec48fbb61e25f) |
| `proxy_config.https.tls_cert_params.no_mtls` | [proxy_config.https.tls_cert_params.no_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-92cfa3e06a3aa555ea2769be636086ea2361e0631f38094a469a1287b2fa8626) |
| `proxy_config.https.tls_cert_params.tls_config` | [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-fdd7c41ef9ab58c73a7b5493d547f7ceca5998a15ffaa87cb4c24cc8a43d36c9) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security` | [proxy_config.https.tls_cert_params.tls_config.custom_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-22cdf5f4c0084cfe07da48c299aba7699ae129420aa9ebccfad1797f736131da) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.cipher_suites` | [proxy_config.https.tls_cert_params.tls_config.custom_security.cipher_suites](data-sources--bigip_http_proxy--reference--group-003.md#canonical-93eaef405c1541dea463240928b4a414bd15cad0d25f0f0276a5d3e954b18b81) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.max_version` | [proxy_config.https.tls_cert_params.tls_config.custom_security.max_version](data-sources--bigip_http_proxy--reference--group-003.md#canonical-6c1b6c9fcfc2b617affecfd8e598c9795d135c150e9f6a158f3f03f5ec5d4ef1) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.min_version` | [proxy_config.https.tls_cert_params.tls_config.custom_security.min_version](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ef278476fa0becb653fb48a0300aec1216bbf77cb2ff67db4e846c54fe83a2c5) |
| `proxy_config.https.tls_cert_params.tls_config.default_security` | [proxy_config.https.tls_cert_params.tls_config.default_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-85748ced1bee1f3e10a12fa39991786e38b3f29333d0a4e143eeb1180a50f931) |
| `proxy_config.https.tls_cert_params.tls_config.low_security` | [proxy_config.https.tls_cert_params.tls_config.low_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1086531268895e0a9c9d2a576c780c2aad56121641ac58def9287e55a720cf71) |
| `proxy_config.https.tls_cert_params.tls_config.medium_security` | [proxy_config.https.tls_cert_params.tls_config.medium_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-a48a826a048394e4d09accca26bef4b4a526fcec111578db67cb0a42b05f6b70) |
| `proxy_config.https.tls_cert_params.use_mtls` | [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-fbffd65c11f0b287819064008a0106603475ad221134487763047f31df6108c2) |
| `proxy_config.https.tls_cert_params.use_mtls.client_certificate_optional` | [proxy_config.https.tls_cert_params.use_mtls.client_certificate_optional](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2b2611e990ab9465263f1bc8f04b2b41c805f46fa01d8ce626d22dc2759c3081) |
| `proxy_config.https.tls_cert_params.use_mtls.crl` | [proxy_config.https.tls_cert_params.use_mtls.crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-87b1cb7dd59e861a77f18359da5f94e5ce9fe992a51571a785f0b26c6dd53e47) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.name` | [proxy_config.https.tls_cert_params.use_mtls.crl.name](data-sources--bigip_http_proxy--reference--group-003.md#canonical-40a0880643acbe78132bed6a8e8ae243ea32715b08227613d40bc324720b500b) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.namespace` | [proxy_config.https.tls_cert_params.use_mtls.crl.namespace](data-sources--bigip_http_proxy--reference--group-003.md#canonical-8d41cc7df037d64fa0a04759daeea6d873864fa7d98a13a1b083130f09a76441) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.tenant` | [proxy_config.https.tls_cert_params.use_mtls.crl.tenant](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f39fb51beeb1c29f3a0c452367f37b88991d0e36f76b19534d0ddd6e43427a01) |
| `proxy_config.https.tls_cert_params.use_mtls.no_crl` | [proxy_config.https.tls_cert_params.use_mtls.no_crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-e8b1f086cb7beab7e8412f07b062bbb2087b673d47bfada377c46a8d89a051cd) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca](data-sources--bigip_http_proxy--reference--group-003.md#canonical-78e86b53479cc7ba11186fa9d23584b10f0d9c3649cb1f7bb5caee09b81621ed) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.name` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.name](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f9d926d359d22daece2ca478f04f62a5d6ede7790b813592d95cccf95846cca9) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.namespace` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.namespace](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1f86d16f54fcf2f6295371cb0384538956b8459b76a62d17e72274023d84c45e) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.tenant` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.tenant](data-sources--bigip_http_proxy--reference--group-003.md#canonical-54316ad634518660fc3110085cd67c6327ee70475d99959b4eb64e324b56eece) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca_url` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca_url](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d4074ee8dca765f904a93be3a97ceb7d8723f6778f72e2dfa34bba172554aefa) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled](data-sources--bigip_http_proxy--reference--group-003.md#canonical-bc5e959778ef8730f77b73b2f60ff901a8072c85c329a9bf997da73e2484138d) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_options` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3f046c1ebf3638a0858c7d9078d0845ce11fc433356a76d7d1faf81babcb456a) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1de6bccd1f65d658ace2452fa2599725f4ed1a192d67564120c5ffc81ad38b93) |
| `proxy_config.https.tls_parameters` | [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-616c3576f8ddcbe0b9f675d0038e8d7a78d45d2b4f1a7e956eb5e69f0c969554) |
| `proxy_config.https.tls_parameters.no_mtls` | [proxy_config.https.tls_parameters.no_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-501cfa27afd123dc948b98a32678c2f0c2905ee94c98b1c06f1049eeca513649) |
| `proxy_config.https.tls_parameters.tls_certificates` | [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f8ed05f6a8d9c1731b317ee50da934274e9dc7746aab3209e98db6870df473d3) |
| `proxy_config.https.tls_parameters.tls_certificates.certificate_url` | [proxy_config.https.tls_parameters.tls_certificates.certificate_url](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3333b7597f01e01855df0ce34f29fdcaf09bf17f3b703da54b8949af1f88101e) |
| `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms` | [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms](data-sources--bigip_http_proxy--reference--group-003.md#canonical-59159874ca0574997727ab420b1cd5665a1726c55dc9cbceaad58295870cb466) |
| `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` | [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--bigip_http_proxy--reference--group-003.md#canonical-a02167508df3c2834ccc7bc86abcf344aecef6445ea121d7cbcd6f0bad17c6f9) |
| `proxy_config.https.tls_parameters.tls_certificates.description_spec` | [proxy_config.https.tls_parameters.tls_certificates.description_spec](data-sources--bigip_http_proxy--reference--group-003.md#canonical-10b04008c44f768cddb228fc5d481ce77824a7d3ae64732e1e99db66900feef0) |
| `proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling` | [proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling](data-sources--bigip_http_proxy--reference--group-003.md#canonical-d615d3b515d9481ea20f9dd59191dbe3673f703f9e3851305d789c2c57b3efa1) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key` | [proxy_config.https.tls_parameters.tls_certificates.private_key](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0eafb96d97e189c015b2728668122db5376f55c8c72536cf55a7037bee33dbf1) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--bigip_http_proxy--reference--group-003.md#canonical-a02f905df5a36a4855936b53857a2a35d89d76b788d84465aba0a27d70d82728) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--bigip_http_proxy--reference--group-003.md#canonical-e390e67c2c6f60b4dee55afd9e8f1feb00e9206ebf43755d94e1cdbd2cf53646) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location](data-sources--bigip_http_proxy--reference--group-003.md#canonical-df8f08a784509358b45532877bf0e53fa614260bcefd75447daf95c5507bcb05) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--bigip_http_proxy--reference--group-003.md#canonical-033f72565abea0539b2cd9b4b665c992166bb42cd024f4468af6c618bbe16c4a) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--bigip_http_proxy--reference--group-003.md#canonical-a32c34a570f339a45a0c285cd9021b625c1b3e4337c5c974ae712be4cb8f22da) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--bigip_http_proxy--reference--group-003.md#canonical-260eb859e7a8a9efd1bcbd41ec9cf71eb095220ddc02dcdc618bd6ed5d7f33c6) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9fb601a012444488e9320052bb5130184440d4423e1037096814a7f40880bda4) |
| `proxy_config.https.tls_parameters.tls_certificates.use_system_defaults` | [proxy_config.https.tls_parameters.tls_certificates.use_system_defaults](data-sources--bigip_http_proxy--reference--group-003.md#canonical-9d186f08a180eb5faa48197d1d092092cf2767ee870fd050cf038463cc0f6a4a) |
| `proxy_config.https.tls_parameters.tls_config` | [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-99b1a5e02623cd1d34fc531e98203bd0a768414cd8cda669cb2dbb066f9427eb) |
| `proxy_config.https.tls_parameters.tls_config.custom_security` | [proxy_config.https.tls_parameters.tls_config.custom_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-dab74c818c4054c0c4d8faf2505db7b64529942469a949e9c8c7ebd21c6a363d) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.cipher_suites` | [proxy_config.https.tls_parameters.tls_config.custom_security.cipher_suites](data-sources--bigip_http_proxy--reference--group-003.md#canonical-cd066ff0879df58c3e76146a7ffe5a18d9124cdf03a0054818e875f9b26e9b61) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.max_version` | [proxy_config.https.tls_parameters.tls_config.custom_security.max_version](data-sources--bigip_http_proxy--reference--group-003.md#canonical-08982dc8b3eec74b7d8de02401e83f777844bb9bdc3f8c48becfaccad15b0ca8) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.min_version` | [proxy_config.https.tls_parameters.tls_config.custom_security.min_version](data-sources--bigip_http_proxy--reference--group-003.md#canonical-f9823dfece0c40f23af44eece687cb5d1db21ac1f8257b06624493749a7ae486) |
| `proxy_config.https.tls_parameters.tls_config.default_security` | [proxy_config.https.tls_parameters.tls_config.default_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-8628d1e0dde400350dea5d17041c91fcfde3a91d386ed85ff663907db1b2f8b8) |
| `proxy_config.https.tls_parameters.tls_config.low_security` | [proxy_config.https.tls_parameters.tls_config.low_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-68b5605ff70509746a3aa996744a2f4f16b665e6268c85fc228aeb92d4d6b5ba) |
| `proxy_config.https.tls_parameters.tls_config.medium_security` | [proxy_config.https.tls_parameters.tls_config.medium_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-c59a90728296f4bee9e7132a3890c2ffe3dc3c7695f3f38718b347e12cf45da9) |
| `proxy_config.https.tls_parameters.use_mtls` | [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-646ac82f0848ff79909c59a0584afaa46ed95ca6deb78715afca511ccc27ff62) |
| `proxy_config.https.tls_parameters.use_mtls.client_certificate_optional` | [proxy_config.https.tls_parameters.use_mtls.client_certificate_optional](data-sources--bigip_http_proxy--reference--group-003.md#canonical-6aa523923112c9e3b317960a51a578529eb155f301923472e4a90242f866d56b) |
| `proxy_config.https.tls_parameters.use_mtls.crl` | [proxy_config.https.tls_parameters.use_mtls.crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-ac282277f0bfac738e807036bc78f94d0fa4f31c48f83f6e2c059f052da40b3e) |
| `proxy_config.https.tls_parameters.use_mtls.crl.name` | [proxy_config.https.tls_parameters.use_mtls.crl.name](data-sources--bigip_http_proxy--reference--group-003.md#canonical-088cb962e8083098b42993fc7769fe21401a492f276ac5efe7e627b8bc9f1180) |
| `proxy_config.https.tls_parameters.use_mtls.crl.namespace` | [proxy_config.https.tls_parameters.use_mtls.crl.namespace](data-sources--bigip_http_proxy--reference--group-003.md#canonical-a193b3fd0778488c6d732b5e8e814998215c68ec4156075cad3cd91301cc9886) |
| `proxy_config.https.tls_parameters.use_mtls.crl.tenant` | [proxy_config.https.tls_parameters.use_mtls.crl.tenant](data-sources--bigip_http_proxy--reference--group-003.md#canonical-acae7262f19349826969a4e4b975cc320376b3d7e5946de7bf2c7cf6c123ae71) |
| `proxy_config.https.tls_parameters.use_mtls.no_crl` | [proxy_config.https.tls_parameters.use_mtls.no_crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-6d651873c3744c970e0cf679680fd08e0ec6cbf3864dd22005fe74f5d8aab46b) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca](data-sources--bigip_http_proxy--reference--group-004.md#canonical-ec620ed4e1ceb922a41bc128e01244aeb0fb1aa8623271798d2ef1a91d157fe0) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.name` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.name](data-sources--bigip_http_proxy--reference--group-004.md#canonical-7c8e38d2ac3632dc782a06491206e36a12ad9d448b120239ee43bfd9ea27365a) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.namespace` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.namespace](data-sources--bigip_http_proxy--reference--group-004.md#canonical-cfc89534b27cc3ccb20932fef8ba9d1ab55f2b664ebd7521101229abe80df15b) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.tenant` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.tenant](data-sources--bigip_http_proxy--reference--group-004.md#canonical-ea69f670383837a7d08e818e549e30186184511ebc01fd55102a560dc6c57709) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca_url` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca_url](data-sources--bigip_http_proxy--reference--group-003.md#canonical-23f63b3bb5124a716a10c3a548bd5e6ed44cb4070c8c44589ef422826ed2ae31) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_disabled` | [proxy_config.https.tls_parameters.use_mtls.xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-c566a03e11427042823e090396dab37a064b816e7a48ccc8597dd71e5ec61286) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_options` | [proxy_config.https.tls_parameters.use_mtls.xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-bcaaab5fecf7103bbbc411b4950be077e8ee04e5bce81f2015e202487039407c) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements](data-sources--bigip_http_proxy--reference--group-004.md#canonical-a677d94a10e6dac851af0f440efedaa8060a47b153a4ce37b4ceec210709e9cc) |
| `proxy_config.https_auto_cert` | [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3ef7c36b7533a3872aefaec38bda05195fab6b72e8121c8e1d6a0e72748ced09) |
| `proxy_config.https_auto_cert.add_hsts` | [proxy_config.https_auto_cert.add_hsts](data-sources--bigip_http_proxy--reference--group-004.md#canonical-11ef11f244c032c128274e0d8b33e9c33f3e9165d4ec6e756a70710d8d226690) |
| `proxy_config.https_auto_cert.append_server_name` | [proxy_config.https_auto_cert.append_server_name](data-sources--bigip_http_proxy--reference--group-004.md#canonical-63e6def0fd14e9812233c0c633eba736ba8f5dad4cd9c005c0f4411e1e0e01fc) |
| `proxy_config.https_auto_cert.coalescing_options` | [proxy_config.https_auto_cert.coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-bbb5cec6e4b3a7739d2082fc84ba3ea56bcbc0c177d061a1f55974bc76faf6ea) |
| `proxy_config.https_auto_cert.coalescing_options.default_coalescing` | [proxy_config.https_auto_cert.coalescing_options.default_coalescing](data-sources--bigip_http_proxy--reference--group-004.md#canonical-a071341990ceb7e43f3b40fbefdfd2e5a9a34ca0e3a4e409a8307f83e54753cf) |
| `proxy_config.https_auto_cert.coalescing_options.strict_coalescing` | [proxy_config.https_auto_cert.coalescing_options.strict_coalescing](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d6602addce1efc876ec1ca52777748472aec7afde4bfff18c072d48c5654ed00) |
| `proxy_config.https_auto_cert.connection_idle_timeout` | [proxy_config.https_auto_cert.connection_idle_timeout](data-sources--bigip_http_proxy--reference--group-004.md#canonical-438ed26a3fdce77bf3bad266aec8219c634da7463fa0f7be32bf74c34504ed70) |
| `proxy_config.https_auto_cert.default_header` | [proxy_config.https_auto_cert.default_header](data-sources--bigip_http_proxy--reference--group-004.md#canonical-68a7514bc893ffc99777dbe92341adca96b07f7b4253efbcf771b59c0c8aa822) |
| `proxy_config.https_auto_cert.default_loadbalancer` | [proxy_config.https_auto_cert.default_loadbalancer](data-sources--bigip_http_proxy--reference--group-004.md#canonical-bab057f1e649f6380a4829e1e0b2bb71b407d0ccd4edde69534f5dd5cc0342bd) |
| `proxy_config.https_auto_cert.disable_path_normalize` | [proxy_config.https_auto_cert.disable_path_normalize](data-sources--bigip_http_proxy--reference--group-004.md#canonical-fddab9fbf88f3ff36a7336a03db0d2ec17b6477418f6228329bcad08a4449460) |
| `proxy_config.https_auto_cert.enable_path_normalize` | [proxy_config.https_auto_cert.enable_path_normalize](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3719345eb6af3051400c1d8ff89f930fad7629d99a6299357e4bb43afe3dec77) |
| `proxy_config.https_auto_cert.http_protocol_options` | [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-8c8e0d531f4bbe24219dfa3f08c4cf6803f4a90ba3520f6c615b0d4962d82cb4) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-7dfa1a82f6ed530298162a343af328933659ca2870d7ecaa2a1d976da5718c44) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-03d42450565ad2264a99a380d6ebd2afb5eaa4a8fdf4152a8ba6b5bc99e8abc2) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-67a1e7ca424f152f15a604cb9e9ac6c7ea7ea89c75b63d0415f761473f4bf0c8) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-a57b478ee4cb6bb39562163494a4d20203ec92acb232ac0a08eba91e33e8b9ce) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-334154287faaa9133a6dabb103451b6b887245fbb1659c288bec322aef5ff631) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--reference--group-004.md#canonical-4c49064506ddda758ded0110b3d9161a8e3073cf3d9bbe238dff246afc3e35f0) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-f0e9f8b0bb40b58745ae1b62c6160ae352b8ec117eda19ba20c5668cf2f0bca1) |
| `proxy_config.https_auto_cert.http_redirect` | [proxy_config.https_auto_cert.http_redirect](data-sources--bigip_http_proxy--reference--group-004.md#canonical-e3dbd983b0f2a1880cbae2b1e5d4f6b0094c41b35df6244e57a046a238c1d0de) |
| `proxy_config.https_auto_cert.no_mtls` | [proxy_config.https_auto_cert.no_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-bc9f5cadd0b39139fac8ac5bb8a65857a83c71b32770725c722785da1fbb4d48) |
| `proxy_config.https_auto_cert.non_default_loadbalancer` | [proxy_config.https_auto_cert.non_default_loadbalancer](data-sources--bigip_http_proxy--reference--group-004.md#canonical-b6525bfc55510fbb7cfccd920e0764fd5408eb2a2289f8c13db5fd69f34aa8f7) |
| `proxy_config.https_auto_cert.pass_through` | [proxy_config.https_auto_cert.pass_through](data-sources--bigip_http_proxy--reference--group-004.md#canonical-c6e3563e7bc2d3575a8f74ceabae2d560aed7d9f4fa4eede30f5627af296bc9f) |
| `proxy_config.https_auto_cert.port` | [proxy_config.https_auto_cert.port](data-sources--bigip_http_proxy--reference--group-004.md#canonical-bd8f7704b75afbfb9f986bdf5d5e2ef972caddb74e304d86ba93cd79783dc760) |
| `proxy_config.https_auto_cert.port_ranges` | [proxy_config.https_auto_cert.port_ranges](data-sources--bigip_http_proxy--reference--group-004.md#canonical-e79a28b026201900ab970350b70111fefc855bd3863e689e7925d1be31e559cf) |
| `proxy_config.https_auto_cert.server_name` | [proxy_config.https_auto_cert.server_name](data-sources--bigip_http_proxy--reference--group-004.md#canonical-c3fbc1da1922128477e527330361c0ae0384272843da89ee35fe592c0a401f05) |
| `proxy_config.https_auto_cert.tls_config` | [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-a35e8126cfda6aadcc344ef3307215e5672c81e249e1ddc35b2547d32a8b0599) |
| `proxy_config.https_auto_cert.tls_config.custom_security` | [proxy_config.https_auto_cert.tls_config.custom_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-c1ff9624bea553594b16f98585cb5a011521d4b1c8272ed48883febcc1ea71f4) |
| `proxy_config.https_auto_cert.tls_config.custom_security.cipher_suites` | [proxy_config.https_auto_cert.tls_config.custom_security.cipher_suites](data-sources--bigip_http_proxy--reference--group-004.md#canonical-6af76e7ff875a52d9f5334e5c4f894b9f784ea2c616095ee0474ea581e9a87c1) |
| `proxy_config.https_auto_cert.tls_config.custom_security.max_version` | [proxy_config.https_auto_cert.tls_config.custom_security.max_version](data-sources--bigip_http_proxy--reference--group-004.md#canonical-fe871a0a492246a95eaf3f9d3e612d72ff2e3682ec13c3f881444406c22456c8) |
| `proxy_config.https_auto_cert.tls_config.custom_security.min_version` | [proxy_config.https_auto_cert.tls_config.custom_security.min_version](data-sources--bigip_http_proxy--reference--group-004.md#canonical-69fd529f4ad4e69d4cba3f58682852528acaf23f5c4b98e5b633a4dbe48b1bb4) |
| `proxy_config.https_auto_cert.tls_config.default_security` | [proxy_config.https_auto_cert.tls_config.default_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-5cd7b6fff6c17f0d8945190fc6a51417e51cd6b1510dd999d3750e732ed93cc1) |
| `proxy_config.https_auto_cert.tls_config.low_security` | [proxy_config.https_auto_cert.tls_config.low_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-51cd6c49c397113c7d9654f076a5b1285babef9a3c194293ebe1f99089939649) |
| `proxy_config.https_auto_cert.tls_config.medium_security` | [proxy_config.https_auto_cert.tls_config.medium_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-9b7ce7bdd92207f2f268d8468ecde3f620d93bf316fd0e4c2a6d02e497e27c54) |
| `proxy_config.https_auto_cert.use_mtls` | [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d0594eddbafbb60eb094435280b88c5d9924a63e3c2f06196e1240709e017e3e) |
| `proxy_config.https_auto_cert.use_mtls.client_certificate_optional` | [proxy_config.https_auto_cert.use_mtls.client_certificate_optional](data-sources--bigip_http_proxy--reference--group-004.md#canonical-b336fd6da0d447195d74038f94238cd41e451ef8180fccde78458702174ae638) |
| `proxy_config.https_auto_cert.use_mtls.crl` | [proxy_config.https_auto_cert.use_mtls.crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-aca153d795edc745a50e74983fb0ccc156f05684b0079826168b36cad00e8b57) |
| `proxy_config.https_auto_cert.use_mtls.crl.name` | [proxy_config.https_auto_cert.use_mtls.crl.name](data-sources--bigip_http_proxy--reference--group-004.md#canonical-30c01d252bcb923cdc8d468646a08faec63573709ced712fe5378bc856dc2a09) |
| `proxy_config.https_auto_cert.use_mtls.crl.namespace` | [proxy_config.https_auto_cert.use_mtls.crl.namespace](data-sources--bigip_http_proxy--reference--group-004.md#canonical-00a8f0bf60ef21ea28dcebb8633175d5dd09c30f3b822831251765966ea906a6) |
| `proxy_config.https_auto_cert.use_mtls.crl.tenant` | [proxy_config.https_auto_cert.use_mtls.crl.tenant](data-sources--bigip_http_proxy--reference--group-004.md#canonical-64736012219c8c50f99730a24c368b044b7347f2266e951ff3814d2c91387aae) |
| `proxy_config.https_auto_cert.use_mtls.no_crl` | [proxy_config.https_auto_cert.use_mtls.no_crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-c1e718b10e3ecded681ca3c1d8fa663555fa1e95668b2b8f6f5f96959e2f427a) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca` | [proxy_config.https_auto_cert.use_mtls.trusted_ca](data-sources--bigip_http_proxy--reference--group-004.md#canonical-808a9d760b6e2a15bc2ccfa325774e97629d00dc834cc3d816ffc2cb02b6c21c) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.name` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.name](data-sources--bigip_http_proxy--reference--group-004.md#canonical-901879cf5b60ec2ec06d6b8675699ae758106005e45d35be09091d891b8eda21) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.namespace` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.namespace](data-sources--bigip_http_proxy--reference--group-004.md#canonical-fb8780a6919416e2aa5954d2e5412401e17d2b4edaa9164af714e4b9d5acbe56) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.tenant` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.tenant](data-sources--bigip_http_proxy--reference--group-004.md#canonical-00c14dcbd46aebc2cdaa86a92653c46ae33abaec81e00d843e82ae9b6a9ed111) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca_url` | [proxy_config.https_auto_cert.use_mtls.trusted_ca_url](data-sources--bigip_http_proxy--reference--group-004.md#canonical-4ce2812effc4b4261a9a4f34ddc76c24b6688284d43dafe808faf25b97050355) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_disabled` | [proxy_config.https_auto_cert.use_mtls.xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-f0a43bbde557cb445f489e69d7bd8a58ff91827490fe8607c73dbd694c4ebd0f) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_options` | [proxy_config.https_auto_cert.use_mtls.xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-8fa2d528f11ea7f6af06a9e047eebad3e86971d16755f9a1887a5a7b6db64afa) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements](data-sources--bigip_http_proxy--reference--group-004.md#canonical-d21b30765a9e3839590c3dba743bcb2e331387dd3852aaf91d367ee478c63829) |

<a id="canonical-4d563bba3e683ad608f22d42c29b0d107bc284f95fd53cdbc5a02035abba1ca7"></a>

## Next pages — Property reference / 2150a054bb1b / 11

- [advanced_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-250c9f8b07036fc33dc6f787805f3dca2d5b508170e66d907330dacafe89c52b)
- [ddos_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-22b749b9380e691ba8bfd877b93a99f2cc8cabdce17fb4a0bda6eb12c4b504aa)
- [irules](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1a339565aa9a15a2645607264cca4d58d053f037f8bd0e3e846158f64d436fb3)
- [lb_algorithm](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0ed3b2daa8855618fec8a4a5e7a5ce44b295dceba3b83a658f4d469b9cb0d5c1)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-c26b7c8b16a4eda57dc835a575d989b32c52cf2b530dca551eec308ac24d31db)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-002.md#canonical-9adae87539ddf68461f98c5221e9bb56542c9758439b42deb575137fbcb1d13a)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-250c9f8b07036fc33dc6f787805f3dca2d5b508170e66d907330dacafe89c52b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-494e9b5a4d1610291c04987a9e60725c709668c0c09109e78e1b6c027e84084e"></a>

## advanced_profile — advanced_profile / 79c018a21e11 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- advanced_profile

<a id="canonical-6e2425007626682d10c3948dce4fd63e6ef4a94644c9681331a0bd5c369608f4"></a>

Type: `"single"`. Computed.

Defines various advanced Profile OPTIONS for a Loadbalancer.

Upstream description:

This defines various advanced Profile OPTIONS for a Loadbalancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"disable\",\"enable_default_profile\"]"
}
```

<a id="canonical-1e33ca061c26005a5f87777852146815695afa5941f87d4ed31664891dae6c98"></a>

## Direct properties — advanced_profile / 79c018a21e11 / 3

- [disable_spec](data-sources--bigip_http_proxy--reference--group-001.md#canonical-85e44722647fddbf7a1c793da78a201e7bb211d8b015c74045e17dd98251c389): complete subsection reference.

- [enable_default_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-c97107662fdadf8c4acd0f2a957afc6f7470368a79c9adfae3f378adea820880): complete subsection reference.

<a id="canonical-36e8545c5ff7beb0d7ac6b2cfdcf7c0313a077215971438568d9d02149cf3f63"></a>

## Next pages — advanced_profile / 79c018a21e11 / 4

- [advanced_profile.disable_spec](data-sources--bigip_http_proxy--reference--group-001.md#canonical-85e44722647fddbf7a1c793da78a201e7bb211d8b015c74045e17dd98251c389)
- [advanced_profile.enable_default_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-c97107662fdadf8c4acd0f2a957afc6f7470368a79c9adfae3f378adea820880)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-85e44722647fddbf7a1c793da78a201e7bb211d8b015c74045e17dd98251c389"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98adf8e73db187cb0da8b96e29cc580497739cfa199bfa041cec21ae35d55179"></a>

## advanced_profile.disable_spec — advanced_profile.disable_spec / f844026153f7 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [advanced_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-250c9f8b07036fc33dc6f787805f3dca2d5b508170e66d907330dacafe89c52b)
- advanced_profile.disable_spec

<a id="canonical-c970693d1e47d753d5109009b0bb3452eb973e14a1bb15c48527ff843dd233fd"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-0ad37c88b8e56231dabf064be20220b3eef3f1a204be2ffd1be2ac67743a55b1"></a>

## Direct properties — advanced_profile.disable_spec / f844026153f7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eb5deb38767c56e2dd49e4b77d2dcd6175ecf92edb1f1af1715614d7284301dd"></a>

## Next pages — advanced_profile.disable_spec / f844026153f7 / 4

- [advanced_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-250c9f8b07036fc33dc6f787805f3dca2d5b508170e66d907330dacafe89c52b)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-c97107662fdadf8c4acd0f2a957afc6f7470368a79c9adfae3f378adea820880"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-646bd21cbfb2a17a2ce688630aa98f9e812c5575e45d447b3d9bc37011fe9fbb"></a>

## advanced_profile.enable_default_profile — advanced_profile.enable_default_profile / f4347e29ac9e / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [advanced_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-250c9f8b07036fc33dc6f787805f3dca2d5b508170e66d907330dacafe89c52b)
- advanced_profile.enable_default_profile

<a id="canonical-f51e8b4fd59d7c5a80bcdcf205babc734ca7d8bc867ca44840864ac6ea437e60"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable default profile.

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

<a id="canonical-81836f0b44eeced4d02c56c94daa9f19c5766ff94f9454373324f526c2c88caf"></a>

## Direct properties — advanced_profile.enable_default_profile / f4347e29ac9e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3047ca156f528abc960b74f6c22d8434cf1524bdbf187c375d4de4142a2eb579"></a>

## Next pages — advanced_profile.enable_default_profile / f4347e29ac9e / 4

- [advanced_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-250c9f8b07036fc33dc6f787805f3dca2d5b508170e66d907330dacafe89c52b)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-22b749b9380e691ba8bfd877b93a99f2cc8cabdce17fb4a0bda6eb12c4b504aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10ae3c8164595cdd41aa7a048e56b0b8163c05cf3100b8b6b9e94603c83d41ff"></a>

## ddos_profile — ddos_profile / bccc8d3c7879 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- ddos_profile

<a id="canonical-f1332d9aaa34f7b5afa58b65ea9cdde426d7f01e04ff6dac5d157186c625fb07"></a>

Type: `"single"`. Computed.

Configuration parameter for ddos profile.

Upstream description:

BIG-IP DDoS Protection Rules.

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

<a id="canonical-a4fc14c69d757579f30a9e0a535d2b5c2c4b66962f6bd395ba83f7d7dacee1ee"></a>

## Direct properties — ddos_profile / bccc8d3c7879 / 3

- [disable_ddos_mitigation](data-sources--bigip_http_proxy--reference--group-001.md#canonical-425ca0ffe939ba761430b3c3b7740743165e50dec3b6151ce982d7b5d63d7527): complete subsection reference.

- [enable_ddos_mitigation](data-sources--bigip_http_proxy--reference--group-001.md#canonical-a9a327f365190942908f51123164bdbe3fdf40ba11db43094fc854a14be39baf): complete subsection reference.

<a id="canonical-8646df873da0b9128002a529f3ce0cbf80cfb5bbbd816f1d9d8fd1e9924c6e70"></a>

## Next pages — ddos_profile / bccc8d3c7879 / 4

- [ddos_profile.disable_ddos_mitigation](data-sources--bigip_http_proxy--reference--group-001.md#canonical-425ca0ffe939ba761430b3c3b7740743165e50dec3b6151ce982d7b5d63d7527)
- [ddos_profile.enable_ddos_mitigation](data-sources--bigip_http_proxy--reference--group-001.md#canonical-a9a327f365190942908f51123164bdbe3fdf40ba11db43094fc854a14be39baf)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-425ca0ffe939ba761430b3c3b7740743165e50dec3b6151ce982d7b5d63d7527"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b92b3af77421ece6fe8c8b8de2b21ba1c2a8f50d2bcbf59d3ffb6e6ce5476be"></a>

## ddos_profile.disable_ddos_mitigation — ddos_profile.disable_ddos_mitigation / f8c4ddef67f8 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [ddos_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-22b749b9380e691ba8bfd877b93a99f2cc8cabdce17fb4a0bda6eb12c4b504aa)
- ddos_profile.disable_ddos_mitigation

<a id="canonical-38d597a3379d2f6c1dca4249f560ac1aa6f02a05c8f66b6d58768c35e7c2a66f"></a>

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

<a id="canonical-aa19724ae9f92d52b0c15dba56dfa424da6240f08c34713f7e16a0bba91e2d6b"></a>

## Direct properties — ddos_profile.disable_ddos_mitigation / f8c4ddef67f8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-483e32125ba649e4ac6f95b65222e164f3b409507e52925852d876f5eeacc12e"></a>

## Next pages — ddos_profile.disable_ddos_mitigation / f8c4ddef67f8 / 4

- [ddos_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-22b749b9380e691ba8bfd877b93a99f2cc8cabdce17fb4a0bda6eb12c4b504aa)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-a9a327f365190942908f51123164bdbe3fdf40ba11db43094fc854a14be39baf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b1ee734948c8162af4927db6db76cd355c78ed6aa8c08295f24507e1e336112"></a>

## ddos_profile.enable_ddos_mitigation — ddos_profile.enable_ddos_mitigation / 50de318f6c1f / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [ddos_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-22b749b9380e691ba8bfd877b93a99f2cc8cabdce17fb4a0bda6eb12c4b504aa)
- ddos_profile.enable_ddos_mitigation

<a id="canonical-e7b7684df6299ca5c34a1c8043f9f8620cad60f39efb38409df2472b6aae105a"></a>

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

<a id="canonical-f6e2f4bba6662fdac8114abac5ce4537d379656db4d45ad3179c011d051f51a9"></a>

## Direct properties — ddos_profile.enable_ddos_mitigation / 50de318f6c1f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bc61af12cb1f828e0c4f869addfd6426eea201294cc1aa3f1c533bb9e8ce95bb"></a>

## Next pages — ddos_profile.enable_ddos_mitigation / 50de318f6c1f / 4

- [ddos_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-22b749b9380e691ba8bfd877b93a99f2cc8cabdce17fb4a0bda6eb12c4b504aa)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-1a339565aa9a15a2645607264cca4d58d053f037f8bd0e3e846158f64d436fb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ad9d1e1e1967f305fb308ebbffe791c5239aec00c5135a01bc0106da60f1d0e"></a>

## irules — irules / 42bf155c5710 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- irules

<a id="canonical-eca18e93ce30aebe953dc739dfe58fc83b49d595ac19385f06bbf300ae44061b"></a>

Type: `"single"`. Computed.

IRules Configuration for downstream connections.

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

<a id="canonical-9c248f857804c903cf96f51f53a3ed24056065123d76d7af4aabca31d7b376c1"></a>

## Direct properties — irules / 42bf155c5710 / 3

- [irules](data-sources--bigip_http_proxy--reference--group-001.md#canonical-4a445d27381ca6d69010ccb3b72b404d793209ecf6214de50f43af9305551b2f): complete subsection reference.

<a id="canonical-c8ab3bb174c48272640466f8b760a60ba7bc50ea6b8b4033eb2aacd06d6d1655"></a>

## Next pages — irules / 42bf155c5710 / 4

- [irules.irules](data-sources--bigip_http_proxy--reference--group-001.md#canonical-4a445d27381ca6d69010ccb3b72b404d793209ecf6214de50f43af9305551b2f)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-4a445d27381ca6d69010ccb3b72b404d793209ecf6214de50f43af9305551b2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6587aa1f33817045ca79f715df246e4e8007a5db32bb743d897955c6c9170fdf"></a>

## irules.irules — irules.irules / 8eb92576858d / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [irules](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1a339565aa9a15a2645607264cca4d58d053f037f8bd0e3e846158f64d436fb3)
- irules.irules

<a id="canonical-09d34740298a272e9a8966fa8ca7225e4aac821a8111af2950e189f01d7ae826"></a>

Type: `"list"`. Computed.

OPTIONS for attaching iRules to BIG-IP HTTP Proxy.

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

<a id="canonical-fbd8c6bd2fd15dfcbb883d048b4bd735eb0d2173c64ab63d9bee18f3a2b271af"></a>

## Direct properties — irules.irules / 8eb92576858d / 3

<a id="canonical-94988e7bffb7261e2f7067375a3b79917d4e3fea60f7e06208981c41419d31dd"></a>

<a id="canonical-b9b8533d603a4c71a4243000af7bab1d1d1697eb4d5b0a9b778edc0ef22a569c"></a>

## name property — irules.irules / 8eb92576858d / 4

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

<a id="canonical-ffcb7693a5f3f703bf4272b2d13c5c9f221fe182c3052c73a87ec76b40d5da42"></a>

<a id="canonical-3002b4d33827079018b58110b1297207707f077b88b370357e64828c1d2d8689"></a>

## namespace property — irules.irules / 8eb92576858d / 5

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

<a id="canonical-a108d8b8d4b13d3b1cd68907b4f33b2e39fa903cc79ddf4d7da57119fb0077e2"></a>

<a id="canonical-36d2197f7e65cd9339a44ed4bc4b0f29dfb43adba499c9f7c4769a845c18944c"></a>

## tenant property — irules.irules / 8eb92576858d / 6

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

<a id="canonical-d687f988dcf01c80c740a47726886f27dd8d236f436419a1906f11e434866a1e"></a>

## Next pages — irules.irules / 8eb92576858d / 7

- [irules](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1a339565aa9a15a2645607264cca4d58d053f037f8bd0e3e846158f64d436fb3)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-0ed3b2daa8855618fec8a4a5e7a5ce44b295dceba3b83a658f4d469b9cb0d5c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1cc403460c71833246747117204f5d8384f29c36ba8ed01584315a9d5724305"></a>

## lb_algorithm — lb_algorithm / 1f4cfa9c7811 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- lb_algorithm

<a id="canonical-4db29f6077710634cbc03be6d95e80e1eaf89b8c9f831425327b42e29db1eef6"></a>

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

<a id="canonical-c7a764614e84effc4eadd36433cadca7464e64babd875e45f489ad45a22b599c"></a>

## Direct properties — lb_algorithm / 1f4cfa9c7811 / 3

- [round_robin](data-sources--bigip_http_proxy--reference--group-001.md#canonical-aa79b4b93462f475eb7ae4bfcbd2d1342d1f5c9f60ec1609c1b3af0495f56422): complete subsection reference.

<a id="canonical-9b0db5e0547670d9aa621c57f8b9b764973caaa64f2823f5b281d070ee8b863d"></a>

## Next pages — lb_algorithm / 1f4cfa9c7811 / 4

- [lb_algorithm.round_robin](data-sources--bigip_http_proxy--reference--group-001.md#canonical-aa79b4b93462f475eb7ae4bfcbd2d1342d1f5c9f60ec1609c1b3af0495f56422)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-aa79b4b93462f475eb7ae4bfcbd2d1342d1f5c9f60ec1609c1b3af0495f56422"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee3a7d0ed7ebd49780b8ef10c2b53436facf314df3a3f80a89448f5d54f81ef4"></a>

## lb_algorithm.round_robin — lb_algorithm.round_robin / 41b9cd8939f7 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [lb_algorithm](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0ed3b2daa8855618fec8a4a5e7a5ce44b295dceba3b83a658f4d469b9cb0d5c1)
- lb_algorithm.round_robin

<a id="canonical-a257d7c48971dad9a095eb18c00726fcf705da9c9acbd40d24fa302d64b9416d"></a>

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

<a id="canonical-f194dbc91b7637c737986fc8e9ec7fd8aecfc4e22c13b294c49cc6a5b50aff89"></a>

## Direct properties — lb_algorithm.round_robin / 41b9cd8939f7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5eec57f8c1aa2d9f00813393d8e989a604f41123c27f11ffff00c1c2145f498d"></a>

## Next pages — lb_algorithm.round_robin / 41b9cd8939f7 / 4

- [lb_algorithm](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0ed3b2daa8855618fec8a4a5e7a5ce44b295dceba3b83a658f4d469b9cb0d5c1)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4d17508dd500fcc62b3c61a1e4c2082fe4ffee2195e51b7ca6e9d56bb74d15f"></a>

## origin_pools — origin_pools / 96a68391e0e6 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- origin_pools

<a id="canonical-d98b635e4877934e47542e8fde8ff65642ea9933538a0335a9d06cd44e7f9741"></a>

Type: `"single"`. Computed.

Configuration parameter for origin pools.

Upstream description:

List of Origin Pools.

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

<a id="canonical-4dad3ebc4afb4e35841c09e79f01265295d486e2cfbece1a76f2654720f2485c"></a>

## Direct properties — origin_pools / 96a68391e0e6 / 3

- [pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a): complete subsection reference.

<a id="canonical-dc71717d86821cf182519a51cb11ad6a58da5bc53ccf082b688e63d6217fbb22"></a>

## Next pages — origin_pools / 96a68391e0e6 / 4

- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f68a22c7a9d43b3150840a2499e3cb144649dfa0026b4857b925aa83746bef1e"></a>

## origin_pools.pools — origin_pools.pools / 0ebb00a56d63 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e)
- origin_pools.pools

<a id="canonical-60ef50d97aa68385c529dc2d2851cd4f6749b6c7e8d9bc8f84af0d7067c4968f"></a>

Type: `"list"`. Computed.

Origin Pools. List of Origin Pools.

Upstream description:

List of Origin Pools.

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-dff77c2a1c72f5b6113e471ba61cbae820668daee43f876dd73faf967344c07e"></a>

## Direct properties — origin_pools.pools / 0ebb00a56d63 / 3

<a id="canonical-109de3f27cc8145458fd5f42c6bfd2245337f935fffaca622288fc83224de8d9"></a>

<a id="canonical-3b25c0cb8fab63f98dcd86b1650fc0ea71ed83f6107c4a548fd5a03383773508"></a>

## name property — origin_pools.pools / 0ebb00a56d63 / 4

Type: `"string"`. Computed.

Name. Name of the origin pool.

Upstream description:

Name of the origin pool.

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

- [origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d): complete subsection reference.

<a id="canonical-9b3b6f538036ee7a761b8517b1f37eaee47e2e22b143884121699ade57bc0bb0"></a>

<a id="canonical-41342091e0c67130ad974b5e8ba10be47498ce8494636684cd8a790e17a41300"></a>

## priority property — origin_pools.pools / 0ebb00a56d63 / 5

Type: `"number"`. Computed.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool. When active origin pool is not available, lower priority origin
pools are made active as per the increasing priority.

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool. When active origin pool is not available, lower priority origin
pools are made active as per the increasing priority.

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
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-51a602ae27991150ae4aa6c4ab7a5445e215df744e206bd0779a1f1154f59ed7"></a>

<a id="canonical-be1809e5fddbb2700235e1e33fe7353816d87ff786ee40a3bd9a3c8f5def362c"></a>

## weight property — origin_pools.pools / 0ebb00a56d63 / 6

Type: `"number"`. Computed.

Weight of this origin pool, valid only with multiple origin pools. Value of 0 will disable the pool.

Upstream description:

Weight of this origin pool, valid only with multiple origin pools. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f1dcd2dd6d793d42548401e519da5daa19c138ca41f338840db4674529a4c62b"></a>

## Next pages — origin_pools.pools / 0ebb00a56d63 / 7

- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74eec1de186cec4dafd84ac6672b9c25b03bf51dbbaf59bdb7847be880d07cda"></a>

## origin_pools.pools.origin_servers — origin_pools.pools.origin_servers / ea91ed5bced2 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a)
- origin_pools.pools.origin_servers

<a id="canonical-4cc0fcc14bbd64516a356f69bf54a519d0324b1a87fbe57cd86e447d18fcd8d7"></a>

Type: `"single"`. Computed.

List of origin Servers for the BIG-IP HTTP Proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"automatic_port\",\"lb_port\",\"port\"]"
}
```

<a id="canonical-b9074469e30d607355dec660d8917adab7363f2fd5ba05ec543b42e2680d819f"></a>

## Direct properties — origin_pools.pools.origin_servers / ea91ed5bced2 / 3

- [automatic_port](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1927e85cbf72a371246ae4ae8983ec46f85c579c53d33ba9d7a403c6bde93584): complete subsection reference.

- [health_checks](data-sources--bigip_http_proxy--reference--group-001.md#canonical-5e52b07d1897a4a240c2ded9625edb68e1da365bccee4269134f3bf9e1b95b8a): complete subsection reference.

- [lb_port](data-sources--bigip_http_proxy--reference--group-001.md#canonical-9f7280eca68fe7a0a5c065150caafe90d41f06f99ddfe7b6b6280aa7f14550fb): complete subsection reference.

- [origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-355d4acc28374bc67f053c73c5fc1c3e216af67fb65258137fc871ddb6b9a11c): complete subsection reference.

<a id="canonical-668c76ceac2344a8e13a3c59880be784b4736d56bce7a346fa7154e31868c691"></a>

<a id="canonical-4a5fafef7fa24709972bc698f14fd0f4ec1ba2ba5bc8d27aaa26cc8116097fd5"></a>

## port property — origin_pools.pools.origin_servers / ea91ed5bced2 / 4

Type: `"number"`. Computed.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

Upstream description:

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

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

<a id="canonical-bd6aa0ad8998b9e6fd2bd31e61cf697bccd6c09fe16076260a3924f65f62710d"></a>

## Next pages — origin_pools.pools.origin_servers / ea91ed5bced2 / 5

- [origin_pools.pools.origin_servers.automatic_port](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1927e85cbf72a371246ae4ae8983ec46f85c579c53d33ba9d7a403c6bde93584)
- [origin_pools.pools.origin_servers.health_checks](data-sources--bigip_http_proxy--reference--group-001.md#canonical-5e52b07d1897a4a240c2ded9625edb68e1da365bccee4269134f3bf9e1b95b8a)
- [origin_pools.pools.origin_servers.lb_port](data-sources--bigip_http_proxy--reference--group-001.md#canonical-9f7280eca68fe7a0a5c065150caafe90d41f06f99ddfe7b6b6280aa7f14550fb)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-355d4acc28374bc67f053c73c5fc1c3e216af67fb65258137fc871ddb6b9a11c)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-1927e85cbf72a371246ae4ae8983ec46f85c579c53d33ba9d7a403c6bde93584"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cf75d9518bf8b4bf4af95e5f197a1013c3a9003009caec40e63291d8af43a1f"></a>

## origin_pools.pools.origin_servers.automatic_port — origin_pools.pools.origin_servers.automatic_port / 705a937ac7a0 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- origin_pools.pools.origin_servers.automatic_port

<a id="canonical-b95648a18b7ba9534f767436d523e69278405250bb0b8e2387446c15fc217c24"></a>

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

<a id="canonical-7fcca1f0e1e88d0be5e6705abacc178da97f7c2e1a5b33ed7f216263260375e9"></a>

## Direct properties — origin_pools.pools.origin_servers.automatic_port / 705a937ac7a0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-47188bff80e8379df2d483a337a210eb077bfb67339c212e421bc6b0a5a04e9c"></a>

## Next pages — origin_pools.pools.origin_servers.automatic_port / 705a937ac7a0 / 4

- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-5e52b07d1897a4a240c2ded9625edb68e1da365bccee4269134f3bf9e1b95b8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-681a4720d7342f8d0790388b290616f8ab77c79c321277babcb73d88afea3516"></a>

## origin_pools.pools.origin_servers.health_checks — origin_pools.pools.origin_servers.health_checks / 4fd3c39ccba2 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- origin_pools.pools.origin_servers.health_checks

<a id="canonical-7d94e81d40f83db9b0dcb0a759d1c1ec2010da77acf3272709f22ae6221a11dc"></a>

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

<a id="canonical-45393d317d9cdc7ac4d3a67cf7076ce47d21030ebd436fcb3222ff6818a95cac"></a>

## Direct properties — origin_pools.pools.origin_servers.health_checks / 4fd3c39ccba2 / 3

- [health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-4c9c6d3e1722e5c73374f4f4e987686187e850077782d4a15adf5cc8910c0eb6): complete subsection reference.

<a id="canonical-1536c70a3e0aa005c6cb509df121fd3ef7912933321c006aa172273c26031577"></a>

<a id="canonical-e49ffdc31e62e9efe15c653571bcd2205cb31d1a0f43dadf07c4e69617ee57a4"></a>

## healthy_threshold property — origin_pools.pools.origin_servers.health_checks / 4fd3c39ccba2 / 4

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

<a id="canonical-ba46146150729436f985f84b3ba8c6446c7d5dcb79ad0354374aed8fa4a5d531"></a>

<a id="canonical-816af117d158c1e85ca284625b435218cf122d3c1fb2b0b53c30698b8e0fe8e5"></a>

## interval property — origin_pools.pools.origin_servers.health_checks / 4fd3c39ccba2 / 5

Type: `"number"`. Computed.

Time interval in seconds between two health check requests.

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

<a id="canonical-898be77d5767c4d3dfdee9c97d438385688ee7ff8cc8ed97c4565f656fedb4a0"></a>

<a id="canonical-983c6fc93142e1ee04ff04d6524aab135a7e2d3cda823df41408316a00a38e5c"></a>

## timeout property — origin_pools.pools.origin_servers.health_checks / 4fd3c39ccba2 / 6

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

<a id="canonical-7631fdccbef4535c4e9b501ff9cf1a0ea740316259f1cc4936be5a5172995415"></a>

<a id="canonical-d6c0ab95ab90b584faf0f99776e5076ad93b7e7733ddcfd476795914a62ba716"></a>

## unhealthy_threshold property — origin_pools.pools.origin_servers.health_checks / 4fd3c39ccba2 / 7

Type: `"number"`. Computed.

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health check
if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

Upstream description:

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health check
if a host responds with 503 this threshold is ignored and the host is considered unhealthy
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

<a id="canonical-8bf96b4ddf3042b72ae76cd96c43130965f174c046f9055b48a0449d6b758705"></a>

## Next pages — origin_pools.pools.origin_servers.health_checks / 4fd3c39ccba2 / 8

- [origin_pools.pools.origin_servers.health_checks.health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-4c9c6d3e1722e5c73374f4f4e987686187e850077782d4a15adf5cc8910c0eb6)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-4c9c6d3e1722e5c73374f4f4e987686187e850077782d4a15adf5cc8910c0eb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da5d9db93589192e91ebd614dc9d0d1488227eb3562356a73921b0287472afc9"></a>

## origin_pools.pools.origin_servers.health_checks.health_check — origin_pools.pools.origin_servers.health_checks.health_check / bdf1c3932c98 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- [origin_pools.pools.origin_servers.health_checks](data-sources--bigip_http_proxy--reference--group-001.md#canonical-5e52b07d1897a4a240c2ded9625edb68e1da365bccee4269134f3bf9e1b95b8a)
- origin_pools.pools.origin_servers.health_checks.health_check

<a id="canonical-f8e8cb7f729522e28f866dcaa5c1c36f1bd5097c9d3d520ff5c9e4f78225424f"></a>

Type: `"list"`. Computed.

List of Health Checks. List of Health Checks.

Upstream description:

List of Health Checks.

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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4fcb7ce53737713a787e5bea7fc7bd823ddb76e8e990e4c40f38ca6fcecee081"></a>

## Direct properties — origin_pools.pools.origin_servers.health_checks.health_check / bdf1c3932c98 / 3

- [icmp_health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-53f672203689c721acbcf90bb53d6a24c550ce3ed4aa687fbe37f220406b7be7): complete subsection reference.

- [tcp_health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-ce91f88ca141b2c9929da1a6d880633c2eef75130d2987d66b96a06a88593223): complete subsection reference.

<a id="canonical-55b5228ee93470bb58742930eeae063d93ff30c8edf0c07dced4ddd103dbee56"></a>

## Next pages — origin_pools.pools.origin_servers.health_checks.health_check / bdf1c3932c98 / 4

- [origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-53f672203689c721acbcf90bb53d6a24c550ce3ed4aa687fbe37f220406b7be7)
- [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-ce91f88ca141b2c9929da1a6d880633c2eef75130d2987d66b96a06a88593223)
- [origin_pools.pools.origin_servers.health_checks](data-sources--bigip_http_proxy--reference--group-001.md#canonical-5e52b07d1897a4a240c2ded9625edb68e1da365bccee4269134f3bf9e1b95b8a)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-53f672203689c721acbcf90bb53d6a24c550ce3ed4aa687fbe37f220406b7be7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5299199c652bdc52af763ae00bae8587f0b5a2b9ccc610fb46d1d57bffe4f214"></a>

## origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check — origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check / 6e4498f8717b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- [origin_pools.pools.origin_servers.health_checks](data-sources--bigip_http_proxy--reference--group-001.md#canonical-5e52b07d1897a4a240c2ded9625edb68e1da365bccee4269134f3bf9e1b95b8a)
- [origin_pools.pools.origin_servers.health_checks.health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-4c9c6d3e1722e5c73374f4f4e987686187e850077782d4a15adf5cc8910c0eb6)
- origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check

<a id="canonical-4608ba84d2ef94274309d51c31645c93af3b1b1b425b5ef30c929df5d680ce4c"></a>

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

<a id="canonical-7ea5218da312abd9d2c9d68fd7ce29b7834b4330e186fc0274a84597ad832695"></a>

## Direct properties — origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check / 6e4498f8717b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d8ea8ca07b2c766775da634e05db7e77dc0528be0cd91385e5899d92e7e963ad"></a>

## Next pages — origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check / 6e4498f8717b / 4

- [origin_pools.pools.origin_servers.health_checks.health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-4c9c6d3e1722e5c73374f4f4e987686187e850077782d4a15adf5cc8910c0eb6)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-ce91f88ca141b2c9929da1a6d880633c2eef75130d2987d66b96a06a88593223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc5ad516f05ac09dbec48298f743222a4508c08ae1e0f273018397fc0e24885c"></a>

## origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check — origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check / b64fb377b55f / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- [origin_pools.pools.origin_servers.health_checks](data-sources--bigip_http_proxy--reference--group-001.md#canonical-5e52b07d1897a4a240c2ded9625edb68e1da365bccee4269134f3bf9e1b95b8a)
- [origin_pools.pools.origin_servers.health_checks.health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-4c9c6d3e1722e5c73374f4f4e987686187e850077782d4a15adf5cc8910c0eb6)
- origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check

<a id="canonical-213aeea9391e73c994a272adb4b2f12e16956cae5771e20e41355a072c7572ff"></a>

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

<a id="canonical-446f2adcc086cc52fa09cb2100322ec8ba7e0f271816251c377c093ec579aabd"></a>

## Direct properties — origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check / b64fb377b55f / 3

<a id="canonical-ccdb4f197002de3071e084dde3bf55b4f6633180f84efd9479d2ba64bb7b1784"></a>

<a id="canonical-30be47531fe927b9e0a4aae9e87a747bea69d70af59386ce28d9d3e37bdd6ec5"></a>

## expected_response property — origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check / b64fb377b55f / 4

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

<a id="canonical-024fa6ed46b3a523ecb7ba0083e9429ffd4650cff3d3f90a1b627dc3f140a1bc"></a>

<a id="canonical-124fa022471158d60d73812ade2a1a4461baf72aa23c9830c111ae817f565922"></a>

## send_payload property — origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check / b64fb377b55f / 5

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

<a id="canonical-3775222e84b9d5232950f7e57e7b93429bce10cbf2172fd45d3ef7e32b295b49"></a>

## Next pages — origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check / b64fb377b55f / 6

- [origin_pools.pools.origin_servers.health_checks.health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-4c9c6d3e1722e5c73374f4f4e987686187e850077782d4a15adf5cc8910c0eb6)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-9f7280eca68fe7a0a5c065150caafe90d41f06f99ddfe7b6b6280aa7f14550fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-742f657de5a5ef1aab2d0feeb5994f3e08ec83cdd5e50d240abebe12bb7ac8f1"></a>

## origin_pools.pools.origin_servers.lb_port — origin_pools.pools.origin_servers.lb_port / 149968fd1de3 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- origin_pools.pools.origin_servers.lb_port

<a id="canonical-abcf26e9312bc007e2ba92c8f851799879404388a665c0ea2a36bafbcf5f3c3b"></a>

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

<a id="canonical-cf99c93a1f82de4dde2dbed550284012720b60c1f14b96649bb9d03bb49e7037"></a>

## Direct properties — origin_pools.pools.origin_servers.lb_port / 149968fd1de3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f1cb6d31c5c9043f8bfb859766e58eeb5a483059ae399ed5411bf62b0e49042f"></a>

## Next pages — origin_pools.pools.origin_servers.lb_port / 149968fd1de3 / 4

- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-355d4acc28374bc67f053c73c5fc1c3e216af67fb65258137fc871ddb6b9a11c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f86359dd5868c32b601e2d97c3271a45d8b5c79f4c26731166b6a4a82c1e0df5"></a>

## origin_pools.pools.origin_servers.origin_servers — origin_pools.pools.origin_servers.origin_servers / 439e5be5b940 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- origin_pools.pools.origin_servers.origin_servers

<a id="canonical-fb5279501fba94f37f6a64b2522665e6ac1878440b62c6d8c8f7bb9dd024593c"></a>

Type: `"list"`. Computed.

List of Origin Servers. List of origin servers for Proxy.

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

<a id="canonical-5cfd3b53021102f3134be32c2158c1706d7eece2b0f13fb11db4d981e37f3246"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers / 439e5be5b940 / 3

- [k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-ad8c69057a368a4754b2eacaead7fd11ee390da5de2907a9f7f44cd0d6ad1a5f): complete subsection reference.

- [private_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-832942dd886c73b7346edac16bea596436b1d31f35d284df7c7cc9b59b7983d4): complete subsection reference.

- [public_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-31c9303db90064228ff9ffba7eaeae620e85940732abf3a9f5085b3f6f77383d): complete subsection reference.

- [public_name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-ddf8319ce01f461c3078058268f5cd0be7ea6b998d76f53760f4552754ee191d): complete subsection reference.

<a id="canonical-e356327972b9c3b6efcbe33f5d3e62d41dff0c2289a91fd74c45341a653877f0"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers / 439e5be5b940 / 4

- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-ad8c69057a368a4754b2eacaead7fd11ee390da5de2907a9f7f44cd0d6ad1a5f)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-832942dd886c73b7346edac16bea596436b1d31f35d284df7c7cc9b59b7983d4)
- [origin_pools.pools.origin_servers.origin_servers.public_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-31c9303db90064228ff9ffba7eaeae620e85940732abf3a9f5085b3f6f77383d)
- [origin_pools.pools.origin_servers.origin_servers.public_name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-ddf8319ce01f461c3078058268f5cd0be7ea6b998d76f53760f4552754ee191d)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-ad8c69057a368a4754b2eacaead7fd11ee390da5de2907a9f7f44cd0d6ad1a5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a27ed62127d1157cd725373bf4811e7a2922ced35c92fefa05f9190a96f7aadd"></a>

## origin_pools.pools.origin_servers.origin_servers.k8s_service — origin_pools.pools.origin_servers.origin_servers.k8s_service / e0aa1f29835e / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-355d4acc28374bc67f053c73c5fc1c3e216af67fb65258137fc871ddb6b9a11c)
- origin_pools.pools.origin_servers.origin_servers.k8s_service

<a id="canonical-0d13de34f9eed505a093fb50be90c5b504ab6bd31299da3739cdf3f4882db888"></a>

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

<a id="canonical-4cde554635ee081392aa32a999eedf5699391c0f273d6109331db37f72e78a53"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.k8s_service / e0aa1f29835e / 3

- [inside_network](data-sources--bigip_http_proxy--reference--group-001.md#canonical-e902044b97461ee64a3c24d4e7e7a07b59bc81c522561fab86f40ed637b436b6): complete subsection reference.

- [outside_network](data-sources--bigip_http_proxy--reference--group-001.md#canonical-6600d3e7c05310664b4e61a1a5a3e321e5586bfbb7d75029b3f7d937f036d65c): complete subsection reference.

<a id="canonical-f3e096629bcee808a04821b2673eeb2eee894a7c6fd81ba49d707bf0993d34e5"></a>

<a id="canonical-4c51d2ddcfa8720b820eeaf6ddaaee71e6d56491183ec945831df5de582a21ae"></a>

## protocol property — origin_pools.pools.origin_servers.origin_servers.k8s_service / e0aa1f29835e / 4

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

<a id="canonical-e463cf99ce14bb2229f504df45d8a68929ab085a1fd2f81dfa99ad6f2ec08d21"></a>

<a id="canonical-07a8944364cc40cf5824b6a1c2db31c8b9589a7c7f6bd1133525229a9f8e3888"></a>

## service_name property — origin_pools.pools.origin_servers.origin_servers.k8s_service / e0aa1f29835e / 5

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

- [site_locator](data-sources--bigip_http_proxy--reference--group-001.md#canonical-c704877228234cc1f9475310717e9c65666d09e95cb9bd8410f94c7bf9cb247f): complete subsection reference.

- [snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-86d788efda4db0a8653d64636e9d6019a25d2f70ef1a3652b85d59d22a41eae6): complete subsection reference.

- [vk8s_networks](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3ff8a6b38c12505a95eba50a5c58604b8ccda77692052832d96fc87e554cbad5): complete subsection reference.

<a id="canonical-74b5702898466ca456bbb4c111f705d3fcea0e23582515d8ef6e78a57f504fcb"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.k8s_service / e0aa1f29835e / 6

- [origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network](data-sources--bigip_http_proxy--reference--group-001.md#canonical-e902044b97461ee64a3c24d4e7e7a07b59bc81c522561fab86f40ed637b436b6)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network](data-sources--bigip_http_proxy--reference--group-001.md#canonical-6600d3e7c05310664b4e61a1a5a3e321e5586bfbb7d75029b3f7d937f036d65c)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](data-sources--bigip_http_proxy--reference--group-001.md#canonical-c704877228234cc1f9475310717e9c65666d09e95cb9bd8410f94c7bf9cb247f)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-86d788efda4db0a8653d64636e9d6019a25d2f70ef1a3652b85d59d22a41eae6)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3ff8a6b38c12505a95eba50a5c58604b8ccda77692052832d96fc87e554cbad5)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-355d4acc28374bc67f053c73c5fc1c3e216af67fb65258137fc871ddb6b9a11c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-e902044b97461ee64a3c24d4e7e7a07b59bc81c522561fab86f40ed637b436b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fdacc1345e3f182498ef32d3b4e24b4fae6cf9df6ed01b201406316a33b6eda0"></a>

## origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network — origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network / 21a142ca0918 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-355d4acc28374bc67f053c73c5fc1c3e216af67fb65258137fc871ddb6b9a11c)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-ad8c69057a368a4754b2eacaead7fd11ee390da5de2907a9f7f44cd0d6ad1a5f)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network

<a id="canonical-7009de49a95e5495824730ce2d490d421e414c3024dc4ad43effb3277ad29c1a"></a>

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

<a id="canonical-54b29a1ac7f3b0ee616c3f2a08893a907dfb86c695ac9ec5aaeca389233ba7d1"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network / 21a142ca0918 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4ba237ef138913f58b40d1f9dd8abedf6edc28a5905829fd6f6a43470233c4b4"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network / 21a142ca0918 / 4

- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-ad8c69057a368a4754b2eacaead7fd11ee390da5de2907a9f7f44cd0d6ad1a5f)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-6600d3e7c05310664b4e61a1a5a3e321e5586bfbb7d75029b3f7d937f036d65c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04e37652983e5ffa4d6f5022da07f61283a0fb34b0bf2ffef3824759a2d59630"></a>

## origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network — origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network / 9e2a7d4f4bcb / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2df954e5860f42fdbeec446e3c7f45668490100aef3c7068dd9998870397582e)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-f7dcdb0ab6ff2013ed4d8f2610d3b5f7ebe69f9b2143451bc9c15ed42565a30a)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-917c2f01bcc2d35a113d73418c3aa09d8928082ea3f392d72dd36248378aa00d)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-355d4acc28374bc67f053c73c5fc1c3e216af67fb65258137fc871ddb6b9a11c)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-ad8c69057a368a4754b2eacaead7fd11ee390da5de2907a9f7f44cd0d6ad1a5f)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network

<a id="canonical-29b6a0f5e0ccc800cd3d787a35ac0c4ffcd4779ba1827f28abeadd02b8867cb6"></a>

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

<a id="canonical-336ecf556d7d3201eb82acc4a37d3ae4ab73aac0ff5f20702074f7cd078c19a4"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network / 9e2a7d4f4bcb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e63f86d34eafa9ec5d16741e777ec3a420fe2d2ac3c4133705b33cd155fcab74"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network / 9e2a7d4f4bcb / 4

- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-ad8c69057a368a4754b2eacaead7fd11ee390da5de2907a9f7f44cd0d6ad1a5f)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-c704877228234cc1f9475310717e9c65666d09e95cb9bd8410f94c7bf9cb247f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
