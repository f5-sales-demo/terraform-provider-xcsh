---
page_title: "xcsh_app_type reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_type reference."
---

# xcsh_app_type reference

<a id="canonical-3221103121003133-1230232300121110-2122110130133103-3202001113101201-3311021023133001-3112111111202133-2003201302132001-0322123211030221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-2200113311021331-0200020232011230-0011223212102321-1001021222113033-0011022023131301-0311022330302220-2111320322332313-1010123013021130)
- Property reference

<a id="canonical-3303020030332212-2222102331321320-1012300102100231-1321330331023130-3232302323323112-2113320030302122-0211031313112320-3313022022210320"></a>

### Direct properties for `xcsh_app_type`

<a id="canonical-2131302023100233-2013100322112021-2012003031020302-3312000331012232-3200130130230010-0300012211100123-2220011003221032-3031320203133032"></a>

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

- [business_logic_markup_setting](data-sources--app_type--reference--group-001.md#canonical-2323122110222310-1232231121122003-3322001112111020-3333201210332002-0003210220103112-1211021101213100-0313132102013202-3311311013201011): complete subsection reference.

<a id="canonical-3331121322030312-2012103103211221-1001212021203133-0231203211010213-2320113312210023-1021100133203320-2122020233131132-2020113212120222"></a>

<a id="canonical-0012111013232331-1111012123002232-3201303233303312-3303100203303210-3032233001010102-2210232321311110-3111322220202002-1301311101220013"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the AppType.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [features](data-sources--app_type--reference--group-001.md#canonical-3022321230021323-1100011333312000-3101300310230021-2220221313113330-0300330220330011-1133023330331010-1100131101013300-3031001121200112): complete subsection reference.

<a id="canonical-3320330031111221-1222301311210231-1113110013111020-0330103330121223-3232302030220122-3001322312221300-2220203002201103-0122200213231312"></a>

<a id="canonical-3033122333230331-2301220120031321-0332321000011132-3233011312312111-3010123013203331-1313030111222021-1003120330201120-0333022330321301"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1002033132002103-3301030303302211-1111123033222312-1121222011002313-2330212021213311-0303130232333332-1223103020213303-3203211302130012"></a>

<a id="canonical-0300212312230011-1110003202331222-0130133300133120-2212110010022031-0300021202003333-3132112333332212-3001033132100301-1002313231000223"></a>

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

<a id="canonical-1330120022111013-1330300102312020-0313030303231222-1223110031112201-1312020203103101-2002310310321221-0021223131111212-0202233212331302"></a>

<a id="canonical-0312232031221312-3202213223010002-3303222101212102-0031121122320221-2312022000123011-0222312000023302-0100212320321220-3231322201133111"></a>

#### `name` property

Type: `"string"`. Required.

Name of the AppType.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3223232222313212-1021223022323210-1301211000321222-2311033131222110-3032110230320000-1201221321220021-3221223011020222-0231211300123322"></a>

<a id="canonical-1111220221301213-1100223003020032-1033033331233230-3110122031303330-0021220001130130-2211301011001322-2223000313002133-0013110331332010"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the AppType exists.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0012120303103303-0233102330031121-1020103221102230-0311233122013112-0013311132010000-0223230012030303-3110220130322233-2133120333112020"></a>

### All schema paths for `xcsh_app_type`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--app_type--reference--group-001.md#canonical-2131302023100233-2013100322112021-2012003031020302-3312000331012232-3200130130230010-0300012211100123-2220011003221032-3031320203133032) |
| `business_logic_markup_setting` | [business_logic_markup_setting](data-sources--app_type--reference--group-001.md#canonical-2321011131120120-3112012303101320-3331103202033020-2013021033000033-1212210111022321-1130011012000120-3030223201302003-0200010323313132) |
| `business_logic_markup_setting.disable_spec` | [business_logic_markup_setting.disable_spec](data-sources--app_type--reference--group-001.md#canonical-0001231101020020-3120232011010203-2012231010031210-1130023331311231-1112010221233332-3121313332322011-3213300202201301-2220320010223013) |
| `business_logic_markup_setting.discovered_api_settings` | [business_logic_markup_setting.discovered_api_settings](data-sources--app_type--reference--group-001.md#canonical-0333103200321223-1332213321320123-2002223022010223-3020133010022333-2230322102202023-2221223021121120-0323000330230123-3113233322212300) |
| `business_logic_markup_setting.discovered_api_settings.purge_duration_for_inactive_discovered_apis` | [business_logic_markup_setting.discovered_api_settings.purge_duration_for_inactive_discovered_apis](data-sources--app_type--reference--group-001.md#canonical-1313121322231020-2111201312231002-0010110230122330-1313022232132031-1230131212301323-3301232130112101-3122020332321132-0123101002230312) |
| `business_logic_markup_setting.enable` | [business_logic_markup_setting.enable](data-sources--app_type--reference--group-001.md#canonical-0320101302322112-3302301112203320-1101231301031213-3002200121021103-3301333120332321-2020330210302131-2321131323120022-3132100122132122) |
| `description` | [description](data-sources--app_type--reference--group-001.md#canonical-3331121322030312-2012103103211221-1001212021203133-0231203211010213-2320113312210023-1021100133203320-2122020233131132-2020113212120222) |
| `features` | [features](data-sources--app_type--reference--group-001.md#canonical-0231233100331200-1030002201233312-3220333312312113-1013003020000331-1031000223031022-2311103311213212-2003012223322220-0121312320100223) |
| `features.type` | [features.type](data-sources--app_type--reference--group-001.md#canonical-0332123121002010-2223220210221313-0022001113023121-3010013321121102-0200011103130302-3122023320213310-0213311032312123-0300111300210201) |
| `id` | [ID](data-sources--app_type--reference--group-001.md#canonical-3320330031111221-1222301311210231-1113110013111020-0330103330121223-3232302030220122-3001322312221300-2220203002201103-0122200213231312) |
| `labels` | [labels](data-sources--app_type--reference--group-001.md#canonical-1002033132002103-3301030303302211-1111123033222312-1121222011002313-2330212021213311-0303130232333332-1223103020213303-3203211302130012) |
| `name` | [name](data-sources--app_type--reference--group-001.md#canonical-1330120022111013-1330300102312020-0313030303231222-1223110031112201-1312020203103101-2002310310321221-0021223131111212-0202233212331302) |
| `namespace` | [namespace](data-sources--app_type--reference--group-001.md#canonical-3223232222313212-1021223022323210-1301211000321222-2311033131222110-3032110230320000-1201221321220021-3221223011020222-0231211300123322) |

<a id="canonical-2323122110222310-1232231121122003-3322001112111020-3333201210332002-0003210220103112-1211021101213100-0313132102013202-3311311013201011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `business_logic_markup_setting` properties

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-2200113311021331-0200020232011230-0011223212102321-1001021222113033-0011022023131301-0311022330302220-2111320322332313-1010123013021130)
- [Property reference](data-sources--app_type--reference--group-001.md#canonical-3221103121003133-1230232300121110-2122110130133103-3202001113101201-3311021023133001-3112111111202133-2003201302132001-0322123211030221)
- business_logic_markup_setting

<a id="canonical-2321011131120120-3112012303101320-3331103202033020-2013021033000033-1212210111022321-1130011012000120-3030223201302003-0200010323313132"></a>

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
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable\",\"enable\"]"
}
```

<a id="canonical-3133230121310103-0132303011320121-0111023302032031-2133230203211032-1203122201221110-1333110202221300-0302203330121223-3120322101323220"></a>

### Direct properties for `business_logic_markup_setting`

- [disable_spec](data-sources--app_type--reference--group-001.md#canonical-2301103100301020-0121322203203032-1301031121220131-0323122212210331-1133131103012300-1112133302323202-0321200321011023-0113300200030113): complete subsection reference.

- [discovered_api_settings](data-sources--app_type--reference--group-001.md#canonical-2032222223210300-2133102122221101-0302332213001120-0000102200020222-2012301022111102-0013310120101020-2013121102320313-0300133021102221): complete subsection reference.

- [enable](data-sources--app_type--reference--group-001.md#canonical-0331323010021032-0333133230201012-0312222203021131-3123301121112221-2110312123102331-1132012330031201-0322120301202203-3230122112113201): complete subsection reference.

<a id="canonical-2301103100301020-0121322203203032-1301031121220131-0323122212210331-1133131103012300-1112133302323202-0321200321011023-0113300200030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `business_logic_markup_setting.disable_spec` properties

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-2200113311021331-0200020232011230-0011223212102321-1001021222113033-0011022023131301-0311022330302220-2111320322332313-1010123013021130)
- [Property reference](data-sources--app_type--reference--group-001.md#canonical-3221103121003133-1230232300121110-2122110130133103-3202001113101201-3311021023133001-3112111111202133-2003201302132001-0322123211030221)
- [business_logic_markup_setting](data-sources--app_type--reference--group-001.md#canonical-2323122110222310-1232231121122003-3322001112111020-3333201210332002-0003210220103112-1211021101213100-0313132102013202-3311311013201011)
- business_logic_markup_setting.disable_spec

<a id="canonical-0001231101020020-3120232011010203-2012231010031210-1130023331311231-1112010221233332-3121313332322011-3213300202201301-2220320010223013"></a>

Type: `["object", {}]`. Computed.

Enable this option

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032222223210300-2133102122221101-0302332213001120-0000102200020222-2012301022111102-0013310120101020-2013121102320313-0300133021102221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `business_logic_markup_setting.discovered_api_settings` properties

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-2200113311021331-0200020232011230-0011223212102321-1001021222113033-0011022023131301-0311022330302220-2111320322332313-1010123013021130)
- [Property reference](data-sources--app_type--reference--group-001.md#canonical-3221103121003133-1230232300121110-2122110130133103-3202001113101201-3311021023133001-3112111111202133-2003201302132001-0322123211030221)
- [business_logic_markup_setting](data-sources--app_type--reference--group-001.md#canonical-2323122110222310-1232231121122003-3322001112111020-3333201210332002-0003210220103112-1211021101213100-0313132102013202-3311311013201011)
- business_logic_markup_setting.discovered_api_settings

<a id="canonical-0333103200321223-1332213321320123-2002223022010223-3020133010022333-2230322102202023-2221223021121120-0323000330230123-3113233322212300"></a>

Type: `"single"`. Computed.

Discovered API Settings. Configure Discovered API Settings.

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

<a id="canonical-3013310321023222-0132120233232333-2132200030210123-3130122320200013-0130213311031032-0200221230330330-3212010213321132-2031300103133000"></a>

### Direct properties for `business_logic_markup_setting.discovered_api_settings`

<a id="canonical-1313121322231020-2111201312231002-0010110230122330-1313022232132031-1230131212301323-3301232130112101-3122020332321132-0123101002230312"></a>

#### `business_logic_markup_setting.discovered_api_settings.purge_duration_for_inactive_discovered_apis` property

Type: `"number"`. Computed.

Inactive discovered API will be deleted after configured duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

<a id="canonical-0331323010021032-0333133230201012-0312222203021131-3123301121112221-2110312123102331-1132012330031201-0322120301202203-3230122112113201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `business_logic_markup_setting.enable` properties

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-2200113311021331-0200020232011230-0011223212102321-1001021222113033-0011022023131301-0311022330302220-2111320322332313-1010123013021130)
- [Property reference](data-sources--app_type--reference--group-001.md#canonical-3221103121003133-1230232300121110-2122110130133103-3202001113101201-3311021023133001-3112111111202133-2003201302132001-0322123211030221)
- [business_logic_markup_setting](data-sources--app_type--reference--group-001.md#canonical-2323122110222310-1232231121122003-3322001112111020-3333201210332002-0003210220103112-1211021101213100-0313132102013202-3311311013201011)
- business_logic_markup_setting.enable

<a id="canonical-0320101302322112-3302301112203320-1101231301031213-3002200121021103-3301333120332321-2020330210302131-2321131323120022-3132100122132122"></a>

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

<a id="canonical-3022321230021323-1100011333312000-3101300310230021-2220221313113330-0300330220330011-1133023330331010-1100131101013300-3031001121200112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `features` properties

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-2200113311021331-0200020232011230-0011223212102321-1001021222113033-0011022023131301-0311022330302220-2111320322332313-1010123013021130)
- [Property reference](data-sources--app_type--reference--group-001.md#canonical-3221103121003133-1230232300121110-2122110130133103-3202001113101201-3311021023133001-3112111111202133-2003201302132001-0322123211030221)
- features

<a id="canonical-0231233100331200-1030002201233312-3220333312312113-1013003020000331-1031000223031022-2311103311213212-2003012223322220-0121312320100223"></a>

Type: `"list"`. Computed.

List of various advanced security features enabled.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3212223212211023-3031130113113200-3033322223312222-3100123121322120-2013000201100333-1002220333210231-3221320230210010-2330110233223031"></a>

### Direct properties for `features`

<a id="canonical-0332123121002010-2223220210221313-0022001113023121-3010013321121102-0200011103130302-3122023320213310-0213311032312123-0300111300210201"></a>

#### `features.type` property

Type: `"string"`. Computed.

\[Enum:
BUSINESS\_LOGIC\_MARKUP|TIMESERIES\_ANOMALY\_DETECTION|PER\_REQ\_ANOMALY\_DETECTION|USER\_BEHAVIOR\_ANALYSIS\]
Enumeration for advanced security features supported API Discovery enables generation of model for
various API interactions between services of App type. Enable analysis of timeseries for various
metric collected like requests, errors, latency etc. Enable anomaly detection per API request, i.e.
Possible values are \`BUSINESS\_LOGIC\_MARKUP\`, \`TIMESERIES\_ANOMALY\_DETECTION\`,
\`PER\_REQ\_ANOMALY\_DETECTION\`, \`USER\_BEHAVIOR\_ANALYSIS\`. Defaults to
\`BUSINESS\_LOGIC\_MARKUP\`.

Additional upstream details:

Enumeration for advanced security features supported

API Discovery enables generation of model for various API interactions between services of App type.
The probability density function (PDF) charts generation for API endpoints Enable user behavior
analysis.

Receipt-pinned upstream constraints:

```json
{
  "default": "BUSINESS_LOGIC_MARKUP",
  "enum": [
    "BUSINESS_LOGIC_MARKUP",
    "TIMESERIES_ANOMALY_DETECTION",
    "PER_REQ_ANOMALY_DETECTION",
    "USER_BEHAVIOR_ANALYSIS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
