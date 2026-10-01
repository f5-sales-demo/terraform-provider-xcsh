---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d51996acf56bb8d473a9e499efb0f91dfa0ea629a565c073997785c33ad691f"></a>

## Property reference — Property reference / bca274cfeba0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- Property reference

<a id="canonical-451a0dc4f44f1f39745aa3ceed67b9c6b24e3c9317e0dbb83eab446d2b7504e1"></a>

## Direct properties — Property reference / bca274cfeba0 / 3

- [active_service_policies](resources--cdn_loadbalancer--reference--group-003.md#canonical-dc4a01731534c5c0943f21101a5334d07234e019abe85d9e5d146dbd9704d4fc): complete subsection reference.

<a id="canonical-31ce353d5a0866d851f9139206883c4ecae645ae2b82ed49bbf086eae90e4825"></a>

<a id="canonical-7b8c05d0f49011e30123f467b0d2d371a87c248add9c5ddbf8e4f62d6cd4a98f"></a>

## annotations property — Property reference / bca274cfeba0 / 4

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

- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68): complete subsection reference.

- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936): complete subsection reference.

- [app_firewall](resources--cdn_loadbalancer--reference--group-006.md#canonical-3897308533814e8cc81ad5c3f6a6b8078cda368c7632d9c91a66707448652ed8): complete subsection reference.

- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-118168ffb743afdb5fc4125b1eb1ed93aa625aae7cdbf8c967248821b6a03d2e): complete subsection reference.

- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110): complete subsection reference.

- [captcha_challenge](resources--cdn_loadbalancer--reference--group-009.md#canonical-1c2555638ed3b0821ac370fa14c5cddaae5615d0e10f75df287474926bd68764): complete subsection reference.

- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6): complete subsection reference.

- [cors_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-61274ef5700bf47a6640898b8f6b648c3daecd806a5fa82ea3a5c5d02fec2f66): complete subsection reference.

- [csrf_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-407303fdc296fc1cb771072072f9f38707eff9e1eb19854190f5f811b1803371): complete subsection reference.

- [custom_cache_rule](resources--cdn_loadbalancer--reference--group-009.md#canonical-a4f8e5202487add45d34bd5e279421532fa89e3abc164a114fdf93433b0f2134): complete subsection reference.

- [data_guard_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1dccf7057060794e294ec86ddd031213e249d14c153ed77e26fdb890afbba54f): complete subsection reference.

- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-5746a5502dfc8e0895cb2a3ef2a31f60a71459706a1427f5bec63040fb069f07): complete subsection reference.

- [default_cache_action](resources--cdn_loadbalancer--reference--group-010.md#canonical-b025aa98d3860c4c9152491bc9a93a336955de2ba2530bd0bc7f447722ac7db7): complete subsection reference.

- [default_sensitive_data_policy](resources--cdn_loadbalancer--reference--group-010.md#canonical-b4c6982c2159d86d5d460cc99c33456ce32b20be8c4e006cef5f974d2fbaf7bf): complete subsection reference.

<a id="canonical-68446349b9a28ab08120f272e164e421f47cd1a41c9d71d0df67745bbb01198f"></a>

<a id="canonical-a3e754ef4dd87ccb86f320af89a8fb2d32ba57feae17d3fd29e44b1bcee5d45d"></a>

## description property — Property reference / bca274cfeba0 / 5

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

<a id="canonical-44cc05fbd0a85489ee1b24fa4d1a5e4365ec7b46bf95bd2210af7b7ee71a3253"></a>

<a id="canonical-f5ef690baeb98d8892dfdb2352e8446217cb3a72f249ea81982370bf95d94ecf"></a>

## disable property — Property reference / bca274cfeba0 / 6

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

- [disable_api_definition](resources--cdn_loadbalancer--reference--group-010.md#canonical-f6b52b1f02e5223d4d1e8c0b11a436aa90f3703cbdf600691fd6fc74df939284): complete subsection reference.

- [disable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-9c91e5fa1b5de026eca2b449600d794d171a250675f3db407f4024485d04df43): complete subsection reference.

- [disable_client_side_defense](resources--cdn_loadbalancer--reference--group-010.md#canonical-3ad968aeeec75f14307d18b7058a7031d761b4b11eacc418a22217bc9e7ca4d8): complete subsection reference.

- [disable_ip_reputation](resources--cdn_loadbalancer--reference--group-010.md#canonical-4d062999b89e2ec7b48c891eea34217d6db7453eb76eb5c1ca992f27c0e6c039): complete subsection reference.

- [disable_malicious_user_detection](resources--cdn_loadbalancer--reference--group-010.md#canonical-fbcfe12bdcce7e716f099f50148c97b40fcb7c7bb76a476ee7b227537a527741): complete subsection reference.

- [disable_rate_limit](resources--cdn_loadbalancer--reference--group-010.md#canonical-919699123b08df44a4a2d6060b273b4e4eccbe2b03bfe046d34d4af96e8f2dc9): complete subsection reference.

- [disable_threat_mesh](resources--cdn_loadbalancer--reference--group-010.md#canonical-30068124851c4c974db1422b8003b7e1bd2902bf0a902448ee113ea0058278ed): complete subsection reference.

- [disable_waf](resources--cdn_loadbalancer--reference--group-010.md#canonical-582f1a4704469fb691c6c973eb67b1da8f6e5db5a115a88613597e82c377b485): complete subsection reference.

<a id="canonical-8b32fc1e7b79a48c4e09936cdf90361316def9f5b5328b27317112460b612c9e"></a>

<a id="canonical-62bcce3be34f4cb2d26e10f2bfb2e588dab2d6729dd94ce41bda76d8cb2392a0"></a>

## domains property — Property reference / bca274cfeba0 / 7

Type: `["list", "string"]`. Required.

List of fully qualified domain names. The CDN Distribution will be setup for these FQDN name(s).
\[This can be a domain or a sub-domain\].

Upstream description:

A list of fully qualified domain names. The CDN Distribution will be setup for these FQDN name(s).
\[This can be a domain or a sub-domain\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
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
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "[\\\\.]+[A-Za-z]+",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "[\\\\.]+[A-Za-z]+",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-b3410bb7f2b0c9f28d0bb252433589a6024286ed4d766c0627313c10a630f1bb): complete subsection reference.

- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-be72efa8da41d7b196abe830ae2dc8afde3c2becf8166936a3708d45c8cf0a73): complete subsection reference.

- [enable_ip_reputation](resources--cdn_loadbalancer--reference--group-010.md#canonical-de2142fc296e6a30b369e4c9b13a4ed016e7ed01af30d672900c574d3983bcda): complete subsection reference.

- [enable_malicious_user_detection](resources--cdn_loadbalancer--reference--group-010.md#canonical-739cd48cdacabeb3c20e33071d28210102da0caafd6e525f2c7bc6a4f191a530): complete subsection reference.

- [enable_threat_mesh](resources--cdn_loadbalancer--reference--group-010.md#canonical-f1b1fa9bfde00138f0ebbee2539ef4ce9300d5d0a630096b677d0a200b99ba0c): complete subsection reference.

- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1dd230b542892f64d8052e5db7f27223d60233a84fe969e8ce6e411444895cab): complete subsection reference.

- [http](resources--cdn_loadbalancer--reference--group-010.md#canonical-f81a9f81cb120d5f06895975b67987a4ceae629b66a48aeaa3ba2a2ea3bb414d): complete subsection reference.

- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595): complete subsection reference.

- [https_auto_cert](resources--cdn_loadbalancer--reference--group-011.md#canonical-bfe4be2d539814187ee6209f6c1e2551e3a3a993ebb2b6595877043d81965a7b): complete subsection reference.

<a id="canonical-a3261fe5771a0f859f7a90e57e430808c65fdcc5be1f1073bcdd55380ea00e6b"></a>

<a id="canonical-377817e6afc9ef297b326873cd898deb26f1d6f964ca9dbf32ab24fe43af4ddf"></a>

## id property — Property reference / bca274cfeba0 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

- [js_challenge](resources--cdn_loadbalancer--reference--group-011.md#canonical-076789b43998a8bee94a64569f821cf1ab5cde4dac942c574ba859bdb49943b9): complete subsection reference.

- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d): complete subsection reference.

- [l7_ddos_action_block](resources--cdn_loadbalancer--reference--group-012.md#canonical-722c0532d8572d916e96bd684dfab040b2fc03c7412dd033ee3fd232b8eca199): complete subsection reference.

- [l7_ddos_action_default](resources--cdn_loadbalancer--reference--group-012.md#canonical-acf231ce5ffcd5d0095980e96c8d1e272272f831c0f9d1a7e3ea58cd5daaf982): complete subsection reference.

- [l7_ddos_action_js_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-ec16815ce3b83fcc24968fa670a2b071b426e9de1613736dc0f5ac839f891325): complete subsection reference.

<a id="canonical-a400b68631e55abe11ed260d087b446b3617e3c0d09bfc58973e00847f6c5538"></a>

<a id="canonical-2d710671c6b779bead70af487c7ac0b60d4685d14d9d52a3ba6bd624b6ab5637"></a>

## labels property — Property reference / bca274cfeba0 / 9

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

<a id="canonical-a182084fa9f0100ef39f6b9b5be0bbf1795761b674537ac9c71a9a4c292cf07f"></a>

<a id="canonical-4c041d1ec132c2e073fb82a789f476b818d20fdd2bff01e047b7b6dceadc4ec6"></a>

## name property — Property reference / bca274cfeba0 / 10

Type: `"string"`. Required.

Name of the CDN Load Balancer. Must be unique within the namespace.

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

<a id="canonical-bf8518aa29584f676125977f25d67ed4df1e8af7c41e7790ec2d4d81bdf1a5d9"></a>

<a id="canonical-1673b29c975157652067e67e543e571d9dc8ded39f58634c8035c56273ad9431"></a>

## namespace property — Property reference / bca274cfeba0 / 11

Type: `"string"`. Required.

Namespace where the CDN Load Balancer is created.

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

- [no_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-f6005363adc29c9893da440ca4e14b77a49e5969955121dcce8465038b52940d): complete subsection reference.

- [no_service_policies](resources--cdn_loadbalancer--reference--group-012.md#canonical-db5bd1441ea30ff501c1ef1d5ef4672bf5bb2236a075287daff00db31ea5a53b): complete subsection reference.

- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-603a6ee1a88486640208f36e01e4e00f68db830979076e4ad137090fd9373676): complete subsection reference.

- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-6237c1513ed4eb47b747c88f4987c86621425202c792c064ab4bfdfacbad388a): complete subsection reference.

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f): complete subsection reference.

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a): complete subsection reference.

- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957): complete subsection reference.

- [sensitive_data_policy](resources--cdn_loadbalancer--reference--group-014.md#canonical-d1c756684a9b6aed13b1e1004dcd60244387d594a41653bbd78e71f0554cae31): complete subsection reference.

- [service_policies_from_namespace](resources--cdn_loadbalancer--reference--group-014.md#canonical-2e87a80d3c690f75e98f2c8ac9c5a16eec12e21cfe3135d22b5ec68445fbc25f): complete subsection reference.

- [slow_ddos_mitigation](resources--cdn_loadbalancer--reference--group-014.md#canonical-6f0c88a559028359020b9885e0a124c2f6941cc75507bfe4ee947054f22f2ed6): complete subsection reference.

- [system_default_timeouts](resources--cdn_loadbalancer--reference--group-014.md#canonical-72cb0a186e7fd30aa2b9bf4141b26dff8238c3cd30698e990babc911234f1e00): complete subsection reference.

- [timeouts](resources--cdn_loadbalancer--reference--group-014.md#canonical-58e15a9529dae4f990093f1f19e485d5d6a8e14c3eee8aa8361abfecf86ca89f): complete subsection reference.

- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-fa607075ce0e765101a422c4bb6d254a483e196f2d69d9529aeeeb3395be1872): complete subsection reference.

- [user_id_client_ip](resources--cdn_loadbalancer--reference--group-014.md#canonical-f2820c8e4f1688393af0fa5b74f4a574dcfd9a57773098a177d5ad9db86f162d): complete subsection reference.

- [user_identification](resources--cdn_loadbalancer--reference--group-014.md#canonical-56b11c6aea01027457a9e3e86571f28a84dccc0c36e4c80325f3f4976c5aec1e): complete subsection reference.

- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-49d1159484a5c47de40732ac1f168236175c7884528ef71566d61749b0a9427a): complete subsection reference.
