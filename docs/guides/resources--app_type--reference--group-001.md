---
page_title: "xcsh_app_type reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_type reference."
---

# xcsh_app_type reference

<a id="canonical-2010013132232302-1121312303131032-1120000202313211-0303313333201213-1313211102301122-0021132003012310-2330313333002232-1320001121301320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-2220100013322301-3233230230220131-2230301203210321-0110213003232211-2331220333323130-0302001023303001-2320103333130202-1122210103122102)
- Property reference

<a id="canonical-3021121012201330-2012131031312213-3021203002301211-2230221231230223-0100013033233121-0222002103121220-3310321311022233-2321323333011330"></a>

### Direct properties for `xcsh_app_type`

<a id="canonical-2010303301101302-3010211132222332-0213222221111133-0013022022103010-3111302200230212-2213313322333003-0202331032310031-1322103023230333"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

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

- [business_logic_markup_setting](resources--app_type--reference--group-001.md#canonical-3223322110003203-1122212310332100-2122030232112232-1223332130013002-0311230312202212-3223203333012301-0103213121203013-3021020310331223): complete subsection reference.

<a id="canonical-3023302133212010-2123232230223213-1122220200212220-3101333222312200-1122331031110200-3131020300203312-1221001233023132-0132322111232100"></a>

<a id="canonical-0010303011223313-0322102230020002-1211330120323000-1310220112211022-0212232330100300-3221222112003033-2111110220312300-3233223112331101"></a>

#### `description` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2212101031003113-3023323321013311-3303330120022203-1012332231233211-1120323321020220-3311111033111202-3230233301013202-3221221011321230"></a>

<a id="canonical-3331223120201300-1112332211230310-1033020102331021-1322200230000123-1211020203120300-2331121133133131-1032033200121101-3202020133131003"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

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

- [features](resources--app_type--reference--group-001.md#canonical-1311121001333100-0000000111012131-2121330230031311-1213111133111000-3123303122012022-0203021310311320-3322202132100311-1222302220321220): complete subsection reference.

<a id="canonical-3233032221232022-2202113101011222-2111203103022211-0131130033322011-2030331011321000-3322102222221323-3130333223213100-1210111313312132"></a>

<a id="canonical-1121122012111302-2100330303233222-2232100020111303-2110233131232032-0221021212210133-3333302131301031-3132321102333103-0033310211313023"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0132030123030212-0000031031301011-1232130023203302-0321223322220121-2330333112311100-2131031112222002-2202132231332110-3001103322030302"></a>

<a id="canonical-2222113110022330-1133302120210210-2310122231231001-2231302233302011-2013133322013030-1322301201230013-1111212300222210-0223300331003102"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-1211321322313121-0332132323201321-1011111300013103-2120030102313202-2033302100100300-0230021200003002-1020331133010332-0101200032203210"></a>

<a id="canonical-0200111200322303-1332222102301103-3300020203110321-3202323022022221-2330302331233103-0130033331222012-2010212211011213-0123310232010313"></a>

#### `name` property

Type: `"string"`. Required.

Name of the App Type. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1123203303120012-1120030022230103-1133200233201303-2333013110300223-0300033022021301-1111113003110212-2031011211003023-2001010211211030"></a>

<a id="canonical-3303311103321312-2213332110102310-0212030233300202-2020312030232210-1321310132010200-0112221001103223-2030301330302230-3031311032120020"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the App Type is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [timeouts](resources--app_type--reference--group-001.md#canonical-3312022023000301-2001323303001121-2233203100323000-0110321210321131-2022000300323231-3033133330000133-2323211331333321-2011010020010022): complete subsection reference.

<a id="canonical-1210311101123100-3101130003223022-1003201213212202-0121322023131202-3330203113220033-3030133103231233-2011012123120022-2132332300313323"></a>

### All schema paths for `xcsh_app_type`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--app_type--reference--group-001.md#canonical-2010303301101302-3010211132222332-0213222221111133-0013022022103010-3111302200230212-2213313322333003-0202331032310031-1322103023230333) |
| `business_logic_markup_setting` | [business_logic_markup_setting](resources--app_type--reference--group-001.md#canonical-3302211102203312-3021232322221011-2203132010122201-0321220032220320-3031233331332222-2130320121301000-3232310130222311-1333131323121022) |
| `business_logic_markup_setting.disable_spec` | [business_logic_markup_setting.disable_spec](resources--app_type--reference--group-001.md#canonical-3302323222023212-2031301101032032-2111003032010231-2303230033323031-3033031301333212-1331200223101010-3112312313213330-2321210121222133) |
| `business_logic_markup_setting.discovered_api_settings` | [business_logic_markup_setting.discovered_api_settings](resources--app_type--reference--group-001.md#canonical-3011033323121113-3222302132331012-3331222210100332-2201102300133322-0110323101022333-2303032310323103-2111201302230200-3003001022033121) |
| `business_logic_markup_setting.discovered_api_settings.purge_duration_for_inactive_discovered_apis` | [business_logic_markup_setting.discovered_api_settings.purge_duration_for_inactive_discovered_apis](resources--app_type--reference--group-001.md#canonical-0321211010101321-2021001011010033-2223211311133012-2332202202310030-2030113322021321-0032211222210333-1021312322111231-0020122331203320) |
| `business_logic_markup_setting.enable` | [business_logic_markup_setting.enable](resources--app_type--reference--group-001.md#canonical-3222230212023121-3033103301223300-2133223033123120-0132220233303320-0222300101200200-1231122322011202-2122133321020200-2231221320031030) |
| `description` | [description](resources--app_type--reference--group-001.md#canonical-3023302133212010-2123232230223213-1122220200212220-3101333222312200-1122331031110200-3131020300203312-1221001233023132-0132322111232100) |
| `disable` | [disable](resources--app_type--reference--group-001.md#canonical-2212101031003113-3023323321013311-3303330120022203-1012332231233211-1120323321020220-3311111033111202-3230233301013202-3221221011321230) |
| `features` | [features](resources--app_type--reference--group-001.md#canonical-0031200033100323-1123130211311102-2000120110321231-1033120112121113-1312021212111301-2200133312202320-3013001003223330-1231311203301230) |
| `features.type` | [features.type](resources--app_type--reference--group-001.md#canonical-3113310212203033-3322021022131032-0112102112331200-3033223121121321-1012301012223113-0222111211112023-2232011013132300-0232300101230302) |
| `id` | [ID](resources--app_type--reference--group-001.md#canonical-3233032221232022-2202113101011222-2111203103022211-0131130033322011-2030331011321000-3322102222221323-3130333223213100-1210111313312132) |
| `labels` | [labels](resources--app_type--reference--group-001.md#canonical-0132030123030212-0000031031301011-1232130023203302-0321223322220121-2330333112311100-2131031112222002-2202132231332110-3001103322030302) |
| `name` | [name](resources--app_type--reference--group-001.md#canonical-1211321322313121-0332132323201321-1011111300013103-2120030102313202-2033302100100300-0230021200003002-1020331133010332-0101200032203210) |
| `namespace` | [namespace](resources--app_type--reference--group-001.md#canonical-1123203303120012-1120030022230103-1133200233201303-2333013110300223-0300033022021301-1111113003110212-2031011211003023-2001010211211030) |
| `timeouts` | [timeouts](resources--app_type--reference--group-001.md#canonical-3201131213333020-3221311100302233-1010232011002213-1222130223233332-3202333301233232-1111231112110212-3333010003222103-2112300223001302) |
| `timeouts.create` | [timeouts.create](resources--app_type--reference--group-001.md#canonical-0132023313103031-3221300013011203-2100202110310232-3102020020223233-2030201202300200-1012221311021100-2020220033010300-1213322002300232) |
| `timeouts.delete` | [timeouts.delete](resources--app_type--reference--group-001.md#canonical-3231333132220033-2103201302201031-1222103120121300-0201002200122133-3100122202313123-2301113213120123-0332012122311310-1031200002010333) |
| `timeouts.read` | [timeouts.read](resources--app_type--reference--group-001.md#canonical-3300303020000033-2333201300210002-1133302230232000-2010220010331010-0012022233300112-0233321110123021-2132331021011230-0001102313002303) |
| `timeouts.update` | [timeouts.update](resources--app_type--reference--group-001.md#canonical-2201002320303311-0331000200302110-3221130010312301-2113013203021200-3232033220002222-0132133113330201-1013003231103330-1322000312002202) |

<a id="canonical-3223322110003203-1122212310332100-2122030232112232-1223332130013002-0311230312202212-3223203333012301-0103213121203013-3021020310331223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `business_logic_markup_setting` properties

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-2220100013322301-3233230230220131-2230301203210321-0110213003232211-2331220333323130-0302001023303001-2320103333130202-1122210103122102)
- [Property reference](resources--app_type--reference--group-001.md#canonical-2010013132232302-1121312303131032-1120000202313211-0303313333201213-1313211102301122-0021132003012310-2330313333002232-1320001121301320)
- business_logic_markup_setting

<a id="canonical-3302211102203312-3021232322221011-2203132010122201-0321220032220320-3031233331332222-2130320121301000-3232310130222311-1333131323121022"></a>

Type: `"object"`. single nested block, Optional.

Settings specifying how API Discovery will be performed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
business_logic_markup_setting {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232102223021113-3232223231230100-0130001001232222-2130211012231233-1103213330121200-2311221102010330-2331130011221333-0331300230103220"></a>

### Direct properties for `business_logic_markup_setting`

- [disable_spec](resources--app_type--reference--group-001.md#canonical-3110203113220000-1323002033203131-2233003102110102-1320022033123322-0220122321031223-2320010331301231-1331322212312100-3133320312132033): complete subsection reference.

- [discovered_api_settings](resources--app_type--reference--group-001.md#canonical-0103023122331110-0120203002302001-3010013203311021-0330032023101200-0111112330010133-2222302332010300-2311301201130012-2130323303212100): complete subsection reference.

- [enable](resources--app_type--reference--group-001.md#canonical-3212031320200323-1032122133100010-2132303001322023-0010131003113211-2213333112013221-1030000033220230-1310021320210103-1030132323201200): complete subsection reference.

<a id="canonical-3110203113220000-1323002033203131-2233003102110102-1320022033123322-0220122321031223-2320010331301231-1331322212312100-3133320312132033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `business_logic_markup_setting.disable_spec` properties

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-2220100013322301-3233230230220131-2230301203210321-0110213003232211-2331220333323130-0302001023303001-2320103333130202-1122210103122102)
- [Property reference](resources--app_type--reference--group-001.md#canonical-2010013132232302-1121312303131032-1120000202313211-0303313333201213-1313211102301122-0021132003012310-2330313333002232-1320001121301320)
- [business_logic_markup_setting](resources--app_type--reference--group-001.md#canonical-3223322110003203-1122212310332100-2122030232112232-1223332130013002-0311230312202212-3223203333012301-0103213121203013-3021020310331223)
- business_logic_markup_setting.disable_spec

<a id="canonical-3302323222023212-2031301101032032-2111003032010231-2303230033323031-3033031301333212-1331200223101010-3112312313213330-2321210121222133"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103023122331110-0120203002302001-3010013203311021-0330032023101200-0111112330010133-2222302332010300-2311301201130012-2130323303212100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `business_logic_markup_setting.discovered_api_settings` properties

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-2220100013322301-3233230230220131-2230301203210321-0110213003232211-2331220333323130-0302001023303001-2320103333130202-1122210103122102)
- [Property reference](resources--app_type--reference--group-001.md#canonical-2010013132232302-1121312303131032-1120000202313211-0303313333201213-1313211102301122-0021132003012310-2330313333002232-1320001121301320)
- [business_logic_markup_setting](resources--app_type--reference--group-001.md#canonical-3223322110003203-1122212310332100-2122030232112232-1223332130013002-0311230312202212-3223203333012301-0103213121203013-3021020310331223)
- business_logic_markup_setting.discovered_api_settings

<a id="canonical-3011033323121113-3222302132331012-3331222210100332-2201102300133322-0110323101022333-2303032310323103-2111201302230200-3003001022033121"></a>

Type: `"object"`. single nested block, Optional.

Discovered API Settings. Configure Discovered API Settings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("purge_duration_for_inactive_discovered_apis")}
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
discovered_api_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211112002123301-0132132100333132-0201100010103313-3013132223332020-3300100121003130-2220330110032013-3021320330133132-3120111311332013"></a>

### Direct properties for `business_logic_markup_setting.discovered_api_settings`

<a id="canonical-0321211010101321-2021001011010033-2223211311133012-2332202202310030-2030113322021321-0032211222210333-1021312322111231-0020122331203320"></a>

#### `business_logic_markup_setting.discovered_api_settings.purge_duration_for_inactive_discovered_apis` property

Type: `"number"`. Optional.

Inactive discovered API will be deleted after configured duration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 7),
}
```

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3212031320200323-1032122133100010-2132303001322023-0010131003113211-2213333112013221-1030000033220230-1310021320210103-1030132323201200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `business_logic_markup_setting.enable` properties

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-2220100013322301-3233230230220131-2230301203210321-0110213003232211-2331220333323130-0302001023303001-2320103333130202-1122210103122102)
- [Property reference](resources--app_type--reference--group-001.md#canonical-2010013132232302-1121312303131032-1120000202313211-0303313333201213-1313211102301122-0021132003012310-2330313333002232-1320001121301320)
- [business_logic_markup_setting](resources--app_type--reference--group-001.md#canonical-3223322110003203-1122212310332100-2122030232112232-1223332130013002-0311230312202212-3223203333012301-0103213121203013-3021020310331223)
- business_logic_markup_setting.enable

<a id="canonical-3222230212023121-3033103301223300-2133223033123120-0132220233303320-0222300101200200-1231122322011202-2122133321020200-2231221320031030"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311121001333100-0000000111012131-2121330230031311-1213111133111000-3123303122012022-0203021310311320-3322202132100311-1222302220321220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `features` properties

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-2220100013322301-3233230230220131-2230301203210321-0110213003232211-2331220333323130-0302001023303001-2320103333130202-1122210103122102)
- [Property reference](resources--app_type--reference--group-001.md#canonical-2010013132232302-1121312303131032-1120000202313211-0303313333201213-1313211102301122-0021132003012310-2330313333002232-1320001121301320)
- features

<a id="canonical-0031200033100323-1123130211311102-2000120110321231-1033120112121113-1312021212111301-2200133312202320-3013001003223330-1231311203301230"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
features {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233012030010212-1232120303220113-3112113013333112-3213313021331223-2011331220120320-2203013101032200-2022231310211300-3000123020120113"></a>

### Direct properties for `features`

<a id="canonical-3113310212203033-3322021022131032-0112102112331200-3033223121121321-1012301012223113-0222111211112023-2232011013132300-0232300101230302"></a>

#### `features.type` property

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["BUSINESS_LOGIC_MARKUP","PER_REQ_ANOMALY_DETECTION","TIMESERIES_ANOMALY_DETECTION","USER_BEHAVIOR_ANALYSIS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("BUSINESS_LOGIC_MARKUP",
    "TIMESERIES_ANOMALY_DETECTION",
    "PER_REQ_ANOMALY_DETECTION",
    "USER_BEHAVIOR_ANALYSIS"),
}
```

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

<a id="canonical-3312022023000301-2001323303001121-2233203100323000-0110321210321131-2022000300323231-3033133330000133-2323211331333321-2011010020010022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-2220100013322301-3233230230220131-2230301203210321-0110213003232211-2331220333323130-0302001023303001-2320103333130202-1122210103122102)
- [Property reference](resources--app_type--reference--group-001.md#canonical-2010013132232302-1121312303131032-1120000202313211-0303313333201213-1313211102301122-0021132003012310-2330313333002232-1320001121301320)
- timeouts

<a id="canonical-3201131213333020-3221311100302233-1010232011002213-1222130223233332-3202333301233232-1111231112110212-3333010003222103-2112300223001302"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130122120310032-2130011323201312-3132013333303310-2202022033112120-0120313020211220-1112322312130213-2222110212121231-1133100020030132"></a>

### Direct properties for `timeouts`

<a id="canonical-0132023313103031-3221300013011203-2100202110310232-3102020020223233-2030201202300200-1012221311021100-2020220033010300-1213322002300232"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3231333132220033-2103201302201031-1222103120121300-0201002200122133-3100122202313123-2301113213120123-0332012122311310-1031200002010333"></a>

<a id="canonical-3333102221332303-1030021231232111-1200000032230333-2132011102113131-0122032103211310-2030030023022012-1333221022023202-2331130313232001"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3300303020000033-2333201300210002-1133302230232000-2010220010331010-0012022233300112-0233321110123021-2132331021011230-0001102313002303"></a>

<a id="canonical-1313333210211101-3002103033320100-0232122232213031-3321321212330220-2301002332202232-3021312032012322-2200121023000121-2230100120233110"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2201002320303311-0331000200302110-3221130010312301-2113013203021200-3232033220002222-0132133113330201-1013003231103330-1322000312002202"></a>

<a id="canonical-3023322002031313-2201301200013320-3123132313320113-3220213223003120-3313333203320101-0002102300012123-1112130332310100-0311003212333020"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
