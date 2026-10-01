---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a978238cb1c894524f9938ead6eea510e0cc265d7b28feb1b360f5423ef0b0e"></a>

## Property reference — Property reference / 22597ce1fe7c / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- Property reference

<a id="canonical-a837bafbc34e0645075bccd8b418752956751635764debb2fcd47bcff996bae7"></a>

## Direct properties — Property reference / 22597ce1fe7c / 3

<a id="canonical-bbc2c83b4e966416749f0f1e651f91af4e039dcbe9485e71a8d9e6f5c4b95304"></a>

<a id="canonical-ff440c4d48f84393ef907395755b1695d811ab5f2b321207e8b3bc4a2191b69b"></a>

## annotations property — Property reference / 22597ce1fe7c / 4

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

- [audit_logs](resources--global_log_receiver--reference--group-001.md#canonical-ed77b985cced6a9e6caaa8ade6368f0dadca11b7d3e334c3d25cf40f5e593b46): complete subsection reference.

- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f): complete subsection reference.

- [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-31bb2a7088509e10c5d9bc0ff4558a7cd09ca1905ff420ddd31ec84beba5c7f6): complete subsection reference.

- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-d3cec99e849c282a720b0675fb155592bf8cb034e8e127ea487395b83385380d): complete subsection reference.

- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-db0eb1eadc3b3cbc91f20800c8e0c2215dcd8305de5e84d6d41a37901c086f9f): complete subsection reference.

<a id="canonical-11c05b32e69a59e59fc5aee2f9f401890d91ce1a6369dc505ece67283d4ff061"></a>

<a id="canonical-cdeb2d8246e640c6b7cc98db68d73326e6de520c014919370c55e43f04a673bc"></a>

## description property — Property reference / 22597ce1fe7c / 5

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

<a id="canonical-df337c75f59036c2cac08ed95baf6146075159454e51b0012fffd5f36261f948"></a>

<a id="canonical-ff5139d5d4f56bd70bed933108952807a5cbaacc777bae402c4863a713eb57fc"></a>

## disable property — Property reference / 22597ce1fe7c / 6

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

- [dns_logs](resources--global_log_receiver--reference--group-002.md#canonical-e5bfd6707af5d068121540614c989041b57c8f1b208696a94558f4252c6fc4b7): complete subsection reference.

- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-88f4e7ab594faa932dc1655ae2a56d356602da8ca018ed5cc5e9a181411c85cd): complete subsection reference.

- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf): complete subsection reference.

<a id="canonical-3d674a270a6f2ea3d16ddf9d08676c9b9f94cd11878cd1985247e8256964d760"></a>

<a id="canonical-6284b15e8a8b8c46d633d5a542f59a7f2a1b8f985173a28acd942499c8f74df2"></a>

## id property — Property reference / 22597ce1fe7c / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96): complete subsection reference.

<a id="canonical-88afb2ea6fd2b301691a40ccff0fd1c010f42480927d58fd89075067d33b31e6"></a>

<a id="canonical-83ab56c89139114931e9ff4552edb1adc3b0f72a32d359aa847010935b7b5d6d"></a>

## labels property — Property reference / 22597ce1fe7c / 8

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

<a id="canonical-b15819948412e70ab857be453e28ef687b0ab9cdd5233506970e6df692d51572"></a>

<a id="canonical-d26be2a45d9757d0995a88b892b1a58b736f3c031df106215f436038bf349733"></a>

## name property — Property reference / 22597ce1fe7c / 9

Type: `"string"`. Required.

Name of the Global Log Receiver. Must be unique within the namespace.

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

<a id="canonical-f60c363979d65e2049939ecca802dd23816b2f86f72f45c345a2be142476cde3"></a>

<a id="canonical-ac68cf308eb01626f4a348802c0069894e91ab17fa052e9e8b9749b2f40a27f1"></a>

## namespace property — Property reference / 22597ce1fe7c / 10

Type: `"string"`. Required.

Namespace where the Global Log Receiver is created.

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

- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-5094ca81331d2853ad9a7287be56bcb5f0fc071accea0025c24e2362b7697388): complete subsection reference.

- [ns_all](resources--global_log_receiver--reference--group-003.md#canonical-94ac97fe81951dfbc63beb0bdb84d4b3cfbea1343ac31d194e4cfe249e881911): complete subsection reference.

- [ns_current](resources--global_log_receiver--reference--group-003.md#canonical-5d4a9b8030a880ba327e5064918fb59a3565b4a15a507da773657ce9558b8eaf): complete subsection reference.

- [ns_list](resources--global_log_receiver--reference--group-003.md#canonical-7b1e2acca773fcf6415527dd41a12a76c20c635cd71ada5012b1e73e227f97e3): complete subsection reference.

- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be): complete subsection reference.

- [request_logs](resources--global_log_receiver--reference--group-004.md#canonical-6e273b67339140e7d510ed896f832f6e797c61fd0cdd910df908cae2e7cda8a5): complete subsection reference.

- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-bf56c64665781b931da6f4c5b53da70995427fed0ac9dcdc9cc103366e6dce5a): complete subsection reference.

- [security_events](resources--global_log_receiver--reference--group-004.md#canonical-e5289c97403a4948d700a4317e634154970ec974a599dc422156e5ba5ab04c9a): complete subsection reference.

- [splunk_receiver](resources--global_log_receiver--reference--group-004.md#canonical-7d67c978d6feb3f69e05312941f8bb07957fbcfa0f320a8bdc357182fe47d0dc): complete subsection reference.

- [sumo_logic_receiver](resources--global_log_receiver--reference--group-004.md#canonical-df52257dd419eb2d77d607a0f7da206f3cfaeade103803f055d0759b33652f1c): complete subsection reference.

- [timeouts](resources--global_log_receiver--reference--group-004.md#canonical-b033c7a246a27eb370c6377276894c7f70ffadc4a31122de2469c3da4f67269e): complete subsection reference.

<a id="canonical-85c84730f52926feec13e23bbb2ccc6a5165156dfa22fb069c876e1d7b7d5a12"></a>

## All schema paths — Property reference / 22597ce1fe7c / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--global_log_receiver--reference--group-001.md#canonical-bbc2c83b4e966416749f0f1e651f91af4e039dcbe9485e71a8d9e6f5c4b95304) |
| `audit_logs` | [audit_logs](resources--global_log_receiver--reference--group-001.md#canonical-11bf8c72675ef03587022c3fc826e12624118be689fffe9c02f6e9bd0bd1168f) |
| `aws_cloud_watch_receiver` | [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-0435c60661e00b6ebd7bf410341885fa5899d1ebc9d802dabf6d655c266a6312) |
| `aws_cloud_watch_receiver.aws_cred` | [aws_cloud_watch_receiver.aws_cred](resources--global_log_receiver--reference--group-001.md#canonical-ae7891b0c9b785b25c4247a92943f4cb2e3b9774a7b3ce119d34cdfb7c8e7c39) |
| `aws_cloud_watch_receiver.aws_cred.name` | [aws_cloud_watch_receiver.aws_cred.name](resources--global_log_receiver--reference--group-001.md#canonical-3900795eceef01770e02e0e186a3edf65a416adc801b9dbad01b1d7c2d087c87) |
| `aws_cloud_watch_receiver.aws_cred.namespace` | [aws_cloud_watch_receiver.aws_cred.namespace](resources--global_log_receiver--reference--group-001.md#canonical-dbccea2267252d5d0a90322a119eebae1caa213441bd19b3f165e8ec22dad98f) |
| `aws_cloud_watch_receiver.aws_cred.tenant` | [aws_cloud_watch_receiver.aws_cred.tenant](resources--global_log_receiver--reference--group-001.md#canonical-014800280c8d47b1dcf937e9b633d703c0dad2de21f23ad70ad1c8ef151ef16b) |
| `aws_cloud_watch_receiver.aws_region` | [aws_cloud_watch_receiver.aws_region](resources--global_log_receiver--reference--group-001.md#canonical-8730c9853f6863274ded542197a581fa4324f1034dc61ab04205fa48cd9d6410) |
| `aws_cloud_watch_receiver.batch` | [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-a3f62fa0fb2640bf3c5a0c7fbecc945b068e08e3bb3a48ba2a43c9834caaca9d) |
| `aws_cloud_watch_receiver.batch.max_bytes` | [aws_cloud_watch_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-001.md#canonical-be237f814c5ba1ca49e70625ef679160b630c0dab70b6c9d67565405409d61f9) |
| `aws_cloud_watch_receiver.batch.max_bytes_disabled` | [aws_cloud_watch_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-001.md#canonical-4f597c654e9c685225fd80d7a481567838d1e432eb7df9f1f405216402be89a7) |
| `aws_cloud_watch_receiver.batch.max_events` | [aws_cloud_watch_receiver.batch.max_events](resources--global_log_receiver--reference--group-001.md#canonical-0bbb094f94b284ab317460c258dcd8c9e1028cc5eceea4447bc4af61964225ec) |
| `aws_cloud_watch_receiver.batch.max_events_disabled` | [aws_cloud_watch_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-001.md#canonical-a2436b3f8467759fd63f43f83f7ecbded5ab89133471221303eaaa926704a647) |
| `aws_cloud_watch_receiver.batch.timeout_seconds` | [aws_cloud_watch_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-001.md#canonical-eacdef4b128e8d9568a5b25f7447ae907a82e79ed649d445f2c9aaa130e78d27) |
| `aws_cloud_watch_receiver.batch.timeout_seconds_default` | [aws_cloud_watch_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-001.md#canonical-262165dd3e9eccab67870cff0f5aa28964c2a0effa21c0cb4e06bc53b6ff69cf) |
| `aws_cloud_watch_receiver.compression` | [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-adce353312ca2a08817394439904e31e64a30622aa0e484b410d189fe121a684) |
| `aws_cloud_watch_receiver.compression.compression_default` | [aws_cloud_watch_receiver.compression.compression_default](resources--global_log_receiver--reference--group-001.md#canonical-30b4e1638daa7a2fcb9def861003c2e9fb0f3372a378921a46355bd9ffdf39b7) |
| `aws_cloud_watch_receiver.compression.compression_gzip` | [aws_cloud_watch_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-001.md#canonical-ecb0bb5a86aebbfccac59212bae382aab96c11a3cf33596dcad6c6cbc357ca8f) |
| `aws_cloud_watch_receiver.compression.compression_none` | [aws_cloud_watch_receiver.compression.compression_none](resources--global_log_receiver--reference--group-001.md#canonical-ada744d8fc958b8904ba0c64fd3fc6c931f4e913915f56ab047de938230c2365) |
| `aws_cloud_watch_receiver.group_name` | [aws_cloud_watch_receiver.group_name](resources--global_log_receiver--reference--group-001.md#canonical-fac31f4a10834165053e6e8eafdbe050d8ba14876851abd520304a3ab2f008e6) |
| `aws_cloud_watch_receiver.stream_name` | [aws_cloud_watch_receiver.stream_name](resources--global_log_receiver--reference--group-001.md#canonical-f3a53f2de5a2078838c99ab634f8768323f14ab46697152eed5e9be1101f7b21) |
| `azure_event_hubs_receiver` | [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-b881e14c6dfa052502f651b8a706084c45f750454aca6f7e13d12cd08d9b0602) |
| `azure_event_hubs_receiver.connection_string` | [azure_event_hubs_receiver.connection_string](resources--global_log_receiver--reference--group-001.md#canonical-3347b61e117b94dd49ae4dc57d072cd3f0d49878f500e9ccef4fb055652ed57b) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info](resources--global_log_receiver--reference--group-001.md#canonical-8169490d77b230fdebc003ed6d66f174235ad488ce1548c6ad2c6f4e2711bf0d) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.decryption_provider` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-001.md#canonical-ad4b5b90de81e1b2fe01105cc224a71f20083fd27f9748c4ae3daedb762f9313) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.location` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.location](resources--global_log_receiver--reference--group-001.md#canonical-28f5711cca0d56a29c56f4bb37b391cd4cb57971c758e4f3295c8d4b3eedfd25) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.store_provider` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-001.md#canonical-4c9b19cf0626ccd5a142df9bc7f6d946efa038017727da65d08b18d4d0d61ba7) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info` | [azure_event_hubs_receiver.connection_string.clear_secret_info](resources--global_log_receiver--reference--group-001.md#canonical-5f30419505ebc4bac30ba30846302e5b4ba5a4c093998e2813cc797488b04348) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info.provider_ref` | [azure_event_hubs_receiver.connection_string.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-001.md#canonical-f7108ad2eb564efa88b2ff162167e7b20429e27ebb9356f6b7f20818ab97b448) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info.url` | [azure_event_hubs_receiver.connection_string.clear_secret_info.url](resources--global_log_receiver--reference--group-001.md#canonical-37373ef071db4c17b283ee5005975c8d562cbdba5ee3b8d322e5f3a67eb79a73) |
| `azure_event_hubs_receiver.instance` | [azure_event_hubs_receiver.instance](resources--global_log_receiver--reference--group-001.md#canonical-cc5442fb91f8076600ab88b7f8c5b370f87c3fbc96f2c91f16f69e680a22408a) |
| `azure_event_hubs_receiver.namespace` | [azure_event_hubs_receiver.namespace](resources--global_log_receiver--reference--group-001.md#canonical-865d3df0f45c4d3abf9e45b080bfcca4e928133da851922c4658af8e4b18638a) |
| `azure_receiver` | [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-c782dfc873188e384141ccd9e5104133b8539660526f46ac90e75f3272573d6d) |
| `azure_receiver.batch` | [azure_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-0622d5592358bc82e1aebc69146a12f80fd5e811d9af21ee7aaca92a03b54fcc) |
| `azure_receiver.batch.max_bytes` | [azure_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-001.md#canonical-efc2c0d542bb962658778f2cac1600283414dd674e9e6346cec72aa54f160483) |
| `azure_receiver.batch.max_bytes_disabled` | [azure_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-002.md#canonical-a16449a9ff7068ac34ec3e69f24fd11cc168a1d66d3625b0b0ad09890bf03ccd) |
| `azure_receiver.batch.max_events` | [azure_receiver.batch.max_events](resources--global_log_receiver--reference--group-001.md#canonical-8b7e16cca09879e6daff0195a84303e60dcf0459d6c85480b62fcd782713c57d) |
| `azure_receiver.batch.max_events_disabled` | [azure_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-002.md#canonical-bd5fb777a6e125a6b9faebb9055b4e5f83be8922f6c7bf296c0a55a52a8151f4) |
| `azure_receiver.batch.timeout_seconds` | [azure_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-001.md#canonical-5f83da3dba00fb7192597acf9a42d097ffcd689ba60d3696ed9f619448625e79) |
| `azure_receiver.batch.timeout_seconds_default` | [azure_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-002.md#canonical-c4b6c998ef13987adaab1425feb501bbce353cd768b6d3b54cfd6971a47d0105) |
| `azure_receiver.compression` | [azure_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-7f04b32ecf26c206a098b2065b8fc05d35d61d2797d6c63eebca4713973b158b) |
| `azure_receiver.compression.compression_default` | [azure_receiver.compression.compression_default](resources--global_log_receiver--reference--group-002.md#canonical-8330716f4518e939791ca77bf23665ca98d6e27eac53f940bc354acdb23ef6c2) |
| `azure_receiver.compression.compression_gzip` | [azure_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-002.md#canonical-6b24279e55efc7dc3e2d16410d92796a9f4251f618ab8c8ced8f3998d70ce71c) |
| `azure_receiver.compression.compression_none` | [azure_receiver.compression.compression_none](resources--global_log_receiver--reference--group-002.md#canonical-3a17e538c83314bc30567dfb48c1f2f10e4bccc7f276953f92737c6f1dd10506) |
| `azure_receiver.connection_string` | [azure_receiver.connection_string](resources--global_log_receiver--reference--group-002.md#canonical-f41af0bf1a6169b8854d485aea2ce797685fd5e8fd625f174464711e86342c48) |
| `azure_receiver.connection_string.blindfold_secret_info` | [azure_receiver.connection_string.blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-a9fbc1e4a6cc2033cf645cf14cfe7a2990036cc443fa6afa72375ebeb67dfe68) |
| `azure_receiver.connection_string.blindfold_secret_info.decryption_provider` | [azure_receiver.connection_string.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-002.md#canonical-7e5c04a54c3a6ba2f016efe01ce2cefbcd2c8a6d0636aa913d320b6744227407) |
| `azure_receiver.connection_string.blindfold_secret_info.location` | [azure_receiver.connection_string.blindfold_secret_info.location](resources--global_log_receiver--reference--group-002.md#canonical-214ada24ec38974afc1630977bebdc11346a35b50f9320597130e661bdc7c34d) |
| `azure_receiver.connection_string.blindfold_secret_info.store_provider` | [azure_receiver.connection_string.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-002.md#canonical-e28249ec944fbff340667dbf15d3331168a6694b3163a02781c0bbd5ef977eb8) |
| `azure_receiver.connection_string.clear_secret_info` | [azure_receiver.connection_string.clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-58e5bc2b3d9e33f0d5c7565e8a9eb58c7cb7d05c1d2ef487d58601091c4cb9e8) |
| `azure_receiver.connection_string.clear_secret_info.provider_ref` | [azure_receiver.connection_string.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-002.md#canonical-ee98c7d5216818d595f643773848072ec30b1b720564d8d4768d67b2384f1d66) |
| `azure_receiver.connection_string.clear_secret_info.url` | [azure_receiver.connection_string.clear_secret_info.url](resources--global_log_receiver--reference--group-002.md#canonical-9b5ca7e2e3fbc4cd0cb9ab9ef3f066f46e49a074afdd5452b397e291a1bfa473) |
| `azure_receiver.container_name` | [azure_receiver.container_name](resources--global_log_receiver--reference--group-001.md#canonical-b34f7671b7c7395cf2c0b363e6f3c6406847ff10fd1bb759f2bf67887294b3fe) |
| `azure_receiver.filename_options` | [azure_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-352e1c515ed8ab6e9360579212169bcc9602419dd0fc7eea6d6f2dc68e1e8e95) |
| `azure_receiver.filename_options.custom_folder` | [azure_receiver.filename_options.custom_folder](resources--global_log_receiver--reference--group-002.md#canonical-a970d8bde8f99c91911016f8adf61d249b3c3ff0e5a8252aa4bf7cf3cb6ebdbc) |
| `azure_receiver.filename_options.log_type_folder` | [azure_receiver.filename_options.log_type_folder](resources--global_log_receiver--reference--group-002.md#canonical-f2bd7a700626d61adcc42a3934f71e7397b946cbf7ab134be199539c9a3d9088) |
| `azure_receiver.filename_options.no_folder` | [azure_receiver.filename_options.no_folder](resources--global_log_receiver--reference--group-002.md#canonical-c13003e98f940438651b699d0bdf784a2a5bb8c5ebf4583810629e7442b41f99) |
| `datadog_receiver` | [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-033e07e78970548fb9be688f0413eb139106cc6737071553d6e7a4e28928697e) |
| `datadog_receiver.batch` | [datadog_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-52ee87c1d221334b64615f85bb404f46a96cd8520c2c242703ac0c3e2f4d1688) |
| `datadog_receiver.batch.max_bytes` | [datadog_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-002.md#canonical-f56faf37e34ef6a21583d6f4006a26833aa4c9019a675a2e33e7c2a7de23cf21) |
| `datadog_receiver.batch.max_bytes_disabled` | [datadog_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-002.md#canonical-159ce68226b41babb547aed73f36c2cf9cc0c4946709c9b769d9c0fe7aad2ae2) |
| `datadog_receiver.batch.max_events` | [datadog_receiver.batch.max_events](resources--global_log_receiver--reference--group-002.md#canonical-91b1df73df21233a1606c684e43e5996c90a6640797841168be6d701df5617c1) |
| `datadog_receiver.batch.max_events_disabled` | [datadog_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-002.md#canonical-8034608c93c3970338eda1d0ef4bc2aeae8c5695036e684c25a81193e5d9df8c) |
| `datadog_receiver.batch.timeout_seconds` | [datadog_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-002.md#canonical-68b297395a72f3a8cfe9321f41413b2eb8f97c84401399df15316cdb601e32eb) |
| `datadog_receiver.batch.timeout_seconds_default` | [datadog_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-002.md#canonical-6a5b2396bbe465633fd794cd186e14f4d7a2dd73a1f4798cfc74e397e9c4ef64) |
| `datadog_receiver.compression` | [datadog_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-1023b7fc89b9f0ad4f399c4dbb54bdf3f6611ad611389069ed92543b92cff303) |
| `datadog_receiver.compression.compression_default` | [datadog_receiver.compression.compression_default](resources--global_log_receiver--reference--group-002.md#canonical-e077cac5a8d726fb74d208e4626646bfa26d7c5a350a049c4eede109aaf5e8c0) |
| `datadog_receiver.compression.compression_gzip` | [datadog_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-002.md#canonical-4634678b6c36998c0af582b2b6c4ded2be6631202f923e1b5b02cb218a2fa1bb) |
| `datadog_receiver.compression.compression_none` | [datadog_receiver.compression.compression_none](resources--global_log_receiver--reference--group-002.md#canonical-51785e809defc8346937776b4bdad10a2bb95fb7565afc441a425a3537c28d87) |
| `datadog_receiver.datadog_api_key` | [datadog_receiver.datadog_api_key](resources--global_log_receiver--reference--group-002.md#canonical-355010e07fefcccb2794c897f77e7554b9c57be5cb8aec3fa48bf3b12cbddaee) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info` | [datadog_receiver.datadog_api_key.blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-e6678bff85960b95e154db863b6f0ea299c0a9619df64d17eb90a8b1c94b5e23) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.decryption_provider` | [datadog_receiver.datadog_api_key.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-002.md#canonical-94a07a20b7b94309dd5247383f9c449f34327e07fdee5c43b6e6d6e1db89d238) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.location` | [datadog_receiver.datadog_api_key.blindfold_secret_info.location](resources--global_log_receiver--reference--group-002.md#canonical-c875d999e9e8f0b70ad58b736ba5ff51be271fa2bd90dfb53826be0d20188d78) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.store_provider` | [datadog_receiver.datadog_api_key.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-002.md#canonical-bf25030156eb65e82a500dcac62a8223fc605574ffc506fb849a35d10261a847) |
| `datadog_receiver.datadog_api_key.clear_secret_info` | [datadog_receiver.datadog_api_key.clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-26985fa196f62aa440b49c1d3326d70e61548872241c3b859993c7c0b3e2774f) |
| `datadog_receiver.datadog_api_key.clear_secret_info.provider_ref` | [datadog_receiver.datadog_api_key.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-002.md#canonical-a6ab129a4a0182001da6bdc11f694874325ef18e50e8f05792f17b8ffc344d2e) |
| `datadog_receiver.datadog_api_key.clear_secret_info.url` | [datadog_receiver.datadog_api_key.clear_secret_info.url](resources--global_log_receiver--reference--group-002.md#canonical-e437b791f027c0d4f0f14357f80fa96be456cf8f9f97708b4d8b14d9009d5f6a) |
| `datadog_receiver.endpoint` | [datadog_receiver.endpoint](resources--global_log_receiver--reference--group-002.md#canonical-9978a77149a296bfd2af86888e9980ae7113393390c35b49cdaf982d62429e45) |
| `datadog_receiver.no_tls` | [datadog_receiver.no_tls](resources--global_log_receiver--reference--group-002.md#canonical-bb320ed6d64990b983850e63915840f3988d3730a451bb5425c4abd6424bc240) |
| `datadog_receiver.site` | [datadog_receiver.site](resources--global_log_receiver--reference--group-002.md#canonical-eef05199ea4e0e63e5681f9c17dd54d048e6a3ca365f792eeccd908fd4d5bce0) |
| `datadog_receiver.use_tls` | [datadog_receiver.use_tls](resources--global_log_receiver--reference--group-002.md#canonical-e8f1324e24d4d4d7ccbcf32c6b13774f64c78a9103c0c341cd1bea77e290f715) |
| `datadog_receiver.use_tls.disable_verify_certificate` | [datadog_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-002.md#canonical-d7a309ab9c25311be111c5ce2eaa6a36c680bb3d0acbb038cbd3428e0cdff0f4) |
| `datadog_receiver.use_tls.disable_verify_hostname` | [datadog_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-002.md#canonical-edcce273abc11f5897eb2ffbaaf70ef7792a3318b88da89fc34c30c928892443) |
| `datadog_receiver.use_tls.enable_verify_certificate` | [datadog_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-002.md#canonical-8dd983ed6f048f8fb44df891954ae843eaef11911402713b772baba3b3620e43) |
| `datadog_receiver.use_tls.enable_verify_hostname` | [datadog_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-002.md#canonical-7ce555a34bcf49903ad0abe63b0a78f67363d74f6cd3d201987ba0f0fa5b0360) |
| `datadog_receiver.use_tls.mtls_disabled` | [datadog_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-002.md#canonical-8ce3295c2ebc66d7b25ccf83a5cdbc4e76731145e58cad3c831a85e9e468f29f) |
| `datadog_receiver.use_tls.mtls_enable` | [datadog_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-002.md#canonical-656c432bd442e3f991b5b74138b64a1c963d97209ec736202f642c7432555456) |
| `datadog_receiver.use_tls.mtls_enable.certificate` | [datadog_receiver.use_tls.mtls_enable.certificate](resources--global_log_receiver--reference--group-002.md#canonical-ed5cc5a7bc3659b3815ea4bbf723ebfe406c4a230127357ede32d0447d0812d5) |
| `datadog_receiver.use_tls.mtls_enable.key_url` | [datadog_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-002.md#canonical-03167634fac131804fb10b32b6c9a512ddaa251b4483cfeffd47ef44d4b18ec0) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-08d6288f7f6f29e41b91b6d98f6278237f8db1be87a9c1283b0555f762308033) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-002.md#canonical-3c64c386a4e59118fdcaa7a1ac2265d34345b94d4c6aef13e14ec9083742c514) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](resources--global_log_receiver--reference--group-002.md#canonical-6cfa2c7ef59285e56cf54883c8619a37132dd5eb7fa1cacf90a7a9b82b86356f) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-002.md#canonical-b7d0096c78c16383753bb7acce8afed1ce209d7a2292717ff3512db7f126307d) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-f11dc8c182beae6ad9ff05ef2a174f1855b3f6ab696e4a9f956b78bea5716509) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-002.md#canonical-d6cf2c4aa066fa5d1240d5cc57dba5e85bc250be0343e9d757a532b24201b67d) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](resources--global_log_receiver--reference--group-002.md#canonical-e1b10fb6bc6f6ddfe72a75b181f61e38db9f59e558e3df01a353518e67cbdb0e) |
| `datadog_receiver.use_tls.no_ca` | [datadog_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-002.md#canonical-6ff8d4f11ae7810fb90012c3a48ee1b9bd4a5ca5785e738af3d52236f08251e1) |
| `datadog_receiver.use_tls.trusted_ca_url` | [datadog_receiver.use_tls.trusted_ca_url](resources--global_log_receiver--reference--group-002.md#canonical-c084ae136f66f124dbab2c74000e733daeb243156c5295d0e148d88fb18baab8) |
| `description` | [description](resources--global_log_receiver--reference--group-001.md#canonical-11c05b32e69a59e59fc5aee2f9f401890d91ce1a6369dc505ece67283d4ff061) |
| `disable` | [disable](resources--global_log_receiver--reference--group-001.md#canonical-df337c75f59036c2cac08ed95baf6146075159454e51b0012fffd5f36261f948) |
| `dns_logs` | [dns_logs](resources--global_log_receiver--reference--group-002.md#canonical-9660a0fa62627fdf04ca090bee24c9b62e6e590aa3dbbaded094ca4ed6ef7acd) |
| `gcp_bucket_receiver` | [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-69146732819d2a49dcb135b7658185ce2cbee211fdc7b6292822455a6f476f55) |
| `gcp_bucket_receiver.batch` | [gcp_bucket_receiver.batch](resources--global_log_receiver--reference--group-002.md#canonical-8554be09cbc28c20aa5295a933c008a7fdf039d7b11c77313f6ee2bd828dc459) |
| `gcp_bucket_receiver.batch.max_bytes` | [gcp_bucket_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-002.md#canonical-2b1f6b609320cca7eb90d37c2f1ab10a94c26bbd7ded3637ba7278f8ce5dc712) |
| `gcp_bucket_receiver.batch.max_bytes_disabled` | [gcp_bucket_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-002.md#canonical-32e583794dc8aaf14eb55b67e84c961a047526f91f9eb82b1f9362832c661880) |
| `gcp_bucket_receiver.batch.max_events` | [gcp_bucket_receiver.batch.max_events](resources--global_log_receiver--reference--group-002.md#canonical-4cd9b7b6450841221a62b65a9a6f794e3821d1d209974c4554b40d503e1e7138) |
| `gcp_bucket_receiver.batch.max_events_disabled` | [gcp_bucket_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-002.md#canonical-2022e57657ca3f6c05bc8c50773e03874026a1639c0223bdc01d9e526bdc7b18) |
| `gcp_bucket_receiver.batch.timeout_seconds` | [gcp_bucket_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-002.md#canonical-326b9553b272d12290bdc49d84b9fc1c77f1d68442cd5b219abccf648b6573bd) |
| `gcp_bucket_receiver.batch.timeout_seconds_default` | [gcp_bucket_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-002.md#canonical-dfde10dfec28ed6ec8042e975cc45f642d84f46044c81232a61304d9bd56fd9d) |
| `gcp_bucket_receiver.bucket` | [gcp_bucket_receiver.bucket](resources--global_log_receiver--reference--group-002.md#canonical-dcb61536b7399d819f01c00fc72b54558b142e1a045a4ea9b0bc8fdca58223df) |
| `gcp_bucket_receiver.compression` | [gcp_bucket_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-57411ffecd2201661716fa992495e09c59d84af2027cac12880ed2798179b102) |
| `gcp_bucket_receiver.compression.compression_default` | [gcp_bucket_receiver.compression.compression_default](resources--global_log_receiver--reference--group-002.md#canonical-022fd83d13ec8ff0e707077d278a7a86463e5c64af0d4dc80e0ffbedb971ddc3) |
| `gcp_bucket_receiver.compression.compression_gzip` | [gcp_bucket_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-002.md#canonical-c5955efc51493c71d000ee8a354b9e13bb205b7c7e13018e1537aef315200d64) |
| `gcp_bucket_receiver.compression.compression_none` | [gcp_bucket_receiver.compression.compression_none](resources--global_log_receiver--reference--group-002.md#canonical-d6e62f7a3af29b6a4fda0236fc9e7961602de9cced9a19ae045c862d53bf6ad6) |
| `gcp_bucket_receiver.filename_options` | [gcp_bucket_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-8c7ea572e7f8d091a0e7981489856586838b661a3d26cf0c56b8e78be681bd2a) |
| `gcp_bucket_receiver.filename_options.custom_folder` | [gcp_bucket_receiver.filename_options.custom_folder](resources--global_log_receiver--reference--group-002.md#canonical-affe54b3910a53fca91242167e26ce68f39f009015f2b9e7e86dcd6116be6d78) |
| `gcp_bucket_receiver.filename_options.log_type_folder` | [gcp_bucket_receiver.filename_options.log_type_folder](resources--global_log_receiver--reference--group-002.md#canonical-26928ed492bb6f782d71bedabb5604df7b5d30cc4db91cd0d1108a8e1c575c6c) |
| `gcp_bucket_receiver.filename_options.no_folder` | [gcp_bucket_receiver.filename_options.no_folder](resources--global_log_receiver--reference--group-002.md#canonical-0253cba75462b665353e623ed72c0eebb1e52652f88d4e17d0d2007f076547e3) |
| `gcp_bucket_receiver.gcp_cred` | [gcp_bucket_receiver.gcp_cred](resources--global_log_receiver--reference--group-002.md#canonical-2caf1b9443f8c24fb1f3774974fcbcdca1773e229345f6be1f939460d4466af0) |
| `gcp_bucket_receiver.gcp_cred.name` | [gcp_bucket_receiver.gcp_cred.name](resources--global_log_receiver--reference--group-002.md#canonical-0ce457cbc0dac13f90d0dbe91eccb098f6378bd230fcbc0022848554cde6e62a) |
| `gcp_bucket_receiver.gcp_cred.namespace` | [gcp_bucket_receiver.gcp_cred.namespace](resources--global_log_receiver--reference--group-002.md#canonical-6588d6274d3f905c308559f8c3f86aaf1777ca05c392c236fa44b181e6afcfbc) |
| `gcp_bucket_receiver.gcp_cred.tenant` | [gcp_bucket_receiver.gcp_cred.tenant](resources--global_log_receiver--reference--group-002.md#canonical-405df0acdcb40d9c0c39e1d02161e2dffb37206015cd0f0629e134e8f9d4ae92) |
| `http_receiver` | [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-68e14810058d293f224852c9c18238dd7d415dd26f2dc869cc6e7dfa4c053942) |
| `http_receiver.auth_basic` | [http_receiver.auth_basic](resources--global_log_receiver--reference--group-002.md#canonical-e6d3e96aee895228cdd375eef79aa1a083e655de47b30b4f6d3e208af499ef40) |
| `http_receiver.auth_basic.password` | [http_receiver.auth_basic.password](resources--global_log_receiver--reference--group-002.md#canonical-acd56852fb35b864dd295d172c560c8012f99269cfcb5dde9301169a32b2b529) |
| `http_receiver.auth_basic.password.blindfold_secret_info` | [http_receiver.auth_basic.password.blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-890ac224e72a7879da809d3b706e70fc4db1fe0d57491e8af5e0546aa184d812) |
| `http_receiver.auth_basic.password.blindfold_secret_info.decryption_provider` | [http_receiver.auth_basic.password.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-002.md#canonical-1d478cbf8fe4c7f2806881a5e3dd20f32df96ca8620805420378f87adbf63693) |
| `http_receiver.auth_basic.password.blindfold_secret_info.location` | [http_receiver.auth_basic.password.blindfold_secret_info.location](resources--global_log_receiver--reference--group-002.md#canonical-64538a27fb865d61b44dbdede6c76237b36496b9ab8f81b80f57722d0b5b298b) |
| `http_receiver.auth_basic.password.blindfold_secret_info.store_provider` | [http_receiver.auth_basic.password.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-002.md#canonical-75755a96156f81a97d25213d4b66b2aec72efaefd2bcf8bcfe139ff5e7203ee8) |
| `http_receiver.auth_basic.password.clear_secret_info` | [http_receiver.auth_basic.password.clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-65deb3328ab7fcc737ef9f0b5fff7c13632ac1a43f80cfd15d19d22088d3a02a) |
| `http_receiver.auth_basic.password.clear_secret_info.provider_ref` | [http_receiver.auth_basic.password.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-002.md#canonical-9e57380e3ea5b53606a23310b4b043a322d63cb99e6fd5d0bee255b8b554e741) |
| `http_receiver.auth_basic.password.clear_secret_info.url` | [http_receiver.auth_basic.password.clear_secret_info.url](resources--global_log_receiver--reference--group-002.md#canonical-42467b433f4d06ad08049161330437ecda7a53bc9e9d77b9daa43dfe251a45c7) |
| `http_receiver.auth_basic.user_name` | [http_receiver.auth_basic.user_name](resources--global_log_receiver--reference--group-002.md#canonical-d1ff9a80ba6124fb7272165b9ff25f1bf4c8c34ad471a2bf2a57cc6b744ecf27) |
| `http_receiver.auth_none` | [http_receiver.auth_none](resources--global_log_receiver--reference--group-002.md#canonical-3cd5d7f9792094add0e15f21d80206e5b5c9a6731b8bc9eb518f5feabba1f56d) |
| `http_receiver.auth_token` | [http_receiver.auth_token](resources--global_log_receiver--reference--group-002.md#canonical-86fba031f98e0371f941564a8847186741669b46282b6c7fd5d28d0f1f8c6457) |
| `http_receiver.auth_token.token` | [http_receiver.auth_token.token](resources--global_log_receiver--reference--group-002.md#canonical-1db2fa5b2402817ca097a0d084aeff56456e325a7c8e8a84c65f86035048c376) |
| `http_receiver.auth_token.token.blindfold_secret_info` | [http_receiver.auth_token.token.blindfold_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-96ac986644b309a7c8dbe55941597199d337d76ec2cf57b58b31f304bf68d3c9) |
| `http_receiver.auth_token.token.blindfold_secret_info.decryption_provider` | [http_receiver.auth_token.token.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-002.md#canonical-dbf8ccfa3ff8aa41c41abb887d5b63662b58720911168ca0a6beaac276439622) |
| `http_receiver.auth_token.token.blindfold_secret_info.location` | [http_receiver.auth_token.token.blindfold_secret_info.location](resources--global_log_receiver--reference--group-002.md#canonical-f4aea66d5d7a24ac748afe5127d28ec457c37f687ad7118f423f2206af3d67fc) |
| `http_receiver.auth_token.token.blindfold_secret_info.store_provider` | [http_receiver.auth_token.token.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-002.md#canonical-81a5326e052610828d48bd9d6f6da3f98d4558baf5825544763032606b1ffc96) |
| `http_receiver.auth_token.token.clear_secret_info` | [http_receiver.auth_token.token.clear_secret_info](resources--global_log_receiver--reference--group-002.md#canonical-6c08846ce7738b68cbf2406e10d6c82ebc237c938f4478f678f1fc6ac9286b91) |
| `http_receiver.auth_token.token.clear_secret_info.provider_ref` | [http_receiver.auth_token.token.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-002.md#canonical-9cebc16a3af3093504f1ea9d2ecdd0686f1b7bf8df844bafb9fb2a5414130e90) |
| `http_receiver.auth_token.token.clear_secret_info.url` | [http_receiver.auth_token.token.clear_secret_info.url](resources--global_log_receiver--reference--group-003.md#canonical-62aa7c3f33f7b45fe8621ecb4a731227c194b331b592b672fbfa9877cf86139b) |
| `http_receiver.batch` | [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-265174117249c41d73ed244dee506d563ef19d1d06afbee6c0b10d7e52d9fe27) |
| `http_receiver.batch.max_bytes` | [http_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-003.md#canonical-4572099d4d0a9690bff63322630f1a46648a7ce10aa330559f282658253305f3) |
| `http_receiver.batch.max_bytes_disabled` | [http_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-28c46774b17b04fcb68364411a6caa387ac89f932d58558ec894d5e44c101d03) |
| `http_receiver.batch.max_events` | [http_receiver.batch.max_events](resources--global_log_receiver--reference--group-003.md#canonical-b10f7e8d121fb5732ec6f366c94dcf440c65951f2609ab22f7e84f158d8aa381) |
| `http_receiver.batch.max_events_disabled` | [http_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-34adbbd77b539b3fcc2c864b59e94ee3aa72cab02577c5122dd58a665b4b38b0) |
| `http_receiver.batch.timeout_seconds` | [http_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-003.md#canonical-d07fe985a7ac7e0b5c480c424e9b3009ecd50d0a31f1703b5c1236d83e88ce11) |
| `http_receiver.batch.timeout_seconds_default` | [http_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-030081acd975e3483ce32411bd784cc047f8060860bed0bc07d4ce8964a06562) |
| `http_receiver.compression` | [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-b83b86b3333e83201fe91a9abe5937fb4de066c4c92b32e0cf43c7ab47de1f6e) |
| `http_receiver.compression.compression_default` | [http_receiver.compression.compression_default](resources--global_log_receiver--reference--group-003.md#canonical-1f44f0023c891d79e7ef64716a3107e02446701e27abf758a47bf231cfe007ec) |
| `http_receiver.compression.compression_gzip` | [http_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-fdee09ed184e1484cc13d23af1018f5e7fc25f224026a6701ba2e9b4f9485835) |
| `http_receiver.compression.compression_none` | [http_receiver.compression.compression_none](resources--global_log_receiver--reference--group-003.md#canonical-3c57af04e0a21ac228feaabf0fd2cfe370fdcb7db8257a7ce443ad1a2d7f3a37) |
| `http_receiver.no_tls` | [http_receiver.no_tls](resources--global_log_receiver--reference--group-003.md#canonical-0913081908c9173301d9fd42da6d879fa9a9e65c38c17bd8325e931975bd5e62) |
| `http_receiver.uri` | [http_receiver.uri](resources--global_log_receiver--reference--group-002.md#canonical-4f2136021b9f318cf3f409ab7302ff6fc02ab161bf71e068f741911c9a96beb3) |
| `http_receiver.use_tls` | [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-da823b430abae4f83379d6bebbefe5b487b8bc28900ca55c0c079586ffcd6676) |
| `http_receiver.use_tls.disable_verify_certificate` | [http_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-6c1587ffa4f5b09118af3ee7cc54b100b2862d28084ae8a53298ef16de45deb7) |
| `http_receiver.use_tls.disable_verify_hostname` | [http_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-0a65a6361ac1e28abcc46d1b69b9d65155c3b79d0b73dc9ea41a0e4949c7ba92) |
| `http_receiver.use_tls.enable_verify_certificate` | [http_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-e7d6568978db5e8db1944a8e12021140e37eff94f283750f3cedf006f3312030) |
| `http_receiver.use_tls.enable_verify_hostname` | [http_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-9857d6d8423ce08b1dbfaf7c3721358d82393e4b81d46e5bd4222ed0fd8d2c2e) |
| `http_receiver.use_tls.mtls_disabled` | [http_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-003.md#canonical-fdddf224b0658e2786c4dc917ade3a70c5c9deb4e6bf0aa1e4cbb2b90d706c70) |
| `http_receiver.use_tls.mtls_enable` | [http_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-ed0f3191afc33ae9591bc93ca15ab527036207533996cabcb3e8453c0f6048d2) |
| `http_receiver.use_tls.mtls_enable.certificate` | [http_receiver.use_tls.mtls_enable.certificate](resources--global_log_receiver--reference--group-003.md#canonical-1cfe03693aa2f9c35fd567371f986983386e2651e2c178ddb48e2199195a4d1b) |
| `http_receiver.use_tls.mtls_enable.key_url` | [http_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-63d19fa58d91c691a326cd544ecfdbf9ad67f000cf8e504ed47c82a824ab8aac) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0a8d1948204027f0291065ba458ba1af149837b659057b2c7bc5e2971506543f) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-003.md#canonical-16ba461fe7c6e147eda28f6dccf5e0b3121b3ddcdda4b159f78cd2d5b072cd28) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](resources--global_log_receiver--reference--group-003.md#canonical-d8177e33c399d486470733b2da4945343f29edf28037605013cdcef709ff0391) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-003.md#canonical-5502da4075f91726723f0109815182ace83b4edb48aa7835ff4f129308b6586d) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-09d644c23bb3ab5998307af547a40f163e1f1a11822ff3bd1e26486fb4b2f98e) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-003.md#canonical-0d714e528a46c0a2af3194bafedb1b62a9b4912b7cd5743697bdbf2d7ca2242d) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](resources--global_log_receiver--reference--group-003.md#canonical-26c0bb8f9ec97e1cb3176fd691f39b204d68c74a5489a7f63f9438c7546fc2b9) |
| `http_receiver.use_tls.no_ca` | [http_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-003.md#canonical-38cc4af7d34304c708168d1b143c7eeaefdfd8de8724c6b37e2fb50c6576cf62) |
| `http_receiver.use_tls.trusted_ca_url` | [http_receiver.use_tls.trusted_ca_url](resources--global_log_receiver--reference--group-003.md#canonical-3ac6d04b08336314ffd44d1585bd836c365a6c3366623a8424b71e0e75c03baf) |
| `id` | [id](resources--global_log_receiver--reference--group-001.md#canonical-3d674a270a6f2ea3d16ddf9d08676c9b9f94cd11878cd1985247e8256964d760) |
| `kafka_receiver` | [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-30974c0ff3bb7bbe13c8e4b6fdd227339653da44184623051eff1f0043a6b046) |
| `kafka_receiver.batch` | [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-963a6ba94cfc28a6f2efb026eb35c4f2385ae27c8d8b109f8b1da9b3a37e6c54) |
| `kafka_receiver.batch.max_bytes` | [kafka_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-003.md#canonical-117dc8605914da4820916b12d4dbfbdca0c72d7f83ac4297177926d989ddbc99) |
| `kafka_receiver.batch.max_bytes_disabled` | [kafka_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-56a6748d774bd3896fd9a9e1dc32aff3ade779c4f52e8406bfdf9a295e411329) |
| `kafka_receiver.batch.max_events` | [kafka_receiver.batch.max_events](resources--global_log_receiver--reference--group-003.md#canonical-04f485a24ee8e2ab64a7c8e301d3b0b5c884bd40da07168612adb464a79bc365) |
| `kafka_receiver.batch.max_events_disabled` | [kafka_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-2afb15967b8cc8a8d31ca3e3a3e20fcb829ed19513cd9b0c8c2751ed5817f4a7) |
| `kafka_receiver.batch.timeout_seconds` | [kafka_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-003.md#canonical-47945ad1af0cba0139bd6b85c370e983ce9ca663a6ee07e1775664bd54b87fe0) |
| `kafka_receiver.batch.timeout_seconds_default` | [kafka_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-9ff7d9d5375134e1f689735cc4d40e2df34fc4ee3c2a5e624f89bcf6e200e7a2) |
| `kafka_receiver.bootstrap_servers` | [kafka_receiver.bootstrap_servers](resources--global_log_receiver--reference--group-003.md#canonical-6991fcf1dde293e7176fc9d2cb81576fd904e807fe54ab5a66d6a8a3e3e59c7b) |
| `kafka_receiver.compression` | [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-9deaaf69949c258db0e01ea4908575aec8d42b0f3d575cfd92b2c9feeeeea3d0) |
| `kafka_receiver.compression.compression_default` | [kafka_receiver.compression.compression_default](resources--global_log_receiver--reference--group-003.md#canonical-4752d790f746f1f24332c0dab8615b2bc5e79a942dd9072a72e114f6a02790d7) |
| `kafka_receiver.compression.compression_gzip` | [kafka_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-a4804e8fcd71cff07724ba34ecab165e8283dfe6669061d1befd2783dd42d4df) |
| `kafka_receiver.compression.compression_none` | [kafka_receiver.compression.compression_none](resources--global_log_receiver--reference--group-003.md#canonical-f0c69de074e62caade164d81978666e1310c86b93b6de25f93f22941afc4bccd) |
| `kafka_receiver.kafka_topic` | [kafka_receiver.kafka_topic](resources--global_log_receiver--reference--group-003.md#canonical-2282ccbe31771db16be8ba1be3539e6545bda85edf555c2c67064a73b55af75c) |
| `kafka_receiver.no_tls` | [kafka_receiver.no_tls](resources--global_log_receiver--reference--group-003.md#canonical-03ff61fef6f9738b90ff4e2da618cf6a3e6c7601c5c2e15af77e1b2244aea015) |
| `kafka_receiver.use_tls` | [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-d9f92407f496c0c5f4f7581665895ec6aa11e6a1cb7cc5052d3626097f9dbdcd) |
| `kafka_receiver.use_tls.disable_verify_certificate` | [kafka_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-622d450ee9d37690261226de1a1bdf9a24326629f8cb8d9e62df81e0ebe67fc2) |
| `kafka_receiver.use_tls.disable_verify_hostname` | [kafka_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-9f9e6f5c59b29ca295414a6aa0be43e667e62584defa1ff180441a3b0bf8a94b) |
| `kafka_receiver.use_tls.enable_verify_certificate` | [kafka_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-d55a7883cfc079b353d05d171a7323cc1b55a1e9fe6182fa748ad7c6991bc1e4) |
| `kafka_receiver.use_tls.enable_verify_hostname` | [kafka_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-519dcb2d3132d21c7c140d13b2be75440fd7cf282156788ed20ceb1bdb66ecd8) |
| `kafka_receiver.use_tls.mtls_disabled` | [kafka_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-003.md#canonical-5c609aa0e27278c5165ab781c2e0fe6a8ed7497c44b4f05809e2f31fd5b2a604) |
| `kafka_receiver.use_tls.mtls_enable` | [kafka_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-5d70aed363e7d721744ff9e3c22126058109a9652fdc783c2b2fcec0d78070ce) |
| `kafka_receiver.use_tls.mtls_enable.certificate` | [kafka_receiver.use_tls.mtls_enable.certificate](resources--global_log_receiver--reference--group-003.md#canonical-14b58c60cd8b5965d94c91762fde4dae3997d8423ea30d7b391018f7f55c3ad6) |
| `kafka_receiver.use_tls.mtls_enable.key_url` | [kafka_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-f08e1c06063c1b1ab3116aa5ec243cf853f0c03763742e7b7a4d2acf72cbf1ff) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-3c8bdd47a90070b84b0bfe7fa2f4c4d8812185dbf6d5010cbd506ad6dbd77d43) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-003.md#canonical-33468ce842df2c2914005d30288a444846c29af8835b32584aa816aad9721666) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](resources--global_log_receiver--reference--group-003.md#canonical-7fce96883dd3ad4054f6bdb5d20c8da8f206100af2b2b6fb0ecb8464fb83b713) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-003.md#canonical-d53af5aa8fc959a377913e83e7dd6b44293074fa0f52d1468428a81d8ee96126) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-90327d476ce0cc58b3eec63d068a635e5a518554a4b9a538405b2a2fe7e81b38) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-003.md#canonical-0c5b98a9e22e9df1a985ea247ab515e31eaecc9b44da05c8ae6b9e8ff2eaff02) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](resources--global_log_receiver--reference--group-003.md#canonical-02022316140abddc673f2ef90fee5a1adc99be8b0555c57c63b490140f5e8f9a) |
| `kafka_receiver.use_tls.no_ca` | [kafka_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-003.md#canonical-35669eabd41c9d9492efedafc66c56f448f71b5c604d81f400bc99ee1a0aa8fd) |
| `kafka_receiver.use_tls.trusted_ca_url` | [kafka_receiver.use_tls.trusted_ca_url](resources--global_log_receiver--reference--group-003.md#canonical-56598bd4f356d247a04ae1fe4ed15df04e5a2b24858f427d44ef586123a01d17) |
| `labels` | [labels](resources--global_log_receiver--reference--group-001.md#canonical-88afb2ea6fd2b301691a40ccff0fd1c010f42480927d58fd89075067d33b31e6) |
| `name` | [name](resources--global_log_receiver--reference--group-001.md#canonical-b15819948412e70ab857be453e28ef687b0ab9cdd5233506970e6df692d51572) |
| `namespace` | [namespace](resources--global_log_receiver--reference--group-001.md#canonical-f60c363979d65e2049939ecca802dd23816b2f86f72f45c345a2be142476cde3) |
| `new_relic_receiver` | [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-6c19c494b563c2d60abdc9410439454b737deb2ef492e68a3bff9eec490830d5) |
| `new_relic_receiver.api_key` | [new_relic_receiver.api_key](resources--global_log_receiver--reference--group-003.md#canonical-2b264cb4868e53777a6ab1fcffb99bb127045f1d6a6e14775fb0295454eb0ca2) |
| `new_relic_receiver.api_key.blindfold_secret_info` | [new_relic_receiver.api_key.blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-8e6141350e7bf8fa5e5a2f7aa87012fee73aa3532be69ba70f96a9f0f381be78) |
| `new_relic_receiver.api_key.blindfold_secret_info.decryption_provider` | [new_relic_receiver.api_key.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-003.md#canonical-7436a86f5294be1e619d49f552098e3e50cf8e7e83ca35b468ed589b8a12d9ac) |
| `new_relic_receiver.api_key.blindfold_secret_info.location` | [new_relic_receiver.api_key.blindfold_secret_info.location](resources--global_log_receiver--reference--group-003.md#canonical-ab8178a4dcf5ddef2b7fa41f201ede4e2dbad8f5ec16c4c5711e3615b9ec1763) |
| `new_relic_receiver.api_key.blindfold_secret_info.store_provider` | [new_relic_receiver.api_key.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-003.md#canonical-5f5a9883fa7bfdee2404edc2844d44985b5e52d908d339340c74d5222e2f1597) |
| `new_relic_receiver.api_key.clear_secret_info` | [new_relic_receiver.api_key.clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0cd49d3c1d5e41e76e12a6715ae208b036a01bbcd0ca577117291011c5a0336e) |
| `new_relic_receiver.api_key.clear_secret_info.provider_ref` | [new_relic_receiver.api_key.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-003.md#canonical-002a0657116d25c20cf59dda534f57e2449d103079e5e75d269dd98fd2620cca) |
| `new_relic_receiver.api_key.clear_secret_info.url` | [new_relic_receiver.api_key.clear_secret_info.url](resources--global_log_receiver--reference--group-003.md#canonical-9dfd0bb6291cdb0c350756905f5dcd0726d3558d4d763e31d127db1eeaff583e) |
| `new_relic_receiver.eu` | [new_relic_receiver.eu](resources--global_log_receiver--reference--group-003.md#canonical-92b8eb502fc087745662d2f24d107c227c7a8376c7e3024df2db8709854c28ad) |
| `new_relic_receiver.us` | [new_relic_receiver.us](resources--global_log_receiver--reference--group-003.md#canonical-4351f3cb11d29176bb6e65617ebab6c3c341f487c0aceebb2576538ef01b94e0) |
| `ns_all` | [ns_all](resources--global_log_receiver--reference--group-003.md#canonical-c27f9e5c2cba312b5473ae4b3229b22d564f8299e9f7c4e79e35919e2517a660) |
| `ns_current` | [ns_current](resources--global_log_receiver--reference--group-003.md#canonical-31ea42d689d2504cf6aa86775e94506b829f188ffb8cb066ef06b562e932d6c9) |
| `ns_list` | [ns_list](resources--global_log_receiver--reference--group-003.md#canonical-3d8776edceb7fd01ca4ca000da91ef806f6e280746404127792a14e1d813cd90) |
| `ns_list.namespaces` | [ns_list.namespaces](resources--global_log_receiver--reference--group-003.md#canonical-f57c3d4696e0f5d54322ba7e0b3c80ca4656f4f61894682af2e9107c56cc1c21) |
| `qradar_receiver` | [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-fb8d7c7d87103f6fdf99293b8e667c3be11670787d98f3d89f9641bb9f2b4f5c) |
| `qradar_receiver.batch` | [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-89f6b3ae8d775713f1b99e617e71009649aedc3d46a63069e0dafd1486d82fbc) |
| `qradar_receiver.batch.max_bytes` | [qradar_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-003.md#canonical-ade6c2fd590b14093a78241888f0255dbc2a7cffa1309168164075e518a1ba03) |
| `qradar_receiver.batch.max_bytes_disabled` | [qradar_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-9324f6ee38dd37b5ba7cb5b8d9d2b48158493aefa4d651abf20d74d3a580d058) |
| `qradar_receiver.batch.max_events` | [qradar_receiver.batch.max_events](resources--global_log_receiver--reference--group-003.md#canonical-290776540be479258eeb7034db60b29168db272d9e8ec00c7700d24a7ad1b00b) |
| `qradar_receiver.batch.max_events_disabled` | [qradar_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-9f266a5bc481b831e319e610a657ef7a9a80654212ae2b0f17b8deab69fe3e0f) |
| `qradar_receiver.batch.timeout_seconds` | [qradar_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-003.md#canonical-bd39b1eec30cf059139000ee20d11b29a4e50977ef09abaddb5aa593d1b4412a) |
| `qradar_receiver.batch.timeout_seconds_default` | [qradar_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-6ef395f56608dfdf7aee92b9de073781e9d241f42c1dec2769d9086c162bc15e) |
| `qradar_receiver.compression` | [qradar_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-306dcb2624983ce2e01787c91252caabccef8af00c84cd570d5b93af0e024a7d) |
| `qradar_receiver.compression.compression_default` | [qradar_receiver.compression.compression_default](resources--global_log_receiver--reference--group-003.md#canonical-aac82c3a70da25647c2bd2e9fe1a425e0bf70ce0697eee59589fd10202a9106e) |
| `qradar_receiver.compression.compression_gzip` | [qradar_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-c190d0fd5813fb55ac422094ef2acb21bfce6aae9fc4565b2f743ce5a2f04ca0) |
| `qradar_receiver.compression.compression_none` | [qradar_receiver.compression.compression_none](resources--global_log_receiver--reference--group-003.md#canonical-02213be623e582aa51b001c964024a6d375210517c2bf508022d2dc8c54fa766) |
| `qradar_receiver.no_tls` | [qradar_receiver.no_tls](resources--global_log_receiver--reference--group-003.md#canonical-8545974d89fa8cd5ef85d928bf99dc2f6aabc862d17d94baa22402156e03f81d) |
| `qradar_receiver.uri` | [qradar_receiver.uri](resources--global_log_receiver--reference--group-003.md#canonical-05eee3e05588c7e0b159e040592ca066a14aa2f3800f5a898d291c7ad76e60a7) |
| `qradar_receiver.use_tls` | [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2fceb74b19c331681087b72070381ca9c81817ea640f3cd83b5f6c86d55dcb96) |
| `qradar_receiver.use_tls.disable_verify_certificate` | [qradar_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-dea4fd7ebb1a920087e754ac6817a2d536db93476bf91d2046844353587d59a7) |
| `qradar_receiver.use_tls.disable_verify_hostname` | [qradar_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-004.md#canonical-4f9b833de127ee59b2229bc1b4e6622c2d0d1526fca8969b8c4b2bb72a21448d) |
| `qradar_receiver.use_tls.enable_verify_certificate` | [qradar_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-004.md#canonical-d32705e9616d742f8b5360459e09b8304f56a09a27910b1de3dc1783dbc41670) |
| `qradar_receiver.use_tls.enable_verify_hostname` | [qradar_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-004.md#canonical-d68c229a317c9fbe11c0266984635fdcd28f4b02d7fc9f7ed9b38392a820e83c) |
| `qradar_receiver.use_tls.mtls_disabled` | [qradar_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-004.md#canonical-2ccf5178084023329694cc9be9ba0cdd6d4f3e1fd87b763d2a34d0334c283341) |
| `qradar_receiver.use_tls.mtls_enable` | [qradar_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-004.md#canonical-4fd443a8823e3b8d2db1105b7ee5a1605f2d76ba93011dde30b0a8463a068935) |
| `qradar_receiver.use_tls.mtls_enable.certificate` | [qradar_receiver.use_tls.mtls_enable.certificate](resources--global_log_receiver--reference--group-004.md#canonical-288e09876078c9c7e41732c1c913e15df6bdae41626295c7d1fe455e248cb601) |
| `qradar_receiver.use_tls.mtls_enable.key_url` | [qradar_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-004.md#canonical-ddd89468fa2fc7b7e756bd08f7ae3d7d0668a913fadf290165c8d225c7bd285f) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-e1f3a5055a60a03a45f360fe86de5eadfdc419d2a8650d4500387928bb0a4893) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-004.md#canonical-f39a35629264016bc29edcd5349c45b77601608c8eba4ae1aa53d4ad8c3f7351) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](resources--global_log_receiver--reference--group-004.md#canonical-2e37cb2e1939a543d7e220582f8bb6c515324993e9b2fb042131ecd904571750) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-004.md#canonical-7f112b088cae47d575c21e312c8c5ab879047676e4be8509bb423d8d59c4c517) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-ba777ebebe1e762bab61fefec3f2e718983411c6f8c75f0ceee6e3de4be6d3d2) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-004.md#canonical-e7e41ea75efbbc6eb8b0a1c8ae74e701d88a5c11cbd4a3471741fd3b8676f3db) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](resources--global_log_receiver--reference--group-004.md#canonical-4ea7321cbbaa30978e08b6430cbef1bedcadb919643d25915320da73e38e9659) |
| `qradar_receiver.use_tls.no_ca` | [qradar_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-004.md#canonical-48c73cf291156d0aee472374eef6f3d78c46ad2daf836077640228d47a24c454) |
| `qradar_receiver.use_tls.trusted_ca_url` | [qradar_receiver.use_tls.trusted_ca_url](resources--global_log_receiver--reference--group-003.md#canonical-e36baa59a1e43aaa6c189a3198b82ab809d3cd6b86cfed48f28f8a7a25f7b3c5) |
| `request_logs` | [request_logs](resources--global_log_receiver--reference--group-004.md#canonical-96ffebd451cfc47ca3cf925b357317d8d707c87a71bd79f2455049437290fa25) |
| `request_logs.sampled` | [request_logs.sampled](resources--global_log_receiver--reference--group-004.md#canonical-49288d41fbbc0005931dd8b84d537ab8efd7dad5764a0b3b06b940fe090a5068) |
| `request_logs.unsampled` | [request_logs.unsampled](resources--global_log_receiver--reference--group-004.md#canonical-539479257245b2ac7ef9072fa9880adb71eb2e8bda9147516f66e30816aba888) |
| `s3_receiver` | [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-98cdd3955d05bfc85f5d3bccc0c658806fb0dc88e48cebda63bbd2b3f7fb1a31) |
| `s3_receiver.aws_cred` | [s3_receiver.aws_cred](resources--global_log_receiver--reference--group-004.md#canonical-360da9dd58f72ef7414fdc7e065f5800de27df80a9de5e7502c222c41c4b57ce) |
| `s3_receiver.aws_cred.name` | [s3_receiver.aws_cred.name](resources--global_log_receiver--reference--group-004.md#canonical-b82e89a0a13105428e709babbd60d9ad61fa4c2ce92a826ba6431f245c7037d9) |
| `s3_receiver.aws_cred.namespace` | [s3_receiver.aws_cred.namespace](resources--global_log_receiver--reference--group-004.md#canonical-d9c8d4dd96ac25b9526c9c83293cdac9ce223075a86a8ca0d90fd225ea8def76) |
| `s3_receiver.aws_cred.tenant` | [s3_receiver.aws_cred.tenant](resources--global_log_receiver--reference--group-004.md#canonical-c1be9fe782ef9d4a8b8af3b60b6a7e0e3979899f638e408dadc9666ad35f1a26) |
| `s3_receiver.aws_region` | [s3_receiver.aws_region](resources--global_log_receiver--reference--group-004.md#canonical-6d21b580fb56aaadb89b712b9c747e177041cf64f58c634a2e99011a380d7606) |
| `s3_receiver.batch` | [s3_receiver.batch](resources--global_log_receiver--reference--group-004.md#canonical-bdf8d81a9cebcdae80cbdbbc8fad8f5bb8b603e4dbb50fd6467d427ba3c11908) |
| `s3_receiver.batch.max_bytes` | [s3_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-004.md#canonical-17447f59dc56619bf93243782e5d1f640099c67d69a9c441024c68ee533c425c) |
| `s3_receiver.batch.max_bytes_disabled` | [s3_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-004.md#canonical-cbb2a8ba7a3a0dbf97e4aa47e74132fd0778b02460a89afae24f6e4e560a4492) |
| `s3_receiver.batch.max_events` | [s3_receiver.batch.max_events](resources--global_log_receiver--reference--group-004.md#canonical-40fedb927c48c0c9cdcdb8649927ca42fff3ef01a39f82afebeb2b5ccd60372a) |
| `s3_receiver.batch.max_events_disabled` | [s3_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-004.md#canonical-0aa9a475a0eab4f753db62fc914c1d1fb1de9c5b5dafc703b8f799e8aaa1483b) |
| `s3_receiver.batch.timeout_seconds` | [s3_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-004.md#canonical-ef9d7456399de37923513d4d09ee6ccebeaa988b81050cbaa72f70f45943b2ab) |
| `s3_receiver.batch.timeout_seconds_default` | [s3_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-004.md#canonical-b6fada739c10de424d81fae0e3a4a721030d60918fc2017776ec7d7e5ebd76d5) |
| `s3_receiver.bucket` | [s3_receiver.bucket](resources--global_log_receiver--reference--group-004.md#canonical-a797e04524711153ff5cd63f3c98c2ede426dd2dffbb1a6e2700ee219a955cad) |
| `s3_receiver.compression` | [s3_receiver.compression](resources--global_log_receiver--reference--group-004.md#canonical-4590c6532e21e05ee5fd2b3c9836d3bbf913e2340cc47fb8e2be9b305f84b076) |
| `s3_receiver.compression.compression_default` | [s3_receiver.compression.compression_default](resources--global_log_receiver--reference--group-004.md#canonical-c7b1efcde10d1778b36208f0299ff09a18af1763e519e01c68e2c8231a577221) |
| `s3_receiver.compression.compression_gzip` | [s3_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-004.md#canonical-4e087c2708ac279e0483d974147609c51ef7c2e10f86ed7cc3a32750cd0e4f0c) |
| `s3_receiver.compression.compression_none` | [s3_receiver.compression.compression_none](resources--global_log_receiver--reference--group-004.md#canonical-758601a3bebe9feba5612dfc189eeb7644d21929b67704253396b136a958e0a6) |
| `s3_receiver.filename_options` | [s3_receiver.filename_options](resources--global_log_receiver--reference--group-004.md#canonical-86e3ebf382b11b5eeded21c840fddf2b276a706b528ff3c7c0bf87b8d6fe8beb) |
| `s3_receiver.filename_options.custom_folder` | [s3_receiver.filename_options.custom_folder](resources--global_log_receiver--reference--group-004.md#canonical-0a1bae5b8759f660604956a5842184ce78846f24f2703ca14c5e333b218a0404) |
| `s3_receiver.filename_options.log_type_folder` | [s3_receiver.filename_options.log_type_folder](resources--global_log_receiver--reference--group-004.md#canonical-382f0914670ed237ff1629d0cb2d7866f86c4f0d57187fcd9a61eb25c970f997) |
| `s3_receiver.filename_options.no_folder` | [s3_receiver.filename_options.no_folder](resources--global_log_receiver--reference--group-004.md#canonical-ef88a8c0d36dba5ea49fb0246de27d5541fab013170aa8594041189f4c32db1d) |
| `security_events` | [security_events](resources--global_log_receiver--reference--group-004.md#canonical-659531f836a283f5296910f5d1b17bd5537ae8c7d9e849c962f9784e820ed92f) |
| `splunk_receiver` | [splunk_receiver](resources--global_log_receiver--reference--group-004.md#canonical-fb1b9a8e316ed284345f44f335ee6101b9a45ab3fe45ee70bcf5b8e162f38684) |
| `splunk_receiver.batch` | [splunk_receiver.batch](resources--global_log_receiver--reference--group-004.md#canonical-c1471a21b791cf1bb6a2c863e16adfd752f23f357c31ef2f9f173567305e1c9d) |
| `splunk_receiver.batch.max_bytes` | [splunk_receiver.batch.max_bytes](resources--global_log_receiver--reference--group-004.md#canonical-eaf3cb71e9a4e312888811f1528bd842ca2cda18b86a6115d5a556c0b06a32ec) |
| `splunk_receiver.batch.max_bytes_disabled` | [splunk_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-004.md#canonical-5c4fae142eaa826ae8be2c12fa2811e570c6aedd9bc3301b7fd9bc3a16458620) |
| `splunk_receiver.batch.max_events` | [splunk_receiver.batch.max_events](resources--global_log_receiver--reference--group-004.md#canonical-72d188c991408e0ec24a1ac287e519723ec074c105526ea7d58cd33a6a59b68b) |
| `splunk_receiver.batch.max_events_disabled` | [splunk_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-004.md#canonical-25e006c7b18fc39898fb1dee5c38391c17f865456cada20ba9d2b94c58e5d6c7) |
| `splunk_receiver.batch.timeout_seconds` | [splunk_receiver.batch.timeout_seconds](resources--global_log_receiver--reference--group-004.md#canonical-d262aefb22a399726a9cdcb840d96cd62768bacb2beda8d0997b84e39a420fb3) |
| `splunk_receiver.batch.timeout_seconds_default` | [splunk_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-004.md#canonical-029661d6a5df9aba6c6500f22d7fd257e6f1e85faf5dbc4c236b8b544cc87946) |
| `splunk_receiver.compression` | [splunk_receiver.compression](resources--global_log_receiver--reference--group-004.md#canonical-1b493ef3b1b46fef2d7d03a282836703a63aea025f0d78d0fc6308a5d70ea70c) |
| `splunk_receiver.compression.compression_default` | [splunk_receiver.compression.compression_default](resources--global_log_receiver--reference--group-004.md#canonical-dfa316a46dae8cacb9e6f214b3f95d0f67cd1b04f6fe9f28b27ac21d66e58875) |
| `splunk_receiver.compression.compression_gzip` | [splunk_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-004.md#canonical-cfe25a144e62b6bb06b1c2f2924509e56f71fed60e3f8a42819b40556d441d21) |
| `splunk_receiver.compression.compression_none` | [splunk_receiver.compression.compression_none](resources--global_log_receiver--reference--group-004.md#canonical-e7b573e722897774b1c3bf03ba3708204736b3fd0aa72d43341e0d07b90edcbc) |
| `splunk_receiver.endpoint` | [splunk_receiver.endpoint](resources--global_log_receiver--reference--group-004.md#canonical-e5851b75a525da0092d6935fa27402027ae96bbcf06309ff30be488f8d1001a5) |
| `splunk_receiver.no_tls` | [splunk_receiver.no_tls](resources--global_log_receiver--reference--group-004.md#canonical-b281fbaea4c06ff5d7d18dde54d97f90237349ad13bef7453d3b14c45bf89abe) |
| `splunk_receiver.splunk_hec_token` | [splunk_receiver.splunk_hec_token](resources--global_log_receiver--reference--group-004.md#canonical-2682a8104a181900fc201f21f24d9457a0938ddefe6389cc002bffc8ee409c89) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info` | [splunk_receiver.splunk_hec_token.blindfold_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-57461e608facc239debefbb723332acbd4ad14e1ac0364e3730e1d740b3ce52a) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.decryption_provider` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-004.md#canonical-940f9ada56b643b49e54e003bf07de5edf32fbeb8f46196de574c4460eb26f08) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.location` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.location](resources--global_log_receiver--reference--group-004.md#canonical-b51a10ad9d354ab7fbc11d4607d270f3b0a222e5368f3e991b460c40d2912d1a) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.store_provider` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-004.md#canonical-121f28ba70c8faa98484498c2e4ecf51d639d6749e0f3f7520182808fb3106d6) |
| `splunk_receiver.splunk_hec_token.clear_secret_info` | [splunk_receiver.splunk_hec_token.clear_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-e83e26cf4a2f57d8ab52b9121ef66da240bd98326caf9702fba425082c3f1578) |
| `splunk_receiver.splunk_hec_token.clear_secret_info.provider_ref` | [splunk_receiver.splunk_hec_token.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-004.md#canonical-473eb03f325ecd6f4de42066a46926d84849aecb87f8478b685768e0ec332985) |
| `splunk_receiver.splunk_hec_token.clear_secret_info.url` | [splunk_receiver.splunk_hec_token.clear_secret_info.url](resources--global_log_receiver--reference--group-004.md#canonical-904411f569e303c99e65116433215d7ac5b5c2d05a391ac50a1a2c106d7f00ce) |
| `splunk_receiver.use_tls` | [splunk_receiver.use_tls](resources--global_log_receiver--reference--group-004.md#canonical-7118da65c46bcb27083a7c2113124a1011ea27e103fce4f6a4e07176cf8a6f3f) |
| `splunk_receiver.use_tls.disable_verify_certificate` | [splunk_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-004.md#canonical-42f8dca5d2d72391b3197e5792b6128ab58ea99508067df58668bae2796d6730) |
| `splunk_receiver.use_tls.disable_verify_hostname` | [splunk_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-004.md#canonical-3729d93ffc5c174d0c6fe5e5a7ec4d1bc39439e354bd532346f48bdf4796772d) |
| `splunk_receiver.use_tls.enable_verify_certificate` | [splunk_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-004.md#canonical-bd46ced2f7755992356f32697fe2f0c65019bf5bb0259f19cc3299988eabbe45) |
| `splunk_receiver.use_tls.enable_verify_hostname` | [splunk_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-004.md#canonical-f4ee9544287aed0cc15e0373b3d41c80fde1419dfdd21a8cc7639fcbe3964043) |
| `splunk_receiver.use_tls.mtls_disabled` | [splunk_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-004.md#canonical-0864644e4f60cd1403d0d0fef033eb1a8117a6e0c797659906e81a618f9c9429) |
| `splunk_receiver.use_tls.mtls_enable` | [splunk_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-004.md#canonical-24b78d71d3f4b039ba68644a6cf5d8745b36af98880f831f8f766246b63175fc) |
| `splunk_receiver.use_tls.mtls_enable.certificate` | [splunk_receiver.use_tls.mtls_enable.certificate](resources--global_log_receiver--reference--group-004.md#canonical-fb33cc747b25f1f877fc4f7d245dcb4d2371e525c60c618e78cb6678b9e43366) |
| `splunk_receiver.use_tls.mtls_enable.key_url` | [splunk_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-004.md#canonical-41ccbb0fe7b78fcff2204af295ad003330d9140e3fd228d39d6679e978abfb04) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-c9ddeb955d97e9fc656c63db00643a21f9ab3a17349f7776c06c8bf4e8a43b2e) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-004.md#canonical-c33abc8f9e60ff16dee592184afbc06590fb4e830ca587625fe65830ab846a1d) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](resources--global_log_receiver--reference--group-004.md#canonical-fc42af6e99c47db47dbaefbdaba3625ddbae87ae5769aa724f4aab567a7cb68a) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-004.md#canonical-54787a02bd5ee4ba94caa9151a1968a1146879adbec40be99cba27333724dd8d) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-acfb92f4416909150b63b228ee2ea8951080118297595129945851956c5866ff) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-004.md#canonical-49d6705015364669ca2f9e36921feaacd455d70fb5eecdfa20aa2ba2acb68790) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](resources--global_log_receiver--reference--group-004.md#canonical-a43b8dcb5726e90f3df1fa5cb97017114d5dbb6421e65ef4cfe9d7389da2d8e9) |
| `splunk_receiver.use_tls.no_ca` | [splunk_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-004.md#canonical-062a1d9f5f92804e70b7e4e4f9d33394c0ab1bf800344b81db7abe19426b7e89) |
| `splunk_receiver.use_tls.trusted_ca_url` | [splunk_receiver.use_tls.trusted_ca_url](resources--global_log_receiver--reference--group-004.md#canonical-28170a913c4dbf1420e7e6c974a930d0561432f4fa303acc4e966258f6f88cc9) |
| `sumo_logic_receiver` | [sumo_logic_receiver](resources--global_log_receiver--reference--group-004.md#canonical-57a21bb6254769a1d2dbfcd710557e3a02117ebe180a98dab044b7e56a1cd57a) |
| `sumo_logic_receiver.url` | [sumo_logic_receiver.url](resources--global_log_receiver--reference--group-004.md#canonical-3d9899cec85e0af859bb3625d82e7bf53153d782e183faf1256b300988afcea7) |
| `sumo_logic_receiver.url.blindfold_secret_info` | [sumo_logic_receiver.url.blindfold_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-c48f5a01429254024d35757e6a64c9dd8c35eeb091c3623781dcf071fa758168) |
| `sumo_logic_receiver.url.blindfold_secret_info.decryption_provider` | [sumo_logic_receiver.url.blindfold_secret_info.decryption_provider](resources--global_log_receiver--reference--group-004.md#canonical-557011d8be9f502df7b4d48f474f279cca7df60c80b4fb5ca93a6d6cae72cbaf) |
| `sumo_logic_receiver.url.blindfold_secret_info.location` | [sumo_logic_receiver.url.blindfold_secret_info.location](resources--global_log_receiver--reference--group-004.md#canonical-ccb715f32ca0dbec020f267e9aff477d70409c84d0fc5f08ff59c6ac5920ba0a) |
| `sumo_logic_receiver.url.blindfold_secret_info.store_provider` | [sumo_logic_receiver.url.blindfold_secret_info.store_provider](resources--global_log_receiver--reference--group-004.md#canonical-ea1fce7fa6816ad3d92b7c3594ad703c2aba2ce9c1dcb5833c35f073cd82e6b9) |
| `sumo_logic_receiver.url.clear_secret_info` | [sumo_logic_receiver.url.clear_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-ec7cd7a0c7ef02691629cb04e2b7daef48b87e2872490280c88a052518674ddc) |
| `sumo_logic_receiver.url.clear_secret_info.provider_ref` | [sumo_logic_receiver.url.clear_secret_info.provider_ref](resources--global_log_receiver--reference--group-004.md#canonical-b5a17ee2fef8f3307cb491173855f504161c0d00c4a1d77b40dc8d216087a005) |
| `sumo_logic_receiver.url.clear_secret_info.url` | [sumo_logic_receiver.url.clear_secret_info.url](resources--global_log_receiver--reference--group-004.md#canonical-b96b908c64bb4aaf146d471472120c4d61710248301713c680fadf35db8f6e42) |
| `timeouts` | [timeouts](resources--global_log_receiver--reference--group-004.md#canonical-bb47a848a11d9b79d4850a974bbd6f26976cee7d8d0507ef2e5d09056637be46) |
| `timeouts.create` | [timeouts.create](resources--global_log_receiver--reference--group-004.md#canonical-9c40a0322f216b50634cd8e029b9fad8448a4723710a512c1fad4ce33cbf3ec7) |
| `timeouts.delete` | [timeouts.delete](resources--global_log_receiver--reference--group-004.md#canonical-d9677401247c9b3f864faffd7516b38bf55d775ca3dbd30def1ffa5d93566cc2) |
| `timeouts.read` | [timeouts.read](resources--global_log_receiver--reference--group-004.md#canonical-6929bfd1c3d65bd24f8f4c65d5437286ad423f708f691ee865d82defc1aad248) |
| `timeouts.update` | [timeouts.update](resources--global_log_receiver--reference--group-004.md#canonical-37163117f6099b0724406c3ed50e8319cf499ab579d8eb84c292423f92762ad2) |

<a id="canonical-d659a9f7c5170c45275614110ac4a5921e2b30c6fb9e459b0a6f61fdf61f7bcd"></a>

## Next pages — Property reference / 22597ce1fe7c / 12

- [audit_logs](resources--global_log_receiver--reference--group-001.md#canonical-ed77b985cced6a9e6caaa8ade6368f0dadca11b7d3e334c3d25cf40f5e593b46)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f)
- [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-31bb2a7088509e10c5d9bc0ff4558a7cd09ca1905ff420ddd31ec84beba5c7f6)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-d3cec99e849c282a720b0675fb155592bf8cb034e8e127ea487395b83385380d)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-db0eb1eadc3b3cbc91f20800c8e0c2215dcd8305de5e84d6d41a37901c086f9f)
- [dns_logs](resources--global_log_receiver--reference--group-002.md#canonical-e5bfd6707af5d068121540614c989041b57c8f1b208696a94558f4252c6fc4b7)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-88f4e7ab594faa932dc1655ae2a56d356602da8ca018ed5cc5e9a181411c85cd)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-bf2394cf5d8419274cfd823dd262328c8899eadae65313d7b15a65f3204709bf)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-ee10d6aa6bae7f44f1fe9b11b6baa91894549e346d1cced532d36d362b306e96)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-5094ca81331d2853ad9a7287be56bcb5f0fc071accea0025c24e2362b7697388)
- [ns_all](resources--global_log_receiver--reference--group-003.md#canonical-94ac97fe81951dfbc63beb0bdb84d4b3cfbea1343ac31d194e4cfe249e881911)
- [ns_current](resources--global_log_receiver--reference--group-003.md#canonical-5d4a9b8030a880ba327e5064918fb59a3565b4a15a507da773657ce9558b8eaf)
- [ns_list](resources--global_log_receiver--reference--group-003.md#canonical-7b1e2acca773fcf6415527dd41a12a76c20c635cd71ada5012b1e73e227f97e3)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-afb382cb3e4d5e62efc324cc347c6ce43343131b41944fdb2dca84647cc5f8be)
- [request_logs](resources--global_log_receiver--reference--group-004.md#canonical-6e273b67339140e7d510ed896f832f6e797c61fd0cdd910df908cae2e7cda8a5)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-bf56c64665781b931da6f4c5b53da70995427fed0ac9dcdc9cc103366e6dce5a)
- [security_events](resources--global_log_receiver--reference--group-004.md#canonical-e5289c97403a4948d700a4317e634154970ec974a599dc422156e5ba5ab04c9a)
- [splunk_receiver](resources--global_log_receiver--reference--group-004.md#canonical-7d67c978d6feb3f69e05312941f8bb07957fbcfa0f320a8bdc357182fe47d0dc)
- [sumo_logic_receiver](resources--global_log_receiver--reference--group-004.md#canonical-df52257dd419eb2d77d607a0f7da206f3cfaeade103803f055d0759b33652f1c)
- [timeouts](resources--global_log_receiver--reference--group-004.md#canonical-b033c7a246a27eb370c6377276894c7f70ffadc4a31122de2469c3da4f67269e)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-ed77b985cced6a9e6caaa8ade6368f0dadca11b7d3e334c3d25cf40f5e593b46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bf90caf03a3e7bcd63222f0e18c36ad158046640cee0ce6e1a34760d609839b"></a>

## audit_logs — audit_logs / 86b8d1f98e6a / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- audit_logs

<a id="canonical-11bf8c72675ef03587022c3fc826e12624118be689fffe9c02f6e9bd0bd1168f"></a>

Type: `["object", {}]`. Optional.

\[OneOf: audit\_logs, dns\_logs, request\_logs, security\_events\] Enable this option

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

OneOf alternatives in this subsection:

- [audit_logs](resources--global_log_receiver--reference--group-001.md#canonical-11bf8c72675ef03587022c3fc826e12624118be689fffe9c02f6e9bd0bd1168f)
- [dns_logs](resources--global_log_receiver--reference--group-002.md#canonical-9660a0fa62627fdf04ca090bee24c9b62e6e590aa3dbbaded094ca4ed6ef7acd)
- [request_logs](resources--global_log_receiver--reference--group-004.md#canonical-96ffebd451cfc47ca3cf925b357317d8d707c87a71bd79f2455049437290fa25)
- [security_events](resources--global_log_receiver--reference--group-004.md#canonical-659531f836a283f5296910f5d1b17bd5537ae8c7d9e849c962f9784e820ed92f)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
audit_logs = {}
```

<a id="canonical-ce2204c7f6720b0c12a29a15ebffde089fbf50e74d5aae98fc735de995e4d754"></a>

## Direct properties — audit_logs / 86b8d1f98e6a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8b0892271be5b9303bce3d0930cc3ea2a641f83f034f75be91f4ea523bef5fdc"></a>

## Next pages — audit_logs / 86b8d1f98e6a / 4

- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2a36580c64d93c5f5fb52f1228db3b0ffab210a6a941908d4e4279dae6aab15"></a>

## aws_cloud_watch_receiver — aws_cloud_watch_receiver / 7647b10f65b5 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- aws_cloud_watch_receiver

<a id="canonical-0435c60661e00b6ebd7bf410341885fa5899d1ebc9d802dabf6d655c266a6312"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: aws\_cloud\_watch\_receiver, azure\_event\_hubs\_receiver, azure\_receiver,
datadog\_receiver, gcp\_bucket\_receiver, http\_receiver, kafka\_receiver, new\_relic\_receiver,
qradar\_receiver, s3\_receiver, splunk\_receiver, sumo\_logic\_receiver\] AWS Cloudwatch Logs
Configuration for Global Log Receiver.

Upstream description:

AWS Cloudwatch Logs Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_region",
    "group_name",
    "stream_name")}
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

OneOf alternatives in this subsection:

- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-0435c60661e00b6ebd7bf410341885fa5899d1ebc9d802dabf6d655c266a6312)
- [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-b881e14c6dfa052502f651b8a706084c45f750454aca6f7e13d12cd08d9b0602)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-c782dfc873188e384141ccd9e5104133b8539660526f46ac90e75f3272573d6d)
- [datadog_receiver](resources--global_log_receiver--reference--group-002.md#canonical-033e07e78970548fb9be688f0413eb139106cc6737071553d6e7a4e28928697e)
- [gcp_bucket_receiver](resources--global_log_receiver--reference--group-002.md#canonical-69146732819d2a49dcb135b7658185ce2cbee211fdc7b6292822455a6f476f55)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-68e14810058d293f224852c9c18238dd7d415dd26f2dc869cc6e7dfa4c053942)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-30974c0ff3bb7bbe13c8e4b6fdd227339653da44184623051eff1f0043a6b046)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-6c19c494b563c2d60abdc9410439454b737deb2ef492e68a3bff9eec490830d5)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-fb8d7c7d87103f6fdf99293b8e667c3be11670787d98f3d89f9641bb9f2b4f5c)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-98cdd3955d05bfc85f5d3bccc0c658806fb0dc88e48cebda63bbd2b3f7fb1a31)
- [splunk_receiver](resources--global_log_receiver--reference--group-004.md#canonical-fb1b9a8e316ed284345f44f335ee6101b9a45ab3fe45ee70bcf5b8e162f38684)
- [sumo_logic_receiver](resources--global_log_receiver--reference--group-004.md#canonical-57a21bb6254769a1d2dbfcd710557e3a02117ebe180a98dab044b7e56a1cd57a)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
aws_cloud_watch_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-2d7e724b856272b053aa1b5acf1f5003be7ed5061840fa9dc6ab1fff70857de3"></a>

## Direct properties — aws_cloud_watch_receiver / 7647b10f65b5 / 3

- [aws_cred](resources--global_log_receiver--reference--group-001.md#canonical-d92667a2de9c58ff02edffb7d9f09eebb3b281fb5702d1c55af4034c5a3deb31): complete subsection reference.

<a id="canonical-8730c9853f6863274ded542197a581fa4324f1034dc61ab04205fa48cd9d6410"></a>

<a id="canonical-b93191c3df51beab2b1a7779558d88c058ffdcfe093421ec2b576dbba3f1db95"></a>

## aws_region property — aws_cloud_watch_receiver / 7647b10f65b5 / 4

Type: `"string"`. Optional.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
AWS Region. AWS Region Name. Possible values are \`ap-northeast-1\`, \`ap-southeast-1\`,
\`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`, \`us-east-2\`,
\`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`, \`ap-northeast-2\`,
\`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`, \`me-south-1\`, \`us-west-1\`,
\`ap-southeast-3\`.

Upstream description:

AWS Region Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  }
}
```

- [batch](resources--global_log_receiver--reference--group-001.md#canonical-47aed24580815b6c3dc47b3efa5f0fd20136e6dd1b34d1fbbaeb32a4fcb1d602): complete subsection reference.

- [compression](resources--global_log_receiver--reference--group-001.md#canonical-42f1b43dd12dd17f4efd5bcfe819c7e609a0add4b9b08905fceede0296b705e1): complete subsection reference.

<a id="canonical-fac31f4a10834165053e6e8eafdbe050d8ba14876851abd520304a3ab2f008e6"></a>

<a id="canonical-ca387034a0a2fb18cc8561ad717344b6a8fcbabbd9587b592ced293100367681"></a>

## group_name property — aws_cloud_watch_receiver / 7647b10f65b5 / 5

Type: `"string"`. Optional.

The group name of the target Cloudwatch Logs stream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[\\\\.\\\\-_/#A-Za-z0-9]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[\\\\.\\\\-_/#A-Za-z0-9]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[\\\\.\\\\-_/#A-Za-z0-9]+$"
  }
}
```

<a id="canonical-f3a53f2de5a2078838c99ab634f8768323f14ab46697152eed5e9be1101f7b21"></a>

<a id="canonical-8798589dfcb5b6e1cbb978cbf964dbae992f47308ffee4c8b4d6ee5cefc36300"></a>

## stream_name property — aws_cloud_watch_receiver / 7647b10f65b5 / 6

Type: `"string"`. Optional.

The stream name of the target Cloudwatch Logs stream. Note that there can only be one writer to a
log stream at a time.

Upstream description:

The stream name of the target Cloudwatch Logs stream. Note that there can only be one writer to a
log stream at a time.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[^:*]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[^:*]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[^:*]*$"
  }
}
```

<a id="canonical-40874fe9c5826e4147577d802e709ebc792a10a9bf43f495413183d1a942a2d4"></a>

## Next pages — aws_cloud_watch_receiver / 7647b10f65b5 / 7

- [aws_cloud_watch_receiver.aws_cred](resources--global_log_receiver--reference--group-001.md#canonical-d92667a2de9c58ff02edffb7d9f09eebb3b281fb5702d1c55af4034c5a3deb31)
- [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-47aed24580815b6c3dc47b3efa5f0fd20136e6dd1b34d1fbbaeb32a4fcb1d602)
- [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-42f1b43dd12dd17f4efd5bcfe819c7e609a0add4b9b08905fceede0296b705e1)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-d92667a2de9c58ff02edffb7d9f09eebb3b281fb5702d1c55af4034c5a3deb31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c0500a30f22a1d5c7d4c38bb456cf30a727e85db98be40a05c2546a375a681a"></a>

## aws_cloud_watch_receiver.aws_cred — aws_cloud_watch_receiver.aws_cred / 4b7893b36745 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f)
- aws_cloud_watch_receiver.aws_cred

<a id="canonical-ae7891b0c9b785b25c4247a92943f4cb2e3b9774a7b3ce119d34cdfb7c8e7c39"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
aws_cred {
  # Configure direct properties listed below.
}
```

<a id="canonical-bddbdd303be4e978c699a99c044e43d1c3106f8f023a7d2a053c52b707a27d18"></a>

## Direct properties — aws_cloud_watch_receiver.aws_cred / 4b7893b36745 / 3

<a id="canonical-3900795eceef01770e02e0e186a3edf65a416adc801b9dbad01b1d7c2d087c87"></a>

<a id="canonical-35a893e3f907c380920e918756c5a8460b0fed820eaf1bc5a0a58b92948574f4"></a>

## name property — aws_cloud_watch_receiver.aws_cred / 4b7893b36745 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-dbccea2267252d5d0a90322a119eebae1caa213441bd19b3f165e8ec22dad98f"></a>

<a id="canonical-5611ea71d853c5e95a1b1f7a43e9f66f22f566262b861e0ef6a947d69d9ca38f"></a>

## namespace property — aws_cloud_watch_receiver.aws_cred / 4b7893b36745 / 5

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
}
```

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

<a id="canonical-014800280c8d47b1dcf937e9b633d703c0dad2de21f23ad70ad1c8ef151ef16b"></a>

<a id="canonical-16b19000324b552463e086301380b21f2bf73a43fdab31b7f4700fd527e732cf"></a>

## tenant property — aws_cloud_watch_receiver.aws_cred / 4b7893b36745 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-94e42655cc83865ed223206ec2a42f8489d8729e506c04993374602bf0a47431"></a>

## Next pages — aws_cloud_watch_receiver.aws_cred / 4b7893b36745 / 7

- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-47aed24580815b6c3dc47b3efa5f0fd20136e6dd1b34d1fbbaeb32a4fcb1d602"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a781ed3ede73de76901c097620c9e091ed45b9a690ca9653920f9c3b67b1365"></a>

## aws_cloud_watch_receiver.batch — aws_cloud_watch_receiver.batch / 157ed99e8d2d / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f)
- aws_cloud_watch_receiver.batch

<a id="canonical-a3f62fa0fb2640bf3c5a0c7fbecc945b068e08e3bb3a48ba2a43c9834caaca9d"></a>

Type: `"object"`. single nested block, Optional.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_bytes",
    "max_bytes_disabled"),
  validators.ConflictingObjectAttributes("max_events",
    "max_events_disabled"),
  validators.ConflictingObjectAttributes("timeout_seconds",
    "timeout_seconds_default")}
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
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

Terraform syntax:

```terraform
batch {
  # Configure direct properties listed below.
}
```

<a id="canonical-1e20ad03315a5eafa3f96bb377e5d67e9adc7f2e0325f8e3ca86dafb9d3cd931"></a>

## Direct properties — aws_cloud_watch_receiver.batch / 157ed99e8d2d / 3

<a id="canonical-be237f814c5ba1ca49e70625ef679160b630c0dab70b6c9d67565405409d61f9"></a>

<a id="canonical-0234bed5cc87c043b5832616c425409409a568a2a60b9e6ff1695d4f8d57b907"></a>

## max_bytes property — aws_cloud_watch_receiver.batch / 157ed99e8d2d / 4

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(4096, 10485760),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](resources--global_log_receiver--reference--group-001.md#canonical-16fe29d5c6d27f9633652676b11f47af283c5157192b87e36254a2c41cf285d4): complete subsection reference.

<a id="canonical-0bbb094f94b284ab317460c258dcd8c9e1028cc5eceea4447bc4af61964225ec"></a>

<a id="canonical-59e20303f3b5326e6ec3c7b9780e8687e292ad8e41cdb6d781e47ff2f68cdaa5"></a>

## max_events property — aws_cloud_watch_receiver.batch / 157ed99e8d2d / 5

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(32, 2000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](resources--global_log_receiver--reference--group-001.md#canonical-fb602c0f3dcff17c0dbc4fb640e16df9c9adb3c6c9435e34015b7de8207cd481): complete subsection reference.

<a id="canonical-eacdef4b128e8d9568a5b25f7447ae907a82e79ed649d445f2c9aaa130e78d27"></a>

<a id="canonical-57317e07658086bc4f70317d1013f5db7558611d29dc2a445a99ac8127e9490b"></a>

## timeout_seconds property — aws_cloud_watch_receiver.batch / 157ed99e8d2d / 6

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](resources--global_log_receiver--reference--group-001.md#canonical-7c21254e9133a13b8953d40556c1f57a2a89e2d353af1633d73f9e9ebe2acce0): complete subsection reference.

<a id="canonical-55bff7275a80b4a963540b30a7b2c3eca2089d037bc6a56b60a7d1686fad7c93"></a>

## Next pages — aws_cloud_watch_receiver.batch / 157ed99e8d2d / 7

- [aws_cloud_watch_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-001.md#canonical-16fe29d5c6d27f9633652676b11f47af283c5157192b87e36254a2c41cf285d4)
- [aws_cloud_watch_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-001.md#canonical-fb602c0f3dcff17c0dbc4fb640e16df9c9adb3c6c9435e34015b7de8207cd481)
- [aws_cloud_watch_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-001.md#canonical-7c21254e9133a13b8953d40556c1f57a2a89e2d353af1633d73f9e9ebe2acce0)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-16fe29d5c6d27f9633652676b11f47af283c5157192b87e36254a2c41cf285d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07302217b7b477846128444180b38182ada842e40624c92f07f49b36cea39be3"></a>

## aws_cloud_watch_receiver.batch.max_bytes_disabled — aws_cloud_watch_receiver.batch.max_bytes_disabled / 84bc3b2d5650 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f)
- [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-47aed24580815b6c3dc47b3efa5f0fd20136e6dd1b34d1fbbaeb32a4fcb1d602)
- aws_cloud_watch_receiver.batch.max_bytes_disabled

<a id="canonical-4f597c654e9c685225fd80d7a481567838d1e432eb7df9f1f405216402be89a7"></a>

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
max_bytes_disabled = {}
```

<a id="canonical-125c6d9854cb5e1097a88b365420ae06ad1e9b5a399978b1bdac9d1811cec886"></a>

## Direct properties — aws_cloud_watch_receiver.batch.max_bytes_disabled / 84bc3b2d5650 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1a48865472ed144f2ab3e54e70419d7bea99f9faa660432000f71fa8ab7a26af"></a>

## Next pages — aws_cloud_watch_receiver.batch.max_bytes_disabled / 84bc3b2d5650 / 4

- [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-47aed24580815b6c3dc47b3efa5f0fd20136e6dd1b34d1fbbaeb32a4fcb1d602)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-fb602c0f3dcff17c0dbc4fb640e16df9c9adb3c6c9435e34015b7de8207cd481"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-728070a11e273690d61fe9e535ed60e6270229e6df2d53a7dd37b14df3d7893a"></a>

## aws_cloud_watch_receiver.batch.max_events_disabled — aws_cloud_watch_receiver.batch.max_events_disabled / a979de6984ea / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f)
- [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-47aed24580815b6c3dc47b3efa5f0fd20136e6dd1b34d1fbbaeb32a4fcb1d602)
- aws_cloud_watch_receiver.batch.max_events_disabled

<a id="canonical-a2436b3f8467759fd63f43f83f7ecbded5ab89133471221303eaaa926704a647"></a>

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
max_events_disabled = {}
```

<a id="canonical-ed8bec4daeda3e17c904038c50b1676b4a3a5deecdc53c5444d47a4b2c85cf5a"></a>

## Direct properties — aws_cloud_watch_receiver.batch.max_events_disabled / a979de6984ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3ca7443e44b4b7467372e5445beb4247b3ae20b918590fd68fb6a33ccb8ce28b"></a>

## Next pages — aws_cloud_watch_receiver.batch.max_events_disabled / a979de6984ea / 4

- [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-47aed24580815b6c3dc47b3efa5f0fd20136e6dd1b34d1fbbaeb32a4fcb1d602)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-7c21254e9133a13b8953d40556c1f57a2a89e2d353af1633d73f9e9ebe2acce0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0588304f4d090323dc38ee1ae3082f650a911795370bbbd1fcb0522f428d87b"></a>

## aws_cloud_watch_receiver.batch.timeout_seconds_default — aws_cloud_watch_receiver.batch.timeout_seconds_default / 1f6a2e6ef6fa / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f)
- [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-47aed24580815b6c3dc47b3efa5f0fd20136e6dd1b34d1fbbaeb32a4fcb1d602)
- aws_cloud_watch_receiver.batch.timeout_seconds_default

<a id="canonical-262165dd3e9eccab67870cff0f5aa28964c2a0effa21c0cb4e06bc53b6ff69cf"></a>

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
timeout_seconds_default = {}
```

<a id="canonical-65a94a1eec67de7fa9a147b18eab0e1e458a34bc3c22f920cfda19577a4ce8bc"></a>

## Direct properties — aws_cloud_watch_receiver.batch.timeout_seconds_default / 1f6a2e6ef6fa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e65a9d429fe54478f0c7f53d6859e390715056769c6d8db28e9e942ff6d2ddf6"></a>

## Next pages — aws_cloud_watch_receiver.batch.timeout_seconds_default / 1f6a2e6ef6fa / 4

- [aws_cloud_watch_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-47aed24580815b6c3dc47b3efa5f0fd20136e6dd1b34d1fbbaeb32a4fcb1d602)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-42f1b43dd12dd17f4efd5bcfe819c7e609a0add4b9b08905fceede0296b705e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7b7bb31ef0a0eab002c0e70f62d6fceb6c4c2a6ce0ed110a6f697f93a823c54"></a>

## aws_cloud_watch_receiver.compression — aws_cloud_watch_receiver.compression / 8354b46cd50f / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f)
- aws_cloud_watch_receiver.compression

<a id="canonical-adce353312ca2a08817394439904e31e64a30622aa0e484b410d189fe121a684"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Upstream description:

Compression Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("compression_default",
    "compression_gzip"),
  validators.ConflictingObjectAttributes("compression_default",
    "compression_none"),
  validators.ConflictingObjectAttributes("compression_gzip",
    "compression_none")}
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
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

Terraform syntax:

```terraform
compression {
  # Configure direct properties listed below.
}
```

<a id="canonical-49460b7d9ad201f918ff7e9270e3f7b411c8cdd8a598396c6f98fe0a753f2377"></a>

## Direct properties — aws_cloud_watch_receiver.compression / 8354b46cd50f / 3

- [compression_default](resources--global_log_receiver--reference--group-001.md#canonical-3ee08d55904c05abfdcc7fda975c880e3ac85cff193b9d25d5f5440df607cb0d): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-001.md#canonical-621d86866e81b00d028fee88c22db82b8f025496371e1365c834acaabbd9958a): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-001.md#canonical-832b7d0e9e7e12fa82228ce60fa1a842ca909cb0e16cc25e35830cc69bacd9b3): complete subsection reference.

<a id="canonical-393eaa33955c8719d6decfbe744cd1989e5a01f007b4d1862f4af15417f0d63a"></a>

## Next pages — aws_cloud_watch_receiver.compression / 8354b46cd50f / 4

- [aws_cloud_watch_receiver.compression.compression_default](resources--global_log_receiver--reference--group-001.md#canonical-3ee08d55904c05abfdcc7fda975c880e3ac85cff193b9d25d5f5440df607cb0d)
- [aws_cloud_watch_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-001.md#canonical-621d86866e81b00d028fee88c22db82b8f025496371e1365c834acaabbd9958a)
- [aws_cloud_watch_receiver.compression.compression_none](resources--global_log_receiver--reference--group-001.md#canonical-832b7d0e9e7e12fa82228ce60fa1a842ca909cb0e16cc25e35830cc69bacd9b3)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-3ee08d55904c05abfdcc7fda975c880e3ac85cff193b9d25d5f5440df607cb0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-694a98b8399789a5f0bd88ca78e1e00dceae92036bf70ea4e90ec085cb1644d1"></a>

## aws_cloud_watch_receiver.compression.compression_default — aws_cloud_watch_receiver.compression.compression_default / 1b34fb74e5a3 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f)
- [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-42f1b43dd12dd17f4efd5bcfe819c7e609a0add4b9b08905fceede0296b705e1)
- aws_cloud_watch_receiver.compression.compression_default

<a id="canonical-30b4e1638daa7a2fcb9def861003c2e9fb0f3372a378921a46355bd9ffdf39b7"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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
compression_default = {}
```

<a id="canonical-5842377c00db27a95392aafa1ab3950b489358d9d232cd22b6a234e378827b7d"></a>

## Direct properties — aws_cloud_watch_receiver.compression.compression_default / 1b34fb74e5a3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-994fdf2454a64aba68b2bfb7723a474276d66bbe5ba6028e92c781759bd0012f"></a>

## Next pages — aws_cloud_watch_receiver.compression.compression_default / 1b34fb74e5a3 / 4

- [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-42f1b43dd12dd17f4efd5bcfe819c7e609a0add4b9b08905fceede0296b705e1)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-621d86866e81b00d028fee88c22db82b8f025496371e1365c834acaabbd9958a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11b3c08165963e808922085225dbdf49bc24263c76e42c98a26a7bf8158049b7"></a>

## aws_cloud_watch_receiver.compression.compression_gzip — aws_cloud_watch_receiver.compression.compression_gzip / 1d9151c131d0 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f)
- [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-42f1b43dd12dd17f4efd5bcfe819c7e609a0add4b9b08905fceede0296b705e1)
- aws_cloud_watch_receiver.compression.compression_gzip

<a id="canonical-ecb0bb5a86aebbfccac59212bae382aab96c11a3cf33596dcad6c6cbc357ca8f"></a>

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
compression_gzip = {}
```

<a id="canonical-91bea4c332b01159dab49d669d8170525372959f4187f3fb350ee38c24befb54"></a>

## Direct properties — aws_cloud_watch_receiver.compression.compression_gzip / 1d9151c131d0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-43bc16d9e145917018b444e2d6d34142bccc1a24fec69fd3697f04d508922897"></a>

## Next pages — aws_cloud_watch_receiver.compression.compression_gzip / 1d9151c131d0 / 4

- [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-42f1b43dd12dd17f4efd5bcfe819c7e609a0add4b9b08905fceede0296b705e1)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-832b7d0e9e7e12fa82228ce60fa1a842ca909cb0e16cc25e35830cc69bacd9b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-466537cdbdd20fe541d6872a00c0c4dfe772021dffe78c98cd1110842e08b6cc"></a>

## aws_cloud_watch_receiver.compression.compression_none — aws_cloud_watch_receiver.compression.compression_none / 4169697377ed / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [aws_cloud_watch_receiver](resources--global_log_receiver--reference--group-001.md#canonical-e925cb8c2d030ceefbf832b960fed2885dc2df3af2340b99002823aa4167e53f)
- [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-42f1b43dd12dd17f4efd5bcfe819c7e609a0add4b9b08905fceede0296b705e1)
- aws_cloud_watch_receiver.compression.compression_none

<a id="canonical-ada744d8fc958b8904ba0c64fd3fc6c931f4e913915f56ab047de938230c2365"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression none.

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
compression_none = {}
```

<a id="canonical-804cb77880ea12a81a3ba7609174baa613c8ffd3b2c69feabeb7e4726a66ffc2"></a>

## Direct properties — aws_cloud_watch_receiver.compression.compression_none / 4169697377ed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bdeedeadc79c474f633dbf24fc142473ca14cd814b4956d2b4d0d4bf78cb4b27"></a>

## Next pages — aws_cloud_watch_receiver.compression.compression_none / 4169697377ed / 4

- [aws_cloud_watch_receiver.compression](resources--global_log_receiver--reference--group-001.md#canonical-42f1b43dd12dd17f4efd5bcfe819c7e609a0add4b9b08905fceede0296b705e1)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-31bb2a7088509e10c5d9bc0ff4558a7cd09ca1905ff420ddd31ec84beba5c7f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df2b3b20473fe7848311664d0e9a4feec76020c75a3b05a0c443456a3e1c9d2d"></a>

## azure_event_hubs_receiver — azure_event_hubs_receiver / 1f3d2c9f4206 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- azure_event_hubs_receiver

<a id="canonical-b881e14c6dfa052502f651b8a706084c45f750454aca6f7e13d12cd08d9b0602"></a>

Type: `"object"`. single nested block, Optional.

Azure Event Hubs Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("instance",
    "namespace")}
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
azure_event_hubs_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-c990f21a3fd0252854867826ff60ff6fa14d333d534fd8a4f8520f77b0c4ed40"></a>

## Direct properties — azure_event_hubs_receiver / 1f3d2c9f4206 / 3

- [connection_string](resources--global_log_receiver--reference--group-001.md#canonical-b168542af544f8fe02d3e064e5a647e83192d10f7d7fb7192b9ec66f40447b9a): complete subsection reference.

<a id="canonical-cc5442fb91f8076600ab88b7f8c5b370f87c3fbc96f2c91f16f69e680a22408a"></a>

<a id="canonical-3d9f34eae171ee2787fa0461509323b4b81ca7056b6698d62e1ff90d153a0efa"></a>

## instance property — azure_event_hubs_receiver / 1f3d2c9f4206 / 4

Type: `"string"`. Optional.

Event Hubs Instance name into which logs should be stored.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  }
}
```

<a id="canonical-865d3df0f45c4d3abf9e45b080bfcca4e928133da851922c4658af8e4b18638a"></a>

<a id="canonical-70d1aed96ebf3fdd329574a2d8182fbe828700c9223e71ce258c78079b4587f8"></a>

## namespace property — azure_event_hubs_receiver / 1f3d2c9f4206 / 5

Type: `"string"`. Optional, Computed.

Event Hubs Namespace is namespace with instance into which logs should be stored.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 3,
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$",
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
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  }
}
```

<a id="canonical-6ada25a740e31a3b56ba528951e53f12568dd3d4377c1a39b6848352f64f8e6b"></a>

## Next pages — azure_event_hubs_receiver / 1f3d2c9f4206 / 6

- [azure_event_hubs_receiver.connection_string](resources--global_log_receiver--reference--group-001.md#canonical-b168542af544f8fe02d3e064e5a647e83192d10f7d7fb7192b9ec66f40447b9a)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-b168542af544f8fe02d3e064e5a647e83192d10f7d7fb7192b9ec66f40447b9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa3a313cca39af0d7f6bc56112400bcd9543b87a5bd29a59cf06ae094bc84e44"></a>

## azure_event_hubs_receiver.connection_string — azure_event_hubs_receiver.connection_string / ff04d83c24b4 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-31bb2a7088509e10c5d9bc0ff4558a7cd09ca1905ff420ddd31ec84beba5c7f6)
- azure_event_hubs_receiver.connection_string

<a id="canonical-3347b61e117b94dd49ae4dc57d072cd3f0d49878f500e9ccef4fb055652ed57b"></a>

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
connection_string {
  # Configure direct properties listed below.
}
```

<a id="canonical-caf79cb2041f95aeb4a3b9808c4b2a3634b895d6e0f0c3c01c39cbb1354b33ce"></a>

## Direct properties — azure_event_hubs_receiver.connection_string / ff04d83c24b4 / 3

- [blindfold_secret_info](resources--global_log_receiver--reference--group-001.md#canonical-08bb6c87699b92aff606c89e51091e997174352d0f2a649b444d956575fe8be3): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-001.md#canonical-705fb3983fd9d49644f8d14ec72d622702272fc458793810d7310947bf98a07b): complete subsection reference.

<a id="canonical-7c9a284463e9e4490fa00baf332baa4fcbf7a79823e921a2832df6ac34c2b2d6"></a>

## Next pages — azure_event_hubs_receiver.connection_string / ff04d83c24b4 / 4

- [azure_event_hubs_receiver.connection_string.blindfold_secret_info](resources--global_log_receiver--reference--group-001.md#canonical-08bb6c87699b92aff606c89e51091e997174352d0f2a649b444d956575fe8be3)
- [azure_event_hubs_receiver.connection_string.clear_secret_info](resources--global_log_receiver--reference--group-001.md#canonical-705fb3983fd9d49644f8d14ec72d622702272fc458793810d7310947bf98a07b)
- [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-31bb2a7088509e10c5d9bc0ff4558a7cd09ca1905ff420ddd31ec84beba5c7f6)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-08bb6c87699b92aff606c89e51091e997174352d0f2a649b444d956575fe8be3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-066a8d030b4c6a56d65294e0720fd06fd01b14c9cff942d7c1ad85feaf95da22"></a>

## azure_event_hubs_receiver.connection_string.blindfold_secret_info — azure_event_hubs_receiver.connection_string.blindfold_secret_info / 4c7cad21e220 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-31bb2a7088509e10c5d9bc0ff4558a7cd09ca1905ff420ddd31ec84beba5c7f6)
- [azure_event_hubs_receiver.connection_string](resources--global_log_receiver--reference--group-001.md#canonical-b168542af544f8fe02d3e064e5a647e83192d10f7d7fb7192b9ec66f40447b9a)
- azure_event_hubs_receiver.connection_string.blindfold_secret_info

<a id="canonical-8169490d77b230fdebc003ed6d66f174235ad488ce1548c6ad2c6f4e2711bf0d"></a>

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

<a id="canonical-423bd604f8709f4fc25daab6df974203e09cf32e7f990e9fb9751c646d74f55b"></a>

## Direct properties — azure_event_hubs_receiver.connection_string.blindfold_secret_info / 4c7cad21e220 / 3

<a id="canonical-ad4b5b90de81e1b2fe01105cc224a71f20083fd27f9748c4ae3daedb762f9313"></a>

<a id="canonical-44163ee69633445ff32241deeab1e39c3fbb3e9be08f5a0cfdb53dd45404caeb"></a>

## decryption_provider property — azure_event_hubs_receiver.connection_string.blindfold_secret_info / 4c7cad21e220 / 4

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

<a id="canonical-28f5711cca0d56a29c56f4bb37b391cd4cb57971c758e4f3295c8d4b3eedfd25"></a>

<a id="canonical-0c545aedef7ef053f9cac73f0aa2d29fb99203e66d4c44d93b3ca1bf41d75a17"></a>

## location property — azure_event_hubs_receiver.connection_string.blindfold_secret_info / 4c7cad21e220 / 5

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

<a id="canonical-4c9b19cf0626ccd5a142df9bc7f6d946efa038017727da65d08b18d4d0d61ba7"></a>

<a id="canonical-a4cf37d44c36db020bdb040d47edf24b6faf9d6eabdf18c35dfd3624021f5bc4"></a>

## store_provider property — azure_event_hubs_receiver.connection_string.blindfold_secret_info / 4c7cad21e220 / 6

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

<a id="canonical-290e9a75445206b21df2651ed5e0253e77506bd3d47bcbab7640831def406a0b"></a>

## Next pages — azure_event_hubs_receiver.connection_string.blindfold_secret_info / 4c7cad21e220 / 7

- [azure_event_hubs_receiver.connection_string](resources--global_log_receiver--reference--group-001.md#canonical-b168542af544f8fe02d3e064e5a647e83192d10f7d7fb7192b9ec66f40447b9a)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-705fb3983fd9d49644f8d14ec72d622702272fc458793810d7310947bf98a07b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f26972bce0076ef7b8093818407bed59cdccd08d3bb3cd648b6c0ab4a717be06"></a>

## azure_event_hubs_receiver.connection_string.clear_secret_info — azure_event_hubs_receiver.connection_string.clear_secret_info / 7167dee4795b / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [azure_event_hubs_receiver](resources--global_log_receiver--reference--group-001.md#canonical-31bb2a7088509e10c5d9bc0ff4558a7cd09ca1905ff420ddd31ec84beba5c7f6)
- [azure_event_hubs_receiver.connection_string](resources--global_log_receiver--reference--group-001.md#canonical-b168542af544f8fe02d3e064e5a647e83192d10f7d7fb7192b9ec66f40447b9a)
- azure_event_hubs_receiver.connection_string.clear_secret_info

<a id="canonical-5f30419505ebc4bac30ba30846302e5b4ba5a4c093998e2813cc797488b04348"></a>

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

<a id="canonical-4f97f47b04203a58a1519cf2892c44c128baa02110ddc918f31419fe358ab1f7"></a>

## Direct properties — azure_event_hubs_receiver.connection_string.clear_secret_info / 7167dee4795b / 3

<a id="canonical-f7108ad2eb564efa88b2ff162167e7b20429e27ebb9356f6b7f20818ab97b448"></a>

<a id="canonical-25d33b313e5afd012374248a52cb20a510cfe07c9254171d4a60544f4730f073"></a>

## provider_ref property — azure_event_hubs_receiver.connection_string.clear_secret_info / 7167dee4795b / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-37373ef071db4c17b283ee5005975c8d562cbdba5ee3b8d322e5f3a67eb79a73"></a>

<a id="canonical-98294fbd83fc8206c61f538854ef38d70a88159b4223fc1134b1a3e977306d25"></a>

## url property — azure_event_hubs_receiver.connection_string.clear_secret_info / 7167dee4795b / 5

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

<a id="canonical-55768a29942acaa0f47216d611ae4faa0788f8b71ee04f2d251b0de615923463"></a>

## Next pages — azure_event_hubs_receiver.connection_string.clear_secret_info / 7167dee4795b / 6

- [azure_event_hubs_receiver.connection_string](resources--global_log_receiver--reference--group-001.md#canonical-b168542af544f8fe02d3e064e5a647e83192d10f7d7fb7192b9ec66f40447b9a)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-d3cec99e849c282a720b0675fb155592bf8cb034e8e127ea487395b83385380d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d7a72fe562ac3c022d66b21c064c08cb9eda224a33ea4f5dbe027e661593152"></a>

## azure_receiver — azure_receiver / 9814ff9d63af / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- azure_receiver

<a id="canonical-c782dfc873188e384141ccd9e5104133b8539660526f46ac90e75f3272573d6d"></a>

Type: `"object"`. single nested block, Optional.

Azure Blob Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("container_name")}
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
azure_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-430d97a826ea0f2f4f3088a2a1b495445e8fedde948d64410e0a6d885c9199f2"></a>

## Direct properties — azure_receiver / 9814ff9d63af / 3

- [batch](resources--global_log_receiver--reference--group-001.md#canonical-36bbfd6dc535d03dcbd9233430f59c29bdc8952e7d2e39ba1f40b2b6827f6464): complete subsection reference.

- [compression](resources--global_log_receiver--reference--group-002.md#canonical-a6b4a288aef6e41730bf5863f09d93d85f0ea498742759fb417576d742339024): complete subsection reference.

- [connection_string](resources--global_log_receiver--reference--group-002.md#canonical-f968946f3bc02042a66321e9adca48e1d9d0d5c8158366f9983066c326dbb29d): complete subsection reference.

<a id="canonical-b34f7671b7c7395cf2c0b363e6f3c6406847ff10fd1bb759f2bf67887294b3fe"></a>

<a id="canonical-49be4d3dd2d845975e3b302bd90168f1cf7c2f384ae124d4f2965127cf4f6803"></a>

## container_name property — azure_receiver / 9814ff9d63af / 4

Type: `"string"`. Optional.

Container Name is the name of the container into which logs should be stored.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  }
}
```

- [filename_options](resources--global_log_receiver--reference--group-002.md#canonical-d434b60f67b137a3945ca4e11e352dce2afec3e1b803d2b3ef379084944d3101): complete subsection reference.

<a id="canonical-c8e5029e0a338afe1145a3c37642e303ed73cedfa53312161c06021275695889"></a>

## Next pages — azure_receiver / 9814ff9d63af / 5

- [azure_receiver.batch](resources--global_log_receiver--reference--group-001.md#canonical-36bbfd6dc535d03dcbd9233430f59c29bdc8952e7d2e39ba1f40b2b6827f6464)
- [azure_receiver.compression](resources--global_log_receiver--reference--group-002.md#canonical-a6b4a288aef6e41730bf5863f09d93d85f0ea498742759fb417576d742339024)
- [azure_receiver.connection_string](resources--global_log_receiver--reference--group-002.md#canonical-f968946f3bc02042a66321e9adca48e1d9d0d5c8158366f9983066c326dbb29d)
- [azure_receiver.filename_options](resources--global_log_receiver--reference--group-002.md#canonical-d434b60f67b137a3945ca4e11e352dce2afec3e1b803d2b3ef379084944d3101)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-36bbfd6dc535d03dcbd9233430f59c29bdc8952e7d2e39ba1f40b2b6827f6464"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-547a100aeaa2a67c43862961665595670da6b13f36c04666c8a42e088b1cc71d"></a>

## azure_receiver.batch — azure_receiver.batch / 4e5b1bea2748 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-7af3c00623500b670d8cb63b589e72d65b2843d1315a7848a75cd621def04808)
- [azure_receiver](resources--global_log_receiver--reference--group-001.md#canonical-d3cec99e849c282a720b0675fb155592bf8cb034e8e127ea487395b83385380d)
- azure_receiver.batch

<a id="canonical-0622d5592358bc82e1aebc69146a12f80fd5e811d9af21ee7aaca92a03b54fcc"></a>

Type: `"object"`. single nested block, Optional.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_bytes",
    "max_bytes_disabled"),
  validators.ConflictingObjectAttributes("max_events",
    "max_events_disabled"),
  validators.ConflictingObjectAttributes("timeout_seconds",
    "timeout_seconds_default")}
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
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

Terraform syntax:

```terraform
batch {
  # Configure direct properties listed below.
}
```

<a id="canonical-4ab7d9eb1d0dfc1b5d107e68ce0f0cff5d7e15376ac98c4bc639783f53798006"></a>

## Direct properties — azure_receiver.batch / 4e5b1bea2748 / 3

<a id="canonical-efc2c0d542bb962658778f2cac1600283414dd674e9e6346cec72aa54f160483"></a>

<a id="canonical-ff1b22ed9996549e86a721c7b08fab4c6ea1b436440e01d75a35c9b636a48114"></a>

## max_bytes property — azure_receiver.batch / 4e5b1bea2748 / 4

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(4096, 10485760),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](resources--global_log_receiver--reference--group-002.md#canonical-bbbac1e7926f169f0485535b15c5334c4dad4d60f2323e580f5f4b5b1bb65244): complete subsection reference.

<a id="canonical-8b7e16cca09879e6daff0195a84303e60dcf0459d6c85480b62fcd782713c57d"></a>

<a id="canonical-03b6eb347cbe69ef017b5455061f4aa4786dec5a7146da48234b2161991c5b7c"></a>

## max_events property — azure_receiver.batch / 4e5b1bea2748 / 5

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(32, 2000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](resources--global_log_receiver--reference--group-002.md#canonical-e27004711cc06faa10778160948f9073c17102323c76f6c8cc50d0927f97f79f): complete subsection reference.

<a id="canonical-5f83da3dba00fb7192597acf9a42d097ffcd689ba60d3696ed9f619448625e79"></a>

<a id="canonical-21c9ce1baf47ad35e2b677228eb8b88ce53766a38a4a03d4b2b0a07941dfd193"></a>

## timeout_seconds property — azure_receiver.batch / 4e5b1bea2748 / 6

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](resources--global_log_receiver--reference--group-002.md#canonical-be8f6114d4ad0d6c8612b75b777f46e6fcf1fd8c2e05ee5edcd47149a902593e): complete subsection reference.
