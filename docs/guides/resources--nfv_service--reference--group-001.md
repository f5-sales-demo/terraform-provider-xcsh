---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5aaf99c87a1792bc1806a721e6531db912c57aab90fa0604cd56f5a473ef8f16"></a>

## Property reference — Property reference / b32dfc5e9333 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- Property reference

<a id="canonical-2409fa443cfd7e848f271b37ae38376571283b74e0203bdfbf30540611d5cb9d"></a>

## Direct properties — Property reference / b32dfc5e9333 / 3

<a id="canonical-b2978896b781b87a3fff32e13e855d7b07afda2a419146de63cb5c83b536a6c8"></a>

<a id="canonical-eb345d8c04b067f79c4bf6173dc9ca0b0e12d0ad4f5e85cbb41690c0aed54be8"></a>

## annotations property — Property reference / b32dfc5e9333 / 4

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

<a id="canonical-2e49e4df5ed1e262c539846cef3734f510571ef1658894853768405bc6bcd583"></a>

<a id="canonical-be01d9c2611bc38645bc8edc87e023b4c20acb2bf1c34b2b63df4ea3e4c5f17b"></a>

## description property — Property reference / b32dfc5e9333 / 5

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

<a id="canonical-28180a749ea4a11833ac3675b8df865a7195b2c1eef903bd328196dc41f5a862"></a>

<a id="canonical-d04b65e7c199789084e18511e72848fbfe17eeb34ed6e99e1f40c6428e0b8524"></a>

## disable property — Property reference / b32dfc5e9333 / 6

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

- [disable_https_management](resources--nfv_service--reference--group-001.md#canonical-6725ac95b37b2f5e8f60a7b62e92fcfb81391da7dd183f8b11259ecd909c2bd0): complete subsection reference.

- [disable_ssh_access](resources--nfv_service--reference--group-001.md#canonical-59b17828f1392d7b77df4c061dc7266ffbbf235c736ab86ffe900270e0b5c4ce): complete subsection reference.

- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-6fef4a6194444c7af1b6d53343496678464efb8500bbd53ba2714a605fd63804): complete subsection reference.

- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4): complete subsection reference.

- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51): complete subsection reference.

<a id="canonical-a2903b099bf7ca67a81ed51a2fdd1bb69ec44bd5fda68d42d0f20191ecd0395c"></a>

<a id="canonical-b25112e706492bbae25196816880e0fa26b23a5e8afdf57b47b35e0bf39da500"></a>

## id property — Property reference / b32dfc5e9333 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-db8f6c88495cf9e9d44072956238517d62121d94c5bc73205248fc391cd71024"></a>

<a id="canonical-f870a80b97c672fb9e70b5d85a0eea4d3615d2d61db35de0e7f3bdcae98642e2"></a>

## labels property — Property reference / b32dfc5e9333 / 8

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

<a id="canonical-17d8743bc18fcde39ca586e09fbad1ff9d1eab6b01b3faea5707abe8f59ed6ec"></a>

<a id="canonical-717aec5ff39c08a0aa9536813afafdffcb1575335aeff46c336e78c8f7d0f11a"></a>

## name property — Property reference / b32dfc5e9333 / 9

Type: `"string"`. Required.

Name of the Nfv Service. Must be unique within the namespace.

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

<a id="canonical-9e97ed29393906a10dcc2b427f50a8b48c9c1ab30132a543ee63790db0c1baca"></a>

<a id="canonical-c520c351ee8812a47aea719da09d29f1689f6105a4cd5cbcd5af865cb100a8d6"></a>

## namespace property — Property reference / b32dfc5e9333 / 10

Type: `"string"`. Required.

Namespace where the Nfv Service is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf): complete subsection reference.

- [timeouts](resources--nfv_service--reference--group-004.md#canonical-22b89a5a007d5d26888d6c80645783ec0135cd1ced1c3d4d29ec4c3cdc0ebce1): complete subsection reference.

<a id="canonical-b03e9622cfcb4ddb3a744a00a9a60aa66f3752b56b8d24d15759e2471667bd83"></a>

## All schema paths — Property reference / b32dfc5e9333 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--nfv_service--reference--group-001.md#canonical-b2978896b781b87a3fff32e13e855d7b07afda2a419146de63cb5c83b536a6c8) |
| `description` | [description](resources--nfv_service--reference--group-001.md#canonical-2e49e4df5ed1e262c539846cef3734f510571ef1658894853768405bc6bcd583) |
| `disable` | [disable](resources--nfv_service--reference--group-001.md#canonical-28180a749ea4a11833ac3675b8df865a7195b2c1eef903bd328196dc41f5a862) |
| `disable_https_management` | [disable_https_management](resources--nfv_service--reference--group-001.md#canonical-d6d9c1c4276b1dba6fb22647aeddc98c0a711c4f114d4b1f9cc65952716eb654) |
| `disable_ssh_access` | [disable_ssh_access](resources--nfv_service--reference--group-001.md#canonical-1577f6dcb281d05115836289f252e63d045d9ac1a5aa5e5c88e9a1fba796d15b) |
| `enabled_ssh_access` | [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-21ca4809441746237ba51aa036221a5e06bd4b462d112224ce9efc59fbb1ca93) |
| `enabled_ssh_access.advertise_on_sli` | [enabled_ssh_access.advertise_on_sli](resources--nfv_service--reference--group-001.md#canonical-031d045d5a861b2a41a07323d365f48e8726b17104e129dc6c37ec6515ecd84e) |
| `enabled_ssh_access.advertise_on_slo` | [enabled_ssh_access.advertise_on_slo](resources--nfv_service--reference--group-001.md#canonical-d730517363dbe5215b6ffdc2a8284e56de2a4bb1bad6e55618f6103332673314) |
| `enabled_ssh_access.advertise_on_slo_sli` | [enabled_ssh_access.advertise_on_slo_sli](resources--nfv_service--reference--group-001.md#canonical-b16d6e3c7f0a6ee10c1f5be35079140283e75c0554838cae2d7226036dbcac2e) |
| `enabled_ssh_access.domain_suffix` | [enabled_ssh_access.domain_suffix](resources--nfv_service--reference--group-001.md#canonical-88877cf04488f92b9160c2fe4bb32bbbf0b676af405c954815be89e9a44bf1b5) |
| `enabled_ssh_access.node_ssh_ports` | [enabled_ssh_access.node_ssh_ports](resources--nfv_service--reference--group-001.md#canonical-dd9b73e94db883cea6541a526fefc5c3e1ea35c6826124786c6ed6432a89fd53) |
| `enabled_ssh_access.node_ssh_ports.node_name` | [enabled_ssh_access.node_ssh_ports.node_name](resources--nfv_service--reference--group-001.md#canonical-352fcaabadda12fd7449f41afb05ace4a984317107147ec13d713b6172d70528) |
| `enabled_ssh_access.node_ssh_ports.ssh_port` | [enabled_ssh_access.node_ssh_ports.ssh_port](resources--nfv_service--reference--group-001.md#canonical-e7250fd238e29cc88540125cfa00a2ef13c2868d69951175353242565ce8c5f8) |
| `f5_big_ip_aws_service` | [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2142caf3816f79bc0b5a8dc3fdd65f9222695b01ae57af4744368f7ffdbceeb9) |
| `f5_big_ip_aws_service.admin_password` | [f5_big_ip_aws_service.admin_password](resources--nfv_service--reference--group-001.md#canonical-08fa7f6691a357f79a9846f12f4c6a5ae6484f9cd718f0ae9541b6344b5a3cd6) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info](resources--nfv_service--reference--group-001.md#canonical-f12ac6666dc171212d4f8284a74bc05dd88ca8b1062eee71fddc9f9ce870b230) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.decryption_provider` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-001.md#canonical-b9c3e4ccbb56ec57a20e80d9326566921777ae110a4f6a3f6c5edc9f288f053c) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.location` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.location](resources--nfv_service--reference--group-001.md#canonical-f5cee47bdce04278fe1411a61d03b8076840d291ac108717c38e38a8b8b38ec1) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.store_provider` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-001.md#canonical-b44f6ab2de7e1f53796f00a1e6df8667cb2d596861a2c529f3bdafdcc043a05d) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info` | [f5_big_ip_aws_service.admin_password.clear_secret_info](resources--nfv_service--reference--group-001.md#canonical-f7466a1404176aa1ae4d2537bce7ae981bfee603126afda95f8e1624e37a3672) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info.provider_ref` | [f5_big_ip_aws_service.admin_password.clear_secret_info.provider_ref](resources--nfv_service--reference--group-001.md#canonical-bb8a031940fdcf901a33dbd85deeace8fae5c13ad2fabc8cde163df6738a8b61) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info.url` | [f5_big_ip_aws_service.admin_password.clear_secret_info.url](resources--nfv_service--reference--group-001.md#canonical-4f9ff36ccbe1ca15240af9071e2e4dbe9d353aa81f2ba709079b221805fc1530) |
| `f5_big_ip_aws_service.admin_username` | [f5_big_ip_aws_service.admin_username](resources--nfv_service--reference--group-001.md#canonical-33ce4b0704b55befc01cb2d1569d2b02b33f81b5f8f6b969c9c22e4bc6669aa8) |
| `f5_big_ip_aws_service.aws_tgw_site_params` | [f5_big_ip_aws_service.aws_tgw_site_params](resources--nfv_service--reference--group-001.md#canonical-f8dd381e131179115134244684f4b289bf93554a0c357026c16c7f2dbd36347f) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site](resources--nfv_service--reference--group-001.md#canonical-03399a97055910b2b1ce7fb514a52ecf4f294f796ef9f94857db9f6496450228) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.name` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.name](resources--nfv_service--reference--group-001.md#canonical-44f67026a93a2850486be21a79d99d09eb74d7488b60a2dcebc78ed30c3245b6) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.namespace` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.namespace](resources--nfv_service--reference--group-001.md#canonical-35acce15bda479152a832a1266dd1ca9049adcd01a937a288ac685d8493b87e6) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.tenant` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.tenant](resources--nfv_service--reference--group-001.md#canonical-9cfc4db49ab7d2c763100b88c85fdd38f204201c7d7a5553cfe6b7f69577eaaa) |
| `f5_big_ip_aws_service.endpoint_service` | [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-4ee6168cd62d41c8989922d8712c7d377b7718e86dd78a73e345093d82ecb417) |
| `f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip` | [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip](resources--nfv_service--reference--group-001.md#canonical-89208ab8c6ba8968e2fc4959cd55699e86567a7196e3e445ad9d654908a4f0c3) |
| `f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external` | [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external](resources--nfv_service--reference--group-001.md#canonical-f4e65b9f6bd44d3243d0a8ca7e0b08f2fcb9076b9c7a1a31e38bbb239b4a47ee) |
| `f5_big_ip_aws_service.endpoint_service.automatic_vip` | [f5_big_ip_aws_service.endpoint_service.automatic_vip](resources--nfv_service--reference--group-001.md#canonical-3697b50ff3fa370d2f805177355f0a925e107688ed2b882e080c9e76f1b76795) |
| `f5_big_ip_aws_service.endpoint_service.configured_vip` | [f5_big_ip_aws_service.endpoint_service.configured_vip](resources--nfv_service--reference--group-001.md#canonical-82d18cdb1aad812c7dcca9dc550071f56d5ba05bb0a00fd545f5c3527bc75beb) |
| `f5_big_ip_aws_service.endpoint_service.custom_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports](resources--nfv_service--reference--group-001.md#canonical-681f58eafa7a57e8eabe9fd3da7b0c67125eb3ecb3f1d891aadbea94f6f65a62) |
| `f5_big_ip_aws_service.endpoint_service.custom_tcp_ports.ports` | [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports.ports](resources--nfv_service--reference--group-001.md#canonical-006120b2a17c6bdb6ca1ac726c0cd1f8152f482b5183ed3f8d890776c05ef066) |
| `f5_big_ip_aws_service.endpoint_service.custom_udp_ports` | [f5_big_ip_aws_service.endpoint_service.custom_udp_ports](resources--nfv_service--reference--group-001.md#canonical-b0fffef94c27f48543349cc860edd0dac5ee8d138dae156e7c507f9fd8f6a173) |
| `f5_big_ip_aws_service.endpoint_service.custom_udp_ports.ports` | [f5_big_ip_aws_service.endpoint_service.custom_udp_ports.ports](resources--nfv_service--reference--group-001.md#canonical-46251af9ebd0a00ace50606cbfb6e25afd9f27d2b2d62370274f3ded7c2d0a8e) |
| `f5_big_ip_aws_service.endpoint_service.default_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.default_tcp_ports](resources--nfv_service--reference--group-001.md#canonical-8c2e4d9781ef667f06a0f82bbf70b502ee0576a95907ac428cd673f7299f1ca8) |
| `f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip` | [f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip](resources--nfv_service--reference--group-001.md#canonical-d2ad5e9d5c6702de80eae4395d6d79d932fc91c2062dc7cbf9a657f1d160d22c) |
| `f5_big_ip_aws_service.endpoint_service.http_port` | [f5_big_ip_aws_service.endpoint_service.http_port](resources--nfv_service--reference--group-001.md#canonical-1a1270bf9d72506f82d42b2bfa9a191f82194b714b440b6bbd7724a22c291983) |
| `f5_big_ip_aws_service.endpoint_service.https_port` | [f5_big_ip_aws_service.endpoint_service.https_port](resources--nfv_service--reference--group-001.md#canonical-e7802b3e5600fab499a0850d37fdc9ea78745ae1572870168ea4ea10baecd41e) |
| `f5_big_ip_aws_service.endpoint_service.no_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.no_tcp_ports](resources--nfv_service--reference--group-001.md#canonical-0d8d7df88027ebd71f68e4c209688575de358e698d7ab4c637f86566ce6a5c06) |
| `f5_big_ip_aws_service.endpoint_service.no_udp_ports` | [f5_big_ip_aws_service.endpoint_service.no_udp_ports](resources--nfv_service--reference--group-001.md#canonical-b257df1ad829e4a78623d78dd1dcecf9f090059a786e6de41ea4ed8fc628c91d) |
| `f5_big_ip_aws_service.market_place_image` | [f5_big_ip_aws_service.market_place_image](resources--nfv_service--reference--group-002.md#canonical-1cc63432f816a5d3af8b64ef59998f89108524fa5e92abf4305e4599f05903d1) |
| `f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps` | [f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps](resources--nfv_service--reference--group-002.md#canonical-c7084c8d2b646cb723d4f92d063559267e61c74af78144ae56035468784bf987) |
| `f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps` | [f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps](resources--nfv_service--reference--group-002.md#canonical-374ef557714400b204f2f477f3379f0a1118e2dd1396a8e5544f3201f5673045) |
| `f5_big_ip_aws_service.nodes` | [f5_big_ip_aws_service.nodes](resources--nfv_service--reference--group-002.md#canonical-fadbd30abbf68d6cfa975fcd4435fe90915961020e3673a5d9ede792c007691f) |
| `f5_big_ip_aws_service.nodes.automatic_prefix` | [f5_big_ip_aws_service.nodes.automatic_prefix](resources--nfv_service--reference--group-002.md#canonical-39aa4c2677e729f4e674b22177da306e201450bc4bd4b2c1082cb2b2cc37ef7b) |
| `f5_big_ip_aws_service.nodes.aws_az_name` | [f5_big_ip_aws_service.nodes.aws_az_name](resources--nfv_service--reference--group-002.md#canonical-9cd8d4056bf96870ae54d1191208c86ccaf68c42eca01d8a3f43123539be34ea) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet` | [f5_big_ip_aws_service.nodes.mgmt_subnet](resources--nfv_service--reference--group-002.md#canonical-1a08c5015a18a055dda2b28cedf19408fbeb3a12c0f860ec490c26a6a99c5b93) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.existing_subnet_id` | [f5_big_ip_aws_service.nodes.mgmt_subnet.existing_subnet_id](resources--nfv_service--reference--group-002.md#canonical-e545c95d05f4deb9522d42fdb8079cc47ac407d2eaa14207a230e757b8d4916c) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param` | [f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param](resources--nfv_service--reference--group-002.md#canonical-17d97287c39533e97427b2f5a0d0b0eb8642638796ad32a82e3630aa63f4a311) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param.ipv4` | [f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param.ipv4](resources--nfv_service--reference--group-002.md#canonical-3d92be3a4f0c0a1254ff29326cba629f7285a67616f3fd31bb63a09d53a97592) |
| `f5_big_ip_aws_service.nodes.node_name` | [f5_big_ip_aws_service.nodes.node_name](resources--nfv_service--reference--group-002.md#canonical-cb308a3a26a0952be331e62c11e5219e47d8fab9c4e12c582fde65a67bbdb5fb) |
| `f5_big_ip_aws_service.nodes.reserved_mgmt_subnet` | [f5_big_ip_aws_service.nodes.reserved_mgmt_subnet](resources--nfv_service--reference--group-002.md#canonical-35ede095a92e65d39dd73627d9cbd5cbff403eac877d5e5f1e742f7f44d8d4f9) |
| `f5_big_ip_aws_service.nodes.tunnel_prefix` | [f5_big_ip_aws_service.nodes.tunnel_prefix](resources--nfv_service--reference--group-002.md#canonical-4395c9ede84d44a752ecccf5a89c6c69ac57304dcc6172799c8fc2ff7f9b5c98) |
| `f5_big_ip_aws_service.ssh_key` | [f5_big_ip_aws_service.ssh_key](resources--nfv_service--reference--group-001.md#canonical-9ae902042fb560fa50f0798c4dd202800bc16c99ec7ed680c9208cd3c2dfb405) |
| `f5_big_ip_aws_service.tags` | [f5_big_ip_aws_service.tags](resources--nfv_service--reference--group-001.md#canonical-b06a7adcba66eb225949433b2f299bb9d8f06a97a1ceb97ca087a0c47ab946c9) |
| `https_management` | [https_management](resources--nfv_service--reference--group-002.md#canonical-b4518939056897104846d93f414610b8e7ef75ac36012d24571976aa6b8b425c) |
| `https_management.advertise_on_internet` | [https_management.advertise_on_internet](resources--nfv_service--reference--group-002.md#canonical-3b10091445e68467f2dc70d0d1e36723c851e08499c661d3d62687bc350f0b60) |
| `https_management.advertise_on_internet.public_ip` | [https_management.advertise_on_internet.public_ip](resources--nfv_service--reference--group-002.md#canonical-782b2e3131ed654806a4e98f59de9d0dd592a2d2656a083326ebfc69a0d16308) |
| `https_management.advertise_on_internet.public_ip.name` | [https_management.advertise_on_internet.public_ip.name](resources--nfv_service--reference--group-002.md#canonical-9295477d0db70f40eb79fca519d6f150bf2d25f30b8899112393a660c834547b) |
| `https_management.advertise_on_internet.public_ip.namespace` | [https_management.advertise_on_internet.public_ip.namespace](resources--nfv_service--reference--group-002.md#canonical-d31aa133f06800231f7f56705f79df54f4d0298bbc715f06a909fdeb81168a02) |
| `https_management.advertise_on_internet.public_ip.tenant` | [https_management.advertise_on_internet.public_ip.tenant](resources--nfv_service--reference--group-002.md#canonical-c568dfdb0491536c6043d483fb49105362515c562828666e4d4d2a78feb0af29) |
| `https_management.advertise_on_internet_default_vip` | [https_management.advertise_on_internet_default_vip](resources--nfv_service--reference--group-002.md#canonical-ceca158e74c183b3d2d2204bde0a31fe379637faac00de2a0c03cd4ffc696ae5) |
| `https_management.advertise_on_sli_vip` | [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3f882978171e511f48b1ba30574f7529eec5107bc5af8911be3b8dcdb064387f) |
| `https_management.advertise_on_sli_vip.no_mtls` | [https_management.advertise_on_sli_vip.no_mtls](resources--nfv_service--reference--group-002.md#canonical-6ce6f81f6b020ffd3327a917e9f83d3bc2dc5061ae8466930bb3afa05fc9f415) |
| `https_management.advertise_on_sli_vip.tls_certificates` | [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-3d9f35642287f71fe43f45f66446a74e75d105530fcc7d6bea80594f54ae23ff) |
| `https_management.advertise_on_sli_vip.tls_certificates.certificate_url` | [https_management.advertise_on_sli_vip.tls_certificates.certificate_url](resources--nfv_service--reference--group-002.md#canonical-b150106cdf741889867f6845e2dc440db97e1fcd4ec3a4d8ade2f748533ecd72) |
| `https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms](resources--nfv_service--reference--group-002.md#canonical-2ea395e648a2473491d7a16060a837c93bf975d8915fd2bf6ffc6d865d77403d) |
| `https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--nfv_service--reference--group-002.md#canonical-12357436615cb4dd496eb0494a567f399e2f238cd4d50784da0450b8111331e0) |
| `https_management.advertise_on_sli_vip.tls_certificates.description_spec` | [https_management.advertise_on_sli_vip.tls_certificates.description_spec](resources--nfv_service--reference--group-002.md#canonical-7336ccbb4a1779c5a2a775f85d1d8316590b89e9239a776740b5702914e861fc) |
| `https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling](resources--nfv_service--reference--group-002.md#canonical-6b30ffcba6c173e586c221fe002c997a40af5be1a2d42c58f653b62317a3b193) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key` | [https_management.advertise_on_sli_vip.tls_certificates.private_key](resources--nfv_service--reference--group-002.md#canonical-f2e2876452be96b1fb2573cbb77e4831ac9ca1a66538b25a22b207925f3171a8) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--reference--group-002.md#canonical-7e8c2d36965687a7d7eb3a256515d13bbdeae78648362b537ff692ea6f9b1066) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-002.md#canonical-f390615e467cc8fce1bac5b84f1d17b8b67067d9731a1bb4be9fc1a3b6076f16) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.location](resources--nfv_service--reference--group-002.md#canonical-c12d63df38050a701274610d00170d53360a15eb37c57c10778ae9c2d38a3afa) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-002.md#canonical-9f953212973c1959001b3ef5a900a74dda57527c47588e5f3a82bd120738f754) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--reference--group-002.md#canonical-5d039e5829dc1747f7c1eafcaaba4ee03e8974a3a0b68bfe6fef416021f82cfb) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.provider_ref](resources--nfv_service--reference--group-002.md#canonical-5b657f86dc67feb541d822ea48622d8ab3243dfccc4c041e9fe1db4c746b85a6) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.url](resources--nfv_service--reference--group-002.md#canonical-f2e185e9d3da5701b2f0ac0d5ad531697a410901aec02590614a54cf2c43c94a) |
| `https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults](resources--nfv_service--reference--group-002.md#canonical-21d71019f6509ab7713e77c8c4c61f9e0fa7b103675d0667c1fccf0355a4d322) |
| `https_management.advertise_on_sli_vip.tls_config` | [https_management.advertise_on_sli_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-d1007795e2ed398441bfd1f3ec0ba65ff1d216e59200c07cd0077b6003edcebe) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security` | [https_management.advertise_on_sli_vip.tls_config.custom_security](resources--nfv_service--reference--group-002.md#canonical-396d63c6998b4786856bf7f9a498f6218258897995d9544f58edc2962f8318a7) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_sli_vip.tls_config.custom_security.cipher_suites](resources--nfv_service--reference--group-002.md#canonical-cad1fdd4574b8f4da20fd4aafac5027bf0c39a6a43d1a3a3ca7c5a8bc9094544) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_sli_vip.tls_config.custom_security.max_version](resources--nfv_service--reference--group-002.md#canonical-4db9a5d704a682aaaa7d9c23b8fa2a433c5e06451d95d1806b04e4f675556474) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_sli_vip.tls_config.custom_security.min_version](resources--nfv_service--reference--group-002.md#canonical-fd011b44902f2d9b6ea9695728dc8c77ed56945fba86e9723eaed89efffad02e) |
| `https_management.advertise_on_sli_vip.tls_config.default_security` | [https_management.advertise_on_sli_vip.tls_config.default_security](resources--nfv_service--reference--group-002.md#canonical-c2498e069c5e5f1b5e6a1e0a9936bd3c76c7a133287c1fa94ec10ca2de6c1892) |
| `https_management.advertise_on_sli_vip.tls_config.low_security` | [https_management.advertise_on_sli_vip.tls_config.low_security](resources--nfv_service--reference--group-002.md#canonical-b0713211203b5df2b384800d6ea3b65b4cbd69cb6393b76e24416f0a8eec9c3c) |
| `https_management.advertise_on_sli_vip.tls_config.medium_security` | [https_management.advertise_on_sli_vip.tls_config.medium_security](resources--nfv_service--reference--group-002.md#canonical-469ace6133b155baf459cd23a6f7672bccc574852c240974e32a9e2e4272296d) |
| `https_management.advertise_on_sli_vip.use_mtls` | [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-b0514d0ec1db1c0428aeb098a8d4395a219dbe44861f5a29dc90ab96e828a08c) |
| `https_management.advertise_on_sli_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_sli_vip.use_mtls.client_certificate_optional](resources--nfv_service--reference--group-002.md#canonical-7a2fb951508c67e93bf4d2717cf6d140539d8173f25dc9fe1fcbefc0a4a0fd43) |
| `https_management.advertise_on_sli_vip.use_mtls.crl` | [https_management.advertise_on_sli_vip.use_mtls.crl](resources--nfv_service--reference--group-002.md#canonical-48f539fd2f28a8cc222b6e4c5fdf8542601d3e191cb733da4c67b523a2e62df0) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.name` | [https_management.advertise_on_sli_vip.use_mtls.crl.name](resources--nfv_service--reference--group-002.md#canonical-e7113fedeee1fd743eb1126a52111bb84b1d4052311b2e3cec3250be0bb36bac) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.namespace` | [https_management.advertise_on_sli_vip.use_mtls.crl.namespace](resources--nfv_service--reference--group-002.md#canonical-a7c8c2755651e6eb088860076630c1c15084c6f3c7327bfb6ef76beb663e77c1) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.tenant` | [https_management.advertise_on_sli_vip.use_mtls.crl.tenant](resources--nfv_service--reference--group-002.md#canonical-785085bc37a9bd46be707e2cfea9f4de00b67658c3396865585109adb366132a) |
| `https_management.advertise_on_sli_vip.use_mtls.no_crl` | [https_management.advertise_on_sli_vip.use_mtls.no_crl](resources--nfv_service--reference--group-002.md#canonical-97bd1ab411dca56becfac6631a7783de0226668dacb72c805da23706c729fe0c) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca](resources--nfv_service--reference--group-002.md#canonical-559678e572c7aa08a07c920a647fa9e60555651a07da394a4c9aedb3af899647) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.name](resources--nfv_service--reference--group-002.md#canonical-e8f2a004bab40706989237bbc2d694f633bf215b45c9c811faa357d5e9ab78cf) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.namespace](resources--nfv_service--reference--group-002.md#canonical-2da27d96403d003cf1639908959cda77e9a83ca6060c4808b7026a5817de1414) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.tenant](resources--nfv_service--reference--group-002.md#canonical-49a3f39e0a20732157f1ada4848581795610ef67e771947da6c51b8cc3950f0e) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca_url](resources--nfv_service--reference--group-002.md#canonical-242a86d2392c33e6242d93798c73263483cceb08aba34216a5a51788099c1b82) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled](resources--nfv_service--reference--group-002.md#canonical-3c5815595dff603936215bc5f8318fb842f148a17f7e99648ea2b7ffcca1ad4b) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_options` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_options](resources--nfv_service--reference--group-002.md#canonical-fb512985260795fd1c0cc4799711030e224a425c405ffe727011d6cf06f5542b) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_options.xfcc_header_elements](resources--nfv_service--reference--group-002.md#canonical-f695ad14a2c01b14887dda2c55eba1786331cf8d46522d545e89f36f270686c0) |
| `https_management.advertise_on_slo_internet_vip` | [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-000f5f6ab55ffd53b27bf332ae67a8bb00cf69f053e1fecd012a61caf6dab875) |
| `https_management.advertise_on_slo_internet_vip.no_mtls` | [https_management.advertise_on_slo_internet_vip.no_mtls](resources--nfv_service--reference--group-002.md#canonical-e5f3b0f001919704b049448b2e7a0b7067d066358ed7a514e97280687328baa9) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates` | [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-3f978cad766147c3a88d222418b99b63af2c330fe53ca6a634411a53db67870f) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.certificate_url` | [https_management.advertise_on_slo_internet_vip.tls_certificates.certificate_url](resources--nfv_service--reference--group-002.md#canonical-2f4ece0bc719c5541a0650e1264cbb54beb0462038cf08ac29f801c6bfb77b34) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms](resources--nfv_service--reference--group-002.md#canonical-71a7270ad0146bc368d3039d5dd4883a4f46f2339ffdae36f67cd4309d0cd399) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--nfv_service--reference--group-002.md#canonical-2263253804d8a554fe7f8b324fa6b6c58ed7014c3b97e1ccb6eb1a5ad9c5c0f0) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.description_spec` | [https_management.advertise_on_slo_internet_vip.tls_certificates.description_spec](resources--nfv_service--reference--group-002.md#canonical-3af7b41d8d10309b278922da334bf7ee0e1c103f9a04b5c221e02dae1c3474cf) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling](resources--nfv_service--reference--group-002.md#canonical-030fd2cb1f48a310b8db555074259487df3a64b9d24b235ef1f3f3a9934adb0d) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](resources--nfv_service--reference--group-002.md#canonical-cd90a3ec39661810f0ed3d28226e25d1bf9164d7a3c62da8280526d9779dd30d) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--reference--group-002.md#canonical-bb7286b22e7c70a35e84661ff4c521a464c4dd6ae52947632370805cbfe78543) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-002.md#canonical-bf869e8f447ae0abe94cbf1cbe7984b5e82e220972133283a3d0c5bd8d9c3506) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.location](resources--nfv_service--reference--group-002.md#canonical-b68a1cb45844d2ca3e19a4f994a7ed30f58d4cbdf904118cb1f805bd3fd5e379) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-002.md#canonical-9bf1ec2c62b32873868cb869290822bc786e4701e5b2f6a23c1fe3b005dcf8bd) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--reference--group-002.md#canonical-c8a8cc4a713602948b7777447730082095a2f6948fbacd994c6c3dd162f06119) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.provider_ref](resources--nfv_service--reference--group-002.md#canonical-39fb675559a22b2f2a9975e19997ae4caeb85fb5de9a954ca06edf21ac397450) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.url](resources--nfv_service--reference--group-002.md#canonical-90754b90a9c50a83d588234887364a069217346ce4b37a935c7ad2a2ebbb9295) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults](resources--nfv_service--reference--group-002.md#canonical-285b5caf9a9977872fb61c9156aedc628e17504cadfd571794da69e8b6c41530) |
| `https_management.advertise_on_slo_internet_vip.tls_config` | [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-7a4264429185ac5cbee11a32acf90af12bd5ee04265683a8755f14397de8dab8) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security](resources--nfv_service--reference--group-002.md#canonical-54993e8047f3a1b7fb616cd7ea12583d98b70d5d75d616ce5aa66768591db9af) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.cipher_suites](resources--nfv_service--reference--group-002.md#canonical-a178516ed5942be1ebc3b0199d083ffbd2aa552d092e42ce7c3b0be39da18509) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.max_version](resources--nfv_service--reference--group-002.md#canonical-112f0db57388611891b59cb23bbd7a43990d9938031bf0bab2ba40f2399aa253) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.min_version](resources--nfv_service--reference--group-002.md#canonical-4cccec009779c07c949028c91278d994aa7118d8cee0cba6f1f7c13ca0f9df03) |
| `https_management.advertise_on_slo_internet_vip.tls_config.default_security` | [https_management.advertise_on_slo_internet_vip.tls_config.default_security](resources--nfv_service--reference--group-002.md#canonical-9527220466bde32fcf88ca2b5bd2fba29f2f010b2f54e208e670a492284d31d9) |
| `https_management.advertise_on_slo_internet_vip.tls_config.low_security` | [https_management.advertise_on_slo_internet_vip.tls_config.low_security](resources--nfv_service--reference--group-002.md#canonical-6f4c0e32310fe40980b9d6f8ff788202e98b5d3fc2c69128334dd6de35e5202b) |
| `https_management.advertise_on_slo_internet_vip.tls_config.medium_security` | [https_management.advertise_on_slo_internet_vip.tls_config.medium_security](resources--nfv_service--reference--group-002.md#canonical-2e2bc386781befbde2919ff6ec17f84173905b45fdfcf5c1c93f6d838dad3a62) |
| `https_management.advertise_on_slo_internet_vip.use_mtls` | [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-238239b99e0b186337027f89b497ff97e1e443be95bde40e0c1a4a2fb79031c6) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_internet_vip.use_mtls.client_certificate_optional](resources--nfv_service--reference--group-002.md#canonical-a5087ee407ed7bdb70fb6b529eefd9c3b3ad143513e57b72696d1570d6f7918e) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl](resources--nfv_service--reference--group-002.md#canonical-cc457d9583dbe52a37abe10fe2fc40db837b42d94664ecb811053a56fa1ccf8a) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.name` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.name](resources--nfv_service--reference--group-002.md#canonical-0c7e64eea5fd33b579482d80f7d8f7283188bae34812f55c283a15451d465abd) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.namespace` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.namespace](resources--nfv_service--reference--group-002.md#canonical-4f5063317c9b71b0226c530a9b9b882d45cfa6437deefc73460384f92f366f62) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.tenant` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.tenant](resources--nfv_service--reference--group-002.md#canonical-520882b7c29eadcc9a81e3eb88eb901db2b61df769bce1f8d3eade9e94d9e771) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.no_crl` | [https_management.advertise_on_slo_internet_vip.use_mtls.no_crl](resources--nfv_service--reference--group-002.md#canonical-1d3f5a0415a5f101c780089a936cfedfe03b37bdd2a0c8c6c8b6796d95698909) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca](resources--nfv_service--reference--group-002.md#canonical-e6814ab6b9764ab73914070e1ec90991be3ce4f57514d5d6683923fb9c8756e6) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.name](resources--nfv_service--reference--group-002.md#canonical-e660aaeafb36729fbca234023e6741acda7d9403c3abddf5bde70d3e5e35ae20) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.namespace](resources--nfv_service--reference--group-002.md#canonical-8c8bf9e443dd5125fa3fde937ce5fc88515a7541170599f6654269599d320679) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.tenant](resources--nfv_service--reference--group-002.md#canonical-cc5ac10d92a7cebe7d4e2d1e62920560d6aa28d8f440da037da664a660f4f6b1) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca_url](resources--nfv_service--reference--group-002.md#canonical-64a4bcc97fd004a2060e92e33287e1e449af093750d4b4e4085b293c596a8872) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled](resources--nfv_service--reference--group-002.md#canonical-decda601ba3a53169fe3362b578fd0988f98372ba453065c3c2b95a49a1abf57) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options](resources--nfv_service--reference--group-003.md#canonical-e696d1ac2cac47dc8c900cc9b0cf61631e6c4b670d5ce55a47356ab1242ca61f) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options.xfcc_header_elements](resources--nfv_service--reference--group-003.md#canonical-94c244e360da27714c01cb855c0af874a43eb2637739e167c313f2bf4e491a45) |
| `https_management.advertise_on_slo_sli` | [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-6ad2c4c38ca51dbdc0f2f54da0cc8dba14054e49bf85b81b7eaf5854bc31192c) |
| `https_management.advertise_on_slo_sli.no_mtls` | [https_management.advertise_on_slo_sli.no_mtls](resources--nfv_service--reference--group-003.md#canonical-3a3c4d5c0ae03661d7148af7eeb43e7830b0c6e2a369632e3b7c02940a8e9ecb) |
| `https_management.advertise_on_slo_sli.tls_certificates` | [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-aa796af23ed9aa6f385192234d04a302db2615a7c8c2b0d3034d9fa16b89d461) |
| `https_management.advertise_on_slo_sli.tls_certificates.certificate_url` | [https_management.advertise_on_slo_sli.tls_certificates.certificate_url](resources--nfv_service--reference--group-003.md#canonical-149d06d940e98998393c8945a35d42e0c3e4283fc6f129bf406517393ca73050) |
| `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-0b1d23a02f771aa560bb99087839d7c89ff931a460a208261afcc09b8c836504) |
| `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-4338113a94f076a297daa4c079a1ff5cc7d409bd486ada61170262f7ddc5ee6f) |
| `https_management.advertise_on_slo_sli.tls_certificates.description_spec` | [https_management.advertise_on_slo_sli.tls_certificates.description_spec](resources--nfv_service--reference--group-003.md#canonical-448a52b525745f5b42d525f68152fac0212afa96f41d302c7f979da21456ef87) |
| `https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling](resources--nfv_service--reference--group-003.md#canonical-c5bd0739c7a83bad2c4198fa544a5a780fea74a3bf904fe246b67829d94d97c3) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key` | [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-25a6ee25874ce5fbd310c488a848abcee050c025d7d55b5c639f0ec462dbf5db) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-1ac7c1c3eee7fa5873d1dc3b2f5cec5f07584c91ea5060753f209915e80fc6b0) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-003.md#canonical-ea0d6f004df46f3ada61d2873be8e63a52aa0b510f6dce8450fbaeb162838088) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.location](resources--nfv_service--reference--group-003.md#canonical-aeda5bfdd05af566154c66dafe1ac0428804e892046b808c41c6b49b7abe5de4) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-003.md#canonical-4b413584033989401741aaf86a65e5e013a6ebca14e90952172807728b9845b6) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-2b51d377d9bebe3874becc73a6c54a079885afcfbaee4288a8422cab5a81cb70) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.provider_ref](resources--nfv_service--reference--group-003.md#canonical-dea13368c828078ca6f26613795c72fb7e0b869021cb085474731eed1b4703d9) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.url](resources--nfv_service--reference--group-003.md#canonical-262934a37915e127e1960f37ebcd004d7307acee46cf3088f9e6235d955c23ff) |
| `https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults](resources--nfv_service--reference--group-003.md#canonical-a0c6155fb728a1d2d75080009d39664baaae49f275922e875cebe3810fcf9adc) |
| `https_management.advertise_on_slo_sli.tls_config` | [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-cd8544e8be3dd22fd07c05f02839352fffff955e215f9b0b2e31b77dd7fd852a) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security` | [https_management.advertise_on_slo_sli.tls_config.custom_security](resources--nfv_service--reference--group-003.md#canonical-c062269bc28137a8e983da4de769bad521b674e5ab7346de82924c81b03e7d77) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_sli.tls_config.custom_security.cipher_suites](resources--nfv_service--reference--group-003.md#canonical-7224521419785ef923ecf7ff120cac61ad2695464c924c14d99dd1f438fa8eab) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_sli.tls_config.custom_security.max_version](resources--nfv_service--reference--group-003.md#canonical-8b0230d32c7b0f75c11d6c26d435e31eab6d1c4de9a2bc24641f3d5ddd67d4de) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_sli.tls_config.custom_security.min_version](resources--nfv_service--reference--group-003.md#canonical-bd701b2d019b75aa32054e906064391f777c65ef33dc47eadaa77560c0e0b779) |
| `https_management.advertise_on_slo_sli.tls_config.default_security` | [https_management.advertise_on_slo_sli.tls_config.default_security](resources--nfv_service--reference--group-003.md#canonical-a5555bf93d1dee72f894c0c86b288023506100181b5a3a943e94f51c93a1b339) |
| `https_management.advertise_on_slo_sli.tls_config.low_security` | [https_management.advertise_on_slo_sli.tls_config.low_security](resources--nfv_service--reference--group-003.md#canonical-c5d3ac53c0c074c5d1a683d3eb5a22cbfa5d77d025d96a321d9147264fd46f8e) |
| `https_management.advertise_on_slo_sli.tls_config.medium_security` | [https_management.advertise_on_slo_sli.tls_config.medium_security](resources--nfv_service--reference--group-003.md#canonical-8a436f12937504d5a232265eb31f1211754528c69a17456e4799d523fe17ec8b) |
| `https_management.advertise_on_slo_sli.use_mtls` | [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-1108693ddc7d3e5e1a96b6c2dd0639d21bb95beb0ebc978f51f39bb6f9d63673) |
| `https_management.advertise_on_slo_sli.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_sli.use_mtls.client_certificate_optional](resources--nfv_service--reference--group-003.md#canonical-62902d3d547a034789e39abe70c6e6c00a28ab2635384d840371c6a3edac218e) |
| `https_management.advertise_on_slo_sli.use_mtls.crl` | [https_management.advertise_on_slo_sli.use_mtls.crl](resources--nfv_service--reference--group-003.md#canonical-92af364fc3ef5bcb74c130f4b506fb1e3fb7bb9552566a445c3c6dfaaa0b52ad) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.name` | [https_management.advertise_on_slo_sli.use_mtls.crl.name](resources--nfv_service--reference--group-003.md#canonical-c52a3a9e5e55e3c55c2a1cebbddb28f0b72f5b6fb81ab8d87d8dc548b0772643) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.namespace` | [https_management.advertise_on_slo_sli.use_mtls.crl.namespace](resources--nfv_service--reference--group-003.md#canonical-ac3c1b8f983edaedfe8489ecfed42314c2d223b1fe43ee221fcaae23fc37ac53) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.tenant` | [https_management.advertise_on_slo_sli.use_mtls.crl.tenant](resources--nfv_service--reference--group-003.md#canonical-3ecf2a033ea96842b475f129b24f2cabb1af01c3ef1da084a9d2639b93cb0335) |
| `https_management.advertise_on_slo_sli.use_mtls.no_crl` | [https_management.advertise_on_slo_sli.use_mtls.no_crl](resources--nfv_service--reference--group-003.md#canonical-0c1acabd82d5c3aa051a870f4cefd1682b017aeeadcdcea4e19d1ff6a879c3c5) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca](resources--nfv_service--reference--group-003.md#canonical-dd2c4b022d0a3eb15e1b5ff3741362e638758a2ca99b0510c5b7b51a76acf994) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.name](resources--nfv_service--reference--group-003.md#canonical-71c4c07308ee22f677fd73ed94ed1274ca406e617a87355f6d306d5d1bc5532f) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.namespace](resources--nfv_service--reference--group-003.md#canonical-72738dff86dce06560a799359237d4f927cd41c763d1458de861ae3ef68177ae) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.tenant](resources--nfv_service--reference--group-003.md#canonical-299269b493fa7854ce84bab7cc613e889493085b076e3d0203ae3e8c1544bf42) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca_url](resources--nfv_service--reference--group-003.md#canonical-30496d380617931f9371a4a5428be8b199e7cd9ffb9cc78729836698c4a504c7) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled](resources--nfv_service--reference--group-003.md#canonical-2c24d043f536442190167d3a8eb47c46cfcf797f17515db04cdb1fd15e88896a) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_options` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_options](resources--nfv_service--reference--group-003.md#canonical-38ae173a7133fe2a98ddcc16969193103920277175584d5cfbf1a949959668ac) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_options.xfcc_header_elements](resources--nfv_service--reference--group-003.md#canonical-6e0b30a1e85e32ebda9d9f4dc9e36c83d976b1666ad29d8953809177613458c9) |
| `https_management.advertise_on_slo_vip` | [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3531b5c834ecb32b9e38f4efc093de765d809a806e2a740169f7efaaf764206e) |
| `https_management.advertise_on_slo_vip.no_mtls` | [https_management.advertise_on_slo_vip.no_mtls](resources--nfv_service--reference--group-003.md#canonical-2b4929fbd87561bd855b8b75491587792bf121657378171098209ef10f613d3c) |
| `https_management.advertise_on_slo_vip.tls_certificates` | [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-8fd4c68ed449a0ec050a39af4e95903a9a75ffec20de722562f4a410f12e8bc0) |
| `https_management.advertise_on_slo_vip.tls_certificates.certificate_url` | [https_management.advertise_on_slo_vip.tls_certificates.certificate_url](resources--nfv_service--reference--group-003.md#canonical-5838ffcce3448515b0655a8a279edfdfeaf68f30b3583c397425d876606e56ab) |
| `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-5b5b17b671aab13747e56d3d9c7c1dbebea803b371b44c74d88e5f85ada90e74) |
| `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-e7b1aae1306ce55c4af9dab21175bd7d72764fe78df43d9e6812a9324e508ebf) |
| `https_management.advertise_on_slo_vip.tls_certificates.description_spec` | [https_management.advertise_on_slo_vip.tls_certificates.description_spec](resources--nfv_service--reference--group-003.md#canonical-7152c6b4c0d6f37ed954d080c4b01c51518dde59275068ddbbef5a39cbeab228) |
| `https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling](resources--nfv_service--reference--group-003.md#canonical-2965c9c329c90b457ef2e38700e6f277a64d37e4479a5fe77c58656f3bc3cad6) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key` | [https_management.advertise_on_slo_vip.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-9b1197291d728060c51a9aa4530085b51530df2b3e86277c75f71ba661a55f50) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-7d579e8fe6766374ad2c9fbaff0d18887c7b273958e438c6ac2d5567e33d52bb) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-003.md#canonical-e53e392fc9413e65179f3be8832a2a651d6bdabd01a7387e6526f1dd3d4e7c1e) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.location](resources--nfv_service--reference--group-003.md#canonical-5e9fd1503101eca30bc77d8867c1386c3ce45c851e09a472b19ad7341b0624b3) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-003.md#canonical-3ca515216ff9f3895f458e62a44a3a3c9b7a841b67f9e882870abf48ee409c4c) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-d68bf7388433afa5dcbe11ba69c8e461a25edf6e51577cfcf990b419f8b5102e) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.provider_ref](resources--nfv_service--reference--group-003.md#canonical-bcd29b0c8f0c3c6de8b75b3169d72cfec4768ab9dbe72f258be28876fc185c31) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.url](resources--nfv_service--reference--group-003.md#canonical-7c0170f46996e2a0fededd199faa5ceb68da829761b208ef8ac4b3ec75399999) |
| `https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults](resources--nfv_service--reference--group-003.md#canonical-8d1ba086845ea53aa1a2b0229366bd7a04924f882490d0a549b0415309613852) |
| `https_management.advertise_on_slo_vip.tls_config` | [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-778600701aca3cc945ab71d9ea9912e0a2d10ae8b8fa5e944cf3ebd800a0c00d) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security` | [https_management.advertise_on_slo_vip.tls_config.custom_security](resources--nfv_service--reference--group-003.md#canonical-51ee633e44aeec128ffbcf858f572750ef1b41cacdd8a94e09d703fb664f3dac) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_vip.tls_config.custom_security.cipher_suites](resources--nfv_service--reference--group-003.md#canonical-3db08b8a600f81980f3415d86226cd6476c87b3bb8325af83653352c0428f99d) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_vip.tls_config.custom_security.max_version](resources--nfv_service--reference--group-003.md#canonical-ea8191063cb2062f6c810021dc07f73cf89bc4c63c3e8f2f97517bbdb6406c4c) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_vip.tls_config.custom_security.min_version](resources--nfv_service--reference--group-003.md#canonical-b907185baf9e64d4aa84029b626c54bcbc8439422e33b95a9a99f5357c48bf72) |
| `https_management.advertise_on_slo_vip.tls_config.default_security` | [https_management.advertise_on_slo_vip.tls_config.default_security](resources--nfv_service--reference--group-003.md#canonical-fe30954c57479a8f21b23b89a7270dc5a565934dab0f096c46c2697c9af27b5d) |
| `https_management.advertise_on_slo_vip.tls_config.low_security` | [https_management.advertise_on_slo_vip.tls_config.low_security](resources--nfv_service--reference--group-003.md#canonical-40bfab942dcc19c6c3923ede13181a5d4be2159816d37e770f806ded9e3768e5) |
| `https_management.advertise_on_slo_vip.tls_config.medium_security` | [https_management.advertise_on_slo_vip.tls_config.medium_security](resources--nfv_service--reference--group-003.md#canonical-0bcc7fe65b417deca064a8b3b8cfee50900ff66aed671df81ccd9805f05723fa) |
| `https_management.advertise_on_slo_vip.use_mtls` | [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-a058339644255cc34b12f45e16622453c8a7610525973fdd0527448879bf2904) |
| `https_management.advertise_on_slo_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_vip.use_mtls.client_certificate_optional](resources--nfv_service--reference--group-003.md#canonical-b7f391c2e21c2f2aef5c7a518ff4230eec01e809a02dded5a938c872abda2022) |
| `https_management.advertise_on_slo_vip.use_mtls.crl` | [https_management.advertise_on_slo_vip.use_mtls.crl](resources--nfv_service--reference--group-003.md#canonical-9c0adfe477a9f53262cf89cd58aa413f97d717ca2b525cbe811b0fb44fb1f562) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.name` | [https_management.advertise_on_slo_vip.use_mtls.crl.name](resources--nfv_service--reference--group-003.md#canonical-02e7bede44c7332db85c270748c682dfb0e20ef653b8a6bc585e9f46246317ac) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.namespace` | [https_management.advertise_on_slo_vip.use_mtls.crl.namespace](resources--nfv_service--reference--group-003.md#canonical-f69a894ee16ca97c8fce7eb6c6045cafc39cb4f9cedf682bb2c5e3bc9ae64b71) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.tenant` | [https_management.advertise_on_slo_vip.use_mtls.crl.tenant](resources--nfv_service--reference--group-003.md#canonical-d382ba1156d991af2656bb16a1ba0f48e5d6f3a310cc4804f7df83ab35cb8083) |
| `https_management.advertise_on_slo_vip.use_mtls.no_crl` | [https_management.advertise_on_slo_vip.use_mtls.no_crl](resources--nfv_service--reference--group-003.md#canonical-9ef83c13d7142da78d7a009232713dd5dd5548f485a74d9a89ac96462fb243e1) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca](resources--nfv_service--reference--group-003.md#canonical-cd3105831bd9cf983b42b83ecc2606973143c29280b1650fa6030a00d39402b8) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.name](resources--nfv_service--reference--group-003.md#canonical-734c5e7a3b9952bd0af1b942b7ac13b5c5072c9fbdff0d09a9dddd89ccc130bf) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.namespace](resources--nfv_service--reference--group-003.md#canonical-e51c8cecc51ef90fb8cf82f667cede501d1d72709c0be588c5fdc88d43e54182) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.tenant](resources--nfv_service--reference--group-003.md#canonical-6ad41ef4192875b507bc061990b2c0483faf16ef29df4f878b1871fc46abd6c9) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca_url](resources--nfv_service--reference--group-003.md#canonical-1f1843db987714a4d0d4520377eaa28a9a67413556759cd420e32f3563aeb5fa) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled](resources--nfv_service--reference--group-003.md#canonical-8ee97d0f2be9cb7facd806fa0dd55ff6a6f7c52e04c5c2d2bdd0d38513a74972) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_options` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_options](resources--nfv_service--reference--group-003.md#canonical-bfb6defb9f2a4323c3836436f0937744a856bafde7af20b85b009e784453cd73) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_options.xfcc_header_elements](resources--nfv_service--reference--group-003.md#canonical-afb7a24ac3224edafdf5940c257bdc4223221fc12111192a545e238880100791) |
| `https_management.default_https_port` | [https_management.default_https_port](resources--nfv_service--reference--group-003.md#canonical-3894cf08a900164b908fcb9ff5b3d8413a0e9be08a34feff162eb595cae8e202) |
| `https_management.domain_suffix` | [https_management.domain_suffix](resources--nfv_service--reference--group-002.md#canonical-dde6155ccfaa0e6d618b72ed7f5b47333a61d150d8b6aaa0c69c16ac8bf1fa4b) |
| `https_management.https_port` | [https_management.https_port](resources--nfv_service--reference--group-002.md#canonical-1b2cebb33b94189f870fa07a4eda9b6e96ab7bc92c8b8b5fdf46aa0080f80501) |
| `id` | [id](resources--nfv_service--reference--group-001.md#canonical-a2903b099bf7ca67a81ed51a2fdd1bb69ec44bd5fda68d42d0f20191ecd0395c) |
| `labels` | [labels](resources--nfv_service--reference--group-001.md#canonical-db8f6c88495cf9e9d44072956238517d62121d94c5bc73205248fc391cd71024) |
| `name` | [name](resources--nfv_service--reference--group-001.md#canonical-17d8743bc18fcde39ca586e09fbad1ff9d1eab6b01b3faea5707abe8f59ed6ec) |
| `namespace` | [namespace](resources--nfv_service--reference--group-001.md#canonical-9e97ed29393906a10dcc2b427f50a8b48c9c1ab30132a543ee63790db0c1baca) |
| `palo_alto_fw_service` | [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-409735399e3d2be8e2f34d7145e44073824af1c2148dd24576cf8c92863b1104) |
| `palo_alto_fw_service.auto_setup` | [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-f4c86c8e1e1632d1724101dc4457dd9dae42e8c25ddaf04bdee5d21e0a7bfc3a) |
| `palo_alto_fw_service.auto_setup.admin_password` | [palo_alto_fw_service.auto_setup.admin_password](resources--nfv_service--reference--group-003.md#canonical-9ff11183e8534b6f1990b81542da2981c1f0b50fbcd6389be94307fc2d03df45) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-f6e42f77395b0a1d6a4308281d87c983b65e3f0ae8fd551fee485e9a01071b34) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-003.md#canonical-faef2b0ccbae06875480ff758ce62c6ef09fc7e072ba2dd1495a7b7320d34d7e) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.location` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.location](resources--nfv_service--reference--group-003.md#canonical-fcdbf12f1809220226f2a0f347a5cbae43a8f48e1e50a30776cd2293fad007c3) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.store_provider` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-003.md#canonical-8b50636744e9c723c9a3917d98a8e8a5d57b9990af7acb7fa5b79950b5df2682) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-ec221483893bc68fb5afe59fa02f92e3e1c2be08a15d72a437ad96858cc0decf) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.provider_ref` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.provider_ref](resources--nfv_service--reference--group-003.md#canonical-e51ded7440edba3af7b44f54fcd155f5abd33281a74b952a8646ddf16ab2575e) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.url` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.url](resources--nfv_service--reference--group-003.md#canonical-1ed5c2f73f14ebcda659153569e128cf9cc139a3833d14edd3f5a4552f658a8e) |
| `palo_alto_fw_service.auto_setup.admin_username` | [palo_alto_fw_service.auto_setup.admin_username](resources--nfv_service--reference--group-003.md#canonical-804d3e51d1e12b4a3465ee2b89eec0ae71e8d900c7432ac28ed363cc842c2bb0) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys` | [palo_alto_fw_service.auto_setup.manual_ssh_keys](resources--nfv_service--reference--group-003.md#canonical-2f6a959a99036fbb247fe691ddea6f65fa4b170d7fde29be419061b84acf2c97) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](resources--nfv_service--reference--group-003.md#canonical-be9c868da4cc8fcd2dec68114743f65c9b28ac5ef6d6daf01a9cef6f2944d4a1) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-6ba4f9d338374a9feff11dd54b618f9278ce692bd0347b867eb3040309b2ad65) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-004.md#canonical-47b1187f3dcda3b6fa162b1b6aab769fdfab3b49fa3c7e0ff7bd4422181233df) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.location` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.location](resources--nfv_service--reference--group-004.md#canonical-7416049b2a6dddd35de75eb2fc9cb35c9b5afd41f80ad369b4d9701154f2143c) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.store_provider` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-004.md#canonical-c3b1ce4b8c312640ac81f0627c765be7cac34718a919dbee76834304d22c743c) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-d62311ee0f6873838214453f8462545c1e2ca30eb4ebb1a54a35cfdc69abfb0f) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.provider_ref` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.provider_ref](resources--nfv_service--reference--group-004.md#canonical-4aae2f48d1e4ea7240a685c3946b5bdd9398406a6dcfaf96f73bc050ff5ea8d2) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.url` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.url](resources--nfv_service--reference--group-004.md#canonical-ce4b7b21a6a3a1ce365da0e667c983638ee14d9e42671a1142088fdbd8cd07eb) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.public_key` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.public_key](resources--nfv_service--reference--group-003.md#canonical-dd567610e9bf5188e3c958750c2069debbd5cb832d34499cf6ea3d297a3b17f9) |
| `palo_alto_fw_service.aws_tgw_site` | [palo_alto_fw_service.aws_tgw_site](resources--nfv_service--reference--group-004.md#canonical-06fb95e7aa6dede5ce324762abc7457cb971e96dbb1a76fc9e050b85740c49aa) |
| `palo_alto_fw_service.aws_tgw_site.name` | [palo_alto_fw_service.aws_tgw_site.name](resources--nfv_service--reference--group-004.md#canonical-6c7e21ed60de62cc861575ea2d4442dc195ebb2d7ac5ced1c80a0b7d5ecbe3f5) |
| `palo_alto_fw_service.aws_tgw_site.namespace` | [palo_alto_fw_service.aws_tgw_site.namespace](resources--nfv_service--reference--group-004.md#canonical-dad0fcf09c249a25fd60dbc0ed2e6a60a49d44194e646513a941bb79110bc396) |
| `palo_alto_fw_service.aws_tgw_site.tenant` | [palo_alto_fw_service.aws_tgw_site.tenant](resources--nfv_service--reference--group-004.md#canonical-9ade89855a8cf733c131fd48ce9a7fd7bb176c84b468369acaddbc0fd15b7259) |
| `palo_alto_fw_service.disable_panaroma` | [palo_alto_fw_service.disable_panaroma](resources--nfv_service--reference--group-004.md#canonical-c21e39b36de4a2e3578d2cd7a8fd908cd26c30e1d23dde0232d2b64dc0ae6670) |
| `palo_alto_fw_service.instance_type` | [palo_alto_fw_service.instance_type](resources--nfv_service--reference--group-003.md#canonical-a76b746ee398a519d7abb7127e51b340798ee98ec4d46b61dc3ae10543c8cadd) |
| `palo_alto_fw_service.pan_ami_bundle1` | [palo_alto_fw_service.pan_ami_bundle1](resources--nfv_service--reference--group-004.md#canonical-7a56bae860591078e7e5798b9521b4de14942370395a086b03aec160d73d22e7) |
| `palo_alto_fw_service.pan_ami_bundle2` | [palo_alto_fw_service.pan_ami_bundle2](resources--nfv_service--reference--group-004.md#canonical-5a14a412eb31509e25439ce526615379bd5cbea702ce2ed8462aff74dbd96cb6) |
| `palo_alto_fw_service.panorama_server` | [palo_alto_fw_service.panorama_server](resources--nfv_service--reference--group-004.md#canonical-d367e7f39b3d910eb1713b0fe454f2d55c32800fb9a045a277dce3704693f29a) |
| `palo_alto_fw_service.panorama_server.authorization_key` | [palo_alto_fw_service.panorama_server.authorization_key](resources--nfv_service--reference--group-004.md#canonical-4c0d99a8630ee8137ec0fa6bb6e1cbdfc41159d590e4351a19562837b3383a62) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-6428b5cf3cfe6a0473b99878f749bc3ff0d7f43c2f85415123ee46d938522dc9) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-004.md#canonical-971413d99e57529afa4bc414f63e860e6df5bed0d302b69800e3b12f111bde69) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.location` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.location](resources--nfv_service--reference--group-004.md#canonical-8146d74bcd2858d1dab58432c331441f782c04ccf3243a0e636c96b87832c834) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.store_provider` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-004.md#canonical-a44e7227699d65365176b6fb0a1bc18fa3990256244dbe9b981adb3925ba2ae4) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-2e84232c05d795390c14466bf502302b73bf86afdd2a418c4b129e753c09a370) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.provider_ref` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.provider_ref](resources--nfv_service--reference--group-004.md#canonical-87bd3e37ad89547ef42f316e123ce66a8735288c3485f6143a57225845f18aef) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.url` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.url](resources--nfv_service--reference--group-004.md#canonical-049ab609317fc0c222c0583699c67ee46f4999cd828db90fcb027f32a732f1d4) |
| `palo_alto_fw_service.panorama_server.device_group_name` | [palo_alto_fw_service.panorama_server.device_group_name](resources--nfv_service--reference--group-004.md#canonical-d30d7aa5f847d278617534c713f0099cb9bd788f9ca76d2482bdc9183639f072) |
| `palo_alto_fw_service.panorama_server.server` | [palo_alto_fw_service.panorama_server.server](resources--nfv_service--reference--group-004.md#canonical-6e913aebe09c93d7062936b5fa3ba9bc357a1bd1fa8ec9a426e18deddde65238) |
| `palo_alto_fw_service.panorama_server.template_stack_name` | [palo_alto_fw_service.panorama_server.template_stack_name](resources--nfv_service--reference--group-004.md#canonical-0bd3cc822ce9dce4c2f4bae5dab7910249e40641ef3795165276944d8a7ccbdc) |
| `palo_alto_fw_service.service_nodes` | [palo_alto_fw_service.service_nodes](resources--nfv_service--reference--group-004.md#canonical-72e64a43fbb90dbf2f5723ec26c747305a7dc378e24b4a51b262380ce816370a) |
| `palo_alto_fw_service.service_nodes.nodes` | [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--reference--group-004.md#canonical-a45394d945aae609dcab8272ea96f8a4dff274534e81808b0005d3af98c60f2f) |
| `palo_alto_fw_service.service_nodes.nodes.aws_az_name` | [palo_alto_fw_service.service_nodes.nodes.aws_az_name](resources--nfv_service--reference--group-004.md#canonical-34f512f9488f5da8b2fefd7b60d32e53cf54084e706d5a8b594ad9788b3ae5c9) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-f0ce6950ebabbe86a9f32a48446eac00937acbc1b3a78418ebc136b88e689f70) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.existing_subnet_id` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.existing_subnet_id](resources--nfv_service--reference--group-004.md#canonical-3cc89bcaa3bd0c69c9a536c4202b5a70590691e32d71526a4823628f378f2bdd) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param](resources--nfv_service--reference--group-004.md#canonical-6dd099fe25611dc2860495a878613d10cc0efd48dadf9a6a1792bdc4a868f82f) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param.ipv4` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param.ipv4](resources--nfv_service--reference--group-004.md#canonical-adc624dfb24fbda6ac0e4fe0f40607985a4917e46363f578275421488ffc3d7a) |
| `palo_alto_fw_service.service_nodes.nodes.node_name` | [palo_alto_fw_service.service_nodes.nodes.node_name](resources--nfv_service--reference--group-004.md#canonical-b92f7dd2fb935770a643cd2e0557113696a919e75a70c826b55052aa9d704ee7) |
| `palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet` | [palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-e70eda96bf31c1ad2f7bf668627cc5ec520812d29867a54f8edd34d9f48ce5de) |
| `palo_alto_fw_service.ssh_key` | [palo_alto_fw_service.ssh_key](resources--nfv_service--reference--group-003.md#canonical-06e9d71b182ee53300cb5ba69196cf3709d569ef154907f1f86828e51357db69) |
| `palo_alto_fw_service.tags` | [palo_alto_fw_service.tags](resources--nfv_service--reference--group-003.md#canonical-aa82a4b6a78a00bcba62b64625124cfc7d637a63e73a5948c26dc0dd54e22f6e) |
| `palo_alto_fw_service.version` | [palo_alto_fw_service.version](resources--nfv_service--reference--group-003.md#canonical-df873e96e04ff6469b73b68430a8994f4a57501f64657cc266b3342cad01c574) |
| `timeouts` | [timeouts](resources--nfv_service--reference--group-004.md#canonical-ae9e0c0d2d16c9868cadf460e69bb17c9e1e22f0e0c3c5c6ffe73466fe9ef49a) |
| `timeouts.create` | [timeouts.create](resources--nfv_service--reference--group-004.md#canonical-eaba720097bc356f1e0eff85431f614c806744ca7d3f7f4dbc238ac990ea9ac6) |
| `timeouts.delete` | [timeouts.delete](resources--nfv_service--reference--group-004.md#canonical-801d2ae594eda82dfe6af70557aaa2c01812c8afbdf8f2b7acccb839dfde8539) |
| `timeouts.read` | [timeouts.read](resources--nfv_service--reference--group-004.md#canonical-734f0daabe3ed3e505ef32ffec0b5cb3c153389f615edc930a6c8c619a2aebbb) |
| `timeouts.update` | [timeouts.update](resources--nfv_service--reference--group-004.md#canonical-aeefb5859002af36a87292f211b5e4ffcad6ef54ac594361cb627c95cbb9a7d0) |

<a id="canonical-fe92e27f7e9759d44185af39553798e09f9ad23c6b3ca5c93a8242c07a0bfc87"></a>

## Next pages — Property reference / b32dfc5e9333 / 12

- [disable_https_management](resources--nfv_service--reference--group-001.md#canonical-6725ac95b37b2f5e8f60a7b62e92fcfb81391da7dd183f8b11259ecd909c2bd0)
- [disable_ssh_access](resources--nfv_service--reference--group-001.md#canonical-59b17828f1392d7b77df4c061dc7266ffbbf235c736ab86ffe900270e0b5c4ce)
- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-6fef4a6194444c7af1b6d53343496678464efb8500bbd53ba2714a605fd63804)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-683b0e043ad34e16d220e00d2c0ad7d7794cd7a996943065814c81f0d29e03cf)
- [timeouts](resources--nfv_service--reference--group-004.md#canonical-22b89a5a007d5d26888d6c80645783ec0135cd1ced1c3d4d29ec4c3cdc0ebce1)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-6725ac95b37b2f5e8f60a7b62e92fcfb81391da7dd183f8b11259ecd909c2bd0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b2c658c5a83e5fd5319cedba50fe08ff841bbc6bc24cac7d4efda1a9b5b748d"></a>

## disable_https_management — disable_https_management / 8b53592d76c9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- disable_https_management

<a id="canonical-d6d9c1c4276b1dba6fb22647aeddc98c0a711c4f114d4b1f9cc65952716eb654"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_https\_management, https\_management; Default: disable\_https\_management\]
Configuration parameter for disable https management.

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

- [disable_https_management](resources--nfv_service--reference--group-001.md#canonical-d6d9c1c4276b1dba6fb22647aeddc98c0a711c4f114d4b1f9cc65952716eb654)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-b4518939056897104846d93f414610b8e7ef75ac36012d24571976aa6b8b425c)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_https_management = {}
```

<a id="canonical-4848ce787fe7f9106c4e1ad4a8e7b167054295235d6026cfb22c29aa7dbdeca0"></a>

## Direct properties — disable_https_management / 8b53592d76c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c77eb023dccbc764c8f9c41886e01fe0460af3ebc1ec5588739c2c8cb51aa9a4"></a>

## Next pages — disable_https_management / 8b53592d76c9 / 4

- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-59b17828f1392d7b77df4c061dc7266ffbbf235c736ab86ffe900270e0b5c4ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-051ea45dafa724cf971ae33645c8bde8ad1ee24802fe4d0c0fef729feca90752"></a>

## disable_ssh_access — disable_ssh_access / a2443e622cbf / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- disable_ssh_access

<a id="canonical-1577f6dcb281d05115836289f252e63d045d9ac1a5aa5e5c88e9a1fba796d15b"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_ssh\_access, enabled\_ssh\_access; Default: disable\_ssh\_access\] Configuration
parameter for disable ssh access.

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

- [disable_ssh_access](resources--nfv_service--reference--group-001.md#canonical-1577f6dcb281d05115836289f252e63d045d9ac1a5aa5e5c88e9a1fba796d15b)
- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-21ca4809441746237ba51aa036221a5e06bd4b462d112224ce9efc59fbb1ca93)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ssh_access = {}
```

<a id="canonical-bb5286c09c7c483e61c4cde11d2afbe762dbcb1ed7b87514dca09e4ef60d390a"></a>

## Direct properties — disable_ssh_access / a2443e622cbf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0fadf2e805baac16dbb035780a64d618681085b7dbabd1cfa40726196dfd754d"></a>

## Next pages — disable_ssh_access / a2443e622cbf / 4

- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-6fef4a6194444c7af1b6d53343496678464efb8500bbd53ba2714a605fd63804"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-771df266fc49389f8109a442794f05af816fe934711c5faeddd9916060c1ce09"></a>

## enabled_ssh_access — enabled_ssh_access / 2ea847f766f0 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- enabled_ssh_access

<a id="canonical-21ca4809441746237ba51aa036221a5e06bd4b462d112224ce9efc59fbb1ca93"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enabled ssh access.

Upstream description:

SSH based configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domain_suffix",
    "node_ssh_ports"),
  validators.ConflictingObjectAttributes("advertise_on_sli",
    "advertise_on_slo"),
  validators.ConflictingObjectAttributes("advertise_on_sli",
    "advertise_on_slo_sli"),
  validators.ConflictingObjectAttributes("advertise_on_slo",
    "advertise_on_slo_sli")}
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
  "x-ves-oneof-field-advertise_choice": "[\"advertise_on_sli\",\"advertise_on_slo\",\"advertise_on_slo_sli\"]"
}
```

Terraform syntax:

```terraform
enabled_ssh_access {
  # Configure direct properties listed below.
}
```

<a id="canonical-4710e0f7ef420c4f9611ab385d10249215d547616fcfacf893a446488e060783"></a>

## Direct properties — enabled_ssh_access / 2ea847f766f0 / 3

- [advertise_on_sli](resources--nfv_service--reference--group-001.md#canonical-114aea97d0f5e9d5da8464aca2e5a0c42741a436c73a05d2e322ebbd56839e4c): complete subsection reference.

- [advertise_on_slo](resources--nfv_service--reference--group-001.md#canonical-b76bcf31538f6467895c33d7de0dc3afae8ba6d75097a10e0f6e74ee78163c20): complete subsection reference.

- [advertise_on_slo_sli](resources--nfv_service--reference--group-001.md#canonical-7c346903a8cbc2a3dbc3eac112da2b875528331c0726be8fc7b155f303f2a60e): complete subsection reference.

<a id="canonical-88877cf04488f92b9160c2fe4bb32bbbf0b676af405c954815be89e9a44bf1b5"></a>

<a id="canonical-958cdcf39c435be976042e4becb17991ac852bcdf89b3880b892da5a0acda31c"></a>

## domain_suffix property — enabled_ssh_access / 2ea847f766f0 / 4

Type: `"string"`. Optional.

Domain suffix will be used along with node name to form the hostname for SSH node management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [node_ssh_ports](resources--nfv_service--reference--group-001.md#canonical-b858178c33cabfa48b86f8868d27282919791973b449327acea0979bd8e32992): complete subsection reference.

<a id="canonical-6571241649d66a9eb6440f6d2588727ae2e0bd43226fab44afcbc1089787a64c"></a>

## Next pages — enabled_ssh_access / 2ea847f766f0 / 5

- [enabled_ssh_access.advertise_on_sli](resources--nfv_service--reference--group-001.md#canonical-114aea97d0f5e9d5da8464aca2e5a0c42741a436c73a05d2e322ebbd56839e4c)
- [enabled_ssh_access.advertise_on_slo](resources--nfv_service--reference--group-001.md#canonical-b76bcf31538f6467895c33d7de0dc3afae8ba6d75097a10e0f6e74ee78163c20)
- [enabled_ssh_access.advertise_on_slo_sli](resources--nfv_service--reference--group-001.md#canonical-7c346903a8cbc2a3dbc3eac112da2b875528331c0726be8fc7b155f303f2a60e)
- [enabled_ssh_access.node_ssh_ports](resources--nfv_service--reference--group-001.md#canonical-b858178c33cabfa48b86f8868d27282919791973b449327acea0979bd8e32992)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-114aea97d0f5e9d5da8464aca2e5a0c42741a436c73a05d2e322ebbd56839e4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-190b52062ef5681db6bcc6767a302e36c153ca3a97c0efd6451454b7ecb0eae4"></a>

## enabled_ssh_access.advertise_on_sli — enabled_ssh_access.advertise_on_sli / 15d2217eb6bf / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-6fef4a6194444c7af1b6d53343496678464efb8500bbd53ba2714a605fd63804)
- enabled_ssh_access.advertise_on_sli

<a id="canonical-031d045d5a861b2a41a07323d365f48e8726b17104e129dc6c37ec6515ecd84e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for advertise on sli.

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
advertise_on_sli = {}
```

<a id="canonical-b4ab5f7c8069288504a2f3d767cf3344f57c543b35dfcdabe720f8d0ff404159"></a>

## Direct properties — enabled_ssh_access.advertise_on_sli / 15d2217eb6bf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bece67753be622e7bdfd7a52c5a018fa9619206541e375c14c806d4fcc5567c4"></a>

## Next pages — enabled_ssh_access.advertise_on_sli / 15d2217eb6bf / 4

- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-6fef4a6194444c7af1b6d53343496678464efb8500bbd53ba2714a605fd63804)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-b76bcf31538f6467895c33d7de0dc3afae8ba6d75097a10e0f6e74ee78163c20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ca312e5b9c285b007aa74890d5ea998cc89464770cdcf83b16f4d04548e184e"></a>

## enabled_ssh_access.advertise_on_slo — enabled_ssh_access.advertise_on_slo / 61abc86147fa / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-6fef4a6194444c7af1b6d53343496678464efb8500bbd53ba2714a605fd63804)
- enabled_ssh_access.advertise_on_slo

<a id="canonical-d730517363dbe5215b6ffdc2a8284e56de2a4bb1bad6e55618f6103332673314"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for advertise on slo.

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
advertise_on_slo = {}
```

<a id="canonical-103f4fb5ecb428ebf13dc1ce088432c060807f3d64464cedeca22dca7b8441f6"></a>

## Direct properties — enabled_ssh_access.advertise_on_slo / 61abc86147fa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-16b0771d5308b998cf533adb9d23838b9b0ba921cbfef9529410318b04ba39d7"></a>

## Next pages — enabled_ssh_access.advertise_on_slo / 61abc86147fa / 4

- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-6fef4a6194444c7af1b6d53343496678464efb8500bbd53ba2714a605fd63804)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-7c346903a8cbc2a3dbc3eac112da2b875528331c0726be8fc7b155f303f2a60e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07fc5e8fe77e7561acfa20f06346b7da8b0ded634c8cbcb94657b425ed348f7e"></a>

## enabled_ssh_access.advertise_on_slo_sli — enabled_ssh_access.advertise_on_slo_sli / 185b4fd658a9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-6fef4a6194444c7af1b6d53343496678464efb8500bbd53ba2714a605fd63804)
- enabled_ssh_access.advertise_on_slo_sli

<a id="canonical-b16d6e3c7f0a6ee10c1f5be35079140283e75c0554838cae2d7226036dbcac2e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for advertise on slo sli.

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
advertise_on_slo_sli = {}
```

<a id="canonical-16ee92ee475fdda78f6c0457a8bf7cce6fbeaeb8559676a047deeee691e2c587"></a>

## Direct properties — enabled_ssh_access.advertise_on_slo_sli / 185b4fd658a9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-36848047f2ebd1ad9308c011f08829968144bdf8bbbfbc287a3751e208cabb7b"></a>

## Next pages — enabled_ssh_access.advertise_on_slo_sli / 185b4fd658a9 / 4

- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-6fef4a6194444c7af1b6d53343496678464efb8500bbd53ba2714a605fd63804)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-b858178c33cabfa48b86f8868d27282919791973b449327acea0979bd8e32992"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ff59b651221e7ad00241f2f72aad695cccd2733a146a568a4a49e279cd59167"></a>

## enabled_ssh_access.node_ssh_ports — enabled_ssh_access.node_ssh_ports / e20cb302bde2 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-6fef4a6194444c7af1b6d53343496678464efb8500bbd53ba2714a605fd63804)
- enabled_ssh_access.node_ssh_ports

<a id="canonical-dd9b73e94db883cea6541a526fefc5c3e1ea35c6826124786c6ed6432a89fd53"></a>

Type: `"object"`. list nested block, Optional.

Management Node SSH Port. Enter TCP port and node name per node.

Upstream description:

Enter TCP port and node name per node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("node_name",
    "ssh_port")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 2,
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
    "ves.io.schema.rules.repeated.max_items": "2"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "2"
  }
}
```

Terraform syntax:

```terraform
node_ssh_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-f3ab6f9162a38a216d1dc0832fb4686886c48a28a6b26b425aaddd9cebd05640"></a>

## Direct properties — enabled_ssh_access.node_ssh_ports / e20cb302bde2 / 3

<a id="canonical-352fcaabadda12fd7449f41afb05ace4a984317107147ec13d713b6172d70528"></a>

<a id="canonical-9ab4cd9f038b69edf900a14948a18420ece10982be4ac59fdd2e2d62d7b6d457"></a>

## node_name property — enabled_ssh_access.node_ssh_ports / e20cb302bde2 / 4

Type: `"string"`. Optional.

Node name will be used to match a particular node with the desired TCP port.

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

<a id="canonical-e7250fd238e29cc88540125cfa00a2ef13c2868d69951175353242565ce8c5f8"></a>

<a id="canonical-b49e15b314b0632cc86e458593d3542fcbcebf6553b7bee4100ecf3af602ae52"></a>

## ssh_port property — enabled_ssh_access.node_ssh_ports / e20cb302bde2 / 5

Type: `"number"`. Optional.

SSH Port. Enter TCP port per node.

Upstream description:

Enter TCP port per node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1024, 65535),
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
    "minimum": 1024
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-a083e12a2015d7f5715e7bc9653a87f635f1ddf38656a1f381cb21a136efe499"></a>

## Next pages — enabled_ssh_access.node_ssh_ports / e20cb302bde2 / 6

- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-6fef4a6194444c7af1b6d53343496678464efb8500bbd53ba2714a605fd63804)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c8289d7544761bce7edad3ec5be53695636ce2ce3d2f55dfae7f2a2736acbfc"></a>

## f5_big_ip_aws_service — f5_big_ip_aws_service / d0432981823c / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- f5_big_ip_aws_service

<a id="canonical-2142caf3816f79bc0b5a8dc3fdd65f9222695b01ae57af4744368f7ffdbceeb9"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: f5\_big\_ip\_aws\_service, palo\_alto\_fw\_service\] Virtual BIG-IP AWS. Virtual BIG-IP
specification for AWS.

Upstream description:

Virtual BIG-IP specification for AWS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("admin_username",
    "nodes",
    "ssh_key")}
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
  "x-ves-oneof-field-image_choice": "[\"market_place_image\"]",
  "x-ves-oneof-field-site_type_choice": "[\"aws_tgw_site_params\"]"
}
```

OneOf alternatives in this subsection:

- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2142caf3816f79bc0b5a8dc3fdd65f9222695b01ae57af4744368f7ffdbceeb9)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-409735399e3d2be8e2f34d7145e44073824af1c2148dd24576cf8c92863b1104)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
f5_big_ip_aws_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-13d0713b30290b6f39aaaa0b4702688bda9114e2e47dc5bc8c5aa6b973dc854e"></a>

## Direct properties — f5_big_ip_aws_service / d0432981823c / 3

- [admin_password](resources--nfv_service--reference--group-001.md#canonical-0421e522b2c40b190f687c7bdd92b8c14169b71f32b7559395cc37737a1679dc): complete subsection reference.

<a id="canonical-33ce4b0704b55befc01cb2d1569d2b02b33f81b5f8f6b969c9c22e4bc6669aa8"></a>

<a id="canonical-d366f96a917acf1c45041bb283839b689064e3cc2d153b6eac00289ef290243d"></a>

## admin_username property — f5_big_ip_aws_service / d0432981823c / 4

Type: `"string"`. Optional.

Admin Username. Admin Username for BIG-IP.

Upstream description:

Admin Username for BIG-IP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [aws_tgw_site_params](resources--nfv_service--reference--group-001.md#canonical-afd6a080907c4e229ae43f4329cccdd21898342797d2a034dc6c3ae64906a6a4): complete subsection reference.

- [endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50): complete subsection reference.

- [market_place_image](resources--nfv_service--reference--group-001.md#canonical-d9a50ca9b65a69d0b0323e4f0b7e41ed3ab5bac567fd056d418be98ac15f29e2): complete subsection reference.

- [nodes](resources--nfv_service--reference--group-002.md#canonical-339df16e21e9e29d792fcf340e9c363ec54743f68cfd223b9539dd23fe1fa661): complete subsection reference.

<a id="canonical-9ae902042fb560fa50f0798c4dd202800bc16c99ec7ed680c9208cd3c2dfb405"></a>

<a id="canonical-5989ff1bc13c16a51e41649ebd04a4e11e24206015cbe84262a3e0f921bc10ce"></a>

## ssh_key property — f5_big_ip_aws_service / d0432981823c / 5

Type: `"string"`. Optional.

Public SSH key for accessing the Big IP nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-b06a7adcba66eb225949433b2f299bb9d8f06a97a1ceb97ca087a0c47ab946c9"></a>

<a id="canonical-5abf72ea2b4eee69b8c26658ee1a751acc1d3d68f28a98ba194b5c4876114748"></a>

## tags property — f5_big_ip_aws_service / d0432981823c / 6

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

<a id="canonical-7ed08e867b1a7dd28b5de329667b69aadd3c7b44bfb058489aae54445ba43124"></a>

## Next pages — f5_big_ip_aws_service / d0432981823c / 7

- [f5_big_ip_aws_service.admin_password](resources--nfv_service--reference--group-001.md#canonical-0421e522b2c40b190f687c7bdd92b8c14169b71f32b7559395cc37737a1679dc)
- [f5_big_ip_aws_service.aws_tgw_site_params](resources--nfv_service--reference--group-001.md#canonical-afd6a080907c4e229ae43f4329cccdd21898342797d2a034dc6c3ae64906a6a4)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- [f5_big_ip_aws_service.market_place_image](resources--nfv_service--reference--group-001.md#canonical-d9a50ca9b65a69d0b0323e4f0b7e41ed3ab5bac567fd056d418be98ac15f29e2)
- [f5_big_ip_aws_service.nodes](resources--nfv_service--reference--group-002.md#canonical-339df16e21e9e29d792fcf340e9c363ec54743f68cfd223b9539dd23fe1fa661)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-0421e522b2c40b190f687c7bdd92b8c14169b71f32b7559395cc37737a1679dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ceab21029ee302589e9036471cd895e6466bff74a8810422209835b992e13ca"></a>

## f5_big_ip_aws_service.admin_password — f5_big_ip_aws_service.admin_password / 38bf7b07708f / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- f5_big_ip_aws_service.admin_password

<a id="canonical-08fa7f6691a357f79a9846f12f4c6a5ae6484f9cd718f0ae9541b6344b5a3cd6"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
admin_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-f1eae3bbb7fd1ad6d1312ab787f2f278fc19745281d8d604a0d5ab910db29ad5"></a>

## Direct properties — f5_big_ip_aws_service.admin_password / 38bf7b07708f / 3

- [blindfold_secret_info](resources--nfv_service--reference--group-001.md#canonical-bd03e87df1e739d5e1e74cb81d09ccc3d19194f2e7ead8f9f5986805161f31a7): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-001.md#canonical-a89aed0dcb466843fa886cc3725cd88a3e36ab3e721021fc2ff8634872151102): complete subsection reference.

<a id="canonical-370dd9731fcf4214c5601b30efc5f49ccf4ef8a145649da02483bd731935d0a3"></a>

## Next pages — f5_big_ip_aws_service.admin_password / 38bf7b07708f / 4

- [f5_big_ip_aws_service.admin_password.blindfold_secret_info](resources--nfv_service--reference--group-001.md#canonical-bd03e87df1e739d5e1e74cb81d09ccc3d19194f2e7ead8f9f5986805161f31a7)
- [f5_big_ip_aws_service.admin_password.clear_secret_info](resources--nfv_service--reference--group-001.md#canonical-a89aed0dcb466843fa886cc3725cd88a3e36ab3e721021fc2ff8634872151102)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-bd03e87df1e739d5e1e74cb81d09ccc3d19194f2e7ead8f9f5986805161f31a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2e7d6c88d335b085f39142a0f853a8e6e76e710ee6cb6ebbbe66e5e2c4ee7b4"></a>

## f5_big_ip_aws_service.admin_password.blindfold_secret_info — f5_big_ip_aws_service.admin_password.blindfold_secret_info / ece21f325e16 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.admin_password](resources--nfv_service--reference--group-001.md#canonical-0421e522b2c40b190f687c7bdd92b8c14169b71f32b7559395cc37737a1679dc)
- f5_big_ip_aws_service.admin_password.blindfold_secret_info

<a id="canonical-f12ac6666dc171212d4f8284a74bc05dd88ca8b1062eee71fddc9f9ce870b230"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-8d843c26acf1af8aed149d11464c4ab3e8caf6b1eb50bfa109de303ff003092f"></a>

## Direct properties — f5_big_ip_aws_service.admin_password.blindfold_secret_info / ece21f325e16 / 3

<a id="canonical-b9c3e4ccbb56ec57a20e80d9326566921777ae110a4f6a3f6c5edc9f288f053c"></a>

<a id="canonical-88d5138cdc247a75f57291da564bfb0d809ac95ffc4702ab0847ce2da52eaf7e"></a>

## decryption_provider property — f5_big_ip_aws_service.admin_password.blindfold_secret_info / ece21f325e16 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-f5cee47bdce04278fe1411a61d03b8076840d291ac108717c38e38a8b8b38ec1"></a>

<a id="canonical-50323af5627bc1e398766466498f6b12d42d53f0ef77d9ced9c11376514e0c29"></a>

## location property — f5_big_ip_aws_service.admin_password.blindfold_secret_info / ece21f325e16 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-b44f6ab2de7e1f53796f00a1e6df8667cb2d596861a2c529f3bdafdcc043a05d"></a>

<a id="canonical-dae9279e4bff05445db0d909972ba053b58f607c67c01177208621fd89979f51"></a>

## store_provider property — f5_big_ip_aws_service.admin_password.blindfold_secret_info / ece21f325e16 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-3a9f1650f1dc456084d0995c299a00e02cf237c61c86402468940acd7a43ee40"></a>

## Next pages — f5_big_ip_aws_service.admin_password.blindfold_secret_info / ece21f325e16 / 7

- [f5_big_ip_aws_service.admin_password](resources--nfv_service--reference--group-001.md#canonical-0421e522b2c40b190f687c7bdd92b8c14169b71f32b7559395cc37737a1679dc)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-a89aed0dcb466843fa886cc3725cd88a3e36ab3e721021fc2ff8634872151102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7eb2ee39b2c4a0f82c07435f73930424518019ab60508b160f6ac3e1dac6c020"></a>

## f5_big_ip_aws_service.admin_password.clear_secret_info — f5_big_ip_aws_service.admin_password.clear_secret_info / 67967f70815b / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.admin_password](resources--nfv_service--reference--group-001.md#canonical-0421e522b2c40b190f687c7bdd92b8c14169b71f32b7559395cc37737a1679dc)
- f5_big_ip_aws_service.admin_password.clear_secret_info

<a id="canonical-f7466a1404176aa1ae4d2537bce7ae981bfee603126afda95f8e1624e37a3672"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-09855f06b588f79d5d4c07b5ed2f8c57c8b86e318624fe10cb262473b04d4e25"></a>

## Direct properties — f5_big_ip_aws_service.admin_password.clear_secret_info / 67967f70815b / 3

<a id="canonical-bb8a031940fdcf901a33dbd85deeace8fae5c13ad2fabc8cde163df6738a8b61"></a>

<a id="canonical-fb044e6eb9ae0d01835c76d87ab407d9b07a61f7c56c803972e17e41b01c41f9"></a>

## provider_ref property — f5_big_ip_aws_service.admin_password.clear_secret_info / 67967f70815b / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-4f9ff36ccbe1ca15240af9071e2e4dbe9d353aa81f2ba709079b221805fc1530"></a>

<a id="canonical-5074c9dfd2f2b9e037372f54d93a3af5847c1dfd550ecafc7bc8edbb7868e3cd"></a>

## url property — f5_big_ip_aws_service.admin_password.clear_secret_info / 67967f70815b / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-01d84fcb41c02454581a2dc124a11a7d4e25fa9fe74886ea2e671c358585fbe2"></a>

## Next pages — f5_big_ip_aws_service.admin_password.clear_secret_info / 67967f70815b / 6

- [f5_big_ip_aws_service.admin_password](resources--nfv_service--reference--group-001.md#canonical-0421e522b2c40b190f687c7bdd92b8c14169b71f32b7559395cc37737a1679dc)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-afd6a080907c4e229ae43f4329cccdd21898342797d2a034dc6c3ae64906a6a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee581fd8ff3ec70170a3fc13522bebe144a7c7c686f5be47bf5915e300057948"></a>

## f5_big_ip_aws_service.aws_tgw_site_params — f5_big_ip_aws_service.aws_tgw_site_params / c249cc539dd9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- f5_big_ip_aws_service.aws_tgw_site_params

<a id="canonical-f8dd381e131179115134244684f4b289bf93554a0c357026c16c7f2dbd36347f"></a>

Type: `"object"`. single nested block, Optional.

BIG-IP AWS TGW Site. BIG-IP AWS TGW site specification.

Upstream description:

BIG-IP AWS TGW site specification.

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
aws_tgw_site_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-39ae4b9f22268374d212878d74b637467576d38fbf348da369e28091c713a589"></a>

## Direct properties — f5_big_ip_aws_service.aws_tgw_site_params / c249cc539dd9 / 3

- [aws_tgw_site](resources--nfv_service--reference--group-001.md#canonical-704a362e2eb9a17771204134be82c6397d010ba60dded3c6a9792aeff75243bb): complete subsection reference.

<a id="canonical-b353d2dc565317cad923ce90c4e12d93eb1f007b7beaf73a2ef1fec20a1895fe"></a>

## Next pages — f5_big_ip_aws_service.aws_tgw_site_params / c249cc539dd9 / 4

- [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site](resources--nfv_service--reference--group-001.md#canonical-704a362e2eb9a17771204134be82c6397d010ba60dded3c6a9792aeff75243bb)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-704a362e2eb9a17771204134be82c6397d010ba60dded3c6a9792aeff75243bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55943e7f744a01ab9755072a7b93e123615bb8919f129fe811e7777bde2698d6"></a>

## f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site — f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site / 5f862ecce244 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.aws_tgw_site_params](resources--nfv_service--reference--group-001.md#canonical-afd6a080907c4e229ae43f4329cccdd21898342797d2a034dc6c3ae64906a6a4)
- f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site

<a id="canonical-03399a97055910b2b1ce7fb514a52ecf4f294f796ef9f94857db9f6496450228"></a>

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
aws_tgw_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-33ca4e0c01a191ef33c4c3bbc999048f2090b659efa5b49f1640fee6764c73a7"></a>

## Direct properties — f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site / 5f862ecce244 / 3

<a id="canonical-44f67026a93a2850486be21a79d99d09eb74d7488b60a2dcebc78ed30c3245b6"></a>

<a id="canonical-bbbb3439eb5b61edd02de98684b144af81fa34e3a17a1058ab2b2b54d638f81e"></a>

## name property — f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site / 5f862ecce244 / 4

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

<a id="canonical-35acce15bda479152a832a1266dd1ca9049adcd01a937a288ac685d8493b87e6"></a>

<a id="canonical-eba7b84bc0444911361a4b23ba0e3a271ad0511b5fcdd7dba832afed53840839"></a>

## namespace property — f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site / 5f862ecce244 / 5

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

<a id="canonical-9cfc4db49ab7d2c763100b88c85fdd38f204201c7d7a5553cfe6b7f69577eaaa"></a>

<a id="canonical-e7c5ab0c30d4135d178f7dad507e2f49d62318dbcd9fd09236ec5b597b24ee21"></a>

## tenant property — f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site / 5f862ecce244 / 6

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

<a id="canonical-9672efb915c0bfe07742ea6744794c06a92f613172a14c2e8004131c4788a9f5"></a>

## Next pages — f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site / 5f862ecce244 / 7

- [f5_big_ip_aws_service.aws_tgw_site_params](resources--nfv_service--reference--group-001.md#canonical-afd6a080907c4e229ae43f4329cccdd21898342797d2a034dc6c3ae64906a6a4)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c68e5cba9dd906660b6fb6f8d93e8fafd81a6a4e25225f7cdd21e3f4baca968"></a>

## f5_big_ip_aws_service.endpoint_service — f5_big_ip_aws_service.endpoint_service / 2c697c8b292a / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- f5_big_ip_aws_service.endpoint_service

<a id="canonical-4ee6168cd62d41c8989922d8712c7d377b7718e86dd78a73e345093d82ecb417"></a>

Type: `"object"`. single nested block, Optional.

Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies
the destination with a new destination address.

Upstream description:

Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies
the destination with a new destination address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_on_slo_ip",
    "advertise_on_slo_ip_external"),
  validators.ConflictingObjectAttributes("advertise_on_slo_ip",
    "disable_advertise_on_slo_ip"),
  validators.ConflictingObjectAttributes("advertise_on_slo_ip_external",
    "disable_advertise_on_slo_ip"),
  validators.ConflictingObjectAttributes("automatic_vip",
    "configured_vip"),
  validators.ConflictingObjectAttributes("custom_tcp_ports",
    "default_tcp_ports"),
  validators.ConflictingObjectAttributes("custom_tcp_ports",
    "http_port"),
  validators.ConflictingObjectAttributes("custom_tcp_ports",
    "https_port"),
  validators.ConflictingObjectAttributes("custom_tcp_ports",
    "no_tcp_ports"),
  validators.ConflictingObjectAttributes("custom_udp_ports",
    "no_udp_ports"),
  validators.ConflictingObjectAttributes("default_tcp_ports",
    "http_port"),
  validators.ConflictingObjectAttributes("default_tcp_ports",
    "https_port"),
  validators.ConflictingObjectAttributes("default_tcp_ports",
    "no_tcp_ports"),
  validators.ConflictingObjectAttributes("http_port",
    "https_port"),
  validators.ConflictingObjectAttributes("http_port",
    "no_tcp_ports"),
  validators.ConflictingObjectAttributes("https_port",
    "no_tcp_ports")}
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
  "x-ves-oneof-field-external_vip_choice": "[\"advertise_on_slo_ip\",\"advertise_on_slo_ip_external\",\"disable_advertise_on_slo_ip\"]",
  "x-ves-oneof-field-inside_vip_choice": "[\"automatic_vip\",\"configured_vip\"]",
  "x-ves-oneof-field-tcp_port_choice": "[\"custom_tcp_ports\",\"default_tcp_ports\",\"http_port\",\"https_port\",\"no_tcp_ports\"]",
  "x-ves-oneof-field-udp_port_choice": "[\"custom_udp_ports\",\"no_udp_ports\"]"
}
```

Terraform syntax:

```terraform
endpoint_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-d7c53ba3cfcdb4b90caa81852f40adf19778060f989ae15612dba36c26e404be"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service / 2c697c8b292a / 3

- [advertise_on_slo_ip](resources--nfv_service--reference--group-001.md#canonical-822e0705e9350e7b4b3fd6b7ec9c448e99958f3893cd37a0fbb80d784ae5cac6): complete subsection reference.

- [advertise_on_slo_ip_external](resources--nfv_service--reference--group-001.md#canonical-7cc1ff84a786622ea1bf0102a08d40b25afd6cfa5454775a4cb600e98ab25e22): complete subsection reference.

- [automatic_vip](resources--nfv_service--reference--group-001.md#canonical-27a01dccff36d8fbb82246f25bb493126b9f608b90ff39a4ea59cf320b8a4c04): complete subsection reference.

<a id="canonical-82d18cdb1aad812c7dcca9dc550071f56d5ba05bb0a00fd545f5c3527bc75beb"></a>

<a id="canonical-ea1dedc18851838d1bdbaf5c69e46780b98614c62af52c735af580adce2ba094"></a>

## configured_vip property — f5_big_ip_aws_service.endpoint_service / 2c697c8b292a / 4

Type: `"string"`. Optional.

Exclusive with \[automatic\_vip\] Enter IP address for the default VIP.

Upstream description:

Exclusive with \[automatic\_vip\] Enter IP address for the default VIP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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
    "ves.io.schema.rules.string.ip": "true",
    "ves.io.schema.rules.string.not_in": "192.0.2.26"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true",
    "ves.io.schema.rules.string.not_in": "192.0.2.26"
  }
}
```

- [custom_tcp_ports](resources--nfv_service--reference--group-001.md#canonical-88f9a9242489460fbb1355c2c88b20e8dd53fb778c8bbf62a1e268fe098daee7): complete subsection reference.

- [custom_udp_ports](resources--nfv_service--reference--group-001.md#canonical-d9c5e16550833222b6a73d93e5b35f0be437151d0b3119f692ae9b4751e969bd): complete subsection reference.

- [default_tcp_ports](resources--nfv_service--reference--group-001.md#canonical-28381722993e44be09910240f91a27423fffbe33f5367048cf51d646bd83eb17): complete subsection reference.

- [disable_advertise_on_slo_ip](resources--nfv_service--reference--group-001.md#canonical-5f1862625cf922d502e1989dcb6eb51ba9e9984a03413cfd0c96e314b7d94b3d): complete subsection reference.

- [http_port](resources--nfv_service--reference--group-001.md#canonical-f8af9d4bfc3bf9f3ad878ac1954454be1d8d28e02904e4f416dd76089802d0dd): complete subsection reference.

- [https_port](resources--nfv_service--reference--group-001.md#canonical-57a7a0057734608f9bcea3dd4b451f2f6a2112db9c192cb732e8b387d46bd15c): complete subsection reference.

- [no_tcp_ports](resources--nfv_service--reference--group-001.md#canonical-bbb3ace79f2a1d7f074fc3cfdac4c4ddb3eaf53cc25ec5cf2cf768a66ec63ddf): complete subsection reference.

- [no_udp_ports](resources--nfv_service--reference--group-001.md#canonical-0e95a308c3ee609f980f014546ac43074ac905defcb4d9ddf82b21c414072501): complete subsection reference.

<a id="canonical-aca5d3dc13648f5ce1221ba8f54657c5241ad8439a0babc91ea2b77b7e5a470d"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service / 2c697c8b292a / 5

- [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip](resources--nfv_service--reference--group-001.md#canonical-822e0705e9350e7b4b3fd6b7ec9c448e99958f3893cd37a0fbb80d784ae5cac6)
- [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external](resources--nfv_service--reference--group-001.md#canonical-7cc1ff84a786622ea1bf0102a08d40b25afd6cfa5454775a4cb600e98ab25e22)
- [f5_big_ip_aws_service.endpoint_service.automatic_vip](resources--nfv_service--reference--group-001.md#canonical-27a01dccff36d8fbb82246f25bb493126b9f608b90ff39a4ea59cf320b8a4c04)
- [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports](resources--nfv_service--reference--group-001.md#canonical-88f9a9242489460fbb1355c2c88b20e8dd53fb778c8bbf62a1e268fe098daee7)
- [f5_big_ip_aws_service.endpoint_service.custom_udp_ports](resources--nfv_service--reference--group-001.md#canonical-d9c5e16550833222b6a73d93e5b35f0be437151d0b3119f692ae9b4751e969bd)
- [f5_big_ip_aws_service.endpoint_service.default_tcp_ports](resources--nfv_service--reference--group-001.md#canonical-28381722993e44be09910240f91a27423fffbe33f5367048cf51d646bd83eb17)
- [f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip](resources--nfv_service--reference--group-001.md#canonical-5f1862625cf922d502e1989dcb6eb51ba9e9984a03413cfd0c96e314b7d94b3d)
- [f5_big_ip_aws_service.endpoint_service.http_port](resources--nfv_service--reference--group-001.md#canonical-f8af9d4bfc3bf9f3ad878ac1954454be1d8d28e02904e4f416dd76089802d0dd)
- [f5_big_ip_aws_service.endpoint_service.https_port](resources--nfv_service--reference--group-001.md#canonical-57a7a0057734608f9bcea3dd4b451f2f6a2112db9c192cb732e8b387d46bd15c)
- [f5_big_ip_aws_service.endpoint_service.no_tcp_ports](resources--nfv_service--reference--group-001.md#canonical-bbb3ace79f2a1d7f074fc3cfdac4c4ddb3eaf53cc25ec5cf2cf768a66ec63ddf)
- [f5_big_ip_aws_service.endpoint_service.no_udp_ports](resources--nfv_service--reference--group-001.md#canonical-0e95a308c3ee609f980f014546ac43074ac905defcb4d9ddf82b21c414072501)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-822e0705e9350e7b4b3fd6b7ec9c448e99958f3893cd37a0fbb80d784ae5cac6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9bb1010b936a79db3f051fa646420b2300b7f3457ec6144d7fa1e8dde2f34f0"></a>

## f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip — f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip / 50dafff6bc47 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip

<a id="canonical-89208ab8c6ba8968e2fc4959cd55699e86567a7196e3e445ad9d654908a4f0c3"></a>

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
advertise_on_slo_ip = {}
```

<a id="canonical-377423fc72caded54693ee3349ddf7cd0ce1ed309fcbff405f8b0aa469c81bbe"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip / 50dafff6bc47 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cf180638da43ec6e58d98b971cecc3812d68b11694d43d6b897985860f7f7f29"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip / 50dafff6bc47 / 4

- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-7cc1ff84a786622ea1bf0102a08d40b25afd6cfa5454775a4cb600e98ab25e22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b61870f23beeb218eded5b4f48e7d8830823b30741918041f0fa2e12c9262d7e"></a>

## f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external — f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external / ab4fd793c71f / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external

<a id="canonical-f4e65b9f6bd44d3243d0a8ca7e0b08f2fcb9076b9c7a1a31e38bbb239b4a47ee"></a>

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
advertise_on_slo_ip_external = {}
```

<a id="canonical-59291fe72bfe99516e38bd1fe450116bdc86cc5b9ce4b9e1ba12d745b9c1b6f5"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external / ab4fd793c71f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-865b64d2634b44aa70000a63fe5a4589675bd9e6a767f3eadabb9492cc10c870"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external / ab4fd793c71f / 4

- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-27a01dccff36d8fbb82246f25bb493126b9f608b90ff39a4ea59cf320b8a4c04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa37077ec7bdc828163e367509e02a7b2cd2273e93390ae1b8a4ac7201a719dd"></a>

## f5_big_ip_aws_service.endpoint_service.automatic_vip — f5_big_ip_aws_service.endpoint_service.automatic_vip / 9d38eb794367 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- f5_big_ip_aws_service.endpoint_service.automatic_vip

<a id="canonical-3697b50ff3fa370d2f805177355f0a925e107688ed2b882e080c9e76f1b76795"></a>

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
automatic_vip = {}
```

<a id="canonical-7e440f53c15fde58b197690f1ff4552adb7c38a89b406fbd6f02e2f7a1ff1b0c"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.automatic_vip / 9d38eb794367 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-db72de917883f38024e3194d31daba597827814157df4b94acd37b226e8e0ec9"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.automatic_vip / 9d38eb794367 / 4

- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-88f9a9242489460fbb1355c2c88b20e8dd53fb778c8bbf62a1e268fe098daee7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13e95357955d869b1fc517a86704eb43419b4bc950816bf81ca1e5f4ad47944a"></a>

## f5_big_ip_aws_service.endpoint_service.custom_tcp_ports — f5_big_ip_aws_service.endpoint_service.custom_tcp_ports / 402888651a15 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- f5_big_ip_aws_service.endpoint_service.custom_tcp_ports

<a id="canonical-681f58eafa7a57e8eabe9fd3da7b0c67125eb3ecb3f1d891aadbea94f6f65a62"></a>

Type: `"object"`. single nested block, Optional.

Port Range List. List of port ranges.

Upstream description:

List of port ranges.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
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
custom_tcp_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-ec0ea77097b7a2f94ca6a0e982fface532473076b5430ea6dcd3f0dca3b6cdbe"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.custom_tcp_ports / 402888651a15 / 3

<a id="canonical-006120b2a17c6bdb6ca1ac726c0cd1f8152f482b5183ed3f8d890776c05ef066"></a>

<a id="canonical-2c616855f800955d335dd9ede9e00d173cce97704523f5ac803b949b40318526"></a>

## ports property — f5_big_ip_aws_service.endpoint_service.custom_tcp_ports / 402888651a15 / 4

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-82f88f1eaba101d6ee5e6ba9d7f7c8876080ec84e847bccaee715b6f6f1664c3"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.custom_tcp_ports / 402888651a15 / 5

- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-d9c5e16550833222b6a73d93e5b35f0be437151d0b3119f692ae9b4751e969bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2d7840a74c7d3f5a2f6e860bf9210da61a5fc9638a3ada4d2e18156b1e63f35"></a>

## f5_big_ip_aws_service.endpoint_service.custom_udp_ports — f5_big_ip_aws_service.endpoint_service.custom_udp_ports / 3365cad868c9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- f5_big_ip_aws_service.endpoint_service.custom_udp_ports

<a id="canonical-b0fffef94c27f48543349cc860edd0dac5ee8d138dae156e7c507f9fd8f6a173"></a>

Type: `"object"`. single nested block, Optional.

Port Range List. List of port ranges.

Upstream description:

List of port ranges.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
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
custom_udp_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-bec54b57448262958e58ad4d620d7bdd56781e650f171b53bd834f0391c937ae"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.custom_udp_ports / 3365cad868c9 / 3

<a id="canonical-46251af9ebd0a00ace50606cbfb6e25afd9f27d2b2d62370274f3ded7c2d0a8e"></a>

<a id="canonical-f31c9f187d22e54b266d24669c17ebefc16aa9314f3b35ad490e45180753de8a"></a>

## ports property — f5_big_ip_aws_service.endpoint_service.custom_udp_ports / 3365cad868c9 / 4

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-0ead365029f7b377c7082f3faa796734197971910766da0ff91498f322a5bac2"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.custom_udp_ports / 3365cad868c9 / 5

- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-28381722993e44be09910240f91a27423fffbe33f5367048cf51d646bd83eb17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1d7ad20ddab6a3ed755bc604f0ed41155b24291cb9a2006036f7c469e36fe03"></a>

## f5_big_ip_aws_service.endpoint_service.default_tcp_ports — f5_big_ip_aws_service.endpoint_service.default_tcp_ports / 8dc456c5d74b / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- f5_big_ip_aws_service.endpoint_service.default_tcp_ports

<a id="canonical-8c2e4d9781ef667f06a0f82bbf70b502ee0576a95907ac428cd673f7299f1ca8"></a>

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
default_tcp_ports = {}
```

<a id="canonical-7788cfd060b39a32cfae818138cf865437a9981dcae85c59b4d335a9e129730d"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.default_tcp_ports / 8dc456c5d74b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a8a0492acf52473b4060777d28f5a9079fc84484d989640d0ca8c94b0facf099"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.default_tcp_ports / 8dc456c5d74b / 4

- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-5f1862625cf922d502e1989dcb6eb51ba9e9984a03413cfd0c96e314b7d94b3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58909f426db1674531f8bd3786f440e924ce83d7d650e318f1714f98130c75a5"></a>

## f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip — f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip / fa06b644b87b / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip

<a id="canonical-d2ad5e9d5c6702de80eae4395d6d79d932fc91c2062dc7cbf9a657f1d160d22c"></a>

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
disable_advertise_on_slo_ip = {}
```

<a id="canonical-c846c4c88b4dfad182dfb2d1fba61c2f57e1a871d323b1c95029dfb5d56a59dd"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip / fa06b644b87b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-884a486cd29de25485b28be90f8a69a2fb1028faa9ea15a9bddb9ae55c855738"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip / fa06b644b87b / 4

- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-f8af9d4bfc3bf9f3ad878ac1954454be1d8d28e02904e4f416dd76089802d0dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf30b92806c4d73822b50853438d0a6439fd971470a6b3bd7ed3b16308042051"></a>

## f5_big_ip_aws_service.endpoint_service.http_port — f5_big_ip_aws_service.endpoint_service.http_port / b954c20cc31f / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- f5_big_ip_aws_service.endpoint_service.http_port

<a id="canonical-1a1270bf9d72506f82d42b2bfa9a191f82194b714b440b6bbd7724a22c291983"></a>

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
http_port = {}
```

<a id="canonical-e2810d2a9938c4a150e02d4ba5b0bc8f07a26c76fd1c32aee96848325a58dc64"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.http_port / b954c20cc31f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-224e5e3324213f3e6587f069884f27ac8bed21cefa509fde27a8a35b3046a55c"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.http_port / b954c20cc31f / 4

- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-57a7a0057734608f9bcea3dd4b451f2f6a2112db9c192cb732e8b387d46bd15c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60a6b9118d18c914efed49ccd27cb0532d7d7f1571c768f62084dc7acc172b14"></a>

## f5_big_ip_aws_service.endpoint_service.https_port — f5_big_ip_aws_service.endpoint_service.https_port / 0f45161c5155 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- f5_big_ip_aws_service.endpoint_service.https_port

<a id="canonical-e7802b3e5600fab499a0850d37fdc9ea78745ae1572870168ea4ea10baecd41e"></a>

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
https_port = {}
```

<a id="canonical-41ef7353c95947b368a18fc859c64e2e56d4de7678cee82f39f31150c16aea6f"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.https_port / 0f45161c5155 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d89f1eeaf99d6c5cb7c44052d090061b20c487da2a27cb110f76f88b95b156da"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.https_port / 0f45161c5155 / 4

- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-bbb3ace79f2a1d7f074fc3cfdac4c4ddb3eaf53cc25ec5cf2cf768a66ec63ddf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe96a90f41ba5c4501bb7f44941a3e67fa302be183887872dd4973b6c97b8843"></a>

## f5_big_ip_aws_service.endpoint_service.no_tcp_ports — f5_big_ip_aws_service.endpoint_service.no_tcp_ports / d4a1573ae5d7 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- f5_big_ip_aws_service.endpoint_service.no_tcp_ports

<a id="canonical-0d8d7df88027ebd71f68e4c209688575de358e698d7ab4c637f86566ce6a5c06"></a>

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
no_tcp_ports = {}
```

<a id="canonical-8ccc92b01d7e284d22ad0baf1ad06bbbe371081d58a5e879fc4d6c29e8c26c92"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.no_tcp_ports / d4a1573ae5d7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-beb7cec734f9f57430a79e8ac9e263ac523345a24ee828edc2c5e983ce3d5051"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.no_tcp_ports / d4a1573ae5d7 / 4

- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-0e95a308c3ee609f980f014546ac43074ac905defcb4d9ddf82b21c414072501"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-157448b2f28a5bbf56c484728c14e0f5b131c1aa453774ef571c895f0de5e401"></a>

## f5_big_ip_aws_service.endpoint_service.no_udp_ports — f5_big_ip_aws_service.endpoint_service.no_udp_ports / 858ee511c8ce / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- f5_big_ip_aws_service.endpoint_service.no_udp_ports

<a id="canonical-b257df1ad829e4a78623d78dd1dcecf9f090059a786e6de41ea4ed8fc628c91d"></a>

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
no_udp_ports = {}
```

<a id="canonical-6a7c70dc98063b372d4e6ca6277c5f946271665ea025b50946c2f15708ed475b"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.no_udp_ports / 858ee511c8ce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3668f455d3f36e697133b02b7465f2e0f44233c31b55be5d53d8e62683fbdff2"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.no_udp_ports / 858ee511c8ce / 4

- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-443a0926b767e5b414eeac67a51b203bb5a8cc161a2c7dc2eb5e9bf2a1efdf50)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-d9a50ca9b65a69d0b0323e4f0b7e41ed3ab5bac567fd056d418be98ac15f29e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
