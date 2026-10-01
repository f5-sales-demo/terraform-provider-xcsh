---
page_title: "xcsh_route reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route reference."
---

# xcsh_route reference

<a id="canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-732f74974a9d080cec5b3046c980d6932dd637c3c690f2a75447b94c7b855282"></a>

## Property reference — Property reference / 04890cbd4f28 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- Property reference

<a id="canonical-95d48edbbb8287461d6461a9a80805ca275568366a9b31948299e27c05fbb744"></a>

## Direct properties — Property reference / 04890cbd4f28 / 3

<a id="canonical-326df19403f7479415b5610bd6936157901801e14f3b1da8a1ee1fb1002e5c1e"></a>

<a id="canonical-1105766d97cc3907a9eda11097bac30603da9c33d01f68bf1ca5f45ca4e234bb"></a>

## annotations property — Property reference / 04890cbd4f28 / 4

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

<a id="canonical-40f4eda67f653c6a6d88cbdac8151c73b529129eb5d7e70266a47cc89b625fa7"></a>

<a id="canonical-507f4db5a7a49c311e34ac45e298205f16f506d7e91b8e8834736570fa39d327"></a>

## description property — Property reference / 04890cbd4f28 / 5

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

<a id="canonical-1baaf692f8fb06b8adaea88d7e4173faef9aefe7ea08e23a04b4b0d1ecb642fa"></a>

<a id="canonical-05b943a466bf52c0a4eca57a2b16cc2bee12dadb2661929002e69691bc40e9b8"></a>

## disable property — Property reference / 04890cbd4f28 / 6

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

<a id="canonical-ba502415db4d5fed1b7c5a9668e90e2028a994a715a7c98e864ba10e7115a60d"></a>

<a id="canonical-8ce0c932ff22ef2180be6788f25efed8d835f5314e3b162fc1609a9e7d1b67f1"></a>

## id property — Property reference / 04890cbd4f28 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-db27a1a98550964bffa6a627b7f246e29078df73b17cd149922fa76d3cdb2a7e"></a>

<a id="canonical-4c48c608b911d6c74c841bb09f9222e158d75f2b5f6dd17e2b4b6c59e830ae53"></a>

## labels property — Property reference / 04890cbd4f28 / 8

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

<a id="canonical-f97c0597ed42695e9b2c6e6211560fc964d93f7241251c0423c53f0e95eca9d0"></a>

<a id="canonical-efdc17c4740f84c14691cf67660682777860c36c0102beb8d8cbde9240a9f6ca"></a>

## name property — Property reference / 04890cbd4f28 / 9

Type: `"string"`. Required.

Name of the Route. Must be unique within the namespace.

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

<a id="canonical-e01b50229dc77c6fabc76b0127967879f0040a66f8277684fe576249070791fa"></a>

<a id="canonical-fb25dc31c5df054415f43a9b4e1e1deda45d65e6fa9c6b3f2a9c97da323c888e"></a>

## namespace property — Property reference / 04890cbd4f28 / 10

Type: `"string"`. Required.

Namespace where the Route is created.

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

- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c): complete subsection reference.

- [timeouts](resources--route--reference--group-003.md#canonical-806a5af33b8df63de7fc42ab58f766c42f08c0fd462d03b1e06b3e9ed42e23ac): complete subsection reference.

<a id="canonical-cc3d95b63dead41b39e03d40fba59828608e79bc7b4d99e7370ab39cde626754"></a>

## All schema paths — Property reference / 04890cbd4f28 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--route--reference--group-001.md#canonical-326df19403f7479415b5610bd6936157901801e14f3b1da8a1ee1fb1002e5c1e) |
| `description` | [description](resources--route--reference--group-001.md#canonical-40f4eda67f653c6a6d88cbdac8151c73b529129eb5d7e70266a47cc89b625fa7) |
| `disable` | [disable](resources--route--reference--group-001.md#canonical-1baaf692f8fb06b8adaea88d7e4173faef9aefe7ea08e23a04b4b0d1ecb642fa) |
| `id` | [id](resources--route--reference--group-001.md#canonical-ba502415db4d5fed1b7c5a9668e90e2028a994a715a7c98e864ba10e7115a60d) |
| `labels` | [labels](resources--route--reference--group-001.md#canonical-db27a1a98550964bffa6a627b7f246e29078df73b17cd149922fa76d3cdb2a7e) |
| `name` | [name](resources--route--reference--group-001.md#canonical-f97c0597ed42695e9b2c6e6211560fc964d93f7241251c0423c53f0e95eca9d0) |
| `namespace` | [namespace](resources--route--reference--group-001.md#canonical-e01b50229dc77c6fabc76b0127967879f0040a66f8277684fe576249070791fa) |
| `routes` | [routes](resources--route--reference--group-001.md#canonical-1fa4dfa72628ad53a1a3d8dd57a551d8435dd2c7c08e17e87bfb2c93e2621a51) |
| `routes.bot_defense_javascript_injection` | [routes.bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-94c0cafb9b6b1b95750032ded40f1199f5ded17f08e6bfc01460a5f4725695b6) |
| `routes.bot_defense_javascript_injection.javascript_location` | [routes.bot_defense_javascript_injection.javascript_location](resources--route--reference--group-001.md#canonical-3fc247b880b57b72e64ad5d61d4739c5af4bba7287424fc200aaa230f2ccb198) |
| `routes.bot_defense_javascript_injection.javascript_tags` | [routes.bot_defense_javascript_injection.javascript_tags](resources--route--reference--group-001.md#canonical-cb009e057c4bc5872b109925e4b50db75d30e077941035a12fda78b7b2ba9261) |
| `routes.bot_defense_javascript_injection.javascript_tags.javascript_url` | [routes.bot_defense_javascript_injection.javascript_tags.javascript_url](resources--route--reference--group-001.md#canonical-9d72669a1bf33d7842bfe8e660edc72cc2c7fb5dddafca733c4a07d4346ef4df) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes](resources--route--reference--group-001.md#canonical-b933bcf011bb43183a8a29de1f7f93d5188aa07f2fc85123d0f422464bce1b6a) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag](resources--route--reference--group-001.md#canonical-c7507b72577deb4c5605c36124f246bd133ed228d699904b6f87958e1ed1b7fe) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value](resources--route--reference--group-001.md#canonical-83c6a66ad06e60a576246f829c15215de2a1e929fdea48e7464dfd42eec25dad) |
| `routes.disable_location_add` | [routes.disable_location_add](resources--route--reference--group-001.md#canonical-db74ab88546a70c3b5dd31a84281910f1ef1c7a2f5163cef38494ebd75d35a72) |
| `routes.inherited_bot_defense_javascript_injection` | [routes.inherited_bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-cd53f990f5cc16be6a8ba1f7f525c557a3f60b0bfb7499b96cb11524a90ca807) |
| `routes.inherited_waf_exclusion` | [routes.inherited_waf_exclusion](resources--route--reference--group-001.md#canonical-2aee3d63da6d7b758ce1b5bdf1a84f5a053e6a849c4bb1089e695bd22e7d2c8b) |
| `routes.match` | [routes.match](resources--route--reference--group-001.md#canonical-66b856fd32205a41f3a0395342c677c381f1f772e86e2769605f03574e2bd1c6) |
| `routes.match.headers` | [routes.match.headers](resources--route--reference--group-001.md#canonical-92a18cc6ef9abe90fd95d3f316a33fdcffeecb15437c7fed26e52e3fe818774a) |
| `routes.match.headers.exact` | [routes.match.headers.exact](resources--route--reference--group-001.md#canonical-5099fd9ea8281358af2d5dfa728f347ffc5546c5f35d809075529ad34158e7c7) |
| `routes.match.headers.invert_match` | [routes.match.headers.invert_match](resources--route--reference--group-001.md#canonical-0f8f1b1313226d57e6660c7bc683269a05c55cd1e4ca62496ec0fba85f3c6fa1) |
| `routes.match.headers.name` | [routes.match.headers.name](resources--route--reference--group-001.md#canonical-b3034d8a6917ba113db04a2961614414c2c49345e8c5d952d60ed2e1af447be1) |
| `routes.match.headers.presence` | [routes.match.headers.presence](resources--route--reference--group-001.md#canonical-8666b35f5e5a3ec1570f5dd8c331fe18a7c33b7bfad4a2c7aa82a2c1d00ed6bc) |
| `routes.match.headers.regex` | [routes.match.headers.regex](resources--route--reference--group-001.md#canonical-b234b50b5f958be467243819f9a6b37dd4d1098909c0c45c66c783c19b1c95cb) |
| `routes.match.http_method` | [routes.match.http_method](resources--route--reference--group-001.md#canonical-70f91849c82960be70f0d34ea5415206ca2357b903fd597870f2e46f16804d90) |
| `routes.match.incoming_port` | [routes.match.incoming_port](resources--route--reference--group-001.md#canonical-fadffe6f6b8ff95193f78e792093f39a45a093eb1542fd47b913b5ac53726510) |
| `routes.match.incoming_port.no_port_match` | [routes.match.incoming_port.no_port_match](resources--route--reference--group-001.md#canonical-a33c2f60761a88efa5245f2ebe0e3da44ef44fa71b0b80b9593eda8cc7099ccf) |
| `routes.match.incoming_port.port` | [routes.match.incoming_port.port](resources--route--reference--group-001.md#canonical-5190550350e307bc4f1a77a01780f0edac370a99b236ab4e9111befb8dbdbc49) |
| `routes.match.incoming_port.port_ranges` | [routes.match.incoming_port.port_ranges](resources--route--reference--group-001.md#canonical-684480b7d08e78468740b4b145b8fb75f85cfbff99f6bd17fd2c4c6a3f9c3b14) |
| `routes.match.path` | [routes.match.path](resources--route--reference--group-001.md#canonical-b44fb2279d9ec892e82ab40443ec8797cfcf366d05d8afe1630dca89f380fd0a) |
| `routes.match.path.path` | [routes.match.path.path](resources--route--reference--group-001.md#canonical-fc2ce8204828d2515d5745d6e4c97574429531b840695648472578b1f80d1049) |
| `routes.match.path.prefix` | [routes.match.path.prefix](resources--route--reference--group-001.md#canonical-3dd953ee81d52702740321576baddd589fcf78d6d4184e9a2813eac556bd5e0c) |
| `routes.match.path.regex` | [routes.match.path.regex](resources--route--reference--group-001.md#canonical-e6ccee598c35bf7585b25f616def04db4c4ee8c5a404f13befc3ae5c0bc5a6c1) |
| `routes.match.query_params` | [routes.match.query_params](resources--route--reference--group-001.md#canonical-7982de1160eac5fe30f0cfbaab4bafe4042f4069f4d20b8460cd755821afcaeb) |
| `routes.match.query_params.exact` | [routes.match.query_params.exact](resources--route--reference--group-001.md#canonical-e334a032dde107e6fdc981cbe811bf8d4eb5f44d3f44736f204363f24bd20936) |
| `routes.match.query_params.key` | [routes.match.query_params.key](resources--route--reference--group-001.md#canonical-9873cfb71e6d5ab7f946921474c77edf942bfa367c6d27d5b995e45d1e0ffe86) |
| `routes.match.query_params.regex` | [routes.match.query_params.regex](resources--route--reference--group-001.md#canonical-6a2a73266c82d56c91d1d440afe5e0c09e8dfd037cf993599f5964928bd70b30) |
| `routes.request_cookies_to_add` | [routes.request_cookies_to_add](resources--route--reference--group-001.md#canonical-27de8d8bfb56c062449e4ad98ecfa3d875099af52fc2a9575d584dd0cababc08) |
| `routes.request_cookies_to_add.name` | [routes.request_cookies_to_add.name](resources--route--reference--group-001.md#canonical-ebe3dbab488ca777eea820c368ac64f47b94741752bc6997d1076205339ea186) |
| `routes.request_cookies_to_add.overwrite` | [routes.request_cookies_to_add.overwrite](resources--route--reference--group-001.md#canonical-138a1842381f3759d547273c98cba5036e192b9314ea82dccc8a2b62e02a69da) |
| `routes.request_cookies_to_add.secret_value` | [routes.request_cookies_to_add.secret_value](resources--route--reference--group-001.md#canonical-4005b88294eac9ae4ea07c15ab929e13f564e4a0bd1d4d97227dd3e01a3852cb) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info](resources--route--reference--group-001.md#canonical-b44ab3c7b1fdf97bec90261535d2133a66ae2925a6aa894c3eeea84cb976d1f5) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--route--reference--group-001.md#canonical-5c47571f857f223fb73bc08684129fcf29c412be16b965c57f48851902afd11a) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.location](resources--route--reference--group-001.md#canonical-d5b8d385db3a2f93528bc08f963dd71d2bc6f8f9a6921b67e7a1f81fad768031) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--route--reference--group-001.md#canonical-b2853e6aaec806f22b34c655e6ad37b81c572c90728e0097c461ace103c92bcd) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info` | [routes.request_cookies_to_add.secret_value.clear_secret_info](resources--route--reference--group-001.md#canonical-7213a3c2022c5d67c37a79bb8b732b2c52842f976c01066245e8b01acad5badb) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [routes.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--route--reference--group-001.md#canonical-31e882f43e4b4cc422179f639fe40da5d0a1aafbc8d901d3bef085d6ae4c5da3) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info.url` | [routes.request_cookies_to_add.secret_value.clear_secret_info.url](resources--route--reference--group-001.md#canonical-e2810fd764a7feb524f7c420bdfa56382c4a47312991a5692b0275bec044f769) |
| `routes.request_cookies_to_add.value` | [routes.request_cookies_to_add.value](resources--route--reference--group-001.md#canonical-78701ad630e610c51f81a5865aefd95ab6c7750cf2ff6aaeec16e8fc1edb762b) |
| `routes.request_cookies_to_remove` | [routes.request_cookies_to_remove](resources--route--reference--group-001.md#canonical-c6560ab78fa9d8c06341610ea2fbbb2c934d38658e1e82f5e3974bd83d2a7b86) |
| `routes.request_headers_to_add` | [routes.request_headers_to_add](resources--route--reference--group-001.md#canonical-5e8c41fbc8d3b7ff01e349c10dcdd61032c456f3f7d566e5dd91b4253604708e) |
| `routes.request_headers_to_add.append` | [routes.request_headers_to_add.append](resources--route--reference--group-001.md#canonical-6d532631a78ee46692a95d94b7f020c35dec7eb42a6d4b5ff95bb5ce78648c7f) |
| `routes.request_headers_to_add.name` | [routes.request_headers_to_add.name](resources--route--reference--group-001.md#canonical-3dbdd5922ff3198b302113137ea5f9b337819d8e50b6ee9c9ade6e55bb60b0b1) |
| `routes.request_headers_to_add.secret_value` | [routes.request_headers_to_add.secret_value](resources--route--reference--group-001.md#canonical-5288358233be7aa12d0e66909f59ce2e3ff87120213212c2dd5a5d165c654746) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info` | [routes.request_headers_to_add.secret_value.blindfold_secret_info](resources--route--reference--group-001.md#canonical-09eb76b0397b9e3008decba1fd3c64d1f81592d69552e37c5f4865450d054021) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--route--reference--group-001.md#canonical-4a4d902f9741fa656cc6578328e427f505f85f1ab65be506528e0feb75532ac4) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.location` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.location](resources--route--reference--group-001.md#canonical-a22eb163303d2a416692cc24cfa58bc1ccb6b5b05a23649c19d5aaab6c0c4e15) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--route--reference--group-001.md#canonical-f72721aa7e8c7f7e3f094e87b55dac0cd9644439b77c2278b5756f87f6b059fb) |
| `routes.request_headers_to_add.secret_value.clear_secret_info` | [routes.request_headers_to_add.secret_value.clear_secret_info](resources--route--reference--group-001.md#canonical-a0b57f52b88da5409685e96d143b852be43d6488625e992117611c29ed0bf273) |
| `routes.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [routes.request_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--route--reference--group-001.md#canonical-7b2f55e67c1d51baf6ea23d4e153abf1301b1fc8dc9b4ce88e8f5a0646ebf981) |
| `routes.request_headers_to_add.secret_value.clear_secret_info.url` | [routes.request_headers_to_add.secret_value.clear_secret_info.url](resources--route--reference--group-001.md#canonical-ec2ce7edaf1ceffd38fa18fc0800419d0b00f09059e0c1ada5de9b0664bc54d7) |
| `routes.request_headers_to_add.value` | [routes.request_headers_to_add.value](resources--route--reference--group-001.md#canonical-3e99e19390e09b0d5514545b3dd67a3292407cb4435e72b4beab7d4efbe18299) |
| `routes.request_headers_to_remove` | [routes.request_headers_to_remove](resources--route--reference--group-001.md#canonical-807896249b94ceffd733ce3a7410f92535df7b50106761943cdde863ea9b84d3) |
| `routes.response_cookies_to_add` | [routes.response_cookies_to_add](resources--route--reference--group-001.md#canonical-07fd928eb654ed5f287820ef5b21f9d944cbea7f27814b89add5437e1184f5fc) |
| `routes.response_cookies_to_add.add_domain` | [routes.response_cookies_to_add.add_domain](resources--route--reference--group-001.md#canonical-6cc7e58cfb7c93062b8ba5138f29c77ea3a1228553a0be0ebae64f329aa3bfdd) |
| `routes.response_cookies_to_add.add_expiry` | [routes.response_cookies_to_add.add_expiry](resources--route--reference--group-001.md#canonical-b5dbd1e904cbea620d964cd54940ba79c8ded3607459c8720b86dba3936edec8) |
| `routes.response_cookies_to_add.add_httponly` | [routes.response_cookies_to_add.add_httponly](resources--route--reference--group-001.md#canonical-f767477760ba00b8c4d72d479ac207aff37ddcee5c7749a53ed65a50e1d110eb) |
| `routes.response_cookies_to_add.add_partitioned` | [routes.response_cookies_to_add.add_partitioned](resources--route--reference--group-001.md#canonical-e249f9712e59a25b9811416a0cbb02777e491d7c0b433b8d8f612914accc08ba) |
| `routes.response_cookies_to_add.add_path` | [routes.response_cookies_to_add.add_path](resources--route--reference--group-001.md#canonical-a4a045a1f09c68ed25ff658e2a6b9d861366af77ac9b50dbe3b0c70e353e639b) |
| `routes.response_cookies_to_add.add_secure` | [routes.response_cookies_to_add.add_secure](resources--route--reference--group-001.md#canonical-8882192c917b12fac9d0f80ac6919205fdf0f786df2ed60f25d8f2f03ee684e8) |
| `routes.response_cookies_to_add.ignore_domain` | [routes.response_cookies_to_add.ignore_domain](resources--route--reference--group-001.md#canonical-3aac7b78166731a78fab42a4e8e0fc3ff3a99bd0c53557506c9ba0c1ff86a9bb) |
| `routes.response_cookies_to_add.ignore_expiry` | [routes.response_cookies_to_add.ignore_expiry](resources--route--reference--group-001.md#canonical-2bd41275e146d4599b607af35ab924eef44b500178405f1cea65fa811075156f) |
| `routes.response_cookies_to_add.ignore_httponly` | [routes.response_cookies_to_add.ignore_httponly](resources--route--reference--group-001.md#canonical-44c08f1083e600b9bf9338ec37ac92853cf23eca729a47269ddbde0258a42874) |
| `routes.response_cookies_to_add.ignore_max_age` | [routes.response_cookies_to_add.ignore_max_age](resources--route--reference--group-002.md#canonical-f5730eddb2b740271ed14cc8b268463e2cf939f86cb651265150419c59a76c01) |
| `routes.response_cookies_to_add.ignore_partitioned` | [routes.response_cookies_to_add.ignore_partitioned](resources--route--reference--group-002.md#canonical-7299ad3fedaa85ae2d4274899295f8c0b7d08d0e7327b0369789b26ff885d5f2) |
| `routes.response_cookies_to_add.ignore_path` | [routes.response_cookies_to_add.ignore_path](resources--route--reference--group-002.md#canonical-fff8dd84901ce1a7d9d803001c84d1938317356271a910ab6cd1bf0378092fb6) |
| `routes.response_cookies_to_add.ignore_samesite` | [routes.response_cookies_to_add.ignore_samesite](resources--route--reference--group-002.md#canonical-539ea52edc8605db6ff1e54c74333eb56cb7a3ff785e38095561ef0c7b02cd4a) |
| `routes.response_cookies_to_add.ignore_secure` | [routes.response_cookies_to_add.ignore_secure](resources--route--reference--group-002.md#canonical-c89875a4bdcd238bd279680defd59c15ce56fad2bd058f91a97a201500d0814c) |
| `routes.response_cookies_to_add.ignore_value` | [routes.response_cookies_to_add.ignore_value](resources--route--reference--group-002.md#canonical-d37ed5511963d42a292f59466df35fbd5d0b92d154f3954e71974b70b9431c40) |
| `routes.response_cookies_to_add.max_age_value` | [routes.response_cookies_to_add.max_age_value](resources--route--reference--group-001.md#canonical-ba7212e6b168140d9342821166570975de53e0bc4793fc7aa16a00ce27ce829a) |
| `routes.response_cookies_to_add.name` | [routes.response_cookies_to_add.name](resources--route--reference--group-001.md#canonical-a214c8320a34ed82908d144c456e8367c7d9d8da58efaad7bf56e687445eea9e) |
| `routes.response_cookies_to_add.overwrite` | [routes.response_cookies_to_add.overwrite](resources--route--reference--group-001.md#canonical-9db105956b7f874ce8aa19b36d4d95fa5ed969e2432208c281d72b4371245272) |
| `routes.response_cookies_to_add.samesite_lax` | [routes.response_cookies_to_add.samesite_lax](resources--route--reference--group-002.md#canonical-1607228bbbcfe57c0286c4a678029da785746ba2841044756666ca066b99477f) |
| `routes.response_cookies_to_add.samesite_none` | [routes.response_cookies_to_add.samesite_none](resources--route--reference--group-002.md#canonical-f9cb48985a61fe595b3e9ad05a6c421fa70f8dfea129c9c35b1ecaf3eafa560d) |
| `routes.response_cookies_to_add.samesite_strict` | [routes.response_cookies_to_add.samesite_strict](resources--route--reference--group-002.md#canonical-1c926807e1c573b07bc7c92b7b06ff39755d9b675440d3a28859e01a9e65afd2) |
| `routes.response_cookies_to_add.secret_value` | [routes.response_cookies_to_add.secret_value](resources--route--reference--group-002.md#canonical-ff1a2b6467b435d7c7522809dc9f5f6b722065de61d6b71ea6532bad2a72627d) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info](resources--route--reference--group-002.md#canonical-335eb48a47aa7bb60ff748600f5940baeb6ef394607d46c206ef21b2dca57fff) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--route--reference--group-002.md#canonical-7882e5a934475797ba0cf6e73aa5e0cf1b58c8606761a6314f3434f4d2204da1) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.location](resources--route--reference--group-002.md#canonical-fee403070c29703f7fd5edb32e2cd60e6dd5004f4f0aee8f9c5ed267a0698dde) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--route--reference--group-002.md#canonical-23d462486101936360f9363b63b9641f909644e7ce263500fd69040c0409230b) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info` | [routes.response_cookies_to_add.secret_value.clear_secret_info](resources--route--reference--group-002.md#canonical-10f316a63e0a22e058a1ee73744acf4566fee6539ce8803fb140f818dac73138) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [routes.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--route--reference--group-002.md#canonical-2ad5ea253ba528ae8e8e56c6681d7fdc779dfae7d30e6fa88c555f9a486fb484) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info.url` | [routes.response_cookies_to_add.secret_value.clear_secret_info.url](resources--route--reference--group-002.md#canonical-6542a32b334a7a7a421f315d4f1a45aa2607825de116f0de5d370098b0493ca6) |
| `routes.response_cookies_to_add.value` | [routes.response_cookies_to_add.value](resources--route--reference--group-001.md#canonical-57706fa921fe841ccd992aee0df8434df4945b5ead879cc310a58a32a4fbe31f) |
| `routes.response_cookies_to_remove` | [routes.response_cookies_to_remove](resources--route--reference--group-001.md#canonical-c06da782c3708cb3540e3c27983584fad0a45bb2664e56906b7e1a0390db5845) |
| `routes.response_headers_to_add` | [routes.response_headers_to_add](resources--route--reference--group-002.md#canonical-c6e4d622e2491269233ea717d52c0a0a4d8aab6f0f41c5db4f5ff591ada8b8c0) |
| `routes.response_headers_to_add.append` | [routes.response_headers_to_add.append](resources--route--reference--group-002.md#canonical-3b4c39d68c712e49466068f1b0af88c80a887fd20084e96c8f412d95aae83276) |
| `routes.response_headers_to_add.name` | [routes.response_headers_to_add.name](resources--route--reference--group-002.md#canonical-5963b015b9ab14223ea5d4811e6b055b1dc97cb3be855b63fe31213d794dd1d0) |
| `routes.response_headers_to_add.secret_value` | [routes.response_headers_to_add.secret_value](resources--route--reference--group-002.md#canonical-0029da3019e515f5e7d467c84feaa577ca87919522eddcdc3bb8ee0bba18201d) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info` | [routes.response_headers_to_add.secret_value.blindfold_secret_info](resources--route--reference--group-002.md#canonical-7ec79c2e0f3f6c62f016c5a0fa43f7fbdad9464b7bd58bb0bf2008a63161cf81) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--route--reference--group-002.md#canonical-1f714712f587d6c649ccf0e48aedb2afd30f3fe82dbed966866d9eea53630d31) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.location` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.location](resources--route--reference--group-002.md#canonical-f0aca7f7002390694e8ed1e0929a95e491b4773ab11c2fba4ea25979c99f2f47) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--route--reference--group-002.md#canonical-b28201a4970c7d261f01e4064d3d121ae0ec9d82cd6a82dd9eb20f5ebfb10b4a) |
| `routes.response_headers_to_add.secret_value.clear_secret_info` | [routes.response_headers_to_add.secret_value.clear_secret_info](resources--route--reference--group-002.md#canonical-7dd9a1c77c87e75a2a226f568a5b40961c53a9c491af0665829e04e3f10fe6ee) |
| `routes.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [routes.response_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--route--reference--group-002.md#canonical-1d07019eca84c2891bd8ba579675482483cd0159f66850880ea0cbb08c407bb5) |
| `routes.response_headers_to_add.secret_value.clear_secret_info.url` | [routes.response_headers_to_add.secret_value.clear_secret_info.url](resources--route--reference--group-002.md#canonical-2b113c14258634c38667b037c74fdfbc25a9811e81b71158a557ab6702a07cd5) |
| `routes.response_headers_to_add.value` | [routes.response_headers_to_add.value](resources--route--reference--group-002.md#canonical-3dd3d83171dc83239de7eab7769204ef397e66f3934587fca72dc1d722ea141c) |
| `routes.response_headers_to_remove` | [routes.response_headers_to_remove](resources--route--reference--group-001.md#canonical-89e6e0ce4839d5f6442ce1acedbcd8c48590a68c65797ab4483e2064c9779af1) |
| `routes.route_destination` | [routes.route_destination](resources--route--reference--group-002.md#canonical-e27c522a04b37e751d571dcf74a6e43927b7b12ef1c880a2d3ff410e38722ca1) |
| `routes.route_destination.auto_host_rewrite` | [routes.route_destination.auto_host_rewrite](resources--route--reference--group-002.md#canonical-487a4ab7ed483f97d4969f260c05ced345ca1efba614172dad909b04e97f9146) |
| `routes.route_destination.buffer_policy` | [routes.route_destination.buffer_policy](resources--route--reference--group-002.md#canonical-fe79dcd4f327b4e33a77b2392dcbe2984adbcfc52e47060997c935cb51bd9c9c) |
| `routes.route_destination.buffer_policy.disabled` | [routes.route_destination.buffer_policy.disabled](resources--route--reference--group-002.md#canonical-3c6dfb0742d2d77cd45c10b997cc14685184a9b250c195efc5717ea8bad295fa) |
| `routes.route_destination.buffer_policy.max_request_bytes` | [routes.route_destination.buffer_policy.max_request_bytes](resources--route--reference--group-002.md#canonical-1fc3f79066d3e4f8b086386276b5e9ddd55ae1560e8ecd62e69d28445b8122ce) |
| `routes.route_destination.cors_policy` | [routes.route_destination.cors_policy](resources--route--reference--group-002.md#canonical-42045d04796b5888b76ea574f859fbe8ac57633751dc949e1fb73a3ccd0a0ba4) |
| `routes.route_destination.cors_policy.allow_credentials` | [routes.route_destination.cors_policy.allow_credentials](resources--route--reference--group-002.md#canonical-fb3d7c51320d6cb39d1a29bc8bd303b8688bb91b808b1626baa7060ad194a89f) |
| `routes.route_destination.cors_policy.allow_headers` | [routes.route_destination.cors_policy.allow_headers](resources--route--reference--group-002.md#canonical-e8a283e09a12951eef34e7ed30e9ecd71d33286e5f5b4d54ca37e013e077d2a3) |
| `routes.route_destination.cors_policy.allow_methods` | [routes.route_destination.cors_policy.allow_methods](resources--route--reference--group-002.md#canonical-2c6e7d17f5ac3e0440ce3fcc6f8aedfe7d3cdeb7304e974cc23126d05349555c) |
| `routes.route_destination.cors_policy.allow_origin` | [routes.route_destination.cors_policy.allow_origin](resources--route--reference--group-002.md#canonical-5467223716f31a996c8dfc45a3ff1efc336b2c80016a841c078ec0f62a308bd5) |
| `routes.route_destination.cors_policy.allow_origin_regex` | [routes.route_destination.cors_policy.allow_origin_regex](resources--route--reference--group-002.md#canonical-55810e51152f011343377b77dfda5c90cc9a07dedf59f14d1d48dfca2081384e) |
| `routes.route_destination.cors_policy.disabled` | [routes.route_destination.cors_policy.disabled](resources--route--reference--group-002.md#canonical-58d9b66b66aa30b3993905af3741d0ccbcbd176781824ba50e104be601a39efb) |
| `routes.route_destination.cors_policy.expose_headers` | [routes.route_destination.cors_policy.expose_headers](resources--route--reference--group-002.md#canonical-ab597efad87ae8cce7dfae84019de4b29bb572f8dac095562a988801ddfdd889) |
| `routes.route_destination.cors_policy.maximum_age` | [routes.route_destination.cors_policy.maximum_age](resources--route--reference--group-002.md#canonical-fdd7973536be3d81c72b6e9378789845ed314659ac7fa8b019a4f95c75d20a29) |
| `routes.route_destination.csrf_policy` | [routes.route_destination.csrf_policy](resources--route--reference--group-002.md#canonical-61457478bbc5b52560b1c6aff1c6fca1c9909c5e49c41dd990491aebafcb5219) |
| `routes.route_destination.csrf_policy.all_load_balancer_domains` | [routes.route_destination.csrf_policy.all_load_balancer_domains](resources--route--reference--group-002.md#canonical-954679b01c313ebf7e6eb602d54dfdf15b5ffff746712308570f6a8a441add82) |
| `routes.route_destination.csrf_policy.custom_domain_list` | [routes.route_destination.csrf_policy.custom_domain_list](resources--route--reference--group-002.md#canonical-858175b04aa2bc8d60ff6026ff01602f20766c0d4f99c816aee918334419e860) |
| `routes.route_destination.csrf_policy.custom_domain_list.domains` | [routes.route_destination.csrf_policy.custom_domain_list.domains](resources--route--reference--group-002.md#canonical-971b77d36c593313cd193f2913807da7e2f8447f05bbb3ac3f50d525a513276f) |
| `routes.route_destination.csrf_policy.disabled` | [routes.route_destination.csrf_policy.disabled](resources--route--reference--group-002.md#canonical-9cb7d56b66889255b448993ba4a553ee11c9c37669c96aba7f3b712931d27f4b) |
| `routes.route_destination.destinations` | [routes.route_destination.destinations](resources--route--reference--group-002.md#canonical-56247b79176a421f1fc8b11864d96e440a3c6759e140434931656869cf65ec12) |
| `routes.route_destination.destinations.cluster` | [routes.route_destination.destinations.cluster](resources--route--reference--group-002.md#canonical-d63b9aab4425f16557d2b3ddde70463a83da9243070de7380f5a740b68de0b22) |
| `routes.route_destination.destinations.cluster.kind` | [routes.route_destination.destinations.cluster.kind](resources--route--reference--group-002.md#canonical-8b956820105fdf29aa4e90dda5e4e73970b2b6fdec899a20c2a6f6a45d319118) |
| `routes.route_destination.destinations.cluster.name` | [routes.route_destination.destinations.cluster.name](resources--route--reference--group-002.md#canonical-d0119abfc40799ab27315488cb9e0cf8586d89a965ffe2d34a4ba4aa5188dfb0) |
| `routes.route_destination.destinations.cluster.namespace` | [routes.route_destination.destinations.cluster.namespace](resources--route--reference--group-002.md#canonical-65548e04b5569eecb83cfb1769a142c3961dcd0069ce1d90a48e135dff2bb71c) |
| `routes.route_destination.destinations.cluster.tenant` | [routes.route_destination.destinations.cluster.tenant](resources--route--reference--group-002.md#canonical-bd96e9cd32bc89527ec7e02fc9ab21ab802843a2e463b77ca65a72bc5aac5434) |
| `routes.route_destination.destinations.cluster.uid` | [routes.route_destination.destinations.cluster.uid](resources--route--reference--group-002.md#canonical-f96d7a8a910116afd4224bc542477aed624df8faccaeca5498f3e523051ce420) |
| `routes.route_destination.destinations.endpoint_subsets` | [routes.route_destination.destinations.endpoint_subsets](resources--route--reference--group-002.md#canonical-cb5421d63d1c8c43e6ec214256aecb91df05a8ea1006385003619cee6b8e911b) |
| `routes.route_destination.destinations.priority` | [routes.route_destination.destinations.priority](resources--route--reference--group-002.md#canonical-2d968cb4cafd987d431e94cf028addcb2513e3c4cac65f8d716b0c0a8f674bc2) |
| `routes.route_destination.destinations.weight` | [routes.route_destination.destinations.weight](resources--route--reference--group-002.md#canonical-a56b88ae2f7634d026fb76ba1ea1124f1cae3e9388441a52340802d8f3d3f32f) |
| `routes.route_destination.do_not_retract_cluster` | [routes.route_destination.do_not_retract_cluster](resources--route--reference--group-002.md#canonical-df72791a4dcd09c5552b3cfc36192a7e01b96a3a3e4fe69ddf8a80aa4971a068) |
| `routes.route_destination.endpoint_subsets` | [routes.route_destination.endpoint_subsets](resources--route--reference--group-002.md#canonical-9f7b6fdc30b59d937ea7a0c495efc8cd9def602e0ea9a58ea5df64c4e5e8586f) |
| `routes.route_destination.hash_policy` | [routes.route_destination.hash_policy](resources--route--reference--group-002.md#canonical-476d4a93073ae71c321ecaa70afc256e953d4c6fa6fbc845032f4d53b7500122) |
| `routes.route_destination.hash_policy.cookie` | [routes.route_destination.hash_policy.cookie](resources--route--reference--group-002.md#canonical-19790a290beca93da0eb07dde3416ae708e591ba00cdac7e979802b89c059da9) |
| `routes.route_destination.hash_policy.cookie.add_httponly` | [routes.route_destination.hash_policy.cookie.add_httponly](resources--route--reference--group-002.md#canonical-e6b6119c47c1661d9642179650cf71c862892bef342f23af4d4703a88c64cac1) |
| `routes.route_destination.hash_policy.cookie.add_secure` | [routes.route_destination.hash_policy.cookie.add_secure](resources--route--reference--group-002.md#canonical-b699cbdde9f533fccb680c1e86562f7a2b61204316f72107568450cf31fdbeb0) |
| `routes.route_destination.hash_policy.cookie.ignore_httponly` | [routes.route_destination.hash_policy.cookie.ignore_httponly](resources--route--reference--group-002.md#canonical-7718dfe36220ef61313e1ebfed071cf7c117b860054372e9f3e9144638e57b23) |
| `routes.route_destination.hash_policy.cookie.ignore_samesite` | [routes.route_destination.hash_policy.cookie.ignore_samesite](resources--route--reference--group-002.md#canonical-46caafca5d7c88f7ae8ae68740205efc56f6ff3507b6c796d1eadc3583466e51) |
| `routes.route_destination.hash_policy.cookie.ignore_secure` | [routes.route_destination.hash_policy.cookie.ignore_secure](resources--route--reference--group-002.md#canonical-00e07401f24b5551959d57f2b380bbdd2c041960903345a8f5100a3e8c876417) |
| `routes.route_destination.hash_policy.cookie.name` | [routes.route_destination.hash_policy.cookie.name](resources--route--reference--group-002.md#canonical-b2330a4bbd3093d1763831472159d25c49a8e0269451aeb07647aa07d8590632) |
| `routes.route_destination.hash_policy.cookie.path` | [routes.route_destination.hash_policy.cookie.path](resources--route--reference--group-002.md#canonical-a298446cd775bcba039ec7301fce06b43a183a41355959d4e73ec4f98b8ff2e4) |
| `routes.route_destination.hash_policy.cookie.samesite_lax` | [routes.route_destination.hash_policy.cookie.samesite_lax](resources--route--reference--group-002.md#canonical-a83dfecafa8be4389753bcd0f1e5acf3f998c7ee92f9ae5e49746267f2f826e8) |
| `routes.route_destination.hash_policy.cookie.samesite_none` | [routes.route_destination.hash_policy.cookie.samesite_none](resources--route--reference--group-002.md#canonical-deef3e0480f6e2d06ef1ec413c35087af4b1567203f15e1069c4d21637366b09) |
| `routes.route_destination.hash_policy.cookie.samesite_strict` | [routes.route_destination.hash_policy.cookie.samesite_strict](resources--route--reference--group-002.md#canonical-8136c825705fadf015c3c48e0a662d35c3807d0c90debb5daca8e2847a0986e2) |
| `routes.route_destination.hash_policy.cookie.ttl` | [routes.route_destination.hash_policy.cookie.ttl](resources--route--reference--group-002.md#canonical-3fc9ddd9c3a391f2322f4dcff17d8c8a07980f8182d3b14b6cf8d9ee57ae7fec) |
| `routes.route_destination.hash_policy.header_name` | [routes.route_destination.hash_policy.header_name](resources--route--reference--group-002.md#canonical-ccf92578c85eeea6d1be3d6b442790e0fa3bbe6b0f59bee91ae41b1228c8d84f) |
| `routes.route_destination.hash_policy.source_ip` | [routes.route_destination.hash_policy.source_ip](resources--route--reference--group-002.md#canonical-ea816efd1e05514f4d2170628c24304751a92da0724f05fe239436ff74bfd2de) |
| `routes.route_destination.hash_policy.terminal` | [routes.route_destination.hash_policy.terminal](resources--route--reference--group-002.md#canonical-c06e26156c2e45d4cb24b30987905c9fe8e32d7072d52bd7d94c3ba7b56a5013) |
| `routes.route_destination.host_rewrite` | [routes.route_destination.host_rewrite](resources--route--reference--group-002.md#canonical-2763aea5bbf819acc2ad5ae1515d946dffa61c79e0695548fa39555d2eb3edf1) |
| `routes.route_destination.mirror_policy` | [routes.route_destination.mirror_policy](resources--route--reference--group-002.md#canonical-d34e8f8af1b01b03c6cdb511abc466a2a534bf95811fbae0edd3a2755a1d764a) |
| `routes.route_destination.mirror_policy.cluster` | [routes.route_destination.mirror_policy.cluster](resources--route--reference--group-002.md#canonical-4c566544fb605b6d9fdcb15642d37b79a93d3dfa7716a8b8dfa7adf96a441cca) |
| `routes.route_destination.mirror_policy.cluster.kind` | [routes.route_destination.mirror_policy.cluster.kind](resources--route--reference--group-002.md#canonical-a9365021683c7ad5b2345b58ce78a6ba2eeae75d1e78df93b1f9a23ce859e083) |
| `routes.route_destination.mirror_policy.cluster.name` | [routes.route_destination.mirror_policy.cluster.name](resources--route--reference--group-002.md#canonical-236a2b47df18809075f396e2fb10baee9ffdb38523de0ce75b8c7a0c56d44c3b) |
| `routes.route_destination.mirror_policy.cluster.namespace` | [routes.route_destination.mirror_policy.cluster.namespace](resources--route--reference--group-002.md#canonical-49d0818284bc7e84e34d62963c1fb638de46545ca8532849581d403700b23699) |
| `routes.route_destination.mirror_policy.cluster.tenant` | [routes.route_destination.mirror_policy.cluster.tenant](resources--route--reference--group-002.md#canonical-e7e4710634b35daa3322daf79cc19b8137b8b9cd847a341bdcc033e03a4e8fe1) |
| `routes.route_destination.mirror_policy.cluster.uid` | [routes.route_destination.mirror_policy.cluster.uid](resources--route--reference--group-002.md#canonical-2c1e620d79adbfbf4d3d1aa9943284a481e287d2689074172168ee7751737695) |
| `routes.route_destination.mirror_policy.percent` | [routes.route_destination.mirror_policy.percent](resources--route--reference--group-002.md#canonical-691679b5159f8fcf20b216497dbdeb4493c48ad91d67a234d7ee1e5a1a023f8c) |
| `routes.route_destination.mirror_policy.percent.denominator` | [routes.route_destination.mirror_policy.percent.denominator](resources--route--reference--group-002.md#canonical-3a79a474fefe27cc00c706a7fac25bf7f2ba708337c71e3878dc43f804a5f3b9) |
| `routes.route_destination.mirror_policy.percent.numerator` | [routes.route_destination.mirror_policy.percent.numerator](resources--route--reference--group-002.md#canonical-98bd4b3e1e32ddfa823f0fe0c7cbd8a51f2eb19f2f3ac78b9fe0ca4a70124524) |
| `routes.route_destination.prefix_rewrite` | [routes.route_destination.prefix_rewrite](resources--route--reference--group-002.md#canonical-f0b6c4e5dd1a3cb9adfce1cb8e86ec9f54f8cd77afdd218a5f88e1ed43349a6a) |
| `routes.route_destination.priority` | [routes.route_destination.priority](resources--route--reference--group-002.md#canonical-be7c50cba3916f83a80785f185d6ba5d4374b718731aaf7516af776bc5c6034a) |
| `routes.route_destination.query_params` | [routes.route_destination.query_params](resources--route--reference--group-002.md#canonical-fa4bcb7df0d63878006914d53327ff7b14722607406a401e7509504585b11db3) |
| `routes.route_destination.query_params.remove_all_params` | [routes.route_destination.query_params.remove_all_params](resources--route--reference--group-002.md#canonical-bafbe1442fbc5fbb9b142bcc1cbab4bbfa65cac8387ead721162665922f3b5ee) |
| `routes.route_destination.query_params.replace_params` | [routes.route_destination.query_params.replace_params](resources--route--reference--group-002.md#canonical-b43d18a2354bdc9f0e8174ca7a90d41c4d2974110bbf7698c51415a70c4af37c) |
| `routes.route_destination.query_params.retain_all_params` | [routes.route_destination.query_params.retain_all_params](resources--route--reference--group-002.md#canonical-35729542fcc10afc25224b1408a5381cdb48e9b23fd788570fea11d98ef28850) |
| `routes.route_destination.regex_rewrite` | [routes.route_destination.regex_rewrite](resources--route--reference--group-002.md#canonical-4765cd2686bc20ff18d1c39d82da8f213fe9884c9734376407622740ee7271ee) |
| `routes.route_destination.regex_rewrite.pattern` | [routes.route_destination.regex_rewrite.pattern](resources--route--reference--group-002.md#canonical-359149b82c8e7926f4093829febbee9e636f11940599d2cb6b2d8e1739f648fd) |
| `routes.route_destination.regex_rewrite.substitution` | [routes.route_destination.regex_rewrite.substitution](resources--route--reference--group-002.md#canonical-de6ac97d0dfca38f9743ab1de405040ea8a897ac90f8a4e55bc9a5eb454f2827) |
| `routes.route_destination.retract_cluster` | [routes.route_destination.retract_cluster](resources--route--reference--group-002.md#canonical-df3e12476841dd3d4b4c0ed63a8ff0e815cf696086838b42bdfb294f549f1b64) |
| `routes.route_destination.retry_policy` | [routes.route_destination.retry_policy](resources--route--reference--group-002.md#canonical-3e623ebc2d0617b3dd4cf866587c7ad36228857f198efea55433187a33fb047c) |
| `routes.route_destination.retry_policy.back_off` | [routes.route_destination.retry_policy.back_off](resources--route--reference--group-002.md#canonical-bab4527c92789cdfde32b446a4667127203982a90dcb2b0bd3da3a9388e8322f) |
| `routes.route_destination.retry_policy.back_off.base_interval` | [routes.route_destination.retry_policy.back_off.base_interval](resources--route--reference--group-002.md#canonical-a7c8b0edec586a8d51a6ea8557aca08f21a62fdcc49b3fadf54bc4d59d8aef2d) |
| `routes.route_destination.retry_policy.back_off.max_interval` | [routes.route_destination.retry_policy.back_off.max_interval](resources--route--reference--group-002.md#canonical-36cf399f34a2880b346f3c07bbb4049f971f8511375795aa6a6554dd95c49641) |
| `routes.route_destination.retry_policy.num_retries` | [routes.route_destination.retry_policy.num_retries](resources--route--reference--group-002.md#canonical-fcdf4c1cf6d36101dc2b50e92ca12bdebb879890a4ad0c04ac8d21a2ef966c58) |
| `routes.route_destination.retry_policy.per_try_timeout` | [routes.route_destination.retry_policy.per_try_timeout](resources--route--reference--group-002.md#canonical-4166305d254c8dc131f9db965d14f9f441f87cf08dc4775c155254c11f0aecd1) |
| `routes.route_destination.retry_policy.retriable_status_codes` | [routes.route_destination.retry_policy.retriable_status_codes](resources--route--reference--group-002.md#canonical-a327040ee3576613588e02dece091dc35662cae63fdc58322e65bb62ff0d3ab0) |
| `routes.route_destination.retry_policy.retry_condition` | [routes.route_destination.retry_policy.retry_condition](resources--route--reference--group-002.md#canonical-2686bbd1ee5c103d51937b86e78dbb80a9995dfafad48f4f441fca237cd5155d) |
| `routes.route_destination.spdy_config` | [routes.route_destination.spdy_config](resources--route--reference--group-002.md#canonical-0c0e64b0730a3ab206cbc799fa0e7887b2a74923fe724f013645fead069d00d0) |
| `routes.route_destination.spdy_config.use_spdy` | [routes.route_destination.spdy_config.use_spdy](resources--route--reference--group-002.md#canonical-dbd7032f9b5096b186929adcc664ea55fc909c03098ea055d006de87b3e23fbb) |
| `routes.route_destination.timeout` | [routes.route_destination.timeout](resources--route--reference--group-002.md#canonical-0b976bff03ec2356095840dc4f6040ac6a92ce76578e1e2162ad21a0dc67e86b) |
| `routes.route_destination.web_socket_config` | [routes.route_destination.web_socket_config](resources--route--reference--group-002.md#canonical-d10aed8565d21029bb719cb924d0ed9608115d906e6ad4a3acbb9dfaf65dea7f) |
| `routes.route_destination.web_socket_config.use_websocket` | [routes.route_destination.web_socket_config.use_websocket](resources--route--reference--group-002.md#canonical-575617fc4b48bca7d6b97b1fa5c681c7ea448630fa2cd05ae3ee0a68184e95fd) |
| `routes.route_direct_response` | [routes.route_direct_response](resources--route--reference--group-002.md#canonical-b2c1b764e74f6e2cf405a410b96d25c56299c4959394fed00059af91811d2485) |
| `routes.route_direct_response.response_body_encoded` | [routes.route_direct_response.response_body_encoded](resources--route--reference--group-002.md#canonical-2acb323af657eb9fd7d625f931cda43391875d01ebca862defaa7bfd455b27fd) |
| `routes.route_direct_response.response_code` | [routes.route_direct_response.response_code](resources--route--reference--group-002.md#canonical-5950d0c627b65ef69bf8812934315c327062932d190b0b158c309ffa358dcfe2) |
| `routes.route_redirect` | [routes.route_redirect](resources--route--reference--group-002.md#canonical-9ffb37669778bfe7ee6a2df730dc37efc515dfe4044b11da8552185fab602fea) |
| `routes.route_redirect.host_redirect` | [routes.route_redirect.host_redirect](resources--route--reference--group-002.md#canonical-111233432cf1b9951a4a2e3a0c92e9895860ef98cc9322bacf4739f6cd24c37d) |
| `routes.route_redirect.path_redirect` | [routes.route_redirect.path_redirect](resources--route--reference--group-002.md#canonical-b96a3833b3ff06950581a2ef73465c90b2255230e3e8f638506f262b96822f3f) |
| `routes.route_redirect.prefix_rewrite` | [routes.route_redirect.prefix_rewrite](resources--route--reference--group-002.md#canonical-9fb5c4b10cfdc63e63e942ac9a7b9bfcc4f2717ab3bf24f4f3ee06d019961d34) |
| `routes.route_redirect.proto_redirect` | [routes.route_redirect.proto_redirect](resources--route--reference--group-002.md#canonical-95981b60c52fc0aad00ac4fef37365f25d1e8163cc457db3bf8630078d6de34b) |
| `routes.route_redirect.remove_all_params` | [routes.route_redirect.remove_all_params](resources--route--reference--group-003.md#canonical-5c59449e3b5a5b006e25df7dd172874d15c7c338104c5265e76806476ca99851) |
| `routes.route_redirect.replace_params` | [routes.route_redirect.replace_params](resources--route--reference--group-002.md#canonical-741a722d091c5501016e5a3b64b4d17c3b09947a588ee7f3129df663c5a9c781) |
| `routes.route_redirect.response_code` | [routes.route_redirect.response_code](resources--route--reference--group-002.md#canonical-9c8e8b2218c6e3eb633849aa56c7013586adb60de8f30003333320fc20e954b9) |
| `routes.route_redirect.retain_all_params` | [routes.route_redirect.retain_all_params](resources--route--reference--group-003.md#canonical-ef56de22d39c4b6040fdcf17188d0b503a6943d5c9868a39c5ddb27dcffce860) |
| `routes.service_policy` | [routes.service_policy](resources--route--reference--group-003.md#canonical-8beeaa8b5a4402111f28c39dd40a483729ac77497dd94e22d590418d842b7883) |
| `routes.service_policy.disable_spec` | [routes.service_policy.disable_spec](resources--route--reference--group-003.md#canonical-26bf8e4c61756301fe32b7cc8a901ea727a5e0a348c6fff88eb473ac98660b2b) |
| `routes.waf_exclusion_policy` | [routes.waf_exclusion_policy](resources--route--reference--group-003.md#canonical-821f7060a5b01bd0e708b132035d79d768449a7bb4d0fc9d8fd773daf3c4983d) |
| `routes.waf_exclusion_policy.name` | [routes.waf_exclusion_policy.name](resources--route--reference--group-003.md#canonical-3bfd4e5e79b4240d3e5f6df299342a775ebe0601b5ec006a18175c2e9d318f01) |
| `routes.waf_exclusion_policy.namespace` | [routes.waf_exclusion_policy.namespace](resources--route--reference--group-003.md#canonical-f8ae1d766cfa0a5e8b94e327b5601b927c86f444f647fa800e471da6ed8bc143) |
| `routes.waf_exclusion_policy.tenant` | [routes.waf_exclusion_policy.tenant](resources--route--reference--group-003.md#canonical-16cfba421611ee50889f0777b7fa340354968df4c5c7cb1238614958282a4175) |
| `routes.waf_type` | [routes.waf_type](resources--route--reference--group-003.md#canonical-81e5777777854566daf6e2aabd89ae4829620b02e0f241bd3a4b244fd43d9182) |
| `routes.waf_type.app_firewall` | [routes.waf_type.app_firewall](resources--route--reference--group-003.md#canonical-ad26da0f2fbdaa07556f8ab1ca94d0fe304246727d0b6570f22aab420939d8dd) |
| `routes.waf_type.app_firewall.app_firewall` | [routes.waf_type.app_firewall.app_firewall](resources--route--reference--group-003.md#canonical-ac927a9e466432b33b032a13f97f092ee5b46cfb748dafddf55cf2e027089ead) |
| `routes.waf_type.app_firewall.app_firewall.kind` | [routes.waf_type.app_firewall.app_firewall.kind](resources--route--reference--group-003.md#canonical-2983a08b554120240a192ba59d4ccab583c37ef16dd70a3604f87e44661a2043) |
| `routes.waf_type.app_firewall.app_firewall.name` | [routes.waf_type.app_firewall.app_firewall.name](resources--route--reference--group-003.md#canonical-97f9f258684ef833acf7e0420be877de7f79baa9490bfd988f203e8c7fe5eda8) |
| `routes.waf_type.app_firewall.app_firewall.namespace` | [routes.waf_type.app_firewall.app_firewall.namespace](resources--route--reference--group-003.md#canonical-cbb35b75244e1c1fcb9b067da9262dbdf20e0a457e757a588a3d040b0effa7ac) |
| `routes.waf_type.app_firewall.app_firewall.tenant` | [routes.waf_type.app_firewall.app_firewall.tenant](resources--route--reference--group-003.md#canonical-425c7064fb0d0715f7e7355ba97b2f0e7bbf6326f7d726bd166cb4c561d67af2) |
| `routes.waf_type.app_firewall.app_firewall.uid` | [routes.waf_type.app_firewall.app_firewall.uid](resources--route--reference--group-003.md#canonical-99c9974ac95612beb88adbf73af3c89c625c36d319d31f4da4ad818cbf69cd18) |
| `routes.waf_type.disable_waf` | [routes.waf_type.disable_waf](resources--route--reference--group-003.md#canonical-c50ce6d4bf3744fa7f9c96d717f36709555bd7593f100982d859fe78a136b0bd) |
| `routes.waf_type.inherit_waf` | [routes.waf_type.inherit_waf](resources--route--reference--group-003.md#canonical-6d9d9a730bd213796297278d2e8e7f668663a67e1e4fb7952fb433ce99819601) |
| `timeouts` | [timeouts](resources--route--reference--group-003.md#canonical-d5e4a8878e621a3a71b42f300db5c0fc81cfed7dac9a7c7d8cfedc58dc9b86be) |
| `timeouts.create` | [timeouts.create](resources--route--reference--group-003.md#canonical-5c06c6c469e7b7f53eac5c4e6a35e3d6362d5f804ca12421fd17c0d6501b52c1) |
| `timeouts.delete` | [timeouts.delete](resources--route--reference--group-003.md#canonical-3bdbf867dcf7b6f2315966f2f002144096846218325667b047770924aa23ed26) |
| `timeouts.read` | [timeouts.read](resources--route--reference--group-003.md#canonical-514d802565f0f8b466ae6bc25506e8a119e333044eeb0e19caab3bc5eb315254) |
| `timeouts.update` | [timeouts.update](resources--route--reference--group-003.md#canonical-1b4f78984d9fe878bf5d700c53a568fc136508dbdeac85e6008de73f7e1d3dac) |

<a id="canonical-435f6d8128de2e9b3c72a553cdfc6dd47a8888336da500dd7e237d6480d5997a"></a>

## Next pages — Property reference / 04890cbd4f28 / 12

- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [timeouts](resources--route--reference--group-003.md#canonical-806a5af33b8df63de7fc42ab58f766c42f08c0fd462d03b1e06b3e9ed42e23ac)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c51a9de9ba263ef102a7452ddca400565b9d65b6ee63e4d42a79a132081896f"></a>

## routes — routes / f2cae90dc028 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- routes

<a id="canonical-1fa4dfa72628ad53a1a3d8dd57a551d8435dd2c7c08e17e87bfb2c93e2621a51"></a>

Type: `"object"`. list nested block, Optional.

List of routes to match for incoming request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bot_defense_javascript_injection",
    "inherited_bot_defense_javascript_injection"),
  validators.ConflictingListObjectAttributes("inherited_waf_exclusion",
    "waf_exclusion_policy"),
  validators.ConflictingListObjectAttributes("route_destination",
    "route_direct_response"),
  validators.ConflictingListObjectAttributes("route_destination",
    "route_redirect"),
  validators.ConflictingListObjectAttributes("route_direct_response",
    "route_redirect")}
```

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

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-bea083a139cec51490b40b9b59d5a9101df3b8cc10376e1bdbe10cce9f5858b2"></a>

## Direct properties — routes / f2cae90dc028 / 3

- [bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-c67772aa819ebc76af3021d1835fc9f184b54804f13d02a69de46b9613506bca): complete subsection reference.

<a id="canonical-db74ab88546a70c3b5dd31a84281910f1ef1c7a2f5163cef38494ebd75d35a72"></a>

<a id="canonical-9c34f6040b0df3da43dde2908db82fd873ccdccef6a299eafb366130bed17c63"></a>

## disable_location_add property — routes / f2cae90dc028 / 4

Type: `"bool"`. Optional.

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

- [inherited_bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-9bdf2af1ff1ecf60aff81ed76cff0367f2dcb849dc784b8f2a5366b38b900e3f): complete subsection reference.

- [inherited_waf_exclusion](resources--route--reference--group-001.md#canonical-0ec91e105d9e99146c1612f6de9a9f6fd89deac94bbe6b690814445b4c59d007): complete subsection reference.

- [match](resources--route--reference--group-001.md#canonical-594bde7204dcf4bdbfd4cf2808be3010c78a18ec6bddf74e6b61a716479ccdf3): complete subsection reference.

- [request_cookies_to_add](resources--route--reference--group-001.md#canonical-6ad5316b8b92cb26e92b6aec85033fefa161cb013a3c01450620605977066c48): complete subsection reference.

<a id="canonical-c6560ab78fa9d8c06341610ea2fbbb2c934d38658e1e82f5e3974bd83d2a7b86"></a>

<a id="canonical-ee5a0daf171d6bd2026cb469a242c176d4bd2a2f5a52a15b3a964123927d12a3"></a>

## request_cookies_to_remove property — routes / f2cae90dc028 / 5

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

- [request_headers_to_add](resources--route--reference--group-001.md#canonical-d47213f6772f9be9e17b3859c8906e9d9017d3533122157b0e77084349aef91f): complete subsection reference.

<a id="canonical-807896249b94ceffd733ce3a7410f92535df7b50106761943cdde863ea9b84d3"></a>

<a id="canonical-7ecc33d316d5b05d249b36a9cb2eae04092545712ae1bac5b7bccee85b499d82"></a>

## request_headers_to_remove property — routes / f2cae90dc028 / 6

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

- [response_cookies_to_add](resources--route--reference--group-001.md#canonical-464e949935c94a1ff63a4aafb9943678d4019e7bad5b8f6dd87322599792ba10): complete subsection reference.

<a id="canonical-c06da782c3708cb3540e3c27983584fad0a45bb2664e56906b7e1a0390db5845"></a>

<a id="canonical-d77e29e338a29e115cf0400167b6268a750c860874b38954fd21e4ab7b24f89e"></a>

## response_cookies_to_remove property — routes / f2cae90dc028 / 7

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

- [response_headers_to_add](resources--route--reference--group-002.md#canonical-4f3cc5438aa982835bc4ef8277313209d669f04339d635509b9803149e3646a8): complete subsection reference.

<a id="canonical-89e6e0ce4839d5f6442ce1acedbcd8c48590a68c65797ab4483e2064c9779af1"></a>

<a id="canonical-af20acbc9e5ce9033a04f0e9d689d5c0e82d7214ba3bef02f451ce4a75f0e589"></a>

## response_headers_to_remove property — routes / f2cae90dc028 / 8

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

- [route_destination](resources--route--reference--group-002.md#canonical-2871daf806087e40bf977e51a547a7cbb4adffefe092c0bb3f12391a581c2574): complete subsection reference.

- [route_direct_response](resources--route--reference--group-002.md#canonical-0f7ee7a872556398c0ca553ec5fd68100e6292c851fb1b554e3f2b6b7b314d6f): complete subsection reference.

- [route_redirect](resources--route--reference--group-002.md#canonical-cc8df830ae3aec0989cb6b87862fe4ab8278de3c6b9e73cfefc8a0fd73b69209): complete subsection reference.

- [service_policy](resources--route--reference--group-003.md#canonical-2bd35375527a1e930c28cac7e72cf42e4a2fe36d743fe2a5d4dbfc63a6c81b8c): complete subsection reference.

- [waf_exclusion_policy](resources--route--reference--group-003.md#canonical-d7dfdc5c8a8e6b5bfc8075b9735ae0313e06ea42f789a3c37bb178cbc56c8081): complete subsection reference.

- [waf_type](resources--route--reference--group-003.md#canonical-7a68f2dc278917495dc0b199478b7b2527048260831acee90d85c0c54ead1eb1): complete subsection reference.

<a id="canonical-5fce58bb76967b644b298694cfba346af5527c640dffed3b63e9ff984ffee47a"></a>

## Next pages — routes / f2cae90dc028 / 9

- [routes.bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-c67772aa819ebc76af3021d1835fc9f184b54804f13d02a69de46b9613506bca)
- [routes.inherited_bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-9bdf2af1ff1ecf60aff81ed76cff0367f2dcb849dc784b8f2a5366b38b900e3f)
- [routes.inherited_waf_exclusion](resources--route--reference--group-001.md#canonical-0ec91e105d9e99146c1612f6de9a9f6fd89deac94bbe6b690814445b4c59d007)
- [routes.match](resources--route--reference--group-001.md#canonical-594bde7204dcf4bdbfd4cf2808be3010c78a18ec6bddf74e6b61a716479ccdf3)
- [routes.request_cookies_to_add](resources--route--reference--group-001.md#canonical-6ad5316b8b92cb26e92b6aec85033fefa161cb013a3c01450620605977066c48)
- [routes.request_headers_to_add](resources--route--reference--group-001.md#canonical-d47213f6772f9be9e17b3859c8906e9d9017d3533122157b0e77084349aef91f)
- [routes.response_cookies_to_add](resources--route--reference--group-001.md#canonical-464e949935c94a1ff63a4aafb9943678d4019e7bad5b8f6dd87322599792ba10)
- [routes.response_headers_to_add](resources--route--reference--group-002.md#canonical-4f3cc5438aa982835bc4ef8277313209d669f04339d635509b9803149e3646a8)
- [routes.route_destination](resources--route--reference--group-002.md#canonical-2871daf806087e40bf977e51a547a7cbb4adffefe092c0bb3f12391a581c2574)
- [routes.route_direct_response](resources--route--reference--group-002.md#canonical-0f7ee7a872556398c0ca553ec5fd68100e6292c851fb1b554e3f2b6b7b314d6f)
- [routes.route_redirect](resources--route--reference--group-002.md#canonical-cc8df830ae3aec0989cb6b87862fe4ab8278de3c6b9e73cfefc8a0fd73b69209)
- [routes.service_policy](resources--route--reference--group-003.md#canonical-2bd35375527a1e930c28cac7e72cf42e4a2fe36d743fe2a5d4dbfc63a6c81b8c)
- [routes.waf_exclusion_policy](resources--route--reference--group-003.md#canonical-d7dfdc5c8a8e6b5bfc8075b9735ae0313e06ea42f789a3c37bb178cbc56c8081)
- [routes.waf_type](resources--route--reference--group-003.md#canonical-7a68f2dc278917495dc0b199478b7b2527048260831acee90d85c0c54ead1eb1)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-c67772aa819ebc76af3021d1835fc9f184b54804f13d02a69de46b9613506bca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca4b7f79a2db6c7a536e000089dfe43b3b7d59626dbe8d0955e004d2c6ec25ae"></a>

## routes.bot_defense_javascript_injection — routes.bot_defense_javascript_injection / 785d59c5f860 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- routes.bot_defense_javascript_injection

<a id="canonical-94c0cafb9b6b1b95750032ded40f1199f5ded17f08e6bfc01460a5f4725695b6"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Javascript Injection Configuration for inline bot defense deployments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("javascript_tags")}
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
bot_defense_javascript_injection {
  # Configure direct properties listed below.
}
```

<a id="canonical-3a9eb1ce8920d9178b53d553d4b863864f343bee28bf0d04269d8b76e110fa60"></a>

## Direct properties — routes.bot_defense_javascript_injection / 785d59c5f860 / 3

<a id="canonical-3fc247b880b57b72e64ad5d61d4739c5af4bba7287424fc200aaa230f2ccb198"></a>

<a id="canonical-7deea11312a373310f958c48ff42d46de6a5ef969cf8fa8c68d2e1848cfa13b0"></a>

## javascript_location property — routes.bot_defense_javascript_injection / 785d59c5f860 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

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

- [javascript_tags](resources--route--reference--group-001.md#canonical-cb545ca64b1d821a693c856cbc6f641598633c851082c9cbbbd8153c62febb37): complete subsection reference.

<a id="canonical-2170e0457575c9c5a20c93df53a575a18a41fda47f83aabef169173a80654c8a"></a>

## Next pages — routes.bot_defense_javascript_injection / 785d59c5f860 / 5

- [routes.bot_defense_javascript_injection.javascript_tags](resources--route--reference--group-001.md#canonical-cb545ca64b1d821a693c856cbc6f641598633c851082c9cbbbd8153c62febb37)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-cb545ca64b1d821a693c856cbc6f641598633c851082c9cbbbd8153c62febb37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da69949b355f1ce781075c1ea8081f7093253ce0561f4f2910835471b03a7fca"></a>

## routes.bot_defense_javascript_injection.javascript_tags — routes.bot_defense_javascript_injection.javascript_tags / 4f33e0160a18 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-c67772aa819ebc76af3021d1835fc9f184b54804f13d02a69de46b9613506bca)
- routes.bot_defense_javascript_injection.javascript_tags

<a id="canonical-cb009e057c4bc5872b109925e4b50db75d30e077941035a12fda78b7b2ba9261"></a>

Type: `"object"`. list nested block, Optional.

Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot
Javascript should be added first.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("javascript_url")}
```

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

Terraform syntax:

```terraform
javascript_tags {
  # Configure direct properties listed below.
}
```

<a id="canonical-3839b40fdd3b01b2587f0054697d8e67e8ce45c8812fa191c0434d4cd0c0abf2"></a>

## Direct properties — routes.bot_defense_javascript_injection.javascript_tags / 4f33e0160a18 / 3

<a id="canonical-9d72669a1bf33d7842bfe8e660edc72cc2c7fb5dddafca733c4a07d4346ef4df"></a>

<a id="canonical-749e167d80e36b4816efe1d91e0989fc108b8d2c66a95c4dbc0ea34f8633821b"></a>

## javascript_url property — routes.bot_defense_javascript_injection.javascript_tags / 4f33e0160a18 / 4

Type: `"string"`. Optional.

Please enter the full URL (include domain and path), or relative path.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

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

- [tag_attributes](resources--route--reference--group-001.md#canonical-556a56fe5c9b2922426986639a8f8a23b67008d064c3ce37f83a3a33deb9c407): complete subsection reference.

<a id="canonical-c570ffc4eb72ae515a13e04d9b0dcbaafc74d41ddc72263e1ac173df7191ead4"></a>

## Next pages — routes.bot_defense_javascript_injection.javascript_tags / 4f33e0160a18 / 5

- [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes](resources--route--reference--group-001.md#canonical-556a56fe5c9b2922426986639a8f8a23b67008d064c3ce37f83a3a33deb9c407)
- [routes.bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-c67772aa819ebc76af3021d1835fc9f184b54804f13d02a69de46b9613506bca)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-556a56fe5c9b2922426986639a8f8a23b67008d064c3ce37f83a3a33deb9c407"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f8de58a92d93d115d7da8cf284bc35dd053a3bd5696dc614216a5cb6eedd7b7"></a>

## routes.bot_defense_javascript_injection.javascript_tags.tag_attributes — routes.bot_defense_javascript_injection.javascript_tags.tag_attributes / 8fb7aad2f80d / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-c67772aa819ebc76af3021d1835fc9f184b54804f13d02a69de46b9613506bca)
- [routes.bot_defense_javascript_injection.javascript_tags](resources--route--reference--group-001.md#canonical-cb545ca64b1d821a693c856cbc6f641598633c851082c9cbbbd8153c62febb37)
- routes.bot_defense_javascript_injection.javascript_tags.tag_attributes

<a id="canonical-b933bcf011bb43183a8a29de1f7f93d5188aa07f2fc85123d0f422464bce1b6a"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
tag_attributes {
  # Configure direct properties listed below.
}
```

<a id="canonical-65a69046467ace29bc3de412762503d2945fedf4eae08f0f9b4d248682970f71"></a>

## Direct properties — routes.bot_defense_javascript_injection.javascript_tags.tag_attributes / 8fb7aad2f80d / 3

<a id="canonical-c7507b72577deb4c5605c36124f246bd133ed228d699904b6f87958e1ed1b7fe"></a>

<a id="canonical-eeb7616ff4164d9bb088961435341e819d536e5d94a9f528630d111e713ec532"></a>

## javascript_tag property — routes.bot_defense_javascript_injection.javascript_tags.tag_attributes / 8fb7aad2f80d / 4

Type: `"string"`. Optional.

\[Enum:
JS\_ATTR\_ID|JS\_ATTR\_CID|JS\_ATTR\_CN|JS\_ATTR\_API\_DOMAIN|JS\_ATTR\_API\_URL|JS\_ATTR\_API\_PATH|JS\_ATTR\_ASYNC|JS\_ATTR\_DEFER\]
Select from one of the predefined tag attributes. Possible values are \`JS\_ATTR\_ID\`,
\`JS\_ATTR\_CID\`, \`JS\_ATTR\_CN\`, \`JS\_ATTR\_API\_DOMAIN\`, \`JS\_ATTR\_API\_URL\`,
\`JS\_ATTR\_API\_PATH\`, \`JS\_ATTR\_ASYNC\`, \`JS\_ATTR\_DEFER\`. Defaults to \`JS\_ATTR\_ID\`.

Upstream description:

Select from one of the predefined tag attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("JS_ATTR_ID",
    "JS_ATTR_CID",
    "JS_ATTR_CN",
    "JS_ATTR_API_DOMAIN",
    "JS_ATTR_API_URL",
    "JS_ATTR_API_PATH",
    "JS_ATTR_ASYNC",
    "JS_ATTR_DEFER"),
}
```

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

<a id="canonical-83c6a66ad06e60a576246f829c15215de2a1e929fdea48e7464dfd42eec25dad"></a>

<a id="canonical-fdeb3545d8f3fb5ac3f98d84a48454d1c5bd0895d405ca22f314fb8ff2d73001"></a>

## tag_value property — routes.bot_defense_javascript_injection.javascript_tags.tag_attributes / 8fb7aad2f80d / 5

Type: `"string"`. Optional.

Value. Add the tag attribute value.

Upstream description:

Add the tag attribute value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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

<a id="canonical-6c84c35ebdea9380658da75bf5ad1b37f77063eaa1b8e40e5dc06684b8912310"></a>

## Next pages — routes.bot_defense_javascript_injection.javascript_tags.tag_attributes / 8fb7aad2f80d / 6

- [routes.bot_defense_javascript_injection.javascript_tags](resources--route--reference--group-001.md#canonical-cb545ca64b1d821a693c856cbc6f641598633c851082c9cbbbd8153c62febb37)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-9bdf2af1ff1ecf60aff81ed76cff0367f2dcb849dc784b8f2a5366b38b900e3f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-138b7b50faebf44e7083b24c50b77558b300a9d03dee985e62f5410dbe007446"></a>

## routes.inherited_bot_defense_javascript_injection — routes.inherited_bot_defense_javascript_injection / 5c31f5cb52d2 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- routes.inherited_bot_defense_javascript_injection

<a id="canonical-cd53f990f5cc16be6a8ba1f7f525c557a3f60b0bfb7499b96cb11524a90ca807"></a>

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
inherited_bot_defense_javascript_injection = {}
```

<a id="canonical-48f7a0c64c403ded5fe0133fc718a367679d409a6d5e71647e36be1005041b0b"></a>

## Direct properties — routes.inherited_bot_defense_javascript_injection / 5c31f5cb52d2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-13c5b8645eeb9edf9aac84ddc993fe3311e49199c608f6cfadc6fb6e8a4358d9"></a>

## Next pages — routes.inherited_bot_defense_javascript_injection / 5c31f5cb52d2 / 4

- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-0ec91e105d9e99146c1612f6de9a9f6fd89deac94bbe6b690814445b4c59d007"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17f4ff68c17c79861684d81d70d0301ed8795dfb3b7299344a93a3041265a441"></a>

## routes.inherited_waf_exclusion — routes.inherited_waf_exclusion / 26bc97f45e47 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- routes.inherited_waf_exclusion

<a id="canonical-2aee3d63da6d7b758ce1b5bdf1a84f5a053e6a849c4bb1089e695bd22e7d2c8b"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
inherited_waf_exclusion = {}
```

<a id="canonical-499aa814f9c736223c54bb0f58f6446e4fcbcf28c04f0c5e09e669843a18e382"></a>

## Direct properties — routes.inherited_waf_exclusion / 26bc97f45e47 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-473ab612c6ef3231fade2b2dc5cdcf3eab443c474a957993e217082d2c0ea8a6"></a>

## Next pages — routes.inherited_waf_exclusion / 26bc97f45e47 / 4

- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-594bde7204dcf4bdbfd4cf2808be3010c78a18ec6bddf74e6b61a716479ccdf3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bcbb8fb27eaefb1fcf73943d1c8c666152b710ef097af7816bf492b68249598"></a>

## routes.match — routes.match / 579d7bb93a5a / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- routes.match

<a id="canonical-66b856fd32205a41f3a0395342c677c381f1f772e86e2769605f03574e2bd1c6"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
match {
  # Configure direct properties listed below.
}
```

<a id="canonical-5642b585fc7b6822207ab0a6d746644b69a4d39d781a31952834ffc50a5b21b8"></a>

## Direct properties — routes.match / 579d7bb93a5a / 3

- [headers](resources--route--reference--group-001.md#canonical-6d07ea39e206b2455a10051f3e62531a880f51d5fc5a575765d116998a2ece60): complete subsection reference.

<a id="canonical-70f91849c82960be70f0d34ea5415206ca2357b903fd597870f2e46f16804d90"></a>

<a id="canonical-2780dc6361650a1ab3381493c91fb190f81e810452df66654057af47d565ac37"></a>

## http_method property — routes.match / 579d7bb93a5a / 4

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

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

- [incoming_port](resources--route--reference--group-001.md#canonical-3ccf88921dca6a571fd732f0e4aa0cab67cf8c48da018070db3b2954a8f25c09): complete subsection reference.

- [path](resources--route--reference--group-001.md#canonical-42d9a6ea8f8fcfb67b6d77768e26a280a834857ec0870cbbff9dabd916b4f48f): complete subsection reference.

- [query_params](resources--route--reference--group-001.md#canonical-4e10ff2ed3eac616b7b8aadb7b92d677b8bb10f3506992a444199269c9625026): complete subsection reference.

<a id="canonical-78c236e6342122d37e04bd8e7826252e9cb07b78ddafed08416b8f4239154217"></a>

## Next pages — routes.match / 579d7bb93a5a / 5

- [routes.match.headers](resources--route--reference--group-001.md#canonical-6d07ea39e206b2455a10051f3e62531a880f51d5fc5a575765d116998a2ece60)
- [routes.match.incoming_port](resources--route--reference--group-001.md#canonical-3ccf88921dca6a571fd732f0e4aa0cab67cf8c48da018070db3b2954a8f25c09)
- [routes.match.path](resources--route--reference--group-001.md#canonical-42d9a6ea8f8fcfb67b6d77768e26a280a834857ec0870cbbff9dabd916b4f48f)
- [routes.match.query_params](resources--route--reference--group-001.md#canonical-4e10ff2ed3eac616b7b8aadb7b92d677b8bb10f3506992a444199269c9625026)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-6d07ea39e206b2455a10051f3e62531a880f51d5fc5a575765d116998a2ece60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e97689873825f61bd3fc4bb4edfa271481c30b440fca869ea3ddc87f2c0cfdc"></a>

## routes.match.headers — routes.match.headers / 4ccb6dbbc097 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.match](resources--route--reference--group-001.md#canonical-594bde7204dcf4bdbfd4cf2808be3010c78a18ec6bddf74e6b61a716479ccdf3)
- routes.match.headers

<a id="canonical-92a18cc6ef9abe90fd95d3f316a33fdcffeecb15437c7fed26e52e3fe818774a"></a>

Type: `"object"`. list nested block, Optional.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
    "regex")}
```

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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-8341b3962c2e659cbc5dc5899c5ef17a9ba63821c9642b431278e96faab3c5e9"></a>

## Direct properties — routes.match.headers / 4ccb6dbbc097 / 3

<a id="canonical-5099fd9ea8281358af2d5dfa728f347ffc5546c5f35d809075529ad34158e7c7"></a>

<a id="canonical-a394fcc1eddb452ccb9bfb119ae3ea5bcdf758ad59f5bed0edcdd20f6c8f994b"></a>

## exact property — routes.match.headers / 4ccb6dbbc097 / 4

Type: `"string"`. Optional.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

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

<a id="canonical-0f8f1b1313226d57e6660c7bc683269a05c55cd1e4ca62496ec0fba85f3c6fa1"></a>

<a id="canonical-1c337026fd86dfad2619b8a8582e4516bc61264f493be4f99c8159a5db77ad87"></a>

## invert_match property — routes.match.headers / 4ccb6dbbc097 / 5

Type: `"bool"`. Optional.

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

<a id="canonical-b3034d8a6917ba113db04a2961614414c2c49345e8c5d952d60ed2e1af447be1"></a>

<a id="canonical-c51d9302ec86416dd8909c59ea47c4bc2bf1f86649dc064e27de108c36d6e549"></a>

## name property — routes.match.headers / 4ccb6dbbc097 / 6

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-8666b35f5e5a3ec1570f5dd8c331fe18a7c33b7bfad4a2c7aa82a2c1d00ed6bc"></a>

<a id="canonical-80ffc73110acc5fe169aa2801c27e0bba105d3bc255b3b5c5c450500ef22171b"></a>

## presence property — routes.match.headers / 4ccb6dbbc097 / 7

Type: `"bool"`. Optional.

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

<a id="canonical-b234b50b5f958be467243819f9a6b37dd4d1098909c0c45c66c783c19b1c95cb"></a>

<a id="canonical-a56bf73d352a3da208dc06ecbd3ccfb3bf180fc58f53c2aac35e5d49d6b6773c"></a>

## regex property — routes.match.headers / 4ccb6dbbc097 / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

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

<a id="canonical-66e35375463e4e41e378a01bc741979b03d6ca89560d07ad84dee7803bb6fe52"></a>

## Next pages — routes.match.headers / 4ccb6dbbc097 / 9

- [routes.match](resources--route--reference--group-001.md#canonical-594bde7204dcf4bdbfd4cf2808be3010c78a18ec6bddf74e6b61a716479ccdf3)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-3ccf88921dca6a571fd732f0e4aa0cab67cf8c48da018070db3b2954a8f25c09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad0b88d7d200297bab0df62c14406b2fda4b117ac345b2567baefbe995a611fa"></a>

## routes.match.incoming_port — routes.match.incoming_port / b55d7a596aa9 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.match](resources--route--reference--group-001.md#canonical-594bde7204dcf4bdbfd4cf2808be3010c78a18ec6bddf74e6b61a716479ccdf3)
- routes.match.incoming_port

<a id="canonical-fadffe6f6b8ff95193f78e792093f39a45a093eb1542fd47b913b5ac53726510"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
incoming_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-952b8ec9dcfb30df4080e9ef0d7d6c04c72830c3ce0ab322b2fa52e115d39b85"></a>

## Direct properties — routes.match.incoming_port / b55d7a596aa9 / 3

- [no_port_match](resources--route--reference--group-001.md#canonical-dc3f0181b22328ce00c839d1be6bd1bf13713eaeebc5181018b67b61c33a818c): complete subsection reference.

<a id="canonical-5190550350e307bc4f1a77a01780f0edac370a99b236ab4e9111befb8dbdbc49"></a>

<a id="canonical-af22c3dd6358f2cb41a8e52bb031cf9e2c5e1639b7a0d73499005a1861bfa9df"></a>

## port property — routes.match.incoming_port / b55d7a596aa9 / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
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

<a id="canonical-684480b7d08e78468740b4b145b8fb75f85cfbff99f6bd17fd2c4c6a3f9c3b14"></a>

<a id="canonical-c891f59feca864d5373820d3b671bb5e3e278c19b6e94af9d0197ca4041e1e42"></a>

## port_ranges property — routes.match.incoming_port / b55d7a596aa9 / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

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

<a id="canonical-f8fdbff726602d62f0508255818419bd95393e5bae918f597894811030db8426"></a>

## Next pages — routes.match.incoming_port / b55d7a596aa9 / 6

- [routes.match.incoming_port.no_port_match](resources--route--reference--group-001.md#canonical-dc3f0181b22328ce00c839d1be6bd1bf13713eaeebc5181018b67b61c33a818c)
- [routes.match](resources--route--reference--group-001.md#canonical-594bde7204dcf4bdbfd4cf2808be3010c78a18ec6bddf74e6b61a716479ccdf3)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-dc3f0181b22328ce00c839d1be6bd1bf13713eaeebc5181018b67b61c33a818c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d88f975815059b8d41e9dbea3e21eea6100f76dc56d16b656167cbd1ab4d2074"></a>

## routes.match.incoming_port.no_port_match — routes.match.incoming_port.no_port_match / 94a17a310478 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.match](resources--route--reference--group-001.md#canonical-594bde7204dcf4bdbfd4cf2808be3010c78a18ec6bddf74e6b61a716479ccdf3)
- [routes.match.incoming_port](resources--route--reference--group-001.md#canonical-3ccf88921dca6a571fd732f0e4aa0cab67cf8c48da018070db3b2954a8f25c09)
- routes.match.incoming_port.no_port_match

<a id="canonical-a33c2f60761a88efa5245f2ebe0e3da44ef44fa71b0b80b9593eda8cc7099ccf"></a>

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
no_port_match = {}
```

<a id="canonical-7de7da69f915897bed1833fc1dd6ea30bf97de81229763578526aa1bd80f5842"></a>

## Direct properties — routes.match.incoming_port.no_port_match / 94a17a310478 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-54bff1c4490bf536bfe34d9f67064d8be61a8b4bd2750ea1b2879c96df3d8922"></a>

## Next pages — routes.match.incoming_port.no_port_match / 94a17a310478 / 4

- [routes.match.incoming_port](resources--route--reference--group-001.md#canonical-3ccf88921dca6a571fd732f0e4aa0cab67cf8c48da018070db3b2954a8f25c09)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-42d9a6ea8f8fcfb67b6d77768e26a280a834857ec0870cbbff9dabd916b4f48f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-744b0ff6e25c55784d3161a4365d967278a28bcfba0e449b5cf5e91808cdd8ef"></a>

## routes.match.path — routes.match.path / f2b6b18c7872 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.match](resources--route--reference--group-001.md#canonical-594bde7204dcf4bdbfd4cf2808be3010c78a18ec6bddf74e6b61a716479ccdf3)
- routes.match.path

<a id="canonical-b44fb2279d9ec892e82ab40443ec8797cfcf366d05d8afe1630dca89f380fd0a"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-bc86e7e12c1189ad151ba10428773f67c4a1876db13923c2d5939816818e93c0"></a>

## Direct properties — routes.match.path / f2b6b18c7872 / 3

<a id="canonical-fc2ce8204828d2515d5745d6e4c97574429531b840695648472578b1f80d1049"></a>

<a id="canonical-405bc7d05ccfcd783df7000f6705fb3a449c6c0bb7341e023edd0c7f1910abfb"></a>

## path property — routes.match.path / f2b6b18c7872 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-3dd953ee81d52702740321576baddd589fcf78d6d4184e9a2813eac556bd5e0c"></a>

<a id="canonical-017b61e2c60aa735f14d3c12ccd05a994a20a08f67ce4d69bc54536b792b0b41"></a>

## prefix property — routes.match.path / f2b6b18c7872 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-e6ccee598c35bf7585b25f616def04db4c4ee8c5a404f13befc3ae5c0bc5a6c1"></a>

<a id="canonical-d39060df0bb88f209f75952b60d8f60339cd062a794d2446ce48fcfb7646c1d1"></a>

## regex property — routes.match.path / f2b6b18c7872 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-f73d4675bdfeb08b66165eda88e939ec6edccc3bcf8e5b8fb7cd3e287e0b837d"></a>

## Next pages — routes.match.path / f2b6b18c7872 / 7

- [routes.match](resources--route--reference--group-001.md#canonical-594bde7204dcf4bdbfd4cf2808be3010c78a18ec6bddf74e6b61a716479ccdf3)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-4e10ff2ed3eac616b7b8aadb7b92d677b8bb10f3506992a444199269c9625026"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b41efac8052fbda8253a3d627e2b7c443ccb384ba2e11b2ac02fb1edebd9208f"></a>

## routes.match.query_params — routes.match.query_params / 923598763c04 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.match](resources--route--reference--group-001.md#canonical-594bde7204dcf4bdbfd4cf2808be3010c78a18ec6bddf74e6b61a716479ccdf3)
- routes.match.query_params

<a id="canonical-7982de1160eac5fe30f0cfbaab4bafe4042f4069f4d20b8460cd755821afcaeb"></a>

Type: `"object"`. list nested block, Optional.

Query Parameters. List of (key, value) query parameters.

Upstream description:

List of (key, value) query parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("exact",
    "regex")}
```

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

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-33ce8fa3d240d893bf44db4e0f60c69107404887c588b1f5e8c2bb94318cd986"></a>

## Direct properties — routes.match.query_params / 923598763c04 / 3

<a id="canonical-e334a032dde107e6fdc981cbe811bf8d4eb5f44d3f44736f204363f24bd20936"></a>

<a id="canonical-cdcdcbbb71f782bfc291f98bfe3b9ffece88130ccfd0a29c4e3bc87479e3a22f"></a>

## exact property — routes.match.query_params / 923598763c04 / 4

Type: `"string"`. Optional.

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

<a id="canonical-9873cfb71e6d5ab7f946921474c77edf942bfa367c6d27d5b995e45d1e0ffe86"></a>

<a id="canonical-6b6ba9e00907d510facc9fbbbfade1c1e4e52bafc3310de5994ca66cdc2b95dc"></a>

## key property — routes.match.query_params / 923598763c04 / 5

Type: `"string"`. Optional.

Query parameter key In the above example, assignee\_username is the key.

Upstream description:

Query parameter key In the above example, assignee\_username is the key.

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

<a id="canonical-6a2a73266c82d56c91d1d440afe5e0c09e8dfd037cf993599f5964928bd70b30"></a>

<a id="canonical-485401f5aa1ac093ea4836165dae6164ba6b6442761b2087a323c4b6dfb959d4"></a>

## regex property — routes.match.query_params / 923598763c04 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\] Regex match value for the query parameter key.

Upstream description:

Exclusive with \[exact\] Regex match value for the query parameter key.

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

<a id="canonical-86ae1989f126575c5658609e9a6578c8e12e26e91449fb9681afa1d678cde58f"></a>

## Next pages — routes.match.query_params / 923598763c04 / 7

- [routes.match](resources--route--reference--group-001.md#canonical-594bde7204dcf4bdbfd4cf2808be3010c78a18ec6bddf74e6b61a716479ccdf3)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-6ad5316b8b92cb26e92b6aec85033fefa161cb013a3c01450620605977066c48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3790d1678969bd3ab8e69ec5c0aa5daab1017f9a33557e2e29a0db1bae9008a2"></a>

## routes.request_cookies_to_add — routes.request_cookies_to_add / 49f6ac74e2f8 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- routes.request_cookies_to_add

<a id="canonical-27de8d8bfb56c062449e4ad98ecfa3d875099af52fc2a9575d584dd0cababc08"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
```

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

Terraform syntax:

```terraform
request_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-96622fc95e55515a8a3fbc56666a39782e8fea720d08c49a1c3c8a8289a65509"></a>

## Direct properties — routes.request_cookies_to_add / 49f6ac74e2f8 / 3

<a id="canonical-ebe3dbab488ca777eea820c368ac64f47b94741752bc6997d1076205339ea186"></a>

<a id="canonical-bed24c2fe22a82fa4dc20bb1cb86a183551a73aa6c8518c8efb251187e5b905c"></a>

## name property — routes.request_cookies_to_add / 49f6ac74e2f8 / 4

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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

<a id="canonical-138a1842381f3759d547273c98cba5036e192b9314ea82dccc8a2b62e02a69da"></a>

<a id="canonical-37a4eb14200ce8285c1991e11705426d2c1b5b66b3e759e4a640ff7a7ad9be08"></a>

## overwrite property — routes.request_cookies_to_add / 49f6ac74e2f8 / 5

Type: `"bool"`. Optional.

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

- [secret_value](resources--route--reference--group-001.md#canonical-1d251255880696b41c1e715d84e2f8982cf501842b58fd5cdac234bdcd48c76c): complete subsection reference.

<a id="canonical-78701ad630e610c51f81a5865aefd95ab6c7750cf2ff6aaeec16e8fc1edb762b"></a>

<a id="canonical-7656b3d8deb6e7eb490ecc6bd15b40d374c3000e6b1a5535cb28d8caa94eb7de"></a>

## value property — routes.request_cookies_to_add / 49f6ac74e2f8 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

<a id="canonical-a4540fb39f0e4dd7b0765ea82164c39679e65bb678cfbcc6a745da63784c0e46"></a>

## Next pages — routes.request_cookies_to_add / 49f6ac74e2f8 / 7

- [routes.request_cookies_to_add.secret_value](resources--route--reference--group-001.md#canonical-1d251255880696b41c1e715d84e2f8982cf501842b58fd5cdac234bdcd48c76c)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-1d251255880696b41c1e715d84e2f8982cf501842b58fd5cdac234bdcd48c76c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb8191437bddcbcf1cc6df29f635fdd4215bbcd7765e41245d131ba483af08af"></a>

## routes.request_cookies_to_add.secret_value — routes.request_cookies_to_add.secret_value / 02f39a388e42 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.request_cookies_to_add](resources--route--reference--group-001.md#canonical-6ad5316b8b92cb26e92b6aec85033fefa161cb013a3c01450620605977066c48)
- routes.request_cookies_to_add.secret_value

<a id="canonical-4005b88294eac9ae4ea07c15ab929e13f564e4a0bd1d4d97227dd3e01a3852cb"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-1ce6b9752b27361ca7dde91770eddcf590875913576876945a881922dd248e74"></a>

## Direct properties — routes.request_cookies_to_add.secret_value / 02f39a388e42 / 3

- [blindfold_secret_info](resources--route--reference--group-001.md#canonical-8370f224129f174c0c0b41b1a6a4c30d634073ddcf4607b9a21e8d8e86146a3d): complete subsection reference.

- [clear_secret_info](resources--route--reference--group-001.md#canonical-3e20a1471ed85562b736cd77e01d4e01e21f21d93c00bf782ac32d8d53c7cd8a): complete subsection reference.

<a id="canonical-16b1067a6ed1acb68b9feba143dbda75797ea3acb8a56e54cb315730601d3f77"></a>

## Next pages — routes.request_cookies_to_add.secret_value / 02f39a388e42 / 4

- [routes.request_cookies_to_add.secret_value.blindfold_secret_info](resources--route--reference--group-001.md#canonical-8370f224129f174c0c0b41b1a6a4c30d634073ddcf4607b9a21e8d8e86146a3d)
- [routes.request_cookies_to_add.secret_value.clear_secret_info](resources--route--reference--group-001.md#canonical-3e20a1471ed85562b736cd77e01d4e01e21f21d93c00bf782ac32d8d53c7cd8a)
- [routes.request_cookies_to_add](resources--route--reference--group-001.md#canonical-6ad5316b8b92cb26e92b6aec85033fefa161cb013a3c01450620605977066c48)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-8370f224129f174c0c0b41b1a6a4c30d634073ddcf4607b9a21e8d8e86146a3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8730c66e8d4770b4016db6915739d14b8d9f90d71154433360973332d9716a5e"></a>

## routes.request_cookies_to_add.secret_value.blindfold_secret_info — routes.request_cookies_to_add.secret_value.blindfold_secret_info / ef7af8a67d2c / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.request_cookies_to_add](resources--route--reference--group-001.md#canonical-6ad5316b8b92cb26e92b6aec85033fefa161cb013a3c01450620605977066c48)
- [routes.request_cookies_to_add.secret_value](resources--route--reference--group-001.md#canonical-1d251255880696b41c1e715d84e2f8982cf501842b58fd5cdac234bdcd48c76c)
- routes.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-b44ab3c7b1fdf97bec90261535d2133a66ae2925a6aa894c3eeea84cb976d1f5"></a>

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

<a id="canonical-0aaa3386658da6a8c1ed04e9258a818ced7767ec16993dcafb449ca7ad47eace"></a>

## Direct properties — routes.request_cookies_to_add.secret_value.blindfold_secret_info / ef7af8a67d2c / 3

<a id="canonical-5c47571f857f223fb73bc08684129fcf29c412be16b965c57f48851902afd11a"></a>

<a id="canonical-92dc8b2e0afbd5b7dadb0aa871204f44480f09ae4f9a43cee63d39c8b2a7033a"></a>

## decryption_provider property — routes.request_cookies_to_add.secret_value.blindfold_secret_info / ef7af8a67d2c / 4

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

<a id="canonical-d5b8d385db3a2f93528bc08f963dd71d2bc6f8f9a6921b67e7a1f81fad768031"></a>

<a id="canonical-b64f2419c6c94b23884fec7810d9d746ddc06e4fdd63283ecd68d7ac2eb95f58"></a>

## location property — routes.request_cookies_to_add.secret_value.blindfold_secret_info / ef7af8a67d2c / 5

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

<a id="canonical-b2853e6aaec806f22b34c655e6ad37b81c572c90728e0097c461ace103c92bcd"></a>

<a id="canonical-885a27444aaa586afe873974b3476ee7e4bdc025aaffd760574c61023b3c79d3"></a>

## store_provider property — routes.request_cookies_to_add.secret_value.blindfold_secret_info / ef7af8a67d2c / 6

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

<a id="canonical-2840c8e0ca2a119bd1ac7fb8697f6b468407fbcf16c12418117efdf4f4a407bd"></a>

## Next pages — routes.request_cookies_to_add.secret_value.blindfold_secret_info / ef7af8a67d2c / 7

- [routes.request_cookies_to_add.secret_value](resources--route--reference--group-001.md#canonical-1d251255880696b41c1e715d84e2f8982cf501842b58fd5cdac234bdcd48c76c)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-3e20a1471ed85562b736cd77e01d4e01e21f21d93c00bf782ac32d8d53c7cd8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-367899073e90b444cb14b182ee6a158175aaf797c0547df65af3b2a73ea36cf3"></a>

## routes.request_cookies_to_add.secret_value.clear_secret_info — routes.request_cookies_to_add.secret_value.clear_secret_info / ad46690213b0 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.request_cookies_to_add](resources--route--reference--group-001.md#canonical-6ad5316b8b92cb26e92b6aec85033fefa161cb013a3c01450620605977066c48)
- [routes.request_cookies_to_add.secret_value](resources--route--reference--group-001.md#canonical-1d251255880696b41c1e715d84e2f8982cf501842b58fd5cdac234bdcd48c76c)
- routes.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-7213a3c2022c5d67c37a79bb8b732b2c52842f976c01066245e8b01acad5badb"></a>

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

<a id="canonical-e8c5cdfbddd6393d4415b5a18eae284cd20b6aff44a9523014b22055318a6f81"></a>

## Direct properties — routes.request_cookies_to_add.secret_value.clear_secret_info / ad46690213b0 / 3

<a id="canonical-31e882f43e4b4cc422179f639fe40da5d0a1aafbc8d901d3bef085d6ae4c5da3"></a>

<a id="canonical-579ab95575984fc89c67de82e616109f8e636e8098e17eabfc59e2f5d9d38925"></a>

## provider_ref property — routes.request_cookies_to_add.secret_value.clear_secret_info / ad46690213b0 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-e2810fd764a7feb524f7c420bdfa56382c4a47312991a5692b0275bec044f769"></a>

<a id="canonical-8b2ff9f40cbdb51f8da905c5a097a8943d5b7576ec101c1790057b7d32bc03ce"></a>

## url property — routes.request_cookies_to_add.secret_value.clear_secret_info / ad46690213b0 / 5

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

<a id="canonical-f966de461a4e11f16614496f675097f9a3eb4473b8d43206565a4589c57e26d5"></a>

## Next pages — routes.request_cookies_to_add.secret_value.clear_secret_info / ad46690213b0 / 6

- [routes.request_cookies_to_add.secret_value](resources--route--reference--group-001.md#canonical-1d251255880696b41c1e715d84e2f8982cf501842b58fd5cdac234bdcd48c76c)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-d47213f6772f9be9e17b3859c8906e9d9017d3533122157b0e77084349aef91f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f45e199f54b0315a6d6f73a8dbfff57eb6ea52baa71f50b2e59db92db34a597"></a>

## routes.request_headers_to_add — routes.request_headers_to_add / 1af2a846b7ca / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- routes.request_headers_to_add

<a id="canonical-5e8c41fbc8d3b7ff01e349c10dcdd61032c456f3f7d566e5dd91b4253604708e"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP requests being sent towards upstream. Headers
specified at this level are applied before headers from the enclosing VirtualHost object level.

Upstream description:

Headers are key-value pairs to be added to HTTP requests being sent towards upstream. Headers
specified at this level are applied before headers from the enclosing VirtualHost object level.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
```

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

Terraform syntax:

```terraform
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-fbb88193363364723ca5d68e6fdc3d1c334bd0129895e127935fdac225a3a5f3"></a>

## Direct properties — routes.request_headers_to_add / 1af2a846b7ca / 3

<a id="canonical-6d532631a78ee46692a95d94b7f020c35dec7eb42a6d4b5ff95bb5ce78648c7f"></a>

<a id="canonical-d909a9449a1880d81f31d55ff6f09e4b3da1ed0c3b05be7be6d033a1d761502e"></a>

## append property — routes.request_headers_to_add / 1af2a846b7ca / 4

Type: `"bool"`. Optional.

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

<a id="canonical-3dbdd5922ff3198b302113137ea5f9b337819d8e50b6ee9c9ade6e55bb60b0b1"></a>

<a id="canonical-8251e83d470715075c6577feea664bf89a8d212f1e53aea795680d0cfaa72449"></a>

## name property — routes.request_headers_to_add / 1af2a846b7ca / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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

- [secret_value](resources--route--reference--group-001.md#canonical-eb65012dca65b77784593004366737707611efb2323c3f8a57b72baab5c208f5): complete subsection reference.

<a id="canonical-3e99e19390e09b0d5514545b3dd67a3292407cb4435e72b4beab7d4efbe18299"></a>

<a id="canonical-074baa631e088c5cab9b99bc8ce757c933bd81c4406a37bc286cfe460224c367"></a>

## value property — routes.request_headers_to_add / 1af2a846b7ca / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

<a id="canonical-602bbbc6af80556452490e165963f98e1b121085f4f7d6da45c79dc3a6749623"></a>

## Next pages — routes.request_headers_to_add / 1af2a846b7ca / 7

- [routes.request_headers_to_add.secret_value](resources--route--reference--group-001.md#canonical-eb65012dca65b77784593004366737707611efb2323c3f8a57b72baab5c208f5)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-eb65012dca65b77784593004366737707611efb2323c3f8a57b72baab5c208f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-239c8f09a5b8ebb72620028b01ef1d70bacec50340a88bad86ae7e8eea686c6d"></a>

## routes.request_headers_to_add.secret_value — routes.request_headers_to_add.secret_value / 9fb601fd0b4b / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.request_headers_to_add](resources--route--reference--group-001.md#canonical-d47213f6772f9be9e17b3859c8906e9d9017d3533122157b0e77084349aef91f)
- routes.request_headers_to_add.secret_value

<a id="canonical-5288358233be7aa12d0e66909f59ce2e3ff87120213212c2dd5a5d165c654746"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-d58c7992b9116977c71bbac47aad296e0debd7a932728be1717bccbe73a58b2b"></a>

## Direct properties — routes.request_headers_to_add.secret_value / 9fb601fd0b4b / 3

- [blindfold_secret_info](resources--route--reference--group-001.md#canonical-0a124d9ad2adaa83ee9a3dc8791f0a58a1a6edf082f9aae1e0640d2a1dcb4c8b): complete subsection reference.

- [clear_secret_info](resources--route--reference--group-001.md#canonical-136c4f59cfe1ac7f3352e9e3c6daecaea323dbf226e412e281c68a6ab60d09ee): complete subsection reference.

<a id="canonical-c9ee6b459901defd8d327d6604610f31f53ff27807a6abe96a6c37e69da93b43"></a>

## Next pages — routes.request_headers_to_add.secret_value / 9fb601fd0b4b / 4

- [routes.request_headers_to_add.secret_value.blindfold_secret_info](resources--route--reference--group-001.md#canonical-0a124d9ad2adaa83ee9a3dc8791f0a58a1a6edf082f9aae1e0640d2a1dcb4c8b)
- [routes.request_headers_to_add.secret_value.clear_secret_info](resources--route--reference--group-001.md#canonical-136c4f59cfe1ac7f3352e9e3c6daecaea323dbf226e412e281c68a6ab60d09ee)
- [routes.request_headers_to_add](resources--route--reference--group-001.md#canonical-d47213f6772f9be9e17b3859c8906e9d9017d3533122157b0e77084349aef91f)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-0a124d9ad2adaa83ee9a3dc8791f0a58a1a6edf082f9aae1e0640d2a1dcb4c8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b1db79879cba5105bf3828a7d23ac7afa10c2800d08bc971aa11a4c8d6a8e7b"></a>

## routes.request_headers_to_add.secret_value.blindfold_secret_info — routes.request_headers_to_add.secret_value.blindfold_secret_info / b26d98faff3d / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.request_headers_to_add](resources--route--reference--group-001.md#canonical-d47213f6772f9be9e17b3859c8906e9d9017d3533122157b0e77084349aef91f)
- [routes.request_headers_to_add.secret_value](resources--route--reference--group-001.md#canonical-eb65012dca65b77784593004366737707611efb2323c3f8a57b72baab5c208f5)
- routes.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-09eb76b0397b9e3008decba1fd3c64d1f81592d69552e37c5f4865450d054021"></a>

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

<a id="canonical-0d4c75eb001ebd194fb022152ce7f70cc260df4a7ee4d24bd294ebd967bb2635"></a>

## Direct properties — routes.request_headers_to_add.secret_value.blindfold_secret_info / b26d98faff3d / 3

<a id="canonical-4a4d902f9741fa656cc6578328e427f505f85f1ab65be506528e0feb75532ac4"></a>

<a id="canonical-1ed60ec38d838a30425c465f19e74c85229dbeceb07d343d94b5b4a4c3b903e5"></a>

## decryption_provider property — routes.request_headers_to_add.secret_value.blindfold_secret_info / b26d98faff3d / 4

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

<a id="canonical-a22eb163303d2a416692cc24cfa58bc1ccb6b5b05a23649c19d5aaab6c0c4e15"></a>

<a id="canonical-3888f9b602347f0daccadf6fbe4d4d5aa656ebca50c9a893f954e853b1b388c3"></a>

## location property — routes.request_headers_to_add.secret_value.blindfold_secret_info / b26d98faff3d / 5

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

<a id="canonical-f72721aa7e8c7f7e3f094e87b55dac0cd9644439b77c2278b5756f87f6b059fb"></a>

<a id="canonical-35e75811a3c7f2ad914b1b1566a4eadfd2fa63f5564f7f10953b27db996fc363"></a>

## store_provider property — routes.request_headers_to_add.secret_value.blindfold_secret_info / b26d98faff3d / 6

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

<a id="canonical-6ae421133d09c09fed9fdca398d1f46a2630fec392ea92096525855825c6e520"></a>

## Next pages — routes.request_headers_to_add.secret_value.blindfold_secret_info / b26d98faff3d / 7

- [routes.request_headers_to_add.secret_value](resources--route--reference--group-001.md#canonical-eb65012dca65b77784593004366737707611efb2323c3f8a57b72baab5c208f5)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-136c4f59cfe1ac7f3352e9e3c6daecaea323dbf226e412e281c68a6ab60d09ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bed2353654e61825deea2ac971134732c9a806da5cec7a1eca8994bf713c3d1b"></a>

## routes.request_headers_to_add.secret_value.clear_secret_info — routes.request_headers_to_add.secret_value.clear_secret_info / e739595c4063 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.request_headers_to_add](resources--route--reference--group-001.md#canonical-d47213f6772f9be9e17b3859c8906e9d9017d3533122157b0e77084349aef91f)
- [routes.request_headers_to_add.secret_value](resources--route--reference--group-001.md#canonical-eb65012dca65b77784593004366737707611efb2323c3f8a57b72baab5c208f5)
- routes.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-a0b57f52b88da5409685e96d143b852be43d6488625e992117611c29ed0bf273"></a>

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

<a id="canonical-9e7f20a14f0412f863f1b42e3c0e3cfcc8f0f17d35c912fd8690aac844c8e7d4"></a>

## Direct properties — routes.request_headers_to_add.secret_value.clear_secret_info / e739595c4063 / 3

<a id="canonical-7b2f55e67c1d51baf6ea23d4e153abf1301b1fc8dc9b4ce88e8f5a0646ebf981"></a>

<a id="canonical-59a4cc973103ad6300ba2d294cf9ecc5d827e52dc7eef9077d47a511b4724c3f"></a>

## provider_ref property — routes.request_headers_to_add.secret_value.clear_secret_info / e739595c4063 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-ec2ce7edaf1ceffd38fa18fc0800419d0b00f09059e0c1ada5de9b0664bc54d7"></a>

<a id="canonical-b54e044a376de62ca4abe0b1c7a5e508c1a6f126a4ab5c7cf9384665e76658d0"></a>

## url property — routes.request_headers_to_add.secret_value.clear_secret_info / e739595c4063 / 5

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

<a id="canonical-99554384a72592316527c903413c2eea764fe39f6b8908c26a1b135f73150793"></a>

## Next pages — routes.request_headers_to_add.secret_value.clear_secret_info / e739595c4063 / 6

- [routes.request_headers_to_add.secret_value](resources--route--reference--group-001.md#canonical-eb65012dca65b77784593004366737707611efb2323c3f8a57b72baab5c208f5)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-464e949935c94a1ff63a4aafb9943678d4019e7bad5b8f6dd87322599792ba10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10e4bf8585c58cf7a73e95fffdbb2af3d9f3df3b1d8d56a90d72824a686943e8"></a>

## routes.response_cookies_to_add — routes.response_cookies_to_add / 96936ea27276 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- routes.response_cookies_to_add

<a id="canonical-07fd928eb654ed5f287820ef5b21f9d944cbea7f27814b89add5437e1184f5fc"></a>

Type: `"object"`. list nested block, Optional.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_domain",
    "ignore_domain"),
  validators.ConflictingListObjectAttributes("add_expiry",
    "ignore_expiry"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_partitioned",
    "ignore_partitioned"),
  validators.ConflictingListObjectAttributes("add_path",
    "ignore_path"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "secret_value"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "value"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
```

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

Terraform syntax:

```terraform
response_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-ce4af94ea470ef7e86b9dc403729072b3a318f8d56174ca6e7c8f6543a1df1aa"></a>

## Direct properties — routes.response_cookies_to_add / 96936ea27276 / 3

<a id="canonical-6cc7e58cfb7c93062b8ba5138f29c77ea3a1228553a0be0ebae64f329aa3bfdd"></a>

<a id="canonical-55e61f2cf101c39dc440ea5474d7295f85a7d68fab8a73ca1f7307a12dadebb2"></a>

## add_domain property — routes.response_cookies_to_add / 96936ea27276 / 4

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

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

<a id="canonical-b5dbd1e904cbea620d964cd54940ba79c8ded3607459c8720b86dba3936edec8"></a>

<a id="canonical-3f1e4f73df48c2fda98984f2839f3a3e4590dff51d2a104209616825c9d023a1"></a>

## add_expiry property — routes.response_cookies_to_add / 96936ea27276 / 5

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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

- [add_httponly](resources--route--reference--group-001.md#canonical-53a46f631a124de3d0781735b81a418f9c432ba0630306349c448b5a637ffe69): complete subsection reference.

- [add_partitioned](resources--route--reference--group-001.md#canonical-3f1c7b6b0e861040d377cac8098dd13859ac43e337bd11826ac2724333be29a4): complete subsection reference.

<a id="canonical-a4a045a1f09c68ed25ff658e2a6b9d861366af77ac9b50dbe3b0c70e353e639b"></a>

<a id="canonical-0bbcebc9aea16e9853d2fd51310e450cc144faed13d08d4fec41042282c4ee69"></a>

## add_path property — routes.response_cookies_to_add / 96936ea27276 / 6

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

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

- [add_secure](resources--route--reference--group-001.md#canonical-dfad3063c52fd7f8aa85af017ddf5c4d4a69ed21a8ea5b1eb138a7f8622242c3): complete subsection reference.

- [ignore_domain](resources--route--reference--group-001.md#canonical-a9ee8c52ba4b37f296fa46796121e79a11a41037a59424a15123fb37c3186bb0): complete subsection reference.

- [ignore_expiry](resources--route--reference--group-001.md#canonical-27f3b25194c9d1d26e15c4b2e54bcad3136578c2c77dbb7e0b445d5eed4e71df): complete subsection reference.

- [ignore_httponly](resources--route--reference--group-001.md#canonical-0096281b2d7af1d7e834d701c6d8ba44459195dd78963e72d47f76a2992fb50c): complete subsection reference.

- [ignore_max_age](resources--route--reference--group-002.md#canonical-0f9ac1fd52b2fa3e66e2dcde7aa5cec053a8437e62e575cfbd464af5a20968a1): complete subsection reference.

- [ignore_partitioned](resources--route--reference--group-002.md#canonical-10a79aff8c4b97ddb8bd4b4b2e4fb7d449626f85f8411c3cbec6f83d695bac29): complete subsection reference.

- [ignore_path](resources--route--reference--group-002.md#canonical-0e2bbed26355268415119b2b72bcf7344f111b6a37215d66872e06a5f7e2d6c9): complete subsection reference.

- [ignore_samesite](resources--route--reference--group-002.md#canonical-8155b429727fc99ec50cfce18ffa6ccd6902bb377aadb3722df45500b413f995): complete subsection reference.

- [ignore_secure](resources--route--reference--group-002.md#canonical-d02f5ffbe7da1039980feba567c88188cf7a7468463939858fd6519048ff1e99): complete subsection reference.

- [ignore_value](resources--route--reference--group-002.md#canonical-cf6ab0b77c20c718ab754990a88444f8e735bb09b80f28e620e1f1b83d2fb6ea): complete subsection reference.

<a id="canonical-ba7212e6b168140d9342821166570975de53e0bc4793fc7aa16a00ce27ce829a"></a>

<a id="canonical-deb8bbbf469cd9bad6a3decaf8ca0fc4018c5cbfb110633ac1a6808ea8e4d0d6"></a>

## max_age_value property — routes.response_cookies_to_add / 96936ea27276 / 7

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

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

<a id="canonical-a214c8320a34ed82908d144c456e8367c7d9d8da58efaad7bf56e687445eea9e"></a>

<a id="canonical-d6a9b6a12a87ebd198697456780af251b14bf7a039bc0d07170f578a1f32c67e"></a>

## name property — routes.response_cookies_to_add / 96936ea27276 / 8

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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

<a id="canonical-9db105956b7f874ce8aa19b36d4d95fa5ed969e2432208c281d72b4371245272"></a>

<a id="canonical-1833c5e7da0de6fd9873a04e6bc373dfe7ac9a415bd2a5aa65f11dde95981cd9"></a>

## overwrite property — routes.response_cookies_to_add / 96936ea27276 / 9

Type: `"bool"`. Optional.

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

- [samesite_lax](resources--route--reference--group-002.md#canonical-b22544525e03863f27d9df506cc0c07f1e18fcb690e0543c8cdb5a99ff09d5a6): complete subsection reference.

- [samesite_none](resources--route--reference--group-002.md#canonical-bc7c2a00c3dcaf580fd186385a031458414e24a6888f258a45a30f7ff2e9d899): complete subsection reference.

- [samesite_strict](resources--route--reference--group-002.md#canonical-f7a26a0a3907578dac0eee34ebe152e3b57b76e29fb636eb553acaf28fa3be79): complete subsection reference.

- [secret_value](resources--route--reference--group-002.md#canonical-138d466099a3a3b72fc65d7cb5eb691af770b4a2913f9aad22bb9f40fd8086c0): complete subsection reference.

<a id="canonical-57706fa921fe841ccd992aee0df8434df4945b5ead879cc310a58a32a4fbe31f"></a>

<a id="canonical-780befd4cd8d0d18911515168adf091b71b9199f3de4733fe51733a4f8c1af58"></a>

## value property — routes.response_cookies_to_add / 96936ea27276 / 10

Type: `"string"`. Optional.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

<a id="canonical-80b7229c41baf770b85e5894a5b4d54fcb0b0d26f4bd392be9d6dc20032e19f6"></a>

## Next pages — routes.response_cookies_to_add / 96936ea27276 / 11

- [routes.response_cookies_to_add.add_httponly](resources--route--reference--group-001.md#canonical-53a46f631a124de3d0781735b81a418f9c432ba0630306349c448b5a637ffe69)
- [routes.response_cookies_to_add.add_partitioned](resources--route--reference--group-001.md#canonical-3f1c7b6b0e861040d377cac8098dd13859ac43e337bd11826ac2724333be29a4)
- [routes.response_cookies_to_add.add_secure](resources--route--reference--group-001.md#canonical-dfad3063c52fd7f8aa85af017ddf5c4d4a69ed21a8ea5b1eb138a7f8622242c3)
- [routes.response_cookies_to_add.ignore_domain](resources--route--reference--group-001.md#canonical-a9ee8c52ba4b37f296fa46796121e79a11a41037a59424a15123fb37c3186bb0)
- [routes.response_cookies_to_add.ignore_expiry](resources--route--reference--group-001.md#canonical-27f3b25194c9d1d26e15c4b2e54bcad3136578c2c77dbb7e0b445d5eed4e71df)
- [routes.response_cookies_to_add.ignore_httponly](resources--route--reference--group-001.md#canonical-0096281b2d7af1d7e834d701c6d8ba44459195dd78963e72d47f76a2992fb50c)
- [routes.response_cookies_to_add.ignore_max_age](resources--route--reference--group-002.md#canonical-0f9ac1fd52b2fa3e66e2dcde7aa5cec053a8437e62e575cfbd464af5a20968a1)
- [routes.response_cookies_to_add.ignore_partitioned](resources--route--reference--group-002.md#canonical-10a79aff8c4b97ddb8bd4b4b2e4fb7d449626f85f8411c3cbec6f83d695bac29)
- [routes.response_cookies_to_add.ignore_path](resources--route--reference--group-002.md#canonical-0e2bbed26355268415119b2b72bcf7344f111b6a37215d66872e06a5f7e2d6c9)
- [routes.response_cookies_to_add.ignore_samesite](resources--route--reference--group-002.md#canonical-8155b429727fc99ec50cfce18ffa6ccd6902bb377aadb3722df45500b413f995)
- [routes.response_cookies_to_add.ignore_secure](resources--route--reference--group-002.md#canonical-d02f5ffbe7da1039980feba567c88188cf7a7468463939858fd6519048ff1e99)
- [routes.response_cookies_to_add.ignore_value](resources--route--reference--group-002.md#canonical-cf6ab0b77c20c718ab754990a88444f8e735bb09b80f28e620e1f1b83d2fb6ea)
- [routes.response_cookies_to_add.samesite_lax](resources--route--reference--group-002.md#canonical-b22544525e03863f27d9df506cc0c07f1e18fcb690e0543c8cdb5a99ff09d5a6)
- [routes.response_cookies_to_add.samesite_none](resources--route--reference--group-002.md#canonical-bc7c2a00c3dcaf580fd186385a031458414e24a6888f258a45a30f7ff2e9d899)
- [routes.response_cookies_to_add.samesite_strict](resources--route--reference--group-002.md#canonical-f7a26a0a3907578dac0eee34ebe152e3b57b76e29fb636eb553acaf28fa3be79)
- [routes.response_cookies_to_add.secret_value](resources--route--reference--group-002.md#canonical-138d466099a3a3b72fc65d7cb5eb691af770b4a2913f9aad22bb9f40fd8086c0)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-53a46f631a124de3d0781735b81a418f9c432ba0630306349c448b5a637ffe69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2605e02d3bcae6856537b4c45da4d3a4828b88b3e2489165c0a379edb00307fa"></a>

## routes.response_cookies_to_add.add_httponly — routes.response_cookies_to_add.add_httponly / 02953dbdb294 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.response_cookies_to_add](resources--route--reference--group-001.md#canonical-464e949935c94a1ff63a4aafb9943678d4019e7bad5b8f6dd87322599792ba10)
- routes.response_cookies_to_add.add_httponly

<a id="canonical-f767477760ba00b8c4d72d479ac207aff37ddcee5c7749a53ed65a50e1d110eb"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
add_httponly = {}
```

<a id="canonical-f8410b58892af5de1f1edad85e9a188b1e094d3a7323247bc1746185b0612b25"></a>

## Direct properties — routes.response_cookies_to_add.add_httponly / 02953dbdb294 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-51c8b0c2c2b0d9755ca3e8c84bc625f4a9b69909a5a0ed585f39bfef89e4ef55"></a>

## Next pages — routes.response_cookies_to_add.add_httponly / 02953dbdb294 / 4

- [routes.response_cookies_to_add](resources--route--reference--group-001.md#canonical-464e949935c94a1ff63a4aafb9943678d4019e7bad5b8f6dd87322599792ba10)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-3f1c7b6b0e861040d377cac8098dd13859ac43e337bd11826ac2724333be29a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d612cb39ccfcbc412e334bff2d386dcafc27ae92f8fb7e84deffef185e4ee281"></a>

## routes.response_cookies_to_add.add_partitioned — routes.response_cookies_to_add.add_partitioned / 374bc0817fe7 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.response_cookies_to_add](resources--route--reference--group-001.md#canonical-464e949935c94a1ff63a4aafb9943678d4019e7bad5b8f6dd87322599792ba10)
- routes.response_cookies_to_add.add_partitioned

<a id="canonical-e249f9712e59a25b9811416a0cbb02777e491d7c0b433b8d8f612914accc08ba"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
add_partitioned = {}
```

<a id="canonical-8ed0e679435d392bb74c9363018cc790f7ecc92f11b3f10330e0c983dbc8c49a"></a>

## Direct properties — routes.response_cookies_to_add.add_partitioned / 374bc0817fe7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c9d9013416e8b1e5e6ad75c41392df53979427a641d0863ea3e5d044c530ebe6"></a>

## Next pages — routes.response_cookies_to_add.add_partitioned / 374bc0817fe7 / 4

- [routes.response_cookies_to_add](resources--route--reference--group-001.md#canonical-464e949935c94a1ff63a4aafb9943678d4019e7bad5b8f6dd87322599792ba10)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-dfad3063c52fd7f8aa85af017ddf5c4d4a69ed21a8ea5b1eb138a7f8622242c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-094ba165ff08866f188037e0a8cc9d61e389b6aaef19d1f4b8f60e30774d2499"></a>

## routes.response_cookies_to_add.add_secure — routes.response_cookies_to_add.add_secure / 7fcbfa4fe76b / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.response_cookies_to_add](resources--route--reference--group-001.md#canonical-464e949935c94a1ff63a4aafb9943678d4019e7bad5b8f6dd87322599792ba10)
- routes.response_cookies_to_add.add_secure

<a id="canonical-8882192c917b12fac9d0f80ac6919205fdf0f786df2ed60f25d8f2f03ee684e8"></a>

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
add_secure = {}
```

<a id="canonical-720c63a7f81f95a534c73fa09d4f9dbc90602b5ee84d244e539e6b39382e7cb4"></a>

## Direct properties — routes.response_cookies_to_add.add_secure / 7fcbfa4fe76b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8f31b5941c957ee26a9211e4f66c2c967058f76cac0f85a2e2d0d5a38a2dbe92"></a>

## Next pages — routes.response_cookies_to_add.add_secure / 7fcbfa4fe76b / 4

- [routes.response_cookies_to_add](resources--route--reference--group-001.md#canonical-464e949935c94a1ff63a4aafb9943678d4019e7bad5b8f6dd87322599792ba10)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-a9ee8c52ba4b37f296fa46796121e79a11a41037a59424a15123fb37c3186bb0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71c48ac83c8f61dd2f34bb86f03f1da2b16427058d6a5cf502d2608801844d31"></a>

## routes.response_cookies_to_add.ignore_domain — routes.response_cookies_to_add.ignore_domain / 6dbd8e066a09 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.response_cookies_to_add](resources--route--reference--group-001.md#canonical-464e949935c94a1ff63a4aafb9943678d4019e7bad5b8f6dd87322599792ba10)
- routes.response_cookies_to_add.ignore_domain

<a id="canonical-3aac7b78166731a78fab42a4e8e0fc3ff3a99bd0c53557506c9ba0c1ff86a9bb"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_domain = {}
```

<a id="canonical-0ec1a97dd9feaacb6909ed0e42783d13b55d706a118ff7748be8529df7eac584"></a>

## Direct properties — routes.response_cookies_to_add.ignore_domain / 6dbd8e066a09 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2b361f7dc54f86ff31e5e051025d05fb9b6d1a1bd9aa13feabb41a287ef736ef"></a>

## Next pages — routes.response_cookies_to_add.ignore_domain / 6dbd8e066a09 / 4

- [routes.response_cookies_to_add](resources--route--reference--group-001.md#canonical-464e949935c94a1ff63a4aafb9943678d4019e7bad5b8f6dd87322599792ba10)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-27f3b25194c9d1d26e15c4b2e54bcad3136578c2c77dbb7e0b445d5eed4e71df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71a61b15a4553cfe78c19bcd9811033ab1cbdea62d08cfec93d4c94e2221919a"></a>

## routes.response_cookies_to_add.ignore_expiry — routes.response_cookies_to_add.ignore_expiry / dec737c072db / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.response_cookies_to_add](resources--route--reference--group-001.md#canonical-464e949935c94a1ff63a4aafb9943678d4019e7bad5b8f6dd87322599792ba10)
- routes.response_cookies_to_add.ignore_expiry

<a id="canonical-2bd41275e146d4599b607af35ab924eef44b500178405f1cea65fa811075156f"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_expiry = {}
```

<a id="canonical-87c824a778eb18290b78e46fa2a7a6ac73e5980f6a532383ec2083fbb81cc156"></a>

## Direct properties — routes.response_cookies_to_add.ignore_expiry / dec737c072db / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64860dd7e0c88950c3f1e20f584b4d1fddd93145802fa3f49ead3c0c6e2840d9"></a>

## Next pages — routes.response_cookies_to_add.ignore_expiry / dec737c072db / 4

- [routes.response_cookies_to_add](resources--route--reference--group-001.md#canonical-464e949935c94a1ff63a4aafb9943678d4019e7bad5b8f6dd87322599792ba10)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-0096281b2d7af1d7e834d701c6d8ba44459195dd78963e72d47f76a2992fb50c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e0f6ae6f5e940a15c24cfc5ae139e2951fd3b64f02d66578ebc59a58faa8463"></a>

## routes.response_cookies_to_add.ignore_httponly — routes.response_cookies_to_add.ignore_httponly / 9fc126122798 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Property reference](resources--route--reference--group-001.md#canonical-ea4cc975dcd01d0e8a8ede1dd2ddf7fda96d13581cd5a17a04a3cb134c98e22e)
- [routes](resources--route--reference--group-001.md#canonical-db6099d1dbbaf2e52801a091b3c1872ef4a57e4e5a91fbeef8f20e669892379c)
- [routes.response_cookies_to_add](resources--route--reference--group-001.md#canonical-464e949935c94a1ff63a4aafb9943678d4019e7bad5b8f6dd87322599792ba10)
- routes.response_cookies_to_add.ignore_httponly

<a id="canonical-44c08f1083e600b9bf9338ec37ac92853cf23eca729a47269ddbde0258a42874"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_httponly = {}
```
