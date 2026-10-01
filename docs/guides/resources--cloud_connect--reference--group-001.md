---
page_title: "xcsh_cloud_connect reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_connect reference."
---

# xcsh_cloud_connect reference

<a id="canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ee7e56466177fa21ae932e42584eb0ed0c6de598a798eee2bf506c3d1d019b2"></a>

## Property reference — Property reference / 98c4d7ea8e1c / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- Property reference

<a id="canonical-6a6a5c63a8eb9b135fd56164829212231965081c05edc187a115835de2c559cb"></a>

## Direct properties — Property reference / 98c4d7ea8e1c / 3

<a id="canonical-7c9959b28e8951e784144122fc5c4e0745c8b05ddb5a637675610d6805a927b3"></a>

<a id="canonical-dfcaa1f9e5342cca08fc526a0112a47d9dfd3e4317fb6f642487404543156a2e"></a>

## annotations property — Property reference / 98c4d7ea8e1c / 4

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

- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183): complete subsection reference.

- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43): complete subsection reference.

<a id="canonical-251910c9c5947a36199037f1f3cf6b1b80fc1ca71b801e1187d743cbd1731d56"></a>

<a id="canonical-fd0dfa8c5c1a345449dbed47a55499bc15300e44d92fd928212519e2e395bf59"></a>

## description property — Property reference / 98c4d7ea8e1c / 5

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

<a id="canonical-824c566d9d9c9fec09e3d7c939fb9f811c25c0356e681477414bdab4fdc0eaeb"></a>

<a id="canonical-53e2f406ca03f27cfd777807aab1dee464d2daf1acf70efd681b3bdc3fe62761"></a>

## disable property — Property reference / 98c4d7ea8e1c / 6

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

<a id="canonical-51708efad332d7be15613dcf97a8a6681899dc9be9236ce7349c4bc4547dac2e"></a>

<a id="canonical-02399367995091fa527e93b54693d8d3171b36c1f7da1364ce63730f534f088e"></a>

## id property — Property reference / 98c4d7ea8e1c / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-842bbb32c62a171b84155bf10ddadfd61ac11bc2f9b9b7580a11429f3ac02184"></a>

<a id="canonical-d149ccc225a95d62eeab9211d779769d46a6a0125b0c833cb503657468d4f58a"></a>

## labels property — Property reference / 98c4d7ea8e1c / 8

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

<a id="canonical-095e4e8f7fcad2e33f64a7279373733bfc7d7e48df6486cecd15d53b80eac0c6"></a>

<a id="canonical-2629dd1d161aa0fa81e7a017cc9723eb90f475452910bf5193cb41c72ce21d4c"></a>

## name property — Property reference / 98c4d7ea8e1c / 9

Type: `"string"`. Required.

Name of the Cloud Connect. Must be unique within the namespace.

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

<a id="canonical-bc33282238b010330802c6c5bcdfe6bebe9c350565987a7cc8a4327f02be589f"></a>

<a id="canonical-67097a0a60cedf3999367665c5dcc77cc548bd0be963914af61d74b6ea767012"></a>

## namespace property — Property reference / 98c4d7ea8e1c / 10

Type: `"string"`. Required.

Namespace where the Cloud Connect is created.

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

- [segment](resources--cloud_connect--reference--group-001.md#canonical-26858322626c00d5922e9a6ee025eab4987bc0ae4eb5e0a79f4eb3ac39f170ae): complete subsection reference.

- [timeouts](resources--cloud_connect--reference--group-001.md#canonical-120208e71f2e652ac918b8a30a42d0ca791c2ea433e8d57c4de612854efd777e): complete subsection reference.

<a id="canonical-f69f401d5f86386178c34dd35240ce8013f590d4831e0eaa6599a0f90e4b52aa"></a>

## All schema paths — Property reference / 98c4d7ea8e1c / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cloud_connect--reference--group-001.md#canonical-7c9959b28e8951e784144122fc5c4e0745c8b05ddb5a637675610d6805a927b3) |
| `aws_provider` | [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-ac195748983c5e95abc7df8ff62f37f097d184920e4c2812c763c4995e61585d) |
| `aws_provider.aws_tgw_site` | [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-b2b3b53f721988ba6b78feb7fa83ed58ffca5f79985e93f7f6b563864ad2731e) |
| `aws_provider.aws_tgw_site.cred` | [aws_provider.aws_tgw_site.cred](resources--cloud_connect--reference--group-001.md#canonical-8966a5eee672cc937aaca4aba170c31577ef7985455710ee59582898ed12b792) |
| `aws_provider.aws_tgw_site.cred.name` | [aws_provider.aws_tgw_site.cred.name](resources--cloud_connect--reference--group-001.md#canonical-87da9d521da0e859ab87507885c9038139bdc0371173a1129af5741963e85a84) |
| `aws_provider.aws_tgw_site.cred.namespace` | [aws_provider.aws_tgw_site.cred.namespace](resources--cloud_connect--reference--group-001.md#canonical-5c6e14717870260885aec52ff80f479b44adcbd8bfb732ab210399c7e9524a09) |
| `aws_provider.aws_tgw_site.cred.tenant` | [aws_provider.aws_tgw_site.cred.tenant](resources--cloud_connect--reference--group-001.md#canonical-27cb09069090dd7099c91463d0bd81a9d0a9ad894b35f15e4eb92a03cee81d73) |
| `aws_provider.aws_tgw_site.site` | [aws_provider.aws_tgw_site.site](resources--cloud_connect--reference--group-001.md#canonical-4149b1a3bd6a8f45a8572f3fa3d5257b3cc23e51b5cdfc1ba7413edaab96b5f5) |
| `aws_provider.aws_tgw_site.site.name` | [aws_provider.aws_tgw_site.site.name](resources--cloud_connect--reference--group-001.md#canonical-856174654010b41e13dd1cf1bb6ebb0377865b570a5007c2350edeb72521d76c) |
| `aws_provider.aws_tgw_site.site.namespace` | [aws_provider.aws_tgw_site.site.namespace](resources--cloud_connect--reference--group-001.md#canonical-427a8cbcb21b8adca7a7c9dda59a5e4cebdf6446fc7024fef93a30c4c0041d94) |
| `aws_provider.aws_tgw_site.site.tenant` | [aws_provider.aws_tgw_site.site.tenant](resources--cloud_connect--reference--group-001.md#canonical-615d1a6116de346092867d43b02057e1a37adda5782006550e96e176423b8419) |
| `aws_provider.aws_tgw_site.vpc_attachments` | [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-625f5b95f64b6a9af97978128a06a843063967b250c76ca23d12160d72cea2d4) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-12b6f23a40c3b9ef81e8494193a489949a0b11074ce2cb23c99bc8eb98875487) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-8861b635a0a24c982ec204a79f4d9da07700bcd5a5e385b57e77b95241118db7) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables](resources--cloud_connect--reference--group-001.md#canonical-921afc59be907e92b5170242436786232a2f5f3c0015c956f4e4d6936dbeca0f) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.route_table_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.route_table_id](resources--cloud_connect--reference--group-001.md#canonical-e6d08c755aea0bff8b7aa17786e5f74820e49eaa6b7cea11b6d71e15e1e19fe1) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.static_routes` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.static_routes](resources--cloud_connect--reference--group-001.md#canonical-f54ec7d3a66e578485b0c317ba0855277b8b3e3423e3dcb7e3e5e68f71d9f8e2) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-a705470377a3615560f33957ec3cb0e09ae9bd3d72a1514f11099853f2098901) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables](resources--cloud_connect--reference--group-001.md#canonical-8227089d96277a3afa0dc37f139404ff8b5ac10ede24334fca73ac9e0b46f894) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables](resources--cloud_connect--reference--group-001.md#canonical-cea32f75f3d6f845ac779caa18957a3bd2c2974c4c257aa502c2a71d1363c928) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables.route_table_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables.route_table_id](resources--cloud_connect--reference--group-001.md#canonical-c50e86aa36f9921cddf1fa3493e834f19a98132f883a0dceec9277a74ab442c0) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels](resources--cloud_connect--reference--group-001.md#canonical-965572ba3545606777547372a26bceeeee2dc8ecab5234f6c9ea4ac0255079ef) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing](resources--cloud_connect--reference--group-001.md#canonical-299acc2d6a86f3b8bc0104d7f594d52b1445e589852f6317b934aa1f0cf32fe1) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.vpc_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.vpc_id](resources--cloud_connect--reference--group-001.md#canonical-1ee8719d011b347be4fd348b7b4b1e84c2e309c9c6d12c339f418bf7a6b53f86) |
| `azure_vnet_site` | [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-f8ba508baae5e7d74c974ed301d3ef03f113f89bc96acbfa79ec3174e8f00a66) |
| `azure_vnet_site.site` | [azure_vnet_site.site](resources--cloud_connect--reference--group-001.md#canonical-2aefc4abf03423e1eaf440e54589553802d013563dbba8fac3332e9c0e20ecc0) |
| `azure_vnet_site.site.name` | [azure_vnet_site.site.name](resources--cloud_connect--reference--group-001.md#canonical-bffab6a46facd91cdf928e06810503f98f25b878e3ff084bff7215dedaa56f8c) |
| `azure_vnet_site.site.namespace` | [azure_vnet_site.site.namespace](resources--cloud_connect--reference--group-001.md#canonical-22dfd1b806784442e618ca10ba6b83b0653b508418087b3d7c2c64e3671f3e0d) |
| `azure_vnet_site.site.tenant` | [azure_vnet_site.site.tenant](resources--cloud_connect--reference--group-001.md#canonical-d71db0adad8c563e4fc9a6ac1f3312b1d94530dc24b87f834ece57dd175c661a) |
| `azure_vnet_site.vnet_attachments` | [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-c3922d5543c74965f926ebb7e2a61f785bc5fce38caf52b36c55bde7794fc50e) |
| `azure_vnet_site.vnet_attachments.vnet_list` | [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-3327e369bba0f8cb3ee193701e1fd4e398e0ac41ef60975aced4efd9760ecbf3) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-a53212e3807a9b65f18383bb163942a51555c053daf2a62853481dae0ca7093d) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables](resources--cloud_connect--reference--group-001.md#canonical-99c607148056ab4255743d9b490fd32c09debf724b8b126a62b60bc9075d73f4) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.route_table_id` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.route_table_id](resources--cloud_connect--reference--group-001.md#canonical-e53e36049ac6b61a58d7e967b3f2415292a7025496178e09491bad444fe9029d) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.static_routes` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.static_routes](resources--cloud_connect--reference--group-001.md#canonical-db230333a0625576972fa1eb9e0449022bc14e6ef75452f8537cf4d7a3e10e17) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route` | [azure_vnet_site.vnet_attachments.vnet_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-d88ffbcc4eb21ca5b5ff7f399e5ee601f99da2c968866033a763a16201e128f3) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables](resources--cloud_connect--reference--group-001.md#canonical-f584a6877e4c1b7dab7ce5d3b76915c027570617c031e6d494e6d1bf90f6d8bf) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables](resources--cloud_connect--reference--group-001.md#canonical-1c9f55a5aa55808d1ed1802be993674fa76e2ce7d87e2a9afbb346ad4fddef72) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables.route_table_id` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables.route_table_id](resources--cloud_connect--reference--group-001.md#canonical-3c25b3c7eca9acfc9a1f613b9250b9fc5da5f7e492e06b6252b541c33e659c59) |
| `azure_vnet_site.vnet_attachments.vnet_list.labels` | [azure_vnet_site.vnet_attachments.vnet_list.labels](resources--cloud_connect--reference--group-001.md#canonical-2a098bf12b5f8b30191df07c5e8c4900dfe962103a0dd819f75fd7292ef79670) |
| `azure_vnet_site.vnet_attachments.vnet_list.manual_routing` | [azure_vnet_site.vnet_attachments.vnet_list.manual_routing](resources--cloud_connect--reference--group-001.md#canonical-a0738cf5eecb0b7aed17f3770124dba590687ce0dddc625829aae61335ef1402) |
| `azure_vnet_site.vnet_attachments.vnet_list.subscription_id` | [azure_vnet_site.vnet_attachments.vnet_list.subscription_id](resources--cloud_connect--reference--group-001.md#canonical-becd1d64965e2564c1a1b98b7be266914c9596b54bafd9f01e3af09c2968ae7b) |
| `azure_vnet_site.vnet_attachments.vnet_list.vnet_id` | [azure_vnet_site.vnet_attachments.vnet_list.vnet_id](resources--cloud_connect--reference--group-001.md#canonical-db581eeb673f8e2f78c39ea0a99067f33fbc20d1c2377923d5f76be533c2f1d9) |
| `description` | [description](resources--cloud_connect--reference--group-001.md#canonical-251910c9c5947a36199037f1f3cf6b1b80fc1ca71b801e1187d743cbd1731d56) |
| `disable` | [disable](resources--cloud_connect--reference--group-001.md#canonical-824c566d9d9c9fec09e3d7c939fb9f811c25c0356e681477414bdab4fdc0eaeb) |
| `id` | [id](resources--cloud_connect--reference--group-001.md#canonical-51708efad332d7be15613dcf97a8a6681899dc9be9236ce7349c4bc4547dac2e) |
| `labels` | [labels](resources--cloud_connect--reference--group-001.md#canonical-842bbb32c62a171b84155bf10ddadfd61ac11bc2f9b9b7580a11429f3ac02184) |
| `name` | [name](resources--cloud_connect--reference--group-001.md#canonical-095e4e8f7fcad2e33f64a7279373733bfc7d7e48df6486cecd15d53b80eac0c6) |
| `namespace` | [namespace](resources--cloud_connect--reference--group-001.md#canonical-bc33282238b010330802c6c5bcdfe6bebe9c350565987a7cc8a4327f02be589f) |
| `segment` | [segment](resources--cloud_connect--reference--group-001.md#canonical-697ce5a1b833b2b846c62ce34d960c0e126fbb1a6d81645bbb267242b855839b) |
| `segment.name` | [segment.name](resources--cloud_connect--reference--group-001.md#canonical-00431dd3798be33f6ee5fbdc96467e56509937e5ef643c4736a9680e54ab2e38) |
| `segment.namespace` | [segment.namespace](resources--cloud_connect--reference--group-001.md#canonical-b5850f3036636a37990b13c846f297e1fe9a810d238551f4ca009dea991d41d8) |
| `segment.tenant` | [segment.tenant](resources--cloud_connect--reference--group-001.md#canonical-d6118005356ca880a19fd7b7b474b8462067eafe1953df34a54900bbb51a487a) |
| `timeouts` | [timeouts](resources--cloud_connect--reference--group-001.md#canonical-916da7551befb29d84365520eca4f474cca8825399f98dc77d881a83d5973948) |
| `timeouts.create` | [timeouts.create](resources--cloud_connect--reference--group-001.md#canonical-f2ead0aa4006ce0692b09e2ed53c39966b6dfc642258797e802509bf17929431) |
| `timeouts.delete` | [timeouts.delete](resources--cloud_connect--reference--group-001.md#canonical-f6f102ffd5ade80916ee8d977393870370fd8c0efb94f3a13b6bff3ce879a45a) |
| `timeouts.read` | [timeouts.read](resources--cloud_connect--reference--group-001.md#canonical-c98b87041d6430da463a334c983a9c1664eabc6a940c0bcc517b61adc1eadabb) |
| `timeouts.update` | [timeouts.update](resources--cloud_connect--reference--group-001.md#canonical-a4bf29cfc566b14f6c58c89fed01d4ad1b7028cff70caf02fb6ffc7ed37d8625) |

<a id="canonical-0eee82ec0047a7dcaa6d1f37cb46e79b3518bb6fb3edcf5041b4ae085dc4555d"></a>

## Next pages — Property reference / 98c4d7ea8e1c / 12

- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43)
- [segment](resources--cloud_connect--reference--group-001.md#canonical-26858322626c00d5922e9a6ee025eab4987bc0ae4eb5e0a79f4eb3ac39f170ae)
- [timeouts](resources--cloud_connect--reference--group-001.md#canonical-120208e71f2e652ac918b8a30a42d0ca791c2ea433e8d57c4de612854efd777e)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78057383c6ba2e61ad3b5d204e514ecf4f1da23c8892d015df7d829c27e1a0a5"></a>

## aws_provider — aws_provider / 7fc0f774815c / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- aws_provider

<a id="canonical-ac195748983c5e95abc7df8ff62f37f097d184920e4c2812c763c4995e61585d"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: aws\_provider, azure\_vnet\_site\] Configuration parameter for aws provider.

Upstream description:

Cloud Connect with AWS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-site_type": "[\"aws_tgw_site\"]"
}
```

OneOf alternatives in this subsection:

- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-ac195748983c5e95abc7df8ff62f37f097d184920e4c2812c763c4995e61585d)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-f8ba508baae5e7d74c974ed301d3ef03f113f89bc96acbfa79ec3174e8f00a66)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
aws_provider {
  # Configure direct properties listed below.
}
```

<a id="canonical-51df3d860e46b04d022945bb43babb90dd25560f37402c0e67b006a9cbaf2329"></a>

## Direct properties — aws_provider / 7fc0f774815c / 3

- [aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3): complete subsection reference.

<a id="canonical-974ba3ff0992ee26603417658a5dab3c269aaf20fcfc71e1ff08cc7c9008c4a7"></a>

## Next pages — aws_provider / 7fc0f774815c / 4

- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5aa739f19f26f019956fcba02ba45a6d06ee86ebc3a9bd2b65a571a8ace6e35b"></a>

## aws_provider.aws_tgw_site — aws_provider.aws_tgw_site / 784b2f17e349 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183)
- aws_provider.aws_tgw_site

<a id="canonical-b2b3b53f721988ba6b78feb7fa83ed58ffca5f79985e93f7f6b563864ad2731e"></a>

Type: `"object"`. single nested block, Optional.

AWS TGW Site Type. Cloud Connect AWS TGW Site Type.

Upstream description:

Cloud Connect AWS TGW Site Type.

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

<a id="canonical-a3062dda2c0959f80d73810c0f0f446744bf5cbe79f220162129f1b7cdf4fcdb"></a>

## Direct properties — aws_provider.aws_tgw_site / 784b2f17e349 / 3

- [cred](resources--cloud_connect--reference--group-001.md#canonical-f6fb99d4e7f14983f469457d6de7afeeef8986fc5f00d5d5576de939eb61329d): complete subsection reference.

- [site](resources--cloud_connect--reference--group-001.md#canonical-2ff1246136c85584f598e1acdbda2b8af97c2f5f711f2793d801ea49ad8e5380): complete subsection reference.

- [vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-86e98d47902999a315048aa07ede1c2ac10023c805b397a5613d673d507b8314): complete subsection reference.

<a id="canonical-55eb10b3a555c54553e925813ec52a087d3c096e3c9316806b05e142fba31ee2"></a>

## Next pages — aws_provider.aws_tgw_site / 784b2f17e349 / 4

- [aws_provider.aws_tgw_site.cred](resources--cloud_connect--reference--group-001.md#canonical-f6fb99d4e7f14983f469457d6de7afeeef8986fc5f00d5d5576de939eb61329d)
- [aws_provider.aws_tgw_site.site](resources--cloud_connect--reference--group-001.md#canonical-2ff1246136c85584f598e1acdbda2b8af97c2f5f711f2793d801ea49ad8e5380)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-86e98d47902999a315048aa07ede1c2ac10023c805b397a5613d673d507b8314)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-f6fb99d4e7f14983f469457d6de7afeeef8986fc5f00d5d5576de939eb61329d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-529e6cbd2c113ac469b8dedfe54323767fa656ba095259bbb449d7b5e95acb5b"></a>

## aws_provider.aws_tgw_site.cred — aws_provider.aws_tgw_site.cred / a48baa5e518c / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- aws_provider.aws_tgw_site.cred

<a id="canonical-8966a5eee672cc937aaca4aba170c31577ef7985455710ee59582898ed12b792"></a>

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
cred {
  # Configure direct properties listed below.
}
```

<a id="canonical-5958509ae3787ca26d239b34a41fa8d4be93e6e28182055ca3739f9187d21179"></a>

## Direct properties — aws_provider.aws_tgw_site.cred / a48baa5e518c / 3

<a id="canonical-87da9d521da0e859ab87507885c9038139bdc0371173a1129af5741963e85a84"></a>

<a id="canonical-ff08947b44186a29fe1be659cccfd156039dd30f9cc6368c7320447365b859e0"></a>

## name property — aws_provider.aws_tgw_site.cred / a48baa5e518c / 4

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

<a id="canonical-5c6e14717870260885aec52ff80f479b44adcbd8bfb732ab210399c7e9524a09"></a>

<a id="canonical-cf6514472f1a61907c43b386a6ba668a12e474755c1c95a8fb02cbc866bbb64d"></a>

## namespace property — aws_provider.aws_tgw_site.cred / a48baa5e518c / 5

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

<a id="canonical-27cb09069090dd7099c91463d0bd81a9d0a9ad894b35f15e4eb92a03cee81d73"></a>

<a id="canonical-519eeb7899ae908d76a23722929245eef28978770c54ecb601c6e58d0dbd61e7"></a>

## tenant property — aws_provider.aws_tgw_site.cred / a48baa5e518c / 6

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

<a id="canonical-c978b5a9ebd365e9cadaf60674a3ae6b289a63e2c5cc8e2a603b65d44aae7f1f"></a>

## Next pages — aws_provider.aws_tgw_site.cred / a48baa5e518c / 7

- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-2ff1246136c85584f598e1acdbda2b8af97c2f5f711f2793d801ea49ad8e5380"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef2fd29f159fae1b41edb7fc7930067006eac0774113648cfb43b61c15695349"></a>

## aws_provider.aws_tgw_site.site — aws_provider.aws_tgw_site.site / c302264d3a29 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- aws_provider.aws_tgw_site.site

<a id="canonical-4149b1a3bd6a8f45a8572f3fa3d5257b3cc23e51b5cdfc1ba7413edaab96b5f5"></a>

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

<a id="canonical-9fac920e4f8d5c157b53c76f456b8ec3ea5d1b168aa5e8da71681d81e477469d"></a>

## Direct properties — aws_provider.aws_tgw_site.site / c302264d3a29 / 3

<a id="canonical-856174654010b41e13dd1cf1bb6ebb0377865b570a5007c2350edeb72521d76c"></a>

<a id="canonical-442cd010ad7b1e1d604dcc745672d542c59249124ab0c4c7d4ec9209f98eeadb"></a>

## name property — aws_provider.aws_tgw_site.site / c302264d3a29 / 4

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

<a id="canonical-427a8cbcb21b8adca7a7c9dda59a5e4cebdf6446fc7024fef93a30c4c0041d94"></a>

<a id="canonical-f76d0bd367dc322d72ae25e70880db74ed267e3cab59aa1a02043a5517912aef"></a>

## namespace property — aws_provider.aws_tgw_site.site / c302264d3a29 / 5

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

<a id="canonical-615d1a6116de346092867d43b02057e1a37adda5782006550e96e176423b8419"></a>

<a id="canonical-0af8fc6d094e313f0eb104f875661a5a1331312c451be0313514e6237657d417"></a>

## tenant property — aws_provider.aws_tgw_site.site / c302264d3a29 / 6

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

<a id="canonical-4010bf4d531d511417f51669da7aac02ba9247d94eb8b4baddf293b8f37d29bd"></a>

## Next pages — aws_provider.aws_tgw_site.site / c302264d3a29 / 7

- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-86e98d47902999a315048aa07ede1c2ac10023c805b397a5613d673d507b8314"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b77cb70fdcba8ddec9bb7e93ccf79d13469e2e0ee22e20ca5515ec0816b5d944"></a>

## aws_provider.aws_tgw_site.vpc_attachments — aws_provider.aws_tgw_site.vpc_attachments / 910629bd1459 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- aws_provider.aws_tgw_site.vpc_attachments

<a id="canonical-625f5b95f64b6a9af97978128a06a843063967b250c76ca23d12160d72cea2d4"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vpc attachments.

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
vpc_attachments {
  # Configure direct properties listed below.
}
```

<a id="canonical-51f0657ab9decb790408ea301af4da680e2cc694eb5fa6a40313601fcd7d4651"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments / 910629bd1459 / 3

- [vpc_list](resources--cloud_connect--reference--group-001.md#canonical-e0c264673a4f9876cdfa0c601c75c9d819671adf522d529d5e4975160cb8d511): complete subsection reference.

<a id="canonical-170825bf80a01ac3181e97308ef79b38ec68d6a0cdf361aa12ea23c9f41045eb"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments / 910629bd1459 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-e0c264673a4f9876cdfa0c601c75c9d819671adf522d529d5e4975160cb8d511)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-e0c264673a4f9876cdfa0c601c75c9d819671adf522d529d5e4975160cb8d511"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-701d56c8a481d0ceb19ab66b15f04cda3cf80cb9baf278d86b7d75b273d0c562"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list — aws_provider.aws_tgw_site.vpc_attachments.vpc_list / e1c36336f37c / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-86e98d47902999a315048aa07ede1c2ac10023c805b397a5613d673d507b8314)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list

<a id="canonical-12b6f23a40c3b9ef81e8494193a489949a0b11074ce2cb23c99bc8eb98875487"></a>

Type: `"object"`. list nested block, Optional.

VPC List. Collection of items or values

Upstream description:

Collection of items or values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("vpc_id"),
  validators.ConflictingListObjectAttributes("custom_routing",
    "default_route"),
  validators.ConflictingListObjectAttributes("custom_routing",
    "manual_routing"),
  validators.ConflictingListObjectAttributes("default_route",
    "manual_routing")}
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

Terraform syntax:

```terraform
vpc_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-36205d227d69a59ca64c17405721cb3fb04ecf539a9b2b73e2df4d939e48d5b9"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list / e1c36336f37c / 3

- [custom_routing](resources--cloud_connect--reference--group-001.md#canonical-d78bf53652375c6f7ca9e6d0c4dee110a560f0a1116caf8142e29bdfcd85f189): complete subsection reference.

- [default_route](resources--cloud_connect--reference--group-001.md#canonical-efd0fe58946304996a9a7645e58fb34a3b444485ce3d160a790519bd7b909d37): complete subsection reference.

- [labels](resources--cloud_connect--reference--group-001.md#canonical-34fe5129c386693f62b7b8fcc1a585411831b6fe58c827af9bb3b7f0dadbdf34): complete subsection reference.

- [manual_routing](resources--cloud_connect--reference--group-001.md#canonical-485ffc918c881e29c4e3f652b1f00ba63b8b44c8a8a5381d56c81909b821c8ed): complete subsection reference.

<a id="canonical-1ee8719d011b347be4fd348b7b4b1e84c2e309c9c6d12c339f418bf7a6b53f86"></a>

<a id="canonical-7ec22e9755111a2486e3006e7786e85c1668a47dc1b2f7ab35a9c02facda04c6"></a>

## vpc_id property — aws_provider.aws_tgw_site.vpc_attachments.vpc_list / e1c36336f37c / 4

Type: `"string"`. Optional.

Enter the VPC ID of the VPC to be attached.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-c7f9e5dfd4591c4943ec81709f3c5fe028926690df0859bb57d188a864e1633f"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list / e1c36336f37c / 5

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-d78bf53652375c6f7ca9e6d0c4dee110a560f0a1116caf8142e29bdfcd85f189)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-efd0fe58946304996a9a7645e58fb34a3b444485ce3d160a790519bd7b909d37)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels](resources--cloud_connect--reference--group-001.md#canonical-34fe5129c386693f62b7b8fcc1a585411831b6fe58c827af9bb3b7f0dadbdf34)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing](resources--cloud_connect--reference--group-001.md#canonical-485ffc918c881e29c4e3f652b1f00ba63b8b44c8a8a5381d56c81909b821c8ed)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-86e98d47902999a315048aa07ede1c2ac10023c805b397a5613d673d507b8314)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-d78bf53652375c6f7ca9e6d0c4dee110a560f0a1116caf8142e29bdfcd85f189"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1dcc06985ac3e4f23df767d8041384332465e068fe4e871840f901adf1e411a9"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing / c20bec4cedfc / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-86e98d47902999a315048aa07ede1c2ac10023c805b397a5613d673d507b8314)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-e0c264673a4f9876cdfa0c601c75c9d819671adf522d529d5e4975160cb8d511)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing

<a id="canonical-8861b635a0a24c982ec204a79f4d9da07700bcd5a5e385b57e77b95241118db7"></a>

Type: `"object"`. single nested block, Optional.

AWS Route Table List. AWS Route Table List.

Upstream description:

AWS Route Table List.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("route_tables")}
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
custom_routing {
  # Configure direct properties listed below.
}
```

<a id="canonical-e2303d0a97aa662cb11c9667ffed57775abe72e06e2216e58fd679640dedb6ba"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing / c20bec4cedfc / 3

- [route_tables](resources--cloud_connect--reference--group-001.md#canonical-9fd4e778c7ff611e670527098f9598d4ac26ff4c6c5ad2ad19297712cd33d940): complete subsection reference.

<a id="canonical-8d9fe96505b9e17c9c91cc8b376d0bd92c1a777d958ae1cd1e51c6997cf9986d"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing / c20bec4cedfc / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables](resources--cloud_connect--reference--group-001.md#canonical-9fd4e778c7ff611e670527098f9598d4ac26ff4c6c5ad2ad19297712cd33d940)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-e0c264673a4f9876cdfa0c601c75c9d819671adf522d529d5e4975160cb8d511)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-9fd4e778c7ff611e670527098f9598d4ac26ff4c6c5ad2ad19297712cd33d940"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d56bf11d30acdd9c71c6f20b975dc57de8a0a2a1c5d76d72887fa16d0b74ed49"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables / 571b067359e1 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-86e98d47902999a315048aa07ede1c2ac10023c805b397a5613d673d507b8314)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-e0c264673a4f9876cdfa0c601c75c9d819671adf522d529d5e4975160cb8d511)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-d78bf53652375c6f7ca9e6d0c4dee110a560f0a1116caf8142e29bdfcd85f189)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables

<a id="canonical-921afc59be907e92b5170242436786232a2f5f3c0015c956f4e4d6936dbeca0f"></a>

Type: `"object"`. list nested block, Optional.

List of route tables. Route Tables.

Upstream description:

Route Tables.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("static_routes")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 200,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 200,
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
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
route_tables {
  # Configure direct properties listed below.
}
```

<a id="canonical-0093f208fbf68338164d490731ed4735e5e6c55ab8e1d281fc104ab8c37e8077"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables / 571b067359e1 / 3

<a id="canonical-e6d08c755aea0bff8b7aa17786e5f74820e49eaa6b7cea11b6d71e15e1e19fe1"></a>

<a id="canonical-38937b03d84b631ab7df39942d3e2972af77d70d3cc2c11c73c7e30c462e35da"></a>

## route_table_id property — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables / 571b067359e1 / 4

Type: `"string"`. Optional.

Route table ID. Route table ID.

Upstream description:

Route table ID.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-f54ec7d3a66e578485b0c317ba0855277b8b3e3423e3dcb7e3e5e68f71d9f8e2"></a>

<a id="canonical-15e517ffdd280e101ccbf77dec8875d671d7d3491c73ac0a98a1d96ca8550ac1"></a>

## static_routes property — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables / 571b067359e1 / 5

Type: `["list", "string"]`. Optional.

Static Routes. List of Static Routes.

Upstream description:

List of Static Routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
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
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1d7e5bed7a6b22ce909c03150cea0794575d3b2970a634a5b24fb43a1b6d1851"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables / 571b067359e1 / 6

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-d78bf53652375c6f7ca9e6d0c4dee110a560f0a1116caf8142e29bdfcd85f189)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-efd0fe58946304996a9a7645e58fb34a3b444485ce3d160a790519bd7b909d37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afbebf85d4fb34691026381e1ad8bd268a41d3a4b8f4c9b27171844f52b573f8"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route / 50f0002ca00f / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-86e98d47902999a315048aa07ede1c2ac10023c805b397a5613d673d507b8314)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-e0c264673a4f9876cdfa0c601c75c9d819671adf522d529d5e4975160cb8d511)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route

<a id="canonical-a705470377a3615560f33957ec3cb0e09ae9bd3d72a1514f11099853f2098901"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

Select Override Default Route Choice.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_route_tables",
    "selective_route_tables")}
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
  "x-ves-oneof-field-default_route_choice": "[\"all_route_tables\",\"selective_route_tables\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-e017c8e28a7b2864dbe3618a01cd110d07e14950e7a676e17e8bd4ef2a0548a4"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route / 50f0002ca00f / 3

- [all_route_tables](resources--cloud_connect--reference--group-001.md#canonical-8493c7c657f59a25b195524a282c6f95af3f349dfc37c4953cdaffc4df054851): complete subsection reference.

- [selective_route_tables](resources--cloud_connect--reference--group-001.md#canonical-a9ad320defcdf7db75c7b25e2ade4138199cb7908c211a65ea64c9275632cc0c): complete subsection reference.

<a id="canonical-7789d87f81329ba52ec75845ac059ec85f14e7ea8f6b1e4989fcdee8e8b0f05c"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route / 50f0002ca00f / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables](resources--cloud_connect--reference--group-001.md#canonical-8493c7c657f59a25b195524a282c6f95af3f349dfc37c4953cdaffc4df054851)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables](resources--cloud_connect--reference--group-001.md#canonical-a9ad320defcdf7db75c7b25e2ade4138199cb7908c211a65ea64c9275632cc0c)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-e0c264673a4f9876cdfa0c601c75c9d819671adf522d529d5e4975160cb8d511)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-8493c7c657f59a25b195524a282c6f95af3f349dfc37c4953cdaffc4df054851"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e828a5d224aa112667783b713bbfaef41aac88e891ccf1ff37ec2042c3f2ea6"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_table / a8f8f5193618 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-86e98d47902999a315048aa07ede1c2ac10023c805b397a5613d673d507b8314)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-e0c264673a4f9876cdfa0c601c75c9d819671adf522d529d5e4975160cb8d511)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-efd0fe58946304996a9a7645e58fb34a3b444485ce3d160a790519bd7b909d37)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables

<a id="canonical-8227089d96277a3afa0dc37f139404ff8b5ac10ede24334fca73ac9e0b46f894"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all route tables.

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
all_route_tables = {}
```

<a id="canonical-bfca8e59e7c165d7b31934b0fbbedc39f3f61d677796785f3e29059f7b84c95c"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_table / a8f8f5193618 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-00148bec0a1b5accbe01c5869e41f77ac21b1bb8f4a7c5305b7b98e6c8030c0b"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_table / a8f8f5193618 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-efd0fe58946304996a9a7645e58fb34a3b444485ce3d160a790519bd7b909d37)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-a9ad320defcdf7db75c7b25e2ade4138199cb7908c211a65ea64c9275632cc0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e8a360003e343d25d60cd1e235e490a64556890c953dff36a9f56d34b0b5256"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route / d838766d3452 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-86e98d47902999a315048aa07ede1c2ac10023c805b397a5613d673d507b8314)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-e0c264673a4f9876cdfa0c601c75c9d819671adf522d529d5e4975160cb8d511)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-efd0fe58946304996a9a7645e58fb34a3b444485ce3d160a790519bd7b909d37)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables

<a id="canonical-cea32f75f3d6f845ac779caa18957a3bd2c2974c4c257aa502c2a71d1363c928"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for selective route tables.

Upstream description:

AWS Route Table.

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
selective_route_tables {
  # Configure direct properties listed below.
}
```

<a id="canonical-d755007a9fd6e893ef7dec8375673cc85a15d2c133dd2804356b3954917fd52c"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route / d838766d3452 / 3

<a id="canonical-c50e86aa36f9921cddf1fa3493e834f19a98132f883a0dceec9277a74ab442c0"></a>

<a id="canonical-c240282e5e93143ee556e48f6bde3b8de45f23df729c985c74c42a21f2e2245d"></a>

## route_table_id property — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route / d838766d3452 / 4

Type: `["list", "string"]`. Optional.

Route table ID. Route table ID.

Upstream description:

Route table ID.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b81a859e82e3ac0b8e3c6dde02e188927b678137eb9a49c15f8f3393997af4f6"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route / d838766d3452 / 5

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-efd0fe58946304996a9a7645e58fb34a3b444485ce3d160a790519bd7b909d37)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-34fe5129c386693f62b7b8fcc1a585411831b6fe58c827af9bb3b7f0dadbdf34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab33259f9b370a58774dfefa35e085881028acc0a3d00f3edccf7185070c7a7d"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels / 67a139417148 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-86e98d47902999a315048aa07ede1c2ac10023c805b397a5613d673d507b8314)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-e0c264673a4f9876cdfa0c601c75c9d819671adf522d529d5e4975160cb8d511)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels

<a id="canonical-965572ba3545606777547372a26bceeeee2dc8ecab5234f6c9ea4ac0255079ef"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
labels {}
```

<a id="canonical-d0387113153399bfa90ce3e488ffddb0ebe32c49c20432f5e34c9317005c33fc"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels / 67a139417148 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-39c4eb1bc460235f26cca384cdee637381e8af3ebd4df90efff05e899c7ab48b"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels / 67a139417148 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-e0c264673a4f9876cdfa0c601c75c9d819671adf522d529d5e4975160cb8d511)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-485ffc918c881e29c4e3f652b1f00ba63b8b44c8a8a5381d56c81909b821c8ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-328537441cfb164c3406a7c2dfd51b7ad4a0ec7fefb50938e691dcb3daa20e69"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing / 246d2d0971d5 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-d3468a70eaff1b49db395794229aaf948232e3910aa564bab064b7248c0cf183)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-50e12b7a4fdc1a4c974dbc5eee01103645c7a416c5dd4219cdf695f5b7963aa3)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-86e98d47902999a315048aa07ede1c2ac10023c805b397a5613d673d507b8314)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-e0c264673a4f9876cdfa0c601c75c9d819671adf522d529d5e4975160cb8d511)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing

<a id="canonical-299acc2d6a86f3b8bc0104d7f594d52b1445e589852f6317b934aa1f0cf32fe1"></a>

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
manual_routing = {}
```

<a id="canonical-5bb6f8a3cf413925154329802d0cd7fd85e6b3d7e51a264b9b35c6bd76d03772"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing / 246d2d0971d5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4bcc7fed60d6f217ab489e2ced0a839509e8731bc2981e71cf6531209fef8958"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing / 246d2d0971d5 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-e0c264673a4f9876cdfa0c601c75c9d819671adf522d529d5e4975160cb8d511)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85cfdf7a9da3ff525edf108010e4a117fbccc5cb6948aadd3d851095cc9743e0"></a>

## azure_vnet_site — azure_vnet_site / 6e1d91178d66 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- azure_vnet_site

<a id="canonical-f8ba508baae5e7d74c974ed301d3ef03f113f89bc96acbfa79ec3174e8f00a66"></a>

Type: `"object"`. single nested block, Optional.

Azure VNet Site Type. Cloud Connect Azure VNet Site Type.

Upstream description:

Cloud Connect Azure VNet Site Type.

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
azure_vnet_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-80867b6879f9e0eddf7e5d76e5455bb617303f4e5b678fde17c7504b98c98a37"></a>

## Direct properties — azure_vnet_site / 6e1d91178d66 / 3

- [site](resources--cloud_connect--reference--group-001.md#canonical-64437e779702c3a7ad90e8cb1edfe8547e455fbc9d6a458f6c93ae4d250a4226): complete subsection reference.

- [vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-ca3cfe872b8c1b0b716b1d5cfb41e19667e04270edd645c96009b920beff4010): complete subsection reference.

<a id="canonical-f0cfe8969bdcda0b7464b0e4df7599a204ecef6d8120a5eb4c961511c007d377"></a>

## Next pages — azure_vnet_site / 6e1d91178d66 / 4

- [azure_vnet_site.site](resources--cloud_connect--reference--group-001.md#canonical-64437e779702c3a7ad90e8cb1edfe8547e455fbc9d6a458f6c93ae4d250a4226)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-ca3cfe872b8c1b0b716b1d5cfb41e19667e04270edd645c96009b920beff4010)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-64437e779702c3a7ad90e8cb1edfe8547e455fbc9d6a458f6c93ae4d250a4226"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a628624be767b85c415e05297fb3d9020bd7c4807183e0fe8884af56221b3e6"></a>

## azure_vnet_site.site — azure_vnet_site.site / 78d83adc122d / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43)
- azure_vnet_site.site

<a id="canonical-2aefc4abf03423e1eaf440e54589553802d013563dbba8fac3332e9c0e20ecc0"></a>

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

<a id="canonical-c64beabdbd308b2cc65d76b617bf58012ccd2dfbaeeb84caeb929108988be89c"></a>

## Direct properties — azure_vnet_site.site / 78d83adc122d / 3

<a id="canonical-bffab6a46facd91cdf928e06810503f98f25b878e3ff084bff7215dedaa56f8c"></a>

<a id="canonical-03c527fb3a70ea8afd7f5f984af0db8a9b1a37c8f871d25999d43b90061625c4"></a>

## name property — azure_vnet_site.site / 78d83adc122d / 4

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

<a id="canonical-22dfd1b806784442e618ca10ba6b83b0653b508418087b3d7c2c64e3671f3e0d"></a>

<a id="canonical-f3aeb2b6bcdb93261658410fe580a14c616cf7c17d8126548ef3190404842ae9"></a>

## namespace property — azure_vnet_site.site / 78d83adc122d / 5

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

<a id="canonical-d71db0adad8c563e4fc9a6ac1f3312b1d94530dc24b87f834ece57dd175c661a"></a>

<a id="canonical-04edad762b2da68f061548e770fa452f5045ae88cf67652bd067230fd0e55701"></a>

## tenant property — azure_vnet_site.site / 78d83adc122d / 6

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

<a id="canonical-3ae13c5266b37fdeefdb4abec02f626648e58fd36313801e11f67ab7db351510"></a>

## Next pages — azure_vnet_site.site / 78d83adc122d / 7

- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-ca3cfe872b8c1b0b716b1d5cfb41e19667e04270edd645c96009b920beff4010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9abf569e1421a4017011e084b23fefd3e2c7f6916e46f5306b6064dae295a621"></a>

## azure_vnet_site.vnet_attachments — azure_vnet_site.vnet_attachments / c0db2d9aa7e5 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43)
- azure_vnet_site.vnet_attachments

<a id="canonical-c3922d5543c74965f926ebb7e2a61f785bc5fce38caf52b36c55bde7794fc50e"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vnet attachments.

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
vnet_attachments {
  # Configure direct properties listed below.
}
```

<a id="canonical-1aa383691cd10024e8709fa498e76b708cf4e00dd26c9c6eb7d2095f5cab1436"></a>

## Direct properties — azure_vnet_site.vnet_attachments / c0db2d9aa7e5 / 3

- [vnet_list](resources--cloud_connect--reference--group-001.md#canonical-938a29b431fdffec77a91374943bd5da7245b16da8cbd0f965308df6b7c0882a): complete subsection reference.

<a id="canonical-0e48bfc02d562d55d09d87f3bcff7f92db316394cf4ee653f46f217e8e66cbda"></a>

## Next pages — azure_vnet_site.vnet_attachments / c0db2d9aa7e5 / 4

- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-938a29b431fdffec77a91374943bd5da7245b16da8cbd0f965308df6b7c0882a)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-938a29b431fdffec77a91374943bd5da7245b16da8cbd0f965308df6b7c0882a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5544c1fb6c0ad8ffbcfeb41900bd9c0c5a11dc8de094bf29ba54b1ebb031255c"></a>

## azure_vnet_site.vnet_attachments.vnet_list — azure_vnet_site.vnet_attachments.vnet_list / a794acbb50f9 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-ca3cfe872b8c1b0b716b1d5cfb41e19667e04270edd645c96009b920beff4010)
- azure_vnet_site.vnet_attachments.vnet_list

<a id="canonical-3327e369bba0f8cb3ee193701e1fd4e398e0ac41ef60975aced4efd9760ecbf3"></a>

Type: `"object"`. list nested block, Optional.

VNet List. Collection of items or values

Upstream description:

Collection of items or values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("subscription_id",
    "vnet_id"),
  validators.ConflictingListObjectAttributes("custom_routing",
    "default_route"),
  validators.ConflictingListObjectAttributes("custom_routing",
    "manual_routing"),
  validators.ConflictingListObjectAttributes("default_route",
    "manual_routing")}
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

Terraform syntax:

```terraform
vnet_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-490d67d88b6cc4d90ac7fbf7d85531789a2c64e61c45b96bec2eb56e6b796e5a"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list / a794acbb50f9 / 3

- [custom_routing](resources--cloud_connect--reference--group-001.md#canonical-a050099425ada4751ee15b1be2cf76d5ce5b91d6c3c818992b81f473e845d586): complete subsection reference.

- [default_route](resources--cloud_connect--reference--group-001.md#canonical-e8878d6388dccdb27c9ea758f62b8c0d25b959f35eed54fa99011ab2254e5847): complete subsection reference.

- [labels](resources--cloud_connect--reference--group-001.md#canonical-a2b1a9f7b63a926d6cdffc0a80f64db442bd928fa7c7d1efc36141d796769b08): complete subsection reference.

- [manual_routing](resources--cloud_connect--reference--group-001.md#canonical-e439e85bb2af5233bc7207d91f83be4e021926be14e6291868eac1ccd723266a): complete subsection reference.

<a id="canonical-becd1d64965e2564c1a1b98b7be266914c9596b54bafd9f01e3af09c2968ae7b"></a>

<a id="canonical-acc59557fc7142a7f64f45e8d1f05e7c5aa84d48f77b9e3dd26d6bcb50118b5c"></a>

## subscription_id property — azure_vnet_site.vnet_attachments.vnet_list / a794acbb50f9 / 4

Type: `"string"`. Optional.

Enter the Subscription ID of the VNet to be attached.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-db581eeb673f8e2f78c39ea0a99067f33fbc20d1c2377923d5f76be533c2f1d9"></a>

<a id="canonical-f42cc471d39a4bf82deb29c75264c75bca5b368ca4799ea23ab307e3bc6db6c8"></a>

## vnet_id property — azure_vnet_site.vnet_attachments.vnet_list / a794acbb50f9 / 5

Type: `"string"`. Optional.

Enter the VNet ID of the VNet to be attached in format
/&lt;resource-group-name&gt;/&lt;VNet-name&gt;.

Upstream description:

Enter the VNet ID of the VNet to be attached in format
/&lt;resource-group-name&gt;/&lt;VNet-name&gt;

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

<a id="canonical-a496b40a0bd1e78c3de4b697b98044e05bf57b8ea1e834ff8378f3647134252a"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list / a794acbb50f9 / 6

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-a050099425ada4751ee15b1be2cf76d5ce5b91d6c3c818992b81f473e845d586)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-e8878d6388dccdb27c9ea758f62b8c0d25b959f35eed54fa99011ab2254e5847)
- [azure_vnet_site.vnet_attachments.vnet_list.labels](resources--cloud_connect--reference--group-001.md#canonical-a2b1a9f7b63a926d6cdffc0a80f64db442bd928fa7c7d1efc36141d796769b08)
- [azure_vnet_site.vnet_attachments.vnet_list.manual_routing](resources--cloud_connect--reference--group-001.md#canonical-e439e85bb2af5233bc7207d91f83be4e021926be14e6291868eac1ccd723266a)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-ca3cfe872b8c1b0b716b1d5cfb41e19667e04270edd645c96009b920beff4010)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-a050099425ada4751ee15b1be2cf76d5ce5b91d6c3c818992b81f473e845d586"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-173e6f41165ce25989d18e203d87770c6b5525ca22428b159090876f5ba24b22"></a>

## azure_vnet_site.vnet_attachments.vnet_list.custom_routing — azure_vnet_site.vnet_attachments.vnet_list.custom_routing / acd9c6946837 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-ca3cfe872b8c1b0b716b1d5cfb41e19667e04270edd645c96009b920beff4010)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-938a29b431fdffec77a91374943bd5da7245b16da8cbd0f965308df6b7c0882a)
- azure_vnet_site.vnet_attachments.vnet_list.custom_routing

<a id="canonical-a53212e3807a9b65f18383bb163942a51555c053daf2a62853481dae0ca7093d"></a>

Type: `"object"`. single nested block, Optional.

List Azure Route Table with Static Route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("route_tables")}
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
custom_routing {
  # Configure direct properties listed below.
}
```

<a id="canonical-6b991e7af1403f3e31f6a74bec08407a3c1ffb2d9267556d32423ea91399ac07"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list.custom_routing / acd9c6946837 / 3

- [route_tables](resources--cloud_connect--reference--group-001.md#canonical-7f22a8a736ed6265ee7059f26b070810a97c929e209828dfa93e9ed8c5dfc4bc): complete subsection reference.

<a id="canonical-3bc9bc58adc04685c229bea3ea807c1ed0fd928db4839327faf4001a475fb4db"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list.custom_routing / acd9c6946837 / 4

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables](resources--cloud_connect--reference--group-001.md#canonical-7f22a8a736ed6265ee7059f26b070810a97c929e209828dfa93e9ed8c5dfc4bc)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-938a29b431fdffec77a91374943bd5da7245b16da8cbd0f965308df6b7c0882a)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-7f22a8a736ed6265ee7059f26b070810a97c929e209828dfa93e9ed8c5dfc4bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74b28cbf9582608296a73b3024ce1cb1d7ea4b877cea8f267a1611349a972a84"></a>

## azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables — azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables / cfba4ce52560 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-ca3cfe872b8c1b0b716b1d5cfb41e19667e04270edd645c96009b920beff4010)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-938a29b431fdffec77a91374943bd5da7245b16da8cbd0f965308df6b7c0882a)
- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-a050099425ada4751ee15b1be2cf76d5ce5b91d6c3c818992b81f473e845d586)
- azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables

<a id="canonical-99c607148056ab4255743d9b490fd32c09debf724b8b126a62b60bc9075d73f4"></a>

Type: `"object"`. list nested block, Optional.

List of route tables with static routes. Route Tables with static routes.

Upstream description:

Route Tables with static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("static_routes")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 200,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 200,
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
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
route_tables {
  # Configure direct properties listed below.
}
```

<a id="canonical-2465e9cd38626de78c807e1ec6f57d352a062b62119b8f9562b921162dcdb87e"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables / cfba4ce52560 / 3

<a id="canonical-e53e36049ac6b61a58d7e967b3f2415292a7025496178e09491bad444fe9029d"></a>

<a id="canonical-b9e3f9f9517137a1e3aa46fe08c73fed6b010b5f2c31b51ba32d9e7fdc57dfb5"></a>

## route_table_id property — azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables / cfba4ce52560 / 4

Type: `"string"`. Optional.

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;.

Upstream description:

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;

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
    },
    "pattern": "^\\\\/[-\\\\w\\\\._\\\\(\\\\)]+\\\\/[a-zA-Z0-9][a-zA-Z0-9-._]+[a-zA-Z0-9_]$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^\\\\/[-\\\\w\\\\._\\\\(\\\\)]+\\\\/[a-zA-Z0-9][a-zA-Z0-9-._]+[a-zA-Z0-9_]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^\\\\/[-\\\\w\\\\._\\\\(\\\\)]+\\\\/[a-zA-Z0-9][a-zA-Z0-9-._]+[a-zA-Z0-9_]$"
  }
}
```

<a id="canonical-db230333a0625576972fa1eb9e0449022bc14e6ef75452f8537cf4d7a3e10e17"></a>

<a id="canonical-2e1ee81292beb526c6e89cff3c40b4da251ad1489f8b6582dab40ff75a0116f4"></a>

## static_routes property — azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables / cfba4ce52560 / 5

Type: `["list", "string"]`. Optional.

Static Routes. List of Static Routes.

Upstream description:

List of Static Routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
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
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-d7802999daba45b6b395f7d59ad03dcc2d0cb3b84190d588e42c5808f80fd537"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables / cfba4ce52560 / 6

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-a050099425ada4751ee15b1be2cf76d5ce5b91d6c3c818992b81f473e845d586)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-e8878d6388dccdb27c9ea758f62b8c0d25b959f35eed54fa99011ab2254e5847"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5a2cdd880c9eb8eb6a40885376c48e3feb38f194e2aeb0488673edd6111e364"></a>

## azure_vnet_site.vnet_attachments.vnet_list.default_route — azure_vnet_site.vnet_attachments.vnet_list.default_route / 9317dfdc7d24 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-ca3cfe872b8c1b0b716b1d5cfb41e19667e04270edd645c96009b920beff4010)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-938a29b431fdffec77a91374943bd5da7245b16da8cbd0f965308df6b7c0882a)
- azure_vnet_site.vnet_attachments.vnet_list.default_route

<a id="canonical-d88ffbcc4eb21ca5b5ff7f399e5ee601f99da2c968866033a763a16201e128f3"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

Select Override Default Route Choice.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_route_tables",
    "selective_route_tables")}
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
  "x-ves-oneof-field-default_route_choice": "[\"all_route_tables\",\"selective_route_tables\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-df5f79a26e57ae24a7b5c52aab547ff82838001c9123461990124096bd4a11f1"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list.default_route / 9317dfdc7d24 / 3

- [all_route_tables](resources--cloud_connect--reference--group-001.md#canonical-b9d0aa522542b2cb9109abad2b7eeadf1af151d61a60a332eeeba501dd2a5e71): complete subsection reference.

- [selective_route_tables](resources--cloud_connect--reference--group-001.md#canonical-7cec497c2b79609b75db7753431c9b300de1d67ebcc46ccb775da1096dcf9e5c): complete subsection reference.

<a id="canonical-2c611b641599e9f79ec7ea6c34172d07015ef285c935f71333fa28bde0e87684"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list.default_route / 9317dfdc7d24 / 4

- [azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables](resources--cloud_connect--reference--group-001.md#canonical-b9d0aa522542b2cb9109abad2b7eeadf1af151d61a60a332eeeba501dd2a5e71)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables](resources--cloud_connect--reference--group-001.md#canonical-7cec497c2b79609b75db7753431c9b300de1d67ebcc46ccb775da1096dcf9e5c)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-938a29b431fdffec77a91374943bd5da7245b16da8cbd0f965308df6b7c0882a)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-b9d0aa522542b2cb9109abad2b7eeadf1af151d61a60a332eeeba501dd2a5e71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cdf394b01e0cb1e94704bd35222b5f6e649474f3342b5c6e6397b31886b90b6"></a>

## azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables — azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables / 684c6996188b / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-ca3cfe872b8c1b0b716b1d5cfb41e19667e04270edd645c96009b920beff4010)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-938a29b431fdffec77a91374943bd5da7245b16da8cbd0f965308df6b7c0882a)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-e8878d6388dccdb27c9ea758f62b8c0d25b959f35eed54fa99011ab2254e5847)
- azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables

<a id="canonical-f584a6877e4c1b7dab7ce5d3b76915c027570617c031e6d494e6d1bf90f6d8bf"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all route tables.

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
all_route_tables = {}
```

<a id="canonical-fcbe385fdffc59e206d3f1a13f0c4db7c2a77653b934b442b268175fbc82cfd3"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables / 684c6996188b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9ea7747e9c3f56767b52a2f938ad8e3dea24b3f27d44a99b4efbb433286a6ecc"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables / 684c6996188b / 4

- [azure_vnet_site.vnet_attachments.vnet_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-e8878d6388dccdb27c9ea758f62b8c0d25b959f35eed54fa99011ab2254e5847)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-7cec497c2b79609b75db7753431c9b300de1d67ebcc46ccb775da1096dcf9e5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2fddfe3c81ef6639247fd220f2cd592487e39cbd873f0cc1ce0bde326720ee07"></a>

## azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables — azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables / 9e9e81ad7e92 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-ca3cfe872b8c1b0b716b1d5cfb41e19667e04270edd645c96009b920beff4010)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-938a29b431fdffec77a91374943bd5da7245b16da8cbd0f965308df6b7c0882a)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-e8878d6388dccdb27c9ea758f62b8c0d25b959f35eed54fa99011ab2254e5847)
- azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables

<a id="canonical-1c9f55a5aa55808d1ed1802be993674fa76e2ce7d87e2a9afbb346ad4fddef72"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for selective route tables.

Upstream description:

Azure Route Table.

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
selective_route_tables {
  # Configure direct properties listed below.
}
```

<a id="canonical-4b4e59274bbc53d9bee61dd3ae3387249c890bcd7c73e6b8eadbb33a987a1d10"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables / 9e9e81ad7e92 / 3

<a id="canonical-3c25b3c7eca9acfc9a1f613b9250b9fc5da5f7e492e06b6252b541c33e659c59"></a>

<a id="canonical-aea84a9bf5669467f2b28ebc9be5ce60d8187b9a439e3e4f96857a01b2d44937"></a>

## route_table_id property — azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables / 9e9e81ad7e92 / 4

Type: `["list", "string"]`. Optional.

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;.

Upstream description:

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f86d7f2cc70cd2bdca8692e9f4ff888ca9a8ad74d08644e9e2f74a65f670b509"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables / 9e9e81ad7e92 / 5

- [azure_vnet_site.vnet_attachments.vnet_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-e8878d6388dccdb27c9ea758f62b8c0d25b959f35eed54fa99011ab2254e5847)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-a2b1a9f7b63a926d6cdffc0a80f64db442bd928fa7c7d1efc36141d796769b08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc31a6d51633cf8e40d2a29afbdf5d82b0048d22c3fe8552526420cab6f6fc20"></a>

## azure_vnet_site.vnet_attachments.vnet_list.labels — azure_vnet_site.vnet_attachments.vnet_list.labels / aea5fdf6886a / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-ca3cfe872b8c1b0b716b1d5cfb41e19667e04270edd645c96009b920beff4010)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-938a29b431fdffec77a91374943bd5da7245b16da8cbd0f965308df6b7c0882a)
- azure_vnet_site.vnet_attachments.vnet_list.labels

<a id="canonical-2a098bf12b5f8b30191df07c5e8c4900dfe962103a0dd819f75fd7292ef79670"></a>

Type: `"object"`. single nested block, Optional.

Add labels for the VNet attachments. These labels can then be used in policies such as enhanced
firewall policies.

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
labels {}
```

<a id="canonical-a288b06f870f3f903f925a8da806612ea964455ef55d6677fa24e3c1b28f3357"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list.labels / aea5fdf6886a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c5c580c94214852b94127d133cfc4122acbcb5184510b416d848a854fcfc4bb1"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list.labels / aea5fdf6886a / 4

- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-938a29b431fdffec77a91374943bd5da7245b16da8cbd0f965308df6b7c0882a)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-e439e85bb2af5233bc7207d91f83be4e021926be14e6291868eac1ccd723266a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68df22988fd6ac038a9b3506d8bc555ac52978c3372c14873f9ef7b5ab7a612d"></a>

## azure_vnet_site.vnet_attachments.vnet_list.manual_routing — azure_vnet_site.vnet_attachments.vnet_list.manual_routing / 119ac54cd4b2 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-c94fa738600c41d904230e9036e8dad2e8ec2dd382c3bbd309a9098af0b5fd43)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-ca3cfe872b8c1b0b716b1d5cfb41e19667e04270edd645c96009b920beff4010)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-938a29b431fdffec77a91374943bd5da7245b16da8cbd0f965308df6b7c0882a)
- azure_vnet_site.vnet_attachments.vnet_list.manual_routing

<a id="canonical-a0738cf5eecb0b7aed17f3770124dba590687ce0dddc625829aae61335ef1402"></a>

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
manual_routing = {}
```

<a id="canonical-98a0c6e5cdd11603e9472bb1d8c15bd911ca2b640727145049701bb3d2fbcf30"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list.manual_routing / 119ac54cd4b2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e8dbd995849e7189f5f14046d8203caf53ac0231c477a48735e324a25e474d19"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list.manual_routing / 119ac54cd4b2 / 4

- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-938a29b431fdffec77a91374943bd5da7245b16da8cbd0f965308df6b7c0882a)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-26858322626c00d5922e9a6ee025eab4987bc0ae4eb5e0a79f4eb3ac39f170ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a60d816336bc7595ce4e5bd2de5e2867f720d0e0f1be5f854a6bc8cb27c01526"></a>

## segment — segment / c9b33b1e0208 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- segment

<a id="canonical-697ce5a1b833b2b846c62ce34d960c0e126fbb1a6d81645bbb267242b855839b"></a>

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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-14f19ddda4adeb2987a74109be60251c1589c90d47a968df9a3e53c9892c974f"></a>

## Direct properties — segment / c9b33b1e0208 / 3

<a id="canonical-00431dd3798be33f6ee5fbdc96467e56509937e5ef643c4736a9680e54ab2e38"></a>

<a id="canonical-7888394f20b17167f9804cf806d9b8dbffb0159f1893ef9c120c63266bbc6e65"></a>

## name property — segment / c9b33b1e0208 / 4

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

<a id="canonical-b5850f3036636a37990b13c846f297e1fe9a810d238551f4ca009dea991d41d8"></a>

<a id="canonical-7a7c5bdd13edc070e4628c5afe720206f9f0644ea7157cef5321a2308776224a"></a>

## namespace property — segment / c9b33b1e0208 / 5

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

<a id="canonical-d6118005356ca880a19fd7b7b474b8462067eafe1953df34a54900bbb51a487a"></a>

<a id="canonical-4f4a7b02e7cdc46336b8fa0796f8ea48fc3424a49a56e09705a848279b075b2f"></a>

## tenant property — segment / c9b33b1e0208 / 6

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

<a id="canonical-1034a61c768ce9f0458868932bd63d6f94580785fd4ec61154ae4c0925234837"></a>

## Next pages — segment / c9b33b1e0208 / 7

- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-120208e71f2e652ac918b8a30a42d0ca791c2ea433e8d57c4de612854efd777e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58cf0e2258a9e2fe1cf97efecda87e359e25f84f92da70fde4d4114bd4068b42"></a>

## timeouts — timeouts / c745cf647c1d / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- timeouts

<a id="canonical-916da7551befb29d84365520eca4f474cca8825399f98dc77d881a83d5973948"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-f4eece0b8ab452d0455f186e61c4e49068eafafa052cb6ffa4fb805a7b585b04"></a>

## Direct properties — timeouts / c745cf647c1d / 3

<a id="canonical-f2ead0aa4006ce0692b09e2ed53c39966b6dfc642258797e802509bf17929431"></a>

<a id="canonical-d9b2607b718a10d95c6d2805ac1ea15e84227d197e787e3819ce84b2a234a22f"></a>

## create property — timeouts / c745cf647c1d / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f6f102ffd5ade80916ee8d977393870370fd8c0efb94f3a13b6bff3ce879a45a"></a>

<a id="canonical-b5c98ba83e285667208a7236f5d6ac12399737d00eb7e0fd1458d66bfe68e027"></a>

## delete property — timeouts / c745cf647c1d / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-c98b87041d6430da463a334c983a9c1664eabc6a940c0bcc517b61adc1eadabb"></a>

<a id="canonical-539795fe406331ce71fe3c25bd9c772fab9151483c21bbff2d053e8ecaadffa0"></a>

## read property — timeouts / c745cf647c1d / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-a4bf29cfc566b14f6c58c89fed01d4ad1b7028cff70caf02fb6ffc7ed37d8625"></a>

<a id="canonical-4a6618a27a0dc05c89ae51cf4770738c8d730ab96577635a4e9955f93d884899"></a>

## update property — timeouts / c745cf647c1d / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-504d08d0ffcf1e44f071ad34c4af458387a4f9990b04222b213159374c066b50"></a>

## Next pages — timeouts / c745cf647c1d / 8

- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
