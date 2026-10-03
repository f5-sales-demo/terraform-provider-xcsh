---
page_title: "xcsh_api_definition reference"
subcategory: "API Management"
description: "Complete grouped canonical reference for xcsh_api_definition reference."
---

# xcsh_api_definition reference

<a id="canonical-1013013311212220-2100012121010012-2012032113201303-3011001133121233-0013201202222002-0113112031031113-2000220121302133-2321231322222211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000321301011131-0210310010332013-2000020312230030-1122321110130203-3310212010121321-0203330321033001-3020103233310003-0112231122123030"></a>

## Property reference — Property reference / 320323313321 / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)
- Property reference

<a id="canonical-1011111021220023-3222011112330222-2311113232001112-1131020100321131-3230213211230133-0223202321332222-3032213103221311-3331231032131322"></a>

## Direct properties — Property reference / 320323313321 / 3

<a id="canonical-3102233200013303-1111100131032012-2100033110201321-0201100101011301-1103320212203011-2323221333010122-0313020330323322-1111232302213313"></a>

<a id="canonical-1032020310313123-2312313032320002-0333323021222322-2012332303310012-1223011022000300-3020323332133211-1230322313100310-1330002332203222"></a>

## annotations property — Property reference / 320323313321 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
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

- [api_inventory_exclusion_list](resources--api_definition--reference--group-001.md#canonical-0322111332303200-0102230312330331-2230231331102322-0111003030300213-3103033311100113-2312223222013222-0332133333003331-1230011112023110): complete subsection reference.

- [api_inventory_inclusion_list](resources--api_definition--reference--group-001.md#canonical-0323203300310302-0203231011212201-3122301133231102-1200330320231103-3013021220022111-2110033120132001-0322200121032302-1301222111101101): complete subsection reference.

<a id="canonical-1112202122030001-0110002103230202-3031112222220330-0113202333110003-3220220101100121-2113232130000023-0323102132322200-3032303332101313"></a>

<a id="canonical-2022233123020221-0221323110103331-2033301302202023-2022313231103321-0320322121133123-2230222231001203-1003310323101231-1102002113122013"></a>

## description property — Property reference / 320323313321 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1013000220232233-2132303201100201-0010223232311322-1102231133310223-3033000011131121-0232333111000003-3311310033013133-2221013231112031"></a>

<a id="canonical-3123313021310113-0211211230111231-0300321110310122-3013030030233220-1322113012012120-1300312221113133-1323113013300221-2011100223232001"></a>

## disable property — Property reference / 320323313321 / 6

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

<a id="canonical-0113223200321130-1021103333203201-3031113103200110-3033031020313003-3222303000302122-0020112333300303-3331023102220132-0033011221203300"></a>

<a id="canonical-2111212112302213-2011230031031010-1113312020030030-3313313003311203-0010332100102211-3003332121321003-2110201120031023-2222130113110201"></a>

## ID property — Property reference / 320323313321 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0321010310303020-3323012320322330-0133223201220010-2031023310222323-3230232312222303-0323233231223312-1132021130332122-3032030120030211"></a>

<a id="canonical-0120233033001313-1210000103231333-3333032110221011-2200333302002212-0132310203302221-0222120302312003-1201112222202230-1132300213300323"></a>

## labels property — Property reference / 320323313321 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
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

- [mixed_schema_origin](resources--api_definition--reference--group-001.md#canonical-2013020123013310-2332123212130022-0023333010011221-1230113002222301-1130100022112333-2321100033121332-2322203201132102-0001311332202103): complete subsection reference.

<a id="canonical-0300120212333211-0322121311110000-3322010323221200-3220311212123131-1313221230123311-3313133102000123-3033132130131003-0200211020033230"></a>

<a id="canonical-1331122121213020-1322012121301220-3212030032321300-2323122230223232-0120200002210122-1330131200311032-2101101102001120-2301211120101303"></a>

## name property — Property reference / 320323313321 / 9

Type: `"string"`. Required.

Name of the API Definition. Must be unique within the namespace.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3103332303330101-3311300211110011-1123001320030323-2303321103211100-3013310233222003-2301033122023021-0003313113330210-2223233033101102"></a>

<a id="canonical-0232021000021013-0003320321200113-1120321112313213-1321013221303133-1311020100102111-2100320112320033-2121300313020003-2330002320323301"></a>

## namespace property — Property reference / 320323313321 / 10

Type: `"string"`. Required.

Namespace where the API Definition is created.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [non_api_endpoints](resources--api_definition--reference--group-001.md#canonical-3013111320133022-1120132002103332-1322121200101212-1332122003230001-3132113222020231-0022312331223301-2130330000300003-1313003013011323): complete subsection reference.

- [strict_schema_origin](resources--api_definition--reference--group-001.md#canonical-2310302013203313-0231100001203113-1203121222003301-2311001121000010-0000132332111022-1101302100003131-0301013302332203-1033132331200113): complete subsection reference.

<a id="canonical-3131321103330221-2001032231110232-3201213033133312-1102100231022003-1120100300111230-0013300211000233-1233000213223113-0000212321300201"></a>

<a id="canonical-3211211221131323-3033031121033100-3330320011231031-3330031122330230-2211232133032001-1320203213212322-2032311103200201-3231023221001323"></a>

## swagger_specs property — Property reference / 320323313321 / 11

Type: `["list", "string"]`. Optional, Computed.

URLs of versioned OpenAPI files uploaded through Web App &amp; API Protection &gt; Files &gt;
Swagger Files. The 512-byte item limit is a URL-length limit; inline string:/// OpenAPI content is
rejected and does not create an API-definition object. Defaults to \`\[\]\`. Server applies default
when omitted.

Upstream description:

URLs of versioned OpenAPI files uploaded through Web App &amp; API Protection &gt; Files &gt;
Swagger Files. The 512-byte item limit is a URL-length limit; inline string:/// OpenAPI content is
rejected and does not create an API-definition object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "512",
    "ves.io.schema.rules.repeated.items.string.pattern": "/api/object_store/namespaces/([a-z]([-a-z0-9]*[a-z0-9])?)/stored_objects/swagger/([a-z]([-a-z0-9]*[a-z0-9])?)/(v|V)[0-9]+(-[0-9]{2}){3}$",
    "ves.io.schema.rules.repeated.max_items": "20",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [timeouts](resources--api_definition--reference--group-001.md#canonical-0012033303210001-2011322101331022-1222132032212313-2010223123333221-2221230230221111-0103330330311200-2001332103301321-0102023222320312): complete subsection reference.

<a id="canonical-0230011030021013-3310123000213133-3120123010113122-3320100111002020-0322301011210333-0113213312032111-1332130023112122-3010212021233033"></a>

## All schema paths — Property reference / 320323313321 / 12

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--api_definition--reference--group-001.md#canonical-3102233200013303-1111100131032012-2100033110201321-0201100101011301-1103320212203011-2323221333010122-0313020330323322-1111232302213313) |
| `api_inventory_exclusion_list` | [api_inventory_exclusion_list](resources--api_definition--reference--group-001.md#canonical-1132033332122122-1312011331033211-3133230330230222-0231002123101322-1031210302022121-1121023311122023-3301203020023231-2222233023200022) |
| `api_inventory_exclusion_list.method` | [api_inventory_exclusion_list.method](resources--api_definition--reference--group-001.md#canonical-0233221231212102-3331220232303031-2230300122222120-3323200031300021-0123103311023033-1110203211310130-1013111233232330-0031221113111022) |
| `api_inventory_exclusion_list.path` | [api_inventory_exclusion_list.path](resources--api_definition--reference--group-001.md#canonical-3113213002132003-1203300301212013-3020113230133121-1312232222233323-0222323201213301-1031332221211113-2302211032013010-2303033131131300) |
| `api_inventory_inclusion_list` | [api_inventory_inclusion_list](resources--api_definition--reference--group-001.md#canonical-2222132221132332-0311100003112202-1123111310121210-0301030231301112-2020231300103122-1002131331300111-0303101211133103-1131320230121131) |
| `api_inventory_inclusion_list.method` | [api_inventory_inclusion_list.method](resources--api_definition--reference--group-001.md#canonical-1333310323321131-1321223201000033-3302213113121022-0112220220231032-1013001113303310-3123022331033232-3300211120302301-0230233131321332) |
| `api_inventory_inclusion_list.path` | [api_inventory_inclusion_list.path](resources--api_definition--reference--group-001.md#canonical-3212220032230133-0201202010010033-0311313223201133-0131003023330332-0123121122202010-0301123220221310-1031310102010210-1001311322302303) |
| `description` | [description](resources--api_definition--reference--group-001.md#canonical-1112202122030001-0110002103230202-3031112222220330-0113202333110003-3220220101100121-2113232130000023-0323102132322200-3032303332101313) |
| `disable` | [disable](resources--api_definition--reference--group-001.md#canonical-1013000220232233-2132303201100201-0010223232311322-1102231133310223-3033000011131121-0232333111000003-3311310033013133-2221013231112031) |
| `id` | [ID](resources--api_definition--reference--group-001.md#canonical-0113223200321130-1021103333203201-3031113103200110-3033031020313003-3222303000302122-0020112333300303-3331023102220132-0033011221203300) |
| `labels` | [labels](resources--api_definition--reference--group-001.md#canonical-0321010310303020-3323012320322330-0133223201220010-2031023310222323-3230232312222303-0323233231223312-1132021130332122-3032030120030211) |
| `mixed_schema_origin` | [mixed_schema_origin](resources--api_definition--reference--group-001.md#canonical-1033312332033231-0003011113321323-3122112220102231-1131332211300213-3230102132322120-1230330331101030-3322230021110033-1321202322001323) |
| `name` | [name](resources--api_definition--reference--group-001.md#canonical-0300120212333211-0322121311110000-3322010323221200-3220311212123131-1313221230123311-3313133102000123-3033132130131003-0200211020033230) |
| `namespace` | [namespace](resources--api_definition--reference--group-001.md#canonical-3103332303330101-3311300211110011-1123001320030323-2303321103211100-3013310233222003-2301033122023021-0003313113330210-2223233033101102) |
| `non_api_endpoints` | [non_api_endpoints](resources--api_definition--reference--group-001.md#canonical-1331000310221200-0332300220323320-2322200010310301-2300331201002332-3012201122203200-1200201103021122-2033032312000220-3323212330133303) |
| `non_api_endpoints.method` | [non_api_endpoints.method](resources--api_definition--reference--group-001.md#canonical-3103332022003331-3122133023223001-3122301000001110-3303000233021032-0030320303300210-2221121202122332-3130131300000023-2203303233300221) |
| `non_api_endpoints.path` | [non_api_endpoints.path](resources--api_definition--reference--group-001.md#canonical-2123122103132320-2322100331312331-2011121201030002-0013030213231220-2032030131310333-1001210133130312-2103133200330101-0111003033013323) |
| `strict_schema_origin` | [strict_schema_origin](resources--api_definition--reference--group-001.md#canonical-0211100032033120-3223121012022133-1122320010213230-0100210221122222-0013201320333123-3032113101333223-3231200232032303-0313003310111221) |
| `swagger_specs` | [swagger_specs](resources--api_definition--reference--group-001.md#canonical-3131321103330221-2001032231110232-3201213033133312-1102100231022003-1120100300111230-0013300211000233-1233000213223113-0000212321300201) |
| `timeouts` | [timeouts](resources--api_definition--reference--group-001.md#canonical-0033101020111030-0031231013221322-2323132123022022-0231032232021102-0022112333131220-0231020200233312-1321222013103100-3200323021233022) |
| `timeouts.create` | [timeouts.create](resources--api_definition--reference--group-001.md#canonical-0130223311323011-0332021302123132-2102112110020100-0122301031030300-2011000010132011-1101200311110012-1012233310102222-1321002231132010) |
| `timeouts.delete` | [timeouts.delete](resources--api_definition--reference--group-001.md#canonical-0200212033203121-2330200232332022-0123020230211022-0011123212302103-1202231022303033-2113000312311102-1223113302110120-3311103130332322) |
| `timeouts.read` | [timeouts.read](resources--api_definition--reference--group-001.md#canonical-0331230122300222-0210303122302201-1010211200212301-0013112323300120-2000123200302011-0033212211133012-0000103212213111-3111303323311210) |
| `timeouts.update` | [timeouts.update](resources--api_definition--reference--group-001.md#canonical-3101203321021231-0321100113332001-1130011211323002-0301222332213231-1003120301033211-1220323311122201-0133031122132101-1301110222121332) |

<a id="canonical-2002303111121211-0312202312320322-1310210122311332-0121333232221312-2112301013110030-0032133010011032-3330203321231201-1210310302101032"></a>

## Next pages — Property reference / 320323313321 / 13

- [api_inventory_exclusion_list](resources--api_definition--reference--group-001.md#canonical-0322111332303200-0102230312330331-2230231331102322-0111003030300213-3103033311100113-2312223222013222-0332133333003331-1230011112023110)
- [api_inventory_inclusion_list](resources--api_definition--reference--group-001.md#canonical-0323203300310302-0203231011212201-3122301133231102-1200330320231103-3013021220022111-2110033120132001-0322200121032302-1301222111101101)
- [mixed_schema_origin](resources--api_definition--reference--group-001.md#canonical-2013020123013310-2332123212130022-0023333010011221-1230113002222301-1130100022112333-2321100033121332-2322203201132102-0001311332202103)
- [non_api_endpoints](resources--api_definition--reference--group-001.md#canonical-3013111320133022-1120132002103332-1322121200101212-1332122003230001-3132113222020231-0022312331223301-2130330000300003-1313003013011323)
- [strict_schema_origin](resources--api_definition--reference--group-001.md#canonical-2310302013203313-0231100001203113-1203121222003301-2311001121000010-0000132332111022-1101302100003131-0301013302332203-1033132331200113)
- [timeouts](resources--api_definition--reference--group-001.md#canonical-0012033303210001-2011322101331022-1222132032212313-2010223123333221-2221230230221111-0103330330311200-2001332103301321-0102023222320312)
- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)

<a id="canonical-0322111332303200-0102230312330331-2230231331102322-0111003030300213-3103033311100113-2312223222013222-0332133333003331-1230011112023110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021321111222110-0300000211312110-0330322222032303-0102120312120110-3122020300231222-2011103033131310-0323032113321130-1222030231200001"></a>

## api_inventory_exclusion_list — api_inventory_exclusion_list / 102101330312 / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)
- [Property reference](resources--api_definition--reference--group-001.md#canonical-1013013311212220-2100012121010012-2012032113201303-3011001133121233-0013201202222002-0113112031031113-2000220121302133-2321231322222211)
- api_inventory_exclusion_list

<a id="canonical-1132033332122122-1312011331033211-3133230330230222-0231002123101322-1031210302022121-1121023311122023-3301203020023231-2222233023200022"></a>

Type: `"object"`. list nested block, Optional.

List of API Endpoints excluded from the API Inventory. Defaults to \`\[\]\`. Server applies default
when omitted.

Upstream description:

List of API Endpoints excluded from the API Inventory.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("path")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
api_inventory_exclusion_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2022320012230210-0123020102201330-1223023012032233-2300010101010130-3320020302111202-2320113300321012-0022020113123031-1230320232212023"></a>

## Direct properties — api_inventory_exclusion_list / 102101330312 / 3

<a id="canonical-0233221231212102-3331220232303031-2230300122222120-3323200031300021-0123103311023033-1110203211310130-1013111233232330-0031221113111022"></a>

<a id="canonical-0222003121203013-3000221131001100-1103201213012212-3123310300310210-3230130233101123-2031310320131011-3031021021201202-2302023030321313"></a>

## method property — api_inventory_exclusion_list / 102101330312 / 4

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

<a id="canonical-3113213002132003-1203300301212013-3020113230133121-1312232222233323-0222323201213301-1031332221211113-2302211032013010-2303033131131300"></a>

<a id="canonical-2010010301113212-1301111020113320-3023113302112220-1203030022022200-1031210120201103-3011022211012300-1323032300230300-2333313332003013"></a>

## path property — api_inventory_exclusion_list / 102101330312 / 5

Type: `"string"`. Optional.

Endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC 3986
and may have parameters according to OpenAPI specification.

Upstream description:

An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC
3986 and may have parameters according to OpenAPI specification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-0223213321113022-0312331202231020-2311231312120110-1323211111000022-2023323002203001-3203232223232133-2302310333123102-2030331032221010"></a>

## Next pages — api_inventory_exclusion_list / 102101330312 / 6

- [Property reference](resources--api_definition--reference--group-001.md#canonical-1013013311212220-2100012121010012-2012032113201303-3011001133121233-0013201202222002-0113112031031113-2000220121302133-2321231322222211)
- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)

<a id="canonical-0323203300310302-0203231011212201-3122301133231102-1200330320231103-3013021220022111-2110033120132001-0322200121032302-1301222111101101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001011013001001-1132131231130131-0231220213300022-0023001332300201-3020023211121332-1121023203322010-1123330312133231-3321130121010211"></a>

## api_inventory_inclusion_list — api_inventory_inclusion_list / 332012110212 / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)
- [Property reference](resources--api_definition--reference--group-001.md#canonical-1013013311212220-2100012121010012-2012032113201303-3011001133121233-0013201202222002-0113112031031113-2000220121302133-2321231322222211)
- api_inventory_inclusion_list

<a id="canonical-2222132221132332-0311100003112202-1123111310121210-0301030231301112-2020231300103122-1002131331300111-0303101211133103-1131320230121131"></a>

Type: `"object"`. list nested block, Optional.

List of API Endpoints included in the API Inventory. Typically, discovered API endpoints are added
to the API Inventory using this list. Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

List of API Endpoints included in the API Inventory. Typically, discovered API endpoints are added
to the API Inventory using this list.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("path")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
api_inventory_inclusion_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211300110133023-1032122013210333-3333122311002232-0123222031022211-0232332330332113-1130333331230011-1032310130231222-0303212321013013"></a>

## Direct properties — api_inventory_inclusion_list / 332012110212 / 3

<a id="canonical-1333310323321131-1321223201000033-3302213113121022-0112220220231032-1013001113303310-3123022331033232-3300211120302301-0230233131321332"></a>

<a id="canonical-2132122233323011-0012110021030101-3221211313003023-3033022120011330-0031020323221003-1332011121310131-1301113220020003-3121011333033310"></a>

## method property — api_inventory_inclusion_list / 332012110212 / 4

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

<a id="canonical-3212220032230133-0201202010010033-0311313223201133-0131003023330332-0123121122202010-0301123220221310-1031310102010210-1001311322302303"></a>

<a id="canonical-0322122321222022-1013001103202301-0312003133303102-0303022310213023-0333123303302202-2110201103223132-0122112220331221-3020010312002213"></a>

## path property — api_inventory_inclusion_list / 332012110212 / 5

Type: `"string"`. Optional.

Endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC 3986
and may have parameters according to OpenAPI specification.

Upstream description:

An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC
3986 and may have parameters according to OpenAPI specification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-3300032130201333-2102321021101112-2121013122002132-0302011110223120-3201201313100211-2033111203022332-0313020032132201-3211013321213310"></a>

## Next pages — api_inventory_inclusion_list / 332012110212 / 6

- [Property reference](resources--api_definition--reference--group-001.md#canonical-1013013311212220-2100012121010012-2012032113201303-3011001133121233-0013201202222002-0113112031031113-2000220121302133-2321231322222211)
- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)

<a id="canonical-2013020123013310-2332123212130022-0023333010011221-1230113002222301-1130100022112333-2321100033121332-2322203201132102-0001311332202103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332201211122310-1032120030211020-1011010022011121-0301101231313222-2123003013031300-0313211233123000-1320121201102212-3102132303233222"></a>

## mixed_schema_origin — mixed_schema_origin / 310123100301 / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)
- [Property reference](resources--api_definition--reference--group-001.md#canonical-1013013311212220-2100012121010012-2012032113201303-3011001133121233-0013201202222002-0113112031031113-2000220121302133-2321231322222211)
- mixed_schema_origin

<a id="canonical-1033312332033231-0003011113321323-3122112220102231-1131332211300213-3230102132322120-1230330331101030-3322230021110033-1321202322001323"></a>

Type: `["object", {}]`. Optional.

\[OneOf: mixed\_schema\_origin, strict\_schema\_origin\] Configuration parameter for mixed schema
origin.

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

- [mixed_schema_origin](resources--api_definition--reference--group-001.md#canonical-1033312332033231-0003011113321323-3122112220102231-1131332211300213-3230102132322120-1230330331101030-3322230021110033-1321202322001323)
- [strict_schema_origin](resources--api_definition--reference--group-001.md#canonical-0211100032033120-3223121012022133-1122320010213230-0100210221122222-0013201320333123-3032113101333223-3231200232032303-0313003310111221)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
mixed_schema_origin = {}
```

<a id="canonical-0201132213111301-0012131010131311-2131003311133130-0131022123001020-0022302122303110-1203011030310312-2002301222222232-1222223322103030"></a>

## Direct properties — mixed_schema_origin / 310123100301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222210220323330-2110322102323221-0320123223313130-2313220132220120-3231233313222323-3101103021001213-1231201322212131-0232301010232112"></a>

## Next pages — mixed_schema_origin / 310123100301 / 4

- [Property reference](resources--api_definition--reference--group-001.md#canonical-1013013311212220-2100012121010012-2012032113201303-3011001133121233-0013201202222002-0113112031031113-2000220121302133-2321231322222211)
- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)

<a id="canonical-3013111320133022-1120132002103332-1322121200101212-1332122003230001-3132113222020231-0022312331223301-2130330000300003-1313003013011323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231232201211112-2110302003320331-1001023202323113-2023202211221031-3100222230321122-3100323023101111-0033123211001112-2310301110221331"></a>

## non_api_endpoints — non_api_endpoints / 133012011222 / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)
- [Property reference](resources--api_definition--reference--group-001.md#canonical-1013013311212220-2100012121010012-2012032113201303-3011001133121233-0013201202222002-0113112031031113-2000220121302133-2321231322222211)
- non_api_endpoints

<a id="canonical-1331000310221200-0332300220323320-2322200010310301-2300331201002332-3012201122203200-1200201103021122-2033032312000220-3323212330133303"></a>

Type: `"object"`. list nested block, Optional.

API Discovery Exclusion List. List of Non-API Endpoints. Defaults to \`\[\]\`. Server applies
default when omitted.

Upstream description:

List of Non-API Endpoints.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("path")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
non_api_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-1331320312000310-1112321130202111-1212333323213312-1200300013002000-2002322112330303-2103033030121233-1002223312133230-1212303203211003"></a>

## Direct properties — non_api_endpoints / 133012011222 / 3

<a id="canonical-3103332022003331-3122133023223001-3122301000001110-3303000233021032-0030320303300210-2221121202122332-3130131300000023-2203303233300221"></a>

<a id="canonical-3133013032203133-1202320231133302-1203021203313112-2001312220012210-0031131231301033-2231011221230311-3310132013202301-1002133223033311"></a>

## method property — non_api_endpoints / 133012011222 / 4

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

<a id="canonical-2123122103132320-2322100331312331-2011121201030002-0013030213231220-2032030131310333-1001210133130312-2103133200330101-0111003033013323"></a>

<a id="canonical-1202131120211100-0211023302112012-1131321110301230-3301121222011323-1030302211032233-3020021122331332-2323021230323101-1002013332210320"></a>

## path property — non_api_endpoints / 133012011222 / 5

Type: `"string"`. Optional.

Endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC 3986
and may have parameters according to OpenAPI specification.

Upstream description:

An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC
3986 and may have parameters according to OpenAPI specification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-2001020333203202-3301112111132003-1102203213123021-1113311310022130-1031320123321103-0202233000133300-1000132310213211-0022002122202212"></a>

## Next pages — non_api_endpoints / 133012011222 / 6

- [Property reference](resources--api_definition--reference--group-001.md#canonical-1013013311212220-2100012121010012-2012032113201303-3011001133121233-0013201202222002-0113112031031113-2000220121302133-2321231322222211)
- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)

<a id="canonical-2310302013203313-0231100001203113-1203121222003301-2311001121000010-0000132332111022-1101302100003131-0301013302332203-1033132331200113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102322020011333-1210123332030012-1300300330210103-3333000123020132-0222110230122133-0013133110313233-1323220022203333-1201031020202020"></a>

## strict_schema_origin — strict_schema_origin / 212212110031 / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)
- [Property reference](resources--api_definition--reference--group-001.md#canonical-1013013311212220-2100012121010012-2012032113201303-3011001133121233-0013201202222002-0113112031031113-2000220121302133-2321231322222211)
- strict_schema_origin

<a id="canonical-0211100032033120-3223121012022133-1122320010213230-0100210221122222-0013201320333123-3032113101333223-3231200232032303-0313003310111221"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for strict schema origin. Defaults to \`map\[\]\`. Server applies default
when omitted.

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
strict_schema_origin = {}
```

<a id="canonical-0210012222201003-0121023320303232-0031031032230000-0212120223221102-1110113310110032-2111100022202120-3002100312300333-3120021002011331"></a>

## Direct properties — strict_schema_origin / 212212110031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221222130203100-1021220201031310-0100311323322111-3212313021301112-3121300202133301-0211033323120101-1223101232110233-3031201113301221"></a>

## Next pages — strict_schema_origin / 212212110031 / 4

- [Property reference](resources--api_definition--reference--group-001.md#canonical-1013013311212220-2100012121010012-2012032113201303-3011001133121233-0013201202222002-0113112031031113-2000220121302133-2321231322222211)
- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)

<a id="canonical-0012033303210001-2011322101331022-1222132032212313-2010223123333221-2221230230221111-0103330330311200-2001332103301321-0102023222320312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112231222013021-3122030211130120-3103222313130123-1020320322231132-1201133223122111-0123121033131112-3111203120302021-3220032101232000"></a>

## timeouts — timeouts / 222323333002 / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)
- [Property reference](resources--api_definition--reference--group-001.md#canonical-1013013311212220-2100012121010012-2012032113201303-3011001133121233-0013201202222002-0113112031031113-2000220121302133-2321231322222211)
- timeouts

<a id="canonical-0033101020111030-0031231013221322-2323132123022022-0231032232021102-0022112333131220-0231020200233312-1321222013103100-3200323021233022"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133122032223201-1220113211122013-1232311022321122-0331021022222033-1112011103232222-3131213132012001-0030120121303212-3032330131020131"></a>

## Direct properties — timeouts / 222323333002 / 3

<a id="canonical-0130223311323011-0332021302123132-2102112110020100-0122301031030300-2011000010132011-1101200311110012-1012233310102222-1321002231132010"></a>

<a id="canonical-3311130233210310-2101313232013223-1131321002230323-0132213102032213-2022332030110320-3200222213223123-2302232003133000-3211220001200130"></a>

## create property — timeouts / 222323333002 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0200212033203121-2330200232332022-0123020230211022-0011123212302103-1202231022303033-2113000312311102-1223113302110120-3311103130332322"></a>

<a id="canonical-3110022120320001-1011202132031320-0311313101101311-2011312331310220-1223031013101210-3202110023110302-3121100312220320-3030313030312232"></a>

## delete property — timeouts / 222323333002 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0331230122300222-0210303122302201-1010211200212301-0013112323300120-2000123200302011-0033212211133012-0000103212213111-3111303323311210"></a>

<a id="canonical-0221212020021323-2211303210023031-2301103231033321-0321120312032110-2212000121322300-3111131331130301-0002020323020322-2222210000023111"></a>

## read property — timeouts / 222323333002 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3101203321021231-0321100113332001-1130011211323002-0301222332213231-1003120301033211-1220323311122201-0133031122132101-1301110222121332"></a>

<a id="canonical-0033312230303210-2231233321110001-1121112300002330-2102010101011200-1310033132331322-0310200132133022-0303330203123110-2230333212311232"></a>

## update property — timeouts / 222323333002 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2103330212203303-0320002021220312-3012133320021021-1211300110131102-2010302200201230-2210113000022313-1133311100033210-3100003201231321"></a>

## Next pages — timeouts / 222323333002 / 8

- [Property reference](resources--api_definition--reference--group-001.md#canonical-1013013311212220-2100012121010012-2012032113201303-3011001133121233-0013201202222002-0113112031031113-2000220121302133-2321231322222211)
- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)
