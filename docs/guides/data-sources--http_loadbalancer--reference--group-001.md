---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7d3e1abf19ec7fc0269517eef0f35f15163f884fd33e9366927c2d231a554ad"></a>

## Property reference — Property reference / 15e40702d135 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- Property reference

<a id="canonical-d6821fae1d2c75f35c84dbcc0b281c5533f2eaa8b361e8bb188089804f2eaeed"></a>

## Direct properties — Property reference / 15e40702d135 / 3

- [active_service_policies](data-sources--http_loadbalancer--reference--group-004.md#canonical-4c31a152608d05332ae3cdc066fe5046c57a37f8e39b3a3aa7e6e60b8f1fbeec): complete subsection reference.

<a id="canonical-f928edd837b6a737a6ac00de96dfec538709db48c89bad5e02ca824f24ed555a"></a>

<a id="canonical-efd565a6d75877a3aa6deae5e7d8d56242cd0a75cd32ae5c96a26b9cfc07da81"></a>

## add_location property — Property reference / 15e40702d135 / 4

Type: `"bool"`. Computed.

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

- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-b912ea97a62c576c3a2059c88cbd7db3a9fd21a1411cfe52c7a58eb370683c80): complete subsection reference.

- [advertise_dualstack_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-ac57a9fc5fde577900d5ede85b4bf3856dabbec4129150043e0a51657476d525): complete subsection reference.

- [advertise_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-94522cc2e6d2d47a177c43de6d04c236d1248ca9e55579fe225e014227bd5f2c): complete subsection reference.

- [advertise_on_public_default_vip](data-sources--http_loadbalancer--reference--group-004.md#canonical-0b24aa92bc0857dbfc9171287aaa35018704e6c4a5f87af95a3320c4c54bb08f): complete subsection reference.

- [advertise_v6_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-ef441ea1c447c1246c6b24e448b776f55b2a43c6019779c89b04b9c61eb8d9b4): complete subsection reference.

<a id="canonical-6a867d504f16c9ad0331b14d0581bc8c9ea223f43477c523f1819abac0df3ef6"></a>

<a id="canonical-b01aae2734a9bb5f6ca5517af2f1fb274541a3984a3ed53d054ca37d3bb7e6dc"></a>

## annotations property — Property reference / 15e40702d135 / 5

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

- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318): complete subsection reference.

- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5): complete subsection reference.

- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73): complete subsection reference.

- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76): complete subsection reference.

- [app_firewall](data-sources--http_loadbalancer--reference--group-010.md#canonical-a2c8a93200e97df3ae4c03b87e5920ac3ccf3a34ff325cacbecfe6c9a5e575f2): complete subsection reference.

- [blocked_clients](data-sources--http_loadbalancer--reference--group-010.md#canonical-e46710ba001bf8ddabcfccfbb5094215ce20da3a35b0ea48470159cc0e080303): complete subsection reference.

- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90): complete subsection reference.

- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d): complete subsection reference.

- [caching_policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-a6fb2379e952189b6bcc8f8566767a9972f615481793db4515c819834aa1c952): complete subsection reference.

- [captcha_challenge](data-sources--http_loadbalancer--reference--group-013.md#canonical-bf8524ae85110951f9bb08b81c5e0e46ea159fc5e7d34366a39bcf3de81f7c7d): complete subsection reference.

- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f): complete subsection reference.

- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-014.md#canonical-acac9167259cfc120ba2f9c96337e59fc39aa9b471956e6aec6e6cd51b54cebb): complete subsection reference.

- [cors_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-4ffd9159bc5a97623e1ac711bca5b5865d2f43de09a0b5b7552e4670e7a446be): complete subsection reference.

- [csrf_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-f1e6471143974860830fd021444acd48e74e2b1b6f98ec945d6dd4b24d617417): complete subsection reference.

- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-d2ef521c5a4c109fc00a5a94ff27b673fdc44185c8a89cd3ede83061de1f3343): complete subsection reference.

- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-31832e4dcb109b76102f24b4362682d28eb3957fb05c9480436f8c3c2c1d0201): complete subsection reference.

- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d): complete subsection reference.

- [default_pool_list](data-sources--http_loadbalancer--reference--group-016.md#canonical-98278c5aa3948605fd01fc43b6fea08aee92fa3c57a3fbe8bb93568f2c8dce49): complete subsection reference.

- [default_route_pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-9be3fca6f6542504cece244b3240752dde0016a01ffb4d8c39941639d5bba0f6): complete subsection reference.

- [default_sensitive_data_policy](data-sources--http_loadbalancer--reference--group-017.md#canonical-15b893a4a3c0f97fe3a61685768548528500e3486c5d0f3aec45a9409148202b): complete subsection reference.

<a id="canonical-1efe0832979017442fb702d17ca3495ed644ecaaa835410375879fd7306b3d8b"></a>

<a id="canonical-b10e7dd4cce086c766828e547491dcb1610f128b5b1b72590d695d5ba3872af2"></a>

## description property — Property reference / 15e40702d135 / 6

Type: `"string"`. Computed.

Description of the HTTPLoadBalancer.

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

- [disable_api_definition](data-sources--http_loadbalancer--reference--group-017.md#canonical-d3e1f6a451e30b0e93bc8549fd2dbeb90aafb192811d671ec70ac7e5beac5f8a): complete subsection reference.

- [disable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-72f2d8912b0e440553040e939c1c7c570deff6b547de214884cc46b0b9e5aba1): complete subsection reference.

- [disable_api_testing](data-sources--http_loadbalancer--reference--group-017.md#canonical-ff803d8b6937c7510caf2b002088badbec57751bcc7d1f09b82ea1227edec199): complete subsection reference.

- [disable_bot_defense](data-sources--http_loadbalancer--reference--group-017.md#canonical-025c607a7db6cf103750aee084ffecd0bb52e176f09763341fbda5b42f40b997): complete subsection reference.

- [disable_caching](data-sources--http_loadbalancer--reference--group-017.md#canonical-02c9534d5c7cf318702bcb048091f816fa8ecfdee8e128dd3512d341989d9953): complete subsection reference.

- [disable_client_side_defense](data-sources--http_loadbalancer--reference--group-017.md#canonical-9bf9f7383d97a540bcb94e62bc865711c55889bb695677d1cec4d59eeb47ac0d): complete subsection reference.

- [disable_ip_reputation](data-sources--http_loadbalancer--reference--group-017.md#canonical-1674ba81c74af0eb269b77b9e757b1b7d36b13b46456f35ed9410bc40de849ae): complete subsection reference.

- [disable_malicious_user_detection](data-sources--http_loadbalancer--reference--group-017.md#canonical-7cedce4a45d8ac191cec74f6233758f2984112f3a55123fdb66eb1ccc4b5ffbb): complete subsection reference.

- [disable_malware_protection](data-sources--http_loadbalancer--reference--group-017.md#canonical-b57e296971fc4052df817977dfea27b735768e109893ad2b7100587e0126c8cf): complete subsection reference.

- [disable_rate_limit](data-sources--http_loadbalancer--reference--group-017.md#canonical-28e55b9e2b6b6493d430059ab19205160e6ac3bb8c4402689d3b95c786c22b0d): complete subsection reference.

- [disable_threat_mesh](data-sources--http_loadbalancer--reference--group-017.md#canonical-78bb759eccbac27147998b2f88d803fd565a71c69942e5d39ca6ed8e300eca05): complete subsection reference.

- [disable_trust_client_ip_headers](data-sources--http_loadbalancer--reference--group-017.md#canonical-c79bdd7560149381f27ac6617a5e2c14f63acf9bb45638c01728fdcb0bd6a0c4): complete subsection reference.

- [disable_waf](data-sources--http_loadbalancer--reference--group-017.md#canonical-67933adc1825eb46b905495bf5963c44e2169f663f6b3c8ea5b9b08c6abe87c3): complete subsection reference.

- [do_not_advertise](data-sources--http_loadbalancer--reference--group-017.md#canonical-362e3ea7ff37ef9971de06d9656e93e613d193bd2a68f5912c4f7f01f05203aa): complete subsection reference.

<a id="canonical-9a05d8601f2f1c55c73aa74561bd4d5d48be4027a898a66a4cc2d636450fe824"></a>

<a id="canonical-d1c6e0990f708087cc42feeb8c64a2ebb4c2af7b0aa9dd55341ea7d6763a9e34"></a>

## domains property — Property reference / 15e40702d135 / 7

Type: `["list", "string"]`. Computed.

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

- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab): complete subsection reference.

- [enable_challenge](data-sources--http_loadbalancer--reference--group-017.md#canonical-0580b32e533934838810b561171d27d15f1c857edc4a222427c359305befa4aa): complete subsection reference.

- [enable_ip_reputation](data-sources--http_loadbalancer--reference--group-017.md#canonical-aaa2e62f5cdd092b54ec1c2b8d102d67ce63fd12ca4f464b132fd31428945e25): complete subsection reference.

- [enable_malicious_user_detection](data-sources--http_loadbalancer--reference--group-017.md#canonical-d6e47ac52e62bc579bc00d01b31e62ac48677812d725d736decd3be81c8e489a): complete subsection reference.

- [enable_threat_mesh](data-sources--http_loadbalancer--reference--group-017.md#canonical-55b4a75f7dde05f4f4c851830cf84dcb2c4adca68704256b98a57e131fbdd9f9): complete subsection reference.

- [enable_trust_client_ip_headers](data-sources--http_loadbalancer--reference--group-017.md#canonical-2bcad775662bd18661990a48eac1087fb2944081ffe2b61081ec4abf1649fa10): complete subsection reference.

- [graphql_rules](data-sources--http_loadbalancer--reference--group-017.md#canonical-1370b5a83b8a2ba11050926c0526c4373a074187e59fec6d2d3dff150b688eae): complete subsection reference.

- [http](data-sources--http_loadbalancer--reference--group-017.md#canonical-6901a453d09cfcee9e11fc8486b63aaba8056dfe0273a80fdfafc2c0920ef5e9): complete subsection reference.

- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171): complete subsection reference.

- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-e2ff5feee3a1ae293158939a97729795c3d8a7d9b4042a95c2f12867960d7cdd): complete subsection reference.

<a id="canonical-f51d2bb7a67a01a2716dced37c46b0e5fb1608b6b892b4baa726523886805e2f"></a>

<a id="canonical-3f3e8fc9e2375c17b851ba8e15b7c9501d586db53ca4b94929db12a3bdd0cb90"></a>

## id property — Property reference / 15e40702d135 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

- [js_challenge](data-sources--http_loadbalancer--reference--group-019.md#canonical-fa678d1509d4d7d2c29fe2e84b5922b2ead4a2d56a62adf3743b9a2f83301b55): complete subsection reference.

- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1f0af2cf2799ca3b7b6fad197d52e9638d499a1703cece2fc253002e639c281c): complete subsection reference.

- [l7_ddos_action_block](data-sources--http_loadbalancer--reference--group-019.md#canonical-d92e3a4521449089119b5bb699e13cff1b3943b6f7568f00be8b265d9c7717e4): complete subsection reference.

- [l7_ddos_action_default](data-sources--http_loadbalancer--reference--group-019.md#canonical-7924678dc1dacf310a0f4ab8a556d9fa849a33458912b776c881f6de8b9f4724): complete subsection reference.

- [l7_ddos_action_js_challenge](data-sources--http_loadbalancer--reference--group-019.md#canonical-f1dc8977853c97729be074e0269adb9eb7a2e6bc5823405bd4a7c49281adb1ed): complete subsection reference.

- [l7_ddos_protection](data-sources--http_loadbalancer--reference--group-019.md#canonical-cfab91794a41686085527af31572416ff30d617ac122c946145dfbd07397f8df): complete subsection reference.

<a id="canonical-f2f808edb4b581e98a3bdd01819edb194640d8b2fbd0a5997452deb8a83d56b5"></a>

<a id="canonical-aa894c34fc80d0fb7b297dbe2b18777d9017f202c1dc4bf14e144d2ae2b36dd6"></a>

## labels property — Property reference / 15e40702d135 / 9

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

- [least_active](data-sources--http_loadbalancer--reference--group-019.md#canonical-c91bb498e28c038e9f29bf34345a12b03b3d7ce5a4ea833f9318fb1aeed8cb21): complete subsection reference.

- [malware_protection_settings](data-sources--http_loadbalancer--reference--group-019.md#canonical-fb08d7c2412a19079628ab7e581422d69d6c9e8af87c9b983e2d06ea711a5a97): complete subsection reference.

- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-c86959ecdb53dc4726f527165c6759f712fa32338df3ccbc9d9bb906e68f6611): complete subsection reference.

- [multi_lb_app](data-sources--http_loadbalancer--reference--group-020.md#canonical-f63d7c36427b75f5ef6b3b3593a8aa83c0bcb29a8bbfcebe7f26ba2449bbf8f4): complete subsection reference.

<a id="canonical-cefb79ef08a86336f1c06cced23c9d0efca649083447346c18d31c0673950bf5"></a>

<a id="canonical-e53a882cb31f5133ea2944369fbc80e40527869d8e32b2e3ca9dfd4698659e88"></a>

## name property — Property reference / 15e40702d135 / 10

Type: `"string"`. Required.

Name of the HTTPLoadBalancer.

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

<a id="canonical-8f4134de42e55a4d256e2e16d4e293d0b35ee4bbc38d3642b852a597cc1f265c"></a>

<a id="canonical-4a2802422cad2a7c4f5fc4f1936a1adef31d2ccf4504b4385e6b3bedf27a0f08"></a>

## namespace property — Property reference / 15e40702d135 / 11

Type: `"string"`. Required.

Namespace where the HTTPLoadBalancer exists.

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

- [no_challenge](data-sources--http_loadbalancer--reference--group-020.md#canonical-0fe873ff4626cd57cc3aa6d49c886ef59c1538a9e081104dac8a319e8662f2f2): complete subsection reference.

- [no_service_policies](data-sources--http_loadbalancer--reference--group-020.md#canonical-46a2dec00244f94cd523ebdc05f6f2e192e43a716c1ec25fec8129e80aa69839): complete subsection reference.

- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-020.md#canonical-5ae23167a0dfbe57af9a3a473287c8a34fb3bdec8278ec98801ae5b3dff108f6): complete subsection reference.

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-4b7d31aa70160b8a4cb22d1d279816764ff6c44143f1ed5549e2dfbd97ced374): complete subsection reference.

- [protected_cookies](data-sources--http_loadbalancer--reference--group-022.md#canonical-be4cbb05d0eb2a144c0092ef2c3da5c3e2fe5e7fb7bc9d7641e25f48fda71855): complete subsection reference.

- [random](data-sources--http_loadbalancer--reference--group-022.md#canonical-e5bf189bc969780bd93b95faed2206bf49f5f1b3fc912f32403ed032f262b529): complete subsection reference.

- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-a2773b65b59fb5112f862f175b75084921aa8340dfdf9a9677b28b502ef4e739): complete subsection reference.

- [ring_hash](data-sources--http_loadbalancer--reference--group-022.md#canonical-eb0f6502296a7d3ad32cbe5d7b650ad8904c4ad46d9151d0118a8a74b08e8f3b): complete subsection reference.

- [round_robin](data-sources--http_loadbalancer--reference--group-022.md#canonical-9026a0dc1016de72d40836f2636cb220ee7076668244e68df27aabfdc22a32d5): complete subsection reference.

- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0): complete subsection reference.

- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--reference--group-025.md#canonical-b6c3a1417be308a40152934c38c279e30755355f2674475995819812f65a9671): complete subsection reference.

- [sensitive_data_policy](data-sources--http_loadbalancer--reference--group-025.md#canonical-555e1fd9db4016e1051d09b85d6284a1b66657e93495db6fab8e93b1b137257d): complete subsection reference.

- [service_policies_from_namespace](data-sources--http_loadbalancer--reference--group-025.md#canonical-ce1121f145bd5fca71bb8b705f4d8310b91f390aa749824440eb6b34acfb0fde): complete subsection reference.

- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4): complete subsection reference.

- [slow_ddos_mitigation](data-sources--http_loadbalancer--reference--group-025.md#canonical-5b08a6f056656710a9d8870bdb14818e3364d4d3b2f17386ae851550fede5001): complete subsection reference.

- [source_ip_stickiness](data-sources--http_loadbalancer--reference--group-025.md#canonical-5e72e062dee812d3a84a0f29d22dbb11853853af05d0b9684665c211ae9eeda8): complete subsection reference.

- [system_default_timeouts](data-sources--http_loadbalancer--reference--group-025.md#canonical-425fe42824262babbc98f50fc32dde2d8970ab573dfe5fa4c70f101bd9509ebd): complete subsection reference.

- [trusted_clients](data-sources--http_loadbalancer--reference--group-025.md#canonical-dcb51cdd2d49f8b7dbd34a96c17f533bb9ec118232e063bb7e15a5f11b0fb9c8): complete subsection reference.

- [user_id_client_ip](data-sources--http_loadbalancer--reference--group-025.md#canonical-ab8c8d0d78dee7b8447509f4d2cb563d77206d05b230ebcefad4e0a53423b0f0): complete subsection reference.

- [user_identification](data-sources--http_loadbalancer--reference--group-025.md#canonical-f09f430b73b0f125c1db603c889dfb4e71bc501f71ee4876284156502d4dc20a): complete subsection reference.

- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe): complete subsection reference.
