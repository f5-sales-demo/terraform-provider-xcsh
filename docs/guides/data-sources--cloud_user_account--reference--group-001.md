---
page_title: "xcsh_cloud_user_account reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_user_account reference."
---

# xcsh_cloud_user_account reference

<a id="canonical-2100003020103021-2211222133012112-2230221333311312-3001210132102132-0320331210120102-2221113213322121-0212112321020330-1011211032013131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-0121211201030010-3301133113320101-0332001123222101-1302203102030320-0332022000333203-1320110312000030-2303311122303231-1222130130102021)
- Property reference

<a id="canonical-0212200231101213-2020330101131021-1132022202102103-1120110213213302-1220320333033310-1033111310123013-1130222130300133-0310110233232200"></a>

### Direct properties for `xcsh_cloud_user_account`

<a id="canonical-1120233310013101-1030332220020301-1222013031201122-3133011110332220-2223102303221012-0322011032112111-3102300033333010-0021122320200221"></a>

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

- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-0102211100331230-1321103331002200-1232102100033030-2321023132121332-3120112320232012-1031323313003332-2202122131032201-3331030302331011): complete subsection reference.

<a id="canonical-3103001031022123-0233113201013323-2323123131202211-1120233311130023-1322100011130330-1100020233101201-1121210001201002-2332302011032003"></a>

<a id="canonical-2232320330210102-3233122310312102-2321033031302313-3013313301213230-2320222331230021-0321313332120300-2203031022311030-2022320021102333"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the CloudUserAccount.

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

<a id="canonical-0212022302002202-0203321010210303-1313003110331003-2203302332030023-3232300110232333-1301122302312320-2032232132222220-2333130301021031"></a>

<a id="canonical-3223320201210030-1313313211313022-1013100003022233-2113323033303021-2020012031120002-2102103010201223-2332000221233302-0112000131203122"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0213011103113311-3313020031110113-3121011010012210-2110233121232102-3012230211200000-0311202232021002-0210210202103200-2323310303312122"></a>

<a id="canonical-1211110300032220-3010202103023130-0030300210023003-1312203023123023-3232312121331021-3233121023003123-1123110313121313-0123312000113003"></a>

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

<a id="canonical-0302223333120320-1021121102232330-0001130020303120-3012310232320030-0013321231100223-1302211023001200-0300210000223331-0000212023212033"></a>

<a id="canonical-2333213131320011-0313133210302020-1011231212022200-1310321010212001-1011200131021032-0111000013332110-3333301010323120-0012131231111320"></a>

#### `name` property

Type: `"string"`. Required.

Name of the CloudUserAccount.

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

<a id="canonical-0232300231211112-1002110201212031-1111232332312331-0101101010301011-3112001132200122-2332030103313013-2332332213113310-3311212130200133"></a>

<a id="canonical-0110331213201030-3311302002103031-2113110112002301-0010200303120110-0120221110111200-3210232333033101-2133313300312012-1113003320102032"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the CloudUserAccount exists.

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

<a id="canonical-3103120231033313-3202212302011013-1103120313032132-0123131320323313-3212303102333100-0010122233202002-2220101312203201-0312231203103011"></a>

### All schema paths for `xcsh_cloud_user_account`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cloud_user_account--reference--group-001.md#canonical-1120233310013101-1030332220020301-1222013031201122-3133011110332220-2223102303221012-0322011032112111-3102300033333010-0021122320200221) |
| `aws_provider` | [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-1301113203000012-3123333011222313-3032130311011001-0131310221023312-1333203012022130-3013132132210221-2221121110222111-3012213120212031) |
| `aws_provider.aws_account_number` | [aws_provider.aws_account_number](data-sources--cloud_user_account--reference--group-001.md#canonical-3232112023330011-2303310031130333-2112100310301200-2103110020021200-0223320203221001-3202330010323023-1022231213212301-1112100123101031) |
| `aws_provider.aws_assume_role` | [aws_provider.aws_assume_role](data-sources--cloud_user_account--reference--group-001.md#canonical-3022102200000123-1012133121322331-3212322011222200-2223320021022323-0220330113310112-2023132033030000-3223023031120321-2220113203002213) |
| `aws_provider.aws_assume_role.custom_external_id` | [aws_provider.aws_assume_role.custom_external_id](data-sources--cloud_user_account--reference--group-001.md#canonical-0032210022311332-3110023203311001-0212203323131023-0202203321011213-3231011010101331-2221221232211301-2122230301333202-0112132200203211) |
| `aws_provider.aws_assume_role.duration_seconds` | [aws_provider.aws_assume_role.duration_seconds](data-sources--cloud_user_account--reference--group-001.md#canonical-2321103232103230-0301323103020000-2010223302121030-2232021112120110-3322321001000020-3133330220111102-3212203233021122-0230032321321210) |
| `aws_provider.aws_assume_role.external_id_is_optional` | [aws_provider.aws_assume_role.external_id_is_optional](data-sources--cloud_user_account--reference--group-001.md#canonical-0031320323210222-0132001203102210-3003021331011130-3333033213231221-2122012210033131-3213320220213303-1231012212111210-3201223133222132) |
| `aws_provider.aws_assume_role.external_id_is_tenant_id` | [aws_provider.aws_assume_role.external_id_is_tenant_id](data-sources--cloud_user_account--reference--group-001.md#canonical-2313033213111332-3333200112112133-1123033312001211-2212132313321110-3333222210002030-2123030223030001-3311202030211312-2033232010031201) |
| `aws_provider.aws_assume_role.role_arn` | [aws_provider.aws_assume_role.role_arn](data-sources--cloud_user_account--reference--group-001.md#canonical-0033311231130232-3202001333000233-0211221320320003-2113232011100233-3312132110232133-0111332023213120-2311232012121120-1010022020112031) |
| `aws_provider.aws_assume_role.session_name` | [aws_provider.aws_assume_role.session_name](data-sources--cloud_user_account--reference--group-001.md#canonical-0331110213301010-3302232131300020-2101211010202013-2032102133201300-1030100322212033-0111132121033203-2132000130011012-1323312110230211) |
| `aws_provider.aws_assume_role.session_tags` | [aws_provider.aws_assume_role.session_tags](data-sources--cloud_user_account--reference--group-001.md#canonical-2230110111003210-0301012331110313-3113322002121302-3303131313211213-2333211231113213-3322322321333010-1200303122322123-1021322110102023) |
| `aws_provider.aws_secret_key` | [aws_provider.aws_secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-3302330122301020-0301002212203200-0202200030302323-1021212230223320-0331001203111133-2310132202232132-1000312313201121-1002233213020321) |
| `aws_provider.aws_secret_key.access_key` | [aws_provider.aws_secret_key.access_key](data-sources--cloud_user_account--reference--group-001.md#canonical-3210322022021110-3113221023210122-2110031211010223-2230011201330131-0021232202122330-3320300213233232-3030122213303331-2322313230213002) |
| `aws_provider.aws_secret_key.secret_key` | [aws_provider.aws_secret_key.secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-1231122010313213-1020221021233332-0110230010111031-1223002330303123-2002322033121333-3300301011013032-3322213010031032-0103211212123301) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info](data-sources--cloud_user_account--reference--group-001.md#canonical-3130003222000321-2101302212003320-0000312133302231-2222123333110010-2033131332011302-3323210320203023-2032323201200212-1113301303211220) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.decryption_provider` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.decryption_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-2030322320222211-3002201110330131-2003320000012333-2130001121320333-0023221320330001-0102102200311220-1312120001010330-1133012210323301) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.location` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.location](data-sources--cloud_user_account--reference--group-001.md#canonical-1301311320013002-2111303001332310-0333020033110311-1333320123230323-3001132311333021-2123133232213132-2101323223210303-2330313303200220) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.store_provider` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.store_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-1100332313100122-2123322102222202-0010021013212030-2032201122000111-3331120000201131-3221312032112121-1021220002123133-0023000032110221) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info` | [aws_provider.aws_secret_key.secret_key.clear_secret_info](data-sources--cloud_user_account--reference--group-001.md#canonical-2330221202021112-3120131013123032-3120132020202100-2011321003000013-1012331123212013-2321202012201210-1100121000302200-2103302233230102) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info.provider_ref` | [aws_provider.aws_secret_key.secret_key.clear_secret_info.provider_ref](data-sources--cloud_user_account--reference--group-001.md#canonical-3211111021001220-1033130202223320-0100000022133223-2001303233310203-1101132211120221-0020213121010103-0112021301230030-1201120103330330) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info.url` | [aws_provider.aws_secret_key.secret_key.clear_secret_info.url](data-sources--cloud_user_account--reference--group-001.md#canonical-2030101221330320-1031322133302121-0201233011003331-2032023001013232-3301120230210220-1202021122210130-2122031330233310-0201022203023121) |
| `description` | [description](data-sources--cloud_user_account--reference--group-001.md#canonical-3103001031022123-0233113201013323-2323123131202211-1120233311130023-1322100011130330-1100020233101201-1121210001201002-2332302011032003) |
| `id` | [ID](data-sources--cloud_user_account--reference--group-001.md#canonical-0212022302002202-0203321010210303-1313003110331003-2203302332030023-3232300110232333-1301122302312320-2032232132222220-2333130301021031) |
| `labels` | [labels](data-sources--cloud_user_account--reference--group-001.md#canonical-0213011103113311-3313020031110113-3121011010012210-2110233121232102-3012230211200000-0311202232021002-0210210202103200-2323310303312122) |
| `name` | [name](data-sources--cloud_user_account--reference--group-001.md#canonical-0302223333120320-1021121102232330-0001130020303120-3012310232320030-0013321231100223-1302211023001200-0300210000223331-0000212023212033) |
| `namespace` | [namespace](data-sources--cloud_user_account--reference--group-001.md#canonical-0232300231211112-1002110201212031-1111232332312331-0101101010301011-3112001132200122-2332030103313013-2332332213113310-3311212130200133) |

<a id="canonical-0102211100331230-1321103331002200-1232102100033030-2321023132121332-3120112320232012-1031323313003332-2202122131032201-3331030302331011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_provider` properties

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-0121211201030010-3301133113320101-0332001123222101-1302203102030320-0332022000333203-1320110312000030-2303311122303231-1222130130102021)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-2100003020103021-2211222133012112-2230221333311312-3001210132102132-0320331210120102-2221113213322121-0212112321020330-1011211032013131)
- aws_provider

<a id="canonical-1301113203000012-3123333011222313-3032130311011001-0131310221023312-1333203012022130-3013132132210221-2221121110222111-3012213120212031"></a>

Type: `"single"`. Computed.

Configuration parameter for aws provider.

Additional upstream details:

Create AWS Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-aws_authentication_type": "[\"aws_assume_role\",\"aws_secret_key\"]"
}
```

<a id="canonical-1001302100002303-3303222211311001-0010021232132133-3011111100232301-2001131022230010-1101210213022133-3113102110103213-2300103121023002"></a>

### Direct properties for `aws_provider`

<a id="canonical-3232112023330011-2303310031130333-2112100310301200-2103110020021200-0223320203221001-3202330010323023-1022231213212301-1112100123101031"></a>

#### `aws_provider.aws_account_number` property

Type: `"string"`. Computed.

Account Number. 12 Digit Account Number.

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
    "ves.io.schema.rules.uint64.gte": "100000000000",
    "ves.io.schema.rules.uint64.lte": "999999999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "100000000000",
    "ves.io.schema.rules.uint64.lte": "999999999999"
  }
}
```

- [aws_assume_role](data-sources--cloud_user_account--reference--group-001.md#canonical-2220311311210332-2112023333100233-2311303110212122-2110301013321303-1131211211030133-0031333223321010-3331013110000201-2202113120120211): complete subsection reference.

- [aws_secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-3032231322023332-0223202202222220-3003320111200233-2000232003120003-3021113313232312-1212031313231100-0211021021232023-2213300022133122): complete subsection reference.

<a id="canonical-2220311311210332-2112023333100233-2311303110212122-2110301013321303-1131211211030133-0031333223321010-3331013110000201-2202113120120211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_provider.aws_assume_role` properties

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-0121211201030010-3301133113320101-0332001123222101-1302203102030320-0332022000333203-1320110312000030-2303311122303231-1222130130102021)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-2100003020103021-2211222133012112-2230221333311312-3001210132102132-0320331210120102-2221113213322121-0212112321020330-1011211032013131)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-0102211100331230-1321103331002200-1232102100033030-2321023132121332-3120112320232012-1031323313003332-2202122131032201-3331030302331011)
- aws_provider.aws_assume_role

<a id="canonical-3022102200000123-1012133121322331-3212322011222200-2223320021022323-0220330113310112-2023132033030000-3223023031120321-2220113203002213"></a>

Type: `"single"`. Computed.

AWS Assume Role to Handle Delegated Access.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-external_id": "[\"custom_external_id\",\"external_id_is_optional\",\"external_id_is_tenant_id\"]"
}
```

<a id="canonical-2320021323030033-0112013100222122-2310013321321111-1232301011332002-3300333001033020-3301133101110312-1023330133322222-3132212303002112"></a>

### Direct properties for `aws_provider.aws_assume_role`

<a id="canonical-0032210022311332-3110023203311001-0212203323131023-0202203321011213-3231011010101331-2221221232211301-2122230301333202-0112132200203211"></a>

#### `aws_provider.aws_assume_role.custom_external_id` property

Type: `"string"`. Computed.

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2"
  }
}
```

<a id="canonical-2321103232103230-0301323103020000-2010223302121030-2232021112120110-3322321001000020-3133330220111102-3212203233021122-0230032321321210"></a>

<a id="canonical-0310232311310120-3111330100003332-1002202302202010-3103203002223331-3020001322211030-0233322330003332-1032230320010002-2130313112112033"></a>

#### `aws_provider.aws_assume_role.duration_seconds` property

Type: `"number"`. Computed.

The duration, in seconds of the role session.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 43200,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 3600
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "43200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "43200"
  }
}
```

- [external_id_is_optional](data-sources--cloud_user_account--reference--group-001.md#canonical-3303301011330122-3302232103013003-0102211012221231-1012221113313231-1213333032013122-1101113310222120-3001022122032023-1201000220213301): complete subsection reference.

- [external_id_is_tenant_id](data-sources--cloud_user_account--reference--group-001.md#canonical-0232203312231103-1203200130102001-3223202131300132-3202310113303031-0112032313002002-2321303203301201-1313001133102122-3211130330101201): complete subsection reference.

<a id="canonical-0033311231130232-3202001333000233-0211221320320003-2113232011100233-3312132110232133-0111332023213120-2311232012121120-1010022020112031"></a>

<a id="canonical-2123031113130031-1020323023302220-0322133201210333-0201013110333031-1311232233331211-0033113110133330-2221213320133020-2323321100021312"></a>

#### `aws_provider.aws_assume_role.role_arn` property

Type: `"string"`. Computed.

IAM Role ARN. IAM Role ARN to assume the role.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 20,
    "pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "20",
    "ves.io.schema.rules.string.pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "20",
    "ves.io.schema.rules.string.pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  }
}
```

<a id="canonical-0331110213301010-3302232131300020-2101211010202013-2032102133201300-1030100322212033-0111132121033203-2132000130011012-1323312110230211"></a>

<a id="canonical-2111233001312220-2322322120312301-1013331200230203-0011320120131112-0310332231222312-0322323022321013-3133330302111333-2002322313010110"></a>

#### `aws_provider.aws_assume_role.session_name` property

Type: `"string"`. Computed.

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 2,
    "pattern": "[\\\\w+=,.@-]*"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2",
    "ves.io.schema.rules.string.pattern": "[\\\\w+=,.@-]*"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2",
    "ves.io.schema.rules.string.pattern": "[\\\\w+=,.@-]*"
  }
}
```

<a id="canonical-2230110111003210-0301012331110313-3113322002121302-3303131313211213-2333211231113213-3322322321333010-1200303122322123-1021322110102023"></a>

<a id="canonical-2020222112220323-1132331103131310-2003122122323103-3133113212011102-3132200021002103-1023320213122011-1010031301210233-2330230011310203"></a>

#### `aws_provider.aws_assume_role.session_tags` property

Type: `["map", "string"]`. Computed.

Session tags are key-value pair attributes that you pass when you assume an IAM role.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 40
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 127,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "127",
      "ves.io.schema.rules.map.max_pairs": "40",
      "ves.io.schema.rules.map.values.string.max_len": "255"
    },
    "values": {
      "maxLength": 255,
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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

<a id="canonical-3303301011330122-3302232103013003-0102211012221231-1012221113313231-1213333032013122-1101113310222120-3001022122032023-1201000220213301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_provider.aws_assume_role.external_id_is_optional` properties

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-0121211201030010-3301133113320101-0332001123222101-1302203102030320-0332022000333203-1320110312000030-2303311122303231-1222130130102021)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-2100003020103021-2211222133012112-2230221333311312-3001210132102132-0320331210120102-2221113213322121-0212112321020330-1011211032013131)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-0102211100331230-1321103331002200-1232102100033030-2321023132121332-3120112320232012-1031323313003332-2202122131032201-3331030302331011)
- [aws_provider.aws_assume_role](data-sources--cloud_user_account--reference--group-001.md#canonical-2220311311210332-2112023333100233-2311303110212122-2110301013321303-1131211211030133-0031333223321010-3331013110000201-2202113120120211)
- aws_provider.aws_assume_role.external_id_is_optional

<a id="canonical-0031320323210222-0132001203102210-3003021331011130-3333033213231221-2122012210033131-3213320220213303-1231012212111210-3201223133222132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for external ID is optional.

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

<a id="canonical-0232203312231103-1203200130102001-3223202131300132-3202310113303031-0112032313002002-2321303203301201-1313001133102122-3211130330101201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_provider.aws_assume_role.external_id_is_tenant_id` properties

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-0121211201030010-3301133113320101-0332001123222101-1302203102030320-0332022000333203-1320110312000030-2303311122303231-1222130130102021)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-2100003020103021-2211222133012112-2230221333311312-3001210132102132-0320331210120102-2221113213322121-0212112321020330-1011211032013131)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-0102211100331230-1321103331002200-1232102100033030-2321023132121332-3120112320232012-1031323313003332-2202122131032201-3331030302331011)
- [aws_provider.aws_assume_role](data-sources--cloud_user_account--reference--group-001.md#canonical-2220311311210332-2112023333100233-2311303110212122-2110301013321303-1131211211030133-0031333223321010-3331013110000201-2202113120120211)
- aws_provider.aws_assume_role.external_id_is_tenant_id

<a id="canonical-2313033213111332-3333200112112133-1123033312001211-2212132313321110-3333222210002030-2123030223030001-3311202030211312-2033232010031201"></a>

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

<a id="canonical-3032231322023332-0223202202222220-3003320111200233-2000232003120003-3021113313232312-1212031313231100-0211021021232023-2213300022133122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_provider.aws_secret_key` properties

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-0121211201030010-3301133113320101-0332001123222101-1302203102030320-0332022000333203-1320110312000030-2303311122303231-1222130130102021)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-2100003020103021-2211222133012112-2230221333311312-3001210132102132-0320331210120102-2221113213322121-0212112321020330-1011211032013131)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-0102211100331230-1321103331002200-1232102100033030-2321023132121332-3120112320232012-1031323313003332-2202122131032201-3331030302331011)
- aws_provider.aws_secret_key

<a id="canonical-3302330122301020-0301002212203200-0202200030302323-1021212230223320-0331001203111133-2310132202232132-1000312313201121-1002233213020321"></a>

Type: `"single"`. Computed.

AWS Programmatic Access Credentials type.

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

<a id="canonical-0231231010112231-3132320012000233-3003221212312311-1312203313222101-1223001112213310-1330213223223101-1130311100310331-2110100033231320"></a>

### Direct properties for `aws_provider.aws_secret_key`

<a id="canonical-3210322022021110-3113221023210122-2110031211010223-2230011201330131-0021232202122330-3320300213233232-3030122213303331-2322313230213002"></a>

#### `aws_provider.aws_secret_key.access_key` property

Type: `"string"`. Computed.

Access Key ID. Access key ID for your AWS account.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-2201123032031030-3313132200301120-3003330322021010-0313203203031020-1133033322002321-3133033310333203-1103133012332130-0212002320112023): complete subsection reference.

<a id="canonical-2201123032031030-3313132200301120-3003330322021010-0313203203031020-1133033322002321-3133033310333203-1103133012332130-0212002320112023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_provider.aws_secret_key.secret_key` properties

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-0121211201030010-3301133113320101-0332001123222101-1302203102030320-0332022000333203-1320110312000030-2303311122303231-1222130130102021)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-2100003020103021-2211222133012112-2230221333311312-3001210132102132-0320331210120102-2221113213322121-0212112321020330-1011211032013131)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-0102211100331230-1321103331002200-1232102100033030-2321023132121332-3120112320232012-1031323313003332-2202122131032201-3331030302331011)
- [aws_provider.aws_secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-3032231322023332-0223202202222220-3003320111200233-2000232003120003-3021113313232312-1212031313231100-0211021021232023-2213300022133122)
- aws_provider.aws_secret_key.secret_key

<a id="canonical-1231122010313213-1020221021233332-0110230010111031-1223002330303123-2002322033121333-3300301011013032-3322213010031032-0103211212123301"></a>

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

<a id="canonical-1002122210213031-2330011033321333-0110221112312332-2122013121313313-2130300023313010-3103322112003031-0130032232302332-1030211010302133"></a>

### Direct properties for `aws_provider.aws_secret_key.secret_key`

- [blindfold_secret_info](data-sources--cloud_user_account--reference--group-001.md#canonical-1211102001222132-0200202012121113-2212011211322031-1113023322232200-0033233131323210-0010220302322310-3201212301012323-0121322012101112): complete subsection reference.

- [clear_secret_info](data-sources--cloud_user_account--reference--group-001.md#canonical-3013012200302331-1310101220122301-0303010001103030-3010323111332130-0330032313231320-3022121201001033-0111013303312311-1010231211221332): complete subsection reference.

<a id="canonical-1211102001222132-0200202012121113-2212011211322031-1113023322232200-0033233131323210-0010220302322310-3201212301012323-0121322012101112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_provider.aws_secret_key.secret_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-0121211201030010-3301133113320101-0332001123222101-1302203102030320-0332022000333203-1320110312000030-2303311122303231-1222130130102021)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-2100003020103021-2211222133012112-2230221333311312-3001210132102132-0320331210120102-2221113213322121-0212112321020330-1011211032013131)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-0102211100331230-1321103331002200-1232102100033030-2321023132121332-3120112320232012-1031323313003332-2202122131032201-3331030302331011)
- [aws_provider.aws_secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-3032231322023332-0223202202222220-3003320111200233-2000232003120003-3021113313232312-1212031313231100-0211021021232023-2213300022133122)
- [aws_provider.aws_secret_key.secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-2201123032031030-3313132200301120-3003330322021010-0313203203031020-1133033322002321-3133033310333203-1103133012332130-0212002320112023)
- aws_provider.aws_secret_key.secret_key.blindfold_secret_info

<a id="canonical-3130003222000321-2101302212003320-0000312133302231-2222123333110010-2033131332011302-3323210320203023-2032323201200212-1113301303211220"></a>

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

<a id="canonical-3211110331000013-3333013300312020-3013322122201231-0333030200012223-3301310210222332-2012103303303321-2010220223102332-1121102011222302"></a>

### Direct properties for `aws_provider.aws_secret_key.secret_key.blindfold_secret_info`

<a id="canonical-2030322320222211-3002201110330131-2003320000012333-2130001121320333-0023221320330001-0102102200311220-1312120001010330-1133012210323301"></a>

#### `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1301311320013002-2111303001332310-0333020033110311-1333320123230323-3001132311333021-2123133232213132-2101323223210303-2330313303200220"></a>

<a id="canonical-1021333233112020-0133012331030031-1122133212111320-2222110323031300-1320331213203200-0100323203210111-1222333221310132-3033313302010132"></a>

#### `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1100332313100122-2123322102222202-0010021013212030-2032201122000111-3331120000201131-3221312032112121-1021220002123133-0023000032110221"></a>

<a id="canonical-3112213320310111-2110331102222131-3012302000330203-2112323030213320-2330000321133300-2011331201202232-0123222021301100-2022113102323230"></a>

#### `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3013012200302331-1310101220122301-0303010001103030-3010323111332130-0330032313231320-3022121201001033-0111013303312311-1010231211221332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_provider.aws_secret_key.secret_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-0121211201030010-3301133113320101-0332001123222101-1302203102030320-0332022000333203-1320110312000030-2303311122303231-1222130130102021)
- [Property reference](data-sources--cloud_user_account--reference--group-001.md#canonical-2100003020103021-2211222133012112-2230221333311312-3001210132102132-0320331210120102-2221113213322121-0212112321020330-1011211032013131)
- [aws_provider](data-sources--cloud_user_account--reference--group-001.md#canonical-0102211100331230-1321103331002200-1232102100033030-2321023132121332-3120112320232012-1031323313003332-2202122131032201-3331030302331011)
- [aws_provider.aws_secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-3032231322023332-0223202202222220-3003320111200233-2000232003120003-3021113313232312-1212031313231100-0211021021232023-2213300022133122)
- [aws_provider.aws_secret_key.secret_key](data-sources--cloud_user_account--reference--group-001.md#canonical-2201123032031030-3313132200301120-3003330322021010-0313203203031020-1133033322002321-3133033310333203-1103133012332130-0212002320112023)
- aws_provider.aws_secret_key.secret_key.clear_secret_info

<a id="canonical-2330221202021112-3120131013123032-3120132020202100-2011321003000013-1012331123212013-2321202012201210-1100121000302200-2103302233230102"></a>

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

<a id="canonical-2132112001022020-0320030333131222-3300213132320232-3033011103013021-3003201131221133-2102210012131111-2100333123200121-3222323021013301"></a>

### Direct properties for `aws_provider.aws_secret_key.secret_key.clear_secret_info`

<a id="canonical-3211111021001220-1033130202223320-0100000022133223-2001303233310203-1101132211120221-0020213121010103-0112021301230030-1201120103330330"></a>

#### `aws_provider.aws_secret_key.secret_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2030101221330320-1031322133302121-0201233011003331-2032023001013232-3301120230210220-1202021122210130-2122031330233310-0201022203023121"></a>

<a id="canonical-1011100223120221-1232112023013023-3031213223001002-1311122320021131-1210223033333322-1200133313331301-3303230020201031-0012231233110011"></a>

#### `aws_provider.aws_secret_key.secret_key.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
