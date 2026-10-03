---
page_title: "xcsh_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery reference."
---

# xcsh_discovery reference

<a id="canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021201133032301-0033000310101332-3033013303012231-3202202111033130-0132101323032001-1302302212213232-1032120103233100-1102331011333003"></a>

## Property reference — Property reference / 132033111333 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- Property reference

<a id="canonical-0112210012221233-1102320303221032-2200301033222202-1300011130222200-0110033330002101-2302220023022322-3110010330202112-3233111122323321"></a>

## Direct properties — Property reference / 132033111333 / 3

<a id="canonical-3232323001123233-2113012211000020-1102230202032102-1233132122210022-2020013210330010-1032023113332300-1203002131000031-0301132121313210"></a>

<a id="canonical-3313120030220033-1112211201313221-3131200110200321-0111322322002033-3103300211310011-0123302120110310-1212220311033222-2333312001300203"></a>

## annotations property — Property reference / 132033111333 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

<a id="canonical-0002012200122012-0013302230101010-1202212030113023-0110320230102331-0100331230122321-2002112022301221-0301203023000002-1011210203311030"></a>

<a id="canonical-3213101301221032-0231011131230311-2121023232223312-1220331000110220-0123232300331100-2101112202030323-2111211212100220-2012321311311133"></a>

## cluster_id property — Property reference / 132033111333 / 5

Type: `"string"`. Optional, Computed.

\[OneOf: cluster\_id, no\_cluster\_id; Default: no\_cluster\_id\] Exclusive with \[no\_cluster\_id\]
Specify identifier for discovery cluster. This identifier can be specified in endpoint object to
discover only from this discovery object.

Upstream description:

Exclusive with \[no\_cluster\_id\] Specify identifier for discovery cluster. This identifier can be
specified in endpoint object to discover only from this discovery object.

Provider validators and defaults (from schema source):

```go
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

OneOf alternatives in this subsection:

- [cluster_id](resources--discovery--reference--group-001.md#canonical-0002012200122012-0013302230101010-1202212030113023-0110320230102331-0100331230122321-2002112022301221-0301203023000002-1011210203311030)
- [no_cluster_id](resources--discovery--reference--group-001.md#canonical-3002323030220322-0020200231302202-3111002300320220-3022101033030022-2021313300131030-0300131003010012-3302102103231001-3130001221002120)

Select alternatives according to the provider validators above.

<a id="canonical-0230211332023021-3100200311332232-2212222123221211-0101203210112011-1322303232132122-1222130300202322-2321003120221322-1200003321313230"></a>

<a id="canonical-3301123122111011-0212310113132220-3331320321232222-1301123102122003-3100021201032100-2200222203313301-1220322002231113-2233231230010013"></a>

## description property — Property reference / 132033111333 / 6

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

<a id="canonical-1213310010120200-2200221131030022-0303100303302333-2002003130000210-0030030303202302-1330020022331012-2132111311023310-3333220230232100"></a>

<a id="canonical-1103213013033233-1320120321100200-2000101310010332-2302102211202020-1203202310003111-2002233011123233-2302213103321333-0320103113320002"></a>

## disable property — Property reference / 132033111333 / 7

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

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

- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233): complete subsection reference.

- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000): complete subsection reference.

<a id="canonical-2331201330020310-3011330213123030-3203322003002111-0100133132101320-3300312013110232-3010033210133322-1221231133311301-0210330122333311"></a>

<a id="canonical-3111102313102201-0133322203202123-2211000121211133-1230001013221212-2131302011303223-2123323100320013-3110231030003333-1011231030332033"></a>

## ID property — Property reference / 132033111333 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2220101102111200-2323310113332112-0132002200231312-1123011022011023-3200210221001203-1231003230231300-3321210213113200-3311121321002110"></a>

<a id="canonical-2132331002002220-1303200130020003-3102133110032010-2012311222013120-1201321222203232-1312100322013323-1211131102032022-2001222222310033"></a>

## labels property — Property reference / 132033111333 / 9

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-0231103223133302-1310022012112113-1303120002132223-1103111031233210-2302112001113222-2101033321321100-3033002001210122-1111331023213201"></a>

<a id="canonical-1111223233213031-1333212012120312-3001003101302111-1311030122223011-2221221013033310-0213001030103131-1310300101313102-0103003103012110"></a>

## name property — Property reference / 132033111333 / 10

Type: `"string"`. Required.

Name of the Discovery. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3031213032120321-0013232210133312-1333331030312323-1113322220213333-0100132022331333-1012202221103020-0130001103021021-0011233300302001"></a>

<a id="canonical-2200220320122201-3300101011302311-3310032112101021-3011112302332320-1103210310102032-0032333131303112-3330032320102210-1322022102011201"></a>

## namespace property — Property reference / 132033111333 / 11

Type: `"string"`. Required.

Namespace where the Discovery is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
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

- [no_cluster_id](resources--discovery--reference--group-001.md#canonical-2332001313120022-1313323002112231-2021132023022331-0131012222332133-3231233201102020-0102032212311132-1210113121202130-1000203132133103): complete subsection reference.

- [timeouts](resources--discovery--reference--group-001.md#canonical-0133122212320132-3323032213033022-3031001120032311-3120311312312302-0221333302003213-0021323013310132-3013211230313033-1331213120121300): complete subsection reference.

- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010): complete subsection reference.

<a id="canonical-0131133321032010-0021332210000000-3102123232231032-1002012131210011-3233120333100311-3110313013101213-3103103323322032-1020013221311320"></a>

## All schema paths — Property reference / 132033111333 / 12

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--discovery--reference--group-001.md#canonical-3232323001123233-2113012211000020-1102230202032102-1233132122210022-2020013210330010-1032023113332300-1203002131000031-0301132121313210) |
| `cluster_id` | [cluster_id](resources--discovery--reference--group-001.md#canonical-0002012200122012-0013302230101010-1202212030113023-0110320230102331-0100331230122321-2002112022301221-0301203023000002-1011210203311030) |
| `description` | [description](resources--discovery--reference--group-001.md#canonical-0230211332023021-3100200311332232-2212222123221211-0101203210112011-1322303232132122-1222130300202322-2321003120221322-1200003321313230) |
| `disable` | [disable](resources--discovery--reference--group-001.md#canonical-1213310010120200-2200221131030022-0303100303302333-2002003130000210-0030030303202302-1330020022331012-2132111311023310-3333220230232100) |
| `discovery_consul` | [discovery_consul](resources--discovery--reference--group-001.md#canonical-1030303101331321-3203302302330023-0021113131100321-2133033101120310-3100033000111203-0311120132222222-0133133213320000-3201332120011001) |
| `discovery_consul.access_info` | [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-3330303202231313-2123032003102222-3032212230121130-0301112230022000-2212130032202323-3110113201232220-3021233310013000-2301113132100311) |
| `discovery_consul.access_info.connection_info` | [discovery_consul.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-2020132232200102-1310012132320031-1021131030332030-2013223213222213-3312101131332222-1301103131223123-2211233010320322-0311221010021010) |
| `discovery_consul.access_info.connection_info.api_server` | [discovery_consul.access_info.connection_info.api_server](resources--discovery--reference--group-001.md#canonical-1133303222102323-0012213223111233-1113032003012133-2103311322122202-3132112121331001-3302222002301120-2231110213121300-0123032330103110) |
| `discovery_consul.access_info.connection_info.tls_info` | [discovery_consul.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-3120130210213221-1310202020132212-0213200313211210-3122322021102032-2100231122310012-0120100032011203-2101213021321110-2333303010331310) |
| `discovery_consul.access_info.connection_info.tls_info.certificate` | [discovery_consul.access_info.connection_info.tls_info.certificate](resources--discovery--reference--group-001.md#canonical-2301230213323312-0221202303110310-3210322121200223-2213110100103000-0022012313121201-1132321100311310-2223003132033133-2001330203110021) |
| `discovery_consul.access_info.connection_info.tls_info.key_url` | [discovery_consul.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-0322032013110131-0131332133302010-3123000133213333-1333130200110113-1110002010120010-2020011120113302-0121123102310132-1313121333132320) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-2133230200133322-2000110300110011-1023130121032311-1100023030022303-0322201311102200-0213112201201231-1211131033111303-0301303111221222) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider](resources--discovery--reference--group-001.md#canonical-3220220212302121-1332221330331101-3201012320231022-2110023003011013-2112321333123010-3111022003102332-3000000132303031-2101033021013012) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location](resources--discovery--reference--group-001.md#canonical-1322103030131202-2130131313200102-3323232230203010-0112231030100222-2112012201222011-1110212320101300-0213100003222030-0220313220023222) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider](resources--discovery--reference--group-001.md#canonical-1323130030102231-3333322333103321-3331133000021012-3231232013312032-1112010203110202-2230111312121132-1333211310313233-1212121023331123) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-3313200120113232-2230301010132103-1313101013210010-0110210210300221-2233210010110322-0302301230020323-0320312100210112-0201130202311120) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref](resources--discovery--reference--group-001.md#canonical-2210033212200030-3021323131010122-1130303131332011-1031311131323233-3011300302131103-0100122231122020-1320100120001221-1223023122223122) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.url` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.url](resources--discovery--reference--group-001.md#canonical-2130222322301313-3002130111120031-3222311113002231-0203122013213003-0023032300333330-0113221133030331-2121231202110031-2233201213203113) |
| `discovery_consul.access_info.connection_info.tls_info.server_name` | [discovery_consul.access_info.connection_info.tls_info.server_name](resources--discovery--reference--group-001.md#canonical-2113211120130033-0131121333100022-1122032023010223-3223112000111332-3313121303113112-1120002202000223-3002302202010030-3222213120231000) |
| `discovery_consul.access_info.connection_info.tls_info.trusted_ca_url` | [discovery_consul.access_info.connection_info.tls_info.trusted_ca_url](resources--discovery--reference--group-001.md#canonical-1212223230301001-0330020212032222-2202301102212312-0030220220033301-1300023323230301-2213030321203121-3200132003301130-2113321223033200) |
| `discovery_consul.access_info.http_basic_auth_info` | [discovery_consul.access_info.http_basic_auth_info](resources--discovery--reference--group-001.md#canonical-0123313323210131-0230300022001113-3102310033300331-1101113223332333-2320232301123003-2010032111102013-3220221323113300-0123102111233213) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url` | [discovery_consul.access_info.http_basic_auth_info.passwd_url](resources--discovery--reference--group-001.md#canonical-1232301021111123-2202332222031210-0200101321230103-1222032233121033-3210130233332030-0103323203220232-3020133322103003-0003321010333010) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-2011002002332312-0112022132233001-3223200101312203-3223010302220121-2210133233100231-2303132002220213-0112302002221111-3233120133123102) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.decryption_provider` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.decryption_provider](resources--discovery--reference--group-001.md#canonical-0132013020121313-0330012133100232-1330210312011222-0102213222102200-2121110232320132-3032312003211003-3201033201011203-0113130203013313) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.location` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.location](resources--discovery--reference--group-001.md#canonical-1231020032012213-2131223322102112-1032111330313020-1201331102010213-2133311220301300-3133131003331101-0033322121121223-1300223300010110) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.store_provider` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.store_provider](resources--discovery--reference--group-001.md#canonical-0232022130001320-1321112000032222-2001123003103002-2322031203323323-3303003000113310-2121220013110123-2210231223011012-0201012230200321) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-3020110303003233-3100031300030231-2231220032300221-1333031203003321-3103232131210110-0310200100122323-0300331232200213-0131223120013103) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.provider_ref` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.provider_ref](resources--discovery--reference--group-001.md#canonical-0201021033121121-2001022100022323-1122233001201200-1032100122210300-2332022233110131-0122213333032233-3103131303121320-2113032233320032) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.url` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.url](resources--discovery--reference--group-001.md#canonical-3331212301021302-1123121333002311-0231310222131210-2201310223011013-0333222010320311-0232333021021100-2001032330221310-0120130000211310) |
| `discovery_consul.access_info.http_basic_auth_info.user_name` | [discovery_consul.access_info.http_basic_auth_info.user_name](resources--discovery--reference--group-001.md#canonical-2132202320222323-1121113303312011-2112233233021003-2220210011122031-1212111313230010-1030200133101213-1110332030111032-2221020223100303) |
| `discovery_consul.publish_info` | [discovery_consul.publish_info](resources--discovery--reference--group-001.md#canonical-0112121310302232-1031213213023031-1232201113331111-3221201010312220-0013300131203133-3300221321011323-0321033123022213-2330323000232211) |
| `discovery_consul.publish_info.disable_spec` | [discovery_consul.publish_info.disable_spec](resources--discovery--reference--group-001.md#canonical-3221112130031131-0211203101310202-1330101002321330-1312123122202003-0233002322031110-1122021123332010-3031230031110032-1323321212000121) |
| `discovery_consul.publish_info.publish` | [discovery_consul.publish_info.publish](resources--discovery--reference--group-001.md#canonical-1201030231033011-1122023123320031-0321020002313233-1121213310311302-2132133110023212-2310332013221222-3202132223101021-1231032201231331) |
| `discovery_k8s` | [discovery_k8s](resources--discovery--reference--group-001.md#canonical-0303032133222231-3132132031112130-2102022332112001-0233203030233332-3011132230201022-1131031022000310-1002230130133033-1310212302021220) |
| `discovery_k8s.access_info` | [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-0130013031001132-2223032223121222-1313121231220323-2300100012211233-1002101110133103-2212320113310231-0310211110122230-1021232103100020) |
| `discovery_k8s.access_info.connection_info` | [discovery_k8s.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-3011220101012032-1302231313202322-2332330323223100-3321211321321131-1113032201322313-2033223123031221-0123210020223332-0003211332111332) |
| `discovery_k8s.access_info.connection_info.api_server` | [discovery_k8s.access_info.connection_info.api_server](resources--discovery--reference--group-001.md#canonical-0200221101102201-3303000322031220-3232032223310002-3030013002320013-0312303111333222-1232212030023222-3310021101303120-0030021132033230) |
| `discovery_k8s.access_info.connection_info.tls_info` | [discovery_k8s.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-3030022132102000-2031102201323102-3113100222010330-3300213113203031-1101233123002001-3010120103132320-0020103003230310-1103311223122223) |
| `discovery_k8s.access_info.connection_info.tls_info.certificate` | [discovery_k8s.access_info.connection_info.tls_info.certificate](resources--discovery--reference--group-001.md#canonical-3110323021212022-1000313012233123-3013103311303101-0022001131132121-0133110020322001-1021330121111203-2012232320222023-2330210010101202) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url` | [discovery_k8s.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-2220020133132011-0310001033322331-2233100020320200-0103130122203010-3230030213012100-3333220232330233-2133233331013120-2003202131013020) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-3130131223001003-1123001123132131-1313231121301233-2012103101020121-0100203311321032-1032132333233322-3323301030122332-2022231021202003) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider](resources--discovery--reference--group-001.md#canonical-0312230222100310-1021131022101323-3321222222113132-0100021301102121-2002323132202102-1301132331000323-2231132100030301-2120121230021032) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location](resources--discovery--reference--group-001.md#canonical-1221323132122311-3130132201123331-2013320121201032-1113020022123320-2021022223300311-2211211113213212-1003230013130002-3200123230121313) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider](resources--discovery--reference--group-001.md#canonical-2330021003221202-3300032202331212-1111131000233330-0121121313320222-3111332000033313-3322332212223201-1021032332313221-2230100330021202) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-0121112223310011-2200111031121311-0213131133330232-2101300310013320-3102130120120332-3203020232330101-3100212001023302-0013331132323113) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref](resources--discovery--reference--group-001.md#canonical-2331322022300123-2310321011223113-3131200021312131-0330011303211121-2121130130311123-0210310122201110-1310033203232322-3310111311233032) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.url` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.url](resources--discovery--reference--group-001.md#canonical-2301132332001300-2200123320212000-0233323102223302-3122121032003130-1033300121330231-3003323030133210-2320031031121303-3103021313103101) |
| `discovery_k8s.access_info.connection_info.tls_info.server_name` | [discovery_k8s.access_info.connection_info.tls_info.server_name](resources--discovery--reference--group-001.md#canonical-0011013301122113-0011002320213102-3010301133220331-0232231320123020-0200222132013102-3021022010301003-1212013131210111-1103232123211313) |
| `discovery_k8s.access_info.connection_info.tls_info.trusted_ca_url` | [discovery_k8s.access_info.connection_info.tls_info.trusted_ca_url](resources--discovery--reference--group-001.md#canonical-0100313202032302-1130311202120020-3110313022132332-2011101001301232-1231232222133121-0003233221333223-0011333231121330-1003210213133130) |
| `discovery_k8s.access_info.isolated` | [discovery_k8s.access_info.isolated](resources--discovery--reference--group-001.md#canonical-3211300221101310-3111000233302102-0003303033023023-0013221003331123-1010310010222310-2323120000222312-2003211221021101-3103201311333023) |
| `discovery_k8s.access_info.kubeconfig_url` | [discovery_k8s.access_info.kubeconfig_url](resources--discovery--reference--group-001.md#canonical-1033220201210013-0311033123222002-3212333323131300-0213032312230030-0101322010020222-0302103230121203-2033313012320120-3023033010100033) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-3102230331112221-2100230112012011-1213331102102001-1033013222121133-1321300232220200-3103211011110120-3201221221313211-1102010012101102) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.decryption_provider` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.decryption_provider](resources--discovery--reference--group-001.md#canonical-1103303221030211-2100031021202022-1103001330010000-1112000113220022-1131001210023200-3201010100303000-1012131100132131-2113231232320102) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.location` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.location](resources--discovery--reference--group-001.md#canonical-2032131212111000-2113001030322321-3103321103310131-3132300103123211-1000103021322300-3212300331122232-1110320333030020-3110313212203301) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.store_provider` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.store_provider](resources--discovery--reference--group-001.md#canonical-0300312312030123-0300210331211232-0102212022211232-0330031300100231-2331101300312113-1101112230112111-3332110333331023-1301032122010120) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-0211013121120001-1030303311113122-0201210030321322-1123032022123210-0330030221311111-0321211123331033-1302021022031320-2133101012221301) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.provider_ref` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info.provider_ref](resources--discovery--reference--group-001.md#canonical-2321131111120312-3223313210213103-3202010223011233-3322230223003033-0213110012232220-0333301230031312-3101233002100232-2230232033203233) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.url` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info.url](resources--discovery--reference--group-001.md#canonical-2301313313331230-3323110211231223-1221121122111023-2201111121121103-1002122012011313-2013200302123112-1302330001103103-1213111230320231) |
| `discovery_k8s.access_info.reachable` | [discovery_k8s.access_info.reachable](resources--discovery--reference--group-001.md#canonical-2112131220231232-0301102320201020-1232302112130232-1133221121220011-2330130313333122-3230001102122021-1022033201022213-1332332002220132) |
| `discovery_k8s.default_all` | [discovery_k8s.default_all](resources--discovery--reference--group-001.md#canonical-2033200123113032-3001322221200211-2100033312102313-0312103132132013-1011101130212313-0001000000233301-2110230201220220-1320132000312113) |
| `discovery_k8s.namespace_mapping` | [discovery_k8s.namespace_mapping](resources--discovery--reference--group-001.md#canonical-1130310213023212-1332013022102022-1111331100303002-2021113202231020-1230020021122231-3322220200003000-2322230130000120-0232111221000111) |
| `discovery_k8s.namespace_mapping.items` | [discovery_k8s.namespace_mapping.items](resources--discovery--reference--group-001.md#canonical-3112000132011301-1222010321120031-3203310201120101-2133032230221032-1210122030302003-2113022110300221-2203221203330020-3021321032100331) |
| `discovery_k8s.namespace_mapping.items.namespace` | [discovery_k8s.namespace_mapping.items.namespace](resources--discovery--reference--group-001.md#canonical-3201112330030203-1133302310200330-1102230113001203-3002032231013323-3222112223222010-1031010002130102-3301331110032301-2122202331201332) |
| `discovery_k8s.namespace_mapping.items.namespace_regex` | [discovery_k8s.namespace_mapping.items.namespace_regex](resources--discovery--reference--group-001.md#canonical-1130100022032112-2301113201110213-3223002220033003-0303132103013200-0001103312021020-0221103001202012-2132111201101330-3200113123101210) |
| `discovery_k8s.publish_info` | [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-3122222231121203-3012003210123330-0330311123322320-2213313323121323-0113123300203332-1113032311231020-0323112032313021-3223031122220121) |
| `discovery_k8s.publish_info.disable_spec` | [discovery_k8s.publish_info.disable_spec](resources--discovery--reference--group-001.md#canonical-1131112011113030-1210133212003130-3311313320312101-1010030313001132-3013331233013300-1312020132001211-2322023011103123-0311120302310100) |
| `discovery_k8s.publish_info.dns_delegation` | [discovery_k8s.publish_info.dns_delegation](resources--discovery--reference--group-001.md#canonical-0013022331002220-1313110232311321-3210232322133213-0021213003322120-1312002013122030-3330223222122301-3013202002112203-0303031222230100) |
| `discovery_k8s.publish_info.dns_delegation.dns_mode` | [discovery_k8s.publish_info.dns_delegation.dns_mode](resources--discovery--reference--group-001.md#canonical-1332300010111012-3302320013202330-0111130212331013-0130321000210330-3310023321020232-0010333012131112-1002231112021123-2320002323102121) |
| `discovery_k8s.publish_info.dns_delegation.subdomain` | [discovery_k8s.publish_info.dns_delegation.subdomain](resources--discovery--reference--group-001.md#canonical-1022211331330030-1332003123210002-3333231111322323-3323233032003011-2333223123033210-2202100211311321-3310101233321103-2333210311000310) |
| `discovery_k8s.publish_info.publish` | [discovery_k8s.publish_info.publish](resources--discovery--reference--group-001.md#canonical-1313310112011001-2301122021130021-2233122213001303-3132203303203013-0132212011210030-1110203023102312-2032012331000200-0031120213233032) |
| `discovery_k8s.publish_info.publish.namespace` | [discovery_k8s.publish_info.publish.namespace](resources--discovery--reference--group-001.md#canonical-2301110233320203-0021313232121333-1220131003100110-3221033032110200-2302200212010212-1113301212023221-0210121133120020-1213010011022120) |
| `discovery_k8s.publish_info.publish_fqdns` | [discovery_k8s.publish_info.publish_fqdns](resources--discovery--reference--group-001.md#canonical-1031022023310213-0232313203112320-1213302020313332-0311311033232210-2302323330201030-2113123312203013-2323321312021030-2222203323321201) |
| `id` | [ID](resources--discovery--reference--group-001.md#canonical-2331201330020310-3011330213123030-3203322003002111-0100133132101320-3300312013110232-3010033210133322-1221231133311301-0210330122333311) |
| `labels` | [labels](resources--discovery--reference--group-001.md#canonical-2220101102111200-2323310113332112-0132002200231312-1123011022011023-3200210221001203-1231003230231300-3321210213113200-3311121321002110) |
| `name` | [name](resources--discovery--reference--group-001.md#canonical-0231103223133302-1310022012112113-1303120002132223-1103111031233210-2302112001113222-2101033321321100-3033002001210122-1111331023213201) |
| `namespace` | [namespace](resources--discovery--reference--group-001.md#canonical-3031213032120321-0013232210133312-1333331030312323-1113322220213333-0100132022331333-1012202221103020-0130001103021021-0011233300302001) |
| `no_cluster_id` | [no_cluster_id](resources--discovery--reference--group-001.md#canonical-3002323030220322-0020200231302202-3111002300320220-3022101033030022-2021313300131030-0300131003010012-3302102103231001-3130001221002120) |
| `timeouts` | [timeouts](resources--discovery--reference--group-001.md#canonical-0012101032103333-1011032300121101-1320200130220320-3123221213101133-3033312102212232-3230131201211022-0330321023111232-2312201331323113) |
| `timeouts.create` | [timeouts.create](resources--discovery--reference--group-001.md#canonical-1112033131022101-2022013210222232-1032111211310023-3103132312130230-0332013311310200-1320023323331223-2123323113333110-2102223332202222) |
| `timeouts.delete` | [timeouts.delete](resources--discovery--reference--group-001.md#canonical-2202202032033022-2300320013310321-3032002000003122-3311310323233233-1110211012322012-2223330110301111-0021011112302010-3303323300133200) |
| `timeouts.read` | [timeouts.read](resources--discovery--reference--group-001.md#canonical-1011000201012302-1011300001301023-1101211332122121-0012132102332023-0132133332030322-2020132020301311-1330332300202211-1221120303031121) |
| `timeouts.update` | [timeouts.update](resources--discovery--reference--group-001.md#canonical-3231100130200331-3321310232220101-2113300013331202-2123333311110333-0103113330100211-0002013301300101-3333023313130001-0202100110313023) |
| `where` | [where](resources--discovery--reference--group-001.md#canonical-1321103302233201-0332121301121100-3101133031230333-0133220331333201-0030002311302101-3220122332233201-0213131123032010-1102233222333213) |
| `where.site` | [where.site](resources--discovery--reference--group-001.md#canonical-2012013100322022-1022120320103220-1213011202221212-0031301113211120-0121000101211123-3131302320112311-0011001103132122-2312321232311020) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](resources--discovery--reference--group-001.md#canonical-2130130203303330-3032002112302222-3303311021103102-3220300023331330-2220110300021020-3130203123022331-2211123013020023-2212023112133033) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](resources--discovery--reference--group-002.md#canonical-2013232001120033-0302203321031212-1133020011332302-1333012221313120-2000331022223130-1101021100312211-0222032331230000-0230320102112321) |
| `where.site.network_type` | [where.site.network_type](resources--discovery--reference--group-001.md#canonical-1223122302203100-3200331030230313-3033220330302303-2130131323322120-1023013321001130-2221000311303032-1010022102102223-2111232233001112) |
| `where.site.ref` | [where.site.ref](resources--discovery--reference--group-002.md#canonical-2033133220303202-3013313000222302-2301200020221123-1122313121332100-3130122201311311-0000112123213332-1303231310133313-1202031022130122) |
| `where.site.ref.kind` | [where.site.ref.kind](resources--discovery--reference--group-002.md#canonical-1113130301320303-0102201023122131-3030032223203001-1230013232303200-0011222101202102-0123132303213222-3202320010210311-3121021213212102) |
| `where.site.ref.name` | [where.site.ref.name](resources--discovery--reference--group-002.md#canonical-3031320213210302-3220122121323100-3210233133033321-1323022011022011-1200100211232201-0330332302310320-3033221000311312-1101131220220030) |
| `where.site.ref.namespace` | [where.site.ref.namespace](resources--discovery--reference--group-002.md#canonical-2200112313310013-2202311201323300-3333100133332133-0002320320303203-2123202220032331-2330031322020203-2012123023330220-3100130010330302) |
| `where.site.ref.tenant` | [where.site.ref.tenant](resources--discovery--reference--group-002.md#canonical-2330123230100021-3011022000123301-3010211231231122-2233001123111213-0111131030321332-2012031031200121-3220312312031010-3200110301112200) |
| `where.site.ref.uid` | [where.site.ref.uid](resources--discovery--reference--group-002.md#canonical-3302231323230211-1230022212113010-2132212230313210-1112211220021102-2201112103222210-2121122200113332-1232212331001002-3133130123122231) |
| `where.virtual_network` | [where.virtual_network](resources--discovery--reference--group-002.md#canonical-0232303211103312-3331210310230232-2233302213331130-0103210120311211-0000201320122332-3000300231333221-2002100223333131-2311110230300022) |
| `where.virtual_network.ref` | [where.virtual_network.ref](resources--discovery--reference--group-002.md#canonical-0313030022322130-0201012102111211-1323203223021023-3221032231313023-2111033001103001-1132232322001222-1111311222123012-0332300301000303) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](resources--discovery--reference--group-002.md#canonical-2310303223212320-0321110030100012-0212032031110030-1130130311202121-1230312032102123-2131033233032120-1233201003302012-2213110133030220) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](resources--discovery--reference--group-002.md#canonical-1101213121031030-1010311321101310-2001033232001130-0101113320121100-1011000021311202-0020323232033133-3100213301032230-3332132120111020) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](resources--discovery--reference--group-002.md#canonical-2230332102131303-2302330121300100-2031333221222002-3221200012111112-2203001021012013-2112222202300313-2100033221301220-3322312033320301) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](resources--discovery--reference--group-002.md#canonical-2303110203312302-2123030210132103-3330002012000123-0313233002221332-2100021101111202-2113323111313320-0311032233113213-0133322332013222) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](resources--discovery--reference--group-002.md#canonical-1330031302023030-1020030301031332-3120221321302000-2120021120311320-1033000332231101-1131022222011001-1133321202003102-1201031121313323) |
| `where.virtual_site` | [where.virtual_site](resources--discovery--reference--group-002.md#canonical-0322133000130021-2333312212222111-0020101213231023-2030310233221013-0020031031000102-3020113021303003-1120101000110032-0110002030131332) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](resources--discovery--reference--group-002.md#canonical-2312101123222101-2011011201100032-2221012332123301-3002130230322301-1332302020000333-1302213311003311-3230331320112310-2330320003131110) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](resources--discovery--reference--group-002.md#canonical-3021002022120123-1022010032122021-1110031333111120-2230322030033022-2320033033023333-2330031110211102-3331023012233013-1113232210103211) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](resources--discovery--reference--group-002.md#canonical-2000131222330332-3332323003332032-2221323000032331-0203112013033230-3122310103120230-2102103122211223-3222321131122303-0133011000320313) |
| `where.virtual_site.ref` | [where.virtual_site.ref](resources--discovery--reference--group-002.md#canonical-2223002323111232-1312332333333231-1100100132112221-0012333303020122-3102332001111121-1232303103331221-2230332023123001-3031110313130223) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](resources--discovery--reference--group-002.md#canonical-2310233223102032-3200231202102012-0332102020222202-2321213203231033-0131300210132221-2122310230200102-3213023230113112-2123333320120000) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](resources--discovery--reference--group-002.md#canonical-2321321030101011-1330320131023013-0333321210013022-1000011201113301-1300033212122000-3213323232321100-0210203301323011-0222000031001100) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](resources--discovery--reference--group-002.md#canonical-1310133233122030-3311122223101211-1231133000123210-1233003002231310-0232031212202101-0000111202002232-0013132332033210-2002201110021133) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](resources--discovery--reference--group-002.md#canonical-0130012332031021-3010032323230100-1022321222223111-3033330120201333-1313333323313330-0131100032320203-2032021301312030-0011231233211132) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](resources--discovery--reference--group-002.md#canonical-3211223011011211-1221223322021300-1303303103320033-1103022230301222-3320010130100231-0101312233120020-2030010012112132-1120233103310133) |

<a id="canonical-1232013301103023-1321332021321321-0123213111000033-2113303202033320-0003011302030203-3321331101300313-0133323000231001-2222221202230130"></a>

## Next pages — Property reference / 132033111333 / 13

- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [no_cluster_id](resources--discovery--reference--group-001.md#canonical-2332001313120022-1313323002112231-2021132023022331-0131012222332133-3231233201102020-0102032212311132-1210113121202130-1000203132133103)
- [timeouts](resources--discovery--reference--group-001.md#canonical-0133122212320132-3323032213033022-3031001120032311-3120311312312302-0221333302003213-0021323013310132-3013211230313033-1331213120121300)
- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323213102130102-2003331212101022-2333203233112221-0212333101311233-0211203323120130-1330320230001310-3320010110232003-3032030313220213"></a>

## discovery_consul — discovery_consul / 321123322023 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- discovery_consul

<a id="canonical-1030303101331321-3203302302330023-0021113131100321-2133033101120310-3100033000111203-0311120132222222-0133133213320000-3201332120011001"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: discovery\_consul, discovery\_k8s\] Discovery configuration for Hashicorp Consul.

Upstream description:

Discovery configuration for Hashicorp Consul.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-namespace_mapping_choice": "[]"
}
```

OneOf alternatives in this subsection:

- [discovery_consul](resources--discovery--reference--group-001.md#canonical-1030303101331321-3203302302330023-0021113131100321-2133033101120310-3100033000111203-0311120132222222-0133133213320000-3201332120011001)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-0303032133222231-3132132031112130-2102022332112001-0233203030233332-3011132230201022-1131031022000310-1002230130133033-1310212302021220)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
discovery_consul {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213001312001222-3313110323232112-2303033311130101-2021311023231022-0320332032322220-3013120120012310-0012122220222222-1111011022031110"></a>

## Direct properties — discovery_consul / 321123322023 / 3

- [access_info](resources--discovery--reference--group-001.md#canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232): complete subsection reference.

- [publish_info](resources--discovery--reference--group-001.md#canonical-3113220101333013-1213033110231132-1230122212322132-3100032133301212-2102011203013113-2321020301320012-3012220232023130-3302132230103012): complete subsection reference.

<a id="canonical-3222333012210311-0313023231133200-2203101100021313-3320102103230031-3231002103132000-3223112112033121-1301102021232211-1222021011010201"></a>

## Next pages — discovery_consul / 321123322023 / 4

- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232)
- [discovery_consul.publish_info](resources--discovery--reference--group-001.md#canonical-3113220101333013-1213033110231132-1230122212322132-3100032133301212-2102011203013113-2321020301320012-3012220232023130-3302132230103012)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133223300101111-1101112012100033-2023223023203203-1233021031023310-3203110002211110-0012320001003133-3330023212000300-3230232013312300"></a>

## discovery_consul.access_info — access_info / 131221233103 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- discovery_consul.access_info

<a id="canonical-3330303202231313-2123032003102222-3032212230121130-0301112230022000-2212130032202323-3110113201232220-3021233310013000-2301113132100311"></a>

Type: `"object"`. single nested block, Optional.

Hashicorp Consul API server information.

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
access_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013230231332200-0131103313302300-2203310320311200-0011000212031201-2120113133213112-1120220033212123-1022033120012333-2102310330020220"></a>

## Direct properties — access_info / 131221233103 / 3

- [connection_info](resources--discovery--reference--group-001.md#canonical-1021002230130101-2331301020313031-0203232222310221-2120111231021320-3112330001001112-0221300302102100-3333022202330012-2022101023021301): complete subsection reference.

- [http_basic_auth_info](resources--discovery--reference--group-001.md#canonical-0001133210323132-2031212030200231-2311331203333003-0221103130123101-1021100123320222-0102311230120010-1211202323201033-1230230013200100): complete subsection reference.

<a id="canonical-0001032331102301-2323333033203330-3120013012300133-0023331330321013-1021122123230303-2212210202212320-3130012030212202-3212102223020332"></a>

## Next pages — access_info / 131221233103 / 4

- [discovery_consul.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-1021002230130101-2331301020313031-0203232222310221-2120111231021320-3112330001001112-0221300302102100-3333022202330012-2022101023021301)
- [discovery_consul.access_info.http_basic_auth_info](resources--discovery--reference--group-001.md#canonical-0001133210323132-2031212030200231-2311331203333003-0221103130123101-1021100123320222-0102311230120010-1211202323201033-1230230013200100)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-1021002230130101-2331301020313031-0203232222310221-2120111231021320-3112330001001112-0221300302102100-3333022202330012-2022101023021301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002232000123021-2330031000221221-0231113333333300-3220020200010031-0001200312003131-2223223000131232-1303121110202101-2312310102101020"></a>

## discovery_consul.access_info.connection_info — connection_info / 000011322333 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232)
- discovery_consul.access_info.connection_info

<a id="canonical-2020132232200102-1310012132320031-1021131030332030-2013223213222213-3312101131332222-1301103131223123-2211233010320322-0311221010021010"></a>

Type: `"object"`. single nested block, Optional.

Configuration details to access discovery service REST API.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_server")}
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
connection_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333333121012333-3110221210313131-2030201320333033-3221102120010302-0332301211022230-0110302130311122-1103312103320022-2101223013001232"></a>

## Direct properties — connection_info / 000011322333 / 3

<a id="canonical-1133303222102323-0012213223111233-1113032003012133-2103311322122202-3132112121331001-3302222002301120-2231110213121300-0123032330103110"></a>

<a id="canonical-3103332300113211-1203123223121103-2200123131331300-2221203320111021-3112300300311133-0002002002230133-0310113222000212-0303312220331030"></a>

## api_server property — connection_info / 000011322333 / 4

Type: `"string"`. Optional.

API server must be a fully qualified domain string and port specified as host:port pair.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

- [tls_info](resources--discovery--reference--group-001.md#canonical-3312320212313202-1013300212310122-0312013131230323-0003132322330031-2321232011003102-3013022223003111-3330323320230303-1013032230130010): complete subsection reference.

<a id="canonical-2131230322231333-0103302121323300-3313211003223102-2322323030233211-1101020320001012-2313312010321332-0220233302132003-3200313313031010"></a>

## Next pages — connection_info / 000011322333 / 5

- [discovery_consul.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-3312320212313202-1013300212310122-0312013131230323-0003132322330031-2321232011003102-3013022223003111-3330323320230303-1013032230130010)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-3312320212313202-1013300212310122-0312013131230323-0003132322330031-2321232011003102-3013022223003111-3330323320230303-1013032230130010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003101223302113-1012131113133220-0330233030131111-2200020101011223-1100320210033000-2232003332303312-1232331220013321-0302303233220131"></a>

## discovery_consul.access_info.connection_info.tls_info — tls_info / 313123000122 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232)
- [discovery_consul.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-1021002230130101-2331301020313031-0203232222310221-2120111231021320-3112330001001112-0221300302102100-3333022202330012-2022101023021301)
- discovery_consul.access_info.connection_info.tls_info

<a id="canonical-3120130210213221-1310202020132212-0213200313211210-3122322021102032-2100231122310012-0120100032011203-2101213021321110-2333303010331310"></a>

Type: `"object"`. single nested block, Optional.

TLS config for client of discovery service.

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
tls_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3323212103001300-3301133311120220-3310221312001130-2121131113113320-3303320000123032-0011220321222311-2110212013302313-3203132303112120"></a>

## Direct properties — tls_info / 313123000122 / 3

<a id="canonical-2301230213323312-0221202303110310-3210322121200223-2213110100103000-0022012313121201-1132321100311310-2223003132033133-2001330203110021"></a>

<a id="canonical-0213101101330230-2031311131222012-3330332322203121-1020120322121232-0212232301203121-0331303022000212-3311213022220113-3333322331123200"></a>

## certificate property — tls_info / 313123000122 / 4

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(100, 131072),
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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](resources--discovery--reference--group-001.md#canonical-3312001233202222-0102220231102331-1220013030301102-0332020032020211-1002132302112303-2211310303103122-0320323122111200-2033023113201222): complete subsection reference.

<a id="canonical-2113211120130033-0131121333100022-1122032023010223-3223112000111332-3313121303113112-1120002202000223-3002302202010030-3222213120231000"></a>

<a id="canonical-1312002033101200-0323202003323330-3213000013100310-1222010133032230-0311222203033202-2010210120322102-3323123333333122-2323010233332211"></a>

## server_name property — tls_info / 313123000122 / 5

Type: `"string"`. Optional.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

Upstream description:

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

Provider validators and defaults (from schema source):

```go
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
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1212223230301001-0330020212032222-2202301102212312-0030220220033301-1300023323230301-2213030321203121-3200132003301130-2113321223033200"></a>

<a id="canonical-3000133201323230-1233022111212011-3010020222102103-2032101032033203-0111302100203031-0202222012112311-2213200113313213-2223110310212221"></a>

## trusted_ca_url property — tls_info / 313123000122 / 6

Type: `"string"`. Optional.

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

Upstream description:

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1022013332111332-1110130012001111-2310333132001310-3010330121212231-0131121110230323-2010231222313301-1331020322202011-0212220310203302"></a>

## Next pages — tls_info / 313123000122 / 7

- [discovery_consul.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-3312001233202222-0102220231102331-1220013030301102-0332020032020211-1002132302112303-2211310303103122-0320323122111200-2033023113201222)
- [discovery_consul.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-1021002230130101-2331301020313031-0203232222310221-2120111231021320-3112330001001112-0221300302102100-3333022202330012-2022101023021301)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-3312001233202222-0102220231102331-1220013030301102-0332020032020211-1002132302112303-2211310303103122-0320323122111200-2033023113201222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120212111332111-1231320100001221-0301213033212021-1003123331011001-2000331103301212-2232011110312020-1101301200112113-0033113321203113"></a>

## discovery_consul.access_info.connection_info.tls_info.key_url — key_url / 321312332013 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232)
- [discovery_consul.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-1021002230130101-2331301020313031-0203232222310221-2120111231021320-3112330001001112-0221300302102100-3333022202330012-2022101023021301)
- [discovery_consul.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-3312320212313202-1013300212310122-0312013131230323-0003132322330031-2321232011003102-3013022223003111-3330323320230303-1013032230130010)
- discovery_consul.access_info.connection_info.tls_info.key_url

<a id="canonical-0322032013110131-0131332133302010-3123000133213333-1333130200110113-1110002010120010-2020011120113302-0121123102310132-1313121333132320"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
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
key_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-1112130322013010-0312221212021211-3002231122020313-2330121333031210-2213222111223213-3120031200030110-1012333212132013-0331230233332013"></a>

## Direct properties — key_url / 321312332013 / 3

- [blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-3033312123211023-0313222123120020-1201011133331312-2300001302333112-0113220221322012-0003030320231231-3232113211001131-0022230313030122): complete subsection reference.

- [clear_secret_info](resources--discovery--reference--group-001.md#canonical-0012201201132201-2113310323100122-1032223311130201-1111200003201301-1212221232011100-0213123113301323-3302213002211023-0012003322201122): complete subsection reference.

<a id="canonical-0331221121011113-2113203320220003-0222111223132321-2230020220323323-2303220120003003-0113202000103203-1100001302120000-2110120103112020"></a>

## Next pages — key_url / 321312332013 / 4

- [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-3033312123211023-0313222123120020-1201011133331312-2300001302333112-0113220221322012-0003030320231231-3232113211001131-0022230313030122)
- [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-0012201201132201-2113310323100122-1032223311130201-1111200003201301-1212221232011100-0213123113301323-3302213002211023-0012003322201122)
- [discovery_consul.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-3312320212313202-1013300212310122-0312013131230323-0003132322330031-2321232011003102-3013022223003111-3330323320230303-1013032230130010)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-3033312123211023-0313222123120020-1201011133331312-2300001302333112-0113220221322012-0003030320231231-3232113211001131-0022230313030122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002333022123111-1223223023212212-2020013231301011-0223230100113000-1323321230012031-3012202301300231-1313322331123231-3222322131313330"></a>

## discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info — blindfold_secret_info / 302332201100 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232)
- [discovery_consul.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-1021002230130101-2331301020313031-0203232222310221-2120111231021320-3112330001001112-0221300302102100-3333022202330012-2022101023021301)
- [discovery_consul.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-3312320212313202-1013300212310122-0312013131230323-0003132322330031-2321232011003102-3013022223003111-3330323320230303-1013032230130010)
- [discovery_consul.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-3312001233202222-0102220231102331-1220013030301102-0332020032020211-1002132302112303-2211310303103122-0320323122111200-2033023113201222)
- discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info

<a id="canonical-2133230200133322-2000110300110011-1023130121032311-1100023030022303-0322201311102200-0213112201201231-1211131033111303-0301303111221222"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0322112120030110-2132030321231222-0022131303003320-3203212222331223-0200330012011222-1120132221132102-0022121203313323-0222221210211000"></a>

## Direct properties — blindfold_secret_info / 302332201100 / 3

<a id="canonical-3220220212302121-1332221330331101-3201012320231022-2110023003011013-2112321333123010-3111022003102332-3000000132303031-2101033021013012"></a>

<a id="canonical-2201011313303232-3030123320330201-0303302113122211-3020202311300123-3231021032312223-3330303133133023-0223222311111032-0312221303123111"></a>

## decryption_provider property — blindfold_secret_info / 302332201100 / 4

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

<a id="canonical-1322103030131202-2130131313200102-3323232230203010-0112231030100222-2112012201222011-1110212320101300-0213100003222030-0220313220023222"></a>

<a id="canonical-1013223110330120-1100020302110202-0013233000202300-0031320010223013-1120200001221002-2332321020033010-3230301103201013-1213123012231332"></a>

## location property — blindfold_secret_info / 302332201100 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1323130030102231-3333322333103321-3331133000021012-3231232013312032-1112010203110202-2230111312121132-1333211310313233-1212121023331123"></a>

<a id="canonical-3110321032203321-1322322221102130-2301002001011131-2213213101321030-1303003202113111-2030211010020301-3120100221100023-1021022112323202"></a>

## store_provider property — blindfold_secret_info / 302332201100 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-1122023300101003-3323030112220302-3003101220213130-0100200213230331-1310212013303203-3231212203020332-1221233101102302-0231203030300023"></a>

## Next pages — blindfold_secret_info / 302332201100 / 7

- [discovery_consul.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-3312001233202222-0102220231102331-1220013030301102-0332020032020211-1002132302112303-2211310303103122-0320323122111200-2033023113201222)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-0012201201132201-2113310323100122-1032223311130201-1111200003201301-1212221232011100-0213123113301323-3302213002211023-0012003322201122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311223013023323-0321103212223312-1302303211013233-1001121230021003-0223230131332221-1002133223333013-3330313023121211-1232320013023101"></a>

## discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info — clear_secret_info / 310103323232 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232)
- [discovery_consul.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-1021002230130101-2331301020313031-0203232222310221-2120111231021320-3112330001001112-0221300302102100-3333022202330012-2022101023021301)
- [discovery_consul.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-3312320212313202-1013300212310122-0312013131230323-0003132322330031-2321232011003102-3013022223003111-3330323320230303-1013032230130010)
- [discovery_consul.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-3312001233202222-0102220231102331-1220013030301102-0332020032020211-1002132302112303-2211310303103122-0320323122111200-2033023113201222)
- discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info

<a id="canonical-3313200120113232-2230301010132103-1313101013210010-0110210210300221-2233210010110322-0302301230020323-0320312100210112-0201130202311120"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0332321222312323-3311300011011131-3201233110323200-3332213003010311-3023313101032022-3203120122303300-0212200020001201-3232232331122033"></a>

## Direct properties — clear_secret_info / 310103323232 / 3

<a id="canonical-2210033212200030-3021323131010122-1130303131332011-1031311131323233-3011300302131103-0100122231122020-1320100120001221-1223023122223122"></a>

<a id="canonical-0232001222132221-0123013333332122-2122012111022201-0300033131130312-1201200230232010-2022110313001311-2312310010032323-3110230113103130"></a>

## provider_ref property — clear_secret_info / 310103323232 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2130222322301313-3002130111120031-3222311113002231-0203122013213003-0023032300333330-0113221133030331-2121231202110031-2233201213203113"></a>

<a id="canonical-3330011312120010-2202102110012011-2001132013030120-2021112202211033-1212221232002301-1032001033333021-2310222000213310-2210022130313002"></a>

## URL property — clear_secret_info / 310103323232 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0031222022210002-2210231322102331-3121313333102231-2133001302302031-1220320000131123-2001001131210122-1020232021122220-3223322033321311"></a>

## Next pages — clear_secret_info / 310103323232 / 6

- [discovery_consul.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-3312001233202222-0102220231102331-1220013030301102-0332020032020211-1002132302112303-2211310303103122-0320323122111200-2033023113201222)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-0001133210323132-2031212030200231-2311331203333003-0221103130123101-1021100123320222-0102311230120010-1211202323201033-1230230013200100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031020112211302-0031022112003123-3103112213211011-3313123302202021-2302212211130321-1031030231130003-0023231103233132-2222310002320120"></a>

## discovery_consul.access_info.http_basic_auth_info — http_basic_auth_info / 212330323331 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232)
- discovery_consul.access_info.http_basic_auth_info

<a id="canonical-0123313323210131-0230300022001113-3102310033300331-1101113223332333-2320232301123003-2010032111102013-3220221323113300-0123102111233213"></a>

Type: `"object"`. single nested block, Optional.

Authentication parameters to access Hashicorp Consul.

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
http_basic_auth_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003010231133013-3231301031320010-2011122222210023-0133131312303102-3211312213131320-2333211323113003-2033311131312201-3220301022030300"></a>

## Direct properties — http_basic_auth_info / 212330323331 / 3

- [passwd_url](resources--discovery--reference--group-001.md#canonical-1120100130133100-3333202313112000-2201131010110123-3312220211222320-3230132002122323-1232211201113033-1112032210023213-3230323120221220): complete subsection reference.

<a id="canonical-2132202320222323-1121113303312011-2112233233021003-2220210011122031-1212111313230010-1030200133101213-1110332030111032-2221020223100303"></a>

<a id="canonical-1323112033021312-2302330332311221-1113222130221300-1130120111303032-1233133222020033-0130020131232130-3311223123202202-0332131010313311"></a>

## user_name property — http_basic_auth_info / 212330323331 / 4

Type: `"string"`. Optional.

username. Username in consul.

Upstream description:

Username in consul.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
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

<a id="canonical-3303111332110333-1232230120010232-1232332300130203-2101321032103030-3311022300232003-3003100303103001-1311311212222321-3002022332303001"></a>

## Next pages — http_basic_auth_info / 212330323331 / 5

- [discovery_consul.access_info.http_basic_auth_info.passwd_url](resources--discovery--reference--group-001.md#canonical-1120100130133100-3333202313112000-2201131010110123-3312220211222320-3230132002122323-1232211201113033-1112032210023213-3230323120221220)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-1120100130133100-3333202313112000-2201131010110123-3312220211222320-3230132002122323-1232211201113033-1112032210023213-3230323120221220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120313012321023-3313110020033012-3220223020213130-3300301021023020-0110123020202221-0312211211122131-3031223312020021-0111023310231123"></a>

## discovery_consul.access_info.http_basic_auth_info.passwd_url — passwd_url / 103230230122 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232)
- [discovery_consul.access_info.http_basic_auth_info](resources--discovery--reference--group-001.md#canonical-0001133210323132-2031212030200231-2311331203333003-0221103130123101-1021100123320222-0102311230120010-1211202323201033-1230230013200100)
- discovery_consul.access_info.http_basic_auth_info.passwd_url

<a id="canonical-1232301021111123-2202332222031210-0200101321230103-1222032233121033-3210130233332030-0103323203220232-3020133322103003-0003321010333010"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
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
passwd_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031222011032213-3213103121222122-1232033220302223-1303230221231202-2200022003200031-0210313330021311-3021310212133333-3323212333111003"></a>

## Direct properties — passwd_url / 103230230122 / 3

- [blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-3011232101103130-2100323200110301-0332013100002313-2232130012031033-2132232233222120-0130212231002000-2311000313002301-3010312133231230): complete subsection reference.

- [clear_secret_info](resources--discovery--reference--group-001.md#canonical-3201033321323120-1203211203210122-1333133303100022-2331320320133023-0233000311233122-3131001020320321-0101223213221310-3131110011322031): complete subsection reference.

<a id="canonical-1111112333113030-0332211321113203-1313011102203222-1232002132303000-3210131220303113-1321101111210202-1213321000222201-3310013231113303"></a>

## Next pages — passwd_url / 103230230122 / 4

- [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-3011232101103130-2100323200110301-0332013100002313-2232130012031033-2132232233222120-0130212231002000-2311000313002301-3010312133231230)
- [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-3201033321323120-1203211203210122-1333133303100022-2331320320133023-0233000311233122-3131001020320321-0101223213221310-3131110011322031)
- [discovery_consul.access_info.http_basic_auth_info](resources--discovery--reference--group-001.md#canonical-0001133210323132-2031212030200231-2311331203333003-0221103130123101-1021100123320222-0102311230120010-1211202323201033-1230230013200100)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-3011232101103130-2100323200110301-0332013100002313-2232130012031033-2132232233222120-0130212231002000-2311000313002301-3010312133231230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223220331120302-3101220120312100-2013101031021113-3032210321223021-0100302322121222-1233311112111103-2322210131120223-1300330021011013"></a>

## discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info — blindfold_secret_info / 120302023132 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232)
- [discovery_consul.access_info.http_basic_auth_info](resources--discovery--reference--group-001.md#canonical-0001133210323132-2031212030200231-2311331203333003-0221103130123101-1021100123320222-0102311230120010-1211202323201033-1230230013200100)
- [discovery_consul.access_info.http_basic_auth_info.passwd_url](resources--discovery--reference--group-001.md#canonical-1120100130133100-3333202313112000-2201131010110123-3312220211222320-3230132002122323-1232211201113033-1112032210023213-3230323120221220)
- discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info

<a id="canonical-2011002002332312-0112022132233001-3223200101312203-3223010302220121-2210133233100231-2303132002220213-0112302002221111-3233120133123102"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2131221321120000-3021000223210302-0032020213302233-0030100301200033-3131100131323011-1110311222223210-3001312031032320-3103202120303031"></a>

## Direct properties — blindfold_secret_info / 120302023132 / 3

<a id="canonical-0132013020121313-0330012133100232-1330210312011222-0102213222102200-2121110232320132-3032312003211003-3201033201011203-0113130203013313"></a>

<a id="canonical-3213133112230233-0031303311230000-3330310133210123-0020222230331223-0111223112110003-2030001000102013-0300010202211323-3010200032013110"></a>

## decryption_provider property — blindfold_secret_info / 120302023132 / 4

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

<a id="canonical-1231020032012213-2131223322102112-1032111330313020-1201331102010213-2133311220301300-3133131003331101-0033322121121223-1300223300010110"></a>

<a id="canonical-3110322312231330-1323210033231011-3033133302212302-0112331212130333-2111103211213332-3222321031310032-0101033212101323-2111230000201000"></a>

## location property — blindfold_secret_info / 120302023132 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0232022130001320-1321112000032222-2001123003103002-2322031203323323-3303003000113310-2121220013110123-2210231223011012-0201012230200321"></a>

<a id="canonical-2121233022201322-1122322013232020-0103312323220100-2111312022102031-2130212022113332-2201132013313201-3123122223212030-0311120232031121"></a>

## store_provider property — blindfold_secret_info / 120302023132 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-2033201320332020-2011033000330201-2112101022321322-2022031102100331-2010112221102221-3021011211222122-2321202302333300-2200231213122133"></a>

## Next pages — blindfold_secret_info / 120302023132 / 7

- [discovery_consul.access_info.http_basic_auth_info.passwd_url](resources--discovery--reference--group-001.md#canonical-1120100130133100-3333202313112000-2201131010110123-3312220211222320-3230132002122323-1232211201113033-1112032210023213-3230323120221220)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-3201033321323120-1203211203210122-1333133303100022-2331320320133023-0233000311233122-3131001020320321-0101223213221310-3131110011322031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233232010212111-2003101232320111-1031110022331021-0123103201320330-1211120031103231-3220133210222012-3212002022213201-3222101330302332"></a>

## discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info — clear_secret_info / 021112111031 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- [discovery_consul.access_info](resources--discovery--reference--group-001.md#canonical-2332031222030011-1013001033031322-3220012120232003-1031221221012300-2223030213002312-1333223111012301-0123302220223011-2032212310123232)
- [discovery_consul.access_info.http_basic_auth_info](resources--discovery--reference--group-001.md#canonical-0001133210323132-2031212030200231-2311331203333003-0221103130123101-1021100123320222-0102311230120010-1211202323201033-1230230013200100)
- [discovery_consul.access_info.http_basic_auth_info.passwd_url](resources--discovery--reference--group-001.md#canonical-1120100130133100-3333202313112000-2201131010110123-3312220211222320-3230132002122323-1232211201113033-1112032210023213-3230323120221220)
- discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info

<a id="canonical-3020110303003233-3100031300030231-2231220032300221-1333031203003321-3103232131210110-0310200100122323-0300331232200213-0131223120013103"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2201300020201202-2302320133021313-0133303211111002-3313100222332122-3100121221233323-0003223312330202-2122303200303023-1000310332310130"></a>

## Direct properties — clear_secret_info / 021112111031 / 3

<a id="canonical-0201021033121121-2001022100022323-1122233001201200-1032100122210300-2332022233110131-0122213333032233-3103131303121320-2113032233320032"></a>

<a id="canonical-3231333030322211-3303000330001301-0322323321100020-3133211200120302-0233330232032020-0030033120113203-3302321212323223-2320033210222131"></a>

## provider_ref property — clear_secret_info / 021112111031 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3331212301021302-1123121333002311-0231310222131210-2201310223011013-0333222010320311-0232333021021100-2001032330221310-0120130000211310"></a>

<a id="canonical-2133131220021031-1231032123120323-1332000102023203-3132133132103022-0033223003322013-0303322301130232-1023313232332203-3131121312212132"></a>

## URL property — clear_secret_info / 021112111031 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2223122131312223-2030033120101322-0333231111113331-3301222331003033-2223231132112023-3302102313133111-3332122233311210-1201110030111123"></a>

## Next pages — clear_secret_info / 021112111031 / 6

- [discovery_consul.access_info.http_basic_auth_info.passwd_url](resources--discovery--reference--group-001.md#canonical-1120100130133100-3333202313112000-2201131010110123-3312220211222320-3230132002122323-1232211201113033-1112032210023213-3230323120221220)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-3113220101333013-1213033110231132-1230122212322132-3100032133301212-2102011203013113-2321020301320012-3012220232023130-3302132230103012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122221022100313-1130311221111032-1103010022021321-1031231213001031-0122022123011311-2023332210102330-0303131323211100-1201101022001320"></a>

## discovery_consul.publish_info — publish_info / 232230001013 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- discovery_consul.publish_info

<a id="canonical-0112121310302232-1031213213023031-1232201113331111-3221201010312220-0013300131203133-3300221321011323-0321033123022213-2330323000232211"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for publish info.

Upstream description:

Consul Configuration to publish VIPs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "publish")}
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
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"publish\"]"
}
```

Terraform syntax:

```terraform
publish_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000033113331120-3113332311323033-2122201331301101-3221001121213132-0222330221131000-3010102131121201-2303320111212013-2020133011120312"></a>

## Direct properties — publish_info / 232230001013 / 3

- [disable_spec](resources--discovery--reference--group-001.md#canonical-0222333221022122-1122123320000221-2023033223002203-3131103021112113-2232212130211311-2212120222322211-0132313102110001-1222130133221302): complete subsection reference.

- [publish](resources--discovery--reference--group-001.md#canonical-2132011113020320-3201332222122110-2333032232133130-1012112010113222-2000133120330330-3023220101123113-0012230332222123-3223033230302310): complete subsection reference.

<a id="canonical-0101331210222123-3123313203032101-3302332032010020-0233303202211030-0231020001122303-0233223023321310-3033120111231232-1030131103333320"></a>

## Next pages — publish_info / 232230001013 / 4

- [discovery_consul.publish_info.disable_spec](resources--discovery--reference--group-001.md#canonical-0222333221022122-1122123320000221-2023033223002203-3131103021112113-2232212130211311-2212120222322211-0132313102110001-1222130133221302)
- [discovery_consul.publish_info.publish](resources--discovery--reference--group-001.md#canonical-2132011113020320-3201332222122110-2333032232133130-1012112010113222-2000133120330330-3023220101123113-0012230332222123-3223033230302310)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-0222333221022122-1122123320000221-2023033223002203-3131103021112113-2232212130211311-2212120222322211-0132313102110001-1222130133221302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300020011302302-1132112310311132-3333231100023113-3233220132313320-3022230211223313-0120113003203100-3112300233112102-1321130300212230"></a>

## discovery_consul.publish_info.disable_spec — disable_spec / 003230221111 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- [discovery_consul.publish_info](resources--discovery--reference--group-001.md#canonical-3113220101333013-1213033110231132-1230122212322132-3100032133301212-2102011203013113-2321020301320012-3012220232023130-3302132230103012)
- discovery_consul.publish_info.disable_spec

<a id="canonical-3221112130031131-0211203101310202-1330101002321330-1312123122202003-0233002322031110-1122021123332010-3031230031110032-1323321212000121"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-3322113222103302-1110300330300001-1332303302001311-0122131333202001-1232312013311000-0020203302011303-1211002122233111-2300011320212133"></a>

## Direct properties — disable_spec / 003230221111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210023233313113-1333011122211102-0121212111302303-2203332033112323-1312133312012010-1213200211323233-3300222301211021-3113200021023112"></a>

## Next pages — disable_spec / 003230221111 / 4

- [discovery_consul.publish_info](resources--discovery--reference--group-001.md#canonical-3113220101333013-1213033110231132-1230122212322132-3100032133301212-2102011203013113-2321020301320012-3012220232023130-3302132230103012)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2132011113020320-3201332222122110-2333032232133130-1012112010113222-2000133120330330-3023220101123113-0012230332222123-3223033230302310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121331312020313-3102230031311323-2032311010203300-3202032233303303-2111332322312330-3202112230202011-0321001003301102-1301311102203103"></a>

## discovery_consul.publish_info.publish — publish / 323002130333 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_consul](resources--discovery--reference--group-001.md#canonical-3002310222320131-3033001100122133-1233311322020320-0021101131131130-2300331220022023-0003100131221220-1002321100123201-2133101310111233)
- [discovery_consul.publish_info](resources--discovery--reference--group-001.md#canonical-3113220101333013-1213033110231132-1230122212322132-3100032133301212-2102011203013113-2321020301320012-3012220232023130-3302132230103012)
- discovery_consul.publish_info.publish

<a id="canonical-1201030231033011-1122023123320031-0321020002313233-1121213310311302-2132133110023212-2310332013221222-3202132223101021-1231032201231331"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
publish = {}
```

<a id="canonical-0110211232213330-0231022221012211-1021133033322310-2222100220032220-0033022312201021-0222002310123131-1031122033213000-1011002103100311"></a>

## Direct properties — publish / 323002130333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312223231303310-2101113201232123-0322220022033210-2203331211131331-1102201333102311-0100011130232013-0001233212222133-1003133030221101"></a>

## Next pages — publish / 323002130333 / 4

- [discovery_consul.publish_info](resources--discovery--reference--group-001.md#canonical-3113220101333013-1213033110231132-1230122212322132-3100032133301212-2102011203013113-2321020301320012-3012220232023130-3302132230103012)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222102203231110-2333301230200113-2232323132102333-2301210233132303-1201022121120030-0322200333212313-3012200210101222-1221222012210203"></a>

## discovery_k8s — discovery_k8s / 010011011300 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- discovery_k8s

<a id="canonical-0303032133222231-3132132031112130-2102022332112001-0233203030233332-3011132230201022-1131031022000310-1002230130133033-1310212302021220"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for discovery k8s.

Upstream description:

Discovery configuration for K8s.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_all",
    "namespace_mapping")}
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
  "x-ves-oneof-field-namespace_mapping_choice": "[\"default_all\",\"namespace_mapping\"]"
}
```

Terraform syntax:

```terraform
discovery_k8s {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320303033112002-1013203232101003-0100201220101200-1001101331320332-0300013011320113-0200331313232322-0123120113221032-3200112002032303"></a>

## Direct properties — discovery_k8s / 010011011300 / 3

- [access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320): complete subsection reference.

- [default_all](resources--discovery--reference--group-001.md#canonical-2112000203012301-3300021232321103-2332023112322323-2303202002010111-0021011211110313-0120122221203123-0121321110201012-0332133223123202): complete subsection reference.

- [namespace_mapping](resources--discovery--reference--group-001.md#canonical-0321010022030131-3311200301332101-3132022310131230-1233220021130131-3210333300201303-3321221200212203-2130012322120030-1013021110331232): complete subsection reference.

- [publish_info](resources--discovery--reference--group-001.md#canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113): complete subsection reference.

<a id="canonical-2111233303022213-1030022231000303-3323321203310000-2112111231332233-1022210321310133-1113120231233222-2322310102222130-1302320313210321"></a>

## Next pages — discovery_k8s / 010011011300 / 4

- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- [discovery_k8s.default_all](resources--discovery--reference--group-001.md#canonical-2112000203012301-3300021232321103-2332023112322323-2303202002010111-0021011211110313-0120122221203123-0121321110201012-0332133223123202)
- [discovery_k8s.namespace_mapping](resources--discovery--reference--group-001.md#canonical-0321010022030131-3311200301332101-3132022310131230-1233220021130131-3210333300201303-3321221200212203-2130012322120030-1013021110331232)
- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120203131112033-0101300223212220-3103231110033123-3332020222120022-0223310030000021-3203131300033301-3331211332001302-1223232332000210"></a>

## discovery_k8s.access_info — access_info / 000311100012 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- discovery_k8s.access_info

<a id="canonical-0130013031001132-2223032223121222-1313121231220323-2300100012211233-1002101110133103-2212320113310231-0310211110122230-1021232103100020"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for access info.

Upstream description:

K8s API server access.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("connection_info",
    "kubeconfig_url"),
  validators.ConflictingObjectAttributes("isolated",
    "reachable")}
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
  "x-ves-oneof-field-config_type": "[\"connection_info\",\"kubeconfig_url\"]",
  "x-ves-oneof-field-k8s_pod_network_choice": "[\"isolated\",\"reachable\"]"
}
```

Terraform syntax:

```terraform
access_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323121220120221-2003302212020233-2111302023010123-3010102220212130-0201222322033300-2211213222310201-0020232311223323-3030231312233121"></a>

## Direct properties — access_info / 000311100012 / 3

- [connection_info](resources--discovery--reference--group-001.md#canonical-3102311222101313-3002123010110023-0001213313233230-1111231130222303-3323223313013213-1003311301200300-2031022333113221-0021113332233333): complete subsection reference.

- [isolated](resources--discovery--reference--group-001.md#canonical-3332310112212022-3122012220002330-2121220211123313-1033230023102333-3330312121222111-2032023323130013-0003110121320123-2320322123230011): complete subsection reference.

- [kubeconfig_url](resources--discovery--reference--group-001.md#canonical-2103233030101201-2311203221133210-0220202230223032-3211312010223020-0101303231022121-2330110001021033-1320230101103202-0211001030103130): complete subsection reference.

- [reachable](resources--discovery--reference--group-001.md#canonical-2233200231201012-1110023212011100-0220223121300122-2221312123303311-3102232213022203-2332332222223031-3301300320131201-2133101011010213): complete subsection reference.

<a id="canonical-2201221301121020-1123230130113200-2133302301332113-0113322113111303-2220030130320211-3330213032332022-3122330323212120-0133031232021302"></a>

## Next pages — access_info / 000311100012 / 4

- [discovery_k8s.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-3102311222101313-3002123010110023-0001213313233230-1111231130222303-3323223313013213-1003311301200300-2031022333113221-0021113332233333)
- [discovery_k8s.access_info.isolated](resources--discovery--reference--group-001.md#canonical-3332310112212022-3122012220002330-2121220211123313-1033230023102333-3330312121222111-2032023323130013-0003110121320123-2320322123230011)
- [discovery_k8s.access_info.kubeconfig_url](resources--discovery--reference--group-001.md#canonical-2103233030101201-2311203221133210-0220202230223032-3211312010223020-0101303231022121-2330110001021033-1320230101103202-0211001030103130)
- [discovery_k8s.access_info.reachable](resources--discovery--reference--group-001.md#canonical-2233200231201012-1110023212011100-0220223121300122-2221312123303311-3102232213022203-2332332222223031-3301300320131201-2133101011010213)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-3102311222101313-3002123010110023-0001213313233230-1111231130222303-3323223313013213-1003311301200300-2031022333113221-0021113332233333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003020332033302-3332101332212021-1023103213121132-3220022210213121-1000211232311301-1212302300232202-1113211322323322-1221011101310131"></a>

## discovery_k8s.access_info.connection_info — connection_info / 201020231032 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- discovery_k8s.access_info.connection_info

<a id="canonical-3011220101012032-1302231313202322-2332330323223100-3321211321321131-1113032201322313-2033223123031221-0123210020223332-0003211332111332"></a>

Type: `"object"`. single nested block, Optional.

Configuration details to access discovery service REST API.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_server")}
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
connection_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021032320332100-2320321001112010-1331310113132030-2322101121120000-1312022213330020-3132210331213203-1010021323012202-0112111123013310"></a>

## Direct properties — connection_info / 201020231032 / 3

<a id="canonical-0200221101102201-3303000322031220-3232032223310002-3030013002320013-0312303111333222-1232212030023222-3310021101303120-0030021132033230"></a>

<a id="canonical-3021300133230300-0320100110303003-3021313331332211-3322222131310221-3130300023030313-2203121131223003-0331221302211223-1103310223211332"></a>

## api_server property — connection_info / 201020231032 / 4

Type: `"string"`. Optional.

API server must be a fully qualified domain string and port specified as host:port pair.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

- [tls_info](resources--discovery--reference--group-001.md#canonical-0030001103023010-3113301100121002-0331223212330233-0002303323120111-3030202130321000-2210212100221313-3220133122323211-1100113002001310): complete subsection reference.

<a id="canonical-0321213010032133-1123300320003032-0101002231131023-0312111223121313-0123322221222023-3123112101001112-2220211120020322-3310333130302202"></a>

## Next pages — connection_info / 201020231032 / 5

- [discovery_k8s.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-0030001103023010-3113301100121002-0331223212330233-0002303323120111-3030202130321000-2210212100221313-3220133122323211-1100113002001310)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-0030001103023010-3113301100121002-0331223212330233-0002303323120111-3030202130321000-2210212100221313-3220133122323211-1100113002001310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023321022213122-1323221113313133-2320132012212001-3310102022031312-3212032003023023-2122001331013210-3302130201210323-1322100211113110"></a>

## discovery_k8s.access_info.connection_info.tls_info — tls_info / 330102023030 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- [discovery_k8s.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-3102311222101313-3002123010110023-0001213313233230-1111231130222303-3323223313013213-1003311301200300-2031022333113221-0021113332233333)
- discovery_k8s.access_info.connection_info.tls_info

<a id="canonical-3030022132102000-2031102201323102-3113100222010330-3300213113203031-1101233123002001-3010120103132320-0020103003230310-1103311223122223"></a>

Type: `"object"`. single nested block, Optional.

TLS config for client of discovery service.

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
tls_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032300211033302-0230111333303231-2013100013112023-3031020220313112-1030321211111333-1330223011233032-2003310313330013-0103312332332311"></a>

## Direct properties — tls_info / 330102023030 / 3

<a id="canonical-3110323021212022-1000313012233123-3013103311303101-0022001131132121-0133110020322001-1021330121111203-2012232320222023-2330210010101202"></a>

<a id="canonical-2132223202212022-3133123102321333-0321301001020020-0113310133213130-2221233001311031-0102100133303322-0021101003333131-2323301220220023"></a>

## certificate property — tls_info / 330102023030 / 4

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(100, 131072),
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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](resources--discovery--reference--group-001.md#canonical-2321032223321012-0011230032331000-3111030313111103-0032322020022302-2223322130213223-2122210113321310-3032121331333211-0013201232200111): complete subsection reference.

<a id="canonical-0011013301122113-0011002320213102-3010301133220331-0232231320123020-0200222132013102-3021022010301003-1212013131210111-1103232123211313"></a>

<a id="canonical-1122321333230212-0331031011333320-2010201201100003-2322110320302323-2120312130223111-3010120310230011-0032021020032222-0110322002001023"></a>

## server_name property — tls_info / 330102023030 / 5

Type: `"string"`. Optional.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

Upstream description:

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

Provider validators and defaults (from schema source):

```go
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
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0100313202032302-1130311202120020-3110313022132332-2011101001301232-1231232222133121-0003233221333223-0011333231121330-1003210213133130"></a>

<a id="canonical-2101122121113232-1021330011232011-3103112333100103-3320330230130321-2021223230300012-1133131202233110-3322023230020220-1220111223011231"></a>

## trusted_ca_url property — tls_info / 330102023030 / 6

Type: `"string"`. Optional.

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

Upstream description:

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2201031103331021-1331301122300302-0322031233210031-0301020231320123-3001303203322333-3031122121021020-0320313123101032-3313100300002201"></a>

## Next pages — tls_info / 330102023030 / 7

- [discovery_k8s.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-2321032223321012-0011230032331000-3111030313111103-0032322020022302-2223322130213223-2122210113321310-3032121331333211-0013201232200111)
- [discovery_k8s.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-3102311222101313-3002123010110023-0001213313233230-1111231130222303-3323223313013213-1003311301200300-2031022333113221-0021113332233333)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2321032223321012-0011230032331000-3111030313111103-0032322020022302-2223322130213223-2122210113321310-3032121331333211-0013201232200111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201321213033003-0322201032133223-2331322102133131-2332323022311121-3320012000300100-0232100001102132-1112112200000301-0333101020023121"></a>

## discovery_k8s.access_info.connection_info.tls_info.key_url — key_url / 213230331300 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- [discovery_k8s.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-3102311222101313-3002123010110023-0001213313233230-1111231130222303-3323223313013213-1003311301200300-2031022333113221-0021113332233333)
- [discovery_k8s.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-0030001103023010-3113301100121002-0331223212330233-0002303323120111-3030202130321000-2210212100221313-3220133122323211-1100113002001310)
- discovery_k8s.access_info.connection_info.tls_info.key_url

<a id="canonical-2220020133132011-0310001033322331-2233100020320200-0103130122203010-3230030213012100-3333220232330233-2133233331013120-2003202131013020"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
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
key_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221213121231230-2300121020222233-2331201100000301-1010332323131203-2322001023312302-2113201310100221-0332122022003222-1033301013000312"></a>

## Direct properties — key_url / 213230331300 / 3

- [blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-2113310213120031-2301332323203120-3122221113212131-2121100101323132-2311103230222101-3200130200100322-3323313211220022-1130030010011210): complete subsection reference.

- [clear_secret_info](resources--discovery--reference--group-001.md#canonical-3120331100301131-1103201232213112-2101021323213213-3311300303012310-0231313111011120-2230200101110230-0022110203002200-3132302012232302): complete subsection reference.

<a id="canonical-1101111332333213-0323112332200333-0122131132221131-3331223011230103-2113123023032020-3200220201023012-3212132131201121-3113102211112123"></a>

## Next pages — key_url / 213230331300 / 4

- [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-2113310213120031-2301332323203120-3122221113212131-2121100101323132-2311103230222101-3200130200100322-3323313211220022-1130030010011210)
- [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-3120331100301131-1103201232213112-2101021323213213-3311300303012310-0231313111011120-2230200101110230-0022110203002200-3132302012232302)
- [discovery_k8s.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-0030001103023010-3113301100121002-0331223212330233-0002303323120111-3030202130321000-2210212100221313-3220133122323211-1100113002001310)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2113310213120031-2301332323203120-3122221113212131-2121100101323132-2311103230222101-3200130200100322-3323313211220022-1130030010011210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120203332011031-2212112233010331-1031013332320232-0032312131311013-1033000102331230-2010021013022012-0023301121333310-0022102122001132"></a>

## discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info — blindfold_secret_info / 112000000002 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- [discovery_k8s.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-3102311222101313-3002123010110023-0001213313233230-1111231130222303-3323223313013213-1003311301200300-2031022333113221-0021113332233333)
- [discovery_k8s.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-0030001103023010-3113301100121002-0331223212330233-0002303323120111-3030202130321000-2210212100221313-3220133122323211-1100113002001310)
- [discovery_k8s.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-2321032223321012-0011230032331000-3111030313111103-0032322020022302-2223322130213223-2122210113321310-3032121331333211-0013201232200111)
- discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info

<a id="canonical-3130131223001003-1123001123132131-1313231121301233-2012103101020121-0100203311321032-1032132333233322-3323301030122332-2022231021202003"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0200303031120121-3032133103130321-3113322333132222-3002331001232200-1212132330022202-0032302221321333-1100221013233133-3303022322212231"></a>

## Direct properties — blindfold_secret_info / 112000000002 / 3

<a id="canonical-0312230222100310-1021131022101323-3321222222113132-0100021301102121-2002323132202102-1301132331000323-2231132100030301-2120121230021032"></a>

<a id="canonical-0311331112002111-0012020212102213-1010223331031012-2111011321323122-1202223203201312-3110230322211211-0201332113102032-2030112320002210"></a>

## decryption_provider property — blindfold_secret_info / 112000000002 / 4

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

<a id="canonical-1221323132122311-3130132201123331-2013320121201032-1113020022123320-2021022223300311-2211211113213212-1003230013130002-3200123230121313"></a>

<a id="canonical-1212313333213301-2301001331022322-3032110131011003-3102012223032221-1011232113010220-3000300202013302-0203222002123103-0331330333200220"></a>

## location property — blindfold_secret_info / 112000000002 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2330021003221202-3300032202331212-1111131000233330-0121121313320222-3111332000033313-3322332212223201-1021032332313221-2230100330021202"></a>

<a id="canonical-3203030033131121-1030021010220102-2303032113212302-1030103023132313-0212223012322203-3201202100301232-0030313303300012-2111121103000311"></a>

## store_provider property — blindfold_secret_info / 112000000002 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-0000231000132323-0220210311012201-3330322312203103-2330023221320220-2222111202203033-3112012322110313-3200211011003122-0231110212212022"></a>

## Next pages — blindfold_secret_info / 112000000002 / 7

- [discovery_k8s.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-2321032223321012-0011230032331000-3111030313111103-0032322020022302-2223322130213223-2122210113321310-3032121331333211-0013201232200111)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-3120331100301131-1103201232213112-2101021323213213-3311300303012310-0231313111011120-2230200101110230-0022110203002200-3132302012232302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332201011030320-1221201300303232-0001011132211033-1323201122103213-3223310112122320-1012112210131322-3001333013023123-3022031233213211"></a>

## discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info — clear_secret_info / 030203033011 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- [discovery_k8s.access_info.connection_info](resources--discovery--reference--group-001.md#canonical-3102311222101313-3002123010110023-0001213313233230-1111231130222303-3323223313013213-1003311301200300-2031022333113221-0021113332233333)
- [discovery_k8s.access_info.connection_info.tls_info](resources--discovery--reference--group-001.md#canonical-0030001103023010-3113301100121002-0331223212330233-0002303323120111-3030202130321000-2210212100221313-3220133122323211-1100113002001310)
- [discovery_k8s.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-2321032223321012-0011230032331000-3111030313111103-0032322020022302-2223322130213223-2122210113321310-3032121331333211-0013201232200111)
- discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info

<a id="canonical-0121112223310011-2200111031121311-0213131133330232-2101300310013320-3102130120120332-3203020232330101-3100212001023302-0013331132323113"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1321223111211332-1033213110013201-2200012211203112-1311221313212110-0013022332113121-3311001011221110-2332201303022223-3220101120320132"></a>

## Direct properties — clear_secret_info / 030203033011 / 3

<a id="canonical-2331322022300123-2310321011223113-3131200021312131-0330011303211121-2121130130311123-0210310122201110-1310033203232322-3310111311233032"></a>

<a id="canonical-1302302123111302-1212221113202011-2231312330331010-3110031321211222-3202030202321330-3123031223100003-1130330300202301-1011021231231032"></a>

## provider_ref property — clear_secret_info / 030203033011 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2301132332001300-2200123320212000-0233323102223302-3122121032003130-1033300121330231-3003323030133210-2320031031121303-3103021313103101"></a>

<a id="canonical-1222333002203313-2031212102323231-3011012332003011-2222221020102123-0323110113122122-2212211333312330-1311213011223122-2313113233010333"></a>

## URL property — clear_secret_info / 030203033011 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2212200211000113-1030303323120203-2111123212012133-0131100211210123-3002122303231233-1021123200223330-0013030300312023-3333130111013011"></a>

## Next pages — clear_secret_info / 030203033011 / 6

- [discovery_k8s.access_info.connection_info.tls_info.key_url](resources--discovery--reference--group-001.md#canonical-2321032223321012-0011230032331000-3111030313111103-0032322020022302-2223322130213223-2122210113321310-3032121331333211-0013201232200111)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-3332310112212022-3122012220002330-2121220211123313-1033230023102333-3330312121222111-2032023323130013-0003110121320123-2320322123230011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312133121313333-3000213200021203-1112022111321130-2333111311112123-1132221001333123-2213310312331132-1203103012201011-1232002200113100"></a>

## discovery_k8s.access_info.isolated — isolated / 210123003332 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- discovery_k8s.access_info.isolated

<a id="canonical-3211300221101310-3111000233302102-0003303033023023-0013221003331123-1010310010222310-2323120000222312-2003211221021101-3103201311333023"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
isolated = {}
```

<a id="canonical-1101312022020023-2223120221320221-0003211322022121-1212133212222333-2312011111110323-3013320313030113-0103331000311322-3211320331323113"></a>

## Direct properties — isolated / 210123003332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012023022121012-2222202302131002-1231232103011100-0113233102310101-3310302200102320-3311101331332121-2123202012220001-0231102330031232"></a>

## Next pages — isolated / 210123003332 / 4

- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2103233030101201-2311203221133210-0220202230223032-3211312010223020-0101303231022121-2330110001021033-1320230101103202-0211001030103130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103200102313011-2011113110132231-2032022300133100-3211020220310133-1110330313310210-3322130202103322-2002000323212203-3202211000100120"></a>

## discovery_k8s.access_info.kubeconfig_url — kubeconfig_url / 213110113303 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- discovery_k8s.access_info.kubeconfig_url

<a id="canonical-1033220201210013-0311033123222002-3212333323131300-0213032312230030-0101322010020222-0302103230121203-2033313012320120-3023033010100033"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
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
kubeconfig_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231333032201001-1320332302001211-3003313112332130-1001332232110133-1213231313031300-3330312033211330-1130220300103133-2232301203121132"></a>

## Direct properties — kubeconfig_url / 213110113303 / 3

- [blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-1121112003223220-2111211201312211-3313302223212113-0312231220131131-0222122203020203-0231232101302010-1303022030200300-3333302311300130): complete subsection reference.

- [clear_secret_info](resources--discovery--reference--group-001.md#canonical-0100213213212331-2003313330013022-3230232303212010-2201001131100021-2221103310311031-2131103302233031-0300232300023101-0000331310202031): complete subsection reference.

<a id="canonical-3311300300002132-2110101212020221-1020322300311121-1003031002022221-2003223202222023-3231222120323220-2300201113222210-1223221211101111"></a>

## Next pages — kubeconfig_url / 213110113303 / 4

- [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info](resources--discovery--reference--group-001.md#canonical-1121112003223220-2111211201312211-3313302223212113-0312231220131131-0222122203020203-0231232101302010-1303022030200300-3333302311300130)
- [discovery_k8s.access_info.kubeconfig_url.clear_secret_info](resources--discovery--reference--group-001.md#canonical-0100213213212331-2003313330013022-3230232303212010-2201001131100021-2221103310311031-2131103302233031-0300232300023101-0000331310202031)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-1121112003223220-2111211201312211-3313302223212113-0312231220131131-0222122203020203-0231232101302010-1303022030200300-3333302311300130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223313122102332-2220222100300321-0310201332003033-3200001110233122-2212030000221112-0332331320333133-2333221102002023-0332031331133033"></a>

## discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info — blindfold_secret_info / 123301233122 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- [discovery_k8s.access_info.kubeconfig_url](resources--discovery--reference--group-001.md#canonical-2103233030101201-2311203221133210-0220202230223032-3211312010223020-0101303231022121-2330110001021033-1320230101103202-0211001030103130)
- discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info

<a id="canonical-3102230331112221-2100230112012011-1213331102102001-1033013222121133-1321300232220200-3103211011110120-3201221221313211-1102010012101102"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1130003121111100-3100332220131331-1231110033123013-2310333331101003-2020002022302003-3322321103123113-1130102233033030-2233012102100113"></a>

## Direct properties — blindfold_secret_info / 123301233122 / 3

<a id="canonical-1103303221030211-2100031021202022-1103001330010000-1112000113220022-1131001210023200-3201010100303000-1012131100132131-2113231232320102"></a>

<a id="canonical-3313031010321223-2021200223121131-0002132232032101-1123213033012221-1233321033003133-0001000000322203-2120212111100011-1320321121313012"></a>

## decryption_provider property — blindfold_secret_info / 123301233122 / 4

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

<a id="canonical-2032131212111000-2113001030322321-3103321103310131-3132300103123211-1000103021322300-3212300331122232-1110320333030020-3110313212203301"></a>

<a id="canonical-2120030030333222-1201120213331100-3022121220111300-1012210223321023-2222012022223313-2320331313220022-0222211112122331-3100122103222302"></a>

## location property — blindfold_secret_info / 123301233122 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0300312312030123-0300210331211232-0102212022211232-0330031300100231-2331101300312113-1101112230112111-3332110333331023-1301032122010120"></a>

<a id="canonical-0230111332133312-1312123310221320-3031012102230113-2211212133113000-3211300003211213-3110201220302223-2203211031000130-0113110023100313"></a>

## store_provider property — blindfold_secret_info / 123301233122 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-0131111120200113-3311300113000223-2123101032333321-1320311202123203-1112223331200123-3131003331210232-1223111212132022-0312212211001330"></a>

## Next pages — blindfold_secret_info / 123301233122 / 7

- [discovery_k8s.access_info.kubeconfig_url](resources--discovery--reference--group-001.md#canonical-2103233030101201-2311203221133210-0220202230223032-3211312010223020-0101303231022121-2330110001021033-1320230101103202-0211001030103130)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-0100213213212331-2003313330013022-3230232303212010-2201001131100021-2221103310311031-2131103302233031-0300232300023101-0000331310202031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331220233032312-3001310213200102-1010120300022331-2211331022321211-2122201022301111-1032311230302221-3132202122222011-0211000202002212"></a>

## discovery_k8s.access_info.kubeconfig_url.clear_secret_info — clear_secret_info / 113320221312 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- [discovery_k8s.access_info.kubeconfig_url](resources--discovery--reference--group-001.md#canonical-2103233030101201-2311203221133210-0220202230223032-3211312010223020-0101303231022121-2330110001021033-1320230101103202-0211001030103130)
- discovery_k8s.access_info.kubeconfig_url.clear_secret_info

<a id="canonical-0211013121120001-1030303311113122-0201210030321322-1123032022123210-0330030221311111-0321211123331033-1302021022031320-2133101012221301"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2123203033030022-3310021011321002-1321100120211111-0301220110023331-1031203010300010-2201231020301013-1220032303111330-3310210030132020"></a>

## Direct properties — clear_secret_info / 113320221312 / 3

<a id="canonical-2321131111120312-3223313210213103-3202010223011233-3322230223003033-0213110012232220-0333301230031312-3101233002100232-2230232033203233"></a>

<a id="canonical-1301102002212303-3122031220101202-2311231130312000-2031202221310133-3003102210230321-1100010012012003-0200030222232101-3312320022330330"></a>

## provider_ref property — clear_secret_info / 113320221312 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2301313313331230-3323110211231223-1221121122111023-2201111121121103-1002122012011313-2013200302123112-1302330001103103-1213111230320231"></a>

<a id="canonical-0322200111023203-3330333203030112-3021031330112013-3302221210310312-2132222100210110-2110300033213321-0322033300311130-3212320131031100"></a>

## URL property — clear_secret_info / 113320221312 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2230132021132133-3033211032033111-3232102213203011-0230131321220220-1010121322231123-2221032100320323-3221131010110113-2202011302133031"></a>

## Next pages — clear_secret_info / 113320221312 / 6

- [discovery_k8s.access_info.kubeconfig_url](resources--discovery--reference--group-001.md#canonical-2103233030101201-2311203221133210-0220202230223032-3211312010223020-0101303231022121-2330110001021033-1320230101103202-0211001030103130)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2233200231201012-1110023212011100-0220223121300122-2221312123303311-3102232213022203-2332332222223031-3301300320131201-2133101011010213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132300102232120-2002220230101101-0123001221023202-2313122300330132-0100332030232033-0230022032321122-2001210233123111-1111302131132122"></a>

## discovery_k8s.access_info.reachable — reachable / 000213023323 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- discovery_k8s.access_info.reachable

<a id="canonical-2112131220231232-0301102320201020-1232302112130232-1133221121220011-2330130313333122-3230001102122021-1022033201022213-1332332002220132"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
reachable = {}
```

<a id="canonical-1030033030222032-0221120002012301-2332210132011300-0210112022033111-3031002132000100-2322011111020222-3320101200133031-1302301020332202"></a>

## Direct properties — reachable / 000213023323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311020202023212-0130331233220003-1202332003030012-2102133200012203-1101231332032220-1111122123222132-2322321303021010-1113102132210213"></a>

## Next pages — reachable / 000213023323 / 4

- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2112000203012301-3300021232321103-2332023112322323-2303202002010111-0021011211110313-0120122221203123-0121321110201012-0332133223123202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101133031322030-0212101103210120-1122113001002111-2002230323303332-3331203322030220-3123033220203213-3031233301031223-1122303312130312"></a>

## discovery_k8s.default_all — default_all / 121313023022 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- discovery_k8s.default_all

<a id="canonical-2033200123113032-3001322221200211-2100033312102313-0312103132132013-1011101130212313-0001000000233301-2110230201220220-1320132000312113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default all.

Upstream description:

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
default_all = {}
```

<a id="canonical-1320302102200031-2302303310113201-1100302112211301-0312003103220002-1213032002320033-3310221020301211-2113200301233232-3200001303320300"></a>

## Direct properties — default_all / 121313023022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201310310222103-1111020331311301-3311302030130233-1202311321212313-0321311033232210-3033303023033223-2301020200323130-3220100033230313"></a>

## Next pages — default_all / 121313023022 / 4

- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-0321010022030131-3311200301332101-3132022310131230-1233220021130131-3210333300201303-3321221200212203-2130012322120030-1013021110331232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002130212032223-1332132121001013-2320333022003012-0012103010313020-1200032232303001-2302031032022133-2122020030313211-3231310223230231"></a>

## discovery_k8s.namespace_mapping — namespace_mapping / 001100002202 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- discovery_k8s.namespace_mapping

<a id="canonical-1130310213023212-1332013022102022-1111331100303002-2021113202231020-1230020021122231-3322220200003000-2322230130000120-0232111221000111"></a>

Type: `"object"`. single nested block, Optional.

Select the mapping between K8s namespaces from which services will be discovered and App Namespace
to which the discovered services will be shared.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("items")}
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
namespace_mapping {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203012303012331-3003112313012310-0132001312133023-0200210330022202-0220123232122012-2203130212002202-3000030001223320-2230320302223322"></a>

## Direct properties — namespace_mapping / 001100002202 / 3

- [items](resources--discovery--reference--group-001.md#canonical-3002302011033021-2032112031301321-3113010021213321-1211130323222133-0222322210032313-2033311001220112-2322102303032012-1211323103133123): complete subsection reference.

<a id="canonical-3023333023010211-3010300322123010-2032202133130020-2231130133103211-2321101131203303-1203321231023300-3301210102021211-0120313123220132"></a>

## Next pages — namespace_mapping / 001100002202 / 4

- [discovery_k8s.namespace_mapping.items](resources--discovery--reference--group-001.md#canonical-3002302011033021-2032112031301321-3113010021213321-1211130323222133-0222322210032313-2033311001220112-2322102303032012-1211323103133123)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-3002302011033021-2032112031301321-3113010021213321-1211130323222133-0222322210032313-2033311001220112-2322102303032012-1211323103133123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110013112112131-1032022130113303-3233202111113110-0123231002331032-0113133020211003-1121312132103323-2030012230300333-0003103230122312"></a>

## discovery_k8s.namespace_mapping.items — items / 200123212113 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.namespace_mapping](resources--discovery--reference--group-001.md#canonical-0321010022030131-3311200301332101-3132022310131230-1233220021130131-3210333300201303-3321221200212203-2130012322120030-1013021110331232)
- discovery_k8s.namespace_mapping.items

<a id="canonical-3112000132011301-1222010321120031-3203310201120101-2133032230221032-1210122030302003-2113022110300221-2203221203330020-3021321032100331"></a>

Type: `"object"`. list nested block, Optional.

Map K8s namespace(s) to App Namespaces. In Shared Configuration, Discovered Services can only be
mapped to a single App Namespace, which is determined by the first matched regular expression.

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
items {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131332121223110-1012322223313121-1210111110010013-2013330011302311-0210102133221333-3113032211002200-1110103020112331-2121332023113322"></a>

## Direct properties — items / 200123212113 / 3

<a id="canonical-3201112330030203-1133302310200330-1102230113001203-3002032231013323-3222112223222010-1031010002130102-3301331110032301-2122202331201332"></a>

<a id="canonical-0132032132132232-0321320200302000-1231112232121003-3020122001200002-3123302100030310-1130333310230321-1010011320012111-0002000321313320"></a>

## namespace property — items / 200123212113 / 4

Type: `"string"`. Optional, Computed.

F5XC Application Namespaces. Select a namespace.

Upstream description:

Select a namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-1130100022032112-2301113201110213-3223002220033003-0303132103013200-0001103312021020-0221103001202012-2132111201101330-3200113123101210"></a>

<a id="canonical-2333211130032100-1211202132322013-1201030202003321-1002113301231232-2312102000121231-1212211013221000-2303222320111012-3231312000233121"></a>

## namespace_regex property — items / 200123212113 / 5

Type: `"string"`. Optional.

The regular expression here will be used to match K8s namespace(s).

Provider validators and defaults (from schema source):

```go
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1202211332011103-0230133332310133-1313111033123303-2030233030222030-2130030302111213-3022120210232233-3330023223211133-2200030110221302"></a>

## Next pages — items / 200123212113 / 6

- [discovery_k8s.namespace_mapping](resources--discovery--reference--group-001.md#canonical-0321010022030131-3311200301332101-3132022310131230-1233220021130131-3210333300201303-3321221200212203-2130012322120030-1013021110331232)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230120202310122-1100012123022121-3201101312103310-3230300302002012-1101102311031120-1310332211313022-2231322302301022-3030212132331112"></a>

## discovery_k8s.publish_info — publish_info / 300203032120 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- discovery_k8s.publish_info

<a id="canonical-3122222231121203-3012003210123330-0330311123322320-2213313323121323-0113123300203332-1113032311231020-0323112032313021-3223031122220121"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for publish info.

Upstream description:

K8s Configuration to publish VIPs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "dns_delegation"),
  validators.ConflictingObjectAttributes("disable_spec",
    "publish"),
  validators.ConflictingObjectAttributes("disable_spec",
    "publish_fqdns"),
  validators.ConflictingObjectAttributes("dns_delegation",
    "publish"),
  validators.ConflictingObjectAttributes("dns_delegation",
    "publish_fqdns"),
  validators.ConflictingObjectAttributes("publish",
    "publish_fqdns")}
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
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"dns_delegation\",\"publish\",\"publish_fqdns\"]"
}
```

Terraform syntax:

```terraform
publish_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232102302021022-1301211313313232-0221030010002202-2131102321123123-3002231312102320-2122301320313333-0103232212220203-3221012011000120"></a>

## Direct properties — publish_info / 300203032120 / 3

- [disable_spec](resources--discovery--reference--group-001.md#canonical-3011020302001133-2202023030221333-1231203031313222-1323230201300200-0033331312113001-1023022000002023-1113301030132032-0031131021301222): complete subsection reference.

- [dns_delegation](resources--discovery--reference--group-001.md#canonical-0213033233320133-2313331011023030-3300221322303212-3012012303031332-2322003103212311-2232013223120231-0223002211302120-2121012230012320): complete subsection reference.

- [publish](resources--discovery--reference--group-001.md#canonical-2201132213231022-0230303111210031-0130220302231300-2323030120010232-2210223132002011-2002001223211111-3210202213010231-0112131333302012): complete subsection reference.

- [publish_fqdns](resources--discovery--reference--group-001.md#canonical-2320301002300300-0102133000123033-1212202333001332-0013302020101020-3320231121013113-0232002100202122-2121011332331332-2121202110113103): complete subsection reference.

<a id="canonical-2233100210213023-3321032103112202-1310111002001010-3130012131131012-0303132311310332-2033201231033131-2312232322020002-1100223123011111"></a>

## Next pages — publish_info / 300203032120 / 4

- [discovery_k8s.publish_info.disable_spec](resources--discovery--reference--group-001.md#canonical-3011020302001133-2202023030221333-1231203031313222-1323230201300200-0033331312113001-1023022000002023-1113301030132032-0031131021301222)
- [discovery_k8s.publish_info.dns_delegation](resources--discovery--reference--group-001.md#canonical-0213033233320133-2313331011023030-3300221322303212-3012012303031332-2322003103212311-2232013223120231-0223002211302120-2121012230012320)
- [discovery_k8s.publish_info.publish](resources--discovery--reference--group-001.md#canonical-2201132213231022-0230303111210031-0130220302231300-2323030120010232-2210223132002011-2002001223211111-3210202213010231-0112131333302012)
- [discovery_k8s.publish_info.publish_fqdns](resources--discovery--reference--group-001.md#canonical-2320301002300300-0102133000123033-1212202333001332-0013302020101020-3320231121013113-0232002100202122-2121011332331332-2121202110113103)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-3011020302001133-2202023030221333-1231203031313222-1323230201300200-0033331312113001-1023022000002023-1113301030132032-0031131021301222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303112213203030-2130230123313120-2110013113332203-3333013000121313-3210023321331213-3000210321020023-1023020100220222-2222201223031211"></a>

## discovery_k8s.publish_info.disable_spec — disable_spec / 323023332021 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113)
- discovery_k8s.publish_info.disable_spec

<a id="canonical-1131112011113030-1210133212003130-3311313320312101-1010030313001132-3013331233013300-1312020132001211-2322023011103123-0311120302310100"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-1302330303122021-3211113133131032-3330220212133003-0330012232110212-1331311003322013-0130201322000121-2133210200133220-1201212303313121"></a>

## Direct properties — disable_spec / 323023332021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022232233112230-0130301110123201-1201012332221220-2231330322000310-2332231130031113-0221313033033123-1032220122232331-2010212022022131"></a>

## Next pages — disable_spec / 323023332021 / 4

- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-0213033233320133-2313331011023030-3300221322303212-3012012303031332-2322003103212311-2232013223120231-0223002211302120-2121012230012320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132203212213332-3233011123102333-3133001200031320-1223230111132131-3121002332213230-0111000031003200-0223102201300233-1112011220201333"></a>

## discovery_k8s.publish_info.dns_delegation — dns_delegation / 221301012203 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113)
- discovery_k8s.publish_info.dns_delegation

<a id="canonical-0013022331002220-1313110232311321-3210232322133213-0021213003322120-1312002013122030-3330223222122301-3013202002112203-0303031222230100"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for DNS delegation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subdomain")}
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
dns_delegation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3001232222022321-2011332132210212-3303111021103222-2331302212012001-3130223302311111-3021233130312132-0311203022032000-1010201300300132"></a>

## Direct properties — dns_delegation / 221301012203 / 3

<a id="canonical-1332300010111012-3302320013202330-0111130212331013-0130321000210330-3310023321020232-0010333012131112-1002231112021123-2320002323102121"></a>

<a id="canonical-3111030201222102-3110303031012102-0330212030021302-2000232003323232-1000300031010102-0123022130033012-0133331013210002-0220320313231032"></a>

## dns_mode property — dns_delegation / 221301012203 / 4

Type: `"string"`. Optional.

\[Enum: CORE\_DNS|KUBE\_DNS\] Two modes are possible CoreDNS: Whether external K8s cluster is
running core-DNS KubeDNS: External K8s cluster is running kube-DNS. Possible values are
\`CORE\_DNS\`, \`KUBE\_DNS\`. Defaults to \`CORE\_DNS\`.

Upstream description:

Two modes are possible

CoreDNS: Whether external K8s cluster is running core-DNS KubeDNS: External K8s cluster is running
kube-DNS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CORE_DNS",
    "KUBE_DNS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CORE_DNS",
  "enum": [
    "CORE_DNS",
    "KUBE_DNS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1022211331330030-1332003123210002-3333231111322323-3323233032003011-2333223123033210-2202100211311321-3310101233321103-2333210311000310"></a>

<a id="canonical-0232033133321330-0101001123221331-0103220111320213-1332000121113232-1321120320321200-0103223022002311-2103103212002332-3031210230013130"></a>

## subdomain property — dns_delegation / 221301012203 / 5

Type: `"string"`. Optional.

The DNS subdomain for which F5XC will respond to DNS queries.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2012030010010232-3032031132322032-2311012300000021-0120101311000133-0100213221001333-1312333000313320-3200311212023003-3103302133300132"></a>

## Next pages — dns_delegation / 221301012203 / 6

- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2201132213231022-0230303111210031-0130220302231300-2323030120010232-2210223132002011-2002001223211111-3210202213010231-0112131333302012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212002102030132-0031103231130313-0103302300010333-3313322032112201-0321010231032101-1230000100320222-2010132010122100-2122303222012111"></a>

## discovery_k8s.publish_info.publish — publish / 101201323110 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113)
- discovery_k8s.publish_info.publish

<a id="canonical-1313310112011001-2301122021130021-2233122213001303-3132203303203013-0132212011210030-1110203023102312-2032012331000200-0031120213233032"></a>

Type: `"object"`. single nested block, Optional.

K8SPublishType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("namespace")}
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
publish {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133122133212220-2231303003011131-0322321332101230-1013233332132013-0131222012010300-0131331031012331-1223003200023113-0322001021210013"></a>

## Direct properties — publish / 101201323110 / 3

<a id="canonical-2301110233320203-0021313232121333-1220131003100110-3221033032110200-2302200212010212-1113301212023221-0210121133120020-1213010011022120"></a>

<a id="canonical-2210231202320212-1220322123332111-0033133123013100-3303220111213223-1120011100022233-3020300130301220-3221023301333222-0233111300023300"></a>

## namespace property — publish / 101201323110 / 4

Type: `"string"`. Optional, Computed.

The namespace where the service/endpoints need to be created if it's not included in the domain. The
external K8s administrator needs to ensure that the namespace exists.

Upstream description:

The namespace where the service/endpoints need to be created if it's not included in the domain. The
external K8s administrator needs to ensure that the namespace exists.

Provider validators and defaults (from schema source):

```go
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
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-3323102210012221-1030110000223022-3033110323013003-2331312211210000-1031220210020212-2302322223010320-3033220020332113-2031003003031312"></a>

## Next pages — publish / 101201323110 / 5

- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2320301002300300-0102133000123033-1212202333001332-0013302020101020-3320231121013113-0232002100202122-2121011332331332-2121202110113103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023022303113102-0200220221220000-0031330003130233-2321323022200211-3213302330330030-2103103312123301-1311101320130032-2230332313122001"></a>

## discovery_k8s.publish_info.publish_fqdns — publish_fqdns / 103311031232 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113)
- discovery_k8s.publish_info.publish_fqdns

<a id="canonical-1031022023310213-0232313203112320-1213302020313332-0311311033232210-2302323330201030-2113123312203013-2323321312021030-2222203323321201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for publish fqdns.

Upstream description:

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
publish_fqdns = {}
```

<a id="canonical-3132022203313322-0322033330203303-1102113312013110-2123113113001322-0332133203322323-0300000211331103-3003202030300010-1320111332132200"></a>

## Direct properties — publish_fqdns / 103311031232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233021130212123-1202113030131003-1033020311111002-3322303033211001-1032232131211301-2113103102012202-2301211312310132-1223100333133131"></a>

## Next pages — publish_fqdns / 103311031232 / 4

- [discovery_k8s.publish_info](resources--discovery--reference--group-001.md#canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2332001313120022-1313323002112231-2021132023022331-0131012222332133-3231233201102020-0102032212311132-1210113121202130-1000203132133103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103221133110232-2021132001033202-0020002022112110-1000203303021300-0000121010021202-2102230210221200-1031113131331323-1300011123323123"></a>

## no_cluster_id — no_cluster_id / 302132223201 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- no_cluster_id

<a id="canonical-3002323030220322-0020200231302202-3111002300320220-3022101033030022-2021313300131030-0300131003010012-3302102103231001-3130001221002120"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
no_cluster_id = {}
```

<a id="canonical-0300131112330100-0120233011230201-3331123321132010-1231003011002202-1012201200032231-1312020123200232-1212113112310001-1232322333112321"></a>

## Direct properties — no_cluster_id / 302132223201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103223203123013-2300200222013312-3130220011122221-3003322231001102-3002230213221002-0322031120101121-3302320213222032-0333202232111033"></a>

## Next pages — no_cluster_id / 302132223201 / 4

- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-0133122212320132-3323032213033022-3031001120032311-3120311312312302-0221333302003213-0021323013310132-3013211230313033-1331213120121300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332223100213303-1223222212222101-3123110331221210-1332210000203222-1201303001203212-3232130100013301-0310221201012121-1222323300131220"></a>

## timeouts — timeouts / 223211133220 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- timeouts

<a id="canonical-0012101032103333-1011032300121101-1320200130220320-3123221213101133-3033312102212232-3230131201211022-0330321023111232-2312201331323113"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311232203213221-0223010331121013-1123002301311311-2203021113132110-3300202202112303-1301012220101320-2010210203111113-2022301020133330"></a>

## Direct properties — timeouts / 223211133220 / 3

<a id="canonical-1112033131022101-2022013210222232-1032111211310023-3103132312130230-0332013311310200-1320023323331223-2123323113333110-2102223332202222"></a>

<a id="canonical-3222123220210310-2003133130222002-0103102112302132-1203311133301211-2020220322102312-3330022203030013-1311023020202101-3231030003000001"></a>

## create property — timeouts / 223211133220 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2202202032033022-2300320013310321-3032002000003122-3311310323233233-1110211012322012-2223330110301111-0021011112302010-3303323300133200"></a>

<a id="canonical-1301100033110110-1202212100000300-3012120022002212-2322311033001203-1203031103101221-0330010023011133-1321301202312133-2203203130301132"></a>

## delete property — timeouts / 223211133220 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1011000201012302-1011300001301023-1101211332122121-0012132102332023-0132133332030322-2020132020301311-1330332300202211-1221120303031121"></a>

<a id="canonical-3202302113011013-1331201313220002-0202311130130131-1201011113222101-3213113030220023-3022032022323332-3320302121100331-0213133302102000"></a>

## read property — timeouts / 223211133220 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3231100130200331-3321310232220101-2113300013331202-2123333311110333-0103113330100211-0002013301300101-3333023313130001-0202100110313023"></a>

<a id="canonical-3330000033220111-2303202111130002-0011303321023323-1013213223221123-0023202133111131-1101010231300212-2121122102321012-1333022310103233"></a>

## update property — timeouts / 223211133220 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0130011112303312-1322212112312333-1320313101003230-0200031111222323-3211032211111323-0201203123011111-3231220112030033-1100012003232221"></a>

## Next pages — timeouts / 223211133220 / 8

- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032123011230031-2030312203122113-2302100022101232-1003222232120031-3103133201211302-0321013120201111-2302310022321010-0303111120201010"></a>

## where — where / 323101133103 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- where

<a id="canonical-1321103302233201-0332121301121100-3101133031230333-0133220331333201-0030002311302101-3220122332233201-0213131123032010-1102233222333213"></a>

Type: `"object"`. single nested block, Optional.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local..

Upstream description:

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingObjectAttributes("virtual_network",
    "virtual_site")}
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
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
where {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202313221332303-2002312102133132-0332110000230022-1102012010130131-2233121001202320-2212322031030330-2002003232123332-1230320200130220"></a>

## Direct properties — where / 323101133103 / 3

- [site](resources--discovery--reference--group-001.md#canonical-2111032303022223-2300132122211213-0100203302230222-1011230233321231-1103310111023200-0013013033313032-1103302322310010-1210312011210200): complete subsection reference.

- [virtual_network](resources--discovery--reference--group-002.md#canonical-1110120031123232-1103133033221300-2333022012020332-2112320331303032-0100001133321112-2231130320002221-2033320232220130-0031211123210313): complete subsection reference.

- [virtual_site](resources--discovery--reference--group-002.md#canonical-2223322232211101-1111031212021233-2331110221020110-1113201232123120-1032003111222123-1221012223203322-0130200312133221-2222033300110330): complete subsection reference.

<a id="canonical-0333212013200001-1130031232123322-2023001300311102-2003201200122011-3231313200003203-0202233333022320-3221020010321032-0220013330001322"></a>

## Next pages — where / 323101133103 / 4

- [where.site](resources--discovery--reference--group-001.md#canonical-2111032303022223-2300132122211213-0100203302230222-1011230233321231-1103310111023200-0013013033313032-1103302322310010-1210312011210200)
- [where.virtual_network](resources--discovery--reference--group-002.md#canonical-1110120031123232-1103133033221300-2333022012020332-2112320331303032-0100001133321112-2231130320002221-2033320232220130-0031211123210313)
- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-2223322232211101-1111031212021233-2331110221020110-1113201232123120-1032003111222123-1221012223203322-0130200312133221-2222033300110330)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2111032303022223-2300132122211213-0100203302230222-1011230233321231-1103310111023200-0013013033313032-1103302322310010-1210312011210200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001332330111000-0210021310011302-0200010321220301-3231021112300313-1333300020102230-2321200023332123-1022101022002201-1213100103102223"></a>

## where.site — site / 012130210213 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- where.site

<a id="canonical-2012013100322022-1022120320103220-1213011202221212-0031301113211120-0121000101211123-3131302320112311-0011001103132122-2312321232311020"></a>

Type: `"object"`. single nested block, Optional.

Specifies a direct reference to a site configuration object.

Upstream description:

This specifies a direct reference to a site configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001010210130201-0133120123311112-1231223132230233-1203201132201320-0232231031131232-2000320011310002-3011313311303211-1200223103230011"></a>

## Direct properties — site / 012130210213 / 3

- [disable_internet_vip](resources--discovery--reference--group-001.md#canonical-1132310133301223-1201221101101031-0330233322032030-0023233223120110-0230220023201310-2333220211300312-0303123103120213-0001032331220103): complete subsection reference.

- [enable_internet_vip](resources--discovery--reference--group-001.md#canonical-2110231323102212-0020122101231302-3330133200323223-2201010003300321-2322213312332232-3102021022010321-1320030113212012-2303013132303030): complete subsection reference.

<a id="canonical-1223122302203100-3200331030230313-3033220330302303-2130131323322120-1023013321001130-2221000311303032-1010022102102223-2111232233001112"></a>

<a id="canonical-3203010123001323-3020223102003012-1102002022332200-0033233122201132-3021200302000123-0033213321020232-2222131120301121-0331113112313110"></a>

## network_type property — site / 012130210213 / 4

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](resources--discovery--reference--group-002.md#canonical-2032210203231300-1003003000132213-3101202330201122-0001100213311020-0131210232302123-3012213102000232-2331302321030002-1313330003132301): complete subsection reference.

<a id="canonical-3133033020011320-2102002212330011-1000330101203123-3312020311300132-0322313003031333-0202011223201033-2211112022131112-1321223231032221"></a>

## Next pages — site / 012130210213 / 5

- [where.site.disable_internet_vip](resources--discovery--reference--group-001.md#canonical-1132310133301223-1201221101101031-0330233322032030-0023233223120110-0230220023201310-2333220211300312-0303123103120213-0001032331220103)
- [where.site.enable_internet_vip](resources--discovery--reference--group-001.md#canonical-2110231323102212-0020122101231302-3330133200323223-2201010003300321-2322213312332232-3102021022010321-1320030113212012-2303013132303030)
- [where.site.ref](resources--discovery--reference--group-002.md#canonical-2032210203231300-1003003000132213-3101202330201122-0001100213311020-0131210232302123-3012213102000232-2331302321030002-1313330003132301)
- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-1132310133301223-1201221101101031-0330233322032030-0023233223120110-0230220023201310-2333220211300312-0303123103120213-0001032331220103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122333031312112-2211211221000120-2222303213222131-2131000122202221-1323320332202233-0122000120122200-2212313330222120-2123022302110111"></a>

## where.site.disable_internet_vip — disable_internet_vip / 032303201101 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-001.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [where.site](resources--discovery--reference--group-001.md#canonical-2111032303022223-2300132122211213-0100203302230222-1011230233321231-1103310111023200-0013013033313032-1103302322310010-1210312011210200)
- where.site.disable_internet_vip

<a id="canonical-2130130203303330-3032002112302222-3303311021103102-3220300023331330-2220110300021020-3130203123022331-2211123013020023-2212023112133033"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
disable_internet_vip = {}
```

<a id="canonical-3230302033113212-1313300101211121-1122203033013201-1100300122002211-2020123113320223-2021312202211233-3212020132321313-1333221231220121"></a>

## Direct properties — disable_internet_vip / 032303201101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130120330223111-3331003233132112-0221103231130322-0321202303312113-1123330012203223-0323111320122320-2231221212003233-2121022130012032"></a>

## Next pages — disable_internet_vip / 032303201101 / 4

- [where.site](resources--discovery--reference--group-001.md#canonical-2111032303022223-2300132122211213-0100203302230222-1011230233321231-1103310111023200-0013013033313032-1103302322310010-1210312011210200)
- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)

<a id="canonical-2110231323102212-0020122101231302-3330133200323223-2201010003300321-2322213312332232-3102021022010321-1320030113212012-2303013132303030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
