---
page_title: "xcsh_route reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route reference."
---

# xcsh_route reference

<a id="canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b987204cb51f8dca068640f685ad08e7dde3762c32dd36571c214ddfbeb1368"></a>

## Property reference — Property reference / 972b71ee7c1b / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- Property reference

<a id="canonical-5985236e2e16e879a07a3e8b7475a1451ac3b6cb3162989f3c370e681b441f81"></a>

## Direct properties — Property reference / 972b71ee7c1b / 3

<a id="canonical-c863f80fb0edc65bbd47ded6ee9173188eb803fd924a007db2da3f5ebcd9cefb"></a>

<a id="canonical-d4ea5c216bda874c22d7f704c2c7639bbab341843f3dd3103c034def45ec0ec6"></a>

## annotations property — Property reference / 972b71ee7c1b / 4

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

<a id="canonical-3ecc1e6d495cf79cc848ccddc83ce3d62aa7cddb791995d2da5df0203f571e7d"></a>

<a id="canonical-4d06fd2e240a6023cea73a2e830b60289a74ee242e120a07a64f003686e3f3dc"></a>

## description property — Property reference / 972b71ee7c1b / 5

Type: `"string"`. Computed.

Description of the Route.

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

<a id="canonical-0f83168b3f28467887a903c3be5b21690cb6314c71a4bf43eeb2ce93bd5ae2c4"></a>

<a id="canonical-c7fc501cb0b86c09c5facb2026f510ee7ef09b5efaabd05effc80bbffb637289"></a>

## id property — Property reference / 972b71ee7c1b / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3917353c742b6b530bc5afece8099142eb7453609f302fbe01b833cf47611ba5"></a>

<a id="canonical-bf49a2656d79a95b676135555f963b883e7495ce47b935e38b4faa388f9f17d6"></a>

## labels property — Property reference / 972b71ee7c1b / 7

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

<a id="canonical-7b7d24b7bf78c749c2f7e250d9c43b26ea3b27969aefc9e1c47c265b5f1c50cd"></a>

<a id="canonical-dbdeb55e46bfe03f178f80d976044a2a60f593cb7cc5ac77e37316cabd9b88ea"></a>

## name property — Property reference / 972b71ee7c1b / 8

Type: `"string"`. Required.

Name of the Route.

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

<a id="canonical-444e2579ff3d8715f67aba2ae8f1919aa447d3e5c3e0704170de0085905a9805"></a>

<a id="canonical-8f7f6324f9418c6be6f3d463fc6e3c4d901a3bac69f08e85cb9c74f7f0b6a235"></a>

## namespace property — Property reference / 972b71ee7c1b / 9

Type: `"string"`. Required.

Namespace where the Route exists.

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

- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49): complete subsection reference.

<a id="canonical-f5561416efa3a39f5fae79904b5fce080e5be712f23224b06e9b1c75d1a1fe74"></a>

## All schema paths — Property reference / 972b71ee7c1b / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--route--reference--group-001.md#canonical-c863f80fb0edc65bbd47ded6ee9173188eb803fd924a007db2da3f5ebcd9cefb) |
| `description` | [description](data-sources--route--reference--group-001.md#canonical-3ecc1e6d495cf79cc848ccddc83ce3d62aa7cddb791995d2da5df0203f571e7d) |
| `id` | [id](data-sources--route--reference--group-001.md#canonical-0f83168b3f28467887a903c3be5b21690cb6314c71a4bf43eeb2ce93bd5ae2c4) |
| `labels` | [labels](data-sources--route--reference--group-001.md#canonical-3917353c742b6b530bc5afece8099142eb7453609f302fbe01b833cf47611ba5) |
| `name` | [name](data-sources--route--reference--group-001.md#canonical-7b7d24b7bf78c749c2f7e250d9c43b26ea3b27969aefc9e1c47c265b5f1c50cd) |
| `namespace` | [namespace](data-sources--route--reference--group-001.md#canonical-444e2579ff3d8715f67aba2ae8f1919aa447d3e5c3e0704170de0085905a9805) |
| `routes` | [routes](data-sources--route--reference--group-001.md#canonical-3818f9702f646d3138646235c3b68cbcd9e90f2a8fe07d0b31344c291afc9eee) |
| `routes.bot_defense_javascript_injection` | [routes.bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-09391bd1fbb8fe42bb054c778c3313f8c93c41b21537e62ce4340ee0553963a6) |
| `routes.bot_defense_javascript_injection.javascript_location` | [routes.bot_defense_javascript_injection.javascript_location](data-sources--route--reference--group-001.md#canonical-e1aafe698138042fee4080337f58aae31131ee10d2bdc56f329df4f0d14d892c) |
| `routes.bot_defense_javascript_injection.javascript_tags` | [routes.bot_defense_javascript_injection.javascript_tags](data-sources--route--reference--group-001.md#canonical-5d804b902de1eeb2b32236ee112c641a4a74fb0ec402c59fbf557a04b703653b) |
| `routes.bot_defense_javascript_injection.javascript_tags.javascript_url` | [routes.bot_defense_javascript_injection.javascript_tags.javascript_url](data-sources--route--reference--group-001.md#canonical-35ffa6285c5c349baad59792af51da1c216bd5a74316489afa1d8de6f2495e9a) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes](data-sources--route--reference--group-001.md#canonical-defde4960a44f77efe68219039f8083b43bc498e07022e3c513ef928ef0097ae) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag](data-sources--route--reference--group-001.md#canonical-7d0ec7ec23788834f2f76de9d391843c12f95db345e0800544f5a603baf0c82a) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value](data-sources--route--reference--group-001.md#canonical-9ce296a2a69235243a52231e58004bb101fd447931c3e820dbcd50a8755fd31b) |
| `routes.disable_location_add` | [routes.disable_location_add](data-sources--route--reference--group-001.md#canonical-79d202c1be3df709cde1f1ed8400102350cf11b9843219540de920b22fb13de8) |
| `routes.inherited_bot_defense_javascript_injection` | [routes.inherited_bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-eb8314fdcbbb8e5f5af67b17368af644fdf9b0e98e4e5c35ee727f30e783f3da) |
| `routes.inherited_waf_exclusion` | [routes.inherited_waf_exclusion](data-sources--route--reference--group-001.md#canonical-02a735e0e99f2f89e29f63aaa1acbb78e98c1e1a219b4b44f4c11e6dcd65b359) |
| `routes.match` | [routes.match](data-sources--route--reference--group-001.md#canonical-077244628db60d731679da110ec40d95b9b27739ac105a98d12615b024bfc200) |
| `routes.match.headers` | [routes.match.headers](data-sources--route--reference--group-001.md#canonical-fbb0dda8a0fdb710d3ec3ea61ecba15fc242acedf7f77138defe03e2ba6f9329) |
| `routes.match.headers.exact` | [routes.match.headers.exact](data-sources--route--reference--group-001.md#canonical-b4d6922f6a2285389fc3a0fbfafea45bf39cdfc669978697659eaf5666f5b4df) |
| `routes.match.headers.invert_match` | [routes.match.headers.invert_match](data-sources--route--reference--group-001.md#canonical-119d8418ec0e8cc31d6c03e234738282689e5d86ed2f2ba622b9acb32a256c8f) |
| `routes.match.headers.name` | [routes.match.headers.name](data-sources--route--reference--group-001.md#canonical-0e1c774f1a9f2ccee48034b4a7016704914a675e4b39c285a2d5ddd0566507f0) |
| `routes.match.headers.presence` | [routes.match.headers.presence](data-sources--route--reference--group-001.md#canonical-1cc49bf5be7b80bdde342721d6b38b7399b6ab301bce7da681d051ef4dfb977f) |
| `routes.match.headers.regex` | [routes.match.headers.regex](data-sources--route--reference--group-001.md#canonical-eafeb058c4161cfe61fc7db743cdd6068afea0572851e198ab445b9ea79075e3) |
| `routes.match.http_method` | [routes.match.http_method](data-sources--route--reference--group-001.md#canonical-aec1d0a9a245c795dd57c16745b19070248edd88d88e2fee1bdbc6a275816036) |
| `routes.match.incoming_port` | [routes.match.incoming_port](data-sources--route--reference--group-001.md#canonical-c8b4a189c56fd076868d6cdaf799e621b5d696e4ddd93cf34cae325b4c6c2f9a) |
| `routes.match.incoming_port.no_port_match` | [routes.match.incoming_port.no_port_match](data-sources--route--reference--group-001.md#canonical-675a550d4eb0dd3a83e836599282ce05d5408ab85692f4410d1de2ab6bec02d6) |
| `routes.match.incoming_port.port` | [routes.match.incoming_port.port](data-sources--route--reference--group-001.md#canonical-e2ca73b2b37634fbcdf319761295e11ce939e30c28e58a3747556075fbbe2f59) |
| `routes.match.incoming_port.port_ranges` | [routes.match.incoming_port.port_ranges](data-sources--route--reference--group-001.md#canonical-9e7f6a8610d37e963e72605aafbd94b8793b0a8f559298b63bb948c90a83df12) |
| `routes.match.path` | [routes.match.path](data-sources--route--reference--group-001.md#canonical-c25a07b85b3d723f7897cde2c2c02844dd7e9785f183c139364aa2250243f463) |
| `routes.match.path.path` | [routes.match.path.path](data-sources--route--reference--group-001.md#canonical-0850806d06c5116d0105bd2f16ad73a01b238b90679a8e1b32a0afac4b1c43d1) |
| `routes.match.path.prefix` | [routes.match.path.prefix](data-sources--route--reference--group-001.md#canonical-d1daed5070510236ee4c6211eb904e753060d2c4a4252fdaddd2107d810c3a03) |
| `routes.match.path.regex` | [routes.match.path.regex](data-sources--route--reference--group-001.md#canonical-17be8e041ea57f9c3fcb759d1b683e491a836a39412e7392219add1b96bd3daa) |
| `routes.match.query_params` | [routes.match.query_params](data-sources--route--reference--group-001.md#canonical-628a149bf8c5ec16faf4ad26d932df706e701b03b58dccab34c3e9174dfe3287) |
| `routes.match.query_params.exact` | [routes.match.query_params.exact](data-sources--route--reference--group-001.md#canonical-84ba224d392d660870b1648f0899eaa112777b7a05546dc811ab80449d3387e2) |
| `routes.match.query_params.key` | [routes.match.query_params.key](data-sources--route--reference--group-001.md#canonical-f06b901c7704bca4590b89448304679637d3c2b6d8cfcaca2f9f7f8d6617c6e4) |
| `routes.match.query_params.regex` | [routes.match.query_params.regex](data-sources--route--reference--group-001.md#canonical-f88638e274a9b5017df4d4d1043e140eb9403bc283016811d75b89bcf2a29bab) |
| `routes.request_cookies_to_add` | [routes.request_cookies_to_add](data-sources--route--reference--group-001.md#canonical-6b646bfced6ec70f3524aadf8b9c575ab4d08ca820edf375b2015049c18cafe8) |
| `routes.request_cookies_to_add.name` | [routes.request_cookies_to_add.name](data-sources--route--reference--group-001.md#canonical-4651daa0bf640582c7e90df1b1f7a85fb9f432f0f6b94e02d5618c85fac1af10) |
| `routes.request_cookies_to_add.overwrite` | [routes.request_cookies_to_add.overwrite](data-sources--route--reference--group-001.md#canonical-97819d859a9b2dfbfdf7150eac7a71d9f52882e9f7859a89216f06ebdfd7857e) |
| `routes.request_cookies_to_add.secret_value` | [routes.request_cookies_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-f37aa9c1d3807d4d06bdc1fa731c81b091a035f65f3905a343f25eb9f29f52b1) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--route--reference--group-001.md#canonical-a3557aac408c18e23d052f53c8242b4c138a66b92674399bb3fbc1f580970b3b) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--route--reference--group-001.md#canonical-4c0311dc1746d41ec21bdcdd88786f050e37df83982cf3d2d1fc277b9d49b638) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--route--reference--group-001.md#canonical-d84acbecaad92a463a4fdbde891f5a84cfa3c2dda7d0c28c1cd28e3ae86f782d) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--route--reference--group-001.md#canonical-4ef2b9a9e2a9dd5b7d14594c17e83dcc1a78f815adac8505c0a50852dc4c7895) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info` | [routes.request_cookies_to_add.secret_value.clear_secret_info](data-sources--route--reference--group-001.md#canonical-6d2a18d2a12a9b7c7d4ae6b4742f41ba83f211a436216d5a96033147feda0f12) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [routes.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--route--reference--group-001.md#canonical-160b6f9d9eda68c3fe51628a65c19dcb3b0026ca3d83b91be4f95e52d7fa623c) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info.url` | [routes.request_cookies_to_add.secret_value.clear_secret_info.url](data-sources--route--reference--group-001.md#canonical-2ebf0c0a449193ca02f955a3a6e9ae3bcda14a931e1ed6824eef4699899cc0cf) |
| `routes.request_cookies_to_add.value` | [routes.request_cookies_to_add.value](data-sources--route--reference--group-001.md#canonical-0a24905c0176b0126290a481f40f8730d66718c40dd1172b92570b8d2dd03a5d) |
| `routes.request_cookies_to_remove` | [routes.request_cookies_to_remove](data-sources--route--reference--group-001.md#canonical-cf2b2c514b44eb409e2c543c31e223f2f593290f199e1c00624bf4c556715a5e) |
| `routes.request_headers_to_add` | [routes.request_headers_to_add](data-sources--route--reference--group-001.md#canonical-c32df7f7b50a14b673ddd6e89cbd58bf258a18b0884b384e785a15ddb6a44aa0) |
| `routes.request_headers_to_add.append` | [routes.request_headers_to_add.append](data-sources--route--reference--group-001.md#canonical-e87d571fa2c40bbcdfa524b8840aed7623f3a5f6d588e6d28a87b17d7862e171) |
| `routes.request_headers_to_add.name` | [routes.request_headers_to_add.name](data-sources--route--reference--group-001.md#canonical-4fe472de5809cef236cd681d89e822830e5f810ce35d19fafdc410c1b0c53c0e) |
| `routes.request_headers_to_add.secret_value` | [routes.request_headers_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-f8bc640bb08b44708e49b065b70b16c2ac6aca5d8a387c4edbf4f1d38d5cfbb9) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info` | [routes.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--route--reference--group-001.md#canonical-40662a722c36eb16e0f777151940accfd4d5a4828101ab51cc45171b0751598c) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--route--reference--group-001.md#canonical-ed2e47215b0f6c7db924d34115882d39023c575f2a7fb3202095b788732fe401) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.location` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--route--reference--group-001.md#canonical-5edea939ec828dd53430adb7d5d68fcd24ea7a940dfc3b0da1cfc4331035c6f1) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--route--reference--group-001.md#canonical-2d2457d9a109064ee66c5dfb5df6fd8d10ffc49f7f2e481e87c15877f577a076) |
| `routes.request_headers_to_add.secret_value.clear_secret_info` | [routes.request_headers_to_add.secret_value.clear_secret_info](data-sources--route--reference--group-001.md#canonical-e96f75ba9614f86a6d7b5ea52a054cce4e2dc64bcadb48d212da9f5da4ef2945) |
| `routes.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [routes.request_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--route--reference--group-001.md#canonical-e88dd9d865ca7a155c97350ab142dd52f58576e0d26c25e0d12780757b0eb652) |
| `routes.request_headers_to_add.secret_value.clear_secret_info.url` | [routes.request_headers_to_add.secret_value.clear_secret_info.url](data-sources--route--reference--group-001.md#canonical-cc889e67303577098be46d477b2db247eb07e1192d83ea08636ad86213f4f372) |
| `routes.request_headers_to_add.value` | [routes.request_headers_to_add.value](data-sources--route--reference--group-001.md#canonical-807087a6c11435229ce842cb8c146c820bc456bfbc5856ad7ab1be3276b7751d) |
| `routes.request_headers_to_remove` | [routes.request_headers_to_remove](data-sources--route--reference--group-001.md#canonical-94f4f9c94130f2238f86de61e455118bd096568b783a6a753f2cdb23b1396126) |
| `routes.response_cookies_to_add` | [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e22dc8ccc9b9c902c4b8b67f74b9347779d175abe41249a2488cec6798dfa62f) |
| `routes.response_cookies_to_add.add_domain` | [routes.response_cookies_to_add.add_domain](data-sources--route--reference--group-001.md#canonical-d02a36301ecebc16bc396052e3d3e7627066a7212f816380c2fcdc39cb2544e2) |
| `routes.response_cookies_to_add.add_expiry` | [routes.response_cookies_to_add.add_expiry](data-sources--route--reference--group-001.md#canonical-e9ee6b87f0f77eb2d141156bf5e40d31dadbd67a692636e175bf39c332a905cc) |
| `routes.response_cookies_to_add.add_httponly` | [routes.response_cookies_to_add.add_httponly](data-sources--route--reference--group-001.md#canonical-7378748ecb2dc85b0c1be817dd9b457393670691e2b49a00cc4c9f10c8753548) |
| `routes.response_cookies_to_add.add_partitioned` | [routes.response_cookies_to_add.add_partitioned](data-sources--route--reference--group-001.md#canonical-56fd5ae4c35d4c05e4fa887a955a7d344aa770bf61a0babb9368f85678a0f998) |
| `routes.response_cookies_to_add.add_path` | [routes.response_cookies_to_add.add_path](data-sources--route--reference--group-001.md#canonical-3c498016660620778bd971ee2f6c384c461e43c5fdd79ac9ede498eb83994966) |
| `routes.response_cookies_to_add.add_secure` | [routes.response_cookies_to_add.add_secure](data-sources--route--reference--group-001.md#canonical-111c77178d0d671b56663efa3e22b071bf4ddc39c0d78f4ce80199a26c9680ff) |
| `routes.response_cookies_to_add.ignore_domain` | [routes.response_cookies_to_add.ignore_domain](data-sources--route--reference--group-001.md#canonical-cd27c8344681e1a0f7e273ebe9a8fbfa106e2d19727afa74c5fec2b931199439) |
| `routes.response_cookies_to_add.ignore_expiry` | [routes.response_cookies_to_add.ignore_expiry](data-sources--route--reference--group-001.md#canonical-0fdf96edef3b7af8a311859a8379d981c28929281f707180199af7af7ff8b522) |
| `routes.response_cookies_to_add.ignore_httponly` | [routes.response_cookies_to_add.ignore_httponly](data-sources--route--reference--group-001.md#canonical-c357f73075849990912020f23f9263815e0c7d5596ebced724dfc3b20f3b98fe) |
| `routes.response_cookies_to_add.ignore_max_age` | [routes.response_cookies_to_add.ignore_max_age](data-sources--route--reference--group-001.md#canonical-d8afaa6917acfa8349aa1548b14bd5389c4773303d9260cc4628a3784ec1b72a) |
| `routes.response_cookies_to_add.ignore_partitioned` | [routes.response_cookies_to_add.ignore_partitioned](data-sources--route--reference--group-001.md#canonical-b0788674f651227f00e47ffa3e25d873b77d73b4a42d3e5d13e5515bb73ce697) |
| `routes.response_cookies_to_add.ignore_path` | [routes.response_cookies_to_add.ignore_path](data-sources--route--reference--group-001.md#canonical-0299cd1499ed2bbe50a3487b99a1ab643d0e55f22c344d536f93a43e4f6698d5) |
| `routes.response_cookies_to_add.ignore_samesite` | [routes.response_cookies_to_add.ignore_samesite](data-sources--route--reference--group-002.md#canonical-8d1471098ca4bfdf4118d30b27e2aaf4f3e05bfe2da966ba8133525fdc573ecb) |
| `routes.response_cookies_to_add.ignore_secure` | [routes.response_cookies_to_add.ignore_secure](data-sources--route--reference--group-002.md#canonical-ee0f4c2ec7374365b946180e6e7fec2cdff944c3e906fcbdb8b7b09844dba1d3) |
| `routes.response_cookies_to_add.ignore_value` | [routes.response_cookies_to_add.ignore_value](data-sources--route--reference--group-002.md#canonical-66755cef50f6160f1bfb85883a15da37816055c4281b6b853e1ebb90b6f8916b) |
| `routes.response_cookies_to_add.max_age_value` | [routes.response_cookies_to_add.max_age_value](data-sources--route--reference--group-001.md#canonical-f6a230e162fba652d18bed0e0f216f1c97ebd2a5861d9735b739e8ccf99dc3b3) |
| `routes.response_cookies_to_add.name` | [routes.response_cookies_to_add.name](data-sources--route--reference--group-001.md#canonical-df7e29343f740a361c0a6b08a212613413d43c7fa7185683daa0d4344fefdbd9) |
| `routes.response_cookies_to_add.overwrite` | [routes.response_cookies_to_add.overwrite](data-sources--route--reference--group-001.md#canonical-54c39677f9fed452d6892daa91fcd73f2fe3e65ad3fda3c5916b8cdae928ae78) |
| `routes.response_cookies_to_add.samesite_lax` | [routes.response_cookies_to_add.samesite_lax](data-sources--route--reference--group-002.md#canonical-f2279b41ba0800ab3f5f647a8cbd1b8678330af61800fbb3ae3bea0485985cea) |
| `routes.response_cookies_to_add.samesite_none` | [routes.response_cookies_to_add.samesite_none](data-sources--route--reference--group-002.md#canonical-9cb509c555360a358891a59ca5c83d31bcb0bdd2880638e6f9da57776fd3f72e) |
| `routes.response_cookies_to_add.samesite_strict` | [routes.response_cookies_to_add.samesite_strict](data-sources--route--reference--group-002.md#canonical-5d9c534d283f5c8045c6340e87e013cd19305c7da49e0d66fd716bae75757ffb) |
| `routes.response_cookies_to_add.secret_value` | [routes.response_cookies_to_add.secret_value](data-sources--route--reference--group-002.md#canonical-07ba7bfdb06cb61e730566bd98df1e4fb68b70ec15bf201798aa7c692009cbf2) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--route--reference--group-002.md#canonical-fdf254d701a708b25b72c747e53196d5fda28f770a73a4018294adb4b67b8edd) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--route--reference--group-002.md#canonical-d7e10e22b57f1f73bbf0da8570519e417ce0d9710e8ede96d3f589818c33e670) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--route--reference--group-002.md#canonical-8d5a91301cf1d69b6c310c7f98c55204c88223ed38088b2c4f6b57dc5c3138bf) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--route--reference--group-002.md#canonical-52c0d51e7d6abbfd6afd0e038a98ed5c0d4dae52c76adcccdaff5b8f903a6b88) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info` | [routes.response_cookies_to_add.secret_value.clear_secret_info](data-sources--route--reference--group-002.md#canonical-e74a779b617c507acb72a7003c217fc9ca5ba93acee0096d96ca29a28fa617fe) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [routes.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--route--reference--group-002.md#canonical-8be2dee45e7ce85285e1e9ab89439787d5a85e1540976aa4d5ebcb8eb830af88) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info.url` | [routes.response_cookies_to_add.secret_value.clear_secret_info.url](data-sources--route--reference--group-002.md#canonical-c9fca8ad92541a57984fc400f0e44fc2736ba614f3c77bf1733806b8017107f7) |
| `routes.response_cookies_to_add.value` | [routes.response_cookies_to_add.value](data-sources--route--reference--group-001.md#canonical-b2bd1c31052bcb661111797f9694e921d27378ddd34184435d1e28188aae8ca5) |
| `routes.response_cookies_to_remove` | [routes.response_cookies_to_remove](data-sources--route--reference--group-001.md#canonical-9a04448ae670bdc9d3e160c41870010ee46d9df15c12079fb212b7d358ba389f) |
| `routes.response_headers_to_add` | [routes.response_headers_to_add](data-sources--route--reference--group-002.md#canonical-23059c6be6391dac6cb201b4eff3897618effd8e5570c864fb6a29b502828376) |
| `routes.response_headers_to_add.append` | [routes.response_headers_to_add.append](data-sources--route--reference--group-002.md#canonical-10ce99d1494e269bc8e5866749732b7a2c73be5de3a2861938c9ca2a2a05d724) |
| `routes.response_headers_to_add.name` | [routes.response_headers_to_add.name](data-sources--route--reference--group-002.md#canonical-0811e2e41ee33da9ba41cc108062c2d9abc069482820cd78ad2447ed84f30f14) |
| `routes.response_headers_to_add.secret_value` | [routes.response_headers_to_add.secret_value](data-sources--route--reference--group-002.md#canonical-95c2a4d8fcddcf42053d217652e4226282466ceda9646a996e96b4c408bf0249) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info` | [routes.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--route--reference--group-002.md#canonical-f307c3580b1ef34e05cb4e1d07699afdbfe74b07a2508a339c2631f3af4e9a6b) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--route--reference--group-002.md#canonical-fd8d3ac77343ba1a91ff697c81ccf82279def4557e27467c641370451c3be82b) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.location` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--route--reference--group-002.md#canonical-d7e04b7f6f74d42227a0c5d6502b92976817d0adabf9bece6b87b6f2df4b60c4) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--route--reference--group-002.md#canonical-a1aa5327191a99875193f09a2a2c23700602ff896f161856d822a90185fdb777) |
| `routes.response_headers_to_add.secret_value.clear_secret_info` | [routes.response_headers_to_add.secret_value.clear_secret_info](data-sources--route--reference--group-002.md#canonical-ebb7439697383fa923bfca4cd1d58c763567d60ba5ade0da79bdc55df2b81474) |
| `routes.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [routes.response_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--route--reference--group-002.md#canonical-ef4afe96e42684db930cd0e02b5c126f02b46410b8fcb97e4e993d099323c699) |
| `routes.response_headers_to_add.secret_value.clear_secret_info.url` | [routes.response_headers_to_add.secret_value.clear_secret_info.url](data-sources--route--reference--group-002.md#canonical-0cb4d0126f74634cc3d87e0a78ee61974a65e7ca3e11c00787e2b93a51b63d45) |
| `routes.response_headers_to_add.value` | [routes.response_headers_to_add.value](data-sources--route--reference--group-002.md#canonical-51acefade0b42f4e9c634bee6a13a5eec1fb3c0593c8e825460d8ee2e0c54c1d) |
| `routes.response_headers_to_remove` | [routes.response_headers_to_remove](data-sources--route--reference--group-001.md#canonical-10ef445b60c6052f17fa04ea585b231dca88fb899d040545b167108a200302f7) |
| `routes.route_destination` | [routes.route_destination](data-sources--route--reference--group-002.md#canonical-86fd3c9260585113b938c8f32c5eadfb4395f47ee06b2280d5d78feb695468fb) |
| `routes.route_destination.auto_host_rewrite` | [routes.route_destination.auto_host_rewrite](data-sources--route--reference--group-002.md#canonical-f84c927fb8ca657d3f2f42c789d7826b52bc63e0c1ac664a0e6092ddbddbd762) |
| `routes.route_destination.buffer_policy` | [routes.route_destination.buffer_policy](data-sources--route--reference--group-002.md#canonical-72bc69c6be723b3e0a3f96b91f41191676d0291d57453f776e6a752fb9c8039e) |
| `routes.route_destination.buffer_policy.disabled` | [routes.route_destination.buffer_policy.disabled](data-sources--route--reference--group-002.md#canonical-02a0eeb374f13137b711f262603a17a2b9463e294a43932c7e0f5b343fd6b179) |
| `routes.route_destination.buffer_policy.max_request_bytes` | [routes.route_destination.buffer_policy.max_request_bytes](data-sources--route--reference--group-002.md#canonical-d6260284404ed8c380e7c15faddbf20f8a0626a9e4bb6941972a6e3713ec2015) |
| `routes.route_destination.cors_policy` | [routes.route_destination.cors_policy](data-sources--route--reference--group-002.md#canonical-6af05e8219973110fefb2c44ac1e1e7e50a380ccf49a4328c7d2a4e9cb0de104) |
| `routes.route_destination.cors_policy.allow_credentials` | [routes.route_destination.cors_policy.allow_credentials](data-sources--route--reference--group-002.md#canonical-80e2094aa0fe76365144dc2c0fdfb1f6899a7173dcebcfbce4c5a49b11978d8c) |
| `routes.route_destination.cors_policy.allow_headers` | [routes.route_destination.cors_policy.allow_headers](data-sources--route--reference--group-002.md#canonical-6e39cfbcbaf7b3186ba25f552711eca90e44a2fbaab7cf65380f8767fb4a1a1e) |
| `routes.route_destination.cors_policy.allow_methods` | [routes.route_destination.cors_policy.allow_methods](data-sources--route--reference--group-002.md#canonical-60826de1f60ee3513b63f094429eb08b5b65808a70a844b056c937c49fe52c23) |
| `routes.route_destination.cors_policy.allow_origin` | [routes.route_destination.cors_policy.allow_origin](data-sources--route--reference--group-002.md#canonical-7d0c6b647e1f62fbf4d86355378b2096df53ee2bb909e68a9ea08dfcf5430b9e) |
| `routes.route_destination.cors_policy.allow_origin_regex` | [routes.route_destination.cors_policy.allow_origin_regex](data-sources--route--reference--group-002.md#canonical-8a2dc4a03bb8e5923600cd49b7732e6d995b50f08756c97145c2df9e5c86ef39) |
| `routes.route_destination.cors_policy.disabled` | [routes.route_destination.cors_policy.disabled](data-sources--route--reference--group-002.md#canonical-1ca30877d8dbbcce2888ad1e172c4a2ce9dd14956ea8cfeea3237818e4e3c630) |
| `routes.route_destination.cors_policy.expose_headers` | [routes.route_destination.cors_policy.expose_headers](data-sources--route--reference--group-002.md#canonical-6286e71e93a571441c08a8f0c2aae043f6ab9b4e7e182dfa907f36f150a8af1d) |
| `routes.route_destination.cors_policy.maximum_age` | [routes.route_destination.cors_policy.maximum_age](data-sources--route--reference--group-002.md#canonical-d089addf6c81fd2a50eb5ef345a51efd11c0947f817fc21bf0fbe5fd20270916) |
| `routes.route_destination.csrf_policy` | [routes.route_destination.csrf_policy](data-sources--route--reference--group-002.md#canonical-06a40ca49870d454868bb44db7e195fca002d9e9779e237a4553245c7fb51f0b) |
| `routes.route_destination.csrf_policy.all_load_balancer_domains` | [routes.route_destination.csrf_policy.all_load_balancer_domains](data-sources--route--reference--group-002.md#canonical-4e02e24ea3a1c63d1dc7777557696f077331ca0be7e693dc6f7dfa82b403051b) |
| `routes.route_destination.csrf_policy.custom_domain_list` | [routes.route_destination.csrf_policy.custom_domain_list](data-sources--route--reference--group-002.md#canonical-74b03ca53839dd83323c8c174ed250e7c8652a5bd5bf74514a09642614d38fb0) |
| `routes.route_destination.csrf_policy.custom_domain_list.domains` | [routes.route_destination.csrf_policy.custom_domain_list.domains](data-sources--route--reference--group-002.md#canonical-a73305096596cb110656dde85bfe761a0578fb11719fff60029e4cc719d39027) |
| `routes.route_destination.csrf_policy.disabled` | [routes.route_destination.csrf_policy.disabled](data-sources--route--reference--group-002.md#canonical-b16bea6baeeb59ea5af3a28aa9f0241572b2754deeb65a685d79150259d49d7e) |
| `routes.route_destination.destinations` | [routes.route_destination.destinations](data-sources--route--reference--group-002.md#canonical-1b4f08701fa01592dbdc93d185f6cf9383754e08183775b78ed309d868fcd58f) |
| `routes.route_destination.destinations.cluster` | [routes.route_destination.destinations.cluster](data-sources--route--reference--group-002.md#canonical-edb16c06e59f184394f809ff111d7a94e18d150d2e24b9c7a30f122223821b18) |
| `routes.route_destination.destinations.cluster.kind` | [routes.route_destination.destinations.cluster.kind](data-sources--route--reference--group-002.md#canonical-8b8cbea979b1616ac5a5e3a3f213d01838fb3cdd8349892691f0b59c10ca36fb) |
| `routes.route_destination.destinations.cluster.name` | [routes.route_destination.destinations.cluster.name](data-sources--route--reference--group-002.md#canonical-d6d6996437e80f7ae9a707a4ae7891f768a40580bfbc167fa7c392c8822a7cb7) |
| `routes.route_destination.destinations.cluster.namespace` | [routes.route_destination.destinations.cluster.namespace](data-sources--route--reference--group-002.md#canonical-5ef9dd8a06652147331c789303239b864cc7ecdc13b77bdea13b3d9f9c60adad) |
| `routes.route_destination.destinations.cluster.tenant` | [routes.route_destination.destinations.cluster.tenant](data-sources--route--reference--group-002.md#canonical-763cdfbdbf6f054ef646aac923052480270fe30a407b7de55f776dc823fd7558) |
| `routes.route_destination.destinations.cluster.uid` | [routes.route_destination.destinations.cluster.uid](data-sources--route--reference--group-002.md#canonical-d0d084cdfa140b0f39dce5e46d2933eb0f835b1bb90c9dfacf49d57f8215d137) |
| `routes.route_destination.destinations.endpoint_subsets` | [routes.route_destination.destinations.endpoint_subsets](data-sources--route--reference--group-002.md#canonical-e9dc4c1becb73e364f734ed5d7815f615c8676b100a115c423a8b926a39ace19) |
| `routes.route_destination.destinations.priority` | [routes.route_destination.destinations.priority](data-sources--route--reference--group-002.md#canonical-532e3caf53081c4a2ae8972d5e79ce2174ed95f8a9fe1e533d70a0def6f8abcf) |
| `routes.route_destination.destinations.weight` | [routes.route_destination.destinations.weight](data-sources--route--reference--group-002.md#canonical-8a94d56da62b7d2707bea6972675d8a04f33d21e66d85c059d608a757827d752) |
| `routes.route_destination.do_not_retract_cluster` | [routes.route_destination.do_not_retract_cluster](data-sources--route--reference--group-002.md#canonical-f00b7541ca1eaf76f086327b58fa6de28cb7d5fc3d780642b6b58979658c9ea4) |
| `routes.route_destination.endpoint_subsets` | [routes.route_destination.endpoint_subsets](data-sources--route--reference--group-002.md#canonical-15a1dbe8854bb7f7569c502d663798ec9103920817a1d837a508b520e42bb58c) |
| `routes.route_destination.hash_policy` | [routes.route_destination.hash_policy](data-sources--route--reference--group-002.md#canonical-118eedc61e5544d464ec3a562d8496ce1d332fd88d42fcdc63a6479d041d8533) |
| `routes.route_destination.hash_policy.cookie` | [routes.route_destination.hash_policy.cookie](data-sources--route--reference--group-002.md#canonical-7cd93c21237ca122127150fcff668c0c2f39ea7a74f8a84646d7d14d4414b0b9) |
| `routes.route_destination.hash_policy.cookie.add_httponly` | [routes.route_destination.hash_policy.cookie.add_httponly](data-sources--route--reference--group-002.md#canonical-a688c4a2b03a74151f2d4e2cdca124ddb955f6e9a2cf75e3dfe27d78290d3256) |
| `routes.route_destination.hash_policy.cookie.add_secure` | [routes.route_destination.hash_policy.cookie.add_secure](data-sources--route--reference--group-002.md#canonical-3643b896dcb8faadcafc7a6cd0d2d2556b2b16d5172596a19262aa85f97c6cc4) |
| `routes.route_destination.hash_policy.cookie.ignore_httponly` | [routes.route_destination.hash_policy.cookie.ignore_httponly](data-sources--route--reference--group-002.md#canonical-1981b49c84b95919b8a0fc159b0f1fb58cec3ce3aa5f78de006730568f1329d1) |
| `routes.route_destination.hash_policy.cookie.ignore_samesite` | [routes.route_destination.hash_policy.cookie.ignore_samesite](data-sources--route--reference--group-002.md#canonical-5727f788761fdc0d7743910646be75b54994a2c846a44e5747ae73fbf8b8f57d) |
| `routes.route_destination.hash_policy.cookie.ignore_secure` | [routes.route_destination.hash_policy.cookie.ignore_secure](data-sources--route--reference--group-002.md#canonical-60a1ad6e507c46a2c4da09abac9cf5f2a50e87a3eb25346ee432b26a625870fa) |
| `routes.route_destination.hash_policy.cookie.name` | [routes.route_destination.hash_policy.cookie.name](data-sources--route--reference--group-002.md#canonical-5eb8056ab3c8aa2fe8a3a1e424f9519dfdd21c7113b8f3f11016ec6e9c967ed2) |
| `routes.route_destination.hash_policy.cookie.path` | [routes.route_destination.hash_policy.cookie.path](data-sources--route--reference--group-002.md#canonical-f2df271641f3ba4aea2c9f49eb969048176589ad7e1e1b80e1f39a861f24e0eb) |
| `routes.route_destination.hash_policy.cookie.samesite_lax` | [routes.route_destination.hash_policy.cookie.samesite_lax](data-sources--route--reference--group-002.md#canonical-513237940ccb00b5e03635c2a84087251bfed374660705cc1e8939a4914fdd7e) |
| `routes.route_destination.hash_policy.cookie.samesite_none` | [routes.route_destination.hash_policy.cookie.samesite_none](data-sources--route--reference--group-002.md#canonical-88d3bfb1153f90f3f33887b48b022d0758aa5bdf549574a36f7fec6546b147e4) |
| `routes.route_destination.hash_policy.cookie.samesite_strict` | [routes.route_destination.hash_policy.cookie.samesite_strict](data-sources--route--reference--group-002.md#canonical-ea87675aa078701d288dcbb27cb1e74548840ec1b9ef7ce9510e022254306c85) |
| `routes.route_destination.hash_policy.cookie.ttl` | [routes.route_destination.hash_policy.cookie.ttl](data-sources--route--reference--group-002.md#canonical-ef2cf430cc38ae9c990f83f3e2c7f70ca23c4103d5f2c4bb5fb858fe57a54027) |
| `routes.route_destination.hash_policy.header_name` | [routes.route_destination.hash_policy.header_name](data-sources--route--reference--group-002.md#canonical-fd43a07f53ed802cb7a03daa731f76d5efc56e81063fc455b9dd3bfa203417d3) |
| `routes.route_destination.hash_policy.source_ip` | [routes.route_destination.hash_policy.source_ip](data-sources--route--reference--group-002.md#canonical-264611856cf9bf7f22e609c8c9bfdf97aecddef62071685e5a215f1e6da78bc8) |
| `routes.route_destination.hash_policy.terminal` | [routes.route_destination.hash_policy.terminal](data-sources--route--reference--group-002.md#canonical-562f74ad6291be6c1a7de50a6a9121541faa725988168e64a02ac16a3d94d616) |
| `routes.route_destination.host_rewrite` | [routes.route_destination.host_rewrite](data-sources--route--reference--group-002.md#canonical-9fd06b30d9e9645771734acfe7bbff498f00764039acfa8a991cd70a63291123) |
| `routes.route_destination.mirror_policy` | [routes.route_destination.mirror_policy](data-sources--route--reference--group-002.md#canonical-c69ab366b458742b13c96003795bc5678aa7d763b8dc733469ec4603c0c1c516) |
| `routes.route_destination.mirror_policy.cluster` | [routes.route_destination.mirror_policy.cluster](data-sources--route--reference--group-002.md#canonical-bf3ae52f6d4093639fe110a24b808930f73403a0f7353055a139ae787f3aef86) |
| `routes.route_destination.mirror_policy.cluster.kind` | [routes.route_destination.mirror_policy.cluster.kind](data-sources--route--reference--group-002.md#canonical-7ff0e8cd3559053a388068c9ee680074cf7607dee927d89afd5bff5017fd6482) |
| `routes.route_destination.mirror_policy.cluster.name` | [routes.route_destination.mirror_policy.cluster.name](data-sources--route--reference--group-002.md#canonical-359a59d5023de2fec0eb86f51c68b8282f81b90b0393d49040f919dd695f49f9) |
| `routes.route_destination.mirror_policy.cluster.namespace` | [routes.route_destination.mirror_policy.cluster.namespace](data-sources--route--reference--group-002.md#canonical-e601a4db809ee8a12effb99aa48b08789df3e16dcce6bfacd3b43a1fd89d3adc) |
| `routes.route_destination.mirror_policy.cluster.tenant` | [routes.route_destination.mirror_policy.cluster.tenant](data-sources--route--reference--group-002.md#canonical-2f8a7a525bdcdab3ae3c8e05dcedb25d55a0b18c12751bb013c07a87d12ac7b9) |
| `routes.route_destination.mirror_policy.cluster.uid` | [routes.route_destination.mirror_policy.cluster.uid](data-sources--route--reference--group-002.md#canonical-a2c97969ee21783f6842b442180ff842950079de4c6457bb2151f3e3d17544d9) |
| `routes.route_destination.mirror_policy.percent` | [routes.route_destination.mirror_policy.percent](data-sources--route--reference--group-002.md#canonical-065a6efc3673054e65e1cf2a5eae0b17bbe2a94ec3ee30eb5728d859518184bb) |
| `routes.route_destination.mirror_policy.percent.denominator` | [routes.route_destination.mirror_policy.percent.denominator](data-sources--route--reference--group-002.md#canonical-cd383683283fc3aa973680357c2ba8a78c450cc7c301c024be8f281797713b70) |
| `routes.route_destination.mirror_policy.percent.numerator` | [routes.route_destination.mirror_policy.percent.numerator](data-sources--route--reference--group-002.md#canonical-122b765e982359d126ca9185a8b8104e88ca243f7ca7712cf90428b294da4b01) |
| `routes.route_destination.prefix_rewrite` | [routes.route_destination.prefix_rewrite](data-sources--route--reference--group-002.md#canonical-fb0ba61350cfbd2541eacfe7b35a5e3fa07716066ae030aa8d0c968257fd48fa) |
| `routes.route_destination.priority` | [routes.route_destination.priority](data-sources--route--reference--group-002.md#canonical-cf244a1b3fa36ede9e5bab163aa4c570a8056ae626fe51a6dc8f8a0714ffe75d) |
| `routes.route_destination.query_params` | [routes.route_destination.query_params](data-sources--route--reference--group-002.md#canonical-f9518de88720230cc505d7bb6d2d8b7647fc175b3f4aacd77be512f4cf21fc0b) |
| `routes.route_destination.query_params.remove_all_params` | [routes.route_destination.query_params.remove_all_params](data-sources--route--reference--group-002.md#canonical-916438db2b24839785f0680656f246a11bd05a7949026b3670d64e39d7025522) |
| `routes.route_destination.query_params.replace_params` | [routes.route_destination.query_params.replace_params](data-sources--route--reference--group-002.md#canonical-231021853fb02a848bb92d58a5f2f5fa7d9bad42531f4e988ab5a1e847e91b6c) |
| `routes.route_destination.query_params.retain_all_params` | [routes.route_destination.query_params.retain_all_params](data-sources--route--reference--group-002.md#canonical-d477835ede72418cd18ccbb20574616bb32b6891e7da5b7d499e7ac912afeb12) |
| `routes.route_destination.regex_rewrite` | [routes.route_destination.regex_rewrite](data-sources--route--reference--group-002.md#canonical-57fd0e9f35ad07c87657b2bad61e99d01e1b492119681df2b503d1ee7a9d9fc0) |
| `routes.route_destination.regex_rewrite.pattern` | [routes.route_destination.regex_rewrite.pattern](data-sources--route--reference--group-002.md#canonical-29731a3fe380421e06772854377694366d4d8d098a4e6da2b1628a069707f71e) |
| `routes.route_destination.regex_rewrite.substitution` | [routes.route_destination.regex_rewrite.substitution](data-sources--route--reference--group-002.md#canonical-60acf295fd67289cf86a6a03bbc0221017b5bf4fe06cf4c32c84fa9f49e742e1) |
| `routes.route_destination.retract_cluster` | [routes.route_destination.retract_cluster](data-sources--route--reference--group-002.md#canonical-c13cdc71b270c21b4a9b289d17c425ab6be7afda7507a790f942371fb97d2057) |
| `routes.route_destination.retry_policy` | [routes.route_destination.retry_policy](data-sources--route--reference--group-002.md#canonical-d06e080a310422bd78836708b49c38917a062cc54b27a7b77c5d0d7c7371616c) |
| `routes.route_destination.retry_policy.back_off` | [routes.route_destination.retry_policy.back_off](data-sources--route--reference--group-002.md#canonical-6fdd297dc39437a74d3680adce27c4071d7d0916bbd9b0145c5e9cb8dee56f55) |
| `routes.route_destination.retry_policy.back_off.base_interval` | [routes.route_destination.retry_policy.back_off.base_interval](data-sources--route--reference--group-002.md#canonical-306685f04e63695004902cd5b1a34776d759030bf5cacaa484efd7760a547448) |
| `routes.route_destination.retry_policy.back_off.max_interval` | [routes.route_destination.retry_policy.back_off.max_interval](data-sources--route--reference--group-002.md#canonical-7873b7b6faaabab800f94cbafe639249e0a9d830e220f90cfd200cfce33bb272) |
| `routes.route_destination.retry_policy.num_retries` | [routes.route_destination.retry_policy.num_retries](data-sources--route--reference--group-002.md#canonical-269f697c211f85ecc5363a4992e5a2f24728c163f3fb67be3e5d84023324b08c) |
| `routes.route_destination.retry_policy.per_try_timeout` | [routes.route_destination.retry_policy.per_try_timeout](data-sources--route--reference--group-002.md#canonical-9215f1b21cf36effa0e9482f7be3c02f8e1e8a72b541ddce3bc87f9da91855e5) |
| `routes.route_destination.retry_policy.retriable_status_codes` | [routes.route_destination.retry_policy.retriable_status_codes](data-sources--route--reference--group-002.md#canonical-11458ea3f08070a9cb1a33f1e818d18ee90d80e4527e5326bc0cfd9cb5a39e1a) |
| `routes.route_destination.retry_policy.retry_condition` | [routes.route_destination.retry_policy.retry_condition](data-sources--route--reference--group-002.md#canonical-ad51bf5ee324d44f1d58be9c33cc0e1cb11a16f4d8aa0e731d50a030294b5888) |
| `routes.route_destination.spdy_config` | [routes.route_destination.spdy_config](data-sources--route--reference--group-002.md#canonical-a8ff963a46d71f796ad3cdb32597fb0265db95ef99f6448ba5c73ba73fa8a140) |
| `routes.route_destination.spdy_config.use_spdy` | [routes.route_destination.spdy_config.use_spdy](data-sources--route--reference--group-002.md#canonical-b0032f45906212ed3cd05db5f2be20ed0de911fd4a31a760c462d1f0348700fb) |
| `routes.route_destination.timeout` | [routes.route_destination.timeout](data-sources--route--reference--group-002.md#canonical-4a4874141028a06f49c74318304093860e7681b9140dad1e9b29d8a0f073c134) |
| `routes.route_destination.web_socket_config` | [routes.route_destination.web_socket_config](data-sources--route--reference--group-002.md#canonical-95f214b295850102844c7c4ffe51bf91411dc806323eeec4444c0eec4b71748b) |
| `routes.route_destination.web_socket_config.use_websocket` | [routes.route_destination.web_socket_config.use_websocket](data-sources--route--reference--group-002.md#canonical-6796043c610e3aed71405785fefc406d6a8b90131fdb18bf420afffdac973d56) |
| `routes.route_direct_response` | [routes.route_direct_response](data-sources--route--reference--group-002.md#canonical-66e9d4ddfde825a4f2d87cd18afaf009bdf7495286b95c2e6e43076cd2c3c8fc) |
| `routes.route_direct_response.response_body_encoded` | [routes.route_direct_response.response_body_encoded](data-sources--route--reference--group-002.md#canonical-5afccce7b24d5d783c8d0350a59575238c13aa667ab671674588aeef05ded056) |
| `routes.route_direct_response.response_code` | [routes.route_direct_response.response_code](data-sources--route--reference--group-002.md#canonical-7d0da583aff2cf6727bdac71139b4597ccc1ebad868d5dbfed6bed6e0c920146) |
| `routes.route_redirect` | [routes.route_redirect](data-sources--route--reference--group-002.md#canonical-922989ee4c170cb535e44a2be046d61c904c14a58644dfcafebbb916f0671ad8) |
| `routes.route_redirect.host_redirect` | [routes.route_redirect.host_redirect](data-sources--route--reference--group-002.md#canonical-7f58dff392aa18c7783bf9999345edbf78cdc784d19871b67e5d74521067419c) |
| `routes.route_redirect.path_redirect` | [routes.route_redirect.path_redirect](data-sources--route--reference--group-002.md#canonical-2324076231b049fe9fe38c7f1bfb8751de9f0b86328622dc0f6060356071971c) |
| `routes.route_redirect.prefix_rewrite` | [routes.route_redirect.prefix_rewrite](data-sources--route--reference--group-002.md#canonical-2bc9721ff52379784bbe8061f19c16940e341a2d7010e343ce0a827c7b9d2c6e) |
| `routes.route_redirect.proto_redirect` | [routes.route_redirect.proto_redirect](data-sources--route--reference--group-002.md#canonical-3191dbc6b546fba46ae038f8ec796aab83d1888d1eb9ceb8a425306981fa5127) |
| `routes.route_redirect.remove_all_params` | [routes.route_redirect.remove_all_params](data-sources--route--reference--group-002.md#canonical-8e63a794405bee163436a518a27098cf9a3a05a0af180c66b3a4f26e45694f94) |
| `routes.route_redirect.replace_params` | [routes.route_redirect.replace_params](data-sources--route--reference--group-002.md#canonical-54b8f9afb284c683c463b6d57b9a9b6828ae52afedf2f7cc4695cc559f25c09f) |
| `routes.route_redirect.response_code` | [routes.route_redirect.response_code](data-sources--route--reference--group-002.md#canonical-92fe474946956f4026577ab71a3e109f7e5fa2d2d230096aebbf18ca99abe601) |
| `routes.route_redirect.retain_all_params` | [routes.route_redirect.retain_all_params](data-sources--route--reference--group-002.md#canonical-8fe6047a4e5ee4d7c058a68cec761cc44430f16f59a9b43b5fdc827eedd53702) |
| `routes.service_policy` | [routes.service_policy](data-sources--route--reference--group-002.md#canonical-18954c8f168ef3adf40c4197db6e8d5fb5f58b488bd020907a3fa40ca515758e) |
| `routes.service_policy.disable_spec` | [routes.service_policy.disable_spec](data-sources--route--reference--group-002.md#canonical-1320c0acd47cb6ee5b877d336b9b352c063a469ef117035161e444358a8a6d46) |
| `routes.waf_exclusion_policy` | [routes.waf_exclusion_policy](data-sources--route--reference--group-002.md#canonical-122789159edad8b9c72e8c840322458a88824472bfc2f4d75de78506d40c1f6e) |
| `routes.waf_exclusion_policy.name` | [routes.waf_exclusion_policy.name](data-sources--route--reference--group-002.md#canonical-7c5488c3f7b58c271f8b07286fb624385ff45956da1237369730cfb3026f9ac2) |
| `routes.waf_exclusion_policy.namespace` | [routes.waf_exclusion_policy.namespace](data-sources--route--reference--group-002.md#canonical-6cace4fd9d1d488c448c4f91f98cf542d7aa01f564371e846f87aaf27c2aba81) |
| `routes.waf_exclusion_policy.tenant` | [routes.waf_exclusion_policy.tenant](data-sources--route--reference--group-002.md#canonical-01fb399102e9e26ab4f22a446964615853e6fb60583d5f37ad91abfa5b455a15) |
| `routes.waf_type` | [routes.waf_type](data-sources--route--reference--group-002.md#canonical-008ee915769b9a965fbbb0192bc8ece7c93d9de07caa3c569422732d4fc65fb8) |
| `routes.waf_type.app_firewall` | [routes.waf_type.app_firewall](data-sources--route--reference--group-003.md#canonical-08b7da3bcdf4f42058fd372a654de5fdc4e1bed93787ada684356d253c7344fc) |
| `routes.waf_type.app_firewall.app_firewall` | [routes.waf_type.app_firewall.app_firewall](data-sources--route--reference--group-003.md#canonical-ff7eb447a7555ded0623d37c2b379b9df32a081d3c87e2c65006a3a2a2f052fd) |
| `routes.waf_type.app_firewall.app_firewall.kind` | [routes.waf_type.app_firewall.app_firewall.kind](data-sources--route--reference--group-003.md#canonical-999fc33da0f95410d9775965c47cb3bf51c04a20eefd199d72b713bb5d354b69) |
| `routes.waf_type.app_firewall.app_firewall.name` | [routes.waf_type.app_firewall.app_firewall.name](data-sources--route--reference--group-003.md#canonical-d4626025295332003c78381472dd925c45f1802c996bf11549e74ea358fc2ce7) |
| `routes.waf_type.app_firewall.app_firewall.namespace` | [routes.waf_type.app_firewall.app_firewall.namespace](data-sources--route--reference--group-003.md#canonical-969aa455b217b7fe4cf0500ef68fba943ab363c17f3949505eebb7189ecace14) |
| `routes.waf_type.app_firewall.app_firewall.tenant` | [routes.waf_type.app_firewall.app_firewall.tenant](data-sources--route--reference--group-003.md#canonical-2e71e4a96cc0fb91702c2a70ebc10f3e2c221ed83db894495226ad79e56be2b6) |
| `routes.waf_type.app_firewall.app_firewall.uid` | [routes.waf_type.app_firewall.app_firewall.uid](data-sources--route--reference--group-003.md#canonical-b0c9bf5328dfb73ba4b0ec27dbac1e0e670f3d595e2d5729a06e2ff7b8c0fc42) |
| `routes.waf_type.disable_waf` | [routes.waf_type.disable_waf](data-sources--route--reference--group-003.md#canonical-75c7cb34fa00beeb2daba6031f0e4eb54c14d719a6d3ae1fe3fcc857772c3469) |
| `routes.waf_type.inherit_waf` | [routes.waf_type.inherit_waf](data-sources--route--reference--group-003.md#canonical-17ddb98cd621b57d1bfd18867c4b38ee45bdbfe3bcdc848418aa814027ba6f10) |

<a id="canonical-b6d737439a1533c99a200bab2b0fc166cbc959db026e9ce0cc7a6baf4b7a2921"></a>

## Next pages — Property reference / 972b71ee7c1b / 11

- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3bac57d0b11e83be20ab454f508524cfcd3e92fbe65ca02ced27c26c8d25db9"></a>

## routes — routes / 2070498554c4 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- routes

<a id="canonical-3818f9702f646d3138646235c3b68cbcd9e90f2a8fe07d0b31344c291afc9eee"></a>

Type: `"list"`. Computed.

List of routes to match for incoming request.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 257,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 257,
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
    "ves.io.schema.rules.repeated.max_items": "257"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "257"
  }
}
```

<a id="canonical-4b06f63a7e8978490d7bb1cd37885018b1255131e2e09a787e5f68eaf7ea7159"></a>

## Direct properties — routes / 2070498554c4 / 3

- [bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-c2ec8eaf5c4fee182c5f4d3d30af7c79731051a3596796f547265902a7dcfd22): complete subsection reference.

<a id="canonical-79d202c1be3df709cde1f1ed8400102350cf11b9843219540de920b22fb13de8"></a>

<a id="canonical-1bfa710501ca0781e09d1bf624dabcc07a1774644649464972fe794137f4fd46"></a>

## disable_location_add property — routes / 2070498554c4 / 4

Type: `"bool"`. Computed.

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

Upstream description:

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

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

- [inherited_bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-73284df79d3594d5fbc5783b129a32c7d48daae43f13e9ab34714fb52843dea0): complete subsection reference.

- [inherited_waf_exclusion](data-sources--route--reference--group-001.md#canonical-c80363e9c26bac22ab025c48db1ef4e2f2a7b14720d9f7657b6e1cf658c4d137): complete subsection reference.

- [match](data-sources--route--reference--group-001.md#canonical-a2257940392ee77e0b7cc0101765684d7b06d8037bd3729204513dcbcb5c25f6): complete subsection reference.

- [request_cookies_to_add](data-sources--route--reference--group-001.md#canonical-fe0cf0995676155ec1ccb3ae9747442486d275b7596e31e7404a8cc9ecffcdaa): complete subsection reference.

<a id="canonical-cf2b2c514b44eb409e2c543c31e223f2f593290f199e1c00624bf4c556715a5e"></a>

<a id="canonical-fefa571f9220fcb299c9aed834ca8184ffb7a41927e56ecd6a07f992dc892f4f"></a>

## request_cookies_to_remove property — routes / 2070498554c4 / 5

Type: `["list", "string"]`. Computed.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](data-sources--route--reference--group-001.md#canonical-0865f1acc903ad9680c665112e61b6d96f02f8e9a7d9cbeed841654e49236b4f): complete subsection reference.

<a id="canonical-94f4f9c94130f2238f86de61e455118bd096568b783a6a753f2cdb23b1396126"></a>

<a id="canonical-59550d59bdcd4bb9f9a52e8962abf1e5d74838475535aaf73a4823393b34272e"></a>

## request_headers_to_remove property — routes / 2070498554c4 / 6

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

- [response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d): complete subsection reference.

<a id="canonical-9a04448ae670bdc9d3e160c41870010ee46d9df15c12079fb212b7d358ba389f"></a>

<a id="canonical-586795591dcd8171f78d74988fbea46785c348bfda35369140ed1d203eeb4781"></a>

## response_cookies_to_remove property — routes / 2070498554c4 / 7

Type: `["list", "string"]`. Computed.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](data-sources--route--reference--group-002.md#canonical-8d786a77ad164087c3e05a9a97cf1feddcd8b9906963893a648384aeec6374f1): complete subsection reference.

<a id="canonical-10ef445b60c6052f17fa04ea585b231dca88fb899d040545b167108a200302f7"></a>

<a id="canonical-7fed12a580b1a4d5161ecef9b7d0da63529dc68146f33bf9d70f6e9837b90d11"></a>

## response_headers_to_remove property — routes / 2070498554c4 / 8

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

- [route_destination](data-sources--route--reference--group-002.md#canonical-62196179190494302a68d593b134fba295962b67974de8bff861b007ba9f9dec): complete subsection reference.

- [route_direct_response](data-sources--route--reference--group-002.md#canonical-584c45695d1623f2c6198b272a75875291bfe763e53d0ba307888a33af25815a): complete subsection reference.

- [route_redirect](data-sources--route--reference--group-002.md#canonical-cd0f5be797549d439537ec1c55d5cd9098a551adfdf30ac18b20d6d021b67053): complete subsection reference.

- [service_policy](data-sources--route--reference--group-002.md#canonical-1a0cd68ab6678dc5ee8c0d95b09eb643ffc0d591cd4967766cdf58feecc75778): complete subsection reference.

- [waf_exclusion_policy](data-sources--route--reference--group-002.md#canonical-9c2da572f6774afcd533806a7075bfa9703c4c1efdc6739bc72cfa6e2bb44373): complete subsection reference.

- [waf_type](data-sources--route--reference--group-002.md#canonical-7362ef3ec8de94401ce424467749d400c77686ad5c2a78680790aca10481fc58): complete subsection reference.

<a id="canonical-49a12153c43d05149bc4bd4e4c5793342458986ef93efe6a910d280025fd60a2"></a>

## Next pages — routes / 2070498554c4 / 9

- [routes.bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-c2ec8eaf5c4fee182c5f4d3d30af7c79731051a3596796f547265902a7dcfd22)
- [routes.inherited_bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-73284df79d3594d5fbc5783b129a32c7d48daae43f13e9ab34714fb52843dea0)
- [routes.inherited_waf_exclusion](data-sources--route--reference--group-001.md#canonical-c80363e9c26bac22ab025c48db1ef4e2f2a7b14720d9f7657b6e1cf658c4d137)
- [routes.match](data-sources--route--reference--group-001.md#canonical-a2257940392ee77e0b7cc0101765684d7b06d8037bd3729204513dcbcb5c25f6)
- [routes.request_cookies_to_add](data-sources--route--reference--group-001.md#canonical-fe0cf0995676155ec1ccb3ae9747442486d275b7596e31e7404a8cc9ecffcdaa)
- [routes.request_headers_to_add](data-sources--route--reference--group-001.md#canonical-0865f1acc903ad9680c665112e61b6d96f02f8e9a7d9cbeed841654e49236b4f)
- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- [routes.response_headers_to_add](data-sources--route--reference--group-002.md#canonical-8d786a77ad164087c3e05a9a97cf1feddcd8b9906963893a648384aeec6374f1)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-62196179190494302a68d593b134fba295962b67974de8bff861b007ba9f9dec)
- [routes.route_direct_response](data-sources--route--reference--group-002.md#canonical-584c45695d1623f2c6198b272a75875291bfe763e53d0ba307888a33af25815a)
- [routes.route_redirect](data-sources--route--reference--group-002.md#canonical-cd0f5be797549d439537ec1c55d5cd9098a551adfdf30ac18b20d6d021b67053)
- [routes.service_policy](data-sources--route--reference--group-002.md#canonical-1a0cd68ab6678dc5ee8c0d95b09eb643ffc0d591cd4967766cdf58feecc75778)
- [routes.waf_exclusion_policy](data-sources--route--reference--group-002.md#canonical-9c2da572f6774afcd533806a7075bfa9703c4c1efdc6739bc72cfa6e2bb44373)
- [routes.waf_type](data-sources--route--reference--group-002.md#canonical-7362ef3ec8de94401ce424467749d400c77686ad5c2a78680790aca10481fc58)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-c2ec8eaf5c4fee182c5f4d3d30af7c79731051a3596796f547265902a7dcfd22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64201b64637a0d52460b1be9b71b3ad332dce0e7fb9f9a8597311465b6a39294"></a>

## routes.bot_defense_javascript_injection — routes.bot_defense_javascript_injection / 78306d924b2d / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- routes.bot_defense_javascript_injection

<a id="canonical-09391bd1fbb8fe42bb054c778c3313f8c93c41b21537e62ce4340ee0553963a6"></a>

Type: `"single"`. Computed.

Bot Defense Javascript Injection Configuration for inline bot defense deployments.

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

<a id="canonical-bcb49122b4edd2fe1339eac077695dca1d50309f941f7ad3d13500bb8c4190a6"></a>

## Direct properties — routes.bot_defense_javascript_injection / 78306d924b2d / 3

<a id="canonical-e1aafe698138042fee4080337f58aae31131ee10d2bdc56f329df4f0d14d892c"></a>

<a id="canonical-5a0a1b36b4abc2ce77cc4c0ebebd2715254d5733a05e8569bf101b97cdefc23b"></a>

## javascript_location property — routes.bot_defense_javascript_injection / 78306d924b2d / 4

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [javascript_tags](data-sources--route--reference--group-001.md#canonical-0e21685a46724fd12cd395b2c1c453787cfcd3cfc29bd632e2428507e4c09182): complete subsection reference.

<a id="canonical-dbb822c607dc01ad9829f837ff9bfc7cc67d2b155a8a7765ece8808ed1d1f4dc"></a>

## Next pages — routes.bot_defense_javascript_injection / 78306d924b2d / 5

- [routes.bot_defense_javascript_injection.javascript_tags](data-sources--route--reference--group-001.md#canonical-0e21685a46724fd12cd395b2c1c453787cfcd3cfc29bd632e2428507e4c09182)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-0e21685a46724fd12cd395b2c1c453787cfcd3cfc29bd632e2428507e4c09182"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c57b9b7602a1d7d931f108d4d85d0f9e0dae31c1dd4d8fbc5357b58868ab1ecf"></a>

## routes.bot_defense_javascript_injection.javascript_tags — routes.bot_defense_javascript_injection.javascript_tags / 64488c8070b6 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-c2ec8eaf5c4fee182c5f4d3d30af7c79731051a3596796f547265902a7dcfd22)
- routes.bot_defense_javascript_injection.javascript_tags

<a id="canonical-5d804b902de1eeb2b32236ee112c641a4a74fb0ec402c59fbf557a04b703653b"></a>

Type: `"list"`. Computed.

Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot
Javascript should be added first.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a922754ceffbd44be66da624271c2601865158667174a1d0bddd61fd6ba7509c"></a>

## Direct properties — routes.bot_defense_javascript_injection.javascript_tags / 64488c8070b6 / 3

<a id="canonical-35ffa6285c5c349baad59792af51da1c216bd5a74316489afa1d8de6f2495e9a"></a>

<a id="canonical-e9c33416cce61d7f7099e9fd04e733869cffed0ccde28a67ace3a185445d6bd0"></a>

## javascript_url property — routes.bot_defense_javascript_injection.javascript_tags / 64488c8070b6 / 4

Type: `"string"`. Computed.

Please enter the full URL (include domain and path), or relative path.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 2048,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [tag_attributes](data-sources--route--reference--group-001.md#canonical-9d92ebc13e984ac8d32b70c87dd6988dfc4771c2ed9a67bb3c8e51b5390c62f7): complete subsection reference.

<a id="canonical-85aef86e876da5e3ecd35f6dadaf19b9382fff5bb75c06ae832905786ad4d0ba"></a>

## Next pages — routes.bot_defense_javascript_injection.javascript_tags / 64488c8070b6 / 5

- [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes](data-sources--route--reference--group-001.md#canonical-9d92ebc13e984ac8d32b70c87dd6988dfc4771c2ed9a67bb3c8e51b5390c62f7)
- [routes.bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-c2ec8eaf5c4fee182c5f4d3d30af7c79731051a3596796f547265902a7dcfd22)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-9d92ebc13e984ac8d32b70c87dd6988dfc4771c2ed9a67bb3c8e51b5390c62f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e77ed9bd88325289709f0fb8c108b4f92d023b577483fbd3c8fb74793198a701"></a>

## routes.bot_defense_javascript_injection.javascript_tags.tag_attributes — routes.bot_defense_javascript_injection.javascript_tags.tag_attributes / baf241ab8957 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-c2ec8eaf5c4fee182c5f4d3d30af7c79731051a3596796f547265902a7dcfd22)
- [routes.bot_defense_javascript_injection.javascript_tags](data-sources--route--reference--group-001.md#canonical-0e21685a46724fd12cd395b2c1c453787cfcd3cfc29bd632e2428507e4c09182)
- routes.bot_defense_javascript_injection.javascript_tags.tag_attributes

<a id="canonical-defde4960a44f77efe68219039f8083b43bc498e07022e3c513ef928ef0097ae"></a>

Type: `"list"`. Computed.

Add the tag attributes you want to include in your Javascript tag.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-92ce653448282de56d9a4b4538f1e6d4d75e61f5c505b2d03836e1a06f62edc6"></a>

## Direct properties — routes.bot_defense_javascript_injection.javascript_tags.tag_attributes / baf241ab8957 / 3

<a id="canonical-7d0ec7ec23788834f2f76de9d391843c12f95db345e0800544f5a603baf0c82a"></a>

<a id="canonical-9401880118b0ac52a996f5796400ad36609501ce55250df1a62b854171a04af5"></a>

## javascript_tag property — routes.bot_defense_javascript_injection.javascript_tags.tag_attributes / baf241ab8957 / 4

Type: `"string"`. Computed.

\[Enum:
JS\_ATTR\_ID|JS\_ATTR\_CID|JS\_ATTR\_CN|JS\_ATTR\_API\_DOMAIN|JS\_ATTR\_API\_URL|JS\_ATTR\_API\_PATH|JS\_ATTR\_ASYNC|JS\_ATTR\_DEFER\]
Select from one of the predefined tag attributes. Possible values are \`JS\_ATTR\_ID\`,
\`JS\_ATTR\_CID\`, \`JS\_ATTR\_CN\`, \`JS\_ATTR\_API\_DOMAIN\`, \`JS\_ATTR\_API\_URL\`,
\`JS\_ATTR\_API\_PATH\`, \`JS\_ATTR\_ASYNC\`, \`JS\_ATTR\_DEFER\`. Defaults to \`JS\_ATTR\_ID\`.

Upstream description:

Select from one of the predefined tag attributes.

Receipt-pinned upstream constraints:

```json
{
  "default": "JS_ATTR_ID",
  "enum": [
    "JS_ATTR_ID",
    "JS_ATTR_CID",
    "JS_ATTR_CN",
    "JS_ATTR_API_DOMAIN",
    "JS_ATTR_API_URL",
    "JS_ATTR_API_PATH",
    "JS_ATTR_ASYNC",
    "JS_ATTR_DEFER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9ce296a2a69235243a52231e58004bb101fd447931c3e820dbcd50a8755fd31b"></a>

<a id="canonical-1bb0fb5dfbad961712d92fc1ffa67f172f61427571d7588bf4fe6f07d9edc44c"></a>

## tag_value property — routes.bot_defense_javascript_injection.javascript_tags.tag_attributes / baf241ab8957 / 5

Type: `"string"`. Computed.

Value. Add the tag attribute value.

Upstream description:

Add the tag attribute value.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-07cca36bc30785e715f16b4206f1c517821824f74eb84860d8a6b015a9ef4417"></a>

## Next pages — routes.bot_defense_javascript_injection.javascript_tags.tag_attributes / baf241ab8957 / 6

- [routes.bot_defense_javascript_injection.javascript_tags](data-sources--route--reference--group-001.md#canonical-0e21685a46724fd12cd395b2c1c453787cfcd3cfc29bd632e2428507e4c09182)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-73284df79d3594d5fbc5783b129a32c7d48daae43f13e9ab34714fb52843dea0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22e2b3de7a23fdf43354f09b5e4a32877aab82c2f7dc7761f735156c6ebeb622"></a>

## routes.inherited_bot_defense_javascript_injection — routes.inherited_bot_defense_javascript_injection / fdc96e167606 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- routes.inherited_bot_defense_javascript_injection

<a id="canonical-eb8314fdcbbb8e5f5af67b17368af644fdf9b0e98e4e5c35ee727f30e783f3da"></a>

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

<a id="canonical-b6c4747bd7168e1e6337f89eacafb609a290f3d17cc14a4f15fd9f335f0744e6"></a>

## Direct properties — routes.inherited_bot_defense_javascript_injection / fdc96e167606 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0f857719d8be3bbfd5c6c027153435c1787ea6fd2154da31e473158be381ff2e"></a>

## Next pages — routes.inherited_bot_defense_javascript_injection / fdc96e167606 / 4

- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-c80363e9c26bac22ab025c48db1ef4e2f2a7b14720d9f7657b6e1cf658c4d137"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-057ab073b0bbbaf5d4b9234fbb3bcd4982c28b68351d52a8c7a7f8e155d9b6a2"></a>

## routes.inherited_waf_exclusion — routes.inherited_waf_exclusion / 276fae82e42e / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- routes.inherited_waf_exclusion

<a id="canonical-02a735e0e99f2f89e29f63aaa1acbb78e98c1e1a219b4b44f4c11e6dcd65b359"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inherited waf exclusion.

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

<a id="canonical-ddd11aee39e62cb27baa18e46e5b282210f62da6ebbfe302352dfcca93e86677"></a>

## Direct properties — routes.inherited_waf_exclusion / 276fae82e42e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb8e69412e1878350758a765ccbb24e27c9930f47ae6eb4bc86b9f4c1fcb46ef"></a>

## Next pages — routes.inherited_waf_exclusion / 276fae82e42e / 4

- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-a2257940392ee77e0b7cc0101765684d7b06d8037bd3729204513dcbcb5c25f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee172395f6e7e2bb4e1a79eefceb963be1c7f7bca55764010b62c6ad56711d08"></a>

## routes.match — routes.match / 6ef77b729710 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- routes.match

<a id="canonical-077244628db60d731679da110ec40d95b9b27739ac105a98d12615b024bfc200"></a>

Type: `"list"`. Computed.

Match. Route match condition.

Upstream description:

Route match condition.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-c4e43e1424cfb2a8e19f3637ae44c79607b68aaad628fddbc0e529aa37cd4fd1"></a>

## Direct properties — routes.match / 6ef77b729710 / 3

- [headers](data-sources--route--reference--group-001.md#canonical-b9b98398209c01399fe55c202b312188b05ffd69144b3f8ff8ae39c1fa23db56): complete subsection reference.

<a id="canonical-aec1d0a9a245c795dd57c16745b19070248edd88d88e2fee1bdbc6a275816036"></a>

<a id="canonical-ef9d0a775e66e52683ea3b65cd5e1016d3654717e74ad204223965082ab8d0eb"></a>

## http_method property — routes.match / 6ef77b729710 / 4

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](data-sources--route--reference--group-001.md#canonical-8d4004f9ff66ea68687d9222b65358d3083cfc4c1a4f7c63432267e2452b89e8): complete subsection reference.

- [path](data-sources--route--reference--group-001.md#canonical-714b79af5881789ae22abfe0b4c3584e09d42259cfa0fba5ac98f4021e4f8835): complete subsection reference.

- [query_params](data-sources--route--reference--group-001.md#canonical-272f71fa352ea14858a3b2beeced0df8e255a590baf445d9718bb18bef1b262b): complete subsection reference.

<a id="canonical-542c209d17fd6c6f66180e6ac9b394cec737857675651ca0ffbc358ca6cb8a19"></a>

## Next pages — routes.match / 6ef77b729710 / 5

- [routes.match.headers](data-sources--route--reference--group-001.md#canonical-b9b98398209c01399fe55c202b312188b05ffd69144b3f8ff8ae39c1fa23db56)
- [routes.match.incoming_port](data-sources--route--reference--group-001.md#canonical-8d4004f9ff66ea68687d9222b65358d3083cfc4c1a4f7c63432267e2452b89e8)
- [routes.match.path](data-sources--route--reference--group-001.md#canonical-714b79af5881789ae22abfe0b4c3584e09d42259cfa0fba5ac98f4021e4f8835)
- [routes.match.query_params](data-sources--route--reference--group-001.md#canonical-272f71fa352ea14858a3b2beeced0df8e255a590baf445d9718bb18bef1b262b)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-b9b98398209c01399fe55c202b312188b05ffd69144b3f8ff8ae39c1fa23db56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44a01caf9a7617d85d31ac1b73572588d8d45fbc7b5ef0a54d2190c9676188b3"></a>

## routes.match.headers — routes.match.headers / 1f073bc9e50f / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.match](data-sources--route--reference--group-001.md#canonical-a2257940392ee77e0b7cc0101765684d7b06d8037bd3729204513dcbcb5c25f6)
- routes.match.headers

<a id="canonical-fbb0dda8a0fdb710d3ec3ea61ecba15fc242acedf7f77138defe03e2ba6f9329"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

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
    "minItems": 0,
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

<a id="canonical-323e80af1575903502915aa9b5bd4ee74151f9a2754f5ca5031fafb2a5c55f7e"></a>

## Direct properties — routes.match.headers / 1f073bc9e50f / 3

<a id="canonical-b4d6922f6a2285389fc3a0fbfafea45bf39cdfc669978697659eaf5666f5b4df"></a>

<a id="canonical-18cd14c61e59e25fc6844e9167c13e8e198d9f631693c5e66b4c76923b0e5db2"></a>

## exact property — routes.match.headers / 1f073bc9e50f / 4

Type: `"string"`. Computed.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-119d8418ec0e8cc31d6c03e234738282689e5d86ed2f2ba622b9acb32a256c8f"></a>

<a id="canonical-0d815121295861bac4da4e87522fef43ad3c9672c8f838f043123040c3b7b772"></a>

## invert_match property — routes.match.headers / 1f073bc9e50f / 5

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-0e1c774f1a9f2ccee48034b4a7016704914a675e4b39c285a2d5ddd0566507f0"></a>

<a id="canonical-60cac603d78298fe26430cc22edc0e2a9121ef6a2fae435ce0cbfca840063e75"></a>

## name property — routes.match.headers / 1f073bc9e50f / 6

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1cc49bf5be7b80bdde342721d6b38b7399b6ab301bce7da681d051ef4dfb977f"></a>

<a id="canonical-63d257c60a9eda500a65281328c66d8d1b41bf2e258d340a1904848c8a9196db"></a>

## presence property — routes.match.headers / 1f073bc9e50f / 7

Type: `"bool"`. Computed.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

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

<a id="canonical-eafeb058c4161cfe61fc7db743cdd6068afea0572851e198ab445b9ea79075e3"></a>

<a id="canonical-bde7f606ffa77d59b0e3dd2e42354c1af6dad911badf2b747c68a050b1881ad5"></a>

## regex property — routes.match.headers / 1f073bc9e50f / 8

Type: `"string"`. Computed.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-517fc0ad9b89e474a6f21b6c2283211c895e1bd2b0561c33197e7afa85aad029"></a>

## Next pages — routes.match.headers / 1f073bc9e50f / 9

- [routes.match](data-sources--route--reference--group-001.md#canonical-a2257940392ee77e0b7cc0101765684d7b06d8037bd3729204513dcbcb5c25f6)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-8d4004f9ff66ea68687d9222b65358d3083cfc4c1a4f7c63432267e2452b89e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-823d50f2e5666d9b1600049ed7de009e98a2551cc34ea0696505b407b4cf54cf"></a>

## routes.match.incoming_port — routes.match.incoming_port / 7f2dccb7abae / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.match](data-sources--route--reference--group-001.md#canonical-a2257940392ee77e0b7cc0101765684d7b06d8037bd3729204513dcbcb5c25f6)
- routes.match.incoming_port

<a id="canonical-c8b4a189c56fd076868d6cdaf799e621b5d696e4ddd93cf34cae325b4c6c2f9a"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-a84a927510d2d290edb74eee18869e513fa3db3fb2a854b7edf3ad5f3e06c3bf"></a>

## Direct properties — routes.match.incoming_port / 7f2dccb7abae / 3

- [no_port_match](data-sources--route--reference--group-001.md#canonical-b2c53f07a6176e3a485538814f174805efb9e9d0e52a348217318ebc5acb977c): complete subsection reference.

<a id="canonical-e2ca73b2b37634fbcdf319761295e11ce939e30c28e58a3747556075fbbe2f59"></a>

<a id="canonical-9c45184d78a33f8e52ec0cccd116e8fd1452069426eb6e3f18cc346bca7f4082"></a>

## port property — routes.match.incoming_port / 7f2dccb7abae / 4

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-9e7f6a8610d37e963e72605aafbd94b8793b0a8f559298b63bb948c90a83df12"></a>

<a id="canonical-db41ab020c169ccb8b1dd6dad382cb7c002a07c0239aa457f2fa480605eb87f5"></a>

## port_ranges property — routes.match.incoming_port / 7f2dccb7abae / 5

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-4c627407a3deb252cd211ac95e4cf74e7f59c8ee1bc6542c294b0c2dec4a517e"></a>

## Next pages — routes.match.incoming_port / 7f2dccb7abae / 6

- [routes.match.incoming_port.no_port_match](data-sources--route--reference--group-001.md#canonical-b2c53f07a6176e3a485538814f174805efb9e9d0e52a348217318ebc5acb977c)
- [routes.match](data-sources--route--reference--group-001.md#canonical-a2257940392ee77e0b7cc0101765684d7b06d8037bd3729204513dcbcb5c25f6)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-b2c53f07a6176e3a485538814f174805efb9e9d0e52a348217318ebc5acb977c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99c955588f4f30620d49017a5e1023dd092c8bdd995dcca2e33c2c2924cf4b7c"></a>

## routes.match.incoming_port.no_port_match — routes.match.incoming_port.no_port_match / 99fa3524692c / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.match](data-sources--route--reference--group-001.md#canonical-a2257940392ee77e0b7cc0101765684d7b06d8037bd3729204513dcbcb5c25f6)
- [routes.match.incoming_port](data-sources--route--reference--group-001.md#canonical-8d4004f9ff66ea68687d9222b65358d3083cfc4c1a4f7c63432267e2452b89e8)
- routes.match.incoming_port.no_port_match

<a id="canonical-675a550d4eb0dd3a83e836599282ce05d5408ab85692f4410d1de2ab6bec02d6"></a>

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

<a id="canonical-9bcb997887350de796ac6f77e4178ab02058d976093b3ad0d6cce7118dcfbc1d"></a>

## Direct properties — routes.match.incoming_port.no_port_match / 99fa3524692c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f467818e9e6efcb4566a5e09435ea8fdc5201198a2d93bb6e22c58a068683f9b"></a>

## Next pages — routes.match.incoming_port.no_port_match / 99fa3524692c / 4

- [routes.match.incoming_port](data-sources--route--reference--group-001.md#canonical-8d4004f9ff66ea68687d9222b65358d3083cfc4c1a4f7c63432267e2452b89e8)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-714b79af5881789ae22abfe0b4c3584e09d42259cfa0fba5ac98f4021e4f8835"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6384b69647c228a9b26b87d28df58e002830576ecb96786a99e362b97b3f36d0"></a>

## routes.match.path — routes.match.path / 1b2f5258e269 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.match](data-sources--route--reference--group-001.md#canonical-a2257940392ee77e0b7cc0101765684d7b06d8037bd3729204513dcbcb5c25f6)
- routes.match.path

<a id="canonical-c25a07b85b3d723f7897cde2c2c02844dd7e9785f183c139364aa2250243f463"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-771b03dec13aa55dd1b36cd55879100dc04826690e54571c3c05f35b447e03f8"></a>

## Direct properties — routes.match.path / 1b2f5258e269 / 3

<a id="canonical-0850806d06c5116d0105bd2f16ad73a01b238b90679a8e1b32a0afac4b1c43d1"></a>

<a id="canonical-03a4333a551a2ae8d9508d000ec610ae96214b3542f95313329e278b9a59e33e"></a>

## path property — routes.match.path / 1b2f5258e269 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-d1daed5070510236ee4c6211eb904e753060d2c4a4252fdaddd2107d810c3a03"></a>

<a id="canonical-11755f72fae09469221c4200e2bfbb591da9a5a3467b7703588e51f01f1cf534"></a>

## prefix property — routes.match.path / 1b2f5258e269 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-17be8e041ea57f9c3fcb759d1b683e491a836a39412e7392219add1b96bd3daa"></a>

<a id="canonical-17cace80c9e97dbf06494409439554b80f0316239dc80324ce2d83e873da69ef"></a>

## regex property — routes.match.path / 1b2f5258e269 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1e430f7eee087f20b5eadf7b87dffff94a05642f741367cd9dce1a63f023024b"></a>

## Next pages — routes.match.path / 1b2f5258e269 / 7

- [routes.match](data-sources--route--reference--group-001.md#canonical-a2257940392ee77e0b7cc0101765684d7b06d8037bd3729204513dcbcb5c25f6)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-272f71fa352ea14858a3b2beeced0df8e255a590baf445d9718bb18bef1b262b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae37925d1a2cb67cb6620031477973db66b96d27f968f6f46b9cb7b69f2634f0"></a>

## routes.match.query_params — routes.match.query_params / 262800636080 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.match](data-sources--route--reference--group-001.md#canonical-a2257940392ee77e0b7cc0101765684d7b06d8037bd3729204513dcbcb5c25f6)
- routes.match.query_params

<a id="canonical-628a149bf8c5ec16faf4ad26d932df706e701b03b58dccab34c3e9174dfe3287"></a>

Type: `"list"`. Computed.

Query Parameters. List of (key, value) query parameters.

Upstream description:

List of (key, value) query parameters.

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

<a id="canonical-d79b7b8584a8ad9ba240dca5a12e32aa556c7b21c6e001da56a18d14548d1f31"></a>

## Direct properties — routes.match.query_params / 262800636080 / 3

<a id="canonical-84ba224d392d660870b1648f0899eaa112777b7a05546dc811ab80449d3387e2"></a>

<a id="canonical-fd8f6123f3d4206ec910bc9ea42c0bc94a3cbb461be9686f0050c07475da7b23"></a>

## exact property — routes.match.query_params / 262800636080 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\] Exact match value for the query parameter key.

Upstream description:

Exclusive with \[regex\] Exact match value for the query parameter key.

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

<a id="canonical-f06b901c7704bca4590b89448304679637d3c2b6d8cfcaca2f9f7f8d6617c6e4"></a>

<a id="canonical-98f78f893fe1d5e62b11338fb1a9693d9753bb692dc965e934e866a14599dc41"></a>

## key property — routes.match.query_params / 262800636080 / 5

Type: `"string"`. Computed.

Query parameter key In the above example, assignee\_username is the key.

Upstream description:

Query parameter key In the above example, assignee\_username is the key.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-f88638e274a9b5017df4d4d1043e140eb9403bc283016811d75b89bcf2a29bab"></a>

<a id="canonical-885d07918b10aba4428d9ebc2389d975a81514188dcb46bc30e23c0d43c900d2"></a>

## regex property — routes.match.query_params / 262800636080 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\] Regex match value for the query parameter key.

Upstream description:

Exclusive with \[exact\] Regex match value for the query parameter key.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2a8a462f6ea2a469fde167bf3d7e68c724a66846330661bd1b3e11528cc47fce"></a>

## Next pages — routes.match.query_params / 262800636080 / 7

- [routes.match](data-sources--route--reference--group-001.md#canonical-a2257940392ee77e0b7cc0101765684d7b06d8037bd3729204513dcbcb5c25f6)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-fe0cf0995676155ec1ccb3ae9747442486d275b7596e31e7404a8cc9ecffcdaa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5f2684974a492f2d06303bc7798f03330ce6cd42de90c8e23f44d00c52d1529"></a>

## routes.request_cookies_to_add — routes.request_cookies_to_add / 8abe522a70ed / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- routes.request_cookies_to_add

<a id="canonical-6b646bfced6ec70f3524aadf8b9c575ab4d08ca820edf375b2015049c18cafe8"></a>

Type: `"list"`. Computed.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-fcbbb48a45a29d217d011c16194ad938cfdb39040708670c62cf4305178371ad"></a>

## Direct properties — routes.request_cookies_to_add / 8abe522a70ed / 3

<a id="canonical-4651daa0bf640582c7e90df1b1f7a85fb9f432f0f6b94e02d5618c85fac1af10"></a>

<a id="canonical-391977fa0f01270177fcc19d69448f289c8ac98bc2ef8f8a09bda08a60109516"></a>

## name property — routes.request_cookies_to_add / 8abe522a70ed / 4

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-97819d859a9b2dfbfdf7150eac7a71d9f52882e9f7859a89216f06ebdfd7857e"></a>

<a id="canonical-f1f2131e18d534d78ef6ce1bde5de86ab3c23492bcf58b418d74955f24e91104"></a>

## overwrite property — routes.request_cookies_to_add / 8abe522a70ed / 5

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

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

- [secret_value](data-sources--route--reference--group-001.md#canonical-e26ed9a0f2d042b0a961fe3d59421f7675f6af542851c9a7a306aa58d02c8918): complete subsection reference.

<a id="canonical-0a24905c0176b0126290a481f40f8730d66718c40dd1172b92570b8d2dd03a5d"></a>

<a id="canonical-d8d44758924d6ceb4513f84149d4d399f57a69834a07bea77dac7ea4d13cfae8"></a>

## value property — routes.request_cookies_to_add / 8abe522a70ed / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-67838fafd8ad6f50505d8003ac55449afea01b8f4687ed05b9839b1951467a4c"></a>

## Next pages — routes.request_cookies_to_add / 8abe522a70ed / 7

- [routes.request_cookies_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-e26ed9a0f2d042b0a961fe3d59421f7675f6af542851c9a7a306aa58d02c8918)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-e26ed9a0f2d042b0a961fe3d59421f7675f6af542851c9a7a306aa58d02c8918"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9b12be44067fde0ec39e1a72b06b9875bf16039f639ea3c445b50c5d1816635"></a>

## routes.request_cookies_to_add.secret_value — routes.request_cookies_to_add.secret_value / 794c13159d7d / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.request_cookies_to_add](data-sources--route--reference--group-001.md#canonical-fe0cf0995676155ec1ccb3ae9747442486d275b7596e31e7404a8cc9ecffcdaa)
- routes.request_cookies_to_add.secret_value

<a id="canonical-f37aa9c1d3807d4d06bdc1fa731c81b091a035f65f3905a343f25eb9f29f52b1"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-4f467ab05cb932b22c8d054ff5b0f8a5b351b8d02ac50dda2f72ad4472739e06"></a>

## Direct properties — routes.request_cookies_to_add.secret_value / 794c13159d7d / 3

- [blindfold_secret_info](data-sources--route--reference--group-001.md#canonical-0bd4680c97f48adfc7690f68ddd9a1dbd035585b5051b2b0ed3b2d53a4f90794): complete subsection reference.

- [clear_secret_info](data-sources--route--reference--group-001.md#canonical-e558c734fb869b99306c20284e2f860f4e1bb0bfe1d83ba987d966c58a80cdb6): complete subsection reference.

<a id="canonical-7d8957b6359aa42b2f06f567924d2bf16733bf3781f37c139882373d87bb370a"></a>

## Next pages — routes.request_cookies_to_add.secret_value / 794c13159d7d / 4

- [routes.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--route--reference--group-001.md#canonical-0bd4680c97f48adfc7690f68ddd9a1dbd035585b5051b2b0ed3b2d53a4f90794)
- [routes.request_cookies_to_add.secret_value.clear_secret_info](data-sources--route--reference--group-001.md#canonical-e558c734fb869b99306c20284e2f860f4e1bb0bfe1d83ba987d966c58a80cdb6)
- [routes.request_cookies_to_add](data-sources--route--reference--group-001.md#canonical-fe0cf0995676155ec1ccb3ae9747442486d275b7596e31e7404a8cc9ecffcdaa)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-0bd4680c97f48adfc7690f68ddd9a1dbd035585b5051b2b0ed3b2d53a4f90794"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edaba16bc068241b7271839c3d20be0755f405a62d23a3fa84db8c9e51a162ab"></a>

## routes.request_cookies_to_add.secret_value.blindfold_secret_info — routes.request_cookies_to_add.secret_value.blindfold_secret_info / c392d0bb6ac5 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.request_cookies_to_add](data-sources--route--reference--group-001.md#canonical-fe0cf0995676155ec1ccb3ae9747442486d275b7596e31e7404a8cc9ecffcdaa)
- [routes.request_cookies_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-e26ed9a0f2d042b0a961fe3d59421f7675f6af542851c9a7a306aa58d02c8918)
- routes.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-a3557aac408c18e23d052f53c8242b4c138a66b92674399bb3fbc1f580970b3b"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-f0ca8176195b66e1654773e5fd322e6d012127d78b6ea6f838b94bef08c95249"></a>

## Direct properties — routes.request_cookies_to_add.secret_value.blindfold_secret_info / c392d0bb6ac5 / 3

<a id="canonical-4c0311dc1746d41ec21bdcdd88786f050e37df83982cf3d2d1fc277b9d49b638"></a>

<a id="canonical-c79a364cddc546551ed333e6ec7f30ec783fa0c7513e198d6a7a8ab3d4a9177a"></a>

## decryption_provider property — routes.request_cookies_to_add.secret_value.blindfold_secret_info / c392d0bb6ac5 / 4

Type: `"string"`. Computed.

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

<a id="canonical-d84acbecaad92a463a4fdbde891f5a84cfa3c2dda7d0c28c1cd28e3ae86f782d"></a>

<a id="canonical-213f9fa1e8346b68b66c3b4ae453628f09184feee3fb3e7f9eddd98d8b9f511c"></a>

## location property — routes.request_cookies_to_add.secret_value.blindfold_secret_info / c392d0bb6ac5 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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

<a id="canonical-4ef2b9a9e2a9dd5b7d14594c17e83dcc1a78f815adac8505c0a50852dc4c7895"></a>

<a id="canonical-16f9de1b11a063ae48e781a306794482339063ed154f054889af06cadfcdf46e"></a>

## store_provider property — routes.request_cookies_to_add.secret_value.blindfold_secret_info / c392d0bb6ac5 / 6

Type: `"string"`. Computed.

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

<a id="canonical-8b51195ef5e393bc6f3c9bd6271d06ceb41f21bdbc18bb190932f097366eb3c6"></a>

## Next pages — routes.request_cookies_to_add.secret_value.blindfold_secret_info / c392d0bb6ac5 / 7

- [routes.request_cookies_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-e26ed9a0f2d042b0a961fe3d59421f7675f6af542851c9a7a306aa58d02c8918)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-e558c734fb869b99306c20284e2f860f4e1bb0bfe1d83ba987d966c58a80cdb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ec0752ddfe1232d7b4cbddc9737a447588a68b44fef8db1750a5a0bad05d081"></a>

## routes.request_cookies_to_add.secret_value.clear_secret_info — routes.request_cookies_to_add.secret_value.clear_secret_info / 4e31691dbce5 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.request_cookies_to_add](data-sources--route--reference--group-001.md#canonical-fe0cf0995676155ec1ccb3ae9747442486d275b7596e31e7404a8cc9ecffcdaa)
- [routes.request_cookies_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-e26ed9a0f2d042b0a961fe3d59421f7675f6af542851c9a7a306aa58d02c8918)
- routes.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-6d2a18d2a12a9b7c7d4ae6b4742f41ba83f211a436216d5a96033147feda0f12"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-db40a28f5974cba54c315d560c4325aa795935046e0270e24dff094a107ab05a"></a>

## Direct properties — routes.request_cookies_to_add.secret_value.clear_secret_info / 4e31691dbce5 / 3

<a id="canonical-160b6f9d9eda68c3fe51628a65c19dcb3b0026ca3d83b91be4f95e52d7fa623c"></a>

<a id="canonical-997dc062f7000d5d46c5e37dbcb581c9113265eef593cd263524a6227161d8d3"></a>

## provider_ref property — routes.request_cookies_to_add.secret_value.clear_secret_info / 4e31691dbce5 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2ebf0c0a449193ca02f955a3a6e9ae3bcda14a931e1ed6824eef4699899cc0cf"></a>

<a id="canonical-a587f26edd76015b8685ead3312817064769fd27372cead00f136cfc618300a8"></a>

## url property — routes.request_cookies_to_add.secret_value.clear_secret_info / 4e31691dbce5 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-8710a16083546bfd7954e92a0e4f7fdd9e3767d524e304fd21c9ec4f3ebf73aa"></a>

## Next pages — routes.request_cookies_to_add.secret_value.clear_secret_info / 4e31691dbce5 / 6

- [routes.request_cookies_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-e26ed9a0f2d042b0a961fe3d59421f7675f6af542851c9a7a306aa58d02c8918)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-0865f1acc903ad9680c665112e61b6d96f02f8e9a7d9cbeed841654e49236b4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34aaf11521bff3b779cd2c403f72e92c7913db506a5fdb386f314823809317e5"></a>

## routes.request_headers_to_add — routes.request_headers_to_add / 834f94905bec / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- routes.request_headers_to_add

<a id="canonical-c32df7f7b50a14b673ddd6e89cbd58bf258a18b0884b384e785a15ddb6a44aa0"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP requests being sent towards upstream. Headers
specified at this level are applied before headers from the enclosing VirtualHost object level.

Upstream description:

Headers are key-value pairs to be added to HTTP requests being sent towards upstream. Headers
specified at this level are applied before headers from the enclosing VirtualHost object level.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-eaecb5d98a7ed1abb271e4f6f413358c0ed7ef477989aafd1945bd3f6370f34c"></a>

## Direct properties — routes.request_headers_to_add / 834f94905bec / 3

<a id="canonical-e87d571fa2c40bbcdfa524b8840aed7623f3a5f6d588e6d28a87b17d7862e171"></a>

<a id="canonical-1d0cb00cb398660b6fd4ea42a91c506b39138541185ae0fb011b50509d43dc2e"></a>

## append property — routes.request_headers_to_add / 834f94905bec / 4

Type: `"bool"`. Computed.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

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

<a id="canonical-4fe472de5809cef236cd681d89e822830e5f810ce35d19fafdc410c1b0c53c0e"></a>

<a id="canonical-a5fcc158e8acde357c957a394cf25f4386733a69aa95a379cfc52de8055b5a32"></a>

## name property — routes.request_headers_to_add / 834f94905bec / 5

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](data-sources--route--reference--group-001.md#canonical-2500909fa0fe05bc43e5601fd6033df14a4254546f6489da2d390b65dfc8a100): complete subsection reference.

<a id="canonical-807087a6c11435229ce842cb8c146c820bc456bfbc5856ad7ab1be3276b7751d"></a>

<a id="canonical-f487d414d22cfa9c9dc0eb76f4342b60cba3e56f4b53cd41d09f6c49a2af47ab"></a>

## value property — routes.request_headers_to_add / 834f94905bec / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-baa6d9f21bd07fc6fb3aad22585ee8eeec46dedd548a2467e159d3e4a7c823f0"></a>

## Next pages — routes.request_headers_to_add / 834f94905bec / 7

- [routes.request_headers_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-2500909fa0fe05bc43e5601fd6033df14a4254546f6489da2d390b65dfc8a100)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-2500909fa0fe05bc43e5601fd6033df14a4254546f6489da2d390b65dfc8a100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6209ddf5ad1e5d27fc22fc4b66ba31756d3fb4df1b4cd87598e1f01b334a71e"></a>

## routes.request_headers_to_add.secret_value — routes.request_headers_to_add.secret_value / 5a0c581c1383 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.request_headers_to_add](data-sources--route--reference--group-001.md#canonical-0865f1acc903ad9680c665112e61b6d96f02f8e9a7d9cbeed841654e49236b4f)
- routes.request_headers_to_add.secret_value

<a id="canonical-f8bc640bb08b44708e49b065b70b16c2ac6aca5d8a387c4edbf4f1d38d5cfbb9"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-30b35ecc44c876e08810d9d958fb150df1df7a08f479335bcae40a643fdd7ca7"></a>

## Direct properties — routes.request_headers_to_add.secret_value / 5a0c581c1383 / 3

- [blindfold_secret_info](data-sources--route--reference--group-001.md#canonical-b918a94442993ad57e9cb98eb2ddb292b1dafc9c80c40515ebda4845dd8d97bd): complete subsection reference.

- [clear_secret_info](data-sources--route--reference--group-001.md#canonical-cd56274b6d575a3c814f1e6169bd833935915246d450a08b3aa67dbb98482603): complete subsection reference.

<a id="canonical-e39d1f2ff907010352fbf4c1b3bf4efe58c76a8a6426ef7695a32917a1cba6ef"></a>

## Next pages — routes.request_headers_to_add.secret_value / 5a0c581c1383 / 4

- [routes.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--route--reference--group-001.md#canonical-b918a94442993ad57e9cb98eb2ddb292b1dafc9c80c40515ebda4845dd8d97bd)
- [routes.request_headers_to_add.secret_value.clear_secret_info](data-sources--route--reference--group-001.md#canonical-cd56274b6d575a3c814f1e6169bd833935915246d450a08b3aa67dbb98482603)
- [routes.request_headers_to_add](data-sources--route--reference--group-001.md#canonical-0865f1acc903ad9680c665112e61b6d96f02f8e9a7d9cbeed841654e49236b4f)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-b918a94442993ad57e9cb98eb2ddb292b1dafc9c80c40515ebda4845dd8d97bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a424b31058e8230d452554e364f3a9015c822b3378cc372083f8f5b215373fbc"></a>

## routes.request_headers_to_add.secret_value.blindfold_secret_info — routes.request_headers_to_add.secret_value.blindfold_secret_info / 5441afdb6f3b / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.request_headers_to_add](data-sources--route--reference--group-001.md#canonical-0865f1acc903ad9680c665112e61b6d96f02f8e9a7d9cbeed841654e49236b4f)
- [routes.request_headers_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-2500909fa0fe05bc43e5601fd6033df14a4254546f6489da2d390b65dfc8a100)
- routes.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-40662a722c36eb16e0f777151940accfd4d5a4828101ab51cc45171b0751598c"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-e87ed3c937b9c819d5cfff316c47a84f23dc2318a7147ea1d9325bd150f16c29"></a>

## Direct properties — routes.request_headers_to_add.secret_value.blindfold_secret_info / 5441afdb6f3b / 3

<a id="canonical-ed2e47215b0f6c7db924d34115882d39023c575f2a7fb3202095b788732fe401"></a>

<a id="canonical-26dfd0d3e3ca1107104d3c3826b11cdcb8904aa51e753c887be58fcff264d956"></a>

## decryption_provider property — routes.request_headers_to_add.secret_value.blindfold_secret_info / 5441afdb6f3b / 4

Type: `"string"`. Computed.

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

<a id="canonical-5edea939ec828dd53430adb7d5d68fcd24ea7a940dfc3b0da1cfc4331035c6f1"></a>

<a id="canonical-b226218824902860a6dc002486aadf5fd01823fb697907f86f3a5cb08d982db8"></a>

## location property — routes.request_headers_to_add.secret_value.blindfold_secret_info / 5441afdb6f3b / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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

<a id="canonical-2d2457d9a109064ee66c5dfb5df6fd8d10ffc49f7f2e481e87c15877f577a076"></a>

<a id="canonical-b4dcffcdd38b907ec0487dab14f05f8b6830458665f8a172418531d639f3820e"></a>

## store_provider property — routes.request_headers_to_add.secret_value.blindfold_secret_info / 5441afdb6f3b / 6

Type: `"string"`. Computed.

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

<a id="canonical-228ddcd18266de520f99c7af4147a19f0182ce781b2daa1e3821a8431250e37d"></a>

## Next pages — routes.request_headers_to_add.secret_value.blindfold_secret_info / 5441afdb6f3b / 7

- [routes.request_headers_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-2500909fa0fe05bc43e5601fd6033df14a4254546f6489da2d390b65dfc8a100)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-cd56274b6d575a3c814f1e6169bd833935915246d450a08b3aa67dbb98482603"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29fa3c7c6e626a3712ef2bd9cfbb7153be59498b0dc1b2309dcbab4393ee98c7"></a>

## routes.request_headers_to_add.secret_value.clear_secret_info — routes.request_headers_to_add.secret_value.clear_secret_info / 5b7b88072bca / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.request_headers_to_add](data-sources--route--reference--group-001.md#canonical-0865f1acc903ad9680c665112e61b6d96f02f8e9a7d9cbeed841654e49236b4f)
- [routes.request_headers_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-2500909fa0fe05bc43e5601fd6033df14a4254546f6489da2d390b65dfc8a100)
- routes.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-e96f75ba9614f86a6d7b5ea52a054cce4e2dc64bcadb48d212da9f5da4ef2945"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-bca84a32e905f0d84363704d222ee71c7015d6a2159d6c5ffee89d7f339d983c"></a>

## Direct properties — routes.request_headers_to_add.secret_value.clear_secret_info / 5b7b88072bca / 3

<a id="canonical-e88dd9d865ca7a155c97350ab142dd52f58576e0d26c25e0d12780757b0eb652"></a>

<a id="canonical-2b0edeeef9cf83a05416c1e4c94dac98f121b5051aac44b84a65709a2aa5f40b"></a>

## provider_ref property — routes.request_headers_to_add.secret_value.clear_secret_info / 5b7b88072bca / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-cc889e67303577098be46d477b2db247eb07e1192d83ea08636ad86213f4f372"></a>

<a id="canonical-bbe422a7d511e8b03f2a2cbfadc488117a8d298ddd7076623921d3b8849840f9"></a>

## url property — routes.request_headers_to_add.secret_value.clear_secret_info / 5b7b88072bca / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-f46baa002f085e93c396fb05f156344207257335af987c217b96ec9f14c6e0d5"></a>

## Next pages — routes.request_headers_to_add.secret_value.clear_secret_info / 5b7b88072bca / 6

- [routes.request_headers_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-2500909fa0fe05bc43e5601fd6033df14a4254546f6489da2d390b65dfc8a100)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1c7720c47de5ecf0b09a7b52c95403a24a8e5ae7a644a2f6cfe0bee32ad248c"></a>

## routes.response_cookies_to_add — routes.response_cookies_to_add / a1d48781ac16 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- routes.response_cookies_to_add

<a id="canonical-e22dc8ccc9b9c902c4b8b67f74b9347779d175abe41249a2488cec6798dfa62f"></a>

Type: `"list"`. Computed.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4c86b743437c182aa37879db05470d31dffa175b50006f3392dc21b99c64131e"></a>

## Direct properties — routes.response_cookies_to_add / a1d48781ac16 / 3

<a id="canonical-d02a36301ecebc16bc396052e3d3e7627066a7212f816380c2fcdc39cb2544e2"></a>

<a id="canonical-8947f38dbfe2aa7576aaf8df5755f6f511ff3797ab7b6c82da9c33ed2ca73ee7"></a>

## add_domain property — routes.response_cookies_to_add / a1d48781ac16 / 4

Type: `"string"`. Computed.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

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

<a id="canonical-e9ee6b87f0f77eb2d141156bf5e40d31dadbd67a692636e175bf39c332a905cc"></a>

<a id="canonical-4dec2937da32e419a4fde051c3552d5c38542a379b954efbd29d8f86fe0468fb"></a>

## add_expiry property — routes.response_cookies_to_add / a1d48781ac16 / 5

Type: `"string"`. Computed.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_httponly](data-sources--route--reference--group-001.md#canonical-62944db3b41489a459db2a292fce0d1213b7ae924c3b1db0fb16711b55d23da3): complete subsection reference.

- [add_partitioned](data-sources--route--reference--group-001.md#canonical-d956414a3b4633f1b7df831936b561a36024092aae03894e732a61e938915572): complete subsection reference.

<a id="canonical-3c498016660620778bd971ee2f6c384c461e43c5fdd79ac9ede498eb83994966"></a>

<a id="canonical-0d52f044c4cd1809ecd0ba6decf33ff593e7dd80823ba105a950d06d27013584"></a>

## add_path property — routes.response_cookies_to_add / a1d48781ac16 / 6

Type: `"string"`. Computed.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_secure](data-sources--route--reference--group-001.md#canonical-8271ab2b55172df90dad0e472293afac144f6ed718bed7e9bc879a30873cb0f5): complete subsection reference.

- [ignore_domain](data-sources--route--reference--group-001.md#canonical-d264de6a0d75c25928a13439c8f5420da839cd932b59f9f334531612cd781981): complete subsection reference.

- [ignore_expiry](data-sources--route--reference--group-001.md#canonical-28dd1593a59c24179c35397371543d9b4247ac3ed7e2ca2b7c987f4a5dd5f779): complete subsection reference.

- [ignore_httponly](data-sources--route--reference--group-001.md#canonical-07296ee33d006e2a6c774af261641078621866e3dba34060459375635f057ac7): complete subsection reference.

- [ignore_max_age](data-sources--route--reference--group-001.md#canonical-e3e68f08f9a68d0e962ff51851df5a3983255fb5ca84b52bd56bcd2b4375527d): complete subsection reference.

- [ignore_partitioned](data-sources--route--reference--group-001.md#canonical-ae2646a1c13fa21606f7b49638403cf9d8e457bdb02d5e10f37886cb1b78ddb7): complete subsection reference.

- [ignore_path](data-sources--route--reference--group-001.md#canonical-cc956b4c4e0b144773de65af99bca28c31b0ededf986260e651df1d948009e25): complete subsection reference.

- [ignore_samesite](data-sources--route--reference--group-001.md#canonical-6a9d642dbb4e762ecbe49828f9626f4a3b56d536f071ae46f6cae1079e7e0093): complete subsection reference.

- [ignore_secure](data-sources--route--reference--group-002.md#canonical-7236da4e5bf425f37eb9eef5267a595f88578808686781f9887e4578ab6432a1): complete subsection reference.

- [ignore_value](data-sources--route--reference--group-002.md#canonical-fff8a5e49ed8efb22f1fc98236ee523220c3c131d936ffa4b0fd6dd8c789e3a8): complete subsection reference.

<a id="canonical-f6a230e162fba652d18bed0e0f216f1c97ebd2a5861d9735b739e8ccf99dc3b3"></a>

<a id="canonical-475f148d4c4d2a77f1876c9c117feefb5f4865529794c1a8f095c582ec6c3dbd"></a>

## max_age_value property — routes.response_cookies_to_add / a1d48781ac16 / 7

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-df7e29343f740a361c0a6b08a212613413d43c7fa7185683daa0d4344fefdbd9"></a>

<a id="canonical-abfdeeddffc149912b5af8f473993da36579ff710e823c748a75a55e9d4b807c"></a>

## name property — routes.response_cookies_to_add / a1d48781ac16 / 8

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-54c39677f9fed452d6892daa91fcd73f2fe3e65ad3fda3c5916b8cdae928ae78"></a>

<a id="canonical-57737f1fc74632a64c1719096b739ff9367c80abe507121d91202d1a339b3ad9"></a>

## overwrite property — routes.response_cookies_to_add / a1d48781ac16 / 9

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

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

- [samesite_lax](data-sources--route--reference--group-002.md#canonical-614b757e087dcf8be7a4ceea25b36bbd9994d024428dde53209c9ea2d36da151): complete subsection reference.

- [samesite_none](data-sources--route--reference--group-002.md#canonical-3f4e20aafb616cc4111fabe8f0d882168c4c98289ec389182d3e4066c305e619): complete subsection reference.

- [samesite_strict](data-sources--route--reference--group-002.md#canonical-039de0f54941490413cf33783629949605a5efd5ef3bd58ec107a8ea45c00b50): complete subsection reference.

- [secret_value](data-sources--route--reference--group-002.md#canonical-3e0e6266d47e6f5eb9a81e7ad5a9315751379d8e89514b47d9a01bf2081f2faa): complete subsection reference.

<a id="canonical-b2bd1c31052bcb661111797f9694e921d27378ddd34184435d1e28188aae8ca5"></a>

<a id="canonical-2dfd8b75b2caeebf1d7d3d4326bad9f3d3c110229e7aa457687b074ba1995cf4"></a>

## value property — routes.response_cookies_to_add / a1d48781ac16 / 10

Type: `"string"`. Computed.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-8f754856651103dd4339a48069d74aa0f3c286b8b054f63d21df3c664f68445f"></a>

## Next pages — routes.response_cookies_to_add / a1d48781ac16 / 11

- [routes.response_cookies_to_add.add_httponly](data-sources--route--reference--group-001.md#canonical-62944db3b41489a459db2a292fce0d1213b7ae924c3b1db0fb16711b55d23da3)
- [routes.response_cookies_to_add.add_partitioned](data-sources--route--reference--group-001.md#canonical-d956414a3b4633f1b7df831936b561a36024092aae03894e732a61e938915572)
- [routes.response_cookies_to_add.add_secure](data-sources--route--reference--group-001.md#canonical-8271ab2b55172df90dad0e472293afac144f6ed718bed7e9bc879a30873cb0f5)
- [routes.response_cookies_to_add.ignore_domain](data-sources--route--reference--group-001.md#canonical-d264de6a0d75c25928a13439c8f5420da839cd932b59f9f334531612cd781981)
- [routes.response_cookies_to_add.ignore_expiry](data-sources--route--reference--group-001.md#canonical-28dd1593a59c24179c35397371543d9b4247ac3ed7e2ca2b7c987f4a5dd5f779)
- [routes.response_cookies_to_add.ignore_httponly](data-sources--route--reference--group-001.md#canonical-07296ee33d006e2a6c774af261641078621866e3dba34060459375635f057ac7)
- [routes.response_cookies_to_add.ignore_max_age](data-sources--route--reference--group-001.md#canonical-e3e68f08f9a68d0e962ff51851df5a3983255fb5ca84b52bd56bcd2b4375527d)
- [routes.response_cookies_to_add.ignore_partitioned](data-sources--route--reference--group-001.md#canonical-ae2646a1c13fa21606f7b49638403cf9d8e457bdb02d5e10f37886cb1b78ddb7)
- [routes.response_cookies_to_add.ignore_path](data-sources--route--reference--group-001.md#canonical-cc956b4c4e0b144773de65af99bca28c31b0ededf986260e651df1d948009e25)
- [routes.response_cookies_to_add.ignore_samesite](data-sources--route--reference--group-001.md#canonical-6a9d642dbb4e762ecbe49828f9626f4a3b56d536f071ae46f6cae1079e7e0093)
- [routes.response_cookies_to_add.ignore_secure](data-sources--route--reference--group-002.md#canonical-7236da4e5bf425f37eb9eef5267a595f88578808686781f9887e4578ab6432a1)
- [routes.response_cookies_to_add.ignore_value](data-sources--route--reference--group-002.md#canonical-fff8a5e49ed8efb22f1fc98236ee523220c3c131d936ffa4b0fd6dd8c789e3a8)
- [routes.response_cookies_to_add.samesite_lax](data-sources--route--reference--group-002.md#canonical-614b757e087dcf8be7a4ceea25b36bbd9994d024428dde53209c9ea2d36da151)
- [routes.response_cookies_to_add.samesite_none](data-sources--route--reference--group-002.md#canonical-3f4e20aafb616cc4111fabe8f0d882168c4c98289ec389182d3e4066c305e619)
- [routes.response_cookies_to_add.samesite_strict](data-sources--route--reference--group-002.md#canonical-039de0f54941490413cf33783629949605a5efd5ef3bd58ec107a8ea45c00b50)
- [routes.response_cookies_to_add.secret_value](data-sources--route--reference--group-002.md#canonical-3e0e6266d47e6f5eb9a81e7ad5a9315751379d8e89514b47d9a01bf2081f2faa)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-62944db3b41489a459db2a292fce0d1213b7ae924c3b1db0fb16711b55d23da3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a32a42f850800f7254c6e16020fa352d426a007fe2c7b39bd61fe17a3bb14dc0"></a>

## routes.response_cookies_to_add.add_httponly — routes.response_cookies_to_add.add_httponly / eb446ada9f80 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- routes.response_cookies_to_add.add_httponly

<a id="canonical-7378748ecb2dc85b0c1be817dd9b457393670691e2b49a00cc4c9f10c8753548"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add httponly.

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

<a id="canonical-593d0be6c8a9e80201068c3b85c03f2c80cfb73e61a82f50ff513e2f7c2b82ce"></a>

## Direct properties — routes.response_cookies_to_add.add_httponly / eb446ada9f80 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b4d398790662f7cb78e98786cbd8435f531ff477e74d0a24c2964cfb5fbdf56c"></a>

## Next pages — routes.response_cookies_to_add.add_httponly / eb446ada9f80 / 4

- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-d956414a3b4633f1b7df831936b561a36024092aae03894e732a61e938915572"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2f7de623875360bd6673b5525a5b623e4edc89dd013493fc67b1969dc92ad1e"></a>

## routes.response_cookies_to_add.add_partitioned — routes.response_cookies_to_add.add_partitioned / 7f2e5977ce71 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- routes.response_cookies_to_add.add_partitioned

<a id="canonical-56fd5ae4c35d4c05e4fa887a955a7d344aa770bf61a0babb9368f85678a0f998"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add partitioned.

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

<a id="canonical-1f8080241bed34abbd35540f4f2c1e8d131d563cac51175dfc78f74735bdd7d2"></a>

## Direct properties — routes.response_cookies_to_add.add_partitioned / 7f2e5977ce71 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-16324452c2fbd4872ddf4b4d9bfcae2e067610b5d7a2d1c80bfb8623e4e22275"></a>

## Next pages — routes.response_cookies_to_add.add_partitioned / 7f2e5977ce71 / 4

- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-8271ab2b55172df90dad0e472293afac144f6ed718bed7e9bc879a30873cb0f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16de4d1cf7b2bd4bc31def53a70e63d931a4dc9c4308eaf8079be2d42d55db2f"></a>

## routes.response_cookies_to_add.add_secure — routes.response_cookies_to_add.add_secure / cbee8bd40b06 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- routes.response_cookies_to_add.add_secure

<a id="canonical-111c77178d0d671b56663efa3e22b071bf4ddc39c0d78f4ce80199a26c9680ff"></a>

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

<a id="canonical-94c4a194484a74b1422d0a56fbfb65f340f583e3f0def2637687a455d2dec5e2"></a>

## Direct properties — routes.response_cookies_to_add.add_secure / cbee8bd40b06 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6f0d0cccd78d4b7b3c7474995978444a02fd5849a9c6c060152726401e1d394f"></a>

## Next pages — routes.response_cookies_to_add.add_secure / cbee8bd40b06 / 4

- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-d264de6a0d75c25928a13439c8f5420da839cd932b59f9f334531612cd781981"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9020bf9dd7cb29703b12f993fdb60ea1f0a9f2ffb0b033e69bf4a10fd7b4e4f0"></a>

## routes.response_cookies_to_add.ignore_domain — routes.response_cookies_to_add.ignore_domain / 17da6991be70 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- routes.response_cookies_to_add.ignore_domain

<a id="canonical-cd27c8344681e1a0f7e273ebe9a8fbfa106e2d19727afa74c5fec2b931199439"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore domain.

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

<a id="canonical-d2b747891c0ec8998f41163f6dfa62126b22601c9edc6775e53c898d29b580f6"></a>

## Direct properties — routes.response_cookies_to_add.ignore_domain / 17da6991be70 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-923c1ddb19ba7e8a9500b30e5389a821da2f8392e1b5705e642778462878aa83"></a>

## Next pages — routes.response_cookies_to_add.ignore_domain / 17da6991be70 / 4

- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-28dd1593a59c24179c35397371543d9b4247ac3ed7e2ca2b7c987f4a5dd5f779"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c56b97a7b15fb0ff6c5abf67ada39870d910247dc6fc82ad340fe4f33c74ee3"></a>

## routes.response_cookies_to_add.ignore_expiry — routes.response_cookies_to_add.ignore_expiry / 6303e8fd2ce6 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- routes.response_cookies_to_add.ignore_expiry

<a id="canonical-0fdf96edef3b7af8a311859a8379d981c28929281f707180199af7af7ff8b522"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore expiry.

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

<a id="canonical-ef49670c3a4c6f63c8f116d765b836b0fad0a2200888c291acfd353532d64fbc"></a>

## Direct properties — routes.response_cookies_to_add.ignore_expiry / 6303e8fd2ce6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a0c803b474666875809a6cbe32096c249923d4b7594851272eb7a26409fdb1b6"></a>

## Next pages — routes.response_cookies_to_add.ignore_expiry / 6303e8fd2ce6 / 4

- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-07296ee33d006e2a6c774af261641078621866e3dba34060459375635f057ac7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df20cd67e75ca66114c599959637f30a73d571842233c003849845fc1c0e7814"></a>

## routes.response_cookies_to_add.ignore_httponly — routes.response_cookies_to_add.ignore_httponly / 470cf593d744 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- routes.response_cookies_to_add.ignore_httponly

<a id="canonical-c357f73075849990912020f23f9263815e0c7d5596ebced724dfc3b20f3b98fe"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore httponly.

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

<a id="canonical-e5e0762e2f1252a51fdc154784bd5f40221c7391274fdbde73bed2fb27a70e8b"></a>

## Direct properties — routes.response_cookies_to_add.ignore_httponly / 470cf593d744 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8a4a765df8e850bbff6e3390d6e6d206b6a3edb595e50f193fc9a722abeeb603"></a>

## Next pages — routes.response_cookies_to_add.ignore_httponly / 470cf593d744 / 4

- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-e3e68f08f9a68d0e962ff51851df5a3983255fb5ca84b52bd56bcd2b4375527d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd5f33e11093a86c745a6b344a269d8b1a21c718ad91a1f6c3fc79f1181e330f"></a>

## routes.response_cookies_to_add.ignore_max_age — routes.response_cookies_to_add.ignore_max_age / 1184e847b847 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- routes.response_cookies_to_add.ignore_max_age

<a id="canonical-d8afaa6917acfa8349aa1548b14bd5389c4773303d9260cc4628a3784ec1b72a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore max age.

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

<a id="canonical-2e108e44284ef69ff658d1923d5dee91395adcae866d0f3c706bbc9df4e3a26f"></a>

## Direct properties — routes.response_cookies_to_add.ignore_max_age / 1184e847b847 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0e9d6a4e8eb52daf1353e5078909b69cdc58a2f5a0edd377e3d21e761877ac58"></a>

## Next pages — routes.response_cookies_to_add.ignore_max_age / 1184e847b847 / 4

- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-ae2646a1c13fa21606f7b49638403cf9d8e457bdb02d5e10f37886cb1b78ddb7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c0e0c9d988b9005505f18b92fb207b63e8ed63569ab54c0a5e9d2e3246e73ca"></a>

## routes.response_cookies_to_add.ignore_partitioned — routes.response_cookies_to_add.ignore_partitioned / badf4dc5aad9 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- routes.response_cookies_to_add.ignore_partitioned

<a id="canonical-b0788674f651227f00e47ffa3e25d873b77d73b4a42d3e5d13e5515bb73ce697"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore partitioned.

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

<a id="canonical-45443af796f49027273d78f952af15d61d88289a3d9834a8171a2441eedb46cf"></a>

## Direct properties — routes.response_cookies_to_add.ignore_partitioned / badf4dc5aad9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9636c2d278a0f1cad07be044fb501f2ed7f5ec0778cc29062f85cce83ea52cae"></a>

## Next pages — routes.response_cookies_to_add.ignore_partitioned / badf4dc5aad9 / 4

- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-cc956b4c4e0b144773de65af99bca28c31b0ededf986260e651df1d948009e25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87e7f2a0f8ff63c6692442d6c922ead82ab96fdbbade32f7697a5201e9d4a777"></a>

## routes.response_cookies_to_add.ignore_path — routes.response_cookies_to_add.ignore_path / 37d1f2791995 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Property reference](data-sources--route--reference--group-001.md#canonical-a6c77d69b7834e989d425e83281601a25261f73228caae96eeda7871db19dfbf)
- [routes](data-sources--route--reference--group-001.md#canonical-19922725b52fed5452e47b4b9e3033774938f366da9529fca3df5d26e12a8f49)
- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- routes.response_cookies_to_add.ignore_path

<a id="canonical-0299cd1499ed2bbe50a3487b99a1ab643d0e55f22c344d536f93a43e4f6698d5"></a>

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

<a id="canonical-72c664dbac1ecb3fb1c74f7cd08d0030239b5b16c855adb370af22019b833559"></a>

## Direct properties — routes.response_cookies_to_add.ignore_path / 37d1f2791995 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cf3d04aa4baed9693527e4415571a2994b03d6009cbe77a099d02ec2aa578d4c"></a>

## Next pages — routes.response_cookies_to_add.ignore_path / 37d1f2791995 / 4

- [routes.response_cookies_to_add](data-sources--route--reference--group-001.md#canonical-e039119e22e362619c78c87daa5fea72f1bb1251f36985ef7df0d528a38e994d)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-6a9d642dbb4e762ecbe49828f9626f4a3b56d536f071ae46f6cae1079e7e0093"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
