---
page_title: "xcsh_cminstance reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cminstance reference."
---

# xcsh_cminstance reference

<a id="canonical-0122030300032322-3101313310110231-2201200223333032-1111223213311301-3211302322021031-2210200012133330-3303132132031200-1011331323321121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-3030301003002223-1121312121232112-1232303010012000-0213330232003222-0030121133101123-1130200110322121-0013030301203132-0203031312323211)
- Property reference

<a id="canonical-1110033030010202-3133121120023313-2122221021320202-0212200230020130-3210212313020100-3122130310233001-1223000210130032-3101113002311111"></a>

### Direct properties for `xcsh_cminstance`

<a id="canonical-3002220320321201-0111231333103000-2100101231202311-2312130200032023-2023203333001103-1022201332332201-3320032212110213-0032201222030232"></a>

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

- [api_token](resources--cminstance--reference--group-001.md#canonical-3222030313100232-0200020221122131-3103013213301233-2303233102233002-1131233212323133-1112110023232123-2100323101300011-1201030210101112): complete subsection reference.

<a id="canonical-3330101230003223-1230021131301010-0331220123002223-0301222101102312-1001311222232033-2120302000111013-2233211232021030-0302323131110201"></a>

<a id="canonical-0232010231302223-1103333333131110-2000130130030120-3030110020002323-0000201302200202-0100112233132121-3023223020212201-2020311120312112"></a>

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

<a id="canonical-3311201123111302-2301312021321323-0222321130101131-0102010132302200-2032002232333102-1121120033302320-1000012030202132-3020030332121301"></a>

<a id="canonical-3313331020032200-1200303021101203-1232322131020130-0100202032123000-0033021112221313-1331112210332013-2300012231321021-2210133231210211"></a>

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

<a id="canonical-2133023103312033-2301030232102231-1113113023001032-3011310220102313-3230232001122202-0001132022300223-3220213321302321-2220112232010332"></a>

<a id="canonical-3112231221131223-2300302232122110-3300310220202332-1220020110102010-2021123102321302-2133121222310231-2210212210123311-1030020011101221"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip](resources--cminstance--reference--group-001.md#canonical-1002002022300323-3132232220002130-1211230032132212-1000211102313130-2110020102001122-2320323121310312-3201200233102200-0011203231010231): complete subsection reference.

<a id="canonical-2233132312331102-1301003111122113-0231003211200121-0220332130102201-0013231132122133-0220021211201021-1221212002332111-3203311133111231"></a>

<a id="canonical-1213012323003331-0100320333130313-3303030113300100-1302313313032201-1031322202223010-2031322323332302-0102111113200221-2200220222311200"></a>

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

<a id="canonical-0003301102213133-2101133212231200-2230022103230012-2223002030333032-1200333012111022-1100131132220120-1010133103121310-1233330013233111"></a>

<a id="canonical-0332103211001003-0202232220310000-3300002221232330-2011111130330300-3000131102031311-0313233113302031-0031320300101122-1020100011110210"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Cminstance. Must be unique within the namespace.

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

<a id="canonical-1102311321010113-1123301311003011-3011230220033133-3323032211030202-2110011330121123-0230203300101220-1301021020001313-1112222103233020"></a>

<a id="canonical-0220120131231332-2132320213231223-1231321333002210-3202222010011112-2223322221001332-1201322221201013-3021230303231213-1100202031323300"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Cminstance is created.

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

- [password](resources--cminstance--reference--group-001.md#canonical-2133031010321330-0110012311313220-1021213332102332-1032201111113230-0123321302231123-0130032231323311-3003311121000102-2321032213111120): complete subsection reference.

<a id="canonical-2220120011203301-0010013311320122-2333313111113101-1222222233112133-0032333011012011-0211133132300310-2101110331013103-0100231232312032"></a>

<a id="canonical-1033321002300333-1021210100333321-0013131221313303-1132122021103313-0130313222123330-1032001102132101-2322301220100231-1113103233120010"></a>

#### `port` property

Type: `"number"`. Required.

Port of the Central Manager instance to connect to.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [timeouts](resources--cminstance--reference--group-001.md#canonical-3012012022112030-3310230301133232-1001311133212011-3332112331333221-3032000313020313-0020131223101003-2012031130303010-3030332222211210): complete subsection reference.

<a id="canonical-3220103122321210-3021013112222102-2331132201320221-3210200223321101-3312211123200311-2113300121320212-1331303012123002-2301221132230101"></a>

<a id="canonical-2301033302033131-3202212203003121-0002010323300032-2023210020021320-2012032033102102-0223012102110311-3333210013231333-0011123211021023"></a>

#### `username` property

Type: `"string"`. Required.

Username for the Central Manager instance.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 4,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "4"
  }
}
```

<a id="canonical-0303032031301113-2203100002003012-3231312233113102-3010121103000213-1110023113113300-0000333231111202-1100110032013112-1013023302131322"></a>

### All schema paths for `xcsh_cminstance`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cminstance--reference--group-001.md#canonical-3002220320321201-0111231333103000-2100101231202311-2312130200032023-2023203333001103-1022201332332201-3320032212110213-0032201222030232) |
| `api_token` | [api_token](resources--cminstance--reference--group-001.md#canonical-1313230022221312-2320101310333121-3023113313230312-0120333003310300-1331200333000133-0200012003102211-1030230200312123-2103032130222311) |
| `api_token.blindfold_secret_info` | [api_token.blindfold_secret_info](resources--cminstance--reference--group-001.md#canonical-2131013303010131-3112311200331020-3310131201100222-1310022203100323-2021313013013221-0302101333101032-1212023110122103-2002331130333332) |
| `api_token.blindfold_secret_info.decryption_provider` | [api_token.blindfold_secret_info.decryption_provider](resources--cminstance--reference--group-001.md#canonical-2332223021211132-2103011003222313-2301230200121010-3033021331312131-1323103201131123-3021201201112321-1100032131310311-1112010001133213) |
| `api_token.blindfold_secret_info.location` | [api_token.blindfold_secret_info.location](resources--cminstance--reference--group-001.md#canonical-0002202100131103-3222301102102310-1030301003013232-2100023201011203-1111103003133130-2210322021200000-3022321211221011-3233123133201032) |
| `api_token.blindfold_secret_info.store_provider` | [api_token.blindfold_secret_info.store_provider](resources--cminstance--reference--group-001.md#canonical-1020201312110200-0011001232111301-1123312110112333-0313130230200000-3322023013121022-1322133031203313-0033120121222302-3301200022133320) |
| `api_token.clear_secret_info` | [api_token.clear_secret_info](resources--cminstance--reference--group-001.md#canonical-0021230020231311-3213201031030333-3222131321212311-2101122032131032-3311230313201203-0332031111032210-0232300303133113-3332113130210323) |
| `api_token.clear_secret_info.provider_ref` | [api_token.clear_secret_info.provider_ref](resources--cminstance--reference--group-001.md#canonical-0203002223222012-1223201302033203-1022010130230021-1210223121023113-0313012313100031-3020211100313132-2101020133213123-2210103121031020) |
| `api_token.clear_secret_info.url` | [api_token.clear_secret_info.url](resources--cminstance--reference--group-001.md#canonical-3103313121212133-1232130300223231-2032131102330331-2121300033111330-0331102200211121-2113330300330300-3001201012211101-0322011013012311) |
| `description` | [description](resources--cminstance--reference--group-001.md#canonical-3330101230003223-1230021131301010-0331220123002223-0301222101102312-1001311222232033-2120302000111013-2233211232021030-0302323131110201) |
| `disable` | [disable](resources--cminstance--reference--group-001.md#canonical-3311201123111302-2301312021321323-0222321130101131-0102010132302200-2032002232333102-1121120033302320-1000012030202132-3020030332121301) |
| `id` | [ID](resources--cminstance--reference--group-001.md#canonical-2133023103312033-2301030232102231-1113113023001032-3011310220102313-3230232001122202-0001132022300223-3220213321302321-2220112232010332) |
| `ip` | [ip](resources--cminstance--reference--group-001.md#canonical-0200323302101033-3000231232020123-2220103000201020-3300102210300211-0312301312001211-1310132003011330-0213020101033123-0211121130112220) |
| `ip.addr` | [ip.addr](resources--cminstance--reference--group-001.md#canonical-1331123221230203-0033303213331222-0100030002110013-1033232021213220-1300202120002012-3231103321012000-1313302130102232-3201012330100001) |
| `labels` | [labels](resources--cminstance--reference--group-001.md#canonical-2233132312331102-1301003111122113-0231003211200121-0220332130102201-0013231132122133-0220021211201021-1221212002332111-3203311133111231) |
| `name` | [name](resources--cminstance--reference--group-001.md#canonical-0003301102213133-2101133212231200-2230022103230012-2223002030333032-1200333012111022-1100131132220120-1010133103121310-1233330013233111) |
| `namespace` | [namespace](resources--cminstance--reference--group-001.md#canonical-1102311321010113-1123301311003011-3011230220033133-3323032211030202-2110011330121123-0230203300101220-1301021020001313-1112222103233020) |
| `password` | [password](resources--cminstance--reference--group-001.md#canonical-1331031232030132-0300332322012130-3331100323320213-3220113112222123-3221212302223000-3212202013021233-3330330010032301-1312113100130133) |
| `password.blindfold_secret_info` | [password.blindfold_secret_info](resources--cminstance--reference--group-001.md#canonical-2300220322233132-2030221202312032-2322303203321132-0301132123022123-0213101110120233-1231230121303330-1122002331033022-1302231003220020) |
| `password.blindfold_secret_info.decryption_provider` | [password.blindfold_secret_info.decryption_provider](resources--cminstance--reference--group-001.md#canonical-1132031112222120-0010221020110330-1000311020200021-1013323002131202-1031301303113110-3302301301321133-0133303112333213-1210311131102121) |
| `password.blindfold_secret_info.location` | [password.blindfold_secret_info.location](resources--cminstance--reference--group-001.md#canonical-0323110033202211-1103223130231100-0212032013122330-2012122322312230-2222312011221211-3010123020013120-0010210210322232-3102111021231322) |
| `password.blindfold_secret_info.store_provider` | [password.blindfold_secret_info.store_provider](resources--cminstance--reference--group-001.md#canonical-1221302320311102-1322003213223313-2311213222222023-1121013032101212-3200323013310101-2212311311131123-3331122223200203-0133202112203213) |
| `password.clear_secret_info` | [password.clear_secret_info](resources--cminstance--reference--group-001.md#canonical-0212031133121131-3133033321311331-3311100203130311-3112232231002023-2231220023223203-2011331020032221-0202220001330012-2310221123123113) |
| `password.clear_secret_info.provider_ref` | [password.clear_secret_info.provider_ref](resources--cminstance--reference--group-001.md#canonical-3122333001001122-2132102302221301-2230233333332032-1102100113133012-3123300120200222-2213201311232200-3001222110200230-2201212102131021) |
| `password.clear_secret_info.url` | [password.clear_secret_info.url](resources--cminstance--reference--group-001.md#canonical-1310111230010330-2013001221013313-3321332223100021-0332323321023031-3233013220210031-1030001111111201-1220212221232113-0011201231132003) |
| `port` | [port](resources--cminstance--reference--group-001.md#canonical-2220120011203301-0010013311320122-2333313111113101-1222222233112133-0032333011012011-0211133132300310-2101110331013103-0100231232312032) |
| `timeouts` | [timeouts](resources--cminstance--reference--group-001.md#canonical-1031300310232103-2212332332301212-0231132112011000-0231213021121031-2003323202020011-3131311101101210-2202311103000212-1011303123021031) |
| `timeouts.create` | [timeouts.create](resources--cminstance--reference--group-001.md#canonical-3330021233132032-3233101313301221-3131021322121101-2103332233000200-0333220012302131-2333310113202313-0103010200211320-3012100211101222) |
| `timeouts.delete` | [timeouts.delete](resources--cminstance--reference--group-001.md#canonical-2011210022023333-2300320310123321-1123213333203022-1310223330113310-3023303013200010-3232112333320203-3332112131303111-1210010012133221) |
| `timeouts.read` | [timeouts.read](resources--cminstance--reference--group-001.md#canonical-2131000010312030-2023113202332122-1332002202121010-3300231023330300-0011221110233232-1130311012112132-1321333303220201-0012213121101223) |
| `timeouts.update` | [timeouts.update](resources--cminstance--reference--group-001.md#canonical-2130013000130110-3030132210112131-3310211331110221-1131311113100323-2231003212023313-1002011220223320-0233203113110123-3110321133220121) |
| `username` | [username](resources--cminstance--reference--group-001.md#canonical-3220103122321210-3021013112222102-2331132201320221-3210200223321101-3312211123200311-2113300121320212-1331303012123002-2301221132230101) |

<a id="canonical-3222030313100232-0200020221122131-3103013213301233-2303233102233002-1131233212323133-1112110023232123-2100323101300011-1201030210101112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_token` properties

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-3030301003002223-1121312121232112-1232303010012000-0213330232003222-0030121133101123-1130200110322121-0013030301203132-0203031312323211)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-0122030300032322-3101313310110231-2201200223333032-1111223213311301-3211302322021031-2210200012133330-3303132132031200-1011331323321121)
- api_token

<a id="canonical-1313230022221312-2320101310333121-3023113313230312-0120333003310300-1331200333000133-0200012003102211-1030230200312123-2103032130222311"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
api_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313113302021120-1222002101313020-2233120311333311-2002010123231203-3211022130123101-2222320022310023-0021133302330020-0000013333323123"></a>

### Direct properties for `api_token`

- [blindfold_secret_info](resources--cminstance--reference--group-001.md#canonical-0122133031100330-0220300330030032-2212101130122033-0231102033113010-0113210121323012-2021331023301333-3222302223103012-0322220023223223): complete subsection reference.

- [clear_secret_info](resources--cminstance--reference--group-001.md#canonical-1233220322220122-0220202202101221-3300121202220003-2322022332111122-2321321313331303-2001210000011203-0312100330023102-3300101320032202): complete subsection reference.

<a id="canonical-0122133031100330-0220300330030032-2212101130122033-0231102033113010-0113210121323012-2021331023301333-3222302223103012-0322220023223223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-3030301003002223-1121312121232112-1232303010012000-0213330232003222-0030121133101123-1130200110322121-0013030301203132-0203031312323211)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-0122030300032322-3101313310110231-2201200223333032-1111223213311301-3211302322021031-2210200012133330-3303132132031200-1011331323321121)
- [api_token](resources--cminstance--reference--group-001.md#canonical-3222030313100232-0200020221122131-3103013213301233-2303233102233002-1131233212323133-1112110023232123-2100323101300011-1201030210101112)
- api_token.blindfold_secret_info

<a id="canonical-2131013303010131-3112311200331020-3310131201100222-1310022203100323-2021313013013221-0302101333101032-1212023110122103-2002331130333332"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003122102030100-0313200023311021-1101103201011003-3212130231012012-1312033123022210-1210320133023110-0032023220201112-3102133213121002"></a>

### Direct properties for `api_token.blindfold_secret_info`

<a id="canonical-2332223021211132-2103011003222313-2301230200121010-3033021331312131-1323103201131123-3021201201112321-1100032131310311-1112010001133213"></a>

#### `api_token.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0002202100131103-3222301102102310-1030301003013232-2100023201011203-1111103003133130-2210322021200000-3022321211221011-3233123133201032"></a>

<a id="canonical-3302232222230132-0213330232313230-0331301120000301-0123223220121230-3131202022010103-2030331302200123-2110313332012230-3133022003223310"></a>

#### `api_token.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1020201312110200-0011001232111301-1123312110112333-0313130230200000-3322023013121022-1322133031203313-0033120121222302-3301200022133320"></a>

<a id="canonical-2302323223100333-0313000121123101-2133013331020033-1120221310303311-1100133131000310-1000023301312301-1301320022201032-1120023013223103"></a>

#### `api_token.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1233220322220122-0220202202101221-3300121202220003-2322022332111122-2321321313331303-2001210000011203-0312100330023102-3300101320032202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-3030301003002223-1121312121232112-1232303010012000-0213330232003222-0030121133101123-1130200110322121-0013030301203132-0203031312323211)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-0122030300032322-3101313310110231-2201200223333032-1111223213311301-3211302322021031-2210200012133330-3303132132031200-1011331323321121)
- [api_token](resources--cminstance--reference--group-001.md#canonical-3222030313100232-0200020221122131-3103013213301233-2303233102233002-1131233212323133-1112110023232123-2100323101300011-1201030210101112)
- api_token.clear_secret_info

<a id="canonical-0021230020231311-3213201031030333-3222131321212311-2101122032131032-3311230313201203-0332031111032210-0232300303133113-3332113130210323"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110000113302003-0003022323301132-3011103330231122-1223302213133110-1022011031333102-0201322102200313-0211323203322200-3123122302110013"></a>

### Direct properties for `api_token.clear_secret_info`

<a id="canonical-0203002223222012-1223201302033203-1022010130230021-1210223121023113-0313012313100031-3020211100313132-2101020133213123-2210103121031020"></a>

#### `api_token.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3103313121212133-1232130300223231-2032131102330331-2121300033111330-0331102200211121-2113330300330300-3001201012211101-0322011013012311"></a>

<a id="canonical-2302011201221213-1210231232110111-2100001323101220-3312303220210210-3230013223310231-1311223113311023-3301302313203101-0131212202011323"></a>

#### `api_token.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1002002022300323-3132232220002130-1211230032132212-1000211102313130-2110020102001122-2320323121310312-3201200233102200-0011203231010231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ip` properties

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-3030301003002223-1121312121232112-1232303010012000-0213330232003222-0030121133101123-1130200110322121-0013030301203132-0203031312323211)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-0122030300032322-3101313310110231-2201200223333032-1111223213311301-3211302322021031-2210200012133330-3303132132031200-1011331323321121)
- ip

<a id="canonical-0200323302101033-3000231232020123-2220103000201020-3300102210300211-0312301312001211-1310132003011330-0213020101033123-0211121130112220"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Additional upstream details:

IPv4 Address in dot-decimal notation.

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
ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102020031223133-0221302111202002-3022330311112101-0023003121313032-3211202202000000-1213230331233232-1211132013231030-0011231223302313"></a>

### Direct properties for `ip`

<a id="canonical-1331123221230203-0033303213331222-0100030002110013-1033232021213220-1300202120002012-3231103321012000-1313302130102232-3201012330100001"></a>

#### `ip.addr` property

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2133031010321330-0110012311313220-1021213332102332-1032201111113230-0123321302231123-0130032231323311-3003311121000102-2321032213111120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `password` properties

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-3030301003002223-1121312121232112-1232303010012000-0213330232003222-0030121133101123-1130200110322121-0013030301203132-0203031312323211)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-0122030300032322-3101313310110231-2201200223333032-1111223213311301-3211302322021031-2210200012133330-3303132132031200-1011331323321121)
- password

<a id="canonical-1331031232030132-0300332322012130-3331100323320213-3220113112222123-3221212302223000-3212202013021233-3330330010032301-1312113100130133"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030100110301000-1323221312030211-2202321131232300-1211231032322031-3212321221322303-0131001303221220-0100120023303200-3010300311211320"></a>

### Direct properties for `password`

- [blindfold_secret_info](resources--cminstance--reference--group-001.md#canonical-1311130101033113-2301223330003302-2300110230013231-0312110021321130-1223212303222120-1200122012232332-1032203210120020-2222331103302001): complete subsection reference.

- [clear_secret_info](resources--cminstance--reference--group-001.md#canonical-1112131122013200-0021332010030123-1313021323332023-2300333033003101-1123021313320103-2212213130111210-1222321200332033-2323233301112111): complete subsection reference.

<a id="canonical-1311130101033113-2301223330003302-2300110230013231-0312110021321130-1223212303222120-1200122012232332-1032203210120020-2222331103302001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-3030301003002223-1121312121232112-1232303010012000-0213330232003222-0030121133101123-1130200110322121-0013030301203132-0203031312323211)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-0122030300032322-3101313310110231-2201200223333032-1111223213311301-3211302322021031-2210200012133330-3303132132031200-1011331323321121)
- [password](resources--cminstance--reference--group-001.md#canonical-2133031010321330-0110012311313220-1021213332102332-1032201111113230-0123321302231123-0130032231323311-3003311121000102-2321032213111120)
- password.blindfold_secret_info

<a id="canonical-2300220322233132-2030221202312032-2322303203321132-0301132123022123-0213101110120233-1231230121303330-1122002331033022-1302231003220020"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321222112323232-0233233132103133-0012021212331231-1331313110111133-0133130030120313-3301003310232110-0023310131023010-1120222031031300"></a>

### Direct properties for `password.blindfold_secret_info`

<a id="canonical-1132031112222120-0010221020110330-1000311020200021-1013323002131202-1031301303113110-3302301301321133-0133303112333213-1210311131102121"></a>

#### `password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0323110033202211-1103223130231100-0212032013122330-2012122322312230-2222312011221211-3010123020013120-0010210210322232-3102111021231322"></a>

<a id="canonical-0231131021223011-1333331111232200-0023333010313001-2231200312130333-0220033323312013-2023230301010103-3032321331222123-2303012211001132"></a>

#### `password.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1221302320311102-1322003213223313-2311213222222023-1121013032101212-3200323013310101-2212311311131123-3331122223200203-0133202112203213"></a>

<a id="canonical-3320013221213230-3031210103223210-1213101311021321-3203213133212110-3201300100223201-1003130030130202-0202222320102233-3311031231030221"></a>

#### `password.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1112131122013200-0021332010030123-1313021323332023-2300333033003101-1123021313320103-2212213130111210-1222321200332033-2323233301112111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-3030301003002223-1121312121232112-1232303010012000-0213330232003222-0030121133101123-1130200110322121-0013030301203132-0203031312323211)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-0122030300032322-3101313310110231-2201200223333032-1111223213311301-3211302322021031-2210200012133330-3303132132031200-1011331323321121)
- [password](resources--cminstance--reference--group-001.md#canonical-2133031010321330-0110012311313220-1021213332102332-1032201111113230-0123321302231123-0130032231323311-3003311121000102-2321032213111120)
- password.clear_secret_info

<a id="canonical-0212031133121131-3133033321311331-3311100203130311-3112232231002023-2231220023223203-2011331020032221-0202220001330012-2310221123123113"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302031220112221-3221031013320203-3232333112211221-3322233332101013-0233132111211032-3213321311330222-2123021123330013-3231032020130203"></a>

### Direct properties for `password.clear_secret_info`

<a id="canonical-3122333001001122-2132102302221301-2230233333332032-1102100113133012-3123300120200222-2213201311232200-3001222110200230-2201212102131021"></a>

#### `password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1310111230010330-2013001221013313-3321332223100021-0332323321023031-3233013220210031-1030001111111201-1220212221232113-0011201231132003"></a>

<a id="canonical-2011303001323333-1222300231003212-1020013131323303-2101202231130021-2120312133123030-0102113311032223-3121100110100223-0012201312221030"></a>

#### `password.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3012012022112030-3310230301133232-1001311133212011-3332112331333221-3032000313020313-0020131223101003-2012031130303010-3030332222211210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-3030301003002223-1121312121232112-1232303010012000-0213330232003222-0030121133101123-1130200110322121-0013030301203132-0203031312323211)
- [Property reference](resources--cminstance--reference--group-001.md#canonical-0122030300032322-3101313310110231-2201200223333032-1111223213311301-3211302322021031-2210200012133330-3303132132031200-1011331323321121)
- timeouts

<a id="canonical-1031300310232103-2212332332301212-0231132112011000-0231213021121031-2003323202020011-3131311101101210-2202311103000212-1011303123021031"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1022332212100001-1210323202333323-1033133121231010-3123113131013013-2011303112312211-0110112123032312-0223122230220323-0320020122133021"></a>

### Direct properties for `timeouts`

<a id="canonical-3330021233132032-3233101313301221-3131021322121101-2103332233000200-0333220012302131-2333310113202313-0103010200211320-3012100211101222"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2011210022023333-2300320310123321-1123213333203022-1310223330113310-3023303013200010-3232112333320203-3332112131303111-1210010012133221"></a>

<a id="canonical-0030310232233121-0131213200000223-3100131003113331-3313301020111323-0321103120200201-2002033103330300-1302232300320130-0201322111302021"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2131000010312030-2023113202332122-1332002202121010-3300231023330300-0011221110233232-1130311012112132-1321333303220201-0012213121101223"></a>

<a id="canonical-3131000330301231-0311102212203013-1100032330023110-1213302322203213-0330211123303312-3230013312331023-3123201231021133-0233331120122021"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2130013000130110-3030132210112131-3310211331110221-1131311113100323-2231003212023313-1002011220223320-0233203113110123-3110321133220121"></a>

<a id="canonical-2220123032012332-2102131003132202-2330131011221213-1102102230102312-3212033022222010-0003220210010133-1001131212110323-0101110110101120"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
