---
page_title: "xcsh_dns_lb_pool reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_pool reference."
---

# xcsh_dns_lb_pool reference

<a id="canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- Property reference

<a id="canonical-1220131332311202-2333223000030332-0022321330312023-3223302100212231-1310210213110130-0101302332312120-0232300113201010-1201222320210112"></a>

### Direct properties for `xcsh_dns_lb_pool`

- [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-0123311331023030-2123210223210232-2020121111022213-0033300301133130-3013123123122303-2220212232330101-0132302013203032-0000203310030023): complete subsection reference.

- [aaaa_pool](resources--dns_lb_pool--reference--group-001.md#canonical-2212212022130300-3213012311313203-2210121222221333-2013311220021222-0022110313202023-1323111030012311-3031202030303213-0113222132323033): complete subsection reference.

<a id="canonical-1210111123133210-2211111210000320-2310032103021021-2233320033213222-0113333012102313-3032123310001300-0011300300103331-2130000111023332"></a>

<a id="canonical-1001330112212003-2220133021330023-3123133313323000-3323203020010100-0031012012220103-2101133320012302-3323330321300111-1110112222231012"></a>

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

- [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-2102130031203200-0130012310030313-1203320121212201-1330323012031100-2012202003023332-0303332132013312-0103012331231303-1323123021230220): complete subsection reference.

<a id="canonical-1221121131313112-3022221102131321-0221210232003123-0210000332002030-2111232312013123-2000123313001312-1121001103023320-1022010213111110"></a>

<a id="canonical-2303303131032020-1330231300311231-3201303233101113-0333303200132213-1311232332023221-0203212310113120-1312022220032331-2211030012130001"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2122311033233010-2212233120100103-3321302021101002-1233133131333233-0122330230201031-3230030001220230-1033301323312011-0033102131322301"></a>

<a id="canonical-0220211122113111-2133100230033030-2212113103013003-2122203002023002-1331213322012323-1200231222312213-2030021221230121-3203001031303021"></a>

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

<a id="canonical-3331130300100213-0310013203301101-1031010132232103-2231001311230110-0220010222221123-1230320312002201-2230001213021233-0110011333001321"></a>

<a id="canonical-2033221201322110-3121132122033132-0220231310121332-1323310302231331-0301101210001232-3010011033122003-0220330000003021-2013203132221322"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2101323200313021-2303211333303212-0232233212200301-1301123232300221-3111002022213030-3023232121230022-2001213131100123-2322012320310212"></a>

<a id="canonical-2013300132023212-1333033011202112-1202032320120002-2101313122010321-1132220233221200-1011033203012030-1323022000100330-0103303021022332"></a>

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

<a id="canonical-0121001122132133-3122322130221331-3113203023302232-2000313102211122-3211130313301110-1111320212313313-1021122221002101-3203233033313020"></a>

<a id="canonical-1002210323313201-0023220102122221-2020023111131310-1023030032321312-2233301101133130-1303311202132230-2313320012210223-2302021220221332"></a>

#### `load_balancing_mode` property

Type: `"string"`. Optional, Computed.

\[Enum: ROUND\_ROBIN|RATIO\_MEMBER|STATIC\_PERSIST|PRIORITY\] - ROUND\_ROBIN: Round-Robin Round
Robin will ensure random equal distribution of requests among all pool members in a pool. -
RATIO\_MEMBER: Ratio-Member Ratio-Member performs load balancing of requests across the pool members
based on the ratio assigned to each pool member - STATIC\_PERSIST.. Possible values are
\`ROUND\_ROBIN\`, \`RATIO\_MEMBER\`, \`STATIC\_PERSIST\`, \`PRIORITY\`. Defaults to
\`ROUND\_ROBIN\`.

Additional upstream details:

&#8203;- ROUND\_ROBIN: Round-Robin

Round Robin will ensure random equal distribution of requests among all pool members in a pool.
&#8203;- RATIO\_MEMBER: Ratio-Member

Ratio-Member performs load balancing of requests across the pool members based on the ratio assigned
to each pool member &#8203;- STATIC\_PERSIST: Static-Persist

The Static Persist load balancing method uses the persist mask, with the source IP address of the
Local Domain Name Server (LDNS), in a deterministic algorithm to send requests to a specific pool
member. If the DNS resolver passes ECS (EDNS-Client-Subnet) information, then a hash of it will be
used, to send the client to the same pool member &#8203;- PRIORITY: Priority

The Priority load balancing method returns all available endpoints in a pool with the highest
priority. Pool Members have a priority value, starting from zero, where a lower value means a higher
priority.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["PRIORITY","RATIO_MEMBER","ROUND_ROBIN","STATIC_PERSIST"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ROUND_ROBIN",
    "RATIO_MEMBER",
    "STATIC_PERSIST",
    "PRIORITY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ROUND_ROBIN",
  "enum": [
    "ROUND_ROBIN",
    "RATIO_MEMBER",
    "STATIC_PERSIST",
    "PRIORITY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [mx_pool](resources--dns_lb_pool--reference--group-001.md#canonical-2212032032130133-2301313011222230-0123131011123021-0200203312302320-1232223301102320-1012121332111002-1123122031330210-2100311112312320): complete subsection reference.

<a id="canonical-1313120011002132-2230032001101322-0331001313223222-3232301211102123-0321302031200011-2023210232010233-2131322003210111-2133201323022202"></a>

<a id="canonical-1112000210120313-0320000230123131-1210033133031321-1101310101212112-1033120101012011-3120202001222311-1230330131130200-0222132222203101"></a>

#### `name` property

Type: `"string"`. Required.

Name of the DNS LB Pool. Must be unique within the namespace.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0220130200032203-2121130032332311-2211313103301010-1023222103121013-3311203221012200-3220330222200003-2000211133131013-1303201131032202"></a>

<a id="canonical-2131220202012132-2233132312313201-0232233323213022-2202103120000332-1323313001120320-2222022222300032-0321221221321222-0000223110223200"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace for the DNS LB Pool. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
EnumExtractionComplete: false
EnumValidators: [{"version":1,"validator":"OneOf","values":["system"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [srv_pool](resources--dns_lb_pool--reference--group-001.md#canonical-3000311322032000-2023201012210313-3122222330203102-3333210122112323-1033111323322312-3122231232003322-3300111021330330-2220331310311332): complete subsection reference.

- [timeouts](resources--dns_lb_pool--reference--group-001.md#canonical-0301322011101121-0101320022303020-2233133301222122-3123332332021333-3021322031131312-1333202003120133-0311330103332133-1222331311003000): complete subsection reference.

<a id="canonical-1331300301120230-3101310120203322-3322113323011220-0021331301102333-0202320103220032-1212112213130313-3100112333023102-0323231212232212"></a>

<a id="canonical-3033100101301220-3020303331020001-1021010330111021-3333033003111001-3210131320322122-1023330012132331-3332312101121123-0310013110301121"></a>

#### `ttl` property

Type: `"number"`. Optional, Computed.

\[OneOf: TTL, use\_rrset\_ttl\] Exclusive with \[use\_rrset\_ttl\] Custom TTL in seconds (default
&#8203;30) for responses from this pool.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

OneOf alternatives in this subsection:

- [TTL](resources--dns_lb_pool--reference--group-001.md#canonical-1331300301120230-3101310120203322-3322113323011220-0021331301102333-0202320103220032-1212112213130313-3100112333023102-0323231212232212)
- [use_rrset_ttl](resources--dns_lb_pool--reference--group-001.md#canonical-1110333213023111-2321322223321202-2311131320220203-2111101222113310-0212111202332331-0302323323222331-0331013011332111-1133033111203113)

Select alternatives according to the provider validators above.

- [use_rrset_ttl](resources--dns_lb_pool--reference--group-001.md#canonical-1201132013313011-0032030003133200-1120231001123032-1302100130003232-3310032232121310-1200122301101332-3100010223322121-0201133221211331): complete subsection reference.

<a id="canonical-2233221013320020-1101222300012001-1220213302121132-2030211001210102-3020313110102021-2021031101131010-3023302202321112-0131122222220033"></a>

### All schema paths for `xcsh_dns_lb_pool`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `a_pool` | [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-3233100323022233-0123010032011212-0122212003020230-3111022302320302-3010220211312232-1323231022212203-0103011300303323-1232130130031001) |
| `a_pool.disable_health_check` | [a_pool.disable_health_check](resources--dns_lb_pool--reference--group-001.md#canonical-2211013022000201-1330111311320313-3101220030201311-3321311331101101-1032102211230330-2220222223002312-1302101011313203-3200133313333122) |
| `a_pool.health_check` | [a_pool.health_check](resources--dns_lb_pool--reference--group-001.md#canonical-1333122310212023-3223300103002330-1111023120031101-1120230113320232-3031020220331110-3332212313111101-2120121012120300-1011022322301232) |
| `a_pool.health_check.name` | [a_pool.health_check.name](resources--dns_lb_pool--reference--group-001.md#canonical-0203000013010030-2000301233222002-0321303002322332-3222021303311112-0100331303312330-0321103123200313-3323300201033102-3201101123223212) |
| `a_pool.health_check.namespace` | [a_pool.health_check.namespace](resources--dns_lb_pool--reference--group-001.md#canonical-0021003132211223-0100102331221113-2312103330031000-3010030023031320-0120013231233322-2000311202001103-3331020323003033-0212130232030130) |
| `a_pool.health_check.tenant` | [a_pool.health_check.tenant](resources--dns_lb_pool--reference--group-001.md#canonical-2010132021130112-0310130002230100-0132100222130200-2233303120032111-1011220201133123-1311233233233301-3203022112312212-3102203323311333) |
| `a_pool.max_answers` | [a_pool.max_answers](resources--dns_lb_pool--reference--group-001.md#canonical-0122332113130303-3020231301000133-0012111202121023-1321220020123110-2322010131331313-1331100300130212-0230330220310301-3110130133132001) |
| `a_pool.members` | [a_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-1110023220310033-0010311113030201-3232222202032300-1300201332310113-2202232031203213-2003112111121230-3131003033023101-2330321313303101) |
| `a_pool.members.disable_spec` | [a_pool.members.disable_spec](resources--dns_lb_pool--reference--group-001.md#canonical-3103002321130201-3223330212303233-0033110301012031-2100202013201022-0200122332233123-3121323231000202-2133222111012200-0333001000130123) |
| `a_pool.members.ip_endpoint` | [a_pool.members.ip_endpoint](resources--dns_lb_pool--reference--group-001.md#canonical-0303200202321110-3201110021132001-1210131000020232-2211212000203311-0120212220232120-2010310002013110-1311302221301232-1012002112213230) |
| `a_pool.members.name` | [a_pool.members.name](resources--dns_lb_pool--reference--group-001.md#canonical-1310103123320232-2332321333202021-1222211101122123-0020023323113101-2231221013112221-1120033212331030-0231020023303320-0311130231002002) |
| `a_pool.members.priority` | [a_pool.members.priority](resources--dns_lb_pool--reference--group-001.md#canonical-3330322033213101-1103311303103122-0212300312211200-1333302323233110-0321031132003020-0331333010120300-0120223210200201-1003333222233011) |
| `a_pool.members.ratio` | [a_pool.members.ratio](resources--dns_lb_pool--reference--group-001.md#canonical-1110033030021233-3303110123212113-2203300010011211-1132200110120313-1023200132332102-2120310202011032-1030013202223122-2011010220020122) |
| `aaaa_pool` | [aaaa_pool](resources--dns_lb_pool--reference--group-001.md#canonical-0303103131133323-1020303203213322-3033133312011032-2010312120323223-3002021320230221-2023013321032330-0203203210103331-1323121111113302) |
| `aaaa_pool.max_answers` | [aaaa_pool.max_answers](resources--dns_lb_pool--reference--group-001.md#canonical-3022111311322010-2310103001200022-1322220222212133-1223332231313101-3112013112021101-0311321030332233-2031032320320220-1123230113302332) |
| `aaaa_pool.members` | [aaaa_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-0223322323321233-3010122103023230-3030003000333310-0321211322010013-1303111220311330-0203202332300303-3323323331233203-0203023011003000) |
| `aaaa_pool.members.disable_spec` | [aaaa_pool.members.disable_spec](resources--dns_lb_pool--reference--group-001.md#canonical-0223013111000213-2313010220020232-0023112223300220-1201020003230230-3333201302012300-1023030032100230-2231301021011303-3323023312023000) |
| `aaaa_pool.members.ip_endpoint` | [aaaa_pool.members.ip_endpoint](resources--dns_lb_pool--reference--group-001.md#canonical-1231232013023321-3033331310020033-0102023010310221-1122001023201210-3233310312032032-2303301103211321-3010333230333021-2233202222131200) |
| `aaaa_pool.members.name` | [aaaa_pool.members.name](resources--dns_lb_pool--reference--group-001.md#canonical-0111001132310023-2002323032233203-1121221231300112-2002330333030121-0101121210000300-2031311012113123-2012213110313232-2203002122310222) |
| `aaaa_pool.members.priority` | [aaaa_pool.members.priority](resources--dns_lb_pool--reference--group-001.md#canonical-0000333232313003-3110112331102130-1120121223230112-2222022033110021-0220203130120222-0121233110002030-0012031212021122-2223223132013211) |
| `aaaa_pool.members.ratio` | [aaaa_pool.members.ratio](resources--dns_lb_pool--reference--group-001.md#canonical-1000131231202103-2112133311300102-0020233312312312-3023211210111233-1000303210132130-2311011200000020-3031312211110132-3131030010212013) |
| `annotations` | [annotations](resources--dns_lb_pool--reference--group-001.md#canonical-1210111123133210-2211111210000320-2310032103021021-2233320033213222-0113333012102313-3032123310001300-0011300300103331-2130000111023332) |
| `cname_pool` | [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-1102013100110222-1003231012010002-1133112100231111-0303132130313220-2222201031010130-3313111310133311-2132322201221201-3130012321101032) |
| `cname_pool.disable_health_check` | [cname_pool.disable_health_check](resources--dns_lb_pool--reference--group-001.md#canonical-2011102101303300-0212233232002101-2200221301121100-0312103323012221-1223210113221131-0330301132111131-3323303102000212-0231312020223231) |
| `cname_pool.health_check` | [cname_pool.health_check](resources--dns_lb_pool--reference--group-001.md#canonical-1321011122120221-3203100211011303-0301300133112220-2023303220123202-1020121020113011-1210211332130212-0112313312103203-1011303223021021) |
| `cname_pool.health_check.name` | [cname_pool.health_check.name](resources--dns_lb_pool--reference--group-001.md#canonical-0220120022010302-1220223310212131-0330321212101221-1222230300101103-3032123303100210-2130031110220020-0331112333033302-3223231012222030) |
| `cname_pool.health_check.namespace` | [cname_pool.health_check.namespace](resources--dns_lb_pool--reference--group-001.md#canonical-1301220031313131-2133133333131132-0320130333101301-1231300212213023-0223221023111100-1232110010013202-0021102023113133-0102300032001010) |
| `cname_pool.health_check.tenant` | [cname_pool.health_check.tenant](resources--dns_lb_pool--reference--group-001.md#canonical-3110010200312110-1221031232210002-0203331233130303-3233013033203010-3020003000313111-3220202020311201-0032203123321112-3212133232131121) |
| `cname_pool.members` | [cname_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-3130032103000222-1011103203200111-3322133132311001-2030132223132010-1131322222111111-1310110331033301-2021212003310110-3200131002220011) |
| `cname_pool.members.domain` | [cname_pool.members.domain](resources--dns_lb_pool--reference--group-001.md#canonical-0023321322111013-3032000321303333-2122223001122233-3300000200213222-3033211322300132-3013122030020333-1120131201023122-2301033120302011) |
| `cname_pool.members.final_translation` | [cname_pool.members.final_translation](resources--dns_lb_pool--reference--group-001.md#canonical-0211212203300330-3110212133300101-2110112012032013-3310212123201121-0012211213111321-2233122023000213-1303021013210003-0213303303233202) |
| `cname_pool.members.name` | [cname_pool.members.name](resources--dns_lb_pool--reference--group-001.md#canonical-3331123003212032-0320121311010220-2132303033103011-3131200330223310-3132013003332211-0110310222130320-0011020132310231-3020302113033320) |
| `cname_pool.members.priority` | [cname_pool.members.priority](resources--dns_lb_pool--reference--group-001.md#canonical-0301010310231333-3213332201112122-1131303320233202-3302033103310230-1013233330303013-2311322220312321-3213211002032203-0013320232212330) |
| `cname_pool.members.ratio` | [cname_pool.members.ratio](resources--dns_lb_pool--reference--group-001.md#canonical-2023002322313010-1222021210203031-1201332213302001-2322002001211100-2012113022203203-3332333102333222-0023022133032332-1023021320032210) |
| `description` | [description](resources--dns_lb_pool--reference--group-001.md#canonical-1221121131313112-3022221102131321-0221210232003123-0210000332002030-2111232312013123-2000123313001312-1121001103023320-1022010213111110) |
| `disable` | [disable](resources--dns_lb_pool--reference--group-001.md#canonical-2122311033233010-2212233120100103-3321302021101002-1233133131333233-0122330230201031-3230030001220230-1033301323312011-0033102131322301) |
| `id` | [ID](resources--dns_lb_pool--reference--group-001.md#canonical-3331130300100213-0310013203301101-1031010132232103-2231001311230110-0220010222221123-1230320312002201-2230001213021233-0110011333001321) |
| `labels` | [labels](resources--dns_lb_pool--reference--group-001.md#canonical-2101323200313021-2303211333303212-0232233212200301-1301123232300221-3111002022213030-3023232121230022-2001213131100123-2322012320310212) |
| `load_balancing_mode` | [load_balancing_mode](resources--dns_lb_pool--reference--group-001.md#canonical-0121001122132133-3122322130221331-3113203023302232-2000313102211122-3211130313301110-1111320212313313-1021122221002101-3203233033313020) |
| `mx_pool` | [mx_pool](resources--dns_lb_pool--reference--group-001.md#canonical-2100123321012023-3313133333301013-2131313230200003-3313332331322203-2310021100232201-2003031322230202-1320301323312233-1320211103320332) |
| `mx_pool.max_answers` | [mx_pool.max_answers](resources--dns_lb_pool--reference--group-001.md#canonical-0310110113000210-2323022131111101-1130330321132200-1313132110210331-2233030131320012-1223032311230221-0311032133031130-2131023122331122) |
| `mx_pool.members` | [mx_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-2021002222021132-0302010122212200-0230103011021022-1010310312021130-0322330103320013-2332223123232122-3033213121321232-1321233020331201) |
| `mx_pool.members.domain` | [mx_pool.members.domain](resources--dns_lb_pool--reference--group-001.md#canonical-2002230003312332-0321130231133231-3201231303232120-1232313202311131-0213300010022031-0101012110213200-0220100131123203-0123320331010001) |
| `mx_pool.members.name` | [mx_pool.members.name](resources--dns_lb_pool--reference--group-001.md#canonical-0310212033330210-1100102230333231-2021120123220100-0320210122122320-2210231123212002-1102103303111123-1101002311113033-2300211323121133) |
| `mx_pool.members.priority` | [mx_pool.members.priority](resources--dns_lb_pool--reference--group-001.md#canonical-2031330212201130-2121100230210110-3323112000320303-2100232111230122-1013101313100212-3111300310321102-3313012112233011-0221100323122002) |
| `mx_pool.members.ratio` | [mx_pool.members.ratio](resources--dns_lb_pool--reference--group-001.md#canonical-1203120201033310-3133030122010002-0323020313133011-1120130103130213-2301132032001020-1202320301212320-0223302112120221-2310300312332200) |
| `name` | [name](resources--dns_lb_pool--reference--group-001.md#canonical-1313120011002132-2230032001101322-0331001313223222-3232301211102123-0321302031200011-2023210232010233-2131322003210111-2133201323022202) |
| `namespace` | [namespace](resources--dns_lb_pool--reference--group-001.md#canonical-0220130200032203-2121130032332311-2211313103301010-1023222103121013-3311203221012200-3220330222200003-2000211133131013-1303201131032202) |
| `srv_pool` | [srv_pool](resources--dns_lb_pool--reference--group-001.md#canonical-0132111133211010-3120131331231103-2313111210022023-1323221320203033-2120133010011311-3013000201212213-0303023302203031-2310132012022023) |
| `srv_pool.max_answers` | [srv_pool.max_answers](resources--dns_lb_pool--reference--group-001.md#canonical-2021020311010103-3213022022101310-0233101003332132-0300220121310012-2332003102102111-1133022313101033-3310212031100233-0312220213200310) |
| `srv_pool.members` | [srv_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-0223330020311313-2201311200122113-2132011120313323-3300132223300120-3022011333203010-3200330102220300-0120331220002132-0130222120133100) |
| `srv_pool.members.final_translation` | [srv_pool.members.final_translation](resources--dns_lb_pool--reference--group-001.md#canonical-0202200210232101-2033112221012022-3103301300031031-3012121201002033-0223231101110332-2013300202233202-1003320130023002-3132220010021323) |
| `srv_pool.members.name` | [srv_pool.members.name](resources--dns_lb_pool--reference--group-001.md#canonical-0012110233010103-1322013233012301-0101323321202232-1021323333112211-0103211311301133-3232031021020103-1220322222200132-3102030330021020) |
| `srv_pool.members.port` | [srv_pool.members.port](resources--dns_lb_pool--reference--group-001.md#canonical-0013332121203113-1321301121220100-2332203230121122-3001231020121332-0331312210313301-2101333202102031-2030230302011200-3332130300110223) |
| `srv_pool.members.priority` | [srv_pool.members.priority](resources--dns_lb_pool--reference--group-001.md#canonical-0213100322302132-2013113230000312-3201211102122023-3231030121231313-1332220022213212-2323002131022300-3221310033030211-1100232323332033) |
| `srv_pool.members.ratio` | [srv_pool.members.ratio](resources--dns_lb_pool--reference--group-001.md#canonical-0102233233231223-3113022102223212-3123032320103323-1112231321111030-2203133120113102-3120311321123001-0131121110012101-1133311013230332) |
| `srv_pool.members.target` | [srv_pool.members.target](resources--dns_lb_pool--reference--group-001.md#canonical-0021330011031312-1123202123103013-2111312331321121-1201102020102011-3322131221101033-2301032223221002-3202130232323311-3333112201300130) |
| `srv_pool.members.weight` | [srv_pool.members.weight](resources--dns_lb_pool--reference--group-001.md#canonical-1312321002330213-3221130131031012-2212311111023300-2310302132221000-0211030203310011-1312111012132301-2012313032203223-3323223213111321) |
| `timeouts` | [timeouts](resources--dns_lb_pool--reference--group-001.md#canonical-1233100210123023-1300301130302300-0221113231132310-2003120110031211-0130113201330231-0101123212113203-3330331310333300-3102001213330100) |
| `timeouts.create` | [timeouts.create](resources--dns_lb_pool--reference--group-001.md#canonical-0001122301000122-2123022123030021-2311130000130003-0321003122323232-2131223303122311-1331131012321022-0121120230321212-0331003120302321) |
| `timeouts.delete` | [timeouts.delete](resources--dns_lb_pool--reference--group-001.md#canonical-1211310203022000-1302111321201003-0120132300023020-2020313212113320-3101201210110323-1223222133003333-2132132210131211-2103330222322333) |
| `timeouts.read` | [timeouts.read](resources--dns_lb_pool--reference--group-001.md#canonical-3202121323320210-0301003023101101-2021110112110102-2103323033102201-3333010100122213-3221321000312301-2000032310232000-0103120201032302) |
| `timeouts.update` | [timeouts.update](resources--dns_lb_pool--reference--group-001.md#canonical-2102032131233013-1101131000032000-0130331321022230-2222011010133233-1221120033113122-3210302311211101-3012223131111221-3122320312110321) |
| `ttl` | [TTL](resources--dns_lb_pool--reference--group-001.md#canonical-1331300301120230-3101310120203322-3322113323011220-0021331301102333-0202320103220032-1212112213130313-3100112333023102-0323231212232212) |
| `use_rrset_ttl` | [use_rrset_ttl](resources--dns_lb_pool--reference--group-001.md#canonical-1110333213023111-2321322223321202-2311131320220203-2111101222113310-0212111202332331-0302323323222331-0331013011332111-1133033111203113) |

<a id="canonical-0123311331023030-2123210223210232-2020121111022213-0033300301133130-3013123123122303-2220212232330101-0132302013203032-0000203310030023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `a_pool` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- a_pool

<a id="canonical-3233100323022233-0123010032011212-0122212003020230-3111022302320302-3010220211312232-1323231022212203-0103011300303323-1232130130031001"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: a\_pool, aaaa\_pool, cname\_pool, mx\_pool, srv\_pool\] Pool for A Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("max_answers",
    "members"),
  validators.ConflictingObjectAttributes("disable_health_check",
    "health_check")}
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
  "x-ves-oneof-field-health_check_choice": "[\"disable_health_check\",\"health_check\"]"
}
```

OneOf alternatives in this subsection:

- [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-3233100323022233-0123010032011212-0122212003020230-3111022302320302-3010220211312232-1323231022212203-0103011300303323-1232130130031001)
- [aaaa_pool](resources--dns_lb_pool--reference--group-001.md#canonical-0303103131133323-1020303203213322-3033133312011032-2010312120323223-3002021320230221-2023013321032330-0203203210103331-1323121111113302)
- [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-1102013100110222-1003231012010002-1133112100231111-0303132130313220-2222201031010130-3313111310133311-2132322201221201-3130012321101032)
- [mx_pool](resources--dns_lb_pool--reference--group-001.md#canonical-2100123321012023-3313133333301013-2131313230200003-3313332331322203-2310021100232201-2003031322230202-1320301323312233-1320211103320332)
- [srv_pool](resources--dns_lb_pool--reference--group-001.md#canonical-0132111133211010-3120131331231103-2313111210022023-1323221320203033-2120133010011311-3013000201212213-0303023302203031-2310132012022023)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
a_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131231210213301-2020120212113032-2100113013002233-2302301002221333-2300030011032130-0023111113013310-1032220330002302-2110013020000332"></a>

### Direct properties for `a_pool`

- [disable_health_check](resources--dns_lb_pool--reference--group-001.md#canonical-0210213002213000-2202232100103103-0110323232101021-3003200012313322-1310021301300000-1313202210221010-1332112122100030-3031332233202111): complete subsection reference.

- [health_check](resources--dns_lb_pool--reference--group-001.md#canonical-0323310313033220-2023103012110012-1332101203113022-2330223030302001-2101022032313221-1330201213313330-2331102103012132-3132302020113221): complete subsection reference.

<a id="canonical-0122332113130303-3020231301000133-0012111202121023-1321220020123110-2322010131331313-1331100300130212-0230330220310301-3110130133132001"></a>

<a id="canonical-2231102201101220-2112311023013111-3021113122322322-1220210211132131-2011321031130321-3102021301213211-0133122001030122-0101032020202021"></a>

#### `a_pool.max_answers` property

Type: `"number"`. Optional.

Limit on number of Resource Records to be included in the response to query.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](resources--dns_lb_pool--reference--group-001.md#canonical-2122022111130320-0001001323110101-3210200302200311-2323100330302110-0013332212221000-3332301300232022-2211213222110222-1102213000320130): complete subsection reference.

<a id="canonical-0210213002213000-2202232100103103-0110323232101021-3003200012313322-1310021301300000-1313202210221010-1332112122100030-3031332233202111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `a_pool.disable_health_check` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-0123311331023030-2123210223210232-2020121111022213-0033300301133130-3013123123122303-2220212232330101-0132302013203032-0000203310030023)
- a_pool.disable_health_check

<a id="canonical-2211013022000201-1330111311320313-3101220030201311-3321311331101101-1032102211230330-2220222223002312-1302101011313203-3200133313333122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable health check.

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
disable_health_check = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323310313033220-2023103012110012-1332101203113022-2330223030302001-2101022032313221-1330201213313330-2331102103012132-3132302020113221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `a_pool.health_check` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-0123311331023030-2123210223210232-2020121111022213-0033300301133130-3013123123122303-2220212232330101-0132302013203032-0000203310030023)
- a_pool.health_check

<a id="canonical-1333122310212023-3223300103002330-1111023120031101-1120230113320232-3031020220331110-3332212313111101-2120121012120300-1011022322301232"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111211223201130-0002130313332122-0213033000013310-2330230230010100-3323022112321320-2222220021210203-0301002130210332-2322220200302220"></a>

### Direct properties for `a_pool.health_check`

<a id="canonical-0203000013010030-2000301233222002-0321303002322332-3222021303311112-0100331303312330-0321103123200313-3323300201033102-3201101123223212"></a>

#### `a_pool.health_check.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0021003132211223-0100102331221113-2312103330031000-3010030023031320-0120013231233322-2000311202001103-3331020323003033-0212130232030130"></a>

<a id="canonical-0331010001331200-0010320232013302-2101222203013331-2131000010203230-2312232031011323-0132130123132131-1302101302332120-2112320323031113"></a>

#### `a_pool.health_check.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2010132021130112-0310130002230100-0132100222130200-2233303120032111-1011220201133123-1311233233233301-3203022112312212-3102203323311333"></a>

<a id="canonical-2322300102223233-2310203220323130-1230311233031230-1330113111303231-1102003120212221-0010232212020101-3032201212130200-3021113220031021"></a>

#### `a_pool.health_check.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2122022111130320-0001001323110101-3210200302200311-2323100330302110-0013332212221000-3332301300232022-2211213222110222-1102213000320130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `a_pool.members` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-0123311331023030-2123210223210232-2020121111022213-0033300301133130-3013123123122303-2220212232330101-0132302013203032-0000203310030023)
- a_pool.members

<a id="canonical-1110023220310033-0010311113030201-3232222202032300-1300201332310113-2202232031203213-2003112111121230-3131003033023101-2330321313303101"></a>

Type: `"object"`. list nested block, Optional.

Pool Members. Configuration parameter for members

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_endpoint")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
members {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300022021211011-1213123031002101-1013320131001110-2233102323330300-1313331030202222-1300232311213310-1122202031113220-0223231221233231"></a>

### Direct properties for `a_pool.members`

<a id="canonical-3103002321130201-3223330212303233-0033110301012031-2100202013201022-0200122332233123-3121323231000202-2133222111012200-0333001000130123"></a>

#### `a_pool.members.disable_spec` property

Type: `"bool"`. Optional.

Value of true will disable the pool-member.

<a id="canonical-0303200202321110-3201110021132001-1210131000020232-2211212000203311-0120212220232120-2010310002013110-1311302221301232-1012002112213230"></a>

<a id="canonical-2311102001112032-0310002332222233-2010332001121210-1130101220011002-1102230302221031-3110301302202103-1000310223010220-3330001321020113"></a>

#### `a_pool.members.ip_endpoint` property

Type: `"string"`. Optional.

Public IP. Public IP address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-1310103123320232-2332321333202021-1222211101122123-0020023323113101-2231221013112221-1120033212331030-0231020023303320-0311130231002002"></a>

<a id="canonical-1133113202233032-1102213011222310-3123221002100122-2101203302323301-0320233331331203-1011101311301001-1231013103021112-3321101233312030"></a>

#### `a_pool.members.name` property

Type: `"string"`. Optional.

Name. Pool member name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3330322033213101-1103311303103122-0212300312211200-1333302323233110-0321031132003020-0331333010120300-0120223210200201-1003333222233011"></a>

<a id="canonical-1310230032211011-2322302121331033-1313300310230222-0123210100331013-0322012100021001-1321222232322330-1220110202101303-2012100001310203"></a>

#### `a_pool.members.priority` property

Type: `"number"`. Optional.

Used if the pool’s load balancing mode is set to Priority.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-1110033030021233-3303110123212113-2203300010011211-1132200110120313-1023200132332102-2120310202011032-1030013202223122-2011010220020122"></a>

<a id="canonical-0130220303112120-2222032332033002-3221230133000103-0300323021013123-2311133231222213-2003200232013030-2202032231303320-2110221030233203"></a>

#### `a_pool.members.ratio` property

Type: `"number"`. Optional.

Used if the pool’s load balancing mode is set to Ratio-Member.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-2212212022130300-3213012311313203-2210121222221333-2013311220021222-0022110313202023-1323111030012311-3031202030303213-0113222132323033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aaaa_pool` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- aaaa_pool

<a id="canonical-0303103131133323-1020303203213322-3033133312011032-2010312120323223-3002021320230221-2023013321032330-0203203210103331-1323121111113302"></a>

Type: `"object"`. single nested block, Optional.

Pool for AAAA Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("max_answers",
    "members")}
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
aaaa_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032032201120223-2021221120112301-2123013033331332-1010121300022120-2203301121330232-1131323032032120-1323022111323012-3003011022332323"></a>

### Direct properties for `aaaa_pool`

<a id="canonical-3022111311322010-2310103001200022-1322220222212133-1223332231313101-3112013112021101-0311321030332233-2031032320320220-1123230113302332"></a>

#### `aaaa_pool.max_answers` property

Type: `"number"`. Optional.

Limit on number of Resource Records to be included in the response to query.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](resources--dns_lb_pool--reference--group-001.md#canonical-2211031203232202-2210321003132211-0301303333203323-2312120220221020-0111000122011332-2010012000313200-0132321112312312-1303031000331121): complete subsection reference.

<a id="canonical-2211031203232202-2210321003132211-0301303333203323-2312120220221020-0111000122011332-2010012000313200-0132321112312312-1303031000331121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aaaa_pool.members` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- [aaaa_pool](resources--dns_lb_pool--reference--group-001.md#canonical-2212212022130300-3213012311313203-2210121222221333-2013311220021222-0022110313202023-1323111030012311-3031202030303213-0113222132323033)
- aaaa_pool.members

<a id="canonical-0223322323321233-3010122103023230-3030003000333310-0321211322010013-1303111220311330-0203202332300303-3323323331233203-0203023011003000"></a>

Type: `"object"`. list nested block, Optional.

Pool Members. Configuration parameter for members

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_endpoint")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
members {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020112122332320-3101220101322001-0302221220130031-0202230200221330-0211203113030231-1322302233131111-0000200112011120-2101033130312320"></a>

### Direct properties for `aaaa_pool.members`

<a id="canonical-0223013111000213-2313010220020232-0023112223300220-1201020003230230-3333201302012300-1023030032100230-2231301021011303-3323023312023000"></a>

#### `aaaa_pool.members.disable_spec` property

Type: `"bool"`. Optional.

Value of true will disable the pool-member.

<a id="canonical-1231232013023321-3033331310020033-0102023010310221-1122001023201210-3233310312032032-2303301103211321-3010333230333021-2233202222131200"></a>

<a id="canonical-0003223123210010-2130000301022101-3211323002310300-0021333112311113-2001013103030302-3211133111211113-3111123311323331-3321010003113011"></a>

#### `aaaa_pool.members.ip_endpoint` property

Type: `"string"`. Optional.

Public IP. Public IP address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-0111001132310023-2002323032233203-1121221231300112-2002330333030121-0101121210000300-2031311012113123-2012213110313232-2203002122310222"></a>

<a id="canonical-2031202122202320-0110030322332120-2102220130110003-2332312323302003-0122112021331122-2313100330200302-0220121231212300-1032212233012113"></a>

#### `aaaa_pool.members.name` property

Type: `"string"`. Optional.

Name. Pool member name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0000333232313003-3110112331102130-1120121223230112-2222022033110021-0220203130120222-0121233110002030-0012031212021122-2223223132013211"></a>

<a id="canonical-0001200101203003-3112213120203032-1231002013123123-0130233313301031-3021323123320121-0230333112120333-2321102301211132-3020000230233013"></a>

#### `aaaa_pool.members.priority` property

Type: `"number"`. Optional.

Used if the pool’s load balancing mode is set to Priority.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-1000131231202103-2112133311300102-0020233312312312-3023211210111233-1000303210132130-2311011200000020-3031312211110132-3131030010212013"></a>

<a id="canonical-2020300100331321-2030211222220003-0320313222102302-0330303010222220-2303200001010202-2003121231300223-3010112230311322-1003001031022131"></a>

#### `aaaa_pool.members.ratio` property

Type: `"number"`. Optional.

Used if the pool’s load balancing mode is set to Ratio-Member.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-2102130031203200-0130012310030313-1203320121212201-1330323012031100-2012202003023332-0303332132013312-0103012331231303-1323123021230220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cname_pool` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- cname_pool

<a id="canonical-1102013100110222-1003231012010002-1133112100231111-0303132130313220-2222201031010130-3313111310133311-2132322201221201-3130012321101032"></a>

Type: `"object"`. single nested block, Optional.

Pool for CNAME Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("members"),
  validators.ConflictingObjectAttributes("disable_health_check",
    "health_check")}
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
  "x-ves-oneof-field-health_check_choice": "[\"disable_health_check\",\"health_check\"]"
}
```

Terraform syntax:

```terraform
cname_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1022210000332302-1333130030111133-1300303102212333-3232133310301310-2012301121213131-1103210030130300-3220321130032012-1121112013003030"></a>

### Direct properties for `cname_pool`

- [disable_health_check](resources--dns_lb_pool--reference--group-001.md#canonical-2213312121000132-2321302000121322-3231020123322033-2011210210232100-3222320312103232-2001130103001013-1223301213121001-2023133333112302): complete subsection reference.

- [health_check](resources--dns_lb_pool--reference--group-001.md#canonical-3320003110222211-2232310102031032-3020111002212112-2001331112330220-1303211201131321-0111101310221021-2202021203003121-3320332310322003): complete subsection reference.

- [members](resources--dns_lb_pool--reference--group-001.md#canonical-1201013212131122-1333312112201001-1233033132210330-1122201133023222-1200233333310033-0023320210210233-2211213313311002-3302101302303120): complete subsection reference.

<a id="canonical-2213312121000132-2321302000121322-3231020123322033-2011210210232100-3222320312103232-2001130103001013-1223301213121001-2023133333112302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cname_pool.disable_health_check` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-2102130031203200-0130012310030313-1203320121212201-1330323012031100-2012202003023332-0303332132013312-0103012331231303-1323123021230220)
- cname_pool.disable_health_check

<a id="canonical-2011102101303300-0212233232002101-2200221301121100-0312103323012221-1223210113221131-0330301132111131-3323303102000212-0231312020223231"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable health check.

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
disable_health_check = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320003110222211-2232310102031032-3020111002212112-2001331112330220-1303211201131321-0111101310221021-2202021203003121-3320332310322003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cname_pool.health_check` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-2102130031203200-0130012310030313-1203320121212201-1330323012031100-2012202003023332-0303332132013312-0103012331231303-1323123021230220)
- cname_pool.health_check

<a id="canonical-1321011122120221-3203100211011303-0301300133112220-2023303220123202-1020121020113011-1210211332130212-0112313312103203-1011303223021021"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203201312102203-0303233113332112-3330322231232032-2313113001030310-2132301030020221-0313020332220002-2320113110203200-0300233320312031"></a>

### Direct properties for `cname_pool.health_check`

<a id="canonical-0220120022010302-1220223310212131-0330321212101221-1222230300101103-3032123303100210-2130031110220020-0331112333033302-3223231012222030"></a>

#### `cname_pool.health_check.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1301220031313131-2133133333131132-0320130333101301-1231300212213023-0223221023111100-1232110010013202-0021102023113133-0102300032001010"></a>

<a id="canonical-3022220333200223-3301133213012212-1120031022132212-1003212210230012-0231200102312232-2333122022132101-3320010011232210-0321131021303120"></a>

#### `cname_pool.health_check.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3110010200312110-1221031232210002-0203331233130303-3233013033203010-3020003000313111-3220202020311201-0032203123321112-3212133232131121"></a>

<a id="canonical-1000222132303111-1010203320123030-3301031030230320-0111132003120210-0330310213021122-1310323001331023-3333133121332011-0223301103003320"></a>

#### `cname_pool.health_check.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1201013212131122-1333312112201001-1233033132210330-1122201133023222-1200233333310033-0023320210210233-2211213313311002-3302101302303120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cname_pool.members` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-2102130031203200-0130012310030313-1203320121212201-1330323012031100-2012202003023332-0303332132013312-0103012331231303-1323123021230220)
- cname_pool.members

<a id="canonical-3130032103000222-1011103203200111-3322133132311001-2030132223132010-1131322222111111-1310110331033301-2021212003310110-3200131002220011"></a>

Type: `"object"`. list nested block, Optional.

Pool Members. Configuration parameter for members

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
members {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132120300330331-0211002233102031-2330202112111320-3001311101230232-0333131313103112-2112230232131210-0300213310110113-2102203002303033"></a>

### Direct properties for `cname_pool.members`

<a id="canonical-0023321322111013-3032000321303333-2122223001122233-3300000200213222-3033211322300132-3013122030020333-1120131201023122-2301033120302011"></a>

#### `cname_pool.members.domain` property

Type: `"string"`. Optional.

Specifies the fully qualified domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-0211212203300330-3110212133300101-2110112012032013-3310212123201121-0012211213111321-2233122023000213-1303021013210003-0213303303233202"></a>

<a id="canonical-3021133020132003-0013130203113112-2123131212103032-3213021210333121-1232232211010101-2101001221132110-2023233320203010-3132303233302001"></a>

#### `cname_pool.members.final_translation` property

Type: `"bool"`. Optional.

If this flag is true, the CNAME record will not be translated further.

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

<a id="canonical-3331123003212032-0320121311010220-2132303033103011-3131200330223310-3132013003332211-0110310222130320-0011020132310231-3020302113033320"></a>

<a id="canonical-0303232323020001-3032311203312123-0102012303112310-2120323213203130-0330103312221321-1101120320132313-3131032113310130-0102112133013113"></a>

#### `cname_pool.members.name` property

Type: `"string"`. Optional.

Name. Pool member name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0301010310231333-3213332201112122-1131303320233202-3302033103310230-1013233330303013-2311322220312321-3213211002032203-0013320232212330"></a>

<a id="canonical-3002123203313111-1302120030312021-1220223132010121-0200333032220203-2201122233012300-3130201233101102-1132313123313313-3233302211302330"></a>

#### `cname_pool.members.priority` property

Type: `"number"`. Optional.

Used if the pool’s load balancing mode is set to Priority. Determines the order in which traffic is
routed to pool members. The lower the number, the higher the priority, making those members active
while higher-numbered members act as backups.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-2023002322313010-1222021210203031-1201332213302001-2322002001211100-2012113022203203-3332333102333222-0023022133032332-1023021320032210"></a>

<a id="canonical-3310232201022210-0221213303000321-3311323123002213-3000311020320020-0121113200010031-1322033223303133-0321323112112312-0220111301332110"></a>

#### `cname_pool.members.ratio` property

Type: `"number"`. Optional.

Used if the pool’s load balancing mode is set to Ratio-Member.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-2212032032130133-2301313011222230-0123131011123021-0200203312302320-1232223301102320-1012121332111002-1123122031330210-2100311112312320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mx_pool` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- mx_pool

<a id="canonical-2100123321012023-3313133333301013-2131313230200003-3313332331322203-2310021100232201-2003031322230202-1320301323312233-1320211103320332"></a>

Type: `"object"`. single nested block, Optional.

Pool for MX Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("max_answers",
    "members")}
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
mx_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1002230023303112-2222221130300011-1101120003132021-0132212311310100-0230102132212333-1233013110022221-1322310330122112-2102002313300203"></a>

### Direct properties for `mx_pool`

<a id="canonical-0310110113000210-2323022131111101-1130330321132200-1313132110210331-2233030131320012-1223032311230221-0311032133031130-2131023122331122"></a>

#### `mx_pool.max_answers` property

Type: `"number"`. Optional.

Limit on number of Resource Records to be included in the response to query.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](resources--dns_lb_pool--reference--group-001.md#canonical-3102332222023330-2212221133033033-1010313321202203-2230023330301010-1133121222312112-3101020211302122-3323300231003121-0222103330303023): complete subsection reference.

<a id="canonical-3102332222023330-2212221133033033-1010313321202203-2230023330301010-1133121222312112-3101020211302122-3323300231003121-0222103330303023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mx_pool.members` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- [mx_pool](resources--dns_lb_pool--reference--group-001.md#canonical-2212032032130133-2301313011222230-0123131011123021-0200203312302320-1232223301102320-1012121332111002-1123122031330210-2100311112312320)
- mx_pool.members

<a id="canonical-2021002222021132-0302010122212200-0230103011021022-1010310312021130-0322330103320013-2332223123232122-3033213121321232-1321233020331201"></a>

Type: `"object"`. list nested block, Optional.

Pool Members. Configuration parameter for members

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
members {
  # Configure direct properties listed below.
}
```

<a id="canonical-2132210113112312-3323130110231321-3330033121021133-2231123222201330-3101020232000223-0213321310032120-1331100303103232-3203330212002132"></a>

### Direct properties for `mx_pool.members`

<a id="canonical-2002230003312332-0321130231133231-3201231303232120-1232313202311131-0213300010022031-0101012110213200-0220100131123203-0123320331010001"></a>

#### `mx_pool.members.domain` property

Type: `"string"`. Optional.

Domain name for routing and identification.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-0310212033330210-1100102230333231-2021120123220100-0320210122122320-2210231123212002-1102103303111123-1101002311113033-2300211323121133"></a>

<a id="canonical-2013310032322123-2000100333331010-3000031013130303-0132210020200213-2230101211101331-2000301023113321-2012201010312012-3223212100300003"></a>

#### `mx_pool.members.name` property

Type: `"string"`. Optional.

Name. Pool member name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2031330212201130-2121100230210110-3323112000320303-2100232111230122-1013101313100212-3111300310321102-3313012112233011-0221100323122002"></a>

<a id="canonical-2132033122001023-0102232231323313-0131023122013113-2210313311323111-3320122130330301-1021221202213133-0231333211132011-1301311301321231"></a>

#### `mx_pool.members.priority` property

Type: `"number"`. Optional.

MX Record Priority. MX Record priority.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1203120201033310-3133030122010002-0323020313133011-1120130103130213-2301132032001020-1202320301212320-0223302112120221-2310300312332200"></a>

<a id="canonical-3132310033223203-2000013120311331-2000033100303021-3120123122131212-0212301330321122-0311113001213002-0201021323111302-2223230023123302"></a>

#### `mx_pool.members.ratio` property

Type: `"number"`. Optional.

Load Balancing Ratio. Load Balancing Ratio.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-3000311322032000-2023201012210313-3122222330203102-3333210122112323-1033111323322312-3122231232003322-3300111021330330-2220331310311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `srv_pool` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- srv_pool

<a id="canonical-0132111133211010-3120131331231103-2313111210022023-1323221320203033-2120133010011311-3013000201212213-0303023302203031-2310132012022023"></a>

Type: `"object"`. single nested block, Optional.

Pool for SRV Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("max_answers",
    "members")}
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
srv_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120322213011111-1033310213300303-0232132233101130-2313133031330022-3323212322311201-0032131032031300-0023312122312213-3013300023100301"></a>

### Direct properties for `srv_pool`

<a id="canonical-2021020311010103-3213022022101310-0233101003332132-0300220121310012-2332003102102111-1133022313101033-3310212031100233-0312220213200310"></a>

#### `srv_pool.max_answers` property

Type: `"number"`. Optional.

Limit on number of Resource Records to be included in the response to query.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](resources--dns_lb_pool--reference--group-001.md#canonical-2222100011301023-2013331101303131-0100021202102233-2232032010003231-1111232023232301-3013023232233113-2110301003003233-1333220210322232): complete subsection reference.

<a id="canonical-2222100011301023-2013331101303131-0100021202102233-2232032010003231-1111232023232301-3013023232233113-2110301003003233-1333220210322232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `srv_pool.members` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- [srv_pool](resources--dns_lb_pool--reference--group-001.md#canonical-3000311322032000-2023201012210313-3122222330203102-3333210122112323-1033111323322312-3122231232003322-3300111021330330-2220331310311332)
- srv_pool.members

<a id="canonical-0223330020311313-2201311200122113-2132011120313323-3300132223300120-3022011333203010-3200330102220300-0120331220002132-0130222120133100"></a>

Type: `"object"`. list nested block, Optional.

Pool Members. Configuration parameter for members

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("port",
    "priority",
    "target",
    "weight")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
members {
  # Configure direct properties listed below.
}
```

<a id="canonical-1020323330231212-2120321031311233-1200232220033320-0031231302300230-1330122101000213-3221332100032302-3331303123000100-0112011131313222"></a>

### Direct properties for `srv_pool.members`

<a id="canonical-0202200210232101-2033112221012022-3103301300031031-3012121201002033-0223231101110332-2013300202233202-1003320130023002-3132220010021323"></a>

#### `srv_pool.members.final_translation` property

Type: `"bool"`. Optional.

If this flag is true, the SRV record will not be translated further.

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

<a id="canonical-0012110233010103-1322013233012301-0101323321202232-1021323333112211-0103211311301133-3232031021020103-1220322222200132-3102030330021020"></a>

<a id="canonical-3113003110203133-1332133132003301-1201000331233102-3221320213232113-2013221300000111-3331332010020013-2300110103231232-3323001220200220"></a>

#### `srv_pool.members.name` property

Type: `"string"`. Optional.

Name. Pool member name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0013332121203113-1321301121220100-2332203230121122-3001231020121332-0331312210313301-2101333202102031-2030230302011200-3332130300110223"></a>

<a id="canonical-1223123232010202-2003010303331131-0102210221312130-2210101121313213-0033330210332221-2102023332133210-0313020102211032-3120203301102102"></a>

#### `srv_pool.members.port` property

Type: `"number"`. Optional.

Port. Port on which the service can be found.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0213100322302132-2013113230000312-3201211102122023-3231030121231313-1332220022213212-2323002131022300-3221310033030211-1100232323332033"></a>

<a id="canonical-3011210111320313-3303301030023320-3233130023022211-2112231030021130-3133330230323213-0303120332030222-3001010211330313-0011332003322231"></a>

#### `srv_pool.members.priority` property

Type: `"number"`. Optional.

Priority of the target. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0102233233231223-3113022102223212-3123032320103323-1112231321111030-2203133120113102-3120311321123001-0131121110012101-1133311013230332"></a>

<a id="canonical-1111312203320313-2033320003121002-0211330102232210-2330233020013023-3301112332303001-1232201110002030-3321203011012003-1211031211122332"></a>

#### `srv_pool.members.ratio` property

Type: `"number"`. Optional.

Load Balancing Ratio. Configuration parameter for ratio

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-0021330011031312-1123202123103013-2111312331321121-1201102020102011-3322131221101033-2301032223221002-3202130232323311-3333112201300130"></a>

<a id="canonical-2331321012023331-2221312100010300-3311033112200101-3113010003311111-2202021302331002-1132011232301312-0030022010200100-1230010220200232"></a>

#### `srv_pool.members.target` property

Type: `"string"`. Optional.

Domain name of the machine providing the service.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  }
}
```

<a id="canonical-1312321002330213-3221130131031012-2212311111023300-2310302132221000-0211030203310011-1312111012132301-2012313032203223-3323223213111321"></a>

<a id="canonical-0003200212220032-2203121301313100-0003230320020301-3120012132113233-2011331210023310-3321200113232222-0201031011212110-0222121123210112"></a>

#### `srv_pool.members.weight` property

Type: `"number"`. Optional.

Weight of the target. A higher number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0301322011101121-0101320022303020-2233133301222122-3123332332021333-3021322031131312-1333202003120133-0311330103332133-1222331311003000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- timeouts

<a id="canonical-1233100210123023-1300301130302300-0221113231132310-2003120110031211-0130113201330231-0101123212113203-3330331310333300-3102001213330100"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321233300301003-3320131313312011-3121131311203102-2012211111122310-2101303222313031-2120210130221320-1122131233212320-0113003033313313"></a>

### Direct properties for `timeouts`

<a id="canonical-0001122301000122-2123022123030021-2311130000130003-0321003122323232-2131223303122311-1331131012321022-0121120230321212-0331003120302321"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1211310203022000-1302111321201003-0120132300023020-2020313212113320-3101201210110323-1223222133003333-2132132210131211-2103330222322333"></a>

<a id="canonical-0323231030010031-1020330031210031-0030333322203330-3003223031333113-0300011133221221-0033020220121033-3323220122200123-2313222203223023"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3202121323320210-0301003023101101-2021110112110102-2103323033102201-3333010100122213-3221321000312301-2000032310232000-0103120201032302"></a>

<a id="canonical-1130201333001120-0022220130223001-1312002211100112-1020203130303211-2103302231310331-0313320021110003-2230223222333132-1103303323200100"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2102032131233013-1101131000032000-0130331321022230-2222011010133233-1221120033113122-3210302311211101-3012223131111221-3122320312110321"></a>

<a id="canonical-0001122220002232-3120332300013332-2022220131321033-2311132120031010-0312210123201310-3203313112232203-3122011013112132-0101221300010332"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1201132013313011-0032030003133200-1120231001123032-1302100130003232-3310032232121310-1200122301101332-3100010223322121-0201133221211331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_rrset_ttl` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- use_rrset_ttl

<a id="canonical-1110333213023111-2321322223321202-2311131320220203-2111101222113310-0212111202332331-0302323323222331-0331013011332111-1133033111203113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use rrset TTL.

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
use_rrset_ttl = {}
```

This is an empty object or choice marker. It has no direct properties.
