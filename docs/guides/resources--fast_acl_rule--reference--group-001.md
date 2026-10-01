---
page_title: "xcsh_fast_acl_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl_rule reference."
---

# xcsh_fast_acl_rule reference

<a id="canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dce61ad62e449452d46a612d63385e8ae01bd4691bb874c59a2843e85a98265f"></a>

## Property reference — Property reference / 403360bd7bd7 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- Property reference

<a id="canonical-c966669a40c9b78b80dad5b7acbd56ede2ac4935d2036bef0b3784254b765a62"></a>

## Direct properties — Property reference / 403360bd7bd7 / 3

- [action](resources--fast_acl_rule--reference--group-001.md#canonical-bb7af3453939b520faab92272ab21159966aa35654806b0d1301cc24b1d30bee): complete subsection reference.

<a id="canonical-17a4d80662c6d79d5ca0f3662d5fc2e7e600bdd08ed77d894f6fd89e8e41293e"></a>

<a id="canonical-e0322cbdfb92bbd710b073fe96e1075eaec141c14a5575bc4aa4bf09f0421aea"></a>

## annotations property — Property reference / 403360bd7bd7 / 4

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

<a id="canonical-c823397f29edd06b006445934bcf98fa60b604362505d23815e18b73a5ad9847"></a>

<a id="canonical-bcce14b5db0ab7847e32221e49e21fc59707f4134f4d9b69a20689533ca061d2"></a>

## description property — Property reference / 403360bd7bd7 / 5

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

<a id="canonical-6c6d2bf6bded1d43a9a3131cf8af0fbc36d7e35afdb4a9b692b91cea962c20ff"></a>

<a id="canonical-afb5eed0ec2a8e663e35015296b311b4c90e570e0cc7f5f29af8e579617ecf11"></a>

## disable property — Property reference / 403360bd7bd7 / 6

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

<a id="canonical-abd833551ea07584caed81f0c587b0e62e9f43aa102e066038e447f5ec7bff4d"></a>

<a id="canonical-e5abf7559dfe975aba465966d4d87bb764c1e1969d8d24aa68fb4406e6adc364"></a>

## id property — Property reference / 403360bd7bd7 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_prefix_set](resources--fast_acl_rule--reference--group-001.md#canonical-37c3e6fe0c8d5f37d9ece480319029b58c2437e46cbec31be989485f8cc844ae): complete subsection reference.

<a id="canonical-fb8cc099defb555506fa73f99e06313b07fe911f1a88152a833357c2ed363263"></a>

<a id="canonical-e9bf272710d5074797de06cefd5e9b82cd96298166bb535d807aecf0837a3f55"></a>

## labels property — Property reference / 403360bd7bd7 / 8

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

<a id="canonical-77d62c77199979061137ff84955c50ad660dd0e26b1fe3eb5ab5a555def039f2"></a>

<a id="canonical-16bd9b2b17f84c13a4e91cf3a12732c2728140bd304a0987cc1a9173212947b1"></a>

## name property — Property reference / 403360bd7bd7 / 9

Type: `"string"`. Required.

Name of the Fast ACL Rule. Must be unique within the namespace.

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

<a id="canonical-67a04d67acd61f730e308601994f12b8438b6bfd84fa99aba0374e7723a2e45d"></a>

<a id="canonical-369d09de1a2c68f2efa8813f8dfe2b3fef2011623771584d1b74b84403a9bed2"></a>

## namespace property — Property reference / 403360bd7bd7 / 10

Type: `"string"`. Required.

Namespace where the Fast ACL Rule is created.

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

- [port](resources--fast_acl_rule--reference--group-001.md#canonical-4e0c0cd24f8a48e4ea23298739c1f4f9f6355794e2d684f2af5e5cc3fd1c0f1c): complete subsection reference.

- [prefix](resources--fast_acl_rule--reference--group-001.md#canonical-3778892bcc7977441bbfb8b4b6cdeabb63f35b65e6aeaba3738300f3a05bf037): complete subsection reference.

- [timeouts](resources--fast_acl_rule--reference--group-001.md#canonical-d07b244e2dbd714897c2d94c17f415a8af9f10db961aa85f9b5371d38aaf2321): complete subsection reference.

<a id="canonical-269c38e8c8d562476541e741de4cd11555a9ecbe1ce800f86b688ff445ce6c40"></a>

## All schema paths — Property reference / 403360bd7bd7 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](resources--fast_acl_rule--reference--group-001.md#canonical-7e0d981420d62182573af2d5f775bfa6694227a54aa29e5b3d9376582c76f8a9) |
| `action.policer_action` | [action.policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-b85a795e5fcf7c7007851378fcebab9d638b1eb32c8569daeae185092733db88) |
| `action.policer_action.ref` | [action.policer_action.ref](resources--fast_acl_rule--reference--group-001.md#canonical-35b52a7ca0dafb1f2d4b4cc03e07bb5a79f19edc9468b5f29d4400ddc06a4647) |
| `action.policer_action.ref.kind` | [action.policer_action.ref.kind](resources--fast_acl_rule--reference--group-001.md#canonical-9a8cb36d4d43ffd2dc58118f07844b60b108f00fa11983df1dbe1f9c0611c9d2) |
| `action.policer_action.ref.name` | [action.policer_action.ref.name](resources--fast_acl_rule--reference--group-001.md#canonical-ff7b987588e7da4c57c19c7259815187a6c41e7dadf4f508b41e42ee56711538) |
| `action.policer_action.ref.namespace` | [action.policer_action.ref.namespace](resources--fast_acl_rule--reference--group-001.md#canonical-b4b688442650f89e4da3ff0f474eb17566fe2e6e99d1c5d5fac8c3d84e3fd240) |
| `action.policer_action.ref.tenant` | [action.policer_action.ref.tenant](resources--fast_acl_rule--reference--group-001.md#canonical-c27bf0751d1a76e05495ec5189ca4212caf0726626271406d366c012cf9854a6) |
| `action.policer_action.ref.uid` | [action.policer_action.ref.uid](resources--fast_acl_rule--reference--group-001.md#canonical-b0bfe26ff304035198c4f85ee6fbc9378f039c01bd9ffbf95c2c61fdf89f2f4c) |
| `action.protocol_policer_action` | [action.protocol_policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-e0d7906a94486d3f48b974dc56b8a1c0fb6cde5615f1854454a2006e1d84ccfc) |
| `action.protocol_policer_action.ref` | [action.protocol_policer_action.ref](resources--fast_acl_rule--reference--group-001.md#canonical-b4ba85df620c20ecaa7af65848d4875cb615552765f176db45913c6c0564f179) |
| `action.protocol_policer_action.ref.kind` | [action.protocol_policer_action.ref.kind](resources--fast_acl_rule--reference--group-001.md#canonical-6e59cbbe64853a52552e3f938fd090319058ac7e12bc8e17412a6e450d413218) |
| `action.protocol_policer_action.ref.name` | [action.protocol_policer_action.ref.name](resources--fast_acl_rule--reference--group-001.md#canonical-50a82e3f59cb5833e20fc1ac3ea13aca1ef6cf75a4214d39a46e4a47983d6523) |
| `action.protocol_policer_action.ref.namespace` | [action.protocol_policer_action.ref.namespace](resources--fast_acl_rule--reference--group-001.md#canonical-1bbab98d0daa90f17bafa253c6042878f756448279223e49f82ac3d50a0325bf) |
| `action.protocol_policer_action.ref.tenant` | [action.protocol_policer_action.ref.tenant](resources--fast_acl_rule--reference--group-001.md#canonical-a679d23982e05f6b3e95561ae21bc622a49b2e353cdfd3ff07466e5032ee411d) |
| `action.protocol_policer_action.ref.uid` | [action.protocol_policer_action.ref.uid](resources--fast_acl_rule--reference--group-001.md#canonical-06755ee6c135b7848aba69098325ddc8220b7e594b8886e7dde3333bc25c79ae) |
| `action.simple_action` | [action.simple_action](resources--fast_acl_rule--reference--group-001.md#canonical-71340a5f8d70cd78f52a5951258c20e313554a87c710877c81975756527555a5) |
| `annotations` | [annotations](resources--fast_acl_rule--reference--group-001.md#canonical-17a4d80662c6d79d5ca0f3662d5fc2e7e600bdd08ed77d894f6fd89e8e41293e) |
| `description` | [description](resources--fast_acl_rule--reference--group-001.md#canonical-c823397f29edd06b006445934bcf98fa60b604362505d23815e18b73a5ad9847) |
| `disable` | [disable](resources--fast_acl_rule--reference--group-001.md#canonical-6c6d2bf6bded1d43a9a3131cf8af0fbc36d7e35afdb4a9b692b91cea962c20ff) |
| `id` | [id](resources--fast_acl_rule--reference--group-001.md#canonical-abd833551ea07584caed81f0c587b0e62e9f43aa102e066038e447f5ec7bff4d) |
| `ip_prefix_set` | [ip_prefix_set](resources--fast_acl_rule--reference--group-001.md#canonical-785bbb7999094e552f246522a832216556ac945d4cecdb95c9ad1afdca532918) |
| `ip_prefix_set.ref` | [ip_prefix_set.ref](resources--fast_acl_rule--reference--group-001.md#canonical-d6470e5f60002a0b5afd6329207875b06d1d1f8c28b3857b4048ac100bda6039) |
| `ip_prefix_set.ref.kind` | [ip_prefix_set.ref.kind](resources--fast_acl_rule--reference--group-001.md#canonical-f7b56bef59718f139dc013035c37307a2f4e694b3989b0a62e5890bf02990c12) |
| `ip_prefix_set.ref.name` | [ip_prefix_set.ref.name](resources--fast_acl_rule--reference--group-001.md#canonical-eb8e722049f2ea1592b9a7d8efc6c5f8b57469b0c77f90c15855a06d59631472) |
| `ip_prefix_set.ref.namespace` | [ip_prefix_set.ref.namespace](resources--fast_acl_rule--reference--group-001.md#canonical-cc85ceba8ff9a4c2536056fb91c0371c7cf9962ae81b33ff8325fced1240b841) |
| `ip_prefix_set.ref.tenant` | [ip_prefix_set.ref.tenant](resources--fast_acl_rule--reference--group-001.md#canonical-27c7da1948fcf852b62b2856463c412e6916346d9bdab69c2312454ece15d5e5) |
| `ip_prefix_set.ref.uid` | [ip_prefix_set.ref.uid](resources--fast_acl_rule--reference--group-001.md#canonical-af918fbc4e7675f6a05ec2f0dcfa03e26e9cdbd840b5f36755ecad6249b216f5) |
| `labels` | [labels](resources--fast_acl_rule--reference--group-001.md#canonical-fb8cc099defb555506fa73f99e06313b07fe911f1a88152a833357c2ed363263) |
| `name` | [name](resources--fast_acl_rule--reference--group-001.md#canonical-77d62c77199979061137ff84955c50ad660dd0e26b1fe3eb5ab5a555def039f2) |
| `namespace` | [namespace](resources--fast_acl_rule--reference--group-001.md#canonical-67a04d67acd61f730e308601994f12b8438b6bfd84fa99aba0374e7723a2e45d) |
| `port` | [port](resources--fast_acl_rule--reference--group-001.md#canonical-7739b04953af5de6bd730868a3140f8812b7a59fc9d23440566c72e18a2fa448) |
| `port.all` | [port.all](resources--fast_acl_rule--reference--group-001.md#canonical-887866cd6839ac189b02e47446943a902331f3d58e3dd1a3e8b86835baebe586) |
| `port.dns` | [port.dns](resources--fast_acl_rule--reference--group-001.md#canonical-538947b36222a540137bc592be0404e23e742ecad2034fc846da3f934a635f38) |
| `port.user_defined` | [port.user_defined](resources--fast_acl_rule--reference--group-001.md#canonical-b886c30fa66d1dca06816baccce7c6b9afd7e9e5cdc4d1837c5830ce379bafd9) |
| `prefix` | [prefix](resources--fast_acl_rule--reference--group-001.md#canonical-922d584e4884762a209eb6b2795b739d6e2b1218806bd9d2793c8b6b470086ba) |
| `prefix.prefix` | [prefix.prefix](resources--fast_acl_rule--reference--group-001.md#canonical-f49a557c8358095e81298b7f17a79d8a9380c8ef1d721d1b1fc2efc2e5f10a59) |
| `timeouts` | [timeouts](resources--fast_acl_rule--reference--group-001.md#canonical-8bfe1681632d305c79b7f395adfc7a951bec52c2cb29056751bd7ae6988760b8) |
| `timeouts.create` | [timeouts.create](resources--fast_acl_rule--reference--group-001.md#canonical-ba6386c268f3c0cc2bf13060686ee273ea5817b9800f25ce12d4a40a70913c72) |
| `timeouts.delete` | [timeouts.delete](resources--fast_acl_rule--reference--group-001.md#canonical-a31883ba3913652f7594752feaf420a8a55405050c2eafd0ce9c85285951efe9) |
| `timeouts.read` | [timeouts.read](resources--fast_acl_rule--reference--group-001.md#canonical-132900ef65882f6cd0aaae1b1bc9343171dc6c76bf02a99d6385e7fcabca4137) |
| `timeouts.update` | [timeouts.update](resources--fast_acl_rule--reference--group-001.md#canonical-bf50c93dcac7a9ba79c5b26e62b67f6dff9ad17d71be0a938b29f3007e128f27) |

<a id="canonical-55830bd60b2f3eb8283d78b4600b5e43d376fae3b76e85cddc3b31999148e4a6"></a>

## Next pages — Property reference / 403360bd7bd7 / 12

- [action](resources--fast_acl_rule--reference--group-001.md#canonical-bb7af3453939b520faab92272ab21159966aa35654806b0d1301cc24b1d30bee)
- [ip_prefix_set](resources--fast_acl_rule--reference--group-001.md#canonical-37c3e6fe0c8d5f37d9ece480319029b58c2437e46cbec31be989485f8cc844ae)
- [port](resources--fast_acl_rule--reference--group-001.md#canonical-4e0c0cd24f8a48e4ea23298739c1f4f9f6355794e2d684f2af5e5cc3fd1c0f1c)
- [prefix](resources--fast_acl_rule--reference--group-001.md#canonical-3778892bcc7977441bbfb8b4b6cdeabb63f35b65e6aeaba3738300f3a05bf037)
- [timeouts](resources--fast_acl_rule--reference--group-001.md#canonical-d07b244e2dbd714897c2d94c17f415a8af9f10db961aa85f9b5371d38aaf2321)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)

<a id="canonical-bb7af3453939b520faab92272ab21159966aa35654806b0d1301cc24b1d30bee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bc888691a92983c05a1def51c3d2255673a23e14034f2cf546f7538afe07be8"></a>

## action — action / 4617c31c10c9 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- action

<a id="canonical-7e0d981420d62182573af2d5f775bfa6694227a54aa29e5b3d9376582c76f8a9"></a>

Type: `"object"`. single nested block, Optional.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Upstream description:

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("policer_action",
    "protocol_policer_action"),
  validators.ConflictingObjectAttributes("policer_action",
    "simple_action"),
  validators.ConflictingObjectAttributes("protocol_policer_action",
    "simple_action")}
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
  "x-ves-oneof-field-action": "[\"policer_action\",\"protocol_policer_action\",\"simple_action\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-3e82f1197e5491bc9be64426fbec4b49b28a08473618bb860df97a6ebf0031fc"></a>

## Direct properties — action / 4617c31c10c9 / 3

- [policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-68907b945310feada0f4d64e4425ad8616bcd704bd44f9624ae3ce7929474f43): complete subsection reference.

- [protocol_policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-0e2fe60c49b715acab24d8f4b9b68be0294c26acfd4064e96c40ed7bb41b8aca): complete subsection reference.

<a id="canonical-71340a5f8d70cd78f52a5951258c20e313554a87c710877c81975756527555a5"></a>

<a id="canonical-e8ebbae29a4ced5666c87c3f510044708a148edc9e5fc97253115a03be5840d7"></a>

## simple_action property — action / 4617c31c10c9 / 4

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

FastAclRuleSimpleAction specifies simple action like PASS or DENY

Drop the traffic Forward the traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-cd26cf31a18fc035316abf9828c9657c1dbc7bc24effb0bdfb23e54b0569262c"></a>

## Next pages — action / 4617c31c10c9 / 5

- [action.policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-68907b945310feada0f4d64e4425ad8616bcd704bd44f9624ae3ce7929474f43)
- [action.protocol_policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-0e2fe60c49b715acab24d8f4b9b68be0294c26acfd4064e96c40ed7bb41b8aca)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)

<a id="canonical-68907b945310feada0f4d64e4425ad8616bcd704bd44f9624ae3ce7929474f43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41a1835746fa57899cb1bce5f976549bbd388248f0fa665efd3ecd8fb2cba07e"></a>

## action.policer_action — action.policer_action / aa77cea8b94c / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-bb7af3453939b520faab92272ab21159966aa35654806b0d1301cc24b1d30bee)
- action.policer_action

<a id="canonical-b85a795e5fcf7c7007851378fcebab9d638b1eb32c8569daeae185092733db88"></a>

Type: `"object"`. single nested block, Optional.

Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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
policer_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-ec1c46c58b151ff70c83b29023093f410b1a029afb8ba0ea731652c3e02cf0a6"></a>

## Direct properties — action.policer_action / aa77cea8b94c / 3

- [ref](resources--fast_acl_rule--reference--group-001.md#canonical-b2046280b35a38aac797e798ba917d614adb5044e2d7c06119402aed0c5b47b8): complete subsection reference.

<a id="canonical-cb5df2af156c7741ecd8d309a08b6447b96d73629aa0c9e27060f5d3eaa512f1"></a>

## Next pages — action.policer_action / aa77cea8b94c / 4

- [action.policer_action.ref](resources--fast_acl_rule--reference--group-001.md#canonical-b2046280b35a38aac797e798ba917d614adb5044e2d7c06119402aed0c5b47b8)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-bb7af3453939b520faab92272ab21159966aa35654806b0d1301cc24b1d30bee)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)

<a id="canonical-b2046280b35a38aac797e798ba917d614adb5044e2d7c06119402aed0c5b47b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa0c07fc66e801a100c85244f3dbfa033e1a356a4d843c7cc95c371d1e54646a"></a>

## action.policer_action.ref — action.policer_action.ref / 2a5d8ea6b5aa / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-bb7af3453939b520faab92272ab21159966aa35654806b0d1301cc24b1d30bee)
- [action.policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-68907b945310feada0f4d64e4425ad8616bcd704bd44f9624ae3ce7929474f43)
- action.policer_action.ref

<a id="canonical-35b52a7ca0dafb1f2d4b4cc03e07bb5a79f19edc9468b5f29d4400ddc06a4647"></a>

Type: `"object"`. list nested block, Optional.

Reference. A policer direct reference.

Upstream description:

A policer direct reference.

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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-517aa6910cfc9e1459439d6fcb8647db5db74fc151ca0c7724348d9d3eff7ce2"></a>

## Direct properties — action.policer_action.ref / 2a5d8ea6b5aa / 3

<a id="canonical-9a8cb36d4d43ffd2dc58118f07844b60b108f00fa11983df1dbe1f9c0611c9d2"></a>

<a id="canonical-a8380a8e92ba54ca4757134e23bedcc62a69564f6106a0b4847028af4364eec2"></a>

## kind property — action.policer_action.ref / 2a5d8ea6b5aa / 4

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

<a id="canonical-ff7b987588e7da4c57c19c7259815187a6c41e7dadf4f508b41e42ee56711538"></a>

<a id="canonical-402c361f8a2ac7e5351817a944d711844c9bd70807ff120cd6aebcb2e107eaa8"></a>

## name property — action.policer_action.ref / 2a5d8ea6b5aa / 5

Type: `"string"`. Optional.

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

<a id="canonical-b4b688442650f89e4da3ff0f474eb17566fe2e6e99d1c5d5fac8c3d84e3fd240"></a>

<a id="canonical-d07d79b10b554e168c8881f2f7b1205f71a94dc9bb32ad9c327eb3ef94c486a0"></a>

## namespace property — action.policer_action.ref / 2a5d8ea6b5aa / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-c27bf0751d1a76e05495ec5189ca4212caf0726626271406d366c012cf9854a6"></a>

<a id="canonical-c342a71142fe22b95dad67b7852a33d8b844c9630ddda9776bf816493c590174"></a>

## tenant property — action.policer_action.ref / 2a5d8ea6b5aa / 7

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

<a id="canonical-b0bfe26ff304035198c4f85ee6fbc9378f039c01bd9ffbf95c2c61fdf89f2f4c"></a>

<a id="canonical-82b480ea3ae7f9e8f8b7e1f63ebff4ef637de759a553eac7ae248aa47bcbda8d"></a>

## uid property — action.policer_action.ref / 2a5d8ea6b5aa / 8

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

<a id="canonical-3dc1b48ccc840a021963fc0dde1bd9f49b573e5a6fdda764ba99a8435ed0579c"></a>

## Next pages — action.policer_action.ref / 2a5d8ea6b5aa / 9

- [action.policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-68907b945310feada0f4d64e4425ad8616bcd704bd44f9624ae3ce7929474f43)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)

<a id="canonical-0e2fe60c49b715acab24d8f4b9b68be0294c26acfd4064e96c40ed7bb41b8aca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16676b5a6bc415d1a787001b1b64fde03d01e1a23b7f3f7e2a1b8640e8e1cae4"></a>

## action.protocol_policer_action — action.protocol_policer_action / 15fe12690283 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-bb7af3453939b520faab92272ab21159966aa35654806b0d1301cc24b1d30bee)
- action.protocol_policer_action

<a id="canonical-e0d7906a94486d3f48b974dc56b8a1c0fb6cde5615f1854454a2006e1d84ccfc"></a>

Type: `"object"`. single nested block, Optional.

Protocol Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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
protocol_policer_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-778240f83915510058736ee3664443a5c98ce48022c10cb3346deb4b69b38469"></a>

## Direct properties — action.protocol_policer_action / 15fe12690283 / 3

- [ref](resources--fast_acl_rule--reference--group-001.md#canonical-e4eedf275c50238d2a079f941c86400d9a9fe890424f89d30c0a501882b165c6): complete subsection reference.

<a id="canonical-94baf0fa1304f3ec6119ac33c66d9d3f33a653fd17ec61e78986609a1604f0ac"></a>

## Next pages — action.protocol_policer_action / 15fe12690283 / 4

- [action.protocol_policer_action.ref](resources--fast_acl_rule--reference--group-001.md#canonical-e4eedf275c50238d2a079f941c86400d9a9fe890424f89d30c0a501882b165c6)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-bb7af3453939b520faab92272ab21159966aa35654806b0d1301cc24b1d30bee)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)

<a id="canonical-e4eedf275c50238d2a079f941c86400d9a9fe890424f89d30c0a501882b165c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68860f2260c398426d0878dae169b7e6eddaace81ce04db111c58589282601c6"></a>

## action.protocol_policer_action.ref — action.protocol_policer_action.ref / 97fb898afdfd / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-bb7af3453939b520faab92272ab21159966aa35654806b0d1301cc24b1d30bee)
- [action.protocol_policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-0e2fe60c49b715acab24d8f4b9b68be0294c26acfd4064e96c40ed7bb41b8aca)
- action.protocol_policer_action.ref

<a id="canonical-b4ba85df620c20ecaa7af65848d4875cb615552765f176db45913c6c0564f179"></a>

Type: `"object"`. list nested block, Optional.

Protocol policer Reference. Reference to protocol policer object.

Upstream description:

Reference to protocol policer object.

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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-c9187bc940af533df7b72138da2d5762cca846f9799b742336aac6298fce54b7"></a>

## Direct properties — action.protocol_policer_action.ref / 97fb898afdfd / 3

<a id="canonical-6e59cbbe64853a52552e3f938fd090319058ac7e12bc8e17412a6e450d413218"></a>

<a id="canonical-f4a424dab6c7ad545d56af9210a61e32bf566f6b9d8a965cfad9cb8ab4cb06ea"></a>

## kind property — action.protocol_policer_action.ref / 97fb898afdfd / 4

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

<a id="canonical-50a82e3f59cb5833e20fc1ac3ea13aca1ef6cf75a4214d39a46e4a47983d6523"></a>

<a id="canonical-2c17822a492f194490277083662f0ae777e0524f9658c2c2d517dda7544a6b6b"></a>

## name property — action.protocol_policer_action.ref / 97fb898afdfd / 5

Type: `"string"`. Optional.

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

<a id="canonical-1bbab98d0daa90f17bafa253c6042878f756448279223e49f82ac3d50a0325bf"></a>

<a id="canonical-d9fe711d1bc43ae19a5b1b1f32074427ed861022ffe09aa6aae49c88159bfb7d"></a>

## namespace property — action.protocol_policer_action.ref / 97fb898afdfd / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-a679d23982e05f6b3e95561ae21bc622a49b2e353cdfd3ff07466e5032ee411d"></a>

<a id="canonical-e64bcd891bd3d8ba948641be9e4310c47bfd4a9ba2c83468996d4ed5af4f9190"></a>

## tenant property — action.protocol_policer_action.ref / 97fb898afdfd / 7

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

<a id="canonical-06755ee6c135b7848aba69098325ddc8220b7e594b8886e7dde3333bc25c79ae"></a>

<a id="canonical-31ad8255d066ebf3a3d55249de0cb2ca6d06e63d606f33cf0e3d1cfcbb248294"></a>

## uid property — action.protocol_policer_action.ref / 97fb898afdfd / 8

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

<a id="canonical-3b50f4ecb18d1519b1aa70109cd873a975b20fc5f71d692ad5c97ee2b63519a6"></a>

## Next pages — action.protocol_policer_action.ref / 97fb898afdfd / 9

- [action.protocol_policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-0e2fe60c49b715acab24d8f4b9b68be0294c26acfd4064e96c40ed7bb41b8aca)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)

<a id="canonical-37c3e6fe0c8d5f37d9ece480319029b58c2437e46cbec31be989485f8cc844ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-315ed6265ca63756b107491b067fe074037789f6dd05f85096c4b9ac3761069f"></a>

## ip_prefix_set — ip_prefix_set / 1815bcb3a111 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- ip_prefix_set

<a id="canonical-785bbb7999094e552f246522a832216556ac945d4cecdb95c9ad1afdca532918"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ip\_prefix\_set, prefix\] List of references to ip\_prefix\_set objects.

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

OneOf alternatives in this subsection:

- [ip_prefix_set](resources--fast_acl_rule--reference--group-001.md#canonical-785bbb7999094e552f246522a832216556ac945d4cecdb95c9ad1afdca532918)
- [prefix](resources--fast_acl_rule--reference--group-001.md#canonical-922d584e4884762a209eb6b2795b739d6e2b1218806bd9d2793c8b6b470086ba)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-c5ed043387b0515125d613c63cb2a4505016f7caddd924ccdbf7062b8d63c82e"></a>

## Direct properties — ip_prefix_set / 1815bcb3a111 / 3

- [ref](resources--fast_acl_rule--reference--group-001.md#canonical-09d9962567a6c13c9c10d79bcf18c26d0a5c4976ca8257611757ae5db75a993e): complete subsection reference.

<a id="canonical-60131576db65b24679a6b2e1c4aa0f3e8c4aa7360734dba3e10dbd1059805a80"></a>

## Next pages — ip_prefix_set / 1815bcb3a111 / 4

- [ip_prefix_set.ref](resources--fast_acl_rule--reference--group-001.md#canonical-09d9962567a6c13c9c10d79bcf18c26d0a5c4976ca8257611757ae5db75a993e)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)

<a id="canonical-09d9962567a6c13c9c10d79bcf18c26d0a5c4976ca8257611757ae5db75a993e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1aae85ea7e1b4d2fd8e9bdc2c21115871f47f0fb492895e2db0fa17332093f77"></a>

## ip_prefix_set.ref — ip_prefix_set.ref / 7c0c63c36a8a / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- [ip_prefix_set](resources--fast_acl_rule--reference--group-001.md#canonical-37c3e6fe0c8d5f37d9ece480319029b58c2437e46cbec31be989485f8cc844ae)
- ip_prefix_set.ref

<a id="canonical-d6470e5f60002a0b5afd6329207875b06d1d1f8c28b3857b4048ac100bda6039"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-8900fb6629db9f0ab02cc2b652de973775915c09385a40e8270895505b2d40e7"></a>

## Direct properties — ip_prefix_set.ref / 7c0c63c36a8a / 3

<a id="canonical-f7b56bef59718f139dc013035c37307a2f4e694b3989b0a62e5890bf02990c12"></a>

<a id="canonical-368948260d97e3cc0424f94c7827d9fdbaaaf23b8ca59fc5ba537a2b398e7c0c"></a>

## kind property — ip_prefix_set.ref / 7c0c63c36a8a / 4

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

<a id="canonical-eb8e722049f2ea1592b9a7d8efc6c5f8b57469b0c77f90c15855a06d59631472"></a>

<a id="canonical-80f78753a5a32045bab6635f0c6297387c8d88ca561454065c1105ac92df30c5"></a>

## name property — ip_prefix_set.ref / 7c0c63c36a8a / 5

Type: `"string"`. Optional.

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

<a id="canonical-cc85ceba8ff9a4c2536056fb91c0371c7cf9962ae81b33ff8325fced1240b841"></a>

<a id="canonical-9f50507e4e9ecea86260b17d1843859baf6c34fc2ab47a989342eb2ae2265d7b"></a>

## namespace property — ip_prefix_set.ref / 7c0c63c36a8a / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-27c7da1948fcf852b62b2856463c412e6916346d9bdab69c2312454ece15d5e5"></a>

<a id="canonical-d0cf4ad3cc76d58979421128077124a429342a0a9ce9d2fab7d3f0459ce53579"></a>

## tenant property — ip_prefix_set.ref / 7c0c63c36a8a / 7

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

<a id="canonical-af918fbc4e7675f6a05ec2f0dcfa03e26e9cdbd840b5f36755ecad6249b216f5"></a>

<a id="canonical-e824006389d9ef874f41e00f2dfe9dc58e36763fe3aca37f1f78d0b691439be7"></a>

## uid property — ip_prefix_set.ref / 7c0c63c36a8a / 8

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

<a id="canonical-1a5c6a7257b234648928baa7922d7e15213be412a006cc033511efb1801d17cb"></a>

## Next pages — ip_prefix_set.ref / 7c0c63c36a8a / 9

- [ip_prefix_set](resources--fast_acl_rule--reference--group-001.md#canonical-37c3e6fe0c8d5f37d9ece480319029b58c2437e46cbec31be989485f8cc844ae)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)

<a id="canonical-4e0c0cd24f8a48e4ea23298739c1f4f9f6355794e2d684f2af5e5cc3fd1c0f1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-677a2baa8e8352a44b122280083c443fdf2b765fdf79499c04c4a73ae34a950e"></a>

## port — port / 46031508a40e / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- port

<a id="canonical-7739b04953af5de6bd730868a3140f8812b7a59fc9d23440566c72e18a2fa448"></a>

Type: `"object"`. list nested block, Optional.

Source Ports. L4 port numbers to match.

Upstream description:

L4 port numbers to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all",
    "dns"),
  validators.ConflictingListObjectAttributes("all",
    "user_defined"),
  validators.ConflictingListObjectAttributes("dns",
    "user_defined")}
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-d6005bc5dcde3528997cbf315460eb2e452bc6a64f46945ae24f8181a5c8baa4"></a>

## Direct properties — port / 46031508a40e / 3

- [all](resources--fast_acl_rule--reference--group-001.md#canonical-faac8e8020645eb81584e90f0d2f13dd7db4d5968a6b258dc1c4688dd9a17945): complete subsection reference.

- [dns](resources--fast_acl_rule--reference--group-001.md#canonical-b9d0f9ea646e57ec34863e851768829f356b9b1cddaf8e1f75bf80c4f0a3fcaf): complete subsection reference.

<a id="canonical-b886c30fa66d1dca06816baccce7c6b9afd7e9e5cdc4d1837c5830ce379bafd9"></a>

<a id="canonical-0fe1d9d1bf73059b7ae1d5da8423a5f56bade9ec230f45344685d272abfd83c3"></a>

## user_defined property — port / 46031508a40e / 4

Type: `"number"`. Optional.

Exclusive with \[all DNS\] Matches the user defined port.

Upstream description:

Exclusive with \[all DNS\] Matches the user defined port.

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-ce6784e0da8dd155135d314c401472febe6ba9c38f5d0e762d0f768ea03b1fec"></a>

## Next pages — port / 46031508a40e / 5

- [port.all](resources--fast_acl_rule--reference--group-001.md#canonical-faac8e8020645eb81584e90f0d2f13dd7db4d5968a6b258dc1c4688dd9a17945)
- [port.dns](resources--fast_acl_rule--reference--group-001.md#canonical-b9d0f9ea646e57ec34863e851768829f356b9b1cddaf8e1f75bf80c4f0a3fcaf)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)

<a id="canonical-faac8e8020645eb81584e90f0d2f13dd7db4d5968a6b258dc1c4688dd9a17945"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-270d369c369eacf5d2ac62246309deb91b97a8a933d79ff406d5a3836570248f"></a>

## port.all — port.all / 68fca8836949 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- [port](resources--fast_acl_rule--reference--group-001.md#canonical-4e0c0cd24f8a48e4ea23298739c1f4f9f6355794e2d684f2af5e5cc3fd1c0f1c)
- port.all

<a id="canonical-887866cd6839ac189b02e47446943a902331f3d58e3dd1a3e8b86835baebe586"></a>

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
all = {}
```

<a id="canonical-a3e7981b4646fd72539120f04d94938e1c24057da116455f89abb6db15a91776"></a>

## Direct properties — port.all / 68fca8836949 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bdfec4cf709e3ec2784fe374218505b2a5050171771c9194f6d221c2d38f2430"></a>

## Next pages — port.all / 68fca8836949 / 4

- [port](resources--fast_acl_rule--reference--group-001.md#canonical-4e0c0cd24f8a48e4ea23298739c1f4f9f6355794e2d684f2af5e5cc3fd1c0f1c)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)

<a id="canonical-b9d0f9ea646e57ec34863e851768829f356b9b1cddaf8e1f75bf80c4f0a3fcaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-046719c5c93eb950fc1bc50f26888caf620ea99cc489823092b8c93e287eb9b1"></a>

## port.dns — port.dns / b99569a96771 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- [port](resources--fast_acl_rule--reference--group-001.md#canonical-4e0c0cd24f8a48e4ea23298739c1f4f9f6355794e2d684f2af5e5cc3fd1c0f1c)
- port.dns

<a id="canonical-538947b36222a540137bc592be0404e23e742ecad2034fc846da3f934a635f38"></a>

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
dns = {}
```

<a id="canonical-aac71c40ae82bd476fbe3736f426d6f0668706e3d0e29a4247339d310fc083f9"></a>

## Direct properties — port.dns / b99569a96771 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-208b7e48a49511da0ecc1b194f31ca2a977faa41c70deb493e15530120ad7151"></a>

## Next pages — port.dns / b99569a96771 / 4

- [port](resources--fast_acl_rule--reference--group-001.md#canonical-4e0c0cd24f8a48e4ea23298739c1f4f9f6355794e2d684f2af5e5cc3fd1c0f1c)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)

<a id="canonical-3778892bcc7977441bbfb8b4b6cdeabb63f35b65e6aeaba3738300f3a05bf037"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fbcc79848ee1729d1f2e7981b611a123a0a9a6db63287a8abfc60919ebf8c39"></a>

## prefix — prefix / e4446b3388fd / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- prefix

<a id="canonical-922d584e4884762a209eb6b2795b739d6e2b1218806bd9d2793c8b6b470086ba"></a>

Type: `"object"`. single nested block, Optional.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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
prefix {
  # Configure direct properties listed below.
}
```

<a id="canonical-61adbba7fce1d94ee57179d6b7b4e3bfd8acaebfb89fc4abb70e3fb264730d00"></a>

## Direct properties — prefix / e4446b3388fd / 3

<a id="canonical-f49a557c8358095e81298b7f17a79d8a9380c8ef1d721d1b1fc2efc2e5f10a59"></a>

<a id="canonical-aff68348d17fecbb93cd4bd420c0431fdd60c333e12c6e4410ecced7c19193e3"></a>

## prefix property — prefix / e4446b3388fd / 4

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-72bd8d920d3ff4c942ebd38ce09e5680cca6a04223e38ce637e4904ff56c445a"></a>

## Next pages — prefix / e4446b3388fd / 5

- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)

<a id="canonical-d07b244e2dbd714897c2d94c17f415a8af9f10db961aa85f9b5371d38aaf2321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0026fed8ff86e85a941f7d94727512fd97903a549987326b1ab77e31d916a421"></a>

## timeouts — timeouts / d219cf526c3c / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- timeouts

<a id="canonical-8bfe1681632d305c79b7f395adfc7a951bec52c2cb29056751bd7ae6988760b8"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-e3cb464a045b0dd0080881856733e127b89dc8d0b86cc1be4970acf9ff0c503b"></a>

## Direct properties — timeouts / d219cf526c3c / 3

<a id="canonical-ba6386c268f3c0cc2bf13060686ee273ea5817b9800f25ce12d4a40a70913c72"></a>

<a id="canonical-12933e05b5798da19ea14b04768e9395b7c5070488f758ce87a81790cfb2fe17"></a>

## create property — timeouts / d219cf526c3c / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a31883ba3913652f7594752feaf420a8a55405050c2eafd0ce9c85285951efe9"></a>

<a id="canonical-ed1fe4f559a29c0e78ab917ae492f5fa834ae5a3f5ffabd085c0b2cbc8a3ca3f"></a>

## delete property — timeouts / d219cf526c3c / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-132900ef65882f6cd0aaae1b1bc9343171dc6c76bf02a99d6385e7fcabca4137"></a>

<a id="canonical-4cd189713211a4f9ca8b25bbd79f733af28ab5de8b13eb17f0eec2ff27475477"></a>

## read property — timeouts / d219cf526c3c / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-bf50c93dcac7a9ba79c5b26e62b67f6dff9ad17d71be0a938b29f3007e128f27"></a>

<a id="canonical-017df5f3f3c2131f5412182f476857a2961d6695afaacd4237c23f4ea08aee50"></a>

## update property — timeouts / d219cf526c3c / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-240231352efeac2c0b7b767be04a0615ee099ce683981517517552302f77e5b6"></a>

## Next pages — timeouts / d219cf526c3c / 8

- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
