---
page_title: "xcsh_cloud_connect reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_connect reference."
---

# xcsh_cloud_connect reference

<a id="canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37f4d51030393d02a145e9abac71ff0aa9289829cf7b153424558891739f624a"></a>

## Property reference — Property reference / bdee5a34a2e5 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- Property reference

<a id="canonical-2e0956062c317338147a8b2d68f6775de72f645668c47ee3235e34e30095c333"></a>

## Direct properties — Property reference / bdee5a34a2e5 / 3

<a id="canonical-9dc5ba114be126d421f947d5cc0f23ffa26174473b9c903c8b7d42fbef5221cf"></a>

<a id="canonical-a610fa5722e73d7e2486736989474d22233eb618732034e6d8008c2a2dfa9008"></a>

## annotations property — Property reference / bdee5a34a2e5 / 4

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

- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660): complete subsection reference.

- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e): complete subsection reference.

<a id="canonical-3ac722ec401080bebed75aac37712bff384f1c109d992b782464b37302d3e025"></a>

<a id="canonical-faa24b6322bd65fb392c458170f5ad9feb3ab7f3ded6e75b181268e33753fbc3"></a>

## description property — Property reference / bdee5a34a2e5 / 5

Type: `"string"`. Computed.

Description of the CloudConnect.

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

<a id="canonical-8a1619b3b6c504f6ebddca5186cd7ad68c4aa09ce8ccc1149b21027cd23d16c7"></a>

<a id="canonical-b7967add3134662346f1e0886a2c3d98c07c3053eee73f89e9aae209492a9782"></a>

## id property — Property reference / bdee5a34a2e5 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0fe1f2473a0f59efc12ad29b88e21d089b7f355b9ada84bf2c9ec1f4ce84605e"></a>

<a id="canonical-35f0b22a3b6ea3bb0ce314fb46ef7470b715749657909e0666bed9b85673dc9d"></a>

## labels property — Property reference / bdee5a34a2e5 / 7

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

<a id="canonical-7fdfba1cae0a8c28a068cceab973de048790f0e9a15d48fb2bf10fca9ce061cf"></a>

<a id="canonical-68aeff8b58be3e78d87f10cd81e1969c76465423e684b5da205b41e5ddd802a9"></a>

## name property — Property reference / bdee5a34a2e5 / 8

Type: `"string"`. Required.

Name of the CloudConnect.

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

<a id="canonical-ac235b3e1faaa551be502dc6945dfa9334db7b82a1793f507e91f41be7a78bab"></a>

<a id="canonical-68820cbe11d00c93001291f65b59750a2f7b7c2c57464a9be889e330d4179673"></a>

## namespace property — Property reference / bdee5a34a2e5 / 9

Type: `"string"`. Required.

Namespace where the CloudConnect exists.

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

- [segment](data-sources--cloud_connect--reference--group-001.md#canonical-2dd4db43791f10dcf78cba7f07887e11b8819442a3176afb6303be1a60516481): complete subsection reference.

<a id="canonical-2e4caf3cc95be2fd48a56f60db44f460b50b3ee3ecc8f89ad65de8e418b5aa14"></a>

## All schema paths — Property reference / bdee5a34a2e5 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cloud_connect--reference--group-001.md#canonical-9dc5ba114be126d421f947d5cc0f23ffa26174473b9c903c8b7d42fbef5221cf) |
| `aws_provider` | [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3b888f9603d0919a3ffd813504d976e9a44dda6952f149179132ca445f3b7f8a) |
| `aws_provider.aws_tgw_site` | [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-3b454c64a3489b551eef961f0495a1fa8295417aa7de112ecb2f7eadd7e1d829) |
| `aws_provider.aws_tgw_site.cred` | [aws_provider.aws_tgw_site.cred](data-sources--cloud_connect--reference--group-001.md#canonical-e88ad256aeebdf4f37b6f701cb48430e7d3df6ad77a3cad6744082079bff655b) |
| `aws_provider.aws_tgw_site.cred.name` | [aws_provider.aws_tgw_site.cred.name](data-sources--cloud_connect--reference--group-001.md#canonical-686b82a11265b0c977da6fe80ee8aaf6456e9ea0693eceb5d17818d5ab6db2c3) |
| `aws_provider.aws_tgw_site.cred.namespace` | [aws_provider.aws_tgw_site.cred.namespace](data-sources--cloud_connect--reference--group-001.md#canonical-8a1581cb3f411b63f0f3fe13109d193549f45a08b9b1b7644dbf9880fb00ef00) |
| `aws_provider.aws_tgw_site.cred.tenant` | [aws_provider.aws_tgw_site.cred.tenant](data-sources--cloud_connect--reference--group-001.md#canonical-cba665233c929e51a743458c906dfe5826b0ae7605db2ea8874d8e62d8da9aa4) |
| `aws_provider.aws_tgw_site.site` | [aws_provider.aws_tgw_site.site](data-sources--cloud_connect--reference--group-001.md#canonical-8d640e09432d1c18e3129b0b899df00c2637df3cf2bd66dcd9504b88fb20ad5e) |
| `aws_provider.aws_tgw_site.site.name` | [aws_provider.aws_tgw_site.site.name](data-sources--cloud_connect--reference--group-001.md#canonical-21ae6c8b358ece48438d8a45a2667152f4ff974cb418a1d3fa13973c482a67b9) |
| `aws_provider.aws_tgw_site.site.namespace` | [aws_provider.aws_tgw_site.site.namespace](data-sources--cloud_connect--reference--group-001.md#canonical-77d167f08ed6d02352a7c369bef313e75d472917b476c74438f7c85f0ba0c7ac) |
| `aws_provider.aws_tgw_site.site.tenant` | [aws_provider.aws_tgw_site.site.tenant](data-sources--cloud_connect--reference--group-001.md#canonical-08886b093a4a0e93cde12b4383b718037bc17546e5bf0f4c4c167db9a55484f6) |
| `aws_provider.aws_tgw_site.vpc_attachments` | [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-74eb06d744a0b955552cfe4ee2ec8df6990f5122a87239dcb94270b71afc4131) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-c37a5c0ed0113ab1fcc4be537f4e9f6ad0dfd96e030c3598da00367a8cd61842) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-10ac35c17b3b9e4a6a09f8c9e3f63e0b997f353f02d82929155f72ccab7d33ba) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-c4963d89729e77ed0529a978f9b6ba917e8a9aeeaea9291fb06f21a78a3401d5) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.route_table_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.route_table_id](data-sources--cloud_connect--reference--group-001.md#canonical-d2f2821ff41e9f3217b6d170271ce180f087972e5f316219a3301dfa5ab070ee) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.static_routes` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.static_routes](data-sources--cloud_connect--reference--group-001.md#canonical-4d941bcd7cf3d96cc3e26128b2ebb713108ea2ec72a1a79cc0473c8e9c5a4c24) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-e69de772962dc13c8c8df95f86427e8d40ce23bf6116b8d99fc4212d1c7d24c8) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-dd11d2902f84c0c760b5a4f331ea02b68b5b429e9e042da96b5bf98d758d46e0) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-bce247e4ba086dcf40fc4172d2e3f6122d0b4698fcda7c0676a55bad3997d63e) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables.route_table_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables.route_table_id](data-sources--cloud_connect--reference--group-001.md#canonical-9a52cce2afe126ebb3715a4219bb6a72bd6fa9d51f9a2c3fd0f7daa322730b79) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels](data-sources--cloud_connect--reference--group-001.md#canonical-71ec46545dd8e0acc625a1176ee336a636bcb22c07d321569d56caa965ba85c6) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing](data-sources--cloud_connect--reference--group-001.md#canonical-a5ea8dc68b70c06aeca332f108aec6106613c10635774fe6a46c0660f30fcc3b) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.vpc_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.vpc_id](data-sources--cloud_connect--reference--group-001.md#canonical-1e67459e0b748ed31fe434bc815700fec438a7a92fba23d72275b34e85e7e788) |
| `azure_vnet_site` | [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-b3f5de28598e2fac824cfce07c6af8f7af0bf863a4f301fcd618d84a26db73f0) |
| `azure_vnet_site.site` | [azure_vnet_site.site](data-sources--cloud_connect--reference--group-001.md#canonical-8c963232b3a3c76d9ed28290146723eb06c44774f0c1eb46409b923df2be06ca) |
| `azure_vnet_site.site.name` | [azure_vnet_site.site.name](data-sources--cloud_connect--reference--group-001.md#canonical-522339440548c776f79c8097e1f7fae1af8664112c2e8b82aa213578f6f28146) |
| `azure_vnet_site.site.namespace` | [azure_vnet_site.site.namespace](data-sources--cloud_connect--reference--group-001.md#canonical-5b442f1f109670d7d0f64ceaf045d9741e6a66dcdf8b966c4630ffde4f8c43fd) |
| `azure_vnet_site.site.tenant` | [azure_vnet_site.site.tenant](data-sources--cloud_connect--reference--group-001.md#canonical-2c8d8c4574f5c30e756f9e6e0f5fdb6a492d6414b0485ec3f8fcdbb9ee232009) |
| `azure_vnet_site.vnet_attachments` | [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-a0afa1eddf6449d5452b278e972e96299e0a94d3abab2f8f29f801ccd18d7f01) |
| `azure_vnet_site.vnet_attachments.vnet_list` | [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-c3d0204bcd0563cecbaaed988429706f09f4e4c342f024e8f576fffdd158df46) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-62a7e2e6095440e23e65d5919f6676ae88194cdef3811402b0f115f525c339ba) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-3a89b8a3f8e9ad5ac574b043856849517107b035f0e03d831a6b2afd1ee38807) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.route_table_id` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.route_table_id](data-sources--cloud_connect--reference--group-001.md#canonical-eb21414f1fb83b6c0dd5e8073b00d2b8be1ac7e9968cdc51e11d882184c65388) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.static_routes` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.static_routes](data-sources--cloud_connect--reference--group-001.md#canonical-8d60dd9e0fba3ccf731fdea572c819c7647149a1c5baf5bba1109f9126b58fa9) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route` | [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-aca8ec81e2436c80ccbbebdf326acf5fb52d4ca7b23568235a231751b840dd16) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-2d6d6520c805cd301462417bc0d87e3e26ac62f3cfa3de44410adb3e346acc29) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-675830ca175ff86cf7394cd88752bb6fe4239a070346adccf668aeb27a1b4a68) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables.route_table_id` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables.route_table_id](data-sources--cloud_connect--reference--group-001.md#canonical-7f5576d407b8502fa506928e8d4de11dd6834657a1e616adc84a7d75cdfaada7) |
| `azure_vnet_site.vnet_attachments.vnet_list.labels` | [azure_vnet_site.vnet_attachments.vnet_list.labels](data-sources--cloud_connect--reference--group-001.md#canonical-4b993aee1d641aa6362f225631781a4e24d2f93ee29317fe0a254173507dcd45) |
| `azure_vnet_site.vnet_attachments.vnet_list.manual_routing` | [azure_vnet_site.vnet_attachments.vnet_list.manual_routing](data-sources--cloud_connect--reference--group-001.md#canonical-745cad365c4d1636337902232751cd481b6b842cccf512b9c835742856bf820d) |
| `azure_vnet_site.vnet_attachments.vnet_list.subscription_id` | [azure_vnet_site.vnet_attachments.vnet_list.subscription_id](data-sources--cloud_connect--reference--group-001.md#canonical-28e78325cb8b13ad522ab6ccf47907940917c409c7748568eb49fede97369a34) |
| `azure_vnet_site.vnet_attachments.vnet_list.vnet_id` | [azure_vnet_site.vnet_attachments.vnet_list.vnet_id](data-sources--cloud_connect--reference--group-001.md#canonical-3bc943bfc1eab2f53efc99ac6626c0256ad00feb32e13afe28cb9446e8db57b5) |
| `description` | [description](data-sources--cloud_connect--reference--group-001.md#canonical-3ac722ec401080bebed75aac37712bff384f1c109d992b782464b37302d3e025) |
| `id` | [id](data-sources--cloud_connect--reference--group-001.md#canonical-8a1619b3b6c504f6ebddca5186cd7ad68c4aa09ce8ccc1149b21027cd23d16c7) |
| `labels` | [labels](data-sources--cloud_connect--reference--group-001.md#canonical-0fe1f2473a0f59efc12ad29b88e21d089b7f355b9ada84bf2c9ec1f4ce84605e) |
| `name` | [name](data-sources--cloud_connect--reference--group-001.md#canonical-7fdfba1cae0a8c28a068cceab973de048790f0e9a15d48fb2bf10fca9ce061cf) |
| `namespace` | [namespace](data-sources--cloud_connect--reference--group-001.md#canonical-ac235b3e1faaa551be502dc6945dfa9334db7b82a1793f507e91f41be7a78bab) |
| `segment` | [segment](data-sources--cloud_connect--reference--group-001.md#canonical-79960451edbffdb34bfc7f74b628cc679a19347416f3d33ca4f17f5328b16e46) |
| `segment.name` | [segment.name](data-sources--cloud_connect--reference--group-001.md#canonical-584cf8a5766cce73995acd0137492ec82d3a37e86b9653d25292acde6ba85453) |
| `segment.namespace` | [segment.namespace](data-sources--cloud_connect--reference--group-001.md#canonical-9d641c9fa4bb4e89a670bf6e849e20164b23b4c2e0dc78a8e299fab2a0e0d127) |
| `segment.tenant` | [segment.tenant](data-sources--cloud_connect--reference--group-001.md#canonical-c699a8b1ff0284a2e110ead01db43a09edc324b35fbc39ec3342ebf7fe5128ea) |

<a id="canonical-2b8fbf518747cc94013be5916cd36afee851d8d2675a318392a16d5f34d824de"></a>

## Next pages — Property reference / bdee5a34a2e5 / 11

- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e)
- [segment](data-sources--cloud_connect--reference--group-001.md#canonical-2dd4db43791f10dcf78cba7f07887e11b8819442a3176afb6303be1a60516481)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43d2072518c9020a574d78c964bd1fac15d9e4bd4f3f8ef1828f33775c781590"></a>

## aws_provider — aws_provider / 9b5b32d1a14f / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- aws_provider

<a id="canonical-3b888f9603d0919a3ffd813504d976e9a44dda6952f149179132ca445f3b7f8a"></a>

Type: `"single"`. Computed.

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

- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3b888f9603d0919a3ffd813504d976e9a44dda6952f149179132ca445f3b7f8a)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-b3f5de28598e2fac824cfce07c6af8f7af0bf863a4f301fcd618d84a26db73f0)

Select alternatives according to the provider validators above.

<a id="canonical-8f50e4404964c483faa6ef601433332dc7c24274c4ef89bf1980ec69982ac99b"></a>

## Direct properties — aws_provider / 9b5b32d1a14f / 3

- [aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62): complete subsection reference.

<a id="canonical-8e4ffb0866df68a6ea58c7d32d7735cff18136fc484c5dfcce72f77815177c46"></a>

## Next pages — aws_provider / 9b5b32d1a14f / 4

- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ee4fa43fb5f2986273802fad3a70a5c36ac595e8f3e74291d3894ced3071a4c"></a>

## aws_provider.aws_tgw_site — aws_provider.aws_tgw_site / a07f46500d41 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660)
- aws_provider.aws_tgw_site

<a id="canonical-3b454c64a3489b551eef961f0495a1fa8295417aa7de112ecb2f7eadd7e1d829"></a>

Type: `"single"`. Computed.

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

<a id="canonical-13099538db7fb92984670f1b3d34fa499309fa01863f07e9171934cc1055f66c"></a>

## Direct properties — aws_provider.aws_tgw_site / a07f46500d41 / 3

- [cred](data-sources--cloud_connect--reference--group-001.md#canonical-cc3a844e7cabac2a5f405d05ec4b559f6d3e42fc21d51e83c1504636995de6fe): complete subsection reference.

- [site](data-sources--cloud_connect--reference--group-001.md#canonical-3d64f567101ca53f1008c66a2671ebfc2b74d13a0cf04196d53ba7245a55ea31): complete subsection reference.

- [vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-e0d630ffe600742752a7c51b9e883c5ae2d244b8d82898d8d45636bc6fbe3077): complete subsection reference.

<a id="canonical-ee3fce788d0c0b26bd565b6677993c8dab65d11a355e85f4a0b5fbaaf04b5f9a"></a>

## Next pages — aws_provider.aws_tgw_site / a07f46500d41 / 4

- [aws_provider.aws_tgw_site.cred](data-sources--cloud_connect--reference--group-001.md#canonical-cc3a844e7cabac2a5f405d05ec4b559f6d3e42fc21d51e83c1504636995de6fe)
- [aws_provider.aws_tgw_site.site](data-sources--cloud_connect--reference--group-001.md#canonical-3d64f567101ca53f1008c66a2671ebfc2b74d13a0cf04196d53ba7245a55ea31)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-e0d630ffe600742752a7c51b9e883c5ae2d244b8d82898d8d45636bc6fbe3077)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-cc3a844e7cabac2a5f405d05ec4b559f6d3e42fc21d51e83c1504636995de6fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21b5951d1dda9da23e878a521cafaaa3e12e0bf8a0c90a9a838397ff0ce6da37"></a>

## aws_provider.aws_tgw_site.cred — aws_provider.aws_tgw_site.cred / defa245eb916 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- aws_provider.aws_tgw_site.cred

<a id="canonical-e88ad256aeebdf4f37b6f701cb48430e7d3df6ad77a3cad6744082079bff655b"></a>

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

<a id="canonical-d3875cd78686a0b9cee739b35b40d97af483f04c473b2d17cfd82c0e0d214b41"></a>

## Direct properties — aws_provider.aws_tgw_site.cred / defa245eb916 / 3

<a id="canonical-686b82a11265b0c977da6fe80ee8aaf6456e9ea0693eceb5d17818d5ab6db2c3"></a>

<a id="canonical-c16c3ed404ee97b17d14511cd4386eb7b8dc5ae5e1ac80b8cfa134e16aad12f0"></a>

## name property — aws_provider.aws_tgw_site.cred / defa245eb916 / 4

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

<a id="canonical-8a1581cb3f411b63f0f3fe13109d193549f45a08b9b1b7644dbf9880fb00ef00"></a>

<a id="canonical-bc1bb42c86139b967abc3ada300aba87dd4192ef0ae7ae7e01a851ea225333ca"></a>

## namespace property — aws_provider.aws_tgw_site.cred / defa245eb916 / 5

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

<a id="canonical-cba665233c929e51a743458c906dfe5826b0ae7605db2ea8874d8e62d8da9aa4"></a>

<a id="canonical-d6f56f6e2e572fcb315feda213726fe41d3b89195c7e440d3b93e4ce1e1a2c9d"></a>

## tenant property — aws_provider.aws_tgw_site.cred / defa245eb916 / 6

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

<a id="canonical-0117b89c34cff9303f09fa6cb64fc5caefd7527a8c8dd4f8a420dcc2c4f2f33f"></a>

## Next pages — aws_provider.aws_tgw_site.cred / defa245eb916 / 7

- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-3d64f567101ca53f1008c66a2671ebfc2b74d13a0cf04196d53ba7245a55ea31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1403c711be1a09328bfb412cd0efda02fae7695a16b932bcecba89f9f88e95e"></a>

## aws_provider.aws_tgw_site.site — aws_provider.aws_tgw_site.site / 2b2f28eafbe0 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- aws_provider.aws_tgw_site.site

<a id="canonical-8d640e09432d1c18e3129b0b899df00c2637df3cf2bd66dcd9504b88fb20ad5e"></a>

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

<a id="canonical-18e4e6f46b68b2fbdf30bfab8b7db251be1bbfea95825d35a8e65eff4d51771a"></a>

## Direct properties — aws_provider.aws_tgw_site.site / 2b2f28eafbe0 / 3

<a id="canonical-21ae6c8b358ece48438d8a45a2667152f4ff974cb418a1d3fa13973c482a67b9"></a>

<a id="canonical-643c2914f6db44406871b0a0b714874ed3eb0680921f933ad68f5667002d417d"></a>

## name property — aws_provider.aws_tgw_site.site / 2b2f28eafbe0 / 4

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

<a id="canonical-77d167f08ed6d02352a7c369bef313e75d472917b476c74438f7c85f0ba0c7ac"></a>

<a id="canonical-2faf09c8868f2a350b397725920caecb41f9c335aa0a305731fbc072bb3f3b9d"></a>

## namespace property — aws_provider.aws_tgw_site.site / 2b2f28eafbe0 / 5

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

<a id="canonical-08886b093a4a0e93cde12b4383b718037bc17546e5bf0f4c4c167db9a55484f6"></a>

<a id="canonical-d5ced96e82d8a8febdf628748bbcbcc2558409252b0ecca249d7133174f4dde0"></a>

## tenant property — aws_provider.aws_tgw_site.site / 2b2f28eafbe0 / 6

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

<a id="canonical-9d4e4239f52c0abb3d8a66a2cffb9a2d22d9ab74c296ff72541bbd103a36d295"></a>

## Next pages — aws_provider.aws_tgw_site.site / 2b2f28eafbe0 / 7

- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-e0d630ffe600742752a7c51b9e883c5ae2d244b8d82898d8d45636bc6fbe3077"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-427ae4ea47f727c3d452f69ae4f456d27a4442f39d58ff6302ee7421a94ae1c6"></a>

## aws_provider.aws_tgw_site.vpc_attachments — aws_provider.aws_tgw_site.vpc_attachments / d8b0391de7cb / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- aws_provider.aws_tgw_site.vpc_attachments

<a id="canonical-74eb06d744a0b955552cfe4ee2ec8df6990f5122a87239dcb94270b71afc4131"></a>

Type: `"single"`. Computed.

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

<a id="canonical-6e6fafd840f84e4e20e72ca3ddc43ebde8c68c25bb2ffaa69e7be17a6c2bf6d7"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments / d8b0391de7cb / 3

- [vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-dcd46012f3a8c031918fa412f8a6540d38146b4eba2e2fdb4779b212fc9ddb13): complete subsection reference.

<a id="canonical-90bc47c8dab29bf770834340be588ccc1220b50caabc71338de43165dfb813f3"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments / d8b0391de7cb / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-dcd46012f3a8c031918fa412f8a6540d38146b4eba2e2fdb4779b212fc9ddb13)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-dcd46012f3a8c031918fa412f8a6540d38146b4eba2e2fdb4779b212fc9ddb13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c49890ce631db38a726ba9f703c16d6f46c73593f008e007aa106103d96d413b"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list — aws_provider.aws_tgw_site.vpc_attachments.vpc_list / 94da5c828a8d / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-e0d630ffe600742752a7c51b9e883c5ae2d244b8d82898d8d45636bc6fbe3077)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list

<a id="canonical-c37a5c0ed0113ab1fcc4be537f4e9f6ad0dfd96e030c3598da00367a8cd61842"></a>

Type: `"list"`. Computed.

VPC List. Collection of items or values

Upstream description:

Collection of items or values

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

<a id="canonical-efe062361498abb201a98c2fbae1d6e9dfbc8b1ad198249ec1672c8351230418"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list / 94da5c828a8d / 3

- [custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-dd413ac4b5e5dff1e7e2e0a862f3ca503831f922d2e55dc6226136178ef5646a): complete subsection reference.

- [default_route](data-sources--cloud_connect--reference--group-001.md#canonical-3fbe211d40a8c76b2230b039de6f0eff8425453f18f675893bd054ff89d60a5c): complete subsection reference.

- [labels](data-sources--cloud_connect--reference--group-001.md#canonical-bda505305734363ac4ff29b94f8942d0f5af411e18c21f22fb355d0b425bfe09): complete subsection reference.

- [manual_routing](data-sources--cloud_connect--reference--group-001.md#canonical-804e4ffd7eef3cd4310de1a798ce7a6b9fec284faf9918da1188b4f031e7dd1d): complete subsection reference.

<a id="canonical-1e67459e0b748ed31fe434bc815700fec438a7a92fba23d72275b34e85e7e788"></a>

<a id="canonical-c0ff556a234ee9748ef96f78888ca997ed15aff8b78280e79b701498ad6e4cdc"></a>

## vpc_id property — aws_provider.aws_tgw_site.vpc_attachments.vpc_list / 94da5c828a8d / 4

Type: `"string"`. Computed.

Enter the VPC ID of the VPC to be attached.

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

<a id="canonical-f780ea1f69e57234d6d39e0055999051187b683cff82ac1d3092e8fa61f574ae"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list / 94da5c828a8d / 5

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-dd413ac4b5e5dff1e7e2e0a862f3ca503831f922d2e55dc6226136178ef5646a)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-3fbe211d40a8c76b2230b039de6f0eff8425453f18f675893bd054ff89d60a5c)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels](data-sources--cloud_connect--reference--group-001.md#canonical-bda505305734363ac4ff29b94f8942d0f5af411e18c21f22fb355d0b425bfe09)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing](data-sources--cloud_connect--reference--group-001.md#canonical-804e4ffd7eef3cd4310de1a798ce7a6b9fec284faf9918da1188b4f031e7dd1d)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-e0d630ffe600742752a7c51b9e883c5ae2d244b8d82898d8d45636bc6fbe3077)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-dd413ac4b5e5dff1e7e2e0a862f3ca503831f922d2e55dc6226136178ef5646a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e020342d4d7b8a0ef09c27c1633e7c8a7dcecf1fd61515033a1f2f8144b88003"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing / 839770e30b02 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-e0d630ffe600742752a7c51b9e883c5ae2d244b8d82898d8d45636bc6fbe3077)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-dcd46012f3a8c031918fa412f8a6540d38146b4eba2e2fdb4779b212fc9ddb13)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing

<a id="canonical-10ac35c17b3b9e4a6a09f8c9e3f63e0b997f353f02d82929155f72ccab7d33ba"></a>

Type: `"single"`. Computed.

AWS Route Table List. AWS Route Table List.

Upstream description:

AWS Route Table List.

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

<a id="canonical-d2ea93032155ff1e5c8178fc13c725b512171ec9ae0ebca93828809690c1894b"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing / 839770e30b02 / 3

- [route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-354019aa6cff4d99c96bd2774906c040ae97822496214c5aa4482a235d5375c1): complete subsection reference.

<a id="canonical-5d91e2f51ad04cc312de185d4c1b2721d406ddf1be99d9b480bb6e118e53b63e"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing / 839770e30b02 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-354019aa6cff4d99c96bd2774906c040ae97822496214c5aa4482a235d5375c1)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-dcd46012f3a8c031918fa412f8a6540d38146b4eba2e2fdb4779b212fc9ddb13)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-354019aa6cff4d99c96bd2774906c040ae97822496214c5aa4482a235d5375c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4e18a38789fa6810eb293afab6470ad5cb5acf957a09274bbde36491f067a92"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables / c0484d18124c / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-e0d630ffe600742752a7c51b9e883c5ae2d244b8d82898d8d45636bc6fbe3077)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-dcd46012f3a8c031918fa412f8a6540d38146b4eba2e2fdb4779b212fc9ddb13)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-dd413ac4b5e5dff1e7e2e0a862f3ca503831f922d2e55dc6226136178ef5646a)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables

<a id="canonical-c4963d89729e77ed0529a978f9b6ba917e8a9aeeaea9291fb06f21a78a3401d5"></a>

Type: `"list"`. Computed.

List of route tables. Route Tables.

Upstream description:

Route Tables.

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

<a id="canonical-a4a69c6356d0e7a4f187905f47575aaa05f97ff9346fe573c79f550d25ff5f6d"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables / c0484d18124c / 3

<a id="canonical-d2f2821ff41e9f3217b6d170271ce180f087972e5f316219a3301dfa5ab070ee"></a>

<a id="canonical-fcd808963c10703fed0535be8187834a79ecaa5f97438b385b9b9922b0426eea"></a>

## route_table_id property — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables / c0484d18124c / 4

Type: `"string"`. Computed.

Route table ID. Route table ID.

Upstream description:

Route table ID.

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

<a id="canonical-4d941bcd7cf3d96cc3e26128b2ebb713108ea2ec72a1a79cc0473c8e9c5a4c24"></a>

<a id="canonical-92ce6e51f075e1435e3c3d4400b6091d3311645280224b29174eee13e9fddad3"></a>

## static_routes property — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables / c0484d18124c / 5

Type: `["list", "string"]`. Computed.

Static Routes. List of Static Routes.

Upstream description:

List of Static Routes.

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

<a id="canonical-22a7e74def3824c5edf17c3822df7b487c30a119a958a48e95430954147fbf71"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables / c0484d18124c / 6

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-dd413ac4b5e5dff1e7e2e0a862f3ca503831f922d2e55dc6226136178ef5646a)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-3fbe211d40a8c76b2230b039de6f0eff8425453f18f675893bd054ff89d60a5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05c236ba329b6337ecac20705ba5533fa7812b7c10e30845ea4a2c805fb2884d"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route / 969a8efc033e / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-e0d630ffe600742752a7c51b9e883c5ae2d244b8d82898d8d45636bc6fbe3077)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-dcd46012f3a8c031918fa412f8a6540d38146b4eba2e2fdb4779b212fc9ddb13)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route

<a id="canonical-e69de772962dc13c8c8df95f86427e8d40ce23bf6116b8d99fc4212d1c7d24c8"></a>

Type: `"single"`. Computed.

Configuration parameter for default route.

Upstream description:

Select Override Default Route Choice.

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

<a id="canonical-b346566cbdff1b9bc2ae3c4745389f25d9b83b00ff35a4318a5a6ad87fcddabb"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route / 969a8efc033e / 3

- [all_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-19146b716833020ecd9c27d03cb151ccd1efbb2ae5de32c25b28314cbd8a4fe4): complete subsection reference.

- [selective_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-3f3b3426229c48630a809679afb68dc82175f895931ae30319ffd816f20e325a): complete subsection reference.

<a id="canonical-6d650b194ecf6af1c8d51f242a146135dbf82d7688121a05840676c819f0656d"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route / 969a8efc033e / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-19146b716833020ecd9c27d03cb151ccd1efbb2ae5de32c25b28314cbd8a4fe4)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-3f3b3426229c48630a809679afb68dc82175f895931ae30319ffd816f20e325a)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-dcd46012f3a8c031918fa412f8a6540d38146b4eba2e2fdb4779b212fc9ddb13)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-19146b716833020ecd9c27d03cb151ccd1efbb2ae5de32c25b28314cbd8a4fe4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3831b63a36ea4f7fe1cf12d51a391c2f5745338e27a8ee5b2506ac1b0f9a2bd7"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_table / b86fa5e2f5d4 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-e0d630ffe600742752a7c51b9e883c5ae2d244b8d82898d8d45636bc6fbe3077)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-dcd46012f3a8c031918fa412f8a6540d38146b4eba2e2fdb4779b212fc9ddb13)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-3fbe211d40a8c76b2230b039de6f0eff8425453f18f675893bd054ff89d60a5c)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables

<a id="canonical-dd11d2902f84c0c760b5a4f331ea02b68b5b429e9e042da96b5bf98d758d46e0"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1158727e8619246403208914bc301d4b57d265830223148de95ec150174e6289"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_table / b86fa5e2f5d4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-143b7309312add61ff1d45e93ebbb06a1f1a2eea307d25403c41d41dd8fdec09"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_table / b86fa5e2f5d4 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-3fbe211d40a8c76b2230b039de6f0eff8425453f18f675893bd054ff89d60a5c)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-3f3b3426229c48630a809679afb68dc82175f895931ae30319ffd816f20e325a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27a14186bfa508e8a426a70b06758a8423687320eb59a87ac23201530122e64a"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route / ebd5410ba62f / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-e0d630ffe600742752a7c51b9e883c5ae2d244b8d82898d8d45636bc6fbe3077)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-dcd46012f3a8c031918fa412f8a6540d38146b4eba2e2fdb4779b212fc9ddb13)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-3fbe211d40a8c76b2230b039de6f0eff8425453f18f675893bd054ff89d60a5c)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables

<a id="canonical-bce247e4ba086dcf40fc4172d2e3f6122d0b4698fcda7c0676a55bad3997d63e"></a>

Type: `"single"`. Computed.

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

<a id="canonical-19bee49ca979b0b464f040c796595e1addd347e1b8bc69069565ec5cf320b2a4"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route / ebd5410ba62f / 3

<a id="canonical-9a52cce2afe126ebb3715a4219bb6a72bd6fa9d51f9a2c3fd0f7daa322730b79"></a>

<a id="canonical-e7c3ae36a5b2460fe1c9ddfaceca455d15fcbd7aaece3ddedf1af9dcb771daf2"></a>

## route_table_id property — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route / ebd5410ba62f / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-4fb1f234f85f2e19212f6fe0524f4fb0a8ca5502cab6c3a0402a74b886f5dec9"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route / ebd5410ba62f / 5

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-3fbe211d40a8c76b2230b039de6f0eff8425453f18f675893bd054ff89d60a5c)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-bda505305734363ac4ff29b94f8942d0f5af411e18c21f22fb355d0b425bfe09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63ec7f09452314b71433eedc349fc93f40fd3d558456b8f7a86db9d26d85fdff"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels / 2a4a2feea60c / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-e0d630ffe600742752a7c51b9e883c5ae2d244b8d82898d8d45636bc6fbe3077)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-dcd46012f3a8c031918fa412f8a6540d38146b4eba2e2fdb4779b212fc9ddb13)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels

<a id="canonical-71ec46545dd8e0acc625a1176ee336a636bcb22c07d321569d56caa965ba85c6"></a>

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

<a id="canonical-8a10ace96c424bb60e0828f738692c926212c0e112101550591636556c318575"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels / 2a4a2feea60c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d189973440b6e22644d80eca19a51fae22de72659ca2861ecd2335b10baa44da"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels / 2a4a2feea60c / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-dcd46012f3a8c031918fa412f8a6540d38146b4eba2e2fdb4779b212fc9ddb13)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-804e4ffd7eef3cd4310de1a798ce7a6b9fec284faf9918da1188b4f031e7dd1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a0ca081a5f3052826ce76268ae4e4593a9373038a992803b8869e8a8fb9152f"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing / 4f7422d79d5f / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-fa937f2fe3dc10514f711a0077349e5ac5b6679299ff72145b5b12ef62cc4660)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-64846d056254e5cb98e7bd855165934517b1fc1aed266c596c7f9f39ec012c62)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-e0d630ffe600742752a7c51b9e883c5ae2d244b8d82898d8d45636bc6fbe3077)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-dcd46012f3a8c031918fa412f8a6540d38146b4eba2e2fdb4779b212fc9ddb13)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing

<a id="canonical-a5ea8dc68b70c06aeca332f108aec6106613c10635774fe6a46c0660f30fcc3b"></a>

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

<a id="canonical-d401b010d0f8ced26e986e8867d8110de63c457efde3d52fca96697ec0585731"></a>

## Direct properties — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing / 4f7422d79d5f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-72076f979f6b4ed81fa1768cc38934196cbc34f4552292a1697fd41b1f596321"></a>

## Next pages — aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing / 4f7422d79d5f / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-dcd46012f3a8c031918fa412f8a6540d38146b4eba2e2fdb4779b212fc9ddb13)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79852372d00c97e15e14564ab7a5662dd4d35b3ee3c46e1ee0808b827435c5b3"></a>

## azure_vnet_site — azure_vnet_site / 2c0cd7cf049f / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- azure_vnet_site

<a id="canonical-b3f5de28598e2fac824cfce07c6af8f7af0bf863a4f301fcd618d84a26db73f0"></a>

Type: `"single"`. Computed.

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

<a id="canonical-bf8fbceee3cbd727573aea6f937c7374253b2f9500371e41c17518bcbb578666"></a>

## Direct properties — azure_vnet_site / 2c0cd7cf049f / 3

- [site](data-sources--cloud_connect--reference--group-001.md#canonical-d22668abe84f13b4a033efd50f37b642d905fc228de1dd33327e0364d2ee24f4): complete subsection reference.

- [vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-2ce367bba074251c2ff0ecdf3d3aa3978f5eb4f4b487bce51290abdaa48643ad): complete subsection reference.

<a id="canonical-cb77783e9454d9eee62bc5f4977333e11f1dda49da949b2dd0201b87a1a5b2c5"></a>

## Next pages — azure_vnet_site / 2c0cd7cf049f / 4

- [azure_vnet_site.site](data-sources--cloud_connect--reference--group-001.md#canonical-d22668abe84f13b4a033efd50f37b642d905fc228de1dd33327e0364d2ee24f4)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-2ce367bba074251c2ff0ecdf3d3aa3978f5eb4f4b487bce51290abdaa48643ad)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-d22668abe84f13b4a033efd50f37b642d905fc228de1dd33327e0364d2ee24f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c436e695df32524a4fb6c52ebf74452e0a7485090c8d5114395021436c4148c7"></a>

## azure_vnet_site.site — azure_vnet_site.site / 0dc5cd3103f5 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e)
- azure_vnet_site.site

<a id="canonical-8c963232b3a3c76d9ed28290146723eb06c44774f0c1eb46409b923df2be06ca"></a>

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

<a id="canonical-65a97583a09ec9c35e2e96a9fe9d437f68617be56e8d8a09572b29bf4bb66177"></a>

## Direct properties — azure_vnet_site.site / 0dc5cd3103f5 / 3

<a id="canonical-522339440548c776f79c8097e1f7fae1af8664112c2e8b82aa213578f6f28146"></a>

<a id="canonical-535214a34036f46072a5d7b2c6bc48cba95fce04e8b3e87b98d23f295d354cd1"></a>

## name property — azure_vnet_site.site / 0dc5cd3103f5 / 4

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

<a id="canonical-5b442f1f109670d7d0f64ceaf045d9741e6a66dcdf8b966c4630ffde4f8c43fd"></a>

<a id="canonical-166346b439d42b2c773a9961f3c4eda7fa86698397cdea8229c6f0b3cd243ceb"></a>

## namespace property — azure_vnet_site.site / 0dc5cd3103f5 / 5

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

<a id="canonical-2c8d8c4574f5c30e756f9e6e0f5fdb6a492d6414b0485ec3f8fcdbb9ee232009"></a>

<a id="canonical-a46a79ea9dd8954a01bdcdaa106cb6f806dbbd04eb7d71aa6e3cdc93de852c85"></a>

## tenant property — azure_vnet_site.site / 0dc5cd3103f5 / 6

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

<a id="canonical-450431b6e6f75539a774138f4a8315769fab5d4d1076b22f0a3af25b76f85c93"></a>

## Next pages — azure_vnet_site.site / 0dc5cd3103f5 / 7

- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-2ce367bba074251c2ff0ecdf3d3aa3978f5eb4f4b487bce51290abdaa48643ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1059667a960866ece7b4d1ca5c5d8276df6fc00dde8a074e362a7643e6d51d07"></a>

## azure_vnet_site.vnet_attachments — azure_vnet_site.vnet_attachments / e85cc8ca97e2 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e)
- azure_vnet_site.vnet_attachments

<a id="canonical-a0afa1eddf6449d5452b278e972e96299e0a94d3abab2f8f29f801ccd18d7f01"></a>

Type: `"single"`. Computed.

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

<a id="canonical-f7405be5c86709d0338ea1a43d3ecbdf96e08d11213d048a333882aecd594b42"></a>

## Direct properties — azure_vnet_site.vnet_attachments / e85cc8ca97e2 / 3

- [vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-8d347a2fc91fcb48739a5ebfe8f52bc7c0ddc7895c4a9ea288c30a23847f3ff0): complete subsection reference.

<a id="canonical-bdfb7c123239bf68399779f356c42e520a156de2dff8fe4f92a0d8524d49f6ca"></a>

## Next pages — azure_vnet_site.vnet_attachments / e85cc8ca97e2 / 4

- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-8d347a2fc91fcb48739a5ebfe8f52bc7c0ddc7895c4a9ea288c30a23847f3ff0)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-8d347a2fc91fcb48739a5ebfe8f52bc7c0ddc7895c4a9ea288c30a23847f3ff0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5e817dc5bfd1898f195442b231dc73fca182a71cfb825482e24ae21d7cdffb6"></a>

## azure_vnet_site.vnet_attachments.vnet_list — azure_vnet_site.vnet_attachments.vnet_list / 13131c3c45c9 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-2ce367bba074251c2ff0ecdf3d3aa3978f5eb4f4b487bce51290abdaa48643ad)
- azure_vnet_site.vnet_attachments.vnet_list

<a id="canonical-c3d0204bcd0563cecbaaed988429706f09f4e4c342f024e8f576fffdd158df46"></a>

Type: `"list"`. Computed.

VNet List. Collection of items or values

Upstream description:

Collection of items or values

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

<a id="canonical-1d4df7c237765316c1330a5d45c58916f0b8cf88959e762c06c6a32308a7a5d2"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list / 13131c3c45c9 / 3

- [custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-cb0055cde8c13a5e8debb374a7c055399c70c0ba7f6675e8fb41b9a060807d18): complete subsection reference.

- [default_route](data-sources--cloud_connect--reference--group-001.md#canonical-926d59300662daabc2e242d7c9aba0360182f1dd07773728cd505a1b5bb81e44): complete subsection reference.

- [labels](data-sources--cloud_connect--reference--group-001.md#canonical-2cd50cd9925df9f0b6f943acb48a8e786f6f120f91dc7e383eb6cfb1dfcabd91): complete subsection reference.

- [manual_routing](data-sources--cloud_connect--reference--group-001.md#canonical-e305345190315320d122514809cb91849a985fc8c436d85d2fc7eda13f724ec6): complete subsection reference.

<a id="canonical-28e78325cb8b13ad522ab6ccf47907940917c409c7748568eb49fede97369a34"></a>

<a id="canonical-fd569bb05e237bf71cd98ada5a8ae892e007b1c3b4a556a5c91de2c34fb5c55f"></a>

## subscription_id property — azure_vnet_site.vnet_attachments.vnet_list / 13131c3c45c9 / 4

Type: `"string"`. Computed.

Enter the Subscription ID of the VNet to be attached.

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

<a id="canonical-3bc943bfc1eab2f53efc99ac6626c0256ad00feb32e13afe28cb9446e8db57b5"></a>

<a id="canonical-20547d6220e87fb06bd432a50c6ffabf45e2fe32387160354d7623bff88dc911"></a>

## vnet_id property — azure_vnet_site.vnet_attachments.vnet_list / 13131c3c45c9 / 5

Type: `"string"`. Computed.

Enter the VNet ID of the VNet to be attached in format
/&lt;resource-group-name&gt;/&lt;VNet-name&gt;.

Upstream description:

Enter the VNet ID of the VNet to be attached in format
/&lt;resource-group-name&gt;/&lt;VNet-name&gt;

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

<a id="canonical-f6e386934980ce4314c53fc014d7ed049863a1419903ec8b7b1f46cbacd0290d"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list / 13131c3c45c9 / 6

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-cb0055cde8c13a5e8debb374a7c055399c70c0ba7f6675e8fb41b9a060807d18)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-926d59300662daabc2e242d7c9aba0360182f1dd07773728cd505a1b5bb81e44)
- [azure_vnet_site.vnet_attachments.vnet_list.labels](data-sources--cloud_connect--reference--group-001.md#canonical-2cd50cd9925df9f0b6f943acb48a8e786f6f120f91dc7e383eb6cfb1dfcabd91)
- [azure_vnet_site.vnet_attachments.vnet_list.manual_routing](data-sources--cloud_connect--reference--group-001.md#canonical-e305345190315320d122514809cb91849a985fc8c436d85d2fc7eda13f724ec6)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-2ce367bba074251c2ff0ecdf3d3aa3978f5eb4f4b487bce51290abdaa48643ad)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-cb0055cde8c13a5e8debb374a7c055399c70c0ba7f6675e8fb41b9a060807d18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbae4dfa3352c253e24803249264a50d4f8552118bfe6b712fe27f99232f4337"></a>

## azure_vnet_site.vnet_attachments.vnet_list.custom_routing — azure_vnet_site.vnet_attachments.vnet_list.custom_routing / c3f9919ad7eb / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-2ce367bba074251c2ff0ecdf3d3aa3978f5eb4f4b487bce51290abdaa48643ad)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-8d347a2fc91fcb48739a5ebfe8f52bc7c0ddc7895c4a9ea288c30a23847f3ff0)
- azure_vnet_site.vnet_attachments.vnet_list.custom_routing

<a id="canonical-62a7e2e6095440e23e65d5919f6676ae88194cdef3811402b0f115f525c339ba"></a>

Type: `"single"`. Computed.

List Azure Route Table with Static Route.

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

<a id="canonical-f238df7be4dbd3aad2e0e862b029f5c94ae642addb716c19ca2c83caeb80bdc3"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list.custom_routing / c3f9919ad7eb / 3

- [route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-ae808327af691e3708913f361d766022b7e8e75e112755843dee2f0ccb888fbc): complete subsection reference.

<a id="canonical-bf2f1836866a355fc37c0719b7830e0b09fb58baa6761f614d2ba3ee95b292b1"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list.custom_routing / c3f9919ad7eb / 4

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-ae808327af691e3708913f361d766022b7e8e75e112755843dee2f0ccb888fbc)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-8d347a2fc91fcb48739a5ebfe8f52bc7c0ddc7895c4a9ea288c30a23847f3ff0)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-ae808327af691e3708913f361d766022b7e8e75e112755843dee2f0ccb888fbc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-caeb23e0aa49db0978c4efb2453320245c578463bbe882fbc1324fc7aa97a981"></a>

## azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables — azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables / 9138ba91865f / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-2ce367bba074251c2ff0ecdf3d3aa3978f5eb4f4b487bce51290abdaa48643ad)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-8d347a2fc91fcb48739a5ebfe8f52bc7c0ddc7895c4a9ea288c30a23847f3ff0)
- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-cb0055cde8c13a5e8debb374a7c055399c70c0ba7f6675e8fb41b9a060807d18)
- azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables

<a id="canonical-3a89b8a3f8e9ad5ac574b043856849517107b035f0e03d831a6b2afd1ee38807"></a>

Type: `"list"`. Computed.

List of route tables with static routes. Route Tables with static routes.

Upstream description:

Route Tables with static routes.

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

<a id="canonical-998dc442db95dbd28aecc27af45f5c11f4134a33ce5ae208015ca34827bc8393"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables / 9138ba91865f / 3

<a id="canonical-eb21414f1fb83b6c0dd5e8073b00d2b8be1ac7e9968cdc51e11d882184c65388"></a>

<a id="canonical-c4134ab0d8f586d145ebf482d0e9672fe930eb31f2890bc6b12934fd52d4fda3"></a>

## route_table_id property — azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables / 9138ba91865f / 4

Type: `"string"`. Computed.

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;.

Upstream description:

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;

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

<a id="canonical-8d60dd9e0fba3ccf731fdea572c819c7647149a1c5baf5bba1109f9126b58fa9"></a>

<a id="canonical-81a7daf1062b8021f48ffa1fa8345321767a316fd6ee1913e4df7f465a58f902"></a>

## static_routes property — azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables / 9138ba91865f / 5

Type: `["list", "string"]`. Computed.

Static Routes. List of Static Routes.

Upstream description:

List of Static Routes.

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

<a id="canonical-2308742426cf9aa21b43a72138be45a15701536c040a8921e7f09a40c47c4efa"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables / 9138ba91865f / 6

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-cb0055cde8c13a5e8debb374a7c055399c70c0ba7f6675e8fb41b9a060807d18)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-926d59300662daabc2e242d7c9aba0360182f1dd07773728cd505a1b5bb81e44"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb18433ecb47a2fb76d5835a1b6de84a6c5810a887e32fbd885c405232df625d"></a>

## azure_vnet_site.vnet_attachments.vnet_list.default_route — azure_vnet_site.vnet_attachments.vnet_list.default_route / 37295d383639 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-2ce367bba074251c2ff0ecdf3d3aa3978f5eb4f4b487bce51290abdaa48643ad)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-8d347a2fc91fcb48739a5ebfe8f52bc7c0ddc7895c4a9ea288c30a23847f3ff0)
- azure_vnet_site.vnet_attachments.vnet_list.default_route

<a id="canonical-aca8ec81e2436c80ccbbebdf326acf5fb52d4ca7b23568235a231751b840dd16"></a>

Type: `"single"`. Computed.

Configuration parameter for default route.

Upstream description:

Select Override Default Route Choice.

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

<a id="canonical-9cd52ec19f604ca30ec8ce8a5e3fcac5d9e5a1d9924ff63a3e16816de6fb4b6f"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list.default_route / 37295d383639 / 3

- [all_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-677d6f29310423704569711663a41e2ba20474ac174524b1cbdd9109d23246a5): complete subsection reference.

- [selective_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-08048a63fb354552eb798b97cdb3f0438aa7852728f5061bfe532b427aff59bd): complete subsection reference.

<a id="canonical-10f3e33ef32ab1a1a5702e571f60241ec8ef4e11743ba6d07e2805c6d87f6ac5"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list.default_route / 37295d383639 / 4

- [azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-677d6f29310423704569711663a41e2ba20474ac174524b1cbdd9109d23246a5)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-08048a63fb354552eb798b97cdb3f0438aa7852728f5061bfe532b427aff59bd)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-8d347a2fc91fcb48739a5ebfe8f52bc7c0ddc7895c4a9ea288c30a23847f3ff0)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-677d6f29310423704569711663a41e2ba20474ac174524b1cbdd9109d23246a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e99a7829e491b248ab01b60f41dfbeaab8bb99db6c9ad25d57c50712f606ef4b"></a>

## azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables — azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables / 47d0e1e52c5c / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-2ce367bba074251c2ff0ecdf3d3aa3978f5eb4f4b487bce51290abdaa48643ad)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-8d347a2fc91fcb48739a5ebfe8f52bc7c0ddc7895c4a9ea288c30a23847f3ff0)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-926d59300662daabc2e242d7c9aba0360182f1dd07773728cd505a1b5bb81e44)
- azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables

<a id="canonical-2d6d6520c805cd301462417bc0d87e3e26ac62f3cfa3de44410adb3e346acc29"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-7a56f510b11750a349ec38608c79201b1c287eef347da862185748c6306fc132"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables / 47d0e1e52c5c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ba11ab8ebc7ccd0228fab81a5294f281f3bed4928c3ea9ea4da6cd5437fe94d6"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables / 47d0e1e52c5c / 4

- [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-926d59300662daabc2e242d7c9aba0360182f1dd07773728cd505a1b5bb81e44)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-08048a63fb354552eb798b97cdb3f0438aa7852728f5061bfe532b427aff59bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41a0ef534997d9e0b2c65d1c89ada9fdc27e7469eb358e428759ea9a4e882120"></a>

## azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables — azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables / e3a92f68268c / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-2ce367bba074251c2ff0ecdf3d3aa3978f5eb4f4b487bce51290abdaa48643ad)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-8d347a2fc91fcb48739a5ebfe8f52bc7c0ddc7895c4a9ea288c30a23847f3ff0)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-926d59300662daabc2e242d7c9aba0360182f1dd07773728cd505a1b5bb81e44)
- azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables

<a id="canonical-675830ca175ff86cf7394cd88752bb6fe4239a070346adccf668aeb27a1b4a68"></a>

Type: `"single"`. Computed.

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

<a id="canonical-880dc1fa97c3a7737fd21534e7679786bad38fcdfaeb95d6ddefe8041c5e30db"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables / e3a92f68268c / 3

<a id="canonical-7f5576d407b8502fa506928e8d4de11dd6834657a1e616adc84a7d75cdfaada7"></a>

<a id="canonical-c2a5ffd594d490bc8ecadd14ed4d89e43b62bd807aa74c8eb2b82bb57cb93f33"></a>

## route_table_id property — azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables / e3a92f68268c / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-94061f14123392fdc70c5159b28817b8c7e7caa012ea6d652c6bc193d48357bf"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables / e3a92f68268c / 5

- [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-926d59300662daabc2e242d7c9aba0360182f1dd07773728cd505a1b5bb81e44)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-2cd50cd9925df9f0b6f943acb48a8e786f6f120f91dc7e383eb6cfb1dfcabd91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49731db334523f6e85a264be710543067b8780754efa17602eb8b368ef1734b1"></a>

## azure_vnet_site.vnet_attachments.vnet_list.labels — azure_vnet_site.vnet_attachments.vnet_list.labels / ab7b5c9c3dd8 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-2ce367bba074251c2ff0ecdf3d3aa3978f5eb4f4b487bce51290abdaa48643ad)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-8d347a2fc91fcb48739a5ebfe8f52bc7c0ddc7895c4a9ea288c30a23847f3ff0)
- azure_vnet_site.vnet_attachments.vnet_list.labels

<a id="canonical-4b993aee1d641aa6362f225631781a4e24d2f93ee29317fe0a254173507dcd45"></a>

Type: `"single"`. Computed.

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

<a id="canonical-a20bb233b4ea92f80b0829715218abd6d0a6b8302423c497f1f65bd1fb02d44a"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list.labels / ab7b5c9c3dd8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e0c7a1aa47bb89d6aa7a52309af889d7c1a249eae545ab778cac3cbf9bdb45e3"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list.labels / ab7b5c9c3dd8 / 4

- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-8d347a2fc91fcb48739a5ebfe8f52bc7c0ddc7895c4a9ea288c30a23847f3ff0)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-e305345190315320d122514809cb91849a985fc8c436d85d2fc7eda13f724ec6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebb5f66ed36cb7dae07a5f29a1f877c080f6ba816b69f6d958d2fe1f0b648eae"></a>

## azure_vnet_site.vnet_attachments.vnet_list.manual_routing — azure_vnet_site.vnet_attachments.vnet_list.manual_routing / d8e265a285d0 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-749ec122046b0ba8cc0d1abdcf9ae27321bba8daf365b5acefaa3fbdaa92375e)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-2ce367bba074251c2ff0ecdf3d3aa3978f5eb4f4b487bce51290abdaa48643ad)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-8d347a2fc91fcb48739a5ebfe8f52bc7c0ddc7895c4a9ea288c30a23847f3ff0)
- azure_vnet_site.vnet_attachments.vnet_list.manual_routing

<a id="canonical-745cad365c4d1636337902232751cd481b6b842cccf512b9c835742856bf820d"></a>

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

<a id="canonical-08695f5202003e4cfd0d9a94d00642a749959132f0f23dc06dbc56aeaa5eb37b"></a>

## Direct properties — azure_vnet_site.vnet_attachments.vnet_list.manual_routing / d8e265a285d0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3cc2e9ef7df54812c36cad855129bfa1254c47dec490736f89755b527f80c922"></a>

## Next pages — azure_vnet_site.vnet_attachments.vnet_list.manual_routing / d8e265a285d0 / 4

- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-8d347a2fc91fcb48739a5ebfe8f52bc7c0ddc7895c4a9ea288c30a23847f3ff0)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-2dd4db43791f10dcf78cba7f07887e11b8819442a3176afb6303be1a60516481"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-252768f9f679f6c79665e8180bc00fc4f2add6fd778c90932636c7edb280a86a"></a>

## segment — segment / f13b520684a9 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- segment

<a id="canonical-79960451edbffdb34bfc7f74b628cc679a19347416f3d33ca4f17f5328b16e46"></a>

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

<a id="canonical-c2e89b5ba65f62948d35c1276c25938aaf19d94e01d2c5d238a408ab8a0ef89f"></a>

## Direct properties — segment / f13b520684a9 / 3

<a id="canonical-584cf8a5766cce73995acd0137492ec82d3a37e86b9653d25292acde6ba85453"></a>

<a id="canonical-17db44291058b7ab689ebe0a566d971bfcdd6e64f40e4cc4ba393692445bcd1c"></a>

## name property — segment / f13b520684a9 / 4

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

<a id="canonical-9d641c9fa4bb4e89a670bf6e849e20164b23b4c2e0dc78a8e299fab2a0e0d127"></a>

<a id="canonical-ea527e5f0ec110bfdec19da6b16170d10a8b62e6863291b9d046c28535e025bc"></a>

## namespace property — segment / f13b520684a9 / 5

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

<a id="canonical-c699a8b1ff0284a2e110ead01db43a09edc324b35fbc39ec3342ebf7fe5128ea"></a>

<a id="canonical-126d8518896b0523d0f52da65e6ef3074cff1e2124ea1f482eee6c5404abf72b"></a>

## tenant property — segment / f13b520684a9 / 6

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

<a id="canonical-a41d251c71d7557bd4f3a83c00aeba70a70d568018eb478293b828f12008aefd"></a>

## Next pages — segment / f13b520684a9 / 7

- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-2a3a55770310a476342fc6bf216a73b51a9d15218221552469e3779bf0dfeb60)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
