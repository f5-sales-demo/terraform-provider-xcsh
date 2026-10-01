---
page_title: "xcsh_app_setting reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_setting reference."
---

# xcsh_app_setting reference

<a id="canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff4a567b9ce5188f9ed66f930e52469a2e1c87eb0b22c22a95796283aec7cc8c"></a>

## Property reference — Property reference / b536487e6ba7 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- Property reference

<a id="canonical-b7e5b371cf655263cdad4ed06fba20ddb959b8bcf0a2b0af73a4f1e0ff637993"></a>

## Direct properties — Property reference / b536487e6ba7 / 3

<a id="canonical-05c67d363710c1913d4ddf799754671507d073ae20f7832e3045b39dc0159b53"></a>

<a id="canonical-496af482fd0138ef4db06d9b6970a60ea2d8580ea0548392de7e1ee644dab9be"></a>

## annotations property — Property reference / b536487e6ba7 / 4

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

- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd): complete subsection reference.

<a id="canonical-15d34d40e9d8140171a454ce7fcce73c712694bf023947e2ac5d9dfaaea822d7"></a>

<a id="canonical-3dd55b87882b1a0dbb228dade18882fe6f18ccd333f80d7d8e749de94ee5223a"></a>

## description property — Property reference / b536487e6ba7 / 5

Type: `"string"`. Computed.

Description of the AppSetting.

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

<a id="canonical-00e6c8929f1bba22273b75284dde3e653682feaf2c778ecbcfd0e91da4ebfe8c"></a>

<a id="canonical-ff69b93dc68885a460c190bd01dcccac8e26bc2507d59dfaf90dc169cdc603b9"></a>

## id property — Property reference / b536487e6ba7 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-c9b98a97e2016268a5044f33cd12de26ad0db623c2675bcfb32e2d3a3bcbaf46"></a>

<a id="canonical-e4c8b4d6581f4ce76e08c0d60dfa160fac746881622fc840698aa6c2b6a31e23"></a>

## labels property — Property reference / b536487e6ba7 / 7

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

<a id="canonical-85e625374728786c79dcb40e91d4fe734f3b1d858a18e404eb45c7b7dd843955"></a>

<a id="canonical-0ceb7502b275796edf013659cfc07f903b59742fc2cb1bd60d94827449400aab"></a>

## name property — Property reference / b536487e6ba7 / 8

Type: `"string"`. Required.

Name of the AppSetting.

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

<a id="canonical-a31a3eda72572eaef481e42d773aa64a88ad4989b0505fa9530ba582023c3b96"></a>

<a id="canonical-07dc83473b5b7bc028b2cbbfaf9036b3aff71962fef3c592a8e2ef64bf812240"></a>

## namespace property — Property reference / b536487e6ba7 / 9

Type: `"string"`. Required.

Namespace where the AppSetting exists.

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

<a id="canonical-bf8cbdfc75b857f4829bddb629045f2172fd7b96c8f3c7eacbeaac97cc824e2c"></a>

## All schema paths — Property reference / b536487e6ba7 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--app_setting--reference--group-001.md#canonical-05c67d363710c1913d4ddf799754671507d073ae20f7832e3045b39dc0159b53) |
| `app_type_settings` | [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-210def4d3fabd9fe042194cfc1601c398e417c724682cbeb123108571604cb77) |
| `app_type_settings.app_type_ref` | [app_type_settings.app_type_ref](data-sources--app_setting--reference--group-001.md#canonical-1bdde31e739f7686922a37e3b789ffe152bfc97fa36d246b06d2bd65e79a56c1) |
| `app_type_settings.app_type_ref.kind` | [app_type_settings.app_type_ref.kind](data-sources--app_setting--reference--group-001.md#canonical-e066be21e255079614835155b795e5c8142de419a845e6c63b341ccc5f8d7179) |
| `app_type_settings.app_type_ref.name` | [app_type_settings.app_type_ref.name](data-sources--app_setting--reference--group-001.md#canonical-d294c0632e30e182b4c92840af4d4569fa021b5c1cdc68ba7842397cb978585a) |
| `app_type_settings.app_type_ref.namespace` | [app_type_settings.app_type_ref.namespace](data-sources--app_setting--reference--group-001.md#canonical-70741bdf73b1f2d75cfa8e036068d92773947b0369b8e31c803f996990f44606) |
| `app_type_settings.app_type_ref.tenant` | [app_type_settings.app_type_ref.tenant](data-sources--app_setting--reference--group-001.md#canonical-b7ad8663bbb218d81867d211581e32c0291e71881ae80b421f0448456867366a) |
| `app_type_settings.app_type_ref.uid` | [app_type_settings.app_type_ref.uid](data-sources--app_setting--reference--group-001.md#canonical-d43e59631db6e5b2f43218e9dd48698017e367a9ee91daf87f117af4a3360fa9) |
| `app_type_settings.business_logic_markup_setting` | [app_type_settings.business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-6f6371980484120705576d38867d983396fe8e3f51d363cb359274125bb98797) |
| `app_type_settings.business_logic_markup_setting.disable_spec` | [app_type_settings.business_logic_markup_setting.disable_spec](data-sources--app_setting--reference--group-001.md#canonical-73524c79662346538ea52b585fb169debe1fc201467ce6a76ccddbcfc3b7d333) |
| `app_type_settings.business_logic_markup_setting.enable` | [app_type_settings.business_logic_markup_setting.enable](data-sources--app_setting--reference--group-001.md#canonical-ac49644d05edd72d9e223f9a6f6427218aff5352f793926e29623c1f125a342e) |
| `app_type_settings.timeseries_analyses_setting` | [app_type_settings.timeseries_analyses_setting](data-sources--app_setting--reference--group-001.md#canonical-9c10645de88906ef8cd8dbbcb3cac78a7aab6c6e2136329597b7c3e4324f9a26) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors` | [app_type_settings.timeseries_analyses_setting.metric_selectors](data-sources--app_setting--reference--group-001.md#canonical-bfb0514ba491a2b7f610f7e870c9822b9cb139f20ac36ff1c342417badb29b45) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors.metric` | [app_type_settings.timeseries_analyses_setting.metric_selectors.metric](data-sources--app_setting--reference--group-001.md#canonical-dcbf8ea93e9d67ac16f59916750ea6e9a371d27f0cb28d34404b24ca51254f5a) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source` | [app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source](data-sources--app_setting--reference--group-001.md#canonical-79a7554df015af16001737bf41a151c88f11e8916e1f6789825b28a17327c9ac) |
| `app_type_settings.user_behavior_analysis_setting` | [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-c50e06623cb5075b8ae0bfe2a30c8971532cf22244ebc177bfd1a5d98a3b0810) |
| `app_type_settings.user_behavior_analysis_setting.disable_detection` | [app_type_settings.user_behavior_analysis_setting.disable_detection](data-sources--app_setting--reference--group-001.md#canonical-ff00c84e07a2e930219cdab47263b53f482fe388cb995323d674c27672f4797f) |
| `app_type_settings.user_behavior_analysis_setting.disable_learning` | [app_type_settings.user_behavior_analysis_setting.disable_learning](data-sources--app_setting--reference--group-001.md#canonical-8a3dd066a611056a58cc2902f1f07138007c247fbc0ff76ebcfeb3ef70089b3c) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection` | [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-bc85555f6a0820bca834e558f636292d60f2dab1735666d37485b09f3055b407) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic` | [app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic](data-sources--app_setting--reference--group-001.md#canonical-6db7228876059f732d9524ed1fac4ef92dc89f20e172f5b36bb683b83068542f) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period` | [app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period](data-sources--app_setting--reference--group-001.md#canonical-3fed33963da6c5c7916242433d1ea38eb335eb1bcd3001d687e8cfa10e8bbc66) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection](data-sources--app_setting--reference--group-001.md#canonical-78335db740912c6f2e26986166d71e7dc5e7f7287296a0937db5d27c876de350) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity](data-sources--app_setting--reference--group-001.md#canonical-7945bdb6d7507ff52f97a180880042a2b61c7a8e09e29f9cfb5ef539a4eefb9a) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity](data-sources--app_setting--reference--group-001.md#canonical-3091560e8a3e2d3d4a7955eb25d8b0c25c1bbfd9c4a299bb8a8dd06edb1093cb) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity](data-sources--app_setting--reference--group-001.md#canonical-9c8e4707f33001ad86e413cc88f9d2486ddbd313d02be1fae48336334ed11557) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation](data-sources--app_setting--reference--group-001.md#canonical-18ddf6dcb04d28f9344c8b428d4df2299aca7f208c826342d7f9dc874d1819cc) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity](data-sources--app_setting--reference--group-001.md#canonical-1d7103188cda3a7672f8f5cbc3e870f093f413210f6d916f5934727d8db7e60a) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit](data-sources--app_setting--reference--group-001.md#canonical-98421de04436ed9a40ada6d835ff448de73651f1d7ce1c9314122926b068b4d9) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity](data-sources--app_setting--reference--group-001.md#canonical-c61f3be91bde41941721ba29b365c630e56de542b1444c4b9a008d8b5378414a) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity](data-sources--app_setting--reference--group-001.md#canonical-1a1063bd53dee8611ef36c436d758948a1c4880b9ccaea1352ca914d7dcd4e78) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity](data-sources--app_setting--reference--group-001.md#canonical-3a51a610478fc4634dd86facbfaea985a2d66e70459ce44e1a89d4d9744e503f) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold](data-sources--app_setting--reference--group-001.md#canonical-8bb08fb315f940ee4dbf64328a87bf1d4128710912357c38df0fb631d114abce) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity](data-sources--app_setting--reference--group-001.md#canonical-e5e0c73812d80aa6291397ab85520f389bcf8598183d15ace8aeb846a1f001f3) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold](data-sources--app_setting--reference--group-001.md#canonical-7e27aea08a8089e48b304e7080732e39850e7211f53087fe96db5091249c53b3) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation](data-sources--app_setting--reference--group-001.md#canonical-0665bad202e24d6a18af734f14136949c965e937b5286310603577a365fc1d8b) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-16a4bcb9f61ddb9ee99ccd920232c0e9518391c18d951c5a61770fd161889b98) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high](data-sources--app_setting--reference--group-001.md#canonical-ddc405a82e3211ca7d914a201985dad4201511b656e3dd7cab2c5a764f85621b) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low](data-sources--app_setting--reference--group-001.md#canonical-41cea5f716e452b621676dadc45cd51bb730fb83c5d140ed7e8e70a178a1c2ac) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium](data-sources--app_setting--reference--group-001.md#canonical-ab77c3910f17bd520ab1b97ae2281eac5735996fcc060672e4b0739aac488d7a) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom](data-sources--app_setting--reference--group-001.md#canonical-f6f41b3f2fac59f6a9d068e34773df2eac6861c7bd55ee5730be6fd8ecab001d) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold](data-sources--app_setting--reference--group-001.md#canonical-cf44df987e0160160a759c0d04859ef2f1d7c75708638dd90ae62ab1e35f2304) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit](data-sources--app_setting--reference--group-001.md#canonical-d88862c337c875137bb4cf32a6fd77129dabd60b9dc4f7cf952df3a484c9c436) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity](data-sources--app_setting--reference--group-001.md#canonical-b3ba13f7212696dcc9fbad8abc70cd97329eba6dba83016966191f06e7dc279e) |
| `app_type_settings.user_behavior_analysis_setting.enable_learning` | [app_type_settings.user_behavior_analysis_setting.enable_learning](data-sources--app_setting--reference--group-001.md#canonical-9a41fc06b0c32533b3dfc11eeea13d248c8af11aec9d6b623613ee6da28f989f) |
| `description` | [description](data-sources--app_setting--reference--group-001.md#canonical-15d34d40e9d8140171a454ce7fcce73c712694bf023947e2ac5d9dfaaea822d7) |
| `id` | [id](data-sources--app_setting--reference--group-001.md#canonical-00e6c8929f1bba22273b75284dde3e653682feaf2c778ecbcfd0e91da4ebfe8c) |
| `labels` | [labels](data-sources--app_setting--reference--group-001.md#canonical-c9b98a97e2016268a5044f33cd12de26ad0db623c2675bcfb32e2d3a3bcbaf46) |
| `name` | [name](data-sources--app_setting--reference--group-001.md#canonical-85e625374728786c79dcb40e91d4fe734f3b1d858a18e404eb45c7b7dd843955) |
| `namespace` | [namespace](data-sources--app_setting--reference--group-001.md#canonical-a31a3eda72572eaef481e42d773aa64a88ad4989b0505fa9530ba582023c3b96) |

<a id="canonical-cb16866832345c135d766fd5dbcb4cbc81c0c06b02b485e59f9e253480cf51d7"></a>

## Next pages — Property reference / b536487e6ba7 / 11

- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ebf89f5210fc39be62c29e015bbfafc21bfb0d7d0897b19e19ce5ff1adf9120"></a>

## app_type_settings — app_type_settings / d0d7c117d794 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- app_type_settings

<a id="canonical-210def4d3fabd9fe042194cfc1601c398e417c724682cbeb123108571604cb77"></a>

Type: `"list"`. Computed.

List of settings to enable for each AppType, given instance of AppType Exist in this Namespace.

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

<a id="canonical-ab1525e497c09b78af4f5601af7f4ff71a70f7742de8413f35a7315af753391b"></a>

## Direct properties — app_type_settings / d0d7c117d794 / 3

- [app_type_ref](data-sources--app_setting--reference--group-001.md#canonical-5064cf58464cbc5e5f7f000bbc5b489c5d7beaa40a6c4aa4ccda4a4de5d5bcc9): complete subsection reference.

- [business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-a05f10690ea19b13c7788522700f56acbd10ce64c2f7722de9609faff408d6fb): complete subsection reference.

- [timeseries_analyses_setting](data-sources--app_setting--reference--group-001.md#canonical-7ab6c8445bfa70fb601160a646df7fe7a3fcc636fcd77479c56c3cae493ac9bb): complete subsection reference.

- [user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82): complete subsection reference.

<a id="canonical-24bcb3e5e09b33737fa260f8f1d082a753bdf8a6021e582771f92ef29ef0e035"></a>

## Next pages — app_type_settings / d0d7c117d794 / 4

- [app_type_settings.app_type_ref](data-sources--app_setting--reference--group-001.md#canonical-5064cf58464cbc5e5f7f000bbc5b489c5d7beaa40a6c4aa4ccda4a4de5d5bcc9)
- [app_type_settings.business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-a05f10690ea19b13c7788522700f56acbd10ce64c2f7722de9609faff408d6fb)
- [app_type_settings.timeseries_analyses_setting](data-sources--app_setting--reference--group-001.md#canonical-7ab6c8445bfa70fb601160a646df7fe7a3fcc636fcd77479c56c3cae493ac9bb)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-5064cf58464cbc5e5f7f000bbc5b489c5d7beaa40a6c4aa4ccda4a4de5d5bcc9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf689f6f1f8025d92c9732296ab704e18d2f6b2c3b7649cea90f6ff977c5c04b"></a>

## app_type_settings.app_type_ref — app_type_settings.app_type_ref / 55838213d047 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- app_type_settings.app_type_ref

<a id="canonical-1bdde31e739f7686922a37e3b789ffe152bfc97fa36d246b06d2bd65e79a56c1"></a>

Type: `"list"`. Computed.

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

<a id="canonical-b38e9c8c2bbabb00bd5e12f198776ebc1175dc88630039dde983bc6e3d1b580b"></a>

## Direct properties — app_type_settings.app_type_ref / 55838213d047 / 3

<a id="canonical-e066be21e255079614835155b795e5c8142de419a845e6c63b341ccc5f8d7179"></a>

<a id="canonical-1da31bbb6c7c0e19a01e81609b2a348505e60c3489d4cfa83927a179db6c736f"></a>

## kind property — app_type_settings.app_type_ref / 55838213d047 / 4

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

<a id="canonical-d294c0632e30e182b4c92840af4d4569fa021b5c1cdc68ba7842397cb978585a"></a>

<a id="canonical-1003a290f05ac6b6925ed69cda24b0cd040f563d56e90137f079d3e2d9dcab8c"></a>

## name property — app_type_settings.app_type_ref / 55838213d047 / 5

Type: `"string"`. Computed.

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

<a id="canonical-70741bdf73b1f2d75cfa8e036068d92773947b0369b8e31c803f996990f44606"></a>

<a id="canonical-eb58120b3812d4e09ab2ed6cba420171f66f700c95f436785810724c8863c98b"></a>

## namespace property — app_type_settings.app_type_ref / 55838213d047 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-b7ad8663bbb218d81867d211581e32c0291e71881ae80b421f0448456867366a"></a>

<a id="canonical-ab0fb8b38390b7f44989e38cf21f7e1452e4758cd1d53592bd6f75d69b09696e"></a>

## tenant property — app_type_settings.app_type_ref / 55838213d047 / 7

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

<a id="canonical-d43e59631db6e5b2f43218e9dd48698017e367a9ee91daf87f117af4a3360fa9"></a>

<a id="canonical-fe3a2c46a8ea215c45cd2146dd6329d8e55dc7d51d8b277e982503b8d6f70729"></a>

## uid property — app_type_settings.app_type_ref / 55838213d047 / 8

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

<a id="canonical-276f0c0c296a299d80654e3d1d909b6c15b60a8ec65aa95e2c30dff4535b20ae"></a>

## Next pages — app_type_settings.app_type_ref / 55838213d047 / 9

- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-a05f10690ea19b13c7788522700f56acbd10ce64c2f7722de9609faff408d6fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f1c794cdcb319effea8928d85339c0598943c15932e05c641ad5dd2d0827b33"></a>

## app_type_settings.business_logic_markup_setting — app_type_settings.business_logic_markup_setting / 79079274f4eb / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- app_type_settings.business_logic_markup_setting

<a id="canonical-6f6371980484120705576d38867d983396fe8e3f51d363cb359274125bb98797"></a>

Type: `"single"`. Computed.

Settings specifying how API Discovery will be performed.

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

<a id="canonical-13021126fb25d8693a804f8be55bae741401ab26ce4603da9ff6a68ce44be7fc"></a>

## Direct properties — app_type_settings.business_logic_markup_setting / 79079274f4eb / 3

- [disable_spec](data-sources--app_setting--reference--group-001.md#canonical-20db558610e576b91d6e9536347481ae46a8b8a878b84ac37697d1d4d7bd21e2): complete subsection reference.

- [enable](data-sources--app_setting--reference--group-001.md#canonical-e779a949e3afcd0a4ad1fd0111c400aa6be696b6d31b2ae591d26b48ddf6d8d1): complete subsection reference.

<a id="canonical-5ee96dd37178fd3f1f466488c0bb6dcaf9f33cef8faa2b19019cbcfb3329913c"></a>

## Next pages — app_type_settings.business_logic_markup_setting / 79079274f4eb / 4

- [app_type_settings.business_logic_markup_setting.disable_spec](data-sources--app_setting--reference--group-001.md#canonical-20db558610e576b91d6e9536347481ae46a8b8a878b84ac37697d1d4d7bd21e2)
- [app_type_settings.business_logic_markup_setting.enable](data-sources--app_setting--reference--group-001.md#canonical-e779a949e3afcd0a4ad1fd0111c400aa6be696b6d31b2ae591d26b48ddf6d8d1)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-20db558610e576b91d6e9536347481ae46a8b8a878b84ac37697d1d4d7bd21e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-068d1fd8e7b633cabe9438bb6192d83006a289c7acbf50bfcf3d9f4a335033c1"></a>

## app_type_settings.business_logic_markup_setting.disable_spec — app_type_settings.business_logic_markup_setting.disable_spec / 2d7b400a3019 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-a05f10690ea19b13c7788522700f56acbd10ce64c2f7722de9609faff408d6fb)
- app_type_settings.business_logic_markup_setting.disable_spec

<a id="canonical-73524c79662346538ea52b585fb169debe1fc201467ce6a76ccddbcfc3b7d333"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-0599d9f0353a4421f7b0093623cace89128b41575a254db4f6236322f9e6def7"></a>

## Direct properties — app_type_settings.business_logic_markup_setting.disable_spec / 2d7b400a3019 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-507f324beb5ce61ca758ddd34537db9583e91423613a84232497c71ebe1b1fd0"></a>

## Next pages — app_type_settings.business_logic_markup_setting.disable_spec / 2d7b400a3019 / 4

- [app_type_settings.business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-a05f10690ea19b13c7788522700f56acbd10ce64c2f7722de9609faff408d6fb)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-e779a949e3afcd0a4ad1fd0111c400aa6be696b6d31b2ae591d26b48ddf6d8d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9cbeb781eeaab02856a2b2a57c5a1cdfe16ec9eae2a52e58f36440b4dd2e0960"></a>

## app_type_settings.business_logic_markup_setting.enable — app_type_settings.business_logic_markup_setting.enable / f431461ae3ef / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-a05f10690ea19b13c7788522700f56acbd10ce64c2f7722de9609faff408d6fb)
- app_type_settings.business_logic_markup_setting.enable

<a id="canonical-ac49644d05edd72d9e223f9a6f6427218aff5352f793926e29623c1f125a342e"></a>

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

<a id="canonical-a67223e39c847cab6a255b3562afe2f2d2ae429d5f200ae07c34e5e99dc0ec60"></a>

## Direct properties — app_type_settings.business_logic_markup_setting.enable / f431461ae3ef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f08a317e45cd286963fa1ef5f0b2456981f2a7f2242094a27917e02f674bdec5"></a>

## Next pages — app_type_settings.business_logic_markup_setting.enable / f431461ae3ef / 4

- [app_type_settings.business_logic_markup_setting](data-sources--app_setting--reference--group-001.md#canonical-a05f10690ea19b13c7788522700f56acbd10ce64c2f7722de9609faff408d6fb)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-7ab6c8445bfa70fb601160a646df7fe7a3fcc636fcd77479c56c3cae493ac9bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8067077d66dfccbfb1e93bbd3b4a16f5c6412b763e4882f65ff64e91b0a20a41"></a>

## app_type_settings.timeseries_analyses_setting — app_type_settings.timeseries_analyses_setting / 0acdc36508b3 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- app_type_settings.timeseries_analyses_setting

<a id="canonical-9c10645de88906ef8cd8dbbcb3cac78a7aab6c6e2136329597b7c3e4324f9a26"></a>

Type: `"single"`. Computed.

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

<a id="canonical-6d0fdbdf6307d97aadacf8b7c4eb2983fdc03bdd1ffc2140320503abf8c61a38"></a>

## Direct properties — app_type_settings.timeseries_analyses_setting / 0acdc36508b3 / 3

- [metric_selectors](data-sources--app_setting--reference--group-001.md#canonical-53d954d98c14bc2eb61d021f981da054085ac0cf549f27226cf794235cdd45ec): complete subsection reference.

<a id="canonical-6e47c8c080cfb639871543f5e1e992a27037b8fb79cf3543155dd75c42da88ba"></a>

## Next pages — app_type_settings.timeseries_analyses_setting / 0acdc36508b3 / 4

- [app_type_settings.timeseries_analyses_setting.metric_selectors](data-sources--app_setting--reference--group-001.md#canonical-53d954d98c14bc2eb61d021f981da054085ac0cf549f27226cf794235cdd45ec)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-53d954d98c14bc2eb61d021f981da054085ac0cf549f27226cf794235cdd45ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00b1fe236c3528bdc76b4fa882710b5e8a127d40cdf5fa55abd45e148841cdb5"></a>

## app_type_settings.timeseries_analyses_setting.metric_selectors — app_type_settings.timeseries_analyses_setting.metric_selectors / e0cfef356060 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.timeseries_analyses_setting](data-sources--app_setting--reference--group-001.md#canonical-7ab6c8445bfa70fb601160a646df7fe7a3fcc636fcd77479c56c3cae493ac9bb)
- app_type_settings.timeseries_analyses_setting.metric_selectors

<a id="canonical-bfb0514ba491a2b7f610f7e870c9822b9cb139f20ac36ff1c342417badb29b45"></a>

Type: `"list"`. Computed.

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

<a id="canonical-6cc0d928316a293cf901f8f40489ba07ea7c6910121960e23a82db6073e9aec1"></a>

## Direct properties — app_type_settings.timeseries_analyses_setting.metric_selectors / e0cfef356060 / 3

<a id="canonical-dcbf8ea93e9d67ac16f59916750ea6e9a371d27f0cb28d34404b24ca51254f5a"></a>

<a id="canonical-1b3b363eca3e2631fdad69c3402924f4d7d834fc50e096ee7ce22595525d437a"></a>

## metric property — app_type_settings.timeseries_analyses_setting.metric_selectors / e0cfef356060 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-79a7554df015af16001737bf41a151c88f11e8916e1f6789825b28a17327c9ac"></a>

<a id="canonical-629c407254f6f3154718d3242db6142d2f08ef5d0b082ffab98c071fb480e33c"></a>

## metrics_source property — app_type_settings.timeseries_analyses_setting.metric_selectors / e0cfef356060 / 5

Type: `"string"`. Computed.

\[Enum: NONE|NODES|EDGES|VIRTUAL\_HOSTS\] Supported sources from which Metrics can be analyzed All
edges in the service mesh graph. Metrics are analyzed separately between all source and destination
service combinations. Possible values are \`NONE\`, \`NODES\`, \`EDGES\`, \`VIRTUAL\_HOSTS\`.

Upstream description:

Supported sources from which Metrics can be analyzed

All edges in the service mesh graph. Metrics are analyzed separately between all source and
destination service combinations.

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

<a id="canonical-860c9dd998d632573e796ebda5f9141b2740e86815e54c20ec15aa8dea731c8a"></a>

## Next pages — app_type_settings.timeseries_analyses_setting.metric_selectors / e0cfef356060 / 6

- [app_type_settings.timeseries_analyses_setting](data-sources--app_setting--reference--group-001.md#canonical-7ab6c8445bfa70fb601160a646df7fe7a3fcc636fcd77479c56c3cae493ac9bb)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcf38bb437fa847ac71f7803c245de42d4d60011323711f35e52d581585be28a"></a>

## app_type_settings.user_behavior_analysis_setting — app_type_settings.user_behavior_analysis_setting / 63b65b1dbd49 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- app_type_settings.user_behavior_analysis_setting

<a id="canonical-c50e06623cb5075b8ae0bfe2a30c8971532cf22244ebc177bfd1a5d98a3b0810"></a>

Type: `"single"`. Computed.

Configuration for user behavior analysis.

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

<a id="canonical-7e94c7ddb3759551f9b6f15e53ae27062bf245b5d63c8cd0af23ecfffc2dd8c2"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting / 63b65b1dbd49 / 3

- [disable_detection](data-sources--app_setting--reference--group-001.md#canonical-c1483461da3e54838c7db04680e92151256e383026c4768b4de5c5c2f0663bac): complete subsection reference.

- [disable_learning](data-sources--app_setting--reference--group-001.md#canonical-0289a8efee4616a3e08682c3a70071b3cca4b520fa661e606ac88e28ea065890): complete subsection reference.

- [enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731): complete subsection reference.

- [enable_learning](data-sources--app_setting--reference--group-001.md#canonical-c2e060c30c0017163d7ae27b37eb3dcc05706ffa0e7559f400eea7dff188cdc4): complete subsection reference.

<a id="canonical-afaeee2372bf750e6090cb596f9f85d21b23399542592a220dda63a67741e5c6"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting / 63b65b1dbd49 / 4

- [app_type_settings.user_behavior_analysis_setting.disable_detection](data-sources--app_setting--reference--group-001.md#canonical-c1483461da3e54838c7db04680e92151256e383026c4768b4de5c5c2f0663bac)
- [app_type_settings.user_behavior_analysis_setting.disable_learning](data-sources--app_setting--reference--group-001.md#canonical-0289a8efee4616a3e08682c3a70071b3cca4b520fa661e606ac88e28ea065890)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [app_type_settings.user_behavior_analysis_setting.enable_learning](data-sources--app_setting--reference--group-001.md#canonical-c2e060c30c0017163d7ae27b37eb3dcc05706ffa0e7559f400eea7dff188cdc4)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-c1483461da3e54838c7db04680e92151256e383026c4768b4de5c5c2f0663bac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-970b58e5bcbc5f4fc8f97ba1bc2bbbef9fdad7869051e540c6e20288a376f3fa"></a>

## app_type_settings.user_behavior_analysis_setting.disable_detection — app_type_settings.user_behavior_analysis_setting.disable_detection / 3a3f1090d58e / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- app_type_settings.user_behavior_analysis_setting.disable_detection

<a id="canonical-ff00c84e07a2e930219cdab47263b53f482fe388cb995323d674c27672f4797f"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-12259d68d81ca923d8d9743df0516e65c0ee5ed73f63f2ced96ccc10c9dd8d6e"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.disable_detection / 3a3f1090d58e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-242c7d9337fd84a39f48965a596e079835e05a3556bbda253b7abd0d211fa83b"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.disable_detection / 3a3f1090d58e / 4

- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-0289a8efee4616a3e08682c3a70071b3cca4b520fa661e606ac88e28ea065890"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18803fde55c94362b3f7f4b9501cb293c5eb31de9098e78039c9bfe69cb5e5b1"></a>

## app_type_settings.user_behavior_analysis_setting.disable_learning — app_type_settings.user_behavior_analysis_setting.disable_learning / 8b24ab3ab1cf / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- app_type_settings.user_behavior_analysis_setting.disable_learning

<a id="canonical-8a3dd066a611056a58cc2902f1f07138007c247fbc0ff76ebcfeb3ef70089b3c"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-fa487af1e8e5a2e95646fbdc010e68bf24ef6d0e613ee07721b1c18e10d156bc"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.disable_learning / 8b24ab3ab1cf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ab05900f28e51781cb90b89b6f81949683416bad36a854b6fca8dc9fe123e688"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.disable_learning / 8b24ab3ab1cf / 4

- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b08d8b57160bd2852e5fa8ed1eb138d30b2164c9cb5c2848be68ba68d295941"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection — app_type_settings.user_behavior_analysis_setting.enable_detection / 197489c3f37b / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- app_type_settings.user_behavior_analysis_setting.enable_detection

<a id="canonical-bc85555f6a0820bca834e558f636292d60f2dab1735666d37485b09f3055b407"></a>

Type: `"single"`. Computed.

Various factors about user activity are monitored and analysed to determine malicious users. These
settings allow tuning those factors used by the system to detect malicious users.

Upstream description:

Various factors about user activity are monitored and analysed to determine malicious users. These
settings allow tuning those factors used by the system to detect malicious users.

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

<a id="canonical-951a5a6464463aaabf5f101de825eee54b71da753ff8184190b0d0fad4e8053a"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection / 197489c3f37b / 3

- [bola_detection_automatic](data-sources--app_setting--reference--group-001.md#canonical-43ec66f3350757d65a6c75ef520acbab6804eb2f5d66757ffc96cbbdd19bb38d): complete subsection reference.

<a id="canonical-3fed33963da6c5c7916242433d1ea38eb335eb1bcd3001d687e8cfa10e8bbc66"></a>

<a id="canonical-3a1d708762b5bcfc30c1655729a6e41f057494fdce71070948e4bac59c39cbaf"></a>

## cooling_off_period property — app_type_settings.user_behavior_analysis_setting.enable_detection / 197489c3f37b / 4

Type: `"number"`. Computed.

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

- [exclude_bola_detection](data-sources--app_setting--reference--group-001.md#canonical-4962576894af7c0600b4bad35fcad543d3b0b29e3df62c95451c40b25d77a905): complete subsection reference.

- [exclude_bot_defense_activity](data-sources--app_setting--reference--group-001.md#canonical-503fc4af86ce0ad630baaa9b440cb7c71a13d5e8ef1bf8f43b61152a62de0469): complete subsection reference.

- [exclude_failed_login_activity](data-sources--app_setting--reference--group-001.md#canonical-026a7dd3ce7f476c1a4315cab03c9119011b756cd09ca2c0c6754297e6ac554d): complete subsection reference.

- [exclude_forbidden_activity](data-sources--app_setting--reference--group-001.md#canonical-bd3f2efab69e90009db4a26d436dc5ed000583b075fda840e74ef2ed4927f413): complete subsection reference.

- [exclude_ip_reputation](data-sources--app_setting--reference--group-001.md#canonical-57976cf63155165bbf1e3d3b7860b6159cde072e4ec4612c1d9dd82bdf1d5767): complete subsection reference.

- [exclude_non_existent_url_activity](data-sources--app_setting--reference--group-001.md#canonical-e4430615debeb59fc256ddc70e3f755518d50b7b3f342681cf2ec4ddb0cd53ac): complete subsection reference.

- [exclude_rate_limit](data-sources--app_setting--reference--group-001.md#canonical-ccfaf5ae5f3e5bfd18ec0cd7ded61a27563acab72710eb8463b3ec2d7697596b): complete subsection reference.

- [exclude_waf_activity](data-sources--app_setting--reference--group-001.md#canonical-a71927c2e8123765a8b954a7532dc3d9068265cc54c13f0daa18eb4e183f1d2f): complete subsection reference.

- [include_bot_defense_activity](data-sources--app_setting--reference--group-001.md#canonical-f8613ed9c38621042602270ba55cc47c2e31e2f70b64ef6c2a0d3b6cea28de80): complete subsection reference.

- [include_failed_login_activity](data-sources--app_setting--reference--group-001.md#canonical-030baa93591e4699c9a87739a46200e3a67955dced27aed90e4e9a16c9785647): complete subsection reference.

- [include_forbidden_activity](data-sources--app_setting--reference--group-001.md#canonical-f3e58573c1c8fc2a2ab21830933a073570669cb86594a49effca13e9b736b01d): complete subsection reference.

- [include_ip_reputation](data-sources--app_setting--reference--group-001.md#canonical-14b4cd85158abf5b1020f196bf6288db358c073a2b8bfea7e1d916d07a2b1acf): complete subsection reference.

- [include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-d1411678995c34e8e8de5781e61ddb6a0bd8f3de96c4ff1d88886eda2e52e794): complete subsection reference.

- [include_non_existent_url_activity_custom](data-sources--app_setting--reference--group-001.md#canonical-57f87219e88e6c908e1de754b6009266562ca7c389fb91427945e125d1c949dc): complete subsection reference.

- [include_rate_limit](data-sources--app_setting--reference--group-001.md#canonical-64d6bc75e77134f22b9edaaf0b9f89182555a5b0de87bafaf0485fd0078b494c): complete subsection reference.

- [include_waf_activity](data-sources--app_setting--reference--group-001.md#canonical-6c754dce44526b5758e7d055d012562ac9f8cd3ffc5172839597814f716ecce9): complete subsection reference.

<a id="canonical-02a7bbb15a317ddca8ce59acf1749e6fd4c4b7f99d3a8b6a71dd7cf5ae955e5e"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection / 197489c3f37b / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic](data-sources--app_setting--reference--group-001.md#canonical-43ec66f3350757d65a6c75ef520acbab6804eb2f5d66757ffc96cbbdd19bb38d)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection](data-sources--app_setting--reference--group-001.md#canonical-4962576894af7c0600b4bad35fcad543d3b0b29e3df62c95451c40b25d77a905)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity](data-sources--app_setting--reference--group-001.md#canonical-503fc4af86ce0ad630baaa9b440cb7c71a13d5e8ef1bf8f43b61152a62de0469)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity](data-sources--app_setting--reference--group-001.md#canonical-026a7dd3ce7f476c1a4315cab03c9119011b756cd09ca2c0c6754297e6ac554d)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity](data-sources--app_setting--reference--group-001.md#canonical-bd3f2efab69e90009db4a26d436dc5ed000583b075fda840e74ef2ed4927f413)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation](data-sources--app_setting--reference--group-001.md#canonical-57976cf63155165bbf1e3d3b7860b6159cde072e4ec4612c1d9dd82bdf1d5767)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity](data-sources--app_setting--reference--group-001.md#canonical-e4430615debeb59fc256ddc70e3f755518d50b7b3f342681cf2ec4ddb0cd53ac)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit](data-sources--app_setting--reference--group-001.md#canonical-ccfaf5ae5f3e5bfd18ec0cd7ded61a27563acab72710eb8463b3ec2d7697596b)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity](data-sources--app_setting--reference--group-001.md#canonical-a71927c2e8123765a8b954a7532dc3d9068265cc54c13f0daa18eb4e183f1d2f)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity](data-sources--app_setting--reference--group-001.md#canonical-f8613ed9c38621042602270ba55cc47c2e31e2f70b64ef6c2a0d3b6cea28de80)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity](data-sources--app_setting--reference--group-001.md#canonical-030baa93591e4699c9a87739a46200e3a67955dced27aed90e4e9a16c9785647)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity](data-sources--app_setting--reference--group-001.md#canonical-f3e58573c1c8fc2a2ab21830933a073570669cb86594a49effca13e9b736b01d)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation](data-sources--app_setting--reference--group-001.md#canonical-14b4cd85158abf5b1020f196bf6288db358c073a2b8bfea7e1d916d07a2b1acf)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-d1411678995c34e8e8de5781e61ddb6a0bd8f3de96c4ff1d88886eda2e52e794)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom](data-sources--app_setting--reference--group-001.md#canonical-57f87219e88e6c908e1de754b6009266562ca7c389fb91427945e125d1c949dc)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit](data-sources--app_setting--reference--group-001.md#canonical-64d6bc75e77134f22b9edaaf0b9f89182555a5b0de87bafaf0485fd0078b494c)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity](data-sources--app_setting--reference--group-001.md#canonical-6c754dce44526b5758e7d055d012562ac9f8cd3ffc5172839597814f716ecce9)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-43ec66f3350757d65a6c75ef520acbab6804eb2f5d66757ffc96cbbdd19bb38d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd12df995dc32f7d3b94045b201fb7660d44c80534a10ba0f82120e2071a04e9"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic — app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection / 86fa6927b5ab / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic

<a id="canonical-6db7228876059f732d9524ed1fac4ef92dc89f20e172f5b36bb683b83068542f"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-20ceeec49ae66543ff7fe523863ad9c40226de4bfa80e5c7df9344d6f3a2fb61"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection / 86fa6927b5ab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0eb2982ac012177454736a4aa826445b1feb2c496dd139abb397381d5965b214"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection / 86fa6927b5ab / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-4962576894af7c0600b4bad35fcad543d3b0b29e3df62c95451c40b25d77a905"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82bbf4a9a4b4bcd77ac2faecb3a803e740fa5f1cc658f7978ee61f3ff37d1970"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_d / 8125c4863f83 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection

<a id="canonical-78335db740912c6f2e26986166d71e7dc5e7f7287296a0937db5d27c876de350"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-afc7285bfbdde4635fdd8b94a6555a7a7b1d1de39dbc622476a783848ee2786c"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_d / 8125c4863f83 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-37945f1f2d90eae6e2fd44f1929fbd3cfd29ddd280499b216b65ccd9c4897ea7"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_d / 8125c4863f83 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-503fc4af86ce0ad630baaa9b440cb7c71a13d5e8ef1bf8f43b61152a62de0469"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a2c16556a27a9a1fee3f7fbff64750d52badb5e7c7871a2df35d4ecf29191f5"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_de / 8f6a26e974b7 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity

<a id="canonical-7945bdb6d7507ff52f97a180880042a2b61c7a8e09e29f9cfb5ef539a4eefb9a"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-c8ace71c2689f38eac8a19df50b50111b890254668bc3c5d3ea0b870b7a8d4c2"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_de / 8f6a26e974b7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e270895f5d076324b66d66588a5cb16122295e0d13a965d51249ce96a0093c89"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_de / 8f6a26e974b7 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-026a7dd3ce7f476c1a4315cab03c9119011b756cd09ca2c0c6754297e6ac554d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57e2564fcdcd779d06060f73ba37cfae30cad1c319b6b8612d83b2992d304fc9"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed / 9feb3f9a37cb / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity

<a id="canonical-3091560e8a3e2d3d4a7955eb25d8b0c25c1bbfd9c4a299bb8a8dd06edb1093cb"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-855c16f2d92a0286fca787d6f1e488e7c7e933aff6afb9a84d116a95c325a5fc"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed / 9feb3f9a37cb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12e0ef4150e0699bfc892e60b73991e746f6d184f6fcb99a178e8e327a2270e2"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed / 9feb3f9a37cb / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-bd3f2efab69e90009db4a26d436dc5ed000583b075fda840e74ef2ed4927f413"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4727f31c77f5bf74f0036627d1409879f973ef23133ecc26e5a63c1935ccda6"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbid / cb3f9d52656b / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity

<a id="canonical-9c8e4707f33001ad86e413cc88f9d2486ddbd313d02be1fae48336334ed11557"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-e13e576a7dc5f64d38f7d7731f0b8d3a753a40cf85ec6f881efc482ffe5176ce"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbid / cb3f9d52656b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8ac8650c4217c16b01e175830cf733c0a4f85deaf3309970540839a0573ac8d1"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbid / cb3f9d52656b / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-57976cf63155165bbf1e3d3b7860b6159cde072e4ec4612c1d9dd82bdf1d5767"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f558870d6d27fa2d1382bea4fd9b4d214a38627011a233cc01828107b7b730c6"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_rep / 22f7872ff651 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation

<a id="canonical-18ddf6dcb04d28f9344c8b428d4df2299aca7f208c826342d7f9dc874d1819cc"></a>

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

<a id="canonical-b79d491b9e74fa6ee6839d0c662ce1035f960e7f1f5dea98db4d3f00b5c9227f"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_rep / 22f7872ff651 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0f7ac901724b556e74cab337f7822910b1dbf73e64fdd5093c7c072bf9eb2304"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_rep / 22f7872ff651 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-e4430615debeb59fc256ddc70e3f755518d50b7b3f342681cf2ec4ddb0cd53ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fdc595b215775630690d5cc5a2ab74696590c311c35aeddad06f31b2b5e3452e"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_ex / 4cfe96667724 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity

<a id="canonical-1d7103188cda3a7672f8f5cbc3e870f093f413210f6d916f5934727d8db7e60a"></a>

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

<a id="canonical-7392ab21e18b1f20ceccaa4591816173f4e6863f5385ab9fcaba7a99b7247db5"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_ex / 4cfe96667724 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-339a3123446ac5dd4445db38cbc29f38517f2b31fe319e391ef405cf16657760"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_ex / 4cfe96667724 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-ccfaf5ae5f3e5bfd18ec0cd7ded61a27563acab72710eb8463b3ec2d7697596b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ea10111c273bebbbe4c26350f8a6566158618856939fdf9f2823cae000fe656"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_l / 94e236656fba / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit

<a id="canonical-98421de04436ed9a40ada6d835ff448de73651f1d7ce1c9314122926b068b4d9"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-87c462b60259343cdd339df9df569bf733f69f86d1930b16ef5715f01fd69511"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_l / 94e236656fba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e6524fc17dffb84dbfab969bd6ecf2926467b9a1ce0cf22e5607dc5026a96959"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_l / 94e236656fba / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-a71927c2e8123765a8b954a7532dc3d9068265cc54c13f0daa18eb4e183f1d2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-167bc709a0596bb8ac691e71badfc2e2f1ca7d921056817fd3a28d52df527766"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_ac / 8cebc884be31 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity

<a id="canonical-c61f3be91bde41941721ba29b365c630e56de542b1444c4b9a008d8b5378414a"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-29a0d461de688c5f464731edfbc7db96a37af425634b72588485175a01bd6d7b"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_ac / 8cebc884be31 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f624cd9c4dbf165e7d18ec7e363a27a8900927841a5a92ebc9fdaf3c66d91adb"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_ac / 8cebc884be31 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-f8613ed9c38621042602270ba55cc47c2e31e2f70b64ef6c2a0d3b6cea28de80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d2a5de2ad61c9764000c5240f69ed29031bee41e9f602a5e4d4cef95cbb7e28"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_de / 036c723e1c14 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity

<a id="canonical-1a1063bd53dee8611ef36c436d758948a1c4880b9ccaea1352ca914d7dcd4e78"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-5b20d6bbb8012a364eda9d96afdeb85616512f0beb37a7f5838a19916ab73d41"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_de / 036c723e1c14 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f86bab694ba45228e8373303aa60f0cdea2fa3a842ab8a1283f23dd1fb7a330e"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_de / 036c723e1c14 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-030baa93591e4699c9a87739a46200e3a67955dced27aed90e4e9a16c9785647"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-321253a283bc6decb51d397fb2b1a9a798f4b1adf9fe4d0f1b8927082b887bf6"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed / 47d11a2ec4c3 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity

<a id="canonical-3a51a610478fc4634dd86facbfaea985a2d66e70459ce44e1a89d4d9744e503f"></a>

Type: `"single"`. Computed.

When enabled, the system monitors persistent failed login attempts from a user. A failed login is
detected if a request results in a response code of 401. These settings specify how to use failed
login activity to determine suspicious behavior.

Upstream description:

When enabled, the system monitors persistent failed login attempts from a user. A failed login is
detected if a request results in a response code of 401. These settings specify how to use failed
login activity to determine suspicious behavior.

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

<a id="canonical-10032616e4581e2316557ea2d2affaa7cd62d5d3e22d53fd81a0a71aa46b5710"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed / 47d11a2ec4c3 / 3

<a id="canonical-8bb08fb315f940ee4dbf64328a87bf1d4128710912357c38df0fb631d114abce"></a>

<a id="canonical-560224cf8dffe24cfd12758cf22f4fe21d266b7bf6825d28e1bf71fc4354977a"></a>

## login_failures_threshold property — app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed / 47d11a2ec4c3 / 4

Type: `"number"`. Computed.

The number of failed logins beyond which the system will flag this user as malicious.

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

<a id="canonical-dccc56991c91d44d72e3604b37c1e9790ad822cc48eb499ef2d51f0448239330"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed / 47d11a2ec4c3 / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-f3e58573c1c8fc2a2ab21830933a073570669cb86594a49effca13e9b736b01d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e99e8e6da6cf2797d10b2e5cc8c9f6857b8500a26735db2f466bf3e1e0aef485"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbid / 8acf86f31d4b / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity

<a id="canonical-e5e0c73812d80aa6291397ab85520f389bcf8598183d15ace8aeb846a1f001f3"></a>

Type: `"single"`. Computed.

When L7 policy rules are set up to disallow certain types of requests, the system monitors
persistent attempts from a user to send requests which result in policy denies. These settings
specify how to use disallowed request activity from a user to determine suspicious behavior.

Upstream description:

When L7 policy rules are set up to disallow certain types of requests, the system monitors
persistent attempts from a user to send requests which result in policy denies. These settings
specify how to use disallowed request activity from a user to determine suspicious behavior.

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

<a id="canonical-3ead668dd1255d3a406e4ea614675a56864a8e9c12a428a574d811517e129578"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbid / 8acf86f31d4b / 3

<a id="canonical-7e27aea08a8089e48b304e7080732e39850e7211f53087fe96db5091249c53b3"></a>

<a id="canonical-5d0e79647ecca3b21787a18840fa6ee5ea01b6213eaa2a453d3b0197278950df"></a>

## forbidden_requests_threshold property — app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbid / 8acf86f31d4b / 4

Type: `"number"`. Computed.

The number of forbidden requests beyond which the system will flag this user as malicious.

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

<a id="canonical-b35b5aba88b6886b31d329c05165d5529a8f586f94fbcc589360c3d032bd8eff"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbid / 8acf86f31d4b / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-14b4cd85158abf5b1020f196bf6288db358c073a2b8bfea7e1d916d07a2b1acf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e355e94c81e146effcbc8e34d42a6a6515ca78202f956eca74c30f88358a7fb7"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation — app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_rep / 4819f34640a7 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation

<a id="canonical-0665bad202e24d6a18af734f14136949c965e937b5286310603577a365fc1d8b"></a>

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

<a id="canonical-b4374fd52a9898e95848e65660a8253177828be2a57d37f4641fa9ada218eb78"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_rep / 4819f34640a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fda42de3c938a1079eb6141f2705517e6aa9957d0c78f66246d45972152d157f"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_rep / 4819f34640a7 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-d1411678995c34e8e8de5781e61ddb6a0bd8f3de96c4ff1d88886eda2e52e794"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecdbe1c085d8285c566e31a9a393c59a6e9226a16660c1423d0ec2800f917d64"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / c35b95a75545 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic

<a id="canonical-16a4bcb9f61ddb9ee99ccd920232c0e9518391c18d951c5a61770fd161889b98"></a>

Type: `"single"`. Computed.

Non-existent URL Automatic Activity Settings.

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

<a id="canonical-0edb6b27a88cde9e181b3ebe356894a9dfec98faab3026960103f4b1556f90af"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / c35b95a75545 / 3

- [high](data-sources--app_setting--reference--group-001.md#canonical-40612320d058badcf8de7f0690d3dcd524e2fa3b5bbb56c2126ce6b5a941312b): complete subsection reference.

- [low](data-sources--app_setting--reference--group-001.md#canonical-81b91beae8e6a0bbaffbeed99200de3c1ba82978d25a5d9aeaa76da2ee95d8a6): complete subsection reference.

- [medium](data-sources--app_setting--reference--group-001.md#canonical-58c71dd42be2732e5ee3e0bebcba1a45e55b6e3f9dd5813b1cd724734d3a7c73): complete subsection reference.

<a id="canonical-a56cebff499a84aeb8baeb63da9e032726f9557ad0f202a43b76d7ff4a6b1584"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / c35b95a75545 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high](data-sources--app_setting--reference--group-001.md#canonical-40612320d058badcf8de7f0690d3dcd524e2fa3b5bbb56c2126ce6b5a941312b)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low](data-sources--app_setting--reference--group-001.md#canonical-81b91beae8e6a0bbaffbeed99200de3c1ba82978d25a5d9aeaa76da2ee95d8a6)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium](data-sources--app_setting--reference--group-001.md#canonical-58c71dd42be2732e5ee3e0bebcba1a45e55b6e3f9dd5813b1cd724734d3a7c73)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-40612320d058badcf8de7f0690d3dcd524e2fa3b5bbb56c2126ce6b5a941312b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5dd7ce79ee1a5168825b26f2fe84a1c6145a382e8695e9f6ee505c22091bac49"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 49458b4e9661 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-d1411678995c34e8e8de5781e61ddb6a0bd8f3de96c4ff1d88886eda2e52e794)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high

<a id="canonical-ddc405a82e3211ca7d914a201985dad4201511b656e3dd7cab2c5a764f85621b"></a>

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

<a id="canonical-3e9f8700c3e683336a8633faea53713855ddc91581d118ac646bc1ebc09f18e5"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 49458b4e9661 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-29e4aa6c8137bf6c3cb2b6827ac26f850916ca64a67c1666da5ea896093653a8"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 49458b4e9661 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-d1411678995c34e8e8de5781e61ddb6a0bd8f3de96c4ff1d88886eda2e52e794)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-81b91beae8e6a0bbaffbeed99200de3c1ba82978d25a5d9aeaa76da2ee95d8a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68c5aa04f6a48843afff56a6b2546914b03931932638df355e847ad08dc4d549"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 19a188ddfc39 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-d1411678995c34e8e8de5781e61ddb6a0bd8f3de96c4ff1d88886eda2e52e794)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low

<a id="canonical-41cea5f716e452b621676dadc45cd51bb730fb83c5d140ed7e8e70a178a1c2ac"></a>

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

<a id="canonical-4e47a0be5d3e4cd076ed868c05889e907c743b28b01d40878eaa52cbb1c9d66a"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 19a188ddfc39 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e88fa6474db128e95a3f2c458cfe1bb1a0c1db87e906d41153cffe197e1b558e"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 19a188ddfc39 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-d1411678995c34e8e8de5781e61ddb6a0bd8f3de96c4ff1d88886eda2e52e794)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-58c71dd42be2732e5ee3e0bebcba1a45e55b6e3f9dd5813b1cd724734d3a7c73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c49cf9a8f99d08730bdeae3e3410794ab39aa44322f62877d4623a47b1e565d8"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 31cc1750db7a / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-d1411678995c34e8e8de5781e61ddb6a0bd8f3de96c4ff1d88886eda2e52e794)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium

<a id="canonical-ab77c3910f17bd520ab1b97ae2281eac5735996fcc060672e4b0739aac488d7a"></a>

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

<a id="canonical-97513a34abadbd8ad17720b949dceb3033eeb42994ce48c5a1e838ca5e5271c1"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 31cc1750db7a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ef12d1e7fcbd9486ebad31ca92edb56f174885ec35fe4ea5bfc06b1025d97cd1"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / 31cc1750db7a / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--reference--group-001.md#canonical-d1411678995c34e8e8de5781e61ddb6a0bd8f3de96c4ff1d88886eda2e52e794)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-57f87219e88e6c908e1de754b6009266562ca7c389fb91427945e125d1c949dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d13c5b713fa68b1c5e5848f1c4cb32832a550140c1a0e23895735c6b1574cc3"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / e50f475a7ea4 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom

<a id="canonical-f6f41b3f2fac59f6a9d068e34773df2eac6861c7bd55ee5730be6fd8ecab001d"></a>

Type: `"single"`. Computed.

Non-existent URL Custom Activity Setting.

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

<a id="canonical-a360e0c919c0157576b0056b3b6f1ff6568fc301d9fcd67f4a2861cf5c294682"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / e50f475a7ea4 / 3

<a id="canonical-cf44df987e0160160a759c0d04859ef2f1d7c75708638dd90ae62ab1e35f2304"></a>

<a id="canonical-d6aae00dc474752342ec6e25f84676012ae823de15e586f03c79867320be0433"></a>

## nonexistent_requests_threshold property — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / e50f475a7ea4 / 4

Type: `"number"`. Computed.

The percentage of non-existent requests beyond which the system will flag this user as malicious.

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

<a id="canonical-479bba9dab7d5e60b3168c94ce22f7a7a0e10bd8861e901c2fbbd0fd12aca793"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_ex / e50f475a7ea4 / 5

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-64d6bc75e77134f22b9edaaf0b9f89182555a5b0de87bafaf0485fd0078b494c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3c583ecaa2abfb82f54a640c43450f36adfc8eae69c2ccddf5fa75b9dac49df"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit — app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_l / 9bc74db0f1b5 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit

<a id="canonical-d88862c337c875137bb4cf32a6fd77129dabd60b9dc4f7cf952df3a484c9c436"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-f1b137f5cd8d46b91ad7f7f2b9ff55fdbcf8cce00117755f77e8f1eec5b2075e"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_l / 9bc74db0f1b5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-82eb4915511dea15863c62eded1f070d385d587124bc72655b6937aa1d8a4786"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_l / 9bc74db0f1b5 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-6c754dce44526b5758e7d055d012562ac9f8cd3ffc5172839597814f716ecce9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d54c941bb85e32278fef962dc340e9b4f52cf87db574a1117abbcba1173c0b84"></a>

## app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity — app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_ac / 3f2cca6e31f6 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity

<a id="canonical-b3ba13f7212696dcc9fbad8abc70cd97329eba6dba83016966191f06e7dc279e"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-07b13c8265cfefdf73ebc6205438f25decead55a2df18d50a087a057fe14a7fa"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_ac / 3f2cca6e31f6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e54b3f25337cbbe986583795c5ebc08b7b5e9a460bc3f47f15e78c22958ba30a"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_ac / 3f2cca6e31f6 / 4

- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--reference--group-001.md#canonical-35ed78a672241142eae153350767a22a9288a604238659d2ef1c0c89772ad731)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-c2e060c30c0017163d7ae27b37eb3dcc05706ffa0e7559f400eea7dff188cdc4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b20532b8042ac50134d851bfe5f7ba49d5dffc96c69972313b879b3470903b4"></a>

## app_type_settings.user_behavior_analysis_setting.enable_learning — app_type_settings.user_behavior_analysis_setting.enable_learning / e7f9e0e48b0a / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Property reference](data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [app_type_settings](data-sources--app_setting--reference--group-001.md#canonical-c16aab867d88641c3cf2f50442adacc8a9e44c67c6d35586fdeb7285cf24b3fd)
- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- app_type_settings.user_behavior_analysis_setting.enable_learning

<a id="canonical-9a41fc06b0c32533b3dfc11eeea13d248c8af11aec9d6b623613ee6da28f989f"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-525008a7e2c9da817d9faf1d78408cceb90b93670a5381f9babe73bae0cfe121"></a>

## Direct properties — app_type_settings.user_behavior_analysis_setting.enable_learning / e7f9e0e48b0a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1ff9e6cdbd0e6eed6a22d1df0b273acd2907e3df86f863d60c42c0e6a34b3ea9"></a>

## Next pages — app_type_settings.user_behavior_analysis_setting.enable_learning / e7f9e0e48b0a / 4

- [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--reference--group-001.md#canonical-f4ffd02257d2ac93638df7d1e16a1a1cae591646fdf97f4ce299c0cc06981d82)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
