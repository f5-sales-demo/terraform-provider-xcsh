---
page_title: "xcsh_tunnel reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tunnel reference."
---

# xcsh_tunnel reference

<a id="canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- Property reference

<a id="canonical-3310000213132222-2130313331233302-2332003020130112-1211111103101012-2310123323221023-3200312032213133-1110333333203202-3002032011301202"></a>

### Direct properties for `xcsh_tunnel`

<a id="canonical-2232233100322033-0212020030100032-3101113111300001-2002200202110130-1122310130021000-2022112233130001-2333323220312300-3322001032020302"></a>

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

<a id="canonical-0031212212233122-1313301130112231-3231002233320023-2331220133223102-1100333121002201-2222311113020223-3321032023223312-3200122100110310"></a>

<a id="canonical-2130021031013202-1110233012312130-2231020202132003-2300000022301101-0332102211033101-0232301221113221-0300301033021030-3230232311333020"></a>

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

<a id="canonical-2030311013130230-2312102210113201-3000333002300312-2211123203000123-3121323330021320-0300221122300232-1022200013120030-2021011321230023"></a>

<a id="canonical-1012223210020002-3230031011103101-1003310011111311-2032022033203202-2020011001331321-1002320220232030-3323223103122333-3223223113012302"></a>

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

<a id="canonical-3323001302220001-2221023313201323-2003001310201332-3013030022031030-0232210130002200-1223030323321212-1232323302233112-2000013311113130"></a>

<a id="canonical-3323310230120102-2303212012032203-2013020332123001-2023101133213101-3003002301011312-3013211032023321-2322122002120021-1200020211022233"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0232003112130033-3221032000301123-1121031100200100-0003131321012011-3300232113013230-0130220133131121-3031212123120222-2110301031103013"></a>

<a id="canonical-2112221122230222-1010002212322213-2213213110101333-2020313303302112-2221031012023201-1301323311113122-1012330033310001-3221232200013232"></a>

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

- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013): complete subsection reference.

<a id="canonical-3100301333132322-3000311000003213-3113220321000300-0210011021203312-3112332130001011-0200321301111220-2123013111012203-1110230212111013"></a>

<a id="canonical-0003012000100133-3320203321313211-3302121213221023-1321101002213321-1021323130013000-2013010102130330-3002012100211022-1313100330031010"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Tunnel. Must be unique within the namespace.

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

<a id="canonical-1321231001022330-3113012102100013-0203103102012113-0231312231023223-2220031330202320-0231232320310133-1010011332030031-0110122111202013"></a>

<a id="canonical-2032023123233302-0002200231330101-1211120211320101-3322110102100301-0102001001121101-3022211331001223-2210030233031303-1332222210110200"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Tunnel is created.

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

- [params](resources--tunnel--reference--group-001.md#canonical-3112221221213033-1303222023230230-0222322021220323-2332312322123003-2200133123232101-3021120033320130-3200322033133310-2131310310120213): complete subsection reference.

- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0020131231323103-2233031133231022-3221122302102332-0010301321220022-1200003103313200-0000212120201011-3323033323302300-1213020310102310): complete subsection reference.

- [timeouts](resources--tunnel--reference--group-001.md#canonical-0232112222121303-1002201122301213-3101023123123300-0132011300202221-1333302123121302-2332110031022232-0013211020002311-1323011300230131): complete subsection reference.

<a id="canonical-0022333131222112-3332000032211323-0122030103222311-3231032213131221-0231003212002002-2201112233112202-1321013203133101-0213332032002323"></a>

<a id="canonical-2032233232102011-0010100332231120-1302320102203021-0130121300333201-3332030303302201-3211012203201221-0301322010000133-2101122011113320"></a>

#### `tunnel_type` property

Type: `"string"`. Optional, Computed.

\[Enum: IPSEC\_PSK|GRE\] Supported tunnel types are IPsec IPsec tunnel type with PSK GRE tunnel
type. Possible values are \`IPSEC\_PSK\`, \`GRE\`. Defaults to \`IPSEC\_PSK\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["GRE","IPSEC_PSK"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("IPSEC_PSK",
    "GRE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "IPSEC_PSK",
  "enum": [
    "IPSEC_PSK",
    "GRE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0021320110100131-2011003331303011-1230310320231231-3032012113232203-0301110130103002-0210300020313300-1010020211330013-1332322230313311"></a>

### All schema paths for `xcsh_tunnel`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--tunnel--reference--group-001.md#canonical-2232233100322033-0212020030100032-3101113111300001-2002200202110130-1122310130021000-2022112233130001-2333323220312300-3322001032020302) |
| `description` | [description](resources--tunnel--reference--group-001.md#canonical-0031212212233122-1313301130112231-3231002233320023-2331220133223102-1100333121002201-2222311113020223-3321032023223312-3200122100110310) |
| `disable` | [disable](resources--tunnel--reference--group-001.md#canonical-2030311013130230-2312102210113201-3000333002300312-2211123203000123-3121323330021320-0300221122300232-1022200013120030-2021011321230023) |
| `id` | [ID](resources--tunnel--reference--group-001.md#canonical-3323001302220001-2221023313201323-2003001310201332-3013030022031030-0232210130002200-1223030323321212-1232323302233112-2000013311113130) |
| `labels` | [labels](resources--tunnel--reference--group-001.md#canonical-0232003112130033-3221032000301123-1121031100200100-0003131321012011-3300232113013230-0130220133131121-3031212123120222-2110301031103013) |
| `local_ip` | [local_ip](resources--tunnel--reference--group-001.md#canonical-1202220013103103-1322102031303020-3032110113301101-1010010031002202-1011313311020312-1221131133212020-2212312012111111-2232002222210112) |
| `local_ip.intf` | [local_ip.intf](resources--tunnel--reference--group-001.md#canonical-2232000333210012-2012120000123302-3220322313001122-2023131121232202-3330223221012212-3213220301231321-1321312002230010-0010130023130121) |
| `local_ip.intf.local_intf` | [local_ip.intf.local_intf](resources--tunnel--reference--group-001.md#canonical-0333303222203010-2121132211233312-3210300220100322-0013130112203000-1002213230210322-1020313022322021-2310130323203311-3222312322313011) |
| `local_ip.intf.local_intf.kind` | [local_ip.intf.local_intf.kind](resources--tunnel--reference--group-001.md#canonical-0103120020311111-3332202102323132-0211003112330020-3023122000011311-1300132030302210-3330021210131311-3322113130321122-2110313313221130) |
| `local_ip.intf.local_intf.name` | [local_ip.intf.local_intf.name](resources--tunnel--reference--group-001.md#canonical-2100033131303000-2323220003110333-3112202320231203-3211300121000112-3312212310203301-1202220020232223-2030202210132030-1101121133232311) |
| `local_ip.intf.local_intf.namespace` | [local_ip.intf.local_intf.namespace](resources--tunnel--reference--group-001.md#canonical-0102232002031021-3330301231000313-3300020100033011-1203200330303110-0131133210012112-2100200203131110-3033313020221330-1311223012101332) |
| `local_ip.intf.local_intf.tenant` | [local_ip.intf.local_intf.tenant](resources--tunnel--reference--group-001.md#canonical-0013022201123012-0332100211203003-2002002103213101-0203203311021031-0312223330213002-2121223201330032-3312010232110120-3222320132223131) |
| `local_ip.intf.local_intf.uid` | [local_ip.intf.local_intf.uid](resources--tunnel--reference--group-001.md#canonical-1112121221102313-0132330111332330-3233012120022311-1332233001122023-3300033022212200-1100330322120310-0012202131201311-0021131031012333) |
| `local_ip.ip_address` | [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-3211322333131020-1231123311323113-2311200003110211-2013102323123313-2010312330020320-2021312103122033-3120022131001133-1330213331003310) |
| `local_ip.ip_address.auto` | [local_ip.ip_address.auto](resources--tunnel--reference--group-001.md#canonical-0313203200203132-0211111113102102-0310212011101200-2213312303213121-2032323132101033-0033122231133213-2121332223320001-0131321211211232) |
| `local_ip.ip_address.ip_address` | [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-1102201013213133-2302103202012100-0232211120113031-2021113132332132-1030333100130020-2222202302300011-1231311210110030-0133303100330011) |
| `local_ip.ip_address.ip_address.dual_stack` | [local_ip.ip_address.ip_address.dual_stack](resources--tunnel--reference--group-001.md#canonical-1322222010010012-3211113113002130-3230123333031302-0203300230230301-0233121322033000-0230132112021333-1003210300112203-1113303310032131) |
| `local_ip.ip_address.ip_address.dual_stack.ipv4` | [local_ip.ip_address.ip_address.dual_stack.ipv4](resources--tunnel--reference--group-001.md#canonical-3332220223221100-2210201300001223-2030102102122303-3020000232103333-3221011123133011-0123011333232322-3013132222320021-1002132201202131) |
| `local_ip.ip_address.ip_address.dual_stack.ipv4.addr` | [local_ip.ip_address.ip_address.dual_stack.ipv4.addr](resources--tunnel--reference--group-001.md#canonical-1013133133121202-0013032103013323-0223121210130121-0002013211021310-1322232102121013-3011120000111103-0221101012130231-3232002113232021) |
| `local_ip.ip_address.ip_address.dual_stack.ipv6` | [local_ip.ip_address.ip_address.dual_stack.ipv6](resources--tunnel--reference--group-001.md#canonical-3003120023102003-3023030233023111-3133030030302322-3232202100202030-1121212312032023-1233100300102300-3032003023313232-0023032221201121) |
| `local_ip.ip_address.ip_address.dual_stack.ipv6.addr` | [local_ip.ip_address.ip_address.dual_stack.ipv6.addr](resources--tunnel--reference--group-001.md#canonical-3112330320103301-1121021303123100-3130222003000231-0201000330033100-0321122131103123-2222210002002023-0032102322023032-0230021302220320) |
| `local_ip.ip_address.ip_address.ipv4` | [local_ip.ip_address.ip_address.ipv4](resources--tunnel--reference--group-001.md#canonical-3010221100010132-1023212101322033-3021203233131020-3033100212101020-1111321013333302-2320323132203123-2130100321321113-2301121230321333) |
| `local_ip.ip_address.ip_address.ipv4.addr` | [local_ip.ip_address.ip_address.ipv4.addr](resources--tunnel--reference--group-001.md#canonical-2110200122110023-3332032333120233-0130102001323122-1202113232121020-3211300121033232-1211321121323131-0120033013002233-0021003130313223) |
| `local_ip.ip_address.ip_address.ipv6` | [local_ip.ip_address.ip_address.ipv6](resources--tunnel--reference--group-001.md#canonical-3210322211210301-3211202321102010-1022310333100230-2101022222002220-2303320311201221-3021313202033002-0131003310332323-2023200203301230) |
| `local_ip.ip_address.ip_address.ipv6.addr` | [local_ip.ip_address.ip_address.ipv6.addr](resources--tunnel--reference--group-001.md#canonical-0010320321331030-0111301010212032-1100323313310323-3030133003030223-2321203330001130-3202032130132023-1101113132100220-2301313333333201) |
| `local_ip.ip_address.virtual_network_type` | [local_ip.ip_address.virtual_network_type](resources--tunnel--reference--group-001.md#canonical-1103101120333011-2022122103301321-1023101301333110-1333220231132122-3022100122013312-0310031032231320-0110031111033033-3101102330210113) |
| `local_ip.ip_address.virtual_network_type.public` | [local_ip.ip_address.virtual_network_type.public](resources--tunnel--reference--group-001.md#canonical-2203202332021331-3202030011232211-2330100312022323-0303322102102121-3311332223221210-3212302331232212-2221320221111023-2123123013102332) |
| `local_ip.ip_address.virtual_network_type.site_local` | [local_ip.ip_address.virtual_network_type.site_local](resources--tunnel--reference--group-001.md#canonical-0212131023100201-2010002102301200-2312322132300222-3232212133111020-3332010011312130-0213302201032113-0111001013003030-0100210332200123) |
| `local_ip.ip_address.virtual_network_type.site_local_inside` | [local_ip.ip_address.virtual_network_type.site_local_inside](resources--tunnel--reference--group-001.md#canonical-3330031120222220-1220333002323010-0133030321021302-3130300130101120-2011032003321131-0131000232031333-1001211111202331-1313330211132331) |
| `name` | [name](resources--tunnel--reference--group-001.md#canonical-3100301333132322-3000311000003213-3113220321000300-0210011021203312-3112332130001011-0200321301111220-2123013111012203-1110230212111013) |
| `namespace` | [namespace](resources--tunnel--reference--group-001.md#canonical-1321231001022330-3113012102100013-0203103102012113-0231312231023223-2220031330202320-0231232320310133-1010011332030031-0110122111202013) |
| `params` | [params](resources--tunnel--reference--group-001.md#canonical-1300020203312121-2031220322300310-3322012003033220-3122131101023202-2210103221213330-1112132322220111-3323231020320010-2232030300311203) |
| `params.ipsec` | [params.ipsec](resources--tunnel--reference--group-001.md#canonical-2202212330001330-1311313333333331-1331010312300322-1331131313213033-0010030122031001-1032022123220300-0111300212331210-0302123213130332) |
| `params.ipsec.ipsec_psk` | [params.ipsec.ipsec_psk](resources--tunnel--reference--group-001.md#canonical-0211110222012133-2211120103112113-3313223103333302-1001233120310021-3313310313331323-1113122213311211-1103131113220221-2131330220313210) |
| `params.ipsec.ipsec_psk.blindfold_secret_info` | [params.ipsec.ipsec_psk.blindfold_secret_info](resources--tunnel--reference--group-001.md#canonical-3002031102001033-0012222320011132-2132211133300230-3233201120112233-0302123313101120-3332002310313203-2310303333011131-3213120201303201) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.decryption_provider` | [params.ipsec.ipsec_psk.blindfold_secret_info.decryption_provider](resources--tunnel--reference--group-001.md#canonical-2112032132203213-2212212100012120-2211100210023010-3332233133200332-1211111332013220-0302213202323111-2102301301133110-2301321113210002) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.location` | [params.ipsec.ipsec_psk.blindfold_secret_info.location](resources--tunnel--reference--group-001.md#canonical-0203101102100221-3011013031303002-3132001132122202-2210201301132121-0320330202021232-0331222013020023-1013012203101133-3101013332013020) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.store_provider` | [params.ipsec.ipsec_psk.blindfold_secret_info.store_provider](resources--tunnel--reference--group-001.md#canonical-3003103001302130-0112010122111002-2131212311120132-0013321013001113-3231313210113230-3011123212021003-2112122220301212-1133011132113012) |
| `params.ipsec.ipsec_psk.clear_secret_info` | [params.ipsec.ipsec_psk.clear_secret_info](resources--tunnel--reference--group-001.md#canonical-3130332012201312-1121331233232102-2031310101003331-2213111112032210-3133003113311131-2011330320130100-0123213003011121-1301211302033203) |
| `params.ipsec.ipsec_psk.clear_secret_info.provider_ref` | [params.ipsec.ipsec_psk.clear_secret_info.provider_ref](resources--tunnel--reference--group-001.md#canonical-1103131120220003-2131121132300232-1322223123120021-3301233123221133-0321300133123300-0102020212210030-0203302330133122-1211000302031202) |
| `params.ipsec.ipsec_psk.clear_secret_info.url` | [params.ipsec.ipsec_psk.clear_secret_info.url](resources--tunnel--reference--group-001.md#canonical-0000331100010220-1232122120201301-2000310023002102-2223122101110000-1013200010322322-2120212332211330-1032320233222102-0021132300030021) |
| `remote_ip` | [remote_ip](resources--tunnel--reference--group-001.md#canonical-1102101210301330-3303010122231300-0323211303100031-3030003231221213-2022001201231223-0100202312303102-2121222003103132-0231103022201212) |
| `remote_ip.endpoints` | [remote_ip.endpoints](resources--tunnel--reference--group-001.md#canonical-1113313210331231-3013030110330231-2201323001211123-0131131330301123-0101330003112001-3102312033002133-1023222331212323-0223132112220300) |
| `remote_ip.endpoints.endpoints` | [remote_ip.endpoints.endpoints](resources--tunnel--reference--group-001.md#canonical-0220330101100332-0310020221231123-2032012021202333-0023120202113232-0111203201022212-0322323223102232-3233113310202201-1133010023212122) |
| `remote_ip.ip` | [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-3222222110011123-0331202323320211-2223210223211200-2202331122221003-1023313130131322-0210202011222322-0322333320230001-0331003011131210) |
| `remote_ip.ip.dual_stack` | [remote_ip.ip.dual_stack](resources--tunnel--reference--group-001.md#canonical-3301103133321333-3203303030321100-2100011331313212-0022113201103231-3023122213233232-3111311220122121-1321022311131112-3133202311122002) |
| `remote_ip.ip.dual_stack.ipv4` | [remote_ip.ip.dual_stack.ipv4](resources--tunnel--reference--group-001.md#canonical-0221132310021102-0233232133320121-0310003311330123-1220113111002031-0022010110013302-1032220200000330-2131032332201322-2213122101003111) |
| `remote_ip.ip.dual_stack.ipv4.addr` | [remote_ip.ip.dual_stack.ipv4.addr](resources--tunnel--reference--group-001.md#canonical-0311302312122133-0002213132023103-2131302321330221-0013023123230113-3130130131232300-1123010032230230-3310323033331111-1132131001213132) |
| `remote_ip.ip.dual_stack.ipv6` | [remote_ip.ip.dual_stack.ipv6](resources--tunnel--reference--group-001.md#canonical-0230323232332330-3313011331321313-1131303222030203-2333123122011302-1120322012300231-0110200003030300-1030000313122032-2220321321301303) |
| `remote_ip.ip.dual_stack.ipv6.addr` | [remote_ip.ip.dual_stack.ipv6.addr](resources--tunnel--reference--group-001.md#canonical-3010233323011222-2301120013332331-0130031320001113-0120300230113103-3030322302022110-3331231333101110-3111200022222131-0111330012033031) |
| `remote_ip.ip.ipv4` | [remote_ip.ip.ipv4](resources--tunnel--reference--group-001.md#canonical-2022223321220223-0113212330330230-2010333000223313-1123131223032312-1303112311000110-1233100302222332-1220123000312013-2330330100302032) |
| `remote_ip.ip.ipv4.addr` | [remote_ip.ip.ipv4.addr](resources--tunnel--reference--group-001.md#canonical-2302303131303320-0302300030033031-1031013012200222-3031311121003221-1130331032002020-2003231103111230-3321023212222102-1310223102122103) |
| `remote_ip.ip.ipv6` | [remote_ip.ip.ipv6](resources--tunnel--reference--group-001.md#canonical-1312133122302313-3023110001322232-0211222001230002-3101203231021111-3110020211220323-0213102221230102-3212311233033033-0011332000002313) |
| `remote_ip.ip.ipv6.addr` | [remote_ip.ip.ipv6.addr](resources--tunnel--reference--group-001.md#canonical-1020132313120222-1203003101131202-2300113310110322-3021230202010110-0321211132012132-3222202211213101-0023122312123313-1122323012130331) |
| `timeouts` | [timeouts](resources--tunnel--reference--group-001.md#canonical-2203011112121111-2102101030131031-0232212122220302-1033312300213033-1022020033022311-0111023133000222-2213300203301030-0201312121232000) |
| `timeouts.create` | [timeouts.create](resources--tunnel--reference--group-001.md#canonical-1111032002331233-3011110202013331-2100302120201323-3322333303332330-0013332110113221-2311222011233100-2332021123001022-3022310010012200) |
| `timeouts.delete` | [timeouts.delete](resources--tunnel--reference--group-001.md#canonical-3003300231121032-1010333210303020-1300230311103133-1320033333203300-2222310200110222-2021201232110223-2132220131131011-1311323313300213) |
| `timeouts.read` | [timeouts.read](resources--tunnel--reference--group-001.md#canonical-3122001321021033-3321201131213102-1313303010202332-1032320021110000-1030201021012030-2030222211203031-2201113312201233-1311213103303320) |
| `timeouts.update` | [timeouts.update](resources--tunnel--reference--group-001.md#canonical-1030121211200000-1013330003323000-1101101100330112-2303110112000031-0122013132231223-3313212131313223-0012333212033201-0122321202311103) |
| `tunnel_type` | [tunnel_type](resources--tunnel--reference--group-001.md#canonical-0022333131222112-3332000032211323-0122030103222311-3231032213131221-0231003212002002-2201112233112202-1321013203133101-0213332032002323) |

<a id="canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- local_ip

<a id="canonical-1202220013103103-1322102031303020-3032110113301101-1010010031002202-1011313311020312-1221131133212020-2212312012111111-2232002222210112"></a>

Type: `"object"`. single nested block, Optional.

Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS
available are - 1. Local Interface - Network Interface from which IP address and network will be
selected 2. IP Address - IP address and network can be configured explicitly.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("intf",
    "ip_address")}
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
  "x-ves-oneof-field-type": "[\"intf\",\"ip_address\"]"
}
```

Terraform syntax:

```terraform
local_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122120302133223-2210101321212021-1020221231323120-1133010232002031-0110021330312213-3030122312313032-1130000031331311-0133031002031002"></a>

### Direct properties for `local_ip`

- [intf](resources--tunnel--reference--group-001.md#canonical-2223230302300101-0122121121132103-2321020002221033-3320003233333123-3313223101333111-2000011230222230-0300113212001320-1033032031000213): complete subsection reference.

- [ip_address](resources--tunnel--reference--group-001.md#canonical-0101220131022022-1000210322332102-0021030332331230-1223222200000203-3112120121021221-1333110220021312-3320303000313110-3032130312212310): complete subsection reference.

<a id="canonical-2223230302300101-0122121121132103-2321020002221033-3320003233333123-3313223101333111-2000011230222230-0300113212001320-1033032031000213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.intf` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013)
- local_ip.intf

<a id="canonical-2232000333210012-2012120000123302-3220322313001122-2023131121232202-3330223221012212-3213220301231321-1321312002230010-0010130023130121"></a>

Type: `"object"`. single nested block, Optional.

Provides the local interface to pick up source IP and network for transporting encapsulated packet.

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
intf {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110130211011232-0220330330022120-1112131311130302-2230201100222030-3100301121311303-0131133211102112-1201013131303312-1322231123130001"></a>

### Direct properties for `local_ip.intf`

- [local_intf](resources--tunnel--reference--group-001.md#canonical-0033232322122220-1001102021210033-3113312202103221-1132231210330230-3031302022112310-0322232330223010-0001022110123012-3232111312130330): complete subsection reference.

<a id="canonical-0033232322122220-1001102021210033-3113312202103221-1132231210330230-3031302022112310-0322232330223010-0001022110123012-3232111312130330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.intf.local_intf` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013)
- [local_ip.intf](resources--tunnel--reference--group-001.md#canonical-2223230302300101-0122121121132103-2321020002221033-3320003233333123-3313223101333111-2000011230222230-0300113212001320-1033032031000213)
- local_ip.intf.local_intf

<a id="canonical-0333303222203010-2121132211233312-3210300220100322-0013130112203000-1002213230210322-1020313022322021-2310130323203311-3222312322313011"></a>

Type: `"object"`. list nested block, Optional.

Local interface to be used for filling in source information of IP and network for transport.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
local_intf {
  # Configure direct properties listed below.
}
```

<a id="canonical-0202320110233311-3313021302333233-1031311130201002-3110102031131121-0323303233102203-3113311130121302-1303311013311202-2310333023103102"></a>

### Direct properties for `local_ip.intf.local_intf`

<a id="canonical-0103120020311111-3332202102323132-0211003112330020-3023122000011311-1300132030302210-3330021210131311-3322113130321122-2110313313221130"></a>

#### `local_ip.intf.local_intf.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-2100033131303000-2323220003110333-3112202320231203-3211300121000112-3312212310203301-1202220020232223-2030202210132030-1101121133232311"></a>

<a id="canonical-3012332113322220-2121211200313332-1003212030333111-3010220202233231-1302202213132302-1201103220232303-0320310101222012-3010100032101222"></a>

#### `local_ip.intf.local_intf.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-0102232002031021-3330301231000313-3300020100033011-1203200330303110-0131133210012112-2100200203131110-3033313020221330-1311223012101332"></a>

<a id="canonical-0220220211223232-0111312033303333-1102213130020303-0220110130102021-0323003231112311-3303221001203023-0311201123301223-2210300202013332"></a>

#### `local_ip.intf.local_intf.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0013022201123012-0332100211203003-2002002103213101-0203203311021031-0312223330213002-2121223201330032-3312010232110120-3222320132223131"></a>

<a id="canonical-1113303032033313-3333133031030231-2223202120201220-0100001120012320-2310221300301010-0013122101132131-0331102102220332-2232232223311230"></a>

#### `local_ip.intf.local_intf.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-1112121221102313-0132330111332330-3233012120022311-1332233001122023-3300033022212200-1100330322120310-0012202131201311-0021131031012333"></a>

<a id="canonical-0021323020030131-0323233130023030-0013201021011000-3000113330330101-3331301221112012-3100020021203301-1222110003110212-1112310312002310"></a>

#### `local_ip.intf.local_intf.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-0101220131022022-1000210322332102-0021030332331230-1223222200000203-3112120121021221-1333110220021312-3320303000313110-3032130312212310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013)
- local_ip.ip_address

<a id="canonical-3211322333131020-1231123311323113-2311200003110211-2013102323123313-2010312330020320-2021312103122033-3120022131001133-1330213331003310"></a>

Type: `"object"`. single nested block, Optional.

Provides the configuration to pick up source IP and network for transporting encapsulated packet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto",
    "ip_address")}
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
  "x-ves-oneof-field-type": "[\"auto\",\"ip_address\"]"
}
```

Terraform syntax:

```terraform
ip_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-2201022100223333-1231133101110211-1022031302332010-3303102231323332-2132133223210133-1210123220300003-1101301320223203-0323010222231322"></a>

### Direct properties for `local_ip.ip_address`

- [auto](resources--tunnel--reference--group-001.md#canonical-2310113003120102-2123032023112022-1033021130133210-3202032303213133-0320330211322033-0121101031222013-0020231300333130-1032320033221121): complete subsection reference.

- [ip_address](resources--tunnel--reference--group-001.md#canonical-0123320310030102-3000002332100213-3023203210330130-0211220103322203-2133231313021313-2320203201313302-0230120313001220-0212201223230022): complete subsection reference.

- [virtual_network_type](resources--tunnel--reference--group-001.md#canonical-0101202023032121-3321102122230302-3102130011232222-1203321222212021-3013332320102113-3332333233110211-3213321132100201-3333213310222103): complete subsection reference.

<a id="canonical-2310113003120102-2123032023112022-1033021130133210-3202032303213133-0320330211322033-0121101031222013-0020231300333130-1032320033221121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.auto` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-0101220131022022-1000210322332102-0021030332331230-1223222200000203-3112120121021221-1333110220021312-3320303000313110-3032130312212310)
- local_ip.ip_address.auto

<a id="canonical-0313203200203132-0211111113102102-0310212011101200-2213312303213121-2032323132101033-0033122231133213-2121332223320001-0131321211211232"></a>

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
auto = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123320310030102-3000002332100213-3023203210330130-0211220103322203-2133231313021313-2320203201313302-0230120313001220-0212201223230022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.ip_address` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-0101220131022022-1000210322332102-0021030332331230-1223222200000203-3112120121021221-1333110220021312-3320303000313110-3032130312212310)
- local_ip.ip_address.ip_address

<a id="canonical-1102201013213133-2302103202012100-0232211120113031-2021113132332132-1030333100130020-2222202302300011-1231311210110030-0133303100330011"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
ip_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112101232110212-2003130301332102-3212003233101202-3000030123223330-1231231230203203-2123331120233301-3233213002313232-3212132212222321"></a>

### Direct properties for `local_ip.ip_address.ip_address`

- [dual_stack](resources--tunnel--reference--group-001.md#canonical-3013212302112132-1030133231022030-3222012011111002-1112033112113322-3030322233312100-1312312201211031-1020310211101203-1321022203113030): complete subsection reference.

- [IPv4](resources--tunnel--reference--group-001.md#canonical-3002322102010000-2000122100201021-1013220022011213-1103122211213231-3221313033132330-3201012030112122-0112000032131332-2030102123001222): complete subsection reference.

- [IPv6](resources--tunnel--reference--group-001.md#canonical-2210203103113323-2132313122312023-1211333201012213-0332303201320120-3203021021002223-0131010200221011-2320103330121231-0320020103111203): complete subsection reference.

<a id="canonical-3013212302112132-1030133231022030-3222012011111002-1112033112113322-3030322233312100-1312312201211031-1020310211101203-1321022203113030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.ip_address.dual_stack` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-0101220131022022-1000210322332102-0021030332331230-1223222200000203-3112120121021221-1333110220021312-3320303000313110-3032130312212310)
- [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-0123320310030102-3000002332100213-3023203210330130-0211220103322203-2133231313021313-2320203201313302-0230120313001220-0212201223230022)
- local_ip.ip_address.ip_address.dual_stack

<a id="canonical-1322222010010012-3211113113002130-3230123333031302-0203300230230301-0233121322033000-0230132112021333-1003210300112203-1113303310032131"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

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
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120131130320012-2320323213012111-3300002233233120-3012003003210103-2022302100031012-3232333231323013-1022313200303002-0030123010021020"></a>

### Direct properties for `local_ip.ip_address.ip_address.dual_stack`

- [IPv4](resources--tunnel--reference--group-001.md#canonical-1213031103013210-3121000233320202-1021232303012301-1233031111231013-3122111013213111-2130212100312002-3233103103202011-3001012330132312): complete subsection reference.

- [IPv6](resources--tunnel--reference--group-001.md#canonical-0333000332031203-2331130013010313-0212302020100120-1103033213303323-0333100123233002-2103010131322101-1312003022110323-3222330331022303): complete subsection reference.

<a id="canonical-1213031103013210-3121000233320202-1021232303012301-1233031111231013-3122111013213111-2130212100312002-3233103103202011-3001012330132312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.ip_address.dual_stack.ipv4` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-0101220131022022-1000210322332102-0021030332331230-1223222200000203-3112120121021221-1333110220021312-3320303000313110-3032130312212310)
- [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-0123320310030102-3000002332100213-3023203210330130-0211220103322203-2133231313021313-2320203201313302-0230120313001220-0212201223230022)
- [local_ip.ip_address.ip_address.dual_stack](resources--tunnel--reference--group-001.md#canonical-3013212302112132-1030133231022030-3222012011111002-1112033112113322-3030322233312100-1312312201211031-1020310211101203-1321022203113030)
- local_ip.ip_address.ip_address.dual_stack.IPv4

<a id="canonical-3332220223221100-2210201300001223-2030102102122303-3020000232103333-3221011123133011-0123011333232322-3013132222320021-1002132201202131"></a>

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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321313233200111-2022213133330330-0213133022023102-2323312021120110-0022001103220201-1302210132210223-2211303111311201-2230001230002210"></a>

### Direct properties for `local_ip.ip_address.ip_address.dual_stack.ipv4`

<a id="canonical-1013133133121202-0013032103013323-0223121210130121-0002013211021310-1322232102121013-3011120000111103-0221101012130231-3232002113232021"></a>

#### `local_ip.ip_address.ip_address.dual_stack.ipv4.addr` property

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0333000332031203-2331130013010313-0212302020100120-1103033213303323-0333100123233002-2103010131322101-1312003022110323-3222330331022303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.ip_address.dual_stack.ipv6` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-0101220131022022-1000210322332102-0021030332331230-1223222200000203-3112120121021221-1333110220021312-3320303000313110-3032130312212310)
- [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-0123320310030102-3000002332100213-3023203210330130-0211220103322203-2133231313021313-2320203201313302-0230120313001220-0212201223230022)
- [local_ip.ip_address.ip_address.dual_stack](resources--tunnel--reference--group-001.md#canonical-3013212302112132-1030133231022030-3222012011111002-1112033112113322-3030322233312100-1312312201211031-1020310211101203-1321022203113030)
- local_ip.ip_address.ip_address.dual_stack.IPv6

<a id="canonical-3003120023102003-3023030233023111-3133030030302322-3232202100202030-1121212312032023-1233100300102300-3032003023313232-0023032221201121"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302210123102233-1213030022112330-3201033023010220-0122331032000132-2333022210201210-2030012030230122-0033011030021323-3230002200322310"></a>

### Direct properties for `local_ip.ip_address.ip_address.dual_stack.ipv6`

<a id="canonical-3112330320103301-1121021303123100-3130222003000231-0201000330033100-0321122131103123-2222210002002023-0032102322023032-0230021302220320"></a>

#### `local_ip.ip_address.ip_address.dual_stack.ipv6.addr` property

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3002322102010000-2000122100201021-1013220022011213-1103122211213231-3221313033132330-3201012030112122-0112000032131332-2030102123001222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.ip_address.ipv4` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-0101220131022022-1000210322332102-0021030332331230-1223222200000203-3112120121021221-1333110220021312-3320303000313110-3032130312212310)
- [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-0123320310030102-3000002332100213-3023203210330130-0211220103322203-2133231313021313-2320203201313302-0230120313001220-0212201223230022)
- local_ip.ip_address.ip_address.IPv4

<a id="canonical-3010221100010132-1023212101322033-3021203233131020-3033100212101020-1111321013333302-2320323132203123-2130100321321113-2301121230321333"></a>

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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302023022102323-0302300103110122-2110031203200333-1312321132203131-2322122131101303-2123221311311220-3332332313122101-2010322222010121"></a>

### Direct properties for `local_ip.ip_address.ip_address.ipv4`

<a id="canonical-2110200122110023-3332032333120233-0130102001323122-1202113232121020-3211300121033232-1211321121323131-0120033013002233-0021003130313223"></a>

#### `local_ip.ip_address.ip_address.ipv4.addr` property

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2210203103113323-2132313122312023-1211333201012213-0332303201320120-3203021021002223-0131010200221011-2320103330121231-0320020103111203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.ip_address.ipv6` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-0101220131022022-1000210322332102-0021030332331230-1223222200000203-3112120121021221-1333110220021312-3320303000313110-3032130312212310)
- [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-0123320310030102-3000002332100213-3023203210330130-0211220103322203-2133231313021313-2320203201313302-0230120313001220-0212201223230022)
- local_ip.ip_address.ip_address.IPv6

<a id="canonical-3210322211210301-3211202321102010-1022310333100230-2101022222002220-2303320311201221-3021313202033002-0131003310332323-2023200203301230"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322322100330111-1213311233321200-1010031001230020-3101313212101101-0121133211133112-2211003330021120-1101032021012312-0001321302213133"></a>

### Direct properties for `local_ip.ip_address.ip_address.ipv6`

<a id="canonical-0010320321331030-0111301010212032-1100323313310323-3030133003030223-2321203330001130-3202032130132023-1101113132100220-2301313333333201"></a>

#### `local_ip.ip_address.ip_address.ipv6.addr` property

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0101202023032121-3321102122230302-3102130011232222-1203321222212021-3013332320102113-3332333233110211-3213321132100201-3333213310222103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.virtual_network_type` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-0101220131022022-1000210322332102-0021030332331230-1223222200000203-3112120121021221-1333110220021312-3320303000313110-3032130312212310)
- local_ip.ip_address.virtual_network_type

<a id="canonical-1103101120333011-2022122103301321-1023101301333110-1333220231132122-3022100122013312-0310031032231320-0110031111033033-3101102330210113"></a>

Type: `"object"`. single nested block, Optional.

Different types of virtual networks understood by the system.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("public",
    "site_local"),
  validators.ConflictingObjectAttributes("public",
    "site_local_inside"),
  validators.ConflictingObjectAttributes("site_local",
    "site_local_inside")}
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
  "x-ves-oneof-field-vn_type_choice": "[\"public\",\"site_local\",\"site_local_inside\"]"
}
```

Terraform syntax:

```terraform
virtual_network_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330323103021010-0302130012221202-1003032312010211-2013102030032010-1033222303001213-2333121102202201-1133133230313300-0010012312321001"></a>

### Direct properties for `local_ip.ip_address.virtual_network_type`

- [public](resources--tunnel--reference--group-001.md#canonical-3120121023033100-3222230032120221-1021310320322210-0312030123110301-1011120112130130-2012031022232022-0113211323030110-3212132332233233): complete subsection reference.

- [site_local](resources--tunnel--reference--group-001.md#canonical-3311200211120100-2120200220132030-0303323022032102-2210030333200002-3300210100312310-0223102133023131-3220110222300230-1201202100310110): complete subsection reference.

- [site_local_inside](resources--tunnel--reference--group-001.md#canonical-1313331030331221-2220333332103303-3233113300321011-0323323233330111-3132323200121322-2211010011030012-3333331223223021-1012032222120121): complete subsection reference.

<a id="canonical-3120121023033100-3222230032120221-1021310320322210-0312030123110301-1011120112130130-2012031022232022-0113211323030110-3212132332233233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.virtual_network_type.public` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-0101220131022022-1000210322332102-0021030332331230-1223222200000203-3112120121021221-1333110220021312-3320303000313110-3032130312212310)
- [local_ip.ip_address.virtual_network_type](resources--tunnel--reference--group-001.md#canonical-0101202023032121-3321102122230302-3102130011232222-1203321222212021-3013332320102113-3332333233110211-3213321132100201-3333213310222103)
- local_ip.ip_address.virtual_network_type.public

<a id="canonical-2203202332021331-3202030011232211-2330100312022323-0303322102102121-3311332223221210-3212302331232212-2221320221111023-2123123013102332"></a>

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
public = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311200211120100-2120200220132030-0303323022032102-2210030333200002-3300210100312310-0223102133023131-3220110222300230-1201202100310110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.virtual_network_type.site_local` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-0101220131022022-1000210322332102-0021030332331230-1223222200000203-3112120121021221-1333110220021312-3320303000313110-3032130312212310)
- [local_ip.ip_address.virtual_network_type](resources--tunnel--reference--group-001.md#canonical-0101202023032121-3321102122230302-3102130011232222-1203321222212021-3013332320102113-3332333233110211-3213321132100201-3333213310222103)
- local_ip.ip_address.virtual_network_type.site_local

<a id="canonical-0212131023100201-2010002102301200-2312322132300222-3232212133111020-3332010011312130-0213302201032113-0111001013003030-0100210332200123"></a>

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
site_local = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313331030331221-2220333332103303-3233113300321011-0323323233330111-3132323200121322-2211010011030012-3333331223223021-1012032222120121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.virtual_network_type.site_local_inside` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-0101220131022022-1000210322332102-0021030332331230-1223222200000203-3112120121021221-1333110220021312-3320303000313110-3032130312212310)
- [local_ip.ip_address.virtual_network_type](resources--tunnel--reference--group-001.md#canonical-0101202023032121-3321102122230302-3102130011232222-1203321222212021-3013332320102113-3332333233110211-3213321132100201-3333213310222103)
- local_ip.ip_address.virtual_network_type.site_local_inside

<a id="canonical-3330031120222220-1220333002323010-0133030321021302-3130300130101120-2011032003321131-0131000232031333-1001211111202331-1313330211132331"></a>

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
site_local_inside = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112221221213033-1303222023230230-0222322021220323-2332312322123003-2200133123232101-3021120033320130-3200322033133310-2131310310120213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `params` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- params

<a id="canonical-1300020203312121-2031220322300310-3322012003033220-3122131101023202-2210103221213330-1112132322220111-3323231020320010-2232030300311203"></a>

Type: `"object"`. single nested block, Optional.

Tunnel configuration parameters for supported encapsulation 1. IPsec is supported with PSK for which
PSK can be configured.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"ipsec\"]"
}
```

Terraform syntax:

```terraform
params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233112301031330-1202230100201220-2111220021110003-0300011123301012-3332013311000132-3212023002203001-1332010103012223-1002032201323302"></a>

### Direct properties for `params`

- [ipsec](resources--tunnel--reference--group-001.md#canonical-1333220212012230-1111003122111213-2201303031102303-3212013001210022-0333212002213133-3030003012100230-3000310333301101-1011203123230131): complete subsection reference.

<a id="canonical-1333220212012230-1111003122111213-2201303031102303-3212013001210022-0333212002213133-3030003012100230-3000310333301101-1011203123230131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `params.ipsec` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [params](resources--tunnel--reference--group-001.md#canonical-3112221221213033-1303222023230230-0222322021220323-2332312322123003-2200133123232101-3021120033320130-3200322033133310-2131310310120213)
- params.ipsec

<a id="canonical-2202212330001330-1311313333333331-1331010312300322-1331131313213033-0010030122031001-1032022123220300-0111300212331210-0302123213130332"></a>

Type: `"object"`. single nested block, Optional.

Configuration for IPsec encapsulation are: 1. PSK - pre shared key to be used by IKE.

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
ipsec {
  # Configure direct properties listed below.
}
```

<a id="canonical-0021312100311120-2223132113330212-2130121213120013-2103201102011033-0002202333202010-2320011310302323-0021131012012330-0300113022012313"></a>

### Direct properties for `params.ipsec`

- [ipsec_psk](resources--tunnel--reference--group-001.md#canonical-0120100121011131-0222212233122303-0033310323102030-0111122220023121-0012322203220132-3213012300321002-0300020001321011-2000312230323110): complete subsection reference.

<a id="canonical-0120100121011131-0222212233122303-0033310323102030-0111122220023121-0012322203220132-3213012300321002-0300020001321011-2000312230323110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `params.ipsec.ipsec_psk` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [params](resources--tunnel--reference--group-001.md#canonical-3112221221213033-1303222023230230-0222322021220323-2332312322123003-2200133123232101-3021120033320130-3200322033133310-2131310310120213)
- [params.ipsec](resources--tunnel--reference--group-001.md#canonical-1333220212012230-1111003122111213-2201303031102303-3212013001210022-0333212002213133-3030003012100230-3000310333301101-1011203123230131)
- params.ipsec.ipsec_psk

<a id="canonical-0211110222012133-2211120103112113-3313223103333302-1001233120310021-3313310313331323-1113122213311211-1103131113220221-2131330220313210"></a>

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
ipsec_psk {
  # Configure direct properties listed below.
}
```

<a id="canonical-3012022001313233-2320230202022312-1120300230133303-1111022310121132-0122103312222110-2311000233330302-2131103203323113-2113300023331311"></a>

### Direct properties for `params.ipsec.ipsec_psk`

- [blindfold_secret_info](resources--tunnel--reference--group-001.md#canonical-2111211130232311-3131010002301010-1002113100220030-2030323201332221-3011122003103232-1100313220111032-3220301301303131-2303231102102231): complete subsection reference.

- [clear_secret_info](resources--tunnel--reference--group-001.md#canonical-1331213300122202-3101201130032022-0222212220301031-1030113301222102-1301232003013011-0000113021123220-0013130331220321-3331313100132101): complete subsection reference.

<a id="canonical-2111211130232311-3131010002301010-1002113100220030-2030323201332221-3011122003103232-1100313220111032-3220301301303131-2303231102102231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `params.ipsec.ipsec_psk.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [params](resources--tunnel--reference--group-001.md#canonical-3112221221213033-1303222023230230-0222322021220323-2332312322123003-2200133123232101-3021120033320130-3200322033133310-2131310310120213)
- [params.ipsec](resources--tunnel--reference--group-001.md#canonical-1333220212012230-1111003122111213-2201303031102303-3212013001210022-0333212002213133-3030003012100230-3000310333301101-1011203123230131)
- [params.ipsec.ipsec_psk](resources--tunnel--reference--group-001.md#canonical-0120100121011131-0222212233122303-0033310323102030-0111122220023121-0012322203220132-3213012300321002-0300020001321011-2000312230323110)
- params.ipsec.ipsec_psk.blindfold_secret_info

<a id="canonical-3002031102001033-0012222320011132-2132211133300230-3233201120112233-0302123313101120-3332002310313203-2310303333011131-3213120201303201"></a>

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

<a id="canonical-0030011333301110-3030231122102020-2130200232223013-1023233032222101-2010230221023330-0212231032331033-1303121030103331-1133000301121323"></a>

### Direct properties for `params.ipsec.ipsec_psk.blindfold_secret_info`

<a id="canonical-2112032132203213-2212212100012120-2211100210023010-3332233133200332-1211111332013220-0302213202323111-2102301301133110-2301321113210002"></a>

#### `params.ipsec.ipsec_psk.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0203101102100221-3011013031303002-3132001132122202-2210201301132121-0320330202021232-0331222013020023-1013012203101133-3101013332013020"></a>

<a id="canonical-1333123022031202-3313301322321200-2322032111012132-2021033223021032-3011123113121222-1222202113230010-0301332333331111-1032302331223103"></a>

#### `params.ipsec.ipsec_psk.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3003103001302130-0112010122111002-2131212311120132-0013321013001113-3231313210113230-3011123212021003-2112122220301212-1133011132113012"></a>

<a id="canonical-2223300133331020-1302002222003013-0100321120022313-3220012001313121-1331232232003233-3332332013100002-2030130012120211-0233012320221102"></a>

#### `params.ipsec.ipsec_psk.blindfold_secret_info.store_provider` property

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

<a id="canonical-1331213300122202-3101201130032022-0222212220301031-1030113301222102-1301232003013011-0000113021123220-0013130331220321-3331313100132101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `params.ipsec.ipsec_psk.clear_secret_info` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [params](resources--tunnel--reference--group-001.md#canonical-3112221221213033-1303222023230230-0222322021220323-2332312322123003-2200133123232101-3021120033320130-3200322033133310-2131310310120213)
- [params.ipsec](resources--tunnel--reference--group-001.md#canonical-1333220212012230-1111003122111213-2201303031102303-3212013001210022-0333212002213133-3030003012100230-3000310333301101-1011203123230131)
- [params.ipsec.ipsec_psk](resources--tunnel--reference--group-001.md#canonical-0120100121011131-0222212233122303-0033310323102030-0111122220023121-0012322203220132-3213012300321002-0300020001321011-2000312230323110)
- params.ipsec.ipsec_psk.clear_secret_info

<a id="canonical-3130332012201312-1121331233232102-2031310101003331-2213111112032210-3133003113311131-2011330320130100-0123213003011121-1301211302033203"></a>

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

<a id="canonical-0011333331032202-0302332110101231-0003320023033111-0022013100033302-1023030203321200-0122301000003013-2133131011301103-0200101022223301"></a>

### Direct properties for `params.ipsec.ipsec_psk.clear_secret_info`

<a id="canonical-1103131120220003-2131121132300232-1322223123120021-3301233123221133-0321300133123300-0102020212210030-0203302330133122-1211000302031202"></a>

#### `params.ipsec.ipsec_psk.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0000331100010220-1232122120201301-2000310023002102-2223122101110000-1013200010322322-2120212332211330-1032320233222102-0021132300030021"></a>

<a id="canonical-2100033121100003-1101200122221023-0213030100112303-1302032223101121-3300110103130120-3223000031111111-1002230312013121-2103001022312132"></a>

#### `params.ipsec.ipsec_psk.clear_secret_info.url` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0020131231323103-2233031133231022-3221122302102332-0010301321220022-1200003103313200-0000212120201011-3323033323302300-1213020310102310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- remote_ip

<a id="canonical-1102101210301330-3303010122231300-0323211303100031-3030003231221213-2022001201231223-0100202312303102-2121222003103132-0231103022201212"></a>

Type: `"object"`. single nested block, Optional.

Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP
Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map
of IP address on per ver node basis.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("endpoints",
    "ip")}
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
  "x-ves-oneof-field-type": "[\"endpoints\",\"ip\"]"
}
```

Terraform syntax:

```terraform
remote_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323333033223102-2021320120132032-2212122230233110-2123332320221323-1112123210231123-1213333020001321-3102202011313232-1331033013031233"></a>

### Direct properties for `remote_ip`

- [endpoints](resources--tunnel--reference--group-001.md#canonical-3131101033002030-1000022121022120-0013113323333100-0321112100201110-2032211213331201-1111000223013220-3300302303113101-1020222312210221): complete subsection reference.

- [ip](resources--tunnel--reference--group-001.md#canonical-2321302332112121-1213230323021130-1030011012230210-2302003120132120-0031302222212223-0133023030001032-1102010220022301-0132123301002110): complete subsection reference.

<a id="canonical-3131101033002030-1000022121022120-0013113323333100-0321112100201110-2032211213331201-1111000223013220-3300302303113101-1020222312210221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.endpoints` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0020131231323103-2233031133231022-3221122302102332-0010301321220022-1200003103313200-0000212120201011-3323033323302300-1213020310102310)
- remote_ip.endpoints

<a id="canonical-1113313210331231-3013030110330231-2201323001211123-0131131330301123-0101330003112001-3102312033002133-1023222331212323-0223132112220300"></a>

Type: `"object"`. single nested block, Optional.

Provides a map of ver node name to remote node attributes Ver node should use these attributes to
configure as remote tunnel.

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
endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112002222113102-0031000211031000-3211201102020113-3010023311311031-1101232331123210-2001320131333313-1300322220312221-2111222023033111"></a>

### Direct properties for `remote_ip.endpoints`

- [endpoints](resources--tunnel--reference--group-001.md#canonical-2200131102023311-1120313333001220-0022101232030223-3300100200331222-2231332133100300-0011012131311200-3113211213310313-1011113233033201): complete subsection reference.

<a id="canonical-2200131102023311-1120313333001220-0022101232030223-3300100200331222-2231332133100300-0011012131311200-3113211213310313-1011113233033201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.endpoints.endpoints` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0020131231323103-2233031133231022-3221122302102332-0010301321220022-1200003103313200-0000212120201011-3323033323302300-1213020310102310)
- [remote_ip.endpoints](resources--tunnel--reference--group-001.md#canonical-3131101033002030-1000022121022120-0013113323333100-0321112100201110-2032211213331201-1111000223013220-3300302303113101-1020222312210221)
- remote_ip.endpoints.endpoints

<a id="canonical-0220330101100332-0310020221231123-2032012021202333-0023120202113232-0111203201022212-0322323223102232-3233113310202201-1133010023212122"></a>

Type: `"object"`. single nested block, Optional.

Map of remote attributes to which tunnel will be established on per site node basis Every node can
have a different attributes and IP address to connect to Key is ver node name and value is Remote
node attributes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 256,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "256",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "128"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

Terraform syntax:

```terraform
endpoints {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321302332112121-1213230323021130-1030011012230210-2302003120132120-0031302222212223-0133023030001032-1102010220022301-0132123301002110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.ip` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0020131231323103-2233031133231022-3221122302102332-0010301321220022-1200003103313200-0000212120201011-3323033323302300-1213020310102310)
- remote_ip.ip

<a id="canonical-3222222110011123-0331202323320211-2223210223211200-2202331122221003-1023313130131322-0210202011222322-0322333320230001-0331003011131210"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311330203210332-0310223132021222-2032023230130010-3200310212112020-3201221101303301-2001302013330333-2121321311120222-2000103220003021"></a>

### Direct properties for `remote_ip.ip`

- [dual_stack](resources--tunnel--reference--group-001.md#canonical-0213230203000003-2112030131320032-2223233132332012-0102130123133003-3213101001120131-1112202321122013-3122131200031223-2130210323322303): complete subsection reference.

- [IPv4](resources--tunnel--reference--group-001.md#canonical-0232311001033320-3221302023010310-0230010301033330-2120122223312001-3303320321011111-1311120220323233-3002200023233330-1220200100232321): complete subsection reference.

- [IPv6](resources--tunnel--reference--group-001.md#canonical-3133331110133123-2233113021010030-0100320223013202-3232120312223210-2203200110210320-1033023312030323-0102013231220210-1011132322113201): complete subsection reference.

<a id="canonical-0213230203000003-2112030131320032-2223233132332012-0102130123133003-3213101001120131-1112202321122013-3122131200031223-2130210323322303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.ip.dual_stack` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0020131231323103-2233031133231022-3221122302102332-0010301321220022-1200003103313200-0000212120201011-3323033323302300-1213020310102310)
- [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-2321302332112121-1213230323021130-1030011012230210-2302003120132120-0031302222212223-0133023030001032-1102010220022301-0132123301002110)
- remote_ip.ip.dual_stack

<a id="canonical-3301103133321333-3203303030321100-2100011331313212-0022113201103231-3023122213233232-3111311220122121-1321022311131112-3133202311122002"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

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
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220020020221231-0332000013031001-2302203233021122-0012311033300123-2120311030303020-3332110313010000-1303310003113312-0031313223012032"></a>

### Direct properties for `remote_ip.ip.dual_stack`

- [IPv4](resources--tunnel--reference--group-001.md#canonical-2123321021002333-3132312022203020-1301032033223020-1320322002102312-1332301012120210-2222223130313212-2230312132311301-1300203211211122): complete subsection reference.

- [IPv6](resources--tunnel--reference--group-001.md#canonical-1203001020233122-1323023223003031-1201103302033313-2111131132003131-0232331002113203-2021010112230221-3210130112011213-3130021333122210): complete subsection reference.

<a id="canonical-2123321021002333-3132312022203020-1301032033223020-1320322002102312-1332301012120210-2222223130313212-2230312132311301-1300203211211122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.ip.dual_stack.ipv4` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0020131231323103-2233031133231022-3221122302102332-0010301321220022-1200003103313200-0000212120201011-3323033323302300-1213020310102310)
- [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-2321302332112121-1213230323021130-1030011012230210-2302003120132120-0031302222212223-0133023030001032-1102010220022301-0132123301002110)
- [remote_ip.ip.dual_stack](resources--tunnel--reference--group-001.md#canonical-0213230203000003-2112030131320032-2223233132332012-0102130123133003-3213101001120131-1112202321122013-3122131200031223-2130210323322303)
- remote_ip.ip.dual_stack.IPv4

<a id="canonical-0221132310021102-0233232133320121-0310003311330123-1220113111002031-0022010110013302-1032220200000330-2131032332201322-2213122101003111"></a>

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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221031302311232-0332010322110130-0210310133312030-0033303233013023-3033131033232033-1333010101033122-1100233133313322-3333223121010001"></a>

### Direct properties for `remote_ip.ip.dual_stack.ipv4`

<a id="canonical-0311302312122133-0002213132023103-2131302321330221-0013023123230113-3130130131232300-1123010032230230-3310323033331111-1132131001213132"></a>

#### `remote_ip.ip.dual_stack.ipv4.addr` property

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1203001020233122-1323023223003031-1201103302033313-2111131132003131-0232331002113203-2021010112230221-3210130112011213-3130021333122210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.ip.dual_stack.ipv6` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0020131231323103-2233031133231022-3221122302102332-0010301321220022-1200003103313200-0000212120201011-3323033323302300-1213020310102310)
- [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-2321302332112121-1213230323021130-1030011012230210-2302003120132120-0031302222212223-0133023030001032-1102010220022301-0132123301002110)
- [remote_ip.ip.dual_stack](resources--tunnel--reference--group-001.md#canonical-0213230203000003-2112030131320032-2223233132332012-0102130123133003-3213101001120131-1112202321122013-3122131200031223-2130210323322303)
- remote_ip.ip.dual_stack.IPv6

<a id="canonical-0230323232332330-3313011331321313-1131303222030203-2333123122011302-1120322012300231-0110200003030300-1030000313122032-2220321321301303"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130123000010321-2313121021133021-2133201300301312-0212131323203102-2031230312122333-1230303130220023-1012203321000311-0003231213200031"></a>

### Direct properties for `remote_ip.ip.dual_stack.ipv6`

<a id="canonical-3010233323011222-2301120013332331-0130031320001113-0120300230113103-3030322302022110-3331231333101110-3111200022222131-0111330012033031"></a>

#### `remote_ip.ip.dual_stack.ipv6.addr` property

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0232311001033320-3221302023010310-0230010301033330-2120122223312001-3303320321011111-1311120220323233-3002200023233330-1220200100232321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.ip.ipv4` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0020131231323103-2233031133231022-3221122302102332-0010301321220022-1200003103313200-0000212120201011-3323033323302300-1213020310102310)
- [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-2321302332112121-1213230323021130-1030011012230210-2302003120132120-0031302222212223-0133023030001032-1102010220022301-0132123301002110)
- remote_ip.ip.IPv4

<a id="canonical-2022223321220223-0113212330330230-2010333000223313-1123131223032312-1303112311000110-1233100302222332-1220123000312013-2330330100302032"></a>

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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210202203212032-3102333100311011-1103313133000010-0010210133130221-2311132213301331-3122323031020322-0110223020011302-3221220102132213"></a>

### Direct properties for `remote_ip.ip.ipv4`

<a id="canonical-2302303131303320-0302300030033031-1031013012200222-3031311121003221-1130331032002020-2003231103111230-3321023212222102-1310223102122103"></a>

#### `remote_ip.ip.ipv4.addr` property

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3133331110133123-2233113021010030-0100320223013202-3232120312223210-2203200110210320-1033023312030323-0102013231220210-1011132322113201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.ip.ipv6` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0020131231323103-2233031133231022-3221122302102332-0010301321220022-1200003103313200-0000212120201011-3323033323302300-1213020310102310)
- [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-2321302332112121-1213230323021130-1030011012230210-2302003120132120-0031302222212223-0133023030001032-1102010220022301-0132123301002110)
- remote_ip.ip.IPv6

<a id="canonical-1312133122302313-3023110001322232-0211222001230002-3101203231021111-3110020211220323-0213102221230102-3212311233033033-0011332000002313"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310203330021331-2202231333323032-0301122012232023-0312320323332311-0300130311221112-3002331022313313-0100232002123101-3322310131100020"></a>

### Direct properties for `remote_ip.ip.ipv6`

<a id="canonical-1020132313120222-1203003101131202-2300113310110322-3021230202010110-0321211132012132-3222202211213101-0023122312123313-1122323012130331"></a>

#### `remote_ip.ip.ipv6.addr` property

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0232112222121303-1002201122301213-3101023123123300-0132011300202221-1333302123121302-2332110031022232-0013211020002311-1323011300230131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- timeouts

<a id="canonical-2203011112121111-2102101030131031-0232212122220302-1033312300213033-1022020033022311-0111023133000222-2213300203301030-0201312121232000"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123222011321010-0130101020103021-1122330313113122-3332120123331333-3301232100332213-2322000120111210-2133301000121303-0302313101133200"></a>

### Direct properties for `timeouts`

<a id="canonical-1111032002331233-3011110202013331-2100302120201323-3322333303332330-0013332110113221-2311222011233100-2332021123001022-3022310010012200"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3003300231121032-1010333210303020-1300230311103133-1320033333203300-2222310200110222-2021201232110223-2132220131131011-1311323313300213"></a>

<a id="canonical-3320003131301033-2302022332212210-1231033100322030-0221020133020320-3221330331102112-1112333110132110-1230211212212323-0311211230311330"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3122001321021033-3321201131213102-1313303010202332-1032320021110000-1030201021012030-2030222211203031-2201113312201233-1311213103303320"></a>

<a id="canonical-1102111322202201-2321322202131223-1030333103021122-0313230021313020-3223210130203011-1231302120102113-3301332221213213-2032021213112020"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1030121211200000-1013330003323000-1101101100330112-2303110112000031-0122013132231223-3313212131313223-0012333212033201-0122321202311103"></a>

<a id="canonical-3031021202011020-0102120200331310-3021301030110101-0101102113211033-2300030200003221-2331212101221013-1222212203313213-0302202212101113"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
