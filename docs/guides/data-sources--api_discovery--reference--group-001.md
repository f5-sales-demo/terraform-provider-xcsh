---
page_title: "xcsh_api_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_discovery reference."
---

# xcsh_api_discovery reference

<a id="canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132321030333032-0010210232222330-0211320201230220-3321121210223333-1303023330301123-1302131211123131-3330100221202320-0221230213303322"></a>

## Property reference — Property reference / 300303222322 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- Property reference

<a id="canonical-2303122201020110-3222323301200021-1220133121220320-0132223102222203-2311330223321100-0231131230001002-1301330313001121-0222113023201303"></a>

## Direct properties — Property reference / 300303222322 / 3

<a id="canonical-2111110031210221-3102021321301101-2031002333312332-0120033101201322-0033302010030310-2333111130311002-2211330120001330-2022011003000023"></a>

<a id="canonical-3132312111112203-3123003001131230-1230200220211311-3223332202100033-0103132212102200-1022010332332130-0030101010023030-3200020221122222"></a>

## annotations property — Property reference / 300303222322 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
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

- [custom_auth_types](data-sources--api_discovery--reference--group-001.md#canonical-0030011001132310-0130023101132022-3322210030220302-0321012110320203-0313220020332330-0100101001100333-1013101323033200-3133301202233011): complete subsection reference.

<a id="canonical-0212132202101303-0103030121021322-1133023020322232-1113131223013121-2113020333231213-0202122122211031-3323331222032300-2313212033310302"></a>

<a id="canonical-1120231233300031-2323211311122100-1322212121030021-3232102320213221-3203011023112331-3032133020130201-3112101020120323-0201010032132002"></a>

## description property — Property reference / 300303222322 / 5

Type: `"string"`. Computed.

Description of the APIDiscovery.

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

<a id="canonical-3230223321132011-0133301210302111-0202220331310332-2300322222322133-1233300103122111-3133313303100331-0323230303001203-3301112103011122"></a>

<a id="canonical-2232210200122001-3303321311030311-2121133332012123-0233332300331323-2111022102231201-0222211003323122-0001122211231031-1012220223100233"></a>

## ID property — Property reference / 300303222322 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1131030220020301-3131203200203333-0133313032121223-0001222203001223-1301222302033210-1230110211101031-3111323323232302-1110103022000111"></a>

<a id="canonical-3110322032133100-1033331322210222-1223303302330331-3301221120022322-0133113033023102-3110002013223211-2113132012131322-3301330313233022"></a>

## labels property — Property reference / 300303222322 / 7

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

<a id="canonical-2210222232220122-3312312232201231-0001310123112022-3303001221101233-0001311203311303-2102303233000021-0330030313101232-1102311132103011"></a>

<a id="canonical-0030110120033213-1122030202221331-1211022211102010-0333300030333303-3201113121130010-0202320003311221-2003130101000200-3333121110200320"></a>

## name property — Property reference / 300303222322 / 8

Type: `"string"`. Required.

Name of the APIDiscovery.

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

<a id="canonical-0000033223301033-0010020211023132-0132202200333201-1020320302002023-3233233101323111-1003330210220220-1303202133113113-0030223010032101"></a>

<a id="canonical-2102232131001300-1120311132201203-0221223123223013-1031121300113013-1230032001031230-1303201022001230-2320131202002102-0232233132021301"></a>

## namespace property — Property reference / 300303222322 / 9

Type: `"string"`. Required.

Namespace where the APIDiscovery exists.

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

- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133): complete subsection reference.

<a id="canonical-0200303310301101-1302230023023321-0200233323001013-2020112100231121-3211210231023313-0203330322113112-1122123220302200-1303223033202332"></a>

## All schema paths — Property reference / 300303222322 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--api_discovery--reference--group-001.md#canonical-2111110031210221-3102021321301101-2031002333312332-0120033101201322-0033302010030310-2333111130311002-2211330120001330-2022011003000023) |
| `custom_auth_types` | [custom_auth_types](data-sources--api_discovery--reference--group-001.md#canonical-0003003120201223-2121123011122303-0220322321333021-0313033221122332-0122323312311103-2123123232021321-1302210111132012-3201031310201131) |
| `custom_auth_types.parameter_name` | [custom_auth_types.parameter_name](data-sources--api_discovery--reference--group-001.md#canonical-3213023002223022-1030032101210302-0210331000103220-1023200102220012-2001322201233021-2111312131330100-3130113332212032-0030211102131123) |
| `custom_auth_types.parameter_type` | [custom_auth_types.parameter_type](data-sources--api_discovery--reference--group-001.md#canonical-3223321203002321-0220132123321011-1211133122202023-1033321100233001-2012312132103201-2011312133202202-1333212332131231-3311303203330100) |
| `description` | [description](data-sources--api_discovery--reference--group-001.md#canonical-0212132202101303-0103030121021322-1133023020322232-1113131223013121-2113020333231213-0202122122211031-3323331222032300-2313212033310302) |
| `id` | [ID](data-sources--api_discovery--reference--group-001.md#canonical-3230223321132011-0133301210302111-0202220331310332-2300322222322133-1233300103122111-3133313303100331-0323230303001203-3301112103011122) |
| `labels` | [labels](data-sources--api_discovery--reference--group-001.md#canonical-1131030220020301-3131203200203333-0133313032121223-0001222203001223-1301222302033210-1230110211101031-3111323323232302-1110103022000111) |
| `name` | [name](data-sources--api_discovery--reference--group-001.md#canonical-2210222232220122-3312312232201231-0001310123112022-3303001221101233-0001311203311303-2102303233000021-0330030313101232-1102311132103011) |
| `namespace` | [namespace](data-sources--api_discovery--reference--group-001.md#canonical-0000033223301033-0010020211023132-0132202200333201-1020320302002023-3233233101323111-1003330210220220-1303202133113113-0030223010032101) |
| `user_defined_api_discovery_policy` | [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-2201312222331013-1121310232323033-1032102202122303-0030023020333133-2311230211323231-0103101133020030-1031200303003231-3213311233122331) |
| `user_defined_api_discovery_policy.discovery_rules` | [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-3032233201323030-0321121222301000-2100202111122211-1032133213122003-1102032210210223-1200311021332231-2132012223130222-0330100112330221) |
| `user_defined_api_discovery_policy.discovery_rules.labels` | [user_defined_api_discovery_policy.discovery_rules.labels](data-sources--api_discovery--reference--group-001.md#canonical-3332123100112311-1031222111211020-2203030012132000-3123012212001122-0213311112233112-0003112100230200-1022033332210220-0113010132320102) |
| `user_defined_api_discovery_policy.discovery_rules.metadata` | [user_defined_api_discovery_policy.discovery_rules.metadata](data-sources--api_discovery--reference--group-001.md#canonical-1021303133000301-0033123023012230-1022112020220330-3003012133111102-2210101030320221-1223100231033030-2122222021232102-3132330322230112) |
| `user_defined_api_discovery_policy.discovery_rules.metadata.description_spec` | [user_defined_api_discovery_policy.discovery_rules.metadata.description_spec](data-sources--api_discovery--reference--group-001.md#canonical-3030222112000331-0123202302312123-2223312130110210-2302002121032213-2201002333100002-3233020223310313-0031212331201123-3312011032312301) |
| `user_defined_api_discovery_policy.discovery_rules.metadata.name` | [user_defined_api_discovery_policy.discovery_rules.metadata.name](data-sources--api_discovery--reference--group-001.md#canonical-1113231211313120-2103330122001000-3110301313222211-1102331321201331-1333323211300301-1031231002322122-0131231130101311-2110301020120012) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties` | [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-2321313133101321-3300031020120121-1102310303232100-2301020012032330-3211213000031203-1002231200100300-0011220322012330-1102002313203002) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](data-sources--api_discovery--reference--group-001.md#canonical-1111312311233102-2220301213012002-0200133123130221-1220322100013333-2311101312020311-2022110021222202-1103300010320231-2220032100212023) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive](data-sources--api_discovery--reference--group-001.md#canonical-1021001020312100-0110311023210113-3333221230130300-2110000200200213-3322211310100113-1102303022301101-0100211213030323-2331312232201003) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore](data-sources--api_discovery--reference--group-001.md#canonical-3102301010233233-0030223010033000-1112311323033300-0131031112011223-2002311313110113-2301123123301303-3201111023222031-0133321102331333) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria](data-sources--api_discovery--reference--group-001.md#canonical-3313113220133202-0232213331310201-2001020021022210-0133003100313221-1332200032211310-0301221011223001-2100303232023101-1023210220310222) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.field_name` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.field_name](data-sources--api_discovery--reference--group-001.md#canonical-1302231133122212-1323120020033321-2212011202213000-0311030222003000-2133111000120310-1001133130120031-3032310100002103-2303122020032022) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.location` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.location](data-sources--api_discovery--reference--group-001.md#canonical-2101123011102302-1321202200101103-3120132312320320-0002021032020303-3101023311323331-1002001132130100-2130203012001021-3110122113011132) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.match_type` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.match_type](data-sources--api_discovery--reference--group-001.md#canonical-3303202110023223-2103213321103011-2313012022220333-1320113310121032-0102002222231220-3010213120121021-2012003330033320-0030313233012222) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.value` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.value](data-sources--api_discovery--reference--group-001.md#canonical-2211022132000113-3202221201121300-3202323231221001-3003020212131201-1021323010113013-0103122121020311-2120001112221133-0222130231301013) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion](data-sources--api_discovery--reference--group-001.md#canonical-3330232230332331-2302212100031023-3223001032102231-2001213323122132-3333221212202001-0111313212221131-3303131131311121-3003320011301022) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.pattern` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.pattern](data-sources--api_discovery--reference--group-001.md#canonical-3321030031233303-0121131321322313-3012321123121020-0130202300231313-0011222310131230-0200213320210331-1312202100102331-2301010201032102) |
| `user_defined_api_discovery_policy.exclusive` | [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--reference--group-001.md#canonical-2312030003110120-1323331300320032-3021001331200032-0111111123100013-1033222102020112-0133322221132121-2312223222100313-0301233011012110) |
| `user_defined_api_discovery_policy.exclusive.archive` | [user_defined_api_discovery_policy.exclusive.archive](data-sources--api_discovery--reference--group-001.md#canonical-3121113303223022-3300321111101303-3120301220322123-2101122013322200-0210012023030320-2121012010222301-3201200311200131-2222032132313202) |
| `user_defined_api_discovery_policy.exclusive.ignore` | [user_defined_api_discovery_policy.exclusive.ignore](data-sources--api_discovery--reference--group-001.md#canonical-2333011132001130-2333101122133330-1310323303311222-1130231330022031-1130323021130101-0233103323233122-2131112221032310-2330031301302121) |
| `user_defined_api_discovery_policy.inclusive` | [user_defined_api_discovery_policy.inclusive](data-sources--api_discovery--reference--group-001.md#canonical-1320222130100200-3012303112033213-2020321300211202-3200230113130222-1230101221130302-1231011020332323-3131122122223300-0000011011113313) |

<a id="canonical-0303220320032131-3022303331211101-3023333233203301-0322221110320111-1222133223111001-2202111223233001-3233121022310131-3030213112231330"></a>

## Next pages — Property reference / 300303222322 / 11

- [custom_auth_types](data-sources--api_discovery--reference--group-001.md#canonical-0030011001132310-0130023101132022-3322210030220302-0321012110320203-0313220020332330-0100101001100333-1013101323033200-3133301202233011)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-0030011001132310-0130023101132022-3322210030220302-0321012110320203-0313220020332330-0100101001100333-1013101323033200-3133301202233011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300032210123130-0233112330213312-2002222330022333-2201201132031131-2112001133111333-1332100002132120-0132121312203112-0232100021033020"></a>

## custom_auth_types — custom_auth_types / 102102302322 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- custom_auth_types

<a id="canonical-0003003120201223-2121123011122303-0220322321333021-0313033221122332-0122323312311103-2123123232021321-1302210111132012-3201031310201131"></a>

Type: `"list"`. Computed.

Select your custom authentication types to be detected in the API discovery. Defaults to \`\[\]\`.
Server applies default when omitted.

Upstream description:

Select your custom authentication types to be detected in the API discovery.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "10"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "10"
  }
}
```

<a id="canonical-3200312331002312-3133030021203003-2302333100222330-0220000302210203-3212331310330022-0103323212132111-2101130233330310-2203331003200000"></a>

## Direct properties — custom_auth_types / 102102302322 / 3

<a id="canonical-3213023002223022-1030032101210302-0210331000103220-1023200102220012-2001322201233021-2111312131330100-3130113332212032-0030211102131123"></a>

<a id="canonical-0322332311002313-1011331331100011-2212232331310120-0321223130020310-0211202310231302-3003310131223330-3120123302210003-3132223122131202"></a>

## parameter_name property — custom_auth_types / 102102302322 / 4

Type: `"string"`. Computed.

Parameter Name. The authentication parameter name.

Upstream description:

The authentication parameter name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^[!#$%&'*+\\\\-.^_`|~0-9A-Za-z]+$"
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
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.pattern": "^[!#$%&'*+\\\\-.^_`|~0-9A-Za-z]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.pattern": "^[!#$%&'*+\\\\-.^_`|~0-9A-Za-z]+$"
  }
}
```

<a id="canonical-3223321203002321-0220132123321011-1211133122202023-1033321100233001-2012312132103201-2011312133202202-1333212332131231-3311303203330100"></a>

<a id="canonical-2030101311232031-0310323330123330-3210113310203011-1102021223121323-2233332130032320-2213113013301311-2123211001233300-0001121202023100"></a>

## parameter_type property — custom_auth_types / 102102302322 / 5

Type: `"string"`. Computed.

\[Enum: QUERY\_PARAMETER|HEADER|COOKIE\] Enumeration for authentication parameter types. Possible
values are \`QUERY\_PARAMETER\`, \`HEADER\`, \`COOKIE\`. Defaults to \`QUERY\_PARAMETER\`.

Upstream description:

Enumeration for authentication parameter types.

Receipt-pinned upstream constraints:

```json
{
  "default": "QUERY_PARAMETER",
  "enum": [
    "QUERY_PARAMETER",
    "HEADER",
    "COOKIE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2222111012213022-2123013130212223-1111231021210321-0113022133020012-3131111320132023-3133201222002220-2301322211200200-3021022133102123"></a>

## Next pages — custom_auth_types / 102102302322 / 6

- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010103010323011-2132210001030201-1302322111103003-1031000122120231-0330313232110312-1020320033133110-1230111300011311-3103332100202132"></a>

## user_defined_api_discovery_policy — user_defined_api_discovery_policy / 021111212222 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- user_defined_api_discovery_policy

<a id="canonical-2201312222331013-1121310232323033-1032102202122303-0030023020333133-2311230211323231-0103101133020030-1031200303003231-3213311233122331"></a>

Type: `"single"`. Computed.

Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be
discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_behavior_choice": "[\"exclusive\",\"inclusive\"]"
}
```

<a id="canonical-0322200331232203-2332011320031103-2110301210002201-0111313330013212-2113010323113210-2033101030023300-3310012032212220-1213111110221212"></a>

## Direct properties — user_defined_api_discovery_policy / 021111212222 / 3

- [discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333): complete subsection reference.

- [exclusive](data-sources--api_discovery--reference--group-001.md#canonical-2300322012120002-0201232221212210-2123211221021011-2200001022101221-2331331221020112-0010131103023101-3003002313313310-3030102012302313): complete subsection reference.

- [inclusive](data-sources--api_discovery--reference--group-001.md#canonical-0330322003032013-1130101013033123-0030320021111200-2010331110331223-1321323331332301-3132233220120201-3023322100000123-1301321323211020): complete subsection reference.

<a id="canonical-0210002312101000-3322232110031232-0012302232030211-3301232033301221-3320200320032031-3112211310113320-0100110322123131-1232113312110223"></a>

## Next pages — user_defined_api_discovery_policy / 021111212222 / 4

- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333)
- [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--reference--group-001.md#canonical-2300322012120002-0201232221212210-2123211221021011-2200001022101221-2331331221020112-0010131103023101-3003002313313310-3030102012302313)
- [user_defined_api_discovery_policy.inclusive](data-sources--api_discovery--reference--group-001.md#canonical-0330322003032013-1130101013033123-0030320021111200-2010331110331223-1321323331332301-3132233220120201-3023322100000123-1301321323211020)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202200332120303-1330122102332103-0333111331330211-3213100122223300-3000123201200021-1001112330311230-2333120301001201-2312120022311021"></a>

## user_defined_api_discovery_policy.discovery_rules — discovery_rules / 031131130123 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- user_defined_api_discovery_policy.discovery_rules

<a id="canonical-3032233201323030-0321121222301000-2100202111122211-1032133213122003-1102032210210223-1200311021332231-2132012223130222-0330100112330221"></a>

Type: `"list"`. Computed.

Define rules to include or exclude endpoints by path, domain, or header. Rules run top to bottom;
unmatched endpoints follow the default action. Defaults to \`\[\]\`. Server applies default when
omitted.

Upstream description:

Define rules to include or exclude endpoints by path, domain, or header. Rules run top to bottom;
unmatched endpoints follow the default action.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3033322312320002-0010033211202030-2022132202130121-0003013333311133-1112223001302031-2303020232113111-1213101230110001-0203331311102020"></a>

## Direct properties — discovery_rules / 031131130123 / 3

- [labels](data-sources--api_discovery--reference--group-001.md#canonical-1131022301210223-3012031113213103-2013203312321000-0222111110021202-3101003110022310-2213112031003113-0213111221132101-2320203002031110): complete subsection reference.

- [metadata](data-sources--api_discovery--reference--group-001.md#canonical-3122231211211030-1132230033000222-0131003213231223-3030113101121210-0201211301123000-0331121111331213-0213330001100032-3310233023300002): complete subsection reference.

- [rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-3323303011223032-3320130223201120-1023021232331211-1310102121212133-1130012211020030-3020120001300323-2232230022103321-0131120302102312): complete subsection reference.

<a id="canonical-2001002310302332-2323120123322132-2010230203110033-2012110301032330-3232003022311110-0223031120301130-3302032132123230-1330331210310230"></a>

## Next pages — discovery_rules / 031131130123 / 4

- [user_defined_api_discovery_policy.discovery_rules.labels](data-sources--api_discovery--reference--group-001.md#canonical-1131022301210223-3012031113213103-2013203312321000-0222111110021202-3101003110022310-2213112031003113-0213111221132101-2320203002031110)
- [user_defined_api_discovery_policy.discovery_rules.metadata](data-sources--api_discovery--reference--group-001.md#canonical-3122231211211030-1132230033000222-0131003213231223-3030113101121210-0201211301123000-0331121111331213-0213330001100032-3310233023300002)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-3323303011223032-3320130223201120-1023021232331211-1310102121212133-1130012211020030-3020120001300323-2232230022103321-0131120302102312)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-1131022301210223-3012031113213103-2013203312321000-0222111110021202-3101003110022310-2213112031003113-0213111221132101-2320203002031110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333112132031012-1120013101022222-0010302010333000-2231101301311023-2332103102110000-0202100212321311-2302233212022302-3201010011000230"></a>

## user_defined_api_discovery_policy.discovery_rules.labels — labels / 032003230110 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333)
- user_defined_api_discovery_policy.discovery_rules.labels

<a id="canonical-3332123100112311-1031222111211020-2203030012132000-3123012212001122-0213311112233112-0003112100230200-1022033332210220-0113010132320102"></a>

Type: `"single"`. Computed.

Map of string keys and values that can be used to organize and categorize the rule.

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

<a id="canonical-1202333023320113-3203113021220132-2130120010320002-2200233020321020-1022203000001101-3002231011223222-3222333131322001-3211011202101202"></a>

## Direct properties — labels / 032003230110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210312110103323-0323300023220131-0003301301010301-1322131103210100-3231221013113331-0002201202033220-0231003213003312-3110201101122121"></a>

## Next pages — labels / 032003230110 / 4

- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-3122231211211030-1132230033000222-0131003213231223-3030113101121210-0201211301123000-0331121111331213-0213330001100032-3310233023300002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013013210200201-0310221120201331-0302133232131321-0120010200222222-0201322310012330-3003331010113330-1220212120111122-1330220102030231"></a>

## user_defined_api_discovery_policy.discovery_rules.metadata — metadata / 033123213133 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333)
- user_defined_api_discovery_policy.discovery_rules.metadata

<a id="canonical-1021303133000301-0033123023012230-1022112020220330-3003012133111102-2210101030320221-1223100231033030-2122222021232102-3132330322230112"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-0100003031223232-2212012300003320-0131133133233321-3203323001102300-0233031310301202-0301111203320113-1101201010203200-2030012203111030"></a>

## Direct properties — metadata / 033123213133 / 3

<a id="canonical-3030222112000331-0123202302312123-2223312130110210-2302002121032213-2201002333100002-3233020223310313-0031212331201123-3312011032312301"></a>

<a id="canonical-0112312300013111-3100032302100310-0003023232300023-0021101201322030-0123013303112301-1022312121103102-3332030122223233-3033200123310222"></a>

## description_spec property — metadata / 033123213133 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1113231211313120-2103330122001000-3110301313222211-1102331321201331-1333323211300301-1031231002322122-0131231130101311-2110301020120012"></a>

<a id="canonical-2311200023112002-2200011133131212-3030322012011021-1332013010130032-1200003220013032-3010210003000333-0000223212002021-0100112210020323"></a>

## name property — metadata / 033123213133 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0322213100112120-1313033323200020-0201101333032002-0333120212311013-1113201002123030-0201011111210221-0130230232301212-2130310212011331"></a>

## Next pages — metadata / 033123213133 / 6

- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-3323303011223032-3320130223201120-1023021232331211-1310102121212133-1130012211020030-3020120001300323-2232230022103321-0131120302102312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310322212011122-0321211302302203-0310130203213102-2330133001321112-0312303021212320-1030030033212012-1000122131231310-1002300022211202"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties — rule_properties / 033031331310 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333)
- user_defined_api_discovery_policy.discovery_rules.rule_properties

<a id="canonical-2321313133101321-3300031020120121-1102310303232100-2301020012032330-3211213000031203-1002231200100300-0011220322012330-1102002313203002"></a>

Type: `"single"`. Computed.

Determines whether matching endpoints are included in API Discovery or excluded.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-criteria": "[\"http_header_criteria\",\"pattern\"]",
  "x-ves-oneof-field-rule_type_choice": "[\"exclusion\",\"inclusion\"]"
}
```

<a id="canonical-0110313123220212-1222233100302231-1223222100122022-2301011122221201-3033311203331202-1222221202101021-0332113332131212-1122332020233222"></a>

## Direct properties — rule_properties / 033031331310 / 3

- [exclusion](data-sources--api_discovery--reference--group-001.md#canonical-3221310322202333-1221230121213123-1112313232331200-2332201220132023-3332123133120032-0311132320213302-0122010033020011-0111010011132221): complete subsection reference.

- [http_header_criteria](data-sources--api_discovery--reference--group-001.md#canonical-0222020233023033-3023331301211103-3221220323212121-2202212123211223-2023200012200111-0210120031121320-2030310021331111-2233003311001121): complete subsection reference.

- [inclusion](data-sources--api_discovery--reference--group-001.md#canonical-2202002321131002-3330102300321213-1021103211211010-0031121133203110-2000320221223110-0001030331112102-3320032102023210-1133122003011331): complete subsection reference.

<a id="canonical-3321030031233303-0121131321322313-3012321123121020-0130202300231313-0011222310131230-0200213320210331-1312202100102331-2301010201032102"></a>

<a id="canonical-0033300012311020-2103030322132300-3033003333323210-0130012313311133-1210020132122322-3323210113232023-1100033300333131-0000113302332220"></a>

## pattern property — rule_properties / 033031331310 / 4

Type: `"string"`. Computed.

Exclusive with \[http\_header\_criteria\] Patterns are matched against the request path to identify
endpoints by path structure, file extension, or version prefix. Endpoints that match this pattern
are affected by the rule.

Upstream description:

Exclusive with \[http\_header\_criteria\] Patterns are matched against the request path to identify
endpoints by path structure, file extension, or version prefix. Endpoints that match this pattern
are affected by the rule.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 512
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-3001023330210101-1330011132022030-3302330103231002-0223300101130302-3312110230203213-3232132223010001-3033121212003220-2132300213233010"></a>

## Next pages — rule_properties / 033031331310 / 5

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](data-sources--api_discovery--reference--group-001.md#canonical-3221310322202333-1221230121213123-1112313232331200-2332201220132023-3332123133120032-0311132320213302-0122010033020011-0111010011132221)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria](data-sources--api_discovery--reference--group-001.md#canonical-0222020233023033-3023331301211103-3221220323212121-2202212123211223-2023200012200111-0210120031121320-2030310021331111-2233003311001121)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion](data-sources--api_discovery--reference--group-001.md#canonical-2202002321131002-3330102300321213-1021103211211010-0031121133203110-2000320221223110-0001030331112102-3320032102023210-1133122003011331)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-3221310322202333-1221230121213123-1112313232331200-2332201220132023-3332123133120032-0311132320213302-0122010033020011-0111010011132221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223313110131121-0302311013230130-1230230200333313-0010310001131102-0111302320230211-2122023220222123-3220333023010122-0200331203223132"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion — exclusion / 320330112332 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-3323303011223032-3320130223201120-1023021232331211-1310102121212133-1130012211020030-3020120001300323-2232230022103321-0131120302102312)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion

<a id="canonical-1111312311233102-2220301213012002-0200133123130221-1220322100013333-2311101312020311-2022110021222202-1103300010320231-2220032100212023"></a>

Type: `"single"`. Computed.

Exclusion Configuration. Configuration for exclusion action.

Upstream description:

Configuration for exclusion action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"archive\",\"ignore\"]"
}
```

<a id="canonical-0210203311332112-0003223210310323-3013033232322113-3022203322003221-1212121303001032-0012222031022012-0212312010332123-1200231010130322"></a>

## Direct properties — exclusion / 320330112332 / 3

- [archive](data-sources--api_discovery--reference--group-001.md#canonical-1211030321223333-3011323023231023-2120113223210331-3000003313222033-0100011110320222-1123033013132313-1302202322203212-3220033203002032): complete subsection reference.

- [ignore](data-sources--api_discovery--reference--group-001.md#canonical-2223202221010110-1020301033200130-3311232122330202-3201032022232230-2232011131230312-1213132103322232-3310112002112113-2130212132121212): complete subsection reference.

<a id="canonical-0311321211010100-1310203203231032-2100133020303332-3122002230303120-1313012322200030-1100212330131020-3200220212131113-3132031120322022"></a>

## Next pages — exclusion / 320330112332 / 4

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive](data-sources--api_discovery--reference--group-001.md#canonical-1211030321223333-3011323023231023-2120113223210331-3000003313222033-0100011110320222-1123033013132313-1302202322203212-3220033203002032)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore](data-sources--api_discovery--reference--group-001.md#canonical-2223202221010110-1020301033200130-3311232122330202-3201032022232230-2232011131230312-1213132103322232-3310112002112113-2130212132121212)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-3323303011223032-3320130223201120-1023021232331211-1310102121212133-1130012211020030-3020120001300323-2232230022103321-0131120302102312)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-1211030321223333-3011323023231023-2120113223210331-3000003313222033-0100011110320222-1123033013132313-1302202322203212-3220033203002032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011320232211000-2100300310113122-2202020302002122-3300001120233010-1223020231121301-0011101130201011-2012012320002323-3231313310312121"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive — archive / 133312222311 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-3323303011223032-3320130223201120-1023021232331211-1310102121212133-1130012211020030-3020120001300323-2232230022103321-0131120302102312)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](data-sources--api_discovery--reference--group-001.md#canonical-3221310322202333-1221230121213123-1112313232331200-2332201220132023-3332123133120032-0311132320213302-0122010033020011-0111010011132221)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive

<a id="canonical-1021001020312100-0110311023210113-3333221230130300-2110000200200213-3322211310100113-1102303022301101-0100211213030323-2331312232201003"></a>

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

<a id="canonical-2001312331013013-2130202201301101-3102233031322220-1310121111320230-2310013112000321-0311203102322331-1223020232321010-3321002131310133"></a>

## Direct properties — archive / 133312222311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013332101120322-3202231333220111-0030211323323030-2300232312133110-2001321200110031-3230220003223002-2223000332211122-2312110221013311"></a>

## Next pages — archive / 133312222311 / 4

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](data-sources--api_discovery--reference--group-001.md#canonical-3221310322202333-1221230121213123-1112313232331200-2332201220132023-3332123133120032-0311132320213302-0122010033020011-0111010011132221)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-2223202221010110-1020301033200130-3311232122330202-3201032022232230-2232011131230312-1213132103322232-3310112002112113-2130212132121212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031121123011310-0130123010221010-0331000001312330-1122201032023012-2120313320333131-1111123223312002-2303011211331033-2121123330333033"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore — ignore / 113103300032 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-3323303011223032-3320130223201120-1023021232331211-1310102121212133-1130012211020030-3020120001300323-2232230022103321-0131120302102312)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](data-sources--api_discovery--reference--group-001.md#canonical-3221310322202333-1221230121213123-1112313232331200-2332201220132023-3332123133120032-0311132320213302-0122010033020011-0111010011132221)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore

<a id="canonical-3102301010233233-0030223010033000-1112311323033300-0131031112011223-2002311313110113-2301123123301303-3201111023222031-0133321102331333"></a>

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

<a id="canonical-0213212132213002-3003130331320112-1203003003023112-0132133323300013-2300203220202023-1101020030203100-3011310300211332-0120201130010300"></a>

## Direct properties — ignore / 113103300032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211203201230230-2201103112113131-0100333030112332-2330310111232222-3231202033030010-3302133301201032-0100201013003222-3122131133101121"></a>

## Next pages — ignore / 113103300032 / 4

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](data-sources--api_discovery--reference--group-001.md#canonical-3221310322202333-1221230121213123-1112313232331200-2332201220132023-3332123133120032-0311132320213302-0122010033020011-0111010011132221)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-0222020233023033-3023331301211103-3221220323212121-2202212123211223-2023200012200111-0210120031121320-2030310021331111-2233003311001121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303123013020311-0030320101120301-1011020120233130-1021222313031321-3331020213222211-0233123131231112-3331212110322001-0133012020333022"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria — http_header_criteria / 300220300133 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-3323303011223032-3320130223201120-1023021232331211-1310102121212133-1130012211020030-3020120001300323-2232230022103321-0131120302102312)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria

<a id="canonical-3313113220133202-0232213331310201-2001020021022210-0133003100313221-1332200032211310-0301221011223001-2100303232023101-1023210220310222"></a>

Type: `"single"`. Computed.

Configuration parameter for http header criteria.

Upstream description:

Criteria for matching HTTP headers.

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

<a id="canonical-0312013131002222-0021223000111211-2222203112201313-3033001320011312-3000032112100323-2032233231121001-3000120112332200-1310101023223323"></a>

## Direct properties — http_header_criteria / 300220300133 / 3

<a id="canonical-1302231133122212-1323120020033321-2212011202213000-0311030222003000-2133111000120310-1001133130120031-3032310100002103-2303122020032022"></a>

<a id="canonical-0133232222102313-3223300133331131-2100201013112000-2102103102231003-2220031010003003-3032011021200133-2331130301003020-3301320102131132"></a>

## field_name property — http_header_criteria / 300220300133 / 4

Type: `"string"`. Computed.

HTTP Header Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2101123011102302-1321202200101103-3120132312320320-0002021032020303-3101023311323331-1002001132130100-2130203012001021-3110122113011132"></a>

<a id="canonical-2322321013111111-3002032111023130-0103330032201003-3110322002321031-1101202223223232-1112202323201232-3103321310332223-2333301330023033"></a>

## location property — http_header_criteria / 300220300133 / 5

Type: `"string"`. Computed.

\[Enum: REQUEST|RESPONSE\] Specifies whether the rule criteria should be evaluated against request
or response Applies the rule to incoming traffic from the client. Applies the rule to outgoing
traffic sent back to the client. Possible values are \`REQUEST\`, \`RESPONSE\`. Defaults to
\`REQUEST\`.

Upstream description:

Specifies whether the rule criteria should be evaluated against request or response

Applies the rule to incoming traffic from the client. Applies the rule to outgoing traffic sent back
to the client.

Receipt-pinned upstream constraints:

```json
{
  "default": "REQUEST",
  "enum": [
    "REQUEST",
    "RESPONSE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3303202110023223-2103213321103011-2313012022220333-1320113310121032-0102002222231220-3010213120121021-2012003330033320-0030313233012222"></a>

<a id="canonical-2020330100010222-0120130030211122-2301213201231012-3221100302012110-2130023320203312-0320233210312200-2122001112110221-0021212333231120"></a>

## match_type property — http_header_criteria / 300220300133 / 6

Type: `"string"`. Computed.

\[Enum: EXACT\_MATCH|SUBSTRING|REGEX\] Specifies how the value should be matched. Possible values
are \`EXACT\_MATCH\`, \`SUBSTRING\`, \`REGEX\`. Defaults to \`EXACT\_MATCH\`.

Upstream description:

Specifies how the value should be matched.

Receipt-pinned upstream constraints:

```json
{
  "default": "EXACT_MATCH",
  "enum": [
    "EXACT_MATCH",
    "SUBSTRING",
    "REGEX"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2211022132000113-3202221201121300-3202323231221001-3003020212131201-1021323010113013-0103122121020311-2120001112221133-0222130231301013"></a>

<a id="canonical-1230321312102111-3002312112323321-0323221011230123-2102202213101312-3321321101022211-3221211102200032-0012303113113020-2221102020223321"></a>

## value property — http_header_criteria / 300220300133 / 7

Type: `"string"`. Computed.

Value. Configuration parameter for value

Upstream description:

Configuration parameter for value

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-1232113011120000-0213100233002221-1211001112033110-1122032313310212-1121122012023132-2101202021011130-1210222031012121-0313231303211223"></a>

## Next pages — http_header_criteria / 300220300133 / 8

- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-3323303011223032-3320130223201120-1023021232331211-1310102121212133-1130012211020030-3020120001300323-2232230022103321-0131120302102312)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-2202002321131002-3330102300321213-1021103211211010-0031121133203110-2000320221223110-0001030331112102-3320032102023210-1133122003011331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010113210103032-0013002022131212-1130013111030123-3021220230322232-3022232111020002-3021101121323112-2331213223013211-0120032001123022"></a>

## user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion — inclusion / 113221113330 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-3323303011223032-3320130223201120-1023021232331211-1310102121212133-1130012211020030-3020120001300323-2232230022103321-0131120302102312)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion

<a id="canonical-3330232230332331-2302212100031023-3223001032102231-2001213323122132-3333221212202001-0111313212221131-3303131131311121-3003320011301022"></a>

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

<a id="canonical-0212313231230223-0313233033011332-3331303023312130-2101130212011202-0110332002112031-2112112030011000-0113020020111122-3123033133000011"></a>

## Direct properties — inclusion / 113221113330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221103101301100-0111210302120302-0232300110032021-0232330223110131-3322033130132211-1021221332100203-2221202120330000-2010131120301113"></a>

## Next pages — inclusion / 113221113330 / 4

- [user_defined_api_discovery_policy.discovery_rules.rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-3323303011223032-3320130223201120-1023021232331211-1310102121212133-1130012211020030-3020120001300323-2232230022103321-0131120302102312)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-2300322012120002-0201232221212210-2123211221021011-2200001022101221-2331331221020112-0010131103023101-3003002313313310-3030102012302313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233023300121121-1130132300201323-0112100111302000-1332002120101101-3333230030003300-0103020113300111-0221030020212333-0303021121301000"></a>

## user_defined_api_discovery_policy.exclusive — exclusive / 003331232121 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- user_defined_api_discovery_policy.exclusive

<a id="canonical-2312030003110120-1323331300320032-3021001331200032-0111111123100013-1033222102020112-0133322221132121-2312223222100313-0301233011012110"></a>

Type: `"single"`. Computed.

Exclusion Configuration. Configuration for exclusion action.

Upstream description:

Configuration for exclusion action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"archive\",\"ignore\"]"
}
```

<a id="canonical-0202121120210130-3133121332301003-3002001221331322-1100120001100201-3232323003123231-2202303133110131-1003332311010203-1330003002223333"></a>

## Direct properties — exclusive / 003331232121 / 3

- [archive](data-sources--api_discovery--reference--group-001.md#canonical-1031032321002013-3131013331022200-3212002300223302-2211032220121221-3213220330313320-3011123203030133-2120021033132130-3213331332220113): complete subsection reference.

- [ignore](data-sources--api_discovery--reference--group-001.md#canonical-0200303112032033-1320213032202301-0222031203313012-0230323332212013-1001123000133311-2121233111121231-2312232131232101-2221323230312221): complete subsection reference.

<a id="canonical-3330323122211120-2232303310000103-2300102130112222-2230030203321132-2120310201220313-3132311021302310-3002103302011101-0110011132033110"></a>

## Next pages — exclusive / 003331232121 / 4

- [user_defined_api_discovery_policy.exclusive.archive](data-sources--api_discovery--reference--group-001.md#canonical-1031032321002013-3131013331022200-3212002300223302-2211032220121221-3213220330313320-3011123203030133-2120021033132130-3213331332220113)
- [user_defined_api_discovery_policy.exclusive.ignore](data-sources--api_discovery--reference--group-001.md#canonical-0200303112032033-1320213032202301-0222031203313012-0230323332212013-1001123000133311-2121233111121231-2312232131232101-2221323230312221)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-1031032321002013-3131013331022200-3212002300223302-2211032220121221-3213220330313320-3011123203030133-2120021033132130-3213331332220113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311220320003021-2123102311212310-3130033200210023-0302322332330111-1212333002231300-2102123302033110-1210330332330330-3032232000300332"></a>

## user_defined_api_discovery_policy.exclusive.archive — archive / 322100023131 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--reference--group-001.md#canonical-2300322012120002-0201232221212210-2123211221021011-2200001022101221-2331331221020112-0010131103023101-3003002313313310-3030102012302313)
- user_defined_api_discovery_policy.exclusive.archive

<a id="canonical-3121113303223022-3300321111101303-3120301220322123-2101122013322200-0210012023030320-2121012010222301-3201200311200131-2222032132313202"></a>

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

<a id="canonical-2120221110231232-2331133230013303-3022300001132322-1101230110022200-1332310122100210-3200211331300132-0023332032313003-1031033023022100"></a>

## Direct properties — archive / 322100023131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133303330331222-3330103123113310-3321130121332321-0100032321310330-0003133113011212-2033000103303112-0010021203333231-1202012132021231"></a>

## Next pages — archive / 322100023131 / 4

- [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--reference--group-001.md#canonical-2300322012120002-0201232221212210-2123211221021011-2200001022101221-2331331221020112-0010131103023101-3003002313313310-3030102012302313)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-0200303112032033-1320213032202301-0222031203313012-0230323332212013-1001123000133311-2121233111121231-2312232131232101-2221323230312221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200322012022123-2312313132231312-1300202330212110-2303322103330123-0210223021221131-2131321301011222-2232012223032232-0321222233030130"></a>

## user_defined_api_discovery_policy.exclusive.ignore — ignore / 321033200213 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--reference--group-001.md#canonical-2300322012120002-0201232221212210-2123211221021011-2200001022101221-2331331221020112-0010131103023101-3003002313313310-3030102012302313)
- user_defined_api_discovery_policy.exclusive.ignore

<a id="canonical-2333011132001130-2333101122133330-1310323303311222-1130231330022031-1130323021130101-0233103323233122-2131112221032310-2330031301302121"></a>

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

<a id="canonical-2031003010203131-2303102101130032-1330322002121231-0033212330023233-1010213221002013-0133330103031130-0032110201321323-1102120220032133"></a>

## Direct properties — ignore / 321033200213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021113032331203-2102211030330110-2213001032201100-0001221210300132-2333330321112333-2331200202132330-2202130303230321-1333311033012001"></a>

## Next pages — ignore / 321033200213 / 4

- [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--reference--group-001.md#canonical-2300322012120002-0201232221212210-2123211221021011-2200001022101221-2331331221020112-0010131103023101-3003002313313310-3030102012302313)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)

<a id="canonical-0330322003032013-1130101013033123-0030320021111200-2010331110331223-1321323331332301-3132233220120201-3023322100000123-1301321323211020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201220201330200-3030322022311133-0232001210111233-0303113111123202-1010122020332131-1131111221102301-3030130202001320-1112311121322022"></a>

## user_defined_api_discovery_policy.inclusive — inclusive / 110001310310 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- user_defined_api_discovery_policy.inclusive

<a id="canonical-1320222130100200-3012303112033213-2020321300211202-3200230113130222-1230101221130302-1231011020332323-3131122122223300-0000011011113313"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-1130003223100302-0230022321111020-0102213311122003-1110201012122110-2002221212112113-0032111233112100-3302032330102121-1030312021211211"></a>

## Direct properties — inclusive / 110001310310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302331201013101-1321000201301212-2010211302233323-3032330300323311-3333222333030200-0111021123031130-2121002130321002-1103021300333201"></a>

## Next pages — inclusive / 110001310310 / 4

- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
