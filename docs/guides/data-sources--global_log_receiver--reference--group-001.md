---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71ca6054badd9ed308aa917c6cc36d4d17d1aea50eff9172e68c0a6c9a0ba815"></a>

## Property reference — Property reference / ee2aefa3af0a / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- Property reference

<a id="canonical-b8949203fc5e4ae17dd6dcab6974bfad8d4c74a9a7cc94691294c17842834a36"></a>

## Direct properties — Property reference / ee2aefa3af0a / 3

<a id="canonical-08170eeb7504824b8c9af262ebefe5fd84b970006ad2153180221346c7e14248"></a>

<a id="canonical-515477c1e099be1e9f803df89d83f154fdf83716b0153d65d2cccf41cef144ca"></a>

## annotations property — Property reference / ee2aefa3af0a / 4

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

- [audit_logs](data-sources--global_log_receiver--reference--group-001.md#canonical-df9883b23b9a3363e082d3c4b3d4c306c14b6778efb738cf51239d9fa1e7fbfb): complete subsection reference.

- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e): complete subsection reference.

- [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1e93724fa97f305e8285ab20a3de95bcda72cb3dd4afba2b5d1f659e5b82aa39): complete subsection reference.

- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6): complete subsection reference.

- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052): complete subsection reference.

<a id="canonical-64743360726e0dc5f3e15b484cd52c6ea168c0516ce2f4023f136aed2f13cc97"></a>

<a id="canonical-c5d58d1dec89933d0abbc981811de5f579e92bdd0bf20f8da77c3ea15585a87b"></a>

## description property — Property reference / ee2aefa3af0a / 5

Type: `"string"`. Computed.

Description of the GlobalLogReceiver.

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

- [dns_logs](data-sources--global_log_receiver--reference--group-002.md#canonical-b7191829d34da96f0435f5aec41d4e561e626957b75cdf343f72630ff0e93240): complete subsection reference.

- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2): complete subsection reference.

- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8): complete subsection reference.

<a id="canonical-a056c0c2a1b6d69f71077a4e06f64f2f7300bac72b7367ef24d8a1f158ecff39"></a>

<a id="canonical-e4792d68bdf6be121ac40161198d1980136a3fccc2f68f3835b30b23a0626da9"></a>

## id property — Property reference / ee2aefa3af0a / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd): complete subsection reference.

<a id="canonical-94995592684160b7dab0af8037ad0a3d507fae58b962efb8ad5de382db15f243"></a>

<a id="canonical-266e81d781681ace70522e9cb831d7b0d638a9cf505c6a0c0c355afe6b6ede4d"></a>

## labels property — Property reference / ee2aefa3af0a / 7

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

<a id="canonical-007760a76f5a884982d2dbbcac8bc948eaa6b9f7d166fade141e16d3e1012f0b"></a>

<a id="canonical-7d48f8550451076dc27993ab85628f73dcd73ba7744457187ad3745c0163c808"></a>

## name property — Property reference / ee2aefa3af0a / 8

Type: `"string"`. Required.

Name of the GlobalLogReceiver.

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

<a id="canonical-4584b30ea2792d16597d6c62636f22527cfd57472f11f1e00887d0dfe0b066d9"></a>

<a id="canonical-887b5abc919e2794932749b4521082f48b7fe259f853cdf7695ffc20a7d2d965"></a>

## namespace property — Property reference / ee2aefa3af0a / 9

Type: `"string"`. Required.

Namespace where the GlobalLogReceiver exists.

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

- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-6a8ca15a6755e8c761d98fa0efc7d887c67f7c2cf83ede5ae159146cee36621c): complete subsection reference.

- [ns_all](data-sources--global_log_receiver--reference--group-003.md#canonical-eac6c2cc015966b8d0a1b560e162a5757c3992ef61ab6e3555f7a8641fdb1d88): complete subsection reference.

- [ns_current](data-sources--global_log_receiver--reference--group-003.md#canonical-e7d88a8682a21bc20dc8d78ab563c10f0d3a2ee20a4b1399394c1bda0161b874): complete subsection reference.

- [ns_list](data-sources--global_log_receiver--reference--group-003.md#canonical-f360f7b122a79b7d1c6f6f9d22020aaec9eb1157401d37099f911acc32f3529a): complete subsection reference.

- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3): complete subsection reference.

- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-78587f316fc037494d3bdda63e54544614ce095432c2e1335fabfd6ee00f891b): complete subsection reference.

- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2): complete subsection reference.

- [security_events](data-sources--global_log_receiver--reference--group-004.md#canonical-4234f3746f0500e78349b4be62bcfa5b49542d6219128d146579072b81afbd0b): complete subsection reference.

- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070): complete subsection reference.

- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-60223cb26da3c4491af2cd9271ba7071122485608016a658e918ce37e12dc0c1): complete subsection reference.

<a id="canonical-56b48faa70a80de8420c6df73ed69e8079aa0017ab60e275677f6d2bd6f4b80f"></a>

## All schema paths — Property reference / ee2aefa3af0a / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--global_log_receiver--reference--group-001.md#canonical-08170eeb7504824b8c9af262ebefe5fd84b970006ad2153180221346c7e14248) |
| `audit_logs` | [audit_logs](data-sources--global_log_receiver--reference--group-001.md#canonical-9fa6306bdbb6069dd4b9aa668f4e2c661e6f3e04d5ff23da549dd8b195a9df42) |
| `aws_cloud_watch_receiver` | [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-fd5167464041e5466356ca926c5af5970183f82d36c7bb3345d6187ab1da0f97) |
| `aws_cloud_watch_receiver.aws_cred` | [aws_cloud_watch_receiver.aws_cred](data-sources--global_log_receiver--reference--group-001.md#canonical-d94456f96c37418dc58a53822fcb1d626e6b35b5ae90f4453180333e702e552e) |
| `aws_cloud_watch_receiver.aws_cred.name` | [aws_cloud_watch_receiver.aws_cred.name](data-sources--global_log_receiver--reference--group-001.md#canonical-22729212426288499175d384f069c8e3bf45966b7eed9e2008006fecef1e2e29) |
| `aws_cloud_watch_receiver.aws_cred.namespace` | [aws_cloud_watch_receiver.aws_cred.namespace](data-sources--global_log_receiver--reference--group-001.md#canonical-539553a479f51cf160ef6106331fc18c42a8ee71510475bb1eb9ee59d48e8c0f) |
| `aws_cloud_watch_receiver.aws_cred.tenant` | [aws_cloud_watch_receiver.aws_cred.tenant](data-sources--global_log_receiver--reference--group-001.md#canonical-c30db574f1137de546b01b378d382eb0de325dc899055896defd79173af7ce21) |
| `aws_cloud_watch_receiver.aws_region` | [aws_cloud_watch_receiver.aws_region](data-sources--global_log_receiver--reference--group-001.md#canonical-f54a0193cece959c174850819f25161e5410ea3af613e40f78614f5d6f9b2659) |
| `aws_cloud_watch_receiver.batch` | [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-6ff77b87244ade51af7ac58c52492e902432ddba9d0927829908e60d171a1d9c) |
| `aws_cloud_watch_receiver.batch.max_bytes` | [aws_cloud_watch_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-001.md#canonical-d5ab853c25bd71f11cfbc5ff3e4d2c69d73b77ab03ae7e9dff37063d1fdccd1f) |
| `aws_cloud_watch_receiver.batch.max_bytes_disabled` | [aws_cloud_watch_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-743ebb45bed6ad9279e9ffa0ff45b6c064f171e36f36b6114491fab958711eb9) |
| `aws_cloud_watch_receiver.batch.max_events` | [aws_cloud_watch_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-001.md#canonical-9c08e76aa58273d5ae1342787410951c6ed55b84b57629cbc279153dc096a8e8) |
| `aws_cloud_watch_receiver.batch.max_events_disabled` | [aws_cloud_watch_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-b252eb40bbab88586206c8b7cf7aaada46f72f4f2c09786b3f78fdd09720d9ee) |
| `aws_cloud_watch_receiver.batch.timeout_seconds` | [aws_cloud_watch_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-001.md#canonical-08972e6997e67a014c0422cb6b58e8984fb0323a802871ffd965351e972d9434) |
| `aws_cloud_watch_receiver.batch.timeout_seconds_default` | [aws_cloud_watch_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-001.md#canonical-b6625abf8ed7349dc06839f48b387e7ddd6795d11d592a2da20c5f358d79c690) |
| `aws_cloud_watch_receiver.compression` | [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-9465830781e13d9893dfb02b396655b711827d008d8a17992b98f1543ed52d87) |
| `aws_cloud_watch_receiver.compression.compression_default` | [aws_cloud_watch_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-001.md#canonical-4c3ba4ea5fd5f580b3a4baa08ba4dbf8bdc2c1c0d783c71ea689d95bd36e67f2) |
| `aws_cloud_watch_receiver.compression.compression_gzip` | [aws_cloud_watch_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-001.md#canonical-6bae98de3de949c0e2b45aefc1052ef76756895fd1de86a3ef46c7cbe668520c) |
| `aws_cloud_watch_receiver.compression.compression_none` | [aws_cloud_watch_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-001.md#canonical-06e14a5887e30c610536c4d48ca5d6527e2f4beddde77f4d624adbe4633b7cbe) |
| `aws_cloud_watch_receiver.group_name` | [aws_cloud_watch_receiver.group_name](data-sources--global_log_receiver--reference--group-001.md#canonical-07faf075fb4c52ffb782b86bc7f8dbaafcfe53fd28604ab373a8b19310a58272) |
| `aws_cloud_watch_receiver.stream_name` | [aws_cloud_watch_receiver.stream_name](data-sources--global_log_receiver--reference--group-001.md#canonical-4afd4f5b4d1013473e50d6aa425922047cd4286571d67a984e1cd2e833f8c273) |
| `azure_event_hubs_receiver` | [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-5f8d310c2746fa4c17c4a70016a2a5fe97191aed64d1cd597a98b1a3d89f2fee) |
| `azure_event_hubs_receiver.connection_string` | [azure_event_hubs_receiver.connection_string](data-sources--global_log_receiver--reference--group-001.md#canonical-3f4267535f141fd1a61562c1e0281c646018d972811f37bc7fc41881d1335b8e) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info](data-sources--global_log_receiver--reference--group-001.md#canonical-6f67f1ba265aaabb7c65c4e12e588552289d7a681048453950c98228d176f93c) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.decryption_provider` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-001.md#canonical-9d3fb97a20e0287fc8733cce7e99b615b43888071665c2e5ab918eb1c845531f) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.location` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-001.md#canonical-a96d2751c331a5b27a2d669e88f3e65f7e58faaead83ad3d50b7ab051d4cb5de) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.store_provider` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-001.md#canonical-2286724c776f3ccadca1619b723ae477ebb57504a0d5634a76282a70ef2b54af) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info` | [azure_event_hubs_receiver.connection_string.clear_secret_info](data-sources--global_log_receiver--reference--group-001.md#canonical-5cbf8a4e79f70be17a7fa9697fa49e4424bbf3f87780af9aeae2cb93063240fe) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info.provider_ref` | [azure_event_hubs_receiver.connection_string.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-001.md#canonical-fd3c0d07adece5898726453ec277bee236d4738668cda42d54446fca120defd4) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info.url` | [azure_event_hubs_receiver.connection_string.clear_secret_info.url](data-sources--global_log_receiver--reference--group-001.md#canonical-937ccedaf14e3024919b28bf2f3dbb59916b5bfb0f12d8f2b2c5fafa3a8e2867) |
| `azure_event_hubs_receiver.instance` | [azure_event_hubs_receiver.instance](data-sources--global_log_receiver--reference--group-001.md#canonical-ef0d165096b5060e460a8e50184d826db12d6d88bfb67c322554a23b8fa4f960) |
| `azure_event_hubs_receiver.namespace` | [azure_event_hubs_receiver.namespace](data-sources--global_log_receiver--reference--group-001.md#canonical-e6190f350141c41cfbe65991353ebc370fa046012b133b20c6a779b96aba5f0a) |
| `azure_receiver` | [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-33d0516820774d65b0cb95a87591491c90c8a78ee4996ddde7ec41aba356117d) |
| `azure_receiver.batch` | [azure_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-bba588ffa53013eb0467f505e13897fadcf8f3f2ba3467b76226c095abdfc527) |
| `azure_receiver.batch.max_bytes` | [azure_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-001.md#canonical-ba331a499368037efaf4c6958b82d55fbc9fd9462f2c62c6a9eb202f1861b076) |
| `azure_receiver.batch.max_bytes_disabled` | [azure_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-ece2b962d6b190267bd988002a60311ca84c5613e7553673a1ac20610dc9f719) |
| `azure_receiver.batch.max_events` | [azure_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-001.md#canonical-a81532423032e39a56057c03dd4bddcdf357669c88aea1ce8c81a10f303687a6) |
| `azure_receiver.batch.max_events_disabled` | [azure_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-22db6b551a8506d2080636f35d98a33da53e95858c5d6c811af3c13c551e15d6) |
| `azure_receiver.batch.timeout_seconds` | [azure_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-001.md#canonical-d43735d6688c92d67acce19fa8432bfd9a8e99cf5ef24b6e6b0a9ba8bd4566dc) |
| `azure_receiver.batch.timeout_seconds_default` | [azure_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-5b31381aef3de6031f6b10cff74861e476d214c0f187a72adf69e3e94d69dc8b) |
| `azure_receiver.compression` | [azure_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-d2738f44957da101adde93ab23d91ca58bc172fb4c8798282b46621f00d67597) |
| `azure_receiver.compression.compression_default` | [azure_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-002.md#canonical-086288c086c32802a48bd111599752ee62c14655880db336654c3d9cb03f2435) |
| `azure_receiver.compression.compression_gzip` | [azure_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-002.md#canonical-8f0646fe1ad8c33548e76b5dc1943a7a82609c23691e5d80761e32777e035288) |
| `azure_receiver.compression.compression_none` | [azure_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-002.md#canonical-c1cf177ac80c70e866c17182bf34a06edea969f5c1b9915e34b7c3af97a0e816) |
| `azure_receiver.connection_string` | [azure_receiver.connection_string](data-sources--global_log_receiver--reference--group-002.md#canonical-9a8ab2b67481917b6269a6805c892ff08b34454c6553603e9b1d289629620eb7) |
| `azure_receiver.connection_string.blindfold_secret_info` | [azure_receiver.connection_string.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-9b62ae478506050f71f7fc7bc9d3e4408d7ae700e0162609e71ef22e1e22a85a) |
| `azure_receiver.connection_string.blindfold_secret_info.decryption_provider` | [azure_receiver.connection_string.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-a744ebc9550404dd61dcea5bf66772b90ce57d076e11589fa435dcf40978b101) |
| `azure_receiver.connection_string.blindfold_secret_info.location` | [azure_receiver.connection_string.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-002.md#canonical-a9a97fd5f79a84523a0049a312619d06f06c3de7bd0f28b6b6e8fe2e0ced3568) |
| `azure_receiver.connection_string.blindfold_secret_info.store_provider` | [azure_receiver.connection_string.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-54df62adb361edbeec6fecd7b0a9b62a6544759b6d6dc0e8516862672955919d) |
| `azure_receiver.connection_string.clear_secret_info` | [azure_receiver.connection_string.clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-926c7a50ad8d0a896ce6777b5734de431fb118c5ef72ed0ec9a97401d8959e1e) |
| `azure_receiver.connection_string.clear_secret_info.provider_ref` | [azure_receiver.connection_string.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-002.md#canonical-facb44a29e39230970a32f42c195dcb032350f5dcbc3b042b2b94bf0cfa39dd5) |
| `azure_receiver.connection_string.clear_secret_info.url` | [azure_receiver.connection_string.clear_secret_info.url](data-sources--global_log_receiver--reference--group-002.md#canonical-8dad2fa3a7256b4d1522a2a7b83b4d9bcb65269c7be85f4acb018dd0f050062c) |
| `azure_receiver.container_name` | [azure_receiver.container_name](data-sources--global_log_receiver--reference--group-001.md#canonical-5a682b0e414665ce244919459fa07d09a594838086df125e77e5fefa822e8304) |
| `azure_receiver.filename_options` | [azure_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-d3e22600ed554da8cfaaea2fc08b5a052ac2a723867835c0276889a548aa8000) |
| `azure_receiver.filename_options.custom_folder` | [azure_receiver.filename_options.custom_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-25913b195b678e1c346147d6998f11b4fc7b93dbd4dde9c079d92e0bfb5610fd) |
| `azure_receiver.filename_options.log_type_folder` | [azure_receiver.filename_options.log_type_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-8b206a92d0a96a4a933c67f6ccdb4eabc53984ce64e88b3164cc1a3fa4e4cc23) |
| `azure_receiver.filename_options.no_folder` | [azure_receiver.filename_options.no_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-a5c03b4b29286e10a758840756f91c035ba813671cbe3b7bcae84faadcf0cb80) |
| `datadog_receiver` | [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-4c6892ae68db2e1b186f351c385d1019fefce792671cb3df607e0554ec2ddc23) |
| `datadog_receiver.batch` | [datadog_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-8ecf8fe844568640d846280cb3a0ce7dbc084d5104a8e3e9f01c9c2eb5c70b15) |
| `datadog_receiver.batch.max_bytes` | [datadog_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-002.md#canonical-a18701e55c6e13a2cf1dd7e6ad1ef31c0644cc1a3d85fb083143d0621d111ede) |
| `datadog_receiver.batch.max_bytes_disabled` | [datadog_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-30637618a2a579faf2ba0f1de44caa34b7520f55b744e1b15ea7fcf1d94c89bc) |
| `datadog_receiver.batch.max_events` | [datadog_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-002.md#canonical-cdcd43eec8d6b858a9ebd27e5440b0367156e09e3907460598279755ee3195b3) |
| `datadog_receiver.batch.max_events_disabled` | [datadog_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-63b7bab2b36f2ac5fb91b3c4bd94e4ccd37c11789152f2c059e8395ccfc3ff1d) |
| `datadog_receiver.batch.timeout_seconds` | [datadog_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-002.md#canonical-59ddef35fd30dcd9c2b1244eedfd147c7486e94fde5ddef5df41e4d6d516befc) |
| `datadog_receiver.batch.timeout_seconds_default` | [datadog_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-b55f689a0c07b6da9a2a30289d24e56333b3c8cf2e58b81e9deea3db36a4062f) |
| `datadog_receiver.compression` | [datadog_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-951ad4f64255801c1bbe1236d09be1aa809874f2dd5455e4bf0391f6610c9cb9) |
| `datadog_receiver.compression.compression_default` | [datadog_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-002.md#canonical-e43da377baf4adaed52f9bf8799918b4fb6d37bf56dcb4451ec694f157175bfb) |
| `datadog_receiver.compression.compression_gzip` | [datadog_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-002.md#canonical-43dd18c502649a825da456d90cfc1fb475675c4df863728a82bbc1978029da4a) |
| `datadog_receiver.compression.compression_none` | [datadog_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-002.md#canonical-8e9cf508f57f3d052f6a4a5c3ec9269e35d1e246edd34202dfb8bf3d19b75155) |
| `datadog_receiver.datadog_api_key` | [datadog_receiver.datadog_api_key](data-sources--global_log_receiver--reference--group-002.md#canonical-3b4764f72c0139bb246e25e69b1c49172d79a115dc699afa8aebc35d23e51dd6) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info` | [datadog_receiver.datadog_api_key.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-fe83c79b3b5a2a87778a0bba976b8aa21b2371cc049fa2c9e43b45cc456fc56c) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.decryption_provider` | [datadog_receiver.datadog_api_key.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-06793d20358936e63313d7c3b3d11fd1d791f28c903e376e3d0f094d3e796802) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.location` | [datadog_receiver.datadog_api_key.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-002.md#canonical-193837d52a6573f6f2c01a654e754e2d1635f3232757d415ab49d03667890bb0) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.store_provider` | [datadog_receiver.datadog_api_key.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-3d12e1b7d780d698f1d7ab07ce2c4f2839d35edc301a38d1507f0b69af58d952) |
| `datadog_receiver.datadog_api_key.clear_secret_info` | [datadog_receiver.datadog_api_key.clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-1f9417400ee0546c88a35a96ab86efb120b36233c50c0e8175278996ff480cf3) |
| `datadog_receiver.datadog_api_key.clear_secret_info.provider_ref` | [datadog_receiver.datadog_api_key.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-002.md#canonical-0988e20d2c700e9c4bb198c13b22c4508c21cd40ea55d677a731936c9da75348) |
| `datadog_receiver.datadog_api_key.clear_secret_info.url` | [datadog_receiver.datadog_api_key.clear_secret_info.url](data-sources--global_log_receiver--reference--group-002.md#canonical-18a53425696b0cb46dff9524c18a00c376e9e13af0af5f9847a54c5224a6a0a7) |
| `datadog_receiver.endpoint` | [datadog_receiver.endpoint](data-sources--global_log_receiver--reference--group-002.md#canonical-32e4c9055a23446d8b7669634b242c92d7193d89e342b97e1af0e9d033c63647) |
| `datadog_receiver.no_tls` | [datadog_receiver.no_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-8bddf8deae3d8010a00b52e06e73acd601fb8b978ce18652636dd280d8c4ce39) |
| `datadog_receiver.site` | [datadog_receiver.site](data-sources--global_log_receiver--reference--group-002.md#canonical-5f00d9798ce5f10dfa4fcda10cae8bf619585e963754f81cf4e3deda5aa2f64e) |
| `datadog_receiver.use_tls` | [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-d3ec2a9b3a86cef3bbc3253d1e4c17ebc021362047bf28e4ec948704a41f95c7) |
| `datadog_receiver.use_tls.disable_verify_certificate` | [datadog_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-002.md#canonical-e51991fd316ad751946194cb45351a404124461893fc381578495a9af12b5e2b) |
| `datadog_receiver.use_tls.disable_verify_hostname` | [datadog_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-002.md#canonical-bdfae2de990360c37d6d31230df3b3830d0d16a1156b53d9d55db22267216a45) |
| `datadog_receiver.use_tls.enable_verify_certificate` | [datadog_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-002.md#canonical-9773376843b913e9e41565bc24bfc07ba2da6ac70f1f85f3ea36a4bad418ef09) |
| `datadog_receiver.use_tls.enable_verify_hostname` | [datadog_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-002.md#canonical-a5921b2bc139d637ffeb4827cf7088024989142b79824731724c9c0d9026667a) |
| `datadog_receiver.use_tls.mtls_disabled` | [datadog_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-5ee0eefaf8ab38af5e1f102d5d789c2bb55f42f945de35e6c7b6f2bffde61589) |
| `datadog_receiver.use_tls.mtls_enable` | [datadog_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-002.md#canonical-aa4734897cfe92c5c9c6617dd19a7263c59f3ffca3863e76cd2c598702b07ca4) |
| `datadog_receiver.use_tls.mtls_enable.certificate` | [datadog_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--reference--group-002.md#canonical-14dbf16c960950556a755cb0f85d2f10a2b1e418117b46617ab0e9b3bb44b206) |
| `datadog_receiver.use_tls.mtls_enable.key_url` | [datadog_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-002.md#canonical-b17a03c634b6cb250c4c6730fa7f8be7d179fa2e3fae5975161876bd6783a656) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-fbea8ff3604339cd1032d95248e8701d7c86f391b586318e8260c714c5503238) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-71666d17008ac5c410cced558c067a2a6a8c48c069c110af97d7d632b85d3a47) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-002.md#canonical-bd8ff3c7bb88e5b7cc55874a3e8dd18130e5204ddd02baad0dcd034f574a717a) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-dcf6b54e5941ededbfcbe14c0df740a5364e9534896d42abd811a135ee509928) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-817eb7fe5c07b09131147f957abab9349fa67fed0733f3c727705cee91673d2c) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-002.md#canonical-57e2ab3b560dfad4133e8a72ee6fb2229bb48d6918671c98d9557b9b25798b4c) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--reference--group-002.md#canonical-97a88cb14e10cd4fbfc91d0861e9adb9bc816979613e5219ae37ccfb3da97776) |
| `datadog_receiver.use_tls.no_ca` | [datadog_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-002.md#canonical-8a89e65daa4a2c84b0b05f9f32c5d7f0c0ac0edb524e3c6f4a47b6b4947d77d1) |
| `datadog_receiver.use_tls.trusted_ca_url` | [datadog_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--reference--group-002.md#canonical-f079ce660e3481ad4433bc031d11254834dee7b14a1a0eb759b9ae9841279d65) |
| `description` | [description](data-sources--global_log_receiver--reference--group-001.md#canonical-64743360726e0dc5f3e15b484cd52c6ea168c0516ce2f4023f136aed2f13cc97) |
| `dns_logs` | [dns_logs](data-sources--global_log_receiver--reference--group-002.md#canonical-3c0cf4cfae7a37e36a098964bdb7ac68726ba850bef0307a663b00b39428e737) |
| `gcp_bucket_receiver` | [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1128734e088b037c58c9603f737bff23c86df51c827cfc4f1932c9f3cf5ed364) |
| `gcp_bucket_receiver.batch` | [gcp_bucket_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-b60b7fc7098719dcbb05bd18ccf5b42822e76a931101db459cea2c6c927d3a99) |
| `gcp_bucket_receiver.batch.max_bytes` | [gcp_bucket_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-002.md#canonical-81ad38888e85f02ca1447590b568cf903a758c5464a0e1ba37a1e39916dc59fc) |
| `gcp_bucket_receiver.batch.max_bytes_disabled` | [gcp_bucket_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-edd512431d149cc6d8d74a94071804e1c087f519285a4fd4da99a727efda1084) |
| `gcp_bucket_receiver.batch.max_events` | [gcp_bucket_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-002.md#canonical-6b9291527985f632a629818a6d9e9bc7f3ec23ff57533fb170773a5d3d8636ba) |
| `gcp_bucket_receiver.batch.max_events_disabled` | [gcp_bucket_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-41278e2ff25ec6ccef86e8017b4964126ce35de5c548450e72cf852a3cb608c3) |
| `gcp_bucket_receiver.batch.timeout_seconds` | [gcp_bucket_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-002.md#canonical-c4ef0f3ed46534b17bda510943e96532d5971cd421dc898e97bb2d4cccf1f603) |
| `gcp_bucket_receiver.batch.timeout_seconds_default` | [gcp_bucket_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-22434fc273c739d43c31bc44965c48c9cfb522ffd013ed5ca75d4bdeb9aeab73) |
| `gcp_bucket_receiver.bucket` | [gcp_bucket_receiver.bucket](data-sources--global_log_receiver--reference--group-002.md#canonical-c6194b69f5e744f8b3d20234e8adc6281e06ba05e9ee2ac5fc7e84896cf86492) |
| `gcp_bucket_receiver.compression` | [gcp_bucket_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-fbfc0f6fb8c878954065cdeb459f311f202261ae70b9fe876789f9bf680d0705) |
| `gcp_bucket_receiver.compression.compression_default` | [gcp_bucket_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-002.md#canonical-847f5fd5904cda37f611bf0a290cff3ed2c5fbac75eb3d96739cdacd63ad8e4f) |
| `gcp_bucket_receiver.compression.compression_gzip` | [gcp_bucket_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-002.md#canonical-0f5423887de056959545ed6540659231afd490d2f4bdb9436b9a502dc6e7da4f) |
| `gcp_bucket_receiver.compression.compression_none` | [gcp_bucket_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-002.md#canonical-415d117efdb1a3cad59447afb8489555eeaaf24a6771f91e970bda40180b2d3c) |
| `gcp_bucket_receiver.filename_options` | [gcp_bucket_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-f4e52b63f489e36499eedc242691bf4039e4a60219e5760737085866128630d1) |
| `gcp_bucket_receiver.filename_options.custom_folder` | [gcp_bucket_receiver.filename_options.custom_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-7df8ea6a557623f4bde559320187ac7a4f50ded1d98899e53f46dc18ecc5bdd6) |
| `gcp_bucket_receiver.filename_options.log_type_folder` | [gcp_bucket_receiver.filename_options.log_type_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-821ba9dc5ad1754203bb7e8ee30b8b48c8a2e83d29ec1e31bd644a42ec88e02d) |
| `gcp_bucket_receiver.filename_options.no_folder` | [gcp_bucket_receiver.filename_options.no_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-4bd851431589b06ac623a71a18fe0eb7c56c3e4393b41f0f3fa0200c677b0299) |
| `gcp_bucket_receiver.gcp_cred` | [gcp_bucket_receiver.gcp_cred](data-sources--global_log_receiver--reference--group-002.md#canonical-aa836cb86a1ce4924fdd598a108474605169187fdf189a89e3a72fcaae3c33a9) |
| `gcp_bucket_receiver.gcp_cred.name` | [gcp_bucket_receiver.gcp_cred.name](data-sources--global_log_receiver--reference--group-002.md#canonical-1aa6ded0200a2f42c76148303d6419283bf2a97c68afb62fb20a7f76a2f483d3) |
| `gcp_bucket_receiver.gcp_cred.namespace` | [gcp_bucket_receiver.gcp_cred.namespace](data-sources--global_log_receiver--reference--group-002.md#canonical-4e867ba467d430eb854764492612e68ba4a413d36b6baf1eb0114ca7debebdae) |
| `gcp_bucket_receiver.gcp_cred.tenant` | [gcp_bucket_receiver.gcp_cred.tenant](data-sources--global_log_receiver--reference--group-002.md#canonical-11470f0aaf7da9afce507a3a7512eff548a9f896365a6de14b2a0fac0636acd9) |
| `http_receiver` | [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1128fb5d3d174ae457010bed22d2085ce139b22d5a8cc9a82b46392b50accfab) |
| `http_receiver.auth_basic` | [http_receiver.auth_basic](data-sources--global_log_receiver--reference--group-002.md#canonical-58acaf9a0fbd810a4eb6afb7f9dd39d3156cd8752b7e8d7eea1666f869a922ee) |
| `http_receiver.auth_basic.password` | [http_receiver.auth_basic.password](data-sources--global_log_receiver--reference--group-002.md#canonical-c0125065ecf75fa9f1dddc29d858199d093e8dbf998aa8adb276ea2766de3d81) |
| `http_receiver.auth_basic.password.blindfold_secret_info` | [http_receiver.auth_basic.password.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-21cb144e32d156083fb56db24f7dc3213ebc85c2dbb50ea2105abf966f38d7c5) |
| `http_receiver.auth_basic.password.blindfold_secret_info.decryption_provider` | [http_receiver.auth_basic.password.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-3c52ed6dcb408ff9a11ce0a703f3ecd53b22ca4457331178d14522aa06cb2cd0) |
| `http_receiver.auth_basic.password.blindfold_secret_info.location` | [http_receiver.auth_basic.password.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-002.md#canonical-5089fb9a7f79ab2b14afa16b30872727260979e05cf65eee9fb6a0dbac6fe62f) |
| `http_receiver.auth_basic.password.blindfold_secret_info.store_provider` | [http_receiver.auth_basic.password.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-caea25fba2ebbd7984ee329cb830ecf71145a968e0cb59280128314a90a95132) |
| `http_receiver.auth_basic.password.clear_secret_info` | [http_receiver.auth_basic.password.clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-0760a48cc9e4330a5ef077135d96fd56583d3308cd101729e9e534c0b4953135) |
| `http_receiver.auth_basic.password.clear_secret_info.provider_ref` | [http_receiver.auth_basic.password.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-002.md#canonical-b7dc171f74029b26a6eea8218796d2f55d42f2a4e81674506858516e78d724fd) |
| `http_receiver.auth_basic.password.clear_secret_info.url` | [http_receiver.auth_basic.password.clear_secret_info.url](data-sources--global_log_receiver--reference--group-002.md#canonical-6ad8789e73c15e6f8d1a77233995b779b0f924bbd301bfba05895af0f3080e11) |
| `http_receiver.auth_basic.user_name` | [http_receiver.auth_basic.user_name](data-sources--global_log_receiver--reference--group-002.md#canonical-3cc30d5948f55a5dc362c3702ad64e11970e515ebb389ae15e90bbef5c9dca6f) |
| `http_receiver.auth_none` | [http_receiver.auth_none](data-sources--global_log_receiver--reference--group-002.md#canonical-01513ad57d8235f3bfebd49bfca8fba002802cff8dc3fb29186ab4cb5dd58a25) |
| `http_receiver.auth_token` | [http_receiver.auth_token](data-sources--global_log_receiver--reference--group-002.md#canonical-a30225a85f2219910dbf42f18d8ab75bbe92ee9bf83086659cc65a94d26a2e26) |
| `http_receiver.auth_token.token` | [http_receiver.auth_token.token](data-sources--global_log_receiver--reference--group-002.md#canonical-212714edc1f0c237204594a28c58e49ab0f6e860365434f4e797b451dc54135d) |
| `http_receiver.auth_token.token.blindfold_secret_info` | [http_receiver.auth_token.token.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-461b1cc272eb18994019e3576da6fc74eadfbad679bda4e2635a78d0c1e51672) |
| `http_receiver.auth_token.token.blindfold_secret_info.decryption_provider` | [http_receiver.auth_token.token.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-7139a81111a08aa2ae64a8e6f57e411b98de84ce67ca03ae64db1dae64191976) |
| `http_receiver.auth_token.token.blindfold_secret_info.location` | [http_receiver.auth_token.token.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-002.md#canonical-2f6c7fb752f6ef91c882f02751482b8eab988d37381e19993f27291386843d3a) |
| `http_receiver.auth_token.token.blindfold_secret_info.store_provider` | [http_receiver.auth_token.token.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-880478675d62ab19f6102e7d5b200e787677ee63499eab852c414c472875e819) |
| `http_receiver.auth_token.token.clear_secret_info` | [http_receiver.auth_token.token.clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-51132397636dc8f3e8697e429737107d1b58bd09bdb750448d0b6316e347db1a) |
| `http_receiver.auth_token.token.clear_secret_info.provider_ref` | [http_receiver.auth_token.token.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-002.md#canonical-df0027f393ad7732156561351b77d9d36ba81c887178d1d96481b4d92fa89869) |
| `http_receiver.auth_token.token.clear_secret_info.url` | [http_receiver.auth_token.token.clear_secret_info.url](data-sources--global_log_receiver--reference--group-002.md#canonical-ff152eb7f27ee50bcdd347e166e3ed1c907751210713d49bae9bd59bef073f91) |
| `http_receiver.batch` | [http_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-cf0bc96796259fa40eafcbd02e580a711195bfa0959a1d485c68356561b21890) |
| `http_receiver.batch.max_bytes` | [http_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-002.md#canonical-d8c364a6a300974a5800494b4b190177aca6112710e1a57944f2564814832330) |
| `http_receiver.batch.max_bytes_disabled` | [http_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-0301b27b63919cd15c5bf72226d99de8090673c65d63b875b9885e637b6f1fd5) |
| `http_receiver.batch.max_events` | [http_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-002.md#canonical-a9470306f908b1367a9c4143e0d7043208823c5491d02b87c04b93535cb1da12) |
| `http_receiver.batch.max_events_disabled` | [http_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-477b5254a598775f0a0afbcfb3dc450666dcf59fc419258c1bd683e374b11b00) |
| `http_receiver.batch.timeout_seconds` | [http_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-002.md#canonical-120df0f073d7a4a7e6fae3730d1d41c2a7238895459c504c90857edd353bc6bf) |
| `http_receiver.batch.timeout_seconds_default` | [http_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-003.md#canonical-0e9d0c4960e1c8c66948d26deb3076da91b971d6ae0ebd37aa364c4ed017d487) |
| `http_receiver.compression` | [http_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-6c50d54b945021f0bdb741480c9cebdc1965f378f118522a51b2784abe7b97db) |
| `http_receiver.compression.compression_default` | [http_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-003.md#canonical-5065dd38ee2bbe8cacfe882c751b30859186541dd69f7bbcd04f19516accc6ba) |
| `http_receiver.compression.compression_gzip` | [http_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-003.md#canonical-b181489196ef22dd3bc1bddaac4aa20eeeed67a40af1b028fdf931ecd031476c) |
| `http_receiver.compression.compression_none` | [http_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-003.md#canonical-f92af2ac5d27141a3a0392c3bc013640a3bef8245ea45bae976d625ea43d1ee3) |
| `http_receiver.no_tls` | [http_receiver.no_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-7572b41242e440e6f6f4bb5cdf52118e1e78c2dc2a3148fe81fe679ee30794b8) |
| `http_receiver.uri` | [http_receiver.uri](data-sources--global_log_receiver--reference--group-002.md#canonical-a853fd4459b3bd75e7e8aabb18dbef36d8ca17e39d1e568b827290abda57a64c) |
| `http_receiver.use_tls` | [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-aa4e8b8d8d38465ca5a6fd8f3ec7dc36981977b167cf2705a2bd00c48acfcf98) |
| `http_receiver.use_tls.disable_verify_certificate` | [http_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-46d7ba665220528fbb56e00e55445f977dc42ee231f6810243f3a8cf4ace3be8) |
| `http_receiver.use_tls.disable_verify_hostname` | [http_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-b029dbb33f2ff2fcf14b77b8c3d35ef838dac463835453604660a5ab896a0610) |
| `http_receiver.use_tls.enable_verify_certificate` | [http_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-eec642e6269ba513a0e83aee691e912859e583b5be6bb3177d372318be53029d) |
| `http_receiver.use_tls.enable_verify_hostname` | [http_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-f8e3737b10451e6c376af765f799c5d8c194f6b690634014f3e027fc05b4e22c) |
| `http_receiver.use_tls.mtls_disabled` | [http_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-04ab152c359632c0c484f1cd3db30fa367439c070883e6e49726351ef2bf3a86) |
| `http_receiver.use_tls.mtls_enable` | [http_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-976329e75b15969ce6bfb460f4c83e19ebbbce612e89ff12091ba4f8b1d65ac8) |
| `http_receiver.use_tls.mtls_enable.certificate` | [http_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-c5bbd499bcf3a08fa1fdd2a6db2eac07dba6f106f18879eee2eacaffe70bf9aa) |
| `http_receiver.use_tls.mtls_enable.key_url` | [http_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-57caad8d342718d189f8cd6e6e71542a33060b56331b59f8e2986557e0593f12) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-777904f0a020cc856bb741ef138c820d839ab3a0f73b560739800e334e089251) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-003.md#canonical-792ab75b7de294b95ac4e22ac67af053de3768cb4bcfac36b0af34ac1ff745a1) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-003.md#canonical-c3c23170fc40dd16db9896bef93705e4cd57a33303c23bfc1f4cd16d9e843410) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-003.md#canonical-0926218d8c16f367888bed8f9b45d88cdc6efa9fade3431b5ed3144fba057a5f) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-e76a95671c4710ae4b95572185b18640078c4cf137759d7fdf5ac2c269f557e2) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-003.md#canonical-466dbd2ac17fde3041f6a47258db106fa616d63eedd9296acb0c135c2ae08ffd) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--reference--group-003.md#canonical-89656200f21bbb540241abaf7cebf4c2bbe9e962906f10d8c6037dcec302c0a6) |
| `http_receiver.use_tls.no_ca` | [http_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-003.md#canonical-5c2bf436a8a202cf82b7c978d7f29587d4cc2448a43ce9ecedc6c1217541c2e1) |
| `http_receiver.use_tls.trusted_ca_url` | [http_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--reference--group-003.md#canonical-32df6a960ccdf166c9a071f9db1ddbb3e5bcb94e8c9952961f58c8180340e6ce) |
| `id` | [id](data-sources--global_log_receiver--reference--group-001.md#canonical-a056c0c2a1b6d69f71077a4e06f64f2f7300bac72b7367ef24d8a1f158ecff39) |
| `kafka_receiver` | [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-a31e1cd1ccf0f2715372a58f5a0fc58c6a79d744538de04ea640d384c8395979) |
| `kafka_receiver.batch` | [kafka_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-53e9bf94b6c6aac7ccef2340cf936c3620f2979dc650a42e8905dbc602150071) |
| `kafka_receiver.batch.max_bytes` | [kafka_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-003.md#canonical-86f08a44657eb6e76050db3089764ed44e263dcb4d865ae555c3154ccb1d4dcc) |
| `kafka_receiver.batch.max_bytes_disabled` | [kafka_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-d6c9ec7b49f0b2a215bbff623073a3424c29c7f2cfadec61b35f7a0e2361da69) |
| `kafka_receiver.batch.max_events` | [kafka_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-003.md#canonical-f3bb912c262ac130008fbb167c6a32075ec2263ef669be881d7b26e4f448ae9d) |
| `kafka_receiver.batch.max_events_disabled` | [kafka_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-fc83403393425e1344f73dc2a2421b1ce58238060dc9cf3e7113e461e84e0308) |
| `kafka_receiver.batch.timeout_seconds` | [kafka_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-003.md#canonical-b5d48fff30e6b4320b6769c66324bcac0cb4611e29d4fa7fff3b587ecc02dfee) |
| `kafka_receiver.batch.timeout_seconds_default` | [kafka_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-003.md#canonical-de9d06e1d3dad50b6982f4df6af621115325fe21701277f10098f70a0e9325b3) |
| `kafka_receiver.bootstrap_servers` | [kafka_receiver.bootstrap_servers](data-sources--global_log_receiver--reference--group-003.md#canonical-7e2cd0638734932dbfa0adfbeddf7c4d5c1911d62517780cabad8dc84eb5aa7d) |
| `kafka_receiver.compression` | [kafka_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-dadf47410aa391ddf47afa8ea17ab1746f66e8666ba9358bf26aa3bb7df367e8) |
| `kafka_receiver.compression.compression_default` | [kafka_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-003.md#canonical-c8436c1f40229d3a3f12a43978b664f211275957632585b6ba8b39a227103813) |
| `kafka_receiver.compression.compression_gzip` | [kafka_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-003.md#canonical-855b25b4f3642ff8a85dd5eec79358bba4c81b03002a9642d891b501a703f605) |
| `kafka_receiver.compression.compression_none` | [kafka_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-003.md#canonical-e1eac692a08e188a817fb4ff376811de7723fa6d4e6c6d0d254dc311bed66203) |
| `kafka_receiver.kafka_topic` | [kafka_receiver.kafka_topic](data-sources--global_log_receiver--reference--group-003.md#canonical-86fd8dd45d23424f0efec599305967471a959ee70fee186b837ae6692ae91f1b) |
| `kafka_receiver.no_tls` | [kafka_receiver.no_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-bd04ced755f2bd79e01964c4ff110b46efc7bef0e198a24c4c13a851266cea4a) |
| `kafka_receiver.use_tls` | [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-fc7a0410084a7323afbf3d3432f45fa98857d4cd46fea4673d8314262ebe8d7d) |
| `kafka_receiver.use_tls.disable_verify_certificate` | [kafka_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-6b869ded0d6fd18c4ce1882aa8db864cbfce7c34e19782fcadb8f106a2c898b0) |
| `kafka_receiver.use_tls.disable_verify_hostname` | [kafka_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-4909e83b6a5c710921b55c398053237589737feea8831ac90ad0396db844a40f) |
| `kafka_receiver.use_tls.enable_verify_certificate` | [kafka_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-539556431e452db7e8d30dbbc66125b08ec40c9d2c2a7a67ff731dde9395cab2) |
| `kafka_receiver.use_tls.enable_verify_hostname` | [kafka_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-f21707b40187ac8e273a44b8e64aab24dabe9bb8cffabc08b3f7afbfa14e32a4) |
| `kafka_receiver.use_tls.mtls_disabled` | [kafka_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-86403a836f2b7c58623b124f9df0fc74b080094ce22d379f838616ee66c32717) |
| `kafka_receiver.use_tls.mtls_enable` | [kafka_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-4f82aee5b8986387dd0a106587743aae6329f8c08936de64966b0e84a2099b22) |
| `kafka_receiver.use_tls.mtls_enable.certificate` | [kafka_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-1a6b774764971831650f041835945bb5a33d8ee8238530189e5db1b53dbd0092) |
| `kafka_receiver.use_tls.mtls_enable.key_url` | [kafka_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-f56b685179c49e34e8efc65c61cf6adbf0dbb9d1133fdb8a6f079dfcad45856f) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-0516261b0949d016e4fa60764d64a1d836f82a4a0b729d3d02c23152ca1eb885) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-003.md#canonical-6a598e77b7343bdbb07574676628c5ab0ba54ddfd2c755a9695045f87ec9d465) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-003.md#canonical-a40f121f3ff7e7cccdf70ef25fb2428d593a2048486fa2117278d4ffb0e657dc) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-003.md#canonical-b3a4a19f341cdeeab62fb551e37776ed7ff80c94ae19c30f686dbeec2e64a344) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-a114592615df6133f0aaf34c9cf40ef0cf88280f732b7b3f58f032af81d53bf0) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-003.md#canonical-99584070aa26d4031e82ef1f5565a405f0cf4fa6a0f990c676608d2cbe19f359) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--reference--group-003.md#canonical-8c76f0a6989b2634fb5c8a0d92bf72f158d6f45dd36f112d3be7b48adb977d63) |
| `kafka_receiver.use_tls.no_ca` | [kafka_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-003.md#canonical-a3c91c6c718fd42b4e06946653535eafcef04e43cc5cccaaa372a0bfd915e652) |
| `kafka_receiver.use_tls.trusted_ca_url` | [kafka_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--reference--group-003.md#canonical-2d8f708a6168f399bb8e044c07091bc1aff3895282c30c4a2cc6d7a473c21274) |
| `labels` | [labels](data-sources--global_log_receiver--reference--group-001.md#canonical-94995592684160b7dab0af8037ad0a3d507fae58b962efb8ad5de382db15f243) |
| `name` | [name](data-sources--global_log_receiver--reference--group-001.md#canonical-007760a76f5a884982d2dbbcac8bc948eaa6b9f7d166fade141e16d3e1012f0b) |
| `namespace` | [namespace](data-sources--global_log_receiver--reference--group-001.md#canonical-4584b30ea2792d16597d6c62636f22527cfd57472f11f1e00887d0dfe0b066d9) |
| `new_relic_receiver` | [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-de4a336b302a3b488fb87be5ddc2114a24a5cc1ba754c0ed424c2c8c71365141) |
| `new_relic_receiver.api_key` | [new_relic_receiver.api_key](data-sources--global_log_receiver--reference--group-003.md#canonical-70959860813e842268b5c73806c213433e44710f27ee47f4b5ae4c8cad1dc771) |
| `new_relic_receiver.api_key.blindfold_secret_info` | [new_relic_receiver.api_key.blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-6d4f84f71264dacc88d3fc0abbffe91102fb1d2c1b984707656c3d035ed2f383) |
| `new_relic_receiver.api_key.blindfold_secret_info.decryption_provider` | [new_relic_receiver.api_key.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-003.md#canonical-de82c3b32fc5fa2bce3ab23d8b5c5cbb06e83d47aab0a6fa519d8b8768e4434a) |
| `new_relic_receiver.api_key.blindfold_secret_info.location` | [new_relic_receiver.api_key.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-003.md#canonical-0b233eac3d86a24b5c76c0cfc3c38e3e721bd97e5104769a73672286a4fb60c9) |
| `new_relic_receiver.api_key.blindfold_secret_info.store_provider` | [new_relic_receiver.api_key.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-003.md#canonical-b3fa2c78844c49dd7524cd3a32f324c532ddb711162c76f4d469bf09e59074a4) |
| `new_relic_receiver.api_key.clear_secret_info` | [new_relic_receiver.api_key.clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-bbef8af3f744997fb19a474aa17488cfd6d35fdee771b8b394c45f6b76326141) |
| `new_relic_receiver.api_key.clear_secret_info.provider_ref` | [new_relic_receiver.api_key.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-003.md#canonical-0993bb241a1066539e70aea7adc033ef2150b5f5db3c07fb8b1daf3dae4e0a09) |
| `new_relic_receiver.api_key.clear_secret_info.url` | [new_relic_receiver.api_key.clear_secret_info.url](data-sources--global_log_receiver--reference--group-003.md#canonical-656a13b8277ed158fe98c12acc1d34f32d3e349aed2df172f5b62e7e3fb3b9a6) |
| `new_relic_receiver.eu` | [new_relic_receiver.eu](data-sources--global_log_receiver--reference--group-003.md#canonical-a97238c0fe6143038aef5dcb1f5cb04b653ef091e11b6bfc311a9bd564666e2a) |
| `new_relic_receiver.us` | [new_relic_receiver.us](data-sources--global_log_receiver--reference--group-003.md#canonical-f267c9272a349908094306debf89123593f1f52c89f843ed9ef97ba7a0209904) |
| `ns_all` | [ns_all](data-sources--global_log_receiver--reference--group-003.md#canonical-4c61613dac1b7423de7e703f6ebfc191efeb081005fa60f1948e50325101b0c7) |
| `ns_current` | [ns_current](data-sources--global_log_receiver--reference--group-003.md#canonical-a0e14c04d927e045973f735ce6e43e5cab4df82b3725fd35a09f0773ca47fc02) |
| `ns_list` | [ns_list](data-sources--global_log_receiver--reference--group-003.md#canonical-c5020c3ff01ccbcb7f51e4dd0018c982755920a0d4dfa0b182cbc771682ad584) |
| `ns_list.namespaces` | [ns_list.namespaces](data-sources--global_log_receiver--reference--group-003.md#canonical-b7bc4b2fd247f874259644c28d852e75bb063e2e306a68f8aaa03ebbdc7b2546) |
| `qradar_receiver` | [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-63e1be9603a2f1112924bcec75ff8b9f9b9b2825fccf9f55c0d1c4a43b8caf74) |
| `qradar_receiver.batch` | [qradar_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-18d61cb5d591896c9ae80eb6d9e52e9c210d02fb4dd296d4570162f356e9f8c1) |
| `qradar_receiver.batch.max_bytes` | [qradar_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-003.md#canonical-f8156477011b20209d669f357f9e64dc712c4a11019b962ac887a02c270d7b03) |
| `qradar_receiver.batch.max_bytes_disabled` | [qradar_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-e2eb3fc8f3b0ff4b61accfd0c69d8638a00216964eadfbfd1bf7753604e34179) |
| `qradar_receiver.batch.max_events` | [qradar_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-003.md#canonical-17a67128efe8c65264683d106db6986d0392430cb7d51e85e996c6225720b1da) |
| `qradar_receiver.batch.max_events_disabled` | [qradar_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-d10152deb0ec74139d2b01578eec86f350f8e2f953cf9850d680551f30cb79ed) |
| `qradar_receiver.batch.timeout_seconds` | [qradar_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-003.md#canonical-5106ec2de2aeec40ba6c0f213fcb939f44e9364a89361faeba90177890e78558) |
| `qradar_receiver.batch.timeout_seconds_default` | [qradar_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-003.md#canonical-87570d7c5eadacfce5890c91f27d7211a89e6588ab86f012f823c0a8114df7b7) |
| `qradar_receiver.compression` | [qradar_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-333b8d781a27bb4f2ca2436d507ddb1fbe6473e28b5c4c300067c5710e2d5f82) |
| `qradar_receiver.compression.compression_default` | [qradar_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-003.md#canonical-4f5d489c596a20bac41f9f817b75f3adf68d092297cfe5b552e96a6b019d9a83) |
| `qradar_receiver.compression.compression_gzip` | [qradar_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-003.md#canonical-efe622676ccb3473753dae8b15cabb2acf751fd60b422d47cdda9f1a03790ca4) |
| `qradar_receiver.compression.compression_none` | [qradar_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-003.md#canonical-4c8c45721e36189f9ccb2d754b1e28078dd899fecbd3e64a612910d1d6fabccb) |
| `qradar_receiver.no_tls` | [qradar_receiver.no_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-75a44f1d1031a8471baa0fd8a92ac2fe1f6c495e30fbdd73b3127432703625aa) |
| `qradar_receiver.uri` | [qradar_receiver.uri](data-sources--global_log_receiver--reference--group-003.md#canonical-7647ab2b5fccea37a28b2aee902ae6b9a441378cd8f69e288091130b0b71a504) |
| `qradar_receiver.use_tls` | [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-aca335a47ecfe2fb33c34eab49101bc1a1029749d745bb8e64a4cbf13b287a72) |
| `qradar_receiver.use_tls.disable_verify_certificate` | [qradar_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-44514df5ca5709d389e614cd04a33f01d8f264e9f5abef71b706b91fd25a5ed8) |
| `qradar_receiver.use_tls.disable_verify_hostname` | [qradar_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-0f39c62fab243475d1d23aa99084782cb61357585962b55264df06512a271c61) |
| `qradar_receiver.use_tls.enable_verify_certificate` | [qradar_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-d3db25716d78ef3fc901f11e7ebbedfb653e6ec6244877b8e1240ea8e6f5f4c3) |
| `qradar_receiver.use_tls.enable_verify_hostname` | [qradar_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-be0cf896738c5e38bd1612313a4c1a0f3212e9f26b3926ff03dd0073b30597c3) |
| `qradar_receiver.use_tls.mtls_disabled` | [qradar_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-4d09553a085ec5fe3b9732a823327c4e0339d5c73050e91cb8d1a7d96b9f178c) |
| `qradar_receiver.use_tls.mtls_enable` | [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-5203b1edc5f8da69ebea5e9cdff7bad9eecd3bd33da6a7566b4729d5e9c67068) |
| `qradar_receiver.use_tls.mtls_enable.certificate` | [qradar_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-4bb979d800e90d819b4f0ef572613be12191da15ae8865b39a8ce33443d1fb7a) |
| `qradar_receiver.use_tls.mtls_enable.key_url` | [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-c7aa5b3930e519b01d85c4e114d15eb07a231e8102c1c51e465aaa05427942c7) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-0fd96d72a5ac05a5c8841ea837aba13bb49e3b9431f8f34725a204b3347f2064) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-3af39a7d4779b04e9193ed4c07e0dfc08edec8760cd2dbf97d1ccb13f2be653f) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-004.md#canonical-e9b8ca6ec72623cbba7b9ddba0081b9d314d0f89caf15294107782a102b5a58a) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-626be832aa74e93743b31b7c1bcde5fa9a5fd61ab51880d8a905c7abce039846) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-a1936f621178a74aaa10b9928b1c78d2f62fa15acf4857d48d3a76f3de174e32) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-004.md#canonical-1fb2d311a83b4845f1c0c7bed9a9a547828c166b1029d3e45600de696c8bcff5) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--reference--group-004.md#canonical-df9593fc2f1f916508281b6afc2f3bdc3a33105cabcffd8c543e8725800af392) |
| `qradar_receiver.use_tls.no_ca` | [qradar_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-004.md#canonical-8adb64c76c96940ebe498a228b0ab2ace2556cc658f955660a2edd4501ffae1b) |
| `qradar_receiver.use_tls.trusted_ca_url` | [qradar_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--reference--group-003.md#canonical-fb06c0d68e5a89ccab5a28718a3e2307d38b62212265f8244b2e8346d0a21732) |
| `request_logs` | [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-4e2818b42395658eeb7c84785af4bf6b05a6a0bb11bef566421f96787bc2b52d) |
| `request_logs.sampled` | [request_logs.sampled](data-sources--global_log_receiver--reference--group-004.md#canonical-2700c0cb72e82e3c51ba6bf17f14cdaa2efb7da8c2bd97f35dee1cef21c11658) |
| `request_logs.unsampled` | [request_logs.unsampled](data-sources--global_log_receiver--reference--group-004.md#canonical-b2799c59167c5d2a69da737ee2143a2369004e5fc66edd7b671afc294694dfbc) |
| `s3_receiver` | [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-e6569dee683a809947ec9064923b603e50489bad872e1b88bf8fec828a194c03) |
| `s3_receiver.aws_cred` | [s3_receiver.aws_cred](data-sources--global_log_receiver--reference--group-004.md#canonical-e218bb346ee7f6d594c482696a312c164acbf6a5022d8cdaf4ad10f26ed6b855) |
| `s3_receiver.aws_cred.name` | [s3_receiver.aws_cred.name](data-sources--global_log_receiver--reference--group-004.md#canonical-3a8ab356ce3acc9cbceaa924a6f01b6227bbe3611c0f725459a95bd32608d2ea) |
| `s3_receiver.aws_cred.namespace` | [s3_receiver.aws_cred.namespace](data-sources--global_log_receiver--reference--group-004.md#canonical-705c04f98bbe168cd7e12a6297276290851bc9ffde7109e684299bd5b97e988a) |
| `s3_receiver.aws_cred.tenant` | [s3_receiver.aws_cred.tenant](data-sources--global_log_receiver--reference--group-004.md#canonical-77a207df33189e84bafe6781179130c012031593edebccffad7c671c13b65c88) |
| `s3_receiver.aws_region` | [s3_receiver.aws_region](data-sources--global_log_receiver--reference--group-004.md#canonical-dc58aa1b3a12744c8c068e45ea5de745c9274d0996fa328687dad41a0c31f924) |
| `s3_receiver.batch` | [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-c8b264cf9f82e39dcc4ada282650ae84579eb4df611cf8ddfb539c8d8b353f9d) |
| `s3_receiver.batch.max_bytes` | [s3_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-004.md#canonical-6e0f86194c793191f0d93310c6fe5075487c4fc28cac7f00f2c326f06da8f329) |
| `s3_receiver.batch.max_bytes_disabled` | [s3_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-4ddbdb5f8c994a75895518eb944f7aa20735e06945deca9c6b342b1cb6d4d488) |
| `s3_receiver.batch.max_events` | [s3_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-004.md#canonical-6903799bd4d481785dfa88f505f7f3c15257bb5916c190b7598a4ec06a393a43) |
| `s3_receiver.batch.max_events_disabled` | [s3_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-dbf83ad9a7d045024b2d2950ae6fa9169f5cbd418924e4730172de376726d0b8) |
| `s3_receiver.batch.timeout_seconds` | [s3_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-004.md#canonical-60675b24dae7f21318aa3102eed3991118bae1fd59bdb28d47caeeefed8c4147) |
| `s3_receiver.batch.timeout_seconds_default` | [s3_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-6bd760ceb829fb4969103ab06bc019e3f5b499462ee31956344b3484e831df37) |
| `s3_receiver.bucket` | [s3_receiver.bucket](data-sources--global_log_receiver--reference--group-004.md#canonical-049ee00c58c3abf4894c8cf5e8a7def624b9d156daed9d444ec13a868e1210f1) |
| `s3_receiver.compression` | [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-c1fe7d2192cafb05ceb0b440ee5022b15206de9f8c4fee42c2e2a1e1361d9c27) |
| `s3_receiver.compression.compression_default` | [s3_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-9115ec844c3bcafebd60303c09e849af9384d8584cf769da79a90437a21ea604) |
| `s3_receiver.compression.compression_gzip` | [s3_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-680365dfa46a41d9c2dbb40dbe2bd251486924c93ab5ccc586ede4530852766d) |
| `s3_receiver.compression.compression_none` | [s3_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-2a27204842edaaec9929d7b4802e982575f1613ac5006309c58905c720a9c536) |
| `s3_receiver.filename_options` | [s3_receiver.filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-4eb0d9f5b081ae8232b16d159840e808b152790389fb38557affa33ad738d2c6) |
| `s3_receiver.filename_options.custom_folder` | [s3_receiver.filename_options.custom_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-ece9e20f1a3047ae68813023df4cbd133f3f9ab1e617b07aec01226213a1aaba) |
| `s3_receiver.filename_options.log_type_folder` | [s3_receiver.filename_options.log_type_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-babf622d80618bffd70b1e54007b2aeede94fd3aee7bcfb9a2dad24afa5819e9) |
| `s3_receiver.filename_options.no_folder` | [s3_receiver.filename_options.no_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-e5698fb7d9e33b0c2a77a9e282e83891ea358ed14c54cd8094393ab25ea3f344) |
| `security_events` | [security_events](data-sources--global_log_receiver--reference--group-004.md#canonical-998417cd4d62d05d5ed56692f9c2eee42b1e26b6d72d2278edfff8d76688f244) |
| `splunk_receiver` | [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-9cea3aa814ee71326f9bd3e207fbff67b5828a4c37f9c9f8eefe23a77dc8bd31) |
| `splunk_receiver.batch` | [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-34a2df733d0986cd3156d9a64bdd36da00f1e324c8764d9b601ce6d101e5b6f0) |
| `splunk_receiver.batch.max_bytes` | [splunk_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-004.md#canonical-776fc6e70314e46fd02d554c5e7138cb0bd5978b253b4387abaa5aab132b183e) |
| `splunk_receiver.batch.max_bytes_disabled` | [splunk_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-6e2f37d1c855e45b12debbf7b019c6f49a4b2d9eb651431c0b680d7a93bb7625) |
| `splunk_receiver.batch.max_events` | [splunk_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-004.md#canonical-3e068b35448217058c7b3208e5e168674babc83794d9013fbae8bd3b9ec61d65) |
| `splunk_receiver.batch.max_events_disabled` | [splunk_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-758275a0ed18e57c894658ed4fff05162cdbb38a6b7a7a21df9a325320f3767b) |
| `splunk_receiver.batch.timeout_seconds` | [splunk_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-004.md#canonical-24c9ff7b82d3935fb4c5ebd7ab232c12809da4f450aa0f11e4c1da50db2488a8) |
| `splunk_receiver.batch.timeout_seconds_default` | [splunk_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-12d29d50f9c214f458879e12a035b24bb7d7dd9c8781826f08d19151e7eb7cd3) |
| `splunk_receiver.compression` | [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-5eee87c1216fe43a4fd3c796e79c7e9a30199b5bc8e47e6efb9ff8f24c3a97ce) |
| `splunk_receiver.compression.compression_default` | [splunk_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-f5f4d6b226d8a2af36388ba49e8f494d13d4127a42ab800be54f66e729d83ef2) |
| `splunk_receiver.compression.compression_gzip` | [splunk_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-9e73327291839cdc29e5594aaf0c836bfbd7853549a14a067338d6a3e54296d0) |
| `splunk_receiver.compression.compression_none` | [splunk_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-a94013c4d2eb50e7948ab142d6e91f813185cdd57ac347829813d58086cde01e) |
| `splunk_receiver.endpoint` | [splunk_receiver.endpoint](data-sources--global_log_receiver--reference--group-004.md#canonical-1ad26994d6accd5bae3acbf0cc389f900e16db179300327e110ff93607d7cb29) |
| `splunk_receiver.no_tls` | [splunk_receiver.no_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-7be2ff8bbf7dd1e630b7707ec28bc7b02d81d3aaebea54a485034a3a7602ea32) |
| `splunk_receiver.splunk_hec_token` | [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-044e2dcbc57b8f62129b6426c180c925b5e20308fa3ec8dfcb4e062683f44e49) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info` | [splunk_receiver.splunk_hec_token.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-de9221762147d6176ab8c8246c67fc3b6c452b8f08cb85a7b303f2d2cece57d8) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.decryption_provider` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-b345e3a32559b8bbe550c8231dae519c54fa46c5bed31f9324dc6e275578eed3) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.location` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-004.md#canonical-13359d5dfbdca6ebe3555c27b0b72c222e633f8e513e3b71538ae494c72005e3) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.store_provider` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-75748dce5fdc52d1cb36abc510c089d1757ae5f3994269c65a41188fdf6e81fc) |
| `splunk_receiver.splunk_hec_token.clear_secret_info` | [splunk_receiver.splunk_hec_token.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-202ef4d0c52ec01c10c6adfdd944ec782dc82c7ad10a9b144e1fad139ae2e257) |
| `splunk_receiver.splunk_hec_token.clear_secret_info.provider_ref` | [splunk_receiver.splunk_hec_token.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-004.md#canonical-7fddef634373fa41eabcacc6eb6cde15b940369b5acd989f0fba2763916799e3) |
| `splunk_receiver.splunk_hec_token.clear_secret_info.url` | [splunk_receiver.splunk_hec_token.clear_secret_info.url](data-sources--global_log_receiver--reference--group-004.md#canonical-2d736ecc5cdac838b855c2629d6e4c5378568366215045719a8c6ffff9db327f) |
| `splunk_receiver.use_tls` | [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-4d99715e85fec97e9bd6aff8fd89b9c95792a8daed8e69dc71e3587b77dbe4d2) |
| `splunk_receiver.use_tls.disable_verify_certificate` | [splunk_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-abb7e0b70e2abd366dadfb109b660e27cfbaa253ba3b8ead5b41f74dc3e4a54d) |
| `splunk_receiver.use_tls.disable_verify_hostname` | [splunk_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-a82b9462cb3d11d7d7fa805695d852dab1f201c3d7bb5678dc7b82ff82438ece) |
| `splunk_receiver.use_tls.enable_verify_certificate` | [splunk_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-2ceaca628858702cdb41a68d5f289dda8d63dde178acbc7f490eac8486de74f9) |
| `splunk_receiver.use_tls.enable_verify_hostname` | [splunk_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-b46c1dac09c40b284fb42e09aab194f30f2b606e9df403c8f09293dabb01728a) |
| `splunk_receiver.use_tls.mtls_disabled` | [splunk_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-c60ee2a1d21d974a2e545ec9809e6618df78ffff3b433f5754e053fc53038507) |
| `splunk_receiver.use_tls.mtls_enable` | [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-84c913066466246d55d5f8a88979343ec606a4ea4bcef1c99154ebf41170aaef) |
| `splunk_receiver.use_tls.mtls_enable.certificate` | [splunk_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-1811a5840278a7188ec0d5f1d0a5c9e3989f94e63bcafd57fc2ab3d38ebeeba8) |
| `splunk_receiver.use_tls.mtls_enable.key_url` | [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-d786d848070a35d93f6e2b2dd75b9b8e5e5bb54365896daed911e13fc64fbc5d) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-98e5aefc296f96fa068108af27931ca91875859c1f87ec1583e5f1bbef71cc1a) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-d2e97925f3c484802181d46258617a374eac6db0dc3fb34eb9a21ba6148353b1) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-004.md#canonical-3a51c28d4c72ce2e91c0c01f7bd39083af12defcc3e4cada6ccd615e0333c71c) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-7975a309044e820703085f20b8312c37c2e6ab157f071b7b7b722af305ee048a) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-ee4f95880b5768442a7371a5a7c37686103dbe672fa10742f996359703e2eaaf) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-004.md#canonical-317d1fb57c20c83cb0f520804963901841283c437cc3b47299df8539d9a1d9ea) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--reference--group-004.md#canonical-c62ef53090ef57d9b81fae6539bf2c0bfb8eeced53a4120ba067e24aaff7f93f) |
| `splunk_receiver.use_tls.no_ca` | [splunk_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-004.md#canonical-1960c66bf5a637fa3d0703679b811caebce661c566aab9c43f75747cb56a97db) |
| `splunk_receiver.use_tls.trusted_ca_url` | [splunk_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--reference--group-004.md#canonical-4f815177ec88baea9133b6f243ef318d9b42ac2b1543fbc2add7e0b428b585a2) |
| `sumo_logic_receiver` | [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-95863243449b46ca8ba7bf4ea6f398b35ba9875c97fd3d58d4c6811920c9a0a1) |
| `sumo_logic_receiver.url` | [sumo_logic_receiver.url](data-sources--global_log_receiver--reference--group-004.md#canonical-cd5ed7f98721b5cfe5d7a06d3b8a10232ccec0455d9e782020bc57a294cfb487) |
| `sumo_logic_receiver.url.blindfold_secret_info` | [sumo_logic_receiver.url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-bb24244f5fd3a931d8c6bbfde05acc9f8d03181b32c217d765bf00f2b6c10dc1) |
| `sumo_logic_receiver.url.blindfold_secret_info.decryption_provider` | [sumo_logic_receiver.url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-59bbe6f718ab3cd7b5aaacbc9a70813f725372cc9a251ea7d311a49714911d0e) |
| `sumo_logic_receiver.url.blindfold_secret_info.location` | [sumo_logic_receiver.url.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-004.md#canonical-b0f2dd09721aaf1b7475de5e952526890831b1c35f39b61012c7d44b5666b085) |
| `sumo_logic_receiver.url.blindfold_secret_info.store_provider` | [sumo_logic_receiver.url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-31af75e245dacacc8a8fe8344b49fdd6589e9acfa2f551eefaa4971ec7eccb47) |
| `sumo_logic_receiver.url.clear_secret_info` | [sumo_logic_receiver.url.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-098ebdf4a57abe43e3c77acf8d241f3b1b88b22709eca7dc4b5c905b776489c2) |
| `sumo_logic_receiver.url.clear_secret_info.provider_ref` | [sumo_logic_receiver.url.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-004.md#canonical-73c14d2eb6766b1fe5f820db4654175ce1ca0e1aed5cfe3238fe488df55aaa86) |
| `sumo_logic_receiver.url.clear_secret_info.url` | [sumo_logic_receiver.url.clear_secret_info.url](data-sources--global_log_receiver--reference--group-004.md#canonical-6775b126b13e527e2ae655d822f2232d38804c4e5265f89e94b4e68f789d0a23) |

<a id="canonical-926ad7e40a48f27ac3967b91dd34d2894cc075c6f54e115ec783626f6625ed52"></a>

## Next pages — Property reference / ee2aefa3af0a / 11

- [audit_logs](data-sources--global_log_receiver--reference--group-001.md#canonical-df9883b23b9a3363e082d3c4b3d4c306c14b6778efb738cf51239d9fa1e7fbfb)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1e93724fa97f305e8285ab20a3de95bcda72cb3dd4afba2b5d1f659e5b82aa39)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [dns_logs](data-sources--global_log_receiver--reference--group-002.md#canonical-b7191829d34da96f0435f5aec41d4e561e626957b75cdf343f72630ff0e93240)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-6a8ca15a6755e8c761d98fa0efc7d887c67f7c2cf83ede5ae159146cee36621c)
- [ns_all](data-sources--global_log_receiver--reference--group-003.md#canonical-eac6c2cc015966b8d0a1b560e162a5757c3992ef61ab6e3555f7a8641fdb1d88)
- [ns_current](data-sources--global_log_receiver--reference--group-003.md#canonical-e7d88a8682a21bc20dc8d78ab563c10f0d3a2ee20a4b1399394c1bda0161b874)
- [ns_list](data-sources--global_log_receiver--reference--group-003.md#canonical-f360f7b122a79b7d1c6f6f9d22020aaec9eb1157401d37099f911acc32f3529a)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-78587f316fc037494d3bdda63e54544614ce095432c2e1335fabfd6ee00f891b)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- [security_events](data-sources--global_log_receiver--reference--group-004.md#canonical-4234f3746f0500e78349b4be62bcfa5b49542d6219128d146579072b81afbd0b)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-60223cb26da3c4491af2cd9271ba7071122485608016a658e918ce37e12dc0c1)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-df9883b23b9a3363e082d3c4b3d4c306c14b6778efb738cf51239d9fa1e7fbfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b91213571fe0930011290404ced9481f29605df28e796b511571773756243e30"></a>

## audit_logs — audit_logs / 1ff9661b7242 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- audit_logs

<a id="canonical-9fa6306bdbb6069dd4b9aa668f4e2c661e6f3e04d5ff23da549dd8b195a9df42"></a>

Type: `["object", {}]`. Computed.

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

- [audit_logs](data-sources--global_log_receiver--reference--group-001.md#canonical-9fa6306bdbb6069dd4b9aa668f4e2c661e6f3e04d5ff23da549dd8b195a9df42)
- [dns_logs](data-sources--global_log_receiver--reference--group-002.md#canonical-3c0cf4cfae7a37e36a098964bdb7ac68726ba850bef0307a663b00b39428e737)
- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-4e2818b42395658eeb7c84785af4bf6b05a6a0bb11bef566421f96787bc2b52d)
- [security_events](data-sources--global_log_receiver--reference--group-004.md#canonical-998417cd4d62d05d5ed56692f9c2eee42b1e26b6d72d2278edfff8d76688f244)

Select alternatives according to the provider validators above.

<a id="canonical-e59955a4e8d76e35f6f3b11e3d536dafb7532fa515f777f17d0fa23fda7e3f19"></a>

## Direct properties — audit_logs / 1ff9661b7242 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7fa7d0377e210144d07caa29ab8d8e002f94a7853521b8bef986e7196da2dcdf"></a>

## Next pages — audit_logs / 1ff9661b7242 / 4

- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42bc14b08338c78ae291b1d2951c260f19ae9167680849c4f8eccbcafaf053e9"></a>

## aws_cloud_watch_receiver — aws_cloud_watch_receiver / bc4d2b2f0c3c / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- aws_cloud_watch_receiver

<a id="canonical-fd5167464041e5466356ca926c5af5970183f82d36c7bb3345d6187ab1da0f97"></a>

Type: `"single"`. Computed.

\[OneOf: aws\_cloud\_watch\_receiver, azure\_event\_hubs\_receiver, azure\_receiver,
datadog\_receiver, gcp\_bucket\_receiver, http\_receiver, kafka\_receiver, new\_relic\_receiver,
qradar\_receiver, s3\_receiver, splunk\_receiver, sumo\_logic\_receiver\] AWS Cloudwatch Logs
Configuration for Global Log Receiver.

Upstream description:

AWS Cloudwatch Logs Configuration for Global Log Receiver.

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

- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-fd5167464041e5466356ca926c5af5970183f82d36c7bb3345d6187ab1da0f97)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-5f8d310c2746fa4c17c4a70016a2a5fe97191aed64d1cd597a98b1a3d89f2fee)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-33d0516820774d65b0cb95a87591491c90c8a78ee4996ddde7ec41aba356117d)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-4c6892ae68db2e1b186f351c385d1019fefce792671cb3df607e0554ec2ddc23)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1128734e088b037c58c9603f737bff23c86df51c827cfc4f1932c9f3cf5ed364)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1128fb5d3d174ae457010bed22d2085ce139b22d5a8cc9a82b46392b50accfab)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-a31e1cd1ccf0f2715372a58f5a0fc58c6a79d744538de04ea640d384c8395979)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-de4a336b302a3b488fb87be5ddc2114a24a5cc1ba754c0ed424c2c8c71365141)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-63e1be9603a2f1112924bcec75ff8b9f9b9b2825fccf9f55c0d1c4a43b8caf74)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-e6569dee683a809947ec9064923b603e50489bad872e1b88bf8fec828a194c03)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-9cea3aa814ee71326f9bd3e207fbff67b5828a4c37f9c9f8eefe23a77dc8bd31)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-95863243449b46ca8ba7bf4ea6f398b35ba9875c97fd3d58d4c6811920c9a0a1)

Select alternatives according to the provider validators above.

<a id="canonical-69a94f1155200e6214c1eabbb0ad82ba2d2cfa8868038b7835847696d5965d0e"></a>

## Direct properties — aws_cloud_watch_receiver / bc4d2b2f0c3c / 3

- [aws_cred](data-sources--global_log_receiver--reference--group-001.md#canonical-3d541a8c2991a2a69fdd3dc41ba6747736cb653461fd620126c416cdc2dc55fc): complete subsection reference.

<a id="canonical-f54a0193cece959c174850819f25161e5410ea3af613e40f78614f5d6f9b2659"></a>

<a id="canonical-a7d1164c09f72f9e53e06adec88415aa908400a46a0a4918dfe505ebd8e871cc"></a>

## aws_region property — aws_cloud_watch_receiver / bc4d2b2f0c3c / 4

Type: `"string"`. Computed.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
AWS Region. AWS Region Name. Possible values are \`ap-northeast-1\`, \`ap-southeast-1\`,
\`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`, \`us-east-2\`,
\`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`, \`ap-northeast-2\`,
\`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`, \`me-south-1\`, \`us-west-1\`,
\`ap-southeast-3\`.

Upstream description:

AWS Region Name.

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

- [batch](data-sources--global_log_receiver--reference--group-001.md#canonical-e0af9183ee489c2211fa5ade2ad069999fc1819b782cbe767f43679fd3e67478): complete subsection reference.

- [compression](data-sources--global_log_receiver--reference--group-001.md#canonical-8e8aedf1590d974b9d62aae8188b22fed053703badf6e150a9c07fcaf6849f57): complete subsection reference.

<a id="canonical-07faf075fb4c52ffb782b86bc7f8dbaafcfe53fd28604ab373a8b19310a58272"></a>

<a id="canonical-3ab8720477cde01d44ce5b6cbb42c9a99d1db949cf59187fae9ab02933190ba0"></a>

## group_name property — aws_cloud_watch_receiver / bc4d2b2f0c3c / 5

Type: `"string"`. Computed.

The group name of the target Cloudwatch Logs stream.

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

<a id="canonical-4afd4f5b4d1013473e50d6aa425922047cd4286571d67a984e1cd2e833f8c273"></a>

<a id="canonical-df285d68510598b95cace9ec3db5304ed6a8c6bddff41fdb97ee3351fb435f14"></a>

## stream_name property — aws_cloud_watch_receiver / bc4d2b2f0c3c / 6

Type: `"string"`. Computed.

The stream name of the target Cloudwatch Logs stream. Note that there can only be one writer to a
log stream at a time.

Upstream description:

The stream name of the target Cloudwatch Logs stream. Note that there can only be one writer to a
log stream at a time.

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

<a id="canonical-3e5d4062a7cf7c938f62b2ea352920065c322e723e5bc9a3b56a494a679c0e69"></a>

## Next pages — aws_cloud_watch_receiver / bc4d2b2f0c3c / 7

- [aws_cloud_watch_receiver.aws_cred](data-sources--global_log_receiver--reference--group-001.md#canonical-3d541a8c2991a2a69fdd3dc41ba6747736cb653461fd620126c416cdc2dc55fc)
- [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-e0af9183ee489c2211fa5ade2ad069999fc1819b782cbe767f43679fd3e67478)
- [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-8e8aedf1590d974b9d62aae8188b22fed053703badf6e150a9c07fcaf6849f57)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-3d541a8c2991a2a69fdd3dc41ba6747736cb653461fd620126c416cdc2dc55fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b5d2846ce0fdcd4bc1957ecf9c8fba2a5c79cfb9558e9858928b0f508e1eb43"></a>

## aws_cloud_watch_receiver.aws_cred — aws_cloud_watch_receiver.aws_cred / 9b0a545162fa / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e)
- aws_cloud_watch_receiver.aws_cred

<a id="canonical-d94456f96c37418dc58a53822fcb1d626e6b35b5ae90f4453180333e702e552e"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-d017141c74a69a215aa515e7f91347f3f6dcf6af29ca51ae254a7dc8e1662b6a"></a>

## Direct properties — aws_cloud_watch_receiver.aws_cred / 9b0a545162fa / 3

<a id="canonical-22729212426288499175d384f069c8e3bf45966b7eed9e2008006fecef1e2e29"></a>

<a id="canonical-45d8edb3902780c6942e0df80b72ee3acb0394384936c3aaf60b662abcbaf2a1"></a>

## name property — aws_cloud_watch_receiver.aws_cred / 9b0a545162fa / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-539553a479f51cf160ef6106331fc18c42a8ee71510475bb1eb9ee59d48e8c0f"></a>

<a id="canonical-ba977c631c350d7b3df94e851f4f997b20fa6fe3309f3999a0a441bcf6d77fef"></a>

## namespace property — aws_cloud_watch_receiver.aws_cred / 9b0a545162fa / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-c30db574f1137de546b01b378d382eb0de325dc899055896defd79173af7ce21"></a>

<a id="canonical-4779faeb27c2aaa57348a351001860bdcb90cf4f0734ad4e6c071563db9b2ee3"></a>

## tenant property — aws_cloud_watch_receiver.aws_cred / 9b0a545162fa / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-a1fa7186ffa0ac23a9f6c6ca53d222785f72b3903f68d0bc86a3db412dd0d692"></a>

## Next pages — aws_cloud_watch_receiver.aws_cred / 9b0a545162fa / 7

- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-e0af9183ee489c2211fa5ade2ad069999fc1819b782cbe767f43679fd3e67478"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-864086cb12a5f6579207fa1b798a85fd7531495fcc221c85a36d21215ba70e5f"></a>

## aws_cloud_watch_receiver.batch — aws_cloud_watch_receiver.batch / fea8c8d58d52 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e)
- aws_cloud_watch_receiver.batch

<a id="canonical-6ff77b87244ade51af7ac58c52492e902432ddba9d0927829908e60d171a1d9c"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

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

<a id="canonical-fb864a4078c1b8806b02f24fd2b04c9cf234cb24f0bf081b648ae05d3192dce5"></a>

## Direct properties — aws_cloud_watch_receiver.batch / fea8c8d58d52 / 3

<a id="canonical-d5ab853c25bd71f11cfbc5ff3e4d2c69d73b77ab03ae7e9dff37063d1fdccd1f"></a>

<a id="canonical-ec6afe2df6fb8fc17253395b633c8887a9ee80b305dedf33f49dc16be65ad989"></a>

## max_bytes property — aws_cloud_watch_receiver.batch / fea8c8d58d52 / 4

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

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

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-2f7de895b09020ece48800f5aa5b03ca9ac203c2ec0c524af28dbf6fa0bed084): complete subsection reference.

<a id="canonical-9c08e76aa58273d5ae1342787410951c6ed55b84b57629cbc279153dc096a8e8"></a>

<a id="canonical-a0085ea26eddcbce1340878859e9abf5a05dd74648a31f8346bc856525f60ec7"></a>

## max_events property — aws_cloud_watch_receiver.batch / fea8c8d58d52 / 5

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

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

- [max_events_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-05480253834e4a72af481c4e209715baa963c42ebe27dc35f49f8f01f43f9350): complete subsection reference.

<a id="canonical-08972e6997e67a014c0422cb6b58e8984fb0323a802871ffd965351e972d9434"></a>

<a id="canonical-89bca529896db9a672fcdaeddca2c75fd90403df33b499b20e1e3e30ac282842"></a>

## timeout_seconds property — aws_cloud_watch_receiver.batch / fea8c8d58d52 / 6

Type: `"string"`. Computed.

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

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-001.md#canonical-c33bcee763befe91190d25301905e9c262ca3edf10c18cb868b559e88e2ee30f): complete subsection reference.

<a id="canonical-0d978a9e84651814e42da6c3fa09a3547268deeb1083c99a952cdc2b35676034"></a>

## Next pages — aws_cloud_watch_receiver.batch / fea8c8d58d52 / 7

- [aws_cloud_watch_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-2f7de895b09020ece48800f5aa5b03ca9ac203c2ec0c524af28dbf6fa0bed084)
- [aws_cloud_watch_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-05480253834e4a72af481c4e209715baa963c42ebe27dc35f49f8f01f43f9350)
- [aws_cloud_watch_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-001.md#canonical-c33bcee763befe91190d25301905e9c262ca3edf10c18cb868b559e88e2ee30f)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-2f7de895b09020ece48800f5aa5b03ca9ac203c2ec0c524af28dbf6fa0bed084"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1b53f923a97a6d3bc0f74ac45205caba72eb56e66f953e6a72bae8b00d20ada"></a>

## aws_cloud_watch_receiver.batch.max_bytes_disabled — aws_cloud_watch_receiver.batch.max_bytes_disabled / cfe6c939e41f / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e)
- [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-e0af9183ee489c2211fa5ade2ad069999fc1819b782cbe767f43679fd3e67478)
- aws_cloud_watch_receiver.batch.max_bytes_disabled

<a id="canonical-743ebb45bed6ad9279e9ffa0ff45b6c064f171e36f36b6114491fab958711eb9"></a>

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

<a id="canonical-2ba308d64985619cbac583a08e7e6312c8651ae8458074350dba7cc5dd56de2e"></a>

## Direct properties — aws_cloud_watch_receiver.batch.max_bytes_disabled / cfe6c939e41f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-33e9e1768435cc61b1d570d38bf4a115fe55b9509ed5981fe529070ae0aa2198"></a>

## Next pages — aws_cloud_watch_receiver.batch.max_bytes_disabled / cfe6c939e41f / 4

- [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-e0af9183ee489c2211fa5ade2ad069999fc1819b782cbe767f43679fd3e67478)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-05480253834e4a72af481c4e209715baa963c42ebe27dc35f49f8f01f43f9350"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d72bd6acfb47e7bccbd44de62867bc4247374c98c58349f3d53a1670e96ebc87"></a>

## aws_cloud_watch_receiver.batch.max_events_disabled — aws_cloud_watch_receiver.batch.max_events_disabled / a8f7a5aa6cd3 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e)
- [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-e0af9183ee489c2211fa5ade2ad069999fc1819b782cbe767f43679fd3e67478)
- aws_cloud_watch_receiver.batch.max_events_disabled

<a id="canonical-b252eb40bbab88586206c8b7cf7aaada46f72f4f2c09786b3f78fdd09720d9ee"></a>

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

<a id="canonical-9e348455ad8eb74215c65b8fae4ff2b4d9bf309368647682a21daa1fd64ec472"></a>

## Direct properties — aws_cloud_watch_receiver.batch.max_events_disabled / a8f7a5aa6cd3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c22614a4770e236a068bcbf7ae940ac2e88e1325531398b7ecea67dcd34896de"></a>

## Next pages — aws_cloud_watch_receiver.batch.max_events_disabled / a8f7a5aa6cd3 / 4

- [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-e0af9183ee489c2211fa5ade2ad069999fc1819b782cbe767f43679fd3e67478)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-c33bcee763befe91190d25301905e9c262ca3edf10c18cb868b559e88e2ee30f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32634bd9b489bd89b8b37c1c6422bbca9613ffd53c9ae42ebf42ef0dca8e75ba"></a>

## aws_cloud_watch_receiver.batch.timeout_seconds_default — aws_cloud_watch_receiver.batch.timeout_seconds_default / f1553af710ca / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e)
- [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-e0af9183ee489c2211fa5ade2ad069999fc1819b782cbe767f43679fd3e67478)
- aws_cloud_watch_receiver.batch.timeout_seconds_default

<a id="canonical-b6625abf8ed7349dc06839f48b387e7ddd6795d11d592a2da20c5f358d79c690"></a>

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

<a id="canonical-11b92705b4ab9dd944deb1fc54bd4421da24bf24f8a51ec3119c8ddd26d53bb3"></a>

## Direct properties — aws_cloud_watch_receiver.batch.timeout_seconds_default / f1553af710ca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-792d8cfd6d93951d2194215270c7b3b2ac0a1b48f339e62fc71f46cfca068bac"></a>

## Next pages — aws_cloud_watch_receiver.batch.timeout_seconds_default / f1553af710ca / 4

- [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-e0af9183ee489c2211fa5ade2ad069999fc1819b782cbe767f43679fd3e67478)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-8e8aedf1590d974b9d62aae8188b22fed053703badf6e150a9c07fcaf6849f57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6133f38d5297d3ae63652213ba1cf37b4f34ff63309b7da57ffd7bbb90f4e75"></a>

## aws_cloud_watch_receiver.compression — aws_cloud_watch_receiver.compression / e01f3082b98d / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e)
- aws_cloud_watch_receiver.compression

<a id="canonical-9465830781e13d9893dfb02b396655b711827d008d8a17992b98f1543ed52d87"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Upstream description:

Compression Type.

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

<a id="canonical-23e26a536dd826554f589928321da5865f185ddf868375db4f0dbd571ca38934"></a>

## Direct properties — aws_cloud_watch_receiver.compression / e01f3082b98d / 3

- [compression_default](data-sources--global_log_receiver--reference--group-001.md#canonical-9f003994f84a2d1c90ebf6434e7c346c2570caf5611b515ee78fcb7c033fe5c2): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-001.md#canonical-6f7223593d14606be9110b4cb7b6f920b8bccbebd044ed50f32b6fc2dc21eb81): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-001.md#canonical-c0374011381d73515c4b20ffcde7e422e9b466b0cf5c91475d2cdb687cfc7ad4): complete subsection reference.

<a id="canonical-b6bb30841ed0b60797ade148da61812325ffb840b96db0a85345db957f3786b1"></a>

## Next pages — aws_cloud_watch_receiver.compression / e01f3082b98d / 4

- [aws_cloud_watch_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-001.md#canonical-9f003994f84a2d1c90ebf6434e7c346c2570caf5611b515ee78fcb7c033fe5c2)
- [aws_cloud_watch_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-001.md#canonical-6f7223593d14606be9110b4cb7b6f920b8bccbebd044ed50f32b6fc2dc21eb81)
- [aws_cloud_watch_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-001.md#canonical-c0374011381d73515c4b20ffcde7e422e9b466b0cf5c91475d2cdb687cfc7ad4)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-9f003994f84a2d1c90ebf6434e7c346c2570caf5611b515ee78fcb7c033fe5c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c62e46f6a27efde3e126b34e4a93875167974d044b6b8f28630ad1aa296bf886"></a>

## aws_cloud_watch_receiver.compression.compression_default — aws_cloud_watch_receiver.compression.compression_default / e628a0dd261e / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e)
- [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-8e8aedf1590d974b9d62aae8188b22fed053703badf6e150a9c07fcaf6849f57)
- aws_cloud_watch_receiver.compression.compression_default

<a id="canonical-4c3ba4ea5fd5f580b3a4baa08ba4dbf8bdc2c1c0d783c71ea689d95bd36e67f2"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-15d60d4ef098b8c8f2b077d1d769c95aee8da01b9f355090b8bab6ccc11e9508"></a>

## Direct properties — aws_cloud_watch_receiver.compression.compression_default / e628a0dd261e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-154d5b1b93bb7880f32c123f8f2d86c5d5b905178b72af56d9a0cab2af4bc1e1"></a>

## Next pages — aws_cloud_watch_receiver.compression.compression_default / e628a0dd261e / 4

- [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-8e8aedf1590d974b9d62aae8188b22fed053703badf6e150a9c07fcaf6849f57)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-6f7223593d14606be9110b4cb7b6f920b8bccbebd044ed50f32b6fc2dc21eb81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca9a30c90c717cdefb5718bd7a03f32cba8c7ab7cdadb28477221858f4c1efcd"></a>

## aws_cloud_watch_receiver.compression.compression_gzip — aws_cloud_watch_receiver.compression.compression_gzip / 7a2971890d9c / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e)
- [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-8e8aedf1590d974b9d62aae8188b22fed053703badf6e150a9c07fcaf6849f57)
- aws_cloud_watch_receiver.compression.compression_gzip

<a id="canonical-6bae98de3de949c0e2b45aefc1052ef76756895fd1de86a3ef46c7cbe668520c"></a>

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

<a id="canonical-e4c29c934888a15314c82493f23c3c1be7c0dc2602b022f7496a7550f6d2d1fc"></a>

## Direct properties — aws_cloud_watch_receiver.compression.compression_gzip / 7a2971890d9c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da43a73f6b6f7825f69e99b14a0a56a38e931d525cbdba3f3babea30b583e3fb"></a>

## Next pages — aws_cloud_watch_receiver.compression.compression_gzip / 7a2971890d9c / 4

- [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-8e8aedf1590d974b9d62aae8188b22fed053703badf6e150a9c07fcaf6849f57)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-c0374011381d73515c4b20ffcde7e422e9b466b0cf5c91475d2cdb687cfc7ad4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d22107d08167f7215681c8049e4852c84f81ed3bb2f564f134363a1576e6787"></a>

## aws_cloud_watch_receiver.compression.compression_none — aws_cloud_watch_receiver.compression.compression_none / 3bb799a77a11 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-659fcb0f0e6c47c121190f6a4617113f6128ac54b24e886490fecb98d43a2d8e)
- [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-8e8aedf1590d974b9d62aae8188b22fed053703badf6e150a9c07fcaf6849f57)
- aws_cloud_watch_receiver.compression.compression_none

<a id="canonical-06e14a5887e30c610536c4d48ca5d6527e2f4beddde77f4d624adbe4633b7cbe"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-48b56f3fba1bb967cd3fcb8409d6861bef7b006a916a17cd8707f777eae8738a"></a>

## Direct properties — aws_cloud_watch_receiver.compression.compression_none / 3bb799a77a11 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6a710e517479290dc3bca4b2d7519dd2f2c3b89a96d0b4737320d783849e1a46"></a>

## Next pages — aws_cloud_watch_receiver.compression.compression_none / 3bb799a77a11 / 4

- [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-8e8aedf1590d974b9d62aae8188b22fed053703badf6e150a9c07fcaf6849f57)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-1e93724fa97f305e8285ab20a3de95bcda72cb3dd4afba2b5d1f659e5b82aa39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45416b011bd78cb07df56dbdf9cc300ea3d3b783d8a8ad05be8678b08e47746a"></a>

## azure_event_hubs_receiver — azure_event_hubs_receiver / e70749346071 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- azure_event_hubs_receiver

<a id="canonical-5f8d310c2746fa4c17c4a70016a2a5fe97191aed64d1cd597a98b1a3d89f2fee"></a>

Type: `"single"`. Computed.

Azure Event Hubs Configuration for Global Log Receiver.

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

<a id="canonical-16c303ac9593b6b03b6b6a952a4ef8ce944d31057586ab2f26ad752af376dd84"></a>

## Direct properties — azure_event_hubs_receiver / e70749346071 / 3

- [connection_string](data-sources--global_log_receiver--reference--group-001.md#canonical-abb40162489c5d60ca4f8bac4d87a6a712d58b8626d6e323895171c9a7273b43): complete subsection reference.

<a id="canonical-ef0d165096b5060e460a8e50184d826db12d6d88bfb67c322554a23b8fa4f960"></a>

<a id="canonical-5ff44fd6443e4dce9f3796be2f1718f024238aece1c4644f9f721658fc7365f8"></a>

## instance property — azure_event_hubs_receiver / e70749346071 / 4

Type: `"string"`. Computed.

Event Hubs Instance name into which logs should be stored.

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

<a id="canonical-e6190f350141c41cfbe65991353ebc370fa046012b133b20c6a779b96aba5f0a"></a>

<a id="canonical-a33ccffa32f6a8fba8a51c5e8f200d8113b2f0b7817bcd366c23493930cd0057"></a>

## namespace property — azure_event_hubs_receiver / e70749346071 / 5

Type: `"string"`. Computed.

Event Hubs Namespace is namespace with instance into which logs should be stored.

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

<a id="canonical-521bbd43708398624f67b610d8f4f6c4030599d1893bdf29bfbb0a64f299e7e8"></a>

## Next pages — azure_event_hubs_receiver / e70749346071 / 6

- [azure_event_hubs_receiver.connection_string](data-sources--global_log_receiver--reference--group-001.md#canonical-abb40162489c5d60ca4f8bac4d87a6a712d58b8626d6e323895171c9a7273b43)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-abb40162489c5d60ca4f8bac4d87a6a712d58b8626d6e323895171c9a7273b43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73bb9c2f27eca4e444e41436fbaeca352307f149d16470e9e930e2375e3224bb"></a>

## azure_event_hubs_receiver.connection_string — azure_event_hubs_receiver.connection_string / 49324f350e26 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1e93724fa97f305e8285ab20a3de95bcda72cb3dd4afba2b5d1f659e5b82aa39)
- azure_event_hubs_receiver.connection_string

<a id="canonical-3f4267535f141fd1a61562c1e0281c646018d972811f37bc7fc41881d1335b8e"></a>

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

<a id="canonical-f515fb451cea04389d386707a839d00831d99bf30ecea7f61c31b786703ef88b"></a>

## Direct properties — azure_event_hubs_receiver.connection_string / 49324f350e26 / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-001.md#canonical-4bf3d64de6b35b576a03865ec80dfbe3196cfd77b49c129ad47db2723d900867): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-001.md#canonical-a9e68274fd0c831a766d03c2bc3a9f49f39156a3840ae53d17ccc7f08b00a96d): complete subsection reference.

<a id="canonical-9de40996d5c924fe423e46e2ba3a138ed1c8dcec99c65bfb2ab66848fe714c16"></a>

## Next pages — azure_event_hubs_receiver.connection_string / 49324f350e26 / 4

- [azure_event_hubs_receiver.connection_string.blindfold_secret_info](data-sources--global_log_receiver--reference--group-001.md#canonical-4bf3d64de6b35b576a03865ec80dfbe3196cfd77b49c129ad47db2723d900867)
- [azure_event_hubs_receiver.connection_string.clear_secret_info](data-sources--global_log_receiver--reference--group-001.md#canonical-a9e68274fd0c831a766d03c2bc3a9f49f39156a3840ae53d17ccc7f08b00a96d)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1e93724fa97f305e8285ab20a3de95bcda72cb3dd4afba2b5d1f659e5b82aa39)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-4bf3d64de6b35b576a03865ec80dfbe3196cfd77b49c129ad47db2723d900867"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-334b8f4d3f0362d4555a1c6bc636269de6d334dc8ec9922fdd8876b339d8816e"></a>

## azure_event_hubs_receiver.connection_string.blindfold_secret_info — azure_event_hubs_receiver.connection_string.blindfold_secret_info / 7c69d26c0c34 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1e93724fa97f305e8285ab20a3de95bcda72cb3dd4afba2b5d1f659e5b82aa39)
- [azure_event_hubs_receiver.connection_string](data-sources--global_log_receiver--reference--group-001.md#canonical-abb40162489c5d60ca4f8bac4d87a6a712d58b8626d6e323895171c9a7273b43)
- azure_event_hubs_receiver.connection_string.blindfold_secret_info

<a id="canonical-6f67f1ba265aaabb7c65c4e12e588552289d7a681048453950c98228d176f93c"></a>

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

<a id="canonical-aa8a2fe7c9a0e884c1677843c66c5b21219f493d8b8d4af71fa1ca5bd03a515d"></a>

## Direct properties — azure_event_hubs_receiver.connection_string.blindfold_secret_info / 7c69d26c0c34 / 3

<a id="canonical-9d3fb97a20e0287fc8733cce7e99b615b43888071665c2e5ab918eb1c845531f"></a>

<a id="canonical-d2bad6d8ab553a80ea3dd720863dedd84c7769c6ea6e4c8de4c8d16fbc30b414"></a>

## decryption_provider property — azure_event_hubs_receiver.connection_string.blindfold_secret_info / 7c69d26c0c34 / 4

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

<a id="canonical-a96d2751c331a5b27a2d669e88f3e65f7e58faaead83ad3d50b7ab051d4cb5de"></a>

<a id="canonical-a6df9b574375c7db45064a9113a502a0eb4abb8a971197f25f2f370c1f914c36"></a>

## location property — azure_event_hubs_receiver.connection_string.blindfold_secret_info / 7c69d26c0c34 / 5

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

<a id="canonical-2286724c776f3ccadca1619b723ae477ebb57504a0d5634a76282a70ef2b54af"></a>

<a id="canonical-5cbbb641193583d26e6eda157e271b12e6f4b6438753514212868881fbfa5638"></a>

## store_provider property — azure_event_hubs_receiver.connection_string.blindfold_secret_info / 7c69d26c0c34 / 6

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

<a id="canonical-429e19994d8a4e0edf22cf205a69d1400e55ab1d55b81974f080509fa18e7d30"></a>

## Next pages — azure_event_hubs_receiver.connection_string.blindfold_secret_info / 7c69d26c0c34 / 7

- [azure_event_hubs_receiver.connection_string](data-sources--global_log_receiver--reference--group-001.md#canonical-abb40162489c5d60ca4f8bac4d87a6a712d58b8626d6e323895171c9a7273b43)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-a9e68274fd0c831a766d03c2bc3a9f49f39156a3840ae53d17ccc7f08b00a96d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-331ff87993deace29d03bdabab019efa1fb3e496a28c7ff730d597fc98042671"></a>

## azure_event_hubs_receiver.connection_string.clear_secret_info — azure_event_hubs_receiver.connection_string.clear_secret_info / 1f7e2ccc46f6 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1e93724fa97f305e8285ab20a3de95bcda72cb3dd4afba2b5d1f659e5b82aa39)
- [azure_event_hubs_receiver.connection_string](data-sources--global_log_receiver--reference--group-001.md#canonical-abb40162489c5d60ca4f8bac4d87a6a712d58b8626d6e323895171c9a7273b43)
- azure_event_hubs_receiver.connection_string.clear_secret_info

<a id="canonical-5cbf8a4e79f70be17a7fa9697fa49e4424bbf3f87780af9aeae2cb93063240fe"></a>

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

<a id="canonical-ac0058114bfb1b14cfc6bbf6e9618c8f5c42bc0e10ac1ffb4085690f65774a13"></a>

## Direct properties — azure_event_hubs_receiver.connection_string.clear_secret_info / 1f7e2ccc46f6 / 3

<a id="canonical-fd3c0d07adece5898726453ec277bee236d4738668cda42d54446fca120defd4"></a>

<a id="canonical-678d6834c8918a2a5212075a1dac57dcbe956964fd00846707510e95cda58c2f"></a>

## provider_ref property — azure_event_hubs_receiver.connection_string.clear_secret_info / 1f7e2ccc46f6 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-937ccedaf14e3024919b28bf2f3dbb59916b5bfb0f12d8f2b2c5fafa3a8e2867"></a>

<a id="canonical-d52f661879d4ef49b2740990db5909d3e38d119d052ade96f794a6e704ac37e8"></a>

## url property — azure_event_hubs_receiver.connection_string.clear_secret_info / 1f7e2ccc46f6 / 5

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

<a id="canonical-24ec03ec942e849bbe5b851a78b8460cff1e60207f24e8f156787acc4e928123"></a>

## Next pages — azure_event_hubs_receiver.connection_string.clear_secret_info / 1f7e2ccc46f6 / 6

- [azure_event_hubs_receiver.connection_string](data-sources--global_log_receiver--reference--group-001.md#canonical-abb40162489c5d60ca4f8bac4d87a6a712d58b8626d6e323895171c9a7273b43)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9772637760aaee6ceba5727ea72ccf4a84730c22a381f697219c4c18603d3d43"></a>

## azure_receiver — azure_receiver / aea6dfd433d2 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- azure_receiver

<a id="canonical-33d0516820774d65b0cb95a87591491c90c8a78ee4996ddde7ec41aba356117d"></a>

Type: `"single"`. Computed.

Azure Blob Configuration for Global Log Receiver.

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

<a id="canonical-13d069caf26a0fd73ce635f4e85756522e2f52abd26f91e55c9eab0182254c57"></a>

## Direct properties — azure_receiver / aea6dfd433d2 / 3

- [batch](data-sources--global_log_receiver--reference--group-001.md#canonical-94b5d6cae427d6372c732fb99af31e900e01b115c060f422b7aab9e92d0a220e): complete subsection reference.

- [compression](data-sources--global_log_receiver--reference--group-002.md#canonical-71a878ee1d36ee303f1eda176a382c17ae04e2922f1d2d38102d6f0138b7d5fe): complete subsection reference.

- [connection_string](data-sources--global_log_receiver--reference--group-002.md#canonical-b1d4378fc6ac36df67d7e84d8d66a74cc58a0fe156db7dbb626ca26b57bbe164): complete subsection reference.

<a id="canonical-5a682b0e414665ce244919459fa07d09a594838086df125e77e5fefa822e8304"></a>

<a id="canonical-28f15aca12045ec0113e4ce673c44ebc2b157bd3548a8e32f9d857786c33757d"></a>

## container_name property — azure_receiver / aea6dfd433d2 / 4

Type: `"string"`. Computed.

Container Name is the name of the container into which logs should be stored.

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

- [filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-67e28dd59f9f134efb8387a9fd4b71f4c99385a9db5a25fa28ea0da0aa22330f): complete subsection reference.

<a id="canonical-5ba4629ade5413fd8dfc4485b79b829eb20807d5afb4db3116f80d12bba6cf6e"></a>

## Next pages — azure_receiver / aea6dfd433d2 / 5

- [azure_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-94b5d6cae427d6372c732fb99af31e900e01b115c060f422b7aab9e92d0a220e)
- [azure_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-71a878ee1d36ee303f1eda176a382c17ae04e2922f1d2d38102d6f0138b7d5fe)
- [azure_receiver.connection_string](data-sources--global_log_receiver--reference--group-002.md#canonical-b1d4378fc6ac36df67d7e84d8d66a74cc58a0fe156db7dbb626ca26b57bbe164)
- [azure_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-67e28dd59f9f134efb8387a9fd4b71f4c99385a9db5a25fa28ea0da0aa22330f)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-94b5d6cae427d6372c732fb99af31e900e01b115c060f422b7aab9e92d0a220e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5a3f630ffba1181b2903ee05cdd4bc86c7bd13762137c4ee455569e32bde8dc"></a>

## azure_receiver.batch — azure_receiver.batch / 70ab8873061e / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- azure_receiver.batch

<a id="canonical-bba588ffa53013eb0467f505e13897fadcf8f3f2ba3467b76226c095abdfc527"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

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

<a id="canonical-02a77d8771c0e3c5868cbccbbaef2eb5803c13204fa57cc08ca513500058aafb"></a>

## Direct properties — azure_receiver.batch / 70ab8873061e / 3

<a id="canonical-ba331a499368037efaf4c6958b82d55fbc9fd9462f2c62c6a9eb202f1861b076"></a>

<a id="canonical-8b026c0bbd53bf584e0503541ad5999ebefcc97e7fd419a0f8b4e4bdbce27a7c"></a>

## max_bytes property — azure_receiver.batch / 70ab8873061e / 4

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

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

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-33570fa762d51df98e4c40a44111d319e82f7d260eb9627f6e951d83de01e3f3): complete subsection reference.

<a id="canonical-a81532423032e39a56057c03dd4bddcdf357669c88aea1ce8c81a10f303687a6"></a>

<a id="canonical-e00b6d38844c4db874583974e7d1cf3cf3b8af4e46448f8755d701d4ca67a19f"></a>

## max_events property — azure_receiver.batch / 70ab8873061e / 5

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

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

- [max_events_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-b376c86bc9b6981145ff733476ef5f7b6783709693e6d277323033fb482368a7): complete subsection reference.

<a id="canonical-d43735d6688c92d67acce19fa8432bfd9a8e99cf5ef24b6e6b0a9ba8bd4566dc"></a>

<a id="canonical-45d3ea1ac18e580876624a8b34e5dc78dac6805a6bc61d24c8532ff0f299da6f"></a>

## timeout_seconds property — azure_receiver.batch / 70ab8873061e / 6

Type: `"string"`. Computed.

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

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-def1ef14667fdcb585f83fae4b18ed18e48417029785d272b1f408aa4e16a157): complete subsection reference.

<a id="canonical-35bc004006356b31e57f49d137d4492d022fa4b2ebf505b223307e766de58b5b"></a>

## Next pages — azure_receiver.batch / 70ab8873061e / 7

- [azure_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-33570fa762d51df98e4c40a44111d319e82f7d260eb9627f6e951d83de01e3f3)
- [azure_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-b376c86bc9b6981145ff733476ef5f7b6783709693e6d277323033fb482368a7)
- [azure_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-def1ef14667fdcb585f83fae4b18ed18e48417029785d272b1f408aa4e16a157)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-33570fa762d51df98e4c40a44111d319e82f7d260eb9627f6e951d83de01e3f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59a4fe794bad4e17bc80258ce2a6d3917333a511988166e13545960d4afb5ff6"></a>

## azure_receiver.batch.max_bytes_disabled — azure_receiver.batch.max_bytes_disabled / ec7530eb1ecd / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [azure_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-94b5d6cae427d6372c732fb99af31e900e01b115c060f422b7aab9e92d0a220e)
- azure_receiver.batch.max_bytes_disabled

<a id="canonical-ece2b962d6b190267bd988002a60311ca84c5613e7553673a1ac20610dc9f719"></a>

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

<a id="canonical-dcf3fa57ff02fe43423fcdf7de13bf34c129eab55b52c7d5019f1fd4b060e276"></a>

## Direct properties — azure_receiver.batch.max_bytes_disabled / ec7530eb1ecd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ef2c7c2492b38030c63b0e58098d6f940a8d7be1658fa5ef8476f888418b609f"></a>

## Next pages — azure_receiver.batch.max_bytes_disabled / ec7530eb1ecd / 4

- [azure_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-94b5d6cae427d6372c732fb99af31e900e01b115c060f422b7aab9e92d0a220e)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-b376c86bc9b6981145ff733476ef5f7b6783709693e6d277323033fb482368a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ca0425f2e1f7c167b7bb4346e5b56b43fa824faa1c0614a566d7e263a7f79cc"></a>

## azure_receiver.batch.max_events_disabled — azure_receiver.batch.max_events_disabled / 73ca8935150f / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [azure_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-94b5d6cae427d6372c732fb99af31e900e01b115c060f422b7aab9e92d0a220e)
- azure_receiver.batch.max_events_disabled

<a id="canonical-22db6b551a8506d2080636f35d98a33da53e95858c5d6c811af3c13c551e15d6"></a>

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

<a id="canonical-1caa8e6408816e51f895809fcf750552e3a40f3a7050c3f2ca3947406d855a35"></a>

## Direct properties — azure_receiver.batch.max_events_disabled / 73ca8935150f / 3

This is an empty object or choice marker. It has no direct properties.
