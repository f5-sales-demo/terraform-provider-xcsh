---
page_title: "xcsh_service_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy reference."
---

# xcsh_service_policy reference

<a id="canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- Property reference

<a id="canonical-1112013330003300-2133021322211022-2321210221123330-2201303203011111-3231321023201221-3000333111032032-0220221231110032-2122011333031101"></a>

### Direct properties for `xcsh_service_policy`

- [allow_all_requests](resources--service_policy--reference--group-001.md#canonical-0323122132121121-2032121222012200-2021300003103320-2200300021120320-3213122032133013-1012030312210210-1131212101302130-1131330232103211): complete subsection reference.

- [allow_list](resources--service_policy--reference--group-001.md#canonical-2103300311300332-3332200122100221-1303003012310332-1301221201020110-2111112310210320-2231312202202012-2331211230323103-2101301311212102): complete subsection reference.

<a id="canonical-3001323130001222-3332322111003220-1100300100121302-2003202100213300-1303000013313022-2320233212011012-0023320301213203-3022122112333320"></a>

<a id="canonical-2103310232100111-1110322001123003-1310121313212312-3130030002031202-0222111131123111-0332233313131300-3100002123132230-3012231033030220"></a>

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

- [any_server](resources--service_policy--reference--group-001.md#canonical-2312320122102201-0133030001231211-3021113302202221-0231313332122131-1331023031332202-2100201311220333-3033202032033122-2311001313122213): complete subsection reference.

- [deny_all_requests](resources--service_policy--reference--group-001.md#canonical-3221222222001232-0223003221133031-1130311332020010-2103301202310211-1113302100001102-0301330330032012-1001212010020311-3300313001303012): complete subsection reference.

- [deny_list](resources--service_policy--reference--group-001.md#canonical-3300232031103221-3132103232200031-0132033333013022-1032110132202122-0000103233010213-1232110230300121-0103203200203111-0001311012323210): complete subsection reference.

<a id="canonical-1302311212221301-1310103302120232-2132221321101212-2113113121120303-0100132122021012-3312131321311002-3311232011000012-2120113310100221"></a>

<a id="canonical-0223121000013031-0320303231233302-1310031001100112-0132310033023030-3122213001202112-3232310202302132-1221303312233222-3011200121121010"></a>

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

<a id="canonical-1133121223221123-1120002202202221-0002112233231100-1120303301103020-1111210301033131-1132130320232301-3322223222112302-0302130112133323"></a>

<a id="canonical-3031322321110001-1000132312221333-3113100230202001-2222003323201021-0231100223212030-0003131023220203-0231220210101120-3023212233333301"></a>

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

<a id="canonical-1020101210221131-2002323212013003-1113323111303303-1002231322231110-1110200213233023-2212030220102122-2320113310322020-3322102020130323"></a>

<a id="canonical-1200130302232332-0011301330022123-0213111220030013-2222202231302022-0131330111222202-3002303300113020-3103032201301121-0102131123322212"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3132302113200002-3003131101313310-0000300321230130-3102003112321303-3101232001012002-0101122333310220-1112030300220110-1112222012310330"></a>

<a id="canonical-3111103000333320-3221001320012100-0100032321322313-2122223120321011-0221333012002112-1202101223203011-3000210032212133-2100333123110132"></a>

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

<a id="canonical-1312102120133113-1000012300011301-2212322211111313-1020212320212312-2212021032210022-3103313233021100-1221130230303023-2213311002021103"></a>

<a id="canonical-0221102133332231-1223122213333101-0231111101131310-1313110030110103-2222230033222130-2330330010120021-0212321301303101-3003221212330300"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Service Policy. Must be unique within the namespace.

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

<a id="canonical-3310021011323222-0223223222003333-3222230033321322-2321321232233132-3302312112230100-3003201000020302-2322212003011210-0300302313032131"></a>

<a id="canonical-3332303102330211-2322231103032101-1002331111133132-3221322320130310-0013231131022100-2131110112331012-0031312213023013-1120313132200023"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Service Policy is created.

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

- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100): complete subsection reference.

<a id="canonical-0133223010313332-1110211023313330-0010322022330102-0322023303232002-3232002130100231-3032013031322121-1301220213200220-0101110033331123"></a>

<a id="canonical-1301033130202301-0122001120331131-1132033210331213-3323011021231333-0032301323210010-0120030221310133-3002201013200232-1022330322201010"></a>

#### `server_name` property

Type: `"string"`. Optional, Computed.

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server to which the request API is directed. The actual names for the server are extracted from the
HTTP Host header and the name of the virtual\_host to which the request is directed. If the request
is directed to a virtual K8s service, the actual names also contain the name of that service. The
predicate evaluates to true if any of the actual names is the same as the expected server name.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [server_name_matcher](resources--service_policy--reference--group-003.md#canonical-3132031220110012-1111121333101001-0122230010332012-3320330110332133-3010320201321321-1211023113112212-1300122212023203-2020020300213013): complete subsection reference.

- [server_selector](resources--service_policy--reference--group-003.md#canonical-3032002123132023-0202122031333133-0202113320332213-2000203232033130-0232320232222201-3022133310223233-0213130313331202-2000111311332223): complete subsection reference.

- [timeouts](resources--service_policy--reference--group-003.md#canonical-1030022313211210-1123201233300031-2301231122113220-0310203223011103-2122310011221001-2023101232323101-1031133331010301-1220301100022212): complete subsection reference.

<a id="canonical-3230122213232022-0122332230300003-2313221110302030-3033102311311213-3003322202122213-3032220103221230-1022122032110231-3112203001311313"></a>

### All schema paths for `xcsh_service_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_requests` | [allow_all_requests](resources--service_policy--reference--group-001.md#canonical-3203302312022132-1033132230002030-2122211232332320-0331003202232001-1301111103130220-3331003033201213-0231110210230311-0011302321201113) |
| `allow_list` | [allow_list](resources--service_policy--reference--group-001.md#canonical-1021110011030020-2000303211110001-2211321101133001-2322031003001032-0220201300302131-3232100200132030-1200230000110101-1003100222010023) |
| `allow_list.asn_list` | [allow_list.asn_list](resources--service_policy--reference--group-001.md#canonical-1001102103011112-1100132211110033-3021312130211230-0232210333130200-0111230330121102-2233223222132110-1011103012032232-2110312013232130) |
| `allow_list.asn_list.as_numbers` | [allow_list.asn_list.as_numbers](resources--service_policy--reference--group-001.md#canonical-0002101112202000-0313112133231111-3222121311110103-0123111133312311-3120131102323001-1331331133112300-2300332131333111-1123100132100020) |
| `allow_list.asn_set` | [allow_list.asn_set](resources--service_policy--reference--group-001.md#canonical-1201113321321311-1230313201033103-1211330102111321-0112201200322310-1301333320033133-1332211233103023-1201210031023220-0121231101333323) |
| `allow_list.asn_set.name` | [allow_list.asn_set.name](resources--service_policy--reference--group-001.md#canonical-2102112202111302-1032112132103102-0010101123102232-3211032011001203-2210121101130213-1211023013332210-1021200100211130-3111213131210313) |
| `allow_list.asn_set.namespace` | [allow_list.asn_set.namespace](resources--service_policy--reference--group-001.md#canonical-1021023233321320-2100332013002330-3221103112302231-3130000033210231-3200000302213233-3000211332223011-2223232132301030-2212021022000031) |
| `allow_list.asn_set.tenant` | [allow_list.asn_set.tenant](resources--service_policy--reference--group-001.md#canonical-1003112013131201-3001330323200330-3211120033132312-2020032110333323-0032203020020220-3001122202121213-3221331032123032-1300200103203122) |
| `allow_list.country_list` | [allow_list.country_list](resources--service_policy--reference--group-001.md#canonical-3300031100203020-1110110300122003-1210213200212101-1313100030012032-2103300323120122-0013022102000323-0303023030003101-0011233211033023) |
| `allow_list.default_action_allow` | [allow_list.default_action_allow](resources--service_policy--reference--group-001.md#canonical-0032220000203011-3033301322320311-0321001130202210-2233212103133231-0213112320132220-3331130130212232-0313122003330000-2323232023030120) |
| `allow_list.default_action_deny` | [allow_list.default_action_deny](resources--service_policy--reference--group-001.md#canonical-3200033300231231-2031100211210333-0101300302213323-3111323012233203-3112012212331213-3133211333011001-3313022313110031-0130010101301021) |
| `allow_list.default_action_next_policy` | [allow_list.default_action_next_policy](resources--service_policy--reference--group-001.md#canonical-0330020331031103-2312200032322032-1333110323333120-3203132330003102-0231220110023320-3230333301101013-3301023201101112-0002011302233100) |
| `allow_list.ip_prefix_set` | [allow_list.ip_prefix_set](resources--service_policy--reference--group-001.md#canonical-3021001031301230-0233322022003232-0210312002232123-3221331101002001-0333130231103103-3133132232320031-3200102320322330-0301213000000002) |
| `allow_list.ip_prefix_set.name` | [allow_list.ip_prefix_set.name](resources--service_policy--reference--group-001.md#canonical-0110020331203121-1020120002102132-1302120332212212-1220103113022001-1221113020200031-3021003031300300-2120130212302211-3313202213211201) |
| `allow_list.ip_prefix_set.namespace` | [allow_list.ip_prefix_set.namespace](resources--service_policy--reference--group-001.md#canonical-2331322222123101-1320232311030121-2023201231001220-1211030213000312-0001201232313333-3123112130023010-0011310211011321-1303133112111311) |
| `allow_list.ip_prefix_set.tenant` | [allow_list.ip_prefix_set.tenant](resources--service_policy--reference--group-001.md#canonical-1301110231122301-3131211323111120-2013321100100123-0010121102113013-1003323313023302-2103303010100213-1320202133312003-1110133230221333) |
| `allow_list.prefix_list` | [allow_list.prefix_list](resources--service_policy--reference--group-001.md#canonical-0033011332111321-0030120220011120-3030223131130113-1201101301332223-1012313002013322-3111030302302123-1230333222213232-1131310233301311) |
| `allow_list.prefix_list.prefixes` | [allow_list.prefix_list.prefixes](resources--service_policy--reference--group-001.md#canonical-2202030320030211-3301133123300010-0121131113233311-1302110100200122-0121112010303132-2120122222023013-1322333222311013-1231221012302311) |
| `allow_list.tls_fingerprint_classes` | [allow_list.tls_fingerprint_classes](resources--service_policy--reference--group-001.md#canonical-1121122332310221-2310102303223311-3320012100231000-0331223310113301-2322122110101313-3301322230010331-3102313102021233-2301303101132310) |
| `allow_list.tls_fingerprint_values` | [allow_list.tls_fingerprint_values](resources--service_policy--reference--group-001.md#canonical-3031230112013110-3213130223312133-0012100300002101-1030303220313100-2301310121131132-2303002131111002-0210232110110202-0311323031121313) |
| `annotations` | [annotations](resources--service_policy--reference--group-001.md#canonical-3001323130001222-3332322111003220-1100300100121302-2003202100213300-1303000013313022-2320233212011012-0023320301213203-3022122112333320) |
| `any_server` | [any_server](resources--service_policy--reference--group-001.md#canonical-1010022313220032-0013322123301331-3333333232310222-3131013132311222-0312133020002231-1333231312221231-2021000112022302-2222001122223122) |
| `deny_all_requests` | [deny_all_requests](resources--service_policy--reference--group-001.md#canonical-3000321302123202-0112320203222012-3033013210133012-3313121213112032-3300203233300233-1031222223201033-1233002203021101-2001220123301012) |
| `deny_list` | [deny_list](resources--service_policy--reference--group-001.md#canonical-2130322123310112-0302010020132022-3323113023231331-3210232201121112-2122111023320011-2022222301103232-0130032022112322-1323301103100122) |
| `deny_list.asn_list` | [deny_list.asn_list](resources--service_policy--reference--group-001.md#canonical-1210020332123113-3200112301330231-1130230301301020-3323312033023322-1323202223100123-1233012222120002-0330131202122023-2010301103103102) |
| `deny_list.asn_list.as_numbers` | [deny_list.asn_list.as_numbers](resources--service_policy--reference--group-001.md#canonical-0110022231111223-1022230021001002-1223230010011131-2203320221003130-3003003221203233-2321323022133111-1232212110133231-3301122322012303) |
| `deny_list.asn_set` | [deny_list.asn_set](resources--service_policy--reference--group-001.md#canonical-1230233302031012-1223103221010302-3310321112222101-1100110210232021-1000223330122113-2010010033131000-1311313103021033-2212213120132022) |
| `deny_list.asn_set.name` | [deny_list.asn_set.name](resources--service_policy--reference--group-001.md#canonical-1312001320211012-2033312002013013-0030231231033012-0312131121220213-1233233110131231-3302100221210303-3321202210333222-2220201312003200) |
| `deny_list.asn_set.namespace` | [deny_list.asn_set.namespace](resources--service_policy--reference--group-001.md#canonical-3130303331003313-2313130201120031-3301100122011110-3211133212012213-1102321013021032-3011201213302231-0221113012000021-0313103300100003) |
| `deny_list.asn_set.tenant` | [deny_list.asn_set.tenant](resources--service_policy--reference--group-001.md#canonical-1121223211021103-2002211003333202-0232312122232232-2022333130112222-0110213303001011-1032313113002210-3023023221011031-2301211130022210) |
| `deny_list.country_list` | [deny_list.country_list](resources--service_policy--reference--group-001.md#canonical-2332022333132001-3213110000113013-3032022133023113-2033020133103230-0030111222002221-3021120221120312-3013331100023232-2130131130213013) |
| `deny_list.default_action_allow` | [deny_list.default_action_allow](resources--service_policy--reference--group-001.md#canonical-1123123120200202-1333122233120111-2300001310021032-1213202023023331-2130323013110133-1221102220111101-2121320002033201-2320112121033202) |
| `deny_list.default_action_deny` | [deny_list.default_action_deny](resources--service_policy--reference--group-001.md#canonical-2211203123032121-3020033011013203-1000120203221121-1130111231233113-3033132032030112-3202020233011103-3103211233130332-0213212011223310) |
| `deny_list.default_action_next_policy` | [deny_list.default_action_next_policy](resources--service_policy--reference--group-001.md#canonical-3211123131233101-2320010303003001-0000201003311023-0030202210231301-0203232310003303-2332020002331323-1200331022123120-3222111133131002) |
| `deny_list.ip_prefix_set` | [deny_list.ip_prefix_set](resources--service_policy--reference--group-001.md#canonical-2202032310300203-3330321232120201-1330011323301231-1300213233221321-1221012001013203-1300121023223023-3023130333121003-0311102331232200) |
| `deny_list.ip_prefix_set.name` | [deny_list.ip_prefix_set.name](resources--service_policy--reference--group-001.md#canonical-2012203201302133-2100031202020203-1001031012112101-0121302022203230-3112230000021311-0001312302221102-0203212232323010-0121333101100111) |
| `deny_list.ip_prefix_set.namespace` | [deny_list.ip_prefix_set.namespace](resources--service_policy--reference--group-001.md#canonical-2212032131001010-0123210032101030-2121130221213202-0133333302102132-1220100210032113-0333331230230123-2203301200220011-3232100133013000) |
| `deny_list.ip_prefix_set.tenant` | [deny_list.ip_prefix_set.tenant](resources--service_policy--reference--group-001.md#canonical-1232012101310123-1201022013203202-1311313020100012-1201001313331223-2323103111000121-0230232200011222-1301312223310033-0231320110201202) |
| `deny_list.prefix_list` | [deny_list.prefix_list](resources--service_policy--reference--group-001.md#canonical-1012302130003113-2313102211131222-3020202113013121-2303103323103211-0313322330013131-1120223211132310-2013202003002113-2330112201122312) |
| `deny_list.prefix_list.prefixes` | [deny_list.prefix_list.prefixes](resources--service_policy--reference--group-001.md#canonical-0310110333331220-2320312123210211-0221300232303101-1212311000312030-3212113013331201-2133200000223023-1211201110121213-1021112102030112) |
| `deny_list.tls_fingerprint_classes` | [deny_list.tls_fingerprint_classes](resources--service_policy--reference--group-001.md#canonical-0112121103222303-1103111201221201-1312131211030123-0332030302021130-3032101230300020-0030211201203120-3020120302131122-2013302222100021) |
| `deny_list.tls_fingerprint_values` | [deny_list.tls_fingerprint_values](resources--service_policy--reference--group-001.md#canonical-0121311230110031-2102031301031021-0320330221201221-3113313301030302-2130313221331023-1101302323021232-3102200232321122-2121120001113300) |
| `description` | [description](resources--service_policy--reference--group-001.md#canonical-1302311212221301-1310103302120232-2132221321101212-2113113121120303-0100132122021012-3312131321311002-3311232011000012-2120113310100221) |
| `disable` | [disable](resources--service_policy--reference--group-001.md#canonical-1133121223221123-1120002202202221-0002112233231100-1120303301103020-1111210301033131-1132130320232301-3322223222112302-0302130112133323) |
| `id` | [ID](resources--service_policy--reference--group-001.md#canonical-1020101210221131-2002323212013003-1113323111303303-1002231322231110-1110200213233023-2212030220102122-2320113310322020-3322102020130323) |
| `labels` | [labels](resources--service_policy--reference--group-001.md#canonical-3132302113200002-3003131101313310-0000300321230130-3102003112321303-3101232001012002-0101122333310220-1112030300220110-1112222012310330) |
| `name` | [name](resources--service_policy--reference--group-001.md#canonical-1312102120133113-1000012300011301-2212322211111313-1020212320212312-2212021032210022-3103313233021100-1221130230303023-2213311002021103) |
| `namespace` | [namespace](resources--service_policy--reference--group-001.md#canonical-3310021011323222-0223223222003333-3222230033321322-2321321232233132-3302312112230100-3003201000020302-2322212003011210-0300302313032131) |
| `rule_list` | [rule_list](resources--service_policy--reference--group-001.md#canonical-1210231001020020-2003112213000033-0122001321300001-0102112010302033-0220031223101333-1311303212120321-2000002302220220-3220012033313130) |
| `rule_list.rules` | [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-0310120133110223-0333200202303103-1223331131100013-3011023100202112-1213123201130331-0010300112020120-0002002231122013-2202212103010103) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](resources--service_policy--reference--group-001.md#canonical-0230311032002310-2102032112120112-2023200102322132-2310322113133112-3222121130312231-3310132000303013-2133301100022332-3233301022110100) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](resources--service_policy--reference--group-001.md#canonical-2213223033003010-3323321321032322-3121230130222022-1211130101200023-2200110133103211-1211110001021222-1310121211321332-0100322011203121) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](resources--service_policy--reference--group-001.md#canonical-1211233011023032-2321000320202212-3132012111311222-0000220331201001-0330130131133002-2332303010231322-0130230322122032-2230331200313022) |
| `rule_list.rules.spec` | [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-3023121133013333-1102330102213132-3332203320113130-2203301321011130-3102223123000331-1231200011033022-2013222033333002-0012300022003303) |
| `rule_list.rules.spec.action` | [rule_list.rules.spec.action](resources--service_policy--reference--group-001.md#canonical-3121232333100200-1201323101301233-1010300112213213-0320300311002320-1322020000001322-2033020001220330-0001030303100300-2320311233233001) |
| `rule_list.rules.spec.any_asn` | [rule_list.rules.spec.any_asn](resources--service_policy--reference--group-001.md#canonical-3201031031122211-3121221200332103-2000210302033231-3110022213222232-2101203221223023-0111310320201312-3002013310322000-3303311333202223) |
| `rule_list.rules.spec.any_client` | [rule_list.rules.spec.any_client](resources--service_policy--reference--group-001.md#canonical-2010321331330322-0300301220123213-0031320001012313-2030112202213333-0101323201222300-1110211110213132-0100222101331031-3111330020303313) |
| `rule_list.rules.spec.any_ip` | [rule_list.rules.spec.any_ip](resources--service_policy--reference--group-001.md#canonical-0133202313123122-2122033223100213-0013313301201220-3032312113213010-0011233131303002-1311301023210120-1101123102121232-0021202332102302) |
| `rule_list.rules.spec.api_group_matcher` | [rule_list.rules.spec.api_group_matcher](resources--service_policy--reference--group-001.md#canonical-3300111010320122-1112033221323132-0233011010201112-0212321302330022-3130203013001023-1003202121102131-0113123230103130-1002020012231211) |
| `rule_list.rules.spec.api_group_matcher.invert_matcher` | [rule_list.rules.spec.api_group_matcher.invert_matcher](resources--service_policy--reference--group-001.md#canonical-0111022111211203-2103332032331222-2032112332113321-1131202300231030-3333020121131030-1302233101203131-2113331312300032-1120123122123110) |
| `rule_list.rules.spec.api_group_matcher.match` | [rule_list.rules.spec.api_group_matcher.match](resources--service_policy--reference--group-001.md#canonical-2203223001013233-0000303110202300-0220203101321330-0123130133100223-2020331310020303-3103000103120202-2210232223113212-1112031132001331) |
| `rule_list.rules.spec.arg_matchers` | [rule_list.rules.spec.arg_matchers](resources--service_policy--reference--group-001.md#canonical-1130310212222121-2230312101133330-2010303330030301-2303212000030232-1102220003120120-0303320120320000-0113202112232313-1032230133231032) |
| `rule_list.rules.spec.arg_matchers.check_not_present` | [rule_list.rules.spec.arg_matchers.check_not_present](resources--service_policy--reference--group-001.md#canonical-2120301202100202-0113031220133011-3031313210132032-3333230132130112-3203121101202232-3221211201301130-0310000211020300-0011002013023303) |
| `rule_list.rules.spec.arg_matchers.check_present` | [rule_list.rules.spec.arg_matchers.check_present](resources--service_policy--reference--group-001.md#canonical-3032112111212202-0202230110023133-3123012032202203-3222100210203113-2300013222131100-1030202023113101-3103302013031301-2230110123112203) |
| `rule_list.rules.spec.arg_matchers.invert_matcher` | [rule_list.rules.spec.arg_matchers.invert_matcher](resources--service_policy--reference--group-001.md#canonical-2011323231013222-3302132301311201-3030101001312031-2102031103213312-3012023233010330-0300013230022001-2330002021022211-1322122331000223) |
| `rule_list.rules.spec.arg_matchers.item` | [rule_list.rules.spec.arg_matchers.item](resources--service_policy--reference--group-001.md#canonical-3003233002221113-3333112023231203-3313031330313310-1130113000321230-3020233111103301-3110131113113201-2230130130201011-2333313232121013) |
| `rule_list.rules.spec.arg_matchers.item.exact_values` | [rule_list.rules.spec.arg_matchers.item.exact_values](resources--service_policy--reference--group-001.md#canonical-3013133302132001-2310100002302003-0132033023011020-2113112230333113-1302030011201332-2211000112213001-0001303000133101-1110212212302233) |
| `rule_list.rules.spec.arg_matchers.item.regex_values` | [rule_list.rules.spec.arg_matchers.item.regex_values](resources--service_policy--reference--group-001.md#canonical-1213031220112002-1012312312333333-3110031330121223-3121111221222211-3203213120210331-0111121313123231-0020133122231213-1101013222320332) |
| `rule_list.rules.spec.arg_matchers.item.transformers` | [rule_list.rules.spec.arg_matchers.item.transformers](resources--service_policy--reference--group-001.md#canonical-0003331321220202-2322221310000223-0120010023311213-1211031103101312-1121211110111033-1122113012320100-3121323110301132-3231221301303021) |
| `rule_list.rules.spec.arg_matchers.name` | [rule_list.rules.spec.arg_matchers.name](resources--service_policy--reference--group-001.md#canonical-3011031033120301-0120122120013011-0010203333320022-0312120130200120-3311312321130203-1213112110303221-2302012303302220-2021013003000010) |
| `rule_list.rules.spec.asn_list` | [rule_list.rules.spec.asn_list](resources--service_policy--reference--group-001.md#canonical-0112220131300210-0000232113301130-1223211111311303-3213010331012012-1030320012010301-3101333000130132-0012230002312010-3111131320010201) |
| `rule_list.rules.spec.asn_list.as_numbers` | [rule_list.rules.spec.asn_list.as_numbers](resources--service_policy--reference--group-001.md#canonical-2131233223131302-3213311010331210-3320021200230231-3123331121200213-1002220212332003-3130331023113002-0012013322031223-1331101133230130) |
| `rule_list.rules.spec.asn_matcher` | [rule_list.rules.spec.asn_matcher](resources--service_policy--reference--group-001.md#canonical-0313232001300103-1123001113133303-0031021200132120-0120213223020223-3212211131200330-3212030313231201-3331332203302112-2212013330103033) |
| `rule_list.rules.spec.asn_matcher.asn_sets` | [rule_list.rules.spec.asn_matcher.asn_sets](resources--service_policy--reference--group-002.md#canonical-3002022331323320-1223010221333030-1232321121220102-1211213203301123-1221001020311203-3030131120033132-1000220012130230-3230210113102001) |
| `rule_list.rules.spec.asn_matcher.asn_sets.kind` | [rule_list.rules.spec.asn_matcher.asn_sets.kind](resources--service_policy--reference--group-002.md#canonical-3013121001110201-3121330022333310-3032312310303200-0012331200121213-3133221220110312-1203313123201022-2202200002121101-3301312321100100) |
| `rule_list.rules.spec.asn_matcher.asn_sets.name` | [rule_list.rules.spec.asn_matcher.asn_sets.name](resources--service_policy--reference--group-002.md#canonical-2200311223230301-3211210233332022-1332110300031323-3211113022023130-3020303210120133-3030022213011200-1121100111330023-1330221113022130) |
| `rule_list.rules.spec.asn_matcher.asn_sets.namespace` | [rule_list.rules.spec.asn_matcher.asn_sets.namespace](resources--service_policy--reference--group-002.md#canonical-2100013200133312-1110123113210012-2121213111021011-2232212322320030-2121323102011120-1301101003122301-2020232301020110-3310300120223220) |
| `rule_list.rules.spec.asn_matcher.asn_sets.tenant` | [rule_list.rules.spec.asn_matcher.asn_sets.tenant](resources--service_policy--reference--group-002.md#canonical-0013301213123311-3032032323110202-3101321112022300-3003312132130330-1221001000130020-1203012033102022-1330233111322330-0331110030212132) |
| `rule_list.rules.spec.asn_matcher.asn_sets.uid` | [rule_list.rules.spec.asn_matcher.asn_sets.uid](resources--service_policy--reference--group-002.md#canonical-2231001300032230-3220021301232210-3002222031133020-2113030200022222-0103100022213302-1033133311103012-0032233201233312-1220303210321122) |
| `rule_list.rules.spec.body_matcher` | [rule_list.rules.spec.body_matcher](resources--service_policy--reference--group-002.md#canonical-1132300332130030-2212330302311231-3031330013112203-1300311030020032-1210003020030112-0232210320320300-1002111312131231-2212011203311320) |
| `rule_list.rules.spec.body_matcher.exact_values` | [rule_list.rules.spec.body_matcher.exact_values](resources--service_policy--reference--group-002.md#canonical-0320100001331130-1330310130320122-2102032302221013-3013103111030331-3013033112332111-3010202220222202-2120321210131101-2322032310122201) |
| `rule_list.rules.spec.body_matcher.regex_values` | [rule_list.rules.spec.body_matcher.regex_values](resources--service_policy--reference--group-002.md#canonical-0323120321021113-0323311220323013-0112312110100202-3312332112103333-0033013213101021-3300202013000110-1032300002001333-0211220323230131) |
| `rule_list.rules.spec.body_matcher.transformers` | [rule_list.rules.spec.body_matcher.transformers](resources--service_policy--reference--group-002.md#canonical-0003222100230331-2202133133320111-3233030222221100-1331032111201332-3010220213030320-2323133101300130-1313130233322020-0233002100223010) |
| `rule_list.rules.spec.bot_action` | [rule_list.rules.spec.bot_action](resources--service_policy--reference--group-002.md#canonical-0033121133202131-2133212303020203-1332123113310000-2021312303302022-3213001013120202-3102223120100231-3133123201011300-2210030303112131) |
| `rule_list.rules.spec.bot_action.bot_skip_processing` | [rule_list.rules.spec.bot_action.bot_skip_processing](resources--service_policy--reference--group-002.md#canonical-2312132211322012-3021201303020020-0031003231301221-1100002310333210-0000033012303022-3212122033101021-0221131330330032-3022033213132302) |
| `rule_list.rules.spec.bot_action.none` | [rule_list.rules.spec.bot_action.none](resources--service_policy--reference--group-002.md#canonical-0130312213301320-2303112020231313-2031231323301301-3320330131120203-1011130332030030-3113013321133122-2033223110113023-2230021331111330) |
| `rule_list.rules.spec.client_name` | [rule_list.rules.spec.client_name](resources--service_policy--reference--group-001.md#canonical-0122000033030321-3322111233303312-0230321130020300-0333230102333221-1132122333032022-2002011130020311-2001020131300113-3000023320031112) |
| `rule_list.rules.spec.client_name_matcher` | [rule_list.rules.spec.client_name_matcher](resources--service_policy--reference--group-002.md#canonical-1122221102210123-2032002011013332-0232112002321231-2130223110233202-0212232233323101-2230300132312311-3232023001011201-0201231012221233) |
| `rule_list.rules.spec.client_name_matcher.exact_values` | [rule_list.rules.spec.client_name_matcher.exact_values](resources--service_policy--reference--group-002.md#canonical-0311313132010320-3200301001320000-2210112133000211-2230023302301130-0022132003301233-3101330031313130-3133010102113213-1301033023321103) |
| `rule_list.rules.spec.client_name_matcher.regex_values` | [rule_list.rules.spec.client_name_matcher.regex_values](resources--service_policy--reference--group-002.md#canonical-2100013303002233-3020112023211021-2313103033111200-0012132012031301-3331130020310001-2013210032012302-2002112321220200-1120030123031011) |
| `rule_list.rules.spec.client_name_matcher.transformers` | [rule_list.rules.spec.client_name_matcher.transformers](resources--service_policy--reference--group-002.md#canonical-0320300110001300-2133011003200221-0202030200323133-1210302233302020-2330320021322130-3312002201203120-0220012120133111-3023331012012020) |
| `rule_list.rules.spec.client_selector` | [rule_list.rules.spec.client_selector](resources--service_policy--reference--group-002.md#canonical-2003111002023221-1103220223332023-3013333202303122-2231213020002130-3001012030303012-0233321330221322-0300321113203130-2002033233332211) |
| `rule_list.rules.spec.client_selector.expressions` | [rule_list.rules.spec.client_selector.expressions](resources--service_policy--reference--group-002.md#canonical-1332221020302120-2032210223023031-0130011330132321-0313130301312022-0020233102232332-0003222310303112-0212110011302122-1302002100102120) |
| `rule_list.rules.spec.cookie_matchers` | [rule_list.rules.spec.cookie_matchers](resources--service_policy--reference--group-002.md#canonical-2112011002113230-2211013332221222-2300302112132333-2010300033210300-3231300000303200-2212222112010320-0332213213332311-1230023222023331) |
| `rule_list.rules.spec.cookie_matchers.check_not_present` | [rule_list.rules.spec.cookie_matchers.check_not_present](resources--service_policy--reference--group-002.md#canonical-2100331203021313-1020101120103220-3230123331101200-1123000203222120-0303120230100112-0302022130001311-3320312132213033-2022032213003132) |
| `rule_list.rules.spec.cookie_matchers.check_present` | [rule_list.rules.spec.cookie_matchers.check_present](resources--service_policy--reference--group-002.md#canonical-3032100303112312-2220033032301231-3210201221223331-1022100012113020-3133023112100331-0000022003010301-0211000322202023-0103122200213112) |
| `rule_list.rules.spec.cookie_matchers.invert_matcher` | [rule_list.rules.spec.cookie_matchers.invert_matcher](resources--service_policy--reference--group-002.md#canonical-3213331130021230-0233233123032033-0332121002120023-3120323302002330-1231030211322023-3211301330203001-1232303033303132-0103133022103122) |
| `rule_list.rules.spec.cookie_matchers.item` | [rule_list.rules.spec.cookie_matchers.item](resources--service_policy--reference--group-002.md#canonical-3231132302133321-0200311113212233-3301231031112302-1133202000230333-3030302012032311-3121012010202020-3202322100300212-3302132301001302) |
| `rule_list.rules.spec.cookie_matchers.item.exact_values` | [rule_list.rules.spec.cookie_matchers.item.exact_values](resources--service_policy--reference--group-002.md#canonical-0022331112020300-3012021202230030-1131121223123110-3203233321233110-3020302003313110-2202331300312023-2300131302132200-0120201000223113) |
| `rule_list.rules.spec.cookie_matchers.item.regex_values` | [rule_list.rules.spec.cookie_matchers.item.regex_values](resources--service_policy--reference--group-002.md#canonical-2310122021011012-3122221312330131-1231233222131212-1131300221002023-2232202002123311-0120011011010030-3212231030001302-0311002211301303) |
| `rule_list.rules.spec.cookie_matchers.item.transformers` | [rule_list.rules.spec.cookie_matchers.item.transformers](resources--service_policy--reference--group-002.md#canonical-3111123013322003-0031333203013332-0032213333211330-3312202011012102-1003100121233100-2312001000310201-3120223013023132-2300031311233223) |
| `rule_list.rules.spec.cookie_matchers.name` | [rule_list.rules.spec.cookie_matchers.name](resources--service_policy--reference--group-002.md#canonical-1311021021022212-1020220132233122-3333023203332220-2312210312123113-2011211110230313-2232311321203310-3121210010200133-2332120313123100) |
| `rule_list.rules.spec.domain_matcher` | [rule_list.rules.spec.domain_matcher](resources--service_policy--reference--group-002.md#canonical-3013030122221221-0321233330320223-3223222031223102-2010233030133031-3010100103013111-3021230133303300-2203123232133122-3131020003001022) |
| `rule_list.rules.spec.domain_matcher.exact_values` | [rule_list.rules.spec.domain_matcher.exact_values](resources--service_policy--reference--group-002.md#canonical-3202131312112222-1323211210231002-0111002131012332-0203010203303122-3102310030202222-2320300310300132-0201022100220013-3120011321113133) |
| `rule_list.rules.spec.domain_matcher.regex_values` | [rule_list.rules.spec.domain_matcher.regex_values](resources--service_policy--reference--group-002.md#canonical-2213010313213022-0210311331222022-1323013113310031-3200011223132010-0203322313022031-3022301101210032-1302231030210132-2321230332202102) |
| `rule_list.rules.spec.domain_matcher.transformers` | [rule_list.rules.spec.domain_matcher.transformers](resources--service_policy--reference--group-002.md#canonical-3301301003231322-3300012022231202-0203302333223321-2013030110222013-3021213232310311-3103102131033032-1300212331312220-1013110010320323) |
| `rule_list.rules.spec.expiration_timestamp` | [rule_list.rules.spec.expiration_timestamp](resources--service_policy--reference--group-001.md#canonical-1013211313113212-1100103301202121-1321121100210222-0221132322010130-1231210310120030-2013222130323111-1233231221102101-3332202313023121) |
| `rule_list.rules.spec.headers` | [rule_list.rules.spec.headers](resources--service_policy--reference--group-002.md#canonical-0213120231022333-1331110113330110-3302230212201103-0212013302220011-3200303113333210-0202311002002032-0113132010112303-1022031230212220) |
| `rule_list.rules.spec.headers.check_not_present` | [rule_list.rules.spec.headers.check_not_present](resources--service_policy--reference--group-002.md#canonical-3222332233120023-0031033121231313-1002303013323121-0213131030313002-3332130123220130-1023010232300312-3230321321102030-3020012302110000) |
| `rule_list.rules.spec.headers.check_present` | [rule_list.rules.spec.headers.check_present](resources--service_policy--reference--group-002.md#canonical-1333113230013120-2200102130022102-1031331222030131-2211231023120313-3103001032231210-3223321012102121-1020320111211232-3212212130300100) |
| `rule_list.rules.spec.headers.invert_matcher` | [rule_list.rules.spec.headers.invert_matcher](resources--service_policy--reference--group-002.md#canonical-0220313230022012-3103100122130010-1022001130120230-2112001330102211-0330213101233313-0313213301211200-1110001300003012-3000230233032322) |
| `rule_list.rules.spec.headers.item` | [rule_list.rules.spec.headers.item](resources--service_policy--reference--group-002.md#canonical-1312310200301223-3210102303211310-3223031201102033-0210102021001002-3321012000031100-0133310132133223-1013201001110201-0330311223023101) |
| `rule_list.rules.spec.headers.item.exact_values` | [rule_list.rules.spec.headers.item.exact_values](resources--service_policy--reference--group-002.md#canonical-3123333112013232-3100030201003232-0230110130123223-2003332302032213-1221231111110210-1131200223312001-0333100200333213-2231220331232230) |
| `rule_list.rules.spec.headers.item.regex_values` | [rule_list.rules.spec.headers.item.regex_values](resources--service_policy--reference--group-002.md#canonical-3302120023100222-1113133221031002-3320330201233103-0222302302320212-0220320031020100-3200133232320011-1222001321333302-3312033300203013) |
| `rule_list.rules.spec.headers.item.transformers` | [rule_list.rules.spec.headers.item.transformers](resources--service_policy--reference--group-002.md#canonical-2223200003130022-2033321033031100-1220133221033333-0002312210313303-1200232300323300-3330110100322222-3122201003010133-3200323000322130) |
| `rule_list.rules.spec.headers.name` | [rule_list.rules.spec.headers.name](resources--service_policy--reference--group-002.md#canonical-1131322331000101-2232300020230331-3102210220030021-2201111033321100-1211023321111013-1202022121330103-1011302012131202-1120012312032201) |
| `rule_list.rules.spec.http_method` | [rule_list.rules.spec.http_method](resources--service_policy--reference--group-002.md#canonical-1021002231113020-0212311202231001-0023222012220312-0121323112302230-2012230031310122-1013021011230333-3002300311231320-2033033231313113) |
| `rule_list.rules.spec.http_method.invert_matcher` | [rule_list.rules.spec.http_method.invert_matcher](resources--service_policy--reference--group-002.md#canonical-3000121122320203-1013130121101033-1102012310033200-2222300232000333-0213002111102211-3132232230221211-0123133323020330-2133231313120211) |
| `rule_list.rules.spec.http_method.methods` | [rule_list.rules.spec.http_method.methods](resources--service_policy--reference--group-002.md#canonical-0331131330322111-0232323132211210-0320012133202000-1020121000233122-1011320022212313-0300000310001231-2121123210200213-1121000201330322) |
| `rule_list.rules.spec.ip_matcher` | [rule_list.rules.spec.ip_matcher](resources--service_policy--reference--group-002.md#canonical-2300301111202111-2100311132323011-1112011132022031-2110302023001212-2101210122110321-2230231001030221-0100221012300200-3322220122330323) |
| `rule_list.rules.spec.ip_matcher.invert_matcher` | [rule_list.rules.spec.ip_matcher.invert_matcher](resources--service_policy--reference--group-002.md#canonical-2130232212313311-1322122202121310-1332122210303230-0331203300330123-0212220331232331-0213103322030222-2122111200020002-0213130223000020) |
| `rule_list.rules.spec.ip_matcher.prefix_sets` | [rule_list.rules.spec.ip_matcher.prefix_sets](resources--service_policy--reference--group-002.md#canonical-3013331101120231-3330123011110330-1320232201033322-0201100213312011-3130110303121120-0322130232213030-1132001221222210-3120323203331031) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.kind` | [rule_list.rules.spec.ip_matcher.prefix_sets.kind](resources--service_policy--reference--group-002.md#canonical-0022322313323221-0102131330313210-0011312323030003-1130022230010120-0111202222201122-1031121233223021-2332003332312021-0332332222033202) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.name` | [rule_list.rules.spec.ip_matcher.prefix_sets.name](resources--service_policy--reference--group-002.md#canonical-3321030231013221-2211002013033121-2230231223203300-2022321010100332-1133231220122122-3131330312033313-2223100030233102-2020210103111021) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.namespace` | [rule_list.rules.spec.ip_matcher.prefix_sets.namespace](resources--service_policy--reference--group-002.md#canonical-0321112330133013-2311230200322112-0122002323111122-1121132033030223-3302321122203202-3122013203100013-0322332212111030-1103220130003030) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.tenant` | [rule_list.rules.spec.ip_matcher.prefix_sets.tenant](resources--service_policy--reference--group-002.md#canonical-1210031232301203-2111013202213111-3010233033313203-3233102312210113-3122221322321210-3333111200123233-0103003030221233-3210031122131000) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.uid` | [rule_list.rules.spec.ip_matcher.prefix_sets.uid](resources--service_policy--reference--group-002.md#canonical-0112120110031132-1220100221221130-3202122120123313-1231330213103130-0303103133121012-0212301032011021-0102232222032200-2322202200221300) |
| `rule_list.rules.spec.ip_prefix_list` | [rule_list.rules.spec.ip_prefix_list](resources--service_policy--reference--group-002.md#canonical-2312122221023221-3203212000322101-0313001101033121-0310111103332002-0030010013303333-0331300011022321-1311311213320200-2030123213321312) |
| `rule_list.rules.spec.ip_prefix_list.invert_match` | [rule_list.rules.spec.ip_prefix_list.invert_match](resources--service_policy--reference--group-002.md#canonical-2203220011330210-0231203013231030-0013312020001002-2113113203301221-0100322313210312-1022113202001123-1310120213231220-0122322230302033) |
| `rule_list.rules.spec.ip_prefix_list.ip_prefixes` | [rule_list.rules.spec.ip_prefix_list.ip_prefixes](resources--service_policy--reference--group-002.md#canonical-0023233302301232-3333103103323230-2113200133113000-2122332230133231-0010022013223011-2011023002311111-2100333331210112-0121223013020001) |
| `rule_list.rules.spec.ip_threat_category_list` | [rule_list.rules.spec.ip_threat_category_list](resources--service_policy--reference--group-002.md#canonical-1312030330312331-3321100332010323-2022233012332322-1000100030230000-3101010320232101-3012122020033300-1021022213303332-0010322032221300) |
| `rule_list.rules.spec.ip_threat_category_list.ip_threat_categories` | [rule_list.rules.spec.ip_threat_category_list.ip_threat_categories](resources--service_policy--reference--group-002.md#canonical-3300001222002020-2101003021130010-2212031201222311-2203211023222302-3113331330303330-0123121112032203-2001323113003113-0300110211312133) |
| `rule_list.rules.spec.ja4_tls_fingerprint` | [rule_list.rules.spec.ja4_tls_fingerprint](resources--service_policy--reference--group-002.md#canonical-2200331211100333-2110313001203310-0011223222311212-2333230223322132-2302012110023130-1103003310303213-3012210111302003-3303132020211123) |
| `rule_list.rules.spec.ja4_tls_fingerprint.exact_values` | [rule_list.rules.spec.ja4_tls_fingerprint.exact_values](resources--service_policy--reference--group-002.md#canonical-0100320003130232-2231313322333333-0032032221301002-0122022312300202-2032001001031022-3330011203233003-0203210320112111-2133203202333131) |
| `rule_list.rules.spec.jwt_claims` | [rule_list.rules.spec.jwt_claims](resources--service_policy--reference--group-002.md#canonical-1000201331321112-2103012033023310-1122030000010121-1000300302013312-1101332122310000-3011223330020211-2021133023212232-0032231221311110) |
| `rule_list.rules.spec.jwt_claims.check_not_present` | [rule_list.rules.spec.jwt_claims.check_not_present](resources--service_policy--reference--group-002.md#canonical-1000010230030210-0322122332332233-3101302030202331-3013001000010232-1203032210033133-2231202001333102-1003311130101020-1001233001201322) |
| `rule_list.rules.spec.jwt_claims.check_present` | [rule_list.rules.spec.jwt_claims.check_present](resources--service_policy--reference--group-002.md#canonical-1333301220120303-2000323100113200-2003313311313303-3231302232010213-2231120310133101-1331200031000031-0122032003132103-0123113230223321) |
| `rule_list.rules.spec.jwt_claims.invert_matcher` | [rule_list.rules.spec.jwt_claims.invert_matcher](resources--service_policy--reference--group-002.md#canonical-0031133331031031-1321010301330100-1033020011302112-1222022013131011-1131010201320211-0303212012003320-3122332313200113-2022233012011331) |
| `rule_list.rules.spec.jwt_claims.item` | [rule_list.rules.spec.jwt_claims.item](resources--service_policy--reference--group-002.md#canonical-1210213111000001-1123321320023333-2123323131200330-1312013303103312-0100303013001310-0200321010222023-0333103010311331-2032220233001202) |
| `rule_list.rules.spec.jwt_claims.item.exact_values` | [rule_list.rules.spec.jwt_claims.item.exact_values](resources--service_policy--reference--group-002.md#canonical-1213220120020323-0033203133311133-0302233311123220-0202110233102231-1222132012112023-2111121033302131-1302313213301113-0011113101003112) |
| `rule_list.rules.spec.jwt_claims.item.regex_values` | [rule_list.rules.spec.jwt_claims.item.regex_values](resources--service_policy--reference--group-002.md#canonical-3332201012122120-1223101313313001-1233310312022310-3331032030220313-1030201333301231-3231321322121010-2120112132030103-1020103021130222) |
| `rule_list.rules.spec.jwt_claims.item.transformers` | [rule_list.rules.spec.jwt_claims.item.transformers](resources--service_policy--reference--group-002.md#canonical-2013011113222133-3201301120112011-1223202113311111-3321003003133203-2230322133123313-3110030033320103-0313003202210230-1132331012323222) |
| `rule_list.rules.spec.jwt_claims.name` | [rule_list.rules.spec.jwt_claims.name](resources--service_policy--reference--group-002.md#canonical-3223111323113330-2112123222131122-1303322201210103-2310033012221230-1131111012101103-0110302310223221-1220320200033112-0013131232301120) |
| `rule_list.rules.spec.label_matcher` | [rule_list.rules.spec.label_matcher](resources--service_policy--reference--group-002.md#canonical-2231321110123331-2012213321010303-2303111210130201-1332001113000330-3100302101222213-0301023003010033-1201110222203211-2203003013121223) |
| `rule_list.rules.spec.label_matcher.keys` | [rule_list.rules.spec.label_matcher.keys](resources--service_policy--reference--group-002.md#canonical-2322031213102022-0121220000021030-3322103300103101-3200223230103012-3230320110330131-0020130331230330-0012221311211031-0323311112121300) |
| `rule_list.rules.spec.log_rule_evaluation` | [rule_list.rules.spec.log_rule_evaluation](resources--service_policy--reference--group-001.md#canonical-1332010122103030-1301232132210311-3123233112322122-3202013022321331-1310200212220303-1332112233002301-0013230100221323-1213122231131320) |
| `rule_list.rules.spec.mum_action` | [rule_list.rules.spec.mum_action](resources--service_policy--reference--group-002.md#canonical-3130111322110332-1003313330213002-0202120002122032-1332322232212331-1130231322331111-0210221103222231-0222131231130223-1022023311230020) |
| `rule_list.rules.spec.mum_action.default` | [rule_list.rules.spec.mum_action.default](resources--service_policy--reference--group-002.md#canonical-1203210221011221-1231322030122330-3011213013320322-1023213231300112-1113310122001202-2323212221133313-1033121310302010-0321210112030102) |
| `rule_list.rules.spec.mum_action.skip_processing` | [rule_list.rules.spec.mum_action.skip_processing](resources--service_policy--reference--group-002.md#canonical-1220022123030333-3220201200011133-2213221212020322-2032213023303312-3031203203210220-2303231103121320-0011231230023221-2033211322012222) |
| `rule_list.rules.spec.path` | [rule_list.rules.spec.path](resources--service_policy--reference--group-002.md#canonical-3103230311331210-0300023330003333-3212322321221123-2330313301220312-2230332311233331-3120022111220210-2201010211121101-2013010022000210) |
| `rule_list.rules.spec.path.encoded_path_matcher` | [rule_list.rules.spec.path.encoded_path_matcher](resources--service_policy--reference--group-002.md#canonical-3222322321200301-1332010330122220-3021301301320111-0312300000331213-1023212310013120-0320203012331023-1012121200223120-3123313333013101) |
| `rule_list.rules.spec.path.exact_values` | [rule_list.rules.spec.path.exact_values](resources--service_policy--reference--group-002.md#canonical-3210333102012033-3131323100311030-3001002321213112-2201333211232032-3301321110102133-3020122200011001-3022321002333323-2230302032021222) |
| `rule_list.rules.spec.path.invert_matcher` | [rule_list.rules.spec.path.invert_matcher](resources--service_policy--reference--group-002.md#canonical-1003023133300002-1120000233122310-1330220211201110-1030331331010323-2102222332231122-1323222323101020-2101031011223333-1222202321301221) |
| `rule_list.rules.spec.path.prefix_values` | [rule_list.rules.spec.path.prefix_values](resources--service_policy--reference--group-002.md#canonical-2331002311030320-2230313213123023-1013222322211130-3312031331210313-3113130000023231-2333120323220101-2332320333311213-3013221130100331) |
| `rule_list.rules.spec.path.regex_values` | [rule_list.rules.spec.path.regex_values](resources--service_policy--reference--group-002.md#canonical-1123122013100312-0311103102303323-1300000223011211-3202333313123230-3103231110302321-0220112133321301-2023002102011032-2121231032221013) |
| `rule_list.rules.spec.path.suffix_values` | [rule_list.rules.spec.path.suffix_values](resources--service_policy--reference--group-002.md#canonical-0100023222031213-3010233320022000-2001302010201112-3032010123310202-0023300220222220-1303230021021010-1213122032002022-1122300330213003) |
| `rule_list.rules.spec.path.transformers` | [rule_list.rules.spec.path.transformers](resources--service_policy--reference--group-002.md#canonical-1230231102330211-2311300033000311-0322220020110023-0110303301233101-1222020333212303-2123300022021211-0322001210320332-2230013113320212) |
| `rule_list.rules.spec.port_matcher` | [rule_list.rules.spec.port_matcher](resources--service_policy--reference--group-002.md#canonical-0322300223131313-2202330032010311-1302231231102213-2323120213102321-3210322230000132-0221322233331332-1201020132320013-1220332013322130) |
| `rule_list.rules.spec.port_matcher.invert_matcher` | [rule_list.rules.spec.port_matcher.invert_matcher](resources--service_policy--reference--group-002.md#canonical-1102222233230023-2003132031021323-1331003211012101-0110330011132130-2321311033100330-2113121032312211-0133000132023030-3203000001230311) |
| `rule_list.rules.spec.port_matcher.ports` | [rule_list.rules.spec.port_matcher.ports](resources--service_policy--reference--group-002.md#canonical-2120332231000300-3331130013300021-3313323011221211-1320223000113102-1331220121321220-0000100103311022-1311001313033331-0121123032101302) |
| `rule_list.rules.spec.query_params` | [rule_list.rules.spec.query_params](resources--service_policy--reference--group-002.md#canonical-3313303312311300-0033303303001212-3310323232203103-0021112110221120-0121003322120210-0231003023330323-2330221000013103-1220103001003120) |
| `rule_list.rules.spec.query_params.check_not_present` | [rule_list.rules.spec.query_params.check_not_present](resources--service_policy--reference--group-002.md#canonical-2213311101101232-1311032112333231-1032121203300013-1302132020312221-1203331211003003-1221331121130103-1300110311112211-3202110200122122) |
| `rule_list.rules.spec.query_params.check_present` | [rule_list.rules.spec.query_params.check_present](resources--service_policy--reference--group-002.md#canonical-3320122212110131-3133133000022033-0012111013212211-2330221333212002-3302320300303030-3033112132321121-0332211011122123-3320112223302312) |
| `rule_list.rules.spec.query_params.invert_matcher` | [rule_list.rules.spec.query_params.invert_matcher](resources--service_policy--reference--group-002.md#canonical-2000211101313102-1233220222331033-2003322310110202-3111112301010102-2012023233330220-3021122310023100-2021321111021233-1232033030210312) |
| `rule_list.rules.spec.query_params.item` | [rule_list.rules.spec.query_params.item](resources--service_policy--reference--group-002.md#canonical-1202132022201203-0312121113221100-3102011200212320-2210112330323031-3101111001102213-2233332203130333-2211103201021122-2103032021111002) |
| `rule_list.rules.spec.query_params.item.exact_values` | [rule_list.rules.spec.query_params.item.exact_values](resources--service_policy--reference--group-002.md#canonical-2102110330320230-1312321310120202-1232201112111230-3220031122121120-3323010001110100-2223120220023210-3312130020112121-2030033320002211) |
| `rule_list.rules.spec.query_params.item.regex_values` | [rule_list.rules.spec.query_params.item.regex_values](resources--service_policy--reference--group-002.md#canonical-0022222102111200-2031211300222203-1032201022322230-1233203213131222-0020020022103322-0121311000311001-1131011203011131-0111103322033111) |
| `rule_list.rules.spec.query_params.item.transformers` | [rule_list.rules.spec.query_params.item.transformers](resources--service_policy--reference--group-002.md#canonical-3033131211220301-3103011222321111-2202301031012003-3310110130302321-0221311210203002-1012320022321132-1300010230002033-0101210201120220) |
| `rule_list.rules.spec.query_params.key` | [rule_list.rules.spec.query_params.key](resources--service_policy--reference--group-002.md#canonical-3022020032001113-1221011220231301-0212001222333231-1111010230033102-2103310321133333-0103323013312032-0033303032320332-0021120233131111) |
| `rule_list.rules.spec.request_constraints` | [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-3333302322101323-0201111313220103-0020221122010103-1212101111101001-2023210312120022-1132203131330102-3000302301311000-1110312003233023) |
| `rule_list.rules.spec.request_constraints.max_cookie_count_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_count_exceeds](resources--service_policy--reference--group-002.md#canonical-3203200013110120-3210122032320032-1023211320310003-3201203132002321-0302323111131223-2320111221303221-2221132020310313-2312002030002020) |
| `rule_list.rules.spec.request_constraints.max_cookie_count_none` | [rule_list.rules.spec.request_constraints.max_cookie_count_none](resources--service_policy--reference--group-002.md#canonical-2033331210013202-1301200033023110-3033023210233200-0302110203221311-1220113021303013-0220112200201323-1202133232320022-3201301133031123) |
| `rule_list.rules.spec.request_constraints.max_cookie_key_size_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_key_size_exceeds](resources--service_policy--reference--group-002.md#canonical-0333323330233211-2131021103331213-3001003020031221-2031010111122201-2313223133223103-1332011330000301-2223131111330000-0102332013303223) |
| `rule_list.rules.spec.request_constraints.max_cookie_key_size_none` | [rule_list.rules.spec.request_constraints.max_cookie_key_size_none](resources--service_policy--reference--group-002.md#canonical-1322302200033123-1131111212300303-2032233013003123-0321101002232202-2111302012322112-3322311310310010-0322012121223303-3211120220000103) |
| `rule_list.rules.spec.request_constraints.max_cookie_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_value_size_exceeds](resources--service_policy--reference--group-002.md#canonical-0130201012313000-0021311132313200-0111023020000112-1332313212211123-3331001003310231-0031333113000223-2221012302000012-2100200110113201) |
| `rule_list.rules.spec.request_constraints.max_cookie_value_size_none` | [rule_list.rules.spec.request_constraints.max_cookie_value_size_none](resources--service_policy--reference--group-002.md#canonical-3100112101300020-0201221301112022-2302330012220132-2211112023333231-2133310233130032-0031201202130210-2012202300310311-0210003230001023) |
| `rule_list.rules.spec.request_constraints.max_header_count_exceeds` | [rule_list.rules.spec.request_constraints.max_header_count_exceeds](resources--service_policy--reference--group-002.md#canonical-1112210022221221-2313233002012221-0102022213231231-2203123232012013-0213021221201311-0132133013310021-0033023123323021-2303100300203331) |
| `rule_list.rules.spec.request_constraints.max_header_count_none` | [rule_list.rules.spec.request_constraints.max_header_count_none](resources--service_policy--reference--group-002.md#canonical-3231131322130322-2322120113213002-3323331321003102-3233000130312032-3302111222211022-0000322033033110-3133033133030003-0330320331322213) |
| `rule_list.rules.spec.request_constraints.max_header_key_size_exceeds` | [rule_list.rules.spec.request_constraints.max_header_key_size_exceeds](resources--service_policy--reference--group-002.md#canonical-0113100021100122-3032110310322311-3233101200221210-1221111302321200-3033011113120202-2102011103133213-2322321202320202-3311302303122313) |
| `rule_list.rules.spec.request_constraints.max_header_key_size_none` | [rule_list.rules.spec.request_constraints.max_header_key_size_none](resources--service_policy--reference--group-002.md#canonical-2311321222221211-2202013013113133-1210132321200222-2333313033032131-1212032303301333-1203320020310031-0003112033312031-1002332221110221) |
| `rule_list.rules.spec.request_constraints.max_header_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_header_value_size_exceeds](resources--service_policy--reference--group-002.md#canonical-1111023220010330-2302131021033122-0212100322130101-0220303212301003-2220211011220330-1110013010010122-1111220311103113-3331211221113100) |
| `rule_list.rules.spec.request_constraints.max_header_value_size_none` | [rule_list.rules.spec.request_constraints.max_header_value_size_none](resources--service_policy--reference--group-002.md#canonical-0200202312023131-3221023120000003-1022133301231332-1032131001023223-0300111321312032-3102330222211120-0111230313322321-0320000122002211) |
| `rule_list.rules.spec.request_constraints.max_parameter_count_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_count_exceeds](resources--service_policy--reference--group-002.md#canonical-3213101212031103-1033001312322000-2111321122111110-0032202123332323-0103032121212313-0323112200122003-0133302222232213-1222201312220130) |
| `rule_list.rules.spec.request_constraints.max_parameter_count_none` | [rule_list.rules.spec.request_constraints.max_parameter_count_none](resources--service_policy--reference--group-002.md#canonical-1011003033220003-2231200213222112-1301323310323030-2201010130331310-3202122333003000-2000311032021302-2213203011033310-2332210032333022) |
| `rule_list.rules.spec.request_constraints.max_parameter_name_size_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_name_size_exceeds](resources--service_policy--reference--group-002.md#canonical-3301222111120221-3303003013202123-3021012001303030-0231023311211002-1022033323021001-3030002100123320-3030120130332111-3030111122111102) |
| `rule_list.rules.spec.request_constraints.max_parameter_name_size_none` | [rule_list.rules.spec.request_constraints.max_parameter_name_size_none](resources--service_policy--reference--group-002.md#canonical-1002200331203103-1221101103003231-3231221331020301-1310323010120031-3331230220333111-0203001303130220-3310300221310330-0100310333320332) |
| `rule_list.rules.spec.request_constraints.max_parameter_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_value_size_exceeds](resources--service_policy--reference--group-002.md#canonical-1332233100201321-3321120120321310-1010000310013023-1132033123201110-1323232202003300-0000103302212012-1012111203010120-1310233330332123) |
| `rule_list.rules.spec.request_constraints.max_parameter_value_size_none` | [rule_list.rules.spec.request_constraints.max_parameter_value_size_none](resources--service_policy--reference--group-002.md#canonical-1232133101130103-2132120111110022-1130120331112013-2303013230120131-3111332322200001-3002103222212012-1320033011210231-3221001232333330) |
| `rule_list.rules.spec.request_constraints.max_query_size_exceeds` | [rule_list.rules.spec.request_constraints.max_query_size_exceeds](resources--service_policy--reference--group-002.md#canonical-2313203203310333-0021121100033120-2210320220010032-0231103221333313-3302023232323011-0111120231033211-2203202023222332-3210022022020113) |
| `rule_list.rules.spec.request_constraints.max_query_size_none` | [rule_list.rules.spec.request_constraints.max_query_size_none](resources--service_policy--reference--group-002.md#canonical-2220223202000320-2313110303213033-1011313210202330-0032032201011313-0033131310303033-2321033102112233-1132323233001102-1103103023332223) |
| `rule_list.rules.spec.request_constraints.max_request_line_size_exceeds` | [rule_list.rules.spec.request_constraints.max_request_line_size_exceeds](resources--service_policy--reference--group-002.md#canonical-1111333020123223-2003023013001313-1303133323210310-2213012202322203-1023032101301031-0222220020330311-0302102121012331-1311320300300331) |
| `rule_list.rules.spec.request_constraints.max_request_line_size_none` | [rule_list.rules.spec.request_constraints.max_request_line_size_none](resources--service_policy--reference--group-002.md#canonical-1321210222121100-2232102302000320-1201232200012221-1022201300033120-0212223323111031-1012300122330001-3122002210110322-2300103212232203) |
| `rule_list.rules.spec.request_constraints.max_request_size_exceeds` | [rule_list.rules.spec.request_constraints.max_request_size_exceeds](resources--service_policy--reference--group-002.md#canonical-3033133010232202-0232321002220220-0112102333333230-0121110000112013-1100330221323020-0231312323002010-0013000002310332-3303001211110031) |
| `rule_list.rules.spec.request_constraints.max_request_size_none` | [rule_list.rules.spec.request_constraints.max_request_size_none](resources--service_policy--reference--group-002.md#canonical-1203310032203211-3200210000223202-1202030320200112-1133330220010203-1310011202111322-3320130300131030-1020200313321202-1223211320331332) |
| `rule_list.rules.spec.request_constraints.max_url_size_exceeds` | [rule_list.rules.spec.request_constraints.max_url_size_exceeds](resources--service_policy--reference--group-002.md#canonical-3112020102101133-0131033313220321-2110030302222023-0012111102232002-0222232132322312-1022020002330202-2223032032030221-0100130332013213) |
| `rule_list.rules.spec.request_constraints.max_url_size_none` | [rule_list.rules.spec.request_constraints.max_url_size_none](resources--service_policy--reference--group-002.md#canonical-3331101010322200-1032332023221003-1013000330313113-2302113330213310-2333323333113030-0222003132132002-1122031332000212-3101320130031302) |
| `rule_list.rules.spec.segment_policy` | [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-002.md#canonical-3123012000130222-0021031033231122-1021210223003202-0321331233310102-0310003132123300-2220102021201131-2300100313011321-0220311300302121) |
| `rule_list.rules.spec.segment_policy.dst_any` | [rule_list.rules.spec.segment_policy.dst_any](resources--service_policy--reference--group-002.md#canonical-1012233013201321-0113011332231330-0302110123312323-3023220002210313-2113120201231301-1120132132033200-2332013221313031-1323113322120203) |
| `rule_list.rules.spec.segment_policy.dst_segments` | [rule_list.rules.spec.segment_policy.dst_segments](resources--service_policy--reference--group-002.md#canonical-0132200022132130-3300330010322312-1032121303001121-2130111312331302-3102000010312323-0333100201100331-0310212123312312-3110131320102302) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments` | [rule_list.rules.spec.segment_policy.dst_segments.segments](resources--service_policy--reference--group-002.md#canonical-2332110131120102-2220301302310211-0031031223031213-2301301133122223-3320123101012320-2323001302332220-1030120300230221-2321030013100123) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.name` | [rule_list.rules.spec.segment_policy.dst_segments.segments.name](resources--service_policy--reference--group-002.md#canonical-0020203332031121-1111112312110101-0031013003001013-1311213202121032-0333310200300331-1121113103202301-1222220021311031-3303233113312001) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.namespace` | [rule_list.rules.spec.segment_policy.dst_segments.segments.namespace](resources--service_policy--reference--group-002.md#canonical-0123001233200023-3120321003310010-2022330322100231-0233220300022322-2111201332131331-0210300013121022-1222021130130331-2101123330331311) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.tenant` | [rule_list.rules.spec.segment_policy.dst_segments.segments.tenant](resources--service_policy--reference--group-002.md#canonical-2200112020130223-3310233030031100-3100232310033212-2101310320220112-3211333332312110-1120003311113321-2233321320310103-2022111302131002) |
| `rule_list.rules.spec.segment_policy.intra_segment` | [rule_list.rules.spec.segment_policy.intra_segment](resources--service_policy--reference--group-002.md#canonical-1210033002102202-1012321302012001-1221021120312231-0210110320120122-1301110110203120-1220333211010130-1313010111312021-0223323120003321) |
| `rule_list.rules.spec.segment_policy.src_any` | [rule_list.rules.spec.segment_policy.src_any](resources--service_policy--reference--group-002.md#canonical-1021110201210013-2301132113303322-1212121222233110-1330302030322332-0131130002000303-1022331220210313-3023002332330031-2330102232022201) |
| `rule_list.rules.spec.segment_policy.src_segments` | [rule_list.rules.spec.segment_policy.src_segments](resources--service_policy--reference--group-002.md#canonical-3332231220130101-1011320302021122-1303123303132300-0311302332120211-2132013011112123-2210133031301012-2210112211011003-0032013310233112) |
| `rule_list.rules.spec.segment_policy.src_segments.segments` | [rule_list.rules.spec.segment_policy.src_segments.segments](resources--service_policy--reference--group-002.md#canonical-0102023111220023-0000302012323012-1121022023130210-1303022321020333-2130213231100210-1301213210011131-2201020001200002-3003302202213232) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.name` | [rule_list.rules.spec.segment_policy.src_segments.segments.name](resources--service_policy--reference--group-002.md#canonical-0332131131300213-0301113323313102-2320303301300302-3311002230202122-0120031101330203-2121121022333231-2001120002132013-2220122002122233) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.namespace` | [rule_list.rules.spec.segment_policy.src_segments.segments.namespace](resources--service_policy--reference--group-002.md#canonical-0220020133112222-1003310100122112-0010320333012211-0212320110000111-3203133101030311-2011303103301130-3333100202203011-1233013013210321) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.tenant` | [rule_list.rules.spec.segment_policy.src_segments.segments.tenant](resources--service_policy--reference--group-002.md#canonical-1112111121033002-0220100101230130-3312331303110222-3302102001312012-2222300323302333-1300203332303132-2112021123121200-0212101311120103) |
| `rule_list.rules.spec.tls_fingerprint_matcher` | [rule_list.rules.spec.tls_fingerprint_matcher](resources--service_policy--reference--group-002.md#canonical-2211302223202211-0133312332123122-3001120111203221-2001331203120000-2122122033200113-0203222022113013-3310203232101331-2101130030200002) |
| `rule_list.rules.spec.tls_fingerprint_matcher.classes` | [rule_list.rules.spec.tls_fingerprint_matcher.classes](resources--service_policy--reference--group-002.md#canonical-0331220121023231-0021112103101300-0031313222021121-3210131000332312-2331212321201033-0321332301010021-1201203310300313-0303030100323313) |
| `rule_list.rules.spec.tls_fingerprint_matcher.exact_values` | [rule_list.rules.spec.tls_fingerprint_matcher.exact_values](resources--service_policy--reference--group-002.md#canonical-1320002201100011-3203312220113103-2312130030332322-0302001320032322-3311013320011030-1330311112132300-1323120212100113-0001120200032202) |
| `rule_list.rules.spec.tls_fingerprint_matcher.excluded_values` | [rule_list.rules.spec.tls_fingerprint_matcher.excluded_values](resources--service_policy--reference--group-002.md#canonical-0131312323300001-1110011110132013-3101210201212221-0212001023323310-3011303221022330-1220301200110033-2003201232202112-0010001321232023) |
| `rule_list.rules.spec.user_identity_matcher` | [rule_list.rules.spec.user_identity_matcher](resources--service_policy--reference--group-002.md#canonical-1232010221230100-0121021000010310-2213323203322133-0302300133131212-2033103123311120-0012331103220112-1112122302102113-0011130020320302) |
| `rule_list.rules.spec.user_identity_matcher.exact_values` | [rule_list.rules.spec.user_identity_matcher.exact_values](resources--service_policy--reference--group-002.md#canonical-1321310003332203-0023133101222121-3312023221111313-3123033321102023-3112220231302223-1031201210222330-3003333102230021-0011001103133301) |
| `rule_list.rules.spec.user_identity_matcher.regex_values` | [rule_list.rules.spec.user_identity_matcher.regex_values](resources--service_policy--reference--group-002.md#canonical-2033313221121303-2030322301120303-2030310003323020-0210012222022101-3002231102231012-2323020331211031-2031302201102303-0221121113033230) |
| `rule_list.rules.spec.waf_action` | [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-2323131031010332-3002023002320130-2322222211000312-3130213100113011-2320002112103301-2101033310021200-1130113033130203-0233221333232023) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control` | [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-3230320101223013-1011031223033032-0101122123300100-3100203130020222-2332330322003133-3033201132303311-2212003122110223-2113023230100000) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts](resources--service_policy--reference--group-003.md#canonical-3203111012121003-3002323231231031-2230103200131210-1113000133301133-2333002220003102-1320230003323300-0213021033331300-0022112312120100) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context](resources--service_policy--reference--group-003.md#canonical-2233211310100100-1030322201001210-0323223223213122-1022322332310120-0113000201323230-1320223132001333-0233202331000013-1203220332032300) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name](resources--service_policy--reference--group-003.md#canonical-1301331011100031-2132301110033000-3313120300231121-1100330301201021-0020030022203020-2023103330122233-2101322211332013-1023233320232200) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](resources--service_policy--reference--group-003.md#canonical-0211231312311010-0130013111211202-2012300323311011-1322213322330330-1210111222303312-3112033330133302-2022033203003321-2213310301213103) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts](resources--service_policy--reference--group-003.md#canonical-3300111112133121-0310032000322201-3013331110020120-2302211202200011-1320321012202131-1123103101113102-2331200000102300-2102322321020130) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](resources--service_policy--reference--group-003.md#canonical-2000131203123232-0320021223212232-3130133213210003-3313113003003033-1120103000001011-0123300020231310-0221230101002223-2312002313111003) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts](resources--service_policy--reference--group-003.md#canonical-1013000012311210-3230012300322003-2031223220130102-0133321333310221-2101203023211023-2002202011021323-2203312011002113-3003112002113233) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context](resources--service_policy--reference--group-003.md#canonical-1000001223231030-1132101302013320-1311210102013211-3120113202210103-2330332231000211-2233230301232011-1303030332013231-0110321130010113) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name](resources--service_policy--reference--group-003.md#canonical-3203212010030220-2230123130120200-3110101022210220-3211100030330013-0130133121031221-1121312013031212-2232333013103302-3320023102223303) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id](resources--service_policy--reference--group-003.md#canonical-1330222030012132-1103030312020202-2033310031321100-0301013131001133-3302320311322332-1130311021003322-2111100323133302-2321322111302021) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts](resources--service_policy--reference--group-003.md#canonical-1100221322321231-2212130303000130-2203333230031211-3212022103122302-1320212320000101-3120322232302131-3220111232313130-2213211331131032) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context](resources--service_policy--reference--group-003.md#canonical-1301232333003013-0312112312303321-1112110003213320-0330231101001300-1333102030210100-3131121302101202-2001121313123323-3212101311103000) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name](resources--service_policy--reference--group-003.md#canonical-1303003230231210-0012101312223201-2013011200200300-0000113223012300-3333013231110312-2301120211313211-1111003011233203-2331012000321100) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](resources--service_policy--reference--group-003.md#canonical-3313311302333121-3233332223103332-0200323223120033-3223303330111031-0333332212110131-3021032223223300-3322131203301003-3112021223010323) |
| `rule_list.rules.spec.waf_action.none` | [rule_list.rules.spec.waf_action.none](resources--service_policy--reference--group-003.md#canonical-1130211132010033-0012023333330123-3122312222333202-2303000132033133-2030012230211010-0123300221231223-1013103130011323-3332101101000120) |
| `rule_list.rules.spec.waf_action.waf_skip_processing` | [rule_list.rules.spec.waf_action.waf_skip_processing](resources--service_policy--reference--group-003.md#canonical-0332332021232010-3301000202212332-3313331233100032-0132212001212200-3331300332220011-2032001311031303-2120131033213332-1222312203332022) |
| `server_name` | [server_name](resources--service_policy--reference--group-001.md#canonical-0133223010313332-1110211023313330-0010322022330102-0322023303232002-3232002130100231-3032013031322121-1301220213200220-0101110033331123) |
| `server_name_matcher` | [server_name_matcher](resources--service_policy--reference--group-003.md#canonical-2000320332113331-2121031221100100-1221133200122002-3322012030312113-3103022233103020-1100222220211102-1020331122111223-0203320320001212) |
| `server_name_matcher.exact_values` | [server_name_matcher.exact_values](resources--service_policy--reference--group-003.md#canonical-0332001021322311-2131220233010003-1000102232330210-3023121000321032-1122011030103022-3221023131331303-0222331231030333-2200112232233231) |
| `server_name_matcher.regex_values` | [server_name_matcher.regex_values](resources--service_policy--reference--group-003.md#canonical-1130030021202301-1031122212120132-2131301011322220-3301123133131033-1331010012013311-2320113302032322-1231120123232022-1213102020330010) |
| `server_selector` | [server_selector](resources--service_policy--reference--group-003.md#canonical-2100103232022233-0113101000311301-3031331021032201-3032212030030022-3321011201111130-3223030011011331-2010300313013200-1101302322210021) |
| `server_selector.expressions` | [server_selector.expressions](resources--service_policy--reference--group-003.md#canonical-0223003301030203-2231011011012312-3230302101132103-0203220212030132-0133300202003202-1033110312221230-0011131113030210-0212001131033220) |
| `timeouts` | [timeouts](resources--service_policy--reference--group-003.md#canonical-2332330300313210-3032220230231002-2210001000213311-1120030113311201-3323010032130212-0301322332310212-1223300320310131-1002210230210203) |
| `timeouts.create` | [timeouts.create](resources--service_policy--reference--group-003.md#canonical-0130223333312220-1300211333313231-3021121330032221-2302232003113011-3203233212213213-0200122002320223-2213001103032131-2310210302000030) |
| `timeouts.delete` | [timeouts.delete](resources--service_policy--reference--group-003.md#canonical-1031302123101121-1101013020003133-0032312221131121-1100303013031110-1322133010222220-3001330301302222-1303300033011013-3211003211212222) |
| `timeouts.read` | [timeouts.read](resources--service_policy--reference--group-003.md#canonical-3222332311310031-1300301003212221-1121213321301231-2303133033030033-0131120223312201-0130113013130102-2223121101102113-3010302222133212) |
| `timeouts.update` | [timeouts.update](resources--service_policy--reference--group-003.md#canonical-2223033213323213-0023113311030323-2230232220032313-0212011012110310-2023321002030132-0232001232202000-0121200122002113-2300001313013322) |

<a id="canonical-0323122132121121-2032121222012200-2021300003103320-2200300021120320-3213122032133013-1012030312210210-1131212101302130-1131330232103211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_all_requests` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- allow_all_requests

<a id="canonical-3203302312022132-1033132230002030-2122211232332320-0331003202232001-1301111103130220-3331003033201213-0231110210230311-0011302321201113"></a>

Type: `["object", {}]`. Optional.

\[OneOf: allow\_all\_requests, allow\_list, deny\_all\_requests, deny\_list, rule\_list\]
Configuration parameter for allow all requests.

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

OneOf alternatives in this subsection:

- [allow_all_requests](resources--service_policy--reference--group-001.md#canonical-3203302312022132-1033132230002030-2122211232332320-0331003202232001-1301111103130220-3331003033201213-0231110210230311-0011302321201113)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-1021110011030020-2000303211110001-2211321101133001-2322031003001032-0220201300302131-3232100200132030-1200230000110101-1003100222010023)
- [deny_all_requests](resources--service_policy--reference--group-001.md#canonical-3000321302123202-0112320203222012-3033013210133012-3313121213112032-3300203233300233-1031222223201033-1233002203021101-2001220123301012)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-2130322123310112-0302010020132022-3323113023231331-3210232201121112-2122111023320011-2022222301103232-0130032022112322-1323301103100122)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-1210231001020020-2003112213000033-0122001321300001-0102112010302033-0220031223101333-1311303212120321-2000002302220220-3220012033313130)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all_requests = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103300311300332-3332200122100221-1303003012310332-1301221201020110-2111112310210320-2231312202202012-2331211230323103-2101301311212102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- allow_list

<a id="canonical-1021110011030020-2000303211110001-2211321101133001-2322031003001032-0220201300302131-3232100200132030-1200230000110101-1003100222010023"></a>

Type: `"object"`. single nested block, Optional.

List of sources. A request belongs to this list if it satisfies any of the match criteria.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_deny"),
  validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_next_policy"),
  validators.ConflictingObjectAttributes("default_action_deny",
    "default_action_next_policy")}
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
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

Terraform syntax:

```terraform
allow_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230220031033221-2223010101302010-3201221302033003-3331023233022130-1020110101210020-1131310032020132-2112111210302222-2311033101333032"></a>

### Direct properties for `allow_list`

- [asn_list](resources--service_policy--reference--group-001.md#canonical-3132110210003212-2331132232213331-1112123020302130-1202030332101021-2113013220113122-3103002132013200-2211120020233101-2032322303231033): complete subsection reference.

- [asn_set](resources--service_policy--reference--group-001.md#canonical-3322131022003321-1102103331111032-3000011101310303-3022231110302022-0300320130213212-3132013103330201-0010310300002333-1313301232213332): complete subsection reference.

<a id="canonical-3300031100203020-1110110300122003-1210213200212101-1313100030012032-2103300323120122-0013022102000323-0303023030003101-0011233211033023"></a>

<a id="canonical-1032321333213332-0020302032331010-0333122213211023-2013313102212202-0111200111123302-0003023003122230-3300012201233223-1102102333113121"></a>

#### `allow_list.country_list` property

Type: `["list", "string"]`. Optional.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Addresses that belong to one of the countries in the given list The country is obtained by
performing a lookup for the source IPv4 Address in a GeoIP DB. Possible values are
\`COUNTRY\_NONE\`, \`COUNTRY\_AD\`, \`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`,
\`COUNTRY\_AI\`, \`COUNTRY\_AL\`, \`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`,
\`COUNTRY\_AQ\`, \`COUNTRY\_AR\`, \`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`,
\`COUNTRY\_AW\`, \`COUNTRY\_AX\`, \`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`,
\`COUNTRY\_BD\`, \`COUNTRY\_BE\`, \`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`,
\`COUNTRY\_BI\`, \`COUNTRY\_BJ\`, \`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`,
\`COUNTRY\_BO\`, \`COUNTRY\_BQ\`, \`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`,
\`COUNTRY\_BV\`, \`COUNTRY\_BW\`, \`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`,
\`COUNTRY\_CC\`, \`COUNTRY\_CD\`, \`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`,
\`COUNTRY\_CI\`, \`COUNTRY\_CK\`, \`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`,
\`COUNTRY\_CO\`, \`COUNTRY\_CR\`, \`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`,
\`COUNTRY\_CW\`, \`COUNTRY\_CX\`, \`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`,
\`COUNTRY\_DJ\`, \`COUNTRY\_DK\`, \`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`,
\`COUNTRY\_EC\`, \`COUNTRY\_EE\`, \`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`,
\`COUNTRY\_ES\`, \`COUNTRY\_ET\`, \`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`,
\`COUNTRY\_FM\`, \`COUNTRY\_FO\`, \`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`,
\`COUNTRY\_GD\`, \`COUNTRY\_GE\`, \`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`,
\`COUNTRY\_GI\`, \`COUNTRY\_GL\`, \`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`,
\`COUNTRY\_GQ\`, \`COUNTRY\_GR\`, \`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`,
\`COUNTRY\_GW\`, \`COUNTRY\_GY\`, \`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`,
\`COUNTRY\_HR\`, \`COUNTRY\_HT\`, \`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`,
\`COUNTRY\_IL\`, \`COUNTRY\_IM\`, \`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`,
\`COUNTRY\_IR\`, \`COUNTRY\_IS\`, \`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`,
\`COUNTRY\_JO\`, \`COUNTRY\_JP\`, \`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`,
\`COUNTRY\_KI\`, \`COUNTRY\_KM\`, \`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`,
\`COUNTRY\_KW\`, \`COUNTRY\_KY\`, \`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`,
\`COUNTRY\_LC\`, \`COUNTRY\_LI\`, \`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`,
\`COUNTRY\_LT\`, \`COUNTRY\_LU\`, \`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`,
\`COUNTRY\_MC\`, \`COUNTRY\_MD\`, \`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`,
\`COUNTRY\_MH\`, \`COUNTRY\_MK\`, \`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`,
\`COUNTRY\_MO\`, \`COUNTRY\_MP\`, \`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`,
\`COUNTRY\_MT\`, \`COUNTRY\_MU\`, \`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`,
\`COUNTRY\_MY\`, \`COUNTRY\_MZ\`, \`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`,
\`COUNTRY\_NF\`, \`COUNTRY\_NG\`, \`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`,
\`COUNTRY\_NP\`, \`COUNTRY\_NR\`, \`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`,
\`COUNTRY\_PA\`, \`COUNTRY\_PE\`, \`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`,
\`COUNTRY\_PK\`, \`COUNTRY\_PL\`, \`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`,
\`COUNTRY\_PS\`, \`COUNTRY\_PT\`, \`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`,
\`COUNTRY\_RE\`, \`COUNTRY\_RO\`, \`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`,
\`COUNTRY\_SA\`, \`COUNTRY\_SB\`, \`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`,
\`COUNTRY\_SG\`, \`COUNTRY\_SH\`, \`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`,
\`COUNTRY\_SL\`, \`COUNTRY\_SM\`, \`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`,
\`COUNTRY\_SS\`, \`COUNTRY\_ST\`, \`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`,
\`COUNTRY\_SZ\`, \`COUNTRY\_TC\`, \`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`,
\`COUNTRY\_TH\`, \`COUNTRY\_TJ\`, \`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`,
\`COUNTRY\_TN\`, \`COUNTRY\_TO\`, \`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`,
\`COUNTRY\_TW\`, \`COUNTRY\_TZ\`, \`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`,
\`COUNTRY\_US\`, \`COUNTRY\_UY\`, \`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`,
\`COUNTRY\_VE\`, \`COUNTRY\_VG\`, \`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`,
\`COUNTRY\_WF\`, \`COUNTRY\_WS\`, \`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`,
\`COUNTRY\_YT\`, \`COUNTRY\_ZA\`, \`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_action_allow](resources--service_policy--reference--group-001.md#canonical-3230002222313121-3230132232023100-2132221313103213-0130303232302310-0332001112012001-2331130201000100-3121200023032130-2333230000221102): complete subsection reference.

- [default_action_deny](resources--service_policy--reference--group-001.md#canonical-1223032033013321-1311120122122222-3011310221133212-2310230102322211-1300020320320303-2100203112101221-1132023332200223-2203311203030230): complete subsection reference.

- [default_action_next_policy](resources--service_policy--reference--group-001.md#canonical-0221322033010110-0232101212332032-3023130311222032-1021212223022322-2221300230001100-0332222122113033-1101302100320031-1000300130321122): complete subsection reference.

- [ip_prefix_set](resources--service_policy--reference--group-001.md#canonical-3222211210323201-0233022113132331-2211021001020331-2133220321121022-1112112121320111-2031201022013131-2331022133100133-2133103302331333): complete subsection reference.

- [prefix_list](resources--service_policy--reference--group-001.md#canonical-2310322313323233-0002101030100031-0020303111000323-0221321202210310-0333200101000223-3210210000131302-2312223010320223-3130230113013031): complete subsection reference.

<a id="canonical-1121122332310221-2310102303223311-3320012100231000-0331223310113301-2322122110101313-3301322230010331-3102313102021233-2301303101132310"></a>

<a id="canonical-2032203000220100-1313220201001211-0232330003213130-0100012011001330-1203211020212200-3133032110002001-3332033000012320-0020301022210210"></a>

#### `allow_list.tls_fingerprint_classes` property

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3031230112013110-3213130223312133-0012100300002101-1030303220313100-2301310121131132-2303002131111002-0210232110110202-0311323031121313"></a>

<a id="canonical-0110333212001301-0310132012313311-3303303210023102-2123303223032001-0232220332223321-1202303211313103-2212011301330320-2222330003013310"></a>

#### `allow_list.tls_fingerprint_values` property

Type: `["list", "string"]`. Optional.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3132110210003212-2331132232213331-1112123020302130-1202030332101021-2113013220113122-3103002132013200-2211120020233101-2032322303231033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.asn_list` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-2103300311300332-3332200122100221-1303003012310332-1301221201020110-2111112310210320-2231312202202012-2331211230323103-2101301311212102)
- allow_list.asn_list

<a id="canonical-1001102103011112-1100132211110033-3021312130211230-0232210333130200-0111230330121102-2233223222132110-1011103012032232-2110312013232130"></a>

Type: `"object"`. single nested block, Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223013010121030-2002301331132321-1300030332201332-0032002212300020-0232011101003313-2303100320101033-0212231123033031-1003032023231202"></a>

### Direct properties for `allow_list.asn_list`

<a id="canonical-0002101112202000-0313112133231111-3222121311110103-0123111133312311-3120131102323001-1331331133112300-2300332131333111-1123100132100020"></a>

#### `allow_list.asn_list.as_numbers` property

Type: `["list", "number"]`. Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3322131022003321-1102103331111032-3000011101310303-3022231110302022-0300320130213212-3132013103330201-0010310300002333-1313301232213332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.asn_set` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-2103300311300332-3332200122100221-1303003012310332-1301221201020110-2111112310210320-2231312202202012-2331211230323103-2101301311212102)
- allow_list.asn_set

<a id="canonical-1201113321321311-1230313201033103-1211330102111321-0112201200322310-1301333320033133-1332211233103023-1201210031023220-0121231101333323"></a>

Type: `"object"`. list nested block, Optional.

Addresses that belong to the ASNs in the given bgp\_asn\_set The ASN is obtained by performing a
lookup for the source IPv4 Address in a GeoIP DB.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
asn_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201232200311202-2021211301223320-1313111301300332-0330033202330313-1322220110110210-2323210001021002-0002102133213100-1320223231313233"></a>

### Direct properties for `allow_list.asn_set`

<a id="canonical-2102112202111302-1032112132103102-0010101123102232-3211032011001203-2210121101130213-1211023013332210-1021200100211130-3111213131210313"></a>

#### `allow_list.asn_set.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1021023233321320-2100332013002330-3221103112302231-3130000033210231-3200000302213233-3000211332223011-2223232132301030-2212021022000031"></a>

<a id="canonical-2221210023012231-0123303321303322-3220113211311210-0032220102332002-3323200021132113-2331312010030130-2133031221112330-1230212320102330"></a>

#### `allow_list.asn_set.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1003112013131201-3001330323200330-3211120033132312-2020032110333323-0032203020020220-3001122202121213-3221331032123032-1300200103203122"></a>

<a id="canonical-3102131122023332-3220110312333210-3210210003010000-0300002101301031-2303300200321223-0303011221203211-0022132200130313-0102203013022313"></a>

#### `allow_list.asn_set.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3230002222313121-3230132232023100-2132221313103213-0130303232302310-0332001112012001-2331130201000100-3121200023032130-2333230000221102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.default_action_allow` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-2103300311300332-3332200122100221-1303003012310332-1301221201020110-2111112310210320-2231312202202012-2331211230323103-2101301311212102)
- allow_list.default_action_allow

<a id="canonical-0032220000203011-3033301322320311-0321001130202210-2233212103133231-0213112320132220-3331130130212232-0313122003330000-2323232023030120"></a>

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
default_action_allow = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223032033013321-1311120122122222-3011310221133212-2310230102322211-1300020320320303-2100203112101221-1132023332200223-2203311203030230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.default_action_deny` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-2103300311300332-3332200122100221-1303003012310332-1301221201020110-2111112310210320-2231312202202012-2331211230323103-2101301311212102)
- allow_list.default_action_deny

<a id="canonical-3200033300231231-2031100211210333-0101300302213323-3111323012233203-3112012212331213-3133211333011001-3313022313110031-0130010101301021"></a>

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
default_action_deny = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221322033010110-0232101212332032-3023130311222032-1021212223022322-2221300230001100-0332222122113033-1101302100320031-1000300130321122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.default_action_next_policy` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-2103300311300332-3332200122100221-1303003012310332-1301221201020110-2111112310210320-2231312202202012-2331211230323103-2101301311212102)
- allow_list.default_action_next_policy

<a id="canonical-0330020331031103-2312200032322032-1333110323333120-3203132330003102-0231220110023320-3230333301101013-3301023201101112-0002011302233100"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
default_action_next_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222211210323201-0233022113132331-2211021001020331-2133220321121022-1112112121320111-2031201022013131-2331022133100133-2133103302331333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-2103300311300332-3332200122100221-1303003012310332-1301221201020110-2111112310210320-2231312202202012-2331211230323103-2101301311212102)
- allow_list.ip_prefix_set

<a id="canonical-3021001031301230-0233322022003232-0210312002232123-3221331101002001-0333130231103103-3133132232320031-3200102320322330-0301213000000002"></a>

Type: `"object"`. list nested block, Optional.

Addresses that are covered by the prefixes in the given ip\_prefix\_set.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030220202322101-2330023110022302-1010102110113133-1311020123201303-2123203331201100-2022320112113003-3002102230311113-0012212213223130"></a>

### Direct properties for `allow_list.ip_prefix_set`

<a id="canonical-0110020331203121-1020120002102132-1302120332212212-1220103113022001-1221113020200031-3021003031300300-2120130212302211-3313202213211201"></a>

#### `allow_list.ip_prefix_set.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2331322222123101-1320232311030121-2023201231001220-1211030213000312-0001201232313333-3123112130023010-0011310211011321-1303133112111311"></a>

<a id="canonical-1131322223103323-0323203220002211-3021130101030311-2222111001101221-1023030211023302-3123333122100313-3001300003012310-0023303030002133"></a>

#### `allow_list.ip_prefix_set.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1301110231122301-3131211323111120-2013321100100123-0010121102113013-1003323313023302-2103303010100213-1320202133312003-1110133230221333"></a>

<a id="canonical-1333103232001120-3332220021322120-1230110210303232-1132021000232130-1222323332223011-1023131300100110-0311010322213203-0322223110123220"></a>

#### `allow_list.ip_prefix_set.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2310322313323233-0002101030100031-0020303111000323-0221321202210310-0333200101000223-3210210000131302-2312223010320223-3130230113013031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.prefix_list` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-2103300311300332-3332200122100221-1303003012310332-1301221201020110-2111112310210320-2231312202202012-2331211230323103-2101301311212102)
- allow_list.prefix_list

<a id="canonical-0033011332111321-0030120220011120-3030223131130113-1201101301332223-1012313002013322-3111030302302123-1230333222213232-1131310233301311"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0002313203120001-3313210003322310-1030232310332033-1220133212231202-0022102013113111-3321112030203320-1223002121202200-1333022203003200"></a>

### Direct properties for `allow_list.prefix_list`

<a id="canonical-2202030320030211-3301133123300010-0121131113233311-1302110100200122-0121112010303132-2120122222023013-1322333222311013-1231221012302311"></a>

#### `allow_list.prefix_list.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2312320122102201-0133030001231211-3021113302202221-0231313332122131-1331023031332202-2100201311220333-3033202032033122-2311001313122213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `any_server` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- any_server

<a id="canonical-1010022313220032-0013322123301331-3333333232310222-3131013132311222-0312133020002231-1333231312221231-2021000112022302-2222001122223122"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: any\_server, server\_name, server\_name\_matcher, server\_selector\] Enable this option.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [any_server](resources--service_policy--reference--group-001.md#canonical-1010022313220032-0013322123301331-3333333232310222-3131013132311222-0312133020002231-1333231312221231-2021000112022302-2222001122223122)
- [server_name](resources--service_policy--reference--group-001.md#canonical-0133223010313332-1110211023313330-0010322022330102-0322023303232002-3232002130100231-3032013031322121-1301220213200220-0101110033331123)
- [server_name_matcher](resources--service_policy--reference--group-003.md#canonical-2000320332113331-2121031221100100-1221133200122002-3322012030312113-3103022233103020-1100222220211102-1020331122111223-0203320320001212)
- [server_selector](resources--service_policy--reference--group-003.md#canonical-2100103232022233-0113101000311301-3031331021032201-3032212030030022-3321011201111130-3223030011011331-2010300313013200-1101302322210021)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_server = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221222222001232-0223003221133031-1130311332020010-2103301202310211-1113302100001102-0301330330032012-1001212010020311-3300313001303012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_all_requests` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- deny_all_requests

<a id="canonical-3000321302123202-0112320203222012-3033013210133012-3313121213112032-3300203233300233-1031222223201033-1233002203021101-2001220123301012"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for deny all requests.

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
deny_all_requests = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300232031103221-3132103232200031-0132033333013022-1032110132202122-0000103233010213-1232110230300121-0103203200203111-0001311012323210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- deny_list

<a id="canonical-2130322123310112-0302010020132022-3323113023231331-3210232201121112-2122111023320011-2022222301103232-0130032022112322-1323301103100122"></a>

Type: `"object"`. single nested block, Optional.

List of sources. A request belongs to this list if it satisfies any of the match criteria.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_deny"),
  validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_next_policy"),
  validators.ConflictingObjectAttributes("default_action_deny",
    "default_action_next_policy")}
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
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

Terraform syntax:

```terraform
deny_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302212003010111-1121111121122022-2311110111210110-0130100330332233-3001013000000100-1320122001231223-3021103232203033-3001003101201103"></a>

### Direct properties for `deny_list`

- [asn_list](resources--service_policy--reference--group-001.md#canonical-3200203300011211-1023112333103122-0102113010000111-3120132133113130-0300132113120321-2302010330032011-2122133210202013-3131203003022003): complete subsection reference.

- [asn_set](resources--service_policy--reference--group-001.md#canonical-1312120322122303-2013021003303311-3212020223301220-2131211113331303-3233130311121031-2310032133332302-3030030010203223-1010220011103122): complete subsection reference.

<a id="canonical-2332022333132001-3213110000113013-3032022133023113-2033020133103230-0030111222002221-3021120221120312-3013331100023232-2130131130213013"></a>

<a id="canonical-1233221332330030-3022111323303323-3330021132321310-2313330322102003-1103031231311100-0213203311322031-1102323233301131-1310030203010233"></a>

#### `deny_list.country_list` property

Type: `["list", "string"]`. Optional.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Addresses that belong to one of the countries in the given list The country is obtained by
performing a lookup for the source IPv4 Address in a GeoIP DB. Possible values are
\`COUNTRY\_NONE\`, \`COUNTRY\_AD\`, \`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`,
\`COUNTRY\_AI\`, \`COUNTRY\_AL\`, \`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`,
\`COUNTRY\_AQ\`, \`COUNTRY\_AR\`, \`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`,
\`COUNTRY\_AW\`, \`COUNTRY\_AX\`, \`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`,
\`COUNTRY\_BD\`, \`COUNTRY\_BE\`, \`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`,
\`COUNTRY\_BI\`, \`COUNTRY\_BJ\`, \`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`,
\`COUNTRY\_BO\`, \`COUNTRY\_BQ\`, \`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`,
\`COUNTRY\_BV\`, \`COUNTRY\_BW\`, \`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`,
\`COUNTRY\_CC\`, \`COUNTRY\_CD\`, \`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`,
\`COUNTRY\_CI\`, \`COUNTRY\_CK\`, \`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`,
\`COUNTRY\_CO\`, \`COUNTRY\_CR\`, \`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`,
\`COUNTRY\_CW\`, \`COUNTRY\_CX\`, \`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`,
\`COUNTRY\_DJ\`, \`COUNTRY\_DK\`, \`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`,
\`COUNTRY\_EC\`, \`COUNTRY\_EE\`, \`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`,
\`COUNTRY\_ES\`, \`COUNTRY\_ET\`, \`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`,
\`COUNTRY\_FM\`, \`COUNTRY\_FO\`, \`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`,
\`COUNTRY\_GD\`, \`COUNTRY\_GE\`, \`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`,
\`COUNTRY\_GI\`, \`COUNTRY\_GL\`, \`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`,
\`COUNTRY\_GQ\`, \`COUNTRY\_GR\`, \`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`,
\`COUNTRY\_GW\`, \`COUNTRY\_GY\`, \`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`,
\`COUNTRY\_HR\`, \`COUNTRY\_HT\`, \`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`,
\`COUNTRY\_IL\`, \`COUNTRY\_IM\`, \`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`,
\`COUNTRY\_IR\`, \`COUNTRY\_IS\`, \`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`,
\`COUNTRY\_JO\`, \`COUNTRY\_JP\`, \`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`,
\`COUNTRY\_KI\`, \`COUNTRY\_KM\`, \`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`,
\`COUNTRY\_KW\`, \`COUNTRY\_KY\`, \`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`,
\`COUNTRY\_LC\`, \`COUNTRY\_LI\`, \`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`,
\`COUNTRY\_LT\`, \`COUNTRY\_LU\`, \`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`,
\`COUNTRY\_MC\`, \`COUNTRY\_MD\`, \`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`,
\`COUNTRY\_MH\`, \`COUNTRY\_MK\`, \`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`,
\`COUNTRY\_MO\`, \`COUNTRY\_MP\`, \`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`,
\`COUNTRY\_MT\`, \`COUNTRY\_MU\`, \`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`,
\`COUNTRY\_MY\`, \`COUNTRY\_MZ\`, \`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`,
\`COUNTRY\_NF\`, \`COUNTRY\_NG\`, \`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`,
\`COUNTRY\_NP\`, \`COUNTRY\_NR\`, \`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`,
\`COUNTRY\_PA\`, \`COUNTRY\_PE\`, \`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`,
\`COUNTRY\_PK\`, \`COUNTRY\_PL\`, \`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`,
\`COUNTRY\_PS\`, \`COUNTRY\_PT\`, \`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`,
\`COUNTRY\_RE\`, \`COUNTRY\_RO\`, \`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`,
\`COUNTRY\_SA\`, \`COUNTRY\_SB\`, \`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`,
\`COUNTRY\_SG\`, \`COUNTRY\_SH\`, \`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`,
\`COUNTRY\_SL\`, \`COUNTRY\_SM\`, \`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`,
\`COUNTRY\_SS\`, \`COUNTRY\_ST\`, \`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`,
\`COUNTRY\_SZ\`, \`COUNTRY\_TC\`, \`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`,
\`COUNTRY\_TH\`, \`COUNTRY\_TJ\`, \`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`,
\`COUNTRY\_TN\`, \`COUNTRY\_TO\`, \`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`,
\`COUNTRY\_TW\`, \`COUNTRY\_TZ\`, \`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`,
\`COUNTRY\_US\`, \`COUNTRY\_UY\`, \`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`,
\`COUNTRY\_VE\`, \`COUNTRY\_VG\`, \`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`,
\`COUNTRY\_WF\`, \`COUNTRY\_WS\`, \`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`,
\`COUNTRY\_YT\`, \`COUNTRY\_ZA\`, \`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_action_allow](resources--service_policy--reference--group-001.md#canonical-1200123000111120-1333010310201310-1203333213321020-3130230230330000-3212211323113001-2321002010212123-1202012303132200-3102311000023001): complete subsection reference.

- [default_action_deny](resources--service_policy--reference--group-001.md#canonical-2232323011112102-1211320130232110-2203230310101302-2131212022022002-0101130031233203-2112103012231132-0031013211132120-3023102000120010): complete subsection reference.

- [default_action_next_policy](resources--service_policy--reference--group-001.md#canonical-3223120202222231-1331120003020321-3112311303312223-0230231100211210-1320033102021313-2202130303113110-3001033231220021-3131011023010211): complete subsection reference.

- [ip_prefix_set](resources--service_policy--reference--group-001.md#canonical-2123102323031321-1312023110203003-1210220201001320-2201031300123220-1130110123112013-2331210321231023-2302032202132013-3322322322311133): complete subsection reference.

- [prefix_list](resources--service_policy--reference--group-001.md#canonical-3103001302301023-1313031100301202-2112323103321301-3002032322123032-2332113132010100-1232003013221100-0303102322302200-3213300303103010): complete subsection reference.

<a id="canonical-0112121103222303-1103111201221201-1312131211030123-0332030302021130-3032101230300020-0030211201203120-3020120302131122-2013302222100021"></a>

<a id="canonical-2333300330300302-2211302233032031-0003131123112120-3321033233302332-1323110110301313-0030003133021111-2200211102020232-1331103003030203"></a>

#### `deny_list.tls_fingerprint_classes` property

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0121311230110031-2102031301031021-0320330221201221-3113313301030302-2130313221331023-1101302323021232-3102200232321122-2121120001113300"></a>

<a id="canonical-1121332222131023-0132100102210210-3311323323303102-3103333223301222-2122313111221101-2030100233020001-0031200232111120-2033221312133210"></a>

#### `deny_list.tls_fingerprint_values` property

Type: `["list", "string"]`. Optional.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3200203300011211-1023112333103122-0102113010000111-3120132133113130-0300132113120321-2302010330032011-2122133210202013-3131203003022003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.asn_list` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-3300232031103221-3132103232200031-0132033333013022-1032110132202122-0000103233010213-1232110230300121-0103203200203111-0001311012323210)
- deny_list.asn_list

<a id="canonical-1210020332123113-3200112301330231-1130230301301020-3323312033023322-1323202223100123-1233012222120002-0330131202122023-2010301103103102"></a>

Type: `"object"`. single nested block, Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302031033313020-0012331021202322-3120021302122230-2033101121113100-1123230301302003-3330113310131111-2220312311200001-1323323021212102"></a>

### Direct properties for `deny_list.asn_list`

<a id="canonical-0110022231111223-1022230021001002-1223230010011131-2203320221003130-3003003221203233-2321323022133111-1232212110133231-3301122322012303"></a>

#### `deny_list.asn_list.as_numbers` property

Type: `["list", "number"]`. Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1312120322122303-2013021003303311-3212020223301220-2131211113331303-3233130311121031-2310032133332302-3030030010203223-1010220011103122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.asn_set` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-3300232031103221-3132103232200031-0132033333013022-1032110132202122-0000103233010213-1232110230300121-0103203200203111-0001311012323210)
- deny_list.asn_set

<a id="canonical-1230233302031012-1223103221010302-3310321112222101-1100110210232021-1000223330122113-2010010033131000-1311313103021033-2212213120132022"></a>

Type: `"object"`. list nested block, Optional.

Addresses that belong to the ASNs in the given bgp\_asn\_set The ASN is obtained by performing a
lookup for the source IPv4 Address in a GeoIP DB.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
asn_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102031231232102-3232110210231232-1202133300010303-3203321031323322-2103210122220320-3032203220211300-2001223330001220-0331122011120210"></a>

### Direct properties for `deny_list.asn_set`

<a id="canonical-1312001320211012-2033312002013013-0030231231033012-0312131121220213-1233233110131231-3302100221210303-3321202210333222-2220201312003200"></a>

#### `deny_list.asn_set.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3130303331003313-2313130201120031-3301100122011110-3211133212012213-1102321013021032-3011201213302231-0221113012000021-0313103300100003"></a>

<a id="canonical-0323303322213233-1000012132021210-2001213023013212-1112011133022303-2003121232001202-2203022333003100-0113222011313202-1020011113221230"></a>

#### `deny_list.asn_set.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1121223211021103-2002211003333202-0232312122232232-2022333130112222-0110213303001011-1032313113002210-3023023221011031-2301211130022210"></a>

<a id="canonical-0002131113223312-2212201011303021-0230310200202301-2222130112311103-1101001131211231-2131033203013232-2223010010331103-0010112031322210"></a>

#### `deny_list.asn_set.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1200123000111120-1333010310201310-1203333213321020-3130230230330000-3212211323113001-2321002010212123-1202012303132200-3102311000023001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.default_action_allow` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-3300232031103221-3132103232200031-0132033333013022-1032110132202122-0000103233010213-1232110230300121-0103203200203111-0001311012323210)
- deny_list.default_action_allow

<a id="canonical-1123123120200202-1333122233120111-2300001310021032-1213202023023331-2130323013110133-1221102220111101-2121320002033201-2320112121033202"></a>

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
default_action_allow = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232323011112102-1211320130232110-2203230310101302-2131212022022002-0101130031233203-2112103012231132-0031013211132120-3023102000120010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.default_action_deny` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-3300232031103221-3132103232200031-0132033333013022-1032110132202122-0000103233010213-1232110230300121-0103203200203111-0001311012323210)
- deny_list.default_action_deny

<a id="canonical-2211203123032121-3020033011013203-1000120203221121-1130111231233113-3033132032030112-3202020233011103-3103211233130332-0213212011223310"></a>

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
default_action_deny = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223120202222231-1331120003020321-3112311303312223-0230231100211210-1320033102021313-2202130303113110-3001033231220021-3131011023010211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.default_action_next_policy` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-3300232031103221-3132103232200031-0132033333013022-1032110132202122-0000103233010213-1232110230300121-0103203200203111-0001311012323210)
- deny_list.default_action_next_policy

<a id="canonical-3211123131233101-2320010303003001-0000201003311023-0030202210231301-0203232310003303-2332020002331323-1200331022123120-3222111133131002"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
default_action_next_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123102323031321-1312023110203003-1210220201001320-2201031300123220-1130110123112013-2331210321231023-2302032202132013-3322322322311133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-3300232031103221-3132103232200031-0132033333013022-1032110132202122-0000103233010213-1232110230300121-0103203200203111-0001311012323210)
- deny_list.ip_prefix_set

<a id="canonical-2202032310300203-3330321232120201-1330011323301231-1300213233221321-1221012001013203-1300121023223023-3023130333121003-0311102331232200"></a>

Type: `"object"`. list nested block, Optional.

Addresses that are covered by the prefixes in the given ip\_prefix\_set.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-1330333110202023-0010200000301100-2010212201032013-0102300201301131-2313213330320030-0332010110130021-2020102032312311-3100133112010000"></a>

### Direct properties for `deny_list.ip_prefix_set`

<a id="canonical-2012203201302133-2100031202020203-1001031012112101-0121302022203230-3112230000021311-0001312302221102-0203212232323010-0121333101100111"></a>

#### `deny_list.ip_prefix_set.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2212032131001010-0123210032101030-2121130221213202-0133333302102132-1220100210032113-0333331230230123-2203301200220011-3232100133013000"></a>

<a id="canonical-0212201121221100-0200001232001111-1220232320203020-0320032222113103-2330210312230301-1113313113233133-3010203230301010-2213201111322332"></a>

#### `deny_list.ip_prefix_set.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1232012101310123-1201022013203202-1311313020100012-1201001313331223-2323103111000121-0230232200011222-1301312223310033-0231320110201202"></a>

<a id="canonical-2120222301010032-0302133331111131-3012101123010321-3031313133331220-3333302001320231-1130100020103330-0322221331330213-3110200221333121"></a>

#### `deny_list.ip_prefix_set.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3103001302301023-1313031100301202-2112323103321301-3002032322123032-2332113132010100-1232003013221100-0303102322302200-3213300303103010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.prefix_list` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-3300232031103221-3132103232200031-0132033333013022-1032110132202122-0000103233010213-1232110230300121-0103203200203111-0001311012323210)
- deny_list.prefix_list

<a id="canonical-1012302130003113-2313102211131222-3020202113013121-2303103323103211-0313322330013131-1120223211132310-2013202003002113-2330112201122312"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133303310331303-3110103111320131-1312332133222321-3332020223133002-2112131212210011-2022302333032302-1202120303333332-2230210033003031"></a>

### Direct properties for `deny_list.prefix_list`

<a id="canonical-0310110333331220-2320312123210211-0221300232303101-1212311000312030-3212113013331201-2133200000223023-1211201110121213-1021112102030112"></a>

#### `deny_list.prefix_list.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- rule_list

<a id="canonical-1210231001020020-2003112213000033-0122001321300001-0102112010302033-0220031223101333-1311303212120321-2000002302220220-3220012033313130"></a>

Type: `"object"`. single nested block, Optional.

Ordered service-policy rules for non-geographic predicates and actions. Do not use country\_list for
a geo-only rule here: the platform adds match-all any\_ip and any\_asn selectors on readback, so the
rule can match all traffic. Use deny\_list or allow\_list with country\_list for geographic source
matching.

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
rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113302022232001-2330232021101131-3111312200033203-0301321020110003-2022100220002121-0221300330301102-2303322132112333-3330002303033211"></a>

### Direct properties for `rule_list`

- [rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231): complete subsection reference.

<a id="canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- rule_list.rules

<a id="canonical-0310120133110223-0333200202303103-1223331131100013-3011023100202112-1213123201130331-0010300112020120-0002002231122013-2202212103010103"></a>

Type: `"object"`. list nested block, Optional.

Define the list of rules (with an order) that should be evaluated by this service policy. Rules are
evaluated from top to bottom in the list.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032311020230010-1203010000033210-2211331132000102-3313331301202230-0001321123230232-0303010232030032-1321301013030313-3023103102020223"></a>

### Direct properties for `rule_list.rules`

- [metadata](resources--service_policy--reference--group-001.md#canonical-1123212232333000-0012232111202110-2313113110232310-0211313020100131-3123310032200122-1103003133231212-3231033221232220-3301221013221102): complete subsection reference.

- [spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212): complete subsection reference.

<a id="canonical-1123212232333000-0012232111202110-2313113110232310-0211313020100131-3123310032200122-1103003133231212-3231033221232220-3301221013221102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.metadata` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- rule_list.rules.metadata

<a id="canonical-0230311032002310-2102032112120112-2023200102322132-2310322113133112-3222121130312231-3310132000303013-2133301100022332-3233301022110100"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330321313311303-0230203013233023-2030223133210200-0000330310100021-2102131210020231-3022121310012220-0322321021013313-3102003110210332"></a>

### Direct properties for `rule_list.rules.metadata`

<a id="canonical-2213223033003010-3323321321032322-3121230130222022-1211130101200023-2200110133103211-1211110001021222-1310121211321332-0100322011203121"></a>

#### `rule_list.rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1211233011023032-2321000320202212-3132012111311222-0000220331201001-0330130131133002-2332303010231322-0130230322122032-2230331200313022"></a>

<a id="canonical-2322120221123132-1202320223221300-0220223122100101-1122021023030002-3023221302030031-0200012101310323-1312213021120221-3231131023020321"></a>

#### `rule_list.rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- rule_list.rules.spec

<a id="canonical-3023121133013333-1102330102213132-3332203320113130-2203301321011130-3102223123000331-1231200011033022-2013222033333002-0012300022003303"></a>

Type: `"object"`. single nested block, Optional.

Shape of service\_policy\_rule in the storage backend.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("action",
    "waf_action"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_client",
    "client_name"),
  validators.ConflictingObjectAttributes("any_client",
    "client_name_matcher"),
  validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("client_name",
    "client_name_matcher"),
  validators.ConflictingObjectAttributes("client_name",
    "client_selector"),
  validators.ConflictingObjectAttributes("client_name",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("client_name_matcher",
    "client_selector"),
  validators.ConflictingObjectAttributes("client_name_matcher",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("ja4_tls_fingerprint",
    "tls_fingerprint_matcher")}
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
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_name\",\"client_name_matcher\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-dst_asn_choice": "[]",
  "x-ves-oneof-field-dst_ip_choice": "[]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"ja4_tls_fingerprint\",\"tls_fingerprint_matcher\"]"
}
```

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000131030333211-2302032211012022-0031213013120110-1311313311310012-0113200011202130-2002131232323021-3012013001220231-1233232013221130"></a>

### Direct properties for `rule_list.rules.spec`

<a id="canonical-3121232333100200-1201323101301233-1010300112213213-0320300311002320-1322020000001322-2033020001220330-0001030303100300-2320311233233001"></a>

#### `rule_list.rules.spec.action` property

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW|NEXT\_POLICY\] The rule action determines the disposition of the input request
API. If a policy matches a rule with an ALLOW action, the processing of the request proceeds
forward. If it matches a rule with a DENY action, the processing of the request is terminated and an
appropriate message/code returned to.. Possible values are \`DENY\`, \`ALLOW\`, \`NEXT\_POLICY\`.
Defaults to \`DENY\`.

Additional upstream details:

The rule action determines the disposition of the input request API. If it matches a rule with a
DENY action, the processing of the request is terminated and an appropriate message/code returned to
the originator. If it matches a rule with a NEXT\_POLICY\_SET action, evaluation of the current
policy set terminates and evaluation of the next policy set in the chain begins.

&#8203;- DENY: DENY

Deny the request. &#8203;- ALLOW: ALLOW

Allow the request to proceed. &#8203;- NEXT\_POLICY\_SET: NEXT\_POLICY\_SET

Terminate evaluation of the current policy set and begin evaluating the next policy set in the
chain. Note that the evaluation of any remaining policies in the current policy set is skipped.
&#8203;- NEXT\_POLICY: NEXT\_POLICY

Terminate evaluation of the current policy and begin evaluating the next policy in the policy set.
Note that the evaluation of any remaining rules in the current policy is skipped. &#8203;-
LAST\_POLICY: LAST\_POLICY

Terminate evaluation of the current policy and begin evaluating the last policy in the policy set.
Note that the evaluation of any remaining rules in the current policy is skipped. &#8203;-
GOTO\_POLICY: GOTO\_POLICY

Terminate evaluation of the current policy and begin evaluating a specific policy in the policy set.
The policy is specified using the goto\_policy field in the rule and must be after the current
policy in the policy set.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALLOW","DENY","NEXT_POLICY"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW",
    "NEXT_POLICY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW",
    "NEXT_POLICY"
  ],
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  }
}
```

- [any_asn](resources--service_policy--reference--group-001.md#canonical-3022103210312210-3010031330131113-2130110221003111-2021012020002312-2331332221031223-0112303222232021-0003300100022103-1230223102001233): complete subsection reference.

- [any_client](resources--service_policy--reference--group-001.md#canonical-3330312321323020-3301230110321320-3001233100202200-3112211003000110-2101113021032321-2200202001121232-3023120301012012-2022231320231223): complete subsection reference.

- [any_ip](resources--service_policy--reference--group-001.md#canonical-1122112300110001-3220231100322013-0323213231220023-3020323213003111-3230131122231222-2011022001301111-3231213310201133-1111033333033013): complete subsection reference.

- [api_group_matcher](resources--service_policy--reference--group-001.md#canonical-1200123013332212-3203303010013021-3203002133013210-1003313101132211-0113302233322111-2222033012113113-0111001320232211-3203213221111311): complete subsection reference.

- [arg_matchers](resources--service_policy--reference--group-001.md#canonical-0322010120330120-3111202333230122-2110113300222122-2320031332000032-3313112022233210-3020313211100101-0221022311002011-0012110003033112): complete subsection reference.

- [asn_list](resources--service_policy--reference--group-001.md#canonical-1003002011102032-0102011123003030-1132330003012031-3203021300320213-2030112010013101-0300310022321303-0032221212002220-2230223122023011): complete subsection reference.

- [asn_matcher](resources--service_policy--reference--group-001.md#canonical-0232211031132121-3002321133223230-3303210013302032-0323302103003332-0131313120223121-3031312302333221-3012100033212200-2323100313231113): complete subsection reference.

- [body_matcher](resources--service_policy--reference--group-002.md#canonical-0030321002132020-0123311113120033-3011301123233201-0303332212102210-0332111031220132-2100030113212110-1130111021302232-1020213123301102): complete subsection reference.

- [bot_action](resources--service_policy--reference--group-002.md#canonical-0302200302211013-0331210113301001-2003331301230333-2332131010012033-3111300123223102-2333001133220102-0110113210223131-3201213110101121): complete subsection reference.

<a id="canonical-0122000033030321-3322111233303312-0230321130020300-0333230102333221-1132122333032022-2002011130020311-2001020131300113-3000023320031112"></a>

<a id="canonical-0133013221210012-0332033023021311-0211333302021132-2133313202113002-0123331013033212-1020100100203221-0222003133320200-3313233211121220"></a>

#### `rule_list.rules.spec.client_name` property

Type: `"string"`. Optional.

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [client_name_matcher](resources--service_policy--reference--group-002.md#canonical-1013211000212013-2212232120132201-3311130211322132-3320030201030231-3310233022020230-1023302200222133-1112120002313123-1031200213012010): complete subsection reference.

- [client_selector](resources--service_policy--reference--group-002.md#canonical-0211201021232311-0030301003210011-2310000123031310-0130033112310121-3032233333123100-0131113122312203-3011123320102131-0203220020213233): complete subsection reference.

- [cookie_matchers](resources--service_policy--reference--group-002.md#canonical-3121212120020200-2301203330120221-0103311213020123-2031130122012111-2001303101221200-3022101002012312-1221331031322310-2110303302010321): complete subsection reference.

- [domain_matcher](resources--service_policy--reference--group-002.md#canonical-0002120030013310-3002023311102221-2221120010233003-0112023021102232-2110300301110330-3222223323210112-1002023003211130-1212030032102211): complete subsection reference.

<a id="canonical-1013211313113212-1100103301202121-1321121100210222-0221132322010130-1231210310120030-2013222130323111-1233231221102101-3332202313023121"></a>

<a id="canonical-3322102101302220-3031030300232210-2033331033211103-1112030320231121-0203030003210210-2222123223313211-0203223330211001-1023120010313120"></a>

#### `rule_list.rules.spec.expiration_timestamp` property

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [headers](resources--service_policy--reference--group-002.md#canonical-3303211112230231-3131220013211212-3131022320033101-1130301232233213-0301112303322203-0023113002110131-0203122003122231-0230111120022301): complete subsection reference.

- [http_method](resources--service_policy--reference--group-002.md#canonical-3200303221303321-0332123202022111-0032113310031110-3333200123232130-3321331111121012-1030202033031002-1030130232323200-1112120101333313): complete subsection reference.

- [ip_matcher](resources--service_policy--reference--group-002.md#canonical-1030111002320303-0232312311303320-3110022123122122-1021111312200023-0113312211223231-0101120320103223-1232331122313210-1312313222312210): complete subsection reference.

- [ip_prefix_list](resources--service_policy--reference--group-002.md#canonical-1320110103211332-0010132210330323-2221011311200322-1022311233301132-3210220223000122-3203222200021011-1130101332123213-1211103331231223): complete subsection reference.

- [ip_threat_category_list](resources--service_policy--reference--group-002.md#canonical-3333211202130202-0133211023331201-3301313020033000-0233020133130232-2320303320330222-2010201313212003-1110123210010000-1101102021012033): complete subsection reference.

- [ja4_tls_fingerprint](resources--service_policy--reference--group-002.md#canonical-3030030233300323-2201301300112210-1201003030313223-2320311230330003-2013232332000022-0131000212013130-1121021010010202-1011100231001020): complete subsection reference.

- [jwt_claims](resources--service_policy--reference--group-002.md#canonical-0232333201113100-3110011212311231-2210020330303010-1321112033033300-2312211113223103-1101003302101330-3313322102011022-2013130320332331): complete subsection reference.

- [label_matcher](resources--service_policy--reference--group-002.md#canonical-3232100012122002-1210310013101032-0310322320030203-3203213320013100-3120100320020220-0100032322102131-1013203100213003-2110320203300311): complete subsection reference.

<a id="canonical-1332010122103030-1301232132210311-3123233112322122-3202013022321331-1310200212220303-1332112233002301-0013230100221323-1213122231131320"></a>

<a id="canonical-3133310031330300-0003130331021322-1101221130021021-2032210203331010-2213002330012023-1120313110033123-0120132312001321-3020022330210331"></a>

#### `rule_list.rules.spec.log_rule_evaluation` property

Type: `"bool"`. Optional.

Log the rule match details along with the request and continue to evaluate rules in the sequence.

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

- [mum_action](resources--service_policy--reference--group-002.md#canonical-1031203103211101-1320130203110223-1203222000022323-0200122233332223-3330121000123133-2233010210331321-1120111112011131-2031122333013131): complete subsection reference.

- [path](resources--service_policy--reference--group-002.md#canonical-3100030012203210-1100201310311210-1001120230313212-2000131113131300-3023102000302131-3003123100333233-0213221032313101-0303033021301201): complete subsection reference.

- [port_matcher](resources--service_policy--reference--group-002.md#canonical-3032102120213313-0312100211232033-3112223313112201-0031231231213331-3302120020022021-0000132323100020-3012123121223223-2302111000133001): complete subsection reference.

- [query_params](resources--service_policy--reference--group-002.md#canonical-2020223332102111-0033130221100121-2321221112100331-3032233101002012-0013223201032330-1010113011301300-2103120333221013-1121231031201231): complete subsection reference.

- [request_constraints](resources--service_policy--reference--group-002.md#canonical-1111011013002100-0111303211122113-1013330011103010-3133132133031332-1000100313033123-0012220113001030-3102222331133010-0103111313230222): complete subsection reference.

- [segment_policy](resources--service_policy--reference--group-002.md#canonical-0133001032102232-2111212101023102-3020011012032031-1020332111321303-0132322221302200-0121210202031101-1133213100300333-0033133312321223): complete subsection reference.

- [tls_fingerprint_matcher](resources--service_policy--reference--group-002.md#canonical-1020320201300000-0002323231201110-2333023122032003-3301321123001032-2102310103222032-0133000133300130-0123200131201022-1013312312033202): complete subsection reference.

- [user_identity_matcher](resources--service_policy--reference--group-002.md#canonical-3001003323033212-0123331232232220-3131203230202012-2011331002320011-1212313031310311-0101011121113132-2102222322031110-3333021300010302): complete subsection reference.

- [waf_action](resources--service_policy--reference--group-003.md#canonical-1330120330031000-3202113320223100-2303020333222232-1113322320113320-2131320031302300-3021303213200102-0321303213312133-3121220001012010): complete subsection reference.

<a id="canonical-3022103210312210-3010031330131113-2130110221003111-2021012020002312-2331332221031223-0112303222232021-0003300100022103-1230223102001233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.any_asn` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.any_asn

<a id="canonical-3201031031122211-3121221200332103-2000210302033231-3110022213222232-2101203221223023-0111310320201312-3002013310322000-3303311333202223"></a>

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
any_asn = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330312321323020-3301230110321320-3001233100202200-3112211003000110-2101113021032321-2200202001121232-3023120301012012-2022231320231223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.any_client` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.any_client

<a id="canonical-2010321331330322-0300301220123213-0031320001012313-2030112202213333-0101323201222300-1110211110213132-0100222101331031-3111330020303313"></a>

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
any_client = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122112300110001-3220231100322013-0323213231220023-3020323213003111-3230131122231222-2011022001301111-3231213310201133-1111033333033013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.any_ip` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.any_ip

<a id="canonical-0133202313123122-2122033223100213-0013313301201220-3032312113213010-0011233131303002-1311301023210120-1101123102121232-0021202332102302"></a>

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
any_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200123013332212-3203303010013021-3203002133013210-1003313101132211-0113302233322111-2222033012113113-0111001320232211-3203213221111311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.api_group_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.api_group_matcher

<a id="canonical-3300111010320122-1112033221323132-0233011010201112-0212321302330022-3130203013001023-1003202121102131-0113123230103130-1002020012231211"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies a list of values for matching an input string. The match is considered
successful if the input value is present in the list. The result of the match is inverted if
invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("match")}
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
api_group_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313320223223101-0332032322320322-1222022001012033-0300023002332300-1103032032300131-2021132131233312-0002013002123223-3333123102303121"></a>

### Direct properties for `rule_list.rules.spec.api_group_matcher`

<a id="canonical-0111022111211203-2103332032331222-2032112332113321-1131202300231030-3333020121131030-1302233101203131-2113331312300032-1120123122123110"></a>

#### `rule_list.rules.spec.api_group_matcher.invert_matcher` property

Type: `"bool"`. Optional.

Invert String Matcher. Invert the match result.

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

<a id="canonical-2203223001013233-0000303110202300-0220203101321330-0123130133100223-2020331310020303-3103000103120202-2210232223113212-1112031132001331"></a>

<a id="canonical-3220102223032011-2320200320330320-3130102031333122-1231122020210122-0103230122203102-0212331322220100-1201212323311302-1313023022233001"></a>

#### `rule_list.rules.spec.api_group_matcher.match` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "63",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "63",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0322010120330120-3111202333230122-2110113300222122-2320031332000032-3313112022233210-3020313211100101-0221022311002011-0012110003033112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.arg_matchers` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.arg_matchers

<a id="canonical-1130310212222121-2230312101133330-2010303330030301-2303212000030232-1102220003120120-0303320120320000-0113202112232313-1032230133231032"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
arg_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232301122020212-2000200110203213-1113311203101021-3102021033222121-2010222202220132-3321102313131203-0123010122103203-0201001211103131"></a>

### Direct properties for `rule_list.rules.spec.arg_matchers`

- [check_not_present](resources--service_policy--reference--group-001.md#canonical-3213333222333011-1112222202300103-2110301320131131-2220032011303223-3232323330313310-3311322322000200-2020200323231121-0303232322032001): complete subsection reference.

- [check_present](resources--service_policy--reference--group-001.md#canonical-0230123203131202-2311032232000230-1022233213323302-1030233220301323-3030030202311032-3230100123200000-0332312100031121-0100300321003011): complete subsection reference.

<a id="canonical-2011323231013222-3302132301311201-3030101001312031-2102031103213312-3012023233010330-0300013230022001-2330002021022211-1322122331000223"></a>

<a id="canonical-1011212211121333-2023313102332310-1022132000122231-1033133313011232-1023033233020012-2320001223203023-0231111130310321-2231233310210123"></a>

#### `rule_list.rules.spec.arg_matchers.invert_matcher` property

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

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

- [item](resources--service_policy--reference--group-001.md#canonical-1110000231321010-3223302220332100-3202212032212223-1021132301333001-0303300033313333-3032033012323101-2112021102211320-0232002211200321): complete subsection reference.

<a id="canonical-3011031033120301-0120122120013011-0010203333320022-0312120130200120-3311312321130203-1213112110303221-2302012303302220-2021013003000010"></a>

<a id="canonical-0320122202203101-2200301321012220-3000111201200132-1113221001312010-3030010122113310-1213031232031022-0233212001131212-1133223202300200"></a>

#### `rule_list.rules.spec.arg_matchers.name` property

Type: `"string"`. Optional.

A case-sensitive JSON path in the HTTP request body.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3213333222333011-1112222202300103-2110301320131131-2220032011303223-3232323330313310-3311322322000200-2020200323231121-0303232322032001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.arg_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.arg_matchers](resources--service_policy--reference--group-001.md#canonical-0322010120330120-3111202333230122-2110113300222122-2320031332000032-3313112022233210-3020313211100101-0221022311002011-0012110003033112)
- rule_list.rules.spec.arg_matchers.check_not_present

<a id="canonical-2120301202100202-0113031220133011-3031313210132032-3333230132130112-3203121101202232-3221211201301130-0310000211020300-0011002013023303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230123203131202-2311032232000230-1022233213323302-1030233220301323-3030030202311032-3230100123200000-0332312100031121-0100300321003011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.arg_matchers.check_present` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.arg_matchers](resources--service_policy--reference--group-001.md#canonical-0322010120330120-3111202333230122-2110113300222122-2320031332000032-3313112022233210-3020313211100101-0221022311002011-0012110003033112)
- rule_list.rules.spec.arg_matchers.check_present

<a id="canonical-3032112111212202-0202230110023133-3123012032202203-3222100210203113-2300013222131100-1030202023113101-3103302013031301-2230110123112203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110000231321010-3223302220332100-3202212032212223-1021132301333001-0303300033313333-3032033012323101-2112021102211320-0232002211200321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.arg_matchers.item` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- [rule_list.rules.spec.arg_matchers](resources--service_policy--reference--group-001.md#canonical-0322010120330120-3111202333230122-2110113300222122-2320031332000032-3313112022233210-3020313211100101-0221022311002011-0012110003033112)
- rule_list.rules.spec.arg_matchers.item

<a id="canonical-3003233002221113-3333112023231203-3313031330313310-1130113000321230-3020233111103301-3110131113113201-2230130130201011-2333313232121013"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330303301033011-2312222031231312-0222320330203102-3001230130322330-1221233220323003-0302220011200100-3310010113211020-3211130121331300"></a>

### Direct properties for `rule_list.rules.spec.arg_matchers.item`

<a id="canonical-3013133302132001-2310100002302003-0132033023011020-2113112230333113-1302030011201332-2211000112213001-0001303000133101-1110212212302233"></a>

#### `rule_list.rules.spec.arg_matchers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1213031220112002-1012312312333333-3110031330121223-3121111221222211-3203213120210331-0111121313123231-0020133122231213-1101013222320332"></a>

<a id="canonical-1000210031031233-2323003021101121-3101000223113012-1323111332203033-0301312133121301-1132230321221133-2222032333013302-0022231021023313"></a>

#### `rule_list.rules.spec.arg_matchers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0003331321220202-2322221310000223-0120010023311213-1211031103101312-1121211110111033-1122113012320100-3121323110301132-3231221301303021"></a>

<a id="canonical-0323133011030012-2321001032233012-3201331212331221-2110122013233222-3102130021233023-3030223012310120-2123331110032232-2133113110203100"></a>

#### `rule_list.rules.spec.arg_matchers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1003002011102032-0102011123003030-1132330003012031-3203021300320213-2030112010013101-0300310022321303-0032221212002220-2230223122023011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.asn_list` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.asn_list

<a id="canonical-0112220131300210-0000232113301130-1223211111311303-3213010331012012-1030320012010301-3101333000130132-0012230002312010-3111131320010201"></a>

Type: `"object"`. single nested block, Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003232200032332-3001210020310102-0123112303001021-0111232222120333-3330032002333003-2222301030210103-3220033120011120-3322000022033300"></a>

### Direct properties for `rule_list.rules.spec.asn_list`

<a id="canonical-2131233223131302-3213311010331210-3320021200230231-3123331121200213-1002220212332003-3130331023113002-0012013322031223-1331101133230130"></a>

#### `rule_list.rules.spec.asn_list.as_numbers` property

Type: `["list", "number"]`. Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0232211031132121-3002321133223230-3303210013302032-0323302103003332-0131313120223121-3031312302333221-3012100033212200-2323100313231113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.asn_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-2003033211210223-1233232013203033-1213012320103003-1220003210002322-3220300031211001-2003332333132300-0301022032112012-0132021330223231)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-2200022213120332-0220211101301113-3032022313032233-3330130333233030-0010110110311113-1112333223123300-0232313012010010-0302001123322212)
- rule_list.rules.spec.asn_matcher

<a id="canonical-0313232001300103-1123001113133303-0031021200132120-0120213223020223-3212211131200330-3212030313231201-3331332203302112-2212013330103033"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003023002112022-0010323121323013-0130102200313213-3233210002120021-0112011013011330-2321120122212210-3320301232000200-1322232023311220"></a>

### Direct properties for `rule_list.rules.spec.asn_matcher`

- [asn_sets](resources--service_policy--reference--group-002.md#canonical-3013111121202120-1033033012323330-1020021320132301-3202002003210211-0333123103122320-1100213133023033-2021101000133111-2002101210222223): complete subsection reference.
