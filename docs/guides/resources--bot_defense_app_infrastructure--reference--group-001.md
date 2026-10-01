---
page_title: "xcsh_bot_defense_app_infrastructure reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_defense_app_infrastructure reference."
---

# xcsh_bot_defense_app_infrastructure reference

<a id="canonical-ab45a834092534caee16f8d6620b6949464c5f1b4a9001ce716fc0b8c3b7aec4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99a1b921f6d8b0535b4e8280eb9b5e6ac8d296d1e6065e45021904943a6b4fba"></a>

## Property reference — Property reference / 18c3965a2f2c / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)
- Property reference

<a id="canonical-0d7643a647ef24868716c22cc8bc500af6d33262df956c077ff89bbbd8d6c20c"></a>

## Direct properties — Property reference / 18c3965a2f2c / 3

<a id="canonical-ec918dde478d463f116c732bc4b4763eb120d03241e689f4eef718da20a2e8bc"></a>

<a id="canonical-60894d24380728b828e6e3759f0cc6f0f417bf79be98e20e1d485be67182138d"></a>

## annotations property — Property reference / 18c3965a2f2c / 4

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

- [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2302f8dca87d727516af79b1026bae4b4453188e8c9eba021d245494dafe317f): complete subsection reference.

- [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-5f420df28410131e89e452d98c433c00b3407ad146136ae6f8e7ce1d9eb0d1c1): complete subsection reference.

<a id="canonical-b7652231665e5aa0e4a0ceda61a51e5312c67a7f3af5df766c555a3c09f8dc68"></a>

<a id="canonical-d711b38ceeb7dad5e84c333cb5b51b1e1e2cdcc5dc649a3393243e668cf4881c"></a>

## description property — Property reference / 18c3965a2f2c / 5

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

<a id="canonical-83fd01625f2408cd4ef8e9f3945b1c34f67e5c438e7765f531c3a6a4e1cecce3"></a>

<a id="canonical-cb021420c35abf333787073d1dbe0445c2880a91d887bc896bc3e9bdea34ee6c"></a>

## disable property — Property reference / 18c3965a2f2c / 6

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

<a id="canonical-7e1deeb240440c3be18a1cb8005f313a80b418156e6b4922ecfb2c60c35a152c"></a>

<a id="canonical-25c2546b44e7a131430c2275e29e0eff98b7f392efa79a0b4a959c13c15cc81f"></a>

## environment_type property — Property reference / 18c3965a2f2c / 7

Type: `"string"`. Optional, Computed.

\[Enum: PRODUCTION|TESTING\] Environment Type Production environment Testing environment. Possible
values are \`PRODUCTION\`, \`TESTING\`. Defaults to \`PRODUCTION\`.

Upstream description:

Environment Type

Production environment Testing environment.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PRODUCTION",
    "TESTING"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PRODUCTION",
  "enum": [
    "PRODUCTION",
    "TESTING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-69654f098352c9e592091808792d857e1fbd842249f22ac5bbcb979b3cfcc3fb"></a>

<a id="canonical-e135a9cb5019642336e3817f2335bcaf924ecc1cfd3defb08d8659ad155f32eb"></a>

## id property — Property reference / 18c3965a2f2c / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-5a5e9f8a20f88a397a08ef1e62e32997493ccb2f6302ab654f35b0481ae62555"></a>

<a id="canonical-e0efe097eff0d92c835a16b4a27d3c001eed6176f202e69050ed9418547007e5"></a>

## labels property — Property reference / 18c3965a2f2c / 9

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

<a id="canonical-4f5cec617145da80293b73f6e2f81fe213ece65e74d02bc77954a4081cdcedf4"></a>

<a id="canonical-5c82cd0658d32286e910ad11a78da57690e2eaeb86f08c099a08b20274ee0827"></a>

## name property — Property reference / 18c3965a2f2c / 10

Type: `"string"`. Required.

Name of the Bot Defense App Infrastructure. Must be unique within the namespace.

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

<a id="canonical-c1e7d5eb32b407c7e9232d831d9ad003ebc6b310d5a5e94147569a7148a203d7"></a>

<a id="canonical-c89a15995e2492b8c5963c23a09ce07dd7b470302a22576911c8e4f8477c7089"></a>

## namespace property — Property reference / 18c3965a2f2c / 11

Type: `"string"`. Required.

Namespace where the Bot Defense App Infrastructure is created.

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

- [timeouts](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-34c5003b8b2e9306176db5f499f3dd03f0d68525e71a4952c3e2b33a3e1d5d94): complete subsection reference.

<a id="canonical-ccf893e4fa9ee03878a8efa30c8cc28e9889ac5b6f6a1f70904342d0e46d4db7"></a>

<a id="canonical-ba0fe0641797e2a9ca3139457b397ff04a54ee268b31fc9b172b8102ebc97f6a"></a>

## traffic_type property — Property reference / 18c3965a2f2c / 12

Type: `"string"`. Optional, Computed.

\[Enum: WEB|MOBILE\] Traffic Type Web traffic Mobile traffic. Possible values are \`WEB\`,
\`MOBILE\`. Defaults to \`WEB\`.

Upstream description:

Traffic Type

Web traffic Mobile traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("WEB",
    "MOBILE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "WEB",
  "enum": [
    "WEB",
    "MOBILE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ab0d9a56ede3a6128cb925ac46dc0ddd70e0244f74ad335af61b90b02dd0bae7"></a>

## All schema paths — Property reference / 18c3965a2f2c / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ec918dde478d463f116c732bc4b4763eb120d03241e689f4eef718da20a2e8bc) |
| `cloud_hosted` | [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-daa30f7d259cea0b191d3445bd29c0086bc40af62b05a1c6638cae176e036e44) |
| `cloud_hosted.egress` | [cloud_hosted.egress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-72e316928363f63609d1e75d9afed23bcdaf27aeddd583e61f6d9bda0e75513b) |
| `cloud_hosted.egress.ip_address` | [cloud_hosted.egress.ip_address](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-31efebd2a7c2446a19d2c96a3d754f219b532a469866117bdd39f575365a6f9e) |
| `cloud_hosted.egress.location` | [cloud_hosted.egress.location](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-f914d849602bd05cb0411467d2bd02ecca1c7abefca9be9f7bd92f41a13b8508) |
| `cloud_hosted.infra_host_name` | [cloud_hosted.infra_host_name](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-8e5fda6db0bb990de80349897b27c76e00a392af0fd83997f43eeacbf10249a2) |
| `cloud_hosted.ingress` | [cloud_hosted.ingress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-cb2a49c42521574dd81ea73af58edc5d954531bff8133633292a7687103c142e) |
| `cloud_hosted.ingress.host_name` | [cloud_hosted.ingress.host_name](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-897c81ca51fbc6e9dfd04386b43d1162eddbcef19e2bdc7cc2193a331e8a354e) |
| `cloud_hosted.ingress.ip_address` | [cloud_hosted.ingress.ip_address](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ffa9cbc88b013d7b5d1b1dbf6f0f211186c1c7af799ce514f927526180a936c6) |
| `cloud_hosted.ingress.location` | [cloud_hosted.ingress.location](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-f4bce6935e5a583cbb854515f8cede76ff0edb30752ffe07c1ba1f5a1198f5be) |
| `cloud_hosted.region` | [cloud_hosted.region](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-41414c152a613ff270a1e631d1efeb1274ea3eff43d048dff75c25abde179217) |
| `data_center_hosted` | [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-6a747faed787c699ab86f0e7212c9830d9eb93f348bbd397c9d2ca29f026d293) |
| `data_center_hosted.egress` | [data_center_hosted.egress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-931b9432ea8a6699251cb9bd6c4035192883999e33d94595ebdb4d018f7c2221) |
| `data_center_hosted.egress.ip_address` | [data_center_hosted.egress.ip_address](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2ea54da07448187fa374bc6ab18d57ed3dc0cebdbae2121f002528479abe55e4) |
| `data_center_hosted.egress.location` | [data_center_hosted.egress.location](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-64890bf3eb90365530e00d7fdaab8677a22a20c8a875f6d5138c44524de69771) |
| `data_center_hosted.infra_host_name` | [data_center_hosted.infra_host_name](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-4116909ff6145841e4370db4e3a979e6d13773c25948b769c5037aaf07887fcc) |
| `data_center_hosted.ingress` | [data_center_hosted.ingress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-e97cd2b46e1c740e8990df3a20fe71deaf2c896e3da6f9a83ada94a57cc6f2a8) |
| `data_center_hosted.ingress.host_name` | [data_center_hosted.ingress.host_name](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-af72bca54e973e9240481b3475d450b44266415a669b69c3fc78aeacbcb8ddb9) |
| `data_center_hosted.ingress.ip_address` | [data_center_hosted.ingress.ip_address](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-9ea84e8a8ea72edfbc53fbe745efacfd9581fe7403faef1f2730aa217b9bda63) |
| `data_center_hosted.ingress.location` | [data_center_hosted.ingress.location](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2a1ff49ad0c3595843211a4bad9d453305f23bab254a8fff8610266d692b171e) |
| `data_center_hosted.region` | [data_center_hosted.region](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-9ee16dfee2d6b5b195daff51ad2c35be80f87b241d7239b587ec20517c87ea9f) |
| `description` | [description](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-b7652231665e5aa0e4a0ceda61a51e5312c67a7f3af5df766c555a3c09f8dc68) |
| `disable` | [disable](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-83fd01625f2408cd4ef8e9f3945b1c34f67e5c438e7765f531c3a6a4e1cecce3) |
| `environment_type` | [environment_type](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-7e1deeb240440c3be18a1cb8005f313a80b418156e6b4922ecfb2c60c35a152c) |
| `id` | [id](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-69654f098352c9e592091808792d857e1fbd842249f22ac5bbcb979b3cfcc3fb) |
| `labels` | [labels](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-5a5e9f8a20f88a397a08ef1e62e32997493ccb2f6302ab654f35b0481ae62555) |
| `name` | [name](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-4f5cec617145da80293b73f6e2f81fe213ece65e74d02bc77954a4081cdcedf4) |
| `namespace` | [namespace](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-c1e7d5eb32b407c7e9232d831d9ad003ebc6b310d5a5e94147569a7148a203d7) |
| `timeouts` | [timeouts](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-8252f739d135326ed1ce44595cc06f36b23f32da1c22b5adf703aff8c414ae9f) |
| `timeouts.create` | [timeouts.create](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-5a4e0695ce4933f6e1d410a2caf9fcb62f0aba13d2b4ae221b9af6aad7b6852f) |
| `timeouts.delete` | [timeouts.delete](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-9328f3f84dcd03d01a5851d4947ea277bd30170a987d71265236eac6c0bf03d7) |
| `timeouts.read` | [timeouts.read](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-67f658cccf8ef8e608e2c1a649cdb1001da7fdf8c72ae3fb6f2138529e28abb5) |
| `timeouts.update` | [timeouts.update](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-9253154cd2688c582e737e20499df11cfa05f2d45988bfbe5d4ab323e8c526cb) |
| `traffic_type` | [traffic_type](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ccf893e4fa9ee03878a8efa30c8cc28e9889ac5b6f6a1f70904342d0e46d4db7) |

<a id="canonical-dbaeccf055471ad48cfe43e08022e637b3e8fd1416211a5f05e98dea6e99b1a9"></a>

## Next pages — Property reference / 18c3965a2f2c / 14

- [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2302f8dca87d727516af79b1026bae4b4453188e8c9eba021d245494dafe317f)
- [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-5f420df28410131e89e452d98c433c00b3407ad146136ae6f8e7ce1d9eb0d1c1)
- [timeouts](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-34c5003b8b2e9306176db5f499f3dd03f0d68525e71a4952c3e2b33a3e1d5d94)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)

<a id="canonical-2302f8dca87d727516af79b1026bae4b4453188e8c9eba021d245494dafe317f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a34da9dc6ba1c9ed8a50928b2a9ddff956c0c4d7de49a140405793bf21bf0950"></a>

## cloud_hosted — cloud_hosted / d9839358a933 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ab45a834092534caee16f8d6620b6949464c5f1b4a9001ce716fc0b8c3b7aec4)
- cloud_hosted

<a id="canonical-daa30f7d259cea0b191d3445bd29c0086bc40af62b05a1c6638cae176e036e44"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: cloud\_hosted, data\_center\_hosted\] F5 Hosted. Infra F5 Hosted.

Upstream description:

Infra F5 Hosted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("egress",
    "infra_host_name",
    "ingress")}
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

OneOf alternatives in this subsection:

- [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-daa30f7d259cea0b191d3445bd29c0086bc40af62b05a1c6638cae176e036e44)
- [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-6a747faed787c699ab86f0e7212c9830d9eb93f348bbd397c9d2ca29f026d293)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
cloud_hosted {
  # Configure direct properties listed below.
}
```

<a id="canonical-14d854ce48f5c4d71cdb22194ab8ff9ba88cba8e13310ccd27b87c2bfe80d897"></a>

## Direct properties — cloud_hosted / d9839358a933 / 3

- [egress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-4e967e965b8f27fd2a38388330153dd1bb7e3bfd8ab69c2151f8e1ea0b1019d2): complete subsection reference.

<a id="canonical-8e5fda6db0bb990de80349897b27c76e00a392af0fd83997f43eeacbf10249a2"></a>

<a id="canonical-b6acec1b09951688a70df0c2e555c3b8a5428a4c553fd854dc8e243333dd0ca4"></a>

## infra_host_name property — cloud_hosted / d9839358a933 / 4

Type: `"string"`. Optional.

Infra Host Name. Infra Host Name.

Upstream description:

Infra Host Name.

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

- [ingress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-fb2e3648602aafef2becf948b35a92c5a43d6815f37a623cf47939ebc13e1bf3): complete subsection reference.

<a id="canonical-41414c152a613ff270a1e631d1efeb1274ea3eff43d048dff75c25abde179217"></a>

<a id="canonical-b24bdf0ecffeb99190ff9b5df838dfe724bbb8f4c8af6b74dc40260a0de31c74"></a>

## region property — cloud_hosted / d9839358a933 / 5

Type: `"string"`. Optional.

\[Enum: US|EU|ASIA\] Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU
European Union region - ASIA: ASIA Asia region. Possible values are \`US\`, \`EU\`, \`ASIA\`.
Defaults to \`US\`.

Upstream description:

Defines a selection for Bot Defense Advanced region

&#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("US",
    "EU",
    "ASIA"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "US",
  "enum": [
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6b0b234c2d7e06e3554824512c2ad237646fb0d340e001bd8bfa53fc8bf7e3af"></a>

## Next pages — cloud_hosted / d9839358a933 / 6

- [cloud_hosted.egress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-4e967e965b8f27fd2a38388330153dd1bb7e3bfd8ab69c2151f8e1ea0b1019d2)
- [cloud_hosted.ingress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-fb2e3648602aafef2becf948b35a92c5a43d6815f37a623cf47939ebc13e1bf3)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ab45a834092534caee16f8d6620b6949464c5f1b4a9001ce716fc0b8c3b7aec4)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)

<a id="canonical-4e967e965b8f27fd2a38388330153dd1bb7e3bfd8ab69c2151f8e1ea0b1019d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12efe073da833d61373249e8c8af6218d0e775f21a25fcf7b249f4107ed52cd9"></a>

## cloud_hosted.egress — cloud_hosted.egress / a1115d9723c4 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ab45a834092534caee16f8d6620b6949464c5f1b4a9001ce716fc0b8c3b7aec4)
- [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2302f8dca87d727516af79b1026bae4b4453188e8c9eba021d245494dafe317f)
- cloud_hosted.egress

<a id="canonical-72e316928363f63609d1e75d9afed23bcdaf27aeddd583e61f6d9bda0e75513b"></a>

Type: `"object"`. list nested block, Optional.

Egress. Egress

Upstream description:

Egress

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_address")}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
egress {
  # Configure direct properties listed below.
}
```

<a id="canonical-b535f4d506578ec6e0b38e1b38c6f738fec93fb80b691e3a0da10c96a0bd8f03"></a>

## Direct properties — cloud_hosted.egress / a1115d9723c4 / 3

<a id="canonical-31efebd2a7c2446a19d2c96a3d754f219b532a469866117bdd39f575365a6f9e"></a>

<a id="canonical-26f533ea1558ddf8ff1b1ea935d852a1cc92e6b153aad6c1c1eab07bcf50b763"></a>

## ip_address property — cloud_hosted.egress / a1115d9723c4 / 4

Type: `"string"`. Optional.

IP Address. Egress IP address.

Upstream description:

Egress IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-f914d849602bd05cb0411467d2bd02ecca1c7abefca9be9f7bd92f41a13b8508"></a>

<a id="canonical-18968c723c112d214edb73b389ca6e5e8e21185b157d9225543182b53e71d0f5"></a>

## location property — cloud_hosted.egress / a1115d9723c4 / 5

Type: `"string"`. Optional.

\[Enum:
AWS\_AP\_NORTHEAST\_1|AWS\_AP\_NORTHEAST\_3|AWS\_AP\_SOUTH\_1|AWS\_AP\_SOUTH\_2|AWS\_AP\_SOUTHEAST\_1|AWS\_AP\_SOUTHEAST\_2|AWS\_AP\_SOUTHEAST\_3|AWS\_EU\_CENTRAL\_1|AWS\_EU\_NORTH\_1|AWS\_EU\_WEST\_1|AWS\_ME\_SOUTH\_1|AWS\_SA\_EAST\_1|AWS\_US\_EAST\_1|AWS\_US\_EAST\_2|AWS\_US\_WEST\_1|AWS\_US\_WEST\_2|GCP\_ASIA\_EAST\_1|GCP\_ASIA\_EAST\_2|GCP\_ASIA\_NORTHEAST\_1|GCP\_ASIA\_NORTHEAST\_2|GCP\_ASIA\_NORTHEAST\_3|GCP\_ASIA\_SOUTH\_1|GCP\_ASIA\_SOUTHEAST\_1|GCP\_ASIA\_SOUTHEAST\_2|GCP\_AUSTRALIA\_SOUTHEAST\_1|GCP\_EUROPE\_WEST\_1|GCP\_EUROPE\_WEST\_2|GCP\_EUROPE\_WEST\_3|GCP\_NORTHAMERICA\_NORTHEAST\_1|GCP\_NORTHAMERICA\_NORTHEAST\_2|GCP\_SOUTHAMERICA\_EAST\_1|GCP\_SOUTHAMERICA\_WEST\_1|GCP\_US\_CENTRAL\_1|GCP\_US\_EAST\_1|GCP\_US\_EAST\_4|GCP\_US\_WEST\_1|GCP\_US\_WEST\_2\]
Region location AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1.. Possible values are
\`AWS\_AP\_NORTHEAST\_1\`, \`AWS\_AP\_NORTHEAST\_3\`, \`AWS\_AP\_SOUTH\_1\`, \`AWS\_AP\_SOUTH\_2\`,
\`AWS\_AP\_SOUTHEAST\_1\`, \`AWS\_AP\_SOUTHEAST\_2\`, \`AWS\_AP\_SOUTHEAST\_3\`,
\`AWS\_EU\_CENTRAL\_1\`, \`AWS\_EU\_NORTH\_1\`, \`AWS\_EU\_WEST\_1\`, \`AWS\_ME\_SOUTH\_1\`,
\`AWS\_SA\_EAST\_1\`, \`AWS\_US\_EAST\_1\`, \`AWS\_US\_EAST\_2\`, \`AWS\_US\_WEST\_1\`,
\`AWS\_US\_WEST\_2\`, \`GCP\_ASIA\_EAST\_1\`, \`GCP\_ASIA\_EAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_1\`,
\`GCP\_ASIA\_NORTHEAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_3\`, \`GCP\_ASIA\_SOUTH\_1\`,
\`GCP\_ASIA\_SOUTHEAST\_1\`, \`GCP\_ASIA\_SOUTHEAST\_2\`, \`GCP\_AUSTRALIA\_SOUTHEAST\_1\`,
\`GCP\_EUROPE\_WEST\_1\`, \`GCP\_EUROPE\_WEST\_2\`, \`GCP\_EUROPE\_WEST\_3\`,
\`GCP\_NORTHAMERICA\_NORTHEAST\_1\`, \`GCP\_NORTHAMERICA\_NORTHEAST\_2\`,
\`GCP\_SOUTHAMERICA\_EAST\_1\`, \`GCP\_SOUTHAMERICA\_WEST\_1\`, \`GCP\_US\_CENTRAL\_1\`,
\`GCP\_US\_EAST\_1\`, \`GCP\_US\_EAST\_4\`, \`GCP\_US\_WEST\_1\`, \`GCP\_US\_WEST\_2\`. Defaults to
\`AWS\_AP\_NORTHEAST\_1\`.

Upstream description:

Region location

AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1 GCP\_ASIA\_EAST\_2
GCP\_ASIA\_NORTHEAST\_1 GCP\_ASIA\_NORTHEAST\_2 GCP\_ASIA\_NORTHEAST\_3 GCP\_ASIA\_SOUTH\_1
GCP\_ASIA\_SOUTHEAST\_1 GCP\_ASIA\_SOUTHEAST\_2 GCP\_AUSTRALIA\_SOUTHEAST\_1 GCP\_EUROPE\_WEST\_1
GCP\_EUROPE\_WEST\_2 GCP\_EUROPE\_WEST\_3 GCP\_NORTHAMERICA\_NORTHEAST\_1
GCP\_NORTHAMERICA\_NORTHEAST\_2 GCP\_SOUTHAMERICA\_EAST\_1 GCP\_SOUTHAMERICA\_WEST\_1
GCP\_US\_CENTRAL\_1 GCP\_US\_EAST\_1 GCP\_US\_EAST\_4 GCP\_US\_WEST\_1 GCP\_US\_WEST\_2.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AWS_AP_NORTHEAST_1",
  "enum": [
    "AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e2d0d49a995abf51aef9811d115dccf6712c4f3eff7be761c137f79e61edbb22"></a>

## Next pages — cloud_hosted.egress / a1115d9723c4 / 6

- [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2302f8dca87d727516af79b1026bae4b4453188e8c9eba021d245494dafe317f)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)

<a id="canonical-fb2e3648602aafef2becf948b35a92c5a43d6815f37a623cf47939ebc13e1bf3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99447cf00a4f8e3f2291f7299f4b49ed528b0ad63c0b135ecb3215f8f40a2463"></a>

## cloud_hosted.ingress — cloud_hosted.ingress / e89885cb3868 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ab45a834092534caee16f8d6620b6949464c5f1b4a9001ce716fc0b8c3b7aec4)
- [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2302f8dca87d727516af79b1026bae4b4453188e8c9eba021d245494dafe317f)
- cloud_hosted.ingress

<a id="canonical-cb2a49c42521574dd81ea73af58edc5d954531bff8133633292a7687103c142e"></a>

Type: `"object"`. list nested block, Optional.

Ingress. Ingress

Upstream description:

Ingress

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("host_name",
    "ip_address")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
ingress {
  # Configure direct properties listed below.
}
```

<a id="canonical-fb319eba486a79d03bde1cc4be096147b3664ab208f489bb37d4f9da370674cd"></a>

## Direct properties — cloud_hosted.ingress / e89885cb3868 / 3

<a id="canonical-897c81ca51fbc6e9dfd04386b43d1162eddbcef19e2bdc7cc2193a331e8a354e"></a>

<a id="canonical-c0e54a34f17b548da9376fd1d0be69a6d924e7cba2f446e9a4e544d8a9de7dfd"></a>

## host_name property — cloud_hosted.ingress / e89885cb3868 / 4

Type: `"string"`. Optional.

Exclusive with \[ip\_address\] Ingress Host Name.

Upstream description:

Exclusive with \[ip\_address\] Ingress Host Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-ffa9cbc88b013d7b5d1b1dbf6f0f211186c1c7af799ce514f927526180a936c6"></a>

<a id="canonical-6975e9fe09eb0e7455c881110cb9bf23e7a92ee1532a63cf47cd5b311047ca57"></a>

## ip_address property — cloud_hosted.ingress / e89885cb3868 / 5

Type: `"string"`. Optional.

Exclusive with \[host\_name\] Ingress IP Address.

Upstream description:

Exclusive with \[host\_name\] Ingress IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-f4bce6935e5a583cbb854515f8cede76ff0edb30752ffe07c1ba1f5a1198f5be"></a>

<a id="canonical-d3022de8ac6050235dffb66527824f1298f0089d6700c66f2c1295ac016f5d9d"></a>

## location property — cloud_hosted.ingress / e89885cb3868 / 6

Type: `"string"`. Optional.

\[Enum:
AWS\_AP\_NORTHEAST\_1|AWS\_AP\_NORTHEAST\_3|AWS\_AP\_SOUTH\_1|AWS\_AP\_SOUTH\_2|AWS\_AP\_SOUTHEAST\_1|AWS\_AP\_SOUTHEAST\_2|AWS\_AP\_SOUTHEAST\_3|AWS\_EU\_CENTRAL\_1|AWS\_EU\_NORTH\_1|AWS\_EU\_WEST\_1|AWS\_ME\_SOUTH\_1|AWS\_SA\_EAST\_1|AWS\_US\_EAST\_1|AWS\_US\_EAST\_2|AWS\_US\_WEST\_1|AWS\_US\_WEST\_2|GCP\_ASIA\_EAST\_1|GCP\_ASIA\_EAST\_2|GCP\_ASIA\_NORTHEAST\_1|GCP\_ASIA\_NORTHEAST\_2|GCP\_ASIA\_NORTHEAST\_3|GCP\_ASIA\_SOUTH\_1|GCP\_ASIA\_SOUTHEAST\_1|GCP\_ASIA\_SOUTHEAST\_2|GCP\_AUSTRALIA\_SOUTHEAST\_1|GCP\_EUROPE\_WEST\_1|GCP\_EUROPE\_WEST\_2|GCP\_EUROPE\_WEST\_3|GCP\_NORTHAMERICA\_NORTHEAST\_1|GCP\_NORTHAMERICA\_NORTHEAST\_2|GCP\_SOUTHAMERICA\_EAST\_1|GCP\_SOUTHAMERICA\_WEST\_1|GCP\_US\_CENTRAL\_1|GCP\_US\_EAST\_1|GCP\_US\_EAST\_4|GCP\_US\_WEST\_1|GCP\_US\_WEST\_2\]
Region location AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1.. Possible values are
\`AWS\_AP\_NORTHEAST\_1\`, \`AWS\_AP\_NORTHEAST\_3\`, \`AWS\_AP\_SOUTH\_1\`, \`AWS\_AP\_SOUTH\_2\`,
\`AWS\_AP\_SOUTHEAST\_1\`, \`AWS\_AP\_SOUTHEAST\_2\`, \`AWS\_AP\_SOUTHEAST\_3\`,
\`AWS\_EU\_CENTRAL\_1\`, \`AWS\_EU\_NORTH\_1\`, \`AWS\_EU\_WEST\_1\`, \`AWS\_ME\_SOUTH\_1\`,
\`AWS\_SA\_EAST\_1\`, \`AWS\_US\_EAST\_1\`, \`AWS\_US\_EAST\_2\`, \`AWS\_US\_WEST\_1\`,
\`AWS\_US\_WEST\_2\`, \`GCP\_ASIA\_EAST\_1\`, \`GCP\_ASIA\_EAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_1\`,
\`GCP\_ASIA\_NORTHEAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_3\`, \`GCP\_ASIA\_SOUTH\_1\`,
\`GCP\_ASIA\_SOUTHEAST\_1\`, \`GCP\_ASIA\_SOUTHEAST\_2\`, \`GCP\_AUSTRALIA\_SOUTHEAST\_1\`,
\`GCP\_EUROPE\_WEST\_1\`, \`GCP\_EUROPE\_WEST\_2\`, \`GCP\_EUROPE\_WEST\_3\`,
\`GCP\_NORTHAMERICA\_NORTHEAST\_1\`, \`GCP\_NORTHAMERICA\_NORTHEAST\_2\`,
\`GCP\_SOUTHAMERICA\_EAST\_1\`, \`GCP\_SOUTHAMERICA\_WEST\_1\`, \`GCP\_US\_CENTRAL\_1\`,
\`GCP\_US\_EAST\_1\`, \`GCP\_US\_EAST\_4\`, \`GCP\_US\_WEST\_1\`, \`GCP\_US\_WEST\_2\`. Defaults to
\`AWS\_AP\_NORTHEAST\_1\`.

Upstream description:

Region location

AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1 GCP\_ASIA\_EAST\_2
GCP\_ASIA\_NORTHEAST\_1 GCP\_ASIA\_NORTHEAST\_2 GCP\_ASIA\_NORTHEAST\_3 GCP\_ASIA\_SOUTH\_1
GCP\_ASIA\_SOUTHEAST\_1 GCP\_ASIA\_SOUTHEAST\_2 GCP\_AUSTRALIA\_SOUTHEAST\_1 GCP\_EUROPE\_WEST\_1
GCP\_EUROPE\_WEST\_2 GCP\_EUROPE\_WEST\_3 GCP\_NORTHAMERICA\_NORTHEAST\_1
GCP\_NORTHAMERICA\_NORTHEAST\_2 GCP\_SOUTHAMERICA\_EAST\_1 GCP\_SOUTHAMERICA\_WEST\_1
GCP\_US\_CENTRAL\_1 GCP\_US\_EAST\_1 GCP\_US\_EAST\_4 GCP\_US\_WEST\_1 GCP\_US\_WEST\_2.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AWS_AP_NORTHEAST_1",
  "enum": [
    "AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-eea540fff9270de18a765944ef4799948fa671b0292d5b7e7e89e0b1aa6e057c"></a>

## Next pages — cloud_hosted.ingress / e89885cb3868 / 7

- [cloud_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2302f8dca87d727516af79b1026bae4b4453188e8c9eba021d245494dafe317f)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)

<a id="canonical-5f420df28410131e89e452d98c433c00b3407ad146136ae6f8e7ce1d9eb0d1c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57e76d70f86a3a64ea13fc38d55d3f57d10a1d6bafac8664192db7746dfd3ef7"></a>

## data_center_hosted — data_center_hosted / 8787503b68c4 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ab45a834092534caee16f8d6620b6949464c5f1b4a9001ce716fc0b8c3b7aec4)
- data_center_hosted

<a id="canonical-6a747faed787c699ab86f0e7212c9830d9eb93f348bbd397c9d2ca29f026d293"></a>

Type: `"object"`. single nested block, Optional.

F5 Hosted. Infra F5 Hosted.

Upstream description:

Infra F5 Hosted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("egress",
    "infra_host_name",
    "ingress")}
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
data_center_hosted {
  # Configure direct properties listed below.
}
```

<a id="canonical-a86d07f8c23d8664dcebfddf74f9e9a78ff4a133c5b374a38e7b24ed7f5a4f39"></a>

## Direct properties — data_center_hosted / 8787503b68c4 / 3

- [egress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3e45699c36e093c64f4d15c6041f5a774208fefbdcbff1b5b5b8a3ec023d9225): complete subsection reference.

<a id="canonical-4116909ff6145841e4370db4e3a979e6d13773c25948b769c5037aaf07887fcc"></a>

<a id="canonical-60d3ca64a8b83a2a1702ba71c9a1a53ec3734eb2ff79fa71f18bfe9d66c6dee3"></a>

## infra_host_name property — data_center_hosted / 8787503b68c4 / 4

Type: `"string"`. Optional.

Infra Host Name. Infra Host Name.

Upstream description:

Infra Host Name.

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

- [ingress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-a1ed7eb915dd7060c2443466cb6cf10027bb4d00b89b40a5d266adfd339ce47c): complete subsection reference.

<a id="canonical-9ee16dfee2d6b5b195daff51ad2c35be80f87b241d7239b587ec20517c87ea9f"></a>

<a id="canonical-dae8db1304098a4799f0bfb0c8c0a3d72f2bb0ac7e4e28519c9c51e43b8530d6"></a>

## region property — data_center_hosted / 8787503b68c4 / 5

Type: `"string"`. Optional.

\[Enum: US|EU|ASIA\] Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU
European Union region - ASIA: ASIA Asia region. Possible values are \`US\`, \`EU\`, \`ASIA\`.
Defaults to \`US\`.

Upstream description:

Defines a selection for Bot Defense Advanced region

&#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("US",
    "EU",
    "ASIA"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "US",
  "enum": [
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-48ccddad5225320f3e45277593b4546c7f94dd957bbcc34433632f5d2db5b0db"></a>

## Next pages — data_center_hosted / 8787503b68c4 / 6

- [data_center_hosted.egress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-3e45699c36e093c64f4d15c6041f5a774208fefbdcbff1b5b5b8a3ec023d9225)
- [data_center_hosted.ingress](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-a1ed7eb915dd7060c2443466cb6cf10027bb4d00b89b40a5d266adfd339ce47c)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ab45a834092534caee16f8d6620b6949464c5f1b4a9001ce716fc0b8c3b7aec4)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)

<a id="canonical-3e45699c36e093c64f4d15c6041f5a774208fefbdcbff1b5b5b8a3ec023d9225"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39413ffdf7bb8bd55315c14a4f1c455b15e6634220e97af47025c9a26af34853"></a>

## data_center_hosted.egress — data_center_hosted.egress / 11280758dde5 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ab45a834092534caee16f8d6620b6949464c5f1b4a9001ce716fc0b8c3b7aec4)
- [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-5f420df28410131e89e452d98c433c00b3407ad146136ae6f8e7ce1d9eb0d1c1)
- data_center_hosted.egress

<a id="canonical-931b9432ea8a6699251cb9bd6c4035192883999e33d94595ebdb4d018f7c2221"></a>

Type: `"object"`. list nested block, Optional.

Egress. Egress

Upstream description:

Egress

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_address")}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
egress {
  # Configure direct properties listed below.
}
```

<a id="canonical-8adbe26a37bc18980cd29c506f56fd76fc7d91c0aac3acd5f09378fc3402eca8"></a>

## Direct properties — data_center_hosted.egress / 11280758dde5 / 3

<a id="canonical-2ea54da07448187fa374bc6ab18d57ed3dc0cebdbae2121f002528479abe55e4"></a>

<a id="canonical-94385e96f0d9610314023f49e98ce3e63c83f11e524bb3d0191c23edf13eeb68"></a>

## ip_address property — data_center_hosted.egress / 11280758dde5 / 4

Type: `"string"`. Optional.

IP Address. Egress IP address.

Upstream description:

Egress IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-64890bf3eb90365530e00d7fdaab8677a22a20c8a875f6d5138c44524de69771"></a>

<a id="canonical-de9b0d5f9570079d0751381d96322b8c56974a41d53230879e19eb4c46d580f6"></a>

## location property — data_center_hosted.egress / 11280758dde5 / 5

Type: `"string"`. Optional.

\[Enum:
AWS\_AP\_NORTHEAST\_1|AWS\_AP\_NORTHEAST\_3|AWS\_AP\_SOUTH\_1|AWS\_AP\_SOUTH\_2|AWS\_AP\_SOUTHEAST\_1|AWS\_AP\_SOUTHEAST\_2|AWS\_AP\_SOUTHEAST\_3|AWS\_EU\_CENTRAL\_1|AWS\_EU\_NORTH\_1|AWS\_EU\_WEST\_1|AWS\_ME\_SOUTH\_1|AWS\_SA\_EAST\_1|AWS\_US\_EAST\_1|AWS\_US\_EAST\_2|AWS\_US\_WEST\_1|AWS\_US\_WEST\_2|GCP\_ASIA\_EAST\_1|GCP\_ASIA\_EAST\_2|GCP\_ASIA\_NORTHEAST\_1|GCP\_ASIA\_NORTHEAST\_2|GCP\_ASIA\_NORTHEAST\_3|GCP\_ASIA\_SOUTH\_1|GCP\_ASIA\_SOUTHEAST\_1|GCP\_ASIA\_SOUTHEAST\_2|GCP\_AUSTRALIA\_SOUTHEAST\_1|GCP\_EUROPE\_WEST\_1|GCP\_EUROPE\_WEST\_2|GCP\_EUROPE\_WEST\_3|GCP\_NORTHAMERICA\_NORTHEAST\_1|GCP\_NORTHAMERICA\_NORTHEAST\_2|GCP\_SOUTHAMERICA\_EAST\_1|GCP\_SOUTHAMERICA\_WEST\_1|GCP\_US\_CENTRAL\_1|GCP\_US\_EAST\_1|GCP\_US\_EAST\_4|GCP\_US\_WEST\_1|GCP\_US\_WEST\_2\]
Region location AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1.. Possible values are
\`AWS\_AP\_NORTHEAST\_1\`, \`AWS\_AP\_NORTHEAST\_3\`, \`AWS\_AP\_SOUTH\_1\`, \`AWS\_AP\_SOUTH\_2\`,
\`AWS\_AP\_SOUTHEAST\_1\`, \`AWS\_AP\_SOUTHEAST\_2\`, \`AWS\_AP\_SOUTHEAST\_3\`,
\`AWS\_EU\_CENTRAL\_1\`, \`AWS\_EU\_NORTH\_1\`, \`AWS\_EU\_WEST\_1\`, \`AWS\_ME\_SOUTH\_1\`,
\`AWS\_SA\_EAST\_1\`, \`AWS\_US\_EAST\_1\`, \`AWS\_US\_EAST\_2\`, \`AWS\_US\_WEST\_1\`,
\`AWS\_US\_WEST\_2\`, \`GCP\_ASIA\_EAST\_1\`, \`GCP\_ASIA\_EAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_1\`,
\`GCP\_ASIA\_NORTHEAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_3\`, \`GCP\_ASIA\_SOUTH\_1\`,
\`GCP\_ASIA\_SOUTHEAST\_1\`, \`GCP\_ASIA\_SOUTHEAST\_2\`, \`GCP\_AUSTRALIA\_SOUTHEAST\_1\`,
\`GCP\_EUROPE\_WEST\_1\`, \`GCP\_EUROPE\_WEST\_2\`, \`GCP\_EUROPE\_WEST\_3\`,
\`GCP\_NORTHAMERICA\_NORTHEAST\_1\`, \`GCP\_NORTHAMERICA\_NORTHEAST\_2\`,
\`GCP\_SOUTHAMERICA\_EAST\_1\`, \`GCP\_SOUTHAMERICA\_WEST\_1\`, \`GCP\_US\_CENTRAL\_1\`,
\`GCP\_US\_EAST\_1\`, \`GCP\_US\_EAST\_4\`, \`GCP\_US\_WEST\_1\`, \`GCP\_US\_WEST\_2\`. Defaults to
\`AWS\_AP\_NORTHEAST\_1\`.

Upstream description:

Region location

AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1 GCP\_ASIA\_EAST\_2
GCP\_ASIA\_NORTHEAST\_1 GCP\_ASIA\_NORTHEAST\_2 GCP\_ASIA\_NORTHEAST\_3 GCP\_ASIA\_SOUTH\_1
GCP\_ASIA\_SOUTHEAST\_1 GCP\_ASIA\_SOUTHEAST\_2 GCP\_AUSTRALIA\_SOUTHEAST\_1 GCP\_EUROPE\_WEST\_1
GCP\_EUROPE\_WEST\_2 GCP\_EUROPE\_WEST\_3 GCP\_NORTHAMERICA\_NORTHEAST\_1
GCP\_NORTHAMERICA\_NORTHEAST\_2 GCP\_SOUTHAMERICA\_EAST\_1 GCP\_SOUTHAMERICA\_WEST\_1
GCP\_US\_CENTRAL\_1 GCP\_US\_EAST\_1 GCP\_US\_EAST\_4 GCP\_US\_WEST\_1 GCP\_US\_WEST\_2.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AWS_AP_NORTHEAST_1",
  "enum": [
    "AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-78133f592d6e58e7cd64ae09ac42149bd18da9a0dbc0102c3633645aeb3bca67"></a>

## Next pages — data_center_hosted.egress / 11280758dde5 / 6

- [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-5f420df28410131e89e452d98c433c00b3407ad146136ae6f8e7ce1d9eb0d1c1)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)

<a id="canonical-a1ed7eb915dd7060c2443466cb6cf10027bb4d00b89b40a5d266adfd339ce47c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2393bd6334192828e8c71fe03bd070f45913f88585174951036652eee3c95c80"></a>

## data_center_hosted.ingress — data_center_hosted.ingress / 2dd3c8ffa8eb / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ab45a834092534caee16f8d6620b6949464c5f1b4a9001ce716fc0b8c3b7aec4)
- [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-5f420df28410131e89e452d98c433c00b3407ad146136ae6f8e7ce1d9eb0d1c1)
- data_center_hosted.ingress

<a id="canonical-e97cd2b46e1c740e8990df3a20fe71deaf2c896e3da6f9a83ada94a57cc6f2a8"></a>

Type: `"object"`. list nested block, Optional.

Ingress. Ingress

Upstream description:

Ingress

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("host_name",
    "ip_address")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
ingress {
  # Configure direct properties listed below.
}
```

<a id="canonical-ec3b5a664b577f9d4237af4b49ffbec103c111591f6ec1d258461ab5927b6b81"></a>

## Direct properties — data_center_hosted.ingress / 2dd3c8ffa8eb / 3

<a id="canonical-af72bca54e973e9240481b3475d450b44266415a669b69c3fc78aeacbcb8ddb9"></a>

<a id="canonical-c45708624e4536c92ab1f0e0b3a75f80c79759cfdc556ed0d8de7d972ab37188"></a>

## host_name property — data_center_hosted.ingress / 2dd3c8ffa8eb / 4

Type: `"string"`. Optional.

Exclusive with \[ip\_address\] Ingress Host Name.

Upstream description:

Exclusive with \[ip\_address\] Ingress Host Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-9ea84e8a8ea72edfbc53fbe745efacfd9581fe7403faef1f2730aa217b9bda63"></a>

<a id="canonical-f9ed0c250262cdc3d9428306476bc7d8dec7fe43eb502efad5e34120924c3740"></a>

## ip_address property — data_center_hosted.ingress / 2dd3c8ffa8eb / 5

Type: `"string"`. Optional.

Exclusive with \[host\_name\] Ingress IP Address.

Upstream description:

Exclusive with \[host\_name\] Ingress IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2a1ff49ad0c3595843211a4bad9d453305f23bab254a8fff8610266d692b171e"></a>

<a id="canonical-46b10ca40a1679ec622808b36a0bdde6b24a436a802f10b125f46cdf11f95d25"></a>

## location property — data_center_hosted.ingress / 2dd3c8ffa8eb / 6

Type: `"string"`. Optional.

\[Enum:
AWS\_AP\_NORTHEAST\_1|AWS\_AP\_NORTHEAST\_3|AWS\_AP\_SOUTH\_1|AWS\_AP\_SOUTH\_2|AWS\_AP\_SOUTHEAST\_1|AWS\_AP\_SOUTHEAST\_2|AWS\_AP\_SOUTHEAST\_3|AWS\_EU\_CENTRAL\_1|AWS\_EU\_NORTH\_1|AWS\_EU\_WEST\_1|AWS\_ME\_SOUTH\_1|AWS\_SA\_EAST\_1|AWS\_US\_EAST\_1|AWS\_US\_EAST\_2|AWS\_US\_WEST\_1|AWS\_US\_WEST\_2|GCP\_ASIA\_EAST\_1|GCP\_ASIA\_EAST\_2|GCP\_ASIA\_NORTHEAST\_1|GCP\_ASIA\_NORTHEAST\_2|GCP\_ASIA\_NORTHEAST\_3|GCP\_ASIA\_SOUTH\_1|GCP\_ASIA\_SOUTHEAST\_1|GCP\_ASIA\_SOUTHEAST\_2|GCP\_AUSTRALIA\_SOUTHEAST\_1|GCP\_EUROPE\_WEST\_1|GCP\_EUROPE\_WEST\_2|GCP\_EUROPE\_WEST\_3|GCP\_NORTHAMERICA\_NORTHEAST\_1|GCP\_NORTHAMERICA\_NORTHEAST\_2|GCP\_SOUTHAMERICA\_EAST\_1|GCP\_SOUTHAMERICA\_WEST\_1|GCP\_US\_CENTRAL\_1|GCP\_US\_EAST\_1|GCP\_US\_EAST\_4|GCP\_US\_WEST\_1|GCP\_US\_WEST\_2\]
Region location AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1.. Possible values are
\`AWS\_AP\_NORTHEAST\_1\`, \`AWS\_AP\_NORTHEAST\_3\`, \`AWS\_AP\_SOUTH\_1\`, \`AWS\_AP\_SOUTH\_2\`,
\`AWS\_AP\_SOUTHEAST\_1\`, \`AWS\_AP\_SOUTHEAST\_2\`, \`AWS\_AP\_SOUTHEAST\_3\`,
\`AWS\_EU\_CENTRAL\_1\`, \`AWS\_EU\_NORTH\_1\`, \`AWS\_EU\_WEST\_1\`, \`AWS\_ME\_SOUTH\_1\`,
\`AWS\_SA\_EAST\_1\`, \`AWS\_US\_EAST\_1\`, \`AWS\_US\_EAST\_2\`, \`AWS\_US\_WEST\_1\`,
\`AWS\_US\_WEST\_2\`, \`GCP\_ASIA\_EAST\_1\`, \`GCP\_ASIA\_EAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_1\`,
\`GCP\_ASIA\_NORTHEAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_3\`, \`GCP\_ASIA\_SOUTH\_1\`,
\`GCP\_ASIA\_SOUTHEAST\_1\`, \`GCP\_ASIA\_SOUTHEAST\_2\`, \`GCP\_AUSTRALIA\_SOUTHEAST\_1\`,
\`GCP\_EUROPE\_WEST\_1\`, \`GCP\_EUROPE\_WEST\_2\`, \`GCP\_EUROPE\_WEST\_3\`,
\`GCP\_NORTHAMERICA\_NORTHEAST\_1\`, \`GCP\_NORTHAMERICA\_NORTHEAST\_2\`,
\`GCP\_SOUTHAMERICA\_EAST\_1\`, \`GCP\_SOUTHAMERICA\_WEST\_1\`, \`GCP\_US\_CENTRAL\_1\`,
\`GCP\_US\_EAST\_1\`, \`GCP\_US\_EAST\_4\`, \`GCP\_US\_WEST\_1\`, \`GCP\_US\_WEST\_2\`. Defaults to
\`AWS\_AP\_NORTHEAST\_1\`.

Upstream description:

Region location

AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1 GCP\_ASIA\_EAST\_2
GCP\_ASIA\_NORTHEAST\_1 GCP\_ASIA\_NORTHEAST\_2 GCP\_ASIA\_NORTHEAST\_3 GCP\_ASIA\_SOUTH\_1
GCP\_ASIA\_SOUTHEAST\_1 GCP\_ASIA\_SOUTHEAST\_2 GCP\_AUSTRALIA\_SOUTHEAST\_1 GCP\_EUROPE\_WEST\_1
GCP\_EUROPE\_WEST\_2 GCP\_EUROPE\_WEST\_3 GCP\_NORTHAMERICA\_NORTHEAST\_1
GCP\_NORTHAMERICA\_NORTHEAST\_2 GCP\_SOUTHAMERICA\_EAST\_1 GCP\_SOUTHAMERICA\_WEST\_1
GCP\_US\_CENTRAL\_1 GCP\_US\_EAST\_1 GCP\_US\_EAST\_4 GCP\_US\_WEST\_1 GCP\_US\_WEST\_2.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AWS_AP_NORTHEAST_1",
  "enum": [
    "AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a64e59914eae166d9f9cd0d8d48205bb89bb9cb1582367eb299cb1ad8e0cc115"></a>

## Next pages — data_center_hosted.ingress / 2dd3c8ffa8eb / 7

- [data_center_hosted](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-5f420df28410131e89e452d98c433c00b3407ad146136ae6f8e7ce1d9eb0d1c1)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)

<a id="canonical-34c5003b8b2e9306176db5f499f3dd03f0d68525e71a4952c3e2b33a3e1d5d94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5929312ef49b5cadd2e79aff22bce35556194a47e6762c59c21a983c08327286"></a>

## timeouts — timeouts / 67d7db6ee5d2 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)
- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ab45a834092534caee16f8d6620b6949464c5f1b4a9001ce716fc0b8c3b7aec4)
- timeouts

<a id="canonical-8252f739d135326ed1ce44595cc06f36b23f32da1c22b5adf703aff8c414ae9f"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-70e9a54e5f5c4f8e4fce391aeac8d7cd5dd0ef00d75e58e1fe1eb6952016df0e"></a>

## Direct properties — timeouts / 67d7db6ee5d2 / 3

<a id="canonical-5a4e0695ce4933f6e1d410a2caf9fcb62f0aba13d2b4ae221b9af6aad7b6852f"></a>

<a id="canonical-2181c8aa2d0ecd2ece22463a21ddb093be739b95c8edb799002d94d53bc1b4e0"></a>

## create property — timeouts / 67d7db6ee5d2 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-9328f3f84dcd03d01a5851d4947ea277bd30170a987d71265236eac6c0bf03d7"></a>

<a id="canonical-9e4e65998708fd5d10d564a6d119106279e810de486551ff4650df8ebf72dc00"></a>

## delete property — timeouts / 67d7db6ee5d2 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-67f658cccf8ef8e608e2c1a649cdb1001da7fdf8c72ae3fb6f2138529e28abb5"></a>

<a id="canonical-023102e2bd5c9d2bae5bcda05b9b488a8af355826a959a14880249e6a3bc6e84"></a>

## read property — timeouts / 67d7db6ee5d2 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-9253154cd2688c582e737e20499df11cfa05f2d45988bfbe5d4ab323e8c526cb"></a>

<a id="canonical-66ba0ba329c72bebe276a835fce71f2598b5ed6ef15ca7e2c540d0a2d55ed8b0"></a>

## update property — timeouts / 67d7db6ee5d2 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-48383e3eda5bb0799cf6d3c06b21d6e7f4bd059338706a6b571b1c6cb92723ef"></a>

## Next pages — timeouts / 67d7db6ee5d2 / 8

- [Property reference](resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ab45a834092534caee16f8d6620b6949464c5f1b4a9001ce716fc0b8c3b7aec4)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-630ddf7375046a55c019beba5ab35ecf1937656c254050d05d694d11d7ea1599)
