---
page_title: "xcsh_tunnel reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tunnel reference."
---

# xcsh_tunnel reference

<a id="canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- Property reference

<a id="canonical-3130130001220311-0313011333013113-0111212203332321-0231233033230101-3110031110233330-1230033132322122-0202010300203110-2302220202021100"></a>

### Direct properties for `xcsh_tunnel`

<a id="canonical-1210313032130100-0300211321200322-3231132021131313-0202132331223331-0321213133133201-2220202102331031-1100112322111011-2032022123333013"></a>

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

<a id="canonical-2313322010122012-0322312130322120-1233203233021023-3031233230311312-0031033122013303-1120101230221231-3323301211031222-3002330020110221"></a>

<a id="canonical-1112112122233013-3033232013130021-1002030211013020-1230033200211013-3330300233233321-1103113223112120-3213122333212101-0300112012323220"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Tunnel.

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

<a id="canonical-2312001210102112-0211000133300121-0333033012211211-0130333011230031-2312221323010323-3313132110230230-1002123000113313-2133211232201311"></a>

<a id="canonical-2011301031333211-0223023232330312-3010312333132230-2220233021033230-0330132333110222-1131131102202130-2122000120122110-1112210210233330"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0011123122133213-3220013100110023-3233130222021030-0221322200001212-2012020021102300-1303020211302120-2013333020120111-1210323301201330"></a>

<a id="canonical-1223010331121011-1212123131122111-3033101123003000-3330233013320210-0332103222203033-0110302030213301-1222212323220313-1003330003201301"></a>

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

- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112): complete subsection reference.

<a id="canonical-1313103130001010-2102123233030210-1103203220320220-2122212210113233-1330010100313013-3103122301211000-0202202220311320-0320323223023301"></a>

<a id="canonical-2310002323113223-0202120000110000-3033033100031031-1131202232012100-2030213211002201-0200300302113210-1112103222122003-0003021223012130"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Tunnel.

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

<a id="canonical-3332230131322210-0301013311203201-3222211033200301-0012100232303311-0200133330133212-1202031002102001-1221312323122221-1212222021333211"></a>

<a id="canonical-2103110120002032-0333133133202100-1111230112130010-0010203312113103-3220321021032221-0102323313032220-3100112012220032-3000311232213220"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Tunnel exists.

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

- [params](data-sources--tunnel--reference--group-001.md#canonical-3320302133032200-2020030010332122-2202231300331030-1223313231331123-0302322303310321-0113101001321032-1203312301033232-3033013232312313): complete subsection reference.

- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-3132312023230012-3013110002122031-3132111221301201-3022331010323223-1232301300132031-2222202103210231-2231122010211301-2302103301333120): complete subsection reference.

<a id="canonical-3212203002032133-1213001023210021-2012022332212023-2033303110003110-2011210101223130-2221230332132003-0221013311202001-2002322213302102"></a>

<a id="canonical-1210020131131112-3102002311113122-1130332001122001-0211202102320120-3001213121231320-0313002333112032-1033112001022012-2200200110012220"></a>

#### `tunnel_type` property

Type: `"string"`. Computed.

\[Enum: IPSEC\_PSK|GRE\] Supported tunnel types are IPsec IPsec tunnel type with PSK GRE tunnel
type. Possible values are \`IPSEC\_PSK\`, \`GRE\`. Defaults to \`IPSEC\_PSK\`.

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

<a id="canonical-2213311202000320-2022330312023203-3213120021213013-2121300031023231-2321133002122002-2300132220222113-1323231120002122-0200310211120010"></a>

### All schema paths for `xcsh_tunnel`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--tunnel--reference--group-001.md#canonical-1210313032130100-0300211321200322-3231132021131313-0202132331223331-0321213133133201-2220202102331031-1100112322111011-2032022123333013) |
| `description` | [description](data-sources--tunnel--reference--group-001.md#canonical-2313322010122012-0322312130322120-1233203233021023-3031233230311312-0031033122013303-1120101230221231-3323301211031222-3002330020110221) |
| `id` | [ID](data-sources--tunnel--reference--group-001.md#canonical-2312001210102112-0211000133300121-0333033012211211-0130333011230031-2312221323010323-3313132110230230-1002123000113313-2133211232201311) |
| `labels` | [labels](data-sources--tunnel--reference--group-001.md#canonical-0011123122133213-3220013100110023-3233130222021030-0221322200001212-2012020021102300-1303020211302120-2013333020120111-1210323301201330) |
| `local_ip` | [local_ip](data-sources--tunnel--reference--group-001.md#canonical-2113323130303032-0012101210122220-2123001200300301-1121133223021211-2303120301312123-3332000123233323-3300230221303110-1223001021031121) |
| `local_ip.intf` | [local_ip.intf](data-sources--tunnel--reference--group-001.md#canonical-0031213112310030-1023300100111113-2333332122132102-0000223203322030-2103123111332022-0010210331321323-1323210220013202-1331203303211211) |
| `local_ip.intf.local_intf` | [local_ip.intf.local_intf](data-sources--tunnel--reference--group-001.md#canonical-2230001120202221-0111212203330311-2012112102031222-3112021033232000-2311231002302303-3323102311113031-2020330013310220-2311331301032110) |
| `local_ip.intf.local_intf.kind` | [local_ip.intf.local_intf.kind](data-sources--tunnel--reference--group-001.md#canonical-3222301220101302-3322223321210300-0200012110112022-1013221011013213-1120131132120233-1001003120201322-0211301312000232-1333012220311203) |
| `local_ip.intf.local_intf.name` | [local_ip.intf.local_intf.name](data-sources--tunnel--reference--group-001.md#canonical-3001133101113122-1020310222110130-2102101121012111-0213330020201000-1122223311001310-2303332213033213-2131230202123312-3232010001202312) |
| `local_ip.intf.local_intf.namespace` | [local_ip.intf.local_intf.namespace](data-sources--tunnel--reference--group-001.md#canonical-3132001113302132-0101311332021110-1031221233211303-1321211020231023-1131001113231321-2110113021130120-3321330012103110-2332001310213131) |
| `local_ip.intf.local_intf.tenant` | [local_ip.intf.local_intf.tenant](data-sources--tunnel--reference--group-001.md#canonical-2001103200210221-2211202321123331-3103303203113220-3122131033133231-0301302211001311-2211001002113020-1011231321330110-3030110230212101) |
| `local_ip.intf.local_intf.uid` | [local_ip.intf.local_intf.uid](data-sources--tunnel--reference--group-001.md#canonical-2303130022102210-3111311112313022-1232013100102212-0121022122031320-3333220232103212-1233022013131010-3330301112223022-0312230123212230) |
| `local_ip.ip_address` | [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-1011133210023003-0020313113020102-2021221131211300-3313300203203210-3003303212013130-3203032203131003-3032310332123201-2012001331030001) |
| `local_ip.ip_address.auto` | [local_ip.ip_address.auto](data-sources--tunnel--reference--group-001.md#canonical-3300121131211021-3202203233333023-3300313013030232-0022112212233321-3200332020323113-3312231210220021-0121332223101331-3003033131023132) |
| `local_ip.ip_address.ip_address` | [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-2232231212222020-3020312132103022-2303000030113210-0233110110322112-0201120031001323-2231122300103223-0132223002221010-0120110102033021) |
| `local_ip.ip_address.ip_address.dual_stack` | [local_ip.ip_address.ip_address.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-3111032221031332-1110102022133012-3301102222323202-1213010113231213-1001030300022013-0302333101102331-1230003303210200-2121302131000311) |
| `local_ip.ip_address.ip_address.dual_stack.ipv4` | [local_ip.ip_address.ip_address.dual_stack.ipv4](data-sources--tunnel--reference--group-001.md#canonical-0002313110013231-0013103313300010-0001301310300203-0222130103213202-1210110123221302-0313000311220331-3302030121232013-1100103110122033) |
| `local_ip.ip_address.ip_address.dual_stack.ipv4.addr` | [local_ip.ip_address.ip_address.dual_stack.ipv4.addr](data-sources--tunnel--reference--group-001.md#canonical-3013020210230333-1000100333033113-1232201121320222-0213231222123011-2002022000123300-2220322031013310-3110001020212021-2020001212303021) |
| `local_ip.ip_address.ip_address.dual_stack.ipv6` | [local_ip.ip_address.ip_address.dual_stack.ipv6](data-sources--tunnel--reference--group-001.md#canonical-2013131310000132-3300130133221100-1331201310112202-2212330011222000-1102311222030120-3131231010323003-0121313310131332-1112023321032030) |
| `local_ip.ip_address.ip_address.dual_stack.ipv6.addr` | [local_ip.ip_address.ip_address.dual_stack.ipv6.addr](data-sources--tunnel--reference--group-001.md#canonical-2003322012122212-0021102233220132-1133202032200221-1031210200002102-1331233220120022-3323012011210112-1203032303330302-3111311230310131) |
| `local_ip.ip_address.ip_address.ipv4` | [local_ip.ip_address.ip_address.ipv4](data-sources--tunnel--reference--group-001.md#canonical-0121012103111311-3300110020313000-3223113123032320-0030101130321120-3103010321223320-0331002221123321-2331332003102011-0113202310201233) |
| `local_ip.ip_address.ip_address.ipv4.addr` | [local_ip.ip_address.ip_address.ipv4.addr](data-sources--tunnel--reference--group-001.md#canonical-3322122311130322-3031030130131120-0203022102313100-0321221210012001-0201202001003333-2313302110303313-1122213212111102-0020013233231120) |
| `local_ip.ip_address.ip_address.ipv6` | [local_ip.ip_address.ip_address.ipv6](data-sources--tunnel--reference--group-001.md#canonical-1313002032210110-0300332013103230-0212112013200221-3103121223131022-3230333211032031-0033202121031030-1310203121133330-3221130100022000) |
| `local_ip.ip_address.ip_address.ipv6.addr` | [local_ip.ip_address.ip_address.ipv6.addr](data-sources--tunnel--reference--group-001.md#canonical-0320200303001210-2020201033311013-1111102010013131-3011233121012200-0012030212003131-2330031212133330-0333133223011130-0101030301222100) |
| `local_ip.ip_address.virtual_network_type` | [local_ip.ip_address.virtual_network_type](data-sources--tunnel--reference--group-001.md#canonical-2202112001100303-3013103010331111-0300132010212002-3112111223003211-3010012101332322-0101330021031210-3033031313103022-0233232121233230) |
| `local_ip.ip_address.virtual_network_type.public` | [local_ip.ip_address.virtual_network_type.public](data-sources--tunnel--reference--group-001.md#canonical-2303110121120133-2020100201312012-0023133003022322-2221212312302203-2301220221213220-2130001012001211-2101101221102111-2330202100033322) |
| `local_ip.ip_address.virtual_network_type.site_local` | [local_ip.ip_address.virtual_network_type.site_local](data-sources--tunnel--reference--group-001.md#canonical-2313333132332213-1100013232101221-2113201323321301-0302321002102221-0231100211310033-1030123232011122-0101110321330331-3330033233033310) |
| `local_ip.ip_address.virtual_network_type.site_local_inside` | [local_ip.ip_address.virtual_network_type.site_local_inside](data-sources--tunnel--reference--group-001.md#canonical-3200100102130000-2210132111000232-3223331101302102-3332020300033302-2103322332033100-2330123313300001-2101133101332031-2221021111330101) |
| `name` | [name](data-sources--tunnel--reference--group-001.md#canonical-1313103130001010-2102123233030210-1103203220320220-2122212210113233-1330010100313013-3103122301211000-0202202220311320-0320323223023301) |
| `namespace` | [namespace](data-sources--tunnel--reference--group-001.md#canonical-3332230131322210-0301013311203201-3222211033200301-0012100232303311-0200133330133212-1202031002102001-1221312323122221-1212222021333211) |
| `params` | [params](data-sources--tunnel--reference--group-001.md#canonical-0213123223111121-0101331310100111-0211020232323022-1231110133230010-0000030333122021-0102301322012303-1300100200010311-1030302030332323) |
| `params.ipsec` | [params.ipsec](data-sources--tunnel--reference--group-001.md#canonical-0022113123321111-0123122212032031-0113023303013022-1131130223003302-3232120130010122-3120201102311220-0313033303301133-0233330323020003) |
| `params.ipsec.ipsec_psk` | [params.ipsec.ipsec_psk](data-sources--tunnel--reference--group-001.md#canonical-0300130232101131-3233322112201102-2232331313002211-3212133020121311-1202133110123233-2131102210033201-3312102331222131-0221323300303022) |
| `params.ipsec.ipsec_psk.blindfold_secret_info` | [params.ipsec.ipsec_psk.blindfold_secret_info](data-sources--tunnel--reference--group-001.md#canonical-3020210101303112-2122033223111221-1333300022131310-3033310012220203-3012021123222003-3210230202121131-1201210112132310-2212003331332131) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.decryption_provider` | [params.ipsec.ipsec_psk.blindfold_secret_info.decryption_provider](data-sources--tunnel--reference--group-001.md#canonical-2212212030322231-2100100101220202-2100102131211222-0333322030233100-1032032333002300-0031121020233002-1021200110112330-3131121233031331) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.location` | [params.ipsec.ipsec_psk.blindfold_secret_info.location](data-sources--tunnel--reference--group-001.md#canonical-0003213011122113-0031200112313301-1223003033321010-2032210111222311-1220030223231133-0322302230320330-3131122001132202-2120213033030123) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.store_provider` | [params.ipsec.ipsec_psk.blindfold_secret_info.store_provider](data-sources--tunnel--reference--group-001.md#canonical-1031120311332320-1200101313320123-1021003222002103-1020333012313002-1022101202100010-0211031022121223-3223312202333320-3011223123023203) |
| `params.ipsec.ipsec_psk.clear_secret_info` | [params.ipsec.ipsec_psk.clear_secret_info](data-sources--tunnel--reference--group-001.md#canonical-3320201033123221-1100101300320021-2011220110020001-2112031311333123-0321332102311230-0220123211221113-0232132321312330-3103302103013110) |
| `params.ipsec.ipsec_psk.clear_secret_info.provider_ref` | [params.ipsec.ipsec_psk.clear_secret_info.provider_ref](data-sources--tunnel--reference--group-001.md#canonical-2200300002112032-1102003311332221-0212233312030121-3313121232201212-3123213230211131-1213133200122002-0101020020233100-1010213320030223) |
| `params.ipsec.ipsec_psk.clear_secret_info.url` | [params.ipsec.ipsec_psk.clear_secret_info.url](data-sources--tunnel--reference--group-001.md#canonical-0221201313303232-2322001203223033-3222022122321233-1331203333313222-3301203313210113-3303003133110031-1211030030221323-2223322302211031) |
| `remote_ip` | [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-0002221103121333-1301120320323321-1231210312322020-0020321102201100-1120202120010223-2131101112332021-1101310103120302-2130333031331000) |
| `remote_ip.endpoints` | [remote_ip.endpoints](data-sources--tunnel--reference--group-001.md#canonical-0231113000023103-2121222130212012-1320031010230320-0312121210112201-0311213231130100-3122302020320231-1111213022213021-2232032331302331) |
| `remote_ip.endpoints.endpoints` | [remote_ip.endpoints.endpoints](data-sources--tunnel--reference--group-001.md#canonical-1231200202332032-1303233322003223-3100002332200303-1001132303012231-0230221033003223-1003111303201013-3330202333333210-0133202210320133) |
| `remote_ip.ip` | [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-0233220113231002-3323212233130101-0220133103331100-3303130320231112-2221001320113203-0330131031223130-3322113311322330-2212231020213001) |
| `remote_ip.ip.dual_stack` | [remote_ip.ip.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-0011132113300111-2203023203200202-2300013230032132-3122211221220200-0022310213213212-2230213022111030-0222102013012330-2333001023002101) |
| `remote_ip.ip.dual_stack.ipv4` | [remote_ip.ip.dual_stack.ipv4](data-sources--tunnel--reference--group-001.md#canonical-0021123223000100-2003221130020333-0333233301302311-0221001121302320-0312213322211230-3323002331322012-2213031210323022-0223232121032320) |
| `remote_ip.ip.dual_stack.ipv4.addr` | [remote_ip.ip.dual_stack.ipv4.addr](data-sources--tunnel--reference--group-001.md#canonical-2233222112323110-1120001103211001-1111213300122003-3132233032130132-1003210321230223-0330323232033010-0330233303303230-3032003002212111) |
| `remote_ip.ip.dual_stack.ipv6` | [remote_ip.ip.dual_stack.ipv6](data-sources--tunnel--reference--group-001.md#canonical-0100133132010000-1131332101110022-3030113030021301-1233031012213223-3310332312211312-1122313323321023-2021331012302102-0131200031010023) |
| `remote_ip.ip.dual_stack.ipv6.addr` | [remote_ip.ip.dual_stack.ipv6.addr](data-sources--tunnel--reference--group-001.md#canonical-0233023111022000-2212202202123020-1013303013331132-0113103213002012-3102000013203230-2232221301132233-1231311001003300-0003101310223332) |
| `remote_ip.ip.ipv4` | [remote_ip.ip.ipv4](data-sources--tunnel--reference--group-001.md#canonical-3031231132101210-0013201010021222-2012011121201302-0333102201213113-1103132211201020-2202233332332332-1311310001320232-0023133003231202) |
| `remote_ip.ip.ipv4.addr` | [remote_ip.ip.ipv4.addr](data-sources--tunnel--reference--group-001.md#canonical-3313003231232312-0331313233132323-0331033212103002-2322003203021213-3012203301221323-2220000223023012-0323023003231023-3100333012103232) |
| `remote_ip.ip.ipv6` | [remote_ip.ip.ipv6](data-sources--tunnel--reference--group-001.md#canonical-1000220211131302-2022101110032133-0133332013222231-2200223123010311-1133201232013211-1013332011021222-3321112201100323-0020231210303103) |
| `remote_ip.ip.ipv6.addr` | [remote_ip.ip.ipv6.addr](data-sources--tunnel--reference--group-001.md#canonical-2102322210113132-0232203221102200-1133213131222123-0131031222010303-1210333320332101-0223330230300303-2212013120213133-1133232221221311) |
| `tunnel_type` | [tunnel_type](data-sources--tunnel--reference--group-001.md#canonical-3212203002032133-1213001023210021-2012022332212023-2033303110003110-2011210101223130-2221230332132003-0221013311202001-2002322213302102) |

<a id="canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- local_ip

<a id="canonical-2113323130303032-0012101210122220-2123001200300301-1121133223021211-2303120301312123-3332000123233323-3300230221303110-1223001021031121"></a>

Type: `"single"`. Computed.

Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS
available are - 1. Local Interface - Network Interface from which IP address and network will be
selected 2. IP Address - IP address and network can be configured explicitly.

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

<a id="canonical-1333200232310112-1013112002223302-3113202001111230-3231232313020101-2131021231223130-0130300232110011-2322333133133130-0220113121120313"></a>

### Direct properties for `local_ip`

- [intf](data-sources--tunnel--reference--group-001.md#canonical-0103231232320033-2333120131222023-2101231331312011-2020233233100113-0133001301010031-1030110211003132-1012301110300323-1312230201230113): complete subsection reference.

- [ip_address](data-sources--tunnel--reference--group-001.md#canonical-0220323001102331-0213311030330111-2111232213210111-1033120330031000-0211330211223312-1331320130322020-0330300311013221-0200103200320110): complete subsection reference.

<a id="canonical-0103231232320033-2333120131222023-2101231331312011-2020233233100113-0133001301010031-1030110211003132-1012301110300323-1312230201230113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.intf` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112)
- local_ip.intf

<a id="canonical-0031213112310030-1023300100111113-2333332122132102-0000223203322030-2103123111332022-0010210331321323-1323210220013202-1331203303211211"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0213011333320322-0131213133312010-2022113111103203-0301100231120201-0323300321301330-0133103220001122-3022103322133332-1232322033121100"></a>

### Direct properties for `local_ip.intf`

- [local_intf](data-sources--tunnel--reference--group-001.md#canonical-1211332200023320-0013131310010332-0023021121000032-1331033122313123-3032123001200201-2021131212111313-1211311203311321-1312121021222110): complete subsection reference.

<a id="canonical-1211332200023320-0013131310010332-0023021121000032-1331033122313123-3032123001200201-2021131212111313-1211311203311321-1312121021222110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.intf.local_intf` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112)
- [local_ip.intf](data-sources--tunnel--reference--group-001.md#canonical-0103231232320033-2333120131222023-2101231331312011-2020233233100113-0133001301010031-1030110211003132-1012301110300323-1312230201230113)
- local_ip.intf.local_intf

<a id="canonical-2230001120202221-0111212203330311-2012112102031222-3112021033232000-2311231002302303-3323102311113031-2020330013310220-2311331301032110"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0220230310001111-0121332010222101-0321031301112302-2021231030313310-2301301013011023-1123321120030223-0000332300031223-3213123302031311"></a>

### Direct properties for `local_ip.intf.local_intf`

<a id="canonical-3222301220101302-3322223321210300-0200012110112022-1013221011013213-1120131132120233-1001003120201322-0211301312000232-1333012220311203"></a>

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

<a id="canonical-3001133101113122-1020310222110130-2102101121012111-0213330020201000-1122223311001310-2303332213033213-2131230202123312-3232010001202312"></a>

<a id="canonical-0212313111321211-2202321223101121-3311110322121233-3310203313203212-2102012200212313-3033312333023221-3123121122300101-0310130022122331"></a>

#### `local_ip.intf.local_intf.name` property

Type: `"string"`. Computed.

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

<a id="canonical-3132001113302132-0101311332021110-1031221233211303-1321211020231023-1131001113231321-2110113021130120-3321330012103110-2332001310213131"></a>

<a id="canonical-3131010011103030-3233330131011032-2222333212210013-3333031312233301-1310010300123213-3010132302300131-1131300033233223-1222033320103213"></a>

#### `local_ip.intf.local_intf.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-2001103200210221-2211202321123331-3103303203113220-3122131033133231-0301302211001311-2211001002113020-1011231321330110-3030110230212101"></a>

<a id="canonical-3312220310323312-1100132011002232-3220323102010330-1310310231323120-2200321100122213-2313010312331302-0300230003210113-1313120110102000"></a>

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

<a id="canonical-2303130022102210-3111311112313022-1232013100102212-0121022122031320-3333220232103212-1233022013131010-3330301112223022-0312230123212230"></a>

<a id="canonical-1311232311021113-0321121322211322-3220132130011002-2000011310220331-1310200321212332-0230202233012312-0302212223320220-0003213232303213"></a>

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

<a id="canonical-0220323001102331-0213311030330111-2111232213210111-1033120330031000-0211330211223312-1331320130322020-0330300311013221-0200103200320110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112)
- local_ip.ip_address

<a id="canonical-1011133210023003-0020313113020102-2021221131211300-3313300203203210-3003303212013130-3203032203131003-3032310332123201-2012001331030001"></a>

Type: `"single"`. Computed.

Provides the configuration to pick up source IP and network for transporting encapsulated packet.

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

<a id="canonical-1200003111103031-2131021230201303-0110323222313303-1003101202210122-2120100112113312-1032110302303300-1330231133223313-0321132131002011"></a>

### Direct properties for `local_ip.ip_address`

- [auto](data-sources--tunnel--reference--group-001.md#canonical-1303212222101010-3010130330100323-3102220231313301-2131331303222303-2203233023212203-2102230022021230-3330011221021322-1310121023121000): complete subsection reference.

- [ip_address](data-sources--tunnel--reference--group-001.md#canonical-0320300121131123-3233122030332323-2221232210010123-1032030203023203-1100033312030023-2110312123120310-3201220101022212-3310120120000131): complete subsection reference.

- [virtual_network_type](data-sources--tunnel--reference--group-001.md#canonical-3330000213100001-2310200332122010-2031312312233110-2322231331011002-0323311022322333-1123031011101001-0000010302321030-0221121113120030): complete subsection reference.

<a id="canonical-1303212222101010-3010130330100323-3102220231313301-2131331303222303-2203233023212203-2102230022021230-3330011221021322-1310121023121000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.auto` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0220323001102331-0213311030330111-2111232213210111-1033120330031000-0211330211223312-1331320130322020-0330300311013221-0200103200320110)
- local_ip.ip_address.auto

<a id="canonical-3300121131211021-3202203233333023-3300313013030232-0022112212233321-3200332020323113-3312231210220021-0121332223101331-3003033131023132"></a>

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

<a id="canonical-0320300121131123-3233122030332323-2221232210010123-1032030203023203-1100033312030023-2110312123120310-3201220101022212-3310120120000131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.ip_address` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0220323001102331-0213311030330111-2111232213210111-1033120330031000-0211330211223312-1331320130322020-0330300311013221-0200103200320110)
- local_ip.ip_address.ip_address

<a id="canonical-2232231212222020-3020312132103022-2303000030113210-0233110110322112-0201120031001323-2231122300103223-0132223002221010-0120110102033021"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

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

<a id="canonical-1033010331322330-2032311033312301-1212022230313232-2333232300320132-0303133010122032-0023333312331013-2110232313200032-2230122002103220"></a>

### Direct properties for `local_ip.ip_address.ip_address`

- [dual_stack](data-sources--tunnel--reference--group-001.md#canonical-3003130231321020-2030123013313211-2302003030330112-2313031221232311-0331202213231103-0123200122330100-0202033131003311-0313232311123133): complete subsection reference.

- [IPv4](data-sources--tunnel--reference--group-001.md#canonical-0032111122203203-0301102312131232-3320121312000123-1001221013110213-0223122313121123-3122023122123312-3323233031121011-1222001002213021): complete subsection reference.

- [IPv6](data-sources--tunnel--reference--group-001.md#canonical-3010231203230221-1232203103301011-3001010331003011-0320131123302302-2311322310331330-3303123120112001-2002231232133013-2230322213031211): complete subsection reference.

<a id="canonical-3003130231321020-2030123013313211-2302003030330112-2313031221232311-0331202213231103-0123200122330100-0202033131003311-0313232311123133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.ip_address.dual_stack` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0220323001102331-0213311030330111-2111232213210111-1033120330031000-0211330211223312-1331320130322020-0330300311013221-0200103200320110)
- [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0320300121131123-3233122030332323-2221232210010123-1032030203023203-1100033312030023-2110312123120310-3201220101022212-3310120120000131)
- local_ip.ip_address.ip_address.dual_stack

<a id="canonical-3111032221031332-1110102022133012-3301102222323202-1213010113231213-1001030300022013-0302333101102331-1230003303210200-2121302131000311"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3213131000031202-2232001121132223-1313322331200202-0113133321002223-3030201311333322-3203300132100023-2020231203131333-0122310333122021"></a>

### Direct properties for `local_ip.ip_address.ip_address.dual_stack`

- [IPv4](data-sources--tunnel--reference--group-001.md#canonical-1220111100011221-0012000010032330-1003020133112210-2200310113213213-1320011210122032-0322330031313132-1210111031230320-2021311111020220): complete subsection reference.

- [IPv6](data-sources--tunnel--reference--group-001.md#canonical-3122203321330212-3023020120101123-0033323113302201-3303301132112232-1203232323122133-0111210211010132-1302032030310310-2103213211203202): complete subsection reference.

<a id="canonical-1220111100011221-0012000010032330-1003020133112210-2200310113213213-1320011210122032-0322330031313132-1210111031230320-2021311111020220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.ip_address.dual_stack.ipv4` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0220323001102331-0213311030330111-2111232213210111-1033120330031000-0211330211223312-1331320130322020-0330300311013221-0200103200320110)
- [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0320300121131123-3233122030332323-2221232210010123-1032030203023203-1100033312030023-2110312123120310-3201220101022212-3310120120000131)
- [local_ip.ip_address.ip_address.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-3003130231321020-2030123013313211-2302003030330112-2313031221232311-0331202213231103-0123200122330100-0202033131003311-0313232311123133)
- local_ip.ip_address.ip_address.dual_stack.IPv4

<a id="canonical-0002313110013231-0013103313300010-0001301310300203-0222130103213202-1210110123221302-0313000311220331-3302030121232013-1100103110122033"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0001332213131320-1111210311223223-1132132311031113-3020333203301210-0211300010303023-3320122130112320-2201320332233011-2030300312213101"></a>

### Direct properties for `local_ip.ip_address.ip_address.dual_stack.ipv4`

<a id="canonical-3013020210230333-1000100333033113-1232201121320222-0213231222123011-2002022000123300-2220322031013310-3110001020212021-2020001212303021"></a>

#### `local_ip.ip_address.ip_address.dual_stack.ipv4.addr` property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-3122203321330212-3023020120101123-0033323113302201-3303301132112232-1203232323122133-0111210211010132-1302032030310310-2103213211203202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.ip_address.dual_stack.ipv6` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0220323001102331-0213311030330111-2111232213210111-1033120330031000-0211330211223312-1331320130322020-0330300311013221-0200103200320110)
- [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0320300121131123-3233122030332323-2221232210010123-1032030203023203-1100033312030023-2110312123120310-3201220101022212-3310120120000131)
- [local_ip.ip_address.ip_address.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-3003130231321020-2030123013313211-2302003030330112-2313031221232311-0331202213231103-0123200122330100-0202033131003311-0313232311123133)
- local_ip.ip_address.ip_address.dual_stack.IPv6

<a id="canonical-2013131310000132-3300130133221100-1331201310112202-2212330011222000-1102311222030120-3131231010323003-0121313310131332-1112023321032030"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0020200110231003-0020010300221023-1313312031000211-3310333110100113-0012231313000131-2123000331203230-2331002011033322-3231330133300002"></a>

### Direct properties for `local_ip.ip_address.ip_address.dual_stack.ipv6`

<a id="canonical-2003322012122212-0021102233220132-1133202032200221-1031210200002102-1331233220120022-3323012011210112-1203032303330302-3111311230310131"></a>

#### `local_ip.ip_address.ip_address.dual_stack.ipv6.addr` property

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0032111122203203-0301102312131232-3320121312000123-1001221013110213-0223122313121123-3122023122123312-3323233031121011-1222001002213021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.ip_address.ipv4` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0220323001102331-0213311030330111-2111232213210111-1033120330031000-0211330211223312-1331320130322020-0330300311013221-0200103200320110)
- [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0320300121131123-3233122030332323-2221232210010123-1032030203023203-1100033312030023-2110312123120310-3201220101022212-3310120120000131)
- local_ip.ip_address.ip_address.IPv4

<a id="canonical-0121012103111311-3300110020313000-3223113123032320-0030101130321120-3103010321223320-0331002221123321-2331332003102011-0113202310201233"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2120301031200332-1101312313303030-1233210330112231-3030103003000321-1033022220320223-3202320021103120-2012301010210032-1122223310110323"></a>

### Direct properties for `local_ip.ip_address.ip_address.ipv4`

<a id="canonical-3322122311130322-3031030130131120-0203022102313100-0321221210012001-0201202001003333-2313302110303313-1122213212111102-0020013233231120"></a>

#### `local_ip.ip_address.ip_address.ipv4.addr` property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-3010231203230221-1232203103301011-3001010331003011-0320131123302302-2311322310331330-3303123120112001-2002231232133013-2230322213031211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.ip_address.ipv6` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0220323001102331-0213311030330111-2111232213210111-1033120330031000-0211330211223312-1331320130322020-0330300311013221-0200103200320110)
- [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0320300121131123-3233122030332323-2221232210010123-1032030203023203-1100033312030023-2110312123120310-3201220101022212-3310120120000131)
- local_ip.ip_address.ip_address.IPv6

<a id="canonical-1313002032210110-0300332013103230-0212112013200221-3103121223131022-3230333211032031-0033202121031030-1310203121133330-3221130100022000"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3320301122003121-0012232303300301-1101220232133301-3330032302111212-2230002230310011-0303230231001230-1233311321213213-2111121132132100"></a>

### Direct properties for `local_ip.ip_address.ip_address.ipv6`

<a id="canonical-0320200303001210-2020201033311013-1111102010013131-3011233121012200-0012030212003131-2330031212133330-0333133223011130-0101030301222100"></a>

#### `local_ip.ip_address.ip_address.ipv6.addr` property

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3330000213100001-2310200332122010-2031312312233110-2322231331011002-0323311022322333-1123031011101001-0000010302321030-0221121113120030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.virtual_network_type` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0220323001102331-0213311030330111-2111232213210111-1033120330031000-0211330211223312-1331320130322020-0330300311013221-0200103200320110)
- local_ip.ip_address.virtual_network_type

<a id="canonical-2202112001100303-3013103010331111-0300132010212002-3112111223003211-3010012101332322-0101330021031210-3033031313103022-0233232121233230"></a>

Type: `"single"`. Computed.

Different types of virtual networks understood by the system.

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

<a id="canonical-1002011311332030-2000333120321012-0000132332302003-1121022102133100-3302121021323110-0001130201030002-1000303110301001-2213220030112013"></a>

### Direct properties for `local_ip.ip_address.virtual_network_type`

- [public](data-sources--tunnel--reference--group-001.md#canonical-1011331311112300-0300223112201200-0231122021202302-2130032312103221-3012032003313311-2132023212303121-3013033002122202-0202020321213332): complete subsection reference.

- [site_local](data-sources--tunnel--reference--group-001.md#canonical-0330130312122223-1210312012022212-2312030100210111-1133102110310000-3231233301123220-2333030010101332-1020230310012011-0310200112331213): complete subsection reference.

- [site_local_inside](data-sources--tunnel--reference--group-001.md#canonical-2002310311201302-2233010123321210-1011030103302203-2103202213202131-1311202233223221-3100031300120112-1120330103212122-0321002131122020): complete subsection reference.

<a id="canonical-1011331311112300-0300223112201200-0231122021202302-2130032312103221-3012032003313311-2132023212303121-3013033002122202-0202020321213332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.virtual_network_type.public` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0220323001102331-0213311030330111-2111232213210111-1033120330031000-0211330211223312-1331320130322020-0330300311013221-0200103200320110)
- [local_ip.ip_address.virtual_network_type](data-sources--tunnel--reference--group-001.md#canonical-3330000213100001-2310200332122010-2031312312233110-2322231331011002-0323311022322333-1123031011101001-0000010302321030-0221121113120030)
- local_ip.ip_address.virtual_network_type.public

<a id="canonical-2303110121120133-2020100201312012-0023133003022322-2221212312302203-2301220221213220-2130001012001211-2101101221102111-2330202100033322"></a>

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

<a id="canonical-0330130312122223-1210312012022212-2312030100210111-1133102110310000-3231233301123220-2333030010101332-1020230310012011-0310200112331213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.virtual_network_type.site_local` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0220323001102331-0213311030330111-2111232213210111-1033120330031000-0211330211223312-1331320130322020-0330300311013221-0200103200320110)
- [local_ip.ip_address.virtual_network_type](data-sources--tunnel--reference--group-001.md#canonical-3330000213100001-2310200332122010-2031312312233110-2322231331011002-0323311022322333-1123031011101001-0000010302321030-0221121113120030)
- local_ip.ip_address.virtual_network_type.site_local

<a id="canonical-2313333132332213-1100013232101221-2113201323321301-0302321002102221-0231100211310033-1030123232011122-0101110321330331-3330033233033310"></a>

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

<a id="canonical-2002310311201302-2233010123321210-1011030103302203-2103202213202131-1311202233223221-3100031300120112-1120330103212122-0321002131122020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_ip.ip_address.virtual_network_type.site_local_inside` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-0220323001102331-0213311030330111-2111232213210111-1033120330031000-0211330211223312-1331320130322020-0330300311013221-0200103200320110)
- [local_ip.ip_address.virtual_network_type](data-sources--tunnel--reference--group-001.md#canonical-3330000213100001-2310200332122010-2031312312233110-2322231331011002-0323311022322333-1123031011101001-0000010302321030-0221121113120030)
- local_ip.ip_address.virtual_network_type.site_local_inside

<a id="canonical-3200100102130000-2210132111000232-3223331101302102-3332020300033302-2103322332033100-2330123313300001-2101133101332031-2221021111330101"></a>

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

<a id="canonical-3320302133032200-2020030010332122-2202231300331030-1223313231331123-0302322303310321-0113101001321032-1203312301033232-3033013232312313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `params` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- params

<a id="canonical-0213123223111121-0101331310100111-0211020232323022-1231110133230010-0000030333122021-0102301322012303-1300100200010311-1030302030332323"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1130222200101021-3010221013321102-0213102003113323-1212313131022031-0232200220313303-2220212030123231-1303303103120332-3310112000123320"></a>

### Direct properties for `params`

- [ipsec](data-sources--tunnel--reference--group-001.md#canonical-0030323213022232-2312233320100221-2220112221301302-0010102320222200-0302013300201022-0133213203111210-1313323212113030-1102320330331201): complete subsection reference.

<a id="canonical-0030323213022232-2312233320100221-2220112221301302-0010102320222200-0302013300201022-0133213203111210-1313323212113030-1102320330331201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `params.ipsec` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [params](data-sources--tunnel--reference--group-001.md#canonical-3320302133032200-2020030010332122-2202231300331030-1223313231331123-0302322303310321-0113101001321032-1203312301033232-3033013232312313)
- params.ipsec

<a id="canonical-0022113123321111-0123122212032031-0113023303013022-1131130223003302-3232120130010122-3120201102311220-0313033303301133-0233330323020003"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1211212230302123-0302323222122131-0131022023111232-3111001000330023-1032223301323212-2211132001320231-2202133221322320-2230011120302323"></a>

### Direct properties for `params.ipsec`

- [ipsec_psk](data-sources--tunnel--reference--group-001.md#canonical-1022210110031320-3302000230230122-0221122220020322-3311033123210001-1101002303020330-3121102321331322-3013000332221322-1123330312311021): complete subsection reference.

<a id="canonical-1022210110031320-3302000230230122-0221122220020322-3311033123210001-1101002303020330-3121102321331322-3013000332221322-1123330312311021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `params.ipsec.ipsec_psk` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [params](data-sources--tunnel--reference--group-001.md#canonical-3320302133032200-2020030010332122-2202231300331030-1223313231331123-0302322303310321-0113101001321032-1203312301033232-3033013232312313)
- [params.ipsec](data-sources--tunnel--reference--group-001.md#canonical-0030323213022232-2312233320100221-2220112221301302-0010102320222200-0302013300201022-0133213203111210-1313323212113030-1102320330331201)
- params.ipsec.ipsec_psk

<a id="canonical-0300130232101131-3233322112201102-2232331313002211-3212133020121311-1202133110123233-2131102210033201-3312102331222131-0221323300303022"></a>

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

<a id="canonical-2321100232311222-3121023022223013-1213033100102321-2200333133121330-2333203130020010-0130231313133131-3033230303330322-2223230310233201"></a>

### Direct properties for `params.ipsec.ipsec_psk`

- [blindfold_secret_info](data-sources--tunnel--reference--group-001.md#canonical-0030022211030133-1113330010002233-1202301322102213-2022113213320030-2212011123111221-3023113302302002-3223301130023120-3203330010230123): complete subsection reference.

- [clear_secret_info](data-sources--tunnel--reference--group-001.md#canonical-1333031020101320-3213123011321322-3022223022131301-2331133012310012-2133312221212321-2010310320233222-1332201032302210-0022102221230201): complete subsection reference.

<a id="canonical-0030022211030133-1113330010002233-1202301322102213-2022113213320030-2212011123111221-3023113302302002-3223301130023120-3203330010230123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `params.ipsec.ipsec_psk.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [params](data-sources--tunnel--reference--group-001.md#canonical-3320302133032200-2020030010332122-2202231300331030-1223313231331123-0302322303310321-0113101001321032-1203312301033232-3033013232312313)
- [params.ipsec](data-sources--tunnel--reference--group-001.md#canonical-0030323213022232-2312233320100221-2220112221301302-0010102320222200-0302013300201022-0133213203111210-1313323212113030-1102320330331201)
- [params.ipsec.ipsec_psk](data-sources--tunnel--reference--group-001.md#canonical-1022210110031320-3302000230230122-0221122220020322-3311033123210001-1101002303020330-3121102321331322-3013000332221322-1123330312311021)
- params.ipsec.ipsec_psk.blindfold_secret_info

<a id="canonical-3020210101303112-2122033223111221-1333300022131310-3033310012220203-3012021123222003-3210230202121131-1201210112132310-2212003331332131"></a>

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

<a id="canonical-2102233310010220-3122333021031312-2132123300203112-1331210000200331-1133023200300100-3022003313130332-2032320201302202-2320220130201203"></a>

### Direct properties for `params.ipsec.ipsec_psk.blindfold_secret_info`

<a id="canonical-2212212030322231-2100100101220202-2100102131211222-0333322030233100-1032032333002300-0031121020233002-1021200110112330-3131121233031331"></a>

#### `params.ipsec.ipsec_psk.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0003213011122113-0031200112313301-1223003033321010-2032210111222311-1220030223231133-0322302230320330-3131122001132202-2120213033030123"></a>

<a id="canonical-1211020021300330-0300312133101312-3220223123331002-2300312002011100-3313311120003110-2100013023200332-3203232222332231-2133111232030130"></a>

#### `params.ipsec.ipsec_psk.blindfold_secret_info.location` property

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

<a id="canonical-1031120311332320-1200101313320123-1021003222002103-1020333012313002-1022101202100010-0211031022121223-3223312202333320-3011223123023203"></a>

<a id="canonical-0120311322013222-1230123330232201-2122113032030123-0232101031221033-1111322121211313-0210031223301131-1101301002302212-2101203002221021"></a>

#### `params.ipsec.ipsec_psk.blindfold_secret_info.store_provider` property

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

<a id="canonical-1333031020101320-3213123011321322-3022223022131301-2331133012310012-2133312221212321-2010310320233222-1332201032302210-0022102221230201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `params.ipsec.ipsec_psk.clear_secret_info` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [params](data-sources--tunnel--reference--group-001.md#canonical-3320302133032200-2020030010332122-2202231300331030-1223313231331123-0302322303310321-0113101001321032-1203312301033232-3033013232312313)
- [params.ipsec](data-sources--tunnel--reference--group-001.md#canonical-0030323213022232-2312233320100221-2220112221301302-0010102320222200-0302013300201022-0133213203111210-1313323212113030-1102320330331201)
- [params.ipsec.ipsec_psk](data-sources--tunnel--reference--group-001.md#canonical-1022210110031320-3302000230230122-0221122220020322-3311033123210001-1101002303020330-3121102321331322-3013000332221322-1123330312311021)
- params.ipsec.ipsec_psk.clear_secret_info

<a id="canonical-3320201033123221-1100101300320021-2011220110020001-2112031311333123-0321332102311230-0220123211221113-0232132321312330-3103302103013110"></a>

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

<a id="canonical-0103103002000332-1112223120312101-0312300013102102-0322121001203122-0111031303101021-2021201211002111-2121221113310233-2112122332123312"></a>

### Direct properties for `params.ipsec.ipsec_psk.clear_secret_info`

<a id="canonical-2200300002112032-1102003311332221-0212233312030121-3313121232201212-3123213230211131-1213133200122002-0101020020233100-1010213320030223"></a>

#### `params.ipsec.ipsec_psk.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0221201313303232-2322001203223033-3222022122321233-1331203333313222-3301203313210113-3303003133110031-1211030030221323-2223322302211031"></a>

<a id="canonical-2122013320132010-1213111023222230-1301003113002230-0020330203202331-3231123212001102-2122320313013330-1133010230120232-1020002231210113"></a>

#### `params.ipsec.ipsec_psk.clear_secret_info.url` property

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

<a id="canonical-3132312023230012-3013110002122031-3132111221301201-3022331010323223-1232301300132031-2222202103210231-2231122010211301-2302103301333120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- remote_ip

<a id="canonical-0002221103121333-1301120320323321-1231210312322020-0020321102201100-1120202120010223-2131101112332021-1101310103120302-2130333031331000"></a>

Type: `"single"`. Computed.

Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP
Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map
of IP address on per ver node basis.

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

<a id="canonical-1300103200301111-1013101203111113-2020212010000131-1233202022212031-2031212001220033-2313211323031220-1023133230321102-0130203223102123"></a>

### Direct properties for `remote_ip`

- [endpoints](data-sources--tunnel--reference--group-001.md#canonical-3003001003012231-2020031310333033-1120113023030332-2131213220100033-3331220330322202-1332221101013030-1121210230102320-0332331033123313): complete subsection reference.

- [ip](data-sources--tunnel--reference--group-001.md#canonical-0100022122121121-3310333301331022-0111103001223122-2110321310213003-2301030200221301-1101031131133022-0321332003031133-2321211011212332): complete subsection reference.

<a id="canonical-3003001003012231-2020031310333033-1120113023030332-2131213220100033-3331220330322202-1332221101013030-1121210230102320-0332331033123313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.endpoints` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-3132312023230012-3013110002122031-3132111221301201-3022331010323223-1232301300132031-2222202103210231-2231122010211301-2302103301333120)
- remote_ip.endpoints

<a id="canonical-0231113000023103-2121222130212012-1320031010230320-0312121210112201-0311213231130100-3122302020320231-1111213022213021-2232032331302331"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2022121230022111-2102132201200212-1113211100020313-3123322011102300-2123032200000100-3233210200001130-0222203011300331-1133323013110331"></a>

### Direct properties for `remote_ip.endpoints`

- [endpoints](data-sources--tunnel--reference--group-001.md#canonical-0001021301213332-0110321233030033-3003203203122222-1221300300330300-3321031210001310-3303113213022033-1121323231023320-0330333210330033): complete subsection reference.

<a id="canonical-0001021301213332-0110321233030033-3003203203122222-1221300300330300-3321031210001310-3303113213022033-1121323231023320-0330333210330033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.endpoints.endpoints` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-3132312023230012-3013110002122031-3132111221301201-3022331010323223-1232301300132031-2222202103210231-2231122010211301-2302103301333120)
- [remote_ip.endpoints](data-sources--tunnel--reference--group-001.md#canonical-3003001003012231-2020031310333033-1120113023030332-2131213220100033-3331220330322202-1332221101013030-1121210230102320-0332331033123313)
- remote_ip.endpoints.endpoints

<a id="canonical-1231200202332032-1303233322003223-3100002332200303-1001132303012231-0230221033003223-1003111303201013-3330202333333210-0133202210320133"></a>

Type: `"single"`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100022122121121-3310333301331022-0111103001223122-2110321310213003-2301030200221301-1101031131133022-0321332003031133-2321211011212332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.ip` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-3132312023230012-3013110002122031-3132111221301201-3022331010323223-1232301300132031-2222202103210231-2231122010211301-2302103301333120)
- remote_ip.ip

<a id="canonical-0233220113231002-3323212233130101-0220133103331100-3303130320231112-2221001320113203-0330131031223130-3322113311322330-2212231020213001"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

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

<a id="canonical-2031311331201300-0203323021122130-1301201112213103-1333102333213031-3133303222022310-3032211221302122-0202320110021003-0133312121210213"></a>

### Direct properties for `remote_ip.ip`

- [dual_stack](data-sources--tunnel--reference--group-001.md#canonical-1120111111010222-2003123012021023-3201112202311112-2312233301113303-0022211331010123-0323003013210130-3312120123003223-2201300212031121): complete subsection reference.

- [IPv4](data-sources--tunnel--reference--group-001.md#canonical-3110233220123120-2201113132033203-3301120103133021-2000121133221013-1131100110221232-2300231321113120-1203120102313010-3013331101110223): complete subsection reference.

- [IPv6](data-sources--tunnel--reference--group-001.md#canonical-1022222121030033-2223121110330010-3120302333000231-2210120020223331-2222220003010010-0231102000311313-2213201321323023-0120301100230101): complete subsection reference.

<a id="canonical-1120111111010222-2003123012021023-3201112202311112-2312233301113303-0022211331010123-0323003013210130-3312120123003223-2201300212031121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.ip.dual_stack` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-3132312023230012-3013110002122031-3132111221301201-3022331010323223-1232301300132031-2222202103210231-2231122010211301-2302103301333120)
- [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-0100022122121121-3310333301331022-0111103001223122-2110321310213003-2301030200221301-1101031131133022-0321332003031133-2321211011212332)
- remote_ip.ip.dual_stack

<a id="canonical-0011132113300111-2203023203200202-2300013230032132-3122211221220200-0022310213213212-2230213022111030-0222102013012330-2333001023002101"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3312131232202311-1031003223103201-0302320032100031-1031031012123023-0001210233212011-2323033212102231-2220200313132222-1010203101102223"></a>

### Direct properties for `remote_ip.ip.dual_stack`

- [IPv4](data-sources--tunnel--reference--group-001.md#canonical-0203232131120301-1322213203122221-0133120231010321-1233213333012001-3310131001000230-3230113011311133-0210200011002321-2211323330321120): complete subsection reference.

- [IPv6](data-sources--tunnel--reference--group-001.md#canonical-0332102200022230-0300330220221132-3112030131322310-3200230112122133-3021033223111301-1210320113300323-2033333112011331-2103330122313332): complete subsection reference.

<a id="canonical-0203232131120301-1322213203122221-0133120231010321-1233213333012001-3310131001000230-3230113011311133-0210200011002321-2211323330321120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.ip.dual_stack.ipv4` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-3132312023230012-3013110002122031-3132111221301201-3022331010323223-1232301300132031-2222202103210231-2231122010211301-2302103301333120)
- [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-0100022122121121-3310333301331022-0111103001223122-2110321310213003-2301030200221301-1101031131133022-0321332003031133-2321211011212332)
- [remote_ip.ip.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-1120111111010222-2003123012021023-3201112202311112-2312233301113303-0022211331010123-0323003013210130-3312120123003223-2201300212031121)
- remote_ip.ip.dual_stack.IPv4

<a id="canonical-0021123223000100-2003221130020333-0333233301302311-0221001121302320-0312213322211230-3323002331322012-2213031210323022-0223232121032320"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0201212121322121-0311223330001121-2022322332332013-1331021223110102-3132230331000303-3001323330312133-1001001112213301-2303103011303033"></a>

### Direct properties for `remote_ip.ip.dual_stack.ipv4`

<a id="canonical-2233222112323110-1120001103211001-1111213300122003-3132233032130132-1003210321230223-0330323232033010-0330233303303230-3032003002212111"></a>

#### `remote_ip.ip.dual_stack.ipv4.addr` property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-0332102200022230-0300330220221132-3112030131322310-3200230112122133-3021033223111301-1210320113300323-2033333112011331-2103330122313332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.ip.dual_stack.ipv6` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-3132312023230012-3013110002122031-3132111221301201-3022331010323223-1232301300132031-2222202103210231-2231122010211301-2302103301333120)
- [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-0100022122121121-3310333301331022-0111103001223122-2110321310213003-2301030200221301-1101031131133022-0321332003031133-2321211011212332)
- [remote_ip.ip.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-1120111111010222-2003123012021023-3201112202311112-2312233301113303-0022211331010123-0323003013210130-3312120123003223-2201300212031121)
- remote_ip.ip.dual_stack.IPv6

<a id="canonical-0100133132010000-1131332101110022-3030113030021301-1233031012213223-3310332312211312-1122313323321023-2021331012302102-0131200031010023"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1121332001121313-1122212302010230-1321201012210223-1023013120210013-3311133132022121-1133303110302010-1320330000120032-2211122330232221"></a>

### Direct properties for `remote_ip.ip.dual_stack.ipv6`

<a id="canonical-0233023111022000-2212202202123020-1013303013331132-0113103213002012-3102000013203230-2232221301132233-1231311001003300-0003101310223332"></a>

#### `remote_ip.ip.dual_stack.ipv6.addr` property

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3110233220123120-2201113132033203-3301120103133021-2000121133221013-1131100110221232-2300231321113120-1203120102313010-3013331101110223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.ip.ipv4` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-3132312023230012-3013110002122031-3132111221301201-3022331010323223-1232301300132031-2222202103210231-2231122010211301-2302103301333120)
- [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-0100022122121121-3310333301331022-0111103001223122-2110321310213003-2301030200221301-1101031131133022-0321332003031133-2321211011212332)
- remote_ip.ip.IPv4

<a id="canonical-3031231132101210-0013201010021222-2012011121201302-0333102201213113-1103132211201020-2202233332332332-1311310001320232-0023133003231202"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2022311301321322-3232221212211001-1030021331223330-3121321331023211-1211312202012203-0011302030112110-3202300100300313-3131200331212233"></a>

### Direct properties for `remote_ip.ip.ipv4`

<a id="canonical-3313003231232312-0331313233132323-0331033212103002-2322003203021213-3012203301221323-2220000223023012-0323023003231023-3100333012103232"></a>

#### `remote_ip.ip.ipv4.addr` property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-1022222121030033-2223121110330010-3120302333000231-2210120020223331-2222220003010010-0231102000311313-2213201321323023-0120301100230101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `remote_ip.ip.ipv6` properties

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-3132312023230012-3013110002122031-3132111221301201-3022331010323223-1232301300132031-2222202103210231-2231122010211301-2302103301333120)
- [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-0100022122121121-3310333301331022-0111103001223122-2110321310213003-2301030200221301-1101031131133022-0321332003031133-2321211011212332)
- remote_ip.ip.IPv6

<a id="canonical-1000220211131302-2022101110032133-0133332013222231-2200223123010311-1133201232013211-1013332011021222-3321112201100323-0020231210303103"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0133320213201310-2032312032030123-0122000022202221-1133233120323131-0232330330311132-3300300122002133-2003300110332002-1301031130001121"></a>

### Direct properties for `remote_ip.ip.ipv6`

<a id="canonical-2102322210113132-0232203221102200-1133213131222123-0131031222010303-1210333320332101-0223330230300303-2212013120213133-1133232221221311"></a>

#### `remote_ip.ip.ipv6.addr` property

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```
