---
page_title: "xcsh_dns_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy reference."
---

# xcsh_dns_proxy reference

<a id="canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- Property reference

<a id="canonical-0322310303331321-3030121210321122-0332232113211233-0113212030230233-3122003232212331-0203201222120202-0022100232022112-0001023102323211"></a>

### Direct properties for `xcsh_dns_proxy`

<a id="canonical-3120111330231113-3032100122320111-3222221110113200-0001000033322110-3332010121121032-1313033222213133-2001211313133132-3033003123100111"></a>

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

- [cache_profile](resources--dns_proxy--reference--group-001.md#canonical-3223201131001301-3012133233101211-1220300001232100-3132313032022201-0320021233121203-3212310110212220-1013310330311333-2031313023321333): complete subsection reference.

- [ddos_profile](resources--dns_proxy--reference--group-001.md#canonical-2110003101023121-2222231331102321-2213230320013033-1313102003322220-0010011312320132-2221021103130201-1303331100212311-3230313220231100): complete subsection reference.

<a id="canonical-2310302221000111-0300230212013100-0102023233012200-2312030322301102-1212322121110100-3331113112033231-2331112001030020-0333331001020132"></a>

<a id="canonical-1131310000122012-0213122110101312-3221321021210013-1010120001112100-1310002201033210-3332230320333203-2023211023032201-3323031322113331"></a>

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

<a id="canonical-1022013310021000-0100332002023033-3330222202321301-2332003200001120-1013221113211232-3223113222103301-3310323010002313-2002301032223213"></a>

<a id="canonical-1312102031010221-1102320120231023-3310231213021010-1333211012221020-3030002033301021-3031130223100212-1101220110313133-0322112131321022"></a>

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

<a id="canonical-1223320333013133-3102123300010231-3011303313122111-1021131103103202-2112001101320132-2332030120102333-1320120113330213-3013103203320332"></a>

<a id="canonical-2312200133233221-1130120322322310-0323332233332012-1030212222321020-3113302313220030-3013232003132020-1031011113331031-0201110231111210"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](resources--dns_proxy--reference--group-001.md#canonical-3220021102300002-2030313113130310-0020121013020112-3132302120121303-2302301233333100-2110200322313022-2310200003321211-3322021001010231): complete subsection reference.

<a id="canonical-0110032113212002-0202013123332331-0102020221230132-2331210020122121-1123202303101112-1322231233210212-1213003120100203-0313233211213321"></a>

<a id="canonical-2100301112002322-2100011320302232-1020033003123002-2301333101032212-0211010212233121-1303212220002101-0203012313223223-3223132130213022"></a>

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

- [lb_algorithm](resources--dns_proxy--reference--group-001.md#canonical-0123302030312300-1300310310103200-2010132201320120-3002100002202030-2300121210113102-3033032012230111-1301131132112112-3013012210331233): complete subsection reference.

<a id="canonical-0120333301321310-3133023321030331-0022233300132103-0320322102231333-0113322111212123-3030322002303112-2132233033213013-3310201203001012"></a>

<a id="canonical-1210323103332210-3101311020223133-1010200012222130-1323101303030220-1030322010030213-2112203221111323-1002030020032031-0212323030132223"></a>

#### `name` property

Type: `"string"`. Required.

Name of the DNS Proxy. Must be unique within the namespace.

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

<a id="canonical-3222300011002011-3012122102100302-0313033131112101-2221033113113020-1100233232002312-2001302333003302-3323001310201010-3022333030221301"></a>

<a id="canonical-0023203110121223-3103313213000333-3133311133210110-2133200130020202-3303322031233311-1010032331332303-1221322030123021-0003212322322211"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace for the DNS Proxy. The F5 XC API restricts this resource to the system namespace; it
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

- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111): complete subsection reference.

- [protocol_inspection](resources--dns_proxy--reference--group-001.md#canonical-3220331331332203-0133330223020233-2010201120310110-3031102331001233-2301333130020221-1130222023012123-1332133111310213-1132022312202301): complete subsection reference.

- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230): complete subsection reference.

- [timeouts](resources--dns_proxy--reference--group-002.md#canonical-3323333303000202-2222033001303102-1032113131211303-1311033002122012-1002203323202112-3112220230111220-3311101211320012-3313320002112211): complete subsection reference.

<a id="canonical-2020130323313303-2120202101111013-3122222001020223-1300313303302220-3032320322322311-2021222103200231-2103200133311231-0021100322031220"></a>

<a id="canonical-3333122132011011-0203320310201011-1013133132010210-2132013110313202-0303212203130131-3332321002311233-3201330011030113-0100032110013112"></a>

#### `transport_type` property

Type: `"string"`. Optional, Computed.

\[Enum: UDP|TCP|BothTCPAndUDP\] Transport Type - UDP: UDP - TCP: TCP - BothTCPAndUDP: Both TCP and
UDP. Possible values are \`UDP\`, \`TCP\`, \`BothTCPAndUDP\`. Defaults to \`UDP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["BothTCPAndUDP","TCP","UDP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("UDP",
    "TCP",
    "BothTCPAndUDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UDP",
  "enum": [
    "UDP",
    "TCP",
    "BothTCPAndUDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2311310113232301-2011103320333021-3020132331302110-3202120201011111-1013333003020020-3303022233032111-3200330120131302-1022120202112233"></a>

### All schema paths for `xcsh_dns_proxy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--dns_proxy--reference--group-001.md#canonical-3120111330231113-3032100122320111-3222221110113200-0001000033322110-3332010121121032-1313033222213133-2001211313133132-3033003123100111) |
| `cache_profile` | [cache_profile](resources--dns_proxy--reference--group-001.md#canonical-3220021031130210-3131301212201133-3033103130122030-0321132012103330-3023103233110211-3130201110321230-1010312011322201-0320111120211211) |
| `cache_profile.cache_size` | [cache_profile.cache_size](resources--dns_proxy--reference--group-001.md#canonical-0003333103300233-2322033323021131-0320331123300200-3213202111030211-2020210023233212-0111020312300232-3333332230130010-3313210120133032) |
| `cache_profile.disable_cache_profile` | [cache_profile.disable_cache_profile](resources--dns_proxy--reference--group-001.md#canonical-3003312112023322-1100332121130003-1111003132022010-3011023112201121-2311100333002031-2301210132012203-2232213213303100-1012002021312211) |
| `ddos_profile` | [ddos_profile](resources--dns_proxy--reference--group-001.md#canonical-0001011310222100-1020210210233023-1200232102213131-1000322223020223-1331123313210120-1002310301133320-0031011000032222-3220122200231030) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](resources--dns_proxy--reference--group-001.md#canonical-3010220101223011-1120320110023122-1131311322311010-0320020012320231-3211211033031011-1121203023101130-1213101222102233-2222302232231230) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](resources--dns_proxy--reference--group-001.md#canonical-3113303033123303-0212222032132212-3321201223310000-0330332000002111-2033131011201313-3213330213023222-1321010333211003-3230300132122330) |
| `description` | [description](resources--dns_proxy--reference--group-001.md#canonical-2310302221000111-0300230212013100-0102023233012200-2312030322301102-1212322121110100-3331113112033231-2331112001030020-0333331001020132) |
| `disable` | [disable](resources--dns_proxy--reference--group-001.md#canonical-1022013310021000-0100332002023033-3330222202321301-2332003200001120-1013221113211232-3223113222103301-3310323010002313-2002301032223213) |
| `id` | [ID](resources--dns_proxy--reference--group-001.md#canonical-1223320333013133-3102123300010231-3011303313122111-1021131103103202-2112001101320132-2332030120102333-1320120113330213-3013103203320332) |
| `irules` | [irules](resources--dns_proxy--reference--group-001.md#canonical-3331300121110021-0220312112133330-0132310121122230-1012012103030110-1333220103230102-0030001203132003-1012100022110213-3121103222310020) |
| `irules.name` | [irules.name](resources--dns_proxy--reference--group-001.md#canonical-2102022231223132-3303221133131022-3303013112323333-2121220313123132-1221211003112211-0120023303310012-1100023232000213-0102321220100322) |
| `irules.namespace` | [irules.namespace](resources--dns_proxy--reference--group-001.md#canonical-0131333322111313-1010302312323123-2133320222122232-1112313113132132-1302102330033313-3323021123010220-0212130000031113-1332313011112320) |
| `irules.tenant` | [irules.tenant](resources--dns_proxy--reference--group-001.md#canonical-2123103000130133-3202320223000301-3133121123330210-3200333322002020-2310133102023212-0320013200323132-2222132321322233-2310201232000120) |
| `labels` | [labels](resources--dns_proxy--reference--group-001.md#canonical-0110032113212002-0202013123332331-0102020221230132-2331210020122121-1123202303101112-1322231233210212-1213003120100203-0313233211213321) |
| `lb_algorithm` | [lb_algorithm](resources--dns_proxy--reference--group-001.md#canonical-0310101102223321-0010010211322010-3033212333223301-2203001033302331-2322031120212212-1221201131331322-0033301230012233-2222011032132001) |
| `lb_algorithm.round_robin` | [lb_algorithm.round_robin](resources--dns_proxy--reference--group-001.md#canonical-3300322330131112-1103122121322122-0322222113213230-0023032132011312-3110132121003200-0233100331011331-3302012302223211-0210130001133320) |
| `name` | [name](resources--dns_proxy--reference--group-001.md#canonical-0120333301321310-3133023321030331-0022233300132103-0320322102231333-0113322111212123-3030322002303112-2132233033213013-3310201203001012) |
| `namespace` | [namespace](resources--dns_proxy--reference--group-001.md#canonical-3222300011002011-3012122102100302-0313033131112101-2221033113113020-1100233232002312-2001302333003302-3323001310201010-3022333030221301) |
| `origin_servers` | [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-3221133002023211-1230131123032230-2022230311222031-0303102020012112-0020303022221020-1002203322130011-3121320111332202-0101333100002012) |
| `origin_servers.health_checks` | [origin_servers.health_checks](resources--dns_proxy--reference--group-001.md#canonical-0100203221200313-3321002013303203-3120022232310311-0210013033101312-1322220202230201-2102233110232232-3002130201223112-1300320222102010) |
| `origin_servers.health_checks.health_check` | [origin_servers.health_checks.health_check](resources--dns_proxy--reference--group-001.md#canonical-1212113033010213-2111220032212302-1023222113313110-1021213003012031-0011300313010013-3310103313202313-1312300332331102-1030010320123203) |
| `origin_servers.health_checks.health_check.dns_health_check` | [origin_servers.health_checks.health_check.dns_health_check](resources--dns_proxy--reference--group-001.md#canonical-0101023132323010-0313313211022112-1132200012310002-3110320100331010-2233202131211201-3203012302100221-1331201330122021-0212033333201301) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_rcode` | [origin_servers.health_checks.health_check.dns_health_check.expected_rcode](resources--dns_proxy--reference--group-001.md#canonical-2101110102013222-3330103310001131-0132131123113203-1033311230101100-3122310023203032-3313113122022330-3232310120203023-3022113122001021) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_record_type` | [origin_servers.health_checks.health_check.dns_health_check.expected_record_type](resources--dns_proxy--reference--group-001.md#canonical-1033123223131031-1132332130033100-2230302132222201-1202122230022011-2121113032001222-0133133000332330-0100013020031332-1030330213000100) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_response` | [origin_servers.health_checks.health_check.dns_health_check.expected_response](resources--dns_proxy--reference--group-001.md#canonical-2032020101230121-1023120121233031-2011330113201123-3331330020113102-3023001210212012-2320103020011120-0302112332220100-2200011333321213) |
| `origin_servers.health_checks.health_check.dns_health_check.query_name` | [origin_servers.health_checks.health_check.dns_health_check.query_name](resources--dns_proxy--reference--group-001.md#canonical-1222033020301132-1020123331313332-3113001230030332-3012203222213332-3230203302221200-1331100330311330-1113022110333232-1321313223301301) |
| `origin_servers.health_checks.health_check.dns_health_check.query_type` | [origin_servers.health_checks.health_check.dns_health_check.query_type](resources--dns_proxy--reference--group-001.md#canonical-2323331220021021-2131030301223330-0212322311012323-0132011311322033-0202212130020223-2202110102131322-3200102232021222-3012133331211302) |
| `origin_servers.health_checks.health_check.dns_health_check.reverse` | [origin_servers.health_checks.health_check.dns_health_check.reverse](resources--dns_proxy--reference--group-001.md#canonical-0120001113020213-2110333000212103-2031221212333013-2321210131030133-0031303021122232-3020100021111321-0223022132320032-0223232200322102) |
| `origin_servers.health_checks.health_check.icmp_health_check` | [origin_servers.health_checks.health_check.icmp_health_check](resources--dns_proxy--reference--group-001.md#canonical-3321213201311130-3223303133331012-0212232310331121-3023100112231323-2300312000021121-1312111230323012-0110320031233131-3002031311032122) |
| `origin_servers.health_checks.health_check.tcp_health_check` | [origin_servers.health_checks.health_check.tcp_health_check](resources--dns_proxy--reference--group-001.md#canonical-2110130130311022-2221333123220030-2033213102002030-2112000212331113-0100200132113303-1233110101001211-2210111312300313-0133020313123332) |
| `origin_servers.health_checks.health_check.tcp_health_check.expected_response` | [origin_servers.health_checks.health_check.tcp_health_check.expected_response](resources--dns_proxy--reference--group-001.md#canonical-2030222011030013-0033032202101012-1310202003333113-0301012131021131-3203202313000010-0102331223320332-0303130113011300-0020023201303032) |
| `origin_servers.health_checks.health_check.tcp_health_check.send_payload` | [origin_servers.health_checks.health_check.tcp_health_check.send_payload](resources--dns_proxy--reference--group-001.md#canonical-0103021110300030-3213100331232313-0003221323123221-2022330200202012-3312310012203332-0020130023232033-2023323311203200-2132203301213031) |
| `origin_servers.health_checks.healthy_threshold` | [origin_servers.health_checks.healthy_threshold](resources--dns_proxy--reference--group-001.md#canonical-2010232332221211-1023233231320311-2132333200301303-2122313003221321-3112312010101122-1222122032113123-0212011101320313-3213122301323122) |
| `origin_servers.health_checks.interval` | [origin_servers.health_checks.interval](resources--dns_proxy--reference--group-001.md#canonical-3132020222032100-0102112313230123-0112332330133031-3203000030003202-1210313222021311-1311312002331031-2110202101110133-2031211030023113) |
| `origin_servers.health_checks.timeout` | [origin_servers.health_checks.timeout](resources--dns_proxy--reference--group-001.md#canonical-0223113311100220-1333030331321213-0003122112031221-0012303321231020-1023300113023222-0030303210201123-0200132033212302-3332321323200230) |
| `origin_servers.health_checks.unhealthy_threshold` | [origin_servers.health_checks.unhealthy_threshold](resources--dns_proxy--reference--group-001.md#canonical-2102133111203303-0112103202022321-0132100333133001-0021011320223003-2122100230333123-1220322203331310-3110220231101331-3302103103230332) |
| `origin_servers.origin_servers` | [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-3332031133332102-1210200123202301-3131332220113013-3002120010112310-1320131201213202-1232112031021311-1113123301211013-1323021123222323) |
| `origin_servers.origin_servers.k8s_service` | [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-2311031010122023-0020203300022113-3113231100221022-0202232100323120-2130220120312121-3313203322211200-3000003133031003-1211200322230000) |
| `origin_servers.origin_servers.k8s_service.inside_network` | [origin_servers.origin_servers.k8s_service.inside_network](resources--dns_proxy--reference--group-001.md#canonical-0331130101111121-3111230320133002-3132033200202303-0221003323232102-3102113302020232-0013002202303201-1113221321223021-1031322310210202) |
| `origin_servers.origin_servers.k8s_service.outside_network` | [origin_servers.origin_servers.k8s_service.outside_network](resources--dns_proxy--reference--group-001.md#canonical-0313112012333031-2101230211010210-0201013201033333-3220100130303121-1213000312022200-0321123330332330-1330202313312020-3321133312321333) |
| `origin_servers.origin_servers.k8s_service.protocol` | [origin_servers.origin_servers.k8s_service.protocol](resources--dns_proxy--reference--group-001.md#canonical-0203320220301132-2211210333300231-2100202230012110-1112322331203023-1110131000000001-0202121301111023-1021211200033230-0212001132100310) |
| `origin_servers.origin_servers.k8s_service.service_name` | [origin_servers.origin_servers.k8s_service.service_name](resources--dns_proxy--reference--group-001.md#canonical-1333213032322021-1121312333300221-1031122323312122-1130221231313210-0301013201101331-0323201231301230-0300003133132100-3230230221101231) |
| `origin_servers.origin_servers.k8s_service.site_locator` | [origin_servers.origin_servers.k8s_service.site_locator](resources--dns_proxy--reference--group-001.md#canonical-0131123131302023-2003033111112021-1312322222231113-3330321122312331-3310100200203002-3011132123222003-2121002123312232-0220020111321002) |
| `origin_servers.origin_servers.k8s_service.site_locator.site` | [origin_servers.origin_servers.k8s_service.site_locator.site](resources--dns_proxy--reference--group-001.md#canonical-3012023321122213-0203013310002233-0330230002111231-3011132231012010-1021100111211303-3210101030111100-1310122321033220-0030032222113320) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.name` | [origin_servers.origin_servers.k8s_service.site_locator.site.name](resources--dns_proxy--reference--group-001.md#canonical-3133331112201330-2221323000032231-3322312201110303-2332322333113122-2121213121313232-3331332120232230-0032203311130133-2013002110212222) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.namespace` | [origin_servers.origin_servers.k8s_service.site_locator.site.namespace](resources--dns_proxy--reference--group-001.md#canonical-2212111320331020-3100123221103200-0300020333030303-2313313033213221-1221132300011003-0221012331333003-2332101322233133-0301333213330223) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.tenant` | [origin_servers.origin_servers.k8s_service.site_locator.site.tenant](resources--dns_proxy--reference--group-001.md#canonical-1033021203033110-0110322113233301-3022023230012232-0002002302322333-1211120213002300-3010123302122301-2230122203031123-3100200203101130) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site](resources--dns_proxy--reference--group-001.md#canonical-0322203112313002-0130113313212101-2020011003331333-0302120303312211-1202000023223110-3321020311223231-1213230121011022-3220001101312203) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name](resources--dns_proxy--reference--group-001.md#canonical-2311031212223221-1213202210320320-2110101201100213-1110012202320032-3021102113301323-2221332210021300-1131011031303233-1120122210001201) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace](resources--dns_proxy--reference--group-001.md#canonical-0000000232323103-3313310332113203-2103132221133311-0030322033130010-3131021122103120-3102232323020113-0212322222000103-0133333102000313) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant](resources--dns_proxy--reference--group-001.md#canonical-3301300223110303-3101233313202020-0030333312113000-0310320332302122-0113221331013223-0213201020302030-2032033032133302-0121030311002011) |
| `origin_servers.origin_servers.k8s_service.snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool](resources--dns_proxy--reference--group-001.md#canonical-3012010102013131-1120232112231130-1001013011013012-2112122131332232-0030102221203312-0321111102332132-1221120033222300-1130030100033033) |
| `origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](resources--dns_proxy--reference--group-001.md#canonical-1010031322233330-3301033023121320-3212322021301330-3203220133322013-2213202322210223-1312012112333010-1031323112031201-0131210321003222) |
| `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](resources--dns_proxy--reference--group-001.md#canonical-0110213033211201-3220233013112013-2330011021021030-2032123232222232-3000023210212300-1332202023210133-1332320111103130-0203121132232203) |
| `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes](resources--dns_proxy--reference--group-001.md#canonical-3232202221133012-0330232022321213-0232110102203000-2302331030032030-2320222222221013-2333303123322131-1131331311120312-0222220000132031) |
| `origin_servers.origin_servers.k8s_service.vk8s_networks` | [origin_servers.origin_servers.k8s_service.vk8s_networks](resources--dns_proxy--reference--group-001.md#canonical-2311301330310123-3032333113020100-1003203223223133-0021332302013232-3132323232311112-2023030031332020-0332032231103200-1330123130221000) |
| `origin_servers.origin_servers.no_preference` | [origin_servers.origin_servers.no_preference](resources--dns_proxy--reference--group-001.md#canonical-0302230133011112-0313003031302002-3211113031213231-1222203330120100-0113030131230113-0030021030120101-1113321020022101-3300332200023233) |
| `origin_servers.origin_servers.public_ip` | [origin_servers.origin_servers.public_ip](resources--dns_proxy--reference--group-001.md#canonical-1033301121023032-2221003220202103-2003300000301031-3023313211313202-0003213001201023-2300311112222023-1231012312312300-0131201001201103) |
| `origin_servers.origin_servers.public_ip.ip` | [origin_servers.origin_servers.public_ip.ip](resources--dns_proxy--reference--group-001.md#canonical-1310200022000010-2011333000322233-3100020313121221-2023031222112332-0303001220020111-2331232321333002-1003012213130033-1212113003221330) |
| `origin_servers.origin_servers.public_name` | [origin_servers.origin_servers.public_name](resources--dns_proxy--reference--group-001.md#canonical-2310230113110121-3233002112131232-0310130223033222-0200223001221330-2320002233020210-0102000312212213-1301200221100020-2320132013030200) |
| `origin_servers.origin_servers.public_name.dns_name` | [origin_servers.origin_servers.public_name.dns_name](resources--dns_proxy--reference--group-001.md#canonical-1112213313131102-3202313100210021-0200013023330011-1322121013300022-1232131222312011-2222013310311100-0032111022022112-2020023022331302) |
| `origin_servers.origin_servers.public_name.refresh_interval` | [origin_servers.origin_servers.public_name.refresh_interval](resources--dns_proxy--reference--group-001.md#canonical-1031332023310330-3311302011121020-3010103312213101-0210202200220101-1213331002032112-2032103213101022-1233231302212003-1300222130101002) |
| `origin_servers.origin_servers.site_preferences` | [origin_servers.origin_servers.site_preferences](resources--dns_proxy--reference--group-001.md#canonical-2310210330121011-3222112222131231-1003321110121111-0012012001200313-3312031013310331-2032130120003330-0112320121310203-3130221331331011) |
| `origin_servers.origin_servers.site_preferences.refs` | [origin_servers.origin_servers.site_preferences.refs](resources--dns_proxy--reference--group-001.md#canonical-2102131103302310-3002111230231211-1200120301311023-0311322033302030-1212023203303310-3210320330123013-3211201022103012-2113231232121012) |
| `origin_servers.origin_servers.site_preferences.refs.name` | [origin_servers.origin_servers.site_preferences.refs.name](resources--dns_proxy--reference--group-001.md#canonical-0320110110100310-0223021231130130-1322333100223211-3121033001112020-1310020130202210-0030231102103133-0102131332130303-2111300210031232) |
| `origin_servers.origin_servers.site_preferences.refs.namespace` | [origin_servers.origin_servers.site_preferences.refs.namespace](resources--dns_proxy--reference--group-001.md#canonical-1311201110031220-2301231233011211-2332101103032313-1130313000130213-3230132210130002-1130113131323332-0032100031310233-0223000111303010) |
| `origin_servers.origin_servers.site_preferences.refs.tenant` | [origin_servers.origin_servers.site_preferences.refs.tenant](resources--dns_proxy--reference--group-001.md#canonical-3203032112211002-2310011112032010-0231201232123133-3230111310132002-2102030300022132-2101300123310121-0023120020112012-1013020313202110) |
| `protocol_inspection` | [protocol_inspection](resources--dns_proxy--reference--group-001.md#canonical-2310103002211302-2333002012333323-1211022012123320-2332232132213310-2021213021300202-2331332311032210-1202133230230131-1313312311321230) |
| `protocol_inspection.name` | [protocol_inspection.name](resources--dns_proxy--reference--group-001.md#canonical-1000221230311033-3001102011222123-0300232313300212-0301223221111212-2232302010233231-3031312102121331-3300101001310300-3112133231322211) |
| `protocol_inspection.namespace` | [protocol_inspection.namespace](resources--dns_proxy--reference--group-001.md#canonical-2321310020010123-0321003311122312-0322321030030001-3132321313031233-3222023001321201-1201110223203202-1310330201033332-3233210112123033) |
| `protocol_inspection.tenant` | [protocol_inspection.tenant](resources--dns_proxy--reference--group-001.md#canonical-0332230333311013-3301121001011313-3101023110310031-1112102023130101-3213031333031321-0200111313333233-3112233301200212-1131333132302011) |
| `proxy_advertisement` | [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-0021201222031123-2101223120210311-2332222111003122-3220102022320323-2122023022130101-3103100010231030-3231211331312003-0311310010121100) |
| `proxy_advertisement.advertise_custom` | [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3213132112221312-3300012330211231-0133123031022133-1130231001332223-2132023132230013-3322130122320323-2223311312322333-2231310233300000) |
| `proxy_advertisement.advertise_custom.advertise_where` | [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-3003223223300310-3011211031123323-1201331300020311-1200221300002310-1212102011313022-0133233033320023-2032300332031222-1201323323221333) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--dns_proxy--reference--group-001.md#canonical-2213121001221203-2023133202223113-3100312111202303-3220233321331230-2213113221011002-2310132211312121-2201212333230202-0130311213022110) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--dns_proxy--reference--group-001.md#canonical-0101002010331332-0033103031331232-1312121201100222-0020210021311301-1233320003333001-3301011112310231-3203012033213210-0013123330300321) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](resources--dns_proxy--reference--group-001.md#canonical-0101002110311020-1001021310031220-2320202330103030-2101100111223322-2302031312221303-3011230113120211-2211130020003201-1101230030001103) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](resources--dns_proxy--reference--group-001.md#canonical-1030323332322320-2021200012012203-3332111303333123-0210301101133002-2001231002300022-1013311323022303-2303210021010023-1202323333302122) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](resources--dns_proxy--reference--group-001.md#canonical-2302112001123302-0121212112132102-1031121001313002-0002321223032133-3031121103001002-3330300130223121-0230300003212123-0120332130322332) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--dns_proxy--reference--group-001.md#canonical-2301300312103302-0312310200132131-1313010010301011-1323023113302020-1110311312303132-2221122200022310-2032020303211003-3122233120202031) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](resources--dns_proxy--reference--group-001.md#canonical-1111130223301320-1101111331101331-1300133030201233-2212331321111303-0133100220231212-0023302213121000-0302213313022322-3211021320203112) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name](resources--dns_proxy--reference--group-001.md#canonical-3231101301213200-2231022302002202-0230230233232333-3202200331023123-2030111222313233-3120210002203011-3001313330003001-3313003202001113) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](resources--dns_proxy--reference--group-001.md#canonical-2202202230001320-1332000303213111-0132213132310300-3231222033100212-3211220121220011-2003013311113103-1221022330300122-0313233220011333) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](resources--dns_proxy--reference--group-001.md#canonical-3323202010223220-1302223012100233-0120231330303101-3222232213210133-3102023123031323-3021332112001022-1210332332311232-1013113131111231) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--dns_proxy--reference--group-001.md#canonical-2212301320331322-0123000231200203-1111121232133032-1032121112310320-0330020003110130-3112110310131020-1200311003003121-0231322301330021) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--dns_proxy--reference--group-001.md#canonical-2221201302323300-3333022111320103-2211320000312123-1102131012110121-0033020020232232-2213320101321122-2101111211211011-2323310003012000) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](resources--dns_proxy--reference--group-001.md#canonical-3201010212011110-1101233311202011-0202000000123211-2133332101113121-0121222122221120-1333201021002302-0133122020322231-3211001033233100) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](resources--dns_proxy--reference--group-001.md#canonical-3102122002113122-3300300131120122-1113122333230200-2302010010011131-1131221322111010-2310330201003212-3332202133022003-1321301132100023) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](resources--dns_proxy--reference--group-001.md#canonical-2020300013132010-0233030301132132-3121212223130211-1000323302303121-3002233022302032-2311201203111032-2030010310310003-1022031331113003) |
| `proxy_advertisement.advertise_custom.advertise_where.port` | [proxy_advertisement.advertise_custom.advertise_where.port](resources--dns_proxy--reference--group-001.md#canonical-1130132230110332-0321221103110123-1232002310000133-0231111301020130-3313133233232222-0200023030020230-3220013132333001-3132312310322112) |
| `proxy_advertisement.advertise_custom.advertise_where.port_ranges` | [proxy_advertisement.advertise_custom.advertise_where.port_ranges](resources--dns_proxy--reference--group-001.md#canonical-3033003112303210-0110122122110233-1123031230332003-0211010200133120-0221230301003030-0311033232101012-1101200131002012-1311022033000223) |
| `proxy_advertisement.advertise_custom.advertise_where.site` | [proxy_advertisement.advertise_custom.advertise_where.site](resources--dns_proxy--reference--group-001.md#canonical-2313301120110301-3332023233032012-1201123231210211-2012022203130103-3213311223123101-0212222202132033-2000330003321103-2221302003022032) |
| `proxy_advertisement.advertise_custom.advertise_where.site.ip` | [proxy_advertisement.advertise_custom.advertise_where.site.ip](resources--dns_proxy--reference--group-001.md#canonical-3201111200013122-0102203133302121-3032230313013310-1123302131020131-0030211101103103-0132213022132003-3220202302223332-3010122021002122) |
| `proxy_advertisement.advertise_custom.advertise_where.site.network` | [proxy_advertisement.advertise_custom.advertise_where.site.network](resources--dns_proxy--reference--group-001.md#canonical-1101031132033003-1110311021030231-1320231223211013-2002220113303210-3211201230003320-1301320032012110-3001310132330301-0033120312332011) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site` | [proxy_advertisement.advertise_custom.advertise_where.site.site](resources--dns_proxy--reference--group-001.md#canonical-3121012021321022-1030012101200203-3130002231220300-1110031221000032-2203321213002333-1011121123133032-0232203123300102-2233132331232310) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.name` | [proxy_advertisement.advertise_custom.advertise_where.site.site.name](resources--dns_proxy--reference--group-001.md#canonical-1313223333010210-2012223303213231-3023002210222033-3201203101313220-3232223323323312-1211020201033010-1103102132333020-0313212220230210) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.site.site.namespace](resources--dns_proxy--reference--group-001.md#canonical-1231323310120323-1232220000320223-1111022232022230-1321321310220332-2310202310032330-3322310010001023-0001323201023231-2312110103212232) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.site.site.tenant](resources--dns_proxy--reference--group-001.md#canonical-0100000320101122-0000210011223003-2133220103122002-1202320200133322-1011221010112223-2303312121101122-0102211031033322-3030130323231133) |
| `proxy_advertisement.advertise_custom.advertise_where.use_default_port` | [proxy_advertisement.advertise_custom.advertise_where.use_default_port](resources--dns_proxy--reference--group-001.md#canonical-1133113321203103-3101030303002133-1022220023012200-1012022331030232-0211212130232033-2332001201010201-3320012112230023-2210110123321301) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-001.md#canonical-3022012000222210-1003230222202323-3003020331002222-2220021330120300-0113301010130111-0100211103333212-0303320203012333-1121320112020023) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--dns_proxy--reference--group-002.md#canonical-2203220020003122-1223331223330131-1213210021002211-1331300220113020-2320333000310102-3131200201120111-2213013021320012-2310133311103010) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](resources--dns_proxy--reference--group-002.md#canonical-3002013123313002-2002221302033222-2120010212131001-1113202112223310-1223122033121211-2003231332122221-2010302100230332-2020102010213031) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip](resources--dns_proxy--reference--group-002.md#canonical-3110212102132200-3120232310321203-2322303302230101-3130230123100321-1000222030302202-3201130033020301-2003010211213011-0003112111133110) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip](resources--dns_proxy--reference--group-002.md#canonical-1230000013311031-2131303130201300-1003313310130233-2121111000210002-0102300220101132-2320122211322330-1231310112321213-2001310200000012) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-0222213132202020-1121330003202230-2031323101023201-0012200222211231-2201122200130000-2031211101010131-2121002312020001-2010333323321223) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name](resources--dns_proxy--reference--group-002.md#canonical-3000030120100023-2221123130313031-2233311332003232-0002031013010222-3312213033231011-3011100011222330-2101100112231230-1230121003332212) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace](resources--dns_proxy--reference--group-002.md#canonical-1213201320010032-2323020212020103-0021021333011123-0011330010320222-1331222122213122-1110132200221330-1303303033121313-2201310210311213) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant](resources--dns_proxy--reference--group-002.md#canonical-2202003331322310-2223003132110232-3000010131230310-1332011033030300-1223232303013021-1030010201221032-3223001101011003-2131113113313200) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-2000331302020031-0312102230303301-0303023121001132-0321022210120003-3313010333322232-0213101331032233-1332020130010321-2122023120111232) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.network](resources--dns_proxy--reference--group-002.md#canonical-1113203110221003-1131202132321021-0233311312113320-3101003021013022-2002201311231113-3323232113113300-3203323133303200-2003333221311120) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-0021100002030003-2312230001232001-1122222010311010-0330101233331201-2212002313313212-1022301200032232-0321210320221323-0200300030302020) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name](resources--dns_proxy--reference--group-002.md#canonical-3020311221030122-1202003213021230-3131032021021123-1023202122213202-0112122213201330-1200010010221332-0211121231213322-3111110322011133) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace](resources--dns_proxy--reference--group-002.md#canonical-0013220031130313-3013001031010213-2221032300333201-2012030233223110-2301003302233122-3333203233303012-1322210120302322-1132001132021111) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant](resources--dns_proxy--reference--group-002.md#canonical-0103203210211310-2021030133010203-1121110111130102-0333113032013112-0232113332232031-0303302210312110-0312013303012330-0302203133303123) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--dns_proxy--reference--group-002.md#canonical-2322000121112300-0011031322121120-3021203331300001-3022021012123321-3132222132112302-0123301231010300-3031032201200212-0333213323230201) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip](resources--dns_proxy--reference--group-002.md#canonical-2132210203323022-2002010100023111-3312332103331320-1131203323222103-2312132320020230-0111211230011121-2230222232010302-1001133313323231) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network](resources--dns_proxy--reference--group-002.md#canonical-2200023110002111-3233330110110202-0310101130032002-1223121223231232-3012003001121201-0231330012333000-1101000133210332-3332221232103223) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-3133103110122322-3103021010203231-1100232131203130-0030320020132023-1101302110231321-3311101130112220-0113111310313232-1332312100121212) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](resources--dns_proxy--reference--group-002.md#canonical-2320122331023032-1102021312221201-1022013013303200-3030120130123203-0330203120331201-1102103331203103-2003001032031310-1320202001333232) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](resources--dns_proxy--reference--group-002.md#canonical-1011122203221030-1123113212013023-3322120131223133-2123123031031331-0122203112213131-3301123121211030-2223133212310030-2120202120310203) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](resources--dns_proxy--reference--group-002.md#canonical-1021133303223230-2201001230213313-1112001120332232-1321330111301111-3202320310003300-3023223333232321-0310200333311020-0012303210323222) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-0130112030213102-0210233031303222-3021132321111132-0112031102232131-3010302133123322-0202211020200131-2330222113123311-0032222201313303) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](resources--dns_proxy--reference--group-002.md#canonical-2001123130003223-3102010022300233-1000113102130232-3002030030222122-2333113202212310-2233230001012130-1131333012200223-3011013031031102) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name](resources--dns_proxy--reference--group-002.md#canonical-0001221132022202-1222203001212323-3211203320130021-2131031222033322-1212101201101313-0012112101013012-1032032122330303-0302313133212012) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace](resources--dns_proxy--reference--group-002.md#canonical-1030111032030311-1010221012023113-3010030331100301-3110133120232013-3332000133321222-3213321122302233-1003311131301231-0320010323102123) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant](resources--dns_proxy--reference--group-002.md#canonical-1330200132333031-0302212303001322-1303022012212031-0313203120323233-3210211020321100-2300031121332202-2113331302021211-2301122231333021) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-0220022230213123-0211200020100231-3333333112201112-3000110013002000-2203230312011032-0311013030310231-2212331121000202-1330233300102112) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name](resources--dns_proxy--reference--group-002.md#canonical-0020112322220221-3202032303112203-0202110203223033-2032020031200100-0120202232310132-1122110130303202-2031212203011100-2300112220022021) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](resources--dns_proxy--reference--group-002.md#canonical-0232011111001203-2020010031301132-1212112131222022-2313313221321230-0331111331103203-3003203131313330-2033212132120232-3021223233113030) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](resources--dns_proxy--reference--group-002.md#canonical-0003331312302013-3110210300130033-2303030030031110-1210200012231303-3013110103323223-1003020133322202-1123101122331232-3213101110333123) |
| `proxy_advertisement.advertise_dualstack_on_public` | [proxy_advertisement.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-3020020013121233-3031311033231313-0030211030221130-1123003331111033-0111302233303300-1221131132300211-2232120103231112-3122030121232213) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_dualstack_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-2203212130210320-2323202330103120-0202202120201120-3320313010330321-1301111200321033-1010121211013230-3210212103332222-2002311303202022) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.name](resources--dns_proxy--reference--group-002.md#canonical-2221313111301231-0111330101312110-3211320220030333-2000131202113000-1031330103121303-3321320231010021-0230310222220113-1220131023312000) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace](resources--dns_proxy--reference--group-002.md#canonical-2300101312112133-3031003200211203-2100223222111003-2133032111210030-3120201122220132-1230102223320023-1113203231300323-1302301113211133) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant](resources--dns_proxy--reference--group-002.md#canonical-3022011221103222-3320122312311232-2222231213323312-1023320232231113-0212213020022200-3300200113133311-1120320021023333-0000030333122013) |
| `proxy_advertisement.advertise_on_public` | [proxy_advertisement.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-1110322033101122-3003022311212312-0022132232200202-1221003100230212-2223131233220202-0031012212233122-3122230130002102-0322200120311221) |
| `proxy_advertisement.advertise_on_public.public_ip` | [proxy_advertisement.advertise_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-3001021030310113-2101133213311130-3331323200230030-1121232033133013-3220300001300301-3321313313311032-3132220030033003-1333201122131330) |
| `proxy_advertisement.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_on_public.public_ip.name](resources--dns_proxy--reference--group-002.md#canonical-2313202322003112-1102301011313203-2311320032022130-2221103210232303-3023232120302031-3223220122231203-1123113301121000-2230020203320313) |
| `proxy_advertisement.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_on_public.public_ip.namespace](resources--dns_proxy--reference--group-002.md#canonical-3100121110222121-0102320103011120-0003323132001220-0321302022202230-1320111201310311-2023332123332223-2330313023331301-3211001112033302) |
| `proxy_advertisement.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_on_public.public_ip.tenant](resources--dns_proxy--reference--group-002.md#canonical-1102230123021113-3230233331221020-1313003133033123-0033330311122023-0230211001222132-3032131130013123-3133023020123313-3111320003231111) |
| `proxy_advertisement.advertise_on_public_default_dualstack_vip` | [proxy_advertisement.advertise_on_public_default_dualstack_vip](resources--dns_proxy--reference--group-002.md#canonical-2010310231010331-1113120123100230-1130133211101302-1323300011022312-1313011332032302-2231232320233211-0231121223022222-1212133303212022) |
| `proxy_advertisement.advertise_on_public_default_ipv6_vip` | [proxy_advertisement.advertise_on_public_default_ipv6_vip](resources--dns_proxy--reference--group-002.md#canonical-1332102031200011-1133321222023301-3310333101100033-2011031112313130-1001012102311000-0022332003011103-1313122030303011-3210312303100033) |
| `proxy_advertisement.advertise_on_public_default_vip` | [proxy_advertisement.advertise_on_public_default_vip](resources--dns_proxy--reference--group-002.md#canonical-2302022130021210-0120233122130103-0221120000331100-0112232110212303-2120231312000302-1323232301203220-2030203122001030-2130001213223123) |
| `proxy_advertisement.advertise_v6_on_public` | [proxy_advertisement.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-1102121313013002-0112123033211102-0303312200120302-0221001121202211-3103001303023120-1123311230231321-2011302033231102-0320313103333302) |
| `proxy_advertisement.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_v6_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-3300021010030330-3100133331003023-0121323203322311-1312201312012223-1301023321100203-0330230103120023-1103011002032102-2001032110001221) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_v6_on_public.public_ip.name](resources--dns_proxy--reference--group-002.md#canonical-0133333023031101-0000101021012011-0111120102110222-0202003232121211-0200200300303112-3332020210223220-1113031231320200-1111331330003330) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_v6_on_public.public_ip.namespace](resources--dns_proxy--reference--group-002.md#canonical-3203131020011022-1121133111110032-3000123031131033-0013111301132233-1001022103323120-2112020112030122-1303311302322230-0333122021200121) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_v6_on_public.public_ip.tenant](resources--dns_proxy--reference--group-002.md#canonical-3100002311122231-3102221233003330-1110030121121230-3122111111200211-1102313031220033-0032102122101122-3001100320020231-1123030231103222) |
| `proxy_advertisement.do_not_advertise` | [proxy_advertisement.do_not_advertise](resources--dns_proxy--reference--group-002.md#canonical-3011330200112223-3321223312121211-1330323001012213-3130112113101221-0301013010310310-2213000202032202-0211201223130122-2310021302333213) |
| `timeouts` | [timeouts](resources--dns_proxy--reference--group-002.md#canonical-2121133002213012-3323330102321310-3111120012100112-0102120022101130-3121022013002130-1103013032222110-1232313023220331-0300331230201033) |
| `timeouts.create` | [timeouts.create](resources--dns_proxy--reference--group-002.md#canonical-0200202011322323-3301012301121211-0332101003323232-3020211300301023-2002321300210213-0313220111310002-2301323013213202-3110021122030001) |
| `timeouts.delete` | [timeouts.delete](resources--dns_proxy--reference--group-002.md#canonical-1022130011203113-2100212123031332-2322101203002233-3211030031322230-2232023030102023-3303322301132131-0230201111033022-3010310001203220) |
| `timeouts.read` | [timeouts.read](resources--dns_proxy--reference--group-002.md#canonical-2011132030031233-3212321333011301-0031131131013303-1022211311020122-1312022132030312-1020120131003010-2203321131321232-1020100130020332) |
| `timeouts.update` | [timeouts.update](resources--dns_proxy--reference--group-002.md#canonical-3231113033232332-3300021010010210-2011120231102311-1202222300233212-2123111000230301-3220131110112002-2301021133203033-0222020322033210) |
| `transport_type` | [transport_type](resources--dns_proxy--reference--group-001.md#canonical-2020130323313303-2120202101111013-3122222001020223-1300313303302220-3032320322322311-2021222103200231-2103200133311231-0021100322031220) |

<a id="canonical-3223201131001301-3012133233101211-1220300001232100-3132313032022201-0320021233121203-3212310110212220-1013310330311333-2031313023321333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_profile` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- cache_profile

<a id="canonical-3220021031130210-3131301212201133-3033103130122030-0321132012103330-3023103233110211-3130201110321230-1010312011322201-0320111120211211"></a>

Type: `"object"`. single nested block, Optional.

DNS Cache specifies cache configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("cache_size",
    "disable_cache_profile")}
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
  "x-ves-oneof-field-cache_profile_choice": "[\"cache_size\",\"disable_cache_profile\"]"
}
```

Terraform syntax:

```terraform
cache_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200111131033000-0212213102202002-3213221202330203-3103220232212321-2303020233210322-3332211113223222-3030320200101131-0012101301320202"></a>

### Direct properties for `cache_profile`

<a id="canonical-0003333103300233-2322033323021131-0320331123300200-3213202111030211-2020210023233212-0111020312300232-3333332230130010-3313210120133032"></a>

#### `cache_profile.cache_size` property

Type: `"number"`. Optional.

Exclusive with \[disable\_cache\_profile\] cache size.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 10240),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10240,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10240"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10240"
  }
}
```

- [disable_cache_profile](resources--dns_proxy--reference--group-001.md#canonical-0212030122111310-0001322001032333-3111220300130012-1112202203333311-3001120122013032-2311032113110123-1110302030132313-0301000310012230): complete subsection reference.

<a id="canonical-0212030122111310-0001322001032333-3111220300130012-1112202203333311-3001120122013032-2311032113110123-1110302030132313-0301000310012230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_profile.disable_cache_profile` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [cache_profile](resources--dns_proxy--reference--group-001.md#canonical-3223201131001301-3012133233101211-1220300001232100-3132313032022201-0320021233121203-3212310110212220-1013310330311333-2031313023321333)
- cache_profile.disable_cache_profile

<a id="canonical-3003312112023322-1100332121130003-1111003132022010-3011023112201121-2311100333002031-2301210132012203-2232213213303100-1012002021312211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable cache profile.

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
disable_cache_profile = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110003101023121-2222231331102321-2213230320013033-1313102003322220-0010011312320132-2221021103130201-1303331100212311-3230313220231100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_profile` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- ddos_profile

<a id="canonical-0001011310222100-1020210210233023-1200232102213131-1000322223020223-1331123313210120-1002310301133320-0031011000032222-3220122200231030"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ddos profile.

Additional upstream details:

DDoS Protection Rule for DNS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_ddos_mitigation",
    "enable_ddos_mitigation")}
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
  "x-ves-oneof-field-ddos_mitigation_choice": "[\"disable_ddos_mitigation\",\"enable_ddos_mitigation\"]"
}
```

Terraform syntax:

```terraform
ddos_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202032210013213-0033002331202112-3131312311013331-1120313211211320-0300230132120313-3232311022303321-0130231131302233-0100333221303230"></a>

### Direct properties for `ddos_profile`

- [disable_ddos_mitigation](resources--dns_proxy--reference--group-001.md#canonical-2223032100103321-3003333220221323-0312212120210010-1010111021120133-3320220312033101-1331331110112102-3203221130110220-0332132303020203): complete subsection reference.

- [enable_ddos_mitigation](resources--dns_proxy--reference--group-001.md#canonical-0113322033110122-3010012031011023-0120030310322113-1011300230111202-3001303212100313-3020111302222203-2333101200222321-0223001130313013): complete subsection reference.

<a id="canonical-2223032100103321-3003333220221323-0312212120210010-1010111021120133-3320220312033101-1331331110112102-3203221130110220-0332132303020203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_profile.disable_ddos_mitigation` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [ddos_profile](resources--dns_proxy--reference--group-001.md#canonical-2110003101023121-2222231331102321-2213230320013033-1313102003322220-0010011312320132-2221021103130201-1303331100212311-3230313220231100)
- ddos_profile.disable_ddos_mitigation

<a id="canonical-3010220101223011-1120320110023122-1131311322311010-0320020012320231-3211211033031011-1121203023101130-1213101222102233-2222302232231230"></a>

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
disable_ddos_mitigation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113322033110122-3010012031011023-0120030310322113-1011300230111202-3001303212100313-3020111302222203-2333101200222321-0223001130313013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_profile.enable_ddos_mitigation` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [ddos_profile](resources--dns_proxy--reference--group-001.md#canonical-2110003101023121-2222231331102321-2213230320013033-1313102003322220-0010011312320132-2221021103130201-1303331100212311-3230313220231100)
- ddos_profile.enable_ddos_mitigation

<a id="canonical-3113303033123303-0212222032132212-3321201223310000-0330332000002111-2033131011201313-3213330213023222-1321010333211003-3230300132122330"></a>

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
enable_ddos_mitigation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220021102300002-2030313113130310-0020121013020112-3132302120121303-2302301233333100-2110200322313022-2310200003321211-3322021001010231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `irules` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- irules

<a id="canonical-3331300121110021-0220312112133330-0132310121122230-1012012103030110-1333220103230102-0030001203132003-1012100022110213-3121103222310020"></a>

Type: `"object"`. list nested block, Optional.

OPTIONS for attaching iRules to DNS proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
irules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221011313202022-0200020213213120-2032210000101330-2203033301211321-1203320311230111-1311113033010201-2231113122213230-2212021111233101"></a>

### Direct properties for `irules`

<a id="canonical-2102022231223132-3303221133131022-3303013112323333-2121220313123132-1221211003112211-0120023303310012-1100023232000213-0102321220100322"></a>

#### `irules.name` property

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

<a id="canonical-0131333322111313-1010302312323123-2133320222122232-1112313113132132-1302102330033313-3323021123010220-0212130000031113-1332313011112320"></a>

<a id="canonical-3032230203232211-1101131332012110-1332311030010021-0300100113303103-3000101131333031-2321023230021332-0302202112231010-0223103333033212"></a>

#### `irules.namespace` property

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

<a id="canonical-2123103000130133-3202320223000301-3133121123330210-3200333322002020-2310133102023212-0320013200323132-2222132321322233-2310201232000120"></a>

<a id="canonical-1032211102121011-3330222132113000-3211031310312110-0033121310212122-3022223220202111-2032002213030000-0222131130100101-3230000320313322"></a>

#### `irules.tenant` property

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

<a id="canonical-0123302030312300-1300310310103200-2010132201320120-3002100002202030-2300121210113102-3033032012230111-1301131132112112-3013012210331233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `lb_algorithm` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- lb_algorithm

<a id="canonical-0310101102223321-0010010211322010-3033212333223301-2203001033302331-2322031120212212-1221201131331322-0033301230012233-2222011032132001"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for lb algorithm.

Additional upstream details:

Load Balancing Algorithm Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lb_algorithm_choice": "[\"round_robin\"]"
}
```

Terraform syntax:

```terraform
lb_algorithm {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311321333300230-2313210103132013-1323003002112031-3232200010333102-2221313120003113-2200230203111111-3031301000013331-1001233322031013"></a>

### Direct properties for `lb_algorithm`

- [round_robin](resources--dns_proxy--reference--group-001.md#canonical-2332223102321133-1211332210000331-0313300331013111-2112332101333331-2330231002323332-2023323330023123-2110102021103332-2321321311022232): complete subsection reference.

<a id="canonical-2332223102321133-1211332210000331-0313300331013111-2112332101333331-2330231002323332-2023323330023123-2110102021103332-2321321311022232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `lb_algorithm.round_robin` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [lb_algorithm](resources--dns_proxy--reference--group-001.md#canonical-0123302030312300-1300310310103200-2010132201320120-3002100002202030-2300121210113102-3033032012230111-1301131132112112-3013012210331233)
- lb_algorithm.round_robin

<a id="canonical-3300322330131112-1103122121322122-0322222113213230-0023032132011312-3110132121003200-0233100331011331-3302012302223211-0210130001133320"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for round robin.

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
round_robin {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- origin_servers

<a id="canonical-3221133002023211-1230131123032230-2022230311222031-0303102020012112-0020303022221020-1002203322130011-3121320111332202-0101333100002012"></a>

Type: `"object"`. single nested block, Optional.

List of origin Servers for the DNS proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_servers")}
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
origin_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121233300312330-1201022210131000-3202012112332123-1322103201102003-0010013013201102-3321020133322203-0002130211132220-3202010130301233"></a>

### Direct properties for `origin_servers`

- [health_checks](resources--dns_proxy--reference--group-001.md#canonical-0223310301131311-0133221320103023-2133332022010100-1313113133020000-2002230211313010-1203013220211012-3312303211332022-2130113100232230): complete subsection reference.

- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222): complete subsection reference.

<a id="canonical-0223310301131311-0133221320103023-2133332022010100-1313113133020000-2002230211313010-1203013220211012-3312303211332022-2130113100232230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.health_checks` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- origin_servers.health_checks

<a id="canonical-0100203221200313-3321002013303203-3120022232310311-0210013033101312-1322220202230201-2102233110232232-3002130201223112-1300320222102010"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for health checks.

Additional upstream details:

Origin Server Health Checks.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check",
    "healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold")}
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
health_checks {
  # Configure direct properties listed below.
}
```

<a id="canonical-3012231302220310-2131133213232303-2212330111203220-0030333001012333-0121232232231303-3200232123120201-1133310202212222-0032222002210030"></a>

### Direct properties for `origin_servers.health_checks`

- [health_check](resources--dns_proxy--reference--group-001.md#canonical-0011102230222312-1022032120311221-3030020231301220-1033133112220102-1212330123002211-3011313200110103-2111001133202211-3000032121011230): complete subsection reference.

<a id="canonical-2010232332221211-1023233231320311-2132333200301303-2122313003221321-3112312010101122-1222122032113123-0212011101320313-3213122301323122"></a>

<a id="canonical-1210310233133030-3311121132221312-1123201010200200-0233120120222302-3001011133113000-0211201231122101-1010203330331021-3130221112001203"></a>

#### `origin_servers.health_checks.healthy_threshold` property

Type: `"number"`. Optional.

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-3132020222032100-0102112313230123-0112332330133031-3203000030003202-1210313222021311-1311312002331031-2110202101110133-2031211030023113"></a>

<a id="canonical-0320012022122120-2121313230230221-2223131110101103-3003111302013103-1033012221010203-3020033112323332-3201031220230210-1120013330131213"></a>

#### `origin_servers.health_checks.interval` property

Type: `"number"`. Optional.

Time interval in seconds between two healthcheck requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-0223113311100220-1333030331321213-0003122112031221-0012303321231020-1023300113023222-0030303210201123-0200132033212302-3332321323200230"></a>

<a id="canonical-1123310333300133-2011322010000133-0321021011221310-3220302122113201-2301112303202322-3220232202210211-2220330313030322-3212210013010002"></a>

#### `origin_servers.health_checks.timeout` property

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-2102133111203303-0112103202022321-0132100333133001-0021011320223003-2122100230333123-1220322203331310-3110220231101331-3302103103230332"></a>

<a id="canonical-1321112213212012-0012100213322002-1030123133100001-3323303200121232-1101002131012222-0322031302313230-2133311100022201-1202001303333210"></a>

#### `origin_servers.health_checks.unhealthy_threshold` property

Type: `"number"`. Optional.

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health
checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-0011102230222312-1022032120311221-3030020231301220-1033133112220102-1212330123002211-3011313200110103-2111001133202211-3000032121011230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.health_checks.health_check` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.health_checks](resources--dns_proxy--reference--group-001.md#canonical-0223310301131311-0133221320103023-2133332022010100-1313113133020000-2002230211313010-1203013220211012-3312303211332022-2130113100232230)
- origin_servers.health_checks.health_check

<a id="canonical-1212113033010213-2111220032212302-1023222113313110-1021213003012031-0011300313010013-3310103313202313-1312300332331102-1030010320123203"></a>

Type: `"object"`. list nested block, Optional.

List of Health Checks. List of Health Checks.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("dns_health_check",
    "icmp_health_check"),
  validators.ConflictingListObjectAttributes("dns_health_check",
    "tcp_health_check"),
  validators.ConflictingListObjectAttributes("icmp_health_check",
    "tcp_health_check")}
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minItems": 0,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003203012310011-3201002130021021-2023113012311001-2301023100322302-1220032030212211-2212033011132202-3321300132310113-0310231111113013"></a>

### Direct properties for `origin_servers.health_checks.health_check`

- [dns_health_check](resources--dns_proxy--reference--group-001.md#canonical-0311002023021013-3123212131013012-3232102023122332-1133132222223232-2100231333113223-0323120223013321-3223033220200132-2100333133312221): complete subsection reference.

- [icmp_health_check](resources--dns_proxy--reference--group-001.md#canonical-1301103100122311-0002330231333233-1133020323101023-0121031110302310-1330212303312023-3211013220210023-3003023230300030-2231211333122033): complete subsection reference.

- [tcp_health_check](resources--dns_proxy--reference--group-001.md#canonical-3103123232330303-0311113021101123-3331030210020302-3003030010121122-0203210221221131-3133212122221102-1323333123103031-1122233122330031): complete subsection reference.

<a id="canonical-0311002023021013-3123212131013012-3232102023122332-1133132222223232-2100231333113223-0323120223013321-3223033220200132-2100333133312221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.health_checks.health_check.dns_health_check` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.health_checks](resources--dns_proxy--reference--group-001.md#canonical-0223310301131311-0133221320103023-2133332022010100-1313113133020000-2002230211313010-1203013220211012-3312303211332022-2130113100232230)
- [origin_servers.health_checks.health_check](resources--dns_proxy--reference--group-001.md#canonical-0011102230222312-1022032120311221-3030020231301220-1033133112220102-1212330123002211-3011313200110103-2111001133202211-3000032121011230)
- origin_servers.health_checks.health_check.dns_health_check

<a id="canonical-0101023132323010-0313313211022112-1132200012310002-3110320100331010-2233202131211201-3203012302100221-1331201330122021-0212033333201301"></a>

Type: `"object"`. single nested block, Optional.

DNS health check reports healthy if DNS query is successful and response header and answer matches
the given value.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expected_response",
    "query_name")}
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
dns_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022201333201203-2331232311203210-2021023012032202-2032102311323122-2011230130003323-2020310331111320-3010021000212323-0303200221221011"></a>

### Direct properties for `origin_servers.health_checks.health_check.dns_health_check`

<a id="canonical-2101110102013222-3330103310001131-0132131123113203-1033311230101100-3122310023203032-3313113122022330-3232310120203023-3022113122001021"></a>

#### `origin_servers.health_checks.health_check.dns_health_check.expected_rcode` property

Type: `"string"`. Optional.

\[Enum: DNS\_RES\_RCODE\_NOERROR|DNS\_RES\_RCODE\_ANY\] Expected DNS Response Rcode Type -
DNS\_RES\_RCODE\_NOERROR: Rcode NOERROR - DNS\_RES\_RCODE\_ANY: RCODE ANY. Possible values are
\`DNS\_RES\_RCODE\_NOERROR\`, \`DNS\_RES\_RCODE\_ANY\`. Defaults to \`DNS\_RES\_RCODE\_NOERROR\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DNS_RES_RCODE_ANY","DNS_RES_RCODE_NOERROR"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("DNS_RES_RCODE_NOERROR",
    "DNS_RES_RCODE_ANY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DNS_RES_RCODE_NOERROR",
  "enum": [
    "DNS_RES_RCODE_NOERROR",
    "DNS_RES_RCODE_ANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1033123223131031-1132332130033100-2230302132222201-1202122230022011-2121113032001222-0133133000332330-0100013020031332-1030330213000100"></a>

<a id="canonical-1100102231121020-2101101103012313-2111000231302012-2200032021223003-1020333321323222-3211300003232213-2220233311332312-3130230123203132"></a>

#### `origin_servers.health_checks.health_check.dns_health_check.expected_record_type` property

Type: `"string"`. Optional.

\[Enum: DNS\_REQUESTED\_QUERY\_TYPE|DNS\_RES\_RECORD\_TYPE\_ANY\] DNS Response Record Type -
DNS\_REQUESTED\_QUERY\_TYPE: Requested Query Type - DNS\_RES\_RECORD\_TYPE\_ANY: Any Record Type.
Possible values are \`DNS\_REQUESTED\_QUERY\_TYPE\`, \`DNS\_RES\_RECORD\_TYPE\_ANY\`. Defaults to
\`DNS\_REQUESTED\_QUERY\_TYPE\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DNS_REQUESTED_QUERY_TYPE","DNS_RES_RECORD_TYPE_ANY"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("DNS_REQUESTED_QUERY_TYPE",
    "DNS_RES_RECORD_TYPE_ANY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DNS_REQUESTED_QUERY_TYPE",
  "enum": [
    "DNS_REQUESTED_QUERY_TYPE",
    "DNS_RES_RECORD_TYPE_ANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2032020101230121-1023120121233031-2011330113201123-3331330020113102-3023001210212012-2320103020011120-0302112332220100-2200011333321213"></a>

<a id="canonical-0032021033233311-2203022131222223-1012121332230333-2131320001100120-2022102230021023-3101213002021210-1013211021223101-2232322321202321"></a>

#### `origin_servers.health_checks.health_check.dns_health_check.expected_response` property

Type: `"string"`. Optional.

Specifies an IPv4 or IPv6 address in the answer section of DNS Response.

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
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_len": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  }
}
```

<a id="canonical-1222033020301132-1020123331313332-3113001230030332-3012203222213332-3230203302221200-1331100330311330-1113022110333232-1321313223301301"></a>

<a id="canonical-0220233131210303-1320303111220222-1020323210032111-2102333231113320-0001122122033211-2302133203221222-0103220221213020-2322130210132232"></a>

#### `origin_servers.health_checks.health_check.dns_health_check.query_name` property

Type: `"string"`. Optional.

The query name that the monitor sends a DNS query for.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2323331220021021-2131030301223330-0212322311012323-0132011311322033-0202212130020223-2202110102131322-3200102232021222-3012133331211302"></a>

<a id="canonical-3131333232102101-2011111233300112-2012323333320020-2312322232330120-0020230313132311-2000111100212301-0003312330302100-2110110303231111"></a>

#### `origin_servers.health_checks.health_check.dns_health_check.query_type` property

Type: `"string"`. Optional.

\[Enum: DNS\_QTYPE\_A|DNS\_QTYPE\_AAAA\] DNS Query Type - DNS\_QTYPE\_A: Query Type A -
DNS\_QTYPE\_AAAA: Query Type AAAA. Possible values are \`DNS\_QTYPE\_A\`, \`DNS\_QTYPE\_AAAA\`.
Defaults to \`DNS\_QTYPE\_A\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DNS_QTYPE_A","DNS_QTYPE_AAAA"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("DNS_QTYPE_A",
    "DNS_QTYPE_AAAA"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DNS_QTYPE_A",
  "enum": [
    "DNS_QTYPE_A",
    "DNS_QTYPE_AAAA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0120001113020213-2110333000212103-2031221212333013-2321210131030133-0031303021122232-3020100021111321-0223022132320032-0223232200322102"></a>

<a id="canonical-2032111131020322-0230130032302033-3033021301223122-1111013013312223-3032301230210213-3213032213022321-0313001333130020-2123101110312310"></a>

#### `origin_servers.health_checks.health_check.dns_health_check.reverse` property

Type: `"bool"`. Optional.

Enable the monitor operation in reverse mode. When the monitor is in reverse mode, a successful
receive string match marks the monitored object down instead of up.

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

<a id="canonical-1301103100122311-0002330231333233-1133020323101023-0121031110302310-1330212303312023-3211013220210023-3003023230300030-2231211333122033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.health_checks.health_check.icmp_health_check` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.health_checks](resources--dns_proxy--reference--group-001.md#canonical-0223310301131311-0133221320103023-2133332022010100-1313113133020000-2002230211313010-1203013220211012-3312303211332022-2130113100232230)
- [origin_servers.health_checks.health_check](resources--dns_proxy--reference--group-001.md#canonical-0011102230222312-1022032120311221-3030020231301220-1033133112220102-1212330123002211-3011313200110103-2111001133202211-3000032121011230)
- origin_servers.health_checks.health_check.icmp_health_check

<a id="canonical-3321213201311130-3223303133331012-0212232310331121-3023100112231323-2300312000021121-1312111230323012-0110320031233131-3002031311032122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for icmp health check.

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
icmp_health_check = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103123232330303-0311113021101123-3331030210020302-3003030010121122-0203210221221131-3133212122221102-1323333123103031-1122233122330031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.health_checks.health_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.health_checks](resources--dns_proxy--reference--group-001.md#canonical-0223310301131311-0133221320103023-2133332022010100-1313113133020000-2002230211313010-1203013220211012-3312303211332022-2130113100232230)
- [origin_servers.health_checks.health_check](resources--dns_proxy--reference--group-001.md#canonical-0011102230222312-1022032120311221-3030020231301220-1033133112220102-1212330123002211-3011313200110103-2111001133202211-3000032121011230)
- origin_servers.health_checks.health_check.tcp_health_check

<a id="canonical-2110130130311022-2221333123220030-2033213102002030-2112000212331113-0100200132113303-1233110101001211-2210111312300313-0133020313123332"></a>

Type: `"object"`. single nested block, Optional.

Monitor reports healthy status if UDP connection is successful and response payload matches expected
response pattern.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expected_response",
    "send_payload")}
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
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202130110231101-2001101000032222-0132101020203102-1101310031122012-0303202301303223-0330031302120230-2332202020000021-3313033213011332"></a>

### Direct properties for `origin_servers.health_checks.health_check.tcp_health_check`

<a id="canonical-2030222011030013-0033032202101012-1310202003333113-0301012131021131-3203202313000010-0102331223320332-0303130113011300-0020023201303032"></a>

#### `origin_servers.health_checks.health_check.tcp_health_check.expected_response` property

Type: `"string"`. Optional.

Specifies a regular expression pattern which will be matched against response payload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-0103021110300030-3213100331232313-0003221323123221-2022330200202012-3312310012203332-0020130023232033-2023323311203200-2132203301213031"></a>

<a id="canonical-3133331301122200-1310033101223022-2013003102232010-2123301211302001-0132023120103200-1320000121003121-0111132320223200-3303202002033000"></a>

#### `origin_servers.health_checks.health_check.tcp_health_check.send_payload` property

Type: `"string"`. Optional.

Send string. Text string sent in the request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- origin_servers.origin_servers

<a id="canonical-3332031133332102-1210200123202301-3131332220113013-3002120010112310-1320131201213202-1232112031021311-1113123301211013-1323021123222323"></a>

Type: `"object"`. list nested block, Optional.

List Of Origin Servers. List of origin servers for Proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("k8s_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("no_preference",
    "site_preferences"),
  validators.ConflictingListObjectAttributes("public_ip",
    "public_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010303033331332-0032133223331303-0222000210010331-2003122100012131-2012303101013010-1202321300102121-1231323321012010-0202023002131001"></a>

### Direct properties for `origin_servers.origin_servers`

- [k8s_service](resources--dns_proxy--reference--group-001.md#canonical-2121232230332301-1202331100030212-1333322300022120-1210312000311032-1101002332032103-0210032120113132-3000133133022003-3122210211120331): complete subsection reference.

- [no_preference](resources--dns_proxy--reference--group-001.md#canonical-2203011113133033-0310102221123301-2011031132213022-0101011033223133-3032322323030112-0332103330110031-0200130113202113-0122220022321102): complete subsection reference.

- [public_ip](resources--dns_proxy--reference--group-001.md#canonical-0332203222013220-0130321101032111-2012303000311332-3112321103103132-1013310013011203-1231231132310313-1133221133303333-1132322120103311): complete subsection reference.

- [public_name](resources--dns_proxy--reference--group-001.md#canonical-1112101213302101-0223233022003002-3131030320013332-1303322131203320-0032301113313020-0123023112021322-0320202221133021-1100001201010321): complete subsection reference.

- [site_preferences](resources--dns_proxy--reference--group-001.md#canonical-2312211011120111-3323213032322320-2331232323202001-1223333121003231-0131002122113102-2002010310202132-2002130313312330-3113200213003311): complete subsection reference.

<a id="canonical-2121232230332301-1202331100030212-1333322300022120-1210312000311032-1101002332032103-0210032120113132-3000133133022003-3122210211120331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.k8s_service` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- origin_servers.origin_servers.k8s_service

<a id="canonical-2311031010122023-0020203300022113-3113231100221022-0202232100323120-2130220120312121-3313203322211200-3000003133031003-1211200322230000"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with K8s service name and site information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "vk8s_networks"),
  validators.ConflictingObjectAttributes("outside_network",
    "vk8s_networks")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

Terraform syntax:

```terraform
k8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012230001130011-2331330112130010-3202233021012130-3110322113331011-2000200320033022-1031320032211201-3230022133130223-2302322332310202"></a>

### Direct properties for `origin_servers.origin_servers.k8s_service`

- [inside_network](resources--dns_proxy--reference--group-001.md#canonical-3302031120131102-3130220213011032-1021223131113222-0032201131230100-1101121320232333-0301011010033120-3120230213332113-2321003021200003): complete subsection reference.

- [outside_network](resources--dns_proxy--reference--group-001.md#canonical-3220222322013210-3331022101000320-1010021033313222-2312211133332323-0212322000213330-3021130202312212-1122231102122210-1210231131020132): complete subsection reference.

<a id="canonical-0203320220301132-2211210333300231-2100202230012110-1112322331203023-1110131000000001-0202121301111023-1021211200033230-0212001132100310"></a>

<a id="canonical-2020111132013330-3311120301023212-1111130121301311-0101311130130101-1120031022210121-2322112132132002-0221232233102222-1130310223333022"></a>

#### `origin_servers.origin_servers.k8s_service.protocol` property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["PROTOCOL_TCP","PROTOCOL_UDP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1333213032322021-1121312333300221-1031122323312122-1130221231313210-0301013201101331-0323201231301230-0300003133132100-3230230221101231"></a>

<a id="canonical-0201213022121023-2113031013200323-1112003012011313-3130021233100303-1033113330222310-0223202002230122-2123301032230302-1323211101301300"></a>

#### `origin_servers.origin_servers.k8s_service.service_name` property

Type: `"string"`. Optional.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Additional upstream details:

For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](resources--dns_proxy--reference--group-001.md#canonical-2232012322211121-2112113112132320-3231202010133100-3010212130302022-0132211100120101-2112113332221323-0223103133302211-1332303121030322): complete subsection reference.

- [snat_pool](resources--dns_proxy--reference--group-001.md#canonical-1030031312233302-1100213103030132-0201110210122330-0220020133023231-3223231120331101-1111302323112012-2303202003012112-3000202200330022): complete subsection reference.

- [vk8s_networks](resources--dns_proxy--reference--group-001.md#canonical-1223130113001301-1130210223200122-0021100111031233-1212011002322211-2111000211101231-0131311330130000-0312122121212211-3020120032000022): complete subsection reference.

<a id="canonical-3302031120131102-3130220213011032-1021223131113222-0032201131230100-1101121320232333-0301011010033120-3120230213332113-2321003021200003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.k8s_service.inside_network` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-2121232230332301-1202331100030212-1333322300022120-1210312000311032-1101002332032103-0210032120113132-3000133133022003-3122210211120331)
- origin_servers.origin_servers.k8s_service.inside_network

<a id="canonical-0331130101111121-3111230320133002-3132033200202303-0221003323232102-3102113302020232-0013002202303201-1113221321223021-1031322310210202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220222322013210-3331022101000320-1010021033313222-2312211133332323-0212322000213330-3021130202312212-1122231102122210-1210231131020132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.k8s_service.outside_network` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-2121232230332301-1202331100030212-1333322300022120-1210312000311032-1101002332032103-0210032120113132-3000133133022003-3122210211120331)
- origin_servers.origin_servers.k8s_service.outside_network

<a id="canonical-0313112012333031-2101230211010210-0201013201033333-3220100130303121-1213000312022200-0321123330332330-1330202313312020-3321133312321333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232012322211121-2112113112132320-3231202010133100-3010212130302022-0132211100120101-2112113332221323-0223103133302211-1332303121030322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.k8s_service.site_locator` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-2121232230332301-1202331100030212-1333322300022120-1210312000311032-1101002332032103-0210032120113132-3000133133022003-3122210211120331)
- origin_servers.origin_servers.k8s_service.site_locator

<a id="canonical-0131123131302023-2003033111112021-1312322222231113-3330321122312331-3310100200203002-3011132123222003-2121002123312232-0220020111321002"></a>

Type: `"object"`. single nested block, Optional.

This message defines a reference to a site or virtual site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100112312011203-0222033312100112-3210332200212212-0213123032101332-0320001333103200-3233230032302102-2001310133330213-3130011332001333"></a>

### Direct properties for `origin_servers.origin_servers.k8s_service.site_locator`

- [site](resources--dns_proxy--reference--group-001.md#canonical-0200002003023121-1020331310011031-1230320031303200-0313123332010210-1233221203223103-2112321100233303-3221302031323100-1101133322200300): complete subsection reference.

- [virtual_site](resources--dns_proxy--reference--group-001.md#canonical-1002120310013321-0203231210112113-3202231311230211-2200221331203110-3023010122031101-0113211330223111-3302333120101332-3101331322012300): complete subsection reference.

<a id="canonical-0200002003023121-1020331310011031-1230320031303200-0313123332010210-1233221203223103-2112321100233303-3221302031323100-1101133322200300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.k8s_service.site_locator.site` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-2121232230332301-1202331100030212-1333322300022120-1210312000311032-1101002332032103-0210032120113132-3000133133022003-3122210211120331)
- [origin_servers.origin_servers.k8s_service.site_locator](resources--dns_proxy--reference--group-001.md#canonical-2232012322211121-2112113112132320-3231202010133100-3010212130302022-0132211100120101-2112113332221323-0223103133302211-1332303121030322)
- origin_servers.origin_servers.k8s_service.site_locator.site

<a id="canonical-3012023321122213-0203013310002233-0330230002111231-3011132231012010-1021100111211303-3210101030111100-1310122321033220-0030032222113320"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121030012210222-3000200333010033-0321120332213310-3120003112230011-2020122023233311-1103220320120113-3102323233002100-1121331121120132"></a>

### Direct properties for `origin_servers.origin_servers.k8s_service.site_locator.site`

<a id="canonical-3133331112201330-2221323000032231-3322312201110303-2332322333113122-2121213121313232-3331332120232230-0032203311130133-2013002110212222"></a>

#### `origin_servers.origin_servers.k8s_service.site_locator.site.name` property

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

<a id="canonical-2212111320331020-3100123221103200-0300020333030303-2313313033213221-1221132300011003-0221012331333003-2332101322233133-0301333213330223"></a>

<a id="canonical-2120232003201213-3133330011230003-1331012032013113-3230211220233322-1200033230032101-3312101200022131-0313021023003303-2303302132310332"></a>

#### `origin_servers.origin_servers.k8s_service.site_locator.site.namespace` property

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

<a id="canonical-1033021203033110-0110322113233301-3022023230012232-0002002302322333-1211120213002300-3010123302122301-2230122203031123-3100200203101130"></a>

<a id="canonical-0233110211023220-1203311010011211-3010313130101100-0121023220233221-0110312033210223-1012212021331201-1102221222011323-3301023131223300"></a>

#### `origin_servers.origin_servers.k8s_service.site_locator.site.tenant` property

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

<a id="canonical-1002120310013321-0203231210112113-3202231311230211-2200221331203110-3023010122031101-0113211330223111-3302333120101332-3101331322012300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.k8s_service.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-2121232230332301-1202331100030212-1333322300022120-1210312000311032-1101002332032103-0210032120113132-3000133133022003-3122210211120331)
- [origin_servers.origin_servers.k8s_service.site_locator](resources--dns_proxy--reference--group-001.md#canonical-2232012322211121-2112113112132320-3231202010133100-3010212130302022-0132211100120101-2112113332221323-0223103133302211-1332303121030322)
- origin_servers.origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-0322203112313002-0130113313212101-2020011003331333-0302120303312211-1202000023223110-3321020311223231-1213230121011022-3220001101312203"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303013220311201-0320333203201003-2122102012332200-1123210210111130-0230331300123122-0133101011020222-2232321001213300-1331100132223013"></a>

### Direct properties for `origin_servers.origin_servers.k8s_service.site_locator.virtual_site`

<a id="canonical-2311031212223221-1213202210320320-2110101201100213-1110012202320032-3021102113301323-2221332210021300-1131011031303233-1120122210001201"></a>

#### `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` property

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

<a id="canonical-0000000232323103-3313310332113203-2103132221133311-0030322033130010-3131021122103120-3102232323020113-0212322222000103-0133333102000313"></a>

<a id="canonical-3331001323023203-0002221232211121-2111231300301131-0203131122201200-1303320100230102-0032133131101232-0101122022300232-1132130010003233"></a>

#### `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` property

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

<a id="canonical-3301300223110303-3101233313202020-0030333312113000-0310320332302122-0113221331013223-0213201020302030-2032033032133302-0121030311002011"></a>

<a id="canonical-3102333300000313-1131222311132311-0113200112320212-2320323013223323-3222020202133130-1310203221231032-3203321001230013-2200330221030330"></a>

#### `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` property

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

<a id="canonical-1030031312233302-1100213103030132-0201110210122330-0220020133023231-3223231120331101-1111302323112012-2303202003012112-3000202200330022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.k8s_service.snat_pool` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-2121232230332301-1202331100030212-1333322300022120-1210312000311032-1101002332032103-0210032120113132-3000133133022003-3122210211120331)
- origin_servers.origin_servers.k8s_service.snat_pool

<a id="canonical-3012010102013131-1120232112231130-1001013011013012-2112122131332232-0030102221203312-0321111102332132-1221120033222300-1130030100033033"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030013330311202-3101322010100010-3013332011003012-1112130330132100-2322023002133033-1313021220202022-2010120100121031-1300013003011120"></a>

### Direct properties for `origin_servers.origin_servers.k8s_service.snat_pool`

- [no_snat_pool](resources--dns_proxy--reference--group-001.md#canonical-2312100302012320-3110320303100331-1233100103222012-0202210311321313-2020130311131303-1121333123002012-3000310312302122-1011330200312232): complete subsection reference.

- [snat_pool](resources--dns_proxy--reference--group-001.md#canonical-0310023313013120-0013121001322102-3111232210211221-1102223131023201-0232121120323020-0312332301003003-2123130201121330-0230300010333303): complete subsection reference.

<a id="canonical-2312100302012320-3110320303100331-1233100103222012-0202210311321313-2020130311131303-1121333123002012-3000310312302122-1011330200312232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-2121232230332301-1202331100030212-1333322300022120-1210312000311032-1101002332032103-0210032120113132-3000133133022003-3122210211120331)
- [origin_servers.origin_servers.k8s_service.snat_pool](resources--dns_proxy--reference--group-001.md#canonical-1030031312233302-1100213103030132-0201110210122330-0220020133023231-3223231120331101-1111302323112012-2303202003012112-3000202200330022)
- origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-1010031322233330-3301033023121320-3212322021301330-3203220133322013-2213202322210223-1312012112333010-1031323112031201-0131210321003222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310023313013120-0013121001322102-3111232210211221-1102223131023201-0232121120323020-0312332301003003-2123130201121330-0230300010333303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-2121232230332301-1202331100030212-1333322300022120-1210312000311032-1101002332032103-0210032120113132-3000133133022003-3122210211120331)
- [origin_servers.origin_servers.k8s_service.snat_pool](resources--dns_proxy--reference--group-001.md#canonical-1030031312233302-1100213103030132-0201110210122330-0220020133023231-3223231120331101-1111302323112012-2303202003012112-3000202200330022)
- origin_servers.origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-0110213033211201-3220233013112013-2330011021021030-2032123232222232-3000023210212300-1332202023210133-1332320111103130-0203121132232203"></a>

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102033021302032-2121022211120111-2221322031013323-1302230323002123-1230002012302201-0030222212312113-2210323101300210-2323212122211221"></a>

### Direct properties for `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool`

<a id="canonical-3232202221133012-0330232022321213-0232110102203000-2302331030032030-2320222222221013-2333303123322131-1131331311120312-0222220000132031"></a>

#### `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1223130113001301-1130210223200122-0021100111031233-1212011002322211-2111000211101231-0131311330130000-0312122121212211-3020120032000022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.k8s_service.vk8s_networks` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--reference--group-001.md#canonical-2121232230332301-1202331100030212-1333322300022120-1210312000311032-1101002332032103-0210032120113132-3000133133022003-3122210211120331)
- origin_servers.origin_servers.k8s_service.vk8s_networks

<a id="canonical-2311301330310123-3032333113020100-1003203223223133-0021332302013232-3132323232311112-2023030031332020-0332032231103200-1330123130221000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vk8s networks.

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
vk8s_networks = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203011113133033-0310102221123301-2011031132213022-0101011033223133-3032322323030112-0332103330110031-0200130113202113-0122220022321102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.no_preference` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- origin_servers.origin_servers.no_preference

<a id="canonical-0302230133011112-0313003031302002-3211113031213231-1222203330120100-0113030131230113-0030021030120101-1113321020022101-3300332200023233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no preference.

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
no_preference = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332203222013220-0130321101032111-2012303000311332-3112321103103132-1013310013011203-1231231132310313-1133221133303333-1132322120103311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.public_ip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- origin_servers.origin_servers.public_ip

<a id="canonical-1033301121023032-2221003220202103-2003300000301031-3023313211313202-0003213001201023-2300311112222023-1231012312312300-0131201001201103"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210313133133233-1211111132200010-1201013223112300-2211132130023332-3112212321020232-2121230113202032-0332212001112133-1231330132121130"></a>

### Direct properties for `origin_servers.origin_servers.public_ip`

<a id="canonical-1310200022000010-2011333000322233-3100020313121221-2023031222112332-0303001220020111-2331232321333002-1003012213130033-1212113003221330"></a>

#### `origin_servers.origin_servers.public_ip.ip` property

Type: `"string"`. Optional.

Public IPv4. Exclusive with \[\] Public IPv4 address.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1112101213302101-0223233022003002-3131030320013332-1303322131203320-0032301113313020-0123023112021322-0320202221133021-1100001201010321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.public_name` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- origin_servers.origin_servers.public_name

<a id="canonical-2310230113110121-3233002112131232-0310130223033222-0200223001221330-2320002233020210-0102000312212213-1301200221100020-2320132013030200"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
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
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130122230113301-3210221302213213-2022021021111313-0200213220301220-3033210001211111-3332000301312200-1320311320312101-1223303110020012"></a>

### Direct properties for `origin_servers.origin_servers.public_name`

<a id="canonical-1112213313131102-3202313100210021-0200013023330011-1322121013300022-1232131222312011-2222013310311100-0032111022022112-2020023022331302"></a>

#### `origin_servers.origin_servers.public_name.dns_name` property

Type: `"string"`. Optional.

DNS Name. DNS Name

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1031332023310330-3311302011121020-3010103312213101-0210202200220101-1213331002032112-2032103213101022-1233231302212003-1300222130101002"></a>

<a id="canonical-0333233020312033-1000113300102200-0120212001210202-1200030020201103-0202210002322102-1200030331232311-0313110232023321-1321130312002331"></a>

#### `origin_servers.origin_servers.public_name.refresh_interval` property

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-2312211011120111-3323213032322320-2331232323202001-1223333121003231-0131002122113102-2002010310202132-2002130313312330-3113200213003311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.site_preferences` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- origin_servers.origin_servers.site_preferences

<a id="canonical-2310210330121011-3222112222131231-1003321110121111-0012012001200313-3312031013310331-2032130120003330-0112320121310203-3130221331331011"></a>

Type: `"object"`. single nested block, Optional.

Carries the references to one or more sites.

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
site_preferences {
  # Configure direct properties listed below.
}
```

<a id="canonical-0312101311211213-3330113213210230-0131220021031121-3303321120323123-2210131123030330-2231230100212122-0112221201212332-3321122011201020"></a>

### Direct properties for `origin_servers.origin_servers.site_preferences`

- [refs](resources--dns_proxy--reference--group-001.md#canonical-2120300301031303-1233002111003102-1302013213232131-0132102101333022-1221013123231233-2321202333002232-3002010333110100-3201212301332100): complete subsection reference.

<a id="canonical-2120300301031303-1233002111003102-1302013213232131-0132102101333022-1221013123231233-2321202333002232-3002010333110100-3201212301332100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.origin_servers.site_preferences.refs` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [origin_servers](resources--dns_proxy--reference--group-001.md#canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111)
- [origin_servers.origin_servers](resources--dns_proxy--reference--group-001.md#canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222)
- [origin_servers.origin_servers.site_preferences](resources--dns_proxy--reference--group-001.md#canonical-2312211011120111-3323213032322320-2331232323202001-1223333121003231-0131002122113102-2002010310202132-2002130313312330-3113200213003311)
- origin_servers.origin_servers.site_preferences.refs

<a id="canonical-2102131103302310-3002111230231211-1200120301311023-0311322033302030-1212023203303310-3210320330123013-3211201022103012-2113231232121012"></a>

Type: `"object"`. list nested block, Optional.

Site References. Reference to one or more sites.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
    "ves.io.schema.rules.repeated.max_items": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64"
  }
}
```

Terraform syntax:

```terraform
refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220023332321133-0100310302202301-3101211201213122-1221332223031020-2012032233301202-0222230232130133-0203210101122210-1331022331003001"></a>

### Direct properties for `origin_servers.origin_servers.site_preferences.refs`

<a id="canonical-0320110110100310-0223021231130130-1322333100223211-3121033001112020-1310020130202210-0030231102103133-0102131332130303-2111300210031232"></a>

#### `origin_servers.origin_servers.site_preferences.refs.name` property

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

<a id="canonical-1311201110031220-2301231233011211-2332101103032313-1130313000130213-3230132210130002-1130113131323332-0032100031310233-0223000111303010"></a>

<a id="canonical-3020311223032321-3211012221332110-0122133301202123-2230020013331033-2222022000013230-1001030110130001-2133311032100323-1333232323011123"></a>

#### `origin_servers.origin_servers.site_preferences.refs.namespace` property

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

<a id="canonical-3203032112211002-2310011112032010-0231201232123133-3230111310132002-2102030300022132-2101300123310121-0023120020112012-1013020313202110"></a>

<a id="canonical-3120212003000301-1321000322221302-3003221003121322-3311110010102313-0130111332233111-2133003100323221-1020201331111013-3201100130011023"></a>

#### `origin_servers.origin_servers.site_preferences.refs.tenant` property

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

<a id="canonical-3220331331332203-0133330223020233-2010201120310110-3031102331001233-2301333130020221-1130222023012123-1332133111310213-1132022312202301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_inspection` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- protocol_inspection

<a id="canonical-2310103002211302-2333002012333323-1211022012123320-2332232132213310-2021213021300202-2331332311032210-1202133230230131-1313312311321230"></a>

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
protocol_inspection {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233210213211210-0212000123100301-0203130203001323-1301321231033213-3113133313231330-0232200121332233-1001001200232000-1020302203132031"></a>

### Direct properties for `protocol_inspection`

<a id="canonical-1000221230311033-3001102011222123-0300232313300212-0301223221111212-2232302010233231-3031312102121331-3300101001310300-3112133231322211"></a>

#### `protocol_inspection.name` property

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

<a id="canonical-2321310020010123-0321003311122312-0322321030030001-3132321313031233-3222023001321201-1201110223203202-1310330201033332-3233210112123033"></a>

<a id="canonical-3031300200113221-1230112001201012-3023302321212233-0332120221113311-3320121120310302-3222010030122023-2231131210030011-3221223111011232"></a>

#### `protocol_inspection.namespace` property

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

<a id="canonical-0332230333311013-3301121001011313-3101023110310031-1112102023130101-3213031333031321-0200111313333233-3112233301200212-1131333132302011"></a>

<a id="canonical-1330110331222031-3133002113113321-0321212223112210-3033000033000111-0210132003233302-2112002011133231-0121313031211103-0122110101101210"></a>

#### `protocol_inspection.tenant` property

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

<a id="canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- proxy_advertisement

<a id="canonical-0021201222031123-2101223120210311-2332222111003122-3220102022320323-2122023022130101-3103100010231030-3231211331312003-0311310010121100"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for proxy advertisement.

Additional upstream details:

Proxy Advertisement Type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_dualstack_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public_default_dualstack_vip"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public_default_ipv6_vip"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public_default_dualstack_vip"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public_default_ipv6_vip"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "advertise_on_public_default_dualstack_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "advertise_on_public_default_ipv6_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_dualstack_vip",
    "advertise_on_public_default_ipv6_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_dualstack_vip",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_dualstack_vip",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_dualstack_vip",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_ipv6_vip",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_ipv6_vip",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_ipv6_vip",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_vip",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_vip",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_v6_on_public",
    "do_not_advertise")}
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
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_dualstack_on_public\",\"advertise_on_public\",\"advertise_on_public_default_dualstack_vip\",\"advertise_on_public_default_ipv6_vip\",\"advertise_on_public_default_vip\",\"advertise_v6_on_public\",\"do_not_advertise\"]"
}
```

Terraform syntax:

```terraform
proxy_advertisement {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313131333112312-1303323213011013-0133223111211110-2322201101312103-2121111312202102-2012002332101120-1101113303023030-2212301232303233"></a>

### Direct properties for `proxy_advertisement`

- [advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101): complete subsection reference.

- [advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-0312321001033003-2000102210132110-2013301302220000-1320023122223111-1132300310003200-1203011203201332-2211111321321133-0310222130312302): complete subsection reference.

- [advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-2220303323213103-2110110021231302-1232003120023021-3003122123002120-0021020320000023-3232200110132031-0222010121032220-2103301122232110): complete subsection reference.

- [advertise_on_public_default_dualstack_vip](resources--dns_proxy--reference--group-002.md#canonical-3121000120202100-0310113221300320-2102133012003031-0222233000032201-3001032020133122-2001303133003202-3101011113003301-0002310313213122): complete subsection reference.

- [advertise_on_public_default_ipv6_vip](resources--dns_proxy--reference--group-002.md#canonical-2220312223313100-0021103312020302-1331000023323330-1112111302001112-1233121002000332-3311110001302010-1233221013011120-2221310113023300): complete subsection reference.

- [advertise_on_public_default_vip](resources--dns_proxy--reference--group-002.md#canonical-1222203232301300-1301320120330311-0033333301311202-3100200223211001-3002302013120130-1132130033200011-3033130130111121-3021122101310213): complete subsection reference.

- [advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-0010212203201301-0332200213312212-1220010333030212-2201021203301200-0030321311132322-0001222133122102-1212100023322133-0113000303320123): complete subsection reference.

- [do_not_advertise](resources--dns_proxy--reference--group-002.md#canonical-0323231032132032-3103023010222030-3122002213313020-2200330021120102-0010211000032013-0320332021031131-0020201322001221-1333331311212313): complete subsection reference.

<a id="canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.advertise_custom

<a id="canonical-3213132112221312-3300012330211231-0133123031022133-1130231001332223-2132023132230013-3322130122320323-2223311312322333-2231310233300000"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a VIP on specific sites.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where")}
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
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302200023100230-2000111230101333-2011011312322123-2201332231133220-0101230301033211-3031330121323102-2230012231211102-0022020221123303"></a>

### Direct properties for `proxy_advertisement.advertise_custom`

- [advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220): complete subsection reference.

<a id="canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- proxy_advertisement.advertise_custom.advertise_where

<a id="canonical-3003223223300310-3011211031123323-1201331300020311-1200221300002310-1212102011313022-0133233033320023-2032300332031222-1201323323221333"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingListObjectAttributes("port",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("port_ranges",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site_with_vip",
    "vk8s_service")}
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
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003300001221301-0233101223313210-3231101323220312-0031212222232301-0200111001212212-3223022213013032-2212123220222032-1011213202110001"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where`

- [advertise_dualstack_on_public](resources--dns_proxy--reference--group-001.md#canonical-1210302231200223-3200232113022130-0220223102302232-3000323111130131-1101121322121130-1100201211301122-3202031033020112-3030021320132313): complete subsection reference.

- [advertise_on_public](resources--dns_proxy--reference--group-001.md#canonical-2330201123331113-3300203000032202-3113132212311201-3233330303301201-0121303030001130-3030122132211110-3013222131331201-0131020100202303): complete subsection reference.

- [advertise_v6_on_public](resources--dns_proxy--reference--group-001.md#canonical-1120333322121220-3021032102210113-0133220302303212-3220302231211022-1231132230302013-3013301002003330-1210223200002302-3213023211201323): complete subsection reference.

<a id="canonical-1130132230110332-0321221103110123-1232002310000133-0231111301020130-3313133233232222-0200023030020230-3220013132333001-3132312310322112"></a>

<a id="canonical-0232313222320012-1321122021101330-0001321200300000-3020003130223002-1231113322312122-3100232301330032-3213202022230100-0123231312031113"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3033003112303210-0110122122110233-1123031230332003-0211010200133120-0221230301003030-0311033232101012-1101200131002012-1311022033000223"></a>

<a id="canonical-2030303033030210-1133030010223321-2022001320100011-2133331120321210-1032103200001001-3022200323301310-3131300101022310-1313333201123013"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

- [site](resources--dns_proxy--reference--group-001.md#canonical-0333220002313122-1230023033012223-0011103000031311-1102313210121022-3120111101003332-0300211233200311-2023302120013213-3202030012302023): complete subsection reference.

- [use_default_port](resources--dns_proxy--reference--group-001.md#canonical-2112301032313320-2121103122013203-2030031013320313-0002330002202003-2330120231030121-1130322300222121-3013201211233313-2002013122031333): complete subsection reference.

- [virtual_network](resources--dns_proxy--reference--group-001.md#canonical-2310102012101212-1313110121233032-0222012221233332-0332003311302302-3001031232200013-1030223210030220-3330302311103220-0202103201323322): complete subsection reference.

- [virtual_site](resources--dns_proxy--reference--group-002.md#canonical-2121010310201023-0012033031302111-2200021322330221-2310000013030030-0102122021131301-1220332330110022-3322322113031011-0121112303002232): complete subsection reference.

- [virtual_site_with_vip](resources--dns_proxy--reference--group-002.md#canonical-3120210322311313-1031233122013201-3010103131230133-1220113202032033-0022231001022221-1121230121113102-2203301120120222-1130010230032210): complete subsection reference.

- [vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-3331020230023233-2321211121222102-3023033122212123-2101133133002010-1023311023330020-2220002321331011-0202230202222210-2111222102020333): complete subsection reference.

<a id="canonical-1210302231200223-3200232113022130-0220223102302232-3000323111130131-1101121322121130-1100201211301122-3202031033020112-3030021320132313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-2213121001221203-2023133202223113-3100312111202303-3220233321331230-2213113221011002-2310132211312121-2201212333230202-0130311213022110"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_dualstack_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021132320102121-2133133330111110-3102323120312002-2102310132023132-2330113120212132-3222133010311003-0101001110323112-3313102221222033"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public`

- [public_ip](resources--dns_proxy--reference--group-001.md#canonical-0113000330020233-1322203022320013-2100203102010001-0130233121233201-1202030312212132-0201103131322201-0032031332323000-3111001223011203): complete subsection reference.

<a id="canonical-0113000330020233-1322203022320013-2100203102010001-0130233121233201-1202030312212132-0201103131322201-0032031332323000-3111001223011203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--dns_proxy--reference--group-001.md#canonical-1210302231200223-3200232113022130-0220223102302232-3000323111130131-1101121322121130-1100201211301122-3202031033020112-3030021320132313)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-0101002010331332-0033103031331232-1312121201100222-0020210021311301-1233320003333001-3301011112310231-3203012033213210-0013123330300321"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113233313011331-0231302211321003-2301230122130010-3031223301132131-0332233200000102-3232232123121120-2132233010332101-0103313221102000"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip`

<a id="canonical-0101002110311020-1001021310031220-2320202330103030-2101100111223322-2302031312221303-3011230113120211-2211130020003201-1101230030001103"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` property

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

<a id="canonical-1030323332322320-2021200012012203-3332111303333123-0210301101133002-2001231002300022-1013311323022303-2303210021010023-1202323333302122"></a>

<a id="canonical-0003200301322103-0211302232311233-3220111100102332-2103111323110312-0220223112003301-1103130300012313-0223023210012212-2311203033330231"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` property

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

<a id="canonical-2302112001123302-0121212112132102-1031121001313002-0002321223032133-3031121103001002-3330300130223121-0230300003212123-0120332130322332"></a>

<a id="canonical-1213303002121203-2331330301232330-2133002022322213-2121323023313122-3201221013332331-1322000231330131-0022213100303321-0302203312133120"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` property

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

<a id="canonical-2330201123331113-3300203000032202-3113132212311201-3233330303301201-0121303030001130-3030122132211110-3013222131331201-0131020100202303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public

<a id="canonical-2301300312103302-0312310200132131-1313010010301011-1323023113302020-1110311312303132-2221122200022310-2032020303211003-3122233120202031"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021011223112020-1301123111131013-3103212211120113-3030210330003100-0112322221333112-0131003320031012-0112131312212010-3233321320110100"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public`

- [public_ip](resources--dns_proxy--reference--group-001.md#canonical-3113231113312211-0320001121010222-2001203011233030-0333120020303233-3302101011020303-0012012331100231-2222310323021102-2233313122321223): complete subsection reference.

<a id="canonical-3113231113312211-0320001121010222-2001203011233030-0333120020303233-3302101011020303-0012012331100231-2222310323021102-2233313122321223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--dns_proxy--reference--group-001.md#canonical-2330201123331113-3300203000032202-3113132212311201-3233330303301201-0121303030001130-3030122132211110-3013222131331201-0131020100202303)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-1111130223301320-1101111331101331-1300133030201233-2212331321111303-0133100220231212-0023302213121000-0302213313022322-3211021320203112"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303102332202021-0312122010030333-0130110231313022-0333013311113133-3120003123131310-2002031320112231-0320331112011032-3230113211101102"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip`

<a id="canonical-3231101301213200-2231022302002202-0230230233232333-3202200331023123-2030111222313233-3120210002203011-3001313330003001-3313003202001113"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` property

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

<a id="canonical-2202202230001320-1332000303213111-0132213132310300-3231222033100212-3211220121220011-2003013311113103-1221022330300122-0313233220011333"></a>

<a id="canonical-3203012031232300-3231103001001101-3002201023103130-3133313322112203-2213120030222002-3330233311221331-3333203200121121-1302012000100111"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` property

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

<a id="canonical-3323202010223220-1302223012100233-0120231330303101-3222232213210133-3102023123031323-3021332112001022-1210332332311232-1013113131111231"></a>

<a id="canonical-1233032112203203-1020233233132223-0210332312210211-1020211000131023-3023222010200110-0313133231322020-2103320322120310-1200102212011203"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` property

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

<a id="canonical-1120333322121220-3021032102210113-0133220302303212-3220302231211022-1231132230302013-3013301002003330-1210223200002302-3213023211201323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-2212301320331322-0123000231200203-1111121232133032-1032121112310320-0330020003110130-3112110310131020-1200311003003121-0231322301330021"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_v6_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302310120211210-0130210230333003-1300102230331003-3103333031330301-3000113210230212-0221021011220333-1303322001303230-3100013100111320"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public`

- [public_ip](resources--dns_proxy--reference--group-001.md#canonical-3201202033131101-0031331311311313-0310023021330133-2101300100231130-0221321121001222-1011011310332110-3030033111232320-2010233012213031): complete subsection reference.

<a id="canonical-3201202033131101-0031331311311313-0310023021330133-2101300100231130-0221321121001222-1011011310332110-3030033111232320-2010233012213031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--dns_proxy--reference--group-001.md#canonical-1120333322121220-3021032102210113-0133220302303212-3220302231211022-1231132230302013-3013301002003330-1210223200002302-3213023211201323)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-2221201302323300-3333022111320103-2211320000312123-1102131012110121-0033020020232232-2213320101321122-2101111211211011-2323310003012000"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123000130332032-0312333201123103-3320002121020201-0032122321220120-3132032013122220-1321130323213221-3111331112331333-0232303100033333"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip`

<a id="canonical-3201010212011110-1101233311202011-0202000000123211-2133332101113121-0121222122221120-1333201021002302-0133122020322231-3211001033233100"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` property

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

<a id="canonical-3102122002113122-3300300131120122-1113122333230200-2302010010011131-1131221322111010-2310330201003212-3332202133022003-1321301132100023"></a>

<a id="canonical-2313221121222121-3203113310101300-1113202012330101-1023021100030022-0000102113020003-1210011321200112-0231112233101300-3332103121120232"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` property

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

<a id="canonical-2020300013132010-0233030301132132-3121212223130211-1000323302303121-3002233022302032-2311201203111032-2030010310310003-1022031331113003"></a>

<a id="canonical-3220213310231303-2122320113320003-0301032301132131-0220223200333330-3010130103332133-1210211113123122-0202311313110200-3133310130100313"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` property

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

<a id="canonical-0333220002313122-1230023033012223-0011103000031311-1102313210121022-3120111101003332-0300211233200311-2023302120013213-3202030012302023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.site` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.site

<a id="canonical-2313301120110301-3332023233032012-1201123231210211-2012022203130103-3213311223123101-0212222202132033-2000330003321103-2221302003022032"></a>

Type: `"object"`. single nested block, Optional.

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033313133220203-2213030131322303-2013022321102122-0031121002003232-0100023020033301-3003302101022212-2321310031130220-1322111131201230"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.site`

<a id="canonical-3201111200013122-0102203133302121-3032230313013310-1123302131020131-0030211101103103-0132213022132003-3220202302223332-3010122021002122"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.ip` property

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1101031132033003-1110311021030231-1320231223211013-2002220113303210-3211201230003320-1301320032012110-3001310132330301-0033120312332011"></a>

<a id="canonical-0323303221220220-2332102110000131-2121123031312311-2102210202130213-2302221223223022-3232113001203103-0113320132011022-1332102021202010"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.network` property

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_NETWORK_INSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_IP_FABRIC","SITE_NETWORK_OUTSIDE","SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_SERVICE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](resources--dns_proxy--reference--group-001.md#canonical-1121030032133221-0203020103230201-2311002322003103-2203232022211100-1133333333013223-2220010301221110-1323320012020221-0123303220133132): complete subsection reference.

<a id="canonical-1121030032133221-0203020103230201-2311002322003103-2203232022211100-1133333333013223-2220010301221110-1323320012020221-0123303220133132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.site.site` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.site](resources--dns_proxy--reference--group-001.md#canonical-0333220002313122-1230023033012223-0011103000031311-1102313210121022-3120111101003332-0300211233200311-2023302120013213-3202030012302023)
- proxy_advertisement.advertise_custom.advertise_where.site.site

<a id="canonical-3121012021321022-1030012101200203-3130002231220300-1110031221000032-2203321213002333-1011121123133032-0232203123300102-2233132331232310"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021313232332332-0200023032111022-2332210212100010-1123010003001230-0200001111320122-2130223101031123-1331100232113102-2211201312103330"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.site.site`

<a id="canonical-1313223333010210-2012223303213231-3023002210222033-3201203101313220-3232223323323312-1211020201033010-1103102132333020-0313212220230210"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.site.name` property

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

<a id="canonical-1231323310120323-1232220000320223-1111022232022230-1321321310220332-2310202310032330-3322310010001023-0001323201023231-2312110103212232"></a>

<a id="canonical-1023211221103031-3002020332220330-2101223020303031-0122012230221303-3120110220102102-2020313003212101-0203010021200230-3320221021310323"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` property

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

<a id="canonical-0100000320101122-0000210011223003-2133220103122002-1202320200133322-1011221010112223-2303312121101122-0102211031033322-3030130323231133"></a>

<a id="canonical-3013123110022201-1130233133302310-1212320302333311-1103223100112303-0210303000321312-2223033103203020-2113320213300232-0102213222022033"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` property

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

<a id="canonical-2112301032313320-2121103122013203-2030031013320313-0002330002202003-2330120231030121-1130322300222121-3013201211233313-2002013122031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.use_default_port` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.use_default_port

<a id="canonical-1133113321203103-3101030303002133-1022220023012200-1012022331030232-0211212130232033-2332001201010201-3320012112230023-2210110123321301"></a>

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
use_default_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2310102012101212-1313110121233032-0222012221233332-0332003311302302-3001031232200013-1030223210030220-3330302311103220-0202103201323322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_network` properties

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network

<a id="canonical-3022012000222210-1003230222202323-3003020331002222-2220021330120300-0113301010130111-0100211103333212-0303320203012333-1121320112020023"></a>

Type: `"object"`. single nested block, Optional.

Parameters to advertise on a given virtual network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_v6_vip",
    "specific_v6_vip"),
  validators.ConflictingObjectAttributes("default_vip",
    "specific_vip")}
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
  "x-ves-oneof-field-v6_vip_choice": "[\"default_v6_vip\",\"specific_v6_vip\"]",
  "x-ves-oneof-field-vip_choice": "[\"default_vip\",\"specific_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_network {
  # Configure direct properties listed below.
}
```
