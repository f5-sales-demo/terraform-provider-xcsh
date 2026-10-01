---
page_title: "xcsh_app_setting reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_setting reference."
---

# xcsh_app_setting reference

<a id="canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1307ad176ad3e80dae6b709945d3fc728304c88c7c12a99775bbe87df06b26cb"></a>

## Property reference — Property reference / 119fb504d73f / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- Property reference

<a id="canonical-e3aab9287993c73b3ab42f94e63ec5eeae0fd1d0d4214ac6b8ba8979e799fe83"></a>

## Direct properties — Property reference / 119fb504d73f / 3

<a id="canonical-3385610e1c3d59b1e9ff6810c2211ec9cf88cdfc4998be4eb961903cda6f061f"></a>

<a id="canonical-d3101e4b5b116698cb65ee8ac9c7efe6fb48cdd9be578f65631afe49021a90f7"></a>

## annotations property — Property reference / 119fb504d73f / 4

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

- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356): complete subsection reference.

<a id="canonical-0a313fc5537582d0076a0d45b22459bf4d4638da2c750fde76615d992eed26a7"></a>

<a id="canonical-ff5414067aea844ca20aafd88ab783479abf810f4beb0ac56bacf3f5516c5e11"></a>

## description property — Property reference / 119fb504d73f / 5

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

<a id="canonical-d905c83a40fd141491580190a25c538217cb7cf9f8fcc5b640aa1180d7692ba4"></a>

<a id="canonical-959540d87de0e2bd7916dd89e38791bfe4e1ae6bfd426d0691a5940f21e73d8d"></a>

## disable property — Property reference / 119fb504d73f / 6

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

<a id="canonical-30475c1af09a11d363618bc9b98f2f428d6d0e5a986d4901a08b2fc5127de1b7"></a>

<a id="canonical-5ef4defb71bb1fb6978eb8edf3ee72c58b7bf260ffad5e50f695562451ccb243"></a>

## id property — Property reference / 119fb504d73f / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-d8d54e22ea7f264b13b8b259710eb0e6ae21fcf44962e9a4aef114128bf920ff"></a>

<a id="canonical-8a5770963a64fa5bd891013c1b31071b5649b507a7275cb5f9dec48bdb0bb6ba"></a>

## labels property — Property reference / 119fb504d73f / 8

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

<a id="canonical-00c231b07778d1680d5cb96658ffc067df4b8478963dfb17f9a0b970b254071d"></a>

<a id="canonical-98540c8ad3ba3616f3a99936abb54de48cda8921e3fe2428364c770eaed0fa0f"></a>

## name property — Property reference / 119fb504d73f / 9

Type: `"string"`. Required.

Name of the App Setting. Must be unique within the namespace.

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

<a id="canonical-9d95c5884ef44b706a40c71f3182c7f84bccd7dd734d8c88574e46dce3a623f5"></a>

<a id="canonical-e6bad3b7c90db16d3feba9d83272702adda5a13abfd66d841b8fda701242d3c6"></a>

## namespace property — Property reference / 119fb504d73f / 10

Type: `"string"`. Required.

Namespace where the App Setting is created.

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

- [timeouts](resources--app_setting--reference--group-001.md#canonical-b6bf006bee2af16444ad1d452e09cf84122225ae3f0f01656f31a62f5a1b158e): complete subsection reference.

<a id="canonical-42b0bcb7533a85d63036de36f454ad1a74a749f8575a486860694656365dd443"></a>

## All schema paths — Property reference / 119fb504d73f / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--app_setting--reference--group-001.md#canonical-3385610e1c3d59b1e9ff6810c2211ec9cf88cdfc4998be4eb961903cda6f061f) |
| `app_type_settings` | [app_type_settings](resources--app_setting--reference--group-001.md#canonical-4a47deb61957e28de040aa75f97b8e532d63e53407e2b8cf5d25c67111712663) |
| `app_type_settings.app_type_ref` | [app_type_settings.app_type_ref](resources--app_setting--reference--group-001.md#canonical-b2ce6ab4dcceaf2a3176ec5982a048dde53da5983e8ee99d1707575657a12a04) |
| `app_type_settings.app_type_ref.kind` | [app_type_settings.app_type_ref.kind](resources--app_setting--reference--group-001.md#canonical-361203a0f136781141afedeecaa1ba20eac8dcb7b80a19a300eb26f537d180df) |
| `app_type_settings.app_type_ref.name` | [app_type_settings.app_type_ref.name](resources--app_setting--reference--group-001.md#canonical-db0324bc5e8efcbcd3a9d068d08e84fab326a0a4109707402f0f4b6380eb00dd) |
| `app_type_settings.app_type_ref.namespace` | [app_type_settings.app_type_ref.namespace](resources--app_setting--reference--group-001.md#canonical-f3594ddfee694c85354e6c09259965524a1951a6c3f2532761a5917641c55270) |
| `app_type_settings.app_type_ref.tenant` | [app_type_settings.app_type_ref.tenant](resources--app_setting--reference--group-001.md#canonical-c03d324bd88c2622d6d29c384a951ef402cdc6b480241e76cacdddd4f8744cdc) |
| `app_type_settings.app_type_ref.uid` | [app_type_settings.app_type_ref.uid](resources--app_setting--reference--group-001.md#canonical-01df61bd5cd86af3b12596b7598ab79355c79a4399d649c8753cce35229846db) |
| `app_type_settings.business_logic_markup_setting` | [app_type_settings.business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-b64e891dd6a371ecda65aa5f99d0dbf97770f1495b09086759bc4a0737e759f1) |
| `app_type_settings.business_logic_markup_setting.disable_spec` | [app_type_settings.business_logic_markup_setting.disable_spec](resources--app_setting--reference--group-001.md#canonical-383223c3a40cb8e8fc783e8a5a1249e99d025ed8373fdcbb34c16325d53d7076) |
| `app_type_settings.business_logic_markup_setting.enable` | [app_type_settings.business_logic_markup_setting.enable](resources--app_setting--reference--group-001.md#canonical-2fc85b63e14691a05aefe8bfaf50f5a28e929f76de90ab3a68420c472ab552a1) |
| `app_type_settings.timeseries_analyses_setting` | [app_type_settings.timeseries_analyses_setting](resources--app_setting--reference--group-001.md#canonical-a501ea728b97e59308d75d3c20acfd81d4ddf328a4563e8623777b5246cbadc2) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors` | [app_type_settings.timeseries_analyses_setting.metric_selectors](resources--app_setting--reference--group-001.md#canonical-a71bcfcc974b495b3164dd42d99a4053b419ff08d3a71c704913edcf1d425634) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors.metric` | [app_type_settings.timeseries_analyses_setting.metric_selectors.metric](resources--app_setting--reference--group-001.md#canonical-314c4e29ed7b06a649ec31d7067308a53c06c979e998ae5d49c5f3fba371a5d7) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source` | [app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source](resources--app_setting--reference--group-001.md#canonical-2634bfcd601e57400fac76e183849eb34bd3560eda589702dbdaa2ae8b31d75e) |
| `app_type_settings.user_behavior_analysis_setting` | [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a95b9ddfb094890c9f9d4bd360e2e9d2838b6ca73adddd1a15b0482c169b60fa) |
| `app_type_settings.user_behavior_analysis_setting.disable_detection` | [app_type_settings.user_behavior_analysis_setting.disable_detection](resources--app_setting--reference--group-001.md#canonical-f64fc6655935330c8016584e0be182277befa8bc2730f6277b457eadb95bfb9f) |
| `app_type_settings.user_behavior_analysis_setting.disable_learning` | [app_type_settings.user_behavior_analysis_setting.disable_learning](resources--app_setting--reference--group-001.md#canonical-5c42e09df28c12b1fe9c355ecf32d42c5c0a9d0f57e30409226c42f99f3b78ca) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection` | [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-400ce4b40c3733e9121789b57deaabcd432d2a1033bf702044aa90d575623721) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic` | [app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic](resources--app_setting--reference--group-001.md#canonical-4a39622499de50363929a7476c736c7612f5baf2bba156b63c7309db131c9406) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period` | [app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period](resources--app_setting--reference--group-001.md#canonical-61b8767049c9f95b870fdc69152b5a96af558405f63e7a417dbc85d7aebcacee) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection](resources--app_setting--reference--group-001.md#canonical-d3e66e10d555c43738de98344ca93b58fc55f982786bcaca633ade01b2bd8c22) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity](resources--app_setting--reference--group-001.md#canonical-ff3dc5f8c0ac27dd7a777dcd921c91dd448be7e9a05aa8c109d129ded2bbfd81) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity](resources--app_setting--reference--group-001.md#canonical-bae5b5aef22977b9c512461aa6e3ee05c6af5a856153eff6dc008bb9153ca3f5) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity](resources--app_setting--reference--group-001.md#canonical-055aaee013ee47c8da952ca757cb6f16ca6b62cb56a2703809d3b1924c9227a5) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation](resources--app_setting--reference--group-001.md#canonical-64a65a9a7bbfa2d9e9b952174f3345ace232b9482ecb0b4de421e1af310cebc6) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity](resources--app_setting--reference--group-001.md#canonical-892729e56e3da5eae13d868cd170e2de916118166a357bc9906c670011e9cddc) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit](resources--app_setting--reference--group-001.md#canonical-7c89ddf27eda6b15df7c5864bc8a086bfd5293f780dbc60141aa312c68b28032) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity](resources--app_setting--reference--group-001.md#canonical-be51d8aa950b355900f47ebabddaf65e23a11f1c1d754068bc70ac3870194606) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity](resources--app_setting--reference--group-001.md#canonical-de451c060ed5349afd195a16ec8e130454450b41f6c19cd7fe823a4ee5f88be6) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity](resources--app_setting--reference--group-001.md#canonical-1ab4f3403976fd4bf8429eb30cf6ac5b76c2f04c23c364e6cd2934bb0d04459a) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold](resources--app_setting--reference--group-001.md#canonical-de84a99b005cd2b5a07c425439896a3d5d5ca0af3ce43554b63304324c5922b0) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity](resources--app_setting--reference--group-001.md#canonical-c99a8aec36705697e74045bc659bbbe7eca2af4fbf3c9ff1b5a7b6346a36ae8c) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold](resources--app_setting--reference--group-001.md#canonical-139d367f036c04fa7acc9fee109bb8286137838447bbcaec2d540ffe9011e122) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation](resources--app_setting--reference--group-001.md#canonical-8ed0d496152743f966d0a70d526527f0aae52c14e690190561c8d72341011def) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-dca5763c54c675063f514c307eb53856f11473731ee60222923ff5efc618a09f) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high](resources--app_setting--reference--group-001.md#canonical-bcfb9800a9427a77deb06ed1ee85860c08a4a48dcbde20b3989de6d43d39f186) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low](resources--app_setting--reference--group-001.md#canonical-0bafa5b6d66dc599eef7ac1deaa2ae93cbece751a4003026d8341f5dba55ac98) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium](resources--app_setting--reference--group-001.md#canonical-822b460b996ee8dfa374049717aab6b79f7ea0942c3c94c44377a32bee843771) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom](resources--app_setting--reference--group-001.md#canonical-83f0ddf8f8bdab5035a9e32c2ea49aa39db231db13099c82938afe95645a2791) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold](resources--app_setting--reference--group-001.md#canonical-836064832f5f292f9c58d2f3d94c65236e426a5b36144dd6065fe9bb001047fe) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit](resources--app_setting--reference--group-001.md#canonical-95325c36050af5c6f4912bbb941ac7b6b9c84f49b4859850c694416af3ebcfb7) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity](resources--app_setting--reference--group-001.md#canonical-c8ffda23bce7d985e23cc069c96a6397521ea01e4ec5fa045aa15720243eebf4) |
| `app_type_settings.user_behavior_analysis_setting.enable_learning` | [app_type_settings.user_behavior_analysis_setting.enable_learning](resources--app_setting--reference--group-001.md#canonical-d07699123055523581526a89343a1cc3372bde8066f4e63124d0e7701bcc5c13) |
| `description` | [description](resources--app_setting--reference--group-001.md#canonical-0a313fc5537582d0076a0d45b22459bf4d4638da2c750fde76615d992eed26a7) |
| `disable` | [disable](resources--app_setting--reference--group-001.md#canonical-d905c83a40fd141491580190a25c538217cb7cf9f8fcc5b640aa1180d7692ba4) |
| `id` | [id](resources--app_setting--reference--group-001.md#canonical-30475c1af09a11d363618bc9b98f2f428d6d0e5a986d4901a08b2fc5127de1b7) |
| `labels` | [labels](resources--app_setting--reference--group-001.md#canonical-d8d54e22ea7f264b13b8b259710eb0e6ae21fcf44962e9a4aef114128bf920ff) |
| `name` | [name](resources--app_setting--reference--group-001.md#canonical-00c231b07778d1680d5cb96658ffc067df4b8478963dfb17f9a0b970b254071d) |
| `namespace` | [namespace](resources--app_setting--reference--group-001.md#canonical-9d95c5884ef44b706a40c71f3182c7f84bccd7dd734d8c88574e46dce3a623f5) |
| `timeouts` | [timeouts](resources--app_setting--reference--group-001.md#canonical-c8788ad92ab8c25486b501baacbdcd9e45d2481c2c8fc852221bed8e3b0136fc) |
| `timeouts.create` | [timeouts.create](resources--app_setting--reference--group-001.md#canonical-717ca7cdd15db15a6712df2c251a0c1a24bc87603d29d5e55b63bfe9f2f53300) |
| `timeouts.delete` | [timeouts.delete](resources--app_setting--reference--group-001.md#canonical-ac5a7cf64f4662deef75a4e292cf0acbe2c60880484191a4c3777070f830f3d7) |
| `timeouts.read` | [timeouts.read](resources--app_setting--reference--group-001.md#canonical-7faee6be1bb2b66561873e366d3961dd72a70b05a4bff752225813316301c9f8) |
| `timeouts.update` | [timeouts.update](resources--app_setting--reference--group-001.md#canonical-333fc64ee989c59e99aec03be8a2b65d9f516fe2dc691a39f047558218c29e22) |

<a id="canonical-4efd248d419f59dad81c90432350d90bc2ce74da1a66bb48e7843cc8251a7108"></a>

## Next pages — Property reference / 119fb504d73f / 12

- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [timeouts](resources--app_setting--reference--group-001.md#canonical-b6bf006bee2af16444ad1d452e09cf84122225ae3f0f01656f31a62f5a1b158e)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b343f24e9251473f93b1f888656c1445ada93afe2974de0fdb98312c0be5a1d"></a>

## app_type_settings — app_type_settings / a30f829f02e8 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- app_type_settings

<a id="canonical-4a47deb61957e28de040aa75f97b8e532d63e53407e2b8cf5d25c67111712663"></a>

Type: `"object"`. list nested block, Optional.

List of settings to enable for each AppType, given instance of AppType Exist in this Namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("app_type_ref")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
app_type_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-c84c90cdae7e4e37d051b846bcbd4c22171ddd2d35a76163c5aaa96f3cb2d52a"></a>

## Direct properties — app_type_settings / a30f829f02e8 / 3

- [app_type_ref](resources--app_setting--reference--group-001.md#canonical-0d4c9492f1c2f847ffad7d10a3e91cfc15ce7d81c54d26e793735033ab0c2025): complete subsection reference.

- [business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-36a2be0a9b32b6297811c65c9d994ea33bd57de3ad674979dbf54de31e625e0e): complete subsection reference.

- [timeseries_analyses_setting](resources--app_setting--reference--group-001.md#canonical-8ef1350a06bca9d5f9e3014dd9f61b2f79fb8010643ce097d12de338b62a4148): complete subsection reference.

- [user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906): complete subsection reference.

<a id="canonical-b1980c5e79e9aa4029de6b45eb8a80a26778df2ca27be440abafd7cbe5d83c52"></a>

## Next pages — app_type_settings / a30f829f02e8 / 4

- [app_type_settings.app_type_ref](resources--app_setting--reference--group-001.md#canonical-0d4c9492f1c2f847ffad7d10a3e91cfc15ce7d81c54d26e793735033ab0c2025)
- [app_type_settings.business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-36a2be0a9b32b6297811c65c9d994ea33bd57de3ad674979dbf54de31e625e0e)
- [app_type_settings.timeseries_analyses_setting](resources--app_setting--reference--group-001.md#canonical-8ef1350a06bca9d5f9e3014dd9f61b2f79fb8010643ce097d12de338b62a4148)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-0d4c9492f1c2f847ffad7d10a3e91cfc15ce7d81c54d26e793735033ab0c2025"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd61a1c1f0528a916956f423bf8ff3983518b82aab68b0bcd0970ca0da2499c8"></a>

## app_type_settings.app_type_ref — app_type_settings.app_type_ref / 3158c3df9fa5 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- app_type_settings.app_type_ref

<a id="canonical-b2ce6ab4dcceaf2a3176ec5982a048dde53da5983e8ee99d1707575657a12a04"></a>

Type: `"object"`. list nested block, Optional.

The AppType of App instance in current Namespace. Associating an AppType reference, will enable
analysis on this instance's generated data.

Upstream description:

The AppType of App instance in current Namespace. Associating an AppType reference, will enable
analysis on this instance's generated data.

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
app_type_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-fdf90cc190d03bdb4473fe86a5fda97c71924826a024513d4fb6ef87686ca281"></a>

## Direct properties — app_type_settings.app_type_ref / 3158c3df9fa5 / 3

<a id="canonical-361203a0f136781141afedeecaa1ba20eac8dcb7b80a19a300eb26f537d180df"></a>

<a id="canonical-a4b67de43c0cf1fe0475f72e54eaa73b0b9505f8fda1ec0466ff29789ee2441c"></a>

## kind property — app_type_settings.app_type_ref / 3158c3df9fa5 / 4

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

<a id="canonical-db0324bc5e8efcbcd3a9d068d08e84fab326a0a4109707402f0f4b6380eb00dd"></a>

<a id="canonical-9cb8db538c02f8c825a2f2fbf0eb14a16a62fe5c795443883499b296e0c430bf"></a>

## name property — app_type_settings.app_type_ref / 3158c3df9fa5 / 5

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

<a id="canonical-f3594ddfee694c85354e6c09259965524a1951a6c3f2532761a5917641c55270"></a>

<a id="canonical-5eb319fcb471e8bcab096216e8d2e2933a1100c4a12974fb9f9b29df5eb8f54a"></a>

## namespace property — app_type_settings.app_type_ref / 3158c3df9fa5 / 6

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

<a id="canonical-c03d324bd88c2622d6d29c384a951ef402cdc6b480241e76cacdddd4f8744cdc"></a>

<a id="canonical-0bbf4c3b97a21f0f905855ff08e32ade557117179a0ffc4a273aaae95a8e7d2e"></a>

## tenant property — app_type_settings.app_type_ref / 3158c3df9fa5 / 7

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

<a id="canonical-01df61bd5cd86af3b12596b7598ab79355c79a4399d649c8753cce35229846db"></a>

<a id="canonical-74e885162c4ce0d294189fb80ee2de62ee04c20e64190bede1f5d64699ed5f39"></a>

## uid property — app_type_settings.app_type_ref / 3158c3df9fa5 / 8

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

<a id="canonical-b4a21e1fd483958975490eb8d6970f34f1f962c98e4086aa68335dfffc79f36c"></a>

## Next pages — app_type_settings.app_type_ref / 3158c3df9fa5 / 9

- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-36a2be0a9b32b6297811c65c9d994ea33bd57de3ad674979dbf54de31e625e0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-670ab7cd9b4b4fd0c1d291194861d9bca31e6bc023447ffa1cbde9134dbac46f"></a>

## app_type_settings.business_logic_markup_setting — app_type_settings.business_logic_markup_setting / 171ec0c47c79 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- app_type_settings.business_logic_markup_setting

<a id="canonical-b64e891dd6a371ecda65aa5f99d0dbf97770f1495b09086759bc4a0737e759f1"></a>

Type: `"object"`. single nested block, Optional.

Settings specifying how API Discovery will be performed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-learn_from_namespace": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
business_logic_markup_setting {
  # Configure direct properties listed below.
}
```

<a id="canonical-f4d6636b017937727101b5f62b032f9f395e44eb06a69ef4d0e4c0323a160637"></a>

## Direct properties — app_type_settings.business_logic_markup_setting / 171ec0c47c79 / 3

- [disable_spec](resources--app_setting--reference--group-001.md#canonical-934b73d092bbd12501a4fd3387ac3088227286827467d0f12e57d422211b051d): complete subsection reference.

- [enable](resources--app_setting--reference--group-001.md#canonical-77a357dbae85320abf23ad5d68d07f69b64e1d49943c2434a0de05cddddd25e3): complete subsection reference.

<a id="canonical-3388a81336377714fb7025e39cdd0faa0cb0d5e57606f95fa3aa6b4f9402eabb"></a>

## Next pages — app_type_settings.business_logic_markup_setting / 171ec0c47c79 / 4

- [app_type_settings.business_logic_markup_setting.disable_spec](resources--app_setting--reference--group-001.md#canonical-934b73d092bbd12501a4fd3387ac3088227286827467d0f12e57d422211b051d)
- [app_type_settings.business_logic_markup_setting.enable](resources--app_setting--reference--group-001.md#canonical-77a357dbae85320abf23ad5d68d07f69b64e1d49943c2434a0de05cddddd25e3)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-934b73d092bbd12501a4fd3387ac3088227286827467d0f12e57d422211b051d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb54c8f2b4a7dda192be07c13725369735317ddab76b0fc9f5da9cd0301be64d"></a>

## app_type_settings.business_logic_markup_setting.disable_spec — app_type_settings.business_logic_markup_setting.disable_spec / 4ee5e1a7f8fa / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-36a2be0a9b32b6297811c65c9d994ea33bd57de3ad674979dbf54de31e625e0e)
- app_type_settings.business_logic_markup_setting.disable_spec

<a id="canonical-383223c3a40cb8e8fc783e8a5a1249e99d025ed8373fdcbb34c16325d53d7076"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-b0206b168a76aa4a25836ee873e137238ea7873d8c8794e63175457e151052bb"></a>

## Direct properties — app_type_settings.business_logic_markup_setting.disable_spec / 4ee5e1a7f8fa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f35742cb7963681cea61684b754baed061b654664d49a28d48605dc1772e847"></a>

## Next pages — app_type_settings.business_logic_markup_setting.disable_spec / 4ee5e1a7f8fa / 4

- [app_type_settings.business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-36a2be0a9b32b6297811c65c9d994ea33bd57de3ad674979dbf54de31e625e0e)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-77a357dbae85320abf23ad5d68d07f69b64e1d49943c2434a0de05cddddd25e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b4c9499bdd3a02ebfedd5c18f41a7ff7208a29ab93c4953a131d450ee267b78"></a>

## app_type_settings.business_logic_markup_setting.enable — app_type_settings.business_logic_markup_setting.enable / 5536af574ac8 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-36a2be0a9b32b6297811c65c9d994ea33bd57de3ad674979dbf54de31e625e0e)
- app_type_settings.business_logic_markup_setting.enable

<a id="canonical-2fc85b63e14691a05aefe8bfaf50f5a28e929f76de90ab3a68420c472ab552a1"></a>

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
enable = {}
```

<a id="canonical-574de06ff0b6cb6a34a51a134df4f1d1c9a0a21260d6f248a9e4c5536fbe7b46"></a>

## Direct properties — app_type_settings.business_logic_markup_setting.enable / 5536af574ac8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9692f734a0537e78f58597537be4d651904fb201252f5fd0e3ad5b337400215d"></a>

## Next pages — app_type_settings.business_logic_markup_setting.enable / 5536af574ac8 / 4

- [app_type_settings.business_logic_markup_setting](resources--app_setting--reference--group-001.md#canonical-36a2be0a9b32b6297811c65c9d994ea33bd57de3ad674979dbf54de31e625e0e)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-8ef1350a06bca9d5f9e3014dd9f61b2f79fb8010643ce097d12de338b62a4148"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cee7a05d7f31c0931980e6e3343c996cd2bc149db96d21b6bf7311d73b88df6c"></a>

## app_type_settings.timeseries_analyses_setting — app_type_settings.timeseries_analyses_setting / 131bba055a85 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- app_type_settings.timeseries_analyses_setting

<a id="canonical-a501ea728b97e59308d75d3c20acfd81d4ddf328a4563e8623777b5246cbadc2"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for timeseries analyses setting.

Upstream description:

Configuration for DDoS Detection.

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
timeseries_analyses_setting {
  # Configure direct properties listed below.
}
```

<a id="canonical-333f9c55326ed4dfe88463a168ee196879e3885221c260425e78adc8fbee51ea"></a>

## Direct properties — app_type_settings.timeseries_analyses_setting / 131bba055a85 / 3

- [metric_selectors](resources--app_setting--reference--group-001.md#canonical-699aed2f867843e691b0727f5fe443ee4818f2a6932e169707ee5cbea2c5715b): complete subsection reference.

<a id="canonical-6e1bb9829c50ae7a6d6783a674862284df97a991d594ea4d3b3bd47eec0a342b"></a>

## Next pages — app_type_settings.timeseries_analyses_setting / 131bba055a85 / 4

- [app_type_settings.timeseries_analyses_setting.metric_selectors](resources--app_setting--reference--group-001.md#canonical-699aed2f867843e691b0727f5fe443ee4818f2a6932e169707ee5cbea2c5715b)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-699aed2f867843e691b0727f5fe443ee4818f2a6932e169707ee5cbea2c5715b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29518030cb8098fb4827f9356f0d9ec5ff7c5948e8717f0cec0729730b8e7547"></a>

## app_type_settings.timeseries_analyses_setting.metric_selectors — app_type_settings.timeseries_analyses_setting.metric_selectors / d55879fa30af / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.timeseries_analyses_setting](resources--app_setting--reference--group-001.md#canonical-8ef1350a06bca9d5f9e3014dd9f61b2f79fb8010643ce097d12de338b62a4148)
- app_type_settings.timeseries_analyses_setting.metric_selectors

<a id="canonical-a71bcfcc974b495b3164dd42d99a4053b419ff08d3a71c704913edcf1d425634"></a>

Type: `"object"`. list nested block, Optional.

Define the metric selection criteria, i.e. The metrics source and the actual metrics that should be
included in the detection logic.

Upstream description:

Define the metric selection criteria, i.e. The metrics source and the actual metrics that should be
included in the detection logic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
metric_selectors {
  # Configure direct properties listed below.
}
```

<a id="canonical-b0044f8a582173d3f56336af6617f8e296ebb2a51ecf52516e36215a198e7c5d"></a>

## Direct properties — app_type_settings.timeseries_analyses_setting.metric_selectors / d55879fa30af / 3

<a id="canonical-314c4e29ed7b06a649ec31d7067308a53c06c979e998ae5d49c5f3fba371a5d7"></a>

<a id="canonical-60999a2734b8bf229ea808991c93a8ed29418b217a2d6f8a9505e9b97e8f5d3f"></a>

## metric property — app_type_settings.timeseries_analyses_setting.metric_selectors / d55879fa30af / 4

Type: `["list", "string"]`. Optional.

\[Enum: NO\_METRICS|REQUEST\_RATE|ERROR\_RATE|LATENCY|THROUGHPUT\] Choose one or more metrics to be
included in the detection logic. Possible values are \`NO\_METRICS\`, \`REQUEST\_RATE\`,
\`ERROR\_RATE\`, \`LATENCY\`, \`THROUGHPUT\`. Defaults to \`NO\_METRICS\`.

Upstream description:

Choose one or more metrics to be included in the detection logic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2634bfcd601e57400fac76e183849eb34bd3560eda589702dbdaa2ae8b31d75e"></a>

<a id="canonical-8c8632c5eade4196adff0303dcaa1a3ce800087d2b960e0547e39aa3652ba90b"></a>

## metrics_source property — app_type_settings.timeseries_analyses_setting.metric_selectors / d55879fa30af / 5

Type: `"string"`. Optional.

\[Enum: NONE|NODES|EDGES|VIRTUAL\_HOSTS\] Supported sources from which Metrics can be analyzed All
edges in the service mesh graph. Metrics are analyzed separately between all source and destination
service combinations. Possible values are \`NONE\`, \`NODES\`, \`EDGES\`, \`VIRTUAL\_HOSTS\`.

Upstream description:

Supported sources from which Metrics can be analyzed

All edges in the service mesh graph. Metrics are analyzed separately between all source and
destination service combinations.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NONE",
    "NODES",
    "EDGES",
    "VIRTUAL_HOSTS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NONE",
  "enum": [
    "NONE",
    "NODES",
    "EDGES",
    "VIRTUAL_HOSTS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5ab836cb173ff600c10ea6feb50a4595c6ec50bbba3df404caed6152719b2e5d"></a>

## Next pages — app_type_settings.timeseries_analyses_setting.metric_selectors / d55879fa30af / 6

- [app_type_settings.timeseries_analyses_setting](resources--app_setting--reference--group-001.md#canonical-8ef1350a06bca9d5f9e3014dd9f61b2f79fb8010643ce097d12de338b62a4148)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5d9c237a7001752c4fb43a82a9829d52451f5dc9f40dcae2dcbb099090e2b6e"></a>

## app_type_settings.user_behavior_analysis_setting — app_type_settings.user_behavior_analysis_setting / 33853f248d53 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- app_type_settings.user_behavior_analysis_setting

<a id="canonical-a95b9ddfb094890c9f9d4bd360e2e9d2838b6ca73adddd1a15b0482c169b60fa"></a>

Type: `"object"`. single nested block, Optional.

Configuration for user behavior analysis.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_detection",
    "enable_detection"),
  validators.ConflictingObjectAttributes("disable_learning",
    "enable_learning")}
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
  "x-ves-oneof-field-learn_from_namespace": "[\"disable_learning\",\"enable_learning\"]",
  "x-ves-oneof-field-malicious_user_detection": "[\"disable_detection\",\"enable_detection\"]"
}
```

Terraform syntax:

```terraform
user_behavior_analysis_setting {
  # Configure direct properties listed below.
}
```

<a id="canonical-593913b8116051108beb33668400659be852da4958e45e9fbc0ca5562f74f511"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting / 33853f248d53 / 3

- [disable_detection](resources--app_setting--reference--group-001.md#canonical-57e720e997c118bc8ccee355bad15c81f01a805ccfff47f0baf1ccf03c9117ac): complete subsection reference.

- [disable_learning](resources--app_setting--reference--group-001.md#canonical-9fe185fe9eeccb6bfdea6bb7f878d339da4cc63d44811ee79237bf62165722a3): complete subsection reference.

- [enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7): complete subsection reference.

- [enable_learning](resources--app_setting--reference--group-001.md#canonical-d01137f84d2f346f488d52eef965b5b98f4c0d6fe2e3b9399dbb4ecec344c68a): complete subsection reference.

<a id="canonical-29963a68aee34294e89629504587653c3f689ceb49618c58056cc1a4b4ec0fe3"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting / 33853f248d53 / 4

- [app_type_settings.user_behavior_analysis_setting.disable_detection](resources--app_setting--reference--group-001.md#canonical-57e720e997c118bc8ccee355bad15c81f01a805ccfff47f0baf1ccf03c9117ac)
- [app_type_settings.user_behavior_analysis_setting.disable_learning](resources--app_setting--reference--group-001.md#canonical-9fe185fe9eeccb6bfdea6bb7f878d339da4cc63d44811ee79237bf62165722a3)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [app_type_settings.user_behavior_analysis_setting.enable_learning](resources--app_setting--reference--group-001.md#canonical-d01137f84d2f346f488d52eef965b5b98f4c0d6fe2e3b9399dbb4ecec344c68a)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-57e720e997c118bc8ccee355bad15c81f01a805ccfff47f0baf1ccf03c9117ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe2f28562c77d578cb315297ed62a4ba53580d771ebda7959e274f9c1761480f"></a>

## app_type_settings.user_behavior_analysis_setting.disable_detection — app_type_settings.user_behavior_analysis_setting.disable_detection / fceaf1fb2c2a / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- app_type_settings.user_behavior_analysis_setting.disable_detection

<a id="canonical-f64fc6655935330c8016584e0be182277befa8bc2730f6277b457eadb95bfb9f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable detection.

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
disable_detection = {}
```

<a id="canonical-c7e990ececa040c17adc4520a0ce7ebb967da57ccf49cafbef25bc4d5aaa212b"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.disable_detection / fceaf1fb2c2a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233430b7a6f3b0d560c40bf4e12c30ef3e9b1851cd90f4e1af48bcd24a88d0d"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.disable_detection / fceaf1fb2c2a / 4

- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-9fe185fe9eeccb6bfdea6bb7f878d339da4cc63d44811ee79237bf62165722a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5b9b2c347d2848561176794038d5ef043ee296775ca62df2edc110fbd1c1036"></a>

## app_type_settings.user_behavior_analysis_setting.disable_learning — app_type_settings.user_behavior_analysis_setting.disable_learning / a0f9949fd1b1 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- app_type_settings.user_behavior_analysis_setting.disable_learning

<a id="canonical-5c42e09df28c12b1fe9c355ecf32d42c5c0a9d0f57e30409226c42f99f3b78ca"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable learning.

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
disable_learning = {}
```

<a id="canonical-5140088e803cdaf6e781389829ec7a8d021d8feb839add088bd6784d08550312"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.disable_learning / a0f9949fd1b1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8511b522517e67fb7284deb3c9e8cba75595b10b6527291b5e9eb65b47429b2e"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.disable_learning / a0f9949fd1b1 / 4

- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c47a96c8e795cef23697c838565750edd54178613a0533782b8e399991fd4fa7"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection — app_type_settings.user_behavior_analysis_setting.enable_detection / c0e64fc5831f / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- app_type_settings.user_behavior_analysis_setting.enable_detection

<a id="canonical-400ce4b40c3733e9121789b57deaabcd432d2a1033bf702044aa90d575623721"></a>

Type: `"object"`. single nested block, Optional.

Various factors about user activity are monitored and analysed to determine malicious users. These
settings allow tuning those factors used by the system to detect malicious users.

Upstream description:

Various factors about user activity are monitored and analysed to determine malicious users. These
settings allow tuning those factors used by the system to detect malicious users.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("bola_detection_automatic",
    "exclude_bola_detection"),
  validators.ConflictingObjectAttributes("exclude_bot_defense_activity",
    "include_bot_defense_activity"),
  validators.ConflictingObjectAttributes("exclude_failed_login_activity",
    "include_failed_login_activity"),
  validators.ConflictingObjectAttributes("exclude_forbidden_activity",
    "include_forbidden_activity"),
  validators.ConflictingObjectAttributes("exclude_ip_reputation",
    "include_ip_reputation"),
  validators.ConflictingObjectAttributes("exclude_non_existent_url_activity",
    "include_non_existent_url_activity_automatic"),
  validators.ConflictingObjectAttributes("exclude_non_existent_url_activity",
    "include_non_existent_url_activity_custom"),
  validators.ConflictingObjectAttributes("exclude_rate_limit",
    "include_rate_limit"),
  validators.ConflictingObjectAttributes("exclude_waf_activity",
    "include_waf_activity"),
  validators.ConflictingObjectAttributes("include_non_existent_url_activity_automatic",
    "include_non_existent_url_activity_custom")}
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
  "x-ves-oneof-field-bola_activity_choice": "[\"bola_detection_automatic\",\"exclude_bola_detection\"]",
  "x-ves-oneof-field-bot_defense_activity_choice": "[\"exclude_bot_defense_activity\",\"include_bot_defense_activity\"]",
  "x-ves-oneof-field-cooling_off_period_setting": "[\"cooling_off_period\"]",
  "x-ves-oneof-field-failed_login_activity_choice": "[\"exclude_failed_login_activity\",\"include_failed_login_activity\"]",
  "x-ves-oneof-field-forbidden_activity_choice": "[\"exclude_forbidden_activity\",\"include_forbidden_activity\"]",
  "x-ves-oneof-field-ip_reputation_choice": "[\"exclude_ip_reputation\",\"include_ip_reputation\"]",
  "x-ves-oneof-field-non_existent_url_activity_choice": "[\"exclude_non_existent_url_activity\",\"include_non_existent_url_activity_automatic\",\"include_non_existent_url_activity_custom\"]",
  "x-ves-oneof-field-rate_limit_choice": "[\"exclude_rate_limit\",\"include_rate_limit\"]",
  "x-ves-oneof-field-waf_activity_choice": "[\"exclude_waf_activity\",\"include_waf_activity\"]"
}
```

Terraform syntax:

```terraform
enable_detection {
  # Configure direct properties listed below.
}
```

<a id="canonical-8d79c11f004e51f190f2b48febf81b5d2c5ad76a26e3aaf1f24231fe6110e6bf"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection / c0e64fc5831f / 3

- [bola_detection_automatic](resources--app_setting--reference--group-001.md#canonical-8c0b897ccd2e760e64b7a01504c6496415b67a12615fab80524d11ae2d9be0c5): complete subsection reference.

<a id="canonical-61b8767049c9f95b870fdc69152b5a96af558405f63e7a417dbc85d7aebcacee"></a>

<a id="canonical-43efe8ce3e1d82da2e428fed011b0e8df7291b5824d3c8f17db7e7631bc3184e"></a>

## cooling_off_period property — app_type_settings.user_behavior_analysis_setting.enable_detection / c0e64fc5831f / 4

Type: `"number"`. Optional.

Exclusive with \[\] Malicious user detection assigns a threat level to each user based on their
activity. Once a threat level is assigned, the system continues tracking activity from this user and
if no further malicious activity is seen, it gradually reduces the threat assessment to lower
levels..

Upstream description:

Exclusive with \[\] Malicious user detection assigns a threat level to each user based on their
activity. Once a threat level is assigned, the system continues tracking activity from this user and
if no further malicious activity is seen, it gradually reduces the threat assessment to lower
levels. This field specifies the time period, in minutes, used by the system to decay a user's
threat level from a high to medium or medium to low or low to none.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(5, 120),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 120,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 5
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "120"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "120"
  }
}
```

- [exclude_bola_detection](resources--app_setting--reference--group-001.md#canonical-4a4735720ddae208b084184b0313c7d6852f357b57d4d1ad0e54c45f625f4923): complete subsection reference.

- [exclude_bot_defense_activity](resources--app_setting--reference--group-001.md#canonical-e6a9d17786a70fb93abdf0228b3abb8a59000236846c279bd33560e08bac8b82): complete subsection reference.

- [exclude_failed_login_activity](resources--app_setting--reference--group-001.md#canonical-db405d5dc6cf8dd528fa4d125896761d97c73bc79e9eb40bec9d6a4a3025e38d): complete subsection reference.

- [exclude_forbidden_activity](resources--app_setting--reference--group-001.md#canonical-4cbec45cd1172b884b82a7e753ce85a7b3a60a2b089a232334ea2e803661b6b5): complete subsection reference.

- [exclude_ip_reputation](resources--app_setting--reference--group-001.md#canonical-e89f5feccbb15c790cb9755dc87939b52dba5a0720e5438858c58dee309798b3): complete subsection reference.

- [exclude_non_existent_url_activity](resources--app_setting--reference--group-001.md#canonical-fa2d5cb2615e22fd4daec7de7170a1a1383749b2e795ae8d56ad87366893e684): complete subsection reference.

- [exclude_rate_limit](resources--app_setting--reference--group-001.md#canonical-8064da0e2e844df91ed036a65ac0b915e9921aebc73305c1ea941f9d83d5adce): complete subsection reference.

- [exclude_waf_activity](resources--app_setting--reference--group-001.md#canonical-40c84c61625bb0d41b3e2f925038315d72ff6d0b9649fbc8a5132dac40acfd9c): complete subsection reference.

- [include_bot_defense_activity](resources--app_setting--reference--group-001.md#canonical-bdfd3c4e9528639085c1c02cb3271c6860cf8e8c9422bec3e38ff00938528dfb): complete subsection reference.

- [include_failed_login_activity](resources--app_setting--reference--group-001.md#canonical-65421d22cb003bf12dde77aebfd4df147a22f200376078b4b7cae1cf081aefdb): complete subsection reference.

- [include_forbidden_activity](resources--app_setting--reference--group-001.md#canonical-59b85f16ce1b4a9cfed0acfb83c8887322a4941b2058e2f7e357243205fa31e4): complete subsection reference.

- [include_ip_reputation](resources--app_setting--reference--group-001.md#canonical-209971f904c9e70e82b2f1100768f68a51b1820eeecd7f73fa9b36904ae796fb): complete subsection reference.

- [include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-be5abca0cd2f042af95a31ac194cac0c64757ac4bf2c18ba3c067f93ae3eb6d1): complete subsection reference.

- [include_non_existent_url_activity_custom](resources--app_setting--reference--group-001.md#canonical-cfa14ec7bb1c3cf14fb5bd291e0fe016ebc69c99f949302fdf7743911bff9b9b): complete subsection reference.

- [include_rate_limit](resources--app_setting--reference--group-001.md#canonical-e20ee432d778f12f7942ec87a89dfa69b49fefcc9fe5ee90fd4cada018a811c0): complete subsection reference.

- [include_waf_activity](resources--app_setting--reference--group-001.md#canonical-b9d46803fbfb042e14ae681e28b89a8bdf98f5abd2b339510d88c9ff67827cd5): complete subsection reference.

<a id="canonical-6e52ab2ee43435dbcb305ff593a3432e9dd0086fcde66ba1c80d9bb6a1c61382"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection / c0e64fc5831f / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic](resources--app_setting--reference--group-001.md#canonical-8c0b897ccd2e760e64b7a01504c6496415b67a12615fab80524d11ae2d9be0c5)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection](resources--app_setting--reference--group-001.md#canonical-4a4735720ddae208b084184b0313c7d6852f357b57d4d1ad0e54c45f625f4923)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity](resources--app_setting--reference--group-001.md#canonical-e6a9d17786a70fb93abdf0228b3abb8a59000236846c279bd33560e08bac8b82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity](resources--app_setting--reference--group-001.md#canonical-db405d5dc6cf8dd528fa4d125896761d97c73bc79e9eb40bec9d6a4a3025e38d)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity](resources--app_setting--reference--group-001.md#canonical-4cbec45cd1172b884b82a7e753ce85a7b3a60a2b089a232334ea2e803661b6b5)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation](resources--app_setting--reference--group-001.md#canonical-e89f5feccbb15c790cb9755dc87939b52dba5a0720e5438858c58dee309798b3)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity](resources--app_setting--reference--group-001.md#canonical-fa2d5cb2615e22fd4daec7de7170a1a1383749b2e795ae8d56ad87366893e684)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit](resources--app_setting--reference--group-001.md#canonical-8064da0e2e844df91ed036a65ac0b915e9921aebc73305c1ea941f9d83d5adce)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity](resources--app_setting--reference--group-001.md#canonical-40c84c61625bb0d41b3e2f925038315d72ff6d0b9649fbc8a5132dac40acfd9c)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity](resources--app_setting--reference--group-001.md#canonical-bdfd3c4e9528639085c1c02cb3271c6860cf8e8c9422bec3e38ff00938528dfb)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity](resources--app_setting--reference--group-001.md#canonical-65421d22cb003bf12dde77aebfd4df147a22f200376078b4b7cae1cf081aefdb)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity](resources--app_setting--reference--group-001.md#canonical-59b85f16ce1b4a9cfed0acfb83c8887322a4941b2058e2f7e357243205fa31e4)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation](resources--app_setting--reference--group-001.md#canonical-209971f904c9e70e82b2f1100768f68a51b1820eeecd7f73fa9b36904ae796fb)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-be5abca0cd2f042af95a31ac194cac0c64757ac4bf2c18ba3c067f93ae3eb6d1)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom](resources--app_setting--reference--group-001.md#canonical-cfa14ec7bb1c3cf14fb5bd291e0fe016ebc69c99f949302fdf7743911bff9b9b)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit](resources--app_setting--reference--group-001.md#canonical-e20ee432d778f12f7942ec87a89dfa69b49fefcc9fe5ee90fd4cada018a811c0)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity](resources--app_setting--reference--group-001.md#canonical-b9d46803fbfb042e14ae681e28b89a8bdf98f5abd2b339510d88c9ff67827cd5)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-8c0b897ccd2e760e64b7a01504c6496415b67a12615fab80524d11ae2d9be0c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e7a28295413057a7ee7429786369670f756fb7f686ac53cbaab2096b87aeee7"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic — app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection / e311e5c5b2b9 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic

<a id="canonical-4a39622499de50363929a7476c736c7612f5baf2bba156b63c7309db131c9406"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for bola detection automatic.

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
bola_detection_automatic = {}
```

<a id="canonical-3d90c22774a8afc18d01440a4fb39229f08f5a3fa90b4de54987a53943977acd"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection / e311e5c5b2b9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1921da5295f1ce7bbdb6b6399a2a3719fde986c309b7adcb28492002ed1d6636"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection / e311e5c5b2b9 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-4a4735720ddae208b084184b0313c7d6852f357b57d4d1ad0e54c45f625f4923"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e6d7895b9a1d53c1ef68236921d909026e9677e5b1d54eed6937a2503bc5346"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_d / 179bd1cd8518 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection

<a id="canonical-d3e66e10d555c43738de98344ca93b58fc55f982786bcaca633ade01b2bd8c22"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exclude bola detection.

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
exclude_bola_detection = {}
```

<a id="canonical-ea3e4cf5ba30d10800be96483e322740a00c3edb8cd33fabc4d2a0505607a19b"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_d / 179bd1cd8518 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4abbe1ffb2b8978bb0c52fa55b4bd0902196d2b8df4160f80a8e0f2ea7c1ae92"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_d / 179bd1cd8518 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-e6a9d17786a70fb93abdf0228b3abb8a59000236846c279bd33560e08bac8b82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e27ed17e8daff2f8814c9907181a34f23bad813decd245fc9ba805dd2a24f005"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_de / 00314365fd55 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity

<a id="canonical-ff3dc5f8c0ac27dd7a777dcd921c91dd448be7e9a05aa8c109d129ded2bbfd81"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exclude bot defense activity.

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
exclude_bot_defense_activity = {}
```

<a id="canonical-4f80570b348a27dab7c444ce12cd3ce74d5c227bbe8cb782279ebf9fb6f43dc2"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_de / 00314365fd55 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dd7c69f4b5160739cc626e58bdd7cfbcc74f49b6163d95ab2579c19c74479374"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_de / 00314365fd55 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-db405d5dc6cf8dd528fa4d125896761d97c73bc79e9eb40bec9d6a4a3025e38d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a865f88c24bb68642aa889724822401c5f4e3aff199ce26ac800af53a58b87a"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed / cc74ea48d6c0 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity

<a id="canonical-bae5b5aef22977b9c512461aa6e3ee05c6af5a856153eff6dc008bb9153ca3f5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exclude failed login activity.

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
exclude_failed_login_activity = {}
```

<a id="canonical-dbbbe0bbafdca5a863dc2c13ce7cdbf407974ae71fbd986c070cde400666d526"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed / cc74ea48d6c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-201685fc76198f614ae90a5b05a445dbdb2951bb1f59fec86dd5bb5d4789d0d8"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed / cc74ea48d6c0 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-4cbec45cd1172b884b82a7e753ce85a7b3a60a2b089a232334ea2e803661b6b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2f8856f82e7b0d7381452d45fb385f0359e2578548fcca9ea84a806f6ff88be"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbid / e589936c030b / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity

<a id="canonical-055aaee013ee47c8da952ca757cb6f16ca6b62cb56a2703809d3b1924c9227a5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exclude forbidden activity.

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
exclude_forbidden_activity = {}
```

<a id="canonical-aa6ef01db1d70b161ea5a48c938f4a45fd7858582e631454c6b0c4c9c8d091a4"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbid / e589936c030b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f7f218b5d171574686f244641546dfce6d568b2c8b57ef9ff6e517003d3fe7e6"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbid / e589936c030b / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-e89f5feccbb15c790cb9755dc87939b52dba5a0720e5438858c58dee309798b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a57baa8ed8d364fbb5113dd4671b41722729f2a21d664e83a3539310442cb791"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_rep / c2cb8c200dda / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation

<a id="canonical-64a65a9a7bbfa2d9e9b952174f3345ace232b9482ecb0b4de421e1af310cebc6"></a>

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
exclude_ip_reputation = {}
```

<a id="canonical-1baa3ab2073f041c2752307a151c28543bf0c3ec1b5d970163b821c27ea107c7"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_rep / c2cb8c200dda / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-82b8a396eb720e86d68c2aa8ff7184d435584863937ed608fc7f5dc12d3a0289"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_rep / c2cb8c200dda / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-fa2d5cb2615e22fd4daec7de7170a1a1383749b2e795ae8d56ad87366893e684"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c15d47a6b9facfad5beb1e09ffbe3a7a43dbf01490902b98fbaf5ec652d1d34"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_ex / feb3e28efcc6 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity

<a id="canonical-892729e56e3da5eae13d868cd170e2de916118166a357bc9906c670011e9cddc"></a>

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
exclude_non_existent_url_activity = {}
```

<a id="canonical-7ab7187264a7094a1899cb82321cd8e12dc646849176ff57fe2502ac43b80ab2"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_ex / feb3e28efcc6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1e72078caa1cc2e26a55a8cfce486ea29333217774c2ed21e0375066aac21d25"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_ex / feb3e28efcc6 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-8064da0e2e844df91ed036a65ac0b915e9921aebc73305c1ea941f9d83d5adce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bd9a165575ab3de4c1b914827378b956266fb649e205664396d680e0314be04"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_l / 7a9690997adc / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit

<a id="canonical-7c89ddf27eda6b15df7c5864bc8a086bfd5293f780dbc60141aa312c68b28032"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exclude rate limit.

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
exclude_rate_limit = {}
```

<a id="canonical-8fa0c606095832625c4cc863c95c817178d228afd5c3512959b50e83db2d1c58"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_l / 7a9690997adc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c76b59587aef8f912088651e7ad8cb518bd145f9387fb733a26779bee685d84f"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_l / 7a9690997adc / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-40c84c61625bb0d41b3e2f925038315d72ff6d0b9649fbc8a5132dac40acfd9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00468deb2336b2fc04517b460cfe9d1f0ba249255cd45393bdf88a3cf4979555"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_ac / 1daaab3beec2 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity

<a id="canonical-be51d8aa950b355900f47ebabddaf65e23a11f1c1d754068bc70ac3870194606"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exclude waf activity.

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
exclude_waf_activity = {}
```

<a id="canonical-7b73276807066a5c832ad80efcf54902dde996302ec533b7f477ec7bb91094ef"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_ac / 1daaab3beec2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-453d5365ca5f221ac784a0c0b49f84758a1bf259ffadf1e9679dcd860ea8d682"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_ac / 1daaab3beec2 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-bdfd3c4e9528639085c1c02cb3271c6860cf8e8c9422bec3e38ff00938528dfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c4c724663b9bc0cf4b6ac57b100000088a85ecb673efcf61250a0035dafc3f5"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_de / 31862386371e / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity

<a id="canonical-de451c060ed5349afd195a16ec8e130454450b41f6c19cd7fe823a4ee5f88be6"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for include bot defense activity.

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
include_bot_defense_activity = {}
```

<a id="canonical-e175c46789559eef75ca4776e8fa10f56e3fabdcb33b18b52bfc7a1783264e09"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_de / 31862386371e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-521462b981c2bf83bab67a91d45327ba3ab1ccc83517630d162115213eee650d"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_de / 31862386371e / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-65421d22cb003bf12dde77aebfd4df147a22f200376078b4b7cae1cf081aefdb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e123ca770341d5279e1fb2ef2a1960f9ddfae8bd63ca7dd8f6ffa01f3ec6163"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed / c05a9c094b7e / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity

<a id="canonical-1ab4f3403976fd4bf8429eb30cf6ac5b76c2f04c23c364e6cd2934bb0d04459a"></a>

Type: `"object"`. single nested block, Optional.

When enabled, the system monitors persistent failed login attempts from a user. A failed login is
detected if a request results in a response code of 401. These settings specify how to use failed
login activity to determine suspicious behavior.

Upstream description:

When enabled, the system monitors persistent failed login attempts from a user. A failed login is
detected if a request results in a response code of 401. These settings specify how to use failed
login activity to determine suspicious behavior.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("login_failures_threshold")}
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
include_failed_login_activity {
  # Configure direct properties listed below.
}
```

<a id="canonical-cddb056147e444c0d532b7a4686c73e1a8848c23af0a9fdb03ea311e31f58c7f"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed / c05a9c094b7e / 3

<a id="canonical-de84a99b005cd2b5a07c425439896a3d5d5ca0af3ce43554b63304324c5922b0"></a>

<a id="canonical-bf4f1c0e1e43b7c4dab78fdd477e56503eb1aacfd877a487bc4ffbfb9cc7b5ba"></a>

## login_failures_threshold property — app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed / c05a9c094b7e / 4

Type: `"number"`. Optional.

The number of failed logins beyond which the system will flag this user as malicious.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

<a id="canonical-ad8fa5f024523a794768e300d3cb2d1b48ca6ab3ee768c8ab41e0ab644b6cd48"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed / c05a9c094b7e / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-59b85f16ce1b4a9cfed0acfb83c8887322a4941b2058e2f7e357243205fa31e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8f5582e926ee220eb8f6e3b42832c71e9bbb048f1939ee69023e7a46b956b29"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbid / 88efed2c2452 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity

<a id="canonical-c99a8aec36705697e74045bc659bbbe7eca2af4fbf3c9ff1b5a7b6346a36ae8c"></a>

Type: `"object"`. single nested block, Optional.

When L7 policy rules are set up to disallow certain types of requests, the system monitors
persistent attempts from a user to send requests which result in policy denies. These settings
specify how to use disallowed request activity from a user to determine suspicious behavior.

Upstream description:

When L7 policy rules are set up to disallow certain types of requests, the system monitors
persistent attempts from a user to send requests which result in policy denies. These settings
specify how to use disallowed request activity from a user to determine suspicious behavior.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forbidden_requests_threshold")}
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
include_forbidden_activity {
  # Configure direct properties listed below.
}
```

<a id="canonical-352518f112e47fd66e83093228079059722933dec9537f6048c0b6ce6df51ec9"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbid / 88efed2c2452 / 3

<a id="canonical-139d367f036c04fa7acc9fee109bb8286137838447bbcaec2d540ffe9011e122"></a>

<a id="canonical-4b6362346296394f95fe29dd47ac3d11c4e10ba923fc4403af202b1996d094ef"></a>

## forbidden_requests_threshold property — app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbid / 88efed2c2452 / 4

Type: `"number"`. Optional.

The number of forbidden requests beyond which the system will flag this user as malicious.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

<a id="canonical-fef093afe1a6d6d445c54b43e35c8e2a28f09d1dbfb8b20ac8f70b94f54daf1f"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbid / 88efed2c2452 / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-209971f904c9e70e82b2f1100768f68a51b1820eeecd7f73fa9b36904ae796fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb075bfc10d1644a8949726e0a0ee8ef66f7d22561d0ba928e79d9c4c9bca9ad"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation — app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_rep / 51e7c20e9ed2 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation

<a id="canonical-8ed0d496152743f966d0a70d526527f0aae52c14e690190561c8d72341011def"></a>

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
include_ip_reputation = {}
```

<a id="canonical-d16c412bd6a818813968a7389e881843ce6a66fa366beefa5ff6a34f21fc5456"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_rep / 51e7c20e9ed2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2c7c82b9ba90766145c7a99031d2470cee476ef4ee60ee0f4e150772e65cadb4"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_rep / 51e7c20e9ed2 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-be5abca0cd2f042af95a31ac194cac0c64757ac4bf2c18ba3c067f93ae3eb6d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec81f8727d6dfac45efca93be9b7aaa702bff68dd5fc04faf153c47e7b2670f7"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / da3e6ef40522 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic

<a id="canonical-dca5763c54c675063f514c307eb53856f11473731ee60222923ff5efc618a09f"></a>

Type: `"object"`. single nested block, Optional.

Non-existent URL Automatic Activity Settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("high",
    "low"),
  validators.ConflictingObjectAttributes("high",
    "medium"),
  validators.ConflictingObjectAttributes("low",
    "medium")}
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
  "x-ves-oneof-field-sensitivity": "[\"high\",\"low\",\"medium\"]"
}
```

Terraform syntax:

```terraform
include_non_existent_url_activity_automatic {
  # Configure direct properties listed below.
}
```

<a id="canonical-968e3d96d1e3b4824122f3019d2dc8cd89598676acea1952c18fefa48b9226b6"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / da3e6ef40522 / 3

- [high](resources--app_setting--reference--group-001.md#canonical-f3b048be366e08f39b15bed3ee0bb4ac1e9b2d361e5fec5517e7b719bc55a38b): complete subsection reference.

- [low](resources--app_setting--reference--group-001.md#canonical-93436bbc6d1f5784562a5abbc54ec471891d6913c197a9d4523b54ee39f47e7f): complete subsection reference.

- [medium](resources--app_setting--reference--group-001.md#canonical-544224adaf6f9fe0556d1379bd65648c422718d11ad39868c4029c583b31e320): complete subsection reference.

<a id="canonical-c7d721acd994da9d5bfa97e13c3e7bd03a830f361895666920169de193e0b97c"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / da3e6ef40522 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high](resources--app_setting--reference--group-001.md#canonical-f3b048be366e08f39b15bed3ee0bb4ac1e9b2d361e5fec5517e7b719bc55a38b)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low](resources--app_setting--reference--group-001.md#canonical-93436bbc6d1f5784562a5abbc54ec471891d6913c197a9d4523b54ee39f47e7f)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium](resources--app_setting--reference--group-001.md#canonical-544224adaf6f9fe0556d1379bd65648c422718d11ad39868c4029c583b31e320)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-f3b048be366e08f39b15bed3ee0bb4ac1e9b2d361e5fec5517e7b719bc55a38b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-699317ff63047e25ca74a4948a9fc77371f9f4dc04c6cfbb9eef4c6006317f2c"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 1882e8087d3c / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-be5abca0cd2f042af95a31ac194cac0c64757ac4bf2c18ba3c067f93ae3eb6d1)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high

<a id="canonical-bcfb9800a9427a77deb06ed1ee85860c08a4a48dcbde20b3989de6d43d39f186"></a>

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
high = {}
```

<a id="canonical-edaf7b63ffeef504667f2e9154019ade8db5d01090a6b0e9aab3fe4a43eb154c"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 1882e8087d3c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-90da1856c02008b6ecb9bdd443177f08ca281764a1167d47799c49f1429d6c9c"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 1882e8087d3c / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-be5abca0cd2f042af95a31ac194cac0c64757ac4bf2c18ba3c067f93ae3eb6d1)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-93436bbc6d1f5784562a5abbc54ec471891d6913c197a9d4523b54ee39f47e7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-133fdb301efc03026cb60ba6dfe8d7a4920084412a217e97ae30eab5da28e305"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / d780d1571afa / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-be5abca0cd2f042af95a31ac194cac0c64757ac4bf2c18ba3c067f93ae3eb6d1)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low

<a id="canonical-0bafa5b6d66dc599eef7ac1deaa2ae93cbece751a4003026d8341f5dba55ac98"></a>

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
low = {}
```

<a id="canonical-ae0ff2afaee4f87ff712a78262ea9e6bf5340340e9945e9c414fcd4345ccec59"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / d780d1571afa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9499941e0958a655f2d2074d48bf93c118b63a5fefe311c2faf7ccddfed8fea"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / d780d1571afa / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-be5abca0cd2f042af95a31ac194cac0c64757ac4bf2c18ba3c067f93ae3eb6d1)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-544224adaf6f9fe0556d1379bd65648c422718d11ad39868c4029c583b31e320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9cd3f8ceb518484a9f454fb1b63083a0901144c2ae419580d5f356215f84a70d"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 577416ccedb4 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-be5abca0cd2f042af95a31ac194cac0c64757ac4bf2c18ba3c067f93ae3eb6d1)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium

<a id="canonical-822b460b996ee8dfa374049717aab6b79f7ea0942c3c94c44377a32bee843771"></a>

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
medium = {}
```

<a id="canonical-3d0d5f74e7536691e94716aeb4404447bee6f4cd74f0c7e1216c687491ca784d"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 577416ccedb4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cf319968d66c6efa6fed4a4f674f80ab7fe86acc6db83d728a2daff2940eca65"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 577416ccedb4 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--reference--group-001.md#canonical-be5abca0cd2f042af95a31ac194cac0c64757ac4bf2c18ba3c067f93ae3eb6d1)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-cfa14ec7bb1c3cf14fb5bd291e0fe016ebc69c99f949302fdf7743911bff9b9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e01d456db9993f842bba5a4d4359cafddd7083e921ef7412b1378cc20bd22eb"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 23bc07486d17 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom

<a id="canonical-83f0ddf8f8bdab5035a9e32c2ea49aa39db231db13099c82938afe95645a2791"></a>

Type: `"object"`. single nested block, Optional.

Non-existent URL Custom Activity Setting.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("nonexistent_requests_threshold")}
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
include_non_existent_url_activity_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-3ff7384f3ecfa16125376b56e932d651c342cf6434e1d9b9b0c8bbc554b56e24"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 23bc07486d17 / 3

<a id="canonical-836064832f5f292f9c58d2f3d94c65236e426a5b36144dd6065fe9bb001047fe"></a>

<a id="canonical-fd90b643f6ddde45985934194c4b906af3a5e2cb2ffe37cd6576925d51315186"></a>

## nonexistent_requests_threshold property — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 23bc07486d17 / 4

Type: `"number"`. Optional.

The percentage of non-existent requests beyond which the system will flag this user as malicious.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-2eab68cca3d9fef4d1898153e462b2a2ed86d34bcc5deebcb865bd567c28207e"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 23bc07486d17 / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-e20ee432d778f12f7942ec87a89dfa69b49fefcc9fe5ee90fd4cada018a811c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c4c326d8dbbc0bbe47536d02f535199e0dad11260ed0f6ed0b3cbd6910ce115"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit — app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_l / 24697374cd4c / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit

<a id="canonical-95325c36050af5c6f4912bbb941ac7b6b9c84f49b4859850c694416af3ebcfb7"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for include rate limit.

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
include_rate_limit = {}
```

<a id="canonical-e96c42db04dacfd4031ae0fa2a959c9a5870e5de8f8bc4788d54f18ffbf2c723"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_l / 24697374cd4c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5faa8bc37165cd808405ec68d206b55cfb706c31499ac52623d97af969cf393c"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_l / 24697374cd4c / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-b9d46803fbfb042e14ae681e28b89a8bdf98f5abd2b339510d88c9ff67827cd5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-304794428d09fe6987a05ed1aeb1ab2930cc723833b0f7b3370d45b1408f984e"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_ac / 36852fa8488b / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity

<a id="canonical-c8ffda23bce7d985e23cc069c96a6397521ea01e4ec5fa045aa15720243eebf4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for include waf activity.

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
include_waf_activity = {}
```

<a id="canonical-cde4104dfd9600d317a355406c3582ddd4f731514dfdad0c3003e0663bae05e0"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_ac / 36852fa8488b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a819fd71901e8d6bd2718aec06ebeac1ad31bf85a261e91d1d0e0d37eccb5cb"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_ac / 36852fa8488b / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--reference--group-001.md#canonical-71bc16553c0d1c17e9dd0f711fc0ebbcb5733b172907bdd8bfd9d07152743cc7)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-d01137f84d2f346f488d52eef965b5b98f4c0d6fe2e3b9399dbb4ecec344c68a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4f748545c1e4a80188703ffaeea5a684a161812811c07cfd57cf244a32a45e2"></a>

## app_type_settings.user_behavior_analysis_setting.enable_learning — app_type_settings.user_behavior_analysis_setting.enable_learning / a72945b3c46c / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [app_type_settings](resources--app_setting--reference--group-001.md#canonical-6e6e9d5933f46c38e5071f676e56063b373a360c2c053c3b274df33ad8c03356)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- app_type_settings.user_behavior_analysis_setting.enable_learning

<a id="canonical-d07699123055523581526a89343a1cc3372bde8066f4e63124d0e7701bcc5c13"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable learning.

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
enable_learning = {}
```

<a id="canonical-e3811167ac6a81aea6508e43662df94ef4c7b4acdfbf9fa298b098776e6e9518"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_learning / a72945b3c46c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a1725603d7c8361cbaadbffb6394aa0767d050eff69c2fa25ce3374d9adaf0b"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_learning / a72945b3c46c / 4

- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--reference--group-001.md#canonical-a0b0eb4e61e218c4cbf49160c75fb0b92da42f6d217517c5a81ef36c1b3c4906)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-b6bf006bee2af16444ad1d452e09cf84122225ae3f0f01656f31a62f5a1b158e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58bdb562521f31ed6352fdf57acbee760d6c9d426f7def3ca7b5d62ee9053707"></a>

## timeouts — timeouts / f4c0ddc450e1 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- timeouts

<a id="canonical-c8788ad92ab8c25486b501baacbdcd9e45d2481c2c8fc852221bed8e3b0136fc"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-bdea50f4526f46d32cd727c69c7279c7583e2708f201d2e3b59c8fc630bd12ff"></a>

## Direct properties — timeouts / f4c0ddc450e1 / 3

<a id="canonical-717ca7cdd15db15a6712df2c251a0c1a24bc87603d29d5e55b63bfe9f2f53300"></a>

<a id="canonical-923b4717bdb17c9caa25de4910552923b83314ebf4f073e15321376b5b07e8ab"></a>

## create property — timeouts / f4c0ddc450e1 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-ac5a7cf64f4662deef75a4e292cf0acbe2c60880484191a4c3777070f830f3d7"></a>

<a id="canonical-155f66f1139a7802c99aaff1ef0bbd6240de78b9e711f63553b9f601aef1bd19"></a>

## delete property — timeouts / f4c0ddc450e1 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-7faee6be1bb2b66561873e366d3961dd72a70b05a4bff752225813316301c9f8"></a>

<a id="canonical-21d8b70ad96502f30478b2a00ca4b89ac06ae6a28aaedad28a68630499f79463"></a>

## read property — timeouts / f4c0ddc450e1 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-333fc64ee989c59e99aec03be8a2b65d9f516fe2dc691a39f047558218c29e22"></a>

<a id="canonical-5bcc8f00e23d1226e82ed75589ce076a42f737f12a5d9fc4556965bd45a555f3"></a>

## update property — timeouts / f4c0ddc450e1 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-c1d83313bad232067cedf99ff19520537c8ee79606340bee00923abdb18109c5"></a>

## Next pages — timeouts / f4c0ddc450e1 / 8

- [Property reference](resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
