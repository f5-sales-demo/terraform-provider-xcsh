---
page_title: "xcsh_api_crawler reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_crawler reference."
---

# xcsh_api_crawler reference

<a id="canonical-3010222002133020-2211321021131000-2301002232300112-1123303232312202-2000011002203003-2003112030230322-2103211310110013-1020232220123011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-3133010202200222-3213011120000312-1321322311223200-2311232132101333-2211031002320321-3221011221222302-0233012202301010-3201312221013200)
- Property reference

<a id="canonical-1113332031022202-2121200012310110-1301210331213113-3022200212220123-0133311032220010-2033231302320210-0232002103113130-0000211013310100"></a>

### Direct properties for `xcsh_api_crawler`

<a id="canonical-2133303113110200-1120213111130011-3313133123001122-2211332130211311-1300301221210213-0101102321102223-2221203112323332-0210000212002111"></a>

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

<a id="canonical-1201103130201231-0113110000210001-2330221102002210-3221110323233030-0102010012100232-0310212330213121-3203333232213213-2112102202331312"></a>

<a id="canonical-3003210310023023-1211222331200003-2002013300302231-3201212300112311-1302232123311030-3323023021321321-1021220223123120-2003132221202003"></a>

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

<a id="canonical-2002013311100022-2212303312130230-0101221131121310-3033213300312101-3002333023121310-0302221130303230-1013321230001100-1333110233111221"></a>

<a id="canonical-2011220033202210-2120033321133113-0213000022100111-0010130000123131-0330032201300120-0302313203011320-1133231222203002-0310021121020022"></a>

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

- [domains](resources--api_crawler--reference--group-001.md#canonical-2023202200302200-3220100302000201-1103011010332303-2230201122011301-2110202200310302-0233212012122331-1231130120112020-1311130230231300): complete subsection reference.

<a id="canonical-0213021020213131-3102213231130103-2032310231121222-3032323121020022-0210221001132212-0121133230212111-1030112010133000-3023103300031122"></a>

<a id="canonical-2200011230200232-3121123020033302-2111303021232233-0210012220321220-1030333330021102-0131013130323320-3003033311303310-1002323323103322"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3210133033321331-1302221303131332-0111021121333321-2130002221000111-2211223313330031-0121020012021120-0231213031030231-3311223203310212"></a>

<a id="canonical-2220102222000013-2103013311311003-2033223233333330-1012111001121310-0321111003221012-1000211332311123-2230111102310101-3000100033331301"></a>

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

<a id="canonical-0030203301222300-0122331020221103-2330130231101331-2232200210202331-3100031013133022-3010332330303320-1303121003221031-0122030101321212"></a>

<a id="canonical-1210232110023021-0120013112200201-0001320001133103-3231012312110111-2023202020310100-0221321023022301-0100203032030303-3121303023202102"></a>

#### `name` property

Type: `"string"`. Required.

Name of the API Crawler. Must be unique within the namespace.

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

<a id="canonical-0212313131012032-2232011012302023-2021300230123102-3212032211322103-1313111121303121-2032211010221130-1113110333223330-2110303031331012"></a>

<a id="canonical-0313312120302200-1303230132030113-2002303333311100-1001101000111320-1021303221033023-0200132220322230-3121203332323212-1023313322032113"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the API Crawler is created.

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

- [timeouts](resources--api_crawler--reference--group-001.md#canonical-1221322320313023-3033221323112112-0101320132131100-1031213300132313-1133210112333322-3012001210003233-2332323300331112-0300021002202111): complete subsection reference.

<a id="canonical-3011330130031200-1202310003231020-3213310313020031-0333012223211003-2100123011131120-2133323103302022-0101322030220211-3233221001101003"></a>

### All schema paths for `xcsh_api_crawler`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--api_crawler--reference--group-001.md#canonical-2133303113110200-1120213111130011-3313133123001122-2211332130211311-1300301221210213-0101102321102223-2221203112323332-0210000212002111) |
| `description` | [description](resources--api_crawler--reference--group-001.md#canonical-1201103130201231-0113110000210001-2330221102002210-3221110323233030-0102010012100232-0310212330213121-3203333232213213-2112102202331312) |
| `disable` | [disable](resources--api_crawler--reference--group-001.md#canonical-2002013311100022-2212303312130230-0101221131121310-3033213300312101-3002333023121310-0302221130303230-1013321230001100-1333110233111221) |
| `domains` | [domains](resources--api_crawler--reference--group-001.md#canonical-0203031000021223-1210213133033013-2222030320112121-1323020122132211-1231230133013231-3113103020133300-3310010220023301-0311000211213032) |
| `domains.domain` | [domains.domain](resources--api_crawler--reference--group-001.md#canonical-3010103013213310-0331022130112010-2010131011102223-2122210322210021-1230133010321113-1320302123221221-1011132111311022-0122311021003310) |
| `domains.simple_login` | [domains.simple_login](resources--api_crawler--reference--group-001.md#canonical-1123333301320321-2033120033330100-2122031020203133-1210011202113122-0312033210311003-0232320332223131-1223000130012103-0130011103310100) |
| `domains.simple_login.password` | [domains.simple_login.password](resources--api_crawler--reference--group-001.md#canonical-3031312223300003-0123213313110032-3222212111322001-2230112101212221-0112133121033222-3322110001213100-1300320303232312-3313013321103230) |
| `domains.simple_login.password.blindfold_secret_info` | [domains.simple_login.password.blindfold_secret_info](resources--api_crawler--reference--group-001.md#canonical-2123201102303002-3123001213011010-0232332113121103-2130021232313021-1220220113231000-2010312302011233-3302202132300221-3203020120110123) |
| `domains.simple_login.password.blindfold_secret_info.decryption_provider` | [domains.simple_login.password.blindfold_secret_info.decryption_provider](resources--api_crawler--reference--group-001.md#canonical-3111132200002301-0300210103313313-3311013132033332-0123123102211110-0333233121323103-3212223203101021-0301032002031112-2300303113033123) |
| `domains.simple_login.password.blindfold_secret_info.location` | [domains.simple_login.password.blindfold_secret_info.location](resources--api_crawler--reference--group-001.md#canonical-3311012311313031-0230032102300222-0130203132300013-1220201013330000-3302110222212203-1022130302010220-3122120330212010-3322023312110321) |
| `domains.simple_login.password.blindfold_secret_info.store_provider` | [domains.simple_login.password.blindfold_secret_info.store_provider](resources--api_crawler--reference--group-001.md#canonical-3121031221320033-2133022223100300-2130330221121202-3120332233323211-3300013003133132-3100302202231301-1030233030331123-3011220103310220) |
| `domains.simple_login.password.clear_secret_info` | [domains.simple_login.password.clear_secret_info](resources--api_crawler--reference--group-001.md#canonical-2100330101022301-1322113210203001-2100323011031313-3203233200223233-0101002133020201-0111303303200220-3102013011010313-3113323020132222) |
| `domains.simple_login.password.clear_secret_info.provider_ref` | [domains.simple_login.password.clear_secret_info.provider_ref](resources--api_crawler--reference--group-001.md#canonical-2321221310203020-0330003332020131-2102010130222303-2300320210233303-1300302031020211-0123023020123333-3031120301210302-0123202201123221) |
| `domains.simple_login.password.clear_secret_info.url` | [domains.simple_login.password.clear_secret_info.url](resources--api_crawler--reference--group-001.md#canonical-3030312301011122-0113202133221211-2231312312020103-1303023230321123-1330213131312031-2010312333212230-3323120211303021-1202022210013212) |
| `domains.simple_login.user` | [domains.simple_login.user](resources--api_crawler--reference--group-001.md#canonical-1202330003012001-1311133321133030-1011320300201012-0102100201013230-2020010203211212-2232113121223221-2310312312031202-1122202012220120) |
| `id` | [ID](resources--api_crawler--reference--group-001.md#canonical-0213021020213131-3102213231130103-2032310231121222-3032323121020022-0210221001132212-0121133230212111-1030112010133000-3023103300031122) |
| `labels` | [labels](resources--api_crawler--reference--group-001.md#canonical-3210133033321331-1302221303131332-0111021121333321-2130002221000111-2211223313330031-0121020012021120-0231213031030231-3311223203310212) |
| `name` | [name](resources--api_crawler--reference--group-001.md#canonical-0030203301222300-0122331020221103-2330130231101331-2232200210202331-3100031013133022-3010332330303320-1303121003221031-0122030101321212) |
| `namespace` | [namespace](resources--api_crawler--reference--group-001.md#canonical-0212313131012032-2232011012302023-2021300230123102-3212032211322103-1313111121303121-2032211010221130-1113110333223330-2110303031331012) |
| `timeouts` | [timeouts](resources--api_crawler--reference--group-001.md#canonical-3322233110300121-1020101121032021-0020011302133120-3031033212333330-1301323102222331-3221111320333112-2220330131001200-3211312100302003) |
| `timeouts.create` | [timeouts.create](resources--api_crawler--reference--group-001.md#canonical-0310003221200001-1030313110223101-2302220013031022-1323312100200113-1302130222023221-3331202031200230-1031020200033010-2310112113323023) |
| `timeouts.delete` | [timeouts.delete](resources--api_crawler--reference--group-001.md#canonical-2200320232221111-3103120233110032-1333313213022100-2213100232021110-2311223332220230-0122211100003131-2012202222323013-0011111301322331) |
| `timeouts.read` | [timeouts.read](resources--api_crawler--reference--group-001.md#canonical-1203121300233300-3103201221213322-2301021030230302-1103010010120102-0303303120031010-0211222002123211-0112122111213311-3223010122310212) |
| `timeouts.update` | [timeouts.update](resources--api_crawler--reference--group-001.md#canonical-3103222113320133-0320223132113023-0020100203213302-3123000220000202-3001030312120122-2331101000031031-3312031112011310-0210123132132100) |

<a id="canonical-2023202200302200-3220100302000201-1103011010332303-2230201122011301-2110202200310302-0233212012122331-1231130120112020-1311130230231300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains` properties

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-3133010202200222-3213011120000312-1321322311223200-2311232132101333-2211031002320321-3221011221222302-0233012202301010-3201312221013200)
- [Property reference](resources--api_crawler--reference--group-001.md#canonical-3010222002133020-2211321021131000-2301002232300112-1123303232312202-2000011002203003-2003112030230322-2103211310110013-1020232220123011)
- domains

<a id="canonical-0203031000021223-1210213133033013-2222030320112121-1323020122132211-1231230133013231-3113103020133300-3310010220023301-0311000211213032"></a>

Type: `"object"`. list nested block, Optional.

API Crawler. API Crawler Configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003100200003101-0212000220200231-1001300212130023-0112222033111331-3123202222033202-1301223032131332-3121303021010331-0133010301011331"></a>

### Direct properties for `domains`

<a id="canonical-3010103013213310-0331022130112010-2010131011102223-2122210322210021-1230133010321113-1320302123221221-1011132111311022-0122311021003310"></a>

#### `domains.domain` property

Type: `"string"`. Optional.

Select the domain to execute API Crawling with given credentials.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](resources--api_crawler--reference--group-001.md#canonical-0212033122200100-2231032213222313-3323330121120013-1233000131220132-0210012302311202-1330033102322111-1121332301231231-0323102010222013): complete subsection reference.

<a id="canonical-0212033122200100-2231032213222313-3323330121120013-1233000131220132-0210012302311202-1330033102322111-1121332301231231-0323102010222013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.simple_login` properties

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-3133010202200222-3213011120000312-1321322311223200-2311232132101333-2211031002320321-3221011221222302-0233012202301010-3201312221013200)
- [Property reference](resources--api_crawler--reference--group-001.md#canonical-3010222002133020-2211321021131000-2301002232300112-1123303232312202-2000011002203003-2003112030230322-2103211310110013-1020232220123011)
- [domains](resources--api_crawler--reference--group-001.md#canonical-2023202200302200-3220100302000201-1103011010332303-2230201122011301-2110202200310302-0233212012122331-1231130120112020-1311130230231300)
- domains.simple_login

<a id="canonical-1123333301320321-2033120033330100-2122031020203133-1210011202113122-0312033210311003-0232320332223131-1223000130012103-0130011103310100"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for simple login.

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
simple_login {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130033121101223-2311212332312123-1322011210331021-2231110030212230-3002332302130032-0003131233303120-3011130330010000-1020321123112211"></a>

### Direct properties for `domains.simple_login`

- [password](resources--api_crawler--reference--group-001.md#canonical-2130123112313201-0122233002333311-0312121203002003-3030211023311213-1111011212323013-2032020232011131-1330210012221201-3212030030212111): complete subsection reference.

<a id="canonical-1202330003012001-1311133321133030-1011320300201012-0102100201013230-2020010203211212-2232113121223221-2310312312031202-1122202012220120"></a>

<a id="canonical-1222100002002101-3031000321233131-3130103332022200-0021331022221330-3021111031013213-2230022132012212-1120213102111003-3102321013323001"></a>

#### `domains.simple_login.user` property

Type: `"string"`. Optional.

Enter the username to assign credentials for the selected domain to crawl.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2130123112313201-0122233002333311-0312121203002003-3030211023311213-1111011212323013-2032020232011131-1330210012221201-3212030030212111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.simple_login.password` properties

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-3133010202200222-3213011120000312-1321322311223200-2311232132101333-2211031002320321-3221011221222302-0233012202301010-3201312221013200)
- [Property reference](resources--api_crawler--reference--group-001.md#canonical-3010222002133020-2211321021131000-2301002232300112-1123303232312202-2000011002203003-2003112030230322-2103211310110013-1020232220123011)
- [domains](resources--api_crawler--reference--group-001.md#canonical-2023202200302200-3220100302000201-1103011010332303-2230201122011301-2110202200310302-0233212012122331-1231130120112020-1311130230231300)
- [domains.simple_login](resources--api_crawler--reference--group-001.md#canonical-0212033122200100-2231032213222313-3323330121120013-1233000131220132-0210012302311202-1330033102322111-1121332301231231-0323102010222013)
- domains.simple_login.password

<a id="canonical-3031312223300003-0123213313110032-3222212111322001-2230112101212221-0112133121033222-3322110001213100-1300320303232312-3313013321103230"></a>

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

<a id="canonical-3220323302211221-2103113311120300-3230233310103023-2332210220302221-1231221230031102-2222203303223212-1103223012201312-3103023232200303"></a>

### Direct properties for `domains.simple_login.password`

- [blindfold_secret_info](resources--api_crawler--reference--group-001.md#canonical-2030121111022321-2100200211121321-0020203211333321-0330002232033332-3030333200121313-0002012312033101-3210020033032133-1313203323120231): complete subsection reference.

- [clear_secret_info](resources--api_crawler--reference--group-001.md#canonical-2121322202022200-2313022130230120-3000130131210230-3323003130011332-2322320302121312-2312231210321230-1101022021232311-2310202120332223): complete subsection reference.

<a id="canonical-2030121111022321-2100200211121321-0020203211333321-0330002232033332-3030333200121313-0002012312033101-3210020033032133-1313203323120231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.simple_login.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-3133010202200222-3213011120000312-1321322311223200-2311232132101333-2211031002320321-3221011221222302-0233012202301010-3201312221013200)
- [Property reference](resources--api_crawler--reference--group-001.md#canonical-3010222002133020-2211321021131000-2301002232300112-1123303232312202-2000011002203003-2003112030230322-2103211310110013-1020232220123011)
- [domains](resources--api_crawler--reference--group-001.md#canonical-2023202200302200-3220100302000201-1103011010332303-2230201122011301-2110202200310302-0233212012122331-1231130120112020-1311130230231300)
- [domains.simple_login](resources--api_crawler--reference--group-001.md#canonical-0212033122200100-2231032213222313-3323330121120013-1233000131220132-0210012302311202-1330033102322111-1121332301231231-0323102010222013)
- [domains.simple_login.password](resources--api_crawler--reference--group-001.md#canonical-2130123112313201-0122233002333311-0312121203002003-3030211023311213-1111011212323013-2032020232011131-1330210012221201-3212030030212111)
- domains.simple_login.password.blindfold_secret_info

<a id="canonical-2123201102303002-3123001213011010-0232332113121103-2130021232313021-1220220113231000-2010312302011233-3302202132300221-3203020120110123"></a>

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

<a id="canonical-2113132033023232-3011320013132223-0133232121330111-1320123110223232-2001300133331211-3221232331130010-1023021223031200-3211003000030110"></a>

### Direct properties for `domains.simple_login.password.blindfold_secret_info`

<a id="canonical-3111132200002301-0300210103313313-3311013132033332-0123123102211110-0333233121323103-3212223203101021-0301032002031112-2300303113033123"></a>

#### `domains.simple_login.password.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3311012311313031-0230032102300222-0130203132300013-1220201013330000-3302110222212203-1022130302010220-3122120330212010-3322023312110321"></a>

<a id="canonical-3131231213132001-1002022110130320-3103013232310221-3320323012332102-0310313312331213-0120323121232203-3031100022321222-0321330300230133"></a>

#### `domains.simple_login.password.blindfold_secret_info.location` property

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

<a id="canonical-3121031221320033-2133022223100300-2130330221121202-3120332233323211-3300013003133132-3100302202231301-1030233030331123-3011220103310220"></a>

<a id="canonical-3010110132111110-1230002111000303-3332233223332321-2321302033320302-0000220201203233-1132131200312210-0002133133033123-1120310302312330"></a>

#### `domains.simple_login.password.blindfold_secret_info.store_provider` property

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

<a id="canonical-2121322202022200-2313022130230120-3000130131210230-3323003130011332-2322320302121312-2312231210321230-1101022021232311-2310202120332223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.simple_login.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-3133010202200222-3213011120000312-1321322311223200-2311232132101333-2211031002320321-3221011221222302-0233012202301010-3201312221013200)
- [Property reference](resources--api_crawler--reference--group-001.md#canonical-3010222002133020-2211321021131000-2301002232300112-1123303232312202-2000011002203003-2003112030230322-2103211310110013-1020232220123011)
- [domains](resources--api_crawler--reference--group-001.md#canonical-2023202200302200-3220100302000201-1103011010332303-2230201122011301-2110202200310302-0233212012122331-1231130120112020-1311130230231300)
- [domains.simple_login](resources--api_crawler--reference--group-001.md#canonical-0212033122200100-2231032213222313-3323330121120013-1233000131220132-0210012302311202-1330033102322111-1121332301231231-0323102010222013)
- [domains.simple_login.password](resources--api_crawler--reference--group-001.md#canonical-2130123112313201-0122233002333311-0312121203002003-3030211023311213-1111011212323013-2032020232011131-1330210012221201-3212030030212111)
- domains.simple_login.password.clear_secret_info

<a id="canonical-2100330101022301-1322113210203001-2100323011031313-3203233200223233-0101002133020201-0111303303200220-3102013011010313-3113323020132222"></a>

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

<a id="canonical-3032203031233003-1222231203003323-2230123131300113-1211233002200313-3102333232202121-1021013020330321-3223030023131302-3021022331222121"></a>

### Direct properties for `domains.simple_login.password.clear_secret_info`

<a id="canonical-2321221310203020-0330003332020131-2102010130222303-2300320210233303-1300302031020211-0123023020123333-3031120301210302-0123202201123221"></a>

#### `domains.simple_login.password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3030312301011122-0113202133221211-2231312312020103-1303023230321123-1330213131312031-2010312333212230-3323120211303021-1202022210013212"></a>

<a id="canonical-2300302221001303-0033230022003003-3032221033302303-1130222331100311-1012212020020311-1231130022323223-3210322110103331-3012032212112131"></a>

#### `domains.simple_login.password.clear_secret_info.url` property

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

<a id="canonical-1221322320313023-3033221323112112-0101320132131100-1031213300132313-1133210112333322-3012001210003233-2332323300331112-0300021002202111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-3133010202200222-3213011120000312-1321322311223200-2311232132101333-2211031002320321-3221011221222302-0233012202301010-3201312221013200)
- [Property reference](resources--api_crawler--reference--group-001.md#canonical-3010222002133020-2211321021131000-2301002232300112-1123303232312202-2000011002203003-2003112030230322-2103211310110013-1020232220123011)
- timeouts

<a id="canonical-3322233110300121-1020101121032021-0020011302133120-3031033212333330-1301323102222331-3221111320333112-2220330131001200-3211312100302003"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011022222313013-0223100301232122-2221323203230133-0121311032300333-0023212203232211-2112233123031213-1130002010312321-1031113333121113"></a>

### Direct properties for `timeouts`

<a id="canonical-0310003221200001-1030313110223101-2302220013031022-1323312100200113-1302130222023221-3331202031200230-1031020200033010-2310112113323023"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2200320232221111-3103120233110032-1333313213022100-2213100232021110-2311223332220230-0122211100003131-2012202222323013-0011111301322331"></a>

<a id="canonical-0322221032130022-3222302322233212-1032201113133001-2001332112030102-1001103203032133-2202211012032133-3101332013130233-2321322202132320"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1203121300233300-3103201221213322-2301021030230302-1103010010120102-0303303120031010-0211222002123211-0112122111213311-3223010122310212"></a>

<a id="canonical-1130113132010103-2312310213110002-3321130031322002-3212200010210003-3123130310012030-2200212220203132-2001320322112010-2221123130101202"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3103222113320133-0320223132113023-0020100203213302-3123000220000202-3001030312120122-2331101000031031-3312031112011310-0210123132132100"></a>

<a id="canonical-0211303033312121-1102120213220321-1003213200221011-0012101212031102-2313331011321020-1220102300012302-1212121122201100-1310023313332003"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
