---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1fd6db45299eb51cf36d21446b25b6afe74df8799646350fc9b5b84e44b2a57"></a>

## Property reference — Property reference / e36a9d4438cd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- Property reference

<a id="canonical-b6ffe58a1262d07ab7e15e715aadc4a93665d07bdef46d7d70710efa57cb1699"></a>

## Direct properties — Property reference / e36a9d4438cd / 3

- [active_service_policies](resources--http_loadbalancer--reference--group-004.md#canonical-45c816f1b70a39912fe4d4cd1e4922a2d552905b7502d2c7b4e5d846d8dbba3e): complete subsection reference.

<a id="canonical-9cb1054cf28ac28b57b3877cb2ec5f9b6905a4afc113bb35d616ed235e1ddbb3"></a>

<a id="canonical-0a159d28fdf185ce747f2dcf7114c1f5a1ef99a2d1ac92eac8056906efc35295"></a>

## add_location property — Property reference / e36a9d4438cd / 4

Type: `"bool"`. Optional, Computed.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses. This configuration is ignored on CE sites. Defaults to \`false\`. Server applies
default when omitted.

Upstream description:

X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; in responses.
This configuration is ignored on CE sites.

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

- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-889ffa3e60dd17feee86ae3bcee07ee8938f2abb4e13aedfc088e05b482b2185): complete subsection reference.

- [advertise_dualstack_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-2b6b25016f9c9da773dad6bae4734bfe7b891c3106abbaf0558aef1d29c4d71c): complete subsection reference.

- [advertise_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-8e3272b86153911cab6a4209bcdf57278316716883e499f3b33c289070b441b5): complete subsection reference.

- [advertise_on_public_default_vip](resources--http_loadbalancer--reference--group-004.md#canonical-da3a88c278b67c687361856cade6be473dd54723d187c266d2f0210405903bda): complete subsection reference.

- [advertise_v6_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-9fa1ab7d46fd1cd4bbb0f4a75e7c47b560ec8fae253abe3988bc2e1df850ea24): complete subsection reference.

<a id="canonical-10aaef5202242e8f91a6a3a11363f04f92482996d6685e0c63471334b4c23983"></a>

<a id="canonical-3322407d5dae6cb7397527e7bead08aab9669ad45cf6ead207a643b36e38fb75"></a>

## annotations property — Property reference / e36a9d4438cd / 5

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

- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c): complete subsection reference.

- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592): complete subsection reference.

- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10): complete subsection reference.

- [api_testing](resources--http_loadbalancer--reference--group-010.md#canonical-699d08bedad7c6a3db855143b0dfa0f5040e8efa4d08cd58dccbf39afde96a71): complete subsection reference.

- [app_firewall](resources--http_loadbalancer--reference--group-010.md#canonical-bff9954daed615e67cb6a3932ebf85a6050831ca3225573c0776c74dd5eb1f8f): complete subsection reference.

- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-3a2e92765af87c3d276d5f20a62319b47649d669158cdc32c4fb394dd3035716): complete subsection reference.

- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc): complete subsection reference.

- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-012.md#canonical-c97868542280c86c9322eccdb6092ef736f7e9c0f70d01a8ccbbbfeabcdc263d): complete subsection reference.

- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-d20150bcbc1a30a287cae322ba14c8cf07c2866aab4bca3a45ecca53fbc870fd): complete subsection reference.

- [captcha_challenge](resources--http_loadbalancer--reference--group-014.md#canonical-fd05a1137cab68c6922bcd30d9888aeee26a67af28a835d3b19f4d7b6e67bbb2): complete subsection reference.

- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa): complete subsection reference.

- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95): complete subsection reference.

- [cors_policy](resources--http_loadbalancer--reference--group-014.md#canonical-98b1bc48ab39bfd42acdddd7722ced6b2d8018703544313e54793d9a49e8306d): complete subsection reference.

- [csrf_policy](resources--http_loadbalancer--reference--group-014.md#canonical-c9a870f8ec486129a3bfb92e157af90b7ed2f407741b40a6ff596b41269bd975): complete subsection reference.

- [data_guard_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3442b1c895a9f57fa1bf6b5e124c769e15a335cd5a6ebc494c98b2dc3d1f9ca9): complete subsection reference.

- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-015.md#canonical-d727335d761aa12d3d7e0cdefe2c946a177e7d8479702d7506c618125c21fc92): complete subsection reference.

- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345): complete subsection reference.

- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-2e2c6f117545e8507f9a557e78ac7c1395c119a6fefa3b39e4e67a70cc292675): complete subsection reference.

- [default_route_pools](resources--http_loadbalancer--reference--group-017.md#canonical-2b8c60fa9dceb82aef881e64259e895cadc5174aabddb1c7bc55794ad8d2c41f): complete subsection reference.

- [default_sensitive_data_policy](resources--http_loadbalancer--reference--group-017.md#canonical-865cc01ec6425b69dc0f336f2b1df4508cba7e06e1863d551ccd67564a9f8773): complete subsection reference.

<a id="canonical-946220f7bfd68cabd201b52e1d70f728bec64098d8bded6af54439372c6a5e79"></a>

<a id="canonical-c9c85fb9c5348d2b70b1ae25604cba75fe3b17e6aa7ec54578d20d18f54bd64b"></a>

## description property — Property reference / e36a9d4438cd / 6

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

<a id="canonical-f2cab1ec389f070ee508b3e799e092d3b1a9a6408b8676bc92947052a4427adc"></a>

<a id="canonical-9f0be39dd7145a1d390450447917f193faf861d8ff5e18c11fbb7aeb3cc0c0d8"></a>

## disable property — Property reference / e36a9d4438cd / 7

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

- [disable_api_definition](resources--http_loadbalancer--reference--group-017.md#canonical-1b55f3059b0d40854d39c4144a9207fcf3f11031231709c548cd57f59eedb026): complete subsection reference.

- [disable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-2e06259edaa0293b9b5a461b956cd0a71e61ff8a95f34885a81c11c6239f00ff): complete subsection reference.

- [disable_api_testing](resources--http_loadbalancer--reference--group-017.md#canonical-f789378c079cb5e3adaa3102def219654d096f04b8c8a2cb49e577b8a9d534a7): complete subsection reference.

- [disable_bot_defense](resources--http_loadbalancer--reference--group-017.md#canonical-db0f00840fedd3d4d10fca6b6e6a38c90e6d3c86f97d953ed4754b50821630fb): complete subsection reference.

- [disable_caching](resources--http_loadbalancer--reference--group-017.md#canonical-b944ef368907204f33191c57e912db818db4f0d91232de5993ce8c168e2f8289): complete subsection reference.

- [disable_client_side_defense](resources--http_loadbalancer--reference--group-017.md#canonical-9ae7b6d81df251cf2b47e0945f1a03143010e544df39ca300e07b415f04488f0): complete subsection reference.

- [disable_ip_reputation](resources--http_loadbalancer--reference--group-017.md#canonical-626bc9a9e1f5a8dbcff2cfc47f51f44452f307d56b7e4446911d766b3cdaaa42): complete subsection reference.

- [disable_malicious_user_detection](resources--http_loadbalancer--reference--group-017.md#canonical-c6b6bcee92fa30fbd299d474868a4374a98620ff5e7a8c91390225b379dfd71a): complete subsection reference.

- [disable_malware_protection](resources--http_loadbalancer--reference--group-017.md#canonical-73fd4dd9ed35740311b05e7c058fa47c122713897fcd26a1db5a9492c44f1311): complete subsection reference.

- [disable_rate_limit](resources--http_loadbalancer--reference--group-017.md#canonical-c48e1d3b25523022258a53db20abda4b84439f175704d24090277c6e2ae78317): complete subsection reference.

- [disable_threat_mesh](resources--http_loadbalancer--reference--group-017.md#canonical-a144aa0b53927f9806853d0586a2b3c5b599fac28431d3cf714604a41e03584b): complete subsection reference.

- [disable_trust_client_ip_headers](resources--http_loadbalancer--reference--group-017.md#canonical-8472f96112920b65bd27a1c5e159bbf163ddeb4a5ca3802a1d9872ad3c0607ab): complete subsection reference.

- [disable_waf](resources--http_loadbalancer--reference--group-017.md#canonical-2d3dceada3af713c60e3a5c97f2df4fffd6ce15af8979fc63fd04464363b4849): complete subsection reference.

- [do_not_advertise](resources--http_loadbalancer--reference--group-017.md#canonical-b29f363d3a20a25f933614c6079ff33efdf1edf8965dfa2ff873b90ad86f557d): complete subsection reference.

<a id="canonical-73b74109d9a83a993f67dcc5ce9dd42df91233391d69ba08cd52b7577f0b8880"></a>

<a id="canonical-d969beda9d433423251c5e70837c74a6e86806a5187452fc6f52aa452e117716"></a>

## domains property — Property reference / e36a9d4438cd / 8

Type: `["list", "string"]`. Required.

List of Domains (host/authority header) that will be matched to load balancer. Supported Domains and
search order: 1. Exact Domain names: www&#46;example.com. 2.

Upstream description:

A list of Domains (host/authority header) that will be matched to load balancer.

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if the Loadbalancer type is HTTPS. Domains also indicate the
list of names for which DNS resolution will be automatically resolved to IP addresses by the system.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739): complete subsection reference.

- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-402b08ef1e68bc5d4cd755fa7d308044f7c196ea49abd373321f605d815aebee): complete subsection reference.

- [enable_ip_reputation](resources--http_loadbalancer--reference--group-018.md#canonical-ad9abe471b9e2c134289301619d5e2e343addf7ac9c75b2ef21b2ab519e4a2bb): complete subsection reference.

- [enable_malicious_user_detection](resources--http_loadbalancer--reference--group-018.md#canonical-92c71e787d6f8b942b2de2f570db9e12d25adfadafb90be4dab1428b72d246e8): complete subsection reference.

- [enable_threat_mesh](resources--http_loadbalancer--reference--group-018.md#canonical-d5e845d5425c1311e52b352f9fa4ea398545c1b66ec415406320486cbd25daef): complete subsection reference.

- [enable_trust_client_ip_headers](resources--http_loadbalancer--reference--group-018.md#canonical-8b81ab46828f74c3866c9e65d4068c4c9adac39de02d860dc0b15ed239a98e8f): complete subsection reference.

- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-13765b09fae8733d02c20565cbeec906e4a91070c52e2eac1b36668592698cd6): complete subsection reference.

- [http](resources--http_loadbalancer--reference--group-018.md#canonical-e1f14d7c9b8e8fb184d02bb2a0fe678d962de127c1841a842aff87500721227b): complete subsection reference.

- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012): complete subsection reference.

- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8): complete subsection reference.

<a id="canonical-e0a070fe83beb7497f3a08fc597b077634f7cb92a662b6b81604bee4308d821b"></a>

<a id="canonical-32ebd3866ffce66ccfc61cdb2f4ffa1728235fd1556487c6855b3efb060e968b"></a>

## id property — Property reference / e36a9d4438cd / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

- [js_challenge](resources--http_loadbalancer--reference--group-019.md#canonical-1a9d02dfd2c57fd3a9e31ecff2bb82960607e7bd4a61e0a621dd4cbbf874f729): complete subsection reference.

- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff): complete subsection reference.

- [l7_ddos_action_block](resources--http_loadbalancer--reference--group-020.md#canonical-7851e2010afcdabe3e06b86c00ca2119a78dc076e9788767f1d31b1ba8a5e3f0): complete subsection reference.

- [l7_ddos_action_default](resources--http_loadbalancer--reference--group-020.md#canonical-dacd16fe43e99c16d648b347417ec881594427f7192ce07760ccb0a434933ebf): complete subsection reference.

- [l7_ddos_action_js_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-5c2a0de006a63bfb8d7a46ca1a9ddaeb0faf0bc4cc2c75baf932466b1c4b9a8d): complete subsection reference.

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-a5261dd749d091cf720a2544f603c93cba5bd20eafe6068a8958d6f4e7b2ac25): complete subsection reference.

<a id="canonical-c81013319e241b63ffa38e888de2b32d143ec1a5c1dd5c5051fcb8352777304f"></a>

<a id="canonical-79dbd3d4a01ee391918597e98f496dfc7f9d399154f032cb8c9416eed597d52a"></a>

## labels property — Property reference / e36a9d4438cd / 10

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

- [least_active](resources--http_loadbalancer--reference--group-020.md#canonical-d01e18784c2ea24048b46a4e90708b3edcd19fab94a6a493e19cf7822cb11fb8): complete subsection reference.

- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-82051da6fa8592d3703f6f8ebe7726692a2cc0a847b7f51741c8239f1a7a7fba): complete subsection reference.

- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee): complete subsection reference.

- [multi_lb_app](resources--http_loadbalancer--reference--group-021.md#canonical-0ae28b3b2b8df4f8ce2c8e04c0176280fe9c9128fb14f292dee21e7534b1c785): complete subsection reference.

<a id="canonical-0ae0a02b0490ea85f6248345d30e6832cbff1908feadbedefb7e34c9657b32ee"></a>

<a id="canonical-969130477f9672416057070d3636e4529c8e076d751a2b9000bf042dc24b4d98"></a>

## name property — Property reference / e36a9d4438cd / 11

Type: `"string"`. Required.

Name of the HTTP Load Balancer. Must be unique within the namespace.

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

<a id="canonical-0ef0f979913fddb8df4ac380214ef1a460f3081857292c39366faaeb05f45670"></a>

<a id="canonical-68c95b7e7f8001c178508df05b4796595b76552fca72ed2a3c6ef239b7a81320"></a>

## namespace property — Property reference / e36a9d4438cd / 12

Type: `"string"`. Required.

Namespace where the HTTP Load Balancer is created.

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

- [no_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-91668e9b3365ea876e1816b4b40472e96059f8303e3a32daf41ba574857d4efe): complete subsection reference.

- [no_service_policies](resources--http_loadbalancer--reference--group-021.md#canonical-67a5b67e1c9362160c1235a278e79e0e6a99c2fc8595e33cad1c274c07dc323d): complete subsection reference.

- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670): complete subsection reference.

- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715): complete subsection reference.

- [protected_cookies](resources--http_loadbalancer--reference--group-022.md#canonical-31e37d7fd8ec4dc8fe24670a3a9a6042ac21493ba23b1d0f7d8652cf6147642f): complete subsection reference.

- [random](resources--http_loadbalancer--reference--group-023.md#canonical-766379fb930c44ad5e5e2c04133a6bc22837ba32ed08ce5c0423109576169e15): complete subsection reference.

- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-7cb6dc89026b3071da1cc7ef34ceaf2acccfcbfeb82b7751ead0e13de0ea160b): complete subsection reference.

- [ring_hash](resources--http_loadbalancer--reference--group-023.md#canonical-f3ff5d1ad828b5311f35bfb1f7876549c6ba1748e43ab9dbeef4e660faa6143c): complete subsection reference.

- [round_robin](resources--http_loadbalancer--reference--group-023.md#canonical-d88720185138bfbe63c4c4e5b159d2d0fbff7267924d86c7e5b8c5c2744a23f9): complete subsection reference.

- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50): complete subsection reference.

- [sensitive_data_disclosure_rules](resources--http_loadbalancer--reference--group-026.md#canonical-c47009539eef3bdcdb8d2b7bb2e345b80b42d65bebf56615f0a60ba75fe78afe): complete subsection reference.

- [sensitive_data_policy](resources--http_loadbalancer--reference--group-026.md#canonical-f0e4b12924ffd6a5474ec9ccc1e5f3bb10984f997a330e17cd29b09c63f46f08): complete subsection reference.

- [service_policies_from_namespace](resources--http_loadbalancer--reference--group-026.md#canonical-f8e76b68a32f7686ffecdddc1f58ed7947137696e500423d54f0161e3b096a55): complete subsection reference.

- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63): complete subsection reference.

- [slow_ddos_mitigation](resources--http_loadbalancer--reference--group-026.md#canonical-ffd5840bac3442e8871f0e6732a39bf72a5cf7ca9d57813bb3e3f7dac2c7c03a): complete subsection reference.

- [source_ip_stickiness](resources--http_loadbalancer--reference--group-026.md#canonical-1855ab42bff48be1c839dbc02e86ba75469ba9ba572e6e12d5a69afa2ba6d8c0): complete subsection reference.

- [system_default_timeouts](resources--http_loadbalancer--reference--group-026.md#canonical-d544b47e171fe74c0dd1c3e1bb390b65307469826b276b138f3f4d3d59ebc155): complete subsection reference.

- [timeouts](resources--http_loadbalancer--reference--group-026.md#canonical-0d031c5e02b9b01edc14906dd7d479341ac318e527a84ba537a24a689885ff3a): complete subsection reference.

- [trusted_clients](resources--http_loadbalancer--reference--group-026.md#canonical-e0af00365f8f0ebfccab32942f6b2726b0806a55ec8ea4b3b523dd7cebb7e4c9): complete subsection reference.

- [user_id_client_ip](resources--http_loadbalancer--reference--group-026.md#canonical-fe1f51e76f8a4491a2c7ef8f9cc82d6a6df68859896b3b9775334e6f613cac00): complete subsection reference.

- [user_identification](resources--http_loadbalancer--reference--group-026.md#canonical-0e0670b13d60febf0c438e4f8a73b5794f67bb35ecad84e1001cef3948fb2a2b): complete subsection reference.

- [waf_exclusion](resources--http_loadbalancer--reference--group-026.md#canonical-86635926dff9a406ce1b9e922504e4198e3c10c37ace3a589dca358b0134d051): complete subsection reference.
