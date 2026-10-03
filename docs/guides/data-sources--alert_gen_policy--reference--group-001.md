---
page_title: "xcsh_alert_gen_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_gen_policy reference."
---

# xcsh_alert_gen_policy reference

<a id="canonical-0232203323003000-0030100311001232-1132202033111113-0322101230133330-0330113102210010-3303033232211100-3200031120003321-1033030031110112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331321022323110-3303333001201222-0232110303323213-3203221113213300-1003233022100202-2101313103102220-0221300120213020-3012200011211320"></a>

## Property reference — Property reference / 322200012212 / 2

Breadcrumbs:

- [xcsh_alert_gen_policy](../data-sources/alert_gen_policy.md#canonical-1212223311231031-2303130101102331-2121323302132120-3001003212230233-2233103321211211-1300022131031130-1101122301003332-3330023301222220)
- Property reference

<a id="canonical-3110110002302102-2322100210130112-2332122213130212-2020102023133100-0130002103131011-2201321312112131-3122002032312321-0233103023002333"></a>

## Direct properties — Property reference / 322200012212 / 3

<a id="canonical-1201320300231332-3123131110232110-2211232333012201-3213330222003033-0000220010301003-0023001333321232-2030213320133013-1310310021212203"></a>

<a id="canonical-2021202220111230-3230110023322202-2100010132111320-1302002110320321-1222020122020131-3320200013020001-0003002012201031-1133210323033203"></a>

## alert_status property — Property reference / 322200012212 / 4

Type: `"string"`. Computed.

\[Enum: ALERT\_ACTIVE|ALERT\_INACTIVE\] Alert Status. List of alert statuses Active Inactive.
Possible values are \`ALERT\_ACTIVE\`, \`ALERT\_INACTIVE\`. Defaults to \`ALERT\_ACTIVE\`.

Upstream description:

List of alert statuses

Active Inactive.

Receipt-pinned upstream constraints:

```json
{
  "default": "ALERT_ACTIVE",
  "enum": [
    "ALERT_ACTIVE",
    "ALERT_INACTIVE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1101230023130221-0212310210121220-1331101312331330-2102010330000312-0313330012310233-3101112001312102-3103033203113010-1100221132301110"></a>

<a id="canonical-3210301131101230-3003133303333120-0231131221120211-2101110012022332-1302220100110111-1313001020132022-0330123001102200-1020332020001232"></a>

## annotations property — Property reference / 322200012212 / 5

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

<a id="canonical-0310003123230332-0220130333333031-1321301310320232-2333133100321311-0200303013332112-1132303033100212-1220010301120330-1132120311211232"></a>

<a id="canonical-0013311103020002-0333012122322210-0001320300120200-1223303301311012-0200121013230322-3132011021011001-2031321222121221-0203111001002213"></a>

## description property — Property reference / 322200012212 / 6

Type: `"string"`. Computed.

Description of the AlertGenPolicy.

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

- [details](data-sources--alert_gen_policy--reference--group-001.md#canonical-2332002311203110-1011103103320020-2010133001222313-2112130031230013-3211100301131023-3010312121111300-1110110101120033-0323222303032110): complete subsection reference.

<a id="canonical-3100222102211022-2301132312013022-1212320032111232-0123002022321302-0100233012012133-0303120303032222-3222031103133030-3332011333221101"></a>

<a id="canonical-2033331322200011-3303112021122133-0022222221101212-0301330213022013-3022120113112202-0032323130201333-0021230312213301-1012032001100321"></a>

## ID property — Property reference / 322200012212 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0111333111231310-2110233010321130-0122010002132332-0132000302103102-0212320100123101-2110231033033032-1320132102013221-3213003010002012"></a>

<a id="canonical-3210001210230233-0213230122132301-2322120123311203-0002231103111311-3122232301222022-2212220230102123-2023121322001333-0100333233321102"></a>

## labels property — Property reference / 322200012212 / 8

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

<a id="canonical-0200021211023030-2102222020311102-3013002033330222-2032012012202211-2233000000122002-2113231221023321-1120123322011201-2223022333110312"></a>

<a id="canonical-3011100300232233-1023311212012023-0303331212121003-3030213331210312-0123333221323310-0022031013112021-1212310122321221-1122311303233030"></a>

## name property — Property reference / 322200012212 / 9

Type: `"string"`. Required.

Name of the AlertGenPolicy.

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

<a id="canonical-2022303133331222-2103223013031002-2200031213313122-2122302133220332-2011001323021211-1212000211000211-1220102033011033-0221100132001221"></a>

<a id="canonical-0022331111010123-2331002310333122-2213022002000231-2122322222101321-0022323233202230-1221003112100111-3222231221313223-2201001231232312"></a>

## namespace property — Property reference / 322200012212 / 10

Type: `"string"`. Required.

Namespace where the AlertGenPolicy exists.

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

<a id="canonical-2032213313303320-1312323131223223-3112103012221322-0101300200110203-2213321230302231-3300303202030002-2210002131023032-1313101211221000"></a>

## All schema paths — Property reference / 322200012212 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `alert_status` | [alert_status](data-sources--alert_gen_policy--reference--group-001.md#canonical-1201320300231332-3123131110232110-2211232333012201-3213330222003033-0000220010301003-0023001333321232-2030213320133013-1310310021212203) |
| `annotations` | [annotations](data-sources--alert_gen_policy--reference--group-001.md#canonical-1101230023130221-0212310210121220-1331101312331330-2102010330000312-0313330012310233-3101112001312102-3103033203113010-1100221132301110) |
| `description` | [description](data-sources--alert_gen_policy--reference--group-001.md#canonical-0310003123230332-0220130333333031-1321301310320232-2333133100321311-0200303013332112-1132303033100212-1220010301120330-1132120311211232) |
| `details` | [details](data-sources--alert_gen_policy--reference--group-001.md#canonical-2200103033222030-2210332322320110-2230313233031311-3200101132202120-1323310001002331-2230211100330022-3330201133313032-3210002303320120) |
| `details.alert_message` | [details.alert_message](data-sources--alert_gen_policy--reference--group-001.md#canonical-3001212021130132-3301023222223100-0310333002200102-2210002332210333-0320113330212320-3000323300023200-1032021013300221-0003131033002211) |
| `details.alert_message_details` | [details.alert_message_details](data-sources--alert_gen_policy--reference--group-001.md#canonical-3033013102202130-1322122201020023-0132112331131222-3312301102021332-2301002311220121-1110032303130120-0312133012321030-1030122022233202) |
| `details.alert_name` | [details.alert_name](data-sources--alert_gen_policy--reference--group-001.md#canonical-2001202021103032-3103333232102112-0123202331112331-3131101122220300-1031222121100333-2323023111332212-2300330123201002-2123333033112013) |
| `details.severity` | [details.severity](data-sources--alert_gen_policy--reference--group-001.md#canonical-2210212332033120-0002133331220313-1310212303233100-3003122001010223-3220310122323031-0220333332331313-3333321322031123-2210023020330200) |
| `id` | [ID](data-sources--alert_gen_policy--reference--group-001.md#canonical-3100222102211022-2301132312013022-1212320032111232-0123002022321302-0100233012012133-0303120303032222-3222031103133030-3332011333221101) |
| `labels` | [labels](data-sources--alert_gen_policy--reference--group-001.md#canonical-0111333111231310-2110233010321130-0122010002132332-0132000302103102-0212320100123101-2110231033033032-1320132102013221-3213003010002012) |
| `name` | [name](data-sources--alert_gen_policy--reference--group-001.md#canonical-0200021211023030-2102222020311102-3013002033330222-2032012012202211-2233000000122002-2113231221023321-1120123322011201-2223022333110312) |
| `namespace` | [namespace](data-sources--alert_gen_policy--reference--group-001.md#canonical-2022303133331222-2103223013031002-2200031213313122-2122302133220332-2011001323021211-1212000211000211-1220102033011033-0221100132001221) |

<a id="canonical-3212331330220132-0010230333222103-0232213212011223-3313202113012310-3013112100201323-2313232022122112-3312012102022211-3312212100323022"></a>

## Next pages — Property reference / 322200012212 / 12

- [details](data-sources--alert_gen_policy--reference--group-001.md#canonical-2332002311203110-1011103103320020-2010133001222313-2112130031230013-3211100301131023-3010312121111300-1110110101120033-0323222303032110)
- [xcsh_alert_gen_policy](../data-sources/alert_gen_policy.md#canonical-1212223311231031-2303130101102331-2121323302132120-3001003212230233-2233103321211211-1300022131031130-1101122301003332-3330023301222220)

<a id="canonical-2332002311203110-1011103103320020-2010133001222313-2112130031230013-3211100301131023-3010312121111300-1110110101120033-0323222303032110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112102332120000-3332330220200331-0201023102213012-0132212022013223-0111002213002101-3022313001100010-3332332023330330-3230323222111312"></a>

## details — details / 002230111230 / 2

Breadcrumbs:

- [xcsh_alert_gen_policy](../data-sources/alert_gen_policy.md#canonical-1212223311231031-2303130101102331-2121323302132120-3001003212230233-2233103321211211-1300022131031130-1101122301003332-3330023301222220)
- [Property reference](data-sources--alert_gen_policy--reference--group-001.md#canonical-0232203323003000-0030100311001232-1132202033111113-0322101230133330-0330113102210010-3303033232211100-3200031120003321-1033030031110112)
- details

<a id="canonical-2200103033222030-2210332322320110-2230313233031311-3200101132202120-1323310001002331-2230211100330022-3330201133313032-3210002303320120"></a>

Type: `"single"`. Computed.

Notification Details. Notification Details.

Upstream description:

Notification Details.

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

<a id="canonical-0013100121113103-3030012311233313-1022313123210331-2122333221011021-0011330003333212-1103203011230031-2100111112203133-0332303123311023"></a>

## Direct properties — details / 002230111230 / 3

<a id="canonical-3001212021130132-3301023222223100-0310333002200102-2210002332210333-0320113330212320-3000323300023200-1032021013300221-0003131033002211"></a>

<a id="canonical-1102212110232023-3333131303031331-3132112223132001-3220013303012113-2330312101302320-0203202200311010-1210332323133301-1033023023131321"></a>

## alert_message property — details / 002230111230 / 4

Type: `"string"`. Computed.

Alert Message. Alert Message.

Upstream description:

Alert Message.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-3033013102202130-1322122201020023-0132112331131222-3312301102021332-2301002311220121-1110032303130120-0312133012321030-1030122022233202"></a>

<a id="canonical-1323221322113321-1231202011122020-3033213000001332-1031130112221223-2110030203233200-3310321312330013-1022101310202010-1323302331022112"></a>

## alert_message_details property — details / 002230111230 / 5

Type: `"string"`. Computed.

Alert Message Details. Detailed message of the alert.

Upstream description:

Detailed message of the alert.

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
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2001202021103032-3103333232102112-0123202331112331-3131101122220300-1031222121100333-2323023111332212-2300330123201002-2123333033112013"></a>

<a id="canonical-0033303231320310-3013230021131213-1032002112000130-3030130330131001-1201321103002103-0023102103220221-3122013111233301-2112030011101310"></a>

## alert_name property — details / 002230111230 / 6

Type: `"string"`. Computed.

Alert Name. Alert Name.

Upstream description:

Alert Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 16,
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
    "ves.io.schema.rules.string.max_len": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "16"
  }
}
```

<a id="canonical-2210212332033120-0002133331220313-1310212303233100-3003122001010223-3220310122323031-0220333332331313-3333321322031123-2210023020330200"></a>

<a id="canonical-3313122130110310-3302203331233130-3033223311001223-1212312301201031-3032023133332300-0223332012331021-0232120000033032-3103332020020102"></a>

## severity property — details / 002230111230 / 7

Type: `"string"`. Computed.

\[Enum: MINOR|MAJOR|CRITICAL\] List of alert severities Minor Major Critical. Possible values are
\`MINOR\`, \`MAJOR\`, \`CRITICAL\`. Defaults to \`MINOR\`.

Upstream description:

List of alert severities

Minor Major Critical.

Receipt-pinned upstream constraints:

```json
{
  "default": "MINOR",
  "enum": [
    "MINOR",
    "MAJOR",
    "CRITICAL"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0303212110203003-0311000010322310-2321330230200333-1032302221121010-0323320212110231-3003033121122100-2333201002212112-1132012203033002"></a>

## Next pages — details / 002230111230 / 8

- [Property reference](data-sources--alert_gen_policy--reference--group-001.md#canonical-0232203323003000-0030100311001232-1132202033111113-0322101230133330-0330113102210010-3303033232211100-3200031120003321-1033030031110112)
- [xcsh_alert_gen_policy](../data-sources/alert_gen_policy.md#canonical-1212223311231031-2303130101102331-2121323302132120-3001003212230233-2233103321211211-1300022131031130-1101122301003332-3330023301222220)
