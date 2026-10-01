---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62a602a8926a01b814bd30d50c12c2a9853e86bc07ecd3004e23ff7d06b6fd69"></a>

## Property reference — Property reference / 1ce8a0b86159 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- Property reference

<a id="canonical-74841d8a717a53b272f6cc9eb2492bfc53cb3b817634ced114d9cebf51352ef5"></a>

## Direct properties — Property reference / 1ce8a0b86159 / 3

- [active_service_policies](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-22eef3517bca4cba6d0e677182fafd166348ed4151c77e52d0e2a870362dc6e4): complete subsection reference.

<a id="canonical-e2c4e68edea39a71077e15e2da29bb482f07e1ad0f10b046e33066c41762d38e"></a>

<a id="canonical-f5fe9d3883c221b35d8396b68f13fd3b3695108b0b424d9c96becf719e8def44"></a>

## annotations property — Property reference / 1ce8a0b86159 / 4

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

- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-66b75078bf3f57cca6e480dadc7b071b3ab3423f328cf97b20e0eea3e110564b): complete subsection reference.

- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-362f6b7defed1ca0331731f6c193bdce3bc10a6cb21aa4e2772f46d6b5078255): complete subsection reference.

- [app_firewall](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-17341d6985723395421d8c31d46b1a17278f274f772d5fba259047f22ec78ced): complete subsection reference.

- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-46045500a1e78445615a53dd5334ad33ce2d454451e88c6a3c566c24ac8b3826): complete subsection reference.

- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0): complete subsection reference.

- [captcha_challenge](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-530bcc216ddcdaff338b41ea0b5a54b2c40d76c0a6e10e7b8e7d1673d6d269a7): complete subsection reference.

- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d): complete subsection reference.

- [cors_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-fc417be8204c4b8728a483a1863650a03902296f2284d0eebabd5417667dd223): complete subsection reference.

- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-75eebc6e30f92011b907ed5c9d3ab2a72625997eccf90b773ec6b505f338396c): complete subsection reference.

- [custom_cache_rule](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-06d10c59251d252d7fce9adb3b74d3979ae27b09a55d4941edd7d01e37b4af78): complete subsection reference.

- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ee57a81a6530732d8c51e83ef1e9b9b1d71709c8413ab7557b32a271eb2a7e65): complete subsection reference.

- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-23fb934fdc8c11aec4fa330f9fdabf8e39499a5f76a8b4537af5d91bb4c53edc): complete subsection reference.

- [default_cache_action](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-7e7f4c028619bb61b3a22fba07eaf909a5ea14f5fa6a39e44d8d2aec7140a9c7): complete subsection reference.

- [default_sensitive_data_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-44182d76825e878f69f595ae497de7899cf9189729eada12d24e7570c7aceadc): complete subsection reference.

<a id="canonical-465c026cf62676efd4b783f210cb71acb0e4efeb38ea0aa71769d4501e073080"></a>

<a id="canonical-b08672a6e0eb245aaedb3e7112e33777a41659525d3262718c4b140a29593d41"></a>

## description property — Property reference / 1ce8a0b86159 / 5

Type: `"string"`. Computed.

Description of the CDNLoadBalancer.

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

- [disable_api_definition](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-df66c2908ab3c0ecf93a8cd22af687b9430c25b9971c0b8a9e6ba728a2001a33): complete subsection reference.

- [disable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-a4d325e03dcc681fa8b89708c9bdb4988252e05b85d3696de9e0565db4361b9c): complete subsection reference.

- [disable_client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-f470ae3f8b5dcb8ac16e8a1165c62a28e5a7546a1d389d547cc001325e5a2b47): complete subsection reference.

- [disable_ip_reputation](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-6cf2ba0eb819fccf9e4e9e9f5993fa8e2dc5fafdb640cb4ddf1a83840aab34d0): complete subsection reference.

- [disable_malicious_user_detection](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-606be885391d214dacd98bf078dd9ed9168af7b9d43bf8d452c8c2ed54e1002d): complete subsection reference.

- [disable_rate_limit](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-93bf6479464cd8a577214ea9fd9a2749eb0acbcf622374fdd47b73cd19a4d798): complete subsection reference.

- [disable_threat_mesh](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-bc358ac1840523be4841d320886d9879cc831b43c07b4c9741eb51161800e999): complete subsection reference.

- [disable_waf](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-45be9a9d0453fcbf1d5b7a9e640b3ddb6c47e33543d135435339fa8f45dc2a25): complete subsection reference.

<a id="canonical-fdbeed3596d26b617d9b26e93506c752ea35524c7cd7f9b868ca002f20b27c7e"></a>

<a id="canonical-382fe7bcc7c1960880740ff61a554eb8bba13977590ae79ad82dce8b59623d2a"></a>

## domains property — Property reference / 1ce8a0b86159 / 6

Type: `["list", "string"]`. Computed.

List of fully qualified domain names. The CDN Distribution will be setup for these FQDN name(s).
\[This can be a domain or a sub-domain\].

Upstream description:

A list of fully qualified domain names. The CDN Distribution will be setup for these FQDN name(s).
\[This can be a domain or a sub-domain\]

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

- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267): complete subsection reference.

- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d08c9244a176c7a71959859b471fc4f32017bb25c0f7babbdfcb3dc9f911841a): complete subsection reference.

- [enable_ip_reputation](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f791575cc7afb803eae1edab6bd8161ce0ceef15091e48b7ee7c3fa6b3ba4eb0): complete subsection reference.

- [enable_malicious_user_detection](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-63dc44aae6483b2cb22e58f32ac31fd9657588bb024fedd084e3781a61ba5de3): complete subsection reference.

- [enable_threat_mesh](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f09d16a1eeb59b3ba5721dc4983536e2eaf4623b14596e4f6b810709ffc6ae2f): complete subsection reference.

- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f1b85ea1fa6a6e2c26e23ee36f23c36b07e6cbe74339a722ebf7694908ccc9c5): complete subsection reference.

- [http](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-008b300c77154c1484b58fff86572204ae1ec82ce5d78ccf747628a43b11a831): complete subsection reference.

- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af): complete subsection reference.

- [https_auto_cert](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-4a52b5f55a6121f3992d8b42133a29c221108de5d07075f61e7f8570d8cc5970): complete subsection reference.

<a id="canonical-6699cd77d135ddde755fafcec0ef83dc186398003458000fde5a45c8764c6466"></a>

<a id="canonical-df44fb3de521af3bfd290f1c28001dc83a677663303676b3cd641c8c4ed4c70b"></a>

## id property — Property reference / 1ce8a0b86159 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [js_challenge](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-b2d3b8b7553e766c73a39014b67c60a27e520c4fcf4518c9a44961b4ab3fd0d4): complete subsection reference.

- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4): complete subsection reference.

- [l7_ddos_action_block](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-7ffe063e878f3a3744dbd23f60e50c97276352019aab855a5dbf85a4d8b05f75): complete subsection reference.

- [l7_ddos_action_default](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-60c1d2cae5f8d7ef4588ef22dc10e409f573ad59f7fa0749e60a4cc7bb905003): complete subsection reference.

- [l7_ddos_action_js_challenge](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-4064035bd1d3552bb9ccb82ffbae60db2624cc9c79d7f8b305061c5af59d3f48): complete subsection reference.

<a id="canonical-b6d0bfa503d2f2dbe8ffa3eda48403d64abf4032f7c3158eab92a132ab411d52"></a>

<a id="canonical-a515064cb2a1e5c177640c00df46dd8b95bc387944bd89ac9f472e2871e8c4f0"></a>

## labels property — Property reference / 1ce8a0b86159 / 8

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

<a id="canonical-2a471580d36c8774686fe829d3c10f4a0c5780bd730f3e3bd53db48d02fc6d22"></a>

<a id="canonical-5b79f97500a06efc70cae41fe57287fea1650ad3416b83f009dcc75d7af4a22c"></a>

## name property — Property reference / 1ce8a0b86159 / 9

Type: `"string"`. Required.

Name of the CDNLoadBalancer.

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

<a id="canonical-588f86a6505b74f05521bb512463227a0a20beca3447eb7d80c1282374e58b94"></a>

<a id="canonical-820969f25b8867da18b81af269016e5a2dc3b3a279f9ba2bac5a449297db55b2"></a>

## namespace property — Property reference / 1ce8a0b86159 / 10

Type: `"string"`. Required.

Namespace where the CDNLoadBalancer exists.

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

- [no_challenge](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-6e529dfa5d955f8b4335cc866733193c137c5e9d565daf1b27a3fa9022f7ec64): complete subsection reference.

- [no_service_policies](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-107e8e679b873906e8381ab7c4ccbea59f6bd95936693ad2aff2271ba7a5e3a6): complete subsection reference.

- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654): complete subsection reference.

- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3d6cc47dd4ab6896f9c02762148d92d2ca1f16b4daa0bc5357e63017b529c4b4): complete subsection reference.

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-4324e730ba27d6afbaf87c0b8e02942e3555725d734051ed25120d8bd87efb8e): complete subsection reference.

- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-8bbb1c9226b8a8e1a08b621c2d00de5afe37a24525ca33acd988ea10eee90deb): complete subsection reference.

- [rate_limit](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-12bed4f6c3fd4f4da68f6b705b22bddc00b50871ecdae9367b04e4df3ea421b3): complete subsection reference.

- [sensitive_data_policy](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-bb2647b3428854666f974efe2d4accde407916594bdd794de0a14630e2c93d77): complete subsection reference.

- [service_policies_from_namespace](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-87808731574bc92e863e99557fbbe8621d30dfe501227638348645fb3d6ae8c4): complete subsection reference.

- [slow_ddos_mitigation](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-01c67b15d3a42bb9a1e992d885e2aaad2b664fbfb5f776a3c0c6d39c912ab599): complete subsection reference.

- [system_default_timeouts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-fa38e48238d43a1d52379d97249f1ee098d2c713f9459bf92678b51c370bbf0f): complete subsection reference.

- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-42424e14e1fa4134094c259fc2bdc475057248d372dc4b9a2620501636ee872e): complete subsection reference.

- [user_id_client_ip](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-51bae569a01abc71d9356bedb70367bcacc4ff50a745eb8a820ce86d9ebea852): complete subsection reference.

- [user_identification](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-cc94a6d81f55468ce844b7b39e7af9c0e4a0c8d3b1eb5554b18b65b5a24008a8): complete subsection reference.

- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-03b49820156dc509ea8c20ab1c05f13581583c48a90846de0afcd4c0e6e2bc37): complete subsection reference.
