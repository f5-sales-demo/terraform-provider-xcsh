---
page_title: "xcsh_log_receiver reference"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_log_receiver reference."
---

# xcsh_log_receiver reference

<a id="canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202020333203233-3311231222031003-2010001023100120-3213100313112232-0030021213132031-1002110131301220-3200000133013330-1203310030103131"></a>

## Property reference — Property reference / 232031230101 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- Property reference

<a id="canonical-2231312323011222-1130012210113111-1022230300321322-1230313112232230-2112313213313010-2231322210022002-3213322233122021-2201222320301201"></a>

## Direct properties — Property reference / 232031230101 / 3

<a id="canonical-2213201110211310-3212322101232101-2313130233201231-1221221323000100-1330123321031021-3203330311323302-1022021312330303-2133002100230002"></a>

<a id="canonical-0121022323231112-2003112031132033-3112032212023311-2113323201211132-3213221123222112-2001120112131300-2303111120223133-3122323313201211"></a>

## annotations property — Property reference / 232031230101 / 4

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

<a id="canonical-2021312322130200-2022303300310320-3020321222201302-3211232221113302-3203121232303020-2100110321211120-1011033131121022-1132233002220311"></a>

<a id="canonical-0221113203110100-1322111231020212-1202023320302330-2221302012030221-1123311013231301-3301311110200123-3001330001223232-2311101332021330"></a>

## description property — Property reference / 232031230101 / 5

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

<a id="canonical-2002002100113122-1013020202110221-1010230332213131-3003311033333000-0330321032202011-3023101300332221-2003303122301221-0213203000111222"></a>

<a id="canonical-0113213100321131-3021020100101033-0120113111223221-3001031120211300-3111313013203331-0012100211233200-2311203310221021-2130301110210311"></a>

## disable property — Property reference / 232031230101 / 6

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

<a id="canonical-0122200323213320-2003303113100221-0220323231102110-2301202011320133-3231330120312131-3012011312033312-2322013030221121-3303311013320202"></a>

<a id="canonical-3013210032020010-0100022303131213-1311230333123110-0312322223032123-1322003000011312-2101323221001101-1310123021301321-3011110220110122"></a>

## ID property — Property reference / 232031230101 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3132313022102022-3300232121213031-0312111312022021-1021303122212232-0131012120032101-0312331001233203-2003312023322123-0232003002122023"></a>

<a id="canonical-2213312030321003-2030231133232213-2113213010203102-0312011321010210-3121301332013010-0113302032022221-3211223131031122-3102020303233132"></a>

## labels property — Property reference / 232031230101 / 8

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

<a id="canonical-3203303231222220-1233221001002322-1013110133331303-0010112020111002-1200203012122221-0223333230103103-2113320120112012-0012002123030103"></a>

<a id="canonical-1020110200331213-2332302030103102-0112022230030231-0330313213200020-0331300123133003-1020031022000132-2103132301122333-2020203130220223"></a>

## name property — Property reference / 232031230101 / 9

Type: `"string"`. Required.

Name of the Log Receiver. Must be unique within the namespace.

Upstream description:

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

<a id="canonical-3100102313313332-2203323123103232-3332222021321333-2322002132130102-0313023002302020-3203133030330130-2201303100230212-2333011212123012"></a>

<a id="canonical-2331023211100212-3121233032100312-0333133311201001-0123120133200011-1231022100220213-2300002320130202-3131230032302232-3131322021301301"></a>

## namespace property — Property reference / 232031230101 / 10

Type: `"string"`. Required.

Namespace where the Log Receiver is created.

Upstream description:

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

- [site_local](resources--log_receiver--reference--group-001.md#canonical-2123333131201320-0023300212320202-2031023130110212-3311321012212320-0202201223021122-3302013111313113-1302212021021210-0203022212321031): complete subsection reference.

- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231): complete subsection reference.

- [timeouts](resources--log_receiver--reference--group-001.md#canonical-2130131322321001-1002201220120233-2233232321020120-3201102003221011-2022123000123110-2333132333113230-2022223022200022-1032320120302321): complete subsection reference.

<a id="canonical-0200320330003303-1321010013101010-3130201102131232-2100211201122212-0211021331131131-1011123122111221-3301322301131222-3122001300311003"></a>

## All schema paths — Property reference / 232031230101 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--log_receiver--reference--group-001.md#canonical-2213201110211310-3212322101232101-2313130233201231-1221221323000100-1330123321031021-3203330311323302-1022021312330303-2133002100230002) |
| `description` | [description](resources--log_receiver--reference--group-001.md#canonical-2021312322130200-2022303300310320-3020321222201302-3211232221113302-3203121232303020-2100110321211120-1011033131121022-1132233002220311) |
| `disable` | [disable](resources--log_receiver--reference--group-001.md#canonical-2002002100113122-1013020202110221-1010230332213131-3003311033333000-0330321032202011-3023101300332221-2003303122301221-0213203000111222) |
| `id` | [ID](resources--log_receiver--reference--group-001.md#canonical-0122200323213320-2003303113100221-0220323231102110-2301202011320133-3231330120312131-3012011312033312-2322013030221121-3303311013320202) |
| `labels` | [labels](resources--log_receiver--reference--group-001.md#canonical-3132313022102022-3300232121213031-0312111312022021-1021303122212232-0131012120032101-0312331001233203-2003312023322123-0232003002122023) |
| `name` | [name](resources--log_receiver--reference--group-001.md#canonical-3203303231222220-1233221001002322-1013110133331303-0010112020111002-1200203012122221-0223333230103103-2113320120112012-0012002123030103) |
| `namespace` | [namespace](resources--log_receiver--reference--group-001.md#canonical-3100102313313332-2203323123103232-3332222021321333-2322002132130102-0313023002302020-3203133030330130-2201303100230212-2333011212123012) |
| `site_local` | [site_local](resources--log_receiver--reference--group-001.md#canonical-3203223331213211-1233233020023033-1000002111113023-3311322112102133-2000202333013331-0003110312203301-1210011223211110-2101013220320223) |
| `syslog` | [syslog](resources--log_receiver--reference--group-001.md#canonical-3123113221322213-1100032231313132-2203001020313022-2100203100001201-1311111121022111-3333020322312231-3220002033102212-3023013323133201) |
| `syslog.syslog_rfc5424` | [syslog.syslog_rfc5424](resources--log_receiver--reference--group-001.md#canonical-1131012122212021-2032330210012222-2312121110332133-0312210201133031-0131322213312232-1231201112020311-1032212310332230-0202212322330103) |
| `syslog.tcp_server` | [syslog.tcp_server](resources--log_receiver--reference--group-001.md#canonical-1211233201000232-2101222213211010-3302223112221221-3101203301112300-2111030223030121-2213132320232210-0100220110310302-1301021213331013) |
| `syslog.tcp_server.port` | [syslog.tcp_server.port](resources--log_receiver--reference--group-001.md#canonical-3110220212123210-0221013330001032-3100013102222223-1312131311302120-2002302012333022-0033112101101033-3032000301113311-0110133101003303) |
| `syslog.tcp_server.server_name` | [syslog.tcp_server.server_name](resources--log_receiver--reference--group-001.md#canonical-0013230223111130-0301001000133100-1100211201130012-3001333122203323-3022321231120023-0113112230132322-0210302321121200-2022231113002020) |
| `syslog.tls_server` | [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-0020023302001203-2233301310211022-3323120001303023-0213110012010310-0233031321013002-0001203023203232-2212233122213231-0001130311023320) |
| `syslog.tls_server.default_https_port` | [syslog.tls_server.default_https_port](resources--log_receiver--reference--group-001.md#canonical-1012312013211300-3221032102101313-1003223132312223-3123132220120021-0000011021212121-2311200012233231-3332210102230023-3313121320101323) |
| `syslog.tls_server.default_syslog_tls_port` | [syslog.tls_server.default_syslog_tls_port](resources--log_receiver--reference--group-001.md#canonical-3222313010303000-3032212132132020-0213333011303011-2031233303013011-0000223132011120-1122320112210132-0102130300122230-2122010321300313) |
| `syslog.tls_server.mtls_disabled` | [syslog.tls_server.mtls_disabled](resources--log_receiver--reference--group-001.md#canonical-3331130130002302-1020301230023101-3232300211021321-2321313021102023-1130100010210033-0112002021232102-0211200333333023-3223133210230331) |
| `syslog.tls_server.mtls_enable` | [syslog.tls_server.mtls_enable](resources--log_receiver--reference--group-001.md#canonical-0332023303013301-3022011103203220-0201120130313020-0213011220232223-2012130132231112-0033333032011030-1020311232222012-3232223021103103) |
| `syslog.tls_server.mtls_enable.certificate` | [syslog.tls_server.mtls_enable.certificate](resources--log_receiver--reference--group-001.md#canonical-3110101130011313-2012102331223221-0212112033223311-1132123202032002-2321202031323211-3313121013302022-1332230112230103-0131211233113202) |
| `syslog.tls_server.mtls_enable.key_url` | [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--reference--group-001.md#canonical-2021123211000111-3121133203232103-0313321201123002-2310032313202013-1333033221201223-2113233131202132-0201212210103102-2301023332011003) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info](resources--log_receiver--reference--group-001.md#canonical-3030122031120301-1031330013211113-1302330102121110-2113332132311130-2010223001233122-3112032000230221-2113303300033310-0213003323202111) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.decryption_provider](resources--log_receiver--reference--group-001.md#canonical-1103110020131000-1010032231012332-0131231303110302-1130310231110210-2332230313202010-3022020023111331-2031001312210210-3100212133210001) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.location` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.location](resources--log_receiver--reference--group-001.md#canonical-2123120103003103-1323210201222321-0303312123220233-1123210122320012-3110003030311201-0213122000322301-1302030103202311-2131323112200021) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.store_provider` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.store_provider](resources--log_receiver--reference--group-001.md#canonical-3120211021013201-3303211003220013-3110320201221301-0213022001211223-1300332312032202-3202221103311332-2020032231103311-1010332312330300) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info](resources--log_receiver--reference--group-001.md#canonical-0103103311310111-2223330012323321-1133110012202032-2003233310012301-2001013110032122-1323131210310302-3201032202302330-0023020013111213) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info.provider_ref` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info.provider_ref](resources--log_receiver--reference--group-001.md#canonical-2201030130030210-2300231121322121-1312121112111311-3101110231313210-1313120123213331-3213332211032232-3302322200300030-2012301110332121) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info.url` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info.url](resources--log_receiver--reference--group-001.md#canonical-2320302103121113-2131302322311203-3202121203200120-0203123120303103-3230230331303312-3011313102003120-0012001013033013-3302323003133330) |
| `syslog.tls_server.port` | [syslog.tls_server.port](resources--log_receiver--reference--group-001.md#canonical-1221303200022120-1312023310211122-0203222210013100-1100312311130113-3120123103121221-1312312201233220-3021332001300301-1212120122322230) |
| `syslog.tls_server.server_name` | [syslog.tls_server.server_name](resources--log_receiver--reference--group-001.md#canonical-2013113212320010-1112002010003210-3111302323221101-3133112120132021-1121301032011101-3012210313120033-0201013132311203-3003002120203301) |
| `syslog.tls_server.trusted_ca_url` | [syslog.tls_server.trusted_ca_url](resources--log_receiver--reference--group-001.md#canonical-0103013302112311-0201012230012001-0000213233113021-3213013211220101-3013321011323310-2002313213103230-2230130210312312-0000232002132103) |
| `syslog.tls_server.volterra_ca` | [syslog.tls_server.volterra_ca](resources--log_receiver--reference--group-001.md#canonical-0330202220100300-0333210333133310-0322012101100011-0302201013023012-3113132203101220-3331223211112200-3133331100101311-0033302110220300) |
| `syslog.udp_server` | [syslog.udp_server](resources--log_receiver--reference--group-001.md#canonical-3222232131211113-2312132322233112-1301323311131000-0210121312101002-2321331131123212-0112132132133103-1232132213103032-1033102103130313) |
| `syslog.udp_server.port` | [syslog.udp_server.port](resources--log_receiver--reference--group-001.md#canonical-0233031211203031-2032213010100022-0110100320230312-0323131130022221-2113003331103211-2331122000332110-3202223322123302-2001020023320302) |
| `syslog.udp_server.server_name` | [syslog.udp_server.server_name](resources--log_receiver--reference--group-001.md#canonical-1130121111032103-1112230102331200-1221310320301013-2221001130023001-2210101213122310-3130110322232031-3012013212211231-1301020320313121) |
| `timeouts` | [timeouts](resources--log_receiver--reference--group-001.md#canonical-2002112103220323-2002111103030222-1203320302331133-0321223322200222-3133300230101001-2003201333102022-2102112111303000-3033332331323203) |
| `timeouts.create` | [timeouts.create](resources--log_receiver--reference--group-001.md#canonical-1030310022213212-1111320202303331-1332010323300102-0103322201033203-0013330233222231-0121033231301022-0233302102322113-3230230223122302) |
| `timeouts.delete` | [timeouts.delete](resources--log_receiver--reference--group-001.md#canonical-2030323000223333-2321133110013011-2102310002312111-2200230333001021-2210220203022302-0221233213103133-1222000003033323-2301011101130123) |
| `timeouts.read` | [timeouts.read](resources--log_receiver--reference--group-001.md#canonical-2123101212231330-2022002202300212-2320123332232231-0003233001312230-2310003120010230-1112223313222101-3320320201203203-3003323033210002) |
| `timeouts.update` | [timeouts.update](resources--log_receiver--reference--group-001.md#canonical-0333131231213231-2330013011122022-0321001100002122-2002320301200331-3232013231331202-2322301312230322-0113132001232020-0102111222100203) |

<a id="canonical-1113302102020133-1021301313010130-0101302330231013-1013221231032312-3102120123011032-1203222300131122-3020131000202012-3121121031102202"></a>

## Next pages — Property reference / 232031230101 / 12

- [site_local](resources--log_receiver--reference--group-001.md#canonical-2123333131201320-0023300212320202-2031023130110212-3311321012212320-0202201223021122-3302013111313113-1302212021021210-0203022212321031)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- [timeouts](resources--log_receiver--reference--group-001.md#canonical-2130131322321001-1002201220120233-2233232321020120-3201102003221011-2022123000123110-2333132333113230-2022223022200022-1032320120302321)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)

<a id="canonical-2123333131201320-0023300212320202-2031023130110212-3311321012212320-0202201223021122-3302013111313113-1302212021021210-0203022212321031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120003200020202-1303112211320123-3313132123212121-0332123332213112-0333103011310010-3131033132330132-1120333322311320-3223021132011103"></a>

## site_local — site_local / 203000220221 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- site_local

<a id="canonical-3203223331213211-1233233020023033-1000002111113023-3311322112102133-2000202333013331-0003110312203301-1210011223211110-2101013220320223"></a>

Type: `"object"`. single nested block, Optional.

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
site_local {}
```

<a id="canonical-0031002002033210-3322211210031022-0130203100212131-1103011130110212-2001320002302103-3100201223201113-3303000010300122-1030003220321233"></a>

## Direct properties — site_local / 203000220221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023032321230131-3233033222032000-3113132033122211-1203223023121000-3002330020332022-3332002110313211-1101020312031333-3110002100033013"></a>

## Next pages — site_local / 203000220221 / 4

- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)

<a id="canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322100332122312-1120023202002300-3111102113003112-1022112112320132-3221310312033322-1332031121330330-1312033230202033-0231200030000201"></a>

## syslog — syslog / 101220323112 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- syslog

<a id="canonical-3123113221322213-1100032231313132-2203001020313022-2100203100001201-1311111121022111-3333020322312231-3220002033102212-3023013323133201"></a>

Type: `"object"`. single nested block, Optional.

Syslog Server Configuration. Configuration for syslog server.

Upstream description:

Configuration for syslog server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("tcp_server",
    "tls_server"),
  validators.ConflictingObjectAttributes("tcp_server",
    "udp_server"),
  validators.ConflictingObjectAttributes("tls_server",
    "udp_server")}
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
  "x-ves-oneof-field-format_choice": "[\"syslog_rfc5424\"]",
  "x-ves-oneof-field-mode_choice": "[\"tcp_server\",\"tls_server\",\"udp_server\"]"
}
```

Terraform syntax:

```terraform
syslog {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313100002001313-0132111310123100-1201223011110132-3301002223211031-3033003320301322-2200133103203212-2122201112320030-3322332003100023"></a>

## Direct properties — syslog / 101220323112 / 3

<a id="canonical-1131012122212021-2032330210012222-2312121110332133-0312210201133031-0131322213312232-1231201112020311-1032212310332230-0202212322330103"></a>

<a id="canonical-3310102122130301-1202100031330112-1213221133311130-3133000033220010-2302332322031323-0130020322103000-1201122321002303-1300002123110000"></a>

## syslog_rfc5424 property — syslog / 101220323112 / 4

Type: `"number"`. Optional.

Exclusive with \[\] Select RFC5424 syslog format and maximum message length.

Upstream description:

Exclusive with \[\] Select RFC5424 syslog format and maximum message length.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(408, 268435456),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 268435456,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 408
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "408",
    "ves.io.schema.rules.uint32.lte": "268435456"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "408",
    "ves.io.schema.rules.uint32.lte": "268435456"
  }
}
```

- [tcp_server](resources--log_receiver--reference--group-001.md#canonical-0101301221203233-1030333032122010-0122011331313322-1202311033113131-1213131201012121-2031303120132032-1303112231003003-0200003312202312): complete subsection reference.

- [tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132): complete subsection reference.

- [udp_server](resources--log_receiver--reference--group-001.md#canonical-0033222222322202-2110311020211201-1022101220030331-1003230221200312-3322222030111310-3110113212301301-3213331102312231-3020011200301320): complete subsection reference.

<a id="canonical-2030110100321123-1002330101103311-0312310012120213-0231022012211032-3310302333033111-1302122231220331-0110121223132301-1211120131332110"></a>

## Next pages — syslog / 101220323112 / 5

- [syslog.tcp_server](resources--log_receiver--reference--group-001.md#canonical-0101301221203233-1030333032122010-0122011331313322-1202311033113131-1213131201012121-2031303120132032-1303112231003003-0200003312202312)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132)
- [syslog.udp_server](resources--log_receiver--reference--group-001.md#canonical-0033222222322202-2110311020211201-1022101220030331-1003230221200312-3322222030111310-3110113212301301-3213331102312231-3020011200301320)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)

<a id="canonical-0101301221203233-1030333032122010-0122011331313322-1202311033113131-1213131201012121-2031303120132032-1303112231003003-0200003312202312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120102031231130-2001003220230100-0323213311000330-3013313330012230-3102121320030013-0030012220122120-3012230201101222-2321232332303130"></a>

## syslog.tcp_server — tcp_server / 230113100232 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- syslog.tcp_server

<a id="canonical-1211233201000232-2101222213211010-3302223112221221-3101203301112300-2111030223030121-2213132320232210-0100220110310302-1301021213331013"></a>

Type: `"object"`. single nested block, Optional.

TCP Server name and Port Number. Name and port number for a TCP server.

Upstream description:

Name and port number for a TCP server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("port",
    "server_name")}
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
tcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220321101100213-0220101321303111-0100312213200312-3001011201121322-2000333201203320-1033102023220001-1311030112233203-1303013133030023"></a>

## Direct properties — tcp_server / 230113100232 / 3

<a id="canonical-3110220212123210-0221013330001032-3100013102222223-1312131311302120-2002302012333022-0033112101101033-3032000301113311-0110133101003303"></a>

<a id="canonical-2110231121223123-1031331332321233-1213213021003122-3210022112123010-1303330203231000-0113111310233302-3121311221030202-2131132211103330"></a>

## port property — tcp_server / 230113100232 / 4

Type: `"number"`. Optional.

Port Number. Port number used for communication.

Upstream description:

Port number used for communication.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0013230223111130-0301001000133100-1100211201130012-3001333122203323-3022321231120023-0113112230132322-0210302321121200-2022231113002020"></a>

<a id="canonical-3210110023011330-3112021321202311-0303202331312303-3332112022001303-3200233131322320-3120101323011030-3020031101103322-3212021320103011"></a>

## server_name property — tcp_server / 230113100232 / 5

Type: `"string"`. Optional.

Server name is fully qualified domain name or IP address of the server.

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
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0221101000232010-0000233231011111-3312030130311120-0102120133313000-3031122211332220-3213133021111322-0322003210101033-0223033200103233"></a>

## Next pages — tcp_server / 230113100232 / 6

- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)

<a id="canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211100001331133-1112002301302010-3323302300123301-0201320011001101-3122130023203333-1020011221013220-1220001133033202-0222122101312003"></a>

## syslog.tls_server — tls_server / 300112011311 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- syslog.tls_server

<a id="canonical-0020023302001203-2233301310211022-3323120001303023-0213110012010310-0233031321013002-0001203023203232-2212233122213231-0001130311023320"></a>

Type: `"object"`. single nested block, Optional.

TLS config for client of discovery service.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("server_name"),
  validators.ConflictingObjectAttributes("default_https_port",
    "default_syslog_tls_port"),
  validators.ConflictingObjectAttributes("default_https_port",
    "port"),
  validators.ConflictingObjectAttributes("default_syslog_tls_port",
    "port"),
  validators.ConflictingObjectAttributes("mtls_disabled",
    "mtls_enable"),
  validators.ConflictingObjectAttributes("trusted_ca_url",
    "volterra_ca")}
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
  "x-ves-oneof-field-ca_choice": "[\"trusted_ca_url\",\"volterra_ca\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-port_choice": "[\"default_https_port\",\"default_syslog_tls_port\",\"port\"]"
}
```

Terraform syntax:

```terraform
tls_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203300332132323-2121133333031023-1100330120121130-3322121213001133-2200000332032123-2122030221123110-1210110001032031-1023002133021132"></a>

## Direct properties — tls_server / 300112011311 / 3

- [default_https_port](resources--log_receiver--reference--group-001.md#canonical-0232232230222332-3110010222321130-3220100133311033-2003212032000031-2113201123003231-1110300222022300-3112013022000011-0002113200201323): complete subsection reference.

- [default_syslog_tls_port](resources--log_receiver--reference--group-001.md#canonical-2113213030110131-1230233321312321-3132222001332230-1011110310311222-1223120113231112-3313122220330311-3013130232303211-3211302230223023): complete subsection reference.

- [mtls_disabled](resources--log_receiver--reference--group-001.md#canonical-2131320132203000-1012003011033012-0331200120113131-2313003030011020-2302030303123211-3020202232022032-1111333322210032-2113333333213310): complete subsection reference.

- [mtls_enable](resources--log_receiver--reference--group-001.md#canonical-3203021321120330-3110200201001212-0323122101231312-0302231322322102-1300112232231022-1313133123001000-0123202003310332-3312303322102112): complete subsection reference.

<a id="canonical-1221303200022120-1312023310211122-0203222210013100-1100312311130113-3120123103121221-1312312201233220-3021332001300301-1212120122322230"></a>

<a id="canonical-3301320100233120-1201203120320121-2200312032011000-1120332203021311-1030023221321123-0110231003010230-0032202132232332-1010103202300113"></a>

## port property — tls_server / 300112011311 / 4

Type: `"number"`. Optional.

Exclusive with \[default\_https\_port default\_syslog\_tls\_port\] Custom port number used for
communication.

Upstream description:

Exclusive with \[default\_https\_port default\_syslog\_tls\_port\] Custom port number used for
communication.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2013113212320010-1112002010003210-3111302323221101-3133112120132021-1121301032011101-3012210313120033-0201013132311203-3003002120203301"></a>

<a id="canonical-3212320223320310-0112333123113113-3023122230321201-3202303211133331-3012133331003320-0223313122021301-0133331133011213-2023000231011130"></a>

## server_name property — tls_server / 300112011311 / 5

Type: `"string"`. Optional.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against.

Upstream description:

ServerName is passed to the server for SNI and is used in the client to check server certificates
against.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0103013302112311-0201012230012001-0000213233113021-3213013211220101-3013321011323310-2002313213103230-2230130210312312-0000232002132103"></a>

<a id="canonical-2313313122022231-3300200231211322-3021201302030020-2312202212132031-0101011101111112-1233302222100230-3320312103330033-0020131320030333"></a>

## trusted_ca_url property — tls_server / 300112011311 / 6

Type: `"string"`. Optional.

Exclusive with \[volterra\_ca\] The URL or value for trusted Server CA certificate or certificate
chain Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[volterra\_ca\] The URL or value for trusted Server CA certificate or certificate
chain Certificates in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

- [volterra_ca](resources--log_receiver--reference--group-001.md#canonical-2131020222320222-0300320321011210-1013021323220103-0323223020032323-3212132220311230-0303111323221010-2001102231110201-2321302030313120): complete subsection reference.

<a id="canonical-2112211023302231-2323030333211303-2201101230120221-1131022013120230-0313203200233313-1102010001032310-1232001133131103-0310021212112202"></a>

## Next pages — tls_server / 300112011311 / 7

- [syslog.tls_server.default_https_port](resources--log_receiver--reference--group-001.md#canonical-0232232230222332-3110010222321130-3220100133311033-2003212032000031-2113201123003231-1110300222022300-3112013022000011-0002113200201323)
- [syslog.tls_server.default_syslog_tls_port](resources--log_receiver--reference--group-001.md#canonical-2113213030110131-1230233321312321-3132222001332230-1011110310311222-1223120113231112-3313122220330311-3013130232303211-3211302230223023)
- [syslog.tls_server.mtls_disabled](resources--log_receiver--reference--group-001.md#canonical-2131320132203000-1012003011033012-0331200120113131-2313003030011020-2302030303123211-3020202232022032-1111333322210032-2113333333213310)
- [syslog.tls_server.mtls_enable](resources--log_receiver--reference--group-001.md#canonical-3203021321120330-3110200201001212-0323122101231312-0302231322322102-1300112232231022-1313133123001000-0123202003310332-3312303322102112)
- [syslog.tls_server.volterra_ca](resources--log_receiver--reference--group-001.md#canonical-2131020222320222-0300320321011210-1013021323220103-0323223020032323-3212132220311230-0303111323221010-2001102231110201-2321302030313120)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)

<a id="canonical-0232232230222332-3110010222321130-3220100133311033-2003212032000031-2113201123003231-1110300222022300-3112013022000011-0002113200201323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131330020220122-0110121311232031-3231322003302312-0133333212003330-2031010220320322-3131021313303121-0302320010103120-0033032211322113"></a>

## syslog.tls_server.default_https_port — default_https_port / 223002010303 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132)
- syslog.tls_server.default_https_port

<a id="canonical-1012312013211300-3221032102101313-1003223132312223-3123132220120021-0000011021212121-2311200012233231-3332210102230023-3313121320101323"></a>

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
default_https_port = {}
```

<a id="canonical-0022000303113123-0320132212033012-3321210023303021-3102200211201121-3031201121102203-0112223311020011-1123123120311100-3301020010023032"></a>

## Direct properties — default_https_port / 223002010303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331330122032110-3000232313130202-0021321220230010-3200322130111301-0022112110011222-3301020303013233-3031223133011233-2133102231012012"></a>

## Next pages — default_https_port / 223002010303 / 4

- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)

<a id="canonical-2113213030110131-1230233321312321-3132222001332230-1011110310311222-1223120113231112-3313122220330311-3013130232303211-3211302230223023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320031213031321-0321200023000203-0013000032131311-2130320101001011-2301123222130331-2321112122203122-1200121211311011-3333232011233032"></a>

## syslog.tls_server.default_syslog_tls_port — default_syslog_tls_port / 300002112122 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132)
- syslog.tls_server.default_syslog_tls_port

<a id="canonical-3222313010303000-3032212132132020-0213333011303011-2031233303013011-0000223132011120-1122320112210132-0102130300122230-2122010321300313"></a>

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
default_syslog_tls_port = {}
```

<a id="canonical-3121003112023233-1122200311232120-3022002303033112-0023302120033022-3110012232031213-1001131021202231-3330330121022210-2320001103102213"></a>

## Direct properties — default_syslog_tls_port / 300002112122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121310313001000-1112331212202212-1302323230212211-1101210021112003-1301103021133133-0030223003112103-0301121002312200-0320322123113032"></a>

## Next pages — default_syslog_tls_port / 300002112122 / 4

- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)

<a id="canonical-2131320132203000-1012003011033012-0331200120113131-2313003030011020-2302030303123211-3020202232022032-1111333322210032-2113333333213310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121122310323223-0210100311130220-1313100301013311-3132323302100103-0221322203011121-2330103111103200-3312220330120110-1231032301333011"></a>

## syslog.tls_server.mtls_disabled — mtls_disabled / 230112323311 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132)
- syslog.tls_server.mtls_disabled

<a id="canonical-3331130130002302-1020301230023101-3232300211021321-2321313021102023-1130100010210033-0112002021232102-0211200333333023-3223133210230331"></a>

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
mtls_disabled = {}
```

<a id="canonical-1121213130312203-0302000233221002-3012110002220123-3000102233211212-0331032331302220-1000322201311133-1002000221311000-0210332020210310"></a>

## Direct properties — mtls_disabled / 230112323311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012103232103100-0022231323320011-1200122033320200-0112031212131303-1101113302331312-2000310332001310-3233100300210100-0321303021312323"></a>

## Next pages — mtls_disabled / 230112323311 / 4

- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)

<a id="canonical-3203021321120330-3110200201001212-0323122101231312-0302231322322102-1300112232231022-1313133123001000-0123202003310332-3312303322102112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231133332212203-2010113033131112-2211332223131331-1133011320023113-1122122310012312-2110022232212001-3012111120322121-2103000032103201"></a>

## syslog.tls_server.mtls_enable — mtls_enable / 212100210012 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132)
- syslog.tls_server.mtls_enable

<a id="canonical-0332023303013301-3022011103203220-0201120130313020-0213011220232223-2012130132231112-0033333032011030-1020311232222012-3232223021103103"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for mtls enable.

Upstream description:

TLS config for client.

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
mtls_enable {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131202012000121-1131233013003211-2002121102330001-2121331031103130-0203001010100002-2211013301300302-3023322302111020-3133000200013320"></a>

## Direct properties — mtls_enable / 212100210012 / 3

<a id="canonical-3110101130011313-2012102331223221-0212112033223311-1132123202032002-2321202031323211-3313121013302022-1332230112230103-0131211233113202"></a>

<a id="canonical-1200120002210021-0123021103102012-3033202003111131-0011333311002003-1233311201131221-3103133312202030-2221321212111033-0210112211032132"></a>

## certificate property — mtls_enable / 212100210012 / 4

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

- [key_url](resources--log_receiver--reference--group-001.md#canonical-3322210101013313-1312220332330211-2001102110333313-2202202111133301-1302332320331112-0231222021202210-2313322201113023-1213300111211310): complete subsection reference.

<a id="canonical-0112322010103301-2130202323103000-1332020133131311-3302231203020033-0200312012221002-2200020021220133-1233103033222102-2201020213333132"></a>

## Next pages — mtls_enable / 212100210012 / 5

- [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--reference--group-001.md#canonical-3322210101013313-1312220332330211-2001102110333313-2202202111133301-1302332320331112-0231222021202210-2313322201113023-1213300111211310)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)

<a id="canonical-3322210101013313-1312220332330211-2001102110333313-2202202111133301-1302332320331112-0231222021202210-2313322201113023-1213300111211310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123213221023010-2030121130231210-3123123211033303-1230121231110223-2323123210311302-0031333301110102-1330330323203123-3300323221302100"></a>

## syslog.tls_server.mtls_enable.key_url — key_url / 112233323312 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132)
- [syslog.tls_server.mtls_enable](resources--log_receiver--reference--group-001.md#canonical-3203021321120330-3110200201001212-0323122101231312-0302231322322102-1300112232231022-1313133123001000-0123202003310332-3312303322102112)
- syslog.tls_server.mtls_enable.key_url

<a id="canonical-2021123211000111-3121133203232103-0313321201123002-2310032313202013-1333033221201223-2113233131202132-0201212210103102-2301023332011003"></a>

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
key_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220220220113233-1120021013013013-1333231200223212-1213310021103010-3232231101101231-1331121003212312-0113033311213031-2131321100230010"></a>

## Direct properties — key_url / 112233323312 / 3

- [blindfold_secret_info](resources--log_receiver--reference--group-001.md#canonical-1300012100320123-1210333223122133-1102203133020331-1333010231030200-2002311033311223-0312100230330301-3313012210332110-3100222223011212): complete subsection reference.

- [clear_secret_info](resources--log_receiver--reference--group-001.md#canonical-3221333010101232-2013013200311303-2131203203320233-3131302201132321-3330332133023110-1233200230101300-1033133203230202-1010210222303110): complete subsection reference.

<a id="canonical-0132011012201010-2030322231310300-2330130202212011-3132312100303101-0123200031032022-0210131022223010-0113030021321132-1302333122132300"></a>

## Next pages — key_url / 112233323312 / 4

- [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info](resources--log_receiver--reference--group-001.md#canonical-1300012100320123-1210333223122133-1102203133020331-1333010231030200-2002311033311223-0312100230330301-3313012210332110-3100222223011212)
- [syslog.tls_server.mtls_enable.key_url.clear_secret_info](resources--log_receiver--reference--group-001.md#canonical-3221333010101232-2013013200311303-2131203203320233-3131302201132321-3330332133023110-1233200230101300-1033133203230202-1010210222303110)
- [syslog.tls_server.mtls_enable](resources--log_receiver--reference--group-001.md#canonical-3203021321120330-3110200201001212-0323122101231312-0302231322322102-1300112232231022-1313133123001000-0123202003310332-3312303322102112)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)

<a id="canonical-1300012100320123-1210333223122133-1102203133020331-1333010231030200-2002311033311223-0312100230330301-3313012210332110-3100222223011212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300030330121313-3001300122102030-2320020221222000-2000303103122220-0301030032203220-2233322120133000-2121033130201301-1321333300313223"></a>

## syslog.tls_server.mtls_enable.key_url.blindfold_secret_info — blindfold_secret_info / 112022130321 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132)
- [syslog.tls_server.mtls_enable](resources--log_receiver--reference--group-001.md#canonical-3203021321120330-3110200201001212-0323122101231312-0302231322322102-1300112232231022-1313133123001000-0123202003310332-3312303322102112)
- [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--reference--group-001.md#canonical-3322210101013313-1312220332330211-2001102110333313-2202202111133301-1302332320331112-0231222021202210-2313322201113023-1213300111211310)
- syslog.tls_server.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-3030122031120301-1031330013211113-1302330102121110-2113332132311130-2010223001233122-3112032000230221-2113303300033310-0213003323202111"></a>

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

<a id="canonical-0000232322233000-1113321212012230-1210011012110003-2120030212013112-3010122223201300-0132013312230202-1031112101130231-2223111101002232"></a>

## Direct properties — blindfold_secret_info / 112022130321 / 3

<a id="canonical-1103110020131000-1010032231012332-0131231303110302-1130310231110210-2332230313202010-3022020023111331-2031001312210210-3100212133210001"></a>

<a id="canonical-3030333203101031-2232131300210311-2102110110323231-2222130220111003-0033113111223200-1220001321132200-3230321021312130-0033213011101331"></a>

## decryption_provider property — blindfold_secret_info / 112022130321 / 4

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

<a id="canonical-2123120103003103-1323210201222321-0303312123220233-1123210122320012-3110003030311201-0213122000322301-1302030103202311-2131323112200021"></a>

<a id="canonical-3033231012121321-1330100003023000-3001131310302201-0302310330130230-3320230322300100-3210213200102020-0203130113013223-0210001312302003"></a>

## location property — blindfold_secret_info / 112022130321 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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

<a id="canonical-3120211021013201-3303211003220013-3110320201221301-0213022001211223-1300332312032202-3202221103311332-2020032231103311-1010332312330300"></a>

<a id="canonical-3132120212220002-2222300303330331-1233103233223032-2312311200023333-3133111002203121-1303232303323232-3120312011000203-2210330321120203"></a>

## store_provider property — blindfold_secret_info / 112022130321 / 6

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

<a id="canonical-0030030110233000-3220012212310332-0002001002031331-1131011000310020-2033302310213033-0302131113010200-2333211220222033-2201233213102310"></a>

## Next pages — blindfold_secret_info / 112022130321 / 7

- [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--reference--group-001.md#canonical-3322210101013313-1312220332330211-2001102110333313-2202202111133301-1302332320331112-0231222021202210-2313322201113023-1213300111211310)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)

<a id="canonical-3221333010101232-2013013200311303-2131203203320233-3131302201132321-3330332133023110-1233200230101300-1033133203230202-1010210222303110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332302021131000-3300022022130232-0101101210121111-3111221010111013-2101132221311113-0013303023121122-1030202322000231-1223021113000222"></a>

## syslog.tls_server.mtls_enable.key_url.clear_secret_info — clear_secret_info / 001012123122 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132)
- [syslog.tls_server.mtls_enable](resources--log_receiver--reference--group-001.md#canonical-3203021321120330-3110200201001212-0323122101231312-0302231322322102-1300112232231022-1313133123001000-0123202003310332-3312303322102112)
- [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--reference--group-001.md#canonical-3322210101013313-1312220332330211-2001102110333313-2202202111133301-1302332320331112-0231222021202210-2313322201113023-1213300111211310)
- syslog.tls_server.mtls_enable.key_url.clear_secret_info

<a id="canonical-0103103311310111-2223330012323321-1133110012202032-2003233310012301-2001013110032122-1323131210310302-3201032202302330-0023020013111213"></a>

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

<a id="canonical-3201021110202232-0323203221111222-0312222110312102-3301201131320121-0303130312220111-3213221000300112-2221031212211021-1121022032313030"></a>

## Direct properties — clear_secret_info / 001012123122 / 3

<a id="canonical-2201030130030210-2300231121322121-1312121112111311-3101110231313210-1313120123213331-3213332211032232-3302322200300030-2012301110332121"></a>

<a id="canonical-3310003302010021-3003201221300331-2211332130111112-3311303222233303-2003100313233012-2303133132113321-0023110002300233-3202111002021023"></a>

## provider_ref property — clear_secret_info / 001012123122 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2320302103121113-2131302322311203-3202121203200120-0203123120303103-3230230331303312-3011313102003120-0012001013033013-3302323003133330"></a>

<a id="canonical-2122201220312103-3210003300210210-3221300330011313-3000022031110011-0031321312121223-0100123111110133-3002003212202003-3113002101311211"></a>

## URL property — clear_secret_info / 001012123122 / 5

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

<a id="canonical-3332302321323322-1033002211123121-1030330201022311-2033121333132100-1201032102003010-3233032310322011-1210002331213100-2302122030303200"></a>

## Next pages — clear_secret_info / 001012123122 / 6

- [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--reference--group-001.md#canonical-3322210101013313-1312220332330211-2001102110333313-2202202111133301-1302332320331112-0231222021202210-2313322201113023-1213300111211310)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)

<a id="canonical-2131020222320222-0300320321011210-1013021323220103-0323223020032323-3212132220311230-0303111323221010-2001102231110201-2321302030313120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222133220000032-0303120201000212-3030302213021000-3330221002232223-2300303231003301-1220100321023311-3101323100133030-0310231120311310"></a>

## syslog.tls_server.volterra_ca — volterra_ca / 232133212210 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132)
- syslog.tls_server.volterra_ca

<a id="canonical-0330202220100300-0333210333133310-0322012101100011-0302201013023012-3113132203101220-3331223211112200-3133331100101311-0033302110220300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra ca.

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
volterra_ca = {}
```

<a id="canonical-2320131113123210-3213132321121133-0212210130001031-1103000111232300-0222230033332200-1312003303121031-1200103120103132-0220003012023003"></a>

## Direct properties — volterra_ca / 232133212210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033333211103300-2202330122332210-1310320203220003-1230300122301203-2303111210302321-3223100113200213-0320312100021033-1001303111133123"></a>

## Next pages — volterra_ca / 232133212210 / 4

- [syslog.tls_server](resources--log_receiver--reference--group-001.md#canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)

<a id="canonical-0033222222322202-2110311020211201-1022101220030331-1003230221200312-3322222030111310-3110113212301301-3213331102312231-3020011200301320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311101020031330-1120033212212022-3321130312222302-3003000321320223-0232100332201001-1332313000232321-2100123212013012-0303201331323320"></a>

## syslog.udp_server — udp_server / 132013101122 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- syslog.udp_server

<a id="canonical-3222232131211113-2312132322233112-1301323311131000-0210121312101002-2321331131123212-0112132132133103-1232132213103032-1033102103130313"></a>

Type: `"object"`. single nested block, Optional.

UDP Server Name and Port Number. Name and port number for a UDP server.

Upstream description:

Name and port number for a UDP server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("port",
    "server_name")}
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
udp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130301212013202-3230332323122312-3211001103103302-1023201333303111-1313331030110220-0121200123021320-3120201031303023-2110132032230212"></a>

## Direct properties — udp_server / 132013101122 / 3

<a id="canonical-0233031211203031-2032213010100022-0110100320230312-0323131130022221-2113003331103211-2331122000332110-3202223322123302-2001020023320302"></a>

<a id="canonical-0212111210020230-2102113013300031-0312332220311123-3010211232202230-1302302120101132-3330233201110022-0231230333010222-3231333112311001"></a>

## port property — udp_server / 132013101122 / 4

Type: `"number"`. Optional.

Port Number. Port number used for communication.

Upstream description:

Port number used for communication.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1130121111032103-1112230102331200-1221310320301013-2221001130023001-2210101213122310-3130110322232031-3012013212211231-1301020320313121"></a>

<a id="canonical-0123223121210120-1310112320300321-2130300010332311-2211301231203321-0032222102230111-2210012312312110-3113323201012310-3101023320013123"></a>

## server_name property — udp_server / 132013101122 / 5

Type: `"string"`. Optional.

Server name is fully qualified domain name or IP address of the server.

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
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2011103121233310-0212302131100113-1030331133320323-1010031012100102-0303201020033020-2201120320010022-0300200213101323-2333011213333012"></a>

## Next pages — udp_server / 132013101122 / 6

- [syslog](resources--log_receiver--reference--group-001.md#canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)

<a id="canonical-2130131322321001-1002201220120233-2233232321020120-3201102003221011-2022123000123110-2333132333113230-2022223022200022-1032320120302321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001021033322331-0212122302022130-2113311003203203-3033012312320300-1221313200232013-1302011121230202-3133312213110222-1203030202013100"></a>

## timeouts — timeouts / 300202321122 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- timeouts

<a id="canonical-2002112103220323-2002111103030222-1203320302331133-0321223322200222-3133300230101001-2003201333102022-2102112111303000-3033332331323203"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0233112122200100-0031122313013121-0203302231111201-3031233303332030-0310233232132210-3012012011100311-1202101030321132-2313310000300230"></a>

## Direct properties — timeouts / 300202321122 / 3

<a id="canonical-1030310022213212-1111320202303331-1332010323300102-0103322201033203-0013330233222231-0121033231301022-0233302102322113-3230230223122302"></a>

<a id="canonical-0123201210101220-0232022212102302-3011021233211333-3132212313113231-2233210221332023-2222230132121112-3310332231223100-2113113103331020"></a>

## create property — timeouts / 300202321122 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2030323000223333-2321133110013011-2102310002312111-2200230333001021-2210220203022302-0221233213103133-1222000003033323-2301011101130123"></a>

<a id="canonical-0220212033012003-3320231130312202-0310101232002110-1111210132230233-3103133233020130-2212102111201302-1321102031133121-2303131323331032"></a>

## delete property — timeouts / 300202321122 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2123101212231330-2022002202300212-2320123332232231-0003233001312230-2310003120010230-1112223313222101-3320320201203203-3003323033210002"></a>

<a id="canonical-2300301203033132-1111131002110123-2320220101020220-2330322110103132-3001220032301030-0113120230023012-0110023302122302-3212013233201313"></a>

## read property — timeouts / 300202321122 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0333131231213231-2330013011122022-0321001100002122-2002320301200331-3232013231331202-2322301312230322-0113132001232020-0102111222100203"></a>

<a id="canonical-0301230001123330-3131011033000212-2133100031033330-3032132110311012-0333310033233220-2313300110020332-3333331012201101-1232123303110131"></a>

## update property — timeouts / 300202321122 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1232310321201101-1323002302002203-1100322213302303-3120002031333323-2012102110021100-0221320000013313-0130332231020320-2121003200200213"></a>

## Next pages — timeouts / 300202321122 / 8

- [Property reference](resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
