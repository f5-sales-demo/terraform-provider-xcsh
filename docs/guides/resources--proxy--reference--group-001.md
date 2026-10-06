---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- Property reference

<a id="canonical-1001313002231110-3220210001330010-2201003330030101-1103212112230101-0103013312003021-1303331302220102-0210002212031223-0211233303130113"></a>

### Direct properties for `xcsh_proxy`

- [active_forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-0003323102113131-0103011023111231-1113001332012010-0301232233010331-0130302132021322-2321100013200013-0033323321132110-1121322121220021): complete subsection reference.

<a id="canonical-1113220030033300-2330301012113223-0032201012033220-3002100210302020-2002133110110000-0000102110023102-2001033221130123-0222312130120312"></a>

<a id="canonical-3203120200113301-0110212123221121-2000330232221101-1333233001300300-2301023221110032-1211013201112101-2201102330000122-2210021323300330"></a>

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

<a id="canonical-1232303023013110-3213111322332101-3212310220203220-2320102011303201-1323131312201012-1100122022213121-1311132023300301-2012102203331302"></a>

<a id="canonical-2300322102122220-1121333220220302-0313132203212323-2303211001312002-1023012003120100-0002022200220030-3213121202220302-2312333221033132"></a>

#### `connection_timeout` property

Type: `"number"`. Optional, Computed.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Additional upstream details:

The default value is 2000 (2 seconds)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(1800000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1800000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

<a id="canonical-1032132031200233-1232001111222113-0120310031313220-0221213200123113-3130103222213100-0321200110323110-3122212310232313-1313201331032203"></a>

<a id="canonical-3002023000111112-3200233311122110-1222000032203312-2012010311012332-1233123001030030-1233133320113032-3111211203302212-3203321211113210"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3123113223330000-0321223321230023-2102220003020112-0133301123032313-1133112230303112-1331121110210003-1130333100122213-2310330113222101"></a>

<a id="canonical-3213111200211010-3120301221323231-3212333102222323-2301121232013021-0011223232311013-3213223312310210-0030113312021313-1003332130320231"></a>

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

- [do_not_advertise](resources--proxy--reference--group-001.md#canonical-1300300333220210-2321100332123100-3100222323020132-2000023232233210-0122031023212311-0203230120332013-3020033130003013-3323122123123232): complete subsection reference.

- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330): complete subsection reference.

- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331): complete subsection reference.

<a id="canonical-3021310010301023-0303200031113011-3232020220213113-2313303000232303-0321312220212333-1232233203313202-1011313201333211-2331300002330210"></a>

<a id="canonical-0211232202023103-2031132221001200-0122123101323330-3233332122033112-1323100102003113-2233102101223313-2221301301211333-1300112002312121"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1232113033323320-1321330113300230-0330200023130232-3303322211233302-2201031330123101-2013031103223033-3130300121311333-2302011131113112"></a>

<a id="canonical-3002012102031130-2233330300002120-2202010010120223-0122333301332322-3122321202321232-3221230020021322-2013301111330210-3311222012310133"></a>

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

<a id="canonical-0113111221131320-3121313220033313-0220213200023113-1212300333123000-2133103200102203-2133102333220320-2113223333133332-1212032200200110"></a>

<a id="canonical-2311030211333332-1220031311020123-2333333021003310-0020030112123232-3000331330301332-2210100330311010-1010032313000000-3133331132012102"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Proxy. Must be unique within the namespace.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0011102331302022-0032213200233212-3102011000121310-3011130113223113-3032031001100002-0102222132032110-3111010021020033-3123112122111210"></a>

<a id="canonical-1322013302112032-2332312332233210-3322333113302133-3302311110303022-3300031111111002-2313001233332113-1020131310303010-1322103300203310"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Proxy is created.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [no_forward_proxy_policy](resources--proxy--reference--group-005.md#canonical-3333131321121110-3221311322032300-1111103312333233-0321012201211000-1102333201111020-0020303330301101-2313211321302213-0221203120310222): complete subsection reference.

- [no_interception](resources--proxy--reference--group-005.md#canonical-3001311021210113-1203322120301211-2003112333102203-1131010301220310-1212123211110331-3201330202033133-1001002102301300-0023323331213323): complete subsection reference.

- [site_local_inside_network](resources--proxy--reference--group-005.md#canonical-2300220130222220-0321012100023110-3312120100301031-2332310123110321-0222103331130303-0122320111223322-2100122213201320-3222032202331000): complete subsection reference.

- [site_local_network](resources--proxy--reference--group-005.md#canonical-1303132001322003-3222233301002011-1312320023221130-0032320032223121-2231221200221010-2211212211123222-1202133222332122-0032211113003013): complete subsection reference.

- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302): complete subsection reference.

- [timeouts](resources--proxy--reference--group-005.md#canonical-2113110122012210-0300103202002221-2012320121222331-2230022312003312-3203002101213002-1312002203112032-1202231032011220-3312233322001132): complete subsection reference.

- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031): complete subsection reference.

<a id="canonical-2320232230211110-2000110222131012-2010210010323113-2032001132112200-0112102321101300-1123202021101012-3322320312010023-3020220220212000"></a>

### All schema paths for `xcsh_proxy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_forward_proxy_policies` | [active_forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-0221211131033121-3202012010201332-1113103330200011-0102030311133213-2111131301330011-2301322012102332-2312002132210112-2020233302023120) |
| `active_forward_proxy_policies.forward_proxy_policies` | [active_forward_proxy_policies.forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-2023203231230201-1323322021130311-2122030003302322-1112013201113131-1113001310002102-2301303302221010-3220122013332323-1131233022222112) |
| `active_forward_proxy_policies.forward_proxy_policies.name` | [active_forward_proxy_policies.forward_proxy_policies.name](resources--proxy--reference--group-001.md#canonical-0122323000223013-2313013011321230-0113223123021133-2322023301100030-1333220110231231-2000220323123311-3311333223201322-0111313031013223) |
| `active_forward_proxy_policies.forward_proxy_policies.namespace` | [active_forward_proxy_policies.forward_proxy_policies.namespace](resources--proxy--reference--group-001.md#canonical-2233210132011010-1312102310211131-2210033133200020-0232002212120023-3101311001132121-1003220003230002-0100201300121203-0002231320332323) |
| `active_forward_proxy_policies.forward_proxy_policies.tenant` | [active_forward_proxy_policies.forward_proxy_policies.tenant](resources--proxy--reference--group-001.md#canonical-1203123310333233-3001103002333323-3312300302233220-0033202323122113-1012332321332231-1231020100123130-3030000221032013-1123033321102313) |
| `annotations` | [annotations](resources--proxy--reference--group-001.md#canonical-1113220030033300-2330301012113223-0032201012033220-3002100210302020-2002133110110000-0000102110023102-2001033221130123-0222312130120312) |
| `connection_timeout` | [connection_timeout](resources--proxy--reference--group-001.md#canonical-1232303023013110-3213111322332101-3212310220203220-2320102011303201-1323131312201012-1100122022213121-1311132023300301-2012102203331302) |
| `description` | [description](resources--proxy--reference--group-001.md#canonical-1032132031200233-1232001111222113-0120310031313220-0221213200123113-3130103222213100-0321200110323110-3122212310232313-1313201331032203) |
| `disable` | [disable](resources--proxy--reference--group-001.md#canonical-3123113223330000-0321223321230023-2102220003020112-0133301123032313-1133112230303112-1331121110210003-1130333100122213-2310330113222101) |
| `do_not_advertise` | [do_not_advertise](resources--proxy--reference--group-001.md#canonical-0312133310103133-2021121111223230-0202102102001031-1211202103032101-3121123313123212-0031113301200102-0123101311303022-3333312123131010) |
| `dynamic_proxy` | [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-0322013013112103-1110312232333000-0001202300202230-1130120121002131-2033132100020232-0121002303211031-1313302121010231-1120332102001132) |
| `dynamic_proxy.disable_dns_masquerade` | [dynamic_proxy.disable_dns_masquerade](resources--proxy--reference--group-001.md#canonical-1332321312013121-3220123333011333-3221202222311303-2112202322103312-0312330113003310-3121130303313003-1130030300221313-0001033032022133) |
| `dynamic_proxy.domains` | [dynamic_proxy.domains](resources--proxy--reference--group-001.md#canonical-2330001000121003-1313222001030002-3222231123332003-1000131020332023-1330020320110022-0232003030231202-1102133333313000-1303210230001010) |
| `dynamic_proxy.enable_dns_masquerade` | [dynamic_proxy.enable_dns_masquerade](resources--proxy--reference--group-001.md#canonical-0320100121210033-0212112313013000-2020201032102121-1120230203133112-1303120032101310-2223323331212223-2111212322301113-0320101223330302) |
| `dynamic_proxy.http_proxy` | [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-0030322113101110-2112322112022013-3132302301210231-1201130310013303-1311221321210103-0002103231302132-0001202001113232-2002130001233222) |
| `dynamic_proxy.http_proxy.more_option` | [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-2012331230003322-1110102322230203-3230302030120130-0300223100313302-1332013220321120-2203020012201113-2022121111103102-1302030003300030) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy` | [dynamic_proxy.http_proxy.more_option.buffer_policy](resources--proxy--reference--group-001.md#canonical-2100310220321112-3301211120113113-2022003303013332-2211132133233021-1131001132202222-1003212130210031-2013010200012030-3003232333302203) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy.disabled` | [dynamic_proxy.http_proxy.more_option.buffer_policy.disabled](resources--proxy--reference--group-001.md#canonical-1312211323203300-2303101221313033-3103113102032222-2300100300131210-2101201221320213-1232032220000002-1113232200202032-3132210311103211) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy.max_request_bytes` | [dynamic_proxy.http_proxy.more_option.buffer_policy.max_request_bytes](resources--proxy--reference--group-001.md#canonical-2221131131333130-2222313200030313-0313031123032023-3012102102113210-3002201023322331-2320211013330232-0301110121110113-1213030100003233) |
| `dynamic_proxy.http_proxy.more_option.compression_params` | [dynamic_proxy.http_proxy.more_option.compression_params](resources--proxy--reference--group-001.md#canonical-0201033213231220-3001033300230203-3311012231233221-1001310312300023-3201210022303231-1331122112113032-2310333221020101-2313320231022123) |
| `dynamic_proxy.http_proxy.more_option.compression_params.content_length` | [dynamic_proxy.http_proxy.more_option.compression_params.content_length](resources--proxy--reference--group-001.md#canonical-0122313201213300-1103023220101303-0230023313000303-2232130120122300-2230302030122132-1210022313113323-0011210130110323-3231003203212022) |
| `dynamic_proxy.http_proxy.more_option.compression_params.content_type` | [dynamic_proxy.http_proxy.more_option.compression_params.content_type](resources--proxy--reference--group-001.md#canonical-3111221020223312-1031303020101233-2313233320203312-3311303100331003-3100000212011003-1101300321200312-2312223000023103-0212232031022232) |
| `dynamic_proxy.http_proxy.more_option.compression_params.disable_on_etag_header` | [dynamic_proxy.http_proxy.more_option.compression_params.disable_on_etag_header](resources--proxy--reference--group-001.md#canonical-3022032330103333-3222310003200332-1333130302122221-3330303013023103-1212232101120022-1232110212000313-2013003120233103-2310213311300032) |
| `dynamic_proxy.http_proxy.more_option.compression_params.remove_accept_encoding_header` | [dynamic_proxy.http_proxy.more_option.compression_params.remove_accept_encoding_header](resources--proxy--reference--group-001.md#canonical-1300312220111301-0001012121333103-3333032212123002-1221031330212002-3232312011233033-2323222232102232-3221213011333023-2303231131030001) |
| `dynamic_proxy.http_proxy.more_option.custom_errors` | [dynamic_proxy.http_proxy.more_option.custom_errors](resources--proxy--reference--group-001.md#canonical-0003033023110332-3200232020113131-3131332333111313-3333031011302003-2102000313130223-0130330013033313-0133031120320333-3320233310002111) |
| `dynamic_proxy.http_proxy.more_option.disable_default_error_pages` | [dynamic_proxy.http_proxy.more_option.disable_default_error_pages](resources--proxy--reference--group-001.md#canonical-0301311113200030-0232132021033021-0011102023320012-3003110313222210-0130023233223220-3121302033301221-1300132300201010-0301011320001202) |
| `dynamic_proxy.http_proxy.more_option.disable_path_normalize` | [dynamic_proxy.http_proxy.more_option.disable_path_normalize](resources--proxy--reference--group-001.md#canonical-1101100113013000-3120132013032211-0323202301031013-3223022122103200-1230221210303211-1021321311132113-2132132123011200-1221003200011230) |
| `dynamic_proxy.http_proxy.more_option.enable_path_normalize` | [dynamic_proxy.http_proxy.more_option.enable_path_normalize](resources--proxy--reference--group-001.md#canonical-0202131213310023-2122123023333221-2331200130200303-0101010202222232-0121301303132332-0313332103321311-2110330023311122-0112120101332101) |
| `dynamic_proxy.http_proxy.more_option.idle_timeout` | [dynamic_proxy.http_proxy.more_option.idle_timeout](resources--proxy--reference--group-001.md#canonical-2210031201222033-1202003032122010-3021013311013113-0103031002202011-2210111100012321-1230303011301310-0230303013203130-0133311133303212) |
| `dynamic_proxy.http_proxy.more_option.max_request_header_size` | [dynamic_proxy.http_proxy.more_option.max_request_header_size](resources--proxy--reference--group-001.md#canonical-1002223310210022-3021220301132011-0220121332012130-3310331010310113-1221220022201210-0313103131222302-0323221231003021-2232231011031121) |
| `dynamic_proxy.http_proxy.more_option.max_requests_per_connection` | [dynamic_proxy.http_proxy.more_option.max_requests_per_connection](resources--proxy--reference--group-001.md#canonical-0300220120123310-3222131032332031-1023332201020312-2101030322212001-1000131201232201-0010320221131221-3122121013032013-2002213322133102) |
| `dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection` | [dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection](resources--proxy--reference--group-001.md#canonical-1231131032321020-1202100030212321-0233122210001010-3330121031203033-1013111022310321-1301113010030220-3021023033330321-1013011320020103) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-001.md#canonical-2022222312021313-2013221003321222-2013222203101023-2113020310032003-2030110132113101-3323330230011010-0121023003330101-3221103230332011) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.name` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.name](resources--proxy--reference--group-001.md#canonical-0210220333113030-2111230002101012-1211131102330133-1303011123120312-1331121033302021-2100122032212322-3302213131223113-2321301003213010) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.overwrite` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.overwrite](resources--proxy--reference--group-001.md#canonical-2301203233021130-1222003031120030-3120332323101333-2002001112222011-3310300111321110-1112122302303203-3331020303221210-3301313020331111) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-0001301310201112-3222103103101101-2020122030230212-3313120311010010-3331210011130011-0132302031130302-2112301113003330-3303011221023112) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-3223110031103030-0322023113110213-3212113131221203-1221233231310202-2203122131232030-1323311101110031-0111133101330202-1320120031332222) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-002.md#canonical-1310103132001013-0331220123000320-0232300020230321-3213032021103012-0311233033323311-1230223230321222-3012130333300232-0311233121231330) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-002.md#canonical-3111021230131311-0101211132100312-1031031200021302-0233123001130010-3212000121122010-3211312231132133-2222233100312000-0020131122311011) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-002.md#canonical-1132322330310030-1332020103201012-2233302120312300-2203032130001102-2133012312020332-0003310130201222-0230321312132221-2001101013113210) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-2221300200230203-1103230202232212-1120100321231101-3103120231310311-3002322311120212-3230100002333023-0003131213111120-1131313112321032) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-002.md#canonical-1213310200312330-3303111332132013-0220312223110131-3023102111320121-3212212202313033-3323330101121120-0221012313032322-2333023133211132) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-002.md#canonical-2330132102030303-2122021321132333-0003131032302103-2231033323102211-2021301013300031-2133133120222330-1130300332313022-0333332010332121) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.value` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.value](resources--proxy--reference--group-001.md#canonical-3232110311003230-0012032100322101-2211203202321130-2111330000232312-2223120323113212-0331301122021120-2112213122030232-0313101330003120) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_remove` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_remove](resources--proxy--reference--group-001.md#canonical-3232313213010203-3130310023232323-1122332030333330-0022222203323323-2023330101023222-3223330233102220-0211301023330330-1221002213222331) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-2202131210002333-2223311202330033-3012001110202212-2302210331120322-0310302231110031-0222323111211312-1221113131000113-1001120301203112) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.append` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.append](resources--proxy--reference--group-002.md#canonical-2210013321203123-1223301312033202-1322303202011203-3101030303300022-2132232022120202-3331220333313313-0311013003021131-0132023130212330) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.name` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.name](resources--proxy--reference--group-002.md#canonical-0303322301202131-3322330220302012-1330123000223011-2313230210003230-1122303022202211-3222010003121111-2303032132032333-2202032122322020) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-0102211321132033-3022303033333203-3220230332101230-1322303331211320-0210112321202113-0202113030232010-3020030230333320-3023122321131022) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-1131302230310321-3001010113301103-0310322123021312-0132023220201031-0003302130123001-0312132210133300-2122322333300002-3230300101213121) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-002.md#canonical-3313331002100310-3203222233132122-3022322110130231-1230222333223010-1031233332333101-0101221211000132-2003220122310110-2331122102322032) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-002.md#canonical-1221001132103121-3012301331013310-1301111213032120-3333203212023022-3130102021220230-2130110101333013-0003233322012221-3023102221100110) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-002.md#canonical-1030331203321001-1331233211101013-1101030123021312-1232322231311332-3032323020012102-3110301220001110-1310120203033232-0221332030322110) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-3232232000112022-3203000010223311-2102021021002032-0003030131112211-1231223110332232-1033023113030313-1323210013331301-3123031311100031) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-002.md#canonical-1020322220112012-1331103312102013-2112030213211033-0121011330123020-0221323101212101-0302120110123311-1012300221310331-1222002130133310) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-002.md#canonical-3203022131030022-1023231022332103-2121110300010113-0320030311100123-0230110331113313-1233311230022033-2020031021001221-3311331130103320) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.value` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.value](resources--proxy--reference--group-002.md#canonical-3322230233023202-1113310222313103-0131222002322030-3231112222300332-3111103111310233-2102031301123003-0211112013032212-1201223323001133) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_remove` | [dynamic_proxy.http_proxy.more_option.request_headers_to_remove](resources--proxy--reference--group-001.md#canonical-0313200131120023-1120113022320011-3033310100301112-1301010002313023-2020302100122010-1302313223312232-3012210330310031-2122220002213100) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-3021311133323301-3301313030222230-2103232331230230-2020232123201303-2201230232101011-2232333220112133-2101320021201103-3201231121111213) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_domain` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_domain](resources--proxy--reference--group-002.md#canonical-2100313122110033-1010230312313003-0012110330301001-1001211032320223-3001103223230303-1100310111110023-2131103230213112-1023321333303223) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_expiry` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_expiry](resources--proxy--reference--group-002.md#canonical-0002121332213331-2321002022022100-0201121011130323-2010113200233231-3220312222223013-3121121033102331-3133120301322010-0201300032000020) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly](resources--proxy--reference--group-002.md#canonical-0132213222030030-2230100100112012-3203011102032002-2232230021210201-2001312131132112-2301301210312231-3023030011232221-0230100002102131) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned](resources--proxy--reference--group-002.md#canonical-1011133030203203-1212123002112132-2221221121302222-1300211221112312-0032301202320100-0220300310020031-1301310010032030-3130102310311311) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_path` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_path](resources--proxy--reference--group-002.md#canonical-0222022313000211-3101310021233113-1133130002223022-1230133100303323-3020202302232110-1321231123231011-1322210021222123-0111220333020311) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure](resources--proxy--reference--group-002.md#canonical-2200113333202300-3201201220321103-3312032233210203-3133012003130021-0203013003322000-2331002030201312-2130210210200332-2100333220003102) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain](resources--proxy--reference--group-002.md#canonical-1020321221322313-0032323310213223-0102000222211333-0121100202121322-3011122202112313-1120201023323030-0100211101011030-3113102312033012) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry](resources--proxy--reference--group-002.md#canonical-2321231121323100-1032111112203313-1223232230323303-0223030120210330-2010030100221332-1112012333022203-1032011031002111-1310010222002020) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly](resources--proxy--reference--group-002.md#canonical-0222131313232122-1100212113122131-3012311313322110-2100020223332330-2031311031323132-1111113001102022-0002300001123013-2103021023212002) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age](resources--proxy--reference--group-002.md#canonical-0211120332130021-0001213320131033-3312012332112221-1220202101331222-1301330223031010-2110122322320211-3133203313222332-0130222213201333) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned](resources--proxy--reference--group-002.md#canonical-2312321101032211-2221330131000212-2123311023211202-2323331110022111-1002033003323001-3120131030122020-2210212132101220-3330313120103010) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path](resources--proxy--reference--group-002.md#canonical-0023322030320302-1323031111323303-1023101111223010-1223111223310313-1011203100012011-3013131312203201-2200123131333312-1131222002113101) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite](resources--proxy--reference--group-002.md#canonical-3103132203323311-2130010101333120-3303002002132002-0112220003103333-2101201021013210-0330231122022032-3020012022210321-2200103020110221) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure](resources--proxy--reference--group-002.md#canonical-2113332000110031-3303031200131001-1323303213212123-1313023302311200-1112203333213000-3123102112132312-3130212022022133-3202332131223131) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value](resources--proxy--reference--group-002.md#canonical-1322201103113013-3210023332323021-3212102010002121-1131131300221211-3210220110323233-3310011301313210-3003101222311123-0102033032312102) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.max_age_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.max_age_value](resources--proxy--reference--group-002.md#canonical-3220330222220321-3010233330132310-1231312230122011-0003133000101112-3002033203310132-2101233223333121-1110010031102100-1322202230321033) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.name` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.name](resources--proxy--reference--group-002.md#canonical-1123002130011300-3200103312031033-3023113101103230-2230323311312200-1212233331110002-0131223311212332-1023121112113021-3123221002112132) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.overwrite` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.overwrite](resources--proxy--reference--group-002.md#canonical-0003320203011302-2012103030001301-1232323300233133-3312221233023313-1323111310231011-1322313110120300-0223223231333220-1003112000130010) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax](resources--proxy--reference--group-002.md#canonical-2111303103112302-2322230001000320-3010100021120311-3203323020120103-1100130330012200-1232301011211323-3313203020022220-1222203302130131) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none](resources--proxy--reference--group-002.md#canonical-0302122300222310-2311210320302011-2033321103023233-3202000031012322-1232203313123031-1230003330233131-1131231203302211-0333111113012120) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict](resources--proxy--reference--group-002.md#canonical-0303100222301012-0033123312320031-1212000003323222-0032132112033021-0122220103113113-1303323011211310-1121220131133323-0111203130230130) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-2300013201200212-3313122223010310-1003333012012002-1102003132021303-0203201000211030-1113113310033120-3130202012311220-3103221330012133) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-2032201022302321-3220001201101001-3121100302131232-3221303122110200-3222103302023111-2033022111030223-1333302003213001-2002213102210113) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-002.md#canonical-3200330221230330-0303003013011101-3201233120132000-0111201003030010-1000003030202301-0012101021320321-0032220321103120-2230312033022203) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-002.md#canonical-2103222111203121-3021223030300201-0031032312213301-0130103131021110-3302312022020301-3311131113330132-3213320121332302-0120231100020033) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-002.md#canonical-1200123303110021-3122121300022111-1030222021213031-2022330131022130-3101223113220112-0022033111113222-0012131000323011-2312333122021032) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-3330123310023133-3320003103012110-2203233331231002-3030333201000232-3210220232010132-1131023322023330-2332012230231230-3223230111200022) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-002.md#canonical-0123003323021323-0110002311323103-2122021320231112-2310202322131033-0011003010013003-0232000013000311-2203232110333231-1311212023230013) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-002.md#canonical-0201123322102021-0332103110120030-0303131202223122-3010101022030332-3022321301233220-2131221001201222-2111101201101301-0330033203113222) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.value](resources--proxy--reference--group-002.md#canonical-3033331112111032-0030310131023121-2333232311231221-0200000323033201-2310113121310201-0320013201302332-2101113012222202-1313103231001231) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_remove` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_remove](resources--proxy--reference--group-001.md#canonical-0011201022031310-1032000222311323-1202322200110210-2312323321000203-2003221333333022-1112022230200030-3222212111310021-3010230222030223) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-002.md#canonical-0202100121330003-1013012330213311-2032200112120101-3210302011102313-3010233130320012-1110001020300000-1310203033211122-0003221103123220) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.append` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.append](resources--proxy--reference--group-002.md#canonical-0313113011332211-1211103020323313-1103320000202200-2311201322321131-1013122100003131-0131131130221133-3220231302031120-0002020010113121) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.name` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.name](resources--proxy--reference--group-002.md#canonical-1001332203111311-3201330300033222-2220330221012232-3003132020323121-0223302002212001-0122033233232232-3120001132233100-2133312213322230) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1323121332033300-1003030231121131-2322211133003010-2303320120232110-0103103102002033-0121102131133030-0230010113210300-0130230201111100) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-2103102221213211-2203011112002232-3222310032320032-3233130132202023-0333202122011130-1303223230120223-3211010120120023-3222033001101313) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-002.md#canonical-0111312023100012-0201030133111300-2131321030212110-2010221001021223-1110103121103301-0112321333220132-0221021002230111-2222133112121320) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-002.md#canonical-3013020020103110-1121110332111103-3232000222113201-0222012213012033-2300303312103011-3203022203202232-3210012221013021-0321330200120230) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-002.md#canonical-1110231012003003-1011111013100011-0122121011100312-1132312303001033-2223121312223300-1203233002321003-1130103323002213-0020232102020112) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-0332313130210102-3232332311010310-0213130222230022-2030032130100231-1130020222200021-3001331331213003-1013133312021210-0102302031000313) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-002.md#canonical-0001323233103000-2131102212210230-0021113103200010-1312233010221122-0030120202322203-1310111302111023-2220311202212032-2201232223022120) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-002.md#canonical-1303101323232033-2311301322330003-3111330303000301-3111332002311010-1230120122211013-2203320132112002-0232002301012013-0110122233321010) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.value` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.value](resources--proxy--reference--group-002.md#canonical-1200002120131302-2003033301113000-2301133331311001-1203311023020133-2333102200310132-0123113201112122-2022220132231233-3000203313233002) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_remove` | [dynamic_proxy.http_proxy.more_option.response_headers_to_remove](resources--proxy--reference--group-001.md#canonical-0320233002132310-1220121303212030-2123030301003223-0111002313320313-3310122023120222-1231111203003133-0110233103101310-3201302013210102) |
| `dynamic_proxy.https_proxy` | [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-2100012333203313-0133032221230031-0122133000311232-3230212201200101-1112310123103300-1212322311331203-1321131033021110-2123132102220131) |
| `dynamic_proxy.https_proxy.more_option` | [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-2233030200133233-3231230013030010-3020020220201332-2131311031203311-3130100232221111-2112323321112031-3220313030202200-2202302301103331) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy` | [dynamic_proxy.https_proxy.more_option.buffer_policy](resources--proxy--reference--group-002.md#canonical-2201121022133030-0101122321201203-0310223122101231-3222113211021310-3120020001122032-0100322200322112-3123223311222113-0313321322310303) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy.disabled` | [dynamic_proxy.https_proxy.more_option.buffer_policy.disabled](resources--proxy--reference--group-002.md#canonical-2002021103022022-3202332002103332-1030101320100133-2112023231303200-1202300313232133-0130201001313220-3102322112132010-2101302313221133) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy.max_request_bytes` | [dynamic_proxy.https_proxy.more_option.buffer_policy.max_request_bytes](resources--proxy--reference--group-002.md#canonical-3110033021333123-0003222013110223-0111313201231221-2223310320111321-0130021202001230-1310013030002031-3131331210002230-2113211320000333) |
| `dynamic_proxy.https_proxy.more_option.compression_params` | [dynamic_proxy.https_proxy.more_option.compression_params](resources--proxy--reference--group-002.md#canonical-3220220331212302-3122301000220211-0130210223110023-2033213010213330-3212312320133211-1211331133022223-1322131311133131-0312333133322013) |
| `dynamic_proxy.https_proxy.more_option.compression_params.content_length` | [dynamic_proxy.https_proxy.more_option.compression_params.content_length](resources--proxy--reference--group-002.md#canonical-2112113302012012-3121020300113223-2321013212003013-3000322232122323-2110022112000032-1310333131231321-0002301202111200-2331102331333002) |
| `dynamic_proxy.https_proxy.more_option.compression_params.content_type` | [dynamic_proxy.https_proxy.more_option.compression_params.content_type](resources--proxy--reference--group-002.md#canonical-0212032303300121-0020113001132301-1001030321000221-2311231200123021-0311033113030233-3022332010200220-3120123111023301-2332330020002312) |
| `dynamic_proxy.https_proxy.more_option.compression_params.disable_on_etag_header` | [dynamic_proxy.https_proxy.more_option.compression_params.disable_on_etag_header](resources--proxy--reference--group-002.md#canonical-0102211203130010-2111323033102112-0020003110322101-1202311110230103-3121230220002222-2320200223101131-3321122032212003-1002110112013021) |
| `dynamic_proxy.https_proxy.more_option.compression_params.remove_accept_encoding_header` | [dynamic_proxy.https_proxy.more_option.compression_params.remove_accept_encoding_header](resources--proxy--reference--group-002.md#canonical-0131212132302121-0123131023033013-1123023000110200-3303301002211331-3222130313123131-3200321012112303-1303232021222033-1322233131211330) |
| `dynamic_proxy.https_proxy.more_option.custom_errors` | [dynamic_proxy.https_proxy.more_option.custom_errors](resources--proxy--reference--group-002.md#canonical-0120301120131103-2132222303132210-2230010301111310-0101301232231221-1230223112031231-1320120223021012-1112221211020313-0203203121120311) |
| `dynamic_proxy.https_proxy.more_option.disable_default_error_pages` | [dynamic_proxy.https_proxy.more_option.disable_default_error_pages](resources--proxy--reference--group-002.md#canonical-1313030123210300-2213213220000330-1032332300131232-3101000213121002-1102230112103203-1102032011203031-0000332210022113-1210232330302000) |
| `dynamic_proxy.https_proxy.more_option.disable_path_normalize` | [dynamic_proxy.https_proxy.more_option.disable_path_normalize](resources--proxy--reference--group-002.md#canonical-2333222322331033-3222333022323123-2202322233120303-1310321031232131-2000011133110311-3210001302331130-2033120301003120-1330001211113310) |
| `dynamic_proxy.https_proxy.more_option.enable_path_normalize` | [dynamic_proxy.https_proxy.more_option.enable_path_normalize](resources--proxy--reference--group-002.md#canonical-3333302232022233-2133110122231321-2211110303020302-0033121013122233-3333203300031203-2322123000101300-2133101101322123-1030302103100201) |
| `dynamic_proxy.https_proxy.more_option.idle_timeout` | [dynamic_proxy.https_proxy.more_option.idle_timeout](resources--proxy--reference--group-002.md#canonical-2020003313312211-3310112200112201-2220101200202223-0231113211011101-3011202211231021-1223331101100130-0310300001331312-0311230313030001) |
| `dynamic_proxy.https_proxy.more_option.max_request_header_size` | [dynamic_proxy.https_proxy.more_option.max_request_header_size](resources--proxy--reference--group-002.md#canonical-1133031030113001-3211122221310122-3100200031331231-0212030022132112-2020111131230310-0110123120311000-2310002321120023-0321130232003202) |
| `dynamic_proxy.https_proxy.more_option.max_requests_per_connection` | [dynamic_proxy.https_proxy.more_option.max_requests_per_connection](resources--proxy--reference--group-002.md#canonical-2111313221213020-2012021110002133-1202231210133100-0133331301032122-2023322331303132-0212101112133301-0322333123101210-3202233221210030) |
| `dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection` | [dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection](resources--proxy--reference--group-002.md#canonical-2202220132301332-0233211122313101-3233330010112233-2221312112311012-0333322201002021-3313002231102022-2010101030303030-1013313030022010) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-2031210233122310-2303310112123101-0121203120323031-3013303230313120-2010100230212300-2003130301332310-0001120002120033-3212021031123220) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.name` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.name](resources--proxy--reference--group-002.md#canonical-0233121311213230-2320122132202220-3310111002302213-0030320122311312-1320111112110000-2203112313300130-1101100233322031-0211212311210330) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.overwrite` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.overwrite](resources--proxy--reference--group-002.md#canonical-1202201130221202-2012020233221231-2230322000333212-2231222321330200-0231112110110110-2100202131103130-0122322320312311-2121213210231121) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-2012032101103300-1323100203330222-0200112130223020-2013112200332301-3231132133123001-3133023101103033-0210133130020133-1331303100213203) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-0032010031122032-0202012313311032-3100202003322221-2122221313301333-3133332213222030-2323212031323322-2233210310101122-2003321232101100) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-002.md#canonical-3122102131003202-0121230121032320-2101111100300120-3021202322122002-1333211111312322-0012321301120022-1233121320011110-0231302011032010) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-002.md#canonical-1232310100202231-3212032220133120-3011112001022022-2022302132000233-1330101323101220-0033010003030121-0123003202231103-3033322302202211) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-002.md#canonical-3021100321202232-0331330102230010-1313201122300013-0111131103301112-0332203231010312-2303100211322132-3210303323031223-0211030330312030) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-3220130301232111-2023321103132200-2000201002201310-2120023030020110-3101302323310310-0301221131200022-3212310233210100-2113102203313331) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-002.md#canonical-1010222013311110-3013312213202221-0010103320222100-2021001133331221-1212111221030030-3310111033102323-3131221122320202-0120213103031011) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-002.md#canonical-2231312211132020-1312013002030110-2332023221121312-3133210232322231-0100332300110131-0033231133322021-0203323230101133-3000000102211231) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.value` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.value](resources--proxy--reference--group-002.md#canonical-3301311032222000-2132003130312102-0222222001302200-2023201332221203-0203000113010010-2210121303320223-2123221123132133-3101311130220000) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_remove` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_remove](resources--proxy--reference--group-002.md#canonical-2233332213222302-2310300122112333-2122003112003030-2312203000103123-2033223300120300-3130203132211022-3311110033231000-2320300003231121) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-3032001302313331-2033020223001322-1322001120103212-1113303122333121-3102203031330223-1320332022312233-2132031002333320-1113133130000102) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.append` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.append](resources--proxy--reference--group-002.md#canonical-3233301321211031-2013200220102310-1312112023102103-3011011010013121-1023013130332311-2223311313330303-1122333330323220-0311300031123102) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.name` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.name](resources--proxy--reference--group-002.md#canonical-3120321211211202-2211101113213230-1322010210302122-3333113012021100-2032222311002333-0233100220320010-2312120201020121-1013232303321021) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1303010212333201-0312112222010130-3120112023110313-0111300333032303-1303110323331133-2112203122322323-2322020222113001-2310023120202323) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-3132331131232130-2101121322323000-2133310112310031-1001212003223330-3321013212313120-3301331230321203-3223030323201323-3230003121001321) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-002.md#canonical-3131320232011121-1103213112011320-1023012021310212-1102312211233320-1320313313330212-2300213220301013-1321123012033100-3113330233303332) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-002.md#canonical-3330110111133132-3331010003313132-0103303201211200-3323213211210211-3110312122020332-0210213222032011-3311222113131121-0200220302123312) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-002.md#canonical-2320033032000231-2212113200133310-0302322232232002-0312301120311223-2020310102030123-1302223033112003-1113331311021233-1300321130020012) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-0310112231113133-2020130021131213-2322210232013012-0221312302133303-1121013010113222-2030022331030323-1131330111233220-2003023313023022) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-002.md#canonical-2303020312220032-2030321103210031-3221323111000222-1111301302223122-3120332101103212-1331311203033300-2233030323311112-1322300302210013) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-002.md#canonical-2103001302111231-3120010233233213-0010232311133301-1323102122312231-1032132331201232-1332231320131312-3101133301323112-1330013111301221) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.value` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.value](resources--proxy--reference--group-002.md#canonical-0032021220222300-1223123002201013-0310321012201013-1302001133331103-3103201101323013-0211331220000001-2221332013231033-1203223223303021) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_remove` | [dynamic_proxy.https_proxy.more_option.request_headers_to_remove](resources--proxy--reference--group-002.md#canonical-0333222211023020-0202111230101111-0013132323101331-0333002103200233-3113331100111002-0330101021222100-1311231023010320-1030133022011333) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-0131201032122021-3010202001203131-3333211331121103-2312023020312012-2001020033333010-1212022102233300-1323220210123323-0022013312202311) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_domain` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_domain](resources--proxy--reference--group-002.md#canonical-3312132021303220-3012010312122123-3121230220022103-0313223031003232-0020013312033112-3122201031312331-3201321203300333-3302021323313103) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_expiry` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_expiry](resources--proxy--reference--group-002.md#canonical-1332221222231313-2221030223213213-3013110020203232-0100103301222001-1301120023222130-2203303112110112-2223023212202101-1200331131302302) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly](resources--proxy--reference--group-002.md#canonical-3331321132111110-1303201230300302-0133132303130120-1301301310003332-3030312203032331-1230323220232211-3201201113312202-3302003011000303) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned](resources--proxy--reference--group-002.md#canonical-1022300231122233-2300033323102210-3210210032322033-0103000321223020-0003332020021001-2033302332220231-1233303333300203-3330021000233031) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_path` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_path](resources--proxy--reference--group-002.md#canonical-0321121023130330-0101333112232023-2302301202323130-2333210020122303-1010111313310300-1133230111330113-0033203203302001-0200211333102031) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure](resources--proxy--reference--group-002.md#canonical-0320101201012130-3320000121222322-0031313133130131-0102313220030212-2221321203303301-2221132310233213-0120211032210332-0310102022311012) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain](resources--proxy--reference--group-002.md#canonical-1102102133022220-2301331301132113-2122123132110202-0101023032010320-2101011033101130-1320110001122221-2033332310203100-2131112031122131) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry](resources--proxy--reference--group-002.md#canonical-1311232102322203-1001312301000031-2020313310322201-3002221320223020-0130032101211000-2010202011303310-3121132222001120-2120230312230221) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly](resources--proxy--reference--group-002.md#canonical-0331322312001122-3332031002303031-1012002121213213-0302133020212111-1221110222013022-2002032021221011-2211023131100310-1221102211323200) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age](resources--proxy--reference--group-002.md#canonical-0021033021202111-0202011120112302-0320213102111313-3001303133012031-2331211301200103-0311322321202030-1120312230312322-3321202312011030) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned](resources--proxy--reference--group-002.md#canonical-2330223031212303-2210033121111302-2330203200131033-2013001021133011-2211132331013130-1311213321021200-0303332103030133-1220121321013321) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path](resources--proxy--reference--group-002.md#canonical-2210222201222220-3322123223010022-2320221122133010-0022221332023210-1221230022310211-3313301302313232-1011223331131213-2120030132122021) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite](resources--proxy--reference--group-002.md#canonical-1132103301333331-2100322212221230-3303203330321333-3232023010201303-1230300230322233-2000300022011223-1302133020123210-3201302300202320) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure](resources--proxy--reference--group-002.md#canonical-3022331022100322-2320213332033311-2211032211200223-1201013302021000-2200012201012321-3220210333333301-0112310322101133-3121201211011300) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value](resources--proxy--reference--group-002.md#canonical-0201122322233323-0301231210021300-0301220222232332-3310031130132213-3103120322000122-2021323310123102-0303333313212213-0303022313321112) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.max_age_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.max_age_value](resources--proxy--reference--group-002.md#canonical-0330131020302331-0032121212210102-0012320311313033-1210032030322203-3021012300033202-2311021312302011-0010110133021010-0220130013221313) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.name` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.name](resources--proxy--reference--group-002.md#canonical-1133300320123320-1022133111203132-2131111313122011-1013023322113102-3122111001012002-3312321033010320-0323011201131113-1011312133323020) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.overwrite` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.overwrite](resources--proxy--reference--group-002.md#canonical-2310321330021330-0020320321130101-3223002321200100-2031300320233031-2202210323210322-3113223122132121-0032323111202323-1123201020300312) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax](resources--proxy--reference--group-003.md#canonical-0201011101112200-2113213311121101-2133131131202020-1300100130302132-3303122121112022-2311211012332301-0032301221111112-2301033133023030) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none](resources--proxy--reference--group-003.md#canonical-3222132023011210-1110103201003321-0222320123101312-2322113123302302-0222031222120111-0203301211003210-0100033103201312-2102122012300110) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict](resources--proxy--reference--group-003.md#canonical-0132131101102101-1101030130221002-0012310102200113-0330011100223211-2320212321000130-3010002330101332-0203102331210211-1203120123031020) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-0112032002023223-3312231333031222-3133332123303123-1030201130032211-1213023132333002-2123122111221122-3010203201213311-2001302223102101) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-0023033123100330-3111121320022312-2103302222000220-3230123020213020-2132310322100001-1113111222123103-2301322031030011-0132203213012112) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-003.md#canonical-1100232022331223-0002321011302200-3212300113210210-2123003310033223-1133303133133012-0032100100002330-3032332312113113-2331020320030012) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-003.md#canonical-3333233213211113-2002103012133211-2112200233100331-2102103113113031-2122022310320210-0301332001310023-3333202003101233-2322210220020002) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-003.md#canonical-0222013012303013-0301332323122310-2031323023210102-2000320302310113-1123103121011001-1013103012101122-1120023203311203-0310131330332233) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-003.md#canonical-0330322112003110-1202120010330112-2200021201333301-3310233132223113-2223220313310003-2003101132010233-2123121333332100-0210100231002100) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-003.md#canonical-2211032110332013-0032012131230101-2210320323232101-3230000221302302-0200303333301120-0332313103101101-0032132220100012-0112030031312133) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-003.md#canonical-0102200001102300-2020103012130203-3330122112101200-1100011133120231-3020110320030213-1310202000033023-1002022133102210-2222123313223031) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.value](resources--proxy--reference--group-002.md#canonical-1323110321001120-2110022133101211-3102123123300333-3111123130313331-2023020331210132-0230010210203323-2210322131222213-1022112330132010) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_remove` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_remove](resources--proxy--reference--group-002.md#canonical-1333121220313131-1002010123102033-0023123312002331-2320312030221110-1302103230333132-0313031310100112-1302321211021203-1102211302130302) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-003.md#canonical-1133210313300331-3002323303321202-2131123113011113-0110002013030120-2100112013231330-3113313302020303-0332122111010121-0110203300323321) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.append` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.append](resources--proxy--reference--group-003.md#canonical-3111120033103020-2333023022131332-0111132031133213-2303123230121321-3132012311132323-3311330101020022-2012302330232330-0100301302203310) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.name` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.name](resources--proxy--reference--group-003.md#canonical-3001231000010331-3331012131330110-0332113320031220-2330222222303001-0022010023001213-0102312320332223-0323131223301203-2210101131130103) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-2100222111320221-1333312123133233-2303301010221011-1223021003231021-3330021032021021-0010112103110323-0313231210122311-0023011233310133) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-2003233321123031-3012301212203311-3033330033123331-1000030311233000-0313313133330030-2222322321203102-3130123020132011-2221020002303113) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-003.md#canonical-3013113120323133-2232022010013102-0200121231020022-3003020233321311-1212013322320122-2011011233000203-3300201022001100-0133003123032202) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-003.md#canonical-1230011002021323-1103233102112000-0101031123023023-3202111321330231-3200321302101122-3333333111101203-0203203031000313-3102210211100020) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-003.md#canonical-3201112332231131-3123113303111200-3313233112021133-3113322012010121-3003313331012212-1222221310001332-0321101100123302-0213112131103232) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-003.md#canonical-0332223311302233-3302313000111223-0311223120003201-2210033322331222-3223130311331012-3133110203231213-3332110031021320-2132111112303312) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-003.md#canonical-2333011103100102-3200121312301022-2321130001333031-3003023032001312-2010031020021322-0000330111203320-3210130211001302-1233032332121232) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-003.md#canonical-3321231202133103-1023233330103133-0022321012110012-1100120213320032-2333320000202003-1011123103232101-1333112233100113-3123031233200223) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.value` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.value](resources--proxy--reference--group-003.md#canonical-2130101231322032-0321232233331222-2220020211122113-0001102323020020-2120300000231113-3220020011231211-1333133030213011-1003033130310120) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_remove` | [dynamic_proxy.https_proxy.more_option.response_headers_to_remove](resources--proxy--reference--group-002.md#canonical-3200310321210311-1212323313231030-0033331230231131-1112300133012131-2202201310220203-2312320203032301-0232302010123030-1003233103031130) |
| `dynamic_proxy.https_proxy.tls_params` | [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0201002202231223-2330122223021330-3010321311130202-2221331101011331-0031101033333202-2313100200131213-1302213121122201-3203001001203013) |
| `dynamic_proxy.https_proxy.tls_params.no_mtls` | [dynamic_proxy.https_proxy.tls_params.no_mtls](resources--proxy--reference--group-003.md#canonical-3110323201120332-2313011321232032-0321031121303201-1202230302302310-1300031030210202-2130321210311330-1333033103230300-2212320102203110) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates` | [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-2031120113301231-2000031022332001-1022302121112230-0030201113230120-0130310002102030-2213013201220022-0303112131331320-2323322111003301) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.certificate_url` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.certificate_url](resources--proxy--reference--group-003.md#canonical-2323332021332122-3032132312220122-0021301003312032-0120323002120121-3101333201111221-0113200100120131-3231112210133213-3123323102213231) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms](resources--proxy--reference--group-003.md#canonical-3133010331132323-3100230200101231-2003233212011100-2010012212133233-1332113111100023-2002203322210110-2033320021110310-0202210313011331) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--proxy--reference--group-003.md#canonical-1203112202302002-0133332013010003-1321322021312221-0303012201311200-2020101021300131-0311023111023102-3031232223012121-3333030123130231) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.description_spec` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.description_spec](resources--proxy--reference--group-003.md#canonical-1233032203323230-2121123121301102-1102203133223212-1002120302000000-0200233211020032-1322001213303030-3321131012302222-0000113022220210) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling](resources--proxy--reference--group-003.md#canonical-2302023312121112-2211301132203021-3131311123221312-0303220132023132-0222010110330221-2231303022120331-0323033232133021-3301230011212332) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](resources--proxy--reference--group-003.md#canonical-3130303022320203-0010103300122112-0020313132130223-1213200333310201-2300301133313320-3222303320323113-1223201120300112-2331120323323323) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-1022212211203011-3100313310201130-3121031100212132-3200330301222013-1112220032120222-3203202200020133-2132310012222110-1121020103123331) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-003.md#canonical-2013032021321303-0320331131310021-0000100310020132-2002223332301131-0102120312230321-2320021022132221-1200113230011121-1111212331213310) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.location](resources--proxy--reference--group-003.md#canonical-2212321321032031-0311313212033211-0031102203310011-0320232013131020-3110222303003100-2001010130221223-2313103220201233-3122113123232312) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--proxy--reference--group-003.md#canonical-2012220303321212-0023123010222310-3303211223103332-2023113022122300-0111132313220221-0122312133333122-1210033033222133-2232011210213112) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info](resources--proxy--reference--group-003.md#canonical-3132322130131110-1210001320210331-2021202123111033-3132223321121030-2113133211030302-3022030310212301-0320210202233323-1233113313012311) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.provider_ref](resources--proxy--reference--group-003.md#canonical-1230322311311232-3003312200311330-2002111321022301-0233312002033033-2123011310112000-1200321001112030-1213031101302121-3232101021130230) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.url` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.url](resources--proxy--reference--group-003.md#canonical-1003023330110121-2102020312232021-1022331001003300-2213320130221021-0103022231020200-2312321230010322-3321312201103330-3132333331113210) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults](resources--proxy--reference--group-003.md#canonical-0332122102233300-3331002010330122-1010131302320320-2312300020321102-2313010120131300-0200013030112300-0133230310000122-0332222300321322) |
| `dynamic_proxy.https_proxy.tls_params.tls_config` | [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-1231102332233303-1002131231012201-1222213131131203-2213033133313332-1202333030332210-2301022102113303-0120203220131132-0120132312233031) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security](resources--proxy--reference--group-003.md#canonical-0122203202132121-3102133100012102-3033333030010120-1210322232123303-2033133030222302-2221233233013110-0030030213000033-3221121133011312) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.cipher_suites` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.cipher_suites](resources--proxy--reference--group-003.md#canonical-3202110222133320-1231121103021103-2021130200121110-1213223033131101-0333210300010003-0112231100033122-1311030331120302-0020121110032011) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.max_version` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.max_version](resources--proxy--reference--group-003.md#canonical-2001303222012110-1030312002011333-1322212222001321-0201022311303100-1020212120023332-1200013313020302-0130332133201103-3202323213023123) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.min_version` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.min_version](resources--proxy--reference--group-003.md#canonical-0131113013211331-2130331233213031-3232213301032122-3233133121323031-1231130310131223-0022020032032031-0020032300102232-2102320130233310) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.default_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.default_security](resources--proxy--reference--group-003.md#canonical-3022110123203303-0330221223113221-0200222033313103-3320231223031032-0313310113031103-1232021230333222-0111201101112211-2030221200113303) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.low_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.low_security](resources--proxy--reference--group-003.md#canonical-3010021210222012-1031100021203223-2101112110133132-3323332130120203-2331332202301021-2032122212101032-2333212002202330-1221022013021101) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.medium_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.medium_security](resources--proxy--reference--group-003.md#canonical-2230000203213000-1120212323122122-3131222220031213-2122333132112022-3033123113032102-1103102203301100-1133032112312220-0312222132210211) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls` | [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-1213230003223333-2310120333333212-3222201000132322-0211123031123233-2300311201230113-2300132032310132-0323103122310011-3211110111022030) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.client_certificate_optional` | [dynamic_proxy.https_proxy.tls_params.use_mtls.client_certificate_optional](resources--proxy--reference--group-003.md#canonical-2213331233003232-1123001122303213-1302120303331310-2100102113001111-2113211130023010-3120111112220200-3033210300220122-0100012102221023) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl](resources--proxy--reference--group-003.md#canonical-3000233230301000-1310120333101011-3222232112301300-3330010330030100-0122231311200121-0323121320102013-1003111131323002-1121200321020130) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.name` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.name](resources--proxy--reference--group-003.md#canonical-2100012213311002-3112111331232000-1232231331102003-1332310010303213-1313112232112310-2102322101210121-1203110302320212-1123312222032112) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.namespace` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.namespace](resources--proxy--reference--group-003.md#canonical-2011013033332322-2012013022331210-2232130323320033-0030312310210131-3021133003101032-0003013231221030-0122010103103023-0130330003111213) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.tenant` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.tenant](resources--proxy--reference--group-003.md#canonical-2301200221321002-2213201030300011-0230112203033202-2310320032022121-0210332320310012-0233223032223002-1331133331312313-0330332122330022) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl` | [dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl](resources--proxy--reference--group-003.md#canonical-0012332021113300-0101132122322010-1233111203122131-3233321011020222-2220213101211133-2022013323301130-0203100030230012-3020202233320013) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca](resources--proxy--reference--group-003.md#canonical-2101133330131000-1210300310201203-2120330333320222-3220313320012132-0033101031132320-3020302310233202-0333113201311211-3011310001320203) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.name` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.name](resources--proxy--reference--group-003.md#canonical-1113223012213120-1312200010332033-1303302220020303-1303320023202333-1333132130311333-0310321130220013-3200303122210012-3102331130023212) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.namespace` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.namespace](resources--proxy--reference--group-003.md#canonical-2020320002332323-3121111023011020-0120202321222220-3022130211300131-3311302020132212-3122311021011222-3321311230212123-3022203003201100) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.tenant` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.tenant](resources--proxy--reference--group-003.md#canonical-2131203132123300-1331311113102133-0030321222032120-0010212302033301-1002230022222033-3230221211331332-3013223131101321-1030321000111212) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca_url` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca_url](resources--proxy--reference--group-003.md#canonical-2211221112201113-0130230100233301-2332111312231113-1330122333302003-1122111001202302-1211111012300233-0303213023202221-1333000330033201) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled](resources--proxy--reference--group-003.md#canonical-3103231300321003-2332320233132101-0112303220110113-0013222023313233-1022010303023100-2222310231010132-1113121132121211-1333311111230012) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options](resources--proxy--reference--group-003.md#canonical-3211002013213020-0021331212032032-2110133202312122-3130330121020202-1020020322133333-1022130000232201-0301301131200330-2203002331301331) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options.xfcc_header_elements` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options.xfcc_header_elements](resources--proxy--reference--group-003.md#canonical-0303132300303032-3323122220133003-2113000112303033-2123202220313310-0231133312333303-0123223021330130-2301310210230033-1003331303013100) |
| `dynamic_proxy.sni_proxy` | [dynamic_proxy.sni_proxy](resources--proxy--reference--group-003.md#canonical-1200030100100302-3313101203023312-3310213031113310-3011223200220232-2103312102230333-1313221202010332-0111200212103102-0112023213123033) |
| `dynamic_proxy.sni_proxy.idle_timeout` | [dynamic_proxy.sni_proxy.idle_timeout](resources--proxy--reference--group-003.md#canonical-3102213023212303-0330313003231333-3320111333210230-1103302013332103-1112333310132031-2121023231131130-0232232223121321-3301102012232102) |
| `http_proxy` | [http_proxy](resources--proxy--reference--group-003.md#canonical-0023030222010311-3300322230002013-2001102311323223-1122022322221322-2020120013210033-0032330013323132-3222112033103212-3223122333232132) |
| `http_proxy.enable_http` | [http_proxy.enable_http](resources--proxy--reference--group-003.md#canonical-2200301101203123-0013221023302022-1232013120121021-3011201200132231-1230013223110322-3123022010130320-0222232030103110-2221202002230222) |
| `http_proxy.more_option` | [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-2221001223210112-1100323203103002-1221130002010130-3330211303220210-3321303213313233-3133102003321312-1003210223123213-2133113122313131) |
| `http_proxy.more_option.buffer_policy` | [http_proxy.more_option.buffer_policy](resources--proxy--reference--group-004.md#canonical-2030331223113202-1301222121300032-1333023200302230-2223210310123030-2103212211031020-3020122131311010-3212011120111103-0222130023203220) |
| `http_proxy.more_option.buffer_policy.disabled` | [http_proxy.more_option.buffer_policy.disabled](resources--proxy--reference--group-004.md#canonical-3030213210301303-3333030201220321-0230101011202023-2003022131230110-2101302003303110-3001120103222012-3333331030333213-3200333203112001) |
| `http_proxy.more_option.buffer_policy.max_request_bytes` | [http_proxy.more_option.buffer_policy.max_request_bytes](resources--proxy--reference--group-004.md#canonical-3022321331221031-3101331230333301-0021011112311301-0221312013023013-2132031333301013-2331310300322220-0031213220302100-1133310220022010) |
| `http_proxy.more_option.compression_params` | [http_proxy.more_option.compression_params](resources--proxy--reference--group-004.md#canonical-0320033203010212-1211021031113132-2000000320222000-3110300022022233-1020031200101233-0020321320102303-1032023200311321-1231210233003101) |
| `http_proxy.more_option.compression_params.content_length` | [http_proxy.more_option.compression_params.content_length](resources--proxy--reference--group-004.md#canonical-2231331100321200-3031101132033022-3311333302312010-0302120312200032-0303300133131332-1202231322023211-0320033210230112-1230101110200303) |
| `http_proxy.more_option.compression_params.content_type` | [http_proxy.more_option.compression_params.content_type](resources--proxy--reference--group-004.md#canonical-0132030320221310-0320321031021113-0212032212130010-1101103000110030-0001123021221321-3003332333033103-3001333000010201-1132113220310322) |
| `http_proxy.more_option.compression_params.disable_on_etag_header` | [http_proxy.more_option.compression_params.disable_on_etag_header](resources--proxy--reference--group-004.md#canonical-1011211133321111-1303000003000102-1123222321310302-1200221131213310-3120023031123200-0120200020000312-3031232212120311-3212222213031032) |
| `http_proxy.more_option.compression_params.remove_accept_encoding_header` | [http_proxy.more_option.compression_params.remove_accept_encoding_header](resources--proxy--reference--group-004.md#canonical-2132023021032303-3122203000311123-0313311221113132-1120321210130112-1200101120313313-1232323222000222-0321031221010231-2230333022231131) |
| `http_proxy.more_option.custom_errors` | [http_proxy.more_option.custom_errors](resources--proxy--reference--group-004.md#canonical-0313022031120112-0330201322330102-2202030133331203-0310301212223021-0011100232010303-2121001310020313-3100201212011233-0003031130000313) |
| `http_proxy.more_option.disable_default_error_pages` | [http_proxy.more_option.disable_default_error_pages](resources--proxy--reference--group-004.md#canonical-3231320000303110-1012233021103302-2102122331023303-0231021323111200-3100110000332321-3223322102131121-3110020330303030-3210220023021330) |
| `http_proxy.more_option.disable_path_normalize` | [http_proxy.more_option.disable_path_normalize](resources--proxy--reference--group-004.md#canonical-0112221233100331-0113012122102211-3011200303303011-1031032220303131-1100323202123211-2233030223201213-1331122233020302-2000333322130033) |
| `http_proxy.more_option.enable_path_normalize` | [http_proxy.more_option.enable_path_normalize](resources--proxy--reference--group-004.md#canonical-0203101203312302-2312202110311310-2220302001001003-0320023221131300-3001112322010233-1212200111231300-3222203231201223-2301302213220111) |
| `http_proxy.more_option.idle_timeout` | [http_proxy.more_option.idle_timeout](resources--proxy--reference--group-004.md#canonical-0112113002130301-2002113131330321-0132330230130300-3133311110212013-3200102133232212-3103213231203200-2121130310333101-3332133011113012) |
| `http_proxy.more_option.max_request_header_size` | [http_proxy.more_option.max_request_header_size](resources--proxy--reference--group-004.md#canonical-1310120112213001-2201330230223221-0020012310332112-2012013230030010-3312203223100101-3030122030323330-3130333110003021-1321322201333201) |
| `http_proxy.more_option.max_requests_per_connection` | [http_proxy.more_option.max_requests_per_connection](resources--proxy--reference--group-004.md#canonical-1002130211130201-2101032201222021-1120321121103302-1300220103010130-3100122213330011-2133111200121022-1121310331111122-2202110311000302) |
| `http_proxy.more_option.no_request_limit_per_connection` | [http_proxy.more_option.no_request_limit_per_connection](resources--proxy--reference--group-004.md#canonical-2013003000303112-2201123200223002-3133201302102000-3033010231020010-0210033222213323-0301020130232020-0101322331132030-1301132223330003) |
| `http_proxy.more_option.request_cookies_to_add` | [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-1210000331112233-3112003330210012-0230201313032113-0123320223110231-2301003312222303-3032133310110022-2302212333333231-3033003101121103) |
| `http_proxy.more_option.request_cookies_to_add.name` | [http_proxy.more_option.request_cookies_to_add.name](resources--proxy--reference--group-004.md#canonical-3102202110110102-0012310312222312-3333030331110311-3302131230122302-0211312030300201-3313231310101332-0313231013210101-1211101233123310) |
| `http_proxy.more_option.request_cookies_to_add.overwrite` | [http_proxy.more_option.request_cookies_to_add.overwrite](resources--proxy--reference--group-004.md#canonical-0203322032102113-3202110212331020-2201301320303113-3203320100333101-2121131123132332-0221202001110130-0222030111200311-3113022000101310) |
| `http_proxy.more_option.request_cookies_to_add.secret_value` | [http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-0032302311112022-3220210031202120-2100222202123221-3110333132333112-1110031030233223-3233100312232031-0032130020101300-1000002031133020) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-1010010223100223-0203100100122031-2300201230110212-0332010211032022-2111001132323003-2122122113130210-1112021023210010-1331200210011110) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-004.md#canonical-3331331233222231-0021031113232222-2111002202000012-2112002322323321-2320000303231120-2323112312132033-3112210201332202-1001011312302103) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-004.md#canonical-1010032322301332-3211332133013000-1303332212023012-3320331120001233-1003020302001033-1221120001200103-1000012133013212-0112011030212002) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-004.md#canonical-3310200010310031-0330112202113002-3203031210003220-0233121203200023-1013132132323131-0323001231331010-1200332010123130-1010000210300011) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-004.md#canonical-3021131101030110-2220002212111122-2111220200213300-1201102330123030-0112213023022100-3232202333202031-3211222223310122-2013100302113113) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-004.md#canonical-3220111300202111-1331000010032330-3032021200011123-0303133110102211-1200123231213011-0012011202002103-1333302102320303-1313321003132013) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-004.md#canonical-2211233212012220-2333132133213302-0321220200333222-3302300010320200-2133103300132103-1222023032333011-0330311130002031-3322022202002202) |
| `http_proxy.more_option.request_cookies_to_add.value` | [http_proxy.more_option.request_cookies_to_add.value](resources--proxy--reference--group-004.md#canonical-2111203323003033-0030313313112322-2201132003320111-3201033321221330-3023101110133332-1002300320131011-0210301200013301-2002113232202012) |
| `http_proxy.more_option.request_cookies_to_remove` | [http_proxy.more_option.request_cookies_to_remove](resources--proxy--reference--group-004.md#canonical-0120232230030030-2321001331333320-2021302031113212-3132332013031220-3332323000313301-1212230310331123-3013320321310111-3312000112021333) |
| `http_proxy.more_option.request_headers_to_add` | [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-1230012112232311-2322032002321033-2203021003032312-3332130230101301-3110110320203223-1023021203021101-0233010213021222-1302103230000110) |
| `http_proxy.more_option.request_headers_to_add.append` | [http_proxy.more_option.request_headers_to_add.append](resources--proxy--reference--group-004.md#canonical-3013220321331202-1030322313102032-2203300013232001-0112212112332122-0223101002302202-2221313003312030-1302223201031223-0212323003301231) |
| `http_proxy.more_option.request_headers_to_add.name` | [http_proxy.more_option.request_headers_to_add.name](resources--proxy--reference--group-004.md#canonical-1203012110003233-1113331323000333-3130231333031022-0102322122211300-2020331011200310-0333301012303131-2120333330313030-3102321001231111) |
| `http_proxy.more_option.request_headers_to_add.secret_value` | [http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-2300231120000222-2210203002022322-3020322312201133-2232223001000320-1202201021113203-2021320210113032-2130000233333201-3232300313103313) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-1031322130332031-2313202200012232-3312032121110312-1213002022212332-1112332000003313-1031100100223332-0201222221111003-2320233232230333) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-004.md#canonical-3113013301200113-2021031312203012-0022322200302100-0020312111223002-1031120031200211-1232020302301121-2012302332031111-2220030203211113) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-004.md#canonical-3231111013321022-1000303002003201-2100312113112020-0211213311002030-3330223130233302-3231112232331213-3222331202001300-2120311122110320) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-004.md#canonical-2013303323012021-2131220223212323-3120221321111313-2313020222111130-0221003320201211-1032000122202230-1230312022310210-3022000233320020) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-004.md#canonical-2311311000202303-3023321133211030-0112113010311230-1332312133230231-2130111202112000-3123031200113022-3122200313203010-2001000200300322) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-004.md#canonical-0012113100031233-3332000210230320-2303202321030130-1223201322011211-1330121030311232-1222103331130032-2313211013211210-2020021013101101) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-004.md#canonical-0310201103103212-1101311330010203-3332331312231012-0013322012220211-0213232212022301-0322200213103303-0003330120312123-1010132020302110) |
| `http_proxy.more_option.request_headers_to_add.value` | [http_proxy.more_option.request_headers_to_add.value](resources--proxy--reference--group-004.md#canonical-2003232011312130-3313211102303323-0310323312310323-0230112133100213-2033323010231020-2131132121133023-0331010232001333-3111100201002021) |
| `http_proxy.more_option.request_headers_to_remove` | [http_proxy.more_option.request_headers_to_remove](resources--proxy--reference--group-004.md#canonical-0123101101130310-1130221002002033-2220012002122330-3223202111300120-0333333200033323-3003032122220102-0221101233303313-0210333023012213) |
| `http_proxy.more_option.response_cookies_to_add` | [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-1110113233031013-0303232013113321-1111012232011023-0231212322222302-2311002212132112-3102031331333331-1020030303221002-3023103221320021) |
| `http_proxy.more_option.response_cookies_to_add.add_domain` | [http_proxy.more_option.response_cookies_to_add.add_domain](resources--proxy--reference--group-004.md#canonical-3013001312222112-1301311032110131-3120013310333030-2121033123320011-2023201201220213-3320201321221022-1113122020132030-2022003033110113) |
| `http_proxy.more_option.response_cookies_to_add.add_expiry` | [http_proxy.more_option.response_cookies_to_add.add_expiry](resources--proxy--reference--group-004.md#canonical-3232101222113000-0201002030012012-2131211312020101-0123201313133123-0130311313032102-2230001130101220-3122310131012033-3013210000333232) |
| `http_proxy.more_option.response_cookies_to_add.add_httponly` | [http_proxy.more_option.response_cookies_to_add.add_httponly](resources--proxy--reference--group-004.md#canonical-2300221132220032-1212303011022123-0102011133303233-2212221022123222-0003001022000033-2000130100210121-1022013010231130-2213221221301330) |
| `http_proxy.more_option.response_cookies_to_add.add_partitioned` | [http_proxy.more_option.response_cookies_to_add.add_partitioned](resources--proxy--reference--group-004.md#canonical-0000220000131020-0112002310222230-1312332010332001-3220230022310100-2332132203133132-1012330313310301-0311322023311130-0303301332131221) |
| `http_proxy.more_option.response_cookies_to_add.add_path` | [http_proxy.more_option.response_cookies_to_add.add_path](resources--proxy--reference--group-004.md#canonical-0023311121131322-3210023031020001-1011201132313100-0312212132010010-2123131221122001-3120323223130033-1332310300323022-2000232301231002) |
| `http_proxy.more_option.response_cookies_to_add.add_secure` | [http_proxy.more_option.response_cookies_to_add.add_secure](resources--proxy--reference--group-004.md#canonical-1333133032032232-2020023020031330-2002000023320230-1303000230212332-0120301210200303-1121000321113200-1311023211220102-1231131312230222) |
| `http_proxy.more_option.response_cookies_to_add.ignore_domain` | [http_proxy.more_option.response_cookies_to_add.ignore_domain](resources--proxy--reference--group-004.md#canonical-0221201230023312-0221323211011323-3332120332233110-2332022022123031-0122021132331132-2111021020131301-1102230013332100-0311030103022130) |
| `http_proxy.more_option.response_cookies_to_add.ignore_expiry` | [http_proxy.more_option.response_cookies_to_add.ignore_expiry](resources--proxy--reference--group-004.md#canonical-3301221122213231-1331033131130203-1323002020111003-3031230221203231-0311102313120022-2001212033133011-2000121011022112-1133021311323230) |
| `http_proxy.more_option.response_cookies_to_add.ignore_httponly` | [http_proxy.more_option.response_cookies_to_add.ignore_httponly](resources--proxy--reference--group-004.md#canonical-3021201202113032-3001303132200023-1000220310322101-0031023123032102-3231233121232202-2333023031120023-3120213203232020-2303132202332332) |
| `http_proxy.more_option.response_cookies_to_add.ignore_max_age` | [http_proxy.more_option.response_cookies_to_add.ignore_max_age](resources--proxy--reference--group-004.md#canonical-1220101022102300-1003133022313203-2003121220233210-0212233101212032-1221100331213001-1131010230211301-3033132101002103-3022323303022301) |
| `http_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [http_proxy.more_option.response_cookies_to_add.ignore_partitioned](resources--proxy--reference--group-004.md#canonical-2100101233323032-1130213211301030-0001300120122332-1310030330221022-2220012210311133-1330210133022130-0231030201313003-3000211013303303) |
| `http_proxy.more_option.response_cookies_to_add.ignore_path` | [http_proxy.more_option.response_cookies_to_add.ignore_path](resources--proxy--reference--group-004.md#canonical-3130232232002203-0101102011220030-0223203312101321-2133223023012313-1201330320002233-3231221133000232-2231230233021000-0220222132200313) |
| `http_proxy.more_option.response_cookies_to_add.ignore_samesite` | [http_proxy.more_option.response_cookies_to_add.ignore_samesite](resources--proxy--reference--group-004.md#canonical-3133123012001302-3003231110111001-2013202131123130-2110021011013110-0131123300023112-1131230123222203-2013312012310022-3223111003033310) |
| `http_proxy.more_option.response_cookies_to_add.ignore_secure` | [http_proxy.more_option.response_cookies_to_add.ignore_secure](resources--proxy--reference--group-004.md#canonical-1233333310313010-1110313202123232-3022121322303321-1321330103210013-2200223232312120-1000103133110022-3123132130001131-3210030332122332) |
| `http_proxy.more_option.response_cookies_to_add.ignore_value` | [http_proxy.more_option.response_cookies_to_add.ignore_value](resources--proxy--reference--group-004.md#canonical-2000011001011323-1300131221132110-3303122103213333-1011011021303021-2012212223232201-0003232211233331-1132133013322310-0020230200220121) |
| `http_proxy.more_option.response_cookies_to_add.max_age_value` | [http_proxy.more_option.response_cookies_to_add.max_age_value](resources--proxy--reference--group-004.md#canonical-2310032332322001-3220200203003332-1022010100011112-3220231322321131-1212130200012222-1000130121302203-3332313100123320-1002111212310323) |
| `http_proxy.more_option.response_cookies_to_add.name` | [http_proxy.more_option.response_cookies_to_add.name](resources--proxy--reference--group-004.md#canonical-2032122312200131-2111332013212203-1021331031300313-2313311211103332-3302230322232122-3022012230002121-0331330113000212-1233110320112012) |
| `http_proxy.more_option.response_cookies_to_add.overwrite` | [http_proxy.more_option.response_cookies_to_add.overwrite](resources--proxy--reference--group-004.md#canonical-3123302332311112-0231212123321032-3300030101300110-0210112033110321-1130233333233013-3210230323301131-3320313300313120-2012330330210030) |
| `http_proxy.more_option.response_cookies_to_add.samesite_lax` | [http_proxy.more_option.response_cookies_to_add.samesite_lax](resources--proxy--reference--group-004.md#canonical-0323000230312231-2003211031102031-1103001100233000-1202311101333212-3020232303303111-2003230311033002-2012113111100201-2122332202210211) |
| `http_proxy.more_option.response_cookies_to_add.samesite_none` | [http_proxy.more_option.response_cookies_to_add.samesite_none](resources--proxy--reference--group-004.md#canonical-0333030220112210-2021023111232103-0122120132213233-2033202133101020-1110312213101233-3331210333132111-3133213211111101-0130332211011103) |
| `http_proxy.more_option.response_cookies_to_add.samesite_strict` | [http_proxy.more_option.response_cookies_to_add.samesite_strict](resources--proxy--reference--group-004.md#canonical-1111222130212203-0131223230132311-3010311232233220-3111231010232322-3301221010120130-3211003012131223-0303220300123312-0302102210001222) |
| `http_proxy.more_option.response_cookies_to_add.secret_value` | [http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-1300211320131301-0203023131301210-1303101312013001-1112303003212202-0120331100100003-2001013102011323-2100133111233203-3101011302130220) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-005.md#canonical-3333123130100000-2100030320313133-1023322130300210-2110122230231023-1003110310131130-2231022121323302-0201202132100000-0313013132223032) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-005.md#canonical-1322101213233323-0021003222021021-0313103233202203-0312133311320113-0131331130102022-0121103111122102-0001233021330003-3030033331111011) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-005.md#canonical-3310112200003233-0330310212312012-1112233102021320-3111222012311013-0301010202320300-1023021312220311-0032032112202232-0312221232123101) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-005.md#canonical-0121031200311302-0221100023112130-1130220012202102-0122110212010003-3322131220212233-1210200213131113-1113030212001213-2303100300012211) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-005.md#canonical-2100103200111022-1232311232303131-3302200200002321-0030031313122311-0130110020122031-3223311213201232-1103222003022021-3112313133223303) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-005.md#canonical-0332120021321111-2101201130210302-1221133321311313-1120312111011021-0302231010213031-0211111202131032-2210310133311312-1322333013301002) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-005.md#canonical-1232330320220220-0010100010231021-1003301110333302-2103220211230100-2203023302020221-0020320032110313-2301113230320311-1110331111003023) |
| `http_proxy.more_option.response_cookies_to_add.value` | [http_proxy.more_option.response_cookies_to_add.value](resources--proxy--reference--group-004.md#canonical-2313220302022013-3032113230332112-3111022103330310-1100300002332301-1102033220111211-1020122231321022-0100032223213231-0133023300113320) |
| `http_proxy.more_option.response_cookies_to_remove` | [http_proxy.more_option.response_cookies_to_remove](resources--proxy--reference--group-004.md#canonical-2100130031301331-1211022003131011-2331122102111202-1033032130122033-0002322110223220-1221031322133213-0311213110022020-3030322110233211) |
| `http_proxy.more_option.response_headers_to_add` | [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-005.md#canonical-1333331213213220-1313002212301213-0311101201203133-0102130131011013-3101203013301303-1021301033120000-3213210013232103-1221322223212331) |
| `http_proxy.more_option.response_headers_to_add.append` | [http_proxy.more_option.response_headers_to_add.append](resources--proxy--reference--group-005.md#canonical-3303320120322021-1301322213313021-1233212312311013-0323322032320113-1212323001011330-0302021030210220-2201132011212212-1312031111030303) |
| `http_proxy.more_option.response_headers_to_add.name` | [http_proxy.more_option.response_headers_to_add.name](resources--proxy--reference--group-005.md#canonical-3313221033320223-3302221133110232-3022030031120130-3320111101111302-3122332312323210-1122212303010201-0120121100301133-2101033112331321) |
| `http_proxy.more_option.response_headers_to_add.secret_value` | [http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-005.md#canonical-3101123131023332-1232313122120131-2221032013303333-0122202011000103-3213012213321111-1320331103202313-1202132103012100-2111002322232122) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-005.md#canonical-0121133002111022-1002000221032031-2000113021022120-3102302022222122-1120232323303023-1333312301201303-1230031121122202-1210320112231313) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-005.md#canonical-2110212123001032-1101302133002011-2102223322230132-0323011323112000-3312012012213031-2230213230111003-3221111023332321-0310132232232210) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-005.md#canonical-3231122103112003-3313302033100331-1321100220000233-1013220000002203-1331121230000032-0110231310013311-0002211203010200-3310312003113222) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-005.md#canonical-3013103000030102-0023123021032023-1321310333023132-3122032233111331-2202023032013132-3212233000312313-2230002200220123-0223012100001311) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-005.md#canonical-2200213112232300-3313331032330221-1202320001001221-0020102133011231-1300131333022001-0220201012233223-1332132221012203-3003301310020330) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-005.md#canonical-1020230230001113-2222202233133031-3223103120332120-3002322321003020-2102310212230102-3213032033033320-3232300030202000-0122013112300020) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-005.md#canonical-0330010312031012-1031211200003032-1113101120121000-0103322101000003-1123231100201303-1330122212321002-2013223333021300-3311313330112122) |
| `http_proxy.more_option.response_headers_to_add.value` | [http_proxy.more_option.response_headers_to_add.value](resources--proxy--reference--group-005.md#canonical-2230222201003303-1330320021123112-0020112223121310-0231110111313132-3022211301313222-3131033301313020-0031121301023313-2013321012033022) |
| `http_proxy.more_option.response_headers_to_remove` | [http_proxy.more_option.response_headers_to_remove](resources--proxy--reference--group-004.md#canonical-2002112100110131-0103300002133100-0010332111230100-1233301232121321-2232103211230233-3002313103012231-1011131310023122-0323031120030323) |
| `id` | [ID](resources--proxy--reference--group-001.md#canonical-3021310010301023-0303200031113011-3232020220213113-2313303000232303-0321312220212333-1232233203313202-1011313201333211-2331300002330210) |
| `labels` | [labels](resources--proxy--reference--group-001.md#canonical-1232113033323320-1321330113300230-0330200023130232-3303322211233302-2201031330123101-2013031103223033-3130300121311333-2302011131113112) |
| `name` | [name](resources--proxy--reference--group-001.md#canonical-0113111221131320-3121313220033313-0220213200023113-1212300333123000-2133103200102203-2133102333220320-2113223333133332-1212032200200110) |
| `namespace` | [namespace](resources--proxy--reference--group-001.md#canonical-0011102331302022-0032213200233212-3102011000121310-3011130113223113-3032031001100002-0102222132032110-3111010021020033-3123112122111210) |
| `no_forward_proxy_policy` | [no_forward_proxy_policy](resources--proxy--reference--group-005.md#canonical-1302311121130102-2131203200103320-1303323313202210-0113023213112321-0121211321110201-2311112312303120-1201103232201010-1112001213313132) |
| `no_interception` | [no_interception](resources--proxy--reference--group-005.md#canonical-3332023313312012-3012310221032321-3101200211301200-2222112211033303-1033311031230203-0102333033323201-3101321210231120-3233331221000130) |
| `site_local_inside_network` | [site_local_inside_network](resources--proxy--reference--group-005.md#canonical-3121123313321101-0001303002003321-2112130332133322-1102033212123313-1303202113031023-1012333112101221-2212211032231111-0323301233320300) |
| `site_local_network` | [site_local_network](resources--proxy--reference--group-005.md#canonical-3220322333321001-2032010012232201-0231323320030211-1123130120121020-2021030112012033-2213000303013330-0001000213111211-3213131133222023) |
| `site_virtual_sites` | [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-0333011221000001-0111132130112111-0231020012230322-3301122013323311-0123011330302122-0120213032111302-3301221321211312-3112303021033110) |
| `site_virtual_sites.advertise_where` | [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-2100110201320221-3330213111022300-3111200010313311-0331110213021031-2123233022023023-3121012211012322-3222213131320121-0203133310233020) |
| `site_virtual_sites.advertise_where.port` | [site_virtual_sites.advertise_where.port](resources--proxy--reference--group-005.md#canonical-2313331113002231-0223313032111320-2121123130113111-2112002101331020-3013211030130110-0320211001020010-0312303030130123-3100331002223113) |
| `site_virtual_sites.advertise_where.site` | [site_virtual_sites.advertise_where.site](resources--proxy--reference--group-005.md#canonical-3322321212133033-3111230213233223-2022230112131332-3202102321032031-0110133333110320-3302032210313332-0303331212200032-3310010110232220) |
| `site_virtual_sites.advertise_where.site.ip` | [site_virtual_sites.advertise_where.site.ip](resources--proxy--reference--group-005.md#canonical-3020123232013323-3210201103102201-2103012332120020-3002111323022032-2000010312231311-3232103100233123-1120022203113223-0110122221133311) |
| `site_virtual_sites.advertise_where.site.network` | [site_virtual_sites.advertise_where.site.network](resources--proxy--reference--group-005.md#canonical-3332301132001121-2330200212011001-1133312132310230-2123020113231313-3210101012031113-1322322130033302-3210222023131233-1023120311110100) |
| `site_virtual_sites.advertise_where.site.site` | [site_virtual_sites.advertise_where.site.site](resources--proxy--reference--group-005.md#canonical-2103321233003120-2201200120122210-3131312013011000-2031220232012110-0201110232321202-3133001211302030-3132303203030301-1123130121013213) |
| `site_virtual_sites.advertise_where.site.site.name` | [site_virtual_sites.advertise_where.site.site.name](resources--proxy--reference--group-005.md#canonical-3121310320130220-2331301102112020-1220113322032121-0031302332011122-0020033121333001-0212302110201103-2231100123111333-0232333022132113) |
| `site_virtual_sites.advertise_where.site.site.namespace` | [site_virtual_sites.advertise_where.site.site.namespace](resources--proxy--reference--group-005.md#canonical-1010221100200201-2213001313323102-2121021201110331-3202330100130002-3311023221213023-3111003032100031-2330030133310310-1111130201031233) |
| `site_virtual_sites.advertise_where.site.site.tenant` | [site_virtual_sites.advertise_where.site.site.tenant](resources--proxy--reference--group-005.md#canonical-3133131122003012-1302032100213122-0310302322101230-1222213220330320-1110231123323312-2230211231012031-0120313203331100-2112002210310032) |
| `site_virtual_sites.advertise_where.use_default_port` | [site_virtual_sites.advertise_where.use_default_port](resources--proxy--reference--group-005.md#canonical-0122120302113102-2312133200030003-0320023221321311-1211202100300300-3232310132131300-3003202022122110-0012312121103301-1103211230011310) |
| `site_virtual_sites.advertise_where.virtual_site` | [site_virtual_sites.advertise_where.virtual_site](resources--proxy--reference--group-005.md#canonical-0212121332300203-2110123100212012-1102113201011032-2333332232213032-3333021310230323-0012103113020323-1322021231120310-2011323023331022) |
| `site_virtual_sites.advertise_where.virtual_site.network` | [site_virtual_sites.advertise_where.virtual_site.network](resources--proxy--reference--group-005.md#canonical-1023312001111300-1100323013130010-3132301100320232-2101021300021111-2000311030233303-3122230221012130-3023003123201320-2330313321010010) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site` | [site_virtual_sites.advertise_where.virtual_site.virtual_site](resources--proxy--reference--group-005.md#canonical-0221220331332212-2030311010233103-3111323303330033-1303122020131120-3012203103332302-1021100230031222-2303032031301320-0012210232123021) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.name` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.name](resources--proxy--reference--group-005.md#canonical-0333011010230103-0210213231232323-3232310311111101-0013203301102123-2231332122102002-1113123133120112-3300230110310210-3311023102322321) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.namespace` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.namespace](resources--proxy--reference--group-005.md#canonical-3010310031323201-1202312201220132-3102011121220122-1000310331002102-2111221323121221-3200312222231333-3230132003111322-1213110232113303) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.tenant` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.tenant](resources--proxy--reference--group-005.md#canonical-1232220200003130-1331033010203112-1233030100221122-0301122111200133-0110300231213002-1203020023003011-0203111220123131-3223201233230230) |
| `timeouts` | [timeouts](resources--proxy--reference--group-005.md#canonical-3100221132113223-0221031321312001-3112300110111021-3322203020203310-3012010232232321-0222122130021011-3230012320013131-0303023223020122) |
| `timeouts.create` | [timeouts.create](resources--proxy--reference--group-005.md#canonical-1103120221321111-3002033010233002-3110210222233021-0231311012103102-1123033223001123-3001110112012100-2021320110011201-0300330321102300) |
| `timeouts.delete` | [timeouts.delete](resources--proxy--reference--group-005.md#canonical-1221303323102020-1110010213210303-2331222111130233-2013220310022102-2323220301220031-0022100223230213-3211022101003231-2032033220301300) |
| `timeouts.read` | [timeouts.read](resources--proxy--reference--group-005.md#canonical-2320022221231000-1231021201033201-1023021213332301-1230312111300001-0323033312020123-1023201001033331-1133021011021300-2032110033212110) |
| `timeouts.update` | [timeouts.update](resources--proxy--reference--group-005.md#canonical-2232312313013221-0311031132223112-0030011111200223-1103230231301131-3200311300003133-1212112321110313-3121322030010203-3133001203330223) |
| `tls_intercept` | [tls_intercept](resources--proxy--reference--group-005.md#canonical-2300212110133011-2303221131021333-1221330301300022-0231220012001111-0213123123321222-0000030203023101-1203222021230313-0321231013011230) |
| `tls_intercept.custom_certificate` | [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-0130022112311203-0002331003312132-1033021013303233-2012221301021003-0123002202003302-3320330131212202-3310321213331132-0333330300202030) |
| `tls_intercept.custom_certificate.certificate_url` | [tls_intercept.custom_certificate.certificate_url](resources--proxy--reference--group-005.md#canonical-2201022013301311-0213100320031220-1002032113200202-1310102202002013-2330221000130002-3323111100231222-0301212033232303-1222032322333202) |
| `tls_intercept.custom_certificate.custom_hash_algorithms` | [tls_intercept.custom_certificate.custom_hash_algorithms](resources--proxy--reference--group-005.md#canonical-0010221002113203-1320123333021111-2131113310302131-2223121311002101-0211030313032012-1213232231303030-0332013011332201-2012220031313001) |
| `tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` | [tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms](resources--proxy--reference--group-005.md#canonical-1123210200321201-0103130010111311-3011202013331200-2330210113233011-2232220100323300-0310023113112211-2121012020332202-2231231231203112) |
| `tls_intercept.custom_certificate.description_spec` | [tls_intercept.custom_certificate.description_spec](resources--proxy--reference--group-005.md#canonical-3133130011332203-2030002102133223-2312210333013330-3000211323231223-2002111120220221-1101123210002112-0201001310000103-1232113321303100) |
| `tls_intercept.custom_certificate.disable_ocsp_stapling` | [tls_intercept.custom_certificate.disable_ocsp_stapling](resources--proxy--reference--group-005.md#canonical-0211213312302113-3322133023012121-0200203231303133-3323110331113331-3302203033022223-1013333112320101-2000123033001203-1130200012111122) |
| `tls_intercept.custom_certificate.private_key` | [tls_intercept.custom_certificate.private_key](resources--proxy--reference--group-005.md#canonical-2200200320331121-1333131312130012-3322323323023132-0203231323032321-3111332333001312-1230102301032332-3202021313322200-0330202120220030) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info](resources--proxy--reference--group-005.md#canonical-1303000210122130-0033031020131202-2223232022111303-0101023310321310-0103201102201023-2200101021211100-3303302032331312-2300211330320103) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-005.md#canonical-2311302231311013-1313022212223011-3320300010200122-2103232330000310-0211000232310332-0220232233231302-2110310213001311-0310203102110013) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.location](resources--proxy--reference--group-005.md#canonical-2302122000031322-1230230313132022-1102003222122101-2002320132312312-2002103012330122-3310002312111221-1031233102023120-1020121323132022) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider](resources--proxy--reference--group-005.md#canonical-3310333120321022-1021132212211032-2331030010132110-0030113122233032-3102312133000121-1111213102300120-0012303330102203-2333331113101201) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info` | [tls_intercept.custom_certificate.private_key.clear_secret_info](resources--proxy--reference--group-005.md#canonical-3211200113130130-2323212132131212-2010223130002213-1210211300333223-2300220012101110-2313201212010302-2010300002320333-3221010210032210) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` | [tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref](resources--proxy--reference--group-005.md#canonical-3220233012210213-0003323133131222-2021100300232131-3201323313110031-0330213312331311-0113212001233312-3233213220032102-3010222023312203) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info.url` | [tls_intercept.custom_certificate.private_key.clear_secret_info.url](resources--proxy--reference--group-005.md#canonical-2213022103112120-1123211032221101-0033212213001120-1021332101023210-2031003323031010-1220310103320001-0333000210100022-0231120000011000) |
| `tls_intercept.custom_certificate.use_system_defaults` | [tls_intercept.custom_certificate.use_system_defaults](resources--proxy--reference--group-005.md#canonical-1110203332211003-3020032332111020-3013002222220323-1032201121213320-0022011103023313-3231220113303122-2211023121012311-2000111032121001) |
| `tls_intercept.enable_for_all_domains` | [tls_intercept.enable_for_all_domains](resources--proxy--reference--group-005.md#canonical-3303033331133302-0302031220133100-3101211302231333-0133120110321330-2112331230232232-2002001213311121-2122333311103001-3031221331320202) |
| `tls_intercept.policy` | [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-3011100031221211-1203111210222303-3201101010103122-1012213130310031-1131102203013211-1200222033133233-3123020100000232-3001100100320300) |
| `tls_intercept.policy.interception_rules` | [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-3312120023110110-1021033230023313-3012032301110231-0312332203030120-1302020330230020-2013110232000331-3202313011323212-0202011032302323) |
| `tls_intercept.policy.interception_rules.disable_interception` | [tls_intercept.policy.interception_rules.disable_interception](resources--proxy--reference--group-005.md#canonical-0131220000203202-2012032110322310-1130312113203001-3202200031223021-2111112001303320-3031213121103212-2212320133201212-1022020211033011) |
| `tls_intercept.policy.interception_rules.domain_match` | [tls_intercept.policy.interception_rules.domain_match](resources--proxy--reference--group-005.md#canonical-0222232012012321-1012022110323022-2130020212210131-3222232100323101-0110033332333101-1213131322200231-2331333113132111-1121000230312002) |
| `tls_intercept.policy.interception_rules.domain_match.exact_value` | [tls_intercept.policy.interception_rules.domain_match.exact_value](resources--proxy--reference--group-005.md#canonical-3201000310002110-3111121001031331-1220221031121301-1313022123030033-0020000223300020-3023033002213233-3030301112233131-0123033223222203) |
| `tls_intercept.policy.interception_rules.domain_match.regex_value` | [tls_intercept.policy.interception_rules.domain_match.regex_value](resources--proxy--reference--group-005.md#canonical-1232322302030023-2132213220131130-0032303303133001-2311010030202200-1301230211311323-2222010132121103-1023231021323013-2031332131031030) |
| `tls_intercept.policy.interception_rules.domain_match.suffix_value` | [tls_intercept.policy.interception_rules.domain_match.suffix_value](resources--proxy--reference--group-005.md#canonical-0230330021220112-1120022222202102-1300102121033322-0231001100213030-1111320121100221-3120221233230210-0011102121321221-3002121303032233) |
| `tls_intercept.policy.interception_rules.enable_interception` | [tls_intercept.policy.interception_rules.enable_interception](resources--proxy--reference--group-005.md#canonical-3013000203331130-3003021302221020-2112333323201001-0011311321113103-0002221001030022-2312213230213010-0312201033202321-3231302301120002) |
| `tls_intercept.trusted_ca_url` | [tls_intercept.trusted_ca_url](resources--proxy--reference--group-005.md#canonical-2302310120010303-0110132203322000-0001021231100100-2330023331030131-0023031130300232-3200023320232110-1113013233133323-1310310320311202) |
| `tls_intercept.volterra_certificate` | [tls_intercept.volterra_certificate](resources--proxy--reference--group-005.md#canonical-2121122002311101-2001321331222012-1332131220210100-0323202103032100-2011030332222001-3112331133023222-1302120010030002-0311103030022013) |
| `tls_intercept.volterra_trusted_ca` | [tls_intercept.volterra_trusted_ca](resources--proxy--reference--group-005.md#canonical-1013120230302223-0033003022013000-2223010222321303-2021123331333122-1122330221121323-1302122111221013-3302102203012022-0003002030322113) |

<a id="canonical-0003323102113131-0103011023111231-1113001332012010-0301232233010331-0130302132021322-2321100013200013-0033323321132110-1121322121220021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_forward_proxy_policies` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- active_forward_proxy_policies

<a id="canonical-0221211131033121-3202012010201332-1113103330200011-0102030311133213-2111131301330011-2301322012102332-2312002132210112-2020233302023120"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_forward\_proxy\_policies, no\_forward\_proxy\_policy; Default:
no\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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

OneOf alternatives in this subsection:

- [active_forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-0221211131033121-3202012010201332-1113103330200011-0102030311133213-2111131301330011-2301322012102332-2312002132210112-2020233302023120)
- [no_forward_proxy_policy](resources--proxy--reference--group-005.md#canonical-1302311121130102-2131203200103320-1303323313202210-0113023213112321-0121211321110201-2311112312303120-1201103232201010-1112001213313132)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333131301322103-1321010022011220-2100110222021130-2100102001332013-3232230221222201-3023322200311232-2123302212301210-3001323221230012"></a>

### Direct properties for `active_forward_proxy_policies`

- [forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-1230113222323330-0220310102200332-1332230021002120-2103311313110211-0300001122200322-1223330220000033-3313100130133003-1331303003011122): complete subsection reference.

<a id="canonical-1230113222323330-0220310102200332-1332230021002120-2103311313110211-0300001122200322-1223330220000033-3313100130133003-1331303003011122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_forward_proxy_policies.forward_proxy_policies` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [active_forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-0003323102113131-0103011023111231-1113001332012010-0301232233010331-0130302132021322-2321100013200013-0033323321132110-1121322121220021)
- active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-2023203231230201-1323322021130311-2122030003302322-1112013201113131-1113001310002102-2301303302221010-3220122013332323-1131233022222112"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213132103312233-2031322312012221-1122132201123110-0322222001132121-3130013111112110-3123021103020211-0121130033121202-0312303330313300"></a>

### Direct properties for `active_forward_proxy_policies.forward_proxy_policies`

<a id="canonical-0122323000223013-2313013011321230-0113223123021133-2322023301100030-1333220110231231-2000220323123311-3311333223201322-0111313031013223"></a>

#### `active_forward_proxy_policies.forward_proxy_policies.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2233210132011010-1312102310211131-2210033133200020-0232002212120023-3101311001132121-1003220003230002-0100201300121203-0002231320332323"></a>

<a id="canonical-3131000132300320-1210310213230313-2113122312203122-2023002203232211-3112013030001131-2332012223121223-2310232231303322-0102310011303323"></a>

#### `active_forward_proxy_policies.forward_proxy_policies.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1203123310333233-3001103002333323-3312300302233220-0033202323122113-1012332321332231-1231020100123130-3030000221032013-1123033321102313"></a>

<a id="canonical-3031131320312323-2310231023010133-1103310211332311-3230100100001331-2311121313231121-0330103003112011-2101333230223112-1022021230030310"></a>

#### `active_forward_proxy_policies.forward_proxy_policies.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1300300333220210-2321100332123100-3100222323020132-2000023232233210-0122031023212311-0203230120332013-3020033130003013-3323122123123232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `do_not_advertise` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- do_not_advertise

<a id="canonical-0312133310103133-2021121111223230-0202102102001031-1211202103032101-3121123313123212-0031113301200102-0123101311303022-3333312123131010"></a>

Type: `["object", {}]`. Optional.

\[OneOf: do\_not\_advertise, site\_virtual\_sites\] Configuration parameter for do not advertise.

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

- [do_not_advertise](resources--proxy--reference--group-001.md#canonical-0312133310103133-2021121111223230-0202102102001031-1211202103032101-3121123313123212-0031113301200102-0123101311303022-3333312123131010)
- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-0333011221000001-0111132130112111-0231020012230322-3301122013323311-0123011330302122-0120213032111302-3301221321211312-3112303021033110)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
do_not_advertise = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- dynamic_proxy

<a id="canonical-0322013013112103-1110312232333000-0001202300202230-1130120121002131-2033132100020232-0121002303211031-1313302121010231-1120332102001132"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dynamic\_proxy, http\_proxy\] Configuration parameter for dynamic proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("disable_dns_masquerade",
    "enable_dns_masquerade"),
  validators.ConflictingObjectAttributes("http_proxy",
    "https_proxy"),
  validators.ConflictingObjectAttributes("http_proxy",
    "sni_proxy"),
  validators.ConflictingObjectAttributes("https_proxy",
    "sni_proxy")}
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
  "x-ves-oneof-field-dns_choice": "[\"disable_dns_masquerade\",\"enable_dns_masquerade\"]",
  "x-ves-oneof-field-proxy_choice": "[\"http_proxy\",\"https_proxy\",\"sni_proxy\"]"
}
```

OneOf alternatives in this subsection:

- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-0322013013112103-1110312232333000-0001202300202230-1130120121002131-2033132100020232-0121002303211031-1313302121010231-1120332102001132)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-0023030222010311-3300322230002013-2001102311323223-1122022322221322-2020120013210033-0032330013323132-3222112033103212-3223122333232132)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dynamic_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102121320010113-2222023031002320-2303213233333200-2021011312200221-3031122312332003-1022111031022132-3123300321133231-2120230312323101"></a>

### Direct properties for `dynamic_proxy`

- [disable_dns_masquerade](resources--proxy--reference--group-001.md#canonical-2121102311231213-2311003112320322-1121001322011303-0100312301001031-1020032301132203-3023111013301030-1020021222031100-1202311330120120): complete subsection reference.

<a id="canonical-2330001000121003-1313222001030002-3222231123332003-1000131020332023-1330020320110022-0232003030231202-1102133333313000-1303210230001010"></a>

<a id="canonical-3203323232012023-2200002001022313-2002210303230203-3012101033312020-1231132303301000-1213023303012330-0213221210131111-1001122021022131"></a>

#### `dynamic_proxy.domains` property

Type: `["list", "string"]`. Optional.

A list of Domains to be proxied. Wildcard hosts are supported in the suffix or prefix form

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [enable_dns_masquerade](resources--proxy--reference--group-001.md#canonical-2221013333003111-1001100121300013-3300223323312232-1321103333110323-0321130111110133-0320232201021232-1000321200100103-3131012122222030): complete subsection reference.

- [http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103): complete subsection reference.

- [https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203): complete subsection reference.

- [sni_proxy](resources--proxy--reference--group-003.md#canonical-3112020001021313-3001013121113231-2312300311233300-0101002320313231-2321103212111010-1200232211133121-0211133000013012-3022312230220132): complete subsection reference.

<a id="canonical-2121102311231213-2311003112320322-1121001322011303-0100312301001031-1020032301132203-3023111013301030-1020021222031100-1202311330120120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.disable_dns_masquerade` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- dynamic_proxy.disable_dns_masquerade

<a id="canonical-1332321312013121-3220123333011333-3221202222311303-2112202322103312-0312330113003310-3121130303313003-1130030300221313-0001033032022133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable DNS masquerade.

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
disable_dns_masquerade = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221013333003111-1001100121300013-3300223323312232-1321103333110323-0321130111110133-0320232201021232-1000321200100103-3131012122222030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.enable_dns_masquerade` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- dynamic_proxy.enable_dns_masquerade

<a id="canonical-0320100121210033-0212112313013000-2020201032102121-1120230203133112-1303120032101310-2223323331212223-2111212322301113-0320101223330302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable DNS masquerade.

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
enable_dns_masquerade = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- dynamic_proxy.http_proxy

<a id="canonical-0030322113101110-2112322112022013-3132302301210231-1201130310013303-1311221321210103-0002103231302132-0001202001113232-2002130001233222"></a>

Type: `"object"`. single nested block, Optional.

Dynamic HTTP Proxy Type. Parameters for dynamic HTTP proxy.

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
http_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030310220121120-2331011131102131-0101003031023030-2012003331032110-0303311313313220-2322300301110122-1213012113012020-1203213331013133"></a>

### Direct properties for `dynamic_proxy.http_proxy`

- [more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232): complete subsection reference.

<a id="canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- dynamic_proxy.http_proxy.more_option

<a id="canonical-2012331230003322-1110102322230203-3230302030120130-0300223100313302-1332013220321120-2203020012201113-2022121111103102-1302030003300030"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to define a route.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("max_requests_per_connection",
    "no_request_limit_per_connection")}
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
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-strict_sni_host_header_check_choice": "[]"
}
```

Terraform syntax:

```terraform
more_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003002123301313-2012331123202012-0213122000210320-3312303121232310-0101313011020231-0203323022313032-3010100333323321-0221212312203121"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option`

- [buffer_policy](resources--proxy--reference--group-001.md#canonical-0123200322031111-0030222301032211-0033120333103232-3013100311131132-1213322330211003-0213213110100201-0220103031330103-2320302122132211): complete subsection reference.

- [compression_params](resources--proxy--reference--group-001.md#canonical-2200031212103131-3310131223321321-1013221221132310-2010222013231121-1313201321220130-0202321300102200-0211121130203102-2332023130201030): complete subsection reference.

<a id="canonical-0003033023110332-3200232020113131-3131332333111313-3333031011302003-2102000313130223-0130330013033313-0133031120320333-3320233310002111"></a>

<a id="canonical-2210220321103211-3012002120023020-2333331313120212-1332020033332120-1121032230103221-1310310311032333-2002102121210200-3100101230222033"></a>

#### `dynamic_proxy.http_proxy.more_option.custom_errors` property

Type: `["map", "string"]`. Optional.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP
responses. Specific response code takes preference when both response code and response code class
matches for a request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"ranges\":[[3,3],[4,4],[5,5],[300,599]],\"type\":\"uint32-string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.uint32.ranges\":\"3,4,5,300-599\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"65536\",\"ves.io.schema.rules.map.values.string.uri_ref\":\"true\"},\"values\":{\"format\":\"uri-reference\",\"maxLength\":65536,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "ranges": [
        [
          3,
          3
        ],
        [
          4,
          4
        ],
        [
          5,
          5
        ],
        [
          300,
          599
        ]
      ],
      "type": "uint32-string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "65536",
      "ves.io.schema.rules.map.values.string.uri_ref": "true"
    },
    "values": {
      "format": "uri-reference",
      "maxLength": 65536,
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
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  }
}
```

<a id="canonical-0301311113200030-0232132021033021-0011102023320012-3003110313222210-0130023233223220-3121302033301221-1300132300201010-0301011320001202"></a>

<a id="canonical-1222030020302332-1210310221030132-0311111132202222-2301013033032312-2013322022013020-1320203111112002-3103331032020102-1010122320000130"></a>

#### `dynamic_proxy.http_proxy.more_option.disable_default_error_pages` property

Type: `"bool"`. Optional.

Disable the use of default F5XC error pages.

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

- [disable_path_normalize](resources--proxy--reference--group-001.md#canonical-3010323103131011-2202132221302001-3230101000320102-1213223032032320-1331201133331230-0321310100011302-0131321202201220-0211331012332312): complete subsection reference.

- [enable_path_normalize](resources--proxy--reference--group-001.md#canonical-2213331021323322-3030210022220213-2001313310032320-2310212103302212-2123120201001231-2223111022333000-1233101102313112-0132110222331132): complete subsection reference.

<a id="canonical-2210031201222033-1202003032122010-3021013311013113-0103031002202011-2210111100012321-1230303011301310-0230303013203130-0133311133303212"></a>

<a id="canonical-2330011300210233-1300200103012011-3333201202201023-3230333323033232-0022020322102030-3112223201302330-1222101213011320-2132212203112312"></a>

#### `dynamic_proxy.http_proxy.more_option.idle_timeout` property

Type: `"number"`. Optional.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with an HTTP 504 (Gateway Timeout) error code if no upstream response
header has been received, otherwise the stream is reset.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(3600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  }
}
```

<a id="canonical-1002223310210022-3021220301132011-0220121332012130-3310331010310113-1221220022201210-0313103131222302-0323221231003021-2232231011031121"></a>

<a id="canonical-2300031321230212-3002101203212331-1233001023220322-1303003131130203-2031131332123013-1121212222101101-3313131221100103-1210313020110300"></a>

#### `dynamic_proxy.http_proxy.more_option.max_request_header_size` property

Type: `"number"`. Optional.

The maximum request header size for downstream connections, in KiB. An HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  }
}
```

<a id="canonical-0300220120123310-3222131032332031-1023332201020312-2101030322212001-1000131201232201-0010320221131221-3122121013032013-2002213322133102"></a>

<a id="canonical-1130030322322013-1203300012311031-2102121203322301-1013332133002333-1003331312133302-2112021120303011-3002302332111200-0232223313310021"></a>

#### `dynamic_proxy.http_proxy.more_option.max_requests_per_connection` property

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [no_request_limit_per_connection](resources--proxy--reference--group-001.md#canonical-1300132130201131-0303023011210331-2213112112010120-1033210011003131-0103131233101201-1213132101102322-3112320303333013-0033212231102002): complete subsection reference.

- [request_cookies_to_add](resources--proxy--reference--group-001.md#canonical-2202201333112001-1113233232320310-1221001002302311-0312120031201222-3230230022213231-3133322233323010-3332110221230120-0230320130222220): complete subsection reference.

<a id="canonical-3232313213010203-3130310023232323-1122332030333330-0022222203323323-2023330101023222-3223330233102220-0211301023330330-1221002213222331"></a>

<a id="canonical-3013023322211210-0221131100010121-0200100320233212-2100201103100100-2323121100232101-1102333233220212-0313212310110102-3003120330300021"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](resources--proxy--reference--group-002.md#canonical-1002022321310321-1002330233323203-2202030110100111-2310321303122121-2011321201120202-2233101100223123-3311230010321231-3331122113232132): complete subsection reference.

<a id="canonical-0313200131120023-1120113022320011-3033310100301112-1301010002313023-2020302100122010-1302313223312232-3012210330310031-2122220002213100"></a>

<a id="canonical-1121312030103230-1020213122012330-3323003310330113-1230112132103123-1313102202222300-2101132121233323-2302003032231223-0213123021333313"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002): complete subsection reference.

<a id="canonical-0011201022031310-1032000222311323-1202322200110210-2312323321000203-2003221333333022-1112022230200030-3222212111310021-3010230222030223"></a>

<a id="canonical-3002321103021322-2202212230210033-1322210223302031-1202331032012202-0221333211032001-0030011223132023-1221321113332302-1201023200022023"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_remove` property

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](resources--proxy--reference--group-002.md#canonical-1212230120213332-0202331232012331-1300210112213133-3300202222202130-1033203202131323-2032233032303103-2021201331031231-3300133123003000): complete subsection reference.

<a id="canonical-0320233002132310-1220121303212030-2123030301003223-0111002313320313-3310122023120222-1231111203003133-0110233103101310-3201302013210102"></a>

<a id="canonical-3033212110132020-3001001031011212-1203313202032333-1223000113310133-0321020021112101-3123121113211030-1201120111213301-1230210201133020"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0123200322031111-0030222301032211-0033120333103232-3013100311131132-1213322330211003-0213213110100201-0220103031330103-2320302122132211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.buffer_policy` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- dynamic_proxy.http_proxy.more_option.buffer_policy

<a id="canonical-2100310220321112-3301211120113113-2022003303013332-2211132133233021-1131001132202222-1003212130210031-2013010200012030-3003232333302203"></a>

Type: `"object"`. single nested block, Optional.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer
configuration takes precedence.

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
buffer_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100230301321031-2001103312230033-0010211210110300-2200102203221301-1201221110323013-2000221333021221-2131213021331130-3301030230311330"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.buffer_policy`

<a id="canonical-1312211323203300-2303101221313033-3103113102032222-2300100300131210-2101201221320213-1232032220000002-1113232200202032-3132210311103211"></a>

#### `dynamic_proxy.http_proxy.more_option.buffer_policy.disabled` property

Type: `"bool"`. Optional.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="canonical-2221131131333130-2222313200030313-0313031123032023-3012102102113210-3002201023322331-2320211013330232-0301110121110113-1213030100003233"></a>

<a id="canonical-3112102232223010-0220030013222202-1111211301220020-3011232301311023-1033102230013323-0213323012011322-3322210100123232-1311223310010133"></a>

#### `dynamic_proxy.http_proxy.more_option.buffer_policy.max_request_bytes` property

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(10485760),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

<a id="canonical-2200031212103131-3310131223321321-1013221221132310-2010222013231121-1313201321220130-0202321300102200-0211121130203102-2332023130201030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.compression_params` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- dynamic_proxy.http_proxy.more_option.compression_params

<a id="canonical-0201033213231220-3001033300230203-3311012231233221-1001310312300023-3201210022303231-1331122112113032-2310333221020101-2313320231022123"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

By default compression will be skipped when:

A request does NOT contain accept-encoding header. A request includes accept-encoding header, but it
does not contain “gzip” or “\*”. A request includes accept-encoding with “gzip” or “\*” with the
weight “q=0”. Note that the “gzip” will have a higher weight then “\*”. For example, if
accept-encoding is “gzip;q=0,\*;q=1”, the filter will not compress. But if the header is set to
“\*;q=0,gzip;q=1”, the filter will compress. A request whose accept-encoding header includes
“identity”. A response contains a content-encoding header. A response contains a cache-control
header whose value includes “no-transform”. A response contains a transfer-encoding header whose
value includes “gzip”. A response does not contain a content-type value that matches one of the
selected mime-types, which default to application/JavaScript, application/JSON,
application/xhtml+XML, image/svg+XML, text/CSS, text/HTML, text/plain, text/XML. Neither
content-length nor transfer-encoding headers are present in the response. Response size is smaller
than 30 bytes (only applicable when transfer-encoding is not chunked).

When compression is applied:

The content-length is removed from response headers. Response headers contain “transfer-encoding:
chunked” and do not contain “content-encoding” header. The “vary: accept-encoding” header is
inserted on every response.

GZIP Compression Level:

A value which is optimal balance between speed of compression and amount of compression is chosen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("content_length")}
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
compression_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303020121132001-2002333010210200-0223200230221131-0033033303222001-2231322303300133-2200032111133330-0223130302113132-2222021312212321"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.compression_params`

<a id="canonical-0122313201213300-1103023220101303-0230023313000303-2232130120122300-2230302030122132-1210022313113323-0011210130110323-3231003203212022"></a>

#### `dynamic_proxy.http_proxy.more_option.compression_params.content_length` property

Type: `"number"`. Optional.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Additional upstream details:

The default value is 30.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(30),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 30
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  }
}
```

<a id="canonical-3111221020223312-1031303020101233-2313233320203312-3311303100331003-3100000212011003-1101300321200312-2312223000023103-0212232031022232"></a>

<a id="canonical-0202223012201112-1322232312200111-3230221232321011-0331312201213111-2113220330220031-0010020300210002-1211320111003321-3003212323213122"></a>

#### `dynamic_proxy.http_proxy.more_option.compression_params.content_type` property

Type: `["list", "string"]`. Optional.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/JavaScript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Additional upstream details:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/JavaScript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3022032330103333-3222310003200332-1333130302122221-3330303013023103-1212232101120022-1232110212000313-2013003120233103-2310213311300032"></a>

<a id="canonical-1122210210100310-3230032121123230-3332300310223201-0133111023023331-0203322013110100-2120123232032111-1102002303033302-1030032303103213"></a>

#### `dynamic_proxy.http_proxy.more_option.compression_params.disable_on_etag_header` property

Type: `"bool"`. Optional.

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

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

<a id="canonical-1300312220111301-0001012121333103-3333032212123002-1221031330212002-3232312011233033-2323222232102232-3221213011333023-2303231131030001"></a>

<a id="canonical-2122133122011202-3020133002301222-2310210033130300-3230333303300233-1211203100300113-2021211103210020-1022320122323103-1233311222303021"></a>

#### `dynamic_proxy.http_proxy.more_option.compression_params.remove_accept_encoding_header` property

Type: `"bool"`. Optional.

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

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

<a id="canonical-3010323103131011-2202132221302001-3230101000320102-1213223032032320-1331201133331230-0321310100011302-0131321202201220-0211331012332312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- dynamic_proxy.http_proxy.more_option.disable_path_normalize

<a id="canonical-1101100113013000-3120132013032211-0323202301031013-3223022122103200-1230221210303211-1021321311132113-2132132123011200-1221003200011230"></a>

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
disable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213331021323322-3030210022220213-2001313310032320-2310212103302212-2123120201001231-2223111022333000-1233101102313112-0132110222331132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- dynamic_proxy.http_proxy.more_option.enable_path_normalize

<a id="canonical-0202131213310023-2122123023333221-2331200130200303-0101010202222232-0121301303132332-0313332103321311-2110330023311122-0112120101332101"></a>

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
enable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300132130201131-0303023011210331-2213112112010120-1033210011003131-0103131233101201-1213132101102322-3112320303333013-0033212231102002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection

<a id="canonical-1231131032321020-1202100030212321-0233122210001010-3330121031203033-1013111022310321-1301113010030220-3021023033330321-1013011320020103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

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
no_request_limit_per_connection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202201333112001-1113233232320310-1221001002302311-0312120031201222-3230230022213231-3133322233323010-3332110221230120-0230320130222220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_cookies_to_add` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add

<a id="canonical-2022222312021313-2013221003321222-2013222203101023-2113020310032003-2030110132113101-3323330230011010-0121023003330101-3221103230332011"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010213231030122-0323333201032130-1030213303023320-2132021311221111-1321023020021120-2102211111022122-3210231102100003-1321121212310223"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_cookies_to_add`

<a id="canonical-0210220333113030-2111230002101012-1211131102330133-1303011123120312-1331121033302021-2100122032212322-3302213131223113-2321301003213010"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2301203233021130-1222003031120030-3120332323101333-2002001112222011-3310300111321110-1112122302303203-3331020303221210-3301313020331111"></a>

<a id="canonical-0102223022321233-1331211110212012-1111030110331102-0011131311001111-2210230223100020-2101231331303111-0021121121000002-0132000131313322"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.overwrite` property

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Additional upstream details:

If true, the value is overwritten to existing values. Default value is do not overwrite.

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

- [secret_value](resources--proxy--reference--group-002.md#canonical-2221121133321303-1301031031112023-1111231110021110-2223211301300001-1113002112011222-1211111333212102-1212331102122331-2003210210030301): complete subsection reference.

<a id="canonical-3232110311003230-0012032100322101-2211203202321130-2111330000232312-2223120323113212-0331301122021120-2112213122030232-0313101330003120"></a>

<a id="canonical-1003232112002110-3103002331112132-3221331022311322-0331220332113220-3333032103100220-2211131321301013-3203233322023212-0100201101311333"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```
