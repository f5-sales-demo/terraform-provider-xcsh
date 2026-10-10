---
page_title: "xcsh_api_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_discovery reference."
---

# xcsh_api_discovery reference

<a id="canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- Property reference

<a id="canonical-0132321030333032-0010210232222330-0211320201230220-3321121210223333-1303023330301123-1302131211123131-3330100221202320-0221230213303322"></a>

### Direct properties for `xcsh_api_discovery`

<a id="canonical-2111110031210221-3102021321301101-2031002333312332-0120033101201322-0033302010030310-2333111130311002-2211330120001330-2022011003000023"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

<a id="canonical-2303122201020110-3222323301200021-1220133121220320-0132223102222203-2311330223321100-0231131230001002-1301330313001121-0222113023201303"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the APIDiscovery.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3132312111112203-3123003001131230-1230200220211311-3223332202100033-0103132212102200-1022010332332130-0030101010023030-3200020221122222"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1131030220020301-3131203200203333-0133313032121223-0001222203001223-1301222302033210-1230110211101031-3111323323232302-1110103022000111"></a>

<a id="canonical-1120231233300031-2323211311122100-1322212121030021-3232102320213221-3203011023112331-3032133020130201-3112101020120323-0201010032132002"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Additional upstream details:

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

<a id="canonical-2232210200122001-3303321311030311-2121133332012123-0233332300331323-2111022102231201-0222211003323122-0001122211231031-1012220223100233"></a>

#### `name` property

Type: `"string"`. Required.

Name of the APIDiscovery.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3110322032133100-1033331322210222-1223303302330331-3301221120022322-0133113033023102-3110002013223211-2113132012131322-3301330313233022"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the APIDiscovery exists.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0030110120033213-1122030202221331-1211022211102010-0333300030333303-3201113121130010-0202320003311221-2003130101000200-3333121110200320"></a>

### All schema paths for `xcsh_api_discovery`

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

<a id="canonical-0030011001132310-0130023101132022-3322210030220302-0321012110320203-0313220020332330-0100101001100333-1013101323033200-3133301202233011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_auth_types` properties

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- custom_auth_types

<a id="canonical-0003003120201223-2121123011122303-0220322321333021-0313033221122332-0122323312311103-2123123232021321-1302210111132012-3201031310201131"></a>

Type: `"list"`. Computed.

Select your custom authentication types to be detected in the API discovery. Defaults to \`\[\]\`.
Server applies default when omitted.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0300032210123130-0233112330213312-2002222330022333-2201201132031131-2112001133111333-1332100002132120-0132121312203112-0232100021033020"></a>

### Direct properties for `custom_auth_types`

<a id="canonical-3213023002223022-1030032101210302-0210331000103220-1023200102220012-2001322201233021-2111312131330100-3130113332212032-0030211102131123"></a>

#### `custom_auth_types.parameter_name` property

Type: `"string"`. Computed.

Parameter Name. The authentication parameter name.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3200312331002312-3133030021203003-2302333100222330-0220000302210203-3212331310330022-0103323212132111-2101130233330310-2203331003200000"></a>

#### `custom_auth_types.parameter_type` property

Type: `"string"`. Computed.

\[Enum: QUERY\_PARAMETER|HEADER|COOKIE\] Enumeration for authentication parameter types. Possible
values are \`QUERY\_PARAMETER\`, \`HEADER\`, \`COOKIE\`. Defaults to \`QUERY\_PARAMETER\`.

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

<a id="canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy` properties

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

<a id="canonical-0010103010323011-2132210001030201-1302322111103003-1031000122120231-0330313232110312-1020320033133110-1230111300011311-3103332100202132"></a>

### Direct properties for `user_defined_api_discovery_policy`

- [discovery_rules](data-sources--api_discovery--reference--group-001.md#canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333): complete subsection reference.

- [exclusive](data-sources--api_discovery--reference--group-001.md#canonical-2300322012120002-0201232221212210-2123211221021011-2200001022101221-2331331221020112-0010131103023101-3003002313313310-3030102012302313): complete subsection reference.

- [inclusive](data-sources--api_discovery--reference--group-001.md#canonical-0330322003032013-1130101013033123-0030320021111200-2010331110331223-1321323331332301-3132233220120201-3023322100000123-1301321323211020): complete subsection reference.

<a id="canonical-1320033202213233-0300312101013032-1002312311300000-0200003100232023-2122110001022203-2111020233010131-0112310221323210-2101333003223333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules` properties

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2202200332120303-1330122102332103-0333111331330211-3213100122223300-3000123201200021-1001112330311230-2333120301001201-2312120022311021"></a>

### Direct properties for `user_defined_api_discovery_policy.discovery_rules`

- [labels](data-sources--api_discovery--reference--group-001.md#canonical-1131022301210223-3012031113213103-2013203312321000-0222111110021202-3101003110022310-2213112031003113-0213111221132101-2320203002031110): complete subsection reference.

- [metadata](data-sources--api_discovery--reference--group-001.md#canonical-3122231211211030-1132230033000222-0131003213231223-3030113101121210-0201211301123000-0331121111331213-0213330001100032-3310233023300002): complete subsection reference.

- [rule_properties](data-sources--api_discovery--reference--group-001.md#canonical-3323303011223032-3320130223201120-1023021232331211-1310102121212133-1130012211020030-3020120001300323-2232230022103321-0131120302102312): complete subsection reference.

<a id="canonical-1131022301210223-3012031113213103-2013203312321000-0222111110021202-3101003110022310-2213112031003113-0213111221132101-2320203002031110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.labels` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122231211211030-1132230033000222-0131003213231223-3030113101121210-0201211301123000-0331121111331213-0213330001100032-3310233023300002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.metadata` properties

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

<a id="canonical-2013013210200201-0310221120201331-0302133232131321-0120010200222222-0201322310012330-3003331010113330-1220212120111122-1330220102030231"></a>

### Direct properties for `user_defined_api_discovery_policy.discovery_rules.metadata`

<a id="canonical-3030222112000331-0123202302312123-2223312130110210-2302002121032213-2201002333100002-3233020223310313-0031212331201123-3312011032312301"></a>

#### `user_defined_api_discovery_policy.discovery_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1113231211313120-2103330122001000-3110301313222211-1102331321201331-1333323211300301-1031231002322122-0131231130101311-2110301020120012"></a>

<a id="canonical-0100003031223232-2212012300003320-0131133133233321-3203323001102300-0233031310301202-0301111203320113-1101201010203200-2030012203111030"></a>

#### `user_defined_api_discovery_policy.discovery_rules.metadata.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3323303011223032-3320130223201120-1023021232331211-1310102121212133-1130012211020030-3020120001300323-2232230022103321-0131120302102312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.rule_properties` properties

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

<a id="canonical-3310322212011122-0321211302302203-0310130203213102-2330133001321112-0312303021212320-1030030033212012-1000122131231310-1002300022211202"></a>

### Direct properties for `user_defined_api_discovery_policy.discovery_rules.rule_properties`

- [exclusion](data-sources--api_discovery--reference--group-001.md#canonical-3221310322202333-1221230121213123-1112313232331200-2332201220132023-3332123133120032-0311132320213302-0122010033020011-0111010011132221): complete subsection reference.

- [http_header_criteria](data-sources--api_discovery--reference--group-001.md#canonical-0222020233023033-3023331301211103-3221220323212121-2202212123211223-2023200012200111-0210120031121320-2030310021331111-2233003311001121): complete subsection reference.

- [inclusion](data-sources--api_discovery--reference--group-001.md#canonical-2202002321131002-3330102300321213-1021103211211010-0031121133203110-2000320221223110-0001030331112102-3320032102023210-1133122003011331): complete subsection reference.

<a id="canonical-3321030031233303-0121131321322313-3012321123121020-0130202300231313-0011222310131230-0200213320210331-1312202100102331-2301010201032102"></a>

<a id="canonical-0110313123220212-1222233100302231-1223222100122022-2301011122221201-3033311203331202-1222221202101021-0332113332131212-1122332020233222"></a>

#### `user_defined_api_discovery_policy.discovery_rules.rule_properties.pattern` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3221310322202333-1221230121213123-1112313232331200-2332201220132023-3332123133120032-0311132320213302-0122010033020011-0111010011132221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion` properties

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

<a id="canonical-2223313110131121-0302311013230130-1230230200333313-0010310001131102-0111302320230211-2122023220222123-3220333023010122-0200331203223132"></a>

### Direct properties for `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion`

- [archive](data-sources--api_discovery--reference--group-001.md#canonical-1211030321223333-3011323023231023-2120113223210331-3000003313222033-0100011110320222-1123033013132313-1302202322203212-3220033203002032): complete subsection reference.

- [ignore](data-sources--api_discovery--reference--group-001.md#canonical-2223202221010110-1020301033200130-3311232122330202-3201032022232230-2232011131230312-1213132103322232-3310112002112113-2130212132121212): complete subsection reference.

<a id="canonical-1211030321223333-3011323023231023-2120113223210331-3000003313222033-0100011110320222-1123033013132313-1302202322203212-3220033203002032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive` properties

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

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223202221010110-1020301033200130-3311232122330202-3201032022232230-2232011131230312-1213132103322232-3310112002112113-2130212132121212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore` properties

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

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222020233023033-3023331301211103-3221220323212121-2202212123211223-2023200012200111-0210120031121320-2030310021331111-2233003311001121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria` properties

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

Additional upstream details:

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

<a id="canonical-1303123013020311-0030320101120301-1011020120233130-1021222313031321-3331020213222211-0233123131231112-3331212110322001-0133012020333022"></a>

### Direct properties for `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria`

<a id="canonical-1302231133122212-1323120020033321-2212011202213000-0311030222003000-2133111000120310-1001133130120031-3032310100002103-2303122020032022"></a>

#### `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.field_name` property

Type: `"string"`. Computed.

HTTP Header Name. Human-readable name for the resource

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0312013131002222-0021223000111211-2222203112201313-3033001320011312-3000032112100323-2032233231121001-3000120112332200-1310101023223323"></a>

#### `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.location` property

Type: `"string"`. Computed.

\[Enum: REQUEST|RESPONSE\] Specifies whether the rule criteria should be evaluated against request
or response Applies the rule to incoming traffic from the client. Applies the rule to outgoing
traffic sent back to the client. Possible values are \`REQUEST\`, \`RESPONSE\`. Defaults to
\`REQUEST\`.

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

<a id="canonical-0133232222102313-3223300133331131-2100201013112000-2102103102231003-2220031010003003-3032011021200133-2331130301003020-3301320102131132"></a>

#### `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.match_type` property

Type: `"string"`. Computed.

\[Enum: EXACT\_MATCH|SUBSTRING|REGEX\] Specifies how the value should be matched. Possible values
are \`EXACT\_MATCH\`, \`SUBSTRING\`, \`REGEX\`. Defaults to \`EXACT\_MATCH\`.

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

<a id="canonical-2322321013111111-3002032111023130-0103330032201003-3110322002321031-1101202223223232-1112202323201232-3103321310332223-2333301330023033"></a>

#### `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.value` property

Type: `"string"`. Computed.

Value. Configuration parameter for value

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2202002321131002-3330102300321213-1021103211211010-0031121133203110-2000320221223110-0001030331112102-3320032102023210-1133122003011331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion` properties

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

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300322012120002-0201232221212210-2123211221021011-2200001022101221-2331331221020112-0010131103023101-3003002313313310-3030102012302313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.exclusive` properties

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- user_defined_api_discovery_policy.exclusive

<a id="canonical-2312030003110120-1323331300320032-3021001331200032-0111111123100013-1033222102020112-0133322221132121-2312223222100313-0301233011012110"></a>

Type: `"single"`. Computed.

Exclusion Configuration. Configuration for exclusion action.

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

<a id="canonical-0233023300121121-1130132300201323-0112100111302000-1332002120101101-3333230030003300-0103020113300111-0221030020212333-0303021121301000"></a>

### Direct properties for `user_defined_api_discovery_policy.exclusive`

- [archive](data-sources--api_discovery--reference--group-001.md#canonical-1031032321002013-3131013331022200-3212002300223302-2211032220121221-3213220330313320-3011123203030133-2120021033132130-3213331332220113): complete subsection reference.

- [ignore](data-sources--api_discovery--reference--group-001.md#canonical-0200303112032033-1320213032202301-0222031203313012-0230323332212013-1001123000133311-2121233111121231-2312232131232101-2221323230312221): complete subsection reference.

<a id="canonical-1031032321002013-3131013331022200-3212002300223302-2211032220121221-3213220330313320-3011123203030133-2120021033132130-3213331332220113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.exclusive.archive` properties

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--reference--group-001.md#canonical-2300322012120002-0201232221212210-2123211221021011-2200001022101221-2331331221020112-0010131103023101-3003002313313310-3030102012302313)
- user_defined_api_discovery_policy.exclusive.archive

<a id="canonical-3121113303223022-3300321111101303-3120301220322123-2101122013322200-0210012023030320-2121012010222301-3201200311200131-2222032132313202"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200303112032033-1320213032202301-0222031203313012-0230323332212013-1001123000133311-2121233111121231-2312232131232101-2221323230312221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.exclusive.ignore` properties

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- [user_defined_api_discovery_policy.exclusive](data-sources--api_discovery--reference--group-001.md#canonical-2300322012120002-0201232221212210-2123211221021011-2200001022101221-2331331221020112-0010131103023101-3003002313313310-3030102012302313)
- user_defined_api_discovery_policy.exclusive.ignore

<a id="canonical-2333011132001130-2333101122133330-1310323303311222-1130231330022031-1130323021130101-0233103323233122-2131112221032310-2330031301302121"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0330322003032013-1130101013033123-0030320021111200-2010331110331223-1321323331332301-3132233220120201-3023322100000123-1301321323211020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_defined_api_discovery_policy.inclusive` properties

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Property reference](data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [user_defined_api_discovery_policy](data-sources--api_discovery--reference--group-001.md#canonical-0322233233132331-0002320020312122-1132122123302210-3030303231112130-3031032330222000-0012333023101233-3010101123322210-0230223212002133)
- user_defined_api_discovery_policy.inclusive

<a id="canonical-1320222130100200-3012303112033213-2020321300211202-3200230113130222-1230101221130302-1231011020332323-3131122122223300-0000011011113313"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.
