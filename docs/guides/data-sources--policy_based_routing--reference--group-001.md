---
page_title: "xcsh_policy_based_routing reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policy_based_routing reference."
---

# xcsh_policy_based_routing reference

<a id="canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc230f7ed642f7c5806b56b5014eaea57349349ae2fae10ef0a9e0d81ac064c9"></a>

## Property reference — Property reference / da3202ed899c / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- Property reference

<a id="canonical-cdcaeb666c30b807509a5e984dcdc0a6d6b0844b38d595f133adcb199d42e539"></a>

## Direct properties — Property reference / da3202ed899c / 3

<a id="canonical-d65b53af970f8c217c33a56e0481c09de7370d4a30f8471ddd2100ed1a59155e"></a>

<a id="canonical-7cf56caef386e1c6990684f92ebd212c923b0577160cd30f679bd49dbb16111d"></a>

## annotations property — Property reference / da3202ed899c / 4

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

<a id="canonical-8dde88b0a9da616c0206253a70d97af8a43b845feebf2cd34f8d62d84122242d"></a>

<a id="canonical-414f9814d00cc2666f0ac28148418a5adf59473412085f64e33b095153a78d87"></a>

## description property — Property reference / da3202ed899c / 5

Type: `"string"`. Computed.

Description of the PolicyBasedRouting.

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

- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32): complete subsection reference.

- [forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-00df9106da3d9248c38a0af546df4f2548c586620425b7859c7f7583a1e0012e): complete subsection reference.

<a id="canonical-d6980d2f8aea28788e5d05a6b7bcfc7ef0c155a8570292b233a41cfb1101d8c6"></a>

<a id="canonical-e2abe406b121e43f6b16ed7a9a551e7e938e6fa57ed134ac69458d7e614ec590"></a>

## id property — Property reference / da3202ed899c / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3cdfc38d17cc8f947439afe0d4285dc9e5268c04a6b3b37b39295c29e6176b03"></a>

<a id="canonical-5dc4c434e1cdfbcb4cb211beb1e2519c385fbcf41197b4412cfb3fe1756fe716"></a>

## labels property — Property reference / da3202ed899c / 7

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

<a id="canonical-ad3691de90caaa011cd3f7e9500ffca2fc9872df918bdf5796f81766607f11a1"></a>

<a id="canonical-7b4dfb484f400e704eea350d5d4c69868388343cffd4f42c5459555a2e4d3c6b"></a>

## name property — Property reference / da3202ed899c / 8

Type: `"string"`. Required.

Name of the PolicyBasedRouting.

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

<a id="canonical-2f4cae6f96fc7dd77e190279607cf18156d51fe60622a13252a50175715809e1"></a>

<a id="canonical-091567c2b1592326f1a6770cee98addf21aa74600ed8edc06d0c480a23fed675"></a>

## namespace property — Property reference / da3202ed899c / 9

Type: `"string"`. Required.

Namespace where the PolicyBasedRouting exists.

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

- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405): complete subsection reference.

<a id="canonical-3add7c8c07b871da4ea0e99bf680a8993a5d2a8ad77f63a54868abf84b28c50c"></a>

## All schema paths — Property reference / da3202ed899c / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--policy_based_routing--reference--group-001.md#canonical-d65b53af970f8c217c33a56e0481c09de7370d4a30f8471ddd2100ed1a59155e) |
| `description` | [description](data-sources--policy_based_routing--reference--group-001.md#canonical-8dde88b0a9da616c0206253a70d97af8a43b845feebf2cd34f8d62d84122242d) |
| `forward_proxy_pbr` | [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-e9a8854ee1eccb87012ba27ca87c5beb0c5249c29d6f9d8683953c210ace7c21) |
| `forward_proxy_pbr.forward_proxy_pbr_rules` | [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-f9016ce0654089533256e7f737db396dd223d446640825586a332b7a61fedb68) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations` | [forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations](data-sources--policy_based_routing--reference--group-001.md#canonical-0cdaaa07c84eafc0f2c21222cc637de6c4cb8edf884b592aa957efb49f4821a7) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.all_sources` | [forward_proxy_pbr.forward_proxy_pbr_rules.all_sources](data-sources--policy_based_routing--reference--group-001.md#canonical-ea01b709ca013ea77fbc310635f4a77f49b8e41a4387fba09e677e77fc2e77cf) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-13555c9a39125c7ed0e35d5f8a9ce5b2d8adccf1b35526c2c4c263c07fbfba6b) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.name](data-sources--policy_based_routing--reference--group-001.md#canonical-d7be537368b1a889b6a8e9899563f153e524d7086928ff542003c662a35e5f88) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.namespace` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.namespace](data-sources--policy_based_routing--reference--group-001.md#canonical-36fa836f74f1461e6b0d370d69875c2a33f9e6a8997dddee201abb973a64b5d4) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.tenant` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.tenant](data-sources--policy_based_routing--reference--group-001.md#canonical-3b87e2bf6f6a47119d3ef982b8b9ecdbfb6bd1b333ed9705b02f96e568029e64) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-184b4eaeb8a92e05417c01033d10788c5cff91a80c4e6574e63759161b7c05f0) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-8659d1f49311c1f9a8b12820d0e1d1db3f7b29bab29888ddff6ccab8a5f5da98) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path](data-sources--policy_based_routing--reference--group-001.md#canonical-8fbff81e669f86b440c328e6552b59e618d8b6ec26da346907945dda825a430e) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.exact_value](data-sources--policy_based_routing--reference--group-001.md#canonical-9cb2f451b6294f8e944cd6ddfa9df531704f19192639a936985ae4418b5d3c55) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_exact_value](data-sources--policy_based_routing--reference--group-001.md#canonical-e24e7e770e9136625787e96ee4a548b441c9552398cd08f538958b9f0e741301) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_prefix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_prefix_value](data-sources--policy_based_routing--reference--group-001.md#canonical-7eb89cb76d1de2c06d9fbafd839a7e9f65408a2607f8dfeea371915421276e9a) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_regex_value](data-sources--policy_based_routing--reference--group-001.md#canonical-5d52dd0ed76ee492b769ddcbff0d4fd8e3465d476ee702df4983630ccc60eaf7) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.regex_value](data-sources--policy_based_routing--reference--group-001.md#canonical-236ecd62dbfccdb90e3a043a6ad80569dccfbf5f040d7ceb85175ba8d1ccef9c) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.suffix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.suffix_value](data-sources--policy_based_routing--reference--group-001.md#canonical-8d9496d9fbbca5e79d0e3304a11db00994e5247d7faa0e179057902b33510411) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set](data-sources--policy_based_routing--reference--group-001.md#canonical-fc7528b27302c2fb29eddeb742c80ed04bbf13bd55aa74cb49484ae73ed16811) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.name](data-sources--policy_based_routing--reference--group-001.md#canonical-d18e11bd2fcf2817d90fa94c21cb050a6b2b6c9ab9993490c0924c1182da02f9) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.namespace` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.namespace](data-sources--policy_based_routing--reference--group-001.md#canonical-89a6fc8f173287da80252d10383923768dcc53a9a02dad3bcbae5d5dfe3e15b3) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.tenant` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.tenant](data-sources--policy_based_routing--reference--group-001.md#canonical-c98bca840512e209de15817a8c540042ad433ad7f2f46269621bc5affb29d5df) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector` | [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector](data-sources--policy_based_routing--reference--group-001.md#canonical-e2c4812f61066bbd4c9c5e6aaf1741a4d9ad070593701e9091b82c6fbd37b8f5) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector.expressions` | [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector.expressions](data-sources--policy_based_routing--reference--group-001.md#canonical-7a533a8bd5e7f2ffd67e915ae4e2412b0ef84e839ae0d820133ba5b152d7875b) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata](data-sources--policy_based_routing--reference--group-001.md#canonical-882e22f3385e35b608b16777f0a27b04c6060178c8f172ece34b4e781cd44713) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.description_spec` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata.description_spec](data-sources--policy_based_routing--reference--group-001.md#canonical-e4741f57f567f3ddeedc8bdbbf0820147d8a3d48ecc4e7180289813e994bc816) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata.name](data-sources--policy_based_routing--reference--group-001.md#canonical-05ee124fd531cbd928a40066a14e69da1e4dbd7a252b50ce684f07b7a0ee1e2c) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-5ec6f762eb6a3e37a2435c87fcfb3331bc6f28d8a4e274fde3c1c03c340aaa8f) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list.prefixes` | [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list.prefixes](data-sources--policy_based_routing--reference--group-001.md#canonical-3a0e299ee52170cfa2f23a063271cfc3558d3e7816f869811164e9fdc5239775) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](data-sources--policy_based_routing--reference--group-001.md#canonical-fc68e11402e79d093af1147ad8497c182fb8a24fee4c7160344dd7baaf2d2d58) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list](data-sources--policy_based_routing--reference--group-001.md#canonical-cd4bf44e645ca693aa812dde8756fd4889e9e0577d7e0ecb2c9ed8be42e24c98) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.exact_value](data-sources--policy_based_routing--reference--group-001.md#canonical-2c571f151e604358c5aa36f6aa28e95d1fd6894342d6fd145b5550cb0d80fe3b) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.regex_value](data-sources--policy_based_routing--reference--group-001.md#canonical-b3e8a03fe929d48f4995e56e00e25f0ab526e63ccce75d976c552fd4cbf7ff98) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.suffix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.suffix_value](data-sources--policy_based_routing--reference--group-001.md#canonical-27a45b5b54b4d4ec3c584b3a0494d84f60d759a7531cd1d2c873e5e622cbb8f6) |
| `forwarding_class_list` | [forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-7125047be099209e535432dfe5dec37b4afb8ccbfc2bcdce248dacb47bc8e5d8) |
| `forwarding_class_list.name` | [forwarding_class_list.name](data-sources--policy_based_routing--reference--group-001.md#canonical-c737c0f7fbedec79778fe952772559e4bf85c6fc9803792b598a3eb7f631b8d2) |
| `forwarding_class_list.namespace` | [forwarding_class_list.namespace](data-sources--policy_based_routing--reference--group-001.md#canonical-cffbb1f52c1803c9421659478c55c46be4a4355915f454172671f9fedc0f7e06) |
| `forwarding_class_list.tenant` | [forwarding_class_list.tenant](data-sources--policy_based_routing--reference--group-001.md#canonical-67bd05388ffc3a453a391818fa0bffc034f96f2d40cb43d52184df8de6587be8) |
| `id` | [id](data-sources--policy_based_routing--reference--group-001.md#canonical-d6980d2f8aea28788e5d05a6b7bcfc7ef0c155a8570292b233a41cfb1101d8c6) |
| `labels` | [labels](data-sources--policy_based_routing--reference--group-001.md#canonical-3cdfc38d17cc8f947439afe0d4285dc9e5268c04a6b3b37b39295c29e6176b03) |
| `name` | [name](data-sources--policy_based_routing--reference--group-001.md#canonical-ad3691de90caaa011cd3f7e9500ffca2fc9872df918bdf5796f81766607f11a1) |
| `namespace` | [namespace](data-sources--policy_based_routing--reference--group-001.md#canonical-2f4cae6f96fc7dd77e190279607cf18156d51fe60622a13252a50175715809e1) |
| `network_pbr` | [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-fd04ccc60307e88d9463a971f5d7df1a85169d821589db16d46bf2b9b8dbba03) |
| `network_pbr.any` | [network_pbr.any](data-sources--policy_based_routing--reference--group-001.md#canonical-430b041e36eaeb3ab1e087d17a9c353d904864303e2afe1f5f4dbe02a9e4452d) |
| `network_pbr.label_selector` | [network_pbr.label_selector](data-sources--policy_based_routing--reference--group-001.md#canonical-156a1784f386717797403c94f7be1c5b90253e7095c40f7fad8035dbabec8473) |
| `network_pbr.label_selector.expressions` | [network_pbr.label_selector.expressions](data-sources--policy_based_routing--reference--group-001.md#canonical-31990e30529bdbce62c36fa3106f58166e3a99792d01efd5986dba9acb24ceef) |
| `network_pbr.network_pbr_rules` | [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-a3d30a9b1b0ab258d82e4ac2ef12849a2152deb3c09bf35ea25ea0151c2c4d74) |
| `network_pbr.network_pbr_rules.all_tcp_traffic` | [network_pbr.network_pbr_rules.all_tcp_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-82855584f0cf011acfa049273ba57e93c8334a70341872a0383d632be429509e) |
| `network_pbr.network_pbr_rules.all_traffic` | [network_pbr.network_pbr_rules.all_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-4eb2dac0ae9ac5a10725c53ca4dfb394e1d54c3613f74de35afe2ec30b0ba99f) |
| `network_pbr.network_pbr_rules.all_udp_traffic` | [network_pbr.network_pbr_rules.all_udp_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-35fb8193aa430a785a9378f1c35c49272e0d67c17723ab2999760eb7eb26a093) |
| `network_pbr.network_pbr_rules.any` | [network_pbr.network_pbr_rules.any](data-sources--policy_based_routing--reference--group-001.md#canonical-bf893380bc5ec146d7e9b665c23a192c82ef962356091e56dc19dab0df37bccb) |
| `network_pbr.network_pbr_rules.applications` | [network_pbr.network_pbr_rules.applications](data-sources--policy_based_routing--reference--group-001.md#canonical-175a063f9e522387be6b671f3410684bdba1d2906e58c238829e33e1f2e89fdc) |
| `network_pbr.network_pbr_rules.applications.applications` | [network_pbr.network_pbr_rules.applications.applications](data-sources--policy_based_routing--reference--group-001.md#canonical-cd6467c9ca6ee36d1b64de466d1dd3c422703d21ccb851d3876dc170c210da1d) |
| `network_pbr.network_pbr_rules.dns_name` | [network_pbr.network_pbr_rules.dns_name](data-sources--policy_based_routing--reference--group-001.md#canonical-a1ad512aaf920dc41ad9f616f9fe36b8ecf0bafddb5dd3d12c44c583351598c4) |
| `network_pbr.network_pbr_rules.forwarding_class_list` | [network_pbr.network_pbr_rules.forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-26e0ebee13d9739fec706f26b4dd04381ccfb9d3a0f426e5c046cc9f12ff491b) |
| `network_pbr.network_pbr_rules.forwarding_class_list.name` | [network_pbr.network_pbr_rules.forwarding_class_list.name](data-sources--policy_based_routing--reference--group-001.md#canonical-42d260713981fa7aadb4b269580bbfc89b8142d013eeb45164cac65d1cd8d00c) |
| `network_pbr.network_pbr_rules.forwarding_class_list.namespace` | [network_pbr.network_pbr_rules.forwarding_class_list.namespace](data-sources--policy_based_routing--reference--group-001.md#canonical-61bccef6ce0a17fca8cb1e47c8bbdcce032484537630751688c426a20b81e136) |
| `network_pbr.network_pbr_rules.forwarding_class_list.tenant` | [network_pbr.network_pbr_rules.forwarding_class_list.tenant](data-sources--policy_based_routing--reference--group-001.md#canonical-f1ac34138432d0b344a155f0ebcbfc6b4d8ecc95e50d6d7e4659ca0b0b7422b3) |
| `network_pbr.network_pbr_rules.ip_prefix_set` | [network_pbr.network_pbr_rules.ip_prefix_set](data-sources--policy_based_routing--reference--group-001.md#canonical-7ad889af4e73d0ff297bce16035ff989905fcdc6f0ff6fd6e6b2d4041c78d103) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref` | [network_pbr.network_pbr_rules.ip_prefix_set.ref](data-sources--policy_based_routing--reference--group-001.md#canonical-95476b283c9fb2fcbdec29487f9741dfe7ad68e4bc46e11ab44874751efa3dd5) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.kind` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.kind](data-sources--policy_based_routing--reference--group-001.md#canonical-912cd407119ba82eb97e5717247570712d2603c07071f191cd0901283a0c99ea) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.name` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.name](data-sources--policy_based_routing--reference--group-001.md#canonical-9e9edcb5052edd851b1f274af730c02ba94b65c0b5546225e840cd49665e5ffc) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.namespace` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.namespace](data-sources--policy_based_routing--reference--group-001.md#canonical-8926721449c19cfe9df40209a7d0e73239d6f2d24ca416e9a76e43d3c50ea574) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.tenant` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.tenant](data-sources--policy_based_routing--reference--group-001.md#canonical-660cfb277b528b749565200b706e67ec3d6d2f54e6977126d13e1b82e984eb6b) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.uid` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.uid](data-sources--policy_based_routing--reference--group-001.md#canonical-5faa6725e40df0284fc2a52f3309c0d04e6d1e491040e98b56cd088d245b0ff3) |
| `network_pbr.network_pbr_rules.metadata` | [network_pbr.network_pbr_rules.metadata](data-sources--policy_based_routing--reference--group-001.md#canonical-37f9c9c5b5c5dc6021b9352610bcdfa2deec1f967bebfc1903ccde62aa3fbfdb) |
| `network_pbr.network_pbr_rules.metadata.description_spec` | [network_pbr.network_pbr_rules.metadata.description_spec](data-sources--policy_based_routing--reference--group-001.md#canonical-aa28934e23b48f10d21805b0354d8e16d32e9f8d2ab160e3ff247f2bcc4bfcc3) |
| `network_pbr.network_pbr_rules.metadata.name` | [network_pbr.network_pbr_rules.metadata.name](data-sources--policy_based_routing--reference--group-001.md#canonical-8054927c92edb8ea2149f756ed06b3a9846c462dac840500210cea07f566f18d) |
| `network_pbr.network_pbr_rules.prefix_list` | [network_pbr.network_pbr_rules.prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-4083e0eb6e062fde7e8d9432e94e582242a49994320574de554ebec9ce30691c) |
| `network_pbr.network_pbr_rules.prefix_list.prefixes` | [network_pbr.network_pbr_rules.prefix_list.prefixes](data-sources--policy_based_routing--reference--group-001.md#canonical-57909b01ac7931c9a3782ff5b51a11de8db74f621af130ac6a6d9da8a9a4ee06) |
| `network_pbr.network_pbr_rules.protocol_port_range` | [network_pbr.network_pbr_rules.protocol_port_range](data-sources--policy_based_routing--reference--group-001.md#canonical-33208f1a676cbe633844719daa4b20691e12dfe2ff0f1973bd16d9910b7e5dd4) |
| `network_pbr.network_pbr_rules.protocol_port_range.port_ranges` | [network_pbr.network_pbr_rules.protocol_port_range.port_ranges](data-sources--policy_based_routing--reference--group-001.md#canonical-b8d50dbd76780d32b11d12e6224db04a72c847b0ba838256b93737cc5c8906e7) |
| `network_pbr.network_pbr_rules.protocol_port_range.protocol` | [network_pbr.network_pbr_rules.protocol_port_range.protocol](data-sources--policy_based_routing--reference--group-001.md#canonical-c06e8b388a0ac5f95f0d895efc1893a61ccad968cc09777eb3a1db3ea2048791) |
| `network_pbr.prefix_list` | [network_pbr.prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-8d635098eb8ea341d7ba1def8781b9982c5a3db07ff148726ea04de622789628) |
| `network_pbr.prefix_list.prefixes` | [network_pbr.prefix_list.prefixes](data-sources--policy_based_routing--reference--group-001.md#canonical-9f2909e540d7324505f3486fe306bc206a7fd50359896e3009632097108ae4e2) |

<a id="canonical-cbb83aa52681f73635634dc4ae079f35c622aae0f0bf8637ed91dd6af214a6e6"></a>

## Next pages — Property reference / da3202ed899c / 11

- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- [forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-00df9106da3d9248c38a0af546df4f2548c586620425b7859c7f7583a1e0012e)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30ff7a949eff2694468db15723ee4c863e92cf01aad075ea5cefd356714e23d2"></a>

## forward_proxy_pbr — forward_proxy_pbr / ba4362200373 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- forward_proxy_pbr

<a id="canonical-e9a8854ee1eccb87012ba27ca87c5beb0c5249c29d6f9d8683953c210ace7c21"></a>

Type: `"single"`. Computed.

\[OneOf: forward\_proxy\_pbr, network\_pbr\] Configuration parameter for forward proxy pbr.

Upstream description:

Network(L3/L4) routing policy rule.

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

- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-e9a8854ee1eccb87012ba27ca87c5beb0c5249c29d6f9d8683953c210ace7c21)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-fd04ccc60307e88d9463a971f5d7df1a85169d821589db16d46bf2b9b8dbba03)

Select alternatives according to the provider validators above.

<a id="canonical-88c0413a5b4771aa610c68de071ba0bc25b292677f86cf3f8ca8a3443895eba2"></a>

## Direct properties — forward_proxy_pbr / ba4362200373 / 3

- [forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a): complete subsection reference.

<a id="canonical-176019229add2b44acedde76220d3cbc54789bf228d9c922f7445e69ab1d2d6b"></a>

## Next pages — forward_proxy_pbr / ba4362200373 / 4

- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c26d53bcabb9314ef836eb3f01a143859526425a3e0e223088144facdb32edc"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules — forward_proxy_pbr.forward_proxy_pbr_rules / c53be9d9ee46 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- forward_proxy_pbr.forward_proxy_pbr_rules

<a id="canonical-f9016ce0654089533256e7f737db396dd223d446640825586a332b7a61fedb68"></a>

Type: `"list"`. Computed.

L3/L4 routing rules. Network(L3/L4) routing policy rules.

Upstream description:

Network(L3/L4) routing policy rules.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-df034e255216c71f415f5a0c828403c54e3fe78e37c79e855b95a637508a6e8a"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules / c53be9d9ee46 / 3

- [all_destinations](data-sources--policy_based_routing--reference--group-001.md#canonical-6171bb1ebdf67b5399f835c91a31adf5a2bba15741984571d15da638a042a627): complete subsection reference.

- [all_sources](data-sources--policy_based_routing--reference--group-001.md#canonical-e083eb0f522a1c30f3393d80342938858bbffd2ad38646feac3f64736d7b1c6a): complete subsection reference.

- [forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-72201c762aacf19d586743867074af145919058d40a70adaa7c3a9494a09a6ea): complete subsection reference.

- [http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-7bbfc8405b63cdcbae107b961f216c9b3f43cb2a154cd2b7b4c24830032aab47): complete subsection reference.

- [ip_prefix_set](data-sources--policy_based_routing--reference--group-001.md#canonical-2c851c7847ce6d67448b4ca967aeee60e4fa9948ed8298d52db1b85dd410d4a9): complete subsection reference.

- [label_selector](data-sources--policy_based_routing--reference--group-001.md#canonical-f32637a485d2e418bd8573fd08d7d6116a869e8c9eb2d15ad43792d4a8035c36): complete subsection reference.

- [metadata](data-sources--policy_based_routing--reference--group-001.md#canonical-68e47746092962d5ce01e813d1db0b0a5d9537d295fb2b0147daa13dd3d6dae1): complete subsection reference.

- [prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-06522db7fdb799b41b854f8821649a0d513946d7f372e11fd599463e88034595): complete subsection reference.

- [tls_list](data-sources--policy_based_routing--reference--group-001.md#canonical-506deaec95b4fe9cd4594c44d53aff5ab3451cc5bb2487262d51478f8c79f768): complete subsection reference.

<a id="canonical-9762b25f164d5ef545854a340dec4d57532374d4c355d7401b605d20f8f130e4"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules / c53be9d9ee46 / 4

- [forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations](data-sources--policy_based_routing--reference--group-001.md#canonical-6171bb1ebdf67b5399f835c91a31adf5a2bba15741984571d15da638a042a627)
- [forward_proxy_pbr.forward_proxy_pbr_rules.all_sources](data-sources--policy_based_routing--reference--group-001.md#canonical-e083eb0f522a1c30f3393d80342938858bbffd2ad38646feac3f64736d7b1c6a)
- [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-72201c762aacf19d586743867074af145919058d40a70adaa7c3a9494a09a6ea)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-7bbfc8405b63cdcbae107b961f216c9b3f43cb2a154cd2b7b4c24830032aab47)
- [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set](data-sources--policy_based_routing--reference--group-001.md#canonical-2c851c7847ce6d67448b4ca967aeee60e4fa9948ed8298d52db1b85dd410d4a9)
- [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector](data-sources--policy_based_routing--reference--group-001.md#canonical-f32637a485d2e418bd8573fd08d7d6116a869e8c9eb2d15ad43792d4a8035c36)
- [forward_proxy_pbr.forward_proxy_pbr_rules.metadata](data-sources--policy_based_routing--reference--group-001.md#canonical-68e47746092962d5ce01e813d1db0b0a5d9537d295fb2b0147daa13dd3d6dae1)
- [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-06522db7fdb799b41b854f8821649a0d513946d7f372e11fd599463e88034595)
- [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](data-sources--policy_based_routing--reference--group-001.md#canonical-506deaec95b4fe9cd4594c44d53aff5ab3451cc5bb2487262d51478f8c79f768)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-6171bb1ebdf67b5399f835c91a31adf5a2bba15741984571d15da638a042a627"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ac3533fc28dbe0f738bb81eb50e1b215629cfeff858ba1f60288bf6524501cf"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations — forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations / 970c029f921b / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations

<a id="canonical-0cdaaa07c84eafc0f2c21222cc637de6c4cb8edf884b592aa957efb49f4821a7"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all destinations.

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

<a id="canonical-f8cab5d410d8dbb0c01261c0f3c5f7a1ac50233c278a233d7347f3bb6dbbd3fc"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations / 970c029f921b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-763b1696995996aee0c2ce80ab86ceb126ab78ddd3a98df89c4ddb1e921fd9d0"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations / 970c029f921b / 4

- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-e083eb0f522a1c30f3393d80342938858bbffd2ad38646feac3f64736d7b1c6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f734411552a5e3fba49a70d3a5971f63b05a7282a61e4a77bc23bc4bb14f032f"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.all_sources — forward_proxy_pbr.forward_proxy_pbr_rules.all_sources / d51dafa15cd1 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- forward_proxy_pbr.forward_proxy_pbr_rules.all_sources

<a id="canonical-ea01b709ca013ea77fbc310635f4a77f49b8e41a4387fba09e677e77fc2e77cf"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all sources.

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

<a id="canonical-46c51ad8f33ecd842f31f51fbe3b8f417ba0ae6d43621163b431763e1c38cb81"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.all_sources / d51dafa15cd1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e700becbbed88fccf44f633422bbf4316f229438026e0205767130d10051ee60"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.all_sources / d51dafa15cd1 / 4

- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-72201c762aacf19d586743867074af145919058d40a70adaa7c3a9494a09a6ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f961087342e4708b915fdf652ba3e594f7ec81fc66abc45af6b7f234d52d203"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list — forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list / 9dec935d29ec / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list

<a id="canonical-13555c9a39125c7ed0e35d5f8a9ce5b2d8adccf1b35526c2c4c263c07fbfba6b"></a>

Type: `"list"`. Computed.

Ordered list of forwarding Class to be used if no rule match.

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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-4c5a861637ee048e4f19c44a9d4119ebcef656fe76369b026c07bf19254b3103"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list / 9dec935d29ec / 3

<a id="canonical-d7be537368b1a889b6a8e9899563f153e524d7086928ff542003c662a35e5f88"></a>

<a id="canonical-812795cf8c9007e6a9d2fc216514509decf83730423f29d1b792cd5b14d14d6b"></a>

## name property — forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list / 9dec935d29ec / 4

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

<a id="canonical-36fa836f74f1461e6b0d370d69875c2a33f9e6a8997dddee201abb973a64b5d4"></a>

<a id="canonical-5f0d025c7d01b625667551697ab8d9c38175c676f05c8c5120366be2ad418e1c"></a>

## namespace property — forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list / 9dec935d29ec / 5

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

<a id="canonical-3b87e2bf6f6a47119d3ef982b8b9ecdbfb6bd1b333ed9705b02f96e568029e64"></a>

<a id="canonical-cf7798327fd133ca3106214bb0205adff45895214284dbcff8da73c10d14ded4"></a>

## tenant property — forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list / 9dec935d29ec / 6

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

<a id="canonical-ee68d46410ed9934ee954128bd13ab3786b8e0ad376af53f8a74974ba44389d2"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list / 9dec935d29ec / 7

- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-7bbfc8405b63cdcbae107b961f216c9b3f43cb2a154cd2b7b4c24830032aab47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf92bfe8183c40119ebe86bd387d4570dfb3667fc60534e2486d7cbe025894c8"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.http_list — forward_proxy_pbr.forward_proxy_pbr_rules.http_list / 182fdbfd3ae6 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- forward_proxy_pbr.forward_proxy_pbr_rules.http_list

<a id="canonical-184b4eaeb8a92e05417c01033d10788c5cff91a80c4e6574e63759161b7c05f0"></a>

Type: `"single"`. Computed.

URLListType.

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

<a id="canonical-39dae158d4506a1d727be765a40101593d72a94a2e06b7f52da3467fbe973aab"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.http_list / 182fdbfd3ae6 / 3

- [http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-5b3a9931f2b25c13952ea94ab155fcd7f652938b778266cf2bf127c99dbda5c8): complete subsection reference.

<a id="canonical-207f9cd1fd134333e4988e8e7d74d7685b6aa542f8cf5be9c8afdebffb57d1ae"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.http_list / 182fdbfd3ae6 / 4

- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-5b3a9931f2b25c13952ea94ab155fcd7f652938b778266cf2bf127c99dbda5c8)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-5b3a9931f2b25c13952ea94ab155fcd7f652938b778266cf2bf127c99dbda5c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51f41e5f1c3d0b8fee92c066859d341578abb3e467334f0793f8a27f9ef1dd07"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / 692979413375 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-7bbfc8405b63cdcbae107b961f216c9b3f43cb2a154cd2b7b4c24830032aab47)
- forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list

<a id="canonical-8659d1f49311c1f9a8b12820d0e1d1db3f7b29bab29888ddff6ccab8a5f5da98"></a>

Type: `"list"`. Computed.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

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

<a id="canonical-80430a4c489889a9f4fc9f52cb28eb9a145ce971cca052c113131ab0215f9e56"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / 692979413375 / 3

- [any_path](data-sources--policy_based_routing--reference--group-001.md#canonical-fd63b2769ab19b657edef354449101cc2a8e56f57fa0cb2a22621a762639577c): complete subsection reference.

<a id="canonical-9cb2f451b6294f8e944cd6ddfa9df531704f19192639a936985ae4418b5d3c55"></a>

<a id="canonical-7f484b21b0c3726287d5ca33d98f5cf401f735e7677991722efaa5b651827cad"></a>

## exact_value property — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / 692979413375 / 4

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

<a id="canonical-e24e7e770e9136625787e96ee4a548b441c9552398cd08f538958b9f0e741301"></a>

<a id="canonical-6aed5ede6bb9044fcff88fb0232013d09a10f7302cba3ece82c0f374a383759c"></a>

## path_exact_value property — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / 692979413375 / 5

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-7eb89cb76d1de2c06d9fbafd839a7e9f65408a2607f8dfeea371915421276e9a"></a>

<a id="canonical-308799fa69ddc64fcd6b1b5b4a098d1c2b8db3568546016bd0a5587dc7b37df1"></a>

## path_prefix_value property — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / 692979413375 / 6

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-5d52dd0ed76ee492b769ddcbff0d4fd8e3465d476ee702df4983630ccc60eaf7"></a>

<a id="canonical-8f0c2a733be0892cee1be76fb69daa45c5aedc096505ae179a5e6a31902151c0"></a>

## path_regex_value property — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / 692979413375 / 7

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

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

<a id="canonical-236ecd62dbfccdb90e3a043a6ad80569dccfbf5f040d7ceb85175ba8d1ccef9c"></a>

<a id="canonical-7d0bf0547a3c822a41352f6dd3666793ec729ea1a585274cb1abbed5b5716ed9"></a>

## regex_value property — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / 692979413375 / 8

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

<a id="canonical-8d9496d9fbbca5e79d0e3304a11db00994e5247d7faa0e179057902b33510411"></a>

<a id="canonical-9a377dfa62c613c2153c5537ad2011f515f3eb59540f220804481346283ec6e1"></a>

## suffix_value property — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / 692979413375 / 9

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

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

<a id="canonical-db464394f7971275a7571b8e5db8312727e60c102e9f46a2e6b6da1b517e0db5"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / 692979413375 / 10

- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path](data-sources--policy_based_routing--reference--group-001.md#canonical-fd63b2769ab19b657edef354449101cc2a8e56f57fa0cb2a22621a762639577c)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-7bbfc8405b63cdcbae107b961f216c9b3f43cb2a154cd2b7b4c24830032aab47)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-fd63b2769ab19b657edef354449101cc2a8e56f57fa0cb2a22621a762639577c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-813a6bab4752a08f6fd1e62c5e8f92d4ad27280d3bea685e54ddd1e332e89190"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path / b801ee99b64b / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-7bbfc8405b63cdcbae107b961f216c9b3f43cb2a154cd2b7b4c24830032aab47)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-5b3a9931f2b25c13952ea94ab155fcd7f652938b778266cf2bf127c99dbda5c8)
- forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path

<a id="canonical-8fbff81e669f86b440c328e6552b59e618d8b6ec26da346907945dda825a430e"></a>

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

<a id="canonical-f941ece8009f3afcabd054fdf233b41ecc1cf3c3143ce99d4030e6d9d6433373"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path / b801ee99b64b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e98cf9d3b7057d8384c5778807c09444ca472ff6f6261a1a031061a26e77a12e"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path / b801ee99b64b / 4

- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-5b3a9931f2b25c13952ea94ab155fcd7f652938b778266cf2bf127c99dbda5c8)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-2c851c7847ce6d67448b4ca967aeee60e4fa9948ed8298d52db1b85dd410d4a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd3ec68b86a114e8ca106ab3e5185f4447bfec5871e26700c98e148b0782dad5"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set — forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set / 24d497792e75 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set

<a id="canonical-fc7528b27302c2fb29eddeb742c80ed04bbf13bd55aa74cb49484ae73ed16811"></a>

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

<a id="canonical-10e4f7a68d5e03abc3754728b3851d81edd355ec2c78c7bcdf5696cafca793c5"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set / 24d497792e75 / 3

<a id="canonical-d18e11bd2fcf2817d90fa94c21cb050a6b2b6c9ab9993490c0924c1182da02f9"></a>

<a id="canonical-a0a059a07c057ef21f95af29f67730fe35faf640a96a6723df6173847f4b41a6"></a>

## name property — forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set / 24d497792e75 / 4

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

<a id="canonical-89a6fc8f173287da80252d10383923768dcc53a9a02dad3bcbae5d5dfe3e15b3"></a>

<a id="canonical-f1d5c8dacb34d8cea4f787a130962806e23ec30544ffc8a3681c51b2a132b192"></a>

## namespace property — forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set / 24d497792e75 / 5

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

<a id="canonical-c98bca840512e209de15817a8c540042ad433ad7f2f46269621bc5affb29d5df"></a>

<a id="canonical-a19aa72967ef75d713698a63f23d38a9783eb4739e8a19739f6f86ce88cb4502"></a>

## tenant property — forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set / 24d497792e75 / 6

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

<a id="canonical-2073bee5eb5386b78598bd38dbddeafdc768398a397cd278bb85d56894f7dfa0"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set / 24d497792e75 / 7

- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-f32637a485d2e418bd8573fd08d7d6116a869e8c9eb2d15ad43792d4a8035c36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41ce455aa62b75d1785e10a581cf0e0f12c874465666f1995a365f76cef8e7e8"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.label_selector — forward_proxy_pbr.forward_proxy_pbr_rules.label_selector / 633d0ba77d9c / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- forward_proxy_pbr.forward_proxy_pbr_rules.label_selector

<a id="canonical-e2c4812f61066bbd4c9c5e6aaf1741a4d9ad070593701e9091b82c6fbd37b8f5"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-cd42c8fc5a414e0b8c1ee950d7f64e85c14580e3c1e138ce10453175588f7a29"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.label_selector / 633d0ba77d9c / 3

<a id="canonical-7a533a8bd5e7f2ffd67e915ae4e2412b0ef84e839ae0d820133ba5b152d7875b"></a>

<a id="canonical-b9f0015e28d020bd80187d7fe5fdcd549557626185495e59f8b18c85b8e2ece8"></a>

## expressions property — forward_proxy_pbr.forward_proxy_pbr_rules.label_selector / 633d0ba77d9c / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-9d5b4f3757b74ee1239463f0ea58318a43bcb20111248ad6370e758f49f1c7ff"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.label_selector / 633d0ba77d9c / 5

- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-68e47746092962d5ce01e813d1db0b0a5d9537d295fb2b0147daa13dd3d6dae1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ead2402b3ce8c4580eb422686352f4321c6fed3ea97991dcfe60efc6dbbe4baa"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.metadata — forward_proxy_pbr.forward_proxy_pbr_rules.metadata / 1c0eb7e81979 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- forward_proxy_pbr.forward_proxy_pbr_rules.metadata

<a id="canonical-882e22f3385e35b608b16777f0a27b04c6060178c8f172ece34b4e781cd44713"></a>

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

<a id="canonical-6915acedec997e839b08e6ed712c1db70247b025f5be4a248a714ca5fa9a48de"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.metadata / 1c0eb7e81979 / 3

<a id="canonical-e4741f57f567f3ddeedc8bdbbf0820147d8a3d48ecc4e7180289813e994bc816"></a>

<a id="canonical-947e170f4ea365c9e01bfda1b12c5fa2ee6f2fbdd6a371b01e079c2a5a63fb9b"></a>

## description_spec property — forward_proxy_pbr.forward_proxy_pbr_rules.metadata / 1c0eb7e81979 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-05ee124fd531cbd928a40066a14e69da1e4dbd7a252b50ce684f07b7a0ee1e2c"></a>

<a id="canonical-223dca9d0cf965e02c1935934d9ee912f858cf0742b346dffda7fc93e325396c"></a>

## name property — forward_proxy_pbr.forward_proxy_pbr_rules.metadata / 1c0eb7e81979 / 5

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

<a id="canonical-847cc65b0f28e681ad28172e090ee9f5ab352c93c2993ff6225930d492f050b3"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.metadata / 1c0eb7e81979 / 6

- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-06522db7fdb799b41b854f8821649a0d513946d7f372e11fd599463e88034595"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c410a1f1471bc42db782374313dcee4bb1631a2d2d5477103f16fe0d47c93597"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list — forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list / d819ee4a26b6 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list

<a id="canonical-5ec6f762eb6a3e37a2435c87fcfb3331bc6f28d8a4e274fde3c1c03c340aaa8f"></a>

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

<a id="canonical-cf74173fc938bacbf710cdfe64d023d71c521120ed34419e0e99500e0a02fc00"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list / d819ee4a26b6 / 3

<a id="canonical-3a0e299ee52170cfa2f23a063271cfc3558d3e7816f869811164e9fdc5239775"></a>

<a id="canonical-c72c4f67299788a81825bd570da6d541bd8a51b3b1c6073b88aa80e6a6ef7423"></a>

## prefixes property — forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list / d819ee4a26b6 / 4

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

<a id="canonical-dfe0cec495afbf69e7502e51bcdd11b599ed1c288fc405b0a6a818774cde12af"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list / d819ee4a26b6 / 5

- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-506deaec95b4fe9cd4594c44d53aff5ab3451cc5bb2487262d51478f8c79f768"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c44bf68bc357b94d388176c636f65a2548039fa16e0a1053fcb0d5e381fe85af"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.tls_list — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list / a066110008dc / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- forward_proxy_pbr.forward_proxy_pbr_rules.tls_list

<a id="canonical-fc68e11402e79d093af1147ad8497c182fb8a24fee4c7160344dd7baaf2d2d58"></a>

Type: `"single"`. Computed.

DomainListType.

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

<a id="canonical-3112c56fee7276432cc914b18082ca60fa4ea5de1976bc5ff93ad3862c325a7d"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list / a066110008dc / 3

- [tls_list](data-sources--policy_based_routing--reference--group-001.md#canonical-f63e47b6686ab468f068ae12bf3d414ca55fe7e81f0cfc6be42206242c121e73): complete subsection reference.

<a id="canonical-1a0b46e70c731ef2e5a889fe3d0776d558e7f6997517e22190d3ca343104ffc3"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list / a066110008dc / 4

- [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list](data-sources--policy_based_routing--reference--group-001.md#canonical-f63e47b6686ab468f068ae12bf3d414ca55fe7e81f0cfc6be42206242c121e73)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-f63e47b6686ab468f068ae12bf3d414ca55fe7e81f0cfc6be42206242c121e73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76b7fbc2ac008d0a5edca2a9c52397dad0b3062f427e5938d4c9f6d212b8b5ee"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list / c7f69802292b / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-ebbd6e66492851c4fb87ea6b8c4d4dee3ea4e7910e068697b769a204f9ae5d32)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-8fb6e2fa5ba097815707aa1afa480e93d794bcffcaf40bfa6c6cd607089b915a)
- [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](data-sources--policy_based_routing--reference--group-001.md#canonical-506deaec95b4fe9cd4594c44d53aff5ab3451cc5bb2487262d51478f8c79f768)
- forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list

<a id="canonical-cd4bf44e645ca693aa812dde8756fd4889e9e0577d7e0ecb2c9ed8be42e24c98"></a>

Type: `"list"`. Computed.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

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

<a id="canonical-da52c94a41e0bfc5c991fbacc924cde81b0e4f3ae078d5704bd0675567ed29f5"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list / c7f69802292b / 3

<a id="canonical-2c571f151e604358c5aa36f6aa28e95d1fd6894342d6fd145b5550cb0d80fe3b"></a>

<a id="canonical-587f922576e98d6d747a88c00004591baf18af737f9f6d2324dc0e32d665d4a0"></a>

## exact_value property — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list / c7f69802292b / 4

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

<a id="canonical-b3e8a03fe929d48f4995e56e00e25f0ab526e63ccce75d976c552fd4cbf7ff98"></a>

<a id="canonical-18a21250f9f76db0a11e910166db8ba075bff036e5b779958e2923fc862abf48"></a>

## regex_value property — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list / c7f69802292b / 5

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

<a id="canonical-27a45b5b54b4d4ec3c584b3a0494d84f60d759a7531cd1d2c873e5e622cbb8f6"></a>

<a id="canonical-7a76544f2f34dc7bd0c20e4243ac2c38537c3ae6cd1f98d295eca99d479e8771"></a>

## suffix_value property — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list / c7f69802292b / 6

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

<a id="canonical-5411922a0f038a5a86e1e675bb1ab9ddbe31089c58f72293ffa5a2346191a932"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list / c7f69802292b / 7

- [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](data-sources--policy_based_routing--reference--group-001.md#canonical-506deaec95b4fe9cd4594c44d53aff5ab3451cc5bb2487262d51478f8c79f768)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-00df9106da3d9248c38a0af546df4f2548c586620425b7859c7f7583a1e0012e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7aaa5b3aaa3f642e2640221265e4f4131ff862d5dd5736ee7053458f7811f5de"></a>

## forwarding_class_list — forwarding_class_list / 911e217a7e95 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- forwarding_class_list

<a id="canonical-7125047be099209e535432dfe5dec37b4afb8ccbfc2bcdce248dacb47bc8e5d8"></a>

Type: `"list"`. Computed.

Ordered list of forwarding Class to be used if source application match and no rule match.

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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-851be6357c8df1a4a2a19cd183d270e5ed8a662179b27eeb51366d0b4e3f9443"></a>

## Direct properties — forwarding_class_list / 911e217a7e95 / 3

<a id="canonical-c737c0f7fbedec79778fe952772559e4bf85c6fc9803792b598a3eb7f631b8d2"></a>

<a id="canonical-df72f0a813212a3888f63fe72ab85cfb30903e6d1189a45cbdd2a783164c74c1"></a>

## name property — forwarding_class_list / 911e217a7e95 / 4

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

<a id="canonical-cffbb1f52c1803c9421659478c55c46be4a4355915f454172671f9fedc0f7e06"></a>

<a id="canonical-894cc901226ee7bbe90b18e51dfa7c89b7bd79ff77ec0728a7ff04263a5c8476"></a>

## namespace property — forwarding_class_list / 911e217a7e95 / 5

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

<a id="canonical-67bd05388ffc3a453a391818fa0bffc034f96f2d40cb43d52184df8de6587be8"></a>

<a id="canonical-e31e584b63521a0ae31a4332a514966a36927b57698629e1e782119b29e5e943"></a>

## tenant property — forwarding_class_list / 911e217a7e95 / 6

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

<a id="canonical-34b6b100748ee4b4654bd4a1c67e1b66d9e82234715c82794d96749c799d0d01"></a>

## Next pages — forwarding_class_list / 911e217a7e95 / 7

- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa247fc2dfbf25199677b8abd440cb56a0d61265a065a9371247aa7e0a0443c3"></a>

## network_pbr — network_pbr / a06991cf8c48 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- network_pbr

<a id="canonical-fd04ccc60307e88d9463a971f5d7df1a85169d821589db16d46bf2b9b8dbba03"></a>

Type: `"single"`. Computed.

Configuration parameter for network pbr.

Upstream description:

Network(L3/L4) routing policy rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-source_choice": "[\"any\",\"label_selector\",\"prefix_list\"]"
}
```

<a id="canonical-61f07bf1eeb3f8c17ffef4ce1b781d44c60e61d7c3b2f4c25483997381fdd388"></a>

## Direct properties — network_pbr / a06991cf8c48 / 3

- [any](data-sources--policy_based_routing--reference--group-001.md#canonical-c7c3c789fb1dbec7638025637d05679527f51bc52de51da2e719722cfc6b2fb3): complete subsection reference.

- [label_selector](data-sources--policy_based_routing--reference--group-001.md#canonical-1ae9605c277239293201bff4ecd460a74d79e4846fd881a06faef4d338cdc349): complete subsection reference.

- [network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901): complete subsection reference.

- [prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-2e666d0078404f5978dae6c8a24c2acec20583e056d394420b3608a2f79fe274): complete subsection reference.

<a id="canonical-c3d4ce8f969badbb0e12f339ce00ade0528fe0a2f1c4bf466895d5f568a536c3"></a>

## Next pages — network_pbr / a06991cf8c48 / 4

- [network_pbr.any](data-sources--policy_based_routing--reference--group-001.md#canonical-c7c3c789fb1dbec7638025637d05679527f51bc52de51da2e719722cfc6b2fb3)
- [network_pbr.label_selector](data-sources--policy_based_routing--reference--group-001.md#canonical-1ae9605c277239293201bff4ecd460a74d79e4846fd881a06faef4d338cdc349)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- [network_pbr.prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-2e666d0078404f5978dae6c8a24c2acec20583e056d394420b3608a2f79fe274)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-c7c3c789fb1dbec7638025637d05679527f51bc52de51da2e719722cfc6b2fb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ed1716f3d25c5cd615c623dbafb2e77fdff950e39ec9bda33687160af49dccc"></a>

## network_pbr.any — network_pbr.any / 1ebc84423a96 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- network_pbr.any

<a id="canonical-430b041e36eaeb3ab1e087d17a9c353d904864303e2afe1f5f4dbe02a9e4452d"></a>

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

<a id="canonical-04e1a550eabc1f1d56890b5e13347bc51c653ddba77d6fe2eb2ba79bd18e7bdb"></a>

## Direct properties — network_pbr.any / 1ebc84423a96 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-afeadbdb3c18edfdf7ace52fa2ba8d2f6fd8a4b5dba016148e5b6ce7fe6b43bb"></a>

## Next pages — network_pbr.any / 1ebc84423a96 / 4

- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-1ae9605c277239293201bff4ecd460a74d79e4846fd881a06faef4d338cdc349"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d4fd648c00bef4d8a54576db5f289ffb27f77dd8ee604fe853b80f4c9b4845c"></a>

## network_pbr.label_selector — network_pbr.label_selector / 172d21e56bba / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- network_pbr.label_selector

<a id="canonical-156a1784f386717797403c94f7be1c5b90253e7095c40f7fad8035dbabec8473"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-4693b85992952d642eceda787b3a0c50d8dcd30142f20ca522c296226a18832b"></a>

## Direct properties — network_pbr.label_selector / 172d21e56bba / 3

<a id="canonical-31990e30529bdbce62c36fa3106f58166e3a99792d01efd5986dba9acb24ceef"></a>

<a id="canonical-063cee3639245872d7b4af84e6b265b2d9fa05e91e2edbe9f14c500ff7248c85"></a>

## expressions property — network_pbr.label_selector / 172d21e56bba / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-543eb618ce7a5b9d75502895c04b90080aaeb42e5a791ea0371c327c08e4cc83"></a>

## Next pages — network_pbr.label_selector / 172d21e56bba / 5

- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ca51baaa56e7caec38fc7420247e28106af5d61545b18e8cec65d64084a52f2"></a>

## network_pbr.network_pbr_rules — network_pbr.network_pbr_rules / 8e79d5ad98c8 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- network_pbr.network_pbr_rules

<a id="canonical-a3d30a9b1b0ab258d82e4ac2ef12849a2152deb3c09bf35ea25ea0151c2c4d74"></a>

Type: `"list"`. Computed.

L3/L4 Destination Routing Rules. Network(L3/L4) routing policy rule.

Upstream description:

Network(L3/L4) routing policy rule.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-80401e8ba41b38d041120051b8cabee431622d88fb055d711eb48e2dd273ace6"></a>

## Direct properties — network_pbr.network_pbr_rules / 8e79d5ad98c8 / 3

- [all_tcp_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-a26fe445d24c64057cc45c2f4189f91d6a627d4416e201acbaa66b0b239bf171): complete subsection reference.

- [all_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-b80267cd0f32f58dec35b6ba9e2bfa0913d5a9aca700ec8bf32085064b708ba8): complete subsection reference.

- [all_udp_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-1c0e695a087a952ab5eff5b63adf7b76a8d46824c684463813139495513f1379): complete subsection reference.

- [any](data-sources--policy_based_routing--reference--group-001.md#canonical-fe90b13302475b5582f77109a4b23e15938fa1a674eceac711920e38ba79b81f): complete subsection reference.

- [applications](data-sources--policy_based_routing--reference--group-001.md#canonical-9b4c7fcba04fa3439d91ea391ee12a8b739a4cfe277289fb111f86a4c28c15e4): complete subsection reference.

<a id="canonical-a1ad512aaf920dc41ad9f616f9fe36b8ecf0bafddb5dd3d12c44c583351598c4"></a>

<a id="canonical-e86a8ee344c316c84e6405f2792b39e53b7007fbf6955e3ffb0ed6508adb91d6"></a>

## dns_name property — network_pbr.network_pbr_rules / 8e79d5ad98c8 / 4

Type: `"string"`. Computed.

Exclusive with \[any ip\_prefix\_set prefix\_list\] Resolve hostname to GET the IP.

Upstream description:

Exclusive with \[any ip\_prefix\_set prefix\_list\] Resolve hostname to GET the IP.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 1024,
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

- [forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-11b224295f34f2aa650a7b215e5cb5fc035e607375440c66424d5ed66bd9ce88): complete subsection reference.

- [ip_prefix_set](data-sources--policy_based_routing--reference--group-001.md#canonical-de16819c0762a93d7942baf43d98b18b396f3dab8ec9bd375d7ef7d66bc6ca4f): complete subsection reference.

- [metadata](data-sources--policy_based_routing--reference--group-001.md#canonical-258b712bbb53f1846a9389d2abd1047a36eb318dcfdce3a8ad3fa8852a73c27a): complete subsection reference.

- [prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-6102ddf8620fd347ff0e8fda974797faa35d4a12f5c29a7ff4bcbc52a946228a): complete subsection reference.

- [protocol_port_range](data-sources--policy_based_routing--reference--group-001.md#canonical-715cb6f11d295f65fb64f92febd3ac53d3276808be6b9d074bc3d1d2519a5254): complete subsection reference.

<a id="canonical-61363067266b85ea67fa2cf4e1ca4a5d210931722155500ffc501289126c075f"></a>

## Next pages — network_pbr.network_pbr_rules / 8e79d5ad98c8 / 5

- [network_pbr.network_pbr_rules.all_tcp_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-a26fe445d24c64057cc45c2f4189f91d6a627d4416e201acbaa66b0b239bf171)
- [network_pbr.network_pbr_rules.all_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-b80267cd0f32f58dec35b6ba9e2bfa0913d5a9aca700ec8bf32085064b708ba8)
- [network_pbr.network_pbr_rules.all_udp_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-1c0e695a087a952ab5eff5b63adf7b76a8d46824c684463813139495513f1379)
- [network_pbr.network_pbr_rules.any](data-sources--policy_based_routing--reference--group-001.md#canonical-fe90b13302475b5582f77109a4b23e15938fa1a674eceac711920e38ba79b81f)
- [network_pbr.network_pbr_rules.applications](data-sources--policy_based_routing--reference--group-001.md#canonical-9b4c7fcba04fa3439d91ea391ee12a8b739a4cfe277289fb111f86a4c28c15e4)
- [network_pbr.network_pbr_rules.forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-11b224295f34f2aa650a7b215e5cb5fc035e607375440c66424d5ed66bd9ce88)
- [network_pbr.network_pbr_rules.ip_prefix_set](data-sources--policy_based_routing--reference--group-001.md#canonical-de16819c0762a93d7942baf43d98b18b396f3dab8ec9bd375d7ef7d66bc6ca4f)
- [network_pbr.network_pbr_rules.metadata](data-sources--policy_based_routing--reference--group-001.md#canonical-258b712bbb53f1846a9389d2abd1047a36eb318dcfdce3a8ad3fa8852a73c27a)
- [network_pbr.network_pbr_rules.prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-6102ddf8620fd347ff0e8fda974797faa35d4a12f5c29a7ff4bcbc52a946228a)
- [network_pbr.network_pbr_rules.protocol_port_range](data-sources--policy_based_routing--reference--group-001.md#canonical-715cb6f11d295f65fb64f92febd3ac53d3276808be6b9d074bc3d1d2519a5254)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-a26fe445d24c64057cc45c2f4189f91d6a627d4416e201acbaa66b0b239bf171"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56ab4fa27b0f4546caf6488e1d8289be1650fd3d91e467bcc327a7ece59c82f6"></a>

## network_pbr.network_pbr_rules.all_tcp_traffic — network_pbr.network_pbr_rules.all_tcp_traffic / 17b07cda05a3 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- network_pbr.network_pbr_rules.all_tcp_traffic

<a id="canonical-82855584f0cf011acfa049273ba57e93c8334a70341872a0383d632be429509e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all tcp traffic.

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

<a id="canonical-cb7cbb4deafe34df374fc82d1c1f2af94ad10e0fede8c73b4b74db23be4693c9"></a>

## Direct properties — network_pbr.network_pbr_rules.all_tcp_traffic / 17b07cda05a3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c92327f249360c36a1bd740189b1f07748d6b87f24b02f12c475a44aed12f345"></a>

## Next pages — network_pbr.network_pbr_rules.all_tcp_traffic / 17b07cda05a3 / 4

- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-b80267cd0f32f58dec35b6ba9e2bfa0913d5a9aca700ec8bf32085064b708ba8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bde9405e44d1f385d03908ef3c6c7636a5f8feaeec3bc9a5eb8dc666f2de923a"></a>

## network_pbr.network_pbr_rules.all_traffic — network_pbr.network_pbr_rules.all_traffic / 6499b90b537d / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- network_pbr.network_pbr_rules.all_traffic

<a id="canonical-4eb2dac0ae9ac5a10725c53ca4dfb394e1d54c3613f74de35afe2ec30b0ba99f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all traffic.

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

<a id="canonical-7fb02d37698561368f62083549d87d155652773b2323ee515519103bebf1b326"></a>

## Direct properties — network_pbr.network_pbr_rules.all_traffic / 6499b90b537d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-008083d882f588f97942824a174dc7c9369e3efb59c01538f25d094743258a32"></a>

## Next pages — network_pbr.network_pbr_rules.all_traffic / 6499b90b537d / 4

- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-1c0e695a087a952ab5eff5b63adf7b76a8d46824c684463813139495513f1379"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21e58a4e10eb1893f3fac615d0c3c1d59f814998655b0410f2c52911b25524e6"></a>

## network_pbr.network_pbr_rules.all_udp_traffic — network_pbr.network_pbr_rules.all_udp_traffic / 29a4014b5d63 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- network_pbr.network_pbr_rules.all_udp_traffic

<a id="canonical-35fb8193aa430a785a9378f1c35c49272e0d67c17723ab2999760eb7eb26a093"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all udp traffic.

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

<a id="canonical-5022ed435e53c0fdc45f6973ea3136f9d92c8bebea16ddf461f6b5ea9a5c0746"></a>

## Direct properties — network_pbr.network_pbr_rules.all_udp_traffic / 29a4014b5d63 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8d234626d08b83e7681ffa6c103c0bd048867003ae647c6b66942924cbca133a"></a>

## Next pages — network_pbr.network_pbr_rules.all_udp_traffic / 29a4014b5d63 / 4

- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-fe90b13302475b5582f77109a4b23e15938fa1a674eceac711920e38ba79b81f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-722945886835aed071c3794d7cd6f176be2a9ff574aa177bd0feb5815904fdca"></a>

## network_pbr.network_pbr_rules.any — network_pbr.network_pbr_rules.any / beeaef51b066 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- network_pbr.network_pbr_rules.any

<a id="canonical-bf893380bc5ec146d7e9b665c23a192c82ef962356091e56dc19dab0df37bccb"></a>

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

<a id="canonical-40415d72aad293a2659ec313b2857a7db65f95ff94c744aba48971c5e36f5fac"></a>

## Direct properties — network_pbr.network_pbr_rules.any / beeaef51b066 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b90c761049a20f0841ed8d1207da5a0a944af74556361a79281e57e655f846ff"></a>

## Next pages — network_pbr.network_pbr_rules.any / beeaef51b066 / 4

- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-9b4c7fcba04fa3439d91ea391ee12a8b739a4cfe277289fb111f86a4c28c15e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0db23e8f5071c4b4b741cc4b53b4afc91c0ae5060e85848a10cb57bed89740c"></a>

## network_pbr.network_pbr_rules.applications — network_pbr.network_pbr_rules.applications / f81303102e65 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- network_pbr.network_pbr_rules.applications

<a id="canonical-175a063f9e522387be6b671f3410684bdba1d2906e58c238829e33e1f2e89fdc"></a>

Type: `"single"`. Computed.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="canonical-397176226c9c59c5c63f71e6327699d284782b75d284cfac2e7e4dd2ea234db0"></a>

## Direct properties — network_pbr.network_pbr_rules.applications / f81303102e65 / 3

<a id="canonical-cd6467c9ca6ee36d1b64de466d1dd3c422703d21ccb851d3876dc170c210da1d"></a>

<a id="canonical-5a0c5e62fa66ae9e77d3904328eeb23f1c0c8a82076a5562a945d6ee7f76d68e"></a>

## applications property — network_pbr.network_pbr_rules.applications / f81303102e65 / 4

Type: `["list", "string"]`. Computed.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="canonical-22a2d179f0c34e9bbdf48a391df6b26a682a3fe27767c7501a75ce246894141b"></a>

## Next pages — network_pbr.network_pbr_rules.applications / f81303102e65 / 5

- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-11b224295f34f2aa650a7b215e5cb5fc035e607375440c66424d5ed66bd9ce88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afdf7c08f4a8f6a4102681476ffa91bd617246a7b929d60b0ad282b763375c64"></a>

## network_pbr.network_pbr_rules.forwarding_class_list — network_pbr.network_pbr_rules.forwarding_class_list / f561e3410cfe / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- network_pbr.network_pbr_rules.forwarding_class_list

<a id="canonical-26e0ebee13d9739fec706f26b4dd04381ccfb9d3a0f426e5c046cc9f12ff491b"></a>

Type: `"list"`. Computed.

Ordered list of forwarding Class to be used if rule match.

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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-20679d4052062f06c6c5b3c569a42a6412c5662049925b02b7bb6382a3a370c7"></a>

## Direct properties — network_pbr.network_pbr_rules.forwarding_class_list / f561e3410cfe / 3

<a id="canonical-42d260713981fa7aadb4b269580bbfc89b8142d013eeb45164cac65d1cd8d00c"></a>

<a id="canonical-888f025e8ed157c915f276df90d776f9064dbce4dfdd1805a96d067f9c717320"></a>

## name property — network_pbr.network_pbr_rules.forwarding_class_list / f561e3410cfe / 4

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

<a id="canonical-61bccef6ce0a17fca8cb1e47c8bbdcce032484537630751688c426a20b81e136"></a>

<a id="canonical-24f87cf7c9fc61ad521445a9f6d6e1c9c7d49c487b77e501c2d64afd558b3e17"></a>

## namespace property — network_pbr.network_pbr_rules.forwarding_class_list / f561e3410cfe / 5

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

<a id="canonical-f1ac34138432d0b344a155f0ebcbfc6b4d8ecc95e50d6d7e4659ca0b0b7422b3"></a>

<a id="canonical-bf6f224b240b0476a72440879c1a75406bd04eb60f96d126dd33b95e3af5e56b"></a>

## tenant property — network_pbr.network_pbr_rules.forwarding_class_list / f561e3410cfe / 6

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

<a id="canonical-bfb282089649ab06d8c6c3218fa958499ccd85f1befe6ac65fc58c5583fecc81"></a>

## Next pages — network_pbr.network_pbr_rules.forwarding_class_list / f561e3410cfe / 7

- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-de16819c0762a93d7942baf43d98b18b396f3dab8ec9bd375d7ef7d66bc6ca4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15042d14f94ac28ba09fa56901dfea72081a898e2470bf395187ca9a2a7bfeb5"></a>

## network_pbr.network_pbr_rules.ip_prefix_set — network_pbr.network_pbr_rules.ip_prefix_set / ed044f1c2ebe / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- network_pbr.network_pbr_rules.ip_prefix_set

<a id="canonical-7ad889af4e73d0ff297bce16035ff989905fcdc6f0ff6fd6e6b2d4041c78d103"></a>

Type: `"single"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

<a id="canonical-6468b1d56f6efa4dcf2cb8fe03cc6518de36b5dd9390ad7bd93dd6ccf5f2e7eb"></a>

## Direct properties — network_pbr.network_pbr_rules.ip_prefix_set / ed044f1c2ebe / 3

- [ref](data-sources--policy_based_routing--reference--group-001.md#canonical-65cda249e093db8b4810663c36a70536c1bf53f9db6b42307e3b7fdb7222a2d4): complete subsection reference.

<a id="canonical-3ad5107323951b254e0700ed148a8ba8f3881d56913ad63f3a1a990d98cd8ee4"></a>

## Next pages — network_pbr.network_pbr_rules.ip_prefix_set / ed044f1c2ebe / 4

- [network_pbr.network_pbr_rules.ip_prefix_set.ref](data-sources--policy_based_routing--reference--group-001.md#canonical-65cda249e093db8b4810663c36a70536c1bf53f9db6b42307e3b7fdb7222a2d4)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-65cda249e093db8b4810663c36a70536c1bf53f9db6b42307e3b7fdb7222a2d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd70459531bb2f5d3adbac41488ce160143a184d26a1650773b091eaae7d2b05"></a>

## network_pbr.network_pbr_rules.ip_prefix_set.ref — network_pbr.network_pbr_rules.ip_prefix_set.ref / 0a1ac118cd1a / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- [network_pbr.network_pbr_rules.ip_prefix_set](data-sources--policy_based_routing--reference--group-001.md#canonical-de16819c0762a93d7942baf43d98b18b396f3dab8ec9bd375d7ef7d66bc6ca4f)
- network_pbr.network_pbr_rules.ip_prefix_set.ref

<a id="canonical-95476b283c9fb2fcbdec29487f9741dfe7ad68e4bc46e11ab44874751efa3dd5"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-27b54f3c5010816fe3e30de98eddb08da95c09a904a8d2f0b35fcc3eb6f09cae"></a>

## Direct properties — network_pbr.network_pbr_rules.ip_prefix_set.ref / 0a1ac118cd1a / 3

<a id="canonical-912cd407119ba82eb97e5717247570712d2603c07071f191cd0901283a0c99ea"></a>

<a id="canonical-eeb2c2d0f5d03ce5d6ef6585c51f8a37d82870149ce714dde568025a2a148ffa"></a>

## kind property — network_pbr.network_pbr_rules.ip_prefix_set.ref / 0a1ac118cd1a / 4

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

<a id="canonical-9e9edcb5052edd851b1f274af730c02ba94b65c0b5546225e840cd49665e5ffc"></a>

<a id="canonical-31584620c2f2a4af97cf115bfcd37f3d3082f304295f38e0724e13f3d92735d4"></a>

## name property — network_pbr.network_pbr_rules.ip_prefix_set.ref / 0a1ac118cd1a / 5

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

<a id="canonical-8926721449c19cfe9df40209a7d0e73239d6f2d24ca416e9a76e43d3c50ea574"></a>

<a id="canonical-fab9b4089c8c5d7abad7c7c0bc012fb95faeae9e337ac7298e58517da092bec2"></a>

## namespace property — network_pbr.network_pbr_rules.ip_prefix_set.ref / 0a1ac118cd1a / 6

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

<a id="canonical-660cfb277b528b749565200b706e67ec3d6d2f54e6977126d13e1b82e984eb6b"></a>

<a id="canonical-5a833bae5b62954670605b3bb9808dcc2ff6081cce9e46fdf9bdc7286b245ad8"></a>

## tenant property — network_pbr.network_pbr_rules.ip_prefix_set.ref / 0a1ac118cd1a / 7

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

<a id="canonical-5faa6725e40df0284fc2a52f3309c0d04e6d1e491040e98b56cd088d245b0ff3"></a>

<a id="canonical-8d149217ccb7a7e75db998b929d7621ee3ada4ddc01293396ef89f1fd511a39c"></a>

## uid property — network_pbr.network_pbr_rules.ip_prefix_set.ref / 0a1ac118cd1a / 8

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

<a id="canonical-07e2b116c7391482ad9a916973218c750e64361e9382d6bfeba05c490cac2cc3"></a>

## Next pages — network_pbr.network_pbr_rules.ip_prefix_set.ref / 0a1ac118cd1a / 9

- [network_pbr.network_pbr_rules.ip_prefix_set](data-sources--policy_based_routing--reference--group-001.md#canonical-de16819c0762a93d7942baf43d98b18b396f3dab8ec9bd375d7ef7d66bc6ca4f)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-258b712bbb53f1846a9389d2abd1047a36eb318dcfdce3a8ad3fa8852a73c27a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-148ee75118d04568cbb23c48523e5f03b9ea17a24e6ba61304ae5b76dce6481c"></a>

## network_pbr.network_pbr_rules.metadata — network_pbr.network_pbr_rules.metadata / 95be7d58ba33 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- network_pbr.network_pbr_rules.metadata

<a id="canonical-37f9c9c5b5c5dc6021b9352610bcdfa2deec1f967bebfc1903ccde62aa3fbfdb"></a>

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

<a id="canonical-e4b14a4d5d5040c6a003fb74cf9248ce0e744e7fff8c31cc9caa495d456fe756"></a>

## Direct properties — network_pbr.network_pbr_rules.metadata / 95be7d58ba33 / 3

<a id="canonical-aa28934e23b48f10d21805b0354d8e16d32e9f8d2ab160e3ff247f2bcc4bfcc3"></a>

<a id="canonical-3ba152609932815e32ea48e6bc49bc01cc6ac1d65c27e96424524444454a5bfb"></a>

## description_spec property — network_pbr.network_pbr_rules.metadata / 95be7d58ba33 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-8054927c92edb8ea2149f756ed06b3a9846c462dac840500210cea07f566f18d"></a>

<a id="canonical-cda235fab6bf316a4d3d145120169e88e9cff794537fa1743322b341f101ebdd"></a>

## name property — network_pbr.network_pbr_rules.metadata / 95be7d58ba33 / 5

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

<a id="canonical-cdf97e8cf58487c0c2c01ae5c4cf5a2f22218cf57abd64d42d528a111d9530cd"></a>

## Next pages — network_pbr.network_pbr_rules.metadata / 95be7d58ba33 / 6

- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-6102ddf8620fd347ff0e8fda974797faa35d4a12f5c29a7ff4bcbc52a946228a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ad567cdc41b00c80787df7aed4b180d0fde5496b8bf7e9f43f8ba99df71b2e9"></a>

## network_pbr.network_pbr_rules.prefix_list — network_pbr.network_pbr_rules.prefix_list / 12a6789583aa / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- network_pbr.network_pbr_rules.prefix_list

<a id="canonical-4083e0eb6e062fde7e8d9432e94e582242a49994320574de554ebec9ce30691c"></a>

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

<a id="canonical-c9615eaa125061463269937bb3bcf721389721da53f317d65d91f9cabe6a3e9a"></a>

## Direct properties — network_pbr.network_pbr_rules.prefix_list / 12a6789583aa / 3

<a id="canonical-57909b01ac7931c9a3782ff5b51a11de8db74f621af130ac6a6d9da8a9a4ee06"></a>

<a id="canonical-c72b6a5777448b5f6f2f295d585bc43d8ab294a21915a8f8cd77ac0f7d1c90d1"></a>

## prefixes property — network_pbr.network_pbr_rules.prefix_list / 12a6789583aa / 4

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

<a id="canonical-c84230b99ac60c2a061871082470cc29578f111bf52b126526129871745440bc"></a>

## Next pages — network_pbr.network_pbr_rules.prefix_list / 12a6789583aa / 5

- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-715cb6f11d295f65fb64f92febd3ac53d3276808be6b9d074bc3d1d2519a5254"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b345d71530d2967a6712a8cb1b3889576f255cc1785af074d9ceb3a32e0df4b"></a>

## network_pbr.network_pbr_rules.protocol_port_range — network_pbr.network_pbr_rules.protocol_port_range / bd1b23c153f0 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- network_pbr.network_pbr_rules.protocol_port_range

<a id="canonical-33208f1a676cbe633844719daa4b20691e12dfe2ff0f1973bd16d9910b7e5dd4"></a>

Type: `"single"`. Computed.

Protocol and Port. Protocol and Port ranges.

Upstream description:

Protocol and Port ranges.

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

<a id="canonical-61a03eb20eb9bf6e88985bdc5c01613ee06149a8282edb80b2dbabb314916632"></a>

## Direct properties — network_pbr.network_pbr_rules.protocol_port_range / bd1b23c153f0 / 3

<a id="canonical-b8d50dbd76780d32b11d12e6224db04a72c847b0ba838256b93737cc5c8906e7"></a>

<a id="canonical-11c62b82ddad1c65f5e1d972abac0ff4656259bf40958a34f64b39157ad0ffe0"></a>

## port_ranges property — network_pbr.network_pbr_rules.protocol_port_range / bd1b23c153f0 / 4

Type: `["list", "string"]`. Computed.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-c06e8b388a0ac5f95f0d895efc1893a61ccad968cc09777eb3a1db3ea2048791"></a>

<a id="canonical-a3dceedb831965ddea5dafccb840fd9f26dc8f47ba04dc70d9d0a39b1e34ddfc"></a>

## protocol property — network_pbr.network_pbr_rules.protocol_port_range / bd1b23c153f0 / 5

Type: `"string"`. Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-59c3d4e30a75e0eb7915bc17e721d137511204552dfaed9bf15e5b266ff2681a"></a>

## Next pages — network_pbr.network_pbr_rules.protocol_port_range / bd1b23c153f0 / 6

- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-16ff46784595eeb798611194d114aa63a92f53ae9a9a4dc34a11f33cb8ec0901)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-2e666d0078404f5978dae6c8a24c2acec20583e056d394420b3608a2f79fe274"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e16eee024a1ef13684a5c644f6a76ebf8ba2d79f575072719cc22a1e79dda83f"></a>

## network_pbr.prefix_list — network_pbr.prefix_list / 4ab14243d6f8 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- network_pbr.prefix_list

<a id="canonical-8d635098eb8ea341d7ba1def8781b9982c5a3db07ff148726ea04de622789628"></a>

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

<a id="canonical-8f557bcfd031054a92407a78fadc37028befbbcb01c01e7975cf8ce848ecdc05"></a>

## Direct properties — network_pbr.prefix_list / 4ab14243d6f8 / 3

<a id="canonical-9f2909e540d7324505f3486fe306bc206a7fd50359896e3009632097108ae4e2"></a>

<a id="canonical-c910f349125059daf8576ddbe9838369c8ffced73c7deff24ad1452d0a5b13c3"></a>

## prefixes property — network_pbr.prefix_list / 4ab14243d6f8 / 4

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

<a id="canonical-d821d86642b4a97ff3eefa7b83bce9552922deef4d379870dc8152e4e49fbb61"></a>

## Next pages — network_pbr.prefix_list / 4ab14243d6f8 / 5

- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-b46f77fcf0f5c5a0cac6b82f1f42cd7b35cd71305145a9280f8d0046a71b2405)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
