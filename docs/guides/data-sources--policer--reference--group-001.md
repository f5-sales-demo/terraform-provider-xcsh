---
page_title: "xcsh_policer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policer reference."
---

# xcsh_policer reference

<a id="canonical-3320132331212333-3010302313330232-0221223202012002-1222203102211013-2121323323003011-1111232130221313-0133131033222112-3002101300233013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012220312113033-3230200322012212-1133230233322031-3223021213203330-2111033023130110-3312231221110122-2022012011231011-3221310001211200"></a>

## Property reference — Property reference / 132012211033 / 2

Breadcrumbs:

- [xcsh_policer](../data-sources/policer.md#canonical-3330032030013021-1312303333232323-1331010000013021-1132031333330232-2302330331211311-0310120132213301-2033021113022132-2232102121103000)
- Property reference

<a id="canonical-2021201303332030-3330213121003330-3330113000033232-2113200200200003-2323231322111100-3003131101100221-3133233120130300-0112232102110331"></a>

## Direct properties — Property reference / 132012211033 / 3

<a id="canonical-2011310023132220-0030303222211030-0011203200122033-0121212210200311-2211320222112033-0121132121031203-2202131000122000-1301323000023031"></a>

<a id="canonical-1021331220030002-1212303210121311-1321231120112230-3212333232010303-3201112100320011-2121110000000330-1331111032002221-1121103002123312"></a>

## annotations property — Property reference / 132012211033 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

<a id="canonical-2131022211113333-1220301331210102-2331223011220130-1311221232002121-1023332302012220-1022331130320101-1122323311301330-0331200322112120"></a>

<a id="canonical-0222002202222011-0301201133221020-3221003310031022-1331101201200020-0223322300032232-2103200103303231-3020212021202100-3321331213232323"></a>

## burst_size property — Property reference / 132012211033 / 5

Type: `"number"`. Computed.

The maximum size permitted for bursts of data. E.g. 10000 pps burst.

Upstream description:

The maximum size permitted for bursts of data. E.g. 10000 pps burst.

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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

<a id="canonical-1001130113220021-2331333323232331-2123333301103032-2210223232103210-1112011103320312-2311310310102022-0021001120220322-0230323320212112"></a>

<a id="canonical-3220133321330013-1130021231210003-0201131302200323-0012113132010133-1230031010312033-2033010300010303-2233213023300322-1030323120301000"></a>

## committed_information_rate property — Property reference / 132012211033 / 6

Type: `"number"`. Computed.

The committed information rate is the guaranteed packets rate for traffic arriving or departing
under normal conditions. E.g. 10000 pps.

Upstream description:

The committed information rate is the guaranteed packets rate for traffic arriving or departing
under normal conditions. E.g. 10000 pps.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10000000,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10000000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10000000"
  }
}
```

<a id="canonical-3023202311103232-1003330222311023-3031002020201310-2132233221012122-1302011132233220-1123001230322031-0223300312110221-1120333320223210"></a>

<a id="canonical-0332023023133133-3100113032000121-2233110330300030-0211220022232302-1013013333200330-3201232023002211-2130030122132003-3022133133001211"></a>

## description property — Property reference / 132012211033 / 7

Type: `"string"`. Computed.

Description of the Policer.

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

<a id="canonical-1330301131132312-0313333110231011-0333332031321131-3212002311302122-3121113101210030-2131203300110003-3001121301321123-2023210003123021"></a>

<a id="canonical-2012103332031133-0111323233102232-1031330233003021-1301310030220021-2113030213302330-0233030031112022-1222311301223103-3100131120133223"></a>

## ID property — Property reference / 132012211033 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1332332320200001-2220230310203323-0203332200002011-0301002301110013-2233121113302303-0321023010123121-3111123130102201-2122131111121110"></a>

<a id="canonical-3103133330002202-1001320303031220-0232101100003310-0311101331022332-2011133201333222-2123302211312123-3001221221201223-1011122200232310"></a>

## labels property — Property reference / 132012211033 / 9

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

<a id="canonical-1232312200330231-2030312012223001-0131203032123321-0302011331211011-2232013011310012-1211123033202130-2131112022133220-2313111331100303"></a>

<a id="canonical-2123103223202332-3020223002330003-3201023211320002-3021123020200100-0320123331220030-1233123011102312-0230333231033111-0023332000133122"></a>

## name property — Property reference / 132012211033 / 10

Type: `"string"`. Required.

Name of the Policer.

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

<a id="canonical-1320222110001233-3111022011331022-1020321322303203-1131003031110131-1120112001003132-3003301303203102-3230033221010121-0121213120132130"></a>

<a id="canonical-0211023123203333-1032323103021332-0030220202033211-0031001210312121-1113213312221202-0102011213332030-0033330001011111-0122123211110112"></a>

## namespace property — Property reference / 132012211033 / 11

Type: `"string"`. Required.

Namespace where the Policer exists.

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

<a id="canonical-3210201110221233-0113212122333123-3012102312033233-1132321012211011-1021302302223301-1132013030213122-1301332203210201-1311200211100230"></a>

<a id="canonical-1020333132033331-3222102221011101-3020213323212132-3320012020311303-0232021000231330-0021320032312033-3022110002102333-3132302321203121"></a>

## policer_mode property — Property reference / 132012211033 / 12

Type: `"string"`. Computed.

\[Enum: POLICER\_MODE\_NOT\_SHARED|POLICER\_MODE\_SHARED\] - POLICER\_MODE\_NOT\_SHARED: Not Shared
A separate policer instance is created for each reference to the policer - POLICER\_MODE\_SHARED:
Shared A common policer instance is used for for all references to the policer. Possible values are
\`POLICER\_MODE\_NOT\_SHARED\`, \`POLICER\_MODE\_SHARED\`. Defaults to
\`POLICER\_MODE\_NOT\_SHARED\`. Server applies default when omitted.

Upstream description:

&#8203;- POLICER\_MODE\_NOT\_SHARED: Not Shared

A separate policer instance is created for each reference to the policer &#8203;-
POLICER\_MODE\_SHARED: Shared

A common policer instance is used for for all references to the policer.

Receipt-pinned upstream constraints:

```json
{
  "default": "POLICER_MODE_NOT_SHARED",
  "enum": [
    "POLICER_MODE_NOT_SHARED",
    "POLICER_MODE_SHARED"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0201301023103122-0303031200210123-3313111020000310-2120032120000112-2211300330013022-2113003003121002-0101333223313302-2110000300030031"></a>

<a id="canonical-0221212120332323-1000132303123311-2303332213101122-3100010203123131-2320220103132122-1200032311021123-0021230120033331-2102112003232222"></a>

## policer_type property — Property reference / 132012211033 / 13

Type: `"string"`. Computed.

\[Enum: POLICER\_SINGLE\_RATE\_TWO\_COLOR\] Specifies the type of Policer Basic Single-Rate
Two-Color Policer. The only possible value is \`POLICER\_SINGLE\_RATE\_TWO\_COLOR\`. Defaults to
\`POLICER\_SINGLE\_RATE\_TWO\_COLOR\`. Server applies default when omitted.

Upstream description:

Specifies the type of Policer

Basic Single-Rate Two-Color Policer.

Receipt-pinned upstream constraints:

```json
{
  "default": "POLICER_SINGLE_RATE_TWO_COLOR",
  "enum": [
    "POLICER_SINGLE_RATE_TWO_COLOR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3310320002021331-1131313312323303-0210013332010221-0313000323133112-2021101312321132-3312321233203023-1023231121033000-3321110031332132"></a>

## All schema paths — Property reference / 132012211033 / 14

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--policer--reference--group-001.md#canonical-2011310023132220-0030303222211030-0011203200122033-0121212210200311-2211320222112033-0121132121031203-2202131000122000-1301323000023031) |
| `burst_size` | [burst_size](data-sources--policer--reference--group-001.md#canonical-2131022211113333-1220301331210102-2331223011220130-1311221232002121-1023332302012220-1022331130320101-1122323311301330-0331200322112120) |
| `committed_information_rate` | [committed_information_rate](data-sources--policer--reference--group-001.md#canonical-1001130113220021-2331333323232331-2123333301103032-2210223232103210-1112011103320312-2311310310102022-0021001120220322-0230323320212112) |
| `description` | [description](data-sources--policer--reference--group-001.md#canonical-3023202311103232-1003330222311023-3031002020201310-2132233221012122-1302011132233220-1123001230322031-0223300312110221-1120333320223210) |
| `id` | [ID](data-sources--policer--reference--group-001.md#canonical-1330301131132312-0313333110231011-0333332031321131-3212002311302122-3121113101210030-2131203300110003-3001121301321123-2023210003123021) |
| `labels` | [labels](data-sources--policer--reference--group-001.md#canonical-1332332320200001-2220230310203323-0203332200002011-0301002301110013-2233121113302303-0321023010123121-3111123130102201-2122131111121110) |
| `name` | [name](data-sources--policer--reference--group-001.md#canonical-1232312200330231-2030312012223001-0131203032123321-0302011331211011-2232013011310012-1211123033202130-2131112022133220-2313111331100303) |
| `namespace` | [namespace](data-sources--policer--reference--group-001.md#canonical-1320222110001233-3111022011331022-1020321322303203-1131003031110131-1120112001003132-3003301303203102-3230033221010121-0121213120132130) |
| `policer_mode` | [policer_mode](data-sources--policer--reference--group-001.md#canonical-3210201110221233-0113212122333123-3012102312033233-1132321012211011-1021302302223301-1132013030213122-1301332203210201-1311200211100230) |
| `policer_type` | [policer_type](data-sources--policer--reference--group-001.md#canonical-0201301023103122-0303031200210123-3313111020000310-2120032120000112-2211300330013022-2113003003121002-0101333223313302-2110000300030031) |

<a id="canonical-2012332000010300-0123231100000012-2310002012231123-2333202202330133-3300113101312303-0012021123011023-0110201030220020-0010233331312221"></a>

## Next pages — Property reference / 132012211033 / 15

- [xcsh_policer](../data-sources/policer.md#canonical-3330032030013021-1312303333232323-1331010000013021-1132031333330232-2302330331211311-0310120132213301-2033021113022132-2232102121103000)
