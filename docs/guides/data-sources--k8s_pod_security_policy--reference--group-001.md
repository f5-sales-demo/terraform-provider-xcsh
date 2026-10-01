---
page_title: "xcsh_k8s_pod_security_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_policy reference."
---

# xcsh_k8s_pod_security_policy reference

<a id="canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14dab4687b409eaa44ed6d086959b24801aca85c808aaa09bfb8bb51b2ba9628"></a>

## Property reference — Property reference / 210c9c99a4ee / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- Property reference

<a id="canonical-010b6fa4fea60042d17337439106dd02d84f2ddadd6ff8a32973a2f3a6d6f1e3"></a>

## Direct properties — Property reference / 210c9c99a4ee / 3

<a id="canonical-d55db4c55d3cf5f4920dac7cde923be72a920d6e5909be7a5dcc042d79f2f24b"></a>

<a id="canonical-ca6d6cd8a9f8fe594bd95a04fb4cf18bf198da0890b09070a7fa117a57405644"></a>

## annotations property — Property reference / 210c9c99a4ee / 4

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

<a id="canonical-90accaf4521d96fd6f2febad995a504f643b1f80b4dfb892b9abb6ae571911ae"></a>

<a id="canonical-789b5251110f4a8deefb22057792df8c5ae6b360f4601c03db7012d9646b8466"></a>

## description property — Property reference / 210c9c99a4ee / 5

Type: `"string"`. Computed.

Description of the K8SPodSecurityPolicy.

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

<a id="canonical-bd01b4111327208d668c74d91a89d65177497de0c406570597e4c226795c98c6"></a>

<a id="canonical-c3695989e4b607223a38f3ed33a7254a49e85623ddb6a2c8a8b9eaf6cd5343e0"></a>

## id property — Property reference / 210c9c99a4ee / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-c30d2b0ab68b6077e439e961e9ba8fc7e3d362d24e407b4ff94c1ada26abcae3"></a>

<a id="canonical-9209325d95876a9d9eff77f5232bf779db3cbcbef1c5116fedb2c107338aca23"></a>

## labels property — Property reference / 210c9c99a4ee / 7

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

<a id="canonical-7181eaba425600d21d8b2430ede3059ef8b821600657e64ec0d411e4348f5bd2"></a>

<a id="canonical-8f4903966f407a4567ed7e2ad6b32e0f40618e8a27fb2a3f11b0d1402588c7cd"></a>

## name property — Property reference / 210c9c99a4ee / 8

Type: `"string"`. Required.

Name of the K8SPodSecurityPolicy.

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

<a id="canonical-643efaab6dd01d082e59f8056b150ba040b9ff1880f04b4919dc48ad8f738bb1"></a>

<a id="canonical-b31c01cc189c3f170bf6d32b83902217cb0a79c79bff90510799fd98a8ed83cc"></a>

## namespace property — Property reference / 210c9c99a4ee / 9

Type: `"string"`. Required.

Namespace where the K8SPodSecurityPolicy exists.

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

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f): complete subsection reference.

<a id="canonical-5c82b84a89d1f4560b95d94559eae68d43e01fb8fe6db153de99bf13d8278ebf"></a>

<a id="canonical-0b5fb75dd824247703483726e9164043051b2841e7d3750f28d368214d36949c"></a>

## yaml property — Property reference / 210c9c99a4ee / 10

Type: `"string"`. Computed.

Exclusive with \[psp\_spec\] K8s YAML for Pod Security Policy.

Upstream description:

Exclusive with \[psp\_spec\] K8s YAML for Pod Security Policy.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "Valid parseable YAML",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "validation": {
      "customRule": "Must be valid YAML"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-931bbbea130e180556fc7b3f684767efbaacdf0adbc89f81694571eb37f0a74f"></a>

## All schema paths — Property reference / 210c9c99a4ee / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-d55db4c55d3cf5f4920dac7cde923be72a920d6e5909be7a5dcc042d79f2f24b) |
| `description` | [description](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-90accaf4521d96fd6f2febad995a504f643b1f80b4dfb892b9abb6ae571911ae) |
| `id` | [id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-bd01b4111327208d668c74d91a89d65177497de0c406570597e4c226795c98c6) |
| `labels` | [labels](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-c30d2b0ab68b6077e439e961e9ba8fc7e3d362d24e407b4ff94c1ada26abcae3) |
| `name` | [name](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-7181eaba425600d21d8b2430ede3059ef8b821600657e64ec0d411e4348f5bd2) |
| `namespace` | [namespace](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-643efaab6dd01d082e59f8056b150ba040b9ff1880f04b4919dc48ad8f738bb1) |
| `psp_spec` | [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-34e69abc3b1e3ee6a9bb92e4cfdaaa4c41b1c0ab536fa3ba290078581fceaad4) |
| `psp_spec.allow_privilege_escalation` | [psp_spec.allow_privilege_escalation](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-8606c0868d47d96d588cac5ee7d51dbf141a06f09c8ffd73668f001abb87b90e) |
| `psp_spec.allowed_capabilities` | [psp_spec.allowed_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-fdf509710cf183ad5b4d1e8cce1f9ece9a8882f2b8ba3bb8a14cdd34e7479e87) |
| `psp_spec.allowed_capabilities.capabilities` | [psp_spec.allowed_capabilities.capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-d6ac868490eb7a718224a1c4829bb9e4aced57e52bf1c1133b05be7af641978c) |
| `psp_spec.allowed_csi_drivers` | [psp_spec.allowed_csi_drivers](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-39bf838985956fc5c1f6892e86ef34803f4b710ce42621b07f45d05b76d25bb8) |
| `psp_spec.allowed_flex_volumes` | [psp_spec.allowed_flex_volumes](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-9c00ab4c583817084aad9e247a7af930c728ff41d557f6723339073794ba026f) |
| `psp_spec.allowed_host_paths` | [psp_spec.allowed_host_paths](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-9f29e90036fd5c158b733ff73acdd15c26952ed2e183aeca2e04a3751e29f8f8) |
| `psp_spec.allowed_host_paths.path_prefix` | [psp_spec.allowed_host_paths.path_prefix](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3f34ccade0f2b1a4fcef3778377b8c0ce54e4b027530415def4bb7794b00972f) |
| `psp_spec.allowed_host_paths.read_only` | [psp_spec.allowed_host_paths.read_only](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-32857a5d60df2ab22359bbeda02a4d3af3596b0e2e01ace05fb3726d509e53e3) |
| `psp_spec.allowed_proc_mounts` | [psp_spec.allowed_proc_mounts](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-73f1e537f11ad6aba46922a188ec206aaabfa624d7301af48efc7f2df1e7d2c9) |
| `psp_spec.allowed_unsafe_sysctls` | [psp_spec.allowed_unsafe_sysctls](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3883747a9a2f79abc582610ef71c0e60398547ffd59ae9ecca960be3f14baf12) |
| `psp_spec.default_allow_privilege_escalation` | [psp_spec.default_allow_privilege_escalation](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-4721caec112309d634b8ad7cccd84673ceb70d1b7ca6fa6949c5b097292998ed) |
| `psp_spec.default_capabilities` | [psp_spec.default_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-68080bc2b2b8948c9ec1b8de07f34dd5130283c25120e3f78dd4d95b1eec630d) |
| `psp_spec.default_capabilities.capabilities` | [psp_spec.default_capabilities.capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-75e8b0eaced6512172a4fb90c930f0cc0a1322d513012136358b074931afcf67) |
| `psp_spec.drop_capabilities` | [psp_spec.drop_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-c9b31295914fdbc58bcdbf9d163366f8fc7d3e6e03cded4accba3500d5c3a880) |
| `psp_spec.drop_capabilities.capabilities` | [psp_spec.drop_capabilities.capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-959bc37bdf8b6464171d9b14c8247262f048aa91befaa99b96a201072571cf8d) |
| `psp_spec.forbidden_sysctls` | [psp_spec.forbidden_sysctls](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3b9dc243716e1bcd70daefe4b9eacd5d988b16cec67efddb50f11bed7ea7cee4) |
| `psp_spec.fs_group_strategy_options` | [psp_spec.fs_group_strategy_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1be7147fd868436517aa5d1bc29fb6872451cda426449d07a555e7ead991a4b7) |
| `psp_spec.fs_group_strategy_options.id_ranges` | [psp_spec.fs_group_strategy_options.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-7b615fb83c5df4108d8950e170cdcd6c907637546b4d6603718dc5203f93cc38) |
| `psp_spec.fs_group_strategy_options.id_ranges.max_id` | [psp_spec.fs_group_strategy_options.id_ranges.max_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-65ab42b898e880923e589ffd692e4590b5f6e5d8f607ba1c1918568adadd0c65) |
| `psp_spec.fs_group_strategy_options.id_ranges.min_id` | [psp_spec.fs_group_strategy_options.id_ranges.min_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-a366dc03f1f03b850b81433bec25c3b5c01be7f973a35a3f10ebddd31f51d2f9) |
| `psp_spec.fs_group_strategy_options.rule` | [psp_spec.fs_group_strategy_options.rule](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-dca9669a74ad6fed3caa1c7e530f2ce63c72cb3447b8b4252511fa5eb5e9e72a) |
| `psp_spec.host_ipc` | [psp_spec.host_ipc](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-91511a75dfda87cd271bdef7ca95fb71cffe81266239c29ba308b1099e7f8e2f) |
| `psp_spec.host_network` | [psp_spec.host_network](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1b04d3c8fe5951daf397f2c7431c0b8934f4500c1a1c2f556a09b62345a1af3f) |
| `psp_spec.host_pid` | [psp_spec.host_pid](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0fc5d9861b8e563cd2e5d052bc5b04cb8789f26d71b7d73178341b943465fdc4) |
| `psp_spec.host_port_ranges` | [psp_spec.host_port_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-76908419b99d146442e6998d70eb3250acdad91cc6cf2295803f115d1834fa30) |
| `psp_spec.no_allowed_capabilities` | [psp_spec.no_allowed_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-12f4e74ae1fbac18015dfacff7ffb6b1caf50d094671d3d1133be5118ba9d78f) |
| `psp_spec.no_default_capabilities` | [psp_spec.no_default_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-bc44c296b512f3ac983d21be4d728374523e513f325bd4f20ed1c503f811d77d) |
| `psp_spec.no_drop_capabilities` | [psp_spec.no_drop_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-5304c9dc8eb2dc467cd9c7a7ca8fc11695e85b02a9edc31540bfa1515b97bacb) |
| `psp_spec.no_fs_groups` | [psp_spec.no_fs_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-813f822b1750d53f0c1df32bce0108d909f4fabaee42aff3c59528fbab0ebd24) |
| `psp_spec.no_run_as_group` | [psp_spec.no_run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-801004e949d1b8939a65a25b0a95a808b8093ba1358cdcc25c077febce95d2a8) |
| `psp_spec.no_run_as_user` | [psp_spec.no_run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-6d8d38006b2808c79050c85eb84fecc577969b72d6c96ccc51db231184652dcb) |
| `psp_spec.no_runtime_class` | [psp_spec.no_runtime_class](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-e2467cdb00e42f7dc501dfd60ef7c1e7253e6661062295e3358a1a37b82bf0f2) |
| `psp_spec.no_se_linux_options` | [psp_spec.no_se_linux_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3c79c9e1d961072a5ede788abcbb72297509650d596f681a4b0eb5ebbb1e412e) |
| `psp_spec.no_supplemental_groups` | [psp_spec.no_supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-930b0021bdb74f5030b2279262b1b6b754290b5b57090e3f6148c0f14899feb3) |
| `psp_spec.privileged` | [psp_spec.privileged](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3b588cb982e5aaf80b4d0801d1baccbb2cc9ec936c4aa2e8623b006aae1a9d33) |
| `psp_spec.read_only_root_filesystem` | [psp_spec.read_only_root_filesystem](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-11e7d48178613cb9bc741a982bd4fd72d6e09357c013be641577606ee63c516c) |
| `psp_spec.run_as_group` | [psp_spec.run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-4309f036ea5f27ae738777d8c6bf0bfa846f265d786e78c7ae0eff280c8c0ec5) |
| `psp_spec.run_as_group.id_ranges` | [psp_spec.run_as_group.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-02bc732c95b06d84c7775ccb5b9c104421e9bd0c03274d923490a00ba092e40a) |
| `psp_spec.run_as_group.id_ranges.max_id` | [psp_spec.run_as_group.id_ranges.max_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-c7db105d4063f032b1748af25d462230777f484a6cb232f328afd24efcd158ea) |
| `psp_spec.run_as_group.id_ranges.min_id` | [psp_spec.run_as_group.id_ranges.min_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-662ab704c1f3a5cbb8f874e369e42ee6dbe880ed7bcca20d18b150d87c579032) |
| `psp_spec.run_as_group.rule` | [psp_spec.run_as_group.rule](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-4c7df6e8fd9fc7744c8b3aa80022856d9bb2e70c8b12ee8e52731f55d73c5770) |
| `psp_spec.run_as_user` | [psp_spec.run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-196dfed234d9aaa9539524eda28486581757688b54f4998ba21cab63a6491e55) |
| `psp_spec.run_as_user.id_ranges` | [psp_spec.run_as_user.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-a7458e60ad4dcc305058b000e00872c0a588a61f8490ae80690a299374a5c9a6) |
| `psp_spec.run_as_user.id_ranges.max_id` | [psp_spec.run_as_user.id_ranges.max_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-a37f954cd4edb77962d13ba06e6403a35c154800a7cbac9341bea513c3856d95) |
| `psp_spec.run_as_user.id_ranges.min_id` | [psp_spec.run_as_user.id_ranges.min_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3dd75acdf63d930de0ddfd35ede764e741b785f4f4d86f28a76032ee8796617d) |
| `psp_spec.run_as_user.rule` | [psp_spec.run_as_user.rule](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-d4cb969e5706418f180d93dc0fd6b9ba59c18f30b1c1aad2b67e64f40433eb4d) |
| `psp_spec.supplemental_groups` | [psp_spec.supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-c3d9a39fe272415824b4e0ebddaf5fde7b649fdab6995d1683226fca664239dd) |
| `psp_spec.supplemental_groups.id_ranges` | [psp_spec.supplemental_groups.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-811133fd8b28c7fb579475e6a4e205256a1fb28c7433f822ad868677c194c6d5) |
| `psp_spec.supplemental_groups.id_ranges.max_id` | [psp_spec.supplemental_groups.id_ranges.max_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-04c36df85b435e88269175eff558baf2bc8916b195f9c5150e62893dadcb8e34) |
| `psp_spec.supplemental_groups.id_ranges.min_id` | [psp_spec.supplemental_groups.id_ranges.min_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-06bbbd3220fbef753adf35b9c1def7cfad32b9bc7174e2ee7cf0253230bd0ffa) |
| `psp_spec.supplemental_groups.rule` | [psp_spec.supplemental_groups.rule](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-5c3b3ee79fe17c5eb8b271997e43442f93751dc4a60f9148486f804439bc47f4) |
| `psp_spec.volumes` | [psp_spec.volumes](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-23b669488d421432c135e7e108644416b912a8f84183cb224f28b53f1aa97ef8) |
| `yaml` | [yaml](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-5c82b84a89d1f4560b95d94559eae68d43e01fb8fe6db153de99bf13d8278ebf) |

<a id="canonical-e379a3f6ddef24bab974055bd20831f1c213adf2c4678a3f1f9506b50ad4aec2"></a>

## Next pages — Property reference / 210c9c99a4ee / 12

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87c97c8750b7145b384458a1c41c6b3aac852249b42af8fa9776bda1362ffc86"></a>

## psp_spec — psp_spec / c8885caf244c / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- psp_spec

<a id="canonical-34e69abc3b1e3ee6a9bb92e4cfdaaa4c41b1c0ab536fa3ba290078581fceaad4"></a>

Type: `"single"`. Computed.

\[OneOf: psp\_spec, yaml\] Pod Security Policy Specification. Form based pod security specification.

Upstream description:

Form based pod security specification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allowed_capabilities_choice": "[\"allowed_capabilities\",\"no_allowed_capabilities\"]",
  "x-ves-oneof-field-default_capabilities_choice": "[\"default_capabilities\",\"no_default_capabilities\"]",
  "x-ves-oneof-field-drop_capabilities_choice": "[\"drop_capabilities\",\"no_drop_capabilities\"]",
  "x-ves-oneof-field-fs_group_choice": "[\"fs_group_strategy_options\",\"no_fs_groups\"]",
  "x-ves-oneof-field-group_choice": "[\"no_run_as_group\",\"run_as_group\"]",
  "x-ves-oneof-field-runtime_class_choice": "[\"no_runtime_class\"]",
  "x-ves-oneof-field-se_linux_choice": "[\"no_se_linux_options\"]",
  "x-ves-oneof-field-supplemental_group_choice": "[\"no_supplemental_groups\",\"supplemental_groups\"]",
  "x-ves-oneof-field-user_choice": "[\"no_run_as_user\",\"run_as_user\"]"
}
```

OneOf alternatives in this subsection:

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-34e69abc3b1e3ee6a9bb92e4cfdaaa4c41b1c0ab536fa3ba290078581fceaad4)
- [yaml](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-5c82b84a89d1f4560b95d94559eae68d43e01fb8fe6db153de99bf13d8278ebf)

Select alternatives according to the provider validators above.

<a id="canonical-6343d4701ec49e81d05aa0d78f2ba9cfe1fc6815394f5132c11bb2d84a18be0b"></a>

## Direct properties — psp_spec / c8885caf244c / 3

<a id="canonical-8606c0868d47d96d588cac5ee7d51dbf141a06f09c8ffd73668f001abb87b90e"></a>

<a id="canonical-35a83c2d5c0363776cc715bfdd1114cc236a30b95c04436203fb86c1527d10b5"></a>

## allow_privilege_escalation property — psp_spec / c8885caf244c / 4

Type: `"bool"`. Computed.

Pod can request to privilege escalation.

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

- [allowed_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-98c5ee8fe87e79cdcab3dcee8c2a511f6b1e292b7a1a74895b649110a0b4de32): complete subsection reference.

<a id="canonical-39bf838985956fc5c1f6892e86ef34803f4b710ce42621b07f45d05b76d25bb8"></a>

<a id="canonical-1d40577bba9040c9b88aef1c1ad82abb99de00d81252663f9a9babbe13894ae6"></a>

## allowed_csi_drivers property — psp_spec / c8885caf244c / 5

Type: `["list", "string"]`. Computed.

Restrict the available CSI drivers for POD, default all drivers are available.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9c00ab4c583817084aad9e247a7af930c728ff41d557f6723339073794ba026f"></a>

<a id="canonical-cdbe81e24ff31be7b2f68c94f4b02a25537f9f109e8614365622a1bb463cbf4c"></a>

## allowed_flex_volumes property — psp_spec / c8885caf244c / 6

Type: `["list", "string"]`. Computed.

Restrict list of Flex volumes, default all volumes are allowed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [allowed_host_paths](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-113fd41da2a9ecfc40331a6e61190b7abc1df37ac5511617b81e473be491376d): complete subsection reference.

<a id="canonical-73f1e537f11ad6aba46922a188ec206aaabfa624d7301af48efc7f2df1e7d2c9"></a>

<a id="canonical-c4a3d158e1df17e947488d5201a8d220893dc8d100dd14cc81052cf6c732f932"></a>

## allowed_proc_mounts property — psp_spec / c8885caf244c / 7

Type: `["list", "string"]`. Computed.

Allowed list of proc mounts, empty list allows default proc mounts.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3883747a9a2f79abc582610ef71c0e60398547ffd59ae9ecca960be3f14baf12"></a>

<a id="canonical-07d8f5c7a8066b3697884a9d387f6307ff86b3e75be9abbe5eb60b62ee39148f"></a>

## allowed_unsafe_sysctls property — psp_spec / c8885caf244c / 8

Type: `["list", "string"]`. Computed.

Allowed list of unsafe sysctls, empty list allows none. Supports prefix reg-ex.

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

<a id="canonical-4721caec112309d634b8ad7cccd84673ceb70d1b7ca6fa6949c5b097292998ed"></a>

<a id="canonical-94ce50933e3202acc8dfb0810dddd44d240767f562ac81c34ecdb06cc4e5c541"></a>

## default_allow_privilege_escalation property — psp_spec / c8885caf244c / 9

Type: `"bool"`. Computed.

Pod has permission for privilege escalation by default.

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

- [default_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-d2518c9bf23bb7b46ceea909901a560b9b33d09acb7fc02aaf53ecda215bedf2): complete subsection reference.

- [drop_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-30152ad1cc9e813ce8b113a10ef4ae80c8862d6f591e4315b3211ce13163016c): complete subsection reference.

<a id="canonical-3b9dc243716e1bcd70daefe4b9eacd5d988b16cec67efddb50f11bed7ea7cee4"></a>

<a id="canonical-559371ecd027f560dcb983a98096ac410a9e72c246b0e9c2c744568cf5ab3957"></a>

## forbidden_sysctls property — psp_spec / c8885caf244c / 10

Type: `["list", "string"]`. Computed.

Forbidden list of sysctls, empty list forbids none. Supports prefix reg-ex.

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

- [fs_group_strategy_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-51fd42d4b5bc8fd40ed4fd1ad317b7774a6ffc912041eba74f4523520ace9e2d): complete subsection reference.

<a id="canonical-91511a75dfda87cd271bdef7ca95fb71cffe81266239c29ba308b1099e7f8e2f"></a>

<a id="canonical-da1afc52506c973dd7b0e634a12e05e74ce3fc35da117bcf67cb2e160af77617"></a>

## host_ipc property — psp_spec / c8885caf244c / 11

Type: `"bool"`. Computed.

Host IPC determines if the policy allows the use of host IPC in the pod spec.

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

<a id="canonical-1b04d3c8fe5951daf397f2c7431c0b8934f4500c1a1c2f556a09b62345a1af3f"></a>

<a id="canonical-bda5c25d8d3341776ee84a5d8e7f6e8c6e31bd44942a44c17e905124a2e72c3c"></a>

## host_network property — psp_spec / c8885caf244c / 12

Type: `"bool"`. Computed.

Host Network determines if the policy allows the use of host network in the pod spec.

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

<a id="canonical-0fc5d9861b8e563cd2e5d052bc5b04cb8789f26d71b7d73178341b943465fdc4"></a>

<a id="canonical-d0950d31e8892da5de04258af6926217bdeb1aaa29a7f4d3683e1d2ea557939f"></a>

## host_pid property — psp_spec / c8885caf244c / 13

Type: `"bool"`. Computed.

Host PID determines if the policy allows the use of host PID in the pod spec.

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

<a id="canonical-76908419b99d146442e6998d70eb3250acdad91cc6cf2295803f115d1834fa30"></a>

<a id="canonical-bb5f352cb8e2941c290d1bdbd64cea6c343f87caadf76856fa32c3994b1fa8bf"></a>

## host_port_ranges property — psp_spec / c8885caf244c / 14

Type: `"string"`. Computed.

Host port ranges determines which ports ranges are allowed to be exposed.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

- [no_allowed_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-841db374e59e81363400bbfed0efb1293e47e6aae84be8de6b6433d4f5e2ac93): complete subsection reference.

- [no_default_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-c7d4de5ef51bd5a3992551b739673fdebb4f7f6f69797644c1695bc0bbd28420): complete subsection reference.

- [no_drop_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0a4908ff2f6784d6eb84b1f29b4b0236bb48c2235083ffa3590d9bc4286118f3): complete subsection reference.

- [no_fs_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-9866f9d9a79532a089b03e0a3cc49eeb933535fb8cc0c3d8969266568a99c6fb): complete subsection reference.

- [no_run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-41c9a33a6fcdfb6ab57ec21a67959b5ab9c8734c0ad03b6e77aee0d5a84cf4db): complete subsection reference.

- [no_run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2dd0f86c2c62b09639dbd851b75468be17ddb1f8af167d09088fc7ced091c635): complete subsection reference.

- [no_runtime_class](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-b80da28e605174bf8c9ae4bb13cdcbe172e3794d29ed6e7c3ccdf71252286c11): complete subsection reference.

- [no_se_linux_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-cf67529b6fb97da5404fc2dfe034cca85957a9ee33695af79469ff981004b35b): complete subsection reference.

- [no_supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-4059018accd2b392bd013e0e0d592e5b9092256c373f645383ddd993b9a1f07c): complete subsection reference.

<a id="canonical-3b588cb982e5aaf80b4d0801d1baccbb2cc9ec936c4aa2e8623b006aae1a9d33"></a>

<a id="canonical-c8f090391e0fafcb6785113cd774480a005aa3abb24570b6558eac392d9c9fff"></a>

## privileged property — psp_spec / c8885caf244c / 15

Type: `"bool"`. Computed.

Privileged determines if a pod can request to be run as privileged.

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

<a id="canonical-11e7d48178613cb9bc741a982bd4fd72d6e09357c013be641577606ee63c516c"></a>

<a id="canonical-3e0246a0044bdb1a6bbf21fdd5b557f83fb184a695913be5057149b1b77d321e"></a>

## read_only_root_filesystem property — psp_spec / c8885caf244c / 16

Type: `"bool"`. Computed.

Containers can only run with read only root filesystem.

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

- [run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-40f047a707869ebaf47acc5a2d4e28f9e3002f0e0b3101d9e399a453bd12acbf): complete subsection reference.

- [run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-a5c7a070654ac9a6ea582fcd39303366b6aaf3a888f64347ad6f2e5d933e1235): complete subsection reference.

- [supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-36b51f79a7eb887beedd4fba4c6a4c3d3ac83da8b5ce3b054b49099e6e08c558): complete subsection reference.

<a id="canonical-23b669488d421432c135e7e108644416b912a8f84183cb224f28b53f1aa97ef8"></a>

<a id="canonical-2a1c43052f8f2d5507b3ef6f2d23274acf998d81afb720313495899c6bc3f8d0"></a>

## volumes property — psp_spec / c8885caf244c / 17

Type: `["list", "string"]`. Computed.

Allow List of volume plugins. Empty no volumes are allowed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-788b5458e7be0676d170c7f9e49631a1479b2c9fd6ba469090db907dbbd2f583"></a>

## Next pages — psp_spec / c8885caf244c / 18

- [psp_spec.allowed_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-98c5ee8fe87e79cdcab3dcee8c2a511f6b1e292b7a1a74895b649110a0b4de32)
- [psp_spec.allowed_host_paths](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-113fd41da2a9ecfc40331a6e61190b7abc1df37ac5511617b81e473be491376d)
- [psp_spec.default_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-d2518c9bf23bb7b46ceea909901a560b9b33d09acb7fc02aaf53ecda215bedf2)
- [psp_spec.drop_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-30152ad1cc9e813ce8b113a10ef4ae80c8862d6f591e4315b3211ce13163016c)
- [psp_spec.fs_group_strategy_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-51fd42d4b5bc8fd40ed4fd1ad317b7774a6ffc912041eba74f4523520ace9e2d)
- [psp_spec.no_allowed_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-841db374e59e81363400bbfed0efb1293e47e6aae84be8de6b6433d4f5e2ac93)
- [psp_spec.no_default_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-c7d4de5ef51bd5a3992551b739673fdebb4f7f6f69797644c1695bc0bbd28420)
- [psp_spec.no_drop_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0a4908ff2f6784d6eb84b1f29b4b0236bb48c2235083ffa3590d9bc4286118f3)
- [psp_spec.no_fs_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-9866f9d9a79532a089b03e0a3cc49eeb933535fb8cc0c3d8969266568a99c6fb)
- [psp_spec.no_run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-41c9a33a6fcdfb6ab57ec21a67959b5ab9c8734c0ad03b6e77aee0d5a84cf4db)
- [psp_spec.no_run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2dd0f86c2c62b09639dbd851b75468be17ddb1f8af167d09088fc7ced091c635)
- [psp_spec.no_runtime_class](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-b80da28e605174bf8c9ae4bb13cdcbe172e3794d29ed6e7c3ccdf71252286c11)
- [psp_spec.no_se_linux_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-cf67529b6fb97da5404fc2dfe034cca85957a9ee33695af79469ff981004b35b)
- [psp_spec.no_supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-4059018accd2b392bd013e0e0d592e5b9092256c373f645383ddd993b9a1f07c)
- [psp_spec.run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-40f047a707869ebaf47acc5a2d4e28f9e3002f0e0b3101d9e399a453bd12acbf)
- [psp_spec.run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-a5c7a070654ac9a6ea582fcd39303366b6aaf3a888f64347ad6f2e5d933e1235)
- [psp_spec.supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-36b51f79a7eb887beedd4fba4c6a4c3d3ac83da8b5ce3b054b49099e6e08c558)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-98c5ee8fe87e79cdcab3dcee8c2a511f6b1e292b7a1a74895b649110a0b4de32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-189439a50124e1a14adfcf47b5fc3e9543da33c98fb895ba4d7de7208f0dd4bd"></a>

## psp_spec.allowed_capabilities — psp_spec.allowed_capabilities / 8fb33541cd81 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.allowed_capabilities

<a id="canonical-fdf509710cf183ad5b4d1e8cce1f9ece9a8882f2b8ba3bb8a14cdd34e7479e87"></a>

Type: `"single"`. Computed.

List of capabilities that docker container has.

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

<a id="canonical-646359c6381e5879d29fbaf9338252f74db5677c77fd8f366f11091539b8b780"></a>

## Direct properties — psp_spec.allowed_capabilities / 8fb33541cd81 / 3

<a id="canonical-d6ac868490eb7a718224a1c4829bb9e4aced57e52bf1c1133b05be7af641978c"></a>

<a id="canonical-b141afd89e3a1284d6b57cbd717c7c23c671642952d945d947a9e1e04ecc4371"></a>

## capabilities property — psp_spec.allowed_capabilities / 8fb33541cd81 / 4

Type: `["list", "string"]`. Computed.

List of capabilities that docker container has.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-fe589f5b325682b36e3aea9bb55075b40f87094bdc3680fba1e335bba343dd01"></a>

## Next pages — psp_spec.allowed_capabilities / 8fb33541cd81 / 5

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-113fd41da2a9ecfc40331a6e61190b7abc1df37ac5511617b81e473be491376d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8517bb05ee1c193b63531d52a9f0530bfc75aad74b71d715eb5d7836df4da2c"></a>

## psp_spec.allowed_host_paths — psp_spec.allowed_host_paths / 104c284af9db / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.allowed_host_paths

<a id="canonical-9f29e90036fd5c158b733ff73acdd15c26952ed2e183aeca2e04a3751e29f8f8"></a>

Type: `"list"`. Computed.

Restrict list of host paths, default all host paths are allowed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7e816574b15125789949e765f8756f4c253bf5fa373918a84afdaad277c748b8"></a>

## Direct properties — psp_spec.allowed_host_paths / 104c284af9db / 3

<a id="canonical-3f34ccade0f2b1a4fcef3778377b8c0ce54e4b027530415def4bb7794b00972f"></a>

<a id="canonical-ec078b43633bf246bffa2fe187df0fa70199e1b9ac34654746a4de15ad5d046e"></a>

## path_prefix property — psp_spec.allowed_host_paths / 104c284af9db / 4

Type: `"string"`. Computed.

Host path prefix is the path prefix that the host volume must match. It does not support \*.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-32857a5d60df2ab22359bbeda02a4d3af3596b0e2e01ace05fb3726d509e53e3"></a>

<a id="canonical-4c199ebc7c8485e1943345e03ecc512651c269f0b208c390ccc7968a2aa94f3a"></a>

## read_only property — psp_spec.allowed_host_paths / 104c284af9db / 5

Type: `"bool"`. Computed.

Volume will be allowed to mount read only.

Upstream description:

This volume will be allowed to mount read only.

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

<a id="canonical-a4bcba9dae9c365782077722a3e401ee9fe0125ee911d0290e295a0507c1239c"></a>

## Next pages — psp_spec.allowed_host_paths / 104c284af9db / 6

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-d2518c9bf23bb7b46ceea909901a560b9b33d09acb7fc02aaf53ecda215bedf2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-703c6e7a23bcffd2d910b5423ef61814a93d66f5e71f17944341ff6469ef643c"></a>

## psp_spec.default_capabilities — psp_spec.default_capabilities / 7832c837700a / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.default_capabilities

<a id="canonical-68080bc2b2b8948c9ec1b8de07f34dd5130283c25120e3f78dd4d95b1eec630d"></a>

Type: `"single"`. Computed.

List of capabilities that docker container has.

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

<a id="canonical-0796f3aa189565b5567cb4c63a4d401f873da821b7aed38bd1b6575043c5a4df"></a>

## Direct properties — psp_spec.default_capabilities / 7832c837700a / 3

<a id="canonical-75e8b0eaced6512172a4fb90c930f0cc0a1322d513012136358b074931afcf67"></a>

<a id="canonical-c3e246cf616077efac801a96eed1989bfdeed8bd8d71400bce243b32dbd87dd1"></a>

## capabilities property — psp_spec.default_capabilities / 7832c837700a / 4

Type: `["list", "string"]`. Computed.

List of capabilities that docker container has.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-736ce3502904cbaa5a67187afefdeeaeec533170d944520b8af862dbd0410634"></a>

## Next pages — psp_spec.default_capabilities / 7832c837700a / 5

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-30152ad1cc9e813ce8b113a10ef4ae80c8862d6f591e4315b3211ce13163016c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7d7fdad94805a6c28e97532b8737fd5f670a834599f83d18d8240ebabc8f9af"></a>

## psp_spec.drop_capabilities — psp_spec.drop_capabilities / 27ff768947d6 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.drop_capabilities

<a id="canonical-c9b31295914fdbc58bcdbf9d163366f8fc7d3e6e03cded4accba3500d5c3a880"></a>

Type: `"single"`. Computed.

List of capabilities that docker container has.

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

<a id="canonical-e375937fd8190f7b5e8f3e3cfdeffc6fdfd7868166435d5d486b019c85b5e716"></a>

## Direct properties — psp_spec.drop_capabilities / 27ff768947d6 / 3

<a id="canonical-959bc37bdf8b6464171d9b14c8247262f048aa91befaa99b96a201072571cf8d"></a>

<a id="canonical-c92d727993fcced27f7918a19aa25ece949261ca8939cdf47514fab64516e09d"></a>

## capabilities property — psp_spec.drop_capabilities / 27ff768947d6 / 4

Type: `["list", "string"]`. Computed.

List of capabilities that docker container has.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-78dc511ea3dda37800b81d5451d73d3f72db79be5b13aad20341ffc04d09ba4d"></a>

## Next pages — psp_spec.drop_capabilities / 27ff768947d6 / 5

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-51fd42d4b5bc8fd40ed4fd1ad317b7774a6ffc912041eba74f4523520ace9e2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af87247c7063fb40ee17f1904aed19997cf3a12df987a0b17f4acdf153f4126f"></a>

## psp_spec.fs_group_strategy_options — psp_spec.fs_group_strategy_options / 691a843d121c / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.fs_group_strategy_options

<a id="canonical-1be7147fd868436517aa5d1bc29fb6872451cda426449d07a555e7ead991a4b7"></a>

Type: `"single"`. Computed.

Configuration parameter for fs group strategy options.

Upstream description:

ID ranges and rules.

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

<a id="canonical-42892a18a4f6a24f0369f41db1da2001decae7c947e68d9d76345a0d451859fd"></a>

## Direct properties — psp_spec.fs_group_strategy_options / 691a843d121c / 3

- [id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-09a125fd139999e2dbc8e15fb993167973a9af0bb52c25061243894f335ee832): complete subsection reference.

<a id="canonical-dca9669a74ad6fed3caa1c7e530f2ce63c72cb3447b8b4252511fa5eb5e9e72a"></a>

<a id="canonical-e7e6f3f86123e35c0d616486f3ea92dd890c01d646f54adbe505a186239995cb"></a>

## rule property — psp_spec.fs_group_strategy_options / 691a843d121c / 4

Type: `"string"`. Computed.

Rule indicated how the FS group ID range is used.

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-e504febb13b00b0d667166d03efea266e6daaecccf30ee0388236c6554cda65d"></a>

## Next pages — psp_spec.fs_group_strategy_options / 691a843d121c / 5

- [psp_spec.fs_group_strategy_options.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-09a125fd139999e2dbc8e15fb993167973a9af0bb52c25061243894f335ee832)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-09a125fd139999e2dbc8e15fb993167973a9af0bb52c25061243894f335ee832"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c77ede08dcf3fec987263cd583df5ad9a7a4550be3ac810cce2f04b744b614a"></a>

## psp_spec.fs_group_strategy_options.id_ranges — psp_spec.fs_group_strategy_options.id_ranges / dfb99ec9d149 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [psp_spec.fs_group_strategy_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-51fd42d4b5bc8fd40ed4fd1ad317b7774a6ffc912041eba74f4523520ace9e2d)
- psp_spec.fs_group_strategy_options.id_ranges

<a id="canonical-7b615fb83c5df4108d8950e170cdcd6c907637546b4d6603718dc5203f93cc38"></a>

Type: `"list"`. Computed.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

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

<a id="canonical-44c982c07d56a19f8e07e5b47f6ef6b3c85733e560d5804592985be2e3e4d346"></a>

## Direct properties — psp_spec.fs_group_strategy_options.id_ranges / dfb99ec9d149 / 3

<a id="canonical-65ab42b898e880923e589ffd692e4590b5f6e5d8f607ba1c1918568adadd0c65"></a>

<a id="canonical-783753f58adb1b9c98210b02464e98832e1a2e1de7566e9948e77dd982ce7f42"></a>

## max_id property — psp_spec.fs_group_strategy_options.id_ranges / dfb99ec9d149 / 4

Type: `"number"`. Computed.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-a366dc03f1f03b850b81433bec25c3b5c01be7f973a35a3f10ebddd31f51d2f9"></a>

<a id="canonical-611575b7545f4147785a9b2ecdf3f2f26965c793ac748c01d2e322e219817097"></a>

## min_id property — psp_spec.fs_group_strategy_options.id_ranges / dfb99ec9d149 / 5

Type: `"number"`. Computed.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-cece8431c5f411bde77a113b33d2f47731115baf188c49c6c8a38de598ffe5b9"></a>

## Next pages — psp_spec.fs_group_strategy_options.id_ranges / dfb99ec9d149 / 6

- [psp_spec.fs_group_strategy_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-51fd42d4b5bc8fd40ed4fd1ad317b7774a6ffc912041eba74f4523520ace9e2d)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-841db374e59e81363400bbfed0efb1293e47e6aae84be8de6b6433d4f5e2ac93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21bb5ed6b956f8c583ab6fb88bf3b293c088616d6784ec5d63d3144e7588897f"></a>

## psp_spec.no_allowed_capabilities — psp_spec.no_allowed_capabilities / f899ab09be25 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.no_allowed_capabilities

<a id="canonical-12f4e74ae1fbac18015dfacff7ffb6b1caf50d094671d3d1133be5118ba9d78f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no allowed capabilities.

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

<a id="canonical-703f6f38046e36aae21a2d041ba8a3b237da48df981b002fff623c067b368d5b"></a>

## Direct properties — psp_spec.no_allowed_capabilities / f899ab09be25 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5f2062c6c81d3c1eb05da1f44437e597469e5ccd35f635c82d124e0e17ddfcf2"></a>

## Next pages — psp_spec.no_allowed_capabilities / f899ab09be25 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-c7d4de5ef51bd5a3992551b739673fdebb4f7f6f69797644c1695bc0bbd28420"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e09cabb04d94f5bd2cb19b5ddb87fe694fa50a043b6ecc2a6740c9f8f49877d7"></a>

## psp_spec.no_default_capabilities — psp_spec.no_default_capabilities / 11a86b037842 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.no_default_capabilities

<a id="canonical-bc44c296b512f3ac983d21be4d728374523e513f325bd4f20ed1c503f811d77d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no default capabilities.

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

<a id="canonical-5a058673b924f302f4f1ca2c8323a3f8ed503e5308fba5261e613c689f232957"></a>

## Direct properties — psp_spec.no_default_capabilities / 11a86b037842 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aff04c3f16620d6b892b8210ab294915e92642c72a267b3c8e9922942245730f"></a>

## Next pages — psp_spec.no_default_capabilities / 11a86b037842 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-0a4908ff2f6784d6eb84b1f29b4b0236bb48c2235083ffa3590d9bc4286118f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30171897a4290e6397e3575270835cf745bba9c807f90d15c700ef605e6c6816"></a>

## psp_spec.no_drop_capabilities — psp_spec.no_drop_capabilities / ae0933cda934 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.no_drop_capabilities

<a id="canonical-5304c9dc8eb2dc467cd9c7a7ca8fc11695e85b02a9edc31540bfa1515b97bacb"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no drop capabilities.

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

<a id="canonical-604358692662ba1fe115f679889b84e7bcd8f496c0edf543b42daa3ccd9b1a9c"></a>

## Direct properties — psp_spec.no_drop_capabilities / ae0933cda934 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6b2ba73381bdfb0eeb03d3aa5305af4fd9a2ce9db300d396fd01b563a43b75b4"></a>

## Next pages — psp_spec.no_drop_capabilities / ae0933cda934 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-9866f9d9a79532a089b03e0a3cc49eeb933535fb8cc0c3d8969266568a99c6fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a103bdfbc62f3956771874f91d76d4e4568a3c663ca501d845c9ab5306b9de38"></a>

## psp_spec.no_fs_groups — psp_spec.no_fs_groups / a6a69fa02b21 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.no_fs_groups

<a id="canonical-813f822b1750d53f0c1df32bce0108d909f4fabaee42aff3c59528fbab0ebd24"></a>

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

<a id="canonical-ebe1cbfd83af93aa4993013847f9e85cb95e4ab985e6bb3a03063c87fcecb91b"></a>

## Direct properties — psp_spec.no_fs_groups / a6a69fa02b21 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8b61714eca819b69d7c773afae97c3b09f29cd21106bb4d3e249b7ffd9990062"></a>

## Next pages — psp_spec.no_fs_groups / a6a69fa02b21 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-41c9a33a6fcdfb6ab57ec21a67959b5ab9c8734c0ad03b6e77aee0d5a84cf4db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70db550fe25ca17827a6c7a0d6f927e61d50253d88cf38aa9c9c01f5dd542b0b"></a>

## psp_spec.no_run_as_group — psp_spec.no_run_as_group / c261c78b22ef / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.no_run_as_group

<a id="canonical-801004e949d1b8939a65a25b0a95a808b8093ba1358cdcc25c077febce95d2a8"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no run as group.

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

<a id="canonical-013f985c5acc8e9bcbf771221f8c4cdf1ffbcc07567acb918c8d8ccaaf628042"></a>

## Direct properties — psp_spec.no_run_as_group / c261c78b22ef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8cad719ef2dfe38fef6af63d89526e52dac83fecbd10ceb769ea576718f3cca2"></a>

## Next pages — psp_spec.no_run_as_group / c261c78b22ef / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-2dd0f86c2c62b09639dbd851b75468be17ddb1f8af167d09088fc7ced091c635"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6aae747be6d844a3f3252a842bbc64657f29f94e179e4a14da3007bcabb4c00e"></a>

## psp_spec.no_run_as_user — psp_spec.no_run_as_user / a4d49d05d84d / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.no_run_as_user

<a id="canonical-6d8d38006b2808c79050c85eb84fecc577969b72d6c96ccc51db231184652dcb"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no run as user.

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

<a id="canonical-c8185699443f8ffca70fe4da7424c2addc8c0ef00eea8d93b324c0ade10bdd86"></a>

## Direct properties — psp_spec.no_run_as_user / a4d49d05d84d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa751c5bed40dd23ef62e028713f85670bce9a81020ae7b611666acf719f5d21"></a>

## Next pages — psp_spec.no_run_as_user / a4d49d05d84d / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-b80da28e605174bf8c9ae4bb13cdcbe172e3794d29ed6e7c3ccdf71252286c11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33a25c5f0fbc286998c3f225c59e0a58233d67d00f54d6a4309e7e15a34750f7"></a>

## psp_spec.no_runtime_class — psp_spec.no_runtime_class / 3a9a7ea34525 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.no_runtime_class

<a id="canonical-e2467cdb00e42f7dc501dfd60ef7c1e7253e6661062295e3358a1a37b82bf0f2"></a>

Type: `"single"`. Computed.

Configuration parameter for no runtime class.

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

<a id="canonical-34aa71d7ec983bc250fd60d1a53a0c327a9e5111ae5cc0b03054f6a1c5a84cdc"></a>

## Direct properties — psp_spec.no_runtime_class / 3a9a7ea34525 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-049f4fd2f85980d1f5e030672c52a10b16c1258ac8e0f5057dede484f45b0a18"></a>

## Next pages — psp_spec.no_runtime_class / 3a9a7ea34525 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-cf67529b6fb97da5404fc2dfe034cca85957a9ee33695af79469ff981004b35b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e750086e4745cacf7910ed385b1e7c5de5669ebd0b55c2213390e622a865abdd"></a>

## psp_spec.no_se_linux_options — psp_spec.no_se_linux_options / 09e751cd7287 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.no_se_linux_options

<a id="canonical-3c79c9e1d961072a5ede788abcbb72297509650d596f681a4b0eb5ebbb1e412e"></a>

Type: `"single"`. Computed.

Configuration parameter for no se linux options.

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

<a id="canonical-498b36e84841c9e9a96aa996eb861412c3980bb54ee7163e09671c8635055dba"></a>

## Direct properties — psp_spec.no_se_linux_options / 09e751cd7287 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8a738ddc8e5fdc88c78335e49ca9ad6a8036a2078eec71e62d36ea30f9d9f937"></a>

## Next pages — psp_spec.no_se_linux_options / 09e751cd7287 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-4059018accd2b392bd013e0e0d592e5b9092256c373f645383ddd993b9a1f07c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4cdabec99576f856d81f753bd22957cf30f31873f19bded2658ac623b8b65f2"></a>

## psp_spec.no_supplemental_groups — psp_spec.no_supplemental_groups / 2d0941d88f84 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.no_supplemental_groups

<a id="canonical-930b0021bdb74f5030b2279262b1b6b754290b5b57090e3f6148c0f14899feb3"></a>

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

<a id="canonical-cb0180b35d97b40d34ca99bfda486672b8cea8694920e6f610c44e39730c7642"></a>

## Direct properties — psp_spec.no_supplemental_groups / 2d0941d88f84 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6d9e9a0a3e7d989dba174b7324989c016679d9f7d032d52445e1a4786346d200"></a>

## Next pages — psp_spec.no_supplemental_groups / 2d0941d88f84 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-40f047a707869ebaf47acc5a2d4e28f9e3002f0e0b3101d9e399a453bd12acbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa31ddb7f062fe8dbcc42f7f4f8c9340a0ee867f462ea2d476f3a23caccd0c46"></a>

## psp_spec.run_as_group — psp_spec.run_as_group / 49333dc2ca8c / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.run_as_group

<a id="canonical-4309f036ea5f27ae738777d8c6bf0bfa846f265d786e78c7ae0eff280c8c0ec5"></a>

Type: `"single"`. Computed.

Configuration parameter for run as group.

Upstream description:

ID ranges and rules.

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

<a id="canonical-74449c9327ee2f23f41a6b20616845b0c0940c660dcf643a250df1ab118bc042"></a>

## Direct properties — psp_spec.run_as_group / 49333dc2ca8c / 3

- [id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-6237ebf107ee7a43881418f0c659eeeacec1ed01f76f3dbaa5f5d62dfb3817a5): complete subsection reference.

<a id="canonical-4c7df6e8fd9fc7744c8b3aa80022856d9bb2e70c8b12ee8e52731f55d73c5770"></a>

<a id="canonical-988a113a9b550416732e7fe9d84e7e71ec2e4481e2109c1a358dcfd8af1a8762"></a>

## rule property — psp_spec.run_as_group / 49333dc2ca8c / 4

Type: `"string"`. Computed.

Rule indicated how the FS group ID range is used.

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-52886c7f3e4557ef912e7eb06461792dd83482aa34cb99f7d39619667ef1cb4f"></a>

## Next pages — psp_spec.run_as_group / 49333dc2ca8c / 5

- [psp_spec.run_as_group.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-6237ebf107ee7a43881418f0c659eeeacec1ed01f76f3dbaa5f5d62dfb3817a5)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-6237ebf107ee7a43881418f0c659eeeacec1ed01f76f3dbaa5f5d62dfb3817a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-801c6d468a0103dcff52bf05a22f9ef0081623424b51984d11338025948d9248"></a>

## psp_spec.run_as_group.id_ranges — psp_spec.run_as_group.id_ranges / 47191e863be7 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [psp_spec.run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-40f047a707869ebaf47acc5a2d4e28f9e3002f0e0b3101d9e399a453bd12acbf)
- psp_spec.run_as_group.id_ranges

<a id="canonical-02bc732c95b06d84c7775ccb5b9c104421e9bd0c03274d923490a00ba092e40a"></a>

Type: `"list"`. Computed.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

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

<a id="canonical-863830bab2a23c965e442ff4fc2c84113a2e099810a32a0e05368450867060f1"></a>

## Direct properties — psp_spec.run_as_group.id_ranges / 47191e863be7 / 3

<a id="canonical-c7db105d4063f032b1748af25d462230777f484a6cb232f328afd24efcd158ea"></a>

<a id="canonical-33174265b2f040702892619a7f6bced1c492d86a5390b547956a24ac89119c1b"></a>

## max_id property — psp_spec.run_as_group.id_ranges / 47191e863be7 / 4

Type: `"number"`. Computed.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-662ab704c1f3a5cbb8f874e369e42ee6dbe880ed7bcca20d18b150d87c579032"></a>

<a id="canonical-4143b9294f038a62902cab20e9ca66e6f7e7646710ab34c332cac49d8a82ea48"></a>

## min_id property — psp_spec.run_as_group.id_ranges / 47191e863be7 / 5

Type: `"number"`. Computed.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-017db70586439aaa5d4a62377abe5c58050baba5c3818585522f2fc8ff9e0a96"></a>

## Next pages — psp_spec.run_as_group.id_ranges / 47191e863be7 / 6

- [psp_spec.run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-40f047a707869ebaf47acc5a2d4e28f9e3002f0e0b3101d9e399a453bd12acbf)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-a5c7a070654ac9a6ea582fcd39303366b6aaf3a888f64347ad6f2e5d933e1235"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8a804011baabd05e0fe98a532ffda1df50ca9dc1ee362e271a5a9e92f184613"></a>

## psp_spec.run_as_user — psp_spec.run_as_user / 6bbe7c047847 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.run_as_user

<a id="canonical-196dfed234d9aaa9539524eda28486581757688b54f4998ba21cab63a6491e55"></a>

Type: `"single"`. Computed.

Configuration parameter for run as user.

Upstream description:

ID ranges and rules.

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

<a id="canonical-eef8dda121f70c0ea22e2d4dff43dd0db1dd9bbccd56aa17e4ff4e2375c9bd79"></a>

## Direct properties — psp_spec.run_as_user / 6bbe7c047847 / 3

- [id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-e2027ae19179d8d105f050d7016531650929e75063c1d66c6f1fa10446309493): complete subsection reference.

<a id="canonical-d4cb969e5706418f180d93dc0fd6b9ba59c18f30b1c1aad2b67e64f40433eb4d"></a>

<a id="canonical-97abb23535eef17821dec42117034ee84ead495153d754453dee67096c7827e6"></a>

## rule property — psp_spec.run_as_user / 6bbe7c047847 / 4

Type: `"string"`. Computed.

Rule indicated how the FS group ID range is used.

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-e54f770e9e719fce01fdab7a56aa4e934b20e7f3794af7bcadbe79563cb5987f"></a>

## Next pages — psp_spec.run_as_user / 6bbe7c047847 / 5

- [psp_spec.run_as_user.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-e2027ae19179d8d105f050d7016531650929e75063c1d66c6f1fa10446309493)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-e2027ae19179d8d105f050d7016531650929e75063c1d66c6f1fa10446309493"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dac5854047fae5a24c02df34033fb5e34010c9af217603cfadb86c2fe79b1fce"></a>

## psp_spec.run_as_user.id_ranges — psp_spec.run_as_user.id_ranges / d5d17e68518d / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [psp_spec.run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-a5c7a070654ac9a6ea582fcd39303366b6aaf3a888f64347ad6f2e5d933e1235)
- psp_spec.run_as_user.id_ranges

<a id="canonical-a7458e60ad4dcc305058b000e00872c0a588a61f8490ae80690a299374a5c9a6"></a>

Type: `"list"`. Computed.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

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

<a id="canonical-7caa69a06497a2aba541c96cf6cb2e4c658a1dac6a8f03719b3800c46fdd7bb9"></a>

## Direct properties — psp_spec.run_as_user.id_ranges / d5d17e68518d / 3

<a id="canonical-a37f954cd4edb77962d13ba06e6403a35c154800a7cbac9341bea513c3856d95"></a>

<a id="canonical-52fa194d709c867d69b883818403f701e62224318be83fe979036f0708ffa95c"></a>

## max_id property — psp_spec.run_as_user.id_ranges / d5d17e68518d / 4

Type: `"number"`. Computed.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3dd75acdf63d930de0ddfd35ede764e741b785f4f4d86f28a76032ee8796617d"></a>

<a id="canonical-a412858e5fc0c2fc8d63444e610b417df129ff197cc7b1055c9bf1a6a72549dd"></a>

## min_id property — psp_spec.run_as_user.id_ranges / d5d17e68518d / 5

Type: `"number"`. Computed.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-49a8325130f03f1fdd586819c99cb10eea281a899cd853c622361e312ce04b09"></a>

## Next pages — psp_spec.run_as_user.id_ranges / d5d17e68518d / 6

- [psp_spec.run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-a5c7a070654ac9a6ea582fcd39303366b6aaf3a888f64347ad6f2e5d933e1235)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-36b51f79a7eb887beedd4fba4c6a4c3d3ac83da8b5ce3b054b49099e6e08c558"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7aadba0c0cfc203bb92c4531f55a46da75bdfc32805d0a6002a4bcf8be42516"></a>

## psp_spec.supplemental_groups — psp_spec.supplemental_groups / 36e921a91166 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- psp_spec.supplemental_groups

<a id="canonical-c3d9a39fe272415824b4e0ebddaf5fde7b649fdab6995d1683226fca664239dd"></a>

Type: `"single"`. Computed.

ID(User,Group,FSGroup) Strategy. ID ranges and rules.

Upstream description:

ID ranges and rules.

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

<a id="canonical-5c5a6af0df218277a12d7b90831a40e32cbd915fa14dc1b51b2ec281194e453b"></a>

## Direct properties — psp_spec.supplemental_groups / 36e921a91166 / 3

- [id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-ca73af750c6e99c4c341ecc4b93c63221aabb9e9707301cfa6a3dc54bc5b1ff9): complete subsection reference.

<a id="canonical-5c3b3ee79fe17c5eb8b271997e43442f93751dc4a60f9148486f804439bc47f4"></a>

<a id="canonical-d93aef872e31423f6605346a28566cb33c73788b52b3bec4d1be597b021e6442"></a>

## rule property — psp_spec.supplemental_groups / 36e921a91166 / 4

Type: `"string"`. Computed.

Rule indicated how the FS group ID range is used.

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-061ed3bed7af18ee64ad6c302468cd748872a1193ad1b26b41e158a90b225979"></a>

## Next pages — psp_spec.supplemental_groups / 36e921a91166 / 5

- [psp_spec.supplemental_groups.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-ca73af750c6e99c4c341ecc4b93c63221aabb9e9707301cfa6a3dc54bc5b1ff9)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)

<a id="canonical-ca73af750c6e99c4c341ecc4b93c63221aabb9e9707301cfa6a3dc54bc5b1ff9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-656018ca13fcad24368223222ea792c6eb646b4a545f707fd49acb7649fc1629"></a>

## psp_spec.supplemental_groups.id_ranges — psp_spec.supplemental_groups.id_ranges / 968c1791531c / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-184ccdb883ac5057158a4e9183a15e9800253e32c385953e20b54c27443f4ac7)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-be33e1ca168cdded2fe64bb027ec84a68d6a598c6decc8689192d34888aff68f)
- [psp_spec.supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-36b51f79a7eb887beedd4fba4c6a4c3d3ac83da8b5ce3b054b49099e6e08c558)
- psp_spec.supplemental_groups.id_ranges

<a id="canonical-811133fd8b28c7fb579475e6a4e205256a1fb28c7433f822ad868677c194c6d5"></a>

Type: `"list"`. Computed.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

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

<a id="canonical-6b1bda18dc16de4a72acb04d445842bf95c5dd302be0ab3f5ebbd0c4c8ad120f"></a>

## Direct properties — psp_spec.supplemental_groups.id_ranges / 968c1791531c / 3

<a id="canonical-04c36df85b435e88269175eff558baf2bc8916b195f9c5150e62893dadcb8e34"></a>

<a id="canonical-2bfa5308393837ee24f743665c61a101bff27f74337c6b9dc431c5ca11710426"></a>

## max_id property — psp_spec.supplemental_groups.id_ranges / 968c1791531c / 4

Type: `"number"`. Computed.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-06bbbd3220fbef753adf35b9c1def7cfad32b9bc7174e2ee7cf0253230bd0ffa"></a>

<a id="canonical-53142412177e12f2b0fce856c78704d7729d16d6d772fe46e48d36f1a919d68e"></a>

## min_id property — psp_spec.supplemental_groups.id_ranges / 968c1791531c / 5

Type: `"number"`. Computed.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-5248cd877041785a3bf78466f7fb170502f844aca88a90d734a33fb83cf0322e"></a>

## Next pages — psp_spec.supplemental_groups.id_ranges / 968c1791531c / 6

- [psp_spec.supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-36b51f79a7eb887beedd4fba4c6a4c3d3ac83da8b5ce3b054b49099e6e08c558)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-36f29df3c9c08e8147f6d5102e2fa3a66d82fbfbb99ed811ed591f69ee6ab6ec)
