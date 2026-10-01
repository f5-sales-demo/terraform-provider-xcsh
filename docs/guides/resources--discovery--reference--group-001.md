---
page_title: "xcsh_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery reference."
---

# xcsh_discovery reference

<a id="canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8985f3b10f03447ecf1f31ade28953dc1e47b38172ca69ee4e613bd052f45fc3"></a>

## Property reference — Property reference / 9d110a78f57f / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- Property reference

<a id="canonical-16906a6f52e33a4ea0c4faa27015caa0143fc091b2a0b2bad413c896ef55aef9"></a>

## Direct properties — Property reference / 9d110a78f57f / 3

<a id="canonical-eeec16ef971a500852b223926f79a90a881e4f044e2d7fb06309d00d31799de4"></a>

<a id="canonical-f760ca0f56961de9dd81483915eba08fd3c25d051bc9853466a353eabfd81c23"></a>

## annotations property — Property reference / 9d110a78f57f / 4

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

<a id="canonical-021a068607cac4446298c5cb14e2c4bd10f6c6b98258ac69318cb00245923d4c"></a>

<a id="canonical-e7471a4e2d15db35992eeaf668f405281bbb0f50915a233b9596642886e75d5f"></a>

## cluster_id property — Property reference / 9d110a78f57f / 5

Type: `"string"`. Optional, Computed.

\[OneOf: cluster\_id, no\_cluster\_id; Default: no\_cluster\_id\] Exclusive with \[no\_cluster\_id\]
Specify identifier for discovery cluster. This identifier can be specified in endpoint object to
discover only from this discovery object.

Upstream description:

Exclusive with \[no\_cluster\_id\] Specify identifier for discovery cluster. This identifier can be
specified in endpoint object to discover only from this discovery object.

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

OneOf alternatives in this subsection:

- [cluster_id](resources--discovery--reference--group-001.md#canonical-021a068607cac4446298c5cb14e2c4bd10f6c6b98258ac69318cb00245923d4c)
- [no_cluster_id](resources--discovery--reference--group-001.md#canonical-c2ecca3a0882dca2d50b0e28ca44f30a89df074c30743106f2493b41dc069098)

Select alternatives according to the provider validators above.

<a id="canonical-2c97e2c9d0835faea6a9ba65118e45857acee79a6a7308bab90d8a7a600f9dec"></a>

<a id="canonical-f16da54526d177a8fde39baa716d2683d0261390a0aa3df168e82b57afb6c107"></a>

## description property — Property reference / 9d110a78f57f / 6

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

<a id="canonical-67d04620a0a5d30a33433cbf820dc0240c3338b27c20af469e5752f4ffa2cb90"></a>

<a id="canonical-539c73ef786394208047413eb24a5888638b40d582bc56efb29d3e7f384d7e02"></a>

## disable property — Property reference / 9d110a78f57f / 7

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

- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f): complete subsection reference.

- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0): complete subsection reference.

<a id="canonical-bd87c234c5f276cce3e83095107de478f0d8752ec43e47fa69b5fd7124f1aff5"></a>

<a id="canonical-d54b74a11fea389ba501995f6c047a669dc85ceb9bed0e07d4b4c0ff45b4cf8f"></a>

## id property — Property reference / 9d110a78f57f / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-a8452560bbd17f961e0a0b765b14a14be09290636d0ecb70f99275e0f5679094"></a>

<a id="canonical-9ef420a87381c203d27d438486d6a1d861e6a8ee7643a1fb6575238a81aaad0f"></a>

## labels property — Property reference / 9d110a78f57f / 9

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

<a id="canonical-2d4eb7f274286597736027ab5354dbe4b25815ea913f9e50cf08191a55f4b9e1"></a>

<a id="canonical-55aef9cd7f986636c10d1c957531aac5a9a473f42704c4dd74c11dd2130d3194"></a>

## name property — Property reference / 9d110a78f57f / 10

Type: `"string"`. Required.

Name of the Discovery. Must be unique within the namespace.

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

<a id="canonical-cd9ce63907ba47f67ff4cdbb57ea89ff1078af7f468a94c81c05324905bf0c81"></a>

<a id="canonical-a0a386a1f0445cb5f4396449c55b2fb85393448e0efddcd6fc3b84a47a292161"></a>

## namespace property — Property reference / 9d110a78f57f / 11

Type: `"string"`. Required.

Namespace where the Discovery is created.

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

- [no_cluster_id](resources--discovery--reference--group-001.md#canonical-be07760a77ec25ad8978b2bd1d1aaf9fedbe1488123a6d5e645d989c408de7d3): complete subsection reference.

- [timeouts](resources--discovery--reference--group-001.md#canonical-1f6a6e1efb3a73cacd0583b5d8d76db229ff20e709ec7d1ec796cdcf7d9d8670): complete subsection reference.

- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04): complete subsection reference.

<a id="canonical-1d7f938409fa4000d26eeb4e4219d905ef63f435d4dc7467d34fbe8e481e9d78"></a>

## All schema paths — Property reference / 9d110a78f57f / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--discovery--reference--group-001.md#canonical-eeec16ef971a500852b223926f79a90a881e4f044e2d7fb06309d00d31799de4) |
| `cluster_id` | [cluster_id](resources--discovery--reference--group-001.md#canonical-021a068607cac4446298c5cb14e2c4bd10f6c6b98258ac69318cb00245923d4c) |
| `description` | [description](resources--discovery--reference--group-001.md#canonical-2c97e2c9d0835faea6a9ba65118e45857acee79a6a7308bab90d8a7a600f9dec) |
| `disable` | [disable](resources--discovery--reference--group-001.md#canonical-67d04620a0a5d30a33433cbf820dc0240c3338b27c20af469e5752f4ffa2cb90) |
| `discovery_consul` | [discovery_consul](resources--discovery--reference--group-001.md#canonical-4ccd1f79e3cb2f0b095dd4399f3d1634d03c05633561eaaa1f7e7e00e1f98141) |
| `discovery_consul.access_info` | [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-fcce2b779b3834aace9ac65c315ac280a670e8bbd45e1ba8c9bf41c0b15de435) |
| `discovery_consul.access_info.connection_info` | [discovery_consul.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-887ae8127419ee0d4974cf8c87ae7aa7f645dfaa714ddadba5bc4e3a35a44244) |
| `discovery_consul.access_info.connection_info.api_server` | [discovery_consul.access_info.connection_info.api_server](resources--discovery--reference--group-001.md#canonical-5fcea4bb069eb56f5738319f93d7a6a2de599f41f2a82c58ad5276701b3bc4d4) |
| `discovery_consul.access_info.connection_info.tls_info` | [discovery_consul.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-d87249e9748887a627837964dae8948e90b5ad061840e163919c9e54bfcc4f74) |
| `discovery_consul.access_info.connection_info.tls_info.certificate` | [discovery_consul.access_info.connection_info.tls_info.certificate](resources--discovery--reference--group-001.md#canonical-b1b27ef6298b3534e4e9982ba75104c00a1b76615ee50d74ab0de3df81f23509) |
| `discovery_consul.access_info.connection_info.tls_info.key_url` | [discovery_consul.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-3a38751d1df9fc84db01f9ff7f72051754084604881585f2196d2d1e7767f7b8) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-9fb207fa805305054b7193b5502cc2b33a8754a0275a186d6574f57331cd5a6a) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider](resources--discovery--reference--group-001.md#canonical-e8a26c997ea7cf51e11b8b4a942c314796e7f6c4d52834bec001eccd913c91c6) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location](resources--discovery--reference--group-001.md#canonical-7a4cc7629c777812fbbac8c416b4c42a961a1a85549b847027403a8c28de82ea) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider](resources--discovery--reference--group-001.md#canonical-7b70c4adffebf4f9fd7c0246edb87d8e56123522ac57665e7f974def6664bf5b) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-f78185eeacc447937744790414924c29af90453a32c6c23b38d9091621722d58) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref](resources--discovery--reference--group-001.md#canonical-a43e680cc9edd11a5ccddf854dd5deefc5c32753106ad688784180696b2daada) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.url` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.url](resources--discovery--reference--group-001.md#canonical-9cabac77c271560dead570ad236879c30b3b0ffc17a5f33d99b6250daf8678d7) |
| `discovery_consul.access_info.connection_info.tls_info.server_name` | [discovery_consul.access_info.connection_info.tls_info.server_name](resources--discovery--reference--group-001.md#canonical-9795870f1d67f40a5a38b12beb58057ef76735d6580a202bc2ca210cea9d8b40) |
| `discovery_consul.access_info.connection_info.tls_info.trusted_ca_url` | [discovery_consul.access_info.connection_info.tls_info.trusted_ca_url](resources--discovery--reference--group-001.md#canonical-66aecc413c2263aaa2c529b60ca283f1702fbb31a73398d9e0783c5c97e6b3e0) |
| `discovery_consul.access_info.http_basic_auth_info` | [discovery_consul.access_info.http_basic_auth_info](resources--discovery--reference--group-001.md#canonical-1bdfb91d2cc0a057d2d0fc3d515ebfbfb8bb16c384395487e8a7b5f01b495be7) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url` | [discovery_consul.access_info.http_basic_auth_info.passwd_url](resources--discovery--reference--group-001.md#canonical-6ec4955ba2faa36420479b136a3af64fe472ff8c13ee3a2ec87fa4c303e44fc4) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-85082fb61629ebc1eb811da3eb132a19a47ef42db3782a2716c82a55ef61f6d2) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.decryption_provider` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.decryption_provider](resources--discovery--reference--group-001.md#canonical-1e1c86773c19f42e7c93616a129ea4a09952ee1eced83943e13e1163177231f7) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.location` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.location](resources--discovery--reference--group-001.md#canonical-6d20e1a79dafa4964e57cdc861f521279fd68c70df743f510fe9966b70af0114) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.store_provider` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.store_provider](resources--discovery--reference--group-001.md#canonical-2e29c078795803aa816c34c2ba363efbf30c05f499a0751ba4b6b146211ac839) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-c85330efd037032dada0ec297f3630f9d3b9d914348106bb30f6e8271dad81d3) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.provider_ref` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.provider_ref](resources--discovery--reference--group-001.md#canonical-2124f659812902bb5abc18604e41a930be2af51d1a9ff3afd3773678973afe0e) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.url` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.url](resources--discovery--reference--group-001.md#canonical-fd9b12725b67f0b52dd2a764a1d2b1473fa84e352efc9250813bca7418700974) |
| `discovery_consul.access_info.http_basic_auth_info.user_name` | [discovery_consul.access_info.http_basic_auth_info.user_name](resources--discovery--reference--group-001.md#canonical-9e8b8abb595f3d8596bef243a890568d66577b044c81f46754f8c54ea922b433) |
| `discovery_consul.publish_info` | [discovery_consul.publish_info](resources--discovery--reference--group-001.md#canonical-16674cae4d9e72cd6e857f55e9844da807c1d8dff0a7917b393db2a7bcec0ba5) |
| `discovery_consul.publish_info.disable_spec` | [discovery_consul.publish_info.disable_spec](resources--discovery--reference--group-001.md#canonical-e959c35d258d1d227c442e7c766da8832f0ba3545a25bf84cdb0d50e7be66019) |
| `discovery_consul.publish_info.publish` | [discovery_consul.publish_info.publish](resources--discovery--reference--group-001.md#canonical-6132d3c55a2dbe0d39202def599f4d729e7d42e6b4f87a6ae27ab4496d3a1b7d) |
| `discovery_k8s` | [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3339faadde78d59c922be5812f8ccbfec57ac84a5d34a03442b1c7cf749b2268) |
| `discovery_k8s.access_info` | [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-1c1cd05eab3ab66a7766da3bb040696f424547d3a6e17d2d349546ac49b93408) |
| `discovery_k8s.access_info.connection_info` | [discovery_k8s.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-c5a1118e72b778babef3bad0f9979e5d573a1eb78fadb3691b908afe0397e57e) |
| `discovery_k8s.access_info.connection_info.api_server` | [discovery_k8s.access_info.connection_info.api_server](resources--discovery--reference--group-001.md#canonical-20a514a1f303a368ee3abd02cc1c2e0736cd5fea6e98c2eaf4251cd80c25e3ec) |
| `discovery_k8s.access_info.connection_info.tls_info` | [discovery_k8s.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-cc29e4808d4a1ed2d742a13cf09d78cd51bdb081c46137b8084c3b3453d6b6ab) |
| `discovery_k8s.access_info.connection_info.tls_info.certificate` | [discovery_k8s.access_info.connection_info.tls_info.certificate](resources--discovery--reference--group-001.md#canonical-d4ec998a40dc6bdbc74f5cd10a05d7991f508e8149f1956386bb8a8bbc904462) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url` | [discovery_k8s.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-a821f7853404febdaf408e201371a8c4ec327190ffa2ef2f9fbfd1d88389d1c8) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-dc76b0435b05b79d77b59c6f864d1219108f5e4e4e7bfbfafbc4c6be8ab49883) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider](resources--discovery--reference--group-001.md#canonical-36b2a4344974a47bf9aaa5de1027149982ede892717bd03bad7903319866c24e) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location](resources--discovery--reference--group-001.md#canonical-69ede6b5dc7a16fd87e1984e5720a6f8892abc35a59579e643b07702e06ec677) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider](resources--discovery--reference--group-001.md#canonical-bc243a62f03a2f6655740bfc19677e2ad5f803f7fafa6ae1493bede9ac43c262) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-195abd05a054d6752775ff2e91c341f8d271863ee322ef11d09812f207f5eed7) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref](resources--discovery--reference--group-001.md#canonical-bde8ac1bb4e45ad7dd809d9d3c1739599971cd5b24d1a854743e3bbaf4575bce) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.url` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.url](resources--discovery--reference--group-001.md#canonical-b17be070a06f89802fed2af2da64e0dc4fc19f2dc3ecc7e4b834d673d32774d1) |
| `discovery_k8s.access_info.connection_info.tls_info.server_name` | [discovery_k8s.access_info.connection_info.tls_info.server_name](resources--discovery--reference--group-001.md#canonical-051f1697050b89d2c4c5fa3d2eb786c820a9e1d2c9284c43661dd91553b9b977) |
| `discovery_k8s.access_info.connection_info.tls_info.trusted_ca_url` | [discovery_k8s.access_info.connection_info.tls_info.trusted_ca_url](resources--discovery--reference--group-001.md#canonical-10de23b25cd62608d4dca7be85441c6e6dbaa7d903be9feb05fed67c439277dc) |
| `discovery_k8s.access_info.isolated` | [discovery_k8s.access_info.isolated](resources--discovery--reference--group-001.md#canonical-e5c29474d502fc9203ccf2cb07a43f5b44d04ab4bb600ab683969251d3875fcb) |
| `discovery_k8s.access_info.kubeconfig_url` | [discovery_k8s.access_info.kubeconfig_url](resources--discovery--reference--group-001.md#canonical-4fa21907353dba82e6ffb770273b6b0c11e8422a324ec6638fdc6e18cb3c440f) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-d2b3d5a990b1618567f524814f1ea65f79c2ea20d3945518e1a69de552106452) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.decryption_provider` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.decryption_provider](resources--discovery--reference--group-001.md#canonical-53ce93259034988a5307c10056017a0a5d0642e0e1110cc04675079d97b6ee12) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.location` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.location](resources--discovery--reference--group-001.md#canonical-8e7665409704ceb9d3e53d1ddec136e5404c9eb0e6c3d6ae54e3f308d4de68f1) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.store_provider` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.store_provider](resources--discovery--reference--group-001.md#canonical-30db631b3093d96e1298a96e3c37042dbd470d97515ac595fe53ff4b7139a118) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-251d96014ccf55da2190ce7a5b38a6e43c329d553995bf4f7224a3789f446a71) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.provider_ref` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info.provider_ref](resources--discovery--reference--group-001.md#canonical-b9755636ebde49d3e212b16ffab2b0cf27506ba83fc6c376d1bc242eacb8f8ef) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.url` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info.url](resources--discovery--reference--group-001.md#canonical-b1df7f6cfb525b6b6965a54ba155965342686177878326d672f014d36756ce2d) |
| `discovery_k8s.access_info.reachable` | [discovery_k8s.access_info.reachable](resources--discovery--reference--group-001.md#canonical-96768b6e314b88486ec9672e5fa59a05bc737fdaec0526894a3e12a77ef82a1e) |
| `discovery_k8s.default_all` | [discovery_k8s.default_all](resources--discovery--reference--group-001.md#canonical-8f81b5cec1ea9825903f64b7364de7874545c9b701000bf194b21a2878780d97) |
| `discovery_k8s.namespace_mapping` | [discovery_k8s.namespace_mapping](resources--discovery--reference--group-001.md#canonical-5cd272e67e1ca48a55f50cc2895e2b486c2096adfaa200c0bab1c0182e569015) |
| `discovery_k8s.namespace_mapping.items` | [discovery_k8s.namespace_mapping.items](resources--discovery--reference--group-001.md#canonical-d601e1716a13960de3d216119f3aca4e6468cc8397294c29a3a63f08c9e4e43d) |
| `discovery_k8s.namespace_mapping.items.namespace` | [discovery_k8s.namespace_mapping.items.namespace](resources--discovery--reference--group-001.md#canonical-e15bc3235fcb483c52b17063c23ad1fbea5aba844d102712f1f543b19a8bd87e) |
| `discovery_k8s.namespace_mapping.items.namespace_regex` | [discovery_k8s.namespace_mapping.items.namespace_regex](resources--discovery--reference--group-001.md#canonical-5c40a396b15e1527eb0a83c3337931e0014f6248294c18869e56147ce05db464) |
| `discovery_k8s.publish_info` | [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-daaad663c60e46fc3cd5beb8a7dfb67b176f08fe573b5b483b58edc9eb35aa19) |
| `discovery_k8s.publish_info.disable_spec` | [discovery_k8s.publish_info.disable_spec](resources--discovery--reference--group-001.md#canonical-5d5855cc647e60dcf5df8d914433705ec7f6f1f07621e065ba2c54db35632d10) |
| `discovery_k8s.publish_info.dns_delegation` | [discovery_k8s.publish_info.dns_delegation](resources--discovery--reference--group-001.md#canonical-072bd0a87752ed79e4bba7e7099c3e987608768cfcaea6b1c78825a33336ab10) |
| `discovery_k8s.publish_info.dns_delegation.dns_mode` | [discovery_k8s.publish_info.dns_delegation.dns_mode](resources--discovery--reference--group-001.md#canonical-7ec04546f2e078bc15726f471ce4093cf42f922e04fc675642b5625bb80bb499) |
| `discovery_k8s.publish_info.dns_delegation.subdomain` | [discovery_k8s.publish_info.dns_delegation.subdomain](resources--discovery--reference--group-001.md#canonical-4a97df0c7e0db902ffb55ebbfbbce0c5bfadb3e4a2425d79f446fe53bf935034) |
| `discovery_k8s.publish_info.publish` | [discovery_k8s.publish_info.publish](resources--discovery--reference--group-001.md#canonical-77d16141b1689709af6a7073de8f38c71e98590c548cb4b68e1bd0200d627bce) |
| `discovery_k8s.publish_info.publish.namespace` | [discovery_k8s.publish_info.publish.namespace](resources--discovery--reference--group-001.md#canonical-b152fe2309dee67f68743414e93ce520b282612657c662e92465f60867105298) |
| `discovery_k8s.publish_info.publish_fqdns` | [discovery_k8s.publish_info.publish_fqdns](resources--discovery--reference--group-001.md#canonical-4d28bd272ede35b867c88dfe35d4fba4b2efc84c976f68c7bbe7624caa8fbe61) |
| `id` | [id](resources--discovery--reference--group-001.md#canonical-bd87c234c5f276cce3e83095107de478f0d8752ec43e47fa69b5fd7124f1aff5) |
| `labels` | [labels](resources--discovery--reference--group-001.md#canonical-a8452560bbd17f961e0a0b765b14a14be09290636d0ecb70f99275e0f5679094) |
| `name` | [name](resources--discovery--reference--group-001.md#canonical-2d4eb7f274286597736027ab5354dbe4b25815ea913f9e50cf08191a55f4b9e1) |
| `namespace` | [namespace](resources--discovery--reference--group-001.md#canonical-cd9ce63907ba47f67ff4cdbb57ea89ff1078af7f468a94c81c05324905bf0c81) |
| `no_cluster_id` | [no_cluster_id](resources--discovery--reference--group-001.md#canonical-c2ecca3a0882dca2d50b0e28ca44f30a89df074c30743106f2493b41dc069098) |
| `timeouts` | [timeouts](resources--discovery--reference--group-001.md#canonical-0644e4ff453b06517881ca38dba6745fcfd929aeec76194a3ce4b56eb687ded7) |
| `timeouts.create` | [timeouts.create](resources--discovery--reference--group-001.md#canonical-563dd2918a1e4aae4e565d0bd37b672c3e1f5d20782fbf6b9bed7fd492afe8aa) |
| `timeouts.delete` | [timeouts.delete](resources--discovery--reference--group-001.md#canonical-a288e3cab0e07d39ce0800daf5d3bbef54946e86abf14c5509156c84f3ef07e0) |
| `timeouts.read` | [timeouts.read](resources--discovery--reference--group-001.md#canonical-450211b245c01c4b5197e69906792f8b1e7fe33a88788c757cfb08a569633359) |
| `timeouts.update` | [timeouts.update](resources--discovery--reference--group-001.md#canonical-ed41c83df9d2ea1197c07f629bff553f135fc425021f1c11ff2f770122414dcb) |
| `where` | [where](resources--discovery--reference--group-001.md#canonical-794f2be13e671650d17cdb3f1fa3dfe10c0b5c91e86bebe12775b38452beafe7) |
| `where.site` | [where.site](resources--discovery--reference--group-001.md#canonical-861d0e8a4a6384e867162a660dc579581901195bddcb85b50505379ab6e6ed48) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](resources--discovery--reference--group-001.md#canonical-9c723cfcce096caaf3d494d2e8c0bf7ca8530248dc8db2bda56c720ba62d67cf) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](resources--discovery--reference--group-001.md#canonical-87b8160f328f93665f205fb27f1a9dd880f4aadc51250da52a3bdb002ce125b9) |
| `where.site.network_type` | [where.site.network_type](resources--discovery--reference--group-001.md#canonical-6b6b28d0e0f4cb37cfa3ccb39c77be984b1f905ca9035cce442924ab95baf056) |
| `where.site.ref` | [where.site.ref](resources--discovery--reference--group-001.md#canonical-8f7e8ce2c7dc0ab2b1808a5b5add9f90dc6a1d750059b9fe73b747f76234a71a) |
| `where.site.ref.kind` | [where.site.ref.kind](resources--discovery--reference--group-001.md#canonical-57731e331284b69dcc3ab8c16c1eece005a918921b7b39eae2e04935d9267992) |
| `where.site.ref.name` | [where.site.ref.name](resources--discovery--reference--group-001.md#canonical-cde27932e8699ed0e4bdf3f97b28528560425ba13cfb2d38cfa40d7651768a0c) |
| `where.site.ref.namespace` | [where.site.ref.namespace](resources--discovery--reference--group-001.md#canonical-a05b7d07a2d61ef0ff41ff9f02e38ce39b8a83bdbc37a223866cbf28d0704f32) |
| `where.site.ref.tenant` | [where.site.ref.tenant](resources--discovery--reference--group-001.md#canonical-bc6ec409c52806f1c496db5aaf05b5671574ce7e8634d819e8db6344e05315a0) |
| `where.site.ref.uid` | [where.site.ref.uid](resources--discovery--reference--group-001.md#canonical-f2b7bb256c2a65c49e9acde456968252a1593aa4996a05fe6e9bd042df71b6ad) |
| `where.virtual_network` | [where.virtual_network](resources--discovery--reference--group-002.md#canonical-2ece54f6fd934b2eafca7f5c13918d65008786bec0c2dfe98242bfddb552cc0a) |
| `where.virtual_network.ref` | [where.virtual_network.ref](resources--discovery--reference--group-002.md#canonical-3730ae9c211925657b8eb24be93addcb953c14c15ebba06a55d6a6c63ec31033) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](resources--discovery--reference--group-002.md#canonical-b4ceb9b83950c4062638d50c5c7358996cd8e49b9d3ef3986f843c86a751f328) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](resources--discovery--reference--group-002.md#canonical-519d934c44d79474813ee05c115f865045009d6208eee3dfd09f13acfe798548) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](resources--discovery--reference--group-002.md#canonical-acf92773b2f19c108dfe9a82e9806556a304918796aa2c37903e9c68fad8fe31) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](resources--discovery--reference--group-002.md#canonical-b3523db29b324793fc08601b37bc2a7e9025156297ed5df8353af5e71febe1ea) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](resources--discovery--reference--group-002.md#canonical-7c3722cc4833137ed8a79c8098258d784f03eb515d2aa1415fe620d261359dfb) |
| `where.virtual_site` | [where.virtual_site](resources--discovery--reference--group-002.md#canonical-3a7c0709bfda6a9508467b4b8cd2fa470834d012c85c9cc35844050e1408c77e) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](resources--discovery--reference--group-002.md#canonical-b645ba918516140ea91be6f1c272ceb17ec8803f729f50f5ecf785b4bce03754) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](resources--discovery--reference--group-002.md#canonical-c908a61b4a10e6895437f558ace8c3cab83cf2ffbc354952fd2c6bc757ba44e5) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](resources--discovery--reference--group-002.md#canonical-8076af3efeec3f8ea9ec03bd235873ecdad1362c924da96beae5d6b31f140e37) |
| `where.virtual_site.ref` | [where.virtual_site.ref](resources--discovery--reference--group-002.md#canonical-ab0bb56e76fbffed5041e5a906ff321ad2f815596ecd3f69acf8b6c1cd53772b) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](resources--discovery--reference--group-002.md#canonical-b4beb48ee0b624863e488aa2b99e3b4f1dc247a99ad2c812e72ec5d69bff8600) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](resources--discovery--reference--group-002.md#canonical-b9e4c4457ce1d2c73fe641ca401615f1703e6680e7eeee50248f1ec52a00d050) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](resources--discovery--reference--group-002.md#canonical-747ef68cf56ab4656d7c06e46f0c2b742e366891005620ae077be3e48285425f) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](resources--discovery--reference--group-002.md#canonical-1c1be349c43bbb104ae6aad5cff1887f77ffbdfc1d40ee238e271d8c05b6f95e) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](resources--discovery--reference--group-002.md#canonical-e5ac516569afa27073cd3e0f532acc6af811c42d11daf6088c10659e58bd3d1f) |

<a id="canonical-6e1f14cb79f89e791b9d500f97ce23f803172323f9f51c371fec0b41aaa62b1c"></a>

## Next pages — Property reference / 9d110a78f57f / 13

- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [no_cluster_id](resources--discovery--reference--group-001.md#canonical-be07760a77ec25ad8978b2bd1d1aaf9fedbe1488123a6d5e645d989c408de7d3)
- [timeouts](resources--discovery--reference--group-001.md#canonical-1f6a6e1efb3a73cacd0583b5d8d76db229ff20e709ec7d1ec796cdcf7d9d8670)
- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb9d271283f6644abf8ef5a926fd1d6f258fb61c7ce2c074f8114b83ce337a27"></a>

## discovery_consul — discovery_consul / b4104be5be8b / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- discovery_consul

<a id="canonical-4ccd1f79e3cb2f0b095dd4399f3d1634d03c05633561eaaa1f7e7e00e1f98141"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: discovery\_consul, discovery\_k8s\] Discovery configuration for Hashicorp Consul.

Upstream description:

Discovery configuration for Hashicorp Consul.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-namespace_mapping_choice": "[]"
}
```

OneOf alternatives in this subsection:

- [discovery_consul](resources--discovery--reference--group-001.md#canonical-4ccd1f79e3cb2f0b095dd4399f3d1634d03c05633561eaaa1f7e7e00e1f98141)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3339faadde78d59c922be5812f8ccbfec57ac84a5d34a03442b1c7cf749b2268)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
discovery_consul {
  # Configure direct properties listed below.
}
```

<a id="canonical-a707606af753bb96b33f571189d4bb4a38f8eea8c76181b4066a8aaa5514a354"></a>

## Direct properties — discovery_consul / b4104be5be8b / 3

- [access_info](resources--discovery--reference--group-001.md#canonical-be36a3054704f37ae8198b834da691b0ab3270b67fad51b11bca8ac58e9b46ee): complete subsection reference.

- [publish_info](resources--discovery--reference--group-001.md#canonical-d7a11fc7673d4b5e6c6a6e9ed039fc66921631d7b9231e06c6a2e2dcf27ac4c6): complete subsection reference.

<a id="canonical-eafc6935372ed7e0a3450277f8493b0ded093780eb5963d971489ba56a245121"></a>

## Next pages — discovery_consul / b4104be5be8b / 4

- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-be36a3054704f37ae8198b834da691b0ab3270b67fad51b11bca8ac58e9b46ee)
- [discovery_consul.publish_info](resources--discovery--reference--group-001.md#canonical-d7a11fc7673d4b5e6c6a6e9ed039fc66921631d7b9231e06c6a2e2dcf27ac4c6)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-be36a3054704f37ae8198b834da691b0ab3270b67fad51b11bca8ac58e9b46ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1faf04555158640f8bacb8e36f24d2f4e350295406e010dffc2e6030ecb87db0"></a>

## discovery_consul.access_info — discovery_consul.access_info / 180b3b769bd3 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- discovery_consul.access_info

<a id="canonical-fcce2b779b3834aace9ac65c315ac280a670e8bbd45e1ba8c9bf41c0b15de435"></a>

Type: `"object"`. single nested block, Optional.

Hashicorp Consul API server information.

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
access_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-87b2dfa01d4f7cb0a3d38d6005026361985df9d658a0f99b4a3d81bf92d3c228"></a>

## Direct properties — discovery_consul.access_info / 180b3b769bd3 / 3

- [connection_info](resources--discovery--reference--group-001.md#canonical-490ac711bdc48dcd23baad299856d278d6f0105629c32490ff2a2f068a44b271): complete subsection reference.

- [http_basic_auth_info](resources--discovery--reference--group-001.md#canonical-017e4ede8d98c82db5f63fc3294dc6d14941be2a12d6c604658bb84f6cb07810): complete subsection reference.

<a id="canonical-013bd4b1bbfcf8fcd81c6c1f0bf7ce474969bb33a69229b8dc18c9a2e64ab23e"></a>

## Next pages — discovery_consul.access_info / 180b3b769bd3 / 4

- [discovery_consul.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-490ac711bdc48dcd23baad299856d278d6f0105629c32490ff2a2f068a44b271)
- [discovery_consul.access_info.http_basic_auth_info](resources--discovery--reference--group-001.md#canonical-017e4ede8d98c82db5f63fc3294dc6d14941be2a12d6c604658bb84f6cb07810)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-490ac711bdc48dcd23baad299856d278d6f0105629c32490ff2a2f068a44b271"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2b806c9bc340a692d5ffff0e822010d018360ddabac076e73654891b6d12448"></a>

## discovery_consul.access_info.connection_info — discovery_consul.access_info.connection_info / 73a3da005ebf / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-be36a3054704f37ae8198b834da691b0ab3270b67fad51b11bca8ac58e9b46ee)
- discovery_consul.access_info.connection_info

<a id="canonical-887ae8127419ee0d4974cf8c87ae7aa7f645dfaa714ddadba5bc4e3a35a44244"></a>

Type: `"object"`. single nested block, Optional.

Configuration details to access discovery service REST API.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_server")}
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
connection_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3ffd91bfd4a64ddd8c878fcfe94981323ec652ac14c9cd5a53d93e0a91ac706e"></a>

## Direct properties — discovery_consul.access_info.connection_info / 73a3da005ebf / 3

<a id="canonical-5fcea4bb069eb56f5738319f93d7a6a2de599f41f2a82c58ad5276701b3bc4d4"></a>

<a id="canonical-d3fb05e5636eb653a06ddf70a98f8549d6c30d5f02082b1f345ea02633da8f4c"></a>

## api_server property — discovery_consul.access_info.connection_info / 73a3da005ebf / 4

Type: `"string"`. Optional.

API server must be a fully qualified domain string and port specified as host:port pair.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

- [tls_info](resources--discovery--reference--group-001.md#canonical-f6e26de247c26d1a361ddb3b037baf0db9b850d2c72ab0d5fcef8b33473ac704): complete subsection reference.

<a id="canonical-9db3ab7f13c99ef0f7943ad2baeccbe551238046b7d84e7e28bf2783e0df7344"></a>

## Next pages — discovery_consul.access_info.connection_info / 73a3da005ebf / 5

- [discovery_consul.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-f6e26de247c26d1a361ddb3b037baf0db9b850d2c72ab0d5fcef8b33473ac704)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-be36a3054704f37ae8198b834da691b0ab3270b67fad51b11bca8ac58e9b46ee)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-f6e26de247c26d1a361ddb3b037baf0db9b850d2c72ab0d5fcef8b33473ac704"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8346bc97467577e83cbcc755a021116b50e243c0ae0fecf66ef681f932cefa1d"></a>

## discovery_consul.access_info.connection_info.tls_info — discovery_consul.access_info.connection_info.tls_info / 1ada55ddb01a / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-be36a3054704f37ae8198b834da691b0ab3270b67fad51b11bca8ac58e9b46ee)
- [discovery_consul.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-490ac711bdc48dcd23baad299856d278d6f0105629c32490ff2a2f068a44b271)
- discovery_consul.access_info.connection_info.tls_info

<a id="canonical-d87249e9748887a627837964dae8948e90b5ad061840e163919c9e54bfcc4f74"></a>

Type: `"object"`. single nested block, Optional.

TLS config for client of discovery service.

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
tls_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-fb993070f17f5628f4a7605c997575f8f3e006ce05a39ab594987cb7e37b3598"></a>

## Direct properties — discovery_consul.access_info.connection_info.tls_info / 1ada55ddb01a / 3

<a id="canonical-b1b27ef6298b3534e4e9982ba75104c00a1b76615ee50d74ab0de3df81f23509"></a>

<a id="canonical-27451f2c8dd5da86fcfba8d94863a66e26bb18d93dcca026f59caa17ffebd6e0"></a>

## certificate property — discovery_consul.access_info.connection_info.tls_info / 1ada55ddb01a / 4

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(100, 131072),
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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](resources--discovery--reference--group-001.md#canonical-f606f8aa12a2d4bd681ccc523e20e225427b25b3a5d334da38eda5608f2d786a): complete subsection reference.

<a id="canonical-9795870f1d67f40a5a38b12beb58057ef76735d6580a202bc2ca210cea9d8b40"></a>

<a id="canonical-7608f4603b883efce70074346a11f3ac35aa33e284918e92fb6fffdabb12ffa5"></a>

## server_name property — discovery_consul.access_info.connection_info.tls_info / 1ada55ddb01a / 5

Type: `"string"`. Optional.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

Upstream description:

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-66aecc413c2263aaa2c529b60ca283f1702fbb31a73398d9e0783c5c97e6b3e0"></a>

<a id="canonical-c07e1eec6f295985c422a4938e44e3e315c908cd22a865b5a7817de7ab5349a9"></a>

## trusted_ca_url property — discovery_consul.access_info.connection_info.tls_info / 1ada55ddb01a / 6

Type: `"string"`. Optional.

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

Upstream description:

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-4a1fe57e54706055b4fde074c4f199ad1d654b3b84b6adf17d23a88526a348f2"></a>

## Next pages — discovery_consul.access_info.connection_info.tls_info / 1ada55ddb01a / 7

- [discovery_consul.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-f606f8aa12a2d4bd681ccc523e20e225427b25b3a5d334da38eda5608f2d786a)
- [discovery_consul.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-490ac711bdc48dcd23baad299856d278d6f0105629c32490ff2a2f068a44b271)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-f606f8aa12a2d4bd681ccc523e20e225427b25b3a5d334da38eda5608f2d786a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8995f956de10069319cf989436fd14180f53c66ae154d8851c605970f5f98d7"></a>

## discovery_consul.access_info.connection_info.tls_info.key_url — discovery_consul.access_info.connection_info.tls_info.key_url / cdde70e76f87 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-be36a3054704f37ae8198b834da691b0ab3270b67fad51b11bca8ac58e9b46ee)
- [discovery_consul.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-490ac711bdc48dcd23baad299856d278d6f0105629c32490ff2a2f068a44b271)
- [discovery_consul.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-f6e26de247c26d1a361ddb3b037baf0db9b850d2c72ab0d5fcef8b33473ac704)
- discovery_consul.access_info.connection_info.tls_info.key_url

<a id="canonical-3a38751d1df9fc84db01f9ff7f72051754084604881585f2196d2d1e7767f7b8"></a>

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
key_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-5673a1c436a66265c2b5a237bc67f364a7a95ae7d836031446fe67873db2ff87"></a>

## Direct properties — discovery_consul.access_info.connection_info.tls_info.key_url / cdde70e76f87 / 3

- [blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-cfd9b94b37a9b6086115ff76b0072fd617a29e8603338b6dee5e505d0ab3731a): complete subsection reference.

- [clear_secret_info](resources--discovery--reference--group-001.md#canonical-068617a197d3b41a4eaf57215580387166a6e150276d7c7bf29c294b060fa85a): complete subsection reference.

<a id="canonical-3da59157978f8a032a56b7b9ac228efbb3a180c3178804e35007260094613588"></a>

## Next pages — discovery_consul.access_info.connection_info.tls_info.key_url / cdde70e76f87 / 4

- [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-cfd9b94b37a9b6086115ff76b0072fd617a29e8603338b6dee5e505d0ab3731a)
- [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-068617a197d3b41a4eaf57215580387166a6e150276d7c7bf29c294b060fa85a)
- [discovery_consul.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-f6e26de247c26d1a361ddb3b037baf0db9b850d2c72ab0d5fcef8b33473ac704)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-cfd9b94b37a9b6086115ff76b0072fd617a29e8603338b6dee5e505d0ab3731a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2fca6d56bacb9a6881edc452bb105c07be6c18dc68b1c2d77ebd6edeae9ddfc"></a>

## discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info — discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_i / 3a672acbe850 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-be36a3054704f37ae8198b834da691b0ab3270b67fad51b11bca8ac58e9b46ee)
- [discovery_consul.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-490ac711bdc48dcd23baad299856d278d6f0105629c32490ff2a2f068a44b271)
- [discovery_consul.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-f6e26de247c26d1a361ddb3b037baf0db9b850d2c72ab0d5fcef8b33473ac704)
- [discovery_consul.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-f606f8aa12a2d4bd681ccc523e20e225427b25b3a5d334da38eda5608f2d786a)
- discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info

<a id="canonical-9fb207fa805305054b7193b5502cc2b33a8754a0275a186d6574f57331cd5a6a"></a>

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

<a id="canonical-3a5983149e339b6a0a7730f8e39aaf6b20f0616a587a97920a663dfb2aa64940"></a>

## Direct properties — discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_i / 3a672acbe850 / 3

<a id="canonical-e8a26c997ea7cf51e11b8b4a942c314796e7f6c4d52834bec001eccd913c91c6"></a>

<a id="canonical-a1177ceecc6f8f2133c976a5c88b5c1bed24edabfccdf7cb2bab554e36a736d5"></a>

## decryption_provider property — discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_i / 3a672acbe850 / 4

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

<a id="canonical-7a4cc7629c777812fbbac8c416b4c42a961a1a85549b847027403a8c28de82ea"></a>

<a id="canonical-47ad4f185023252207bc08b00de04ac758801a42bee483c4ecc53847676c6b7e"></a>

## location property — discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_i / 3a672acbe850 / 5

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

<a id="canonical-7b70c4adffebf4f9fd7c0246edb87d8e56123522ac57665e7f974def6664bf5b"></a>

<a id="canonical-d4e4e8f97aea949cb108115da79d1e4c730e25d58c944231d842940b49296ee2"></a>

## store_provider property — discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_i / 3a672acbe850 / 6

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

<a id="canonical-5a2f0443fb316a32c34689dc10827b3d74987ce3ed9a323e69bd14b22d8ccc0b"></a>

## Next pages — discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_i / 3a672acbe850 / 7

- [discovery_consul.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-f606f8aa12a2d4bd681ccc523e20e225427b25b3a5d334da38eda5608f2d786a)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-068617a197d3b41a4eaf57215580387166a6e150276d7c7bf29c294b060fa85a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35ac72fb394e6af672ce51ef4166c2432bb1dfa9427ebfc7fcdcb6656ee072d1"></a>

## discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info — discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info / aa5dc1d13eee / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-be36a3054704f37ae8198b834da691b0ab3270b67fad51b11bca8ac58e9b46ee)
- [discovery_consul.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-490ac711bdc48dcd23baad299856d278d6f0105629c32490ff2a2f068a44b271)
- [discovery_consul.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-f6e26de247c26d1a361ddb3b037baf0db9b850d2c72ab0d5fcef8b33473ac704)
- [discovery_consul.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-f606f8aa12a2d4bd681ccc523e20e225427b25b3a5d334da38eda5608f2d786a)
- discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info

<a id="canonical-f78185eeacc447937744790414924c29af90453a32c6c23b38d9091621722d58"></a>

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

<a id="canonical-3ee6adbbf5c0515de1bd4ee0fe9c3135cbdd138ae361acf026808061eebbd68f"></a>

## Direct properties — discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info / aa5dc1d13eee / 3

<a id="canonical-a43e680cc9edd11a5ccddf854dd5deefc5c32753106ad688784180696b2daada"></a>

<a id="canonical-2e06a7a91b1fff9a9a1952a1303dd7366182cb848a537075b6d043bbd4b174dc"></a>

## provider_ref property — discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info / aa5dc1d13eee / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-9cabac77c271560dead570ad236879c30b3b0ffc17a5f33d99b6250daf8678d7"></a>

<a id="canonical-fc176604a249418581787318895a294f66a6e0b14e04ffc9b4a809f4a429cdc2"></a>

## url property — discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info / aa5dc1d13eee / 5

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

<a id="canonical-0da8a902a4b7a4bdd9dff4ad9f072c8d68e0075b8105d91a48b896a8ebe8fe75"></a>

## Next pages — discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info / aa5dc1d13eee / 6

- [discovery_consul.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-f606f8aa12a2d4bd681ccc523e20e225427b25b3a5d334da38eda5608f2d786a)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-017e4ede8d98c82db5f63fc3294dc6d14941be2a12d6c604658bb84f6cb07810"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd2169720d2960dbd35a7945f76f2889b29a57394d32d7030bb53bdeaad02e18"></a>

## discovery_consul.access_info.http_basic_auth_info — discovery_consul.access_info.http_basic_auth_info / ddacbc9bcefd / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-be36a3054704f37ae8198b834da691b0ab3270b67fad51b11bca8ac58e9b46ee)
- discovery_consul.access_info.http_basic_auth_info

<a id="canonical-1bdfb91d2cc0a057d2d0fc3d515ebfbfb8bb16c384395487e8a7b5f01b495be7"></a>

Type: `"object"`. single nested block, Optional.

Authentication parameters to access Hashicorp Consul.

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
http_basic_auth_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-c312d7c7edc4de04856aa90b1f776cd2e5da7778bf97b5c38fd5dda1e8c4a330"></a>

## Direct properties — discovery_consul.access_info.http_basic_auth_info / ddacbc9bcefd / 3

- [passwd_url](resources--discovery--reference--group-001.md#canonical-5841c7d0ff8b7580a174451bf6a25ab8ec7826bb6e9615cf563a42e7eced8a68): complete subsection reference.

<a id="canonical-9e8b8abb595f3d8596bef243a890568d66577b044c81f46754f8c54ea922b433"></a>

<a id="canonical-7b58f276b2f3ed6957a9ca705c615cce6f7ea20f1c21db9cf5adb8a23e744df5"></a>

## user_name property — discovery_consul.access_info.http_basic_auth_info / ddacbc9bcefd / 4

Type: `"string"`. Optional.

User Name. Username in consul.

Upstream description:

Username in consul.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-f357e53f6eb1812e6efb072391e4e4ccf52b0b83c34334c175d66ab9c22becc1"></a>

## Next pages — discovery_consul.access_info.http_basic_auth_info / ddacbc9bcefd / 5

- [discovery_consul.access_info.http_basic_auth_info.passwd_url](resources--discovery--reference--group-001.md#canonical-5841c7d0ff8b7580a174451bf6a25ab8ec7826bb6e9615cf563a42e7eced8a68)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-be36a3054704f37ae8198b834da691b0ab3270b67fad51b11bca8ac58e9b46ee)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-5841c7d0ff8b7580a174451bf6a25ab8ec7826bb6e9615cf563a42e7eced8a68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98dc6e4bf75083c6e8ac89dcf0c492c8146c88a93696569dcdaf6209152f4b5b"></a>

## discovery_consul.access_info.http_basic_auth_info.passwd_url — discovery_consul.access_info.http_basic_auth_info.passwd_url / 21ff594ecb1a / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-be36a3054704f37ae8198b834da691b0ab3270b67fad51b11bca8ac58e9b46ee)
- [discovery_consul.access_info.http_basic_auth_info](resources--discovery--reference--group-001.md#canonical-017e4ede8d98c82db5f63fc3294dc6d14941be2a12d6c604658bb84f6cb07810)
- discovery_consul.access_info.http_basic_auth_info.passwd_url

<a id="canonical-6ec4955ba2faa36420479b136a3af64fe472ff8c13ee3a2ec87fa4c303e44fc4"></a>

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
passwd_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-4da853a7e74d9a9a6e3e8cab73b29b62a028380d24dfc275c9d267fffb9bf543"></a>

## Direct properties — discovery_consul.access_info.http_basic_auth_info.passwd_url / 21ff594ecb1a / 3

- [blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-c5b914dc90ee05313e1d00b7ae70634f9ebafa981c9ad080b50370b1c4d9fb6c): complete subsection reference.

- [clear_secret_info](resources--discovery--reference--group-001.md#canonical-e13f9ed86396391a7f7f340abde387cb2f035bdadd048e3911ae7a74dd505e8d): complete subsection reference.

<a id="canonical-555bf5cc3e9795e3771528ea6e09ecc0e4768cd77945592267e40aa1f41ed5f3"></a>

## Next pages — discovery_consul.access_info.http_basic_auth_info.passwd_url / 21ff594ecb1a / 4

- [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-c5b914dc90ee05313e1d00b7ae70634f9ebafa981c9ad080b50370b1c4d9fb6c)
- [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-e13f9ed86396391a7f7f340abde387cb2f035bdadd048e3911ae7a74dd505e8d)
- [discovery_consul.access_info.http_basic_auth_info](resources--discovery--reference--group-001.md#canonical-017e4ede8d98c82db5f63fc3294dc6d14941be2a12d6c604658bb84f6cb07810)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-c5b914dc90ee05313e1d00b7ae70634f9ebafa981c9ad080b50370b1c4d9fb6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aba3d632d1a18d908744d257ce939ac910cba66a6fd56553ba91d62b70f09147"></a>

## discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info — discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_in / 4fb7fb6322de / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-be36a3054704f37ae8198b834da691b0ab3270b67fad51b11bca8ac58e9b46ee)
- [discovery_consul.access_info.http_basic_auth_info](resources--discovery--reference--group-001.md#canonical-017e4ede8d98c82db5f63fc3294dc6d14941be2a12d6c604658bb84f6cb07810)
- [discovery_consul.access_info.http_basic_auth_info.passwd_url](resources--discovery--reference--group-001.md#canonical-5841c7d0ff8b7580a174451bf6a25ab8ec7826bb6e9615cf563a42e7eced8a68)
- discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info

<a id="canonical-85082fb61629ebc1eb811da3eb132a19a47ef42db3782a2716c82a55ef61f6d2"></a>

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

<a id="canonical-9da79600c902b9320e227caf0c43180fdd41dec554d6aae4c1d8d3b8d3898ccd"></a>

## Direct properties — discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_in / 4fb7fb6322de / 3

<a id="canonical-1e1c86773c19f42e7c93616a129ea4a09952ee1eced83943e13e1163177231f7"></a>

<a id="canonical-e77d6b2f0dcf5b00fcd1f91b08aacf6b15ad65038c0404873012297bc480e1d4"></a>

## decryption_provider property — discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_in / 4fb7fb6322de / 4

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

<a id="canonical-6d20e1a79dafa4964e57cdc861f521279fd68c70df743f510fe9966b70af0114"></a>

<a id="canonical-d4eb6b7c7b90fb45cf7f29b216f6673f954e59feeae4dd0e113e647b95b00840"></a>

## location property — discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_in / 4fb7fb6322de / 5

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

<a id="canonical-2e29c078795803aa816c34c2ba363efbf30c05f499a0751ba4b6b146211ac839"></a>

<a id="canonical-99bca87a5ae87b8813dbba1095d8a48d9c98a5fea1787de1db6ab98c3562e359"></a>

## store_provider property — discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_in / 4fb7fb6322de / 6

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

<a id="canonical-8f878f88853c0f219644ae7a8a35243d845a94a9c9165a9ab98b2ff0a0b6769f"></a>

## Next pages — discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_in / 4fb7fb6322de / 7

- [discovery_consul.access_info.http_basic_auth_info.passwd_url](resources--discovery--reference--group-001.md#canonical-5841c7d0ff8b7580a174451bf6a25ab8ec7826bb6e9615cf563a42e7eced8a68)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-e13f9ed86396391a7f7f340abde387cb2f035bdadd048e3911ae7a74dd505e8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fb849958346ee154d50af491b4e1e3c6560d4ede87e4a86e608a9e1ea47ccbe"></a>

## discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info — discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info / 5b1ac125654d / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-be36a3054704f37ae8198b834da691b0ab3270b67fad51b11bca8ac58e9b46ee)
- [discovery_consul.access_info.http_basic_auth_info](resources--discovery--reference--group-001.md#canonical-017e4ede8d98c82db5f63fc3294dc6d14941be2a12d6c604658bb84f6cb07810)
- [discovery_consul.access_info.http_basic_auth_info.passwd_url](resources--discovery--reference--group-001.md#canonical-5841c7d0ff8b7580a174451bf6a25ab8ec7826bb6e9615cf563a42e7eced8a68)
- discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info

<a id="canonical-c85330efd037032dada0ec297f3630f9d3b9d914348106bb30f6e8271dad81d3"></a>

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

<a id="canonical-a1c08862b2e1f2771fce5542f742af9ad0669bfb03af6f229ace0ccb40d3ed1c"></a>

## Direct properties — discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info / 5b1ac125654d / 3

<a id="canonical-2124f659812902bb5abc18604e41a930be2af51d1a9ff3afd3773678973afe0e"></a>

<a id="canonical-edfccea5f303c0713aef9408df9606322ff2e3880c3d85e3f2e66eebb83e4a9d"></a>

## provider_ref property — discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info / 5b1ac125654d / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-fd9b12725b67f0b52dd2a764a1d2b1473fa84e352efc9250813bca7418700974"></a>

<a id="canonical-9f76824d6d39b63b7e0122e3de7de4ca0fac3e8733eb172e4bdeefa3dd67699e"></a>

## url property — discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info / 5b1ac125654d / 5

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

<a id="canonical-ab69ddab8c3d847a3fb555fdf1abd0cfabb5e58bf24b77d5fe6afd646150c55b"></a>

## Next pages — discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info / 5b1ac125654d / 6

- [discovery_consul.access_info.http_basic_auth_info.passwd_url](resources--discovery--reference--group-001.md#canonical-5841c7d0ff8b7580a174451bf6a25ab8ec7826bb6e9615cf563a42e7eced8a68)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-d7a11fc7673d4b5e6c6a6e9ed039fc66921631d7b9231e06c6a2e2dcf27ac4c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1aa4a4375cd6954e5310a2794db6704d1a29b1758bfa44bc3377b9506144a078"></a>

## discovery_consul.publish_info — discovery_consul.publish_info / c0607cbac047 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- discovery_consul.publish_info

<a id="canonical-16674cae4d9e72cd6e857f55e9844da807c1d8dff0a7917b393db2a7bcec0ba5"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for publish info.

Upstream description:

Consul Configuration to publish VIPs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "publish")}
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
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"publish\"]"
}
```

Terraform syntax:

```terraform
publish_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-403d7f58d7fb5ecf9a87dc51e90599de2af29740c449d661b3e15987887c5636"></a>

## Direct properties — discovery_consul.publish_info / c0607cbac047 / 3

- [disable_spec](resources--discovery--reference--group-001.md#canonical-2afe929a5a6f80298b3eb0a3dd4c9597ae99c975a662aea51edd25016a71fa72): complete subsection reference.

- [publish](resources--discovery--reference--group-001.md#canonical-9e157238e1faa694bf3ae7dc465845ea807d8f3ccba116d706b3ea9beb3eccb4): complete subsection reference.

<a id="canonical-11f64a9bdbde3391f2f8e1082fce294c2d2016b32facbe74cf615b6e4c753ff8"></a>

## Next pages — discovery_consul.publish_info / c0607cbac047 / 4

- [discovery_consul.publish_info.disable_spec](resources--discovery--reference--group-001.md#canonical-2afe929a5a6f80298b3eb0a3dd4c9597ae99c975a662aea51edd25016a71fa72)
- [discovery_consul.publish_info.publish](resources--discovery--reference--group-001.md#canonical-9e157238e1faa694bf3ae7dc465845ea807d8f3ccba116d706b3ea9beb3eccb4)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-2afe929a5a6f80298b3eb0a3dd4c9597ae99c975a662aea51edd25016a71fa72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0205cb25e5b4d5effb502d7efa1edf8cab25af7185c38d0d6c2f592797309ac"></a>

## discovery_consul.publish_info.disable_spec — discovery_consul.publish_info.disable_spec / 98060e0eca55 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- [discovery_consul.publish_info](resources--discovery--reference--group-001.md#canonical-d7a11fc7673d4b5e6c6a6e9ed039fc66921631d7b9231e06c6a2e2dcf27ac4c6)
- discovery_consul.publish_info.disable_spec

<a id="canonical-e959c35d258d1d227c442e7c766da8832f0ba3545a25bf84cdb0d50e7be66019"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-fa5ea4f254c3cc017ecf20751a77f8816ed87d40088f21736509abd5b017899f"></a>

## Direct properties — discovery_consul.publish_info.disable_spec / 98060e0eca55 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e42efdd77f15a95219995cb3a3f8f5bb767f618467825eeff0ab1949d78092d6"></a>

## Next pages — discovery_consul.publish_info.disable_spec / 98060e0eca55 / 4

- [discovery_consul.publish_info](resources--discovery--reference--group-001.md#canonical-d7a11fc7673d4b5e6c6a6e9ed039fc66921631d7b9231e06c6a2e2dcf27ac4c6)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-9e157238e1faa694bf3ae7dc465845ea807d8f3ccba116d706b3ea9beb3eccb4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99f76237d2b0dd7b8ed448f0e23afcf395fbadbce25ac88539043c5271d528d3"></a>

## discovery_consul.publish_info.publish — discovery_consul.publish_info.publish / 462b3dec273f / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-c2d2ae1dcf05069f6fd7a2380945d75cb0f6828b0341da6842e506e19f47456f)
- [discovery_consul.publish_info](resources--discovery--reference--group-001.md#canonical-d7a11fc7673d4b5e6c6a6e9ed039fc66921631d7b9231e06c6a2e2dcf27ac4c6)
- discovery_consul.publish_info.publish

<a id="canonical-6132d3c55a2dbe0d39202def599f4d729e7d42e6b4f87a6ae27ab4496d3a1b7d"></a>

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
publish = {}
```

<a id="canonical-1496e9fc2d2a91a5497cfeb4aa4283a80f2b68492a0b46dd4d68f9c045093435"></a>

## Direct properties — discovery_consul.publish_info.publish / 462b3dec273f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-36aedcf4915e1b9b3aa0a3e4a3f6577d5287f4b51015cb8701be6a9f437cca51"></a>

## Next pages — discovery_consul.publish_info.publish / 462b3dec273f / 4

- [discovery_consul.publish_info](resources--discovery--reference--group-001.md#canonical-d7a11fc7673d4b5e6c6a6e9ed039fc66921631d7b9231e06c6a2e2dcf27ac4c6)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa4a3b54bfc6c817aeede4bfb192f7b36129960c3a83f9b7c682446a69a86923"></a>

## discovery_k8s — discovery_k8s / e8df71105170 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- discovery_k8s

<a id="canonical-3339faadde78d59c922be5812f8ccbfec57ac84a5d34a03442b1c7cf749b2268"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for discovery k8s.

Upstream description:

Discovery configuration for K8s.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_all",
    "namespace_mapping")}
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
  "x-ves-oneof-field-namespace_mapping_choice": "[\"default_all\",\"namespace_mapping\"]"
}
```

Terraform syntax:

```terraform
discovery_k8s {
  # Configure direct properties listed below.
}
```

<a id="canonical-78ccf582478ee443108684604147de3e301c5e1720f77bba1b617a4ee05823b3"></a>

## Direct properties — discovery_k8s / e8df71105170 / 3

- [access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8): complete subsection reference.

- [default_all](resources--discovery--reference--group-001.md#canonical-960231b1f026ee53be2d6ebbb388211509165537186a98db19e548463e7eb6e2): complete subsection reference.

- [namespace_mapping](resources--discovery--reference--group-001.md#canonical-3910a31df5831f91de2b476c6fa0971de4ff0873f9a609a39c1ba60c47254f6e): complete subsection reference.

- [publish_info](resources--discovery--reference--group-001.md#canonical-230097efea8f6b00e6211e070b14704532ce078688bff3eb5d78894b40d0c297): complete subsection reference.

<a id="canonical-95bf32a74c2ad033fbe63d009656dfaf4a939d1f5762dbeabad12a9c72e37939"></a>

## Next pages — discovery_k8s / e8df71105170 / 4

- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- [discovery_k8s.default_all](resources--discovery--reference--group-001.md#canonical-960231b1f026ee53be2d6ebbb388211509165537186a98db19e548463e7eb6e2)
- [discovery_k8s.namespace_mapping](resources--discovery--reference--group-001.md#canonical-3910a31df5831f91de2b476c6fa0971de4ff0873f9a609a39c1ba60c47254f6e)
- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-230097efea8f6b00e6211e070b14704532ce078688bff3eb5d78894b40d0c297)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-188dd58f11c2b9a8d3b543dbfe22a60a2bd0c009e37703f1fd97e0726bbbe024"></a>

## discovery_k8s.access_info — discovery_k8s.access_info / c8d2ad035406 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- discovery_k8s.access_info

<a id="canonical-1c1cd05eab3ab66a7766da3bb040696f424547d3a6e17d2d349546ac49b93408"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for access info.

Upstream description:

K8s API server access.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("connection_info",
    "kubeconfig_url"),
  validators.ConflictingObjectAttributes("isolated",
    "reachable")}
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
  "x-ves-oneof-field-config_type": "[\"connection_info\",\"kubeconfig_url\"]",
  "x-ves-oneof-field-k8s_pod_network_choice": "[\"isolated\",\"reachable\"]"
}
```

Terraform syntax:

```terraform
access_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-7b66862983ca622f95c8b11bc44a899c21aba3f0a59ead2108bb5afbccb76bd9"></a>

## Direct properties — discovery_k8s.access_info / c8d2ad035406 / 3

- [connection_info](resources--discovery--reference--group-001.md#canonical-d2d6a477c26c450b019f7bec55b5cab3fbaf71e743d718308d2bf5e9095febff): complete subsection reference.

- [isolated](resources--discovery--reference--group-001.md#canonical-fed1698ada1a80bc99a256f74fb0b4bffcd99a958e2fb70703519e1bb8e9bb05): complete subsection reference.

- [kubeconfig_url](resources--discovery--reference--group-001.md#canonical-93bcc461b58e97e4288acacee5d84ac811ced299bc50124f78b114e22504c4dc): complete subsection reference.

- [reachable](resources--discovery--reference--group-001.md#canonical-af82d846542e615028ad9c1aa9d9bcf5d2ba72a3befaaacdf1c387619f445127): complete subsection reference.

<a id="canonical-a1a716485bb1c5e09fcb1f9717e97573a831ce25fc9cef8adaf3b9981f36e272"></a>

## Next pages — discovery_k8s.access_info / c8d2ad035406 / 4

- [discovery_k8s.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-d2d6a477c26c450b019f7bec55b5cab3fbaf71e743d718308d2bf5e9095febff)
- [discovery_k8s.access_info.isolated](resources--discovery--reference--group-001.md#canonical-fed1698ada1a80bc99a256f74fb0b4bffcd99a958e2fb70703519e1bb8e9bb05)
- [discovery_k8s.access_info.kubeconfig_url](resources--discovery--reference--group-001.md#canonical-93bcc461b58e97e4288acacee5d84ac811ced299bc50124f78b114e22504c4dc)
- [discovery_k8s.access_info.reachable](resources--discovery--reference--group-001.md#canonical-af82d846542e615028ad9c1aa9d9bcf5d2ba72a3befaaacdf1c387619f445127)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-d2d6a477c26c450b019f7bec55b5cab3fbaf71e743d718308d2bf5e9095febff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4323e3f2fe47e9894b4e765ee82a49d94096ed7166cb0ba25797aefa69151d1d"></a>

## discovery_k8s.access_info.connection_info — discovery_k8s.access_info.connection_info / bf564f848b4e / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- discovery_k8s.access_info.connection_info

<a id="canonical-c5a1118e72b778babef3bad0f9979e5d573a1eb78fadb3691b908afe0397e57e"></a>

Type: `"object"`. single nested block, Optional.

Configuration details to access discovery service REST API.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_server")}
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
connection_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-893b8f90b8e415847dd1778cba459600762a7f08de93d9e34427b1a21655b1f4"></a>

## Direct properties — discovery_k8s.access_info.connection_info / bf564f848b4e / 3

<a id="canonical-20a514a1f303a368ee3abd02cc1c2e0736cd5fea6e98c2eaf4251cd80c25e3ec"></a>

<a id="canonical-c9c1fb3038414cc3c9dfdfa5faa9dd29dcc0b337a365dac33da7296b53d2b97e"></a>

## api_server property — discovery_k8s.access_info.connection_info / bf564f848b4e / 4

Type: `"string"`. Optional.

API server must be a fully qualified domain string and port specified as host:port pair.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

- [tls_info](resources--discovery--reference--group-001.md#canonical-0c0532c4d7c506423dae6f2f02cfb615cc89ce40a4990a77e87daee5505c2074): complete subsection reference.

<a id="canonical-399c439f5bc380ce110ad74b3656b6771bea9a8bdb591056a895823af4fdcca2"></a>

## Next pages — discovery_k8s.access_info.connection_info / bf564f848b4e / 5

- [discovery_k8s.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-0c0532c4d7c506423dae6f2f02cfb615cc89ce40a4990a77e87daee5505c2074)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-0c0532c4d7c506423dae6f2f02cfb615cc89ce40a4990a77e87daee5505c2074"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cbe4a9da7ba57ddfb8786981f448a376e63832cb9a07d1e4f272193b7a4255d4"></a>

## discovery_k8s.access_info.connection_info.tls_info — discovery_k8s.access_info.connection_info.tls_info / 68a5b9f122cc / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- [discovery_k8s.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-d2d6a477c26c450b019f7bec55b5cab3fbaf71e743d718308d2bf5e9095febff)
- discovery_k8s.access_info.connection_info.tls_info

<a id="canonical-cc29e4808d4a1ed2d742a13cf09d78cd51bdb081c46137b8084c3b3453d6b6ab"></a>

Type: `"object"`. single nested block, Optional.

TLS config for client of discovery service.

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
tls_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-8ec253f22c57fced8740758bcd228dd64ce6557f7cac5bce83d37f0713dbefb5"></a>

## Direct properties — discovery_k8s.access_info.connection_info.tls_info / 68a5b9f122cc / 3

<a id="canonical-d4ec998a40dc6bdbc74f5cd10a05d7991f508e8149f1956386bb8a8bbc904462"></a>

<a id="canonical-9eae298adf6d2e7f39c4120817d1f9dca9bc1d4d1241fcfa09443fddbbc68a0b"></a>

## certificate property — discovery_k8s.access_info.connection_info.tls_info / 68a5b9f122cc / 4

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(100, 131072),
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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](resources--discovery--reference--group-001.md#canonical-b93abe4605b0ef40d53375530ee882b2abe9c9eb9a917e74ce67dfe50786e815): complete subsection reference.

<a id="canonical-051f1697050b89d2c4c5fa3d2eb786c820a9e1d2c9284c43661dd91553b9b977"></a>

<a id="canonical-5ae7fb263d345ff884861403ba538cbb98d9cad5c4634b050e2483aa14e8204b"></a>

## server_name property — discovery_k8s.access_info.connection_info.tls_info / 68a5b9f122cc / 5

Type: `"string"`. Optional.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

Upstream description:

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-10de23b25cd62608d4dca7be85441c6e6dbaa7d903be9feb05fed67c439277dc"></a>

<a id="canonical-916995ee49f05b85d35bf413f8f2c73989aecc065f762bd4fa2ec2286856b16d"></a>

## trusted_ca_url property — discovery_k8s.access_info.connection_info.tls_info / 68a5b9f122cc / 6

Type: `"string"`. Optional.

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

Upstream description:

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-a1353f497dc5ac323a36f90d3122de1bc1ce3ebfcd69924838ddb44ef74300a1"></a>

## Next pages — discovery_k8s.access_info.connection_info.tls_info / 68a5b9f122cc / 7

- [discovery_k8s.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-b93abe4605b0ef40d53375530ee882b2abe9c9eb9a917e74ce67dfe50786e815)
- [discovery_k8s.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-d2d6a477c26c450b019f7bec55b5cab3fbaf71e743d718308d2bf5e9095febff)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-b93abe4605b0ef40d53375530ee882b2abe9c9eb9a917e74ce67dfe50786e815"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1e673c33a84e7ebbde927ddbeecad59f8180c102e40149e565a00313f4482d9"></a>

## discovery_k8s.access_info.connection_info.tls_info.key_url — discovery_k8s.access_info.connection_info.tls_info.key_url / b931fb9ecf70 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- [discovery_k8s.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-d2d6a477c26c450b019f7bec55b5cab3fbaf71e743d718308d2bf5e9095febff)
- [discovery_k8s.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-0c0532c4d7c506423dae6f2f02cfb615cc89ce40a4990a77e87daee5505c2074)
- discovery_k8s.access_info.connection_info.tls_info.key_url

<a id="canonical-a821f7853404febdaf408e201371a8c4ec327190ffa2ef2f9fbfd1d88389d1c8"></a>

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
key_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-e99d9b6cb0648aafbd85003144fbb763ba04bdb2978744293e68a0ea4fc47036"></a>

## Direct properties — discovery_k8s.access_info.connection_info.tls_info.key_url / b931fb9ecf70 / 3

- [blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-97d2760db1fbb8d8daa5799d99411edeb54eca91e072043afbde5a0a5c304164): complete subsection reference.

- [clear_secret_info](resources--discovery--reference--group-001.md#canonical-d8f50c5d5386e9d69127b9e7f5c331b42ddd5158ac81152c0a5230a0dec86bb2): complete subsection reference.

<a id="canonical-5157efe73b5be83f1a75ea5dfdac5b13976cb388e0a212c6e679d859d74a559b"></a>

## Next pages — discovery_k8s.access_info.connection_info.tls_info.key_url / b931fb9ecf70 / 4

- [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-97d2760db1fbb8d8daa5799d99411edeb54eca91e072043afbde5a0a5c304164)
- [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-d8f50c5d5386e9d69127b9e7f5c331b42ddd5158ac81152c0a5230a0dec86bb2)
- [discovery_k8s.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-0c0532c4d7c506423dae6f2f02cfb615cc89ce40a4990a77e87daee5505c2074)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-97d2760db1fbb8d8daa5799d99411edeb54eca91e072043afbde5a0a5c304164"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-988fe14da65af13d4d1fee2e0ed9dd474f012f6c842472860bc59ff40a49a05e"></a>

## discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info — discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info / 0f7f17580002 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- [discovery_k8s.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-d2d6a477c26c450b019f7bec55b5cab3fbaf71e743d718308d2bf5e9095febff)
- [discovery_k8s.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-0c0532c4d7c506423dae6f2f02cfb615cc89ce40a4990a77e87daee5505c2074)
- [discovery_k8s.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-b93abe4605b0ef40d53375530ee882b2abe9c9eb9a917e74ce67dfe50786e815)
- discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info

<a id="canonical-dc76b0435b05b79d77b59c6f864d1219108f5e4e4e7bfbfafbc4c6be8ab49883"></a>

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

<a id="canonical-20ccd619ce7d3739d7ebf7aac2f41ba0667bc2a20eca9e7f50a47bdff32ba9ad"></a>

## Direct properties — discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info / 0f7f17580002 / 3

<a id="canonical-36b2a4344974a47bf9aaa5de1027149982ede892717bd03bad7903319866c24e"></a>

<a id="canonical-35f56095062264a744afd34695179eda62ae3876d4b3a96521f9748e8c5b80a4"></a>

## decryption_provider property — discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info / 0f7f17580002 / 4

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

<a id="canonical-69ede6b5dc7a16fd87e1984e5720a6f8892abc35a59579e643b07702e06ec677"></a>

<a id="canonical-66dff9f1b107d2bace51d143d21ab3a945b97128c0c221f223a826d33df3f828"></a>

## location property — discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info / 0f7f17580002 / 5

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

<a id="canonical-bc243a62f03a2f6655740bfc19677e2ad5f803f7fafa6ae1493bede9ac43c262"></a>

<a id="canonical-e330f7594c244a12b33979b24c4cb7b726ac6ea3e1890c6e0cdf3c0695653035"></a>

## store_provider property — discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info / 0f7f17580002 / 6

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

<a id="canonical-00b407bb289351a1fceb68d3bc2e9e28aa5628cfd61ba537e09450da2d52698a"></a>

## Next pages — discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info / 0f7f17580002 / 7

- [discovery_k8s.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-b93abe4605b0ef40d53375530ee882b2abe9c9eb9a917e74ce67dfe50786e815)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-d8f50c5d5386e9d69127b9e7f5c331b42ddd5158ac81152c0a5230a0dec86bb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e84533869870cee0115e94f7b85a4e7ebd166b8465a477ac1fc72dbca36f9e5"></a>

## discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info — discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info / 0aec233233c5 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- [discovery_k8s.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-d2d6a477c26c450b019f7bec55b5cab3fbaf71e743d718308d2bf5e9095febff)
- [discovery_k8s.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-0c0532c4d7c506423dae6f2f02cfb615cc89ce40a4990a77e87daee5505c2074)
- [discovery_k8s.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-b93abe4605b0ef40d53375530ee882b2abe9c9eb9a917e74ce67dfe50786e815)
- discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info

<a id="canonical-195abd05a054d6752775ff2e91c341f8d271863ee322ef11d09812f207f5eed7"></a>

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

<a id="canonical-79ad597e4f9d41e1a01a58d675a77994072be5d9f5045a54be8732abe8458e1e"></a>

## Direct properties — discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info / 0aec233233c5 / 3

<a id="canonical-bde8ac1bb4e45ad7dd809d9d3c1739599971cd5b24d1a854743e3bbaf4575bce"></a>

<a id="canonical-72c9b57266a57885addbcf44d437996ae2322e7cdb36b4035cf308b14526db4e"></a>

## provider_ref property — discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info / 0aec233233c5 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-b17be070a06f89802fed2af2da64e0dc4fc19f2dc3ecc7e4b834d673d32774d1"></a>

<a id="canonical-6afc28f78d992eedc51be0c5aaa4849b3b51769aa697fdbc759c5adab75ef13f"></a>

## url property — discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info / 0aec233233c5 / 5

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

<a id="canonical-a68250174ccfb623956e619f1d42591bc26b3b6f496e0afc07330d8bff7151c5"></a>

## Next pages — discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info / 0aec233233c5 / 6

- [discovery_k8s.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-b93abe4605b0ef40d53375530ee882b2abe9c9eb9a917e74ce67dfe50786e815)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-fed1698ada1a80bc99a256f74fb0b4bffcd99a958e2fb70703519e1bb8e9bb05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b67d9dffc09e026356295e5cbf57559b5ea41fdba7d36f5e634c68456e0a05d0"></a>

## discovery_k8s.access_info.isolated — discovery_k8s.access_info.isolated / 0e0dda91b0fe / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- discovery_k8s.access_info.isolated

<a id="canonical-e5c29474d502fc9203ccf2cb07a43f5b44d04ab4bb600ab683969251d3875fcb"></a>

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
isolated = {}
```

<a id="canonical-51d8a20bab629e290397a299667e6abfb615553bc7e3731713f40d7ae5e3ded7"></a>

## Direct properties — discovery_k8s.access_info.isolated / 0e0dda91b0fe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-862ca646aa8b27426db9315017bd2d11f4ca04b8f547df999b886a012d4bc36e"></a>

## Next pages — discovery_k8s.access_info.isolated / 0e0dda91b0fe / 4

- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-93bcc461b58e97e4288acacee5d84ac811ced299bc50124f78b114e22504c4dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53812dc5855d47ad8e2b07d0e5228d1f54f37d24fa7224fa8203b9a3e2940418"></a>

## discovery_k8s.access_info.kubeconfig_url — discovery_k8s.access_info.kubeconfig_url / 78e4619d45f3 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- discovery_k8s.access_info.kubeconfig_url

<a id="canonical-4fa21907353dba82e6ffb770273b6b0c11e8422a324ec6638fdc6e18cb3c440f"></a>

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
kubeconfig_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-6dfce84178fb2065c3dd6f9c41fae51f67b77370fcd8f97c5ca304dfaec6365e"></a>

## Direct properties — discovery_k8s.access_info.kubeconfig_url / 78e4619d45f3 / 3

- [blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-59583ae895961da5f7cab99736b6875d2a6a32232db91c847328c830ffcb5c1c): complete subsection reference.

- [clear_secret_info](resources--discovery--reference--group-001.md#canonical-109e79bd83dfc1caecbb3984a105d409a94f4d4d9d4f2bcd30bb02d100f7488d): complete subsection reference.

<a id="canonical-f5c3009e9446622948eb0d59433422a983ae2a8beda98ee8b0857aa46ba65455"></a>

## Next pages — discovery_k8s.access_info.kubeconfig_url / 78e4619d45f3 / 4

- [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-59583ae895961da5f7cab99736b6875d2a6a32232db91c847328c830ffcb5c1c)
- [discovery_k8s.access_info.kubeconfig_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-109e79bd83dfc1caecbb3984a105d409a94f4d4d9d4f2bcd30bb02d100f7488d)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-59583ae895961da5f7cab99736b6875d2a6a32232db91c847328c830ffcb5c1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2bdda4bea8a90c393487e0cfe0054bdaa6300a563ef78fdfbfa5208b3e37d7cf"></a>

## discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info — discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info / 75c1e56f1bda / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- [discovery_k8s.access_info.kubeconfig_url](resources--discovery--reference--group-001.md#canonical-93bcc461b58e97e4288acacee5d84ac811ced299bc50124f78b114e22504c4dc)
- discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info

<a id="canonical-d2b3d5a990b1618567f524814f1ea65f79c2ea20d3945518e1a69de552106452"></a>

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

<a id="canonical-5c0d9550d0fa877d6d50f6c7b4ffd4438808ac83fae536d75c4af3ccaf192417"></a>

## Direct properties — discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info / 75c1e56f1bda / 3

<a id="canonical-53ce93259034988a5307c10056017a0a5d0642e0e1110cc04675079d97b6ee12"></a>

<a id="canonical-f7344e6b8982b65d027ae3915b9cf1a96fe4f0df01000ea39899540578e59dc6"></a>

## decryption_provider property — discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info / 75c1e56f1bda / 4

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

<a id="canonical-8e7665409704ceb9d3e53d1ddec136e5404c9eb0e6c3d6ae54e3f308d4de68f1"></a>

<a id="canonical-9830cfea61627f50ca6685704692be4baa18aaf7b8f77a0a2a9566bdd0693ab2"></a>

## location property — discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info / 75c1e56f1bda / 5

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

<a id="canonical-30db631b3093d96e1298a96e3c37042dbd470d97515ac595fe53ff4b7139a118"></a>

<a id="canonical-2c57e7f6766f4a78cd192b17a599f5c0e5c03967d4868caba394d01c1750b437"></a>

## store_provider property — discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info / 75c1e56f1bda / 6

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

<a id="canonical-1d558817f5c1702b9b44eff978d626e356afd81bdd0fd92e6b56678a369a507c"></a>

## Next pages — discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info / 75c1e56f1bda / 7

- [discovery_k8s.access_info.kubeconfig_url](resources--discovery--reference--group-001.md#canonical-93bcc461b58e97e4288acacee5d84ac811ced299bc50124f78b114e22504c4dc)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-109e79bd83dfc1caecbb3984a105d409a94f4d4d9d4f2bcd30bb02d100f7488d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7da2f3b6c1d27812446302bda5f4ae659a84ac554ed6cca9de89aa85250220a6"></a>

## discovery_k8s.access_info.kubeconfig_url.clear_secret_info — discovery_k8s.access_info.kubeconfig_url.clear_secret_info / 48c4195f8a76 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- [discovery_k8s.access_info.kubeconfig_url](resources--discovery--reference--group-001.md#canonical-93bcc461b58e97e4288acacee5d84ac811ced299bc50124f78b114e22504c4dc)
- discovery_k8s.access_info.kubeconfig_url.clear_secret_info

<a id="canonical-251d96014ccf55da2190ce7a5b38a6e43c329d553995bf4f7224a3789f446a71"></a>

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

<a id="canonical-9b8cf30af4245e427941895531a142fd4d8c4c04a1b48c47683b357cf490c788"></a>

## Direct properties — discovery_k8s.access_info.kubeconfig_url.clear_secret_info / 48c4195f8a76 / 3

<a id="canonical-b9755636ebde49d3e212b16ffab2b0cf27506ba83fc6c376d1bc242eacb8f8ef"></a>

<a id="canonical-714829b3da368462b5b5cd808d8a9d1fc34a4b39501061832032ab91f6e0af3c"></a>

## provider_ref property — discovery_k8s.access_info.kubeconfig_url.clear_secret_info / 48c4195f8a76 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-b1df7f6cfb525b6b6965a54ba155965342686177878326d672f014d36756ce2d"></a>

<a id="canonical-3a8152e3fcfe3316c937c587f2a64d369ea9091494c0f9f93a3f0d5ce6e1d350"></a>

## url property — discovery_k8s.access_info.kubeconfig_url.clear_secret_info / 48c4195f8a76 / 5

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

<a id="canonical-ac78979fcf94e3d5ee4a78c52c779a284467ab5ba9390e3be9744517a21727cd"></a>

## Next pages — discovery_k8s.access_info.kubeconfig_url.clear_secret_info / 48c4195f8a76 / 6

- [discovery_k8s.access_info.kubeconfig_url](resources--discovery--reference--group-001.md#canonical-93bcc461b58e97e4288acacee5d84ac811ced299bc50124f78b114e22504c4dc)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-af82d846542e615028ad9c1aa9d9bcf5d2ba72a3befaaacdf1c387619f445127"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ec12b9882a2c4511b0692e2b76b0f1e10f8cb8f2c28ee5a8192f6d555c9d79a"></a>

## discovery_k8s.access_info.reachable — discovery_k8s.access_info.reachable / 47a0c60272fb / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- discovery_k8s.access_info.reachable

<a id="canonical-96768b6e314b88486ec9672e5fa59a05bc737fdaec0526894a3e12a77ef82a1e"></a>

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
reachable = {}
```

<a id="canonical-4c3cca8e296021b1be91e1702458a3d5cd09e010ba15522af84607cd72c48fa2"></a>

## Direct properties — discovery_k8s.access_info.reachable / 47a0c60272fb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f52222e61cf6fa0362f83306927e01a351b7e3a85569ba9ebae732445749e927"></a>

## Next pages — discovery_k8s.access_info.reachable / 47a0c60272fb / 4

- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-96aefa0908dde30cd65ff7c5d0e2dfb13536ff17ff369033d0a6566a8bfac5f8)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-960231b1f026ee53be2d6ebbb388211509165537186a98db19e548463e7eb6e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-517cde8c264539185a5c109582b3bcfefd8fa328db3e88e7cdbf136b5acf6736"></a>

## discovery_k8s.default_all — discovery_k8s.default_all / b3b45f6772ca / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- discovery_k8s.default_all

<a id="canonical-8f81b5cec1ea9825903f64b7364de7874545c9b701000bf194b21a2878780d97"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default all.

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
default_all = {}
```

<a id="canonical-78c9280db2cf45e150c96971360d3a0267382e0ff4a48c6597831beee0073e30"></a>

## Direct properties — discovery_k8s.default_all / b3b45f6772ca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-21d34a935523dd71f5c8c72f62d799b739d4fba4cfccb3ebb1220edce840fb37"></a>

## Next pages — discovery_k8s.default_all / b3b45f6772ca / 4

- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-3910a31df5831f91de2b476c6fa0971de4ff0873f9a609a39c1ba60c47254f6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-027263ab7e799047b8fca0c6064c4dc8603aecc1b234e29f9a20cde5edd2bb2d"></a>

## discovery_k8s.namespace_mapping — discovery_k8s.namespace_mapping / 74101d0500a2 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- discovery_k8s.namespace_mapping

<a id="canonical-5cd272e67e1ca48a55f50cc2895e2b486c2096adfaa200c0bab1c0182e569015"></a>

Type: `"object"`. single nested block, Optional.

Select the mapping between K8s namespaces from which services will be discovered and App Namespace
to which the discovered services will be shared.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("items")}
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
namespace_mapping {
  # Configure direct properties listed below.
}
```

<a id="canonical-631b31bdc35b71b41e0767cb2093c2a2286ee686a37260a2c0301af8ace32afa"></a>

## Direct properties — discovery_k8s.namespace_mapping / 74101d0500a2 / 3

- [items](resources--discovery--reference--group-001.md#canonical-c2c853c98e58dc79d71099f96573ba9f2aea43b78fd41a16ba4b338665ed37db): complete subsection reference.

<a id="canonical-cbfcb125c4c3a6c48e89f708ad71f4e5b945d8f363e6d2f0f191226518ddba1e"></a>

## Next pages — discovery_k8s.namespace_mapping / 74101d0500a2 / 4

- [discovery_k8s.namespace_mapping.items](resources--discovery--reference--group-001.md#canonical-c2c853c98e58dc79d71099f96573ba9f2aea43b78fd41a16ba4b338665ed37db)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-c2c853c98e58dc79d71099f96573ba9f2aea43b78fd41a16ba4b338665ed37db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d41d659d4e29c5f3ef8955d41bb42f4e177c894359d9e4fb8c1acc3f034ec6b6"></a>

## discovery_k8s.namespace_mapping.items — discovery_k8s.namespace_mapping.items / 95bb4c81b997 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.namespace_mapping](resources--discovery--reference--group-001.md#canonical-3910a31df5831f91de2b476c6fa0971de4ff0873f9a609a39c1ba60c47254f6e)
- discovery_k8s.namespace_mapping.items

<a id="canonical-d601e1716a13960de3d216119f3aca4e6468cc8397294c29a3a63f08c9e4e43d"></a>

Type: `"object"`. list nested block, Optional.

Map K8s namespace(s) to App Namespaces. In Shared Configuration, Discovered Services can only be
mapped to a single App Namespace, which is determined by the first matched regex.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
items {
  # Configure direct properties listed below.
}
```

<a id="canonical-9df99ad446eabdd96455410787f05cb52449fa7fd73a50a0544c85bd99f8b5fa"></a>

## Direct properties — discovery_k8s.namespace_mapping.items / 95bb4c81b997 / 3

<a id="canonical-e15bc3235fcb483c52b17063c23ad1fbea5aba844d102712f1f543b19a8bd87e"></a>

<a id="canonical-1e39e7ae39e20c806d5ae643c8681802dbc903345cff4b394417819502039df8"></a>

## namespace property — discovery_k8s.namespace_mapping.items / 95bb4c81b997 / 4

Type: `"string"`. Optional, Computed.

F5XC Application Namespaces. Select a namespace.

Upstream description:

Select a namespace.

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

<a id="canonical-5c40a396b15e1527eb0a83c3337931e0014f6248294c18869e56147ce05db464"></a>

<a id="canonical-bf95c3906589ee87613220f9425f1b6eb648066d66947a40b3ab8546edd80bd9"></a>

## namespace_regex property — discovery_k8s.namespace_mapping.items / 95bb4c81b997 / 5

Type: `"string"`. Optional.

The regex here will be used to match K8s namespace(s).

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-6297e1532c7fed1f7754f6f38cbcca8c9c332567ca624baffc2eb95fa0314a72"></a>

## Next pages — discovery_k8s.namespace_mapping.items / 95bb4c81b997 / 6

- [discovery_k8s.namespace_mapping](resources--discovery--reference--group-001.md#canonical-3910a31df5831f91de2b476c6fa0971de4ff0873f9a609a39c1ba60c47254f6e)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-230097efea8f6b00e6211e070b14704532ce078688bff3eb5d78894b40d0c297"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac622d1a5019b299e14764f4ecc32086514b535874fa5dcaadeb2c4acc99ef56"></a>

## discovery_k8s.publish_info — discovery_k8s.publish_info / 006111c23398 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- discovery_k8s.publish_info

<a id="canonical-daaad663c60e46fc3cd5beb8a7dfb67b176f08fe573b5b483b58edc9eb35aa19"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for publish info.

Upstream description:

K8s Configuration to publish VIPs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "dns_delegation"),
  validators.ConflictingObjectAttributes("disable_spec",
    "publish"),
  validators.ConflictingObjectAttributes("disable_spec",
    "publish_fqdns"),
  validators.ConflictingObjectAttributes("dns_delegation",
    "publish"),
  validators.ConflictingObjectAttributes("dns_delegation",
    "publish_fqdns"),
  validators.ConflictingObjectAttributes("publish",
    "publish_fqdns")}
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
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"dns_delegation\",\"publish\",\"publish_fqdns\"]"
}
```

Terraform syntax:

```terraform
publish_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-ee4b224a71977dee293040a29d4b96dbc2b764b89ac78dff13ba6a23e9185018"></a>

## Direct properties — discovery_k8s.publish_info / 006111c23398 / 3

- [disable_spec](resources--discovery--reference--group-001.md#canonical-c523205fa22cca7f6d8cddea7bb21c200ff765c14b28008b57c4c78e0d749c6a): complete subsection reference.

- [dns_delegation](resources--discovery--reference--group-001.md#canonical-273efe1fb7f452ccf0a7ace6c61b337eba0d39b5ae1eb62d2b0a5c98991ac1b8): complete subsection reference.

- [publish](resources--discovery--reference--group-001.md#canonical-a17a7b4a2ccd590d1ca32b70bb31812ea4ade0858206b955e48a712d1677fc86): complete subsection reference.

- [publish_fqdns](resources--discovery--reference--group-001.md#canonical-b8c42c30127c06cf668bf07e07c88448f8b591d72e09089a9917ef7e998945d3): complete subsection reference.

<a id="canonical-af4249cbf93935a274542044dc19d746337b5d3e8f86d3ddb6bba20250adb155"></a>

## Next pages — discovery_k8s.publish_info / 006111c23398 / 4

- [discovery_k8s.publish_info.disable_spec](resources--discovery--reference--group-001.md#canonical-c523205fa22cca7f6d8cddea7bb21c200ff765c14b28008b57c4c78e0d749c6a)
- [discovery_k8s.publish_info.dns_delegation](resources--discovery--reference--group-001.md#canonical-273efe1fb7f452ccf0a7ace6c61b337eba0d39b5ae1eb62d2b0a5c98991ac1b8)
- [discovery_k8s.publish_info.publish](resources--discovery--reference--group-001.md#canonical-a17a7b4a2ccd590d1ca32b70bb31812ea4ade0858206b955e48a712d1677fc86)
- [discovery_k8s.publish_info.publish_fqdns](resources--discovery--reference--group-001.md#canonical-b8c42c30127c06cf668bf07e07c88448f8b591d72e09089a9917ef7e998945d3)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-c523205fa22cca7f6d8cddea7bb21c200ff765c14b28008b57c4c78e0d749c6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-335a78cc9cb1bdd8941d7fa3ff1c0677e42f9f67c093920b4b210a2aaa86b365"></a>

## discovery_k8s.publish_info.disable_spec — discovery_k8s.publish_info.disable_spec / 82e491ecbf89 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-230097efea8f6b00e6211e070b14704532ce078688bff3eb5d78894b40d0c297)
- discovery_k8s.publish_info.disable_spec

<a id="canonical-5d5855cc647e60dcf5df8d914433705ec7f6f1f07621e065ba2c54db35632d10"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-72f33689e55df74efca267c33c1ae5267dd43e871c87a0199f9207e8619b3dd9"></a>

## Direct properties — discovery_k8s.publish_info.disable_spec / 82e491ecbf89 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cabaf5ac1cc546e1611bea68adf3a034beb5c35729dcf3db4ea1abbd8498a29d"></a>

## Next pages — discovery_k8s.publish_info.disable_spec / 82e491ecbf89 / 4

- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-230097efea8f6b00e6211e070b14704532ce078688bff3eb5d78894b40d0c297)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-273efe1fb7f452ccf0a7ace6c61b337eba0d39b5ae1eb62d2b0a5c98991ac1b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e8e69feef15b4bfdf0603786bb1579dd90be9ec1500d0e02b4a1c2f5616887f"></a>

## discovery_k8s.publish_info.dns_delegation — discovery_k8s.publish_info.dns_delegation / 8c46b9a711a3 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-230097efea8f6b00e6211e070b14704532ce078688bff3eb5d78894b40d0c297)
- discovery_k8s.publish_info.dns_delegation

<a id="canonical-072bd0a87752ed79e4bba7e7099c3e987608768cfcaea6b1c78825a33336ab10"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dns delegation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subdomain")}
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
dns_delegation {
  # Configure direct properties listed below.
}
```

<a id="canonical-c1baa2b985f9e926f35494eabdca6181dcaf2d55c9bdcd9e358ca38044870c1e"></a>

## Direct properties — discovery_k8s.publish_info.dns_delegation / 8c46b9a711a3 / 3

<a id="canonical-7ec04546f2e078bc15726f471ce4093cf42f922e04fc675642b5625bb80bb499"></a>

<a id="canonical-d5321a92d4ccd1923c98c27280b83eee40c0d1121b29c3c61ff4790228e37b4e"></a>

## dns_mode property — discovery_k8s.publish_info.dns_delegation / 8c46b9a711a3 / 4

Type: `"string"`. Optional.

\[Enum: CORE\_DNS|KUBE\_DNS\] Two modes are possible CoreDNS: Whether external K8s cluster is
running core-DNS KubeDNS: External K8s cluster is running kube-DNS. Possible values are
\`CORE\_DNS\`, \`KUBE\_DNS\`. Defaults to \`CORE\_DNS\`.

Upstream description:

Two modes are possible

CoreDNS: Whether external K8s cluster is running core-DNS KubeDNS: External K8s cluster is running
kube-DNS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CORE_DNS",
    "KUBE_DNS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CORE_DNS",
  "enum": [
    "CORE_DNS",
    "KUBE_DNS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4a97df0c7e0db902ffb55ebbfbbce0c5bfadb3e4a2425d79f446fe53bf935034"></a>

<a id="canonical-2e3dfe7c1105ba7d13a15e277e0195ee79638e6013aca0b5934e60becd92c1dc"></a>

## subdomain property — discovery_k8s.publish_info.dns_delegation / 8c46b9a711a3 / 5

Type: `"string"`. Optional.

The DNS subdomain for which F5XC will respond to DNS queries.

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

<a id="canonical-8630412ece35ee8eb51b00091847501f109e907f76fc0df8e0d662c3d3c9fc1e"></a>

## Next pages — discovery_k8s.publish_info.dns_delegation / 8c46b9a711a3 / 6

- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-230097efea8f6b00e6211e070b14704532ce078688bff3eb5d78894b40d0c297)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-a17a7b4a2ccd590d1ca32b70bb31812ea4ade0858206b955e48a712d1677fc86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2609231e0d4ed73713cb013ff7e8e5a13912d3916c010e2a847846909acea195"></a>

## discovery_k8s.publish_info.publish — discovery_k8s.publish_info.publish / fb7541461ed4 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-230097efea8f6b00e6211e070b14704532ce078688bff3eb5d78894b40d0c297)
- discovery_k8s.publish_info.publish

<a id="canonical-77d16141b1689709af6a7073de8f38c71e98590c548cb4b68e1bd0200d627bce"></a>

Type: `"object"`. single nested block, Optional.

K8SPublishType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("namespace")}
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
publish {
  # Configure direct properties listed below.
}
```

<a id="canonical-1f69f9a8adcc315d3ae7e46c47bfe7871da861301df4d1bd6b0e02d73a049907"></a>

## Direct properties — discovery_k8s.publish_info.publish / fb7541461ed4 / 3

<a id="canonical-b152fe2309dee67f68743414e93ce520b282612657c662e92465f60867105298"></a>

<a id="canonical-a4b62e2668e9bf950f7db1d0f3a159eb581502afc8c1cc68e92f1fea2f5702f0"></a>

## namespace property — discovery_k8s.publish_info.publish / fb7541461ed4 / 4

Type: `"string"`. Optional, Computed.

The namespace where the service/endpoints need to be created if it's not included in the domain. The
external K8s administrator needs to ensure that the namespace exists.

Upstream description:

The namespace where the service/endpoints need to be created if it's not included in the domain. The
external K8s administrator needs to ensure that the namespace exists.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-fb4a41a94c500acacf53b1c3bdda59004da24226b2eab138cfa08f978d0c3376"></a>

## Next pages — discovery_k8s.publish_info.publish / fb7541461ed4 / 5

- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-230097efea8f6b00e6211e070b14704532ce078688bff3eb5d78894b40d0c297)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-b8c42c30127c06cf668bf07e07c88448f8b591d72e09089a9917ef7e998945d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b2b35d220a29a000df0372fb9eca825e7cbcf0c934f66f17547870eacfb7681"></a>

## discovery_k8s.publish_info.publish_fqdns — discovery_k8s.publish_info.publish_fqdns / 76822f4f536e / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-ca3da6b618603ba6deb0cacd1cd60eebb3c6e72569428bd03c1a923b41be8bc0)
- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-230097efea8f6b00e6211e070b14704532ce078688bff3eb5d78894b40d0c297)
- discovery_k8s.publish_info.publish_fqdns

<a id="canonical-4d28bd272ede35b867c88dfe35d4fba4b2efc84c976f68c7bbe7624caa8fbe61"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for publish fqdns.

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
publish_fqdns = {}
```

<a id="canonical-de2a3dfa3a3fc8f3525f61d49b5d707a3e7e3ebb30025f53c388cc047857e7a0"></a>

## Direct properties — discovery_k8s.publish_info.publish_fqdns / 76822f4f536e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f25c99b625cc7434f235542faccf9414eb9d971974d21a2b1976d1e6b43f7dd"></a>

## Next pages — discovery_k8s.publish_info.publish_fqdns / 76822f4f536e / 4

- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-230097efea8f6b00e6211e070b14704532ce078688bff3eb5d78894b40d0c297)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-be07760a77ec25ad8978b2bd1d1aaf9fedbe1488123a6d5e645d989c408de7d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93a5f52e897813e20808a594408f32700064426292b24a604d5ddf7b7015bedb"></a>

## no_cluster_id — no_cluster_id / 5523d6c9eae1 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- no_cluster_id

<a id="canonical-c2ecca3a0882dca2d50b0e28ca44f30a89df074c30743106f2493b41dc069098"></a>

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
no_cluster_id = {}
```

<a id="canonical-30756f1018bc5b21fd6f97846d0c50a2468603ad7621b82e665d6d016eebf5b9"></a>

## Direct properties — no_cluster_id / 5523d6c9eae1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-53ae36c7b082a1f6dca056a9c3ead052c2b27a423a358459f2e27a8e3f8ae54f"></a>

## Next pages — no_cluster_id / 5523d6c9eae1 / 4

- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-1f6a6e1efb3a73cacd0583b5d8d76db229ff20e709ec7d1ec796cdcf7d9d8670"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ead09f36baa6a91db53da647e9008ea61cc18e6ee7101f134a611996aef0768"></a>

## timeouts — timeouts / 195113ae57e8 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- timeouts

<a id="canonical-0644e4ff453b06517881ca38dba6745fcfd929aeec76194a3ce4b56eb687ded7"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-b5ba39e92b13d6475b0b1d75a3257794f08a25b3711a8478849235578ac487fc"></a>

## Direct properties — timeouts / 195113ae57e8 / 3

<a id="canonical-563dd2918a1e4aae4e565d0bd37b672c3e1f5d20782fbf6b9bed7fd492afe8aa"></a>

<a id="canonical-ea6e8934837dca8213496c9e63d5fc6588a3a4b6fc2a3307752c8891ed303001"></a>

## create property — timeouts / 195113ae57e8 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a288e3cab0e07d39ce0800daf5d3bbef54946e86abf14c5509156c84f3ef07e0"></a>

<a id="canonical-7140f51462990030c660a0a6bad4f063633534693c10b15f79c62d9fa38dcc5e"></a>

## delete property — timeouts / 195113ae57e8 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-450211b245c01c4b5197e69906792f8b1e7fe33a88788c757cfb08a569633359"></a>

<a id="canonical-e2c971477d877a0222d5c71d61157a91e75cca0bca38aefef8c9943d277f2480"></a>

## read property — timeouts / 195113ae57e8 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-ed41c83df9d2ea1197c07f629bff553f135fc425021f1c11ff2f770122414dcb"></a>

<a id="canonical-fc00fa15b389570205cf92fb479eba5b0b89f55d5112dc2699692e467f2b44ef"></a>

## update property — timeouts / 195113ae57e8 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1c156cf67a996dbf78dd10ec20355abbe53a557b218db155eda1630f50183ba9"></a>

## Next pages — timeouts / 195113ae57e8 / 8

- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e6c5b0d8cda3697b240a46e43aae60dd37e1972391d8855b2d0ae4433558844"></a>

## where — where / f43bd3ed17d3 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- where

<a id="canonical-794f2be13e671650d17cdb3f1fa3dfe10c0b5c91e86bebe12775b38452beafe7"></a>

Type: `"object"`. single nested block, Optional.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local..

Upstream description:

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingObjectAttributes("virtual_network",
    "virtual_site")}
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
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
where {
  # Configure direct properties listed below.
}
```

<a id="canonical-e2de9fb382d927de3e500b0a5218471daf6418b8a6e8d33c820ee6fe6ce20728"></a>

## Direct properties — where / f43bd3ed17d3 / 3

- [site](resources--discovery--reference--group-001.md#canonical-953b32abb079a967108f2b2a45b2fe6d53d152e0071cfdce53cbad0464d85920): complete subsection reference.

- [virtual_network](resources--discovery--reference--group-001.md#canonical-5460d6ee537cfa70bf28623e96e3dcce1005fe56ad7380a98fe2ea1c0d95b937): complete subsection reference.

- [virtual_site](resources--discovery--reference--group-002.md#canonical-abeae9515536626fbd5292145786e6d84e0d5a9b691ab8fa1c8367e9aa3f053c): complete subsection reference.

<a id="canonical-3f9878015c36e6fa8b070d5283860685edde00e322bff2b8e9204e4e281fc07a"></a>

## Next pages — where / f43bd3ed17d3 / 4

- [where.site](resources--discovery--reference--group-001.md#canonical-953b32abb079a967108f2b2a45b2fe6d53d152e0071cfdce53cbad0464d85920)
- [where.virtual_network](resources--discovery--reference--group-001.md#canonical-5460d6ee537cfa70bf28623e96e3dcce1005fe56ad7380a98fe2ea1c0d95b937)
- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-abeae9515536626fbd5292145786e6d84e0d5a9b691ab8fa1c8367e9aa3f053c)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-953b32abb079a967108f2b2a45b2fe6d53d152e0071cfdce53cbad0464d85920"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41fbc5402427417220139a31ed256c377fc084acb980bf9b4a44a0a1674134ab"></a>

## where.site — where.site / 772d5119c927 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04)
- where.site

<a id="canonical-861d0e8a4a6384e867162a660dc579581901195bddcb85b50505379ab6e6ed48"></a>

Type: `"object"`. single nested block, Optional.

Specifies a direct reference to a site configuration object.

Upstream description:

This specifies a direct reference to a site configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-011247211f61bd566dadeb2f6385e8782eb4d76e80e05d02c5df5ce560ad3b05"></a>

## Direct properties — where.site / 772d5119c927 / 3

- [disable_internet_vip](resources--discovery--reference--group-001.md#canonical-5ed1fc6b61a5144d3cbfa38c0bbeb6142ca0b874bfa25c36336d3627013bda13): complete subsection reference.

- [enable_internet_vip](resources--discovery--reference--group-001.md#canonical-94b7b4a608691b72fc7e0eeba1103c39ba9f6faed224a13978317986b31deccc): complete subsection reference.

<a id="canonical-6b6b28d0e0f4cb37cfa3ccb39c77be984b1f905ca9035cce442924ab95baf056"></a>

<a id="canonical-e311b07bc8ad20c65208afa00fbda85ec983201b0f9f922eaa758c593d5d6dd4"></a>

## network_type property — where.site / 772d5119c927 / 4

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](resources--discovery--reference--group-001.md#canonical-8e923b70430c07a7d18bc85a01427d481d92ec9bc69d202ebdcb930277f037b1): complete subsection reference.

<a id="canonical-df3c8178920a6f0540f118dbf6235c1e3adc337f2216b84fa558a75679aed3a9"></a>

## Next pages — where.site / 772d5119c927 / 5

- [where.site.disable_internet_vip](resources--discovery--reference--group-001.md#canonical-5ed1fc6b61a5144d3cbfa38c0bbeb6142ca0b874bfa25c36336d3627013bda13)
- [where.site.enable_internet_vip](resources--discovery--reference--group-001.md#canonical-94b7b4a608691b72fc7e0eeba1103c39ba9f6faed224a13978317986b31deccc)
- [where.site.ref](resources--discovery--reference--group-001.md#canonical-8e923b70430c07a7d18bc85a01427d481d92ec9bc69d202ebdcb930277f037b1)
- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-5ed1fc6b61a5144d3cbfa38c0bbeb6142ca0b874bfa25c36336d3627013bda13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9afcdd96a5969018aace7a9d9d01a8a97be3e8af1a0186a0a6dfca989b2b2515"></a>

## where.site.disable_internet_vip — where.site.disable_internet_vip / b3da253b3851 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04)
- [where.site](resources--discovery--reference--group-001.md#canonical-953b32abb079a967108f2b2a45b2fe6d53d152e0071cfdce53cbad0464d85920)
- where.site.disable_internet_vip

<a id="canonical-9c723cfcce096caaf3d494d2e8c0bf7ca8530248dc8db2bda56c720ba62d67cf"></a>

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
disable_internet_vip = {}
```

<a id="canonical-ecc8f5e677c119595a8cf1e150c1a0a5886d7e2b89da296fe621ee777fa6da19"></a>

## Direct properties — where.site.disable_internet_vip / b3da253b3851 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1c63cad5fd0ef796294ed73a398b3d975bf068eb3b5786b8ada660ef9929c18e"></a>

## Next pages — where.site.disable_internet_vip / b3da253b3851 / 4

- [where.site](resources--discovery--reference--group-001.md#canonical-953b32abb079a967108f2b2a45b2fe6d53d152e0071cfdce53cbad0464d85920)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-94b7b4a608691b72fc7e0eeba1103c39ba9f6faed224a13978317986b31deccc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc74cbccb0f9f733750f761f50ccde8b0b8442b9c321cc1291382eee26906607"></a>

## where.site.enable_internet_vip — where.site.enable_internet_vip / 2c550e2250c8 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04)
- [where.site](resources--discovery--reference--group-001.md#canonical-953b32abb079a967108f2b2a45b2fe6d53d152e0071cfdce53cbad0464d85920)
- where.site.enable_internet_vip

<a id="canonical-87b8160f328f93665f205fb27f1a9dd880f4aadc51250da52a3bdb002ce125b9"></a>

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
enable_internet_vip = {}
```

<a id="canonical-c823bb74b7da616c308019f890113c276a0f150d74a511bebeab2f6ac8d311f3"></a>

## Direct properties — where.site.enable_internet_vip / 2c550e2250c8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2b38e6f03c9cfec4c8dec90d973fe1387627db51ca223542a62b332622b3e4e4"></a>

## Next pages — where.site.enable_internet_vip / 2c550e2250c8 / 4

- [where.site](resources--discovery--reference--group-001.md#canonical-953b32abb079a967108f2b2a45b2fe6d53d152e0071cfdce53cbad0464d85920)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-8e923b70430c07a7d18bc85a01427d481d92ec9bc69d202ebdcb930277f037b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c283c183dd1f755ced8ea19f1a082e1f64c68b7bf563506a909d1fe4d8251663"></a>

## where.site.ref — where.site.ref / 13acf97b51c0 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Property reference](resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [where](resources--discovery--reference--group-001.md#canonical-4bec4bb27f77358af04cb1c0deb0c9bcfec18b0ed212b26940e8646fd0019c04)
- [where.site](resources--discovery--reference--group-001.md#canonical-953b32abb079a967108f2b2a45b2fe6d53d152e0071cfdce53cbad0464d85920)
- where.site.ref

<a id="canonical-8f7e8ce2c7dc0ab2b1808a5b5add9f90dc6a1d750059b9fe73b747f76234a71a"></a>

Type: `"object"`. list nested block, Optional.

Reference. A site direct reference.

Upstream description:

A site direct reference.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
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

<a id="canonical-ce323843f5a6badf4ae3df377e99e0975da094046a5bf54133e09cb7113cd32a"></a>

## Direct properties — where.site.ref / 13acf97b51c0 / 3

<a id="canonical-57731e331284b69dcc3ab8c16c1eece005a918921b7b39eae2e04935d9267992"></a>

<a id="canonical-f8903a4cfd6bcfc85ae4d0424db4571329c48ae37d80ab84fd83c90e4e695611"></a>

## kind property — where.site.ref / 13acf97b51c0 / 4

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

<a id="canonical-cde27932e8699ed0e4bdf3f97b28528560425ba13cfb2d38cfa40d7651768a0c"></a>

<a id="canonical-aee3397109a5ff01951cc4d74bdf0a475a37a4a988c98e12b7470d8b76539beb"></a>

## name property — where.site.ref / 13acf97b51c0 / 5

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

<a id="canonical-a05b7d07a2d61ef0ff41ff9f02e38ce39b8a83bdbc37a223866cbf28d0704f32"></a>

<a id="canonical-74fcd2a0f01bb17514d8c1c8f4bb11e4e4bd02326871e514cb27d2efca03c888"></a>

## namespace property — where.site.ref / 13acf97b51c0 / 6

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

<a id="canonical-bc6ec409c52806f1c496db5aaf05b5671574ce7e8634d819e8db6344e05315a0"></a>

<a id="canonical-06dfda8c7b813fb85d523eb1998b81b3ae3292673081de0e056f9f7eba9e915d"></a>

## tenant property — where.site.ref / 13acf97b51c0 / 7

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

<a id="canonical-f2b7bb256c2a65c49e9acde456968252a1593aa4996a05fe6e9bd042df71b6ad"></a>

<a id="canonical-55b7063b3d36e9648ae941fbcdd49e29ee1eb27a696d0b217d6c3e59139a031b"></a>

## uid property — where.site.ref / 13acf97b51c0 / 8

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

<a id="canonical-bc6b23feda451d99d01f1aee37c87b459f879f81059564efbbb7490e4be79d3f"></a>

## Next pages — where.site.ref / 13acf97b51c0 / 9

- [where.site](resources--discovery--reference--group-001.md#canonical-953b32abb079a967108f2b2a45b2fe6d53d152e0071cfdce53cbad0464d85920)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-5460d6ee537cfa70bf28623e96e3dcce1005fe56ad7380a98fe2ea1c0d95b937"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
